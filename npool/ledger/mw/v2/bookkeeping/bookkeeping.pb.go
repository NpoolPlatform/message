// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.18.1
// source: npool/ledger/mw/v2/bookkeeping/bookkeeping.proto

package bookkeeping

import (
	v1 "github.com/NpoolPlatform/message/npool/basetypes/ledger/v1"
	statement "github.com/NpoolPlatform/message/npool/ledger/mw/v2/statement"
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

type BookKeepingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Infos []*statement.StatementReq `protobuf:"bytes,10,rep,name=Infos,proto3" json:"Infos,omitempty"`
}

func (x *BookKeepingRequest) Reset() {
	*x = BookKeepingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_ledger_mw_v2_bookkeeping_bookkeeping_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BookKeepingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BookKeepingRequest) ProtoMessage() {}

func (x *BookKeepingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_ledger_mw_v2_bookkeeping_bookkeeping_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BookKeepingRequest.ProtoReflect.Descriptor instead.
func (*BookKeepingRequest) Descriptor() ([]byte, []int) {
	return file_npool_ledger_mw_v2_bookkeeping_bookkeeping_proto_rawDescGZIP(), []int{0}
}

func (x *BookKeepingRequest) GetInfos() []*statement.StatementReq {
	if x != nil {
		return x.Infos
	}
	return nil
}

type BookKeepingResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *BookKeepingResponse) Reset() {
	*x = BookKeepingResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_ledger_mw_v2_bookkeeping_bookkeeping_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BookKeepingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BookKeepingResponse) ProtoMessage() {}

func (x *BookKeepingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_ledger_mw_v2_bookkeeping_bookkeeping_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BookKeepingResponse.ProtoReflect.Descriptor instead.
func (*BookKeepingResponse) Descriptor() ([]byte, []int) {
	return file_npool_ledger_mw_v2_bookkeeping_bookkeeping_proto_rawDescGZIP(), []int{1}
}

type LockBalanceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AppID      string `protobuf:"bytes,10,opt,name=AppID,proto3" json:"AppID,omitempty"`
	UserID     string `protobuf:"bytes,20,opt,name=UserID,proto3" json:"UserID,omitempty"`
	CoinTypeID string `protobuf:"bytes,30,opt,name=CoinTypeID,proto3" json:"CoinTypeID,omitempty"`
	Amount     string `protobuf:"bytes,40,opt,name=Amount,proto3" json:"Amount,omitempty"`
}

func (x *LockBalanceRequest) Reset() {
	*x = LockBalanceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_ledger_mw_v2_bookkeeping_bookkeeping_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LockBalanceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LockBalanceRequest) ProtoMessage() {}

func (x *LockBalanceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_ledger_mw_v2_bookkeeping_bookkeeping_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LockBalanceRequest.ProtoReflect.Descriptor instead.
func (*LockBalanceRequest) Descriptor() ([]byte, []int) {
	return file_npool_ledger_mw_v2_bookkeeping_bookkeeping_proto_rawDescGZIP(), []int{2}
}

func (x *LockBalanceRequest) GetAppID() string {
	if x != nil {
		return x.AppID
	}
	return ""
}

func (x *LockBalanceRequest) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

func (x *LockBalanceRequest) GetCoinTypeID() string {
	if x != nil {
		return x.CoinTypeID
	}
	return ""
}

func (x *LockBalanceRequest) GetAmount() string {
	if x != nil {
		return x.Amount
	}
	return ""
}

type LockBalanceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *LockBalanceResponse) Reset() {
	*x = LockBalanceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_ledger_mw_v2_bookkeeping_bookkeeping_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LockBalanceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LockBalanceResponse) ProtoMessage() {}

