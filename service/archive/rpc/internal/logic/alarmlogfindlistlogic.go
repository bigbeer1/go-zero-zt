package logic

import (
	"context"
	"fmt"
	"tpmt-zt/common/tdenginex"
	"tpmt-zt/service/archive/model"
	"tpmt-zt/service/archive/rpc/archiveclient"
	"tpmt-zt/service/archive/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type AlarmLogFindListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAlarmLogFindListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AlarmLogFindListLogic {
	return &AlarmLogFindListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 告警日志
func (l *AlarmLogFindListLogic) AlarmLogFindList(in *archiveclient.AlarmLogFindListReq) (*archiveclient.AlarmLogFindListResp, error) {

	var alarmTddb *model.TdDb

	switch in.AlarmCategory {
	case 1: // 告警预警
		alarmTddb = &model.TdDb{
			DbName:    "alarm_log." + "YcYm",
			TableName: "alarm_log.ptm_alarm_log",
		}
	case 2: // yx变位提醒
		alarmTddb = &model.TdDb{
			DbName:    "alarm_log." + "Yx",
			TableName: "alarm_log.ptm_alarm_log",
		}
	case 3: // 网关下线提醒
		alarmTddb = &model.TdDb{
			DbName:    "alarm_log." + "GateWay",
			TableName: "alarm_log.ptm_alarm_log",
		}
	default:
		return nil, fmt.Errorf("AlarmCategory必须是1,2,3")
	}

	alarmLog := model.AlarmLog{
		Id:           in.Id,
		Mid:          in.Mid,
		AlarmType:    in.AlarmType,
		AlarmGrade:   in.AlarmGrade,
		AlarmContent: "",
		AssetCode:    in.AssetCode,
		AlarmState:   in.AlarmState,
	}

	all, err := alarmLog.FindList(l.ctx, l.svcCtx.Taos, alarmTddb, in.Current, in.PageSize, in.StartTime, in.EndTime)
	if err != nil {
		if err.Error() != tdenginex.ErrNotFoundTable {
			return nil, err
		}
	}

	total := alarmLog.Count(l.ctx, l.svcCtx.Taos, alarmTddb, in.StartTime, in.EndTime)

	var datas []*archiveclient.AlarmLogFindListData
	for _, item := range all {
		datas = append(datas, &archiveclient.AlarmLogFindListData{
			Ts:           item.Ts.UnixMilli(),
			Id:           item.Id,
			Mid:          item.Mid,
			Name:         item.Name,
			AlarmType:    item.AlarmType,
			AlarmGrade:   item.AlarmGrade,
			AlarmContent: item.AlarmContent,
			AssetCode:    item.AssetCode,
			AlarmState:   item.AlarmState,
		})

	}

	return &archiveclient.AlarmLogFindListResp{
		Total: total,
		List:  datas,
	}, nil
}
