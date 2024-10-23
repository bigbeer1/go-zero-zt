package logic

import (
	"fmt"
	"github.com/hibiken/asynq"
	"github.com/zeromicro/go-zero/core/logx"
	"tpmt-zt/service/asynq/jobtype"
)

// 定时告警数据查询
func (l *MqueueScheduler) schedulerAlarmTasksScheduler() {

	task := asynq.NewTask(jobtype.SchedulerAlarm, nil)
	// 从配置文件获取

	every := fmt.Sprintf("@every %vs", l.svcCtx.Config.SchedulerAlarmTasksTime)
	entryID, err := l.svcCtx.Scheduler.Register(every, task)
	if err != nil {
		logx.WithContext(l.ctx).Errorf("!!!MqueueSchedulerErr!!! ====> 【SchedulerAlarmTasks】 registered  err:%+v , task:%+v", err, task)
	}
	fmt.Printf("【SchedulerAlarmTasks】 registered an  entry: %q \n", entryID)

}
