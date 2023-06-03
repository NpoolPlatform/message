// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.18.1
// source: npool/notif/gw/v1/notif/notif.proto

package notif

import (
	v1 "github.com/NpoolPlatform/message/npool/basetypes/v1"
	channel "github.com/NpoolPlatform/message/npool/notif/mgr/v1/channel"
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

type Notif struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID           string               `protobuf:"bytes,10,opt,name=ID,proto3" json:"ID,omitempty"`
	AppID        string               `protobuf:"bytes,20,opt,name=AppID,proto3" json:"AppID,omitempty"`
	AppName      string               `protobuf:"bytes,30,opt,name=AppName,proto3" json:"AppName,omitempty"`
	UserID       string               `protobuf:"bytes,40,opt,name=UserID,proto3" json:"UserID,omitempty"`
	EmailAddress string               `protobuf:"bytes,50,opt,name=EmailAddress,proto3" json:"EmailAddress,omitempty"`
	PhoneNO      string               `protobuf:"bytes,60,opt,name=PhoneNO,proto3" json:"PhoneNO,omitempty"`
	Username     string               `protobuf:"bytes,70,opt,name=Username,proto3" json:"Username,omitempty"`
	EventType    v1.UsedFor           `protobuf:"varint,80,opt,name=EventType,proto3,enum=basetypes.v1.UsedFor" json:"EventType,omitempty"`
	UseTemplate  bool                 `protobuf:"varint,90,opt,name=UseTemplate,proto3" json:"UseTemplate,omitempty"`
	Title        string               `protobuf:"bytes,100,opt,name=Title,proto3" json:"Title,omitempty"`
	Content      string               `protobuf:"bytes,110,opt,name=Content,proto3" json:"Content,omitempty"`
	Channel      channel.NotifChannel `protobuf:"varint,120,opt,name=Channel,proto3,enum=notif.manager.channel.v1.NotifChannel" json:"Channel,omitempty"`
	Notified     bool                 `protobuf:"varint,130,opt,name=Notified,proto3" json:"Notified,omitempty"`
	CreatedAt    uint32               `protobuf:"varint,140,opt,name=CreatedAt,proto3" json:"CreatedAt,omitempty"`
	UpdatedAt    uint32               `protobuf:"varint,150,opt,name=UpdatedAt,proto3" json:"UpdatedAt,omitempty"`
	LangID       string               `protobuf:"bytes,160,opt,name=LangID,proto3" json:"LangID,omitempty"`
	Lang         string               `protobuf:"bytes,170,opt,name=Lang,proto3" json:"Lang,omitempty"`
	EventID      string               `protobuf:"bytes,180,opt,name=EventID,proto3" json:"EventID,omitempty"`
}

func (x *Notif) Reset() {
	*x = Notif{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_notif_gw_v1_notif_notif_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Notif) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Notif) ProtoMessage() {}

func (x *Notif) ProtoReflect() protoreflect.Message {
	mi := &file_npool_notif_gw_v1_notif_notif_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Notif.ProtoReflect.Descriptor instead.
func (*Notif) Descriptor() ([]byte, []int) {
	return file_npool_notif_gw_v1_notif_notif_proto_rawDescGZIP(), []int{0}
}

func (x *Notif) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *Notif) GetAppID() string {
	if x != nil {
		return x.AppID
	}
	return ""
}

func (x *Notif) GetAppName() string {
	if x != nil {
		return x.AppName
	}
	return ""
}

func (x *Notif) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

func (x *Notif) GetEmailAddress() string {
	if x != nil {
		return x.EmailAddress
	}
	return ""
}

func (x *Notif) GetPhoneNO() string {
	if x != nil {
		return x.PhoneNO
	}
	return ""
}

func (x *Notif) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *Notif) GetEventType() v1.UsedFor {
	if x != nil {
		return x.EventType
	}
	return v1.UsedFor(0)
}

func (x *Notif) GetUseTemplate() bool {
	if x != nil {
		return x.UseTemplate
	}
	return false
}

