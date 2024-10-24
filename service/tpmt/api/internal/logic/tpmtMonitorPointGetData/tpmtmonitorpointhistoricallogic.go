package tpmtMonitorPointGetData

import (
	"context"
	"tpmt-zt/common"
	"tpmt-zt/common/msg"
	"tpmt-zt/service/tpmt/rpc/tpmtclient"

	"tpmt-zt/service/tpmt/api/internal/svc"
	"tpmt-zt/service/tpmt/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type TpmtMonitorPointHistoricalLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTpmtMonitorPointHistoricalLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TpmtMonitorPointHistoricalLogic {
	return &TpmtMonitorPointHistoricalLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TpmtMonitorPointHistoricalLogic) TpmtMonitorPointHistorical(req *types.TpmtmonitorPointHistoricalReq) (resp *types.Response, err error) {

	var list []*TpmtMonitorPointHistoricalListResp

	if req.TimeRangeEndDayTime <= req.TimeRangeStartDayTime && req.TimeRangeEndDayTime != 0 && req.TimeRangeStartDayTime != 0 {
		return nil, common.NewDefaultError("结束时间不能小于等于开始时间")
	}

	// 获取ID 去重
	IdsMap := map[int64]int64{}

	for _, id := range req.Ids {
		IdsMap[id] = id
	}

	for _, id := range IdsMap {
		// 请求RPC方法
		res, err := l.svcCtx.TpmtRpc.TpmtMonitorPointHistorical(l.ctx, &tpmtclient.TpmtMonitorPointHistoricalReq{
			Id:                    id,
			TimeRangeStartDayTime: req.TimeRangeStartDayTime,
			TimeRangeEndDayTime:   req.TimeRangeEndDayTime,
		})

		if err != nil {
			return nil, common.NewDefaultError(err.Error())
		}

		list = append(list, &TpmtMonitorPointHistoricalListResp{
			Id:                        res.Id,
			CreatedAt:                 res.CreatedAt,
			UpdatedAt:                 res.UpdatedAt,
			CreatedName:               res.CreatedName,
			UpdatedName:               res.UpdatedName,
			SerialNumber:              res.SerialNumber,
			Name:                      res.Name,
			RegisterAddress:           res.RegisterAddress,
			PointCollectorInstruction: res.PointCollectorInstruction,
			PointAnalysisRule:         res.PointAnalysisRule,
			PointType:                 res.PointType,
			PointCategory:             res.PointCategory,
			PointGroup:                res.PointGroup,
			CircuitType:               res.CircuitType,
			YxDecode:                  res.YxDecode,
			DataBits:                  res.DataBits,
			Coefficient:               res.Coefficient,
			RetainDecimals:            res.RetainDecimals,
			Unit:                      res.Unit,
			AlarmDuration:             res.AlarmDuration,
			AlarmUpValue:              res.AlarmUpValue,
			AlarmDownValue:            res.AlarmDownValue,
			WarningUpValue:            res.WarningUpValue,
			WarningDownValue:          res.WarningDownValue,
			IsDisplacementWarning:     res.IsDisplacementWarning,
			TpmtGatewayId:             res.TpmtGatewayId,
			AssetId:                   res.AssetId,
			AssetName:                 res.AssetName,
			Sort:                      res.Sort,
			Data:                      res.List,
		})
	}

	return &types.Response{
		Code: 0,
		Msg:  msg.Success,
		Data: list,
	}, nil
}

type TpmtMonitorPointHistoricalListResp struct {
	Id                        int64       `json:"id"`                          // 监测点ID
	CreatedAt                 int64       `json:"created_at"`                  // 创建时间
	UpdatedAt                 int64       `json:"updated_at"`                  // 更新时间
	CreatedName               string      `json:"created_name"`                // 创建人
	UpdatedName               string      `json:"updated_name"`                // 更新人
	SerialNumber              string      `json:"serial_number"`               // 编号
	Name                      string      `json:"name"`                        // 监测点名称
	RegisterAddress           string      `json:"register_address"`            // 寄存器地址
	PointCollectorInstruction int64       `json:"point_collector_instruction"` // 采集器指令 1: 02  2:03  3:04
	PointAnalysisRule         int64       `json:"point_analysis_rule"`         // 采集器解析规则 1: 16位无符号/2:单精度浮点数
	PointType                 int64       `json:"point_type"`                  // 类型：1:综保/2:局放/3:测温/4:微水/5:油色谱/6:机器人/7:其他
	PointCategory             int64       `json:"point_category"`              // 类别：1:遥信/2:遥测/3:遥脉
	PointGroup                int64       `json:"point_group"`                 // 分组,
	CircuitType               int64       `json:"circuit_type"`                // 回路类型
	YxDecode                  string      `json:"yx_decode"`                   // 遥信解译
	DataBits                  int64       `json:"data_bits"`                   // 数据位
	Coefficient               float64     `json:"coefficient"`                 // 系数
	RetainDecimals            int64       `json:"retain_decimals"`             // 保留小数位
	Unit                      string      `json:"unit"`                        // 单位
	AlarmDuration             int64       `json:"alarm_duration"`              // 持续时间
	AlarmUpValue              float64     `json:"alarm_up_value"`              // 告警上限
	AlarmDownValue            float64     `json:"alarm_down_value"`            // 告警下限
	WarningUpValue            float64     `json:"warning_up_value"`            // 预警上限
	WarningDownValue          float64     `json:"warning_down_value"`          // 预警下限
	IsDisplacementWarning     int64       `json:"is_displacement_warning"`     // 变位预警 0 不启用 1:启用
	TpmtGatewayId             string      `json:"tpmt_gateway_id"`             // 网关ID
	AssetId                   string      `json:"asset_id"`                    // 资产Id
	AssetName                 string      `json:"asset_name"`                  // 资产名称
	Sort                      int64       `json:"sort"`                        // 排序
	Data                      interface{} `json:"data"`                        // 历史数据
}
