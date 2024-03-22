// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.18.1
// source: npool/miningpool/gw/v1/orderuser/orderuser.proto

package orderuser

import (
	orderuser "github.com/NpoolPlatform/message/npool/miningpool/mw/v1/orderuser"
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

type GetOrderUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EntID string `protobuf:"bytes,10,opt,name=EntID,proto3" json:"EntID,omitempty"`
}

func (x *GetOrderUserRequest) Reset() {
	*x = GetOrderUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_miningpool_gw_v1_orderuser_orderuser_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetOrderUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOrderUserRequest) ProtoMessage() {}

func (x *GetOrderUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_miningpool_gw_v1_orderuser_orderuser_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOrderUserRequest.ProtoReflect.Descriptor instead.
func (*GetOrderUserRequest) Descriptor() ([]byte, []int) {
	return file_npool_miningpool_gw_v1_orderuser_orderuser_proto_rawDescGZIP(), []int{0}
}

func (x *GetOrderUserRequest) GetEntID() string {
	if x != nil {
		return x.EntID
	}
	return ""
}

type GetOrderUserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *orderuser.OrderUser `protobuf:"bytes,10,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (x *GetOrderUserResponse) Reset() {
	*x = GetOrderUserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_miningpool_gw_v1_orderuser_orderuser_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetOrderUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOrderUserResponse) ProtoMessage() {}

func (x *GetOrderUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_miningpool_gw_v1_orderuser_orderuser_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOrderUserResponse.ProtoReflect.Descriptor instead.
func (*GetOrderUserResponse) Descriptor() ([]byte, []int) {
	return file_npool_miningpool_gw_v1_orderuser_orderuser_proto_rawDescGZIP(), []int{1}
}

func (x *GetOrderUserResponse) GetInfo() *orderuser.OrderUser {
	if x != nil {
		return x.Info
	}
	return nil
}

type GetOrderUsersByOrderRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderID string `protobuf:"bytes,10,opt,name=OrderID,proto3" json:"OrderID,omitempty"`
	Offset  int32  `protobuf:"varint,20,opt,name=Offset,proto3" json:"Offset,omitempty"`
	Limit   int32  `protobuf:"varint,30,opt,name=Limit,proto3" json:"Limit,omitempty"`
}

func (x *GetOrderUsersByOrderRequest) Reset() {
	*x = GetOrderUsersByOrderRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_miningpool_gw_v1_orderuser_orderuser_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetOrderUsersByOrderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOrderUsersByOrderRequest) ProtoMessage() {}

func (x *GetOrderUsersByOrderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_miningpool_gw_v1_orderuser_orderuser_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOrderUsersByOrderRequest.ProtoReflect.Descriptor instead.
func (*GetOrderUsersByOrderRequest) Descriptor() ([]byte, []int) {
	return file_npool_miningpool_gw_v1_orderuser_orderuser_proto_rawDescGZIP(), []int{2}
}

func (x *GetOrderUsersByOrderRequest) GetOrderID() string {
	if x != nil {
		return x.OrderID
	}
	return ""
}

func (x *GetOrderUsersByOrderRequest) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *GetOrderUsersByOrderRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type GetOrderUsersByOrderResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Infos []*orderuser.OrderUser `protobuf:"bytes,10,rep,name=Infos,proto3" json:"Infos,omitempty"`
	Total uint32                 `protobuf:"varint,20,opt,name=Total,proto3" json:"Total,omitempty"`
}

func (x *GetOrderUsersByOrderResponse) Reset() {
	*x = GetOrderUsersByOrderResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_miningpool_gw_v1_orderuser_orderuser_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetOrderUsersByOrderResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOrderUsersByOrderResponse) ProtoMessage() {}

func (x *GetOrderUsersByOrderResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_miningpool_gw_v1_orderuser_orderuser_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOrderUsersByOrderResponse.ProtoReflect.Descriptor instead.
func (*GetOrderUsersByOrderResponse) Descriptor() ([]byte, []int) {
	return file_npool_miningpool_gw_v1_orderuser_orderuser_proto_rawDescGZIP(), []int{3}
}

func (x *GetOrderUsersByOrderResponse) GetInfos() []*orderuser.OrderUser {
	if x != nil {
		return x.Infos
	}
	return nil
}

func (x *GetOrderUsersByOrderResponse) GetTotal() uint32 {
	if x != nil {
		return x.Total
	}
	return 0
}

type SetupRevenueAddressRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EntID          string `protobuf:"bytes,10,opt,name=EntID,proto3" json:"EntID,omitempty"`
	RevenueAddress string `protobuf:"bytes,20,opt,name=RevenueAddress,proto3" json:"RevenueAddress,omitempty"`
}

func (x *SetupRevenueAddressRequest) Reset() {
	*x = SetupRevenueAddressRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_miningpool_gw_v1_orderuser_orderuser_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetupRevenueAddressRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetupRevenueAddressRequest) ProtoMessage() {}

func (x *SetupRevenueAddressRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_miningpool_gw_v1_orderuser_orderuser_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetupRevenueAddressRequest.ProtoReflect.Descriptor instead.
func (*SetupRevenueAddressRequest) Descriptor() ([]byte, []int) {
	return file_npool_miningpool_gw_v1_orderuser_orderuser_proto_rawDescGZIP(), []int{4}
}

func (x *SetupRevenueAddressRequest) GetEntID() string {
	if x != nil {
		return x.EntID
	}
	return ""
}

func (x *SetupRevenueAddressRequest) GetRevenueAddress() string {
	if x != nil {
		return x.RevenueAddress
	}
	return ""
}

type SetupRevenueAddressResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *orderuser.OrderUser `protobuf:"bytes,10,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (x *SetupRevenueAddressResponse) Reset() {
	*x = SetupRevenueAddressResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_miningpool_gw_v1_orderuser_orderuser_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetupRevenueAddressResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetupRevenueAddressResponse) ProtoMessage() {}

