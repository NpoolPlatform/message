// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.18.1
// source: npool/good/mw/v1/appgood/appgood.proto

package appgood

import (
	_ "github.com/NpoolPlatform/message/npool"
	appgood "github.com/NpoolPlatform/message/npool/good/mgr/v1/appgood"
	good "github.com/NpoolPlatform/message/npool/good/mgr/v1/good"
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

type Good struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID                    string           `protobuf:"bytes,10,opt,name=ID,proto3" json:"ID,omitempty"`
	AppID                 string           `protobuf:"bytes,20,opt,name=AppID,proto3" json:"AppID,omitempty"`
	GoodID                string           `protobuf:"bytes,30,opt,name=GoodID,proto3" json:"GoodID,omitempty"`
	Online                bool             `protobuf:"varint,40,opt,name=Online,proto3" json:"Online,omitempty"`
	Visible               bool             `protobuf:"varint,50,opt,name=Visible,proto3" json:"Visible,omitempty"`
	Price                 string           `protobuf:"bytes,60,opt,name=Price,proto3" json:"Price,omitempty"`
	DisplayIndex          int32            `protobuf:"varint,70,opt,name=DisplayIndex,proto3" json:"DisplayIndex,omitempty"`
	PurchaseLimit         int32            `protobuf:"varint,80,opt,name=PurchaseLimit,proto3" json:"PurchaseLimit,omitempty"`
	CommissionPercent     int32            `protobuf:"varint,90,opt,name=CommissionPercent,proto3" json:"CommissionPercent,omitempty"`
	DeviceType            string           `protobuf:"bytes,100,opt,name=DeviceType,proto3" json:"DeviceType,omitempty"`
	DeviceManufacturer    string           `protobuf:"bytes,110,opt,name=DeviceManufacturer,proto3" json:"DeviceManufacturer,omitempty"`
	DevicePowerComsuption uint32           `protobuf:"varint,120,opt,name=DevicePowerComsuption,proto3" json:"DevicePowerComsuption,omitempty"`
	DeviceShipmentAt      uint32           `protobuf:"varint,130,opt,name=DeviceShipmentAt,proto3" json:"DeviceShipmentAt,omitempty"`
	DevicePosters         []string         `protobuf:"bytes,140,rep,name=DevicePosters,proto3" json:"DevicePosters,omitempty"`
	DurationDays          int32            `protobuf:"varint,150,opt,name=DurationDays,proto3" json:"DurationDays,omitempty"`
	CoinTypeID            string           `protobuf:"bytes,160,opt,name=CoinTypeID,proto3" json:"CoinTypeID,omitempty"`
	VendorLocationCountry string           `protobuf:"bytes,170,opt,name=VendorLocationCountry,proto3" json:"VendorLocationCountry,omitempty"`
	GoodType              good.GoodType    `protobuf:"varint,180,opt,name=GoodType,proto3,enum=good.manager.good.v1.GoodType" json:"GoodType,omitempty"`
	BenefitType           good.BenefitType `protobuf:"varint,190,opt,name=BenefitType,proto3,enum=good.manager.good.v1.BenefitType" json:"BenefitType,omitempty"`
	Title                 string           `protobuf:"bytes,200,opt,name=Title,proto3" json:"Title,omitempty"`
	Unit                  string           `protobuf:"bytes,210,opt,name=Unit,proto3" json:"Unit,omitempty"`
	UnitAmount            int32            `protobuf:"varint,220,opt,name=UnitAmount,proto3" json:"UnitAmount,omitempty"`
	SupportCoinTypeIDs    []string         `protobuf:"bytes,230,rep,name=SupportCoinTypeIDs,proto3" json:"SupportCoinTypeIDs,omitempty"`
	TestOnly              bool             `protobuf:"varint,240,opt,name=TestOnly,proto3" json:"TestOnly,omitempty"`
	GoodTotal             uint32           `protobuf:"varint,250,opt,name=GoodTotal,proto3" json:"GoodTotal,omitempty"`
	GoodLocked            uint32           `protobuf:"varint,260,opt,name=GoodLocked,proto3" json:"GoodLocked,omitempty"`
	GoodInService         uint32           `protobuf:"varint,270,opt,name=GoodInService,proto3" json:"GoodInService,omitempty"`
	GoodSold              uint32           `protobuf:"varint,280,opt,name=GoodSold,proto3" json:"GoodSold,omitempty"`
	SubGoods              []*Good          `protobuf:"bytes,290,rep,name=SubGoods,proto3" json:"SubGoods,omitempty"`
	StartAt               uint32           `protobuf:"varint,300,opt,name=StartAt,proto3" json:"StartAt,omitempty"`
	CreatedAt             uint32           `protobuf:"varint,310,opt,name=CreatedAt,proto3" json:"CreatedAt,omitempty"`
}

