package tools

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

/**
 * @Description $
 * @Author hll
 * @Date 2024/9/20 下午4:48
 **/

type BlockInfo struct {
	Order map[string]*Block
}

type Block struct {
	Number       uint64 `json:"number,omitempty"`
	PreviousHash string `json:"previous_hash,omitempty"`
	DataHash     string `json:"data_hash,omitempty"`
}

type ChannelInfo struct {
	Channel map[string]*BlockInfo
}

var channelInfo *ChannelInfo

func GetChannelInfo() *ChannelInfo {
	return channelInfo
}
func Init() {
	channelInfo = &ChannelInfo{
		Channel: make(map[string]*BlockInfo),
	}
	filename := "metrics.json"
	// 检查文件是否存在
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		// 文件不存在，创建新文件
		file, err := os.Create(filename)
		if err != nil {
			log.Fatalf("error creating file: %v", err)
		}
		defer file.Close()

	} else {
		// 文件存在，读取文件
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			log.Fatalf("error reading file: %v", err)
		}
		if len(data) > 1 {
			// 解析文件内容到 metrics map
			//var metrics *Metrics
			if err := json.Unmarshal(data, channelInfo); err != nil {
				log.Fatalf("error unmarshaling data: %v", err)
			}

			fmt.Println("当前 Metrics:", channelInfo)
		}
	}
}
