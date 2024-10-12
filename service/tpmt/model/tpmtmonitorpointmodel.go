package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ TpmtMonitorPointModel = (*customTpmtMonitorPointModel)(nil)

type (
	// TpmtMonitorPointModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTpmtMonitorPointModel.

	TpmtMonitorPointModel interface {
		tpmtMonitorPointModel
	}

	customTpmtMonitorPointModel struct {
		*defaultTpmtMonitorPointModel
	}
)

// NewTpmtMonitorPointModel returns a model for the database table.
func NewTpmtMonitorPointModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) TpmtMonitorPointModel {
	return &customTpmtMonitorPointModel{
		defaultTpmtMonitorPointModel: newTpmtMonitorPointModel(conn, c, opts...),
	}
}
