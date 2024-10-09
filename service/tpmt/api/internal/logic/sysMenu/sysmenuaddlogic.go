package sysMenu

import (
	"context"
	"tpmt-zt/common"
	"tpmt-zt/common/jwtx"
	"tpmt-zt/common/msg"
	"tpmt-zt/service/tpmt/rpc/tpmtclient"

	"tpmt-zt/service/tpmt/api/internal/svc"
	"tpmt-zt/service/tpmt/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SysMenuAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSysMenuAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SysMenuAddLogic {
	return &SysMenuAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SysMenuAddLogic) SysMenuAdd(req *types.SysMenuAddRequest) (resp *types.Response, err error) {
	// 用户登录信息
	tokenData := jwtx.ParseToken(l.ctx)

	_, err = l.svcCtx.TpmtRpc.SysMenuAdd(l.ctx, &tpmtclient.SysMenuAddReq{
		MenuType:    req.MenuType,       // 菜单类型(层级关系)
		Name:        req.Name,           // 菜单名称
		Title:       req.Title,          // 标题
		Path:        req.Path,           // 路径
		Component:   req.Component,      // 本地路径
		Redirect:    req.Redirect,       // 跳转
		Sort:        req.Sort,           // sort
		Icon:        req.Icon,           // 图标
		IsHide:      req.IsHide,         // 是否隐藏
		IsKeepAlive: req.IsKeepAlive,    // 是否缓存
		ParentId:    req.ParentId,       // 父ID
		IsHome:      req.IsHome,         // 是否首页
		IsMain:      req.IsMain,         // 是否主菜单
		CreatedName: tokenData.NickName, // 创建人
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
