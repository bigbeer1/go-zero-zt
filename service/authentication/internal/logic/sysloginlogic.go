package logic

import (
	"context"
	"errors"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"tpmt-zt/common/cryptx"
	"tpmt-zt/service/authentication/authenticationclient"
	"tpmt-zt/service/authentication/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type SysLoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSysLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SysLoginLogic {
	return &SysLoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SysLoginLogic) SysLogin(in *authenticationclient.SysLoginReq) (*authenticationclient.SysUserFindOneResp, error) {
	// 根据用户名 去找对应的用户信息
	user, err := l.svcCtx.SysUserModel.FindByAccount(l.ctx, in.Account)
	if err != nil {
		if errors.Is(err, sqlc.ErrNotFound) {
			return nil, fmt.Errorf("用户不存在")
		}
		return nil, err
	}

	// 判断用户的密码对不对
	// sha256 加密
	password := cryptx.PasswordEncrypt(l.svcCtx.Config.Salt+in.Account, in.Password)
	if password != user.Password {
		return nil, fmt.Errorf("用户密码错误")
	}

	switch user.State {

	case 2:
		return nil, fmt.Errorf("用户已停用")
	case 3:
		return nil, fmt.Errorf("用户已封禁")

	}

	return &authenticationclient.SysUserFindOneResp{
		Id:          user.Id,
		Account:     user.Account,
		NickName:    user.NickName,
		State:       user.State,
		CreatedName: user.CreatedName,
		CreatedAt:   user.CreatedAt.UnixMilli(),
		UpdatedName: user.UpdatedName.String,
		UpdatedAt:   user.UpdatedAt.Time.UnixMilli(),
	}, nil
}
