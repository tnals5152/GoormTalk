package config

import (
	"fmt"
	"log"
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

var KAFKA string

func KafkaSetting() bool {
	KAFKA = fmt.Sprintf("%s:%s", os.Getenv("KAFKA_IP"), os.Getenv("KAFKA_PORT"))
	if KAFKA == "" {
		return false
	}
	return true
}
func KafkaConsumer() {

	consumer, err := sarama.NewConsumer([]string{
		KAFKA,
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
		KAFKA,
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
	fmt.Println(partition)
	fmt.Println(offset)

}

func MakeTopic(userID string) {
	broker := sarama.NewBroker(KAFKA)

	config := sarama.NewConfig()
	config.Version = sarama.V2_8_0_0
	err := broker.Open(config)

	connected, err := broker.Connected()
	if !connected {
		log.Println(err)
	}

	request := sarama.CreateTopicsRequest{
		Timeout: time.Second * 10,
		TopicDetails: map[string]*sarama.TopicDetail{
			userID: &sarama.TopicDetail{
				NumPartitions:     int32(1),
				ReplicationFactor: int16(1),
			},
		},
	}
	response, err := broker.CreateTopics(&request)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(response)
}
