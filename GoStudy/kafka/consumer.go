/**
 * @Author: Gong Yanhui
 * @Description: kafka 消费者
 * @Date: 2024-02-26 21:29
 */

package kafka

import (
	"fmt"
	"github.com/Shopify/sarama"
	"time"

	cluster "github.com/bsm/sarama-cluster"
)

var (
	kafkaConsumer *cluster.Consumer
	kafkaTopic    = "topic1"
	groupId       = "groupId1"
)

func initConfig(address []string) {
	// 配置
	var err error
	config := cluster.NewConfig()
	config.Consumer.Return.Errors = true
	config.Group.Return.Notifications = true
	config.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategyRange
	config.Consumer.Offsets.Initial = -2
	config.Consumer.Offsets.CommitInterval = 1 * time.Second
	config.Group.Return.Notifications = true
	// 创建消费者
	kafkaConsumer, err = cluster.NewConsumer(address, groupId, []string{kafkaTopic}, config)
	if err != nil {
		panic(err.Error())
	}
	if kafkaConsumer == nil {
		panic(fmt.Sprintf("consumer is nil. kafka info -> {brokers:%v, topic: %v, group: %v}", address, kafkaTopic, groupId))
	}
	fmt.Printf("kafka init success, consumer: %v, topic: %v, ", kafkaConsumer, kafkaTopic)
}

func Consumer(address []string) {
	initConfig(address)
	for {
		select {
		case msg, ok := <-kafkaConsumer.Messages():
			if ok {
				fmt.Printf("receive message: %s \n", msg.Value)
				kafkaConsumer.MarkOffset(msg, "")
			} else {
				fmt.Printf("kafka receive failed")
			}
		case err, ok := <-kafkaConsumer.Errors():
			if ok {
				fmt.Printf("consumer error: %v \n", err)
			}
		case ntf, ok := <-kafkaConsumer.Notifications():
			if ok {
				fmt.Printf("consumer notification: %v \n", ntf)
			}
		}
	}
}
