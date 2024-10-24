package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/zrpc"
	"tpmt-zt/common/tdenginex"
)

type Config struct {
	zrpc.RpcServerConf

	Mysql struct {
		DataSource string
	}

	CacheRedis cache.CacheConf

	Tdengine tdenginex.TDengineConfig

	Salt string

	CAuth struct {
		AccessSecret string
		AccessExpire int64
	}
}
