// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v3.18.1
// source: npool/chain/mw/v1/fiat/currency/currency.proto

package currency

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

type CurrencyReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID              *uint32              `protobuf:"varint,10,opt,name=ID,proto3,oneof" json:"ID,omitempty"`
	EntID           *string              `protobuf:"bytes,11,opt,name=EntID,proto3,oneof" json:"EntID,omitempty"`
	FiatID          *string              `protobuf:"bytes,20,opt,name=FiatID,proto3,oneof" json:"FiatID,omitempty"`
	FeedType        *v1.CurrencyFeedType `protobuf:"varint,30,opt,name=FeedType,proto3,enum=basetypes.v1.CurrencyFeedType,oneof" json:"FeedType,omitempty"`
	MarketValueHigh *string              `protobuf:"bytes,40,opt,name=MarketValueHigh,proto3,oneof" json:"MarketValueHigh,omitempty"`
	MarketValueLow  *string              `protobuf:"bytes,50,opt,name=MarketValueLow,proto3,oneof" json:"MarketValueLow,omitempty"`
}

func (x *CurrencyReq) Reset() {
	*x = CurrencyReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_chain_mw_v1_fiat_currency_currency_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CurrencyReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CurrencyReq) ProtoMessage() {}

func (x *CurrencyReq) ProtoReflect() protoreflect.Message {
	mi := &file_npool_chain_mw_v1_fiat_currency_currency_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CurrencyReq.ProtoReflect.Descriptor instead.
func (*CurrencyReq) Descriptor() ([]byte, []int) {
	return file_npool_chain_mw_v1_fiat_currency_currency_proto_rawDescGZIP(), []int{0}
}

func (x *CurrencyReq) GetID() uint32 {
	if x != nil && x.ID != nil {
		return *x.ID
	}
	return 0
}

func (x *CurrencyReq) GetEntID() string {
	if x != nil && x.EntID != nil {
		return *x.EntID
	}
	return ""
}

func (x *CurrencyReq) GetFiatID() string {
	if x != nil && x.FiatID != nil {
		return *x.FiatID
	}
	return ""
}

func (x *CurrencyReq) GetFeedType() v1.CurrencyFeedType {
	if x != nil && x.FeedType != nil {
		return *x.FeedType
	}
	return v1.CurrencyFeedType(0)
}

func (x *CurrencyReq) GetMarketValueHigh() string {
	if x != nil && x.MarketValueHigh != nil {
		return *x.MarketValueHigh
	}
	return ""
}

func (x *CurrencyReq) GetMarketValueLow() string {
	if x != nil && x.MarketValueLow != nil {
		return *x.MarketValueLow
	}
	return ""
}

type Currency struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: sql:"id"
	ID uint32 `protobuf:"varint,10,opt,name=ID,proto3" json:"ID,omitempty" sql:"id"`
	// @inject_tag: sql:"ent_id"
	EntID string `protobuf:"bytes,11,opt,name=EntID,proto3" json:"EntID,omitempty" sql:"ent_id"`
	// @inject_tag: sql:"fiat_id"
	FiatID string `protobuf:"bytes,20,opt,name=FiatID,proto3" json:"FiatID,omitempty" sql:"fiat_id"`
	// @inject_tag: sql:"feed_type"
	FeedTypeStr string              `protobuf:"bytes,30,opt,name=FeedTypeStr,proto3" json:"FeedTypeStr,omitempty" sql:"feed_type"`
	FeedType    v1.CurrencyFeedType `protobuf:"varint,40,opt,name=FeedType,proto3,enum=basetypes.v1.CurrencyFeedType" json:"FeedType,omitempty"`
	// @inject_tag: sql:"fiat_name"
	FiatName string `protobuf:"bytes,50,opt,name=FiatName,proto3" json:"FiatName,omitempty" sql:"fiat_name"`
	// @inject_tag: sql:"fiat_logo"
	FiatLogo string `protobuf:"bytes,60,opt,name=FiatLogo,proto3" json:"FiatLogo,omitempty" sql:"fiat_logo"`
	// @inject_tag: sql:"fiat_unit"
	FiatUnit string `protobuf:"bytes,70,opt,name=FiatUnit,proto3" json:"FiatUnit,omitempty" sql:"fiat_unit"`
	// @inject_tag: sql:"market_value_high"
	MarketValueHigh string `protobuf:"bytes,80,opt,name=MarketValueHigh,proto3" json:"MarketValueHigh,omitempty" sql:"market_value_high"`
	// @inject_tag: sql:"market_value_low"
	MarketValueLow string `protobuf:"bytes,90,opt,name=MarketValueLow,proto3" json:"MarketValueLow,omitempty" sql:"market_value_low"`
	// @inject_tag: sql:"created_at"
	CreatedAt uint32 `protobuf:"varint,1000,opt,name=CreatedAt,proto3" json:"CreatedAt,omitempty" sql:"created_at"`
	// @inject_tag: sql:"updated_at"
	UpdatedAt uint32 `protobuf:"varint,1010,opt,name=UpdatedAt,proto3" json:"UpdatedAt,omitempty" sql:"updated_at"`
}

