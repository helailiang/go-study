package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"log"
	"math/big"
	"newstudy/eth/use/store"
)

/**
 * @Description $
 * @Author hll
 * @Date 2024/8/7 下午5:35
 **/

func main() {
	from := common.HexToAddress("0xce552f976c89e3d3310e37789d0c7e1afda85837")
	add := crypto.CreateAddress(from, 0)
	fmt.Println(add)
	var signedTx types.Transaction
	paload := `{
    "type": "0x0",
    "chainId": "0x13882",
    "nonce": "0x2",
    "to": null,
    "gas": "0x19a79",
    "gasPrice": "0x737be7600",
    "maxPriorityFeePerGas": null,
    "maxFeePerGas": null,
    "value": "0x0",
    "input": "0x608060405234801561001057600080fd5b5060ec8061001f6000396000f3fe6080604052600436106049576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff1680632e64cec114604e5780636057361d146076575b600080fd5b348015605957600080fd5b50606060ad565b6040518082815260200191505060405180910390f35b348015608157600080fd5b5060ab60048036036020811015609657600080fd5b810190808035906020019092919050505060b6565b005b60008054905090565b806000819055505056fea165627a7a72305820fecc517871161a44b8e811bcf92c0defb3ecbc3d9f889b6e101359132d99ac420029",
    "v": "0x27128",
    "r": "0xfda6608b3a31bc6066ba917f44869d33a2171f77efd8836744a9875d0e5a6ef1",
    "s": "0x48b5482dea915bfcd21a828623774e1c7587db2b12fb61dee42df3edbbec5cab",
    "hash": "0x8a480b1d6fb11a4b3aea9014269036dfaffc2df5ba09ddfb2881ee9c972617cd"
}`
	err := signedTx.UnmarshalJSON([]byte(paload))
	if err != nil {
		log.Fatal("", err)
	}
	fmt.Println("sign===>", signedTx.ChainId())
	fmt.Println("--------------------------------")
	send()
}

func send() {
	privateKey, err := crypto.HexToECDSA("274cfd629a3e82a786df5df3266c615b7bdbecd6fb8bdff2d9290d8e6cdf0635")
	if err != nil {
		log.Fatal("HexToECDSA==>", err)
	}
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(1))
	if err != nil {
		log.Fatal("auth：", err)
	}
	auth.Nonce = big.NewInt(int64(5))
	auth.Value = big.NewInt(0) // in wei
	auth.GasPrice = big.NewInt(0)
	auth.NoSend = true
	_, tx, _, err := store.DeployStore(auth, nil)
	txjson, _ := tx.MarshalJSON()
	fmt.Println("signTx txjson==>", string(txjson))
}
