package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf

	TpmtRpc zrpc.RpcClientConf

	ArchiveRpc zrpc.RpcClientConf // 日志RPC

	Auth struct {
		AccessSecret string
		AccessExpire int64
	}
}
