// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.18.1
// source: npool/notif/mw/v1/announcement/sendstate/sendstate.proto

package sendstate

import (
	npool "github.com/NpoolPlatform/message/npool"
	announcement "github.com/NpoolPlatform/message/npool/notif/mgr/v1/announcement"
	sendstate "github.com/NpoolPlatform/message/npool/notif/mgr/v1/announcement/sendstate"
	channel "github.com/NpoolPlatform/message/npool/notif/mgr/v1/channel"
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

type SendState struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: sql:"announcement_id"
	AnnouncementID string `protobuf:"bytes,10,opt,name=AnnouncementID,proto3" json:"AnnouncementID,omitempty" sql:"announcement_id"`
	// @inject_tag: sql:"app_id"
	AppID string `protobuf:"bytes,20,opt,name=AppID,proto3" json:"AppID,omitempty" sql:"app_id"`
	// @inject_tag: sql:"user_id"
	UserID string `protobuf:"bytes,30,opt,name=UserID,proto3" json:"UserID,omitempty" sql:"user_id"`
	// @inject_tag: sql:"title"
	Title string `protobuf:"bytes,40,opt,name=Title,proto3" json:"Title,omitempty" sql:"title"`
	// @inject_tag: sql:"content"
	Content string `protobuf:"bytes,50,opt,name=Content,proto3" json:"Content,omitempty" sql:"content"`
	// @inject_tag: sql:"channel"
	ChannelStr string               `protobuf:"bytes,60,opt,name=ChannelStr,proto3" json:"ChannelStr,omitempty" sql:"channel"`
	Channel    channel.NotifChannel `protobuf:"varint,70,opt,name=Channel,proto3,enum=notif.manager.channel.v1.NotifChannel" json:"Channel,omitempty"`
	// @inject_tag: sql:"type"
	AnnouncementTypeStr string                        `protobuf:"bytes,80,opt,name=AnnouncementTypeStr,proto3" json:"AnnouncementTypeStr,omitempty" sql:"type"`
	AnnouncementType    announcement.AnnouncementType `protobuf:"varint,90,opt,name=AnnouncementType,proto3,enum=notif.manager.announcement.v1.AnnouncementType" json:"AnnouncementType,omitempty"`
	// @inject_tag: sql:"created_at"
	CreatedAt uint32 `protobuf:"varint,100,opt,name=CreatedAt,proto3" json:"CreatedAt,omitempty" sql:"created_at"`
	// @inject_tag: sql:"updated_at"
	UpdatedAt uint32 `protobuf:"varint,110,opt,name=UpdatedAt,proto3" json:"UpdatedAt,omitempty" sql:"updated_at"`
}

func (x *SendState) Reset() {
	*x = SendState{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_notif_mw_v1_announcement_sendstate_sendstate_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendState) ProtoMessage() {}

func (x *SendState) ProtoReflect() protoreflect.Message {
	mi := &file_npool_notif_mw_v1_announcement_sendstate_sendstate_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendState.ProtoReflect.Descriptor instead.
func (*SendState) Descriptor() ([]byte, []int) {
	return file_npool_notif_mw_v1_announcement_sendstate_sendstate_proto_rawDescGZIP(), []int{0}
}

func (x *SendState) GetAnnouncementID() string {
	if x != nil {
		return x.AnnouncementID
	}
	return ""
}

func (x *SendState) GetAppID() string {
	if x != nil {
		return x.AppID
	}
	return ""
}

func (x *SendState) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

func (x *SendState) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *SendState) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *SendState) GetChannelStr() string {
	if x != nil {
		return x.ChannelStr
	}
	return ""
}

func (x *SendState) GetChannel() channel.NotifChannel {
	if x != nil {
		return x.Channel
	}
	return channel.NotifChannel(0)
}

func (x *SendState) GetAnnouncementTypeStr() string {
	if x != nil {
		return x.AnnouncementTypeStr
	}
	return ""
}

func (x *SendState) GetAnnouncementType() announcement.AnnouncementType {
	if x != nil {
		return x.AnnouncementType
	}
	return announcement.AnnouncementType(0)
}

