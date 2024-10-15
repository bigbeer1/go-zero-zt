package logic

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	uuid "github.com/satori/go.uuid"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"time"
	"tpmt-zt/common/cryptx"
	"tpmt-zt/service/tpmt/model"
	"tpmt-zt/service/tpmt/rpc/internal/svc"
	"tpmt-zt/service/tpmt/rpc/tpmtclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type SysUserAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSysUserAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SysUserAddLogic {
	return &SysUserAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SysUserAddLogic) SysUserAdd(in *tpmtclient.SysUserAddReq) (*tpmtclient.CommonResp, error) {

	// 查用户名是否又重复
	_, err := l.svcCtx.SysUserModel.FindByAccount(l.ctx, in.Account)
	if err != nil {
		if !errors.Is(err, sqlc.ErrNotFound) {
			return nil, err
		}
	}

	// 已经存在了Account 我们需要返回错误
	if err == nil {
		return nil, fmt.Errorf("用户名:%v已存在", in.Account)
	}

	// 开启事务
	err = l.svcCtx.SysUserModel.TransCtx(l.ctx, func(ctx context.Context, session sqlx.Session) error {

		// sha256 加密
		password := cryptx.PasswordEncrypt(l.svcCtx.Config.Salt+in.Account, in.Password)
		// 添加用户信息
		userId := uuid.NewV4().String()
		_, err = l.svcCtx.SysUserModel.TransInsert(ctx, session, &model.SysUser{
			Id:          userId,         // ID
			CreatedAt:   time.Now(),     // 创建时间
			Account:     in.Account,     // 用户名
			NickName:    in.NickName,    // 姓名
			Password:    password,       // 密码
			State:       in.State,       // 状态 1:正常 2:停用 3:封禁
			CreatedName: in.CreatedName, // 创建人
		})
		if err != nil {
			return err
		}

		// 用户添加角色   当然也不添加
		if in.RoleId != 0 {
			role, err := l.svcCtx.SysRoleModel.FindOne(l.ctx, in.RoleId)
			if err != nil {
				if errors.Is(err, sqlc.ErrNotFound) {
					return fmt.Errorf("SysRole没有该ID:%v", in.RoleId)
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

			// 中间表添加角色和用户的关系
			_, err = l.svcCtx.SysUserRoleModel.TransInsert(l.ctx, session, &model.SysUserRole{
				UserId: userId,
				RoleId: role.Id,
				UserType: sql.NullInt64{
					Int64: 1,
					Valid: true,
				},
				CreatedName: in.CreatedName,
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
