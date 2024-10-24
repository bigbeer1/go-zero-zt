package logic

import (
	"context"
	"fmt"
	"github.com/Masterminds/squirrel"
	"sync"
	"tpmt-zt/common/alarm"
	"tpmt-zt/common/datax"
	"tpmt-zt/common/eval"
	"tpmt-zt/service/tpmt/model"
	"tpmt-zt/service/tpmt/rpc/internal/svc"
	"tpmt-zt/service/tpmt/rpc/tpmtclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type TpmtMonitorPointRealTimeListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewTpmtMonitorPointRealTimeListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TpmtMonitorPointRealTimeListLogic {
	return &TpmtMonitorPointRealTimeListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 分页获取传感器实时数据
func (l *TpmtMonitorPointRealTimeListLogic) TpmtMonitorPointRealTimeList(in *tpmtclient.TpmtMonitorPointRealTimeListReq) (*tpmtclient.TpmtMonitorPointRealTimeListResp, error) {
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

	// todo 添加限流器 30

	/*
		协程并发查询提升速度
	*/

	// 开启协程池

	var wg sync.WaitGroup
	// 协程获取返回数据  管道用于存储数据
	chanData := make(chan *tpmtclient.TpmtMonitorPointRealTimeData, len(all))
	// 协程数控制
	var limit = make(chan struct{}, 200)

	for _, item := range all {
		limit <- struct{}{}
		wg.Add(1)
		go func(item *model.TpmtMonitorPoint) {
			defer func() {
				wg.Done()
				<-limit
			}()
			var redisKey, resultValue string
			var updateTime int64

			redisKey = fmt.Sprintf("action_time:id:%v",
				item.Id)
			// 根据key获取redis中的实时数据
			temp := l.svcCtx.TpmtMonitorPointModel.GetRealData(l.ctx, redisKey)

			// 获取实时数据和时间
			resultValue, updateTime = eval.MonitorValueEval(temp, item.Coefficient)

			// 告警等级  0正常 1预警 2告警
			// 初始化告警状态为 无  0正常  1越上线  2越下线
			var alarmRuleInfo = &tpmtclient.AlarmRuleInfo{
				Level:    0,
				RuleType: 0,
			}

			// 查询告警规则
			resultValueFloat64, err := datax.ToFloat64(resultValue)
			if err == nil {
				info := alarm.CheckAlarmRuleYc(item, resultValueFloat64)
				if info != nil {
					alarmRuleInfo = &tpmtclient.AlarmRuleInfo{
						Level:    info.Level,
						RuleType: info.RuleType,
						RuleData: info.RuleData,
					}
				}
			}

			var assetRes *tpmtclient.TpmtAssetFindOneResp

			if item.AssetId != "" {
				//查询该点位关联的资产信息
				asset, err := l.svcCtx.TpmtAssetModel.FindOne(l.ctx, item.AssetId)
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

			data := &tpmtclient.TpmtMonitorPointRealTimeData{
				Id:                        item.Id,                         //监测点ID
				CreatedAt:                 item.CreatedAt.UnixMilli(),      //创建时间
				UpdatedAt:                 item.UpdatedAt.Time.UnixMilli(), //更新时间
				CreatedName:               item.CreatedName,                //创建人
				UpdatedName:               item.UpdatedName.String,         //更新人
				SerialNumber:              item.SerialNumber,               //编号
				Name:                      item.Name,                       //监测点名称
				RegisterAddress:           item.RegisterAddress,            //寄存器地址
				PointCollectorInstruction: item.PointCollectorInstruction,
				PointAnalysisRule:         item.PointAnalysisRule,
				PointType:                 item.PointType,             //类型：1:综保/2:局放/3:测温/4:微水/5:油色谱/6:机器人/7:其他
				PointCategory:             item.PointCategory,         //类别：1:遥信/2:遥测/3:遥脉
				PointGroup:                item.PointGroup,            //分组
				CircuitType:               item.CircuitType,           //回路类型
				YxDecode:                  item.YxDecode.String,       //遥信解译
				DataBits:                  item.DataBits,              //数据位
				Coefficient:               item.Coefficient,           //系数
				RetainDecimals:            item.RetainDecimals,        //保留小数位
				Unit:                      item.Unit.String,           //单位
				AlarmDuration:             item.AlarmDuration,         //持续时间
				AlarmUpValue:              item.AlarmUpValue,          //告警上限
				AlarmDownValue:            item.AlarmDownValue,        //告警下限
				WarningUpValue:            item.WarningUpValue,        //预警上限
				WarningDownValue:          item.WarningDownValue,      //预警下限
				IsDisplacementWarning:     item.IsDisplacementWarning, //变位预警 0 不启用 1:启用
				TpmtGatewayId:             item.TpmtGatewayId,         //网关ID
				AssetId:                   item.AssetId,
				Sort:                      item.Sort,     //排序
				ResultValue:               resultValue,   //实时数据
				UpdateTime:                updateTime,    //实时数据更新时间
				AlarmRuleInfo:             alarmRuleInfo, //告警内容
				Asset:                     assetRes,
			}
			chanData <- data
		}(item)

	}

	wg.Wait()
	close(chanData)

	var list []*tpmtclient.TpmtMonitorPointRealTimeData

	for i := range chanData {
		list = append(list, i)
	}

	return &tpmtclient.TpmtMonitorPointRealTimeListResp{
		Total: count,
		List:  list,
	}, nil
}
