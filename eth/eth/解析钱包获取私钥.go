/*
@Created by : 2021/1/11 16:06
@Author : hll
@Descripition :
*/
package main

import (
	"encoding/hex"
	"flag"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/crypto"
	"io/ioutil"
)

// 运行程序： go run 解析钱包获取私钥.go

func main() {

	var walletPath = flag.String("walletPath", "./hllwallet", "钱包文件路径")
	var pwd = flag.String("pwd", "", "钱包密码")
	flag.Parse()
	ksfilefullpath := *walletPath
	kspassword := *pwd

	//// 钱包文件路径
	//ksfilefullpath := "E:\\gopath\\src\\hll\\eth\\UTC--2021-12-07T05-37-20.686497589Z--06765711b4d24a1582bc44c88aa77cefc8ae8d45"
	////kstype := ""
	////钱包密码
	//kspassword := ""

	keyjson, err := ioutil.ReadFile(ksfilefullpath)

	if err != nil {

		fmt.Println(err)

	}

	// Decrypt with the correct password
	key, err := keystore.DecryptKey(keyjson, kspassword)

	if err != nil {

		fmt.Printf("test : json key failed to decrypt: %v", err)

	}
	fmt.Println("address==>", key.Address.String())
	fmt.Println("Private Key= " + hex.EncodeToString(crypto.FromECDSA(key.PrivateKey)))
	fmt.Println("public key", hex.EncodeToString(crypto.FromECDSAPub(&key.PrivateKey.PublicKey)))
}
