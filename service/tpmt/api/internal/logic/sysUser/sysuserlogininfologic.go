package sysUser

import (
	"context"
	"github.com/jinzhu/copier"
	"tpmt-zt/common"
	"tpmt-zt/common/jwtx"
	"tpmt-zt/common/msg"
	"tpmt-zt/service/tpmt/api/internal/logic/sysInterface"
	"tpmt-zt/service/tpmt/api/internal/logic/sysMenu"
	"tpmt-zt/service/tpmt/api/internal/logic/sysRole"
	"tpmt-zt/service/tpmt/rpc/tpmtclient"

	"tpmt-zt/service/tpmt/api/internal/svc"
	"tpmt-zt/service/tpmt/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SysUserLoginInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSysUserLoginInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SysUserLoginInfoLogic {
	return &SysUserLoginInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SysUserLoginInfoLogic) SysUserLoginInfo() (resp *types.Response, err error) {
	// token信息
	tokenData := jwtx.ParseToken(l.ctx)

	// 用户信息
	user, err := l.svcCtx.TpmtRpc.SysUserFindOne(l.ctx, &tpmtclient.SysUserFindOneReq{
		Id: tokenData.Uid, // 用户ID
	})
	if err != nil {
		return nil, common.NewDefaultError(err.Error())
	}

	var result SysUserLoginInfoResp

	_ = copier.Copy(&result.User, user)

	if user.RoleId != 0 {
		// 角色信息
		role, err := l.svcCtx.TpmtRpc.SysRoleFindOne(l.ctx, &tpmtclient.SysRoleFindOneReq{
			Id: user.RoleId, // 角色ID
		})
		if err == nil {
			result.Role = &sysRole.SysRoleDataList{
				Id:          role.Id,
				Name:        role.Name,
				Remark:      role.Remark,
				RoleType:    role.RoleType,
				CreatedName: role.CreatedName,
				CreatedAt:   role.CreatedAt,
				UpdatedName: role.UpdatedName,
				UpdatedAt:   role.UpdatedAt,
			}
		}

		// 角色的菜单
		menuResp, err := l.svcCtx.TpmtRpc.SysMenuByRoleId(l.ctx, &tpmtclient.SysMenuByRoleIdReq{
			RoleId: user.RoleId,
		})

		if err == nil {
			_ = copier.Copy(&result.MenuList, menuResp.List)
		}

		// 角色的接口
		interfaceRep, err := l.svcCtx.TpmtRpc.SysInterfaceByRoleId(l.ctx, &tpmtclient.SysInterfaceByRoleIdReq{
			RoleId: user.RoleId,
		})

		if err == nil {
			_ = copier.Copy(&result.InterfaceList, interfaceRep.List)
		}
	}

	return &types.Response{
		Code: 0,
		Msg:  msg.Success,
		Data: result,
	}, nil
}

type SysUserLoginInfoResp struct {
	User          SysUserFindOneResp                   `json:"user"`           // 用户信息
	Role          *sysRole.SysRoleDataList             `json:"role"`           // 角色信息
	MenuList      []*sysMenu.SysMenuDataList           `json:"menu_list"`      // 菜单list
	InterfaceList []*sysInterface.SysInterfaceDataList `json:"interface_list"` // 接口list
}
