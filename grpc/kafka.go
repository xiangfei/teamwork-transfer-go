package grpc

import (
	"context"
	pb "teamwork-transfer-go/proto"
	"teamwork-transfer-go/service"
)

type Message struct {
	service *service.AwifiKafkaService
}


func NewMessage(s *service.AwifiKafkaService) *Message {

	return &Message{service: s}
}






func (c *Message) SendMessage(ctx context.Context, in *pb.KafkaMessageRequest) (*pb.KafkaMessageReply, error) {

	flag := c.service.Deliver_Messages(in.GetTopic(), in.GetMessage())
	if flag {
		return &pb.KafkaMessageReply{Message: "message  success" + in.GetMessage() + in.GetTopic()}, nil
	} else {
		return &pb.KafkaMessageReply{Message: "message  error" + in.GetMessage() + in.GetTopic()}, nil
	}
	return &pb.KafkaMessageReply{Message: "message  error" + in.GetMessage() + in.GetTopic()}, nil
}
