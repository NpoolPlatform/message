// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v3.18.1
// source: npool/good/mw/v1/good/stock/stock.proto

package stock

import (
	v1 "github.com/NpoolPlatform/message/npool/basetypes/good/v1"
	v11 "github.com/NpoolPlatform/message/npool/basetypes/v1"
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

type MiningGoodStockReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EntID          *string       `protobuf:"bytes,10,opt,name=EntID,proto3,oneof" json:"EntID,omitempty"`
	PoolRootUserID *string       `protobuf:"bytes,20,opt,name=PoolRootUserID,proto3,oneof" json:"PoolRootUserID,omitempty"`
	MiningPoolID   *string       `protobuf:"bytes,30,opt,name=MiningPoolID,proto3,oneof" json:"MiningPoolID,omitempty"`
	PoolGoodUserID *string       `protobuf:"bytes,40,opt,name=PoolGoodUserID,proto3,oneof" json:"PoolGoodUserID,omitempty"`
	Total          *string       `protobuf:"bytes,50,opt,name=Total,proto3,oneof" json:"Total,omitempty"`
	State          *v1.GoodState `protobuf:"varint,60,opt,name=State,proto3,enum=basetypes.good.v1.GoodState,oneof" json:"State,omitempty"`
}

func (x *MiningGoodStockReq) Reset() {
	*x = MiningGoodStockReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_good_mw_v1_good_stock_stock_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MiningGoodStockReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MiningGoodStockReq) ProtoMessage() {}

func (x *MiningGoodStockReq) ProtoReflect() protoreflect.Message {
	mi := &file_npool_good_mw_v1_good_stock_stock_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MiningGoodStockReq.ProtoReflect.Descriptor instead.
func (*MiningGoodStockReq) Descriptor() ([]byte, []int) {
	return file_npool_good_mw_v1_good_stock_stock_proto_rawDescGZIP(), []int{0}
}

func (x *MiningGoodStockReq) GetEntID() string {
	if x != nil && x.EntID != nil {
		return *x.EntID
	}
	return ""
}

func (x *MiningGoodStockReq) GetPoolRootUserID() string {
	if x != nil && x.PoolRootUserID != nil {
		return *x.PoolRootUserID
	}
	return ""
}

func (x *MiningGoodStockReq) GetMiningPoolID() string {
	if x != nil && x.MiningPoolID != nil {
		return *x.MiningPoolID
	}
	return ""
}

func (x *MiningGoodStockReq) GetPoolGoodUserID() string {
	if x != nil && x.PoolGoodUserID != nil {
		return *x.PoolGoodUserID
	}
	return ""
}

func (x *MiningGoodStockReq) GetTotal() string {
	if x != nil && x.Total != nil {
		return *x.Total
	}
	return ""
}

func (x *MiningGoodStockReq) GetState() v1.GoodState {
	if x != nil && x.State != nil {
		return *x.State
	}
	return v1.GoodState(0)
}

type MiningGoodStock struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: sql:"id"
	ID uint32 `protobuf:"varint,10,opt,name=ID,proto3" json:"ID,omitempty" sql:"id"`
	// @inject_tag: sql:"ent_id"
	EntID string `protobuf:"bytes,20,opt,name=EntID,proto3" json:"EntID,omitempty" sql:"ent_id"`
	// @inject_tag: sql:"good_id"
	GoodID string `protobuf:"bytes,21,opt,name=GoodID,proto3" json:"GoodID,omitempty" sql:"good_id"`
	// @inject_tag: sql:"good_stock_id"
	GoodStockID string `protobuf:"bytes,30,opt,name=GoodStockID,proto3" json:"GoodStockID,omitempty" sql:"good_stock_id"`
	// @inject_tag: sql:"mining_pool_id"
	MiningPoolID string `protobuf:"bytes,40,opt,name=MiningPoolID,proto3" json:"MiningPoolID,omitempty" sql:"mining_pool_id"`
	// @inject_tag: sql:"pool_good_user_id"
	PoolGoodUserID string `protobuf:"bytes,50,opt,name=PoolGoodUserID,proto3" json:"PoolGoodUserID,omitempty" sql:"pool_good_user_id"`
	// @inject_tag: sql:"pool_root_user_id"
	PoolRootUserID string `protobuf:"bytes,51,opt,name=PoolRootUserID,proto3" json:"PoolRootUserID,omitempty" sql:"pool_root_user_id"`
	// @inject_tag: sql:"total"
	Total string `protobuf:"bytes,60,opt,name=Total,proto3" json:"Total,omitempty" sql:"total"`
	// @inject_tag: sql:"spot_quantity"
	SpotQuantity string `protobuf:"bytes,70,opt,name=SpotQuantity,proto3" json:"SpotQuantity,omitempty" sql:"spot_quantity"`
	// @inject_tag: sql:"locked"
	Locked string `protobuf:"bytes,80,opt,name=Locked,proto3" json:"Locked,omitempty" sql:"locked"`
	// @inject_tag: sql:"wait_start"
	WaitStart string `protobuf:"bytes,90,opt,name=WaitStart,proto3" json:"WaitStart,omitempty" sql:"wait_start"`
	// @inject_tag: sql:"in_service"
	InService string `protobuf:"bytes,100,opt,name=InService,proto3" json:"InService,omitempty" sql:"in_service"`
	// @inject_tag: sql:"sold"
	Sold string `protobuf:"bytes,110,opt,name=Sold,proto3" json:"Sold,omitempty" sql:"sold"`
	// @inject_tag: sql:"state"
	State v1.GoodState `protobuf:"varint,120,opt,name=State,proto3,enum=basetypes.good.v1.GoodState" json:"State,omitempty" sql:"state"`
}

