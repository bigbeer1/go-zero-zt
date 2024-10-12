package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ TpmtGatewayModel = (*customTpmtGatewayModel)(nil)

type (
	// TpmtGatewayModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTpmtGatewayModel.

	TpmtGatewayModel interface {
		tpmtGatewayModel
	}

	customTpmtGatewayModel struct {
		*defaultTpmtGatewayModel
	}
)

// NewTpmtGatewayModel returns a model for the database table.
func NewTpmtGatewayModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) TpmtGatewayModel {
	return &customTpmtGatewayModel{
		defaultTpmtGatewayModel: newTpmtGatewayModel(conn, c, opts...),
	}
}
