package sysDict

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

type SysDictUpLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSysDictUpLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SysDictUpLogic {
	return &SysDictUpLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SysDictUpLogic) SysDictUp(req *types.SysDictUpRequest) (resp *types.Response, err error) {
	// 用户登录信息
	tokenData := jwtx.ParseToken(l.ctx)

	_, err = l.svcCtx.AuthenticationRpc.SysDictUpdate(l.ctx, &authenticationclient.SysDictUpdateReq{
		Id:          req.Id,             // 字典类型ID
		UpdatedName: tokenData.NickName, // 更新人
		DictType:    req.DictType,       // 字典类型
		DictLabel:   req.DictLabel,      // 字典标签
		DictValue:   req.DictValue,      // 字典键值
		Sort:        req.Sort,           // 排序
		Remark:      req.Remark,         // 备注
		State:       req.State,          // 状态
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