func (x *MiningGoodStock) Reset() {
	*x = MiningGoodStock{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_good_mw_v1_good_stock_stock_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MiningGoodStock) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MiningGoodStock) ProtoMessage() {}

func (x *MiningGoodStock) ProtoReflect() protoreflect.Message {
	mi := &file_npool_good_mw_v1_good_stock_stock_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MiningGoodStock.ProtoReflect.Descriptor instead.
func (*MiningGoodStock) Descriptor() ([]byte, []int) {
	return file_npool_good_mw_v1_good_stock_stock_proto_rawDescGZIP(), []int{1}
}

func (x *MiningGoodStock) GetID() uint32 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *MiningGoodStock) GetEntID() string {
	if x != nil {
		return x.EntID
	}
	return ""
}

func (x *MiningGoodStock) GetGoodID() string {
	if x != nil {
		return x.GoodID
	}
	return ""
}

func (x *MiningGoodStock) GetGoodStockID() string {
	if x != nil {
		return x.GoodStockID
	}
	return ""
}

func (x *MiningGoodStock) GetMiningPoolID() string {
	if x != nil {
		return x.MiningPoolID
	}
	return ""
}

func (x *MiningGoodStock) GetPoolGoodUserID() string {
	if x != nil {
		return x.PoolGoodUserID
	}
	return ""
}

func (x *MiningGoodStock) GetPoolRootUserID() string {
	if x != nil {
		return x.PoolRootUserID
	}
	return ""
}

func (x *MiningGoodStock) GetTotal() string {
	if x != nil {
		return x.Total
	}
	return ""
}

func (x *MiningGoodStock) GetSpotQuantity() string {
	if x != nil {
		return x.SpotQuantity
	}
	return ""
}

func (x *MiningGoodStock) GetLocked() string {
	if x != nil {
		return x.Locked
	}
	return ""
}

func (x *MiningGoodStock) GetWaitStart() string {
	if x != nil {
		return x.WaitStart
	}
	return ""
}

func (x *MiningGoodStock) GetInService() string {
	if x != nil {
		return x.InService
	}
	return ""
}

func (x *MiningGoodStock) GetSold() string {
	if x != nil {
		return x.Sold
	}
	return ""
}

func (x *MiningGoodStock) GetState() v1.GoodState {
	if x != nil {
		return x.State
	}
	return v1.GoodState(0)
}

type MiningGoodStockInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: sql:"ent_id"
	EntID string `protobuf:"bytes,10,opt,name=EntID,proto3" json:"EntID,omitempty" sql:"ent_id"`
	// @inject_tag: sql:"good_stock_id"
	GoodStockID string `protobuf:"bytes,20,opt,name=GoodStockID,proto3" json:"GoodStockID,omitempty" sql:"good_stock_id"`
	// @inject_tag: sql:"mining_pool_id"
	MiningPoolID string `protobuf:"bytes,30,opt,name=MiningPoolID,proto3" json:"MiningPoolID,omitempty" sql:"mining_pool_id"`
	// @inject_tag: sql:"pool_good_user_id"
	PoolGoodUserID string `protobuf:"bytes,40,opt,name=PoolGoodUserID,proto3" json:"PoolGoodUserID,omitempty" sql:"pool_good_user_id"`
	// @inject_tag: sql:"total"
	Total string `protobuf:"bytes,50,opt,name=Total,proto3" json:"Total,omitempty" sql:"total"`
	// @inject_tag: sql:"spot_quantity"
	SpotQuantity string `protobuf:"bytes,60,opt,name=SpotQuantity,proto3" json:"SpotQuantity,omitempty" sql:"spot_quantity"`
}

func (x *MiningGoodStockInfo) Reset() {
	*x = MiningGoodStockInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_good_mw_v1_good_stock_stock_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MiningGoodStockInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MiningGoodStockInfo) ProtoMessage() {}

func (x *MiningGoodStockInfo) ProtoReflect() protoreflect.Message {
	mi := &file_npool_good_mw_v1_good_stock_stock_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MiningGoodStockInfo.ProtoReflect.Descriptor instead.
func (*MiningGoodStockInfo) Descriptor() ([]byte, []int) {
	return file_npool_good_mw_v1_good_stock_stock_proto_rawDescGZIP(), []int{2}
}

func (x *MiningGoodStockInfo) GetEntID() string {
	if x != nil {
		return x.EntID
	}
	return ""
}

func (x *MiningGoodStockInfo) GetGoodStockID() string {
	if x != nil {
		return x.GoodStockID
	}
	return ""
}

func (x *MiningGoodStockInfo) GetMiningPoolID() string {
	if x != nil {
		return x.MiningPoolID
	}
	return ""
}

func (x *MiningGoodStockInfo) GetPoolGoodUserID() string {
	if x != nil {
		return x.PoolGoodUserID
	}
	return ""
}

func (x *MiningGoodStockInfo) GetTotal() string {
	if x != nil {
		return x.Total
	}
	return ""
}

func (x *MiningGoodStockInfo) GetSpotQuantity() string {
	if x != nil {
		return x.SpotQuantity
	}
	return ""
}

type Conds struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID             *v11.Uint32Val `protobuf:"bytes,10,opt,name=ID,proto3,oneof" json:"ID,omitempty"`
	EntID          *v11.StringVal `protobuf:"bytes,20,opt,name=EntID,proto3,oneof" json:"EntID,omitempty"`
	GoodStockID    *v11.StringVal `protobuf:"bytes,30,opt,name=GoodStockID,proto3,oneof" json:"GoodStockID,omitempty"`
	MiningPoolID   *v11.StringVal `protobuf:"bytes,40,opt,name=MiningPoolID,proto3,oneof" json:"MiningPoolID,omitempty"`
	PoolGoodUserID *v11.StringVal `protobuf:"bytes,50,opt,name=PoolGoodUserID,proto3,oneof" json:"PoolGoodUserID,omitempty"`
	PoolRootUserID *v11.StringVal `protobuf:"bytes,60,opt,name=PoolRootUserID,proto3,oneof" json:"PoolRootUserID,omitempty"`
	State          *v11.StringVal `protobuf:"bytes,70,opt,name=State,proto3,oneof" json:"State,omitempty"`
}

func (x *Conds) Reset() {
	*x = Conds{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_good_mw_v1_good_stock_stock_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Conds) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Conds) ProtoMessage() {}

func (x *Conds) ProtoReflect() protoreflect.Message {
	mi := &file_npool_good_mw_v1_good_stock_stock_proto_msgTypes[3]
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
	return file_npool_good_mw_v1_good_stock_stock_proto_rawDescGZIP(), []int{3}
}

func (x *Conds) GetID() *v11.Uint32Val {
	if x != nil {
		return x.ID
	}
	return nil
}

func (x *Conds) GetEntID() *v11.StringVal {
	if x != nil {
		return x.EntID
	}
	return nil
}

func (x *Conds) GetGoodStockID() *v11.StringVal {
	if x != nil {
		return x.GoodStockID
	}
	return nil
}

func (x *Conds) GetMiningPoolID() *v11.StringVal {
	if x != nil {
		return x.MiningPoolID
	}
	return nil
}

func (x *Conds) GetPoolGoodUserID() *v11.StringVal {
	if x != nil {
		return x.PoolGoodUserID
	}
	return nil
}

