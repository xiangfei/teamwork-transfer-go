package service

import (
	"bytes"
        //"fmt"
	"encoding/json"
	"net/http"
	"strings"
	"teamwork-transfer-go/client"
)

type AwifiHttpService struct {
	httpclient *client.AwifiHttp
}


func NewHttpServiceFromClient(zkclient *client.AwifiHttp) *AwifiHttpService {

	return &AwifiHttpService{zkclient}

}

func (service *AwifiHttpService) JsonPost(path string, jsonbytes []byte) bool {
	// full_url :=  service.httpclient.Baseurl() + path
	var build strings.Builder
	build.WriteString(service.httpclient.Baseurl())
	build.WriteString(path)
	full_url := build.String()
	resp, err := http.Post(full_url, "application/json", bytes.NewBuffer(jsonbytes))
	if err != nil {
		return false
	} else {
		var res map[string]interface{}
		json.NewDecoder(resp.Body).Decode(&res)
                //fmt.Println(res)
                //fmt.Println(res["code"] ) 
                //fmt.Println(reflect.TypeOf(res["code"]))
                return res["code"]  == float64(200)
                //return true
	}

}

