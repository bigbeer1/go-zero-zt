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

type SysAuthAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSysAuthAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SysAuthAddLogic {
	return &SysAuthAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SysAuthAddLogic) SysAuthAdd(req *types.SysAuthAddRequest) (resp *types.Response, err error) {
	// 用户登录信息
	tokenData := jwtx.ParseToken(l.ctx)

	_, err = l.svcCtx.TpmtRpc.SysAuthAdd(l.ctx, &tpmtclient.SysAuthAddReq{
		CreatedName: tokenData.NickName, // 创建人
		NickName:    req.NickName,       // 机构名
		AuthToken:   req.AuthToken,      // 令牌
		State:       req.State,          // 状态 1:正常 2:停用 3:封禁
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
