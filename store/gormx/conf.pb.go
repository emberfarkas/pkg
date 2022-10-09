// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.7
// source: store/gormx/conf.proto

package gormx

import (
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

type Conf struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Driver   string `protobuf:"bytes,1,opt,name=driver,proto3" json:"driver,omitempty"`
	Source   string `protobuf:"bytes,2,opt,name=source,proto3" json:"source,omitempty"`
	LogLevel int32  `protobuf:"varint,3,opt,name=logLevel,proto3" json:"logLevel,omitempty"`
}

func (x *Conf) Reset() {
	*x = Conf{}
	if protoimpl.UnsafeEnabled {
		mi := &file_store_gormx_conf_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Conf) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Conf) ProtoMessage() {}

func (x *Conf) ProtoReflect() protoreflect.Message {
	mi := &file_store_gormx_conf_proto_msgTypes[0]
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
	return file_store_gormx_conf_proto_rawDescGZIP(), []int{0}
}

func (x *Conf) GetDriver() string {
	if x != nil {
		return x.Driver
	}
	return ""
}

func (x *Conf) GetSource() string {
	if x != nil {
		return x.Source
	}
	return ""
}

func (x *Conf) GetLogLevel() int32 {
	if x != nil {
		return x.LogLevel
	}
	return 0
}

var File_store_gormx_conf_proto protoreflect.FileDescriptor

var file_store_gormx_conf_proto_rawDesc = []byte{
	0x0a, 0x16, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x67, 0x6f, 0x72, 0x6d, 0x78, 0x2f, 0x63, 0x6f,
	0x6e, 0x66, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e,
	0x67, 0x6f, 0x72, 0x6d, 0x78, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x52, 0x0a, 0x04, 0x43, 0x6f, 0x6e, 0x66, 0x12, 0x16, 0x0a,
	0x06, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64,
	0x72, 0x69, 0x76, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x1a, 0x0a,
	0x08, 0x6c, 0x6f, 0x67, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x08, 0x6c, 0x6f, 0x67, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x42, 0x2e, 0x5a, 0x2c, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x66, 0x61, 0x72,
	0x6b, 0x61, 0x73, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x67, 0x6f,
	0x72, 0x6d, 0x78, 0x3b, 0x67, 0x6f, 0x72, 0x6d, 0x78, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_store_gormx_conf_proto_rawDescOnce sync.Once
	file_store_gormx_conf_proto_rawDescData = file_store_gormx_conf_proto_rawDesc
)

func file_store_gormx_conf_proto_rawDescGZIP() []byte {
	file_store_gormx_conf_proto_rawDescOnce.Do(func() {
		file_store_gormx_conf_proto_rawDescData = protoimpl.X.CompressGZIP(file_store_gormx_conf_proto_rawDescData)
	})
	return file_store_gormx_conf_proto_rawDescData
}

var file_store_gormx_conf_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_store_gormx_conf_proto_goTypes = []interface{}{
	(*Conf)(nil), // 0: store.gormx.Conf
}
var file_store_gormx_conf_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_store_gormx_conf_proto_init() }
func file_store_gormx_conf_proto_init() {
	if File_store_gormx_conf_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_store_gormx_conf_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
			RawDescriptor: file_store_gormx_conf_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_store_gormx_conf_proto_goTypes,
		DependencyIndexes: file_store_gormx_conf_proto_depIdxs,
		MessageInfos:      file_store_gormx_conf_proto_msgTypes,
	}.Build()
	File_store_gormx_conf_proto = out.File
	file_store_gormx_conf_proto_rawDesc = nil
	file_store_gormx_conf_proto_goTypes = nil
	file_store_gormx_conf_proto_depIdxs = nil
}
