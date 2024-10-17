package tpmtMonitorPoint

import (
	"context"
	"tpmt-zt/common"
	"tpmt-zt/common/msg"
	"tpmt-zt/service/tpmt/rpc/tpmtclient"

	"tpmt-zt/service/tpmt/api/internal/svc"
	"tpmt-zt/service/tpmt/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type TpmtMonitorPointDelLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTpmtMonitorPointDelLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TpmtMonitorPointDelLogic {
	return &TpmtMonitorPointDelLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TpmtMonitorPointDelLogic) TpmtMonitorPointDel(req *types.TpmtMonitorPointDelRequest) (resp *types.Response, err error) {
	// 用户登录信息
	//tokenData := jwtx.ParseToken(l.ctx)

	_, err = l.svcCtx.TpmtRpc.TpmtMonitorPointDelete(l.ctx, &tpmtclient.TpmtMonitorPointDeleteReq{
		Id: req.Id, // 监测点ID
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
