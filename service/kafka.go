package service

import (
        "teamwork-transfer-go/client"
)

type AwifiKafkaService struct {
        zkclient *client.AwifiKafka
}


func NewKafkaServiceFromClient(zkclient *client.AwifiKafka) *AwifiKafkaService {

        return &AwifiKafkaService{zkclient}

}



func (service *AwifiKafkaService) Deliver_Messages(topic string , message string) bool{


    service.zkclient.DeliverMessage(topic , message)
    return true 

}
