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
	"tpmt-zt/common"
	"tpmt-zt/common/jwtx"
	"tpmt-zt/service/tpmt/model"

	"tpmt-zt/service/tpmt/rpc/internal/svc"
	"tpmt-zt/service/tpmt/rpc/tpmtclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type SysAuthAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSysAuthAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SysAuthAddLogic {
	return &SysAuthAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 第三方用户
func (l *SysAuthAddLogic) SysAuthAdd(in *tpmtclient.SysAuthAddReq) (*tpmtclient.CommonResp, error) {

	// 生成对应的token
	timeUnix := time.Now().Unix()
	id := uuid.NewV4().String()

	token, err := jwtx.GetToken(l.svcCtx.Config.CAuth.AccessSecret, timeUnix, l.svcCtx.Config.CAuth.AccessExpire*90000, id,
		common.AuthTokenType, in.NickName)

	if err != nil {
		return nil, fmt.Errorf("token生成失败")
	}

	// 开启事务

	err = l.svcCtx.SysAuthModel.TransCtx(l.ctx, func(ctx context.Context, session sqlx.Session) error {

		// 添加基础信息
		_, err = l.svcCtx.SysAuthModel.TransInsert(l.ctx, session, &model.SysAuth{
			Id:          id,             // ID
			CreatedAt:   time.Now(),     // 创建时间
			CreatedName: in.CreatedName, // 创建人
			NickName:    in.NickName,    // 机构名
			AuthToken:   token,          // 令牌
			State:       in.State,       // 状态 1:正常 2:停用 3:封禁
		})

		if in.RoleId != 0 {
			role, err := l.svcCtx.SysRoleModel.FindOne(l.ctx, in.RoleId)
			if err != nil {
				if errors.Is(err, sqlc.ErrNotFound) {
					return fmt.Errorf("SysRole没有该ID:%v", in.RoleId)
				}
				return err
			}

			// 判断角色类型是否是普通角色
			if role.RoleType != 3 {
				return fmt.Errorf("该角色类型无法添加到第三方用户")
			}

			// 判断该数据是否被删除
			if role.DeletedAt.Valid == true {
				return fmt.Errorf("SysRole该ID已被删除：%v", role.Id)
			}

			// 中间表添加角色和用户的关系
			_, err = l.svcCtx.SysUserRoleModel.TransInsert(l.ctx, session, &model.SysUserRole{
				UserId: id,
				RoleId: role.Id,
				UserType: sql.NullInt64{
					Int64: 2,
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
