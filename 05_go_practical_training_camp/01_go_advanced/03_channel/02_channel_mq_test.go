package _3_channel

import (
	"fmt"
	"testing"
)

type Broker struct {
	consumers []*Consumer
}

type Consumer struct {
	ch chan string
}

func (b *Broker) produce(msg string) {
	// 轮询给消费者发送消息
	for _, v := range b.consumers {
		v.ch <- msg
	}
}

func (b *Broker) subscribe(consumer *Consumer) {
	b.consumers = append(b.consumers, consumer)
}

func TestMq1(t *testing.T) {
	// 初始化一个Broker节点
	b := &Broker{
		consumers: make([]*Consumer, 0, 4),
	}

	// 创建2个消费者
	consumer1 := &Consumer{
		ch: make(chan string, 1),
	}
	consumer2 := &Consumer{
		ch: make(chan string, 1),
	}

	// 这2个消费者订阅Broker
	b.subscribe(consumer1)
	b.subscribe(consumer2)

	// 生产者发送一个消息
	b.produce("一条消息")

	// 2个消费者拿到了刚才生产者发送的消息
	fmt.Println(<-consumer1.ch)
	fmt.Println(<-consumer2.ch)

	// Output
	// 一条消息
	// 一条消息
}
