package sysDict

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

type SysDictAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSysDictAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SysDictAddLogic {
	return &SysDictAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SysDictAddLogic) SysDictAdd(req *types.SysDictAddRequest) (resp *types.Response, err error) {
	// 用户登录信息
	tokenData := jwtx.ParseToken(l.ctx)

	_, err = l.svcCtx.TpmtRpc.SysDictAdd(l.ctx, &tpmtclient.SysDictAddReq{
		CreatedName: tokenData.NickName, // 创建人
		DictType:    req.DictType,       // 字典类型
		DictLabel:   req.DictLabel,      // 字典标签
		DictValue:   req.DictValue,      // 字典键值
		Sort:        req.Sort,           // 排序
		Remark:      req.Remark,         // 备注
		State:       req.State,          // 状态
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
