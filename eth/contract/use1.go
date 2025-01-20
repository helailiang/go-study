package main

import (
	"context"
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
	"strings"
)

func main() {
	unsignedTxDataFunc()
	//unsignedTxDataCallContract()
	//unsignedTxDataCallContractQuery()
	//send()
}

func send() {
	// 连接到以太坊客户端
	client, err := ethclient.Dial("https://holesky.infura.io/v3/364d773fa11d4081be873f0673002443")
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	signTx := "f88a578403d039cb830493e0946834465217ef45237d81e891930589beb854e43e80a460fe47b100000000000000000000000000000000000000000000000000000000000003e88284f4a0322ca6be1a366caef5d6a75af949bcb2e451228fc5d96e5a7f60ff34b0ece3cca04d38820ab45b06fb590b6ecab722ff628d7b1b767ee7b55408431a1028361e07"
	s, err := hex.DecodeString(signTx)
	if err != nil {
		panic(err)
	}
	var signedTx types.Transaction
	//types.LegacyTx
	err = signedTx.UnmarshalBinary([]byte(s))
	if err != nil {
		panic(err)
	}
	err = client.SendTransaction(context.Background(), &signedTx)
	if err != nil {
		panic(err)
	}
	// 输出交易哈希
	fmt.Printf("Transaction sent! Hash: %s\n", signedTx.Hash().Hex())
}

func unsignedTxDataFunc() {
	// 连接到以太坊客户端
	client, err := ethclient.Dial("https://holesky.infura.io/v3/364d773fa11d4081be873f0673002443")
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	// 定义必要的参数
	privateKey, err := crypto.HexToECDSA("274cfd629a3e82a786df5df3266c615b7bdbecd6fb8bdff2d9290d8e6cdf0635")
	if err != nil {
		log.Fatalf("Failed to load private key: %v\n", err)
	}
	publicKey := crypto.PubkeyToAddress(privateKey.PublicKey)

	nonce, err := client.PendingNonceAt(context.Background(), publicKey)
	if err != nil {
		log.Fatalf("Failed to get nonce: %v\n", err)
	}

	gasLimit := uint64(3000000) // 设置 Gas 上限
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatalf("Failed to get gas price: %v\n", err)
	}

	// 合约二进制代码 (bin)
	contractBin := "608060405234801561001057600080fd5b506040516101203803806101208339818101604052602081101561003357600080fd5b8101908080519060200190929190505050806000819055505060c68061005a6000396000f3fe6080604052348015600f57600080fd5b506004361060325760003560e01c806360fe47b11460375780636d4ce63c146062575b600080fd5b606060048036036020811015604b57600080fd5b8101908080359060200190929190505050607e565b005b60686088565b6040518082815260200191505060405180910390f35b8060008190555050565b6000805490509056fea265627a7a723158206898197779f6c3d78cdbeea055685281cb642e99b2cc3c6d8802292300f1314064736f6c63430005100032"
	//0x0dbe671f + address + uint256

	// 构造函数参数 ABI 编码
	// 示例：构造函数接受一个 uint256 参数和一个地址参数
	abiDefinition := `[
	{
		"inputs": [
			{
				"internalType": "uint256",
				"name": "x",
				"type": "uint256",
				"description": "x"
			}
		],
		"payable": false,
		"stateMutability": "nonpayable",
		"type": "constructor"
	},
	{
		"constant": true,
		"inputs": [],
		"name": "get",
		"outputs": [
			{
				"internalType": "uint256",
				"name": "",
				"type": "uint256",
				"description": "shuchu"
			}
		],
		"payable": false,
		"stateMutability": "view",
		"type": "function"
	},
	{
		"constant": false,
		"inputs": [
			{
				"internalType": "uint256",
				"name": "x",
				"type": "uint256",
                "description": "shuru"
			}
		],
		"name": "set",
		"outputs": [],
		"payable": false,
		"stateMutability": "nonpayable",
		"type": "function"
	}
]`
	parsedABI, err := abi.JSON(strings.NewReader(abiDefinition))
	if err != nil {
		log.Fatalf("Failed to parse ABI: %v\n", err)
	}

	constructorArgs, err := parsedABI.Pack("", big.NewInt(12345)) //big.NewInt(12345)
	if err != nil {
		log.Fatalf("Failed to pack constructor arguments: %v\n", err)
	}

	// 合并 bin 和构造函数参数
	contractData := append(common.FromHex(contractBin), constructorArgs...)
	//contractData := common.FromHex(contractBin)

	// 构造交易对象
	tx := types.NewContractCreation(nonce, big.NewInt(0), gasLimit, gasPrice, contractData)

	// 将交易序列化为 RLP 格式（未签名的原始数据）
	unsignedTxData, err := tx.MarshalBinary()
	if err != nil {
		log.Fatalf("Failed to marshal transaction: %v\n", err)
	}

	// 输出未签名的交易数据
	fmt.Printf("Unsigned Transaction Data: %s\n", hex.EncodeToString(unsignedTxData))
	// 对交易进行签名
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(big.NewInt(17000)), privateKey) // 这里 1 是主网的 Chain ID
	if err != nil {
		log.Fatalf("Failed to sign transaction: %v\n", err)
	}
	// 输出签名后的交易数据
	fmt.Printf("Signed Transaction: %v\n", signedTx)

	// 发送交易
	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatalf("Failed to send transaction: %v\n", err)
	}

	// 输出交易哈希
	fmt.Printf("Transaction sent! Hash: %s\n", signedTx.Hash().Hex())

}

