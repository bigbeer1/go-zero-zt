package svc

import (
	"github.com/hibiken/asynq"
	asynqx "tpmt-zt/common/asynq"
	"tpmt-zt/service/asynq/scheduler/internal/config"
)

type ServiceContext struct {
	Config config.Config

	Scheduler *asynq.Scheduler
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:    c,
		Scheduler: asynqx.NewScheduler(c.CacheRedis),
	}
}
