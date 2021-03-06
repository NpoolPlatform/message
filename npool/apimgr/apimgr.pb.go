// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.18.1
// source: npool/apimgr/apimgr.proto

package apimgr

import (
	npool "github.com/NpoolPlatform/message/npool"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Path struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Method     string `protobuf:"bytes,10,opt,name=Method,proto3" json:"Method,omitempty"`
	Path       string `protobuf:"bytes,20,opt,name=Path,proto3" json:"Path,omitempty"`
	Exported   bool   `protobuf:"varint,30,opt,name=Exported,proto3" json:"Exported,omitempty"`
	MethodName string `protobuf:"bytes,40,opt,name=MethodName,proto3" json:"MethodName,omitempty"`
}

func (x *Path) Reset() {
	*x = Path{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_apimgr_apimgr_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Path) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Path) ProtoMessage() {}

func (x *Path) ProtoReflect() protoreflect.Message {
	mi := &file_npool_apimgr_apimgr_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Path.ProtoReflect.Descriptor instead.
func (*Path) Descriptor() ([]byte, []int) {
	return file_npool_apimgr_apimgr_proto_rawDescGZIP(), []int{0}
}

func (x *Path) GetMethod() string {
	if x != nil {
		return x.Method
	}
	return ""
}

func (x *Path) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *Path) GetExported() bool {
	if x != nil {
		return x.Exported
	}
	return false
}

func (x *Path) GetMethodName() string {
	if x != nil {
		return x.MethodName
	}
	return ""
}

type ServiceApis struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Protocol    string  `protobuf:"bytes,10,opt,name=Protocol,proto3" json:"Protocol,omitempty"`
	ServiceName string  `protobuf:"bytes,20,opt,name=ServiceName,proto3" json:"ServiceName,omitempty"`
	PathPrefix  string  `protobuf:"bytes,30,opt,name=PathPrefix,proto3" json:"PathPrefix,omitempty"`
	Paths       []*Path `protobuf:"bytes,40,rep,name=Paths,proto3" json:"Paths,omitempty"`
}

func (x *ServiceApis) Reset() {
	*x = ServiceApis{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_apimgr_apimgr_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServiceApis) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServiceApis) ProtoMessage() {}

func (x *ServiceApis) ProtoReflect() protoreflect.Message {
	mi := &file_npool_apimgr_apimgr_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServiceApis.ProtoReflect.Descriptor instead.
func (*ServiceApis) Descriptor() ([]byte, []int) {
	return file_npool_apimgr_apimgr_proto_rawDescGZIP(), []int{1}
}

func (x *ServiceApis) GetProtocol() string {
	if x != nil {
		return x.Protocol
	}
	return ""
}

func (x *ServiceApis) GetServiceName() string {
	if x != nil {
		return x.ServiceName
	}
	return ""
}

func (x *ServiceApis) GetPathPrefix() string {
	if x != nil {
		return x.PathPrefix
	}
	return ""
}

func (x *ServiceApis) GetPaths() []*Path {
	if x != nil {
		return x.Paths
	}
	return nil
}

type RegisterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *ServiceApis `protobuf:"bytes,10,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (x *RegisterRequest) Reset() {
	*x = RegisterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_apimgr_apimgr_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterRequest) ProtoMessage() {}

func (x *RegisterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_apimgr_apimgr_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterRequest.ProtoReflect.Descriptor instead.
func (*RegisterRequest) Descriptor() ([]byte, []int) {
	return file_npool_apimgr_apimgr_proto_rawDescGZIP(), []int{2}
}

func (x *RegisterRequest) GetInfo() *ServiceApis {
	if x != nil {
		return x.Info
	}
	return nil
}

type RegisterResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *ServiceApis `protobuf:"bytes,10,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (x *RegisterResponse) Reset() {
	*x = RegisterResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_apimgr_apimgr_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterResponse) ProtoMessage() {}