func (x *LockBalanceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_ledger_mw_v2_bookkeeping_bookkeeping_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LockBalanceResponse.ProtoReflect.Descriptor instead.
func (*LockBalanceResponse) Descriptor() ([]byte, []int) {
	return file_npool_ledger_mw_v2_bookkeeping_bookkeeping_proto_rawDescGZIP(), []int{3}
}

type UnlockBalanceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AppID      string `protobuf:"bytes,10,opt,name=AppID,proto3" json:"AppID,omitempty"`
	UserID     string `protobuf:"bytes,20,opt,name=UserID,proto3" json:"UserID,omitempty"`
	CoinTypeID string `protobuf:"bytes,30,opt,name=CoinTypeID,proto3" json:"CoinTypeID,omitempty"`
	Unlocked   string `protobuf:"bytes,40,opt,name=Unlocked,proto3" json:"Unlocked,omitempty"`
	Outcoming  string `protobuf:"bytes,50,opt,name=Outcoming,proto3" json:"Outcoming,omitempty"`
	// Only Payment or Withdrawal
	IOSubType v1.IOSubType `protobuf:"varint,60,opt,name=IOSubType,proto3,enum=basetypes.ledger.v1.IOSubType" json:"IOSubType,omitempty"`
	IOExtra   string       `protobuf:"bytes,70,opt,name=IOExtra,proto3" json:"IOExtra,omitempty"`
}

func (x *UnlockBalanceRequest) Reset() {
	*x = UnlockBalanceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_ledger_mw_v2_bookkeeping_bookkeeping_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnlockBalanceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnlockBalanceRequest) ProtoMessage() {}

func (x *UnlockBalanceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_ledger_mw_v2_bookkeeping_bookkeeping_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnlockBalanceRequest.ProtoReflect.Descriptor instead.
func (*UnlockBalanceRequest) Descriptor() ([]byte, []int) {
	return file_npool_ledger_mw_v2_bookkeeping_bookkeeping_proto_rawDescGZIP(), []int{4}
}

func (x *UnlockBalanceRequest) GetAppID() string {
	if x != nil {
		return x.AppID
	}
	return ""
}

func (x *UnlockBalanceRequest) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

func (x *UnlockBalanceRequest) GetCoinTypeID() string {
	if x != nil {
		return x.CoinTypeID
	}
	return ""
}

func (x *UnlockBalanceRequest) GetUnlocked() string {
	if x != nil {
		return x.Unlocked
	}
	return ""
}

func (x *UnlockBalanceRequest) GetOutcoming() string {
	if x != nil {
		return x.Outcoming
	}
	return ""
}

func (x *UnlockBalanceRequest) GetIOSubType() v1.IOSubType {
	if x != nil {
		return x.IOSubType
	}
	return v1.IOSubType(0)
}

func (x *UnlockBalanceRequest) GetIOExtra() string {
	if x != nil {
		return x.IOExtra
	}
	return ""
}

type UnlockBalanceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UnlockBalanceResponse) Reset() {
	*x = UnlockBalanceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_ledger_mw_v2_bookkeeping_bookkeeping_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnlockBalanceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnlockBalanceResponse) ProtoMessage() {}

func (x *UnlockBalanceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_ledger_mw_v2_bookkeeping_bookkeeping_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnlockBalanceResponse.ProtoReflect.Descriptor instead.
func (*UnlockBalanceResponse) Descriptor() ([]byte, []int) {
	return file_npool_ledger_mw_v2_bookkeeping_bookkeeping_proto_rawDescGZIP(), []int{5}
}

type FinalPaymentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Infos []*statement.StatementReq `protobuf:"bytes,10,rep,name=Infos,proto3" json:"Infos,omitempty"`
}

func (x *FinalPaymentRequest) Reset() {
	*x = FinalPaymentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_ledger_mw_v2_bookkeeping_bookkeeping_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FinalPaymentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FinalPaymentRequest) ProtoMessage() {}

