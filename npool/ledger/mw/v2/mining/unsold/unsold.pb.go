// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.18.1
// source: npool/ledger/mw/v2/mining/unsold/unsold.proto

package unsold

import (
	unsold "github.com/NpoolPlatform/message/npool/ledger/mgr/v1/mining/unsold"
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

type UnsoldReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID                   *string `protobuf:"bytes,10,opt,name=ID,proto3,oneof" json:"ID,omitempty"`
	GoodID               *string `protobuf:"bytes,20,opt,name=GoodID,proto3,oneof" json:"GoodID,omitempty"`
	CoinTypeID           *string `protobuf:"bytes,30,opt,name=CoinTypeID,proto3,oneof" json:"CoinTypeID,omitempty"`
	Amount               *string `protobuf:"bytes,40,opt,name=Amount,proto3,oneof" json:"Amount,omitempty"`
	BenefitIntervalHours *uint32 `protobuf:"varint,50,opt,name=BenefitIntervalHours,proto3,oneof" json:"BenefitIntervalHours,omitempty"`
	CreatedAt            *uint32 `protobuf:"varint,60,opt,name=CreatedAt,proto3,oneof" json:"CreatedAt,omitempty"`
}

func (x *UnsoldReq) Reset() {
	*x = UnsoldReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_ledger_mw_v2_mining_unsold_unsold_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnsoldReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnsoldReq) ProtoMessage() {}

func (x *UnsoldReq) ProtoReflect() protoreflect.Message {
	mi := &file_npool_ledger_mw_v2_mining_unsold_unsold_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnsoldReq.ProtoReflect.Descriptor instead.
func (*UnsoldReq) Descriptor() ([]byte, []int) {
	return file_npool_ledger_mw_v2_mining_unsold_unsold_proto_rawDescGZIP(), []int{0}
}

func (x *UnsoldReq) GetID() string {
	if x != nil && x.ID != nil {
		return *x.ID
	}
	return ""
}

func (x *UnsoldReq) GetGoodID() string {
	if x != nil && x.GoodID != nil {
		return *x.GoodID
	}
	return ""
}

func (x *UnsoldReq) GetCoinTypeID() string {
	if x != nil && x.CoinTypeID != nil {
		return *x.CoinTypeID
	}
	return ""
}

func (x *UnsoldReq) GetAmount() string {
	if x != nil && x.Amount != nil {
		return *x.Amount
	}
	return ""
}

func (x *UnsoldReq) GetBenefitIntervalHours() uint32 {
	if x != nil && x.BenefitIntervalHours != nil {
		return *x.BenefitIntervalHours
	}
	return 0
}

func (x *UnsoldReq) GetCreatedAt() uint32 {
	if x != nil && x.CreatedAt != nil {
		return *x.CreatedAt
	}
	return 0
}

type CreateUnsoldRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *UnsoldReq `protobuf:"bytes,10,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (x *CreateUnsoldRequest) Reset() {
	*x = CreateUnsoldRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_ledger_mw_v2_mining_unsold_unsold_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateUnsoldRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUnsoldRequest) ProtoMessage() {}

func (x *CreateUnsoldRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_ledger_mw_v2_mining_unsold_unsold_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateUnsoldRequest.ProtoReflect.Descriptor instead.
func (*CreateUnsoldRequest) Descriptor() ([]byte, []int) {
	return file_npool_ledger_mw_v2_mining_unsold_unsold_proto_rawDescGZIP(), []int{1}
}

func (x *CreateUnsoldRequest) GetInfo() *UnsoldReq {
	if x != nil {
		return x.Info
	}
	return nil
}

type CreateUnsoldResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *unsold.Unsold `protobuf:"bytes,10,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (x *CreateUnsoldResponse) Reset() {
	*x = CreateUnsoldResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_ledger_mw_v2_mining_unsold_unsold_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateUnsoldResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUnsoldResponse) ProtoMessage() {}

