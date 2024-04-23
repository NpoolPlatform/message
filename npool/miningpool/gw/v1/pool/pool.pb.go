// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.18.1
// source: npool/miningpool/gw/v1/pool/pool.proto

package pool

import (
	v1 "github.com/NpoolPlatform/message/npool/basetypes/miningpool/v1"
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

type Pool struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID             uint32            `protobuf:"varint,9,opt,name=ID,proto3" json:"ID,omitempty"`
	EntID          string            `protobuf:"bytes,10,opt,name=EntID,proto3" json:"EntID,omitempty"`
	MiningpoolType v1.MiningpoolType `protobuf:"varint,20,opt,name=MiningpoolType,proto3,enum=basetypes.miningpool.v1.MiningpoolType" json:"MiningpoolType,omitempty"`
	Name           string            `protobuf:"bytes,30,opt,name=Name,proto3" json:"Name,omitempty"`
	Site           string            `protobuf:"bytes,40,opt,name=Site,proto3" json:"Site,omitempty"`
	Description    string            `protobuf:"bytes,50,opt,name=Description,proto3" json:"Description,omitempty"`
	Coins          []*Coin           `protobuf:"bytes,60,rep,name=Coins,proto3" json:"Coins,omitempty"`
	FractionRules  []*FractionRule   `protobuf:"bytes,70,rep,name=FractionRules,proto3" json:"FractionRules,omitempty"`
}

func (x *Pool) Reset() {
	*x = Pool{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_miningpool_gw_v1_pool_pool_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Pool) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Pool) ProtoMessage() {}

func (x *Pool) ProtoReflect() protoreflect.Message {
	mi := &file_npool_miningpool_gw_v1_pool_pool_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Pool.ProtoReflect.Descriptor instead.
func (*Pool) Descriptor() ([]byte, []int) {
	return file_npool_miningpool_gw_v1_pool_pool_proto_rawDescGZIP(), []int{0}
}

func (x *Pool) GetID() uint32 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *Pool) GetEntID() string {
	if x != nil {
		return x.EntID
	}
	return ""
}

func (x *Pool) GetMiningpoolType() v1.MiningpoolType {
	if x != nil {
		return x.MiningpoolType
	}
	return v1.MiningpoolType(0)
}

func (x *Pool) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Pool) GetSite() string {
	if x != nil {
		return x.Site
	}
	return ""
}

func (x *Pool) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Pool) GetCoins() []*Coin {
	if x != nil {
		return x.Coins
	}
	return nil
}

func (x *Pool) GetFractionRules() []*FractionRule {
	if x != nil {
		return x.FractionRules
	}
	return nil
}

type Coin struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CoinType         v1.CoinType      `protobuf:"varint,10,opt,name=CoinType,proto3,enum=basetypes.miningpool.v1.CoinType" json:"CoinType,omitempty"`
	RevenueTypes     []v1.RevenueType `protobuf:"varint,20,rep,packed,name=RevenueTypes,proto3,enum=basetypes.miningpool.v1.RevenueType" json:"RevenueTypes,omitempty"`
	FeeRate          string           `protobuf:"bytes,30,opt,name=FeeRate,proto3" json:"FeeRate,omitempty"`
	FixedRevenueAble bool             `protobuf:"varint,40,opt,name=FixedRevenueAble,proto3" json:"FixedRevenueAble,omitempty"`
	Threshold        string           `protobuf:"bytes,50,opt,name=Threshold,proto3" json:"Threshold,omitempty"`
	Remark           string           `protobuf:"bytes,60,opt,name=Remark,proto3" json:"Remark,omitempty"`
}

func (x *Coin) Reset() {
	*x = Coin{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_miningpool_gw_v1_pool_pool_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Coin) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Coin) ProtoMessage() {}

