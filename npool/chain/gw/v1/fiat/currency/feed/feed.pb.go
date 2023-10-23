// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.18.1
// source: npool/chain/gw/v1/fiat/currency/feed/feed.proto

package feed

import (
	v1 "github.com/NpoolPlatform/message/npool/basetypes/v1"
	feed "github.com/NpoolPlatform/message/npool/chain/mw/v1/fiat/currency/feed"
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

type CreateFeedRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FiatID       string              `protobuf:"bytes,10,opt,name=FiatID,proto3" json:"FiatID,omitempty"`
	FeedType     v1.CurrencyFeedType `protobuf:"varint,20,opt,name=FeedType,proto3,enum=basetypes.v1.CurrencyFeedType" json:"FeedType,omitempty"`
	FeedFiatName string              `protobuf:"bytes,30,opt,name=FeedFiatName,proto3" json:"FeedFiatName,omitempty"`
}

func (x *CreateFeedRequest) Reset() {
	*x = CreateFeedRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_chain_gw_v1_fiat_currency_feed_feed_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateFeedRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateFeedRequest) ProtoMessage() {}

func (x *CreateFeedRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_chain_gw_v1_fiat_currency_feed_feed_proto_msgTypes[0]
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
	return file_npool_chain_gw_v1_fiat_currency_feed_feed_proto_rawDescGZIP(), []int{0}
}

func (x *CreateFeedRequest) GetFiatID() string {
	if x != nil {
		return x.FiatID
	}
	return ""
}

func (x *CreateFeedRequest) GetFeedType() v1.CurrencyFeedType {
	if x != nil {
		return x.FeedType
	}
	return v1.CurrencyFeedType(0)
}

func (x *CreateFeedRequest) GetFeedFiatName() string {
	if x != nil {
		return x.FeedFiatName
	}
	return ""
}

type CreateFeedResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *feed.Feed `protobuf:"bytes,10,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (x *CreateFeedResponse) Reset() {
	*x = CreateFeedResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_chain_gw_v1_fiat_currency_feed_feed_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateFeedResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateFeedResponse) ProtoMessage() {}

func (x *CreateFeedResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_chain_gw_v1_fiat_currency_feed_feed_proto_msgTypes[1]
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
	return file_npool_chain_gw_v1_fiat_currency_feed_feed_proto_rawDescGZIP(), []int{1}
}

func (x *CreateFeedResponse) GetInfo() *feed.Feed {
	if x != nil {
		return x.Info
	}
	return nil
}

type UpdateFeedRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID           uint32  `protobuf:"varint,10,opt,name=ID,proto3" json:"ID,omitempty"`
	FeedFiatName *string `protobuf:"bytes,20,opt,name=FeedFiatName,proto3,oneof" json:"FeedFiatName,omitempty"`
	Disabled     *bool   `protobuf:"varint,30,opt,name=Disabled,proto3,oneof" json:"Disabled,omitempty"`
}

func (x *UpdateFeedRequest) Reset() {
	*x = UpdateFeedRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_chain_gw_v1_fiat_currency_feed_feed_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateFeedRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateFeedRequest) ProtoMessage() {}

func (x *UpdateFeedRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_chain_gw_v1_fiat_currency_feed_feed_proto_msgTypes[2]
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
	return file_npool_chain_gw_v1_fiat_currency_feed_feed_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateFeedRequest) GetID() uint32 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *UpdateFeedRequest) GetFeedFiatName() string {
	if x != nil && x.FeedFiatName != nil {
		return *x.FeedFiatName
	}
	return ""
}

func (x *UpdateFeedRequest) GetDisabled() bool {
	if x != nil && x.Disabled != nil {
		return *x.Disabled
	}
	return false
}

type UpdateFeedResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *feed.Feed `protobuf:"bytes,10,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (x *UpdateFeedResponse) Reset() {
	*x = UpdateFeedResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_chain_gw_v1_fiat_currency_feed_feed_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateFeedResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateFeedResponse) ProtoMessage() {}