func (x *FinalPaymentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_ledger_mw_v2_bookkeeping_bookkeeping_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FinalPaymentRequest.ProtoReflect.Descriptor instead.
func (*FinalPaymentRequest) Descriptor() ([]byte, []int) {
	return file_npool_ledger_mw_v2_bookkeeping_bookkeeping_proto_rawDescGZIP(), []int{6}
}

func (x *FinalPaymentRequest) GetInfos() []*statement.StatementReq {
	if x != nil {
		return x.Infos
	}
	return nil
}

type FinalPaymentResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *FinalPaymentResponse) Reset() {
	*x = FinalPaymentResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_ledger_mw_v2_bookkeeping_bookkeeping_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FinalPaymentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FinalPaymentResponse) ProtoMessage() {}

func (x *FinalPaymentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_ledger_mw_v2_bookkeeping_bookkeeping_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FinalPaymentResponse.ProtoReflect.Descriptor instead.
func (*FinalPaymentResponse) Descriptor() ([]byte, []int) {
	return file_npool_ledger_mw_v2_bookkeeping_bookkeeping_proto_rawDescGZIP(), []int{7}
}

var File_npool_ledger_mw_v2_bookkeeping_bookkeeping_proto protoreflect.FileDescriptor

var file_npool_ledger_mw_v2_bookkeeping_bookkeeping_proto_rawDesc = []byte{
	0x0a, 0x30, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2f, 0x6d,
	0x77, 0x2f, 0x76, 0x32, 0x2f, 0x62, 0x6f, 0x6f, 0x6b, 0x6b, 0x65, 0x65, 0x70, 0x69, 0x6e, 0x67,
	0x2f, 0x62, 0x6f, 0x6f, 0x6b, 0x6b, 0x65, 0x65, 0x70, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x20, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c,
	0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x6b, 0x65, 0x65, 0x70, 0x69, 0x6e,
	0x67, 0x2e, 0x76, 0x32, 0x1a, 0x2c, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x6c, 0x65, 0x64, 0x67,
	0x65, 0x72, 0x2f, 0x6d, 0x77, 0x2f, 0x76, 0x32, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x6d, 0x65,
	0x6e, 0x74, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x25, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x2f, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x6e,
	0x75, 0x6d, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x58, 0x0a, 0x12, 0x42, 0x6f, 0x6f,
	0x6b, 0x4b, 0x65, 0x65, 0x70, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x42, 0x0a, 0x05, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2c,
	0x2e, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61,
	0x72, 0x65, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x32, 0x2e,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x52, 0x05, 0x49, 0x6e,
	0x66, 0x6f, 0x73, 0x22, 0x15, 0x0a, 0x13, 0x42, 0x6f, 0x6f, 0x6b, 0x4b, 0x65, 0x65, 0x70, 0x69,
	0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x7a, 0x0a, 0x12, 0x4c, 0x6f,
	0x63, 0x6b, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x14, 0x0a, 0x05, 0x41, 0x70, 0x70, 0x49, 0x44, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x41, 0x70, 0x70, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44,
	0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x1e,
	0x0a, 0x0a, 0x43, 0x6f, 0x69, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x49, 0x44, 0x18, 0x1e, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x43, 0x6f, 0x69, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x49, 0x44, 0x12, 0x16,
	0x0a, 0x06, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x28, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x15, 0x0a, 0x13, 0x4c, 0x6f, 0x63, 0x6b, 0x42, 0x61,
	0x6c, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0xf6, 0x01,
	0x0a, 0x14, 0x55, 0x6e, 0x6c, 0x6f, 0x63, 0x6b, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x41, 0x70, 0x70, 0x49, 0x44, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x41, 0x70, 0x70, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06,
	0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x55, 0x73,
	0x65, 0x72, 0x49, 0x44, 0x12, 0x1e, 0x0a, 0x0a, 0x43, 0x6f, 0x69, 0x6e, 0x54, 0x79, 0x70, 0x65,
	0x49, 0x44, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x43, 0x6f, 0x69, 0x6e, 0x54, 0x79,
	0x70, 0x65, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08, 0x55, 0x6e, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x64,
	0x18, 0x28, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x55, 0x6e, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x64,
	0x12, 0x1c, 0x0a, 0x09, 0x4f, 0x75, 0x74, 0x63, 0x6f, 0x6d, 0x69, 0x6e, 0x67, 0x18, 0x32, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x4f, 0x75, 0x74, 0x63, 0x6f, 0x6d, 0x69, 0x6e, 0x67, 0x12, 0x3c,
	0x0a, 0x09, 0x49, 0x4f, 0x53, 0x75, 0x62, 0x54, 0x79, 0x70, 0x65, 0x18, 0x3c, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x1e, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x6c, 0x65,
	0x64, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x4f, 0x53, 0x75, 0x62, 0x54, 0x79, 0x70,
	0x65, 0x52, 0x09, 0x49, 0x4f, 0x53, 0x75, 0x62, 0x54, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x49, 0x4f, 0x45, 0x78, 0x74, 0x72, 0x61, 0x18, 0x46, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x49,
	0x4f, 0x45, 0x78, 0x74, 0x72, 0x61, 0x22, 0x17, 0x0a, 0x15, 0x55, 0x6e, 0x6c, 0x6f, 0x63, 0x6b,
	0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x59, 0x0a, 0x13, 0x46, 0x69, 0x6e, 0x61, 0x6c, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x42, 0x0a, 0x05, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x18,
	0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2e, 0x6d,
	0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x65, 0x6d,
	0x65, 0x6e, 0x74, 0x2e, 0x76, 0x32, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x6d, 0x65, 0x6e, 0x74,
	0x52, 0x65, 0x71, 0x52, 0x05, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x22, 0x16, 0x0a, 0x14, 0x46, 0x69,
	0x6e, 0x61, 0x6c, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x32, 0x8e, 0x04, 0x0a, 0x0a, 0x4d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72,
	0x65, 0x12, 0x7c, 0x0a, 0x0b, 0x42, 0x6f, 0x6f, 0x6b, 0x4b, 0x65, 0x65, 0x70, 0x69, 0x6e, 0x67,
	0x12, 0x34, 0x2e, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65,
	0x77, 0x61, 0x72, 0x65, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x6b, 0x65, 0x65, 0x70, 0x69, 0x6e, 0x67,
	0x2e, 0x76, 0x32, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x4b, 0x65, 0x65, 0x70, 0x69, 0x6e, 0x67, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x35, 0x2e, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2e,
	0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x6b,
	0x65, 0x65, 0x70, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x32, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x4b, 0x65,
	0x65, 0x70, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x7c, 0x0a, 0x0b, 0x4c, 0x6f, 0x63, 0x6b, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x34,
	0x2e, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61,
	0x72, 0x65, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x6b, 0x65, 0x65, 0x70, 0x69, 0x6e, 0x67, 0x2e, 0x76,
	0x32, 0x2e, 0x4c, 0x6f, 0x63, 0x6b, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x35, 0x2e, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2e, 0x6d, 0x69,
	0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x6b, 0x65, 0x65,
	0x70, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x32, 0x2e, 0x4c, 0x6f, 0x63, 0x6b, 0x42, 0x61, 0x6c, 0x61,
	0x6e, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x82, 0x01,
	0x0a, 0x0d, 0x55, 0x6e, 0x6c, 0x6f, 0x63, 0x6b, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x12,
	0x36, 0x2e, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77,
	0x61, 0x72, 0x65, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x6b, 0x65, 0x65, 0x70, 0x69, 0x6e, 0x67, 0x2e,
	0x76, 0x32, 0x2e, 0x55, 0x6e, 0x6c, 0x6f, 0x63, 0x6b, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x37, 0x2e, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72,
	0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x62, 0x6f, 0x6f, 0x6b,
	0x6b, 0x65, 0x65, 0x70, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x32, 0x2e, 0x55, 0x6e, 0x6c, 0x6f, 0x63,
	0x6b, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x7f, 0x0a, 0x0c, 0x46, 0x69, 0x6e, 0x61, 0x6c, 0x50, 0x61, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x12, 0x35, 0x2e, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2e, 0x6d, 0x69, 0x64, 0x64,
	0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x6b, 0x65, 0x65, 0x70, 0x69,
	0x6e, 0x67, 0x2e, 0x76, 0x32, 0x2e, 0x46, 0x69, 0x6e, 0x61, 0x6c, 0x50, 0x61, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x36, 0x2e, 0x6c, 0x65, 0x64, 0x67,
	0x65, 0x72, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x62, 0x6f,
	0x6f, 0x6b, 0x6b, 0x65, 0x65, 0x70, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x32, 0x2e, 0x46, 0x69, 0x6e,
	0x61, 0x6c, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x42, 0x41, 0x5a, 0x3f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x4e, 0x70, 0x6f, 0x6f, 0x6c, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x6c, 0x65,
	0x64, 0x67, 0x65, 0x72, 0x2f, 0x6d, 0x77, 0x2f, 0x76, 0x32, 0x2f, 0x62, 0x6f, 0x6f, 0x6b, 0x6b,
	0x65, 0x65, 0x70, 0x69, 0x6e, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_npool_ledger_mw_v2_bookkeeping_bookkeeping_proto_rawDescOnce sync.Once
	file_npool_ledger_mw_v2_bookkeeping_bookkeeping_proto_rawDescData = file_npool_ledger_mw_v2_bookkeeping_bookkeeping_proto_rawDesc
)

func file_npool_ledger_mw_v2_bookkeeping_bookkeeping_proto_rawDescGZIP() []byte {
	file_npool_ledger_mw_v2_bookkeeping_bookkeeping_proto_rawDescOnce.Do(func() {
		file_npool_ledger_mw_v2_bookkeeping_bookkeeping_proto_rawDescData = protoimpl.X.CompressGZIP(file_npool_ledger_mw_v2_bookkeeping_bookkeeping_proto_rawDescData)
	})
	return file_npool_ledger_mw_v2_bookkeeping_bookkeeping_proto_rawDescData
}

var file_npool_ledger_mw_v2_bookkeeping_bookkeeping_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_npool_ledger_mw_v2_bookkeeping_bookkeeping_proto_goTypes = []interface{}{
	(*BookKeepingRequest)(nil),     // 0: ledger.middleware.bookkeeping.v2.BookKeepingRequest
	(*BookKeepingResponse)(nil),    // 1: ledger.middleware.bookkeeping.v2.BookKeepingResponse
	(*LockBalanceRequest)(nil),     // 2: ledger.middleware.bookkeeping.v2.LockBalanceRequest
	(*LockBalanceResponse)(nil),    // 3: ledger.middleware.bookkeeping.v2.LockBalanceResponse
	(*UnlockBalanceRequest)(nil),   // 4: ledger.middleware.bookkeeping.v2.UnlockBalanceRequest
	(*UnlockBalanceResponse)(nil),  // 5: ledger.middleware.bookkeeping.v2.UnlockBalanceResponse
	(*FinalPaymentRequest)(nil),    // 6: ledger.middleware.bookkeeping.v2.FinalPaymentRequest
	(*FinalPaymentResponse)(nil),   // 7: ledger.middleware.bookkeeping.v2.FinalPaymentResponse
	(*statement.StatementReq)(nil), // 8: ledger.middleware.statement.v2.StatementReq
	(v1.IOSubType)(0),              // 9: basetypes.ledger.v1.IOSubType
}
var file_npool_ledger_mw_v2_bookkeeping_bookkeeping_proto_depIdxs = []int32{
	8, // 0: ledger.middleware.bookkeeping.v2.BookKeepingRequest.Infos:type_name -> ledger.middleware.statement.v2.StatementReq
	9, // 1: ledger.middleware.bookkeeping.v2.UnlockBalanceRequest.IOSubType:type_name -> basetypes.ledger.v1.IOSubType
	8, // 2: ledger.middleware.bookkeeping.v2.FinalPaymentRequest.Infos:type_name -> ledger.middleware.statement.v2.StatementReq
	0, // 3: ledger.middleware.bookkeeping.v2.Middleware.BookKeeping:input_type -> ledger.middleware.bookkeeping.v2.BookKeepingRequest
	2, // 4: ledger.middleware.bookkeeping.v2.Middleware.LockBalance:input_type -> ledger.middleware.bookkeeping.v2.LockBalanceRequest
	4, // 5: ledger.middleware.bookkeeping.v2.Middleware.UnlockBalance:input_type -> ledger.middleware.bookkeeping.v2.UnlockBalanceRequest
	6, // 6: ledger.middleware.bookkeeping.v2.Middleware.FinalPayment:input_type -> ledger.middleware.bookkeeping.v2.FinalPaymentRequest
	1, // 7: ledger.middleware.bookkeeping.v2.Middleware.BookKeeping:output_type -> ledger.middleware.bookkeeping.v2.BookKeepingResponse
	3, // 8: ledger.middleware.bookkeeping.v2.Middleware.LockBalance:output_type -> ledger.middleware.bookkeeping.v2.LockBalanceResponse
	5, // 9: ledger.middleware.bookkeeping.v2.Middleware.UnlockBalance:output_type -> ledger.middleware.bookkeeping.v2.UnlockBalanceResponse
	7, // 10: ledger.middleware.bookkeeping.v2.Middleware.FinalPayment:output_type -> ledger.middleware.bookkeeping.v2.FinalPaymentResponse
	7, // [7:11] is the sub-list for method output_type
	3, // [3:7] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_npool_ledger_mw_v2_bookkeeping_bookkeeping_proto_init() }
func file_npool_ledger_mw_v2_bookkeeping_bookkeeping_proto_init() {
	if File_npool_ledger_mw_v2_bookkeeping_bookkeeping_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_npool_ledger_mw_v2_bookkeeping_bookkeeping_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BookKeepingRequest); i {
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
		file_npool_ledger_mw_v2_bookkeeping_bookkeeping_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BookKeepingResponse); i {
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
		file_npool_ledger_mw_v2_bookkeeping_bookkeeping_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LockBalanceRequest); i {
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
		file_npool_ledger_mw_v2_bookkeeping_bookkeeping_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LockBalanceResponse); i {
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
		file_npool_ledger_mw_v2_bookkeeping_bookkeeping_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UnlockBalanceRequest); i {
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
		file_npool_ledger_mw_v2_bookkeeping_bookkeeping_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UnlockBalanceResponse); i {
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
		file_npool_ledger_mw_v2_bookkeeping_bookkeeping_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FinalPaymentRequest); i {
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
		file_npool_ledger_mw_v2_bookkeeping_bookkeeping_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FinalPaymentResponse); i {
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
			RawDescriptor: file_npool_ledger_mw_v2_bookkeeping_bookkeeping_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_npool_ledger_mw_v2_bookkeeping_bookkeeping_proto_goTypes,
		DependencyIndexes: file_npool_ledger_mw_v2_bookkeeping_bookkeeping_proto_depIdxs,
		MessageInfos:      file_npool_ledger_mw_v2_bookkeeping_bookkeeping_proto_msgTypes,
	}.Build()
	File_npool_ledger_mw_v2_bookkeeping_bookkeeping_proto = out.File
	file_npool_ledger_mw_v2_bookkeeping_bookkeeping_proto_rawDesc = nil
	file_npool_ledger_mw_v2_bookkeeping_bookkeeping_proto_goTypes = nil
	file_npool_ledger_mw_v2_bookkeeping_bookkeeping_proto_depIdxs = nil
}
