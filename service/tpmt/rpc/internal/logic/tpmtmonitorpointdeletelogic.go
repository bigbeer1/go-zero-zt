package logic

import (
	"context"
	"errors"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlc"

	"tpmt-zt/service/tpmt/rpc/internal/svc"
	"tpmt-zt/service/tpmt/rpc/tpmtclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type TpmtMonitorPointDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewTpmtMonitorPointDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TpmtMonitorPointDeleteLogic {
	return &TpmtMonitorPointDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *TpmtMonitorPointDeleteLogic) TpmtMonitorPointDelete(in *tpmtclient.TpmtMonitorPointDeleteReq) (*tpmtclient.CommonResp, error) {

	res, err := l.svcCtx.TpmtMonitorPointModel.FindOne(l.ctx, in.Id)
	if err != nil {
		if errors.Is(err, sqlc.ErrNotFound) {
			return nil, fmt.Errorf("TpmtMonitorPoint没有该ID:%v", in.Id)
		}
		return nil, err
	}

	err = l.svcCtx.TpmtMonitorPointModel.Delete(l.ctx, res.Id)
	if err != nil {
		return nil, err
	}

	return &tpmtclient.CommonResp{}, nil
}
