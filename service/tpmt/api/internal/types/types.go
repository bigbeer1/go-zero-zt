// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.2

package types

type Response struct {
	Code int64       `json:"code"` // 状态码
	Msg  string      `json:"msg"`  // 消息
	Data interface{} `json:"data"` // 数据
}

type SysAuthAddRequest struct {
	NickName  string `json:"nick_name,optional"`  // 机构名
	AuthToken string `json:"auth_token,optional"` // 令牌
	State     int64  `json:"state,optional"`      // 状态 1:正常 2:停用 3:封禁
}

type SysAuthDelRequest struct {
	Id string `path:"id"` // 第三方用户ID
}

type SysAuthInfoRequest struct {
	Id string `form:"id"` // 第三方用户ID
}

type SysAuthListRequest struct {
	Current   int64  `form:"current,default=1,optional"`    // 页码
	PageSize  int64  `form:"page_size,default=10,optional"` // 页数
	NickName  string `form:"nick_name,optional"`            // 机构名
	AuthToken string `form:"auth_token,optional"`           // 令牌
	State     int64  `form:"state,default=99,optional"`     // 状态 1:正常 2:停用 3:封禁
}

type SysAuthUpRequest struct {
	Id        string `json:"id"`                  // 第三方用户ID
	NickName  string `json:"nick_name,optional"`  // 机构名
	AuthToken string `json:"auth_token,optional"` // 令牌
	State     int64  `json:"state,optional"`      // 状态 1:正常 2:停用 3:封禁
}

type SysDictAddRequest struct {
	DictType  string `json:"dict_type,optional"`  // 字典类型
	DictLabel string `json:"dict_label,optional"` // 字典标签
	DictValue string `json:"dict_value,optional"` // 字典键值
	Sort      int64  `json:"sort,optional"`       // 排序
	Remark    string `json:"remark,optional"`     // 备注
	State     int64  `json:"state,optional"`      // 状态
}

type SysDictDelRequest struct {
	Id int64 `path:"id"` // 字典类型ID
}

type SysDictListRequest struct {
	Current   int64  `form:"current,default=1,optional"`    // 页码
	PageSize  int64  `form:"page_size,default=10,optional"` // 页数
	DictType  string `form:"dict_type,optional"`            // 字典类型
	DictLabel string `form:"dict_label,optional"`           // 字典标签
	DictValue string `form:"dict_value,optional"`           // 字典键值
	Remark    string `form:"remark,optional"`               // 备注
	State     int64  `form:"state,default=99,optional"`     // 状态
}

type SysDictTypeAddRequest struct {
	Name     string `json:"name,optional"`      // 字典名称
	DictType string `json:"dict_type,optional"` // 字典类型
	State    int64  `json:"state,optional"`     // 状态
	Remark   string `json:"remark,optional"`    // 描述
	Sort     int64  `json:"sort,optional"`      // 排序
}

type SysDictTypeDelRequest struct {
	Id int64 `path:"id"` // 字典类型ID
}

type SysDictTypeListRequest struct {
	Current  int64  `form:"current,default=1,optional"`    // 页码
	PageSize int64  `form:"page_size,default=10,optional"` // 页数
	Name     string `form:"name,optional"`                 // 字典名称
	DictType string `form:"dict_type,optional"`            // 字典类型
	State    int64  `form:"state,default=99,optional"`     // 状态
	Remark   string `form:"remark,optional"`               // 描述
}

type SysDictTypeUpRequest struct {
	Id       int64  `json:"id"`                 // 字典类型ID
	Name     string `json:"name,optional"`      // 字典名称
	DictType string `json:"dict_type,optional"` // 字典类型
	State    int64  `json:"state,optional"`     // 状态
	Remark   string `json:"remark,optional"`    // 描述
	Sort     int64  `json:"sort,optional"`      // 排序
}

type SysDictUpRequest struct {
	Id        int64  `json:"id"`                  // 字典类型ID
	DictType  string `json:"dict_type,optional"`  // 字典类型
	DictLabel string `json:"dict_label,optional"` // 字典标签
	DictValue string `json:"dict_value,optional"` // 字典键值
	Sort      int64  `json:"sort,optional"`       // 排序
	Remark    string `json:"remark,optional"`     // 备注
	State     int64  `json:"state,optional"`      // 状态
}

