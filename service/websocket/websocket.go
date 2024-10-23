package main

import (
	"flag"
	"fmt"
	"github.com/valyala/fasthttp"
	zredis "github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"golang.org/x/time/rate"
	"runtime"
	"time"
	"tpmt-zt/common/rpcserver"
	"tpmt-zt/service/websocket/internal/config"
	"tpmt-zt/service/websocket/internal/logic"
	"tpmt-zt/service/websocket/internal/server"
	"tpmt-zt/service/websocket/internal/svc"
	"tpmt-zt/service/websocket/websocketclient"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/websocket.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	cpuNum := runtime.NumCPU() //获得当前设备的cpu核心数
	fmt.Println("websocket任务使用,cpu核心数:", cpuNum)
	runtime.GOMAXPROCS(cpuNum) //设置需要用到的cpu数量

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		websocketclient.RegisterWebsocketServer(grpcServer, server.NewWebsocketServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting websokcet at %s:%s...\n", c.WebsocketHost, c.WebsocketPort)
	fmt.Println(fmt.Sprintf("用户在线状态存放时间：%v 分钟", c.OnlineTime))
	ctx.OnlineTime = time.Duration(c.OnlineTime) * time.Minute

	limit := rate.Every(time.Duration(c.TimeLimit) * time.Millisecond)
	ctx.TaskMaxLimit = rate.NewLimiter(limit, c.TaskMaxLimit)
	fmt.Println(fmt.Sprintf("限流器任务数:%v,速率:%v", c.TaskMaxLimit, c.TimeLimit))

	websocket := logic.NewCronScheduler(ctx)
	requestHandler := websocket.Register()

	server := fasthttp.Server{
		Name:        c.Name,
		Handler:     requestHandler,
		ReadTimeout: time.Minute,
	}

	go server.ListenAndServe(fmt.Sprintf("%v:%v", c.WebsocketHost, c.WebsocketPort))

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
