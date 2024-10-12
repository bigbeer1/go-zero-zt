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

type TpmtGatewayAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewTpmtGatewayAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TpmtGatewayAddLogic {
	return &TpmtGatewayAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 网关
func (l *TpmtGatewayAddLogic) TpmtGatewayAdd(in *tpmtclient.TpmtGatewayAddReq) (*tpmtclient.CommonResp, error) {

	_, err := l.svcCtx.TpmtGatewayModel.Insert(l.ctx, &model.TpmtGateway{
		Id:           uuid.NewV4().String(), // ID
		CreatedAt:    time.Now(),            // 创建时间
		CreatedName:  in.CreatedName,        // 创建人
		GatewayName:  in.GatewayName,        // 网关名称
		GatewayModel: in.GatewayModel,       // 网关型号
		ManuFacturer: in.ManuFacturer,       // 生产厂家
		Agreement:    in.Agreement,          // 协议 默认1:modbus
		BaudRate:     in.BaudRate,           // 波特率
		Parity:       in.Parity,             // 校验
		DataBits:     in.DataBits,           // 数据位
		StopBits:     in.StopBits,           // 停止位
		ComPort:      in.ComPort,            // com端口
		AddressCode:  in.AddressCode,        // 地址码
	})
	if err != nil {
		return nil, err
	}

	return &tpmtclient.CommonResp{}, nil
}
