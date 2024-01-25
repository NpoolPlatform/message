// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.18.1
// source: npool/chain/gw/v1/coin/usedfor/usedfor.proto

package usedfor

import (
	v1 "github.com/NpoolPlatform/message/npool/basetypes/chain/v1"
	usedfor "github.com/NpoolPlatform/message/npool/chain/mw/v1/coin/usedfor"
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

type CreateCoinUsedForRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CoinTypeID string         `protobuf:"bytes,10,opt,name=CoinTypeID,proto3" json:"CoinTypeID,omitempty"`
	UsedFor    v1.CoinUsedFor `protobuf:"varint,20,opt,name=UsedFor,proto3,enum=basetypes.chain.v1.CoinUsedFor" json:"UsedFor,omitempty"`
	Priority   *uint32        `protobuf:"varint,30,opt,name=Priority,proto3,oneof" json:"Priority,omitempty"`
}

func (x *CreateCoinUsedForRequest) Reset() {
	*x = CreateCoinUsedForRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_chain_gw_v1_coin_usedfor_usedfor_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCoinUsedForRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCoinUsedForRequest) ProtoMessage() {}

func (x *CreateCoinUsedForRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_chain_gw_v1_coin_usedfor_usedfor_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCoinUsedForRequest.ProtoReflect.Descriptor instead.
func (*CreateCoinUsedForRequest) Descriptor() ([]byte, []int) {
	return file_npool_chain_gw_v1_coin_usedfor_usedfor_proto_rawDescGZIP(), []int{0}
}

func (x *CreateCoinUsedForRequest) GetCoinTypeID() string {
	if x != nil {
		return x.CoinTypeID
	}
	return ""
}

func (x *CreateCoinUsedForRequest) GetUsedFor() v1.CoinUsedFor {
	if x != nil {
		return x.UsedFor
	}
	return v1.CoinUsedFor(0)
}

func (x *CreateCoinUsedForRequest) GetPriority() uint32 {
	if x != nil && x.Priority != nil {
		return *x.Priority
	}
	return 0
}

type CreateCoinUsedForResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *usedfor.CoinUsedFor `protobuf:"bytes,10,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (x *CreateCoinUsedForResponse) Reset() {
	*x = CreateCoinUsedForResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_chain_gw_v1_coin_usedfor_usedfor_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCoinUsedForResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCoinUsedForResponse) ProtoMessage() {}

func (x *CreateCoinUsedForResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_chain_gw_v1_coin_usedfor_usedfor_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCoinUsedForResponse.ProtoReflect.Descriptor instead.
func (*CreateCoinUsedForResponse) Descriptor() ([]byte, []int) {
	return file_npool_chain_gw_v1_coin_usedfor_usedfor_proto_rawDescGZIP(), []int{1}
}

func (x *CreateCoinUsedForResponse) GetInfo() *usedfor.CoinUsedFor {
	if x != nil {
		return x.Info
	}
	return nil
}

type GetCoinUsedForsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CoinTypeIDs []string         `protobuf:"bytes,10,rep,name=CoinTypeIDs,proto3" json:"CoinTypeIDs,omitempty"`
	UsedFors    []v1.CoinUsedFor `protobuf:"varint,20,rep,packed,name=UsedFors,proto3,enum=basetypes.chain.v1.CoinUsedFor" json:"UsedFors,omitempty"`
	Offset      int32            `protobuf:"varint,30,opt,name=Offset,proto3" json:"Offset,omitempty"`
	Limit       int32            `protobuf:"varint,40,opt,name=Limit,proto3" json:"Limit,omitempty"`
}

func (x *GetCoinUsedForsRequest) Reset() {
	*x = GetCoinUsedForsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_chain_gw_v1_coin_usedfor_usedfor_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCoinUsedForsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCoinUsedForsRequest) ProtoMessage() {}

func (x *GetCoinUsedForsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_chain_gw_v1_coin_usedfor_usedfor_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCoinUsedForsRequest.ProtoReflect.Descriptor instead.
func (*GetCoinUsedForsRequest) Descriptor() ([]byte, []int) {
	return file_npool_chain_gw_v1_coin_usedfor_usedfor_proto_rawDescGZIP(), []int{2}
}

func (x *GetCoinUsedForsRequest) GetCoinTypeIDs() []string {
	if x != nil {
		return x.CoinTypeIDs
	}
	return nil
}

func (x *GetCoinUsedForsRequest) GetUsedFors() []v1.CoinUsedFor {
	if x != nil {
		return x.UsedFors
	}
	return nil
}

func (x *GetCoinUsedForsRequest) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *GetCoinUsedForsRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type GetCoinUsedForsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Infos []*usedfor.CoinUsedFor `protobuf:"bytes,10,rep,name=Infos,proto3" json:"Infos,omitempty"`
	Total uint32                 `protobuf:"varint,20,opt,name=Total,proto3" json:"Total,omitempty"`
}

func (x *GetCoinUsedForsResponse) Reset() {
	*x = GetCoinUsedForsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_chain_gw_v1_coin_usedfor_usedfor_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCoinUsedForsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCoinUsedForsResponse) ProtoMessage() {}

func (x *GetCoinUsedForsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_chain_gw_v1_coin_usedfor_usedfor_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCoinUsedForsResponse.ProtoReflect.Descriptor instead.
func (*GetCoinUsedForsResponse) Descriptor() ([]byte, []int) {
	return file_npool_chain_gw_v1_coin_usedfor_usedfor_proto_rawDescGZIP(), []int{3}
}

func (x *GetCoinUsedForsResponse) GetInfos() []*usedfor.CoinUsedFor {
	if x != nil {
		return x.Infos
	}
	return nil
}

func (x *GetCoinUsedForsResponse) GetTotal() uint32 {
	if x != nil {
		return x.Total
	}
	return 0
}

type DeleteCoinUsedForRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID    uint32 `protobuf:"varint,10,opt,name=ID,proto3" json:"ID,omitempty"`
	EntID string `protobuf:"bytes,20,opt,name=EntID,proto3" json:"EntID,omitempty"`
}

func (x *DeleteCoinUsedForRequest) Reset() {
	*x = DeleteCoinUsedForRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_chain_gw_v1_coin_usedfor_usedfor_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteCoinUsedForRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteCoinUsedForRequest) ProtoMessage() {}

func (x *DeleteCoinUsedForRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_chain_gw_v1_coin_usedfor_usedfor_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteCoinUsedForRequest.ProtoReflect.Descriptor instead.
func (*DeleteCoinUsedForRequest) Descriptor() ([]byte, []int) {
	return file_npool_chain_gw_v1_coin_usedfor_usedfor_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteCoinUsedForRequest) GetID() uint32 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *DeleteCoinUsedForRequest) GetEntID() string {
	if x != nil {
		return x.EntID
	}
	return ""
}

type DeleteCoinUsedForResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *usedfor.CoinUsedFor `protobuf:"bytes,10,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (x *DeleteCoinUsedForResponse) Reset() {
	*x = DeleteCoinUsedForResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_chain_gw_v1_coin_usedfor_usedfor_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteCoinUsedForResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteCoinUsedForResponse) ProtoMessage() {}

