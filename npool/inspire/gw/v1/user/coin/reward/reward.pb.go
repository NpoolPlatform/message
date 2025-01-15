// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.18.1
// source: npool/inspire/gw/v1/user/coin/reward/reward.proto

package reward

import (
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

type UserCoinReward struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID           uint32   `protobuf:"varint,10,opt,name=ID,proto3" json:"ID,omitempty"`
	EntID        string   `protobuf:"bytes,20,opt,name=EntID,proto3" json:"EntID,omitempty"`
	AppID        string   `protobuf:"bytes,30,opt,name=AppID,proto3" json:"AppID,omitempty"`
	UserID       string   `protobuf:"bytes,40,opt,name=UserID,proto3" json:"UserID,omitempty"`
	CoinTypeID   string   `protobuf:"bytes,50,opt,name=CoinTypeID,proto3" json:"CoinTypeID,omitempty"`
	CoinRewards  string   `protobuf:"bytes,60,opt,name=CoinRewards,proto3" json:"CoinRewards,omitempty"`
	CoinName     string   `protobuf:"bytes,70,opt,name=CoinName,proto3" json:"CoinName,omitempty"`
	DisplayNames []string `protobuf:"bytes,80,rep,name=DisplayNames,proto3" json:"DisplayNames,omitempty"`
	CoinLogo     string   `protobuf:"bytes,90,opt,name=CoinLogo,proto3" json:"CoinLogo,omitempty"`
	CoinUnit     string   `protobuf:"bytes,100,opt,name=CoinUnit,proto3" json:"CoinUnit,omitempty"`
	CreatedAt    uint32   `protobuf:"varint,1000,opt,name=CreatedAt,proto3" json:"CreatedAt,omitempty"`
	UpdatedAt    uint32   `protobuf:"varint,1010,opt,name=UpdatedAt,proto3" json:"UpdatedAt,omitempty"`
}

func (x *UserCoinReward) Reset() {
	*x = UserCoinReward{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_inspire_gw_v1_user_coin_reward_reward_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserCoinReward) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserCoinReward) ProtoMessage() {}

func (x *UserCoinReward) ProtoReflect() protoreflect.Message {
	mi := &file_npool_inspire_gw_v1_user_coin_reward_reward_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserCoinReward.ProtoReflect.Descriptor instead.
func (*UserCoinReward) Descriptor() ([]byte, []int) {
	return file_npool_inspire_gw_v1_user_coin_reward_reward_proto_rawDescGZIP(), []int{0}
}

func (x *UserCoinReward) GetID() uint32 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *UserCoinReward) GetEntID() string {
	if x != nil {
		return x.EntID
	}
	return ""
}

func (x *UserCoinReward) GetAppID() string {
	if x != nil {
		return x.AppID
	}
	return ""
}

func (x *UserCoinReward) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

func (x *UserCoinReward) GetCoinTypeID() string {
	if x != nil {
		return x.CoinTypeID
	}
	return ""
}

func (x *UserCoinReward) GetCoinRewards() string {
	if x != nil {
		return x.CoinRewards
	}
	return ""
}

func (x *UserCoinReward) GetCoinName() string {
	if x != nil {
		return x.CoinName
	}
	return ""
}

func (x *UserCoinReward) GetDisplayNames() []string {
	if x != nil {
		return x.DisplayNames
	}
	return nil
}

func (x *UserCoinReward) GetCoinLogo() string {
	if x != nil {
		return x.CoinLogo
	}
	return ""
}

func (x *UserCoinReward) GetCoinUnit() string {
	if x != nil {
		return x.CoinUnit
	}
	return ""
}

func (x *UserCoinReward) GetCreatedAt() uint32 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *UserCoinReward) GetUpdatedAt() uint32 {
	if x != nil {
		return x.UpdatedAt
	}
	return 0
}

type GetMyCoinRewardsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AppID  string `protobuf:"bytes,10,opt,name=AppID,proto3" json:"AppID,omitempty"`
	UserID string `protobuf:"bytes,20,opt,name=UserID,proto3" json:"UserID,omitempty"`
	Offset int32  `protobuf:"varint,30,opt,name=Offset,proto3" json:"Offset,omitempty"`
	Limit  int32  `protobuf:"varint,40,opt,name=Limit,proto3" json:"Limit,omitempty"`
}