func (x *Coin) ProtoReflect() protoreflect.Message {
	mi := &file_npool_miningpool_gw_v1_pool_pool_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Coin.ProtoReflect.Descriptor instead.
func (*Coin) Descriptor() ([]byte, []int) {
	return file_npool_miningpool_gw_v1_pool_pool_proto_rawDescGZIP(), []int{1}
}

func (x *Coin) GetCoinType() v1.CoinType {
	if x != nil {
		return x.CoinType
	}
	return v1.CoinType(0)
}

func (x *Coin) GetRevenueTypes() []v1.RevenueType {
	if x != nil {
		return x.RevenueTypes
	}
	return nil
}

func (x *Coin) GetFeeRate() string {
	if x != nil {
		return x.FeeRate
	}
	return ""
}

func (x *Coin) GetFixedRevenueAble() bool {
	if x != nil {
		return x.FixedRevenueAble
	}
	return false
}

func (x *Coin) GetThreshold() string {
	if x != nil {
		return x.Threshold
	}
	return ""
}

func (x *Coin) GetRemark() string {
	if x != nil {
		return x.Remark
	}
	return ""
}

type FractionRule struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CoinType         v1.CoinType `protobuf:"varint,10,opt,name=CoinType,proto3,enum=basetypes.miningpool.v1.CoinType" json:"CoinType,omitempty"`
	WithdrawInterval uint32      `protobuf:"varint,20,opt,name=WithdrawInterval,proto3" json:"WithdrawInterval,omitempty"`
	MinAmount        float32     `protobuf:"fixed32,30,opt,name=MinAmount,proto3" json:"MinAmount,omitempty"`
	WithdrawRate     float32     `protobuf:"fixed32,40,opt,name=WithdrawRate,proto3" json:"WithdrawRate,omitempty"`
}

func (x *FractionRule) Reset() {
	*x = FractionRule{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_miningpool_gw_v1_pool_pool_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FractionRule) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FractionRule) ProtoMessage() {}

func (x *FractionRule) ProtoReflect() protoreflect.Message {
	mi := &file_npool_miningpool_gw_v1_pool_pool_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FractionRule.ProtoReflect.Descriptor instead.
func (*FractionRule) Descriptor() ([]byte, []int) {
	return file_npool_miningpool_gw_v1_pool_pool_proto_rawDescGZIP(), []int{2}
}

func (x *FractionRule) GetCoinType() v1.CoinType {
	if x != nil {
		return x.CoinType
	}
	return v1.CoinType(0)
}

func (x *FractionRule) GetWithdrawInterval() uint32 {
	if x != nil {
		return x.WithdrawInterval
	}
	return 0
}

func (x *FractionRule) GetMinAmount() float32 {
	if x != nil {
		return x.MinAmount
	}
	return 0
}

func (x *FractionRule) GetWithdrawRate() float32 {
	if x != nil {
		return x.WithdrawRate
	}
	return 0
}

type AdminGetPoolsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Offset int32 `protobuf:"varint,10,opt,name=Offset,proto3" json:"Offset,omitempty"`
	Limit  int32 `protobuf:"varint,20,opt,name=Limit,proto3" json:"Limit,omitempty"`
}

func (x *AdminGetPoolsRequest) Reset() {
	*x = AdminGetPoolsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_miningpool_gw_v1_pool_pool_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdminGetPoolsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdminGetPoolsRequest) ProtoMessage() {}

func (x *AdminGetPoolsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_miningpool_gw_v1_pool_pool_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdminGetPoolsRequest.ProtoReflect.Descriptor instead.
func (*AdminGetPoolsRequest) Descriptor() ([]byte, []int) {
	return file_npool_miningpool_gw_v1_pool_pool_proto_rawDescGZIP(), []int{3}
}

func (x *AdminGetPoolsRequest) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *AdminGetPoolsRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type AdminGetPoolsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Infos []*Pool `protobuf:"bytes,10,rep,name=Infos,proto3" json:"Infos,omitempty"`
	Total uint32  `protobuf:"varint,20,opt,name=Total,proto3" json:"Total,omitempty"`
}

func (x *AdminGetPoolsResponse) Reset() {
	*x = AdminGetPoolsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_miningpool_gw_v1_pool_pool_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdminGetPoolsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdminGetPoolsResponse) ProtoMessage() {}

func (x *AdminGetPoolsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_miningpool_gw_v1_pool_pool_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdminGetPoolsResponse.ProtoReflect.Descriptor instead.
func (*AdminGetPoolsResponse) Descriptor() ([]byte, []int) {
	return file_npool_miningpool_gw_v1_pool_pool_proto_rawDescGZIP(), []int{4}
}

func (x *AdminGetPoolsResponse) GetInfos() []*Pool {
	if x != nil {
		return x.Infos
	}
	return nil
}

func (x *AdminGetPoolsResponse) GetTotal() uint32 {
	if x != nil {
		return x.Total
	}
	return 0
}

type AdminGetPoolRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EntID string `protobuf:"bytes,10,opt,name=EntID,proto3" json:"EntID,omitempty"`
}

