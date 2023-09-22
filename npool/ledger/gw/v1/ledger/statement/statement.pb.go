// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.18.1
// source: npool/ledger/gw/v1/ledger/statement/statement.proto

package statement

import (
	v1 "github.com/NpoolPlatform/message/npool/basetypes/ledger/v1"
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

type Statement struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CoinTypeID   string       `protobuf:"bytes,10,opt,name=CoinTypeID,proto3" json:"CoinTypeID,omitempty"`
	CoinName     string       `protobuf:"bytes,20,opt,name=CoinName,proto3" json:"CoinName,omitempty"`
	DisplayNames []string     `protobuf:"bytes,21,rep,name=DisplayNames,proto3" json:"DisplayNames,omitempty"`
	CoinLogo     string       `protobuf:"bytes,30,opt,name=CoinLogo,proto3" json:"CoinLogo,omitempty"`
	CoinUnit     string       `protobuf:"bytes,40,opt,name=CoinUnit,proto3" json:"CoinUnit,omitempty"`
	IOType       v1.IOType    `protobuf:"varint,50,opt,name=IOType,proto3,enum=basetypes.ledger.v1.IOType" json:"IOType,omitempty"`
	IOSubType    v1.IOSubType `protobuf:"varint,60,opt,name=IOSubType,proto3,enum=basetypes.ledger.v1.IOSubType" json:"IOSubType,omitempty"`
	Amount       string       `protobuf:"bytes,70,opt,name=Amount,proto3" json:"Amount,omitempty"`
	IOExtra      string       `protobuf:"bytes,80,opt,name=IOExtra,proto3" json:"IOExtra,omitempty"`
	CreatedAt    uint32       `protobuf:"varint,90,opt,name=CreatedAt,proto3" json:"CreatedAt,omitempty"`
	UserID       string       `protobuf:"bytes,100,opt,name=UserID,proto3" json:"UserID,omitempty"`
	PhoneNO      string       `protobuf:"bytes,110,opt,name=PhoneNO,proto3" json:"PhoneNO,omitempty"`
	EmailAddress string       `protobuf:"bytes,120,opt,name=EmailAddress,proto3" json:"EmailAddress,omitempty"`
	AppID        string       `protobuf:"bytes,130,opt,name=AppID,proto3" json:"AppID,omitempty"`
	ID           string       `protobuf:"bytes,140,opt,name=ID,proto3" json:"ID,omitempty"`
}

func (x *Statement) Reset() {
	*x = Statement{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_ledger_gw_v1_ledger_statement_statement_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Statement) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Statement) ProtoMessage() {}

func (x *Statement) ProtoReflect() protoreflect.Message {
	mi := &file_npool_ledger_gw_v1_ledger_statement_statement_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Statement.ProtoReflect.Descriptor instead.
func (*Statement) Descriptor() ([]byte, []int) {
	return file_npool_ledger_gw_v1_ledger_statement_statement_proto_rawDescGZIP(), []int{0}
}

func (x *Statement) GetCoinTypeID() string {
	if x != nil {
		return x.CoinTypeID
	}
	return ""
}

func (x *Statement) GetCoinName() string {
	if x != nil {
		return x.CoinName
	}
	return ""
}

func (x *Statement) GetDisplayNames() []string {
	if x != nil {
		return x.DisplayNames
	}
	return nil
}

func (x *Statement) GetCoinLogo() string {
	if x != nil {
		return x.CoinLogo
	}
	return ""
}

func (x *Statement) GetCoinUnit() string {
	if x != nil {
		return x.CoinUnit
	}
	return ""
}

func (x *Statement) GetIOType() v1.IOType {
	if x != nil {
		return x.IOType
	}
	return v1.IOType(0)
}

func (x *Statement) GetIOSubType() v1.IOSubType {
	if x != nil {
		return x.IOSubType
	}
	return v1.IOSubType(0)
}

func (x *Statement) GetAmount() string {
	if x != nil {
		return x.Amount
	}
	return ""
}

func (x *Statement) GetIOExtra() string {
	if x != nil {
		return x.IOExtra
	}
	return ""
}

func (x *Statement) GetCreatedAt() uint32 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *Statement) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

func (x *Statement) GetPhoneNO() string {
	if x != nil {
		return x.PhoneNO
	}
	return ""
}

