package logic

import (
	"fmt"
	"github.com/hibiken/asynq"
	"github.com/zeromicro/go-zero/core/logx"
	"tpmt-zt/common/global"
	"tpmt-zt/service/asynq/jobtype"
)

// 定时实时数据存储到时序数据库
func (l *MqueueScheduler) schedulerDataSetTasksScheduler() {

	task := asynq.NewTask(jobtype.SchedulerDateSet, nil)
	// 从配置文件获取

	every := fmt.Sprintf("@every %vs", global.ComMemoryCycleTime)
	entryID, err := l.svcCtx.Scheduler.Register(every, task)
	if err != nil {
		logx.WithContext(l.ctx).Errorf("!!!MqueueSchedulerErr!!! ====> 【SchedulerAlarmTasks】 registered  err:%+v , task:%+v", err, task)
	}
	fmt.Printf("【SchedulerAlarmTasks】 registered an  entry: %q \n", entryID)

}
