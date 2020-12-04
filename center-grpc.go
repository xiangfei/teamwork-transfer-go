package main

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"google.golang.org/grpc"
	"teamwork-transfer-go/client"
	teamworkgrpc "teamwork-transfer-go/grpc"
	pb "teamwork-transfer-go/proto"
	"teamwork-transfer-go/service"
        "fmt"
)

func main() {
	kafkaclient := client.GetAwifiKafkaSingleton()
	defer kafkaclient.Close()
	kafkaservice := service.NewKafkaServiceFromClient(kafkaclient)
        fmt.Println(kafkaservice)
	grpcServer := grpc.NewServer()
	//myService := &teamworkgrpc.Message{}   //  {kafkaservice}
        myService :=teamworkgrpc.NewMessage(kafkaservice)
	pb.RegisterMessageServer(grpcServer, myService)
	app := iris.New()

	rootApp := mvc.New(app)
	rootApp.Handle(myService, mvc.GRPC{
		Server:      grpcServer, // Required.
		ServiceName: "/",        // Required.
		Strict:      false,
	})
	app.Run(iris.TLS(":443", "./openssl/server.crt", "./openssl/server.key"))

}
