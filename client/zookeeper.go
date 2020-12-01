package client

import (
	"fmt"
	"github.com/go-zookeeper/zk"
	"teamwork-transfer-go/config"
	"time"
        "sync"
)


var (
        singlezkInstance *AwifiZk
        zklock           = &sync.Mutex{}
)


type AwifiZk struct {
	zkurl      []string
	connection *zk.Conn
        username string
        password string 
}
 

func GetAwifiZkSingleton() *AwifiZk {

        if singlezkInstance == nil {
                zklock.Lock()
                defer zklock.Unlock()
                if singlezkInstance == nil {
                        singlezkInstance = initialize_zk()
                        return singlezkInstance
                }
        }
        return singlezkInstance


}


func initialize_zk() *AwifiZk {
	config := config.SingleConfigInstance()
	conn, _, err := zk.Connect(config.Zkurl, time.Second*5)
	if err != nil {
		panic(err)
	}
	s := config.Zkusername + ":" + config.Zkpassword
	//conn.AddAuth("digest", []byte("wifi2zk:awifi456"))
	conn.AddAuth("digest", []byte(s))
	//defer conn.Close()
	instance := &AwifiZk{
		connection: conn,
                zkurl: config.Zkurl,
                username: config.Zkusername,
                password: config.Zkpassword,
	}
	return instance

}

func (zk *AwifiZk) Close() {
	zk.connection.Close()
}

func (zk *AwifiZk) Children(path string) []string {
	children, _, _, err := zk.connection.ChildrenW(path)
	if err != nil {
		panic(err)
	}

	return children
}

func (zk *AwifiZk) Del(path string) bool { 
   _, sate, _ := zk.connection.Get(path)
        err := zk.connection.Delete(path, sate.Version)
        if err != nil {
                return false
        }
        return true


}

func  (zk *AwifiZk) Get(path string) string{
    data, _, err := zk.connection.Get(path)
        if err != nil {
                fmt.Printf("查询%s失败, err: %v\n", path, err)
                return ""
        }
        fmt.Printf("%s 的值为 %s\n", path, string(data))
        return string(data)


}

func (zk *AwifiZk) Exist(path string) bool {
        _, _, err := zk.connection.Get(path)
        if err != nil {
                return false
        }
        return true


}

func (z *AwifiZk) Add(path string , data string) bool  {

        // flags有4种取值：
        // 0:永久，除非手动删除
        // zk.FlagEphemeral = 1:短暂，session断开则该节点也被删除
        // zk.FlagSequence  = 2:会自动在节点后面添加序号
        // 3:Ephemeral和Sequence，即，短暂且自动添加序号
        var flags int32 = 0
        // 获取访问控制权限
        acls := zk.WorldACL(zk.PermAll)
        s, err := z.connection.Create(path, []byte(data), flags, acls)
        if err != nil {
                fmt.Printf("创建失败:   path %s  %v\n  ", path, err)
                return false
        }
        fmt.Printf("创建: %s 成功", s)
        return true

}


func(z *AwifiZk) SequenceAdd(path string, data string) string {

        var flags int32 = 2
        acls := zk.WorldACL(zk.PermAll)
        s, err := z.connection.Create(path, []byte(data), flags, acls)
        if err != nil {
                fmt.Printf("创建失败:   path %s  %v\n  ", path, err)
                panic(err)
        }
        fmt.Printf("创建: %s 成功", s)
        return s


}


