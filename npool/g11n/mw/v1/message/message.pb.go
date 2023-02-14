// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.18.1
// source: npool/g11n/mw/v1/message/message.proto

package message

import (
	_ "github.com/NpoolPlatform/message/npool"
	message "github.com/NpoolPlatform/message/npool/g11n/mgr/v1/message"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: sql:"id"
	ID string `protobuf:"bytes,10,opt,name=ID,proto3" json:"ID,omitempty" sql:"id"`
	// @inject_tag: sql:"app_id"
	AppID string `protobuf:"bytes,20,opt,name=AppID,proto3" json:"AppID,omitempty" sql:"app_id"`
	// @inject_tag: sql:"lang_id"
	LangID string `protobuf:"bytes,30,opt,name=LangID,proto3" json:"LangID,omitempty" sql:"lang_id"`
	// @inject_tag: sql:"lang"
	Lang string `protobuf:"bytes,40,opt,name=Lang,proto3" json:"Lang,omitempty" sql:"lang"`
	// @inject_tag: sql:"message_id"
	MessageID string `protobuf:"bytes,50,opt,name=MessageID,proto3" json:"MessageID,omitempty" sql:"message_id"`
	// @inject_tag: sql:"message"
	Message string `protobuf:"bytes,60,opt,name=Message,proto3" json:"Message,omitempty" sql:"message"`
	// @inject_tag: sql:"get_index"
	GetIndex uint32 `protobuf:"varint,70,opt,name=GetIndex,proto3" json:"GetIndex,omitempty" sql:"get_index"`
	// @inject_tag: sql:"disabled"
	Disabled bool `protobuf:"varint,80,opt,name=Disabled,proto3" json:"Disabled,omitempty" sql:"disabled"`
	// @inject_tag: sql:"created_at"
	CreatedAt uint32 `protobuf:"varint,90,opt,name=CreatedAt,proto3" json:"CreatedAt,omitempty" sql:"created_at"`
	// @inject_tag: sql:"updated_at"
	UpdatedAt uint32 `protobuf:"varint,100,opt,name=UpdatedAt,proto3" json:"UpdatedAt,omitempty" sql:"updated_at"`
}

func (x *Message) Reset() {
	*x = Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_g11n_mw_v1_message_message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_npool_g11n_mw_v1_message_message_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message.ProtoReflect.Descriptor instead.
func (*Message) Descriptor() ([]byte, []int) {
	return file_npool_g11n_mw_v1_message_message_proto_rawDescGZIP(), []int{0}
}

func (x *Message) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *Message) GetAppID() string {
	if x != nil {
		return x.AppID
	}
	return ""
}

func (x *Message) GetLangID() string {
	if x != nil {
		return x.LangID
	}
	return ""
}

func (x *Message) GetLang() string {
	if x != nil {
		return x.Lang
	}
	return ""
}

func (x *Message) GetMessageID() string {
	if x != nil {
		return x.MessageID
	}
	return ""
}

func (x *Message) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *Message) GetGetIndex() uint32 {
	if x != nil {
		return x.GetIndex
	}
	return 0
}

func (x *Message) GetDisabled() bool {
	if x != nil {
		return x.Disabled
	}
	return false
}

func (x *Message) GetCreatedAt() uint32 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *Message) GetUpdatedAt() uint32 {
	if x != nil {
		return x.UpdatedAt
	}
	return 0
}

type CreateMessageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *message.MessageReq `protobuf:"bytes,10,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (x *CreateMessageRequest) Reset() {
	*x = CreateMessageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_g11n_mw_v1_message_message_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateMessageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateMessageRequest) ProtoMessage() {}

func (x *CreateMessageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_g11n_mw_v1_message_message_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateMessageRequest.ProtoReflect.Descriptor instead.
func (*CreateMessageRequest) Descriptor() ([]byte, []int) {
	return file_npool_g11n_mw_v1_message_message_proto_rawDescGZIP(), []int{1}
}

func (x *CreateMessageRequest) GetInfo() *message.MessageReq {
	if x != nil {
		return x.Info
	}
	return nil
}

type CreateMessageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *Message `protobuf:"bytes,10,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (x *CreateMessageResponse) Reset() {
	*x = CreateMessageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_g11n_mw_v1_message_message_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateMessageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateMessageResponse) ProtoMessage() {}

