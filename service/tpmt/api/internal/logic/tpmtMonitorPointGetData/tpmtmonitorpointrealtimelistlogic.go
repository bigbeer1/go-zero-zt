package tpmtMonitorPointGetData

import (
	"context"
	"github.com/jinzhu/copier"
	"tpmt-zt/common"
	"tpmt-zt/common/msg"
	"tpmt-zt/service/tpmt/api/internal/logic/tpmtAsset"
	"tpmt-zt/service/tpmt/rpc/tpmtclient"

	"tpmt-zt/service/tpmt/api/internal/svc"
	"tpmt-zt/service/tpmt/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type TpmtMonitorPointRealTimeListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTpmtMonitorPointRealTimeListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TpmtMonitorPointRealTimeListLogic {
	return &TpmtMonitorPointRealTimeListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TpmtMonitorPointRealTimeListLogic) TpmtMonitorPointRealTimeList(req *types.TpmtMonitorPointRealTimeRequest) (resp *types.Response, err error) {

	all, err := l.svcCtx.TpmtRpc.TpmtMonitorPointRealTimeList(l.ctx, &tpmtclient.TpmtMonitorPointRealTimeListReq{
		Current:       req.Current,       // 页码
		PageSize:      req.PageSize,      // 页数
		SerialNumber:  req.SerialNumber,  // 编号
		Name:          req.Name,          // 监测点名称
		PointType:     req.PointType,     // 类型：1:综保/2:局放/3:测温/4:微水/5:油色谱/6:机器人/7:其他
		PointCategory: req.PointCategory, // 类别：1:遥信/2:遥测/3:遥脉
		PointGroup:    req.PointGroup,    // 分组
		CircuitType:   req.CircuitType,   // 回路类型
		TpmtGatewayId: req.TpmtGatewayId, // 网关ID
		AssetId:       req.AssetId,       // 资产ID
	})
	if err != nil {
		return nil, common.NewDefaultError(err.Error())
	}

	var result TpmtMonitorPointRealTimeListResp
	_ = copier.Copy(&result, all)

	return &types.Response{
		Code: 0,
		Msg:  msg.Success,
		Data: result,
	}, nil

}

type TpmtMonitorPointRealTimeListResp struct {
	Total int64                           `json:"total"` // 总数
	List  []*TpmtMonitorPointRealTimeData `json:"list"`  // 内容
}

// TpmtMonitorPointRealTimeList 列表信息
type TpmtMonitorPointRealTimeData struct {
	Id                        int64                           `json:"id"`                          // 监测点ID
	CreatedAt                 int64                           `json:"created_at"`                  // 创建时间
	UpdatedAt                 int64                           `json:"updated_at"`                  // 更新时间
	CreatedName               string                          `json:"created_name"`                // 创建人
	UpdatedName               string                          `json:"updated_name"`                // 更新人
	SerialNumber              string                          `json:"serial_number"`               // 编号
	Name                      string                          `json:"name"`                        // 监测点名称
	RegisterAddress           string                          `json:"register_address"`            // 寄存器地址
	PointCollectorInstruction int64                           `json:"point_collector_instruction"` // 采集器指令 1: 02  2:03  3:04
	PointAnalysisRule         int64                           `json:"point_analysis_rule"`         // 采集器解析规则 1: 16位无符号/2:单精度浮点数
	PointType                 int64                           `json:"point_type"`                  // 类型：1:综保/2:局放/3:测温/4:微水/5:油色谱/6:机器人/7:其他
	PointCategory             int64                           `json:"point_category"`              // 类别：1:遥信/2:遥测/3:遥脉
	PointGroup                int64                           `json:"point_group"`                 // 分组
	CircuitType               int64                           `json:"circuit_type"`                // 回路类型
	YxDecode                  string                          `json:"yx_decode"`                   // 遥信解译
	DataBits                  int64                           `json:"data_bits"`                   // 数据位
	Coefficient               float64                         `json:"coefficient"`                 // 系数
	RetainDecimals            int64                           `json:"retain_decimals"`             // 保留小数位
	Unit                      string                          `json:"unit"`                        // 单位
	AlarmDuration             int64                           `json:"alarm_duration"`              // 持续时间
	AlarmUpValue              float64                         `json:"alarm_up_value"`              // 告警上限
	AlarmDownValue            float64                         `json:"alarm_down_value"`            // 告警下限
	WarningUpValue            float64                         `json:"warning_up_value"`            // 预警上限
	WarningDownValue          float64                         `json:"warning_down_value"`          // 预警下限
	IsDisplacementWarning     int64                           `json:"is_displacement_warning"`     // 变位预警 0 不启用 1:启用
	TpmtGatewayId             string                          `json:"tpmt_gateway_id"`             // 网关ID
	AssetId                   string                          `json:"asset_id"`                    // 资产ID
	Sort                      int64                           `json:"sort"`                        // 排序
	ResultValue               string                          `json:"result_value"`                // 监测值
	UpdateTime                int64                           `json:"update_time"`                 // 值更新时间
	AlarmRuleInfo             *AlarmRuleInfo                  `json:"alarm_rule_info"`             // 数值状态
	Asset                     *tpmtAsset.TpmtAssetFindOneResp `json:"asset"`                       // 资产信息
}

// 告警内容
type AlarmRuleInfo struct {
	Level    int64  `json:"level"`     // 告警等级
	RuleType int64  `json:"rule_type"` // 1越上限,2越下限,3相同
	RuleData string `json:"rule_data"` // 数值带单位
}