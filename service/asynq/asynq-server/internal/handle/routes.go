package handle

import (
	"context"
	"github.com/hibiken/asynq"
	"tpmt-zt/service/asynq/asynq-server/internal/logic/alarm"
	"tpmt-zt/service/asynq/asynq-server/internal/logic/schedule"
	"tpmt-zt/service/asynq/asynq-server/internal/logic/store"
	"tpmt-zt/service/asynq/asynq-server/internal/svc"
	"tpmt-zt/service/asynq/jobtype"
)

type CronJob struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCronJob(ctx context.Context, svcCtx *svc.ServiceContext) *CronJob {
	return &CronJob{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// register asynq-server
func (l *CronJob) Register() *asynq.ServeMux {

	mux := asynq.NewServeMux()

	// 告警任务分发
	mux.Handle(jobtype.SchedulerAlarm, schedule.NewSchedulerAlarmTasksHandler(l.svcCtx))

	// 数据存储分发
	mux.Handle(jobtype.SchedulerDateSet, schedule.NewSchedulerDateSetTasksHandler(l.svcCtx))
	
	// 数据流程任务存储任务  时序数据库
	mux.Handle(jobtype.DataSet, store.NewDataSetHandler(l.svcCtx))

	// 数据存储到redis中去
	mux.Handle(jobtype.DataRedisSet, store.NewDataRedisSetHandler(l.svcCtx))

	// 网关状态存储任务
	mux.Handle(jobtype.GatewayStateOnline, store.NewGatewayStateHandler(l.svcCtx))

	// DeviceAlarm 告警状态判断
	mux.Handle(jobtype.DeviceAlarm, alarm.NewDeviceAlarmHandler(l.svcCtx))

	return mux
}
