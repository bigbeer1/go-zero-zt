// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.2
// Source: tpmt.proto

package tpmt

import (
	"context"

	"tpmt-zt/service/tpmt/rpc/tpmtclient"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	AlarmRuleInfo               = tpmtclient.AlarmRuleInfo
	CommonResp                  = tpmtclient.CommonResp
	SysAuthAddReq               = tpmtclient.SysAuthAddReq
	SysAuthDeleteReq            = tpmtclient.SysAuthDeleteReq
	SysAuthFindOneReq           = tpmtclient.SysAuthFindOneReq
	SysAuthFindOneResp          = tpmtclient.SysAuthFindOneResp
	SysAuthListData             = tpmtclient.SysAuthListData
	SysAuthListReq              = tpmtclient.SysAuthListReq
	SysAuthListResp             = tpmtclient.SysAuthListResp
	SysAuthUpdateReq            = tpmtclient.SysAuthUpdateReq
	SysDictAddReq               = tpmtclient.SysDictAddReq
	SysDictDeleteReq            = tpmtclient.SysDictDeleteReq
	SysDictFindOneReq           = tpmtclient.SysDictFindOneReq
	SysDictFindOneResp          = tpmtclient.SysDictFindOneResp
	SysDictListData             = tpmtclient.SysDictListData
	SysDictListReq              = tpmtclient.SysDictListReq
	SysDictListResp             = tpmtclient.SysDictListResp
	SysDictTypeAddReq           = tpmtclient.SysDictTypeAddReq
	SysDictTypeDeleteReq        = tpmtclient.SysDictTypeDeleteReq
	SysDictTypeFindOneReq       = tpmtclient.SysDictTypeFindOneReq
	SysDictTypeFindOneResp      = tpmtclient.SysDictTypeFindOneResp
	SysDictTypeListData         = tpmtclient.SysDictTypeListData
	SysDictTypeListReq          = tpmtclient.SysDictTypeListReq
	SysDictTypeListResp         = tpmtclient.SysDictTypeListResp
	SysDictTypeUpdateReq        = tpmtclient.SysDictTypeUpdateReq
	SysDictUpdateReq            = tpmtclient.SysDictUpdateReq
	SysInterfaceAddReq          = tpmtclient.SysInterfaceAddReq
	SysInterfaceByRoleIdReq     = tpmtclient.SysInterfaceByRoleIdReq
	SysInterfaceByRoleIdResp    = tpmtclient.SysInterfaceByRoleIdResp
	SysInterfaceDeleteReq       = tpmtclient.SysInterfaceDeleteReq
	SysInterfaceFindOneReq      = tpmtclient.SysInterfaceFindOneReq
	SysInterfaceFindOneResp     = tpmtclient.SysInterfaceFindOneResp
	SysInterfaceListData        = tpmtclient.SysInterfaceListData
	SysInterfaceListReq         = tpmtclient.SysInterfaceListReq
	SysInterfaceListResp        = tpmtclient.SysInterfaceListResp
	SysInterfaceUpdateReq       = tpmtclient.SysInterfaceUpdateReq
	SysLoginReq                 = tpmtclient.SysLoginReq
	SysMenuAddReq               = tpmtclient.SysMenuAddReq
	SysMenuByRoleIdReq          = tpmtclient.SysMenuByRoleIdReq
	SysMenuByRoleIdResp         = tpmtclient.SysMenuByRoleIdResp
	SysMenuDeleteReq            = tpmtclient.SysMenuDeleteReq
	SysMenuFindOneReq           = tpmtclient.SysMenuFindOneReq
	SysMenuFindOneResp          = tpmtclient.SysMenuFindOneResp
	SysMenuListData             = tpmtclient.SysMenuListData
	SysMenuListReq              = tpmtclient.SysMenuListReq
	SysMenuListResp             = tpmtclient.SysMenuListResp
	SysMenuUpdateReq            = tpmtclient.SysMenuUpdateReq
	SysRoleAddReq               = tpmtclient.SysRoleAddReq
	SysRoleDeleteReq            = tpmtclient.SysRoleDeleteReq
	SysRoleFindOneReq           = tpmtclient.SysRoleFindOneReq
	SysRoleFindOneResp          = tpmtclient.SysRoleFindOneResp
	SysRoleListData             = tpmtclient.SysRoleListData
	SysRoleListReq              = tpmtclient.SysRoleListReq
	SysRoleListResp             = tpmtclient.SysRoleListResp
	SysRoleUpdateReq            = tpmtclient.SysRoleUpdateReq
	SysUserAddReq               = tpmtclient.SysUserAddReq
	SysUserDeleteReq            = tpmtclient.SysUserDeleteReq
	SysUserFindOneReq           = tpmtclient.SysUserFindOneReq
	SysUserFindOneResp          = tpmtclient.SysUserFindOneResp
	SysUserListData             = tpmtclient.SysUserListData
	SysUserListReq              = tpmtclient.SysUserListReq
	SysUserListResp             = tpmtclient.SysUserListResp
	SysUserResetPwdReq          = tpmtclient.SysUserResetPwdReq
	SysUserResetPwdResp         = tpmtclient.SysUserResetPwdResp
	SysUserUpMyPwdReq           = tpmtclient.SysUserUpMyPwdReq
	SysUserUpdateReq            = tpmtclient.SysUserUpdateReq
	TpmtAssetAddReq             = tpmtclient.TpmtAssetAddReq
	TpmtAssetDeleteReq          = tpmtclient.TpmtAssetDeleteReq
	TpmtAssetFindOneReq         = tpmtclient.TpmtAssetFindOneReq
	TpmtAssetFindOneResp        = tpmtclient.TpmtAssetFindOneResp
	TpmtAssetListData           = tpmtclient.TpmtAssetListData
	TpmtAssetListReq            = tpmtclient.TpmtAssetListReq
	TpmtAssetListResp           = tpmtclient.TpmtAssetListResp
	TpmtAssetUpdateReq          = tpmtclient.TpmtAssetUpdateReq
	TpmtGatewayAddReq           = tpmtclient.TpmtGatewayAddReq
	TpmtGatewayDeleteReq        = tpmtclient.TpmtGatewayDeleteReq
	TpmtGatewayFindOneReq       = tpmtclient.TpmtGatewayFindOneReq
	TpmtGatewayFindOneResp      = tpmtclient.TpmtGatewayFindOneResp
	TpmtGatewayListData         = tpmtclient.TpmtGatewayListData
	TpmtGatewayListReq          = tpmtclient.TpmtGatewayListReq
	TpmtGatewayListResp         = tpmtclient.TpmtGatewayListResp
	TpmtGatewayUpdateReq        = tpmtclient.TpmtGatewayUpdateReq
	TpmtMonitorPointAddReq      = tpmtclient.TpmtMonitorPointAddReq
	TpmtMonitorPointDeleteReq   = tpmtclient.TpmtMonitorPointDeleteReq
	TpmtMonitorPointFindOneReq  = tpmtclient.TpmtMonitorPointFindOneReq
	TpmtMonitorPointFindOneResp = tpmtclient.TpmtMonitorPointFindOneResp
	TpmtMonitorPointListData    = tpmtclient.TpmtMonitorPointListData
	TpmtMonitorPointListReq     = tpmtclient.TpmtMonitorPointListReq
	TpmtMonitorPointListResp    = tpmtclient.TpmtMonitorPointListResp
	TpmtMonitorPointUpdateReq   = tpmtclient.TpmtMonitorPointUpdateReq

	Tpmt interface {
		// 用户登录
		SysLogin(ctx context.Context, in *SysLoginReq, opts ...grpc.CallOption) (*SysUserFindOneResp, error)
		// 用户
		SysUserAdd(ctx context.Context, in *SysUserAddReq, opts ...grpc.CallOption) (*CommonResp, error)
		SysUserDelete(ctx context.Context, in *SysUserDeleteReq, opts ...grpc.CallOption) (*CommonResp, error)
		SysUserUpdate(ctx context.Context, in *SysUserUpdateReq, opts ...grpc.CallOption) (*CommonResp, error)
		SysUserFindOne(ctx context.Context, in *SysUserFindOneReq, opts ...grpc.CallOption) (*SysUserFindOneResp, error)
		SysUserList(ctx context.Context, in *SysUserListReq, opts ...grpc.CallOption) (*SysUserListResp, error)
		// 重置用户密码
		SysUserResetPwd(ctx context.Context, in *SysUserResetPwdReq, opts ...grpc.CallOption) (*SysUserResetPwdResp, error)
		// 用户修改自己的密码
		SysUserUpMyPwd(ctx context.Context, in *SysUserUpMyPwdReq, opts ...grpc.CallOption) (*CommonResp, error)
		// 第三方用户
		SysAuthAdd(ctx context.Context, in *SysAuthAddReq, opts ...grpc.CallOption) (*CommonResp, error)
		SysAuthDelete(ctx context.Context, in *SysAuthDeleteReq, opts ...grpc.CallOption) (*CommonResp, error)
		SysAuthUpdate(ctx context.Context, in *SysAuthUpdateReq, opts ...grpc.CallOption) (*CommonResp, error)
		SysAuthFindOne(ctx context.Context, in *SysAuthFindOneReq, opts ...grpc.CallOption) (*SysAuthFindOneResp, error)
		SysAuthList(ctx context.Context, in *SysAuthListReq, opts ...grpc.CallOption) (*SysAuthListResp, error)
		// 角色
		SysRoleAdd(ctx context.Context, in *SysRoleAddReq, opts ...grpc.CallOption) (*CommonResp, error)
		SysRoleDelete(ctx context.Context, in *SysRoleDeleteReq, opts ...grpc.CallOption) (*CommonResp, error)
		SysRoleUpdate(ctx context.Context, in *SysRoleUpdateReq, opts ...grpc.CallOption) (*CommonResp, error)
		SysRoleFindOne(ctx context.Context, in *SysRoleFindOneReq, opts ...grpc.CallOption) (*SysRoleFindOneResp, error)
		SysRoleList(ctx context.Context, in *SysRoleListReq, opts ...grpc.CallOption) (*SysRoleListResp, error)
		// 菜单
		SysMenuAdd(ctx context.Context, in *SysMenuAddReq, opts ...grpc.CallOption) (*CommonResp, error)
		SysMenuDelete(ctx context.Context, in *SysMenuDeleteReq, opts ...grpc.CallOption) (*CommonResp, error)
		SysMenuUpdate(ctx context.Context, in *SysMenuUpdateReq, opts ...grpc.CallOption) (*CommonResp, error)
		SysMenuFindOne(ctx context.Context, in *SysMenuFindOneReq, opts ...grpc.CallOption) (*SysMenuFindOneResp, error)
		SysMenuList(ctx context.Context, in *SysMenuListReq, opts ...grpc.CallOption) (*SysMenuListResp, error)
		// 通过角色ID获取菜单信息
		SysMenuByRoleId(ctx context.Context, in *SysMenuByRoleIdReq, opts ...grpc.CallOption) (*SysMenuByRoleIdResp, error)
		// 接口
		SysInterfaceAdd(ctx context.Context, in *SysInterfaceAddReq, opts ...grpc.CallOption) (*CommonResp, error)
		SysInterfaceDelete(ctx context.Context, in *SysInterfaceDeleteReq, opts ...grpc.CallOption) (*CommonResp, error)
		SysInterfaceUpdate(ctx context.Context, in *SysInterfaceUpdateReq, opts ...grpc.CallOption) (*CommonResp, error)
		SysInterfaceFindOne(ctx context.Context, in *SysInterfaceFindOneReq, opts ...grpc.CallOption) (*SysInterfaceFindOneResp, error)
		SysInterfaceList(ctx context.Context, in *SysInterfaceListReq, opts ...grpc.CallOption) (*SysInterfaceListResp, error)
		// 通过角色ID获取接口信息
		SysInterfaceByRoleId(ctx context.Context, in *SysInterfaceByRoleIdReq, opts ...grpc.CallOption) (*SysInterfaceByRoleIdResp, error)
		// 字典类型
		SysDictTypeAdd(ctx context.Context, in *SysDictTypeAddReq, opts ...grpc.CallOption) (*CommonResp, error)
		SysDictTypeDelete(ctx context.Context, in *SysDictTypeDeleteReq, opts ...grpc.CallOption) (*CommonResp, error)
		SysDictTypeUpdate(ctx context.Context, in *SysDictTypeUpdateReq, opts ...grpc.CallOption) (*CommonResp, error)
		SysDictTypeFindOne(ctx context.Context, in *SysDictTypeFindOneReq, opts ...grpc.CallOption) (*SysDictTypeFindOneResp, error)
		SysDictTypeList(ctx context.Context, in *SysDictTypeListReq, opts ...grpc.CallOption) (*SysDictTypeListResp, error)
		// 字典
		SysDictAdd(ctx context.Context, in *SysDictAddReq, opts ...grpc.CallOption) (*CommonResp, error)
		SysDictDelete(ctx context.Context, in *SysDictDeleteReq, opts ...grpc.CallOption) (*CommonResp, error)
		SysDictUpdate(ctx context.Context, in *SysDictUpdateReq, opts ...grpc.CallOption) (*CommonResp, error)
		SysDictFindOne(ctx context.Context, in *SysDictFindOneReq, opts ...grpc.CallOption) (*SysDictFindOneResp, error)
		SysDictList(ctx context.Context, in *SysDictListReq, opts ...grpc.CallOption) (*SysDictListResp, error)
		// 资产
		TpmtAssetAdd(ctx context.Context, in *TpmtAssetAddReq, opts ...grpc.CallOption) (*CommonResp, error)
		TpmtAssetDelete(ctx context.Context, in *TpmtAssetDeleteReq, opts ...grpc.CallOption) (*CommonResp, error)
		TpmtAssetUpdate(ctx context.Context, in *TpmtAssetUpdateReq, opts ...grpc.CallOption) (*CommonResp, error)
		TpmtAssetFindOne(ctx context.Context, in *TpmtAssetFindOneReq, opts ...grpc.CallOption) (*TpmtAssetFindOneResp, error)
		TpmtAssetList(ctx context.Context, in *TpmtAssetListReq, opts ...grpc.CallOption) (*TpmtAssetListResp, error)
		// 网关
		TpmtGatewayAdd(ctx context.Context, in *TpmtGatewayAddReq, opts ...grpc.CallOption) (*CommonResp, error)
		TpmtGatewayDelete(ctx context.Context, in *TpmtGatewayDeleteReq, opts ...grpc.CallOption) (*CommonResp, error)
		TpmtGatewayUpdate(ctx context.Context, in *TpmtGatewayUpdateReq, opts ...grpc.CallOption) (*CommonResp, error)
		TpmtGatewayFindOne(ctx context.Context, in *TpmtGatewayFindOneReq, opts ...grpc.CallOption) (*TpmtGatewayFindOneResp, error)
		TpmtGatewayList(ctx context.Context, in *TpmtGatewayListReq, opts ...grpc.CallOption) (*TpmtGatewayListResp, error)
		// 监控点
		TpmtMonitorPointAdd(ctx context.Context, in *TpmtMonitorPointAddReq, opts ...grpc.CallOption) (*CommonResp, error)
		TpmtMonitorPointDelete(ctx context.Context, in *TpmtMonitorPointDeleteReq, opts ...grpc.CallOption) (*CommonResp, error)
		TpmtMonitorPointUpdate(ctx context.Context, in *TpmtMonitorPointUpdateReq, opts ...grpc.CallOption) (*CommonResp, error)
		TpmtMonitorPointFindOne(ctx context.Context, in *TpmtMonitorPointFindOneReq, opts ...grpc.CallOption) (*TpmtMonitorPointFindOneResp, error)
		TpmtMonitorPointList(ctx context.Context, in *TpmtMonitorPointListReq, opts ...grpc.CallOption) (*TpmtMonitorPointListResp, error)
	}

	defaultTpmt struct {
		cli zrpc.Client
	}
)

