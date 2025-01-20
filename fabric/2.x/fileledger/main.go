package main

import (
	"encoding/json"
	"fileledger/kafka"
	"fileledger/parse"
	"fileledger/tools"
	"fmt"
	"github.com/golang/protobuf/proto"
	cb "github.com/hyperledger/fabric-protos-go/common"
	"github.com/hyperledger/fabric-protos-go/orderer"
	"github.com/hyperledger/fabric/common/flogging"
	"github.com/hyperledger/fabric/common/ledger/blockledger/fileledger"
	"github.com/hyperledger/fabric/common/metrics/disabled"
	"github.com/hyperledger/fabric/core/ledger/kvledger"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"io"
	"io/ioutil"
	"os"
	"path"
	"strconv"
)

/**
 * @Description $
 * @Author hll
 * @Date 2024/9/3 下午7:25
 **/
//var ldPath = flag.String("ledgerPath", "E:\\go-new-work\\newstudy\\fabric\\2.x\\fileledger\\orderer", "orderer账本路径")
//
//var channelName = flag.String("channel", "hll9", "将删除的应用通道")
//var blockNumber = flag.Uint64("block", 0, "快高")

func init() {
	flogging.Init(flogging.Config{
		Format:  "",
		Writer:  os.Stderr,
		LogSpec: "DEBUG",
	})
}

var rootCmd = &cobra.Command{
	Use:   "hll",
	Short: "修改order/peer账本数据",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("rootCmd ==>Run")
	},
}
var orderCmd = &cobra.Command{Use: "order"}
var peerCmd = &cobra.Command{Use: "peer"}

var orderExitCmd = &cobra.Command{
	Use:   "exit",
	Short: "order exit 裁剪",
	// args 表示命令行中没有被解析为标志或参数的剩余参数部分。这些剩余的参数通常是用户在命令行中输入的非标志和非参数的额外文本。
	// go run .\main.go hello -n tt yy uu
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("order exit 裁剪，%s！\n", name)
		ExitOrderChannel()
	},
}
var orderQueryCmd = &cobra.Command{
	Use:   "query",
	Short: "query 指定快高",
	// args 表示命令行中没有被解析为标志或参数的剩余参数部分。这些剩余的参数通常是用户在命令行中输入的非标志和非参数的额外文本。
	// go run .\main.go hello -n tt yy uu
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("order query，%s！\n", name)
		QueryOrderChannel()
	},
}
var orderChannelQueryCmd = &cobra.Command{
	Use:   "queryAll",
	Short: "query order所有通道",
	// args 表示命令行中没有被解析为标志或参数的剩余参数部分。这些剩余的参数通常是用户在命令行中输入的非标志和非参数的额外文本。
	// go run .\main.go hello -n tt yy uu
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("query order所有通道%s！\n", name)
		orderChannelQueryAll()
	},
}

var orderRestCmd = &cobra.Command{
	Use:   "rollback",
	Short: "rollback 裁剪到指定快高",
	// args 表示命令行中没有被解析为标志或参数的剩余参数部分。这些剩余的参数通常是用户在命令行中输入的非标志和非参数的额外文本。
	// go run .\main.go hello -n tt yy uu
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("order rollback裁剪，%s！\n", name)
		DeleteOrderLedgerByNum()
	},
}

var orderSaveBlockCmd = &cobra.Command{
	Use:   "save",
	Short: "保存指定快高成文件",
	// args 表示命令行中没有被解析为标志或参数的剩余参数部分。这些剩余的参数通常是用户在命令行中输入的非标志和非参数的额外文本。
	// go run .\main.go hello -n tt yy uu
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("order save 保存指定快高成文件，%s！\n", name)
		SaveOrderLedgerByBlockNum()
	},
}
var orderAddBlockCmd = &cobra.Command{
	Use:   "add",
	Short: "通过块文件为order添加账本",
	// args 表示命令行中没有被解析为标志或参数的剩余参数部分。这些剩余的参数通常是用户在命令行中输入的非标志和非参数的额外文本。
	// go run .\main.go hello -n tt yy uu
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("order add 根据指定快高成文件添加账本，%s！\n", name)
		AddOrderLedger()
	},
}
var peerRestCmd = &cobra.Command{
	Use:   "rollback",
	Short: "rollback 裁剪",
	// args 表示命令行中没有被解析为标志或参数的剩余参数部分。这些剩余的参数通常是用户在命令行中输入的非标志和非参数的额外文本。
	// go run .\main.go hello -n tt yy uu
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("peer 裁剪，%s！\n", name)
		DeletePeerLedger()
	},
}

