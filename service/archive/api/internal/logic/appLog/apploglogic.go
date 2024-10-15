package appLog

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
	"tpmt-zt/common"
	"tpmt-zt/common/msg"
	"tpmt-zt/service/archive/api/internal/svc"
	"tpmt-zt/service/archive/api/internal/types"
	"tpmt-zt/service/archive/rpc/archiveclient"
)

type AppLogLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAppLogLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AppLogLogic {
	return &AppLogLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AppLogLogic) AppLog(req *types.AppLogReqest) (resp *types.Response, err error) {
	// 用户登录信息
	if req.CreatedEndTime <= req.CreatedStartTime && req.CreatedEndTime != 0 && req.CreatedStartTime != 0 {
		return nil, common.NewDefaultError("结束时间不能小于等于开始时间")
	}

	res, err := l.svcCtx.ArchiveRpc.AppLoggerFindList(l.ctx, &archiveclient.AppLoggerFindListReq{
		Current:          req.Current,
		PageSize:         req.PageSize,
		StartTime:        req.CreatedStartTime,
		EndTime:          req.CreatedEndTime,
		Uid:              req.Uid,
		Ip:               req.Ip,
		InterfaceType:    req.InterfaceType,
		InterfaceAddress: req.InterfaceAddress,
		IsRequest:        req.IsRequest,
	})
	if err != nil {
		return nil, common.NewDefaultError(err.Error())
	}

	var result AppLoggerFindListResp

	_ = copier.Copy(&result, res)

	return &types.Response{
		Code: 0,
		Msg:  msg.Success,
		Data: result,
	}, nil
}

type AppLoggerFindListResp struct {
	Total int64         `json:"total"` //总数据量
	List  []*LoggerData `json:"list"`  //数据
}

type LoggerData struct {
	Uid              string `json:"uid"`               // 操作人ID
	CreatedTime      int64  `json:"created_time"`      // 创建人名称
	CreatedName      string `json:"created_name"`      // 创建人名称
	Ip               string `json:"ip"`                // 请求Ip
	InterfaceType    string `json:"interface_type"`    // 请求方法
	InterfaceAddress string `json:"interface_address"` // 请求地址
	RequestData      string `json:"request_data"`      // 请求参数
	IsRequest        int64  `json:"is_request"`        // 请求结果
	ResponseData     string `json:"response_data"`     // 返回参数
	Timed            int64  `json:"timed"`             // 运算时间
}
