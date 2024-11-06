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

type SysDictDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSysDictDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SysDictDeleteLogic {
	return &SysDictDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SysDictDeleteLogic) SysDictDelete(in *authenticationclient.SysDictDeleteReq) (*authenticationclient.CommonResp, error) {

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

	res.DeletedAt.Time = time.Now()
	res.DeletedAt.Valid = true
	res.DeletedName.String = in.DeletedName
	res.DeletedName.Valid = true

	err = l.svcCtx.SysDictModel.Update(l.ctx, res)
	if err != nil {
		return nil, err
	}

	return &authenticationclient.CommonResp{}, nil
}
