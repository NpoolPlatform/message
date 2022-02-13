// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.18.1
// source: npool/coininfo/coininfo.proto

package coininfo

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type VersionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info string `protobuf:"bytes,100,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (x *VersionResponse) Reset() {
	*x = VersionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_coininfo_coininfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VersionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VersionResponse) ProtoMessage() {}

func (x *VersionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_coininfo_coininfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VersionResponse.ProtoReflect.Descriptor instead.
func (*VersionResponse) Descriptor() ([]byte, []int) {
	return file_npool_coininfo_coininfo_proto_rawDescGZIP(), []int{0}
}

func (x *VersionResponse) GetInfo() string {
	if x != nil {
		return x.Info
	}
	return ""
}

type CoinInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID             string  `protobuf:"bytes,100,opt,name=ID,proto3" json:"ID,omitempty"`                           // uuid
	PreSale        bool    `protobuf:"varint,110,opt,name=PreSale,proto3" json:"PreSale,omitempty"`                // 是否为预售，false为在售商品
	Name           string  `protobuf:"bytes,120,opt,name=Name,proto3" json:"Name,omitempty"`                       // 币种名称 FIL, BTC
	Unit           string  `protobuf:"bytes,130,opt,name=Unit,proto3" json:"Unit,omitempty"`                       // 单位：FIL
	Logo           string  `protobuf:"bytes,140,opt,name=Logo,proto3" json:"Logo,omitempty"`                       // url; can be empty
	ENV            string  `protobuf:"bytes,150,opt,name=ENV,proto3" json:"ENV,omitempty"`                         // main or test
	ReservedAmount float64 `protobuf:"fixed64,160,opt,name=ReservedAmount,proto3" json:"ReservedAmount,omitempty"` // 预留金额
	CreatedAt      uint32  `protobuf:"varint,170,opt,name=CreatedAt,proto3" json:"CreatedAt,omitempty"`
	UpdatedAt      uint32  `protobuf:"varint,180,opt,name=UpdatedAt,proto3" json:"UpdatedAt,omitempty"`
	ForPay         bool    `protobuf:"varint,190,opt,name=ForPay,proto3" json:"ForPay,omitempty"` // 是否可用作支付货币
}

func (x *CoinInfo) Reset() {
	*x = CoinInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_coininfo_coininfo_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CoinInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CoinInfo) ProtoMessage() {}

func (x *CoinInfo) ProtoReflect() protoreflect.Message {
	mi := &file_npool_coininfo_coininfo_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CoinInfo.ProtoReflect.Descriptor instead.
func (*CoinInfo) Descriptor() ([]byte, []int) {
	return file_npool_coininfo_coininfo_proto_rawDescGZIP(), []int{1}
}

func (x *CoinInfo) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *CoinInfo) GetPreSale() bool {
	if x != nil {
		return x.PreSale
	}
	return false
}

func (x *CoinInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CoinInfo) GetUnit() string {
	if x != nil {
		return x.Unit
	}
	return ""
}

func (x *CoinInfo) GetLogo() string {
	if x != nil {
		return x.Logo
	}
	return ""
}

func (x *CoinInfo) GetENV() string {
	if x != nil {
		return x.ENV
	}
	return ""
}

func (x *CoinInfo) GetReservedAmount() float64 {
	if x != nil {
		return x.ReservedAmount
	}
	return 0
}

func (x *CoinInfo) GetCreatedAt() uint32 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *CoinInfo) GetUpdatedAt() uint32 {
	if x != nil {
		return x.UpdatedAt
	}
	return 0
}

func (x *CoinInfo) GetForPay() bool {
	if x != nil {
		return x.ForPay
	}
	return false
}

type GetCoinInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID   string `protobuf:"bytes,100,opt,name=ID,proto3" json:"ID,omitempty"`
	Name string `protobuf:"bytes,110,opt,name=Name,proto3" json:"Name,omitempty"`
}

func (x *GetCoinInfoRequest) Reset() {
	*x = GetCoinInfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_coininfo_coininfo_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCoinInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCoinInfoRequest) ProtoMessage() {}

func (x *GetCoinInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_coininfo_coininfo_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCoinInfoRequest.ProtoReflect.Descriptor instead.
func (*GetCoinInfoRequest) Descriptor() ([]byte, []int) {
	return file_npool_coininfo_coininfo_proto_rawDescGZIP(), []int{2}
}

func (x *GetCoinInfoRequest) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *GetCoinInfoRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type GetCoinInfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *CoinInfo `protobuf:"bytes,100,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (x *GetCoinInfoResponse) Reset() {
	*x = GetCoinInfoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_coininfo_coininfo_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCoinInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCoinInfoResponse) ProtoMessage() {}