func (x *SendState) GetCreatedAt() uint32 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *SendState) GetUpdatedAt() uint32 {
	if x != nil {
		return x.UpdatedAt
	}
	return 0
}

type CreateSendStateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AppID          string               `protobuf:"bytes,20,opt,name=AppID,proto3" json:"AppID,omitempty"`
	UserID         string               `protobuf:"bytes,30,opt,name=UserID,proto3" json:"UserID,omitempty"`
	AnnouncementID string               `protobuf:"bytes,40,opt,name=AnnouncementID,proto3" json:"AnnouncementID,omitempty"`
	Channel        channel.NotifChannel `protobuf:"varint,50,opt,name=Channel,proto3,enum=notif.manager.channel.v1.NotifChannel" json:"Channel,omitempty"`
}

func (x *CreateSendStateRequest) Reset() {
	*x = CreateSendStateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_notif_mw_v1_announcement_sendstate_sendstate_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateSendStateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSendStateRequest) ProtoMessage() {}

func (x *CreateSendStateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_notif_mw_v1_announcement_sendstate_sendstate_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSendStateRequest.ProtoReflect.Descriptor instead.
func (*CreateSendStateRequest) Descriptor() ([]byte, []int) {
	return file_npool_notif_mw_v1_announcement_sendstate_sendstate_proto_rawDescGZIP(), []int{1}
}

func (x *CreateSendStateRequest) GetAppID() string {
	if x != nil {
		return x.AppID
	}
	return ""
}

func (x *CreateSendStateRequest) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

func (x *CreateSendStateRequest) GetAnnouncementID() string {
	if x != nil {
		return x.AnnouncementID
	}
	return ""
}

func (x *CreateSendStateRequest) GetChannel() channel.NotifChannel {
	if x != nil {
		return x.Channel
	}
	return channel.NotifChannel(0)
}

type CreateSendStateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateSendStateResponse) Reset() {
	*x = CreateSendStateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_notif_mw_v1_announcement_sendstate_sendstate_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateSendStateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSendStateResponse) ProtoMessage() {}

func (x *CreateSendStateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_notif_mw_v1_announcement_sendstate_sendstate_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSendStateResponse.ProtoReflect.Descriptor instead.
func (*CreateSendStateResponse) Descriptor() ([]byte, []int) {
	return file_npool_notif_mw_v1_announcement_sendstate_sendstate_proto_rawDescGZIP(), []int{2}
}

type CreateSendStatesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Infos []*sendstate.SendStateReq `protobuf:"bytes,10,rep,name=Infos,proto3" json:"Infos,omitempty"`
}

func (x *CreateSendStatesRequest) Reset() {
	*x = CreateSendStatesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_notif_mw_v1_announcement_sendstate_sendstate_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateSendStatesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSendStatesRequest) ProtoMessage() {}

func (x *CreateSendStatesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_notif_mw_v1_announcement_sendstate_sendstate_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSendStatesRequest.ProtoReflect.Descriptor instead.
func (*CreateSendStatesRequest) Descriptor() ([]byte, []int) {
	return file_npool_notif_mw_v1_announcement_sendstate_sendstate_proto_rawDescGZIP(), []int{3}
}

func (x *CreateSendStatesRequest) GetInfos() []*sendstate.SendStateReq {
	if x != nil {
		return x.Infos
	}
	return nil
}

type CreateSendStatesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateSendStatesResponse) Reset() {
	*x = CreateSendStatesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_notif_mw_v1_announcement_sendstate_sendstate_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateSendStatesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSendStatesResponse) ProtoMessage() {}