var name string
var fileCode string
var ldPath string //flag.String("ledgerPath", "E:\\go-new-work\\newstudy\\fabric\\2.x\\fileledger\\orderer", "orderer账本路径")

var channelName string // = flag.String("channel", "hll9", "将删除的应用通道")
var blockNumber uint64 //= flag.Uint64("block", 0, "快高")

var orderName string // = flag.String("channel", "hll9", "将删除的应用通道")
var needParse bool   //= flag.Bool("")
func main() {
	rootCmd.PersistentFlags().StringVarP(&ldPath, "ledgerPath", "l", "order1/sharedfiles/orderer/orderer 或者 org/peer1/sharedfiles/production/ledgersData/", "包含chains和index的上一级账本路径")
	rootCmd.PersistentFlags().StringVarP(&channelName, "channel", "c", "hll9", "将删除的应用通道")
	rootCmd.PersistentFlags().Uint64VarP(&blockNumber, "blockNumber", "n", 0, "快高")
	orderChannelQueryCmd.Flags().StringVarP(&orderName, "orderName", "o", "order1", "order名称")
	orderQueryCmd.Flags().BoolVarP(&needParse, "needParse", "p", false, "是否需要解析块")

	rootCmd.AddCommand(orderCmd)
	rootCmd.AddCommand(peerCmd)
	rootCmd.AddCommand(kafka.KfkCmd)
	orderCmd.AddCommand(orderExitCmd)
	orderCmd.AddCommand(orderRestCmd)
	orderCmd.AddCommand(orderQueryCmd)
	orderCmd.AddCommand(orderChannelQueryCmd)
	orderCmd.AddCommand(orderSaveBlockCmd)
	orderCmd.AddCommand(orderAddBlockCmd)

	peerCmd.AddCommand(peerRestCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

//func main() {
//	flag.Parse()
//	//ld := "E:\\go-new-work\\newstudy\\fabric\\2.x\\fileledger\\orderer"
//	channel := *channelName
//	ld := *ldPath
//	lf, err := fileledger.New(ld, &disabled.Provider{})
//	if err != nil {
//		panic(err)
//	}
//	for _, s := range lf.ChannelIDs() {
//		fmt.Println(s)
//	}
//	c, err := lf.GetOrCreate(channel)
//	fmt.Printf("%s Height==>%d \n", channel, c.Height())
//	//it, seq := c.Iterator(&orderer.SeekPosition{Type: &orderer.SeekPosition_Newest{}})
//	it, seq := c.Iterator(&orderer.SeekPosition{Type: &orderer.SeekPosition_Specified{Specified: &orderer.SeekSpecified{Number: 1}}})
//
//	committedBlock, status := it.Next()
//	fmt.Println(seq)
//	fmt.Println(status)
//	parse.PreprocessProtoBlock(committedBlock, channel)
//	it.Close()
//	err = lf.Remove(channel)
//	if err != nil {
//		panic(err)
//	}
//
//}

func ExitOrderChannel() {
	fmt.Printf("channelName: %s \n", channelName)
	fmt.Printf("ldPath: %s \n", ldPath)

	channel := channelName
	ld := ldPath
	lf, err := fileledger.New(ld, &disabled.Provider{})
	if err != nil {
		panic(err)
	}
	//for _, s := range lf.ChannelIDs() {
	//	fmt.Println(s)
	//}
	//c, err := lf.GetOrCreate(channel)
	//fmt.Printf("%s Height==>%d \n", channel, c.Height())
	////it, seq := c.Iterator(&orderer.SeekPosition{Type: &orderer.SeekPosition_Newest{}})
	//it, seq := c.Iterator(&orderer.SeekPosition{Type: &orderer.SeekPosition_Specified{Specified: &orderer.SeekSpecified{Number: blockNumber}}})
	//
	//committedBlock, status := it.Next()
	//fmt.Println(seq)
	//fmt.Println(status)
	//parse.PreprocessProtoBlock(committedBlock, channel)
	//it.Close()
	err = lf.Remove(channel)
	if err != nil {
		panic(err)
	}
}

func orderChannelQueryAll() {
	tools.Init()
	channelInfo := tools.GetChannelInfo()
	fmt.Printf("ldPath: %s \n", ldPath)

	ld := ldPath
	lf, err := fileledger.New(ld, &disabled.Provider{})
	if err != nil {
		panic(err)
	}
	for _, s := range lf.ChannelIDs() {
		c, err := lf.GetOrCreate(s)
		if err != nil {
			panic(fmt.Errorf("GetOrCreate:%s", err))
		}
		maxH := c.Height()
		fmt.Printf("%s最高快高 Height==>%d \n", s, maxH)
		it, seq := c.Iterator(&orderer.SeekPosition{Type: &orderer.SeekPosition_Newest{}})
		//it, seq := c.Iterator(&orderer.SeekPosition{Type: &orderer.SeekPosition_Specified{Specified: &orderer.SeekSpecified{Number: blockNumber}}})

		blk, status := it.Next()
		fmt.Println(seq)
		fmt.Println(status)
		//fmt.Println("blk.Header.Number", blk.Header.Number)
		//fmt.Printf("blk.Header.PreviousHash:[%x]\n", blk.Header.PreviousHash)
		//parse.PreprocessProtoBlock(committedBlock, s)
		it.Close()

		var ok bool
		var blockInfo *tools.BlockInfo
		var block *tools.Block
		if blockInfo, ok = channelInfo.Channel[s]; !ok {
			blockInfo = &tools.BlockInfo{
				Order: make(map[string]*tools.Block),
			}
		}
		channelInfo.Channel[s] = blockInfo
		if block, ok = blockInfo.Order[orderName]; !ok {
			block = &tools.Block{
				Number:       blk.Header.Number,
				PreviousHash: fmt.Sprintf("%x", blk.Header.PreviousHash),
				DataHash:     fmt.Sprintf("%x", blk.Header.DataHash),
			}
		}
		blockInfo.Order[orderName] = block
	}
	data, err := json.Marshal(channelInfo)
	if err != nil {
		panic(fmt.Sprintf("error marshaling metrics data: %v", err))
	}
	f, err := os.OpenFile("metrics.json", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		panic(err)
	}
	_, err = f.Write(data)
	if err1 := f.Close(); err1 != nil && err == nil {
		err = err1
	}
}

func QueryOrderChannel() {
	fmt.Printf("channelName: %s \n", channelName)
	fmt.Printf("ldPath: %s \n", ldPath)
	fmt.Printf("needParse: %t \n", needParse)

	channel := channelName
	ld := ldPath
	lf, err := fileledger.New(ld, &disabled.Provider{})
	if err != nil {
		panic(err)
	}
	//for _, s := range lf.ChannelIDs() {
	//	fmt.Println(s)
	//}
	c, err := lf.GetOrCreate(channel)
	maxH := c.Height()
	fmt.Printf("%s最高快高 Height==>%d \n", channel, maxH)
	if maxH > 0 && maxH <= blockNumber {
		panic(fmt.Errorf("不被允许的快高查询"))
		return
	}
	//it, seq := c.Iterator(&orderer.SeekPosition{Type: &orderer.SeekPosition_Newest{}})
	it, seq := c.Iterator(&orderer.SeekPosition{Type: &orderer.SeekPosition_Specified{Specified: &orderer.SeekSpecified{Number: blockNumber}}})
	defer it.Close()

	committedBlock, status := it.Next()
	fmt.Println(seq)
	fmt.Println(status)
	parse.PreprocessProtoBlock(committedBlock, channel, needParse)

}

func SaveOrderLedgerByBlockNum() {
	fmt.Printf("channelName: %s \n", channelName)
	fmt.Printf("ldPath: %s \n", ldPath)

	channel := channelName
	ld := ldPath
	lf, err := fileledger.New(ld, &disabled.Provider{})
	if err != nil {
		panic(err)
	}
	for _, s := range lf.ChannelIDs() {
		fmt.Println(s)
	}
	c, err := lf.GetOrCreate(channel)
	maxH := c.Height()
	fmt.Printf("%s最高快高 Height==>%d \n", channel, maxH)
	if maxH > 0 && maxH <= blockNumber {
		panic(fmt.Errorf("不被允许的快高查询"))
		return
	}
	//it, seq := c.Iterator(&orderer.SeekPosition{Type: &orderer.SeekPosition_Newest{}})
	it, seq := c.Iterator(&orderer.SeekPosition{Type: &orderer.SeekPosition_Specified{Specified: &orderer.SeekSpecified{Number: blockNumber}}})

	committedBlock, status := it.Next()
	fmt.Println(seq)
	fmt.Println(status)
	it.Close()
	b, err := proto.Marshal(committedBlock)
	if err != nil {
		panic(fmt.Errorf("序列化区块失败：%s", err))
	}
	Number := strconv.FormatUint(blockNumber, 10)
	var file = channel + "_" + Number + ".block"

	if err = ioutil.WriteFile(file, b, 0644); err != nil {
		panic(fmt.Errorf("写入文件失败：%s", err))
	}
}

func AddOrderLedger() {
	fmt.Printf("channelName: %s \n", channelName)
	fmt.Printf("ldPath: %s \n", ldPath)

	channel := channelName
	ld := ldPath
	lf, err := fileledger.New(ld, &disabled.Provider{})
	if err != nil {
		panic(err)
	}
	c, err := lf.GetOrCreate(channel)
	maxH := c.Height()
	fmt.Printf("%s当前快高 Height==>%d \n", channel, maxH)

	//msgType := proto.MessageType("common.Block")
	//if msgType == nil {
	//	panic(fmt.Errorf("message of type %s unknown", msgType))
	//}
	//msg := reflect.New(msgType.Elem()).Interface().(proto.Message)
	Number := strconv.FormatUint(blockNumber, 10)
	var filePath = channel + "_" + Number + ".block"
	fmt.Printf("add block file : %s \n", filePath)
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("打开文件出错:", err)
		return
	}
	defer file.Close()
	in, err := io.ReadAll(file)
	if err != nil {
		panic(fmt.Errorf("error reading input:%s", err))
	}
	var block = &cb.Block{}
	err = proto.Unmarshal(in, block)
	if err != nil {
		panic(fmt.Errorf("error unmarshaling:%s", err))
	}
	err = c.Append(block)
	if err != nil {
		fmt.Println("添加order账本失败:", err)
		return
	}
}

func DeletePeerLedger() {
	//rootFSPath := filepath.Join(coreconfig.GetPath("peer.fileSystemPath"), "ledgersData")
	//rootFSPath := "E:\\go-new-work\\newstudy\\fabric\\2.x\\fileledger\\ledgersData"
	//ldPath = rootFSPath
	err := kvledger.RollbackKVLedger(ldPath, channelName, blockNumber)
	if err != nil {
		panic(err)
	}
}

func DeleteOrderLedgerByNum() {
	//rootFSPath := filepath.Join(coreconfig.GetPath("peer.fileSystemPath"), "ledgersData")
	rootFSPath := "E:\\go-new-work\\newstudy\\fabric\\2.x\\fileledger\\orderer3\\orderer"
	ldPath = rootFSPath
	err := kvledger.RollbackKVLedgerHll(ldPath, channelName, blockNumber)
	if err != nil {
		panic(err)
	}
}

func WriteFile(file string, buf string, perm os.FileMode) error {
	dir := path.Dir(file)
	// Create the directory if it doesn't exist
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0755)
		if err != nil {
			return errors.Wrapf(err, "Failed to create directory '%s' for file '%s'", dir, file)
		}
	}

	f, err := os.OpenFile(file, os.O_WRONLY|os.O_CREATE|os.O_APPEND, perm)
	if err != nil {
		return err
	}
	_, err = f.WriteString(buf + "\n")
	if err1 := f.Close(); err1 != nil && err == nil {
		err = err1
	}
	return err
}
