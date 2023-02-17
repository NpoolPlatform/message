// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.18.1
// source: npool/chain/gw/v1/tx/tx.proto

package tx

import (
	_ "github.com/NpoolPlatform/message/npool"
	account "github.com/NpoolPlatform/message/npool/account/mgr/v1/account"
	v1 "github.com/NpoolPlatform/message/npool/basetypes/v1"
	tx "github.com/NpoolPlatform/message/npool/chain/mgr/v1/tx"
	_ "github.com/NpoolPlatform/message/npool/chain/mw/v1/tx"
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

type Tx struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID            string                 `protobuf:"bytes,10,opt,name=ID,proto3" json:"ID,omitempty"`
	CoinTypeID    string                 `protobuf:"bytes,20,opt,name=CoinTypeID,proto3" json:"CoinTypeID,omitempty"`
	CoinName      string                 `protobuf:"bytes,30,opt,name=CoinName,proto3" json:"CoinName,omitempty"`
	CoinLogo      string                 `protobuf:"bytes,40,opt,name=CoinLogo,proto3" json:"CoinLogo,omitempty"`
	CoinUnit      string                 `protobuf:"bytes,50,opt,name=CoinUnit,proto3" json:"CoinUnit,omitempty"`
	CoinENV       string                 `protobuf:"bytes,60,opt,name=CoinENV,proto3" json:"CoinENV,omitempty"`
	AppID         string                 `protobuf:"bytes,70,opt,name=AppID,proto3" json:"AppID,omitempty"`
	AppName       string                 `protobuf:"bytes,80,opt,name=AppName,proto3" json:"AppName,omitempty"`
	FromAccountID string                 `protobuf:"bytes,90,opt,name=FromAccountID,proto3" json:"FromAccountID,omitempty"`
	FromAddress   string                 `protobuf:"bytes,100,opt,name=FromAddress,proto3" json:"FromAddress,omitempty"`
	FromUsedFor   account.AccountUsedFor `protobuf:"varint,110,opt,name=FromUsedFor,proto3,enum=account.manager.account1.v1.AccountUsedFor" json:"FromUsedFor,omitempty"`
	ToAccountID   string                 `protobuf:"bytes,120,opt,name=ToAccountID,proto3" json:"ToAccountID,omitempty"`
	ToAddress     string                 `protobuf:"bytes,130,opt,name=ToAddress,proto3" json:"ToAddress,omitempty"`
	ToUsedFor     account.AccountUsedFor `protobuf:"varint,140,opt,name=ToUsedFor,proto3,enum=account.manager.account1.v1.AccountUsedFor" json:"ToUsedFor,omitempty"`
	Amount        string                 `protobuf:"bytes,150,opt,name=Amount,proto3" json:"Amount,omitempty"`
	FeeAmount     string                 `protobuf:"bytes,160,opt,name=FeeAmount,proto3" json:"FeeAmount,omitempty"`
	ChainTxID     string                 `protobuf:"bytes,170,opt,name=ChainTxID,proto3" json:"ChainTxID,omitempty"`
	State         tx.TxState             `protobuf:"varint,180,opt,name=State,proto3,enum=chain.manager.tx.v1.TxState" json:"State,omitempty"`
	Extra         string                 `protobuf:"bytes,190,opt,name=Extra,proto3" json:"Extra,omitempty"`
	Type          v1.TxType              `protobuf:"varint,200,opt,name=Type,proto3,enum=basetypes.v1.TxType" json:"Type,omitempty"`
	CreatedAt     uint32                 `protobuf:"varint,210,opt,name=CreatedAt,proto3" json:"CreatedAt,omitempty"`
	UpdatedAt     uint32                 `protobuf:"varint,220,opt,name=UpdatedAt,proto3" json:"UpdatedAt,omitempty"`
}

func (x *Tx) Reset() {
	*x = Tx{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_chain_gw_v1_tx_tx_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Tx) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Tx) ProtoMessage() {}

func (x *Tx) ProtoReflect() protoreflect.Message {
	mi := &file_npool_chain_gw_v1_tx_tx_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Tx.ProtoReflect.Descriptor instead.
func (*Tx) Descriptor() ([]byte, []int) {
	return file_npool_chain_gw_v1_tx_tx_proto_rawDescGZIP(), []int{0}
}

func (x *Tx) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *Tx) GetCoinTypeID() string {
	if x != nil {
		return x.CoinTypeID
	}
	return ""
}

func (x *Tx) GetCoinName() string {
	if x != nil {
		return x.CoinName
	}
	return ""
}

func (x *Tx) GetCoinLogo() string {
	if x != nil {
		return x.CoinLogo
	}
	return ""
}

