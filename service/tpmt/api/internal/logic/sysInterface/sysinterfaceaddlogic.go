package sysInterface

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

type SysInterfaceAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSysInterfaceAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SysInterfaceAddLogic {
	return &SysInterfaceAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SysInterfaceAddLogic) SysInterfaceAdd(req *types.SysInterfaceAddRequest) (resp *types.Response, err error) {
	// 用户登录信息
	tokenData := jwtx.ParseToken(l.ctx)

	_, err = l.svcCtx.AuthenticationRpc.SysInterfaceAdd(l.ctx, &authenticationclient.SysInterfaceAddReq{
		CreatedName:        tokenData.NickName,     // 创建人
		Name:               req.Name,               // 接口名称
		Path:               req.Path,               // 接口地址
		InterfaceType:      req.InterfaceType,      // 接口类型
		InterfaceGroupName: req.InterfaceGroupName, // 接口分组名称
		Remark:             req.Remark,             // 备注
		Sort:               req.Sort,               // sort
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
