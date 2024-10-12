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

type SysAuthListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSysAuthListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SysAuthListLogic {
	return &SysAuthListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SysAuthListLogic) SysAuthList(req *types.SysAuthListRequest) (resp *types.Response, err error) {

	all, err := l.svcCtx.TpmtRpc.SysAuthList(l.ctx, &tpmtclient.SysAuthListReq{
		Current:   req.Current,   // 页码
		PageSize:  req.PageSize,  // 页数
		NickName:  req.NickName,  // 机构名
		AuthToken: req.AuthToken, // 令牌
		State:     req.State,     // 状态 1:正常 2:停用 3:封禁
	})
	if err != nil {
		return nil, common.NewDefaultError(err.Error())
	}

	var result SysAuthListResp
	_ = copier.Copy(&result, all)

	return &types.Response{
		Code: 0,
		Msg:  msg.Success,
		Data: result,
	}, nil
}

type SysAuthListResp struct {
	Total int64              `json:"total"`
	List  []*SysAuthDataList `json:"list"`
}

type SysAuthDataList struct {
	Id          string `json:"id"`           // 第三方用户ID,
	CreatedAt   int64  `json:"created_at"`   // 创建时间,
	UpdatedAt   int64  `json:"updated_at"`   // 更新时间,
	CreatedName string `json:"created_name"` // 创建人,
	UpdatedName string `json:"updated_name"` // 更新人,
	NickName    string `json:"nick_name"`    // 机构名,
	AuthToken   string `json:"auth_token"`   // 令牌,
	State       int64  `json:"state"`        // 状态 1:正常 2:停用 3:封禁
}
