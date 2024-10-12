package logic

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"time"
	"tpmt-zt/service/tpmt/model"

	"tpmt-zt/service/tpmt/rpc/internal/svc"
	"tpmt-zt/service/tpmt/rpc/tpmtclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type SysUserUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSysUserUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SysUserUpdateLogic {
	return &SysUserUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SysUserUpdateLogic) SysUserUpdate(in *tpmtclient.SysUserUpdateReq) (*tpmtclient.CommonResp, error) {
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

	// 姓名
	if len(in.NickName) > 0 {
		res.NickName = in.NickName
	}

	// 状态 1:正常 2:停用 3:封禁
	if in.State != 0 {
		res.State = in.State
	}

	res.UpdatedName.String = in.UpdatedName
	res.UpdatedName.Valid = true
	res.UpdatedAt.Time = time.Now()
	res.UpdatedAt.Valid = true

	// 开启事务
	err = l.svcCtx.SysUserModel.TransCtx(l.ctx, func(ctx context.Context, session sqlx.Session) error {

		// 更新用户信息
		err = l.svcCtx.SysUserModel.TransUpdate(l.ctx, session, res)
		if err != nil {
			return err
		}

		// 如果查询用户角色中间表
		userRole, err := l.svcCtx.SysUserRoleModel.FindByUserId(l.ctx, res.Id)
		if err != nil {
			if !errors.Is(err, sqlc.ErrNotFound) {
				return err
			}
		}

		// 用户角色中间表为空的情况
		if errors.Is(err, sqlc.ErrNotFound) {
			if in.RoleId != 0 {
				// 查询角色信息
				role, err := l.svcCtx.SysRoleModel.FindOne(l.ctx, in.RoleId)
				if err != nil {
					if errors.Is(err, sqlc.ErrNotFound) {
						return fmt.Errorf("SysRole没有该ID:%v", role.Id)
					}
					return err
				}

				// 判断角色类型是否是普通角色
				if role.RoleType != 2 {
					return fmt.Errorf("该角色类型无法添加到user")
				}

				// 判断该数据是否被删除
				if role.DeletedAt.Valid == true {
					return fmt.Errorf("SysRole该ID已被删除：%v", role.Id)
				}

				// 添加用户角色中间表信息
				// 中间表添加角色和用户的关系
				_, err = l.svcCtx.SysUserRoleModel.TransInsert(l.ctx, session, &model.SysUserRole{
					UserId: res.Id,
					RoleId: role.Id,
					UserType: sql.NullInt64{
						Int64: role.RoleType,
						Valid: role.RoleType != 0,
					},
					CreatedName: in.UpdatedName,
					CreatedAt:   time.Now(),
				})

				if err != nil {
					return err
				}
			}
		} else {
			// 如果角色ID改变 我们需要进行更换
			if userRole.RoleId != in.RoleId {
				if in.RoleId == 0 {
					// 直接删除用户角色中间表信息
					// 删除用户角色中间表信息
					err := l.svcCtx.SysUserRoleModel.TransDelete(l.ctx, userRole)
					if err != nil {
						return err
					}
				} else {
					// 查询角色信息
					role, err := l.svcCtx.SysRoleModel.FindOne(l.ctx, in.RoleId)
					if err != nil {
						if errors.Is(err, sqlc.ErrNotFound) {
							return fmt.Errorf("SysRole没有该ID:%v", role.Id)
						}
						return err
					}

					// 判断角色类型是否是普通角色
					if role.RoleType != 2 {
						return fmt.Errorf("该角色类型无法添加到user")
					}

					// 判断该数据是否被删除
					if role.DeletedAt.Valid == true {
						return fmt.Errorf("SysRole该ID已被删除：%v", role.Id)
					}

					// 删除用户角色中间表信息
					err = l.svcCtx.SysUserRoleModel.TransDelete(l.ctx, userRole)
					if err != nil {
						return err
					}

					// 添加用户角色中间表信息
					// 中间表添加角色和用户的关系
					_, err = l.svcCtx.SysUserRoleModel.TransInsert(l.ctx, session, &model.SysUserRole{
						UserId: res.Id,
						RoleId: role.Id,
						UserType: sql.NullInt64{
							Int64: role.RoleType,
							Valid: role.RoleType != 0,
						},
						CreatedName: in.UpdatedName,
						CreatedAt:   time.Now(),
					})

					if err != nil {
						return err
					}

				}
			}
		}
		return nil

	})

	if err != nil {
		return nil, err
	}

	return &tpmtclient.CommonResp{}, nil
}