func (x *Good) Reset() {
	*x = Good{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_good_mw_v1_appgood_appgood_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Good) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Good) ProtoMessage() {}

func (x *Good) ProtoReflect() protoreflect.Message {
	mi := &file_npool_good_mw_v1_appgood_appgood_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Good.ProtoReflect.Descriptor instead.
func (*Good) Descriptor() ([]byte, []int) {
	return file_npool_good_mw_v1_appgood_appgood_proto_rawDescGZIP(), []int{0}
}

func (x *Good) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *Good) GetAppID() string {
	if x != nil {
		return x.AppID
	}
	return ""
}

func (x *Good) GetGoodID() string {
	if x != nil {
		return x.GoodID
	}
	return ""
}

func (x *Good) GetOnline() bool {
	if x != nil {
		return x.Online
	}
	return false
}

func (x *Good) GetVisible() bool {
	if x != nil {
		return x.Visible
	}
	return false
}

func (x *Good) GetPrice() string {
	if x != nil {
		return x.Price
	}
	return ""
}

func (x *Good) GetDisplayIndex() int32 {
	if x != nil {
		return x.DisplayIndex
	}
	return 0
}

func (x *Good) GetPurchaseLimit() int32 {
	if x != nil {
		return x.PurchaseLimit
	}
	return 0
}

func (x *Good) GetCommissionPercent() int32 {
	if x != nil {
		return x.CommissionPercent
	}
	return 0
}

func (x *Good) GetDeviceType() string {
	if x != nil {
		return x.DeviceType
	}
	return ""
}

func (x *Good) GetDeviceManufacturer() string {
	if x != nil {
		return x.DeviceManufacturer
	}
	return ""
}

func (x *Good) GetDevicePowerComsuption() uint32 {
	if x != nil {
		return x.DevicePowerComsuption
	}
	return 0
}

func (x *Good) GetDeviceShipmentAt() uint32 {
	if x != nil {
		return x.DeviceShipmentAt
	}
	return 0
}

func (x *Good) GetDevicePosters() []string {
	if x != nil {
		return x.DevicePosters
	}
	return nil
}

func (x *Good) GetDurationDays() int32 {
	if x != nil {
		return x.DurationDays
	}
	return 0
}

func (x *Good) GetCoinTypeID() string {
	if x != nil {
		return x.CoinTypeID
	}
	return ""
}

func (x *Good) GetVendorLocationCountry() string {
	if x != nil {
		return x.VendorLocationCountry
	}
	return ""
}

func (x *Good) GetGoodType() good.GoodType {
	if x != nil {
		return x.GoodType
	}
	return good.GoodType(0)
}

func (x *Good) GetBenefitType() good.BenefitType {
	if x != nil {
		return x.BenefitType
	}
	return good.BenefitType(0)
}

func (x *Good) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Good) GetUnit() string {
	if x != nil {
		return x.Unit
	}
	return ""
}

func (x *Good) GetUnitAmount() int32 {
	if x != nil {
		return x.UnitAmount
	}
	return 0
}

func (x *Good) GetSupportCoinTypeIDs() []string {
	if x != nil {
		return x.SupportCoinTypeIDs
	}
	return nil
}

func (x *Good) GetTestOnly() bool {
	if x != nil {
		return x.TestOnly
	}
	return false
}

func (x *Good) GetGoodTotal() uint32 {
	if x != nil {
		return x.GoodTotal
	}
	return 0
}

func (x *Good) GetGoodLocked() uint32 {
	if x != nil {
		return x.GoodLocked
	}
	return 0
}

func (x *Good) GetGoodInService() uint32 {
	if x != nil {
		return x.GoodInService
	}
	return 0
}

func (x *Good) GetGoodSold() uint32 {
	if x != nil {
		return x.GoodSold
	}
	return 0
}

func (x *Good) GetSubGoods() []*Good {
	if x != nil {
		return x.SubGoods
	}
	return nil
}