func (x *RegisterResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_apimgr_apimgr_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterResponse.ProtoReflect.Descriptor instead.
func (*RegisterResponse) Descriptor() ([]byte, []int) {
	return file_npool_apimgr_apimgr_proto_rawDescGZIP(), []int{3}
}

func (x *RegisterResponse) GetInfo() *ServiceApis {
	if x != nil {
		return x.Info
	}
	return nil
}

type ServicePath struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID          string   `protobuf:"bytes,10,opt,name=ID,proto3" json:"ID,omitempty"`
	Protocol    string   `protobuf:"bytes,20,opt,name=Protocol,proto3" json:"Protocol,omitempty"`
	ServiceName string   `protobuf:"bytes,30,opt,name=ServiceName,proto3" json:"ServiceName,omitempty"`
	Domains     []string `protobuf:"bytes,40,rep,name=Domains,proto3" json:"Domains,omitempty"`
	PathPrefix  string   `protobuf:"bytes,50,opt,name=PathPrefix,proto3" json:"PathPrefix,omitempty"`
	Method      string   `protobuf:"bytes,60,opt,name=Method,proto3" json:"Method,omitempty"`
	Path        string   `protobuf:"bytes,70,opt,name=Path,proto3" json:"Path,omitempty"`
	Exported    bool     `protobuf:"varint,80,opt,name=Exported,proto3" json:"Exported,omitempty"`
	CreateAt    uint32   `protobuf:"varint,90,opt,name=CreateAt,proto3" json:"CreateAt,omitempty"`
	UpdateAt    uint32   `protobuf:"varint,100,opt,name=UpdateAt,proto3" json:"UpdateAt,omitempty"`
	MethodName  string   `protobuf:"bytes,110,opt,name=MethodName,proto3" json:"MethodName,omitempty"`
}

func (x *ServicePath) Reset() {
	*x = ServicePath{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_apimgr_apimgr_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServicePath) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServicePath) ProtoMessage() {}

func (x *ServicePath) ProtoReflect() protoreflect.Message {
	mi := &file_npool_apimgr_apimgr_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServicePath.ProtoReflect.Descriptor instead.
func (*ServicePath) Descriptor() ([]byte, []int) {
	return file_npool_apimgr_apimgr_proto_rawDescGZIP(), []int{4}
}

func (x *ServicePath) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *ServicePath) GetProtocol() string {
	if x != nil {
		return x.Protocol
	}
	return ""
}

func (x *ServicePath) GetServiceName() string {
	if x != nil {
		return x.ServiceName
	}
	return ""
}

func (x *ServicePath) GetDomains() []string {
	if x != nil {
		return x.Domains
	}
	return nil
}

func (x *ServicePath) GetPathPrefix() string {
	if x != nil {
		return x.PathPrefix
	}
	return ""
}

func (x *ServicePath) GetMethod() string {
	if x != nil {
		return x.Method
	}
	return ""
}

func (x *ServicePath) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *ServicePath) GetExported() bool {
	if x != nil {
		return x.Exported
	}
	return false
}

func (x *ServicePath) GetCreateAt() uint32 {
	if x != nil {
		return x.CreateAt
	}
	return 0
}

func (x *ServicePath) GetUpdateAt() uint32 {
	if x != nil {
		return x.UpdateAt
	}
	return 0
}

func (x *ServicePath) GetMethodName() string {
	if x != nil {
		return x.MethodName
	}
	return ""
}

type GetApisRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetApisRequest) Reset() {
	*x = GetApisRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_apimgr_apimgr_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetApisRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetApisRequest) ProtoMessage() {}

func (x *GetApisRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_apimgr_apimgr_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetApisRequest.ProtoReflect.Descriptor instead.
func (*GetApisRequest) Descriptor() ([]byte, []int) {
	return file_npool_apimgr_apimgr_proto_rawDescGZIP(), []int{5}
}

type GetApisResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Infos []*ServicePath `protobuf:"bytes,10,rep,name=Infos,proto3" json:"Infos,omitempty"`
}

func (x *GetApisResponse) Reset() {
	*x = GetApisResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_apimgr_apimgr_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetApisResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetApisResponse) ProtoMessage() {}

