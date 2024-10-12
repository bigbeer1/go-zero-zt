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

type TpmtGatewayDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewTpmtGatewayDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TpmtGatewayDeleteLogic {
	return &TpmtGatewayDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *TpmtGatewayDeleteLogic) TpmtGatewayDelete(in *tpmtclient.TpmtGatewayDeleteReq) (*tpmtclient.CommonResp, error) {

	res, err := l.svcCtx.TpmtGatewayModel.FindOne(l.ctx, in.Id)
	if err != nil {
		if errors.Is(err, sqlc.ErrNotFound) {
			return nil, fmt.Errorf("TpmtGateway没有该ID:%v", in.Id)
		}
		return nil, err
	}

	err = l.svcCtx.TpmtGatewayModel.Delete(l.ctx, res.Id)
	if err != nil {
		return nil, err
	}

	return &tpmtclient.CommonResp{}, nil
}
