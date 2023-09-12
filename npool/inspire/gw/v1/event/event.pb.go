// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.18.1
// source: npool/inspire/gw/v1/event/event.proto

package event

import (
	v1 "github.com/NpoolPlatform/message/npool/basetypes/v1"
	coupon "github.com/NpoolPlatform/message/npool/inspire/mw/v1/coupon"
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

type Event struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID             string           `protobuf:"bytes,10,opt,name=ID,proto3" json:"ID,omitempty"`
	AppID          string           `protobuf:"bytes,20,opt,name=AppID,proto3" json:"AppID,omitempty"`
	AppName        string           `protobuf:"bytes,30,opt,name=AppName,proto3" json:"AppName,omitempty"`
	EventType      v1.UsedFor       `protobuf:"varint,40,opt,name=EventType,proto3,enum=basetypes.v1.UsedFor" json:"EventType,omitempty"`
	Coupons        []*coupon.Coupon `protobuf:"bytes,50,rep,name=Coupons,proto3" json:"Coupons,omitempty"`
	Credits        string           `protobuf:"bytes,60,opt,name=Credits,proto3" json:"Credits,omitempty"`
	CreditsPerUSD  string           `protobuf:"bytes,70,opt,name=CreditsPerUSD,proto3" json:"CreditsPerUSD,omitempty"`
	MaxConsecutive uint32           `protobuf:"varint,80,opt,name=MaxConsecutive,proto3" json:"MaxConsecutive,omitempty"`
	GoodID         string           `protobuf:"bytes,90,opt,name=GoodID,proto3" json:"GoodID,omitempty"`
	GoodName       string           `protobuf:"bytes,100,opt,name=GoodName,proto3" json:"GoodName,omitempty"`
	InviterLayers  uint32           `protobuf:"varint,110,opt,name=InviterLayers,proto3" json:"InviterLayers,omitempty"`
	CreatedAt      uint32           `protobuf:"varint,200,opt,name=CreatedAt,proto3" json:"CreatedAt,omitempty"`
	UpdatedAt      uint32           `protobuf:"varint,210,opt,name=UpdatedAt,proto3" json:"UpdatedAt,omitempty"`
	DeletedAt      uint32           `protobuf:"varint,220,opt,name=DeletedAt,proto3" json:"DeletedAt,omitempty"`
	AppGoodID      string           `protobuf:"bytes,230,opt,name=AppGoodID,proto3" json:"AppGoodID,omitempty"`
	AppGoodName    string           `protobuf:"bytes,240,opt,name=AppGoodName,proto3" json:"AppGoodName,omitempty"`
}

func (x *Event) Reset() {
	*x = Event{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_inspire_gw_v1_event_event_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Event) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Event) ProtoMessage() {}

func (x *Event) ProtoReflect() protoreflect.Message {
	mi := &file_npool_inspire_gw_v1_event_event_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Event.ProtoReflect.Descriptor instead.
func (*Event) Descriptor() ([]byte, []int) {
	return file_npool_inspire_gw_v1_event_event_proto_rawDescGZIP(), []int{0}
}

func (x *Event) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *Event) GetAppID() string {
	if x != nil {
		return x.AppID
	}
	return ""
}

func (x *Event) GetAppName() string {
	if x != nil {
		return x.AppName
	}
	return ""
}

func (x *Event) GetEventType() v1.UsedFor {
	if x != nil {
		return x.EventType
	}
	return v1.UsedFor(0)
}

func (x *Event) GetCoupons() []*coupon.Coupon {
	if x != nil {
		return x.Coupons
	}
	return nil
}

func (x *Event) GetCredits() string {
	if x != nil {
		return x.Credits
	}
	return ""
}

func (x *Event) GetCreditsPerUSD() string {
	if x != nil {
		return x.CreditsPerUSD
	}
	return ""
}

func (x *Event) GetMaxConsecutive() uint32 {
	if x != nil {
		return x.MaxConsecutive
	}
	return 0
}

