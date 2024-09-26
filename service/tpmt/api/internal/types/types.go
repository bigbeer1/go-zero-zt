// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.2

package types

type Response struct {
	Code int64       `json:"code"` // 状态码
	Msg  string      `json:"msg"`  // 消息
	Data interface{} `json:"data"` // 数据
}

type SysLoginRequest struct {
	Account  string `json:"account"`  // 用户名
	Password string `json:"password"` // 密码
}

type SysUserAddRequest struct {
	Account  string `json:"account"`              // 用户名
	NickName string `json:"nick_name"`            // 姓名
	Password string `json:"password"`             // 密码
	State    int64  `json:"state,options=1|2|3,"` // 状态 1:正常 2:停用 3:封禁
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

type SysUserUpRequest struct {
	Id       string `json:"id"`                  // 用户ID
	NickName string `json:"nick_name,optional"`  // 姓名
	State    int64  `json:"state,options=1|2|3"` // 状态 1:正常 2:停用 3:封禁
}
