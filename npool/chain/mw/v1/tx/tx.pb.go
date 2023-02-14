// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.18.1
// source: npool/chain/mw/v1/tx/tx.proto

package tx

import (
	_ "github.com/NpoolPlatform/message/npool"
	tx "github.com/NpoolPlatform/message/npool/chain/mgr/v1/tx"
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

	// @inject_tag: sql:"id"
	ID string `protobuf:"bytes,10,opt,name=ID,proto3" json:"ID,omitempty" sql:"id"`
	// @inject_tag: sql:"coin_type_id"
	CoinTypeID string `protobuf:"bytes,20,opt,name=CoinTypeID,proto3" json:"CoinTypeID,omitempty" sql:"coin_type_id"`
	// @inject_tag: sql:"coin_name"
	CoinName string `protobuf:"bytes,30,opt,name=CoinName,proto3" json:"CoinName,omitempty" sql:"coin_name"`
	// @inject_tag: sql:"coin_logo"
	CoinLogo string `protobuf:"bytes,40,opt,name=CoinLogo,proto3" json:"CoinLogo,omitempty" sql:"coin_logo"`
	// @inject_tag: sql:"coin_unit"
	CoinUnit string `protobuf:"bytes,50,opt,name=CoinUnit,proto3" json:"CoinUnit,omitempty" sql:"coin_unit"`
	// @inject_tag: sql:"coin_env"
	CoinENV string `protobuf:"bytes,60,opt,name=CoinENV,proto3" json:"CoinENV,omitempty" sql:"coin_env"`
	// @inject_tag: sql:"from_account_id"
	FromAccountID string `protobuf:"bytes,70,opt,name=FromAccountID,proto3" json:"FromAccountID,omitempty" sql:"from_account_id"`
	// @inject_tag: sql:"to_account_id"
	ToAccountID string `protobuf:"bytes,80,opt,name=ToAccountID,proto3" json:"ToAccountID,omitempty" sql:"to_account_id"`
	// @inject_tag: sql:"amount"
	Amount string `protobuf:"bytes,90,opt,name=Amount,proto3" json:"Amount,omitempty" sql:"amount"`
	// @inject_tag: sql:"fee_amount"
	FeeAmount string `protobuf:"bytes,100,opt,name=FeeAmount,proto3" json:"FeeAmount,omitempty" sql:"fee_amount"`
	// @inject_tag: sql:"chain_tx_id"
	ChainTxID string     `protobuf:"bytes,110,opt,name=ChainTxID,proto3" json:"ChainTxID,omitempty" sql:"chain_tx_id"`
	State     tx.TxState `protobuf:"varint,120,opt,name=State,proto3,enum=chain.manager.tx.v1.TxState" json:"State,omitempty"`
	// @inject_tag: sql:"state"
	StateStr string `protobuf:"bytes,130,opt,name=StateStr,proto3" json:"StateStr,omitempty" sql:"state"`
	// @inject_tag: sql:"extra"
	Extra string    `protobuf:"bytes,140,opt,name=Extra,proto3" json:"Extra,omitempty" sql:"extra"`
	Type  tx.TxType `protobuf:"varint,150,opt,name=Type,proto3,enum=chain.manager.tx.v1.TxType" json:"Type,omitempty"`
	// @inject_tag: sql:"type"
	TypeStr string `protobuf:"bytes,160,opt,name=TypeStr,proto3" json:"TypeStr,omitempty" sql:"type"`
	// @inject_tag: sql:"created_at"
	CreatedAt uint32 `protobuf:"varint,170,opt,name=CreatedAt,proto3" json:"CreatedAt,omitempty" sql:"created_at"`
	// @inject_tag: sql:"updated_at"
	UpdatedAt uint32 `protobuf:"varint,180,opt,name=UpdatedAt,proto3" json:"UpdatedAt,omitempty" sql:"updated_at"`
}

func (x *Tx) Reset() {
	*x = Tx{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_chain_mw_v1_tx_tx_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Tx) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Tx) ProtoMessage() {}

