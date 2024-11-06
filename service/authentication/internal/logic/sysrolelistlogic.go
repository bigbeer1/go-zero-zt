package logic

import (
	"context"
	"github.com/Masterminds/squirrel"
	"tpmt-zt/service/authentication/authenticationclient"
	"tpmt-zt/service/authentication/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type SysRoleListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSysRoleListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SysRoleListLogic {
	return &SysRoleListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SysRoleListLogic) SysRoleList(in *authenticationclient.SysRoleListReq) (*authenticationclient.SysRoleListResp, error) {
	whereBuilder := l.svcCtx.SysRoleModel.RowBuilder()

	whereBuilder = whereBuilder.Where("deleted_at is null")
	whereBuilder = whereBuilder.OrderBy("created_at DESC, id DESC")

	// 角色名称
	if len(in.Name) > 0 {
		whereBuilder = whereBuilder.Where(squirrel.Like{
			"name ": "%" + in.Name + "%",
		})
	}
	// 备注
	if len(in.Remark) > 0 {
		whereBuilder = whereBuilder.Where(squirrel.Like{
			"remark ": "%" + in.Remark + "%",
		})
	}
	// 角色类型 1:管理员角色  2:普通角色  3:第三方角色
	if in.RoleType != 99 {
		whereBuilder = whereBuilder.Where(squirrel.Eq{
			"role_type ": in.RoleType,
		})
	}

	all, err := l.svcCtx.SysRoleModel.FindList(l.ctx, whereBuilder, in.Current, in.PageSize)
	if err != nil {
		return nil, err
	}

	countBuilder := l.svcCtx.SysRoleModel.CountBuilder("id")

	countBuilder = countBuilder.Where("deleted_at is null")

	// 角色名称
	if len(in.Name) > 0 {
		countBuilder = countBuilder.Where(squirrel.Like{
			"name ": "%" + in.Name + "%",
		})
	}
	// 备注
	if len(in.Remark) > 0 {
		countBuilder = countBuilder.Where(squirrel.Like{
			"remark ": "%" + in.Remark + "%",
		})
	}
	// 角色类型 1:管理员角色  2:普通角色  3:第三方角色
	if in.RoleType != 99 {
		countBuilder = countBuilder.Where(squirrel.Eq{
			"role_type ": in.RoleType,
		})
	}
	count, err := l.svcCtx.SysRoleModel.FindCount(l.ctx, countBuilder)
	if err != nil {
		return nil, err
	}

	var list []*authenticationclient.SysRoleListData
	for _, item := range all {
		list = append(list, &authenticationclient.SysRoleListData{
			Id:          item.Id,                         //角色ID
			Name:        item.Name,                       //角色名称
			Remark:      item.Remark.String,              //备注
			RoleType:    item.RoleType,                   //角色类型 1:管理员角色  2:普通角色  3:第三方角色
			CreatedName: item.CreatedName,                //创建人
			CreatedAt:   item.CreatedAt.UnixMilli(),      //创建时间
			UpdatedName: item.UpdatedName.String,         //更新人
			UpdatedAt:   item.UpdatedAt.Time.UnixMilli(), //更新时间
		})
	}

	return &authenticationclient.SysRoleListResp{
		Total: count,
		List:  list,
	}, nil
}
