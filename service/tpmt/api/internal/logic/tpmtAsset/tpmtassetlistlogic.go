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

type TpmtAssetListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTpmtAssetListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TpmtAssetListLogic {
	return &TpmtAssetListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TpmtAssetListLogic) TpmtAssetList(req *types.TpmtAssetListRequest) (resp *types.Response, err error) {

	all, err := l.svcCtx.TpmtRpc.TpmtAssetList(l.ctx, &tpmtclient.TpmtAssetListReq{
		Current:      req.Current,      // 页码
		PageSize:     req.PageSize,     // 页数
		AssetType:    req.AssetType,    // 资产类型
		AssetCode:    req.AssetCode,    // 资产编号
		AssetName:    req.AssetName,    // 资产名称
		AssetModel:   req.AssetModel,   // 资产型号
		ManuFacturer: req.ManuFacturer, // 生产厂家
		Voltage:      req.Voltage,      // 电压
		Capacity:     req.Capacity,     // 容量
	})
	if err != nil {
		return nil, common.NewDefaultError(err.Error())
	}

	var result TpmtAssetListResp
	_ = copier.Copy(&result, all)

	return &types.Response{
		Code: 0,
		Msg:  msg.Success,
		Data: result,
	}, nil
}

type TpmtAssetListResp struct {
	Total int64                `json:"total"`
	List  []*TpmtAssetDataList `json:"list"`
}

type TpmtAssetDataList struct {
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
