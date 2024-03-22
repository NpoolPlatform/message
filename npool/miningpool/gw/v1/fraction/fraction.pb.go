// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.18.1
// source: npool/miningpool/gw/v1/fraction/fraction.proto

package fraction

import (
	v1 "github.com/NpoolPlatform/message/npool/basetypes/miningpool/v1"
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

type Fraction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID               uint32           `protobuf:"varint,10,opt,name=ID,proto3" json:"ID,omitempty"`
	EntID            string           `protobuf:"bytes,20,opt,name=EntID,proto3" json:"EntID,omitempty"`
	OrderUserID      string           `protobuf:"bytes,30,opt,name=OrderUserID,proto3" json:"OrderUserID,omitempty"`
	WithdrawStateStr string           `protobuf:"bytes,40,opt,name=WithdrawStateStr,proto3" json:"WithdrawStateStr,omitempty"`
	WithdrawState    v1.WithdrawState `protobuf:"varint,41,opt,name=WithdrawState,proto3,enum=basetypes.miningpool.v1.WithdrawState" json:"WithdrawState,omitempty"`
	WithdrawTime     uint32           `protobuf:"varint,50,opt,name=WithdrawTime,proto3" json:"WithdrawTime,omitempty"`
	PayTime          uint32           `protobuf:"varint,60,opt,name=PayTime,proto3" json:"PayTime,omitempty"`
	CreatedAt        uint32           `protobuf:"varint,70,opt,name=CreatedAt,proto3" json:"CreatedAt,omitempty"`
	UpdatedAt        uint32           `protobuf:"varint,80,opt,name=UpdatedAt,proto3" json:"UpdatedAt,omitempty"`
}

func (x *Fraction) Reset() {
	*x = Fraction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_miningpool_gw_v1_fraction_fraction_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Fraction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Fraction) ProtoMessage() {}

func (x *Fraction) ProtoReflect() protoreflect.Message {
	mi := &file_npool_miningpool_gw_v1_fraction_fraction_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Fraction.ProtoReflect.Descriptor instead.
func (*Fraction) Descriptor() ([]byte, []int) {
	return file_npool_miningpool_gw_v1_fraction_fraction_proto_rawDescGZIP(), []int{0}
}

func (x *Fraction) GetID() uint32 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *Fraction) GetEntID() string {
	if x != nil {
		return x.EntID
	}
	return ""
}

func (x *Fraction) GetOrderUserID() string {
	if x != nil {
		return x.OrderUserID
	}
	return ""
}

func (x *Fraction) GetWithdrawStateStr() string {
	if x != nil {
		return x.WithdrawStateStr
	}
	return ""
}

func (x *Fraction) GetWithdrawState() v1.WithdrawState {
	if x != nil {
		return x.WithdrawState
	}
	return v1.WithdrawState(0)
}

func (x *Fraction) GetWithdrawTime() uint32 {
	if x != nil {
		return x.WithdrawTime
	}
	return 0
}

func (x *Fraction) GetPayTime() uint32 {
	if x != nil {
		return x.PayTime
	}
	return 0
}

func (x *Fraction) GetCreatedAt() uint32 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *Fraction) GetUpdatedAt() uint32 {
	if x != nil {
		return x.UpdatedAt
	}
	return 0
}

type CreateFractionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderUserID *string `protobuf:"bytes,30,opt,name=OrderUserID,proto3,oneof" json:"OrderUserID,omitempty"`
}

func (x *CreateFractionRequest) Reset() {
	*x = CreateFractionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_miningpool_gw_v1_fraction_fraction_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateFractionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateFractionRequest) ProtoMessage() {}

