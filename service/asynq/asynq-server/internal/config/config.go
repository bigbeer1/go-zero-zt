package config

import (
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/zrpc"
	"tpmt-zt/common/tdenginex"
)

type Config struct {
	service.ServiceConf

	// 实时数据保存时间
	RealDataTime int64

	Mysql struct {
		DataSource string
	}

	CacheRedis cache.CacheConf

	// 任务数
	Concurrency int

	// 自定义定时任务协程数
	ScheduledTasksLimit int64

	Tdengine tdenginex.TDengineConfig

	WebsocketRpc zrpc.RpcClientConf

	MqttSendRpc zrpc.RpcClientConf

	// 告警最多执行任务数
	AlarmLimit int64

	// 数据存储到时序数据库最高任务数
	DataSetLimit int64
}
