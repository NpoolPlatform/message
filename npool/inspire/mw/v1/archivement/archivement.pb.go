// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.18.1
// source: npool/inspire/mw/v1/archivement/archivement.proto

package archivement

import (
	v1 "github.com/NpoolPlatform/message/npool/basetypes/v1"
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

type ArchivementReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID              *string `protobuf:"bytes,10,opt,name=ID,proto3,oneof" json:"ID,omitempty"`
	AppID           *string `protobuf:"bytes,20,opt,name=AppID,proto3,oneof" json:"AppID,omitempty"`
	UserID          *string `protobuf:"bytes,30,opt,name=UserID,proto3,oneof" json:"UserID,omitempty"`
	GoodID          *string `protobuf:"bytes,40,opt,name=GoodID,proto3,oneof" json:"GoodID,omitempty"`
	CoinTypeID      *string `protobuf:"bytes,50,opt,name=CoinTypeID,proto3,oneof" json:"CoinTypeID,omitempty"`
	TotalAmount     *string `protobuf:"bytes,60,opt,name=TotalAmount,proto3,oneof" json:"TotalAmount,omitempty"`
	SelfAmount      *string `protobuf:"bytes,70,opt,name=SelfAmount,proto3,oneof" json:"SelfAmount,omitempty"`
	TotalUnits      *string `protobuf:"bytes,80,opt,name=TotalUnits,proto3,oneof" json:"TotalUnits,omitempty"`
	SelfUnits       *string `protobuf:"bytes,90,opt,name=SelfUnits,proto3,oneof" json:"SelfUnits,omitempty"`
	TotalCommission *string `protobuf:"bytes,100,opt,name=TotalCommission,proto3,oneof" json:"TotalCommission,omitempty"`
	SelfCommission  *string `protobuf:"bytes,110,opt,name=SelfCommission,proto3,oneof" json:"SelfCommission,omitempty"`
}

func (x *ArchivementReq) Reset() {
	*x = ArchivementReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_inspire_mw_v1_archivement_archivement_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ArchivementReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ArchivementReq) ProtoMessage() {}

func (x *ArchivementReq) ProtoReflect() protoreflect.Message {
	mi := &file_npool_inspire_mw_v1_archivement_archivement_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ArchivementReq.ProtoReflect.Descriptor instead.
func (*ArchivementReq) Descriptor() ([]byte, []int) {
	return file_npool_inspire_mw_v1_archivement_archivement_proto_rawDescGZIP(), []int{0}
}

func (x *ArchivementReq) GetID() string {
	if x != nil && x.ID != nil {
		return *x.ID
	}
	return ""
}

func (x *ArchivementReq) GetAppID() string {
	if x != nil && x.AppID != nil {
		return *x.AppID
	}
	return ""
}

func (x *ArchivementReq) GetUserID() string {
	if x != nil && x.UserID != nil {
		return *x.UserID
	}
	return ""
}

func (x *ArchivementReq) GetGoodID() string {
	if x != nil && x.GoodID != nil {
		return *x.GoodID
	}
	return ""
}

func (x *ArchivementReq) GetCoinTypeID() string {
	if x != nil && x.CoinTypeID != nil {
		return *x.CoinTypeID
	}
	return ""
}

func (x *ArchivementReq) GetTotalAmount() string {
	if x != nil && x.TotalAmount != nil {
		return *x.TotalAmount
	}
	return ""
}

func (x *ArchivementReq) GetSelfAmount() string {
	if x != nil && x.SelfAmount != nil {
		return *x.SelfAmount
	}
	return ""
}

func (x *ArchivementReq) GetTotalUnits() string {
	if x != nil && x.TotalUnits != nil {
		return *x.TotalUnits
	}
	return ""
}

func (x *ArchivementReq) GetSelfUnits() string {
	if x != nil && x.SelfUnits != nil {
		return *x.SelfUnits
	}
	return ""
}

func (x *ArchivementReq) GetTotalCommission() string {
	if x != nil && x.TotalCommission != nil {
		return *x.TotalCommission
	}
	return ""
}

func (x *ArchivementReq) GetSelfCommission() string {
	if x != nil && x.SelfCommission != nil {
		return *x.SelfCommission
	}
	return ""
}

type Archivement struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID     string `protobuf:"bytes,10,opt,name=ID,proto3" json:"ID,omitempty"`
	AppID  string `protobuf:"bytes,20,opt,name=AppID,proto3" json:"AppID,omitempty"`
	UserID string `protobuf:"bytes,30,opt,name=UserID,proto3" json:"UserID,omitempty"`
	GoodID string `protobuf:"bytes,40,opt,name=GoodID,proto3" json:"GoodID,omitempty"`
	// CoinTypeID of the good
	CoinTypeID string `protobuf:"bytes,50,opt,name=CoinTypeID,proto3" json:"CoinTypeID,omitempty"`
	// Payment amount in USD
	TotalAmount string `protobuf:"bytes,60,opt,name=TotalAmount,proto3" json:"TotalAmount,omitempty"`
	SelfAmount  string `protobuf:"bytes,70,opt,name=SelfAmount,proto3" json:"SelfAmount,omitempty"`
	TotalUnits  string `protobuf:"bytes,80,opt,name=TotalUnits,proto3" json:"TotalUnits,omitempty"`
	SelfUnits   string `protobuf:"bytes,90,opt,name=SelfUnits,proto3" json:"SelfUnits,omitempty"`
	// Commission amount in USD
	TotalCommission string `protobuf:"bytes,100,opt,name=TotalCommission,proto3" json:"TotalCommission,omitempty"`
	SelfCommission  string `protobuf:"bytes,110,opt,name=SelfCommission,proto3" json:"SelfCommission,omitempty"`
}

