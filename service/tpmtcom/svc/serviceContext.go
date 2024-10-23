package svc

import (
	"github.com/hibiken/asynq"
	"github.com/panjf2000/ants/v2"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"golang.org/x/time/rate"
	asynqx "tpmt-zt/common/asynq"
	"tpmt-zt/service/tpmt/model"
	"tpmt-zt/service/tpmtcom/config"
)

type ServiceContext struct {
	Config config.Config
	//字典
	SysDictModel model.SysDictModel

	//TPMT采集器
	TpmtGatewayModel model.TpmtGatewayModel
	//TPMT监测点
	TpmtMonitorPointModel model.TpmtMonitorPointModel

	// asynq客户端
	AsynqClient *asynq.Client

	// websocket最大任务数 限流器
	TaskMaxLimit *rate.Limiter

	// 数据协程池
	DataAntsPool *ants.PoolWithFunc
	// 网关数据协程池
	GateWayAntsPool *ants.PoolWithFunc
	// 通知内容协程池
	SendSocketAntsPool *ants.PoolWithFunc
	// 通知内容协程池带锁
	SendSocketNoLockAntsPool *ants.PoolWithFunc
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)

	return &ServiceContext{
		Config:                c,
		AsynqClient:           asynqx.NewAsynqClient(c.CacheRedis),
		SysDictModel:          model.NewSysDictModel(conn, c.CacheRedis),
		TpmtGatewayModel:      model.NewTpmtGatewayModel(conn, c.CacheRedis),
		TpmtMonitorPointModel: model.NewTpmtMonitorPointModel(conn, c.CacheRedis),
	}
}
