// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.18.1
// source: npool/inspire/gw/v1/coin/config/config.proto

package config

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type CoinConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID           uint32   `protobuf:"varint,10,opt,name=ID,proto3" json:"ID,omitempty"`
	EntID        string   `protobuf:"bytes,20,opt,name=EntID,proto3" json:"EntID,omitempty"`
	AppID        string   `protobuf:"bytes,30,opt,name=AppID,proto3" json:"AppID,omitempty"`
	CoinTypeID   string   `protobuf:"bytes,40,opt,name=CoinTypeID,proto3" json:"CoinTypeID,omitempty"`
	MaxValue     string   `protobuf:"bytes,50,opt,name=MaxValue,proto3" json:"MaxValue,omitempty"`
	Allocated    string   `protobuf:"bytes,60,opt,name=Allocated,proto3" json:"Allocated,omitempty"`
	CoinName     string   `protobuf:"bytes,70,opt,name=CoinName,proto3" json:"CoinName,omitempty"`
	DisplayNames []string `protobuf:"bytes,80,rep,name=DisplayNames,proto3" json:"DisplayNames,omitempty"`
	CoinLogo     string   `protobuf:"bytes,90,opt,name=CoinLogo,proto3" json:"CoinLogo,omitempty"`
	CoinUnit     string   `protobuf:"bytes,100,opt,name=CoinUnit,proto3" json:"CoinUnit,omitempty"`
	CreatedAt    uint32   `protobuf:"varint,1000,opt,name=CreatedAt,proto3" json:"CreatedAt,omitempty"`
	UpdatedAt    uint32   `protobuf:"varint,1010,opt,name=UpdatedAt,proto3" json:"UpdatedAt,omitempty"`
}

func (x *CoinConfig) Reset() {
	*x = CoinConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_inspire_gw_v1_coin_config_config_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CoinConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CoinConfig) ProtoMessage() {}

func (x *CoinConfig) ProtoReflect() protoreflect.Message {
	mi := &file_npool_inspire_gw_v1_coin_config_config_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CoinConfig.ProtoReflect.Descriptor instead.
func (*CoinConfig) Descriptor() ([]byte, []int) {
	return file_npool_inspire_gw_v1_coin_config_config_proto_rawDescGZIP(), []int{0}
}

func (x *CoinConfig) GetID() uint32 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *CoinConfig) GetEntID() string {
	if x != nil {
		return x.EntID
	}
	return ""
}

func (x *CoinConfig) GetAppID() string {
	if x != nil {
		return x.AppID
	}
	return ""
}

func (x *CoinConfig) GetCoinTypeID() string {
	if x != nil {
		return x.CoinTypeID
	}
	return ""
}

func (x *CoinConfig) GetMaxValue() string {
	if x != nil {
		return x.MaxValue
	}
	return ""
}

func (x *CoinConfig) GetAllocated() string {
	if x != nil {
		return x.Allocated
	}
	return ""
}

func (x *CoinConfig) GetCoinName() string {
	if x != nil {
		return x.CoinName
	}
	return ""
}

func (x *CoinConfig) GetDisplayNames() []string {
	if x != nil {
		return x.DisplayNames
	}
	return nil
}

func (x *CoinConfig) GetCoinLogo() string {
	if x != nil {
		return x.CoinLogo
	}
	return ""
}

func (x *CoinConfig) GetCoinUnit() string {
	if x != nil {
		return x.CoinUnit
	}
	return ""
}

func (x *CoinConfig) GetCreatedAt() uint32 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *CoinConfig) GetUpdatedAt() uint32 {
	if x != nil {
		return x.UpdatedAt
	}
	return 0
}

type AdminCreateCoinConfigRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TargetAppID string `protobuf:"bytes,10,opt,name=TargetAppID,proto3" json:"TargetAppID,omitempty"`
	CoinTypeID  string `protobuf:"bytes,20,opt,name=CoinTypeID,proto3" json:"CoinTypeID,omitempty"`
	MaxValue    string `protobuf:"bytes,30,opt,name=MaxValue,proto3" json:"MaxValue,omitempty"`
}

