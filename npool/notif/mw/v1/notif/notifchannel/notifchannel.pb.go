// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.18.1
// source: npool/notif/mw/v1/notif/notifchannel/notifchannel.proto

package notifchannel

import (
	notifchannel "github.com/NpoolPlatform/message/npool/notif/mgr/v1/notif/notifchannel"
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

type GetNotifChannelsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Conds  *notifchannel.Conds `protobuf:"bytes,10,opt,name=Conds,proto3" json:"Conds,omitempty"`
	Offset int32               `protobuf:"varint,20,opt,name=Offset,proto3" json:"Offset,omitempty"`
	Limit  int32               `protobuf:"varint,30,opt,name=Limit,proto3" json:"Limit,omitempty"`
}

func (x *GetNotifChannelsRequest) Reset() {
	*x = GetNotifChannelsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_notif_mw_v1_notif_notifchannel_notifchannel_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetNotifChannelsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetNotifChannelsRequest) ProtoMessage() {}

func (x *GetNotifChannelsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_notif_mw_v1_notif_notifchannel_notifchannel_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetNotifChannelsRequest.ProtoReflect.Descriptor instead.
func (*GetNotifChannelsRequest) Descriptor() ([]byte, []int) {
	return file_npool_notif_mw_v1_notif_notifchannel_notifchannel_proto_rawDescGZIP(), []int{0}
}

func (x *GetNotifChannelsRequest) GetConds() *notifchannel.Conds {
	if x != nil {
		return x.Conds
	}
	return nil
}

func (x *GetNotifChannelsRequest) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *GetNotifChannelsRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type GetNotifChannelsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Infos []*notifchannel.NotifChannel `protobuf:"bytes,10,rep,name=Infos,proto3" json:"Infos,omitempty"`
	Total uint32                       `protobuf:"varint,20,opt,name=Total,proto3" json:"Total,omitempty"`
}

func (x *GetNotifChannelsResponse) Reset() {
	*x = GetNotifChannelsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_notif_mw_v1_notif_notifchannel_notifchannel_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetNotifChannelsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetNotifChannelsResponse) ProtoMessage() {}

func (x *GetNotifChannelsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_notif_mw_v1_notif_notifchannel_notifchannel_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetNotifChannelsResponse.ProtoReflect.Descriptor instead.
func (*GetNotifChannelsResponse) Descriptor() ([]byte, []int) {
	return file_npool_notif_mw_v1_notif_notifchannel_notifchannel_proto_rawDescGZIP(), []int{1}
}

func (x *GetNotifChannelsResponse) GetInfos() []*notifchannel.NotifChannel {
	if x != nil {
		return x.Infos
	}
	return nil
}

func (x *GetNotifChannelsResponse) GetTotal() uint32 {
	if x != nil {
		return x.Total
	}
	return 0
}

type GetNotifChannelOnlyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Conds *notifchannel.Conds `protobuf:"bytes,10,opt,name=Conds,proto3" json:"Conds,omitempty"`
}

func (x *GetNotifChannelOnlyRequest) Reset() {
	*x = GetNotifChannelOnlyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_notif_mw_v1_notif_notifchannel_notifchannel_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetNotifChannelOnlyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetNotifChannelOnlyRequest) ProtoMessage() {}

