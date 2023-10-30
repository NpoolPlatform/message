// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.18.1
// source: npool/ledger/gw/v1/withdraw/withdraw.proto

package withdraw

import (
	v1 "github.com/NpoolPlatform/message/npool/basetypes/ledger/v1"
	v11 "github.com/NpoolPlatform/message/npool/basetypes/v1"
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

type Withdraw struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CoinTypeID    string           `protobuf:"bytes,10,opt,name=CoinTypeID,proto3" json:"CoinTypeID,omitempty"`
	CoinName      string           `protobuf:"bytes,20,opt,name=CoinName,proto3" json:"CoinName,omitempty"`
	DisplayNames  []string         `protobuf:"bytes,21,rep,name=DisplayNames,proto3" json:"DisplayNames,omitempty"`
	CoinLogo      string           `protobuf:"bytes,30,opt,name=CoinLogo,proto3" json:"CoinLogo,omitempty"`
	CoinUnit      string           `protobuf:"bytes,40,opt,name=CoinUnit,proto3" json:"CoinUnit,omitempty"`
	Amount        string           `protobuf:"bytes,50,opt,name=Amount,proto3" json:"Amount,omitempty"`
	CreatedAt     uint32           `protobuf:"varint,60,opt,name=CreatedAt,proto3" json:"CreatedAt,omitempty"`
	Address       string           `protobuf:"bytes,70,opt,name=Address,proto3" json:"Address,omitempty"`
	AddressLabels []string         `protobuf:"bytes,80,rep,name=AddressLabels,proto3" json:"AddressLabels,omitempty"`
	State         v1.WithdrawState `protobuf:"varint,90,opt,name=State,proto3,enum=basetypes.ledger.v1.WithdrawState" json:"State,omitempty"`
	Message       string           `protobuf:"bytes,100,opt,name=Message,proto3" json:"Message,omitempty"`
	AppID         string           `protobuf:"bytes,110,opt,name=AppID,proto3" json:"AppID,omitempty"`
	UserID        string           `protobuf:"bytes,120,opt,name=UserID,proto3" json:"UserID,omitempty"`
	ID            string           `protobuf:"bytes,130,opt,name=ID,proto3" json:"ID,omitempty"`
}

func (x *Withdraw) Reset() {
	*x = Withdraw{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_ledger_gw_v1_withdraw_withdraw_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Withdraw) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Withdraw) ProtoMessage() {}

func (x *Withdraw) ProtoReflect() protoreflect.Message {
	mi := &file_npool_ledger_gw_v1_withdraw_withdraw_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Withdraw.ProtoReflect.Descriptor instead.
func (*Withdraw) Descriptor() ([]byte, []int) {
	return file_npool_ledger_gw_v1_withdraw_withdraw_proto_rawDescGZIP(), []int{0}
}

func (x *Withdraw) GetCoinTypeID() string {
	if x != nil {
		return x.CoinTypeID
	}
	return ""
}

func (x *Withdraw) GetCoinName() string {
	if x != nil {
		return x.CoinName
	}
	return ""
}

func (x *Withdraw) GetDisplayNames() []string {
	if x != nil {
		return x.DisplayNames
	}
	return nil
}

func (x *Withdraw) GetCoinLogo() string {
	if x != nil {
		return x.CoinLogo
	}
	return ""
}

func (x *Withdraw) GetCoinUnit() string {
	if x != nil {
		return x.CoinUnit
	}
	return ""
}

func (x *Withdraw) GetAmount() string {
	if x != nil {
		return x.Amount
	}
	return ""
}

func (x *Withdraw) GetCreatedAt() uint32 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *Withdraw) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *Withdraw) GetAddressLabels() []string {
	if x != nil {
		return x.AddressLabels
	}
	return nil
}

func (x *Withdraw) GetState() v1.WithdrawState {
	if x != nil {
		return x.State
	}
	return v1.WithdrawState(0)
}

func (x *Withdraw) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *Withdraw) GetAppID() string {
	if x != nil {
		return x.AppID
	}
	return ""
}

