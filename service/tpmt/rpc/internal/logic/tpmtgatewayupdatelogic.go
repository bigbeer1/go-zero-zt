package logic

import (
	"context"
	"errors"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"time"

	"tpmt-zt/service/tpmt/rpc/internal/svc"
	"tpmt-zt/service/tpmt/rpc/tpmtclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type TpmtGatewayUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewTpmtGatewayUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TpmtGatewayUpdateLogic {
	return &TpmtGatewayUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *TpmtGatewayUpdateLogic) TpmtGatewayUpdate(in *tpmtclient.TpmtGatewayUpdateReq) (*tpmtclient.CommonResp, error) {

	res, err := l.svcCtx.TpmtGatewayModel.FindOne(l.ctx, in.Id)
	if err != nil {
		if errors.Is(err, sqlc.ErrNotFound) {
			return nil, fmt.Errorf("TpmtGateway没有该ID: %v", in.Id)
		}
		return nil, err
	}

	// 网关名称
	if len(in.GatewayName) > 0 {
		res.GatewayName = in.GatewayName
	}
	// 网关型号
	if len(in.GatewayModel) > 0 {
		res.GatewayModel = in.GatewayModel
	}
	// 生产厂家
	if len(in.ManuFacturer) > 0 {
		res.ManuFacturer = in.ManuFacturer
	}
	// 协议 默认1:modbus
	if in.Agreement != 0 {
		res.Agreement = in.Agreement
	}
	// 波特率
	if in.BaudRate != 0 {
		res.BaudRate = in.BaudRate
	}
	// 校验
	if len(in.Parity) > 0 {
		res.Parity = in.Parity
	}
	// 数据位
	if in.DataBits != 0 {
		res.DataBits = in.DataBits
	}
	// 停止位
	if in.StopBits != 0 {
		res.StopBits = in.StopBits
	}
	// com端口
	if len(in.ComPort) > 0 {
		res.ComPort = in.ComPort
	}
	// 地址码
	if in.AddressCode != 0 {
		res.AddressCode = in.AddressCode
	}

	res.UpdatedName.String = in.UpdatedName
	res.UpdatedName.Valid = true
	res.UpdatedAt.Time = time.Now()
	res.UpdatedAt.Valid = true

	err = l.svcCtx.TpmtGatewayModel.Update(l.ctx, res)

	if err != nil {
		return nil, err
	}
	return &tpmtclient.CommonResp{}, nil

}