func (x *Event) GetGoodID() string {
	if x != nil {
		return x.GoodID
	}
	return ""
}

func (x *Event) GetGoodName() string {
	if x != nil {
		return x.GoodName
	}
	return ""
}

func (x *Event) GetInviterLayers() uint32 {
	if x != nil {
		return x.InviterLayers
	}
	return 0
}

func (x *Event) GetCreatedAt() uint32 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *Event) GetUpdatedAt() uint32 {
	if x != nil {
		return x.UpdatedAt
	}
	return 0
}

func (x *Event) GetDeletedAt() uint32 {
	if x != nil {
		return x.DeletedAt
	}
	return 0
}

func (x *Event) GetAppGoodID() string {
	if x != nil {
		return x.AppGoodID
	}
	return ""
}

func (x *Event) GetAppGoodName() string {
	if x != nil {
		return x.AppGoodName
	}
	return ""
}

type CreateEventRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AppID          string     `protobuf:"bytes,10,opt,name=AppID,proto3" json:"AppID,omitempty"`
	EventType      v1.UsedFor `protobuf:"varint,20,opt,name=EventType,proto3,enum=basetypes.v1.UsedFor" json:"EventType,omitempty"`
	CouponIDs      []string   `protobuf:"bytes,30,rep,name=CouponIDs,proto3" json:"CouponIDs,omitempty"`
	Credits        *string    `protobuf:"bytes,40,opt,name=Credits,proto3,oneof" json:"Credits,omitempty"`
	CreditsPerUSD  *string    `protobuf:"bytes,50,opt,name=CreditsPerUSD,proto3,oneof" json:"CreditsPerUSD,omitempty"`
	MaxConsecutive *uint32    `protobuf:"varint,60,opt,name=MaxConsecutive,proto3,oneof" json:"MaxConsecutive,omitempty"`
	AppGoodID      *string    `protobuf:"bytes,70,opt,name=AppGoodID,proto3,oneof" json:"AppGoodID,omitempty"`
	InviterLayers  *uint32    `protobuf:"varint,80,opt,name=InviterLayers,proto3,oneof" json:"InviterLayers,omitempty"`
}

func (x *CreateEventRequest) Reset() {
	*x = CreateEventRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_inspire_gw_v1_event_event_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateEventRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateEventRequest) ProtoMessage() {}