func unsignedTxDataCallContract() {
	// 连接到以太坊客户端
	client, err := ethclient.Dial("https://holesky.infura.io/v3/364d773fa11d4081be873f0673002443")
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v\n", err)
	}

	// 定义必要的参数
	privateKey, err := crypto.HexToECDSA("274cfd629a3e82a786df5df3266c615b7bdbecd6fb8bdff2d9290d8e6cdf0635")
	if err != nil {
		log.Fatalf("Failed to load private key: %v\n", err)
	}
	publicKey := crypto.PubkeyToAddress(privateKey.PublicKey)

	nonce, err := client.PendingNonceAt(context.Background(), publicKey)
	if err != nil {
		log.Fatalf("Failed to get nonce: %v\n", err)
	}

	gasLimit := uint64(300000) // 设置 Gas 上限
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatalf("Failed to get gas price: %v\n", err)
	}

	// 合约地址 (已部署合约的地址)
	contractAddress := common.HexToAddress("0x6834465217ef45237d81e891930589beb854e43e")

	// ABI 定义（假设你有合约 ABI）
	abiDefinition := `[
	{
		"inputs": [
			{
				"internalType": "uint256",
				"name": "x",
				"type": "uint256",
				"description": "x"
			}
		],
		"payable": false,
		"stateMutability": "nonpayable",
		"type": "constructor"
	},
	{
		"constant": true,
		"inputs": [],
		"name": "get",
		"outputs": [
			{
				"internalType": "uint256",
				"name": "",
				"type": "uint256",
				"description": "shuchu"
			}
		],
		"payable": false,
		"stateMutability": "view",
		"type": "function"
	},
	{
		"constant": false,
		"inputs": [
			{
				"internalType": "uint256",
				"name": "x",
				"type": "uint256",
                "description": "shuru"
			}
		],
		"name": "set",
		"outputs": [],
		"payable": false,
		"stateMutability": "nonpayable",
		"type": "function"
	}
]`
	parsedABI, err := abi.JSON(strings.NewReader(abiDefinition))
	if err != nil {
		log.Fatalf("Failed to parse ABI: %v\n", err)
	}

	// 方法参数：查询某地址的余额，假设查询的地址是 publicKey
	methodName := "set"
	params := []interface{}{big.NewInt(1000)}

	// 编码调用数据
	data, err := parsedABI.Pack(methodName, params...)
	if err != nil {
		log.Fatalf("Failed to pack data: %v\n", err)
	}

	// 构造交易对象
	tx := types.NewTransaction(nonce, contractAddress, big.NewInt(0), gasLimit, gasPrice, data)

	// 输出未签名的交易数据
	unsignedTxData, err := tx.MarshalBinary()
	if err != nil {
		log.Fatalf("Failed to marshal transaction: %v\n", err)
	}

	fmt.Printf("Unsigned Transaction Data: %s\n", hex.EncodeToString(unsignedTxData))

	// 对交易进行签名
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(big.NewInt(17000)), privateKey) // 1 为主网 Chain ID
	if err != nil {
		log.Fatalf("Failed to sign transaction: %v\n", err)
	}
	signedTxAfer, err := signedTx.MarshalBinary()
	if err != nil {
		log.Fatalf("Failed to sign after transaction: %v\n", err)
	}
	fmt.Printf("signed Transaction Data: %s\n", hex.EncodeToString(signedTxAfer))

	//// 发送签名后的交易
	//err = client.SendTransaction(context.Background(), signedTx)
	//if err != nil {
	//	log.Fatalf("Failed to send transaction: %v\n", err)
	//}

	// 输出交易哈希
	fmt.Printf("Transaction sent! Hash: %s\n", signedTx.Hash().Hex())
}