func NewTpmt(cli zrpc.Client) Tpmt {
	return &defaultTpmt{
		cli: cli,
	}
}

// 用户登录
func (m *defaultTpmt) SysLogin(ctx context.Context, in *SysLoginReq, opts ...grpc.CallOption) (*SysUserFindOneResp, error) {
	client := tpmtclient.NewTpmtClient(m.cli.Conn())
	return client.SysLogin(ctx, in, opts...)
}

// 用户
func (m *defaultTpmt) SysUserAdd(ctx context.Context, in *SysUserAddReq, opts ...grpc.CallOption) (*CommonResp, error) {
	client := tpmtclient.NewTpmtClient(m.cli.Conn())
	return client.SysUserAdd(ctx, in, opts...)
}

func (m *defaultTpmt) SysUserDelete(ctx context.Context, in *SysUserDeleteReq, opts ...grpc.CallOption) (*CommonResp, error) {
	client := tpmtclient.NewTpmtClient(m.cli.Conn())
	return client.SysUserDelete(ctx, in, opts...)
}

func (m *defaultTpmt) SysUserUpdate(ctx context.Context, in *SysUserUpdateReq, opts ...grpc.CallOption) (*CommonResp, error) {
	client := tpmtclient.NewTpmtClient(m.cli.Conn())
	return client.SysUserUpdate(ctx, in, opts...)
}

