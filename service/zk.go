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

func (service *AwifiZkService) Create_client_schedule_task(taskid string, mac string, opts string) bool {
	base := "/awifi/task/collect"
	task_path := base + "/" + mac
	full_task_id := task_path + "/" + taskid
	s := service.zkclient.Exist(full_task_id)
	if s {
		service.zkclient.Del( full_task_id)
	}
	return service.zkclient.Add(full_task_id, opts)

}

func (service *AwifiZkService) List_client_schedule_task(mac string) []string {

	base := "/awifi/task/collect"
	task_path := base + "/" + mac
	return service.zkclient.Children(task_path)
}

func (service *AwifiZkService) Destroy_client_schedule_task(taskid string, mac string) bool {

	base := "/awifi/task/collect"
	task_path := base + "/" + mac
	full_task_id := task_path + "/" + taskid
	s := service.zkclient.Exist( full_task_id)
	if s {
		return service.zkclient.Del( full_task_id)
	}
	return true
}

func (service *AwifiZkService) Kickcollect(mac string) bool {

	base := "/awifi/client/collect"
	task_path := base + "/" + mac
	return service.zkclient.Del(task_path)
}

func (service *AwifiZkService) Create_client_once_task(mac, opts string) string {
	full_task_path := "/_zkqueues/" + mac + "/message"
	return service.zkclient.SequenceAdd(full_task_path, opts)
}

func (service *AwifiZkService) Destroy_client_once_task(mac string, taskid string) bool {

	full_task_id := "/_zkqueues/" + mac + "/" + taskid
	return service.zkclient.Del(full_task_id)

}

func (service *AwifiZkService) Create_service_schedule_task(taskid string, opts string) bool {

	path := "/awifi/task/master"
	full_task_id := path + "/" + taskid
	return service.zkclient.Add(full_task_id, opts)

}

func (service *AwifiZkService) Destroy_service_schedule_task(taskid string) bool {
	path := "/awifi/task/master"
	full_task_id := path + "/" + taskid

	return service.zkclient.Del( full_task_id)

}

func (service *AwifiZkService) Create_service_once_task(opts string) string {
	base := "/_zkqueues/gateway"
	return service.zkclient.SequenceAdd(base, opts)
}

func (service *AwifiZkService) Destroy_service_once_task(taskid string) bool {
	full_task_id := "/_zkqueues/gateway/" + taskid
	return service.zkclient.Del( full_task_id)
}

func (service *AwifiZkService) List_client_once_task(mac string) []string {
	base := "/_zkqueues/" + mac
	return service.zkclient.Children( base)

}

func (service *AwifiZkService) List_client_agent_task() []string {
	base := "/awifi/task/agent"
	return service.zkclient.Children(base)

}
func (service *AwifiZkService) List_service_task() []string {
	base := "/_zkqueues/gateway"
	return service.zkclient.Children(base)
}

func (service *AwifiZkService) List_service_schedule_task() []string {

	base := "/awifi/task/master"
	return service.zkclient.Children(base)
}

