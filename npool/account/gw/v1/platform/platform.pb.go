// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.18.1
// source: npool/account/gw/v1/platform/platform.proto

package platform

import (
	v1 "github.com/NpoolPlatform/message/npool/basetypes/v1"
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

type Account struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID         uint32             `protobuf:"varint,9,opt,name=ID,proto3" json:"ID,omitempty"`
	EntID      string             `protobuf:"bytes,10,opt,name=EntID,proto3" json:"EntID,omitempty"`
	CoinTypeID string             `protobuf:"bytes,20,opt,name=CoinTypeID,proto3" json:"CoinTypeID,omitempty"`
	CoinName   string             `protobuf:"bytes,30,opt,name=CoinName,proto3" json:"CoinName,omitempty"`
	CoinUnit   string             `protobuf:"bytes,40,opt,name=CoinUnit,proto3" json:"CoinUnit,omitempty"`
	CoinEnv    string             `protobuf:"bytes,50,opt,name=CoinEnv,proto3" json:"CoinEnv,omitempty"`
	CoinLogo   string             `protobuf:"bytes,60,opt,name=CoinLogo,proto3" json:"CoinLogo,omitempty"`
	UsedFor    v1.AccountUsedFor  `protobuf:"varint,70,opt,name=UsedFor,proto3,enum=basetypes.v1.AccountUsedFor" json:"UsedFor,omitempty"`
	AccountID  string             `protobuf:"bytes,80,opt,name=AccountID,proto3" json:"AccountID,omitempty"`
	Address    string             `protobuf:"bytes,90,opt,name=Address,proto3" json:"Address,omitempty"`
	Backup     bool               `protobuf:"varint,100,opt,name=Backup,proto3" json:"Backup,omitempty"`
	Active     bool               `protobuf:"varint,110,opt,name=Active,proto3" json:"Active,omitempty"`
	Locked     bool               `protobuf:"varint,120,opt,name=Locked,proto3" json:"Locked,omitempty"`
	LockedBy   v1.AccountLockedBy `protobuf:"varint,130,opt,name=LockedBy,proto3,enum=basetypes.v1.AccountLockedBy" json:"LockedBy,omitempty"`
	Blocked    bool               `protobuf:"varint,140,opt,name=Blocked,proto3" json:"Blocked,omitempty"`
	CreatedAt  uint32             `protobuf:"varint,150,opt,name=CreatedAt,proto3" json:"CreatedAt,omitempty"`
	UpdatedAt  uint32             `protobuf:"varint,160,opt,name=UpdatedAt,proto3" json:"UpdatedAt,omitempty"`
}

func (x *Account) Reset() {
	*x = Account{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_account_gw_v1_platform_platform_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Account) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Account) ProtoMessage() {}

func (x *Account) ProtoReflect() protoreflect.Message {
	mi := &file_npool_account_gw_v1_platform_platform_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Account.ProtoReflect.Descriptor instead.
func (*Account) Descriptor() ([]byte, []int) {
	return file_npool_account_gw_v1_platform_platform_proto_rawDescGZIP(), []int{0}
}

func (x *Account) GetID() uint32 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *Account) GetEntID() string {
	if x != nil {
		return x.EntID
	}
	return ""
}

func (x *Account) GetCoinTypeID() string {
	if x != nil {
		return x.CoinTypeID
	}
	return ""
}

func (x *Account) GetCoinName() string {
	if x != nil {
		return x.CoinName
	}
	return ""
}

func (x *Account) GetCoinUnit() string {
	if x != nil {
		return x.CoinUnit
	}
	return ""
}

func (x *Account) GetCoinEnv() string {
	if x != nil {
		return x.CoinEnv
	}
	return ""
}

func (x *Account) GetCoinLogo() string {
	if x != nil {
		return x.CoinLogo
	}
	return ""
}

func (x *Account) GetUsedFor() v1.AccountUsedFor {
	if x != nil {
		return x.UsedFor
	}
	return v1.AccountUsedFor(0)
}

func (x *Account) GetAccountID() string {
	if x != nil {
		return x.AccountID
	}
	return ""
}

