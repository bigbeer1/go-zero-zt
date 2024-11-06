package logic

import (
	"context"
	"errors"
	"github.com/zeromicro/go-zero/core/stores/sqlc"

	"tpmt-zt/service/authentication/authenticationclient"
	"tpmt-zt/service/authentication/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindOneInterfaceByPathAndInterfaceTypeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindOneInterfaceByPathAndInterfaceTypeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindOneInterfaceByPathAndInterfaceTypeLogic {
	return &FindOneInterfaceByPathAndInterfaceTypeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 根据请求路径/请求类型 查询
func (l *FindOneInterfaceByPathAndInterfaceTypeLogic) FindOneInterfaceByPathAndInterfaceType(in *authenticationclient.FindOneInterfaceByPathAndInterfaceTypeReq) (*authenticationclient.FindOneInterfaceByPathAndInterfaceTypeResp, error) {
	res, err := l.svcCtx.SysInterfaceModel.FindOneByPathAndInterfaceType(l.ctx, in.Path, in.InterfaceType)
	if err != nil {
		if err == sqlc.ErrNotFound {
			return nil, errors.New("未找到接口信息,鉴权失败")
		}
		return nil, err
	}

	return &authenticationclient.FindOneInterfaceByPathAndInterfaceTypeResp{
		Id:                 res.Id,
		CreatedAt:          res.CreatedAt.UnixMilli(),
		UpdatedAt:          res.UpdatedAt.Time.UnixMilli(),
		CreatedName:        res.CreatedName,
		UpdatedName:        res.UpdatedName.String,
		Name:               res.Name,
		Path:               res.Path,
		Remark:             res.Remark.String,
		InterfaceType:      res.InterfaceType,
		InterfaceGroupName: res.InterfaceGroupName.String,
		Sort:               res.Sort,
	}, nil
}
