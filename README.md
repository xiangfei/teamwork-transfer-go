# _teamwork-transfer-go_ # 
### centor transfer 
1. 边缘转发http 发送消息到kafka
2. 启动
   go run centor.go
3. 配置文件手动修改代码
4. 默认port 4567


### border transfer

1. 边缘转发http 发送消息到master
2. 启动
   go run border.go
3. 配置文件手动修改代码
4. 默认port 4568


### 配置文件
1. config.yaml


### 转发异常

- ruby 客户端手动处理


### 安装

1. 安装go
2. 下载依赖包

```
# 使用代理
export GO111MODULE=on
export GOPROXY=https://goproxy.io

```
3. go mod download

4. go build border.go

5. go build center.go
