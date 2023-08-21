// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.21.12
// source: client/rabbitmq/conf.proto

package rabbitmq

import (
	_ "github.com/go-kratos/kratos/v2/errors"
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

type ErrorReason int32

const (
	ErrorReason_UNKOWN         ErrorReason = 0
	ErrorReason_DISCONNECT     ErrorReason = 1 // 断链接了
	ErrorReason_CHANNEL_CLOSED ErrorReason = 2 // channel 断了
)

// Enum value maps for ErrorReason.
var (
	ErrorReason_name = map[int32]string{
		0: "UNKOWN",
		1: "DISCONNECT",
		2: "CHANNEL_CLOSED",
	}
	ErrorReason_value = map[string]int32{
		"UNKOWN":         0,
		"DISCONNECT":     1,
		"CHANNEL_CLOSED": 2,
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
	return file_client_rabbitmq_conf_proto_enumTypes[0].Descriptor()
}

func (ErrorReason) Type() protoreflect.EnumType {
	return &file_client_rabbitmq_conf_proto_enumTypes[0]
}

func (x ErrorReason) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ErrorReason.Descriptor instead.
func (ErrorReason) EnumDescriptor() ([]byte, []int) {
	return file_client_rabbitmq_conf_proto_rawDescGZIP(), []int{0}
}

type RabbitConf struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	Host     string `protobuf:"bytes,3,opt,name=host,proto3" json:"host,omitempty"`
	Port     int32  `protobuf:"varint,4,opt,name=port,proto3" json:"port,omitempty"`
	VHost    string `protobuf:"bytes,5,opt,name=vHost,proto3" json:"vHost,omitempty"`
	Address  string `protobuf:"bytes,6,opt,name=address,proto3" json:"address,omitempty"`
}

func (x *RabbitConf) Reset() {
	*x = RabbitConf{}
	if protoimpl.UnsafeEnabled {
		mi := &file_client_rabbitmq_conf_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RabbitConf) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RabbitConf) ProtoMessage() {}

func (x *RabbitConf) ProtoReflect() protoreflect.Message {
	mi := &file_client_rabbitmq_conf_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RabbitConf.ProtoReflect.Descriptor instead.
func (*RabbitConf) Descriptor() ([]byte, []int) {
	return file_client_rabbitmq_conf_proto_rawDescGZIP(), []int{0}
}

func (x *RabbitConf) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *RabbitConf) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *RabbitConf) GetHost() string {
	if x != nil {
		return x.Host
	}
	return ""
}

func (x *RabbitConf) GetPort() int32 {
	if x != nil {
		return x.Port
	}
	return 0
}

func (x *RabbitConf) GetVHost() string {
	if x != nil {
		return x.VHost
	}
	return ""
}

func (x *RabbitConf) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

type ConsumerConf struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Consumer string `protobuf:"bytes,2,opt,name=consumer,proto3" json:"consumer,omitempty"`
}

func (x *ConsumerConf) Reset() {
	*x = ConsumerConf{}
	if protoimpl.UnsafeEnabled {
		mi := &file_client_rabbitmq_conf_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConsumerConf) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConsumerConf) ProtoMessage() {}

func (x *ConsumerConf) ProtoReflect() protoreflect.Message {
	mi := &file_client_rabbitmq_conf_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConsumerConf.ProtoReflect.Descriptor instead.
func (*ConsumerConf) Descriptor() ([]byte, []int) {
	return file_client_rabbitmq_conf_proto_rawDescGZIP(), []int{1}
}

func (x *ConsumerConf) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ConsumerConf) GetConsumer() string {
	if x != nil {
		return x.Consumer
	}
	return ""
}

type ListenerConf struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Rabbit *RabbitConf     `protobuf:"bytes,1,opt,name=rabbit,proto3" json:"rabbit,omitempty"`
	Queues []*ConsumerConf `protobuf:"bytes,2,rep,name=queues,proto3" json:"queues,omitempty"`
}

