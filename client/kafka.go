package client

import (
	"fmt"
	"github.com/Shopify/sarama"
        "teamwork-transfer-go/config"
        "sync"
)

var (
        singlekafkaInstance *AwifiKafka
        lock           = &sync.Mutex{}
)


type AwifiKafka struct {
	kafka_url     []string
	asyncproducer sarama.AsyncProducer
}



func GetAwifiKafkaSingleton() *AwifiKafka {

        if singlekafkaInstance == nil {
                lock.Lock()
                defer lock.Unlock()
                if singlekafkaInstance  == nil {
                        singlekafkaInstance = initialize_kafka()
                        return singlekafkaInstance
                }
        }
        return singlekafkaInstance

}




func initialize_kafka() *AwifiKafka {
	baseconfig := config.SingleConfigInstance()

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
                                fmt.Println("offset: ", suc.Offset , "partitions: ", suc.Partition  , "topic:" , suc.Topic)
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

