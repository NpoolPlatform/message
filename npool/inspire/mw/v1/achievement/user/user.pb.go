// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.18.1
// source: npool/inspire/mw/v1/achievement/user/user.proto

package user

import (
	v1 "github.com/NpoolPlatform/message/npool/basetypes/v1"
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

type AchievementUserReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID    *uint32 `protobuf:"varint,10,opt,name=ID,proto3,oneof" json:"ID,omitempty"`
	EntID *string `protobuf:"bytes,11,opt,name=EntID,proto3,oneof" json:"EntID,omitempty"`
}

func (x *AchievementUserReq) Reset() {
	*x = AchievementUserReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_inspire_mw_v1_achievement_user_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AchievementUserReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AchievementUserReq) ProtoMessage() {}

func (x *AchievementUserReq) ProtoReflect() protoreflect.Message {
	mi := &file_npool_inspire_mw_v1_achievement_user_user_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AchievementUserReq.ProtoReflect.Descriptor instead.
func (*AchievementUserReq) Descriptor() ([]byte, []int) {
	return file_npool_inspire_mw_v1_achievement_user_user_proto_rawDescGZIP(), []int{0}
}

func (x *AchievementUserReq) GetID() uint32 {
	if x != nil && x.ID != nil {
		return *x.ID
	}
	return 0
}

func (x *AchievementUserReq) GetEntID() string {
	if x != nil && x.EntID != nil {
		return *x.EntID
	}
	return ""
}

type AchievementUser struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: sql:"id"
	ID uint32 `protobuf:"varint,10,opt,name=ID,proto3" json:"ID,omitempty" sql:"id"`
	// @inject_tag: sql:"ent_id"
	EntID string `protobuf:"bytes,20,opt,name=EntID,proto3" json:"EntID,omitempty" sql:"ent_id"`
	// @inject_tag: sql:"app_id"
	AppID string `protobuf:"bytes,30,opt,name=AppID,proto3" json:"AppID,omitempty" sql:"app_id"`
	// @inject_tag: sql:"user_id"
	UserID string `protobuf:"bytes,40,opt,name=UserID,proto3" json:"UserID,omitempty" sql:"user_id"`
	// Commission amount in USD
	// @inject_tag: sql:"total_commission"
	TotalCommission string `protobuf:"bytes,50,opt,name=TotalCommission,proto3" json:"TotalCommission,omitempty" sql:"total_commission"`
	// @inject_tag: sql:"self_commission"
	SelfCommission string `protobuf:"bytes,60,opt,name=SelfCommission,proto3" json:"SelfCommission,omitempty" sql:"self_commission"`
	// @inject_tag: sql:"direct_invites"
	DirectInvites uint32 `protobuf:"varint,70,opt,name=DirectInvites,proto3" json:"DirectInvites,omitempty" sql:"direct_invites"`
	// @inject_tag: sql:"indirect_invites"
	IndirectInvites uint32 `protobuf:"varint,80,opt,name=IndirectInvites,proto3" json:"IndirectInvites,omitempty" sql:"indirect_invites"`
	// @inject_tag: sql:"direct_consume_amount"
	DirectConsumeAmount string `protobuf:"bytes,90,opt,name=DirectConsumeAmount,proto3" json:"DirectConsumeAmount,omitempty" sql:"direct_consume_amount"`
	// @inject_tag: sql:"invitee_consume_amount"
	InviteeConsumeAmount string `protobuf:"bytes,100,opt,name=InviteeConsumeAmount,proto3" json:"InviteeConsumeAmount,omitempty" sql:"invitee_consume_amount"`
	// @inject_tag: sql:"created_at"
	CreatedAt uint32 `protobuf:"varint,1000,opt,name=CreatedAt,proto3" json:"CreatedAt,omitempty" sql:"created_at"`
	// @inject_tag: sql:"updated_at"
	UpdatedAt uint32 `protobuf:"varint,1010,opt,name=UpdatedAt,proto3" json:"UpdatedAt,omitempty" sql:"updated_at"`
	// @inject_tag: sql:"deleted_at"
	DeletedAt uint32 `protobuf:"varint,1020,opt,name=DeletedAt,proto3" json:"DeletedAt,omitempty" sql:"deleted_at"`
}