func (x *ListenerConf) Reset() {
	*x = ListenerConf{}
	if protoimpl.UnsafeEnabled {
		mi := &file_client_rabbitmq_conf_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListenerConf) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListenerConf) ProtoMessage() {}

func (x *ListenerConf) ProtoReflect() protoreflect.Message {
	mi := &file_client_rabbitmq_conf_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListenerConf.ProtoReflect.Descriptor instead.
func (*ListenerConf) Descriptor() ([]byte, []int) {
	return file_client_rabbitmq_conf_proto_rawDescGZIP(), []int{2}
}

func (x *ListenerConf) GetRabbit() *RabbitConf {
	if x != nil {
		return x.Rabbit
	}
	return nil
}

func (x *ListenerConf) GetQueues() []*ConsumerConf {
	if x != nil {
		return x.Queues
	}
	return nil
}

type ProducerConf struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Rabbit      *RabbitConf `protobuf:"bytes,1,opt,name=rabbit,proto3" json:"rabbit,omitempty"`
	ContentType string      `protobuf:"bytes,2,opt,name=contentType,proto3" json:"contentType,omitempty"`
}

func (x *ProducerConf) Reset() {
	*x = ProducerConf{}
	if protoimpl.UnsafeEnabled {
		mi := &file_client_rabbitmq_conf_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProducerConf) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProducerConf) ProtoMessage() {}

