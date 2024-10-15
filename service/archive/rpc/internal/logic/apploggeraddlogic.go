package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"time"
	"tpmt-zt/service/archive/model"
	"tpmt-zt/service/archive/rpc/archiveclient"
	"tpmt-zt/service/archive/rpc/internal/svc"
)

type AppLoggerAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAppLoggerAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AppLoggerAddLogic {
	return &AppLoggerAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 用户日志
func (l *AppLoggerAddLogic) AppLoggerAdd(in *archiveclient.AppLoggerAddReq) (*archiveclient.CommonResp, error) {

	// 讲uuid中- 全部替换为_  确保插入和查询成功

	tddb := &model.TdDb{
		TableName: "app_log.tpmt_app_log",
		DbName:    "app_log.logger",
	}

	data := &model.TdAppLog{
		CreatedTime:      time.Now(),
		Uid:              in.Uid,
		CreatedName:      in.CreatedName,
		Ip:               in.Ip,
		InterfaceType:    in.InterfaceType,
		InterfaceAddress: in.InterfaceAddress,
		IsRequest:        in.IsRequest,
		RequestData:      in.RequestData,
		ResponseData:     in.ResponseData,
		Timed:            in.Timed,
	}

	err := data.Insert(l.ctx, l.svcCtx.Taos, tddb)
	if err != nil {
		return nil, err
	}

	return &archiveclient.CommonResp{}, nil
}