func (x *AchievementUser) Reset() {
	*x = AchievementUser{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_inspire_mw_v1_achievement_user_user_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AchievementUser) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AchievementUser) ProtoMessage() {}

func (x *AchievementUser) ProtoReflect() protoreflect.Message {
	mi := &file_npool_inspire_mw_v1_achievement_user_user_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AchievementUser.ProtoReflect.Descriptor instead.
func (*AchievementUser) Descriptor() ([]byte, []int) {
	return file_npool_inspire_mw_v1_achievement_user_user_proto_rawDescGZIP(), []int{1}
}

func (x *AchievementUser) GetID() uint32 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *AchievementUser) GetEntID() string {
	if x != nil {
		return x.EntID
	}
	return ""
}

func (x *AchievementUser) GetAppID() string {
	if x != nil {
		return x.AppID
	}
	return ""
}

func (x *AchievementUser) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

func (x *AchievementUser) GetTotalCommission() string {
	if x != nil {
		return x.TotalCommission
	}
	return ""
}

func (x *AchievementUser) GetSelfCommission() string {
	if x != nil {
		return x.SelfCommission
	}
	return ""
}

func (x *AchievementUser) GetDirectInvites() uint32 {
	if x != nil {
		return x.DirectInvites
	}
	return 0
}

func (x *AchievementUser) GetIndirectInvites() uint32 {
	if x != nil {
		return x.IndirectInvites
	}
	return 0
}

func (x *AchievementUser) GetDirectConsumeAmount() string {
	if x != nil {
		return x.DirectConsumeAmount
	}
	return ""
}

func (x *AchievementUser) GetInviteeConsumeAmount() string {
	if x != nil {
		return x.InviteeConsumeAmount
	}
	return ""
}

func (x *AchievementUser) GetCreatedAt() uint32 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *AchievementUser) GetUpdatedAt() uint32 {
	if x != nil {
		return x.UpdatedAt
	}
	return 0
}

func (x *AchievementUser) GetDeletedAt() uint32 {
	if x != nil {
		return x.DeletedAt
	}
	return 0
}

type Conds struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EntID   *v1.StringVal      `protobuf:"bytes,10,opt,name=EntID,proto3,oneof" json:"EntID,omitempty"`
	AppID   *v1.StringVal      `protobuf:"bytes,20,opt,name=AppID,proto3,oneof" json:"AppID,omitempty"`
	UserID  *v1.StringVal      `protobuf:"bytes,30,opt,name=UserID,proto3,oneof" json:"UserID,omitempty"`
	UserIDs *v1.StringSliceVal `protobuf:"bytes,40,opt,name=UserIDs,proto3,oneof" json:"UserIDs,omitempty"`
}

func (x *Conds) Reset() {
	*x = Conds{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_inspire_mw_v1_achievement_user_user_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Conds) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Conds) ProtoMessage() {}

func (x *Conds) ProtoReflect() protoreflect.Message {
	mi := &file_npool_inspire_mw_v1_achievement_user_user_proto_msgTypes[2]
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
	return file_npool_inspire_mw_v1_achievement_user_user_proto_rawDescGZIP(), []int{2}
}

func (x *Conds) GetEntID() *v1.StringVal {
	if x != nil {
		return x.EntID
	}
	return nil
}

func (x *Conds) GetAppID() *v1.StringVal {
	if x != nil {
		return x.AppID
	}
	return nil
}

func (x *Conds) GetUserID() *v1.StringVal {
	if x != nil {
		return x.UserID
	}
	return nil
}

func (x *Conds) GetUserIDs() *v1.StringSliceVal {
	if x != nil {
		return x.UserIDs
	}
	return nil
}

type GetAchievementUsersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Conds  *Conds `protobuf:"bytes,10,opt,name=Conds,proto3" json:"Conds,omitempty"`
	Offset int32  `protobuf:"varint,20,opt,name=Offset,proto3" json:"Offset,omitempty"`
	Limit  int32  `protobuf:"varint,30,opt,name=Limit,proto3" json:"Limit,omitempty"`
}

func (x *GetAchievementUsersRequest) Reset() {
	*x = GetAchievementUsersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_inspire_mw_v1_achievement_user_user_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAchievementUsersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAchievementUsersRequest) ProtoMessage() {}

func (x *GetAchievementUsersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_inspire_mw_v1_achievement_user_user_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAchievementUsersRequest.ProtoReflect.Descriptor instead.
func (*GetAchievementUsersRequest) Descriptor() ([]byte, []int) {
	return file_npool_inspire_mw_v1_achievement_user_user_proto_rawDescGZIP(), []int{3}
}

func (x *GetAchievementUsersRequest) GetConds() *Conds {
	if x != nil {
		return x.Conds
	}
	return nil
}

func (x *GetAchievementUsersRequest) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *GetAchievementUsersRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type GetAchievementUsersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Infos []*AchievementUser `protobuf:"bytes,10,rep,name=Infos,proto3" json:"Infos,omitempty"`
	Total uint32             `protobuf:"varint,20,opt,name=Total,proto3" json:"Total,omitempty"`
}

func (x *GetAchievementUsersResponse) Reset() {
	*x = GetAchievementUsersResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_inspire_mw_v1_achievement_user_user_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAchievementUsersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAchievementUsersResponse) ProtoMessage() {}

func (x *GetAchievementUsersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_inspire_mw_v1_achievement_user_user_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAchievementUsersResponse.ProtoReflect.Descriptor instead.
func (*GetAchievementUsersResponse) Descriptor() ([]byte, []int) {
	return file_npool_inspire_mw_v1_achievement_user_user_proto_rawDescGZIP(), []int{4}
}

func (x *GetAchievementUsersResponse) GetInfos() []*AchievementUser {
	if x != nil {
		return x.Infos
	}
	return nil
}

func (x *GetAchievementUsersResponse) GetTotal() uint32 {
	if x != nil {
		return x.Total
	}
	return 0
}

type DeleteAchievementUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *AchievementUserReq `protobuf:"bytes,10,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (x *DeleteAchievementUserRequest) Reset() {
	*x = DeleteAchievementUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_inspire_mw_v1_achievement_user_user_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteAchievementUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteAchievementUserRequest) ProtoMessage() {}

func (x *DeleteAchievementUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_inspire_mw_v1_achievement_user_user_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteAchievementUserRequest.ProtoReflect.Descriptor instead.
func (*DeleteAchievementUserRequest) Descriptor() ([]byte, []int) {
	return file_npool_inspire_mw_v1_achievement_user_user_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteAchievementUserRequest) GetInfo() *AchievementUserReq {
	if x != nil {
		return x.Info
	}
	return nil
}

type DeleteAchievementUserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *AchievementUser `protobuf:"bytes,10,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (x *DeleteAchievementUserResponse) Reset() {
	*x = DeleteAchievementUserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_inspire_mw_v1_achievement_user_user_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteAchievementUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteAchievementUserResponse) ProtoMessage() {}

