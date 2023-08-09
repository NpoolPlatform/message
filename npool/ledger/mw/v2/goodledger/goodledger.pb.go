// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.18.1
// source: npool/ledger/mw/v2/goodledger/goodledger.proto

package goodledger

import (
	v1 "github.com/NpoolPlatform/message/npool/basetypes/v1"
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

type GoodLedgerReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID         *string `protobuf:"bytes,10,opt,name=ID,proto3,oneof" json:"ID,omitempty"`
	GoodID     *string `protobuf:"bytes,20,opt,name=GoodID,proto3,oneof" json:"GoodID,omitempty"`
	CoinTypeID *string `protobuf:"bytes,30,opt,name=CoinTypeID,proto3,oneof" json:"CoinTypeID,omitempty"`
	Amount     *string `protobuf:"bytes,40,opt,name=Amount,proto3,oneof" json:"Amount,omitempty"`
	ToPlatform *string `protobuf:"bytes,50,opt,name=ToPlatform,proto3,oneof" json:"ToPlatform,omitempty"`
	ToUser     *string `protobuf:"bytes,60,opt,name=ToUser,proto3,oneof" json:"ToUser,omitempty"`
}

func (x *GoodLedgerReq) Reset() {
	*x = GoodLedgerReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_ledger_mw_v2_goodledger_goodledger_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GoodLedgerReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GoodLedgerReq) ProtoMessage() {}

func (x *GoodLedgerReq) ProtoReflect() protoreflect.Message {
	mi := &file_npool_ledger_mw_v2_goodledger_goodledger_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GoodLedgerReq.ProtoReflect.Descriptor instead.
func (*GoodLedgerReq) Descriptor() ([]byte, []int) {
	return file_npool_ledger_mw_v2_goodledger_goodledger_proto_rawDescGZIP(), []int{0}
}

func (x *GoodLedgerReq) GetID() string {
	if x != nil && x.ID != nil {
		return *x.ID
	}
	return ""
}

func (x *GoodLedgerReq) GetGoodID() string {
	if x != nil && x.GoodID != nil {
		return *x.GoodID
	}
	return ""
}

func (x *GoodLedgerReq) GetCoinTypeID() string {
	if x != nil && x.CoinTypeID != nil {
		return *x.CoinTypeID
	}
	return ""
}

func (x *GoodLedgerReq) GetAmount() string {
	if x != nil && x.Amount != nil {
		return *x.Amount
	}
	return ""
}

func (x *GoodLedgerReq) GetToPlatform() string {
	if x != nil && x.ToPlatform != nil {
		return *x.ToPlatform
	}
	return ""
}

func (x *GoodLedgerReq) GetToUser() string {
	if x != nil && x.ToUser != nil {
		return *x.ToUser
	}
	return ""
}

type GoodLedger struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: sql:"id"
	ID string `protobuf:"bytes,10,opt,name=ID,proto3" json:"ID,omitempty" sql:"id"`
	// @inject_tag: sql:"good_id"
	GoodID string `protobuf:"bytes,20,opt,name=GoodID,proto3" json:"GoodID,omitempty" sql:"good_id"`
	// @inject_tag: sql:"coin_type_id"
	CoinTypeID string `protobuf:"bytes,30,opt,name=CoinTypeID,proto3" json:"CoinTypeID,omitempty" sql:"coin_type_id"`
	// @inject_tag: sql:"amount"
	Amount string `protobuf:"bytes,40,opt,name=Amount,proto3" json:"Amount,omitempty" sql:"amount"`
	// @inject_tag: sql:"to_platform"
	ToPlatform string `protobuf:"bytes,50,opt,name=ToPlatform,proto3" json:"ToPlatform,omitempty" sql:"to_platform"`
	// @inject_tag: sql:"to_user"
	ToUser string `protobuf:"bytes,60,opt,name=ToUser,proto3" json:"ToUser,omitempty" sql:"to_user"`
	// @inject_tag: sql:"created_at"
	CreatedAt uint32 `protobuf:"varint,70,opt,name=CreatedAt,proto3" json:"CreatedAt,omitempty" sql:"created_at"`
	// @inject_tag: sql:"updated_at"
	UpdatedAt uint32 `protobuf:"varint,80,opt,name=UpdatedAt,proto3" json:"UpdatedAt,omitempty" sql:"updated_at"`
}