func (x *GetCoinInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_coininfo_coininfo_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCoinInfoResponse.ProtoReflect.Descriptor instead.
func (*GetCoinInfoResponse) Descriptor() ([]byte, []int) {
	return file_npool_coininfo_coininfo_proto_rawDescGZIP(), []int{3}
}

func (x *GetCoinInfoResponse) GetInfo() *CoinInfo {
	if x != nil {
		return x.Info
	}
	return nil
}

type CreateCoinInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PreSale bool   `protobuf:"varint,100,opt,name=PreSale,proto3" json:"PreSale,omitempty"` // 是否为预售，false为在售商品
	Name    string `protobuf:"bytes,110,opt,name=Name,proto3" json:"Name,omitempty"`        // 币种名称 FIL, BTC
	Unit    string `protobuf:"bytes,120,opt,name=Unit,proto3" json:"Unit,omitempty"`        // 单位：FIL
	Logo    string `protobuf:"bytes,130,opt,name=Logo,proto3" json:"Logo,omitempty"`        // url; can be empty
	ENV     string `protobuf:"bytes,140,opt,name=ENV,proto3" json:"ENV,omitempty"`          // main or test
}

func (x *CreateCoinInfoRequest) Reset() {
	*x = CreateCoinInfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_coininfo_coininfo_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCoinInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCoinInfoRequest) ProtoMessage() {}

