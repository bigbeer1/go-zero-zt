package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"tpmt-zt/common/tdenginex"
	"tpmt-zt/service/archive/model"
	"tpmt-zt/service/archive/rpc/archiveclient"
	"tpmt-zt/service/archive/rpc/internal/svc"
)

type AppLoggerFindListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAppLoggerFindListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AppLoggerFindListLogic {
	return &AppLoggerFindListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AppLoggerFindListLogic) AppLoggerFindList(in *archiveclient.AppLoggerFindListReq) (*archiveclient.AppLoggerFindListResp, error) {

	appLog := model.TdAppLog{
		Uid:              in.Uid,
		Ip:               in.Ip,
		InterfaceType:    in.InterfaceType,
		InterfaceAddress: in.InterfaceAddress,
		IsRequest:        in.IsRequest,
	}

	tdDb := &model.TdDb{
		TableName: "",
		DbName:    "app_log.logger",
	}

	all, err := appLog.FindList(l.ctx, l.svcCtx.Taos, tdDb, in.Current, in.PageSize, in.StartTime, in.EndTime)
	if err != nil {
		if err.Error() != tdenginex.ErrNotFoundTable {
			return nil, err
		}
	}

	total := appLog.Count(l.ctx, l.svcCtx.Taos, tdDb, in.StartTime, in.EndTime)

	var datas []*archiveclient.LoggerData
	for _, item := range all {
		datas = append(datas, &archiveclient.LoggerData{
			Uid:              item.Uid,
			CreatedTime:      item.CreatedTime.UnixMilli(),
			CreatedName:      item.CreatedName,
			Ip:               item.Ip,
			InterfaceType:    item.InterfaceType,
			InterfaceAddress: item.InterfaceAddress,
			RequestData:      item.RequestData,
			IsRequest:        item.IsRequest,
			ResponseData:     item.ResponseData,
			Timed:            item.Timed,
		})

	}

	return &archiveclient.AppLoggerFindListResp{
		Total: total,
		List:  datas,
	}, nil
}