func (m *defaultTpmt) SysUserFindOne(ctx context.Context, in *SysUserFindOneReq, opts ...grpc.CallOption) (*SysUserFindOneResp, error) {
	client := tpmtclient.NewTpmtClient(m.cli.Conn())
	return client.SysUserFindOne(ctx, in, opts...)
}

func (m *defaultTpmt) SysUserList(ctx context.Context, in *SysUserListReq, opts ...grpc.CallOption) (*SysUserListResp, error) {
	client := tpmtclient.NewTpmtClient(m.cli.Conn())
	return client.SysUserList(ctx, in, opts...)
}

// 重置用户密码
func (m *defaultTpmt) SysUserResetPwd(ctx context.Context, in *SysUserResetPwdReq, opts ...grpc.CallOption) (*SysUserResetPwdResp, error) {
	client := tpmtclient.NewTpmtClient(m.cli.Conn())
	return client.SysUserResetPwd(ctx, in, opts...)
}

// 用户修改自己的密码
func (m *defaultTpmt) SysUserUpMyPwd(ctx context.Context, in *SysUserUpMyPwdReq, opts ...grpc.CallOption) (*CommonResp, error) {
	client := tpmtclient.NewTpmtClient(m.cli.Conn())
	return client.SysUserUpMyPwd(ctx, in, opts...)
}