func (x *CreateCoinInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_coininfo_coininfo_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCoinInfoRequest.ProtoReflect.Descriptor instead.
func (*CreateCoinInfoRequest) Descriptor() ([]byte, []int) {
	return file_npool_coininfo_coininfo_proto_rawDescGZIP(), []int{4}
}

func (x *CreateCoinInfoRequest) GetPreSale() bool {
	if x != nil {
		return x.PreSale
	}
	return false
}

func (x *CreateCoinInfoRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateCoinInfoRequest) GetUnit() string {
	if x != nil {
		return x.Unit
	}
	return ""
}

func (x *CreateCoinInfoRequest) GetLogo() string {
	if x != nil {
		return x.Logo
	}
	return ""
}

func (x *CreateCoinInfoRequest) GetENV() string {
	if x != nil {
		return x.ENV
	}
	return ""
}

type CreateCoinInfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *CoinInfo `protobuf:"bytes,100,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (x *CreateCoinInfoResponse) Reset() {
	*x = CreateCoinInfoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_coininfo_coininfo_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCoinInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCoinInfoResponse) ProtoMessage() {}

func (x *CreateCoinInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_coininfo_coininfo_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCoinInfoResponse.ProtoReflect.Descriptor instead.
func (*CreateCoinInfoResponse) Descriptor() ([]byte, []int) {
	return file_npool_coininfo_coininfo_proto_rawDescGZIP(), []int{5}
}

func (x *CreateCoinInfoResponse) GetInfo() *CoinInfo {
	if x != nil {
		return x.Info
	}
	return nil
}

type GetCoinInfosRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PreSale bool   `protobuf:"varint,100,opt,name=PreSale,proto3" json:"PreSale,omitempty"` // 是否为预售，false为在售商品
	Name    string `protobuf:"bytes,110,opt,name=Name,proto3" json:"Name,omitempty"`        // 币种名称 FIL, BTC
	Offset  int32  `protobuf:"varint,120,opt,name=Offset,proto3" json:"Offset,omitempty"`
	Limit   int32  `protobuf:"varint,130,opt,name=Limit,proto3" json:"Limit,omitempty"`
}

func (x *GetCoinInfosRequest) Reset() {
	*x = GetCoinInfosRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_coininfo_coininfo_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCoinInfosRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCoinInfosRequest) ProtoMessage() {}

func (x *GetCoinInfosRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_coininfo_coininfo_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCoinInfosRequest.ProtoReflect.Descriptor instead.
func (*GetCoinInfosRequest) Descriptor() ([]byte, []int) {
	return file_npool_coininfo_coininfo_proto_rawDescGZIP(), []int{6}
}

func (x *GetCoinInfosRequest) GetPreSale() bool {
	if x != nil {
		return x.PreSale
	}
	return false
}

func (x *GetCoinInfosRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GetCoinInfosRequest) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *GetCoinInfosRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type GetCoinInfosResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Total int32       `protobuf:"varint,100,opt,name=Total,proto3" json:"Total,omitempty"`
	Infos []*CoinInfo `protobuf:"bytes,110,rep,name=Infos,proto3" json:"Infos,omitempty"`
}

func (x *GetCoinInfosResponse) Reset() {
	*x = GetCoinInfosResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_coininfo_coininfo_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCoinInfosResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCoinInfosResponse) ProtoMessage() {}

func (x *GetCoinInfosResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_coininfo_coininfo_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCoinInfosResponse.ProtoReflect.Descriptor instead.
func (*GetCoinInfosResponse) Descriptor() ([]byte, []int) {
	return file_npool_coininfo_coininfo_proto_rawDescGZIP(), []int{7}
}

func (x *GetCoinInfosResponse) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *GetCoinInfosResponse) GetInfos() []*CoinInfo {
	if x != nil {
		return x.Infos
	}
	return nil
}

type UpdateCoinInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID             string  `protobuf:"bytes,100,opt,name=ID,proto3" json:"ID,omitempty"`
	PreSale        bool    `protobuf:"varint,110,opt,name=PreSale,proto3" json:"PreSale,omitempty"` // 是否为预售，false为在售商品
	Logo           string  `protobuf:"bytes,120,opt,name=Logo,proto3" json:"Logo,omitempty"`        // url; can be empty
	ReservedAmount float64 `protobuf:"fixed64,130,opt,name=ReservedAmount,proto3" json:"ReservedAmount,omitempty"`
}

func (x *UpdateCoinInfoRequest) Reset() {
	*x = UpdateCoinInfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_coininfo_coininfo_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateCoinInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateCoinInfoRequest) ProtoMessage() {}

func (x *UpdateCoinInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_coininfo_coininfo_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateCoinInfoRequest.ProtoReflect.Descriptor instead.
func (*UpdateCoinInfoRequest) Descriptor() ([]byte, []int) {
	return file_npool_coininfo_coininfo_proto_rawDescGZIP(), []int{8}
}

func (x *UpdateCoinInfoRequest) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *UpdateCoinInfoRequest) GetPreSale() bool {
	if x != nil {
		return x.PreSale
	}
	return false
}

func (x *UpdateCoinInfoRequest) GetLogo() string {
	if x != nil {
		return x.Logo
	}
	return ""
}

func (x *UpdateCoinInfoRequest) GetReservedAmount() float64 {
	if x != nil {
		return x.ReservedAmount
	}
	return 0
}

type UpdateCoinInfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *CoinInfo `protobuf:"bytes,100,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (x *UpdateCoinInfoResponse) Reset() {
	*x = UpdateCoinInfoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_coininfo_coininfo_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateCoinInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateCoinInfoResponse) ProtoMessage() {}