func (x *Statement) GetEmailAddress() string {
	if x != nil {
		return x.EmailAddress
	}
	return ""
}

func (x *Statement) GetAppID() string {
	if x != nil {
		return x.AppID
	}
	return ""
}

func (x *Statement) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

type GetStatementsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AppID   string `protobuf:"bytes,10,opt,name=AppID,proto3" json:"AppID,omitempty"`
	UserID  string `protobuf:"bytes,20,opt,name=UserID,proto3" json:"UserID,omitempty"`
	StartAt uint32 `protobuf:"varint,30,opt,name=StartAt,proto3" json:"StartAt,omitempty"`
	EndAt   uint32 `protobuf:"varint,40,opt,name=EndAt,proto3" json:"EndAt,omitempty"`
	Offset  int32  `protobuf:"varint,50,opt,name=Offset,proto3" json:"Offset,omitempty"`
	Limit   int32  `protobuf:"varint,60,opt,name=Limit,proto3" json:"Limit,omitempty"`
}

func (x *GetStatementsRequest) Reset() {
	*x = GetStatementsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_ledger_gw_v1_ledger_statement_statement_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetStatementsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetStatementsRequest) ProtoMessage() {}

func (x *GetStatementsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_ledger_gw_v1_ledger_statement_statement_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetStatementsRequest.ProtoReflect.Descriptor instead.
func (*GetStatementsRequest) Descriptor() ([]byte, []int) {
	return file_npool_ledger_gw_v1_ledger_statement_statement_proto_rawDescGZIP(), []int{1}
}

func (x *GetStatementsRequest) GetAppID() string {
	if x != nil {
		return x.AppID
	}
	return ""
}

func (x *GetStatementsRequest) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

func (x *GetStatementsRequest) GetStartAt() uint32 {
	if x != nil {
		return x.StartAt
	}
	return 0
}

func (x *GetStatementsRequest) GetEndAt() uint32 {
	if x != nil {
		return x.EndAt
	}
	return 0
}

func (x *GetStatementsRequest) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *GetStatementsRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type GetStatementsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Infos []*Statement `protobuf:"bytes,10,rep,name=Infos,proto3" json:"Infos,omitempty"`
	Total uint32       `protobuf:"varint,20,opt,name=Total,proto3" json:"Total,omitempty"`
}

func (x *GetStatementsResponse) Reset() {
	*x = GetStatementsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_ledger_gw_v1_ledger_statement_statement_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetStatementsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetStatementsResponse) ProtoMessage() {}

func (x *GetStatementsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_ledger_gw_v1_ledger_statement_statement_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetStatementsResponse.ProtoReflect.Descriptor instead.
func (*GetStatementsResponse) Descriptor() ([]byte, []int) {
	return file_npool_ledger_gw_v1_ledger_statement_statement_proto_rawDescGZIP(), []int{2}
}

func (x *GetStatementsResponse) GetInfos() []*Statement {
	if x != nil {
		return x.Infos
	}
	return nil
}

func (x *GetStatementsResponse) GetTotal() uint32 {
	if x != nil {
		return x.Total
	}
	return 0
}

type GetAppStatementsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TargetAppID string `protobuf:"bytes,10,opt,name=TargetAppID,proto3" json:"TargetAppID,omitempty"`
	Offset      int32  `protobuf:"varint,20,opt,name=Offset,proto3" json:"Offset,omitempty"`
	Limit       int32  `protobuf:"varint,30,opt,name=Limit,proto3" json:"Limit,omitempty"`
}

func (x *GetAppStatementsRequest) Reset() {
	*x = GetAppStatementsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_ledger_gw_v1_ledger_statement_statement_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAppStatementsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAppStatementsRequest) ProtoMessage() {}

func (x *GetAppStatementsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_ledger_gw_v1_ledger_statement_statement_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAppStatementsRequest.ProtoReflect.Descriptor instead.
func (*GetAppStatementsRequest) Descriptor() ([]byte, []int) {
	return file_npool_ledger_gw_v1_ledger_statement_statement_proto_rawDescGZIP(), []int{3}
}

func (x *GetAppStatementsRequest) GetTargetAppID() string {
	if x != nil {
		return x.TargetAppID
	}
	return ""
}

func (x *GetAppStatementsRequest) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *GetAppStatementsRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type GetAppStatementsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Infos []*Statement `protobuf:"bytes,10,rep,name=Infos,proto3" json:"Infos,omitempty"`
	Total uint32       `protobuf:"varint,20,opt,name=Total,proto3" json:"Total,omitempty"`
}