func (x *CreateMessageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_g11n_mw_v1_message_message_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateMessageResponse.ProtoReflect.Descriptor instead.
func (*CreateMessageResponse) Descriptor() ([]byte, []int) {
	return file_npool_g11n_mw_v1_message_message_proto_rawDescGZIP(), []int{2}
}

func (x *CreateMessageResponse) GetInfo() *Message {
	if x != nil {
		return x.Info
	}
	return nil
}

type CreateMessagesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Infos []*message.MessageReq `protobuf:"bytes,10,rep,name=Infos,proto3" json:"Infos,omitempty"`
}

func (x *CreateMessagesRequest) Reset() {
	*x = CreateMessagesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_g11n_mw_v1_message_message_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateMessagesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateMessagesRequest) ProtoMessage() {}

func (x *CreateMessagesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_g11n_mw_v1_message_message_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateMessagesRequest.ProtoReflect.Descriptor instead.
func (*CreateMessagesRequest) Descriptor() ([]byte, []int) {
	return file_npool_g11n_mw_v1_message_message_proto_rawDescGZIP(), []int{3}
}

func (x *CreateMessagesRequest) GetInfos() []*message.MessageReq {
	if x != nil {
		return x.Infos
	}
	return nil
}

type CreateMessagesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Infos []*Message `protobuf:"bytes,10,rep,name=Infos,proto3" json:"Infos,omitempty"`
}

func (x *CreateMessagesResponse) Reset() {
	*x = CreateMessagesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_g11n_mw_v1_message_message_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateMessagesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateMessagesResponse) ProtoMessage() {}

func (x *CreateMessagesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_g11n_mw_v1_message_message_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateMessagesResponse.ProtoReflect.Descriptor instead.
func (*CreateMessagesResponse) Descriptor() ([]byte, []int) {
	return file_npool_g11n_mw_v1_message_message_proto_rawDescGZIP(), []int{4}
}

func (x *CreateMessagesResponse) GetInfos() []*Message {
	if x != nil {
		return x.Infos
	}
	return nil
}

type UpdateMessageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *message.MessageReq `protobuf:"bytes,10,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (x *UpdateMessageRequest) Reset() {
	*x = UpdateMessageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_g11n_mw_v1_message_message_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateMessageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateMessageRequest) ProtoMessage() {}

func (x *UpdateMessageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_g11n_mw_v1_message_message_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateMessageRequest.ProtoReflect.Descriptor instead.
func (*UpdateMessageRequest) Descriptor() ([]byte, []int) {
	return file_npool_g11n_mw_v1_message_message_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateMessageRequest) GetInfo() *message.MessageReq {
	if x != nil {
		return x.Info
	}
	return nil
}

type UpdateMessageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *Message `protobuf:"bytes,10,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (x *UpdateMessageResponse) Reset() {
	*x = UpdateMessageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_g11n_mw_v1_message_message_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateMessageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateMessageResponse) ProtoMessage() {}

func (x *UpdateMessageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_g11n_mw_v1_message_message_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateMessageResponse.ProtoReflect.Descriptor instead.
func (*UpdateMessageResponse) Descriptor() ([]byte, []int) {
	return file_npool_g11n_mw_v1_message_message_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateMessageResponse) GetInfo() *Message {
	if x != nil {
		return x.Info
	}
	return nil
}

type GetMessagesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Conds  *message.Conds `protobuf:"bytes,10,opt,name=Conds,proto3" json:"Conds,omitempty"`
	Offset int32          `protobuf:"varint,20,opt,name=Offset,proto3" json:"Offset,omitempty"`
	Limit  int32          `protobuf:"varint,30,opt,name=Limit,proto3" json:"Limit,omitempty"`
}

