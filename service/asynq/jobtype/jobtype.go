package jobtype

// SchedulerScheduledTasks 自定义定时任务
const SchedulerScheduledTasks = "schedule:scheduledTasks"

// SchedulerScheduledTasksFailureRecord 重试任务
const SchedulerScheduledTasksFailureRecord = "schedule:scheduledTasksFailureRecord"

// SchedulerAlarm 告警任务分发
const SchedulerAlarm = "schedule:schedulerAlarm"

// SchedulerDateSet 定时实时数据存储到时序数据库
const SchedulerDateSet = "schedule:schedulerDateSet"

// ScheduledTasks 自定义任务
const ScheduledTasks = "asyncTask:scheduledTasks"

// DataSet 数据流转到时序数据库中去
const DataSet = "asyncTask:dataSet"

// DataRedisSet 数据流转到Redis中去
const DataRedisSet = "asyncTask:dataRedisSet"

// GatewayState 网关日志流转
const GatewayStateOnline = "asyncTask:gatewayStateOnline"

// DeviceAlarm 告警状态判断
const DeviceAlarm = "asyncTask:deviceAlarm"
