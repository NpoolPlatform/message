// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.18.1
// source: npool/inspire/gw/v1/coupon/app/coin/coin.proto

package coin

import (
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

type CouponCoin struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID                 uint32 `protobuf:"varint,10,opt,name=ID,proto3" json:"ID,omitempty"`
	EntID              string `protobuf:"bytes,20,opt,name=EntID,proto3" json:"EntID,omitempty"`
	AppID              string `protobuf:"bytes,30,opt,name=AppID,proto3" json:"AppID,omitempty"`
	CouponID           string `protobuf:"bytes,40,opt,name=CouponID,proto3" json:"CouponID,omitempty"`
	CouponName         string `protobuf:"bytes,50,opt,name=CouponName,proto3" json:"CouponName,omitempty"`
	CouponDenomination string `protobuf:"bytes,60,opt,name=CouponDenomination,proto3" json:"CouponDenomination,omitempty"`
	CoinTypeID         string `protobuf:"bytes,70,opt,name=CoinTypeID,proto3" json:"CoinTypeID,omitempty"`
	CoinName           string `protobuf:"bytes,80,opt,name=CoinName,proto3" json:"CoinName,omitempty"`
	CoinENV            string `protobuf:"bytes,90,opt,name=CoinENV,proto3" json:"CoinENV,omitempty"`
	CreatedAt          uint32 `protobuf:"varint,1000,opt,name=CreatedAt,proto3" json:"CreatedAt,omitempty"`
	UpdatedAt          uint32 `protobuf:"varint,1010,opt,name=UpdatedAt,proto3" json:"UpdatedAt,omitempty"`
}

func (x *CouponCoin) Reset() {
	*x = CouponCoin{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_inspire_gw_v1_coupon_app_coin_coin_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CouponCoin) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CouponCoin) ProtoMessage() {}

func (x *CouponCoin) ProtoReflect() protoreflect.Message {
	mi := &file_npool_inspire_gw_v1_coupon_app_coin_coin_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CouponCoin.ProtoReflect.Descriptor instead.
func (*CouponCoin) Descriptor() ([]byte, []int) {
	return file_npool_inspire_gw_v1_coupon_app_coin_coin_proto_rawDescGZIP(), []int{0}
}

func (x *CouponCoin) GetID() uint32 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *CouponCoin) GetEntID() string {
	if x != nil {
		return x.EntID
	}
	return ""
}

func (x *CouponCoin) GetAppID() string {
	if x != nil {
		return x.AppID
	}
	return ""
}

func (x *CouponCoin) GetCouponID() string {
	if x != nil {
		return x.CouponID
	}
	return ""
}

func (x *CouponCoin) GetCouponName() string {
	if x != nil {
		return x.CouponName
	}
	return ""
}

func (x *CouponCoin) GetCouponDenomination() string {
	if x != nil {
		return x.CouponDenomination
	}
	return ""
}

func (x *CouponCoin) GetCoinTypeID() string {
	if x != nil {
		return x.CoinTypeID
	}
	return ""
}

func (x *CouponCoin) GetCoinName() string {
	if x != nil {
		return x.CoinName
	}
	return ""
}

func (x *CouponCoin) GetCoinENV() string {
	if x != nil {
		return x.CoinENV
	}
	return ""
}

func (x *CouponCoin) GetCreatedAt() uint32 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *CouponCoin) GetUpdatedAt() uint32 {
	if x != nil {
		return x.UpdatedAt
	}
	return 0
}

type GetCouponCoinsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AppID  string `protobuf:"bytes,10,opt,name=AppID,proto3" json:"AppID,omitempty"`
	Offset int32  `protobuf:"varint,20,opt,name=Offset,proto3" json:"Offset,omitempty"`
	Limit  int32  `protobuf:"varint,30,opt,name=Limit,proto3" json:"Limit,omitempty"`
}

func (x *GetCouponCoinsRequest) Reset() {
	*x = GetCouponCoinsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_inspire_gw_v1_coupon_app_coin_coin_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCouponCoinsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCouponCoinsRequest) ProtoMessage() {}

