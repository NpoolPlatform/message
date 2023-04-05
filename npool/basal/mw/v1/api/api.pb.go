// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.18.1
// source: npool/basal/mw/v1/api/api.proto

package api

import (
	_ "github.com/NpoolPlatform/message/npool"
	api "github.com/NpoolPlatform/message/npool/basal/mgr/v1/api"
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

type CreateAPIRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *api.APIReq `protobuf:"bytes,10,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (x *CreateAPIRequest) Reset() {
	*x = CreateAPIRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_basal_mw_v1_api_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateAPIRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAPIRequest) ProtoMessage() {}

func (x *CreateAPIRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_basal_mw_v1_api_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAPIRequest.ProtoReflect.Descriptor instead.
func (*CreateAPIRequest) Descriptor() ([]byte, []int) {
	return file_npool_basal_mw_v1_api_api_proto_rawDescGZIP(), []int{0}
}

func (x *CreateAPIRequest) GetInfo() *api.APIReq {
	if x != nil {
		return x.Info
	}
	return nil
}

type CreateAPIResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *api.API `protobuf:"bytes,10,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (x *CreateAPIResponse) Reset() {
	*x = CreateAPIResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_basal_mw_v1_api_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateAPIResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAPIResponse) ProtoMessage() {}

func (x *CreateAPIResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_basal_mw_v1_api_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAPIResponse.ProtoReflect.Descriptor instead.
func (*CreateAPIResponse) Descriptor() ([]byte, []int) {
	return file_npool_basal_mw_v1_api_api_proto_rawDescGZIP(), []int{1}
}

func (x *CreateAPIResponse) GetInfo() *api.API {
	if x != nil {
		return x.Info
	}
	return nil
}

type CreateAPIsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Infos []*api.APIReq `protobuf:"bytes,10,rep,name=Infos,proto3" json:"Infos,omitempty"`
}

func (x *CreateAPIsRequest) Reset() {
	*x = CreateAPIsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_basal_mw_v1_api_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateAPIsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAPIsRequest) ProtoMessage() {}

func (x *CreateAPIsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_basal_mw_v1_api_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAPIsRequest.ProtoReflect.Descriptor instead.
func (*CreateAPIsRequest) Descriptor() ([]byte, []int) {
	return file_npool_basal_mw_v1_api_api_proto_rawDescGZIP(), []int{2}
}

func (x *CreateAPIsRequest) GetInfos() []*api.APIReq {
	if x != nil {
		return x.Infos
	}
	return nil
}

type CreateAPIsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Infos []*api.API `protobuf:"bytes,10,rep,name=Infos,proto3" json:"Infos,omitempty"`
}

func (x *CreateAPIsResponse) Reset() {
	*x = CreateAPIsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_basal_mw_v1_api_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateAPIsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAPIsResponse) ProtoMessage() {}

func (x *CreateAPIsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_basal_mw_v1_api_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAPIsResponse.ProtoReflect.Descriptor instead.
func (*CreateAPIsResponse) Descriptor() ([]byte, []int) {
	return file_npool_basal_mw_v1_api_api_proto_rawDescGZIP(), []int{3}
}

func (x *CreateAPIsResponse) GetInfos() []*api.API {
	if x != nil {
		return x.Infos
	}
	return nil
}

type UpdateAPIRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *api.APIReq `protobuf:"bytes,10,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (x *UpdateAPIRequest) Reset() {
	*x = UpdateAPIRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_basal_mw_v1_api_api_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateAPIRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateAPIRequest) ProtoMessage() {}

func (x *UpdateAPIRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_basal_mw_v1_api_api_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateAPIRequest.ProtoReflect.Descriptor instead.
func (*UpdateAPIRequest) Descriptor() ([]byte, []int) {
	return file_npool_basal_mw_v1_api_api_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateAPIRequest) GetInfo() *api.APIReq {
	if x != nil {
		return x.Info
	}
	return nil
}

type UpdateAPIResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *api.API `protobuf:"bytes,10,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (x *UpdateAPIResponse) Reset() {
	*x = UpdateAPIResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_basal_mw_v1_api_api_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateAPIResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateAPIResponse) ProtoMessage() {}

