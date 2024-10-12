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

type TpmtGatewayAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTpmtGatewayAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TpmtGatewayAddLogic {
	return &TpmtGatewayAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TpmtGatewayAddLogic) TpmtGatewayAdd(req *types.TpmtGatewayAddRequest) (resp *types.Response, err error) {
	// 用户登录信息
	tokenData := jwtx.ParseToken(l.ctx)

	_, err = l.svcCtx.TpmtRpc.TpmtGatewayAdd(l.ctx, &tpmtclient.TpmtGatewayAddReq{
		CreatedName:  tokenData.NickName, // 创建人
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