func (x *Good) GetStartAt() uint32 {
	if x != nil {
		return x.StartAt
	}
	return 0
}

func (x *Good) GetCreatedAt() uint32 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

type GetGoodsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Conds  *appgood.Conds `protobuf:"bytes,10,opt,name=Conds,proto3" json:"Conds,omitempty"`
	Offset int32          `protobuf:"varint,20,opt,name=Offset,proto3" json:"Offset,omitempty"`
	Limit  int32          `protobuf:"varint,30,opt,name=Limit,proto3" json:"Limit,omitempty"`
}

func (x *GetGoodsRequest) Reset() {
	*x = GetGoodsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_good_mw_v1_appgood_appgood_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetGoodsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetGoodsRequest) ProtoMessage() {}

func (x *GetGoodsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_good_mw_v1_appgood_appgood_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetGoodsRequest.ProtoReflect.Descriptor instead.
func (*GetGoodsRequest) Descriptor() ([]byte, []int) {
	return file_npool_good_mw_v1_appgood_appgood_proto_rawDescGZIP(), []int{1}
}

func (x *GetGoodsRequest) GetConds() *appgood.Conds {
	if x != nil {
		return x.Conds
	}
	return nil
}

func (x *GetGoodsRequest) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *GetGoodsRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type GetGoodsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Infos []*Good `protobuf:"bytes,10,rep,name=Infos,proto3" json:"Infos,omitempty"`
	Total uint32  `protobuf:"varint,20,opt,name=Total,proto3" json:"Total,omitempty"`
}

func (x *GetGoodsResponse) Reset() {
	*x = GetGoodsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_good_mw_v1_appgood_appgood_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetGoodsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetGoodsResponse) ProtoMessage() {}

func (x *GetGoodsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_good_mw_v1_appgood_appgood_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetGoodsResponse.ProtoReflect.Descriptor instead.
func (*GetGoodsResponse) Descriptor() ([]byte, []int) {
	return file_npool_good_mw_v1_appgood_appgood_proto_rawDescGZIP(), []int{2}
}

func (x *GetGoodsResponse) GetInfos() []*Good {
	if x != nil {
		return x.Infos
	}
	return nil
}

func (x *GetGoodsResponse) GetTotal() uint32 {
	if x != nil {
		return x.Total
	}
	return 0
}

type UpdateGoodRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *appgood.AppGoodReq `protobuf:"bytes,10,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (x *UpdateGoodRequest) Reset() {
	*x = UpdateGoodRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_good_mw_v1_appgood_appgood_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateGoodRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateGoodRequest) ProtoMessage() {}

func (x *UpdateGoodRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_good_mw_v1_appgood_appgood_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateGoodRequest.ProtoReflect.Descriptor instead.
func (*UpdateGoodRequest) Descriptor() ([]byte, []int) {
	return file_npool_good_mw_v1_appgood_appgood_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateGoodRequest) GetInfo() *appgood.AppGoodReq {
	if x != nil {
		return x.Info
	}
	return nil
}

type UpdateGoodResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *Good `protobuf:"bytes,10,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (x *UpdateGoodResponse) Reset() {
	*x = UpdateGoodResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_good_mw_v1_appgood_appgood_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateGoodResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateGoodResponse) ProtoMessage() {}

