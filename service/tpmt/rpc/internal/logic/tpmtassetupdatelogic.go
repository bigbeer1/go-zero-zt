package logic

import (
	"context"
	"errors"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"time"

	"tpmt-zt/service/tpmt/rpc/internal/svc"
	"tpmt-zt/service/tpmt/rpc/tpmtclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type TpmtAssetUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewTpmtAssetUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TpmtAssetUpdateLogic {
	return &TpmtAssetUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *TpmtAssetUpdateLogic) TpmtAssetUpdate(in *tpmtclient.TpmtAssetUpdateReq) (*tpmtclient.CommonResp, error) {

	res, err := l.svcCtx.TpmtAssetModel.FindOne(l.ctx, in.Id)
	if err != nil {
		if errors.Is(err, sqlc.ErrNotFound) {
			return nil, fmt.Errorf("TpmtAsset没有该ID: %v", in.Id)
		}
		return nil, err
	}

	// 资产类型
	if in.AssetType != 0 {
		res.AssetType = in.AssetType
	}
	// 资产编号
	if len(in.AssetCode) > 0 {
		res.AssetCode = in.AssetCode
	}
	// 资产名称
	if len(in.AssetName) > 0 {
		res.AssetName = in.AssetName
	}
	// 资产型号
	if len(in.AssetModel) > 0 {
		res.AssetModel = in.AssetModel
	}
	// 生产厂家
	if len(in.ManuFacturer) > 0 {
		res.ManuFacturer = in.ManuFacturer
	}
	// 电压
	if len(in.Voltage) > 0 {
		res.Voltage = in.Voltage
	}
	// 容量
	if len(in.Capacity) > 0 {
		res.Capacity = in.Capacity
	}

	res.UpdatedName.String = in.UpdatedName
	res.UpdatedName.Valid = true
	res.UpdatedAt.Time = time.Now()
	res.UpdatedAt.Valid = true

	err = l.svcCtx.TpmtAssetModel.Update(l.ctx, res)

	if err != nil {
		return nil, err
	}
	return &tpmtclient.CommonResp{}, nil
}
