// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.21.12
// source: log/conf.proto

package log

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

type CoreType int32

const (
	CoreType_Unkown     CoreType = 0
	CoreType_Stdout     CoreType = 1
	CoreType_File       CoreType = 2
	CoreType_Fluent     CoreType = 3
	CoreType_CloudWatch CoreType = 4
	CoreType_AliYun     CoreType = 5
	CoreType_Sys        CoreType = 6
	CoreType_Hc         CoreType = 7
)

// Enum value maps for CoreType.
var (
	CoreType_name = map[int32]string{
		0: "Unkown",
		1: "Stdout",
		2: "File",
		3: "Fluent",
		4: "CloudWatch",
		5: "AliYun",
		6: "Sys",
		7: "Hc",
	}
	CoreType_value = map[string]int32{
		"Unkown":     0,
		"Stdout":     1,
		"File":       2,
		"Fluent":     3,
		"CloudWatch": 4,
		"AliYun":     5,
		"Sys":        6,
		"Hc":         7,
	}
)

func (x CoreType) Enum() *CoreType {
	p := new(CoreType)
	*p = x
	return p
}

func (x CoreType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CoreType) Descriptor() protoreflect.EnumDescriptor {
	return file_log_conf_proto_enumTypes[0].Descriptor()
}

func (CoreType) Type() protoreflect.EnumType {
	return &file_log_conf_proto_enumTypes[0]
}

func (x CoreType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CoreType.Descriptor instead.
func (CoreType) EnumDescriptor() ([]byte, []int) {
	return file_log_conf_proto_rawDescGZIP(), []int{0}
}

type Conf struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type         CoreType `protobuf:"varint,1,opt,name=type,proto3,enum=log.CoreType" json:"type,omitempty"`
	Level        int32    `protobuf:"varint,2,opt,name=level,proto3" json:"level,omitempty"`
	Path         string   `protobuf:"bytes,3,opt,name=path,proto3" json:"path,omitempty"`
	Endpoint     string   `protobuf:"bytes,4,opt,name=endpoint,proto3" json:"endpoint,omitempty"`
	Key          string   `protobuf:"bytes,5,opt,name=key,proto3" json:"key,omitempty"`
	Secret       string   `protobuf:"bytes,6,opt,name=secret,proto3" json:"secret,omitempty"`
	Region       string   `protobuf:"bytes,7,opt,name=region,proto3" json:"region,omitempty"`
	Profile      string   `protobuf:"bytes,8,opt,name=profile,proto3" json:"profile,omitempty"`
	LogGroupName string   `protobuf:"bytes,9,opt,name=logGroupName,proto3" json:"logGroupName,omitempty"`
}

func (x *Conf) Reset() {
	*x = Conf{}
	if protoimpl.UnsafeEnabled {
		mi := &file_log_conf_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Conf) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Conf) ProtoMessage() {}

func (x *Conf) ProtoReflect() protoreflect.Message {
	mi := &file_log_conf_proto_msgTypes[0]
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
	return file_log_conf_proto_rawDescGZIP(), []int{0}
}

func (x *Conf) GetType() CoreType {
	if x != nil {
		return x.Type
	}
	return CoreType_Unkown
}

func (x *Conf) GetLevel() int32 {
	if x != nil {
		return x.Level
	}
	return 0
}

func (x *Conf) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *Conf) GetEndpoint() string {
	if x != nil {
		return x.Endpoint
	}
	return ""
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

func (x *Conf) GetProfile() string {
	if x != nil {
		return x.Profile
	}
	return ""
}

func (x *Conf) GetLogGroupName() string {
	if x != nil {
		return x.LogGroupName
	}
	return ""
}

var File_log_conf_proto protoreflect.FileDescriptor

var file_log_conf_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x6c, 0x6f, 0x67, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x03, 0x6c, 0x6f, 0x67, 0x22, 0xef, 0x01, 0x0a, 0x04, 0x43, 0x6f, 0x6e, 0x66, 0x12, 0x21,
	0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0d, 0x2e, 0x6c,
	0x6f, 0x67, 0x2e, 0x43, 0x6f, 0x72, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x1a, 0x0a, 0x08, 0x65,
	0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x65,
	0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x63,
	0x72, 0x65, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x65, 0x63, 0x72, 0x65,
	0x74, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x72, 0x6f,
	0x66, 0x69, 0x6c, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x66,
	0x69, 0x6c, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x6c, 0x6f, 0x67, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4e,
	0x61, 0x6d, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6c, 0x6f, 0x67, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x4e, 0x61, 0x6d, 0x65, 0x2a, 0x65, 0x0a, 0x08, 0x43, 0x6f, 0x72, 0x65, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x0a, 0x0a, 0x06, 0x55, 0x6e, 0x6b, 0x6f, 0x77, 0x6e, 0x10, 0x00, 0x12,
	0x0a, 0x0a, 0x06, 0x53, 0x74, 0x64, 0x6f, 0x75, 0x74, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x46,
	0x69, 0x6c, 0x65, 0x10, 0x02, 0x12, 0x0a, 0x0a, 0x06, 0x46, 0x6c, 0x75, 0x65, 0x6e, 0x74, 0x10,
	0x03, 0x12, 0x0e, 0x0a, 0x0a, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x57, 0x61, 0x74, 0x63, 0x68, 0x10,
	0x04, 0x12, 0x0a, 0x0a, 0x06, 0x41, 0x6c, 0x69, 0x59, 0x75, 0x6e, 0x10, 0x05, 0x12, 0x07, 0x0a,
	0x03, 0x53, 0x79, 0x73, 0x10, 0x06, 0x12, 0x06, 0x0a, 0x02, 0x48, 0x63, 0x10, 0x07, 0x42, 0x22,
	0x5a, 0x20, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2d,
	0x62, 0x61, 0x6d, 0x62, 0x6f, 0x6f, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x6c, 0x6f, 0x67, 0x3b, 0x6c,
	0x6f, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_log_conf_proto_rawDescOnce sync.Once
	file_log_conf_proto_rawDescData = file_log_conf_proto_rawDesc
)

func file_log_conf_proto_rawDescGZIP() []byte {
	file_log_conf_proto_rawDescOnce.Do(func() {
		file_log_conf_proto_rawDescData = protoimpl.X.CompressGZIP(file_log_conf_proto_rawDescData)
	})
	return file_log_conf_proto_rawDescData
}

var file_log_conf_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_log_conf_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_log_conf_proto_goTypes = []interface{}{
	(CoreType)(0), // 0: log.CoreType
	(*Conf)(nil),  // 1: log.Conf
}
var file_log_conf_proto_depIdxs = []int32{
	0, // 0: log.Conf.type:type_name -> log.CoreType
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_log_conf_proto_init() }
func file_log_conf_proto_init() {
	if File_log_conf_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_log_conf_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
			RawDescriptor: file_log_conf_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_log_conf_proto_goTypes,
		DependencyIndexes: file_log_conf_proto_depIdxs,
		EnumInfos:         file_log_conf_proto_enumTypes,
		MessageInfos:      file_log_conf_proto_msgTypes,
	}.Build()
	File_log_conf_proto = out.File
	file_log_conf_proto_rawDesc = nil
	file_log_conf_proto_goTypes = nil
	file_log_conf_proto_depIdxs = nil
}
