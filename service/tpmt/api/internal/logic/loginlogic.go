package logic

import (
	"context"
	"time"
	"tpmt-zt/common"
	"tpmt-zt/common/jwtx"
	"tpmt-zt/common/msg"
	"tpmt-zt/service/authentication/authenticationclient"

	"tpmt-zt/service/tpmt/api/internal/svc"
	"tpmt-zt/service/tpmt/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.SysLoginRequest) (resp *types.Response, err error) {

	user, err := l.svcCtx.AuthenticationRpc.SysLogin(l.ctx, &authenticationclient.SysLoginReq{
		Account:  req.Account,
		Password: req.Password,
	})

	if err != nil {
		return nil, common.NewDefaultError(err.Error())
	}

	// 生成对应的token
	timeUnix := time.Now().Unix()

	token, err := jwtx.GetToken(l.svcCtx.Config.Auth.AccessSecret, timeUnix, l.svcCtx.Config.Auth.AccessExpire, user.Id,
		common.UserTokenType, user.NickName)

	if err != nil {
		return nil, common.NewCodeError(common.TokenNewErrorCode, "token生成失败", nil)
	}

	var respx LoginResponse

	respx.AccessToken = token
	respx.AccessExpire = timeUnix + l.svcCtx.Config.Auth.AccessExpire

	return &types.Response{
		Code: 0,
		Msg:  msg.Success,
		Data: respx,
	}, nil
}

type LoginResponse struct {
	AccessToken  string `json:"accessToken"`
	AccessExpire int64  `json:"accessExpire"`
}
