package svc

import (
	"errors"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/syncx"
	"golang.org/x/time/rate"
	"time"
	"tpmt-zt/service/tpmt/model"
	"tpmt-zt/service/websocket/internal/config"
)

type ServiceContext struct {
	Config config.Config
	// 用户在线时间保存
	OnlineTime time.Duration
	Redis      cache.Cache
	// 监测点
	TpmtMonitorPointModel model.TpmtMonitorPointModel

	// 最大任务数 限流器
	TaskMaxLimit *rate.Limiter
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:                c,
		Redis:                 cache.New(c.CacheRedis, syncx.NewSingleFlight(), cache.NewStat("any"), errors.New("placeholder")),
		TpmtMonitorPointModel: model.NewTpmtMonitorPointModel(conn, c.CacheRedis),
	}
}