func (x *Tx) GetCoinUnit() string {
	if x != nil {
		return x.CoinUnit
	}
	return ""
}

func (x *Tx) GetCoinENV() string {
	if x != nil {
		return x.CoinENV
	}
	return ""
}

func (x *Tx) GetAppID() string {
	if x != nil {
		return x.AppID
	}
	return ""
}

func (x *Tx) GetAppName() string {
	if x != nil {
		return x.AppName
	}
	return ""
}

func (x *Tx) GetFromAccountID() string {
	if x != nil {
		return x.FromAccountID
	}
	return ""
}

func (x *Tx) GetFromAddress() string {
	if x != nil {
		return x.FromAddress
	}
	return ""
}

func (x *Tx) GetFromUsedFor() account.AccountUsedFor {
	if x != nil {
		return x.FromUsedFor
	}
	return account.AccountUsedFor(0)
}

func (x *Tx) GetToAccountID() string {
	if x != nil {
		return x.ToAccountID
	}
	return ""
}

func (x *Tx) GetToAddress() string {
	if x != nil {
		return x.ToAddress
	}
	return ""
}

func (x *Tx) GetToUsedFor() account.AccountUsedFor {
	if x != nil {
		return x.ToUsedFor
	}
	return account.AccountUsedFor(0)
}

func (x *Tx) GetAmount() string {
	if x != nil {
		return x.Amount
	}
	return ""
}

func (x *Tx) GetFeeAmount() string {
	if x != nil {
		return x.FeeAmount
	}
	return ""
}

func (x *Tx) GetChainTxID() string {
	if x != nil {
		return x.ChainTxID
	}
	return ""
}

func (x *Tx) GetState() tx.TxState {
	if x != nil {
		return x.State
	}
	return tx.TxState(0)
}

func (x *Tx) GetExtra() string {
	if x != nil {
		return x.Extra
	}
	return ""
}

func (x *Tx) GetType() v1.TxType {
	if x != nil {
		return x.Type
	}
	return v1.TxType(0)
}

func (x *Tx) GetCreatedAt() uint32 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *Tx) GetUpdatedAt() uint32 {
	if x != nil {
		return x.UpdatedAt
	}
	return 0
}

type GetTxsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Offset int32 `protobuf:"varint,10,opt,name=Offset,proto3" json:"Offset,omitempty"`
	Limit  int32 `protobuf:"varint,20,opt,name=Limit,proto3" json:"Limit,omitempty"`
}

func (x *GetTxsRequest) Reset() {
	*x = GetTxsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_chain_gw_v1_tx_tx_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTxsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTxsRequest) ProtoMessage() {}

func (x *GetTxsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_chain_gw_v1_tx_tx_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTxsRequest.ProtoReflect.Descriptor instead.
func (*GetTxsRequest) Descriptor() ([]byte, []int) {
	return file_npool_chain_gw_v1_tx_tx_proto_rawDescGZIP(), []int{1}
}

func (x *GetTxsRequest) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *GetTxsRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type GetTxsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Infos []*Tx  `protobuf:"bytes,10,rep,name=Infos,proto3" json:"Infos,omitempty"`
	Total uint32 `protobuf:"varint,20,opt,name=Total,proto3" json:"Total,omitempty"`
}

func (x *GetTxsResponse) Reset() {
	*x = GetTxsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_chain_gw_v1_tx_tx_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTxsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTxsResponse) ProtoMessage() {}

