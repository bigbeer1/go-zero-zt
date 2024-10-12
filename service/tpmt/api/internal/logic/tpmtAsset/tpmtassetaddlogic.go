package tpmtAsset

import (
	"context"
	"tpmt-zt/common"
	"tpmt-zt/common/jwtx"
	"tpmt-zt/common/msg"
	"tpmt-zt/service/tpmt/rpc/tpmtclient"

	"tpmt-zt/service/tpmt/api/internal/svc"
	"tpmt-zt/service/tpmt/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type TpmtAssetAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTpmtAssetAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TpmtAssetAddLogic {
	return &TpmtAssetAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TpmtAssetAddLogic) TpmtAssetAdd(req *types.TpmtAssetAddRequest) (resp *types.Response, err error) {
	// 用户登录信息
	tokenData := jwtx.ParseToken(l.ctx)

	_, err = l.svcCtx.TpmtRpc.TpmtAssetAdd(l.ctx, &tpmtclient.TpmtAssetAddReq{
		AssetType:    req.AssetType,      // 资产类型
		AssetCode:    req.AssetCode,      // 资产编号
		AssetName:    req.AssetName,      // 资产名称
		AssetModel:   req.AssetModel,     // 资产型号
		ManuFacturer: req.ManuFacturer,   // 生产厂家
		Voltage:      req.Voltage,        // 电压
		Capacity:     req.Capacity,       // 容量
		CreatedName:  tokenData.NickName, // 创建人
	})
	if err != nil {
		return nil, common.NewDefaultError(err.Error())
	}
	return &types.Response{
		Code: 0,
		Msg:  msg.Success,
		Data: nil,
	}, nil
}