func (x *Archivement) Reset() {
	*x = Archivement{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_inspire_mw_v1_archivement_archivement_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Archivement) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Archivement) ProtoMessage() {}

func (x *Archivement) ProtoReflect() protoreflect.Message {
	mi := &file_npool_inspire_mw_v1_archivement_archivement_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Archivement.ProtoReflect.Descriptor instead.
func (*Archivement) Descriptor() ([]byte, []int) {
	return file_npool_inspire_mw_v1_archivement_archivement_proto_rawDescGZIP(), []int{1}
}

func (x *Archivement) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *Archivement) GetAppID() string {
	if x != nil {
		return x.AppID
	}
	return ""
}

func (x *Archivement) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

func (x *Archivement) GetGoodID() string {
	if x != nil {
		return x.GoodID
	}
	return ""
}

func (x *Archivement) GetCoinTypeID() string {
	if x != nil {
		return x.CoinTypeID
	}
	return ""
}

func (x *Archivement) GetTotalAmount() string {
	if x != nil {
		return x.TotalAmount
	}
	return ""
}

func (x *Archivement) GetSelfAmount() string {
	if x != nil {
		return x.SelfAmount
	}
	return ""
}

func (x *Archivement) GetTotalUnits() string {
	if x != nil {
		return x.TotalUnits
	}
	return ""
}

func (x *Archivement) GetSelfUnits() string {
	if x != nil {
		return x.SelfUnits
	}
	return ""
}

func (x *Archivement) GetTotalCommission() string {
	if x != nil {
		return x.TotalCommission
	}
	return ""
}

func (x *Archivement) GetSelfCommission() string {
	if x != nil {
		return x.SelfCommission
	}
	return ""
}