func (x *Account) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *Account) GetBackup() bool {
	if x != nil {
		return x.Backup
	}
	return false
}

func (x *Account) GetActive() bool {
	if x != nil {
		return x.Active
	}
	return false
}

func (x *Account) GetLocked() bool {
	if x != nil {
		return x.Locked
	}
	return false
}

func (x *Account) GetLockedBy() v1.AccountLockedBy {
	if x != nil {
		return x.LockedBy
	}
	return v1.AccountLockedBy(0)
}

func (x *Account) GetBlocked() bool {
	if x != nil {
		return x.Blocked
	}
	return false
}

func (x *Account) GetCreatedAt() uint32 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *Account) GetUpdatedAt() uint32 {
	if x != nil {
		return x.UpdatedAt
	}
	return 0
}

type CreateAccountRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CoinTypeID string            `protobuf:"bytes,10,opt,name=CoinTypeID,proto3" json:"CoinTypeID,omitempty"`
	Address    *string           `protobuf:"bytes,20,opt,name=Address,proto3,oneof" json:"Address,omitempty"`
	UsedFor    v1.AccountUsedFor `protobuf:"varint,30,opt,name=UsedFor,proto3,enum=basetypes.v1.AccountUsedFor" json:"UsedFor,omitempty"`
}

func (x *CreateAccountRequest) Reset() {
	*x = CreateAccountRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_account_gw_v1_platform_platform_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateAccountRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAccountRequest) ProtoMessage() {}

func (x *CreateAccountRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_account_gw_v1_platform_platform_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAccountRequest.ProtoReflect.Descriptor instead.
func (*CreateAccountRequest) Descriptor() ([]byte, []int) {
	return file_npool_account_gw_v1_platform_platform_proto_rawDescGZIP(), []int{1}
}

func (x *CreateAccountRequest) GetCoinTypeID() string {
	if x != nil {
		return x.CoinTypeID
	}
	return ""
}

func (x *CreateAccountRequest) GetAddress() string {
	if x != nil && x.Address != nil {
		return *x.Address
	}
	return ""
}

func (x *CreateAccountRequest) GetUsedFor() v1.AccountUsedFor {
	if x != nil {
		return x.UsedFor
	}
	return v1.AccountUsedFor(0)
}

type CreateAccountResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *Account `protobuf:"bytes,10,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (x *CreateAccountResponse) Reset() {
	*x = CreateAccountResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_account_gw_v1_platform_platform_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateAccountResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAccountResponse) ProtoMessage() {}

func (x *CreateAccountResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_account_gw_v1_platform_platform_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAccountResponse.ProtoReflect.Descriptor instead.
func (*CreateAccountResponse) Descriptor() ([]byte, []int) {
	return file_npool_account_gw_v1_platform_platform_proto_rawDescGZIP(), []int{2}
}

func (x *CreateAccountResponse) GetInfo() *Account {
	if x != nil {
		return x.Info
	}
	return nil
}

type UpdateAccountRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID      uint32 `protobuf:"varint,10,opt,name=ID,proto3" json:"ID,omitempty"`
	EntID   string `protobuf:"bytes,11,opt,name=EntID,proto3" json:"EntID,omitempty"`
	Backup  *bool  `protobuf:"varint,20,opt,name=Backup,proto3,oneof" json:"Backup,omitempty"`
	Active  *bool  `protobuf:"varint,30,opt,name=Active,proto3,oneof" json:"Active,omitempty"`
	Blocked *bool  `protobuf:"varint,40,opt,name=Blocked,proto3,oneof" json:"Blocked,omitempty"`
	Locked  *bool  `protobuf:"varint,50,opt,name=Locked,proto3,oneof" json:"Locked,omitempty"`
}

func (x *UpdateAccountRequest) Reset() {
	*x = UpdateAccountRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_account_gw_v1_platform_platform_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateAccountRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateAccountRequest) ProtoMessage() {}

func (x *UpdateAccountRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_account_gw_v1_platform_platform_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateAccountRequest.ProtoReflect.Descriptor instead.
func (*UpdateAccountRequest) Descriptor() ([]byte, []int) {
	return file_npool_account_gw_v1_platform_platform_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateAccountRequest) GetID() uint32 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *UpdateAccountRequest) GetEntID() string {
	if x != nil {
		return x.EntID
	}
	return ""
}

func (x *UpdateAccountRequest) GetBackup() bool {
	if x != nil && x.Backup != nil {
		return *x.Backup
	}
	return false
}

func (x *UpdateAccountRequest) GetActive() bool {
	if x != nil && x.Active != nil {
		return *x.Active
	}
	return false
}

func (x *UpdateAccountRequest) GetBlocked() bool {
	if x != nil && x.Blocked != nil {
		return *x.Blocked
	}
	return false
}

func (x *UpdateAccountRequest) GetLocked() bool {
	if x != nil && x.Locked != nil {
		return *x.Locked
	}
	return false
}

type UpdateAccountResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *Account `protobuf:"bytes,10,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (x *UpdateAccountResponse) Reset() {
	*x = UpdateAccountResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_account_gw_v1_platform_platform_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateAccountResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateAccountResponse) ProtoMessage() {}

