package logic

import (
	"context"
	"database/sql"
	"time"
	"tpmt-zt/service/tpmt/model"

	"tpmt-zt/service/tpmt/rpc/internal/svc"
	"tpmt-zt/service/tpmt/rpc/tpmtclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type SysInterfaceAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSysInterfaceAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SysInterfaceAddLogic {
	return &SysInterfaceAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 接口
func (l *SysInterfaceAddLogic) SysInterfaceAdd(in *tpmtclient.SysInterfaceAddReq) (*tpmtclient.CommonResp, error) {

	_, err := l.svcCtx.SysInterfaceModel.Insert(l.ctx, &model.SysInterface{
		CreatedAt:          time.Now(),                                                                        // 创建时间
		CreatedName:        in.CreatedName,                                                                    // 创建人
		Name:               in.Name,                                                                           // 接口名称
		Path:               in.Path,                                                                           // 接口地址
		InterfaceType:      in.InterfaceType,                                                                  // 接口类型
		InterfaceGroupName: sql.NullString{String: in.InterfaceGroupName, Valid: in.InterfaceGroupName != ""}, // 接口分组名称
		Remark:             sql.NullString{String: in.Remark, Valid: in.Remark != ""},                         // 备注
		Sort:               in.Sort,                                                                           // sort
	})
	if err != nil {
		return nil, err
	}

	return &tpmtclient.CommonResp{}, nil
}
