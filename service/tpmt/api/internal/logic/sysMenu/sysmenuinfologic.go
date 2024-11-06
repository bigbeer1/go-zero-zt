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

type SysMenuInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSysMenuInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SysMenuInfoLogic {
	return &SysMenuInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SysMenuInfoLogic) SysMenuInfo(req *types.SysMenuInfoRequest) (resp *types.Response, err error) {

	res, err := l.svcCtx.AuthenticationRpc.SysMenuFindOne(l.ctx, &authenticationclient.SysMenuFindOneReq{
		Id: req.Id, // 菜单ID
	})
	if err != nil {
		return nil, common.NewDefaultError(err.Error())
	}

	var result SysMenuFindOneResp
	_ = copier.Copy(&result, res)

	return &types.Response{
		Code: 0,
		Msg:  msg.Success,
		Data: result,
	}, nil
}

type SysMenuFindOneResp struct {
	Id          int64  `json:"id"`            // 菜单ID,
	MenuType    int64  `json:"menu_type"`     // 菜单类型(层级关系), 1:目录 2:菜单 3:按钮
	Name        string `json:"name"`          // 菜单名称,
	Title       string `json:"title"`         // 标题,
	Path        string `json:"path"`          // 路径,
	Component   string `json:"component"`     // 本地路径,
	Redirect    string `json:"redirect"`      // 跳转,
	Sort        int64  `json:"sort"`          // sort,
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