func (x *CreateEventRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_inspire_gw_v1_event_event_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateEventRequest.ProtoReflect.Descriptor instead.
func (*CreateEventRequest) Descriptor() ([]byte, []int) {
	return file_npool_inspire_gw_v1_event_event_proto_rawDescGZIP(), []int{1}
}

func (x *CreateEventRequest) GetAppID() string {
	if x != nil {
		return x.AppID
	}
	return ""
}

func (x *CreateEventRequest) GetEventType() v1.UsedFor {
	if x != nil {
		return x.EventType
	}
	return v1.UsedFor(0)
}

func (x *CreateEventRequest) GetCouponIDs() []string {
	if x != nil {
		return x.CouponIDs
	}
	return nil
}

func (x *CreateEventRequest) GetCredits() string {
	if x != nil && x.Credits != nil {
		return *x.Credits
	}
	return ""
}

func (x *CreateEventRequest) GetCreditsPerUSD() string {
	if x != nil && x.CreditsPerUSD != nil {
		return *x.CreditsPerUSD
	}
	return ""
}

func (x *CreateEventRequest) GetMaxConsecutive() uint32 {
	if x != nil && x.MaxConsecutive != nil {
		return *x.MaxConsecutive
	}
	return 0
}

func (x *CreateEventRequest) GetAppGoodID() string {
	if x != nil && x.AppGoodID != nil {
		return *x.AppGoodID
	}
	return ""
}

func (x *CreateEventRequest) GetInviterLayers() uint32 {
	if x != nil && x.InviterLayers != nil {
		return *x.InviterLayers
	}
	return 0
}

type CreateEventResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *Event `protobuf:"bytes,10,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (x *CreateEventResponse) Reset() {
	*x = CreateEventResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_inspire_gw_v1_event_event_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateEventResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateEventResponse) ProtoMessage() {}

func (x *CreateEventResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_inspire_gw_v1_event_event_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateEventResponse.ProtoReflect.Descriptor instead.
func (*CreateEventResponse) Descriptor() ([]byte, []int) {
	return file_npool_inspire_gw_v1_event_event_proto_rawDescGZIP(), []int{2}
}

func (x *CreateEventResponse) GetInfo() *Event {
	if x != nil {
		return x.Info
	}
	return nil
}

type GetEventsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AppID  string `protobuf:"bytes,10,opt,name=AppID,proto3" json:"AppID,omitempty"`
	Offset int32  `protobuf:"varint,20,opt,name=Offset,proto3" json:"Offset,omitempty"`
	Limit  int32  `protobuf:"varint,30,opt,name=Limit,proto3" json:"Limit,omitempty"`
}

func (x *GetEventsRequest) Reset() {
	*x = GetEventsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_inspire_gw_v1_event_event_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetEventsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetEventsRequest) ProtoMessage() {}

func (x *GetEventsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_inspire_gw_v1_event_event_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetEventsRequest.ProtoReflect.Descriptor instead.
func (*GetEventsRequest) Descriptor() ([]byte, []int) {
	return file_npool_inspire_gw_v1_event_event_proto_rawDescGZIP(), []int{3}
}

func (x *GetEventsRequest) GetAppID() string {
	if x != nil {
		return x.AppID
	}
	return ""
}

func (x *GetEventsRequest) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *GetEventsRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type GetEventsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Infos []*Event `protobuf:"bytes,10,rep,name=Infos,proto3" json:"Infos,omitempty"`
	Total uint32   `protobuf:"varint,20,opt,name=Total,proto3" json:"Total,omitempty"`
}

func (x *GetEventsResponse) Reset() {
	*x = GetEventsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_inspire_gw_v1_event_event_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetEventsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetEventsResponse) ProtoMessage() {}

func (x *GetEventsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_inspire_gw_v1_event_event_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetEventsResponse.ProtoReflect.Descriptor instead.
func (*GetEventsResponse) Descriptor() ([]byte, []int) {
	return file_npool_inspire_gw_v1_event_event_proto_rawDescGZIP(), []int{4}
}

func (x *GetEventsResponse) GetInfos() []*Event {
	if x != nil {
		return x.Infos
	}
	return nil
}

func (x *GetEventsResponse) GetTotal() uint32 {
	if x != nil {
		return x.Total
	}
	return 0
}

type UpdateEventRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID             string   `protobuf:"bytes,10,opt,name=ID,proto3" json:"ID,omitempty"`
	AppID          string   `protobuf:"bytes,20,opt,name=AppID,proto3" json:"AppID,omitempty"`
	CouponIDs      []string `protobuf:"bytes,30,rep,name=CouponIDs,proto3" json:"CouponIDs,omitempty"`
	Credits        *string  `protobuf:"bytes,40,opt,name=Credits,proto3,oneof" json:"Credits,omitempty"`
	CreditsPerUSD  *string  `protobuf:"bytes,50,opt,name=CreditsPerUSD,proto3,oneof" json:"CreditsPerUSD,omitempty"`
	MaxConsecutive *uint32  `protobuf:"varint,60,opt,name=MaxConsecutive,proto3,oneof" json:"MaxConsecutive,omitempty"`
	InviterLayers  *uint32  `protobuf:"varint,80,opt,name=InviterLayers,proto3,oneof" json:"InviterLayers,omitempty"`
}

func (x *UpdateEventRequest) Reset() {
	*x = UpdateEventRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_inspire_gw_v1_event_event_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateEventRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateEventRequest) ProtoMessage() {}

func (x *UpdateEventRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_inspire_gw_v1_event_event_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateEventRequest.ProtoReflect.Descriptor instead.
func (*UpdateEventRequest) Descriptor() ([]byte, []int) {
	return file_npool_inspire_gw_v1_event_event_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateEventRequest) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *UpdateEventRequest) GetAppID() string {
	if x != nil {
		return x.AppID
	}
	return ""
}

func (x *UpdateEventRequest) GetCouponIDs() []string {
	if x != nil {
		return x.CouponIDs
	}
	return nil
}

func (x *UpdateEventRequest) GetCredits() string {
	if x != nil && x.Credits != nil {
		return *x.Credits
	}
	return ""
}

func (x *UpdateEventRequest) GetCreditsPerUSD() string {
	if x != nil && x.CreditsPerUSD != nil {
		return *x.CreditsPerUSD
	}
	return ""
}

func (x *UpdateEventRequest) GetMaxConsecutive() uint32 {
	if x != nil && x.MaxConsecutive != nil {
		return *x.MaxConsecutive
	}
	return 0
}

func (x *UpdateEventRequest) GetInviterLayers() uint32 {
	if x != nil && x.InviterLayers != nil {
		return *x.InviterLayers
	}
	return 0
}

type UpdateEventResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *Event `protobuf:"bytes,10,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (x *UpdateEventResponse) Reset() {
	*x = UpdateEventResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_inspire_gw_v1_event_event_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateEventResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateEventResponse) ProtoMessage() {}