type Conds struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID         *v1.StringVal      `protobuf:"bytes,10,opt,name=ID,proto3,oneof" json:"ID,omitempty"`
	AppID      *v1.StringVal      `protobuf:"bytes,20,opt,name=AppID,proto3,oneof" json:"AppID,omitempty"`
	UserID     *v1.StringVal      `protobuf:"bytes,30,opt,name=UserID,proto3,oneof" json:"UserID,omitempty"`
	GoodID     *v1.StringVal      `protobuf:"bytes,40,opt,name=GoodID,proto3,oneof" json:"GoodID,omitempty"`
	CoinTypeID *v1.StringVal      `protobuf:"bytes,50,opt,name=CoinTypeID,proto3,oneof" json:"CoinTypeID,omitempty"`
	UserIDs    *v1.StringSliceVal `protobuf:"bytes,120,opt,name=UserIDs,proto3,oneof" json:"UserIDs,omitempty"`
}

func (x *Conds) Reset() {
	*x = Conds{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_inspire_mw_v1_archivement_archivement_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Conds) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Conds) ProtoMessage() {}

func (x *Conds) ProtoReflect() protoreflect.Message {
	mi := &file_npool_inspire_mw_v1_archivement_archivement_proto_msgTypes[2]
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
	return file_npool_inspire_mw_v1_archivement_archivement_proto_rawDescGZIP(), []int{2}
}

func (x *Conds) GetID() *v1.StringVal {
	if x != nil {
		return x.ID
	}
	return nil
}

func (x *Conds) GetAppID() *v1.StringVal {
	if x != nil {
		return x.AppID
	}
	return nil
}

func (x *Conds) GetUserID() *v1.StringVal {
	if x != nil {
		return x.UserID
	}
	return nil
}

func (x *Conds) GetGoodID() *v1.StringVal {
	if x != nil {
		return x.GoodID
	}
	return nil
}

func (x *Conds) GetCoinTypeID() *v1.StringVal {
	if x != nil {
		return x.CoinTypeID
	}
	return nil
}

func (x *Conds) GetUserIDs() *v1.StringSliceVal {
	if x != nil {
		return x.UserIDs
	}
	return nil
}

type ExpropriateArchivementRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderID string `protobuf:"bytes,10,opt,name=OrderID,proto3" json:"OrderID,omitempty"`
}

func (x *ExpropriateArchivementRequest) Reset() {
	*x = ExpropriateArchivementRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_inspire_mw_v1_archivement_archivement_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExpropriateArchivementRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExpropriateArchivementRequest) ProtoMessage() {}

func (x *ExpropriateArchivementRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_inspire_mw_v1_archivement_archivement_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExpropriateArchivementRequest.ProtoReflect.Descriptor instead.
func (*ExpropriateArchivementRequest) Descriptor() ([]byte, []int) {
	return file_npool_inspire_mw_v1_archivement_archivement_proto_rawDescGZIP(), []int{3}
}

func (x *ExpropriateArchivementRequest) GetOrderID() string {
	if x != nil {
		return x.OrderID
	}
	return ""
}

type ExpropriateArchivementResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ExpropriateArchivementResponse) Reset() {
	*x = ExpropriateArchivementResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_inspire_mw_v1_archivement_archivement_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExpropriateArchivementResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExpropriateArchivementResponse) ProtoMessage() {}

func (x *ExpropriateArchivementResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_inspire_mw_v1_archivement_archivement_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExpropriateArchivementResponse.ProtoReflect.Descriptor instead.
func (*ExpropriateArchivementResponse) Descriptor() ([]byte, []int) {
	return file_npool_inspire_mw_v1_archivement_archivement_proto_rawDescGZIP(), []int{4}
}

type GetArchivementsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Conds  *Conds `protobuf:"bytes,10,opt,name=Conds,proto3" json:"Conds,omitempty"`
	Offset int32  `protobuf:"varint,20,opt,name=Offset,proto3" json:"Offset,omitempty"`
	Limit  int32  `protobuf:"varint,30,opt,name=Limit,proto3" json:"Limit,omitempty"`
}

func (x *GetArchivementsRequest) Reset() {
	*x = GetArchivementsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_inspire_mw_v1_archivement_archivement_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetArchivementsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetArchivementsRequest) ProtoMessage() {}