func (x *CreateSendStatesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_notif_mw_v1_announcement_sendstate_sendstate_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSendStatesResponse.ProtoReflect.Descriptor instead.
func (*CreateSendStatesResponse) Descriptor() ([]byte, []int) {
	return file_npool_notif_mw_v1_announcement_sendstate_sendstate_proto_rawDescGZIP(), []int{4}
}

type Conds struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID             *npool.StringVal      `protobuf:"bytes,10,opt,name=ID,proto3,oneof" json:"ID,omitempty"`
	AppID          *npool.StringVal      `protobuf:"bytes,20,opt,name=AppID,proto3,oneof" json:"AppID,omitempty"`
	UserID         *npool.StringVal      `protobuf:"bytes,30,opt,name=UserID,proto3,oneof" json:"UserID,omitempty"`
	AnnouncementID *npool.StringVal      `protobuf:"bytes,40,opt,name=AnnouncementID,proto3,oneof" json:"AnnouncementID,omitempty"`
	Channel        *npool.Uint32Val      `protobuf:"bytes,50,opt,name=Channel,proto3,oneof" json:"Channel,omitempty"`
	EndAt          *npool.Uint32Val      `protobuf:"bytes,60,opt,name=EndAt,proto3,oneof" json:"EndAt,omitempty"`
	UserIDs        *npool.StringSliceVal `protobuf:"bytes,70,opt,name=UserIDs,proto3,oneof" json:"UserIDs,omitempty"`
}

func (x *Conds) Reset() {
	*x = Conds{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_notif_mw_v1_announcement_sendstate_sendstate_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Conds) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Conds) ProtoMessage() {}

func (x *Conds) ProtoReflect() protoreflect.Message {
	mi := &file_npool_notif_mw_v1_announcement_sendstate_sendstate_proto_msgTypes[5]
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
	return file_npool_notif_mw_v1_announcement_sendstate_sendstate_proto_rawDescGZIP(), []int{5}
}

func (x *Conds) GetID() *npool.StringVal {
	if x != nil {
		return x.ID
	}
	return nil
}

func (x *Conds) GetAppID() *npool.StringVal {
	if x != nil {
		return x.AppID
	}
	return nil
}

func (x *Conds) GetUserID() *npool.StringVal {
	if x != nil {
		return x.UserID
	}
	return nil
}

func (x *Conds) GetAnnouncementID() *npool.StringVal {
	if x != nil {
		return x.AnnouncementID
	}
	return nil
}

func (x *Conds) GetChannel() *npool.Uint32Val {
	if x != nil {
		return x.Channel
	}
	return nil
}

func (x *Conds) GetEndAt() *npool.Uint32Val {
	if x != nil {
		return x.EndAt
	}
	return nil
}

func (x *Conds) GetUserIDs() *npool.StringSliceVal {
	if x != nil {
		return x.UserIDs
	}
	return nil
}

type GetSendStatesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Conds  *Conds `protobuf:"bytes,10,opt,name=Conds,proto3" json:"Conds,omitempty"`
	Offset int32  `protobuf:"varint,30,opt,name=Offset,proto3" json:"Offset,omitempty"`
	Limit  int32  `protobuf:"varint,40,opt,name=Limit,proto3" json:"Limit,omitempty"`
}

func (x *GetSendStatesRequest) Reset() {
	*x = GetSendStatesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_notif_mw_v1_announcement_sendstate_sendstate_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSendStatesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSendStatesRequest) ProtoMessage() {}

func (x *GetSendStatesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_notif_mw_v1_announcement_sendstate_sendstate_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSendStatesRequest.ProtoReflect.Descriptor instead.
func (*GetSendStatesRequest) Descriptor() ([]byte, []int) {
	return file_npool_notif_mw_v1_announcement_sendstate_sendstate_proto_rawDescGZIP(), []int{6}
}

func (x *GetSendStatesRequest) GetConds() *Conds {
	if x != nil {
		return x.Conds
	}
	return nil
}

func (x *GetSendStatesRequest) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *GetSendStatesRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type GetSendStatesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Infos []*SendState `protobuf:"bytes,10,rep,name=Infos,proto3" json:"Infos,omitempty"`
	Total uint32       `protobuf:"varint,20,opt,name=Total,proto3" json:"Total,omitempty"`
}

func (x *GetSendStatesResponse) Reset() {
	*x = GetSendStatesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_notif_mw_v1_announcement_sendstate_sendstate_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSendStatesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSendStatesResponse) ProtoMessage() {}

