package logic

import (
	"context"
	"fmt"
	"time"
	"tpmt-zt/common/tdenginex"
	"tpmt-zt/service/archive/model"
	"tpmt-zt/service/archive/rpc/archiveclient"
	"tpmt-zt/service/archive/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type AlarmUpStateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAlarmUpStateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AlarmUpStateLogic {
	return &AlarmUpStateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 告警更新状态
func (l *AlarmUpStateLogic) AlarmUpState(in *archiveclient.AlarmUpStateReq) (*archiveclient.CommonResp, error) {

	alarmTddb := &model.TdDb{
		DbName:    "alarm_log." + "YcYm",
		TableName: "alarm_log.ptm_alarm_log",
	}
	alarmLog := model.AlarmLog{
		Ts:         time.Time{},
		Id:         in.Id,
		AlarmType:  99,
		AlarmGrade: 99,
		AlarmState: 99,
	}

	all, err := alarmLog.FindOne(l.ctx, l.svcCtx.Taos, alarmTddb)
	if err != nil {
		if err.Error() != tdenginex.ErrNotFoundTable {
			return nil, err
		}
	}

	if len(all) == 0 {
		return nil, fmt.Errorf("不存在这个告警ID")
	}

	// 实际上就一个
	for _, item := range all {
		if item.AlarmState != 0 {
			return &archiveclient.CommonResp{}, nil
		}

		alarmLogData := model.AlarmLog{
			Ts:           item.Ts,
			Id:           item.Id,
			Mid:          item.Mid,
			Name:         item.Name,
			AlarmType:    item.AlarmType,
			AlarmGrade:   item.AlarmGrade,
			AlarmContent: item.AlarmContent,
			AssetCode:    item.AssetCode,
			AlarmState:   1,
		}

		// 添加到时序数据库中去
		err = alarmLogData.Insert(l.ctx, l.svcCtx.Taos, alarmTddb)
		if err != nil {
			logx.Errorf("时序数据库更新告警数据失败！" + err.Error())
		}

	}

	return &archiveclient.CommonResp{}, nil
}
