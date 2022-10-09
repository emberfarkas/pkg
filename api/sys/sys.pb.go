// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.7
// source: api/sys/sys.proto

package sys

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

type CheckResourceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Path   string `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	Method string `protobuf:"bytes,2,opt,name=method,proto3" json:"method,omitempty"`
}

func (x *CheckResourceRequest) Reset() {
	*x = CheckResourceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_sys_sys_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckResourceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckResourceRequest) ProtoMessage() {}

func (x *CheckResourceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_sys_sys_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckResourceRequest.ProtoReflect.Descriptor instead.
func (*CheckResourceRequest) Descriptor() ([]byte, []int) {
	return file_api_sys_sys_proto_rawDescGZIP(), []int{0}
}

func (x *CheckResourceRequest) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *CheckResourceRequest) GetMethod() string {
	if x != nil {
		return x.Method
	}
	return ""
}

type CheckResourceReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ok bool `protobuf:"varint,1,opt,name=ok,proto3" json:"ok,omitempty"`
}

func (x *CheckResourceReply) Reset() {
	*x = CheckResourceReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_sys_sys_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckResourceReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckResourceReply) ProtoMessage() {}

func (x *CheckResourceReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_sys_sys_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckResourceReply.ProtoReflect.Descriptor instead.
func (*CheckResourceReply) Descriptor() ([]byte, []int) {
	return file_api_sys_sys_proto_rawDescGZIP(), []int{1}
}

func (x *CheckResourceReply) GetOk() bool {
	if x != nil {
		return x.Ok
	}
	return false
}

type AuthRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccessToken string `protobuf:"bytes,1,opt,name=AccessToken,proto3" json:"AccessToken,omitempty"`
}

func (x *AuthRequest) Reset() {
	*x = AuthRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_sys_sys_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthRequest) ProtoMessage() {}

func (x *AuthRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_sys_sys_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthRequest.ProtoReflect.Descriptor instead.
func (*AuthRequest) Descriptor() ([]byte, []int) {
	return file_api_sys_sys_proto_rawDescGZIP(), []int{2}
}

func (x *AuthRequest) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

type AuthResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DataScope string `protobuf:"bytes,1,opt,name=dataScope,proto3" json:"dataScope,omitempty"`
	UserId    uint64 `protobuf:"varint,2,opt,name=userId,proto3" json:"userId,omitempty"`
	RoleId    uint64 `protobuf:"varint,3,opt,name=roleId,proto3" json:"roleId,omitempty"`
	RoleKey   string `protobuf:"bytes,4,opt,name=roleKey,proto3" json:"roleKey,omitempty"`
}

func (x *AuthResp) Reset() {
	*x = AuthResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_sys_sys_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthResp) ProtoMessage() {}

func (x *AuthResp) ProtoReflect() protoreflect.Message {
	mi := &file_api_sys_sys_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthResp.ProtoReflect.Descriptor instead.
func (*AuthResp) Descriptor() ([]byte, []int) {
	return file_api_sys_sys_proto_rawDescGZIP(), []int{3}
}

func (x *AuthResp) GetDataScope() string {
	if x != nil {
		return x.DataScope
	}
	return ""
}

func (x *AuthResp) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *AuthResp) GetRoleId() uint64 {
	if x != nil {
		return x.RoleId
	}
	return 0
}

func (x *AuthResp) GetRoleKey() string {
	if x != nil {
		return x.RoleKey
	}
	return ""
}

type CheckRoleLevelRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoleId int32 `protobuf:"varint,1,opt,name=roleId,proto3" json:"roleId,omitempty"`
	Target int32 `protobuf:"varint,2,opt,name=target,proto3" json:"target,omitempty"`
}

func (x *CheckRoleLevelRequest) Reset() {
	*x = CheckRoleLevelRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_sys_sys_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckRoleLevelRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckRoleLevelRequest) ProtoMessage() {}

func (x *CheckRoleLevelRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_sys_sys_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckRoleLevelRequest.ProtoReflect.Descriptor instead.
func (*CheckRoleLevelRequest) Descriptor() ([]byte, []int) {
	return file_api_sys_sys_proto_rawDescGZIP(), []int{4}
}

func (x *CheckRoleLevelRequest) GetRoleId() int32 {
	if x != nil {
		return x.RoleId
	}
	return 0
}

func (x *CheckRoleLevelRequest) GetTarget() int32 {
	if x != nil {
		return x.Target
	}
	return 0
}

type CheckRoleLevelReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CheckRoleLevelReply) Reset() {
	*x = CheckRoleLevelReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_sys_sys_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckRoleLevelReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckRoleLevelReply) ProtoMessage() {}

func (x *CheckRoleLevelReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_sys_sys_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckRoleLevelReply.ProtoReflect.Descriptor instead.
func (*CheckRoleLevelReply) Descriptor() ([]byte, []int) {
	return file_api_sys_sys_proto_rawDescGZIP(), []int{5}
}

var File_api_sys_sys_proto protoreflect.FileDescriptor

var file_api_sys_sys_proto_rawDesc = []byte{
	0x0a, 0x11, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x79, 0x73, 0x2f, 0x73, 0x79, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x07, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x79, 0x73, 0x1a, 0x1c, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e,
	0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x2f, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x54, 0x0a, 0x14, 0x43, 0x68, 0x65,
	0x63, 0x6b, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1b, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x02, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x1f,
	0x0a, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07,
	0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x02, 0x52, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x22,
	0x24, 0x0a, 0x12, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x6f, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x02, 0x6f, 0x6b, 0x22, 0x38, 0x0a, 0x0b, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x29, 0x0a, 0x0b, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02,
	0x10, 0x01, 0x52, 0x0b, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22,
	0x72, 0x0a, 0x08, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x73, 0x70, 0x12, 0x1c, 0x0a, 0x09, 0x64,
	0x61, 0x74, 0x61, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x64, 0x61, 0x74, 0x61, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x6f, 0x6c, 0x65, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x06, 0x72, 0x6f, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x6f, 0x6c,
	0x65, 0x4b, 0x65, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x72, 0x6f, 0x6c, 0x65,
	0x4b, 0x65, 0x79, 0x22, 0x47, 0x0a, 0x15, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x6f, 0x6c, 0x65,
	0x4c, 0x65, 0x76, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06,
	0x72, 0x6f, 0x6c, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x72, 0x6f,
	0x6c, 0x65, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x22, 0x15, 0x0a, 0x13,
	0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x6f, 0x6c, 0x65, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x32, 0xac, 0x02, 0x0a, 0x03, 0x53, 0x79, 0x73, 0x12, 0x58, 0x0a, 0x04, 0x41,
	0x75, 0x74, 0x68, 0x12, 0x14, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x79, 0x73, 0x2e, 0x41, 0x75,
	0x74, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x73, 0x79, 0x73, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x73, 0x70, 0x22, 0x27, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x16, 0x12, 0x14, 0x2f, 0x78, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61,
	0x6c, 0x2f, 0x73, 0x79, 0x73, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x92, 0x41, 0x08, 0x12, 0x06, 0xe6,
	0x8e, 0x88, 0xe6, 0x9d, 0x83, 0x12, 0x7b, 0x0a, 0x0d, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x1d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x79, 0x73,
	0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x79, 0x73, 0x2e,
	0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x22, 0x2e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1d, 0x12, 0x1b, 0x2f, 0x78, 0x2f, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x73, 0x79, 0x73, 0x2f, 0x63, 0x68, 0x65, 0x63,
	0x6b, 0x63, 0x61, 0x73, 0x62, 0x69, 0x6e, 0x92, 0x41, 0x08, 0x12, 0x06, 0xe6, 0x8e, 0x88, 0xe6,
	0x9d, 0x83, 0x12, 0x4e, 0x0a, 0x0e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x6f, 0x6c, 0x65, 0x4c,
	0x65, 0x76, 0x65, 0x6c, 0x12, 0x1e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x79, 0x73, 0x2e, 0x43,
	0x68, 0x65, 0x63, 0x6b, 0x52, 0x6f, 0x6c, 0x65, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x79, 0x73, 0x2e, 0x43,
	0x68, 0x65, 0x63, 0x6b, 0x52, 0x6f, 0x6c, 0x65, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x42, 0x98, 0x03, 0x5a, 0x26, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x66, 0x61, 0x72, 0x6b, 0x61, 0x73, 0x2f, 0x70, 0x6b,
	0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x79, 0x73, 0x3b, 0x73, 0x79, 0x73, 0x92, 0x41, 0xec,
	0x02, 0x12, 0x10, 0x0a, 0x09, 0x73, 0x73, 0x6f, 0xe6, 0x9c, 0x8d, 0xe5, 0x8a, 0xa1, 0x32, 0x03,
	0x31, 0x2e, 0x30, 0x2a, 0x03, 0x01, 0x02, 0x04, 0x32, 0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x32, 0x16, 0x61, 0x70, 0x70, 0x6c,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x78, 0x2d, 0x66, 0x6f, 0x6f, 0x2d, 0x6d, 0x69,
	0x6d, 0x65, 0x3a, 0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f,
	0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x16, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2f, 0x78, 0x2d, 0x66, 0x6f, 0x6f, 0x2d, 0x6d, 0x69, 0x6d, 0x65, 0x52, 0xeb, 0x01, 0x0a,
	0x03, 0x35, 0x30, 0x30, 0x12, 0xe3, 0x01, 0x0a, 0x0c, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x20,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x11, 0x0a, 0x0f, 0x1a, 0x0d, 0x2e, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x73, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x1a, 0xbf, 0x01, 0x0a, 0x10, 0x58, 0x2d, 0x43,
	0x6f, 0x72, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2d, 0x49, 0x64, 0x12, 0xaa, 0x01,
	0x0a, 0x2b, 0x55, 0x6e, 0x69, 0x71, 0x75, 0x65, 0x20, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x20, 0x69,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x20, 0x66, 0x6f, 0x72, 0x20, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x20, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x12, 0x06, 0x73,
	0x74, 0x72, 0x69, 0x6e, 0x67, 0x1a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x32, 0x26, 0x22, 0x32, 0x34,
	0x33, 0x38, 0x61, 0x63, 0x33, 0x63, 0x2d, 0x33, 0x37, 0x65, 0x62, 0x2d, 0x34, 0x39, 0x30, 0x32,
	0x2d, 0x61, 0x64, 0x65, 0x66, 0x2d, 0x65, 0x64, 0x31, 0x36, 0x62, 0x34, 0x34, 0x33, 0x31, 0x30,
	0x33, 0x30, 0x22, 0x6a, 0x45, 0x5e, 0x5b, 0x30, 0x2d, 0x39, 0x41, 0x2d, 0x46, 0x5d, 0x7b, 0x38,
	0x7d, 0x2d, 0x5b, 0x30, 0x2d, 0x39, 0x41, 0x2d, 0x46, 0x5d, 0x7b, 0x34, 0x7d, 0x2d, 0x34, 0x5b,
	0x30, 0x2d, 0x39, 0x41, 0x2d, 0x46, 0x5d, 0x7b, 0x33, 0x7d, 0x2d, 0x5b, 0x38, 0x39, 0x41, 0x42,
	0x5d, 0x5b, 0x30, 0x2d, 0x39, 0x41, 0x2d, 0x46, 0x5d, 0x7b, 0x33, 0x7d, 0x2d, 0x5b, 0x30, 0x2d,
	0x39, 0x41, 0x2d, 0x46, 0x5d, 0x7b, 0x31, 0x32, 0x7d, 0x24, 0x5a, 0x11, 0x0a, 0x0f, 0x0a, 0x09,
	0x42, 0x61, 0x73, 0x69, 0x63, 0x41, 0x75, 0x74, 0x68, 0x12, 0x02, 0x08, 0x01, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_sys_sys_proto_rawDescOnce sync.Once
	file_api_sys_sys_proto_rawDescData = file_api_sys_sys_proto_rawDesc
)

func file_api_sys_sys_proto_rawDescGZIP() []byte {
	file_api_sys_sys_proto_rawDescOnce.Do(func() {
		file_api_sys_sys_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_sys_sys_proto_rawDescData)
	})
	return file_api_sys_sys_proto_rawDescData
}

var file_api_sys_sys_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_api_sys_sys_proto_goTypes = []interface{}{
	(*CheckResourceRequest)(nil),  // 0: api.sys.CheckResourceRequest
	(*CheckResourceReply)(nil),    // 1: api.sys.CheckResourceReply
	(*AuthRequest)(nil),           // 2: api.sys.AuthRequest
	(*AuthResp)(nil),              // 3: api.sys.AuthResp
	(*CheckRoleLevelRequest)(nil), // 4: api.sys.CheckRoleLevelRequest
	(*CheckRoleLevelReply)(nil),   // 5: api.sys.CheckRoleLevelReply
}
var file_api_sys_sys_proto_depIdxs = []int32{
	2, // 0: api.sys.Sys.Auth:input_type -> api.sys.AuthRequest
	0, // 1: api.sys.Sys.CheckResource:input_type -> api.sys.CheckResourceRequest
	4, // 2: api.sys.Sys.CheckRoleLevel:input_type -> api.sys.CheckRoleLevelRequest
	3, // 3: api.sys.Sys.Auth:output_type -> api.sys.AuthResp
	1, // 4: api.sys.Sys.CheckResource:output_type -> api.sys.CheckResourceReply
	5, // 5: api.sys.Sys.CheckRoleLevel:output_type -> api.sys.CheckRoleLevelReply
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_sys_sys_proto_init() }
func file_api_sys_sys_proto_init() {
	if File_api_sys_sys_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_sys_sys_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckResourceRequest); i {
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
		file_api_sys_sys_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckResourceReply); i {
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
		file_api_sys_sys_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthRequest); i {
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
		file_api_sys_sys_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthResp); i {
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
		file_api_sys_sys_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckRoleLevelRequest); i {
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
		file_api_sys_sys_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckRoleLevelReply); i {
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
			RawDescriptor: file_api_sys_sys_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_sys_sys_proto_goTypes,
		DependencyIndexes: file_api_sys_sys_proto_depIdxs,
		MessageInfos:      file_api_sys_sys_proto_msgTypes,
	}.Build()
	File_api_sys_sys_proto = out.File
	file_api_sys_sys_proto_rawDesc = nil
	file_api_sys_sys_proto_goTypes = nil
	file_api_sys_sys_proto_depIdxs = nil
}
