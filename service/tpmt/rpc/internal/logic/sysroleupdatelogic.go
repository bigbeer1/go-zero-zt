package logic

import (
	"context"
	"errors"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"time"

	"tpmt-zt/service/tpmt/rpc/internal/svc"
	"tpmt-zt/service/tpmt/rpc/tpmtclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type SysRoleUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSysRoleUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SysRoleUpdateLogic {
	return &SysRoleUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SysRoleUpdateLogic) SysRoleUpdate(in *tpmtclient.SysRoleUpdateReq) (*tpmtclient.CommonResp, error) {
	res, err := l.svcCtx.SysRoleModel.FindOne(l.ctx, in.Id)
	if err != nil {
		if errors.Is(err, sqlc.ErrNotFound) {
			return nil, fmt.Errorf("SysRole没有该ID: %v", in.Id)
		}
		return nil, err
	}

	// 判断该数据是否被删除
	if res.DeletedAt.Valid == true {
		return nil, fmt.Errorf("SysRole该ID已被删除： %v", in.Id)
	}

	// 角色名称
	if len(in.Name) > 0 {
		res.Name = in.Name
	}
	// 备注
	if len(in.Remark) > 0 {
		res.Remark.String = in.Remark
		res.Remark.Valid = true
	}
	// 角色类型 1:管理员角色  2:普通角色  3:第三方角色
	if in.RoleType != 0 {
		res.RoleType = in.RoleType
	}

	res.UpdatedName.String = in.UpdatedName
	res.UpdatedName.Valid = true
	res.UpdatedAt.Time = time.Now()
	res.UpdatedAt.Valid = true

	err = l.svcCtx.SysRoleModel.Update(l.ctx, res)

	if err != nil {
		return nil, err
	}
	return &tpmtclient.CommonResp{}, nil
}
