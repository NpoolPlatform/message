// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.18.1
// source: npool/subscribemgr/subscribemgr.proto

package subscribemgr

import (
	npool "github.com/NpoolPlatform/message/npool"
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

type EmailSubscriber struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID           string `protobuf:"bytes,10,opt,name=ID,proto3" json:"ID,omitempty"`
	AppID        string `protobuf:"bytes,20,opt,name=AppID,proto3" json:"AppID,omitempty"`
	EmailAddress string `protobuf:"bytes,30,opt,name=EmailAddress,proto3" json:"EmailAddress,omitempty"`
}

func (x *EmailSubscriber) Reset() {
	*x = EmailSubscriber{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_subscribemgr_subscribemgr_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmailSubscriber) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmailSubscriber) ProtoMessage() {}

func (x *EmailSubscriber) ProtoReflect() protoreflect.Message {
	mi := &file_npool_subscribemgr_subscribemgr_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmailSubscriber.ProtoReflect.Descriptor instead.
func (*EmailSubscriber) Descriptor() ([]byte, []int) {
	return file_npool_subscribemgr_subscribemgr_proto_rawDescGZIP(), []int{0}
}

func (x *EmailSubscriber) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *EmailSubscriber) GetAppID() string {
	if x != nil {
		return x.AppID
	}
	return ""
}

func (x *EmailSubscriber) GetEmailAddress() string {
	if x != nil {
		return x.EmailAddress
	}
	return ""
}

type CreateEmailSubscriberRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *EmailSubscriber `protobuf:"bytes,10,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (x *CreateEmailSubscriberRequest) Reset() {
	*x = CreateEmailSubscriberRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_subscribemgr_subscribemgr_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateEmailSubscriberRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateEmailSubscriberRequest) ProtoMessage() {}

func (x *CreateEmailSubscriberRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_subscribemgr_subscribemgr_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateEmailSubscriberRequest.ProtoReflect.Descriptor instead.
func (*CreateEmailSubscriberRequest) Descriptor() ([]byte, []int) {
	return file_npool_subscribemgr_subscribemgr_proto_rawDescGZIP(), []int{1}
}

func (x *CreateEmailSubscriberRequest) GetInfo() *EmailSubscriber {
	if x != nil {
		return x.Info
	}
	return nil
}

type CreateEmailSubscriberResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *EmailSubscriber `protobuf:"bytes,10,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (x *CreateEmailSubscriberResponse) Reset() {
	*x = CreateEmailSubscriberResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_subscribemgr_subscribemgr_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateEmailSubscriberResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateEmailSubscriberResponse) ProtoMessage() {}

