package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"log"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

// BEP-20 合约 ABI
const bep20ABI = `[{"constant":true,"inputs":[{"internalType":"address","name":"account","type":"address"}],"name":"balanceOf","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"}]`

func main() {
	// Binance Smart Chain 节点 URL
	bscURL := "https://bsc.blockrazor.xyz" // 可替换为其他节点服务
	client, err := ethclient.Dial(bscURL)
	if err != nil {
		panic(fmt.Sprintf("连接到 BSC 节点失败: %v", err))
	}

	// BEP-20 合约地址
	tokenAddress := common.HexToAddress("0x55d398326f99059fF775485246999027B3197955") // 替换为目标代币的合约地址

	// 目标账户地址
	walletAddress := common.HexToAddress("0xD183F2BBF8b28d9fec8367cb06FE72B88778C86B") // 替换为目标查询的钱包地址

	parsedABI, err := abi.JSON(strings.NewReader(bep20ABI))
	if err != nil {
		log.Fatalf("Failed to parse ABI: %v\n", err)
	}

	// 方法参数：查询某地址的余额，假设查询的地址是 publicKey
	methodName := "balanceOf"
	params := []interface{}{walletAddress}

	// 编码调用数据
	data, err := parsedABI.Pack(methodName, params...)
	if err != nil {
		log.Fatalf("Failed to pack data: %v\n", err)
	}
	// 调用合约的查询方法
	msg := ethereum.CallMsg{
		To:   &tokenAddress,
		Data: data,
	}

	// 执行查询
	result, err := client.CallContract(context.Background(), msg, nil)
	if err != nil {
		log.Fatalf("Failed to call contract: %v", err)
	}

	// 解析返回值（返回的是 uint256 类型的余额）
	var balance *big.Int
	err = parsedABI.UnpackIntoInterface(&balance, methodName, result)
	if err != nil {
		log.Fatalf("Failed to unpack result: %v", err)
	}

	// 输出查询结果
	fmt.Printf("账户余额: %s\n", balance.String())

}