func (x *UpdateEventResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_inspire_gw_v1_event_event_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateEventResponse.ProtoReflect.Descriptor instead.
func (*UpdateEventResponse) Descriptor() ([]byte, []int) {
	return file_npool_inspire_gw_v1_event_event_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateEventResponse) GetInfo() *Event {
	if x != nil {
		return x.Info
	}
	return nil
}

var File_npool_inspire_gw_v1_event_event_proto protoreflect.FileDescriptor

var file_npool_inspire_gw_v1_event_event_proto_rawDesc = []byte{
	0x0a, 0x25, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x69, 0x6e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2f,
	0x67, 0x77, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2f, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18, 0x69, 0x6e, 0x73, 0x70, 0x69, 0x72, 0x65,
	0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x76,
	0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x27, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x69, 0x6e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2f, 0x6d,
	0x77, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x75, 0x70,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f,
	0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65,
	0x64, 0x66, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9d, 0x04, 0x0a, 0x05, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x41, 0x70, 0x70, 0x49, 0x44, 0x18, 0x14, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x41, 0x70, 0x70, 0x49, 0x44, 0x12, 0x18, 0x0a, 0x07, 0x41, 0x70,
	0x70, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x41, 0x70, 0x70,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x33, 0x0a, 0x09, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70,
	0x65, 0x18, 0x28, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x15, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x64, 0x46, 0x6f, 0x72, 0x52, 0x09,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x3e, 0x0a, 0x07, 0x43, 0x6f, 0x75,
	0x70, 0x6f, 0x6e, 0x73, 0x18, 0x32, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x69, 0x6e, 0x73,
	0x70, 0x69, 0x72, 0x65, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e,
	0x63, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x75, 0x70, 0x6f, 0x6e,
	0x52, 0x07, 0x43, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x43, 0x72, 0x65,
	0x64, 0x69, 0x74, 0x73, 0x18, 0x3c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x43, 0x72, 0x65, 0x64,
	0x69, 0x74, 0x73, 0x12, 0x24, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x64, 0x69, 0x74, 0x73, 0x50, 0x65,
	0x72, 0x55, 0x53, 0x44, 0x18, 0x46, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x43, 0x72, 0x65, 0x64,
	0x69, 0x74, 0x73, 0x50, 0x65, 0x72, 0x55, 0x53, 0x44, 0x12, 0x26, 0x0a, 0x0e, 0x4d, 0x61, 0x78,
	0x43, 0x6f, 0x6e, 0x73, 0x65, 0x63, 0x75, 0x74, 0x69, 0x76, 0x65, 0x18, 0x50, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x0e, 0x4d, 0x61, 0x78, 0x43, 0x6f, 0x6e, 0x73, 0x65, 0x63, 0x75, 0x74, 0x69, 0x76,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x47, 0x6f, 0x6f, 0x64, 0x49, 0x44, 0x18, 0x5a, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x47, 0x6f, 0x6f, 0x64, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08, 0x47, 0x6f, 0x6f,
	0x64, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x64, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x47, 0x6f, 0x6f,
	0x64, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x72,
	0x4c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x18, 0x6e, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0d, 0x49, 0x6e,
	0x76, 0x69, 0x74, 0x65, 0x72, 0x4c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x12, 0x1d, 0x0a, 0x09, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0xc8, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x09, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0xd2, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x09, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0xdc, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x09, 0x41, 0x70, 0x70, 0x47,
	0x6f, 0x6f, 0x64, 0x49, 0x44, 0x18, 0xe6, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x41, 0x70,
	0x70, 0x47, 0x6f, 0x6f, 0x64, 0x49, 0x44, 0x12, 0x21, 0x0a, 0x0b, 0x41, 0x70, 0x70, 0x47, 0x6f,
	0x6f, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0xf0, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x41,
	0x70, 0x70, 0x47, 0x6f, 0x6f, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x93, 0x03, 0x0a, 0x12, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x14, 0x0a, 0x05, 0x41, 0x70, 0x70, 0x49, 0x44, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x41, 0x70, 0x70, 0x49, 0x44, 0x12, 0x33, 0x0a, 0x09, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x54, 0x79, 0x70, 0x65, 0x18, 0x14, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x15, 0x2e, 0x62, 0x61, 0x73,
	0x65, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x64, 0x46, 0x6f,
	0x72, 0x52, 0x09, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1c, 0x0a, 0x09,
	0x43, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x49, 0x44, 0x73, 0x18, 0x1e, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x09, 0x43, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x49, 0x44, 0x73, 0x12, 0x1d, 0x0a, 0x07, 0x43, 0x72,
	0x65, 0x64, 0x69, 0x74, 0x73, 0x18, 0x28, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x07, 0x43,
	0x72, 0x65, 0x64, 0x69, 0x74, 0x73, 0x88, 0x01, 0x01, 0x12, 0x29, 0x0a, 0x0d, 0x43, 0x72, 0x65,
	0x64, 0x69, 0x74, 0x73, 0x50, 0x65, 0x72, 0x55, 0x53, 0x44, 0x18, 0x32, 0x20, 0x01, 0x28, 0x09,
	0x48, 0x01, 0x52, 0x0d, 0x43, 0x72, 0x65, 0x64, 0x69, 0x74, 0x73, 0x50, 0x65, 0x72, 0x55, 0x53,
	0x44, 0x88, 0x01, 0x01, 0x12, 0x2b, 0x0a, 0x0e, 0x4d, 0x61, 0x78, 0x43, 0x6f, 0x6e, 0x73, 0x65,
	0x63, 0x75, 0x74, 0x69, 0x76, 0x65, 0x18, 0x3c, 0x20, 0x01, 0x28, 0x0d, 0x48, 0x02, 0x52, 0x0e,
	0x4d, 0x61, 0x78, 0x43, 0x6f, 0x6e, 0x73, 0x65, 0x63, 0x75, 0x74, 0x69, 0x76, 0x65, 0x88, 0x01,
	0x01, 0x12, 0x21, 0x0a, 0x09, 0x41, 0x70, 0x70, 0x47, 0x6f, 0x6f, 0x64, 0x49, 0x44, 0x18, 0x46,
	0x20, 0x01, 0x28, 0x09, 0x48, 0x03, 0x52, 0x09, 0x41, 0x70, 0x70, 0x47, 0x6f, 0x6f, 0x64, 0x49,
	0x44, 0x88, 0x01, 0x01, 0x12, 0x29, 0x0a, 0x0d, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x72, 0x4c,
	0x61, 0x79, 0x65, 0x72, 0x73, 0x18, 0x50, 0x20, 0x01, 0x28, 0x0d, 0x48, 0x04, 0x52, 0x0d, 0x49,
	0x6e, 0x76, 0x69, 0x74, 0x65, 0x72, 0x4c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x88, 0x01, 0x01, 0x42,
	0x0a, 0x0a, 0x08, 0x5f, 0x43, 0x72, 0x65, 0x64, 0x69, 0x74, 0x73, 0x42, 0x10, 0x0a, 0x0e, 0x5f,
	0x43, 0x72, 0x65, 0x64, 0x69, 0x74, 0x73, 0x50, 0x65, 0x72, 0x55, 0x53, 0x44, 0x42, 0x11, 0x0a,
	0x0f, 0x5f, 0x4d, 0x61, 0x78, 0x43, 0x6f, 0x6e, 0x73, 0x65, 0x63, 0x75, 0x74, 0x69, 0x76, 0x65,
	0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x41, 0x70, 0x70, 0x47, 0x6f, 0x6f, 0x64, 0x49, 0x44, 0x42, 0x10,
	0x0a, 0x0e, 0x5f, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x72, 0x4c, 0x61, 0x79, 0x65, 0x72, 0x73,
	0x22, 0x4a, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x33, 0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x69, 0x6e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e,
	0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31,
	0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x56, 0x0a, 0x10,
	0x47, 0x65, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x14, 0x0a, 0x05, 0x41, 0x70, 0x70, 0x49, 0x44, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x41, 0x70, 0x70, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74,
	0x18, 0x14, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x14,
	0x0a, 0x05, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x4c,
	0x69, 0x6d, 0x69, 0x74, 0x22, 0x60, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x35, 0x0a, 0x05, 0x49, 0x6e, 0x66,
	0x6f, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x69, 0x6e, 0x73, 0x70, 0x69,
	0x72, 0x65, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x2e, 0x76, 0x31, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x05, 0x49, 0x6e, 0x66, 0x6f, 0x73,
	0x12, 0x14, 0x0a, 0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x14, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x22, 0xbd, 0x02, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x49, 0x44, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x12, 0x14, 0x0a,
	0x05, 0x41, 0x70, 0x70, 0x49, 0x44, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x41, 0x70,
	0x70, 0x49, 0x44, 0x12, 0x1c, 0x0a, 0x09, 0x43, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x49, 0x44, 0x73,
	0x18, 0x1e, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x43, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x49, 0x44,
	0x73, 0x12, 0x1d, 0x0a, 0x07, 0x43, 0x72, 0x65, 0x64, 0x69, 0x74, 0x73, 0x18, 0x28, 0x20, 0x01,
	0x28, 0x09, 0x48, 0x00, 0x52, 0x07, 0x43, 0x72, 0x65, 0x64, 0x69, 0x74, 0x73, 0x88, 0x01, 0x01,
	0x12, 0x29, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x64, 0x69, 0x74, 0x73, 0x50, 0x65, 0x72, 0x55, 0x53,
	0x44, 0x18, 0x32, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x0d, 0x43, 0x72, 0x65, 0x64, 0x69,
	0x74, 0x73, 0x50, 0x65, 0x72, 0x55, 0x53, 0x44, 0x88, 0x01, 0x01, 0x12, 0x2b, 0x0a, 0x0e, 0x4d,
	0x61, 0x78, 0x43, 0x6f, 0x6e, 0x73, 0x65, 0x63, 0x75, 0x74, 0x69, 0x76, 0x65, 0x18, 0x3c, 0x20,
	0x01, 0x28, 0x0d, 0x48, 0x02, 0x52, 0x0e, 0x4d, 0x61, 0x78, 0x43, 0x6f, 0x6e, 0x73, 0x65, 0x63,
	0x75, 0x74, 0x69, 0x76, 0x65, 0x88, 0x01, 0x01, 0x12, 0x29, 0x0a, 0x0d, 0x49, 0x6e, 0x76, 0x69,
	0x74, 0x65, 0x72, 0x4c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x18, 0x50, 0x20, 0x01, 0x28, 0x0d, 0x48,
	0x03, 0x52, 0x0d, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x72, 0x4c, 0x61, 0x79, 0x65, 0x72, 0x73,
	0x88, 0x01, 0x01, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x43, 0x72, 0x65, 0x64, 0x69, 0x74, 0x73, 0x42,
	0x10, 0x0a, 0x0e, 0x5f, 0x43, 0x72, 0x65, 0x64, 0x69, 0x74, 0x73, 0x50, 0x65, 0x72, 0x55, 0x53,
	0x44, 0x42, 0x11, 0x0a, 0x0f, 0x5f, 0x4d, 0x61, 0x78, 0x43, 0x6f, 0x6e, 0x73, 0x65, 0x63, 0x75,
	0x74, 0x69, 0x76, 0x65, 0x42, 0x10, 0x0a, 0x0e, 0x5f, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x72,
	0x4c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x22, 0x4a, 0x0a, 0x13, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x33, 0x0a,
	0x04, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x69, 0x6e,
	0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x04, 0x49, 0x6e,
	0x66, 0x6f, 0x32, 0xb4, 0x03, 0x0a, 0x07, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x12, 0x8e,
	0x01, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x2c,
	0x2e, 0x69, 0x6e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79,
	0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2d, 0x2e, 0x69,
	0x6e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x22, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x1c, 0x3a, 0x01, 0x2a, 0x22, 0x17, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x69, 0x6e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x12,
	0x86, 0x01, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x2a, 0x2e,
	0x69, 0x6e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2b, 0x2e, 0x69, 0x6e, 0x73, 0x70,
	0x69, 0x72, 0x65, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x20, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1a, 0x3a, 0x01,
	0x2a, 0x22, 0x15, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x65, 0x74, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x69, 0x6e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x73, 0x12, 0x8e, 0x01, 0x0a, 0x0b, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x2c, 0x2e, 0x69, 0x6e, 0x73, 0x70, 0x69,
	0x72, 0x65, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2d, 0x2e, 0x69, 0x6e, 0x73, 0x70, 0x69, 0x72, 0x65,
	0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x76,
	0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x22, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1c, 0x3a, 0x01, 0x2a,
	0x22, 0x17, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x69, 0x6e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x42, 0x3c, 0x5a, 0x3a, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4e, 0x70, 0x6f, 0x6f, 0x6c, 0x50, 0x6c, 0x61,
	0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x6e, 0x70,
	0x6f, 0x6f, 0x6c, 0x2f, 0x69, 0x6e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2f, 0x67, 0x77, 0x2f, 0x76,
	0x31, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_npool_inspire_gw_v1_event_event_proto_rawDescOnce sync.Once
	file_npool_inspire_gw_v1_event_event_proto_rawDescData = file_npool_inspire_gw_v1_event_event_proto_rawDesc
)

func file_npool_inspire_gw_v1_event_event_proto_rawDescGZIP() []byte {
	file_npool_inspire_gw_v1_event_event_proto_rawDescOnce.Do(func() {
		file_npool_inspire_gw_v1_event_event_proto_rawDescData = protoimpl.X.CompressGZIP(file_npool_inspire_gw_v1_event_event_proto_rawDescData)
	})
	return file_npool_inspire_gw_v1_event_event_proto_rawDescData
}

var file_npool_inspire_gw_v1_event_event_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_npool_inspire_gw_v1_event_event_proto_goTypes = []interface{}{
	(*Event)(nil),               // 0: inspire.gateway.event.v1.Event
	(*CreateEventRequest)(nil),  // 1: inspire.gateway.event.v1.CreateEventRequest
	(*CreateEventResponse)(nil), // 2: inspire.gateway.event.v1.CreateEventResponse
	(*GetEventsRequest)(nil),    // 3: inspire.gateway.event.v1.GetEventsRequest
	(*GetEventsResponse)(nil),   // 4: inspire.gateway.event.v1.GetEventsResponse
	(*UpdateEventRequest)(nil),  // 5: inspire.gateway.event.v1.UpdateEventRequest
	(*UpdateEventResponse)(nil), // 6: inspire.gateway.event.v1.UpdateEventResponse
	(v1.UsedFor)(0),             // 7: basetypes.v1.UsedFor
	(*coupon.Coupon)(nil),       // 8: inspire.middleware.coupon.v1.Coupon
}
var file_npool_inspire_gw_v1_event_event_proto_depIdxs = []int32{
	7, // 0: inspire.gateway.event.v1.Event.EventType:type_name -> basetypes.v1.UsedFor
	8, // 1: inspire.gateway.event.v1.Event.Coupons:type_name -> inspire.middleware.coupon.v1.Coupon
	7, // 2: inspire.gateway.event.v1.CreateEventRequest.EventType:type_name -> basetypes.v1.UsedFor
	0, // 3: inspire.gateway.event.v1.CreateEventResponse.Info:type_name -> inspire.gateway.event.v1.Event
	0, // 4: inspire.gateway.event.v1.GetEventsResponse.Infos:type_name -> inspire.gateway.event.v1.Event
	0, // 5: inspire.gateway.event.v1.UpdateEventResponse.Info:type_name -> inspire.gateway.event.v1.Event
	1, // 6: inspire.gateway.event.v1.Gateway.CreateEvent:input_type -> inspire.gateway.event.v1.CreateEventRequest
	3, // 7: inspire.gateway.event.v1.Gateway.GetEvents:input_type -> inspire.gateway.event.v1.GetEventsRequest
	5, // 8: inspire.gateway.event.v1.Gateway.UpdateEvent:input_type -> inspire.gateway.event.v1.UpdateEventRequest
	2, // 9: inspire.gateway.event.v1.Gateway.CreateEvent:output_type -> inspire.gateway.event.v1.CreateEventResponse
	4, // 10: inspire.gateway.event.v1.Gateway.GetEvents:output_type -> inspire.gateway.event.v1.GetEventsResponse
	6, // 11: inspire.gateway.event.v1.Gateway.UpdateEvent:output_type -> inspire.gateway.event.v1.UpdateEventResponse
	9, // [9:12] is the sub-list for method output_type
	6, // [6:9] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_npool_inspire_gw_v1_event_event_proto_init() }
func file_npool_inspire_gw_v1_event_event_proto_init() {
	if File_npool_inspire_gw_v1_event_event_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_npool_inspire_gw_v1_event_event_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Event); i {
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
		file_npool_inspire_gw_v1_event_event_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateEventRequest); i {
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
		file_npool_inspire_gw_v1_event_event_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateEventResponse); i {
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
		file_npool_inspire_gw_v1_event_event_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetEventsRequest); i {
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
		file_npool_inspire_gw_v1_event_event_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetEventsResponse); i {
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
		file_npool_inspire_gw_v1_event_event_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateEventRequest); i {
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
		file_npool_inspire_gw_v1_event_event_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateEventResponse); i {
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
	file_npool_inspire_gw_v1_event_event_proto_msgTypes[1].OneofWrappers = []interface{}{}
	file_npool_inspire_gw_v1_event_event_proto_msgTypes[5].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_npool_inspire_gw_v1_event_event_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_npool_inspire_gw_v1_event_event_proto_goTypes,
		DependencyIndexes: file_npool_inspire_gw_v1_event_event_proto_depIdxs,
		MessageInfos:      file_npool_inspire_gw_v1_event_event_proto_msgTypes,
	}.Build()
	File_npool_inspire_gw_v1_event_event_proto = out.File
	file_npool_inspire_gw_v1_event_event_proto_rawDesc = nil
	file_npool_inspire_gw_v1_event_event_proto_goTypes = nil
	file_npool_inspire_gw_v1_event_event_proto_depIdxs = nil
}
