/*
@Created by : 2021/1/11 17:22
@Author : hll
@Descripition :
*/
package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/keystore"
)

func main() {
	//d, err := ioutil.TempDir("E:\\go_workspace\\src\\hll\\AWS\\hll", "")
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	ks := keystore.NewPlaintextKeyStore("E:\\go_workspace\\src\\hll\\AWS")
	ks.NewAccount("123456")
	fmt.Printf("===> %+v", ks)

}
