package logic

import (
	"context"
	"errors"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"time"
	"tpmt-zt/service/authentication/authenticationclient"
	"tpmt-zt/service/authentication/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type SysMenuUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSysMenuUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SysMenuUpdateLogic {
	return &SysMenuUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SysMenuUpdateLogic) SysMenuUpdate(in *authenticationclient.SysMenuUpdateReq) (*authenticationclient.CommonResp, error) {

	res, err := l.svcCtx.SysMenuModel.FindOne(l.ctx, in.Id)
	if err != nil {
		if errors.Is(err, sqlc.ErrNotFound) {
			return nil, fmt.Errorf("SysMenu没有该ID: %v", in.Id)
		}
		return nil, err
	}

	// 判断该数据是否被删除
	if res.DeletedAt.Valid == true {
		return nil, fmt.Errorf("SysMenu该ID已被删除：%v", in.Id)
	}

	// 菜单类型(层级关系)
	if in.MenuType != 0 {
		res.MenuType = in.MenuType
	}
	// 菜单名称
	if len(in.Name) > 0 {
		res.Name = in.Name
	}
	// 标题
	if len(in.Title) > 0 {
		res.Title = in.Title
	}
	// 路径
	if len(in.Path) > 0 {
		res.Path = in.Path
	}
	// 本地路径
	if len(in.Component) > 0 {
		res.Component = in.Component
	}
	// 跳转
	if len(in.Redirect) > 0 {
		res.Redirect.String = in.Redirect
		res.Redirect.Valid = true
	}
	// sort
	if in.Sort != 0 {
		res.Sort = in.Sort
	}
	// 图标
	if len(in.Icon) > 0 {
		res.Icon.String = in.Icon
		res.Icon.Valid = true
	}
	// 是否隐藏
	if in.IsHide != 0 {
		res.IsHide = in.IsHide
	}
	// 是否缓存
	if in.IsKeepAlive != 0 {
		res.IsKeepAlive = in.IsKeepAlive
	}
	// 父ID
	if in.ParentId != 0 {
		res.ParentId = in.ParentId
	}
	// 是否首页
	if in.IsHome != 0 {
		res.IsHome = in.IsHome
	}
	// 是否主菜单
	if in.IsMain != 0 {
		res.IsMain = in.IsMain
	}

	res.UpdatedName.String = in.UpdatedName
	res.UpdatedName.Valid = true
	res.UpdatedAt.Time = time.Now()
	res.UpdatedAt.Valid = true

	err = l.svcCtx.SysMenuModel.Update(l.ctx, res)

	if err != nil {
		return nil, err
	}
	return &authenticationclient.CommonResp{}, nil
}
