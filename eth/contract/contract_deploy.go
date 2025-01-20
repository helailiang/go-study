package main

/**
 * @Description $
 * @Author hll
 * @Date 2024/8/2 下午5:42
 **/
import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/rpc"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	store "contract/store" // for demo
)

func main() {
	//var token = "Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiJ9.eyJwZXJtaXNzaW9ucyI6WyJuZXQ6KiIsImV0aDoqIiwiYWRtaW46KiJdLCJ1c2VybmFtZSI6InA2bWIwaSIsImlhdCI6MTcyMzQ0OTEwNywiZXhwIjoxNzIzNDQ5NDA3fQ.wQOG1THGTJTkQ5Eue_RvrtozKNbFzNfWFH3F_x7k2KHy_6K8ouN0bqqc5BCkYoWzoYVLd6jSpzcjbN_5XY0-gdkpzlVIbvaGWrEJK_LX4cDFVkHHjByqpyQ5JKyT-NMO4XPAlAdJceknS1ITyjaxXm4ii1PHBI2Ti7RIa2g2_QwaEEnOFe7Odlf35Y1w8Ls03pl9ghMMTlDfpGp_ygMpG5r49Bdlwjm5ymrO6wewp5QcCujV2uMiVm_hi1J12PU0OfKK9blk_qICk5s52cYKBLI32c9zn7CsEURhvUjGgOVmg1KEJ3KXwy4rlH6ZdMe2Yh7SJRCInfV13A9CG1ly4Q"
	rc, err := rpc.DialOptions(context.Background(), "https://holesky.infura.io/v3/d90449f2a78842d6a3e39aed7c444ac7") // rpc.WithHeader("Authorization", token)
	if err != nil {
		log.Fatal(fmt.Sprintf("构建Besu客户端失败，error: %s", err))
	}
	//rc.SetHeader("Authorization", token)

	client := ethclient.NewClient(rc)

	//client, err := ethclient.Dial("https://ethereum-holesky.publicnode.com")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//274cfd629a3e82a786df5df3266c615b7bdbecd6fb8bdff2d9290d8e6cdf0635
	//besu 94b7c07c0ab8fa6d13ce1d0eaffba86b6cf9c033100072a20a2029c56113f819
	privateKey, err := crypto.HexToECDSA("274cfd629a3e82a786df5df3266c615b7bdbecd6fb8bdff2d9290d8e6cdf0635")
	if err != nil {
		log.Fatal("HexToECDSA==>", err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal("获取chainID失败：", err)
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal("PendingNonceAt==>", err)
	}
	fmt.Println("fromAddress==>", fromAddress)

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal("gasPrice==>", err)
	}

	//auth := bind.NewKeyedTransactor(privateKey)
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		log.Fatal("auth：", err)
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0) // in wei
	//auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice
	//auth.GasPrice = big.NewInt(0)
	auth.NoSend = true
	//input := "1.0"

	address, tx, instance, err := store.DeployStore(auth, client)
	if err != nil {
		log.Fatal("DeployStore==>", err)
	}
	fmt.Println("signTx==>", tx)
	txbin, _ := tx.MarshalBinary()
	fmt.Println("signTx txbin==>", string(txbin))
	txjson, _ := tx.MarshalJSON()
	fmt.Println("signTx txjson==>", string(txjson))
	fmt.Println("contract address==>", address.Hex()) // 0x147B8eb97fD247D06C4006D269c90C1908Fb5D54
	fmt.Println("交易hash==>", tx.Hash().Hex())         // 0xdae8ba5444eefdc99f4d45cd0c4f24056cba6a02cefbf78066ef9f4188ff7dc0

	_ = instance

	address1 := common.HexToAddress("0x147B8eb97fD247D06C4006D269c90C1908Fb5D54")
	instance, err = store.NewStore(address1, client)
	//instance.Store()
}

//contract address==> 0x5E5A6aef275ecf26a7eFb558E1485bfF9D41d425
//交易hash==> 0xd74726e2ee4b632d5efe79d1915d0edd5cc1eb3d2412bded8ab22b2685c0531a
