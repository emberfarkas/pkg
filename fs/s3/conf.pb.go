// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: fs/s3/conf.proto

package s3

import (
	_ "github.com/go-kratos/kratos/v2/errors"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/durationpb"
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
	ErrorReason_NOT_ALLOW_EXT ErrorReason = 0
)

// Enum value maps for ErrorReason.
var (
	ErrorReason_name = map[int32]string{
		0: "NOT_ALLOW_EXT",
	}
	ErrorReason_value = map[string]int32{
		"NOT_ALLOW_EXT": 0,
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
	return file_fs_s3_conf_proto_enumTypes[0].Descriptor()
}

func (ErrorReason) Type() protoreflect.EnumType {
	return &file_fs_s3_conf_proto_enumTypes[0]
}

func (x ErrorReason) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ErrorReason.Descriptor instead.
func (ErrorReason) EnumDescriptor() ([]byte, []int) {
	return file_fs_s3_conf_proto_rawDescGZIP(), []int{0}
}

type Conf struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key    string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Secret string `protobuf:"bytes,2,opt,name=secret,proto3" json:"secret,omitempty"`
	Region string `protobuf:"bytes,3,opt,name=region,proto3" json:"region,omitempty"`
	Bucket string `protobuf:"bytes,4,opt,name=bucket,proto3" json:"bucket,omitempty"`
	Dir    string `protobuf:"bytes,5,opt,name=dir,proto3" json:"dir,omitempty"`
}

func (x *Conf) Reset() {
	*x = Conf{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fs_s3_conf_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Conf) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Conf) ProtoMessage() {}

func (x *Conf) ProtoReflect() protoreflect.Message {
	mi := &file_fs_s3_conf_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Conf.ProtoReflect.Descriptor instead.
func (*Conf) Descriptor() ([]byte, []int) {
	return file_fs_s3_conf_proto_rawDescGZIP(), []int{0}
}

func (x *Conf) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *Conf) GetSecret() string {
	if x != nil {
		return x.Secret
	}
	return ""
}

func (x *Conf) GetRegion() string {
	if x != nil {
		return x.Region
	}
	return ""
}

func (x *Conf) GetBucket() string {
	if x != nil {
		return x.Bucket
	}
	return ""
}

func (x *Conf) GetDir() string {
	if x != nil {
		return x.Dir
	}
	return ""
}

var File_fs_s3_conf_proto protoreflect.FileDescriptor

var file_fs_s3_conf_proto_rawDesc = []byte{
	0x0a, 0x10, 0x66, 0x73, 0x2f, 0x73, 0x33, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x05, 0x66, 0x73, 0x2e, 0x73, 0x33, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x73, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x72,
	0x0a, 0x04, 0x43, 0x6f, 0x6e, 0x66, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x63, 0x72,
	0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74,
	0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x62, 0x75, 0x63, 0x6b,
	0x65, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74,
	0x12, 0x10, 0x0a, 0x03, 0x64, 0x69, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x64,
	0x69, 0x72, 0x2a, 0x26, 0x0a, 0x0b, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x65, 0x61, 0x73, 0x6f,
	0x6e, 0x12, 0x11, 0x0a, 0x0d, 0x4e, 0x4f, 0x54, 0x5f, 0x41, 0x4c, 0x4c, 0x4f, 0x57, 0x5f, 0x45,
	0x58, 0x54, 0x10, 0x00, 0x1a, 0x04, 0xa0, 0x45, 0xf4, 0x03, 0x42, 0x23, 0x5a, 0x21, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2d, 0x62, 0x61, 0x6d, 0x62,
	0x6f, 0x6f, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x66, 0x73, 0x2f, 0x73, 0x33, 0x3b, 0x73, 0x33, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_fs_s3_conf_proto_rawDescOnce sync.Once
	file_fs_s3_conf_proto_rawDescData = file_fs_s3_conf_proto_rawDesc
)

func file_fs_s3_conf_proto_rawDescGZIP() []byte {
	file_fs_s3_conf_proto_rawDescOnce.Do(func() {
		file_fs_s3_conf_proto_rawDescData = protoimpl.X.CompressGZIP(file_fs_s3_conf_proto_rawDescData)
	})
	return file_fs_s3_conf_proto_rawDescData
}

var file_fs_s3_conf_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_fs_s3_conf_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_fs_s3_conf_proto_goTypes = []interface{}{
	(ErrorReason)(0), // 0: fs.s3.ErrorReason
	(*Conf)(nil),     // 1: fs.s3.Conf
}
var file_fs_s3_conf_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_fs_s3_conf_proto_init() }
func file_fs_s3_conf_proto_init() {
	if File_fs_s3_conf_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_fs_s3_conf_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Conf); i {
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
			RawDescriptor: file_fs_s3_conf_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_fs_s3_conf_proto_goTypes,
		DependencyIndexes: file_fs_s3_conf_proto_depIdxs,
		EnumInfos:         file_fs_s3_conf_proto_enumTypes,
		MessageInfos:      file_fs_s3_conf_proto_msgTypes,
	}.Build()
	File_fs_s3_conf_proto = out.File
	file_fs_s3_conf_proto_rawDesc = nil
	file_fs_s3_conf_proto_goTypes = nil
	file_fs_s3_conf_proto_depIdxs = nil
}
