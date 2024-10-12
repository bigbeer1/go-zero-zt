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

type SysDictFindOneLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSysDictFindOneLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SysDictFindOneLogic {
	return &SysDictFindOneLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SysDictFindOneLogic) SysDictFindOne(in *tpmtclient.SysDictFindOneReq) (*tpmtclient.SysDictFindOneResp, error) {

	res, err := l.svcCtx.SysDictModel.FindOne(l.ctx, in.Id)
	if err != nil {
		if errors.Is(err, sqlc.ErrNotFound) {
			return nil, fmt.Errorf("SysDict没有该ID:%v", in.Id)
		}
		return nil, err
	}

	// 判断该数据是否被删除
	if res.DeletedAt.Valid == true {
		return nil, fmt.Errorf("SysDict该ID已被删除：%v", in.Id)
	}

	return &tpmtclient.SysDictFindOneResp{
		Id:          res.Id,                         //字典类型ID
		CreatedAt:   res.CreatedAt.UnixMilli(),      //创建时间
		UpdatedAt:   res.UpdatedAt.Time.UnixMilli(), //更新时间
		CreatedName: res.CreatedName,                //创建人
		UpdatedName: res.UpdatedName.String,         //更新人
		DictType:    res.DictType,                   //字典类型
		DictLabel:   res.DictLabel,                  //字典标签
		DictValue:   res.DictValue,                  //字典键值
		Sort:        res.Sort,                       //排序
		Remark:      res.Remark.String,              //备注
		State:       res.State,                      //状态
	}, nil

}
