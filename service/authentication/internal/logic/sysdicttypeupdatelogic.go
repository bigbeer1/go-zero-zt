package logic

import (
	"context"
	"errors"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"time"
	"tpmt-zt/service/authentication/authenticationclient"
	"tpmt-zt/service/authentication/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type SysDictTypeUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSysDictTypeUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SysDictTypeUpdateLogic {
	return &SysDictTypeUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SysDictTypeUpdateLogic) SysDictTypeUpdate(in *authenticationclient.SysDictTypeUpdateReq) (*authenticationclient.CommonResp, error) {

	res, err := l.svcCtx.SysDictTypeModel.FindOne(l.ctx, in.Id)
	if err != nil {
		if errors.Is(err, sqlc.ErrNotFound) {
			return nil, fmt.Errorf("SysDictType没有该ID: %v", in.Id)
		}
		return nil, err
	}

	// 判断该数据是否被删除
	if res.DeletedAt.Valid == true {
		return nil, fmt.Errorf("SysDictType该ID已被删除：%v", in.Id)
	}

	// 字典名称
	if len(in.Name) > 0 {
		res.Name = in.Name
	}
	// 字典类型
	if len(in.DictType) > 0 {
		res.DictType = in.DictType
	}
	// 状态
	if in.State != 0 {
		res.State = in.State
	}
	// 描述
	if len(in.Remark) > 0 {
		res.Remark.String = in.Remark
		res.Remark.Valid = true
	}
	// 排序
	if in.Sort != 0 {
		res.Sort = in.Sort
	}

	res.UpdatedName.String = in.UpdatedName
	res.UpdatedName.Valid = true
	res.UpdatedAt.Time = time.Now()
	res.UpdatedAt.Valid = true

	err = l.svcCtx.SysDictTypeModel.Update(l.ctx, res)

	if err != nil {
		return nil, err
	}
	return &authenticationclient.CommonResp{}, nil
}
