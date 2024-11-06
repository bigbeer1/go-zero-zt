package sysMenu

import (
	"context"
	"github.com/jinzhu/copier"
	"tpmt-zt/common"
	"tpmt-zt/common/msg"
	"tpmt-zt/service/authentication/authenticationclient"

	"tpmt-zt/service/tpmt/api/internal/svc"
	"tpmt-zt/service/tpmt/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SysMenuListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSysMenuListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SysMenuListLogic {
	return &SysMenuListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SysMenuListLogic) SysMenuList(req *types.SysMenuListRequest) (resp *types.Response, err error) {

	all, err := l.svcCtx.AuthenticationRpc.SysMenuList(l.ctx, &authenticationclient.SysMenuListReq{
		Current:     req.Current,     // 页码
		PageSize:    req.PageSize,    // 页数
		MenuType:    req.MenuType,    // 菜单类型(层级关系)
		Name:        req.Name,        // 菜单名称
		Title:       req.Title,       // 标题
		Path:        req.Path,        // 路径
		Component:   req.Component,   // 本地路径
		Redirect:    req.Redirect,    // 跳转
		Icon:        req.Icon,        // 图标
		IsHide:      req.IsHide,      // 是否隐藏
		IsKeepAlive: req.IsKeepAlive, // 是否缓存
		ParentId:    req.ParentId,    // 父ID
		IsHome:      req.IsHome,      // 是否首页
		IsMain:      req.IsMain,      // 是否主菜单
	})
	if err != nil {
		return nil, common.NewDefaultError(err.Error())
	}

	var result SysMenuListResp
	_ = copier.Copy(&result, all)

	return &types.Response{
		Code: 0,
		Msg:  msg.Success,
		Data: result,
	}, nil
}

type SysMenuListResp struct {
	Total int64              `json:"total"`
	List  []*SysMenuDataList `json:"list"`
}

type SysMenuDataList struct {
	Id          int64  `json:"id"`            // 菜单ID,
	MenuType    int64  `json:"menu_type"`     // 菜单类型(层级关系),
	Name        string `json:"name"`          // 菜单名称,
	Title       string `json:"title"`         // 标题,
	Path        string `json:"path"`          // 路径,
	Component   string `json:"component"`     // 本地路径,
	Redirect    string `json:"redirect"`      // 跳转,
	Icon        string `json:"icon"`          // 图标,
	IsHide      int64  `json:"is_hide"`       // 是否隐藏,
	IsKeepAlive int64  `json:"is_keep_alive"` // 是否缓存,
	ParentId    int64  `json:"parent_id"`     // 父ID,
	IsHome      int64  `json:"is_home"`       // 是否首页,
	IsMain      int64  `json:"is_main"`       // 是否主菜单,
	CreatedName string `json:"created_name"`  // 创建人,
	CreatedAt   int64  `json:"created_at"`    // 创建时间,
	UpdatedName string `json:"updated_name"`  // 更新人,
	UpdatedAt   int64  `json:"updated_at"`    // 更新时间
}
