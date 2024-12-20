// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: websocket.proto

package websocketclient

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

// 通用空返回
type CommonResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CommonResp) Reset() {
	*x = CommonResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_websocket_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommonResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommonResp) ProtoMessage() {}

func (x *CommonResp) ProtoReflect() protoreflect.Message {
	mi := &file_websocket_proto_msgTypes[0]
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
	return file_websocket_proto_rawDescGZIP(), []int{0}
}

type AlarmMessageReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ts           int64  `protobuf:"varint,1,opt,name=ts,proto3" json:"ts,omitempty"`                                        // 告警事件
	Id           string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`                                         // 告警ID
	Mid          string `protobuf:"bytes,3,opt,name=mid,proto3" json:"mid,omitempty"`                                       // 告警设备ID
	Name         string `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`                                     // 告警设备名称
	AlarmType    int64  `protobuf:"varint,5,opt,name=alarm_type,json=alarmType,proto3" json:"alarm_type,omitempty"`         // 告警类型
	AlarmGrade   int64  `protobuf:"varint,6,opt,name=alarm_grade,json=alarmGrade,proto3" json:"alarm_grade,omitempty"`      // 告警等级
	AlarmContent string `protobuf:"bytes,7,opt,name=alarm_content,json=alarmContent,proto3" json:"alarm_content,omitempty"` // 告警内容
	AssetCode    string `protobuf:"bytes,8,opt,name=asset_code,json=assetCode,proto3" json:"asset_code,omitempty"`          // 柜号
}

func (x *AlarmMessageReq) Reset() {
	*x = AlarmMessageReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_websocket_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AlarmMessageReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AlarmMessageReq) ProtoMessage() {}

func (x *AlarmMessageReq) ProtoReflect() protoreflect.Message {
	mi := &file_websocket_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AlarmMessageReq.ProtoReflect.Descriptor instead.
func (*AlarmMessageReq) Descriptor() ([]byte, []int) {
	return file_websocket_proto_rawDescGZIP(), []int{1}
}

func (x *AlarmMessageReq) GetTs() int64 {
	if x != nil {
		return x.Ts
	}
	return 0
}

func (x *AlarmMessageReq) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *AlarmMessageReq) GetMid() string {
	if x != nil {
		return x.Mid
	}
	return ""
}

func (x *AlarmMessageReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AlarmMessageReq) GetAlarmType() int64 {
	if x != nil {
		return x.AlarmType
	}
	return 0
}

func (x *AlarmMessageReq) GetAlarmGrade() int64 {
	if x != nil {
		return x.AlarmGrade
	}
	return 0
}

func (x *AlarmMessageReq) GetAlarmContent() string {
	if x != nil {
		return x.AlarmContent
	}
	return ""
}

func (x *AlarmMessageReq) GetAssetCode() string {
	if x != nil {
		return x.AssetCode
	}
	return ""
}

var File_websocket_proto protoreflect.FileDescriptor

var file_websocket_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x77, 0x65, 0x62, 0x73, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0f, 0x77, 0x65, 0x62, 0x73, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x63, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x22, 0x0c, 0x0a, 0x0a, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x22, 0xdb, 0x01, 0x0a, 0x0f, 0x41, 0x6c, 0x61, 0x72, 0x6d, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x74, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x02, 0x74, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6d, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x6c,
	0x61, 0x72, 0x6d, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09,
	0x61, 0x6c, 0x61, 0x72, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x6c, 0x61,
	0x72, 0x6d, 0x5f, 0x67, 0x72, 0x61, 0x64, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a,
	0x61, 0x6c, 0x61, 0x72, 0x6d, 0x47, 0x72, 0x61, 0x64, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x61, 0x6c,
	0x61, 0x72, 0x6d, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0c, 0x61, 0x6c, 0x61, 0x72, 0x6d, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12,
	0x1d, 0x0a, 0x0a, 0x61, 0x73, 0x73, 0x65, 0x74, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x73, 0x73, 0x65, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x32, 0x5a,
	0x0a, 0x09, 0x77, 0x65, 0x62, 0x73, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x4d, 0x0a, 0x0c, 0x41,
	0x6c, 0x61, 0x72, 0x6d, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x20, 0x2e, 0x77, 0x65,
	0x62, 0x73, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x41, 0x6c,
	0x61, 0x72, 0x6d, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x1b, 0x2e,
	0x77, 0x65, 0x62, 0x73, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e,
	0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x42, 0x13, 0x5a, 0x11, 0x2e, 0x2f,
	0x77, 0x65, 0x62, 0x73, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_websocket_proto_rawDescOnce sync.Once
	file_websocket_proto_rawDescData = file_websocket_proto_rawDesc
)

func file_websocket_proto_rawDescGZIP() []byte {
	file_websocket_proto_rawDescOnce.Do(func() {
		file_websocket_proto_rawDescData = protoimpl.X.CompressGZIP(file_websocket_proto_rawDescData)
	})
	return file_websocket_proto_rawDescData
}

var file_websocket_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_websocket_proto_goTypes = []interface{}{
	(*CommonResp)(nil),      // 0: websocketclient.CommonResp
	(*AlarmMessageReq)(nil), // 1: websocketclient.AlarmMessageReq
}
var file_websocket_proto_depIdxs = []int32{
	1, // 0: websocketclient.websocket.AlarmMessage:input_type -> websocketclient.AlarmMessageReq
	0, // 1: websocketclient.websocket.AlarmMessage:output_type -> websocketclient.CommonResp
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_websocket_proto_init() }
func file_websocket_proto_init() {
	if File_websocket_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_websocket_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_websocket_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AlarmMessageReq); i {
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
			RawDescriptor: file_websocket_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_websocket_proto_goTypes,
		DependencyIndexes: file_websocket_proto_depIdxs,
		MessageInfos:      file_websocket_proto_msgTypes,
	}.Build()
	File_websocket_proto = out.File
	file_websocket_proto_rawDesc = nil
	file_websocket_proto_goTypes = nil
	file_websocket_proto_depIdxs = nil
}
