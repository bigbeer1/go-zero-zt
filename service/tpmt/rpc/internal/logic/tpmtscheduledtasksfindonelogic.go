package logic

import (
	"context"
	"errors"
	"github.com/zeromicro/go-zero/core/stores/sqlc"

	"tpmt-zt/service/tpmt/rpc/internal/svc"
	"tpmt-zt/service/tpmt/rpc/tpmtclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type TpmtScheduledTasksFindOneLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewTpmtScheduledTasksFindOneLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TpmtScheduledTasksFindOneLogic {
	return &TpmtScheduledTasksFindOneLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *TpmtScheduledTasksFindOneLogic) TpmtScheduledTasksFindOne(in *tpmtclient.TpmtScheduledTasksFindOneReq) (*tpmtclient.TpmtScheduledTasksFindOneResp, error) {

	res, err := l.svcCtx.TpmtScheduledTasksModel.FindOne(l.ctx, in.Id)
	if err != nil {
		if errors.Is(err, sqlc.ErrNotFound) {
			return nil, errors.New("TpmtScheduledTasks没有该ID：" + in.Id)
		}
		return nil, err
	}

	return &tpmtclient.TpmtScheduledTasksFindOneResp{
		Id:                  res.Id,                         //定时任务ID
		CreatedAt:           res.CreatedAt.UnixMilli(),      //创建时间
		UpdatedAt:           res.UpdatedAt.Time.UnixMilli(), //更新时间
		CreatedName:         res.CreatedName,                //创建人
		UpdatedName:         res.UpdatedName.String,         //更新人
		SchedulerName:       res.SchedulerName,              //名称
		SchedulerCategory:   res.SchedulerCategory,          //类别 1:已接入任务, 2:自定义任务
		SchedulerTaskNumber: res.SchedulerTaskNumber,        //已接入任务号
		SchedulerType:       res.SchedulerType,              //类型 1:Http任务,2:Webservices任务
		IntervalTime:        res.IntervalTime,               //间隔时间按秒
		ErrorOrder:          res.ErrorOrder,                 //失败重新发送次数1-10次 不可超过10次
		FailIntervalTime:    res.FailIntervalTime,           //失败间隔时间按秒
		State:               res.State,                      //状态 1:启动  2:暂停
		SchedulerData:       res.SchedulerData,              //内容
	}, nil
}
