package main

import (
	"net/http"
        "runtime" 
	"teamwork-transfer-go/api"
	"teamwork-transfer-go/client"
	//"teamwork-transfer-go/config"
	"github.com/kataras/iris/v12"
	"teamwork-transfer-go/service"
)

func main() {
        runtime.GOMAXPROCS(runtime.NumCPU())
	zkclient := client.GetAwifiZkSingleton()
	defer zkclient.Close()
	zkservice := service.NewZkServiceFromClient(zkclient)
	zkapi := api.NewAwifiZkApi(zkservice)
        baseapi :=api.NewAwifiBaseApi()
	app := iris.New()

        // base api
	app.Get("/", baseapi.Index)
	app.Get("/json", baseapi.IndexJson) 

        // http api
        httpclient :=client.GetAwifiHttpSingleton() 
        httpservice := service.NewHttpServiceFromClient(httpclient)
        httpapi :=api.NewAwifiHttpApi(httpservice)
        app.Post("/sendmsg" ,  httpapi.SendMessage)       
        app.Get("/testmsg" , httpapi.TestMessage)       
        // zk api
	app.Post("/kickcollect", zkapi.Kickcollect)
	app.Post("/create_service_task", zkapi.CreateServiceTask)
	app.Post("/create_service_schedule_task", zkapi.CreateServiceScheduleTask)
	app.Post("/destroy_service_schedule_task", zkapi.DestroyServiceScheduleTask)
	app.Post("/create_client_schedule_task", zkapi.CreateClientScheduleTask)
	app.Post("/destroy_client_schedule_task", zkapi.DestroyClientScheduleTask)
	app.Post("/list_client_schedule_task", zkapi.ListClientScheduleTask)
	app.Post("/create_client_once_task", zkapi.CreateClientOnceTask)
	app.Post("/destroy_client_once_task", zkapi.DestroyClientOnceTask)
	app.Post("/list_client_once_task", zkapi.ListClientOnceTask)
	app.Post("/list_client_agent_task", zkapi.ListClientAgentTask)
	app.Post("/list_service_task", zkapi.ListServiceTask)
	app.Post("/list_service_schedule_task", zkapi.ListServiceScheduleTask)

        // server
	srv := &http.Server{Addr: ":4568"}
	app.Run(iris.Server(srv))

}