func (x *GetArchivementsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_inspire_mw_v1_archivement_archivement_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetArchivementsRequest.ProtoReflect.Descriptor instead.
func (*GetArchivementsRequest) Descriptor() ([]byte, []int) {
	return file_npool_inspire_mw_v1_archivement_archivement_proto_rawDescGZIP(), []int{5}
}

func (x *GetArchivementsRequest) GetConds() *Conds {
	if x != nil {
		return x.Conds
	}
	return nil
}

func (x *GetArchivementsRequest) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *GetArchivementsRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type GetArchivementsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Infos []*Archivement `protobuf:"bytes,10,rep,name=Infos,proto3" json:"Infos,omitempty"`
	Total uint32         `protobuf:"varint,20,opt,name=Total,proto3" json:"Total,omitempty"`
}

func (x *GetArchivementsResponse) Reset() {
	*x = GetArchivementsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_inspire_mw_v1_archivement_archivement_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetArchivementsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetArchivementsResponse) ProtoMessage() {}

func (x *GetArchivementsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_inspire_mw_v1_archivement_archivement_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetArchivementsResponse.ProtoReflect.Descriptor instead.
func (*GetArchivementsResponse) Descriptor() ([]byte, []int) {
	return file_npool_inspire_mw_v1_archivement_archivement_proto_rawDescGZIP(), []int{6}
}

func (x *GetArchivementsResponse) GetInfos() []*Archivement {
	if x != nil {
		return x.Infos
	}
	return nil
}

func (x *GetArchivementsResponse) GetTotal() uint32 {
	if x != nil {
		return x.Total
	}
	return 0
}

var File_npool_inspire_mw_v1_archivement_archivement_proto protoreflect.FileDescriptor

var file_npool_inspire_mw_v1_archivement_archivement_proto_rawDesc = []byte{
	0x0a, 0x31, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x69, 0x6e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2f,
	0x6d, 0x77, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x2f, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x21, 0x69, 0x6e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x6d, 0x69, 0x64,
	0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x6d,
	0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x62, 0x61, 0x73, 0x65,
	0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa8, 0x04, 0x0a, 0x0e, 0x41, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65,
	0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x12, 0x13, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x02, 0x49, 0x44, 0x88, 0x01, 0x01, 0x12, 0x19, 0x0a, 0x05,
	0x41, 0x70, 0x70, 0x49, 0x44, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x05, 0x41,
	0x70, 0x70, 0x49, 0x44, 0x88, 0x01, 0x01, 0x12, 0x1b, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49,
	0x44, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x09, 0x48, 0x02, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49,
	0x44, 0x88, 0x01, 0x01, 0x12, 0x1b, 0x0a, 0x06, 0x47, 0x6f, 0x6f, 0x64, 0x49, 0x44, 0x18, 0x28,
	0x20, 0x01, 0x28, 0x09, 0x48, 0x03, 0x52, 0x06, 0x47, 0x6f, 0x6f, 0x64, 0x49, 0x44, 0x88, 0x01,
	0x01, 0x12, 0x23, 0x0a, 0x0a, 0x43, 0x6f, 0x69, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x49, 0x44, 0x18,
	0x32, 0x20, 0x01, 0x28, 0x09, 0x48, 0x04, 0x52, 0x0a, 0x43, 0x6f, 0x69, 0x6e, 0x54, 0x79, 0x70,
	0x65, 0x49, 0x44, 0x88, 0x01, 0x01, 0x12, 0x25, 0x0a, 0x0b, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x41,
	0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x3c, 0x20, 0x01, 0x28, 0x09, 0x48, 0x05, 0x52, 0x0b, 0x54,
	0x6f, 0x74, 0x61, 0x6c, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x88, 0x01, 0x01, 0x12, 0x23, 0x0a,
	0x0a, 0x53, 0x65, 0x6c, 0x66, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x46, 0x20, 0x01, 0x28,
	0x09, 0x48, 0x06, 0x52, 0x0a, 0x53, 0x65, 0x6c, 0x66, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x88,
	0x01, 0x01, 0x12, 0x23, 0x0a, 0x0a, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x55, 0x6e, 0x69, 0x74, 0x73,
	0x18, 0x50, 0x20, 0x01, 0x28, 0x09, 0x48, 0x07, 0x52, 0x0a, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x55,
	0x6e, 0x69, 0x74, 0x73, 0x88, 0x01, 0x01, 0x12, 0x21, 0x0a, 0x09, 0x53, 0x65, 0x6c, 0x66, 0x55,
	0x6e, 0x69, 0x74, 0x73, 0x18, 0x5a, 0x20, 0x01, 0x28, 0x09, 0x48, 0x08, 0x52, 0x09, 0x53, 0x65,
	0x6c, 0x66, 0x55, 0x6e, 0x69, 0x74, 0x73, 0x88, 0x01, 0x01, 0x12, 0x2d, 0x0a, 0x0f, 0x54, 0x6f,
	0x74, 0x61, 0x6c, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x64, 0x20,
	0x01, 0x28, 0x09, 0x48, 0x09, 0x52, 0x0f, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x43, 0x6f, 0x6d, 0x6d,
	0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x88, 0x01, 0x01, 0x12, 0x2b, 0x0a, 0x0e, 0x53, 0x65, 0x6c,
	0x66, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x6e, 0x20, 0x01, 0x28,
	0x09, 0x48, 0x0a, 0x52, 0x0e, 0x53, 0x65, 0x6c, 0x66, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x88, 0x01, 0x01, 0x42, 0x05, 0x0a, 0x03, 0x5f, 0x49, 0x44, 0x42, 0x08, 0x0a,
	0x06, 0x5f, 0x41, 0x70, 0x70, 0x49, 0x44, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x55, 0x73, 0x65, 0x72,
	0x49, 0x44, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x47, 0x6f, 0x6f, 0x64, 0x49, 0x44, 0x42, 0x0d, 0x0a,
	0x0b, 0x5f, 0x43, 0x6f, 0x69, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x49, 0x44, 0x42, 0x0e, 0x0a, 0x0c,
	0x5f, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x0d, 0x0a, 0x0b,
	0x5f, 0x53, 0x65, 0x6c, 0x66, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x0d, 0x0a, 0x0b, 0x5f,
	0x54, 0x6f, 0x74, 0x61, 0x6c, 0x55, 0x6e, 0x69, 0x74, 0x73, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x53,
	0x65, 0x6c, 0x66, 0x55, 0x6e, 0x69, 0x74, 0x73, 0x42, 0x12, 0x0a, 0x10, 0x5f, 0x54, 0x6f, 0x74,
	0x61, 0x6c, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x42, 0x11, 0x0a, 0x0f,
	0x5f, 0x53, 0x65, 0x6c, 0x66, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x22,
	0xd5, 0x02, 0x0a, 0x0b, 0x41, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x12,
	0x14, 0x0a, 0x05, 0x41, 0x70, 0x70, 0x49, 0x44, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x41, 0x70, 0x70, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18,
	0x1e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x16, 0x0a,
	0x06, 0x47, 0x6f, 0x6f, 0x64, 0x49, 0x44, 0x18, 0x28, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x47,
	0x6f, 0x6f, 0x64, 0x49, 0x44, 0x12, 0x1e, 0x0a, 0x0a, 0x43, 0x6f, 0x69, 0x6e, 0x54, 0x79, 0x70,
	0x65, 0x49, 0x44, 0x18, 0x32, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x43, 0x6f, 0x69, 0x6e, 0x54,
	0x79, 0x70, 0x65, 0x49, 0x44, 0x12, 0x20, 0x0a, 0x0b, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x41, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x3c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x54, 0x6f, 0x74, 0x61,
	0x6c, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x53, 0x65, 0x6c, 0x66, 0x41,
	0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x46, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x53, 0x65, 0x6c,
	0x66, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x54, 0x6f, 0x74, 0x61, 0x6c,
	0x55, 0x6e, 0x69, 0x74, 0x73, 0x18, 0x50, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x54, 0x6f, 0x74,
	0x61, 0x6c, 0x55, 0x6e, 0x69, 0x74, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x53, 0x65, 0x6c, 0x66, 0x55,
	0x6e, 0x69, 0x74, 0x73, 0x18, 0x5a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x53, 0x65, 0x6c, 0x66,
	0x55, 0x6e, 0x69, 0x74, 0x73, 0x12, 0x28, 0x0a, 0x0f, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x43, 0x6f,
	0x6d, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x64, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f,
	0x54, 0x6f, 0x74, 0x61, 0x6c, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12,
	0x26, 0x0a, 0x0e, 0x53, 0x65, 0x6c, 0x66, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x18, 0x6e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x53, 0x65, 0x6c, 0x66, 0x43, 0x6f, 0x6d,
	0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x92, 0x03, 0x0a, 0x05, 0x43, 0x6f, 0x6e, 0x64,
	0x73, 0x12, 0x2c, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e,
	0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x72,
	0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x48, 0x00, 0x52, 0x02, 0x49, 0x44, 0x88, 0x01, 0x01, 0x12,
	0x32, 0x0a, 0x05, 0x41, 0x70, 0x70, 0x49, 0x44, 0x18, 0x14, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17,
	0x2e, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74,
	0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x48, 0x01, 0x52, 0x05, 0x41, 0x70, 0x70, 0x49, 0x44,
	0x88, 0x01, 0x01, 0x12, 0x34, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x1e, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x48, 0x02, 0x52, 0x06,
	0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x88, 0x01, 0x01, 0x12, 0x34, 0x0a, 0x06, 0x47, 0x6f, 0x6f,
	0x64, 0x49, 0x44, 0x18, 0x28, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x62, 0x61, 0x73, 0x65,
	0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56,
	0x61, 0x6c, 0x48, 0x03, 0x52, 0x06, 0x47, 0x6f, 0x6f, 0x64, 0x49, 0x44, 0x88, 0x01, 0x01, 0x12,
	0x3c, 0x0a, 0x0a, 0x43, 0x6f, 0x69, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x49, 0x44, 0x18, 0x32, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x48, 0x04, 0x52, 0x0a,
	0x43, 0x6f, 0x69, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x49, 0x44, 0x88, 0x01, 0x01, 0x12, 0x3b, 0x0a,
	0x07, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x73, 0x18, 0x78, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c,
	0x2e, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74,
	0x72, 0x69, 0x6e, 0x67, 0x53, 0x6c, 0x69, 0x63, 0x65, 0x56, 0x61, 0x6c, 0x48, 0x05, 0x52, 0x07,
	0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x73, 0x88, 0x01, 0x01, 0x42, 0x05, 0x0a, 0x03, 0x5f, 0x49,
	0x44, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x41, 0x70, 0x70, 0x49, 0x44, 0x42, 0x09, 0x0a, 0x07, 0x5f,
	0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x47, 0x6f, 0x6f, 0x64, 0x49,
	0x44, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x43, 0x6f, 0x69, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x49, 0x44,
	0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x73, 0x22, 0x39, 0x0a, 0x1d,
	0x45, 0x78, 0x70, 0x72, 0x6f, 0x70, 0x72, 0x69, 0x61, 0x74, 0x65, 0x41, 0x72, 0x63, 0x68, 0x69,
	0x76, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a,
	0x07, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x44, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x44, 0x22, 0x20, 0x0a, 0x1e, 0x45, 0x78, 0x70, 0x72, 0x6f,
	0x70, 0x72, 0x69, 0x61, 0x74, 0x65, 0x41, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x86, 0x01, 0x0a, 0x16, 0x47, 0x65,
	0x74, 0x41, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x3e, 0x0a, 0x05, 0x43, 0x6f, 0x6e, 0x64, 0x73, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x69, 0x6e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x6d, 0x69,
	0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65,
	0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x64, 0x73, 0x52, 0x05, 0x43,
	0x6f, 0x6e, 0x64, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x14,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05,
	0x4c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x4c, 0x69, 0x6d,
	0x69, 0x74, 0x22, 0x75, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x41, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65,
	0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x44, 0x0a,
	0x05, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x69,
	0x6e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72,
	0x65, 0x2e, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31,
	0x2e, 0x41, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x05, 0x49, 0x6e,
	0x66, 0x6f, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x14, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x32, 0xbb, 0x02, 0x0a, 0x0a, 0x4d, 0x69,
	0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x12, 0x9f, 0x01, 0x0a, 0x16, 0x45, 0x78, 0x70,
	0x72, 0x6f, 0x70, 0x72, 0x69, 0x61, 0x74, 0x65, 0x41, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x6d,
	0x65, 0x6e, 0x74, 0x12, 0x40, 0x2e, 0x69, 0x6e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x6d, 0x69,
	0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65,
	0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x78, 0x70, 0x72, 0x6f, 0x70, 0x72, 0x69,
	0x61, 0x74, 0x65, 0x41, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x41, 0x2e, 0x69, 0x6e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e,
	0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x61, 0x72, 0x63, 0x68, 0x69,
	0x76, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x78, 0x70, 0x72, 0x6f, 0x70,
	0x72, 0x69, 0x61, 0x74, 0x65, 0x41, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x6d, 0x65, 0x6e, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x8a, 0x01, 0x0a, 0x0f, 0x47,
	0x65, 0x74, 0x41, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x39,
	0x2e, 0x69, 0x6e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77,
	0x61, 0x72, 0x65, 0x2e, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e,
	0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3a, 0x2e, 0x69, 0x6e, 0x73, 0x70,
	0x69, 0x72, 0x65, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x61,
	0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65,
	0x74, 0x41, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x42, 0x5a, 0x40, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4e, 0x70, 0x6f, 0x6f, 0x6c, 0x50, 0x6c, 0x61, 0x74, 0x66,
	0x6f, 0x72, 0x6d, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x6e, 0x70, 0x6f, 0x6f,
	0x6c, 0x2f, 0x69, 0x6e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2f, 0x6d, 0x77, 0x2f, 0x76, 0x31, 0x2f,
	0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_npool_inspire_mw_v1_archivement_archivement_proto_rawDescOnce sync.Once
	file_npool_inspire_mw_v1_archivement_archivement_proto_rawDescData = file_npool_inspire_mw_v1_archivement_archivement_proto_rawDesc
)

func file_npool_inspire_mw_v1_archivement_archivement_proto_rawDescGZIP() []byte {
	file_npool_inspire_mw_v1_archivement_archivement_proto_rawDescOnce.Do(func() {
		file_npool_inspire_mw_v1_archivement_archivement_proto_rawDescData = protoimpl.X.CompressGZIP(file_npool_inspire_mw_v1_archivement_archivement_proto_rawDescData)
	})
	return file_npool_inspire_mw_v1_archivement_archivement_proto_rawDescData
}

var file_npool_inspire_mw_v1_archivement_archivement_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_npool_inspire_mw_v1_archivement_archivement_proto_goTypes = []interface{}{
	(*ArchivementReq)(nil),                 // 0: inspire.middleware.archivement.v1.ArchivementReq
	(*Archivement)(nil),                    // 1: inspire.middleware.archivement.v1.Archivement
	(*Conds)(nil),                          // 2: inspire.middleware.archivement.v1.Conds
	(*ExpropriateArchivementRequest)(nil),  // 3: inspire.middleware.archivement.v1.ExpropriateArchivementRequest
	(*ExpropriateArchivementResponse)(nil), // 4: inspire.middleware.archivement.v1.ExpropriateArchivementResponse
	(*GetArchivementsRequest)(nil),         // 5: inspire.middleware.archivement.v1.GetArchivementsRequest
	(*GetArchivementsResponse)(nil),        // 6: inspire.middleware.archivement.v1.GetArchivementsResponse
	(*v1.StringVal)(nil),                   // 7: basetypes.v1.StringVal
	(*v1.StringSliceVal)(nil),              // 8: basetypes.v1.StringSliceVal
}
var file_npool_inspire_mw_v1_archivement_archivement_proto_depIdxs = []int32{
	7,  // 0: inspire.middleware.archivement.v1.Conds.ID:type_name -> basetypes.v1.StringVal
	7,  // 1: inspire.middleware.archivement.v1.Conds.AppID:type_name -> basetypes.v1.StringVal
	7,  // 2: inspire.middleware.archivement.v1.Conds.UserID:type_name -> basetypes.v1.StringVal
	7,  // 3: inspire.middleware.archivement.v1.Conds.GoodID:type_name -> basetypes.v1.StringVal
	7,  // 4: inspire.middleware.archivement.v1.Conds.CoinTypeID:type_name -> basetypes.v1.StringVal
	8,  // 5: inspire.middleware.archivement.v1.Conds.UserIDs:type_name -> basetypes.v1.StringSliceVal
	2,  // 6: inspire.middleware.archivement.v1.GetArchivementsRequest.Conds:type_name -> inspire.middleware.archivement.v1.Conds
	1,  // 7: inspire.middleware.archivement.v1.GetArchivementsResponse.Infos:type_name -> inspire.middleware.archivement.v1.Archivement
	3,  // 8: inspire.middleware.archivement.v1.Middleware.ExpropriateArchivement:input_type -> inspire.middleware.archivement.v1.ExpropriateArchivementRequest
	5,  // 9: inspire.middleware.archivement.v1.Middleware.GetArchivements:input_type -> inspire.middleware.archivement.v1.GetArchivementsRequest
	4,  // 10: inspire.middleware.archivement.v1.Middleware.ExpropriateArchivement:output_type -> inspire.middleware.archivement.v1.ExpropriateArchivementResponse
	6,  // 11: inspire.middleware.archivement.v1.Middleware.GetArchivements:output_type -> inspire.middleware.archivement.v1.GetArchivementsResponse
	10, // [10:12] is the sub-list for method output_type
	8,  // [8:10] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_npool_inspire_mw_v1_archivement_archivement_proto_init() }
func file_npool_inspire_mw_v1_archivement_archivement_proto_init() {
	if File_npool_inspire_mw_v1_archivement_archivement_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_npool_inspire_mw_v1_archivement_archivement_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ArchivementReq); i {
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
		file_npool_inspire_mw_v1_archivement_archivement_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Archivement); i {
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
		file_npool_inspire_mw_v1_archivement_archivement_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
		file_npool_inspire_mw_v1_archivement_archivement_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExpropriateArchivementRequest); i {
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
		file_npool_inspire_mw_v1_archivement_archivement_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExpropriateArchivementResponse); i {
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
		file_npool_inspire_mw_v1_archivement_archivement_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetArchivementsRequest); i {
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
		file_npool_inspire_mw_v1_archivement_archivement_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetArchivementsResponse); i {
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
	file_npool_inspire_mw_v1_archivement_archivement_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_npool_inspire_mw_v1_archivement_archivement_proto_msgTypes[2].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_npool_inspire_mw_v1_archivement_archivement_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_npool_inspire_mw_v1_archivement_archivement_proto_goTypes,
		DependencyIndexes: file_npool_inspire_mw_v1_archivement_archivement_proto_depIdxs,
		MessageInfos:      file_npool_inspire_mw_v1_archivement_archivement_proto_msgTypes,
	}.Build()
	File_npool_inspire_mw_v1_archivement_archivement_proto = out.File
	file_npool_inspire_mw_v1_archivement_archivement_proto_rawDesc = nil
	file_npool_inspire_mw_v1_archivement_archivement_proto_goTypes = nil
	file_npool_inspire_mw_v1_archivement_archivement_proto_depIdxs = nil
}
