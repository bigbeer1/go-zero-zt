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

type SysRoleAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSysRoleAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SysRoleAddLogic {
	return &SysRoleAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 角色
func (l *SysRoleAddLogic) SysRoleAdd(in *tpmtclient.SysRoleAddReq) (*tpmtclient.CommonResp, error) {

	err := l.svcCtx.SysRoleModel.TransCtx(l.ctx, func(ctx context.Context, session sqlx.Session) error {

		return nil
	})

	role, err := l.svcCtx.SysRoleModel.Insert(l.ctx, &model.SysRole{
		CreatedAt:   time.Now(),                                                // 创建时间
		Name:        in.Name,                                                   // 角色名称
		Remark:      sql.NullString{String: in.Remark, Valid: in.Remark != ""}, // 备注
		RoleType:    in.RoleType,                                               // 角色类型 1:管理员角色  2:普通角色  3:第三方角色
		CreatedName: in.CreatedName,                                            // 创建人
	})

	if err != nil {
		return nil, err
	}

	// 获取角色ID
	roleId, err := role.LastInsertId()
	if err != nil {
		return nil, err
	}

	// 菜单IDS和角色 添加到  中间表去确定关系
	for _, menuId := range in.MenuIds {
		res, err := l.svcCtx.SysMenuModel.FindOne(l.ctx, menuId)
		if err != nil {
			if errors.Is(err, sqlc.ErrNotFound) {
				return nil, fmt.Errorf("SysMenu没有该ID:%v", menuId)
			}
			return nil, err
		}

		// 判断该数据是否被删除
		if res.DeletedAt.Valid == true {
			return nil, fmt.Errorf("SysMenu该ID已被删除：%v", menuId)
		}

		// 加菜单和角色ID 添加到中间表去
		l.svcCtx.SysRoleMenuModel.Insert(l.ctx, &model.SysRoleMenu{
			RoleId:      roleId,
			MenuId:      menuId,
			CreatedName: in.CreatedName,
			CreatedAt:   time.Now(),
		})

	}

	// 接口IDS和角色 添加到  中间表去确定关系

	if err != nil {
		return nil, err
	}

	return &tpmtclient.CommonResp{}, nil
}
