package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"os"
	"tpmt-zt/service/asynq/scheduler/internal/config"
	"tpmt-zt/service/asynq/scheduler/internal/logic"
	"tpmt-zt/service/asynq/scheduler/internal/svc"
)

var configFile = flag.String("f", "etc/scheduler.yaml", "Specify the config file")

func main() {
	flag.Parse()
	var c config.Config

	conf.MustLoad(*configFile, &c)

	logx.DisableStat()
	// log、prometheus、trace、metricsUrl.
	if err := c.SetUp(); err != nil {
		panic(err)
	}

	svcContext := svc.NewServiceContext(c)
	ctx := context.Background()

	mqueueScheduler := logic.NewCronScheduler(ctx, svcContext)
	mqueueScheduler.Register()

	println(fmt.Sprintf("网关告警检测时间:%v秒", c.SchedulerAlarmTasksTime))

	println(fmt.Sprintf("自定义定时任务执行查询发布时间:%v秒", c.ScheduledTasksTime))

	if err := svcContext.Scheduler.Run(); err != nil {
		logx.Errorf("!!!MqueueSchedulerErr!!!  run err:%+v", err)
		os.Exit(1)
	}

}
