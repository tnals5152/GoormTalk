package config

import (
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/Shopify/sarama"

	"goorm_socket/utils"
)

type ProducerStruct struct {
	ChatProducer sarama.SyncProducer //close필요
}

type Consumer struct {
	Consumer          sarama.Consumer
	PartitionConsumer []sarama.PartitionConsumer
}

var KAFKA string
var Broker *sarama.Broker
var Producer *ProducerStruct

//카프카 연결 IP세팅
func KafkaSetting() bool {
	KAFKA = fmt.Sprintf("%s:%s", os.Getenv("KAFKA_IP"), os.Getenv("KAFKA_PORT"))
	if KAFKA == ":" {
		return false
	}
	return true
}

//카프카 연결 함수
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

//카프카 컨슈머 생성 함수
func KafkaConsumer() {

	consumer, err := sarama.NewConsumer([]string{
		KAFKA,
	}, nil)
	utils.IfErrorMakePanic(err, "can not create consumer")
	c := Consumer{Consumer: consumer}

	//DB에서 채팅방 select하여 ID리스트 반환 -> ID 기준으로 for문

	//ID를 이름으로 토픽 이름이 일치하는게 있는지 확인 및 없으면 토픽 생성(for)

	//해당 토픽이 가지고 있는 파티션 리스트 반환 없으면 에러(for문으로 변경 예정)
	topic := "2"
	partitions, err := consumer.Partitions(topic)
	utils.IfErrorMakePanic(err, fmt.Sprintf("get %s topic partition", topic))

	// timeSource := rand.NewSource(time.Now().UnixNano())
	// random := rand.New(timeSource)

	// partition := random.Intn(len(partitions)-1) + 1
	partitionConsumer, err := consumer.ConsumePartition(topic, partitions[1], sarama.OffsetNewest)
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
					fmt.Println("Consumer  : ", c)
				}
			}
		}(v)
	}
	time.Sleep(10000 * time.Second)
}

//카프카 프로듀서 생성 함수
//서버 하나당 하나 생성
func KafkaProduce() *ProducerStruct {
	config := sarama.NewConfig()
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Return.Successes = true
	c, err := sarama.NewSyncProducer([]string{
		KAFKA,
	}, config)
	utils.IfErrorMakePanic(err, "can not create produce")
	producer := &ProducerStruct{ChatProducer: c}
	Producer = producer
	// time.Sleep(5 * time.Second)
	// producer.Send("soomin!!")
	return producer
}

func (producer *ProducerStruct) Send(topic string, message string) {
	producer.SendWithPartitionNumber(topic, message, 1)
}

func (producer *ProducerStruct) SendWithPartitionNumber(topic string, message string, partition int32) {
	_, _, err := producer.ChatProducer.SendMessage(&sarama.ProducerMessage{
		Topic:     topic,
		Partition: partition,
		Value:     sarama.StringEncoder(message),
	})
	utils.ErrorCheck(err)
}

//채팅방 별 토픽 생성
func MakeTopic(topic string) {

	request := sarama.CreateTopicsRequest{
		Timeout: time.Second * 10,
		TopicDetails: map[string]*sarama.TopicDetail{
			topic: &sarama.TopicDetail{
				NumPartitions:     int32(2), //partition 0 - elasticsearch, partition 1: consumer
				ReplicationFactor: int16(1),
			},
		},
	}
	response, err := Broker.CreateTopics(&request)
	utils.ErrorCheck(err)
	fmt.Println("response: ", response.TopicErrors)
}
