package logic

import (
	"context"
	"errors"
	"fmt"
	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"tpmt-zt/service/tpmt/model"

	"tpmt-zt/service/tpmt/rpc/internal/svc"
	"tpmt-zt/service/tpmt/rpc/tpmtclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type SysUserListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSysUserListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SysUserListLogic {
	return &SysUserListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SysUserListLogic) SysUserList(in *tpmtclient.SysUserListReq) (*tpmtclient.SysUserListResp, error) {
	whereBuilder := l.svcCtx.SysUserModel.RowBuilder()

	whereBuilder = whereBuilder.Where("deleted_at is null")
	whereBuilder = whereBuilder.OrderBy("created_at DESC, id DESC")

	// 姓名
	if len(in.NickName) > 0 {
		whereBuilder = whereBuilder.Where(squirrel.Like{
			"nick_name ": "%" + in.NickName + "%",
		})
	}

	// 状态 1:正常 2:停用 3:封禁
	if in.State != 99 {
		whereBuilder = whereBuilder.Where(squirrel.Eq{
			"state ": in.State,
		})
	}

	all, err := l.svcCtx.SysUserModel.FindList(l.ctx, whereBuilder, in.Current, in.PageSize)
	if err != nil {
		return nil, err
	}

	countBuilder := l.svcCtx.SysUserModel.CountBuilder("id")

	countBuilder = countBuilder.Where("deleted_at is null")

	// 姓名
	if len(in.NickName) > 0 {
		countBuilder = countBuilder.Where(squirrel.Like{
			"nick_name ": "%" + in.NickName + "%",
		})
	}

	// 状态 1:正常 2:停用 3:封禁
	if in.State != 99 {
		countBuilder = countBuilder.Where(squirrel.Eq{
			"state ": in.State,
		})
	}
	count, err := l.svcCtx.SysUserModel.FindCount(l.ctx, countBuilder)
	if err != nil {
		return nil, err
	}

	var list []*tpmtclient.SysUserListData
	for _, item := range all {
		// 查询用户角色信息
		userRole, err := l.svcCtx.SysUserRoleModel.FindByUserId(l.ctx, item.Id)
		if err != nil {
			if !errors.Is(err, sqlc.ErrNotFound) {
				return nil, err
			}
		}

		var roleRes model.SysRole

		if err == nil {
			role, err := l.svcCtx.SysRoleModel.FindOne(l.ctx, userRole.RoleId)
			if err != nil {
				if errors.Is(err, sqlc.ErrNotFound) {
					return nil, fmt.Errorf("SysRole没有该ID:%v", role.Id)
				}
				return nil, err
			}

			// 判断该数据是否被删除
			if role.DeletedAt.Valid == true {
				return nil, fmt.Errorf("SysRole该ID已被删除：%v", role.Id)
			}
			roleRes = model.SysRole{
				Id:          role.Id,
				Name:        role.Name,
				Remark:      role.Remark,
				RoleType:    role.RoleType,
				CreatedName: role.CreatedName,
				CreatedAt:   role.CreatedAt,
				UpdatedName: role.UpdatedName,
				UpdatedAt:   role.UpdatedAt,
				DeletedAt:   role.DeletedAt,
				DeletedName: role.DeletedName,
			}
		}

		list = append(list, &tpmtclient.SysUserListData{
			Id:          item.Id,                         //用户ID
			Account:     item.Account,                    //用户名
			NickName:    item.NickName,                   //姓名
			State:       item.State,                      //状态 1:正常 2:停用 3:封禁
			CreatedName: item.CreatedName,                //创建人
			CreatedAt:   item.CreatedAt.UnixMilli(),      //创建时间
			UpdatedName: item.UpdatedName.String,         //更新人
			UpdatedAt:   item.UpdatedAt.Time.UnixMilli(), //更新时间
			RoleId:      roleRes.Id,                      //角色id
			RoleName:    roleRes.Name,                    //角色名称
			RoleType:    roleRes.RoleType,                //角色类型
		})
	}

	return &tpmtclient.SysUserListResp{
		Total: count,
		List:  list,
	}, nil
}
