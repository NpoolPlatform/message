// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.18.1
// source: npool/chain/gw/v1/coin/coin.proto

package coin

import (
	_ "github.com/NpoolPlatform/message/npool"
	coin "github.com/NpoolPlatform/message/npool/chain/mw/v1/coin"
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

type CreateCoinRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *coin.CoinReq `protobuf:"bytes,10,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (x *CreateCoinRequest) Reset() {
	*x = CreateCoinRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_chain_gw_v1_coin_coin_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCoinRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCoinRequest) ProtoMessage() {}

func (x *CreateCoinRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_chain_gw_v1_coin_coin_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCoinRequest.ProtoReflect.Descriptor instead.
func (*CreateCoinRequest) Descriptor() ([]byte, []int) {
	return file_npool_chain_gw_v1_coin_coin_proto_rawDescGZIP(), []int{0}
}

func (x *CreateCoinRequest) GetInfo() *coin.CoinReq {
	if x != nil {
		return x.Info
	}
	return nil
}

type CreateCoinResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *coin.Coin `protobuf:"bytes,10,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (x *CreateCoinResponse) Reset() {
	*x = CreateCoinResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_chain_gw_v1_coin_coin_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCoinResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCoinResponse) ProtoMessage() {}

func (x *CreateCoinResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_chain_gw_v1_coin_coin_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCoinResponse.ProtoReflect.Descriptor instead.
func (*CreateCoinResponse) Descriptor() ([]byte, []int) {
	return file_npool_chain_gw_v1_coin_coin_proto_rawDescGZIP(), []int{1}
}

func (x *CreateCoinResponse) GetInfo() *coin.Coin {
	if x != nil {
		return x.Info
	}
	return nil
}

type GetCoinsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Conds  *coin.Conds `protobuf:"bytes,10,opt,name=Conds,proto3" json:"Conds,omitempty"`
	Offset int32       `protobuf:"varint,20,opt,name=Offset,proto3" json:"Offset,omitempty"`
	Limit  int32       `protobuf:"varint,30,opt,name=Limit,proto3" json:"Limit,omitempty"`
}

func (x *GetCoinsRequest) Reset() {
	*x = GetCoinsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_chain_gw_v1_coin_coin_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCoinsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCoinsRequest) ProtoMessage() {}

func (x *GetCoinsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_chain_gw_v1_coin_coin_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCoinsRequest.ProtoReflect.Descriptor instead.
func (*GetCoinsRequest) Descriptor() ([]byte, []int) {
	return file_npool_chain_gw_v1_coin_coin_proto_rawDescGZIP(), []int{2}
}

func (x *GetCoinsRequest) GetConds() *coin.Conds {
	if x != nil {
		return x.Conds
	}
	return nil
}

func (x *GetCoinsRequest) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *GetCoinsRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type GetCoinsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Infos []*coin.Coin `protobuf:"bytes,10,rep,name=Infos,proto3" json:"Infos,omitempty"`
	Total uint32       `protobuf:"varint,20,opt,name=Total,proto3" json:"Total,omitempty"`
}

func (x *GetCoinsResponse) Reset() {
	*x = GetCoinsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_chain_gw_v1_coin_coin_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCoinsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCoinsResponse) ProtoMessage() {}

func (x *GetCoinsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_chain_gw_v1_coin_coin_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCoinsResponse.ProtoReflect.Descriptor instead.
func (*GetCoinsResponse) Descriptor() ([]byte, []int) {
	return file_npool_chain_gw_v1_coin_coin_proto_rawDescGZIP(), []int{3}
}

func (x *GetCoinsResponse) GetInfos() []*coin.Coin {
	if x != nil {
		return x.Infos
	}
	return nil
}

func (x *GetCoinsResponse) GetTotal() uint32 {
	if x != nil {
		return x.Total
	}
	return 0
}

type UpdateCoinRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *coin.CoinReq `protobuf:"bytes,10,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (x *UpdateCoinRequest) Reset() {
	*x = UpdateCoinRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_chain_gw_v1_coin_coin_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateCoinRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateCoinRequest) ProtoMessage() {}

func (x *UpdateCoinRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_chain_gw_v1_coin_coin_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateCoinRequest.ProtoReflect.Descriptor instead.
func (*UpdateCoinRequest) Descriptor() ([]byte, []int) {
	return file_npool_chain_gw_v1_coin_coin_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateCoinRequest) GetInfo() *coin.CoinReq {
	if x != nil {
		return x.Info
	}
	return nil
}

type UpdateCoinResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *coin.Coin `protobuf:"bytes,10,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (x *UpdateCoinResponse) Reset() {
	*x = UpdateCoinResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_chain_gw_v1_coin_coin_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateCoinResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateCoinResponse) ProtoMessage() {}

