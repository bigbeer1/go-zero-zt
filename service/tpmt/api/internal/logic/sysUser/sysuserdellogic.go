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

type SysUserDelLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSysUserDelLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SysUserDelLogic {
	return &SysUserDelLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SysUserDelLogic) SysUserDel(req *types.SysUserDelRequest) (resp *types.Response, err error) {
	// 用户登录信息

	_, err = l.svcCtx.TpmtRpc.SysUserDelete(l.ctx, &tpmtclient.SysUserDeleteReq{
		Id:          req.Id,             // 用户ID
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
