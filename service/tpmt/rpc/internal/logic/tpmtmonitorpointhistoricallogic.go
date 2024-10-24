package logic

import (
	"context"
	"errors"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"math/big"
	"strconv"
	"tpmt-zt/common/datax"
	"tpmt-zt/common/tdenginex"
	archivemodel "tpmt-zt/service/archive/model"
	"tpmt-zt/service/tpmt/rpc/internal/svc"
	"tpmt-zt/service/tpmt/rpc/tpmtclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type TpmtMonitorPointHistoricalLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewTpmtMonitorPointHistoricalLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TpmtMonitorPointHistoricalLogic {
	return &TpmtMonitorPointHistoricalLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取历史数据接口
func (l *TpmtMonitorPointHistoricalLogic) TpmtMonitorPointHistorical(in *tpmtclient.TpmtMonitorPointHistoricalReq) (*tpmtclient.TpmtMonitorPointHistoricalResp, error) {
	res, err := l.svcCtx.TpmtMonitorPointModel.FindOne(l.ctx, in.Id)
	if err != nil {
		if errors.Is(err, sqlc.ErrNotFound) {
			return nil, fmt.Errorf("TpmtMonitorPoint没有该ID:%v", in.Id)
		}
		return nil, err
	}

	var assetRes tpmtclient.TpmtAssetFindOneResp

	if res.AssetId != "" {
		//查询该点位关联的资产信息
		asset, err := l.svcCtx.TpmtAssetModel.FindOne(l.ctx, res.AssetId)
		if err == nil {
			assetRes = tpmtclient.TpmtAssetFindOneResp{
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

	// 存储时序数据库key
	dataKey := fmt.Sprintf("sequential_id_%v",
		res.Id)

	tddb := &archivemodel.TdDb{
		DbName:    "monitor_point." + dataKey,
		TableName: "monitor_point.tpmt_monitor_point",
	}

	// 默认加权平均
	every := "2m" // 一小时以内 默认以2分钟 为间隔

	historicalOverDayPoints := int64(300)
	historicalHalfDayMergeMinutes := "5m" // 1小时-12小时 默认以5分钟 为间隔
	historicalDayMergeMinutes := "10m"    //12小时-24小时 默认以10分钟 为间隔

	timex := in.TimeRangeEndDayTime - in.TimeRangeStartDayTime
	if timex > 86400000 {
		// 固定点数
		number := historicalOverDayPoints
		every = fmt.Sprintf("%vm", timex/(3600*1000/60)/number)
	}

	// 为了兼容前端可能出现的 毫秒偏差 这边摆动2毫秒
	// 1小时-12小时
	if 3599998 <= timex && timex < 43200002 {
		every = historicalHalfDayMergeMinutes
	}

	// 12小时-24小时
	if 43200002 <= timex && timex <= 86400002 {
		// 默认加权平均
		every = historicalDayMergeMinutes
	}

	tdMonitor := &archivemodel.TdMonitor{}

	all, err := tdMonitor.FindAll(l.ctx, l.svcCtx.Taos, tddb, in.TimeRangeStartDayTime, in.TimeRangeEndDayTime, every)
	if err != nil {
		if err.Error() != tdenginex.ErrNotFoundTable {
			return nil, err
		}
	}

	var list []*tpmtclient.TpmtMonitorPointHistoricalListData

	for _, item := range all {

		createdtime := item.Wend.UnixMilli()
		// 判断创建时间是否小于范围最后的时间,td存在会查出超出范围的时间用于去除
		if createdtime < in.TimeRangeEndDayTime || in.TimeRangeEndDayTime == 0 {
			// 乘以系数

			float64DataBig := big.NewFloat(item.Data)
			float64DataBig.Mul(float64DataBig, big.NewFloat(res.Coefficient))

			bit := "%." + "2" + "f"
			monitorValueFloat, _ := strconv.ParseFloat(fmt.Sprintf(bit, float64DataBig), 64)

			monitorValue := datax.ToString(monitorValueFloat)

			list = append(list, &tpmtclient.TpmtMonitorPointHistoricalListData{
				MonitorValue: monitorValue,
				CreateTime:   createdtime,
			})
		}

	}
	return &tpmtclient.TpmtMonitorPointHistoricalResp{
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
		TpmtGatewayId:             res.TpmtGatewayId,
		AssetId:                   res.AssetId,
		AssetName:                 assetRes.AssetName,
		Sort:                      res.Sort,
		List:                      list,
	}, nil
}