func unsignedTxDataCallContractQuery() {
	// 连接到以太坊客户端
	client, err := ethclient.Dial("https://holesky.infura.io/v3/364d773fa11d4081be873f0673002443")
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v\n", err)
	}

	// 定义必要的参数
	//privateKey, err := crypto.HexToECDSA("274cfd629a3e82a786df5df3266c615b7bdbecd6fb8bdff2d9290d8e6cdf0635")
	//if err != nil {
	//	log.Fatalf("Failed to load private key: %v\n", err)
	//}
	//publicKey := crypto.PubkeyToAddress(privateKey.PublicKey)

	//nonce, err := client.PendingNonceAt(context.Background(), publicKey)
	//if err != nil {
	//	log.Fatalf("Failed to get nonce: %v\n", err)
	//}
	//
	//gasLimit := uint64(300000) // 设置 Gas 上限
	//gasPrice, err := client.SuggestGasPrice(context.Background())
	//if err != nil {
	//	log.Fatalf("Failed to get gas price: %v\n", err)
	//}

	// 合约地址 (已部署合约的地址)
	contractAddress := common.HexToAddress("0x6834465217ef45237d81e891930589beb854e43e")

	// ABI 定义（假设你有合约 ABI）
	abiDefinition := `[
	{
		"inputs": [
			{
				"internalType": "uint256",
				"name": "x",
				"type": "uint256",
				"description": "x"
			}
		],
		"payable": false,
		"stateMutability": "nonpayable",
		"type": "constructor"
	},
	{
		"constant": true,
		"inputs": [],
		"name": "get",
		"outputs": [
			{
				"internalType": "uint256",
				"name": "",
				"type": "uint256",
				"description": "shuchu"
			}
		],
		"payable": false,
		"stateMutability": "view",
		"type": "function"
	},
	{
		"constant": false,
		"inputs": [
			{
				"internalType": "uint256",
				"name": "x",
				"type": "uint256",
                "description": "shuru"
			}
		],
		"name": "set",
		"outputs": [],
		"payable": false,
		"stateMutability": "nonpayable",
		"type": "function"
	}
]`
	parsedABI, err := abi.JSON(strings.NewReader(abiDefinition))
	if err != nil {
		log.Fatalf("Failed to parse ABI: %v\n", err)
	}

	// 方法参数：查询某地址的余额，假设查询的地址是 publicKey
	methodName := "get"
	//params := []interface{}{big.NewInt(1000)}

	// 编码调用数据
	data, err := parsedABI.Pack(methodName)
	if err != nil {
		log.Fatalf("Failed to pack data: %v\n", err)
	}
	// 调用合约的查询方法
	msg := ethereum.CallMsg{
		To:   &contractAddress,
		Data: data,
	}

	// 执行查询
	result, err := client.CallContract(context.Background(), msg, nil)
	if err != nil {
		log.Fatalf("Failed to call contract: %v", err)
	}

	// 解析返回值（返回的是 uint256 类型的余额）
	var balance *big.Int
	err = parsedABI.UnpackIntoInterface(&balance, "get", result)
	if err != nil {
		log.Fatalf("Failed to unpack result: %v", err)
	}

	// 输出查询结果
	fmt.Printf("Balance: %s\n", balance.String())
}
