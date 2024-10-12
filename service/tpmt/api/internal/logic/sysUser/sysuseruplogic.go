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

type SysUserUpLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSysUserUpLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SysUserUpLogic {
	return &SysUserUpLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SysUserUpLogic) SysUserUp(req *types.SysUserUpRequest) (resp *types.Response, err error) {

	tokenData := jwtx.ParseToken(l.ctx)

	//  不能重置自己
	if tokenData.Uid == req.Id {
		return nil, common.NewDefaultError("该接口不能修改自身")
	}

	_, err = l.svcCtx.TpmtRpc.SysUserUpdate(l.ctx, &tpmtclient.SysUserUpdateReq{
		Id:          req.Id,             // 用户ID
		NickName:    req.NickName,       // 姓名
		State:       req.State,          // 状态 1:正常 2:停用 3:封禁
		UpdatedName: tokenData.NickName, // 更新人
		RoleId:      req.RoldId,         // 角色Id
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
