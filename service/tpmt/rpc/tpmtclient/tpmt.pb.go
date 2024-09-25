// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: tpmt.proto

package tpmtclient

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CommonResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CommonResp) Reset() {
	*x = CommonResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tpmt_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommonResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommonResp) ProtoMessage() {}

func (x *CommonResp) ProtoReflect() protoreflect.Message {
	mi := &file_tpmt_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommonResp.ProtoReflect.Descriptor instead.
func (*CommonResp) Descriptor() ([]byte, []int) {
	return file_tpmt_proto_rawDescGZIP(), []int{0}
}

// SysUser 添加
type SysUserAddReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Account     string `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`                            // 用户名
	NickName    string `protobuf:"bytes,2,opt,name=nick_name,json=nickName,proto3" json:"nick_name,omitempty"`          // 姓名
	Password    string `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`                          // 密码
	State       int64  `protobuf:"varint,4,opt,name=state,proto3" json:"state,omitempty"`                               // 状态 1:正常 2:停用 3:封禁
	CreatedName string `protobuf:"bytes,5,opt,name=created_name,json=createdName,proto3" json:"created_name,omitempty"` // 创建人
}

func (x *SysUserAddReq) Reset() {
	*x = SysUserAddReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tpmt_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SysUserAddReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SysUserAddReq) ProtoMessage() {}

func (x *SysUserAddReq) ProtoReflect() protoreflect.Message {
	mi := &file_tpmt_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SysUserAddReq.ProtoReflect.Descriptor instead.
func (*SysUserAddReq) Descriptor() ([]byte, []int) {
	return file_tpmt_proto_rawDescGZIP(), []int{1}
}

func (x *SysUserAddReq) GetAccount() string {
	if x != nil {
		return x.Account
	}
	return ""
}

func (x *SysUserAddReq) GetNickName() string {
	if x != nil {
		return x.NickName
	}
	return ""
}

func (x *SysUserAddReq) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *SysUserAddReq) GetState() int64 {
	if x != nil {
		return x.State
	}
	return 0
}

func (x *SysUserAddReq) GetCreatedName() string {
	if x != nil {
		return x.CreatedName
	}
	return ""
}

// SysUser 删除
type SysUserDeleteReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`                                      // 用户ID
	DeletedName string `protobuf:"bytes,2,opt,name=deleted_name,json=deletedName,proto3" json:"deleted_name,omitempty"` // 删除人
}

func (x *SysUserDeleteReq) Reset() {
	*x = SysUserDeleteReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tpmt_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SysUserDeleteReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SysUserDeleteReq) ProtoMessage() {}

func (x *SysUserDeleteReq) ProtoReflect() protoreflect.Message {
	mi := &file_tpmt_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SysUserDeleteReq.ProtoReflect.Descriptor instead.
func (*SysUserDeleteReq) Descriptor() ([]byte, []int) {
	return file_tpmt_proto_rawDescGZIP(), []int{2}
}

func (x *SysUserDeleteReq) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *SysUserDeleteReq) GetDeletedName() string {
	if x != nil {
		return x.DeletedName
	}
	return ""
}

// SysUser 更新
type SysUserUpdateReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`                                      // 用户ID
	NickName    string `protobuf:"bytes,3,opt,name=nick_name,json=nickName,proto3" json:"nick_name,omitempty"`          // 姓名
	State       int64  `protobuf:"varint,5,opt,name=state,proto3" json:"state,omitempty"`                               // 状态 1:正常 2:停用 3:封禁
	UpdatedName string `protobuf:"bytes,6,opt,name=updated_name,json=updatedName,proto3" json:"updated_name,omitempty"` // 更新人
}

func (x *SysUserUpdateReq) Reset() {
	*x = SysUserUpdateReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tpmt_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SysUserUpdateReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SysUserUpdateReq) ProtoMessage() {}

func (x *SysUserUpdateReq) ProtoReflect() protoreflect.Message {
	mi := &file_tpmt_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SysUserUpdateReq.ProtoReflect.Descriptor instead.
func (*SysUserUpdateReq) Descriptor() ([]byte, []int) {
	return file_tpmt_proto_rawDescGZIP(), []int{3}
}

func (x *SysUserUpdateReq) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *SysUserUpdateReq) GetNickName() string {
	if x != nil {
		return x.NickName
	}
	return ""
}

