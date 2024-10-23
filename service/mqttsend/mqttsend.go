package main

import (
	"flag"
	"fmt"
	zredis "github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"runtime"
	"time"
	"tpmt-zt/common/rpcserver"
	"tpmt-zt/service/mqttsend/internal/config"
	"tpmt-zt/service/mqttsend/internal/server"
	"tpmt-zt/service/mqttsend/internal/svc"
	"tpmt-zt/service/mqttsend/mqttsendclient"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/mqtt-send.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		mqttsendclient.RegisterMqttSendServer(grpcServer, server.NewMqttSendServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})

	cpuNum := runtime.NumCPU() //获得当前设备的cpu核心数
	fmt.Println("mqttSend使用,cpu核心数:", cpuNum)
	runtime.GOMAXPROCS(cpuNum) //设置需要用到的cpu数量
	defer s.Stop()

	// 设置日志输出 接口慢时间  rpc
	zrpc.SetServerSlowThreshold(time.Second * 2000)
	// redis
	zredis.SetSlowThreshold(time.Second * 2000)
	// sqlx
	sqlx.SetSlowThreshold(time.Second * 2000)

	//rpc log
	s.AddUnaryInterceptors(rpcserver.LoggerInterceptor)

	fmt.Println(c.Name, c.ListenOn)
	s.Start()
}