func (x *GetCouponCoinsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_inspire_gw_v1_coupon_app_coin_coin_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCouponCoinsRequest.ProtoReflect.Descriptor instead.
func (*GetCouponCoinsRequest) Descriptor() ([]byte, []int) {
	return file_npool_inspire_gw_v1_coupon_app_coin_coin_proto_rawDescGZIP(), []int{1}
}

func (x *GetCouponCoinsRequest) GetAppID() string {
	if x != nil {
		return x.AppID
	}
	return ""
}

func (x *GetCouponCoinsRequest) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *GetCouponCoinsRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type GetCouponCoinsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Infos []*CouponCoin `protobuf:"bytes,10,rep,name=Infos,proto3" json:"Infos,omitempty"`
	Total uint32        `protobuf:"varint,20,opt,name=Total,proto3" json:"Total,omitempty"`
}

func (x *GetCouponCoinsResponse) Reset() {
	*x = GetCouponCoinsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_inspire_gw_v1_coupon_app_coin_coin_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCouponCoinsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCouponCoinsResponse) ProtoMessage() {}

func (x *GetCouponCoinsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_inspire_gw_v1_coupon_app_coin_coin_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCouponCoinsResponse.ProtoReflect.Descriptor instead.
func (*GetCouponCoinsResponse) Descriptor() ([]byte, []int) {
	return file_npool_inspire_gw_v1_coupon_app_coin_coin_proto_rawDescGZIP(), []int{2}
}

func (x *GetCouponCoinsResponse) GetInfos() []*CouponCoin {
	if x != nil {
		return x.Infos
	}
	return nil
}

func (x *GetCouponCoinsResponse) GetTotal() uint32 {
	if x != nil {
		return x.Total
	}
	return 0
}

type CreateCouponCoinRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AppID      string `protobuf:"bytes,10,opt,name=AppID,proto3" json:"AppID,omitempty"`
	CouponID   string `protobuf:"bytes,20,opt,name=CouponID,proto3" json:"CouponID,omitempty"`
	CoinTypeID string `protobuf:"bytes,30,opt,name=CoinTypeID,proto3" json:"CoinTypeID,omitempty"`
}

func (x *CreateCouponCoinRequest) Reset() {
	*x = CreateCouponCoinRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_inspire_gw_v1_coupon_app_coin_coin_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCouponCoinRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCouponCoinRequest) ProtoMessage() {}

