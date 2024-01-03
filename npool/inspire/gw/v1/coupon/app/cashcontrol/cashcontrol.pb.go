// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.18.1
// source: npool/inspire/gw/v1/coupon/app/cashcontrol/cashcontrol.proto

package cashcontrol

import (
	v1 "github.com/NpoolPlatform/message/npool/basetypes/inspire/v1"
	cashcontrol "github.com/NpoolPlatform/message/npool/inspire/mw/v1/coupon/cashcontrol"
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

type GetCashControlsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TargetAppID string `protobuf:"bytes,10,opt,name=TargetAppID,proto3" json:"TargetAppID,omitempty"`
	Offset      int32  `protobuf:"varint,20,opt,name=Offset,proto3" json:"Offset,omitempty"`
	Limit       int32  `protobuf:"varint,30,opt,name=Limit,proto3" json:"Limit,omitempty"`
}

func (x *GetCashControlsRequest) Reset() {
	*x = GetCashControlsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_inspire_gw_v1_coupon_app_cashcontrol_cashcontrol_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCashControlsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCashControlsRequest) ProtoMessage() {}

func (x *GetCashControlsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_inspire_gw_v1_coupon_app_cashcontrol_cashcontrol_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCashControlsRequest.ProtoReflect.Descriptor instead.
func (*GetCashControlsRequest) Descriptor() ([]byte, []int) {
	return file_npool_inspire_gw_v1_coupon_app_cashcontrol_cashcontrol_proto_rawDescGZIP(), []int{0}
}

func (x *GetCashControlsRequest) GetTargetAppID() string {
	if x != nil {
		return x.TargetAppID
	}
	return ""
}

func (x *GetCashControlsRequest) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *GetCashControlsRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type GetCashControlsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Infos []*cashcontrol.CashControl `protobuf:"bytes,10,rep,name=Infos,proto3" json:"Infos,omitempty"`
	Total uint32                     `protobuf:"varint,20,opt,name=Total,proto3" json:"Total,omitempty"`
}

func (x *GetCashControlsResponse) Reset() {
	*x = GetCashControlsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_inspire_gw_v1_coupon_app_cashcontrol_cashcontrol_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCashControlsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCashControlsResponse) ProtoMessage() {}

func (x *GetCashControlsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_inspire_gw_v1_coupon_app_cashcontrol_cashcontrol_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCashControlsResponse.ProtoReflect.Descriptor instead.
func (*GetCashControlsResponse) Descriptor() ([]byte, []int) {
	return file_npool_inspire_gw_v1_coupon_app_cashcontrol_cashcontrol_proto_rawDescGZIP(), []int{1}
}

func (x *GetCashControlsResponse) GetInfos() []*cashcontrol.CashControl {
	if x != nil {
		return x.Infos
	}
	return nil
}

func (x *GetCashControlsResponse) GetTotal() uint32 {
	if x != nil {
		return x.Total
	}
	return 0
}

type CreateCashControlRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CouponID        string             `protobuf:"bytes,10,opt,name=CouponID,proto3" json:"CouponID,omitempty"`
	CashControlType v1.CashControlType `protobuf:"varint,20,opt,name=CashControlType,proto3,enum=basetypes.inspire.v1.CashControlType" json:"CashControlType,omitempty"`
	Amount          *string            `protobuf:"bytes,30,opt,name=Amount,proto3,oneof" json:"Amount,omitempty"`
}

func (x *CreateCashControlRequest) Reset() {
	*x = CreateCashControlRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_inspire_gw_v1_coupon_app_cashcontrol_cashcontrol_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCashControlRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCashControlRequest) ProtoMessage() {}