func (x *SysUserUpdateReq) GetState() int64 {
	if x != nil {
		return x.State
	}
	return 0
}

func (x *SysUserUpdateReq) GetUpdatedName() string {
	if x != nil {
		return x.UpdatedName
	}
	return ""
}

// SysUser 单个查询
type SysUserFindOneReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"` // 用户ID
}

func (x *SysUserFindOneReq) Reset() {
	*x = SysUserFindOneReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tpmt_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SysUserFindOneReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SysUserFindOneReq) ProtoMessage() {}

func (x *SysUserFindOneReq) ProtoReflect() protoreflect.Message {
	mi := &file_tpmt_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SysUserFindOneReq.ProtoReflect.Descriptor instead.
func (*SysUserFindOneReq) Descriptor() ([]byte, []int) {
	return file_tpmt_proto_rawDescGZIP(), []int{4}
}

func (x *SysUserFindOneReq) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

// SysUser 单个查询返回
type SysUserFindOneResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`                                      // 用户ID
	Account     string `protobuf:"bytes,2,opt,name=account,proto3" json:"account,omitempty"`                            // 用户名
	NickName    string `protobuf:"bytes,3,opt,name=nick_name,json=nickName,proto3" json:"nick_name,omitempty"`          // 姓名
	Password    string `protobuf:"bytes,4,opt,name=password,proto3" json:"password,omitempty"`                          // 密码
	State       int64  `protobuf:"varint,5,opt,name=state,proto3" json:"state,omitempty"`                               // 状态 1:正常 2:停用 3:封禁
	CreatedName string `protobuf:"bytes,6,opt,name=created_name,json=createdName,proto3" json:"created_name,omitempty"` // 创建人
	CreatedAt   int64  `protobuf:"varint,7,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`      // 创建时间
	UpdatedName string `protobuf:"bytes,8,opt,name=updated_name,json=updatedName,proto3" json:"updated_name,omitempty"` // 更新人
	UpdatedAt   int64  `protobuf:"varint,9,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`      // 更新时间
}

func (x *SysUserFindOneResp) Reset() {
	*x = SysUserFindOneResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tpmt_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SysUserFindOneResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SysUserFindOneResp) ProtoMessage() {}

func (x *SysUserFindOneResp) ProtoReflect() protoreflect.Message {
	mi := &file_tpmt_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SysUserFindOneResp.ProtoReflect.Descriptor instead.
func (*SysUserFindOneResp) Descriptor() ([]byte, []int) {
	return file_tpmt_proto_rawDescGZIP(), []int{5}
}

func (x *SysUserFindOneResp) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *SysUserFindOneResp) GetAccount() string {
	if x != nil {
		return x.Account
	}
	return ""
}

func (x *SysUserFindOneResp) GetNickName() string {
	if x != nil {
		return x.NickName
	}
	return ""
}

func (x *SysUserFindOneResp) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *SysUserFindOneResp) GetState() int64 {
	if x != nil {
		return x.State
	}
	return 0
}

func (x *SysUserFindOneResp) GetCreatedName() string {
	if x != nil {
		return x.CreatedName
	}
	return ""
}

func (x *SysUserFindOneResp) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *SysUserFindOneResp) GetUpdatedName() string {
	if x != nil {
		return x.UpdatedName
	}
	return ""
}

func (x *SysUserFindOneResp) GetUpdatedAt() int64 {
	if x != nil {
		return x.UpdatedAt
	}
	return 0
}

// SysUser 分页查询
type SysUserListReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Current  int64  `protobuf:"varint,1,opt,name=current,proto3" json:"current,omitempty"`                   // 页码
	PageSize int64  `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"` // 页数
	NickName string `protobuf:"bytes,4,opt,name=nick_name,json=nickName,proto3" json:"nick_name,omitempty"`  // 姓名
	State    int64  `protobuf:"varint,6,opt,name=state,proto3" json:"state,omitempty"`                       // 状态 1:正常 2:停用 3:封禁
}

func (x *SysUserListReq) Reset() {
	*x = SysUserListReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tpmt_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SysUserListReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SysUserListReq) ProtoMessage() {}

