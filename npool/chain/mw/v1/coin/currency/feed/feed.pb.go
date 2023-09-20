// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.18.1
// source: npool/chain/mw/v1/coin/currency/feed/feed.proto

package feed

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

type FeedReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID           *string              `protobuf:"bytes,10,opt,name=ID,proto3,oneof" json:"ID,omitempty"`
	CoinTypeID   *string              `protobuf:"bytes,20,opt,name=CoinTypeID,proto3,oneof" json:"CoinTypeID,omitempty"`
	FeedType     *v1.CurrencyFeedType `protobuf:"varint,30,opt,name=FeedType,proto3,enum=basetypes.v1.CurrencyFeedType,oneof" json:"FeedType,omitempty"`
	FeedCoinName *string              `protobuf:"bytes,40,opt,name=FeedCoinName,proto3,oneof" json:"FeedCoinName,omitempty"`
	Disabled     *bool                `protobuf:"varint,50,opt,name=Disabled,proto3,oneof" json:"Disabled,omitempty"`
}

func (x *FeedReq) Reset() {
	*x = FeedReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_chain_mw_v1_coin_currency_feed_feed_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FeedReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FeedReq) ProtoMessage() {}

func (x *FeedReq) ProtoReflect() protoreflect.Message {
	mi := &file_npool_chain_mw_v1_coin_currency_feed_feed_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FeedReq.ProtoReflect.Descriptor instead.
func (*FeedReq) Descriptor() ([]byte, []int) {
	return file_npool_chain_mw_v1_coin_currency_feed_feed_proto_rawDescGZIP(), []int{0}
}

func (x *FeedReq) GetID() string {
	if x != nil && x.ID != nil {
		return *x.ID
	}
	return ""
}

func (x *FeedReq) GetCoinTypeID() string {
	if x != nil && x.CoinTypeID != nil {
		return *x.CoinTypeID
	}
	return ""
}

func (x *FeedReq) GetFeedType() v1.CurrencyFeedType {
	if x != nil && x.FeedType != nil {
		return *x.FeedType
	}
	return v1.CurrencyFeedType(0)
}

func (x *FeedReq) GetFeedCoinName() string {
	if x != nil && x.FeedCoinName != nil {
		return *x.FeedCoinName
	}
	return ""
}

func (x *FeedReq) GetDisabled() bool {
	if x != nil && x.Disabled != nil {
		return *x.Disabled
	}
	return false
}

type Feed struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: sql:"id"
	ID string `protobuf:"bytes,10,opt,name=ID,proto3" json:"ID,omitempty" sql:"id"`
	// @inject_tag: sql:"coin_type_id"
	CoinTypeID string `protobuf:"bytes,20,opt,name=CoinTypeID,proto3" json:"CoinTypeID,omitempty" sql:"coin_type_id"`
	// @inject_tag: sql:"coin_name"
	CoinName string `protobuf:"bytes,30,opt,name=CoinName,proto3" json:"CoinName,omitempty" sql:"coin_name"`
	// @inject_tag: sql:"coin_unit"
	CoinUnit string `protobuf:"bytes,40,opt,name=CoinUnit,proto3" json:"CoinUnit,omitempty" sql:"coin_unit"`
	// @inject_tag: sql:"coin_logo"
	CoinLogo string `protobuf:"bytes,50,opt,name=CoinLogo,proto3" json:"CoinLogo,omitempty" sql:"coin_logo"`
	// @inject_tag: sql:"coin_env"
	CoinENV string `protobuf:"bytes,60,opt,name=CoinENV,proto3" json:"CoinENV,omitempty" sql:"coin_env"`
	// @inject_tag: sql:"feed_type"
	FeedTypeStr string              `protobuf:"bytes,70,opt,name=FeedTypeStr,proto3" json:"FeedTypeStr,omitempty" sql:"feed_type"`
	FeedType    v1.CurrencyFeedType `protobuf:"varint,80,opt,name=FeedType,proto3,enum=basetypes.v1.CurrencyFeedType" json:"FeedType,omitempty"`
	// @inject_tag: sql:"feed_coin_name"
	FeedCoinName string `protobuf:"bytes,90,opt,name=FeedCoinName,proto3" json:"FeedCoinName,omitempty" sql:"feed_coin_name"`
	// @inject_tag: sql:"disabled"
	Disabled bool `protobuf:"varint,100,opt,name=Disabled,proto3" json:"Disabled,omitempty" sql:"disabled"`
	// @inject_tag: sql:"created_at"
	CreatedAt uint32 `protobuf:"varint,1000,opt,name=CreatedAt,proto3" json:"CreatedAt,omitempty" sql:"created_at"`
	// @inject_tag: sql:"updated_at"
	UpdatedAt uint32 `protobuf:"varint,1010,opt,name=UpdatedAt,proto3" json:"UpdatedAt,omitempty" sql:"updated_at"`
}

