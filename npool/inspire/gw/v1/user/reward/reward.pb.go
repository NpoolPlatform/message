// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.18.1
// source: npool/inspire/gw/v1/user/reward/reward.proto

package reward

import (
	reward "github.com/NpoolPlatform/message/npool/inspire/mw/v1/user/reward"
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

type GetMyRewardsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AppID  string `protobuf:"bytes,10,opt,name=AppID,proto3" json:"AppID,omitempty"`
	UserID string `protobuf:"bytes,20,opt,name=UserID,proto3" json:"UserID,omitempty"`
	Offset int32  `protobuf:"varint,30,opt,name=Offset,proto3" json:"Offset,omitempty"`
	Limit  int32  `protobuf:"varint,40,opt,name=Limit,proto3" json:"Limit,omitempty"`
}

func (x *GetMyRewardsRequest) Reset() {
	*x = GetMyRewardsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_inspire_gw_v1_user_reward_reward_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMyRewardsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMyRewardsRequest) ProtoMessage() {}

func (x *GetMyRewardsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_inspire_gw_v1_user_reward_reward_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMyRewardsRequest.ProtoReflect.Descriptor instead.
func (*GetMyRewardsRequest) Descriptor() ([]byte, []int) {
	return file_npool_inspire_gw_v1_user_reward_reward_proto_rawDescGZIP(), []int{0}
}

func (x *GetMyRewardsRequest) GetAppID() string {
	if x != nil {
		return x.AppID
	}
	return ""
}

func (x *GetMyRewardsRequest) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

func (x *GetMyRewardsRequest) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *GetMyRewardsRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type GetMyRewardsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Infos []*reward.UserReward `protobuf:"bytes,10,rep,name=Infos,proto3" json:"Infos,omitempty"`
	Total uint32               `protobuf:"varint,20,opt,name=Total,proto3" json:"Total,omitempty"`
}

func (x *GetMyRewardsResponse) Reset() {
	*x = GetMyRewardsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_inspire_gw_v1_user_reward_reward_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMyRewardsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMyRewardsResponse) ProtoMessage() {}

func (x *GetMyRewardsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_inspire_gw_v1_user_reward_reward_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMyRewardsResponse.ProtoReflect.Descriptor instead.
func (*GetMyRewardsResponse) Descriptor() ([]byte, []int) {
	return file_npool_inspire_gw_v1_user_reward_reward_proto_rawDescGZIP(), []int{1}
}

func (x *GetMyRewardsResponse) GetInfos() []*reward.UserReward {
	if x != nil {
		return x.Infos
	}
	return nil
}

func (x *GetMyRewardsResponse) GetTotal() uint32 {
	if x != nil {
		return x.Total
	}
	return 0
}

type AdminGetUserRewardsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TargetAppID string `protobuf:"bytes,10,opt,name=TargetAppID,proto3" json:"TargetAppID,omitempty"`
	Offset      int32  `protobuf:"varint,30,opt,name=Offset,proto3" json:"Offset,omitempty"`
	Limit       int32  `protobuf:"varint,40,opt,name=Limit,proto3" json:"Limit,omitempty"`
}

func (x *AdminGetUserRewardsRequest) Reset() {
	*x = AdminGetUserRewardsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_inspire_gw_v1_user_reward_reward_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdminGetUserRewardsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdminGetUserRewardsRequest) ProtoMessage() {}

func (x *AdminGetUserRewardsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_inspire_gw_v1_user_reward_reward_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdminGetUserRewardsRequest.ProtoReflect.Descriptor instead.
func (*AdminGetUserRewardsRequest) Descriptor() ([]byte, []int) {
	return file_npool_inspire_gw_v1_user_reward_reward_proto_rawDescGZIP(), []int{2}
}

func (x *AdminGetUserRewardsRequest) GetTargetAppID() string {
	if x != nil {
		return x.TargetAppID
	}
	return ""
}

func (x *AdminGetUserRewardsRequest) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *AdminGetUserRewardsRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type AdminGetUserRewardsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Infos []*reward.UserReward `protobuf:"bytes,10,rep,name=Infos,proto3" json:"Infos,omitempty"`
	Total uint32               `protobuf:"varint,20,opt,name=Total,proto3" json:"Total,omitempty"`
}

func (x *AdminGetUserRewardsResponse) Reset() {
	*x = AdminGetUserRewardsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_inspire_gw_v1_user_reward_reward_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdminGetUserRewardsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdminGetUserRewardsResponse) ProtoMessage() {}