func (x *GetMessagesRequest) Reset() {
	*x = GetMessagesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_g11n_mw_v1_message_message_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMessagesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMessagesRequest) ProtoMessage() {}

func (x *GetMessagesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_g11n_mw_v1_message_message_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMessagesRequest.ProtoReflect.Descriptor instead.
func (*GetMessagesRequest) Descriptor() ([]byte, []int) {
	return file_npool_g11n_mw_v1_message_message_proto_rawDescGZIP(), []int{7}
}

func (x *GetMessagesRequest) GetConds() *message.Conds {
	if x != nil {
		return x.Conds
	}
	return nil
}

func (x *GetMessagesRequest) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *GetMessagesRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type GetMessagesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Infos []*Message `protobuf:"bytes,10,rep,name=Infos,proto3" json:"Infos,omitempty"`
	Total uint32     `protobuf:"varint,20,opt,name=Total,proto3" json:"Total,omitempty"`
}

func (x *GetMessagesResponse) Reset() {
	*x = GetMessagesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_g11n_mw_v1_message_message_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMessagesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMessagesResponse) ProtoMessage() {}

func (x *GetMessagesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_g11n_mw_v1_message_message_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMessagesResponse.ProtoReflect.Descriptor instead.
func (*GetMessagesResponse) Descriptor() ([]byte, []int) {
	return file_npool_g11n_mw_v1_message_message_proto_rawDescGZIP(), []int{8}
}

func (x *GetMessagesResponse) GetInfos() []*Message {
	if x != nil {
		return x.Infos
	}
	return nil
}

func (x *GetMessagesResponse) GetTotal() uint32 {
	if x != nil {
		return x.Total
	}
	return 0
}

type DeleteMessageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID string `protobuf:"bytes,10,opt,name=ID,proto3" json:"ID,omitempty"`
}

func (x *DeleteMessageRequest) Reset() {
	*x = DeleteMessageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_g11n_mw_v1_message_message_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteMessageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteMessageRequest) ProtoMessage() {}

