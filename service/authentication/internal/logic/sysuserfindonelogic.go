package logic

import (
	"context"
	"errors"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"tpmt-zt/service/authentication/authenticationclient"
	"tpmt-zt/service/authentication/internal/svc"
	"tpmt-zt/service/authentication/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type SysUserFindOneLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSysUserFindOneLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SysUserFindOneLogic {
	return &SysUserFindOneLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SysUserFindOneLogic) SysUserFindOne(in *authenticationclient.SysUserFindOneReq) (*authenticationclient.SysUserFindOneResp, error) {
	res, err := l.svcCtx.SysUserModel.FindOne(l.ctx, in.Id)
	if err != nil {
		if errors.Is(err, sqlc.ErrNotFound) {
			return nil, fmt.Errorf("SysUser没有该ID:%v", in.Id)
		}
		return nil, err
	}

	// 判断该数据是否被删除
	if res.DeletedAt.Valid == true {
		return nil, errors.New("SysUser该ID已被删除：" + in.Id)
	}

	// 查询用户角色信息
	userRole, err := l.svcCtx.SysUserRoleModel.FindByUserIdAndUserType(l.ctx, res.Id, 1)
	if err != nil {
		if !errors.Is(err, sqlc.ErrNotFound) {
			return nil, err
		}
	}

	var roleRes model.SysRole

	if err == nil {
		role, err := l.svcCtx.SysRoleModel.FindOne(l.ctx, userRole.RoleId)
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

	return &authenticationclient.SysUserFindOneResp{
		Id:          res.Id,                         //用户ID
		Account:     res.Account,                    //用户名
		NickName:    res.NickName,                   //姓名
		State:       res.State,                      //状态 1:正常 2:停用 3:封禁
		CreatedName: res.CreatedName,                //创建人
		CreatedAt:   res.CreatedAt.UnixMilli(),      //创建时间
		UpdatedName: res.UpdatedName.String,         //更新人
		UpdatedAt:   res.UpdatedAt.Time.UnixMilli(), //更新时间
		RoleId:      roleRes.Id,                     //角色ID
		RoleName:    roleRes.Name,                   //角色名称
		RoleType:    roleRes.RoleType,               //角色类型
	}, nil
}