func (x *SysUserListReq) ProtoReflect() protoreflect.Message {
	mi := &file_tpmt_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SysUserListReq.ProtoReflect.Descriptor instead.
func (*SysUserListReq) Descriptor() ([]byte, []int) {
	return file_tpmt_proto_rawDescGZIP(), []int{6}
}

func (x *SysUserListReq) GetCurrent() int64 {
	if x != nil {
		return x.Current
	}
	return 0
}

func (x *SysUserListReq) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *SysUserListReq) GetNickName() string {
	if x != nil {
		return x.NickName
	}
	return ""
}

func (x *SysUserListReq) GetState() int64 {
	if x != nil {
		return x.State
	}
	return 0
}

// SysUser 分页查询返回
type SysUserListResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Total int64              `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"` // 总数
	List  []*SysUserListData `protobuf:"bytes,2,rep,name=list,proto3" json:"list,omitempty"`    // 内容
}

func (x *SysUserListResp) Reset() {
	*x = SysUserListResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tpmt_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SysUserListResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SysUserListResp) ProtoMessage() {}

func (x *SysUserListResp) ProtoReflect() protoreflect.Message {
	mi := &file_tpmt_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SysUserListResp.ProtoReflect.Descriptor instead.
func (*SysUserListResp) Descriptor() ([]byte, []int) {
	return file_tpmt_proto_rawDescGZIP(), []int{7}
}

func (x *SysUserListResp) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *SysUserListResp) GetList() []*SysUserListData {
	if x != nil {
		return x.List
	}
	return nil
}

// SysUser 列表信息
type SysUserListData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`                                      // 用户ID
	Account     string `protobuf:"bytes,2,opt,name=account,proto3" json:"account,omitempty"`                            // 用户名
	NickName    string `protobuf:"bytes,3,opt,name=nick_name,json=nickName,proto3" json:"nick_name,omitempty"`          // 姓名
	State       int64  `protobuf:"varint,5,opt,name=state,proto3" json:"state,omitempty"`                               // 状态 1:正常 2:停用 3:封禁
	CreatedName string `protobuf:"bytes,6,opt,name=created_name,json=createdName,proto3" json:"created_name,omitempty"` // 创建人
	CreatedAt   int64  `protobuf:"varint,7,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`      // 创建时间
	UpdatedName string `protobuf:"bytes,8,opt,name=updated_name,json=updatedName,proto3" json:"updated_name,omitempty"` // 更新人
	UpdatedAt   int64  `protobuf:"varint,9,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`      // 更新时间
}

func (x *SysUserListData) Reset() {
	*x = SysUserListData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tpmt_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SysUserListData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SysUserListData) ProtoMessage() {}

func (x *SysUserListData) ProtoReflect() protoreflect.Message {
	mi := &file_tpmt_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SysUserListData.ProtoReflect.Descriptor instead.
func (*SysUserListData) Descriptor() ([]byte, []int) {
	return file_tpmt_proto_rawDescGZIP(), []int{8}
}

func (x *SysUserListData) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *SysUserListData) GetAccount() string {
	if x != nil {
		return x.Account
	}
	return ""
}

func (x *SysUserListData) GetNickName() string {
	if x != nil {
		return x.NickName
	}
	return ""
}

func (x *SysUserListData) GetState() int64 {
	if x != nil {
		return x.State
	}
	return 0
}

func (x *SysUserListData) GetCreatedName() string {
	if x != nil {
		return x.CreatedName
	}
	return ""
}

func (x *SysUserListData) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *SysUserListData) GetUpdatedName() string {
	if x != nil {
		return x.UpdatedName
	}
	return ""
}

func (x *SysUserListData) GetUpdatedAt() int64 {
	if x != nil {
		return x.UpdatedAt
	}
	return 0
}

var File_tpmt_proto protoreflect.FileDescriptor

