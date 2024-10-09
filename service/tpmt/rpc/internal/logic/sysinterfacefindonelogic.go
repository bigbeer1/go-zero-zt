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

type SysInterfaceFindOneLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSysInterfaceFindOneLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SysInterfaceFindOneLogic {
	return &SysInterfaceFindOneLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SysInterfaceFindOneLogic) SysInterfaceFindOne(in *tpmtclient.SysInterfaceFindOneReq) (*tpmtclient.SysInterfaceFindOneResp, error) {
	res, err := l.svcCtx.SysInterfaceModel.FindOne(l.ctx, in.Id)
	if err != nil {
		if errors.Is(err, sqlc.ErrNotFound) {
			return nil, fmt.Errorf("SysInterface没有该ID:%v", in.Id)
		}
		return nil, err
	}

	// 判断该数据是否被删除
	if res.DeletedAt.Valid == true {
		return nil, fmt.Errorf("SysInterface该ID已被删除：%v", in.Id)
	}

	return &tpmtclient.SysInterfaceFindOneResp{
		Id:                 res.Id,                         //接口ID
		CreatedAt:          res.CreatedAt.UnixMilli(),      //创建时间
		UpdatedAt:          res.UpdatedAt.Time.UnixMilli(), //更新时间
		CreatedName:        res.CreatedName,                //创建人
		UpdatedName:        res.UpdatedName.String,         //更新人
		Name:               res.Name,                       //接口名称
		Path:               res.Path,                       //接口地址
		InterfaceType:      res.InterfaceType,              //接口类型
		InterfaceGroupName: res.InterfaceGroupName.String,  //接口分组名称
		Remark:             res.Remark.String,              //备注
		Sort:               res.Sort,                       //sort
	}, nil
}
