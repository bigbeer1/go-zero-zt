package sysRole

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

type SysRoleUpLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSysRoleUpLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SysRoleUpLogic {
	return &SysRoleUpLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SysRoleUpLogic) SysRoleUp(req *types.SysRoleUpRequest) (resp *types.Response, err error) {
	// 用户登录信息
	tokenData := jwtx.ParseToken(l.ctx)

	_, err = l.svcCtx.TpmtRpc.SysRoleUpdate(l.ctx, &tpmtclient.SysRoleUpdateReq{
		Id:           req.Id,              // 角色ID
		Name:         req.Name,            // 角色名称
		Remark:       req.Remark,          // 备注
		RoleType:     req.RoleType,        // 角色类型 1:管理员角色  2:普通角色  3:第三方角色
		UpdatedName:  tokenData.NickName,  // 更新人
		MenuIds:      req.SysMenuIds,      // 菜单IDS
		InterfaceIds: req.SysInterfaceIds, // 接口IDS
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
