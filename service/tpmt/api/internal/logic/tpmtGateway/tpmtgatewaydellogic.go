package tpmtGateway

import (
	"context"
	"tpmt-zt/common"
	"tpmt-zt/common/msg"
	"tpmt-zt/service/tpmt/rpc/tpmtclient"

	"tpmt-zt/service/tpmt/api/internal/svc"
	"tpmt-zt/service/tpmt/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type TpmtGatewayDelLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTpmtGatewayDelLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TpmtGatewayDelLogic {
	return &TpmtGatewayDelLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TpmtGatewayDelLogic) TpmtGatewayDel(req *types.TpmtGatewayDelRequest) (resp *types.Response, err error) {
	// 用户登录信息
	//tokenData := jwtx.ParseToken(l.ctx)

	_, err = l.svcCtx.TpmtRpc.TpmtGatewayDelete(l.ctx, &tpmtclient.TpmtGatewayDeleteReq{
		Id: req.Id, // 采集器ID/网关
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
