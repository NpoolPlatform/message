// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.18.1
// source: npool/notif/mw/v1/announcement/readstate/readstate.proto

package announcement

import (
	readstate "github.com/NpoolPlatform/message/npool/notif/mgr/v1/announcement/readstate"
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

type ReadState struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: sql:"id"
	AnnouncementID string `protobuf:"bytes,10,opt,name=AnnouncementID,proto3" json:"AnnouncementID,omitempty" sql:"id"`
	// @inject_tag: sql:"app_id"
	AppID string `protobuf:"bytes,20,opt,name=AppID,proto3" json:"AppID,omitempty" sql:"app_id"`
	// @inject_tag: sql:"user_id"
	UserID string `protobuf:"bytes,30,opt,name=UserID,proto3" json:"UserID,omitempty" sql:"user_id"`
	// @inject_tag: sql:"title"
	Title string `protobuf:"bytes,40,opt,name=Title,proto3" json:"Title,omitempty" sql:"title"`
	// @inject_tag: sql:"content"
	Content string `protobuf:"bytes,50,opt,name=Content,proto3" json:"Content,omitempty" sql:"content"`
	// @inject_tag: sql:"channels"
	Channels []channel.NotifChannel `protobuf:"varint,60,rep,packed,name=Channels,proto3,enum=notif.manager.channel.v1.NotifChannel" json:"Channels,omitempty" sql:"channels"`
	// @inject_tag: sql:"email_send"
	EmailSend   bool   `protobuf:"varint,70,opt,name=EmailSend,proto3" json:"EmailSend,omitempty" sql:"email_send"`
	AlreadyRead bool   `protobuf:"varint,80,opt,name=AlreadyRead,proto3" json:"AlreadyRead,omitempty"`
	CreatedAt   uint32 `protobuf:"varint,90,opt,name=CreatedAt,proto3" json:"CreatedAt,omitempty"`
	UpdatedAt   uint32 `protobuf:"varint,100,opt,name=UpdatedAt,proto3" json:"UpdatedAt,omitempty"`
}

func (x *ReadState) Reset() {
	*x = ReadState{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_notif_mw_v1_announcement_readstate_readstate_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadState) ProtoMessage() {}

func (x *ReadState) ProtoReflect() protoreflect.Message {
	mi := &file_npool_notif_mw_v1_announcement_readstate_readstate_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadState.ProtoReflect.Descriptor instead.
func (*ReadState) Descriptor() ([]byte, []int) {
	return file_npool_notif_mw_v1_announcement_readstate_readstate_proto_rawDescGZIP(), []int{0}
}

func (x *ReadState) GetAnnouncementID() string {
	if x != nil {
		return x.AnnouncementID
	}
	return ""
}

func (x *ReadState) GetAppID() string {
	if x != nil {
		return x.AppID
	}
	return ""
}

func (x *ReadState) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

func (x *ReadState) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *ReadState) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *ReadState) GetChannels() []channel.NotifChannel {
	if x != nil {
		return x.Channels
	}
	return nil
}

func (x *ReadState) GetEmailSend() bool {
	if x != nil {
		return x.EmailSend
	}
	return false
}

func (x *ReadState) GetAlreadyRead() bool {
	if x != nil {
		return x.AlreadyRead
	}
	return false
}

func (x *ReadState) GetCreatedAt() uint32 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *ReadState) GetUpdatedAt() uint32 {
	if x != nil {
		return x.UpdatedAt
	}
	return 0
}

type GetReadStateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AnnouncementID string `protobuf:"bytes,10,opt,name=AnnouncementID,proto3" json:"AnnouncementID,omitempty"`
	UserID         string `protobuf:"bytes,20,opt,name=UserID,proto3" json:"UserID,omitempty"`
}

func (x *GetReadStateRequest) Reset() {
	*x = GetReadStateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_notif_mw_v1_announcement_readstate_readstate_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetReadStateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetReadStateRequest) ProtoMessage() {}

