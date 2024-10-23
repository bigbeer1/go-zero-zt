package main

import (
	"flag"
	"fmt"
	"github.com/panjf2000/ants/v2"
	"github.com/valyala/fasthttp"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"golang.org/x/time/rate"
	"os"
	"runtime"
	"time"
	"tpmt-zt/common/global"
	"tpmt-zt/service/tpmtcom/com"
	"tpmt-zt/service/tpmtcom/config"
	"tpmt-zt/service/tpmtcom/logic"
	"tpmt-zt/service/tpmtcom/svc"
	"tpmt-zt/service/tpmtcom/util"
)

var configFile = flag.String("f", "etc/tpmt-com.yaml", "Specify the config file")

func main() {
	flag.Parse()
	var c config.Config

	conf.MustLoad(*configFile, &c)

	logx.DisableStat()
	if err := c.SetUp(); err != nil {
		panic(err)
	}

	cpuNum := runtime.NumCPU() //获得当前设备的cpu核心数
	fmt.Println("tpmtCom任务使用,cpu核心数:", cpuNum)
	runtime.GOMAXPROCS(cpuNum) //设置需要用到的cpu数量

	ctx := svc.NewServiceContext(c)

	// 开启自动获取modbus数据
	tpmtcom := com.NewComScheduler(ctx)
	go tpmtcom.Start()

	limit := rate.Every(time.Duration(c.TimeLimit) * time.Millisecond)
	ctx.TaskMaxLimit = rate.NewLimiter(limit, c.TaskMaxLimit)
	fmt.Println(fmt.Sprintf("限流器任务数:%v,速率:%v", c.TaskMaxLimit, c.TimeLimit))

	fmt.Println(fmt.Sprintf("RTU超时时间:%v毫秒", c.RtuTimeout))

	websocket := logic.NewCronScheduler(ctx)
	requestHandler := websocket.Register()

	server := fasthttp.Server{
		Name:        c.Name,
		Handler:     requestHandler,
		ReadTimeout: time.Minute,
	}

	ctx.DataAntsPool, _ = ants.NewPoolWithFunc(4000, func(req interface{}) {
		data, _ := req.(global.MonitorCacheUpData)
		tpmtcom.DataFlowRedis(&data)
	})

	ctx.GateWayAntsPool, _ = ants.NewPoolWithFunc(1000, func(req interface{}) {
		data, _ := req.(global.GateWayOnlineUpData)
		tpmtcom.GateWayOnlineDataFlow(&data)
	})

	ctx.SendSocketAntsPool, _ = ants.NewPoolWithFunc(10, func(req interface{}) {
		data, _ := req.(util.SendMessage)

		util.SendSocket(data)
	})

	ctx.SendSocketNoLockAntsPool, _ = ants.NewPoolWithFunc(1000, func(req interface{}) {
		data, _ := req.(util.SendMessage)
		util.SendSocketNoLock(data)
	})

	if err := server.ListenAndServe(fmt.Sprintf("%v:%v", c.Host, c.Port)); err != nil {
		logx.Errorf("!!!WebsocketError: !!!  run err:%+v", err)
		os.Exit(1)
	}

}