type SysInterfaceAddRequest struct {
	Name               string `json:"name"`                 // 接口名称
	Path               string `json:"path"`                 // 接口地址
	InterfaceType      string `json:"interface_type"`       // 接口类型
	InterfaceGroupName string `json:"interface_group_name"` // 接口分组名称
	Remark             string `json:"remark,optional"`      // 备注
	Sort               int64  `json:"sort,optional"`        // sort
}

type SysInterfaceDelRequest struct {
	Id int64 `path:"id"` // 接口ID
}

type SysInterfaceInfoRequest struct {
	Id int64 `form:"id"` // 接口ID
}

type SysInterfaceListRequest struct {
	Current            int64  `form:"current,default=1,optional"`    // 页码
	PageSize           int64  `form:"page_size,default=10,optional"` // 页数
	Name               string `form:"name,optional"`                 // 接口名称
	Path               string `form:"path,optional"`                 // 接口地址
	InterfaceType      string `form:"interface_type,optional"`       // 接口类型
	InterfaceGroupName string `form:"interface_group_name,optional"` // 接口分组名称
	Remark             string `form:"remark,optional"`               // 备注
}

type SysInterfaceUpRequest struct {
	Id                 int64  `json:"id"`                            // 接口ID
	Name               string `json:"name,optional"`                 // 接口名称
	Path               string `json:"path,optional"`                 // 接口地址
	InterfaceType      string `json:"interface_type,optional"`       // 接口类型
	InterfaceGroupName string `json:"interface_group_name,optional"` // 接口分组名称
	Remark             string `json:"remark,optional"`               // 备注
	Sort               int64  `json:"sort,optional"`                 // sort
}

type SysLoginRequest struct {
	Account  string `json:"account"`  // 用户名
	Password string `json:"password"` // 密码
}

type SysMenuAddRequest struct {
	MenuType    int64  `json:"menu_type"`          // 菜单类型(层级关系)
	Name        string `json:"name"`               // 菜单名称
	Title       string `json:"title"`              // 标题
	Path        string `json:"path"`               // 路径
	Component   string `json:"component"`          // 本地路径
	Redirect    string `json:"redirect"`           // 跳转
	Sort        int64  `json:"sort"`               // sort
	Icon        string `json:"icon"`               // 图标
	IsHide      int64  `json:"is_hide"`            // 是否隐藏
	IsKeepAlive int64  `json:"is_keep_alive"`      // 是否缓存
	ParentId    int64  `json:"parent_id,optional"` // 父ID
	IsHome      int64  `json:"is_home"`            // 是否首页
	IsMain      int64  `json:"is_main"`            // 是否主菜单
}

type SysMenuDelRequest struct {
	Id int64 `path:"id"` // 菜单ID
}

type SysMenuInfoRequest struct {
	Id int64 `form:"id"` // 菜单ID
}

type SysMenuListRequest struct {
	Current     int64  `form:"current,default=1,optional"`        // 页码
	PageSize    int64  `form:"page_size,default=10,optional"`     // 页数
	MenuType    int64  `form:"menu_type,default=99,optional"`     // 菜单类型(层级关系)
	Name        string `form:"name,optional"`                     // 菜单名称
	Title       string `form:"title,optional"`                    // 标题
	Path        string `form:"path,optional"`                     // 路径
	Component   string `form:"component,optional"`                // 本地路径
	Redirect    string `form:"redirect,optional"`                 // 跳转
	Icon        string `form:"icon,optional"`                     // 图标
	IsHide      int64  `form:"is_hide,default=99,optional"`       // 是否隐藏
	IsKeepAlive int64  `form:"is_keep_alive,default=99,optional"` // 是否缓存
	ParentId    int64  `form:"parent_id,default=99,optional"`     // 父ID
	IsHome      int64  `form:"is_home,default=99,optional"`       // 是否首页
	IsMain      int64  `form:"is_main,default=99,optional"`       // 是否主菜单
}

