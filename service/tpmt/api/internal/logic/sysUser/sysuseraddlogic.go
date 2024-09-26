package sysUser

import (
	"context"
	"tpmt-zt/common"
	"tpmt-zt/common/msg"
	"tpmt-zt/common/tokenData"
	"tpmt-zt/service/tpmt/rpc/tpmtclient"

	"tpmt-zt/service/tpmt/api/internal/svc"
	"tpmt-zt/service/tpmt/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SysUserAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSysUserAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SysUserAddLogic {
	return &SysUserAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SysUserAddLogic) SysUserAdd(req *types.SysUserAddRequest) (resp *types.Response, err error) {

	_, err = l.svcCtx.TpmtRpc.SysUserAdd(l.ctx, &tpmtclient.SysUserAddReq{
		Account:     req.Account,        // 用户名
		NickName:    req.NickName,       // 姓名
		Password:    req.Password,       // 密码
		State:       req.State,          // 状态 1:正常 2:停用 3:封禁
		CreatedName: tokenData.NickName, // 创建人
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
