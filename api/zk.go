package api

import (
	"encoding/json"
        "fmt"
        "strconv"
	"github.com/kataras/iris/v12"
	"teamwork-transfer-go/service"
)

type AwifiZkApi struct {
	service *service.AwifiZkService
}

// init method
func NewAwifiZkApi(service *service.AwifiZkService) *AwifiZkApi {

	return &AwifiZkApi{service: service}

}

// http post body
type ZookeeperTask struct {
	Mac    string `json: "mac" msgpack:"mac"`
	Taskid int    `json: "taskid" msgpack:"taskid"`
	Opts   Opts   `json: "opts" msgpack:"opts"`
}

type ZookeeperServiceScheduleTask struct {
	Taskid string `json: "taskid" msgpack:"taskid"`
	Opts   Opts   `json: "opts" msgpack:"opts"`
}
type Opts map[string]interface{}

// method

func (zk *AwifiZkApi) Kickcollect(ctx iris.Context) {

	var zookeepertask ZookeeperTask
	ctx.ReadJSON(&zookeepertask)
	v := zk.service.Kickcollect( zookeepertask.Mac)
	if v {
		ctx.JSON(iris.Map{
			"code":    200,
			"message": "success",
			"data":    "",
		})
	} else {

		ctx.JSON(iris.Map{
			"code":    500,
			"message": "error",
			"data":    "",
		})

	}

}

func (zk *AwifiZkApi) CreateServiceTask(ctx iris.Context) {

	var zookeepertask ZookeeperTask
	ctx.ReadJSON(&zookeepertask)
	str, err := json.Marshal(zookeepertask.Opts)
	if err != nil {
		fmt.Println(err)
	}
	v := zk.service.Create_service_once_task( string(str))
	ctx.JSON(iris.Map{
		"code":    200,
		"message": "success",
		"data":    v,
	})

}

func (zk *AwifiZkApi) CreateServiceScheduleTask(ctx iris.Context) {

	var zookeepertask ZookeeperServiceScheduleTask
	ctx.ReadJSON(&zookeepertask)
	str, err := json.Marshal(zookeepertask.Opts)
	if err != nil {
		fmt.Println(err)
	}

	v := zk.service.Create_service_schedule_task( zookeepertask.Taskid, string(str))
	if v {
		ctx.JSON(iris.Map{
			"code":    200,
			"message": "success",
			"data":    "",
		})
	} else {

		ctx.JSON(iris.Map{
			"code":    500,
			"message": "error",
			"data":    "",
		})

	}

}

func (zk *AwifiZkApi) DestroyServiceScheduleTask(ctx iris.Context) {
	var zookeepertask ZookeeperServiceScheduleTask
	ctx.ReadJSON(&zookeepertask)
	v := zk.service.Destroy_service_schedule_task( zookeepertask.Taskid)
	if v {
		ctx.JSON(iris.Map{
			"code":    200,
			"message": "success",
			"data":    "",
		})
	} else {

		ctx.JSON(iris.Map{
			"code":    500,
			"message": "error",
			"data":    "",
		})

	}

}

func (zk *AwifiZkApi) CreateClientScheduleTask(ctx iris.Context) {

	var zookeepertask ZookeeperTask
	ctx.ReadJSON(&zookeepertask)
	str, err := json.Marshal(zookeepertask.Opts)
	if err != nil {
		fmt.Println(err)
	}
	v := zk.service.Create_client_schedule_task(strconv.Itoa(zookeepertask.Taskid), zookeepertask.Mac, string(str))

	if v {
		ctx.JSON(iris.Map{
			"code":    200,
			"message": "success",
			"data":    "",
		})
	} else {

		ctx.JSON(iris.Map{
			"code":    500,
			"message": "error",
			"data":    "",
		})

	}

}

func (zk *AwifiZkApi) DestroyClientScheduleTask(ctx iris.Context) {

	var zookeepertask ZookeeperTask
	ctx.ReadJSON(&zookeepertask)
	v := zk.service.Destroy_client_schedule_task( strconv.Itoa(zookeepertask.Taskid), zookeepertask.Mac)
	if v {
		ctx.JSON(iris.Map{
			"code":    200,
			"message": "success",
			"data":    "",
		})
	} else {

		ctx.JSON(iris.Map{
			"code":    500,
			"message": "error",
			"data":    "",
		})

	}

}

func (zk *AwifiZkApi) ListClientScheduleTask(ctx iris.Context) {
	var zookeepertask ZookeeperTask
	ctx.ReadJSON(&zookeepertask)
	v := zk.service.List_client_schedule_task( zookeepertask.Mac)

	ctx.JSON(iris.Map{
		"code":    200,
		"message": "success",
		"data":    v,
	})

}

func (zk *AwifiZkApi) CreateClientOnceTask(ctx iris.Context) {
	var zookeepertask ZookeeperTask
	ctx.ReadJSON(&zookeepertask)
	str, err := json.Marshal(zookeepertask.Opts)
	if err != nil {
		fmt.Println(err)
	}
	v := zk.service.Create_client_once_task(zookeepertask.Mac, string(str))
	ctx.JSON(iris.Map{
		"code":    200,
		"message": "success",
		"data":    v,
	})

}

func (zk *AwifiZkApi) DestroyClientOnceTask(ctx iris.Context) {

	var zookeepertask ZookeeperTask
	ctx.ReadJSON(&zookeepertask)
	v := zk.service.Destroy_client_once_task(strconv.Itoa(zookeepertask.Taskid), zookeepertask.Mac)

	if v {
		ctx.JSON(iris.Map{
			"code":    200,
			"message": "success",
			"data":    "",
		})
	} else {

		ctx.JSON(iris.Map{
			"code":    500,
			"message": "error",
			"data":    "",
		})

	}

}

func (zk *AwifiZkApi) ListClientOnceTask(ctx iris.Context) {
	var zookeepertask ZookeeperTask
	ctx.ReadJSON(&zookeepertask)
	v := zk.service.List_client_once_task(zookeepertask.Mac)

	ctx.JSON(iris.Map{
		"code":    200,
		"message": "success",
		"data":    v,
	})

}

func (zk *AwifiZkApi) ListClientAgentTask(ctx iris.Context) {
	v := zk.service.List_client_agent_task

	ctx.JSON(iris.Map{
		"code":    200,
		"message": "success",
		"data":    v,
	})

}

func (zk *AwifiZkApi) ListServiceTask(ctx iris.Context) {
	v := zk.service.List_service_task

	ctx.JSON(iris.Map{
		"code":    200,
		"message": "success",
		"data":    v,
	})

}

func (zk *AwifiZkApi) ListServiceScheduleTask(ctx iris.Context) {

	v := zk.service.List_service_schedule_task

	ctx.JSON(iris.Map{
		"code":    200,
		"message": "success",
		"data":    v,
	})

}

