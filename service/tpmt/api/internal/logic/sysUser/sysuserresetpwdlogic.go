package sysUser

import (
	"context"
	"tpmt-zt/common"
	"tpmt-zt/common/jwtx"
	"tpmt-zt/common/msg"
	"tpmt-zt/service/authentication/authenticationclient"

	"tpmt-zt/service/tpmt/api/internal/svc"
	"tpmt-zt/service/tpmt/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SysUserResetPwdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSysUserResetPwdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SysUserResetPwdLogic {
	return &SysUserResetPwdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SysUserResetPwdLogic) SysUserResetPwd(req *types.SysUserResetPwdRequest) (resp *types.Response, err error) {
	tokenData := jwtx.ParseToken(l.ctx)

	//  不能重置自己
	if tokenData.Uid == req.Id {
		return nil, common.NewDefaultError("该接口不能修改自身")
	}

	res, err := l.svcCtx.AuthenticationRpc.SysUserResetPwd(l.ctx, &authenticationclient.SysUserResetPwdReq{
		Id:          req.Id,
		Password:    req.Password,
		UpdatedName: tokenData.NickName,
	})

	if err != nil {
		return nil, common.NewDefaultError(err.Error())
	}

	return &types.Response{
		Code: 0,
		Msg:  msg.Success,
		Data: res.Password,
	}, nil
}
