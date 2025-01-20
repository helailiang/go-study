package parse

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/hyperledger/fabric-protos-go/common"
	"github.com/hyperledger/fabric-protos-go/orderer"
	"github.com/hyperledger/fabric-protos-go/peer"
	"github.com/hyperledger/fabric/common/configtx"
	"github.com/hyperledger/fabric/protoutil"
	"time"
)

// TxValidationFlags is array of transaction validation codes. It is used when committer validates block.
type TxValidationFlags []uint8

// NewTxValidationFlags Create new object-array of validation codes with target size.
// Default values: TxValidationCode_NOT_VALIDATED
func NewTxValidationFlags(size int) TxValidationFlags {
	return newTxValidationFlagsSetValue(size, peer.TxValidationCode_NOT_VALIDATED)
}

func newTxValidationFlagsSetValue(size int, value peer.TxValidationCode) TxValidationFlags {
	inst := make(TxValidationFlags, size)
	for i := range inst {
		inst[i] = uint8(value)
	}

	return inst
}

// SetFlag assigns validation code to specified transaction
func (obj TxValidationFlags) SetFlag(txIndex int, flag peer.TxValidationCode) {
	obj[txIndex] = uint8(flag)
}

// Flag returns validation code at specified transaction
func (obj TxValidationFlags) Flag(txIndex int) peer.TxValidationCode {
	return peer.TxValidationCode(obj[txIndex])
}

// IsValid checks if specified transaction is valid
func (obj TxValidationFlags) IsValid(txIndex int) bool {
	return obj.IsSetTo(txIndex, peer.TxValidationCode_VALID)
}

// IsInvalid checks if specified transaction is invalid
func (obj TxValidationFlags) IsInvalid(txIndex int) bool {
	return !obj.IsValid(txIndex)
}

// IsSetTo returns true if the specified transaction equals flag; false otherwise.
func (obj TxValidationFlags) IsSetTo(txIndex int, flag peer.TxValidationCode) bool {
	return obj.Flag(txIndex) == flag
}

func PreprocessProtoBlock(blk *common.Block, channelId string, needParse bool) {

	//b := &block{num: blk.Header.Number}
	fmt.Println("blk.Header.Number", blk.Header.Number)
	fmt.Printf("blk.Header.PreviousHash:[%x]\n", blk.Header.PreviousHash)
	if !needParse {
		return
	}
	meta, err := protoutil.GetMetadataFromBlock(blk, common.BlockMetadataIndex_ORDERER)
	if err != nil {
		panic(err)
	}
	kfkMeata := &orderer.KafkaMetadata{}
	err = proto.Unmarshal(meta.Value, kfkMeata)
	if err != nil {
		panic(err)
	}
	fmt.Printf("区块号【%d】偏移量：LastOffsetPersisted:%d, LastOriginalOffsetProcessed:%d, LastResubmittedConfigOffset:%d \n",
		blk.Header.Number, kfkMeata.LastOffsetPersisted, kfkMeata.LastOriginalOffsetProcessed, kfkMeata.LastResubmittedConfigOffset)
	// Committer validator has already set validation flags based on well formed tran checks
	txsFilter := TxValidationFlags(blk.Metadata.Metadata[common.BlockMetadataIndex_TRANSACTIONS_FILTER])
	if len(txsFilter) == 0 {
		txsFilter = NewTxValidationFlags(len(blk.Data.Data))
		blk.Metadata.Metadata[common.BlockMetadataIndex_TRANSACTIONS_FILTER] = txsFilter
	}
	for txIndex, envBytes := range blk.Data.Data {
		var env *common.Envelope
		var chdr *common.ChannelHeader
		var payload *common.Payload
		var err error

		if env, err = protoutil.GetEnvelopeFromBlock(envBytes); err == nil {
			if payload, err = protoutil.UnmarshalPayload(env.Payload); err == nil {
				chdr, err = protoutil.UnmarshalChannelHeader(payload.Header.ChannelHeader)
			}
		}
		fmt.Println("txid==>", chdr.GetTxId())
		fmt.Printf("Channel [%s]: Block [%d] Transaction index [%d] TxId [%s]"+
			" time:%s. Reason code [%s]\n",
			chdr.GetChannelId(), blk.Header.Number, txIndex,
			chdr.GetTxId(), time.Unix(chdr.Timestamp.Seconds, int64(chdr.Timestamp.Nanos)).Format(time.DateTime),
			txsFilter.Flag(txIndex).String())
		//if txsFilter.IsInvalid(txIndex) {
		//	// Skipping invalid transaction
		//
		//	continue
		//}
		if err != nil {
			return
		}

		//var txRWSet *rwsetutil.TxRwSet
		txType := common.HeaderType(chdr.Type)
		fmt.Printf("txType=%s\n", txType)
		if txType == common.HeaderType_ENDORSER_TRANSACTION {
			// extract actions from the envelope message
			respPayload, err := protoutil.GetActionFromEnvelope(envBytes)
			if err != nil {
				txsFilter.SetFlag(txIndex, peer.TxValidationCode_NIL_TXACTION)
				continue
			}
			//txStatInfo.ChaincodeID = respPayload.ChaincodeId
			//txStatInfo.ChaincodeEventData = respPayload.Events
			fmt.Printf("ChaincodeID:%s ChaincodeEventData:%s\n", respPayload.ChaincodeId, respPayload.Events)
			/*txRWSet = &rwsetutil.TxRwSet{}
			if err = txRWSet.FromProtoBytes(respPayload.Results); err != nil {
				txsFilter.SetFlag(txIndex, peer.TxValidationCode_INVALID_OTHER_REASON)
				continue
			}*/
		} else if common.HeaderType(chdr.Type) == common.HeaderType_CONFIG {
			configEnvelope, err := configtx.UnmarshalConfigEnvelope(payload.Data)
			if err != nil {
				fmt.Println(err, "error unmarshalling config which passed initial validity checks")
				return
			}
			fmt.Printf("config currently at sequence %d\n", configEnvelope.Config.Sequence)
		}

	}
	return
}