func (x *Currency) Reset() {
	*x = Currency{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_chain_mw_v1_fiat_currency_currency_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Currency) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Currency) ProtoMessage() {}

func (x *Currency) ProtoReflect() protoreflect.Message {
	mi := &file_npool_chain_mw_v1_fiat_currency_currency_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Currency.ProtoReflect.Descriptor instead.
func (*Currency) Descriptor() ([]byte, []int) {
	return file_npool_chain_mw_v1_fiat_currency_currency_proto_rawDescGZIP(), []int{1}
}

func (x *Currency) GetID() uint32 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *Currency) GetEntID() string {
	if x != nil {
		return x.EntID
	}
	return ""
}

func (x *Currency) GetFiatID() string {
	if x != nil {
		return x.FiatID
	}
	return ""
}

func (x *Currency) GetFeedTypeStr() string {
	if x != nil {
		return x.FeedTypeStr
	}
	return ""
}

func (x *Currency) GetFeedType() v1.CurrencyFeedType {
	if x != nil {
		return x.FeedType
	}
	return v1.CurrencyFeedType(0)
}

func (x *Currency) GetFiatName() string {
	if x != nil {
		return x.FiatName
	}
	return ""
}

func (x *Currency) GetFiatLogo() string {
	if x != nil {
		return x.FiatLogo
	}
	return ""
}

func (x *Currency) GetFiatUnit() string {
	if x != nil {
		return x.FiatUnit
	}
	return ""
}

func (x *Currency) GetMarketValueHigh() string {
	if x != nil {
		return x.MarketValueHigh
	}
	return ""
}

func (x *Currency) GetMarketValueLow() string {
	if x != nil {
		return x.MarketValueLow
	}
	return ""
}

func (x *Currency) GetCreatedAt() uint32 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *Currency) GetUpdatedAt() uint32 {
	if x != nil {
		return x.UpdatedAt
	}
	return 0
}

type Conds struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EntID    *v1.StringVal      `protobuf:"bytes,10,opt,name=EntID,proto3,oneof" json:"EntID,omitempty"`
	FiatID   *v1.StringVal      `protobuf:"bytes,20,opt,name=FiatID,proto3,oneof" json:"FiatID,omitempty"`
	FiatIDs  *v1.StringSliceVal `protobuf:"bytes,30,opt,name=FiatIDs,proto3,oneof" json:"FiatIDs,omitempty"`
	FiatName *v1.StringVal      `protobuf:"bytes,40,opt,name=FiatName,proto3,oneof" json:"FiatName,omitempty"`
}

func (x *Conds) Reset() {
	*x = Conds{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_chain_mw_v1_fiat_currency_currency_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Conds) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Conds) ProtoMessage() {}

func (x *Conds) ProtoReflect() protoreflect.Message {
	mi := &file_npool_chain_mw_v1_fiat_currency_currency_proto_msgTypes[2]
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
	return file_npool_chain_mw_v1_fiat_currency_currency_proto_rawDescGZIP(), []int{2}
}

func (x *Conds) GetEntID() *v1.StringVal {
	if x != nil {
		return x.EntID
	}
	return nil
}