func (x *UpdateCoinInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_coininfo_coininfo_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateCoinInfoResponse.ProtoReflect.Descriptor instead.
func (*UpdateCoinInfoResponse) Descriptor() ([]byte, []int) {
	return file_npool_coininfo_coininfo_proto_rawDescGZIP(), []int{9}
}

func (x *UpdateCoinInfoResponse) GetInfo() *CoinInfo {
	if x != nil {
		return x.Info
	}
	return nil
}

var File_npool_coininfo_coininfo_proto protoreflect.FileDescriptor

var file_npool_coininfo_coininfo_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x63, 0x6f, 0x69, 0x6e, 0x69, 0x6e, 0x66, 0x6f,
	0x2f, 0x63, 0x6f, 0x69, 0x6e, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x12, 0x73, 0x70, 0x68, 0x69, 0x6e, 0x78, 0x2e, 0x63, 0x6f, 0x69, 0x6e, 0x69, 0x6e, 0x66, 0x6f,
	0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x25,
	0x0a, 0x0f, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x64, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x85, 0x02, 0x0a, 0x08, 0x43, 0x6f, 0x69, 0x6e, 0x49, 0x6e,
	0x66, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x64, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x49, 0x44, 0x12, 0x18, 0x0a, 0x07, 0x50, 0x72, 0x65, 0x53, 0x61, 0x6c, 0x65, 0x18, 0x6e, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x07, 0x50, 0x72, 0x65, 0x53, 0x61, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x4e, 0x61, 0x6d, 0x65, 0x18, 0x78, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x13, 0x0a, 0x04, 0x55, 0x6e, 0x69, 0x74, 0x18, 0x82, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x55, 0x6e, 0x69, 0x74, 0x12, 0x13, 0x0a, 0x04, 0x4c, 0x6f, 0x67, 0x6f, 0x18, 0x8c, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4c, 0x6f, 0x67, 0x6f, 0x12, 0x11, 0x0a, 0x03, 0x45, 0x4e,
	0x56, 0x18, 0x96, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x45, 0x4e, 0x56, 0x12, 0x27, 0x0a,
	0x0e, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x64, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0xa0, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0e, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x64,
	0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1d, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x18, 0xaa, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x18, 0xb4, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x12, 0x17, 0x0a, 0x06, 0x46, 0x6f, 0x72, 0x50, 0x61, 0x79, 0x18, 0xbe,
	0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x46, 0x6f, 0x72, 0x50, 0x61, 0x79, 0x22, 0x38, 0x0a,
	0x12, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x69, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x64, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x6e, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x47, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x43, 0x6f,
	0x69, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x30,
	0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x64, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x73,
	0x70, 0x68, 0x69, 0x6e, 0x78, 0x2e, 0x63, 0x6f, 0x69, 0x6e, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x76,
	0x31, 0x2e, 0x43, 0x6f, 0x69, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x49, 0x6e, 0x66, 0x6f,
	0x22, 0x81, 0x01, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x69, 0x6e, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x50, 0x72,
	0x65, 0x53, 0x61, 0x6c, 0x65, 0x18, 0x64, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x50, 0x72, 0x65,
	0x53, 0x61, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x6e, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x55, 0x6e, 0x69, 0x74,
	0x18, 0x78, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x55, 0x6e, 0x69, 0x74, 0x12, 0x13, 0x0a, 0x04,
	0x4c, 0x6f, 0x67, 0x6f, 0x18, 0x82, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4c, 0x6f, 0x67,
	0x6f, 0x12, 0x11, 0x0a, 0x03, 0x45, 0x4e, 0x56, 0x18, 0x8c, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x45, 0x4e, 0x56, 0x22, 0x4a, 0x0a, 0x16, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f,
	0x69, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x30,
	0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x64, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x73,
	0x70, 0x68, 0x69, 0x6e, 0x78, 0x2e, 0x63, 0x6f, 0x69, 0x6e, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x76,
	0x31, 0x2e, 0x43, 0x6f, 0x69, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x49, 0x6e, 0x66, 0x6f,
	0x22, 0x72, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x69, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x50, 0x72, 0x65, 0x53, 0x61,
	0x6c, 0x65, 0x18, 0x64, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x50, 0x72, 0x65, 0x53, 0x61, 0x6c,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x6e, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18,
	0x78, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x15, 0x0a,
	0x05, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x82, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x4c,
	0x69, 0x6d, 0x69, 0x74, 0x22, 0x60, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x69, 0x6e, 0x49,
	0x6e, 0x66, 0x6f, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x54, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x64, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x54, 0x6f, 0x74,
	0x61, 0x6c, 0x12, 0x32, 0x0a, 0x05, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x18, 0x6e, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x1c, 0x2e, 0x73, 0x70, 0x68, 0x69, 0x6e, 0x78, 0x2e, 0x63, 0x6f, 0x69, 0x6e, 0x69,
	0x6e, 0x66, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x69, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x05, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x22, 0x7e, 0x0a, 0x15, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x43, 0x6f, 0x69, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x64, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x12,
	0x18, 0x0a, 0x07, 0x50, 0x72, 0x65, 0x53, 0x61, 0x6c, 0x65, 0x18, 0x6e, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x07, 0x50, 0x72, 0x65, 0x53, 0x61, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x4c, 0x6f, 0x67,
	0x6f, 0x18, 0x78, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4c, 0x6f, 0x67, 0x6f, 0x12, 0x27, 0x0a,
	0x0e, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x64, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x82, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0e, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x64,
	0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x4a, 0x0a, 0x16, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x43, 0x6f, 0x69, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x30, 0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x64, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c,
	0x2e, 0x73, 0x70, 0x68, 0x69, 0x6e, 0x78, 0x2e, 0x63, 0x6f, 0x69, 0x6e, 0x69, 0x6e, 0x66, 0x6f,
	0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x69, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x49, 0x6e,
	0x66, 0x6f, 0x32, 0xff, 0x04, 0x0a, 0x0e, 0x53, 0x70, 0x68, 0x69, 0x6e, 0x78, 0x43, 0x6f, 0x69,
	0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x5b, 0x0a, 0x07, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x23, 0x2e, 0x73, 0x70, 0x68, 0x69, 0x6e,
	0x78, 0x2e, 0x63, 0x6f, 0x69, 0x6e, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x56, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x13, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x0d, 0x22, 0x08, 0x2f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x3a,
	0x01, 0x2a, 0x12, 0x87, 0x01, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x69,
	0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x29, 0x2e, 0x73, 0x70, 0x68, 0x69, 0x6e, 0x78, 0x2e, 0x63,
	0x6f, 0x69, 0x6e, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x43, 0x6f, 0x69, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x2a, 0x2e, 0x73, 0x70, 0x68, 0x69, 0x6e, 0x78, 0x2e, 0x63, 0x6f, 0x69, 0x6e, 0x69, 0x6e,
	0x66, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x69, 0x6e,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1e, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x18, 0x22, 0x13, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x2f, 0x63, 0x6f, 0x69, 0x6e, 0x69, 0x6e, 0x66, 0x6f, 0x3a, 0x01, 0x2a, 0x12, 0x7b, 0x0a, 0x0b,
	0x47, 0x65, 0x74, 0x43, 0x6f, 0x69, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x26, 0x2e, 0x73, 0x70,
	0x68, 0x69, 0x6e, 0x78, 0x2e, 0x63, 0x6f, 0x69, 0x6e, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x76, 0x31,
	0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x69, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x73, 0x70, 0x68, 0x69, 0x6e, 0x78, 0x2e, 0x63, 0x6f, 0x69,
	0x6e, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x69, 0x6e,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1b, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x15, 0x22, 0x10, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x65, 0x74, 0x2f, 0x63, 0x6f,
	0x69, 0x6e, 0x69, 0x6e, 0x66, 0x6f, 0x3a, 0x01, 0x2a, 0x12, 0x7f, 0x0a, 0x0c, 0x47, 0x65, 0x74,
	0x43, 0x6f, 0x69, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x12, 0x27, 0x2e, 0x73, 0x70, 0x68, 0x69,
	0x6e, 0x78, 0x2e, 0x63, 0x6f, 0x69, 0x6e, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x47,
	0x65, 0x74, 0x43, 0x6f, 0x69, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x28, 0x2e, 0x73, 0x70, 0x68, 0x69, 0x6e, 0x78, 0x2e, 0x63, 0x6f, 0x69, 0x6e,
	0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x69, 0x6e, 0x49,
	0x6e, 0x66, 0x6f, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1c, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x16, 0x22, 0x11, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x65, 0x74, 0x2f, 0x63, 0x6f,
	0x69, 0x6e, 0x69, 0x6e, 0x66, 0x6f, 0x73, 0x3a, 0x01, 0x2a, 0x12, 0x87, 0x01, 0x0a, 0x0e, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x69, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x29, 0x2e,
	0x73, 0x70, 0x68, 0x69, 0x6e, 0x78, 0x2e, 0x63, 0x6f, 0x69, 0x6e, 0x69, 0x6e, 0x66, 0x6f, 0x2e,
	0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x69, 0x6e, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x73, 0x70, 0x68, 0x69, 0x6e,
	0x78, 0x2e, 0x63, 0x6f, 0x69, 0x6e, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x69, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x18, 0x22, 0x13, 0x2f, 0x76,
	0x31, 0x2f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x63, 0x6f, 0x69, 0x6e, 0x69, 0x6e, 0x66,
	0x6f, 0x3a, 0x01, 0x2a, 0x42, 0x31, 0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x4e, 0x70, 0x6f, 0x6f, 0x6c, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d,
	0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x63,
	0x6f, 0x69, 0x6e, 0x69, 0x6e, 0x66, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_npool_coininfo_coininfo_proto_rawDescOnce sync.Once
	file_npool_coininfo_coininfo_proto_rawDescData = file_npool_coininfo_coininfo_proto_rawDesc
)

