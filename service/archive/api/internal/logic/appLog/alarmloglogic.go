package appLog

import (
	"context"
	"github.com/jinzhu/copier"
	"tpmt-zt/common"
	"tpmt-zt/common/msg"
	"tpmt-zt/service/archive/api/internal/svc"
	"tpmt-zt/service/archive/api/internal/types"
	"tpmt-zt/service/archive/rpc/archiveclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type AlarmLogLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAlarmLogLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AlarmLogLogic {
	return &AlarmLogLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AlarmLogLogic) AlarmLog(req *types.AlarmLogReqest) (resp *types.Response, err error) {

	res, err := l.svcCtx.ArchiveRpc.AlarmLogFindList(l.ctx, &archiveclient.AlarmLogFindListReq{
		Current:       req.Current,
		PageSize:      req.PageSize,
		StartTime:     req.CreatedStartTime,
		EndTime:       req.CreatedEndTime,
		AlarmCategory: req.AlarmCategory,
		Id:            req.Id,
		Mid:           req.Mid,
		AlarmType:     req.AlarmType,
		AlarmGrade:    req.AlarmGrade,
		AssetCode:     req.AssetCode,
		AlarmState:    req.AlarmState,
	})

	if err != nil {
		return nil, common.NewDefaultError(err.Error())
	}

	var result AlarmLogFindListResp

	_ = copier.Copy(&result, res)

	return &types.Response{
		Code: 0,
		Msg:  msg.Success,
		Data: result,
	}, nil
}

type AlarmLogFindListResp struct {
	Total int64                   `json:"total"` //总数据量
	List  []*AlarmLogFindListData `json:"list"`  //数据
}

type AlarmLogFindListData struct {
	Ts           int64  `json:"ts"`            // 开始时间
	Id           string `json:"id"`            // 告警ID
	Mid          string `json:"mid"`           // mid 预警设备ID
	Name         string `json:"name"`          // 设备名称
	AlarmType    int64  `json:"alarm_type"`    // 类型：1 越上限/2 越下限/ 3 变位 / 4 网关下线
	AlarmGrade   int64  `json:"alarm_grade"`   // 等级：1 预警/2 告警 /3 提醒
	AlarmContent string `json:"alarm_content"` // 数据  返回内容
	AssetCode    string `json:"asset_code"`    // 设备名称
	AlarmState   int64  `json:"alarm_state"`   // 状态 0 未读  1已读  2已确认
}
