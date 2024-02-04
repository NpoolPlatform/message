// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.18.1
// source: npool/ledger/mw/v2/simulate/ledger/profit/profit.proto

package profit

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

type ProfitReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID         *uint32 `protobuf:"varint,9,opt,name=ID,proto3,oneof" json:"ID,omitempty"`
	EntID      *string `protobuf:"bytes,10,opt,name=EntID,proto3,oneof" json:"EntID,omitempty"`
	AppID      *string `protobuf:"bytes,20,opt,name=AppID,proto3,oneof" json:"AppID,omitempty"`
	UserID     *string `protobuf:"bytes,30,opt,name=UserID,proto3,oneof" json:"UserID,omitempty"`
	CoinTypeID *string `protobuf:"bytes,40,opt,name=CoinTypeID,proto3,oneof" json:"CoinTypeID,omitempty"`
	Incoming   *string `protobuf:"bytes,50,opt,name=Incoming,proto3,oneof" json:"Incoming,omitempty"`
}

func (x *ProfitReq) Reset() {
	*x = ProfitReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_ledger_mw_v2_simulate_ledger_profit_profit_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProfitReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProfitReq) ProtoMessage() {}

func (x *ProfitReq) ProtoReflect() protoreflect.Message {
	mi := &file_npool_ledger_mw_v2_simulate_ledger_profit_profit_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProfitReq.ProtoReflect.Descriptor instead.
func (*ProfitReq) Descriptor() ([]byte, []int) {
	return file_npool_ledger_mw_v2_simulate_ledger_profit_profit_proto_rawDescGZIP(), []int{0}
}

func (x *ProfitReq) GetID() uint32 {
	if x != nil && x.ID != nil {
		return *x.ID
	}
	return 0
}

func (x *ProfitReq) GetEntID() string {
	if x != nil && x.EntID != nil {
		return *x.EntID
	}
	return ""
}

func (x *ProfitReq) GetAppID() string {
	if x != nil && x.AppID != nil {
		return *x.AppID
	}
	return ""
}

func (x *ProfitReq) GetUserID() string {
	if x != nil && x.UserID != nil {
		return *x.UserID
	}
	return ""
}

func (x *ProfitReq) GetCoinTypeID() string {
	if x != nil && x.CoinTypeID != nil {
		return *x.CoinTypeID
	}
	return ""
}

func (x *ProfitReq) GetIncoming() string {
	if x != nil && x.Incoming != nil {
		return *x.Incoming
	}
	return ""
}

type Profit struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: sql:"id"
	ID uint32 `protobuf:"varint,9,opt,name=ID,proto3" json:"ID,omitempty" sql:"id"`
	// @inject_tag: sql:"ent_id"
	EntID string `protobuf:"bytes,10,opt,name=EntID,proto3" json:"EntID,omitempty" sql:"ent_id"`
	// @inject_tag: sql:"app_id"
	AppID string `protobuf:"bytes,20,opt,name=AppID,proto3" json:"AppID,omitempty" sql:"app_id"`
	// @inject_tag: sql:"user_id"
	UserID string `protobuf:"bytes,30,opt,name=UserID,proto3" json:"UserID,omitempty" sql:"user_id"`
	// @inject_tag: sql:"coin_type_id"
	CoinTypeID string `protobuf:"bytes,40,opt,name=CoinTypeID,proto3" json:"CoinTypeID,omitempty" sql:"coin_type_id"`
	// @inject_tag: sql:"incoming"
	Incoming string `protobuf:"bytes,50,opt,name=Incoming,proto3" json:"Incoming,omitempty" sql:"incoming"`
	// @inject_tag: sql:"created_at"
	CreatedAt uint32 `protobuf:"varint,1000,opt,name=CreatedAt,proto3" json:"CreatedAt,omitempty" sql:"created_at"`
	// @inject_tag: sql:"updated_at"
	UpdatedAt uint32 `protobuf:"varint,1010,opt,name=UpdatedAt,proto3" json:"UpdatedAt,omitempty" sql:"updated_at"`
}

