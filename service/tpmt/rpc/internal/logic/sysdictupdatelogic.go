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

type SysDictUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSysDictUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SysDictUpdateLogic {
	return &SysDictUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SysDictUpdateLogic) SysDictUpdate(in *tpmtclient.SysDictUpdateReq) (*tpmtclient.CommonResp, error) {

	res, err := l.svcCtx.SysDictModel.FindOne(l.ctx, in.Id)
	if err != nil {
		if errors.Is(err, sqlc.ErrNotFound) {
			return nil, fmt.Errorf("SysDict没有该ID: %v", in.Id)
		}
		return nil, err
	}

	// 判断该数据是否被删除
	if res.DeletedAt.Valid == true {
		return nil, fmt.Errorf("SysDict该ID已被删除： %v", in.Id)
	}

	// 字典类型
	if len(in.DictType) > 0 {
		res.DictType = in.DictType
	}
	// 字典标签
	if len(in.DictLabel) > 0 {
		res.DictLabel = in.DictLabel
	}
	// 字典键值
	if len(in.DictValue) > 0 {
		res.DictValue = in.DictValue
	}
	// 排序
	if in.Sort != 0 {
		res.Sort = in.Sort
	}
	// 备注
	if len(in.Remark) > 0 {
		res.Remark.String = in.Remark
		res.Remark.Valid = true
	}
	// 状态
	if in.State != 0 {
		res.State = in.State
	}

	res.UpdatedName.String = in.UpdatedName
	res.UpdatedName.Valid = true
	res.UpdatedAt.Time = time.Now()
	res.UpdatedAt.Valid = true

	err = l.svcCtx.SysDictModel.Update(l.ctx, res)

	if err != nil {
		return nil, err
	}
	return &tpmtclient.CommonResp{}, nil

}
