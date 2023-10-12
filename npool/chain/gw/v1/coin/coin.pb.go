// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.18.1
// source: npool/chain/gw/v1/coin/coin.proto

package coin

import (
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

	Name string `protobuf:"bytes,10,opt,name=Name,proto3" json:"Name,omitempty"`
	Unit string `protobuf:"bytes,20,opt,name=Unit,proto3" json:"Unit,omitempty"`
	ENV  string `protobuf:"bytes,30,opt,name=ENV,proto3" json:"ENV,omitempty"`
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

func (x *CreateCoinRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateCoinRequest) GetUnit() string {
	if x != nil {
		return x.Unit
	}
	return ""
}

func (x *CreateCoinRequest) GetENV() string {
	if x != nil {
		return x.ENV
	}
	return ""
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

	Offset int32 `protobuf:"varint,10,opt,name=Offset,proto3" json:"Offset,omitempty"`
	Limit  int32 `protobuf:"varint,20,opt,name=Limit,proto3" json:"Limit,omitempty"`
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

	ID                          uint32  `protobuf:"varint,10,opt,name=ID,proto3" json:"ID,omitempty"`
	Presale                     *bool   `protobuf:"varint,20,opt,name=Presale,proto3,oneof" json:"Presale,omitempty"`
	ReservedAmount              *string `protobuf:"bytes,30,opt,name=ReservedAmount,proto3,oneof" json:"ReservedAmount,omitempty"`
	ForPay                      *bool   `protobuf:"varint,40,opt,name=ForPay,proto3,oneof" json:"ForPay,omitempty"`
	HomePage                    *string `protobuf:"bytes,50,opt,name=HomePage,proto3,oneof" json:"HomePage,omitempty"`
	Specs                       *string `protobuf:"bytes,60,opt,name=Specs,proto3,oneof" json:"Specs,omitempty"`
	FeeCoinTypeID               *string `protobuf:"bytes,70,opt,name=FeeCoinTypeID,proto3,oneof" json:"FeeCoinTypeID,omitempty"`
	WithdrawFeeByStableUSD      *bool   `protobuf:"varint,80,opt,name=WithdrawFeeByStableUSD,proto3,oneof" json:"WithdrawFeeByStableUSD,omitempty"`
	WithdrawFeeAmount           *string `protobuf:"bytes,90,opt,name=WithdrawFeeAmount,proto3,oneof" json:"WithdrawFeeAmount,omitempty"`
	CollectFeeAmount            *string `protobuf:"bytes,100,opt,name=CollectFeeAmount,proto3,oneof" json:"CollectFeeAmount,omitempty"`
	HotWalletFeeAmount          *string `protobuf:"bytes,110,opt,name=HotWalletFeeAmount,proto3,oneof" json:"HotWalletFeeAmount,omitempty"`
	LowFeeAmount                *string `protobuf:"bytes,120,opt,name=LowFeeAmount,proto3,oneof" json:"LowFeeAmount,omitempty"`
	HotLowFeeAmount             *string `protobuf:"bytes,121,opt,name=HotLowFeeAmount,proto3,oneof" json:"HotLowFeeAmount,omitempty"`
	HotWalletAccountAmount      *string `protobuf:"bytes,130,opt,name=HotWalletAccountAmount,proto3,oneof" json:"HotWalletAccountAmount,omitempty"`
	PaymentAccountCollectAmount *string `protobuf:"bytes,140,opt,name=PaymentAccountCollectAmount,proto3,oneof" json:"PaymentAccountCollectAmount,omitempty"`
	StableUSD                   *bool   `protobuf:"varint,150,opt,name=StableUSD,proto3,oneof" json:"StableUSD,omitempty"`
	Disabled                    *bool   `protobuf:"varint,160,opt,name=Disabled,proto3,oneof" json:"Disabled,omitempty"`
	LeastTransferAmount         *string `protobuf:"bytes,170,opt,name=LeastTransferAmount,proto3,oneof" json:"LeastTransferAmount,omitempty"`
	NeedMemo                    *bool   `protobuf:"varint,180,opt,name=NeedMemo,proto3,oneof" json:"NeedMemo,omitempty"`
	RefreshCurrency             *bool   `protobuf:"varint,190,opt,name=RefreshCurrency,proto3,oneof" json:"RefreshCurrency,omitempty"`
	Logo                        *string `protobuf:"bytes,200,opt,name=Logo,proto3,oneof" json:"Logo,omitempty"`
	CheckNewAddressBalance      *bool   `protobuf:"varint,210,opt,name=CheckNewAddressBalance,proto3,oneof" json:"CheckNewAddressBalance,omitempty"`
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

func (x *UpdateCoinRequest) GetID() uint32 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *UpdateCoinRequest) GetPresale() bool {
	if x != nil && x.Presale != nil {
		return *x.Presale
	}
	return false
}

func (x *UpdateCoinRequest) GetReservedAmount() string {
	if x != nil && x.ReservedAmount != nil {
		return *x.ReservedAmount
	}
	return ""
}

func (x *UpdateCoinRequest) GetForPay() bool {
	if x != nil && x.ForPay != nil {
		return *x.ForPay
	}
	return false
}

func (x *UpdateCoinRequest) GetHomePage() string {
	if x != nil && x.HomePage != nil {
		return *x.HomePage
	}
	return ""
}

func (x *UpdateCoinRequest) GetSpecs() string {
	if x != nil && x.Specs != nil {
		return *x.Specs
	}
	return ""
}

func (x *UpdateCoinRequest) GetFeeCoinTypeID() string {
	if x != nil && x.FeeCoinTypeID != nil {
		return *x.FeeCoinTypeID
	}
	return ""
}

func (x *UpdateCoinRequest) GetWithdrawFeeByStableUSD() bool {
	if x != nil && x.WithdrawFeeByStableUSD != nil {
		return *x.WithdrawFeeByStableUSD
	}
	return false
}

func (x *UpdateCoinRequest) GetWithdrawFeeAmount() string {
	if x != nil && x.WithdrawFeeAmount != nil {
		return *x.WithdrawFeeAmount
	}
	return ""
}

func (x *UpdateCoinRequest) GetCollectFeeAmount() string {
	if x != nil && x.CollectFeeAmount != nil {
		return *x.CollectFeeAmount
	}
	return ""
}

func (x *UpdateCoinRequest) GetHotWalletFeeAmount() string {
	if x != nil && x.HotWalletFeeAmount != nil {
		return *x.HotWalletFeeAmount
	}
	return ""
}

func (x *UpdateCoinRequest) GetLowFeeAmount() string {
	if x != nil && x.LowFeeAmount != nil {
		return *x.LowFeeAmount
	}
	return ""
}

func (x *UpdateCoinRequest) GetHotLowFeeAmount() string {
	if x != nil && x.HotLowFeeAmount != nil {
		return *x.HotLowFeeAmount
	}
	return ""
}

func (x *UpdateCoinRequest) GetHotWalletAccountAmount() string {
	if x != nil && x.HotWalletAccountAmount != nil {
		return *x.HotWalletAccountAmount
	}
	return ""
}

func (x *UpdateCoinRequest) GetPaymentAccountCollectAmount() string {
	if x != nil && x.PaymentAccountCollectAmount != nil {
		return *x.PaymentAccountCollectAmount
	}
	return ""
}

func (x *UpdateCoinRequest) GetStableUSD() bool {
	if x != nil && x.StableUSD != nil {
		return *x.StableUSD
	}
	return false
}

func (x *UpdateCoinRequest) GetDisabled() bool {
	if x != nil && x.Disabled != nil {
		return *x.Disabled
	}
	return false
}

func (x *UpdateCoinRequest) GetLeastTransferAmount() string {
	if x != nil && x.LeastTransferAmount != nil {
		return *x.LeastTransferAmount
	}
	return ""
}

func (x *UpdateCoinRequest) GetNeedMemo() bool {
	if x != nil && x.NeedMemo != nil {
		return *x.NeedMemo
	}
	return false
}

func (x *UpdateCoinRequest) GetRefreshCurrency() bool {
	if x != nil && x.RefreshCurrency != nil {
		return *x.RefreshCurrency
	}
	return false
}

func (x *UpdateCoinRequest) GetLogo() string {
	if x != nil && x.Logo != nil {
		return *x.Logo
	}
	return ""
}

func (x *UpdateCoinRequest) GetCheckNewAddressBalance() bool {
	if x != nil && x.CheckNewAddressBalance != nil {
		return *x.CheckNewAddressBalance
	}
	return false
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
	0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f,
	0x63, 0x68, 0x61, 0x69, 0x6e, 0x2f, 0x6d, 0x77, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x69, 0x6e,
	0x2f, 0x63, 0x6f, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4d, 0x0a, 0x11, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x55, 0x6e, 0x69, 0x74, 0x18, 0x14, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x55, 0x6e, 0x69, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x45, 0x4e, 0x56, 0x18,
	0x1e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x45, 0x4e, 0x56, 0x22, 0x48, 0x0a, 0x12, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x32, 0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e,
	0x2e, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72,
	0x65, 0x2e, 0x63, 0x6f, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x69, 0x6e, 0x52, 0x04,
	0x49, 0x6e, 0x66, 0x6f, 0x22, 0x3f, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x69, 0x6e, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x4f, 0x66, 0x66, 0x73, 0x65,
	0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12,
	0x14, 0x0a, 0x05, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x14, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05,
	0x4c, 0x69, 0x6d, 0x69, 0x74, 0x22, 0x5e, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x69, 0x6e,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x34, 0x0a, 0x05, 0x49, 0x6e, 0x66,
	0x6f, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x63, 0x68, 0x61, 0x69, 0x6e,
	0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x63, 0x6f, 0x69, 0x6e,
	0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x69, 0x6e, 0x52, 0x05, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x12,
	0x14, 0x0a, 0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x14, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05,
	0x54, 0x6f, 0x74, 0x61, 0x6c, 0x22, 0xd7, 0x0a, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x43, 0x6f, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49,
	0x44, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x49, 0x44, 0x12, 0x1d, 0x0a, 0x07, 0x50,
	0x72, 0x65, 0x73, 0x61, 0x6c, 0x65, 0x18, 0x14, 0x20, 0x01, 0x28, 0x08, 0x48, 0x00, 0x52, 0x07,
	0x50, 0x72, 0x65, 0x73, 0x61, 0x6c, 0x65, 0x88, 0x01, 0x01, 0x12, 0x2b, 0x0a, 0x0e, 0x52, 0x65,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x64, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x1e, 0x20, 0x01,
	0x28, 0x09, 0x48, 0x01, 0x52, 0x0e, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x64, 0x41, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x88, 0x01, 0x01, 0x12, 0x1b, 0x0a, 0x06, 0x46, 0x6f, 0x72, 0x50, 0x61,
	0x79, 0x18, 0x28, 0x20, 0x01, 0x28, 0x08, 0x48, 0x02, 0x52, 0x06, 0x46, 0x6f, 0x72, 0x50, 0x61,
	0x79, 0x88, 0x01, 0x01, 0x12, 0x1f, 0x0a, 0x08, 0x48, 0x6f, 0x6d, 0x65, 0x50, 0x61, 0x67, 0x65,
	0x18, 0x32, 0x20, 0x01, 0x28, 0x09, 0x48, 0x03, 0x52, 0x08, 0x48, 0x6f, 0x6d, 0x65, 0x50, 0x61,
	0x67, 0x65, 0x88, 0x01, 0x01, 0x12, 0x19, 0x0a, 0x05, 0x53, 0x70, 0x65, 0x63, 0x73, 0x18, 0x3c,
	0x20, 0x01, 0x28, 0x09, 0x48, 0x04, 0x52, 0x05, 0x53, 0x70, 0x65, 0x63, 0x73, 0x88, 0x01, 0x01,
	0x12, 0x29, 0x0a, 0x0d, 0x46, 0x65, 0x65, 0x43, 0x6f, 0x69, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x49,
	0x44, 0x18, 0x46, 0x20, 0x01, 0x28, 0x09, 0x48, 0x05, 0x52, 0x0d, 0x46, 0x65, 0x65, 0x43, 0x6f,
	0x69, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x49, 0x44, 0x88, 0x01, 0x01, 0x12, 0x3b, 0x0a, 0x16, 0x57,
	0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x46, 0x65, 0x65, 0x42, 0x79, 0x53, 0x74, 0x61, 0x62,
	0x6c, 0x65, 0x55, 0x53, 0x44, 0x18, 0x50, 0x20, 0x01, 0x28, 0x08, 0x48, 0x06, 0x52, 0x16, 0x57,
	0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x46, 0x65, 0x65, 0x42, 0x79, 0x53, 0x74, 0x61, 0x62,
	0x6c, 0x65, 0x55, 0x53, 0x44, 0x88, 0x01, 0x01, 0x12, 0x31, 0x0a, 0x11, 0x57, 0x69, 0x74, 0x68,
	0x64, 0x72, 0x61, 0x77, 0x46, 0x65, 0x65, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x5a, 0x20,
	0x01, 0x28, 0x09, 0x48, 0x07, 0x52, 0x11, 0x57, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x46,
	0x65, 0x65, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x88, 0x01, 0x01, 0x12, 0x2f, 0x0a, 0x10, 0x43,
	0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x46, 0x65, 0x65, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x64, 0x20, 0x01, 0x28, 0x09, 0x48, 0x08, 0x52, 0x10, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74,
	0x46, 0x65, 0x65, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x88, 0x01, 0x01, 0x12, 0x33, 0x0a, 0x12,
	0x48, 0x6f, 0x74, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x46, 0x65, 0x65, 0x41, 0x6d, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x6e, 0x20, 0x01, 0x28, 0x09, 0x48, 0x09, 0x52, 0x12, 0x48, 0x6f, 0x74, 0x57,
	0x61, 0x6c, 0x6c, 0x65, 0x74, 0x46, 0x65, 0x65, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x88, 0x01,
	0x01, 0x12, 0x27, 0x0a, 0x0c, 0x4c, 0x6f, 0x77, 0x46, 0x65, 0x65, 0x41, 0x6d, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x78, 0x20, 0x01, 0x28, 0x09, 0x48, 0x0a, 0x52, 0x0c, 0x4c, 0x6f, 0x77, 0x46, 0x65,
	0x65, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x88, 0x01, 0x01, 0x12, 0x2d, 0x0a, 0x0f, 0x48, 0x6f,
	0x74, 0x4c, 0x6f, 0x77, 0x46, 0x65, 0x65, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x79, 0x20,
	0x01, 0x28, 0x09, 0x48, 0x0b, 0x52, 0x0f, 0x48, 0x6f, 0x74, 0x4c, 0x6f, 0x77, 0x46, 0x65, 0x65,
	0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x88, 0x01, 0x01, 0x12, 0x3c, 0x0a, 0x16, 0x48, 0x6f, 0x74,
	0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x41, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x82, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x0c, 0x52, 0x16, 0x48, 0x6f,
	0x74, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x41, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x88, 0x01, 0x01, 0x12, 0x46, 0x0a, 0x1b, 0x50, 0x61, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74,
	0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x8c, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x0d, 0x52,
	0x1b, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x43,
	0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x88, 0x01, 0x01, 0x12,
	0x22, 0x0a, 0x09, 0x53, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x55, 0x53, 0x44, 0x18, 0x96, 0x01, 0x20,
	0x01, 0x28, 0x08, 0x48, 0x0e, 0x52, 0x09, 0x53, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x55, 0x53, 0x44,
	0x88, 0x01, 0x01, 0x12, 0x20, 0x0a, 0x08, 0x44, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18,
	0xa0, 0x01, 0x20, 0x01, 0x28, 0x08, 0x48, 0x0f, 0x52, 0x08, 0x44, 0x69, 0x73, 0x61, 0x62, 0x6c,
	0x65, 0x64, 0x88, 0x01, 0x01, 0x12, 0x36, 0x0a, 0x13, 0x4c, 0x65, 0x61, 0x73, 0x74, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0xaa, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x48, 0x10, 0x52, 0x13, 0x4c, 0x65, 0x61, 0x73, 0x74, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x66, 0x65, 0x72, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x88, 0x01, 0x01, 0x12, 0x20, 0x0a,
	0x08, 0x4e, 0x65, 0x65, 0x64, 0x4d, 0x65, 0x6d, 0x6f, 0x18, 0xb4, 0x01, 0x20, 0x01, 0x28, 0x08,
	0x48, 0x11, 0x52, 0x08, 0x4e, 0x65, 0x65, 0x64, 0x4d, 0x65, 0x6d, 0x6f, 0x88, 0x01, 0x01, 0x12,
	0x2e, 0x0a, 0x0f, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e,
	0x63, 0x79, 0x18, 0xbe, 0x01, 0x20, 0x01, 0x28, 0x08, 0x48, 0x12, 0x52, 0x0f, 0x52, 0x65, 0x66,
	0x72, 0x65, 0x73, 0x68, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x88, 0x01, 0x01, 0x12,
	0x18, 0x0a, 0x04, 0x4c, 0x6f, 0x67, 0x6f, 0x18, 0xc8, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x13,
	0x52, 0x04, 0x4c, 0x6f, 0x67, 0x6f, 0x88, 0x01, 0x01, 0x12, 0x3c, 0x0a, 0x16, 0x43, 0x68, 0x65,
	0x63, 0x6b, 0x4e, 0x65, 0x77, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x42, 0x61, 0x6c, 0x61,
	0x6e, 0x63, 0x65, 0x18, 0xd2, 0x01, 0x20, 0x01, 0x28, 0x08, 0x48, 0x14, 0x52, 0x16, 0x43, 0x68,
	0x65, 0x63, 0x6b, 0x4e, 0x65, 0x77, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x42, 0x61, 0x6c,
	0x61, 0x6e, 0x63, 0x65, 0x88, 0x01, 0x01, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x50, 0x72, 0x65, 0x73,
	0x61, 0x6c, 0x65, 0x42, 0x11, 0x0a, 0x0f, 0x5f, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x64,
	0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x46, 0x6f, 0x72, 0x50, 0x61,
	0x79, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x48, 0x6f, 0x6d, 0x65, 0x50, 0x61, 0x67, 0x65, 0x42, 0x08,
	0x0a, 0x06, 0x5f, 0x53, 0x70, 0x65, 0x63, 0x73, 0x42, 0x10, 0x0a, 0x0e, 0x5f, 0x46, 0x65, 0x65,
	0x43, 0x6f, 0x69, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x49, 0x44, 0x42, 0x19, 0x0a, 0x17, 0x5f, 0x57,
	0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x46, 0x65, 0x65, 0x42, 0x79, 0x53, 0x74, 0x61, 0x62,
	0x6c, 0x65, 0x55, 0x53, 0x44, 0x42, 0x14, 0x0a, 0x12, 0x5f, 0x57, 0x69, 0x74, 0x68, 0x64, 0x72,
	0x61, 0x77, 0x46, 0x65, 0x65, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x13, 0x0a, 0x11, 0x5f,
	0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x46, 0x65, 0x65, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x42, 0x15, 0x0a, 0x13, 0x5f, 0x48, 0x6f, 0x74, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x46, 0x65,
	0x65, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x0f, 0x0a, 0x0d, 0x5f, 0x4c, 0x6f, 0x77, 0x46,
	0x65, 0x65, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x12, 0x0a, 0x10, 0x5f, 0x48, 0x6f, 0x74,
	0x4c, 0x6f, 0x77, 0x46, 0x65, 0x65, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x19, 0x0a, 0x17,
	0x5f, 0x48, 0x6f, 0x74, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x1e, 0x0a, 0x1c, 0x5f, 0x50, 0x61, 0x79, 0x6d,
	0x65, 0x6e, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63,
	0x74, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x53, 0x74, 0x61, 0x62,
	0x6c, 0x65, 0x55, 0x53, 0x44, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x44, 0x69, 0x73, 0x61, 0x62, 0x6c,
	0x65, 0x64, 0x42, 0x16, 0x0a, 0x14, 0x5f, 0x4c, 0x65, 0x61, 0x73, 0x74, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x66, 0x65, 0x72, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x4e,
	0x65, 0x65, 0x64, 0x4d, 0x65, 0x6d, 0x6f, 0x42, 0x12, 0x0a, 0x10, 0x5f, 0x52, 0x65, 0x66, 0x72,
	0x65, 0x73, 0x68, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x42, 0x07, 0x0a, 0x05, 0x5f,
	0x4c, 0x6f, 0x67, 0x6f, 0x42, 0x19, 0x0a, 0x17, 0x5f, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x4e, 0x65,
	0x77, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x22,
	0x48, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x69, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x32, 0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x6d, 0x69, 0x64, 0x64,
	0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x63, 0x6f, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x43,
	0x6f, 0x69, 0x6e, 0x52, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x32, 0xfe, 0x02, 0x0a, 0x07, 0x47, 0x61,
	0x74, 0x65, 0x77, 0x61, 0x79, 0x12, 0x7d, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43,
	0x6f, 0x69, 0x6e, 0x12, 0x28, 0x2e, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x67, 0x61, 0x74, 0x65,
	0x77, 0x61, 0x79, 0x2e, 0x63, 0x6f, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x43, 0x6f, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e,
	0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x63, 0x6f,
	0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x69, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14,
	0x22, 0x0f, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x2f, 0x63, 0x6f, 0x69,
	0x6e, 0x3a, 0x01, 0x2a, 0x12, 0x75, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x69, 0x6e, 0x73,
	0x12, 0x26, 0x2e, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79,
	0x2e, 0x63, 0x6f, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x69, 0x6e,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x63, 0x68, 0x61, 0x69, 0x6e,
	0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x63, 0x6f, 0x69, 0x6e, 0x2e, 0x76, 0x31,
	0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x69, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x18, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x12, 0x22, 0x0d, 0x2f, 0x76, 0x31, 0x2f, 0x67,
	0x65, 0x74, 0x2f, 0x63, 0x6f, 0x69, 0x6e, 0x73, 0x3a, 0x01, 0x2a, 0x12, 0x7d, 0x0a, 0x0a, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x69, 0x6e, 0x12, 0x28, 0x2e, 0x63, 0x68, 0x61, 0x69,
	0x6e, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x63, 0x6f, 0x69, 0x6e, 0x2e, 0x76,
	0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x67, 0x61, 0x74, 0x65,
	0x77, 0x61, 0x79, 0x2e, 0x63, 0x6f, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x43, 0x6f, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1a,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14, 0x22, 0x0f, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x2f, 0x63, 0x6f, 0x69, 0x6e, 0x3a, 0x01, 0x2a, 0x42, 0x39, 0x5a, 0x37, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4e, 0x70, 0x6f, 0x6f, 0x6c, 0x50, 0x6c,
	0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x6e,
	0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2f, 0x67, 0x77, 0x2f, 0x76, 0x31,
	0x2f, 0x63, 0x6f, 0x69, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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
	(*coin.Coin)(nil),          // 6: chain.middleware.coin.v1.Coin
}
var file_npool_chain_gw_v1_coin_coin_proto_depIdxs = []int32{
	6, // 0: chain.gateway.coin.v1.CreateCoinResponse.Info:type_name -> chain.middleware.coin.v1.Coin
	6, // 1: chain.gateway.coin.v1.GetCoinsResponse.Infos:type_name -> chain.middleware.coin.v1.Coin
	6, // 2: chain.gateway.coin.v1.UpdateCoinResponse.Info:type_name -> chain.middleware.coin.v1.Coin
	0, // 3: chain.gateway.coin.v1.Gateway.CreateCoin:input_type -> chain.gateway.coin.v1.CreateCoinRequest
	2, // 4: chain.gateway.coin.v1.Gateway.GetCoins:input_type -> chain.gateway.coin.v1.GetCoinsRequest
	4, // 5: chain.gateway.coin.v1.Gateway.UpdateCoin:input_type -> chain.gateway.coin.v1.UpdateCoinRequest
	1, // 6: chain.gateway.coin.v1.Gateway.CreateCoin:output_type -> chain.gateway.coin.v1.CreateCoinResponse
	3, // 7: chain.gateway.coin.v1.Gateway.GetCoins:output_type -> chain.gateway.coin.v1.GetCoinsResponse
	5, // 8: chain.gateway.coin.v1.Gateway.UpdateCoin:output_type -> chain.gateway.coin.v1.UpdateCoinResponse
	6, // [6:9] is the sub-list for method output_type
	3, // [3:6] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
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
	file_npool_chain_gw_v1_coin_coin_proto_msgTypes[4].OneofWrappers = []interface{}{}
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
