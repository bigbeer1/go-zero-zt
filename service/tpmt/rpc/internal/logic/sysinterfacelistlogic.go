package logic

import (
	"context"
	"github.com/Masterminds/squirrel"

	"tpmt-zt/service/tpmt/rpc/internal/svc"
	"tpmt-zt/service/tpmt/rpc/tpmtclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type SysInterfaceListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSysInterfaceListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SysInterfaceListLogic {
	return &SysInterfaceListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SysInterfaceListLogic) SysInterfaceList(in *tpmtclient.SysInterfaceListReq) (*tpmtclient.SysInterfaceListResp, error) {

	whereBuilder := l.svcCtx.SysInterfaceModel.RowBuilder()

	whereBuilder = whereBuilder.Where("deleted_at is null")
	whereBuilder = whereBuilder.OrderBy("created_at DESC, id DESC")

	// 接口名称
	if len(in.Name) > 0 {
		whereBuilder = whereBuilder.Where(squirrel.Like{
			"name ": "%" + in.Name + "%",
		})
	}
	// 接口地址
	if len(in.Path) > 0 {
		whereBuilder = whereBuilder.Where(squirrel.Like{
			"path ": "%" + in.Path + "%",
		})
	}
	// 接口类型
	if len(in.InterfaceType) > 0 {
		whereBuilder = whereBuilder.Where(squirrel.Like{
			"interface_type ": "%" + in.InterfaceType + "%",
		})
	}
	// 接口分组名称
	if len(in.InterfaceGroupName) > 0 {
		whereBuilder = whereBuilder.Where(squirrel.Like{
			"interface_group_name ": "%" + in.InterfaceGroupName + "%",
		})
	}
	// 备注
	if len(in.Remark) > 0 {
		whereBuilder = whereBuilder.Where(squirrel.Like{
			"remark ": "%" + in.Remark + "%",
		})
	}

	all, err := l.svcCtx.SysInterfaceModel.FindList(l.ctx, whereBuilder, in.Current, in.PageSize)
	if err != nil {
		return nil, err
	}

	countBuilder := l.svcCtx.SysInterfaceModel.CountBuilder("id")

	countBuilder = countBuilder.Where("deleted_at is null")

	// 接口名称
	if len(in.Name) > 0 {
		countBuilder = countBuilder.Where(squirrel.Like{
			"name ": "%" + in.Name + "%",
		})
	}
	// 接口地址
	if len(in.Path) > 0 {
		countBuilder = countBuilder.Where(squirrel.Like{
			"path ": "%" + in.Path + "%",
		})
	}
	// 接口类型
	if len(in.InterfaceType) > 0 {
		countBuilder = countBuilder.Where(squirrel.Like{
			"interface_type ": "%" + in.InterfaceType + "%",
		})
	}
	// 接口分组名称
	if len(in.InterfaceGroupName) > 0 {
		countBuilder = countBuilder.Where(squirrel.Like{
			"interface_group_name ": "%" + in.InterfaceGroupName + "%",
		})
	}
	// 备注
	if len(in.Remark) > 0 {
		countBuilder = countBuilder.Where(squirrel.Like{
			"remark ": "%" + in.Remark + "%",
		})
	}
	count, err := l.svcCtx.SysInterfaceModel.FindCount(l.ctx, countBuilder)
	if err != nil {
		return nil, err
	}

	var list []*tpmtclient.SysInterfaceListData
	for _, item := range all {
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

	return &tpmtclient.SysInterfaceListResp{
		Total: count,
		List:  list,
	}, nil
}