func (x *GetAppStatementsResponse) Reset() {
	*x = GetAppStatementsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_ledger_gw_v1_ledger_statement_statement_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAppStatementsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAppStatementsResponse) ProtoMessage() {}

func (x *GetAppStatementsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_ledger_gw_v1_ledger_statement_statement_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAppStatementsResponse.ProtoReflect.Descriptor instead.
func (*GetAppStatementsResponse) Descriptor() ([]byte, []int) {
	return file_npool_ledger_gw_v1_ledger_statement_statement_proto_rawDescGZIP(), []int{4}
}

func (x *GetAppStatementsResponse) GetInfos() []*Statement {
	if x != nil {
		return x.Infos
	}
	return nil
}

func (x *GetAppStatementsResponse) GetTotal() uint32 {
	if x != nil {
		return x.Total
	}
	return 0
}

var File_npool_ledger_gw_v1_ledger_statement_statement_proto protoreflect.FileDescriptor

var file_npool_ledger_gw_v1_ledger_statement_statement_proto_rawDesc = []byte{
	0x0a, 0x33, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2f, 0x67,
	0x77, 0x2f, 0x76, 0x31, 0x2f, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2f, 0x73, 0x74, 0x61, 0x74,
	0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x22, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2e, 0x67, 0x61,
	0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2e, 0x73, 0x74, 0x61,
	0x74, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x25, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x62,
	0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2f,
	0x76, 0x31, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe4,
	0x03, 0x0a, 0x09, 0x53, 0x74, 0x61, 0x74, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x1e, 0x0a, 0x0a,
	0x43, 0x6f, 0x69, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x49, 0x44, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x43, 0x6f, 0x69, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08,
	0x43, 0x6f, 0x69, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x43, 0x6f, 0x69, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x44, 0x69, 0x73, 0x70,
	0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x18, 0x15, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0c,
	0x44, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x12, 0x1a, 0x0a, 0x08,
	0x43, 0x6f, 0x69, 0x6e, 0x4c, 0x6f, 0x67, 0x6f, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x43, 0x6f, 0x69, 0x6e, 0x4c, 0x6f, 0x67, 0x6f, 0x12, 0x1a, 0x0a, 0x08, 0x43, 0x6f, 0x69, 0x6e,
	0x55, 0x6e, 0x69, 0x74, 0x18, 0x28, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x43, 0x6f, 0x69, 0x6e,
	0x55, 0x6e, 0x69, 0x74, 0x12, 0x33, 0x0a, 0x06, 0x49, 0x4f, 0x54, 0x79, 0x70, 0x65, 0x18, 0x32,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x1b, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x2e, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x4f, 0x54, 0x79, 0x70,
	0x65, 0x52, 0x06, 0x49, 0x4f, 0x54, 0x79, 0x70, 0x65, 0x12, 0x3c, 0x0a, 0x09, 0x49, 0x4f, 0x53,
	0x75, 0x62, 0x54, 0x79, 0x70, 0x65, 0x18, 0x3c, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1e, 0x2e, 0x62,
	0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2e,
	0x76, 0x31, 0x2e, 0x49, 0x4f, 0x53, 0x75, 0x62, 0x54, 0x79, 0x70, 0x65, 0x52, 0x09, 0x49, 0x4f,
	0x53, 0x75, 0x62, 0x54, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x41, 0x6d, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x46, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12,
	0x18, 0x0a, 0x07, 0x49, 0x4f, 0x45, 0x78, 0x74, 0x72, 0x61, 0x18, 0x50, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x49, 0x4f, 0x45, 0x78, 0x74, 0x72, 0x61, 0x12, 0x1c, 0x0a, 0x09, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x5a, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49,
	0x44, 0x18, 0x64, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12,
	0x18, 0x0a, 0x07, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x4f, 0x18, 0x6e, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x4f, 0x12, 0x22, 0x0a, 0x0c, 0x45, 0x6d, 0x61,
	0x69, 0x6c, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x78, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x15, 0x0a,
	0x05, 0x41, 0x70, 0x70, 0x49, 0x44, 0x18, 0x82, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x41,
	0x70, 0x70, 0x49, 0x44, 0x12, 0x0f, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x8c, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x49, 0x44, 0x22, 0xa2, 0x01, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x53, 0x74, 0x61,
	0x74, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14,
	0x0a, 0x05, 0x41, 0x70, 0x70, 0x49, 0x44, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x41,
	0x70, 0x70, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x14,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x18, 0x0a, 0x07,
	0x53, 0x74, 0x61, 0x72, 0x74, 0x41, 0x74, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x53,
	0x74, 0x61, 0x72, 0x74, 0x41, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x45, 0x6e, 0x64, 0x41, 0x74, 0x18,
	0x28, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x45, 0x6e, 0x64, 0x41, 0x74, 0x12, 0x16, 0x0a, 0x06,
	0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x32, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x4f, 0x66,
	0x66, 0x73, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x3c, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x05, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x22, 0x72, 0x0a, 0x15, 0x47, 0x65,
	0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x43, 0x0a, 0x05, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x18, 0x0a, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2e, 0x67, 0x61, 0x74, 0x65,
	0x77, 0x61, 0x79, 0x2e, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x65,
	0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x52, 0x05, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x6f, 0x74, 0x61,
	0x6c, 0x18, 0x14, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x22, 0x69,
	0x0a, 0x17, 0x47, 0x65, 0x74, 0x41, 0x70, 0x70, 0x53, 0x74, 0x61, 0x74, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x54, 0x61, 0x72,
	0x67, 0x65, 0x74, 0x41, 0x70, 0x70, 0x49, 0x44, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x41, 0x70, 0x70, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x4f,
	0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x14, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x4f, 0x66, 0x66,
	0x73, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x1e, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x05, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x22, 0x75, 0x0a, 0x18, 0x47, 0x65, 0x74,
	0x41, 0x70, 0x70, 0x53, 0x74, 0x61, 0x74, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x43, 0x0a, 0x05, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x18, 0x0a,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2e, 0x67, 0x61,
	0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2e, 0x73, 0x74, 0x61,
	0x74, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x6d,
	0x65, 0x6e, 0x74, 0x52, 0x05, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x6f,
	0x74, 0x61, 0x6c, 0x18, 0x14, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c,
	0x32, 0xdc, 0x02, 0x0a, 0x07, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x12, 0xa0, 0x01, 0x0a,
	0x0d, 0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x38,
	0x2e, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e,
	0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x65, 0x6d, 0x65, 0x6e, 0x74,
	0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x6d, 0x65, 0x6e, 0x74,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x39, 0x2e, 0x6c, 0x65, 0x64, 0x67, 0x65,
	0x72, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72,
	0x2e, 0x73, 0x74, 0x61, 0x74, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65,
	0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x1a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14, 0x3a, 0x01, 0x2a, 0x22, 0x0f,
	0x2f, 0x76, 0x31, 0x2f, 0x67, 0x65, 0x74, 0x2f, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12,
	0xad, 0x01, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x41, 0x70, 0x70, 0x53, 0x74, 0x61, 0x74, 0x65, 0x6d,
	0x65, 0x6e, 0x74, 0x73, 0x12, 0x3b, 0x2e, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2e, 0x67, 0x61,
	0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2e, 0x73, 0x74, 0x61,
	0x74, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x70, 0x70,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x3c, 0x2e, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77,
	0x61, 0x79, 0x2e, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x65, 0x6d,
	0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x70, 0x70, 0x53, 0x74, 0x61,
	0x74, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x1e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x18, 0x3a, 0x01, 0x2a, 0x22, 0x13, 0x2f, 0x76, 0x31, 0x2f,
	0x67, 0x65, 0x74, 0x2f, 0x61, 0x70, 0x70, 0x2f, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x42,
	0x46, 0x5a, 0x44, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4e, 0x70,
	0x6f, 0x6f, 0x6c, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x2f, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72,
	0x2f, 0x67, 0x77, 0x2f, 0x76, 0x31, 0x2f, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2f, 0x73, 0x74,
	0x61, 0x74, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_npool_ledger_gw_v1_ledger_statement_statement_proto_rawDescOnce sync.Once
	file_npool_ledger_gw_v1_ledger_statement_statement_proto_rawDescData = file_npool_ledger_gw_v1_ledger_statement_statement_proto_rawDesc
)

func file_npool_ledger_gw_v1_ledger_statement_statement_proto_rawDescGZIP() []byte {
	file_npool_ledger_gw_v1_ledger_statement_statement_proto_rawDescOnce.Do(func() {
		file_npool_ledger_gw_v1_ledger_statement_statement_proto_rawDescData = protoimpl.X.CompressGZIP(file_npool_ledger_gw_v1_ledger_statement_statement_proto_rawDescData)
	})
	return file_npool_ledger_gw_v1_ledger_statement_statement_proto_rawDescData
}

var file_npool_ledger_gw_v1_ledger_statement_statement_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_npool_ledger_gw_v1_ledger_statement_statement_proto_goTypes = []interface{}{
	(*Statement)(nil),                // 0: ledger.gateway.ledger.statement.v1.Statement
	(*GetStatementsRequest)(nil),     // 1: ledger.gateway.ledger.statement.v1.GetStatementsRequest
	(*GetStatementsResponse)(nil),    // 2: ledger.gateway.ledger.statement.v1.GetStatementsResponse
	(*GetAppStatementsRequest)(nil),  // 3: ledger.gateway.ledger.statement.v1.GetAppStatementsRequest
	(*GetAppStatementsResponse)(nil), // 4: ledger.gateway.ledger.statement.v1.GetAppStatementsResponse
	(v1.IOType)(0),                   // 5: basetypes.ledger.v1.IOType
	(v1.IOSubType)(0),                // 6: basetypes.ledger.v1.IOSubType
}
var file_npool_ledger_gw_v1_ledger_statement_statement_proto_depIdxs = []int32{
	5, // 0: ledger.gateway.ledger.statement.v1.Statement.IOType:type_name -> basetypes.ledger.v1.IOType
	6, // 1: ledger.gateway.ledger.statement.v1.Statement.IOSubType:type_name -> basetypes.ledger.v1.IOSubType
	0, // 2: ledger.gateway.ledger.statement.v1.GetStatementsResponse.Infos:type_name -> ledger.gateway.ledger.statement.v1.Statement
	0, // 3: ledger.gateway.ledger.statement.v1.GetAppStatementsResponse.Infos:type_name -> ledger.gateway.ledger.statement.v1.Statement
	1, // 4: ledger.gateway.ledger.statement.v1.Gateway.GetStatements:input_type -> ledger.gateway.ledger.statement.v1.GetStatementsRequest
	3, // 5: ledger.gateway.ledger.statement.v1.Gateway.GetAppStatements:input_type -> ledger.gateway.ledger.statement.v1.GetAppStatementsRequest
	2, // 6: ledger.gateway.ledger.statement.v1.Gateway.GetStatements:output_type -> ledger.gateway.ledger.statement.v1.GetStatementsResponse
	4, // 7: ledger.gateway.ledger.statement.v1.Gateway.GetAppStatements:output_type -> ledger.gateway.ledger.statement.v1.GetAppStatementsResponse
	6, // [6:8] is the sub-list for method output_type
	4, // [4:6] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_npool_ledger_gw_v1_ledger_statement_statement_proto_init() }
func file_npool_ledger_gw_v1_ledger_statement_statement_proto_init() {
	if File_npool_ledger_gw_v1_ledger_statement_statement_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_npool_ledger_gw_v1_ledger_statement_statement_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Statement); i {
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
		file_npool_ledger_gw_v1_ledger_statement_statement_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetStatementsRequest); i {
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
		file_npool_ledger_gw_v1_ledger_statement_statement_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetStatementsResponse); i {
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
		file_npool_ledger_gw_v1_ledger_statement_statement_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAppStatementsRequest); i {
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
		file_npool_ledger_gw_v1_ledger_statement_statement_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAppStatementsResponse); i {
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
			RawDescriptor: file_npool_ledger_gw_v1_ledger_statement_statement_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_npool_ledger_gw_v1_ledger_statement_statement_proto_goTypes,
		DependencyIndexes: file_npool_ledger_gw_v1_ledger_statement_statement_proto_depIdxs,
		MessageInfos:      file_npool_ledger_gw_v1_ledger_statement_statement_proto_msgTypes,
	}.Build()
	File_npool_ledger_gw_v1_ledger_statement_statement_proto = out.File
	file_npool_ledger_gw_v1_ledger_statement_statement_proto_rawDesc = nil
	file_npool_ledger_gw_v1_ledger_statement_statement_proto_goTypes = nil
	file_npool_ledger_gw_v1_ledger_statement_statement_proto_depIdxs = nil
}