var file_tpmt_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x74, 0x70, 0x6d, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x74, 0x70,
	0x6d, 0x74, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x22, 0x0c, 0x0a, 0x0a, 0x43, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x22, 0x9b, 0x01, 0x0a, 0x0d, 0x53, 0x79, 0x73, 0x55, 0x73,
	0x65, 0x72, 0x41, 0x64, 0x64, 0x52, 0x65, 0x71, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x6e, 0x69, 0x63, 0x6b, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x73,
	0x74, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74,
	0x65, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x4e, 0x61, 0x6d, 0x65, 0x22, 0x45, 0x0a, 0x10, 0x53, 0x79, 0x73, 0x55, 0x73, 0x65, 0x72, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x64, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x64, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x78, 0x0a, 0x10, 0x53,
	0x79, 0x73, 0x55, 0x73, 0x65, 0x72, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x1b, 0x0a, 0x09, 0x6e, 0x69, 0x63, 0x6b, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x73, 0x74, 0x61,
	0x74, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x23, 0x0a, 0x11, 0x53, 0x79, 0x73, 0x55, 0x73, 0x65, 0x72,
	0x46, 0x69, 0x6e, 0x64, 0x4f, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x91, 0x02, 0x0a, 0x12, 0x53,
	0x79, 0x73, 0x55, 0x73, 0x65, 0x72, 0x46, 0x69, 0x6e, 0x64, 0x4f, 0x6e, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x6e,
	0x69, 0x63, 0x6b, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x6e, 0x69, 0x63, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a,
	0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x21, 0x0a, 0x0c,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x7a,
	0x0a, 0x0e, 0x53, 0x79, 0x73, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71,
	0x12, 0x18, 0x0a, 0x07, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x07, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61,
	0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x70,
	0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6e, 0x69, 0x63, 0x6b, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x69, 0x63, 0x6b,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x22, 0x58, 0x0a, 0x0f, 0x53, 0x79,
	0x73, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x14, 0x0a,
	0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x12, 0x2f, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x1b, 0x2e, 0x74, 0x70, 0x6d, 0x74, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x53,
	0x79, 0x73, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04,
	0x6c, 0x69, 0x73, 0x74, 0x22, 0xf2, 0x01, 0x0a, 0x0f, 0x53, 0x79, 0x73, 0x55, 0x73, 0x65, 0x72,
	0x4c, 0x69, 0x73, 0x74, 0x44, 0x61, 0x74, 0x61, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x6e, 0x69, 0x63, 0x6b, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05,
	0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x32, 0xee, 0x02, 0x0a, 0x04, 0x54, 0x70,
	0x6d, 0x74, 0x12, 0x3f, 0x0a, 0x0a, 0x53, 0x79, 0x73, 0x55, 0x73, 0x65, 0x72, 0x41, 0x64, 0x64,
	0x12, 0x19, 0x2e, 0x74, 0x70, 0x6d, 0x74, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x53, 0x79,
	0x73, 0x55, 0x73, 0x65, 0x72, 0x41, 0x64, 0x64, 0x52, 0x65, 0x71, 0x1a, 0x16, 0x2e, 0x74, 0x70,
	0x6d, 0x74, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x12, 0x45, 0x0a, 0x0d, 0x53, 0x79, 0x73, 0x55, 0x73, 0x65, 0x72, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x12, 0x1c, 0x2e, 0x74, 0x70, 0x6d, 0x74, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x2e, 0x53, 0x79, 0x73, 0x55, 0x73, 0x65, 0x72, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52,
	0x65, 0x71, 0x1a, 0x16, 0x2e, 0x74, 0x70, 0x6d, 0x74, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e,
	0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x12, 0x45, 0x0a, 0x0d, 0x53, 0x79,
	0x73, 0x55, 0x73, 0x65, 0x72, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x1c, 0x2e, 0x74, 0x70,
	0x6d, 0x74, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x53, 0x79, 0x73, 0x55, 0x73, 0x65, 0x72,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x16, 0x2e, 0x74, 0x70, 0x6d, 0x74,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x12, 0x4f, 0x0a, 0x0e, 0x53, 0x79, 0x73, 0x55, 0x73, 0x65, 0x72, 0x46, 0x69, 0x6e, 0x64,
	0x4f, 0x6e, 0x65, 0x12, 0x1d, 0x2e, 0x74, 0x70, 0x6d, 0x74, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x2e, 0x53, 0x79, 0x73, 0x55, 0x73, 0x65, 0x72, 0x46, 0x69, 0x6e, 0x64, 0x4f, 0x6e, 0x65, 0x52,
	0x65, 0x71, 0x1a, 0x1e, 0x2e, 0x74, 0x70, 0x6d, 0x74, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e,
	0x53, 0x79, 0x73, 0x55, 0x73, 0x65, 0x72, 0x46, 0x69, 0x6e, 0x64, 0x4f, 0x6e, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x12, 0x46, 0x0a, 0x0b, 0x53, 0x79, 0x73, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x69, 0x73,
	0x74, 0x12, 0x1a, 0x2e, 0x74, 0x70, 0x6d, 0x74, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x53,
	0x79, 0x73, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x1b, 0x2e,
	0x74, 0x70, 0x6d, 0x74, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x53, 0x79, 0x73, 0x55, 0x73,
	0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x42, 0x0e, 0x5a, 0x0c, 0x2e, 0x2f,
	0x74, 0x70, 0x6d, 0x74, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_tpmt_proto_rawDescOnce sync.Once
	file_tpmt_proto_rawDescData = file_tpmt_proto_rawDesc
)

