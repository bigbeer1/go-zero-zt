package sysDict

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

type SysDictListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSysDictListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SysDictListLogic {
	return &SysDictListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SysDictListLogic) SysDictList(req *types.SysDictListRequest) (resp *types.Response, err error) {

	all, err := l.svcCtx.AuthenticationRpc.SysDictList(l.ctx, &authenticationclient.SysDictListReq{
		Current:   req.Current,   // 页码
		PageSize:  req.PageSize,  // 页数
		DictType:  req.DictType,  // 字典类型
		DictLabel: req.DictLabel, // 字典标签
		DictValue: req.DictValue, // 字典键值
		Remark:    req.Remark,    // 备注
		State:     req.State,     // 状态
	})
	if err != nil {
		return nil, common.NewDefaultError(err.Error())
	}

	var result SysDictListResp
	_ = copier.Copy(&result, all)

	return &types.Response{
		Code: 0,
		Msg:  msg.Success,
		Data: result,
	}, nil
}

type SysDictListResp struct {
	Total int64              `json:"total"`
	List  []*SysDictDataList `json:"list"`
}

type SysDictDataList struct {
	Id          int64  `json:"id"`           // 字典类型ID,
	CreatedAt   int64  `json:"created_at"`   // 创建时间,
	UpdatedAt   int64  `json:"updated_at"`   // 更新时间,
	CreatedName string `json:"created_name"` // 创建人,
	UpdatedName string `json:"updated_name"` // 更新人,
	DictType    string `json:"dict_type"`    // 字典类型,
	DictLabel   string `json:"dict_label"`   // 字典标签,
	DictValue   string `json:"dict_value"`   // 字典键值,
	Remark      string `json:"remark"`       // 备注,
	State       int64  `json:"state"`        // 状态
}