func (x *CreateCashControlRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_inspire_gw_v1_coupon_app_cashcontrol_cashcontrol_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCashControlRequest.ProtoReflect.Descriptor instead.
func (*CreateCashControlRequest) Descriptor() ([]byte, []int) {
	return file_npool_inspire_gw_v1_coupon_app_cashcontrol_cashcontrol_proto_rawDescGZIP(), []int{2}
}

func (x *CreateCashControlRequest) GetCouponID() string {
	if x != nil {
		return x.CouponID
	}
	return ""
}

func (x *CreateCashControlRequest) GetCashControlType() v1.CashControlType {
	if x != nil {
		return x.CashControlType
	}
	return v1.CashControlType(0)
}

func (x *CreateCashControlRequest) GetAmount() string {
	if x != nil && x.Amount != nil {
		return *x.Amount
	}
	return ""
}

type CreateCashControlResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *cashcontrol.CashControl `protobuf:"bytes,10,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (x *CreateCashControlResponse) Reset() {
	*x = CreateCashControlResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_inspire_gw_v1_coupon_app_cashcontrol_cashcontrol_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCashControlResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCashControlResponse) ProtoMessage() {}

func (x *CreateCashControlResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_inspire_gw_v1_coupon_app_cashcontrol_cashcontrol_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCashControlResponse.ProtoReflect.Descriptor instead.
func (*CreateCashControlResponse) Descriptor() ([]byte, []int) {
	return file_npool_inspire_gw_v1_coupon_app_cashcontrol_cashcontrol_proto_rawDescGZIP(), []int{3}
}

func (x *CreateCashControlResponse) GetInfo() *cashcontrol.CashControl {
	if x != nil {
		return x.Info
	}
	return nil
}

type DeleteCashControlRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID          uint32 `protobuf:"varint,10,opt,name=ID,proto3" json:"ID,omitempty"`
	EntID       string `protobuf:"bytes,20,opt,name=EntID,proto3" json:"EntID,omitempty"`
	TargetAppID string `protobuf:"bytes,30,opt,name=TargetAppID,proto3" json:"TargetAppID,omitempty"`
}

func (x *DeleteCashControlRequest) Reset() {
	*x = DeleteCashControlRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_inspire_gw_v1_coupon_app_cashcontrol_cashcontrol_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteCashControlRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteCashControlRequest) ProtoMessage() {}

func (x *DeleteCashControlRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_inspire_gw_v1_coupon_app_cashcontrol_cashcontrol_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteCashControlRequest.ProtoReflect.Descriptor instead.
func (*DeleteCashControlRequest) Descriptor() ([]byte, []int) {
	return file_npool_inspire_gw_v1_coupon_app_cashcontrol_cashcontrol_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteCashControlRequest) GetID() uint32 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *DeleteCashControlRequest) GetEntID() string {
	if x != nil {
		return x.EntID
	}
	return ""
}

func (x *DeleteCashControlRequest) GetTargetAppID() string {
	if x != nil {
		return x.TargetAppID
	}
	return ""
}

type DeleteCashControlResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *cashcontrol.CashControl `protobuf:"bytes,10,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (x *DeleteCashControlResponse) Reset() {
	*x = DeleteCashControlResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_inspire_gw_v1_coupon_app_cashcontrol_cashcontrol_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteCashControlResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteCashControlResponse) ProtoMessage() {}

func (x *DeleteCashControlResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_inspire_gw_v1_coupon_app_cashcontrol_cashcontrol_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteCashControlResponse.ProtoReflect.Descriptor instead.
func (*DeleteCashControlResponse) Descriptor() ([]byte, []int) {
	return file_npool_inspire_gw_v1_coupon_app_cashcontrol_cashcontrol_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteCashControlResponse) GetInfo() *cashcontrol.CashControl {
	if x != nil {
		return x.Info
	}
	return nil
}

type GetAppCashControlsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AppID  string `protobuf:"bytes,10,opt,name=AppID,proto3" json:"AppID,omitempty"`
	Offset int32  `protobuf:"varint,20,opt,name=Offset,proto3" json:"Offset,omitempty"`
	Limit  int32  `protobuf:"varint,30,opt,name=Limit,proto3" json:"Limit,omitempty"`
}

