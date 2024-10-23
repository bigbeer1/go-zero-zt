package store

import (
	"context"
	"fmt"
	"github.com/hibiken/asynq"
	"github.com/zeromicro/go-zero/core/jsonx"
	"github.com/zeromicro/go-zero/core/logx"
	"time"
	"tpmt-zt/common/datax"
	"tpmt-zt/common/eval"
	archivemodel "tpmt-zt/service/archive/model"
	"tpmt-zt/service/asynq/asynq-server/internal/svc"
	"tpmt-zt/service/tpmt/model"
)

type dataSet struct {
	svcCtx *svc.ServiceContext
}

func NewDataSetHandler(svcCtx *svc.ServiceContext) *dataSet {
	return &dataSet{
		svcCtx: svcCtx,
	}
}

func (l *dataSet) ProcessTask(ctx context.Context, t *asynq.Task) error {
	var data model.TpmtMonitorPoint
	if err := jsonx.Unmarshal(t.Payload(), &data); err != nil {
		logx.Errorf("报警json DataSet 转换失败 : ", t.Payload())
		return nil
	}

	redisKey := fmt.Sprintf("action_time:id:%v",
		data.Id)
	// 根据key获取redis中的实时数据
	temp := l.svcCtx.TpmtMonitorPointModel.GetRealData(ctx, redisKey)

	// 获取实时数据和时间
	resultValue, updateTime := eval.MonitorValueEval(temp, 1)

	// 如果没有数据则不存储
	if resultValue == "-" || resultValue == "" || updateTime < 0 {
		return nil
	}

	// 存储时序数据库
	dataKey := fmt.Sprintf("sequential_id_%v",
		data.Id)

	tddb := &archivemodel.TdDb{
		DbName:    "monitor_point." + dataKey,
		TableName: "monitor_point.tpmt_monitor_point",
	}

	dataFloat64, _ := datax.ToFloat64(resultValue)

	// 存储时序数据库记录检测值
	tdMonitor := archivemodel.TdMonitor{
		Ts:   time.UnixMilli(updateTime),
		Data: dataFloat64,
	}

	err := tdMonitor.Insert(ctx, l.svcCtx.Taos, tddb)
	if err != nil {
		logx.Errorf("时序数据库存储失败！" + err.Error())
	}

	return nil
}
