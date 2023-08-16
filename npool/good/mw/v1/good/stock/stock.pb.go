// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.18.1
// source: npool/good/mw/v1/good/stock/stock.proto

package stock

import (
	_ "github.com/NpoolPlatform/message/npool/basetypes/good/v1"
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

type StockReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID        *string `protobuf:"bytes,10,opt,name=ID,proto3,oneof" json:"ID,omitempty"`
	GoodID    *string `protobuf:"bytes,20,opt,name=GoodID,proto3,oneof" json:"GoodID,omitempty"`
	Total     *string `protobuf:"bytes,30,opt,name=Total,proto3,oneof" json:"Total,omitempty"`
	Locked    *string `protobuf:"bytes,40,opt,name=Locked,proto3,oneof" json:"Locked,omitempty"`
	WaitStart *string `protobuf:"bytes,50,opt,name=WaitStart,proto3,oneof" json:"WaitStart,omitempty"`
	InService *string `protobuf:"bytes,60,opt,name=InService,proto3,oneof" json:"InService,omitempty"`
	AppLocked *string `protobuf:"bytes,70,opt,name=AppLocked,proto3,oneof" json:"AppLocked,omitempty"`
}

func (x *StockReq) Reset() {
	*x = StockReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_good_mw_v1_good_stock_stock_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StockReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StockReq) ProtoMessage() {}

func (x *StockReq) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use StockReq.ProtoReflect.Descriptor instead.
func (*StockReq) Descriptor() ([]byte, []int) {
	return file_npool_good_mw_v1_good_stock_stock_proto_rawDescGZIP(), []int{0}
}

func (x *StockReq) GetID() string {
	if x != nil && x.ID != nil {
		return *x.ID
	}
	return ""
}

func (x *StockReq) GetGoodID() string {
	if x != nil && x.GoodID != nil {
		return *x.GoodID
	}
	return ""
}

func (x *StockReq) GetTotal() string {
	if x != nil && x.Total != nil {
		return *x.Total
	}
	return ""
}

func (x *StockReq) GetLocked() string {
	if x != nil && x.Locked != nil {
		return *x.Locked
	}
	return ""
}

func (x *StockReq) GetWaitStart() string {
	if x != nil && x.WaitStart != nil {
		return *x.WaitStart
	}
	return ""
}

func (x *StockReq) GetInService() string {
	if x != nil && x.InService != nil {
		return *x.InService
	}
	return ""
}

func (x *StockReq) GetAppLocked() string {
	if x != nil && x.AppLocked != nil {
		return *x.AppLocked
	}
	return ""
}

type Stock struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: sql:"id"
	ID string `protobuf:"bytes,10,opt,name=ID,proto3" json:"ID,omitempty" sql:"id"`
	// @inject_tag: sql:"good_id"
	GoodID string `protobuf:"bytes,20,opt,name=GoodID,proto3" json:"GoodID,omitempty" sql:"good_id"`
	// @inject_tag: sql:"good_name"
	GoodName string `protobuf:"bytes,30,opt,name=GoodName,proto3" json:"GoodName,omitempty" sql:"good_name"`
	// @inject_tag: sql:"total"
	Total string `protobuf:"bytes,40,opt,name=Total,proto3" json:"Total,omitempty" sql:"total"`
	// @inject_tag: sql:"locked"
	Locked string `protobuf:"bytes,50,opt,name=Locked,proto3" json:"Locked,omitempty" sql:"locked"`
	// @inject_tag: sql:"wait_start"
	WaitStart string `protobuf:"bytes,60,opt,name=WaitStart,proto3" json:"WaitStart,omitempty" sql:"wait_start"`
	// @inject_tag: sql:"in_service"
	InService string `protobuf:"bytes,70,opt,name=InService,proto3" json:"InService,omitempty" sql:"in_service"`
	// @inject_tag: sql:"sold"
	Sold string `protobuf:"bytes,80,opt,name=Sold,proto3" json:"Sold,omitempty" sql:"sold"`
	// @inject_tag: sql:"app_locked"
	AppLocked string `protobuf:"bytes,90,opt,name=AppLocked,proto3" json:"AppLocked,omitempty" sql:"app_locked"`
	// @inject_tag: sql:"created_at"
	CreatedAt uint32 `protobuf:"varint,1000,opt,name=CreatedAt,proto3" json:"CreatedAt,omitempty" sql:"created_at"`
	// @inject_tag: sql:"updated_at"
	UpdatedAt uint32 `protobuf:"varint,1010,opt,name=UpdatedAt,proto3" json:"UpdatedAt,omitempty" sql:"updated_at"`
}

