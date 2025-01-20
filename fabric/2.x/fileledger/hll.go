package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	cb "github.com/hyperledger/fabric-protos-go/common"
	"io/ioutil"
	"os"
)

/**
 * @Description $
 * @Author hll
 * @Date 2024/9/30 下午3:23
 **/

func main() {
	// 定义要写入的文件名
	fileName := "payload_data.txt"
	oldb, err := ioutil.ReadFile("a.txt")
	if err != nil {
		panic(err)
	}

	var committedBlock = &cb.Block{}
	err = proto.Unmarshal(oldb, committedBlock)
	if err != nil {
		panic(fmt.Errorf("Unmarshal序列化区块失败：%s", err))
	}
	//b, err := proto.Marshal(committedBlock)
	//if err != nil {
	//	panic(fmt.Errorf("序列化区块失败：%s", err))
	//}
	fmt.Printf("PreviousHash=>%s\n", committedBlock.Header.PreviousHash)
	err = os.WriteFile(fileName, oldb, 0644)
	if err != nil {
		fmt.Println("Error writing file:", err)
		return
	}
	fmt.Println("Payload data written to file successfully.")
}