func (x *CreateFractionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_miningpool_gw_v1_fraction_fraction_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateFractionRequest.ProtoReflect.Descriptor instead.
func (*CreateFractionRequest) Descriptor() ([]byte, []int) {
	return file_npool_miningpool_gw_v1_fraction_fraction_proto_rawDescGZIP(), []int{1}
}

func (x *CreateFractionRequest) GetOrderUserID() string {
	if x != nil && x.OrderUserID != nil {
		return *x.OrderUserID
	}
	return ""
}

type CreateFractionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *Fraction `protobuf:"bytes,10,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (x *CreateFractionResponse) Reset() {
	*x = CreateFractionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_miningpool_gw_v1_fraction_fraction_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateFractionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateFractionResponse) ProtoMessage() {}

func (x *CreateFractionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_miningpool_gw_v1_fraction_fraction_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateFractionResponse.ProtoReflect.Descriptor instead.
func (*CreateFractionResponse) Descriptor() ([]byte, []int) {
	return file_npool_miningpool_gw_v1_fraction_fraction_proto_rawDescGZIP(), []int{2}
}

func (x *CreateFractionResponse) GetInfo() *Fraction {
	if x != nil {
		return x.Info
	}
	return nil
}

type GetFractionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EntID string `protobuf:"bytes,10,opt,name=EntID,proto3" json:"EntID,omitempty"`
}

func (x *GetFractionRequest) Reset() {
	*x = GetFractionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_miningpool_gw_v1_fraction_fraction_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFractionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFractionRequest) ProtoMessage() {}

func (x *GetFractionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_miningpool_gw_v1_fraction_fraction_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFractionRequest.ProtoReflect.Descriptor instead.
func (*GetFractionRequest) Descriptor() ([]byte, []int) {
	return file_npool_miningpool_gw_v1_fraction_fraction_proto_rawDescGZIP(), []int{3}
}

func (x *GetFractionRequest) GetEntID() string {
	if x != nil {
		return x.EntID
	}
	return ""
}

type GetFractionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *Fraction `protobuf:"bytes,10,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (x *GetFractionResponse) Reset() {
	*x = GetFractionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_miningpool_gw_v1_fraction_fraction_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFractionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFractionResponse) ProtoMessage() {}

func (x *GetFractionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_miningpool_gw_v1_fraction_fraction_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFractionResponse.ProtoReflect.Descriptor instead.
func (*GetFractionResponse) Descriptor() ([]byte, []int) {
	return file_npool_miningpool_gw_v1_fraction_fraction_proto_rawDescGZIP(), []int{4}
}

func (x *GetFractionResponse) GetInfo() *Fraction {
	if x != nil {
		return x.Info
	}
	return nil
}

type GetOrderFractionsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderUserID string `protobuf:"bytes,10,opt,name=OrderUserID,proto3" json:"OrderUserID,omitempty"`
	Offset      int32  `protobuf:"varint,20,opt,name=Offset,proto3" json:"Offset,omitempty"`
	Limit       int32  `protobuf:"varint,30,opt,name=Limit,proto3" json:"Limit,omitempty"`
}

func (x *GetOrderFractionsRequest) Reset() {
	*x = GetOrderFractionsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_miningpool_gw_v1_fraction_fraction_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetOrderFractionsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOrderFractionsRequest) ProtoMessage() {}

func (x *GetOrderFractionsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_miningpool_gw_v1_fraction_fraction_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOrderFractionsRequest.ProtoReflect.Descriptor instead.
func (*GetOrderFractionsRequest) Descriptor() ([]byte, []int) {
	return file_npool_miningpool_gw_v1_fraction_fraction_proto_rawDescGZIP(), []int{5}
}

func (x *GetOrderFractionsRequest) GetOrderUserID() string {
	if x != nil {
		return x.OrderUserID
	}
	return ""
}

func (x *GetOrderFractionsRequest) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *GetOrderFractionsRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type GetOrderFractionsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Infos []*Fraction `protobuf:"bytes,10,rep,name=Infos,proto3" json:"Infos,omitempty"`
	Total uint32      `protobuf:"varint,20,opt,name=Total,proto3" json:"Total,omitempty"`
}

func (x *GetOrderFractionsResponse) Reset() {
	*x = GetOrderFractionsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_miningpool_gw_v1_fraction_fraction_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetOrderFractionsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOrderFractionsResponse) ProtoMessage() {}

