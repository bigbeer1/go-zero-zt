package sysDictType

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

type SysDictTypeListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSysDictTypeListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SysDictTypeListLogic {
	return &SysDictTypeListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SysDictTypeListLogic) SysDictTypeList(req *types.SysDictTypeListRequest) (resp *types.Response, err error) {

	all, err := l.svcCtx.AuthenticationRpc.SysDictTypeList(l.ctx, &authenticationclient.SysDictTypeListReq{
		Current:  req.Current,  // 页码
		PageSize: req.PageSize, // 页数
		Name:     req.Name,     // 字典名称
		DictType: req.DictType, // 字典类型
		State:    req.State,    // 状态
		Remark:   req.Remark,   // 描述
	})
	if err != nil {
		return nil, common.NewDefaultError(err.Error())
	}

	var result SysDictTypeListResp
	_ = copier.Copy(&result, all)

	return &types.Response{
		Code: 0,
		Msg:  msg.Success,
		Data: result,
	}, nil
}

type SysDictTypeListResp struct {
	Total int64                  `json:"total"`
	List  []*SysDictTypeDataList `json:"list"`
}

type SysDictTypeDataList struct {
	Id          int64  `json:"id"`           // 字典类型ID,
	CreatedAt   int64  `json:"created_at"`   // 创建时间,
	UpdatedAt   int64  `json:"updated_at"`   // 更新时间,
	CreatedName string `json:"created_name"` // 创建人,
	UpdatedName string `json:"updated_name"` // 更新人,
	Name        string `json:"name"`         // 字典名称,
	DictType    string `json:"dict_type"`    // 字典类型,
	State       int64  `json:"state"`        // 状态,
	Remark      string `json:"remark"`       // 描述
}
