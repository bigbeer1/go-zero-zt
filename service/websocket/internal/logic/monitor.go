package logic

import (
	"context"
	"fmt"
	"github.com/fasthttp/websocket"
	"github.com/zeromicro/go-zero/core/logx"
	"sync"
	"tpmt-zt/common/alarm"
	"tpmt-zt/common/datax"
	"tpmt-zt/common/eval"
	"tpmt-zt/common/global"
	"tpmt-zt/common/jsonx"
	"tpmt-zt/common/jwtx"
	"tpmt-zt/service/tpmt/model"
	"tpmt-zt/service/tpmt/rpc/tpmtclient"
	"tpmt-zt/service/websocket/internal/util"
)

func (l *Websocket) monitorRealData(ctx context.Context, conn *websocket.Conn, dataString string, tokenData *jwtx.TokenData) {
	var in []*global.SocketMonitor

	// 解析json
	err := jsonx.Str2Struct(dataString, &in)
	if err != nil {
		err = util.MonitorSendSocket(conn, fmt.Sprintf("error|json数据格式错误|%s|", dataString))
		if err != nil {
			logx.Errorf(err.Error())
		}
		return
	}

	/*
		协程并发查询提升速度
	*/
	var wg sync.WaitGroup
	// 协程获取返回数据
	chanData := make(chan *global.SocketMonitor, len(in))
	// 协程数控制
	var limit = make(chan struct{}, l.svcCtx.Config.RealTimeLimit)

	var list []*global.SocketMonitor

	for _, v := range in {
		limit <- struct{}{}
		wg.Add(1)
		go func(ctx context.Context, tpmtMonitorPointModel model.TpmtMonitorPointModel, item *global.SocketMonitor) {
			defer func() {
				wg.Done()
				<-limit
			}()
			monitorPoint, err := tpmtMonitorPointModel.FindOne(ctx, item.Id)
			if err != nil {
				return
			}

			var redisKey, resultValue string
			var updateTime int64

			redisKey = fmt.Sprintf("action_time:id:%v",
				item.Id)
			// 根据key获取redis中的实时数据
			temp := l.svcCtx.TpmtMonitorPointModel.GetRealData(ctx, redisKey)

			// 获取实时数据和时间
			resultValue, updateTime = eval.MonitorValueEval(temp, monitorPoint.Coefficient)

			// 告警等级  0正常 1预警 2告警
			// 初始化告警状态为 无  0正常  1越上线  2越下线
			var alarmRuleInfo = &tpmtclient.AlarmRuleInfo{
				Level:    0,
				RuleType: 0,
			}

			// 查询告警规则
			resultValueFloat64, err := datax.ToFloat64(resultValue)
			if err == nil {
				info := alarm.CheckAlarmRuleYc(monitorPoint, resultValueFloat64)
				if info != nil {
					alarmRuleInfo = &tpmtclient.AlarmRuleInfo{
						Level:    info.Level,
						RuleType: info.RuleType,
						RuleData: info.RuleData,
					}
				}
			}

			var data = &global.SocketMonitor{
				Id:            monitorPoint.Id,
				ResultValue:   resultValue,
				UpdateTime:    updateTime,
				Level:         alarmRuleInfo.Level,
				RuleType:      alarmRuleInfo.RuleType,
				PointCategory: monitorPoint.PointCategory,
				Unit:          monitorPoint.Unit.String,
			}
			chanData <- data
		}(ctx, l.svcCtx.TpmtMonitorPointModel, v)

	}

	wg.Wait()
	close(chanData)
	for i := range chanData {
		list = append(list, i)
	}
	reData, _ := jsonx.ToJSONStr(list)

	err = util.MonitorSendSocket(conn, fmt.Sprintf("monitor|%v|", reData))
	if err != nil {
		logx.Errorf(err.Error())
	}
	return

}
