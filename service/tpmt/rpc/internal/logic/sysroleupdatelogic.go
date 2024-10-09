package logic

import (
	"context"
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

type SysRoleUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSysRoleUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SysRoleUpdateLogic {
	return &SysRoleUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SysRoleUpdateLogic) SysRoleUpdate(in *tpmtclient.SysRoleUpdateReq) (*tpmtclient.CommonResp, error) {
	res, err := l.svcCtx.SysRoleModel.FindOne(l.ctx, in.Id)
	if err != nil {
		if errors.Is(err, sqlc.ErrNotFound) {
			return nil, fmt.Errorf("SysRole没有该ID: %v", in.Id)
		}
		return nil, err
	}

	// 判断该数据是否被删除
	if res.DeletedAt.Valid == true {
		return nil, fmt.Errorf("SysRole该ID已被删除： %v", in.Id)
	}

	// 角色名称
	if len(in.Name) > 0 {
		res.Name = in.Name
	}
	// 备注
	if len(in.Remark) > 0 {
		res.Remark.String = in.Remark
		res.Remark.Valid = true
	}
	// 角色类型 1:管理员角色  2:普通角色  3:第三方角色
	if in.RoleType != 0 {
		res.RoleType = in.RoleType
	}

	res.UpdatedName.String = in.UpdatedName
	res.UpdatedName.Valid = true
	res.UpdatedAt.Time = time.Now()
	res.UpdatedAt.Valid = true

	err = l.svcCtx.SysRoleModel.TransCtx(l.ctx, func(ctx context.Context, session sqlx.Session) error {

		// 更新角色数据
		err = l.svcCtx.SysRoleModel.TransUpdate(ctx, session, res)
		if err != nil {
			return err
		}

		// 删除角色和菜单中间表内的关系 然后重写添加
		err = l.svcCtx.SysRoleMenuModel.TransDeleteByRoleId(ctx, session, in.Id)
		if err != nil {
			return err
		}

		// 菜单IDS和角色 添加到  中间表去确定关系
		for _, menuId := range in.MenuIds {
			res, err := l.svcCtx.SysMenuModel.FindOne(ctx, menuId)
			if err != nil {
				if errors.Is(err, sqlc.ErrNotFound) {
					return fmt.Errorf("SysMenu没有该ID:%v", menuId)
				}
				return err
			}

			// 判断该数据是否被删除
			if res.DeletedAt.Valid == true {
				return fmt.Errorf("SysMenu该ID已被删除：%v", menuId)
			}

			// 加菜单和角色ID 添加到中间表去
			_, err = l.svcCtx.SysRoleMenuModel.TransInsert(l.ctx, session, &model.SysRoleMenu{
				RoleId:      in.Id,
				MenuId:      menuId,
				CreatedName: in.UpdatedName,
				CreatedAt:   time.Now(),
			})
			if err != nil {
				return err
			}

		}

		// 删除角色和接口中间表内的关系 然后重写添加
		err = l.svcCtx.SysRoleInterfaceModel.TransDeleteByRoleId(ctx, session, in.Id)
		if err != nil {
			return err
		}

		// 接口IDS和角色 添加到  中间表去确定关系
		for _, interfaceId := range in.InterfaceIds {
			res, err := l.svcCtx.SysInterfaceModel.FindOne(l.ctx, interfaceId)
			if err != nil {
				if errors.Is(err, sqlc.ErrNotFound) {
					return fmt.Errorf("SysInterfaceId没有该ID:%v", interfaceId)
				}
				return err
			}

			// 判断该数据是否被删除
			if res.DeletedAt.Valid == true {
				return fmt.Errorf("SysInterfaceId该ID已被删除：%v", interfaceId)
			}

			// 加菜单和角色ID 添加到中间表去
			_, err = l.svcCtx.SysRoleInterfaceModel.TransInsert(l.ctx, session, &model.SysRoleInterface{
				RoleId:      in.Id,
				InterfaceId: interfaceId,
				CreatedName: in.UpdatedName,
				CreatedAt:   time.Now(),
			})
			if err != nil {
				return err
			}

		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return &tpmtclient.CommonResp{}, nil
}
