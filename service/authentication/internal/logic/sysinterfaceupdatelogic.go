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

type SysInterfaceUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSysInterfaceUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SysInterfaceUpdateLogic {
	return &SysInterfaceUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SysInterfaceUpdateLogic) SysInterfaceUpdate(in *authenticationclient.SysInterfaceUpdateReq) (*authenticationclient.CommonResp, error) {

	res, err := l.svcCtx.SysInterfaceModel.FindOne(l.ctx, in.Id)
	if err != nil {
		if errors.Is(err, sqlc.ErrNotFound) {
			return nil, fmt.Errorf("SysInterface没有该ID: %v", in.Id)
		}
		return nil, err
	}

	// 判断该数据是否被删除
	if res.DeletedAt.Valid == true {
		return nil, fmt.Errorf("SysInterface该ID已被删除：%v", in.Id)
	}

	// 接口名称
	if len(in.Name) > 0 {
		res.Name = in.Name
	}
	// 接口地址
	if len(in.Path) > 0 {
		res.Path = in.Path
	}
	// 接口类型
	if len(in.InterfaceType) > 0 {
		res.InterfaceType = in.InterfaceType
	}
	// 接口分组名称
	if len(in.InterfaceGroupName) > 0 {
		res.InterfaceGroupName.String = in.InterfaceGroupName
		res.InterfaceGroupName.Valid = true
	}
	// 备注
	if len(in.Remark) > 0 {
		res.Remark.String = in.Remark
		res.Remark.Valid = true
	}
	// sort
	if in.Sort != 0 {
		res.Sort = in.Sort
	}

	res.UpdatedName.String = in.UpdatedName
	res.UpdatedName.Valid = true
	res.UpdatedAt.Time = time.Now()
	res.UpdatedAt.Valid = true

	err = l.svcCtx.SysInterfaceModel.Update(l.ctx, res)

	if err != nil {
		return nil, err
	}
	return &authenticationclient.CommonResp{}, nil
}