func (x *Stock) Reset() {
	*x = Stock{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_good_mw_v1_good_stock_stock_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Stock) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Stock) ProtoMessage() {}

func (x *Stock) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use Stock.ProtoReflect.Descriptor instead.
func (*Stock) Descriptor() ([]byte, []int) {
	return file_npool_good_mw_v1_good_stock_stock_proto_rawDescGZIP(), []int{1}
}

func (x *Stock) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *Stock) GetGoodID() string {
	if x != nil {
		return x.GoodID
	}
	return ""
}

func (x *Stock) GetGoodName() string {
	if x != nil {
		return x.GoodName
	}
	return ""
}

func (x *Stock) GetTotal() string {
	if x != nil {
		return x.Total
	}
	return ""
}

func (x *Stock) GetLocked() string {
	if x != nil {
		return x.Locked
	}
	return ""
}

func (x *Stock) GetWaitStart() string {
	if x != nil {
		return x.WaitStart
	}
	return ""
}

func (x *Stock) GetInService() string {
	if x != nil {
		return x.InService
	}
	return ""
}

func (x *Stock) GetSold() string {
	if x != nil {
		return x.Sold
	}
	return ""
}

func (x *Stock) GetAppLocked() string {
	if x != nil {
		return x.AppLocked
	}
	return ""
}

func (x *Stock) GetCreatedAt() uint32 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *Stock) GetUpdatedAt() uint32 {
	if x != nil {
		return x.UpdatedAt
	}
	return 0
}

type Conds struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID      *v1.StringVal      `protobuf:"bytes,10,opt,name=ID,proto3,oneof" json:"ID,omitempty"`
	GoodID  *v1.StringVal      `protobuf:"bytes,20,opt,name=GoodID,proto3,oneof" json:"GoodID,omitempty"`
	GoodIDs *v1.StringSliceVal `protobuf:"bytes,30,opt,name=GoodIDs,proto3,oneof" json:"GoodIDs,omitempty"`
}

func (x *Conds) Reset() {
	*x = Conds{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_good_mw_v1_good_stock_stock_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Conds) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Conds) ProtoMessage() {}

func (x *Conds) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use Conds.ProtoReflect.Descriptor instead.
func (*Conds) Descriptor() ([]byte, []int) {
	return file_npool_good_mw_v1_good_stock_stock_proto_rawDescGZIP(), []int{2}
}

func (x *Conds) GetID() *v1.StringVal {
	if x != nil {
		return x.ID
	}
	return nil
}

func (x *Conds) GetGoodID() *v1.StringVal {
	if x != nil {
		return x.GoodID
	}
	return nil
}

func (x *Conds) GetGoodIDs() *v1.StringSliceVal {
	if x != nil {
		return x.GoodIDs
	}
	return nil
}

type AddStockRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *StockReq `protobuf:"bytes,10,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (x *AddStockRequest) Reset() {
	*x = AddStockRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_good_mw_v1_good_stock_stock_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddStockRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddStockRequest) ProtoMessage() {}

func (x *AddStockRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use AddStockRequest.ProtoReflect.Descriptor instead.
func (*AddStockRequest) Descriptor() ([]byte, []int) {
	return file_npool_good_mw_v1_good_stock_stock_proto_rawDescGZIP(), []int{3}
}

func (x *AddStockRequest) GetInfo() *StockReq {
	if x != nil {
		return x.Info
	}
	return nil
}

type AddStockResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *Stock `protobuf:"bytes,10,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (x *AddStockResponse) Reset() {
	*x = AddStockResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_good_mw_v1_good_stock_stock_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddStockResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddStockResponse) ProtoMessage() {}

func (x *AddStockResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_good_mw_v1_good_stock_stock_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddStockResponse.ProtoReflect.Descriptor instead.
func (*AddStockResponse) Descriptor() ([]byte, []int) {
	return file_npool_good_mw_v1_good_stock_stock_proto_rawDescGZIP(), []int{4}
}

func (x *AddStockResponse) GetInfo() *Stock {
	if x != nil {
		return x.Info
	}
	return nil
}

type SubStockRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *StockReq `protobuf:"bytes,10,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (x *SubStockRequest) Reset() {
	*x = SubStockRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_good_mw_v1_good_stock_stock_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubStockRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubStockRequest) ProtoMessage() {}

func (x *SubStockRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_good_mw_v1_good_stock_stock_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubStockRequest.ProtoReflect.Descriptor instead.
func (*SubStockRequest) Descriptor() ([]byte, []int) {
	return file_npool_good_mw_v1_good_stock_stock_proto_rawDescGZIP(), []int{5}
}

func (x *SubStockRequest) GetInfo() *StockReq {
	if x != nil {
		return x.Info
	}
	return nil
}

type SubStockResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *Stock `protobuf:"bytes,10,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (x *SubStockResponse) Reset() {
	*x = SubStockResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_good_mw_v1_good_stock_stock_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubStockResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubStockResponse) ProtoMessage() {}

