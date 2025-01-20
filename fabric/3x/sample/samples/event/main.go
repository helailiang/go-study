package main

import (
	"crypto/sha256"
	"encoding/asn1"
	"encoding/base64"
	"github.com/hyperledger/fabric-protos-go/common"
	"log"
	"math/big"
	"time"

	"github.com/hyperledger/fabric-sdk-go/pkg/client/event"
	"github.com/hyperledger/fabric-sdk-go/pkg/common/providers/fab"
	"github.com/shitaibin/fabric-sdk-go-sample/cli"
)

const (
	//org1CfgPath = "config/node1sdk-config.yaml"
	org1CfgPath = "config/beijing1sdk-config.yaml"

	//org2CfgPath = "../../config/org2sdk-config.yaml"
)

var (
	//peer0Org1 = "peer1.node1.private.bsnbase.com"
	peer0Org1 = "peer1.gg.ggconnect.bsnbase.com"

	//peer0Org2 = "peer0.org2.example.com"
)

func main() {
	//log.Println(os.Getwd())
	org1Client := cli.New(org1CfgPath, "gg.ggconnect.bsnbase.com", "Admin", "Admin")
	//org2Client := cli.New(org2CfgPath, "Org2", "Admin", "User1")
	defer org1Client.Close()
	//defer org2Client.Close()

	// New event client
	//cp := org1Client.SDK.ChannelContext(org1Client.ChannelID, fabsdk.WithUser(org1Client.OrgUser))

	//ec, err := event.New(
	//	cp,
	//	event.WithBlockEvents(), // 如果没有，会是filtered
	//	// event.WithBlockNum(1), // 从指定区块获取，需要此参数
	//	event.WithSeekType(seek.Newest))
	//if err != nil {
	//	log.Printf("Create event client error: %v", err)
	//}
	//
	//// block event listen
	////监听区块, 结束取消事件监听
	//blockListener(ec)
	//defer ec.Unregister(blockListener(ec))
	//监听链码事件
	//defer ec.Unregister(filteredBlockListener(ec))

	// tx listen
	txIDCh := make(chan string, 100)

	//监听交易事件
	//go txListener(ec, txIDCh)

	// chaincode event listen
	//defer ec.Unregister(chainCodeEventListener(nil, ec))

	DoChainCode(org1Client, txIDCh)
	close(txIDCh)

	time.Sleep(time.Second * 300)
}

// 监听区块
func blockListener(ec *event.Client) fab.Registration {
	// Register monitor block event
	beReg, beCh, err := ec.RegisterBlockEvent()
	if err != nil {
		log.Printf("Register block event error: %v\n", err)
	}
	log.Println(">>>>>>>>>> Registered block event")

	// Receive block event
	go func() {
		for e := range beCh {
			log.Printf("Receive block event:\nSourceURL: %v\nNumber: %v\nHash"+
				": %v\nPreviousHash: %v\n\n",
				e.SourceURL,
				e.Block.Header.Number,
				base64.StdEncoding.EncodeToString(BlockHeaderHashBytes(e.Block.Header)),
				base64.StdEncoding.EncodeToString(e.Block.Header.PreviousHash))

			//expectedEnvelope := &common.Envelope{}
			//err = proto.Unmarshal(e.Block.Data.Data[0], expectedEnvelope)
			//expectedPayload := &common.Payload{}
			//err = proto.Unmarshal(expectedEnvelope.Payload, expectedPayload)
			//log.Printf("====>%s\n", expectedPayload)
		}
	}()

	return beReg
}

// 监听链码事件
func filteredBlockListener(ec *event.Client) fab.Registration {
	// Register monitor filtered block event
	fbeReg, fbeCh, err := ec.RegisterFilteredBlockEvent()
	if err != nil {
		log.Printf("Register filtered block event error: %v", err)
	}
	log.Println("Registered filtered block event")

	// Receive filtered block event
	go func() {
		for e := range fbeCh {
			log.Printf("Receive filterd block event:\nNumber: %v\nlen("+
				"transactions): %v\nSourceURL: %v",
				e.FilteredBlock.Number, len(e.FilteredBlock.
					FilteredTransactions), e.SourceURL)

			for i, tx := range e.FilteredBlock.FilteredTransactions {
				log.Printf("tx index %d: type: %v, txid: %v, "+
					"validation code: %v", i,
					tx.Type, tx.Txid,
					tx.TxValidationCode)
			}
			log.Println() // Just go print empty log for easy to read
		}
	}()

	return fbeReg
}

