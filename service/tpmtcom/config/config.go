package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf

	Mysql struct {
		DataSource string
	}

	CAuth struct {
		AccessSecret string
		AccessExpire int64
	}

	TaskMaxLimit int // 限流数量

	TimeLimit int //  限流时间

	Mock bool

	CacheRedis cache.CacheConf

	RtuTimeout int //  rtu超时时间
}
