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

type TpmtAssetDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewTpmtAssetDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TpmtAssetDeleteLogic {
	return &TpmtAssetDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *TpmtAssetDeleteLogic) TpmtAssetDelete(in *tpmtclient.TpmtAssetDeleteReq) (*tpmtclient.CommonResp, error) {

	res, err := l.svcCtx.TpmtAssetModel.FindOne(l.ctx, in.Id)
	if err != nil {
		if errors.Is(err, sqlc.ErrNotFound) {
			return nil, fmt.Errorf("TpmtAsset没有该ID:%v", in.Id)
		}
		return nil, err
	}

	err = l.svcCtx.TpmtAssetModel.Delete(l.ctx, res.Id)
	if err != nil {
		return nil, err
	}

	return &tpmtclient.CommonResp{}, nil
}