func (x *Profit) Reset() {
	*x = Profit{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_ledger_mw_v2_simulate_ledger_profit_profit_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Profit) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Profit) ProtoMessage() {}

func (x *Profit) ProtoReflect() protoreflect.Message {
	mi := &file_npool_ledger_mw_v2_simulate_ledger_profit_profit_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Profit.ProtoReflect.Descriptor instead.
func (*Profit) Descriptor() ([]byte, []int) {
	return file_npool_ledger_mw_v2_simulate_ledger_profit_profit_proto_rawDescGZIP(), []int{1}
}

func (x *Profit) GetID() uint32 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *Profit) GetEntID() string {
	if x != nil {
		return x.EntID
	}
	return ""
}

func (x *Profit) GetAppID() string {
	if x != nil {
		return x.AppID
	}
	return ""
}

func (x *Profit) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

func (x *Profit) GetCoinTypeID() string {
	if x != nil {
		return x.CoinTypeID
	}
	return ""
}

func (x *Profit) GetIncoming() string {
	if x != nil {
		return x.Incoming
	}
	return ""
}

func (x *Profit) GetCreatedAt() uint32 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *Profit) GetUpdatedAt() uint32 {
	if x != nil {
		return x.UpdatedAt
	}
	return 0
}

type Conds struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EntID      *v1.StringVal `protobuf:"bytes,10,opt,name=EntID,proto3,oneof" json:"EntID,omitempty"`
	AppID      *v1.StringVal `protobuf:"bytes,20,opt,name=AppID,proto3,oneof" json:"AppID,omitempty"`
	UserID     *v1.StringVal `protobuf:"bytes,30,opt,name=UserID,proto3,oneof" json:"UserID,omitempty"`
	CoinTypeID *v1.StringVal `protobuf:"bytes,40,opt,name=CoinTypeID,proto3,oneof" json:"CoinTypeID,omitempty"`
	Incoming   *v1.StringVal `protobuf:"bytes,50,opt,name=Incoming,proto3,oneof" json:"Incoming,omitempty"`
}

func (x *Conds) Reset() {
	*x = Conds{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_ledger_mw_v2_simulate_ledger_profit_profit_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Conds) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Conds) ProtoMessage() {}

func (x *Conds) ProtoReflect() protoreflect.Message {
	mi := &file_npool_ledger_mw_v2_simulate_ledger_profit_profit_proto_msgTypes[2]
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
	return file_npool_ledger_mw_v2_simulate_ledger_profit_profit_proto_rawDescGZIP(), []int{2}
}

func (x *Conds) GetEntID() *v1.StringVal {
	if x != nil {
		return x.EntID
	}
	return nil
}

func (x *Conds) GetAppID() *v1.StringVal {
	if x != nil {
		return x.AppID
	}
	return nil
}

func (x *Conds) GetUserID() *v1.StringVal {
	if x != nil {
		return x.UserID
	}
	return nil
}

func (x *Conds) GetCoinTypeID() *v1.StringVal {
	if x != nil {
		return x.CoinTypeID
	}
	return nil
}

func (x *Conds) GetIncoming() *v1.StringVal {
	if x != nil {
		return x.Incoming
	}
	return nil
}

type GetProfitRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EntID string `protobuf:"bytes,10,opt,name=EntID,proto3" json:"EntID,omitempty"`
}

func (x *GetProfitRequest) Reset() {
	*x = GetProfitRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_ledger_mw_v2_simulate_ledger_profit_profit_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetProfitRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProfitRequest) ProtoMessage() {}