func (x *UpdateAPIResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_basal_mw_v1_api_api_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateAPIResponse.ProtoReflect.Descriptor instead.
func (*UpdateAPIResponse) Descriptor() ([]byte, []int) {
	return file_npool_basal_mw_v1_api_api_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateAPIResponse) GetInfo() *api.API {
	if x != nil {
		return x.Info
	}
	return nil
}

type GetAPIsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Conds  *api.Conds `protobuf:"bytes,10,opt,name=Conds,proto3" json:"Conds,omitempty"`
	Offset int32      `protobuf:"varint,20,opt,name=Offset,proto3" json:"Offset,omitempty"`
	Limit  int32      `protobuf:"varint,30,opt,name=Limit,proto3" json:"Limit,omitempty"`
}

func (x *GetAPIsRequest) Reset() {
	*x = GetAPIsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_basal_mw_v1_api_api_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAPIsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAPIsRequest) ProtoMessage() {}

func (x *GetAPIsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_basal_mw_v1_api_api_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAPIsRequest.ProtoReflect.Descriptor instead.
func (*GetAPIsRequest) Descriptor() ([]byte, []int) {
	return file_npool_basal_mw_v1_api_api_proto_rawDescGZIP(), []int{6}
}

func (x *GetAPIsRequest) GetConds() *api.Conds {
	if x != nil {
		return x.Conds
	}
	return nil
}

func (x *GetAPIsRequest) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *GetAPIsRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type GetAPIsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Infos []*api.API `protobuf:"bytes,10,rep,name=Infos,proto3" json:"Infos,omitempty"`
	Total uint32     `protobuf:"varint,20,opt,name=Total,proto3" json:"Total,omitempty"`
}

func (x *GetAPIsResponse) Reset() {
	*x = GetAPIsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_basal_mw_v1_api_api_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAPIsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAPIsResponse) ProtoMessage() {}

func (x *GetAPIsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_basal_mw_v1_api_api_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAPIsResponse.ProtoReflect.Descriptor instead.
func (*GetAPIsResponse) Descriptor() ([]byte, []int) {
	return file_npool_basal_mw_v1_api_api_proto_rawDescGZIP(), []int{7}
}

func (x *GetAPIsResponse) GetInfos() []*api.API {
	if x != nil {
		return x.Infos
	}
	return nil
}

func (x *GetAPIsResponse) GetTotal() uint32 {
	if x != nil {
		return x.Total
	}
	return 0
}

type GetDomainsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetDomainsRequest) Reset() {
	*x = GetDomainsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_basal_mw_v1_api_api_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDomainsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDomainsRequest) ProtoMessage() {}

func (x *GetDomainsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_basal_mw_v1_api_api_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDomainsRequest.ProtoReflect.Descriptor instead.
func (*GetDomainsRequest) Descriptor() ([]byte, []int) {
	return file_npool_basal_mw_v1_api_api_proto_rawDescGZIP(), []int{8}
}

type GetDomainsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Infos []string `protobuf:"bytes,10,rep,name=Infos,proto3" json:"Infos,omitempty"`
}

func (x *GetDomainsResponse) Reset() {
	*x = GetDomainsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_basal_mw_v1_api_api_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDomainsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDomainsResponse) ProtoMessage() {}