func file_tpmt_proto_rawDescGZIP() []byte {
	file_tpmt_proto_rawDescOnce.Do(func() {
		file_tpmt_proto_rawDescData = protoimpl.X.CompressGZIP(file_tpmt_proto_rawDescData)
	})
	return file_tpmt_proto_rawDescData
}

var file_tpmt_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_tpmt_proto_goTypes = []interface{}{
	(*CommonResp)(nil),         // 0: tpmtclient.CommonResp
	(*SysUserAddReq)(nil),      // 1: tpmtclient.SysUserAddReq
	(*SysUserDeleteReq)(nil),   // 2: tpmtclient.SysUserDeleteReq
	(*SysUserUpdateReq)(nil),   // 3: tpmtclient.SysUserUpdateReq
	(*SysUserFindOneReq)(nil),  // 4: tpmtclient.SysUserFindOneReq
	(*SysUserFindOneResp)(nil), // 5: tpmtclient.SysUserFindOneResp
	(*SysUserListReq)(nil),     // 6: tpmtclient.SysUserListReq
	(*SysUserListResp)(nil),    // 7: tpmtclient.SysUserListResp
	(*SysUserListData)(nil),    // 8: tpmtclient.SysUserListData
}
var file_tpmt_proto_depIdxs = []int32{
	8, // 0: tpmtclient.SysUserListResp.list:type_name -> tpmtclient.SysUserListData
	1, // 1: tpmtclient.Tpmt.SysUserAdd:input_type -> tpmtclient.SysUserAddReq
	2, // 2: tpmtclient.Tpmt.SysUserDelete:input_type -> tpmtclient.SysUserDeleteReq
	3, // 3: tpmtclient.Tpmt.SysUserUpdate:input_type -> tpmtclient.SysUserUpdateReq
	4, // 4: tpmtclient.Tpmt.SysUserFindOne:input_type -> tpmtclient.SysUserFindOneReq
	6, // 5: tpmtclient.Tpmt.SysUserList:input_type -> tpmtclient.SysUserListReq
	0, // 6: tpmtclient.Tpmt.SysUserAdd:output_type -> tpmtclient.CommonResp
	0, // 7: tpmtclient.Tpmt.SysUserDelete:output_type -> tpmtclient.CommonResp
	0, // 8: tpmtclient.Tpmt.SysUserUpdate:output_type -> tpmtclient.CommonResp
	5, // 9: tpmtclient.Tpmt.SysUserFindOne:output_type -> tpmtclient.SysUserFindOneResp
	7, // 10: tpmtclient.Tpmt.SysUserList:output_type -> tpmtclient.SysUserListResp
	6, // [6:11] is the sub-list for method output_type
	1, // [1:6] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_tpmt_proto_init() }
func file_tpmt_proto_init() {
	if File_tpmt_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tpmt_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommonResp); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_tpmt_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SysUserAddReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_tpmt_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SysUserDeleteReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_tpmt_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SysUserUpdateReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_tpmt_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SysUserFindOneReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_tpmt_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SysUserFindOneResp); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_tpmt_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SysUserListReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_tpmt_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SysUserListResp); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_tpmt_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SysUserListData); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_tpmt_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_tpmt_proto_goTypes,
		DependencyIndexes: file_tpmt_proto_depIdxs,
		MessageInfos:      file_tpmt_proto_msgTypes,
	}.Build()
	File_tpmt_proto = out.File
	file_tpmt_proto_rawDesc = nil
	file_tpmt_proto_goTypes = nil
	file_tpmt_proto_depIdxs = nil
}
