package logic

import (
	"context"
	"github.com/Masterminds/squirrel"

	"tpmt-zt/service/tpmt/rpc/internal/svc"
	"tpmt-zt/service/tpmt/rpc/tpmtclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type TpmtGatewayListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewTpmtGatewayListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TpmtGatewayListLogic {
	return &TpmtGatewayListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *TpmtGatewayListLogic) TpmtGatewayList(in *tpmtclient.TpmtGatewayListReq) (*tpmtclient.TpmtGatewayListResp, error) {

	whereBuilder := l.svcCtx.TpmtGatewayModel.RowBuilder()

	whereBuilder = whereBuilder.OrderBy("created_at DESC, id DESC")

	// 网关名称
	if len(in.GatewayName) > 0 {
		whereBuilder = whereBuilder.Where(squirrel.Like{
			"gateway_name ": "%" + in.GatewayName + "%",
		})
	}
	// 网关型号
	if len(in.GatewayModel) > 0 {
		whereBuilder = whereBuilder.Where(squirrel.Like{
			"gateway_model ": "%" + in.GatewayModel + "%",
		})
	}
	// 生产厂家
	if len(in.ManuFacturer) > 0 {
		whereBuilder = whereBuilder.Where(squirrel.Like{
			"manu_facturer ": "%" + in.ManuFacturer + "%",
		})
	}
	// 协议 默认1:modbus
	if in.Agreement != 99 {
		whereBuilder = whereBuilder.Where(squirrel.Eq{
			"agreement ": in.Agreement,
		})
	}
	// 波特率
	if in.BaudRate != 99 {
		whereBuilder = whereBuilder.Where(squirrel.Eq{
			"baud_rate ": in.BaudRate,
		})
	}
	// 校验
	if len(in.Parity) > 0 {
		whereBuilder = whereBuilder.Where(squirrel.Like{
			"parity ": "%" + in.Parity + "%",
		})
	}
	// 数据位
	if in.DataBits != 99 {
		whereBuilder = whereBuilder.Where(squirrel.Eq{
			"data_bits ": in.DataBits,
		})
	}
	// 停止位
	if in.StopBits != 99 {
		whereBuilder = whereBuilder.Where(squirrel.Eq{
			"stop_bits ": in.StopBits,
		})
	}
	// com端口
	if len(in.ComPort) > 0 {
		whereBuilder = whereBuilder.Where(squirrel.Like{
			"com_port ": "%" + in.ComPort + "%",
		})
	}
	// 地址码
	if in.AddressCode != 99 {
		whereBuilder = whereBuilder.Where(squirrel.Eq{
			"address_code ": in.AddressCode,
		})
	}

	all, err := l.svcCtx.TpmtGatewayModel.FindList(l.ctx, whereBuilder, in.Current, in.PageSize)
	if err != nil {
		return nil, err
	}

	countBuilder := l.svcCtx.TpmtGatewayModel.CountBuilder("id")

	// 网关名称
	if len(in.GatewayName) > 0 {
		countBuilder = countBuilder.Where(squirrel.Like{
			"gateway_name ": "%" + in.GatewayName + "%",
		})
	}
	// 网关型号
	if len(in.GatewayModel) > 0 {
		countBuilder = countBuilder.Where(squirrel.Like{
			"gateway_model ": "%" + in.GatewayModel + "%",
		})
	}
	// 生产厂家
	if len(in.ManuFacturer) > 0 {
		countBuilder = countBuilder.Where(squirrel.Like{
			"manu_facturer ": "%" + in.ManuFacturer + "%",
		})
	}
	// 协议 默认1:modbus
	if in.Agreement != 99 {
		countBuilder = countBuilder.Where(squirrel.Eq{
			"agreement ": in.Agreement,
		})
	}
	// 波特率
	if in.BaudRate != 99 {
		countBuilder = countBuilder.Where(squirrel.Eq{
			"baud_rate ": in.BaudRate,
		})
	}
	// 校验
	if len(in.Parity) > 0 {
		countBuilder = countBuilder.Where(squirrel.Like{
			"parity ": "%" + in.Parity + "%",
		})
	}
	// 数据位
	if in.DataBits != 99 {
		countBuilder = countBuilder.Where(squirrel.Eq{
			"data_bits ": in.DataBits,
		})
	}
	// 停止位
	if in.StopBits != 99 {
		countBuilder = countBuilder.Where(squirrel.Eq{
			"stop_bits ": in.StopBits,
		})
	}
	// com端口
	if len(in.ComPort) > 0 {
		countBuilder = countBuilder.Where(squirrel.Like{
			"com_port ": "%" + in.ComPort + "%",
		})
	}
	// 地址码
	if in.AddressCode != 99 {
		countBuilder = countBuilder.Where(squirrel.Eq{
			"address_code ": in.AddressCode,
		})
	}
	count, err := l.svcCtx.TpmtGatewayModel.FindCount(l.ctx, countBuilder)
	if err != nil {
		return nil, err
	}

	var list []*tpmtclient.TpmtGatewayListData
	for _, item := range all {
		list = append(list, &tpmtclient.TpmtGatewayListData{
			Id:           item.Id,                         //采集器ID/网关
			CreatedAt:    item.CreatedAt.UnixMilli(),      //创建时间
			UpdatedAt:    item.UpdatedAt.Time.UnixMilli(), //更新时间
			CreatedName:  item.CreatedName,                //创建人
			UpdatedName:  item.UpdatedName.String,         //更新人
			GatewayName:  item.GatewayName,                //网关名称
			GatewayModel: item.GatewayModel,               //网关型号
			ManuFacturer: item.ManuFacturer,               //生产厂家
			Agreement:    item.Agreement,                  //协议 默认1:modbus
			BaudRate:     item.BaudRate,                   //波特率
			Parity:       item.Parity,                     //校验
			DataBits:     item.DataBits,                   //数据位
			StopBits:     item.StopBits,                   //停止位
			ComPort:      item.ComPort,                    //com端口
			AddressCode:  item.AddressCode,                //地址码
		})
	}

	return &tpmtclient.TpmtGatewayListResp{
		Total: count,
		List:  list,
	}, nil
}
