package config

import (
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/Shopify/sarama"

	"goorm_socket/utils"
)

type Producer struct {
	ChatProducer sarama.SyncProducer //close필요
}

type Consumer struct {
	Consumer          sarama.Consumer
	PartitionConsumer []sarama.PartitionConsumer
}

var KAFKA string
var Broker *sarama.Broker

func KafkaSetting() bool {
	KAFKA = fmt.Sprintf("%s:%s", os.Getenv("KAFKA_IP"), os.Getenv("KAFKA_PORT"))
	if KAFKA == "" {
		return false
	}
	return true
}

func ConnectBroker() {
	broker := sarama.NewBroker(KAFKA)

	config := sarama.NewConfig()
	config.Version = sarama.V2_8_0_0
	err := broker.Open(config)
	utils.IfErrorMakePanic(err, "can not open broker")

	connected, err := broker.Connected()
	if !connected {
		utils.IfErrorMakePanic(err, "can not connect broker")
	}
	Broker = broker
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
	utils.IfErrorMakePanic(err, "can not create consumer")
	c := Consumer{Consumer: consumer}

	//DB에서 채팅방 select하여 ID리스트 반환 -> ID 기준으로 for문

	//ID를 이름으로 토픽 이름이 일치하는게 있는지 확인 및 없으면 토픽 생성(for)

	//해당 토픽이 가지고 있는 파티션 리스트 반환 없으면 에러(for문으로 변경 예정)
	topic := "test"
	partitions, err := consumer.Partitions(topic)
	utils.IfErrorMakePanic(err, fmt.Sprintf("get %s topic partition", topic))

	//파티션 개수는 서버에 연결된 사람 수만큼 생성...?
	// timeSource := rand.NewSource(time.Now().UnixNano())
	// random := rand.New(timeSource)

	// partition := random.Intn(len(partitions)-1) + 1
	//파티션 구독 중인지 알 수 있는지 확인 및 알 수 있으면 사용자 없는 곳에 연결 (해당 채팅방 사람 수만큼 할 수 X <- 한 사람이 여러개 틀 수 있음)
	partitionConsumer, err := consumer.ConsumePartition(topic, partitions[0], sarama.OffsetNewest)
	utils.IfErrorMakePanic(err, fmt.Sprintf("get partition error from %s topic", topic))
	c.PartitionConsumer = append(c.PartitionConsumer, partitionConsumer)

	//모든 파티션에 대해 구독 처리 -> X   하나의 파티션만 구독
	// for i, v := range partitions {
	// 	c.PartitionConsumer[i], err = consumer.ConsumePartition("test", v, sarama.OffsetNewest)
	// 	fmt.Println(3, err)

	// }

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

//서버 하나당 하나 생성
func KafkaProduce() *Producer {
	config := sarama.NewConfig()
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Return.Successes = true
	c, err := sarama.NewSyncProducer([]string{
		KAFKA,
	}, config)
	utils.IfErrorMakePanic(err, "can not create produce")
	producer := &Producer{ChatProducer: c}
	// time.Sleep(5 * time.Second)
	// producer.Send("soomin!!")
	return producer
}

func (producer *Producer) Send(topic string, message string) {
	_, _, err := producer.ChatProducer.SendMessage(&sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(message),
	})
	utils.ErrorCheck(err)

}

func MakeTopic(topic string) {

	request := sarama.CreateTopicsRequest{
		Timeout: time.Second * 10,
		TopicDetails: map[string]*sarama.TopicDetail{
			topic: &sarama.TopicDetail{
				NumPartitions:     int32(1),
				ReplicationFactor: int16(1),
			},
		},
	}
	response, err := Broker.CreateTopics(&request)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(response)
}
