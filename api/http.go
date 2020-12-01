package api

import (
	"encoding/json"
	"github.com/kataras/iris/v12"
	"teamwork-transfer-go/service"
)

type AwifiHttpApi struct {
	service *service.AwifiHttpService
}

// init
func NewAwifiHttpApi(s *service.AwifiHttpService) *AwifiHttpApi {

	return &AwifiHttpApi{service: s}
}

// kafka message  struct
type KafkaMessage struct {
	Topic string `json:"topic" msgpack:"topic"`

	Message string `json:"message" msgpack:"message"`
}

func (http *AwifiHttpApi) SendMessage(ctx iris.Context) {

	var kafkamessage KafkaMessage
	ctx.ReadJSON(&kafkamessage)
	values := map[string]string{"message": kafkamessage.Message, "topic": kafkamessage.Topic}

	jsonValue, _ := json.Marshal(values)

	flag := http.service.JsonPost("/sendmsg", jsonValue)
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