// 监听交易事件
func txListener(ec *event.Client, txIDCh chan string) {
	log.Println("Transaction listener start")
	defer log.Println("Transaction listener exit")

	for id := range txIDCh {
		// Register monitor transaction event
		log.Printf("Register transaction event for: %v\n", id)
		txReg, txCh, err := ec.RegisterTxStatusEvent(id)
		if err != nil {
			log.Printf("Register transaction event error: %v\n", err)
			continue
		}
		defer ec.Unregister(txReg)

		// Receive transaction event
		go func() {
			for e := range txCh {
				log.Printf("Receive transaction event: txid: %v, "+
					"validation code: %v, block number: %v\n",
					e.TxID,
					e.TxValidationCode,
					e.BlockNumber)
			}
		}()
	}
}

func chainCodeEventListener(c *cli.Client, ec *event.Client) fab.Registration {
	eventName := ".*"
	log.Printf("Listen chaincode event: %v\n", eventName)

	var (
		ccReg   fab.Registration
		eventCh <-chan *fab.CCEvent
		err     error
	)
	if c != nil {
		log.Println("Using client to register chaincode event")
		ccReg, eventCh, err = c.RegisterChaincodeEvent("mycc", eventName)
	} else {
		log.Println("Using event client to register chaincode event")
		ccReg, eventCh, err = ec.RegisterChaincodeEvent("mycc", eventName)
	}
	if err != nil {
		log.Printf("Register chaincode event error: %v\n", err.Error())
		return nil
	}

	// consume event
	go func() {
		for e := range eventCh {
			log.Printf("Receive cc event, ccid: %v \neventName: %v\n"+
				"payload: %v \ntxid: %v \nblock: %v \nsourceURL: %v\n",
				e.ChaincodeID, e.EventName, string(e.Payload), e.TxID, e.BlockNumber, e.SourceURL)
		}
	}()

	return ccReg
}

// Install、Deploy、Invoke、Query、Upgrade
func DoChainCode(cli1 *cli.Client, txCh chan<- string) {
	var (
		txid fab.TransactionID
		err  error
	)

	// ccVersion := "v1"
	// if err := cli1.InstallCC(ccVersion, peer0Org1); err != nil {
	// 	log.Panicf("Intall chaincode error: %v", err)
	// }
	// log.Println("Chaincode has been installed on org1's peers")
	//
	// // InstantiateCC chaincode only need once for each channel
	// if txid, err = cli1.InstantiateCC(ccVersion, peer0Org1); err != nil {
	// 	log.Panicf("Instantiated chaincode error: %v", err)
	// }
	// if txid != "" {
	// 	txCh <- string(txid)
	// }
	// log.Println("Chaincode has been instantiated")
	//
	if txid, err = cli1.InvokeCC([]string{peer0Org1}); err != nil {
		log.Panicf("Invoke chaincode error: %v", err)
	}
	if txid != "" {
		txCh <- string(txid)
	}
	log.Println(">>>>>>>> Invoke chaincode success")

	//if txid, err = cli1.InvokeCCDelete([]string{peer0Org1}); err != nil {
	//	log.Printf("Invoke chaincode delete error: %v", err)
	//}
	//
	if err := cli1.QueryCC(peer0Org1, "a"); err != nil {
		log.Panicf("Query chaincode error: %v", err)
	}
	//log.Println("Query chaincode success on peer0.org1")
}

type asn1Header struct {
	Number       *big.Int
	PreviousHash []byte
	DataHash     []byte
}

func BlockHeaderHashBytes(b *common.BlockHeader) []byte {
	asn1Header := asn1Header{
		PreviousHash: b.PreviousHash,
		DataHash:     b.DataHash,
		Number:       new(big.Int).SetUint64(b.Number),
	}
	result, err := asn1.Marshal(asn1Header)
	if err != nil {
		// Errors should only arise for types which cannot be encoded, since the
		// BlockHeader type is known a-priori to contain only encodable types, an
		// error here is fatal and should not be propagated
		panic(err)
	}
	hasher := sha256.New()
	hasher.Write(result)
	return hasher.Sum(nil)
}
