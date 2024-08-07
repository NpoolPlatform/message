// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v3.18.1
// source: npool/order/mw/v1/powerrental/outofgas/outofgas.proto

package outofgas

import (
	_ "github.com/NpoolPlatform/message/npool/basetypes/good/v1"
	_ "github.com/NpoolPlatform/message/npool/basetypes/v1"
	outofgas "github.com/NpoolPlatform/message/npool/order/mw/v1/outofgas"
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

type OutOfGasReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID      *uint32 `protobuf:"varint,10,opt,name=ID,proto3,oneof" json:"ID,omitempty"`
	EntID   *string `protobuf:"bytes,20,opt,name=EntID,proto3,oneof" json:"EntID,omitempty"`
	OrderID *string `protobuf:"bytes,30,opt,name=OrderID,proto3,oneof" json:"OrderID,omitempty"`
	StartAt *uint32 `protobuf:"varint,40,opt,name=StartAt,proto3,oneof" json:"StartAt,omitempty"`
	EndAt   *uint32 `protobuf:"varint,50,opt,name=EndAt,proto3,oneof" json:"EndAt,omitempty"`
}

func (x *OutOfGasReq) Reset() {
	*x = OutOfGasReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_order_mw_v1_powerrental_outofgas_outofgas_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OutOfGasReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OutOfGasReq) ProtoMessage() {}

func (x *OutOfGasReq) ProtoReflect() protoreflect.Message {
	mi := &file_npool_order_mw_v1_powerrental_outofgas_outofgas_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OutOfGasReq.ProtoReflect.Descriptor instead.
func (*OutOfGasReq) Descriptor() ([]byte, []int) {
	return file_npool_order_mw_v1_powerrental_outofgas_outofgas_proto_rawDescGZIP(), []int{0}
}

func (x *OutOfGasReq) GetID() uint32 {
	if x != nil && x.ID != nil {
		return *x.ID
	}
	return 0
}

func (x *OutOfGasReq) GetEntID() string {
	if x != nil && x.EntID != nil {
		return *x.EntID
	}
	return ""
}

func (x *OutOfGasReq) GetOrderID() string {
	if x != nil && x.OrderID != nil {
		return *x.OrderID
	}
	return ""
}

func (x *OutOfGasReq) GetStartAt() uint32 {
	if x != nil && x.StartAt != nil {
		return *x.StartAt
	}
	return 0
}

func (x *OutOfGasReq) GetEndAt() uint32 {
	if x != nil && x.EndAt != nil {
		return *x.EndAt
	}
	return 0
}

type CreateOutOfGasRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *OutOfGasReq `protobuf:"bytes,10,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (x *CreateOutOfGasRequest) Reset() {
	*x = CreateOutOfGasRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_order_mw_v1_powerrental_outofgas_outofgas_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateOutOfGasRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateOutOfGasRequest) ProtoMessage() {}

func (x *CreateOutOfGasRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_order_mw_v1_powerrental_outofgas_outofgas_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateOutOfGasRequest.ProtoReflect.Descriptor instead.
func (*CreateOutOfGasRequest) Descriptor() ([]byte, []int) {
	return file_npool_order_mw_v1_powerrental_outofgas_outofgas_proto_rawDescGZIP(), []int{1}
}

func (x *CreateOutOfGasRequest) GetInfo() *OutOfGasReq {
	if x != nil {
		return x.Info
	}
	return nil
}

type CreateOutOfGasResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *outofgas.OutOfGas `protobuf:"bytes,10,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (x *CreateOutOfGasResponse) Reset() {
	*x = CreateOutOfGasResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_order_mw_v1_powerrental_outofgas_outofgas_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateOutOfGasResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateOutOfGasResponse) ProtoMessage() {}