func (x *GoodLedger) Reset() {
	*x = GoodLedger{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_ledger_mw_v2_goodledger_goodledger_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GoodLedger) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GoodLedger) ProtoMessage() {}

func (x *GoodLedger) ProtoReflect() protoreflect.Message {
	mi := &file_npool_ledger_mw_v2_goodledger_goodledger_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GoodLedger.ProtoReflect.Descriptor instead.
func (*GoodLedger) Descriptor() ([]byte, []int) {
	return file_npool_ledger_mw_v2_goodledger_goodledger_proto_rawDescGZIP(), []int{1}
}

func (x *GoodLedger) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *GoodLedger) GetGoodID() string {
	if x != nil {
		return x.GoodID
	}
	return ""
}

func (x *GoodLedger) GetCoinTypeID() string {
	if x != nil {
		return x.CoinTypeID
	}
	return ""
}

func (x *GoodLedger) GetAmount() string {
	if x != nil {
		return x.Amount
	}
	return ""
}

func (x *GoodLedger) GetToPlatform() string {
	if x != nil {
		return x.ToPlatform
	}
	return ""
}

func (x *GoodLedger) GetToUser() string {
	if x != nil {
		return x.ToUser
	}
	return ""
}

func (x *GoodLedger) GetCreatedAt() uint32 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *GoodLedger) GetUpdatedAt() uint32 {
	if x != nil {
		return x.UpdatedAt
	}
	return 0
}

type Conds struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID         *v1.StringVal `protobuf:"bytes,10,opt,name=ID,proto3,oneof" json:"ID,omitempty"`
	GoodID     *v1.StringVal `protobuf:"bytes,20,opt,name=GoodID,proto3,oneof" json:"GoodID,omitempty"`
	CoinTypeID *v1.StringVal `protobuf:"bytes,30,opt,name=CoinTypeID,proto3,oneof" json:"CoinTypeID,omitempty"`
	Amount     *v1.StringVal `protobuf:"bytes,40,opt,name=Amount,proto3,oneof" json:"Amount,omitempty"`
	ToPlatform *v1.StringVal `protobuf:"bytes,50,opt,name=ToPlatform,proto3,oneof" json:"ToPlatform,omitempty"`
	ToUser     *v1.StringVal `protobuf:"bytes,60,opt,name=ToUser,proto3,oneof" json:"ToUser,omitempty"`
}

func (x *Conds) Reset() {
	*x = Conds{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_ledger_mw_v2_goodledger_goodledger_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Conds) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Conds) ProtoMessage() {}

func (x *Conds) ProtoReflect() protoreflect.Message {
	mi := &file_npool_ledger_mw_v2_goodledger_goodledger_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Conds.ProtoReflect.Descriptor instead.
func (*Conds) Descriptor() ([]byte, []int) {
	return file_npool_ledger_mw_v2_goodledger_goodledger_proto_rawDescGZIP(), []int{2}
}

func (x *Conds) GetID() *v1.StringVal {
	if x != nil {
		return x.ID
	}
	return nil
}

func (x *Conds) GetGoodID() *v1.StringVal {
	if x != nil {
		return x.GoodID
	}
	return nil
}

func (x *Conds) GetCoinTypeID() *v1.StringVal {
	if x != nil {
		return x.CoinTypeID
	}
	return nil
}

func (x *Conds) GetAmount() *v1.StringVal {
	if x != nil {
		return x.Amount
	}
	return nil
}

func (x *Conds) GetToPlatform() *v1.StringVal {
	if x != nil {
		return x.ToPlatform
	}
	return nil
}

func (x *Conds) GetToUser() *v1.StringVal {
	if x != nil {
		return x.ToUser
	}
	return nil
}

type CreateGoodLedgerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *GoodLedgerReq `protobuf:"bytes,10,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (x *CreateGoodLedgerRequest) Reset() {
	*x = CreateGoodLedgerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_ledger_mw_v2_goodledger_goodledger_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateGoodLedgerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateGoodLedgerRequest) ProtoMessage() {}

func (x *CreateGoodLedgerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_ledger_mw_v2_goodledger_goodledger_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateGoodLedgerRequest.ProtoReflect.Descriptor instead.
func (*CreateGoodLedgerRequest) Descriptor() ([]byte, []int) {
	return file_npool_ledger_mw_v2_goodledger_goodledger_proto_rawDescGZIP(), []int{3}
}

func (x *CreateGoodLedgerRequest) GetInfo() *GoodLedgerReq {
	if x != nil {
		return x.Info
	}
	return nil
}

type CreateGoodLedgerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *GoodLedger `protobuf:"bytes,10,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (x *CreateGoodLedgerResponse) Reset() {
	*x = CreateGoodLedgerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_ledger_mw_v2_goodledger_goodledger_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateGoodLedgerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateGoodLedgerResponse) ProtoMessage() {}