func (x *Feed) Reset() {
	*x = Feed{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_chain_mw_v1_coin_currency_feed_feed_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Feed) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Feed) ProtoMessage() {}

func (x *Feed) ProtoReflect() protoreflect.Message {
	mi := &file_npool_chain_mw_v1_coin_currency_feed_feed_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Feed.ProtoReflect.Descriptor instead.
func (*Feed) Descriptor() ([]byte, []int) {
	return file_npool_chain_mw_v1_coin_currency_feed_feed_proto_rawDescGZIP(), []int{1}
}

func (x *Feed) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *Feed) GetCoinTypeID() string {
	if x != nil {
		return x.CoinTypeID
	}
	return ""
}

func (x *Feed) GetCoinName() string {
	if x != nil {
		return x.CoinName
	}
	return ""
}

func (x *Feed) GetCoinUnit() string {
	if x != nil {
		return x.CoinUnit
	}
	return ""
}

func (x *Feed) GetCoinLogo() string {
	if x != nil {
		return x.CoinLogo
	}
	return ""
}

func (x *Feed) GetCoinENV() string {
	if x != nil {
		return x.CoinENV
	}
	return ""
}

func (x *Feed) GetFeedTypeStr() string {
	if x != nil {
		return x.FeedTypeStr
	}
	return ""
}

func (x *Feed) GetFeedType() v1.CurrencyFeedType {
	if x != nil {
		return x.FeedType
	}
	return v1.CurrencyFeedType(0)
}

func (x *Feed) GetFeedCoinName() string {
	if x != nil {
		return x.FeedCoinName
	}
	return ""
}

func (x *Feed) GetDisabled() bool {
	if x != nil {
		return x.Disabled
	}
	return false
}

func (x *Feed) GetCreatedAt() uint32 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *Feed) GetUpdatedAt() uint32 {
	if x != nil {
		return x.UpdatedAt
	}
	return 0
}

type Conds struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID          *v1.StringVal      `protobuf:"bytes,10,opt,name=ID,proto3,oneof" json:"ID,omitempty"`
	CoinTypeID  *v1.StringVal      `protobuf:"bytes,20,opt,name=CoinTypeID,proto3,oneof" json:"CoinTypeID,omitempty"`
	CoinTypeIDs *v1.StringSliceVal `protobuf:"bytes,30,opt,name=CoinTypeIDs,proto3,oneof" json:"CoinTypeIDs,omitempty"`
	FeedType    *v1.Uint32Val      `protobuf:"bytes,40,opt,name=FeedType,proto3,oneof" json:"FeedType,omitempty"`
	Disabled    *v1.BoolVal        `protobuf:"bytes,50,opt,name=Disabled,proto3,oneof" json:"Disabled,omitempty"`
}

func (x *Conds) Reset() {
	*x = Conds{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_chain_mw_v1_coin_currency_feed_feed_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Conds) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Conds) ProtoMessage() {}

func (x *Conds) ProtoReflect() protoreflect.Message {
	mi := &file_npool_chain_mw_v1_coin_currency_feed_feed_proto_msgTypes[2]
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
	return file_npool_chain_mw_v1_coin_currency_feed_feed_proto_rawDescGZIP(), []int{2}
}

func (x *Conds) GetID() *v1.StringVal {
	if x != nil {
		return x.ID
	}
	return nil
}

func (x *Conds) GetCoinTypeID() *v1.StringVal {
	if x != nil {
		return x.CoinTypeID
	}
	return nil
}

func (x *Conds) GetCoinTypeIDs() *v1.StringSliceVal {
	if x != nil {
		return x.CoinTypeIDs
	}
	return nil
}

func (x *Conds) GetFeedType() *v1.Uint32Val {
	if x != nil {
		return x.FeedType
	}
	return nil
}

func (x *Conds) GetDisabled() *v1.BoolVal {
	if x != nil {
		return x.Disabled
	}
	return nil
}

type CreateFeedRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *FeedReq `protobuf:"bytes,10,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (x *CreateFeedRequest) Reset() {
	*x = CreateFeedRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_chain_mw_v1_coin_currency_feed_feed_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateFeedRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateFeedRequest) ProtoMessage() {}

func (x *CreateFeedRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_chain_mw_v1_coin_currency_feed_feed_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateFeedRequest.ProtoReflect.Descriptor instead.
func (*CreateFeedRequest) Descriptor() ([]byte, []int) {
	return file_npool_chain_mw_v1_coin_currency_feed_feed_proto_rawDescGZIP(), []int{3}
}

func (x *CreateFeedRequest) GetInfo() *FeedReq {
	if x != nil {
		return x.Info
	}
	return nil
}

type CreateFeedResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *Feed `protobuf:"bytes,10,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (x *CreateFeedResponse) Reset() {
	*x = CreateFeedResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_chain_mw_v1_coin_currency_feed_feed_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateFeedResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateFeedResponse) ProtoMessage() {}

func (x *CreateFeedResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_chain_mw_v1_coin_currency_feed_feed_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateFeedResponse.ProtoReflect.Descriptor instead.
func (*CreateFeedResponse) Descriptor() ([]byte, []int) {
	return file_npool_chain_mw_v1_coin_currency_feed_feed_proto_rawDescGZIP(), []int{4}
}

func (x *CreateFeedResponse) GetInfo() *Feed {
	if x != nil {
		return x.Info
	}
	return nil
}

type UpdateFeedRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *FeedReq `protobuf:"bytes,10,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (x *UpdateFeedRequest) Reset() {
	*x = UpdateFeedRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_chain_mw_v1_coin_currency_feed_feed_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateFeedRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateFeedRequest) ProtoMessage() {}

func (x *UpdateFeedRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_chain_mw_v1_coin_currency_feed_feed_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateFeedRequest.ProtoReflect.Descriptor instead.
func (*UpdateFeedRequest) Descriptor() ([]byte, []int) {
	return file_npool_chain_mw_v1_coin_currency_feed_feed_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateFeedRequest) GetInfo() *FeedReq {
	if x != nil {
		return x.Info
	}
	return nil
}

type UpdateFeedResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *Feed `protobuf:"bytes,10,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (x *UpdateFeedResponse) Reset() {
	*x = UpdateFeedResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_chain_mw_v1_coin_currency_feed_feed_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateFeedResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateFeedResponse) ProtoMessage() {}

func (x *UpdateFeedResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_chain_mw_v1_coin_currency_feed_feed_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateFeedResponse.ProtoReflect.Descriptor instead.
func (*UpdateFeedResponse) Descriptor() ([]byte, []int) {
	return file_npool_chain_mw_v1_coin_currency_feed_feed_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateFeedResponse) GetInfo() *Feed {
	if x != nil {
		return x.Info
	}
	return nil
}

type GetFeedsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Conds  *Conds `protobuf:"bytes,10,opt,name=Conds,proto3" json:"Conds,omitempty"`
	Offset int32  `protobuf:"varint,20,opt,name=Offset,proto3" json:"Offset,omitempty"`
	Limit  int32  `protobuf:"varint,30,opt,name=Limit,proto3" json:"Limit,omitempty"`
}

func (x *GetFeedsRequest) Reset() {
	*x = GetFeedsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_chain_mw_v1_coin_currency_feed_feed_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFeedsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFeedsRequest) ProtoMessage() {}

func (x *GetFeedsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_chain_mw_v1_coin_currency_feed_feed_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFeedsRequest.ProtoReflect.Descriptor instead.
func (*GetFeedsRequest) Descriptor() ([]byte, []int) {
	return file_npool_chain_mw_v1_coin_currency_feed_feed_proto_rawDescGZIP(), []int{7}
}

func (x *GetFeedsRequest) GetConds() *Conds {
	if x != nil {
		return x.Conds
	}
	return nil
}

func (x *GetFeedsRequest) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *GetFeedsRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type GetFeedsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Infos []*Feed `protobuf:"bytes,10,rep,name=Infos,proto3" json:"Infos,omitempty"`
	Total uint32  `protobuf:"varint,20,opt,name=Total,proto3" json:"Total,omitempty"`
}

func (x *GetFeedsResponse) Reset() {
	*x = GetFeedsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_chain_mw_v1_coin_currency_feed_feed_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFeedsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFeedsResponse) ProtoMessage() {}