func (x *GetApisResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_apimgr_apimgr_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetApisResponse.ProtoReflect.Descriptor instead.
func (*GetApisResponse) Descriptor() ([]byte, []int) {
	return file_npool_apimgr_apimgr_proto_rawDescGZIP(), []int{6}
}

func (x *GetApisResponse) GetInfos() []*ServicePath {
	if x != nil {
		return x.Infos
	}
	return nil
}

type GetServiceMethodAPIRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ServiceName string `protobuf:"bytes,10,opt,name=ServiceName,proto3" json:"ServiceName,omitempty"`
	MethodName  string `protobuf:"bytes,20,opt,name=MethodName,proto3" json:"MethodName,omitempty"`
}

func (x *GetServiceMethodAPIRequest) Reset() {
	*x = GetServiceMethodAPIRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_apimgr_apimgr_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetServiceMethodAPIRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetServiceMethodAPIRequest) ProtoMessage() {}

func (x *GetServiceMethodAPIRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_apimgr_apimgr_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetServiceMethodAPIRequest.ProtoReflect.Descriptor instead.
func (*GetServiceMethodAPIRequest) Descriptor() ([]byte, []int) {
	return file_npool_apimgr_apimgr_proto_rawDescGZIP(), []int{7}
}

func (x *GetServiceMethodAPIRequest) GetServiceName() string {
	if x != nil {
		return x.ServiceName
	}
	return ""
}

func (x *GetServiceMethodAPIRequest) GetMethodName() string {
	if x != nil {
		return x.MethodName
	}
	return ""
}

type GetServiceMethodAPIResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *ServicePath `protobuf:"bytes,10,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (x *GetServiceMethodAPIResponse) Reset() {
	*x = GetServiceMethodAPIResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_apimgr_apimgr_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetServiceMethodAPIResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetServiceMethodAPIResponse) ProtoMessage() {}

func (x *GetServiceMethodAPIResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_apimgr_apimgr_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetServiceMethodAPIResponse.ProtoReflect.Descriptor instead.
func (*GetServiceMethodAPIResponse) Descriptor() ([]byte, []int) {
	return file_npool_apimgr_apimgr_proto_rawDescGZIP(), []int{8}
}

func (x *GetServiceMethodAPIResponse) GetInfo() *ServicePath {
	if x != nil {
		return x.Info
	}
	return nil
}

var File_npool_apimgr_apimgr_proto protoreflect.FileDescriptor

var file_npool_apimgr_apimgr_proto_rawDesc = []byte{
	0x0a, 0x19, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x61, 0x70, 0x69, 0x6d, 0x67, 0x72, 0x2f, 0x61,
	0x70, 0x69, 0x6d, 0x67, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x61, 0x70, 0x69,
	0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x6e, 0x70,
	0x6f, 0x6f, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6e, 0x0a, 0x04, 0x50, 0x61, 0x74,
	0x68, 0x12, 0x16, 0x0a, 0x06, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x50, 0x61, 0x74,
	0x68, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x50, 0x61, 0x74, 0x68, 0x12, 0x1a, 0x0a,
	0x08, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x64, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x08, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x4d, 0x65, 0x74,
	0x68, 0x6f, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x28, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x4d,
	0x65, 0x74, 0x68, 0x6f, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x97, 0x01, 0x0a, 0x0b, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x41, 0x70, 0x69, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x12, 0x20, 0x0a, 0x0b, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x4e, 0x61, 0x6d, 0x65, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x50, 0x61, 0x74, 0x68, 0x50,
	0x72, 0x65, 0x66, 0x69, 0x78, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x50, 0x61, 0x74,
	0x68, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x12, 0x2a, 0x0a, 0x05, 0x50, 0x61, 0x74, 0x68, 0x73,
	0x18, 0x28, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x74, 0x68, 0x52, 0x05, 0x50, 0x61,
	0x74, 0x68, 0x73, 0x22, 0x42, 0x0a, 0x0f, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2f, 0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67,
	0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x41, 0x70, 0x69,
	0x73, 0x52, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x43, 0x0a, 0x10, 0x52, 0x65, 0x67, 0x69, 0x73,
	0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2f, 0x0a, 0x04, 0x49,
	0x6e, 0x66, 0x6f, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x41, 0x70, 0x69, 0x73, 0x52, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0xb5, 0x02, 0x0a,
	0x0b, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x61, 0x74, 0x68, 0x12, 0x0e, 0x0a, 0x02,
	0x49, 0x44, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x12, 0x20, 0x0a, 0x0b, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x44, 0x6f,
	0x6d, 0x61, 0x69, 0x6e, 0x73, 0x18, 0x28, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x44, 0x6f, 0x6d,
	0x61, 0x69, 0x6e, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x50, 0x61, 0x74, 0x68, 0x50, 0x72, 0x65, 0x66,
	0x69, 0x78, 0x18, 0x32, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x50, 0x61, 0x74, 0x68, 0x50, 0x72,
	0x65, 0x66, 0x69, 0x78, 0x12, 0x16, 0x0a, 0x06, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x3c,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x50, 0x61, 0x74, 0x68, 0x18, 0x46, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x50, 0x61, 0x74, 0x68,
	0x12, 0x1a, 0x0a, 0x08, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x64, 0x18, 0x50, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x08, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x64, 0x12, 0x1a, 0x0a, 0x08,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x74, 0x18, 0x5a, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x41, 0x74, 0x18, 0x64, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x41, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x4e, 0x61,
	0x6d, 0x65, 0x18, 0x6e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64,
	0x4e, 0x61, 0x6d, 0x65, 0x22, 0x10, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x41, 0x70, 0x69, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x44, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x41, 0x70, 0x69,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x31, 0x0a, 0x05, 0x49, 0x6e, 0x66,
	0x6f, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x50, 0x61, 0x74, 0x68, 0x52, 0x05, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x22, 0x5e, 0x0a, 0x1a,
	0x47, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64,
	0x41, 0x50, 0x49, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a,
	0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x4e, 0x0a, 0x1b,
	0x47, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64,
	0x41, 0x50, 0x49, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2f, 0x0a, 0x04, 0x49,
	0x6e, 0x66, 0x6f, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x50, 0x61, 0x74, 0x68, 0x52, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x32, 0xc4, 0x03, 0x0a,
	0x0a, 0x41, 0x70, 0x69, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x12, 0x51, 0x0a, 0x07, 0x56,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x19,
	0x2e, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x13, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x0d, 0x22, 0x08, 0x2f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x3a, 0x01, 0x2a, 0x12, 0x66,
	0x0a, 0x08, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x12, 0x1f, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x17, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x11, 0x22, 0x0c, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x67, 0x69, 0x73,
	0x74, 0x65, 0x72, 0x3a, 0x01, 0x2a, 0x12, 0x63, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x41, 0x70, 0x69,
	0x73, 0x12, 0x1e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e,
	0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x70, 0x69, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e,
	0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x70, 0x69, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x17, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x11, 0x22, 0x0c, 0x2f, 0x76, 0x31, 0x2f,
	0x67, 0x65, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x3a, 0x01, 0x2a, 0x12, 0x95, 0x01, 0x0a, 0x13,
	0x47, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64,
	0x41, 0x50, 0x49, 0x12, 0x2a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65,
	0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4d,
	0x65, 0x74, 0x68, 0x6f, 0x64, 0x41, 0x50, 0x49, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x2b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x31,
	0x2e, 0x47, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4d, 0x65, 0x74, 0x68, 0x6f,
	0x64, 0x41, 0x50, 0x49, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x25, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x1f, 0x22, 0x1a, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x65, 0x74, 0x2f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x2f, 0x61, 0x70, 0x69,
	0x3a, 0x01, 0x2a, 0x42, 0x2f, 0x5a, 0x2d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x4e, 0x70, 0x6f, 0x6f, 0x6c, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x61, 0x70,
	0x69, 0x6d, 0x67, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_npool_apimgr_apimgr_proto_rawDescOnce sync.Once
	file_npool_apimgr_apimgr_proto_rawDescData = file_npool_apimgr_apimgr_proto_rawDesc
)

func file_npool_apimgr_apimgr_proto_rawDescGZIP() []byte {
	file_npool_apimgr_apimgr_proto_rawDescOnce.Do(func() {
		file_npool_apimgr_apimgr_proto_rawDescData = protoimpl.X.CompressGZIP(file_npool_apimgr_apimgr_proto_rawDescData)
	})
	return file_npool_apimgr_apimgr_proto_rawDescData
}

var file_npool_apimgr_apimgr_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_npool_apimgr_apimgr_proto_goTypes = []interface{}{
	(*Path)(nil),                        // 0: api.manager.v1.Path
	(*ServiceApis)(nil),                 // 1: api.manager.v1.ServiceApis
	(*RegisterRequest)(nil),             // 2: api.manager.v1.RegisterRequest
	(*RegisterResponse)(nil),            // 3: api.manager.v1.RegisterResponse
	(*ServicePath)(nil),                 // 4: api.manager.v1.ServicePath
	(*GetApisRequest)(nil),              // 5: api.manager.v1.GetApisRequest
	(*GetApisResponse)(nil),             // 6: api.manager.v1.GetApisResponse
	(*GetServiceMethodAPIRequest)(nil),  // 7: api.manager.v1.GetServiceMethodAPIRequest
	(*GetServiceMethodAPIResponse)(nil), // 8: api.manager.v1.GetServiceMethodAPIResponse
	(*emptypb.Empty)(nil),               // 9: google.protobuf.Empty
	(*npool.VersionResponse)(nil),       // 10: npool.v1.VersionResponse
}
var file_npool_apimgr_apimgr_proto_depIdxs = []int32{
	0,  // 0: api.manager.v1.ServiceApis.Paths:type_name -> api.manager.v1.Path
	1,  // 1: api.manager.v1.RegisterRequest.Info:type_name -> api.manager.v1.ServiceApis
	1,  // 2: api.manager.v1.RegisterResponse.Info:type_name -> api.manager.v1.ServiceApis
	4,  // 3: api.manager.v1.GetApisResponse.Infos:type_name -> api.manager.v1.ServicePath
	4,  // 4: api.manager.v1.GetServiceMethodAPIResponse.Info:type_name -> api.manager.v1.ServicePath
	9,  // 5: api.manager.v1.ApiManager.Version:input_type -> google.protobuf.Empty
	2,  // 6: api.manager.v1.ApiManager.Register:input_type -> api.manager.v1.RegisterRequest
	5,  // 7: api.manager.v1.ApiManager.GetApis:input_type -> api.manager.v1.GetApisRequest
	7,  // 8: api.manager.v1.ApiManager.GetServiceMethodAPI:input_type -> api.manager.v1.GetServiceMethodAPIRequest
	10, // 9: api.manager.v1.ApiManager.Version:output_type -> npool.v1.VersionResponse
	3,  // 10: api.manager.v1.ApiManager.Register:output_type -> api.manager.v1.RegisterResponse
	6,  // 11: api.manager.v1.ApiManager.GetApis:output_type -> api.manager.v1.GetApisResponse
	8,  // 12: api.manager.v1.ApiManager.GetServiceMethodAPI:output_type -> api.manager.v1.GetServiceMethodAPIResponse
	9,  // [9:13] is the sub-list for method output_type
	5,  // [5:9] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_npool_apimgr_apimgr_proto_init() }
func file_npool_apimgr_apimgr_proto_init() {
	if File_npool_apimgr_apimgr_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_npool_apimgr_apimgr_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Path); i {
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
		file_npool_apimgr_apimgr_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServiceApis); i {
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
		file_npool_apimgr_apimgr_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterRequest); i {
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
		file_npool_apimgr_apimgr_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterResponse); i {
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
		file_npool_apimgr_apimgr_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServicePath); i {
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
		file_npool_apimgr_apimgr_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetApisRequest); i {
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
		file_npool_apimgr_apimgr_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetApisResponse); i {
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
		file_npool_apimgr_apimgr_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetServiceMethodAPIRequest); i {
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
		file_npool_apimgr_apimgr_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetServiceMethodAPIResponse); i {
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
			RawDescriptor: file_npool_apimgr_apimgr_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_npool_apimgr_apimgr_proto_goTypes,
		DependencyIndexes: file_npool_apimgr_apimgr_proto_depIdxs,
		MessageInfos:      file_npool_apimgr_apimgr_proto_msgTypes,
	}.Build()
	File_npool_apimgr_apimgr_proto = out.File
	file_npool_apimgr_apimgr_proto_rawDesc = nil
	file_npool_apimgr_apimgr_proto_goTypes = nil
	file_npool_apimgr_apimgr_proto_depIdxs = nil
}
