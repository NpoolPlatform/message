// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v3.18.1
// source: npool/account/gw/v1/orderbenefit/orderbenefit.proto

package orderbenefit

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

	ID               uint32            `protobuf:"varint,9,opt,name=ID,proto3" json:"ID,omitempty"`
	EntID            string            `protobuf:"bytes,10,opt,name=EntID,proto3" json:"EntID,omitempty"`
	AppID            string            `protobuf:"bytes,20,opt,name=AppID,proto3" json:"AppID,omitempty"`
	UserID           string            `protobuf:"bytes,30,opt,name=UserID,proto3" json:"UserID,omitempty"`
	CoinTypeID       string            `protobuf:"bytes,40,opt,name=CoinTypeID,proto3" json:"CoinTypeID,omitempty"`
	CoinName         string            `protobuf:"bytes,50,opt,name=CoinName,proto3" json:"CoinName,omitempty"`
	CoinDisplayNames []string          `protobuf:"bytes,60,rep,name=CoinDisplayNames,proto3" json:"CoinDisplayNames,omitempty"`
	CoinUnit         string            `protobuf:"bytes,70,opt,name=CoinUnit,proto3" json:"CoinUnit,omitempty"`
	CoinEnv          string            `protobuf:"bytes,80,opt,name=CoinEnv,proto3" json:"CoinEnv,omitempty"`
	CoinLogo         string            `protobuf:"bytes,90,opt,name=CoinLogo,proto3" json:"CoinLogo,omitempty"`
	AccountID        string            `protobuf:"bytes,100,opt,name=AccountID,proto3" json:"AccountID,omitempty"`
	Address          string            `protobuf:"bytes,110,opt,name=Address,proto3" json:"Address,omitempty"`
	PhoneNO          string            `protobuf:"bytes,120,opt,name=PhoneNO,proto3" json:"PhoneNO,omitempty"`
	EmailAddress     string            `protobuf:"bytes,130,opt,name=EmailAddress,proto3" json:"EmailAddress,omitempty"`
	OrderID          string            `protobuf:"bytes,140,opt,name=OrderID,proto3" json:"OrderID,omitempty"`
	Active           bool              `protobuf:"varint,150,opt,name=Active,proto3" json:"Active,omitempty"`
	Blocked          bool              `protobuf:"varint,160,opt,name=Blocked,proto3" json:"Blocked,omitempty"`
	Locked           bool              `protobuf:"varint,170,opt,name=Locked,proto3" json:"Locked,omitempty"`
	UsedFor          v1.AccountUsedFor `protobuf:"varint,180,opt,name=UsedFor,proto3,enum=basetypes.v1.AccountUsedFor" json:"UsedFor,omitempty"`
	CreatedAt        uint32            `protobuf:"varint,1000,opt,name=CreatedAt,proto3" json:"CreatedAt,omitempty"`
	UpdatedAt        uint32            `protobuf:"varint,1010,opt,name=UpdatedAt,proto3" json:"UpdatedAt,omitempty"`
}

func (x *Account) Reset() {
	*x = Account{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_account_gw_v1_orderbenefit_orderbenefit_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Account) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Account) ProtoMessage() {}

