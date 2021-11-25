package main

import (
	"fmt"
	"github.com/Shopify/sarama"
)

// kafka demo
// 基于sarama第三方库开发的kafka client
func main() {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll          // 发送完整数据
	config.Producer.Partitioner = sarama.NewRandomPartitioner // 新选出一个
	config.Producer.Return.Successes = true                   // 成功交付的消息在success——channel返回
	// 构造一个消息
	msg := &sarama.ProducerMessage{}
	msg.Topic = "web_log"
	msg.Value = sarama.StringEncoder("this is a test log")
	client, err := sarama.NewSyncProducer([]string{"127.0.0.1:9092"}, config)
	if err != nil {
		fmt.Println("producer closed, err:", err)
		return
	}
	fmt.Println("连接kafka成功")
	defer func(client sarama.SyncProducer) {
		err := client.Close()
		if err != nil {

		}
	}(client)
	pid, offset, err := client.SendMessage(msg)
	if err != nil {
		fmt.Println("send msg faild, err:", err)
		return
	}
	fmt.Printf("pid:%v offset:%v\n", pid, offset)
	fmt.Println("发送成功！")
}
