package schedule

import (
	"context"
	"errors"
	"fmt"
	"github.com/hibiken/asynq"
	"github.com/zeromicro/go-zero/core/logx"
	"time"
	"tpmt-zt/common/jsonx"
	"tpmt-zt/service/asynq/asynq-server/internal/logic/task"
	"tpmt-zt/service/asynq/asynq-server/internal/svc"
	"tpmt-zt/service/asynq/jobtype"
	"tpmt-zt/service/tpmt/model"
)

type SchedulerScheduledTasksFailureRecordHandler struct {
	svcCtx *svc.ServiceContext
}

func NewSchedulerScheduledTasksFailureRecordHandler(svcCtx *svc.ServiceContext) *SchedulerScheduledTasksFailureRecordHandler {
	return &SchedulerScheduledTasksFailureRecordHandler{
		svcCtx: svcCtx,
	}
}

func (l *SchedulerScheduledTasksFailureRecordHandler) ProcessTask(ctx context.Context, t *asynq.Task) error {

	// 获取所有重试任务数据
	countBuilder := l.svcCtx.TpmtScheduledTasksFailureRecordModel.CountBuilder("id")

	count, err := l.svcCtx.TpmtScheduledTasksFailureRecordModel.FindCount(ctx, countBuilder)
	if err != nil {
		return errors.New(err.Error())
	}

	whereBuilder := l.svcCtx.TpmtScheduledTasksFailureRecordModel.RowBuilder()
	whereBuilder = whereBuilder.OrderBy("created_at DESC, id DESC")

	all, err := l.svcCtx.TpmtScheduledTasksFailureRecordModel.FindList(ctx, whereBuilder, 1, count)
	if err != nil {
		return errors.New(err.Error())
	}

	var limit = make(chan struct{}, l.svcCtx.Config.ScheduledTasksLimit)
	// 获取当前时间时间戳毫秒
	now := time.Now()

	for _, item := range all {

		// 协程数控制
		limit <- struct{}{}
		// 创建30秒协程
		ctxA, _ := context.WithTimeout(context.Background(), 30*time.Second)
		// 协程写入日志
		go func(ctxA context.Context, item *model.TpmtScheduledTasksFailureRecord, now time.Time) {
			defer func() {
				<-limit
			}()

			// 构建redisKey
			redisKey := fmt.Sprintf("scheduled_tasks_failure_record_id:%v", item.Id)

			if item.ErrorOrder <= item.FailOrder {
				// 删除任务
				err = l.svcCtx.TpmtScheduledTasksFailureRecordModel.Delete(ctxA, item.Id)
				if err != nil {
					logx.Errorf("删除重试任务失败Id:", item.Id)
				}

				logx.Infof("重试任务失败次数用完")
				return
			}

			var lastTimeUnixMilli int64

			timeNowUnixMilli := now.UnixMilli()

			l.svcCtx.Redis.GetCtx(ctxA, redisKey, &lastTimeUnixMilli)

			if lastTimeUnixMilli != 0 {
				// 距离上次任务间隔时间差
				lastDifferenceTime := timeNowUnixMilli - lastTimeUnixMilli

				// 任务间隔时间按秒
				intervalTimeUnixMilli := (item.FailIntervalTime) * 1000

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
				logx.Errorf("重试任务Key: %s存入失败！", redisKey)
				return
			}

			item.FailOrder = item.FailOrder + 1

			// 增加执行次数
			err = l.svcCtx.TpmtScheduledTasksFailureRecordModel.Update(ctxA, item)
			if err != nil {
				logx.Errorf("重试任务更新次数失败Id:", item.Id)
				return
			}

			logx.Infof("重试任务触发")

			req := &task.ScheduledTasksReq{
				TaskType:  2,
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