func (x *DeleteAchievementUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_inspire_mw_v1_achievement_user_user_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteAchievementUserResponse.ProtoReflect.Descriptor instead.
func (*DeleteAchievementUserResponse) Descriptor() ([]byte, []int) {
	return file_npool_inspire_mw_v1_achievement_user_user_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteAchievementUserResponse) GetInfo() *AchievementUser {
	if x != nil {
		return x.Info
	}
	return nil
}

var File_npool_inspire_mw_v1_achievement_user_user_proto protoreflect.FileDescriptor

var file_npool_inspire_mw_v1_achievement_user_user_proto_rawDesc = []byte{
	0x0a, 0x2f, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x69, 0x6e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2f,
	0x6d, 0x77, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x63, 0x68, 0x69, 0x65, 0x76, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x26, 0x69, 0x6e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c,
	0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x61, 0x63, 0x68, 0x69, 0x65, 0x76, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x62,
	0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6e, 0x64,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x55, 0x0a, 0x12, 0x41, 0x63, 0x68, 0x69, 0x65,
	0x76, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x12, 0x13, 0x0a,
	0x02, 0x49, 0x44, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0d, 0x48, 0x00, 0x52, 0x02, 0x49, 0x44, 0x88,
	0x01, 0x01, 0x12, 0x19, 0x0a, 0x05, 0x45, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x0b, 0x20, 0x01, 0x28,
	0x09, 0x48, 0x01, 0x52, 0x05, 0x45, 0x6e, 0x74, 0x49, 0x44, 0x88, 0x01, 0x01, 0x42, 0x05, 0x0a,
	0x03, 0x5f, 0x49, 0x44, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x45, 0x6e, 0x74, 0x49, 0x44, 0x22, 0xca,
	0x03, 0x0a, 0x0f, 0x41, 0x63, 0x68, 0x69, 0x65, 0x76, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x55, 0x73,
	0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02,
	0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x45, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x14, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x45, 0x6e, 0x74, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x41, 0x70, 0x70, 0x49,
	0x44, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x41, 0x70, 0x70, 0x49, 0x44, 0x12, 0x16,
	0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x28, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x28, 0x0a, 0x0f, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x43,
	0x6f, 0x6d, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x32, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0f, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x12, 0x26, 0x0a, 0x0e, 0x53, 0x65, 0x6c, 0x66, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x18, 0x3c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x53, 0x65, 0x6c, 0x66, 0x43, 0x6f,
	0x6d, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x24, 0x0a, 0x0d, 0x44, 0x69, 0x72, 0x65,
	0x63, 0x74, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x73, 0x18, 0x46, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x0d, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x73, 0x12, 0x28,
	0x0a, 0x0f, 0x49, 0x6e, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x65,
	0x73, 0x18, 0x50, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0f, 0x49, 0x6e, 0x64, 0x69, 0x72, 0x65, 0x63,
	0x74, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x73, 0x12, 0x30, 0x0a, 0x13, 0x44, 0x69, 0x72, 0x65,
	0x63, 0x74, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x5a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x13, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x43, 0x6f, 0x6e,
	0x73, 0x75, 0x6d, 0x65, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x32, 0x0a, 0x14, 0x49, 0x6e,
	0x76, 0x69, 0x74, 0x65, 0x65, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x41, 0x6d, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x64, 0x20, 0x01, 0x28, 0x09, 0x52, 0x14, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x65,
	0x65, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1d,
	0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0xe8, 0x07, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a,
	0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0xf2, 0x07, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x09,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0xfc, 0x07, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x09, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x8d, 0x02, 0x0a, 0x05,
	0x43, 0x6f, 0x6e, 0x64, 0x73, 0x12, 0x32, 0x0a, 0x05, 0x45, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x48, 0x00, 0x52,
	0x05, 0x45, 0x6e, 0x74, 0x49, 0x44, 0x88, 0x01, 0x01, 0x12, 0x32, 0x0a, 0x05, 0x41, 0x70, 0x70,
	0x49, 0x44, 0x18, 0x14, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61,
	0x6c, 0x48, 0x01, 0x52, 0x05, 0x41, 0x70, 0x70, 0x49, 0x44, 0x88, 0x01, 0x01, 0x12, 0x34, 0x0a,
	0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e,
	0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x72,
	0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x48, 0x02, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44,
	0x88, 0x01, 0x01, 0x12, 0x3b, 0x0a, 0x07, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x73, 0x18, 0x28,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x53, 0x6c, 0x69, 0x63, 0x65, 0x56,
	0x61, 0x6c, 0x48, 0x03, 0x52, 0x07, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x73, 0x88, 0x01, 0x01,
	0x42, 0x08, 0x0a, 0x06, 0x5f, 0x45, 0x6e, 0x74, 0x49, 0x44, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x41,
	0x70, 0x70, 0x49, 0x44, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x42,
	0x0a, 0x0a, 0x08, 0x5f, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x73, 0x22, 0x8f, 0x01, 0x0a, 0x1a,
	0x47, 0x65, 0x74, 0x41, 0x63, 0x68, 0x69, 0x65, 0x76, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x55, 0x73,
	0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x43, 0x0a, 0x05, 0x43, 0x6f,
	0x6e, 0x64, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x69, 0x6e, 0x73, 0x70,
	0x69, 0x72, 0x65, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x61,
	0x63, 0x68, 0x69, 0x65, 0x76, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e,
	0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x64, 0x73, 0x52, 0x05, 0x43, 0x6f, 0x6e, 0x64, 0x73, 0x12,
	0x16, 0x0a, 0x06, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x14, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x06, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x4c, 0x69, 0x6d, 0x69, 0x74,
	0x18, 0x1e, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x22, 0x82, 0x01,
	0x0a, 0x1b, 0x47, 0x65, 0x74, 0x41, 0x63, 0x68, 0x69, 0x65, 0x76, 0x65, 0x6d, 0x65, 0x6e, 0x74,
	0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4d, 0x0a,
	0x05, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x37, 0x2e, 0x69,
	0x6e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72,
	0x65, 0x2e, 0x61, 0x63, 0x68, 0x69, 0x65, 0x76, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x63, 0x68, 0x69, 0x65, 0x76, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x05, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x12, 0x14, 0x0a, 0x05,
	0x54, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x14, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x54, 0x6f, 0x74,
	0x61, 0x6c, 0x22, 0x6e, 0x0a, 0x1c, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x63, 0x68, 0x69,
	0x65, 0x76, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x4e, 0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x3a, 0x2e, 0x69, 0x6e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c,
	0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x61, 0x63, 0x68, 0x69, 0x65, 0x76, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x63, 0x68, 0x69, 0x65, 0x76,
	0x65, 0x6d, 0x65, 0x6e, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x52, 0x04, 0x49, 0x6e,
	0x66, 0x6f, 0x22, 0x6c, 0x0a, 0x1d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x63, 0x68, 0x69,
	0x65, 0x76, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x4b, 0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x37, 0x2e, 0x69, 0x6e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x6d, 0x69, 0x64, 0x64,
	0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x61, 0x63, 0x68, 0x69, 0x65, 0x76, 0x65, 0x6d, 0x65,
	0x6e, 0x74, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x63, 0x68, 0x69, 0x65,
	0x76, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x49, 0x6e, 0x66, 0x6f,
	0x32, 0xb6, 0x03, 0x0a, 0x0a, 0x4d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x12,
	0xce, 0x01, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x41, 0x63, 0x68, 0x69, 0x65, 0x76, 0x65, 0x6d, 0x65,
	0x6e, 0x74, 0x55, 0x73, 0x65, 0x72, 0x73, 0x12, 0x42, 0x2e, 0x69, 0x6e, 0x73, 0x70, 0x69, 0x72,
	0x65, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x61, 0x63, 0x68,
	0x69, 0x65, 0x76, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31,
	0x2e, 0x47, 0x65, 0x74, 0x41, 0x63, 0x68, 0x69, 0x65, 0x76, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x55,
	0x73, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x43, 0x2e, 0x69, 0x6e,
	0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65,
	0x2e, 0x61, 0x63, 0x68, 0x69, 0x65, 0x76, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x63, 0x68, 0x69, 0x65, 0x76, 0x65, 0x6d,
	0x65, 0x6e, 0x74, 0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x2e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x28, 0x22, 0x23, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x63,
	0x68, 0x69, 0x65, 0x76, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x47, 0x65, 0x74, 0x41, 0x63, 0x68,
	0x69, 0x65, 0x76, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x55, 0x73, 0x65, 0x72, 0x73, 0x3a, 0x01, 0x2a,
	0x12, 0xd6, 0x01, 0x0a, 0x15, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x63, 0x68, 0x69, 0x65,
	0x76, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x55, 0x73, 0x65, 0x72, 0x12, 0x44, 0x2e, 0x69, 0x6e, 0x73,
	0x70, 0x69, 0x72, 0x65, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e,
	0x61, 0x63, 0x68, 0x69, 0x65, 0x76, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x63, 0x68, 0x69, 0x65, 0x76,
	0x65, 0x6d, 0x65, 0x6e, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x45, 0x2e, 0x69, 0x6e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c,
	0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x61, 0x63, 0x68, 0x69, 0x65, 0x76, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x41, 0x63, 0x68, 0x69, 0x65, 0x76, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x30, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2a, 0x22,
	0x25, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x63, 0x68, 0x69, 0x65, 0x76, 0x65, 0x6d, 0x65, 0x6e, 0x74,
	0x5f, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x63, 0x68, 0x69, 0x65, 0x76, 0x65, 0x6d, 0x65,
	0x6e, 0x74, 0x55, 0x73, 0x65, 0x72, 0x3a, 0x01, 0x2a, 0x42, 0x47, 0x5a, 0x45, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4e, 0x70, 0x6f, 0x6f, 0x6c, 0x50, 0x6c, 0x61,
	0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x6e, 0x70,
	0x6f, 0x6f, 0x6c, 0x2f, 0x69, 0x6e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2f, 0x6d, 0x77, 0x2f, 0x76,
	0x31, 0x2f, 0x61, 0x63, 0x68, 0x69, 0x65, 0x76, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x75, 0x73,
	0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_npool_inspire_mw_v1_achievement_user_user_proto_rawDescOnce sync.Once
	file_npool_inspire_mw_v1_achievement_user_user_proto_rawDescData = file_npool_inspire_mw_v1_achievement_user_user_proto_rawDesc
)

func file_npool_inspire_mw_v1_achievement_user_user_proto_rawDescGZIP() []byte {
	file_npool_inspire_mw_v1_achievement_user_user_proto_rawDescOnce.Do(func() {
		file_npool_inspire_mw_v1_achievement_user_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_npool_inspire_mw_v1_achievement_user_user_proto_rawDescData)
	})
	return file_npool_inspire_mw_v1_achievement_user_user_proto_rawDescData
}

