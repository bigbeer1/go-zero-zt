package logic

import (
	"context"
	"github.com/Masterminds/squirrel"

	"tpmt-zt/service/tpmt/rpc/internal/svc"
	"tpmt-zt/service/tpmt/rpc/tpmtclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type TpmtMonitorPointListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewTpmtMonitorPointListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TpmtMonitorPointListLogic {
	return &TpmtMonitorPointListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *TpmtMonitorPointListLogic) TpmtMonitorPointList(in *tpmtclient.TpmtMonitorPointListReq) (*tpmtclient.TpmtMonitorPointListResp, error) {

	whereBuilder := l.svcCtx.TpmtMonitorPointModel.RowBuilder()

	whereBuilder = whereBuilder.OrderBy("created_at DESC, id DESC")

	// 编号
	if len(in.SerialNumber) > 0 {
		whereBuilder = whereBuilder.Where(squirrel.Like{
			"serial_number ": "%" + in.SerialNumber + "%",
		})
	}
	// 监测点名称
	if len(in.Name) > 0 {
		whereBuilder = whereBuilder.Where(squirrel.Like{
			"name ": "%" + in.Name + "%",
		})
	}
	// 寄存器地址
	if len(in.RegisterAddress) > 0 {
		whereBuilder = whereBuilder.Where(squirrel.Like{
			"register_address ": "%" + in.RegisterAddress + "%",
		})
	}
	// 采集器指令  1: 01  2: 02  3:03  4:04
	if in.PointCollectorInstruction != 99 {
		whereBuilder = whereBuilder.Where(squirrel.Eq{
			"point_collector_instruction ": in.PointCollectorInstruction,
		})
	}
	// 采集器解析规则 1: 16位无符号/2:单精度浮点数
	if in.PointAnalysisRule != 99 {
		whereBuilder = whereBuilder.Where(squirrel.Eq{
			"point_analysis_rule ": in.PointAnalysisRule,
		})
	}
	// 类型：1:综保/2:局放/3:测温/4:微水/5:油色谱/6:机器人/7:其他
	if in.PointType != 99 {
		whereBuilder = whereBuilder.Where(squirrel.Eq{
			"point_type ": in.PointType,
		})
	}
	// 类别：1:遥信/2:遥测/3:遥脉
	if in.PointCategory != 99 {
		whereBuilder = whereBuilder.Where(squirrel.Eq{
			"point_category ": in.PointCategory,
		})
	}
	// 分组
	if in.PointGroup != 99 {
		whereBuilder = whereBuilder.Where(squirrel.Eq{
			"point_group ": in.PointGroup,
		})
	}
	// 回路类型
	if in.CircuitType != 99 {
		whereBuilder = whereBuilder.Where(squirrel.Eq{
			"circuit_type ": in.CircuitType,
		})
	}

	// 数据位
	if in.DataBits != 99 {
		whereBuilder = whereBuilder.Where(squirrel.Eq{
			"data_bits ": in.DataBits,
		})
	}

	// 变位预警 0 不启用 1:启用
	if in.IsDisplacementWarning != 99 {
		whereBuilder = whereBuilder.Where(squirrel.Eq{
			"is_displacement_warning ": in.IsDisplacementWarning,
		})
	}
	// 网关ID
	if len(in.TpmtGatewayId) > 0 {
		whereBuilder = whereBuilder.Where(squirrel.Eq{
			"tpmt_gateway_id ": in.TpmtGatewayId,
		})
	}
	// 资产ID
	if len(in.AssetId) > 0 {
		whereBuilder = whereBuilder.Where(squirrel.Eq{
			"asset_id ": in.AssetId,
		})
	}

	all, err := l.svcCtx.TpmtMonitorPointModel.FindList(l.ctx, whereBuilder, in.Current, in.PageSize)
	if err != nil {
		return nil, err
	}

	countBuilder := l.svcCtx.TpmtMonitorPointModel.CountBuilder("id")

	// 编号
	if len(in.SerialNumber) > 0 {
		countBuilder = countBuilder.Where(squirrel.Like{
			"serial_number ": "%" + in.SerialNumber + "%",
		})
	}
	// 监测点名称
	if len(in.Name) > 0 {
		countBuilder = countBuilder.Where(squirrel.Like{
			"name ": "%" + in.Name + "%",
		})
	}
	// 寄存器地址
	if len(in.RegisterAddress) > 0 {
		countBuilder = countBuilder.Where(squirrel.Like{
			"register_address ": "%" + in.RegisterAddress + "%",
		})
	}
	// 采集器指令  1: 01  2: 02  3:03  4:04
	if in.PointCollectorInstruction != 99 {
		countBuilder = countBuilder.Where(squirrel.Eq{
			"point_collector_instruction ": in.PointCollectorInstruction,
		})
	}
	// 采集器解析规则 1: 16位无符号/2:单精度浮点数
	if in.PointAnalysisRule != 99 {
		countBuilder = countBuilder.Where(squirrel.Eq{
			"point_analysis_rule ": in.PointAnalysisRule,
		})
	}
	// 类型：1:综保/2:局放/3:测温/4:微水/5:油色谱/6:机器人/7:其他
	if in.PointType != 99 {
		countBuilder = countBuilder.Where(squirrel.Eq{
			"point_type ": in.PointType,
		})
	}
	// 类别：1:遥信/2:遥测/3:遥脉
	if in.PointCategory != 99 {
		countBuilder = countBuilder.Where(squirrel.Eq{
			"point_category ": in.PointCategory,
		})
	}
	// 分组
	if in.PointGroup != 99 {
		countBuilder = countBuilder.Where(squirrel.Eq{
			"point_group ": in.PointGroup,
		})
	}
	// 回路类型
	if in.CircuitType != 99 {
		countBuilder = countBuilder.Where(squirrel.Eq{
			"circuit_type ": in.CircuitType,
		})
	}

	// 数据位
	if in.DataBits != 99 {
		countBuilder = countBuilder.Where(squirrel.Eq{
			"data_bits ": in.DataBits,
		})
	}

	// 变位预警 0 不启用 1:启用
	if in.IsDisplacementWarning != 99 {
		countBuilder = countBuilder.Where(squirrel.Eq{
			"is_displacement_warning ": in.IsDisplacementWarning,
		})
	}
	// 网关ID
	if len(in.TpmtGatewayId) > 0 {
		countBuilder = countBuilder.Where(squirrel.Eq{
			"tpmt_gateway_id ": in.TpmtGatewayId,
		})
	}
	// 资产ID
	if len(in.AssetId) > 0 {
		countBuilder = countBuilder.Where(squirrel.Eq{
			"asset_id ": in.AssetId,
		})
	}
	count, err := l.svcCtx.TpmtMonitorPointModel.FindCount(l.ctx, countBuilder)
	if err != nil {
		return nil, err
	}

	var list []*tpmtclient.TpmtMonitorPointListData
	for _, item := range all {
		list = append(list, &tpmtclient.TpmtMonitorPointListData{
			Id:                        item.Id,                         //监测点ID
			CreatedAt:                 item.CreatedAt.UnixMilli(),      //创建时间
			UpdatedAt:                 item.UpdatedAt.Time.UnixMilli(), //更新时间
			CreatedName:               item.CreatedName,                //创建人
			UpdatedName:               item.UpdatedName.String,         //更新人
			SerialNumber:              item.SerialNumber,               //编号
			Name:                      item.Name,                       //监测点名称
			RegisterAddress:           item.RegisterAddress,            //寄存器地址
			PointCollectorInstruction: item.PointCollectorInstruction,  //采集器指令  1: 01  2: 02  3:03  4:04
			PointAnalysisRule:         item.PointAnalysisRule,          //采集器解析规则 1: 16位无符号/2:单精度浮点数
			PointType:                 item.PointType,                  //类型：1:综保/2:局放/3:测温/4:微水/5:油色谱/6:机器人/7:其他
			PointCategory:             item.PointCategory,              //类别：1:遥信/2:遥测/3:遥脉
			PointGroup:                item.PointGroup,                 //分组
			CircuitType:               item.CircuitType,                //回路类型
			YxDecode:                  item.YxDecode.String,            //遥信解译
			DataBits:                  item.DataBits,                   //数据位
			Coefficient:               item.Coefficient,                //系数
			RetainDecimals:            item.RetainDecimals,             //保留小数位
			Unit:                      item.Unit.String,                //单位
			AlarmDuration:             item.AlarmDuration,              //持续时间
			AlarmUpValue:              item.AlarmUpValue,               //告警上限
			AlarmDownValue:            item.AlarmDownValue,             //告警下限
			WarningUpValue:            item.WarningUpValue,             //预警上限
			WarningDownValue:          item.WarningDownValue,           //预警下限
			IsDisplacementWarning:     item.IsDisplacementWarning,      //变位预警 0 不启用 1:启用
			TpmtGatewayId:             item.TpmtGatewayId,              //网关ID
			AssetId:                   item.AssetId,                    //资产ID
			Sort:                      item.Sort,                       //排序
		})
	}

	return &tpmtclient.TpmtMonitorPointListResp{
		Total: count,
		List:  list,
	}, nil
}
