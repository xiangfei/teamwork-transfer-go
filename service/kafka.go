package service

import (
        "teamwork-transfer-go/client"
)

type AwifiKafkaService struct {
        //zkclient []*client.AwifiKafka
        zkclient *client.AwifiKafkaList
}


func NewKafkaServiceFromClient(zkclient *client.AwifiKafkaList) *AwifiKafkaService {

        return &AwifiKafkaService{zkclient}

}



func (service *AwifiKafkaService) Deliver_Messages(topic string , message string) bool{

    service.zkclient.RandomGetClient().DeliverMessage(topic , message)
    return true 

}


