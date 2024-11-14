package schedule

import (
	"context"
	"errors"
	"fmt"
	"github.com/Masterminds/squirrel"
	"github.com/hibiken/asynq"
	"github.com/zeromicro/go-zero/core/logx"
	"time"
	"tpmt-zt/common/jsonx"
	"tpmt-zt/service/asynq/asynq-server/internal/logic/task"
	"tpmt-zt/service/asynq/asynq-server/internal/svc"
	"tpmt-zt/service/asynq/jobtype"
	"tpmt-zt/service/tpmt/model"
)

type SchedulerScheduledTasksHandler struct {
	svcCtx *svc.ServiceContext
}

func NewSchedulerScheduledTasksHandler(svcCtx *svc.ServiceContext) *SchedulerScheduledTasksHandler {
	return &SchedulerScheduledTasksHandler{
		svcCtx: svcCtx,
	}
}

func (l *SchedulerScheduledTasksHandler) ProcessTask(ctx context.Context, t *asynq.Task) error {

	// 获取所有定时任务数据
	countBuilder := l.svcCtx.TpmtScheduledTasksModel.CountBuilder("id")

	countBuilder = countBuilder.Where(squirrel.Eq{
		"state ": 1,
	})

	count, err := l.svcCtx.TpmtScheduledTasksModel.FindCount(ctx, countBuilder)
	if err != nil {
		return errors.New(err.Error())
	}

	whereBuilder := l.svcCtx.TpmtScheduledTasksModel.RowBuilder()
	whereBuilder = whereBuilder.OrderBy("created_at DESC, id DESC")

	whereBuilder = whereBuilder.Where(squirrel.Eq{
		"state ": 1,
	})

	all, err := l.svcCtx.TpmtScheduledTasksModel.FindList(ctx, whereBuilder, 1, count)
	if err != nil {
		return errors.New(err.Error())
	}

	var limit = make(chan struct{}, l.svcCtx.Config.ScheduledTasksLimit)

	now := time.Now()

	for _, item := range all {

		// 协程数控制
		limit <- struct{}{}
		// 创建30秒协程
		ctxA, _ := context.WithTimeout(context.Background(), 30*time.Second)
		// 协程写入日志
		go func(ctxA context.Context, item *model.TpmtScheduledTasks, now time.Time) {
			defer func() {
				<-limit
			}()

			var lastTimeUnixMilli int64
			// 获取当前时间时间戳毫秒

			timeNowUnixMilli := now.UnixMilli()

			// 构建redisKey
			redisKey := fmt.Sprintf("scheduled_tasks_id:%v", item.Id)

			l.svcCtx.Redis.GetCtx(ctxA, redisKey, &lastTimeUnixMilli)

			if lastTimeUnixMilli != 0 {
				// 距离上次任务间隔时间差
				lastDifferenceTime := timeNowUnixMilli - lastTimeUnixMilli

				// 任务间隔时间按秒
				intervalTimeUnixMilli := (item.IntervalTime) * 1000

				timeX := lastDifferenceTime - intervalTimeUnixMilli - 2000

				// 判断是否满足间隔时间 启动任务
				notStartBool := timeX < 0

				if notStartBool {
					return
				}

			}

			// 写入reids 执行记录
			err = l.svcCtx.Redis.SetCtx(ctxA, redisKey, timeNowUnixMilli)
			if err != nil {
				logx.Errorf("定时任务Key: %s存入失败！", redisKey)
				return
			}

			logx.Infof("定时任务触发")

			req := &task.ScheduledTasksReq{
				TaskType:  1,
				BeginTime: now,
				Data:      item,
			}

			dataBytes, _ := jsonx.Marshal(req)

			// 触发异步任务asynq
			l.svcCtx.AsynqClient.EnqueueContext(ctxA, asynq.NewTask(jobtype.ScheduledTasks, dataBytes))

		}(ctxA, item, now)
	}
	return nil
}
