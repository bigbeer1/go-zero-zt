package logic

import (
	"context"
	"github.com/Masterminds/squirrel"
	"tpmt-zt/service/tpmt/rpc/internal/svc"
	"tpmt-zt/service/tpmt/rpc/tpmtclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type SysInterfaceByRoleIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSysInterfaceByRoleIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SysInterfaceByRoleIdLogic {
	return &SysInterfaceByRoleIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 通过角色ID获取接口信息
func (l *SysInterfaceByRoleIdLogic) SysInterfaceByRoleId(in *tpmtclient.SysInterfaceByRoleIdReq) (*tpmtclient.SysInterfaceByRoleIdResp, error) {
	// 先去中间表找对应的接口IDS
	whereBuilder := l.svcCtx.SysRoleInterfaceModel.RowBuilder()
	whereBuilder = whereBuilder.OrderBy("created_at DESC, id DESC")
	// 接口名称
	whereBuilder = whereBuilder.Where(squirrel.Eq{
		"role_id ": in.RoleId,
	})

	all, err := l.svcCtx.SysRoleInterfaceModel.FindAll(l.ctx, whereBuilder)
	if err != nil {
		return nil, err
	}

	var list []*tpmtclient.SysInterfaceListData
	for _, v := range all {
		item, err := l.svcCtx.SysInterfaceModel.FindOne(l.ctx, v.InterfaceId)
		if err != nil {
			return nil, err
		}
		list = append(list, &tpmtclient.SysInterfaceListData{
			Id:                 item.Id,                         //接口ID
			CreatedAt:          item.CreatedAt.UnixMilli(),      //创建时间
			UpdatedAt:          item.UpdatedAt.Time.UnixMilli(), //更新时间
			CreatedName:        item.CreatedName,                //创建人
			UpdatedName:        item.UpdatedName.String,         //更新人
			Name:               item.Name,                       //接口名称
			Path:               item.Path,                       //接口地址
			InterfaceType:      item.InterfaceType,              //接口类型
			InterfaceGroupName: item.InterfaceGroupName.String,  //接口分组名称
			Remark:             item.Remark.String,              //备注
			Sort:               item.Sort,                       //sort
		})
	}

	return &tpmtclient.SysInterfaceByRoleIdResp{
		List: list,
	}, nil
}
