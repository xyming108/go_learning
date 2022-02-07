package main

import (
	"fmt"
	"github.com/Shopify/sarama"
	"math/rand"
	"strconv"
)

func InitKafkaProducer() (sarama.AsyncProducer, error) {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll          //等待服务器所有副本都保存成功后的响应
	config.Producer.Partitioner = sarama.NewRandomPartitioner //随机的分区类型
	//是否等待成功和失败后的响应,只有上面的RequireAcks设置不是NoReponse这里才有用
	config.Producer.Return.Successes = true
	config.Producer.Return.Errors = true
	config.Version = sarama.V2_5_0_0 //设置使用的kafka版本,如果低于V0_10_0_0版本,消息中的timestrap没有作用.需要消费和生产同时配置
	return sarama.NewAsyncProducer([]string{"10.0.54.57:9092"}, config)
}

//ProducerMsg 调用该函数生产消息
func ProducerMsg() {
	prod, err := InitKafkaProducer()
	if err != nil {
		panic(err)
	}
	defer prod.Close()
	msg := &sarama.ProducerMessage{
		Topic:     "test",
		Key:       sarama.StringEncoder("test"),
		Partition: 1,
	}
	msgchan := prod.Input()

	for i := 0; i < 10; i++ {
		msg.Value = sarama.StringEncoder("msg id is :" + strconv.Itoa(rand.Intn(100)))
		msgchan <- msg
		select {
		case suc := <-prod.Successes():
			fmt.Println(suc)
		case err := <-prod.Errors():
			fmt.Println(err)
		}
	}
}

func main() {
	ProducerMsg()
}
