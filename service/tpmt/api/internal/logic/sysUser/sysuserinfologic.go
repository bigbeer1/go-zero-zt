package sysUser

import (
	"context"
	"tpmt-zt/common"
	"tpmt-zt/common/msg"
	"tpmt-zt/service/authentication/authenticationclient"

	"github.com/jinzhu/copier"

	"tpmt-zt/service/tpmt/api/internal/svc"
	"tpmt-zt/service/tpmt/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SysUserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSysUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SysUserInfoLogic {
	return &SysUserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SysUserInfoLogic) SysUserInfo(req *types.SysUserInfoRequest) (resp *types.Response, err error) {

	res, err := l.svcCtx.AuthenticationRpc.SysUserFindOne(l.ctx, &authenticationclient.SysUserFindOneReq{
		Id: req.Id, // 用户ID
	})
	if err != nil {
		return nil, common.NewDefaultError(err.Error())
	}

	var result SysUserFindOneResp
	_ = copier.Copy(&result, res)

	return &types.Response{
		Code: 0,
		Msg:  msg.Success,
		Data: result,
	}, nil
}

type SysUserFindOneResp struct {
	Id          string `json:"id"`           // 用户ID,
	Account     string `json:"account"`      // 用户名,
	NickName    string `json:"nick_name"`    // 姓名,
	State       int64  `json:"state"`        // 状态 1:正常 2:停用 3:封禁,
	CreatedName string `json:"created_name"` // 创建人,
	CreatedAt   int64  `json:"created_at"`   // 创建时间,
	UpdatedName string `json:"updated_name"` // 更新人,
	UpdatedAt   int64  `json:"updated_at"`   // 更新时间
	RoleId      int64  `json:"role_id"`      // 角色ID
	RoleName    string `json:"role_name"`    // 角色名称
	RoleType    int64  `json:"role_type"`    // 角色类型
}
