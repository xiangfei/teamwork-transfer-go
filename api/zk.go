package api

import (
	"encoding/json"
        "fmt"
        "strconv"
	"github.com/kataras/iris/v12"
        "teamwork-transfer-go/service"
)

type AwifiZkApi struct {
	svc *service.AwifiZkService
}

// init method
func NewAwifiZkApi(svc *service.AwifiZkService) *AwifiZkApi {

	return &AwifiZkApi{svc: svc}

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
	v := zk.svc.Kickcollect( zookeepertask.Mac)
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
	v := zk.svc.Create_service_once_task( string(str))
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

	v := zk.svc.Create_service_schedule_task( zookeepertask.Taskid, string(str))
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
	v := zk.svc.Destroy_service_schedule_task( zookeepertask.Taskid)
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
	v := zk.svc.Create_client_schedule_task(strconv.Itoa(zookeepertask.Taskid), zookeepertask.Mac, string(str))

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
	v := zk.svc.Destroy_client_schedule_task( strconv.Itoa(zookeepertask.Taskid), zookeepertask.Mac)
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
	v := zk.svc.List_client_schedule_task( zookeepertask.Mac)

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
	v := zk.svc.Create_client_once_task(zookeepertask.Mac, string(str))
	ctx.JSON(iris.Map{
		"code":    200,
		"message": "success",
		"data":    v,
	})

}

func (zk *AwifiZkApi) DestroyClientOnceTask(ctx iris.Context) {

	var zookeepertask ZookeeperTask
	ctx.ReadJSON(&zookeepertask)
	v := zk.svc.Destroy_client_once_task(strconv.Itoa(zookeepertask.Taskid), zookeepertask.Mac)

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
	v := zk.svc.List_client_once_task(zookeepertask.Mac)

	ctx.JSON(iris.Map{
		"code":    200,
		"message": "success",
		"data":    v,
	})

}

func (zk *AwifiZkApi) ListClientAgentTask(ctx iris.Context) {
	v := zk.svc.List_client_agent_task()

	ctx.JSON(iris.Map{
		"code":    200,
		"message": "success",
		"data":    v,
	})

}

func (zk *AwifiZkApi) ListServiceTask(ctx iris.Context) {
	v := zk.svc.List_service_task()

	ctx.JSON(iris.Map{
		"code":    200,
		"message": "success",
		"data":    v,
	})

}

func (zk *AwifiZkApi) ListServiceScheduleTask(ctx iris.Context) {

	v := zk.svc.List_service_schedule_task()

	ctx.JSON(iris.Map{
		"code":    200,
		"message": "success",
		"data":    v,
	})

}

func (zk *AwifiZkApi) ListMasterClients(ctx iris.Context) {

      v := zk.svc.List_master_client()
     
      current := zk.svc.Current_master_client()
   
        ctx.JSON(iris.Map{
                "code":    200,
                "message": "success",
                "data":   iris.Map{"current":  current , "list":  v } ,
        })


}

func (zk *AwifiZkApi) ListControlClients(ctx iris.Context) {

    v := zk.svc.List_control_clients()

        ctx.JSON(iris.Map{
                "code":    200,
                "message": "success",
                "data":    v,
        })


}

func (zk *AwifiZkApi) ListCollectClients(ctx iris.Context) {

    v := zk.svc.List_collect_clients()

        ctx.JSON(iris.Map{
                "code":    200,
                "message": "success",
                "data":    v,
        })


}

func (zk *AwifiZkApi) ListGatewayClients(ctx iris.Context) {

    v := zk.svc.List_gateway_clients()

        ctx.JSON(iris.Map{
                "code":    200,
                "message": "success",
                "data":    v,
        })


}

func (zk *AwifiZkApi) ListHeartbeatClients(ctx iris.Context) {

    v := zk.svc.List_heartbeat_clients()

        ctx.JSON(iris.Map{
                "code":    200,
                "message": "success",
                "data":    v,
        })


}

func (zk *AwifiZkApi) ListNotifyClients(ctx iris.Context) {

    v := zk.svc.List_notify_clients()

        ctx.JSON(iris.Map{
                "code":    200,
                "message": "success",
                "data":    v,
        })


}

func (zk *AwifiZkApi) ListAgentClients(ctx iris.Context) {

    v := zk.svc.List_agent_clients()

        ctx.JSON(iris.Map{
                "code":    200,
                "message": "success",
                "data":    v,
        })


}



