// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package endpointv2

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// MessagingFee is an auto generated low-level Go binding around an user-defined struct.
type MessagingFee struct {
	NativeFee  *big.Int
	LzTokenFee *big.Int
}

// MessagingParams is an auto generated low-level Go binding around an user-defined struct.
type MessagingParams struct {
	DstEid       uint32
	Receiver     [32]byte
	Message      []byte
	Options      []byte
	PayInLzToken bool
}

// MessagingReceipt is an auto generated low-level Go binding around an user-defined struct.
type MessagingReceipt struct {
	Guid  [32]byte
	Nonce uint64
	Fee   MessagingFee
}

// Origin is an auto generated low-level Go binding around an user-defined struct.
type Origin struct {
	SrcEid uint32
	Sender [32]byte
	Nonce  uint64
}

// SetConfigParam is an auto generated low-level Go binding around an user-defined struct.
type SetConfigParam struct {
	Eid        uint32
	ConfigType uint32
	Config     []byte
}

// Endpointv2MetaData contains all meta data concerning the Endpointv2 contract.
var Endpointv2MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_eid\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"LZ_AlreadyRegistered\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LZ_ComposeExists\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"expected\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"actual\",\"type\":\"bytes32\"}],\"name\":\"LZ_ComposeNotFound\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LZ_DefaultReceiveLibUnavailable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LZ_DefaultSendLibUnavailable\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"requiredNative\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"suppliedNative\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requiredLzToken\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"suppliedLzToken\",\"type\":\"uint256\"}],\"name\":\"LZ_InsufficientFee\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LZ_InvalidExpiry\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"}],\"name\":\"LZ_InvalidNonce\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LZ_InvalidPayloadHash\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LZ_InvalidReceiveLibrary\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LZ_LzTokenUnavailable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LZ_OnlyNonDefaultLib\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LZ_OnlyReceiveLib\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LZ_OnlyRegisteredLib\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LZ_OnlyRegisteredOrDefaultLib\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LZ_OnlySendLib\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LZ_PathNotInitializable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LZ_PathNotVerifiable\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"expected\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"actual\",\"type\":\"bytes32\"}],\"name\":\"LZ_PayloadHashNotFound\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LZ_SameValue\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LZ_SendReentrancy\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LZ_Unauthorized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LZ_UnsupportedEid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LZ_UnsupportedInterface\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LZ_ZeroLzTokenFee\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"Transfer_NativeFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Transfer_ToAddressIsZero\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"guid\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"index\",\"type\":\"uint16\"}],\"name\":\"ComposeDelivered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"guid\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"index\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"ComposeSent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"eid\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newLib\",\"type\":\"address\"}],\"name\":\"DefaultReceiveLibrarySet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"eid\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldLib\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"}],\"name\":\"DefaultReceiveLibraryTimeoutSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"eid\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newLib\",\"type\":\"address\"}],\"name\":\"DefaultSendLibrarySet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"delegate\",\"type\":\"address\"}],\"name\":\"DelegateSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"srcEid\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"sender\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"}],\"name\":\"InboundNonceSkipped\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newLib\",\"type\":\"address\"}],\"name\":\"LibraryRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"executor\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"guid\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"index\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"gas\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"reason\",\"type\":\"bytes\"}],\"name\":\"LzComposeAlert\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"executor\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"srcEid\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"sender\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"}],\"indexed\":false,\"internalType\":\"structOrigin\",\"name\":\"origin\",\"type\":\"tuple\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"guid\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"gas\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"reason\",\"type\":\"bytes\"}],\"name\":\"LzReceiveAlert\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"LzTokenSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"srcEid\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"sender\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"payloadHash\",\"type\":\"bytes32\"}],\"name\":\"PacketBurnt\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"srcEid\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"sender\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"}],\"indexed\":false,\"internalType\":\"structOrigin\",\"name\":\"origin\",\"type\":\"tuple\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"PacketDelivered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"srcEid\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"sender\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"payloadHash\",\"type\":\"bytes32\"}],\"name\":\"PacketNilified\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"encodedPayload\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"options\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sendLibrary\",\"type\":\"address\"}],\"name\":\"PacketSent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"srcEid\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"sender\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"}],\"indexed\":false,\"internalType\":\"structOrigin\",\"name\":\"origin\",\"type\":\"tuple\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"payloadHash\",\"type\":\"bytes32\"}],\"name\":\"PacketVerified\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"eid\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newLib\",\"type\":\"address\"}],\"name\":\"ReceiveLibrarySet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"eid\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldLib\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timeout\",\"type\":\"uint256\"}],\"name\":\"ReceiveLibraryTimeoutSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"eid\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newLib\",\"type\":\"address\"}],\"name\":\"SendLibrarySet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"EMPTY_PAYLOAD_HASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"NIL_PAYLOAD_HASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"blockedLibrary\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_oapp\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"_srcEid\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"_sender\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"_nonce\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"_payloadHash\",\"type\":\"bytes32\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_oapp\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"srcEid\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"sender\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"}],\"internalType\":\"structOrigin\",\"name\":\"_origin\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"_guid\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"}],\"name\":\"clear\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"guid\",\"type\":\"bytes32\"},{\"internalType\":\"uint16\",\"name\":\"index\",\"type\":\"uint16\"}],\"name\":\"composeQueue\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"messageHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"srcEid\",\"type\":\"uint32\"}],\"name\":\"defaultReceiveLibrary\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"lib\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"srcEid\",\"type\":\"uint32\"}],\"name\":\"defaultReceiveLibraryTimeout\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"lib\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"dstEid\",\"type\":\"uint32\"}],\"name\":\"defaultSendLibrary\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"lib\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"oapp\",\"type\":\"address\"}],\"name\":\"delegates\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"delegate\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"eid\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_oapp\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_lib\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"_eid\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"_configType\",\"type\":\"uint32\"}],\"name\":\"getConfig\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"config\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"_srcEid\",\"type\":\"uint32\"}],\"name\":\"getReceiveLibrary\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"lib\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isDefault\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRegisteredLibraries\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSendContext\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"_dstEid\",\"type\":\"uint32\"}],\"name\":\"getSendLibrary\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"lib\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"_srcEid\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"_sender\",\"type\":\"bytes32\"}],\"name\":\"inboundNonce\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"srcEid\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"sender\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"inboundNonce\",\"type\":\"uint64\"}],\"name\":\"inboundPayloadHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"payloadHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"srcEid\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"sender\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"}],\"internalType\":\"structOrigin\",\"name\":\"_origin\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"}],\"name\":\"initializable\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"_dstEid\",\"type\":\"uint32\"}],\"name\":\"isDefaultSendLibrary\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"lib\",\"type\":\"address\"}],\"name\":\"isRegisteredLibrary\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isSendingMessage\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_eid\",\"type\":\"uint32\"}],\"name\":\"isSupportedEid\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"_srcEid\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"_actualReceiveLib\",\"type\":\"address\"}],\"name\":\"isValidReceiveLibrary\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"srcEid\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"sender\",\"type\":\"bytes32\"}],\"name\":\"lazyInboundNonce\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_guid\",\"type\":\"bytes32\"},{\"internalType\":\"uint16\",\"name\":\"_index\",\"type\":\"uint16\"},{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_extraData\",\"type\":\"bytes\"}],\"name\":\"lzCompose\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_guid\",\"type\":\"bytes32\"},{\"internalType\":\"uint16\",\"name\":\"_index\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"_gas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_extraData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_reason\",\"type\":\"bytes\"}],\"name\":\"lzComposeAlert\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"srcEid\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"sender\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"}],\"internalType\":\"structOrigin\",\"name\":\"_origin\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_guid\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_extraData\",\"type\":\"bytes\"}],\"name\":\"lzReceive\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"srcEid\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"sender\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"}],\"internalType\":\"structOrigin\",\"name\":\"_origin\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_guid\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_gas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_extraData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_reason\",\"type\":\"bytes\"}],\"name\":\"lzReceiveAlert\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lzToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nativeToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"_dstEid\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"_receiver\",\"type\":\"bytes32\"}],\"name\":\"nextGuid\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_oapp\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"_srcEid\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"_sender\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"_nonce\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"_payloadHash\",\"type\":\"bytes32\"}],\"name\":\"nilify\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"dstEid\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"receiver\",\"type\":\"bytes32\"}],\"name\":\"outboundNonce\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"dstEid\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"receiver\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"options\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"payInLzToken\",\"type\":\"bool\"}],\"internalType\":\"structMessagingParams\",\"name\":\"_params\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"}],\"name\":\"quote\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"nativeFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lzTokenFee\",\"type\":\"uint256\"}],\"internalType\":\"structMessagingFee\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"srcEid\",\"type\":\"uint32\"}],\"name\":\"receiveLibraryTimeout\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"lib\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"recoverToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_lib\",\"type\":\"address\"}],\"name\":\"registerLibrary\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"dstEid\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"receiver\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"options\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"payInLzToken\",\"type\":\"bool\"}],\"internalType\":\"structMessagingParams\",\"name\":\"_params\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"_refundAddress\",\"type\":\"address\"}],\"name\":\"send\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"guid\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"nativeFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lzTokenFee\",\"type\":\"uint256\"}],\"internalType\":\"structMessagingFee\",\"name\":\"fee\",\"type\":\"tuple\"}],\"internalType\":\"structMessagingReceipt\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_guid\",\"type\":\"bytes32\"},{\"internalType\":\"uint16\",\"name\":\"_index\",\"type\":\"uint16\"},{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"}],\"name\":\"sendCompose\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_oapp\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_lib\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"eid\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"configType\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"config\",\"type\":\"bytes\"}],\"internalType\":\"structSetConfigParam[]\",\"name\":\"_params\",\"type\":\"tuple[]\"}],\"name\":\"setConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_eid\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"_newLib\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_gracePeriod\",\"type\":\"uint256\"}],\"name\":\"setDefaultReceiveLibrary\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_eid\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"_lib\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_expiry\",\"type\":\"uint256\"}],\"name\":\"setDefaultReceiveLibraryTimeout\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_eid\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"_newLib\",\"type\":\"address\"}],\"name\":\"setDefaultSendLibrary\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_delegate\",\"type\":\"address\"}],\"name\":\"setDelegate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_lzToken\",\"type\":\"address\"}],\"name\":\"setLzToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_paused\",\"type\":\"bool\"}],\"name\":\"setPaused\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_oapp\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"_eid\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"_newLib\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_gracePeriod\",\"type\":\"uint256\"}],\"name\":\"setReceiveLibrary\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_oapp\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"_eid\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"_lib\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_expiry\",\"type\":\"uint256\"}],\"name\":\"setReceiveLibraryTimeout\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_oapp\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"_eid\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"_newLib\",\"type\":\"address\"}],\"name\":\"setSendLibrary\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_oapp\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"_srcEid\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"_sender\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"_nonce\",\"type\":\"uint64\"}],\"name\":\"skip\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"srcEid\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"sender\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"}],\"internalType\":\"structOrigin\",\"name\":\"_origin\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"}],\"name\":\"verifiable\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"srcEid\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"sender\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"}],\"internalType\":\"structOrigin\",\"name\":\"_origin\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_payloadHash\",\"type\":\"bytes32\"}],\"name\":\"verify\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60c06040526001600d5534801562000015575f80fd5b50604051620056bd380380620056bd8339810160408190526200003891620002d6565b816200004433620000ac565b63ffffffff166080526040516200005b90620002c8565b604051809103905ff08015801562000075573d5f803e3d5ffd5b506001600160a01b031660a08190526200008f90620000fb565b600e805460ff19169055620000a481620000ac565b50506200034c565b5f80546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b6200010562000268565b6040516301ffc9a760e01b81526325fc096160e21b60048201526001600160a01b038216906301ffc9a790602401602060405180830381865afa1580156200014f573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019062000175919062000324565b620001935760405163eb64f35d60e01b815260040160405180910390fd5b6001600160a01b0381165f9081526005602052604090205460ff1615620001cd5760405163457517f360e11b815260040160405180910390fd5b6001600160a01b0381165f818152600560209081526040808320805460ff191660019081179091556004805491820181559093527f8a35acfbc15ff81a39ae7d344fd709f28e8600b4aa8c65c6b64bfe7fe36bd19b90920180546001600160a01b0319168417905590519182527f6b374d56679ca9463f27c85c6311e2bb7fde69bf201d3da39d53f10bd9d78af5910160405180910390a150565b5f546001600160a01b03163314620002c65760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015260640160405180910390fd5b565b6101d580620054e883390190565b5f8060408385031215620002e8575f80fd5b825163ffffffff81168114620002fc575f80fd5b60208401519092506001600160a01b038116811462000319575f80fd5b809150509250929050565b5f6020828403121562000335575f80fd5b8151801515811462000345575f80fd5b9392505050565b60805160a051615157620003915f395f6107ad01525f81816105590152818161276801528181612f4d01528181612fb30152818161367e01526136e401526151575ff3fe608060405260043610610366575f3560e01c80638da5cb5b116101c8578063c9a54a99116100fd578063dc93c8a21161009d578063e8964e811161006d578063e8964e8114610b9e578063ef667aa114610bbd578063f2fde38b14610c06578063f64be4c714610c25575f80fd5b8063dc93c8a214610af2578063ddc28c5814610b3c578063e1758bd814610b68578063e4fe1d9414610b7a575f80fd5b8063cb5026b9116100d8578063cb5026b914610a73578063d4b4ec8f14610a86578063d70b890214610aa5578063dc706a6214610ac4575f80fd5b8063c9a54a99146109f0578063c9fc7bcd14610a0f578063ca5eb5e114610a54575f80fd5b8063a718531b11610168578063aafe5e0711610143578063aafe5e0714610974578063aafea31214610993578063b96a277f146109b2578063c28e0eed146109d1575f80fd5b8063a718531b14610917578063a7229fd914610936578063a825d74714610955575f80fd5b80639535ff30116101a35780639535ff30146108755780639c6d7340146108945780639d7f9775146108d9578063a0dd43fc146108f8575f80fd5b80638da5cb5b146108255780639132e5c31461084157806391d20fa114610862575f80fd5b80635b17bb701161029e5780636e83f5bb1161023e5780637331809111610219578063733180911461079c57806379624ca9146107cf5780637cb59012146107e7578063861e1ca514610806575f80fd5b80636e83f5bb146106f75780636f50a80314610754578063715018a614610788575f80fd5b8063697fe6b611610279578063697fe6b61461067b5780636a14d7151461069a5780636bf73fa3146106b95780636dbd9f90146106d8575f80fd5b80635b17bb70146105dc5780635c975abb146106395780636750cd4c1461065c575f80fd5b80632baf0be711610309578063402f8468116102e4578063402f8468146104eb57806340f8068314610529578063416ecebf14610548578063587cde1e14610590575f80fd5b80632baf0be7146104655780632e80fbf31461048757806335d330b0146104a6575f80fd5b8063183c834f11610344578063183c834f146103db5780632637a450146103fa5780632a56c1b01461041a5780632b3197b914610439575f80fd5b80630c0c389e1461036a57806314f651a91461037f57806316c38b3c146103bc575b5f80fd5b61037d61037836600461409f565b610c59565b005b34801561038a575f80fd5b50610393610d53565b6040805163ffffffff90931683526001600160a01b039091166020830152015b60405180910390f35b3480156103c7575f80fd5b5061037d6103d6366004614140565b610d83565b3480156103e6575f80fd5b5061037d6103f536600461416e565b610da4565b61040d6104083660046141b6565b611041565b6040516103b39190614204565b348015610425575f80fd5b5061037d610434366004614245565b611176565b348015610444575f80fd5b506104586104533660046142af565b611215565b6040516103b3919061434d565b348015610470575f80fd5b506104795f1981565b6040519081526020016103b3565b348015610492575f80fd5b5061037d6104a1366004614375565b6112db565b3480156104b1575f80fd5b506104796104c03660046143d7565b600c60209081525f948552604080862082529385528385208152918452828420909152825290205481565b3480156104f6575f80fd5b5061050a610505366004614416565b611446565b604080516001600160a01b0390931683529015156020830152016103b3565b348015610534575f80fd5b5061037d610543366004614375565b6114c2565b348015610553575f80fd5b5061057b7f000000000000000000000000000000000000000000000000000000000000000081565b60405163ffffffff90911681526020016103b3565b34801561059b575f80fd5b506105c46105aa36600461443e565b600f6020525f90815260409020546001600160a01b031681565b6040516001600160a01b0390911681526020016103b3565b3480156105e7575f80fd5b506106216105f6366004614457565b600160209081525f93845260408085208252928452828420905282529020546001600160401b031681565b6040516001600160401b0390911681526020016103b3565b348015610644575f80fd5b50600e5460ff165b60405190151581526020016103b3565b348015610667575f80fd5b5061064c610676366004614490565b611625565b348015610686575f80fd5b5061037d6106953660046144a9565b611673565b3480156106a5575f80fd5b5061037d6106b436600461416e565b6116e8565b3480156106c4575f80fd5b5061037d6106d336600461459a565b611abb565b3480156106e3575f80fd5b5061037d6106f2366004614673565b611b25565b348015610702575f80fd5b50610735610711366004614490565b600b6020525f9081526040902080546001909101546001600160a01b039091169082565b604080516001600160a01b0390931683526020830191909152016103b3565b34801561075f575f80fd5b506105c461076e366004614490565b600a6020525f90815260409020546001600160a01b031681565b348015610793575f80fd5b5061037d611bc1565b3480156107a7575f80fd5b506105c47f000000000000000000000000000000000000000000000000000000000000000081565b3480156107da575f80fd5b50600d546001141561064c565b3480156107f2575f80fd5b5061037d6108013660046146fc565b611bd4565b348015610811575f80fd5b5061064c610820366004614748565b611cbe565b348015610830575f80fd5b505f546001600160a01b03166105c4565b34801561084c575f80fd5b50610855611d28565b6040516103b39190614771565b61037d6108703660046147bd565b611d88565b348015610880575f80fd5b5061037d61088f366004614860565b611f12565b34801561089f575f80fd5b506106216108ae366004614457565b600360209081525f93845260408085208252928452828420905282529020546001600160401b031681565b3480156108e4575f80fd5b5061064c6108f3366004614860565b61216f565b348015610903575f80fd5b50610621610912366004614457565b612243565b348015610922575f80fd5b5061037d6109313660046148a0565b6122a2565b348015610941575f80fd5b5061037d6109503660046148c9565b6125b8565b348015610960575f80fd5b5061037d61096f3660046148e4565b6125d0565b34801561097f575f80fd5b5061047961098e366004614457565b612719565b34801561099e575f80fd5b5061037d6109ad36600461491e565b612798565b3480156109bd575f80fd5b506105c46109cc366004614416565b6129ba565b3480156109dc575f80fd5b5061037d6109eb36600461443e565b612a27565b3480156109fb575f80fd5b5061064c610a0a366004614748565b612a8c565b348015610a1a575f80fd5b50610479610a29366004614938565b600260209081525f948552604080862082529385528385208152918452828420909152825290205481565b348015610a5f575f80fd5b5061037d610a6e36600461443e565b612aef565b348015610a7e575f80fd5b506104795f81565b348015610a91575f80fd5b5061037d610aa03660046148a0565b612b4e565b348015610ab0575f80fd5b5061037d610abf366004614938565b612d83565b348015610acf575f80fd5b5061064c610ade36600461443e565b60056020525f908152604090205460ff1681565b348015610afd575f80fd5b5061064c610b0c366004614416565b6001600160a01b039182165f90815260066020908152604080832063ffffffff9490941683529290522054161590565b348015610b47575f80fd5b50610b5b610b563660046141b6565b612e71565b6040516103b39190614977565b348015610b73575f80fd5b505f6105c4565b348015610b85575f80fd5b50600e546105c49061010090046001600160a01b031681565b348015610ba9575f80fd5b5061037d610bb836600461443e565b6130d9565b348015610bc8575f80fd5b50610735610bd7366004614416565b600860209081525f9283526040808420909152908252902080546001909101546001600160a01b039091169082565b348015610c11575f80fd5b5061037d610c2036600461443e565b613239565b348015610c30575f80fd5b506105c4610c3f366004614490565b60096020525f90815260409020546001600160a01b031681565b610ca886610c6a60208a018a614490565b60208a0135610c7f60608c0160408d0161498e565b898989604051602001610c94939291906149a7565b6040516020818303038152906040526132af565b506040516313137d6560e01b81526001600160a01b038716906313137d65903490610ce3908b908a908a908a9033908b908b90600401614a21565b5f604051808303818588803b158015610cfa575f80fd5b505af1158015610d0c573d5f803e3d5ffd5b50505050507f3cd5e48f9730b129dc7550f0fcea9c767b7be37837cd10e55eb35f734f4bca048787604051610d42929190614a77565b60405180910390a150505050505050565b5f80610d62600d546001141590565b610d6d575f80610d7b565b610d7b600d5460a081901c91565b915091509091565b610d8b613466565b8015610d9c57610d996134bf565b50565b610d99613519565b6001600160a01b0382165f90815260056020526040902054829060ff16610dde57604051631bc58ef360e11b815260040160405180910390fd5b826001600160a01b03811615610e7e575f816001600160a01b0316631881d94d6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610e2b573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610e4f9190614ab1565b6002811115610e6057610e60614a9d565b03610e7e576040516342756b1b60e11b815260040160405180910390fd5b83856001600160a01b03821615610f18576040516319d4335360e21b815263ffffffff821660048201526001600160a01b03831690636750cd4c90602401602060405180830381865afa158015610ed7573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610efb9190614acf565b610f1857604051630e476aa960e21b815260040160405180910390fd5b610f2188613552565b5f610f2c8989611446565b9150508015610f4e57604051633c075f7560e01b815260040160405180910390fd5b855f03610f94576001600160a01b0389165f90815260086020908152604080832063ffffffff8c168452909152812080546001600160a01b031916815560010155610ff9565b438611610fb4576040516302efcf9160e11b815260040160405180910390fd5b6001600160a01b038981165f90815260086020908152604080832063ffffffff8d168452909152902080546001600160a01b0319169189169190911781556001018690555b7f4e0a5bbfa0c11a64effb1ada324b5437a17272e1aed9320398715ef71bb209288989898960405161102e9493929190614aea565b60405180910390a1505050505050505050565b611049613ff3565b6110516135a2565b61105e6020840184614490565b336001600d54146110825760405163ee120b0960e01b815260040160405180910390fd5b63ffffffff60a01b60a083811b919091166001600160a01b03831617600d556110b090860160808701614140565b80156110ca5750600e5461010090046001600160a01b0316155b156110e857604051632d7b695560e11b815260040160405180910390fd5b5f806110f433886135e8565b915091505f6111003490565b90505f61111b61111660a08b0160808c01614140565b613888565b905061112c84604001518383613924565b600e546040850151602001516111539161010090046001600160a01b03169083868c613972565b6040840151516111659083858b61399f565b50506001600d555095945050505050565b61117f85613552565b5f838383604051602001611195939291906149a7565b60408051601f1981840301815291905290506111d3866111b86020880188614490565b60208801356111cd60608a0160408b0161498e565b856132af565b507f3cd5e48f9730b129dc7550f0fcea9c767b7be37837cd10e55eb35f734f4bca048587604051611205929190614a77565b60405180910390a1505050505050565b6001600160a01b0383165f90815260056020526040902054606090849060ff1661125257604051631bc58ef360e11b815260040160405180910390fd5b604051639c33abf760e01b815263ffffffff80861660048301526001600160a01b0388811660248401529085166044830152861690639c33abf7906064015f60405180830381865afa1580156112aa573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f191682016040526112d19190810190614ba8565b9695505050505050565b6112e485613552565b6001600160a01b0385165f90815260026020908152604080832063ffffffff88168452825280832086845282528083206001600160401b038616845290915290205481811461135557604051637182306f60e01b815260048101829052602481018390526044015b60405180910390fd5b6001600160a01b0386165f90815260016020908152604080832063ffffffff8916845282528083208784529091529020546001600160401b03908116908416118015906113a0575080155b156113c957604051630c09b63560e41b81526001600160401b038416600482015260240161134c565b6001600160a01b0386165f90815260026020908152604080832063ffffffff89168452825280832087845282528083206001600160401b0387168452909152908190205f199055517faf0450c392c4f702515a457a362328c8aa21916048ca6d0419e248b30cb552929061120590879087908a9088908890614bd9565b6001600160a01b038083165f90815260076020908152604080832063ffffffff8616845290915281205490911690816114bb5763ffffffff83165f908152600a60205260409020546001600160a01b03169150816114b757604051633c74268360e11b815260040160405180910390fd5b5060015b9250929050565b6114cb85613552565b6001600160a01b0385165f90815260026020908152604080832063ffffffff88168452825280832086845282528083206001600160401b038616845290915290205481811461153757604051637182306f60e01b8152600481018290526024810183905260440161134c565b80158061158057506001600160a01b0386165f90815260016020908152604080832063ffffffff8916845282528083208784529091529020546001600160401b03908116908416115b156115a957604051630c09b63560e41b81526001600160401b038416600482015260240161134c565b6001600160a01b0386165f90815260026020908152604080832063ffffffff89168452825280832087845282528083206001600160401b038716845290915280822091909155517f7f68a37a6e69a0de35024a234558f9efe4b33b58657753d21eaaa82d51c3510e9061120590879087908a9088908890614bd9565b63ffffffff81165f908152600960205260408120546001600160a01b03161580159061166d575063ffffffff82165f908152600a60205260409020546001600160a01b031615155b92915050565b336001600160a01b03168b6001600160a01b03168d6001600160a01b03167f8a0b1dce321c5c5fb42349bce46d18087c04140de520917661fb923e44a904b98d8d8d8d8d8d8d8d8d8d6040516116d29a99989796959493929190614c16565b60405180910390a4505050505050505050505050565b6001600160a01b0382165f90815260056020526040902054829060ff1615801561171a57506001600160a01b03811615155b156117385760405163a4ff2ec360e01b815260040160405180910390fd5b826001600160a01b038116156117d8575f816001600160a01b0316631881d94d6040518163ffffffff1660e01b8152600401602060405180830381865afa158015611785573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906117a99190614ab1565b60028111156117ba576117ba614a9d565b036117d8576040516342756b1b60e11b815260040160405180910390fd5b83856001600160a01b03821615611872576040516319d4335360e21b815263ffffffff821660048201526001600160a01b03831690636750cd4c90602401602060405180830381865afa158015611831573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906118559190614acf565b61187257604051630e476aa960e21b815260040160405180910390fd5b61187b88613552565b6001600160a01b038089165f90815260076020908152604080832063ffffffff8c168452909152902054811690871681036118c95760405163d0ecb66b60e01b815260040160405180910390fd5b6001600160a01b038981165f81815260076020908152604080832063ffffffff8e168085529083529281902080546001600160a01b031916958d1695861790558051938452908301919091528101919091527fcd6f92f5ac6185a5acfa02c92090746cec64d777269cbcd0ed031e396657a1c29060600160405180910390a18515611a40576001600160a01b038116158061196b57506001600160a01b038716155b1561198957604051633c075f7560e01b815260040160405180910390fd5b5f6040518060400160405280836001600160a01b0316815260200188436119b09190614c92565b90526001600160a01b038b81165f90815260086020908152604080832063ffffffff8f168452825291829020845181546001600160a01b03191694169390931783558301516001909201829055519192507f4e0a5bbfa0c11a64effb1ada324b5437a17272e1aed9320398715ef71bb2092891611a32918d918d918791614aea565b60405180910390a150611ab0565b6001600160a01b0389165f90815260086020908152604080832063ffffffff8c16845290915280822080546001600160a01b0319168155600101829055517f4e0a5bbfa0c11a64effb1ada324b5437a17272e1aed9320398715ef71bb209289161102e918c918c91869190614aea565b505050505050505050565b336001600160a01b03168a6001600160a01b03167f7edfa10fe10193301ad8a8bea7e968c7bcabcc64981f368e3aeada40ce26ae2c8d8c8c8c8c8c8c8c8c8c604051611b109a99989796959493929190614ca5565b60405180910390a35050505050505050505050565b6001600160a01b0383165f90815260056020526040902054839060ff16611b5f57604051631bc58ef360e11b815260040160405180910390fd5b611b6885613552565b604051631077eb9160e11b81526001600160a01b038516906320efd72290611b9890889087908790600401614d03565b5f604051808303815f87803b158015611baf575f80fd5b505af1158015611ab0573d5f803e3d5ffd5b611bc9613466565b611bd25f6139c9565b565b335f908152600c602090815260408083206001600160a01b03891684528252808320878452825280832061ffff8716845290915290205415611c2957604051630542086560e21b815260040160405180910390fd5b8181604051611c39929190614df1565b60408051918290038220335f818152600c60209081528482206001600160a01b038c16835281528482208a8352815284822061ffff8a168352905292909220557f3d52ff888d033fd3dd1d8057da59e850c91d91a72c41dfa445b247dfedeb6dc191611caf919088908890889088908890614e00565b60405180910390a15050505050565b6001600160a01b0381165f9081526001602090815260408220611d2191859185918590611ced90850185614490565b63ffffffff16815260208082019290925260409081015f9081208984013582529092529020546001600160401b0316613a18565b9392505050565b60606004805480602002602001604051908101604052809291908181526020018280548015611d7e57602002820191905f5260205f20905b81546001600160a01b03168152600190910190602001808311611d60575b5050505050905090565b6001600160a01b038089165f908152600c60209081526040808320938b168352928152828220898352815282822061ffff891683529052818120549151611dd29087908790614df1565b60405180910390209050808214611e06576040516335ca595f60e01b8152600481018390526024810182905260440161134c565b6001600160a01b03808b165f908152600c60209081526040808320938d168084529382528083208c8452825280832061ffff8c168452909152908190206001905551630685081360e51b815263d0a10260903490611e74908e908d908c908c9033908d908d90600401614e49565b5f604051808303818588803b158015611e8b575f80fd5b505af1158015611e9d573d5f803e3d5ffd5b50505050507e36c98efcf9e6641dfbc9051f66f405253e8e0c2ab4a24dccda15595b7378c88a8a8a8a604051611efe94939291906001600160a01b039485168152929093166020830152604082015261ffff91909116606082015260800190565b60405180910390a150505050505050505050565b6001600160a01b0381165f90815260056020526040902054819060ff16158015611f4457506001600160a01b03811615155b15611f625760405163a4ff2ec360e01b815260040160405180910390fd5b816001600160a01b03811615612003576001816001600160a01b0316631881d94d6040518163ffffffff1660e01b8152600401602060405180830381865afa158015611fb0573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611fd49190614ab1565b6002811115611fe557611fe5614a9d565b0361200357604051633d00f6f160e11b815260040160405180910390fd5b82846001600160a01b0382161561209d576040516319d4335360e21b815263ffffffff821660048201526001600160a01b03831690636750cd4c90602401602060405180830381865afa15801561205c573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906120809190614acf565b61209d57604051630e476aa960e21b815260040160405180910390fd5b6120a687613552565b6001600160a01b038781165f90815260066020908152604080832063ffffffff8b1684529091529020548187169116036120f35760405163d0ecb66b60e01b815260040160405180910390fd5b6001600160a01b038781165f81815260066020908152604080832063ffffffff8c168085529083529281902080546001600160a01b031916958b1695861790558051938452908301919091528101919091527f4cff966ebee29a156dcb34cf72c1d06231fb1777f6bdf6e8089819232f002b1c90606001610d42565b5f805f61217c8686611446565b91509150816001600160a01b0316846001600160a01b0316036121a457600192505050611d21565b5f816121d7576001600160a01b0387165f90815260086020908152604080832063ffffffff8a16845290915290206121ec565b63ffffffff86165f908152600b602052604090205b6040805180820190915281546001600160a01b03908116808352600190930154602083015290925086161480156122265750438160200151115b156122375760019350505050611d21565b505f9695505050505050565b6001600160a01b0383165f90815260016020908152604080832063ffffffff8616845282528083208484529091528120546001600160401b03165b61228d85858584600101613a99565b1561229a5760010161227e565b949350505050565b6122aa613466565b6001600160a01b0382165f90815260056020526040902054829060ff166122e457604051631bc58ef360e11b815260040160405180910390fd5b826001600160a01b03811615612384575f816001600160a01b0316631881d94d6040518163ffffffff1660e01b8152600401602060405180830381865afa158015612331573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906123559190614ab1565b600281111561236657612366614a9d565b03612384576040516342756b1b60e11b815260040160405180910390fd5b83856001600160a01b0382161561241e576040516319d4335360e21b815263ffffffff821660048201526001600160a01b03831690636750cd4c90602401602060405180830381865afa1580156123dd573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906124019190614acf565b61241e57604051630e476aa960e21b815260040160405180910390fd5b63ffffffff87165f908152600a60205260409020546001600160a01b0390811690871681036124605760405163d0ecb66b60e01b815260040160405180910390fd5b63ffffffff88165f818152600a602090815260409182902080546001600160a01b0319166001600160a01b038c169081179091558251938452908301527fc16891855cffb4a5ac51ac11864a3f3c96ba816cc45fe686c987ae36277de5ec910160405180910390a1851561254c5763ffffffff88165f908152600b6020526040902080546001600160a01b0319166001600160a01b0383161781556125058743614c92565b600182018190556040517f55b28633cdb29709386f555dfc54418592ad475ce7a65a78ac5928af60ffb8f89161253e918c918691614e9b565b60405180910390a1506125ae565b63ffffffff88165f908152600b602052604080822080546001600160a01b0319168155600101829055517f55b28633cdb29709386f555dfc54418592ad475ce7a65a78ac5928af60ffb8f8916125a5918b918591614e9b565b60405180910390a15b5050505050505050565b6125c0613466565b6125cb838383613ae5565b505050565b6125e7826125e16020860186614490565b3361216f565b612604576040516313e9bb2b60e21b815260040160405180910390fd5b6001600160a01b0382165f908152600160209081526040822090829061262c90870187614490565b63ffffffff16815260208082019290925260409081015f9081208784013582529092529020546001600160401b03169050612668848483613a18565b6126855760405163751cb20f60e01b815260040160405180910390fd5b612690848483613b08565b6126ac5760405162bbf0e560e11b815260040160405180910390fd5b6126d8836126bd6020870187614490565b60208701356126d26060890160408a0161498e565b86613bbe565b7f0d87345f3d1c929caba93e1c3821b54ff3512e12b66aa3cfe54b6bcbc17e59b484848460405161270b93929190614ec2565b60405180910390a150505050565b6001600160a01b0383165f90815260036020908152604080832063ffffffff8616845282528083208484529091528120548190612760906001600160401b03166001614eea565b905061278f817f0000000000000000000000000000000000000000000000000000000000000000878787613c23565b95945050505050565b6127a0613466565b6001600160a01b0381165f90815260056020526040902054819060ff166127da57604051631bc58ef360e11b815260040160405180910390fd5b816001600160a01b0381161561287b576001816001600160a01b0316631881d94d6040518163ffffffff1660e01b8152600401602060405180830381865afa158015612828573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061284c9190614ab1565b600281111561285d5761285d614a9d565b0361287b57604051633d00f6f160e11b815260040160405180910390fd5b82846001600160a01b03821615612915576040516319d4335360e21b815263ffffffff821660048201526001600160a01b03831690636750cd4c90602401602060405180830381865afa1580156128d4573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906128f89190614acf565b61291557604051630e476aa960e21b815260040160405180910390fd5b63ffffffff86165f908152600960205260409020546001600160a01b038087169116036129555760405163d0ecb66b60e01b815260040160405180910390fd5b63ffffffff86165f8181526009602090815260409182902080546001600160a01b0319166001600160a01b038a169081179091558251938452908301527f16aa0f528038ab41019e95bae5b418a50ba8532c5800e3b7ea2f517d3fa625f59101611205565b6001600160a01b038083165f90815260066020908152604080832063ffffffff86168452909152902054168061166d575063ffffffff81165f908152600960205260409020546001600160a01b03168061166d57604051636c1ccdb560e01b815260040160405180910390fd5b612a2f613466565b600e8054610100600160a81b0319166101006001600160a01b038416908102919091179091556040519081527fd476ec5ec1ac11cec3714d41e7ea49419471aceb9bd0dff1becfc3e363a62396906020015b60405180910390a150565b6001600160a01b0381165f9081526001602090815260408220611d2191859185918590612abb90850185614490565b63ffffffff16815260208082019290925260409081015f9081208984013582529092529020546001600160401b0316613b08565b335f818152600f602090815260409182902080546001600160a01b0319166001600160a01b0386169081179091558251938452908301527f6ee10e9ed4d6ce9742703a498707862f4b00f1396a87195eb93267b3d79839819101612a81565b6001600160a01b0382165f90815260056020526040902054829060ff16612b8857604051631bc58ef360e11b815260040160405180910390fd5b826001600160a01b03811615612c28575f816001600160a01b0316631881d94d6040518163ffffffff1660e01b8152600401602060405180830381865afa158015612bd5573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190612bf99190614ab1565b6002811115612c0a57612c0a614a9d565b03612c28576040516342756b1b60e11b815260040160405180910390fd5b83856001600160a01b03821615612cc2576040516319d4335360e21b815263ffffffff821660048201526001600160a01b03831690636750cd4c90602401602060405180830381865afa158015612c81573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190612ca59190614acf565b612cc257604051630e476aa960e21b815260040160405180910390fd5b612cca613466565b845f03612cfc5763ffffffff87165f908152600b6020526040812080546001600160a01b031916815560010155612d50565b438511612d1c576040516302efcf9160e11b815260040160405180910390fd5b63ffffffff87165f908152600b6020526040902080546001600160a01b0319166001600160a01b0388161781556001018590555b7f55b28633cdb29709386f555dfc54418592ad475ce7a65a78ac5928af60ffb8f8878787604051610d4293929190614e9b565b612d8c84613552565b612d97848484612243565b612da2906001614eea565b6001600160401b0316816001600160401b031614612dde57604051630c09b63560e41b81526001600160401b038216600482015260240161134c565b6001600160a01b0384165f81815260016020908152604080832063ffffffff8816808552908352818420878552835292819020805467ffffffffffffffff19166001600160401b038716908117909155815193845291830186905282019290925260608101919091527f28f40053783033ef755556a0c3315379141f51a33aed8334174ffbadd90bde489060800161270b565b604080518082019091525f8082526020820152612e9460a0840160808501614140565b8015612eae5750600e5461010090046001600160a01b0316155b15612ecc57604051632d7b695560e11b815260040160405180910390fd5b6001600160a01b0382165f9081526003602090815260408220908290612ef490870187614490565b63ffffffff16815260208082019290925260409081015f908120878401358252909252902054612f2e906001600160401b03166001614eea565b90505f6040518060e00160405280836001600160401b031681526020017f000000000000000000000000000000000000000000000000000000000000000063ffffffff168152602001856001600160a01b03168152602001865f016020810190612f989190614490565b63ffffffff16815260200186602001358152602001612fef847f0000000000000000000000000000000000000000000000000000000000000000888a5f016020810190612fe59190614490565b8b60200135613c23565b81526020016130016040880188614f0a565b8080601f0160208091040260200160405190810160405280939291908181526020018383808284375f9201829052509390945250929350915061304d9050856109cc6020890189614490565b90506001600160a01b03811663d80e9bd98361306c60608a018a614f0a565b61307c60a08c0160808d01614140565b6040518563ffffffff1660e01b815260040161309b9493929190614f4c565b6040805180830381865afa1580156130b5573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906112d1919061503d565b6130e1613466565b6040516301ffc9a760e01b81526325fc096160e21b60048201526001600160a01b038216906301ffc9a790602401602060405180830381865afa15801561312a573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061314e9190614acf565b61316b5760405163eb64f35d60e01b815260040160405180910390fd5b6001600160a01b0381165f9081526005602052604090205460ff16156131a45760405163457517f360e11b815260040160405180910390fd5b6001600160a01b0381165f818152600560209081526040808320805460ff191660019081179091556004805491820181559093527f8a35acfbc15ff81a39ae7d344fd709f28e8600b4aa8c65c6b64bfe7fe36bd19b90920180546001600160a01b0319168417905590519182527f6b374d56679ca9463f27c85c6311e2bb7fde69bf201d3da39d53f10bd9d78af59101612a81565b613241613466565b6001600160a01b0381166132a65760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b606482015260840161134c565b610d99816139c9565b6001600160a01b0385165f90815260016020908152604080832063ffffffff8816845282528083208684529091528120546001600160401b0390811690841681101561339b57600181015b846001600160401b0316816001600160401b03161161334f5761331f88888884613a99565b61334757604051630c09b63560e41b81526001600160401b038216600482015260240161134c565b6001016132fa565b506001600160a01b0387165f90815260016020908152604080832063ffffffff8a16845282528083208884529091529020805467ffffffffffffffff19166001600160401b0386161790555b82516020808501919091206001600160a01b0389165f90815260028352604080822063ffffffff8b168352845280822089835284528082206001600160401b0389168352909352919091205490925080831461341457604051637182306f60e01b8152600481018290526024810184905260440161134c565b50506001600160a01b039095165f90815260026020908152604080832063ffffffff90971683529581528582209482529384528481206001600160401b03909316815291909252918220919091555090565b5f546001600160a01b03163314611bd25760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015260640161134c565b6134c76135a2565b600e805460ff191660011790557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a2586134fc3390565b6040516001600160a01b03909116815260200160405180910390a1565b613521613c9b565b600e805460ff191690557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa336134fc565b336001600160a01b0382161480159061358457506001600160a01b038181165f908152600f6020526040902054163314155b15610d995760405163c4c5259360e01b815260040160405180910390fd5b600e5460ff1615611bd25760405162461bcd60e51b815260206004820152601060248201526f14185d5cd8589b194e881c185d5cd95960821b604482015260640161134c565b6135f0613ff3565b5f8061365f856136036020870187614490565b6001600160a01b03919091165f90815260036020908152604080832063ffffffff949094168352928152828220818901358352905220805467ffffffffffffffff19811660016001600160401b03928316019182161790915590565b90505f6040518060e00160405280836001600160401b031681526020017f000000000000000000000000000000000000000000000000000000000000000063ffffffff168152602001876001600160a01b03168152602001865f0160208101906136c99190614490565b63ffffffff16815260200186602001358152602001613716847f00000000000000000000000000000000000000000000000000000000000000008a8a5f016020810190612fe59190614490565b81526020016137286040880188614f0a565b8080601f0160208091040260200160405190810160405280939291908181526020018383808284375f920182905250939094525092935091506137749050876109cc6020890189614490565b90505f806001600160a01b038316634389e58f8561379560608c018c614f0a565b6137a560a08e0160808f01614140565b6040518563ffffffff1660e01b81526004016137c49493929190614f4c565b5f604051808303815f875af11580156137df573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f191682016040526138069190810190615057565b90925090507f1ab700d4ced0c005b164c0f789fd09fcbb0156d4c2041b8a3bfbcd961cd1567f8161383a60608b018b614f0a565b8660405161384b94939291906150a2565b60405180910390a1506040805160608101825260a09094015184526001600160401b03909416602084015292820192909252925090509250929050565b5f811561391f57600e546040516370a0823160e01b81523060048201526101009091046001600160a01b0316906370a0823190602401602060405180830381865afa1580156138d9573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906138fd91906150e1565b9050805f0361391f57604051631775123760e01b815260040160405180910390fd5b919050565b82518210806139365750808360200151115b156125cb5782516020840151604051634f3ec0d360e01b815260048101929092526024820184905260448201526064810182905260840161134c565b831561398357613983858386613ce4565b82841015613998576139988582868603613ce4565b5050505050565b83156139af576139af8285613d1f565b828410156139c3576139c381858503613d1f565b50505050565b5f80546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b5f80826001600160401b0316118061229a5750604051600162842fc360e01b031981526001600160a01b0384169063ff7bd03d90613a5a9087906004016150f8565b602060405180830381865afa158015613a75573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061229a9190614acf565b6001600160a01b0384165f90815260026020908152604080832063ffffffff87168452825280832085845282528083206001600160401b03851684529091529020541515949350505050565b6001600160a01b038316613afd576125cb8282613d1f565b6125cb838383613ce4565b5f6001600160401b038216613b23606086016040870161498e565b6001600160401b0316118061229a57506001600160a01b0383165f9081526002602090815260408220908290613b5b90880188614490565b63ffffffff1663ffffffff1681526020019081526020015f205f866020013581526020019081526020015f205f866040016020810190613b9b919061498e565b6001600160401b0316815260208101919091526040015f20541415949350505050565b80613bdc576040516304df7fdb60e11b815260040160405180910390fd5b6001600160a01b039094165f90815260026020908152604080832063ffffffff90961683529481528482209382529283528381206001600160401b03909216815291522055565b5f85856001600160a01b03861660405160c09390931b6001600160c01b031916602084015260e091821b6001600160e01b03199081166028850152602c8401919091529085901b16604c8201526050810183905260700160405160208183030381529060405280519060200120905095945050505050565b600e5460ff16611bd25760405162461bcd60e51b815260206004820152601460248201527314185d5cd8589b194e881b9bdd081c185d5cd95960621b604482015260640161134c565b6001600160a01b038216613d0b576040516306b7a93160e41b815260040160405180910390fd5b6125cb6001600160a01b0384168383613dc8565b6001600160a01b038216613d46576040516306b7a93160e41b815260040160405180910390fd5b5f826001600160a01b0316826040515f6040518083038185875af1925050503d805f8114613d8f576040519150601f19603f3d011682016040523d82523d5f602084013e613d94565b606091505b50509050806125cb57604051631196f20d60e21b81526001600160a01b03841660048201526024810183905260440161134c565b604080516001600160a01b03848116602483015260448083018590528351808403909101815260649092018352602080830180516001600160e01b031663a9059cbb60e01b17905283518085019094528084527f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564908401526125cb928692915f91613e57918516908490613ed6565b905080515f1480613e77575080806020019051810190613e779190614acf565b6125cb5760405162461bcd60e51b815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e6044820152691bdd081cdd58d8d9595960b21b606482015260840161134c565b606061229a84845f85855f80866001600160a01b03168587604051613efb9190615106565b5f6040518083038185875af1925050503d805f8114613f35576040519150601f19603f3d011682016040523d82523d5f602084013e613f3a565b606091505b5091509150613f4b87838387613f56565b979650505050505050565b60608315613fc45782515f03613fbd576001600160a01b0385163b613fbd5760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e7472616374000000604482015260640161134c565b508161229a565b61229a8383815115613fd95781518083602001fd5b8060405162461bcd60e51b815260040161134c919061434d565b60405180606001604052805f80191681526020015f6001600160401b0316815260200161403160405180604001604052805f81526020015f81525090565b905290565b5f60608284031215614046575f80fd5b50919050565b80356001600160a01b038116811461391f575f80fd5b5f8083601f840112614072575f80fd5b5081356001600160401b03811115614088575f80fd5b6020830191508360208285010111156114bb575f80fd5b5f805f805f805f60e0888a0312156140b5575f80fd5b6140bf8989614036565b96506140cd6060890161404c565b95506080880135945060a08801356001600160401b03808211156140ef575f80fd5b6140fb8b838c01614062565b909650945060c08a0135915080821115614113575f80fd5b506141208a828b01614062565b989b979a50959850939692959293505050565b8015158114610d99575f80fd5b5f60208284031215614150575f80fd5b8135611d2181614133565b803563ffffffff8116811461391f575f80fd5b5f805f8060808587031215614181575f80fd5b61418a8561404c565b93506141986020860161415b565b92506141a66040860161404c565b9396929550929360600135925050565b5f80604083850312156141c7575f80fd5b82356001600160401b038111156141dc575f80fd5b830160a081860312156141ed575f80fd5b91506141fb6020840161404c565b90509250929050565b5f608082019050825182526001600160401b036020840151166020830152604083015161423e604084018280518252602090810151910152565b5092915050565b5f805f805f60c08688031215614259575f80fd5b6142628661404c565b94506142718760208801614036565b93506080860135925060a08601356001600160401b03811115614292575f80fd5b61429e88828901614062565b969995985093965092949392505050565b5f805f80608085870312156142c2575f80fd5b6142cb8561404c565b93506142d96020860161404c565b92506142e76040860161415b565b91506142f56060860161415b565b905092959194509250565b5f5b8381101561431a578181015183820152602001614302565b50505f910152565b5f8151808452614339816020860160208601614300565b601f01601f19169290920160200192915050565b602081525f611d216020830184614322565b80356001600160401b038116811461391f575f80fd5b5f805f805f60a08688031215614389575f80fd5b6143928661404c565b94506143a06020870161415b565b9350604086013592506143b56060870161435f565b949793965091946080013592915050565b803561ffff8116811461391f575f80fd5b5f805f80608085870312156143ea575f80fd5b6143f38561404c565b93506144016020860161404c565b9250604085013591506142f5606086016143c6565b5f8060408385031215614427575f80fd5b6144308361404c565b91506141fb6020840161415b565b5f6020828403121561444e575f80fd5b611d218261404c565b5f805f60608486031215614469575f80fd5b6144728461404c565b92506144806020850161415b565b9150604084013590509250925092565b5f602082840312156144a0575f80fd5b611d218261415b565b5f805f805f805f805f805f806101208d8f0312156144c5575f80fd5b6144ce8d61404c565b9b506144dc60208e0161404c565b9a5060408d013599506144f160608e016143c6565b985060808d0135975060a08d013596506001600160401b0360c08e01351115614518575f80fd5b6145288e60c08f01358f01614062565b90965094506001600160401b0360e08e01351115614544575f80fd5b6145548e60e08f01358f01614062565b90945092506001600160401b036101008e01351115614571575f80fd5b6145828e6101008f01358f01614062565b81935080925050509295989b509295989b509295989b565b5f805f805f805f805f805f6101408c8e0312156145b5575f80fd5b6145bf8d8d614036565b9a506145cd60608d0161404c565b995060808c0135985060a08c0135975060c08c013596506001600160401b038060e08e013511156145fc575f80fd5b61460c8e60e08f01358f01614062565b90975095506101008d0135811015614622575f80fd5b6146338e6101008f01358f01614062565b90955093506101208d0135811015614649575f80fd5b5061465b8d6101208e01358e01614062565b81935080925050509295989b509295989b9093969950565b5f805f8060608587031215614686575f80fd5b61468f8561404c565b935061469d6020860161404c565b925060408501356001600160401b03808211156146b8575f80fd5b818701915087601f8301126146cb575f80fd5b8135818111156146d9575f80fd5b8860208260051b85010111156146ed575f80fd5b95989497505060200194505050565b5f805f805f60808688031215614710575f80fd5b6147198661404c565b94506020860135935061472e604087016143c6565b925060608601356001600160401b03811115614292575f80fd5b5f8060808385031215614759575f80fd5b6147638484614036565b91506141fb6060840161404c565b602080825282518282018190525f9190848201906040850190845b818110156147b15783516001600160a01b03168352928401929184019160010161478c565b50909695505050505050565b5f805f805f805f8060c0898b0312156147d4575f80fd5b6147dd8961404c565b97506147eb60208a0161404c565b96506040890135955061480060608a016143c6565b945060808901356001600160401b038082111561481b575f80fd5b6148278c838d01614062565b909650945060a08b013591508082111561483f575f80fd5b5061484c8b828c01614062565b999c989b5096995094979396929594505050565b5f805f60608486031215614872575f80fd5b61487b8461404c565b92506148896020850161415b565b91506148976040850161404c565b90509250925092565b5f805f606084860312156148b2575f80fd5b6148bb8461415b565b92506144806020850161404c565b5f805f606084860312156148db575f80fd5b6148bb8461404c565b5f805f60a084860312156148f6575f80fd5b6149008585614036565b925061490e6060850161404c565b9150608084013590509250925092565b5f806040838503121561492f575f80fd5b6141ed8361415b565b5f805f806080858703121561494b575f80fd5b6149548561404c565b93506149626020860161415b565b9250604085013591506142f56060860161435f565b81518152602080830151908201526040810161166d565b5f6020828403121561499e575f80fd5b611d218261435f565b838152818360208301375f910160200190815292915050565b63ffffffff6149ce8261415b565b168252602081013560208301526001600160401b036149ef6040830161435f565b1660408301525050565b81835281816020850137505f828201602090810191909152601f909101601f19169091010190565b614a2b81896149c0565b86606082015260e060808201525f614a4760e0830187896149f9565b6001600160a01b03861660a084015282810360c0840152614a698185876149f9565b9a9950505050505050505050565b60808101614a8582856149c0565b6001600160a01b039290921660609190910152919050565b634e487b7160e01b5f52602160045260245ffd5b5f60208284031215614ac1575f80fd5b815160038110611d21575f80fd5b5f60208284031215614adf575f80fd5b8151611d2181614133565b6001600160a01b03948516815263ffffffff93909316602084015292166040820152606081019190915260800190565b634e487b7160e01b5f52604160045260245ffd5b5f82601f830112614b3d575f80fd5b81516001600160401b0380821115614b5757614b57614b1a565b604051601f8301601f19908116603f01168101908282118183101715614b7f57614b7f614b1a565b81604052838152866020858801011115614b97575f80fd5b6112d1846020830160208901614300565b5f60208284031215614bb8575f80fd5b81516001600160401b03811115614bcd575f80fd5b61229a84828501614b2e565b63ffffffff95909516855260208501939093526001600160a01b039190911660408401526001600160401b03166060830152608082015260a00190565b8a815261ffff8a16602082015288604082015287606082015260e060808201525f614c4560e08301888a6149f9565b82810360a0840152614c588187896149f9565b905082810360c0840152614c6d8185876149f9565b9d9c50505050505050505050505050565b634e487b7160e01b5f52601160045260245ffd5b8082018082111561166d5761166d614c7e565b5f610120614cb3838e6149c0565b8b60608401528a60808401528960a08401528060c0840152614cd8818401898b6149f9565b905082810360e0840152614ced8187896149f9565b9050828103610100840152614c6d8185876149f9565b6001600160a01b0384168152604060208083018290528282018490525f9190606090818501600587901b8601830188865b89811015614de157888303605f190184528135368c9003605e19018112614d59575f80fd5b8b0163ffffffff80614d6a8361415b565b16855280614d7988840161415b565b16878601525087810135601e19823603018112614d94575f80fd5b0185810190356001600160401b03811115614dad575f80fd5b803603821315614dbb575f80fd5b8789860152614dcd88860182846149f9565b958701959450505090840190600101614d34565b50909a9950505050505050505050565b818382375f9101908152919050565b6001600160a01b038781168252861660208201526040810185905261ffff8416606082015260a0608082018190525f90614e3d90830184866149f9565b98975050505050505050565b5f60018060a01b03808a16835288602084015260a06040840152614e7160a08401888a6149f9565b81871660608501528381036080850152614e8c8186886149f9565b9b9a5050505050505050505050565b63ffffffff9390931683526001600160a01b03919091166020830152604082015260600190565b60a08101614ed082866149c0565b6001600160a01b0393909316606082015260800152919050565b6001600160401b0381811683821601908082111561423e5761423e614c7e565b5f808335601e19843603018112614f1f575f80fd5b8301803591506001600160401b03821115614f38575f80fd5b6020019150368190038213156114bb575f80fd5b606080825285516001600160401b031682820152602086015163ffffffff16608083015260408601516001600160a01b031660a08301528501515f90614f9a60c084018263ffffffff169052565b50608086015160e083015260a086015161010083015260c086015160e0610120840152614fcb610140840182614322565b90508281036020840152614fe08186886149f9565b91505061278f604083018415159052565b5f60408284031215615001575f80fd5b604051604081018181106001600160401b038211171561502357615023614b1a565b604052825181526020928301519281019290925250919050565b5f6040828403121561504d575f80fd5b611d218383614ff1565b5f8060608385031215615068575f80fd5b6150728484614ff1565b915060408301516001600160401b0381111561508c575f80fd5b61509885828601614b2e565b9150509250929050565b606081525f6150b46060830187614322565b82810360208401526150c78186886149f9565b91505060018060a01b038316604083015295945050505050565b5f602082840312156150f1575f80fd5b5051919050565b6060810161166d82846149c0565b5f8251615117818460208701614300565b919091019291505056fea264697066735822122026cec8d43f18d6822d2f6622b3452fd02360070870a01881b27044569befc6cb64736f6c63430008140033608060405234801561000f575f80fd5b506101b88061001d5f395ff3fe608060405234801561000f575f80fd5b506004361061004a575f3560e01c806301ffc9a7146100635780631881d94d1461008b57806354fd4d501461009a5780636750cd4c146100c1575b604051632657b6c360e01b815260040160405180910390fd5b61007661007136600461010b565b6100d5565b60405190151581526020015b60405180910390f35b60026040516100829190610139565b6040805167ffffffffffffffff815260ff6020820152600291810191909152606001610082565b6100766100cf36600461015f565b50600190565b5f6001600160e01b031982166325fc096160e21b148061010557506301ffc9a760e01b6001600160e01b03198316145b92915050565b5f6020828403121561011b575f80fd5b81356001600160e01b031981168114610132575f80fd5b9392505050565b602081016003831061015957634e487b7160e01b5f52602160045260245ffd5b91905290565b5f6020828403121561016f575f80fd5b813563ffffffff81168114610132575f80fdfea2646970667358221220525199bcf8a5ce76d360e528b27754176becd7c82a307706653d688a7761089364736f6c63430008140033",
}

