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

type SysDictTypeDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSysDictTypeDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SysDictTypeDeleteLogic {
	return &SysDictTypeDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SysDictTypeDeleteLogic) SysDictTypeDelete(in *authenticationclient.SysDictTypeDeleteReq) (*authenticationclient.CommonResp, error) {

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

	res.DeletedAt.Time = time.Now()
	res.DeletedAt.Valid = true
	res.DeletedName.String = in.DeletedName
	res.DeletedName.Valid = true

	err = l.svcCtx.SysDictTypeModel.Update(l.ctx, res)
	if err != nil {
		return nil, err
	}

	return &authenticationclient.CommonResp{}, nil
}