func (x *CreateOutOfGasResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_order_mw_v1_powerrental_outofgas_outofgas_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateOutOfGasResponse.ProtoReflect.Descriptor instead.
func (*CreateOutOfGasResponse) Descriptor() ([]byte, []int) {
	return file_npool_order_mw_v1_powerrental_outofgas_outofgas_proto_rawDescGZIP(), []int{2}
}

func (x *CreateOutOfGasResponse) GetInfo() *outofgas.OutOfGas {
	if x != nil {
		return x.Info
	}
	return nil
}

type UpdateOutOfGasRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *OutOfGasReq `protobuf:"bytes,10,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (x *UpdateOutOfGasRequest) Reset() {
	*x = UpdateOutOfGasRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_order_mw_v1_powerrental_outofgas_outofgas_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateOutOfGasRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateOutOfGasRequest) ProtoMessage() {}

func (x *UpdateOutOfGasRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_order_mw_v1_powerrental_outofgas_outofgas_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateOutOfGasRequest.ProtoReflect.Descriptor instead.
func (*UpdateOutOfGasRequest) Descriptor() ([]byte, []int) {
	return file_npool_order_mw_v1_powerrental_outofgas_outofgas_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateOutOfGasRequest) GetInfo() *OutOfGasReq {
	if x != nil {
		return x.Info
	}
	return nil
}

type UpdateOutOfGasResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *outofgas.OutOfGas `protobuf:"bytes,10,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (x *UpdateOutOfGasResponse) Reset() {
	*x = UpdateOutOfGasResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_order_mw_v1_powerrental_outofgas_outofgas_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateOutOfGasResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateOutOfGasResponse) ProtoMessage() {}

func (x *UpdateOutOfGasResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_order_mw_v1_powerrental_outofgas_outofgas_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateOutOfGasResponse.ProtoReflect.Descriptor instead.
func (*UpdateOutOfGasResponse) Descriptor() ([]byte, []int) {
	return file_npool_order_mw_v1_powerrental_outofgas_outofgas_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateOutOfGasResponse) GetInfo() *outofgas.OutOfGas {
	if x != nil {
		return x.Info
	}
	return nil
}

type DeleteOutOfGasRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *OutOfGasReq `protobuf:"bytes,10,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (x *DeleteOutOfGasRequest) Reset() {
	*x = DeleteOutOfGasRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_order_mw_v1_powerrental_outofgas_outofgas_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteOutOfGasRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteOutOfGasRequest) ProtoMessage() {}

func (x *DeleteOutOfGasRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_order_mw_v1_powerrental_outofgas_outofgas_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteOutOfGasRequest.ProtoReflect.Descriptor instead.
func (*DeleteOutOfGasRequest) Descriptor() ([]byte, []int) {
	return file_npool_order_mw_v1_powerrental_outofgas_outofgas_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteOutOfGasRequest) GetInfo() *OutOfGasReq {
	if x != nil {
		return x.Info
	}
	return nil
}

type DeleteOutOfGasResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *outofgas.OutOfGas `protobuf:"bytes,10,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (x *DeleteOutOfGasResponse) Reset() {
	*x = DeleteOutOfGasResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_order_mw_v1_powerrental_outofgas_outofgas_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteOutOfGasResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteOutOfGasResponse) ProtoMessage() {}