// 第三方用户
func (m *defaultTpmt) SysAuthAdd(ctx context.Context, in *SysAuthAddReq, opts ...grpc.CallOption) (*CommonResp, error) {
	client := tpmtclient.NewTpmtClient(m.cli.Conn())
	return client.SysAuthAdd(ctx, in, opts...)
}

func (m *defaultTpmt) SysAuthDelete(ctx context.Context, in *SysAuthDeleteReq, opts ...grpc.CallOption) (*CommonResp, error) {
	client := tpmtclient.NewTpmtClient(m.cli.Conn())
	return client.SysAuthDelete(ctx, in, opts...)
}

func (m *defaultTpmt) SysAuthUpdate(ctx context.Context, in *SysAuthUpdateReq, opts ...grpc.CallOption) (*CommonResp, error) {
	client := tpmtclient.NewTpmtClient(m.cli.Conn())
	return client.SysAuthUpdate(ctx, in, opts...)
}

func (m *defaultTpmt) SysAuthFindOne(ctx context.Context, in *SysAuthFindOneReq, opts ...grpc.CallOption) (*SysAuthFindOneResp, error) {
	client := tpmtclient.NewTpmtClient(m.cli.Conn())
	return client.SysAuthFindOne(ctx, in, opts...)
}

func (m *defaultTpmt) SysAuthList(ctx context.Context, in *SysAuthListReq, opts ...grpc.CallOption) (*SysAuthListResp, error) {
	client := tpmtclient.NewTpmtClient(m.cli.Conn())
	return client.SysAuthList(ctx, in, opts...)
}

