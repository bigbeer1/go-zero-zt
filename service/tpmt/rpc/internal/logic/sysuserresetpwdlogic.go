package logic

import (
	"context"
	"errors"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"math/rand"
	"time"
	"tpmt-zt/common/cryptx"

	"tpmt-zt/service/tpmt/rpc/internal/svc"
	"tpmt-zt/service/tpmt/rpc/tpmtclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type SysUserResetPwdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSysUserResetPwdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SysUserResetPwdLogic {
	return &SysUserResetPwdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 重置用户密码
func (l *SysUserResetPwdLogic) SysUserResetPwd(in *tpmtclient.SysUserResetPwdReq) (*tpmtclient.SysUserResetPwdResp, error) {
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

	// 传了密码
	if len(in.Password) > 0 {
		// sha256 加密
		password := cryptx.PasswordEncrypt(l.svcCtx.Config.Salt+res.Account, in.Password)
		res.Password = password
	} else {
		in.Password = fmt.Sprintf("%v", rand.Int31n(99999999))
		// sha256 加密
		password := cryptx.PasswordEncrypt(l.svcCtx.Config.Salt+res.Account, in.Password)
		res.Password = password
	}

	res.UpdatedName.String = in.UpdatedName
	res.UpdatedName.Valid = true
	res.UpdatedAt.Time = time.Now()
	res.UpdatedAt.Valid = true

	err = l.svcCtx.SysUserModel.Update(l.ctx, res)

	if err != nil {
		return nil, err
	}
	return &tpmtclient.SysUserResetPwdResp{
		Password: in.Password,
	}, nil

}
