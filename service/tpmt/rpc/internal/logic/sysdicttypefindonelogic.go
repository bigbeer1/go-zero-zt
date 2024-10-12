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

type SysDictTypeFindOneLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSysDictTypeFindOneLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SysDictTypeFindOneLogic {
	return &SysDictTypeFindOneLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SysDictTypeFindOneLogic) SysDictTypeFindOne(in *tpmtclient.SysDictTypeFindOneReq) (*tpmtclient.SysDictTypeFindOneResp, error) {

	res, err := l.svcCtx.SysDictTypeModel.FindOne(l.ctx, in.Id)
	if err != nil {
		if errors.Is(err, sqlc.ErrNotFound) {
			return nil, fmt.Errorf("SysDictType没有该ID:%v", in.Id)
		}
		return nil, err
	}

	// 判断该数据是否被删除
	if res.DeletedAt.Valid == true {
		return nil, fmt.Errorf("SysDictType该ID已被删除：%v", in.Id)
	}

	return &tpmtclient.SysDictTypeFindOneResp{
		Id:          res.Id,                         //字典类型ID
		CreatedAt:   res.CreatedAt.UnixMilli(),      //创建时间
		UpdatedAt:   res.UpdatedAt.Time.UnixMilli(), //更新时间
		CreatedName: res.CreatedName,                //创建人
		UpdatedName: res.UpdatedName.String,         //更新人
		Name:        res.Name,                       //字典名称
		DictType:    res.DictType,                   //字典类型
		State:       res.State,                      //状态
		Remark:      res.Remark.String,              //描述
		Sort:        res.Sort,                       //排序
	}, nil

}
