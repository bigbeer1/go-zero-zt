package schedule

import (
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/hibiken/asynq"
	"github.com/zeromicro/go-zero/core/jsonx"
	"github.com/zeromicro/go-zero/core/logx"
	"time"
	"tpmt-zt/service/asynq/asynq-server/internal/svc"
	"tpmt-zt/service/asynq/jobtype"
	"tpmt-zt/service/tpmt/model"
)

type SchedulerAlarmTasksHandler struct {
	svcCtx *svc.ServiceContext
}

func NewSchedulerAlarmTasksHandler(svcCtx *svc.ServiceContext) *SchedulerAlarmTasksHandler {
	return &SchedulerAlarmTasksHandler{
		svcCtx: svcCtx,
	}
}

func (l *SchedulerAlarmTasksHandler) ProcessTask(ctx context.Context, t *asynq.Task) error {

	whereBuilder := l.svcCtx.TpmtMonitorPointModel.RowBuilder()

	whereBuilder = whereBuilder.OrderBy("created_at DESC, id DESC")

	// 告警持续时间
	whereBuilder = whereBuilder.Where(squirrel.NotEq{
		"alarm_duration ": 0,
	})

	all, err := l.svcCtx.TpmtMonitorPointModel.FindAll(ctx, whereBuilder)
	if err != nil {
		logx.Errorf("告警:查询监测点错误:%s", err)
		return nil
	}

	var limit = make(chan struct{}, l.svcCtx.Config.AlarmLimit)

	for _, item := range all {
		// 未知类型错误数据
		if item.PointCategory == 2 || item.PointCategory == 3 {
			// 协程数控制
			limit <- struct{}{}
			// 创建30秒协程写入日志
			ctxA, _ := context.WithTimeout(context.Background(), 30*time.Second)
			// 协程写入日志
			go func(ctxA context.Context, item *model.TpmtMonitorPoint) {
				defer func() {
					<-limit
				}()
				dataBytes, _ := jsonx.Marshal(item)
				// 触发异步任务asynq
				l.svcCtx.AsynqClient.EnqueueContext(ctxA, asynq.NewTask(jobtype.DeviceAlarm, dataBytes))
			}(ctxA, item)

		} else {
			continue
		}

	}

	return nil

}
