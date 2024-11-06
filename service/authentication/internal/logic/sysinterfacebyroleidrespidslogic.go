package logic

import (
	"context"
	"github.com/Masterminds/squirrel"

	"tpmt-zt/service/authentication/authenticationclient"
	"tpmt-zt/service/authentication/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type SysInterfaceByRoleIdRespIDsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSysInterfaceByRoleIdRespIDsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SysInterfaceByRoleIdRespIDsLogic {
	return &SysInterfaceByRoleIdRespIDsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 通过角色ID获取接口IDS
func (l *SysInterfaceByRoleIdRespIDsLogic) SysInterfaceByRoleIdRespIDs(in *authenticationclient.SysInterfaceByRoleIdReq) (*authenticationclient.SysInterfaceByRoleIdRespIDsResp, error) {
	// 先去中间表找对应的接口IDS
	whereBuilder := l.svcCtx.SysRoleInterfaceModel.RowBuilder()
	whereBuilder = whereBuilder.OrderBy("created_at DESC, id DESC")
	// 接口名称
	whereBuilder = whereBuilder.Where(squirrel.Eq{
		"role_id ": in.RoleId,
	})

	all, err := l.svcCtx.SysRoleInterfaceModel.FindAll(l.ctx, whereBuilder)
	if err != nil {
		return nil, err
	}

	var ids []int64
	for _, v := range all {
		ids = append(ids, v.InterfaceId)
	}

	return &authenticationclient.SysInterfaceByRoleIdRespIDsResp{
		Ids: ids,
	}, nil
}