func (x *Conds) GetPoolRootUserID() *v11.StringVal {
	if x != nil {
		return x.PoolRootUserID
	}
	return nil
}

func (x *Conds) GetState() *v11.StringVal {
	if x != nil {
		return x.State
	}
	return nil
}

var File_npool_good_mw_v1_good_stock_stock_proto protoreflect.FileDescriptor

var file_npool_good_mw_v1_good_stock_stock_proto_rawDesc = []byte{
	0x0a, 0x27, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x67, 0x6f, 0x6f, 0x64, 0x2f, 0x6d, 0x77, 0x2f,
	0x76, 0x31, 0x2f, 0x67, 0x6f, 0x6f, 0x64, 0x2f, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x2f, 0x73, 0x74,
	0x6f, 0x63, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1e, 0x67, 0x6f, 0x6f, 0x64, 0x2e,
	0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x67, 0x6f, 0x6f, 0x64, 0x31,
	0x2e, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x2e, 0x76, 0x31, 0x1a, 0x23, 0x6e, 0x70, 0x6f, 0x6f, 0x6c,
	0x2f, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x64, 0x2f,
	0x76, 0x31, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e,
	0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f,
	0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xdb,
	0x02, 0x0a, 0x12, 0x4d, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x47, 0x6f, 0x6f, 0x64, 0x53, 0x74, 0x6f,
	0x63, 0x6b, 0x52, 0x65, 0x71, 0x12, 0x19, 0x0a, 0x05, 0x45, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x05, 0x45, 0x6e, 0x74, 0x49, 0x44, 0x88, 0x01, 0x01,
	0x12, 0x2b, 0x0a, 0x0e, 0x50, 0x6f, 0x6f, 0x6c, 0x52, 0x6f, 0x6f, 0x74, 0x55, 0x73, 0x65, 0x72,
	0x49, 0x44, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x0e, 0x50, 0x6f, 0x6f, 0x6c,
	0x52, 0x6f, 0x6f, 0x74, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x88, 0x01, 0x01, 0x12, 0x27, 0x0a,
	0x0c, 0x4d, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x50, 0x6f, 0x6f, 0x6c, 0x49, 0x44, 0x18, 0x1e, 0x20,
	0x01, 0x28, 0x09, 0x48, 0x02, 0x52, 0x0c, 0x4d, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x50, 0x6f, 0x6f,
	0x6c, 0x49, 0x44, 0x88, 0x01, 0x01, 0x12, 0x2b, 0x0a, 0x0e, 0x50, 0x6f, 0x6f, 0x6c, 0x47, 0x6f,
	0x6f, 0x64, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x28, 0x20, 0x01, 0x28, 0x09, 0x48, 0x03,
	0x52, 0x0e, 0x50, 0x6f, 0x6f, 0x6c, 0x47, 0x6f, 0x6f, 0x64, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44,
	0x88, 0x01, 0x01, 0x12, 0x19, 0x0a, 0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x32, 0x20, 0x01,
	0x28, 0x09, 0x48, 0x04, 0x52, 0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x88, 0x01, 0x01, 0x12, 0x37,
	0x0a, 0x05, 0x53, 0x74, 0x61, 0x74, 0x65, 0x18, 0x3c, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1c, 0x2e,
	0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x64, 0x2e, 0x76,
	0x31, 0x2e, 0x47, 0x6f, 0x6f, 0x64, 0x53, 0x74, 0x61, 0x74, 0x65, 0x48, 0x05, 0x52, 0x05, 0x53,
	0x74, 0x61, 0x74, 0x65, 0x88, 0x01, 0x01, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x45, 0x6e, 0x74, 0x49,
	0x44, 0x42, 0x11, 0x0a, 0x0f, 0x5f, 0x50, 0x6f, 0x6f, 0x6c, 0x52, 0x6f, 0x6f, 0x74, 0x55, 0x73,
	0x65, 0x72, 0x49, 0x44, 0x42, 0x0f, 0x0a, 0x0d, 0x5f, 0x4d, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x50,
	0x6f, 0x6f, 0x6c, 0x49, 0x44, 0x42, 0x11, 0x0a, 0x0f, 0x5f, 0x50, 0x6f, 0x6f, 0x6c, 0x47, 0x6f,
	0x6f, 0x64, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x54, 0x6f, 0x74,
	0x61, 0x6c, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x53, 0x74, 0x61, 0x74, 0x65, 0x22, 0xbb, 0x03, 0x0a,
	0x0f, 0x4d, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x47, 0x6f, 0x6f, 0x64, 0x53, 0x74, 0x6f, 0x63, 0x6b,
	0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x49, 0x44,
	0x12, 0x14, 0x0a, 0x05, 0x45, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x45, 0x6e, 0x74, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x47, 0x6f, 0x6f, 0x64, 0x49, 0x44,
	0x18, 0x15, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x47, 0x6f, 0x6f, 0x64, 0x49, 0x44, 0x12, 0x20,
	0x0a, 0x0b, 0x47, 0x6f, 0x6f, 0x64, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x49, 0x44, 0x18, 0x1e, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x47, 0x6f, 0x6f, 0x64, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x49, 0x44,
	0x12, 0x22, 0x0a, 0x0c, 0x4d, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x50, 0x6f, 0x6f, 0x6c, 0x49, 0x44,
	0x18, 0x28, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x4d, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x50, 0x6f,
	0x6f, 0x6c, 0x49, 0x44, 0x12, 0x26, 0x0a, 0x0e, 0x50, 0x6f, 0x6f, 0x6c, 0x47, 0x6f, 0x6f, 0x64,
	0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x32, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x50, 0x6f,
	0x6f, 0x6c, 0x47, 0x6f, 0x6f, 0x64, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x26, 0x0a, 0x0e,
	0x50, 0x6f, 0x6f, 0x6c, 0x52, 0x6f, 0x6f, 0x74, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x33,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x50, 0x6f, 0x6f, 0x6c, 0x52, 0x6f, 0x6f, 0x74, 0x55, 0x73,
	0x65, 0x72, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x3c, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x22, 0x0a, 0x0c, 0x53, 0x70,
	0x6f, 0x74, 0x51, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x46, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x53, 0x70, 0x6f, 0x74, 0x51, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x16,
	0x0a, 0x06, 0x4c, 0x6f, 0x63, 0x6b, 0x65, 0x64, 0x18, 0x50, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x4c, 0x6f, 0x63, 0x6b, 0x65, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x57, 0x61, 0x69, 0x74, 0x53, 0x74,
	0x61, 0x72, 0x74, 0x18, 0x5a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x57, 0x61, 0x69, 0x74, 0x53,
	0x74, 0x61, 0x72, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x49, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x18, 0x64, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x49, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x53, 0x6f, 0x6c, 0x64, 0x18, 0x6e, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x53, 0x6f, 0x6c, 0x64, 0x12, 0x32, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x74, 0x65, 0x18,
	0x78, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1c, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x6f, 0x6f, 0x64, 0x53, 0x74,
	0x61, 0x74, 0x65, 0x52, 0x05, 0x53, 0x74, 0x61, 0x74, 0x65, 0x22, 0xd3, 0x01, 0x0a, 0x13, 0x4d,
	0x69, 0x6e, 0x69, 0x6e, 0x67, 0x47, 0x6f, 0x6f, 0x64, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x49, 0x6e,
	0x66, 0x6f, 0x12, 0x14, 0x0a, 0x05, 0x45, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x45, 0x6e, 0x74, 0x49, 0x44, 0x12, 0x20, 0x0a, 0x0b, 0x47, 0x6f, 0x6f, 0x64,
	0x53, 0x74, 0x6f, 0x63, 0x6b, 0x49, 0x44, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x47,
	0x6f, 0x6f, 0x64, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x49, 0x44, 0x12, 0x22, 0x0a, 0x0c, 0x4d, 0x69,
	0x6e, 0x69, 0x6e, 0x67, 0x50, 0x6f, 0x6f, 0x6c, 0x49, 0x44, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x4d, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x50, 0x6f, 0x6f, 0x6c, 0x49, 0x44, 0x12, 0x26,
	0x0a, 0x0e, 0x50, 0x6f, 0x6f, 0x6c, 0x47, 0x6f, 0x6f, 0x64, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44,
	0x18, 0x28, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x50, 0x6f, 0x6f, 0x6c, 0x47, 0x6f, 0x6f, 0x64,
	0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x18,
	0x32, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x22, 0x0a, 0x0c,
	0x53, 0x70, 0x6f, 0x74, 0x51, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x3c, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0c, 0x53, 0x70, 0x6f, 0x74, 0x51, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x22, 0x8d, 0x04, 0x0a, 0x05, 0x43, 0x6f, 0x6e, 0x64, 0x73, 0x12, 0x2c, 0x0a, 0x02, 0x49, 0x44,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70,
	0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x48,
	0x00, 0x52, 0x02, 0x49, 0x44, 0x88, 0x01, 0x01, 0x12, 0x32, 0x0a, 0x05, 0x45, 0x6e, 0x74, 0x49,
	0x44, 0x18, 0x14, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c,
	0x48, 0x01, 0x52, 0x05, 0x45, 0x6e, 0x74, 0x49, 0x44, 0x88, 0x01, 0x01, 0x12, 0x3e, 0x0a, 0x0b,
	0x47, 0x6f, 0x6f, 0x64, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x49, 0x44, 0x18, 0x1e, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x17, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x48, 0x02, 0x52, 0x0b, 0x47, 0x6f,
	0x6f, 0x64, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x49, 0x44, 0x88, 0x01, 0x01, 0x12, 0x40, 0x0a, 0x0c,
	0x4d, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x50, 0x6f, 0x6f, 0x6c, 0x49, 0x44, 0x18, 0x28, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x17, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x48, 0x03, 0x52, 0x0c, 0x4d,
	0x69, 0x6e, 0x69, 0x6e, 0x67, 0x50, 0x6f, 0x6f, 0x6c, 0x49, 0x44, 0x88, 0x01, 0x01, 0x12, 0x44,
	0x0a, 0x0e, 0x50, 0x6f, 0x6f, 0x6c, 0x47, 0x6f, 0x6f, 0x64, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44,
	0x18, 0x32, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70,
	0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x48,
	0x04, 0x52, 0x0e, 0x50, 0x6f, 0x6f, 0x6c, 0x47, 0x6f, 0x6f, 0x64, 0x55, 0x73, 0x65, 0x72, 0x49,
	0x44, 0x88, 0x01, 0x01, 0x12, 0x44, 0x0a, 0x0e, 0x50, 0x6f, 0x6f, 0x6c, 0x52, 0x6f, 0x6f, 0x74,
	0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x3c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x62,
	0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x72, 0x69,
	0x6e, 0x67, 0x56, 0x61, 0x6c, 0x48, 0x05, 0x52, 0x0e, 0x50, 0x6f, 0x6f, 0x6c, 0x52, 0x6f, 0x6f,
	0x74, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x88, 0x01, 0x01, 0x12, 0x32, 0x0a, 0x05, 0x53, 0x74,
	0x61, 0x74, 0x65, 0x18, 0x46, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x62, 0x61, 0x73, 0x65,
	0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56,
	0x61, 0x6c, 0x48, 0x06, 0x52, 0x05, 0x53, 0x74, 0x61, 0x74, 0x65, 0x88, 0x01, 0x01, 0x42, 0x05,
	0x0a, 0x03, 0x5f, 0x49, 0x44, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x45, 0x6e, 0x74, 0x49, 0x44, 0x42,
	0x0e, 0x0a, 0x0c, 0x5f, 0x47, 0x6f, 0x6f, 0x64, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x49, 0x44, 0x42,
	0x0f, 0x0a, 0x0d, 0x5f, 0x4d, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x50, 0x6f, 0x6f, 0x6c, 0x49, 0x44,
	0x42, 0x11, 0x0a, 0x0f, 0x5f, 0x50, 0x6f, 0x6f, 0x6c, 0x47, 0x6f, 0x6f, 0x64, 0x55, 0x73, 0x65,
	0x72, 0x49, 0x44, 0x42, 0x11, 0x0a, 0x0f, 0x5f, 0x50, 0x6f, 0x6f, 0x6c, 0x52, 0x6f, 0x6f, 0x74,
	0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x53, 0x74, 0x61, 0x74, 0x65,
	0x42, 0x3e, 0x5a, 0x3c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4e,
	0x70, 0x6f, 0x6f, 0x6c, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x2f, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x67, 0x6f, 0x6f, 0x64, 0x2f,
	0x6d, 0x77, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x6f, 0x6f, 0x64, 0x2f, 0x73, 0x74, 0x6f, 0x63, 0x6b,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_npool_good_mw_v1_good_stock_stock_proto_rawDescOnce sync.Once
	file_npool_good_mw_v1_good_stock_stock_proto_rawDescData = file_npool_good_mw_v1_good_stock_stock_proto_rawDesc
)

