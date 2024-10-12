package sysAuth

import (
	"context"
	"github.com/jinzhu/copier"
	"tpmt-zt/common"
	"tpmt-zt/common/msg"
	"tpmt-zt/service/tpmt/rpc/tpmtclient"

	"tpmt-zt/service/tpmt/api/internal/svc"
	"tpmt-zt/service/tpmt/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SysAuthInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSysAuthInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SysAuthInfoLogic {
	return &SysAuthInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SysAuthInfoLogic) SysAuthInfo(req *types.SysAuthInfoRequest) (resp *types.Response, err error) {

	res, err := l.svcCtx.TpmtRpc.SysAuthFindOne(l.ctx, &tpmtclient.SysAuthFindOneReq{
		Id: req.Id, // 第三方用户ID
	})
	if err != nil {
		return nil, common.NewDefaultError(err.Error())
	}

	var result SysAuthFindOneResp
	_ = copier.Copy(&result, res)

	return &types.Response{
		Code: 0,
		Msg:  msg.Success,
		Data: result,
	}, nil
}

type SysAuthFindOneResp struct {
	Id          string `json:"id"`           // 第三方用户ID,
	CreatedAt   int64  `json:"created_at"`   // 创建时间,
	UpdatedAt   int64  `json:"updated_at"`   // 更新时间,
	CreatedName string `json:"created_name"` // 创建人,
	UpdatedName string `json:"updated_name"` // 更新人,
	NickName    string `json:"nick_name"`    // 机构名,
	AuthToken   string `json:"auth_token"`   // 令牌,
	State       int64  `json:"state"`        // 状态 1:正常 2:停用 3:封禁
}