func (x *UpdateAccountResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_account_gw_v1_platform_platform_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateAccountResponse.ProtoReflect.Descriptor instead.
func (*UpdateAccountResponse) Descriptor() ([]byte, []int) {
	return file_npool_account_gw_v1_platform_platform_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateAccountResponse) GetInfo() *Account {
	if x != nil {
		return x.Info
	}
	return nil
}

type GetAccountsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Offset int32 `protobuf:"varint,20,opt,name=Offset,proto3" json:"Offset,omitempty"`
	Limit  int32 `protobuf:"varint,30,opt,name=Limit,proto3" json:"Limit,omitempty"`
}

func (x *GetAccountsRequest) Reset() {
	*x = GetAccountsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_account_gw_v1_platform_platform_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAccountsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAccountsRequest) ProtoMessage() {}

func (x *GetAccountsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_account_gw_v1_platform_platform_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAccountsRequest.ProtoReflect.Descriptor instead.
func (*GetAccountsRequest) Descriptor() ([]byte, []int) {
	return file_npool_account_gw_v1_platform_platform_proto_rawDescGZIP(), []int{5}
}

func (x *GetAccountsRequest) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *GetAccountsRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type GetAccountsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Infos []*Account `protobuf:"bytes,10,rep,name=Infos,proto3" json:"Infos,omitempty"`
	Total uint32     `protobuf:"varint,20,opt,name=Total,proto3" json:"Total,omitempty"`
}

func (x *GetAccountsResponse) Reset() {
	*x = GetAccountsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_account_gw_v1_platform_platform_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAccountsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAccountsResponse) ProtoMessage() {}