func (x *Notif) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Notif) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *Notif) GetChannel() channel.NotifChannel {
	if x != nil {
		return x.Channel
	}
	return channel.NotifChannel(0)
}

func (x *Notif) GetNotified() bool {
	if x != nil {
		return x.Notified
	}
	return false
}

func (x *Notif) GetCreatedAt() uint32 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *Notif) GetUpdatedAt() uint32 {
	if x != nil {
		return x.UpdatedAt
	}
	return 0
}

func (x *Notif) GetLangID() string {
	if x != nil {
		return x.LangID
	}
	return ""
}

func (x *Notif) GetLang() string {
	if x != nil {
		return x.Lang
	}
	return ""
}

func (x *Notif) GetEventID() string {
	if x != nil {
		return x.EventID
	}
	return ""
}

type GetNotifRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID string `protobuf:"bytes,10,opt,name=ID,proto3" json:"ID,omitempty"`
}

func (x *GetNotifRequest) Reset() {
	*x = GetNotifRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_notif_gw_v1_notif_notif_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetNotifRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetNotifRequest) ProtoMessage() {}

func (x *GetNotifRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_notif_gw_v1_notif_notif_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetNotifRequest.ProtoReflect.Descriptor instead.
func (*GetNotifRequest) Descriptor() ([]byte, []int) {
	return file_npool_notif_gw_v1_notif_notif_proto_rawDescGZIP(), []int{1}
}

func (x *GetNotifRequest) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

type GetNotifResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *Notif `protobuf:"bytes,10,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (x *GetNotifResponse) Reset() {
	*x = GetNotifResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_notif_gw_v1_notif_notif_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetNotifResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetNotifResponse) ProtoMessage() {}

func (x *GetNotifResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_notif_gw_v1_notif_notif_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetNotifResponse.ProtoReflect.Descriptor instead.
func (*GetNotifResponse) Descriptor() ([]byte, []int) {
	return file_npool_notif_gw_v1_notif_notif_proto_rawDescGZIP(), []int{2}
}

func (x *GetNotifResponse) GetInfo() *Notif {
	if x != nil {
		return x.Info
	}
	return nil
}

type UpdateNotifsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IDs      []string `protobuf:"bytes,10,rep,name=IDs,proto3" json:"IDs,omitempty"`
	AppID    string   `protobuf:"bytes,20,opt,name=AppID,proto3" json:"AppID,omitempty"`
	UserID   string   `protobuf:"bytes,30,opt,name=UserID,proto3" json:"UserID,omitempty"`
	Notified bool     `protobuf:"varint,40,opt,name=Notified,proto3" json:"Notified,omitempty"`
}

func (x *UpdateNotifsRequest) Reset() {
	*x = UpdateNotifsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_notif_gw_v1_notif_notif_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateNotifsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateNotifsRequest) ProtoMessage() {}

func (x *UpdateNotifsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_notif_gw_v1_notif_notif_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateNotifsRequest.ProtoReflect.Descriptor instead.
func (*UpdateNotifsRequest) Descriptor() ([]byte, []int) {
	return file_npool_notif_gw_v1_notif_notif_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateNotifsRequest) GetIDs() []string {
	if x != nil {
		return x.IDs
	}
	return nil
}

func (x *UpdateNotifsRequest) GetAppID() string {
	if x != nil {
		return x.AppID
	}
	return ""
}

func (x *UpdateNotifsRequest) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

func (x *UpdateNotifsRequest) GetNotified() bool {
	if x != nil {
		return x.Notified
	}
	return false
}

type UpdateNotifsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Infos []*Notif `protobuf:"bytes,10,rep,name=Infos,proto3" json:"Infos,omitempty"`
}

func (x *UpdateNotifsResponse) Reset() {
	*x = UpdateNotifsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_notif_gw_v1_notif_notif_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateNotifsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateNotifsResponse) ProtoMessage() {}

func (x *UpdateNotifsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_notif_gw_v1_notif_notif_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateNotifsResponse.ProtoReflect.Descriptor instead.
func (*UpdateNotifsResponse) Descriptor() ([]byte, []int) {
	return file_npool_notif_gw_v1_notif_notif_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateNotifsResponse) GetInfos() []*Notif {
	if x != nil {
		return x.Infos
	}
	return nil
}

type GetNotifsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AppID  string `protobuf:"bytes,10,opt,name=AppID,proto3" json:"AppID,omitempty"`
	UserID string `protobuf:"bytes,20,opt,name=UserID,proto3" json:"UserID,omitempty"`
	LangID string `protobuf:"bytes,30,opt,name=LangID,proto3" json:"LangID,omitempty"`
	Offset int32  `protobuf:"varint,40,opt,name=Offset,proto3" json:"Offset,omitempty"`
	Limit  int32  `protobuf:"varint,50,opt,name=Limit,proto3" json:"Limit,omitempty"`
}

func (x *GetNotifsRequest) Reset() {
	*x = GetNotifsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_notif_gw_v1_notif_notif_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetNotifsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetNotifsRequest) ProtoMessage() {}

func (x *GetNotifsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_notif_gw_v1_notif_notif_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetNotifsRequest.ProtoReflect.Descriptor instead.
func (*GetNotifsRequest) Descriptor() ([]byte, []int) {
	return file_npool_notif_gw_v1_notif_notif_proto_rawDescGZIP(), []int{5}
}

func (x *GetNotifsRequest) GetAppID() string {
	if x != nil {
		return x.AppID
	}
	return ""
}

func (x *GetNotifsRequest) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

func (x *GetNotifsRequest) GetLangID() string {
	if x != nil {
		return x.LangID
	}
	return ""
}

func (x *GetNotifsRequest) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *GetNotifsRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type GetNotifsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Infos []*Notif `protobuf:"bytes,10,rep,name=Infos,proto3" json:"Infos,omitempty"`
	Total uint32   `protobuf:"varint,20,opt,name=Total,proto3" json:"Total,omitempty"`
}

func (x *GetNotifsResponse) Reset() {
	*x = GetNotifsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_notif_gw_v1_notif_notif_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetNotifsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetNotifsResponse) ProtoMessage() {}

func (x *GetNotifsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_notif_gw_v1_notif_notif_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetNotifsResponse.ProtoReflect.Descriptor instead.
func (*GetNotifsResponse) Descriptor() ([]byte, []int) {
	return file_npool_notif_gw_v1_notif_notif_proto_rawDescGZIP(), []int{6}
}

func (x *GetNotifsResponse) GetInfos() []*Notif {
	if x != nil {
		return x.Infos
	}
	return nil
}

func (x *GetNotifsResponse) GetTotal() uint32 {
	if x != nil {
		return x.Total
	}
	return 0
}

type GetAppNotifsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AppID  string `protobuf:"bytes,10,opt,name=AppID,proto3" json:"AppID,omitempty"`
	Offset int32  `protobuf:"varint,20,opt,name=Offset,proto3" json:"Offset,omitempty"`
	Limit  int32  `protobuf:"varint,30,opt,name=Limit,proto3" json:"Limit,omitempty"`
}

func (x *GetAppNotifsRequest) Reset() {
	*x = GetAppNotifsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_notif_gw_v1_notif_notif_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAppNotifsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAppNotifsRequest) ProtoMessage() {}

func (x *GetAppNotifsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_notif_gw_v1_notif_notif_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAppNotifsRequest.ProtoReflect.Descriptor instead.
func (*GetAppNotifsRequest) Descriptor() ([]byte, []int) {
	return file_npool_notif_gw_v1_notif_notif_proto_rawDescGZIP(), []int{7}
}

func (x *GetAppNotifsRequest) GetAppID() string {
	if x != nil {
		return x.AppID
	}
	return ""
}

func (x *GetAppNotifsRequest) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *GetAppNotifsRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type GetAppNotifsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Infos []*Notif `protobuf:"bytes,10,rep,name=Infos,proto3" json:"Infos,omitempty"`
	Total uint32   `protobuf:"varint,20,opt,name=Total,proto3" json:"Total,omitempty"`
}

func (x *GetAppNotifsResponse) Reset() {
	*x = GetAppNotifsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_notif_gw_v1_notif_notif_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAppNotifsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAppNotifsResponse) ProtoMessage() {}

func (x *GetAppNotifsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_notif_gw_v1_notif_notif_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAppNotifsResponse.ProtoReflect.Descriptor instead.
func (*GetAppNotifsResponse) Descriptor() ([]byte, []int) {
	return file_npool_notif_gw_v1_notif_notif_proto_rawDescGZIP(), []int{8}
}

func (x *GetAppNotifsResponse) GetInfos() []*Notif {
	if x != nil {
		return x.Infos
	}
	return nil
}

func (x *GetAppNotifsResponse) GetTotal() uint32 {
	if x != nil {
		return x.Total
	}
	return 0
}

var File_npool_notif_gw_v1_notif_notif_proto protoreflect.FileDescriptor

var file_npool_notif_gw_v1_notif_notif_proto_rawDesc = []byte{
	0x0a, 0x23, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x2f, 0x67, 0x77,
	0x2f, 0x76, 0x31, 0x2f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x2f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x17, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x2e, 0x67, 0x61, 0x74,
	0x65, 0x77, 0x61, 0x79, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x33, 0x2e, 0x76, 0x31, 0x1a, 0x1c,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x28, 0x6e, 0x70,
	0x6f, 0x6f, 0x6c, 0x2f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x2f, 0x6d, 0x67, 0x72, 0x2f, 0x76, 0x31,
	0x2f, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x2f, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x62, 0x61,
	0x73, 0x65, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x64, 0x66,
	0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa6, 0x04, 0x0a, 0x05, 0x4e, 0x6f, 0x74,
	0x69, 0x66, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x41, 0x70, 0x70, 0x49, 0x44, 0x18, 0x14, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x41, 0x70, 0x70, 0x49, 0x44, 0x12, 0x18, 0x0a, 0x07, 0x41, 0x70, 0x70, 0x4e,
	0x61, 0x6d, 0x65, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x41, 0x70, 0x70, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x28, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x22, 0x0a, 0x0c, 0x45, 0x6d,
	0x61, 0x69, 0x6c, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x32, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x18,
	0x0a, 0x07, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x4f, 0x18, 0x3c, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x4f, 0x12, 0x1a, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x46, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x55, 0x73, 0x65, 0x72,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x33, 0x0a, 0x09, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70,
	0x65, 0x18, 0x50, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x15, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x64, 0x46, 0x6f, 0x72, 0x52, 0x09,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x55, 0x73, 0x65,
	0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x18, 0x5a, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b,
	0x55, 0x73, 0x65, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x54,
	0x69, 0x74, 0x6c, 0x65, 0x18, 0x64, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x54, 0x69, 0x74, 0x6c,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x6e, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x40, 0x0a, 0x07, 0x43,
	0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x18, 0x78, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x26, 0x2e, 0x6e,
	0x6f, 0x74, 0x69, 0x66, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x63, 0x68, 0x61,
	0x6e, 0x6e, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x43, 0x68, 0x61,
	0x6e, 0x6e, 0x65, 0x6c, 0x52, 0x07, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x12, 0x1b, 0x0a,
	0x08, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x65, 0x64, 0x18, 0x82, 0x01, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x08, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x65, 0x64, 0x12, 0x1d, 0x0a, 0x09, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x8c, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x09, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x96, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x17, 0x0a, 0x06, 0x4c, 0x61, 0x6e, 0x67,
	0x49, 0x44, 0x18, 0xa0, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x4c, 0x61, 0x6e, 0x67, 0x49,
	0x44, 0x12, 0x13, 0x0a, 0x04, 0x4c, 0x61, 0x6e, 0x67, 0x18, 0xaa, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x4c, 0x61, 0x6e, 0x67, 0x12, 0x19, 0x0a, 0x07, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x49,
	0x44, 0x18, 0xb4, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x49,
	0x44, 0x22, 0x21, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x49, 0x44, 0x22, 0x46, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x4e, 0x6f, 0x74, 0x69, 0x66,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x32, 0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x2e, 0x67,
	0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x33, 0x2e, 0x76, 0x31,
	0x2e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x52, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x71, 0x0a, 0x13,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x49, 0x44, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x03, 0x49, 0x44, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x41, 0x70, 0x70, 0x49, 0x44, 0x18, 0x14,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x41, 0x70, 0x70, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x55,
	0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x55, 0x73, 0x65,
	0x72, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x65, 0x64, 0x18,
	0x28, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x65, 0x64, 0x22,
	0x4c, 0x0a, 0x14, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x34, 0x0a, 0x05, 0x49, 0x6e, 0x66, 0x6f, 0x73,
	0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x2e, 0x67,
	0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x33, 0x2e, 0x76, 0x31,
	0x2e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x52, 0x05, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x22, 0x86, 0x01,
	0x0a, 0x10, 0x47, 0x65, 0x74, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x41, 0x70, 0x70, 0x49, 0x44, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x41, 0x70, 0x70, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72,
	0x49, 0x44, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44,
	0x12, 0x16, 0x0a, 0x06, 0x4c, 0x61, 0x6e, 0x67, 0x49, 0x44, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x4c, 0x61, 0x6e, 0x67, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x4f, 0x66, 0x66, 0x73,
	0x65, 0x74, 0x18, 0x28, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74,
	0x12, 0x14, 0x0a, 0x05, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x32, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x05, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x22, 0x5f, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x4e, 0x6f, 0x74,
	0x69, 0x66, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x34, 0x0a, 0x05, 0x49,
	0x6e, 0x66, 0x6f, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x6e, 0x6f, 0x74,
	0x69, 0x66, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66,
	0x33, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x52, 0x05, 0x49, 0x6e, 0x66, 0x6f,
	0x73, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x14, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x22, 0x59, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x41, 0x70,
	0x70, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14,
	0x0a, 0x05, 0x41, 0x70, 0x70, 0x49, 0x44, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x41,
	0x70, 0x70, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x14,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05,
	0x4c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x4c, 0x69, 0x6d,
	0x69, 0x74, 0x22, 0x62, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x41, 0x70, 0x70, 0x4e, 0x6f, 0x74, 0x69,
	0x66, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x34, 0x0a, 0x05, 0x49, 0x6e,
	0x66, 0x6f, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x6e, 0x6f, 0x74, 0x69,
	0x66, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x33,
	0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x52, 0x05, 0x49, 0x6e, 0x66, 0x6f, 0x73,
	0x12, 0x14, 0x0a, 0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x14, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x32, 0x9c, 0x04, 0x0a, 0x07, 0x47, 0x61, 0x74, 0x65, 0x77,
	0x61, 0x79, 0x12, 0x79, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x12, 0x28,
	0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x6e,
	0x6f, 0x74, 0x69, 0x66, 0x33, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4e, 0x6f, 0x74, 0x69,
	0x66, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66,
	0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x33, 0x2e,
	0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x18, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x12, 0x3a, 0x01, 0x2a, 0x22, 0x0d,
	0x2f, 0x76, 0x31, 0x2f, 0x67, 0x65, 0x74, 0x2f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x12, 0x89, 0x01,
	0x0a, 0x0c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x73, 0x12, 0x2c,
	0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x6e,
	0x6f, 0x74, 0x69, 0x66, 0x33, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4e,
	0x6f, 0x74, 0x69, 0x66, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2d, 0x2e, 0x6e,
	0x6f, 0x74, 0x69, 0x66, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x6e, 0x6f, 0x74,
	0x69, 0x66, 0x33, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4e, 0x6f, 0x74,
	0x69, 0x66, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1c, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x16, 0x3a, 0x01, 0x2a, 0x22, 0x11, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x2f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x73, 0x12, 0x7d, 0x0a, 0x09, 0x47, 0x65, 0x74,
	0x4e, 0x6f, 0x74, 0x69, 0x66, 0x73, 0x12, 0x29, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x2e, 0x67,
	0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x33, 0x2e, 0x76, 0x31,
	0x2e, 0x47, 0x65, 0x74, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x2a, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61,
	0x79, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x33, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4e,
	0x6f, 0x74, 0x69, 0x66, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x19, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x13, 0x3a, 0x01, 0x2a, 0x22, 0x0e, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x65,
	0x74, 0x2f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x73, 0x12, 0x8a, 0x01, 0x0a, 0x0c, 0x47, 0x65, 0x74,
	0x41, 0x70, 0x70, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x73, 0x12, 0x2c, 0x2e, 0x6e, 0x6f, 0x74, 0x69,
	0x66, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x33,
	0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x70, 0x70, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2d, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x2e,
	0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x33, 0x2e, 0x76,
	0x31, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x70, 0x70, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x17, 0x3a, 0x01,
	0x2a, 0x22, 0x12, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x65, 0x74, 0x2f, 0x61, 0x70, 0x70, 0x2f, 0x6e,
	0x6f, 0x74, 0x69, 0x66, 0x73, 0x42, 0x3a, 0x5a, 0x38, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x4e, 0x70, 0x6f, 0x6f, 0x6c, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72,
	0x6d, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f,
	0x6e, 0x6f, 0x74, 0x69, 0x66, 0x2f, 0x67, 0x77, 0x2f, 0x76, 0x31, 0x2f, 0x6e, 0x6f, 0x74, 0x69,
	0x66, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_npool_notif_gw_v1_notif_notif_proto_rawDescOnce sync.Once
	file_npool_notif_gw_v1_notif_notif_proto_rawDescData = file_npool_notif_gw_v1_notif_notif_proto_rawDesc
)

func file_npool_notif_gw_v1_notif_notif_proto_rawDescGZIP() []byte {
	file_npool_notif_gw_v1_notif_notif_proto_rawDescOnce.Do(func() {
		file_npool_notif_gw_v1_notif_notif_proto_rawDescData = protoimpl.X.CompressGZIP(file_npool_notif_gw_v1_notif_notif_proto_rawDescData)
	})
	return file_npool_notif_gw_v1_notif_notif_proto_rawDescData
}

var file_npool_notif_gw_v1_notif_notif_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_npool_notif_gw_v1_notif_notif_proto_goTypes = []interface{}{
	(*Notif)(nil),                // 0: notif.gateway.notif3.v1.Notif
	(*GetNotifRequest)(nil),      // 1: notif.gateway.notif3.v1.GetNotifRequest
	(*GetNotifResponse)(nil),     // 2: notif.gateway.notif3.v1.GetNotifResponse
	(*UpdateNotifsRequest)(nil),  // 3: notif.gateway.notif3.v1.UpdateNotifsRequest
	(*UpdateNotifsResponse)(nil), // 4: notif.gateway.notif3.v1.UpdateNotifsResponse
	(*GetNotifsRequest)(nil),     // 5: notif.gateway.notif3.v1.GetNotifsRequest
	(*GetNotifsResponse)(nil),    // 6: notif.gateway.notif3.v1.GetNotifsResponse
	(*GetAppNotifsRequest)(nil),  // 7: notif.gateway.notif3.v1.GetAppNotifsRequest
	(*GetAppNotifsResponse)(nil), // 8: notif.gateway.notif3.v1.GetAppNotifsResponse
	(v1.UsedFor)(0),              // 9: basetypes.v1.UsedFor
	(channel.NotifChannel)(0),    // 10: notif.manager.channel.v1.NotifChannel
}
var file_npool_notif_gw_v1_notif_notif_proto_depIdxs = []int32{
	9,  // 0: notif.gateway.notif3.v1.Notif.EventType:type_name -> basetypes.v1.UsedFor
	10, // 1: notif.gateway.notif3.v1.Notif.Channel:type_name -> notif.manager.channel.v1.NotifChannel
	0,  // 2: notif.gateway.notif3.v1.GetNotifResponse.Info:type_name -> notif.gateway.notif3.v1.Notif
	0,  // 3: notif.gateway.notif3.v1.UpdateNotifsResponse.Infos:type_name -> notif.gateway.notif3.v1.Notif
	0,  // 4: notif.gateway.notif3.v1.GetNotifsResponse.Infos:type_name -> notif.gateway.notif3.v1.Notif
	0,  // 5: notif.gateway.notif3.v1.GetAppNotifsResponse.Infos:type_name -> notif.gateway.notif3.v1.Notif
	1,  // 6: notif.gateway.notif3.v1.Gateway.GetNotif:input_type -> notif.gateway.notif3.v1.GetNotifRequest
	3,  // 7: notif.gateway.notif3.v1.Gateway.UpdateNotifs:input_type -> notif.gateway.notif3.v1.UpdateNotifsRequest
	5,  // 8: notif.gateway.notif3.v1.Gateway.GetNotifs:input_type -> notif.gateway.notif3.v1.GetNotifsRequest
	7,  // 9: notif.gateway.notif3.v1.Gateway.GetAppNotifs:input_type -> notif.gateway.notif3.v1.GetAppNotifsRequest
	2,  // 10: notif.gateway.notif3.v1.Gateway.GetNotif:output_type -> notif.gateway.notif3.v1.GetNotifResponse
	4,  // 11: notif.gateway.notif3.v1.Gateway.UpdateNotifs:output_type -> notif.gateway.notif3.v1.UpdateNotifsResponse
	6,  // 12: notif.gateway.notif3.v1.Gateway.GetNotifs:output_type -> notif.gateway.notif3.v1.GetNotifsResponse
	8,  // 13: notif.gateway.notif3.v1.Gateway.GetAppNotifs:output_type -> notif.gateway.notif3.v1.GetAppNotifsResponse
	10, // [10:14] is the sub-list for method output_type
	6,  // [6:10] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_npool_notif_gw_v1_notif_notif_proto_init() }
func file_npool_notif_gw_v1_notif_notif_proto_init() {
	if File_npool_notif_gw_v1_notif_notif_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_npool_notif_gw_v1_notif_notif_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Notif); i {
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
		file_npool_notif_gw_v1_notif_notif_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetNotifRequest); i {
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
		file_npool_notif_gw_v1_notif_notif_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetNotifResponse); i {
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
		file_npool_notif_gw_v1_notif_notif_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateNotifsRequest); i {
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
		file_npool_notif_gw_v1_notif_notif_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateNotifsResponse); i {
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
		file_npool_notif_gw_v1_notif_notif_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetNotifsRequest); i {
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
		file_npool_notif_gw_v1_notif_notif_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetNotifsResponse); i {
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
		file_npool_notif_gw_v1_notif_notif_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAppNotifsRequest); i {
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
		file_npool_notif_gw_v1_notif_notif_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAppNotifsResponse); i {
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
			RawDescriptor: file_npool_notif_gw_v1_notif_notif_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_npool_notif_gw_v1_notif_notif_proto_goTypes,
		DependencyIndexes: file_npool_notif_gw_v1_notif_notif_proto_depIdxs,
		MessageInfos:      file_npool_notif_gw_v1_notif_notif_proto_msgTypes,
	}.Build()
	File_npool_notif_gw_v1_notif_notif_proto = out.File
	file_npool_notif_gw_v1_notif_notif_proto_rawDesc = nil
	file_npool_notif_gw_v1_notif_notif_proto_goTypes = nil
	file_npool_notif_gw_v1_notif_notif_proto_depIdxs = nil
}
