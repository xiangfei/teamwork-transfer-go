package main

import (
	//"net/http"
	"github.com/valyala/tcplisten"

	"teamwork-transfer-go/api"
	"teamwork-transfer-go/client"
	//"teamwork-transfer-go/config"
	//"github.com/valyala/tcplisten"
	"github.com/kataras/iris/v12"
	"teamwork-transfer-go/service"
)

func main() {
	kafkaclient := client.GetAwifiKafkaSingleton()
	defer kafkaclient.Close()
	kafkaservice := service.NewKafkaServiceFromClient(kafkaclient)
	kafkaapi := api.NewAwifiKafkaApi(kafkaservice)

	baseapi := api.NewAwifiBaseApi()

	app := iris.New()

	// base api
	app.Get("/", baseapi.Index)
	app.Get("/json", baseapi.IndexJson)

	// kafka api
	app.Post("/sendmsg", kafkaapi.SendMessage)
	app.Get("/testmsg", kafkaapi.Test)

	//  server config
	//srv := &http.Server{Addr: ":4569"}
	//app.Run(iris.Server(srv))
	listenerCfg := tcplisten.Config{
		ReusePort:   true,
		DeferAccept: true,
		FastOpen:    true,
	}

	l, err := listenerCfg.NewListener("tcp4", ":4569")
	if err != nil {
		app.Logger().Fatal(err)
	}
	app.Run(iris.Listener(l))

}

