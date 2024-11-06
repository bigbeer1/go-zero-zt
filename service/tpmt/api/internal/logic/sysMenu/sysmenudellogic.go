package sysMenu

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

type SysMenuDelLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSysMenuDelLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SysMenuDelLogic {
	return &SysMenuDelLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SysMenuDelLogic) SysMenuDel(req *types.SysMenuDelRequest) (resp *types.Response, err error) {
	// 用户登录信息
	tokenData := jwtx.ParseToken(l.ctx)

	_, err = l.svcCtx.AuthenticationRpc.SysMenuDelete(l.ctx, &authenticationclient.SysMenuDeleteReq{
		Id:          req.Id,             // 菜单ID
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
