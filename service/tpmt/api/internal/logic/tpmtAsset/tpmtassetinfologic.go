package tpmtAsset

import (
	"context"
	"github.com/jinzhu/copier"
	"tpmt-zt/common"
	"tpmt-zt/common/msg"
	"tpmt-zt/service/tpmt/rpc/tpmtclient"

	"tpmt-zt/service/tpmt/api/internal/svc"
	"tpmt-zt/service/tpmt/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type TpmtAssetInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTpmtAssetInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TpmtAssetInfoLogic {
	return &TpmtAssetInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TpmtAssetInfoLogic) TpmtAssetInfo(req *types.TpmtAssetInfoRequest) (resp *types.Response, err error) {

	res, err := l.svcCtx.TpmtRpc.TpmtAssetFindOne(l.ctx, &tpmtclient.TpmtAssetFindOneReq{
		Id: req.Id, // 资产ID
	})
	if err != nil {
		return nil, common.NewDefaultError(err.Error())
	}

	var result TpmtAssetFindOneResp
	_ = copier.Copy(&result, res)

	return &types.Response{
		Code: 0,
		Msg:  msg.Success,
		Data: result,
	}, nil
}

type TpmtAssetFindOneResp struct {
	Id           string `json:"id"`            // 资产ID,
	AssetType    int64  `json:"asset_type"`    // 资产类型,
	AssetCode    string `json:"asset_code"`    // 资产编号,
	AssetName    string `json:"asset_name"`    // 资产名称,
	AssetModel   string `json:"asset_model"`   // 资产型号,
	ManuFacturer string `json:"manu_facturer"` // 生产厂家,
	Voltage      string `json:"voltage"`       // 电压,
	Capacity     string `json:"capacity"`      // 容量,
	CreatedAt    int64  `json:"created_at"`    // 创建时间,
	CreatedName  string `json:"created_name"`  // 创建人,
	UpdatedAt    int64  `json:"updated_at"`    // 更新时间,
	UpdatedName  string `json:"updated_name"`  // 更新人
}