func (x *GetProfitRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_ledger_mw_v2_simulate_ledger_profit_profit_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProfitRequest.ProtoReflect.Descriptor instead.
func (*GetProfitRequest) Descriptor() ([]byte, []int) {
	return file_npool_ledger_mw_v2_simulate_ledger_profit_profit_proto_rawDescGZIP(), []int{3}
}

func (x *GetProfitRequest) GetEntID() string {
	if x != nil {
		return x.EntID
	}
	return ""
}

type GetProfitResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *Profit `protobuf:"bytes,10,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (x *GetProfitResponse) Reset() {
	*x = GetProfitResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_ledger_mw_v2_simulate_ledger_profit_profit_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetProfitResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProfitResponse) ProtoMessage() {}

func (x *GetProfitResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_ledger_mw_v2_simulate_ledger_profit_profit_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProfitResponse.ProtoReflect.Descriptor instead.
func (*GetProfitResponse) Descriptor() ([]byte, []int) {
	return file_npool_ledger_mw_v2_simulate_ledger_profit_profit_proto_rawDescGZIP(), []int{4}
}

func (x *GetProfitResponse) GetInfo() *Profit {
	if x != nil {
		return x.Info
	}
	return nil
}

type GetProfitsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Conds  *Conds `protobuf:"bytes,10,opt,name=Conds,proto3" json:"Conds,omitempty"`
	Offset int32  `protobuf:"varint,20,opt,name=Offset,proto3" json:"Offset,omitempty"`
	Limit  int32  `protobuf:"varint,30,opt,name=Limit,proto3" json:"Limit,omitempty"`
}

func (x *GetProfitsRequest) Reset() {
	*x = GetProfitsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_ledger_mw_v2_simulate_ledger_profit_profit_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetProfitsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProfitsRequest) ProtoMessage() {}

func (x *GetProfitsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_ledger_mw_v2_simulate_ledger_profit_profit_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProfitsRequest.ProtoReflect.Descriptor instead.
func (*GetProfitsRequest) Descriptor() ([]byte, []int) {
	return file_npool_ledger_mw_v2_simulate_ledger_profit_profit_proto_rawDescGZIP(), []int{5}
}

func (x *GetProfitsRequest) GetConds() *Conds {
	if x != nil {
		return x.Conds
	}
	return nil
}

func (x *GetProfitsRequest) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *GetProfitsRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type GetProfitsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Infos []*Profit `protobuf:"bytes,10,rep,name=Infos,proto3" json:"Infos,omitempty"`
	Total uint32    `protobuf:"varint,20,opt,name=Total,proto3" json:"Total,omitempty"`
}

func (x *GetProfitsResponse) Reset() {
	*x = GetProfitsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_ledger_mw_v2_simulate_ledger_profit_profit_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetProfitsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProfitsResponse) ProtoMessage() {}

func (x *GetProfitsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_ledger_mw_v2_simulate_ledger_profit_profit_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProfitsResponse.ProtoReflect.Descriptor instead.
func (*GetProfitsResponse) Descriptor() ([]byte, []int) {
	return file_npool_ledger_mw_v2_simulate_ledger_profit_profit_proto_rawDescGZIP(), []int{6}
}

func (x *GetProfitsResponse) GetInfos() []*Profit {
	if x != nil {
		return x.Infos
	}
	return nil
}

func (x *GetProfitsResponse) GetTotal() uint32 {
	if x != nil {
		return x.Total
	}
	return 0
}

type ExistProfitCondsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Conds *Conds `protobuf:"bytes,10,opt,name=Conds,proto3" json:"Conds,omitempty"`
}

func (x *ExistProfitCondsRequest) Reset() {
	*x = ExistProfitCondsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_ledger_mw_v2_simulate_ledger_profit_profit_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExistProfitCondsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExistProfitCondsRequest) ProtoMessage() {}

func (x *ExistProfitCondsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_ledger_mw_v2_simulate_ledger_profit_profit_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExistProfitCondsRequest.ProtoReflect.Descriptor instead.
func (*ExistProfitCondsRequest) Descriptor() ([]byte, []int) {
	return file_npool_ledger_mw_v2_simulate_ledger_profit_profit_proto_rawDescGZIP(), []int{7}
}

func (x *ExistProfitCondsRequest) GetConds() *Conds {
	if x != nil {
		return x.Conds
	}
	return nil
}

type ExistProfitCondsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info bool `protobuf:"varint,10,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (x *ExistProfitCondsResponse) Reset() {
	*x = ExistProfitCondsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_ledger_mw_v2_simulate_ledger_profit_profit_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExistProfitCondsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExistProfitCondsResponse) ProtoMessage() {}