func (x *GetFeedsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_chain_mw_v1_coin_currency_feed_feed_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFeedsResponse.ProtoReflect.Descriptor instead.
func (*GetFeedsResponse) Descriptor() ([]byte, []int) {
	return file_npool_chain_mw_v1_coin_currency_feed_feed_proto_rawDescGZIP(), []int{8}
}

func (x *GetFeedsResponse) GetInfos() []*Feed {
	if x != nil {
		return x.Infos
	}
	return nil
}

func (x *GetFeedsResponse) GetTotal() uint32 {
	if x != nil {
		return x.Total
	}
	return 0
}

var File_npool_chain_mw_v1_coin_currency_feed_feed_proto protoreflect.FileDescriptor

var file_npool_chain_mw_v1_coin_currency_feed_feed_proto_rawDesc = []byte{
	0x0a, 0x2f, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2f, 0x6d, 0x77,
	0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x69, 0x6e, 0x2f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63,
	0x79, 0x2f, 0x66, 0x65, 0x65, 0x64, 0x2f, 0x66, 0x65, 0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x27, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77,
	0x61, 0x72, 0x65, 0x2e, 0x63, 0x6f, 0x69, 0x6e, 0x2e, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63,
	0x79, 0x31, 0x2e, 0x66, 0x65, 0x65, 0x64, 0x2e, 0x76, 0x31, 0x1a, 0x1e, 0x6e, 0x70, 0x6f, 0x6f,
	0x6c, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63,
	0x6f, 0x6e, 0x64, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x29, 0x6e, 0x70, 0x6f, 0x6f,
	0x6c, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63,
	0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x66, 0x65, 0x65, 0x64, 0x74, 0x79, 0x70, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8f, 0x02, 0x0a, 0x07, 0x46, 0x65, 0x65, 0x64, 0x52, 0x65,
	0x71, 0x12, 0x13, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52,
	0x02, 0x49, 0x44, 0x88, 0x01, 0x01, 0x12, 0x23, 0x0a, 0x0a, 0x43, 0x6f, 0x69, 0x6e, 0x54, 0x79,
	0x70, 0x65, 0x49, 0x44, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x0a, 0x43, 0x6f,
	0x69, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x49, 0x44, 0x88, 0x01, 0x01, 0x12, 0x3f, 0x0a, 0x08, 0x46,
	0x65, 0x65, 0x64, 0x54, 0x79, 0x70, 0x65, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1e, 0x2e,
	0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x75, 0x72,
	0x72, 0x65, 0x6e, 0x63, 0x79, 0x46, 0x65, 0x65, 0x64, 0x54, 0x79, 0x70, 0x65, 0x48, 0x02, 0x52,
	0x08, 0x46, 0x65, 0x65, 0x64, 0x54, 0x79, 0x70, 0x65, 0x88, 0x01, 0x01, 0x12, 0x27, 0x0a, 0x0c,
	0x46, 0x65, 0x65, 0x64, 0x43, 0x6f, 0x69, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x28, 0x20, 0x01,
	0x28, 0x09, 0x48, 0x03, 0x52, 0x0c, 0x46, 0x65, 0x65, 0x64, 0x43, 0x6f, 0x69, 0x6e, 0x4e, 0x61,
	0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x1f, 0x0a, 0x08, 0x44, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65,
	0x64, 0x18, 0x32, 0x20, 0x01, 0x28, 0x08, 0x48, 0x04, 0x52, 0x08, 0x44, 0x69, 0x73, 0x61, 0x62,
	0x6c, 0x65, 0x64, 0x88, 0x01, 0x01, 0x42, 0x05, 0x0a, 0x03, 0x5f, 0x49, 0x44, 0x42, 0x0d, 0x0a,
	0x0b, 0x5f, 0x43, 0x6f, 0x69, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x49, 0x44, 0x42, 0x0b, 0x0a, 0x09,
	0x5f, 0x46, 0x65, 0x65, 0x64, 0x54, 0x79, 0x70, 0x65, 0x42, 0x0f, 0x0a, 0x0d, 0x5f, 0x46, 0x65,
	0x65, 0x64, 0x43, 0x6f, 0x69, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x44,
	0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x22, 0x80, 0x03, 0x0a, 0x04, 0x46, 0x65, 0x65, 0x64,
	0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44,
	0x12, 0x1e, 0x0a, 0x0a, 0x43, 0x6f, 0x69, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x49, 0x44, 0x18, 0x14,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x43, 0x6f, 0x69, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x49, 0x44,
	0x12, 0x1a, 0x0a, 0x08, 0x43, 0x6f, 0x69, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x1e, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x43, 0x6f, 0x69, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x43, 0x6f, 0x69, 0x6e, 0x55, 0x6e, 0x69, 0x74, 0x18, 0x28, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x43, 0x6f, 0x69, 0x6e, 0x55, 0x6e, 0x69, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x43, 0x6f, 0x69, 0x6e,
	0x4c, 0x6f, 0x67, 0x6f, 0x18, 0x32, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x43, 0x6f, 0x69, 0x6e,
	0x4c, 0x6f, 0x67, 0x6f, 0x12, 0x18, 0x0a, 0x07, 0x43, 0x6f, 0x69, 0x6e, 0x45, 0x4e, 0x56, 0x18,
	0x3c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x43, 0x6f, 0x69, 0x6e, 0x45, 0x4e, 0x56, 0x12, 0x20,
	0x0a, 0x0b, 0x46, 0x65, 0x65, 0x64, 0x54, 0x79, 0x70, 0x65, 0x53, 0x74, 0x72, 0x18, 0x46, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x46, 0x65, 0x65, 0x64, 0x54, 0x79, 0x70, 0x65, 0x53, 0x74, 0x72,
	0x12, 0x3a, 0x0a, 0x08, 0x46, 0x65, 0x65, 0x64, 0x54, 0x79, 0x70, 0x65, 0x18, 0x50, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x1e, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x46, 0x65, 0x65, 0x64, 0x54, 0x79,
	0x70, 0x65, 0x52, 0x08, 0x46, 0x65, 0x65, 0x64, 0x54, 0x79, 0x70, 0x65, 0x12, 0x22, 0x0a, 0x0c,
	0x46, 0x65, 0x65, 0x64, 0x43, 0x6f, 0x69, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x5a, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0c, 0x46, 0x65, 0x65, 0x64, 0x43, 0x6f, 0x69, 0x6e, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x44, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x64, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x08, 0x44, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x1d, 0x0a, 0x09,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0xe8, 0x07, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x09, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0xf2, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0xea, 0x02, 0x0a, 0x05, 0x43,
	0x6f, 0x6e, 0x64, 0x73, 0x12, 0x2c, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x17, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x48, 0x00, 0x52, 0x02, 0x49, 0x44, 0x88,
	0x01, 0x01, 0x12, 0x3c, 0x0a, 0x0a, 0x43, 0x6f, 0x69, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x49, 0x44,
	0x18, 0x14, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70,
	0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x48,
	0x01, 0x52, 0x0a, 0x43, 0x6f, 0x69, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x49, 0x44, 0x88, 0x01, 0x01,
	0x12, 0x43, 0x0a, 0x0b, 0x43, 0x6f, 0x69, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x49, 0x44, 0x73, 0x18,
	0x1e, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x53, 0x6c, 0x69, 0x63, 0x65,
	0x56, 0x61, 0x6c, 0x48, 0x02, 0x52, 0x0b, 0x43, 0x6f, 0x69, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x49,
	0x44, 0x73, 0x88, 0x01, 0x01, 0x12, 0x38, 0x0a, 0x08, 0x46, 0x65, 0x65, 0x64, 0x54, 0x79, 0x70,
	0x65, 0x18, 0x28, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c,
	0x48, 0x03, 0x52, 0x08, 0x46, 0x65, 0x65, 0x64, 0x54, 0x79, 0x70, 0x65, 0x88, 0x01, 0x01, 0x12,
	0x36, 0x0a, 0x08, 0x44, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x32, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x15, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x48, 0x04, 0x52, 0x08, 0x44, 0x69, 0x73, 0x61,
	0x62, 0x6c, 0x65, 0x64, 0x88, 0x01, 0x01, 0x42, 0x05, 0x0a, 0x03, 0x5f, 0x49, 0x44, 0x42, 0x0d,
	0x0a, 0x0b, 0x5f, 0x43, 0x6f, 0x69, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x49, 0x44, 0x42, 0x0e, 0x0a,
	0x0c, 0x5f, 0x43, 0x6f, 0x69, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x49, 0x44, 0x73, 0x42, 0x0b, 0x0a,
	0x09, 0x5f, 0x46, 0x65, 0x65, 0x64, 0x54, 0x79, 0x70, 0x65, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x44,
	0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x22, 0x59, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x46, 0x65, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x44, 0x0a, 0x04,
	0x49, 0x6e, 0x66, 0x6f, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x63, 0x68, 0x61,
	0x69, 0x6e, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x63, 0x6f,
	0x69, 0x6e, 0x2e, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x31, 0x2e, 0x66, 0x65, 0x65,
	0x64, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x65, 0x65, 0x64, 0x52, 0x65, 0x71, 0x52, 0x04, 0x49, 0x6e,
	0x66, 0x6f, 0x22, 0x57, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x46, 0x65, 0x65, 0x64,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x41, 0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x6d,
	0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x63, 0x6f, 0x69, 0x6e, 0x2e, 0x63,
	0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x31, 0x2e, 0x66, 0x65, 0x65, 0x64, 0x2e, 0x76, 0x31,
	0x2e, 0x46, 0x65, 0x65, 0x64, 0x52, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x59, 0x0a, 0x11, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x46, 0x65, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x44, 0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x30,
	0x2e, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72,
	0x65, 0x2e, 0x63, 0x6f, 0x69, 0x6e, 0x2e, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x31,
	0x2e, 0x66, 0x65, 0x65, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x65, 0x65, 0x64, 0x52, 0x65, 0x71,
	0x52, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x57, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x46, 0x65, 0x65, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x41, 0x0a, 0x04,
	0x49, 0x6e, 0x66, 0x6f, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x63, 0x68, 0x61,
	0x69, 0x6e, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x63, 0x6f,
	0x69, 0x6e, 0x2e, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x31, 0x2e, 0x66, 0x65, 0x65,
	0x64, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x65, 0x65, 0x64, 0x52, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x22,
	0x85, 0x01, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x46, 0x65, 0x65, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x44, 0x0a, 0x05, 0x43, 0x6f, 0x6e, 0x64, 0x73, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c,
	0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x63, 0x6f, 0x69, 0x6e, 0x2e, 0x63, 0x75, 0x72, 0x72, 0x65,
	0x6e, 0x63, 0x79, 0x31, 0x2e, 0x66, 0x65, 0x65, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6e,
	0x64, 0x73, 0x52, 0x05, 0x43, 0x6f, 0x6e, 0x64, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x4f, 0x66, 0x66,
	0x73, 0x65, 0x74, 0x18, 0x14, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x4f, 0x66, 0x66, 0x73, 0x65,
	0x74, 0x12, 0x14, 0x0a, 0x05, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x05, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x22, 0x6d, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x46, 0x65,
	0x65, 0x64, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x43, 0x0a, 0x05, 0x49,
	0x6e, 0x66, 0x6f, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x63, 0x68, 0x61,
	0x69, 0x6e, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x63, 0x6f,
	0x69, 0x6e, 0x2e, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x31, 0x2e, 0x66, 0x65, 0x65,
	0x64, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x65, 0x65, 0x64, 0x52, 0x05, 0x49, 0x6e, 0x66, 0x6f, 0x73,
	0x12, 0x14, 0x0a, 0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x14, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x32, 0xa4, 0x03, 0x0a, 0x0a, 0x4d, 0x69, 0x64, 0x64, 0x6c,
	0x65, 0x77, 0x61, 0x72, 0x65, 0x12, 0x87, 0x01, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x46, 0x65, 0x65, 0x64, 0x12, 0x3a, 0x2e, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x6d, 0x69, 0x64,
	0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x63, 0x6f, 0x69, 0x6e, 0x2e, 0x63, 0x75, 0x72,
	0x72, 0x65, 0x6e, 0x63, 0x79, 0x31, 0x2e, 0x66, 0x65, 0x65, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x46, 0x65, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x3b, 0x2e, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77,
	0x61, 0x72, 0x65, 0x2e, 0x63, 0x6f, 0x69, 0x6e, 0x2e, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63,
	0x79, 0x31, 0x2e, 0x66, 0x65, 0x65, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x46, 0x65, 0x65, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x87, 0x01, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x46, 0x65, 0x65, 0x64, 0x12, 0x3a,
	0x2e, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72,
	0x65, 0x2e, 0x63, 0x6f, 0x69, 0x6e, 0x2e, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x31,
	0x2e, 0x66, 0x65, 0x65, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x46,
	0x65, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3b, 0x2e, 0x63, 0x68, 0x61,
	0x69, 0x6e, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x63, 0x6f,
	0x69, 0x6e, 0x2e, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x31, 0x2e, 0x66, 0x65, 0x65,
	0x64, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x46, 0x65, 0x65, 0x64, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x81, 0x01, 0x0a, 0x08, 0x47, 0x65,
	0x74, 0x46, 0x65, 0x65, 0x64, 0x73, 0x12, 0x38, 0x2e, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x6d,
	0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x63, 0x6f, 0x69, 0x6e, 0x2e, 0x63,
	0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x31, 0x2e, 0x66, 0x65, 0x65, 0x64, 0x2e, 0x76, 0x31,
	0x2e, 0x47, 0x65, 0x74, 0x46, 0x65, 0x65, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x39, 0x2e, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77,
	0x61, 0x72, 0x65, 0x2e, 0x63, 0x6f, 0x69, 0x6e, 0x2e, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63,
	0x79, 0x31, 0x2e, 0x66, 0x65, 0x65, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x65,
	0x65, 0x64, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x47, 0x5a,
	0x45, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4e, 0x70, 0x6f, 0x6f,
	0x6c, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x2f, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2f, 0x6d, 0x77,
	0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x69, 0x6e, 0x2f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63,
	0x79, 0x2f, 0x66, 0x65, 0x65, 0x64, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_npool_chain_mw_v1_coin_currency_feed_feed_proto_rawDescOnce sync.Once
	file_npool_chain_mw_v1_coin_currency_feed_feed_proto_rawDescData = file_npool_chain_mw_v1_coin_currency_feed_feed_proto_rawDesc
)

func file_npool_chain_mw_v1_coin_currency_feed_feed_proto_rawDescGZIP() []byte {
	file_npool_chain_mw_v1_coin_currency_feed_feed_proto_rawDescOnce.Do(func() {
		file_npool_chain_mw_v1_coin_currency_feed_feed_proto_rawDescData = protoimpl.X.CompressGZIP(file_npool_chain_mw_v1_coin_currency_feed_feed_proto_rawDescData)
	})
	return file_npool_chain_mw_v1_coin_currency_feed_feed_proto_rawDescData
}

var file_npool_chain_mw_v1_coin_currency_feed_feed_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_npool_chain_mw_v1_coin_currency_feed_feed_proto_goTypes = []interface{}{
	(*FeedReq)(nil),            // 0: chain.middleware.coin.currency1.feed.v1.FeedReq
	(*Feed)(nil),               // 1: chain.middleware.coin.currency1.feed.v1.Feed
	(*Conds)(nil),              // 2: chain.middleware.coin.currency1.feed.v1.Conds
	(*CreateFeedRequest)(nil),  // 3: chain.middleware.coin.currency1.feed.v1.CreateFeedRequest
	(*CreateFeedResponse)(nil), // 4: chain.middleware.coin.currency1.feed.v1.CreateFeedResponse
	(*UpdateFeedRequest)(nil),  // 5: chain.middleware.coin.currency1.feed.v1.UpdateFeedRequest
	(*UpdateFeedResponse)(nil), // 6: chain.middleware.coin.currency1.feed.v1.UpdateFeedResponse
	(*GetFeedsRequest)(nil),    // 7: chain.middleware.coin.currency1.feed.v1.GetFeedsRequest
	(*GetFeedsResponse)(nil),   // 8: chain.middleware.coin.currency1.feed.v1.GetFeedsResponse
	(v1.CurrencyFeedType)(0),   // 9: basetypes.v1.CurrencyFeedType
	(*v1.StringVal)(nil),       // 10: basetypes.v1.StringVal
	(*v1.StringSliceVal)(nil),  // 11: basetypes.v1.StringSliceVal
	(*v1.Uint32Val)(nil),       // 12: basetypes.v1.Uint32Val
	(*v1.BoolVal)(nil),         // 13: basetypes.v1.BoolVal
}
var file_npool_chain_mw_v1_coin_currency_feed_feed_proto_depIdxs = []int32{
	9,  // 0: chain.middleware.coin.currency1.feed.v1.FeedReq.FeedType:type_name -> basetypes.v1.CurrencyFeedType
	9,  // 1: chain.middleware.coin.currency1.feed.v1.Feed.FeedType:type_name -> basetypes.v1.CurrencyFeedType
	10, // 2: chain.middleware.coin.currency1.feed.v1.Conds.ID:type_name -> basetypes.v1.StringVal
	10, // 3: chain.middleware.coin.currency1.feed.v1.Conds.CoinTypeID:type_name -> basetypes.v1.StringVal
	11, // 4: chain.middleware.coin.currency1.feed.v1.Conds.CoinTypeIDs:type_name -> basetypes.v1.StringSliceVal
	12, // 5: chain.middleware.coin.currency1.feed.v1.Conds.FeedType:type_name -> basetypes.v1.Uint32Val
	13, // 6: chain.middleware.coin.currency1.feed.v1.Conds.Disabled:type_name -> basetypes.v1.BoolVal
	0,  // 7: chain.middleware.coin.currency1.feed.v1.CreateFeedRequest.Info:type_name -> chain.middleware.coin.currency1.feed.v1.FeedReq
	1,  // 8: chain.middleware.coin.currency1.feed.v1.CreateFeedResponse.Info:type_name -> chain.middleware.coin.currency1.feed.v1.Feed
	0,  // 9: chain.middleware.coin.currency1.feed.v1.UpdateFeedRequest.Info:type_name -> chain.middleware.coin.currency1.feed.v1.FeedReq
	1,  // 10: chain.middleware.coin.currency1.feed.v1.UpdateFeedResponse.Info:type_name -> chain.middleware.coin.currency1.feed.v1.Feed
	2,  // 11: chain.middleware.coin.currency1.feed.v1.GetFeedsRequest.Conds:type_name -> chain.middleware.coin.currency1.feed.v1.Conds
	1,  // 12: chain.middleware.coin.currency1.feed.v1.GetFeedsResponse.Infos:type_name -> chain.middleware.coin.currency1.feed.v1.Feed
	3,  // 13: chain.middleware.coin.currency1.feed.v1.Middleware.CreateFeed:input_type -> chain.middleware.coin.currency1.feed.v1.CreateFeedRequest
	5,  // 14: chain.middleware.coin.currency1.feed.v1.Middleware.UpdateFeed:input_type -> chain.middleware.coin.currency1.feed.v1.UpdateFeedRequest
	7,  // 15: chain.middleware.coin.currency1.feed.v1.Middleware.GetFeeds:input_type -> chain.middleware.coin.currency1.feed.v1.GetFeedsRequest
	4,  // 16: chain.middleware.coin.currency1.feed.v1.Middleware.CreateFeed:output_type -> chain.middleware.coin.currency1.feed.v1.CreateFeedResponse
	6,  // 17: chain.middleware.coin.currency1.feed.v1.Middleware.UpdateFeed:output_type -> chain.middleware.coin.currency1.feed.v1.UpdateFeedResponse
	8,  // 18: chain.middleware.coin.currency1.feed.v1.Middleware.GetFeeds:output_type -> chain.middleware.coin.currency1.feed.v1.GetFeedsResponse
	16, // [16:19] is the sub-list for method output_type
	13, // [13:16] is the sub-list for method input_type
	13, // [13:13] is the sub-list for extension type_name
	13, // [13:13] is the sub-list for extension extendee
	0,  // [0:13] is the sub-list for field type_name
}

func init() { file_npool_chain_mw_v1_coin_currency_feed_feed_proto_init() }
func file_npool_chain_mw_v1_coin_currency_feed_feed_proto_init() {
	if File_npool_chain_mw_v1_coin_currency_feed_feed_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_npool_chain_mw_v1_coin_currency_feed_feed_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FeedReq); i {
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
		file_npool_chain_mw_v1_coin_currency_feed_feed_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Feed); i {
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
		file_npool_chain_mw_v1_coin_currency_feed_feed_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
		file_npool_chain_mw_v1_coin_currency_feed_feed_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateFeedRequest); i {
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
		file_npool_chain_mw_v1_coin_currency_feed_feed_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateFeedResponse); i {
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
		file_npool_chain_mw_v1_coin_currency_feed_feed_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateFeedRequest); i {
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
		file_npool_chain_mw_v1_coin_currency_feed_feed_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateFeedResponse); i {
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
		file_npool_chain_mw_v1_coin_currency_feed_feed_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFeedsRequest); i {
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
		file_npool_chain_mw_v1_coin_currency_feed_feed_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFeedsResponse); i {
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
	file_npool_chain_mw_v1_coin_currency_feed_feed_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_npool_chain_mw_v1_coin_currency_feed_feed_proto_msgTypes[2].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_npool_chain_mw_v1_coin_currency_feed_feed_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_npool_chain_mw_v1_coin_currency_feed_feed_proto_goTypes,
		DependencyIndexes: file_npool_chain_mw_v1_coin_currency_feed_feed_proto_depIdxs,
		MessageInfos:      file_npool_chain_mw_v1_coin_currency_feed_feed_proto_msgTypes,
	}.Build()
	File_npool_chain_mw_v1_coin_currency_feed_feed_proto = out.File
	file_npool_chain_mw_v1_coin_currency_feed_feed_proto_rawDesc = nil
	file_npool_chain_mw_v1_coin_currency_feed_feed_proto_goTypes = nil
	file_npool_chain_mw_v1_coin_currency_feed_feed_proto_depIdxs = nil
}