func (x *Withdraw) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

func (x *Withdraw) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

type CreateWithdrawRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AppID            string         `protobuf:"bytes,10,opt,name=AppID,proto3" json:"AppID,omitempty"`
	UserID           string         `protobuf:"bytes,20,opt,name=UserID,proto3" json:"UserID,omitempty"`
	CoinTypeID       string         `protobuf:"bytes,30,opt,name=CoinTypeID,proto3" json:"CoinTypeID,omitempty"`
	AccountID        string         `protobuf:"bytes,40,opt,name=AccountID,proto3" json:"AccountID,omitempty"`
	Amount           string         `protobuf:"bytes,50,opt,name=Amount,proto3" json:"Amount,omitempty"`
	AccountType      v11.SignMethod `protobuf:"varint,60,opt,name=AccountType,proto3,enum=basetypes.v1.SignMethod" json:"AccountType,omitempty"`
	VerificationCode string         `protobuf:"bytes,80,opt,name=VerificationCode,proto3" json:"VerificationCode,omitempty"`
}

func (x *CreateWithdrawRequest) Reset() {
	*x = CreateWithdrawRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_ledger_gw_v1_withdraw_withdraw_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateWithdrawRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateWithdrawRequest) ProtoMessage() {}

func (x *CreateWithdrawRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_ledger_gw_v1_withdraw_withdraw_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateWithdrawRequest.ProtoReflect.Descriptor instead.
func (*CreateWithdrawRequest) Descriptor() ([]byte, []int) {
	return file_npool_ledger_gw_v1_withdraw_withdraw_proto_rawDescGZIP(), []int{1}
}

func (x *CreateWithdrawRequest) GetAppID() string {
	if x != nil {
		return x.AppID
	}
	return ""
}

func (x *CreateWithdrawRequest) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

func (x *CreateWithdrawRequest) GetCoinTypeID() string {
	if x != nil {
		return x.CoinTypeID
	}
	return ""
}

func (x *CreateWithdrawRequest) GetAccountID() string {
	if x != nil {
		return x.AccountID
	}
	return ""
}

func (x *CreateWithdrawRequest) GetAmount() string {
	if x != nil {
		return x.Amount
	}
	return ""
}

func (x *CreateWithdrawRequest) GetAccountType() v11.SignMethod {
	if x != nil {
		return x.AccountType
	}
	return v11.SignMethod(0)
}

func (x *CreateWithdrawRequest) GetVerificationCode() string {
	if x != nil {
		return x.VerificationCode
	}
	return ""
}

type CreateWithdrawResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *Withdraw `protobuf:"bytes,10,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (x *CreateWithdrawResponse) Reset() {
	*x = CreateWithdrawResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_ledger_gw_v1_withdraw_withdraw_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateWithdrawResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateWithdrawResponse) ProtoMessage() {}

func (x *CreateWithdrawResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_ledger_gw_v1_withdraw_withdraw_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateWithdrawResponse.ProtoReflect.Descriptor instead.
func (*CreateWithdrawResponse) Descriptor() ([]byte, []int) {
	return file_npool_ledger_gw_v1_withdraw_withdraw_proto_rawDescGZIP(), []int{2}
}

func (x *CreateWithdrawResponse) GetInfo() *Withdraw {
	if x != nil {
		return x.Info
	}
	return nil
}

type GetWithdrawsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AppID  string `protobuf:"bytes,10,opt,name=AppID,proto3" json:"AppID,omitempty"`
	UserID string `protobuf:"bytes,20,opt,name=UserID,proto3" json:"UserID,omitempty"`
	Offset int32  `protobuf:"varint,30,opt,name=Offset,proto3" json:"Offset,omitempty"`
	Limit  int32  `protobuf:"varint,40,opt,name=Limit,proto3" json:"Limit,omitempty"`
}

func (x *GetWithdrawsRequest) Reset() {
	*x = GetWithdrawsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_ledger_gw_v1_withdraw_withdraw_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetWithdrawsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetWithdrawsRequest) ProtoMessage() {}

func (x *GetWithdrawsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_ledger_gw_v1_withdraw_withdraw_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetWithdrawsRequest.ProtoReflect.Descriptor instead.
func (*GetWithdrawsRequest) Descriptor() ([]byte, []int) {
	return file_npool_ledger_gw_v1_withdraw_withdraw_proto_rawDescGZIP(), []int{3}
}

func (x *GetWithdrawsRequest) GetAppID() string {
	if x != nil {
		return x.AppID
	}
	return ""
}

func (x *GetWithdrawsRequest) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

func (x *GetWithdrawsRequest) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *GetWithdrawsRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type GetWithdrawsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Infos []*Withdraw `protobuf:"bytes,10,rep,name=Infos,proto3" json:"Infos,omitempty"`
	Total uint32      `protobuf:"varint,20,opt,name=Total,proto3" json:"Total,omitempty"`
}

func (x *GetWithdrawsResponse) Reset() {
	*x = GetWithdrawsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_ledger_gw_v1_withdraw_withdraw_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetWithdrawsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetWithdrawsResponse) ProtoMessage() {}

func (x *GetWithdrawsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_ledger_gw_v1_withdraw_withdraw_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetWithdrawsResponse.ProtoReflect.Descriptor instead.
func (*GetWithdrawsResponse) Descriptor() ([]byte, []int) {
	return file_npool_ledger_gw_v1_withdraw_withdraw_proto_rawDescGZIP(), []int{4}
}

func (x *GetWithdrawsResponse) GetInfos() []*Withdraw {
	if x != nil {
		return x.Infos
	}
	return nil
}

func (x *GetWithdrawsResponse) GetTotal() uint32 {
	if x != nil {
		return x.Total
	}
	return 0
}

type GetAppWithdrawsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AppID  string `protobuf:"bytes,10,opt,name=AppID,proto3" json:"AppID,omitempty"`
	Offset int32  `protobuf:"varint,20,opt,name=Offset,proto3" json:"Offset,omitempty"`
	Limit  int32  `protobuf:"varint,30,opt,name=Limit,proto3" json:"Limit,omitempty"`
}

func (x *GetAppWithdrawsRequest) Reset() {
	*x = GetAppWithdrawsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_ledger_gw_v1_withdraw_withdraw_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAppWithdrawsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAppWithdrawsRequest) ProtoMessage() {}

func (x *GetAppWithdrawsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_ledger_gw_v1_withdraw_withdraw_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAppWithdrawsRequest.ProtoReflect.Descriptor instead.
func (*GetAppWithdrawsRequest) Descriptor() ([]byte, []int) {
	return file_npool_ledger_gw_v1_withdraw_withdraw_proto_rawDescGZIP(), []int{5}
}

func (x *GetAppWithdrawsRequest) GetAppID() string {
	if x != nil {
		return x.AppID
	}
	return ""
}

func (x *GetAppWithdrawsRequest) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *GetAppWithdrawsRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type GetAppWithdrawsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Infos []*Withdraw `protobuf:"bytes,10,rep,name=Infos,proto3" json:"Infos,omitempty"`
	Total uint32      `protobuf:"varint,20,opt,name=Total,proto3" json:"Total,omitempty"`
}

func (x *GetAppWithdrawsResponse) Reset() {
	*x = GetAppWithdrawsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_ledger_gw_v1_withdraw_withdraw_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAppWithdrawsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAppWithdrawsResponse) ProtoMessage() {}

func (x *GetAppWithdrawsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_ledger_gw_v1_withdraw_withdraw_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAppWithdrawsResponse.ProtoReflect.Descriptor instead.
func (*GetAppWithdrawsResponse) Descriptor() ([]byte, []int) {
	return file_npool_ledger_gw_v1_withdraw_withdraw_proto_rawDescGZIP(), []int{6}
}

func (x *GetAppWithdrawsResponse) GetInfos() []*Withdraw {
	if x != nil {
		return x.Infos
	}
	return nil
}

func (x *GetAppWithdrawsResponse) GetTotal() uint32 {
	if x != nil {
		return x.Total
	}
	return 0
}

type GetNAppWithdrawsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TargetAppID string `protobuf:"bytes,10,opt,name=TargetAppID,proto3" json:"TargetAppID,omitempty"`
	Offset      int32  `protobuf:"varint,20,opt,name=Offset,proto3" json:"Offset,omitempty"`
	Limit       int32  `protobuf:"varint,30,opt,name=Limit,proto3" json:"Limit,omitempty"`
}

func (x *GetNAppWithdrawsRequest) Reset() {
	*x = GetNAppWithdrawsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_ledger_gw_v1_withdraw_withdraw_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetNAppWithdrawsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetNAppWithdrawsRequest) ProtoMessage() {}

func (x *GetNAppWithdrawsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_ledger_gw_v1_withdraw_withdraw_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetNAppWithdrawsRequest.ProtoReflect.Descriptor instead.
func (*GetNAppWithdrawsRequest) Descriptor() ([]byte, []int) {
	return file_npool_ledger_gw_v1_withdraw_withdraw_proto_rawDescGZIP(), []int{7}
}

func (x *GetNAppWithdrawsRequest) GetTargetAppID() string {
	if x != nil {
		return x.TargetAppID
	}
	return ""
}

func (x *GetNAppWithdrawsRequest) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *GetNAppWithdrawsRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type GetNAppWithdrawsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Infos []*Withdraw `protobuf:"bytes,10,rep,name=Infos,proto3" json:"Infos,omitempty"`
	Total uint32      `protobuf:"varint,20,opt,name=Total,proto3" json:"Total,omitempty"`
}

func (x *GetNAppWithdrawsResponse) Reset() {
	*x = GetNAppWithdrawsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_ledger_gw_v1_withdraw_withdraw_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetNAppWithdrawsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetNAppWithdrawsResponse) ProtoMessage() {}

func (x *GetNAppWithdrawsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_ledger_gw_v1_withdraw_withdraw_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetNAppWithdrawsResponse.ProtoReflect.Descriptor instead.
func (*GetNAppWithdrawsResponse) Descriptor() ([]byte, []int) {
	return file_npool_ledger_gw_v1_withdraw_withdraw_proto_rawDescGZIP(), []int{8}
}

func (x *GetNAppWithdrawsResponse) GetInfos() []*Withdraw {
	if x != nil {
		return x.Infos
	}
	return nil
}

func (x *GetNAppWithdrawsResponse) GetTotal() uint32 {
	if x != nil {
		return x.Total
	}
	return 0
}

var File_npool_ledger_gw_v1_withdraw_withdraw_proto protoreflect.FileDescriptor

var file_npool_ledger_gw_v1_withdraw_withdraw_proto_rawDesc = []byte{
	0x0a, 0x2a, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2f, 0x67,
	0x77, 0x2f, 0x76, 0x31, 0x2f, 0x77, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x2f, 0x77, 0x69,
	0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1a, 0x6c, 0x65,
	0x64, 0x67, 0x65, 0x72, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x77, 0x69, 0x74,
	0x68, 0x64, 0x72, 0x61, 0x77, 0x2e, 0x76, 0x31, 0x1a, 0x23, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f,
	0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x69, 0x67,
	0x6e, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x25, 0x6e, 0x70, 0x6f,
	0x6f, 0x6c, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x6c, 0x65, 0x64,
	0x67, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xab, 0x03, 0x0a, 0x08, 0x57, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x12,
	0x1e, 0x0a, 0x0a, 0x43, 0x6f, 0x69, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x49, 0x44, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x43, 0x6f, 0x69, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x49, 0x44, 0x12,
	0x1a, 0x0a, 0x08, 0x43, 0x6f, 0x69, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x14, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x43, 0x6f, 0x69, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x44,
	0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x18, 0x15, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x0c, 0x44, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x12,
	0x1a, 0x0a, 0x08, 0x43, 0x6f, 0x69, 0x6e, 0x4c, 0x6f, 0x67, 0x6f, 0x18, 0x1e, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x43, 0x6f, 0x69, 0x6e, 0x4c, 0x6f, 0x67, 0x6f, 0x12, 0x1a, 0x0a, 0x08, 0x43,
	0x6f, 0x69, 0x6e, 0x55, 0x6e, 0x69, 0x74, 0x18, 0x28, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x43,
	0x6f, 0x69, 0x6e, 0x55, 0x6e, 0x69, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x41, 0x6d, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x32, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12,
	0x1c, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x3c, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x18, 0x0a,
	0x07, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x46, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x24, 0x0a, 0x0d, 0x41, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x50, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0d,
	0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x12, 0x38, 0x0a,
	0x05, 0x53, 0x74, 0x61, 0x74, 0x65, 0x18, 0x5a, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x22, 0x2e, 0x62,
	0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2e,
	0x76, 0x31, 0x2e, 0x57, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x53, 0x74, 0x61, 0x74, 0x65,
	0x52, 0x05, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x18, 0x64, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x41, 0x70, 0x70, 0x49, 0x44, 0x18, 0x6e, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x41, 0x70, 0x70, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49,
	0x44, 0x18, 0x78, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12,
	0x0f, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x82, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44,
	0x22, 0x83, 0x02, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x57, 0x69, 0x74, 0x68, 0x64,
	0x72, 0x61, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x41, 0x70,
	0x70, 0x49, 0x44, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x41, 0x70, 0x70, 0x49, 0x44,
	0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x1e, 0x0a, 0x0a, 0x43, 0x6f, 0x69, 0x6e,
	0x54, 0x79, 0x70, 0x65, 0x49, 0x44, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x43, 0x6f,
	0x69, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x49, 0x44, 0x12, 0x1c, 0x0a, 0x09, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x28, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x41, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x32, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x3a,
	0x0a, 0x0b, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x18, 0x3c, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x18, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x52, 0x0b, 0x41,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x2a, 0x0a, 0x10, 0x56, 0x65,
	0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x50,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x43, 0x6f, 0x64, 0x65, 0x22, 0x52, 0x0a, 0x16, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x57, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x38, 0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24,
	0x2e, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e,
	0x77, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x2e, 0x76, 0x31, 0x2e, 0x57, 0x69, 0x74, 0x68,
	0x64, 0x72, 0x61, 0x77, 0x52, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x71, 0x0a, 0x13, 0x47, 0x65,
	0x74, 0x57, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x14, 0x0a, 0x05, 0x41, 0x70, 0x70, 0x49, 0x44, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x41, 0x70, 0x70, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49,
	0x44, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12,
	0x16, 0x0a, 0x06, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x06, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x4c, 0x69, 0x6d, 0x69, 0x74,
	0x18, 0x28, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x22, 0x68, 0x0a,
	0x14, 0x47, 0x65, 0x74, 0x57, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3a, 0x0a, 0x05, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x18, 0x0a,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2e, 0x67, 0x61,
	0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x77, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x2e, 0x76,
	0x31, 0x2e, 0x57, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x52, 0x05, 0x49, 0x6e, 0x66, 0x6f,
	0x73, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x14, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x22, 0x5c, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x41, 0x70,
	0x70, 0x57, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x14, 0x0a, 0x05, 0x41, 0x70, 0x70, 0x49, 0x44, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x41, 0x70, 0x70, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x4f, 0x66, 0x66, 0x73, 0x65,
	0x74, 0x18, 0x14, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12,
	0x14, 0x0a, 0x05, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05,
	0x4c, 0x69, 0x6d, 0x69, 0x74, 0x22, 0x6b, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x41, 0x70, 0x70, 0x57,
	0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x3a, 0x0a, 0x05, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x24, 0x2e, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79,
	0x2e, 0x77, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x2e, 0x76, 0x31, 0x2e, 0x57, 0x69, 0x74,
	0x68, 0x64, 0x72, 0x61, 0x77, 0x52, 0x05, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x12, 0x14, 0x0a, 0x05,
	0x54, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x14, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x54, 0x6f, 0x74,
	0x61, 0x6c, 0x22, 0x69, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x4e, 0x41, 0x70, 0x70, 0x57, 0x69, 0x74,
	0x68, 0x64, 0x72, 0x61, 0x77, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x20, 0x0a,
	0x0b, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x41, 0x70, 0x70, 0x49, 0x44, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x41, 0x70, 0x70, 0x49, 0x44, 0x12,
	0x16, 0x0a, 0x06, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x14, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x06, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x4c, 0x69, 0x6d, 0x69, 0x74,
	0x18, 0x1e, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x22, 0x6c, 0x0a,
	0x18, 0x47, 0x65, 0x74, 0x4e, 0x41, 0x70, 0x70, 0x57, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3a, 0x0a, 0x05, 0x49, 0x6e, 0x66,
	0x6f, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x6c, 0x65, 0x64, 0x67, 0x65,
	0x72, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x77, 0x69, 0x74, 0x68, 0x64, 0x72,
	0x61, 0x77, 0x2e, 0x76, 0x31, 0x2e, 0x57, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x52, 0x05,
	0x49, 0x6e, 0x66, 0x6f, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x14,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x32, 0xf8, 0x04, 0x0a, 0x07,
	0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x12, 0x97, 0x01, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x57, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x12, 0x31, 0x2e, 0x6c, 0x65, 0x64,
	0x67, 0x65, 0x72, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x77, 0x69, 0x74, 0x68,
	0x64, 0x72, 0x61, 0x77, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x57, 0x69,
	0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x32, 0x2e,
	0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x77,
	0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x57, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x1e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x18, 0x3a, 0x01, 0x2a, 0x22, 0x13, 0x2f, 0x76,
	0x31, 0x2f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x2f, 0x77, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61,
	0x77, 0x12, 0x8f, 0x01, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x57, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61,
	0x77, 0x73, 0x12, 0x2f, 0x2e, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2e, 0x67, 0x61, 0x74, 0x65,
	0x77, 0x61, 0x79, 0x2e, 0x77, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x2e, 0x76, 0x31, 0x2e,
	0x47, 0x65, 0x74, 0x57, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x30, 0x2e, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2e, 0x67, 0x61, 0x74,
	0x65, 0x77, 0x61, 0x79, 0x2e, 0x77, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x2e, 0x76, 0x31,
	0x2e, 0x47, 0x65, 0x74, 0x57, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x16, 0x3a, 0x01, 0x2a,
	0x22, 0x11, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x65, 0x74, 0x2f, 0x77, 0x69, 0x74, 0x68, 0x64, 0x72,
	0x61, 0x77, 0x73, 0x12, 0x9c, 0x01, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x41, 0x70, 0x70, 0x57, 0x69,
	0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x73, 0x12, 0x32, 0x2e, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72,
	0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x77, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61,
	0x77, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x70, 0x70, 0x57, 0x69, 0x74, 0x68, 0x64,
	0x72, 0x61, 0x77, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x33, 0x2e, 0x6c, 0x65,
	0x64, 0x67, 0x65, 0x72, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x77, 0x69, 0x74,
	0x68, 0x64, 0x72, 0x61, 0x77, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x70, 0x70, 0x57,
	0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x20, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1a, 0x3a, 0x01, 0x2a, 0x22, 0x15, 0x2f, 0x76, 0x31,
	0x2f, 0x67, 0x65, 0x74, 0x2f, 0x61, 0x70, 0x70, 0x2f, 0x77, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61,
	0x77, 0x73, 0x12, 0xa1, 0x01, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x4e, 0x41, 0x70, 0x70, 0x57, 0x69,
	0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x73, 0x12, 0x33, 0x2e, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72,
	0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x77, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61,
	0x77, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4e, 0x41, 0x70, 0x70, 0x57, 0x69, 0x74, 0x68,
	0x64, 0x72, 0x61, 0x77, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x34, 0x2e, 0x6c,
	0x65, 0x64, 0x67, 0x65, 0x72, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x77, 0x69,
	0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4e, 0x41, 0x70,
	0x70, 0x57, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x22, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1c, 0x3a, 0x01, 0x2a, 0x22, 0x17, 0x2f,
	0x76, 0x31, 0x2f, 0x67, 0x65, 0x74, 0x2f, 0x6e, 0x2f, 0x61, 0x70, 0x70, 0x2f, 0x77, 0x69, 0x74,
	0x68, 0x64, 0x72, 0x61, 0x77, 0x73, 0x42, 0x3e, 0x5a, 0x3c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4e, 0x70, 0x6f, 0x6f, 0x6c, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f,
	0x72, 0x6d, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x6e, 0x70, 0x6f, 0x6f, 0x6c,
	0x2f, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2f, 0x67, 0x77, 0x2f, 0x76, 0x31, 0x2f, 0x77, 0x69,
	0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_npool_ledger_gw_v1_withdraw_withdraw_proto_rawDescOnce sync.Once
	file_npool_ledger_gw_v1_withdraw_withdraw_proto_rawDescData = file_npool_ledger_gw_v1_withdraw_withdraw_proto_rawDesc
)

func file_npool_ledger_gw_v1_withdraw_withdraw_proto_rawDescGZIP() []byte {
	file_npool_ledger_gw_v1_withdraw_withdraw_proto_rawDescOnce.Do(func() {
		file_npool_ledger_gw_v1_withdraw_withdraw_proto_rawDescData = protoimpl.X.CompressGZIP(file_npool_ledger_gw_v1_withdraw_withdraw_proto_rawDescData)
	})
	return file_npool_ledger_gw_v1_withdraw_withdraw_proto_rawDescData
}

var file_npool_ledger_gw_v1_withdraw_withdraw_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_npool_ledger_gw_v1_withdraw_withdraw_proto_goTypes = []interface{}{
	(*Withdraw)(nil),                 // 0: ledger.gateway.withdraw.v1.Withdraw
	(*CreateWithdrawRequest)(nil),    // 1: ledger.gateway.withdraw.v1.CreateWithdrawRequest
	(*CreateWithdrawResponse)(nil),   // 2: ledger.gateway.withdraw.v1.CreateWithdrawResponse
	(*GetWithdrawsRequest)(nil),      // 3: ledger.gateway.withdraw.v1.GetWithdrawsRequest
	(*GetWithdrawsResponse)(nil),     // 4: ledger.gateway.withdraw.v1.GetWithdrawsResponse
	(*GetAppWithdrawsRequest)(nil),   // 5: ledger.gateway.withdraw.v1.GetAppWithdrawsRequest
	(*GetAppWithdrawsResponse)(nil),  // 6: ledger.gateway.withdraw.v1.GetAppWithdrawsResponse
	(*GetNAppWithdrawsRequest)(nil),  // 7: ledger.gateway.withdraw.v1.GetNAppWithdrawsRequest
	(*GetNAppWithdrawsResponse)(nil), // 8: ledger.gateway.withdraw.v1.GetNAppWithdrawsResponse
	(v1.WithdrawState)(0),            // 9: basetypes.ledger.v1.WithdrawState
	(v11.SignMethod)(0),              // 10: basetypes.v1.SignMethod
}
var file_npool_ledger_gw_v1_withdraw_withdraw_proto_depIdxs = []int32{
	9,  // 0: ledger.gateway.withdraw.v1.Withdraw.State:type_name -> basetypes.ledger.v1.WithdrawState
	10, // 1: ledger.gateway.withdraw.v1.CreateWithdrawRequest.AccountType:type_name -> basetypes.v1.SignMethod
	0,  // 2: ledger.gateway.withdraw.v1.CreateWithdrawResponse.Info:type_name -> ledger.gateway.withdraw.v1.Withdraw
	0,  // 3: ledger.gateway.withdraw.v1.GetWithdrawsResponse.Infos:type_name -> ledger.gateway.withdraw.v1.Withdraw
	0,  // 4: ledger.gateway.withdraw.v1.GetAppWithdrawsResponse.Infos:type_name -> ledger.gateway.withdraw.v1.Withdraw
	0,  // 5: ledger.gateway.withdraw.v1.GetNAppWithdrawsResponse.Infos:type_name -> ledger.gateway.withdraw.v1.Withdraw
	1,  // 6: ledger.gateway.withdraw.v1.Gateway.CreateWithdraw:input_type -> ledger.gateway.withdraw.v1.CreateWithdrawRequest
	3,  // 7: ledger.gateway.withdraw.v1.Gateway.GetWithdraws:input_type -> ledger.gateway.withdraw.v1.GetWithdrawsRequest
	5,  // 8: ledger.gateway.withdraw.v1.Gateway.GetAppWithdraws:input_type -> ledger.gateway.withdraw.v1.GetAppWithdrawsRequest
	7,  // 9: ledger.gateway.withdraw.v1.Gateway.GetNAppWithdraws:input_type -> ledger.gateway.withdraw.v1.GetNAppWithdrawsRequest
	2,  // 10: ledger.gateway.withdraw.v1.Gateway.CreateWithdraw:output_type -> ledger.gateway.withdraw.v1.CreateWithdrawResponse
	4,  // 11: ledger.gateway.withdraw.v1.Gateway.GetWithdraws:output_type -> ledger.gateway.withdraw.v1.GetWithdrawsResponse
	6,  // 12: ledger.gateway.withdraw.v1.Gateway.GetAppWithdraws:output_type -> ledger.gateway.withdraw.v1.GetAppWithdrawsResponse
	8,  // 13: ledger.gateway.withdraw.v1.Gateway.GetNAppWithdraws:output_type -> ledger.gateway.withdraw.v1.GetNAppWithdrawsResponse
	10, // [10:14] is the sub-list for method output_type
	6,  // [6:10] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_npool_ledger_gw_v1_withdraw_withdraw_proto_init() }
func file_npool_ledger_gw_v1_withdraw_withdraw_proto_init() {
	if File_npool_ledger_gw_v1_withdraw_withdraw_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_npool_ledger_gw_v1_withdraw_withdraw_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Withdraw); i {
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
		file_npool_ledger_gw_v1_withdraw_withdraw_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateWithdrawRequest); i {
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
		file_npool_ledger_gw_v1_withdraw_withdraw_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateWithdrawResponse); i {
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
		file_npool_ledger_gw_v1_withdraw_withdraw_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetWithdrawsRequest); i {
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
		file_npool_ledger_gw_v1_withdraw_withdraw_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetWithdrawsResponse); i {
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
		file_npool_ledger_gw_v1_withdraw_withdraw_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAppWithdrawsRequest); i {
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
		file_npool_ledger_gw_v1_withdraw_withdraw_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAppWithdrawsResponse); i {
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
		file_npool_ledger_gw_v1_withdraw_withdraw_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetNAppWithdrawsRequest); i {
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
		file_npool_ledger_gw_v1_withdraw_withdraw_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetNAppWithdrawsResponse); i {
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
			RawDescriptor: file_npool_ledger_gw_v1_withdraw_withdraw_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_npool_ledger_gw_v1_withdraw_withdraw_proto_goTypes,
		DependencyIndexes: file_npool_ledger_gw_v1_withdraw_withdraw_proto_depIdxs,
		MessageInfos:      file_npool_ledger_gw_v1_withdraw_withdraw_proto_msgTypes,
	}.Build()
	File_npool_ledger_gw_v1_withdraw_withdraw_proto = out.File
	file_npool_ledger_gw_v1_withdraw_withdraw_proto_rawDesc = nil
	file_npool_ledger_gw_v1_withdraw_withdraw_proto_goTypes = nil
	file_npool_ledger_gw_v1_withdraw_withdraw_proto_depIdxs = nil
}