func (x *GetTxsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_chain_gw_v1_tx_tx_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTxsResponse.ProtoReflect.Descriptor instead.
func (*GetTxsResponse) Descriptor() ([]byte, []int) {
	return file_npool_chain_gw_v1_tx_tx_proto_rawDescGZIP(), []int{2}
}

func (x *GetTxsResponse) GetInfos() []*Tx {
	if x != nil {
		return x.Infos
	}
	return nil
}

func (x *GetTxsResponse) GetTotal() uint32 {
	if x != nil {
		return x.Total
	}
	return 0
}

var File_npool_chain_gw_v1_tx_tx_proto protoreflect.FileDescriptor

var file_npool_chain_gw_v1_tx_tx_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2f, 0x67, 0x77,
	0x2f, 0x76, 0x31, 0x2f, 0x74, 0x78, 0x2f, 0x74, 0x78, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x13, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x74,
	0x78, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1f, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x78, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x6e, 0x70, 0x6f, 0x6f, 0x6c,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x63, 0x68,
	0x61, 0x69, 0x6e, 0x2f, 0x6d, 0x67, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x78, 0x2f, 0x74, 0x78,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x63, 0x68,
	0x61, 0x69, 0x6e, 0x2f, 0x6d, 0x77, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x78, 0x2f, 0x74, 0x78, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2a, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x2f, 0x6d, 0x67, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x82, 0x06, 0x0a, 0x02, 0x54, 0x78, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x12, 0x1e, 0x0a, 0x0a, 0x43, 0x6f, 0x69, 0x6e,
	0x54, 0x79, 0x70, 0x65, 0x49, 0x44, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x43, 0x6f,
	0x69, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08, 0x43, 0x6f, 0x69, 0x6e,
	0x4e, 0x61, 0x6d, 0x65, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x43, 0x6f, 0x69, 0x6e,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x43, 0x6f, 0x69, 0x6e, 0x4c, 0x6f, 0x67, 0x6f,
	0x18, 0x28, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x43, 0x6f, 0x69, 0x6e, 0x4c, 0x6f, 0x67, 0x6f,
	0x12, 0x1a, 0x0a, 0x08, 0x43, 0x6f, 0x69, 0x6e, 0x55, 0x6e, 0x69, 0x74, 0x18, 0x32, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x43, 0x6f, 0x69, 0x6e, 0x55, 0x6e, 0x69, 0x74, 0x12, 0x18, 0x0a, 0x07,
	0x43, 0x6f, 0x69, 0x6e, 0x45, 0x4e, 0x56, 0x18, 0x3c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x43,
	0x6f, 0x69, 0x6e, 0x45, 0x4e, 0x56, 0x12, 0x14, 0x0a, 0x05, 0x41, 0x70, 0x70, 0x49, 0x44, 0x18,
	0x46, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x41, 0x70, 0x70, 0x49, 0x44, 0x12, 0x18, 0x0a, 0x07,
	0x41, 0x70, 0x70, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x50, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x41,
	0x70, 0x70, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x46, 0x72, 0x6f, 0x6d, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x5a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x46,
	0x72, 0x6f, 0x6d, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x44, 0x12, 0x20, 0x0a, 0x0b,
	0x46, 0x72, 0x6f, 0x6d, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x64, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x46, 0x72, 0x6f, 0x6d, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x4d,
	0x0a, 0x0b, 0x46, 0x72, 0x6f, 0x6d, 0x55, 0x73, 0x65, 0x64, 0x46, 0x6f, 0x72, 0x18, 0x6e, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x2b, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x6d, 0x61,
	0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x31, 0x2e, 0x76,
	0x31, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x55, 0x73, 0x65, 0x64, 0x46, 0x6f, 0x72,
	0x52, 0x0b, 0x46, 0x72, 0x6f, 0x6d, 0x55, 0x73, 0x65, 0x64, 0x46, 0x6f, 0x72, 0x12, 0x20, 0x0a,
	0x0b, 0x54, 0x6f, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x78, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x54, 0x6f, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x44, 0x12,
	0x1d, 0x0a, 0x09, 0x54, 0x6f, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x82, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x54, 0x6f, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x4a,
	0x0a, 0x09, 0x54, 0x6f, 0x55, 0x73, 0x65, 0x64, 0x46, 0x6f, 0x72, 0x18, 0x8c, 0x01, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x2b, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x6d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x72, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x31, 0x2e, 0x76, 0x31,
	0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x55, 0x73, 0x65, 0x64, 0x46, 0x6f, 0x72, 0x52,
	0x09, 0x54, 0x6f, 0x55, 0x73, 0x65, 0x64, 0x46, 0x6f, 0x72, 0x12, 0x17, 0x0a, 0x06, 0x41, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x96, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x41, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x1d, 0x0a, 0x09, 0x46, 0x65, 0x65, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0xa0, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x46, 0x65, 0x65, 0x41, 0x6d, 0x6f, 0x75,
	0x6e, 0x74, 0x12, 0x1d, 0x0a, 0x09, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x54, 0x78, 0x49, 0x44, 0x18,
	0xaa, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x54, 0x78, 0x49,
	0x44, 0x12, 0x33, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x74, 0x65, 0x18, 0xb4, 0x01, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x1c, 0x2e, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65,
	0x72, 0x2e, 0x74, 0x78, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x78, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52,
	0x05, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x15, 0x0a, 0x05, 0x45, 0x78, 0x74, 0x72, 0x61, 0x18,
	0xbe, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x45, 0x78, 0x74, 0x72, 0x61, 0x12, 0x29, 0x0a,
	0x04, 0x54, 0x79, 0x70, 0x65, 0x18, 0xc8, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x14, 0x2e, 0x62,
	0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x78, 0x54, 0x79,
	0x70, 0x65, 0x52, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1d, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0xd2, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x18, 0xdc, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x3d, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x54, 0x78, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x4f, 0x66, 0x66, 0x73, 0x65,
	0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12,
	0x14, 0x0a, 0x05, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x14, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05,
	0x4c, 0x69, 0x6d, 0x69, 0x74, 0x22, 0x55, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x54, 0x78, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2d, 0x0a, 0x05, 0x49, 0x6e, 0x66, 0x6f, 0x73,
	0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x67,
	0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x74, 0x78, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x78, 0x52,
	0x05, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x18,
	0x14, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x32, 0x74, 0x0a, 0x07,
	0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x12, 0x69, 0x0a, 0x06, 0x47, 0x65, 0x74, 0x54, 0x78,
	0x73, 0x12, 0x22, 0x2e, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61,
	0x79, 0x2e, 0x74, 0x78, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x78, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x67, 0x61,
	0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x74, 0x78, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x54,
	0x78, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x16, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x10, 0x22, 0x0b, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x65, 0x74, 0x2f, 0x74, 0x78, 0x73, 0x3a,
	0x01, 0x2a, 0x42, 0x37, 0x5a, 0x35, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x4e, 0x70, 0x6f, 0x6f, 0x6c, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x63, 0x68, 0x61,
	0x69, 0x6e, 0x2f, 0x67, 0x77, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x78, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_npool_chain_gw_v1_tx_tx_proto_rawDescOnce sync.Once
	file_npool_chain_gw_v1_tx_tx_proto_rawDescData = file_npool_chain_gw_v1_tx_tx_proto_rawDesc
)

