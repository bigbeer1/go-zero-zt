package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ TpmtScheduledTasksModel = (*customTpmtScheduledTasksModel)(nil)

type (
	// TpmtScheduledTasksModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTpmtScheduledTasksModel.

	TpmtScheduledTasksModel interface {
		tpmtScheduledTasksModel
	}

	customTpmtScheduledTasksModel struct {
		*defaultTpmtScheduledTasksModel
	}
)

// NewTpmtScheduledTasksModel returns a model for the database table.
func NewTpmtScheduledTasksModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) TpmtScheduledTasksModel {
	return &customTpmtScheduledTasksModel{
		defaultTpmtScheduledTasksModel: newTpmtScheduledTasksModel(conn, c, opts...),
	}
}
