package sysDictType

import (
	"context"
	"tpmt-zt/common"
	"tpmt-zt/common/jwtx"
	"tpmt-zt/common/msg"
	"tpmt-zt/service/authentication/authenticationclient"

	"tpmt-zt/service/tpmt/api/internal/svc"
	"tpmt-zt/service/tpmt/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SysDictTypeAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSysDictTypeAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SysDictTypeAddLogic {
	return &SysDictTypeAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SysDictTypeAddLogic) SysDictTypeAdd(req *types.SysDictTypeAddRequest) (resp *types.Response, err error) {
	// 用户登录信息
	tokenData := jwtx.ParseToken(l.ctx)

	_, err = l.svcCtx.AuthenticationRpc.SysDictTypeAdd(l.ctx, &authenticationclient.SysDictTypeAddReq{
		CreatedName: tokenData.NickName, // 创建人
		Name:        req.Name,           // 字典名称
		DictType:    req.DictType,       // 字典类型
		State:       req.State,          // 状态
		Remark:      req.Remark,         // 描述
		Sort:        req.Sort,           // 排序
	})
	if err != nil {
		return nil, common.NewDefaultError(err.Error())
	}
	return &types.Response{
		Code: 0,
		Msg:  msg.Success,
		Data: nil,
	}, nil
}
