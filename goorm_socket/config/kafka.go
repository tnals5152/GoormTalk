package config

import (
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/Shopify/sarama"
)

type Producer struct {
	ChatProducer sarama.SyncProducer
}

type Consumer struct {
	Consumer          sarama.Consumer
	PartitionConsumer [2]sarama.PartitionConsumer
}

func KafkaConsumer() {

	consumer, err := sarama.NewConsumer([]string{
		"54.180.85.30:9092",
	}, nil)
	fmt.Println(2, err)
	c := Consumer{Consumer: consumer}

	partitions, err := consumer.Partitions("test")
	fmt.Println(1, err, partitions)
	for i, v := range partitions {
		c.PartitionConsumer[i], err = consumer.ConsumePartition("test", v, sarama.OffsetNewest)
		fmt.Println(3, err)

	}

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)
	for _, v := range c.PartitionConsumer {
		go func(v sarama.PartitionConsumer) {
			for {
				select {
				case msg := <-v.Messages():
					fmt.Println("Message   : ", string(msg.Value))
					fmt.Println("Partition : ", msg.Partition)
					fmt.Println("Offset    : ", msg.Offset)
				}
			}
		}(v)
	}
	time.Sleep(10000 * time.Second)
}

func KafkaProduce() {
	config := sarama.NewConfig()
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Return.Successes = true
	c, err := sarama.NewSyncProducer([]string{
		"54.180.85.30:9092",
	}, config)
	fmt.Println(err)
	p := &Producer{ChatProducer: c}
	time.Sleep(5 * time.Second)
	p.Send("soomin!!")
}

func (producer *Producer) Send(message string) {
	partition, offset, err := producer.ChatProducer.SendMessage(&sarama.ProducerMessage{
		Topic: "test",
		Value: sarama.StringEncoder(message),
	})
	fmt.Println(err)
	fmt.Println(partition, "????")
	fmt.Println(offset)

}
