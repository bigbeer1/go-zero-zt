package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf

	WebsocketHost string
	WebsocketPort string

	CAuth struct {
		AccessSecret string
		AccessExpire int64
	}

	CacheRedis cache.CacheConf

	// 在线数据redis存放有效期
	OnlineTime int64

	Mysql struct {
		DataSource string
	}

	RealTimeLimit int64

	TaskMaxLimit int

	TimeLimit int
}
