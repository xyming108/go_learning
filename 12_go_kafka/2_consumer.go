package main

import (
	"fmt"
	"github.com/Shopify/sarama"
	"os"
)

func main() {
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true
	config.Version = sarama.V0_11_0_2

	c, err := sarama.NewConsumer([]string{"10.0.54.57:9092"}, config)
	if err != nil {
		fmt.Fprintf(os.Stdout, "sarama.NewConsumer err, message=%v \n", err)
		return
	}
	defer c.Close()

	cp, err := c.ConsumePartition("test", 0, sarama.OffsetOldest)
	if err != nil {
		fmt.Fprintf(os.Stdout, "try create partition_consumer err, message=%v \n", err)
		return
	}
	defer cp.Close()

	for {
		select {
		case msg := <-cp.Messages():
			fmt.Fprintf(os.Stdout, "msg offset: %d, partition: %d, timestamp: %s, value: %s \n", msg.Offset, msg.Partition, msg.Timestamp.String(), string(msg.Value))
		case err := <-cp.Errors():
			fmt.Fprintf(os.Stdout, "err :%v \n", err)
		}
	}
}