func (x *Conds) GetFiatID() *v1.StringVal {
	if x != nil {
		return x.FiatID
	}
	return nil
}

func (x *Conds) GetFiatIDs() *v1.StringSliceVal {
	if x != nil {
		return x.FiatIDs
	}
	return nil
}

func (x *Conds) GetFiatName() *v1.StringVal {
	if x != nil {
		return x.FiatName
	}
	return nil
}

type GetCurrencyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EntID string `protobuf:"bytes,10,opt,name=EntID,proto3" json:"EntID,omitempty"`
}

func (x *GetCurrencyRequest) Reset() {
	*x = GetCurrencyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_chain_mw_v1_fiat_currency_currency_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCurrencyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCurrencyRequest) ProtoMessage() {}

func (x *GetCurrencyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_chain_mw_v1_fiat_currency_currency_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCurrencyRequest.ProtoReflect.Descriptor instead.
func (*GetCurrencyRequest) Descriptor() ([]byte, []int) {
	return file_npool_chain_mw_v1_fiat_currency_currency_proto_rawDescGZIP(), []int{3}
}

func (x *GetCurrencyRequest) GetEntID() string {
	if x != nil {
		return x.EntID
	}
	return ""
}

type GetCurrencyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *Currency `protobuf:"bytes,10,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (x *GetCurrencyResponse) Reset() {
	*x = GetCurrencyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_chain_mw_v1_fiat_currency_currency_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCurrencyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCurrencyResponse) ProtoMessage() {}

func (x *GetCurrencyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_chain_mw_v1_fiat_currency_currency_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCurrencyResponse.ProtoReflect.Descriptor instead.
func (*GetCurrencyResponse) Descriptor() ([]byte, []int) {
	return file_npool_chain_mw_v1_fiat_currency_currency_proto_rawDescGZIP(), []int{4}
}

func (x *GetCurrencyResponse) GetInfo() *Currency {
	if x != nil {
		return x.Info
	}
	return nil
}

type GetCurrenciesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Conds  *Conds `protobuf:"bytes,10,opt,name=Conds,proto3" json:"Conds,omitempty"`
	Offset int32  `protobuf:"varint,20,opt,name=Offset,proto3" json:"Offset,omitempty"`
	Limit  int32  `protobuf:"varint,30,opt,name=Limit,proto3" json:"Limit,omitempty"`
}

func (x *GetCurrenciesRequest) Reset() {
	*x = GetCurrenciesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_chain_mw_v1_fiat_currency_currency_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCurrenciesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCurrenciesRequest) ProtoMessage() {}

func (x *GetCurrenciesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_chain_mw_v1_fiat_currency_currency_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCurrenciesRequest.ProtoReflect.Descriptor instead.
func (*GetCurrenciesRequest) Descriptor() ([]byte, []int) {
	return file_npool_chain_mw_v1_fiat_currency_currency_proto_rawDescGZIP(), []int{5}
}

func (x *GetCurrenciesRequest) GetConds() *Conds {
	if x != nil {
		return x.Conds
	}
	return nil
}

func (x *GetCurrenciesRequest) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *GetCurrenciesRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type GetCurrenciesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Infos []*Currency `protobuf:"bytes,10,rep,name=Infos,proto3" json:"Infos,omitempty"`
	Total uint32      `protobuf:"varint,20,opt,name=Total,proto3" json:"Total,omitempty"`
}

func (x *GetCurrenciesResponse) Reset() {
	*x = GetCurrenciesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_chain_mw_v1_fiat_currency_currency_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCurrenciesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCurrenciesResponse) ProtoMessage() {}

func (x *GetCurrenciesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_chain_mw_v1_fiat_currency_currency_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCurrenciesResponse.ProtoReflect.Descriptor instead.
func (*GetCurrenciesResponse) Descriptor() ([]byte, []int) {
	return file_npool_chain_mw_v1_fiat_currency_currency_proto_rawDescGZIP(), []int{6}
}

func (x *GetCurrenciesResponse) GetInfos() []*Currency {
	if x != nil {
		return x.Infos
	}
	return nil
}

func (x *GetCurrenciesResponse) GetTotal() uint32 {
	if x != nil {
		return x.Total
	}
	return 0
}

var File_npool_chain_mw_v1_fiat_currency_currency_proto protoreflect.FileDescriptor

var file_npool_chain_mw_v1_fiat_currency_currency_proto_rawDesc = []byte{
	0x0a, 0x2e, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2f, 0x6d, 0x77,
	0x2f, 0x76, 0x31, 0x2f, 0x66, 0x69, 0x61, 0x74, 0x2f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63,
	0x79, 0x2f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x21, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61,
	0x72, 0x65, 0x2e, 0x66, 0x69, 0x61, 0x74, 0x2e, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79,
	0x2e, 0x76, 0x31, 0x1a, 0x1e, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x29, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79,
	0x66, 0x65, 0x65, 0x64, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc7,
	0x02, 0x0a, 0x0b, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x52, 0x65, 0x71, 0x12, 0x13,
	0x0a, 0x02, 0x49, 0x44, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0d, 0x48, 0x00, 0x52, 0x02, 0x49, 0x44,
	0x88, 0x01, 0x01, 0x12, 0x19, 0x0a, 0x05, 0x45, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x0b, 0x20, 0x01,
	0x28, 0x09, 0x48, 0x01, 0x52, 0x05, 0x45, 0x6e, 0x74, 0x49, 0x44, 0x88, 0x01, 0x01, 0x12, 0x1b,
	0x0a, 0x06, 0x46, 0x69, 0x61, 0x74, 0x49, 0x44, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x48, 0x02,
	0x52, 0x06, 0x46, 0x69, 0x61, 0x74, 0x49, 0x44, 0x88, 0x01, 0x01, 0x12, 0x3f, 0x0a, 0x08, 0x46,
	0x65, 0x65, 0x64, 0x54, 0x79, 0x70, 0x65, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1e, 0x2e,
	0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x75, 0x72,
	0x72, 0x65, 0x6e, 0x63, 0x79, 0x46, 0x65, 0x65, 0x64, 0x54, 0x79, 0x70, 0x65, 0x48, 0x03, 0x52,
	0x08, 0x46, 0x65, 0x65, 0x64, 0x54, 0x79, 0x70, 0x65, 0x88, 0x01, 0x01, 0x12, 0x2d, 0x0a, 0x0f,
	0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x48, 0x69, 0x67, 0x68, 0x18,
	0x28, 0x20, 0x01, 0x28, 0x09, 0x48, 0x04, 0x52, 0x0f, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x48, 0x69, 0x67, 0x68, 0x88, 0x01, 0x01, 0x12, 0x2b, 0x0a, 0x0e, 0x4d,
	0x61, 0x72, 0x6b, 0x65, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x4c, 0x6f, 0x77, 0x18, 0x32, 0x20,
	0x01, 0x28, 0x09, 0x48, 0x05, 0x52, 0x0e, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x4c, 0x6f, 0x77, 0x88, 0x01, 0x01, 0x42, 0x05, 0x0a, 0x03, 0x5f, 0x49, 0x44, 0x42,
	0x08, 0x0a, 0x06, 0x5f, 0x45, 0x6e, 0x74, 0x49, 0x44, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x46, 0x69,
	0x61, 0x74, 0x49, 0x44, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x46, 0x65, 0x65, 0x64, 0x54, 0x79, 0x70,
	0x65, 0x42, 0x12, 0x0a, 0x10, 0x5f, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x48, 0x69, 0x67, 0x68, 0x42, 0x11, 0x0a, 0x0f, 0x5f, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x4c, 0x6f, 0x77, 0x22, 0x8a, 0x03, 0x0a, 0x08, 0x43, 0x75, 0x72,
	0x72, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x02, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x45, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x0b,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x45, 0x6e, 0x74, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x46,
	0x69, 0x61, 0x74, 0x49, 0x44, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x46, 0x69, 0x61,
	0x74, 0x49, 0x44, 0x12, 0x20, 0x0a, 0x0b, 0x46, 0x65, 0x65, 0x64, 0x54, 0x79, 0x70, 0x65, 0x53,
	0x74, 0x72, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x46, 0x65, 0x65, 0x64, 0x54, 0x79,
	0x70, 0x65, 0x53, 0x74, 0x72, 0x12, 0x3a, 0x0a, 0x08, 0x46, 0x65, 0x65, 0x64, 0x54, 0x79, 0x70,
	0x65, 0x18, 0x28, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1e, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x46,
	0x65, 0x65, 0x64, 0x54, 0x79, 0x70, 0x65, 0x52, 0x08, 0x46, 0x65, 0x65, 0x64, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x46, 0x69, 0x61, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x32, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x46, 0x69, 0x61, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a,
	0x08, 0x46, 0x69, 0x61, 0x74, 0x4c, 0x6f, 0x67, 0x6f, 0x18, 0x3c, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x46, 0x69, 0x61, 0x74, 0x4c, 0x6f, 0x67, 0x6f, 0x12, 0x1a, 0x0a, 0x08, 0x46, 0x69, 0x61,
	0x74, 0x55, 0x6e, 0x69, 0x74, 0x18, 0x46, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x46, 0x69, 0x61,
	0x74, 0x55, 0x6e, 0x69, 0x74, 0x12, 0x28, 0x0a, 0x0f, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x48, 0x69, 0x67, 0x68, 0x18, 0x50, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f,
	0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x48, 0x69, 0x67, 0x68, 0x12,
	0x26, 0x0a, 0x0e, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x4c, 0x6f,
	0x77, 0x18, 0x5a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x4c, 0x6f, 0x77, 0x12, 0x1d, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x18, 0xe8, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x18, 0xf2, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x96, 0x02, 0x0a, 0x05, 0x43, 0x6f, 0x6e, 0x64, 0x73, 0x12,
	0x32, 0x0a, 0x05, 0x45, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17,
	0x2e, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74,
	0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x48, 0x00, 0x52, 0x05, 0x45, 0x6e, 0x74, 0x49, 0x44,
	0x88, 0x01, 0x01, 0x12, 0x34, 0x0a, 0x06, 0x46, 0x69, 0x61, 0x74, 0x49, 0x44, 0x18, 0x14, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x48, 0x01, 0x52, 0x06,
	0x46, 0x69, 0x61, 0x74, 0x49, 0x44, 0x88, 0x01, 0x01, 0x12, 0x3b, 0x0a, 0x07, 0x46, 0x69, 0x61,
	0x74, 0x49, 0x44, 0x73, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x62, 0x61, 0x73,
	0x65, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67,
	0x53, 0x6c, 0x69, 0x63, 0x65, 0x56, 0x61, 0x6c, 0x48, 0x02, 0x52, 0x07, 0x46, 0x69, 0x61, 0x74,
	0x49, 0x44, 0x73, 0x88, 0x01, 0x01, 0x12, 0x38, 0x0a, 0x08, 0x46, 0x69, 0x61, 0x74, 0x4e, 0x61,
	0x6d, 0x65, 0x18, 0x28, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61,
	0x6c, 0x48, 0x03, 0x52, 0x08, 0x46, 0x69, 0x61, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01,
	0x42, 0x08, 0x0a, 0x06, 0x5f, 0x45, 0x6e, 0x74, 0x49, 0x44, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x46,
	0x69, 0x61, 0x74, 0x49, 0x44, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x46, 0x69, 0x61, 0x74, 0x49, 0x44,
	0x73, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x46, 0x69, 0x61, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x2a,
	0x0a, 0x12, 0x47, 0x65, 0x74, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x45, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x45, 0x6e, 0x74, 0x49, 0x44, 0x22, 0x56, 0x0a, 0x13, 0x47, 0x65,
	0x74, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x3f, 0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x2b, 0x2e, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61,
	0x72, 0x65, 0x2e, 0x66, 0x69, 0x61, 0x74, 0x2e, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79,
	0x2e, 0x76, 0x31, 0x2e, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x52, 0x04, 0x49, 0x6e,
	0x66, 0x6f, 0x22, 0x84, 0x01, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e,
	0x63, 0x69, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3e, 0x0a, 0x05, 0x43,
	0x6f, 0x6e, 0x64, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x63, 0x68, 0x61,
	0x69, 0x6e, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x66, 0x69,
	0x61, 0x74, 0x2e, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x43,
	0x6f, 0x6e, 0x64, 0x73, 0x52, 0x05, 0x43, 0x6f, 0x6e, 0x64, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x4f,
	0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x14, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x4f, 0x66, 0x66,
	0x73, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x1e, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x05, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x22, 0x70, 0x0a, 0x15, 0x47, 0x65, 0x74,
	0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x41, 0x0a, 0x05, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x2b, 0x2e, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65,
	0x77, 0x61, 0x72, 0x65, 0x2e, 0x66, 0x69, 0x61, 0x74, 0x2e, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e,
	0x63, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x52, 0x05,
	0x49, 0x6e, 0x66, 0x6f, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x14,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x32, 0x93, 0x02, 0x0a, 0x0a,
	0x4d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x12, 0x7e, 0x0a, 0x0b, 0x47, 0x65,
	0x74, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x35, 0x2e, 0x63, 0x68, 0x61, 0x69,
	0x6e, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x66, 0x69, 0x61,
	0x74, 0x2e, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65,
	0x74, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x36, 0x2e, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77,
	0x61, 0x72, 0x65, 0x2e, 0x66, 0x69, 0x61, 0x74, 0x2e, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63,
	0x79, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x84, 0x01, 0x0a, 0x0d, 0x47,
	0x65, 0x74, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x69, 0x65, 0x73, 0x12, 0x37, 0x2e, 0x63,
	0x68, 0x61, 0x69, 0x6e, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e,
	0x66, 0x69, 0x61, 0x74, 0x2e, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x2e, 0x76, 0x31,
	0x2e, 0x47, 0x65, 0x74, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x69, 0x65, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x38, 0x2e, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x6d, 0x69,
	0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x66, 0x69, 0x61, 0x74, 0x2e, 0x63, 0x75,
	0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x75, 0x72,
	0x72, 0x65, 0x6e, 0x63, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x42, 0x42, 0x5a, 0x40, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x4e, 0x70, 0x6f, 0x6f, 0x6c, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x63, 0x68, 0x61, 0x69,
	0x6e, 0x2f, 0x6d, 0x77, 0x2f, 0x76, 0x31, 0x2f, 0x66, 0x69, 0x61, 0x74, 0x2f, 0x63, 0x75, 0x72,
	0x72, 0x65, 0x6e, 0x63, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_npool_chain_mw_v1_fiat_currency_currency_proto_rawDescOnce sync.Once
	file_npool_chain_mw_v1_fiat_currency_currency_proto_rawDescData = file_npool_chain_mw_v1_fiat_currency_currency_proto_rawDesc
)

func file_npool_chain_mw_v1_fiat_currency_currency_proto_rawDescGZIP() []byte {
	file_npool_chain_mw_v1_fiat_currency_currency_proto_rawDescOnce.Do(func() {
		file_npool_chain_mw_v1_fiat_currency_currency_proto_rawDescData = protoimpl.X.CompressGZIP(file_npool_chain_mw_v1_fiat_currency_currency_proto_rawDescData)
	})
	return file_npool_chain_mw_v1_fiat_currency_currency_proto_rawDescData
}

var file_npool_chain_mw_v1_fiat_currency_currency_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_npool_chain_mw_v1_fiat_currency_currency_proto_goTypes = []interface{}{
	(*CurrencyReq)(nil),           // 0: chain.middleware.fiat.currency.v1.CurrencyReq
	(*Currency)(nil),              // 1: chain.middleware.fiat.currency.v1.Currency
	(*Conds)(nil),                 // 2: chain.middleware.fiat.currency.v1.Conds
	(*GetCurrencyRequest)(nil),    // 3: chain.middleware.fiat.currency.v1.GetCurrencyRequest
	(*GetCurrencyResponse)(nil),   // 4: chain.middleware.fiat.currency.v1.GetCurrencyResponse
	(*GetCurrenciesRequest)(nil),  // 5: chain.middleware.fiat.currency.v1.GetCurrenciesRequest
	(*GetCurrenciesResponse)(nil), // 6: chain.middleware.fiat.currency.v1.GetCurrenciesResponse
	(v1.CurrencyFeedType)(0),      // 7: basetypes.v1.CurrencyFeedType
	(*v1.StringVal)(nil),          // 8: basetypes.v1.StringVal
	(*v1.StringSliceVal)(nil),     // 9: basetypes.v1.StringSliceVal
}
var file_npool_chain_mw_v1_fiat_currency_currency_proto_depIdxs = []int32{
	7,  // 0: chain.middleware.fiat.currency.v1.CurrencyReq.FeedType:type_name -> basetypes.v1.CurrencyFeedType
	7,  // 1: chain.middleware.fiat.currency.v1.Currency.FeedType:type_name -> basetypes.v1.CurrencyFeedType
	8,  // 2: chain.middleware.fiat.currency.v1.Conds.EntID:type_name -> basetypes.v1.StringVal
	8,  // 3: chain.middleware.fiat.currency.v1.Conds.FiatID:type_name -> basetypes.v1.StringVal
	9,  // 4: chain.middleware.fiat.currency.v1.Conds.FiatIDs:type_name -> basetypes.v1.StringSliceVal
	8,  // 5: chain.middleware.fiat.currency.v1.Conds.FiatName:type_name -> basetypes.v1.StringVal
	1,  // 6: chain.middleware.fiat.currency.v1.GetCurrencyResponse.Info:type_name -> chain.middleware.fiat.currency.v1.Currency
	2,  // 7: chain.middleware.fiat.currency.v1.GetCurrenciesRequest.Conds:type_name -> chain.middleware.fiat.currency.v1.Conds
	1,  // 8: chain.middleware.fiat.currency.v1.GetCurrenciesResponse.Infos:type_name -> chain.middleware.fiat.currency.v1.Currency
	3,  // 9: chain.middleware.fiat.currency.v1.Middleware.GetCurrency:input_type -> chain.middleware.fiat.currency.v1.GetCurrencyRequest
	5,  // 10: chain.middleware.fiat.currency.v1.Middleware.GetCurrencies:input_type -> chain.middleware.fiat.currency.v1.GetCurrenciesRequest
	4,  // 11: chain.middleware.fiat.currency.v1.Middleware.GetCurrency:output_type -> chain.middleware.fiat.currency.v1.GetCurrencyResponse
	6,  // 12: chain.middleware.fiat.currency.v1.Middleware.GetCurrencies:output_type -> chain.middleware.fiat.currency.v1.GetCurrenciesResponse
	11, // [11:13] is the sub-list for method output_type
	9,  // [9:11] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_npool_chain_mw_v1_fiat_currency_currency_proto_init() }
func file_npool_chain_mw_v1_fiat_currency_currency_proto_init() {
	if File_npool_chain_mw_v1_fiat_currency_currency_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_npool_chain_mw_v1_fiat_currency_currency_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CurrencyReq); i {
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
		file_npool_chain_mw_v1_fiat_currency_currency_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Currency); i {
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
		file_npool_chain_mw_v1_fiat_currency_currency_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
		file_npool_chain_mw_v1_fiat_currency_currency_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCurrencyRequest); i {
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
		file_npool_chain_mw_v1_fiat_currency_currency_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCurrencyResponse); i {
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
		file_npool_chain_mw_v1_fiat_currency_currency_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCurrenciesRequest); i {
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
		file_npool_chain_mw_v1_fiat_currency_currency_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCurrenciesResponse); i {
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
	file_npool_chain_mw_v1_fiat_currency_currency_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_npool_chain_mw_v1_fiat_currency_currency_proto_msgTypes[2].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_npool_chain_mw_v1_fiat_currency_currency_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_npool_chain_mw_v1_fiat_currency_currency_proto_goTypes,
		DependencyIndexes: file_npool_chain_mw_v1_fiat_currency_currency_proto_depIdxs,
		MessageInfos:      file_npool_chain_mw_v1_fiat_currency_currency_proto_msgTypes,
	}.Build()
	File_npool_chain_mw_v1_fiat_currency_currency_proto = out.File
	file_npool_chain_mw_v1_fiat_currency_currency_proto_rawDesc = nil
	file_npool_chain_mw_v1_fiat_currency_currency_proto_goTypes = nil
	file_npool_chain_mw_v1_fiat_currency_currency_proto_depIdxs = nil
}
