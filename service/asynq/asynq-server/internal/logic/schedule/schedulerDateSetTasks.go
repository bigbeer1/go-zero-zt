package schedule

import (
	"context"
	"github.com/hibiken/asynq"
	"github.com/zeromicro/go-zero/core/jsonx"
	"github.com/zeromicro/go-zero/core/logx"
	"time"
	"tpmt-zt/service/asynq/asynq-server/internal/svc"
	"tpmt-zt/service/asynq/jobtype"
	"tpmt-zt/service/tpmt/model"
)

type SchedulerDateSetTasksHandler struct {
	svcCtx *svc.ServiceContext
}

func NewSchedulerDateSetTasksHandler(svcCtx *svc.ServiceContext) *SchedulerDateSetTasksHandler {
	return &SchedulerDateSetTasksHandler{
		svcCtx: svcCtx,
	}
}

// SchedulerDateSet 定时实时数据存储到时序数据库
func (l *SchedulerDateSetTasksHandler) ProcessTask(ctx context.Context, t *asynq.Task) error {

	whereBuilder := l.svcCtx.TpmtMonitorPointModel.RowBuilder()

	whereBuilder = whereBuilder.OrderBy("created_at DESC, id DESC")

	all, err := l.svcCtx.TpmtMonitorPointModel.FindAll(ctx, whereBuilder)
	if err != nil {
		logx.Errorf("数据存储:查询监测点错误:%s", err)
		return nil
	}

	var limit = make(chan struct{}, l.svcCtx.Config.DataSetLimit)

	for _, item := range all {
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
			l.svcCtx.AsynqClient.EnqueueContext(ctxA, asynq.NewTask(jobtype.DataSet, dataBytes))
		}(ctxA, item)

	}

	return nil

}
