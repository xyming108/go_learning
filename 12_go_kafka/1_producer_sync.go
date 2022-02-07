package main

import (
	"fmt"
	"github.com/Shopify/sarama"
	"os"
	"time"
)

func main() {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.Timeout = 5 * time.Second

	p, err := sarama.NewSyncProducer([]string{"10.0.54.57:9092"}, config)
	if err != nil {
		fmt.Fprintf(os.Stdout, "sarama.NewSyncProducer err, message=%v \n", err)
		return
	}
	defer p.Close()

	for i := 0; i < 10; i++ {
		msg := &sarama.ProducerMessage{
			Topic: "test",
			Value: sarama.ByteEncoder(fmt.Sprintf("sync: this is a message. index=%d", i)),
		}

		if part, offset, err := p.SendMessage(msg); err != nil {
			fmt.Fprintf(os.Stdout, "发送失败，Error: %v \n", err)
		} else {
			fmt.Fprintf(os.Stdout, "发送成功，partition=%d, offset=%d \n", part, offset)
		}
	}
}
