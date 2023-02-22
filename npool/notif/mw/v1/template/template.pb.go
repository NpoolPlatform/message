// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.18.1
// source: npool/notif/mw/v1/template/template.proto

package template

import (
	v1 "github.com/NpoolPlatform/message/npool/basetypes/v1"
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

type TemplateVars struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username  *string `protobuf:"bytes,10,opt,name=Username,proto3,oneof" json:"Username,omitempty"`
	Message   *string `protobuf:"bytes,20,opt,name=Message,proto3,oneof" json:"Message,omitempty"`
	Amount    *string `protobuf:"bytes,30,opt,name=Amount,proto3,oneof" json:"Amount,omitempty"`
	CoinUnit  *string `protobuf:"bytes,40,opt,name=CoinUnit,proto3,oneof" json:"CoinUnit,omitempty"`
	Timestamp *uint32 `protobuf:"varint,50,opt,name=Timestamp,proto3,oneof" json:"Timestamp,omitempty"`
	Address   *string `protobuf:"bytes,60,opt,name=Address,proto3,oneof" json:"Address,omitempty"`
	Code      *string `protobuf:"bytes,70,opt,name=Code,proto3,oneof" json:"Code,omitempty"`
}

func (x *TemplateVars) Reset() {
	*x = TemplateVars{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_notif_mw_v1_template_template_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TemplateVars) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TemplateVars) ProtoMessage() {}

func (x *TemplateVars) ProtoReflect() protoreflect.Message {
	mi := &file_npool_notif_mw_v1_template_template_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TemplateVars.ProtoReflect.Descriptor instead.
func (*TemplateVars) Descriptor() ([]byte, []int) {
	return file_npool_notif_mw_v1_template_template_proto_rawDescGZIP(), []int{0}
}

func (x *TemplateVars) GetUsername() string {
	if x != nil && x.Username != nil {
		return *x.Username
	}
	return ""
}

func (x *TemplateVars) GetMessage() string {
	if x != nil && x.Message != nil {
		return *x.Message
	}
	return ""
}

func (x *TemplateVars) GetAmount() string {
	if x != nil && x.Amount != nil {
		return *x.Amount
	}
	return ""
}

func (x *TemplateVars) GetCoinUnit() string {
	if x != nil && x.CoinUnit != nil {
		return *x.CoinUnit
	}
	return ""
}

func (x *TemplateVars) GetTimestamp() uint32 {
	if x != nil && x.Timestamp != nil {
		return *x.Timestamp
	}
	return 0
}

func (x *TemplateVars) GetAddress() string {
	if x != nil && x.Address != nil {
		return *x.Address
	}
	return ""
}

func (x *TemplateVars) GetCode() string {
	if x != nil && x.Code != nil {
		return *x.Code
	}
	return ""
}

type TextInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Subject  string   `protobuf:"bytes,10,opt,name=Subject,proto3" json:"Subject,omitempty"`
	Content  string   `protobuf:"bytes,20,opt,name=Content,proto3" json:"Content,omitempty"`
	From     string   `protobuf:"bytes,30,opt,name=From,proto3" json:"From,omitempty"`
	ToCCs    []string `protobuf:"bytes,40,rep,name=ToCCs,proto3" json:"ToCCs,omitempty"`
	ReplyTos []string `protobuf:"bytes,50,rep,name=ReplyTos,proto3" json:"ReplyTos,omitempty"`
}

func (x *TextInfo) Reset() {
	*x = TextInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_notif_mw_v1_template_template_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TextInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TextInfo) ProtoMessage() {}

func (x *TextInfo) ProtoReflect() protoreflect.Message {
	mi := &file_npool_notif_mw_v1_template_template_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TextInfo.ProtoReflect.Descriptor instead.
func (*TextInfo) Descriptor() ([]byte, []int) {
	return file_npool_notif_mw_v1_template_template_proto_rawDescGZIP(), []int{1}
}

func (x *TextInfo) GetSubject() string {
	if x != nil {
		return x.Subject
	}
	return ""
}

func (x *TextInfo) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *TextInfo) GetFrom() string {
	if x != nil {
		return x.From
	}
	return ""
}

