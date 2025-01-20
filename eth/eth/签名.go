/*
@Created by : 2021/8/6 15:43 
@Author : hll
@Descripition :
*/
package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"log"
)

func main(){
	privateKey, err := crypto.HexToECDSA("d579f697582e36b25d05cbe25a46d343e8afb3d8f3d37feb6610d53122fcc41a")
	if err != nil {
		log.Fatal(err)
	}
	data := []byte("需要签名的数据")
	hash := crypto.Keccak256(data)
	//fmt.Println(hash.Hex()) // 0x1c8aff950685c2ed4bc3174f3472287b56d9517b9c948127319a09a7a36deac8
	signature, err := crypto.Sign(hash, privateKey)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(hexutil.Encode(signature))
}

