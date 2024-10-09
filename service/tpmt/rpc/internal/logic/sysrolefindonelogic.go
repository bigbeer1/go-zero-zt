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

type SysRoleFindOneLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSysRoleFindOneLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SysRoleFindOneLogic {
	return &SysRoleFindOneLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SysRoleFindOneLogic) SysRoleFindOne(in *tpmtclient.SysRoleFindOneReq) (*tpmtclient.SysRoleFindOneResp, error) {
	res, err := l.svcCtx.SysRoleModel.FindOne(l.ctx, in.Id)
	if err != nil {
		if errors.Is(err, sqlc.ErrNotFound) {
			return nil, fmt.Errorf("SysRole没有该ID:%v", in.Id)
		}
		return nil, err
	}

	// 判断该数据是否被删除
	if res.DeletedAt.Valid == true {
		return nil, fmt.Errorf("SysRole该ID已被删除：%v", in.Id)
	}

	return &tpmtclient.SysRoleFindOneResp{
		Id:          res.Id,                         //角色ID
		Name:        res.Name,                       //角色名称
		Remark:      res.Remark.String,              //备注
		RoleType:    res.RoleType,                   //角色类型 1:管理员角色  2:普通角色  3:第三方角色
		CreatedName: res.CreatedName,                //创建人
		CreatedAt:   res.CreatedAt.UnixMilli(),      //创建时间
		UpdatedName: res.UpdatedName.String,         //更新人
		UpdatedAt:   res.UpdatedAt.Time.UnixMilli(), //更新时间
	}, nil
}
