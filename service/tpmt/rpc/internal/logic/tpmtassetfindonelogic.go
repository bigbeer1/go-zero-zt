package logic

import (
	"context"
	"errors"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlc"

	"tpmt-zt/service/tpmt/rpc/internal/svc"
	"tpmt-zt/service/tpmt/rpc/tpmtclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type TpmtAssetFindOneLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewTpmtAssetFindOneLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TpmtAssetFindOneLogic {
	return &TpmtAssetFindOneLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *TpmtAssetFindOneLogic) TpmtAssetFindOne(in *tpmtclient.TpmtAssetFindOneReq) (*tpmtclient.TpmtAssetFindOneResp, error) {

	res, err := l.svcCtx.TpmtAssetModel.FindOne(l.ctx, in.Id)
	if err != nil {
		if errors.Is(err, sqlc.ErrNotFound) {
			return nil, fmt.Errorf("TpmtAsset没有该ID:%v", in.Id)
		}
		return nil, err
	}

	return &tpmtclient.TpmtAssetFindOneResp{
		Id:           res.Id,                         //资产ID
		AssetType:    res.AssetType,                  //资产类型
		AssetCode:    res.AssetCode,                  //资产编号
		AssetName:    res.AssetName,                  //资产名称
		AssetModel:   res.AssetModel,                 //资产型号
		ManuFacturer: res.ManuFacturer,               //生产厂家
		Voltage:      res.Voltage,                    //电压
		Capacity:     res.Capacity,                   //容量
		CreatedAt:    res.CreatedAt.UnixMilli(),      //创建时间
		CreatedName:  res.CreatedName,                //创建人
		UpdatedAt:    res.UpdatedAt.Time.UnixMilli(), //更新时间
		UpdatedName:  res.UpdatedName.String,         //更新人
	}, nil
}
