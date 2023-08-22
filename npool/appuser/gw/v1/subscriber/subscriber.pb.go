// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.18.1
// source: npool/appuser/gw/v1/subscriber/subscriber.proto

package subscriber

import (
	subscriber "github.com/NpoolPlatform/message/npool/appuser/mw/v1/subscriber"
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

type CreateSubscriberRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AppID          string  `protobuf:"bytes,10,opt,name=AppID,proto3" json:"AppID,omitempty"`
	EmailAddress   string  `protobuf:"bytes,20,opt,name=EmailAddress,proto3" json:"EmailAddress,omitempty"`
	SubscribeAppID *string `protobuf:"bytes,30,opt,name=SubscribeAppID,proto3,oneof" json:"SubscribeAppID,omitempty"`
}

func (x *CreateSubscriberRequest) Reset() {
	*x = CreateSubscriberRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_appuser_gw_v1_subscriber_subscriber_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateSubscriberRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSubscriberRequest) ProtoMessage() {}

func (x *CreateSubscriberRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_appuser_gw_v1_subscriber_subscriber_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSubscriberRequest.ProtoReflect.Descriptor instead.
func (*CreateSubscriberRequest) Descriptor() ([]byte, []int) {
	return file_npool_appuser_gw_v1_subscriber_subscriber_proto_rawDescGZIP(), []int{0}
}

func (x *CreateSubscriberRequest) GetAppID() string {
	if x != nil {
		return x.AppID
	}
	return ""
}

func (x *CreateSubscriberRequest) GetEmailAddress() string {
	if x != nil {
		return x.EmailAddress
	}
	return ""
}

func (x *CreateSubscriberRequest) GetSubscribeAppID() string {
	if x != nil && x.SubscribeAppID != nil {
		return *x.SubscribeAppID
	}
	return ""
}

type CreateSubscriberResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *subscriber.Subscriber `protobuf:"bytes,10,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (x *CreateSubscriberResponse) Reset() {
	*x = CreateSubscriberResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_appuser_gw_v1_subscriber_subscriber_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateSubscriberResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSubscriberResponse) ProtoMessage() {}

func (x *CreateSubscriberResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_appuser_gw_v1_subscriber_subscriber_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSubscriberResponse.ProtoReflect.Descriptor instead.
func (*CreateSubscriberResponse) Descriptor() ([]byte, []int) {
	return file_npool_appuser_gw_v1_subscriber_subscriber_proto_rawDescGZIP(), []int{1}
}

func (x *CreateSubscriberResponse) GetInfo() *subscriber.Subscriber {
	if x != nil {
		return x.Info
	}
	return nil
}

type GetSubscriberesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AppID  string `protobuf:"bytes,10,opt,name=AppID,proto3" json:"AppID,omitempty"`
	Offset int32  `protobuf:"varint,20,opt,name=Offset,proto3" json:"Offset,omitempty"`
	Limit  int32  `protobuf:"varint,30,opt,name=Limit,proto3" json:"Limit,omitempty"`
}

func (x *GetSubscriberesRequest) Reset() {
	*x = GetSubscriberesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_appuser_gw_v1_subscriber_subscriber_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSubscriberesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSubscriberesRequest) ProtoMessage() {}

func (x *GetSubscriberesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_appuser_gw_v1_subscriber_subscriber_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSubscriberesRequest.ProtoReflect.Descriptor instead.
func (*GetSubscriberesRequest) Descriptor() ([]byte, []int) {
	return file_npool_appuser_gw_v1_subscriber_subscriber_proto_rawDescGZIP(), []int{2}
}

func (x *GetSubscriberesRequest) GetAppID() string {
	if x != nil {
		return x.AppID
	}
	return ""
}

func (x *GetSubscriberesRequest) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *GetSubscriberesRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type GetSubscriberesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Infos []*subscriber.Subscriber `protobuf:"bytes,10,rep,name=Infos,proto3" json:"Infos,omitempty"`
	Total uint32                   `protobuf:"varint,20,opt,name=Total,proto3" json:"Total,omitempty"`
}

func (x *GetSubscriberesResponse) Reset() {
	*x = GetSubscriberesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_appuser_gw_v1_subscriber_subscriber_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSubscriberesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSubscriberesResponse) ProtoMessage() {}

func (x *GetSubscriberesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_appuser_gw_v1_subscriber_subscriber_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSubscriberesResponse.ProtoReflect.Descriptor instead.
func (*GetSubscriberesResponse) Descriptor() ([]byte, []int) {
	return file_npool_appuser_gw_v1_subscriber_subscriber_proto_rawDescGZIP(), []int{3}
}

func (x *GetSubscriberesResponse) GetInfos() []*subscriber.Subscriber {
	if x != nil {
		return x.Infos
	}
	return nil
}

func (x *GetSubscriberesResponse) GetTotal() uint32 {
	if x != nil {
		return x.Total
	}
	return 0
}

type DeleteSubscriberRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AppID        string `protobuf:"bytes,10,opt,name=AppID,proto3" json:"AppID,omitempty"`
	EmailAddress string `protobuf:"bytes,20,opt,name=EmailAddress,proto3" json:"EmailAddress,omitempty"`
}

func (x *DeleteSubscriberRequest) Reset() {
	*x = DeleteSubscriberRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_appuser_gw_v1_subscriber_subscriber_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteSubscriberRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteSubscriberRequest) ProtoMessage() {}

func (x *DeleteSubscriberRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_appuser_gw_v1_subscriber_subscriber_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteSubscriberRequest.ProtoReflect.Descriptor instead.
func (*DeleteSubscriberRequest) Descriptor() ([]byte, []int) {
	return file_npool_appuser_gw_v1_subscriber_subscriber_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteSubscriberRequest) GetAppID() string {
	if x != nil {
		return x.AppID
	}
	return ""
}

func (x *DeleteSubscriberRequest) GetEmailAddress() string {
	if x != nil {
		return x.EmailAddress
	}
	return ""
}

type DeleteSubscriberResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *subscriber.Subscriber `protobuf:"bytes,10,opt,name=Info,proto3,oneof" json:"Info,omitempty"`
}

func (x *DeleteSubscriberResponse) Reset() {
	*x = DeleteSubscriberResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_appuser_gw_v1_subscriber_subscriber_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteSubscriberResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteSubscriberResponse) ProtoMessage() {}

func (x *DeleteSubscriberResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_appuser_gw_v1_subscriber_subscriber_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteSubscriberResponse.ProtoReflect.Descriptor instead.
func (*DeleteSubscriberResponse) Descriptor() ([]byte, []int) {
	return file_npool_appuser_gw_v1_subscriber_subscriber_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteSubscriberResponse) GetInfo() *subscriber.Subscriber {
	if x != nil {
		return x.Info
	}
	return nil
}

var File_npool_appuser_gw_v1_subscriber_subscriber_proto protoreflect.FileDescriptor

var file_npool_appuser_gw_v1_subscriber_subscriber_proto_rawDesc = []byte{
	0x0a, 0x2f, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x61, 0x70, 0x70, 0x75, 0x73, 0x65, 0x72, 0x2f,
	0x67, 0x77, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x72,
	0x2f, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x1d, 0x61, 0x70, 0x70, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77,
	0x61, 0x79, 0x2e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x72, 0x2e, 0x76, 0x31,
	0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2f,
	0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x61, 0x70, 0x70, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x6d, 0x77,
	0x2f, 0x76, 0x31, 0x2f, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x72, 0x2f, 0x73,
	0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x93, 0x01, 0x0a, 0x17, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72,
	0x69, 0x62, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x41,
	0x70, 0x70, 0x49, 0x44, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x41, 0x70, 0x70, 0x49,
	0x44, 0x12, 0x22, 0x0a, 0x0c, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x41, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x2b, 0x0a, 0x0e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69,
	0x62, 0x65, 0x41, 0x70, 0x70, 0x49, 0x44, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52,
	0x0e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x41, 0x70, 0x70, 0x49, 0x44, 0x88,
	0x01, 0x01, 0x42, 0x11, 0x0a, 0x0f, 0x5f, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65,
	0x41, 0x70, 0x70, 0x49, 0x44, 0x22, 0x5c, 0x0a, 0x18, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53,
	0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x40, 0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x2c, 0x2e, 0x61, 0x70, 0x70, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65,
	0x77, 0x61, 0x72, 0x65, 0x2e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x72, 0x2e,
	0x76, 0x31, 0x2e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x72, 0x52, 0x04, 0x49,
	0x6e, 0x66, 0x6f, 0x22, 0x5c, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72,
	0x69, 0x62, 0x65, 0x72, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a,
	0x05, 0x41, 0x70, 0x70, 0x49, 0x44, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x41, 0x70,
	0x70, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x14, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x06, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x4c,
	0x69, 0x6d, 0x69, 0x74, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x4c, 0x69, 0x6d, 0x69,
	0x74, 0x22, 0x73, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62,
	0x65, 0x72, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x42, 0x0a, 0x05,
	0x49, 0x6e, 0x66, 0x6f, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x61, 0x70,
	0x70, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65,
	0x2e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x53,
	0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x72, 0x52, 0x05, 0x49, 0x6e, 0x66, 0x6f, 0x73,
	0x12, 0x14, 0x0a, 0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x14, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x22, 0x53, 0x0a, 0x17, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x14, 0x0a, 0x05, 0x41, 0x70, 0x70, 0x49, 0x44, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x41, 0x70, 0x70, 0x49, 0x44, 0x12, 0x22, 0x0a, 0x0c, 0x45, 0x6d, 0x61, 0x69, 0x6c,
	0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x45,
	0x6d, 0x61, 0x69, 0x6c, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0x6a, 0x0a, 0x18, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x72, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x45, 0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x61, 0x70, 0x70, 0x75, 0x73, 0x65, 0x72, 0x2e,
	0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x73, 0x75, 0x62, 0x73, 0x63,
	0x72, 0x69, 0x62, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69,
	0x62, 0x65, 0x72, 0x48, 0x00, 0x52, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x88, 0x01, 0x01, 0x42, 0x07,
	0x0a, 0x05, 0x5f, 0x49, 0x6e, 0x66, 0x6f, 0x32, 0xfd, 0x03, 0x0a, 0x07, 0x47, 0x61, 0x74, 0x65,
	0x77, 0x61, 0x79, 0x12, 0xa5, 0x01, 0x0a, 0x10, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x75,
	0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x72, 0x12, 0x36, 0x2e, 0x61, 0x70, 0x70, 0x75, 0x73,
	0x65, 0x72, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x73, 0x75, 0x62, 0x73, 0x63,
	0x72, 0x69, 0x62, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53,
	0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x37, 0x2e, 0x61, 0x70, 0x70, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77,
	0x61, 0x79, 0x2e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x72, 0x2e, 0x76, 0x31,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x20, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x1a, 0x3a, 0x01, 0x2a, 0x22, 0x15, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x2f, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x72, 0x12, 0xa1, 0x01, 0x0a, 0x0f,
	0x47, 0x65, 0x74, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x72, 0x65, 0x73, 0x12,
	0x35, 0x2e, 0x61, 0x70, 0x70, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61,
	0x79, 0x2e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e,
	0x47, 0x65, 0x74, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x72, 0x65, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x36, 0x2e, 0x61, 0x70, 0x70, 0x75, 0x73, 0x65, 0x72,
	0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69,
	0x62, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72,
	0x69, 0x62, 0x65, 0x72, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1f,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x19, 0x3a, 0x01, 0x2a, 0x22, 0x14, 0x2f, 0x76, 0x31, 0x2f, 0x67,
	0x65, 0x74, 0x2f, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x72, 0x65, 0x73, 0x12,
	0xa5, 0x01, 0x0a, 0x10, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72,
	0x69, 0x62, 0x65, 0x72, 0x12, 0x36, 0x2e, 0x61, 0x70, 0x70, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x67,
	0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65,
	0x72, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x75, 0x62, 0x73, 0x63,
	0x72, 0x69, 0x62, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x37, 0x2e, 0x61,
	0x70, 0x70, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x73,
	0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x72, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x20, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1a, 0x3a, 0x01, 0x2a,
	0x22, 0x15, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x2f, 0x73, 0x75, 0x62,
	0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x72, 0x42, 0x41, 0x5a, 0x3f, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4e, 0x70, 0x6f, 0x6f, 0x6c, 0x50, 0x6c, 0x61, 0x74, 0x66,
	0x6f, 0x72, 0x6d, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x6e, 0x70, 0x6f, 0x6f,
	0x6c, 0x2f, 0x61, 0x70, 0x70, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x67, 0x77, 0x2f, 0x76, 0x31, 0x2f,
	0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_npool_appuser_gw_v1_subscriber_subscriber_proto_rawDescOnce sync.Once
	file_npool_appuser_gw_v1_subscriber_subscriber_proto_rawDescData = file_npool_appuser_gw_v1_subscriber_subscriber_proto_rawDesc
)

func file_npool_appuser_gw_v1_subscriber_subscriber_proto_rawDescGZIP() []byte {
	file_npool_appuser_gw_v1_subscriber_subscriber_proto_rawDescOnce.Do(func() {
		file_npool_appuser_gw_v1_subscriber_subscriber_proto_rawDescData = protoimpl.X.CompressGZIP(file_npool_appuser_gw_v1_subscriber_subscriber_proto_rawDescData)
	})
	return file_npool_appuser_gw_v1_subscriber_subscriber_proto_rawDescData
}

var file_npool_appuser_gw_v1_subscriber_subscriber_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_npool_appuser_gw_v1_subscriber_subscriber_proto_goTypes = []interface{}{
	(*CreateSubscriberRequest)(nil),  // 0: appuser.gateway.subscriber.v1.CreateSubscriberRequest
	(*CreateSubscriberResponse)(nil), // 1: appuser.gateway.subscriber.v1.CreateSubscriberResponse
	(*GetSubscriberesRequest)(nil),   // 2: appuser.gateway.subscriber.v1.GetSubscriberesRequest
	(*GetSubscriberesResponse)(nil),  // 3: appuser.gateway.subscriber.v1.GetSubscriberesResponse
	(*DeleteSubscriberRequest)(nil),  // 4: appuser.gateway.subscriber.v1.DeleteSubscriberRequest
	(*DeleteSubscriberResponse)(nil), // 5: appuser.gateway.subscriber.v1.DeleteSubscriberResponse
	(*subscriber.Subscriber)(nil),    // 6: appuser.middleware.subscriber.v1.Subscriber
}
var file_npool_appuser_gw_v1_subscriber_subscriber_proto_depIdxs = []int32{
	6, // 0: appuser.gateway.subscriber.v1.CreateSubscriberResponse.Info:type_name -> appuser.middleware.subscriber.v1.Subscriber
	6, // 1: appuser.gateway.subscriber.v1.GetSubscriberesResponse.Infos:type_name -> appuser.middleware.subscriber.v1.Subscriber
	6, // 2: appuser.gateway.subscriber.v1.DeleteSubscriberResponse.Info:type_name -> appuser.middleware.subscriber.v1.Subscriber
	0, // 3: appuser.gateway.subscriber.v1.Gateway.CreateSubscriber:input_type -> appuser.gateway.subscriber.v1.CreateSubscriberRequest
	2, // 4: appuser.gateway.subscriber.v1.Gateway.GetSubscriberes:input_type -> appuser.gateway.subscriber.v1.GetSubscriberesRequest
	4, // 5: appuser.gateway.subscriber.v1.Gateway.DeleteSubscriber:input_type -> appuser.gateway.subscriber.v1.DeleteSubscriberRequest
	1, // 6: appuser.gateway.subscriber.v1.Gateway.CreateSubscriber:output_type -> appuser.gateway.subscriber.v1.CreateSubscriberResponse
	3, // 7: appuser.gateway.subscriber.v1.Gateway.GetSubscriberes:output_type -> appuser.gateway.subscriber.v1.GetSubscriberesResponse
	5, // 8: appuser.gateway.subscriber.v1.Gateway.DeleteSubscriber:output_type -> appuser.gateway.subscriber.v1.DeleteSubscriberResponse
	6, // [6:9] is the sub-list for method output_type
	3, // [3:6] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_npool_appuser_gw_v1_subscriber_subscriber_proto_init() }
func file_npool_appuser_gw_v1_subscriber_subscriber_proto_init() {
	if File_npool_appuser_gw_v1_subscriber_subscriber_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_npool_appuser_gw_v1_subscriber_subscriber_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateSubscriberRequest); i {
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
		file_npool_appuser_gw_v1_subscriber_subscriber_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateSubscriberResponse); i {
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
		file_npool_appuser_gw_v1_subscriber_subscriber_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSubscriberesRequest); i {
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
		file_npool_appuser_gw_v1_subscriber_subscriber_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSubscriberesResponse); i {
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
		file_npool_appuser_gw_v1_subscriber_subscriber_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteSubscriberRequest); i {
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
		file_npool_appuser_gw_v1_subscriber_subscriber_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteSubscriberResponse); i {
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
	file_npool_appuser_gw_v1_subscriber_subscriber_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_npool_appuser_gw_v1_subscriber_subscriber_proto_msgTypes[5].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_npool_appuser_gw_v1_subscriber_subscriber_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_npool_appuser_gw_v1_subscriber_subscriber_proto_goTypes,
		DependencyIndexes: file_npool_appuser_gw_v1_subscriber_subscriber_proto_depIdxs,
		MessageInfos:      file_npool_appuser_gw_v1_subscriber_subscriber_proto_msgTypes,
	}.Build()
	File_npool_appuser_gw_v1_subscriber_subscriber_proto = out.File
	file_npool_appuser_gw_v1_subscriber_subscriber_proto_rawDesc = nil
	file_npool_appuser_gw_v1_subscriber_subscriber_proto_goTypes = nil
	file_npool_appuser_gw_v1_subscriber_subscriber_proto_depIdxs = nil
}