func (x *AdminCreateCoinConfigRequest) Reset() {
	*x = AdminCreateCoinConfigRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_inspire_gw_v1_coin_config_config_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdminCreateCoinConfigRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdminCreateCoinConfigRequest) ProtoMessage() {}

func (x *AdminCreateCoinConfigRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_inspire_gw_v1_coin_config_config_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdminCreateCoinConfigRequest.ProtoReflect.Descriptor instead.
func (*AdminCreateCoinConfigRequest) Descriptor() ([]byte, []int) {
	return file_npool_inspire_gw_v1_coin_config_config_proto_rawDescGZIP(), []int{1}
}

func (x *AdminCreateCoinConfigRequest) GetTargetAppID() string {
	if x != nil {
		return x.TargetAppID
	}
	return ""
}

func (x *AdminCreateCoinConfigRequest) GetCoinTypeID() string {
	if x != nil {
		return x.CoinTypeID
	}
	return ""
}

func (x *AdminCreateCoinConfigRequest) GetMaxValue() string {
	if x != nil {
		return x.MaxValue
	}
	return ""
}

type AdminCreateCoinConfigResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *CoinConfig `protobuf:"bytes,10,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (x *AdminCreateCoinConfigResponse) Reset() {
	*x = AdminCreateCoinConfigResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_inspire_gw_v1_coin_config_config_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdminCreateCoinConfigResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdminCreateCoinConfigResponse) ProtoMessage() {}

func (x *AdminCreateCoinConfigResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_inspire_gw_v1_coin_config_config_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdminCreateCoinConfigResponse.ProtoReflect.Descriptor instead.
func (*AdminCreateCoinConfigResponse) Descriptor() ([]byte, []int) {
	return file_npool_inspire_gw_v1_coin_config_config_proto_rawDescGZIP(), []int{2}
}

func (x *AdminCreateCoinConfigResponse) GetInfo() *CoinConfig {
	if x != nil {
		return x.Info
	}
	return nil
}

type AdminUpdateCoinConfigRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID          uint32  `protobuf:"varint,10,opt,name=ID,proto3" json:"ID,omitempty"`
	EntID       string  `protobuf:"bytes,20,opt,name=EntID,proto3" json:"EntID,omitempty"`
	TargetAppID string  `protobuf:"bytes,30,opt,name=TargetAppID,proto3" json:"TargetAppID,omitempty"`
	MaxValue    *string `protobuf:"bytes,40,opt,name=MaxValue,proto3,oneof" json:"MaxValue,omitempty"`
}

func (x *AdminUpdateCoinConfigRequest) Reset() {
	*x = AdminUpdateCoinConfigRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_inspire_gw_v1_coin_config_config_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdminUpdateCoinConfigRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdminUpdateCoinConfigRequest) ProtoMessage() {}

func (x *AdminUpdateCoinConfigRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_inspire_gw_v1_coin_config_config_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdminUpdateCoinConfigRequest.ProtoReflect.Descriptor instead.
func (*AdminUpdateCoinConfigRequest) Descriptor() ([]byte, []int) {
	return file_npool_inspire_gw_v1_coin_config_config_proto_rawDescGZIP(), []int{3}
}

func (x *AdminUpdateCoinConfigRequest) GetID() uint32 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *AdminUpdateCoinConfigRequest) GetEntID() string {
	if x != nil {
		return x.EntID
	}
	return ""
}

func (x *AdminUpdateCoinConfigRequest) GetTargetAppID() string {
	if x != nil {
		return x.TargetAppID
	}
	return ""
}

func (x *AdminUpdateCoinConfigRequest) GetMaxValue() string {
	if x != nil && x.MaxValue != nil {
		return *x.MaxValue
	}
	return ""
}

type AdminUpdateCoinConfigResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *CoinConfig `protobuf:"bytes,10,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (x *AdminUpdateCoinConfigResponse) Reset() {
	*x = AdminUpdateCoinConfigResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_inspire_gw_v1_coin_config_config_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdminUpdateCoinConfigResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdminUpdateCoinConfigResponse) ProtoMessage() {}

func (x *AdminUpdateCoinConfigResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_inspire_gw_v1_coin_config_config_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdminUpdateCoinConfigResponse.ProtoReflect.Descriptor instead.
func (*AdminUpdateCoinConfigResponse) Descriptor() ([]byte, []int) {
	return file_npool_inspire_gw_v1_coin_config_config_proto_rawDescGZIP(), []int{4}
}

func (x *AdminUpdateCoinConfigResponse) GetInfo() *CoinConfig {
	if x != nil {
		return x.Info
	}
	return nil
}

type AdminGetCoinConfigsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TargetAppID string `protobuf:"bytes,10,opt,name=TargetAppID,proto3" json:"TargetAppID,omitempty"`
	Offset      int32  `protobuf:"varint,20,opt,name=Offset,proto3" json:"Offset,omitempty"`
	Limit       int32  `protobuf:"varint,30,opt,name=Limit,proto3" json:"Limit,omitempty"`
}

func (x *AdminGetCoinConfigsRequest) Reset() {
	*x = AdminGetCoinConfigsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_inspire_gw_v1_coin_config_config_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdminGetCoinConfigsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdminGetCoinConfigsRequest) ProtoMessage() {}

func (x *AdminGetCoinConfigsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_inspire_gw_v1_coin_config_config_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdminGetCoinConfigsRequest.ProtoReflect.Descriptor instead.
func (*AdminGetCoinConfigsRequest) Descriptor() ([]byte, []int) {
	return file_npool_inspire_gw_v1_coin_config_config_proto_rawDescGZIP(), []int{5}
}

func (x *AdminGetCoinConfigsRequest) GetTargetAppID() string {
	if x != nil {
		return x.TargetAppID
	}
	return ""
}

func (x *AdminGetCoinConfigsRequest) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *AdminGetCoinConfigsRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type AdminGetCoinConfigsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Infos []*CoinConfig `protobuf:"bytes,10,rep,name=Infos,proto3" json:"Infos,omitempty"`
	Total uint32        `protobuf:"varint,20,opt,name=Total,proto3" json:"Total,omitempty"`
}

func (x *AdminGetCoinConfigsResponse) Reset() {
	*x = AdminGetCoinConfigsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_inspire_gw_v1_coin_config_config_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdminGetCoinConfigsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdminGetCoinConfigsResponse) ProtoMessage() {}

func (x *AdminGetCoinConfigsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_inspire_gw_v1_coin_config_config_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdminGetCoinConfigsResponse.ProtoReflect.Descriptor instead.
func (*AdminGetCoinConfigsResponse) Descriptor() ([]byte, []int) {
	return file_npool_inspire_gw_v1_coin_config_config_proto_rawDescGZIP(), []int{6}
}

func (x *AdminGetCoinConfigsResponse) GetInfos() []*CoinConfig {
	if x != nil {
		return x.Infos
	}
	return nil
}

func (x *AdminGetCoinConfigsResponse) GetTotal() uint32 {
	if x != nil {
		return x.Total
	}
	return 0
}

type AdminDeleteCoinConfigRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID    uint32 `protobuf:"varint,10,opt,name=ID,proto3" json:"ID,omitempty"`
	EntID string `protobuf:"bytes,20,opt,name=EntID,proto3" json:"EntID,omitempty"`
}

func (x *AdminDeleteCoinConfigRequest) Reset() {
	*x = AdminDeleteCoinConfigRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_inspire_gw_v1_coin_config_config_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdminDeleteCoinConfigRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdminDeleteCoinConfigRequest) ProtoMessage() {}

func (x *AdminDeleteCoinConfigRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_inspire_gw_v1_coin_config_config_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdminDeleteCoinConfigRequest.ProtoReflect.Descriptor instead.
func (*AdminDeleteCoinConfigRequest) Descriptor() ([]byte, []int) {
	return file_npool_inspire_gw_v1_coin_config_config_proto_rawDescGZIP(), []int{7}
}

func (x *AdminDeleteCoinConfigRequest) GetID() uint32 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *AdminDeleteCoinConfigRequest) GetEntID() string {
	if x != nil {
		return x.EntID
	}
	return ""
}

type AdminDeleteCoinConfigResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *CoinConfig `protobuf:"bytes,10,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (x *AdminDeleteCoinConfigResponse) Reset() {
	*x = AdminDeleteCoinConfigResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_inspire_gw_v1_coin_config_config_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdminDeleteCoinConfigResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdminDeleteCoinConfigResponse) ProtoMessage() {}

func (x *AdminDeleteCoinConfigResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_inspire_gw_v1_coin_config_config_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdminDeleteCoinConfigResponse.ProtoReflect.Descriptor instead.
func (*AdminDeleteCoinConfigResponse) Descriptor() ([]byte, []int) {
	return file_npool_inspire_gw_v1_coin_config_config_proto_rawDescGZIP(), []int{8}
}

func (x *AdminDeleteCoinConfigResponse) GetInfo() *CoinConfig {
	if x != nil {
		return x.Info
	}
	return nil
}

var File_npool_inspire_gw_v1_coin_config_config_proto protoreflect.FileDescriptor

var file_npool_inspire_gw_v1_coin_config_config_proto_rawDesc = []byte{
	0x0a, 0x2c, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x69, 0x6e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2f,
	0x67, 0x77, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x69, 0x6e, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1e,
	0x69, 0x6e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e,
	0x63, 0x6f, 0x69, 0x6e, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76, 0x31, 0x1a, 0x1c,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd8, 0x02, 0x0a,
	0x0a, 0x43, 0x6f, 0x69, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x0e, 0x0a, 0x02, 0x49,
	0x44, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x45,
	0x6e, 0x74, 0x49, 0x44, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x45, 0x6e, 0x74, 0x49,
	0x44, 0x12, 0x14, 0x0a, 0x05, 0x41, 0x70, 0x70, 0x49, 0x44, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x41, 0x70, 0x70, 0x49, 0x44, 0x12, 0x1e, 0x0a, 0x0a, 0x43, 0x6f, 0x69, 0x6e, 0x54,
	0x79, 0x70, 0x65, 0x49, 0x44, 0x18, 0x28, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x43, 0x6f, 0x69,
	0x6e, 0x54, 0x79, 0x70, 0x65, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08, 0x4d, 0x61, 0x78, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x32, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x4d, 0x61, 0x78, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x41, 0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x65, 0x64,
	0x18, 0x3c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x41, 0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x65,
	0x64, 0x12, 0x1a, 0x0a, 0x08, 0x43, 0x6f, 0x69, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x46, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x43, 0x6f, 0x69, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x22, 0x0a,
	0x0c, 0x44, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x18, 0x50, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x0c, 0x44, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65,
	0x73, 0x12, 0x1a, 0x0a, 0x08, 0x43, 0x6f, 0x69, 0x6e, 0x4c, 0x6f, 0x67, 0x6f, 0x18, 0x5a, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x43, 0x6f, 0x69, 0x6e, 0x4c, 0x6f, 0x67, 0x6f, 0x12, 0x1a, 0x0a,
	0x08, 0x43, 0x6f, 0x69, 0x6e, 0x55, 0x6e, 0x69, 0x74, 0x18, 0x64, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x43, 0x6f, 0x69, 0x6e, 0x55, 0x6e, 0x69, 0x74, 0x12, 0x1d, 0x0a, 0x09, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0xe8, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x09, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0xf2, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x7c, 0x0a, 0x1c, 0x41, 0x64, 0x6d, 0x69, 0x6e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x69, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x54, 0x61, 0x72, 0x67, 0x65,
	0x74, 0x41, 0x70, 0x70, 0x49, 0x44, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x54, 0x61,
	0x72, 0x67, 0x65, 0x74, 0x41, 0x70, 0x70, 0x49, 0x44, 0x12, 0x1e, 0x0a, 0x0a, 0x43, 0x6f, 0x69,
	0x6e, 0x54, 0x79, 0x70, 0x65, 0x49, 0x44, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x43,
	0x6f, 0x69, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08, 0x4d, 0x61, 0x78,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x4d, 0x61, 0x78,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x5f, 0x0a, 0x1d, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x69, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3e, 0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x69, 0x6e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x67,
	0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x63, 0x6f, 0x69, 0x6e, 0x2e, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x69, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x52, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x94, 0x01, 0x0a, 0x1c, 0x41, 0x64, 0x6d, 0x69, 0x6e,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x69, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x02, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x45, 0x6e, 0x74, 0x49, 0x44,
	0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x45, 0x6e, 0x74, 0x49, 0x44, 0x12, 0x20, 0x0a,
	0x0b, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x41, 0x70, 0x70, 0x49, 0x44, 0x18, 0x1e, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x41, 0x70, 0x70, 0x49, 0x44, 0x12,
	0x1f, 0x0a, 0x08, 0x4d, 0x61, 0x78, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x28, 0x20, 0x01, 0x28,
	0x09, 0x48, 0x00, 0x52, 0x08, 0x4d, 0x61, 0x78, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x88, 0x01, 0x01,
	0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x4d, 0x61, 0x78, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x5f, 0x0a,
	0x1d, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x69, 0x6e,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3e,
	0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x69,
	0x6e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x63,
	0x6f, 0x69, 0x6e, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f,
	0x69, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x6c,
	0x0a, 0x1a, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x69, 0x6e, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x0b,
	0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x41, 0x70, 0x70, 0x49, 0x44, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x41, 0x70, 0x70, 0x49, 0x44, 0x12, 0x16,
	0x0a, 0x06, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x14, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06,
	0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x18,
	0x1e, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x22, 0x75, 0x0a, 0x1b,
	0x41, 0x64, 0x6d, 0x69, 0x6e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x69, 0x6e, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x40, 0x0a, 0x05, 0x49,
	0x6e, 0x66, 0x6f, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x69, 0x6e, 0x73,
	0x70, 0x69, 0x72, 0x65, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x63, 0x6f, 0x69,
	0x6e, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x69, 0x6e,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x05, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x12, 0x14, 0x0a,
	0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x14, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x54, 0x6f,
	0x74, 0x61, 0x6c, 0x22, 0x44, 0x0a, 0x1c, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x43, 0x6f, 0x69, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x02, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x45, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x14, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x45, 0x6e, 0x74, 0x49, 0x44, 0x22, 0x5f, 0x0a, 0x1d, 0x41, 0x64, 0x6d,
	0x69, 0x6e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6f, 0x69, 0x6e, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3e, 0x0a, 0x04, 0x49, 0x6e,
	0x66, 0x6f, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x69, 0x6e, 0x73, 0x70, 0x69,
	0x72, 0x65, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x63, 0x6f, 0x69, 0x6e, 0x2e,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x69, 0x6e, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x52, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x32, 0xfd, 0x05, 0x0a, 0x07, 0x47,
	0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x12, 0xbc, 0x01, 0x0a, 0x15, 0x41, 0x64, 0x6d, 0x69, 0x6e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x69, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x12, 0x3c, 0x2e, 0x69, 0x6e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77,
	0x61, 0x79, 0x2e, 0x63, 0x6f, 0x69, 0x6e, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76,
	0x31, 0x2e, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x69,
	0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3d,
	0x2e, 0x69, 0x6e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79,
	0x2e, 0x63, 0x6f, 0x69, 0x6e, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76, 0x31, 0x2e,
	0x41, 0x64, 0x6d, 0x69, 0x6e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x69, 0x6e, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x26, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x20, 0x3a, 0x01, 0x2a, 0x22, 0x1b, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x64,
	0x6d, 0x69, 0x6e, 0x2f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x2f, 0x63, 0x6f, 0x69, 0x6e, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0xbc, 0x01, 0x0a, 0x15, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x69, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12,
	0x3c, 0x2e, 0x69, 0x6e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61,
	0x79, 0x2e, 0x63, 0x6f, 0x69, 0x6e, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76, 0x31,
	0x2e, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x69, 0x6e,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3d, 0x2e,
	0x69, 0x6e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e,
	0x63, 0x6f, 0x69, 0x6e, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x41,
	0x64, 0x6d, 0x69, 0x6e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x69, 0x6e, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x26, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x20, 0x3a, 0x01, 0x2a, 0x22, 0x1b, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x64, 0x6d,
	0x69, 0x6e, 0x2f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x63, 0x6f, 0x69, 0x6e, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x12, 0xb4, 0x01, 0x0a, 0x13, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x47, 0x65,
	0x74, 0x43, 0x6f, 0x69, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x12, 0x3a, 0x2e, 0x69,
	0x6e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x63,
	0x6f, 0x69, 0x6e, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x64,
	0x6d, 0x69, 0x6e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x69, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3b, 0x2e, 0x69, 0x6e, 0x73, 0x70, 0x69,
	0x72, 0x65, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x63, 0x6f, 0x69, 0x6e, 0x2e,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x47,
	0x65, 0x74, 0x43, 0x6f, 0x69, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x24, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1e, 0x3a, 0x01, 0x2a,
	0x22, 0x19, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x67, 0x65, 0x74, 0x2f,
	0x63, 0x6f, 0x69, 0x6e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x12, 0xbc, 0x01, 0x0a, 0x15,
	0x41, 0x64, 0x6d, 0x69, 0x6e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6f, 0x69, 0x6e, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x3c, 0x2e, 0x69, 0x6e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e,
	0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x63, 0x6f, 0x69, 0x6e, 0x2e, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x43, 0x6f, 0x69, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x3d, 0x2e, 0x69, 0x6e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x67, 0x61,
	0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x63, 0x6f, 0x69, 0x6e, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x43, 0x6f, 0x69, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x26, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x20, 0x3a, 0x01, 0x2a, 0x22, 0x1b, 0x2f,
	0x76, 0x31, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x2f,
	0x63, 0x6f, 0x69, 0x6e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x42, 0x42, 0x5a, 0x40, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4e, 0x70, 0x6f, 0x6f, 0x6c, 0x50, 0x6c,
	0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x6e,
	0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x69, 0x6e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2f, 0x67, 0x77, 0x2f,
	0x76, 0x31, 0x2f, 0x63, 0x6f, 0x69, 0x6e, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_npool_inspire_gw_v1_coin_config_config_proto_rawDescOnce sync.Once
	file_npool_inspire_gw_v1_coin_config_config_proto_rawDescData = file_npool_inspire_gw_v1_coin_config_config_proto_rawDesc
)

func file_npool_inspire_gw_v1_coin_config_config_proto_rawDescGZIP() []byte {
	file_npool_inspire_gw_v1_coin_config_config_proto_rawDescOnce.Do(func() {
		file_npool_inspire_gw_v1_coin_config_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_npool_inspire_gw_v1_coin_config_config_proto_rawDescData)
	})
	return file_npool_inspire_gw_v1_coin_config_config_proto_rawDescData
}

var file_npool_inspire_gw_v1_coin_config_config_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_npool_inspire_gw_v1_coin_config_config_proto_goTypes = []interface{}{
	(*CoinConfig)(nil),                    // 0: inspire.gateway.coin.config.v1.CoinConfig
	(*AdminCreateCoinConfigRequest)(nil),  // 1: inspire.gateway.coin.config.v1.AdminCreateCoinConfigRequest
	(*AdminCreateCoinConfigResponse)(nil), // 2: inspire.gateway.coin.config.v1.AdminCreateCoinConfigResponse
	(*AdminUpdateCoinConfigRequest)(nil),  // 3: inspire.gateway.coin.config.v1.AdminUpdateCoinConfigRequest
	(*AdminUpdateCoinConfigResponse)(nil), // 4: inspire.gateway.coin.config.v1.AdminUpdateCoinConfigResponse
	(*AdminGetCoinConfigsRequest)(nil),    // 5: inspire.gateway.coin.config.v1.AdminGetCoinConfigsRequest
	(*AdminGetCoinConfigsResponse)(nil),   // 6: inspire.gateway.coin.config.v1.AdminGetCoinConfigsResponse
	(*AdminDeleteCoinConfigRequest)(nil),  // 7: inspire.gateway.coin.config.v1.AdminDeleteCoinConfigRequest
	(*AdminDeleteCoinConfigResponse)(nil), // 8: inspire.gateway.coin.config.v1.AdminDeleteCoinConfigResponse
}
var file_npool_inspire_gw_v1_coin_config_config_proto_depIdxs = []int32{
	0, // 0: inspire.gateway.coin.config.v1.AdminCreateCoinConfigResponse.Info:type_name -> inspire.gateway.coin.config.v1.CoinConfig
	0, // 1: inspire.gateway.coin.config.v1.AdminUpdateCoinConfigResponse.Info:type_name -> inspire.gateway.coin.config.v1.CoinConfig
	0, // 2: inspire.gateway.coin.config.v1.AdminGetCoinConfigsResponse.Infos:type_name -> inspire.gateway.coin.config.v1.CoinConfig
	0, // 3: inspire.gateway.coin.config.v1.AdminDeleteCoinConfigResponse.Info:type_name -> inspire.gateway.coin.config.v1.CoinConfig
	1, // 4: inspire.gateway.coin.config.v1.Gateway.AdminCreateCoinConfig:input_type -> inspire.gateway.coin.config.v1.AdminCreateCoinConfigRequest
	3, // 5: inspire.gateway.coin.config.v1.Gateway.AdminUpdateCoinConfig:input_type -> inspire.gateway.coin.config.v1.AdminUpdateCoinConfigRequest
	5, // 6: inspire.gateway.coin.config.v1.Gateway.AdminGetCoinConfigs:input_type -> inspire.gateway.coin.config.v1.AdminGetCoinConfigsRequest
	7, // 7: inspire.gateway.coin.config.v1.Gateway.AdminDeleteCoinConfig:input_type -> inspire.gateway.coin.config.v1.AdminDeleteCoinConfigRequest
	2, // 8: inspire.gateway.coin.config.v1.Gateway.AdminCreateCoinConfig:output_type -> inspire.gateway.coin.config.v1.AdminCreateCoinConfigResponse
	4, // 9: inspire.gateway.coin.config.v1.Gateway.AdminUpdateCoinConfig:output_type -> inspire.gateway.coin.config.v1.AdminUpdateCoinConfigResponse
	6, // 10: inspire.gateway.coin.config.v1.Gateway.AdminGetCoinConfigs:output_type -> inspire.gateway.coin.config.v1.AdminGetCoinConfigsResponse
	8, // 11: inspire.gateway.coin.config.v1.Gateway.AdminDeleteCoinConfig:output_type -> inspire.gateway.coin.config.v1.AdminDeleteCoinConfigResponse
	8, // [8:12] is the sub-list for method output_type
	4, // [4:8] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_npool_inspire_gw_v1_coin_config_config_proto_init() }
func file_npool_inspire_gw_v1_coin_config_config_proto_init() {
	if File_npool_inspire_gw_v1_coin_config_config_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_npool_inspire_gw_v1_coin_config_config_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CoinConfig); i {
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
		file_npool_inspire_gw_v1_coin_config_config_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdminCreateCoinConfigRequest); i {
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
		file_npool_inspire_gw_v1_coin_config_config_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdminCreateCoinConfigResponse); i {
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
		file_npool_inspire_gw_v1_coin_config_config_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdminUpdateCoinConfigRequest); i {
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
		file_npool_inspire_gw_v1_coin_config_config_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdminUpdateCoinConfigResponse); i {
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
		file_npool_inspire_gw_v1_coin_config_config_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdminGetCoinConfigsRequest); i {
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
		file_npool_inspire_gw_v1_coin_config_config_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdminGetCoinConfigsResponse); i {
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
		file_npool_inspire_gw_v1_coin_config_config_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdminDeleteCoinConfigRequest); i {
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
		file_npool_inspire_gw_v1_coin_config_config_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdminDeleteCoinConfigResponse); i {
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
	file_npool_inspire_gw_v1_coin_config_config_proto_msgTypes[3].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_npool_inspire_gw_v1_coin_config_config_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_npool_inspire_gw_v1_coin_config_config_proto_goTypes,
		DependencyIndexes: file_npool_inspire_gw_v1_coin_config_config_proto_depIdxs,
		MessageInfos:      file_npool_inspire_gw_v1_coin_config_config_proto_msgTypes,
	}.Build()
	File_npool_inspire_gw_v1_coin_config_config_proto = out.File
	file_npool_inspire_gw_v1_coin_config_config_proto_rawDesc = nil
	file_npool_inspire_gw_v1_coin_config_config_proto_goTypes = nil
	file_npool_inspire_gw_v1_coin_config_config_proto_depIdxs = nil
}