func (x *CreateCouponCoinRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_inspire_gw_v1_coupon_app_coin_coin_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCouponCoinRequest.ProtoReflect.Descriptor instead.
func (*CreateCouponCoinRequest) Descriptor() ([]byte, []int) {
	return file_npool_inspire_gw_v1_coupon_app_coin_coin_proto_rawDescGZIP(), []int{3}
}

func (x *CreateCouponCoinRequest) GetAppID() string {
	if x != nil {
		return x.AppID
	}
	return ""
}

func (x *CreateCouponCoinRequest) GetCouponID() string {
	if x != nil {
		return x.CouponID
	}
	return ""
}

func (x *CreateCouponCoinRequest) GetCoinTypeID() string {
	if x != nil {
		return x.CoinTypeID
	}
	return ""
}

type CreateCouponCoinResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *CouponCoin `protobuf:"bytes,10,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (x *CreateCouponCoinResponse) Reset() {
	*x = CreateCouponCoinResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_inspire_gw_v1_coupon_app_coin_coin_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCouponCoinResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCouponCoinResponse) ProtoMessage() {}

func (x *CreateCouponCoinResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_inspire_gw_v1_coupon_app_coin_coin_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCouponCoinResponse.ProtoReflect.Descriptor instead.
func (*CreateCouponCoinResponse) Descriptor() ([]byte, []int) {
	return file_npool_inspire_gw_v1_coupon_app_coin_coin_proto_rawDescGZIP(), []int{4}
}

func (x *CreateCouponCoinResponse) GetInfo() *CouponCoin {
	if x != nil {
		return x.Info
	}
	return nil
}

type DeleteCouponCoinRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID    uint32 `protobuf:"varint,10,opt,name=ID,proto3" json:"ID,omitempty"`
	EntID string `protobuf:"bytes,20,opt,name=EntID,proto3" json:"EntID,omitempty"`
	AppID string `protobuf:"bytes,30,opt,name=AppID,proto3" json:"AppID,omitempty"`
}

func (x *DeleteCouponCoinRequest) Reset() {
	*x = DeleteCouponCoinRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_inspire_gw_v1_coupon_app_coin_coin_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteCouponCoinRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteCouponCoinRequest) ProtoMessage() {}

func (x *DeleteCouponCoinRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_inspire_gw_v1_coupon_app_coin_coin_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteCouponCoinRequest.ProtoReflect.Descriptor instead.
func (*DeleteCouponCoinRequest) Descriptor() ([]byte, []int) {
	return file_npool_inspire_gw_v1_coupon_app_coin_coin_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteCouponCoinRequest) GetID() uint32 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *DeleteCouponCoinRequest) GetEntID() string {
	if x != nil {
		return x.EntID
	}
	return ""
}

func (x *DeleteCouponCoinRequest) GetAppID() string {
	if x != nil {
		return x.AppID
	}
	return ""
}

type DeleteCouponCoinResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *CouponCoin `protobuf:"bytes,10,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (x *DeleteCouponCoinResponse) Reset() {
	*x = DeleteCouponCoinResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_inspire_gw_v1_coupon_app_coin_coin_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteCouponCoinResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteCouponCoinResponse) ProtoMessage() {}

