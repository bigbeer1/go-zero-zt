.PHONY: ptm
ptm: linux_env authentication mqtt-send archive  tpmt asynq-server  tpmt-com tpmt-websocket


# linux 环境变量
linux_env:
	set GOARCH=amd64
	go env -w GOARCH=amd64
	set GOOS=linux
	go env -w GOOS=linux


# 系统服务
authentication:
	go build -o deploy/golang/out/authentication/authentication-rpc service/authentication/authentication.go


# script 底层重启RPC
mqtt-send:
	go build -o deploy/golang/out/mqttSend/mqtt-send service/mqttsend/mqttsend.go

# tpmt微服务
tpmt:
	go build -o deploy/golang/out/tpmt/rpc/tpmt-rpc service/tpmt/rpc/tpmt.go
	go build -o deploy/golang/out/tpmt/api/tpmt-api service/tpmt/api/tpmt.go

tpmt-com:
	go build -o deploy/golang/out/tpmtcom/tpmt-com service/tpmtcom/tpmtcom.go

tpmt-websocket:
	go build -o deploy/golang/out/websocket/websocket service/websocket/websocket.go


# 定时任务/异步任务服务
asynq-server:
	go build -o deploy/golang/out/other/asynq-server service/asynq/asynq-server/asynq-server.go
	go build -o deploy/golang/out/other/scheduler service/asynq/scheduler/scheduler.go


# script 底层重启RPC
mqtt-send:
	go build -o deploy/golang/out/mqttSend/mqtt-send service/mqttsend/mqttsend.go