func (x *UpdateFeedResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_chain_gw_v1_fiat_currency_feed_feed_proto_msgTypes[3]
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
	return file_npool_chain_gw_v1_fiat_currency_feed_feed_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateFeedResponse) GetInfo() *feed.Feed {
	if x != nil {
		return x.Info
	}
	return nil
}

type GetFeedsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Offset int32 `protobuf:"varint,10,opt,name=Offset,proto3" json:"Offset,omitempty"`
	Limit  int32 `protobuf:"varint,20,opt,name=Limit,proto3" json:"Limit,omitempty"`
}

func (x *GetFeedsRequest) Reset() {
	*x = GetFeedsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_chain_gw_v1_fiat_currency_feed_feed_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFeedsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFeedsRequest) ProtoMessage() {}

func (x *GetFeedsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_chain_gw_v1_fiat_currency_feed_feed_proto_msgTypes[4]
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
	return file_npool_chain_gw_v1_fiat_currency_feed_feed_proto_rawDescGZIP(), []int{4}
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

	Infos []*feed.Feed `protobuf:"bytes,10,rep,name=Infos,proto3" json:"Infos,omitempty"`
	Total uint32       `protobuf:"varint,20,opt,name=Total,proto3" json:"Total,omitempty"`
}

func (x *GetFeedsResponse) Reset() {
	*x = GetFeedsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_chain_gw_v1_fiat_currency_feed_feed_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFeedsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFeedsResponse) ProtoMessage() {}

func (x *GetFeedsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_chain_gw_v1_fiat_currency_feed_feed_proto_msgTypes[5]
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
	return file_npool_chain_gw_v1_fiat_currency_feed_feed_proto_rawDescGZIP(), []int{5}
}