// 角色
func (m *defaultTpmt) SysRoleAdd(ctx context.Context, in *SysRoleAddReq, opts ...grpc.CallOption) (*CommonResp, error) {
	client := tpmtclient.NewTpmtClient(m.cli.Conn())
	return client.SysRoleAdd(ctx, in, opts...)
}

func (m *defaultTpmt) SysRoleDelete(ctx context.Context, in *SysRoleDeleteReq, opts ...grpc.CallOption) (*CommonResp, error) {
	client := tpmtclient.NewTpmtClient(m.cli.Conn())
	return client.SysRoleDelete(ctx, in, opts...)
}

func (m *defaultTpmt) SysRoleUpdate(ctx context.Context, in *SysRoleUpdateReq, opts ...grpc.CallOption) (*CommonResp, error) {
	client := tpmtclient.NewTpmtClient(m.cli.Conn())
	return client.SysRoleUpdate(ctx, in, opts...)
}

func (m *defaultTpmt) SysRoleFindOne(ctx context.Context, in *SysRoleFindOneReq, opts ...grpc.CallOption) (*SysRoleFindOneResp, error) {
	client := tpmtclient.NewTpmtClient(m.cli.Conn())
	return client.SysRoleFindOne(ctx, in, opts...)
}

func (m *defaultTpmt) SysRoleList(ctx context.Context, in *SysRoleListReq, opts ...grpc.CallOption) (*SysRoleListResp, error) {
	client := tpmtclient.NewTpmtClient(m.cli.Conn())
	return client.SysRoleList(ctx, in, opts...)
}

