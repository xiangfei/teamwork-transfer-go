package service

import (
	"teamwork-transfer-go/client"
)

type AwifiZkService struct {
	zkclient *client.AwifiZk
}

func NewZkServiceFromClient(zkclient *client.AwifiZk) *AwifiZkService {

	return &AwifiZkService{zkclient}

}

func (svc *AwifiZkService) Create_client_schedule_task(taskid string, mac string, opts string) bool {
	base := "/awifi/task/collect"
	task_path := base + "/" + mac
	full_task_id := task_path + "/" + taskid
	s := svc.zkclient.Exist(full_task_id)
	if s {
		svc.zkclient.Del(full_task_id)
	}
	return svc.zkclient.Add(full_task_id, opts)

}

func (svc *AwifiZkService) List_client_schedule_task(mac string) []string {

	base := "/awifi/task/collect"
	task_path := base + "/" + mac
	return svc.zkclient.Children(task_path)
}

func (svc *AwifiZkService) Destroy_client_schedule_task(taskid string, mac string) bool {

	base := "/awifi/task/collect"
	task_path := base + "/" + mac
	full_task_id := task_path + "/" + taskid
	s := svc.zkclient.Exist(full_task_id)
	if s {
		return svc.zkclient.Del(full_task_id)
	}
	return true
}

func (svc *AwifiZkService) Kickcollect(mac string) bool {

	base := "/awifi/client/collect"
	task_path := base + "/" + mac
	return svc.zkclient.Del(task_path)
}

func (svc *AwifiZkService) Create_client_once_task(mac, opts string) string {
	full_task_path := "/_zkqueues/" + mac + "/message"
	return svc.zkclient.SequenceAdd(full_task_path, opts)
}

func (svc *AwifiZkService) Destroy_client_once_task(mac string, taskid string) bool {

	full_task_id := "/_zkqueues/" + mac + "/" + taskid
	return svc.zkclient.Del(full_task_id)

}

func (svc *AwifiZkService) Create_service_schedule_task(taskid string, opts string) bool {

	path := "/awifi/task/master"
	full_task_id := path + "/" + taskid
	return svc.zkclient.Add(full_task_id, opts)

}

func (svc *AwifiZkService) Destroy_service_schedule_task(taskid string) bool {
	path := "/awifi/task/master"
	full_task_id := path + "/" + taskid

	return svc.zkclient.Del(full_task_id)

}

func (svc *AwifiZkService) Create_service_once_task(opts string) string {
	base := "/_zkqueues/gateway"
	return svc.zkclient.SequenceAdd(base, opts)
}

func (svc *AwifiZkService) Destroy_service_once_task(taskid string) bool {
	full_task_id := "/_zkqueues/gateway/" + taskid
	return svc.zkclient.Del(full_task_id)
}

func (svc *AwifiZkService) List_client_once_task(mac string) []string {
	base := "/_zkqueues/" + mac
	return svc.zkclient.Children(base)

}

func (svc *AwifiZkService) List_client_agent_task() []string {
	base := "/awifi/task/agent"
	return svc.zkclient.Children(base)

}
func (svc *AwifiZkService) List_service_task() []string {
	base := "/_zkqueues/gateway"
	return svc.zkclient.Children(base)
}

func (svc *AwifiZkService) List_service_schedule_task() []string {

	base := "/awifi/task/master"
	return svc.zkclient.Children(base)
}

func (svc *AwifiZkService) List_master_client() []string { 
        base := "/awifi/client/master"
	return  svc.zkclient.Children(base)
}

func (svc *AwifiZkService) Current_master_client() string {

     return svc.zkclient.Get("/awifi/client/master")
}

func (svc *AwifiZkService) List_gateway_clients() []string {
	base := "/awifi/client/gateway"
	return svc.zkclient.Children(base)

}

func (svc *AwifiZkService) List_control_clients() []string {
	base := "/awifi/client/control"
	return svc.zkclient.Children(base)

}

func (svc *AwifiZkService) List_agent_clients() []string {
	base := "/awifi/client/agent"
	return svc.zkclient.Children(base)

}

func (svc *AwifiZkService) List_heartbeat_clients() []string {
	base := "/awifi/client/heartbeat"
	return svc.zkclient.Children(base)

}

func (svc *AwifiZkService) List_notify_clients() []string {
	base := "/awifi/client/notify"
	return svc.zkclient.Children(base)

}
func (svc *AwifiZkService) List_collect_clients() []string {
	base := "/awifi/client/collect"
	return svc.zkclient.Children(base)

}

