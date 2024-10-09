package sysInterface

import (
	"context"
	"github.com/jinzhu/copier"
	"tpmt-zt/common"
	"tpmt-zt/common/msg"
	"tpmt-zt/service/tpmt/rpc/tpmtclient"

	"tpmt-zt/service/tpmt/api/internal/svc"
	"tpmt-zt/service/tpmt/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SysInterfaceListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSysInterfaceListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SysInterfaceListLogic {
	return &SysInterfaceListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SysInterfaceListLogic) SysInterfaceList(req *types.SysInterfaceListRequest) (resp *types.Response, err error) {

	all, err := l.svcCtx.TpmtRpc.SysInterfaceList(l.ctx, &tpmtclient.SysInterfaceListReq{
		Current:            req.Current,            // 页码
		PageSize:           req.PageSize,           // 页数
		Name:               req.Name,               // 接口名称
		Path:               req.Path,               // 接口地址
		InterfaceType:      req.InterfaceType,      // 接口类型
		InterfaceGroupName: req.InterfaceGroupName, // 接口分组名称
		Remark:             req.Remark,             // 备注
	})
	if err != nil {
		return nil, common.NewDefaultError(err.Error())
	}

	var result SysInterfaceListResp
	_ = copier.Copy(&result, all)

	return &types.Response{
		Code: 0,
		Msg:  msg.Success,
		Data: result,
	}, nil
}

type SysInterfaceListResp struct {
	Total int64                   `json:"total"`
	List  []*SysInterfaceDataList `json:"list"`
}

type SysInterfaceDataList struct {
	Id                 int64  `json:"id"`                   // 接口ID,
	CreatedAt          int64  `json:"created_at"`           // 创建时间,
	UpdatedAt          int64  `json:"updated_at"`           // 更新时间,
	CreatedName        string `json:"created_name"`         // 创建人,
	UpdatedName        string `json:"updated_name"`         // 更新人,
	Name               string `json:"name"`                 // 接口名称,
	Path               string `json:"path"`                 // 接口地址,
	InterfaceType      string `json:"interface_type"`       // 接口类型,
	InterfaceGroupName string `json:"interface_group_name"` // 接口分组名称,
	Remark             string `json:"remark"`               // 备注
}