func (x *Tx) ProtoReflect() protoreflect.Message {
	mi := &file_npool_chain_mw_v1_tx_tx_proto_msgTypes[0]
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
	return file_npool_chain_mw_v1_tx_tx_proto_rawDescGZIP(), []int{0}
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

func (x *Tx) GetFromAccountID() string {
	if x != nil {
		return x.FromAccountID
	}
	return ""
}

func (x *Tx) GetToAccountID() string {
	if x != nil {
		return x.ToAccountID
	}
	return ""
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

func (x *Tx) GetStateStr() string {
	if x != nil {
		return x.StateStr
	}
	return ""
}

func (x *Tx) GetExtra() string {
	if x != nil {
		return x.Extra
	}
	return ""
}

func (x *Tx) GetType() tx.TxType {
	if x != nil {
		return x.Type
	}
	return tx.TxType(0)
}

func (x *Tx) GetTypeStr() string {
	if x != nil {
		return x.TypeStr
	}
	return ""
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

type CreateTxRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *tx.TxReq `protobuf:"bytes,10,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (x *CreateTxRequest) Reset() {
	*x = CreateTxRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_chain_mw_v1_tx_tx_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTxRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTxRequest) ProtoMessage() {}

func (x *CreateTxRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_chain_mw_v1_tx_tx_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTxRequest.ProtoReflect.Descriptor instead.
func (*CreateTxRequest) Descriptor() ([]byte, []int) {
	return file_npool_chain_mw_v1_tx_tx_proto_rawDescGZIP(), []int{1}
}

func (x *CreateTxRequest) GetInfo() *tx.TxReq {
	if x != nil {
		return x.Info
	}
	return nil
}

type CreateTxResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *Tx `protobuf:"bytes,10,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (x *CreateTxResponse) Reset() {
	*x = CreateTxResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_chain_mw_v1_tx_tx_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTxResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTxResponse) ProtoMessage() {}

func (x *CreateTxResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_chain_mw_v1_tx_tx_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTxResponse.ProtoReflect.Descriptor instead.
func (*CreateTxResponse) Descriptor() ([]byte, []int) {
	return file_npool_chain_mw_v1_tx_tx_proto_rawDescGZIP(), []int{2}
}

func (x *CreateTxResponse) GetInfo() *Tx {
	if x != nil {
		return x.Info
	}
	return nil
}

type CreateTxsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Infos []*tx.TxReq `protobuf:"bytes,10,rep,name=Infos,proto3" json:"Infos,omitempty"`
}

func (x *CreateTxsRequest) Reset() {
	*x = CreateTxsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_chain_mw_v1_tx_tx_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTxsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTxsRequest) ProtoMessage() {}

func (x *CreateTxsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_chain_mw_v1_tx_tx_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTxsRequest.ProtoReflect.Descriptor instead.
func (*CreateTxsRequest) Descriptor() ([]byte, []int) {
	return file_npool_chain_mw_v1_tx_tx_proto_rawDescGZIP(), []int{3}
}

func (x *CreateTxsRequest) GetInfos() []*tx.TxReq {
	if x != nil {
		return x.Infos
	}
	return nil
}

type CreateTxsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Infos []*Tx `protobuf:"bytes,10,rep,name=Infos,proto3" json:"Infos,omitempty"`
}

func (x *CreateTxsResponse) Reset() {
	*x = CreateTxsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_chain_mw_v1_tx_tx_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTxsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTxsResponse) ProtoMessage() {}

func (x *CreateTxsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_chain_mw_v1_tx_tx_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTxsResponse.ProtoReflect.Descriptor instead.
func (*CreateTxsResponse) Descriptor() ([]byte, []int) {
	return file_npool_chain_mw_v1_tx_tx_proto_rawDescGZIP(), []int{4}
}

func (x *CreateTxsResponse) GetInfos() []*Tx {
	if x != nil {
		return x.Infos
	}
	return nil
}

type GetTxRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID string `protobuf:"bytes,10,opt,name=ID,proto3" json:"ID,omitempty"`
}

func (x *GetTxRequest) Reset() {
	*x = GetTxRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_chain_mw_v1_tx_tx_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTxRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTxRequest) ProtoMessage() {}

func (x *GetTxRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_chain_mw_v1_tx_tx_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTxRequest.ProtoReflect.Descriptor instead.
func (*GetTxRequest) Descriptor() ([]byte, []int) {
	return file_npool_chain_mw_v1_tx_tx_proto_rawDescGZIP(), []int{5}
}

func (x *GetTxRequest) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

type GetTxResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *Tx `protobuf:"bytes,10,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (x *GetTxResponse) Reset() {
	*x = GetTxResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_chain_mw_v1_tx_tx_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTxResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTxResponse) ProtoMessage() {}

func (x *GetTxResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_chain_mw_v1_tx_tx_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTxResponse.ProtoReflect.Descriptor instead.
func (*GetTxResponse) Descriptor() ([]byte, []int) {
	return file_npool_chain_mw_v1_tx_tx_proto_rawDescGZIP(), []int{6}
}

func (x *GetTxResponse) GetInfo() *Tx {
	if x != nil {
		return x.Info
	}
	return nil
}

type GetTxsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Conds  *tx.Conds `protobuf:"bytes,10,opt,name=Conds,proto3" json:"Conds,omitempty"`
	Offset int32     `protobuf:"varint,20,opt,name=Offset,proto3" json:"Offset,omitempty"`
	Limit  int32     `protobuf:"varint,30,opt,name=Limit,proto3" json:"Limit,omitempty"`
}

func (x *GetTxsRequest) Reset() {
	*x = GetTxsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_chain_mw_v1_tx_tx_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTxsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTxsRequest) ProtoMessage() {}

func (x *GetTxsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_chain_mw_v1_tx_tx_proto_msgTypes[7]
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
	return file_npool_chain_mw_v1_tx_tx_proto_rawDescGZIP(), []int{7}
}

func (x *GetTxsRequest) GetConds() *tx.Conds {
	if x != nil {
		return x.Conds
	}
	return nil
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
		mi := &file_npool_chain_mw_v1_tx_tx_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTxsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTxsResponse) ProtoMessage() {}

func (x *GetTxsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_chain_mw_v1_tx_tx_proto_msgTypes[8]
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
	return file_npool_chain_mw_v1_tx_tx_proto_rawDescGZIP(), []int{8}
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

type UpdateTxRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *tx.TxReq `protobuf:"bytes,10,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (x *UpdateTxRequest) Reset() {
	*x = UpdateTxRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_chain_mw_v1_tx_tx_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateTxRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateTxRequest) ProtoMessage() {}

func (x *UpdateTxRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_chain_mw_v1_tx_tx_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateTxRequest.ProtoReflect.Descriptor instead.
func (*UpdateTxRequest) Descriptor() ([]byte, []int) {
	return file_npool_chain_mw_v1_tx_tx_proto_rawDescGZIP(), []int{9}
}

func (x *UpdateTxRequest) GetInfo() *tx.TxReq {
	if x != nil {
		return x.Info
	}
	return nil
}

type UpdateTxResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *Tx `protobuf:"bytes,10,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (x *UpdateTxResponse) Reset() {
	*x = UpdateTxResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_chain_mw_v1_tx_tx_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateTxResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateTxResponse) ProtoMessage() {}

func (x *UpdateTxResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_chain_mw_v1_tx_tx_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateTxResponse.ProtoReflect.Descriptor instead.
func (*UpdateTxResponse) Descriptor() ([]byte, []int) {
	return file_npool_chain_mw_v1_tx_tx_proto_rawDescGZIP(), []int{10}
}

func (x *UpdateTxResponse) GetInfo() *Tx {
	if x != nil {
		return x.Info
	}
	return nil
}

var File_npool_chain_mw_v1_tx_tx_proto protoreflect.FileDescriptor

var file_npool_chain_mw_v1_tx_tx_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2f, 0x6d, 0x77,
	0x2f, 0x76, 0x31, 0x2f, 0x74, 0x78, 0x2f, 0x74, 0x78, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x16, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72,
	0x65, 0x2e, 0x74, 0x78, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x6e, 0x70, 0x6f,
	0x6f, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f,
	0x63, 0x68, 0x61, 0x69, 0x6e, 0x2f, 0x6d, 0x67, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x78, 0x2f,
	0x74, 0x78, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb1, 0x04, 0x0a, 0x02, 0x54, 0x78, 0x12,
	0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x12,
	0x1e, 0x0a, 0x0a, 0x43, 0x6f, 0x69, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x49, 0x44, 0x18, 0x14, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x43, 0x6f, 0x69, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x49, 0x44, 0x12,
	0x1a, 0x0a, 0x08, 0x43, 0x6f, 0x69, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x1e, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x43, 0x6f, 0x69, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x43,
	0x6f, 0x69, 0x6e, 0x4c, 0x6f, 0x67, 0x6f, 0x18, 0x28, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x43,
	0x6f, 0x69, 0x6e, 0x4c, 0x6f, 0x67, 0x6f, 0x12, 0x1a, 0x0a, 0x08, 0x43, 0x6f, 0x69, 0x6e, 0x55,
	0x6e, 0x69, 0x74, 0x18, 0x32, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x43, 0x6f, 0x69, 0x6e, 0x55,
	0x6e, 0x69, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x43, 0x6f, 0x69, 0x6e, 0x45, 0x4e, 0x56, 0x18, 0x3c,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x43, 0x6f, 0x69, 0x6e, 0x45, 0x4e, 0x56, 0x12, 0x24, 0x0a,
	0x0d, 0x46, 0x72, 0x6f, 0x6d, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x46,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x46, 0x72, 0x6f, 0x6d, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x49, 0x44, 0x12, 0x20, 0x0a, 0x0b, 0x54, 0x6f, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x49, 0x44, 0x18, 0x50, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x54, 0x6f, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x5a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1c, 0x0a,
	0x09, 0x46, 0x65, 0x65, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x64, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x46, 0x65, 0x65, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x43,
	0x68, 0x61, 0x69, 0x6e, 0x54, 0x78, 0x49, 0x44, 0x18, 0x6e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x43, 0x68, 0x61, 0x69, 0x6e, 0x54, 0x78, 0x49, 0x44, 0x12, 0x32, 0x0a, 0x05, 0x53, 0x74, 0x61,
	0x74, 0x65, 0x18, 0x78, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1c, 0x2e, 0x63, 0x68, 0x61, 0x69, 0x6e,
	0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x74, 0x78, 0x2e, 0x76, 0x31, 0x2e, 0x54,
	0x78, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x1b, 0x0a,
	0x08, 0x53, 0x74, 0x61, 0x74, 0x65, 0x53, 0x74, 0x72, 0x18, 0x82, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x53, 0x74, 0x61, 0x74, 0x65, 0x53, 0x74, 0x72, 0x12, 0x15, 0x0a, 0x05, 0x45, 0x78,
	0x74, 0x72, 0x61, 0x18, 0x8c, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x45, 0x78, 0x74, 0x72,
	0x61, 0x12, 0x30, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x18, 0x96, 0x01, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x1b, 0x2e, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72,
	0x2e, 0x74, 0x78, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x78, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x19, 0x0a, 0x07, 0x54, 0x79, 0x70, 0x65, 0x53, 0x74, 0x72, 0x18, 0xa0,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x54, 0x79, 0x70, 0x65, 0x53, 0x74, 0x72, 0x12, 0x1d,
	0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0xaa, 0x01, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a,
	0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0xb4, 0x01, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x41, 0x0a, 0x0f,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x78, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x2e, 0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x74, 0x78,
	0x2e, 0x76, 0x31, 0x2e, 0x54, 0x78, 0x52, 0x65, 0x71, 0x52, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x22,
	0x42, 0x0a, 0x10, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x78, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x2e, 0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65,
	0x77, 0x61, 0x72, 0x65, 0x2e, 0x74, 0x78, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x78, 0x52, 0x04, 0x49,
	0x6e, 0x66, 0x6f, 0x22, 0x44, 0x0a, 0x10, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x78, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x30, 0x0a, 0x05, 0x49, 0x6e, 0x66, 0x6f, 0x73,
	0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x6d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x74, 0x78, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x78, 0x52,
	0x65, 0x71, 0x52, 0x05, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x22, 0x45, 0x0a, 0x11, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x54, 0x78, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x30,
	0x0a, 0x05, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65,
	0x2e, 0x74, 0x78, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x78, 0x52, 0x05, 0x49, 0x6e, 0x66, 0x6f, 0x73,
	0x22, 0x1e, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x54, 0x78, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44,
	0x22, 0x3f, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x54, 0x78, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x2e, 0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61,
	0x72, 0x65, 0x2e, 0x74, 0x78, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x78, 0x52, 0x04, 0x49, 0x6e, 0x66,
	0x6f, 0x22, 0x6f, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x54, 0x78, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x30, 0x0a, 0x05, 0x43, 0x6f, 0x6e, 0x64, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65,
	0x72, 0x2e, 0x74, 0x78, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x64, 0x73, 0x52, 0x05, 0x43,
	0x6f, 0x6e, 0x64, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x14,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05,
	0x4c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x4c, 0x69, 0x6d,
	0x69, 0x74, 0x22, 0x58, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x54, 0x78, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x30, 0x0a, 0x05, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x18, 0x0a, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x6d, 0x69, 0x64, 0x64,
	0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x74, 0x78, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x78, 0x52,
	0x05, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x18,
	0x14, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x22, 0x41, 0x0a, 0x0f,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x78, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x2e, 0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x74, 0x78,
	0x2e, 0x76, 0x31, 0x2e, 0x54, 0x78, 0x52, 0x65, 0x71, 0x52, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x22,
	0x42, 0x0a, 0x10, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x78, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x2e, 0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65,
	0x77, 0x61, 0x72, 0x65, 0x2e, 0x74, 0x78, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x78, 0x52, 0x04, 0x49,
	0x6e, 0x66, 0x6f, 0x32, 0xe5, 0x03, 0x0a, 0x0a, 0x4d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61,
	0x72, 0x65, 0x12, 0x5f, 0x0a, 0x08, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x78, 0x12, 0x27,
	0x2e, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72,
	0x65, 0x2e, 0x74, 0x78, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x78,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e,
	0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x74, 0x78, 0x2e, 0x76, 0x31,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x78, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x62, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x78, 0x73,
	0x12, 0x28, 0x2e, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77,
	0x61, 0x72, 0x65, 0x2e, 0x74, 0x78, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x54, 0x78, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x63, 0x68, 0x61,
	0x69, 0x6e, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x74, 0x78,
	0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x78, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x56, 0x0a, 0x05, 0x47, 0x65, 0x74, 0x54, 0x78,
	0x12, 0x24, 0x2e, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77,
	0x61, 0x72, 0x65, 0x2e, 0x74, 0x78, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x78, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x6d,
	0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x74, 0x78, 0x2e, 0x76, 0x31, 0x2e,
	0x47, 0x65, 0x74, 0x54, 0x78, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x59, 0x0a, 0x06, 0x47, 0x65, 0x74, 0x54, 0x78, 0x73, 0x12, 0x25, 0x2e, 0x63, 0x68, 0x61, 0x69,
	0x6e, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x74, 0x78, 0x2e,
	0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x78, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x26, 0x2e, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77,
	0x61, 0x72, 0x65, 0x2e, 0x74, 0x78, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x78, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x5f, 0x0a, 0x08, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x54, 0x78, 0x12, 0x27, 0x2e, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x6d,
	0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x74, 0x78, 0x2e, 0x76, 0x31, 0x2e,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x78, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x28, 0x2e, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61,
	0x72, 0x65, 0x2e, 0x74, 0x78, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54,
	0x78, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x37, 0x5a, 0x35, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4e, 0x70, 0x6f, 0x6f, 0x6c, 0x50,
	0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f,
	0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2f, 0x6d, 0x77, 0x2f, 0x76,
	0x31, 0x2f, 0x74, 0x78, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_npool_chain_mw_v1_tx_tx_proto_rawDescOnce sync.Once
	file_npool_chain_mw_v1_tx_tx_proto_rawDescData = file_npool_chain_mw_v1_tx_tx_proto_rawDesc
)

func file_npool_chain_mw_v1_tx_tx_proto_rawDescGZIP() []byte {
	file_npool_chain_mw_v1_tx_tx_proto_rawDescOnce.Do(func() {
		file_npool_chain_mw_v1_tx_tx_proto_rawDescData = protoimpl.X.CompressGZIP(file_npool_chain_mw_v1_tx_tx_proto_rawDescData)
	})
	return file_npool_chain_mw_v1_tx_tx_proto_rawDescData
}

var file_npool_chain_mw_v1_tx_tx_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_npool_chain_mw_v1_tx_tx_proto_goTypes = []interface{}{
	(*Tx)(nil),                // 0: chain.middleware.tx.v1.Tx
	(*CreateTxRequest)(nil),   // 1: chain.middleware.tx.v1.CreateTxRequest
	(*CreateTxResponse)(nil),  // 2: chain.middleware.tx.v1.CreateTxResponse
	(*CreateTxsRequest)(nil),  // 3: chain.middleware.tx.v1.CreateTxsRequest
	(*CreateTxsResponse)(nil), // 4: chain.middleware.tx.v1.CreateTxsResponse
	(*GetTxRequest)(nil),      // 5: chain.middleware.tx.v1.GetTxRequest
	(*GetTxResponse)(nil),     // 6: chain.middleware.tx.v1.GetTxResponse
	(*GetTxsRequest)(nil),     // 7: chain.middleware.tx.v1.GetTxsRequest
	(*GetTxsResponse)(nil),    // 8: chain.middleware.tx.v1.GetTxsResponse
	(*UpdateTxRequest)(nil),   // 9: chain.middleware.tx.v1.UpdateTxRequest
	(*UpdateTxResponse)(nil),  // 10: chain.middleware.tx.v1.UpdateTxResponse
	(tx.TxState)(0),           // 11: chain.manager.tx.v1.TxState
	(tx.TxType)(0),            // 12: chain.manager.tx.v1.TxType
	(*tx.TxReq)(nil),          // 13: chain.manager.tx.v1.TxReq
	(*tx.Conds)(nil),          // 14: chain.manager.tx.v1.Conds
}
var file_npool_chain_mw_v1_tx_tx_proto_depIdxs = []int32{
	11, // 0: chain.middleware.tx.v1.Tx.State:type_name -> chain.manager.tx.v1.TxState
	12, // 1: chain.middleware.tx.v1.Tx.Type:type_name -> chain.manager.tx.v1.TxType
	13, // 2: chain.middleware.tx.v1.CreateTxRequest.Info:type_name -> chain.manager.tx.v1.TxReq
	0,  // 3: chain.middleware.tx.v1.CreateTxResponse.Info:type_name -> chain.middleware.tx.v1.Tx
	13, // 4: chain.middleware.tx.v1.CreateTxsRequest.Infos:type_name -> chain.manager.tx.v1.TxReq
	0,  // 5: chain.middleware.tx.v1.CreateTxsResponse.Infos:type_name -> chain.middleware.tx.v1.Tx
	0,  // 6: chain.middleware.tx.v1.GetTxResponse.Info:type_name -> chain.middleware.tx.v1.Tx
	14, // 7: chain.middleware.tx.v1.GetTxsRequest.Conds:type_name -> chain.manager.tx.v1.Conds
	0,  // 8: chain.middleware.tx.v1.GetTxsResponse.Infos:type_name -> chain.middleware.tx.v1.Tx
	13, // 9: chain.middleware.tx.v1.UpdateTxRequest.Info:type_name -> chain.manager.tx.v1.TxReq
	0,  // 10: chain.middleware.tx.v1.UpdateTxResponse.Info:type_name -> chain.middleware.tx.v1.Tx
	1,  // 11: chain.middleware.tx.v1.Middleware.CreateTx:input_type -> chain.middleware.tx.v1.CreateTxRequest
	3,  // 12: chain.middleware.tx.v1.Middleware.CreateTxs:input_type -> chain.middleware.tx.v1.CreateTxsRequest
	5,  // 13: chain.middleware.tx.v1.Middleware.GetTx:input_type -> chain.middleware.tx.v1.GetTxRequest
	7,  // 14: chain.middleware.tx.v1.Middleware.GetTxs:input_type -> chain.middleware.tx.v1.GetTxsRequest
	9,  // 15: chain.middleware.tx.v1.Middleware.UpdateTx:input_type -> chain.middleware.tx.v1.UpdateTxRequest
	2,  // 16: chain.middleware.tx.v1.Middleware.CreateTx:output_type -> chain.middleware.tx.v1.CreateTxResponse
	4,  // 17: chain.middleware.tx.v1.Middleware.CreateTxs:output_type -> chain.middleware.tx.v1.CreateTxsResponse
	6,  // 18: chain.middleware.tx.v1.Middleware.GetTx:output_type -> chain.middleware.tx.v1.GetTxResponse
	8,  // 19: chain.middleware.tx.v1.Middleware.GetTxs:output_type -> chain.middleware.tx.v1.GetTxsResponse
	10, // 20: chain.middleware.tx.v1.Middleware.UpdateTx:output_type -> chain.middleware.tx.v1.UpdateTxResponse
	16, // [16:21] is the sub-list for method output_type
	11, // [11:16] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_npool_chain_mw_v1_tx_tx_proto_init() }
func file_npool_chain_mw_v1_tx_tx_proto_init() {
	if File_npool_chain_mw_v1_tx_tx_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_npool_chain_mw_v1_tx_tx_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_npool_chain_mw_v1_tx_tx_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateTxRequest); i {
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
		file_npool_chain_mw_v1_tx_tx_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateTxResponse); i {
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
		file_npool_chain_mw_v1_tx_tx_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateTxsRequest); i {
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
		file_npool_chain_mw_v1_tx_tx_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateTxsResponse); i {
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
		file_npool_chain_mw_v1_tx_tx_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTxRequest); i {
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
		file_npool_chain_mw_v1_tx_tx_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTxResponse); i {
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
		file_npool_chain_mw_v1_tx_tx_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
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
		file_npool_chain_mw_v1_tx_tx_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
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
		file_npool_chain_mw_v1_tx_tx_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateTxRequest); i {
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
		file_npool_chain_mw_v1_tx_tx_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateTxResponse); i {
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
			RawDescriptor: file_npool_chain_mw_v1_tx_tx_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_npool_chain_mw_v1_tx_tx_proto_goTypes,
		DependencyIndexes: file_npool_chain_mw_v1_tx_tx_proto_depIdxs,
		MessageInfos:      file_npool_chain_mw_v1_tx_tx_proto_msgTypes,
	}.Build()
	File_npool_chain_mw_v1_tx_tx_proto = out.File
	file_npool_chain_mw_v1_tx_tx_proto_rawDesc = nil
	file_npool_chain_mw_v1_tx_tx_proto_goTypes = nil
	file_npool_chain_mw_v1_tx_tx_proto_depIdxs = nil
}