func (x *SetupRevenueAddressResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_miningpool_gw_v1_orderuser_orderuser_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetupRevenueAddressResponse.ProtoReflect.Descriptor instead.
func (*SetupRevenueAddressResponse) Descriptor() ([]byte, []int) {
	return file_npool_miningpool_gw_v1_orderuser_orderuser_proto_rawDescGZIP(), []int{5}
}

func (x *SetupRevenueAddressResponse) GetInfo() *orderuser.OrderUser {
	if x != nil {
		return x.Info
	}
	return nil
}

type SetupAutoPayRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EntID   string `protobuf:"bytes,10,opt,name=EntID,proto3" json:"EntID,omitempty"`
	AutoPay bool   `protobuf:"varint,20,opt,name=AutoPay,proto3" json:"AutoPay,omitempty"`
}

func (x *SetupAutoPayRequest) Reset() {
	*x = SetupAutoPayRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_miningpool_gw_v1_orderuser_orderuser_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetupAutoPayRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetupAutoPayRequest) ProtoMessage() {}

func (x *SetupAutoPayRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_miningpool_gw_v1_orderuser_orderuser_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetupAutoPayRequest.ProtoReflect.Descriptor instead.
func (*SetupAutoPayRequest) Descriptor() ([]byte, []int) {
	return file_npool_miningpool_gw_v1_orderuser_orderuser_proto_rawDescGZIP(), []int{6}
}

func (x *SetupAutoPayRequest) GetEntID() string {
	if x != nil {
		return x.EntID
	}
	return ""
}

func (x *SetupAutoPayRequest) GetAutoPay() bool {
	if x != nil {
		return x.AutoPay
	}
	return false
}

type SetupAutoPayResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *orderuser.OrderUser `protobuf:"bytes,10,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (x *SetupAutoPayResponse) Reset() {
	*x = SetupAutoPayResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_miningpool_gw_v1_orderuser_orderuser_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetupAutoPayResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetupAutoPayResponse) ProtoMessage() {}

func (x *SetupAutoPayResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_miningpool_gw_v1_orderuser_orderuser_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetupAutoPayResponse.ProtoReflect.Descriptor instead.
func (*SetupAutoPayResponse) Descriptor() ([]byte, []int) {
	return file_npool_miningpool_gw_v1_orderuser_orderuser_proto_rawDescGZIP(), []int{7}
}

func (x *SetupAutoPayResponse) GetInfo() *orderuser.OrderUser {
	if x != nil {
		return x.Info
	}
	return nil
}

var File_npool_miningpool_gw_v1_orderuser_orderuser_proto protoreflect.FileDescriptor

var file_npool_miningpool_gw_v1_orderuser_orderuser_proto_rawDesc = []byte{
	0x0a, 0x30, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x6d, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x70, 0x6f,
	0x6f, 0x6c, 0x2f, 0x67, 0x77, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x75, 0x73,
	0x65, 0x72, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x1f, 0x6d, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x70, 0x6f, 0x6f, 0x6c, 0x2e, 0x67,
	0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x75, 0x73, 0x65, 0x72,
	0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x30, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x6d, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x70,
	0x6f, 0x6f, 0x6c, 0x2f, 0x6d, 0x77, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x75,
	0x73, 0x65, 0x72, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x2b, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x55,
	0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x45, 0x6e,
	0x74, 0x49, 0x44, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x45, 0x6e, 0x74, 0x49, 0x44,
	0x22, 0x59, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x41, 0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x6d, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x70,
	0x6f, 0x6f, 0x6c, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x65, 0x0a, 0x1b, 0x47,
	0x65, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x73, 0x42, 0x79, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x49, 0x44, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x14,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05,
	0x4c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x4c, 0x69, 0x6d,
	0x69, 0x74, 0x22, 0x79, 0x0a, 0x1c, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x55, 0x73,
	0x65, 0x72, 0x73, 0x42, 0x79, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x43, 0x0a, 0x05, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x2d, 0x2e, 0x6d, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x70, 0x6f, 0x6f, 0x6c, 0x2e, 0x6d,
	0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x75,
	0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72,
	0x52, 0x05, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c,
	0x18, 0x14, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x22, 0x5a, 0x0a,
	0x1a, 0x53, 0x65, 0x74, 0x75, 0x70, 0x52, 0x65, 0x76, 0x65, 0x6e, 0x75, 0x65, 0x41, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x45,
	0x6e, 0x74, 0x49, 0x44, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x45, 0x6e, 0x74, 0x49,
	0x44, 0x12, 0x26, 0x0a, 0x0e, 0x52, 0x65, 0x76, 0x65, 0x6e, 0x75, 0x65, 0x41, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x52, 0x65, 0x76, 0x65, 0x6e,
	0x75, 0x65, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0x60, 0x0a, 0x1b, 0x53, 0x65, 0x74,
	0x75, 0x70, 0x52, 0x65, 0x76, 0x65, 0x6e, 0x75, 0x65, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x41, 0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x6d, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x70,
	0x6f, 0x6f, 0x6c, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x45, 0x0a, 0x13, 0x53,
	0x65, 0x74, 0x75, 0x70, 0x41, 0x75, 0x74, 0x6f, 0x50, 0x61, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x45, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x45, 0x6e, 0x74, 0x49, 0x44, 0x12, 0x18, 0x0a, 0x07, 0x41, 0x75, 0x74, 0x6f,
	0x50, 0x61, 0x79, 0x18, 0x14, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x41, 0x75, 0x74, 0x6f, 0x50,
	0x61, 0x79, 0x22, 0x59, 0x0a, 0x14, 0x53, 0x65, 0x74, 0x75, 0x70, 0x41, 0x75, 0x74, 0x6f, 0x50,
	0x61, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x41, 0x0a, 0x04, 0x49, 0x6e,
	0x66, 0x6f, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x6d, 0x69, 0x6e, 0x69, 0x6e,
	0x67, 0x70, 0x6f, 0x6f, 0x6c, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65,
	0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x32, 0xb7, 0x05,
	0x0a, 0x07, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x12, 0x99, 0x01, 0x0a, 0x0c, 0x47, 0x65,
	0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x12, 0x34, 0x2e, 0x6d, 0x69, 0x6e,
	0x69, 0x6e, 0x67, 0x70, 0x6f, 0x6f, 0x6c, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x35, 0x2e, 0x6d, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x70, 0x6f, 0x6f, 0x6c, 0x2e, 0x67, 0x61,
	0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x75, 0x73, 0x65, 0x72, 0x2e,
	0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x16, 0x3a,
	0x01, 0x2a, 0x22, 0x11, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x65, 0x74, 0x2f, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x75, 0x73, 0x65, 0x72, 0x12, 0xbb, 0x01, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x73, 0x42, 0x79, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x3c,
	0x2e, 0x6d, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x70, 0x6f, 0x6f, 0x6c, 0x2e, 0x67, 0x61, 0x74, 0x65,
	0x77, 0x61, 0x79, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31,
	0x2e, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x73, 0x42, 0x79,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3d, 0x2e, 0x6d,
	0x69, 0x6e, 0x69, 0x6e, 0x67, 0x70, 0x6f, 0x6f, 0x6c, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61,
	0x79, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47,
	0x65, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x73, 0x42, 0x79, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x26, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x20, 0x3a, 0x01, 0x2a, 0x22, 0x1b, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x65, 0x74, 0x2f,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2f, 0x62, 0x79, 0x2f, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x12, 0xb5, 0x01, 0x0a, 0x13, 0x53, 0x65, 0x74, 0x75, 0x70, 0x52, 0x65, 0x76,
	0x65, 0x6e, 0x75, 0x65, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x3b, 0x2e, 0x6d, 0x69,
	0x6e, 0x69, 0x6e, 0x67, 0x70, 0x6f, 0x6f, 0x6c, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79,
	0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65,
	0x74, 0x75, 0x70, 0x52, 0x65, 0x76, 0x65, 0x6e, 0x75, 0x65, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3c, 0x2e, 0x6d, 0x69, 0x6e, 0x69, 0x6e,
	0x67, 0x70, 0x6f, 0x6f, 0x6c, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x74, 0x75, 0x70,
	0x52, 0x65, 0x76, 0x65, 0x6e, 0x75, 0x65, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x23, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1d, 0x3a, 0x01,
	0x2a, 0x22, 0x18, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x74, 0x75, 0x70, 0x2f, 0x72, 0x65, 0x76,
	0x65, 0x6e, 0x75, 0x65, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x99, 0x01, 0x0a, 0x0c,
	0x53, 0x65, 0x74, 0x75, 0x70, 0x41, 0x75, 0x74, 0x6f, 0x50, 0x61, 0x79, 0x12, 0x34, 0x2e, 0x6d,
	0x69, 0x6e, 0x69, 0x6e, 0x67, 0x70, 0x6f, 0x6f, 0x6c, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61,
	0x79, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x53,
	0x65, 0x74, 0x75, 0x70, 0x41, 0x75, 0x74, 0x6f, 0x50, 0x61, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x35, 0x2e, 0x6d, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x70, 0x6f, 0x6f, 0x6c, 0x2e,
	0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x75, 0x73, 0x65,
	0x72, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x74, 0x75, 0x70, 0x41, 0x75, 0x74, 0x6f, 0x50, 0x61,
	0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1c, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x16, 0x3a, 0x01, 0x2a, 0x22, 0x11, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x74, 0x75, 0x70, 0x2f,
	0x61, 0x75, 0x74, 0x6f, 0x70, 0x61, 0x79, 0x42, 0x43, 0x5a, 0x41, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4e, 0x70, 0x6f, 0x6f, 0x6c, 0x50, 0x6c, 0x61, 0x74, 0x66,
	0x6f, 0x72, 0x6d, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x6e, 0x70, 0x6f, 0x6f,
	0x6c, 0x2f, 0x6d, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x67, 0x77, 0x2f,
	0x76, 0x31, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x75, 0x73, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_npool_miningpool_gw_v1_orderuser_orderuser_proto_rawDescOnce sync.Once
	file_npool_miningpool_gw_v1_orderuser_orderuser_proto_rawDescData = file_npool_miningpool_gw_v1_orderuser_orderuser_proto_rawDesc
)

func file_npool_miningpool_gw_v1_orderuser_orderuser_proto_rawDescGZIP() []byte {
	file_npool_miningpool_gw_v1_orderuser_orderuser_proto_rawDescOnce.Do(func() {
		file_npool_miningpool_gw_v1_orderuser_orderuser_proto_rawDescData = protoimpl.X.CompressGZIP(file_npool_miningpool_gw_v1_orderuser_orderuser_proto_rawDescData)
	})
	return file_npool_miningpool_gw_v1_orderuser_orderuser_proto_rawDescData
}

var file_npool_miningpool_gw_v1_orderuser_orderuser_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_npool_miningpool_gw_v1_orderuser_orderuser_proto_goTypes = []interface{}{
	(*GetOrderUserRequest)(nil),          // 0: miningpool.gateway.orderuser.v1.GetOrderUserRequest
	(*GetOrderUserResponse)(nil),         // 1: miningpool.gateway.orderuser.v1.GetOrderUserResponse
	(*GetOrderUsersByOrderRequest)(nil),  // 2: miningpool.gateway.orderuser.v1.GetOrderUsersByOrderRequest
	(*GetOrderUsersByOrderResponse)(nil), // 3: miningpool.gateway.orderuser.v1.GetOrderUsersByOrderResponse
	(*SetupRevenueAddressRequest)(nil),   // 4: miningpool.gateway.orderuser.v1.SetupRevenueAddressRequest
	(*SetupRevenueAddressResponse)(nil),  // 5: miningpool.gateway.orderuser.v1.SetupRevenueAddressResponse
	(*SetupAutoPayRequest)(nil),          // 6: miningpool.gateway.orderuser.v1.SetupAutoPayRequest
	(*SetupAutoPayResponse)(nil),         // 7: miningpool.gateway.orderuser.v1.SetupAutoPayResponse
	(*orderuser.OrderUser)(nil),          // 8: miningpool.middleware.orderuser.v1.OrderUser
}
var file_npool_miningpool_gw_v1_orderuser_orderuser_proto_depIdxs = []int32{
	8, // 0: miningpool.gateway.orderuser.v1.GetOrderUserResponse.Info:type_name -> miningpool.middleware.orderuser.v1.OrderUser
	8, // 1: miningpool.gateway.orderuser.v1.GetOrderUsersByOrderResponse.Infos:type_name -> miningpool.middleware.orderuser.v1.OrderUser
	8, // 2: miningpool.gateway.orderuser.v1.SetupRevenueAddressResponse.Info:type_name -> miningpool.middleware.orderuser.v1.OrderUser
	8, // 3: miningpool.gateway.orderuser.v1.SetupAutoPayResponse.Info:type_name -> miningpool.middleware.orderuser.v1.OrderUser
	0, // 4: miningpool.gateway.orderuser.v1.Gateway.GetOrderUser:input_type -> miningpool.gateway.orderuser.v1.GetOrderUserRequest
	2, // 5: miningpool.gateway.orderuser.v1.Gateway.GetOrderUsersByOrder:input_type -> miningpool.gateway.orderuser.v1.GetOrderUsersByOrderRequest
	4, // 6: miningpool.gateway.orderuser.v1.Gateway.SetupRevenueAddress:input_type -> miningpool.gateway.orderuser.v1.SetupRevenueAddressRequest
	6, // 7: miningpool.gateway.orderuser.v1.Gateway.SetupAutoPay:input_type -> miningpool.gateway.orderuser.v1.SetupAutoPayRequest
	1, // 8: miningpool.gateway.orderuser.v1.Gateway.GetOrderUser:output_type -> miningpool.gateway.orderuser.v1.GetOrderUserResponse
	3, // 9: miningpool.gateway.orderuser.v1.Gateway.GetOrderUsersByOrder:output_type -> miningpool.gateway.orderuser.v1.GetOrderUsersByOrderResponse
	5, // 10: miningpool.gateway.orderuser.v1.Gateway.SetupRevenueAddress:output_type -> miningpool.gateway.orderuser.v1.SetupRevenueAddressResponse
	7, // 11: miningpool.gateway.orderuser.v1.Gateway.SetupAutoPay:output_type -> miningpool.gateway.orderuser.v1.SetupAutoPayResponse
	8, // [8:12] is the sub-list for method output_type
	4, // [4:8] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_npool_miningpool_gw_v1_orderuser_orderuser_proto_init() }
func file_npool_miningpool_gw_v1_orderuser_orderuser_proto_init() {
	if File_npool_miningpool_gw_v1_orderuser_orderuser_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_npool_miningpool_gw_v1_orderuser_orderuser_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetOrderUserRequest); i {
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
		file_npool_miningpool_gw_v1_orderuser_orderuser_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetOrderUserResponse); i {
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
		file_npool_miningpool_gw_v1_orderuser_orderuser_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetOrderUsersByOrderRequest); i {
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
		file_npool_miningpool_gw_v1_orderuser_orderuser_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetOrderUsersByOrderResponse); i {
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
		file_npool_miningpool_gw_v1_orderuser_orderuser_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetupRevenueAddressRequest); i {
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
		file_npool_miningpool_gw_v1_orderuser_orderuser_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetupRevenueAddressResponse); i {
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
		file_npool_miningpool_gw_v1_orderuser_orderuser_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetupAutoPayRequest); i {
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
		file_npool_miningpool_gw_v1_orderuser_orderuser_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetupAutoPayResponse); i {
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
			RawDescriptor: file_npool_miningpool_gw_v1_orderuser_orderuser_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_npool_miningpool_gw_v1_orderuser_orderuser_proto_goTypes,
		DependencyIndexes: file_npool_miningpool_gw_v1_orderuser_orderuser_proto_depIdxs,
		MessageInfos:      file_npool_miningpool_gw_v1_orderuser_orderuser_proto_msgTypes,
	}.Build()
	File_npool_miningpool_gw_v1_orderuser_orderuser_proto = out.File
	file_npool_miningpool_gw_v1_orderuser_orderuser_proto_rawDesc = nil
	file_npool_miningpool_gw_v1_orderuser_orderuser_proto_goTypes = nil
	file_npool_miningpool_gw_v1_orderuser_orderuser_proto_depIdxs = nil
}
