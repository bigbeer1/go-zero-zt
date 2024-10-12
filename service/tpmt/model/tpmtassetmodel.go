package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ TpmtAssetModel = (*customTpmtAssetModel)(nil)

type (
	// TpmtAssetModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTpmtAssetModel.

	TpmtAssetModel interface {
		tpmtAssetModel
	}

	customTpmtAssetModel struct {
		*defaultTpmtAssetModel
	}
)

// NewTpmtAssetModel returns a model for the database table.
func NewTpmtAssetModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) TpmtAssetModel {
	return &customTpmtAssetModel{
		defaultTpmtAssetModel: newTpmtAssetModel(conn, c, opts...),
	}
}