var file_npool_inspire_mw_v1_achievement_user_user_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_npool_inspire_mw_v1_achievement_user_user_proto_goTypes = []interface{}{
	(*AchievementUserReq)(nil),            // 0: inspire.middleware.achievement.user.v1.AchievementUserReq
	(*AchievementUser)(nil),               // 1: inspire.middleware.achievement.user.v1.AchievementUser
	(*Conds)(nil),                         // 2: inspire.middleware.achievement.user.v1.Conds
	(*GetAchievementUsersRequest)(nil),    // 3: inspire.middleware.achievement.user.v1.GetAchievementUsersRequest
	(*GetAchievementUsersResponse)(nil),   // 4: inspire.middleware.achievement.user.v1.GetAchievementUsersResponse
	(*DeleteAchievementUserRequest)(nil),  // 5: inspire.middleware.achievement.user.v1.DeleteAchievementUserRequest
	(*DeleteAchievementUserResponse)(nil), // 6: inspire.middleware.achievement.user.v1.DeleteAchievementUserResponse
	(*v1.StringVal)(nil),                  // 7: basetypes.v1.StringVal
	(*v1.StringSliceVal)(nil),             // 8: basetypes.v1.StringSliceVal
}
var file_npool_inspire_mw_v1_achievement_user_user_proto_depIdxs = []int32{
	7,  // 0: inspire.middleware.achievement.user.v1.Conds.EntID:type_name -> basetypes.v1.StringVal
	7,  // 1: inspire.middleware.achievement.user.v1.Conds.AppID:type_name -> basetypes.v1.StringVal
	7,  // 2: inspire.middleware.achievement.user.v1.Conds.UserID:type_name -> basetypes.v1.StringVal
	8,  // 3: inspire.middleware.achievement.user.v1.Conds.UserIDs:type_name -> basetypes.v1.StringSliceVal
	2,  // 4: inspire.middleware.achievement.user.v1.GetAchievementUsersRequest.Conds:type_name -> inspire.middleware.achievement.user.v1.Conds
	1,  // 5: inspire.middleware.achievement.user.v1.GetAchievementUsersResponse.Infos:type_name -> inspire.middleware.achievement.user.v1.AchievementUser
	0,  // 6: inspire.middleware.achievement.user.v1.DeleteAchievementUserRequest.Info:type_name -> inspire.middleware.achievement.user.v1.AchievementUserReq
	1,  // 7: inspire.middleware.achievement.user.v1.DeleteAchievementUserResponse.Info:type_name -> inspire.middleware.achievement.user.v1.AchievementUser
	3,  // 8: inspire.middleware.achievement.user.v1.Middleware.GetAchievementUsers:input_type -> inspire.middleware.achievement.user.v1.GetAchievementUsersRequest
	5,  // 9: inspire.middleware.achievement.user.v1.Middleware.DeleteAchievementUser:input_type -> inspire.middleware.achievement.user.v1.DeleteAchievementUserRequest
	4,  // 10: inspire.middleware.achievement.user.v1.Middleware.GetAchievementUsers:output_type -> inspire.middleware.achievement.user.v1.GetAchievementUsersResponse
	6,  // 11: inspire.middleware.achievement.user.v1.Middleware.DeleteAchievementUser:output_type -> inspire.middleware.achievement.user.v1.DeleteAchievementUserResponse
	10, // [10:12] is the sub-list for method output_type
	8,  // [8:10] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_npool_inspire_mw_v1_achievement_user_user_proto_init() }
func file_npool_inspire_mw_v1_achievement_user_user_proto_init() {
	if File_npool_inspire_mw_v1_achievement_user_user_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_npool_inspire_mw_v1_achievement_user_user_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AchievementUserReq); i {
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
		file_npool_inspire_mw_v1_achievement_user_user_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AchievementUser); i {
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
		file_npool_inspire_mw_v1_achievement_user_user_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
		file_npool_inspire_mw_v1_achievement_user_user_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAchievementUsersRequest); i {
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
		file_npool_inspire_mw_v1_achievement_user_user_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAchievementUsersResponse); i {
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
		file_npool_inspire_mw_v1_achievement_user_user_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteAchievementUserRequest); i {
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
		file_npool_inspire_mw_v1_achievement_user_user_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteAchievementUserResponse); i {
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
	file_npool_inspire_mw_v1_achievement_user_user_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_npool_inspire_mw_v1_achievement_user_user_proto_msgTypes[2].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_npool_inspire_mw_v1_achievement_user_user_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_npool_inspire_mw_v1_achievement_user_user_proto_goTypes,
		DependencyIndexes: file_npool_inspire_mw_v1_achievement_user_user_proto_depIdxs,
		MessageInfos:      file_npool_inspire_mw_v1_achievement_user_user_proto_msgTypes,
	}.Build()
	File_npool_inspire_mw_v1_achievement_user_user_proto = out.File
	file_npool_inspire_mw_v1_achievement_user_user_proto_rawDesc = nil
	file_npool_inspire_mw_v1_achievement_user_user_proto_goTypes = nil
	file_npool_inspire_mw_v1_achievement_user_user_proto_depIdxs = nil
}