func (x *GetOrderFractionsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_miningpool_gw_v1_fraction_fraction_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOrderFractionsResponse.ProtoReflect.Descriptor instead.
func (*GetOrderFractionsResponse) Descriptor() ([]byte, []int) {
	return file_npool_miningpool_gw_v1_fraction_fraction_proto_rawDescGZIP(), []int{6}
}

func (x *GetOrderFractionsResponse) GetInfos() []*Fraction {
	if x != nil {
		return x.Infos
	}
	return nil
}

func (x *GetOrderFractionsResponse) GetTotal() uint32 {
	if x != nil {
		return x.Total
	}
	return 0
}

var File_npool_miningpool_gw_v1_fraction_fraction_proto protoreflect.FileDescriptor

var file_npool_miningpool_gw_v1_fraction_fraction_proto_rawDesc = []byte{
	0x0a, 0x2e, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x6d, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x70, 0x6f,
	0x6f, 0x6c, 0x2f, 0x67, 0x77, 0x2f, 0x76, 0x31, 0x2f, 0x66, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x2f, 0x66, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x1e, 0x6d, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x70, 0x6f, 0x6f, 0x6c, 0x2e, 0x67, 0x61, 0x74,
	0x65, 0x77, 0x61, 0x79, 0x2e, 0x66, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31,
	0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x29,
	0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f,
	0x6d, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x6e,
	0x75, 0x6d, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc6, 0x02, 0x0a, 0x08, 0x46, 0x72,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x02, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x45, 0x6e, 0x74, 0x49, 0x44, 0x18,
	0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x45, 0x6e, 0x74, 0x49, 0x44, 0x12, 0x20, 0x0a, 0x0b,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x1e, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x2a,
	0x0a, 0x10, 0x57, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x53, 0x74, 0x61, 0x74, 0x65, 0x53,
	0x74, 0x72, 0x18, 0x28, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x57, 0x69, 0x74, 0x68, 0x64, 0x72,
	0x61, 0x77, 0x53, 0x74, 0x61, 0x74, 0x65, 0x53, 0x74, 0x72, 0x12, 0x4c, 0x0a, 0x0d, 0x57, 0x69,
	0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x53, 0x74, 0x61, 0x74, 0x65, 0x18, 0x29, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x26, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x6d, 0x69,
	0x6e, 0x69, 0x6e, 0x67, 0x70, 0x6f, 0x6f, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x57, 0x69, 0x74, 0x68,
	0x64, 0x72, 0x61, 0x77, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x0d, 0x57, 0x69, 0x74, 0x68, 0x64,
	0x72, 0x61, 0x77, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x57, 0x69, 0x74, 0x68,
	0x64, 0x72, 0x61, 0x77, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x32, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c,
	0x57, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x50, 0x61, 0x79, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x3c, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x50,
	0x61, 0x79, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x18, 0x46, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x18, 0x50, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x22, 0x4e, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x46, 0x72, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x25, 0x0a, 0x0b, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x09,
	0x48, 0x00, 0x52, 0x0b, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x88,
	0x01, 0x01, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72,
	0x49, 0x44, 0x22, 0x56, 0x0a, 0x16, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x46, 0x72, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3c, 0x0a, 0x04,
	0x49, 0x6e, 0x66, 0x6f, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x6d, 0x69, 0x6e,
	0x69, 0x6e, 0x67, 0x70, 0x6f, 0x6f, 0x6c, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e,
	0x66, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x72, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x2a, 0x0a, 0x12, 0x47, 0x65,
	0x74, 0x46, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x14, 0x0a, 0x05, 0x45, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x45, 0x6e, 0x74, 0x49, 0x44, 0x22, 0x53, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x46, 0x72, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3c, 0x0a,
	0x04, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x6d, 0x69,
	0x6e, 0x69, 0x6e, 0x67, 0x70, 0x6f, 0x6f, 0x6c, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79,
	0x2e, 0x66, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x72, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x6a, 0x0a, 0x18, 0x47,
	0x65, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x46, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x4f, 0x66, 0x66,
	0x73, 0x65, 0x74, 0x18, 0x14, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x4f, 0x66, 0x66, 0x73, 0x65,
	0x74, 0x12, 0x14, 0x0a, 0x05, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x05, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x22, 0x71, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x46, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3e, 0x0a, 0x05, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x18, 0x0a, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x6d, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x70, 0x6f, 0x6f, 0x6c,
	0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x66, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x05, 0x49,
	0x6e, 0x66, 0x6f, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x14, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x32, 0xe8, 0x03, 0x0a, 0x0a, 0x4d,
	0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x12, 0x9f, 0x01, 0x0a, 0x0e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x46, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x35, 0x2e, 0x6d,
	0x69, 0x6e, 0x69, 0x6e, 0x67, 0x70, 0x6f, 0x6f, 0x6c, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61,
	0x79, 0x2e, 0x66, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x46, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x36, 0x2e, 0x6d, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x70, 0x6f, 0x6f, 0x6c,
	0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x66, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x46, 0x72, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1e, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x18, 0x3a, 0x01, 0x2a, 0x22, 0x13, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x2f, 0x66, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x93, 0x01, 0x0a, 0x0b,
	0x47, 0x65, 0x74, 0x46, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x32, 0x2e, 0x6d, 0x69,
	0x6e, 0x69, 0x6e, 0x67, 0x70, 0x6f, 0x6f, 0x6c, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79,
	0x2e, 0x66, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74,
	0x46, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x33, 0x2e, 0x6d, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x70, 0x6f, 0x6f, 0x6c, 0x2e, 0x67, 0x61, 0x74,
	0x65, 0x77, 0x61, 0x79, 0x2e, 0x66, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31,
	0x2e, 0x47, 0x65, 0x74, 0x46, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15, 0x3a, 0x01, 0x2a, 0x22,
	0x10, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x65, 0x74, 0x2f, 0x66, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0xa1, 0x01, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x46, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x12, 0x38, 0x2e, 0x6d, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x70, 0x6f, 0x6f, 0x6c, 0x2e,
	0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x66, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x46, 0x72, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x39, 0x2e, 0x6d,
	0x69, 0x6e, 0x69, 0x6e, 0x67, 0x70, 0x6f, 0x6f, 0x6c, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61,
	0x79, 0x2e, 0x66, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65,
	0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x46, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x16, 0x3a,
	0x01, 0x2a, 0x22, 0x11, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x65, 0x74, 0x2f, 0x66, 0x72, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x42, 0x42, 0x5a, 0x40, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x4e, 0x70, 0x6f, 0x6f, 0x6c, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72,
	0x6d, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f,
	0x6d, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x67, 0x77, 0x2f, 0x76, 0x31,
	0x2f, 0x66, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_npool_miningpool_gw_v1_fraction_fraction_proto_rawDescOnce sync.Once
	file_npool_miningpool_gw_v1_fraction_fraction_proto_rawDescData = file_npool_miningpool_gw_v1_fraction_fraction_proto_rawDesc
)

func file_npool_miningpool_gw_v1_fraction_fraction_proto_rawDescGZIP() []byte {
	file_npool_miningpool_gw_v1_fraction_fraction_proto_rawDescOnce.Do(func() {
		file_npool_miningpool_gw_v1_fraction_fraction_proto_rawDescData = protoimpl.X.CompressGZIP(file_npool_miningpool_gw_v1_fraction_fraction_proto_rawDescData)
	})
	return file_npool_miningpool_gw_v1_fraction_fraction_proto_rawDescData
}

var file_npool_miningpool_gw_v1_fraction_fraction_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_npool_miningpool_gw_v1_fraction_fraction_proto_goTypes = []interface{}{
	(*Fraction)(nil),                  // 0: miningpool.gateway.fraction.v1.Fraction
	(*CreateFractionRequest)(nil),     // 1: miningpool.gateway.fraction.v1.CreateFractionRequest
	(*CreateFractionResponse)(nil),    // 2: miningpool.gateway.fraction.v1.CreateFractionResponse
	(*GetFractionRequest)(nil),        // 3: miningpool.gateway.fraction.v1.GetFractionRequest
	(*GetFractionResponse)(nil),       // 4: miningpool.gateway.fraction.v1.GetFractionResponse
	(*GetOrderFractionsRequest)(nil),  // 5: miningpool.gateway.fraction.v1.GetOrderFractionsRequest
	(*GetOrderFractionsResponse)(nil), // 6: miningpool.gateway.fraction.v1.GetOrderFractionsResponse
	(v1.WithdrawState)(0),             // 7: basetypes.miningpool.v1.WithdrawState
}
var file_npool_miningpool_gw_v1_fraction_fraction_proto_depIdxs = []int32{
	7, // 0: miningpool.gateway.fraction.v1.Fraction.WithdrawState:type_name -> basetypes.miningpool.v1.WithdrawState
	0, // 1: miningpool.gateway.fraction.v1.CreateFractionResponse.Info:type_name -> miningpool.gateway.fraction.v1.Fraction
	0, // 2: miningpool.gateway.fraction.v1.GetFractionResponse.Info:type_name -> miningpool.gateway.fraction.v1.Fraction
	0, // 3: miningpool.gateway.fraction.v1.GetOrderFractionsResponse.Infos:type_name -> miningpool.gateway.fraction.v1.Fraction
	1, // 4: miningpool.gateway.fraction.v1.Middleware.CreateFraction:input_type -> miningpool.gateway.fraction.v1.CreateFractionRequest
	3, // 5: miningpool.gateway.fraction.v1.Middleware.GetFraction:input_type -> miningpool.gateway.fraction.v1.GetFractionRequest
	5, // 6: miningpool.gateway.fraction.v1.Middleware.GetFractions:input_type -> miningpool.gateway.fraction.v1.GetOrderFractionsRequest
	2, // 7: miningpool.gateway.fraction.v1.Middleware.CreateFraction:output_type -> miningpool.gateway.fraction.v1.CreateFractionResponse
	4, // 8: miningpool.gateway.fraction.v1.Middleware.GetFraction:output_type -> miningpool.gateway.fraction.v1.GetFractionResponse
	6, // 9: miningpool.gateway.fraction.v1.Middleware.GetFractions:output_type -> miningpool.gateway.fraction.v1.GetOrderFractionsResponse
	7, // [7:10] is the sub-list for method output_type
	4, // [4:7] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_npool_miningpool_gw_v1_fraction_fraction_proto_init() }
func file_npool_miningpool_gw_v1_fraction_fraction_proto_init() {
	if File_npool_miningpool_gw_v1_fraction_fraction_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_npool_miningpool_gw_v1_fraction_fraction_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Fraction); i {
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
		file_npool_miningpool_gw_v1_fraction_fraction_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateFractionRequest); i {
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
		file_npool_miningpool_gw_v1_fraction_fraction_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateFractionResponse); i {
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
		file_npool_miningpool_gw_v1_fraction_fraction_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFractionRequest); i {
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
		file_npool_miningpool_gw_v1_fraction_fraction_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFractionResponse); i {
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
		file_npool_miningpool_gw_v1_fraction_fraction_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetOrderFractionsRequest); i {
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
		file_npool_miningpool_gw_v1_fraction_fraction_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetOrderFractionsResponse); i {
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
	file_npool_miningpool_gw_v1_fraction_fraction_proto_msgTypes[1].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_npool_miningpool_gw_v1_fraction_fraction_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_npool_miningpool_gw_v1_fraction_fraction_proto_goTypes,
		DependencyIndexes: file_npool_miningpool_gw_v1_fraction_fraction_proto_depIdxs,
		MessageInfos:      file_npool_miningpool_gw_v1_fraction_fraction_proto_msgTypes,
	}.Build()
	File_npool_miningpool_gw_v1_fraction_fraction_proto = out.File
	file_npool_miningpool_gw_v1_fraction_fraction_proto_rawDesc = nil
	file_npool_miningpool_gw_v1_fraction_fraction_proto_goTypes = nil
	file_npool_miningpool_gw_v1_fraction_fraction_proto_depIdxs = nil
}
