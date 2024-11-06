package logic

import (
	"context"
	"errors"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"time"
	"tpmt-zt/common/cryptx"
	"tpmt-zt/service/authentication/authenticationclient"
	"tpmt-zt/service/authentication/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type SysUserUpMyPwdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSysUserUpMyPwdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SysUserUpMyPwdLogic {
	return &SysUserUpMyPwdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 用户修改自己的密码
func (l *SysUserUpMyPwdLogic) SysUserUpMyPwd(in *authenticationclient.SysUserUpMyPwdReq) (*authenticationclient.CommonResp, error) {
	// 查到用户信息
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

	// 判断用户的密码对不对
	// sha256 加密
	passwordOld := cryptx.PasswordEncrypt(l.svcCtx.Config.Salt+res.Account, in.OldPassword)
	if passwordOld != res.Password {
		return nil, fmt.Errorf("密码错误")
	}

	// sha256 加密
	passwordNew := cryptx.PasswordEncrypt(l.svcCtx.Config.Salt+res.Account, in.NewPassword)
	res.Password = passwordNew

	res.UpdatedName.String = res.NickName
	res.UpdatedName.Valid = true
	res.UpdatedAt.Time = time.Now()
	res.UpdatedAt.Valid = true

	err = l.svcCtx.SysUserModel.Update(l.ctx, res)

	return &authenticationclient.CommonResp{}, nil
}