func (x *UpdateGoodResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_good_mw_v1_appgood_appgood_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateGoodResponse.ProtoReflect.Descriptor instead.
func (*UpdateGoodResponse) Descriptor() ([]byte, []int) {
	return file_npool_good_mw_v1_appgood_appgood_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateGoodResponse) GetInfo() *Good {
	if x != nil {
		return x.Info
	}
	return nil
}

var File_npool_good_mw_v1_appgood_appgood_proto protoreflect.FileDescriptor

var file_npool_good_mw_v1_appgood_appgood_proto_rawDesc = []byte{
	0x0a, 0x26, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x67, 0x6f, 0x6f, 0x64, 0x2f, 0x6d, 0x77, 0x2f,
	0x76, 0x31, 0x2f, 0x61, 0x70, 0x70, 0x67, 0x6f, 0x6f, 0x64, 0x2f, 0x61, 0x70, 0x70, 0x67, 0x6f,
	0x6f, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1a, 0x67, 0x6f, 0x6f, 0x64, 0x2e, 0x6d,
	0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x61, 0x70, 0x70, 0x67, 0x6f, 0x6f,
	0x64, 0x2e, 0x76, 0x31, 0x1a, 0x21, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x67, 0x6f, 0x6f, 0x64,
	0x2f, 0x6d, 0x67, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x6f, 0x6f, 0x64, 0x2f, 0x67, 0x6f, 0x6f,
	0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x27, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x67,
	0x6f, 0x6f, 0x64, 0x2f, 0x6d, 0x67, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x70, 0x70, 0x67, 0x6f,
	0x6f, 0x64, 0x2f, 0x61, 0x70, 0x70, 0x67, 0x6f, 0x6f, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x11, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xf6, 0x08, 0x0a, 0x04, 0x47, 0x6f, 0x6f, 0x64, 0x12, 0x0e, 0x0a, 0x02,
	0x49, 0x44, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05,
	0x41, 0x70, 0x70, 0x49, 0x44, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x41, 0x70, 0x70,
	0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x47, 0x6f, 0x6f, 0x64, 0x49, 0x44, 0x18, 0x1e, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x47, 0x6f, 0x6f, 0x64, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x4f, 0x6e,
	0x6c, 0x69, 0x6e, 0x65, 0x18, 0x28, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x4f, 0x6e, 0x6c, 0x69,
	0x6e, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x56, 0x69, 0x73, 0x69, 0x62, 0x6c, 0x65, 0x18, 0x32, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x07, 0x56, 0x69, 0x73, 0x69, 0x62, 0x6c, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x50, 0x72, 0x69, 0x63, 0x65, 0x18, 0x3c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x50, 0x72, 0x69,
	0x63, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x44, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x49, 0x6e, 0x64,
	0x65, 0x78, 0x18, 0x46, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x44, 0x69, 0x73, 0x70, 0x6c, 0x61,
	0x79, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x24, 0x0a, 0x0d, 0x50, 0x75, 0x72, 0x63, 0x68, 0x61,
	0x73, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x50, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x50,
	0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x2c, 0x0a, 0x11,
	0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x50, 0x65, 0x72, 0x63, 0x65, 0x6e,
	0x74, 0x18, 0x5a, 0x20, 0x01, 0x28, 0x05, 0x52, 0x11, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x50, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x44, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x18, 0x64, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x2e, 0x0a, 0x12, 0x44, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x4d, 0x61, 0x6e, 0x75, 0x66, 0x61, 0x63, 0x74, 0x75, 0x72, 0x65, 0x72,
	0x18, 0x6e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x4d, 0x61,
	0x6e, 0x75, 0x66, 0x61, 0x63, 0x74, 0x75, 0x72, 0x65, 0x72, 0x12, 0x34, 0x0a, 0x15, 0x44, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x50, 0x6f, 0x77, 0x65, 0x72, 0x43, 0x6f, 0x6d, 0x73, 0x75, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x78, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x15, 0x44, 0x65, 0x76, 0x69, 0x63,
	0x65, 0x50, 0x6f, 0x77, 0x65, 0x72, 0x43, 0x6f, 0x6d, 0x73, 0x75, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x2b, 0x0a, 0x10, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x53, 0x68, 0x69, 0x70, 0x6d, 0x65,
	0x6e, 0x74, 0x41, 0x74, 0x18, 0x82, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x10, 0x44, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x53, 0x68, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x41, 0x74, 0x12, 0x25, 0x0a,
	0x0d, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x65, 0x72, 0x73, 0x18, 0x8c,
	0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0d, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x50, 0x6f, 0x73,
	0x74, 0x65, 0x72, 0x73, 0x12, 0x23, 0x0a, 0x0c, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x44, 0x61, 0x79, 0x73, 0x18, 0x96, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x44, 0x75, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x79, 0x73, 0x12, 0x1f, 0x0a, 0x0a, 0x43, 0x6f, 0x69,
	0x6e, 0x54, 0x79, 0x70, 0x65, 0x49, 0x44, 0x18, 0xa0, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x43, 0x6f, 0x69, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x49, 0x44, 0x12, 0x35, 0x0a, 0x15, 0x56, 0x65,
	0x6e, 0x64, 0x6f, 0x72, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x72, 0x79, 0x18, 0xaa, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x15, 0x56, 0x65, 0x6e, 0x64,
	0x6f, 0x72, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x3b, 0x0a, 0x08, 0x47, 0x6f, 0x6f, 0x64, 0x54, 0x79, 0x70, 0x65, 0x18, 0xb4, 0x01,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x1e, 0x2e, 0x67, 0x6f, 0x6f, 0x64, 0x2e, 0x6d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x72, 0x2e, 0x67, 0x6f, 0x6f, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x6f, 0x6f, 0x64,
	0x54, 0x79, 0x70, 0x65, 0x52, 0x08, 0x47, 0x6f, 0x6f, 0x64, 0x54, 0x79, 0x70, 0x65, 0x12, 0x44,
	0x0a, 0x0b, 0x42, 0x65, 0x6e, 0x65, 0x66, 0x69, 0x74, 0x54, 0x79, 0x70, 0x65, 0x18, 0xbe, 0x01,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x21, 0x2e, 0x67, 0x6f, 0x6f, 0x64, 0x2e, 0x6d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x72, 0x2e, 0x67, 0x6f, 0x6f, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x65, 0x6e, 0x65,
	0x66, 0x69, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0b, 0x42, 0x65, 0x6e, 0x65, 0x66, 0x69, 0x74,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x15, 0x0a, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x18, 0xc8, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x13, 0x0a, 0x04, 0x55,
	0x6e, 0x69, 0x74, 0x18, 0xd2, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x55, 0x6e, 0x69, 0x74,
	0x12, 0x1f, 0x0a, 0x0a, 0x55, 0x6e, 0x69, 0x74, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0xdc,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x55, 0x6e, 0x69, 0x74, 0x41, 0x6d, 0x6f, 0x75, 0x6e,
	0x74, 0x12, 0x2f, 0x0a, 0x12, 0x53, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x43, 0x6f, 0x69, 0x6e,
	0x54, 0x79, 0x70, 0x65, 0x49, 0x44, 0x73, 0x18, 0xe6, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x12,
	0x53, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x43, 0x6f, 0x69, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x49,
	0x44, 0x73, 0x12, 0x1b, 0x0a, 0x08, 0x54, 0x65, 0x73, 0x74, 0x4f, 0x6e, 0x6c, 0x79, 0x18, 0xf0,
	0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x54, 0x65, 0x73, 0x74, 0x4f, 0x6e, 0x6c, 0x79, 0x12,
	0x1d, 0x0a, 0x09, 0x47, 0x6f, 0x6f, 0x64, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0xfa, 0x01, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x09, 0x47, 0x6f, 0x6f, 0x64, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x1f,
	0x0a, 0x0a, 0x47, 0x6f, 0x6f, 0x64, 0x4c, 0x6f, 0x63, 0x6b, 0x65, 0x64, 0x18, 0x84, 0x02, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x0a, 0x47, 0x6f, 0x6f, 0x64, 0x4c, 0x6f, 0x63, 0x6b, 0x65, 0x64, 0x12,
	0x25, 0x0a, 0x0d, 0x47, 0x6f, 0x6f, 0x64, 0x49, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x18, 0x8e, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0d, 0x47, 0x6f, 0x6f, 0x64, 0x49, 0x6e, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x1b, 0x0a, 0x08, 0x47, 0x6f, 0x6f, 0x64, 0x53, 0x6f,
	0x6c, 0x64, 0x18, 0x98, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x47, 0x6f, 0x6f, 0x64, 0x53,
	0x6f, 0x6c, 0x64, 0x12, 0x3d, 0x0a, 0x08, 0x53, 0x75, 0x62, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x18,
	0xa2, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x67, 0x6f, 0x6f, 0x64, 0x2e, 0x6d, 0x69,
	0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x61, 0x70, 0x70, 0x67, 0x6f, 0x6f, 0x64,
	0x2e, 0x76, 0x31, 0x2e, 0x47, 0x6f, 0x6f, 0x64, 0x52, 0x08, 0x53, 0x75, 0x62, 0x47, 0x6f, 0x6f,
	0x64, 0x73, 0x12, 0x19, 0x0a, 0x07, 0x53, 0x74, 0x61, 0x72, 0x74, 0x41, 0x74, 0x18, 0xac, 0x02,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x53, 0x74, 0x61, 0x72, 0x74, 0x41, 0x74, 0x12, 0x1d, 0x0a,
	0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0xb6, 0x02, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x75, 0x0a, 0x0f,
	0x47, 0x65, 0x74, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x34, 0x0a, 0x05, 0x43, 0x6f, 0x6e, 0x64, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e,
	0x2e, 0x67, 0x6f, 0x6f, 0x64, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x61, 0x70,
	0x70, 0x67, 0x6f, 0x6f, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x64, 0x73, 0x52, 0x05,
	0x43, 0x6f, 0x6e, 0x64, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18,
	0x14, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x14, 0x0a,
	0x05, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x4c, 0x69,
	0x6d, 0x69, 0x74, 0x22, 0x60, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x36, 0x0a, 0x05, 0x49, 0x6e, 0x66, 0x6f, 0x73,
	0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x67, 0x6f, 0x6f, 0x64, 0x2e, 0x6d, 0x69,
	0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x61, 0x70, 0x70, 0x67, 0x6f, 0x6f, 0x64,
	0x2e, 0x76, 0x31, 0x2e, 0x47, 0x6f, 0x6f, 0x64, 0x52, 0x05, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x12,
	0x14, 0x0a, 0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x14, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05,
	0x54, 0x6f, 0x74, 0x61, 0x6c, 0x22, 0x4c, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x47,
	0x6f, 0x6f, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x37, 0x0a, 0x04, 0x49, 0x6e,
	0x66, 0x6f, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x67, 0x6f, 0x6f, 0x64, 0x2e,
	0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x70, 0x67, 0x6f, 0x6f, 0x64, 0x2e,
	0x76, 0x31, 0x2e, 0x41, 0x70, 0x70, 0x47, 0x6f, 0x6f, 0x64, 0x52, 0x65, 0x71, 0x52, 0x04, 0x49,
	0x6e, 0x66, 0x6f, 0x22, 0x4a, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x47, 0x6f, 0x6f,
	0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x34, 0x0a, 0x04, 0x49, 0x6e, 0x66,
	0x6f, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x67, 0x6f, 0x6f, 0x64, 0x2e, 0x6d,
	0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x61, 0x70, 0x70, 0x67, 0x6f, 0x6f,
	0x64, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x6f, 0x6f, 0x64, 0x52, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x32,
	0xe4, 0x01, 0x0a, 0x0a, 0x4d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x12, 0x67,
	0x0a, 0x08, 0x47, 0x65, 0x74, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x12, 0x2b, 0x2e, 0x67, 0x6f, 0x6f,
	0x64, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x61, 0x70, 0x70,
	0x67, 0x6f, 0x6f, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x47, 0x6f, 0x6f, 0x64, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2c, 0x2e, 0x67, 0x6f, 0x6f, 0x64, 0x2e, 0x6d,
	0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x61, 0x70, 0x70, 0x67, 0x6f, 0x6f,
	0x64, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x6d, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x47, 0x6f, 0x6f, 0x64, 0x12, 0x2d, 0x2e, 0x67, 0x6f, 0x6f, 0x64, 0x2e, 0x6d, 0x69, 0x64,
	0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x61, 0x70, 0x70, 0x67, 0x6f, 0x6f, 0x64, 0x2e,
	0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x47, 0x6f, 0x6f, 0x64, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x2e, 0x2e, 0x67, 0x6f, 0x6f, 0x64, 0x2e, 0x6d, 0x69, 0x64, 0x64,
	0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x61, 0x70, 0x70, 0x67, 0x6f, 0x6f, 0x64, 0x2e, 0x76,
	0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x47, 0x6f, 0x6f, 0x64, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x3b, 0x5a, 0x39, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4e, 0x70, 0x6f, 0x6f, 0x6c, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f,
	0x72, 0x6d, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x6e, 0x70, 0x6f, 0x6f, 0x6c,
	0x2f, 0x67, 0x6f, 0x6f, 0x64, 0x2f, 0x6d, 0x77, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x70, 0x70, 0x67,
	0x6f, 0x6f, 0x64, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_npool_good_mw_v1_appgood_appgood_proto_rawDescOnce sync.Once
	file_npool_good_mw_v1_appgood_appgood_proto_rawDescData = file_npool_good_mw_v1_appgood_appgood_proto_rawDesc
)

func file_npool_good_mw_v1_appgood_appgood_proto_rawDescGZIP() []byte {
	file_npool_good_mw_v1_appgood_appgood_proto_rawDescOnce.Do(func() {
		file_npool_good_mw_v1_appgood_appgood_proto_rawDescData = protoimpl.X.CompressGZIP(file_npool_good_mw_v1_appgood_appgood_proto_rawDescData)
	})
	return file_npool_good_mw_v1_appgood_appgood_proto_rawDescData
}

var file_npool_good_mw_v1_appgood_appgood_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_npool_good_mw_v1_appgood_appgood_proto_goTypes = []interface{}{
	(*Good)(nil),               // 0: good.middleware.appgood.v1.Good
	(*GetGoodsRequest)(nil),    // 1: good.middleware.appgood.v1.GetGoodsRequest
	(*GetGoodsResponse)(nil),   // 2: good.middleware.appgood.v1.GetGoodsResponse
	(*UpdateGoodRequest)(nil),  // 3: good.middleware.appgood.v1.UpdateGoodRequest
	(*UpdateGoodResponse)(nil), // 4: good.middleware.appgood.v1.UpdateGoodResponse
	(good.GoodType)(0),         // 5: good.manager.good.v1.GoodType
	(good.BenefitType)(0),      // 6: good.manager.good.v1.BenefitType
	(*appgood.Conds)(nil),      // 7: good.manager.appgood.v1.Conds
	(*appgood.AppGoodReq)(nil), // 8: good.manager.appgood.v1.AppGoodReq
}
var file_npool_good_mw_v1_appgood_appgood_proto_depIdxs = []int32{
	5, // 0: good.middleware.appgood.v1.Good.GoodType:type_name -> good.manager.good.v1.GoodType
	6, // 1: good.middleware.appgood.v1.Good.BenefitType:type_name -> good.manager.good.v1.BenefitType
	0, // 2: good.middleware.appgood.v1.Good.SubGoods:type_name -> good.middleware.appgood.v1.Good
	7, // 3: good.middleware.appgood.v1.GetGoodsRequest.Conds:type_name -> good.manager.appgood.v1.Conds
	0, // 4: good.middleware.appgood.v1.GetGoodsResponse.Infos:type_name -> good.middleware.appgood.v1.Good
	8, // 5: good.middleware.appgood.v1.UpdateGoodRequest.Info:type_name -> good.manager.appgood.v1.AppGoodReq
	0, // 6: good.middleware.appgood.v1.UpdateGoodResponse.Info:type_name -> good.middleware.appgood.v1.Good
	1, // 7: good.middleware.appgood.v1.Middleware.GetGoods:input_type -> good.middleware.appgood.v1.GetGoodsRequest
	3, // 8: good.middleware.appgood.v1.Middleware.UpdateGood:input_type -> good.middleware.appgood.v1.UpdateGoodRequest
	2, // 9: good.middleware.appgood.v1.Middleware.GetGoods:output_type -> good.middleware.appgood.v1.GetGoodsResponse
	4, // 10: good.middleware.appgood.v1.Middleware.UpdateGood:output_type -> good.middleware.appgood.v1.UpdateGoodResponse
	9, // [9:11] is the sub-list for method output_type
	7, // [7:9] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_npool_good_mw_v1_appgood_appgood_proto_init() }
func file_npool_good_mw_v1_appgood_appgood_proto_init() {
	if File_npool_good_mw_v1_appgood_appgood_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_npool_good_mw_v1_appgood_appgood_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Good); i {
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
		file_npool_good_mw_v1_appgood_appgood_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetGoodsRequest); i {
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
		file_npool_good_mw_v1_appgood_appgood_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetGoodsResponse); i {
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
		file_npool_good_mw_v1_appgood_appgood_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateGoodRequest); i {
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
		file_npool_good_mw_v1_appgood_appgood_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateGoodResponse); i {
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
			RawDescriptor: file_npool_good_mw_v1_appgood_appgood_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_npool_good_mw_v1_appgood_appgood_proto_goTypes,
		DependencyIndexes: file_npool_good_mw_v1_appgood_appgood_proto_depIdxs,
		MessageInfos:      file_npool_good_mw_v1_appgood_appgood_proto_msgTypes,
	}.Build()
	File_npool_good_mw_v1_appgood_appgood_proto = out.File
	file_npool_good_mw_v1_appgood_appgood_proto_rawDesc = nil
	file_npool_good_mw_v1_appgood_appgood_proto_goTypes = nil
	file_npool_good_mw_v1_appgood_appgood_proto_depIdxs = nil
}