func (x *GetAccountsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_account_gw_v1_platform_platform_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAccountsResponse.ProtoReflect.Descriptor instead.
func (*GetAccountsResponse) Descriptor() ([]byte, []int) {
	return file_npool_account_gw_v1_platform_platform_proto_rawDescGZIP(), []int{6}
}

func (x *GetAccountsResponse) GetInfos() []*Account {
	if x != nil {
		return x.Infos
	}
	return nil
}

func (x *GetAccountsResponse) GetTotal() uint32 {
	if x != nil {
		return x.Total
	}
	return 0
}

var File_npool_account_gw_v1_platform_platform_proto protoreflect.FileDescriptor

var file_npool_account_gw_v1_platform_platform_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2f,
	0x67, 0x77, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x70,
	0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1b, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x70,
	0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x27, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f,
	0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x75, 0x73, 0x65, 0x64, 0x66, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x28, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70,
	0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x6c, 0x6f, 0x63,
	0x6b, 0x65, 0x64, 0x62, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8a, 0x04, 0x0a, 0x07,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x02, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x45, 0x6e, 0x74, 0x49, 0x44,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x45, 0x6e, 0x74, 0x49, 0x44, 0x12, 0x1e, 0x0a,
	0x0a, 0x43, 0x6f, 0x69, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x49, 0x44, 0x18, 0x14, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x43, 0x6f, 0x69, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x49, 0x44, 0x12, 0x1a, 0x0a,
	0x08, 0x43, 0x6f, 0x69, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x43, 0x6f, 0x69, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x43, 0x6f, 0x69,
	0x6e, 0x55, 0x6e, 0x69, 0x74, 0x18, 0x28, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x43, 0x6f, 0x69,
	0x6e, 0x55, 0x6e, 0x69, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x43, 0x6f, 0x69, 0x6e, 0x45, 0x6e, 0x76,
	0x18, 0x32, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x43, 0x6f, 0x69, 0x6e, 0x45, 0x6e, 0x76, 0x12,
	0x1a, 0x0a, 0x08, 0x43, 0x6f, 0x69, 0x6e, 0x4c, 0x6f, 0x67, 0x6f, 0x18, 0x3c, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x43, 0x6f, 0x69, 0x6e, 0x4c, 0x6f, 0x67, 0x6f, 0x12, 0x36, 0x0a, 0x07, 0x55,
	0x73, 0x65, 0x64, 0x46, 0x6f, 0x72, 0x18, 0x46, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1c, 0x2e, 0x62,
	0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x55, 0x73, 0x65, 0x64, 0x46, 0x6f, 0x72, 0x52, 0x07, 0x55, 0x73, 0x65, 0x64,
	0x46, 0x6f, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x44,
	0x18, 0x50, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49,
	0x44, 0x12, 0x18, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x5a, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x42,
	0x61, 0x63, 0x6b, 0x75, 0x70, 0x18, 0x64, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x42, 0x61, 0x63,
	0x6b, 0x75, 0x70, 0x12, 0x16, 0x0a, 0x06, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x18, 0x6e, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x06, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x4c,
	0x6f, 0x63, 0x6b, 0x65, 0x64, 0x18, 0x78, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x4c, 0x6f, 0x63,
	0x6b, 0x65, 0x64, 0x12, 0x3a, 0x0a, 0x08, 0x4c, 0x6f, 0x63, 0x6b, 0x65, 0x64, 0x42, 0x79, 0x18,
	0x82, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1d, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70,
	0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4c, 0x6f, 0x63,
	0x6b, 0x65, 0x64, 0x42, 0x79, 0x52, 0x08, 0x4c, 0x6f, 0x63, 0x6b, 0x65, 0x64, 0x42, 0x79, 0x12,
	0x19, 0x0a, 0x07, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x64, 0x18, 0x8c, 0x01, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x07, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x64, 0x12, 0x1d, 0x0a, 0x09, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x96, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x09, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0xa0, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x99, 0x01, 0x0a, 0x14, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x43, 0x6f, 0x69, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x49, 0x44, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x43, 0x6f, 0x69, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x49,
	0x44, 0x12, 0x1d, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x14, 0x20, 0x01,
	0x28, 0x09, 0x48, 0x00, 0x52, 0x07, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x88, 0x01, 0x01,
	0x12, 0x36, 0x0a, 0x07, 0x55, 0x73, 0x65, 0x64, 0x46, 0x6f, 0x72, 0x18, 0x1e, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x1c, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x55, 0x73, 0x65, 0x64, 0x46, 0x6f, 0x72, 0x52,
	0x07, 0x55, 0x73, 0x65, 0x64, 0x46, 0x6f, 0x72, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x41, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x22, 0x51, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x38, 0x0a,
	0x04, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x70, 0x6c,
	0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x52, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0xdf, 0x01, 0x0a, 0x14, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x49, 0x44,
	0x12, 0x14, 0x0a, 0x05, 0x45, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x45, 0x6e, 0x74, 0x49, 0x44, 0x12, 0x1b, 0x0a, 0x06, 0x42, 0x61, 0x63, 0x6b, 0x75, 0x70,
	0x18, 0x14, 0x20, 0x01, 0x28, 0x08, 0x48, 0x00, 0x52, 0x06, 0x42, 0x61, 0x63, 0x6b, 0x75, 0x70,
	0x88, 0x01, 0x01, 0x12, 0x1b, 0x0a, 0x06, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x18, 0x1e, 0x20,
	0x01, 0x28, 0x08, 0x48, 0x01, 0x52, 0x06, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x88, 0x01, 0x01,
	0x12, 0x1d, 0x0a, 0x07, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x64, 0x18, 0x28, 0x20, 0x01, 0x28,
	0x08, 0x48, 0x02, 0x52, 0x07, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x64, 0x88, 0x01, 0x01, 0x12,
	0x1b, 0x0a, 0x06, 0x4c, 0x6f, 0x63, 0x6b, 0x65, 0x64, 0x18, 0x32, 0x20, 0x01, 0x28, 0x08, 0x48,
	0x03, 0x52, 0x06, 0x4c, 0x6f, 0x63, 0x6b, 0x65, 0x64, 0x88, 0x01, 0x01, 0x42, 0x09, 0x0a, 0x07,
	0x5f, 0x42, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x41, 0x63, 0x74, 0x69,
	0x76, 0x65, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x64, 0x42, 0x09,
	0x0a, 0x07, 0x5f, 0x4c, 0x6f, 0x63, 0x6b, 0x65, 0x64, 0x22, 0x51, 0x0a, 0x15, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x38, 0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x24, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77,
	0x61, 0x79, 0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x41,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x42, 0x0a, 0x12,
	0x47, 0x65, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x14, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x06, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x4c, 0x69,
	0x6d, 0x69, 0x74, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x4c, 0x69, 0x6d, 0x69, 0x74,
	0x22, 0x67, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3a, 0x0a, 0x05, 0x49, 0x6e, 0x66, 0x6f, 0x73,
	0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72,
	0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x05, 0x49, 0x6e,
	0x66, 0x6f, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x14, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x32, 0xe1, 0x03, 0x0a, 0x07, 0x47, 0x61,
	0x74, 0x65, 0x77, 0x61, 0x79, 0x12, 0x9d, 0x01, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x31, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f,
	0x72, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x32, 0x2e, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x70, 0x6c, 0x61,
	0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x25,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1f, 0x3a, 0x01, 0x2a, 0x22, 0x1a, 0x2f, 0x76, 0x31, 0x2f, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x2f, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x9d, 0x01, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x31, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f,
	0x72, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x32, 0x2e, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x70, 0x6c, 0x61,
	0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x25,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1f, 0x3a, 0x01, 0x2a, 0x22, 0x1a, 0x2f, 0x76, 0x31, 0x2f, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x95, 0x01, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x41, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x73, 0x12, 0x2f, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e,
	0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d,
	0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x30, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72,
	0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x23, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1d,
	0x3a, 0x01, 0x2a, 0x22, 0x18, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x65, 0x74, 0x2f, 0x70, 0x6c, 0x61,
	0x74, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x42, 0x3f, 0x5a,
	0x3d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4e, 0x70, 0x6f, 0x6f,
	0x6c, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x2f, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2f,
	0x67, 0x77, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_npool_account_gw_v1_platform_platform_proto_rawDescOnce sync.Once
	file_npool_account_gw_v1_platform_platform_proto_rawDescData = file_npool_account_gw_v1_platform_platform_proto_rawDesc
)

