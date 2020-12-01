package main

import (
        "net/http"
	"teamwork-transfer-go/api"
	"teamwork-transfer-go/client"
	//"teamwork-transfer-go/config"
	"github.com/kataras/iris/v12"
	"teamwork-transfer-go/service"
)

func main() {
	kafkaclient := client.GetAwifiKafkaSingleton()
	defer kafkaclient.Close()
	kafkaservice := service.NewKafkaServiceFromClient(kafkaclient)
	kafkaapi := api.NewAwifiKafkaApi(kafkaservice)
        
        baseapi :=api.NewAwifiBaseApi()

	app := iris.New()

        // base api
	app.Get("/", baseapi.Index)
	app.Get("/json", baseapi.IndexJson) 

        // kafka api
        app.Post("/sendmsg" ,  kafkaapi.SendMessage)       
        app.Get("/testmsg" , kafkaapi.Test) 

        //  server config
        srv := &http.Server{Addr: ":4567"}
	app.Run(iris.Server(srv))

}