func (x *GetDomainsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_basal_mw_v1_api_api_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDomainsResponse.ProtoReflect.Descriptor instead.
func (*GetDomainsResponse) Descriptor() ([]byte, []int) {
	return file_npool_basal_mw_v1_api_api_proto_rawDescGZIP(), []int{9}
}

func (x *GetDomainsResponse) GetInfos() []string {
	if x != nil {
		return x.Infos
	}
	return nil
}

var File_npool_basal_mw_v1_api_api_proto protoreflect.FileDescriptor

var file_npool_basal_mw_v1_api_api_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x62, 0x61, 0x73, 0x61, 0x6c, 0x2f, 0x6d, 0x77,
	0x2f, 0x76, 0x31, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x17, 0x62, 0x61, 0x73, 0x61, 0x6c, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77,
	0x61, 0x72, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x1a, 0x11, 0x6e, 0x70, 0x6f, 0x6f,
	0x6c, 0x2f, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x6e,
	0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x62, 0x61, 0x73, 0x61, 0x6c, 0x2f, 0x6d, 0x67, 0x72, 0x2f, 0x76,
	0x31, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x44, 0x0a, 0x10, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x50, 0x49, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x30, 0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1c, 0x2e, 0x62, 0x61, 0x73, 0x61, 0x6c, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65,
	0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x50, 0x49, 0x52, 0x65, 0x71, 0x52,
	0x04, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x42, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41,
	0x50, 0x49, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2d, 0x0a, 0x04, 0x49, 0x6e,
	0x66, 0x6f, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x62, 0x61, 0x73, 0x61, 0x6c,
	0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e,
	0x41, 0x50, 0x49, 0x52, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x47, 0x0a, 0x11, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x41, 0x50, 0x49, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x32,
	0x0a, 0x05, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e,
	0x62, 0x61, 0x73, 0x61, 0x6c, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x50, 0x49, 0x52, 0x65, 0x71, 0x52, 0x05, 0x49, 0x6e, 0x66,
	0x6f, 0x73, 0x22, 0x45, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x50, 0x49, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2f, 0x0a, 0x05, 0x49, 0x6e, 0x66, 0x6f,
	0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x62, 0x61, 0x73, 0x61, 0x6c, 0x2e,
	0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x41,
	0x50, 0x49, 0x52, 0x05, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x22, 0x44, 0x0a, 0x10, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x41, 0x50, 0x49, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x30, 0x0a,
	0x04, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x62, 0x61,
	0x73, 0x61, 0x6c, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x76, 0x31, 0x2e, 0x41, 0x50, 0x49, 0x52, 0x65, 0x71, 0x52, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x22,
	0x42, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x50, 0x49, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2d, 0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x19, 0x2e, 0x62, 0x61, 0x73, 0x61, 0x6c, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67,
	0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x50, 0x49, 0x52, 0x04, 0x49,
	0x6e, 0x66, 0x6f, 0x22, 0x71, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x41, 0x50, 0x49, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x31, 0x0a, 0x05, 0x43, 0x6f, 0x6e, 0x64, 0x73, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x62, 0x61, 0x73, 0x61, 0x6c, 0x2e, 0x6d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x64,
	0x73, 0x52, 0x05, 0x43, 0x6f, 0x6e, 0x64, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x4f, 0x66, 0x66, 0x73,
	0x65, 0x74, 0x18, 0x14, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74,
	0x12, 0x14, 0x0a, 0x05, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x05, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x22, 0x58, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x41, 0x50, 0x49,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2f, 0x0a, 0x05, 0x49, 0x6e, 0x66,
	0x6f, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x62, 0x61, 0x73, 0x61, 0x6c,
	0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e,
	0x41, 0x50, 0x49, 0x52, 0x05, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x6f,
	0x74, 0x61, 0x6c, 0x18, 0x14, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c,
	0x22, 0x13, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x2a, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x44, 0x6f, 0x6d, 0x61,
	0x69, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x49,
	0x6e, 0x66, 0x6f, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x49, 0x6e, 0x66, 0x6f,
	0x73, 0x32, 0x8a, 0x04, 0x0a, 0x0a, 0x4d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65,
	0x12, 0x64, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x50, 0x49, 0x12, 0x29, 0x2e,
	0x62, 0x61, 0x73, 0x61, 0x6c, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x50,
	0x49, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x62, 0x61, 0x73, 0x61, 0x6c,
	0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x50, 0x49, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x67, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x41, 0x50, 0x49, 0x73, 0x12, 0x2a, 0x2e, 0x62, 0x61, 0x73, 0x61, 0x6c, 0x2e, 0x6d, 0x69, 0x64,
	0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x50, 0x49, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x2b, 0x2e, 0x62, 0x61, 0x73, 0x61, 0x6c, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77,
	0x61, 0x72, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x41, 0x50, 0x49, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x64, 0x0a, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x50, 0x49, 0x12, 0x29, 0x2e, 0x62,
	0x61, 0x73, 0x61, 0x6c, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x50, 0x49,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x62, 0x61, 0x73, 0x61, 0x6c, 0x2e,
	0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76,
	0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x50, 0x49, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x5e, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x41, 0x50, 0x49, 0x73,
	0x12, 0x27, 0x2e, 0x62, 0x61, 0x73, 0x61, 0x6c, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77,
	0x61, 0x72, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x50,
	0x49, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x62, 0x61, 0x73, 0x61,
	0x6c, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x50, 0x49, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x67, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x44, 0x6f, 0x6d, 0x61,
	0x69, 0x6e, 0x73, 0x12, 0x2a, 0x2e, 0x62, 0x61, 0x73, 0x61, 0x6c, 0x2e, 0x6d, 0x69, 0x64, 0x64,
	0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65,
	0x74, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x2b, 0x2e, 0x62, 0x61, 0x73, 0x61, 0x6c, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61,
	0x72, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x44, 0x6f, 0x6d,
	0x61, 0x69, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x38,
	0x5a, 0x36, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4e, 0x70, 0x6f,
	0x6f, 0x6c, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x2f, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x62, 0x61, 0x73, 0x61, 0x6c, 0x2f, 0x6d,
	0x77, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_npool_basal_mw_v1_api_api_proto_rawDescOnce sync.Once
	file_npool_basal_mw_v1_api_api_proto_rawDescData = file_npool_basal_mw_v1_api_api_proto_rawDesc
)

