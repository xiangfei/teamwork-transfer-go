package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"sync"
)

var (
	singleInstance *conf
	lock           = &sync.Mutex{}
)

type conf struct {
	Zkusername      string   `yaml:"zkusername"`
	Zkpassword      string   `yaml:"zkpassword"`
	Zkurl           []string `yaml: zkurl`
	Center_http_url []string `yaml: "center_http_url"`
	Kafka_url       []string `yaml: "kafka_url"`
}

func (c *conf) getConf() *conf {

	yamlFile, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return c
}

func newConfig() *conf {
	var c conf
	return c.getConf()
}

func SingleConfigInstance() *conf {

	if singleInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if singleInstance == nil {
			singleInstance = newConfig()
			return singleInstance
		}
	}
	return singleInstance

}

func init() {
	_ = SingleConfigInstance()

}

//func main() {
//	v := SingleConfigInstance()
//	fmt.Println(v)
//}