func file_npool_chain_gw_v1_tx_tx_proto_rawDescGZIP() []byte {
	file_npool_chain_gw_v1_tx_tx_proto_rawDescOnce.Do(func() {
		file_npool_chain_gw_v1_tx_tx_proto_rawDescData = protoimpl.X.CompressGZIP(file_npool_chain_gw_v1_tx_tx_proto_rawDescData)
	})
	return file_npool_chain_gw_v1_tx_tx_proto_rawDescData
}

var file_npool_chain_gw_v1_tx_tx_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_npool_chain_gw_v1_tx_tx_proto_goTypes = []interface{}{
	(*Tx)(nil),                  // 0: chain.gateway.tx.v1.Tx
	(*GetTxsRequest)(nil),       // 1: chain.gateway.tx.v1.GetTxsRequest
	(*GetTxsResponse)(nil),      // 2: chain.gateway.tx.v1.GetTxsResponse
	(account.AccountUsedFor)(0), // 3: account.manager.account1.v1.AccountUsedFor
	(tx.TxState)(0),             // 4: chain.manager.tx.v1.TxState
	(v1.TxType)(0),              // 5: basetypes.v1.TxType
}
var file_npool_chain_gw_v1_tx_tx_proto_depIdxs = []int32{
	3, // 0: chain.gateway.tx.v1.Tx.FromUsedFor:type_name -> account.manager.account1.v1.AccountUsedFor
	3, // 1: chain.gateway.tx.v1.Tx.ToUsedFor:type_name -> account.manager.account1.v1.AccountUsedFor
	4, // 2: chain.gateway.tx.v1.Tx.State:type_name -> chain.manager.tx.v1.TxState
	5, // 3: chain.gateway.tx.v1.Tx.Type:type_name -> basetypes.v1.TxType
	0, // 4: chain.gateway.tx.v1.GetTxsResponse.Infos:type_name -> chain.gateway.tx.v1.Tx
	1, // 5: chain.gateway.tx.v1.Gateway.GetTxs:input_type -> chain.gateway.tx.v1.GetTxsRequest
	2, // 6: chain.gateway.tx.v1.Gateway.GetTxs:output_type -> chain.gateway.tx.v1.GetTxsResponse
	6, // [6:7] is the sub-list for method output_type
	5, // [5:6] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_npool_chain_gw_v1_tx_tx_proto_init() }
func file_npool_chain_gw_v1_tx_tx_proto_init() {
	if File_npool_chain_gw_v1_tx_tx_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_npool_chain_gw_v1_tx_tx_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Tx); i {
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
		file_npool_chain_gw_v1_tx_tx_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTxsRequest); i {
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
		file_npool_chain_gw_v1_tx_tx_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTxsResponse); i {
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
			RawDescriptor: file_npool_chain_gw_v1_tx_tx_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_npool_chain_gw_v1_tx_tx_proto_goTypes,
		DependencyIndexes: file_npool_chain_gw_v1_tx_tx_proto_depIdxs,
		MessageInfos:      file_npool_chain_gw_v1_tx_tx_proto_msgTypes,
	}.Build()
	File_npool_chain_gw_v1_tx_tx_proto = out.File
	file_npool_chain_gw_v1_tx_tx_proto_rawDesc = nil
	file_npool_chain_gw_v1_tx_tx_proto_goTypes = nil
	file_npool_chain_gw_v1_tx_tx_proto_depIdxs = nil
}
