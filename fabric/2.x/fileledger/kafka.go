package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/golang/protobuf/proto"
	cb "github.com/hyperledger/fabric-protos-go/common"
	ab "github.com/hyperledger/fabric-protos-go/orderer"
	"github.com/hyperledger/fabric/protoutil"
	"log"
	"os"
	"time"
)

var addr = flag.String("addr", "localhost:9092", "kafka 地址")
var topic = flag.String("topic", "", "通道名称")
var offset = flag.Int64("offset", 0, "偏移量")
var partition = flag.Int("partition", 0, "分区")

func main() {
	flag.Parse()
	log.Println("addr=", *addr)
	log.Println("topic=", *topic)
	log.Println("offset=", *offset)
	log.Println("partition=", *partition)
	config := sarama.NewConfig()
	config.Consumer.Offsets.Initial = sarama.OffsetOldest

	// 创建消费者
	client, err := sarama.NewClient([]string{*addr}, config)
	if err != nil {
		log.Fatalf("Error creating Kafka client: %v", err)
	}
	defer client.Close()

	consumer, err := sarama.NewConsumerFromClient(client)
	if err != nil {
		log.Fatalf("Error creating Kafka consumer: %v", err)
	}
	defer consumer.Close()

	topic := *topic
	partition := int32(*partition)
	offset := int64(*offset) // 要读取的指定偏移量

	pc, err := consumer.ConsumePartition(topic, partition, offset)
	if err != nil {
		log.Fatalf("Error consuming partition: %v", err)
	}
	defer pc.Close()

	// 设置超时，防止无限等待
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Second)
	defer cancel()
	msg := new(ab.KafkaMessage)

	// 读取指定偏移量的消息
	for {
		select {
		case in := <-pc.Messages():
			log.Println("")
			log.Printf("==============分割===============Offset:%d", in.Offset)
			//if in.Offset == offset {
			//	log.Printf("Received message at offset %d", in.Offset)
			//	return
			//}
			if err := proto.Unmarshal(in.Value, msg); err != nil {
				// This shouldn't happen, it should be filtered at ingress
				log.Printf("[channel: %s] Unable to unmarshal consumed message = %s", topic, err)
				continue
			}
			//else {
			//	log.Printf("[channel: %s] Successfully unmarshalled consumed message, offset is %d. Inspecting type...", topic, in.Offset)
			//}
			switch msg.Type.(type) {
			case *ab.KafkaMessage_Connect:
				log.Printf("[channel: %s] 【kafka Connect】", topic)
			case *ab.KafkaMessage_TimeToCut:
				log.Printf("[channel: %s] It's a 【time-to-cut message】 for block [%d]", topic, msg.GetTimeToCut().GetBlockNumber())
			case *ab.KafkaMessage_Regular:
				log.Printf("[channel: %s] Processing 【regular Kafka message】 of type %s", topic, msg.GetRegular().Class.String())
				env := &cb.Envelope{}
				if err := proto.Unmarshal(msg.GetRegular().Payload, env); err != nil {
					// This shouldn't happen, it should be filtered at ingress
					fmt.Errorf("failed to unmarshal payload of regular message because = %s", err)
				}

				switch msg.GetRegular().Class {
				case ab.KafkaMessageRegular_UNKNOWN:
					log.Panicf("[channel: %s] Kafka message of type UNKNOWN should have been processed already", topic)

				case ab.KafkaMessageRegular_NORMAL:
					log.Printf("[channel: %s] NORMAL 消息，ConfigSeq：%d", topic, msg.GetRegular().ConfigSeq)
					cHeader, err := protoutil.ChannelHeader(env)
					if err != nil {
						panic(err)
					}
					log.Printf("[channel: %s] TxId:%s", cHeader.ChannelId, cHeader.TxId)

				case ab.KafkaMessageRegular_CONFIG:
					log.Printf("[channel: %s] CONFIG 消息，ConfigSeq：%d", topic, msg.GetRegular().ConfigSeq)

				}

			}
		case <-ctx.Done():
			log.Printf(" exit ")
			os.Exit(0)
		}
	}
}
