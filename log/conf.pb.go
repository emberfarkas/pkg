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

type StdoutConf struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Enable bool  `protobuf:"varint,1,opt,name=enable,proto3" json:"enable,omitempty"`
	Level  int32 `protobuf:"varint,2,opt,name=level,proto3" json:"level,omitempty"`
}

func (x *StdoutConf) Reset() {
	*x = StdoutConf{}
	if protoimpl.UnsafeEnabled {
		mi := &file_log_conf_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StdoutConf) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StdoutConf) ProtoMessage() {}

func (x *StdoutConf) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use StdoutConf.ProtoReflect.Descriptor instead.
func (*StdoutConf) Descriptor() ([]byte, []int) {
	return file_log_conf_proto_rawDescGZIP(), []int{0}
}

func (x *StdoutConf) GetEnable() bool {
	if x != nil {
		return x.Enable
	}
	return false
}

func (x *StdoutConf) GetLevel() int32 {
	if x != nil {
		return x.Level
	}
	return 0
}

type FileConf struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Enable bool   `protobuf:"varint,1,opt,name=enable,proto3" json:"enable,omitempty"`
	Level  int32  `protobuf:"varint,2,opt,name=level,proto3" json:"level,omitempty"`
	Path   string `protobuf:"bytes,3,opt,name=path,proto3" json:"path,omitempty"`
	Name   string `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *FileConf) Reset() {
	*x = FileConf{}
	if protoimpl.UnsafeEnabled {
		mi := &file_log_conf_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileConf) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileConf) ProtoMessage() {}

func (x *FileConf) ProtoReflect() protoreflect.Message {
	mi := &file_log_conf_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileConf.ProtoReflect.Descriptor instead.
func (*FileConf) Descriptor() ([]byte, []int) {
	return file_log_conf_proto_rawDescGZIP(), []int{1}
}

func (x *FileConf) GetEnable() bool {
	if x != nil {
		return x.Enable
	}
	return false
}

func (x *FileConf) GetLevel() int32 {
	if x != nil {
		return x.Level
	}
	return 0
}

func (x *FileConf) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *FileConf) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type FluentConf struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Enable bool   `protobuf:"varint,1,opt,name=enable,proto3" json:"enable,omitempty"`
	Level  int32  `protobuf:"varint,2,opt,name=level,proto3" json:"level,omitempty"`
	Addr   string `protobuf:"bytes,3,opt,name=addr,proto3" json:"addr,omitempty"`
}

func (x *FluentConf) Reset() {
	*x = FluentConf{}
	if protoimpl.UnsafeEnabled {
		mi := &file_log_conf_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FluentConf) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FluentConf) ProtoMessage() {}

func (x *FluentConf) ProtoReflect() protoreflect.Message {
	mi := &file_log_conf_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FluentConf.ProtoReflect.Descriptor instead.
func (*FluentConf) Descriptor() ([]byte, []int) {
	return file_log_conf_proto_rawDescGZIP(), []int{2}
}

func (x *FluentConf) GetEnable() bool {
	if x != nil {
		return x.Enable
	}
	return false
}

func (x *FluentConf) GetLevel() int32 {
	if x != nil {
		return x.Level
	}
	return 0
}

func (x *FluentConf) GetAddr() string {
	if x != nil {
		return x.Addr
	}
	return ""
}

type CloudWatchConf struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Enable       bool   `protobuf:"varint,1,opt,name=enable,proto3" json:"enable,omitempty"`
	Level        int32  `protobuf:"varint,2,opt,name=level,proto3" json:"level,omitempty"`
	Key          string `protobuf:"bytes,3,opt,name=key,proto3" json:"key,omitempty"`
	Secret       string `protobuf:"bytes,4,opt,name=secret,proto3" json:"secret,omitempty"`
	Region       string `protobuf:"bytes,5,opt,name=region,proto3" json:"region,omitempty"`
	Profile      string `protobuf:"bytes,6,opt,name=profile,proto3" json:"profile,omitempty"`
	LogGroupName string `protobuf:"bytes,7,opt,name=logGroupName,proto3" json:"logGroupName,omitempty"`
}

func (x *CloudWatchConf) Reset() {
	*x = CloudWatchConf{}
	if protoimpl.UnsafeEnabled {
		mi := &file_log_conf_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CloudWatchConf) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CloudWatchConf) ProtoMessage() {}

func (x *CloudWatchConf) ProtoReflect() protoreflect.Message {
	mi := &file_log_conf_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CloudWatchConf.ProtoReflect.Descriptor instead.
func (*CloudWatchConf) Descriptor() ([]byte, []int) {
	return file_log_conf_proto_rawDescGZIP(), []int{3}
}

func (x *CloudWatchConf) GetEnable() bool {
	if x != nil {
		return x.Enable
	}
	return false
}

func (x *CloudWatchConf) GetLevel() int32 {
	if x != nil {
		return x.Level
	}
	return 0
}

func (x *CloudWatchConf) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *CloudWatchConf) GetSecret() string {
	if x != nil {
		return x.Secret
	}
	return ""
}

func (x *CloudWatchConf) GetRegion() string {
	if x != nil {
		return x.Region
	}
	return ""
}

func (x *CloudWatchConf) GetProfile() string {
	if x != nil {
		return x.Profile
	}
	return ""
}

func (x *CloudWatchConf) GetLogGroupName() string {
	if x != nil {
		return x.LogGroupName
	}
	return ""
}

type Conf struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Console    *StdoutConf     `protobuf:"bytes,1,opt,name=console,proto3" json:"console,omitempty"`
	File       *FileConf       `protobuf:"bytes,2,opt,name=file,proto3" json:"file,omitempty"`
	Fluent     *FluentConf     `protobuf:"bytes,3,opt,name=fluent,proto3" json:"fluent,omitempty"`
	CloudWatch *CloudWatchConf `protobuf:"bytes,4,opt,name=cloudWatch,proto3" json:"cloudWatch,omitempty"`
}

func (x *Conf) Reset() {
	*x = Conf{}
	if protoimpl.UnsafeEnabled {
		mi := &file_log_conf_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Conf) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Conf) ProtoMessage() {}

func (x *Conf) ProtoReflect() protoreflect.Message {
	mi := &file_log_conf_proto_msgTypes[4]
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
	return file_log_conf_proto_rawDescGZIP(), []int{4}
}

func (x *Conf) GetConsole() *StdoutConf {
	if x != nil {
		return x.Console
	}
	return nil
}

func (x *Conf) GetFile() *FileConf {
	if x != nil {
		return x.File
	}
	return nil
}

func (x *Conf) GetFluent() *FluentConf {
	if x != nil {
		return x.Fluent
	}
	return nil
}

func (x *Conf) GetCloudWatch() *CloudWatchConf {
	if x != nil {
		return x.CloudWatch
	}
	return nil
}

var File_log_conf_proto protoreflect.FileDescriptor

var file_log_conf_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x6c, 0x6f, 0x67, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x03, 0x6c, 0x6f, 0x67, 0x22, 0x3a, 0x0a, 0x0a, 0x53, 0x74, 0x64, 0x6f, 0x75, 0x74, 0x43,
	0x6f, 0x6e, 0x66, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x06, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6c,
	0x65, 0x76, 0x65, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x65, 0x76, 0x65,
	0x6c, 0x22, 0x60, 0x0a, 0x08, 0x46, 0x69, 0x6c, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x12, 0x16, 0x0a,
	0x06, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x65,
	0x6e, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x70,
	0x61, 0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x22, 0x4e, 0x0a, 0x0a, 0x46, 0x6c, 0x75, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x6e,
	0x66, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x06, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x65, 0x76,
	0x65, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x12,
	0x12, 0x0a, 0x04, 0x61, 0x64, 0x64, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x61,
	0x64, 0x64, 0x72, 0x22, 0xbe, 0x01, 0x0a, 0x0e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x57, 0x61, 0x74,
	0x63, 0x68, 0x43, 0x6f, 0x6e, 0x66, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c,
	0x65, 0x76, 0x65, 0x6c, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x12, 0x16,
	0x0a, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c,
	0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65,
	0x12, 0x22, 0x0a, 0x0c, 0x6c, 0x6f, 0x67, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4e, 0x61, 0x6d, 0x65,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6c, 0x6f, 0x67, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x4e, 0x61, 0x6d, 0x65, 0x22, 0xb2, 0x01, 0x0a, 0x04, 0x43, 0x6f, 0x6e, 0x66, 0x12, 0x29, 0x0a,
	0x07, 0x63, 0x6f, 0x6e, 0x73, 0x6f, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f,
	0x2e, 0x6c, 0x6f, 0x67, 0x2e, 0x53, 0x74, 0x64, 0x6f, 0x75, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x52,
	0x07, 0x63, 0x6f, 0x6e, 0x73, 0x6f, 0x6c, 0x65, 0x12, 0x21, 0x0a, 0x04, 0x66, 0x69, 0x6c, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x6c, 0x6f, 0x67, 0x2e, 0x46, 0x69, 0x6c,
	0x65, 0x43, 0x6f, 0x6e, 0x66, 0x52, 0x04, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x27, 0x0a, 0x06, 0x66,
	0x6c, 0x75, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x6c, 0x6f,
	0x67, 0x2e, 0x46, 0x6c, 0x75, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x52, 0x06, 0x66, 0x6c,
	0x75, 0x65, 0x6e, 0x74, 0x12, 0x33, 0x0a, 0x0a, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x57, 0x61, 0x74,
	0x63, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x6c, 0x6f, 0x67, 0x2e, 0x43,
	0x6c, 0x6f, 0x75, 0x64, 0x57, 0x61, 0x74, 0x63, 0x68, 0x43, 0x6f, 0x6e, 0x66, 0x52, 0x0a, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x57, 0x61, 0x74, 0x63, 0x68, 0x42, 0x22, 0x5a, 0x20, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2d, 0x62, 0x61, 0x6d, 0x62, 0x6f,
	0x6f, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x6c, 0x6f, 0x67, 0x3b, 0x6c, 0x6f, 0x67, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
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

var file_log_conf_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_log_conf_proto_goTypes = []interface{}{
	(*StdoutConf)(nil),     // 0: log.StdoutConf
	(*FileConf)(nil),       // 1: log.FileConf
	(*FluentConf)(nil),     // 2: log.FluentConf
	(*CloudWatchConf)(nil), // 3: log.CloudWatchConf
	(*Conf)(nil),           // 4: log.Conf
}
var file_log_conf_proto_depIdxs = []int32{
	0, // 0: log.Conf.console:type_name -> log.StdoutConf
	1, // 1: log.Conf.file:type_name -> log.FileConf
	2, // 2: log.Conf.fluent:type_name -> log.FluentConf
	3, // 3: log.Conf.cloudWatch:type_name -> log.CloudWatchConf
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_log_conf_proto_init() }
func file_log_conf_proto_init() {
	if File_log_conf_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_log_conf_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StdoutConf); i {
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
		file_log_conf_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileConf); i {
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
		file_log_conf_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FluentConf); i {
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
		file_log_conf_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CloudWatchConf); i {
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
		file_log_conf_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
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
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_log_conf_proto_goTypes,
		DependencyIndexes: file_log_conf_proto_depIdxs,
		MessageInfos:      file_log_conf_proto_msgTypes,
	}.Build()
	File_log_conf_proto = out.File
	file_log_conf_proto_rawDesc = nil
	file_log_conf_proto_goTypes = nil
	file_log_conf_proto_depIdxs = nil
}
