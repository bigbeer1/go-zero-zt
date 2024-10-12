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

type SysAuthFindOneLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSysAuthFindOneLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SysAuthFindOneLogic {
	return &SysAuthFindOneLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SysAuthFindOneLogic) SysAuthFindOne(in *tpmtclient.SysAuthFindOneReq) (*tpmtclient.SysAuthFindOneResp, error) {

	res, err := l.svcCtx.SysAuthModel.FindOne(l.ctx, in.Id)
	if err != nil {
		if errors.Is(err, sqlc.ErrNotFound) {
			return nil, fmt.Errorf("SysAuth没有该ID:%v", in.Id)
		}
		return nil, err
	}

	// 判断该数据是否被删除
	if res.DeletedAt.Valid == true {
		return nil, fmt.Errorf("SysAuth该ID已被删除：%v", in.Id)
	}

	return &tpmtclient.SysAuthFindOneResp{
		Id:          res.Id,                         //第三方用户ID
		CreatedAt:   res.CreatedAt.UnixMilli(),      //创建时间
		UpdatedAt:   res.UpdatedAt.Time.UnixMilli(), //更新时间
		CreatedName: res.CreatedName,                //创建人
		UpdatedName: res.UpdatedName.String,         //更新人
		NickName:    res.NickName,                   //机构名
		AuthToken:   res.AuthToken,                  //令牌
		State:       res.State,                      //状态 1:正常 2:停用 3:封禁
	}, nil

}
