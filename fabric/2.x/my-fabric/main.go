package main

import (
	"fmt"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/ledger"
	"github.com/hyperledger/fabric-sdk-go/pkg/core/config"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
	"log"
)

/**
 * @Description $
 * @Author hll
 * @Date 2024/5/7 18:36
 **/

func main() {
	configpath := "E:\\go-new-work\\newstudy\\fabric\\2.x\\my-fabric\\config\\org1sdk-config.yaml"
	sdk, err := fabsdk.New(config.FromFile(configpath))
	if err != nil {
		log.Panicf("failed to create fabric sdk: %s", err)
	}

	client := NewLedgerClient(sdk, "netchannel", "Node1", "Admin", "")
	block, err := client.QueryBlock(10)
	if err != nil {
		log.Panicf("failed to QueryBlock: %s", err)
	}
	fmt.Println(block.Header)
}
func NewLedgerClient(sdk *fabsdk.FabricSDK, channelID, orgName, orgAdmin, OrgUser string) (ld *ledger.Client) {
	var err error
	// create cc
	ccp := sdk.ChannelContext(channelID, fabsdk.WithUser(orgAdmin), fabsdk.WithOrg(orgName))
	ld, err = ledger.New(ccp)
	if err != nil {
		log.Panicf("failed to create channel client: %s", err)
	}
	log.Println("Initialized channel client")

	return ld
}