type SysMenuUpRequest struct {
	Id          int64  `json:"id"`                     // 菜单ID
	MenuType    int64  `json:"menu_type,optional"`     // 菜单类型(层级关系)
	Name        string `json:"name,optional"`          // 菜单名称
	Title       string `json:"title,optional"`         // 标题
	Path        string `json:"path,optional"`          // 路径
	Component   string `json:"component,optional"`     // 本地路径
	Redirect    string `json:"redirect,optional"`      // 跳转
	Sort        int64  `json:"sort,optional"`          // sort
	Icon        string `json:"icon,optional"`          // 图标
	IsHide      int64  `json:"is_hide,optional"`       // 是否隐藏
	IsKeepAlive int64  `json:"is_keep_alive,optional"` // 是否缓存
	ParentId    int64  `json:"parent_id,optional"`     // 父ID
	IsHome      int64  `json:"is_home,optional"`       // 是否首页
	IsMain      int64  `json:"is_main,optional"`       // 是否主菜单
}

type SysRoleAddRequest struct {
	Name            string  `json:"name"`                    // 角色名称
	Remark          string  `json:"remark,optional"`         // 备注
	RoleType        int64   `json:"role_type,options=1|2|3"` // 角色类型 1:管理员角色  2:普通角色  3:第三方角色
	SysMenuIds      []int64 `json:"sys_menu_ids"`            // 菜单IDS
	SysInterfaceIds []int64 `json:"sys_interface_ids"`       // 接口IDS
}

type SysRoleDelRequest struct {
	Id int64 `path:"id"` // 角色ID
}

type SysRoleInfoRequest struct {
	Id int64 `form:"id"` // 角色ID
}

type SysRoleListRequest struct {
	Current  int64  `form:"current,default=1,optional"`    // 页码
	PageSize int64  `form:"page_size,default=10,optional"` // 页数
	Name     string `form:"name,optional"`                 // 角色名称
	Remark   string `form:"remark,optional"`               // 备注
	RoleType int64  `form:"role_type,default=99,optional"` // 角色类型 1:管理员角色  2:普通角色  3:第三方角色
}

type SysRoleUpRequest struct {
	Id              int64   `json:"id"`                                 // 角色ID
	Name            string  `json:"name,optional"`                      // 角色名称
	Remark          string  `json:"remark,optional"`                    // 备注
	RoleType        int64   `json:"role_type,optional,options=0|1|2|3"` // 角色类型 1:管理员角色  2:普通角色  3:第三方角色
	SysMenuIds      []int64 `json:"sys_menu_ids"`                       // 菜单IDS
	SysInterfaceIds []int64 `json:"sys_interface_ids"`                  // 接口IDS
}

type SysUserAddRequest struct {
	Account  string `json:"account"`              // 用户名
	NickName string `json:"nick_name"`            // 姓名
	Password string `json:"password"`             // 密码
	State    int64  `json:"state,options=1|2|3,"` // 状态 1:正常 2:停用 3:封禁
	RoldId   int64  `json:"role_id,optional"`     // 角色ID
}

type SysUserDelRequest struct {
	Id string `path:"id"` // 用户ID
}

type SysUserInfoRequest struct {
	Id string `form:"id"` // 用户ID
}

type SysUserListRequest struct {
	Current  int64  `form:"current,default=1,optional"`    // 页码
	PageSize int64  `form:"page_size,default=10,optional"` // 页数
	NickName string `form:"nick_name,optional"`            // 姓名
	State    int64  `form:"state,default=99,optional"`     // 状态 1:正常 2:停用 3:封禁
}

type SysUserResetPwdRequest struct {
	Id       string `json:"id"`                // 用户ID
	Password string `json:"password,optional"` // 密码
}

type SysUserUpMyInfoRequest struct {
	NickName string `json:"nick_name"` // 姓名
}

type SysUserUpMyPwdRequest struct {
	OldPassword string `json:"old_password"` // 旧密码
	NewPassword string `json:"new_password"` // 新密码
}

type SysUserUpRequest struct {
	Id       string `json:"id"`                  // 用户ID
	NickName string `json:"nick_name,optional"`  // 姓名
	State    int64  `json:"state,options=1|2|3"` // 状态 1:正常 2:停用 3:封禁
	RoldId   int64  `json:"role_id"`             // 角色ID
}

type TpmtAssetAddRequest struct {
	AssetType    int64  `json:"asset_type,optional"`    // 资产类型
	AssetCode    string `json:"asset_code,optional"`    // 资产编号
	AssetName    string `json:"asset_name,optional"`    // 资产名称
	AssetModel   string `json:"asset_model,optional"`   // 资产型号
	ManuFacturer string `json:"manu_facturer,optional"` // 生产厂家
	Voltage      string `json:"voltage,optional"`       // 电压
	Capacity     string `json:"capacity,optional"`      // 容量
}

