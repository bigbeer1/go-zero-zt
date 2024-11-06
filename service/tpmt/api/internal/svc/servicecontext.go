package svc

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
	"tpmt-zt/service/archive/rpc/archive"
	"tpmt-zt/service/authentication/authentication"
	"tpmt-zt/service/tpmt/api/internal/config"
	"tpmt-zt/service/tpmt/api/internal/middleware"
	"tpmt-zt/service/tpmt/rpc/tpmt"
)

type ServiceContext struct {
	Config config.Config

	TpmtRpc tpmt.Tpmt

	// 日志服务
	ArchiveRpc archive.Archive

	// 鉴权RPC
	AuthenticationRpc authentication.Authentication

	// 鉴权中间件
	CheckAuth rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:            c,
		TpmtRpc:           tpmt.NewTpmt(zrpc.MustNewClient(c.TpmtRpc)),
		ArchiveRpc:        archive.NewArchive(zrpc.MustNewClient(c.ArchiveRpc)),
		AuthenticationRpc: authentication.NewAuthentication(zrpc.MustNewClient(c.AuthenticationRpc)),
		CheckAuth: middleware.NewCheckAuthMiddleware(authentication.NewAuthentication(zrpc.MustNewClient(c.AuthenticationRpc)),
			archive.NewArchive(zrpc.MustNewClient(c.ArchiveRpc)), c.Auth.AccessSecret).Handle,
	}
}
