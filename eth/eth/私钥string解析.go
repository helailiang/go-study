/*
@Created by : 2021/6/28 13:53
@Author : hll
@Descripition :
*/
package main

import (
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"strings"
)

func main() {
	privateStringToAddress()
}

func privateStringToAddress() {
	//由私钥字符串转换私钥和地址
	//由私钥字符串转换私钥
	acc1Key, _ := crypto.HexToECDSA("d579f697582e36b25d05cbe25a46d343e8afb3d8f3d37feb6610d53122fcc41a")
	address1 := crypto.PubkeyToAddress(acc1Key.PublicKey)
	fmt.Println("address ", address1.String())

	dummyAddr := common.HexToAddress("9b2055d370f73ec7d8a03e965129118dc8f5bf83")
	fmt.Println("dummyAddr", dummyAddr.String())

	input := "8fa172c6f9208e4e0a05616a588288b6dd524c2633573b7c7ae059177d303c3e"
	privateKey, _ := crypto.HexToECDSA(strings.TrimPrefix(input, "0x"))
	fmt.Println("Private Key= " + hex.EncodeToString(crypto.FromECDSA(privateKey)))
	fmt.Println("public key", hex.EncodeToString(crypto.FromECDSAPub(&privateKey.PublicKey)))
	fmt.Println("address ", crypto.PubkeyToAddress(privateKey.PublicKey).String())

	msg := crypto.Keccak256([]byte("需要签名的数据"))
	//sum := sha256.Sum256([]byte("需要签名的数据"))
	//
	//h := sha256.New()
	//h.Write([]byte("需要签名的数据"))
	//sum := h.Sum(nil)

	sig, err := crypto.Sign(msg, privateKey)
	if err != nil {
		fmt.Errorf("Sign error: %s", err)
	}
	fmt.Println("签名数据为：", hexutil.Encode(sig))
	//recoveredPub, err := crypto.Ecrecover(msg, sig)
	//if err != nil {
	//	fmt.Errorf("ECRecover error: %s", err)
	//}
	//pubKey, _ := crypto.UnmarshalPubkey(recoveredPub)
	//fmt.Println("public key", hex.EncodeToString(crypto.FromECDSAPub(pubKey)))

}
