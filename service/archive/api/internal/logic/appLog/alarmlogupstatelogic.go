package appLog

import (
	"context"
	"tpmt-zt/common"
	"tpmt-zt/common/msg"
	"tpmt-zt/service/archive/api/internal/svc"
	"tpmt-zt/service/archive/api/internal/types"
	"tpmt-zt/service/archive/rpc/archiveclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type AlarmLogUpStateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAlarmLogUpStateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AlarmLogUpStateLogic {
	return &AlarmLogUpStateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AlarmLogUpStateLogic) AlarmLogUpState(req *types.AlarmUpStateReqest) (resp *types.Response, err error) {

	_, err = l.svcCtx.ArchiveRpc.AlarmUpState(l.ctx, &archiveclient.AlarmUpStateReq{
		Id: req.Id,
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
