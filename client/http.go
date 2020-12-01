package client

import (
	"math/rand"
	"sync"
	"teamwork-transfer-go/config"
)

var (
	singleHttpInstance *AwifiHttp
	httplock           = &sync.Mutex{}
)

type AwifiHttp struct {
	base_url []string
}

func GetAwifiHttpSingleton() *AwifiHttp {

	if singleHttpInstance == nil {
		httplock.Lock()
		defer httplock.Unlock()
		if singleHttpInstance == nil {
			singleHttpInstance = initialize_http()
			return singleHttpInstance
		}
	}
	return singleHttpInstance

}

func initialize_http() *AwifiHttp {
	baseconfig := config.SingleConfigInstance()

	instance := &AwifiHttp{
		base_url: baseconfig.Center_http_url,
	}
	return instance

}

func (http *AwifiHttp) Baseurl() string {

	randomIndex := rand.Intn(len(http.base_url))
	randomelemtn := http.base_url[randomIndex]
	return randomelemtn
}