func (x *GetNotifChannelOnlyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_notif_mw_v1_notif_notifchannel_notifchannel_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetNotifChannelOnlyRequest.ProtoReflect.Descriptor instead.
func (*GetNotifChannelOnlyRequest) Descriptor() ([]byte, []int) {
	return file_npool_notif_mw_v1_notif_notifchannel_notifchannel_proto_rawDescGZIP(), []int{2}
}

func (x *GetNotifChannelOnlyRequest) GetConds() *notifchannel.Conds {
	if x != nil {
		return x.Conds
	}
	return nil
}

type GetNotifChannelOnlyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *notifchannel.NotifChannel `protobuf:"bytes,10,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (x *GetNotifChannelOnlyResponse) Reset() {
	*x = GetNotifChannelOnlyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_notif_mw_v1_notif_notifchannel_notifchannel_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetNotifChannelOnlyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetNotifChannelOnlyResponse) ProtoMessage() {}

func (x *GetNotifChannelOnlyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_notif_mw_v1_notif_notifchannel_notifchannel_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetNotifChannelOnlyResponse.ProtoReflect.Descriptor instead.
func (*GetNotifChannelOnlyResponse) Descriptor() ([]byte, []int) {
	return file_npool_notif_mw_v1_notif_notifchannel_notifchannel_proto_rawDescGZIP(), []int{3}
}

func (x *GetNotifChannelOnlyResponse) GetInfo() *notifchannel.NotifChannel {
	if x != nil {
		return x.Info
	}
	return nil
}

var File_npool_notif_mw_v1_notif_notifchannel_notifchannel_proto protoreflect.FileDescriptor

var file_npool_notif_mw_v1_notif_notifchannel_notifchannel_proto_rawDesc = []byte{
	0x0a, 0x37, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x2f, 0x6d, 0x77,
	0x2f, 0x76, 0x31, 0x2f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x2f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x63,
	0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x2f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x63, 0x68, 0x61, 0x6e,
	0x6e, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x21, 0x6e, 0x6f, 0x74, 0x69, 0x66,
	0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x6e, 0x6f, 0x74, 0x69,
	0x66, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x31, 0x2e, 0x76, 0x31, 0x1a, 0x38, 0x6e, 0x70,
	0x6f, 0x6f, 0x6c, 0x2f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x2f, 0x6d, 0x67, 0x72, 0x2f, 0x76, 0x31,
	0x2f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x2f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x63, 0x68, 0x61, 0x6e,
	0x6e, 0x65, 0x6c, 0x2f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8a, 0x01, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x4e, 0x6f,
	0x74, 0x69, 0x66, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x41, 0x0a, 0x05, 0x43, 0x6f, 0x6e, 0x64, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x2b, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65,
	0x72, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x63, 0x68, 0x61,
	0x6e, 0x6e, 0x65, 0x6c, 0x31, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x64, 0x73, 0x52, 0x05,
	0x43, 0x6f, 0x6e, 0x64, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18,
	0x14, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x14, 0x0a,
	0x05, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x4c, 0x69,
	0x6d, 0x69, 0x74, 0x22, 0x7a, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x43,
	0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x48, 0x0a, 0x05, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x32,
	0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x6e,
	0x6f, 0x74, 0x69, 0x66, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65,
	0x6c, 0x31, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x43, 0x68, 0x61, 0x6e, 0x6e,
	0x65, 0x6c, 0x52, 0x05, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x6f, 0x74,
	0x61, 0x6c, 0x18, 0x14, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x22,
	0x5f, 0x0a, 0x1a, 0x47, 0x65, 0x74, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x43, 0x68, 0x61, 0x6e, 0x6e,
	0x65, 0x6c, 0x4f, 0x6e, 0x6c, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x41, 0x0a,
	0x05, 0x43, 0x6f, 0x6e, 0x64, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x6e,
	0x6f, 0x74, 0x69, 0x66, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x6e, 0x6f, 0x74,
	0x69, 0x66, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x31,
	0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x64, 0x73, 0x52, 0x05, 0x43, 0x6f, 0x6e, 0x64, 0x73,
	0x22, 0x65, 0x0a, 0x1b, 0x47, 0x65, 0x74, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x43, 0x68, 0x61, 0x6e,
	0x6e, 0x65, 0x6c, 0x4f, 0x6e, 0x6c, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x46, 0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x32, 0x2e,
	0x6e, 0x6f, 0x74, 0x69, 0x66, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x6e, 0x6f,
	0x74, 0x69, 0x66, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c,
	0x31, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65,
	0x6c, 0x52, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x32, 0xb5, 0x02, 0x0a, 0x0a, 0x4d, 0x69, 0x64, 0x64,
	0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x12, 0x8d, 0x01, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x4e, 0x6f,
	0x74, 0x69, 0x66, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x73, 0x12, 0x3a, 0x2e, 0x6e, 0x6f,
	0x74, 0x69, 0x66, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x6e,
	0x6f, 0x74, 0x69, 0x66, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x31, 0x2e, 0x76, 0x31, 0x2e,
	0x47, 0x65, 0x74, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3b, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x2e,
	0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66,
	0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x31, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4e,
	0x6f, 0x74, 0x69, 0x66, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x96, 0x01, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x4e, 0x6f,
	0x74, 0x69, 0x66, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x4f, 0x6e, 0x6c, 0x79, 0x12, 0x3d,
	0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72,
	0x65, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x31, 0x2e,
	0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x43, 0x68, 0x61, 0x6e, 0x6e,
	0x65, 0x6c, 0x4f, 0x6e, 0x6c, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3e, 0x2e,
	0x6e, 0x6f, 0x74, 0x69, 0x66, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65,
	0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x31, 0x2e, 0x76,
	0x31, 0x2e, 0x47, 0x65, 0x74, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65,
	0x6c, 0x4f, 0x6e, 0x6c, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42,
	0x47, 0x5a, 0x45, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4e, 0x70,
	0x6f, 0x6f, 0x6c, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x2f, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x2f,
	0x6d, 0x77, 0x2f, 0x76, 0x31, 0x2f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x2f, 0x6e, 0x6f, 0x74, 0x69,
	0x66, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_npool_notif_mw_v1_notif_notifchannel_notifchannel_proto_rawDescOnce sync.Once
	file_npool_notif_mw_v1_notif_notifchannel_notifchannel_proto_rawDescData = file_npool_notif_mw_v1_notif_notifchannel_notifchannel_proto_rawDesc
)

func file_npool_notif_mw_v1_notif_notifchannel_notifchannel_proto_rawDescGZIP() []byte {
	file_npool_notif_mw_v1_notif_notifchannel_notifchannel_proto_rawDescOnce.Do(func() {
		file_npool_notif_mw_v1_notif_notifchannel_notifchannel_proto_rawDescData = protoimpl.X.CompressGZIP(file_npool_notif_mw_v1_notif_notifchannel_notifchannel_proto_rawDescData)
	})
	return file_npool_notif_mw_v1_notif_notifchannel_notifchannel_proto_rawDescData
}

var file_npool_notif_mw_v1_notif_notifchannel_notifchannel_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_npool_notif_mw_v1_notif_notifchannel_notifchannel_proto_goTypes = []interface{}{
	(*GetNotifChannelsRequest)(nil),     // 0: notif.middleware.notifchannel1.v1.GetNotifChannelsRequest
	(*GetNotifChannelsResponse)(nil),    // 1: notif.middleware.notifchannel1.v1.GetNotifChannelsResponse
	(*GetNotifChannelOnlyRequest)(nil),  // 2: notif.middleware.notifchannel1.v1.GetNotifChannelOnlyRequest
	(*GetNotifChannelOnlyResponse)(nil), // 3: notif.middleware.notifchannel1.v1.GetNotifChannelOnlyResponse
	(*notifchannel.Conds)(nil),          // 4: notif.manager.notif.notifchannel1.v1.Conds
	(*notifchannel.NotifChannel)(nil),   // 5: notif.manager.notif.notifchannel1.v1.NotifChannel
}
var file_npool_notif_mw_v1_notif_notifchannel_notifchannel_proto_depIdxs = []int32{
	4, // 0: notif.middleware.notifchannel1.v1.GetNotifChannelsRequest.Conds:type_name -> notif.manager.notif.notifchannel1.v1.Conds
	5, // 1: notif.middleware.notifchannel1.v1.GetNotifChannelsResponse.Infos:type_name -> notif.manager.notif.notifchannel1.v1.NotifChannel
	4, // 2: notif.middleware.notifchannel1.v1.GetNotifChannelOnlyRequest.Conds:type_name -> notif.manager.notif.notifchannel1.v1.Conds
	5, // 3: notif.middleware.notifchannel1.v1.GetNotifChannelOnlyResponse.Info:type_name -> notif.manager.notif.notifchannel1.v1.NotifChannel
	0, // 4: notif.middleware.notifchannel1.v1.Middleware.GetNotifChannels:input_type -> notif.middleware.notifchannel1.v1.GetNotifChannelsRequest
	2, // 5: notif.middleware.notifchannel1.v1.Middleware.GetNotifChannelOnly:input_type -> notif.middleware.notifchannel1.v1.GetNotifChannelOnlyRequest
	1, // 6: notif.middleware.notifchannel1.v1.Middleware.GetNotifChannels:output_type -> notif.middleware.notifchannel1.v1.GetNotifChannelsResponse
	3, // 7: notif.middleware.notifchannel1.v1.Middleware.GetNotifChannelOnly:output_type -> notif.middleware.notifchannel1.v1.GetNotifChannelOnlyResponse
	6, // [6:8] is the sub-list for method output_type
	4, // [4:6] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_npool_notif_mw_v1_notif_notifchannel_notifchannel_proto_init() }
func file_npool_notif_mw_v1_notif_notifchannel_notifchannel_proto_init() {
	if File_npool_notif_mw_v1_notif_notifchannel_notifchannel_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_npool_notif_mw_v1_notif_notifchannel_notifchannel_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetNotifChannelsRequest); i {
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
		file_npool_notif_mw_v1_notif_notifchannel_notifchannel_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetNotifChannelsResponse); i {
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
		file_npool_notif_mw_v1_notif_notifchannel_notifchannel_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetNotifChannelOnlyRequest); i {
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
		file_npool_notif_mw_v1_notif_notifchannel_notifchannel_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetNotifChannelOnlyResponse); i {
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
			RawDescriptor: file_npool_notif_mw_v1_notif_notifchannel_notifchannel_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_npool_notif_mw_v1_notif_notifchannel_notifchannel_proto_goTypes,
		DependencyIndexes: file_npool_notif_mw_v1_notif_notifchannel_notifchannel_proto_depIdxs,
		MessageInfos:      file_npool_notif_mw_v1_notif_notifchannel_notifchannel_proto_msgTypes,
	}.Build()
	File_npool_notif_mw_v1_notif_notifchannel_notifchannel_proto = out.File
	file_npool_notif_mw_v1_notif_notifchannel_notifchannel_proto_rawDesc = nil
	file_npool_notif_mw_v1_notif_notifchannel_notifchannel_proto_goTypes = nil
	file_npool_notif_mw_v1_notif_notifchannel_notifchannel_proto_depIdxs = nil
}
