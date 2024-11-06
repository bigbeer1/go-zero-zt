package sysInterface

import (
	"context"
	"github.com/jinzhu/copier"
	"tpmt-zt/common"
	"tpmt-zt/common/msg"
	"tpmt-zt/service/authentication/authenticationclient"

	"tpmt-zt/service/tpmt/api/internal/svc"
	"tpmt-zt/service/tpmt/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SysInterfaceInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSysInterfaceInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SysInterfaceInfoLogic {
	return &SysInterfaceInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SysInterfaceInfoLogic) SysInterfaceInfo(req *types.SysInterfaceInfoRequest) (resp *types.Response, err error) {

	res, err := l.svcCtx.AuthenticationRpc.SysInterfaceFindOne(l.ctx, &authenticationclient.SysInterfaceFindOneReq{
		Id: req.Id, // 接口ID
	})
	if err != nil {
		return nil, common.NewDefaultError(err.Error())
	}

	var result SysInterfaceFindOneResp
	_ = copier.Copy(&result, res)

	return &types.Response{
		Code: 0,
		Msg:  msg.Success,
		Data: result,
	}, nil
}

type SysInterfaceFindOneResp struct {
	Id                 int64  `json:"id"`                   // 接口ID,
	CreatedAt          int64  `json:"created_at"`           // 创建时间,
	UpdatedAt          int64  `json:"updated_at"`           // 更新时间,
	CreatedName        string `json:"created_name"`         // 创建人,
	UpdatedName        string `json:"updated_name"`         // 更新人,
	Name               string `json:"name"`                 // 接口名称,
	Path               string `json:"path"`                 // 接口地址,
	InterfaceType      string `json:"interface_type"`       // 接口类型,
	InterfaceGroupName string `json:"interface_group_name"` // 接口分组名称,
	Remark             string `json:"remark"`               // 备注,
	Sort               int64  `json:"sort"`                 // sort
}
