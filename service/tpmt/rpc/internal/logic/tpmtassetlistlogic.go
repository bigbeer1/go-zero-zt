package logic

import (
	"context"
	"github.com/Masterminds/squirrel"

	"tpmt-zt/service/tpmt/rpc/internal/svc"
	"tpmt-zt/service/tpmt/rpc/tpmtclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type TpmtAssetListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewTpmtAssetListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TpmtAssetListLogic {
	return &TpmtAssetListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *TpmtAssetListLogic) TpmtAssetList(in *tpmtclient.TpmtAssetListReq) (*tpmtclient.TpmtAssetListResp, error) {

	whereBuilder := l.svcCtx.TpmtAssetModel.RowBuilder()

	whereBuilder = whereBuilder.OrderBy("created_at DESC, id DESC")

	// 资产类型
	if in.AssetType != 99 {
		whereBuilder = whereBuilder.Where(squirrel.Eq{
			"asset_type ": in.AssetType,
		})
	}
	// 资产编号
	if len(in.AssetCode) > 0 {
		whereBuilder = whereBuilder.Where(squirrel.Like{
			"asset_code ": "%" + in.AssetCode + "%",
		})
	}
	// 资产名称
	if len(in.AssetName) > 0 {
		whereBuilder = whereBuilder.Where(squirrel.Like{
			"asset_name ": "%" + in.AssetName + "%",
		})
	}
	// 资产型号
	if len(in.AssetModel) > 0 {
		whereBuilder = whereBuilder.Where(squirrel.Like{
			"asset_model ": "%" + in.AssetModel + "%",
		})
	}
	// 生产厂家
	if len(in.ManuFacturer) > 0 {
		whereBuilder = whereBuilder.Where(squirrel.Like{
			"manu_facturer ": "%" + in.ManuFacturer + "%",
		})
	}
	// 电压
	if len(in.Voltage) > 0 {
		whereBuilder = whereBuilder.Where(squirrel.Like{
			"voltage ": "%" + in.Voltage + "%",
		})
	}
	// 容量
	if len(in.Capacity) > 0 {
		whereBuilder = whereBuilder.Where(squirrel.Like{
			"capacity ": "%" + in.Capacity + "%",
		})
	}

	all, err := l.svcCtx.TpmtAssetModel.FindList(l.ctx, whereBuilder, in.Current, in.PageSize)
	if err != nil {
		return nil, err
	}

	countBuilder := l.svcCtx.TpmtAssetModel.CountBuilder("id")

	// 资产类型
	if in.AssetType != 99 {
		countBuilder = countBuilder.Where(squirrel.Eq{
			"asset_type ": in.AssetType,
		})
	}
	// 资产编号
	if len(in.AssetCode) > 0 {
		countBuilder = countBuilder.Where(squirrel.Like{
			"asset_code ": "%" + in.AssetCode + "%",
		})
	}
	// 资产名称
	if len(in.AssetName) > 0 {
		countBuilder = countBuilder.Where(squirrel.Like{
			"asset_name ": "%" + in.AssetName + "%",
		})
	}
	// 资产型号
	if len(in.AssetModel) > 0 {
		countBuilder = countBuilder.Where(squirrel.Like{
			"asset_model ": "%" + in.AssetModel + "%",
		})
	}
	// 生产厂家
	if len(in.ManuFacturer) > 0 {
		countBuilder = countBuilder.Where(squirrel.Like{
			"manu_facturer ": "%" + in.ManuFacturer + "%",
		})
	}
	// 电压
	if len(in.Voltage) > 0 {
		countBuilder = countBuilder.Where(squirrel.Like{
			"voltage ": "%" + in.Voltage + "%",
		})
	}
	// 容量
	if len(in.Capacity) > 0 {
		countBuilder = countBuilder.Where(squirrel.Like{
			"capacity ": "%" + in.Capacity + "%",
		})
	}
	count, err := l.svcCtx.TpmtAssetModel.FindCount(l.ctx, countBuilder)
	if err != nil {
		return nil, err
	}

	var list []*tpmtclient.TpmtAssetListData
	for _, item := range all {
		list = append(list, &tpmtclient.TpmtAssetListData{
			Id:           item.Id,                         //资产ID
			AssetType:    item.AssetType,                  //资产类型
			AssetCode:    item.AssetCode,                  //资产编号
			AssetName:    item.AssetName,                  //资产名称
			AssetModel:   item.AssetModel,                 //资产型号
			ManuFacturer: item.ManuFacturer,               //生产厂家
			Voltage:      item.Voltage,                    //电压
			Capacity:     item.Capacity,                   //容量
			CreatedAt:    item.CreatedAt.UnixMilli(),      //创建时间
			CreatedName:  item.CreatedName,                //创建人
			UpdatedAt:    item.UpdatedAt.Time.UnixMilli(), //更新时间
			UpdatedName:  item.UpdatedName.String,         //更新人
		})
	}

	return &tpmtclient.TpmtAssetListResp{
		Total: count,
		List:  list,
	}, nil
}
