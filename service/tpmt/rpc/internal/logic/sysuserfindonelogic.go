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

type SysUserFindOneLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSysUserFindOneLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SysUserFindOneLogic {
	return &SysUserFindOneLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SysUserFindOneLogic) SysUserFindOne(in *tpmtclient.SysUserFindOneReq) (*tpmtclient.SysUserFindOneResp, error) {
	res, err := l.svcCtx.SysUserModel.FindOne(l.ctx, in.Id)
	if err != nil {
		if errors.Is(err, sqlc.ErrNotFound) {
			return nil, fmt.Errorf("SysUser没有该ID:%v", in.Id)
		}
		return nil, err
	}

	// 判断该数据是否被删除
	if res.DeletedAt.Valid == true {
		return nil, errors.New("SysUser该ID已被删除：" + in.Id)
	}

	return &tpmtclient.SysUserFindOneResp{
		Id:          res.Id,                         //用户ID
		Account:     res.Account,                    //用户名
		NickName:    res.NickName,                   //姓名
		Password:    res.Password,                   //密码
		State:       res.State,                      //状态 1:正常 2:停用 3:封禁
		CreatedName: res.CreatedName,                //创建人
		CreatedAt:   res.CreatedAt.UnixMilli(),      //创建时间
		UpdatedName: res.UpdatedName.String,         //更新人
		UpdatedAt:   res.UpdatedAt.Time.UnixMilli(), //更新时间
	}, nil
}