func (x *ProducerConf) ProtoReflect() protoreflect.Message {
	mi := &file_client_rabbitmq_conf_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProducerConf.ProtoReflect.Descriptor instead.
func (*ProducerConf) Descriptor() ([]byte, []int) {
	return file_client_rabbitmq_conf_proto_rawDescGZIP(), []int{3}
}

func (x *ProducerConf) GetRabbit() *RabbitConf {
	if x != nil {
		return x.Rabbit
	}
	return nil
}

func (x *ProducerConf) GetContentType() string {
	if x != nil {
		return x.ContentType
	}
	return ""
}

type ExchangeConf struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Exchange string `protobuf:"bytes,1,opt,name=exchange,proto3" json:"exchange,omitempty"`
	Key      string `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *ExchangeConf) Reset() {
	*x = ExchangeConf{}
	if protoimpl.UnsafeEnabled {
		mi := &file_client_rabbitmq_conf_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExchangeConf) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExchangeConf) ProtoMessage() {}

func (x *ExchangeConf) ProtoReflect() protoreflect.Message {
	mi := &file_client_rabbitmq_conf_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExchangeConf.ProtoReflect.Descriptor instead.
func (*ExchangeConf) Descriptor() ([]byte, []int) {
	return file_client_rabbitmq_conf_proto_rawDescGZIP(), []int{4}
}

func (x *ExchangeConf) GetExchange() string {
	if x != nil {
		return x.Exchange
	}
	return ""
}

func (x *ExchangeConf) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

type QueueConf struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *QueueConf) Reset() {
	*x = QueueConf{}
	if protoimpl.UnsafeEnabled {
		mi := &file_client_rabbitmq_conf_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueueConf) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueueConf) ProtoMessage() {}

func (x *QueueConf) ProtoReflect() protoreflect.Message {
	mi := &file_client_rabbitmq_conf_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueueConf.ProtoReflect.Descriptor instead.
func (*QueueConf) Descriptor() ([]byte, []int) {
	return file_client_rabbitmq_conf_proto_rawDescGZIP(), []int{5}
}

func (x *QueueConf) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type AdminExchangeConf struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name       string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Kind       string `protobuf:"bytes,2,opt,name=kind,proto3" json:"kind,omitempty"`
	Durable    bool   `protobuf:"varint,3,opt,name=durable,proto3" json:"durable,omitempty"`
	AutoDelete bool   `protobuf:"varint,4,opt,name=autoDelete,proto3" json:"autoDelete,omitempty"`
	Internal   bool   `protobuf:"varint,5,opt,name=internal,proto3" json:"internal,omitempty"`
	NoWait     bool   `protobuf:"varint,6,opt,name=noWait,proto3" json:"noWait,omitempty"`
}

func (x *AdminExchangeConf) Reset() {
	*x = AdminExchangeConf{}
	if protoimpl.UnsafeEnabled {
		mi := &file_client_rabbitmq_conf_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdminExchangeConf) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdminExchangeConf) ProtoMessage() {}

func (x *AdminExchangeConf) ProtoReflect() protoreflect.Message {
	mi := &file_client_rabbitmq_conf_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdminExchangeConf.ProtoReflect.Descriptor instead.
func (*AdminExchangeConf) Descriptor() ([]byte, []int) {
	return file_client_rabbitmq_conf_proto_rawDescGZIP(), []int{6}
}

func (x *AdminExchangeConf) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AdminExchangeConf) GetKind() string {
	if x != nil {
		return x.Kind
	}
	return ""
}

func (x *AdminExchangeConf) GetDurable() bool {
	if x != nil {
		return x.Durable
	}
	return false
}

func (x *AdminExchangeConf) GetAutoDelete() bool {
	if x != nil {
		return x.AutoDelete
	}
	return false
}

func (x *AdminExchangeConf) GetInternal() bool {
	if x != nil {
		return x.Internal
	}
	return false
}

func (x *AdminExchangeConf) GetNoWait() bool {
	if x != nil {
		return x.NoWait
	}
	return false
}

type AdminQueueConf struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name       string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Durable    bool   `protobuf:"varint,3,opt,name=durable,proto3" json:"durable,omitempty"`
	AutoDelete bool   `protobuf:"varint,4,opt,name=autoDelete,proto3" json:"autoDelete,omitempty"`
	Exclusive  bool   `protobuf:"varint,5,opt,name=exclusive,proto3" json:"exclusive,omitempty"`
	NoWait     bool   `protobuf:"varint,6,opt,name=noWait,proto3" json:"noWait,omitempty"`
}

func (x *AdminQueueConf) Reset() {
	*x = AdminQueueConf{}
	if protoimpl.UnsafeEnabled {
		mi := &file_client_rabbitmq_conf_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdminQueueConf) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdminQueueConf) ProtoMessage() {}

func (x *AdminQueueConf) ProtoReflect() protoreflect.Message {
	mi := &file_client_rabbitmq_conf_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdminQueueConf.ProtoReflect.Descriptor instead.
func (*AdminQueueConf) Descriptor() ([]byte, []int) {
	return file_client_rabbitmq_conf_proto_rawDescGZIP(), []int{7}
}

func (x *AdminQueueConf) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AdminQueueConf) GetDurable() bool {
	if x != nil {
		return x.Durable
	}
	return false
}

func (x *AdminQueueConf) GetAutoDelete() bool {
	if x != nil {
		return x.AutoDelete
	}
	return false
}

func (x *AdminQueueConf) GetExclusive() bool {
	if x != nil {
		return x.Exclusive
	}
	return false
}

func (x *AdminQueueConf) GetNoWait() bool {
	if x != nil {
		return x.NoWait
	}
	return false
}

var File_client_rabbitmq_conf_proto protoreflect.FileDescriptor

var file_client_rabbitmq_conf_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x72, 0x61, 0x62, 0x62, 0x69, 0x74, 0x6d,
	0x71, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x63, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x2e, 0x72, 0x61, 0x62, 0x62, 0x69, 0x74, 0x6d, 0x71, 0x1a, 0x13, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x73, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x9c, 0x01, 0x0a, 0x0a, 0x52, 0x61, 0x62, 0x62, 0x69, 0x74, 0x43, 0x6f, 0x6e,
	0x66, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a,
	0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x6f, 0x73,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x6f, 0x72,
	0x74, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x48, 0x6f, 0x73, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x76, 0x48, 0x6f, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x22, 0x3e, 0x0a, 0x0c, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x43, 0x6f, 0x6e,
	0x66, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65,
	0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65,
	0x72, 0x22, 0x7a, 0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x43, 0x6f, 0x6e,
	0x66, 0x12, 0x33, 0x0a, 0x06, 0x72, 0x61, 0x62, 0x62, 0x69, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1b, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x72, 0x61, 0x62, 0x62, 0x69,
	0x74, 0x6d, 0x71, 0x2e, 0x52, 0x61, 0x62, 0x62, 0x69, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x52, 0x06,
	0x72, 0x61, 0x62, 0x62, 0x69, 0x74, 0x12, 0x35, 0x0a, 0x06, 0x71, 0x75, 0x65, 0x75, 0x65, 0x73,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e,
	0x72, 0x61, 0x62, 0x62, 0x69, 0x74, 0x6d, 0x71, 0x2e, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65,
	0x72, 0x43, 0x6f, 0x6e, 0x66, 0x52, 0x06, 0x71, 0x75, 0x65, 0x75, 0x65, 0x73, 0x22, 0x65, 0x0a,
	0x0c, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x12, 0x33, 0x0a,
	0x06, 0x72, 0x61, 0x62, 0x62, 0x69, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x72, 0x61, 0x62, 0x62, 0x69, 0x74, 0x6d, 0x71, 0x2e,
	0x52, 0x61, 0x62, 0x62, 0x69, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x52, 0x06, 0x72, 0x61, 0x62, 0x62,
	0x69, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x54, 0x79, 0x70, 0x65, 0x22, 0x3c, 0x0a, 0x0c, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65,
	0x43, 0x6f, 0x6e, 0x66, 0x12, 0x1a, 0x0a, 0x08, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x22, 0x1f, 0x0a, 0x09, 0x51, 0x75, 0x65, 0x75, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x22, 0xa9, 0x01, 0x0a, 0x11, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x45, 0x78, 0x63,
	0x68, 0x61, 0x6e, 0x67, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x6b, 0x69, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6b, 0x69, 0x6e,
	0x64, 0x12, 0x18, 0x0a, 0x07, 0x64, 0x75, 0x72, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x07, 0x64, 0x75, 0x72, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x61,
	0x75, 0x74, 0x6f, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x0a, 0x61, 0x75, 0x74, 0x6f, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x6f, 0x57, 0x61, 0x69,
	0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x6e, 0x6f, 0x57, 0x61, 0x69, 0x74, 0x22,
	0x94, 0x01, 0x0a, 0x0e, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x51, 0x75, 0x65, 0x75, 0x65, 0x43, 0x6f,
	0x6e, 0x66, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x64, 0x75, 0x72, 0x61, 0x62, 0x6c,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x64, 0x75, 0x72, 0x61, 0x62, 0x6c, 0x65,
	0x12, 0x1e, 0x0a, 0x0a, 0x61, 0x75, 0x74, 0x6f, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x61, 0x75, 0x74, 0x6f, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x12, 0x1c, 0x0a, 0x09, 0x65, 0x78, 0x63, 0x6c, 0x75, 0x73, 0x69, 0x76, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x09, 0x65, 0x78, 0x63, 0x6c, 0x75, 0x73, 0x69, 0x76, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x6e, 0x6f, 0x57, 0x61, 0x69, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06,
	0x6e, 0x6f, 0x57, 0x61, 0x69, 0x74, 0x2a, 0x43, 0x0a, 0x0b, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52,
	0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x0a, 0x0a, 0x06, 0x55, 0x4e, 0x4b, 0x4f, 0x57, 0x4e, 0x10,
	0x00, 0x12, 0x0e, 0x0a, 0x0a, 0x44, 0x49, 0x53, 0x43, 0x4f, 0x4e, 0x4e, 0x45, 0x43, 0x54, 0x10,
	0x01, 0x12, 0x12, 0x0a, 0x0e, 0x43, 0x48, 0x41, 0x4e, 0x4e, 0x45, 0x4c, 0x5f, 0x43, 0x4c, 0x4f,
	0x53, 0x45, 0x44, 0x10, 0x02, 0x1a, 0x04, 0xa0, 0x45, 0xf4, 0x03, 0x42, 0x33, 0x5a, 0x31, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2d, 0x62, 0x61, 0x6d,
	0x62, 0x6f, 0x6f, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x72,
	0x61, 0x62, 0x62, 0x69, 0x74, 0x6d, 0x71, 0x3b, 0x72, 0x61, 0x62, 0x62, 0x69, 0x74, 0x6d, 0x71,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_client_rabbitmq_conf_proto_rawDescOnce sync.Once
	file_client_rabbitmq_conf_proto_rawDescData = file_client_rabbitmq_conf_proto_rawDesc
)

func file_client_rabbitmq_conf_proto_rawDescGZIP() []byte {
	file_client_rabbitmq_conf_proto_rawDescOnce.Do(func() {
		file_client_rabbitmq_conf_proto_rawDescData = protoimpl.X.CompressGZIP(file_client_rabbitmq_conf_proto_rawDescData)
	})
	return file_client_rabbitmq_conf_proto_rawDescData
}

var file_client_rabbitmq_conf_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_client_rabbitmq_conf_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_client_rabbitmq_conf_proto_goTypes = []interface{}{
	(ErrorReason)(0),          // 0: client.rabbitmq.ErrorReason
	(*RabbitConf)(nil),        // 1: client.rabbitmq.RabbitConf
	(*ConsumerConf)(nil),      // 2: client.rabbitmq.ConsumerConf
	(*ListenerConf)(nil),      // 3: client.rabbitmq.ListenerConf
	(*ProducerConf)(nil),      // 4: client.rabbitmq.ProducerConf
	(*ExchangeConf)(nil),      // 5: client.rabbitmq.ExchangeConf
	(*QueueConf)(nil),         // 6: client.rabbitmq.QueueConf
	(*AdminExchangeConf)(nil), // 7: client.rabbitmq.AdminExchangeConf
	(*AdminQueueConf)(nil),    // 8: client.rabbitmq.AdminQueueConf
}
var file_client_rabbitmq_conf_proto_depIdxs = []int32{
	1, // 0: client.rabbitmq.ListenerConf.rabbit:type_name -> client.rabbitmq.RabbitConf
	2, // 1: client.rabbitmq.ListenerConf.queues:type_name -> client.rabbitmq.ConsumerConf
	1, // 2: client.rabbitmq.ProducerConf.rabbit:type_name -> client.rabbitmq.RabbitConf
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_client_rabbitmq_conf_proto_init() }
func file_client_rabbitmq_conf_proto_init() {
	if File_client_rabbitmq_conf_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_client_rabbitmq_conf_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RabbitConf); i {
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
		file_client_rabbitmq_conf_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConsumerConf); i {
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
		file_client_rabbitmq_conf_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListenerConf); i {
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
		file_client_rabbitmq_conf_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProducerConf); i {
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
		file_client_rabbitmq_conf_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExchangeConf); i {
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
		file_client_rabbitmq_conf_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueueConf); i {
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
		file_client_rabbitmq_conf_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdminExchangeConf); i {
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
		file_client_rabbitmq_conf_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdminQueueConf); i {
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
			RawDescriptor: file_client_rabbitmq_conf_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_client_rabbitmq_conf_proto_goTypes,
		DependencyIndexes: file_client_rabbitmq_conf_proto_depIdxs,
		EnumInfos:         file_client_rabbitmq_conf_proto_enumTypes,
		MessageInfos:      file_client_rabbitmq_conf_proto_msgTypes,
	}.Build()
	File_client_rabbitmq_conf_proto = out.File
	file_client_rabbitmq_conf_proto_rawDesc = nil
	file_client_rabbitmq_conf_proto_goTypes = nil
	file_client_rabbitmq_conf_proto_depIdxs = nil
}