type TpmtAssetDelRequest struct {
	Id string `path:"id"` // 资产ID
}

type TpmtAssetInfoRequest struct {
	Id string `form:"id"` // 资产ID
}

type TpmtAssetListRequest struct {
	Current      int64  `form:"current,default=1,optional"`     // 页码
	PageSize     int64  `form:"page_size,default=10,optional"`  // 页数
	AssetType    int64  `form:"asset_type,default=99,optional"` // 资产类型
	AssetCode    string `form:"asset_code,optional"`            // 资产编号
	AssetName    string `form:"asset_name,optional"`            // 资产名称
	AssetModel   string `form:"asset_model,optional"`           // 资产型号
	ManuFacturer string `form:"manu_facturer,optional"`         // 生产厂家
	Voltage      string `form:"voltage,optional"`               // 电压
	Capacity     string `form:"capacity,optional"`              // 容量
}

type TpmtAssetUpRequest struct {
	Id           string `json:"id"`                     // 资产ID
	AssetType    int64  `json:"asset_type,optional"`    // 资产类型
	AssetCode    string `json:"asset_code,optional"`    // 资产编号
	AssetName    string `json:"asset_name,optional"`    // 资产名称
	AssetModel   string `json:"asset_model,optional"`   // 资产型号
	ManuFacturer string `json:"manu_facturer,optional"` // 生产厂家
	Voltage      string `json:"voltage,optional"`       // 电压
	Capacity     string `json:"capacity,optional"`      // 容量
}

type TpmtGatewayAddRequest struct {
	GatewayName  string `json:"gateway_name,optional"`  // 网关名称
	GatewayModel string `json:"gateway_model,optional"` // 网关型号
	ManuFacturer string `json:"manu_facturer,optional"` // 生产厂家
	Agreement    int64  `json:"agreement,optional"`     // 协议 默认1:modbus
	BaudRate     int64  `json:"baud_rate,optional"`     // 波特率
	Parity       string `json:"parity,optional"`        // 校验
	DataBits     int64  `json:"data_bits,optional"`     // 数据位
	StopBits     int64  `json:"stop_bits,optional"`     // 停止位
	ComPort      string `json:"com_port,optional"`      // com端口
	AddressCode  int64  `json:"address_code,optional"`  // 地址码
}

type TpmtGatewayDelRequest struct {
	Id string `path:"id"` // 采集器ID/网关
}

type TpmtGatewayInfoRequest struct {
	Id string `form:"id"` // 采集器ID/网关
}

type TpmtGatewayListRequest struct {
	Current      int64  `form:"current,default=1,optional"`       // 页码
	PageSize     int64  `form:"page_size,default=10,optional"`    // 页数
	GatewayName  string `form:"gateway_name,optional"`            // 网关名称
	GatewayModel string `form:"gateway_model,optional"`           // 网关型号
	ManuFacturer string `form:"manu_facturer,optional"`           // 生产厂家
	Agreement    int64  `form:"agreement,default=99,optional"`    // 协议 默认1:modbus
	BaudRate     int64  `form:"baud_rate,default=99,optional"`    // 波特率
	Parity       string `form:"parity,optional"`                  // 校验
	DataBits     int64  `form:"data_bits,default=99,optional"`    // 数据位
	StopBits     int64  `form:"stop_bits,default=99,optional"`    // 停止位
	ComPort      string `form:"com_port,optional"`                // com端口
	AddressCode  int64  `form:"address_code,default=99,optional"` // 地址码
}

type TpmtGatewayUpRequest struct {
	Id           string `json:"id"`                     // 采集器ID/网关
	GatewayName  string `json:"gateway_name,optional"`  // 网关名称
	GatewayModel string `json:"gateway_model,optional"` // 网关型号
	ManuFacturer string `json:"manu_facturer,optional"` // 生产厂家
	Agreement    int64  `json:"agreement,optional"`     // 协议 默认1:modbus
	BaudRate     int64  `json:"baud_rate,optional"`     // 波特率
	Parity       string `json:"parity,optional"`        // 校验
	DataBits     int64  `json:"data_bits,optional"`     // 数据位
	StopBits     int64  `json:"stop_bits,optional"`     // 停止位
	ComPort      string `json:"com_port,optional"`      // com端口
	AddressCode  int64  `json:"address_code,optional"`  // 地址码
}
