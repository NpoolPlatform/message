// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.18.1
// source: npool/third/mw/v1/notif/notify.proto

package notif

import (
	notif "github.com/NpoolPlatform/message/npool/notif/mgr/v1/notif"
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

// if UseTemplate = true Title and Content Valid
type SendNotifEmailRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AppID           string          `protobuf:"bytes,10,opt,name=AppID,proto3" json:"AppID,omitempty"`
	FromAccount     *string         `protobuf:"bytes,20,opt,name=FromAccount,proto3,oneof" json:"FromAccount,omitempty"`
	UsedFor         notif.EventType `protobuf:"varint,30,opt,name=UsedFor,proto3,enum=notif.manager.notif1.v1.EventType" json:"UsedFor,omitempty"`
	ReceiverAccount string          `protobuf:"bytes,40,opt,name=ReceiverAccount,proto3" json:"ReceiverAccount,omitempty"`
	LangID          string          `protobuf:"bytes,50,opt,name=LangID,proto3" json:"LangID,omitempty"`
	SenderName      *string         `protobuf:"bytes,60,opt,name=SenderName,proto3,oneof" json:"SenderName,omitempty"`
	ReceiverName    *string         `protobuf:"bytes,70,opt,name=ReceiverName,proto3,oneof" json:"ReceiverName,omitempty"`
	Message         *string         `protobuf:"bytes,80,opt,name=Message,proto3,oneof" json:"Message,omitempty"`
	UseTemplate     bool            `protobuf:"varint,90,opt,name=UseTemplate,proto3" json:"UseTemplate,omitempty"`
	Title           string          `protobuf:"bytes,100,opt,name=Title,proto3" json:"Title,omitempty"`
	Content         string          `protobuf:"bytes,110,opt,name=Content,proto3" json:"Content,omitempty"`
}

func (x *SendNotifEmailRequest) Reset() {
	*x = SendNotifEmailRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_third_mw_v1_notif_notify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendNotifEmailRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendNotifEmailRequest) ProtoMessage() {}

func (x *SendNotifEmailRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_third_mw_v1_notif_notify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendNotifEmailRequest.ProtoReflect.Descriptor instead.
func (*SendNotifEmailRequest) Descriptor() ([]byte, []int) {
	return file_npool_third_mw_v1_notif_notify_proto_rawDescGZIP(), []int{0}
}

func (x *SendNotifEmailRequest) GetAppID() string {
	if x != nil {
		return x.AppID
	}
	return ""
}

func (x *SendNotifEmailRequest) GetFromAccount() string {
	if x != nil && x.FromAccount != nil {
		return *x.FromAccount
	}
	return ""
}

func (x *SendNotifEmailRequest) GetUsedFor() notif.EventType {
	if x != nil {
		return x.UsedFor
	}
	return notif.EventType(0)
}

func (x *SendNotifEmailRequest) GetReceiverAccount() string {
	if x != nil {
		return x.ReceiverAccount
	}
	return ""
}

func (x *SendNotifEmailRequest) GetLangID() string {
	if x != nil {
		return x.LangID
	}
	return ""
}

func (x *SendNotifEmailRequest) GetSenderName() string {
	if x != nil && x.SenderName != nil {
		return *x.SenderName
	}
	return ""
}

func (x *SendNotifEmailRequest) GetReceiverName() string {
	if x != nil && x.ReceiverName != nil {
		return *x.ReceiverName
	}
	return ""
}

func (x *SendNotifEmailRequest) GetMessage() string {
	if x != nil && x.Message != nil {
		return *x.Message
	}
	return ""
}

func (x *SendNotifEmailRequest) GetUseTemplate() bool {
	if x != nil {
		return x.UseTemplate
	}
	return false
}

func (x *SendNotifEmailRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *SendNotifEmailRequest) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

type SendNotifEmailResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SendNotifEmailResponse) Reset() {
	*x = SendNotifEmailResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_third_mw_v1_notif_notify_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendNotifEmailResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendNotifEmailResponse) ProtoMessage() {}

