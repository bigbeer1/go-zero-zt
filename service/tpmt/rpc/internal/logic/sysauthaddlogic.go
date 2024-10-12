package logic

import (
	"context"
	uuid "github.com/satori/go.uuid"
	"time"
	"tpmt-zt/service/tpmt/model"

	"tpmt-zt/service/tpmt/rpc/internal/svc"
	"tpmt-zt/service/tpmt/rpc/tpmtclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type SysAuthAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSysAuthAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SysAuthAddLogic {
	return &SysAuthAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 第三方用户
func (l *SysAuthAddLogic) SysAuthAdd(in *tpmtclient.SysAuthAddReq) (*tpmtclient.CommonResp, error) {

	_, err := l.svcCtx.SysAuthModel.Insert(l.ctx, &model.SysAuth{
		Id:          uuid.NewV4().String(), // ID
		CreatedAt:   time.Now(),            // 创建时间
		CreatedName: in.CreatedName,        // 创建人
		NickName:    in.NickName,           // 机构名
		AuthToken:   in.AuthToken,          // 令牌
		State:       in.State,              // 状态 1:正常 2:停用 3:封禁
	})
	if err != nil {
		return nil, err
	}

	return &tpmtclient.CommonResp{}, nil
}
