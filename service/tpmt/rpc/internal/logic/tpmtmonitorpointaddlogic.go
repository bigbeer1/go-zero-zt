package logic

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"time"
	"tpmt-zt/service/tpmt/model"

	"tpmt-zt/service/tpmt/rpc/internal/svc"
	"tpmt-zt/service/tpmt/rpc/tpmtclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type TpmtMonitorPointAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewTpmtMonitorPointAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TpmtMonitorPointAddLogic {
	return &TpmtMonitorPointAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 监控点
func (l *TpmtMonitorPointAddLogic) TpmtMonitorPointAdd(in *tpmtclient.TpmtMonitorPointAddReq) (*tpmtclient.CommonResp, error) {

	// 判断资产是否添加和是否存在
	if in.AssetId != "" {
		_, err := l.svcCtx.TpmtAssetModel.FindOne(l.ctx, in.AssetId)
		if err != nil {
			if errors.Is(err, sqlc.ErrNotFound) {
				return nil, fmt.Errorf("TpmtAsset没有该ID:%v", in.AssetId)
			}
			return nil, err
		}
	}

	_, err := l.svcCtx.TpmtGatewayModel.FindOne(l.ctx, in.TpmtGatewayId)
	if err != nil {
		if errors.Is(err, sqlc.ErrNotFound) {
			return nil, fmt.Errorf("TpmtGatewayId没有该ID:%v", in.TpmtGatewayId)
		}
		return nil, err
	}

	_, err = l.svcCtx.TpmtMonitorPointModel.Insert(l.ctx, &model.TpmtMonitorPoint{
		CreatedAt:                 time.Now(),                                                    // 创建时间
		CreatedName:               in.CreatedName,                                                // 创建人
		SerialNumber:              in.SerialNumber,                                               // 编号
		Name:                      in.Name,                                                       // 监测点名称
		RegisterAddress:           in.RegisterAddress,                                            // 寄存器地址
		PointCollectorInstruction: in.PointCollectorInstruction,                                  // 采集器指令  1: 01  2: 02  3:03  4:04
		PointAnalysisRule:         in.PointAnalysisRule,                                          // 采集器解析规则 1: 16位无符号/2:单精度浮点数
		PointType:                 in.PointType,                                                  // 类型：1:综保/2:局放/3:测温/4:微水/5:油色谱/6:机器人/7:其他
		PointCategory:             in.PointCategory,                                              // 类别：1:遥信/2:遥测/3:遥脉
		PointGroup:                in.PointGroup,                                                 // 分组
		CircuitType:               in.CircuitType,                                                // 回路类型
		YxDecode:                  sql.NullString{String: in.YxDecode, Valid: in.YxDecode != ""}, // 遥信解译
		DataBits:                  in.DataBits,                                                   // 数据位
		Coefficient:               in.Coefficient,                                                // 系数
		RetainDecimals:            in.RetainDecimals,                                             // 保留小数位
		Unit:                      sql.NullString{String: in.Unit, Valid: in.Unit != ""},         // 单位
		AlarmDuration:             in.AlarmDuration,                                              // 持续时间
		AlarmUpValue:              in.AlarmUpValue,                                               // 告警上限
		AlarmDownValue:            in.AlarmDownValue,                                             // 告警下限
		WarningUpValue:            in.WarningUpValue,                                             // 预警上限
		WarningDownValue:          in.WarningDownValue,                                           // 预警下限
		IsDisplacementWarning:     in.IsDisplacementWarning,                                      // 变位预警 0 不启用 1:启用
		TpmtGatewayId:             in.TpmtGatewayId,                                              // 网关ID
		AssetId:                   in.AssetId,                                                    // 资产ID
		Sort:                      in.Sort,                                                       // 排序
	})
	if err != nil {
		return nil, err
	}

	return &tpmtclient.CommonResp{}, nil
}
