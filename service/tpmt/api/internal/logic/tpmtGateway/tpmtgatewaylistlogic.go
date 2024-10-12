package tpmtGateway

import (
	"context"
	"github.com/jinzhu/copier"
	"tpmt-zt/common"
	"tpmt-zt/common/msg"
	"tpmt-zt/service/tpmt/rpc/tpmtclient"

	"tpmt-zt/service/tpmt/api/internal/svc"
	"tpmt-zt/service/tpmt/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type TpmtGatewayListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTpmtGatewayListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TpmtGatewayListLogic {
	return &TpmtGatewayListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TpmtGatewayListLogic) TpmtGatewayList(req *types.TpmtGatewayListRequest) (resp *types.Response, err error) {

	all, err := l.svcCtx.TpmtRpc.TpmtGatewayList(l.ctx, &tpmtclient.TpmtGatewayListReq{
		Current:      req.Current,      // 页码
		PageSize:     req.PageSize,     // 页数
		GatewayName:  req.GatewayName,  // 网关名称
		GatewayModel: req.GatewayModel, // 网关型号
		ManuFacturer: req.ManuFacturer, // 生产厂家
		Agreement:    req.Agreement,    // 协议 默认1:modbus
		BaudRate:     req.BaudRate,     // 波特率
		Parity:       req.Parity,       // 校验
		DataBits:     req.DataBits,     // 数据位
		StopBits:     req.StopBits,     // 停止位
		ComPort:      req.ComPort,      // com端口
		AddressCode:  req.AddressCode,  // 地址码
	})
	if err != nil {
		return nil, common.NewDefaultError(err.Error())
	}

	var result TpmtGatewayListResp
	_ = copier.Copy(&result, all)

	return &types.Response{
		Code: 0,
		Msg:  msg.Success,
		Data: result,
	}, nil
}

type TpmtGatewayListResp struct {
	Total int64                  `json:"total"`
	List  []*TpmtGatewayDataList `json:"list"`
}

type TpmtGatewayDataList struct {
	Id           string `json:"id"`            // 采集器ID/网关,
	CreatedAt    int64  `json:"created_at"`    // 创建时间,
	UpdatedAt    int64  `json:"updated_at"`    // 更新时间,
	CreatedName  string `json:"created_name"`  // 创建人,
	UpdatedName  string `json:"updated_name"`  // 更新人,
	GatewayName  string `json:"gateway_name"`  // 网关名称,
	GatewayModel string `json:"gateway_model"` // 网关型号,
	ManuFacturer string `json:"manu_facturer"` // 生产厂家,
	Agreement    int64  `json:"agreement"`     // 协议 默认1:modbus,
	BaudRate     int64  `json:"baud_rate"`     // 波特率,
	Parity       string `json:"parity"`        // 校验,
	DataBits     int64  `json:"data_bits"`     // 数据位,
	StopBits     int64  `json:"stop_bits"`     // 停止位,
	ComPort      string `json:"com_port"`      // com端口,
	AddressCode  int64  `json:"address_code"`  // 地址码
}
