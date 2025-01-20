package chaincode

import (
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// SmartContract provides functions for managing an Asset
type SmartContract struct {
	contractapi.Contract
}

// CreateAsset issues a new asset to the world state with given details.
func (s *SmartContract) Set(ctx contractapi.TransactionContextInterface, key string, vue string) error {
	fmt.Println("creating set")
	err := ctx.GetStub().PutState(key, []byte(vue))
	if err != nil {
		return fmt.Errorf("failed to set asset: %s", key)
	}
	return nil
}

// ReadAsset returns the asset stored in the world state with given id.
func (s *SmartContract) Get(ctx contractapi.TransactionContextInterface, key string) (string, error) {
	fmt.Println("creating get")
	value, err := ctx.GetStub().GetState(key)
	if err != nil {
		return "", fmt.Errorf("failed to read from world state: %v", err)
	}
	if value == nil {
		return "", fmt.Errorf("the asset %s does not exist", key)
	}
	return string(value), nil
}