func (x *GetFeedsResponse) GetInfos() []*feed.Feed {
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

var File_npool_chain_gw_v1_fiat_currency_feed_feed_proto protoreflect.FileDescriptor

var file_npool_chain_gw_v1_fiat_currency_feed_feed_proto_rawDesc = []byte{
	0x0a, 0x2f, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2f, 0x67, 0x77,
	0x2f, 0x76, 0x31, 0x2f, 0x66, 0x69, 0x61, 0x74, 0x2f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63,
	0x79, 0x2f, 0x66, 0x65, 0x65, 0x64, 0x2f, 0x66, 0x65, 0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x23, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79,
	0x2e, 0x66, 0x69, 0x61, 0x74, 0x2e, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x2e, 0x66,
	0x65, 0x65, 0x64, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x29, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x62, 0x61, 0x73, 0x65,
	0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63,
	0x79, 0x66, 0x65, 0x65, 0x64, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x2f, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2f, 0x6d, 0x77, 0x2f,
	0x76, 0x31, 0x2f, 0x66, 0x69, 0x61, 0x74, 0x2f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79,
	0x2f, 0x66, 0x65, 0x65, 0x64, 0x2f, 0x66, 0x65, 0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x8b, 0x01, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x46, 0x65, 0x65, 0x64, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x46, 0x69, 0x61, 0x74, 0x49, 0x44,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x46, 0x69, 0x61, 0x74, 0x49, 0x44, 0x12, 0x3a,
	0x0a, 0x08, 0x46, 0x65, 0x65, 0x64, 0x54, 0x79, 0x70, 0x65, 0x18, 0x14, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x1e, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x46, 0x65, 0x65, 0x64, 0x54, 0x79, 0x70, 0x65,
	0x52, 0x08, 0x46, 0x65, 0x65, 0x64, 0x54, 0x79, 0x70, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x46, 0x65,
	0x65, 0x64, 0x46, 0x69, 0x61, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x46, 0x65, 0x65, 0x64, 0x46, 0x69, 0x61, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x57,
	0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x46, 0x65, 0x65, 0x64, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x41, 0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c,
	0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x66, 0x69, 0x61, 0x74, 0x2e, 0x63, 0x75, 0x72, 0x72, 0x65,
	0x6e, 0x63, 0x79, 0x31, 0x2e, 0x66, 0x65, 0x65, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x65, 0x65,
	0x64, 0x52, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x8b, 0x01, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x46, 0x65, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x49, 0x44, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x49, 0x44, 0x12, 0x27, 0x0a,
	0x0c, 0x46, 0x65, 0x65, 0x64, 0x46, 0x69, 0x61, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x14, 0x20,
	0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0c, 0x46, 0x65, 0x65, 0x64, 0x46, 0x69, 0x61, 0x74, 0x4e,
	0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x1f, 0x0a, 0x08, 0x44, 0x69, 0x73, 0x61, 0x62, 0x6c,
	0x65, 0x64, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x08, 0x48, 0x01, 0x52, 0x08, 0x44, 0x69, 0x73, 0x61,
	0x62, 0x6c, 0x65, 0x64, 0x88, 0x01, 0x01, 0x42, 0x0f, 0x0a, 0x0d, 0x5f, 0x46, 0x65, 0x65, 0x64,
	0x46, 0x69, 0x61, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x44, 0x69, 0x73,
	0x61, 0x62, 0x6c, 0x65, 0x64, 0x22, 0x57, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x46,
	0x65, 0x65, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x41, 0x0a, 0x04, 0x49,
	0x6e, 0x66, 0x6f, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x63, 0x68, 0x61, 0x69,
	0x6e, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x66, 0x69, 0x61,
	0x74, 0x2e, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x31, 0x2e, 0x66, 0x65, 0x65, 0x64,
	0x2e, 0x76, 0x31, 0x2e, 0x46, 0x65, 0x65, 0x64, 0x52, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x3f,
	0x0a, 0x0f, 0x47, 0x65, 0x74, 0x46, 0x65, 0x65, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x16, 0x0a, 0x06, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x06, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x4c, 0x69, 0x6d,
	0x69, 0x74, 0x18, 0x14, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x22,
	0x6d, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x46, 0x65, 0x65, 0x64, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x43, 0x0a, 0x05, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x18, 0x0a, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c,
	0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x66, 0x69, 0x61, 0x74, 0x2e, 0x63, 0x75, 0x72, 0x72, 0x65,
	0x6e, 0x63, 0x79, 0x31, 0x2e, 0x66, 0x65, 0x65, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x65, 0x65,
	0x64, 0x52, 0x05, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x6f, 0x74, 0x61,
	0x6c, 0x18, 0x14, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x32, 0xf9,
	0x03, 0x0a, 0x07, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x12, 0xa5, 0x01, 0x0a, 0x0a, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x46, 0x65, 0x65, 0x64, 0x12, 0x36, 0x2e, 0x63, 0x68, 0x61, 0x69,
	0x6e, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x66, 0x69, 0x61, 0x74, 0x2e, 0x63,
	0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x2e, 0x66, 0x65, 0x65, 0x64, 0x2e, 0x76, 0x31, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x46, 0x65, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x37, 0x2e, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61,
	0x79, 0x2e, 0x66, 0x69, 0x61, 0x74, 0x2e, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x2e,
	0x66, 0x65, 0x65, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x46, 0x65,
	0x65, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x26, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x20, 0x3a, 0x01, 0x2a, 0x22, 0x1b, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x2f, 0x66, 0x69, 0x61, 0x74, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x66, 0x65,
	0x65, 0x64, 0x12, 0xa5, 0x01, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x46, 0x65, 0x65,
	0x64, 0x12, 0x36, 0x2e, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61,
	0x79, 0x2e, 0x66, 0x69, 0x61, 0x74, 0x2e, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x2e,
	0x66, 0x65, 0x65, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x46, 0x65,
	0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x37, 0x2e, 0x63, 0x68, 0x61, 0x69,
	0x6e, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x66, 0x69, 0x61, 0x74, 0x2e, 0x63,
	0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x2e, 0x66, 0x65, 0x65, 0x64, 0x2e, 0x76, 0x31, 0x2e,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x46, 0x65, 0x65, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x26, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x20, 0x3a, 0x01, 0x2a, 0x22, 0x1b, 0x2f,
	0x76, 0x31, 0x2f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x66, 0x69, 0x61, 0x74, 0x63, 0x75,
	0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x66, 0x65, 0x65, 0x64, 0x12, 0x9d, 0x01, 0x0a, 0x08, 0x47,
	0x65, 0x74, 0x46, 0x65, 0x65, 0x64, 0x73, 0x12, 0x34, 0x2e, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e,
	0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x66, 0x69, 0x61, 0x74, 0x2e, 0x63, 0x75, 0x72,
	0x72, 0x65, 0x6e, 0x63, 0x79, 0x2e, 0x66, 0x65, 0x65, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65,
	0x74, 0x46, 0x65, 0x65, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x35, 0x2e,
	0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x66, 0x69,
	0x61, 0x74, 0x2e, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x2e, 0x66, 0x65, 0x65, 0x64,
	0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x65, 0x65, 0x64, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x24, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1e, 0x3a, 0x01, 0x2a, 0x22,
	0x19, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x65, 0x74, 0x2f, 0x66, 0x69, 0x61, 0x74, 0x63, 0x75, 0x72,
	0x72, 0x65, 0x6e, 0x63, 0x79, 0x66, 0x65, 0x65, 0x64, 0x73, 0x42, 0x47, 0x5a, 0x45, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4e, 0x70, 0x6f, 0x6f, 0x6c, 0x50, 0x6c,
	0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x6e,
	0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2f, 0x67, 0x77, 0x2f, 0x76, 0x31,
	0x2f, 0x66, 0x69, 0x61, 0x74, 0x2f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x2f, 0x66,
	0x65, 0x65, 0x64, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_npool_chain_gw_v1_fiat_currency_feed_feed_proto_rawDescOnce sync.Once
	file_npool_chain_gw_v1_fiat_currency_feed_feed_proto_rawDescData = file_npool_chain_gw_v1_fiat_currency_feed_feed_proto_rawDesc
)

func file_npool_chain_gw_v1_fiat_currency_feed_feed_proto_rawDescGZIP() []byte {
	file_npool_chain_gw_v1_fiat_currency_feed_feed_proto_rawDescOnce.Do(func() {
		file_npool_chain_gw_v1_fiat_currency_feed_feed_proto_rawDescData = protoimpl.X.CompressGZIP(file_npool_chain_gw_v1_fiat_currency_feed_feed_proto_rawDescData)
	})
	return file_npool_chain_gw_v1_fiat_currency_feed_feed_proto_rawDescData
}

var file_npool_chain_gw_v1_fiat_currency_feed_feed_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_npool_chain_gw_v1_fiat_currency_feed_feed_proto_goTypes = []interface{}{
	(*CreateFeedRequest)(nil),  // 0: chain.gateway.fiat.currency.feed.v1.CreateFeedRequest
	(*CreateFeedResponse)(nil), // 1: chain.gateway.fiat.currency.feed.v1.CreateFeedResponse
	(*UpdateFeedRequest)(nil),  // 2: chain.gateway.fiat.currency.feed.v1.UpdateFeedRequest
	(*UpdateFeedResponse)(nil), // 3: chain.gateway.fiat.currency.feed.v1.UpdateFeedResponse
	(*GetFeedsRequest)(nil),    // 4: chain.gateway.fiat.currency.feed.v1.GetFeedsRequest
	(*GetFeedsResponse)(nil),   // 5: chain.gateway.fiat.currency.feed.v1.GetFeedsResponse
	(v1.CurrencyFeedType)(0),   // 6: basetypes.v1.CurrencyFeedType
	(*feed.Feed)(nil),          // 7: chain.middleware.fiat.currency1.feed.v1.Feed
}
var file_npool_chain_gw_v1_fiat_currency_feed_feed_proto_depIdxs = []int32{
	6, // 0: chain.gateway.fiat.currency.feed.v1.CreateFeedRequest.FeedType:type_name -> basetypes.v1.CurrencyFeedType
	7, // 1: chain.gateway.fiat.currency.feed.v1.CreateFeedResponse.Info:type_name -> chain.middleware.fiat.currency1.feed.v1.Feed
	7, // 2: chain.gateway.fiat.currency.feed.v1.UpdateFeedResponse.Info:type_name -> chain.middleware.fiat.currency1.feed.v1.Feed
	7, // 3: chain.gateway.fiat.currency.feed.v1.GetFeedsResponse.Infos:type_name -> chain.middleware.fiat.currency1.feed.v1.Feed
	0, // 4: chain.gateway.fiat.currency.feed.v1.Gateway.CreateFeed:input_type -> chain.gateway.fiat.currency.feed.v1.CreateFeedRequest
	2, // 5: chain.gateway.fiat.currency.feed.v1.Gateway.UpdateFeed:input_type -> chain.gateway.fiat.currency.feed.v1.UpdateFeedRequest
	4, // 6: chain.gateway.fiat.currency.feed.v1.Gateway.GetFeeds:input_type -> chain.gateway.fiat.currency.feed.v1.GetFeedsRequest
	1, // 7: chain.gateway.fiat.currency.feed.v1.Gateway.CreateFeed:output_type -> chain.gateway.fiat.currency.feed.v1.CreateFeedResponse
	3, // 8: chain.gateway.fiat.currency.feed.v1.Gateway.UpdateFeed:output_type -> chain.gateway.fiat.currency.feed.v1.UpdateFeedResponse
	5, // 9: chain.gateway.fiat.currency.feed.v1.Gateway.GetFeeds:output_type -> chain.gateway.fiat.currency.feed.v1.GetFeedsResponse
	7, // [7:10] is the sub-list for method output_type
	4, // [4:7] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_npool_chain_gw_v1_fiat_currency_feed_feed_proto_init() }
func file_npool_chain_gw_v1_fiat_currency_feed_feed_proto_init() {
	if File_npool_chain_gw_v1_fiat_currency_feed_feed_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_npool_chain_gw_v1_fiat_currency_feed_feed_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_npool_chain_gw_v1_fiat_currency_feed_feed_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_npool_chain_gw_v1_fiat_currency_feed_feed_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
		file_npool_chain_gw_v1_fiat_currency_feed_feed_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
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
		file_npool_chain_gw_v1_fiat_currency_feed_feed_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
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
		file_npool_chain_gw_v1_fiat_currency_feed_feed_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
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
	file_npool_chain_gw_v1_fiat_currency_feed_feed_proto_msgTypes[2].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_npool_chain_gw_v1_fiat_currency_feed_feed_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_npool_chain_gw_v1_fiat_currency_feed_feed_proto_goTypes,
		DependencyIndexes: file_npool_chain_gw_v1_fiat_currency_feed_feed_proto_depIdxs,
		MessageInfos:      file_npool_chain_gw_v1_fiat_currency_feed_feed_proto_msgTypes,
	}.Build()
	File_npool_chain_gw_v1_fiat_currency_feed_feed_proto = out.File
	file_npool_chain_gw_v1_fiat_currency_feed_feed_proto_rawDesc = nil
	file_npool_chain_gw_v1_fiat_currency_feed_feed_proto_goTypes = nil
	file_npool_chain_gw_v1_fiat_currency_feed_feed_proto_depIdxs = nil
}