func (x *SubStockResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_good_mw_v1_good_stock_stock_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubStockResponse.ProtoReflect.Descriptor instead.
func (*SubStockResponse) Descriptor() ([]byte, []int) {
	return file_npool_good_mw_v1_good_stock_stock_proto_rawDescGZIP(), []int{6}
}

func (x *SubStockResponse) GetInfo() *Stock {
	if x != nil {
		return x.Info
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
	0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xae,
	0x02, 0x0a, 0x08, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x12, 0x13, 0x0a, 0x02, 0x49,
	0x44, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x02, 0x49, 0x44, 0x88, 0x01, 0x01,
	0x12, 0x1b, 0x0a, 0x06, 0x47, 0x6f, 0x6f, 0x64, 0x49, 0x44, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09,
	0x48, 0x01, 0x52, 0x06, 0x47, 0x6f, 0x6f, 0x64, 0x49, 0x44, 0x88, 0x01, 0x01, 0x12, 0x19, 0x0a,
	0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x09, 0x48, 0x02, 0x52, 0x05,
	0x54, 0x6f, 0x74, 0x61, 0x6c, 0x88, 0x01, 0x01, 0x12, 0x1b, 0x0a, 0x06, 0x4c, 0x6f, 0x63, 0x6b,
	0x65, 0x64, 0x18, 0x28, 0x20, 0x01, 0x28, 0x09, 0x48, 0x03, 0x52, 0x06, 0x4c, 0x6f, 0x63, 0x6b,
	0x65, 0x64, 0x88, 0x01, 0x01, 0x12, 0x21, 0x0a, 0x09, 0x57, 0x61, 0x69, 0x74, 0x53, 0x74, 0x61,
	0x72, 0x74, 0x18, 0x32, 0x20, 0x01, 0x28, 0x09, 0x48, 0x04, 0x52, 0x09, 0x57, 0x61, 0x69, 0x74,
	0x53, 0x74, 0x61, 0x72, 0x74, 0x88, 0x01, 0x01, 0x12, 0x21, 0x0a, 0x09, 0x49, 0x6e, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x18, 0x3c, 0x20, 0x01, 0x28, 0x09, 0x48, 0x05, 0x52, 0x09, 0x49,
	0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x88, 0x01, 0x01, 0x12, 0x21, 0x0a, 0x09, 0x41,
	0x70, 0x70, 0x4c, 0x6f, 0x63, 0x6b, 0x65, 0x64, 0x18, 0x46, 0x20, 0x01, 0x28, 0x09, 0x48, 0x06,
	0x52, 0x09, 0x41, 0x70, 0x70, 0x4c, 0x6f, 0x63, 0x6b, 0x65, 0x64, 0x88, 0x01, 0x01, 0x42, 0x05,
	0x0a, 0x03, 0x5f, 0x49, 0x44, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x47, 0x6f, 0x6f, 0x64, 0x49, 0x44,
	0x42, 0x08, 0x0a, 0x06, 0x5f, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x4c,
	0x6f, 0x63, 0x6b, 0x65, 0x64, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x57, 0x61, 0x69, 0x74, 0x53, 0x74,
	0x61, 0x72, 0x74, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x49, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x41, 0x70, 0x70, 0x4c, 0x6f, 0x63, 0x6b, 0x65, 0x64, 0x22,
	0xa5, 0x02, 0x0a, 0x05, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x47, 0x6f, 0x6f,
	0x64, 0x49, 0x44, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x47, 0x6f, 0x6f, 0x64, 0x49,
	0x44, 0x12, 0x1a, 0x0a, 0x08, 0x47, 0x6f, 0x6f, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x1e, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x47, 0x6f, 0x6f, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x28, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x54, 0x6f,
	0x74, 0x61, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x4c, 0x6f, 0x63, 0x6b, 0x65, 0x64, 0x18, 0x32, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x4c, 0x6f, 0x63, 0x6b, 0x65, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x57,
	0x61, 0x69, 0x74, 0x53, 0x74, 0x61, 0x72, 0x74, 0x18, 0x3c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x57, 0x61, 0x69, 0x74, 0x53, 0x74, 0x61, 0x72, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x49, 0x6e, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x18, 0x46, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x49, 0x6e,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x53, 0x6f, 0x6c, 0x64, 0x18,
	0x50, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x53, 0x6f, 0x6c, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x41,
	0x70, 0x70, 0x4c, 0x6f, 0x63, 0x6b, 0x65, 0x64, 0x18, 0x5a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x41, 0x70, 0x70, 0x4c, 0x6f, 0x63, 0x6b, 0x65, 0x64, 0x12, 0x1d, 0x0a, 0x09, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0xe8, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x09, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0xf2, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0xc6, 0x01, 0x0a, 0x05, 0x43, 0x6f, 0x6e, 0x64,
	0x73, 0x12, 0x2c, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e,
	0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x72,
	0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x48, 0x00, 0x52, 0x02, 0x49, 0x44, 0x88, 0x01, 0x01, 0x12,
	0x34, 0x0a, 0x06, 0x47, 0x6f, 0x6f, 0x64, 0x49, 0x44, 0x18, 0x14, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x17, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53,
	0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x48, 0x01, 0x52, 0x06, 0x47, 0x6f, 0x6f, 0x64,
	0x49, 0x44, 0x88, 0x01, 0x01, 0x12, 0x3b, 0x0a, 0x07, 0x47, 0x6f, 0x6f, 0x64, 0x49, 0x44, 0x73,
	0x18, 0x1e, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70,
	0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x53, 0x6c, 0x69, 0x63,
	0x65, 0x56, 0x61, 0x6c, 0x48, 0x02, 0x52, 0x07, 0x47, 0x6f, 0x6f, 0x64, 0x49, 0x44, 0x73, 0x88,
	0x01, 0x01, 0x42, 0x05, 0x0a, 0x03, 0x5f, 0x49, 0x44, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x47, 0x6f,
	0x6f, 0x64, 0x49, 0x44, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x47, 0x6f, 0x6f, 0x64, 0x49, 0x44, 0x73,
	0x22, 0x4f, 0x0a, 0x0f, 0x41, 0x64, 0x64, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x3c, 0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x28, 0x2e, 0x67, 0x6f, 0x6f, 0x64, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77,
	0x61, 0x72, 0x65, 0x2e, 0x67, 0x6f, 0x6f, 0x64, 0x31, 0x2e, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x2e,
	0x76, 0x31, 0x2e, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x52, 0x04, 0x49, 0x6e, 0x66,
	0x6f, 0x22, 0x4d, 0x0a, 0x10, 0x41, 0x64, 0x64, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x39, 0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x67, 0x6f, 0x6f, 0x64, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c,
	0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x67, 0x6f, 0x6f, 0x64, 0x31, 0x2e, 0x73, 0x74, 0x6f, 0x63,
	0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x52, 0x04, 0x49, 0x6e, 0x66, 0x6f,
	0x22, 0x4f, 0x0a, 0x0f, 0x53, 0x75, 0x62, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x3c, 0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x28, 0x2e, 0x67, 0x6f, 0x6f, 0x64, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77,
	0x61, 0x72, 0x65, 0x2e, 0x67, 0x6f, 0x6f, 0x64, 0x31, 0x2e, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x2e,
	0x76, 0x31, 0x2e, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x52, 0x04, 0x49, 0x6e, 0x66,
	0x6f, 0x22, 0x4d, 0x0a, 0x10, 0x53, 0x75, 0x62, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x39, 0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x67, 0x6f, 0x6f, 0x64, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c,
	0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x67, 0x6f, 0x6f, 0x64, 0x31, 0x2e, 0x73, 0x74, 0x6f, 0x63,
	0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x52, 0x04, 0x49, 0x6e, 0x66, 0x6f,
	0x32, 0xee, 0x01, 0x0a, 0x0a, 0x4d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x12,
	0x6f, 0x0a, 0x08, 0x41, 0x64, 0x64, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x12, 0x2f, 0x2e, 0x67, 0x6f,
	0x6f, 0x64, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x67, 0x6f,
	0x6f, 0x64, 0x31, 0x2e, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x64, 0x64,
	0x53, 0x74, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x30, 0x2e, 0x67,
	0x6f, 0x6f, 0x64, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x67,
	0x6f, 0x6f, 0x64, 0x31, 0x2e, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x64,
	0x64, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x6f, 0x0a, 0x08, 0x53, 0x75, 0x62, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x12, 0x2f, 0x2e, 0x67,
	0x6f, 0x6f, 0x64, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x67,
	0x6f, 0x6f, 0x64, 0x31, 0x2e, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x75,
	0x62, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x30, 0x2e,
	0x67, 0x6f, 0x6f, 0x64, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e,
	0x67, 0x6f, 0x6f, 0x64, 0x31, 0x2e, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x53,
	0x75, 0x62, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x42, 0x3e, 0x5a, 0x3c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x4e, 0x70, 0x6f, 0x6f, 0x6c, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x67, 0x6f, 0x6f, 0x64,
	0x2f, 0x6d, 0x77, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x6f, 0x6f, 0x64, 0x2f, 0x73, 0x74, 0x6f, 0x63,
	0x6b, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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

var file_npool_good_mw_v1_good_stock_stock_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_npool_good_mw_v1_good_stock_stock_proto_goTypes = []interface{}{
	(*StockReq)(nil),          // 0: good.middleware.good1.stock.v1.StockReq
	(*Stock)(nil),             // 1: good.middleware.good1.stock.v1.Stock
	(*Conds)(nil),             // 2: good.middleware.good1.stock.v1.Conds
	(*AddStockRequest)(nil),   // 3: good.middleware.good1.stock.v1.AddStockRequest
	(*AddStockResponse)(nil),  // 4: good.middleware.good1.stock.v1.AddStockResponse
	(*SubStockRequest)(nil),   // 5: good.middleware.good1.stock.v1.SubStockRequest
	(*SubStockResponse)(nil),  // 6: good.middleware.good1.stock.v1.SubStockResponse
	(*v1.StringVal)(nil),      // 7: basetypes.v1.StringVal
	(*v1.StringSliceVal)(nil), // 8: basetypes.v1.StringSliceVal
}
var file_npool_good_mw_v1_good_stock_stock_proto_depIdxs = []int32{
	7, // 0: good.middleware.good1.stock.v1.Conds.ID:type_name -> basetypes.v1.StringVal
	7, // 1: good.middleware.good1.stock.v1.Conds.GoodID:type_name -> basetypes.v1.StringVal
	8, // 2: good.middleware.good1.stock.v1.Conds.GoodIDs:type_name -> basetypes.v1.StringSliceVal
	0, // 3: good.middleware.good1.stock.v1.AddStockRequest.Info:type_name -> good.middleware.good1.stock.v1.StockReq
	1, // 4: good.middleware.good1.stock.v1.AddStockResponse.Info:type_name -> good.middleware.good1.stock.v1.Stock
	0, // 5: good.middleware.good1.stock.v1.SubStockRequest.Info:type_name -> good.middleware.good1.stock.v1.StockReq
	1, // 6: good.middleware.good1.stock.v1.SubStockResponse.Info:type_name -> good.middleware.good1.stock.v1.Stock
	3, // 7: good.middleware.good1.stock.v1.Middleware.AddStock:input_type -> good.middleware.good1.stock.v1.AddStockRequest
	5, // 8: good.middleware.good1.stock.v1.Middleware.SubStock:input_type -> good.middleware.good1.stock.v1.SubStockRequest
	4, // 9: good.middleware.good1.stock.v1.Middleware.AddStock:output_type -> good.middleware.good1.stock.v1.AddStockResponse
	6, // 10: good.middleware.good1.stock.v1.Middleware.SubStock:output_type -> good.middleware.good1.stock.v1.SubStockResponse
	9, // [9:11] is the sub-list for method output_type
	7, // [7:9] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_npool_good_mw_v1_good_stock_stock_proto_init() }
func file_npool_good_mw_v1_good_stock_stock_proto_init() {
	if File_npool_good_mw_v1_good_stock_stock_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_npool_good_mw_v1_good_stock_stock_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StockReq); i {
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
			switch v := v.(*Stock); i {
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
		file_npool_good_mw_v1_good_stock_stock_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddStockRequest); i {
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
		file_npool_good_mw_v1_good_stock_stock_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddStockResponse); i {
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
		file_npool_good_mw_v1_good_stock_stock_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubStockRequest); i {
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
		file_npool_good_mw_v1_good_stock_stock_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubStockResponse); i {
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
	file_npool_good_mw_v1_good_stock_stock_proto_msgTypes[2].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_npool_good_mw_v1_good_stock_stock_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
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