func (x *ExistProfitCondsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_ledger_mw_v2_simulate_ledger_profit_profit_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExistProfitCondsResponse.ProtoReflect.Descriptor instead.
func (*ExistProfitCondsResponse) Descriptor() ([]byte, []int) {
	return file_npool_ledger_mw_v2_simulate_ledger_profit_profit_proto_rawDescGZIP(), []int{8}
}

func (x *ExistProfitCondsResponse) GetInfo() bool {
	if x != nil {
		return x.Info
	}
	return false
}

var File_npool_ledger_mw_v2_simulate_ledger_profit_profit_proto protoreflect.FileDescriptor

var file_npool_ledger_mw_v2_simulate_ledger_profit_profit_proto_rawDesc = []byte{
	0x0a, 0x36, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2f, 0x6d,
	0x77, 0x2f, 0x76, 0x32, 0x2f, 0x73, 0x69, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x2f, 0x6c, 0x65,
	0x64, 0x67, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x66,
	0x69, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x2b, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72,
	0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x73, 0x69, 0x6d, 0x75,
	0x6c, 0x61, 0x74, 0x65, 0x2e, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x66,
	0x69, 0x74, 0x2e, 0x76, 0x32, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xfb, 0x01, 0x0a, 0x09, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x74, 0x52, 0x65,
	0x71, 0x12, 0x13, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0d, 0x48, 0x00, 0x52,
	0x02, 0x49, 0x44, 0x88, 0x01, 0x01, 0x12, 0x19, 0x0a, 0x05, 0x45, 0x6e, 0x74, 0x49, 0x44, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x05, 0x45, 0x6e, 0x74, 0x49, 0x44, 0x88, 0x01,
	0x01, 0x12, 0x19, 0x0a, 0x05, 0x41, 0x70, 0x70, 0x49, 0x44, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09,
	0x48, 0x02, 0x52, 0x05, 0x41, 0x70, 0x70, 0x49, 0x44, 0x88, 0x01, 0x01, 0x12, 0x1b, 0x0a, 0x06,
	0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x09, 0x48, 0x03, 0x52, 0x06,
	0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x88, 0x01, 0x01, 0x12, 0x23, 0x0a, 0x0a, 0x43, 0x6f, 0x69,
	0x6e, 0x54, 0x79, 0x70, 0x65, 0x49, 0x44, 0x18, 0x28, 0x20, 0x01, 0x28, 0x09, 0x48, 0x04, 0x52,
	0x0a, 0x43, 0x6f, 0x69, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x49, 0x44, 0x88, 0x01, 0x01, 0x12, 0x1f,
	0x0a, 0x08, 0x49, 0x6e, 0x63, 0x6f, 0x6d, 0x69, 0x6e, 0x67, 0x18, 0x32, 0x20, 0x01, 0x28, 0x09,
	0x48, 0x05, 0x52, 0x08, 0x49, 0x6e, 0x63, 0x6f, 0x6d, 0x69, 0x6e, 0x67, 0x88, 0x01, 0x01, 0x42,
	0x05, 0x0a, 0x03, 0x5f, 0x49, 0x44, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x45, 0x6e, 0x74, 0x49, 0x44,
	0x42, 0x08, 0x0a, 0x06, 0x5f, 0x41, 0x70, 0x70, 0x49, 0x44, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x55,
	0x73, 0x65, 0x72, 0x49, 0x44, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x43, 0x6f, 0x69, 0x6e, 0x54, 0x79,
	0x70, 0x65, 0x49, 0x44, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x49, 0x6e, 0x63, 0x6f, 0x6d, 0x69, 0x6e,
	0x67, 0x22, 0xd6, 0x01, 0x0a, 0x06, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x49, 0x44, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05,
	0x45, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x45, 0x6e, 0x74,
	0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x41, 0x70, 0x70, 0x49, 0x44, 0x18, 0x14, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x41, 0x70, 0x70, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72,
	0x49, 0x44, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44,
	0x12, 0x1e, 0x0a, 0x0a, 0x43, 0x6f, 0x69, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x49, 0x44, 0x18, 0x28,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x43, 0x6f, 0x69, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x49, 0x44,
	0x12, 0x1a, 0x0a, 0x08, 0x49, 0x6e, 0x63, 0x6f, 0x6d, 0x69, 0x6e, 0x67, 0x18, 0x32, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x49, 0x6e, 0x63, 0x6f, 0x6d, 0x69, 0x6e, 0x67, 0x12, 0x1d, 0x0a, 0x09,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0xe8, 0x07, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x09, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0xf2, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0xd8, 0x02, 0x0a, 0x05, 0x43,
	0x6f, 0x6e, 0x64, 0x73, 0x12, 0x32, 0x0a, 0x05, 0x45, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x48, 0x00, 0x52, 0x05,
	0x45, 0x6e, 0x74, 0x49, 0x44, 0x88, 0x01, 0x01, 0x12, 0x32, 0x0a, 0x05, 0x41, 0x70, 0x70, 0x49,
	0x44, 0x18, 0x14, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c,
	0x48, 0x01, 0x52, 0x05, 0x41, 0x70, 0x70, 0x49, 0x44, 0x88, 0x01, 0x01, 0x12, 0x34, 0x0a, 0x06,
	0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x62,
	0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x72, 0x69,
	0x6e, 0x67, 0x56, 0x61, 0x6c, 0x48, 0x02, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x88,
	0x01, 0x01, 0x12, 0x3c, 0x0a, 0x0a, 0x43, 0x6f, 0x69, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x49, 0x44,
	0x18, 0x28, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70,
	0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x48,
	0x03, 0x52, 0x0a, 0x43, 0x6f, 0x69, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x49, 0x44, 0x88, 0x01, 0x01,
	0x12, 0x38, 0x0a, 0x08, 0x49, 0x6e, 0x63, 0x6f, 0x6d, 0x69, 0x6e, 0x67, 0x18, 0x32, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x17, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x48, 0x04, 0x52, 0x08, 0x49,
	0x6e, 0x63, 0x6f, 0x6d, 0x69, 0x6e, 0x67, 0x88, 0x01, 0x01, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x45,
	0x6e, 0x74, 0x49, 0x44, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x41, 0x70, 0x70, 0x49, 0x44, 0x42, 0x09,
	0x0a, 0x07, 0x5f, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x43, 0x6f,
	0x69, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x49, 0x44, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x49, 0x6e, 0x63,
	0x6f, 0x6d, 0x69, 0x6e, 0x67, 0x22, 0x28, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x66,
	0x69, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x45, 0x6e, 0x74,
	0x49, 0x44, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x45, 0x6e, 0x74, 0x49, 0x44, 0x22,
	0x5c, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x47, 0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x33, 0x2e, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2e, 0x6d, 0x69, 0x64, 0x64,
	0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x73, 0x69, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x2e,
	0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x74, 0x2e, 0x76, 0x32,
	0x2e, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x74, 0x52, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x8b, 0x01,
	0x0a, 0x11, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x48, 0x0a, 0x05, 0x43, 0x6f, 0x6e, 0x64, 0x73, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x32, 0x2e, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2e, 0x6d, 0x69, 0x64, 0x64,
	0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x73, 0x69, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x2e,
	0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x74, 0x2e, 0x76, 0x32,
	0x2e, 0x43, 0x6f, 0x6e, 0x64, 0x73, 0x52, 0x05, 0x43, 0x6f, 0x6e, 0x64, 0x73, 0x12, 0x16, 0x0a,
	0x06, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x14, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x4f,
	0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x1e,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x22, 0x75, 0x0a, 0x12, 0x47,
	0x65, 0x74, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x49, 0x0a, 0x05, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x33, 0x2e, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65,
	0x77, 0x61, 0x72, 0x65, 0x2e, 0x73, 0x69, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x2e, 0x6c, 0x65,
	0x64, 0x67, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x74, 0x2e, 0x76, 0x32, 0x2e, 0x50,
	0x72, 0x6f, 0x66, 0x69, 0x74, 0x52, 0x05, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x12, 0x14, 0x0a, 0x05,
	0x54, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x14, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x54, 0x6f, 0x74,
	0x61, 0x6c, 0x22, 0x63, 0x0a, 0x17, 0x45, 0x78, 0x69, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x66, 0x69,
	0x74, 0x43, 0x6f, 0x6e, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x48, 0x0a,
	0x05, 0x43, 0x6f, 0x6e, 0x64, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x32, 0x2e, 0x6c,
	0x65, 0x64, 0x67, 0x65, 0x72, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65,
	0x2e, 0x73, 0x69, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x2e, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x74, 0x2e, 0x76, 0x32, 0x2e, 0x43, 0x6f, 0x6e, 0x64, 0x73,
	0x52, 0x05, 0x43, 0x6f, 0x6e, 0x64, 0x73, 0x22, 0x2e, 0x0a, 0x18, 0x45, 0x78, 0x69, 0x73, 0x74,
	0x50, 0x72, 0x6f, 0x66, 0x69, 0x74, 0x43, 0x6f, 0x6e, 0x64, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x32, 0xd1, 0x04, 0x0a, 0x0a, 0x4d, 0x69, 0x64, 0x64,
	0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x12, 0xb4, 0x01, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x50, 0x72,
	0x6f, 0x66, 0x69, 0x74, 0x12, 0x3d, 0x2e, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2e, 0x6d, 0x69,
	0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x73, 0x69, 0x6d, 0x75, 0x6c, 0x61, 0x74,
	0x65, 0x2e, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x74, 0x2e,
	0x76, 0x32, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x3e, 0x2e, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2e, 0x6d, 0x69, 0x64,
	0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x73, 0x69, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x65,
	0x2e, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x74, 0x2e, 0x76,
	0x32, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x28, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x22, 0x22, 0x1d, 0x2f, 0x76, 0x31,
	0x2f, 0x73, 0x69, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x74,
	0x5f, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x74, 0x3a, 0x01, 0x2a, 0x12, 0xb8, 0x01,
	0x0a, 0x0a, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x74, 0x73, 0x12, 0x3e, 0x2e, 0x6c,
	0x65, 0x64, 0x67, 0x65, 0x72, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65,
	0x2e, 0x73, 0x69, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x2e, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x74, 0x2e, 0x76, 0x32, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x72,
	0x6f, 0x66, 0x69, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3f, 0x2e, 0x6c,
	0x65, 0x64, 0x67, 0x65, 0x72, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65,
	0x2e, 0x73, 0x69, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x2e, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x74, 0x2e, 0x76, 0x32, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x72,
	0x6f, 0x66, 0x69, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x29, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x23, 0x22, 0x1e, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x69, 0x6d, 0x75, 0x6c,
	0x61, 0x74, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x74, 0x5f, 0x47, 0x65, 0x74, 0x50, 0x72,
	0x6f, 0x66, 0x69, 0x74, 0x73, 0x3a, 0x01, 0x2a, 0x12, 0xd0, 0x01, 0x0a, 0x10, 0x45, 0x78, 0x69,
	0x73, 0x74, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x74, 0x43, 0x6f, 0x6e, 0x64, 0x73, 0x12, 0x44, 0x2e,
	0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72,
	0x65, 0x2e, 0x73, 0x69, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x2e, 0x6c, 0x65, 0x64, 0x67, 0x65,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x74, 0x2e, 0x76, 0x32, 0x2e, 0x45, 0x78, 0x69, 0x73,
	0x74, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x74, 0x43, 0x6f, 0x6e, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x45, 0x2e, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2e, 0x6d, 0x69, 0x64,
	0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x73, 0x69, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x65,
	0x2e, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x74, 0x2e, 0x76,
	0x32, 0x2e, 0x45, 0x78, 0x69, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x74, 0x43, 0x6f, 0x6e,
	0x64, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2f, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x29, 0x22, 0x24, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x69, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x74, 0x5f, 0x45, 0x78, 0x69, 0x73, 0x74, 0x50, 0x72, 0x6f,
	0x66, 0x69, 0x74, 0x43, 0x6f, 0x6e, 0x64, 0x73, 0x3a, 0x01, 0x2a, 0x42, 0x4c, 0x5a, 0x4a, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4e, 0x70, 0x6f, 0x6f, 0x6c, 0x50,
	0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f,
	0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2f, 0x6d, 0x77, 0x2f,
	0x76, 0x32, 0x2f, 0x73, 0x69, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x2f, 0x6c, 0x65, 0x64, 0x67,
	0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_npool_ledger_mw_v2_simulate_ledger_profit_profit_proto_rawDescOnce sync.Once
	file_npool_ledger_mw_v2_simulate_ledger_profit_profit_proto_rawDescData = file_npool_ledger_mw_v2_simulate_ledger_profit_profit_proto_rawDesc
)

func file_npool_ledger_mw_v2_simulate_ledger_profit_profit_proto_rawDescGZIP() []byte {
	file_npool_ledger_mw_v2_simulate_ledger_profit_profit_proto_rawDescOnce.Do(func() {
		file_npool_ledger_mw_v2_simulate_ledger_profit_profit_proto_rawDescData = protoimpl.X.CompressGZIP(file_npool_ledger_mw_v2_simulate_ledger_profit_profit_proto_rawDescData)
	})
	return file_npool_ledger_mw_v2_simulate_ledger_profit_profit_proto_rawDescData
}

var file_npool_ledger_mw_v2_simulate_ledger_profit_profit_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_npool_ledger_mw_v2_simulate_ledger_profit_profit_proto_goTypes = []interface{}{
	(*ProfitReq)(nil),                // 0: ledger.middleware.simulate.ledger.profit.v2.ProfitReq
	(*Profit)(nil),                   // 1: ledger.middleware.simulate.ledger.profit.v2.Profit
	(*Conds)(nil),                    // 2: ledger.middleware.simulate.ledger.profit.v2.Conds
	(*GetProfitRequest)(nil),         // 3: ledger.middleware.simulate.ledger.profit.v2.GetProfitRequest
	(*GetProfitResponse)(nil),        // 4: ledger.middleware.simulate.ledger.profit.v2.GetProfitResponse
	(*GetProfitsRequest)(nil),        // 5: ledger.middleware.simulate.ledger.profit.v2.GetProfitsRequest
	(*GetProfitsResponse)(nil),       // 6: ledger.middleware.simulate.ledger.profit.v2.GetProfitsResponse
	(*ExistProfitCondsRequest)(nil),  // 7: ledger.middleware.simulate.ledger.profit.v2.ExistProfitCondsRequest
	(*ExistProfitCondsResponse)(nil), // 8: ledger.middleware.simulate.ledger.profit.v2.ExistProfitCondsResponse
	(*v1.StringVal)(nil),             // 9: basetypes.v1.StringVal
}
var file_npool_ledger_mw_v2_simulate_ledger_profit_profit_proto_depIdxs = []int32{
	9,  // 0: ledger.middleware.simulate.ledger.profit.v2.Conds.EntID:type_name -> basetypes.v1.StringVal
	9,  // 1: ledger.middleware.simulate.ledger.profit.v2.Conds.AppID:type_name -> basetypes.v1.StringVal
	9,  // 2: ledger.middleware.simulate.ledger.profit.v2.Conds.UserID:type_name -> basetypes.v1.StringVal
	9,  // 3: ledger.middleware.simulate.ledger.profit.v2.Conds.CoinTypeID:type_name -> basetypes.v1.StringVal
	9,  // 4: ledger.middleware.simulate.ledger.profit.v2.Conds.Incoming:type_name -> basetypes.v1.StringVal
	1,  // 5: ledger.middleware.simulate.ledger.profit.v2.GetProfitResponse.Info:type_name -> ledger.middleware.simulate.ledger.profit.v2.Profit
	2,  // 6: ledger.middleware.simulate.ledger.profit.v2.GetProfitsRequest.Conds:type_name -> ledger.middleware.simulate.ledger.profit.v2.Conds
	1,  // 7: ledger.middleware.simulate.ledger.profit.v2.GetProfitsResponse.Infos:type_name -> ledger.middleware.simulate.ledger.profit.v2.Profit
	2,  // 8: ledger.middleware.simulate.ledger.profit.v2.ExistProfitCondsRequest.Conds:type_name -> ledger.middleware.simulate.ledger.profit.v2.Conds
	3,  // 9: ledger.middleware.simulate.ledger.profit.v2.Middleware.GetProfit:input_type -> ledger.middleware.simulate.ledger.profit.v2.GetProfitRequest
	5,  // 10: ledger.middleware.simulate.ledger.profit.v2.Middleware.GetProfits:input_type -> ledger.middleware.simulate.ledger.profit.v2.GetProfitsRequest
	7,  // 11: ledger.middleware.simulate.ledger.profit.v2.Middleware.ExistProfitConds:input_type -> ledger.middleware.simulate.ledger.profit.v2.ExistProfitCondsRequest
	4,  // 12: ledger.middleware.simulate.ledger.profit.v2.Middleware.GetProfit:output_type -> ledger.middleware.simulate.ledger.profit.v2.GetProfitResponse
	6,  // 13: ledger.middleware.simulate.ledger.profit.v2.Middleware.GetProfits:output_type -> ledger.middleware.simulate.ledger.profit.v2.GetProfitsResponse
	8,  // 14: ledger.middleware.simulate.ledger.profit.v2.Middleware.ExistProfitConds:output_type -> ledger.middleware.simulate.ledger.profit.v2.ExistProfitCondsResponse
	12, // [12:15] is the sub-list for method output_type
	9,  // [9:12] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_npool_ledger_mw_v2_simulate_ledger_profit_profit_proto_init() }
func file_npool_ledger_mw_v2_simulate_ledger_profit_profit_proto_init() {
	if File_npool_ledger_mw_v2_simulate_ledger_profit_profit_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_npool_ledger_mw_v2_simulate_ledger_profit_profit_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProfitReq); i {
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
		file_npool_ledger_mw_v2_simulate_ledger_profit_profit_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Profit); i {
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
		file_npool_ledger_mw_v2_simulate_ledger_profit_profit_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
		file_npool_ledger_mw_v2_simulate_ledger_profit_profit_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetProfitRequest); i {
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
		file_npool_ledger_mw_v2_simulate_ledger_profit_profit_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetProfitResponse); i {
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
		file_npool_ledger_mw_v2_simulate_ledger_profit_profit_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetProfitsRequest); i {
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
		file_npool_ledger_mw_v2_simulate_ledger_profit_profit_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetProfitsResponse); i {
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
		file_npool_ledger_mw_v2_simulate_ledger_profit_profit_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExistProfitCondsRequest); i {
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
		file_npool_ledger_mw_v2_simulate_ledger_profit_profit_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExistProfitCondsResponse); i {
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
	file_npool_ledger_mw_v2_simulate_ledger_profit_profit_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_npool_ledger_mw_v2_simulate_ledger_profit_profit_proto_msgTypes[2].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_npool_ledger_mw_v2_simulate_ledger_profit_profit_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_npool_ledger_mw_v2_simulate_ledger_profit_profit_proto_goTypes,
		DependencyIndexes: file_npool_ledger_mw_v2_simulate_ledger_profit_profit_proto_depIdxs,
		MessageInfos:      file_npool_ledger_mw_v2_simulate_ledger_profit_profit_proto_msgTypes,
	}.Build()
	File_npool_ledger_mw_v2_simulate_ledger_profit_profit_proto = out.File
	file_npool_ledger_mw_v2_simulate_ledger_profit_profit_proto_rawDesc = nil
	file_npool_ledger_mw_v2_simulate_ledger_profit_profit_proto_goTypes = nil
	file_npool_ledger_mw_v2_simulate_ledger_profit_profit_proto_depIdxs = nil
}
