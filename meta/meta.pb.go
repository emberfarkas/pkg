// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.21.12
// source: meta/meta.proto

package meta

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "github.com/go-kratos/kratos/v2/errors"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ErrorReason int32

const (
	// 为某个枚举单独设置错误码
	ErrorReason_MD_NOT_FOUND          ErrorReason = 0
	ErrorReason_DP_NOT_FOUND          ErrorReason = 1
	ErrorReason_TOKEN_NOT_FOUND       ErrorReason = 2
	ErrorReason_UA_NOT_FOUND          ErrorReason = 3
	ErrorReason_REMOTE_ADDR_NOT_FOUND ErrorReason = 4
)

// Enum value maps for ErrorReason.
var (
	ErrorReason_name = map[int32]string{
		0: "MD_NOT_FOUND",
		1: "DP_NOT_FOUND",
		2: "TOKEN_NOT_FOUND",
		3: "UA_NOT_FOUND",
		4: "REMOTE_ADDR_NOT_FOUND",
	}
	ErrorReason_value = map[string]int32{
		"MD_NOT_FOUND":          0,
		"DP_NOT_FOUND":          1,
		"TOKEN_NOT_FOUND":       2,
		"UA_NOT_FOUND":          3,
		"REMOTE_ADDR_NOT_FOUND": 4,
	}
)

func (x ErrorReason) Enum() *ErrorReason {
	p := new(ErrorReason)
	*p = x
	return p
}

func (x ErrorReason) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ErrorReason) Descriptor() protoreflect.EnumDescriptor {
	return file_meta_meta_proto_enumTypes[0].Descriptor()
}

func (ErrorReason) Type() protoreflect.EnumType {
	return &file_meta_meta_proto_enumTypes[0]
}

func (x ErrorReason) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ErrorReason.Descriptor instead.
func (ErrorReason) EnumDescriptor() ([]byte, []int) {
	return file_meta_meta_proto_rawDescGZIP(), []int{0}
}

type DataPermission struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId    uint64 `protobuf:"varint,1,opt,name=userId,proto3" json:"userId,omitempty"`
	RoleId    int32  `protobuf:"varint,2,opt,name=roleId,proto3" json:"roleId,omitempty"`
	RoleKey   string `protobuf:"bytes,3,opt,name=roleKey,proto3" json:"roleKey,omitempty"`
	DataScope string `protobuf:"bytes,4,opt,name=dataScope,proto3" json:"dataScope,omitempty"`
}

func (x *DataPermission) Reset() {
	*x = DataPermission{}
	if protoimpl.UnsafeEnabled {
		mi := &file_meta_meta_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataPermission) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataPermission) ProtoMessage() {}

func (x *DataPermission) ProtoReflect() protoreflect.Message {
	mi := &file_meta_meta_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataPermission.ProtoReflect.Descriptor instead.
func (*DataPermission) Descriptor() ([]byte, []int) {
	return file_meta_meta_proto_rawDescGZIP(), []int{0}
}

func (x *DataPermission) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *DataPermission) GetRoleId() int32 {
	if x != nil {
		return x.RoleId
	}
	return 0
}

func (x *DataPermission) GetRoleKey() string {
	if x != nil {
		return x.RoleKey
	}
	return ""
}

func (x *DataPermission) GetDataScope() string {
	if x != nil {
		return x.DataScope
	}
	return ""
}

var File_meta_meta_proto protoreflect.FileDescriptor

var file_meta_meta_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x6d, 0x65, 0x74, 0x61, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65,
	0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e,
	0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x13, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x78, 0x0a, 0x0e, 0x44, 0x61, 0x74, 0x61, 0x50, 0x65, 0x72, 0x6d,
	0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x16,
	0x0a, 0x06, 0x72, 0x6f, 0x6c, 0x65, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06,
	0x72, 0x6f, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x6f, 0x6c, 0x65, 0x4b, 0x65,
	0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x72, 0x6f, 0x6c, 0x65, 0x4b, 0x65, 0x79,
	0x12, 0x1c, 0x0a, 0x09, 0x64, 0x61, 0x74, 0x61, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x64, 0x61, 0x74, 0x61, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x2a, 0x79,
	0x0a, 0x0b, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x10, 0x0a,
	0x0c, 0x4d, 0x44, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x46, 0x4f, 0x55, 0x4e, 0x44, 0x10, 0x00, 0x12,
	0x10, 0x0a, 0x0c, 0x44, 0x50, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x46, 0x4f, 0x55, 0x4e, 0x44, 0x10,
	0x01, 0x12, 0x13, 0x0a, 0x0f, 0x54, 0x4f, 0x4b, 0x45, 0x4e, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x46,
	0x4f, 0x55, 0x4e, 0x44, 0x10, 0x02, 0x12, 0x10, 0x0a, 0x0c, 0x55, 0x41, 0x5f, 0x4e, 0x4f, 0x54,
	0x5f, 0x46, 0x4f, 0x55, 0x4e, 0x44, 0x10, 0x03, 0x12, 0x19, 0x0a, 0x15, 0x52, 0x45, 0x4d, 0x4f,
	0x54, 0x45, 0x5f, 0x41, 0x44, 0x44, 0x52, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x46, 0x4f, 0x55, 0x4e,
	0x44, 0x10, 0x04, 0x1a, 0x04, 0xa0, 0x45, 0xf4, 0x03, 0x42, 0x24, 0x5a, 0x22, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2d, 0x62, 0x61, 0x6d, 0x62, 0x6f,
	0x6f, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x3b, 0x6d, 0x65, 0x74, 0x61, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_meta_meta_proto_rawDescOnce sync.Once
	file_meta_meta_proto_rawDescData = file_meta_meta_proto_rawDesc
)

func file_meta_meta_proto_rawDescGZIP() []byte {
	file_meta_meta_proto_rawDescOnce.Do(func() {
		file_meta_meta_proto_rawDescData = protoimpl.X.CompressGZIP(file_meta_meta_proto_rawDescData)
	})
	return file_meta_meta_proto_rawDescData
}

var file_meta_meta_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_meta_meta_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_meta_meta_proto_goTypes = []interface{}{
	(ErrorReason)(0),       // 0: meta.ErrorReason
	(*DataPermission)(nil), // 1: meta.DataPermission
}
var file_meta_meta_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_meta_meta_proto_init() }
func file_meta_meta_proto_init() {
	if File_meta_meta_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_meta_meta_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DataPermission); i {
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
			RawDescriptor: file_meta_meta_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_meta_meta_proto_goTypes,
		DependencyIndexes: file_meta_meta_proto_depIdxs,
		EnumInfos:         file_meta_meta_proto_enumTypes,
		MessageInfos:      file_meta_meta_proto_msgTypes,
	}.Build()
	File_meta_meta_proto = out.File
	file_meta_meta_proto_rawDesc = nil
	file_meta_meta_proto_goTypes = nil
	file_meta_meta_proto_depIdxs = nil
}
