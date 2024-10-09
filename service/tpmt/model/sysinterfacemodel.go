package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ SysInterfaceModel = (*customSysInterfaceModel)(nil)

type (
	// SysInterfaceModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSysInterfaceModel.

	SysInterfaceModel interface {
		sysInterfaceModel
	}

	customSysInterfaceModel struct {
		*defaultSysInterfaceModel
	}
)

// NewSysInterfaceModel returns a model for the database table.
func NewSysInterfaceModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) SysInterfaceModel {
	return &customSysInterfaceModel{
		defaultSysInterfaceModel: newSysInterfaceModel(conn, c, opts...),
	}
}