// Endpointv2ABI is the input ABI used to generate the binding from.
// Deprecated: Use Endpointv2MetaData.ABI instead.
var Endpointv2ABI = Endpointv2MetaData.ABI

// Endpointv2Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use Endpointv2MetaData.Bin instead.
var Endpointv2Bin = Endpointv2MetaData.Bin

// DeployEndpointv2 deploys a new Ethereum contract, binding an instance of Endpointv2 to it.
func DeployEndpointv2(auth *bind.TransactOpts, backend bind.ContractBackend, _eid uint32, _owner common.Address) (common.Address, *types.Transaction, *Endpointv2, error) {
	parsed, err := Endpointv2MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(Endpointv2Bin), backend, _eid, _owner)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Endpointv2{Endpointv2Caller: Endpointv2Caller{contract: contract}, Endpointv2Transactor: Endpointv2Transactor{contract: contract}, Endpointv2Filterer: Endpointv2Filterer{contract: contract}}, nil
}

// Endpointv2 is an auto generated Go binding around an Ethereum contract.
type Endpointv2 struct {
	Endpointv2Caller     // Read-only binding to the contract
	Endpointv2Transactor // Write-only binding to the contract
	Endpointv2Filterer   // Log filterer for contract events
}

// Endpointv2Caller is an auto generated read-only Go binding around an Ethereum contract.
type Endpointv2Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Endpointv2Transactor is an auto generated write-only Go binding around an Ethereum contract.
type Endpointv2Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Endpointv2Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Endpointv2Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Endpointv2Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Endpointv2Session struct {
	Contract     *Endpointv2       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Endpointv2CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Endpointv2CallerSession struct {
	Contract *Endpointv2Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// Endpointv2TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Endpointv2TransactorSession struct {
	Contract     *Endpointv2Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// Endpointv2Raw is an auto generated low-level Go binding around an Ethereum contract.
type Endpointv2Raw struct {
	Contract *Endpointv2 // Generic contract binding to access the raw methods on
}

// Endpointv2CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Endpointv2CallerRaw struct {
	Contract *Endpointv2Caller // Generic read-only contract binding to access the raw methods on
}

// Endpointv2TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Endpointv2TransactorRaw struct {
	Contract *Endpointv2Transactor // Generic write-only contract binding to access the raw methods on
}

