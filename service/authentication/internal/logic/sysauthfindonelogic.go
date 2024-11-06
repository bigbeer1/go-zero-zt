package logic

import (
	"context"
	"errors"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"tpmt-zt/service/authentication/authentication"
	"tpmt-zt/service/authentication/internal/svc"
	"tpmt-zt/service/authentication/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type SysAuthFindOneLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSysAuthFindOneLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SysAuthFindOneLogic {
	return &SysAuthFindOneLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SysAuthFindOneLogic) SysAuthFindOne(in *authentication.SysAuthFindOneReq) (*authentication.SysAuthFindOneResp, error) {

	res, err := l.svcCtx.SysAuthModel.FindOne(l.ctx, in.Id)
	if err != nil {
		if errors.Is(err, sqlc.ErrNotFound) {
			return nil, fmt.Errorf("SysAuth没有该ID:%v", in.Id)
		}
		return nil, err
	}

	// 判断该数据是否被删除
	if res.DeletedAt.Valid == true {
		return nil, fmt.Errorf("SysAuth该ID已被删除：%v", in.Id)
	}

	// 查询用户角色信息
	AuthRole, err := l.svcCtx.SysUserRoleModel.FindByUserIdAndUserType(l.ctx, res.Id, 2)
	if err != nil {
		if !errors.Is(err, sqlc.ErrNotFound) {
			return nil, err
		}
	}

	var roleRes model.SysRole

	if err == nil {
		role, err := l.svcCtx.SysRoleModel.FindOne(l.ctx, AuthRole.RoleId)
		if err != nil {
			if errors.Is(err, sqlc.ErrNotFound) {
				return nil, fmt.Errorf("SysRole没有该ID:%v", role.Id)
			}
			return nil, err
		}

		// 判断该数据是否被删除
		if role.DeletedAt.Valid == true {
			return nil, fmt.Errorf("SysRole该ID已被删除：%v", role.Id)
		}
		roleRes = model.SysRole{
			Id:          role.Id,
			Name:        role.Name,
			Remark:      role.Remark,
			RoleType:    role.RoleType,
			CreatedName: role.CreatedName,
			CreatedAt:   role.CreatedAt,
			UpdatedName: role.UpdatedName,
			UpdatedAt:   role.UpdatedAt,
			DeletedAt:   role.DeletedAt,
			DeletedName: role.DeletedName,
		}
	}

	return &authentication.SysAuthFindOneResp{
		Id:          res.Id,                         //第三方用户ID
		CreatedAt:   res.CreatedAt.UnixMilli(),      //创建时间
		UpdatedAt:   res.UpdatedAt.Time.UnixMilli(), //更新时间
		CreatedName: res.CreatedName,                //创建人
		UpdatedName: res.UpdatedName.String,         //更新人
		NickName:    res.NickName,                   //机构名
		AuthToken:   res.AuthToken,                  //令牌
		State:       res.State,                      //状态 1:正常 2:停用 3:封禁
		RoleId:      roleRes.Id,
		RoleName:    roleRes.Name,
		RoleType:    roleRes.RoleType,
	}, nil

}
