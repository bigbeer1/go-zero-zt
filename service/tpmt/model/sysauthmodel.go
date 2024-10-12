package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ SysAuthModel = (*customSysAuthModel)(nil)

type (
	// SysAuthModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSysAuthModel.

	SysAuthModel interface {
		sysAuthModel
	}

	customSysAuthModel struct {
		*defaultSysAuthModel
	}
)

// NewSysAuthModel returns a model for the database table.
func NewSysAuthModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) SysAuthModel {
	return &customSysAuthModel{
		defaultSysAuthModel: newSysAuthModel(conn, c, opts...),
	}
}