func file_npool_good_mw_v1_good_stock_stock_proto_rawDescGZIP() []byte {
	file_npool_good_mw_v1_good_stock_stock_proto_rawDescOnce.Do(func() {
		file_npool_good_mw_v1_good_stock_stock_proto_rawDescData = protoimpl.X.CompressGZIP(file_npool_good_mw_v1_good_stock_stock_proto_rawDescData)
	})
	return file_npool_good_mw_v1_good_stock_stock_proto_rawDescData
}

var file_npool_good_mw_v1_good_stock_stock_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_npool_good_mw_v1_good_stock_stock_proto_goTypes = []interface{}{
	(*MiningGoodStockReq)(nil),  // 0: good.middleware.good1.stock.v1.MiningGoodStockReq
	(*MiningGoodStock)(nil),     // 1: good.middleware.good1.stock.v1.MiningGoodStock
	(*MiningGoodStockInfo)(nil), // 2: good.middleware.good1.stock.v1.MiningGoodStockInfo
	(*Conds)(nil),               // 3: good.middleware.good1.stock.v1.Conds
	(v1.GoodState)(0),           // 4: basetypes.good.v1.GoodState
	(*v11.Uint32Val)(nil),       // 5: basetypes.v1.Uint32Val
	(*v11.StringVal)(nil),       // 6: basetypes.v1.StringVal
}
var file_npool_good_mw_v1_good_stock_stock_proto_depIdxs = []int32{
	4, // 0: good.middleware.good1.stock.v1.MiningGoodStockReq.State:type_name -> basetypes.good.v1.GoodState
	4, // 1: good.middleware.good1.stock.v1.MiningGoodStock.State:type_name -> basetypes.good.v1.GoodState
	5, // 2: good.middleware.good1.stock.v1.Conds.ID:type_name -> basetypes.v1.Uint32Val
	6, // 3: good.middleware.good1.stock.v1.Conds.EntID:type_name -> basetypes.v1.StringVal
	6, // 4: good.middleware.good1.stock.v1.Conds.GoodStockID:type_name -> basetypes.v1.StringVal
	6, // 5: good.middleware.good1.stock.v1.Conds.MiningPoolID:type_name -> basetypes.v1.StringVal
	6, // 6: good.middleware.good1.stock.v1.Conds.PoolGoodUserID:type_name -> basetypes.v1.StringVal
	6, // 7: good.middleware.good1.stock.v1.Conds.PoolRootUserID:type_name -> basetypes.v1.StringVal
	6, // 8: good.middleware.good1.stock.v1.Conds.State:type_name -> basetypes.v1.StringVal
	9, // [9:9] is the sub-list for method output_type
	9, // [9:9] is the sub-list for method input_type
	9, // [9:9] is the sub-list for extension type_name
	9, // [9:9] is the sub-list for extension extendee
	0, // [0:9] is the sub-list for field type_name
}

