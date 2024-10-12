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

type SysAuthUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSysAuthUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SysAuthUpdateLogic {
	return &SysAuthUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SysAuthUpdateLogic) SysAuthUpdate(in *tpmtclient.SysAuthUpdateReq) (*tpmtclient.CommonResp, error) {

	res, err := l.svcCtx.SysAuthModel.FindOne(l.ctx, in.Id)
	if err != nil {
		if errors.Is(err, sqlc.ErrNotFound) {
			return nil, fmt.Errorf("SysAuth没有该ID: %v", in.Id)
		}
		return nil, err
	}

	// 判断该数据是否被删除
	if res.DeletedAt.Valid == true {
		return nil, errors.New("SysAuth该ID已被删除：" + in.Id)
	}

	// 机构名
	if len(in.NickName) > 0 {
		res.NickName = in.NickName
	}
	// 令牌
	if len(in.AuthToken) > 0 {
		res.AuthToken = in.AuthToken
	}
	// 状态 1:正常 2:停用 3:封禁
	if in.State != 0 {
		res.State = in.State
	}

	res.UpdatedName.String = in.UpdatedName
	res.UpdatedName.Valid = true
	res.UpdatedAt.Time = time.Now()
	res.UpdatedAt.Valid = true

	err = l.svcCtx.SysAuthModel.Update(l.ctx, res)

	if err != nil {
		return nil, err
	}
	return &tpmtclient.CommonResp{}, nil
}
