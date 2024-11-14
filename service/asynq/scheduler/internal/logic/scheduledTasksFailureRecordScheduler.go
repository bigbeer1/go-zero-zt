package logic

import (
	"fmt"
	"github.com/hibiken/asynq"
	"github.com/zeromicro/go-zero/core/logx"
	"tpmt-zt/service/asynq/jobtype"
)

// 自定义定时任务
func (l *MqueueScheduler) scheduledTasksFailureRecordScheduler() {

	task := asynq.NewTask(jobtype.SchedulerScheduledTasksFailureRecord, nil)
	// 从配置文件获取

	every := fmt.Sprintf("@every %vs", l.svcCtx.Config.ScheduledTasksFailureRecordTime)
	entryID, err := l.svcCtx.Scheduler.Register(every, task)
	if err != nil {
		logx.WithContext(l.ctx).Errorf("!!!MqueueSchedulerErr!!! ====> 【SchedulerScheduledTasksFailureRecord】 registered  err:%+v , task:%+v", err, task)
	}
	fmt.Printf("【SchedulerScheduledTasksFailureRecord】 registered an  entry: %q \n", entryID)

}