func (x *CreateEmailSubscriberResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_subscribemgr_subscribemgr_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateEmailSubscriberResponse.ProtoReflect.Descriptor instead.
func (*CreateEmailSubscriberResponse) Descriptor() ([]byte, []int) {
	return file_npool_subscribemgr_subscribemgr_proto_rawDescGZIP(), []int{2}
}

func (x *CreateEmailSubscriberResponse) GetInfo() *EmailSubscriber {
	if x != nil {
		return x.Info
	}
	return nil
}

type GetEmailSubscribersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AppID string                       `protobuf:"bytes,10,opt,name=AppID,proto3" json:"AppID,omitempty"`
	Conds map[string]*npool.FilterCond `protobuf:"bytes,20,rep,name=Conds,proto3" json:"Conds,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *GetEmailSubscribersRequest) Reset() {
	*x = GetEmailSubscribersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_subscribemgr_subscribemgr_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetEmailSubscribersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetEmailSubscribersRequest) ProtoMessage() {}

func (x *GetEmailSubscribersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_subscribemgr_subscribemgr_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetEmailSubscribersRequest.ProtoReflect.Descriptor instead.
func (*GetEmailSubscribersRequest) Descriptor() ([]byte, []int) {
	return file_npool_subscribemgr_subscribemgr_proto_rawDescGZIP(), []int{3}
}

func (x *GetEmailSubscribersRequest) GetAppID() string {
	if x != nil {
		return x.AppID
	}
	return ""
}

func (x *GetEmailSubscribersRequest) GetConds() map[string]*npool.FilterCond {
	if x != nil {
		return x.Conds
	}
	return nil
}

type GetEmailSubscribersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Infos []*EmailSubscriber `protobuf:"bytes,10,rep,name=Infos,proto3" json:"Infos,omitempty"`
}

func (x *GetEmailSubscribersResponse) Reset() {
	*x = GetEmailSubscribersResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_subscribemgr_subscribemgr_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetEmailSubscribersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetEmailSubscribersResponse) ProtoMessage() {}

func (x *GetEmailSubscribersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_subscribemgr_subscribemgr_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetEmailSubscribersResponse.ProtoReflect.Descriptor instead.
func (*GetEmailSubscribersResponse) Descriptor() ([]byte, []int) {
	return file_npool_subscribemgr_subscribemgr_proto_rawDescGZIP(), []int{4}
}

func (x *GetEmailSubscribersResponse) GetInfos() []*EmailSubscriber {
	if x != nil {
		return x.Infos
	}
	return nil
}

type GetAppEmailSubscribersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TargetAppID string                       `protobuf:"bytes,10,opt,name=TargetAppID,proto3" json:"TargetAppID,omitempty"`
	Conds       map[string]*npool.FilterCond `protobuf:"bytes,20,rep,name=Conds,proto3" json:"Conds,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *GetAppEmailSubscribersRequest) Reset() {
	*x = GetAppEmailSubscribersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_subscribemgr_subscribemgr_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAppEmailSubscribersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAppEmailSubscribersRequest) ProtoMessage() {}

func (x *GetAppEmailSubscribersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_subscribemgr_subscribemgr_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAppEmailSubscribersRequest.ProtoReflect.Descriptor instead.
func (*GetAppEmailSubscribersRequest) Descriptor() ([]byte, []int) {
	return file_npool_subscribemgr_subscribemgr_proto_rawDescGZIP(), []int{5}
}

func (x *GetAppEmailSubscribersRequest) GetTargetAppID() string {
	if x != nil {
		return x.TargetAppID
	}
	return ""
}

func (x *GetAppEmailSubscribersRequest) GetConds() map[string]*npool.FilterCond {
	if x != nil {
		return x.Conds
	}
	return nil
}

type GetAppEmailSubscribersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Infos []*EmailSubscriber `protobuf:"bytes,10,rep,name=Infos,proto3" json:"Infos,omitempty"`
}

func (x *GetAppEmailSubscribersResponse) Reset() {
	*x = GetAppEmailSubscribersResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_subscribemgr_subscribemgr_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAppEmailSubscribersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAppEmailSubscribersResponse) ProtoMessage() {}

