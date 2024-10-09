package sysRole

import (
	"context"
	"github.com/jinzhu/copier"
	"tpmt-zt/common"
	"tpmt-zt/common/msg"
	"tpmt-zt/service/tpmt/rpc/tpmtclient"

	"tpmt-zt/service/tpmt/api/internal/svc"
	"tpmt-zt/service/tpmt/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SysRoleListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSysRoleListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SysRoleListLogic {
	return &SysRoleListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SysRoleListLogic) SysRoleList(req *types.SysRoleListRequest) (resp *types.Response, err error) {

	all, err := l.svcCtx.TpmtRpc.SysRoleList(l.ctx, &tpmtclient.SysRoleListReq{
		Current:  req.Current,  // 页码
		PageSize: req.PageSize, // 页数
		Name:     req.Name,     // 角色名称
		Remark:   req.Remark,   // 备注
		RoleType: req.RoleType, // 角色类型 1:管理员角色  2:普通角色  3:第三方角色
	})
	if err != nil {
		return nil, common.NewDefaultError(err.Error())
	}

	var result SysRoleListResp
	_ = copier.Copy(&result, all)

	return &types.Response{
		Code: 0,
		Msg:  msg.Success,
		Data: result,
	}, nil
}

type SysRoleListResp struct {
	Total int64              `json:"total"`
	List  []*SysRoleDataList `json:"list"`
}

type SysRoleDataList struct {
	Id          int64  `json:"id"`           // 角色ID,
	Name        string `json:"name"`         // 角色名称,
	Remark      string `json:"remark"`       // 备注,
	RoleType    int64  `json:"role_type"`    // 角色类型 1:管理员角色  2:普通角色  3:第三方角色,
	CreatedName string `json:"created_name"` // 创建人,
	CreatedAt   int64  `json:"created_at"`   // 创建时间,
	UpdatedName string `json:"updated_name"` // 更新人,
	UpdatedAt   int64  `json:"updated_at"`   // 更新时间
}