func (x *UpdateCoinResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_chain_gw_v1_coin_coin_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateCoinResponse.ProtoReflect.Descriptor instead.
func (*UpdateCoinResponse) Descriptor() ([]byte, []int) {
	return file_npool_chain_gw_v1_coin_coin_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateCoinResponse) GetInfo() *coin.Coin {
	if x != nil {
		return x.Info
	}
	return nil
}

var File_npool_chain_gw_v1_coin_coin_proto protoreflect.FileDescriptor

var file_npool_chain_gw_v1_coin_coin_proto_rawDesc = []byte{
	0x0a, 0x21, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2f, 0x67, 0x77,
	0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x69, 0x6e, 0x2f, 0x63, 0x6f, 0x69, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x15, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77,
	0x61, 0x79, 0x2e, 0x63, 0x6f, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f,
	0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x6e, 0x70, 0x6f,
	0x6f, 0x6c, 0x2f, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2f, 0x6d, 0x77, 0x2f, 0x76, 0x31, 0x2f, 0x63,
	0x6f, 0x69, 0x6e, 0x2f, 0x63, 0x6f, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4a,
	0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x35, 0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x21, 0x2e, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65,
	0x77, 0x61, 0x72, 0x65, 0x2e, 0x63, 0x6f, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x69,
	0x6e, 0x52, 0x65, 0x71, 0x52, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x48, 0x0a, 0x12, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x32, 0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e,
	0x2e, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72,
	0x65, 0x2e, 0x63, 0x6f, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x69, 0x6e, 0x52, 0x04,
	0x49, 0x6e, 0x66, 0x6f, 0x22, 0x76, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x69, 0x6e, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x35, 0x0a, 0x05, 0x43, 0x6f, 0x6e, 0x64, 0x73,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x6d,
	0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x63, 0x6f, 0x69, 0x6e, 0x2e, 0x76,
	0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x64, 0x73, 0x52, 0x05, 0x43, 0x6f, 0x6e, 0x64, 0x73, 0x12, 0x16,
	0x0a, 0x06, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x14, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06,
	0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x18,
	0x1e, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x22, 0x5e, 0x0a, 0x10,
	0x47, 0x65, 0x74, 0x43, 0x6f, 0x69, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x34, 0x0a, 0x05, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x1e, 0x2e, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61,
	0x72, 0x65, 0x2e, 0x63, 0x6f, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x69, 0x6e, 0x52,
	0x05, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x18,
	0x14, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x22, 0x4a, 0x0a, 0x11,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x35, 0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x21, 0x2e, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61,
	0x72, 0x65, 0x2e, 0x63, 0x6f, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x69, 0x6e, 0x52,
	0x65, 0x71, 0x52, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x48, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x43, 0x6f, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x32,
	0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x63,
	0x68, 0x61, 0x69, 0x6e, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e,
	0x63, 0x6f, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x69, 0x6e, 0x52, 0x04, 0x49, 0x6e,
	0x66, 0x6f, 0x32, 0xfe, 0x02, 0x0a, 0x07, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x12, 0x7d,
	0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x69, 0x6e, 0x12, 0x28, 0x2e, 0x63,
	0x68, 0x61, 0x69, 0x6e, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x63, 0x6f, 0x69,
	0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x69, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x67,
	0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x63, 0x6f, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x1a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14, 0x22, 0x0f, 0x2f, 0x76, 0x31, 0x2f, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x2f, 0x63, 0x6f, 0x69, 0x6e, 0x3a, 0x01, 0x2a, 0x12, 0x75, 0x0a,
	0x08, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x69, 0x6e, 0x73, 0x12, 0x26, 0x2e, 0x63, 0x68, 0x61, 0x69,
	0x6e, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x63, 0x6f, 0x69, 0x6e, 0x2e, 0x76,
	0x31, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x69, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x27, 0x2e, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61,
	0x79, 0x2e, 0x63, 0x6f, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x69,
	0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x18, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x12, 0x22, 0x0d, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x65, 0x74, 0x2f, 0x63, 0x6f, 0x69, 0x6e,
	0x73, 0x3a, 0x01, 0x2a, 0x12, 0x7d, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f,
	0x69, 0x6e, 0x12, 0x28, 0x2e, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77,
	0x61, 0x79, 0x2e, 0x63, 0x6f, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x43, 0x6f, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x63,
	0x68, 0x61, 0x69, 0x6e, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x63, 0x6f, 0x69,
	0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x69, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14, 0x22,
	0x0f, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x63, 0x6f, 0x69, 0x6e,
	0x3a, 0x01, 0x2a, 0x42, 0x39, 0x5a, 0x37, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x4e, 0x70, 0x6f, 0x6f, 0x6c, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x63, 0x68,
	0x61, 0x69, 0x6e, 0x2f, 0x67, 0x77, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x69, 0x6e, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_npool_chain_gw_v1_coin_coin_proto_rawDescOnce sync.Once
	file_npool_chain_gw_v1_coin_coin_proto_rawDescData = file_npool_chain_gw_v1_coin_coin_proto_rawDesc
)

func file_npool_chain_gw_v1_coin_coin_proto_rawDescGZIP() []byte {
	file_npool_chain_gw_v1_coin_coin_proto_rawDescOnce.Do(func() {
		file_npool_chain_gw_v1_coin_coin_proto_rawDescData = protoimpl.X.CompressGZIP(file_npool_chain_gw_v1_coin_coin_proto_rawDescData)
	})
	return file_npool_chain_gw_v1_coin_coin_proto_rawDescData
}

var file_npool_chain_gw_v1_coin_coin_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_npool_chain_gw_v1_coin_coin_proto_goTypes = []interface{}{
	(*CreateCoinRequest)(nil),  // 0: chain.gateway.coin.v1.CreateCoinRequest
	(*CreateCoinResponse)(nil), // 1: chain.gateway.coin.v1.CreateCoinResponse
	(*GetCoinsRequest)(nil),    // 2: chain.gateway.coin.v1.GetCoinsRequest
	(*GetCoinsResponse)(nil),   // 3: chain.gateway.coin.v1.GetCoinsResponse
	(*UpdateCoinRequest)(nil),  // 4: chain.gateway.coin.v1.UpdateCoinRequest
	(*UpdateCoinResponse)(nil), // 5: chain.gateway.coin.v1.UpdateCoinResponse
	(*coin.CoinReq)(nil),       // 6: chain.middleware.coin.v1.CoinReq
	(*coin.Coin)(nil),          // 7: chain.middleware.coin.v1.Coin
	(*coin.Conds)(nil),         // 8: chain.middleware.coin.v1.Conds
}
var file_npool_chain_gw_v1_coin_coin_proto_depIdxs = []int32{
	6, // 0: chain.gateway.coin.v1.CreateCoinRequest.Info:type_name -> chain.middleware.coin.v1.CoinReq
	7, // 1: chain.gateway.coin.v1.CreateCoinResponse.Info:type_name -> chain.middleware.coin.v1.Coin
	8, // 2: chain.gateway.coin.v1.GetCoinsRequest.Conds:type_name -> chain.middleware.coin.v1.Conds
	7, // 3: chain.gateway.coin.v1.GetCoinsResponse.Infos:type_name -> chain.middleware.coin.v1.Coin
	6, // 4: chain.gateway.coin.v1.UpdateCoinRequest.Info:type_name -> chain.middleware.coin.v1.CoinReq
	7, // 5: chain.gateway.coin.v1.UpdateCoinResponse.Info:type_name -> chain.middleware.coin.v1.Coin
	0, // 6: chain.gateway.coin.v1.Gateway.CreateCoin:input_type -> chain.gateway.coin.v1.CreateCoinRequest
	2, // 7: chain.gateway.coin.v1.Gateway.GetCoins:input_type -> chain.gateway.coin.v1.GetCoinsRequest
	4, // 8: chain.gateway.coin.v1.Gateway.UpdateCoin:input_type -> chain.gateway.coin.v1.UpdateCoinRequest
	1, // 9: chain.gateway.coin.v1.Gateway.CreateCoin:output_type -> chain.gateway.coin.v1.CreateCoinResponse
	3, // 10: chain.gateway.coin.v1.Gateway.GetCoins:output_type -> chain.gateway.coin.v1.GetCoinsResponse
	5, // 11: chain.gateway.coin.v1.Gateway.UpdateCoin:output_type -> chain.gateway.coin.v1.UpdateCoinResponse
	9, // [9:12] is the sub-list for method output_type
	6, // [6:9] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_npool_chain_gw_v1_coin_coin_proto_init() }
func file_npool_chain_gw_v1_coin_coin_proto_init() {
	if File_npool_chain_gw_v1_coin_coin_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_npool_chain_gw_v1_coin_coin_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateCoinRequest); i {
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
		file_npool_chain_gw_v1_coin_coin_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateCoinResponse); i {
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
		file_npool_chain_gw_v1_coin_coin_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCoinsRequest); i {
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
		file_npool_chain_gw_v1_coin_coin_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCoinsResponse); i {
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
		file_npool_chain_gw_v1_coin_coin_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateCoinRequest); i {
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
		file_npool_chain_gw_v1_coin_coin_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateCoinResponse); i {
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
			RawDescriptor: file_npool_chain_gw_v1_coin_coin_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_npool_chain_gw_v1_coin_coin_proto_goTypes,
		DependencyIndexes: file_npool_chain_gw_v1_coin_coin_proto_depIdxs,
		MessageInfos:      file_npool_chain_gw_v1_coin_coin_proto_msgTypes,
	}.Build()
	File_npool_chain_gw_v1_coin_coin_proto = out.File
	file_npool_chain_gw_v1_coin_coin_proto_rawDesc = nil
	file_npool_chain_gw_v1_coin_coin_proto_goTypes = nil
	file_npool_chain_gw_v1_coin_coin_proto_depIdxs = nil
}
