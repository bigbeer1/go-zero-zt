package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ TpmtScheduledTasksFailureRecordModel = (*customTpmtScheduledTasksFailureRecordModel)(nil)

type (
	// TpmtScheduledTasksFailureRecordModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTpmtScheduledTasksFailureRecordModel.

	TpmtScheduledTasksFailureRecordModel interface {
		tpmtScheduledTasksFailureRecordModel
	}

	customTpmtScheduledTasksFailureRecordModel struct {
		*defaultTpmtScheduledTasksFailureRecordModel
	}
)

// NewTpmtScheduledTasksFailureRecordModel returns a model for the database table.
func NewTpmtScheduledTasksFailureRecordModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) TpmtScheduledTasksFailureRecordModel {
	return &customTpmtScheduledTasksFailureRecordModel{
		defaultTpmtScheduledTasksFailureRecordModel: newTpmtScheduledTasksFailureRecordModel(conn, c, opts...),
	}
}