func (x *DeleteCoinUsedForResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_chain_gw_v1_coin_usedfor_usedfor_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteCoinUsedForResponse.ProtoReflect.Descriptor instead.
func (*DeleteCoinUsedForResponse) Descriptor() ([]byte, []int) {
	return file_npool_chain_gw_v1_coin_usedfor_usedfor_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteCoinUsedForResponse) GetInfo() *usedfor.CoinUsedFor {
	if x != nil {
		return x.Info
	}
	return nil
}

var File_npool_chain_gw_v1_coin_usedfor_usedfor_proto protoreflect.FileDescriptor

var file_npool_chain_gw_v1_coin_usedfor_usedfor_proto_rawDesc = []byte{
	0x0a, 0x2c, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2f, 0x67, 0x77,
	0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x69, 0x6e, 0x2f, 0x75, 0x73, 0x65, 0x64, 0x66, 0x6f, 0x72,
	0x2f, 0x75, 0x73, 0x65, 0x64, 0x66, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1d,
	0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x63, 0x6f,
	0x69, 0x6e, 0x2e, 0x75, 0x73, 0x65, 0x64, 0x66, 0x6f, 0x72, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2c, 0x6e, 0x70, 0x6f,
	0x6f, 0x6c, 0x2f, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2f, 0x6d, 0x77, 0x2f, 0x76, 0x31, 0x2f, 0x63,
	0x6f, 0x69, 0x6e, 0x2f, 0x75, 0x73, 0x65, 0x64, 0x66, 0x6f, 0x72, 0x2f, 0x75, 0x73, 0x65, 0x64,
	0x66, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x24, 0x6e, 0x70, 0x6f, 0x6f, 0x6c,
	0x2f, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x63, 0x68, 0x61, 0x69, 0x6e,
	0x2f, 0x76, 0x31, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xa3, 0x01, 0x0a, 0x18, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x69, 0x6e, 0x55, 0x73,
	0x65, 0x64, 0x46, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0a,
	0x43, 0x6f, 0x69, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x49, 0x44, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x43, 0x6f, 0x69, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x49, 0x44, 0x12, 0x39, 0x0a, 0x07,
	0x55, 0x73, 0x65, 0x64, 0x46, 0x6f, 0x72, 0x18, 0x14, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1f, 0x2e,
	0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e,
	0x76, 0x31, 0x2e, 0x43, 0x6f, 0x69, 0x6e, 0x55, 0x73, 0x65, 0x64, 0x46, 0x6f, 0x72, 0x52, 0x07,
	0x55, 0x73, 0x65, 0x64, 0x46, 0x6f, 0x72, 0x12, 0x1f, 0x0a, 0x08, 0x50, 0x72, 0x69, 0x6f, 0x72,
	0x69, 0x74, 0x79, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x0d, 0x48, 0x00, 0x52, 0x08, 0x50, 0x72, 0x69,
	0x6f, 0x72, 0x69, 0x74, 0x79, 0x88, 0x01, 0x01, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x50, 0x72, 0x69,
	0x6f, 0x72, 0x69, 0x74, 0x79, 0x22, 0x5e, 0x0a, 0x19, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43,
	0x6f, 0x69, 0x6e, 0x55, 0x73, 0x65, 0x64, 0x46, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x41, 0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x2d, 0x2e, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77,
	0x61, 0x72, 0x65, 0x2e, 0x63, 0x6f, 0x69, 0x6e, 0x2e, 0x75, 0x73, 0x65, 0x64, 0x66, 0x6f, 0x72,
	0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x69, 0x6e, 0x55, 0x73, 0x65, 0x64, 0x46, 0x6f, 0x72, 0x52,
	0x04, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0xa5, 0x01, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x69,
	0x6e, 0x55, 0x73, 0x65, 0x64, 0x46, 0x6f, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x20, 0x0a, 0x0b, 0x43, 0x6f, 0x69, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x49, 0x44, 0x73, 0x18,
	0x0a, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x43, 0x6f, 0x69, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x49,
	0x44, 0x73, 0x12, 0x3b, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x64, 0x46, 0x6f, 0x72, 0x73, 0x18, 0x14,
	0x20, 0x03, 0x28, 0x0e, 0x32, 0x1f, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x2e, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x69, 0x6e, 0x55, 0x73,
	0x65, 0x64, 0x46, 0x6f, 0x72, 0x52, 0x08, 0x55, 0x73, 0x65, 0x64, 0x46, 0x6f, 0x72, 0x73, 0x12,
	0x16, 0x0a, 0x06, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x06, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x4c, 0x69, 0x6d, 0x69, 0x74,
	0x18, 0x28, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x22, 0x74, 0x0a,
	0x17, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x69, 0x6e, 0x55, 0x73, 0x65, 0x64, 0x46, 0x6f, 0x72, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x43, 0x0a, 0x05, 0x49, 0x6e, 0x66, 0x6f,
	0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e,
	0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x63, 0x6f, 0x69, 0x6e, 0x2e,
	0x75, 0x73, 0x65, 0x64, 0x66, 0x6f, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x69, 0x6e, 0x55,
	0x73, 0x65, 0x64, 0x46, 0x6f, 0x72, 0x52, 0x05, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x12, 0x14, 0x0a,
	0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x14, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x54, 0x6f,
	0x74, 0x61, 0x6c, 0x22, 0x40, 0x0a, 0x18, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6f, 0x69,
	0x6e, 0x55, 0x73, 0x65, 0x64, 0x46, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x49, 0x44, 0x12,
	0x14, 0x0a, 0x05, 0x45, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x45, 0x6e, 0x74, 0x49, 0x44, 0x22, 0x5e, 0x0a, 0x19, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43,
	0x6f, 0x69, 0x6e, 0x55, 0x73, 0x65, 0x64, 0x46, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x41, 0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x2d, 0x2e, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77,
	0x61, 0x72, 0x65, 0x2e, 0x63, 0x6f, 0x69, 0x6e, 0x2e, 0x75, 0x73, 0x65, 0x64, 0x66, 0x6f, 0x72,
	0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x69, 0x6e, 0x55, 0x73, 0x65, 0x64, 0x46, 0x6f, 0x72, 0x52,
	0x04, 0x49, 0x6e, 0x66, 0x6f, 0x32, 0x85, 0x04, 0x0a, 0x07, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61,
	0x79, 0x12, 0xa9, 0x01, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x69, 0x6e,
	0x55, 0x73, 0x65, 0x64, 0x46, 0x6f, 0x72, 0x12, 0x37, 0x2e, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e,
	0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x63, 0x6f, 0x69, 0x6e, 0x2e, 0x75, 0x73, 0x65,
	0x64, 0x66, 0x6f, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f,
	0x69, 0x6e, 0x55, 0x73, 0x65, 0x64, 0x46, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x38, 0x2e, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79,
	0x2e, 0x63, 0x6f, 0x69, 0x6e, 0x2e, 0x75, 0x73, 0x65, 0x64, 0x66, 0x6f, 0x72, 0x2e, 0x76, 0x31,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x69, 0x6e, 0x55, 0x73, 0x65, 0x64, 0x46,
	0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x21, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x1b, 0x3a, 0x01, 0x2a, 0x22, 0x16, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x2f, 0x63, 0x6f, 0x69, 0x6e, 0x75, 0x73, 0x65, 0x64, 0x66, 0x6f, 0x72, 0x12, 0xa1, 0x01,
	0x0a, 0x0f, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x69, 0x6e, 0x55, 0x73, 0x65, 0x64, 0x46, 0x6f, 0x72,
	0x73, 0x12, 0x35, 0x2e, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61,
	0x79, 0x2e, 0x63, 0x6f, 0x69, 0x6e, 0x2e, 0x75, 0x73, 0x65, 0x64, 0x66, 0x6f, 0x72, 0x2e, 0x76,
	0x31, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x69, 0x6e, 0x55, 0x73, 0x65, 0x64, 0x46, 0x6f, 0x72,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x36, 0x2e, 0x63, 0x68, 0x61, 0x69, 0x6e,
	0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x63, 0x6f, 0x69, 0x6e, 0x2e, 0x75, 0x73,
	0x65, 0x64, 0x66, 0x6f, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x69, 0x6e,
	0x55, 0x73, 0x65, 0x64, 0x46, 0x6f, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x1f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x19, 0x3a, 0x01, 0x2a, 0x22, 0x14, 0x2f, 0x76, 0x31,
	0x2f, 0x67, 0x65, 0x74, 0x2f, 0x63, 0x6f, 0x69, 0x6e, 0x75, 0x73, 0x65, 0x64, 0x66, 0x6f, 0x72,
	0x73, 0x12, 0xa9, 0x01, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6f, 0x69, 0x6e,
	0x55, 0x73, 0x65, 0x64, 0x46, 0x6f, 0x72, 0x12, 0x37, 0x2e, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e,
	0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x63, 0x6f, 0x69, 0x6e, 0x2e, 0x75, 0x73, 0x65,
	0x64, 0x66, 0x6f, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6f,
	0x69, 0x6e, 0x55, 0x73, 0x65, 0x64, 0x46, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x38, 0x2e, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79,
	0x2e, 0x63, 0x6f, 0x69, 0x6e, 0x2e, 0x75, 0x73, 0x65, 0x64, 0x66, 0x6f, 0x72, 0x2e, 0x76, 0x31,
	0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6f, 0x69, 0x6e, 0x55, 0x73, 0x65, 0x64, 0x46,
	0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x21, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x1b, 0x3a, 0x01, 0x2a, 0x22, 0x16, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x2f, 0x63, 0x6f, 0x69, 0x6e, 0x75, 0x73, 0x65, 0x64, 0x66, 0x6f, 0x72, 0x42, 0x41, 0x5a,
	0x3f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4e, 0x70, 0x6f, 0x6f,
	0x6c, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x2f, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2f, 0x67, 0x77,
	0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x69, 0x6e, 0x2f, 0x75, 0x73, 0x65, 0x64, 0x66, 0x6f, 0x72,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_npool_chain_gw_v1_coin_usedfor_usedfor_proto_rawDescOnce sync.Once
	file_npool_chain_gw_v1_coin_usedfor_usedfor_proto_rawDescData = file_npool_chain_gw_v1_coin_usedfor_usedfor_proto_rawDesc
)

func file_npool_chain_gw_v1_coin_usedfor_usedfor_proto_rawDescGZIP() []byte {
	file_npool_chain_gw_v1_coin_usedfor_usedfor_proto_rawDescOnce.Do(func() {
		file_npool_chain_gw_v1_coin_usedfor_usedfor_proto_rawDescData = protoimpl.X.CompressGZIP(file_npool_chain_gw_v1_coin_usedfor_usedfor_proto_rawDescData)
	})
	return file_npool_chain_gw_v1_coin_usedfor_usedfor_proto_rawDescData
}

var file_npool_chain_gw_v1_coin_usedfor_usedfor_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_npool_chain_gw_v1_coin_usedfor_usedfor_proto_goTypes = []interface{}{
	(*CreateCoinUsedForRequest)(nil),  // 0: chain.gateway.coin.usedfor.v1.CreateCoinUsedForRequest
	(*CreateCoinUsedForResponse)(nil), // 1: chain.gateway.coin.usedfor.v1.CreateCoinUsedForResponse
	(*GetCoinUsedForsRequest)(nil),    // 2: chain.gateway.coin.usedfor.v1.GetCoinUsedForsRequest
	(*GetCoinUsedForsResponse)(nil),   // 3: chain.gateway.coin.usedfor.v1.GetCoinUsedForsResponse
	(*DeleteCoinUsedForRequest)(nil),  // 4: chain.gateway.coin.usedfor.v1.DeleteCoinUsedForRequest
	(*DeleteCoinUsedForResponse)(nil), // 5: chain.gateway.coin.usedfor.v1.DeleteCoinUsedForResponse
	(v1.CoinUsedFor)(0),               // 6: basetypes.chain.v1.CoinUsedFor
	(*usedfor.CoinUsedFor)(nil),       // 7: chain.middleware.coin.usedfor.v1.CoinUsedFor
}
var file_npool_chain_gw_v1_coin_usedfor_usedfor_proto_depIdxs = []int32{
	6, // 0: chain.gateway.coin.usedfor.v1.CreateCoinUsedForRequest.UsedFor:type_name -> basetypes.chain.v1.CoinUsedFor
	7, // 1: chain.gateway.coin.usedfor.v1.CreateCoinUsedForResponse.Info:type_name -> chain.middleware.coin.usedfor.v1.CoinUsedFor
	6, // 2: chain.gateway.coin.usedfor.v1.GetCoinUsedForsRequest.UsedFors:type_name -> basetypes.chain.v1.CoinUsedFor
	7, // 3: chain.gateway.coin.usedfor.v1.GetCoinUsedForsResponse.Infos:type_name -> chain.middleware.coin.usedfor.v1.CoinUsedFor
	7, // 4: chain.gateway.coin.usedfor.v1.DeleteCoinUsedForResponse.Info:type_name -> chain.middleware.coin.usedfor.v1.CoinUsedFor
	0, // 5: chain.gateway.coin.usedfor.v1.Gateway.CreateCoinUsedFor:input_type -> chain.gateway.coin.usedfor.v1.CreateCoinUsedForRequest
	2, // 6: chain.gateway.coin.usedfor.v1.Gateway.GetCoinUsedFors:input_type -> chain.gateway.coin.usedfor.v1.GetCoinUsedForsRequest
	4, // 7: chain.gateway.coin.usedfor.v1.Gateway.DeleteCoinUsedFor:input_type -> chain.gateway.coin.usedfor.v1.DeleteCoinUsedForRequest
	1, // 8: chain.gateway.coin.usedfor.v1.Gateway.CreateCoinUsedFor:output_type -> chain.gateway.coin.usedfor.v1.CreateCoinUsedForResponse
	3, // 9: chain.gateway.coin.usedfor.v1.Gateway.GetCoinUsedFors:output_type -> chain.gateway.coin.usedfor.v1.GetCoinUsedForsResponse
	5, // 10: chain.gateway.coin.usedfor.v1.Gateway.DeleteCoinUsedFor:output_type -> chain.gateway.coin.usedfor.v1.DeleteCoinUsedForResponse
	8, // [8:11] is the sub-list for method output_type
	5, // [5:8] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_npool_chain_gw_v1_coin_usedfor_usedfor_proto_init() }
func file_npool_chain_gw_v1_coin_usedfor_usedfor_proto_init() {
	if File_npool_chain_gw_v1_coin_usedfor_usedfor_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_npool_chain_gw_v1_coin_usedfor_usedfor_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateCoinUsedForRequest); i {
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
		file_npool_chain_gw_v1_coin_usedfor_usedfor_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateCoinUsedForResponse); i {
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
		file_npool_chain_gw_v1_coin_usedfor_usedfor_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCoinUsedForsRequest); i {
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
		file_npool_chain_gw_v1_coin_usedfor_usedfor_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCoinUsedForsResponse); i {
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
		file_npool_chain_gw_v1_coin_usedfor_usedfor_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteCoinUsedForRequest); i {
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
		file_npool_chain_gw_v1_coin_usedfor_usedfor_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteCoinUsedForResponse); i {
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
	file_npool_chain_gw_v1_coin_usedfor_usedfor_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_npool_chain_gw_v1_coin_usedfor_usedfor_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_npool_chain_gw_v1_coin_usedfor_usedfor_proto_goTypes,
		DependencyIndexes: file_npool_chain_gw_v1_coin_usedfor_usedfor_proto_depIdxs,
		MessageInfos:      file_npool_chain_gw_v1_coin_usedfor_usedfor_proto_msgTypes,
	}.Build()
	File_npool_chain_gw_v1_coin_usedfor_usedfor_proto = out.File
	file_npool_chain_gw_v1_coin_usedfor_usedfor_proto_rawDesc = nil
	file_npool_chain_gw_v1_coin_usedfor_usedfor_proto_goTypes = nil
	file_npool_chain_gw_v1_coin_usedfor_usedfor_proto_depIdxs = nil
}