func (x *CreateGoodLedgerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_ledger_mw_v2_goodledger_goodledger_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateGoodLedgerResponse.ProtoReflect.Descriptor instead.
func (*CreateGoodLedgerResponse) Descriptor() ([]byte, []int) {
	return file_npool_ledger_mw_v2_goodledger_goodledger_proto_rawDescGZIP(), []int{4}
}

func (x *CreateGoodLedgerResponse) GetInfo() *GoodLedger {
	if x != nil {
		return x.Info
	}
	return nil
}

type GetGoodLedgerOnlyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Conds *Conds `protobuf:"bytes,10,opt,name=Conds,proto3" json:"Conds,omitempty"`
}

func (x *GetGoodLedgerOnlyRequest) Reset() {
	*x = GetGoodLedgerOnlyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_ledger_mw_v2_goodledger_goodledger_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetGoodLedgerOnlyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetGoodLedgerOnlyRequest) ProtoMessage() {}

func (x *GetGoodLedgerOnlyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_ledger_mw_v2_goodledger_goodledger_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetGoodLedgerOnlyRequest.ProtoReflect.Descriptor instead.
func (*GetGoodLedgerOnlyRequest) Descriptor() ([]byte, []int) {
	return file_npool_ledger_mw_v2_goodledger_goodledger_proto_rawDescGZIP(), []int{5}
}

func (x *GetGoodLedgerOnlyRequest) GetConds() *Conds {
	if x != nil {
		return x.Conds
	}
	return nil
}

type GetGoodLedgerOnlyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *GoodLedger `protobuf:"bytes,10,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (x *GetGoodLedgerOnlyResponse) Reset() {
	*x = GetGoodLedgerOnlyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_ledger_mw_v2_goodledger_goodledger_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetGoodLedgerOnlyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetGoodLedgerOnlyResponse) ProtoMessage() {}

func (x *GetGoodLedgerOnlyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_ledger_mw_v2_goodledger_goodledger_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetGoodLedgerOnlyResponse.ProtoReflect.Descriptor instead.
func (*GetGoodLedgerOnlyResponse) Descriptor() ([]byte, []int) {
	return file_npool_ledger_mw_v2_goodledger_goodledger_proto_rawDescGZIP(), []int{6}
}

func (x *GetGoodLedgerOnlyResponse) GetInfo() *GoodLedger {
	if x != nil {
		return x.Info
	}
	return nil
}

type AddGoodLedgerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *GoodLedgerReq `protobuf:"bytes,10,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (x *AddGoodLedgerRequest) Reset() {
	*x = AddGoodLedgerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_ledger_mw_v2_goodledger_goodledger_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddGoodLedgerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddGoodLedgerRequest) ProtoMessage() {}