func file_npool_account_gw_v1_platform_platform_proto_rawDescGZIP() []byte {
	file_npool_account_gw_v1_platform_platform_proto_rawDescOnce.Do(func() {
		file_npool_account_gw_v1_platform_platform_proto_rawDescData = protoimpl.X.CompressGZIP(file_npool_account_gw_v1_platform_platform_proto_rawDescData)
	})
	return file_npool_account_gw_v1_platform_platform_proto_rawDescData
}

var file_npool_account_gw_v1_platform_platform_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_npool_account_gw_v1_platform_platform_proto_goTypes = []interface{}{
	(*Account)(nil),               // 0: account.gateway.platform.v1.Account
	(*CreateAccountRequest)(nil),  // 1: account.gateway.platform.v1.CreateAccountRequest
	(*CreateAccountResponse)(nil), // 2: account.gateway.platform.v1.CreateAccountResponse
	(*UpdateAccountRequest)(nil),  // 3: account.gateway.platform.v1.UpdateAccountRequest
	(*UpdateAccountResponse)(nil), // 4: account.gateway.platform.v1.UpdateAccountResponse
	(*GetAccountsRequest)(nil),    // 5: account.gateway.platform.v1.GetAccountsRequest
	(*GetAccountsResponse)(nil),   // 6: account.gateway.platform.v1.GetAccountsResponse
	(v1.AccountUsedFor)(0),        // 7: basetypes.v1.AccountUsedFor
	(v1.AccountLockedBy)(0),       // 8: basetypes.v1.AccountLockedBy
}
var file_npool_account_gw_v1_platform_platform_proto_depIdxs = []int32{
	7, // 0: account.gateway.platform.v1.Account.UsedFor:type_name -> basetypes.v1.AccountUsedFor
	8, // 1: account.gateway.platform.v1.Account.LockedBy:type_name -> basetypes.v1.AccountLockedBy
	7, // 2: account.gateway.platform.v1.CreateAccountRequest.UsedFor:type_name -> basetypes.v1.AccountUsedFor
	0, // 3: account.gateway.platform.v1.CreateAccountResponse.Info:type_name -> account.gateway.platform.v1.Account
	0, // 4: account.gateway.platform.v1.UpdateAccountResponse.Info:type_name -> account.gateway.platform.v1.Account
	0, // 5: account.gateway.platform.v1.GetAccountsResponse.Infos:type_name -> account.gateway.platform.v1.Account
	1, // 6: account.gateway.platform.v1.Gateway.CreateAccount:input_type -> account.gateway.platform.v1.CreateAccountRequest
	3, // 7: account.gateway.platform.v1.Gateway.UpdateAccount:input_type -> account.gateway.platform.v1.UpdateAccountRequest
	5, // 8: account.gateway.platform.v1.Gateway.GetAccounts:input_type -> account.gateway.platform.v1.GetAccountsRequest
	2, // 9: account.gateway.platform.v1.Gateway.CreateAccount:output_type -> account.gateway.platform.v1.CreateAccountResponse
	4, // 10: account.gateway.platform.v1.Gateway.UpdateAccount:output_type -> account.gateway.platform.v1.UpdateAccountResponse
	6, // 11: account.gateway.platform.v1.Gateway.GetAccounts:output_type -> account.gateway.platform.v1.GetAccountsResponse
	9, // [9:12] is the sub-list for method output_type
	6, // [6:9] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_npool_account_gw_v1_platform_platform_proto_init() }
func file_npool_account_gw_v1_platform_platform_proto_init() {
	if File_npool_account_gw_v1_platform_platform_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_npool_account_gw_v1_platform_platform_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Account); i {
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
		file_npool_account_gw_v1_platform_platform_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateAccountRequest); i {
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
		file_npool_account_gw_v1_platform_platform_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateAccountResponse); i {
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
		file_npool_account_gw_v1_platform_platform_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateAccountRequest); i {
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
		file_npool_account_gw_v1_platform_platform_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateAccountResponse); i {
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
		file_npool_account_gw_v1_platform_platform_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAccountsRequest); i {
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
		file_npool_account_gw_v1_platform_platform_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAccountsResponse); i {
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
	file_npool_account_gw_v1_platform_platform_proto_msgTypes[1].OneofWrappers = []interface{}{}
	file_npool_account_gw_v1_platform_platform_proto_msgTypes[3].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_npool_account_gw_v1_platform_platform_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_npool_account_gw_v1_platform_platform_proto_goTypes,
		DependencyIndexes: file_npool_account_gw_v1_platform_platform_proto_depIdxs,
		MessageInfos:      file_npool_account_gw_v1_platform_platform_proto_msgTypes,
	}.Build()
	File_npool_account_gw_v1_platform_platform_proto = out.File
	file_npool_account_gw_v1_platform_platform_proto_rawDesc = nil
	file_npool_account_gw_v1_platform_platform_proto_goTypes = nil
	file_npool_account_gw_v1_platform_platform_proto_depIdxs = nil
}
