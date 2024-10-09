package sysRole

import (
	"context"
	"github.com/jinzhu/copier"
	"tpmt-zt/common"
	"tpmt-zt/common/msg"
	"tpmt-zt/service/tpmt/api/internal/logic/sysInterface"
	"tpmt-zt/service/tpmt/api/internal/logic/sysMenu"
	"tpmt-zt/service/tpmt/rpc/tpmtclient"

	"tpmt-zt/service/tpmt/api/internal/svc"
	"tpmt-zt/service/tpmt/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SysRoleInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSysRoleInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SysRoleInfoLogic {
	return &SysRoleInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SysRoleInfoLogic) SysRoleInfo(req *types.SysRoleInfoRequest) (resp *types.Response, err error) {
	res, err := l.svcCtx.TpmtRpc.SysRoleFindOne(l.ctx, &tpmtclient.SysRoleFindOneReq{
		Id: req.Id, // 角色ID
	})

	if err != nil {
		return nil, common.NewDefaultError(err.Error())
	}

	menuResp, err := l.svcCtx.TpmtRpc.SysMenuByRoleId(l.ctx, &tpmtclient.SysMenuByRoleIdReq{
		RoleId: req.Id,
	})

	interfaceRep, err := l.svcCtx.TpmtRpc.SysInterfaceByRoleId(l.ctx, &tpmtclient.SysInterfaceByRoleIdReq{
		RoleId: req.Id,
	})

	var result SysRoleFindOneResp
	_ = copier.Copy(&result, res)

	_ = copier.Copy(&result.MenuList, menuResp.List)
	_ = copier.Copy(&result.InterfaceList, interfaceRep.List)

	return &types.Response{
		Code: 0,
		Msg:  msg.Success,
		Data: result,
	}, nil
}

type SysRoleFindOneResp struct {
	Id            int64                                `json:"id"`             // 角色ID,
	Name          string                               `json:"name"`           // 角色名称,
	Remark        string                               `json:"remark"`         // 备注,
	RoleType      int64                                `json:"role_type"`      // 角色类型 1:管理员角色  2:普通角色  3:第三方角色,
	CreatedName   string                               `json:"created_name"`   // 创建人,
	CreatedAt     int64                                `json:"created_at"`     // 创建时间,
	UpdatedName   string                               `json:"updated_name"`   // 更新人,
	UpdatedAt     int64                                `json:"updated_at"`     // 更新时间
	MenuList      []*sysMenu.SysMenuDataList           `json:"menu_list"`      // 菜单list
	InterfaceList []*sysInterface.SysInterfaceDataList `json:"interface_list"` // 接口list
}