func (x *Account) ProtoReflect() protoreflect.Message {
	mi := &file_npool_account_gw_v1_orderbenefit_orderbenefit_proto_msgTypes[0]
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
	return file_npool_account_gw_v1_orderbenefit_orderbenefit_proto_rawDescGZIP(), []int{0}
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

func (x *Account) GetAppID() string {
	if x != nil {
		return x.AppID
	}
	return ""
}

func (x *Account) GetUserID() string {
	if x != nil {
		return x.UserID
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

func (x *Account) GetCoinDisplayNames() []string {
	if x != nil {
		return x.CoinDisplayNames
	}
	return nil
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

func (x *Account) GetPhoneNO() string {
	if x != nil {
		return x.PhoneNO
	}
	return ""
}

func (x *Account) GetEmailAddress() string {
	if x != nil {
		return x.EmailAddress
	}
	return ""
}

func (x *Account) GetOrderID() string {
	if x != nil {
		return x.OrderID
	}
	return ""
}

func (x *Account) GetActive() bool {
	if x != nil {
		return x.Active
	}
	return false
}

func (x *Account) GetBlocked() bool {
	if x != nil {
		return x.Blocked
	}
	return false
}

func (x *Account) GetLocked() bool {
	if x != nil {
		return x.Locked
	}
	return false
}

func (x *Account) GetUsedFor() v1.AccountUsedFor {
	if x != nil {
		return x.UsedFor
	}
	return v1.AccountUsedFor(0)
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

type GetAccountRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EntID  string `protobuf:"bytes,10,opt,name=EntID,proto3" json:"EntID,omitempty"`
	AppID  string `protobuf:"bytes,20,opt,name=AppID,proto3" json:"AppID,omitempty"`
	UserID string `protobuf:"bytes,30,opt,name=UserID,proto3" json:"UserID,omitempty"`
}

func (x *GetAccountRequest) Reset() {
	*x = GetAccountRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_account_gw_v1_orderbenefit_orderbenefit_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAccountRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAccountRequest) ProtoMessage() {}

func (x *GetAccountRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_account_gw_v1_orderbenefit_orderbenefit_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAccountRequest.ProtoReflect.Descriptor instead.
func (*GetAccountRequest) Descriptor() ([]byte, []int) {
	return file_npool_account_gw_v1_orderbenefit_orderbenefit_proto_rawDescGZIP(), []int{1}
}

func (x *GetAccountRequest) GetEntID() string {
	if x != nil {
		return x.EntID
	}
	return ""
}

func (x *GetAccountRequest) GetAppID() string {
	if x != nil {
		return x.AppID
	}
	return ""
}

func (x *GetAccountRequest) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

type GetAccountResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *Account `protobuf:"bytes,10,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (x *GetAccountResponse) Reset() {
	*x = GetAccountResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_account_gw_v1_orderbenefit_orderbenefit_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAccountResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAccountResponse) ProtoMessage() {}

func (x *GetAccountResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_account_gw_v1_orderbenefit_orderbenefit_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAccountResponse.ProtoReflect.Descriptor instead.
func (*GetAccountResponse) Descriptor() ([]byte, []int) {
	return file_npool_account_gw_v1_orderbenefit_orderbenefit_proto_rawDescGZIP(), []int{2}
}

func (x *GetAccountResponse) GetInfo() *Account {
	if x != nil {
		return x.Info
	}
	return nil
}

type GetAccountsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AppID  string `protobuf:"bytes,10,opt,name=AppID,proto3" json:"AppID,omitempty"`
	UserID string `protobuf:"bytes,20,opt,name=UserID,proto3" json:"UserID,omitempty"`
	Offset int32  `protobuf:"varint,30,opt,name=Offset,proto3" json:"Offset,omitempty"`
	Limit  int32  `protobuf:"varint,40,opt,name=Limit,proto3" json:"Limit,omitempty"`
}

func (x *GetAccountsRequest) Reset() {
	*x = GetAccountsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_account_gw_v1_orderbenefit_orderbenefit_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAccountsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAccountsRequest) ProtoMessage() {}

func (x *GetAccountsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_account_gw_v1_orderbenefit_orderbenefit_proto_msgTypes[3]
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
	return file_npool_account_gw_v1_orderbenefit_orderbenefit_proto_rawDescGZIP(), []int{3}
}

func (x *GetAccountsRequest) GetAppID() string {
	if x != nil {
		return x.AppID
	}
	return ""
}

func (x *GetAccountsRequest) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
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
		mi := &file_npool_account_gw_v1_orderbenefit_orderbenefit_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAccountsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAccountsResponse) ProtoMessage() {}

func (x *GetAccountsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_account_gw_v1_orderbenefit_orderbenefit_proto_msgTypes[4]
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
	return file_npool_account_gw_v1_orderbenefit_orderbenefit_proto_rawDescGZIP(), []int{4}
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

var File_npool_account_gw_v1_orderbenefit_orderbenefit_proto protoreflect.FileDescriptor

var file_npool_account_gw_v1_orderbenefit_orderbenefit_proto_rawDesc = []byte{
	0x0a, 0x33, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2f,
	0x67, 0x77, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x62, 0x65, 0x6e, 0x65, 0x66,
	0x69, 0x74, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x62, 0x65, 0x6e, 0x65, 0x66, 0x69, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x67,
	0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x62, 0x65, 0x6e, 0x65,
	0x66, 0x69, 0x74, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x27, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x62, 0x61, 0x73, 0x65,
	0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x75, 0x73, 0x65, 0x64, 0x66, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xed, 0x04,
	0x0a, 0x07, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x45, 0x6e, 0x74,
	0x49, 0x44, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x45, 0x6e, 0x74, 0x49, 0x44, 0x12,
	0x14, 0x0a, 0x05, 0x41, 0x70, 0x70, 0x49, 0x44, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x41, 0x70, 0x70, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18,
	0x1e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x1e, 0x0a,
	0x0a, 0x43, 0x6f, 0x69, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x49, 0x44, 0x18, 0x28, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x43, 0x6f, 0x69, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x49, 0x44, 0x12, 0x1a, 0x0a,
	0x08, 0x43, 0x6f, 0x69, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x32, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x43, 0x6f, 0x69, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x2a, 0x0a, 0x10, 0x43, 0x6f, 0x69,
	0x6e, 0x44, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x18, 0x3c, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x10, 0x43, 0x6f, 0x69, 0x6e, 0x44, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79,
	0x4e, 0x61, 0x6d, 0x65, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x43, 0x6f, 0x69, 0x6e, 0x55, 0x6e, 0x69,
	0x74, 0x18, 0x46, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x43, 0x6f, 0x69, 0x6e, 0x55, 0x6e, 0x69,
	0x74, 0x12, 0x18, 0x0a, 0x07, 0x43, 0x6f, 0x69, 0x6e, 0x45, 0x6e, 0x76, 0x18, 0x50, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x43, 0x6f, 0x69, 0x6e, 0x45, 0x6e, 0x76, 0x12, 0x1a, 0x0a, 0x08, 0x43,
	0x6f, 0x69, 0x6e, 0x4c, 0x6f, 0x67, 0x6f, 0x18, 0x5a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x43,
	0x6f, 0x69, 0x6e, 0x4c, 0x6f, 0x67, 0x6f, 0x12, 0x1c, 0x0a, 0x09, 0x41, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x49, 0x44, 0x18, 0x64, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x49, 0x44, 0x12, 0x18, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x18, 0x6e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12,
	0x18, 0x0a, 0x07, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x4f, 0x18, 0x78, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x4f, 0x12, 0x23, 0x0a, 0x0c, 0x45, 0x6d, 0x61,
	0x69, 0x6c, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x82, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x19,
	0x0a, 0x07, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x44, 0x18, 0x8c, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x44, 0x12, 0x17, 0x0a, 0x06, 0x41, 0x63, 0x74,
	0x69, 0x76, 0x65, 0x18, 0x96, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x41, 0x63, 0x74, 0x69,
	0x76, 0x65, 0x12, 0x19, 0x0a, 0x07, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x64, 0x18, 0xa0, 0x01,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x64, 0x12, 0x17, 0x0a,
	0x06, 0x4c, 0x6f, 0x63, 0x6b, 0x65, 0x64, 0x18, 0xaa, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06,
	0x4c, 0x6f, 0x63, 0x6b, 0x65, 0x64, 0x12, 0x37, 0x0a, 0x07, 0x55, 0x73, 0x65, 0x64, 0x46, 0x6f,
	0x72, 0x18, 0xb4, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1c, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x55,
	0x73, 0x65, 0x64, 0x46, 0x6f, 0x72, 0x52, 0x07, 0x55, 0x73, 0x65, 0x64, 0x46, 0x6f, 0x72, 0x12,
	0x1d, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0xe8, 0x07, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d,
	0x0a, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0xf2, 0x07, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x57, 0x0a,
	0x11, 0x47, 0x65, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x45, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x45, 0x6e, 0x74, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x41, 0x70, 0x70, 0x49,
	0x44, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x41, 0x70, 0x70, 0x49, 0x44, 0x12, 0x16,
	0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x22, 0x52, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x41, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3c, 0x0a, 0x04,
	0x49, 0x6e, 0x66, 0x6f, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x62, 0x65, 0x6e, 0x65, 0x66, 0x69, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x52, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x70, 0x0a, 0x12, 0x47, 0x65,
	0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x14, 0x0a, 0x05, 0x41, 0x70, 0x70, 0x49, 0x44, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x41, 0x70, 0x70, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44,
	0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x16,
	0x0a, 0x06, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06,
	0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x18,
	0x28, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x22, 0x6b, 0x0a, 0x13,
	0x47, 0x65, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x3e, 0x0a, 0x05, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x18, 0x0a, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x28, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x67, 0x61, 0x74,
	0x65, 0x77, 0x61, 0x79, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x62, 0x65, 0x6e, 0x65, 0x66, 0x69,
	0x74, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x05, 0x49, 0x6e,
	0x66, 0x6f, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x14, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x32, 0xcd, 0x02, 0x0a, 0x07, 0x47, 0x61,
	0x74, 0x65, 0x77, 0x61, 0x79, 0x12, 0x9d, 0x01, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x41, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x12, 0x32, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x67,
	0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x62, 0x65, 0x6e, 0x65,
	0x66, 0x69, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x33, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x62, 0x65, 0x6e, 0x65, 0x66, 0x69, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x26, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x20, 0x3a, 0x01, 0x2a, 0x22, 0x1b, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x65,
	0x74, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x62, 0x65, 0x6e, 0x65, 0x66, 0x69, 0x74, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0xa1, 0x01, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x41, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x73, 0x12, 0x33, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e,
	0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x62, 0x65, 0x6e,
	0x65, 0x66, 0x69, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x34, 0x2e, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x62, 0x65, 0x6e, 0x65, 0x66, 0x69, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x27, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x21, 0x3a, 0x01, 0x2a, 0x22, 0x1c, 0x2f, 0x76, 0x31,
	0x2f, 0x67, 0x65, 0x74, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x62, 0x65, 0x6e, 0x65, 0x66, 0x69,
	0x74, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x42, 0x43, 0x5a, 0x41, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4e, 0x70, 0x6f, 0x6f, 0x6c, 0x50, 0x6c, 0x61,
	0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x6e, 0x70,
	0x6f, 0x6f, 0x6c, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2f, 0x67, 0x77, 0x2f, 0x76,
	0x31, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x62, 0x65, 0x6e, 0x65, 0x66, 0x69, 0x74, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_npool_account_gw_v1_orderbenefit_orderbenefit_proto_rawDescOnce sync.Once
	file_npool_account_gw_v1_orderbenefit_orderbenefit_proto_rawDescData = file_npool_account_gw_v1_orderbenefit_orderbenefit_proto_rawDesc
)

func file_npool_account_gw_v1_orderbenefit_orderbenefit_proto_rawDescGZIP() []byte {
	file_npool_account_gw_v1_orderbenefit_orderbenefit_proto_rawDescOnce.Do(func() {
		file_npool_account_gw_v1_orderbenefit_orderbenefit_proto_rawDescData = protoimpl.X.CompressGZIP(file_npool_account_gw_v1_orderbenefit_orderbenefit_proto_rawDescData)
	})
	return file_npool_account_gw_v1_orderbenefit_orderbenefit_proto_rawDescData
}

var file_npool_account_gw_v1_orderbenefit_orderbenefit_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_npool_account_gw_v1_orderbenefit_orderbenefit_proto_goTypes = []interface{}{
	(*Account)(nil),             // 0: account.gateway.orderbenefit.v1.Account
	(*GetAccountRequest)(nil),   // 1: account.gateway.orderbenefit.v1.GetAccountRequest
	(*GetAccountResponse)(nil),  // 2: account.gateway.orderbenefit.v1.GetAccountResponse
	(*GetAccountsRequest)(nil),  // 3: account.gateway.orderbenefit.v1.GetAccountsRequest
	(*GetAccountsResponse)(nil), // 4: account.gateway.orderbenefit.v1.GetAccountsResponse
	(v1.AccountUsedFor)(0),      // 5: basetypes.v1.AccountUsedFor
}
var file_npool_account_gw_v1_orderbenefit_orderbenefit_proto_depIdxs = []int32{
	5, // 0: account.gateway.orderbenefit.v1.Account.UsedFor:type_name -> basetypes.v1.AccountUsedFor
	0, // 1: account.gateway.orderbenefit.v1.GetAccountResponse.Info:type_name -> account.gateway.orderbenefit.v1.Account
	0, // 2: account.gateway.orderbenefit.v1.GetAccountsResponse.Infos:type_name -> account.gateway.orderbenefit.v1.Account
	1, // 3: account.gateway.orderbenefit.v1.Gateway.GetAccount:input_type -> account.gateway.orderbenefit.v1.GetAccountRequest
	3, // 4: account.gateway.orderbenefit.v1.Gateway.GetAccounts:input_type -> account.gateway.orderbenefit.v1.GetAccountsRequest
	2, // 5: account.gateway.orderbenefit.v1.Gateway.GetAccount:output_type -> account.gateway.orderbenefit.v1.GetAccountResponse
	4, // 6: account.gateway.orderbenefit.v1.Gateway.GetAccounts:output_type -> account.gateway.orderbenefit.v1.GetAccountsResponse
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_npool_account_gw_v1_orderbenefit_orderbenefit_proto_init() }
func file_npool_account_gw_v1_orderbenefit_orderbenefit_proto_init() {
	if File_npool_account_gw_v1_orderbenefit_orderbenefit_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_npool_account_gw_v1_orderbenefit_orderbenefit_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_npool_account_gw_v1_orderbenefit_orderbenefit_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAccountRequest); i {
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
		file_npool_account_gw_v1_orderbenefit_orderbenefit_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAccountResponse); i {
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
		file_npool_account_gw_v1_orderbenefit_orderbenefit_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
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
		file_npool_account_gw_v1_orderbenefit_orderbenefit_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_npool_account_gw_v1_orderbenefit_orderbenefit_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_npool_account_gw_v1_orderbenefit_orderbenefit_proto_goTypes,
		DependencyIndexes: file_npool_account_gw_v1_orderbenefit_orderbenefit_proto_depIdxs,
		MessageInfos:      file_npool_account_gw_v1_orderbenefit_orderbenefit_proto_msgTypes,
	}.Build()
	File_npool_account_gw_v1_orderbenefit_orderbenefit_proto = out.File
	file_npool_account_gw_v1_orderbenefit_orderbenefit_proto_rawDesc = nil
	file_npool_account_gw_v1_orderbenefit_orderbenefit_proto_goTypes = nil
	file_npool_account_gw_v1_orderbenefit_orderbenefit_proto_depIdxs = nil
}
