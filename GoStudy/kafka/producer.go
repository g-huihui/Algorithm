/**
 * @Author: Gong Yanhui
 * @Description: kafka 生产者
 * @Date: 2024-02-26 20:18
 */

package kafka

import (
	"fmt"
	"github.com/Shopify/sarama"
	"time"
)

// SyncProducer 同步消息模式
func SyncProducer(address []string) {
	// 配置
	config := sarama.NewConfig()
	// 属性设置
	config.Producer.Return.Successes = true
	config.Producer.Timeout = 5 * time.Second
	// 创建生产者
	p, err := sarama.NewSyncProducer(address, config)
	if err != nil {
		fmt.Printf("sarama.NewSyncProducer err, message=%s \n", err)
		return
	}
	defer p.Close()

	// 定义topic名称
	topic := "topic1"
	// 发送的消息
	srcValue := "sync: this is a message. index=%d"

	for i := 1; i <= 10; i++ {
		value := fmt.Sprintf(srcValue, i)
		// 构建发送的消息
		msg := &sarama.ProducerMessage{
			Topic: topic,
			Value: sarama.StringEncoder(value),
		}
		// 发送消息
		partition, offset, errIn := p.SendMessage(msg)
		if errIn != nil {
			fmt.Printf("send message(%s) failed, err:%s \n", value, errIn)
		} else {
			fmt.Printf("message sent success partition=%d, offset=%d \n", partition, offset)
		}
		time.Sleep(1 * time.Second)
	}
}

// AsyncProducer 异步消息
func AsyncProducer(address []string) {

	config := sarama.NewConfig()
	// 等待服务器所有副本都保存成功后的响应
	config.Producer.RequiredAcks = sarama.WaitForAll
	// 随机向partition发送消息
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	// 是否等待成功和失败后的响应,只有上面的RequireAcks设置不是NoReponse这里才有用.
	config.Producer.Return.Successes = true
	config.Producer.Return.Errors = true
	// 设置使用的kafka版本,如果低于V0_10_0_0版本,消息中的timestrap没有作用.需要消费和生产同时配置
	// 注意，版本设置不对的话，kafka会返回很奇怪的错误，并且无法成功发送消息
	config.Version = sarama.V0_10_0_1

	fmt.Println("start make producer")
	// 使用配置,新建一个异步生产者
	producer, e := sarama.NewAsyncProducer(address, config)
	if e != nil {
		fmt.Println(e)
		return
	}
	defer producer.AsyncClose()

	// 循环判断哪个通道发送过来数据.
	fmt.Println("start goroutine")
	go func(p sarama.AsyncProducer) {
		for {
			select {
			case success := <-p.Successes():
				fmt.Println("send success offset: ", success.Offset,
					"timestamp: ", success.Timestamp.Format("15:04:05"),
					"partitions: ", success.Partition)
			case fail := <-p.Errors():
				fmt.Println("send failed, err: ", fail.Err)
			}
		}
	}(producer)

	var value string
	for i := 0; ; i++ {
		time.Sleep(500 * time.Millisecond)
		nowTime := time.Now()
		value = "this is a message time " + nowTime.Format("15:04:05")

		// 发送的消息,主题。
		// 注意：这里的msg必须得是新构建的变量，不然你会发现发送过去的消息内容都是一样的，因为批次发送消息的关系。
		msg := &sarama.ProducerMessage{
			Topic: "topic2",
			Value: sarama.ByteEncoder(value), // 将字符串转化为字节数组
		}

		// 使用通道发送
		producer.Input() <- msg
		fmt.Printf("send message: %s\n", value)
	}
}
