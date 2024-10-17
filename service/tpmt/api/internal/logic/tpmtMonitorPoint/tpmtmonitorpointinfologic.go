package tpmtMonitorPoint

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

type TpmtMonitorPointInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTpmtMonitorPointInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TpmtMonitorPointInfoLogic {
	return &TpmtMonitorPointInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TpmtMonitorPointInfoLogic) TpmtMonitorPointInfo(req *types.TpmtMonitorPointInfoRequest) (resp *types.Response, err error) {

	res, err := l.svcCtx.TpmtRpc.TpmtMonitorPointFindOne(l.ctx, &tpmtclient.TpmtMonitorPointFindOneReq{
		Id: req.Id, // 监测点ID
	})
	if err != nil {
		return nil, common.NewDefaultError(err.Error())
	}

	var result TpmtMonitorPointFindOneResp
	_ = copier.Copy(&result, res)

	return &types.Response{
		Code: 0,
		Msg:  msg.Success,
		Data: result,
	}, nil
}

type TpmtMonitorPointFindOneResp struct {
	Id                        int64   `json:"id"`                          // 监测点ID,
	CreatedAt                 int64   `json:"created_at"`                  // 创建时间,
	UpdatedAt                 int64   `json:"updated_at"`                  // 更新时间,
	CreatedName               string  `json:"created_name"`                // 创建人,
	UpdatedName               string  `json:"updated_name"`                // 更新人,
	SerialNumber              string  `json:"serial_number"`               // 编号,
	Name                      string  `json:"name"`                        // 监测点名称,
	RegisterAddress           string  `json:"register_address"`            // 寄存器地址,
	PointCollectorInstruction int64   `json:"point_collector_instruction"` // 采集器指令  1: 01  2: 02  3:03  4:04,
	PointAnalysisRule         int64   `json:"point_analysis_rule"`         // 采集器解析规则 1: 16位无符号/2:单精度浮点数,
	PointType                 int64   `json:"point_type"`                  // 类型：1:综保/2:局放/3:测温/4:微水/5:油色谱/6:机器人/7:其他,
	PointCategory             int64   `json:"point_category"`              // 类别：1:遥信/2:遥测/3:遥脉,
	PointGroup                int64   `json:"point_group"`                 // 分组,
	CircuitType               int64   `json:"circuit_type"`                // 回路类型,
	YxDecode                  string  `json:"yx_decode"`                   // 遥信解译,
	DataBits                  int64   `json:"data_bits"`                   // 数据位,
	Coefficient               float64 `json:"coefficient"`                 // 系数,
	RetainDecimals            int64   `json:"retain_decimals"`             // 保留小数位,
	Unit                      string  `json:"unit"`                        // 单位,
	AlarmDuration             int64   `json:"alarm_duration"`              // 持续时间,
	AlarmUpValue              float64 `json:"alarm_up_value"`              // 告警上限,
	AlarmDownValue            float64 `json:"alarm_down_value"`            // 告警下限,
	WarningUpValue            float64 `json:"warning_up_value"`            // 预警上限,
	WarningDownValue          float64 `json:"warning_down_value"`          // 预警下限,
	IsDisplacementWarning     int64   `json:"is_displacement_warning"`     // 变位预警 0 不启用 1:启用,
	TpmtGatewayId             string  `json:"tpmt_gateway_id"`             // 网关ID,
	AssetId                   string  `json:"asset_id"`                    // 资产ID,
	Sort                      int64   `json:"sort"`                        // 排序
}
