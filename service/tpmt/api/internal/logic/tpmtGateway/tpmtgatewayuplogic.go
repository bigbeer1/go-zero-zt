package tpmtGateway

import (
	"context"
	"tpmt-zt/common"
	"tpmt-zt/common/jwtx"
	"tpmt-zt/common/msg"
	"tpmt-zt/service/tpmt/rpc/tpmtclient"

	"tpmt-zt/service/tpmt/api/internal/svc"
	"tpmt-zt/service/tpmt/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type TpmtGatewayUpLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTpmtGatewayUpLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TpmtGatewayUpLogic {
	return &TpmtGatewayUpLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TpmtGatewayUpLogic) TpmtGatewayUp(req *types.TpmtGatewayUpRequest) (resp *types.Response, err error) {
	// 用户登录信息
	tokenData := jwtx.ParseToken(l.ctx)

	_, err = l.svcCtx.TpmtRpc.TpmtGatewayUpdate(l.ctx, &tpmtclient.TpmtGatewayUpdateReq{
		Id:           req.Id,             // 采集器ID/网关
		UpdatedName:  tokenData.NickName, // 更新人
		GatewayName:  req.GatewayName,    // 网关名称
		GatewayModel: req.GatewayModel,   // 网关型号
		ManuFacturer: req.ManuFacturer,   // 生产厂家
		Agreement:    req.Agreement,      // 协议 默认1:modbus
		BaudRate:     req.BaudRate,       // 波特率
		Parity:       req.Parity,         // 校验
		DataBits:     req.DataBits,       // 数据位
		StopBits:     req.StopBits,       // 停止位
		ComPort:      req.ComPort,        // com端口
		AddressCode:  req.AddressCode,    // 地址码
	})
	if err != nil {
		return nil, common.NewDefaultError(err.Error())
	}
	return &types.Response{
		Code: 0,
		Msg:  msg.Success,
		Data: nil,
	}, nil
}
