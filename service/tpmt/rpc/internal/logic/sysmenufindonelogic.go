package logic

import (
	"context"
	"errors"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlc"

	"tpmt-zt/service/tpmt/rpc/internal/svc"
	"tpmt-zt/service/tpmt/rpc/tpmtclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type SysMenuFindOneLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSysMenuFindOneLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SysMenuFindOneLogic {
	return &SysMenuFindOneLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SysMenuFindOneLogic) SysMenuFindOne(in *tpmtclient.SysMenuFindOneReq) (*tpmtclient.SysMenuFindOneResp, error) {
	res, err := l.svcCtx.SysMenuModel.FindOne(l.ctx, in.Id)
	if err != nil {
		if errors.Is(err, sqlc.ErrNotFound) {
			return nil, fmt.Errorf("SysMenu没有该ID:%v", in.Id)
		}
		return nil, err
	}

	// 判断该数据是否被删除
	if res.DeletedAt.Valid == true {
		return nil, fmt.Errorf("SysMenu该ID已被删除：%v", in.Id)
	}

	return &tpmtclient.SysMenuFindOneResp{
		Id:          res.Id,                         //菜单ID
		MenuType:    res.MenuType,                   //菜单类型(层级关系)
		Name:        res.Name,                       //菜单名称
		Title:       res.Title,                      //标题
		Path:        res.Path,                       //路径
		Component:   res.Component,                  //本地路径
		Redirect:    res.Redirect.String,            //跳转
		Sort:        res.Sort,                       //sort
		Icon:        res.Icon.String,                //图标
		IsHide:      res.IsHide,                     //是否隐藏
		IsKeepAlive: res.IsKeepAlive,                //是否缓存
		ParentId:    res.ParentId,                   //父ID
		IsHome:      res.IsHome,                     //是否首页
		IsMain:      res.IsMain,                     //是否主菜单
		CreatedName: res.CreatedName,                //创建人
		CreatedAt:   res.CreatedAt.UnixMilli(),      //创建时间
		UpdatedName: res.UpdatedName.String,         //更新人
		UpdatedAt:   res.UpdatedAt.Time.UnixMilli(), //更新时间
	}, nil
}
