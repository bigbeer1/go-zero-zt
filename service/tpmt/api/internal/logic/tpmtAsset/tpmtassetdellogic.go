package tpmtAsset

import (
	"context"
	"tpmt-zt/common"
	"tpmt-zt/common/msg"
	"tpmt-zt/service/tpmt/rpc/tpmtclient"

	"tpmt-zt/service/tpmt/api/internal/svc"
	"tpmt-zt/service/tpmt/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type TpmtAssetDelLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTpmtAssetDelLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TpmtAssetDelLogic {
	return &TpmtAssetDelLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TpmtAssetDelLogic) TpmtAssetDel(req *types.TpmtAssetDelRequest) (resp *types.Response, err error) {

	_, err = l.svcCtx.TpmtRpc.TpmtAssetDelete(l.ctx, &tpmtclient.TpmtAssetDeleteReq{
		Id: req.Id, // 资产ID
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