func (x *GetMyCoinRewardsRequest) Reset() {
	*x = GetMyCoinRewardsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_inspire_gw_v1_user_coin_reward_reward_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMyCoinRewardsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMyCoinRewardsRequest) ProtoMessage() {}

func (x *GetMyCoinRewardsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_inspire_gw_v1_user_coin_reward_reward_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMyCoinRewardsRequest.ProtoReflect.Descriptor instead.
func (*GetMyCoinRewardsRequest) Descriptor() ([]byte, []int) {
	return file_npool_inspire_gw_v1_user_coin_reward_reward_proto_rawDescGZIP(), []int{1}
}

func (x *GetMyCoinRewardsRequest) GetAppID() string {
	if x != nil {
		return x.AppID
	}
	return ""
}

func (x *GetMyCoinRewardsRequest) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

func (x *GetMyCoinRewardsRequest) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *GetMyCoinRewardsRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type GetMyCoinRewardsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Infos []*UserCoinReward `protobuf:"bytes,10,rep,name=Infos,proto3" json:"Infos,omitempty"`
	Total uint32            `protobuf:"varint,20,opt,name=Total,proto3" json:"Total,omitempty"`
}

func (x *GetMyCoinRewardsResponse) Reset() {
	*x = GetMyCoinRewardsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_inspire_gw_v1_user_coin_reward_reward_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMyCoinRewardsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMyCoinRewardsResponse) ProtoMessage() {}

func (x *GetMyCoinRewardsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_inspire_gw_v1_user_coin_reward_reward_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMyCoinRewardsResponse.ProtoReflect.Descriptor instead.
func (*GetMyCoinRewardsResponse) Descriptor() ([]byte, []int) {
	return file_npool_inspire_gw_v1_user_coin_reward_reward_proto_rawDescGZIP(), []int{2}
}

func (x *GetMyCoinRewardsResponse) GetInfos() []*UserCoinReward {
	if x != nil {
		return x.Infos
	}
	return nil
}

func (x *GetMyCoinRewardsResponse) GetTotal() uint32 {
	if x != nil {
		return x.Total
	}
	return 0
}

type AdminGetUserCoinRewardsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TargetAppID string `protobuf:"bytes,10,opt,name=TargetAppID,proto3" json:"TargetAppID,omitempty"`
	Offset      int32  `protobuf:"varint,30,opt,name=Offset,proto3" json:"Offset,omitempty"`
	Limit       int32  `protobuf:"varint,40,opt,name=Limit,proto3" json:"Limit,omitempty"`
}

func (x *AdminGetUserCoinRewardsRequest) Reset() {
	*x = AdminGetUserCoinRewardsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_inspire_gw_v1_user_coin_reward_reward_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdminGetUserCoinRewardsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdminGetUserCoinRewardsRequest) ProtoMessage() {}

func (x *AdminGetUserCoinRewardsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_inspire_gw_v1_user_coin_reward_reward_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdminGetUserCoinRewardsRequest.ProtoReflect.Descriptor instead.
func (*AdminGetUserCoinRewardsRequest) Descriptor() ([]byte, []int) {
	return file_npool_inspire_gw_v1_user_coin_reward_reward_proto_rawDescGZIP(), []int{3}
}

func (x *AdminGetUserCoinRewardsRequest) GetTargetAppID() string {
	if x != nil {
		return x.TargetAppID
	}
	return ""
}

func (x *AdminGetUserCoinRewardsRequest) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *AdminGetUserCoinRewardsRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type AdminGetUserCoinRewardsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Infos []*UserCoinReward `protobuf:"bytes,10,rep,name=Infos,proto3" json:"Infos,omitempty"`
	Total uint32            `protobuf:"varint,20,opt,name=Total,proto3" json:"Total,omitempty"`
}

func (x *AdminGetUserCoinRewardsResponse) Reset() {
	*x = AdminGetUserCoinRewardsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_inspire_gw_v1_user_coin_reward_reward_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdminGetUserCoinRewardsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdminGetUserCoinRewardsResponse) ProtoMessage() {}

func (x *AdminGetUserCoinRewardsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_inspire_gw_v1_user_coin_reward_reward_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdminGetUserCoinRewardsResponse.ProtoReflect.Descriptor instead.
func (*AdminGetUserCoinRewardsResponse) Descriptor() ([]byte, []int) {
	return file_npool_inspire_gw_v1_user_coin_reward_reward_proto_rawDescGZIP(), []int{4}
}

func (x *AdminGetUserCoinRewardsResponse) GetInfos() []*UserCoinReward {
	if x != nil {
		return x.Infos
	}
	return nil
}

func (x *AdminGetUserCoinRewardsResponse) GetTotal() uint32 {
	if x != nil {
		return x.Total
	}
	return 0
}

var File_npool_inspire_gw_v1_user_coin_reward_reward_proto protoreflect.FileDescriptor

var file_npool_inspire_gw_v1_user_coin_reward_reward_proto_rawDesc = []byte{
	0x0a, 0x31, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x69, 0x6e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2f,
	0x67, 0x77, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x63, 0x6f, 0x69, 0x6e, 0x2f,
	0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x2f, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x23, 0x69, 0x6e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x67, 0x61, 0x74,
	0x65, 0x77, 0x61, 0x79, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x63, 0x6f, 0x69, 0x6e, 0x2e, 0x72,
	0x65, 0x77, 0x61, 0x72, 0x64, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xdc, 0x02, 0x0a, 0x0e, 0x55, 0x73, 0x65, 0x72, 0x43,
	0x6f, 0x69, 0x6e, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x45, 0x6e, 0x74,
	0x49, 0x44, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x45, 0x6e, 0x74, 0x49, 0x44, 0x12,
	0x14, 0x0a, 0x05, 0x41, 0x70, 0x70, 0x49, 0x44, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x41, 0x70, 0x70, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18,
	0x28, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x1e, 0x0a,
	0x0a, 0x43, 0x6f, 0x69, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x49, 0x44, 0x18, 0x32, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x43, 0x6f, 0x69, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x49, 0x44, 0x12, 0x20, 0x0a,
	0x0b, 0x43, 0x6f, 0x69, 0x6e, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x73, 0x18, 0x3c, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x43, 0x6f, 0x69, 0x6e, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x73, 0x12,
	0x1a, 0x0a, 0x08, 0x43, 0x6f, 0x69, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x46, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x43, 0x6f, 0x69, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x44,
	0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x18, 0x50, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x0c, 0x44, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x12,
	0x1a, 0x0a, 0x08, 0x43, 0x6f, 0x69, 0x6e, 0x4c, 0x6f, 0x67, 0x6f, 0x18, 0x5a, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x43, 0x6f, 0x69, 0x6e, 0x4c, 0x6f, 0x67, 0x6f, 0x12, 0x1a, 0x0a, 0x08, 0x43,
	0x6f, 0x69, 0x6e, 0x55, 0x6e, 0x69, 0x74, 0x18, 0x64, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x43,
	0x6f, 0x69, 0x6e, 0x55, 0x6e, 0x69, 0x74, 0x12, 0x1d, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x18, 0xe8, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x18, 0xf2, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x75, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x4d, 0x79, 0x43, 0x6f,
	0x69, 0x6e, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x14, 0x0a, 0x05, 0x41, 0x70, 0x70, 0x49, 0x44, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x41, 0x70, 0x70, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44,
	0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x16,
	0x0a, 0x06, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06,
	0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x18,
	0x28, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x22, 0x7b, 0x0a, 0x18,
	0x47, 0x65, 0x74, 0x4d, 0x79, 0x43, 0x6f, 0x69, 0x6e, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x49, 0x0a, 0x05, 0x49, 0x6e, 0x66, 0x6f,
	0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x33, 0x2e, 0x69, 0x6e, 0x73, 0x70, 0x69, 0x72,
	0x65, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x63,
	0x6f, 0x69, 0x6e, 0x2e, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73,
	0x65, 0x72, 0x43, 0x6f, 0x69, 0x6e, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x52, 0x05, 0x49, 0x6e,
	0x66, 0x6f, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x14, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x22, 0x70, 0x0a, 0x1e, 0x41, 0x64, 0x6d,
	0x69, 0x6e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x43, 0x6f, 0x69, 0x6e, 0x52, 0x65, 0x77,
	0x61, 0x72, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x54,
	0x61, 0x72, 0x67, 0x65, 0x74, 0x41, 0x70, 0x70, 0x49, 0x44, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x41, 0x70, 0x70, 0x49, 0x44, 0x12, 0x16, 0x0a,
	0x06, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x4f,
	0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x28,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x22, 0x82, 0x01, 0x0a, 0x1f,
	0x41, 0x64, 0x6d, 0x69, 0x6e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x43, 0x6f, 0x69, 0x6e,
	0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x49, 0x0a, 0x05, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x33,
	0x2e, 0x69, 0x6e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x63, 0x6f, 0x69, 0x6e, 0x2e, 0x72, 0x65, 0x77, 0x61, 0x72,
	0x64, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x43, 0x6f, 0x69, 0x6e, 0x52, 0x65, 0x77,
	0x61, 0x72, 0x64, 0x52, 0x05, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x6f,
	0x74, 0x61, 0x6c, 0x18, 0x14, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c,
	0x32, 0x8f, 0x03, 0x0a, 0x07, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x12, 0xce, 0x01, 0x0a,
	0x17, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x43, 0x6f, 0x69,
	0x6e, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x73, 0x12, 0x43, 0x2e, 0x69, 0x6e, 0x73, 0x70, 0x69,
	0x72, 0x65, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e,
	0x63, 0x6f, 0x69, 0x6e, 0x2e, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x41,
	0x64, 0x6d, 0x69, 0x6e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x43, 0x6f, 0x69, 0x6e, 0x52,
	0x65, 0x77, 0x61, 0x72, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x44, 0x2e,
	0x69, 0x6e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x2e, 0x63, 0x6f, 0x69, 0x6e, 0x2e, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64,
	0x2e, 0x76, 0x31, 0x2e, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72,
	0x43, 0x6f, 0x69, 0x6e, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x28, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x22, 0x3a, 0x01, 0x2a, 0x22, 0x1d,
	0x2f, 0x76, 0x31, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x67, 0x65, 0x74, 0x2f, 0x75, 0x73,
	0x65, 0x72, 0x63, 0x6f, 0x69, 0x6e, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x73, 0x12, 0xb2, 0x01,
	0x0a, 0x10, 0x47, 0x65, 0x74, 0x4d, 0x79, 0x43, 0x6f, 0x69, 0x6e, 0x52, 0x65, 0x77, 0x61, 0x72,
	0x64, 0x73, 0x12, 0x3c, 0x2e, 0x69, 0x6e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x67, 0x61, 0x74,
	0x65, 0x77, 0x61, 0x79, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x63, 0x6f, 0x69, 0x6e, 0x2e, 0x72,
	0x65, 0x77, 0x61, 0x72, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x79, 0x43, 0x6f,
	0x69, 0x6e, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x3d, 0x2e, 0x69, 0x6e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77,
	0x61, 0x79, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x63, 0x6f, 0x69, 0x6e, 0x2e, 0x72, 0x65, 0x77,
	0x61, 0x72, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x79, 0x43, 0x6f, 0x69, 0x6e,
	0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x21, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1b, 0x3a, 0x01, 0x2a, 0x22, 0x16, 0x2f, 0x76, 0x31, 0x2f,
	0x67, 0x65, 0x74, 0x2f, 0x6d, 0x79, 0x2f, 0x63, 0x6f, 0x69, 0x6e, 0x72, 0x65, 0x77, 0x61, 0x72,
	0x64, 0x73, 0x42, 0x47, 0x5a, 0x45, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x4e, 0x70, 0x6f, 0x6f, 0x6c, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x69, 0x6e, 0x73,
	0x70, 0x69, 0x72, 0x65, 0x2f, 0x67, 0x77, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f,
	0x63, 0x6f, 0x69, 0x6e, 0x2f, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_npool_inspire_gw_v1_user_coin_reward_reward_proto_rawDescOnce sync.Once
	file_npool_inspire_gw_v1_user_coin_reward_reward_proto_rawDescData = file_npool_inspire_gw_v1_user_coin_reward_reward_proto_rawDesc
)

func file_npool_inspire_gw_v1_user_coin_reward_reward_proto_rawDescGZIP() []byte {
	file_npool_inspire_gw_v1_user_coin_reward_reward_proto_rawDescOnce.Do(func() {
		file_npool_inspire_gw_v1_user_coin_reward_reward_proto_rawDescData = protoimpl.X.CompressGZIP(file_npool_inspire_gw_v1_user_coin_reward_reward_proto_rawDescData)
	})
	return file_npool_inspire_gw_v1_user_coin_reward_reward_proto_rawDescData
}

var file_npool_inspire_gw_v1_user_coin_reward_reward_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_npool_inspire_gw_v1_user_coin_reward_reward_proto_goTypes = []interface{}{
	(*UserCoinReward)(nil),                  // 0: inspire.gateway.user.coin.reward.v1.UserCoinReward
	(*GetMyCoinRewardsRequest)(nil),         // 1: inspire.gateway.user.coin.reward.v1.GetMyCoinRewardsRequest
	(*GetMyCoinRewardsResponse)(nil),        // 2: inspire.gateway.user.coin.reward.v1.GetMyCoinRewardsResponse
	(*AdminGetUserCoinRewardsRequest)(nil),  // 3: inspire.gateway.user.coin.reward.v1.AdminGetUserCoinRewardsRequest
	(*AdminGetUserCoinRewardsResponse)(nil), // 4: inspire.gateway.user.coin.reward.v1.AdminGetUserCoinRewardsResponse
}
var file_npool_inspire_gw_v1_user_coin_reward_reward_proto_depIdxs = []int32{
	0, // 0: inspire.gateway.user.coin.reward.v1.GetMyCoinRewardsResponse.Infos:type_name -> inspire.gateway.user.coin.reward.v1.UserCoinReward
	0, // 1: inspire.gateway.user.coin.reward.v1.AdminGetUserCoinRewardsResponse.Infos:type_name -> inspire.gateway.user.coin.reward.v1.UserCoinReward
	3, // 2: inspire.gateway.user.coin.reward.v1.Gateway.AdminGetUserCoinRewards:input_type -> inspire.gateway.user.coin.reward.v1.AdminGetUserCoinRewardsRequest
	1, // 3: inspire.gateway.user.coin.reward.v1.Gateway.GetMyCoinRewards:input_type -> inspire.gateway.user.coin.reward.v1.GetMyCoinRewardsRequest
	4, // 4: inspire.gateway.user.coin.reward.v1.Gateway.AdminGetUserCoinRewards:output_type -> inspire.gateway.user.coin.reward.v1.AdminGetUserCoinRewardsResponse
	2, // 5: inspire.gateway.user.coin.reward.v1.Gateway.GetMyCoinRewards:output_type -> inspire.gateway.user.coin.reward.v1.GetMyCoinRewardsResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_npool_inspire_gw_v1_user_coin_reward_reward_proto_init() }
func file_npool_inspire_gw_v1_user_coin_reward_reward_proto_init() {
	if File_npool_inspire_gw_v1_user_coin_reward_reward_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_npool_inspire_gw_v1_user_coin_reward_reward_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserCoinReward); i {
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
		file_npool_inspire_gw_v1_user_coin_reward_reward_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMyCoinRewardsRequest); i {
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
		file_npool_inspire_gw_v1_user_coin_reward_reward_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMyCoinRewardsResponse); i {
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
		file_npool_inspire_gw_v1_user_coin_reward_reward_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdminGetUserCoinRewardsRequest); i {
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
		file_npool_inspire_gw_v1_user_coin_reward_reward_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdminGetUserCoinRewardsResponse); i {
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
			RawDescriptor: file_npool_inspire_gw_v1_user_coin_reward_reward_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_npool_inspire_gw_v1_user_coin_reward_reward_proto_goTypes,
		DependencyIndexes: file_npool_inspire_gw_v1_user_coin_reward_reward_proto_depIdxs,
		MessageInfos:      file_npool_inspire_gw_v1_user_coin_reward_reward_proto_msgTypes,
	}.Build()
	File_npool_inspire_gw_v1_user_coin_reward_reward_proto = out.File
	file_npool_inspire_gw_v1_user_coin_reward_reward_proto_rawDesc = nil
	file_npool_inspire_gw_v1_user_coin_reward_reward_proto_goTypes = nil
	file_npool_inspire_gw_v1_user_coin_reward_reward_proto_depIdxs = nil
}