func init() { file_npool_good_mw_v1_good_stock_stock_proto_init() }
func file_npool_good_mw_v1_good_stock_stock_proto_init() {
	if File_npool_good_mw_v1_good_stock_stock_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_npool_good_mw_v1_good_stock_stock_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MiningGoodStockReq); i {
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
		file_npool_good_mw_v1_good_stock_stock_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MiningGoodStock); i {
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
		file_npool_good_mw_v1_good_stock_stock_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MiningGoodStockInfo); i {
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
		file_npool_good_mw_v1_good_stock_stock_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
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
	}
	file_npool_good_mw_v1_good_stock_stock_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_npool_good_mw_v1_good_stock_stock_proto_msgTypes[3].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_npool_good_mw_v1_good_stock_stock_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_npool_good_mw_v1_good_stock_stock_proto_goTypes,
		DependencyIndexes: file_npool_good_mw_v1_good_stock_stock_proto_depIdxs,
		MessageInfos:      file_npool_good_mw_v1_good_stock_stock_proto_msgTypes,
	}.Build()
	File_npool_good_mw_v1_good_stock_stock_proto = out.File
	file_npool_good_mw_v1_good_stock_stock_proto_rawDesc = nil
	file_npool_good_mw_v1_good_stock_stock_proto_goTypes = nil
	file_npool_good_mw_v1_good_stock_stock_proto_depIdxs = nil
}
