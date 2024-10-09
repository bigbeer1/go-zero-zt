package sysUser

import (
	"context"
	"tpmt-zt/common"
	"tpmt-zt/common/jwtx"
	"tpmt-zt/common/msg"
	"tpmt-zt/service/tpmt/rpc/tpmtclient"

	"tpmt-zt/service/tpmt/api/internal/svc"
	"tpmt-zt/service/tpmt/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SysUserUpMyInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSysUserUpMyInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SysUserUpMyInfoLogic {
	return &SysUserUpMyInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SysUserUpMyInfoLogic) SysUserUpMyInfo(req *types.SysUserUpMyInfoRequest) (resp *types.Response, err error) {

	tokenData := jwtx.ParseToken(l.ctx)

	_, err = l.svcCtx.TpmtRpc.SysUserUpdate(l.ctx, &tpmtclient.SysUserUpdateReq{
		Id:          tokenData.Uid, // 用户ID
		NickName:    req.NickName,  // 姓名
		UpdatedName: req.NickName,  // 更新人
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