// 菜单
func (m *defaultTpmt) SysMenuAdd(ctx context.Context, in *SysMenuAddReq, opts ...grpc.CallOption) (*CommonResp, error) {
	client := tpmtclient.NewTpmtClient(m.cli.Conn())
	return client.SysMenuAdd(ctx, in, opts...)
}

func (m *defaultTpmt) SysMenuDelete(ctx context.Context, in *SysMenuDeleteReq, opts ...grpc.CallOption) (*CommonResp, error) {
	client := tpmtclient.NewTpmtClient(m.cli.Conn())
	return client.SysMenuDelete(ctx, in, opts...)
}

func (m *defaultTpmt) SysMenuUpdate(ctx context.Context, in *SysMenuUpdateReq, opts ...grpc.CallOption) (*CommonResp, error) {
	client := tpmtclient.NewTpmtClient(m.cli.Conn())
	return client.SysMenuUpdate(ctx, in, opts...)
}

func (m *defaultTpmt) SysMenuFindOne(ctx context.Context, in *SysMenuFindOneReq, opts ...grpc.CallOption) (*SysMenuFindOneResp, error) {
	client := tpmtclient.NewTpmtClient(m.cli.Conn())
	return client.SysMenuFindOne(ctx, in, opts...)
}

func (m *defaultTpmt) SysMenuList(ctx context.Context, in *SysMenuListReq, opts ...grpc.CallOption) (*SysMenuListResp, error) {
	client := tpmtclient.NewTpmtClient(m.cli.Conn())
	return client.SysMenuList(ctx, in, opts...)
}

// 通过角色ID获取菜单信息
func (m *defaultTpmt) SysMenuByRoleId(ctx context.Context, in *SysMenuByRoleIdReq, opts ...grpc.CallOption) (*SysMenuByRoleIdResp, error) {
	client := tpmtclient.NewTpmtClient(m.cli.Conn())
	return client.SysMenuByRoleId(ctx, in, opts...)
}

// 接口
func (m *defaultTpmt) SysInterfaceAdd(ctx context.Context, in *SysInterfaceAddReq, opts ...grpc.CallOption) (*CommonResp, error) {
	client := tpmtclient.NewTpmtClient(m.cli.Conn())
	return client.SysInterfaceAdd(ctx, in, opts...)
}

func (m *defaultTpmt) SysInterfaceDelete(ctx context.Context, in *SysInterfaceDeleteReq, opts ...grpc.CallOption) (*CommonResp, error) {
	client := tpmtclient.NewTpmtClient(m.cli.Conn())
	return client.SysInterfaceDelete(ctx, in, opts...)
}

func (m *defaultTpmt) SysInterfaceUpdate(ctx context.Context, in *SysInterfaceUpdateReq, opts ...grpc.CallOption) (*CommonResp, error) {
	client := tpmtclient.NewTpmtClient(m.cli.Conn())
	return client.SysInterfaceUpdate(ctx, in, opts...)
}

func (m *defaultTpmt) SysInterfaceFindOne(ctx context.Context, in *SysInterfaceFindOneReq, opts ...grpc.CallOption) (*SysInterfaceFindOneResp, error) {
	client := tpmtclient.NewTpmtClient(m.cli.Conn())
	return client.SysInterfaceFindOne(ctx, in, opts...)
}

func (m *defaultTpmt) SysInterfaceList(ctx context.Context, in *SysInterfaceListReq, opts ...grpc.CallOption) (*SysInterfaceListResp, error) {
	client := tpmtclient.NewTpmtClient(m.cli.Conn())
	return client.SysInterfaceList(ctx, in, opts...)
}