// NewEndpointv2 creates a new instance of Endpointv2, bound to a specific deployed contract.
func NewEndpointv2(address common.Address, backend bind.ContractBackend) (*Endpointv2, error) {
	contract, err := bindEndpointv2(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Endpointv2{Endpointv2Caller: Endpointv2Caller{contract: contract}, Endpointv2Transactor: Endpointv2Transactor{contract: contract}, Endpointv2Filterer: Endpointv2Filterer{contract: contract}}, nil
}

// NewEndpointv2Caller creates a new read-only instance of Endpointv2, bound to a specific deployed contract.
func NewEndpointv2Caller(address common.Address, caller bind.ContractCaller) (*Endpointv2Caller, error) {
	contract, err := bindEndpointv2(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Endpointv2Caller{contract: contract}, nil
}

// NewEndpointv2Transactor creates a new write-only instance of Endpointv2, bound to a specific deployed contract.
func NewEndpointv2Transactor(address common.Address, transactor bind.ContractTransactor) (*Endpointv2Transactor, error) {
	contract, err := bindEndpointv2(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Endpointv2Transactor{contract: contract}, nil
}

// NewEndpointv2Filterer creates a new log filterer instance of Endpointv2, bound to a specific deployed contract.
func NewEndpointv2Filterer(address common.Address, filterer bind.ContractFilterer) (*Endpointv2Filterer, error) {
	contract, err := bindEndpointv2(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Endpointv2Filterer{contract: contract}, nil
}

// bindEndpointv2 binds a generic wrapper to an already deployed contract.
func bindEndpointv2(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := Endpointv2MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Endpointv2 *Endpointv2Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Endpointv2.Contract.Endpointv2Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Endpointv2 *Endpointv2Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Endpointv2.Contract.Endpointv2Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Endpointv2 *Endpointv2Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Endpointv2.Contract.Endpointv2Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Endpointv2 *Endpointv2CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Endpointv2.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Endpointv2 *Endpointv2TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Endpointv2.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Endpointv2 *Endpointv2TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Endpointv2.Contract.contract.Transact(opts, method, params...)
}

// EMPTYPAYLOADHASH is a free data retrieval call binding the contract method 0xcb5026b9.
//
// Solidity: function EMPTY_PAYLOAD_HASH() view returns(bytes32)
func (_Endpointv2 *Endpointv2Caller) EMPTYPAYLOADHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Endpointv2.contract.Call(opts, &out, "EMPTY_PAYLOAD_HASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// EMPTYPAYLOADHASH is a free data retrieval call binding the contract method 0xcb5026b9.
//
// Solidity: function EMPTY_PAYLOAD_HASH() view returns(bytes32)
func (_Endpointv2 *Endpointv2Session) EMPTYPAYLOADHASH() ([32]byte, error) {
	return _Endpointv2.Contract.EMPTYPAYLOADHASH(&_Endpointv2.CallOpts)
}

// EMPTYPAYLOADHASH is a free data retrieval call binding the contract method 0xcb5026b9.
//
// Solidity: function EMPTY_PAYLOAD_HASH() view returns(bytes32)
func (_Endpointv2 *Endpointv2CallerSession) EMPTYPAYLOADHASH() ([32]byte, error) {
	return _Endpointv2.Contract.EMPTYPAYLOADHASH(&_Endpointv2.CallOpts)
}

// NILPAYLOADHASH is a free data retrieval call binding the contract method 0x2baf0be7.
//
// Solidity: function NIL_PAYLOAD_HASH() view returns(bytes32)
func (_Endpointv2 *Endpointv2Caller) NILPAYLOADHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Endpointv2.contract.Call(opts, &out, "NIL_PAYLOAD_HASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// NILPAYLOADHASH is a free data retrieval call binding the contract method 0x2baf0be7.
//
// Solidity: function NIL_PAYLOAD_HASH() view returns(bytes32)
func (_Endpointv2 *Endpointv2Session) NILPAYLOADHASH() ([32]byte, error) {
	return _Endpointv2.Contract.NILPAYLOADHASH(&_Endpointv2.CallOpts)
}

// NILPAYLOADHASH is a free data retrieval call binding the contract method 0x2baf0be7.
//
// Solidity: function NIL_PAYLOAD_HASH() view returns(bytes32)
func (_Endpointv2 *Endpointv2CallerSession) NILPAYLOADHASH() ([32]byte, error) {
	return _Endpointv2.Contract.NILPAYLOADHASH(&_Endpointv2.CallOpts)
}

// BlockedLibrary is a free data retrieval call binding the contract method 0x73318091.
//
// Solidity: function blockedLibrary() view returns(address)
func (_Endpointv2 *Endpointv2Caller) BlockedLibrary(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Endpointv2.contract.Call(opts, &out, "blockedLibrary")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BlockedLibrary is a free data retrieval call binding the contract method 0x73318091.
//
// Solidity: function blockedLibrary() view returns(address)
func (_Endpointv2 *Endpointv2Session) BlockedLibrary() (common.Address, error) {
	return _Endpointv2.Contract.BlockedLibrary(&_Endpointv2.CallOpts)
}

// BlockedLibrary is a free data retrieval call binding the contract method 0x73318091.
//
// Solidity: function blockedLibrary() view returns(address)
func (_Endpointv2 *Endpointv2CallerSession) BlockedLibrary() (common.Address, error) {
	return _Endpointv2.Contract.BlockedLibrary(&_Endpointv2.CallOpts)
}

// ComposeQueue is a free data retrieval call binding the contract method 0x35d330b0.
//
// Solidity: function composeQueue(address from, address to, bytes32 guid, uint16 index) view returns(bytes32 messageHash)
func (_Endpointv2 *Endpointv2Caller) ComposeQueue(opts *bind.CallOpts, from common.Address, to common.Address, guid [32]byte, index uint16) ([32]byte, error) {
	var out []interface{}
	err := _Endpointv2.contract.Call(opts, &out, "composeQueue", from, to, guid, index)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ComposeQueue is a free data retrieval call binding the contract method 0x35d330b0.
//
// Solidity: function composeQueue(address from, address to, bytes32 guid, uint16 index) view returns(bytes32 messageHash)
func (_Endpointv2 *Endpointv2Session) ComposeQueue(from common.Address, to common.Address, guid [32]byte, index uint16) ([32]byte, error) {
	return _Endpointv2.Contract.ComposeQueue(&_Endpointv2.CallOpts, from, to, guid, index)
}

// ComposeQueue is a free data retrieval call binding the contract method 0x35d330b0.
//
// Solidity: function composeQueue(address from, address to, bytes32 guid, uint16 index) view returns(bytes32 messageHash)
func (_Endpointv2 *Endpointv2CallerSession) ComposeQueue(from common.Address, to common.Address, guid [32]byte, index uint16) ([32]byte, error) {
	return _Endpointv2.Contract.ComposeQueue(&_Endpointv2.CallOpts, from, to, guid, index)
}

// DefaultReceiveLibrary is a free data retrieval call binding the contract method 0x6f50a803.
//
// Solidity: function defaultReceiveLibrary(uint32 srcEid) view returns(address lib)
func (_Endpointv2 *Endpointv2Caller) DefaultReceiveLibrary(opts *bind.CallOpts, srcEid uint32) (common.Address, error) {
	var out []interface{}
	err := _Endpointv2.contract.Call(opts, &out, "defaultReceiveLibrary", srcEid)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DefaultReceiveLibrary is a free data retrieval call binding the contract method 0x6f50a803.
//
// Solidity: function defaultReceiveLibrary(uint32 srcEid) view returns(address lib)
func (_Endpointv2 *Endpointv2Session) DefaultReceiveLibrary(srcEid uint32) (common.Address, error) {
	return _Endpointv2.Contract.DefaultReceiveLibrary(&_Endpointv2.CallOpts, srcEid)
}

// DefaultReceiveLibrary is a free data retrieval call binding the contract method 0x6f50a803.
//
// Solidity: function defaultReceiveLibrary(uint32 srcEid) view returns(address lib)
func (_Endpointv2 *Endpointv2CallerSession) DefaultReceiveLibrary(srcEid uint32) (common.Address, error) {
	return _Endpointv2.Contract.DefaultReceiveLibrary(&_Endpointv2.CallOpts, srcEid)
}

// DefaultReceiveLibraryTimeout is a free data retrieval call binding the contract method 0x6e83f5bb.
//
// Solidity: function defaultReceiveLibraryTimeout(uint32 srcEid) view returns(address lib, uint256 expiry)
func (_Endpointv2 *Endpointv2Caller) DefaultReceiveLibraryTimeout(opts *bind.CallOpts, srcEid uint32) (struct {
	Lib    common.Address
	Expiry *big.Int
}, error) {
	var out []interface{}
	err := _Endpointv2.contract.Call(opts, &out, "defaultReceiveLibraryTimeout", srcEid)

	outstruct := new(struct {
		Lib    common.Address
		Expiry *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Lib = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Expiry = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// DefaultReceiveLibraryTimeout is a free data retrieval call binding the contract method 0x6e83f5bb.
//
// Solidity: function defaultReceiveLibraryTimeout(uint32 srcEid) view returns(address lib, uint256 expiry)
func (_Endpointv2 *Endpointv2Session) DefaultReceiveLibraryTimeout(srcEid uint32) (struct {
	Lib    common.Address
	Expiry *big.Int
}, error) {
	return _Endpointv2.Contract.DefaultReceiveLibraryTimeout(&_Endpointv2.CallOpts, srcEid)
}

// DefaultReceiveLibraryTimeout is a free data retrieval call binding the contract method 0x6e83f5bb.
//
// Solidity: function defaultReceiveLibraryTimeout(uint32 srcEid) view returns(address lib, uint256 expiry)
func (_Endpointv2 *Endpointv2CallerSession) DefaultReceiveLibraryTimeout(srcEid uint32) (struct {
	Lib    common.Address
	Expiry *big.Int
}, error) {
	return _Endpointv2.Contract.DefaultReceiveLibraryTimeout(&_Endpointv2.CallOpts, srcEid)
}

// DefaultSendLibrary is a free data retrieval call binding the contract method 0xf64be4c7.
//
// Solidity: function defaultSendLibrary(uint32 dstEid) view returns(address lib)
func (_Endpointv2 *Endpointv2Caller) DefaultSendLibrary(opts *bind.CallOpts, dstEid uint32) (common.Address, error) {
	var out []interface{}
	err := _Endpointv2.contract.Call(opts, &out, "defaultSendLibrary", dstEid)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DefaultSendLibrary is a free data retrieval call binding the contract method 0xf64be4c7.
//
// Solidity: function defaultSendLibrary(uint32 dstEid) view returns(address lib)
func (_Endpointv2 *Endpointv2Session) DefaultSendLibrary(dstEid uint32) (common.Address, error) {
	return _Endpointv2.Contract.DefaultSendLibrary(&_Endpointv2.CallOpts, dstEid)
}

// DefaultSendLibrary is a free data retrieval call binding the contract method 0xf64be4c7.
//
// Solidity: function defaultSendLibrary(uint32 dstEid) view returns(address lib)
func (_Endpointv2 *Endpointv2CallerSession) DefaultSendLibrary(dstEid uint32) (common.Address, error) {
	return _Endpointv2.Contract.DefaultSendLibrary(&_Endpointv2.CallOpts, dstEid)
}

// Delegates is a free data retrieval call binding the contract method 0x587cde1e.
//
// Solidity: function delegates(address oapp) view returns(address delegate)
func (_Endpointv2 *Endpointv2Caller) Delegates(opts *bind.CallOpts, oapp common.Address) (common.Address, error) {
	var out []interface{}
	err := _Endpointv2.contract.Call(opts, &out, "delegates", oapp)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Delegates is a free data retrieval call binding the contract method 0x587cde1e.
//
// Solidity: function delegates(address oapp) view returns(address delegate)
func (_Endpointv2 *Endpointv2Session) Delegates(oapp common.Address) (common.Address, error) {
	return _Endpointv2.Contract.Delegates(&_Endpointv2.CallOpts, oapp)
}

// Delegates is a free data retrieval call binding the contract method 0x587cde1e.
//
// Solidity: function delegates(address oapp) view returns(address delegate)
func (_Endpointv2 *Endpointv2CallerSession) Delegates(oapp common.Address) (common.Address, error) {
	return _Endpointv2.Contract.Delegates(&_Endpointv2.CallOpts, oapp)
}

// Eid is a free data retrieval call binding the contract method 0x416ecebf.
//
// Solidity: function eid() view returns(uint32)
func (_Endpointv2 *Endpointv2Caller) Eid(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Endpointv2.contract.Call(opts, &out, "eid")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// Eid is a free data retrieval call binding the contract method 0x416ecebf.
//
// Solidity: function eid() view returns(uint32)
func (_Endpointv2 *Endpointv2Session) Eid() (uint32, error) {
	return _Endpointv2.Contract.Eid(&_Endpointv2.CallOpts)
}

// Eid is a free data retrieval call binding the contract method 0x416ecebf.
//
// Solidity: function eid() view returns(uint32)
func (_Endpointv2 *Endpointv2CallerSession) Eid() (uint32, error) {
	return _Endpointv2.Contract.Eid(&_Endpointv2.CallOpts)
}

// GetConfig is a free data retrieval call binding the contract method 0x2b3197b9.
//
// Solidity: function getConfig(address _oapp, address _lib, uint32 _eid, uint32 _configType) view returns(bytes config)
func (_Endpointv2 *Endpointv2Caller) GetConfig(opts *bind.CallOpts, _oapp common.Address, _lib common.Address, _eid uint32, _configType uint32) ([]byte, error) {
	var out []interface{}
	err := _Endpointv2.contract.Call(opts, &out, "getConfig", _oapp, _lib, _eid, _configType)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetConfig is a free data retrieval call binding the contract method 0x2b3197b9.
//
// Solidity: function getConfig(address _oapp, address _lib, uint32 _eid, uint32 _configType) view returns(bytes config)
func (_Endpointv2 *Endpointv2Session) GetConfig(_oapp common.Address, _lib common.Address, _eid uint32, _configType uint32) ([]byte, error) {
	return _Endpointv2.Contract.GetConfig(&_Endpointv2.CallOpts, _oapp, _lib, _eid, _configType)
}

// GetConfig is a free data retrieval call binding the contract method 0x2b3197b9.
//
// Solidity: function getConfig(address _oapp, address _lib, uint32 _eid, uint32 _configType) view returns(bytes config)
func (_Endpointv2 *Endpointv2CallerSession) GetConfig(_oapp common.Address, _lib common.Address, _eid uint32, _configType uint32) ([]byte, error) {
	return _Endpointv2.Contract.GetConfig(&_Endpointv2.CallOpts, _oapp, _lib, _eid, _configType)
}

// GetReceiveLibrary is a free data retrieval call binding the contract method 0x402f8468.
//
// Solidity: function getReceiveLibrary(address _receiver, uint32 _srcEid) view returns(address lib, bool isDefault)
func (_Endpointv2 *Endpointv2Caller) GetReceiveLibrary(opts *bind.CallOpts, _receiver common.Address, _srcEid uint32) (struct {
	Lib       common.Address
	IsDefault bool
}, error) {
	var out []interface{}
	err := _Endpointv2.contract.Call(opts, &out, "getReceiveLibrary", _receiver, _srcEid)

	outstruct := new(struct {
		Lib       common.Address
		IsDefault bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Lib = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.IsDefault = *abi.ConvertType(out[1], new(bool)).(*bool)

	return *outstruct, err

}

// GetReceiveLibrary is a free data retrieval call binding the contract method 0x402f8468.
//
// Solidity: function getReceiveLibrary(address _receiver, uint32 _srcEid) view returns(address lib, bool isDefault)
func (_Endpointv2 *Endpointv2Session) GetReceiveLibrary(_receiver common.Address, _srcEid uint32) (struct {
	Lib       common.Address
	IsDefault bool
}, error) {
	return _Endpointv2.Contract.GetReceiveLibrary(&_Endpointv2.CallOpts, _receiver, _srcEid)
}

// GetReceiveLibrary is a free data retrieval call binding the contract method 0x402f8468.
//
// Solidity: function getReceiveLibrary(address _receiver, uint32 _srcEid) view returns(address lib, bool isDefault)
func (_Endpointv2 *Endpointv2CallerSession) GetReceiveLibrary(_receiver common.Address, _srcEid uint32) (struct {
	Lib       common.Address
	IsDefault bool
}, error) {
	return _Endpointv2.Contract.GetReceiveLibrary(&_Endpointv2.CallOpts, _receiver, _srcEid)
}

// GetRegisteredLibraries is a free data retrieval call binding the contract method 0x9132e5c3.
//
// Solidity: function getRegisteredLibraries() view returns(address[])
func (_Endpointv2 *Endpointv2Caller) GetRegisteredLibraries(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Endpointv2.contract.Call(opts, &out, "getRegisteredLibraries")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetRegisteredLibraries is a free data retrieval call binding the contract method 0x9132e5c3.
//
// Solidity: function getRegisteredLibraries() view returns(address[])
func (_Endpointv2 *Endpointv2Session) GetRegisteredLibraries() ([]common.Address, error) {
	return _Endpointv2.Contract.GetRegisteredLibraries(&_Endpointv2.CallOpts)
}

// GetRegisteredLibraries is a free data retrieval call binding the contract method 0x9132e5c3.
//
// Solidity: function getRegisteredLibraries() view returns(address[])
func (_Endpointv2 *Endpointv2CallerSession) GetRegisteredLibraries() ([]common.Address, error) {
	return _Endpointv2.Contract.GetRegisteredLibraries(&_Endpointv2.CallOpts)
}

// GetSendContext is a free data retrieval call binding the contract method 0x14f651a9.
//
// Solidity: function getSendContext() view returns(uint32, address)
func (_Endpointv2 *Endpointv2Caller) GetSendContext(opts *bind.CallOpts) (uint32, common.Address, error) {
	var out []interface{}
	err := _Endpointv2.contract.Call(opts, &out, "getSendContext")

	if err != nil {
		return *new(uint32), *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)
	out1 := *abi.ConvertType(out[1], new(common.Address)).(*common.Address)

	return out0, out1, err

}

// GetSendContext is a free data retrieval call binding the contract method 0x14f651a9.
//
// Solidity: function getSendContext() view returns(uint32, address)
func (_Endpointv2 *Endpointv2Session) GetSendContext() (uint32, common.Address, error) {
	return _Endpointv2.Contract.GetSendContext(&_Endpointv2.CallOpts)
}

// GetSendContext is a free data retrieval call binding the contract method 0x14f651a9.
//
// Solidity: function getSendContext() view returns(uint32, address)
func (_Endpointv2 *Endpointv2CallerSession) GetSendContext() (uint32, common.Address, error) {
	return _Endpointv2.Contract.GetSendContext(&_Endpointv2.CallOpts)
}

// GetSendLibrary is a free data retrieval call binding the contract method 0xb96a277f.
//
// Solidity: function getSendLibrary(address _sender, uint32 _dstEid) view returns(address lib)
func (_Endpointv2 *Endpointv2Caller) GetSendLibrary(opts *bind.CallOpts, _sender common.Address, _dstEid uint32) (common.Address, error) {
	var out []interface{}
	err := _Endpointv2.contract.Call(opts, &out, "getSendLibrary", _sender, _dstEid)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetSendLibrary is a free data retrieval call binding the contract method 0xb96a277f.
//
// Solidity: function getSendLibrary(address _sender, uint32 _dstEid) view returns(address lib)
func (_Endpointv2 *Endpointv2Session) GetSendLibrary(_sender common.Address, _dstEid uint32) (common.Address, error) {
	return _Endpointv2.Contract.GetSendLibrary(&_Endpointv2.CallOpts, _sender, _dstEid)
}

// GetSendLibrary is a free data retrieval call binding the contract method 0xb96a277f.
//
// Solidity: function getSendLibrary(address _sender, uint32 _dstEid) view returns(address lib)
func (_Endpointv2 *Endpointv2CallerSession) GetSendLibrary(_sender common.Address, _dstEid uint32) (common.Address, error) {
	return _Endpointv2.Contract.GetSendLibrary(&_Endpointv2.CallOpts, _sender, _dstEid)
}

// InboundNonce is a free data retrieval call binding the contract method 0xa0dd43fc.
//
// Solidity: function inboundNonce(address _receiver, uint32 _srcEid, bytes32 _sender) view returns(uint64)
func (_Endpointv2 *Endpointv2Caller) InboundNonce(opts *bind.CallOpts, _receiver common.Address, _srcEid uint32, _sender [32]byte) (uint64, error) {
	var out []interface{}
	err := _Endpointv2.contract.Call(opts, &out, "inboundNonce", _receiver, _srcEid, _sender)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// InboundNonce is a free data retrieval call binding the contract method 0xa0dd43fc.
//
// Solidity: function inboundNonce(address _receiver, uint32 _srcEid, bytes32 _sender) view returns(uint64)
func (_Endpointv2 *Endpointv2Session) InboundNonce(_receiver common.Address, _srcEid uint32, _sender [32]byte) (uint64, error) {
	return _Endpointv2.Contract.InboundNonce(&_Endpointv2.CallOpts, _receiver, _srcEid, _sender)
}

// InboundNonce is a free data retrieval call binding the contract method 0xa0dd43fc.
//
// Solidity: function inboundNonce(address _receiver, uint32 _srcEid, bytes32 _sender) view returns(uint64)
func (_Endpointv2 *Endpointv2CallerSession) InboundNonce(_receiver common.Address, _srcEid uint32, _sender [32]byte) (uint64, error) {
	return _Endpointv2.Contract.InboundNonce(&_Endpointv2.CallOpts, _receiver, _srcEid, _sender)
}

// InboundPayloadHash is a free data retrieval call binding the contract method 0xc9fc7bcd.
//
// Solidity: function inboundPayloadHash(address receiver, uint32 srcEid, bytes32 sender, uint64 inboundNonce) view returns(bytes32 payloadHash)
func (_Endpointv2 *Endpointv2Caller) InboundPayloadHash(opts *bind.CallOpts, receiver common.Address, srcEid uint32, sender [32]byte, inboundNonce uint64) ([32]byte, error) {
	var out []interface{}
	err := _Endpointv2.contract.Call(opts, &out, "inboundPayloadHash", receiver, srcEid, sender, inboundNonce)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// InboundPayloadHash is a free data retrieval call binding the contract method 0xc9fc7bcd.
//
// Solidity: function inboundPayloadHash(address receiver, uint32 srcEid, bytes32 sender, uint64 inboundNonce) view returns(bytes32 payloadHash)
func (_Endpointv2 *Endpointv2Session) InboundPayloadHash(receiver common.Address, srcEid uint32, sender [32]byte, inboundNonce uint64) ([32]byte, error) {
	return _Endpointv2.Contract.InboundPayloadHash(&_Endpointv2.CallOpts, receiver, srcEid, sender, inboundNonce)
}

// InboundPayloadHash is a free data retrieval call binding the contract method 0xc9fc7bcd.
//
// Solidity: function inboundPayloadHash(address receiver, uint32 srcEid, bytes32 sender, uint64 inboundNonce) view returns(bytes32 payloadHash)
func (_Endpointv2 *Endpointv2CallerSession) InboundPayloadHash(receiver common.Address, srcEid uint32, sender [32]byte, inboundNonce uint64) ([32]byte, error) {
	return _Endpointv2.Contract.InboundPayloadHash(&_Endpointv2.CallOpts, receiver, srcEid, sender, inboundNonce)
}

// Initializable is a free data retrieval call binding the contract method 0x861e1ca5.
//
// Solidity: function initializable((uint32,bytes32,uint64) _origin, address _receiver) view returns(bool)
func (_Endpointv2 *Endpointv2Caller) Initializable(opts *bind.CallOpts, _origin Origin, _receiver common.Address) (bool, error) {
	var out []interface{}
	err := _Endpointv2.contract.Call(opts, &out, "initializable", _origin, _receiver)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Initializable is a free data retrieval call binding the contract method 0x861e1ca5.
//
// Solidity: function initializable((uint32,bytes32,uint64) _origin, address _receiver) view returns(bool)
func (_Endpointv2 *Endpointv2Session) Initializable(_origin Origin, _receiver common.Address) (bool, error) {
	return _Endpointv2.Contract.Initializable(&_Endpointv2.CallOpts, _origin, _receiver)
}

// Initializable is a free data retrieval call binding the contract method 0x861e1ca5.
//
// Solidity: function initializable((uint32,bytes32,uint64) _origin, address _receiver) view returns(bool)
func (_Endpointv2 *Endpointv2CallerSession) Initializable(_origin Origin, _receiver common.Address) (bool, error) {
	return _Endpointv2.Contract.Initializable(&_Endpointv2.CallOpts, _origin, _receiver)
}

// IsDefaultSendLibrary is a free data retrieval call binding the contract method 0xdc93c8a2.
//
// Solidity: function isDefaultSendLibrary(address _sender, uint32 _dstEid) view returns(bool)
func (_Endpointv2 *Endpointv2Caller) IsDefaultSendLibrary(opts *bind.CallOpts, _sender common.Address, _dstEid uint32) (bool, error) {
	var out []interface{}
	err := _Endpointv2.contract.Call(opts, &out, "isDefaultSendLibrary", _sender, _dstEid)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsDefaultSendLibrary is a free data retrieval call binding the contract method 0xdc93c8a2.
//
// Solidity: function isDefaultSendLibrary(address _sender, uint32 _dstEid) view returns(bool)
func (_Endpointv2 *Endpointv2Session) IsDefaultSendLibrary(_sender common.Address, _dstEid uint32) (bool, error) {
	return _Endpointv2.Contract.IsDefaultSendLibrary(&_Endpointv2.CallOpts, _sender, _dstEid)
}

// IsDefaultSendLibrary is a free data retrieval call binding the contract method 0xdc93c8a2.
//
// Solidity: function isDefaultSendLibrary(address _sender, uint32 _dstEid) view returns(bool)
func (_Endpointv2 *Endpointv2CallerSession) IsDefaultSendLibrary(_sender common.Address, _dstEid uint32) (bool, error) {
	return _Endpointv2.Contract.IsDefaultSendLibrary(&_Endpointv2.CallOpts, _sender, _dstEid)
}

// IsRegisteredLibrary is a free data retrieval call binding the contract method 0xdc706a62.
//
// Solidity: function isRegisteredLibrary(address lib) view returns(bool)
func (_Endpointv2 *Endpointv2Caller) IsRegisteredLibrary(opts *bind.CallOpts, lib common.Address) (bool, error) {
	var out []interface{}
	err := _Endpointv2.contract.Call(opts, &out, "isRegisteredLibrary", lib)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsRegisteredLibrary is a free data retrieval call binding the contract method 0xdc706a62.
//
// Solidity: function isRegisteredLibrary(address lib) view returns(bool)
func (_Endpointv2 *Endpointv2Session) IsRegisteredLibrary(lib common.Address) (bool, error) {
	return _Endpointv2.Contract.IsRegisteredLibrary(&_Endpointv2.CallOpts, lib)
}

// IsRegisteredLibrary is a free data retrieval call binding the contract method 0xdc706a62.
//
// Solidity: function isRegisteredLibrary(address lib) view returns(bool)
func (_Endpointv2 *Endpointv2CallerSession) IsRegisteredLibrary(lib common.Address) (bool, error) {
	return _Endpointv2.Contract.IsRegisteredLibrary(&_Endpointv2.CallOpts, lib)
}

// IsSendingMessage is a free data retrieval call binding the contract method 0x79624ca9.
//
// Solidity: function isSendingMessage() view returns(bool)
func (_Endpointv2 *Endpointv2Caller) IsSendingMessage(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Endpointv2.contract.Call(opts, &out, "isSendingMessage")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsSendingMessage is a free data retrieval call binding the contract method 0x79624ca9.
//
// Solidity: function isSendingMessage() view returns(bool)
func (_Endpointv2 *Endpointv2Session) IsSendingMessage() (bool, error) {
	return _Endpointv2.Contract.IsSendingMessage(&_Endpointv2.CallOpts)
}

// IsSendingMessage is a free data retrieval call binding the contract method 0x79624ca9.
//
// Solidity: function isSendingMessage() view returns(bool)
func (_Endpointv2 *Endpointv2CallerSession) IsSendingMessage() (bool, error) {
	return _Endpointv2.Contract.IsSendingMessage(&_Endpointv2.CallOpts)
}

// IsSupportedEid is a free data retrieval call binding the contract method 0x6750cd4c.
//
// Solidity: function isSupportedEid(uint32 _eid) view returns(bool)
func (_Endpointv2 *Endpointv2Caller) IsSupportedEid(opts *bind.CallOpts, _eid uint32) (bool, error) {
	var out []interface{}
	err := _Endpointv2.contract.Call(opts, &out, "isSupportedEid", _eid)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsSupportedEid is a free data retrieval call binding the contract method 0x6750cd4c.
//
// Solidity: function isSupportedEid(uint32 _eid) view returns(bool)
func (_Endpointv2 *Endpointv2Session) IsSupportedEid(_eid uint32) (bool, error) {
	return _Endpointv2.Contract.IsSupportedEid(&_Endpointv2.CallOpts, _eid)
}

// IsSupportedEid is a free data retrieval call binding the contract method 0x6750cd4c.
//
// Solidity: function isSupportedEid(uint32 _eid) view returns(bool)
func (_Endpointv2 *Endpointv2CallerSession) IsSupportedEid(_eid uint32) (bool, error) {
	return _Endpointv2.Contract.IsSupportedEid(&_Endpointv2.CallOpts, _eid)
}

// IsValidReceiveLibrary is a free data retrieval call binding the contract method 0x9d7f9775.
//
// Solidity: function isValidReceiveLibrary(address _receiver, uint32 _srcEid, address _actualReceiveLib) view returns(bool)
func (_Endpointv2 *Endpointv2Caller) IsValidReceiveLibrary(opts *bind.CallOpts, _receiver common.Address, _srcEid uint32, _actualReceiveLib common.Address) (bool, error) {
	var out []interface{}
	err := _Endpointv2.contract.Call(opts, &out, "isValidReceiveLibrary", _receiver, _srcEid, _actualReceiveLib)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsValidReceiveLibrary is a free data retrieval call binding the contract method 0x9d7f9775.
//
// Solidity: function isValidReceiveLibrary(address _receiver, uint32 _srcEid, address _actualReceiveLib) view returns(bool)
func (_Endpointv2 *Endpointv2Session) IsValidReceiveLibrary(_receiver common.Address, _srcEid uint32, _actualReceiveLib common.Address) (bool, error) {
	return _Endpointv2.Contract.IsValidReceiveLibrary(&_Endpointv2.CallOpts, _receiver, _srcEid, _actualReceiveLib)
}

// IsValidReceiveLibrary is a free data retrieval call binding the contract method 0x9d7f9775.
//
// Solidity: function isValidReceiveLibrary(address _receiver, uint32 _srcEid, address _actualReceiveLib) view returns(bool)
func (_Endpointv2 *Endpointv2CallerSession) IsValidReceiveLibrary(_receiver common.Address, _srcEid uint32, _actualReceiveLib common.Address) (bool, error) {
	return _Endpointv2.Contract.IsValidReceiveLibrary(&_Endpointv2.CallOpts, _receiver, _srcEid, _actualReceiveLib)
}

// LazyInboundNonce is a free data retrieval call binding the contract method 0x5b17bb70.
//
// Solidity: function lazyInboundNonce(address receiver, uint32 srcEid, bytes32 sender) view returns(uint64 nonce)
func (_Endpointv2 *Endpointv2Caller) LazyInboundNonce(opts *bind.CallOpts, receiver common.Address, srcEid uint32, sender [32]byte) (uint64, error) {
	var out []interface{}
	err := _Endpointv2.contract.Call(opts, &out, "lazyInboundNonce", receiver, srcEid, sender)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// LazyInboundNonce is a free data retrieval call binding the contract method 0x5b17bb70.
//
// Solidity: function lazyInboundNonce(address receiver, uint32 srcEid, bytes32 sender) view returns(uint64 nonce)
func (_Endpointv2 *Endpointv2Session) LazyInboundNonce(receiver common.Address, srcEid uint32, sender [32]byte) (uint64, error) {
	return _Endpointv2.Contract.LazyInboundNonce(&_Endpointv2.CallOpts, receiver, srcEid, sender)
}

// LazyInboundNonce is a free data retrieval call binding the contract method 0x5b17bb70.
//
// Solidity: function lazyInboundNonce(address receiver, uint32 srcEid, bytes32 sender) view returns(uint64 nonce)
func (_Endpointv2 *Endpointv2CallerSession) LazyInboundNonce(receiver common.Address, srcEid uint32, sender [32]byte) (uint64, error) {
	return _Endpointv2.Contract.LazyInboundNonce(&_Endpointv2.CallOpts, receiver, srcEid, sender)
}

// LzToken is a free data retrieval call binding the contract method 0xe4fe1d94.
//
// Solidity: function lzToken() view returns(address)
func (_Endpointv2 *Endpointv2Caller) LzToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Endpointv2.contract.Call(opts, &out, "lzToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LzToken is a free data retrieval call binding the contract method 0xe4fe1d94.
//
// Solidity: function lzToken() view returns(address)
func (_Endpointv2 *Endpointv2Session) LzToken() (common.Address, error) {
	return _Endpointv2.Contract.LzToken(&_Endpointv2.CallOpts)
}

// LzToken is a free data retrieval call binding the contract method 0xe4fe1d94.
//
// Solidity: function lzToken() view returns(address)
func (_Endpointv2 *Endpointv2CallerSession) LzToken() (common.Address, error) {
	return _Endpointv2.Contract.LzToken(&_Endpointv2.CallOpts)
}

// NativeToken is a free data retrieval call binding the contract method 0xe1758bd8.
//
// Solidity: function nativeToken() view returns(address)
func (_Endpointv2 *Endpointv2Caller) NativeToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Endpointv2.contract.Call(opts, &out, "nativeToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NativeToken is a free data retrieval call binding the contract method 0xe1758bd8.
//
// Solidity: function nativeToken() view returns(address)
func (_Endpointv2 *Endpointv2Session) NativeToken() (common.Address, error) {
	return _Endpointv2.Contract.NativeToken(&_Endpointv2.CallOpts)
}

// NativeToken is a free data retrieval call binding the contract method 0xe1758bd8.
//
// Solidity: function nativeToken() view returns(address)
func (_Endpointv2 *Endpointv2CallerSession) NativeToken() (common.Address, error) {
	return _Endpointv2.Contract.NativeToken(&_Endpointv2.CallOpts)
}

// NextGuid is a free data retrieval call binding the contract method 0xaafe5e07.
//
// Solidity: function nextGuid(address _sender, uint32 _dstEid, bytes32 _receiver) view returns(bytes32)
func (_Endpointv2 *Endpointv2Caller) NextGuid(opts *bind.CallOpts, _sender common.Address, _dstEid uint32, _receiver [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Endpointv2.contract.Call(opts, &out, "nextGuid", _sender, _dstEid, _receiver)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// NextGuid is a free data retrieval call binding the contract method 0xaafe5e07.
//
// Solidity: function nextGuid(address _sender, uint32 _dstEid, bytes32 _receiver) view returns(bytes32)
func (_Endpointv2 *Endpointv2Session) NextGuid(_sender common.Address, _dstEid uint32, _receiver [32]byte) ([32]byte, error) {
	return _Endpointv2.Contract.NextGuid(&_Endpointv2.CallOpts, _sender, _dstEid, _receiver)
}

// NextGuid is a free data retrieval call binding the contract method 0xaafe5e07.
//
// Solidity: function nextGuid(address _sender, uint32 _dstEid, bytes32 _receiver) view returns(bytes32)
func (_Endpointv2 *Endpointv2CallerSession) NextGuid(_sender common.Address, _dstEid uint32, _receiver [32]byte) ([32]byte, error) {
	return _Endpointv2.Contract.NextGuid(&_Endpointv2.CallOpts, _sender, _dstEid, _receiver)
}

// OutboundNonce is a free data retrieval call binding the contract method 0x9c6d7340.
//
// Solidity: function outboundNonce(address sender, uint32 dstEid, bytes32 receiver) view returns(uint64 nonce)
func (_Endpointv2 *Endpointv2Caller) OutboundNonce(opts *bind.CallOpts, sender common.Address, dstEid uint32, receiver [32]byte) (uint64, error) {
	var out []interface{}
	err := _Endpointv2.contract.Call(opts, &out, "outboundNonce", sender, dstEid, receiver)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// OutboundNonce is a free data retrieval call binding the contract method 0x9c6d7340.
//
// Solidity: function outboundNonce(address sender, uint32 dstEid, bytes32 receiver) view returns(uint64 nonce)
func (_Endpointv2 *Endpointv2Session) OutboundNonce(sender common.Address, dstEid uint32, receiver [32]byte) (uint64, error) {
	return _Endpointv2.Contract.OutboundNonce(&_Endpointv2.CallOpts, sender, dstEid, receiver)
}

// OutboundNonce is a free data retrieval call binding the contract method 0x9c6d7340.
//
// Solidity: function outboundNonce(address sender, uint32 dstEid, bytes32 receiver) view returns(uint64 nonce)
func (_Endpointv2 *Endpointv2CallerSession) OutboundNonce(sender common.Address, dstEid uint32, receiver [32]byte) (uint64, error) {
	return _Endpointv2.Contract.OutboundNonce(&_Endpointv2.CallOpts, sender, dstEid, receiver)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Endpointv2 *Endpointv2Caller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Endpointv2.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Endpointv2 *Endpointv2Session) Owner() (common.Address, error) {
	return _Endpointv2.Contract.Owner(&_Endpointv2.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Endpointv2 *Endpointv2CallerSession) Owner() (common.Address, error) {
	return _Endpointv2.Contract.Owner(&_Endpointv2.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Endpointv2 *Endpointv2Caller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Endpointv2.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Endpointv2 *Endpointv2Session) Paused() (bool, error) {
	return _Endpointv2.Contract.Paused(&_Endpointv2.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Endpointv2 *Endpointv2CallerSession) Paused() (bool, error) {
	return _Endpointv2.Contract.Paused(&_Endpointv2.CallOpts)
}

// Quote is a free data retrieval call binding the contract method 0xddc28c58.
//
// Solidity: function quote((uint32,bytes32,bytes,bytes,bool) _params, address _sender) view returns((uint256,uint256))
func (_Endpointv2 *Endpointv2Caller) Quote(opts *bind.CallOpts, _params MessagingParams, _sender common.Address) (MessagingFee, error) {
	var out []interface{}
	err := _Endpointv2.contract.Call(opts, &out, "quote", _params, _sender)

	if err != nil {
		return *new(MessagingFee), err
	}

	out0 := *abi.ConvertType(out[0], new(MessagingFee)).(*MessagingFee)

	return out0, err

}

// Quote is a free data retrieval call binding the contract method 0xddc28c58.
//
// Solidity: function quote((uint32,bytes32,bytes,bytes,bool) _params, address _sender) view returns((uint256,uint256))
func (_Endpointv2 *Endpointv2Session) Quote(_params MessagingParams, _sender common.Address) (MessagingFee, error) {
	return _Endpointv2.Contract.Quote(&_Endpointv2.CallOpts, _params, _sender)
}

// Quote is a free data retrieval call binding the contract method 0xddc28c58.
//
// Solidity: function quote((uint32,bytes32,bytes,bytes,bool) _params, address _sender) view returns((uint256,uint256))
func (_Endpointv2 *Endpointv2CallerSession) Quote(_params MessagingParams, _sender common.Address) (MessagingFee, error) {
	return _Endpointv2.Contract.Quote(&_Endpointv2.CallOpts, _params, _sender)
}

// ReceiveLibraryTimeout is a free data retrieval call binding the contract method 0xef667aa1.
//
// Solidity: function receiveLibraryTimeout(address receiver, uint32 srcEid) view returns(address lib, uint256 expiry)
func (_Endpointv2 *Endpointv2Caller) ReceiveLibraryTimeout(opts *bind.CallOpts, receiver common.Address, srcEid uint32) (struct {
	Lib    common.Address
	Expiry *big.Int
}, error) {
	var out []interface{}
	err := _Endpointv2.contract.Call(opts, &out, "receiveLibraryTimeout", receiver, srcEid)

	outstruct := new(struct {
		Lib    common.Address
		Expiry *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Lib = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Expiry = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// ReceiveLibraryTimeout is a free data retrieval call binding the contract method 0xef667aa1.
//
// Solidity: function receiveLibraryTimeout(address receiver, uint32 srcEid) view returns(address lib, uint256 expiry)
func (_Endpointv2 *Endpointv2Session) ReceiveLibraryTimeout(receiver common.Address, srcEid uint32) (struct {
	Lib    common.Address
	Expiry *big.Int
}, error) {
	return _Endpointv2.Contract.ReceiveLibraryTimeout(&_Endpointv2.CallOpts, receiver, srcEid)
}

// ReceiveLibraryTimeout is a free data retrieval call binding the contract method 0xef667aa1.
//
// Solidity: function receiveLibraryTimeout(address receiver, uint32 srcEid) view returns(address lib, uint256 expiry)
func (_Endpointv2 *Endpointv2CallerSession) ReceiveLibraryTimeout(receiver common.Address, srcEid uint32) (struct {
	Lib    common.Address
	Expiry *big.Int
}, error) {
	return _Endpointv2.Contract.ReceiveLibraryTimeout(&_Endpointv2.CallOpts, receiver, srcEid)
}

// Verifiable is a free data retrieval call binding the contract method 0xc9a54a99.
//
// Solidity: function verifiable((uint32,bytes32,uint64) _origin, address _receiver) view returns(bool)
func (_Endpointv2 *Endpointv2Caller) Verifiable(opts *bind.CallOpts, _origin Origin, _receiver common.Address) (bool, error) {
	var out []interface{}
	err := _Endpointv2.contract.Call(opts, &out, "verifiable", _origin, _receiver)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Verifiable is a free data retrieval call binding the contract method 0xc9a54a99.
//
// Solidity: function verifiable((uint32,bytes32,uint64) _origin, address _receiver) view returns(bool)
func (_Endpointv2 *Endpointv2Session) Verifiable(_origin Origin, _receiver common.Address) (bool, error) {
	return _Endpointv2.Contract.Verifiable(&_Endpointv2.CallOpts, _origin, _receiver)
}

// Verifiable is a free data retrieval call binding the contract method 0xc9a54a99.
//
// Solidity: function verifiable((uint32,bytes32,uint64) _origin, address _receiver) view returns(bool)
func (_Endpointv2 *Endpointv2CallerSession) Verifiable(_origin Origin, _receiver common.Address) (bool, error) {
	return _Endpointv2.Contract.Verifiable(&_Endpointv2.CallOpts, _origin, _receiver)
}

// Burn is a paid mutator transaction binding the contract method 0x40f80683.
//
// Solidity: function burn(address _oapp, uint32 _srcEid, bytes32 _sender, uint64 _nonce, bytes32 _payloadHash) returns()
func (_Endpointv2 *Endpointv2Transactor) Burn(opts *bind.TransactOpts, _oapp common.Address, _srcEid uint32, _sender [32]byte, _nonce uint64, _payloadHash [32]byte) (*types.Transaction, error) {
	return _Endpointv2.contract.Transact(opts, "burn", _oapp, _srcEid, _sender, _nonce, _payloadHash)
}

// Burn is a paid mutator transaction binding the contract method 0x40f80683.
//
// Solidity: function burn(address _oapp, uint32 _srcEid, bytes32 _sender, uint64 _nonce, bytes32 _payloadHash) returns()
func (_Endpointv2 *Endpointv2Session) Burn(_oapp common.Address, _srcEid uint32, _sender [32]byte, _nonce uint64, _payloadHash [32]byte) (*types.Transaction, error) {
	return _Endpointv2.Contract.Burn(&_Endpointv2.TransactOpts, _oapp, _srcEid, _sender, _nonce, _payloadHash)
}

// Burn is a paid mutator transaction binding the contract method 0x40f80683.
//
// Solidity: function burn(address _oapp, uint32 _srcEid, bytes32 _sender, uint64 _nonce, bytes32 _payloadHash) returns()
func (_Endpointv2 *Endpointv2TransactorSession) Burn(_oapp common.Address, _srcEid uint32, _sender [32]byte, _nonce uint64, _payloadHash [32]byte) (*types.Transaction, error) {
	return _Endpointv2.Contract.Burn(&_Endpointv2.TransactOpts, _oapp, _srcEid, _sender, _nonce, _payloadHash)
}

// Clear is a paid mutator transaction binding the contract method 0x2a56c1b0.
//
// Solidity: function clear(address _oapp, (uint32,bytes32,uint64) _origin, bytes32 _guid, bytes _message) returns()
func (_Endpointv2 *Endpointv2Transactor) Clear(opts *bind.TransactOpts, _oapp common.Address, _origin Origin, _guid [32]byte, _message []byte) (*types.Transaction, error) {
	return _Endpointv2.contract.Transact(opts, "clear", _oapp, _origin, _guid, _message)
}

// Clear is a paid mutator transaction binding the contract method 0x2a56c1b0.
//
// Solidity: function clear(address _oapp, (uint32,bytes32,uint64) _origin, bytes32 _guid, bytes _message) returns()
func (_Endpointv2 *Endpointv2Session) Clear(_oapp common.Address, _origin Origin, _guid [32]byte, _message []byte) (*types.Transaction, error) {
	return _Endpointv2.Contract.Clear(&_Endpointv2.TransactOpts, _oapp, _origin, _guid, _message)
}

// Clear is a paid mutator transaction binding the contract method 0x2a56c1b0.
//
// Solidity: function clear(address _oapp, (uint32,bytes32,uint64) _origin, bytes32 _guid, bytes _message) returns()
func (_Endpointv2 *Endpointv2TransactorSession) Clear(_oapp common.Address, _origin Origin, _guid [32]byte, _message []byte) (*types.Transaction, error) {
	return _Endpointv2.Contract.Clear(&_Endpointv2.TransactOpts, _oapp, _origin, _guid, _message)
}

// LzCompose is a paid mutator transaction binding the contract method 0x91d20fa1.
//
// Solidity: function lzCompose(address _from, address _to, bytes32 _guid, uint16 _index, bytes _message, bytes _extraData) payable returns()
func (_Endpointv2 *Endpointv2Transactor) LzCompose(opts *bind.TransactOpts, _from common.Address, _to common.Address, _guid [32]byte, _index uint16, _message []byte, _extraData []byte) (*types.Transaction, error) {
	return _Endpointv2.contract.Transact(opts, "lzCompose", _from, _to, _guid, _index, _message, _extraData)
}

// LzCompose is a paid mutator transaction binding the contract method 0x91d20fa1.
//
// Solidity: function lzCompose(address _from, address _to, bytes32 _guid, uint16 _index, bytes _message, bytes _extraData) payable returns()
func (_Endpointv2 *Endpointv2Session) LzCompose(_from common.Address, _to common.Address, _guid [32]byte, _index uint16, _message []byte, _extraData []byte) (*types.Transaction, error) {
	return _Endpointv2.Contract.LzCompose(&_Endpointv2.TransactOpts, _from, _to, _guid, _index, _message, _extraData)
}

// LzCompose is a paid mutator transaction binding the contract method 0x91d20fa1.
//
// Solidity: function lzCompose(address _from, address _to, bytes32 _guid, uint16 _index, bytes _message, bytes _extraData) payable returns()
func (_Endpointv2 *Endpointv2TransactorSession) LzCompose(_from common.Address, _to common.Address, _guid [32]byte, _index uint16, _message []byte, _extraData []byte) (*types.Transaction, error) {
	return _Endpointv2.Contract.LzCompose(&_Endpointv2.TransactOpts, _from, _to, _guid, _index, _message, _extraData)
}

// LzComposeAlert is a paid mutator transaction binding the contract method 0x697fe6b6.
//
// Solidity: function lzComposeAlert(address _from, address _to, bytes32 _guid, uint16 _index, uint256 _gas, uint256 _value, bytes _message, bytes _extraData, bytes _reason) returns()
func (_Endpointv2 *Endpointv2Transactor) LzComposeAlert(opts *bind.TransactOpts, _from common.Address, _to common.Address, _guid [32]byte, _index uint16, _gas *big.Int, _value *big.Int, _message []byte, _extraData []byte, _reason []byte) (*types.Transaction, error) {
	return _Endpointv2.contract.Transact(opts, "lzComposeAlert", _from, _to, _guid, _index, _gas, _value, _message, _extraData, _reason)
}

// LzComposeAlert is a paid mutator transaction binding the contract method 0x697fe6b6.
//
// Solidity: function lzComposeAlert(address _from, address _to, bytes32 _guid, uint16 _index, uint256 _gas, uint256 _value, bytes _message, bytes _extraData, bytes _reason) returns()
func (_Endpointv2 *Endpointv2Session) LzComposeAlert(_from common.Address, _to common.Address, _guid [32]byte, _index uint16, _gas *big.Int, _value *big.Int, _message []byte, _extraData []byte, _reason []byte) (*types.Transaction, error) {
	return _Endpointv2.Contract.LzComposeAlert(&_Endpointv2.TransactOpts, _from, _to, _guid, _index, _gas, _value, _message, _extraData, _reason)
}

// LzComposeAlert is a paid mutator transaction binding the contract method 0x697fe6b6.
//
// Solidity: function lzComposeAlert(address _from, address _to, bytes32 _guid, uint16 _index, uint256 _gas, uint256 _value, bytes _message, bytes _extraData, bytes _reason) returns()
func (_Endpointv2 *Endpointv2TransactorSession) LzComposeAlert(_from common.Address, _to common.Address, _guid [32]byte, _index uint16, _gas *big.Int, _value *big.Int, _message []byte, _extraData []byte, _reason []byte) (*types.Transaction, error) {
	return _Endpointv2.Contract.LzComposeAlert(&_Endpointv2.TransactOpts, _from, _to, _guid, _index, _gas, _value, _message, _extraData, _reason)
}

// LzReceive is a paid mutator transaction binding the contract method 0x0c0c389e.
//
// Solidity: function lzReceive((uint32,bytes32,uint64) _origin, address _receiver, bytes32 _guid, bytes _message, bytes _extraData) payable returns()
func (_Endpointv2 *Endpointv2Transactor) LzReceive(opts *bind.TransactOpts, _origin Origin, _receiver common.Address, _guid [32]byte, _message []byte, _extraData []byte) (*types.Transaction, error) {
	return _Endpointv2.contract.Transact(opts, "lzReceive", _origin, _receiver, _guid, _message, _extraData)
}

// LzReceive is a paid mutator transaction binding the contract method 0x0c0c389e.
//
// Solidity: function lzReceive((uint32,bytes32,uint64) _origin, address _receiver, bytes32 _guid, bytes _message, bytes _extraData) payable returns()
func (_Endpointv2 *Endpointv2Session) LzReceive(_origin Origin, _receiver common.Address, _guid [32]byte, _message []byte, _extraData []byte) (*types.Transaction, error) {
	return _Endpointv2.Contract.LzReceive(&_Endpointv2.TransactOpts, _origin, _receiver, _guid, _message, _extraData)
}

// LzReceive is a paid mutator transaction binding the contract method 0x0c0c389e.
//
// Solidity: function lzReceive((uint32,bytes32,uint64) _origin, address _receiver, bytes32 _guid, bytes _message, bytes _extraData) payable returns()
func (_Endpointv2 *Endpointv2TransactorSession) LzReceive(_origin Origin, _receiver common.Address, _guid [32]byte, _message []byte, _extraData []byte) (*types.Transaction, error) {
	return _Endpointv2.Contract.LzReceive(&_Endpointv2.TransactOpts, _origin, _receiver, _guid, _message, _extraData)
}

// LzReceiveAlert is a paid mutator transaction binding the contract method 0x6bf73fa3.
//
// Solidity: function lzReceiveAlert((uint32,bytes32,uint64) _origin, address _receiver, bytes32 _guid, uint256 _gas, uint256 _value, bytes _message, bytes _extraData, bytes _reason) returns()
func (_Endpointv2 *Endpointv2Transactor) LzReceiveAlert(opts *bind.TransactOpts, _origin Origin, _receiver common.Address, _guid [32]byte, _gas *big.Int, _value *big.Int, _message []byte, _extraData []byte, _reason []byte) (*types.Transaction, error) {
	return _Endpointv2.contract.Transact(opts, "lzReceiveAlert", _origin, _receiver, _guid, _gas, _value, _message, _extraData, _reason)
}

// LzReceiveAlert is a paid mutator transaction binding the contract method 0x6bf73fa3.
//
// Solidity: function lzReceiveAlert((uint32,bytes32,uint64) _origin, address _receiver, bytes32 _guid, uint256 _gas, uint256 _value, bytes _message, bytes _extraData, bytes _reason) returns()
func (_Endpointv2 *Endpointv2Session) LzReceiveAlert(_origin Origin, _receiver common.Address, _guid [32]byte, _gas *big.Int, _value *big.Int, _message []byte, _extraData []byte, _reason []byte) (*types.Transaction, error) {
	return _Endpointv2.Contract.LzReceiveAlert(&_Endpointv2.TransactOpts, _origin, _receiver, _guid, _gas, _value, _message, _extraData, _reason)
}

// LzReceiveAlert is a paid mutator transaction binding the contract method 0x6bf73fa3.
//
// Solidity: function lzReceiveAlert((uint32,bytes32,uint64) _origin, address _receiver, bytes32 _guid, uint256 _gas, uint256 _value, bytes _message, bytes _extraData, bytes _reason) returns()
func (_Endpointv2 *Endpointv2TransactorSession) LzReceiveAlert(_origin Origin, _receiver common.Address, _guid [32]byte, _gas *big.Int, _value *big.Int, _message []byte, _extraData []byte, _reason []byte) (*types.Transaction, error) {
	return _Endpointv2.Contract.LzReceiveAlert(&_Endpointv2.TransactOpts, _origin, _receiver, _guid, _gas, _value, _message, _extraData, _reason)
}

// Nilify is a paid mutator transaction binding the contract method 0x2e80fbf3.
//
// Solidity: function nilify(address _oapp, uint32 _srcEid, bytes32 _sender, uint64 _nonce, bytes32 _payloadHash) returns()
func (_Endpointv2 *Endpointv2Transactor) Nilify(opts *bind.TransactOpts, _oapp common.Address, _srcEid uint32, _sender [32]byte, _nonce uint64, _payloadHash [32]byte) (*types.Transaction, error) {
	return _Endpointv2.contract.Transact(opts, "nilify", _oapp, _srcEid, _sender, _nonce, _payloadHash)
}

// Nilify is a paid mutator transaction binding the contract method 0x2e80fbf3.
//
// Solidity: function nilify(address _oapp, uint32 _srcEid, bytes32 _sender, uint64 _nonce, bytes32 _payloadHash) returns()
func (_Endpointv2 *Endpointv2Session) Nilify(_oapp common.Address, _srcEid uint32, _sender [32]byte, _nonce uint64, _payloadHash [32]byte) (*types.Transaction, error) {
	return _Endpointv2.Contract.Nilify(&_Endpointv2.TransactOpts, _oapp, _srcEid, _sender, _nonce, _payloadHash)
}

// Nilify is a paid mutator transaction binding the contract method 0x2e80fbf3.
//
// Solidity: function nilify(address _oapp, uint32 _srcEid, bytes32 _sender, uint64 _nonce, bytes32 _payloadHash) returns()
func (_Endpointv2 *Endpointv2TransactorSession) Nilify(_oapp common.Address, _srcEid uint32, _sender [32]byte, _nonce uint64, _payloadHash [32]byte) (*types.Transaction, error) {
	return _Endpointv2.Contract.Nilify(&_Endpointv2.TransactOpts, _oapp, _srcEid, _sender, _nonce, _payloadHash)
}

// RecoverToken is a paid mutator transaction binding the contract method 0xa7229fd9.
//
// Solidity: function recoverToken(address _token, address _to, uint256 _amount) returns()
func (_Endpointv2 *Endpointv2Transactor) RecoverToken(opts *bind.TransactOpts, _token common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Endpointv2.contract.Transact(opts, "recoverToken", _token, _to, _amount)
}

// RecoverToken is a paid mutator transaction binding the contract method 0xa7229fd9.
//
// Solidity: function recoverToken(address _token, address _to, uint256 _amount) returns()
func (_Endpointv2 *Endpointv2Session) RecoverToken(_token common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Endpointv2.Contract.RecoverToken(&_Endpointv2.TransactOpts, _token, _to, _amount)
}

// RecoverToken is a paid mutator transaction binding the contract method 0xa7229fd9.
//
// Solidity: function recoverToken(address _token, address _to, uint256 _amount) returns()
func (_Endpointv2 *Endpointv2TransactorSession) RecoverToken(_token common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Endpointv2.Contract.RecoverToken(&_Endpointv2.TransactOpts, _token, _to, _amount)
}

// RegisterLibrary is a paid mutator transaction binding the contract method 0xe8964e81.
//
// Solidity: function registerLibrary(address _lib) returns()
func (_Endpointv2 *Endpointv2Transactor) RegisterLibrary(opts *bind.TransactOpts, _lib common.Address) (*types.Transaction, error) {
	return _Endpointv2.contract.Transact(opts, "registerLibrary", _lib)
}

// RegisterLibrary is a paid mutator transaction binding the contract method 0xe8964e81.
//
// Solidity: function registerLibrary(address _lib) returns()
func (_Endpointv2 *Endpointv2Session) RegisterLibrary(_lib common.Address) (*types.Transaction, error) {
	return _Endpointv2.Contract.RegisterLibrary(&_Endpointv2.TransactOpts, _lib)
}

// RegisterLibrary is a paid mutator transaction binding the contract method 0xe8964e81.
//
// Solidity: function registerLibrary(address _lib) returns()
func (_Endpointv2 *Endpointv2TransactorSession) RegisterLibrary(_lib common.Address) (*types.Transaction, error) {
	return _Endpointv2.Contract.RegisterLibrary(&_Endpointv2.TransactOpts, _lib)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Endpointv2 *Endpointv2Transactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Endpointv2.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Endpointv2 *Endpointv2Session) RenounceOwnership() (*types.Transaction, error) {
	return _Endpointv2.Contract.RenounceOwnership(&_Endpointv2.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Endpointv2 *Endpointv2TransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Endpointv2.Contract.RenounceOwnership(&_Endpointv2.TransactOpts)
}

// Send is a paid mutator transaction binding the contract method 0x2637a450.
//
// Solidity: function send((uint32,bytes32,bytes,bytes,bool) _params, address _refundAddress) payable returns((bytes32,uint64,(uint256,uint256)))
func (_Endpointv2 *Endpointv2Transactor) Send(opts *bind.TransactOpts, _params MessagingParams, _refundAddress common.Address) (*types.Transaction, error) {
	return _Endpointv2.contract.Transact(opts, "send", _params, _refundAddress)
}

// Send is a paid mutator transaction binding the contract method 0x2637a450.
//
// Solidity: function send((uint32,bytes32,bytes,bytes,bool) _params, address _refundAddress) payable returns((bytes32,uint64,(uint256,uint256)))
func (_Endpointv2 *Endpointv2Session) Send(_params MessagingParams, _refundAddress common.Address) (*types.Transaction, error) {
	return _Endpointv2.Contract.Send(&_Endpointv2.TransactOpts, _params, _refundAddress)
}

// Send is a paid mutator transaction binding the contract method 0x2637a450.
//
// Solidity: function send((uint32,bytes32,bytes,bytes,bool) _params, address _refundAddress) payable returns((bytes32,uint64,(uint256,uint256)))
func (_Endpointv2 *Endpointv2TransactorSession) Send(_params MessagingParams, _refundAddress common.Address) (*types.Transaction, error) {
	return _Endpointv2.Contract.Send(&_Endpointv2.TransactOpts, _params, _refundAddress)
}

// SendCompose is a paid mutator transaction binding the contract method 0x7cb59012.
//
// Solidity: function sendCompose(address _to, bytes32 _guid, uint16 _index, bytes _message) returns()
func (_Endpointv2 *Endpointv2Transactor) SendCompose(opts *bind.TransactOpts, _to common.Address, _guid [32]byte, _index uint16, _message []byte) (*types.Transaction, error) {
	return _Endpointv2.contract.Transact(opts, "sendCompose", _to, _guid, _index, _message)
}

// SendCompose is a paid mutator transaction binding the contract method 0x7cb59012.
//
// Solidity: function sendCompose(address _to, bytes32 _guid, uint16 _index, bytes _message) returns()
func (_Endpointv2 *Endpointv2Session) SendCompose(_to common.Address, _guid [32]byte, _index uint16, _message []byte) (*types.Transaction, error) {
	return _Endpointv2.Contract.SendCompose(&_Endpointv2.TransactOpts, _to, _guid, _index, _message)
}

// SendCompose is a paid mutator transaction binding the contract method 0x7cb59012.
//
// Solidity: function sendCompose(address _to, bytes32 _guid, uint16 _index, bytes _message) returns()
func (_Endpointv2 *Endpointv2TransactorSession) SendCompose(_to common.Address, _guid [32]byte, _index uint16, _message []byte) (*types.Transaction, error) {
	return _Endpointv2.Contract.SendCompose(&_Endpointv2.TransactOpts, _to, _guid, _index, _message)
}

// SetConfig is a paid mutator transaction binding the contract method 0x6dbd9f90.
//
// Solidity: function setConfig(address _oapp, address _lib, (uint32,uint32,bytes)[] _params) returns()
func (_Endpointv2 *Endpointv2Transactor) SetConfig(opts *bind.TransactOpts, _oapp common.Address, _lib common.Address, _params []SetConfigParam) (*types.Transaction, error) {
	return _Endpointv2.contract.Transact(opts, "setConfig", _oapp, _lib, _params)
}

// SetConfig is a paid mutator transaction binding the contract method 0x6dbd9f90.
//
// Solidity: function setConfig(address _oapp, address _lib, (uint32,uint32,bytes)[] _params) returns()
func (_Endpointv2 *Endpointv2Session) SetConfig(_oapp common.Address, _lib common.Address, _params []SetConfigParam) (*types.Transaction, error) {
	return _Endpointv2.Contract.SetConfig(&_Endpointv2.TransactOpts, _oapp, _lib, _params)
}

// SetConfig is a paid mutator transaction binding the contract method 0x6dbd9f90.
//
// Solidity: function setConfig(address _oapp, address _lib, (uint32,uint32,bytes)[] _params) returns()
func (_Endpointv2 *Endpointv2TransactorSession) SetConfig(_oapp common.Address, _lib common.Address, _params []SetConfigParam) (*types.Transaction, error) {
	return _Endpointv2.Contract.SetConfig(&_Endpointv2.TransactOpts, _oapp, _lib, _params)
}

// SetDefaultReceiveLibrary is a paid mutator transaction binding the contract method 0xa718531b.
//
// Solidity: function setDefaultReceiveLibrary(uint32 _eid, address _newLib, uint256 _gracePeriod) returns()
func (_Endpointv2 *Endpointv2Transactor) SetDefaultReceiveLibrary(opts *bind.TransactOpts, _eid uint32, _newLib common.Address, _gracePeriod *big.Int) (*types.Transaction, error) {
	return _Endpointv2.contract.Transact(opts, "setDefaultReceiveLibrary", _eid, _newLib, _gracePeriod)
}

// SetDefaultReceiveLibrary is a paid mutator transaction binding the contract method 0xa718531b.
//
// Solidity: function setDefaultReceiveLibrary(uint32 _eid, address _newLib, uint256 _gracePeriod) returns()
func (_Endpointv2 *Endpointv2Session) SetDefaultReceiveLibrary(_eid uint32, _newLib common.Address, _gracePeriod *big.Int) (*types.Transaction, error) {
	return _Endpointv2.Contract.SetDefaultReceiveLibrary(&_Endpointv2.TransactOpts, _eid, _newLib, _gracePeriod)
}

// SetDefaultReceiveLibrary is a paid mutator transaction binding the contract method 0xa718531b.
//
// Solidity: function setDefaultReceiveLibrary(uint32 _eid, address _newLib, uint256 _gracePeriod) returns()
func (_Endpointv2 *Endpointv2TransactorSession) SetDefaultReceiveLibrary(_eid uint32, _newLib common.Address, _gracePeriod *big.Int) (*types.Transaction, error) {
	return _Endpointv2.Contract.SetDefaultReceiveLibrary(&_Endpointv2.TransactOpts, _eid, _newLib, _gracePeriod)
}

// SetDefaultReceiveLibraryTimeout is a paid mutator transaction binding the contract method 0xd4b4ec8f.
//
// Solidity: function setDefaultReceiveLibraryTimeout(uint32 _eid, address _lib, uint256 _expiry) returns()
func (_Endpointv2 *Endpointv2Transactor) SetDefaultReceiveLibraryTimeout(opts *bind.TransactOpts, _eid uint32, _lib common.Address, _expiry *big.Int) (*types.Transaction, error) {
	return _Endpointv2.contract.Transact(opts, "setDefaultReceiveLibraryTimeout", _eid, _lib, _expiry)
}

// SetDefaultReceiveLibraryTimeout is a paid mutator transaction binding the contract method 0xd4b4ec8f.
//
// Solidity: function setDefaultReceiveLibraryTimeout(uint32 _eid, address _lib, uint256 _expiry) returns()
func (_Endpointv2 *Endpointv2Session) SetDefaultReceiveLibraryTimeout(_eid uint32, _lib common.Address, _expiry *big.Int) (*types.Transaction, error) {
	return _Endpointv2.Contract.SetDefaultReceiveLibraryTimeout(&_Endpointv2.TransactOpts, _eid, _lib, _expiry)
}

// SetDefaultReceiveLibraryTimeout is a paid mutator transaction binding the contract method 0xd4b4ec8f.
//
// Solidity: function setDefaultReceiveLibraryTimeout(uint32 _eid, address _lib, uint256 _expiry) returns()
func (_Endpointv2 *Endpointv2TransactorSession) SetDefaultReceiveLibraryTimeout(_eid uint32, _lib common.Address, _expiry *big.Int) (*types.Transaction, error) {
	return _Endpointv2.Contract.SetDefaultReceiveLibraryTimeout(&_Endpointv2.TransactOpts, _eid, _lib, _expiry)
}

// SetDefaultSendLibrary is a paid mutator transaction binding the contract method 0xaafea312.
//
// Solidity: function setDefaultSendLibrary(uint32 _eid, address _newLib) returns()
func (_Endpointv2 *Endpointv2Transactor) SetDefaultSendLibrary(opts *bind.TransactOpts, _eid uint32, _newLib common.Address) (*types.Transaction, error) {
	return _Endpointv2.contract.Transact(opts, "setDefaultSendLibrary", _eid, _newLib)
}

// SetDefaultSendLibrary is a paid mutator transaction binding the contract method 0xaafea312.
//
// Solidity: function setDefaultSendLibrary(uint32 _eid, address _newLib) returns()
func (_Endpointv2 *Endpointv2Session) SetDefaultSendLibrary(_eid uint32, _newLib common.Address) (*types.Transaction, error) {
	return _Endpointv2.Contract.SetDefaultSendLibrary(&_Endpointv2.TransactOpts, _eid, _newLib)
}

// SetDefaultSendLibrary is a paid mutator transaction binding the contract method 0xaafea312.
//
// Solidity: function setDefaultSendLibrary(uint32 _eid, address _newLib) returns()
func (_Endpointv2 *Endpointv2TransactorSession) SetDefaultSendLibrary(_eid uint32, _newLib common.Address) (*types.Transaction, error) {
	return _Endpointv2.Contract.SetDefaultSendLibrary(&_Endpointv2.TransactOpts, _eid, _newLib)
}

// SetDelegate is a paid mutator transaction binding the contract method 0xca5eb5e1.
//
// Solidity: function setDelegate(address _delegate) returns()
func (_Endpointv2 *Endpointv2Transactor) SetDelegate(opts *bind.TransactOpts, _delegate common.Address) (*types.Transaction, error) {
	return _Endpointv2.contract.Transact(opts, "setDelegate", _delegate)
}

// SetDelegate is a paid mutator transaction binding the contract method 0xca5eb5e1.
//
// Solidity: function setDelegate(address _delegate) returns()
func (_Endpointv2 *Endpointv2Session) SetDelegate(_delegate common.Address) (*types.Transaction, error) {
	return _Endpointv2.Contract.SetDelegate(&_Endpointv2.TransactOpts, _delegate)
}

// SetDelegate is a paid mutator transaction binding the contract method 0xca5eb5e1.
//
// Solidity: function setDelegate(address _delegate) returns()
func (_Endpointv2 *Endpointv2TransactorSession) SetDelegate(_delegate common.Address) (*types.Transaction, error) {
	return _Endpointv2.Contract.SetDelegate(&_Endpointv2.TransactOpts, _delegate)
}

// SetLzToken is a paid mutator transaction binding the contract method 0xc28e0eed.
//
// Solidity: function setLzToken(address _lzToken) returns()
func (_Endpointv2 *Endpointv2Transactor) SetLzToken(opts *bind.TransactOpts, _lzToken common.Address) (*types.Transaction, error) {
	return _Endpointv2.contract.Transact(opts, "setLzToken", _lzToken)
}

// SetLzToken is a paid mutator transaction binding the contract method 0xc28e0eed.
//
// Solidity: function setLzToken(address _lzToken) returns()
func (_Endpointv2 *Endpointv2Session) SetLzToken(_lzToken common.Address) (*types.Transaction, error) {
	return _Endpointv2.Contract.SetLzToken(&_Endpointv2.TransactOpts, _lzToken)
}

// SetLzToken is a paid mutator transaction binding the contract method 0xc28e0eed.
//
// Solidity: function setLzToken(address _lzToken) returns()
func (_Endpointv2 *Endpointv2TransactorSession) SetLzToken(_lzToken common.Address) (*types.Transaction, error) {
	return _Endpointv2.Contract.SetLzToken(&_Endpointv2.TransactOpts, _lzToken)
}

// SetPaused is a paid mutator transaction binding the contract method 0x16c38b3c.
//
// Solidity: function setPaused(bool _paused) returns()
func (_Endpointv2 *Endpointv2Transactor) SetPaused(opts *bind.TransactOpts, _paused bool) (*types.Transaction, error) {
	return _Endpointv2.contract.Transact(opts, "setPaused", _paused)
}

// SetPaused is a paid mutator transaction binding the contract method 0x16c38b3c.
//
// Solidity: function setPaused(bool _paused) returns()
func (_Endpointv2 *Endpointv2Session) SetPaused(_paused bool) (*types.Transaction, error) {
	return _Endpointv2.Contract.SetPaused(&_Endpointv2.TransactOpts, _paused)
}

// SetPaused is a paid mutator transaction binding the contract method 0x16c38b3c.
//
// Solidity: function setPaused(bool _paused) returns()
func (_Endpointv2 *Endpointv2TransactorSession) SetPaused(_paused bool) (*types.Transaction, error) {
	return _Endpointv2.Contract.SetPaused(&_Endpointv2.TransactOpts, _paused)
}

// SetReceiveLibrary is a paid mutator transaction binding the contract method 0x6a14d715.
//
// Solidity: function setReceiveLibrary(address _oapp, uint32 _eid, address _newLib, uint256 _gracePeriod) returns()
func (_Endpointv2 *Endpointv2Transactor) SetReceiveLibrary(opts *bind.TransactOpts, _oapp common.Address, _eid uint32, _newLib common.Address, _gracePeriod *big.Int) (*types.Transaction, error) {
	return _Endpointv2.contract.Transact(opts, "setReceiveLibrary", _oapp, _eid, _newLib, _gracePeriod)
}

// SetReceiveLibrary is a paid mutator transaction binding the contract method 0x6a14d715.
//
// Solidity: function setReceiveLibrary(address _oapp, uint32 _eid, address _newLib, uint256 _gracePeriod) returns()
func (_Endpointv2 *Endpointv2Session) SetReceiveLibrary(_oapp common.Address, _eid uint32, _newLib common.Address, _gracePeriod *big.Int) (*types.Transaction, error) {
	return _Endpointv2.Contract.SetReceiveLibrary(&_Endpointv2.TransactOpts, _oapp, _eid, _newLib, _gracePeriod)
}

// SetReceiveLibrary is a paid mutator transaction binding the contract method 0x6a14d715.
//
// Solidity: function setReceiveLibrary(address _oapp, uint32 _eid, address _newLib, uint256 _gracePeriod) returns()
func (_Endpointv2 *Endpointv2TransactorSession) SetReceiveLibrary(_oapp common.Address, _eid uint32, _newLib common.Address, _gracePeriod *big.Int) (*types.Transaction, error) {
	return _Endpointv2.Contract.SetReceiveLibrary(&_Endpointv2.TransactOpts, _oapp, _eid, _newLib, _gracePeriod)
}

// SetReceiveLibraryTimeout is a paid mutator transaction binding the contract method 0x183c834f.
//
// Solidity: function setReceiveLibraryTimeout(address _oapp, uint32 _eid, address _lib, uint256 _expiry) returns()
func (_Endpointv2 *Endpointv2Transactor) SetReceiveLibraryTimeout(opts *bind.TransactOpts, _oapp common.Address, _eid uint32, _lib common.Address, _expiry *big.Int) (*types.Transaction, error) {
	return _Endpointv2.contract.Transact(opts, "setReceiveLibraryTimeout", _oapp, _eid, _lib, _expiry)
}

// SetReceiveLibraryTimeout is a paid mutator transaction binding the contract method 0x183c834f.
//
// Solidity: function setReceiveLibraryTimeout(address _oapp, uint32 _eid, address _lib, uint256 _expiry) returns()
func (_Endpointv2 *Endpointv2Session) SetReceiveLibraryTimeout(_oapp common.Address, _eid uint32, _lib common.Address, _expiry *big.Int) (*types.Transaction, error) {
	return _Endpointv2.Contract.SetReceiveLibraryTimeout(&_Endpointv2.TransactOpts, _oapp, _eid, _lib, _expiry)
}

// SetReceiveLibraryTimeout is a paid mutator transaction binding the contract method 0x183c834f.
//
// Solidity: function setReceiveLibraryTimeout(address _oapp, uint32 _eid, address _lib, uint256 _expiry) returns()
func (_Endpointv2 *Endpointv2TransactorSession) SetReceiveLibraryTimeout(_oapp common.Address, _eid uint32, _lib common.Address, _expiry *big.Int) (*types.Transaction, error) {
	return _Endpointv2.Contract.SetReceiveLibraryTimeout(&_Endpointv2.TransactOpts, _oapp, _eid, _lib, _expiry)
}

// SetSendLibrary is a paid mutator transaction binding the contract method 0x9535ff30.
//
// Solidity: function setSendLibrary(address _oapp, uint32 _eid, address _newLib) returns()
func (_Endpointv2 *Endpointv2Transactor) SetSendLibrary(opts *bind.TransactOpts, _oapp common.Address, _eid uint32, _newLib common.Address) (*types.Transaction, error) {
	return _Endpointv2.contract.Transact(opts, "setSendLibrary", _oapp, _eid, _newLib)
}

// SetSendLibrary is a paid mutator transaction binding the contract method 0x9535ff30.
//
// Solidity: function setSendLibrary(address _oapp, uint32 _eid, address _newLib) returns()
func (_Endpointv2 *Endpointv2Session) SetSendLibrary(_oapp common.Address, _eid uint32, _newLib common.Address) (*types.Transaction, error) {
	return _Endpointv2.Contract.SetSendLibrary(&_Endpointv2.TransactOpts, _oapp, _eid, _newLib)
}

// SetSendLibrary is a paid mutator transaction binding the contract method 0x9535ff30.
//
// Solidity: function setSendLibrary(address _oapp, uint32 _eid, address _newLib) returns()
func (_Endpointv2 *Endpointv2TransactorSession) SetSendLibrary(_oapp common.Address, _eid uint32, _newLib common.Address) (*types.Transaction, error) {
	return _Endpointv2.Contract.SetSendLibrary(&_Endpointv2.TransactOpts, _oapp, _eid, _newLib)
}

// Skip is a paid mutator transaction binding the contract method 0xd70b8902.
//
// Solidity: function skip(address _oapp, uint32 _srcEid, bytes32 _sender, uint64 _nonce) returns()
func (_Endpointv2 *Endpointv2Transactor) Skip(opts *bind.TransactOpts, _oapp common.Address, _srcEid uint32, _sender [32]byte, _nonce uint64) (*types.Transaction, error) {
	return _Endpointv2.contract.Transact(opts, "skip", _oapp, _srcEid, _sender, _nonce)
}

// Skip is a paid mutator transaction binding the contract method 0xd70b8902.
//
// Solidity: function skip(address _oapp, uint32 _srcEid, bytes32 _sender, uint64 _nonce) returns()
func (_Endpointv2 *Endpointv2Session) Skip(_oapp common.Address, _srcEid uint32, _sender [32]byte, _nonce uint64) (*types.Transaction, error) {
	return _Endpointv2.Contract.Skip(&_Endpointv2.TransactOpts, _oapp, _srcEid, _sender, _nonce)
}

// Skip is a paid mutator transaction binding the contract method 0xd70b8902.
//
// Solidity: function skip(address _oapp, uint32 _srcEid, bytes32 _sender, uint64 _nonce) returns()
func (_Endpointv2 *Endpointv2TransactorSession) Skip(_oapp common.Address, _srcEid uint32, _sender [32]byte, _nonce uint64) (*types.Transaction, error) {
	return _Endpointv2.Contract.Skip(&_Endpointv2.TransactOpts, _oapp, _srcEid, _sender, _nonce)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Endpointv2 *Endpointv2Transactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Endpointv2.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Endpointv2 *Endpointv2Session) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Endpointv2.Contract.TransferOwnership(&_Endpointv2.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Endpointv2 *Endpointv2TransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Endpointv2.Contract.TransferOwnership(&_Endpointv2.TransactOpts, newOwner)
}

// Verify is a paid mutator transaction binding the contract method 0xa825d747.
//
// Solidity: function verify((uint32,bytes32,uint64) _origin, address _receiver, bytes32 _payloadHash) returns()
func (_Endpointv2 *Endpointv2Transactor) Verify(opts *bind.TransactOpts, _origin Origin, _receiver common.Address, _payloadHash [32]byte) (*types.Transaction, error) {
	return _Endpointv2.contract.Transact(opts, "verify", _origin, _receiver, _payloadHash)
}

// Verify is a paid mutator transaction binding the contract method 0xa825d747.
//
// Solidity: function verify((uint32,bytes32,uint64) _origin, address _receiver, bytes32 _payloadHash) returns()
func (_Endpointv2 *Endpointv2Session) Verify(_origin Origin, _receiver common.Address, _payloadHash [32]byte) (*types.Transaction, error) {
	return _Endpointv2.Contract.Verify(&_Endpointv2.TransactOpts, _origin, _receiver, _payloadHash)
}

// Verify is a paid mutator transaction binding the contract method 0xa825d747.
//
// Solidity: function verify((uint32,bytes32,uint64) _origin, address _receiver, bytes32 _payloadHash) returns()
func (_Endpointv2 *Endpointv2TransactorSession) Verify(_origin Origin, _receiver common.Address, _payloadHash [32]byte) (*types.Transaction, error) {
	return _Endpointv2.Contract.Verify(&_Endpointv2.TransactOpts, _origin, _receiver, _payloadHash)
}

// Endpointv2ComposeDeliveredIterator is returned from FilterComposeDelivered and is used to iterate over the raw logs and unpacked data for ComposeDelivered events raised by the Endpointv2 contract.
type Endpointv2ComposeDeliveredIterator struct {
	Event *Endpointv2ComposeDelivered // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *Endpointv2ComposeDeliveredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Endpointv2ComposeDelivered)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(Endpointv2ComposeDelivered)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *Endpointv2ComposeDeliveredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Endpointv2ComposeDeliveredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Endpointv2ComposeDelivered represents a ComposeDelivered event raised by the Endpointv2 contract.
type Endpointv2ComposeDelivered struct {
	From  common.Address
	To    common.Address
	Guid  [32]byte
	Index uint16
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterComposeDelivered is a free log retrieval operation binding the contract event 0x0036c98efcf9e6641dfbc9051f66f405253e8e0c2ab4a24dccda15595b7378c8.
//
// Solidity: event ComposeDelivered(address from, address to, bytes32 guid, uint16 index)
func (_Endpointv2 *Endpointv2Filterer) FilterComposeDelivered(opts *bind.FilterOpts) (*Endpointv2ComposeDeliveredIterator, error) {

	logs, sub, err := _Endpointv2.contract.FilterLogs(opts, "ComposeDelivered")
	if err != nil {
		return nil, err
	}
	return &Endpointv2ComposeDeliveredIterator{contract: _Endpointv2.contract, event: "ComposeDelivered", logs: logs, sub: sub}, nil
}

// WatchComposeDelivered is a free log subscription operation binding the contract event 0x0036c98efcf9e6641dfbc9051f66f405253e8e0c2ab4a24dccda15595b7378c8.
//
// Solidity: event ComposeDelivered(address from, address to, bytes32 guid, uint16 index)
func (_Endpointv2 *Endpointv2Filterer) WatchComposeDelivered(opts *bind.WatchOpts, sink chan<- *Endpointv2ComposeDelivered) (event.Subscription, error) {

	logs, sub, err := _Endpointv2.contract.WatchLogs(opts, "ComposeDelivered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Endpointv2ComposeDelivered)
				if err := _Endpointv2.contract.UnpackLog(event, "ComposeDelivered", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseComposeDelivered is a log parse operation binding the contract event 0x0036c98efcf9e6641dfbc9051f66f405253e8e0c2ab4a24dccda15595b7378c8.
//
// Solidity: event ComposeDelivered(address from, address to, bytes32 guid, uint16 index)
func (_Endpointv2 *Endpointv2Filterer) ParseComposeDelivered(log types.Log) (*Endpointv2ComposeDelivered, error) {
	event := new(Endpointv2ComposeDelivered)
	if err := _Endpointv2.contract.UnpackLog(event, "ComposeDelivered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Endpointv2ComposeSentIterator is returned from FilterComposeSent and is used to iterate over the raw logs and unpacked data for ComposeSent events raised by the Endpointv2 contract.
type Endpointv2ComposeSentIterator struct {
	Event *Endpointv2ComposeSent // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *Endpointv2ComposeSentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Endpointv2ComposeSent)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(Endpointv2ComposeSent)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *Endpointv2ComposeSentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Endpointv2ComposeSentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Endpointv2ComposeSent represents a ComposeSent event raised by the Endpointv2 contract.
type Endpointv2ComposeSent struct {
	From    common.Address
	To      common.Address
	Guid    [32]byte
	Index   uint16
	Message []byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterComposeSent is a free log retrieval operation binding the contract event 0x3d52ff888d033fd3dd1d8057da59e850c91d91a72c41dfa445b247dfedeb6dc1.
//
// Solidity: event ComposeSent(address from, address to, bytes32 guid, uint16 index, bytes message)
func (_Endpointv2 *Endpointv2Filterer) FilterComposeSent(opts *bind.FilterOpts) (*Endpointv2ComposeSentIterator, error) {

	logs, sub, err := _Endpointv2.contract.FilterLogs(opts, "ComposeSent")
	if err != nil {
		return nil, err
	}
	return &Endpointv2ComposeSentIterator{contract: _Endpointv2.contract, event: "ComposeSent", logs: logs, sub: sub}, nil
}

// WatchComposeSent is a free log subscription operation binding the contract event 0x3d52ff888d033fd3dd1d8057da59e850c91d91a72c41dfa445b247dfedeb6dc1.
//
// Solidity: event ComposeSent(address from, address to, bytes32 guid, uint16 index, bytes message)
func (_Endpointv2 *Endpointv2Filterer) WatchComposeSent(opts *bind.WatchOpts, sink chan<- *Endpointv2ComposeSent) (event.Subscription, error) {

	logs, sub, err := _Endpointv2.contract.WatchLogs(opts, "ComposeSent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Endpointv2ComposeSent)
				if err := _Endpointv2.contract.UnpackLog(event, "ComposeSent", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseComposeSent is a log parse operation binding the contract event 0x3d52ff888d033fd3dd1d8057da59e850c91d91a72c41dfa445b247dfedeb6dc1.
//
// Solidity: event ComposeSent(address from, address to, bytes32 guid, uint16 index, bytes message)
func (_Endpointv2 *Endpointv2Filterer) ParseComposeSent(log types.Log) (*Endpointv2ComposeSent, error) {
	event := new(Endpointv2ComposeSent)
	if err := _Endpointv2.contract.UnpackLog(event, "ComposeSent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Endpointv2DefaultReceiveLibrarySetIterator is returned from FilterDefaultReceiveLibrarySet and is used to iterate over the raw logs and unpacked data for DefaultReceiveLibrarySet events raised by the Endpointv2 contract.
type Endpointv2DefaultReceiveLibrarySetIterator struct {
	Event *Endpointv2DefaultReceiveLibrarySet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *Endpointv2DefaultReceiveLibrarySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Endpointv2DefaultReceiveLibrarySet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(Endpointv2DefaultReceiveLibrarySet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *Endpointv2DefaultReceiveLibrarySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Endpointv2DefaultReceiveLibrarySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Endpointv2DefaultReceiveLibrarySet represents a DefaultReceiveLibrarySet event raised by the Endpointv2 contract.
type Endpointv2DefaultReceiveLibrarySet struct {
	Eid    uint32
	NewLib common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDefaultReceiveLibrarySet is a free log retrieval operation binding the contract event 0xc16891855cffb4a5ac51ac11864a3f3c96ba816cc45fe686c987ae36277de5ec.
//
// Solidity: event DefaultReceiveLibrarySet(uint32 eid, address newLib)
func (_Endpointv2 *Endpointv2Filterer) FilterDefaultReceiveLibrarySet(opts *bind.FilterOpts) (*Endpointv2DefaultReceiveLibrarySetIterator, error) {

	logs, sub, err := _Endpointv2.contract.FilterLogs(opts, "DefaultReceiveLibrarySet")
	if err != nil {
		return nil, err
	}
	return &Endpointv2DefaultReceiveLibrarySetIterator{contract: _Endpointv2.contract, event: "DefaultReceiveLibrarySet", logs: logs, sub: sub}, nil
}

// WatchDefaultReceiveLibrarySet is a free log subscription operation binding the contract event 0xc16891855cffb4a5ac51ac11864a3f3c96ba816cc45fe686c987ae36277de5ec.
//
// Solidity: event DefaultReceiveLibrarySet(uint32 eid, address newLib)
func (_Endpointv2 *Endpointv2Filterer) WatchDefaultReceiveLibrarySet(opts *bind.WatchOpts, sink chan<- *Endpointv2DefaultReceiveLibrarySet) (event.Subscription, error) {

	logs, sub, err := _Endpointv2.contract.WatchLogs(opts, "DefaultReceiveLibrarySet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Endpointv2DefaultReceiveLibrarySet)
				if err := _Endpointv2.contract.UnpackLog(event, "DefaultReceiveLibrarySet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDefaultReceiveLibrarySet is a log parse operation binding the contract event 0xc16891855cffb4a5ac51ac11864a3f3c96ba816cc45fe686c987ae36277de5ec.
//
// Solidity: event DefaultReceiveLibrarySet(uint32 eid, address newLib)
func (_Endpointv2 *Endpointv2Filterer) ParseDefaultReceiveLibrarySet(log types.Log) (*Endpointv2DefaultReceiveLibrarySet, error) {
	event := new(Endpointv2DefaultReceiveLibrarySet)
	if err := _Endpointv2.contract.UnpackLog(event, "DefaultReceiveLibrarySet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Endpointv2DefaultReceiveLibraryTimeoutSetIterator is returned from FilterDefaultReceiveLibraryTimeoutSet and is used to iterate over the raw logs and unpacked data for DefaultReceiveLibraryTimeoutSet events raised by the Endpointv2 contract.
type Endpointv2DefaultReceiveLibraryTimeoutSetIterator struct {
	Event *Endpointv2DefaultReceiveLibraryTimeoutSet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *Endpointv2DefaultReceiveLibraryTimeoutSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Endpointv2DefaultReceiveLibraryTimeoutSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(Endpointv2DefaultReceiveLibraryTimeoutSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *Endpointv2DefaultReceiveLibraryTimeoutSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Endpointv2DefaultReceiveLibraryTimeoutSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Endpointv2DefaultReceiveLibraryTimeoutSet represents a DefaultReceiveLibraryTimeoutSet event raised by the Endpointv2 contract.
type Endpointv2DefaultReceiveLibraryTimeoutSet struct {
	Eid    uint32
	OldLib common.Address
	Expiry *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDefaultReceiveLibraryTimeoutSet is a free log retrieval operation binding the contract event 0x55b28633cdb29709386f555dfc54418592ad475ce7a65a78ac5928af60ffb8f8.
//
// Solidity: event DefaultReceiveLibraryTimeoutSet(uint32 eid, address oldLib, uint256 expiry)
func (_Endpointv2 *Endpointv2Filterer) FilterDefaultReceiveLibraryTimeoutSet(opts *bind.FilterOpts) (*Endpointv2DefaultReceiveLibraryTimeoutSetIterator, error) {

	logs, sub, err := _Endpointv2.contract.FilterLogs(opts, "DefaultReceiveLibraryTimeoutSet")
	if err != nil {
		return nil, err
	}
	return &Endpointv2DefaultReceiveLibraryTimeoutSetIterator{contract: _Endpointv2.contract, event: "DefaultReceiveLibraryTimeoutSet", logs: logs, sub: sub}, nil
}

// WatchDefaultReceiveLibraryTimeoutSet is a free log subscription operation binding the contract event 0x55b28633cdb29709386f555dfc54418592ad475ce7a65a78ac5928af60ffb8f8.
//
// Solidity: event DefaultReceiveLibraryTimeoutSet(uint32 eid, address oldLib, uint256 expiry)
func (_Endpointv2 *Endpointv2Filterer) WatchDefaultReceiveLibraryTimeoutSet(opts *bind.WatchOpts, sink chan<- *Endpointv2DefaultReceiveLibraryTimeoutSet) (event.Subscription, error) {

	logs, sub, err := _Endpointv2.contract.WatchLogs(opts, "DefaultReceiveLibraryTimeoutSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Endpointv2DefaultReceiveLibraryTimeoutSet)
				if err := _Endpointv2.contract.UnpackLog(event, "DefaultReceiveLibraryTimeoutSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDefaultReceiveLibraryTimeoutSet is a log parse operation binding the contract event 0x55b28633cdb29709386f555dfc54418592ad475ce7a65a78ac5928af60ffb8f8.
//
// Solidity: event DefaultReceiveLibraryTimeoutSet(uint32 eid, address oldLib, uint256 expiry)
func (_Endpointv2 *Endpointv2Filterer) ParseDefaultReceiveLibraryTimeoutSet(log types.Log) (*Endpointv2DefaultReceiveLibraryTimeoutSet, error) {
	event := new(Endpointv2DefaultReceiveLibraryTimeoutSet)
	if err := _Endpointv2.contract.UnpackLog(event, "DefaultReceiveLibraryTimeoutSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Endpointv2DefaultSendLibrarySetIterator is returned from FilterDefaultSendLibrarySet and is used to iterate over the raw logs and unpacked data for DefaultSendLibrarySet events raised by the Endpointv2 contract.
type Endpointv2DefaultSendLibrarySetIterator struct {
	Event *Endpointv2DefaultSendLibrarySet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *Endpointv2DefaultSendLibrarySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Endpointv2DefaultSendLibrarySet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(Endpointv2DefaultSendLibrarySet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *Endpointv2DefaultSendLibrarySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Endpointv2DefaultSendLibrarySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Endpointv2DefaultSendLibrarySet represents a DefaultSendLibrarySet event raised by the Endpointv2 contract.
type Endpointv2DefaultSendLibrarySet struct {
	Eid    uint32
	NewLib common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDefaultSendLibrarySet is a free log retrieval operation binding the contract event 0x16aa0f528038ab41019e95bae5b418a50ba8532c5800e3b7ea2f517d3fa625f5.
//
// Solidity: event DefaultSendLibrarySet(uint32 eid, address newLib)
func (_Endpointv2 *Endpointv2Filterer) FilterDefaultSendLibrarySet(opts *bind.FilterOpts) (*Endpointv2DefaultSendLibrarySetIterator, error) {

	logs, sub, err := _Endpointv2.contract.FilterLogs(opts, "DefaultSendLibrarySet")
	if err != nil {
		return nil, err
	}
	return &Endpointv2DefaultSendLibrarySetIterator{contract: _Endpointv2.contract, event: "DefaultSendLibrarySet", logs: logs, sub: sub}, nil
}

// WatchDefaultSendLibrarySet is a free log subscription operation binding the contract event 0x16aa0f528038ab41019e95bae5b418a50ba8532c5800e3b7ea2f517d3fa625f5.
//
// Solidity: event DefaultSendLibrarySet(uint32 eid, address newLib)
func (_Endpointv2 *Endpointv2Filterer) WatchDefaultSendLibrarySet(opts *bind.WatchOpts, sink chan<- *Endpointv2DefaultSendLibrarySet) (event.Subscription, error) {

	logs, sub, err := _Endpointv2.contract.WatchLogs(opts, "DefaultSendLibrarySet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Endpointv2DefaultSendLibrarySet)
				if err := _Endpointv2.contract.UnpackLog(event, "DefaultSendLibrarySet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDefaultSendLibrarySet is a log parse operation binding the contract event 0x16aa0f528038ab41019e95bae5b418a50ba8532c5800e3b7ea2f517d3fa625f5.
//
// Solidity: event DefaultSendLibrarySet(uint32 eid, address newLib)
func (_Endpointv2 *Endpointv2Filterer) ParseDefaultSendLibrarySet(log types.Log) (*Endpointv2DefaultSendLibrarySet, error) {
	event := new(Endpointv2DefaultSendLibrarySet)
	if err := _Endpointv2.contract.UnpackLog(event, "DefaultSendLibrarySet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Endpointv2DelegateSetIterator is returned from FilterDelegateSet and is used to iterate over the raw logs and unpacked data for DelegateSet events raised by the Endpointv2 contract.
type Endpointv2DelegateSetIterator struct {
	Event *Endpointv2DelegateSet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *Endpointv2DelegateSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Endpointv2DelegateSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(Endpointv2DelegateSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *Endpointv2DelegateSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Endpointv2DelegateSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Endpointv2DelegateSet represents a DelegateSet event raised by the Endpointv2 contract.
type Endpointv2DelegateSet struct {
	Sender   common.Address
	Delegate common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterDelegateSet is a free log retrieval operation binding the contract event 0x6ee10e9ed4d6ce9742703a498707862f4b00f1396a87195eb93267b3d7983981.
//
// Solidity: event DelegateSet(address sender, address delegate)
func (_Endpointv2 *Endpointv2Filterer) FilterDelegateSet(opts *bind.FilterOpts) (*Endpointv2DelegateSetIterator, error) {

	logs, sub, err := _Endpointv2.contract.FilterLogs(opts, "DelegateSet")
	if err != nil {
		return nil, err
	}
	return &Endpointv2DelegateSetIterator{contract: _Endpointv2.contract, event: "DelegateSet", logs: logs, sub: sub}, nil
}

// WatchDelegateSet is a free log subscription operation binding the contract event 0x6ee10e9ed4d6ce9742703a498707862f4b00f1396a87195eb93267b3d7983981.
//
// Solidity: event DelegateSet(address sender, address delegate)
func (_Endpointv2 *Endpointv2Filterer) WatchDelegateSet(opts *bind.WatchOpts, sink chan<- *Endpointv2DelegateSet) (event.Subscription, error) {

	logs, sub, err := _Endpointv2.contract.WatchLogs(opts, "DelegateSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Endpointv2DelegateSet)
				if err := _Endpointv2.contract.UnpackLog(event, "DelegateSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDelegateSet is a log parse operation binding the contract event 0x6ee10e9ed4d6ce9742703a498707862f4b00f1396a87195eb93267b3d7983981.
//
// Solidity: event DelegateSet(address sender, address delegate)
func (_Endpointv2 *Endpointv2Filterer) ParseDelegateSet(log types.Log) (*Endpointv2DelegateSet, error) {
	event := new(Endpointv2DelegateSet)
	if err := _Endpointv2.contract.UnpackLog(event, "DelegateSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Endpointv2InboundNonceSkippedIterator is returned from FilterInboundNonceSkipped and is used to iterate over the raw logs and unpacked data for InboundNonceSkipped events raised by the Endpointv2 contract.
type Endpointv2InboundNonceSkippedIterator struct {
	Event *Endpointv2InboundNonceSkipped // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *Endpointv2InboundNonceSkippedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Endpointv2InboundNonceSkipped)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(Endpointv2InboundNonceSkipped)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *Endpointv2InboundNonceSkippedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Endpointv2InboundNonceSkippedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Endpointv2InboundNonceSkipped represents a InboundNonceSkipped event raised by the Endpointv2 contract.
type Endpointv2InboundNonceSkipped struct {
	SrcEid   uint32
	Sender   [32]byte
	Receiver common.Address
	Nonce    uint64
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterInboundNonceSkipped is a free log retrieval operation binding the contract event 0x28f40053783033ef755556a0c3315379141f51a33aed8334174ffbadd90bde48.
//
// Solidity: event InboundNonceSkipped(uint32 srcEid, bytes32 sender, address receiver, uint64 nonce)
func (_Endpointv2 *Endpointv2Filterer) FilterInboundNonceSkipped(opts *bind.FilterOpts) (*Endpointv2InboundNonceSkippedIterator, error) {

	logs, sub, err := _Endpointv2.contract.FilterLogs(opts, "InboundNonceSkipped")
	if err != nil {
		return nil, err
	}
	return &Endpointv2InboundNonceSkippedIterator{contract: _Endpointv2.contract, event: "InboundNonceSkipped", logs: logs, sub: sub}, nil
}

// WatchInboundNonceSkipped is a free log subscription operation binding the contract event 0x28f40053783033ef755556a0c3315379141f51a33aed8334174ffbadd90bde48.
//
// Solidity: event InboundNonceSkipped(uint32 srcEid, bytes32 sender, address receiver, uint64 nonce)
func (_Endpointv2 *Endpointv2Filterer) WatchInboundNonceSkipped(opts *bind.WatchOpts, sink chan<- *Endpointv2InboundNonceSkipped) (event.Subscription, error) {

	logs, sub, err := _Endpointv2.contract.WatchLogs(opts, "InboundNonceSkipped")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Endpointv2InboundNonceSkipped)
				if err := _Endpointv2.contract.UnpackLog(event, "InboundNonceSkipped", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInboundNonceSkipped is a log parse operation binding the contract event 0x28f40053783033ef755556a0c3315379141f51a33aed8334174ffbadd90bde48.
//
// Solidity: event InboundNonceSkipped(uint32 srcEid, bytes32 sender, address receiver, uint64 nonce)
func (_Endpointv2 *Endpointv2Filterer) ParseInboundNonceSkipped(log types.Log) (*Endpointv2InboundNonceSkipped, error) {
	event := new(Endpointv2InboundNonceSkipped)
	if err := _Endpointv2.contract.UnpackLog(event, "InboundNonceSkipped", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Endpointv2LibraryRegisteredIterator is returned from FilterLibraryRegistered and is used to iterate over the raw logs and unpacked data for LibraryRegistered events raised by the Endpointv2 contract.
type Endpointv2LibraryRegisteredIterator struct {
	Event *Endpointv2LibraryRegistered // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *Endpointv2LibraryRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Endpointv2LibraryRegistered)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(Endpointv2LibraryRegistered)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *Endpointv2LibraryRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Endpointv2LibraryRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Endpointv2LibraryRegistered represents a LibraryRegistered event raised by the Endpointv2 contract.
type Endpointv2LibraryRegistered struct {
	NewLib common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterLibraryRegistered is a free log retrieval operation binding the contract event 0x6b374d56679ca9463f27c85c6311e2bb7fde69bf201d3da39d53f10bd9d78af5.
//
// Solidity: event LibraryRegistered(address newLib)
func (_Endpointv2 *Endpointv2Filterer) FilterLibraryRegistered(opts *bind.FilterOpts) (*Endpointv2LibraryRegisteredIterator, error) {

	logs, sub, err := _Endpointv2.contract.FilterLogs(opts, "LibraryRegistered")
	if err != nil {
		return nil, err
	}
	return &Endpointv2LibraryRegisteredIterator{contract: _Endpointv2.contract, event: "LibraryRegistered", logs: logs, sub: sub}, nil
}

// WatchLibraryRegistered is a free log subscription operation binding the contract event 0x6b374d56679ca9463f27c85c6311e2bb7fde69bf201d3da39d53f10bd9d78af5.
//
// Solidity: event LibraryRegistered(address newLib)
func (_Endpointv2 *Endpointv2Filterer) WatchLibraryRegistered(opts *bind.WatchOpts, sink chan<- *Endpointv2LibraryRegistered) (event.Subscription, error) {

	logs, sub, err := _Endpointv2.contract.WatchLogs(opts, "LibraryRegistered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Endpointv2LibraryRegistered)
				if err := _Endpointv2.contract.UnpackLog(event, "LibraryRegistered", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLibraryRegistered is a log parse operation binding the contract event 0x6b374d56679ca9463f27c85c6311e2bb7fde69bf201d3da39d53f10bd9d78af5.
//
// Solidity: event LibraryRegistered(address newLib)
func (_Endpointv2 *Endpointv2Filterer) ParseLibraryRegistered(log types.Log) (*Endpointv2LibraryRegistered, error) {
	event := new(Endpointv2LibraryRegistered)
	if err := _Endpointv2.contract.UnpackLog(event, "LibraryRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Endpointv2LzComposeAlertIterator is returned from FilterLzComposeAlert and is used to iterate over the raw logs and unpacked data for LzComposeAlert events raised by the Endpointv2 contract.
type Endpointv2LzComposeAlertIterator struct {
	Event *Endpointv2LzComposeAlert // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *Endpointv2LzComposeAlertIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Endpointv2LzComposeAlert)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(Endpointv2LzComposeAlert)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *Endpointv2LzComposeAlertIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Endpointv2LzComposeAlertIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Endpointv2LzComposeAlert represents a LzComposeAlert event raised by the Endpointv2 contract.
type Endpointv2LzComposeAlert struct {
	From      common.Address
	To        common.Address
	Executor  common.Address
	Guid      [32]byte
	Index     uint16
	Gas       *big.Int
	Value     *big.Int
	Message   []byte
	ExtraData []byte
	Reason    []byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterLzComposeAlert is a free log retrieval operation binding the contract event 0x8a0b1dce321c5c5fb42349bce46d18087c04140de520917661fb923e44a904b9.
//
// Solidity: event LzComposeAlert(address indexed from, address indexed to, address indexed executor, bytes32 guid, uint16 index, uint256 gas, uint256 value, bytes message, bytes extraData, bytes reason)
func (_Endpointv2 *Endpointv2Filterer) FilterLzComposeAlert(opts *bind.FilterOpts, from []common.Address, to []common.Address, executor []common.Address) (*Endpointv2LzComposeAlertIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var executorRule []interface{}
	for _, executorItem := range executor {
		executorRule = append(executorRule, executorItem)
	}

	logs, sub, err := _Endpointv2.contract.FilterLogs(opts, "LzComposeAlert", fromRule, toRule, executorRule)
	if err != nil {
		return nil, err
	}
	return &Endpointv2LzComposeAlertIterator{contract: _Endpointv2.contract, event: "LzComposeAlert", logs: logs, sub: sub}, nil
}

// WatchLzComposeAlert is a free log subscription operation binding the contract event 0x8a0b1dce321c5c5fb42349bce46d18087c04140de520917661fb923e44a904b9.
//
// Solidity: event LzComposeAlert(address indexed from, address indexed to, address indexed executor, bytes32 guid, uint16 index, uint256 gas, uint256 value, bytes message, bytes extraData, bytes reason)
func (_Endpointv2 *Endpointv2Filterer) WatchLzComposeAlert(opts *bind.WatchOpts, sink chan<- *Endpointv2LzComposeAlert, from []common.Address, to []common.Address, executor []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var executorRule []interface{}
	for _, executorItem := range executor {
		executorRule = append(executorRule, executorItem)
	}

	logs, sub, err := _Endpointv2.contract.WatchLogs(opts, "LzComposeAlert", fromRule, toRule, executorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Endpointv2LzComposeAlert)
				if err := _Endpointv2.contract.UnpackLog(event, "LzComposeAlert", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLzComposeAlert is a log parse operation binding the contract event 0x8a0b1dce321c5c5fb42349bce46d18087c04140de520917661fb923e44a904b9.
//
// Solidity: event LzComposeAlert(address indexed from, address indexed to, address indexed executor, bytes32 guid, uint16 index, uint256 gas, uint256 value, bytes message, bytes extraData, bytes reason)
func (_Endpointv2 *Endpointv2Filterer) ParseLzComposeAlert(log types.Log) (*Endpointv2LzComposeAlert, error) {
	event := new(Endpointv2LzComposeAlert)
	if err := _Endpointv2.contract.UnpackLog(event, "LzComposeAlert", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Endpointv2LzReceiveAlertIterator is returned from FilterLzReceiveAlert and is used to iterate over the raw logs and unpacked data for LzReceiveAlert events raised by the Endpointv2 contract.
type Endpointv2LzReceiveAlertIterator struct {
	Event *Endpointv2LzReceiveAlert // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *Endpointv2LzReceiveAlertIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Endpointv2LzReceiveAlert)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(Endpointv2LzReceiveAlert)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *Endpointv2LzReceiveAlertIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Endpointv2LzReceiveAlertIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Endpointv2LzReceiveAlert represents a LzReceiveAlert event raised by the Endpointv2 contract.
type Endpointv2LzReceiveAlert struct {
	Receiver  common.Address
	Executor  common.Address
	Origin    Origin
	Guid      [32]byte
	Gas       *big.Int
	Value     *big.Int
	Message   []byte
	ExtraData []byte
	Reason    []byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterLzReceiveAlert is a free log retrieval operation binding the contract event 0x7edfa10fe10193301ad8a8bea7e968c7bcabcc64981f368e3aeada40ce26ae2c.
//
// Solidity: event LzReceiveAlert(address indexed receiver, address indexed executor, (uint32,bytes32,uint64) origin, bytes32 guid, uint256 gas, uint256 value, bytes message, bytes extraData, bytes reason)
func (_Endpointv2 *Endpointv2Filterer) FilterLzReceiveAlert(opts *bind.FilterOpts, receiver []common.Address, executor []common.Address) (*Endpointv2LzReceiveAlertIterator, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}
	var executorRule []interface{}
	for _, executorItem := range executor {
		executorRule = append(executorRule, executorItem)
	}

	logs, sub, err := _Endpointv2.contract.FilterLogs(opts, "LzReceiveAlert", receiverRule, executorRule)
	if err != nil {
		return nil, err
	}
	return &Endpointv2LzReceiveAlertIterator{contract: _Endpointv2.contract, event: "LzReceiveAlert", logs: logs, sub: sub}, nil
}

// WatchLzReceiveAlert is a free log subscription operation binding the contract event 0x7edfa10fe10193301ad8a8bea7e968c7bcabcc64981f368e3aeada40ce26ae2c.
//
// Solidity: event LzReceiveAlert(address indexed receiver, address indexed executor, (uint32,bytes32,uint64) origin, bytes32 guid, uint256 gas, uint256 value, bytes message, bytes extraData, bytes reason)
func (_Endpointv2 *Endpointv2Filterer) WatchLzReceiveAlert(opts *bind.WatchOpts, sink chan<- *Endpointv2LzReceiveAlert, receiver []common.Address, executor []common.Address) (event.Subscription, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}
	var executorRule []interface{}
	for _, executorItem := range executor {
		executorRule = append(executorRule, executorItem)
	}

	logs, sub, err := _Endpointv2.contract.WatchLogs(opts, "LzReceiveAlert", receiverRule, executorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Endpointv2LzReceiveAlert)
				if err := _Endpointv2.contract.UnpackLog(event, "LzReceiveAlert", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLzReceiveAlert is a log parse operation binding the contract event 0x7edfa10fe10193301ad8a8bea7e968c7bcabcc64981f368e3aeada40ce26ae2c.
//
// Solidity: event LzReceiveAlert(address indexed receiver, address indexed executor, (uint32,bytes32,uint64) origin, bytes32 guid, uint256 gas, uint256 value, bytes message, bytes extraData, bytes reason)
func (_Endpointv2 *Endpointv2Filterer) ParseLzReceiveAlert(log types.Log) (*Endpointv2LzReceiveAlert, error) {
	event := new(Endpointv2LzReceiveAlert)
	if err := _Endpointv2.contract.UnpackLog(event, "LzReceiveAlert", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Endpointv2LzTokenSetIterator is returned from FilterLzTokenSet and is used to iterate over the raw logs and unpacked data for LzTokenSet events raised by the Endpointv2 contract.
type Endpointv2LzTokenSetIterator struct {
	Event *Endpointv2LzTokenSet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *Endpointv2LzTokenSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Endpointv2LzTokenSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(Endpointv2LzTokenSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *Endpointv2LzTokenSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Endpointv2LzTokenSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Endpointv2LzTokenSet represents a LzTokenSet event raised by the Endpointv2 contract.
type Endpointv2LzTokenSet struct {
	Token common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterLzTokenSet is a free log retrieval operation binding the contract event 0xd476ec5ec1ac11cec3714d41e7ea49419471aceb9bd0dff1becfc3e363a62396.
//
// Solidity: event LzTokenSet(address token)
func (_Endpointv2 *Endpointv2Filterer) FilterLzTokenSet(opts *bind.FilterOpts) (*Endpointv2LzTokenSetIterator, error) {

	logs, sub, err := _Endpointv2.contract.FilterLogs(opts, "LzTokenSet")
	if err != nil {
		return nil, err
	}
	return &Endpointv2LzTokenSetIterator{contract: _Endpointv2.contract, event: "LzTokenSet", logs: logs, sub: sub}, nil
}

// WatchLzTokenSet is a free log subscription operation binding the contract event 0xd476ec5ec1ac11cec3714d41e7ea49419471aceb9bd0dff1becfc3e363a62396.
//
// Solidity: event LzTokenSet(address token)
func (_Endpointv2 *Endpointv2Filterer) WatchLzTokenSet(opts *bind.WatchOpts, sink chan<- *Endpointv2LzTokenSet) (event.Subscription, error) {

	logs, sub, err := _Endpointv2.contract.WatchLogs(opts, "LzTokenSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Endpointv2LzTokenSet)
				if err := _Endpointv2.contract.UnpackLog(event, "LzTokenSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLzTokenSet is a log parse operation binding the contract event 0xd476ec5ec1ac11cec3714d41e7ea49419471aceb9bd0dff1becfc3e363a62396.
//
// Solidity: event LzTokenSet(address token)
func (_Endpointv2 *Endpointv2Filterer) ParseLzTokenSet(log types.Log) (*Endpointv2LzTokenSet, error) {
	event := new(Endpointv2LzTokenSet)
	if err := _Endpointv2.contract.UnpackLog(event, "LzTokenSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Endpointv2OwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Endpointv2 contract.
type Endpointv2OwnershipTransferredIterator struct {
	Event *Endpointv2OwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *Endpointv2OwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Endpointv2OwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(Endpointv2OwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *Endpointv2OwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Endpointv2OwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Endpointv2OwnershipTransferred represents a OwnershipTransferred event raised by the Endpointv2 contract.
type Endpointv2OwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Endpointv2 *Endpointv2Filterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*Endpointv2OwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Endpointv2.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &Endpointv2OwnershipTransferredIterator{contract: _Endpointv2.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Endpointv2 *Endpointv2Filterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *Endpointv2OwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Endpointv2.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Endpointv2OwnershipTransferred)
				if err := _Endpointv2.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Endpointv2 *Endpointv2Filterer) ParseOwnershipTransferred(log types.Log) (*Endpointv2OwnershipTransferred, error) {
	event := new(Endpointv2OwnershipTransferred)
	if err := _Endpointv2.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Endpointv2PacketBurntIterator is returned from FilterPacketBurnt and is used to iterate over the raw logs and unpacked data for PacketBurnt events raised by the Endpointv2 contract.
type Endpointv2PacketBurntIterator struct {
	Event *Endpointv2PacketBurnt // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *Endpointv2PacketBurntIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Endpointv2PacketBurnt)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(Endpointv2PacketBurnt)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *Endpointv2PacketBurntIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Endpointv2PacketBurntIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Endpointv2PacketBurnt represents a PacketBurnt event raised by the Endpointv2 contract.
type Endpointv2PacketBurnt struct {
	SrcEid      uint32
	Sender      [32]byte
	Receiver    common.Address
	Nonce       uint64
	PayloadHash [32]byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterPacketBurnt is a free log retrieval operation binding the contract event 0x7f68a37a6e69a0de35024a234558f9efe4b33b58657753d21eaaa82d51c3510e.
//
// Solidity: event PacketBurnt(uint32 srcEid, bytes32 sender, address receiver, uint64 nonce, bytes32 payloadHash)
func (_Endpointv2 *Endpointv2Filterer) FilterPacketBurnt(opts *bind.FilterOpts) (*Endpointv2PacketBurntIterator, error) {

	logs, sub, err := _Endpointv2.contract.FilterLogs(opts, "PacketBurnt")
	if err != nil {
		return nil, err
	}
	return &Endpointv2PacketBurntIterator{contract: _Endpointv2.contract, event: "PacketBurnt", logs: logs, sub: sub}, nil
}

// WatchPacketBurnt is a free log subscription operation binding the contract event 0x7f68a37a6e69a0de35024a234558f9efe4b33b58657753d21eaaa82d51c3510e.
//
// Solidity: event PacketBurnt(uint32 srcEid, bytes32 sender, address receiver, uint64 nonce, bytes32 payloadHash)
func (_Endpointv2 *Endpointv2Filterer) WatchPacketBurnt(opts *bind.WatchOpts, sink chan<- *Endpointv2PacketBurnt) (event.Subscription, error) {

	logs, sub, err := _Endpointv2.contract.WatchLogs(opts, "PacketBurnt")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Endpointv2PacketBurnt)
				if err := _Endpointv2.contract.UnpackLog(event, "PacketBurnt", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePacketBurnt is a log parse operation binding the contract event 0x7f68a37a6e69a0de35024a234558f9efe4b33b58657753d21eaaa82d51c3510e.
//
// Solidity: event PacketBurnt(uint32 srcEid, bytes32 sender, address receiver, uint64 nonce, bytes32 payloadHash)
func (_Endpointv2 *Endpointv2Filterer) ParsePacketBurnt(log types.Log) (*Endpointv2PacketBurnt, error) {
	event := new(Endpointv2PacketBurnt)
	if err := _Endpointv2.contract.UnpackLog(event, "PacketBurnt", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Endpointv2PacketDeliveredIterator is returned from FilterPacketDelivered and is used to iterate over the raw logs and unpacked data for PacketDelivered events raised by the Endpointv2 contract.
type Endpointv2PacketDeliveredIterator struct {
	Event *Endpointv2PacketDelivered // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *Endpointv2PacketDeliveredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Endpointv2PacketDelivered)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(Endpointv2PacketDelivered)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *Endpointv2PacketDeliveredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Endpointv2PacketDeliveredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Endpointv2PacketDelivered represents a PacketDelivered event raised by the Endpointv2 contract.
type Endpointv2PacketDelivered struct {
	Origin   Origin
	Receiver common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterPacketDelivered is a free log retrieval operation binding the contract event 0x3cd5e48f9730b129dc7550f0fcea9c767b7be37837cd10e55eb35f734f4bca04.
//
// Solidity: event PacketDelivered((uint32,bytes32,uint64) origin, address receiver)
func (_Endpointv2 *Endpointv2Filterer) FilterPacketDelivered(opts *bind.FilterOpts) (*Endpointv2PacketDeliveredIterator, error) {

	logs, sub, err := _Endpointv2.contract.FilterLogs(opts, "PacketDelivered")
	if err != nil {
		return nil, err
	}
	return &Endpointv2PacketDeliveredIterator{contract: _Endpointv2.contract, event: "PacketDelivered", logs: logs, sub: sub}, nil
}

// WatchPacketDelivered is a free log subscription operation binding the contract event 0x3cd5e48f9730b129dc7550f0fcea9c767b7be37837cd10e55eb35f734f4bca04.
//
// Solidity: event PacketDelivered((uint32,bytes32,uint64) origin, address receiver)
func (_Endpointv2 *Endpointv2Filterer) WatchPacketDelivered(opts *bind.WatchOpts, sink chan<- *Endpointv2PacketDelivered) (event.Subscription, error) {

	logs, sub, err := _Endpointv2.contract.WatchLogs(opts, "PacketDelivered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Endpointv2PacketDelivered)
				if err := _Endpointv2.contract.UnpackLog(event, "PacketDelivered", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePacketDelivered is a log parse operation binding the contract event 0x3cd5e48f9730b129dc7550f0fcea9c767b7be37837cd10e55eb35f734f4bca04.
//
// Solidity: event PacketDelivered((uint32,bytes32,uint64) origin, address receiver)
func (_Endpointv2 *Endpointv2Filterer) ParsePacketDelivered(log types.Log) (*Endpointv2PacketDelivered, error) {
	event := new(Endpointv2PacketDelivered)
	if err := _Endpointv2.contract.UnpackLog(event, "PacketDelivered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Endpointv2PacketNilifiedIterator is returned from FilterPacketNilified and is used to iterate over the raw logs and unpacked data for PacketNilified events raised by the Endpointv2 contract.
type Endpointv2PacketNilifiedIterator struct {
	Event *Endpointv2PacketNilified // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *Endpointv2PacketNilifiedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Endpointv2PacketNilified)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(Endpointv2PacketNilified)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *Endpointv2PacketNilifiedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Endpointv2PacketNilifiedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Endpointv2PacketNilified represents a PacketNilified event raised by the Endpointv2 contract.
type Endpointv2PacketNilified struct {
	SrcEid      uint32
	Sender      [32]byte
	Receiver    common.Address
	Nonce       uint64
	PayloadHash [32]byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterPacketNilified is a free log retrieval operation binding the contract event 0xaf0450c392c4f702515a457a362328c8aa21916048ca6d0419e248b30cb55292.
//
// Solidity: event PacketNilified(uint32 srcEid, bytes32 sender, address receiver, uint64 nonce, bytes32 payloadHash)
func (_Endpointv2 *Endpointv2Filterer) FilterPacketNilified(opts *bind.FilterOpts) (*Endpointv2PacketNilifiedIterator, error) {

	logs, sub, err := _Endpointv2.contract.FilterLogs(opts, "PacketNilified")
	if err != nil {
		return nil, err
	}
	return &Endpointv2PacketNilifiedIterator{contract: _Endpointv2.contract, event: "PacketNilified", logs: logs, sub: sub}, nil
}

// WatchPacketNilified is a free log subscription operation binding the contract event 0xaf0450c392c4f702515a457a362328c8aa21916048ca6d0419e248b30cb55292.
//
// Solidity: event PacketNilified(uint32 srcEid, bytes32 sender, address receiver, uint64 nonce, bytes32 payloadHash)
func (_Endpointv2 *Endpointv2Filterer) WatchPacketNilified(opts *bind.WatchOpts, sink chan<- *Endpointv2PacketNilified) (event.Subscription, error) {

	logs, sub, err := _Endpointv2.contract.WatchLogs(opts, "PacketNilified")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Endpointv2PacketNilified)
				if err := _Endpointv2.contract.UnpackLog(event, "PacketNilified", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePacketNilified is a log parse operation binding the contract event 0xaf0450c392c4f702515a457a362328c8aa21916048ca6d0419e248b30cb55292.
//
// Solidity: event PacketNilified(uint32 srcEid, bytes32 sender, address receiver, uint64 nonce, bytes32 payloadHash)
func (_Endpointv2 *Endpointv2Filterer) ParsePacketNilified(log types.Log) (*Endpointv2PacketNilified, error) {
	event := new(Endpointv2PacketNilified)
	if err := _Endpointv2.contract.UnpackLog(event, "PacketNilified", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Endpointv2PacketSentIterator is returned from FilterPacketSent and is used to iterate over the raw logs and unpacked data for PacketSent events raised by the Endpointv2 contract.
type Endpointv2PacketSentIterator struct {
	Event *Endpointv2PacketSent // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *Endpointv2PacketSentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Endpointv2PacketSent)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(Endpointv2PacketSent)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *Endpointv2PacketSentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Endpointv2PacketSentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Endpointv2PacketSent represents a PacketSent event raised by the Endpointv2 contract.
type Endpointv2PacketSent struct {
	EncodedPayload []byte
	Options        []byte
	SendLibrary    common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterPacketSent is a free log retrieval operation binding the contract event 0x1ab700d4ced0c005b164c0f789fd09fcbb0156d4c2041b8a3bfbcd961cd1567f.
//
// Solidity: event PacketSent(bytes encodedPayload, bytes options, address sendLibrary)
func (_Endpointv2 *Endpointv2Filterer) FilterPacketSent(opts *bind.FilterOpts) (*Endpointv2PacketSentIterator, error) {

	logs, sub, err := _Endpointv2.contract.FilterLogs(opts, "PacketSent")
	if err != nil {
		return nil, err
	}
	return &Endpointv2PacketSentIterator{contract: _Endpointv2.contract, event: "PacketSent", logs: logs, sub: sub}, nil
}

// WatchPacketSent is a free log subscription operation binding the contract event 0x1ab700d4ced0c005b164c0f789fd09fcbb0156d4c2041b8a3bfbcd961cd1567f.
//
// Solidity: event PacketSent(bytes encodedPayload, bytes options, address sendLibrary)
func (_Endpointv2 *Endpointv2Filterer) WatchPacketSent(opts *bind.WatchOpts, sink chan<- *Endpointv2PacketSent) (event.Subscription, error) {

	logs, sub, err := _Endpointv2.contract.WatchLogs(opts, "PacketSent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Endpointv2PacketSent)
				if err := _Endpointv2.contract.UnpackLog(event, "PacketSent", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePacketSent is a log parse operation binding the contract event 0x1ab700d4ced0c005b164c0f789fd09fcbb0156d4c2041b8a3bfbcd961cd1567f.
//
// Solidity: event PacketSent(bytes encodedPayload, bytes options, address sendLibrary)
func (_Endpointv2 *Endpointv2Filterer) ParsePacketSent(log types.Log) (*Endpointv2PacketSent, error) {
	event := new(Endpointv2PacketSent)
	if err := _Endpointv2.contract.UnpackLog(event, "PacketSent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Endpointv2PacketVerifiedIterator is returned from FilterPacketVerified and is used to iterate over the raw logs and unpacked data for PacketVerified events raised by the Endpointv2 contract.
type Endpointv2PacketVerifiedIterator struct {
	Event *Endpointv2PacketVerified // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *Endpointv2PacketVerifiedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Endpointv2PacketVerified)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(Endpointv2PacketVerified)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *Endpointv2PacketVerifiedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Endpointv2PacketVerifiedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Endpointv2PacketVerified represents a PacketVerified event raised by the Endpointv2 contract.
type Endpointv2PacketVerified struct {
	Origin      Origin
	Receiver    common.Address
	PayloadHash [32]byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterPacketVerified is a free log retrieval operation binding the contract event 0x0d87345f3d1c929caba93e1c3821b54ff3512e12b66aa3cfe54b6bcbc17e59b4.
//
// Solidity: event PacketVerified((uint32,bytes32,uint64) origin, address receiver, bytes32 payloadHash)
func (_Endpointv2 *Endpointv2Filterer) FilterPacketVerified(opts *bind.FilterOpts) (*Endpointv2PacketVerifiedIterator, error) {

	logs, sub, err := _Endpointv2.contract.FilterLogs(opts, "PacketVerified")
	if err != nil {
		return nil, err
	}
	return &Endpointv2PacketVerifiedIterator{contract: _Endpointv2.contract, event: "PacketVerified", logs: logs, sub: sub}, nil
}

// WatchPacketVerified is a free log subscription operation binding the contract event 0x0d87345f3d1c929caba93e1c3821b54ff3512e12b66aa3cfe54b6bcbc17e59b4.
//
// Solidity: event PacketVerified((uint32,bytes32,uint64) origin, address receiver, bytes32 payloadHash)
func (_Endpointv2 *Endpointv2Filterer) WatchPacketVerified(opts *bind.WatchOpts, sink chan<- *Endpointv2PacketVerified) (event.Subscription, error) {

	logs, sub, err := _Endpointv2.contract.WatchLogs(opts, "PacketVerified")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Endpointv2PacketVerified)
				if err := _Endpointv2.contract.UnpackLog(event, "PacketVerified", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePacketVerified is a log parse operation binding the contract event 0x0d87345f3d1c929caba93e1c3821b54ff3512e12b66aa3cfe54b6bcbc17e59b4.
//
// Solidity: event PacketVerified((uint32,bytes32,uint64) origin, address receiver, bytes32 payloadHash)
func (_Endpointv2 *Endpointv2Filterer) ParsePacketVerified(log types.Log) (*Endpointv2PacketVerified, error) {
	event := new(Endpointv2PacketVerified)
	if err := _Endpointv2.contract.UnpackLog(event, "PacketVerified", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Endpointv2PausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the Endpointv2 contract.
type Endpointv2PausedIterator struct {
	Event *Endpointv2Paused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *Endpointv2PausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Endpointv2Paused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(Endpointv2Paused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *Endpointv2PausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Endpointv2PausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Endpointv2Paused represents a Paused event raised by the Endpointv2 contract.
type Endpointv2Paused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Endpointv2 *Endpointv2Filterer) FilterPaused(opts *bind.FilterOpts) (*Endpointv2PausedIterator, error) {

	logs, sub, err := _Endpointv2.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &Endpointv2PausedIterator{contract: _Endpointv2.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Endpointv2 *Endpointv2Filterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *Endpointv2Paused) (event.Subscription, error) {

	logs, sub, err := _Endpointv2.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Endpointv2Paused)
				if err := _Endpointv2.contract.UnpackLog(event, "Paused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Endpointv2 *Endpointv2Filterer) ParsePaused(log types.Log) (*Endpointv2Paused, error) {
	event := new(Endpointv2Paused)
	if err := _Endpointv2.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Endpointv2ReceiveLibrarySetIterator is returned from FilterReceiveLibrarySet and is used to iterate over the raw logs and unpacked data for ReceiveLibrarySet events raised by the Endpointv2 contract.
type Endpointv2ReceiveLibrarySetIterator struct {
	Event *Endpointv2ReceiveLibrarySet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *Endpointv2ReceiveLibrarySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Endpointv2ReceiveLibrarySet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(Endpointv2ReceiveLibrarySet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *Endpointv2ReceiveLibrarySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Endpointv2ReceiveLibrarySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Endpointv2ReceiveLibrarySet represents a ReceiveLibrarySet event raised by the Endpointv2 contract.
type Endpointv2ReceiveLibrarySet struct {
	Receiver common.Address
	Eid      uint32
	NewLib   common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterReceiveLibrarySet is a free log retrieval operation binding the contract event 0xcd6f92f5ac6185a5acfa02c92090746cec64d777269cbcd0ed031e396657a1c2.
//
// Solidity: event ReceiveLibrarySet(address receiver, uint32 eid, address newLib)
func (_Endpointv2 *Endpointv2Filterer) FilterReceiveLibrarySet(opts *bind.FilterOpts) (*Endpointv2ReceiveLibrarySetIterator, error) {

	logs, sub, err := _Endpointv2.contract.FilterLogs(opts, "ReceiveLibrarySet")
	if err != nil {
		return nil, err
	}
	return &Endpointv2ReceiveLibrarySetIterator{contract: _Endpointv2.contract, event: "ReceiveLibrarySet", logs: logs, sub: sub}, nil
}

// WatchReceiveLibrarySet is a free log subscription operation binding the contract event 0xcd6f92f5ac6185a5acfa02c92090746cec64d777269cbcd0ed031e396657a1c2.
//
// Solidity: event ReceiveLibrarySet(address receiver, uint32 eid, address newLib)
func (_Endpointv2 *Endpointv2Filterer) WatchReceiveLibrarySet(opts *bind.WatchOpts, sink chan<- *Endpointv2ReceiveLibrarySet) (event.Subscription, error) {

	logs, sub, err := _Endpointv2.contract.WatchLogs(opts, "ReceiveLibrarySet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Endpointv2ReceiveLibrarySet)
				if err := _Endpointv2.contract.UnpackLog(event, "ReceiveLibrarySet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseReceiveLibrarySet is a log parse operation binding the contract event 0xcd6f92f5ac6185a5acfa02c92090746cec64d777269cbcd0ed031e396657a1c2.
//
// Solidity: event ReceiveLibrarySet(address receiver, uint32 eid, address newLib)
func (_Endpointv2 *Endpointv2Filterer) ParseReceiveLibrarySet(log types.Log) (*Endpointv2ReceiveLibrarySet, error) {
	event := new(Endpointv2ReceiveLibrarySet)
	if err := _Endpointv2.contract.UnpackLog(event, "ReceiveLibrarySet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Endpointv2ReceiveLibraryTimeoutSetIterator is returned from FilterReceiveLibraryTimeoutSet and is used to iterate over the raw logs and unpacked data for ReceiveLibraryTimeoutSet events raised by the Endpointv2 contract.
type Endpointv2ReceiveLibraryTimeoutSetIterator struct {
	Event *Endpointv2ReceiveLibraryTimeoutSet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *Endpointv2ReceiveLibraryTimeoutSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Endpointv2ReceiveLibraryTimeoutSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(Endpointv2ReceiveLibraryTimeoutSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *Endpointv2ReceiveLibraryTimeoutSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Endpointv2ReceiveLibraryTimeoutSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Endpointv2ReceiveLibraryTimeoutSet represents a ReceiveLibraryTimeoutSet event raised by the Endpointv2 contract.
type Endpointv2ReceiveLibraryTimeoutSet struct {
	Receiver common.Address
	Eid      uint32
	OldLib   common.Address
	Timeout  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterReceiveLibraryTimeoutSet is a free log retrieval operation binding the contract event 0x4e0a5bbfa0c11a64effb1ada324b5437a17272e1aed9320398715ef71bb20928.
//
// Solidity: event ReceiveLibraryTimeoutSet(address receiver, uint32 eid, address oldLib, uint256 timeout)
func (_Endpointv2 *Endpointv2Filterer) FilterReceiveLibraryTimeoutSet(opts *bind.FilterOpts) (*Endpointv2ReceiveLibraryTimeoutSetIterator, error) {

	logs, sub, err := _Endpointv2.contract.FilterLogs(opts, "ReceiveLibraryTimeoutSet")
	if err != nil {
		return nil, err
	}
	return &Endpointv2ReceiveLibraryTimeoutSetIterator{contract: _Endpointv2.contract, event: "ReceiveLibraryTimeoutSet", logs: logs, sub: sub}, nil
}

// WatchReceiveLibraryTimeoutSet is a free log subscription operation binding the contract event 0x4e0a5bbfa0c11a64effb1ada324b5437a17272e1aed9320398715ef71bb20928.
//
// Solidity: event ReceiveLibraryTimeoutSet(address receiver, uint32 eid, address oldLib, uint256 timeout)
func (_Endpointv2 *Endpointv2Filterer) WatchReceiveLibraryTimeoutSet(opts *bind.WatchOpts, sink chan<- *Endpointv2ReceiveLibraryTimeoutSet) (event.Subscription, error) {

	logs, sub, err := _Endpointv2.contract.WatchLogs(opts, "ReceiveLibraryTimeoutSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Endpointv2ReceiveLibraryTimeoutSet)
				if err := _Endpointv2.contract.UnpackLog(event, "ReceiveLibraryTimeoutSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseReceiveLibraryTimeoutSet is a log parse operation binding the contract event 0x4e0a5bbfa0c11a64effb1ada324b5437a17272e1aed9320398715ef71bb20928.
//
// Solidity: event ReceiveLibraryTimeoutSet(address receiver, uint32 eid, address oldLib, uint256 timeout)
func (_Endpointv2 *Endpointv2Filterer) ParseReceiveLibraryTimeoutSet(log types.Log) (*Endpointv2ReceiveLibraryTimeoutSet, error) {
	event := new(Endpointv2ReceiveLibraryTimeoutSet)
	if err := _Endpointv2.contract.UnpackLog(event, "ReceiveLibraryTimeoutSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Endpointv2SendLibrarySetIterator is returned from FilterSendLibrarySet and is used to iterate over the raw logs and unpacked data for SendLibrarySet events raised by the Endpointv2 contract.
type Endpointv2SendLibrarySetIterator struct {
	Event *Endpointv2SendLibrarySet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *Endpointv2SendLibrarySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Endpointv2SendLibrarySet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(Endpointv2SendLibrarySet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *Endpointv2SendLibrarySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Endpointv2SendLibrarySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Endpointv2SendLibrarySet represents a SendLibrarySet event raised by the Endpointv2 contract.
type Endpointv2SendLibrarySet struct {
	Sender common.Address
	Eid    uint32
	NewLib common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSendLibrarySet is a free log retrieval operation binding the contract event 0x4cff966ebee29a156dcb34cf72c1d06231fb1777f6bdf6e8089819232f002b1c.
//
// Solidity: event SendLibrarySet(address sender, uint32 eid, address newLib)
func (_Endpointv2 *Endpointv2Filterer) FilterSendLibrarySet(opts *bind.FilterOpts) (*Endpointv2SendLibrarySetIterator, error) {

	logs, sub, err := _Endpointv2.contract.FilterLogs(opts, "SendLibrarySet")
	if err != nil {
		return nil, err
	}
	return &Endpointv2SendLibrarySetIterator{contract: _Endpointv2.contract, event: "SendLibrarySet", logs: logs, sub: sub}, nil
}

// WatchSendLibrarySet is a free log subscription operation binding the contract event 0x4cff966ebee29a156dcb34cf72c1d06231fb1777f6bdf6e8089819232f002b1c.
//
// Solidity: event SendLibrarySet(address sender, uint32 eid, address newLib)
func (_Endpointv2 *Endpointv2Filterer) WatchSendLibrarySet(opts *bind.WatchOpts, sink chan<- *Endpointv2SendLibrarySet) (event.Subscription, error) {

	logs, sub, err := _Endpointv2.contract.WatchLogs(opts, "SendLibrarySet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Endpointv2SendLibrarySet)
				if err := _Endpointv2.contract.UnpackLog(event, "SendLibrarySet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSendLibrarySet is a log parse operation binding the contract event 0x4cff966ebee29a156dcb34cf72c1d06231fb1777f6bdf6e8089819232f002b1c.
//
// Solidity: event SendLibrarySet(address sender, uint32 eid, address newLib)
func (_Endpointv2 *Endpointv2Filterer) ParseSendLibrarySet(log types.Log) (*Endpointv2SendLibrarySet, error) {
	event := new(Endpointv2SendLibrarySet)
	if err := _Endpointv2.contract.UnpackLog(event, "SendLibrarySet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Endpointv2UnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the Endpointv2 contract.
type Endpointv2UnpausedIterator struct {
	Event *Endpointv2Unpaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *Endpointv2UnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Endpointv2Unpaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(Endpointv2Unpaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *Endpointv2UnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Endpointv2UnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Endpointv2Unpaused represents a Unpaused event raised by the Endpointv2 contract.
type Endpointv2Unpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Endpointv2 *Endpointv2Filterer) FilterUnpaused(opts *bind.FilterOpts) (*Endpointv2UnpausedIterator, error) {

	logs, sub, err := _Endpointv2.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &Endpointv2UnpausedIterator{contract: _Endpointv2.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Endpointv2 *Endpointv2Filterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *Endpointv2Unpaused) (event.Subscription, error) {

	logs, sub, err := _Endpointv2.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Endpointv2Unpaused)
				if err := _Endpointv2.contract.UnpackLog(event, "Unpaused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Endpointv2 *Endpointv2Filterer) ParseUnpaused(log types.Log) (*Endpointv2Unpaused, error) {
	event := new(Endpointv2Unpaused)
	if err := _Endpointv2.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
