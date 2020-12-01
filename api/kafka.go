package api

import (
	"github.com/kataras/iris/v12"
	"teamwork-transfer-go/service"
)

type AwifiKafkaApi struct {
	service *service.AwifiKafkaService
}

// init
func NewAwifiKafkaApi(s *service.AwifiKafkaService) *AwifiKafkaApi {

	return &AwifiKafkaApi{service: s}
}


func (kafka *AwifiKafkaApi) SendMessage(ctx iris.Context) {

	var kafkamessage KafkaMessage
	ctx.ReadJSON(&kafkamessage)

	flag := kafka.service.Deliver_Messages(kafkamessage.Topic, kafkamessage.Message)
	if flag {
		ctx.JSON(iris.Map{

			"code": 200,

			"message": "success",

			"data": "",
		})

	} else {
		ctx.JSON(iris.Map{
			"code":    500,
			"message": "error",
			"data":    "error",
		})

	}

}


func (kafka *AwifiKafkaApi) Test(ctx iris.Context) {


        flag := kafka.service.Deliver_Messages("host.teamwork.performance", "xxxxxxxxxxxxx")
        if flag {
                ctx.JSON(iris.Map{

                        "code": 200,

                        "message": "success",

                        "data": "",
                })

        } else {
                ctx.JSON(iris.Map{
                        "code":    500,
                        "message": "error",
                        "data":    "error",
                })

        }

}