func (x *GetAppCashControlsRequest) Reset() {
	*x = GetAppCashControlsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_inspire_gw_v1_coupon_app_cashcontrol_cashcontrol_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAppCashControlsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAppCashControlsRequest) ProtoMessage() {}

func (x *GetAppCashControlsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_inspire_gw_v1_coupon_app_cashcontrol_cashcontrol_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAppCashControlsRequest.ProtoReflect.Descriptor instead.
func (*GetAppCashControlsRequest) Descriptor() ([]byte, []int) {
	return file_npool_inspire_gw_v1_coupon_app_cashcontrol_cashcontrol_proto_rawDescGZIP(), []int{6}
}

func (x *GetAppCashControlsRequest) GetAppID() string {
	if x != nil {
		return x.AppID
	}
	return ""
}

func (x *GetAppCashControlsRequest) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *GetAppCashControlsRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type GetAppCashControlsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Infos []*cashcontrol.CashControl `protobuf:"bytes,10,rep,name=Infos,proto3" json:"Infos,omitempty"`
	Total uint32                     `protobuf:"varint,20,opt,name=Total,proto3" json:"Total,omitempty"`
}

func (x *GetAppCashControlsResponse) Reset() {
	*x = GetAppCashControlsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_inspire_gw_v1_coupon_app_cashcontrol_cashcontrol_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAppCashControlsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAppCashControlsResponse) ProtoMessage() {}

