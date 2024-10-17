package tpmtMonitorPoint

import (
	"context"
	"tpmt-zt/common"
	"tpmt-zt/common/jwtx"
	"tpmt-zt/common/msg"
	"tpmt-zt/service/tpmt/rpc/tpmtclient"

	"tpmt-zt/service/tpmt/api/internal/svc"
	"tpmt-zt/service/tpmt/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type TpmtMonitorPointUpLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTpmtMonitorPointUpLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TpmtMonitorPointUpLogic {
	return &TpmtMonitorPointUpLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TpmtMonitorPointUpLogic) TpmtMonitorPointUp(req *types.TpmtMonitorPointUpRequest) (resp *types.Response, err error) {
	// 用户登录信息
	tokenData := jwtx.ParseToken(l.ctx)

	_, err = l.svcCtx.TpmtRpc.TpmtMonitorPointUpdate(l.ctx, &tpmtclient.TpmtMonitorPointUpdateReq{
		Id:                        req.Id,                        // 监测点ID
		UpdatedName:               tokenData.NickName,            // 更新人
		SerialNumber:              req.SerialNumber,              // 编号
		Name:                      req.Name,                      // 监测点名称
		RegisterAddress:           req.RegisterAddress,           // 寄存器地址
		PointCollectorInstruction: req.PointCollectorInstruction, // 采集器指令  1: 01  2: 02  3:03  4:04
		PointAnalysisRule:         req.PointAnalysisRule,         // 采集器解析规则 1: 16位无符号/2:单精度浮点数
		PointType:                 req.PointType,                 // 类型：1:综保/2:局放/3:测温/4:微水/5:油色谱/6:机器人/7:其他
		PointCategory:             req.PointCategory,             // 类别：1:遥信/2:遥测/3:遥脉
		PointGroup:                req.PointGroup,                // 分组
		CircuitType:               req.CircuitType,               // 回路类型
		YxDecode:                  req.YxDecode,                  // 遥信解译
		DataBits:                  req.DataBits,                  // 数据位
		Coefficient:               req.Coefficient,               // 系数
		RetainDecimals:            req.RetainDecimals,            // 保留小数位
		Unit:                      req.Unit,                      // 单位
		AlarmDuration:             req.AlarmDuration,             // 持续时间
		AlarmUpValue:              req.AlarmUpValue,              // 告警上限
		AlarmDownValue:            req.AlarmDownValue,            // 告警下限
		WarningUpValue:            req.WarningUpValue,            // 预警上限
		WarningDownValue:          req.WarningDownValue,          // 预警下限
		IsDisplacementWarning:     req.IsDisplacementWarning,     // 变位预警 0 不启用 1:启用
		AssetId:                   req.AssetId,                   // 资产ID
		Sort:                      req.Sort,                      // 排序
	})
	if err != nil {
		return nil, common.NewDefaultError(err.Error())
	}
	return &types.Response{
		Code: 0,
		Msg:  msg.Success,
		Data: nil,
	}, nil
}
