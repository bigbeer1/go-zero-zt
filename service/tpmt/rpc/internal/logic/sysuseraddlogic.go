package logic

import (
	"context"
	"fmt"
	uuid "github.com/satori/go.uuid"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
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

	_, err := l.svcCtx.SysUserModel.FindByAccount(l.ctx, in.Account)
	if err != nil {
		if err != sqlc.ErrNotFound {
			return nil, err
		}
	}

	// 已经存在了Account 我们需要返回错误
	if err == nil {
		return nil, fmt.Errorf("用户名:%v已存在", in.Account)
	}

	// sha256 加密
	password := cryptx.PasswordEncrypt(l.svcCtx.Config.Salt+in.Account, in.Password)

	_, err = l.svcCtx.SysUserModel.Insert(l.ctx, &model.SysUser{
		Id:          uuid.NewV4().String(), // ID
		CreatedAt:   time.Now(),            // 创建时间
		Account:     in.Account,            // 用户名
		NickName:    in.NickName,           // 姓名
		Password:    password,              // 密码
		State:       in.State,              // 状态 1:正常 2:停用 3:封禁
		CreatedName: in.CreatedName,        // 创建人
	})
	if err != nil {
		return nil, err
	}

	return &tpmtclient.CommonResp{}, nil
}
