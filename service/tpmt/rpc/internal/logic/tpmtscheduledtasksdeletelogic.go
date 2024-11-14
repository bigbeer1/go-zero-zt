package logic

import (
	"context"
	"errors"
	"github.com/zeromicro/go-zero/core/stores/sqlc"

	"tpmt-zt/service/tpmt/rpc/internal/svc"
	"tpmt-zt/service/tpmt/rpc/tpmtclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type TpmtScheduledTasksDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewTpmtScheduledTasksDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TpmtScheduledTasksDeleteLogic {
	return &TpmtScheduledTasksDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *TpmtScheduledTasksDeleteLogic) TpmtScheduledTasksDelete(in *tpmtclient.TpmtScheduledTasksDeleteReq) (*tpmtclient.CommonResp, error) {

	_, err := l.svcCtx.TpmtScheduledTasksModel.FindOne(l.ctx, in.Id)
	if err != nil {
		if errors.Is(err, sqlc.ErrNotFound) {
			return nil, errors.New("TpmtScheduledTasks没有该ID：" + in.Id)
		}
		return nil, err
	}

	err = l.svcCtx.TpmtScheduledTasksModel.Delete(l.ctx, in.Id)
	if err != nil {
		return nil, err
	}

	return &tpmtclient.CommonResp{}, nil
}