func (x *DeleteMessageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_g11n_mw_v1_message_message_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteMessageRequest.ProtoReflect.Descriptor instead.
func (*DeleteMessageRequest) Descriptor() ([]byte, []int) {
	return file_npool_g11n_mw_v1_message_message_proto_rawDescGZIP(), []int{9}
}

func (x *DeleteMessageRequest) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

type DeleteMessageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *Message `protobuf:"bytes,10,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (x *DeleteMessageResponse) Reset() {
	*x = DeleteMessageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_g11n_mw_v1_message_message_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteMessageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteMessageResponse) ProtoMessage() {}

func (x *DeleteMessageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_g11n_mw_v1_message_message_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteMessageResponse.ProtoReflect.Descriptor instead.
func (*DeleteMessageResponse) Descriptor() ([]byte, []int) {
	return file_npool_g11n_mw_v1_message_message_proto_rawDescGZIP(), []int{10}
}

func (x *DeleteMessageResponse) GetInfo() *Message {
	if x != nil {
		return x.Info
	}
	return nil
}

var File_npool_g11n_mw_v1_message_message_proto protoreflect.FileDescriptor

var file_npool_g11n_mw_v1_message_message_proto_rawDesc = []byte{
	0x0a, 0x26, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x67, 0x31, 0x31, 0x6e, 0x2f, 0x6d, 0x77, 0x2f,
	0x76, 0x31, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1b, 0x67, 0x31, 0x31, 0x6e, 0x2e, 0x6d,
	0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x31, 0x2e, 0x76, 0x31, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x11, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x27, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x67, 0x31, 0x31,
	0x6e, 0x2f, 0x6d, 0x67, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x87,
	0x02, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x41, 0x70,
	0x70, 0x49, 0x44, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x41, 0x70, 0x70, 0x49, 0x44,
	0x12, 0x16, 0x0a, 0x06, 0x4c, 0x61, 0x6e, 0x67, 0x49, 0x44, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x4c, 0x61, 0x6e, 0x67, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x4c, 0x61, 0x6e, 0x67,
	0x18, 0x28, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4c, 0x61, 0x6e, 0x67, 0x12, 0x1c, 0x0a, 0x09,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x49, 0x44, 0x18, 0x32, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x49, 0x44, 0x12, 0x18, 0x0a, 0x07, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x3c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x64, 0x65, 0x78,
	0x18, 0x46, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x64, 0x65, 0x78,
	0x12, 0x1a, 0x0a, 0x08, 0x44, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x50, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x08, 0x44, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x1c, 0x0a, 0x09,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x5a, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x64, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x50, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x38, 0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24,
	0x2e, 0x67, 0x31, 0x31, 0x6e, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x31, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x52, 0x65, 0x71, 0x52, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x51, 0x0a, 0x15, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x38, 0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x24, 0x2e, 0x67, 0x31, 0x31, 0x6e, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77,
	0x61, 0x72, 0x65, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x31, 0x2e, 0x76, 0x31, 0x2e,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x53, 0x0a,
	0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3a, 0x0a, 0x05, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x18,
	0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x67, 0x31, 0x31, 0x6e, 0x2e, 0x6d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x72, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x31, 0x2e, 0x76, 0x31,
	0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x52, 0x05, 0x49, 0x6e, 0x66,
	0x6f, 0x73, 0x22, 0x54, 0x0a, 0x16, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3a, 0x0a, 0x05,
	0x49, 0x6e, 0x66, 0x6f, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x67, 0x31,
	0x31, 0x6e, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x31, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x52, 0x05, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x22, 0x50, 0x0a, 0x14, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x38, 0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24,
	0x2e, 0x67, 0x31, 0x31, 0x6e, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x31, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x52, 0x65, 0x71, 0x52, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x51, 0x0a, 0x15, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x38, 0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x24, 0x2e, 0x67, 0x31, 0x31, 0x6e, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77,
	0x61, 0x72, 0x65, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x31, 0x2e, 0x76, 0x31, 0x2e,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x79, 0x0a,
	0x12, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x35, 0x0a, 0x05, 0x43, 0x6f, 0x6e, 0x64, 0x73, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x67, 0x31, 0x31, 0x6e, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65,
	0x72, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x31, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f,
	0x6e, 0x64, 0x73, 0x52, 0x05, 0x43, 0x6f, 0x6e, 0x64, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x4f, 0x66,
	0x66, 0x73, 0x65, 0x74, 0x18, 0x14, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x4f, 0x66, 0x66, 0x73,
	0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x1e, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x05, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x22, 0x67, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x3a, 0x0a, 0x05, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24,
	0x2e, 0x67, 0x31, 0x31, 0x6e, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65,
	0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x31, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x52, 0x05, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x54,
	0x6f, 0x74, 0x61, 0x6c, 0x18, 0x14, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x54, 0x6f, 0x74, 0x61,
	0x6c, 0x22, 0x26, 0x0a, 0x14, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x22, 0x51, 0x0a, 0x15, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x38, 0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x24, 0x2e, 0x67, 0x31, 0x31, 0x6e, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61,
	0x72, 0x65, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x31, 0x2e, 0x76, 0x31, 0x2e, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x32, 0xeb, 0x04, 0x0a,
	0x0a, 0x4d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x12, 0x78, 0x0a, 0x0d, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x31, 0x2e, 0x67,
	0x31, 0x31, 0x6e, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x31, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x32, 0x2e, 0x67, 0x31, 0x31, 0x6e, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72,
	0x65, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x31, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x7b, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x12, 0x32, 0x2e, 0x67, 0x31, 0x31, 0x6e, 0x2e, 0x6d,
	0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x31, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x33, 0x2e, 0x67, 0x31,
	0x31, 0x6e, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x31, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x78, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x12, 0x31, 0x2e, 0x67, 0x31, 0x31, 0x6e, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c,
	0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x31, 0x2e, 0x76,
	0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x32, 0x2e, 0x67, 0x31, 0x31, 0x6e, 0x2e, 0x6d, 0x69,
	0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x31, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x72, 0x0a, 0x0b,
	0x47, 0x65, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x12, 0x2f, 0x2e, 0x67, 0x31,
	0x31, 0x6e, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x31, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x30, 0x2e, 0x67,
	0x31, 0x31, 0x6e, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x31, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x78, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x12, 0x31, 0x2e, 0x67, 0x31, 0x31, 0x6e, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77,
	0x61, 0x72, 0x65, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x31, 0x2e, 0x76, 0x31, 0x2e,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x32, 0x2e, 0x67, 0x31, 0x31, 0x6e, 0x2e, 0x6d, 0x69, 0x64, 0x64,
	0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x31, 0x2e,
	0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x3b, 0x5a, 0x39, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4e, 0x70, 0x6f, 0x6f, 0x6c, 0x50, 0x6c,
	0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x6e,
	0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x67, 0x31, 0x31, 0x6e, 0x2f, 0x6d, 0x77, 0x2f, 0x76, 0x31, 0x2f,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_npool_g11n_mw_v1_message_message_proto_rawDescOnce sync.Once
	file_npool_g11n_mw_v1_message_message_proto_rawDescData = file_npool_g11n_mw_v1_message_message_proto_rawDesc
)

func file_npool_g11n_mw_v1_message_message_proto_rawDescGZIP() []byte {
	file_npool_g11n_mw_v1_message_message_proto_rawDescOnce.Do(func() {
		file_npool_g11n_mw_v1_message_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_npool_g11n_mw_v1_message_message_proto_rawDescData)
	})
	return file_npool_g11n_mw_v1_message_message_proto_rawDescData
}

var file_npool_g11n_mw_v1_message_message_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_npool_g11n_mw_v1_message_message_proto_goTypes = []interface{}{
	(*Message)(nil),                // 0: g11n.middleware.message1.v1.Message
	(*CreateMessageRequest)(nil),   // 1: g11n.middleware.message1.v1.CreateMessageRequest
	(*CreateMessageResponse)(nil),  // 2: g11n.middleware.message1.v1.CreateMessageResponse
	(*CreateMessagesRequest)(nil),  // 3: g11n.middleware.message1.v1.CreateMessagesRequest
	(*CreateMessagesResponse)(nil), // 4: g11n.middleware.message1.v1.CreateMessagesResponse
	(*UpdateMessageRequest)(nil),   // 5: g11n.middleware.message1.v1.UpdateMessageRequest
	(*UpdateMessageResponse)(nil),  // 6: g11n.middleware.message1.v1.UpdateMessageResponse
	(*GetMessagesRequest)(nil),     // 7: g11n.middleware.message1.v1.GetMessagesRequest
	(*GetMessagesResponse)(nil),    // 8: g11n.middleware.message1.v1.GetMessagesResponse
	(*DeleteMessageRequest)(nil),   // 9: g11n.middleware.message1.v1.DeleteMessageRequest
	(*DeleteMessageResponse)(nil),  // 10: g11n.middleware.message1.v1.DeleteMessageResponse
	(*message.MessageReq)(nil),     // 11: g11n.manager.message1.v1.MessageReq
	(*message.Conds)(nil),          // 12: g11n.manager.message1.v1.Conds
}
var file_npool_g11n_mw_v1_message_message_proto_depIdxs = []int32{
	11, // 0: g11n.middleware.message1.v1.CreateMessageRequest.Info:type_name -> g11n.manager.message1.v1.MessageReq
	0,  // 1: g11n.middleware.message1.v1.CreateMessageResponse.Info:type_name -> g11n.middleware.message1.v1.Message
	11, // 2: g11n.middleware.message1.v1.CreateMessagesRequest.Infos:type_name -> g11n.manager.message1.v1.MessageReq
	0,  // 3: g11n.middleware.message1.v1.CreateMessagesResponse.Infos:type_name -> g11n.middleware.message1.v1.Message
	11, // 4: g11n.middleware.message1.v1.UpdateMessageRequest.Info:type_name -> g11n.manager.message1.v1.MessageReq
	0,  // 5: g11n.middleware.message1.v1.UpdateMessageResponse.Info:type_name -> g11n.middleware.message1.v1.Message
	12, // 6: g11n.middleware.message1.v1.GetMessagesRequest.Conds:type_name -> g11n.manager.message1.v1.Conds
	0,  // 7: g11n.middleware.message1.v1.GetMessagesResponse.Infos:type_name -> g11n.middleware.message1.v1.Message
	0,  // 8: g11n.middleware.message1.v1.DeleteMessageResponse.Info:type_name -> g11n.middleware.message1.v1.Message
	1,  // 9: g11n.middleware.message1.v1.Middleware.CreateMessage:input_type -> g11n.middleware.message1.v1.CreateMessageRequest
	3,  // 10: g11n.middleware.message1.v1.Middleware.CreateMessages:input_type -> g11n.middleware.message1.v1.CreateMessagesRequest
	5,  // 11: g11n.middleware.message1.v1.Middleware.UpdateMessage:input_type -> g11n.middleware.message1.v1.UpdateMessageRequest
	7,  // 12: g11n.middleware.message1.v1.Middleware.GetMessages:input_type -> g11n.middleware.message1.v1.GetMessagesRequest
	9,  // 13: g11n.middleware.message1.v1.Middleware.DeleteMessage:input_type -> g11n.middleware.message1.v1.DeleteMessageRequest
	2,  // 14: g11n.middleware.message1.v1.Middleware.CreateMessage:output_type -> g11n.middleware.message1.v1.CreateMessageResponse
	4,  // 15: g11n.middleware.message1.v1.Middleware.CreateMessages:output_type -> g11n.middleware.message1.v1.CreateMessagesResponse
	6,  // 16: g11n.middleware.message1.v1.Middleware.UpdateMessage:output_type -> g11n.middleware.message1.v1.UpdateMessageResponse
	8,  // 17: g11n.middleware.message1.v1.Middleware.GetMessages:output_type -> g11n.middleware.message1.v1.GetMessagesResponse
	10, // 18: g11n.middleware.message1.v1.Middleware.DeleteMessage:output_type -> g11n.middleware.message1.v1.DeleteMessageResponse
	14, // [14:19] is the sub-list for method output_type
	9,  // [9:14] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_npool_g11n_mw_v1_message_message_proto_init() }
func file_npool_g11n_mw_v1_message_message_proto_init() {
	if File_npool_g11n_mw_v1_message_message_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_npool_g11n_mw_v1_message_message_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Message); i {
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
		file_npool_g11n_mw_v1_message_message_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateMessageRequest); i {
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
		file_npool_g11n_mw_v1_message_message_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateMessageResponse); i {
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
		file_npool_g11n_mw_v1_message_message_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateMessagesRequest); i {
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
		file_npool_g11n_mw_v1_message_message_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateMessagesResponse); i {
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
		file_npool_g11n_mw_v1_message_message_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateMessageRequest); i {
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
		file_npool_g11n_mw_v1_message_message_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateMessageResponse); i {
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
		file_npool_g11n_mw_v1_message_message_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMessagesRequest); i {
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
		file_npool_g11n_mw_v1_message_message_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMessagesResponse); i {
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
		file_npool_g11n_mw_v1_message_message_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteMessageRequest); i {
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
		file_npool_g11n_mw_v1_message_message_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteMessageResponse); i {
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
			RawDescriptor: file_npool_g11n_mw_v1_message_message_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_npool_g11n_mw_v1_message_message_proto_goTypes,
		DependencyIndexes: file_npool_g11n_mw_v1_message_message_proto_depIdxs,
		MessageInfos:      file_npool_g11n_mw_v1_message_message_proto_msgTypes,
	}.Build()
	File_npool_g11n_mw_v1_message_message_proto = out.File
	file_npool_g11n_mw_v1_message_message_proto_rawDesc = nil
	file_npool_g11n_mw_v1_message_message_proto_goTypes = nil
	file_npool_g11n_mw_v1_message_message_proto_depIdxs = nil
}