func (x *GetAppEmailSubscribersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_subscribemgr_subscribemgr_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAppEmailSubscribersResponse.ProtoReflect.Descriptor instead.
func (*GetAppEmailSubscribersResponse) Descriptor() ([]byte, []int) {
	return file_npool_subscribemgr_subscribemgr_proto_rawDescGZIP(), []int{6}
}

func (x *GetAppEmailSubscribersResponse) GetInfos() []*EmailSubscriber {
	if x != nil {
		return x.Infos
	}
	return nil
}

var File_npool_subscribemgr_subscribemgr_proto protoreflect.FileDescriptor

var file_npool_subscribemgr_subscribemgr_proto_rawDesc = []byte{
	0x0a, 0x25, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62,
	0x65, 0x6d, 0x67, 0x72, 0x2f, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x6d, 0x67,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x14, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69,
	0x62, 0x65, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70,
	0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f,
	0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5b, 0x0a, 0x0f, 0x45,
	0x6d, 0x61, 0x69, 0x6c, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x72, 0x12, 0x0e,
	0x0a, 0x02, 0x49, 0x44, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x12, 0x14,
	0x0a, 0x05, 0x41, 0x70, 0x70, 0x49, 0x44, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x41,
	0x70, 0x70, 0x49, 0x44, 0x12, 0x22, 0x0a, 0x0c, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x41, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x45, 0x6d, 0x61, 0x69,
	0x6c, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0x59, 0x0a, 0x1c, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x39, 0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69,
	0x62, 0x65, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6d,
	0x61, 0x69, 0x6c, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x72, 0x52, 0x04, 0x49,
	0x6e, 0x66, 0x6f, 0x22, 0x5a, 0x0a, 0x1d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x45, 0x6d, 0x61,
	0x69, 0x6c, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x39, 0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x25, 0x2e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x2e, 0x6d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x53,
	0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x72, 0x52, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x22,
	0xd5, 0x01, 0x0a, 0x1a, 0x47, 0x65, 0x74, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x53, 0x75, 0x62, 0x73,
	0x63, 0x72, 0x69, 0x62, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14,
	0x0a, 0x05, 0x41, 0x70, 0x70, 0x49, 0x44, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x41,
	0x70, 0x70, 0x49, 0x44, 0x12, 0x51, 0x0a, 0x05, 0x43, 0x6f, 0x6e, 0x64, 0x73, 0x18, 0x14, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x3b, 0x2e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x2e,
	0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x45, 0x6d,
	0x61, 0x69, 0x6c, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x72, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x43, 0x6f, 0x6e, 0x64, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x52, 0x05, 0x43, 0x6f, 0x6e, 0x64, 0x73, 0x1a, 0x4e, 0x0a, 0x0a, 0x43, 0x6f, 0x6e, 0x64, 0x73,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2a, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2e, 0x76,
	0x31, 0x2e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x64, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x5a, 0x0a, 0x1b, 0x47, 0x65, 0x74, 0x45, 0x6d,
	0x61, 0x69, 0x6c, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x72, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3b, 0x0a, 0x05, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x18,
	0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62,
	0x65, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6d, 0x61,
	0x69, 0x6c, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x72, 0x52, 0x05, 0x49, 0x6e,
	0x66, 0x6f, 0x73, 0x22, 0xe7, 0x01, 0x0a, 0x1d, 0x47, 0x65, 0x74, 0x41, 0x70, 0x70, 0x45, 0x6d,
	0x61, 0x69, 0x6c, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x72, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x41,
	0x70, 0x70, 0x49, 0x44, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x54, 0x61, 0x72, 0x67,
	0x65, 0x74, 0x41, 0x70, 0x70, 0x49, 0x44, 0x12, 0x54, 0x0a, 0x05, 0x43, 0x6f, 0x6e, 0x64, 0x73,
	0x18, 0x14, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3e, 0x2e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69,
	0x62, 0x65, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65,
	0x74, 0x41, 0x70, 0x70, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69,
	0x62, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x43, 0x6f, 0x6e, 0x64,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x05, 0x43, 0x6f, 0x6e, 0x64, 0x73, 0x1a, 0x4e, 0x0a,
	0x0a, 0x43, 0x6f, 0x6e, 0x64, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2a, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x6e,
	0x70, 0x6f, 0x6f, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x43, 0x6f,
	0x6e, 0x64, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x5d, 0x0a,
	0x1e, 0x47, 0x65, 0x74, 0x41, 0x70, 0x70, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x53, 0x75, 0x62, 0x73,
	0x63, 0x72, 0x69, 0x62, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x3b, 0x0a, 0x05, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x25,
	0x2e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67,
	0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x53, 0x75, 0x62, 0x73, 0x63,
	0x72, 0x69, 0x62, 0x65, 0x72, 0x52, 0x05, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x32, 0xe3, 0x04, 0x0a,
	0x10, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65,
	0x72, 0x12, 0x51, 0x0a, 0x07, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x1a, 0x19, 0x2e, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2e, 0x76, 0x31, 0x2e,
	0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x13, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0d, 0x22, 0x08, 0x2f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x3a, 0x01, 0x2a, 0x12, 0xa8, 0x01, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x45,
	0x6d, 0x61, 0x69, 0x6c, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x72, 0x12, 0x32,
	0x2e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67,
	0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x45, 0x6d, 0x61, 0x69,
	0x6c, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x33, 0x2e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x2e, 0x6d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x45, 0x6d, 0x61, 0x69, 0x6c, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x72, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x26, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x20, 0x22,
	0x1b, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x2f, 0x65, 0x6d, 0x61, 0x69,
	0x6c, 0x2f, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x72, 0x3a, 0x01, 0x2a, 0x12,
	0xa0, 0x01, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x53, 0x75, 0x62, 0x73,
	0x63, 0x72, 0x69, 0x62, 0x65, 0x72, 0x73, 0x12, 0x30, 0x2e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72,
	0x69, 0x62, 0x65, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47,
	0x65, 0x74, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65,
	0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x31, 0x2e, 0x73, 0x75, 0x62, 0x73,
	0x63, 0x72, 0x69, 0x62, 0x65, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x31,
	0x2e, 0x47, 0x65, 0x74, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69,
	0x62, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x24, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x1e, 0x22, 0x19, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x65, 0x74, 0x2f, 0x65, 0x6d,
	0x61, 0x69, 0x6c, 0x2f, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x72, 0x73, 0x3a,
	0x01, 0x2a, 0x12, 0xad, 0x01, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x41, 0x70, 0x70, 0x45, 0x6d, 0x61,
	0x69, 0x6c, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x72, 0x73, 0x12, 0x33, 0x2e,
	0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65,
	0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x70, 0x70, 0x45, 0x6d, 0x61, 0x69, 0x6c,
	0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x34, 0x2e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x2e, 0x6d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x70, 0x70,
	0x45, 0x6d, 0x61, 0x69, 0x6c, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x72, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x28, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x22,
	0x22, 0x1d, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x65, 0x74, 0x2f, 0x61, 0x70, 0x70, 0x2f, 0x65, 0x6d,
	0x61, 0x69, 0x6c, 0x2f, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x72, 0x73, 0x3a,
	0x01, 0x2a, 0x42, 0x35, 0x5a, 0x33, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x4e, 0x70, 0x6f, 0x6f, 0x6c, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x73, 0x75, 0x62,
	0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x6d, 0x67, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_npool_subscribemgr_subscribemgr_proto_rawDescOnce sync.Once
	file_npool_subscribemgr_subscribemgr_proto_rawDescData = file_npool_subscribemgr_subscribemgr_proto_rawDesc
)

func file_npool_subscribemgr_subscribemgr_proto_rawDescGZIP() []byte {
	file_npool_subscribemgr_subscribemgr_proto_rawDescOnce.Do(func() {
		file_npool_subscribemgr_subscribemgr_proto_rawDescData = protoimpl.X.CompressGZIP(file_npool_subscribemgr_subscribemgr_proto_rawDescData)
	})
	return file_npool_subscribemgr_subscribemgr_proto_rawDescData
}

var file_npool_subscribemgr_subscribemgr_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_npool_subscribemgr_subscribemgr_proto_goTypes = []interface{}{
	(*EmailSubscriber)(nil),                // 0: subscribe.manager.v1.EmailSubscriber
	(*CreateEmailSubscriberRequest)(nil),   // 1: subscribe.manager.v1.CreateEmailSubscriberRequest
	(*CreateEmailSubscriberResponse)(nil),  // 2: subscribe.manager.v1.CreateEmailSubscriberResponse
	(*GetEmailSubscribersRequest)(nil),     // 3: subscribe.manager.v1.GetEmailSubscribersRequest
	(*GetEmailSubscribersResponse)(nil),    // 4: subscribe.manager.v1.GetEmailSubscribersResponse
	(*GetAppEmailSubscribersRequest)(nil),  // 5: subscribe.manager.v1.GetAppEmailSubscribersRequest
	(*GetAppEmailSubscribersResponse)(nil), // 6: subscribe.manager.v1.GetAppEmailSubscribersResponse
	nil,                                    // 7: subscribe.manager.v1.GetEmailSubscribersRequest.CondsEntry
	nil,                                    // 8: subscribe.manager.v1.GetAppEmailSubscribersRequest.CondsEntry
	(*npool.FilterCond)(nil),               // 9: npool.v1.FilterCond
	(*emptypb.Empty)(nil),                  // 10: google.protobuf.Empty
	(*npool.VersionResponse)(nil),          // 11: npool.v1.VersionResponse
}
var file_npool_subscribemgr_subscribemgr_proto_depIdxs = []int32{
	0,  // 0: subscribe.manager.v1.CreateEmailSubscriberRequest.Info:type_name -> subscribe.manager.v1.EmailSubscriber
	0,  // 1: subscribe.manager.v1.CreateEmailSubscriberResponse.Info:type_name -> subscribe.manager.v1.EmailSubscriber
	7,  // 2: subscribe.manager.v1.GetEmailSubscribersRequest.Conds:type_name -> subscribe.manager.v1.GetEmailSubscribersRequest.CondsEntry
	0,  // 3: subscribe.manager.v1.GetEmailSubscribersResponse.Infos:type_name -> subscribe.manager.v1.EmailSubscriber
	8,  // 4: subscribe.manager.v1.GetAppEmailSubscribersRequest.Conds:type_name -> subscribe.manager.v1.GetAppEmailSubscribersRequest.CondsEntry
	0,  // 5: subscribe.manager.v1.GetAppEmailSubscribersResponse.Infos:type_name -> subscribe.manager.v1.EmailSubscriber
	9,  // 6: subscribe.manager.v1.GetEmailSubscribersRequest.CondsEntry.value:type_name -> npool.v1.FilterCond
	9,  // 7: subscribe.manager.v1.GetAppEmailSubscribersRequest.CondsEntry.value:type_name -> npool.v1.FilterCond
	10, // 8: subscribe.manager.v1.SubscribeManager.Version:input_type -> google.protobuf.Empty
	1,  // 9: subscribe.manager.v1.SubscribeManager.CreateEmailSubscriber:input_type -> subscribe.manager.v1.CreateEmailSubscriberRequest
	3,  // 10: subscribe.manager.v1.SubscribeManager.GetEmailSubscribers:input_type -> subscribe.manager.v1.GetEmailSubscribersRequest
	5,  // 11: subscribe.manager.v1.SubscribeManager.GetAppEmailSubscribers:input_type -> subscribe.manager.v1.GetAppEmailSubscribersRequest
	11, // 12: subscribe.manager.v1.SubscribeManager.Version:output_type -> npool.v1.VersionResponse
	2,  // 13: subscribe.manager.v1.SubscribeManager.CreateEmailSubscriber:output_type -> subscribe.manager.v1.CreateEmailSubscriberResponse
	4,  // 14: subscribe.manager.v1.SubscribeManager.GetEmailSubscribers:output_type -> subscribe.manager.v1.GetEmailSubscribersResponse
	6,  // 15: subscribe.manager.v1.SubscribeManager.GetAppEmailSubscribers:output_type -> subscribe.manager.v1.GetAppEmailSubscribersResponse
	12, // [12:16] is the sub-list for method output_type
	8,  // [8:12] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_npool_subscribemgr_subscribemgr_proto_init() }
func file_npool_subscribemgr_subscribemgr_proto_init() {
	if File_npool_subscribemgr_subscribemgr_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_npool_subscribemgr_subscribemgr_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmailSubscriber); i {
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
		file_npool_subscribemgr_subscribemgr_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateEmailSubscriberRequest); i {
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
		file_npool_subscribemgr_subscribemgr_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateEmailSubscriberResponse); i {
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
		file_npool_subscribemgr_subscribemgr_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetEmailSubscribersRequest); i {
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
		file_npool_subscribemgr_subscribemgr_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetEmailSubscribersResponse); i {
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
		file_npool_subscribemgr_subscribemgr_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAppEmailSubscribersRequest); i {
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
		file_npool_subscribemgr_subscribemgr_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAppEmailSubscribersResponse); i {
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
			RawDescriptor: file_npool_subscribemgr_subscribemgr_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_npool_subscribemgr_subscribemgr_proto_goTypes,
		DependencyIndexes: file_npool_subscribemgr_subscribemgr_proto_depIdxs,
		MessageInfos:      file_npool_subscribemgr_subscribemgr_proto_msgTypes,
	}.Build()
	File_npool_subscribemgr_subscribemgr_proto = out.File
	file_npool_subscribemgr_subscribemgr_proto_rawDesc = nil
	file_npool_subscribemgr_subscribemgr_proto_goTypes = nil
	file_npool_subscribemgr_subscribemgr_proto_depIdxs = nil
}