func file_npool_coininfo_coininfo_proto_rawDescGZIP() []byte {
	file_npool_coininfo_coininfo_proto_rawDescOnce.Do(func() {
		file_npool_coininfo_coininfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_npool_coininfo_coininfo_proto_rawDescData)
	})
	return file_npool_coininfo_coininfo_proto_rawDescData
}

var file_npool_coininfo_coininfo_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_npool_coininfo_coininfo_proto_goTypes = []interface{}{
	(*VersionResponse)(nil),        // 0: sphinx.coininfo.v1.VersionResponse
	(*CoinInfo)(nil),               // 1: sphinx.coininfo.v1.CoinInfo
	(*GetCoinInfoRequest)(nil),     // 2: sphinx.coininfo.v1.GetCoinInfoRequest
	(*GetCoinInfoResponse)(nil),    // 3: sphinx.coininfo.v1.GetCoinInfoResponse
	(*CreateCoinInfoRequest)(nil),  // 4: sphinx.coininfo.v1.CreateCoinInfoRequest
	(*CreateCoinInfoResponse)(nil), // 5: sphinx.coininfo.v1.CreateCoinInfoResponse
	(*GetCoinInfosRequest)(nil),    // 6: sphinx.coininfo.v1.GetCoinInfosRequest
	(*GetCoinInfosResponse)(nil),   // 7: sphinx.coininfo.v1.GetCoinInfosResponse
	(*UpdateCoinInfoRequest)(nil),  // 8: sphinx.coininfo.v1.UpdateCoinInfoRequest
	(*UpdateCoinInfoResponse)(nil), // 9: sphinx.coininfo.v1.UpdateCoinInfoResponse
	(*emptypb.Empty)(nil),          // 10: google.protobuf.Empty
}
var file_npool_coininfo_coininfo_proto_depIdxs = []int32{
	1,  // 0: sphinx.coininfo.v1.GetCoinInfoResponse.Info:type_name -> sphinx.coininfo.v1.CoinInfo
	1,  // 1: sphinx.coininfo.v1.CreateCoinInfoResponse.Info:type_name -> sphinx.coininfo.v1.CoinInfo
	1,  // 2: sphinx.coininfo.v1.GetCoinInfosResponse.Infos:type_name -> sphinx.coininfo.v1.CoinInfo
	1,  // 3: sphinx.coininfo.v1.UpdateCoinInfoResponse.Info:type_name -> sphinx.coininfo.v1.CoinInfo
	10, // 4: sphinx.coininfo.v1.SphinxCoinInfo.Version:input_type -> google.protobuf.Empty
	4,  // 5: sphinx.coininfo.v1.SphinxCoinInfo.CreateCoinInfo:input_type -> sphinx.coininfo.v1.CreateCoinInfoRequest
	2,  // 6: sphinx.coininfo.v1.SphinxCoinInfo.GetCoinInfo:input_type -> sphinx.coininfo.v1.GetCoinInfoRequest
	6,  // 7: sphinx.coininfo.v1.SphinxCoinInfo.GetCoinInfos:input_type -> sphinx.coininfo.v1.GetCoinInfosRequest
	8,  // 8: sphinx.coininfo.v1.SphinxCoinInfo.UpdateCoinInfo:input_type -> sphinx.coininfo.v1.UpdateCoinInfoRequest
	0,  // 9: sphinx.coininfo.v1.SphinxCoinInfo.Version:output_type -> sphinx.coininfo.v1.VersionResponse
	5,  // 10: sphinx.coininfo.v1.SphinxCoinInfo.CreateCoinInfo:output_type -> sphinx.coininfo.v1.CreateCoinInfoResponse
	3,  // 11: sphinx.coininfo.v1.SphinxCoinInfo.GetCoinInfo:output_type -> sphinx.coininfo.v1.GetCoinInfoResponse
	7,  // 12: sphinx.coininfo.v1.SphinxCoinInfo.GetCoinInfos:output_type -> sphinx.coininfo.v1.GetCoinInfosResponse
	9,  // 13: sphinx.coininfo.v1.SphinxCoinInfo.UpdateCoinInfo:output_type -> sphinx.coininfo.v1.UpdateCoinInfoResponse
	9,  // [9:14] is the sub-list for method output_type
	4,  // [4:9] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_npool_coininfo_coininfo_proto_init() }
func file_npool_coininfo_coininfo_proto_init() {
	if File_npool_coininfo_coininfo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_npool_coininfo_coininfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VersionResponse); i {
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
		file_npool_coininfo_coininfo_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CoinInfo); i {
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
		file_npool_coininfo_coininfo_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCoinInfoRequest); i {
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
		file_npool_coininfo_coininfo_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCoinInfoResponse); i {
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
		file_npool_coininfo_coininfo_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateCoinInfoRequest); i {
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
		file_npool_coininfo_coininfo_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateCoinInfoResponse); i {
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
		file_npool_coininfo_coininfo_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCoinInfosRequest); i {
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
		file_npool_coininfo_coininfo_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCoinInfosResponse); i {
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
		file_npool_coininfo_coininfo_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateCoinInfoRequest); i {
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
		file_npool_coininfo_coininfo_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateCoinInfoResponse); i {
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
			RawDescriptor: file_npool_coininfo_coininfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_npool_coininfo_coininfo_proto_goTypes,
		DependencyIndexes: file_npool_coininfo_coininfo_proto_depIdxs,
		MessageInfos:      file_npool_coininfo_coininfo_proto_msgTypes,
	}.Build()
	File_npool_coininfo_coininfo_proto = out.File
	file_npool_coininfo_coininfo_proto_rawDesc = nil
	file_npool_coininfo_coininfo_proto_goTypes = nil
	file_npool_coininfo_coininfo_proto_depIdxs = nil
}
