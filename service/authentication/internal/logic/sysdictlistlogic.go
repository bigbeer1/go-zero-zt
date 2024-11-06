package logic

import (
	"context"
	"github.com/Masterminds/squirrel"
	"tpmt-zt/service/authentication/authenticationclient"
	"tpmt-zt/service/authentication/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type SysDictListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSysDictListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SysDictListLogic {
	return &SysDictListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SysDictListLogic) SysDictList(in *authenticationclient.SysDictListReq) (*authenticationclient.SysDictListResp, error) {

	whereBuilder := l.svcCtx.SysDictModel.RowBuilder()

	whereBuilder = whereBuilder.Where("deleted_at is null")
	whereBuilder = whereBuilder.OrderBy("created_at DESC, id DESC")

	// 字典类型
	if len(in.DictType) > 0 {
		whereBuilder = whereBuilder.Where(squirrel.Like{
			"dict_type ": "%" + in.DictType + "%",
		})
	}
	// 字典标签
	if len(in.DictLabel) > 0 {
		whereBuilder = whereBuilder.Where(squirrel.Like{
			"dict_label ": "%" + in.DictLabel + "%",
		})
	}
	// 字典键值
	if len(in.DictValue) > 0 {
		whereBuilder = whereBuilder.Where(squirrel.Like{
			"dict_value ": "%" + in.DictValue + "%",
		})
	}
	// 备注
	if len(in.Remark) > 0 {
		whereBuilder = whereBuilder.Where(squirrel.Like{
			"remark ": "%" + in.Remark + "%",
		})
	}
	// 状态
	if in.State != 99 {
		whereBuilder = whereBuilder.Where(squirrel.Eq{
			"state ": in.State,
		})
	}

	all, err := l.svcCtx.SysDictModel.FindList(l.ctx, whereBuilder, in.Current, in.PageSize)
	if err != nil {
		return nil, err
	}

	countBuilder := l.svcCtx.SysDictModel.CountBuilder("id")

	countBuilder = countBuilder.Where("deleted_at is null")

	// 字典类型
	if len(in.DictType) > 0 {
		countBuilder = countBuilder.Where(squirrel.Like{
			"dict_type ": "%" + in.DictType + "%",
		})
	}
	// 字典标签
	if len(in.DictLabel) > 0 {
		countBuilder = countBuilder.Where(squirrel.Like{
			"dict_label ": "%" + in.DictLabel + "%",
		})
	}
	// 字典键值
	if len(in.DictValue) > 0 {
		countBuilder = countBuilder.Where(squirrel.Like{
			"dict_value ": "%" + in.DictValue + "%",
		})
	}
	// 备注
	if len(in.Remark) > 0 {
		countBuilder = countBuilder.Where(squirrel.Like{
			"remark ": "%" + in.Remark + "%",
		})
	}
	// 状态
	if in.State != 99 {
		countBuilder = countBuilder.Where(squirrel.Eq{
			"state ": in.State,
		})
	}
	count, err := l.svcCtx.SysDictModel.FindCount(l.ctx, countBuilder)
	if err != nil {
		return nil, err
	}

	var list []*authenticationclient.SysDictListData
	for _, item := range all {
		list = append(list, &authenticationclient.SysDictListData{
			Id:          item.Id,                         //字典类型ID
			CreatedAt:   item.CreatedAt.UnixMilli(),      //创建时间
			UpdatedAt:   item.UpdatedAt.Time.UnixMilli(), //更新时间
			CreatedName: item.CreatedName,                //创建人
			UpdatedName: item.UpdatedName.String,         //更新人
			DictType:    item.DictType,                   //字典类型
			DictLabel:   item.DictLabel,                  //字典标签
			DictValue:   item.DictValue,                  //字典键值
			Sort:        item.Sort,                       //排序
			Remark:      item.Remark.String,              //备注
			State:       item.State,                      //状态
		})
	}

	return &authenticationclient.SysDictListResp{
		Total: count,
		List:  list,
	}, nil
}
