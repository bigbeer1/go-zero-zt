package logic

import (
	"context"
	uuid "github.com/satori/go.uuid"
	"time"
	"tpmt-zt/service/tpmt/model"

	"tpmt-zt/service/tpmt/rpc/internal/svc"
	"tpmt-zt/service/tpmt/rpc/tpmtclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type TpmtAssetAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewTpmtAssetAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TpmtAssetAddLogic {
	return &TpmtAssetAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 资产
func (l *TpmtAssetAddLogic) TpmtAssetAdd(in *tpmtclient.TpmtAssetAddReq) (*tpmtclient.CommonResp, error) {
	_, err := l.svcCtx.TpmtAssetModel.Insert(l.ctx, &model.TpmtAsset{
		Id:           uuid.NewV4().String(), // ID
		CreatedAt:    time.Now(),            // 创建时间
		AssetType:    in.AssetType,          // 资产类型
		AssetCode:    in.AssetCode,          // 资产编号
		AssetName:    in.AssetName,          // 资产名称
		AssetModel:   in.AssetModel,         // 资产型号
		ManuFacturer: in.ManuFacturer,       // 生产厂家
		Voltage:      in.Voltage,            // 电压
		Capacity:     in.Capacity,           // 容量
		CreatedName:  in.CreatedName,        // 创建人
	})
	if err != nil {
		return nil, err
	}

	return &tpmtclient.CommonResp{}, nil
}
