package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	zredis "github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
	"os"
	"runtime"
	"time"
	"tpmt-zt/common/global"
	"tpmt-zt/service/asynq/asynq-server/internal/config"
	"tpmt-zt/service/asynq/asynq-server/internal/handle"
	"tpmt-zt/service/asynq/asynq-server/internal/svc"
)

var configFile = flag.String("f", "etc/asynq-server.yaml", "Specify the config file")

func main() {
	flag.Parse()
	var c config.Config
	// 小时转换成Unix
	conf.MustLoad(*configFile, &c)

	// log、prometheus、trace、metricsUrl
	if err := c.SetUp(); err != nil {
		panic(err)
	}

	logx.DisableStat()

	svcContext := svc.NewServiceContext(c)

	fmt.Println(fmt.Sprintf("实时数据redis存储时长：%v 分钟", c.RealDataTime))
	svcContext.RealDataTime = time.Duration(c.RealDataTime) * time.Minute
	ctx := context.Background()
	cronJob := handle.NewCronJob(ctx, svcContext)
	mux := cronJob.Register()

	// 设置日志输出 接口慢时间  rpc
	zrpc.SetServerSlowThreshold(time.Second * 2000)
	// redis
	zredis.SetSlowThreshold(time.Second * 2000)
	// sqlx
	sqlx.SetSlowThreshold(time.Second * 2000)

	fmt.Println(fmt.Sprintf("ansync并发线程数：%v ", c.Concurrency))

	cpuNum := runtime.NumCPU() //获得当前设备的cpu核心数
	fmt.Println("ansync任务使用,cpu核心数:", cpuNum)
	runtime.GOMAXPROCS(cpuNum) //设置需要用到的cpu数量

	fmt.Println("启动时间:", time.Now().In(global.ShangHaiTime).Format(time.RFC3339Nano))

	if err := svcContext.AsynqServer.Run(mux); err != nil {
		logx.Errorf("!!!CronJobErr!!! run err:%+v", err)
		os.Exit(1)
	}
}
