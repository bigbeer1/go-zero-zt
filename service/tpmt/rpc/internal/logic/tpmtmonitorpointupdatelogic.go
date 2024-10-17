package logic

import (
	"context"
	"errors"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"time"

	"tpmt-zt/service/tpmt/rpc/internal/svc"
	"tpmt-zt/service/tpmt/rpc/tpmtclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type TpmtMonitorPointUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewTpmtMonitorPointUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TpmtMonitorPointUpdateLogic {
	return &TpmtMonitorPointUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *TpmtMonitorPointUpdateLogic) TpmtMonitorPointUpdate(in *tpmtclient.TpmtMonitorPointUpdateReq) (*tpmtclient.CommonResp, error) {

	res, err := l.svcCtx.TpmtMonitorPointModel.FindOne(l.ctx, in.Id)
	if err != nil {
		if errors.Is(err, sqlc.ErrNotFound) {
			return nil, fmt.Errorf("TpmtMonitorPoint没有该ID: %v", in.Id)
		}
		return nil, err
	}

	// 编号
	if len(in.SerialNumber) > 0 {
		res.SerialNumber = in.SerialNumber
	}
	// 监测点名称
	if len(in.Name) > 0 {
		res.Name = in.Name
	}
	// 寄存器地址
	if len(in.RegisterAddress) > 0 {
		res.RegisterAddress = in.RegisterAddress
	}
	// 采集器指令  1: 01  2: 02  3:03  4:04
	if in.PointCollectorInstruction != 0 {
		res.PointCollectorInstruction = in.PointCollectorInstruction
	}
	// 采集器解析规则 1: 16位无符号/2:单精度浮点数
	if in.PointAnalysisRule != 0 {
		res.PointAnalysisRule = in.PointAnalysisRule
	}
	// 类型：1:综保/2:局放/3:测温/4:微水/5:油色谱/6:机器人/7:其他
	if in.PointType != 0 {
		res.PointType = in.PointType
	}
	// 类别：1:遥信/2:遥测/3:遥脉
	if in.PointCategory != 0 {
		res.PointCategory = in.PointCategory
	}
	// 分组
	if in.PointGroup != 0 {
		res.PointGroup = in.PointGroup
	}
	// 回路类型
	if in.CircuitType != 0 {
		res.CircuitType = in.CircuitType
	}
	// 遥信解译
	if len(in.YxDecode) > 0 {
		res.YxDecode.String = in.YxDecode
		res.YxDecode.Valid = true
	}
	// 数据位
	if in.DataBits != 0 {
		res.DataBits = in.DataBits
	}
	// 系数
	if in.Coefficient != 0.0 {
		res.Coefficient = in.Coefficient
	}
	// 保留小数位
	if in.RetainDecimals != 0 {
		res.RetainDecimals = in.RetainDecimals
	}
	// 单位
	if len(in.Unit) > 0 {
		res.Unit.String = in.Unit
		res.Unit.Valid = true
	}
	// 持续时间
	if in.AlarmDuration != 0 {
		res.AlarmDuration = in.AlarmDuration
	}
	// 告警上限
	if in.AlarmUpValue != 0.0 {
		res.AlarmUpValue = in.AlarmUpValue
	}
	// 告警下限
	if in.AlarmDownValue != 0.0 {
		res.AlarmDownValue = in.AlarmDownValue
	}
	// 预警上限
	if in.WarningUpValue != 0.0 {
		res.WarningUpValue = in.WarningUpValue
	}
	// 预警下限
	if in.WarningDownValue != 0.0 {
		res.WarningDownValue = in.WarningDownValue
	}
	// 变位预警 0 不启用 1:启用
	if in.IsDisplacementWarning != 0 {
		res.IsDisplacementWarning = in.IsDisplacementWarning
	}

	// 资产ID
	if res.AssetId != in.AssetId {
		if in.AssetId != "" {
			_, err := l.svcCtx.TpmtAssetModel.FindOne(l.ctx, in.AssetId)
			if err != nil {
				if errors.Is(err, sqlc.ErrNotFound) {
					return nil, fmt.Errorf("TpmtAsset没有该ID:%v", in.AssetId)
				}
				return nil, err
			}
			res.AssetId = in.AssetId
		} else {
			res.AssetId = ""
		}
	}
	// 排序
	if in.Sort != 0 {
		res.Sort = in.Sort
	}

	res.UpdatedName.String = in.UpdatedName
	res.UpdatedName.Valid = true
	res.UpdatedAt.Time = time.Now()
	res.UpdatedAt.Valid = true

	err = l.svcCtx.TpmtMonitorPointModel.Update(l.ctx, res)

	if err != nil {
		return nil, err
	}
	return &tpmtclient.CommonResp{}, nil

}