func (x *GetAppCashControlsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_inspire_gw_v1_coupon_app_cashcontrol_cashcontrol_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAppCashControlsResponse.ProtoReflect.Descriptor instead.
func (*GetAppCashControlsResponse) Descriptor() ([]byte, []int) {
	return file_npool_inspire_gw_v1_coupon_app_cashcontrol_cashcontrol_proto_rawDescGZIP(), []int{7}
}

func (x *GetAppCashControlsResponse) GetInfos() []*cashcontrol.CashControl {
	if x != nil {
		return x.Infos
	}
	return nil
}

func (x *GetAppCashControlsResponse) GetTotal() uint32 {
	if x != nil {
		return x.Total
	}
	return 0
}

var File_npool_inspire_gw_v1_coupon_app_cashcontrol_cashcontrol_proto protoreflect.FileDescriptor

var file_npool_inspire_gw_v1_coupon_app_cashcontrol_cashcontrol_proto_rawDesc = []byte{
	0x0a, 0x3c, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x69, 0x6e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2f,
	0x67, 0x77, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x70,
	0x2f, 0x63, 0x61, 0x73, 0x68, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2f, 0x63, 0x61, 0x73,
	0x68, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x29,
	0x69, 0x6e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e,
	0x63, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x63, 0x61, 0x73, 0x68, 0x63,
	0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x26, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x62,
	0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x69, 0x6e, 0x73, 0x70, 0x69, 0x72, 0x65,
	0x2f, 0x76, 0x31, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x3c, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x69, 0x6e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2f, 0x6d,
	0x77, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x70, 0x2f,
	0x63, 0x61, 0x73, 0x68, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2f, 0x63, 0x61, 0x73, 0x68,
	0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x68, 0x0a,
	0x16, 0x47, 0x65, 0x74, 0x43, 0x61, 0x73, 0x68, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x54, 0x61, 0x72, 0x67, 0x65,
	0x74, 0x41, 0x70, 0x70, 0x49, 0x44, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x54, 0x61,
	0x72, 0x67, 0x65, 0x74, 0x41, 0x70, 0x70, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x4f, 0x66, 0x66,
	0x73, 0x65, 0x74, 0x18, 0x14, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x4f, 0x66, 0x66, 0x73, 0x65,
	0x74, 0x12, 0x14, 0x0a, 0x05, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x05, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x22, 0x7c, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x43, 0x61,
	0x73, 0x68, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x4b, 0x0a, 0x05, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x35, 0x2e, 0x69, 0x6e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x6d, 0x69, 0x64, 0x64,
	0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x63, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x2e, 0x63, 0x61,
	0x73, 0x68, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x61, 0x73,
	0x68, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x52, 0x05, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x12,
	0x14, 0x0a, 0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x14, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05,
	0x54, 0x6f, 0x74, 0x61, 0x6c, 0x22, 0xaf, 0x01, 0x0a, 0x18, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x43, 0x61, 0x73, 0x68, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x43, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x49, 0x44, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x43, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x49, 0x44, 0x12, 0x4f,
	0x0a, 0x0f, 0x43, 0x61, 0x73, 0x68, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x54, 0x79, 0x70,
	0x65, 0x18, 0x14, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x25, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x2e, 0x69, 0x6e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43,
	0x61, 0x73, 0x68, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0f,
	0x43, 0x61, 0x73, 0x68, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x1b, 0x0a, 0x06, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x09, 0x48,
	0x00, 0x52, 0x06, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x88, 0x01, 0x01, 0x42, 0x09, 0x0a, 0x07,
	0x5f, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x66, 0x0a, 0x19, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x43, 0x61, 0x73, 0x68, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x49, 0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x35, 0x2e, 0x69, 0x6e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x6d, 0x69, 0x64,
	0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x63, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x2e, 0x63,
	0x61, 0x73, 0x68, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x61,
	0x73, 0x68, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x52, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x22,
	0x62, 0x0a, 0x18, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x61, 0x73, 0x68, 0x43, 0x6f, 0x6e,
	0x74, 0x72, 0x6f, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49,
	0x44, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x45,
	0x6e, 0x74, 0x49, 0x44, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x45, 0x6e, 0x74, 0x49,
	0x44, 0x12, 0x20, 0x0a, 0x0b, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x41, 0x70, 0x70, 0x49, 0x44,
	0x18, 0x1e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x41, 0x70,
	0x70, 0x49, 0x44, 0x22, 0x66, 0x0a, 0x19, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x61, 0x73,
	0x68, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x49, 0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x35,
	0x2e, 0x69, 0x6e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77,
	0x61, 0x72, 0x65, 0x2e, 0x63, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x2e, 0x63, 0x61, 0x73, 0x68, 0x63,
	0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x61, 0x73, 0x68, 0x43, 0x6f,
	0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x52, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x5f, 0x0a, 0x19, 0x47,
	0x65, 0x74, 0x41, 0x70, 0x70, 0x43, 0x61, 0x73, 0x68, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x41, 0x70, 0x70, 0x49,
	0x44, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x41, 0x70, 0x70, 0x49, 0x44, 0x12, 0x16,
	0x0a, 0x06, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x14, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06,
	0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x18,
	0x1e, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x22, 0x7f, 0x0a, 0x1a,
	0x47, 0x65, 0x74, 0x41, 0x70, 0x70, 0x43, 0x61, 0x73, 0x68, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f,
	0x6c, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4b, 0x0a, 0x05, 0x49, 0x6e,
	0x66, 0x6f, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x35, 0x2e, 0x69, 0x6e, 0x73, 0x70,
	0x69, 0x72, 0x65, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x63,
	0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x2e, 0x63, 0x61, 0x73, 0x68, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f,
	0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x61, 0x73, 0x68, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c,
	0x52, 0x05, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c,
	0x18, 0x14, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x32, 0x95, 0x06,
	0x0a, 0x07, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x12, 0xb9, 0x01, 0x0a, 0x0f, 0x47, 0x65,
	0x74, 0x43, 0x61, 0x73, 0x68, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x73, 0x12, 0x41, 0x2e,
	0x69, 0x6e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e,
	0x63, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x63, 0x61, 0x73, 0x68, 0x63,
	0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x61, 0x73,
	0x68, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x42, 0x2e, 0x69, 0x6e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77,
	0x61, 0x79, 0x2e, 0x63, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x63, 0x61,
	0x73, 0x68, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74,
	0x43, 0x61, 0x73, 0x68, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x19, 0x3a, 0x01, 0x2a, 0x22,
	0x14, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x65, 0x74, 0x2f, 0x63, 0x61, 0x73, 0x68, 0x63, 0x6f, 0x6e,
	0x74, 0x72, 0x6f, 0x6c, 0x73, 0x12, 0xc1, 0x01, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x43, 0x61, 0x73, 0x68, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x12, 0x43, 0x2e, 0x69, 0x6e,
	0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x63, 0x6f,
	0x75, 0x70, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x63, 0x61, 0x73, 0x68, 0x63, 0x6f, 0x6e,
	0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x61,
	0x73, 0x68, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x44, 0x2e, 0x69, 0x6e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77,
	0x61, 0x79, 0x2e, 0x63, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x63, 0x61,
	0x73, 0x68, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x43, 0x61, 0x73, 0x68, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x21, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1b, 0x3a, 0x01,
	0x2a, 0x22, 0x16, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x2f, 0x63, 0x61,
	0x73, 0x68, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x12, 0xc1, 0x01, 0x0a, 0x11, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x43, 0x61, 0x73, 0x68, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x12,
	0x43, 0x2e, 0x69, 0x6e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61,
	0x79, 0x2e, 0x63, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x63, 0x61, 0x73,
	0x68, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x43, 0x61, 0x73, 0x68, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x44, 0x2e, 0x69, 0x6e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x67,
	0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x63, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x2e, 0x61, 0x70,
	0x70, 0x2e, 0x63, 0x61, 0x73, 0x68, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x76, 0x31,
	0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x61, 0x73, 0x68, 0x43, 0x6f, 0x6e, 0x74, 0x72,
	0x6f, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x21, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x1b, 0x3a, 0x01, 0x2a, 0x22, 0x16, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x2f, 0x63, 0x61, 0x73, 0x68, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x12, 0xc5, 0x01,
	0x0a, 0x12, 0x47, 0x65, 0x74, 0x41, 0x70, 0x70, 0x43, 0x61, 0x73, 0x68, 0x43, 0x6f, 0x6e, 0x74,
	0x72, 0x6f, 0x6c, 0x73, 0x12, 0x44, 0x2e, 0x69, 0x6e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x67,
	0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x63, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x2e, 0x61, 0x70,
	0x70, 0x2e, 0x63, 0x61, 0x73, 0x68, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x76, 0x31,
	0x2e, 0x47, 0x65, 0x74, 0x41, 0x70, 0x70, 0x43, 0x61, 0x73, 0x68, 0x43, 0x6f, 0x6e, 0x74, 0x72,
	0x6f, 0x6c, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x45, 0x2e, 0x69, 0x6e, 0x73,
	0x70, 0x69, 0x72, 0x65, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x63, 0x6f, 0x75,
	0x70, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x63, 0x61, 0x73, 0x68, 0x63, 0x6f, 0x6e, 0x74,
	0x72, 0x6f, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x70, 0x70, 0x43, 0x61, 0x73,
	0x68, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x22, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1c, 0x3a, 0x01, 0x2a, 0x22, 0x17, 0x2f, 0x76,
	0x31, 0x2f, 0x67, 0x65, 0x74, 0x2f, 0x61, 0x70, 0x70, 0x63, 0x61, 0x73, 0x68, 0x63, 0x6f, 0x6e,
	0x74, 0x72, 0x6f, 0x6c, 0x73, 0x42, 0x4d, 0x5a, 0x4b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x4e, 0x70, 0x6f, 0x6f, 0x6c, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72,
	0x6d, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f,
	0x69, 0x6e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2f, 0x67, 0x77, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f,
	0x75, 0x70, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x70, 0x2f, 0x63, 0x61, 0x73, 0x68, 0x63, 0x6f, 0x6e,
	0x74, 0x72, 0x6f, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_npool_inspire_gw_v1_coupon_app_cashcontrol_cashcontrol_proto_rawDescOnce sync.Once
	file_npool_inspire_gw_v1_coupon_app_cashcontrol_cashcontrol_proto_rawDescData = file_npool_inspire_gw_v1_coupon_app_cashcontrol_cashcontrol_proto_rawDesc
)

func file_npool_inspire_gw_v1_coupon_app_cashcontrol_cashcontrol_proto_rawDescGZIP() []byte {
	file_npool_inspire_gw_v1_coupon_app_cashcontrol_cashcontrol_proto_rawDescOnce.Do(func() {
		file_npool_inspire_gw_v1_coupon_app_cashcontrol_cashcontrol_proto_rawDescData = protoimpl.X.CompressGZIP(file_npool_inspire_gw_v1_coupon_app_cashcontrol_cashcontrol_proto_rawDescData)
	})
	return file_npool_inspire_gw_v1_coupon_app_cashcontrol_cashcontrol_proto_rawDescData
}

var file_npool_inspire_gw_v1_coupon_app_cashcontrol_cashcontrol_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_npool_inspire_gw_v1_coupon_app_cashcontrol_cashcontrol_proto_goTypes = []interface{}{
	(*GetCashControlsRequest)(nil),     // 0: inspire.gateway.coupon.app.cashcontrol.v1.GetCashControlsRequest
	(*GetCashControlsResponse)(nil),    // 1: inspire.gateway.coupon.app.cashcontrol.v1.GetCashControlsResponse
	(*CreateCashControlRequest)(nil),   // 2: inspire.gateway.coupon.app.cashcontrol.v1.CreateCashControlRequest
	(*CreateCashControlResponse)(nil),  // 3: inspire.gateway.coupon.app.cashcontrol.v1.CreateCashControlResponse
	(*DeleteCashControlRequest)(nil),   // 4: inspire.gateway.coupon.app.cashcontrol.v1.DeleteCashControlRequest
	(*DeleteCashControlResponse)(nil),  // 5: inspire.gateway.coupon.app.cashcontrol.v1.DeleteCashControlResponse
	(*GetAppCashControlsRequest)(nil),  // 6: inspire.gateway.coupon.app.cashcontrol.v1.GetAppCashControlsRequest
	(*GetAppCashControlsResponse)(nil), // 7: inspire.gateway.coupon.app.cashcontrol.v1.GetAppCashControlsResponse
	(*cashcontrol.CashControl)(nil),    // 8: inspire.middleware.coupon.cashcontrol.v1.CashControl
	(v1.CashControlType)(0),            // 9: basetypes.inspire.v1.CashControlType
}
var file_npool_inspire_gw_v1_coupon_app_cashcontrol_cashcontrol_proto_depIdxs = []int32{
	8, // 0: inspire.gateway.coupon.app.cashcontrol.v1.GetCashControlsResponse.Infos:type_name -> inspire.middleware.coupon.cashcontrol.v1.CashControl
	9, // 1: inspire.gateway.coupon.app.cashcontrol.v1.CreateCashControlRequest.CashControlType:type_name -> basetypes.inspire.v1.CashControlType
	8, // 2: inspire.gateway.coupon.app.cashcontrol.v1.CreateCashControlResponse.Info:type_name -> inspire.middleware.coupon.cashcontrol.v1.CashControl
	8, // 3: inspire.gateway.coupon.app.cashcontrol.v1.DeleteCashControlResponse.Info:type_name -> inspire.middleware.coupon.cashcontrol.v1.CashControl
	8, // 4: inspire.gateway.coupon.app.cashcontrol.v1.GetAppCashControlsResponse.Infos:type_name -> inspire.middleware.coupon.cashcontrol.v1.CashControl
	0, // 5: inspire.gateway.coupon.app.cashcontrol.v1.Gateway.GetCashControls:input_type -> inspire.gateway.coupon.app.cashcontrol.v1.GetCashControlsRequest
	2, // 6: inspire.gateway.coupon.app.cashcontrol.v1.Gateway.CreateCashControl:input_type -> inspire.gateway.coupon.app.cashcontrol.v1.CreateCashControlRequest
	4, // 7: inspire.gateway.coupon.app.cashcontrol.v1.Gateway.DeleteCashControl:input_type -> inspire.gateway.coupon.app.cashcontrol.v1.DeleteCashControlRequest
	6, // 8: inspire.gateway.coupon.app.cashcontrol.v1.Gateway.GetAppCashControls:input_type -> inspire.gateway.coupon.app.cashcontrol.v1.GetAppCashControlsRequest
	1, // 9: inspire.gateway.coupon.app.cashcontrol.v1.Gateway.GetCashControls:output_type -> inspire.gateway.coupon.app.cashcontrol.v1.GetCashControlsResponse
	3, // 10: inspire.gateway.coupon.app.cashcontrol.v1.Gateway.CreateCashControl:output_type -> inspire.gateway.coupon.app.cashcontrol.v1.CreateCashControlResponse
	5, // 11: inspire.gateway.coupon.app.cashcontrol.v1.Gateway.DeleteCashControl:output_type -> inspire.gateway.coupon.app.cashcontrol.v1.DeleteCashControlResponse
	7, // 12: inspire.gateway.coupon.app.cashcontrol.v1.Gateway.GetAppCashControls:output_type -> inspire.gateway.coupon.app.cashcontrol.v1.GetAppCashControlsResponse
	9, // [9:13] is the sub-list for method output_type
	5, // [5:9] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_npool_inspire_gw_v1_coupon_app_cashcontrol_cashcontrol_proto_init() }
func file_npool_inspire_gw_v1_coupon_app_cashcontrol_cashcontrol_proto_init() {
	if File_npool_inspire_gw_v1_coupon_app_cashcontrol_cashcontrol_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_npool_inspire_gw_v1_coupon_app_cashcontrol_cashcontrol_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCashControlsRequest); i {
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
		file_npool_inspire_gw_v1_coupon_app_cashcontrol_cashcontrol_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCashControlsResponse); i {
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
		file_npool_inspire_gw_v1_coupon_app_cashcontrol_cashcontrol_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateCashControlRequest); i {
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
		file_npool_inspire_gw_v1_coupon_app_cashcontrol_cashcontrol_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateCashControlResponse); i {
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
		file_npool_inspire_gw_v1_coupon_app_cashcontrol_cashcontrol_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteCashControlRequest); i {
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
		file_npool_inspire_gw_v1_coupon_app_cashcontrol_cashcontrol_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteCashControlResponse); i {
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
		file_npool_inspire_gw_v1_coupon_app_cashcontrol_cashcontrol_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAppCashControlsRequest); i {
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
		file_npool_inspire_gw_v1_coupon_app_cashcontrol_cashcontrol_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAppCashControlsResponse); i {
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
	file_npool_inspire_gw_v1_coupon_app_cashcontrol_cashcontrol_proto_msgTypes[2].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_npool_inspire_gw_v1_coupon_app_cashcontrol_cashcontrol_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_npool_inspire_gw_v1_coupon_app_cashcontrol_cashcontrol_proto_goTypes,
		DependencyIndexes: file_npool_inspire_gw_v1_coupon_app_cashcontrol_cashcontrol_proto_depIdxs,
		MessageInfos:      file_npool_inspire_gw_v1_coupon_app_cashcontrol_cashcontrol_proto_msgTypes,
	}.Build()
	File_npool_inspire_gw_v1_coupon_app_cashcontrol_cashcontrol_proto = out.File
	file_npool_inspire_gw_v1_coupon_app_cashcontrol_cashcontrol_proto_rawDesc = nil
	file_npool_inspire_gw_v1_coupon_app_cashcontrol_cashcontrol_proto_goTypes = nil
	file_npool_inspire_gw_v1_coupon_app_cashcontrol_cashcontrol_proto_depIdxs = nil
}
