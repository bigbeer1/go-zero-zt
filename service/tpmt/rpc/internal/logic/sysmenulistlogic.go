package logic

import (
	"context"
	"github.com/Masterminds/squirrel"

	"tpmt-zt/service/tpmt/rpc/internal/svc"
	"tpmt-zt/service/tpmt/rpc/tpmtclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type SysMenuListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSysMenuListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SysMenuListLogic {
	return &SysMenuListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SysMenuListLogic) SysMenuList(in *tpmtclient.SysMenuListReq) (*tpmtclient.SysMenuListResp, error) {

	whereBuilder := l.svcCtx.SysMenuModel.RowBuilder()

	whereBuilder = whereBuilder.Where("deleted_at is null")
	whereBuilder = whereBuilder.OrderBy("created_at DESC, id DESC")

	// 菜单类型(层级关系)
	if in.MenuType != 99 {
		whereBuilder = whereBuilder.Where(squirrel.Eq{
			"menu_type ": in.MenuType,
		})
	}
	// 菜单名称
	if len(in.Name) > 0 {
		whereBuilder = whereBuilder.Where(squirrel.Like{
			"name ": "%" + in.Name + "%",
		})
	}
	// 标题
	if len(in.Title) > 0 {
		whereBuilder = whereBuilder.Where(squirrel.Like{
			"title ": "%" + in.Title + "%",
		})
	}
	// 路径
	if len(in.Path) > 0 {
		whereBuilder = whereBuilder.Where(squirrel.Like{
			"path ": "%" + in.Path + "%",
		})
	}
	// 本地路径
	if len(in.Component) > 0 {
		whereBuilder = whereBuilder.Where(squirrel.Like{
			"component ": "%" + in.Component + "%",
		})
	}
	// 跳转
	if len(in.Redirect) > 0 {
		whereBuilder = whereBuilder.Where(squirrel.Like{
			"redirect ": "%" + in.Redirect + "%",
		})
	}
	// 图标
	if len(in.Icon) > 0 {
		whereBuilder = whereBuilder.Where(squirrel.Like{
			"icon ": "%" + in.Icon + "%",
		})
	}
	// 是否隐藏
	if in.IsHide != 99 {
		whereBuilder = whereBuilder.Where(squirrel.Eq{
			"is_hide ": in.IsHide,
		})
	}
	// 是否缓存
	if in.IsKeepAlive != 99 {
		whereBuilder = whereBuilder.Where(squirrel.Eq{
			"is_keep_alive ": in.IsKeepAlive,
		})
	}
	// 父ID
	if in.ParentId != 99 {
		whereBuilder = whereBuilder.Where(squirrel.Eq{
			"parent_id ": in.ParentId,
		})
	}
	// 是否首页
	if in.IsHome != 99 {
		whereBuilder = whereBuilder.Where(squirrel.Eq{
			"is_home ": in.IsHome,
		})
	}
	// 是否主菜单
	if in.IsMain != 99 {
		whereBuilder = whereBuilder.Where(squirrel.Eq{
			"is_main ": in.IsMain,
		})
	}

	all, err := l.svcCtx.SysMenuModel.FindList(l.ctx, whereBuilder, in.Current, in.PageSize)
	if err != nil {
		return nil, err
	}

	countBuilder := l.svcCtx.SysMenuModel.CountBuilder("id")

	countBuilder = countBuilder.Where("deleted_at is null")

	// 菜单类型(层级关系)
	if in.MenuType != 99 {
		countBuilder = countBuilder.Where(squirrel.Eq{
			"menu_type ": in.MenuType,
		})
	}
	// 菜单名称
	if len(in.Name) > 0 {
		countBuilder = countBuilder.Where(squirrel.Like{
			"name ": "%" + in.Name + "%",
		})
	}
	// 标题
	if len(in.Title) > 0 {
		countBuilder = countBuilder.Where(squirrel.Like{
			"title ": "%" + in.Title + "%",
		})
	}
	// 路径
	if len(in.Path) > 0 {
		countBuilder = countBuilder.Where(squirrel.Like{
			"path ": "%" + in.Path + "%",
		})
	}
	// 本地路径
	if len(in.Component) > 0 {
		countBuilder = countBuilder.Where(squirrel.Like{
			"component ": "%" + in.Component + "%",
		})
	}
	// 跳转
	if len(in.Redirect) > 0 {
		countBuilder = countBuilder.Where(squirrel.Like{
			"redirect ": "%" + in.Redirect + "%",
		})
	}
	// 图标
	if len(in.Icon) > 0 {
		countBuilder = countBuilder.Where(squirrel.Like{
			"icon ": "%" + in.Icon + "%",
		})
	}
	// 是否隐藏
	if in.IsHide != 99 {
		countBuilder = countBuilder.Where(squirrel.Eq{
			"is_hide ": in.IsHide,
		})
	}
	// 是否缓存
	if in.IsKeepAlive != 99 {
		countBuilder = countBuilder.Where(squirrel.Eq{
			"is_keep_alive ": in.IsKeepAlive,
		})
	}
	// 父ID
	if in.ParentId != 99 {
		countBuilder = countBuilder.Where(squirrel.Eq{
			"parent_id ": in.ParentId,
		})
	}
	// 是否首页
	if in.IsHome != 99 {
		countBuilder = countBuilder.Where(squirrel.Eq{
			"is_home ": in.IsHome,
		})
	}
	// 是否主菜单
	if in.IsMain != 99 {
		countBuilder = countBuilder.Where(squirrel.Eq{
			"is_main ": in.IsMain,
		})
	}
	count, err := l.svcCtx.SysMenuModel.FindCount(l.ctx, countBuilder)
	if err != nil {
		return nil, err
	}

	var list []*tpmtclient.SysMenuListData
	for _, item := range all {
		list = append(list, &tpmtclient.SysMenuListData{
			Id:          item.Id,                         //菜单ID
			MenuType:    item.MenuType,                   //菜单类型(层级关系)
			Name:        item.Name,                       //菜单名称
			Title:       item.Title,                      //标题
			Path:        item.Path,                       //路径
			Component:   item.Component,                  //本地路径
			Redirect:    item.Redirect.String,            //跳转
			Sort:        item.Sort,                       //sort
			Icon:        item.Icon.String,                //图标
			IsHide:      item.IsHide,                     //是否隐藏
			IsKeepAlive: item.IsKeepAlive,                //是否缓存
			ParentId:    item.ParentId,                   //父ID
			IsHome:      item.IsHome,                     //是否首页
			IsMain:      item.IsMain,                     //是否主菜单
			CreatedName: item.CreatedName,                //创建人
			CreatedAt:   item.CreatedAt.UnixMilli(),      //创建时间
			UpdatedName: item.UpdatedName.String,         //更新人
			UpdatedAt:   item.UpdatedAt.Time.UnixMilli(), //更新时间
		})
	}

	return &tpmtclient.SysMenuListResp{
		Total: count,
		List:  list,
	}, nil
}
