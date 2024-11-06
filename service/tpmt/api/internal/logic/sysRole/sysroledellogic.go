package sysRole

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

type SysRoleDelLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSysRoleDelLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SysRoleDelLogic {
	return &SysRoleDelLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SysRoleDelLogic) SysRoleDel(req *types.SysRoleDelRequest) (resp *types.Response, err error) {
	// 用户登录信息
	tokenData := jwtx.ParseToken(l.ctx)

	_, err = l.svcCtx.AuthenticationRpc.SysRoleDelete(l.ctx, &authenticationclient.SysRoleDeleteReq{
		Id:          req.Id,             // 角色ID
		DeletedName: tokenData.NickName, // 删除人
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