func (x *GetSendStatesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_notif_mw_v1_announcement_sendstate_sendstate_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSendStatesResponse.ProtoReflect.Descriptor instead.
func (*GetSendStatesResponse) Descriptor() ([]byte, []int) {
	return file_npool_notif_mw_v1_announcement_sendstate_sendstate_proto_rawDescGZIP(), []int{7}
}

func (x *GetSendStatesResponse) GetInfos() []*SendState {
	if x != nil {
		return x.Infos
	}
	return nil
}

func (x *GetSendStatesResponse) GetTotal() uint32 {
	if x != nil {
		return x.Total
	}
	return 0
}

var File_npool_notif_mw_v1_announcement_sendstate_sendstate_proto protoreflect.FileDescriptor

var file_npool_notif_mw_v1_announcement_sendstate_sendstate_proto_rawDesc = []byte{
	0x0a, 0x38, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x2f, 0x6d, 0x77,
	0x2f, 0x76, 0x31, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x75, 0x6e, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74,
	0x2f, 0x73, 0x65, 0x6e, 0x64, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2f, 0x73, 0x65, 0x6e, 0x64, 0x73,
	0x74, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x2a, 0x6e, 0x6f, 0x74, 0x69,
	0x66, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x61, 0x6e, 0x6e,
	0x6f, 0x75, 0x6e, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x73, 0x65, 0x6e, 0x64, 0x73, 0x74,
	0x61, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x32, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x6e, 0x6f,
	0x74, 0x69, 0x66, 0x2f, 0x6d, 0x67, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x75,
	0x6e, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x75, 0x6e, 0x63, 0x65,
	0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x39, 0x6e, 0x70, 0x6f, 0x6f,
	0x6c, 0x2f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x2f, 0x6d, 0x67, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x75, 0x6e, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x73, 0x65, 0x6e, 0x64,
	0x73, 0x74, 0x61, 0x74, 0x65, 0x2f, 0x73, 0x65, 0x6e, 0x64, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x28, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x6e, 0x6f, 0x74,
	0x69, 0x66, 0x2f, 0x6d, 0x67, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65,
	0x6c, 0x2f, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x11, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xbe, 0x03, 0x0a, 0x09, 0x53, 0x65, 0x6e, 0x64, 0x53, 0x74, 0x61, 0x74, 0x65,
	0x12, 0x26, 0x0a, 0x0e, 0x41, 0x6e, 0x6e, 0x6f, 0x75, 0x6e, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74,
	0x49, 0x44, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x41, 0x6e, 0x6e, 0x6f, 0x75, 0x6e,
	0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x41, 0x70, 0x70, 0x49,
	0x44, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x41, 0x70, 0x70, 0x49, 0x44, 0x12, 0x16,
	0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x18,
	0x28, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x32, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x43,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65,
	0x6c, 0x53, 0x74, 0x72, 0x18, 0x3c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x43, 0x68, 0x61, 0x6e,
	0x6e, 0x65, 0x6c, 0x53, 0x74, 0x72, 0x12, 0x40, 0x0a, 0x07, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65,
	0x6c, 0x18, 0x46, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x26, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x2e,
	0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x2e,
	0x76, 0x31, 0x2e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x52,
	0x07, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x12, 0x30, 0x0a, 0x13, 0x41, 0x6e, 0x6e, 0x6f,
	0x75, 0x6e, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x53, 0x74, 0x72, 0x18,
	0x50, 0x20, 0x01, 0x28, 0x09, 0x52, 0x13, 0x41, 0x6e, 0x6e, 0x6f, 0x75, 0x6e, 0x63, 0x65, 0x6d,
	0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x53, 0x74, 0x72, 0x12, 0x5b, 0x0a, 0x10, 0x41, 0x6e,
	0x6e, 0x6f, 0x75, 0x6e, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x18, 0x5a,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x2f, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x2e, 0x6d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x72, 0x2e, 0x61, 0x6e, 0x6e, 0x6f, 0x75, 0x6e, 0x63, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x6e, 0x6e, 0x6f, 0x75, 0x6e, 0x63, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x10, 0x41, 0x6e, 0x6e, 0x6f, 0x75, 0x6e, 0x63, 0x65, 0x6d,
	0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x18, 0x64, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x18, 0x6e, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x22, 0xb0, 0x01, 0x0a, 0x16, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x65,
	0x6e, 0x64, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14,
	0x0a, 0x05, 0x41, 0x70, 0x70, 0x49, 0x44, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x41,
	0x70, 0x70, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x1e,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x26, 0x0a, 0x0e,
	0x41, 0x6e, 0x6e, 0x6f, 0x75, 0x6e, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x28,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x41, 0x6e, 0x6e, 0x6f, 0x75, 0x6e, 0x63, 0x65, 0x6d, 0x65,
	0x6e, 0x74, 0x49, 0x44, 0x12, 0x40, 0x0a, 0x07, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x18,
	0x32, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x26, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x2e, 0x6d, 0x61,
	0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x2e, 0x76, 0x31,
	0x2e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x52, 0x07, 0x43,
	0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x22, 0x19, 0x0a, 0x17, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x53, 0x65, 0x6e, 0x64, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x59, 0x0a, 0x17, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x65, 0x6e, 0x64, 0x53,
	0x74, 0x61, 0x74, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3e, 0x0a, 0x05,
	0x49, 0x6e, 0x66, 0x6f, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x6e, 0x6f,
	0x74, 0x69, 0x66, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x73, 0x65, 0x6e, 0x64,
	0x73, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x53, 0x74, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x71, 0x52, 0x05, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x22, 0x1a, 0x0a, 0x18,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x65, 0x6e, 0x64, 0x53, 0x74, 0x61, 0x74, 0x65, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0xc3, 0x03, 0x0a, 0x05, 0x43, 0x6f, 0x6e,
	0x64, 0x73, 0x12, 0x28, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13,
	0x2e, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67,
	0x56, 0x61, 0x6c, 0x48, 0x00, 0x52, 0x02, 0x49, 0x44, 0x88, 0x01, 0x01, 0x12, 0x2e, 0x0a, 0x05,
	0x41, 0x70, 0x70, 0x49, 0x44, 0x18, 0x14, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x6e, 0x70,
	0x6f, 0x6f, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c,
	0x48, 0x01, 0x52, 0x05, 0x41, 0x70, 0x70, 0x49, 0x44, 0x88, 0x01, 0x01, 0x12, 0x30, 0x0a, 0x06,
	0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x6e,
	0x70, 0x6f, 0x6f, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61,
	0x6c, 0x48, 0x02, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x88, 0x01, 0x01, 0x12, 0x40,
	0x0a, 0x0e, 0x41, 0x6e, 0x6e, 0x6f, 0x75, 0x6e, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x44,
	0x18, 0x28, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2e, 0x76,
	0x31, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x48, 0x03, 0x52, 0x0e, 0x41,
	0x6e, 0x6e, 0x6f, 0x75, 0x6e, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x44, 0x88, 0x01, 0x01,
	0x12, 0x32, 0x0a, 0x07, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x18, 0x32, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x13, 0x2e, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x69, 0x6e,
	0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x48, 0x04, 0x52, 0x07, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65,
	0x6c, 0x88, 0x01, 0x01, 0x12, 0x2e, 0x0a, 0x05, 0x45, 0x6e, 0x64, 0x41, 0x74, 0x18, 0x3c, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x55,
	0x69, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x48, 0x05, 0x52, 0x05, 0x45, 0x6e, 0x64, 0x41,
	0x74, 0x88, 0x01, 0x01, 0x12, 0x37, 0x0a, 0x07, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x73, 0x18,
	0x46, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2e, 0x76, 0x31,
	0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x53, 0x6c, 0x69, 0x63, 0x65, 0x56, 0x61, 0x6c, 0x48,
	0x06, 0x52, 0x07, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x73, 0x88, 0x01, 0x01, 0x42, 0x05, 0x0a,
	0x03, 0x5f, 0x49, 0x44, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x41, 0x70, 0x70, 0x49, 0x44, 0x42, 0x09,
	0x0a, 0x07, 0x5f, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x42, 0x11, 0x0a, 0x0f, 0x5f, 0x41, 0x6e,
	0x6e, 0x6f, 0x75, 0x6e, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x44, 0x42, 0x0a, 0x0a, 0x08,
	0x5f, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x45, 0x6e, 0x64,
	0x41, 0x74, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x73, 0x22, 0x8d,
	0x01, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x53, 0x65, 0x6e, 0x64, 0x53, 0x74, 0x61, 0x74, 0x65, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x47, 0x0a, 0x05, 0x43, 0x6f, 0x6e, 0x64, 0x73,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x31, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x2e, 0x6d,
	0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x61, 0x6e, 0x6e, 0x6f, 0x75, 0x6e,
	0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x73, 0x65, 0x6e, 0x64, 0x73, 0x74, 0x61, 0x74, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x64, 0x73, 0x52, 0x05, 0x43, 0x6f, 0x6e, 0x64, 0x73,
	0x12, 0x16, 0x0a, 0x06, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x06, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x4c, 0x69, 0x6d, 0x69,
	0x74, 0x18, 0x28, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x22, 0x7a,
	0x0a, 0x15, 0x47, 0x65, 0x74, 0x53, 0x65, 0x6e, 0x64, 0x53, 0x74, 0x61, 0x74, 0x65, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4b, 0x0a, 0x05, 0x49, 0x6e, 0x66, 0x6f, 0x73,
	0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x35, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x2e, 0x6d,
	0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x61, 0x6e, 0x6e, 0x6f, 0x75, 0x6e,
	0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x73, 0x65, 0x6e, 0x64, 0x73, 0x74, 0x61, 0x74, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x49,
	0x6e, 0x66, 0x6f, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x14, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x32, 0xe6, 0x03, 0x0a, 0x0a, 0x4d,
	0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x12, 0x9c, 0x01, 0x0a, 0x0f, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x53, 0x65, 0x6e, 0x64, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x42, 0x2e,
	0x6e, 0x6f, 0x74, 0x69, 0x66, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65,
	0x2e, 0x61, 0x6e, 0x6e, 0x6f, 0x75, 0x6e, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x73, 0x65,
	0x6e, 0x64, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x53, 0x65, 0x6e, 0x64, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x43, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65,
	0x77, 0x61, 0x72, 0x65, 0x2e, 0x61, 0x6e, 0x6e, 0x6f, 0x75, 0x6e, 0x63, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x2e, 0x73, 0x65, 0x6e, 0x64, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x65, 0x6e, 0x64, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x9f, 0x01, 0x0a, 0x10, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x53, 0x65, 0x6e, 0x64, 0x53, 0x74, 0x61, 0x74, 0x65, 0x73, 0x12, 0x43, 0x2e,
	0x6e, 0x6f, 0x74, 0x69, 0x66, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65,
	0x2e, 0x61, 0x6e, 0x6e, 0x6f, 0x75, 0x6e, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x73, 0x65,
	0x6e, 0x64, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x53, 0x65, 0x6e, 0x64, 0x53, 0x74, 0x61, 0x74, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x44, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c,
	0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x61, 0x6e, 0x6e, 0x6f, 0x75, 0x6e, 0x63, 0x65, 0x6d, 0x65,
	0x6e, 0x74, 0x2e, 0x73, 0x65, 0x6e, 0x64, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x65, 0x6e, 0x64, 0x53, 0x74, 0x61, 0x74, 0x65, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x96, 0x01, 0x0a, 0x0d, 0x47,
	0x65, 0x74, 0x53, 0x65, 0x6e, 0x64, 0x53, 0x74, 0x61, 0x74, 0x65, 0x73, 0x12, 0x40, 0x2e, 0x6e,
	0x6f, 0x74, 0x69, 0x66, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e,
	0x61, 0x6e, 0x6e, 0x6f, 0x75, 0x6e, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x73, 0x65, 0x6e,
	0x64, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x65, 0x6e,
	0x64, 0x53, 0x74, 0x61, 0x74, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x41,
	0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72,
	0x65, 0x2e, 0x61, 0x6e, 0x6e, 0x6f, 0x75, 0x6e, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x73,
	0x65, 0x6e, 0x64, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x53,
	0x65, 0x6e, 0x64, 0x53, 0x74, 0x61, 0x74, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x42, 0x4b, 0x5a, 0x49, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x4e, 0x70, 0x6f, 0x6f, 0x6c, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x6e, 0x6f,
	0x74, 0x69, 0x66, 0x2f, 0x6d, 0x77, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x75, 0x6e,
	0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x73, 0x65, 0x6e, 0x64, 0x73, 0x74, 0x61, 0x74, 0x65,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_npool_notif_mw_v1_announcement_sendstate_sendstate_proto_rawDescOnce sync.Once
	file_npool_notif_mw_v1_announcement_sendstate_sendstate_proto_rawDescData = file_npool_notif_mw_v1_announcement_sendstate_sendstate_proto_rawDesc
)

func file_npool_notif_mw_v1_announcement_sendstate_sendstate_proto_rawDescGZIP() []byte {
	file_npool_notif_mw_v1_announcement_sendstate_sendstate_proto_rawDescOnce.Do(func() {
		file_npool_notif_mw_v1_announcement_sendstate_sendstate_proto_rawDescData = protoimpl.X.CompressGZIP(file_npool_notif_mw_v1_announcement_sendstate_sendstate_proto_rawDescData)
	})
	return file_npool_notif_mw_v1_announcement_sendstate_sendstate_proto_rawDescData
}

var file_npool_notif_mw_v1_announcement_sendstate_sendstate_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_npool_notif_mw_v1_announcement_sendstate_sendstate_proto_goTypes = []interface{}{
	(*SendState)(nil),                  // 0: notif.middleware.announcement.sendstate.v1.SendState
	(*CreateSendStateRequest)(nil),     // 1: notif.middleware.announcement.sendstate.v1.CreateSendStateRequest
	(*CreateSendStateResponse)(nil),    // 2: notif.middleware.announcement.sendstate.v1.CreateSendStateResponse
	(*CreateSendStatesRequest)(nil),    // 3: notif.middleware.announcement.sendstate.v1.CreateSendStatesRequest
	(*CreateSendStatesResponse)(nil),   // 4: notif.middleware.announcement.sendstate.v1.CreateSendStatesResponse
	(*Conds)(nil),                      // 5: notif.middleware.announcement.sendstate.v1.Conds
	(*GetSendStatesRequest)(nil),       // 6: notif.middleware.announcement.sendstate.v1.GetSendStatesRequest
	(*GetSendStatesResponse)(nil),      // 7: notif.middleware.announcement.sendstate.v1.GetSendStatesResponse
	(channel.NotifChannel)(0),          // 8: notif.manager.channel.v1.NotifChannel
	(announcement.AnnouncementType)(0), // 9: notif.manager.announcement.v1.AnnouncementType
	(*sendstate.SendStateReq)(nil),     // 10: notif.manager.sendstate.v1.SendStateReq
	(*npool.StringVal)(nil),            // 11: npool.v1.StringVal
	(*npool.Uint32Val)(nil),            // 12: npool.v1.Uint32Val
	(*npool.StringSliceVal)(nil),       // 13: npool.v1.StringSliceVal
}
var file_npool_notif_mw_v1_announcement_sendstate_sendstate_proto_depIdxs = []int32{
	8,  // 0: notif.middleware.announcement.sendstate.v1.SendState.Channel:type_name -> notif.manager.channel.v1.NotifChannel
	9,  // 1: notif.middleware.announcement.sendstate.v1.SendState.AnnouncementType:type_name -> notif.manager.announcement.v1.AnnouncementType
	8,  // 2: notif.middleware.announcement.sendstate.v1.CreateSendStateRequest.Channel:type_name -> notif.manager.channel.v1.NotifChannel
	10, // 3: notif.middleware.announcement.sendstate.v1.CreateSendStatesRequest.Infos:type_name -> notif.manager.sendstate.v1.SendStateReq
	11, // 4: notif.middleware.announcement.sendstate.v1.Conds.ID:type_name -> npool.v1.StringVal
	11, // 5: notif.middleware.announcement.sendstate.v1.Conds.AppID:type_name -> npool.v1.StringVal
	11, // 6: notif.middleware.announcement.sendstate.v1.Conds.UserID:type_name -> npool.v1.StringVal
	11, // 7: notif.middleware.announcement.sendstate.v1.Conds.AnnouncementID:type_name -> npool.v1.StringVal
	12, // 8: notif.middleware.announcement.sendstate.v1.Conds.Channel:type_name -> npool.v1.Uint32Val
	12, // 9: notif.middleware.announcement.sendstate.v1.Conds.EndAt:type_name -> npool.v1.Uint32Val
	13, // 10: notif.middleware.announcement.sendstate.v1.Conds.UserIDs:type_name -> npool.v1.StringSliceVal
	5,  // 11: notif.middleware.announcement.sendstate.v1.GetSendStatesRequest.Conds:type_name -> notif.middleware.announcement.sendstate.v1.Conds
	0,  // 12: notif.middleware.announcement.sendstate.v1.GetSendStatesResponse.Infos:type_name -> notif.middleware.announcement.sendstate.v1.SendState
	1,  // 13: notif.middleware.announcement.sendstate.v1.Middleware.CreateSendState:input_type -> notif.middleware.announcement.sendstate.v1.CreateSendStateRequest
	3,  // 14: notif.middleware.announcement.sendstate.v1.Middleware.CreateSendStates:input_type -> notif.middleware.announcement.sendstate.v1.CreateSendStatesRequest
	6,  // 15: notif.middleware.announcement.sendstate.v1.Middleware.GetSendStates:input_type -> notif.middleware.announcement.sendstate.v1.GetSendStatesRequest
	2,  // 16: notif.middleware.announcement.sendstate.v1.Middleware.CreateSendState:output_type -> notif.middleware.announcement.sendstate.v1.CreateSendStateResponse
	4,  // 17: notif.middleware.announcement.sendstate.v1.Middleware.CreateSendStates:output_type -> notif.middleware.announcement.sendstate.v1.CreateSendStatesResponse
	7,  // 18: notif.middleware.announcement.sendstate.v1.Middleware.GetSendStates:output_type -> notif.middleware.announcement.sendstate.v1.GetSendStatesResponse
	16, // [16:19] is the sub-list for method output_type
	13, // [13:16] is the sub-list for method input_type
	13, // [13:13] is the sub-list for extension type_name
	13, // [13:13] is the sub-list for extension extendee
	0,  // [0:13] is the sub-list for field type_name
}

func init() { file_npool_notif_mw_v1_announcement_sendstate_sendstate_proto_init() }
func file_npool_notif_mw_v1_announcement_sendstate_sendstate_proto_init() {
	if File_npool_notif_mw_v1_announcement_sendstate_sendstate_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_npool_notif_mw_v1_announcement_sendstate_sendstate_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendState); i {
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
		file_npool_notif_mw_v1_announcement_sendstate_sendstate_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateSendStateRequest); i {
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
		file_npool_notif_mw_v1_announcement_sendstate_sendstate_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateSendStateResponse); i {
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
		file_npool_notif_mw_v1_announcement_sendstate_sendstate_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateSendStatesRequest); i {
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
		file_npool_notif_mw_v1_announcement_sendstate_sendstate_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateSendStatesResponse); i {
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
		file_npool_notif_mw_v1_announcement_sendstate_sendstate_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
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
		file_npool_notif_mw_v1_announcement_sendstate_sendstate_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSendStatesRequest); i {
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
		file_npool_notif_mw_v1_announcement_sendstate_sendstate_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSendStatesResponse); i {
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
	file_npool_notif_mw_v1_announcement_sendstate_sendstate_proto_msgTypes[5].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_npool_notif_mw_v1_announcement_sendstate_sendstate_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_npool_notif_mw_v1_announcement_sendstate_sendstate_proto_goTypes,
		DependencyIndexes: file_npool_notif_mw_v1_announcement_sendstate_sendstate_proto_depIdxs,
		MessageInfos:      file_npool_notif_mw_v1_announcement_sendstate_sendstate_proto_msgTypes,
	}.Build()
	File_npool_notif_mw_v1_announcement_sendstate_sendstate_proto = out.File
	file_npool_notif_mw_v1_announcement_sendstate_sendstate_proto_rawDesc = nil
	file_npool_notif_mw_v1_announcement_sendstate_sendstate_proto_goTypes = nil
	file_npool_notif_mw_v1_announcement_sendstate_sendstate_proto_depIdxs = nil
}
