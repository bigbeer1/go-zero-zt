package logic

import (
	"context"
	"github.com/Masterminds/squirrel"
	"tpmt-zt/service/authentication/authenticationclient"
	"tpmt-zt/service/authentication/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type SysMenuByRoleIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSysMenuByRoleIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SysMenuByRoleIdLogic {
	return &SysMenuByRoleIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 通过角色ID获取菜单信息
func (l *SysMenuByRoleIdLogic) SysMenuByRoleId(in *authenticationclient.SysMenuByRoleIdReq) (*authenticationclient.SysMenuByRoleIdResp, error) {
	// 先去中间表找对应的接口IDS
	whereBuilder := l.svcCtx.SysRoleMenuModel.RowBuilder()
	whereBuilder = whereBuilder.OrderBy("created_at DESC, id DESC")
	// 接口名称
	whereBuilder = whereBuilder.Where(squirrel.Eq{
		"role_id ": in.RoleId,
	})

	all, err := l.svcCtx.SysRoleMenuModel.FindAll(l.ctx, whereBuilder)
	if err != nil {
		return nil, err
	}

	var list []*authenticationclient.SysMenuListData
	for _, v := range all {
		item, err := l.svcCtx.SysMenuModel.FindOne(l.ctx, v.MenuId)
		if err != nil {
			return nil, err
		}
		list = append(list, &authenticationclient.SysMenuListData{
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
	return &authenticationclient.SysMenuByRoleIdResp{
		List: list,
	}, nil
}