func (x *AdminGetUserRewardsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_inspire_gw_v1_user_reward_reward_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdminGetUserRewardsResponse.ProtoReflect.Descriptor instead.
func (*AdminGetUserRewardsResponse) Descriptor() ([]byte, []int) {
	return file_npool_inspire_gw_v1_user_reward_reward_proto_rawDescGZIP(), []int{3}
}

func (x *AdminGetUserRewardsResponse) GetInfos() []*reward.UserReward {
	if x != nil {
		return x.Infos
	}
	return nil
}

func (x *AdminGetUserRewardsResponse) GetTotal() uint32 {
	if x != nil {
		return x.Total
	}
	return 0
}

var File_npool_inspire_gw_v1_user_reward_reward_proto protoreflect.FileDescriptor

var file_npool_inspire_gw_v1_user_reward_reward_proto_rawDesc = []byte{
	0x0a, 0x2c, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x69, 0x6e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2f,
	0x67, 0x77, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x72, 0x65, 0x77, 0x61, 0x72,
	0x64, 0x2f, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1e,
	0x69, 0x6e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x2e, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x2e, 0x76, 0x31, 0x1a, 0x1c,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2c, 0x6e, 0x70,
	0x6f, 0x6f, 0x6c, 0x2f, 0x69, 0x6e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2f, 0x6d, 0x77, 0x2f, 0x76,
	0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x2f, 0x72, 0x65,
	0x77, 0x61, 0x72, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x71, 0x0a, 0x13, 0x47, 0x65,
	0x74, 0x4d, 0x79, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x14, 0x0a, 0x05, 0x41, 0x70, 0x70, 0x49, 0x44, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x41, 0x70, 0x70, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49,
	0x44, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12,
	0x16, 0x0a, 0x06, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x06, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x4c, 0x69, 0x6d, 0x69, 0x74,
	0x18, 0x28, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x22, 0x71, 0x0a,
	0x14, 0x47, 0x65, 0x74, 0x4d, 0x79, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x43, 0x0a, 0x05, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x18, 0x0a,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x69, 0x6e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x6d,
	0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x72,
	0x65, 0x77, 0x61, 0x72, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x77,
	0x61, 0x72, 0x64, 0x52, 0x05, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x6f,
	0x74, 0x61, 0x6c, 0x18, 0x14, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c,
	0x22, 0x6c, 0x0a, 0x1a, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72,
	0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x20,
	0x0a, 0x0b, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x41, 0x70, 0x70, 0x49, 0x44, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x41, 0x70, 0x70, 0x49, 0x44,
	0x12, 0x16, 0x0a, 0x06, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x06, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x4c, 0x69, 0x6d, 0x69,
	0x74, 0x18, 0x28, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x22, 0x78,
	0x0a, 0x1b, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65,
	0x77, 0x61, 0x72, 0x64, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x43, 0x0a,
	0x05, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x69,
	0x6e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72,
	0x65, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x2e, 0x76, 0x31,
	0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x52, 0x05, 0x49, 0x6e, 0x66,
	0x6f, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x14, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x32, 0xdb, 0x02, 0x0a, 0x07, 0x47, 0x61, 0x74,
	0x65, 0x77, 0x61, 0x79, 0x12, 0xb4, 0x01, 0x0a, 0x13, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x47, 0x65,
	0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x73, 0x12, 0x3a, 0x2e, 0x69,
	0x6e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x75,
	0x73, 0x65, 0x72, 0x2e, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x64,
	0x6d, 0x69, 0x6e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3b, 0x2e, 0x69, 0x6e, 0x73, 0x70, 0x69,
	0x72, 0x65, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e,
	0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x47,
	0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x24, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1e, 0x3a, 0x01, 0x2a,
	0x22, 0x19, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x67, 0x65, 0x74, 0x2f,
	0x75, 0x73, 0x65, 0x72, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x73, 0x12, 0x98, 0x01, 0x0a, 0x0c,
	0x47, 0x65, 0x74, 0x4d, 0x79, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x73, 0x12, 0x33, 0x2e, 0x69,
	0x6e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x75,
	0x73, 0x65, 0x72, 0x2e, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65,
	0x74, 0x4d, 0x79, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x34, 0x2e, 0x69, 0x6e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x67, 0x61, 0x74, 0x65,
	0x77, 0x61, 0x79, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x2e,
	0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x79, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x17, 0x3a,
	0x01, 0x2a, 0x22, 0x12, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x65, 0x74, 0x2f, 0x6d, 0x79, 0x2f, 0x72,
	0x65, 0x77, 0x61, 0x72, 0x64, 0x73, 0x42, 0x42, 0x5a, 0x40, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4e, 0x70, 0x6f, 0x6f, 0x6c, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f,
	0x72, 0x6d, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x6e, 0x70, 0x6f, 0x6f, 0x6c,
	0x2f, 0x69, 0x6e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2f, 0x67, 0x77, 0x2f, 0x76, 0x31, 0x2f, 0x75,
	0x73, 0x65, 0x72, 0x2f, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_npool_inspire_gw_v1_user_reward_reward_proto_rawDescOnce sync.Once
	file_npool_inspire_gw_v1_user_reward_reward_proto_rawDescData = file_npool_inspire_gw_v1_user_reward_reward_proto_rawDesc
)

func file_npool_inspire_gw_v1_user_reward_reward_proto_rawDescGZIP() []byte {
	file_npool_inspire_gw_v1_user_reward_reward_proto_rawDescOnce.Do(func() {
		file_npool_inspire_gw_v1_user_reward_reward_proto_rawDescData = protoimpl.X.CompressGZIP(file_npool_inspire_gw_v1_user_reward_reward_proto_rawDescData)
	})
	return file_npool_inspire_gw_v1_user_reward_reward_proto_rawDescData
}

var file_npool_inspire_gw_v1_user_reward_reward_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_npool_inspire_gw_v1_user_reward_reward_proto_goTypes = []interface{}{
	(*GetMyRewardsRequest)(nil),         // 0: inspire.gateway.user.reward.v1.GetMyRewardsRequest
	(*GetMyRewardsResponse)(nil),        // 1: inspire.gateway.user.reward.v1.GetMyRewardsResponse
	(*AdminGetUserRewardsRequest)(nil),  // 2: inspire.gateway.user.reward.v1.AdminGetUserRewardsRequest
	(*AdminGetUserRewardsResponse)(nil), // 3: inspire.gateway.user.reward.v1.AdminGetUserRewardsResponse
	(*reward.UserReward)(nil),           // 4: inspire.middleware.user.reward.v1.UserReward
}
var file_npool_inspire_gw_v1_user_reward_reward_proto_depIdxs = []int32{
	4, // 0: inspire.gateway.user.reward.v1.GetMyRewardsResponse.Infos:type_name -> inspire.middleware.user.reward.v1.UserReward
	4, // 1: inspire.gateway.user.reward.v1.AdminGetUserRewardsResponse.Infos:type_name -> inspire.middleware.user.reward.v1.UserReward
	2, // 2: inspire.gateway.user.reward.v1.Gateway.AdminGetUserRewards:input_type -> inspire.gateway.user.reward.v1.AdminGetUserRewardsRequest
	0, // 3: inspire.gateway.user.reward.v1.Gateway.GetMyRewards:input_type -> inspire.gateway.user.reward.v1.GetMyRewardsRequest
	3, // 4: inspire.gateway.user.reward.v1.Gateway.AdminGetUserRewards:output_type -> inspire.gateway.user.reward.v1.AdminGetUserRewardsResponse
	1, // 5: inspire.gateway.user.reward.v1.Gateway.GetMyRewards:output_type -> inspire.gateway.user.reward.v1.GetMyRewardsResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_npool_inspire_gw_v1_user_reward_reward_proto_init() }
func file_npool_inspire_gw_v1_user_reward_reward_proto_init() {
	if File_npool_inspire_gw_v1_user_reward_reward_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_npool_inspire_gw_v1_user_reward_reward_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMyRewardsRequest); i {
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
		file_npool_inspire_gw_v1_user_reward_reward_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMyRewardsResponse); i {
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
		file_npool_inspire_gw_v1_user_reward_reward_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdminGetUserRewardsRequest); i {
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
		file_npool_inspire_gw_v1_user_reward_reward_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdminGetUserRewardsResponse); i {
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
			RawDescriptor: file_npool_inspire_gw_v1_user_reward_reward_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_npool_inspire_gw_v1_user_reward_reward_proto_goTypes,
		DependencyIndexes: file_npool_inspire_gw_v1_user_reward_reward_proto_depIdxs,
		MessageInfos:      file_npool_inspire_gw_v1_user_reward_reward_proto_msgTypes,
	}.Build()
	File_npool_inspire_gw_v1_user_reward_reward_proto = out.File
	file_npool_inspire_gw_v1_user_reward_reward_proto_rawDesc = nil
	file_npool_inspire_gw_v1_user_reward_reward_proto_goTypes = nil
	file_npool_inspire_gw_v1_user_reward_reward_proto_depIdxs = nil
}
