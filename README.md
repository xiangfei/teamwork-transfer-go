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

# 性能

```
[root@autotestweb devops-server]# ab -n 20000 -c 200  http://127.0.0.1:4567/testmsg
This is ApacheBench, Version 2.3 <$Revision: 1430300 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking 127.0.0.1 (be patient)
Completed 2000 requests
Completed 4000 requests
Completed 6000 requests
Completed 8000 requests
Completed 10000 requests
Completed 12000 requests
Completed 14000 requests
Completed 16000 requests
Completed 18000 requests
Completed 20000 requests
Finished 20000 requests


Server Software:        
Server Hostname:        127.0.0.1
Server Port:            4567

Document Path:          /testmsg
Document Length:        56 bytes

Concurrency Level:      200
Time taken for tests:   1.185 seconds
Complete requests:      20000
Failed requests:        0
Write errors:           0
Total transferred:      3580000 bytes
HTML transferred:       1120000 bytes
Requests per second:    16874.22 [#/sec] (mean)
Time per request:       11.852 [ms] (mean)
Time per request:       0.059 [ms] (mean, across all concurrent requests)
Transfer rate:          2949.69 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    3  18.8      2    1001
Processing:     0    7   6.4      5     212
Waiting:        0    5   5.6      4     212
Total:          0    9  19.8      7    1005

Percentage of the requests served within a certain time (ms)
  50%      7
  66%      9
  75%     11
  80%     12
  90%     16
  95%     19
  98%     23
  99%     35
 100%   1005 (longest request)

```
