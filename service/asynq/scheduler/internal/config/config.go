package config

import (
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/core/stores/cache"
)

type Config struct {
	service.ServiceConf

	Mysql struct {
		DataSource string
	}

	CacheRedis                      cache.CacheConf
	SchedulerAlarmTasksTime         int64
	ScheduledTasksTime              int64
	ScheduledTasksFailureRecordTime int64
}
