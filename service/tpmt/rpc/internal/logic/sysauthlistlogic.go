package logic

import (
	"context"
	"github.com/Masterminds/squirrel"

	"tpmt-zt/service/tpmt/rpc/internal/svc"
	"tpmt-zt/service/tpmt/rpc/tpmtclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type SysAuthListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSysAuthListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SysAuthListLogic {
	return &SysAuthListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SysAuthListLogic) SysAuthList(in *tpmtclient.SysAuthListReq) (*tpmtclient.SysAuthListResp, error) {

	whereBuilder := l.svcCtx.SysAuthModel.RowBuilder()

	whereBuilder = whereBuilder.Where("deleted_at is null")
	whereBuilder = whereBuilder.OrderBy("created_at DESC, id DESC")

	// 机构名
	if len(in.NickName) > 0 {
		whereBuilder = whereBuilder.Where(squirrel.Like{
			"nick_name ": "%" + in.NickName + "%",
		})
	}
	// 令牌
	if len(in.AuthToken) > 0 {
		whereBuilder = whereBuilder.Where(squirrel.Like{
			"auth_token ": "%" + in.AuthToken + "%",
		})
	}
	// 状态 1:正常 2:停用 3:封禁
	if in.State != 99 {
		whereBuilder = whereBuilder.Where(squirrel.Eq{
			"state ": in.State,
		})
	}

	all, err := l.svcCtx.SysAuthModel.FindList(l.ctx, whereBuilder, in.Current, in.PageSize)
	if err != nil {
		return nil, err
	}

	countBuilder := l.svcCtx.SysAuthModel.CountBuilder("id")

	countBuilder = countBuilder.Where("deleted_at is null")

	// 机构名
	if len(in.NickName) > 0 {
		countBuilder = countBuilder.Where(squirrel.Like{
			"nick_name ": "%" + in.NickName + "%",
		})
	}
	// 令牌
	if len(in.AuthToken) > 0 {
		countBuilder = countBuilder.Where(squirrel.Like{
			"auth_token ": "%" + in.AuthToken + "%",
		})
	}
	// 状态 1:正常 2:停用 3:封禁
	if in.State != 99 {
		countBuilder = countBuilder.Where(squirrel.Eq{
			"state ": in.State,
		})
	}
	count, err := l.svcCtx.SysAuthModel.FindCount(l.ctx, countBuilder)
	if err != nil {
		return nil, err
	}

	var list []*tpmtclient.SysAuthListData
	for _, item := range all {
		list = append(list, &tpmtclient.SysAuthListData{
			Id:          item.Id,                         //第三方用户ID
			CreatedAt:   item.CreatedAt.UnixMilli(),      //创建时间
			UpdatedAt:   item.UpdatedAt.Time.UnixMilli(), //更新时间
			CreatedName: item.CreatedName,                //创建人
			UpdatedName: item.UpdatedName.String,         //更新人
			NickName:    item.NickName,                   //机构名
			AuthToken:   item.AuthToken,                  //令牌
			State:       item.State,                      //状态 1:正常 2:停用 3:封禁
		})
	}

	return &tpmtclient.SysAuthListResp{
		Total: count,
		List:  list,
	}, nil
}