func (x *CreateUnsoldResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_ledger_mw_v2_mining_unsold_unsold_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateUnsoldResponse.ProtoReflect.Descriptor instead.
func (*CreateUnsoldResponse) Descriptor() ([]byte, []int) {
	return file_npool_ledger_mw_v2_mining_unsold_unsold_proto_rawDescGZIP(), []int{2}
}

func (x *CreateUnsoldResponse) GetInfo() *unsold.Unsold {
	if x != nil {
		return x.Info
	}
	return nil
}

type GetUnsoldOnlyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Conds *unsold.Conds `protobuf:"bytes,10,opt,name=Conds,proto3" json:"Conds,omitempty"`
}

func (x *GetUnsoldOnlyRequest) Reset() {
	*x = GetUnsoldOnlyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_ledger_mw_v2_mining_unsold_unsold_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUnsoldOnlyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUnsoldOnlyRequest) ProtoMessage() {}

func (x *GetUnsoldOnlyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_ledger_mw_v2_mining_unsold_unsold_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUnsoldOnlyRequest.ProtoReflect.Descriptor instead.
func (*GetUnsoldOnlyRequest) Descriptor() ([]byte, []int) {
	return file_npool_ledger_mw_v2_mining_unsold_unsold_proto_rawDescGZIP(), []int{3}
}

func (x *GetUnsoldOnlyRequest) GetConds() *unsold.Conds {
	if x != nil {
		return x.Conds
	}
	return nil
}

type GetUnsoldOnlyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *unsold.Unsold `protobuf:"bytes,10,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (x *GetUnsoldOnlyResponse) Reset() {
	*x = GetUnsoldOnlyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_ledger_mw_v2_mining_unsold_unsold_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUnsoldOnlyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUnsoldOnlyResponse) ProtoMessage() {}

func (x *GetUnsoldOnlyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_ledger_mw_v2_mining_unsold_unsold_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUnsoldOnlyResponse.ProtoReflect.Descriptor instead.
func (*GetUnsoldOnlyResponse) Descriptor() ([]byte, []int) {
	return file_npool_ledger_mw_v2_mining_unsold_unsold_proto_rawDescGZIP(), []int{4}
}

func (x *GetUnsoldOnlyResponse) GetInfo() *unsold.Unsold {
	if x != nil {
		return x.Info
	}
	return nil
}

var File_npool_ledger_mw_v2_mining_unsold_unsold_proto protoreflect.FileDescriptor

var file_npool_ledger_mw_v2_mining_unsold_unsold_proto_rawDesc = []byte{
	0x0a, 0x2d, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2f, 0x6d,
	0x77, 0x2f, 0x76, 0x32, 0x2f, 0x6d, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x2f, 0x75, 0x6e, 0x73, 0x6f,
	0x6c, 0x64, 0x2f, 0x75, 0x6e, 0x73, 0x6f, 0x6c, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x22, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61,
	0x72, 0x65, 0x2e, 0x6d, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x75, 0x6e, 0x73, 0x6f, 0x6c, 0x64,
	0x2e, 0x76, 0x32, 0x1a, 0x2e, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x6c, 0x65, 0x64, 0x67, 0x65,
	0x72, 0x2f, 0x6d, 0x67, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x2f,
	0x75, 0x6e, 0x73, 0x6f, 0x6c, 0x64, 0x2f, 0x75, 0x6e, 0x73, 0x6f, 0x6c, 0x64, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xae, 0x02, 0x0a, 0x09, 0x55, 0x6e, 0x73, 0x6f, 0x6c, 0x64, 0x52, 0x65,
	0x71, 0x12, 0x13, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52,
	0x02, 0x49, 0x44, 0x88, 0x01, 0x01, 0x12, 0x1b, 0x0a, 0x06, 0x47, 0x6f, 0x6f, 0x64, 0x49, 0x44,
	0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x06, 0x47, 0x6f, 0x6f, 0x64, 0x49, 0x44,
	0x88, 0x01, 0x01, 0x12, 0x23, 0x0a, 0x0a, 0x43, 0x6f, 0x69, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x49,
	0x44, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x09, 0x48, 0x02, 0x52, 0x0a, 0x43, 0x6f, 0x69, 0x6e, 0x54,
	0x79, 0x70, 0x65, 0x49, 0x44, 0x88, 0x01, 0x01, 0x12, 0x1b, 0x0a, 0x06, 0x41, 0x6d, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x28, 0x20, 0x01, 0x28, 0x09, 0x48, 0x03, 0x52, 0x06, 0x41, 0x6d, 0x6f, 0x75,
	0x6e, 0x74, 0x88, 0x01, 0x01, 0x12, 0x37, 0x0a, 0x14, 0x42, 0x65, 0x6e, 0x65, 0x66, 0x69, 0x74,
	0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x48, 0x6f, 0x75, 0x72, 0x73, 0x18, 0x32, 0x20,
	0x01, 0x28, 0x0d, 0x48, 0x04, 0x52, 0x14, 0x42, 0x65, 0x6e, 0x65, 0x66, 0x69, 0x74, 0x49, 0x6e,
	0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x48, 0x6f, 0x75, 0x72, 0x73, 0x88, 0x01, 0x01, 0x12, 0x21,
	0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x3c, 0x20, 0x01, 0x28,
	0x0d, 0x48, 0x05, 0x52, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x88, 0x01,
	0x01, 0x42, 0x05, 0x0a, 0x03, 0x5f, 0x49, 0x44, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x47, 0x6f, 0x6f,
	0x64, 0x49, 0x44, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x43, 0x6f, 0x69, 0x6e, 0x54, 0x79, 0x70, 0x65,
	0x49, 0x44, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x17, 0x0a,
	0x15, 0x5f, 0x42, 0x65, 0x6e, 0x65, 0x66, 0x69, 0x74, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61,
	0x6c, 0x48, 0x6f, 0x75, 0x72, 0x73, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x22, 0x58, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x6e,
	0x73, 0x6f, 0x6c, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x41, 0x0a, 0x04, 0x49,
	0x6e, 0x66, 0x6f, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x6c, 0x65, 0x64, 0x67,
	0x65, 0x72, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x6d, 0x69,
	0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x75, 0x6e, 0x73, 0x6f, 0x6c, 0x64, 0x2e, 0x76, 0x32, 0x2e, 0x55,
	0x6e, 0x73, 0x6f, 0x6c, 0x64, 0x52, 0x65, 0x71, 0x52, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x53,
	0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x6e, 0x73, 0x6f, 0x6c, 0x64, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3b, 0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2e, 0x6d, 0x61,
	0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x6d, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x75, 0x6e, 0x73,
	0x6f, 0x6c, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x6e, 0x73, 0x6f, 0x6c, 0x64, 0x52, 0x04, 0x49,
	0x6e, 0x66, 0x6f, 0x22, 0x54, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x55, 0x6e, 0x73, 0x6f, 0x6c, 0x64,
	0x4f, 0x6e, 0x6c, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3c, 0x0a, 0x05, 0x43,
	0x6f, 0x6e, 0x64, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x6c, 0x65, 0x64,
	0x67, 0x65, 0x72, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x6d, 0x69, 0x6e, 0x69,
	0x6e, 0x67, 0x2e, 0x75, 0x6e, 0x73, 0x6f, 0x6c, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6e,
	0x64, 0x73, 0x52, 0x05, 0x43, 0x6f, 0x6e, 0x64, 0x73, 0x22, 0x54, 0x0a, 0x15, 0x47, 0x65, 0x74,
	0x55, 0x6e, 0x73, 0x6f, 0x6c, 0x64, 0x4f, 0x6e, 0x6c, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x3b, 0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x27, 0x2e, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65,
	0x72, 0x2e, 0x6d, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x75, 0x6e, 0x73, 0x6f, 0x6c, 0x64, 0x2e,
	0x76, 0x31, 0x2e, 0x55, 0x6e, 0x73, 0x6f, 0x6c, 0x64, 0x52, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x32,
	0x9b, 0x02, 0x0a, 0x0a, 0x4d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x12, 0x83,
	0x01, 0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x6e, 0x73, 0x6f, 0x6c, 0x64, 0x12,
	0x37, 0x2e, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77,
	0x61, 0x72, 0x65, 0x2e, 0x6d, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x75, 0x6e, 0x73, 0x6f, 0x6c,
	0x64, 0x2e, 0x76, 0x32, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x6e, 0x73, 0x6f, 0x6c,
	0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x38, 0x2e, 0x6c, 0x65, 0x64, 0x67, 0x65,
	0x72, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x6d, 0x69, 0x6e,
	0x69, 0x6e, 0x67, 0x2e, 0x75, 0x6e, 0x73, 0x6f, 0x6c, 0x64, 0x2e, 0x76, 0x32, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x55, 0x6e, 0x73, 0x6f, 0x6c, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x86, 0x01, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x55, 0x6e, 0x73, 0x6f,
	0x6c, 0x64, 0x4f, 0x6e, 0x6c, 0x79, 0x12, 0x38, 0x2e, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2e,
	0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x6d, 0x69, 0x6e, 0x69, 0x6e,
	0x67, 0x2e, 0x75, 0x6e, 0x73, 0x6f, 0x6c, 0x64, 0x2e, 0x76, 0x32, 0x2e, 0x47, 0x65, 0x74, 0x55,
	0x6e, 0x73, 0x6f, 0x6c, 0x64, 0x4f, 0x6e, 0x6c, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x39, 0x2e, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65,
	0x77, 0x61, 0x72, 0x65, 0x2e, 0x6d, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x75, 0x6e, 0x73, 0x6f,
	0x6c, 0x64, 0x2e, 0x76, 0x32, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x6e, 0x73, 0x6f, 0x6c, 0x64, 0x4f,
	0x6e, 0x6c, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x43, 0x5a,
	0x41, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4e, 0x70, 0x6f, 0x6f,
	0x6c, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x2f, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2f, 0x6d,
	0x77, 0x2f, 0x76, 0x32, 0x2f, 0x6d, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x2f, 0x75, 0x6e, 0x73, 0x6f,
	0x6c, 0x64, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_npool_ledger_mw_v2_mining_unsold_unsold_proto_rawDescOnce sync.Once
	file_npool_ledger_mw_v2_mining_unsold_unsold_proto_rawDescData = file_npool_ledger_mw_v2_mining_unsold_unsold_proto_rawDesc
)

func file_npool_ledger_mw_v2_mining_unsold_unsold_proto_rawDescGZIP() []byte {
	file_npool_ledger_mw_v2_mining_unsold_unsold_proto_rawDescOnce.Do(func() {
		file_npool_ledger_mw_v2_mining_unsold_unsold_proto_rawDescData = protoimpl.X.CompressGZIP(file_npool_ledger_mw_v2_mining_unsold_unsold_proto_rawDescData)
	})
	return file_npool_ledger_mw_v2_mining_unsold_unsold_proto_rawDescData
}

var file_npool_ledger_mw_v2_mining_unsold_unsold_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_npool_ledger_mw_v2_mining_unsold_unsold_proto_goTypes = []interface{}{
	(*UnsoldReq)(nil),             // 0: ledger.middleware.mining.unsold.v2.UnsoldReq
	(*CreateUnsoldRequest)(nil),   // 1: ledger.middleware.mining.unsold.v2.CreateUnsoldRequest
	(*CreateUnsoldResponse)(nil),  // 2: ledger.middleware.mining.unsold.v2.CreateUnsoldResponse
	(*GetUnsoldOnlyRequest)(nil),  // 3: ledger.middleware.mining.unsold.v2.GetUnsoldOnlyRequest
	(*GetUnsoldOnlyResponse)(nil), // 4: ledger.middleware.mining.unsold.v2.GetUnsoldOnlyResponse
	(*unsold.Unsold)(nil),         // 5: ledger.manager.mining.unsold.v1.Unsold
	(*unsold.Conds)(nil),          // 6: ledger.manager.mining.unsold.v1.Conds
}
var file_npool_ledger_mw_v2_mining_unsold_unsold_proto_depIdxs = []int32{
	0, // 0: ledger.middleware.mining.unsold.v2.CreateUnsoldRequest.Info:type_name -> ledger.middleware.mining.unsold.v2.UnsoldReq
	5, // 1: ledger.middleware.mining.unsold.v2.CreateUnsoldResponse.Info:type_name -> ledger.manager.mining.unsold.v1.Unsold
	6, // 2: ledger.middleware.mining.unsold.v2.GetUnsoldOnlyRequest.Conds:type_name -> ledger.manager.mining.unsold.v1.Conds
	5, // 3: ledger.middleware.mining.unsold.v2.GetUnsoldOnlyResponse.Info:type_name -> ledger.manager.mining.unsold.v1.Unsold
	1, // 4: ledger.middleware.mining.unsold.v2.Middleware.CreateUnsold:input_type -> ledger.middleware.mining.unsold.v2.CreateUnsoldRequest
	3, // 5: ledger.middleware.mining.unsold.v2.Middleware.GetUnsoldOnly:input_type -> ledger.middleware.mining.unsold.v2.GetUnsoldOnlyRequest
	2, // 6: ledger.middleware.mining.unsold.v2.Middleware.CreateUnsold:output_type -> ledger.middleware.mining.unsold.v2.CreateUnsoldResponse
	4, // 7: ledger.middleware.mining.unsold.v2.Middleware.GetUnsoldOnly:output_type -> ledger.middleware.mining.unsold.v2.GetUnsoldOnlyResponse
	6, // [6:8] is the sub-list for method output_type
	4, // [4:6] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_npool_ledger_mw_v2_mining_unsold_unsold_proto_init() }
func file_npool_ledger_mw_v2_mining_unsold_unsold_proto_init() {
	if File_npool_ledger_mw_v2_mining_unsold_unsold_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_npool_ledger_mw_v2_mining_unsold_unsold_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UnsoldReq); i {
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
		file_npool_ledger_mw_v2_mining_unsold_unsold_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateUnsoldRequest); i {
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
		file_npool_ledger_mw_v2_mining_unsold_unsold_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateUnsoldResponse); i {
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
		file_npool_ledger_mw_v2_mining_unsold_unsold_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUnsoldOnlyRequest); i {
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
		file_npool_ledger_mw_v2_mining_unsold_unsold_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUnsoldOnlyResponse); i {
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
	file_npool_ledger_mw_v2_mining_unsold_unsold_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_npool_ledger_mw_v2_mining_unsold_unsold_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_npool_ledger_mw_v2_mining_unsold_unsold_proto_goTypes,
		DependencyIndexes: file_npool_ledger_mw_v2_mining_unsold_unsold_proto_depIdxs,
		MessageInfos:      file_npool_ledger_mw_v2_mining_unsold_unsold_proto_msgTypes,
	}.Build()
	File_npool_ledger_mw_v2_mining_unsold_unsold_proto = out.File
	file_npool_ledger_mw_v2_mining_unsold_unsold_proto_rawDesc = nil
	file_npool_ledger_mw_v2_mining_unsold_unsold_proto_goTypes = nil
	file_npool_ledger_mw_v2_mining_unsold_unsold_proto_depIdxs = nil
}