func (x *GetReadStateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_notif_mw_v1_announcement_readstate_readstate_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetReadStateRequest.ProtoReflect.Descriptor instead.
func (*GetReadStateRequest) Descriptor() ([]byte, []int) {
	return file_npool_notif_mw_v1_announcement_readstate_readstate_proto_rawDescGZIP(), []int{1}
}

func (x *GetReadStateRequest) GetAnnouncementID() string {
	if x != nil {
		return x.AnnouncementID
	}
	return ""
}

func (x *GetReadStateRequest) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

type GetReadStateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *ReadState `protobuf:"bytes,10,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (x *GetReadStateResponse) Reset() {
	*x = GetReadStateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_notif_mw_v1_announcement_readstate_readstate_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetReadStateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetReadStateResponse) ProtoMessage() {}

func (x *GetReadStateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_notif_mw_v1_announcement_readstate_readstate_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetReadStateResponse.ProtoReflect.Descriptor instead.
func (*GetReadStateResponse) Descriptor() ([]byte, []int) {
	return file_npool_notif_mw_v1_announcement_readstate_readstate_proto_rawDescGZIP(), []int{2}
}

func (x *GetReadStateResponse) GetInfo() *ReadState {
	if x != nil {
		return x.Info
	}
	return nil
}

type GetReadStatesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Conds  *readstate.Conds `protobuf:"bytes,10,opt,name=Conds,proto3" json:"Conds,omitempty"`
	Offset int32            `protobuf:"varint,20,opt,name=Offset,proto3" json:"Offset,omitempty"`
	Limit  int32            `protobuf:"varint,30,opt,name=Limit,proto3" json:"Limit,omitempty"`
}

func (x *GetReadStatesRequest) Reset() {
	*x = GetReadStatesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_notif_mw_v1_announcement_readstate_readstate_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetReadStatesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetReadStatesRequest) ProtoMessage() {}

func (x *GetReadStatesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_notif_mw_v1_announcement_readstate_readstate_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetReadStatesRequest.ProtoReflect.Descriptor instead.
func (*GetReadStatesRequest) Descriptor() ([]byte, []int) {
	return file_npool_notif_mw_v1_announcement_readstate_readstate_proto_rawDescGZIP(), []int{3}
}

func (x *GetReadStatesRequest) GetConds() *readstate.Conds {
	if x != nil {
		return x.Conds
	}
	return nil
}

func (x *GetReadStatesRequest) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *GetReadStatesRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type GetReadStatesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Infos []*ReadState `protobuf:"bytes,10,rep,name=Infos,proto3" json:"Infos,omitempty"`
	Total uint32       `protobuf:"varint,20,opt,name=Total,proto3" json:"Total,omitempty"`
}

func (x *GetReadStatesResponse) Reset() {
	*x = GetReadStatesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_notif_mw_v1_announcement_readstate_readstate_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetReadStatesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetReadStatesResponse) ProtoMessage() {}

func (x *GetReadStatesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_notif_mw_v1_announcement_readstate_readstate_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetReadStatesResponse.ProtoReflect.Descriptor instead.
func (*GetReadStatesResponse) Descriptor() ([]byte, []int) {
	return file_npool_notif_mw_v1_announcement_readstate_readstate_proto_rawDescGZIP(), []int{4}
}

func (x *GetReadStatesResponse) GetInfos() []*ReadState {
	if x != nil {
		return x.Infos
	}
	return nil
}

func (x *GetReadStatesResponse) GetTotal() uint32 {
	if x != nil {
		return x.Total
	}
	return 0
}

var File_npool_notif_mw_v1_announcement_readstate_readstate_proto protoreflect.FileDescriptor

var file_npool_notif_mw_v1_announcement_readstate_readstate_proto_rawDesc = []byte{
	0x0a, 0x38, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x2f, 0x6d, 0x77,
	0x2f, 0x76, 0x31, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x75, 0x6e, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74,
	0x2f, 0x72, 0x65, 0x61, 0x64, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2f, 0x72, 0x65, 0x61, 0x64, 0x73,
	0x74, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x20, 0x6e, 0x6f, 0x74, 0x69,
	0x66, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x61, 0x6e, 0x6e,
	0x6f, 0x75, 0x6e, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x1a, 0x39, 0x6e, 0x70,
	0x6f, 0x6f, 0x6c, 0x2f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x2f, 0x6d, 0x67, 0x72, 0x2f, 0x76, 0x31,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x75, 0x6e, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x72, 0x65,
	0x61, 0x64, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2f, 0x72, 0x65, 0x61, 0x64, 0x73, 0x74, 0x61, 0x74,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x28, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x6e,
	0x6f, 0x74, 0x69, 0x66, 0x2f, 0x6d, 0x67, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x68, 0x61, 0x6e,
	0x6e, 0x65, 0x6c, 0x2f, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xd1, 0x02, 0x0a, 0x09, 0x52, 0x65, 0x61, 0x64, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12,
	0x26, 0x0a, 0x0e, 0x41, 0x6e, 0x6e, 0x6f, 0x75, 0x6e, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x49,
	0x44, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x41, 0x6e, 0x6e, 0x6f, 0x75, 0x6e, 0x63,
	0x65, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x41, 0x70, 0x70, 0x49, 0x44,
	0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x41, 0x70, 0x70, 0x49, 0x44, 0x12, 0x16, 0x0a,
	0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x55,
	0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x28,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x43,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x32, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x43, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x42, 0x0a, 0x08, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c,
	0x73, 0x18, 0x3c, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x26, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x2e,
	0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x2e,
	0x76, 0x31, 0x2e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x52,
	0x08, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x45, 0x6d, 0x61,
	0x69, 0x6c, 0x53, 0x65, 0x6e, 0x64, 0x18, 0x46, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x45, 0x6d,
	0x61, 0x69, 0x6c, 0x53, 0x65, 0x6e, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x41, 0x6c, 0x72, 0x65, 0x61,
	0x64, 0x79, 0x52, 0x65, 0x61, 0x64, 0x18, 0x50, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x41, 0x6c,
	0x72, 0x65, 0x61, 0x64, 0x79, 0x52, 0x65, 0x61, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x5a, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x18, 0x64, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x55, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x52, 0x65, 0x61, 0x64,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x26, 0x0a, 0x0e,
	0x41, 0x6e, 0x6e, 0x6f, 0x75, 0x6e, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x41, 0x6e, 0x6e, 0x6f, 0x75, 0x6e, 0x63, 0x65, 0x6d, 0x65,
	0x6e, 0x74, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x14,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x22, 0x57, 0x0a, 0x14,
	0x47, 0x65, 0x74, 0x52, 0x65, 0x61, 0x64, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3f, 0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c,
	0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x61, 0x6e, 0x6e, 0x6f, 0x75, 0x6e, 0x63, 0x65, 0x6d, 0x65,
	0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52,
	0x04, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x7d, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x52, 0x65, 0x61, 0x64,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x37, 0x0a,
	0x05, 0x43, 0x6f, 0x6e, 0x64, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x6e,
	0x6f, 0x74, 0x69, 0x66, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x72, 0x65, 0x61,
	0x64, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x64, 0x73, 0x52,
	0x05, 0x43, 0x6f, 0x6e, 0x64, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74,
	0x18, 0x14, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x14,
	0x0a, 0x05, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x4c,
	0x69, 0x6d, 0x69, 0x74, 0x22, 0x70, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x52, 0x65, 0x61, 0x64, 0x53,
	0x74, 0x61, 0x74, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x41, 0x0a,
	0x05, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x6e,
	0x6f, 0x74, 0x69, 0x66, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e,
	0x61, 0x6e, 0x6e, 0x6f, 0x75, 0x6e, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e,
	0x52, 0x65, 0x61, 0x64, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x49, 0x6e, 0x66, 0x6f, 0x73,
	0x12, 0x14, 0x0a, 0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x14, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x32, 0x92, 0x02, 0x0a, 0x0a, 0x4d, 0x69, 0x64, 0x64, 0x6c,
	0x65, 0x77, 0x61, 0x72, 0x65, 0x12, 0x7f, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x52, 0x65, 0x61, 0x64,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x35, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x2e, 0x6d, 0x69,
	0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x61, 0x6e, 0x6e, 0x6f, 0x75, 0x6e, 0x63,
	0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x61, 0x64,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x36, 0x2e, 0x6e,
	0x6f, 0x74, 0x69, 0x66, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e,
	0x61, 0x6e, 0x6e, 0x6f, 0x75, 0x6e, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e,
	0x47, 0x65, 0x74, 0x52, 0x65, 0x61, 0x64, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x82, 0x01, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x52, 0x65,
	0x61, 0x64, 0x53, 0x74, 0x61, 0x74, 0x65, 0x73, 0x12, 0x36, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66,
	0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x61, 0x6e, 0x6e, 0x6f,
	0x75, 0x6e, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x52,
	0x65, 0x61, 0x64, 0x53, 0x74, 0x61, 0x74, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x37, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77,
	0x61, 0x72, 0x65, 0x2e, 0x61, 0x6e, 0x6e, 0x6f, 0x75, 0x6e, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74,
	0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x61, 0x64, 0x53, 0x74, 0x61, 0x74, 0x65,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x41, 0x5a, 0x3f, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4e, 0x70, 0x6f, 0x6f, 0x6c, 0x50,
	0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f,
	0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x2f, 0x6d, 0x77, 0x2f, 0x76,
	0x31, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x75, 0x6e, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_npool_notif_mw_v1_announcement_readstate_readstate_proto_rawDescOnce sync.Once
	file_npool_notif_mw_v1_announcement_readstate_readstate_proto_rawDescData = file_npool_notif_mw_v1_announcement_readstate_readstate_proto_rawDesc
)

func file_npool_notif_mw_v1_announcement_readstate_readstate_proto_rawDescGZIP() []byte {
	file_npool_notif_mw_v1_announcement_readstate_readstate_proto_rawDescOnce.Do(func() {
		file_npool_notif_mw_v1_announcement_readstate_readstate_proto_rawDescData = protoimpl.X.CompressGZIP(file_npool_notif_mw_v1_announcement_readstate_readstate_proto_rawDescData)
	})
	return file_npool_notif_mw_v1_announcement_readstate_readstate_proto_rawDescData
}

var file_npool_notif_mw_v1_announcement_readstate_readstate_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_npool_notif_mw_v1_announcement_readstate_readstate_proto_goTypes = []interface{}{
	(*ReadState)(nil),             // 0: notif.middleware.announcement.v1.ReadState
	(*GetReadStateRequest)(nil),   // 1: notif.middleware.announcement.v1.GetReadStateRequest
	(*GetReadStateResponse)(nil),  // 2: notif.middleware.announcement.v1.GetReadStateResponse
	(*GetReadStatesRequest)(nil),  // 3: notif.middleware.announcement.v1.GetReadStatesRequest
	(*GetReadStatesResponse)(nil), // 4: notif.middleware.announcement.v1.GetReadStatesResponse
	(channel.NotifChannel)(0),     // 5: notif.manager.channel.v1.NotifChannel
	(*readstate.Conds)(nil),       // 6: notif.manager.readstate.v1.Conds
}
var file_npool_notif_mw_v1_announcement_readstate_readstate_proto_depIdxs = []int32{
	5, // 0: notif.middleware.announcement.v1.ReadState.Channels:type_name -> notif.manager.channel.v1.NotifChannel
	0, // 1: notif.middleware.announcement.v1.GetReadStateResponse.Info:type_name -> notif.middleware.announcement.v1.ReadState
	6, // 2: notif.middleware.announcement.v1.GetReadStatesRequest.Conds:type_name -> notif.manager.readstate.v1.Conds
	0, // 3: notif.middleware.announcement.v1.GetReadStatesResponse.Infos:type_name -> notif.middleware.announcement.v1.ReadState
	1, // 4: notif.middleware.announcement.v1.Middleware.GetReadState:input_type -> notif.middleware.announcement.v1.GetReadStateRequest
	3, // 5: notif.middleware.announcement.v1.Middleware.GetReadStates:input_type -> notif.middleware.announcement.v1.GetReadStatesRequest
	2, // 6: notif.middleware.announcement.v1.Middleware.GetReadState:output_type -> notif.middleware.announcement.v1.GetReadStateResponse
	4, // 7: notif.middleware.announcement.v1.Middleware.GetReadStates:output_type -> notif.middleware.announcement.v1.GetReadStatesResponse
	6, // [6:8] is the sub-list for method output_type
	4, // [4:6] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_npool_notif_mw_v1_announcement_readstate_readstate_proto_init() }
func file_npool_notif_mw_v1_announcement_readstate_readstate_proto_init() {
	if File_npool_notif_mw_v1_announcement_readstate_readstate_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_npool_notif_mw_v1_announcement_readstate_readstate_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadState); i {
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
		file_npool_notif_mw_v1_announcement_readstate_readstate_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetReadStateRequest); i {
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
		file_npool_notif_mw_v1_announcement_readstate_readstate_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetReadStateResponse); i {
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
		file_npool_notif_mw_v1_announcement_readstate_readstate_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetReadStatesRequest); i {
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
		file_npool_notif_mw_v1_announcement_readstate_readstate_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetReadStatesResponse); i {
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
			RawDescriptor: file_npool_notif_mw_v1_announcement_readstate_readstate_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_npool_notif_mw_v1_announcement_readstate_readstate_proto_goTypes,
		DependencyIndexes: file_npool_notif_mw_v1_announcement_readstate_readstate_proto_depIdxs,
		MessageInfos:      file_npool_notif_mw_v1_announcement_readstate_readstate_proto_msgTypes,
	}.Build()
	File_npool_notif_mw_v1_announcement_readstate_readstate_proto = out.File
	file_npool_notif_mw_v1_announcement_readstate_readstate_proto_rawDesc = nil
	file_npool_notif_mw_v1_announcement_readstate_readstate_proto_goTypes = nil
	file_npool_notif_mw_v1_announcement_readstate_readstate_proto_depIdxs = nil
}