// 通过角色ID获取接口信息
func (m *defaultTpmt) SysInterfaceByRoleId(ctx context.Context, in *SysInterfaceByRoleIdReq, opts ...grpc.CallOption) (*SysInterfaceByRoleIdResp, error) {
	client := tpmtclient.NewTpmtClient(m.cli.Conn())
	return client.SysInterfaceByRoleId(ctx, in, opts...)
}

// 字典类型
func (m *defaultTpmt) SysDictTypeAdd(ctx context.Context, in *SysDictTypeAddReq, opts ...grpc.CallOption) (*CommonResp, error) {
	client := tpmtclient.NewTpmtClient(m.cli.Conn())
	return client.SysDictTypeAdd(ctx, in, opts...)
}

func (m *defaultTpmt) SysDictTypeDelete(ctx context.Context, in *SysDictTypeDeleteReq, opts ...grpc.CallOption) (*CommonResp, error) {
	client := tpmtclient.NewTpmtClient(m.cli.Conn())
	return client.SysDictTypeDelete(ctx, in, opts...)
}

func (m *defaultTpmt) SysDictTypeUpdate(ctx context.Context, in *SysDictTypeUpdateReq, opts ...grpc.CallOption) (*CommonResp, error) {
	client := tpmtclient.NewTpmtClient(m.cli.Conn())
	return client.SysDictTypeUpdate(ctx, in, opts...)
}

func (m *defaultTpmt) SysDictTypeFindOne(ctx context.Context, in *SysDictTypeFindOneReq, opts ...grpc.CallOption) (*SysDictTypeFindOneResp, error) {
	client := tpmtclient.NewTpmtClient(m.cli.Conn())
	return client.SysDictTypeFindOne(ctx, in, opts...)
}

func (m *defaultTpmt) SysDictTypeList(ctx context.Context, in *SysDictTypeListReq, opts ...grpc.CallOption) (*SysDictTypeListResp, error) {
	client := tpmtclient.NewTpmtClient(m.cli.Conn())
	return client.SysDictTypeList(ctx, in, opts...)
}

// 字典
func (m *defaultTpmt) SysDictAdd(ctx context.Context, in *SysDictAddReq, opts ...grpc.CallOption) (*CommonResp, error) {
	client := tpmtclient.NewTpmtClient(m.cli.Conn())
	return client.SysDictAdd(ctx, in, opts...)
}

func (m *defaultTpmt) SysDictDelete(ctx context.Context, in *SysDictDeleteReq, opts ...grpc.CallOption) (*CommonResp, error) {
	client := tpmtclient.NewTpmtClient(m.cli.Conn())
	return client.SysDictDelete(ctx, in, opts...)
}

func (m *defaultTpmt) SysDictUpdate(ctx context.Context, in *SysDictUpdateReq, opts ...grpc.CallOption) (*CommonResp, error) {
	client := tpmtclient.NewTpmtClient(m.cli.Conn())
	return client.SysDictUpdate(ctx, in, opts...)
}

func (m *defaultTpmt) SysDictFindOne(ctx context.Context, in *SysDictFindOneReq, opts ...grpc.CallOption) (*SysDictFindOneResp, error) {
	client := tpmtclient.NewTpmtClient(m.cli.Conn())
	return client.SysDictFindOne(ctx, in, opts...)
}

func (m *defaultTpmt) SysDictList(ctx context.Context, in *SysDictListReq, opts ...grpc.CallOption) (*SysDictListResp, error) {
	client := tpmtclient.NewTpmtClient(m.cli.Conn())
	return client.SysDictList(ctx, in, opts...)
}

// 资产
func (m *defaultTpmt) TpmtAssetAdd(ctx context.Context, in *TpmtAssetAddReq, opts ...grpc.CallOption) (*CommonResp, error) {
	client := tpmtclient.NewTpmtClient(m.cli.Conn())
	return client.TpmtAssetAdd(ctx, in, opts...)
}

func (m *defaultTpmt) TpmtAssetDelete(ctx context.Context, in *TpmtAssetDeleteReq, opts ...grpc.CallOption) (*CommonResp, error) {
	client := tpmtclient.NewTpmtClient(m.cli.Conn())
	return client.TpmtAssetDelete(ctx, in, opts...)
}

func (m *defaultTpmt) TpmtAssetUpdate(ctx context.Context, in *TpmtAssetUpdateReq, opts ...grpc.CallOption) (*CommonResp, error) {
	client := tpmtclient.NewTpmtClient(m.cli.Conn())
	return client.TpmtAssetUpdate(ctx, in, opts...)
}