func file_npool_basal_mw_v1_api_api_proto_rawDescGZIP() []byte {
	file_npool_basal_mw_v1_api_api_proto_rawDescOnce.Do(func() {
		file_npool_basal_mw_v1_api_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_npool_basal_mw_v1_api_api_proto_rawDescData)
	})
	return file_npool_basal_mw_v1_api_api_proto_rawDescData
}

var file_npool_basal_mw_v1_api_api_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_npool_basal_mw_v1_api_api_proto_goTypes = []interface{}{
	(*CreateAPIRequest)(nil),   // 0: basal.middleware.api.v1.CreateAPIRequest
	(*CreateAPIResponse)(nil),  // 1: basal.middleware.api.v1.CreateAPIResponse
	(*CreateAPIsRequest)(nil),  // 2: basal.middleware.api.v1.CreateAPIsRequest
	(*CreateAPIsResponse)(nil), // 3: basal.middleware.api.v1.CreateAPIsResponse
	(*UpdateAPIRequest)(nil),   // 4: basal.middleware.api.v1.UpdateAPIRequest
	(*UpdateAPIResponse)(nil),  // 5: basal.middleware.api.v1.UpdateAPIResponse
	(*GetAPIsRequest)(nil),     // 6: basal.middleware.api.v1.GetAPIsRequest
	(*GetAPIsResponse)(nil),    // 7: basal.middleware.api.v1.GetAPIsResponse
	(*GetDomainsRequest)(nil),  // 8: basal.middleware.api.v1.GetDomainsRequest
	(*GetDomainsResponse)(nil), // 9: basal.middleware.api.v1.GetDomainsResponse
	(*api.APIReq)(nil),         // 10: basal.manager.api.v1.APIReq
	(*api.API)(nil),            // 11: basal.manager.api.v1.API
	(*api.Conds)(nil),          // 12: basal.manager.api.v1.Conds
}
var file_npool_basal_mw_v1_api_api_proto_depIdxs = []int32{
	10, // 0: basal.middleware.api.v1.CreateAPIRequest.Info:type_name -> basal.manager.api.v1.APIReq
	11, // 1: basal.middleware.api.v1.CreateAPIResponse.Info:type_name -> basal.manager.api.v1.API
	10, // 2: basal.middleware.api.v1.CreateAPIsRequest.Infos:type_name -> basal.manager.api.v1.APIReq
	11, // 3: basal.middleware.api.v1.CreateAPIsResponse.Infos:type_name -> basal.manager.api.v1.API
	10, // 4: basal.middleware.api.v1.UpdateAPIRequest.Info:type_name -> basal.manager.api.v1.APIReq
	11, // 5: basal.middleware.api.v1.UpdateAPIResponse.Info:type_name -> basal.manager.api.v1.API
	12, // 6: basal.middleware.api.v1.GetAPIsRequest.Conds:type_name -> basal.manager.api.v1.Conds
	11, // 7: basal.middleware.api.v1.GetAPIsResponse.Infos:type_name -> basal.manager.api.v1.API
	0,  // 8: basal.middleware.api.v1.Middleware.CreateAPI:input_type -> basal.middleware.api.v1.CreateAPIRequest
	2,  // 9: basal.middleware.api.v1.Middleware.CreateAPIs:input_type -> basal.middleware.api.v1.CreateAPIsRequest
	4,  // 10: basal.middleware.api.v1.Middleware.UpdateAPI:input_type -> basal.middleware.api.v1.UpdateAPIRequest
	6,  // 11: basal.middleware.api.v1.Middleware.GetAPIs:input_type -> basal.middleware.api.v1.GetAPIsRequest
	8,  // 12: basal.middleware.api.v1.Middleware.GetDomains:input_type -> basal.middleware.api.v1.GetDomainsRequest
	1,  // 13: basal.middleware.api.v1.Middleware.CreateAPI:output_type -> basal.middleware.api.v1.CreateAPIResponse
	3,  // 14: basal.middleware.api.v1.Middleware.CreateAPIs:output_type -> basal.middleware.api.v1.CreateAPIsResponse
	5,  // 15: basal.middleware.api.v1.Middleware.UpdateAPI:output_type -> basal.middleware.api.v1.UpdateAPIResponse
	7,  // 16: basal.middleware.api.v1.Middleware.GetAPIs:output_type -> basal.middleware.api.v1.GetAPIsResponse
	9,  // 17: basal.middleware.api.v1.Middleware.GetDomains:output_type -> basal.middleware.api.v1.GetDomainsResponse
	13, // [13:18] is the sub-list for method output_type
	8,  // [8:13] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_npool_basal_mw_v1_api_api_proto_init() }
func file_npool_basal_mw_v1_api_api_proto_init() {
	if File_npool_basal_mw_v1_api_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_npool_basal_mw_v1_api_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateAPIRequest); i {
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
		file_npool_basal_mw_v1_api_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateAPIResponse); i {
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
		file_npool_basal_mw_v1_api_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateAPIsRequest); i {
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
		file_npool_basal_mw_v1_api_api_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateAPIsResponse); i {
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
		file_npool_basal_mw_v1_api_api_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateAPIRequest); i {
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
		file_npool_basal_mw_v1_api_api_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateAPIResponse); i {
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
		file_npool_basal_mw_v1_api_api_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAPIsRequest); i {
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
		file_npool_basal_mw_v1_api_api_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAPIsResponse); i {
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
		file_npool_basal_mw_v1_api_api_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDomainsRequest); i {
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
		file_npool_basal_mw_v1_api_api_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDomainsResponse); i {
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
			RawDescriptor: file_npool_basal_mw_v1_api_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_npool_basal_mw_v1_api_api_proto_goTypes,
		DependencyIndexes: file_npool_basal_mw_v1_api_api_proto_depIdxs,
		MessageInfos:      file_npool_basal_mw_v1_api_api_proto_msgTypes,
	}.Build()
	File_npool_basal_mw_v1_api_api_proto = out.File
	file_npool_basal_mw_v1_api_api_proto_rawDesc = nil
	file_npool_basal_mw_v1_api_api_proto_goTypes = nil
	file_npool_basal_mw_v1_api_api_proto_depIdxs = nil
}
