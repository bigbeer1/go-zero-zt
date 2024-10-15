package svc

import (
	"database/sql"
	"tpmt-zt/common/tdenginex"
	"tpmt-zt/service/archive/rpc/internal/config"
)

type ServiceContext struct {
	Config config.Config

	// Td连接
	Taos *sql.DB
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Taos:   tdenginex.NewTDengineManager(c.Tdengine),
	}
}