func (x *AdminGetPoolRequest) Reset() {
	*x = AdminGetPoolRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_miningpool_gw_v1_pool_pool_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdminGetPoolRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdminGetPoolRequest) ProtoMessage() {}

func (x *AdminGetPoolRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_miningpool_gw_v1_pool_pool_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdminGetPoolRequest.ProtoReflect.Descriptor instead.
func (*AdminGetPoolRequest) Descriptor() ([]byte, []int) {
	return file_npool_miningpool_gw_v1_pool_pool_proto_rawDescGZIP(), []int{5}
}

func (x *AdminGetPoolRequest) GetEntID() string {
	if x != nil {
		return x.EntID
	}
	return ""
}

type AdminGetPoolResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *Pool `protobuf:"bytes,10,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (x *AdminGetPoolResponse) Reset() {
	*x = AdminGetPoolResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_miningpool_gw_v1_pool_pool_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdminGetPoolResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdminGetPoolResponse) ProtoMessage() {}

func (x *AdminGetPoolResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_miningpool_gw_v1_pool_pool_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdminGetPoolResponse.ProtoReflect.Descriptor instead.
func (*AdminGetPoolResponse) Descriptor() ([]byte, []int) {
	return file_npool_miningpool_gw_v1_pool_pool_proto_rawDescGZIP(), []int{6}
}

func (x *AdminGetPoolResponse) GetInfo() *Pool {
	if x != nil {
		return x.Info
	}
	return nil
}

var File_npool_miningpool_gw_v1_pool_pool_proto protoreflect.FileDescriptor

var file_npool_miningpool_gw_v1_pool_pool_proto_rawDesc = []byte{
	0x0a, 0x26, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x6d, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x70, 0x6f,
	0x6f, 0x6c, 0x2f, 0x67, 0x77, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x70, 0x6f,
	0x6f, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1a, 0x6d, 0x69, 0x6e, 0x69, 0x6e, 0x67,
	0x70, 0x6f, 0x6f, 0x6c, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x70, 0x6f, 0x6f,
	0x6c, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x29, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x2f, 0x6d, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x76,
	0x31, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xcf, 0x02,
	0x0a, 0x04, 0x50, 0x6f, 0x6f, 0x6c, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x02, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x45, 0x6e, 0x74, 0x49, 0x44, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x45, 0x6e, 0x74, 0x49, 0x44, 0x12, 0x4f, 0x0a, 0x0e,
	0x4d, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x70, 0x6f, 0x6f, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x18, 0x14,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x27, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x2e, 0x6d, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x70, 0x6f, 0x6f, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x4d,
	0x69, 0x6e, 0x69, 0x6e, 0x67, 0x70, 0x6f, 0x6f, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0e, 0x4d,
	0x69, 0x6e, 0x69, 0x6e, 0x67, 0x70, 0x6f, 0x6f, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x53, 0x69, 0x74, 0x65, 0x18, 0x28, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x53, 0x69, 0x74, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x32, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x44, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x36, 0x0a, 0x05, 0x43, 0x6f, 0x69, 0x6e, 0x73,
	0x18, 0x3c, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x6d, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x70,
	0x6f, 0x6f, 0x6c, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x70, 0x6f, 0x6f, 0x6c,
	0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x69, 0x6e, 0x52, 0x05, 0x43, 0x6f, 0x69, 0x6e, 0x73, 0x12,
	0x4e, 0x0a, 0x0d, 0x46, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x75, 0x6c, 0x65, 0x73,
	0x18, 0x46, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x6d, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x70,
	0x6f, 0x6f, 0x6c, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x70, 0x6f, 0x6f, 0x6c,
	0x2e, 0x76, 0x31, 0x2e, 0x46, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x75, 0x6c, 0x65,
	0x52, 0x0d, 0x46, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x75, 0x6c, 0x65, 0x73, 0x22,
	0x8b, 0x02, 0x0a, 0x04, 0x43, 0x6f, 0x69, 0x6e, 0x12, 0x3d, 0x0a, 0x08, 0x43, 0x6f, 0x69, 0x6e,
	0x54, 0x79, 0x70, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x21, 0x2e, 0x62, 0x61, 0x73,
	0x65, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x6d, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x70, 0x6f, 0x6f,
	0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x69, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x08, 0x43,
	0x6f, 0x69, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x48, 0x0a, 0x0c, 0x52, 0x65, 0x76, 0x65, 0x6e,
	0x75, 0x65, 0x54, 0x79, 0x70, 0x65, 0x73, 0x18, 0x14, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x24, 0x2e,
	0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x6d, 0x69, 0x6e, 0x69, 0x6e, 0x67,
	0x70, 0x6f, 0x6f, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x76, 0x65, 0x6e, 0x75, 0x65, 0x54,
	0x79, 0x70, 0x65, 0x52, 0x0c, 0x52, 0x65, 0x76, 0x65, 0x6e, 0x75, 0x65, 0x54, 0x79, 0x70, 0x65,
	0x73, 0x12, 0x18, 0x0a, 0x07, 0x46, 0x65, 0x65, 0x52, 0x61, 0x74, 0x65, 0x18, 0x1e, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x46, 0x65, 0x65, 0x52, 0x61, 0x74, 0x65, 0x12, 0x2a, 0x0a, 0x10, 0x46,
	0x69, 0x78, 0x65, 0x64, 0x52, 0x65, 0x76, 0x65, 0x6e, 0x75, 0x65, 0x41, 0x62, 0x6c, 0x65, 0x18,
	0x28, 0x20, 0x01, 0x28, 0x08, 0x52, 0x10, 0x46, 0x69, 0x78, 0x65, 0x64, 0x52, 0x65, 0x76, 0x65,
	0x6e, 0x75, 0x65, 0x41, 0x62, 0x6c, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x54, 0x68, 0x72, 0x65, 0x73,
	0x68, 0x6f, 0x6c, 0x64, 0x18, 0x32, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x54, 0x68, 0x72, 0x65,
	0x73, 0x68, 0x6f, 0x6c, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x52, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x18,
	0x3c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x52, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x22, 0xbb, 0x01,
	0x0a, 0x0c, 0x46, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x75, 0x6c, 0x65, 0x12, 0x3d,
	0x0a, 0x08, 0x43, 0x6f, 0x69, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x21, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x6d, 0x69, 0x6e,
	0x69, 0x6e, 0x67, 0x70, 0x6f, 0x6f, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x69, 0x6e, 0x54,
	0x79, 0x70, 0x65, 0x52, 0x08, 0x43, 0x6f, 0x69, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x2a, 0x0a,
	0x10, 0x57, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61,
	0x6c, 0x18, 0x14, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x10, 0x57, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61,
	0x77, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x12, 0x1c, 0x0a, 0x09, 0x4d, 0x69, 0x6e,
	0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x02, 0x52, 0x09, 0x4d, 0x69,
	0x6e, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x22, 0x0a, 0x0c, 0x57, 0x69, 0x74, 0x68, 0x64,
	0x72, 0x61, 0x77, 0x52, 0x61, 0x74, 0x65, 0x18, 0x28, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0c, 0x57,
	0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x52, 0x61, 0x74, 0x65, 0x22, 0x44, 0x0a, 0x14, 0x41,
	0x64, 0x6d, 0x69, 0x6e, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x6f, 0x6c, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x06, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x4c,
	0x69, 0x6d, 0x69, 0x74, 0x18, 0x14, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x4c, 0x69, 0x6d, 0x69,
	0x74, 0x22, 0x65, 0x0a, 0x15, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x6f,
	0x6c, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x36, 0x0a, 0x05, 0x49, 0x6e,
	0x66, 0x6f, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x6d, 0x69, 0x6e, 0x69,
	0x6e, 0x67, 0x70, 0x6f, 0x6f, 0x6c, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x70,
	0x6f, 0x6f, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6f, 0x6f, 0x6c, 0x52, 0x05, 0x49, 0x6e, 0x66,
	0x6f, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x14, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x22, 0x2b, 0x0a, 0x13, 0x41, 0x64, 0x6d, 0x69,
	0x6e, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x6f, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x14, 0x0a, 0x05, 0x45, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x45, 0x6e, 0x74, 0x49, 0x44, 0x22, 0x4c, 0x0a, 0x14, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x47, 0x65,
	0x74, 0x50, 0x6f, 0x6f, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x34, 0x0a,
	0x04, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x6d, 0x69,
	0x6e, 0x69, 0x6e, 0x67, 0x70, 0x6f, 0x6f, 0x6c, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79,
	0x2e, 0x70, 0x6f, 0x6f, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6f, 0x6f, 0x6c, 0x52, 0x04, 0x49,
	0x6e, 0x66, 0x6f, 0x32, 0xb8, 0x02, 0x0a, 0x07, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x12,
	0x99, 0x01, 0x0a, 0x12, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x47, 0x65,
	0x74, 0x50, 0x6f, 0x6f, 0x6c, 0x73, 0x12, 0x30, 0x2e, 0x6d, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x70,
	0x6f, 0x6f, 0x6c, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x70, 0x6f, 0x6f, 0x6c,
	0x2e, 0x76, 0x31, 0x2e, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x6f, 0x6c,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x31, 0x2e, 0x6d, 0x69, 0x6e, 0x69, 0x6e,
	0x67, 0x70, 0x6f, 0x6f, 0x6c, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x70, 0x6f,
	0x6f, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x47, 0x65, 0x74, 0x50, 0x6f,
	0x6f, 0x6c, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1e, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x18, 0x3a, 0x01, 0x2a, 0x22, 0x13, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x64, 0x6d, 0x69,
	0x6e, 0x2f, 0x67, 0x65, 0x74, 0x2f, 0x70, 0x6f, 0x6f, 0x6c, 0x73, 0x12, 0x90, 0x01, 0x0a, 0x0c,
	0x41, 0x64, 0x6d, 0x69, 0x6e, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x6f, 0x6c, 0x12, 0x2f, 0x2e, 0x6d,
	0x69, 0x6e, 0x69, 0x6e, 0x67, 0x70, 0x6f, 0x6f, 0x6c, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61,
	0x79, 0x2e, 0x70, 0x6f, 0x6f, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x47,
	0x65, 0x74, 0x50, 0x6f, 0x6f, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x30, 0x2e,
	0x6d, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x70, 0x6f, 0x6f, 0x6c, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77,
	0x61, 0x79, 0x2e, 0x70, 0x6f, 0x6f, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x64, 0x6d, 0x69, 0x6e,
	0x47, 0x65, 0x74, 0x50, 0x6f, 0x6f, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x1d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x17, 0x3a, 0x01, 0x2a, 0x22, 0x12, 0x2f, 0x76, 0x31, 0x2f,
	0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x67, 0x65, 0x74, 0x2f, 0x70, 0x6f, 0x6f, 0x6c, 0x42, 0x3e,
	0x5a, 0x3c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4e, 0x70, 0x6f,
	0x6f, 0x6c, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x2f, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x6d, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x70,
	0x6f, 0x6f, 0x6c, 0x2f, 0x67, 0x77, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x6f, 0x6f, 0x6c, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_npool_miningpool_gw_v1_pool_pool_proto_rawDescOnce sync.Once
	file_npool_miningpool_gw_v1_pool_pool_proto_rawDescData = file_npool_miningpool_gw_v1_pool_pool_proto_rawDesc
)

func file_npool_miningpool_gw_v1_pool_pool_proto_rawDescGZIP() []byte {
	file_npool_miningpool_gw_v1_pool_pool_proto_rawDescOnce.Do(func() {
		file_npool_miningpool_gw_v1_pool_pool_proto_rawDescData = protoimpl.X.CompressGZIP(file_npool_miningpool_gw_v1_pool_pool_proto_rawDescData)
	})
	return file_npool_miningpool_gw_v1_pool_pool_proto_rawDescData
}

var file_npool_miningpool_gw_v1_pool_pool_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_npool_miningpool_gw_v1_pool_pool_proto_goTypes = []interface{}{
	(*Pool)(nil),                  // 0: miningpool.gateway.pool.v1.Pool
	(*Coin)(nil),                  // 1: miningpool.gateway.pool.v1.Coin
	(*FractionRule)(nil),          // 2: miningpool.gateway.pool.v1.FractionRule
	(*AdminGetPoolsRequest)(nil),  // 3: miningpool.gateway.pool.v1.AdminGetPoolsRequest
	(*AdminGetPoolsResponse)(nil), // 4: miningpool.gateway.pool.v1.AdminGetPoolsResponse
	(*AdminGetPoolRequest)(nil),   // 5: miningpool.gateway.pool.v1.AdminGetPoolRequest
	(*AdminGetPoolResponse)(nil),  // 6: miningpool.gateway.pool.v1.AdminGetPoolResponse
	(v1.MiningpoolType)(0),        // 7: basetypes.miningpool.v1.MiningpoolType
	(v1.CoinType)(0),              // 8: basetypes.miningpool.v1.CoinType
	(v1.RevenueType)(0),           // 9: basetypes.miningpool.v1.RevenueType
}
var file_npool_miningpool_gw_v1_pool_pool_proto_depIdxs = []int32{
	7,  // 0: miningpool.gateway.pool.v1.Pool.MiningpoolType:type_name -> basetypes.miningpool.v1.MiningpoolType
	1,  // 1: miningpool.gateway.pool.v1.Pool.Coins:type_name -> miningpool.gateway.pool.v1.Coin
	2,  // 2: miningpool.gateway.pool.v1.Pool.FractionRules:type_name -> miningpool.gateway.pool.v1.FractionRule
	8,  // 3: miningpool.gateway.pool.v1.Coin.CoinType:type_name -> basetypes.miningpool.v1.CoinType
	9,  // 4: miningpool.gateway.pool.v1.Coin.RevenueTypes:type_name -> basetypes.miningpool.v1.RevenueType
	8,  // 5: miningpool.gateway.pool.v1.FractionRule.CoinType:type_name -> basetypes.miningpool.v1.CoinType
	0,  // 6: miningpool.gateway.pool.v1.AdminGetPoolsResponse.Infos:type_name -> miningpool.gateway.pool.v1.Pool
	0,  // 7: miningpool.gateway.pool.v1.AdminGetPoolResponse.Info:type_name -> miningpool.gateway.pool.v1.Pool
	3,  // 8: miningpool.gateway.pool.v1.Gateway.AdminAdminGetPools:input_type -> miningpool.gateway.pool.v1.AdminGetPoolsRequest
	5,  // 9: miningpool.gateway.pool.v1.Gateway.AdminGetPool:input_type -> miningpool.gateway.pool.v1.AdminGetPoolRequest
	4,  // 10: miningpool.gateway.pool.v1.Gateway.AdminAdminGetPools:output_type -> miningpool.gateway.pool.v1.AdminGetPoolsResponse
	6,  // 11: miningpool.gateway.pool.v1.Gateway.AdminGetPool:output_type -> miningpool.gateway.pool.v1.AdminGetPoolResponse
	10, // [10:12] is the sub-list for method output_type
	8,  // [8:10] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_npool_miningpool_gw_v1_pool_pool_proto_init() }
func file_npool_miningpool_gw_v1_pool_pool_proto_init() {
	if File_npool_miningpool_gw_v1_pool_pool_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_npool_miningpool_gw_v1_pool_pool_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Pool); i {
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
		file_npool_miningpool_gw_v1_pool_pool_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Coin); i {
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
		file_npool_miningpool_gw_v1_pool_pool_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FractionRule); i {
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
		file_npool_miningpool_gw_v1_pool_pool_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdminGetPoolsRequest); i {
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
		file_npool_miningpool_gw_v1_pool_pool_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdminGetPoolsResponse); i {
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
		file_npool_miningpool_gw_v1_pool_pool_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdminGetPoolRequest); i {
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
		file_npool_miningpool_gw_v1_pool_pool_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdminGetPoolResponse); i {
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
			RawDescriptor: file_npool_miningpool_gw_v1_pool_pool_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_npool_miningpool_gw_v1_pool_pool_proto_goTypes,
		DependencyIndexes: file_npool_miningpool_gw_v1_pool_pool_proto_depIdxs,
		MessageInfos:      file_npool_miningpool_gw_v1_pool_pool_proto_msgTypes,
	}.Build()
	File_npool_miningpool_gw_v1_pool_pool_proto = out.File
	file_npool_miningpool_gw_v1_pool_pool_proto_rawDesc = nil
	file_npool_miningpool_gw_v1_pool_pool_proto_goTypes = nil
	file_npool_miningpool_gw_v1_pool_pool_proto_depIdxs = nil
}