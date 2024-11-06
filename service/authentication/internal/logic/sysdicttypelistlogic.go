package logic

import (
	"context"
	"github.com/Masterminds/squirrel"
	"tpmt-zt/service/authentication/authenticationclient"
	"tpmt-zt/service/authentication/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type SysDictTypeListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSysDictTypeListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SysDictTypeListLogic {
	return &SysDictTypeListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SysDictTypeListLogic) SysDictTypeList(in *authenticationclient.SysDictTypeListReq) (*authenticationclient.SysDictTypeListResp, error) {

	whereBuilder := l.svcCtx.SysDictTypeModel.RowBuilder()

	whereBuilder = whereBuilder.Where("deleted_at is null")
	whereBuilder = whereBuilder.OrderBy("created_at DESC, id DESC")

	// 字典名称
	if len(in.Name) > 0 {
		whereBuilder = whereBuilder.Where(squirrel.Like{
			"name ": "%" + in.Name + "%",
		})
	}
	// 字典类型
	if len(in.DictType) > 0 {
		whereBuilder = whereBuilder.Where(squirrel.Like{
			"dict_type ": "%" + in.DictType + "%",
		})
	}
	// 状态
	if in.State != 99 {
		whereBuilder = whereBuilder.Where(squirrel.Eq{
			"state ": in.State,
		})
	}
	// 描述
	if len(in.Remark) > 0 {
		whereBuilder = whereBuilder.Where(squirrel.Like{
			"remark ": "%" + in.Remark + "%",
		})
	}

	all, err := l.svcCtx.SysDictTypeModel.FindList(l.ctx, whereBuilder, in.Current, in.PageSize)
	if err != nil {
		return nil, err
	}

	countBuilder := l.svcCtx.SysDictTypeModel.CountBuilder("id")

	countBuilder = countBuilder.Where("deleted_at is null")

	// 字典名称
	if len(in.Name) > 0 {
		countBuilder = countBuilder.Where(squirrel.Like{
			"name ": "%" + in.Name + "%",
		})
	}
	// 字典类型
	if len(in.DictType) > 0 {
		countBuilder = countBuilder.Where(squirrel.Like{
			"dict_type ": "%" + in.DictType + "%",
		})
	}
	// 状态
	if in.State != 99 {
		countBuilder = countBuilder.Where(squirrel.Eq{
			"state ": in.State,
		})
	}
	// 描述
	if len(in.Remark) > 0 {
		countBuilder = countBuilder.Where(squirrel.Like{
			"remark ": "%" + in.Remark + "%",
		})
	}
	count, err := l.svcCtx.SysDictTypeModel.FindCount(l.ctx, countBuilder)
	if err != nil {
		return nil, err
	}

	var list []*authenticationclient.SysDictTypeListData
	for _, item := range all {
		list = append(list, &authenticationclient.SysDictTypeListData{
			Id:          item.Id,                         //字典类型ID
			CreatedAt:   item.CreatedAt.UnixMilli(),      //创建时间
			UpdatedAt:   item.UpdatedAt.Time.UnixMilli(), //更新时间
			CreatedName: item.CreatedName,                //创建人
			UpdatedName: item.UpdatedName.String,         //更新人
			Name:        item.Name,                       //字典名称
			DictType:    item.DictType,                   //字典类型
			State:       item.State,                      //状态
			Remark:      item.Remark.String,              //描述
			Sort:        item.Sort,                       //排序
		})
	}

	return &authenticationclient.SysDictTypeListResp{
		Total: count,
		List:  list,
	}, nil
}
