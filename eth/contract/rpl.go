package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/rlp"
	"log"
)

func main() {
	// 参数定义
	_endpoint := common.HexToAddress("0xf37817f975cf3865abd5f7a860ff4968e3ad23c9")
	_receiveUln301 := common.HexToAddress("0x0000000000000000000000000000000000000000")
	_messageLibs := []common.Address{common.HexToAddress("0xe83d86ed3352a4eb4a8098e3f189ebecfb5e0a22")}
	_priceFeed := common.HexToAddress("0x0000000000000000000000000000000000000000")
	_roleAdmin := common.HexToAddress("0x4e0539F3e0A941dCD0617F636381b4c3BC25F048")
	_admins := []common.Address{common.HexToAddress("0x4e0539F3e0A941dCD0617F636381b4c3BC25F048")}

	// 对每个地址进行 RLP 编码
	rlpEncodedEndpoint, err := rlp.EncodeToBytes(_endpoint)
	if err != nil {
		log.Fatalf("Failed to RLP encode _endpoint: %v", err)
	}

	rlpEncodedReceiveUln301, err := rlp.EncodeToBytes(_receiveUln301)
	if err != nil {
		log.Fatalf("Failed to RLP encode _receiveUln301: %v", err)
	}

	rlpEncodedMessageLibs, err := rlp.EncodeToBytes(_messageLibs)
	if err != nil {
		log.Fatalf("Failed to RLP encode _messageLibs: %v", err)
	}

	rlpEncodedPriceFeed, err := rlp.EncodeToBytes(_priceFeed)
	if err != nil {
		log.Fatalf("Failed to RLP encode _priceFeed: %v", err)
	}

	rlpEncodedRoleAdmin, err := rlp.EncodeToBytes(_roleAdmin)
	if err != nil {
		log.Fatalf("Failed to RLP encode _roleAdmin: %v", err)
	}

	rlpEncodedAdmins, err := rlp.EncodeToBytes(_admins)
	if err != nil {
		log.Fatalf("Failed to RLP encode _admins: %v", err)
	}

	// 将所有编码后的参数拼接在一起
	allEncodedParams := append(rlpEncodedEndpoint, rlpEncodedReceiveUln301...)
	allEncodedParams = append(allEncodedParams, rlpEncodedMessageLibs...)
	allEncodedParams = append(allEncodedParams, rlpEncodedPriceFeed...)
	allEncodedParams = append(allEncodedParams, rlpEncodedRoleAdmin...)
	allEncodedParams = append(allEncodedParams, rlpEncodedAdmins...)

	// 输出最终的 RLP 编码结果
	fmt.Printf("RLP Encoded Parameters: %x\n", allEncodedParams)
}