func (x *DeleteOutOfGasResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_order_mw_v1_powerrental_outofgas_outofgas_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteOutOfGasResponse.ProtoReflect.Descriptor instead.
func (*DeleteOutOfGasResponse) Descriptor() ([]byte, []int) {
	return file_npool_order_mw_v1_powerrental_outofgas_outofgas_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteOutOfGasResponse) GetInfo() *outofgas.OutOfGas {
	if x != nil {
		return x.Info
	}
	return nil
}

var File_npool_order_mw_v1_powerrental_outofgas_outofgas_proto protoreflect.FileDescriptor

var file_npool_order_mw_v1_powerrental_outofgas_outofgas_proto_rawDesc = []byte{
	0x0a, 0x35, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2f, 0x6d, 0x77,
	0x2f, 0x76, 0x31, 0x2f, 0x70, 0x6f, 0x77, 0x65, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x2f,
	0x6f, 0x75, 0x74, 0x6f, 0x66, 0x67, 0x61, 0x73, 0x2f, 0x6f, 0x75, 0x74, 0x6f, 0x66, 0x67, 0x61,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x28, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x6d,
	0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x70, 0x6f, 0x77, 0x65, 0x72, 0x72,
	0x65, 0x6e, 0x74, 0x61, 0x6c, 0x2e, 0x6f, 0x75, 0x74, 0x6f, 0x66, 0x67, 0x61, 0x73, 0x2e, 0x76,
	0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1e, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x23, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x2f, 0x67, 0x6f, 0x6f, 0x64, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x29, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x2f, 0x6d, 0x77, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x75, 0x74, 0x6f, 0x66, 0x67, 0x61, 0x73,
	0x2f, 0x6f, 0x75, 0x74, 0x6f, 0x66, 0x67, 0x61, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xc9, 0x01, 0x0a, 0x0b, 0x4f, 0x75, 0x74, 0x4f, 0x66, 0x47, 0x61, 0x73, 0x52, 0x65, 0x71, 0x12,
	0x13, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0d, 0x48, 0x00, 0x52, 0x02, 0x49,
	0x44, 0x88, 0x01, 0x01, 0x12, 0x19, 0x0a, 0x05, 0x45, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x14, 0x20,
	0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x05, 0x45, 0x6e, 0x74, 0x49, 0x44, 0x88, 0x01, 0x01, 0x12,
	0x1d, 0x0a, 0x07, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x44, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x09,
	0x48, 0x02, 0x52, 0x07, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x44, 0x88, 0x01, 0x01, 0x12, 0x1d,
	0x0a, 0x07, 0x53, 0x74, 0x61, 0x72, 0x74, 0x41, 0x74, 0x18, 0x28, 0x20, 0x01, 0x28, 0x0d, 0x48,
	0x03, 0x52, 0x07, 0x53, 0x74, 0x61, 0x72, 0x74, 0x41, 0x74, 0x88, 0x01, 0x01, 0x12, 0x19, 0x0a,
	0x05, 0x45, 0x6e, 0x64, 0x41, 0x74, 0x18, 0x32, 0x20, 0x01, 0x28, 0x0d, 0x48, 0x04, 0x52, 0x05,
	0x45, 0x6e, 0x64, 0x41, 0x74, 0x88, 0x01, 0x01, 0x42, 0x05, 0x0a, 0x03, 0x5f, 0x49, 0x44, 0x42,
	0x08, 0x0a, 0x06, 0x5f, 0x45, 0x6e, 0x74, 0x49, 0x44, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x49, 0x44, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x53, 0x74, 0x61, 0x72, 0x74, 0x41,
	0x74, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x45, 0x6e, 0x64, 0x41, 0x74, 0x22, 0x62, 0x0a, 0x15, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x75, 0x74, 0x4f, 0x66, 0x47, 0x61, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x49, 0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x35, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c,
	0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x70, 0x6f, 0x77, 0x65, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x61,
	0x6c, 0x2e, 0x6f, 0x75, 0x74, 0x6f, 0x66, 0x67, 0x61, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x75,
	0x74, 0x4f, 0x66, 0x47, 0x61, 0x73, 0x52, 0x65, 0x71, 0x52, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x22,
	0x54, 0x0a, 0x16, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x75, 0x74, 0x4f, 0x66, 0x47, 0x61,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3a, 0x0a, 0x04, 0x49, 0x6e, 0x66,
	0x6f, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e,
	0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x6f, 0x75, 0x74, 0x6f, 0x66,
	0x67, 0x61, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x75, 0x74, 0x4f, 0x66, 0x47, 0x61, 0x73, 0x52,
	0x04, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x62, 0x0a, 0x15, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4f,
	0x75, 0x74, 0x4f, 0x66, 0x47, 0x61, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x49,
	0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x35, 0x2e, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e,
	0x70, 0x6f, 0x77, 0x65, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x2e, 0x6f, 0x75, 0x74, 0x6f,
	0x66, 0x67, 0x61, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x75, 0x74, 0x4f, 0x66, 0x47, 0x61, 0x73,
	0x52, 0x65, 0x71, 0x52, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x54, 0x0a, 0x16, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x4f, 0x75, 0x74, 0x4f, 0x66, 0x47, 0x61, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x3a, 0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x26, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65,
	0x77, 0x61, 0x72, 0x65, 0x2e, 0x6f, 0x75, 0x74, 0x6f, 0x66, 0x67, 0x61, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x4f, 0x75, 0x74, 0x4f, 0x66, 0x47, 0x61, 0x73, 0x52, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x22,
	0x62, 0x0a, 0x15, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4f, 0x75, 0x74, 0x4f, 0x66, 0x47, 0x61,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x49, 0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x35, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x6d,
	0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x70, 0x6f, 0x77, 0x65, 0x72, 0x72,
	0x65, 0x6e, 0x74, 0x61, 0x6c, 0x2e, 0x6f, 0x75, 0x74, 0x6f, 0x66, 0x67, 0x61, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x4f, 0x75, 0x74, 0x4f, 0x66, 0x47, 0x61, 0x73, 0x52, 0x65, 0x71, 0x52, 0x04, 0x49,
	0x6e, 0x66, 0x6f, 0x22, 0x54, 0x0a, 0x16, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4f, 0x75, 0x74,
	0x4f, 0x66, 0x47, 0x61, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3a, 0x0a,
	0x04, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x6f,
	0x75, 0x74, 0x6f, 0x66, 0x67, 0x61, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x75, 0x74, 0x4f, 0x66,
	0x47, 0x61, 0x73, 0x52, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x32, 0xea, 0x04, 0x0a, 0x0a, 0x4d, 0x69,
	0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x12, 0xc7, 0x01, 0x0a, 0x0e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x4f, 0x75, 0x74, 0x4f, 0x66, 0x47, 0x61, 0x73, 0x12, 0x3f, 0x2e, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x70,
	0x6f, 0x77, 0x65, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x2e, 0x6f, 0x75, 0x74, 0x6f, 0x66,
	0x67, 0x61, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x75, 0x74,
	0x4f, 0x66, 0x47, 0x61, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x40, 0x2e, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e,
	0x70, 0x6f, 0x77, 0x65, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x2e, 0x6f, 0x75, 0x74, 0x6f,
	0x66, 0x67, 0x61, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x75,
	0x74, 0x4f, 0x66, 0x47, 0x61, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x32,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2c, 0x3a, 0x01, 0x2a, 0x22, 0x27, 0x2f, 0x76, 0x31, 0x2f, 0x70,
	0x6f, 0x77, 0x65, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x5f, 0x6f, 0x75, 0x74, 0x6f, 0x66,
	0x67, 0x61, 0x73, 0x5f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x75, 0x74, 0x4f, 0x66, 0x47,
	0x61, 0x73, 0x12, 0xc7, 0x01, 0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4f, 0x75, 0x74,
	0x4f, 0x66, 0x47, 0x61, 0x73, 0x12, 0x3f, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x6d, 0x69,
	0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x70, 0x6f, 0x77, 0x65, 0x72, 0x72, 0x65,
	0x6e, 0x74, 0x61, 0x6c, 0x2e, 0x6f, 0x75, 0x74, 0x6f, 0x66, 0x67, 0x61, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4f, 0x75, 0x74, 0x4f, 0x66, 0x47, 0x61, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x40, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x6d,
	0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x70, 0x6f, 0x77, 0x65, 0x72, 0x72,
	0x65, 0x6e, 0x74, 0x61, 0x6c, 0x2e, 0x6f, 0x75, 0x74, 0x6f, 0x66, 0x67, 0x61, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4f, 0x75, 0x74, 0x4f, 0x66, 0x47, 0x61, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x32, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2c,
	0x3a, 0x01, 0x2a, 0x22, 0x27, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x6f, 0x77, 0x65, 0x72, 0x72, 0x65,
	0x6e, 0x74, 0x61, 0x6c, 0x5f, 0x6f, 0x75, 0x74, 0x6f, 0x66, 0x67, 0x61, 0x73, 0x5f, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x4f, 0x75, 0x74, 0x4f, 0x66, 0x47, 0x61, 0x73, 0x12, 0xc7, 0x01, 0x0a,
	0x0e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4f, 0x75, 0x74, 0x4f, 0x66, 0x47, 0x61, 0x73, 0x12,
	0x3f, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61,
	0x72, 0x65, 0x2e, 0x70, 0x6f, 0x77, 0x65, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x2e, 0x6f,
	0x75, 0x74, 0x6f, 0x66, 0x67, 0x61, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x4f, 0x75, 0x74, 0x4f, 0x66, 0x47, 0x61, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x40, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77,
	0x61, 0x72, 0x65, 0x2e, 0x70, 0x6f, 0x77, 0x65, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x2e,
	0x6f, 0x75, 0x74, 0x6f, 0x66, 0x67, 0x61, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x4f, 0x75, 0x74, 0x4f, 0x66, 0x47, 0x61, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x32, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2c, 0x3a, 0x01, 0x2a, 0x22, 0x27, 0x2f,
	0x76, 0x31, 0x2f, 0x70, 0x6f, 0x77, 0x65, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x5f, 0x6f,
	0x75, 0x74, 0x6f, 0x66, 0x67, 0x61, 0x73, 0x5f, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4f, 0x75,
	0x74, 0x4f, 0x66, 0x47, 0x61, 0x73, 0x42, 0x49, 0x5a, 0x47, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4e, 0x70, 0x6f, 0x6f, 0x6c, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f,
	0x72, 0x6d, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x6e, 0x70, 0x6f, 0x6f, 0x6c,
	0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2f, 0x6d, 0x77, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x6f, 0x77,
	0x65, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x2f, 0x6f, 0x75, 0x74, 0x6f, 0x66, 0x67, 0x61,
	0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_npool_order_mw_v1_powerrental_outofgas_outofgas_proto_rawDescOnce sync.Once
	file_npool_order_mw_v1_powerrental_outofgas_outofgas_proto_rawDescData = file_npool_order_mw_v1_powerrental_outofgas_outofgas_proto_rawDesc
)

func file_npool_order_mw_v1_powerrental_outofgas_outofgas_proto_rawDescGZIP() []byte {
	file_npool_order_mw_v1_powerrental_outofgas_outofgas_proto_rawDescOnce.Do(func() {
		file_npool_order_mw_v1_powerrental_outofgas_outofgas_proto_rawDescData = protoimpl.X.CompressGZIP(file_npool_order_mw_v1_powerrental_outofgas_outofgas_proto_rawDescData)
	})
	return file_npool_order_mw_v1_powerrental_outofgas_outofgas_proto_rawDescData
}

var file_npool_order_mw_v1_powerrental_outofgas_outofgas_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_npool_order_mw_v1_powerrental_outofgas_outofgas_proto_goTypes = []interface{}{
	(*OutOfGasReq)(nil),            // 0: order.middleware.powerrental.outofgas.v1.OutOfGasReq
	(*CreateOutOfGasRequest)(nil),  // 1: order.middleware.powerrental.outofgas.v1.CreateOutOfGasRequest
	(*CreateOutOfGasResponse)(nil), // 2: order.middleware.powerrental.outofgas.v1.CreateOutOfGasResponse
	(*UpdateOutOfGasRequest)(nil),  // 3: order.middleware.powerrental.outofgas.v1.UpdateOutOfGasRequest
	(*UpdateOutOfGasResponse)(nil), // 4: order.middleware.powerrental.outofgas.v1.UpdateOutOfGasResponse
	(*DeleteOutOfGasRequest)(nil),  // 5: order.middleware.powerrental.outofgas.v1.DeleteOutOfGasRequest
	(*DeleteOutOfGasResponse)(nil), // 6: order.middleware.powerrental.outofgas.v1.DeleteOutOfGasResponse
	(*outofgas.OutOfGas)(nil),      // 7: order.middleware.outofgas.v1.OutOfGas
}
var file_npool_order_mw_v1_powerrental_outofgas_outofgas_proto_depIdxs = []int32{
	0, // 0: order.middleware.powerrental.outofgas.v1.CreateOutOfGasRequest.Info:type_name -> order.middleware.powerrental.outofgas.v1.OutOfGasReq
	7, // 1: order.middleware.powerrental.outofgas.v1.CreateOutOfGasResponse.Info:type_name -> order.middleware.outofgas.v1.OutOfGas
	0, // 2: order.middleware.powerrental.outofgas.v1.UpdateOutOfGasRequest.Info:type_name -> order.middleware.powerrental.outofgas.v1.OutOfGasReq
	7, // 3: order.middleware.powerrental.outofgas.v1.UpdateOutOfGasResponse.Info:type_name -> order.middleware.outofgas.v1.OutOfGas
	0, // 4: order.middleware.powerrental.outofgas.v1.DeleteOutOfGasRequest.Info:type_name -> order.middleware.powerrental.outofgas.v1.OutOfGasReq
	7, // 5: order.middleware.powerrental.outofgas.v1.DeleteOutOfGasResponse.Info:type_name -> order.middleware.outofgas.v1.OutOfGas
	1, // 6: order.middleware.powerrental.outofgas.v1.Middleware.CreateOutOfGas:input_type -> order.middleware.powerrental.outofgas.v1.CreateOutOfGasRequest
	3, // 7: order.middleware.powerrental.outofgas.v1.Middleware.UpdateOutOfGas:input_type -> order.middleware.powerrental.outofgas.v1.UpdateOutOfGasRequest
	5, // 8: order.middleware.powerrental.outofgas.v1.Middleware.DeleteOutOfGas:input_type -> order.middleware.powerrental.outofgas.v1.DeleteOutOfGasRequest
	2, // 9: order.middleware.powerrental.outofgas.v1.Middleware.CreateOutOfGas:output_type -> order.middleware.powerrental.outofgas.v1.CreateOutOfGasResponse
	4, // 10: order.middleware.powerrental.outofgas.v1.Middleware.UpdateOutOfGas:output_type -> order.middleware.powerrental.outofgas.v1.UpdateOutOfGasResponse
	6, // 11: order.middleware.powerrental.outofgas.v1.Middleware.DeleteOutOfGas:output_type -> order.middleware.powerrental.outofgas.v1.DeleteOutOfGasResponse
	9, // [9:12] is the sub-list for method output_type
	6, // [6:9] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_npool_order_mw_v1_powerrental_outofgas_outofgas_proto_init() }
func file_npool_order_mw_v1_powerrental_outofgas_outofgas_proto_init() {
	if File_npool_order_mw_v1_powerrental_outofgas_outofgas_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_npool_order_mw_v1_powerrental_outofgas_outofgas_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OutOfGasReq); i {
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
		file_npool_order_mw_v1_powerrental_outofgas_outofgas_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateOutOfGasRequest); i {
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
		file_npool_order_mw_v1_powerrental_outofgas_outofgas_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateOutOfGasResponse); i {
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
		file_npool_order_mw_v1_powerrental_outofgas_outofgas_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateOutOfGasRequest); i {
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
		file_npool_order_mw_v1_powerrental_outofgas_outofgas_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateOutOfGasResponse); i {
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
		file_npool_order_mw_v1_powerrental_outofgas_outofgas_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteOutOfGasRequest); i {
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
		file_npool_order_mw_v1_powerrental_outofgas_outofgas_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteOutOfGasResponse); i {
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
	file_npool_order_mw_v1_powerrental_outofgas_outofgas_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_npool_order_mw_v1_powerrental_outofgas_outofgas_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_npool_order_mw_v1_powerrental_outofgas_outofgas_proto_goTypes,
		DependencyIndexes: file_npool_order_mw_v1_powerrental_outofgas_outofgas_proto_depIdxs,
		MessageInfos:      file_npool_order_mw_v1_powerrental_outofgas_outofgas_proto_msgTypes,
	}.Build()
	File_npool_order_mw_v1_powerrental_outofgas_outofgas_proto = out.File
	file_npool_order_mw_v1_powerrental_outofgas_outofgas_proto_rawDesc = nil
	file_npool_order_mw_v1_powerrental_outofgas_outofgas_proto_goTypes = nil
	file_npool_order_mw_v1_powerrental_outofgas_outofgas_proto_depIdxs = nil
}
