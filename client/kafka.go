package client

import (
	"fmt"
	"github.com/Shopify/sarama"
	"sync" 
        "math/rand"
	"teamwork-transfer-go/config"
)

var (
	singlekafkaInstance *AwifiKafka
	lock                = &sync.Mutex{}
)

type AwifiKafka struct {
	kafka_url     []string
	asyncproducer sarama.AsyncProducer
}

type AwifiKafkaList struct {
	awifi_kafka_list []*AwifiKafka
}

func NewAwifiKafkaList() *AwifiKafkaList {

 return &AwifiKafkaList{awifi_kafka_list: []*AwifiKafka{} }
}

func (list *AwifiKafkaList)  AddKafkaClient(client *AwifiKafka) {
  list.awifi_kafka_list = append(list.awifi_kafka_list, client )

} 


func (list *AwifiKafkaList) Close() {

	for _, kafka := range list.awifi_kafka_list {
		kafka.Close()
	}

}

func (list *AwifiKafkaList) RandomGetClient()   *AwifiKafka {

	randomIndex := rand.Intn(len(list.awifi_kafka_list))
	return list.awifi_kafka_list[randomIndex]

}

func GetAwifiKafkaSingleton() *AwifiKafka {

	if singlekafkaInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if singlekafkaInstance == nil {
			singlekafkaInstance = NewAwifiKafka()
			return singlekafkaInstance
		}
	}
	return singlekafkaInstance

}

func NewAwifiKafka() *AwifiKafka {
	baseconfig := config.NewConfig()

	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true
	producer, e := sarama.NewAsyncProducer(baseconfig.Kafka_url, config)
	if e != nil {
		panic(e)
	}
	go func(p sarama.AsyncProducer) {
		for {
			select {
			case suc := <-p.Successes():
				fmt.Println("offset: ", suc.Offset, "partitions: ", suc.Partition, "topic:", suc.Topic)
			case fail := <-p.Errors():
				fmt.Println("err: ", fail.Err)
			}
		}
	}(producer)

	instance := &AwifiKafka{
		kafka_url:     baseconfig.Kafka_url,
		asyncproducer: producer,
	}
	return instance

}

func (kafka *AwifiKafka) Close() {

	kafka.asyncproducer.AsyncClose()

}

func (kafka *AwifiKafka) DeliverMessage(topic string, message string) {
	msg := &sarama.ProducerMessage{
		Topic: topic}
	msg.Value = sarama.ByteEncoder(message)
	kafka.asyncproducer.Input() <- msg

}

//func main() {
//	p := NewProducer()
//	defer p.AsyncClose()
//	Produce(p, "xxx", "xxxx")
//	fmt.Println(p)
//}