func (x *AddGoodLedgerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_ledger_mw_v2_goodledger_goodledger_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddGoodLedgerRequest.ProtoReflect.Descriptor instead.
func (*AddGoodLedgerRequest) Descriptor() ([]byte, []int) {
	return file_npool_ledger_mw_v2_goodledger_goodledger_proto_rawDescGZIP(), []int{7}
}

func (x *AddGoodLedgerRequest) GetInfo() *GoodLedgerReq {
	if x != nil {
		return x.Info
	}
	return nil
}

type AddGoodLedgerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *GoodLedger `protobuf:"bytes,10,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (x *AddGoodLedgerResponse) Reset() {
	*x = AddGoodLedgerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_ledger_mw_v2_goodledger_goodledger_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddGoodLedgerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddGoodLedgerResponse) ProtoMessage() {}

func (x *AddGoodLedgerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_ledger_mw_v2_goodledger_goodledger_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddGoodLedgerResponse.ProtoReflect.Descriptor instead.
func (*AddGoodLedgerResponse) Descriptor() ([]byte, []int) {
	return file_npool_ledger_mw_v2_goodledger_goodledger_proto_rawDescGZIP(), []int{8}
}

func (x *AddGoodLedgerResponse) GetInfo() *GoodLedger {
	if x != nil {
		return x.Info
	}
	return nil
}

var File_npool_ledger_mw_v2_goodledger_goodledger_proto protoreflect.FileDescriptor

var file_npool_ledger_mw_v2_goodledger_goodledger_proto_rawDesc = []byte{
	0x0a, 0x2e, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2f, 0x6d,
	0x77, 0x2f, 0x76, 0x32, 0x2f, 0x67, 0x6f, 0x6f, 0x64, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2f,
	0x67, 0x6f, 0x6f, 0x64, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x1f, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77,
	0x61, 0x72, 0x65, 0x2e, 0x67, 0x6f, 0x6f, 0x64, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2e, 0x76,
	0x32, 0x1a, 0x1e, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70,
	0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x8b, 0x02, 0x0a, 0x0d, 0x47, 0x6f, 0x6f, 0x64, 0x4c, 0x65, 0x64, 0x67, 0x65, 0x72,
	0x52, 0x65, 0x71, 0x12, 0x13, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x48,
	0x00, 0x52, 0x02, 0x49, 0x44, 0x88, 0x01, 0x01, 0x12, 0x1b, 0x0a, 0x06, 0x47, 0x6f, 0x6f, 0x64,
	0x49, 0x44, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x06, 0x47, 0x6f, 0x6f, 0x64,
	0x49, 0x44, 0x88, 0x01, 0x01, 0x12, 0x23, 0x0a, 0x0a, 0x43, 0x6f, 0x69, 0x6e, 0x54, 0x79, 0x70,
	0x65, 0x49, 0x44, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x09, 0x48, 0x02, 0x52, 0x0a, 0x43, 0x6f, 0x69,
	0x6e, 0x54, 0x79, 0x70, 0x65, 0x49, 0x44, 0x88, 0x01, 0x01, 0x12, 0x1b, 0x0a, 0x06, 0x41, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x28, 0x20, 0x01, 0x28, 0x09, 0x48, 0x03, 0x52, 0x06, 0x41, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x88, 0x01, 0x01, 0x12, 0x23, 0x0a, 0x0a, 0x54, 0x6f, 0x50, 0x6c, 0x61,
	0x74, 0x66, 0x6f, 0x72, 0x6d, 0x18, 0x32, 0x20, 0x01, 0x28, 0x09, 0x48, 0x04, 0x52, 0x0a, 0x54,
	0x6f, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x88, 0x01, 0x01, 0x12, 0x1b, 0x0a, 0x06,
	0x54, 0x6f, 0x55, 0x73, 0x65, 0x72, 0x18, 0x3c, 0x20, 0x01, 0x28, 0x09, 0x48, 0x05, 0x52, 0x06,
	0x54, 0x6f, 0x55, 0x73, 0x65, 0x72, 0x88, 0x01, 0x01, 0x42, 0x05, 0x0a, 0x03, 0x5f, 0x49, 0x44,
	0x42, 0x09, 0x0a, 0x07, 0x5f, 0x47, 0x6f, 0x6f, 0x64, 0x49, 0x44, 0x42, 0x0d, 0x0a, 0x0b, 0x5f,
	0x43, 0x6f, 0x69, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x49, 0x44, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x41,
	0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x54, 0x6f, 0x50, 0x6c, 0x61, 0x74,
	0x66, 0x6f, 0x72, 0x6d, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x54, 0x6f, 0x55, 0x73, 0x65, 0x72, 0x22,
	0xe0, 0x01, 0x0a, 0x0a, 0x47, 0x6f, 0x6f, 0x64, 0x4c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x12, 0x0e,
	0x0a, 0x02, 0x49, 0x44, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x12, 0x16,
	0x0a, 0x06, 0x47, 0x6f, 0x6f, 0x64, 0x49, 0x44, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x47, 0x6f, 0x6f, 0x64, 0x49, 0x44, 0x12, 0x1e, 0x0a, 0x0a, 0x43, 0x6f, 0x69, 0x6e, 0x54, 0x79,
	0x70, 0x65, 0x49, 0x44, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x43, 0x6f, 0x69, 0x6e,
	0x54, 0x79, 0x70, 0x65, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x28, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1e,
	0x0a, 0x0a, 0x54, 0x6f, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x18, 0x32, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x54, 0x6f, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x12, 0x16,
	0x0a, 0x06, 0x54, 0x6f, 0x55, 0x73, 0x65, 0x72, 0x18, 0x3c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x54, 0x6f, 0x55, 0x73, 0x65, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x18, 0x46, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x18, 0x50, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x22, 0x99, 0x03, 0x0a, 0x05, 0x43, 0x6f, 0x6e, 0x64, 0x73, 0x12, 0x2c, 0x0a, 0x02,
	0x49, 0x44, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61,
	0x6c, 0x48, 0x00, 0x52, 0x02, 0x49, 0x44, 0x88, 0x01, 0x01, 0x12, 0x34, 0x0a, 0x06, 0x47, 0x6f,
	0x6f, 0x64, 0x49, 0x44, 0x18, 0x14, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x62, 0x61, 0x73,
	0x65, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67,
	0x56, 0x61, 0x6c, 0x48, 0x01, 0x52, 0x06, 0x47, 0x6f, 0x6f, 0x64, 0x49, 0x44, 0x88, 0x01, 0x01,
	0x12, 0x3c, 0x0a, 0x0a, 0x43, 0x6f, 0x69, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x49, 0x44, 0x18, 0x1e,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x48, 0x02, 0x52,
	0x0a, 0x43, 0x6f, 0x69, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x49, 0x44, 0x88, 0x01, 0x01, 0x12, 0x34,
	0x0a, 0x06, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x28, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17,
	0x2e, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74,
	0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x48, 0x03, 0x52, 0x06, 0x41, 0x6d, 0x6f, 0x75, 0x6e,
	0x74, 0x88, 0x01, 0x01, 0x12, 0x3c, 0x0a, 0x0a, 0x54, 0x6f, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f,
	0x72, 0x6d, 0x18, 0x32, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61,
	0x6c, 0x48, 0x04, 0x52, 0x0a, 0x54, 0x6f, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x88,
	0x01, 0x01, 0x12, 0x34, 0x0a, 0x06, 0x54, 0x6f, 0x55, 0x73, 0x65, 0x72, 0x18, 0x3c, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x17, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x48, 0x05, 0x52, 0x06, 0x54,
	0x6f, 0x55, 0x73, 0x65, 0x72, 0x88, 0x01, 0x01, 0x42, 0x05, 0x0a, 0x03, 0x5f, 0x49, 0x44, 0x42,
	0x09, 0x0a, 0x07, 0x5f, 0x47, 0x6f, 0x6f, 0x64, 0x49, 0x44, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x43,
	0x6f, 0x69, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x49, 0x44, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x41, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x54, 0x6f, 0x50, 0x6c, 0x61, 0x74, 0x66,
	0x6f, 0x72, 0x6d, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x54, 0x6f, 0x55, 0x73, 0x65, 0x72, 0x22, 0x5d,
	0x0a, 0x17, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x47, 0x6f, 0x6f, 0x64, 0x4c, 0x65, 0x64, 0x67,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x42, 0x0a, 0x04, 0x49, 0x6e, 0x66,
	0x6f, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72,
	0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x67, 0x6f, 0x6f, 0x64,
	0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x32, 0x2e, 0x47, 0x6f, 0x6f, 0x64, 0x4c, 0x65,
	0x64, 0x67, 0x65, 0x72, 0x52, 0x65, 0x71, 0x52, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x5b, 0x0a,
	0x18, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x47, 0x6f, 0x6f, 0x64, 0x4c, 0x65, 0x64, 0x67, 0x65,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3f, 0x0a, 0x04, 0x49, 0x6e, 0x66,
	0x6f, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72,
	0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x67, 0x6f, 0x6f, 0x64,
	0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x32, 0x2e, 0x47, 0x6f, 0x6f, 0x64, 0x4c, 0x65,
	0x64, 0x67, 0x65, 0x72, 0x52, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x58, 0x0a, 0x18, 0x47, 0x65,
	0x74, 0x47, 0x6f, 0x6f, 0x64, 0x4c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x4f, 0x6e, 0x6c, 0x79, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3c, 0x0a, 0x05, 0x43, 0x6f, 0x6e, 0x64, 0x73, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2e, 0x6d,
	0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x67, 0x6f, 0x6f, 0x64, 0x6c, 0x65,
	0x64, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x32, 0x2e, 0x43, 0x6f, 0x6e, 0x64, 0x73, 0x52, 0x05, 0x43,
	0x6f, 0x6e, 0x64, 0x73, 0x22, 0x5c, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x47, 0x6f, 0x6f, 0x64, 0x4c,
	0x65, 0x64, 0x67, 0x65, 0x72, 0x4f, 0x6e, 0x6c, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x3f, 0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x2b, 0x2e, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77,
	0x61, 0x72, 0x65, 0x2e, 0x67, 0x6f, 0x6f, 0x64, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2e, 0x76,
	0x32, 0x2e, 0x47, 0x6f, 0x6f, 0x64, 0x4c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x52, 0x04, 0x49, 0x6e,
	0x66, 0x6f, 0x22, 0x5a, 0x0a, 0x14, 0x41, 0x64, 0x64, 0x47, 0x6f, 0x6f, 0x64, 0x4c, 0x65, 0x64,
	0x67, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x42, 0x0a, 0x04, 0x49, 0x6e,
	0x66, 0x6f, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x6c, 0x65, 0x64, 0x67, 0x65,
	0x72, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x67, 0x6f, 0x6f,
	0x64, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x32, 0x2e, 0x47, 0x6f, 0x6f, 0x64, 0x4c,
	0x65, 0x64, 0x67, 0x65, 0x72, 0x52, 0x65, 0x71, 0x52, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x58,
	0x0a, 0x15, 0x41, 0x64, 0x64, 0x47, 0x6f, 0x6f, 0x64, 0x4c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3f, 0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2e, 0x6d,
	0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x67, 0x6f, 0x6f, 0x64, 0x6c, 0x65,
	0x64, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x32, 0x2e, 0x47, 0x6f, 0x6f, 0x64, 0x4c, 0x65, 0x64, 0x67,
	0x65, 0x72, 0x52, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x32, 0xaa, 0x03, 0x0a, 0x0a, 0x4d, 0x69, 0x64,
	0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x12, 0x89, 0x01, 0x0a, 0x10, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x47, 0x6f, 0x6f, 0x64, 0x4c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x12, 0x38, 0x2e, 0x6c,
	0x65, 0x64, 0x67, 0x65, 0x72, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65,
	0x2e, 0x67, 0x6f, 0x6f, 0x64, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x32, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x47, 0x6f, 0x6f, 0x64, 0x4c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x39, 0x2e, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2e,
	0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x67, 0x6f, 0x6f, 0x64, 0x6c,
	0x65, 0x64, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x32, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x47,
	0x6f, 0x6f, 0x64, 0x4c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x8c, 0x01, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x47, 0x6f, 0x6f, 0x64, 0x4c,
	0x65, 0x64, 0x67, 0x65, 0x72, 0x4f, 0x6e, 0x6c, 0x79, 0x12, 0x39, 0x2e, 0x6c, 0x65, 0x64, 0x67,
	0x65, 0x72, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x67, 0x6f,
	0x6f, 0x64, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x32, 0x2e, 0x47, 0x65, 0x74, 0x47,
	0x6f, 0x6f, 0x64, 0x4c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x4f, 0x6e, 0x6c, 0x79, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x3a, 0x2e, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2e, 0x6d, 0x69,
	0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x67, 0x6f, 0x6f, 0x64, 0x6c, 0x65, 0x64,
	0x67, 0x65, 0x72, 0x2e, 0x76, 0x32, 0x2e, 0x47, 0x65, 0x74, 0x47, 0x6f, 0x6f, 0x64, 0x4c, 0x65,
	0x64, 0x67, 0x65, 0x72, 0x4f, 0x6e, 0x6c, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x80, 0x01, 0x0a, 0x0d, 0x41, 0x64, 0x64, 0x47, 0x6f, 0x6f, 0x64, 0x4c, 0x65,
	0x64, 0x67, 0x65, 0x72, 0x12, 0x35, 0x2e, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2e, 0x6d, 0x69,
	0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x67, 0x6f, 0x6f, 0x64, 0x6c, 0x65, 0x64,
	0x67, 0x65, 0x72, 0x2e, 0x76, 0x32, 0x2e, 0x41, 0x64, 0x64, 0x47, 0x6f, 0x6f, 0x64, 0x4c, 0x65,
	0x64, 0x67, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x36, 0x2e, 0x6c, 0x65,
	0x64, 0x67, 0x65, 0x72, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e,
	0x67, 0x6f, 0x6f, 0x64, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x32, 0x2e, 0x41, 0x64,
	0x64, 0x47, 0x6f, 0x6f, 0x64, 0x4c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x40, 0x5a, 0x3e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x4e, 0x70, 0x6f, 0x6f, 0x6c, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72,
	0x6d, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f,
	0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2f, 0x6d, 0x77, 0x2f, 0x76, 0x32, 0x2f, 0x67, 0x6f, 0x6f,
	0x64, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_npool_ledger_mw_v2_goodledger_goodledger_proto_rawDescOnce sync.Once
	file_npool_ledger_mw_v2_goodledger_goodledger_proto_rawDescData = file_npool_ledger_mw_v2_goodledger_goodledger_proto_rawDesc
)

func file_npool_ledger_mw_v2_goodledger_goodledger_proto_rawDescGZIP() []byte {
	file_npool_ledger_mw_v2_goodledger_goodledger_proto_rawDescOnce.Do(func() {
		file_npool_ledger_mw_v2_goodledger_goodledger_proto_rawDescData = protoimpl.X.CompressGZIP(file_npool_ledger_mw_v2_goodledger_goodledger_proto_rawDescData)
	})
	return file_npool_ledger_mw_v2_goodledger_goodledger_proto_rawDescData
}

var file_npool_ledger_mw_v2_goodledger_goodledger_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_npool_ledger_mw_v2_goodledger_goodledger_proto_goTypes = []interface{}{
	(*GoodLedgerReq)(nil),             // 0: ledger.middleware.goodledger.v2.GoodLedgerReq
	(*GoodLedger)(nil),                // 1: ledger.middleware.goodledger.v2.GoodLedger
	(*Conds)(nil),                     // 2: ledger.middleware.goodledger.v2.Conds
	(*CreateGoodLedgerRequest)(nil),   // 3: ledger.middleware.goodledger.v2.CreateGoodLedgerRequest
	(*CreateGoodLedgerResponse)(nil),  // 4: ledger.middleware.goodledger.v2.CreateGoodLedgerResponse
	(*GetGoodLedgerOnlyRequest)(nil),  // 5: ledger.middleware.goodledger.v2.GetGoodLedgerOnlyRequest
	(*GetGoodLedgerOnlyResponse)(nil), // 6: ledger.middleware.goodledger.v2.GetGoodLedgerOnlyResponse
	(*AddGoodLedgerRequest)(nil),      // 7: ledger.middleware.goodledger.v2.AddGoodLedgerRequest
	(*AddGoodLedgerResponse)(nil),     // 8: ledger.middleware.goodledger.v2.AddGoodLedgerResponse
	(*v1.StringVal)(nil),              // 9: basetypes.v1.StringVal
}
var file_npool_ledger_mw_v2_goodledger_goodledger_proto_depIdxs = []int32{
	9,  // 0: ledger.middleware.goodledger.v2.Conds.ID:type_name -> basetypes.v1.StringVal
	9,  // 1: ledger.middleware.goodledger.v2.Conds.GoodID:type_name -> basetypes.v1.StringVal
	9,  // 2: ledger.middleware.goodledger.v2.Conds.CoinTypeID:type_name -> basetypes.v1.StringVal
	9,  // 3: ledger.middleware.goodledger.v2.Conds.Amount:type_name -> basetypes.v1.StringVal
	9,  // 4: ledger.middleware.goodledger.v2.Conds.ToPlatform:type_name -> basetypes.v1.StringVal
	9,  // 5: ledger.middleware.goodledger.v2.Conds.ToUser:type_name -> basetypes.v1.StringVal
	0,  // 6: ledger.middleware.goodledger.v2.CreateGoodLedgerRequest.Info:type_name -> ledger.middleware.goodledger.v2.GoodLedgerReq
	1,  // 7: ledger.middleware.goodledger.v2.CreateGoodLedgerResponse.Info:type_name -> ledger.middleware.goodledger.v2.GoodLedger
	2,  // 8: ledger.middleware.goodledger.v2.GetGoodLedgerOnlyRequest.Conds:type_name -> ledger.middleware.goodledger.v2.Conds
	1,  // 9: ledger.middleware.goodledger.v2.GetGoodLedgerOnlyResponse.Info:type_name -> ledger.middleware.goodledger.v2.GoodLedger
	0,  // 10: ledger.middleware.goodledger.v2.AddGoodLedgerRequest.Info:type_name -> ledger.middleware.goodledger.v2.GoodLedgerReq
	1,  // 11: ledger.middleware.goodledger.v2.AddGoodLedgerResponse.Info:type_name -> ledger.middleware.goodledger.v2.GoodLedger
	3,  // 12: ledger.middleware.goodledger.v2.Middleware.CreateGoodLedger:input_type -> ledger.middleware.goodledger.v2.CreateGoodLedgerRequest
	5,  // 13: ledger.middleware.goodledger.v2.Middleware.GetGoodLedgerOnly:input_type -> ledger.middleware.goodledger.v2.GetGoodLedgerOnlyRequest
	7,  // 14: ledger.middleware.goodledger.v2.Middleware.AddGoodLedger:input_type -> ledger.middleware.goodledger.v2.AddGoodLedgerRequest
	4,  // 15: ledger.middleware.goodledger.v2.Middleware.CreateGoodLedger:output_type -> ledger.middleware.goodledger.v2.CreateGoodLedgerResponse
	6,  // 16: ledger.middleware.goodledger.v2.Middleware.GetGoodLedgerOnly:output_type -> ledger.middleware.goodledger.v2.GetGoodLedgerOnlyResponse
	8,  // 17: ledger.middleware.goodledger.v2.Middleware.AddGoodLedger:output_type -> ledger.middleware.goodledger.v2.AddGoodLedgerResponse
	15, // [15:18] is the sub-list for method output_type
	12, // [12:15] is the sub-list for method input_type
	12, // [12:12] is the sub-list for extension type_name
	12, // [12:12] is the sub-list for extension extendee
	0,  // [0:12] is the sub-list for field type_name
}

func init() { file_npool_ledger_mw_v2_goodledger_goodledger_proto_init() }
func file_npool_ledger_mw_v2_goodledger_goodledger_proto_init() {
	if File_npool_ledger_mw_v2_goodledger_goodledger_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_npool_ledger_mw_v2_goodledger_goodledger_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GoodLedgerReq); i {
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
		file_npool_ledger_mw_v2_goodledger_goodledger_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GoodLedger); i {
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
		file_npool_ledger_mw_v2_goodledger_goodledger_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Conds); i {
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
		file_npool_ledger_mw_v2_goodledger_goodledger_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateGoodLedgerRequest); i {
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
		file_npool_ledger_mw_v2_goodledger_goodledger_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateGoodLedgerResponse); i {
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
		file_npool_ledger_mw_v2_goodledger_goodledger_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetGoodLedgerOnlyRequest); i {
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
		file_npool_ledger_mw_v2_goodledger_goodledger_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetGoodLedgerOnlyResponse); i {
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
		file_npool_ledger_mw_v2_goodledger_goodledger_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddGoodLedgerRequest); i {
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
		file_npool_ledger_mw_v2_goodledger_goodledger_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddGoodLedgerResponse); i {
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
	file_npool_ledger_mw_v2_goodledger_goodledger_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_npool_ledger_mw_v2_goodledger_goodledger_proto_msgTypes[2].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_npool_ledger_mw_v2_goodledger_goodledger_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_npool_ledger_mw_v2_goodledger_goodledger_proto_goTypes,
		DependencyIndexes: file_npool_ledger_mw_v2_goodledger_goodledger_proto_depIdxs,
		MessageInfos:      file_npool_ledger_mw_v2_goodledger_goodledger_proto_msgTypes,
	}.Build()
	File_npool_ledger_mw_v2_goodledger_goodledger_proto = out.File
	file_npool_ledger_mw_v2_goodledger_goodledger_proto_rawDesc = nil
	file_npool_ledger_mw_v2_goodledger_goodledger_proto_goTypes = nil
	file_npool_ledger_mw_v2_goodledger_goodledger_proto_depIdxs = nil
}