func (x *SendNotifEmailResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_third_mw_v1_notif_notify_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendNotifEmailResponse.ProtoReflect.Descriptor instead.
func (*SendNotifEmailResponse) Descriptor() ([]byte, []int) {
	return file_npool_third_mw_v1_notif_notify_proto_rawDescGZIP(), []int{1}
}

type SendAnnouncementRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ReceiverAccount []string `protobuf:"bytes,10,rep,name=ReceiverAccount,proto3" json:"ReceiverAccount,omitempty"`
	Title           string   `protobuf:"bytes,20,opt,name=title,proto3" json:"title,omitempty"`
	Content         string   `protobuf:"bytes,30,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *SendAnnouncementRequest) Reset() {
	*x = SendAnnouncementRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_third_mw_v1_notif_notify_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendAnnouncementRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendAnnouncementRequest) ProtoMessage() {}

func (x *SendAnnouncementRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_third_mw_v1_notif_notify_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendAnnouncementRequest.ProtoReflect.Descriptor instead.
func (*SendAnnouncementRequest) Descriptor() ([]byte, []int) {
	return file_npool_third_mw_v1_notif_notify_proto_rawDescGZIP(), []int{2}
}

func (x *SendAnnouncementRequest) GetReceiverAccount() []string {
	if x != nil {
		return x.ReceiverAccount
	}
	return nil
}

func (x *SendAnnouncementRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *SendAnnouncementRequest) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

type SendAnnouncementResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SendAnnouncementResponse) Reset() {
	*x = SendAnnouncementResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_third_mw_v1_notif_notify_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendAnnouncementResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendAnnouncementResponse) ProtoMessage() {}

func (x *SendAnnouncementResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_third_mw_v1_notif_notify_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendAnnouncementResponse.ProtoReflect.Descriptor instead.
func (*SendAnnouncementResponse) Descriptor() ([]byte, []int) {
	return file_npool_third_mw_v1_notif_notify_proto_rawDescGZIP(), []int{3}
}

var File_npool_third_mw_v1_notif_notify_proto protoreflect.FileDescriptor

var file_npool_third_mw_v1_notif_notify_proto_rawDesc = []byte{
	0x0a, 0x24, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x74, 0x68, 0x69, 0x72, 0x64, 0x2f, 0x6d, 0x77,
	0x2f, 0x76, 0x31, 0x2f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x2f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1a, 0x74, 0x68, 0x69, 0x72, 0x64, 0x2e, 0x6d, 0x69,
	0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x31, 0x2e,
	0x76, 0x31, 0x1a, 0x24, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x2f,
	0x6d, 0x67, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x2f, 0x6e, 0x6f, 0x74,
	0x69, 0x66, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xcf, 0x03, 0x0a, 0x15, 0x53, 0x65, 0x6e,
	0x64, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x41, 0x70, 0x70, 0x49, 0x44, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x41, 0x70, 0x70, 0x49, 0x44, 0x12, 0x25, 0x0a, 0x0b, 0x46, 0x72, 0x6f, 0x6d,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52,
	0x0b, 0x46, 0x72, 0x6f, 0x6d, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x88, 0x01, 0x01, 0x12,
	0x3c, 0x0a, 0x07, 0x55, 0x73, 0x65, 0x64, 0x46, 0x6f, 0x72, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x22, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72,
	0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x31, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x54, 0x79, 0x70, 0x65, 0x52, 0x07, 0x55, 0x73, 0x65, 0x64, 0x46, 0x6f, 0x72, 0x12, 0x28, 0x0a,
	0x0f, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x28, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x4c, 0x61, 0x6e, 0x67, 0x49,
	0x44, 0x18, 0x32, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x4c, 0x61, 0x6e, 0x67, 0x49, 0x44, 0x12,
	0x23, 0x0a, 0x0a, 0x53, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x3c, 0x20,
	0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x0a, 0x53, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x4e, 0x61, 0x6d,
	0x65, 0x88, 0x01, 0x01, 0x12, 0x27, 0x0a, 0x0c, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72,
	0x4e, 0x61, 0x6d, 0x65, 0x18, 0x46, 0x20, 0x01, 0x28, 0x09, 0x48, 0x02, 0x52, 0x0c, 0x52, 0x65,
	0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x1d, 0x0a,
	0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x50, 0x20, 0x01, 0x28, 0x09, 0x48, 0x03,
	0x52, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x88, 0x01, 0x01, 0x12, 0x20, 0x0a, 0x0b,
	0x55, 0x73, 0x65, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x18, 0x5a, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x0b, 0x55, 0x73, 0x65, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x64, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x54,
	0x69, 0x74, 0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18,
	0x6e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x42, 0x0e,
	0x0a, 0x0c, 0x5f, 0x46, 0x72, 0x6f, 0x6d, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x0d,
	0x0a, 0x0b, 0x5f, 0x53, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x42, 0x0f, 0x0a,
	0x0d, 0x5f, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x42, 0x0a,
	0x0a, 0x08, 0x5f, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x18, 0x0a, 0x16, 0x53, 0x65,
	0x6e, 0x64, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x73, 0x0a, 0x17, 0x53, 0x65, 0x6e, 0x64, 0x41, 0x6e, 0x6e, 0x6f,
	0x75, 0x6e, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x28, 0x0a, 0x0f, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x41, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0f, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76,
	0x65, 0x72, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x1a, 0x0a, 0x18, 0x53, 0x65, 0x6e,
	0x64, 0x41, 0x6e, 0x6e, 0x6f, 0x75, 0x6e, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0x8a, 0x02, 0x0a, 0x0a, 0x4d, 0x69, 0x64, 0x64, 0x6c, 0x65,
	0x77, 0x61, 0x72, 0x65, 0x12, 0x79, 0x0a, 0x0e, 0x53, 0x65, 0x6e, 0x64, 0x4e, 0x6f, 0x74, 0x69,
	0x66, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x31, 0x2e, 0x74, 0x68, 0x69, 0x72, 0x64, 0x2e, 0x6d,
	0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x31,
	0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x45, 0x6d, 0x61,
	0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x32, 0x2e, 0x74, 0x68, 0x69, 0x72,
	0x64, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x6e, 0x6f, 0x74,
	0x69, 0x66, 0x31, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x4e, 0x6f, 0x74, 0x69, 0x66,
	0x45, 0x6d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x80, 0x01, 0x0a, 0x15, 0x53, 0x65, 0x6e, 0x64, 0x41, 0x6e, 0x6e, 0x6f, 0x75, 0x6e, 0x63, 0x65,
	0x6d, 0x65, 0x6e, 0x74, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x31, 0x2e, 0x74, 0x68, 0x69, 0x72,
	0x64, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x6e, 0x6f, 0x74,
	0x69, 0x66, 0x31, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x4e, 0x6f, 0x74, 0x69, 0x66,
	0x45, 0x6d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x32, 0x2e, 0x74,
	0x68, 0x69, 0x72, 0x64, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e,
	0x6e, 0x6f, 0x74, 0x69, 0x66, 0x31, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x4e, 0x6f,
	0x74, 0x69, 0x66, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x42, 0x3a, 0x5a, 0x38, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x4e, 0x70, 0x6f, 0x6f, 0x6c, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x74, 0x68, 0x69,
	0x72, 0x64, 0x2f, 0x6d, 0x77, 0x2f, 0x76, 0x31, 0x2f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_npool_third_mw_v1_notif_notify_proto_rawDescOnce sync.Once
	file_npool_third_mw_v1_notif_notify_proto_rawDescData = file_npool_third_mw_v1_notif_notify_proto_rawDesc
)

func file_npool_third_mw_v1_notif_notify_proto_rawDescGZIP() []byte {
	file_npool_third_mw_v1_notif_notify_proto_rawDescOnce.Do(func() {
		file_npool_third_mw_v1_notif_notify_proto_rawDescData = protoimpl.X.CompressGZIP(file_npool_third_mw_v1_notif_notify_proto_rawDescData)
	})
	return file_npool_third_mw_v1_notif_notify_proto_rawDescData
}

var file_npool_third_mw_v1_notif_notify_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_npool_third_mw_v1_notif_notify_proto_goTypes = []interface{}{
	(*SendNotifEmailRequest)(nil),    // 0: third.middleware.notif1.v1.SendNotifEmailRequest
	(*SendNotifEmailResponse)(nil),   // 1: third.middleware.notif1.v1.SendNotifEmailResponse
	(*SendAnnouncementRequest)(nil),  // 2: third.middleware.notif1.v1.SendAnnouncementRequest
	(*SendAnnouncementResponse)(nil), // 3: third.middleware.notif1.v1.SendAnnouncementResponse
	(notif.EventType)(0),             // 4: notif.manager.notif1.v1.EventType
}
var file_npool_third_mw_v1_notif_notify_proto_depIdxs = []int32{
	4, // 0: third.middleware.notif1.v1.SendNotifEmailRequest.UsedFor:type_name -> notif.manager.notif1.v1.EventType
	0, // 1: third.middleware.notif1.v1.Middleware.SendNotifEmail:input_type -> third.middleware.notif1.v1.SendNotifEmailRequest
	0, // 2: third.middleware.notif1.v1.Middleware.SendAnnouncementEmail:input_type -> third.middleware.notif1.v1.SendNotifEmailRequest
	1, // 3: third.middleware.notif1.v1.Middleware.SendNotifEmail:output_type -> third.middleware.notif1.v1.SendNotifEmailResponse
	1, // 4: third.middleware.notif1.v1.Middleware.SendAnnouncementEmail:output_type -> third.middleware.notif1.v1.SendNotifEmailResponse
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_npool_third_mw_v1_notif_notify_proto_init() }
func file_npool_third_mw_v1_notif_notify_proto_init() {
	if File_npool_third_mw_v1_notif_notify_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_npool_third_mw_v1_notif_notify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendNotifEmailRequest); i {
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
		file_npool_third_mw_v1_notif_notify_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendNotifEmailResponse); i {
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
		file_npool_third_mw_v1_notif_notify_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendAnnouncementRequest); i {
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
		file_npool_third_mw_v1_notif_notify_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendAnnouncementResponse); i {
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
	file_npool_third_mw_v1_notif_notify_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_npool_third_mw_v1_notif_notify_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_npool_third_mw_v1_notif_notify_proto_goTypes,
		DependencyIndexes: file_npool_third_mw_v1_notif_notify_proto_depIdxs,
		MessageInfos:      file_npool_third_mw_v1_notif_notify_proto_msgTypes,
	}.Build()
	File_npool_third_mw_v1_notif_notify_proto = out.File
	file_npool_third_mw_v1_notif_notify_proto_rawDesc = nil
	file_npool_third_mw_v1_notif_notify_proto_goTypes = nil
	file_npool_third_mw_v1_notif_notify_proto_depIdxs = nil
}
