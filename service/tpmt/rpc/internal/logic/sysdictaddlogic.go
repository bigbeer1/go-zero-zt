package logic

import (
	"context"
	"database/sql"
	"time"
	"tpmt-zt/service/tpmt/model"

	"tpmt-zt/service/tpmt/rpc/internal/svc"
	"tpmt-zt/service/tpmt/rpc/tpmtclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type SysDictAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSysDictAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SysDictAddLogic {
	return &SysDictAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 字典
func (l *SysDictAddLogic) SysDictAdd(in *tpmtclient.SysDictAddReq) (*tpmtclient.CommonResp, error) {
	_, err := l.svcCtx.SysDictModel.Insert(l.ctx, &model.SysDict{
		CreatedAt:   time.Now(),                                                // 创建时间
		CreatedName: in.CreatedName,                                            // 创建人
		DictType:    in.DictType,                                               // 字典类型
		DictLabel:   in.DictLabel,                                              // 字典标签
		DictValue:   in.DictValue,                                              // 字典键值
		Sort:        in.Sort,                                                   // 排序
		Remark:      sql.NullString{String: in.Remark, Valid: in.Remark != ""}, // 备注
		State:       in.State,                                                  // 状态
	})
	if err != nil {
		return nil, err
	}

	return &tpmtclient.CommonResp{}, nil
}