func (x *TextInfo) GetToCCs() []string {
	if x != nil {
		return x.ToCCs
	}
	return nil
}

func (x *TextInfo) GetReplyTos() []string {
	if x != nil {
		return x.ReplyTos
	}
	return nil
}

type GenerateTextRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AppID     string               `protobuf:"bytes,10,opt,name=AppID,proto3" json:"AppID,omitempty"`
	LangID    string               `protobuf:"bytes,20,opt,name=LangID,proto3" json:"LangID,omitempty"`
	Channel   channel.NotifChannel `protobuf:"varint,30,opt,name=Channel,proto3,enum=notif.manager.channel.v1.NotifChannel" json:"Channel,omitempty"`
	EventType v1.UsedFor           `protobuf:"varint,40,opt,name=EventType,proto3,enum=basetypes.v1.UsedFor" json:"EventType,omitempty"`
	Vars      *TemplateVars        `protobuf:"bytes,50,opt,name=Vars,proto3,oneof" json:"Vars,omitempty"`
}

func (x *GenerateTextRequest) Reset() {
	*x = GenerateTextRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_notif_mw_v1_template_template_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenerateTextRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenerateTextRequest) ProtoMessage() {}

func (x *GenerateTextRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_notif_mw_v1_template_template_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenerateTextRequest.ProtoReflect.Descriptor instead.
func (*GenerateTextRequest) Descriptor() ([]byte, []int) {
	return file_npool_notif_mw_v1_template_template_proto_rawDescGZIP(), []int{2}
}

func (x *GenerateTextRequest) GetAppID() string {
	if x != nil {
		return x.AppID
	}
	return ""
}

func (x *GenerateTextRequest) GetLangID() string {
	if x != nil {
		return x.LangID
	}
	return ""
}

func (x *GenerateTextRequest) GetChannel() channel.NotifChannel {
	if x != nil {
		return x.Channel
	}
	return channel.NotifChannel(0)
}

func (x *GenerateTextRequest) GetEventType() v1.UsedFor {
	if x != nil {
		return x.EventType
	}
	return v1.UsedFor(0)
}

func (x *GenerateTextRequest) GetVars() *TemplateVars {
	if x != nil {
		return x.Vars
	}
	return nil
}

type GenerateTextResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *TextInfo `protobuf:"bytes,10,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (x *GenerateTextResponse) Reset() {
	*x = GenerateTextResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_notif_mw_v1_template_template_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenerateTextResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenerateTextResponse) ProtoMessage() {}

func (x *GenerateTextResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_notif_mw_v1_template_template_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenerateTextResponse.ProtoReflect.Descriptor instead.
func (*GenerateTextResponse) Descriptor() ([]byte, []int) {
	return file_npool_notif_mw_v1_template_template_proto_rawDescGZIP(), []int{3}
}

func (x *GenerateTextResponse) GetInfo() *TextInfo {
	if x != nil {
		return x.Info
	}
	return nil
}

var File_npool_notif_mw_v1_template_template_proto protoreflect.FileDescriptor

var file_npool_notif_mw_v1_template_template_proto_rawDesc = []byte{
	0x0a, 0x29, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x2f, 0x6d, 0x77,
	0x2f, 0x76, 0x31, 0x2f, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x2f, 0x74, 0x65, 0x6d,
	0x70, 0x6c, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1c, 0x6e, 0x6f, 0x74,
	0x69, 0x66, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x74, 0x65,
	0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x28, 0x6e, 0x70, 0x6f, 0x6f, 0x6c,
	0x2f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x2f, 0x6d, 0x67, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x68,
	0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x2f, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x64, 0x66, 0x6f, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbb, 0x02, 0x0a, 0x0c, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61,
	0x74, 0x65, 0x56, 0x61, 0x72, 0x73, 0x12, 0x1f, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x08, 0x55, 0x73, 0x65, 0x72,
	0x6e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x1d, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x07, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x88, 0x01, 0x01, 0x12, 0x1b, 0x0a, 0x06, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x1e, 0x20, 0x01, 0x28, 0x09, 0x48, 0x02, 0x52, 0x06, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x88, 0x01, 0x01, 0x12, 0x1f, 0x0a, 0x08, 0x43, 0x6f, 0x69, 0x6e, 0x55, 0x6e, 0x69, 0x74, 0x18,
	0x28, 0x20, 0x01, 0x28, 0x09, 0x48, 0x03, 0x52, 0x08, 0x43, 0x6f, 0x69, 0x6e, 0x55, 0x6e, 0x69,
	0x74, 0x88, 0x01, 0x01, 0x12, 0x21, 0x0a, 0x09, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x18, 0x32, 0x20, 0x01, 0x28, 0x0d, 0x48, 0x04, 0x52, 0x09, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x88, 0x01, 0x01, 0x12, 0x1d, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x18, 0x3c, 0x20, 0x01, 0x28, 0x09, 0x48, 0x05, 0x52, 0x07, 0x41, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x88, 0x01, 0x01, 0x12, 0x17, 0x0a, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x46,
	0x20, 0x01, 0x28, 0x09, 0x48, 0x06, 0x52, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x88, 0x01, 0x01, 0x42,
	0x0b, 0x0a, 0x09, 0x5f, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x0a, 0x0a, 0x08,
	0x5f, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x41, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x43, 0x6f, 0x69, 0x6e, 0x55, 0x6e, 0x69, 0x74,
	0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x0a,
	0x0a, 0x08, 0x5f, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x43,
	0x6f, 0x64, 0x65, 0x22, 0x84, 0x01, 0x0a, 0x08, 0x54, 0x65, 0x78, 0x74, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x18, 0x0a, 0x07, 0x53, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x53, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x43, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x43, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x46, 0x72, 0x6f, 0x6d, 0x18, 0x1e, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x46, 0x72, 0x6f, 0x6d, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x6f, 0x43, 0x43,
	0x73, 0x18, 0x28, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x54, 0x6f, 0x43, 0x43, 0x73, 0x12, 0x1a,
	0x0a, 0x08, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x54, 0x6f, 0x73, 0x18, 0x32, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x08, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x54, 0x6f, 0x73, 0x22, 0x88, 0x02, 0x0a, 0x13, 0x47,
	0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x54, 0x65, 0x78, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x41, 0x70, 0x70, 0x49, 0x44, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x41, 0x70, 0x70, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x4c, 0x61, 0x6e, 0x67,
	0x49, 0x44, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x4c, 0x61, 0x6e, 0x67, 0x49, 0x44,
	0x12, 0x40, 0x0a, 0x07, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x18, 0x1e, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x26, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65,
	0x72, 0x2e, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x6f, 0x74,
	0x69, 0x66, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x52, 0x07, 0x43, 0x68, 0x61, 0x6e, 0x6e,
	0x65, 0x6c, 0x12, 0x33, 0x0a, 0x09, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x18,
	0x28, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x15, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x64, 0x46, 0x6f, 0x72, 0x52, 0x09, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x43, 0x0a, 0x04, 0x56, 0x61, 0x72, 0x73, 0x18,
	0x32, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x2e, 0x6d, 0x69,
	0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x56, 0x61, 0x72,
	0x73, 0x48, 0x00, 0x52, 0x04, 0x56, 0x61, 0x72, 0x73, 0x88, 0x01, 0x01, 0x42, 0x07, 0x0a, 0x05,
	0x5f, 0x56, 0x61, 0x72, 0x73, 0x22, 0x52, 0x0a, 0x14, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74,
	0x65, 0x54, 0x65, 0x78, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3a, 0x0a,
	0x04, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x6e, 0x6f,
	0x74, 0x69, 0x66, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x74,
	0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x65, 0x78, 0x74, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x32, 0x85, 0x01, 0x0a, 0x0a, 0x4d, 0x69,
	0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x12, 0x77, 0x0a, 0x0c, 0x47, 0x65, 0x6e, 0x65,
	0x72, 0x61, 0x74, 0x65, 0x54, 0x65, 0x78, 0x74, 0x12, 0x31, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66,
	0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x74, 0x65, 0x6d, 0x70,
	0x6c, 0x61, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65,
	0x54, 0x65, 0x78, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x32, 0x2e, 0x6e, 0x6f,
	0x74, 0x69, 0x66, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x74,
	0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72,
	0x61, 0x74, 0x65, 0x54, 0x65, 0x78, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x42, 0x3d, 0x5a, 0x3b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x4e, 0x70, 0x6f, 0x6f, 0x6c, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x6e, 0x6f, 0x74, 0x69,
	0x66, 0x2f, 0x6d, 0x77, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_npool_notif_mw_v1_template_template_proto_rawDescOnce sync.Once
	file_npool_notif_mw_v1_template_template_proto_rawDescData = file_npool_notif_mw_v1_template_template_proto_rawDesc
)

func file_npool_notif_mw_v1_template_template_proto_rawDescGZIP() []byte {
	file_npool_notif_mw_v1_template_template_proto_rawDescOnce.Do(func() {
		file_npool_notif_mw_v1_template_template_proto_rawDescData = protoimpl.X.CompressGZIP(file_npool_notif_mw_v1_template_template_proto_rawDescData)
	})
	return file_npool_notif_mw_v1_template_template_proto_rawDescData
}

var file_npool_notif_mw_v1_template_template_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_npool_notif_mw_v1_template_template_proto_goTypes = []interface{}{
	(*TemplateVars)(nil),         // 0: notif.middleware.template.v1.TemplateVars
	(*TextInfo)(nil),             // 1: notif.middleware.template.v1.TextInfo
	(*GenerateTextRequest)(nil),  // 2: notif.middleware.template.v1.GenerateTextRequest
	(*GenerateTextResponse)(nil), // 3: notif.middleware.template.v1.GenerateTextResponse
	(channel.NotifChannel)(0),    // 4: notif.manager.channel.v1.NotifChannel
	(v1.UsedFor)(0),              // 5: basetypes.v1.UsedFor
}
var file_npool_notif_mw_v1_template_template_proto_depIdxs = []int32{
	4, // 0: notif.middleware.template.v1.GenerateTextRequest.Channel:type_name -> notif.manager.channel.v1.NotifChannel
	5, // 1: notif.middleware.template.v1.GenerateTextRequest.EventType:type_name -> basetypes.v1.UsedFor
	0, // 2: notif.middleware.template.v1.GenerateTextRequest.Vars:type_name -> notif.middleware.template.v1.TemplateVars
	1, // 3: notif.middleware.template.v1.GenerateTextResponse.Info:type_name -> notif.middleware.template.v1.TextInfo
	2, // 4: notif.middleware.template.v1.Middleware.GenerateText:input_type -> notif.middleware.template.v1.GenerateTextRequest
	3, // 5: notif.middleware.template.v1.Middleware.GenerateText:output_type -> notif.middleware.template.v1.GenerateTextResponse
	5, // [5:6] is the sub-list for method output_type
	4, // [4:5] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_npool_notif_mw_v1_template_template_proto_init() }
func file_npool_notif_mw_v1_template_template_proto_init() {
	if File_npool_notif_mw_v1_template_template_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_npool_notif_mw_v1_template_template_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TemplateVars); i {
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
		file_npool_notif_mw_v1_template_template_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TextInfo); i {
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
		file_npool_notif_mw_v1_template_template_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenerateTextRequest); i {
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
		file_npool_notif_mw_v1_template_template_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenerateTextResponse); i {
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
	file_npool_notif_mw_v1_template_template_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_npool_notif_mw_v1_template_template_proto_msgTypes[2].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_npool_notif_mw_v1_template_template_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_npool_notif_mw_v1_template_template_proto_goTypes,
		DependencyIndexes: file_npool_notif_mw_v1_template_template_proto_depIdxs,
		MessageInfos:      file_npool_notif_mw_v1_template_template_proto_msgTypes,
	}.Build()
	File_npool_notif_mw_v1_template_template_proto = out.File
	file_npool_notif_mw_v1_template_template_proto_rawDesc = nil
	file_npool_notif_mw_v1_template_template_proto_goTypes = nil
	file_npool_notif_mw_v1_template_template_proto_depIdxs = nil
}
