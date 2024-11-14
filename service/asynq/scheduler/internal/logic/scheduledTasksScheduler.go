package logic

import (
	"fmt"
	"github.com/hibiken/asynq"
	"github.com/zeromicro/go-zero/core/logx"
	"tpmt-zt/service/asynq/jobtype"
)

// 自定义定时任务
func (l *MqueueScheduler) scheduledTasksScheduler() {

	task := asynq.NewTask(jobtype.SchedulerScheduledTasks, nil)
	// 从配置文件获取

	every := fmt.Sprintf("@every %vs", l.svcCtx.Config.ScheduledTasksTime)
	entryID, err := l.svcCtx.Scheduler.Register(every, task)
	if err != nil {
		logx.WithContext(l.ctx).Errorf("!!!MqueueSchedulerErr!!! ====> 【SchedulerScheduledTasks】 registered  err:%+v , task:%+v", err, task)
	}
	fmt.Printf("【SchedulerScheduledTasks】 registered an  entry: %q \n", entryID)

}
