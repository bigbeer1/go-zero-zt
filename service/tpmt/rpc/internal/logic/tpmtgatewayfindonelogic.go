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

type TpmtGatewayFindOneLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewTpmtGatewayFindOneLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TpmtGatewayFindOneLogic {
	return &TpmtGatewayFindOneLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *TpmtGatewayFindOneLogic) TpmtGatewayFindOne(in *tpmtclient.TpmtGatewayFindOneReq) (*tpmtclient.TpmtGatewayFindOneResp, error) {

	res, err := l.svcCtx.TpmtGatewayModel.FindOne(l.ctx, in.Id)
	if err != nil {
		if errors.Is(err, sqlc.ErrNotFound) {
			return nil, fmt.Errorf("TpmtGateway没有该ID:%v", in.Id)
		}
		return nil, err
	}

	return &tpmtclient.TpmtGatewayFindOneResp{
		Id:           res.Id,                         //采集器ID/网关
		CreatedAt:    res.CreatedAt.UnixMilli(),      //创建时间
		UpdatedAt:    res.UpdatedAt.Time.UnixMilli(), //更新时间
		CreatedName:  res.CreatedName,                //创建人
		UpdatedName:  res.UpdatedName.String,         //更新人
		GatewayName:  res.GatewayName,                //网关名称
		GatewayModel: res.GatewayModel,               //网关型号
		ManuFacturer: res.ManuFacturer,               //生产厂家
		Agreement:    res.Agreement,                  //协议 默认1:modbus
		BaudRate:     res.BaudRate,                   //波特率
		Parity:       res.Parity,                     //校验
		DataBits:     res.DataBits,                   //数据位
		StopBits:     res.StopBits,                   //停止位
		ComPort:      res.ComPort,                    //com端口
		AddressCode:  res.AddressCode,                //地址码
	}, nil
}
