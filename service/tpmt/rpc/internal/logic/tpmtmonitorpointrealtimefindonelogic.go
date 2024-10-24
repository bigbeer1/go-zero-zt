package logic

import (
	"context"
	"errors"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"tpmt-zt/common/alarm"
	"tpmt-zt/common/datax"
	"tpmt-zt/common/eval"
	"tpmt-zt/service/tpmt/rpc/internal/svc"
	"tpmt-zt/service/tpmt/rpc/tpmtclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type TpmtMonitorPointRealTimeFindOneLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewTpmtMonitorPointRealTimeFindOneLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TpmtMonitorPointRealTimeFindOneLogic {
	return &TpmtMonitorPointRealTimeFindOneLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取单个传感器数据
func (l *TpmtMonitorPointRealTimeFindOneLogic) TpmtMonitorPointRealTimeFindOne(in *tpmtclient.TpmtMonitorPointRealTimeFindOneReq) (*tpmtclient.TpmtMonitorPointRealTimeFindOneResp, error) {
	res, err := l.svcCtx.TpmtMonitorPointModel.FindOne(l.ctx, in.Id)
	if err != nil {
		if errors.Is(err, sqlc.ErrNotFound) {
			return nil, fmt.Errorf("TpmtMonitorPoint没有该ID:%v", in.Id)
		}
		return nil, err
	}

	var assetRes *tpmtclient.TpmtAssetFindOneResp

	if res.AssetId != "" {
		//查询该点位关联的资产信息
		asset, err := l.svcCtx.TpmtAssetModel.FindOne(l.ctx, res.AssetId)
		if err == nil {
			assetRes = &tpmtclient.TpmtAssetFindOneResp{
				Id:           asset.Id,
				AssetType:    asset.AssetType,
				AssetCode:    asset.AssetCode,
				AssetName:    asset.AssetName,
				AssetModel:   asset.AssetModel,
				ManuFacturer: asset.ManuFacturer,
				Voltage:      asset.Voltage,
				Capacity:     asset.Capacity,
				CreatedAt:    asset.CreatedAt.UnixMilli(),
				CreatedName:  asset.CreatedName,
				UpdatedAt:    asset.UpdatedAt.Time.UnixMilli(),
				UpdatedName:  asset.UpdatedName.String,
			}
		}
	}

	redisKey := fmt.Sprintf("action_time:id:%v",
		res.Id)
	// 根据key获取redis中的实时数据
	temp := l.svcCtx.TpmtMonitorPointModel.GetRealData(l.ctx, redisKey)

	// 获取实时数据和时间
	resultValue, updateTime := eval.MonitorValueEval(temp, res.Coefficient)

	var alarmInfo *tpmtclient.AlarmRuleInfo
	if resultValue != "-" {
		resultValueFloat64, err := datax.ToFloat64(resultValue)
		if err != nil {
			return nil, fmt.Errorf("resultValueFloat64转换失败！" + err.Error())
		}
		info := alarm.CheckAlarmRuleYc(res, resultValueFloat64)
		if info != nil {
			alarmInfo = &tpmtclient.AlarmRuleInfo{
				Level:    info.Level,
				RuleType: info.RuleType,
				RuleData: info.RuleData,
			}

		}
	}

	return &tpmtclient.TpmtMonitorPointRealTimeFindOneResp{
		Id:                        res.Id,                         //监测点ID
		CreatedAt:                 res.CreatedAt.UnixMilli(),      //创建时间
		UpdatedAt:                 res.UpdatedAt.Time.UnixMilli(), //更新时间
		CreatedName:               res.CreatedName,                //创建人
		UpdatedName:               res.UpdatedName.String,         //更新人
		SerialNumber:              res.SerialNumber,               //编号
		Name:                      res.Name,                       //监测点名称
		RegisterAddress:           res.RegisterAddress,            //寄存器地址
		PointCollectorInstruction: res.PointCollectorInstruction,  //采集器指令  1: 01  2: 02  3:03  4:04
		PointAnalysisRule:         res.PointAnalysisRule,          //采集器解析规则 1: 16位无符号/2:单精度浮点数
		PointType:                 res.PointType,                  //类型：1:综保/2:局放/3:测温/4:微水/5:油色谱/6:机器人/7:其他
		PointCategory:             res.PointCategory,              //类别：1:遥信/2:遥测/3:遥脉
		PointGroup:                res.PointGroup,                 //分组
		CircuitType:               res.CircuitType,                //回路类型
		YxDecode:                  res.YxDecode.String,            //遥信解译
		DataBits:                  res.DataBits,                   //数据位
		Coefficient:               res.Coefficient,                //系数
		RetainDecimals:            res.RetainDecimals,             //保留小数位
		Unit:                      res.Unit.String,                //单位
		AlarmDuration:             res.AlarmDuration,              //持续时间
		AlarmUpValue:              res.AlarmUpValue,               //告警上限
		AlarmDownValue:            res.AlarmDownValue,             //告警下限
		WarningUpValue:            res.WarningUpValue,             //预警上限
		WarningDownValue:          res.WarningDownValue,           //预警下限
		IsDisplacementWarning:     res.IsDisplacementWarning,      //变位预警 0 不启用 1:启用
		TpmtGatewayId:             res.TpmtGatewayId,              //网关ID
		AssetId:                   res.AssetId,                    //资产ID
		Sort:                      res.Sort,                       //排序
		ResultValue:               resultValue,
		UpdateTime:                updateTime,
		AlarmRuleInfo:             alarmInfo,
		Asset:                     assetRes,
	}, nil
}