func (m *defaultTpmt) TpmtAssetFindOne(ctx context.Context, in *TpmtAssetFindOneReq, opts ...grpc.CallOption) (*TpmtAssetFindOneResp, error) {
	client := tpmtclient.NewTpmtClient(m.cli.Conn())
	return client.TpmtAssetFindOne(ctx, in, opts...)
}

func (m *defaultTpmt) TpmtAssetList(ctx context.Context, in *TpmtAssetListReq, opts ...grpc.CallOption) (*TpmtAssetListResp, error) {
	client := tpmtclient.NewTpmtClient(m.cli.Conn())
	return client.TpmtAssetList(ctx, in, opts...)
}

// 网关
func (m *defaultTpmt) TpmtGatewayAdd(ctx context.Context, in *TpmtGatewayAddReq, opts ...grpc.CallOption) (*CommonResp, error) {
	client := tpmtclient.NewTpmtClient(m.cli.Conn())
	return client.TpmtGatewayAdd(ctx, in, opts...)
}

func (m *defaultTpmt) TpmtGatewayDelete(ctx context.Context, in *TpmtGatewayDeleteReq, opts ...grpc.CallOption) (*CommonResp, error) {
	client := tpmtclient.NewTpmtClient(m.cli.Conn())
	return client.TpmtGatewayDelete(ctx, in, opts...)
}

func (m *defaultTpmt) TpmtGatewayUpdate(ctx context.Context, in *TpmtGatewayUpdateReq, opts ...grpc.CallOption) (*CommonResp, error) {
	client := tpmtclient.NewTpmtClient(m.cli.Conn())
	return client.TpmtGatewayUpdate(ctx, in, opts...)
}

func (m *defaultTpmt) TpmtGatewayFindOne(ctx context.Context, in *TpmtGatewayFindOneReq, opts ...grpc.CallOption) (*TpmtGatewayFindOneResp, error) {
	client := tpmtclient.NewTpmtClient(m.cli.Conn())
	return client.TpmtGatewayFindOne(ctx, in, opts...)
}

func (m *defaultTpmt) TpmtGatewayList(ctx context.Context, in *TpmtGatewayListReq, opts ...grpc.CallOption) (*TpmtGatewayListResp, error) {
	client := tpmtclient.NewTpmtClient(m.cli.Conn())
	return client.TpmtGatewayList(ctx, in, opts...)
}

// 监控点
func (m *defaultTpmt) TpmtMonitorPointAdd(ctx context.Context, in *TpmtMonitorPointAddReq, opts ...grpc.CallOption) (*CommonResp, error) {
	client := tpmtclient.NewTpmtClient(m.cli.Conn())
	return client.TpmtMonitorPointAdd(ctx, in, opts...)
}

func (m *defaultTpmt) TpmtMonitorPointDelete(ctx context.Context, in *TpmtMonitorPointDeleteReq, opts ...grpc.CallOption) (*CommonResp, error) {
	client := tpmtclient.NewTpmtClient(m.cli.Conn())
	return client.TpmtMonitorPointDelete(ctx, in, opts...)
}

func (m *defaultTpmt) TpmtMonitorPointUpdate(ctx context.Context, in *TpmtMonitorPointUpdateReq, opts ...grpc.CallOption) (*CommonResp, error) {
	client := tpmtclient.NewTpmtClient(m.cli.Conn())
	return client.TpmtMonitorPointUpdate(ctx, in, opts...)
}

func (m *defaultTpmt) TpmtMonitorPointFindOne(ctx context.Context, in *TpmtMonitorPointFindOneReq, opts ...grpc.CallOption) (*TpmtMonitorPointFindOneResp, error) {
	client := tpmtclient.NewTpmtClient(m.cli.Conn())
	return client.TpmtMonitorPointFindOne(ctx, in, opts...)
}

func (m *defaultTpmt) TpmtMonitorPointList(ctx context.Context, in *TpmtMonitorPointListReq, opts ...grpc.CallOption) (*TpmtMonitorPointListResp, error) {
	client := tpmtclient.NewTpmtClient(m.cli.Conn())
	return client.TpmtMonitorPointList(ctx, in, opts...)
}