func (x *DeleteCouponCoinResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_inspire_gw_v1_coupon_app_coin_coin_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteCouponCoinResponse.ProtoReflect.Descriptor instead.
func (*DeleteCouponCoinResponse) Descriptor() ([]byte, []int) {
	return file_npool_inspire_gw_v1_coupon_app_coin_coin_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteCouponCoinResponse) GetInfo() *CouponCoin {
	if x != nil {
		return x.Info
	}
	return nil
}

var File_npool_inspire_gw_v1_coupon_app_coin_coin_proto protoreflect.FileDescriptor

var file_npool_inspire_gw_v1_coupon_app_coin_coin_proto_rawDesc = []byte{
	0x0a, 0x2e, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x69, 0x6e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2f,
	0x67, 0x77, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x70,
	0x2f, 0x63, 0x6f, 0x69, 0x6e, 0x2f, 0x63, 0x6f, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x22, 0x69, 0x6e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61,
	0x79, 0x2e, 0x63, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x63, 0x6f, 0x69,
	0x6e, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xc8, 0x02, 0x0a, 0x0a, 0x43, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x43, 0x6f, 0x69,
	0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x49,
	0x44, 0x12, 0x14, 0x0a, 0x05, 0x45, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x45, 0x6e, 0x74, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x41, 0x70, 0x70, 0x49, 0x44,
	0x18, 0x1e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x41, 0x70, 0x70, 0x49, 0x44, 0x12, 0x1a, 0x0a,
	0x08, 0x43, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x49, 0x44, 0x18, 0x28, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x43, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x49, 0x44, 0x12, 0x1e, 0x0a, 0x0a, 0x43, 0x6f, 0x75,
	0x70, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x32, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x43,
	0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x2e, 0x0a, 0x12, 0x43, 0x6f, 0x75,
	0x70, 0x6f, 0x6e, 0x44, 0x65, 0x6e, 0x6f, 0x6d, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x3c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x43, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x44, 0x65, 0x6e,
	0x6f, 0x6d, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1e, 0x0a, 0x0a, 0x43, 0x6f, 0x69,
	0x6e, 0x54, 0x79, 0x70, 0x65, 0x49, 0x44, 0x18, 0x46, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x43,
	0x6f, 0x69, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08, 0x43, 0x6f, 0x69,
	0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x50, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x43, 0x6f, 0x69,
	0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x43, 0x6f, 0x69, 0x6e, 0x45, 0x4e, 0x56,
	0x18, 0x5a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x43, 0x6f, 0x69, 0x6e, 0x45, 0x4e, 0x56, 0x12,
	0x1d, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0xe8, 0x07, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d,
	0x0a, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0xf2, 0x07, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x5b, 0x0a,
	0x15, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x43, 0x6f, 0x69, 0x6e, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x41, 0x70, 0x70, 0x49, 0x44, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x41, 0x70, 0x70, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06,
	0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x14, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x4f, 0x66,
	0x66, 0x73, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x1e, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x05, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x22, 0x74, 0x0a, 0x16, 0x47, 0x65,
	0x74, 0x43, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x43, 0x6f, 0x69, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x44, 0x0a, 0x05, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x18, 0x0a, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x69, 0x6e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x67, 0x61,
	0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x63, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x70,
	0x2e, 0x63, 0x6f, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x43,
	0x6f, 0x69, 0x6e, 0x52, 0x05, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x6f,
	0x74, 0x61, 0x6c, 0x18, 0x14, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c,
	0x22, 0x6b, 0x0a, 0x17, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x75, 0x70, 0x6f, 0x6e,
	0x43, 0x6f, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x41,
	0x70, 0x70, 0x49, 0x44, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x41, 0x70, 0x70, 0x49,
	0x44, 0x12, 0x1a, 0x0a, 0x08, 0x43, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x49, 0x44, 0x18, 0x14, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x43, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x49, 0x44, 0x12, 0x1e, 0x0a,
	0x0a, 0x43, 0x6f, 0x69, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x49, 0x44, 0x18, 0x1e, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x43, 0x6f, 0x69, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x49, 0x44, 0x22, 0x5e, 0x0a,
	0x18, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x43, 0x6f, 0x69,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x42, 0x0a, 0x04, 0x49, 0x6e, 0x66,
	0x6f, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x69, 0x6e, 0x73, 0x70, 0x69, 0x72,
	0x65, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x63, 0x6f, 0x75, 0x70, 0x6f, 0x6e,
	0x2e, 0x61, 0x70, 0x70, 0x2e, 0x63, 0x6f, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x75,
	0x70, 0x6f, 0x6e, 0x43, 0x6f, 0x69, 0x6e, 0x52, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x55, 0x0a,
	0x17, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x43, 0x6f, 0x69,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x45, 0x6e, 0x74, 0x49,
	0x44, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x45, 0x6e, 0x74, 0x49, 0x44, 0x12, 0x14,
	0x0a, 0x05, 0x41, 0x70, 0x70, 0x49, 0x44, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x41,
	0x70, 0x70, 0x49, 0x44, 0x22, 0x5e, 0x0a, 0x18, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6f,
	0x75, 0x70, 0x6f, 0x6e, 0x43, 0x6f, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x42, 0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2e,
	0x2e, 0x69, 0x6e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79,
	0x2e, 0x63, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x63, 0x6f, 0x69, 0x6e,
	0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x43, 0x6f, 0x69, 0x6e, 0x52, 0x04,
	0x49, 0x6e, 0x66, 0x6f, 0x32, 0x97, 0x04, 0x0a, 0x07, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79,
	0x12, 0xa7, 0x01, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x43, 0x6f,
	0x69, 0x6e, 0x73, 0x12, 0x39, 0x2e, 0x69, 0x6e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x67, 0x61,
	0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x63, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x70,
	0x2e, 0x63, 0x6f, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x75, 0x70,
	0x6f, 0x6e, 0x43, 0x6f, 0x69, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3a,
	0x2e, 0x69, 0x6e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79,
	0x2e, 0x63, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x63, 0x6f, 0x69, 0x6e,
	0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x43, 0x6f, 0x69,
	0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1e, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x18, 0x3a, 0x01, 0x2a, 0x22, 0x13, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x65, 0x74, 0x2f, 0x63,
	0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x63, 0x6f, 0x69, 0x6e, 0x73, 0x12, 0xaf, 0x01, 0x0a, 0x10, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x43, 0x6f, 0x69, 0x6e, 0x12,
	0x3b, 0x2e, 0x69, 0x6e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61,
	0x79, 0x2e, 0x63, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x63, 0x6f, 0x69,
	0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x75, 0x70, 0x6f,
	0x6e, 0x43, 0x6f, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3c, 0x2e, 0x69,
	0x6e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x63,
	0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x63, 0x6f, 0x69, 0x6e, 0x2e, 0x76,
	0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x43, 0x6f,
	0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x20, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x1a, 0x3a, 0x01, 0x2a, 0x22, 0x15, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x2f, 0x63, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x63, 0x6f, 0x69, 0x6e, 0x12, 0xaf, 0x01, 0x0a,
	0x10, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x43, 0x6f, 0x69,
	0x6e, 0x12, 0x3b, 0x2e, 0x69, 0x6e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x67, 0x61, 0x74, 0x65,
	0x77, 0x61, 0x79, 0x2e, 0x63, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x63,
	0x6f, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6f, 0x75,
	0x70, 0x6f, 0x6e, 0x43, 0x6f, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3c,
	0x2e, 0x69, 0x6e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79,
	0x2e, 0x63, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x63, 0x6f, 0x69, 0x6e,
	0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6f, 0x75, 0x70, 0x6f, 0x6e,
	0x43, 0x6f, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x20, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x1a, 0x3a, 0x01, 0x2a, 0x22, 0x15, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x2f, 0x63, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x63, 0x6f, 0x69, 0x6e, 0x42, 0x46,
	0x5a, 0x44, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4e, 0x70, 0x6f,
	0x6f, 0x6c, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x2f, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x69, 0x6e, 0x73, 0x70, 0x69, 0x72, 0x65,
	0x2f, 0x67, 0x77, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x2f, 0x61, 0x70,
	0x70, 0x2f, 0x63, 0x6f, 0x69, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_npool_inspire_gw_v1_coupon_app_coin_coin_proto_rawDescOnce sync.Once
	file_npool_inspire_gw_v1_coupon_app_coin_coin_proto_rawDescData = file_npool_inspire_gw_v1_coupon_app_coin_coin_proto_rawDesc
)

func file_npool_inspire_gw_v1_coupon_app_coin_coin_proto_rawDescGZIP() []byte {
	file_npool_inspire_gw_v1_coupon_app_coin_coin_proto_rawDescOnce.Do(func() {
		file_npool_inspire_gw_v1_coupon_app_coin_coin_proto_rawDescData = protoimpl.X.CompressGZIP(file_npool_inspire_gw_v1_coupon_app_coin_coin_proto_rawDescData)
	})
	return file_npool_inspire_gw_v1_coupon_app_coin_coin_proto_rawDescData
}

var file_npool_inspire_gw_v1_coupon_app_coin_coin_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_npool_inspire_gw_v1_coupon_app_coin_coin_proto_goTypes = []interface{}{
	(*CouponCoin)(nil),               // 0: inspire.gateway.coupon.app.coin.v1.CouponCoin
	(*GetCouponCoinsRequest)(nil),    // 1: inspire.gateway.coupon.app.coin.v1.GetCouponCoinsRequest
	(*GetCouponCoinsResponse)(nil),   // 2: inspire.gateway.coupon.app.coin.v1.GetCouponCoinsResponse
	(*CreateCouponCoinRequest)(nil),  // 3: inspire.gateway.coupon.app.coin.v1.CreateCouponCoinRequest
	(*CreateCouponCoinResponse)(nil), // 4: inspire.gateway.coupon.app.coin.v1.CreateCouponCoinResponse
	(*DeleteCouponCoinRequest)(nil),  // 5: inspire.gateway.coupon.app.coin.v1.DeleteCouponCoinRequest
	(*DeleteCouponCoinResponse)(nil), // 6: inspire.gateway.coupon.app.coin.v1.DeleteCouponCoinResponse
}
var file_npool_inspire_gw_v1_coupon_app_coin_coin_proto_depIdxs = []int32{
	0, // 0: inspire.gateway.coupon.app.coin.v1.GetCouponCoinsResponse.Infos:type_name -> inspire.gateway.coupon.app.coin.v1.CouponCoin
	0, // 1: inspire.gateway.coupon.app.coin.v1.CreateCouponCoinResponse.Info:type_name -> inspire.gateway.coupon.app.coin.v1.CouponCoin
	0, // 2: inspire.gateway.coupon.app.coin.v1.DeleteCouponCoinResponse.Info:type_name -> inspire.gateway.coupon.app.coin.v1.CouponCoin
	1, // 3: inspire.gateway.coupon.app.coin.v1.Gateway.GetCouponCoins:input_type -> inspire.gateway.coupon.app.coin.v1.GetCouponCoinsRequest
	3, // 4: inspire.gateway.coupon.app.coin.v1.Gateway.CreateCouponCoin:input_type -> inspire.gateway.coupon.app.coin.v1.CreateCouponCoinRequest
	5, // 5: inspire.gateway.coupon.app.coin.v1.Gateway.DeleteCouponCoin:input_type -> inspire.gateway.coupon.app.coin.v1.DeleteCouponCoinRequest
	2, // 6: inspire.gateway.coupon.app.coin.v1.Gateway.GetCouponCoins:output_type -> inspire.gateway.coupon.app.coin.v1.GetCouponCoinsResponse
	4, // 7: inspire.gateway.coupon.app.coin.v1.Gateway.CreateCouponCoin:output_type -> inspire.gateway.coupon.app.coin.v1.CreateCouponCoinResponse
	6, // 8: inspire.gateway.coupon.app.coin.v1.Gateway.DeleteCouponCoin:output_type -> inspire.gateway.coupon.app.coin.v1.DeleteCouponCoinResponse
	6, // [6:9] is the sub-list for method output_type
	3, // [3:6] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_npool_inspire_gw_v1_coupon_app_coin_coin_proto_init() }
func file_npool_inspire_gw_v1_coupon_app_coin_coin_proto_init() {
	if File_npool_inspire_gw_v1_coupon_app_coin_coin_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_npool_inspire_gw_v1_coupon_app_coin_coin_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CouponCoin); i {
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
		file_npool_inspire_gw_v1_coupon_app_coin_coin_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCouponCoinsRequest); i {
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
		file_npool_inspire_gw_v1_coupon_app_coin_coin_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCouponCoinsResponse); i {
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
		file_npool_inspire_gw_v1_coupon_app_coin_coin_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateCouponCoinRequest); i {
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
		file_npool_inspire_gw_v1_coupon_app_coin_coin_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateCouponCoinResponse); i {
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
		file_npool_inspire_gw_v1_coupon_app_coin_coin_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteCouponCoinRequest); i {
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
		file_npool_inspire_gw_v1_coupon_app_coin_coin_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteCouponCoinResponse); i {
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
			RawDescriptor: file_npool_inspire_gw_v1_coupon_app_coin_coin_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_npool_inspire_gw_v1_coupon_app_coin_coin_proto_goTypes,
		DependencyIndexes: file_npool_inspire_gw_v1_coupon_app_coin_coin_proto_depIdxs,
		MessageInfos:      file_npool_inspire_gw_v1_coupon_app_coin_coin_proto_msgTypes,
	}.Build()
	File_npool_inspire_gw_v1_coupon_app_coin_coin_proto = out.File
	file_npool_inspire_gw_v1_coupon_app_coin_coin_proto_rawDesc = nil
	file_npool_inspire_gw_v1_coupon_app_coin_coin_proto_goTypes = nil
	file_npool_inspire_gw_v1_coupon_app_coin_coin_proto_depIdxs = nil
}
