package sysAuth

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

type SysAuthUpLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSysAuthUpLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SysAuthUpLogic {
	return &SysAuthUpLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SysAuthUpLogic) SysAuthUp(req *types.SysAuthUpRequest) (resp *types.Response, err error) {
	// 用户登录信息
	tokenData := jwtx.ParseToken(l.ctx)

	_, err = l.svcCtx.TpmtRpc.SysAuthUpdate(l.ctx, &tpmtclient.SysAuthUpdateReq{
		Id:          req.Id,             // 第三方用户ID
		UpdatedName: tokenData.NickName, // 更新人
		State:       req.State,          // 状态 1:正常 2:停用 3:封禁
		RoleId:      req.RoldId,         // 角色ID
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
