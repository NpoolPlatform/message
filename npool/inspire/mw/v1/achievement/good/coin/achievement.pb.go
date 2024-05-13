// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v3.18.1
// source: npool/inspire/mw/v1/achievement/good/coin/achievement.proto

package coin

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

type AchievementReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID             *uint32 `protobuf:"varint,10,opt,name=ID,proto3,oneof" json:"ID,omitempty"`
	EntID          *string `protobuf:"bytes,20,opt,name=EntID,proto3,oneof" json:"EntID,omitempty"`
	GoodCoinTypeID *string `protobuf:"bytes,50,opt,name=GoodCoinTypeID,proto3,oneof" json:"GoodCoinTypeID,omitempty"`
}

func (x *AchievementReq) Reset() {
	*x = AchievementReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_inspire_mw_v1_achievement_good_coin_achievement_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AchievementReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AchievementReq) ProtoMessage() {}

func (x *AchievementReq) ProtoReflect() protoreflect.Message {
	mi := &file_npool_inspire_mw_v1_achievement_good_coin_achievement_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AchievementReq.ProtoReflect.Descriptor instead.
func (*AchievementReq) Descriptor() ([]byte, []int) {
	return file_npool_inspire_mw_v1_achievement_good_coin_achievement_proto_rawDescGZIP(), []int{0}
}

func (x *AchievementReq) GetID() uint32 {
	if x != nil && x.ID != nil {
		return *x.ID
	}
	return 0
}

func (x *AchievementReq) GetEntID() string {
	if x != nil && x.EntID != nil {
		return *x.EntID
	}
	return ""
}

func (x *AchievementReq) GetGoodCoinTypeID() string {
	if x != nil && x.GoodCoinTypeID != nil {
		return *x.GoodCoinTypeID
	}
	return ""
}

type Achievement struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: sql:"id"
	ID uint32 `protobuf:"varint,10,opt,name=ID,proto3" json:"ID,omitempty" sql:"id"`
	// @inject_tag: sql:"ent_id"
	EntID string `protobuf:"bytes,11,opt,name=EntID,proto3" json:"EntID,omitempty" sql:"ent_id"`
	// @inject_tag: sql:"app_id"
	AppID string `protobuf:"bytes,20,opt,name=AppID,proto3" json:"AppID,omitempty" sql:"app_id"`
	// @inject_tag: sql:"user_id"
	UserID string `protobuf:"bytes,30,opt,name=UserID,proto3" json:"UserID,omitempty" sql:"user_id"`
	// @inject_tag: sql:"good_coin_type_id"
	GoodCoinTypeID string `protobuf:"bytes,40,opt,name=GoodCoinTypeID,proto3" json:"GoodCoinTypeID,omitempty" sql:"good_coin_type_id"`
	// Payment amount in USD
	// @inject_tag: sql:"total_amount_usd"
	TotalAmountUSD string `protobuf:"bytes,60,opt,name=TotalAmountUSD,proto3" json:"TotalAmountUSD,omitempty" sql:"total_amount_usd"`
	// @inject_tag: sql:"self_amount_usd"
	SelfAmountUSD string `protobuf:"bytes,70,opt,name=SelfAmountUSD,proto3" json:"SelfAmountUSD,omitempty" sql:"self_amount_usd"`
	// @inject_tag: sql:"total_units_v1"
	TotalUnits string `protobuf:"bytes,80,opt,name=TotalUnits,proto3" json:"TotalUnits,omitempty" sql:"total_units_v1"`
	// @inject_tag: sql:"self_units_v1"
	SelfUnits string `protobuf:"bytes,90,opt,name=SelfUnits,proto3" json:"SelfUnits,omitempty" sql:"self_units_v1"`
	// Commission amount in USD
	// @inject_tag: sql:"total_commission_usd"
	TotalCommissionUSD string `protobuf:"bytes,100,opt,name=TotalCommissionUSD,proto3" json:"TotalCommissionUSD,omitempty" sql:"total_commission_usd"`
	// @inject_tag: sql:"self_commission_usd"
	SelfCommissionUSD string `protobuf:"bytes,110,opt,name=SelfCommissionUSD,proto3" json:"SelfCommissionUSD,omitempty" sql:"self_commission_usd"`
	// @inject_tag: sql:"created_at"
	CreatedAt uint32 `protobuf:"varint,1000,opt,name=CreatedAt,proto3" json:"CreatedAt,omitempty" sql:"created_at"`
	// @inject_tag: sql:"updated_at"
	UpdatedAt uint32 `protobuf:"varint,1010,opt,name=UpdatedAt,proto3" json:"UpdatedAt,omitempty" sql:"updated_at"`
}

func (x *Achievement) Reset() {
	*x = Achievement{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_inspire_mw_v1_achievement_good_coin_achievement_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Achievement) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Achievement) ProtoMessage() {}

func (x *Achievement) ProtoReflect() protoreflect.Message {
	mi := &file_npool_inspire_mw_v1_achievement_good_coin_achievement_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Achievement.ProtoReflect.Descriptor instead.
func (*Achievement) Descriptor() ([]byte, []int) {
	return file_npool_inspire_mw_v1_achievement_good_coin_achievement_proto_rawDescGZIP(), []int{1}
}

func (x *Achievement) GetID() uint32 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *Achievement) GetEntID() string {
	if x != nil {
		return x.EntID
	}
	return ""
}

func (x *Achievement) GetAppID() string {
	if x != nil {
		return x.AppID
	}
	return ""
}

func (x *Achievement) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

func (x *Achievement) GetGoodCoinTypeID() string {
	if x != nil {
		return x.GoodCoinTypeID
	}
	return ""
}

func (x *Achievement) GetTotalAmountUSD() string {
	if x != nil {
		return x.TotalAmountUSD
	}
	return ""
}

func (x *Achievement) GetSelfAmountUSD() string {
	if x != nil {
		return x.SelfAmountUSD
	}
	return ""
}

func (x *Achievement) GetTotalUnits() string {
	if x != nil {
		return x.TotalUnits
	}
	return ""
}

func (x *Achievement) GetSelfUnits() string {
	if x != nil {
		return x.SelfUnits
	}
	return ""
}

func (x *Achievement) GetTotalCommissionUSD() string {
	if x != nil {
		return x.TotalCommissionUSD
	}
	return ""
}

func (x *Achievement) GetSelfCommissionUSD() string {
	if x != nil {
		return x.SelfCommissionUSD
	}
	return ""
}

func (x *Achievement) GetCreatedAt() uint32 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *Achievement) GetUpdatedAt() uint32 {
	if x != nil {
		return x.UpdatedAt
	}
	return 0
}

type Conds struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EntID          *v1.StringVal      `protobuf:"bytes,10,opt,name=EntID,proto3,oneof" json:"EntID,omitempty"`
	AppID          *v1.StringVal      `protobuf:"bytes,20,opt,name=AppID,proto3,oneof" json:"AppID,omitempty"`
	UserID         *v1.StringVal      `protobuf:"bytes,30,opt,name=UserID,proto3,oneof" json:"UserID,omitempty"`
	UserIDs        *v1.StringSliceVal `protobuf:"bytes,40,opt,name=UserIDs,proto3,oneof" json:"UserIDs,omitempty"`
	GoodCoinTypeID *v1.StringVal      `protobuf:"bytes,50,opt,name=GoodCoinTypeID,proto3,oneof" json:"GoodCoinTypeID,omitempty"`
}

func (x *Conds) Reset() {
	*x = Conds{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_inspire_mw_v1_achievement_good_coin_achievement_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Conds) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Conds) ProtoMessage() {}

func (x *Conds) ProtoReflect() protoreflect.Message {
	mi := &file_npool_inspire_mw_v1_achievement_good_coin_achievement_proto_msgTypes[2]
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
	return file_npool_inspire_mw_v1_achievement_good_coin_achievement_proto_rawDescGZIP(), []int{2}
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

func (x *Conds) GetGoodCoinTypeID() *v1.StringVal {
	if x != nil {
		return x.GoodCoinTypeID
	}
	return nil
}

type GetAchievementsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Conds  *Conds `protobuf:"bytes,10,opt,name=Conds,proto3" json:"Conds,omitempty"`
	Offset int32  `protobuf:"varint,20,opt,name=Offset,proto3" json:"Offset,omitempty"`
	Limit  int32  `protobuf:"varint,30,opt,name=Limit,proto3" json:"Limit,omitempty"`
}

func (x *GetAchievementsRequest) Reset() {
	*x = GetAchievementsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_inspire_mw_v1_achievement_good_coin_achievement_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAchievementsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAchievementsRequest) ProtoMessage() {}

func (x *GetAchievementsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_inspire_mw_v1_achievement_good_coin_achievement_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAchievementsRequest.ProtoReflect.Descriptor instead.
func (*GetAchievementsRequest) Descriptor() ([]byte, []int) {
	return file_npool_inspire_mw_v1_achievement_good_coin_achievement_proto_rawDescGZIP(), []int{3}
}

func (x *GetAchievementsRequest) GetConds() *Conds {
	if x != nil {
		return x.Conds
	}
	return nil
}

func (x *GetAchievementsRequest) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *GetAchievementsRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type GetAchievementsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Infos []*Achievement `protobuf:"bytes,10,rep,name=Infos,proto3" json:"Infos,omitempty"`
	Total uint32         `protobuf:"varint,20,opt,name=Total,proto3" json:"Total,omitempty"`
}

func (x *GetAchievementsResponse) Reset() {
	*x = GetAchievementsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_inspire_mw_v1_achievement_good_coin_achievement_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAchievementsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAchievementsResponse) ProtoMessage() {}

func (x *GetAchievementsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_inspire_mw_v1_achievement_good_coin_achievement_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAchievementsResponse.ProtoReflect.Descriptor instead.
func (*GetAchievementsResponse) Descriptor() ([]byte, []int) {
	return file_npool_inspire_mw_v1_achievement_good_coin_achievement_proto_rawDescGZIP(), []int{4}
}

func (x *GetAchievementsResponse) GetInfos() []*Achievement {
	if x != nil {
		return x.Infos
	}
	return nil
}

func (x *GetAchievementsResponse) GetTotal() uint32 {
	if x != nil {
		return x.Total
	}
	return 0
}

type DeleteAchievementRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *AchievementReq `protobuf:"bytes,10,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (x *DeleteAchievementRequest) Reset() {
	*x = DeleteAchievementRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_inspire_mw_v1_achievement_good_coin_achievement_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteAchievementRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteAchievementRequest) ProtoMessage() {}

func (x *DeleteAchievementRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_inspire_mw_v1_achievement_good_coin_achievement_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteAchievementRequest.ProtoReflect.Descriptor instead.
func (*DeleteAchievementRequest) Descriptor() ([]byte, []int) {
	return file_npool_inspire_mw_v1_achievement_good_coin_achievement_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteAchievementRequest) GetInfo() *AchievementReq {
	if x != nil {
		return x.Info
	}
	return nil
}

type DeleteAchievementResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *Achievement `protobuf:"bytes,10,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (x *DeleteAchievementResponse) Reset() {
	*x = DeleteAchievementResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_inspire_mw_v1_achievement_good_coin_achievement_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteAchievementResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteAchievementResponse) ProtoMessage() {}

func (x *DeleteAchievementResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_inspire_mw_v1_achievement_good_coin_achievement_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteAchievementResponse.ProtoReflect.Descriptor instead.
func (*DeleteAchievementResponse) Descriptor() ([]byte, []int) {
	return file_npool_inspire_mw_v1_achievement_good_coin_achievement_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteAchievementResponse) GetInfo() *Achievement {
	if x != nil {
		return x.Info
	}
	return nil
}

var File_npool_inspire_mw_v1_achievement_good_coin_achievement_proto protoreflect.FileDescriptor

var file_npool_inspire_mw_v1_achievement_good_coin_achievement_proto_rawDesc = []byte{
	0x0a, 0x3b, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x69, 0x6e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2f,
	0x6d, 0x77, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x63, 0x68, 0x69, 0x65, 0x76, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x2f, 0x67, 0x6f, 0x6f, 0x64, 0x2f, 0x63, 0x6f, 0x69, 0x6e, 0x2f, 0x61, 0x63, 0x68, 0x69,
	0x65, 0x76, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x2b, 0x69,
	0x6e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72,
	0x65, 0x2e, 0x61, 0x63, 0x68, 0x69, 0x65, 0x76, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x67, 0x6f,
	0x6f, 0x64, 0x2e, 0x63, 0x6f, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f,
	0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6e,
	0x64, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x91, 0x01, 0x0a, 0x0e, 0x41, 0x63, 0x68,
	0x69, 0x65, 0x76, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x12, 0x13, 0x0a, 0x02, 0x49,
	0x44, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0d, 0x48, 0x00, 0x52, 0x02, 0x49, 0x44, 0x88, 0x01, 0x01,
	0x12, 0x19, 0x0a, 0x05, 0x45, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x48,
	0x01, 0x52, 0x05, 0x45, 0x6e, 0x74, 0x49, 0x44, 0x88, 0x01, 0x01, 0x12, 0x2b, 0x0a, 0x0e, 0x47,
	0x6f, 0x6f, 0x64, 0x43, 0x6f, 0x69, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x49, 0x44, 0x18, 0x32, 0x20,
	0x01, 0x28, 0x09, 0x48, 0x02, 0x52, 0x0e, 0x47, 0x6f, 0x6f, 0x64, 0x43, 0x6f, 0x69, 0x6e, 0x54,
	0x79, 0x70, 0x65, 0x49, 0x44, 0x88, 0x01, 0x01, 0x42, 0x05, 0x0a, 0x03, 0x5f, 0x49, 0x44, 0x42,
	0x08, 0x0a, 0x06, 0x5f, 0x45, 0x6e, 0x74, 0x49, 0x44, 0x42, 0x11, 0x0a, 0x0f, 0x5f, 0x47, 0x6f,
	0x6f, 0x64, 0x43, 0x6f, 0x69, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x49, 0x44, 0x22, 0xb1, 0x03, 0x0a,
	0x0b, 0x41, 0x63, 0x68, 0x69, 0x65, 0x76, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x49, 0x44, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05,
	0x45, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x45, 0x6e, 0x74,
	0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x41, 0x70, 0x70, 0x49, 0x44, 0x18, 0x14, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x41, 0x70, 0x70, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72,
	0x49, 0x44, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44,
	0x12, 0x26, 0x0a, 0x0e, 0x47, 0x6f, 0x6f, 0x64, 0x43, 0x6f, 0x69, 0x6e, 0x54, 0x79, 0x70, 0x65,
	0x49, 0x44, 0x18, 0x28, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x47, 0x6f, 0x6f, 0x64, 0x43, 0x6f,
	0x69, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x49, 0x44, 0x12, 0x26, 0x0a, 0x0e, 0x54, 0x6f, 0x74, 0x61,
	0x6c, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x55, 0x53, 0x44, 0x18, 0x3c, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0e, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x55, 0x53, 0x44,
	0x12, 0x24, 0x0a, 0x0d, 0x53, 0x65, 0x6c, 0x66, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x55, 0x53,
	0x44, 0x18, 0x46, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x53, 0x65, 0x6c, 0x66, 0x41, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x55, 0x53, 0x44, 0x12, 0x1e, 0x0a, 0x0a, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x55,
	0x6e, 0x69, 0x74, 0x73, 0x18, 0x50, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x54, 0x6f, 0x74, 0x61,
	0x6c, 0x55, 0x6e, 0x69, 0x74, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x53, 0x65, 0x6c, 0x66, 0x55, 0x6e,
	0x69, 0x74, 0x73, 0x18, 0x5a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x53, 0x65, 0x6c, 0x66, 0x55,
	0x6e, 0x69, 0x74, 0x73, 0x12, 0x2e, 0x0a, 0x12, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x43, 0x6f, 0x6d,
	0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x55, 0x53, 0x44, 0x18, 0x64, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x12, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x55, 0x53, 0x44, 0x12, 0x2c, 0x0a, 0x11, 0x53, 0x65, 0x6c, 0x66, 0x43, 0x6f, 0x6d, 0x6d,
	0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x55, 0x53, 0x44, 0x18, 0x6e, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x11, 0x53, 0x65, 0x6c, 0x66, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x55,
	0x53, 0x44, 0x12, 0x1d, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18,
	0xe8, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x12, 0x1d, 0x0a, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0xf2,
	0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x22, 0xe6, 0x02, 0x0a, 0x05, 0x43, 0x6f, 0x6e, 0x64, 0x73, 0x12, 0x32, 0x0a, 0x05, 0x45, 0x6e,
	0x74, 0x49, 0x44, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x62, 0x61, 0x73, 0x65,
	0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56,
	0x61, 0x6c, 0x48, 0x00, 0x52, 0x05, 0x45, 0x6e, 0x74, 0x49, 0x44, 0x88, 0x01, 0x01, 0x12, 0x32,
	0x0a, 0x05, 0x41, 0x70, 0x70, 0x49, 0x44, 0x18, 0x14, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e,
	0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x72,
	0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x48, 0x01, 0x52, 0x05, 0x41, 0x70, 0x70, 0x49, 0x44, 0x88,
	0x01, 0x01, 0x12, 0x34, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x1e, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x17, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x48, 0x02, 0x52, 0x06, 0x55,
	0x73, 0x65, 0x72, 0x49, 0x44, 0x88, 0x01, 0x01, 0x12, 0x3b, 0x0a, 0x07, 0x55, 0x73, 0x65, 0x72,
	0x49, 0x44, 0x73, 0x18, 0x28, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x62, 0x61, 0x73, 0x65,
	0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x53,
	0x6c, 0x69, 0x63, 0x65, 0x56, 0x61, 0x6c, 0x48, 0x03, 0x52, 0x07, 0x55, 0x73, 0x65, 0x72, 0x49,
	0x44, 0x73, 0x88, 0x01, 0x01, 0x12, 0x44, 0x0a, 0x0e, 0x47, 0x6f, 0x6f, 0x64, 0x43, 0x6f, 0x69,
	0x6e, 0x54, 0x79, 0x70, 0x65, 0x49, 0x44, 0x18, 0x32, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e,
	0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x72,
	0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x48, 0x04, 0x52, 0x0e, 0x47, 0x6f, 0x6f, 0x64, 0x43, 0x6f,
	0x69, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x49, 0x44, 0x88, 0x01, 0x01, 0x42, 0x08, 0x0a, 0x06, 0x5f,
	0x45, 0x6e, 0x74, 0x49, 0x44, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x41, 0x70, 0x70, 0x49, 0x44, 0x42,
	0x09, 0x0a, 0x07, 0x5f, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x55,
	0x73, 0x65, 0x72, 0x49, 0x44, 0x73, 0x42, 0x11, 0x0a, 0x0f, 0x5f, 0x47, 0x6f, 0x6f, 0x64, 0x43,
	0x6f, 0x69, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x49, 0x44, 0x22, 0x90, 0x01, 0x0a, 0x16, 0x47, 0x65,
	0x74, 0x41, 0x63, 0x68, 0x69, 0x65, 0x76, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x48, 0x0a, 0x05, 0x43, 0x6f, 0x6e, 0x64, 0x73, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x32, 0x2e, 0x69, 0x6e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x6d, 0x69,
	0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x61, 0x63, 0x68, 0x69, 0x65, 0x76, 0x65,
	0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x67, 0x6f, 0x6f, 0x64, 0x2e, 0x63, 0x6f, 0x69, 0x6e, 0x2e, 0x76,
	0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x64, 0x73, 0x52, 0x05, 0x43, 0x6f, 0x6e, 0x64, 0x73, 0x12, 0x16,
	0x0a, 0x06, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x14, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06,
	0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x18,
	0x1e, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x22, 0x7f, 0x0a, 0x17,
	0x47, 0x65, 0x74, 0x41, 0x63, 0x68, 0x69, 0x65, 0x76, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4e, 0x0a, 0x05, 0x49, 0x6e, 0x66, 0x6f, 0x73,
	0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x38, 0x2e, 0x69, 0x6e, 0x73, 0x70, 0x69, 0x72, 0x65,
	0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x61, 0x63, 0x68, 0x69,
	0x65, 0x76, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x67, 0x6f, 0x6f, 0x64, 0x2e, 0x63, 0x6f, 0x69,
	0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x63, 0x68, 0x69, 0x65, 0x76, 0x65, 0x6d, 0x65, 0x6e, 0x74,
	0x52, 0x05, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c,
	0x18, 0x14, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x22, 0x6b, 0x0a,
	0x18, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x63, 0x68, 0x69, 0x65, 0x76, 0x65, 0x6d, 0x65,
	0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x4f, 0x0a, 0x04, 0x49, 0x6e, 0x66,
	0x6f, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3b, 0x2e, 0x69, 0x6e, 0x73, 0x70, 0x69, 0x72,
	0x65, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x61, 0x63, 0x68,
	0x69, 0x65, 0x76, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x67, 0x6f, 0x6f, 0x64, 0x2e, 0x63, 0x6f,
	0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x63, 0x68, 0x69, 0x65, 0x76, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x52, 0x65, 0x71, 0x52, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x69, 0x0a, 0x19, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x41, 0x63, 0x68, 0x69, 0x65, 0x76, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4c, 0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x38, 0x2e, 0x69, 0x6e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e,
	0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x61, 0x63, 0x68, 0x69, 0x65,
	0x76, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x67, 0x6f, 0x6f, 0x64, 0x2e, 0x63, 0x6f, 0x69, 0x6e,
	0x2e, 0x76, 0x31, 0x2e, 0x41, 0x63, 0x68, 0x69, 0x65, 0x76, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x52,
	0x04, 0x49, 0x6e, 0x66, 0x6f, 0x32, 0xbe, 0x03, 0x0a, 0x0a, 0x4d, 0x69, 0x64, 0x64, 0x6c, 0x65,
	0x77, 0x61, 0x72, 0x65, 0x12, 0xd2, 0x01, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x41, 0x63, 0x68, 0x69,
	0x65, 0x76, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x43, 0x2e, 0x69, 0x6e, 0x73, 0x70, 0x69,
	0x72, 0x65, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x61, 0x63,
	0x68, 0x69, 0x65, 0x76, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x67, 0x6f, 0x6f, 0x64, 0x2e, 0x63,
	0x6f, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x63, 0x68, 0x69, 0x65, 0x76,
	0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x44, 0x2e,
	0x69, 0x6e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61,
	0x72, 0x65, 0x2e, 0x61, 0x63, 0x68, 0x69, 0x65, 0x76, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x67,
	0x6f, 0x6f, 0x64, 0x2e, 0x63, 0x6f, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41,
	0x63, 0x68, 0x69, 0x65, 0x76, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x34, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2e, 0x3a, 0x01, 0x2a, 0x22, 0x29,
	0x2f, 0x76, 0x31, 0x2f, 0x67, 0x6f, 0x6f, 0x64, 0x5f, 0x63, 0x6f, 0x69, 0x6e, 0x5f, 0x61, 0x63,
	0x68, 0x69, 0x65, 0x76, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x47, 0x65, 0x74, 0x41, 0x63, 0x68,
	0x69, 0x65, 0x76, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0xda, 0x01, 0x0a, 0x11, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x41, 0x63, 0x68, 0x69, 0x65, 0x76, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x12,
	0x45, 0x2e, 0x69, 0x6e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65,
	0x77, 0x61, 0x72, 0x65, 0x2e, 0x61, 0x63, 0x68, 0x69, 0x65, 0x76, 0x65, 0x6d, 0x65, 0x6e, 0x74,
	0x2e, 0x67, 0x6f, 0x6f, 0x64, 0x2e, 0x63, 0x6f, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x41, 0x63, 0x68, 0x69, 0x65, 0x76, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x46, 0x2e, 0x69, 0x6e, 0x73, 0x70, 0x69, 0x72, 0x65,
	0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x61, 0x63, 0x68, 0x69,
	0x65, 0x76, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x67, 0x6f, 0x6f, 0x64, 0x2e, 0x63, 0x6f, 0x69,
	0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x63, 0x68, 0x69, 0x65,
	0x76, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x36,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x30, 0x3a, 0x01, 0x2a, 0x22, 0x2b, 0x2f, 0x76, 0x31, 0x2f, 0x67,
	0x6f, 0x6f, 0x64, 0x5f, 0x63, 0x6f, 0x69, 0x6e, 0x5f, 0x61, 0x63, 0x68, 0x69, 0x65, 0x76, 0x65,
	0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x63, 0x68, 0x69, 0x65,
	0x76, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x42, 0x4c, 0x5a, 0x4a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4e, 0x70, 0x6f, 0x6f, 0x6c, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f,
	0x72, 0x6d, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x6e, 0x70, 0x6f, 0x6f, 0x6c,
	0x2f, 0x69, 0x6e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2f, 0x6d, 0x77, 0x2f, 0x76, 0x31, 0x2f, 0x61,
	0x63, 0x68, 0x69, 0x65, 0x76, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x67, 0x6f, 0x6f, 0x64, 0x2f,
	0x63, 0x6f, 0x69, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_npool_inspire_mw_v1_achievement_good_coin_achievement_proto_rawDescOnce sync.Once
	file_npool_inspire_mw_v1_achievement_good_coin_achievement_proto_rawDescData = file_npool_inspire_mw_v1_achievement_good_coin_achievement_proto_rawDesc
)

func file_npool_inspire_mw_v1_achievement_good_coin_achievement_proto_rawDescGZIP() []byte {
	file_npool_inspire_mw_v1_achievement_good_coin_achievement_proto_rawDescOnce.Do(func() {
		file_npool_inspire_mw_v1_achievement_good_coin_achievement_proto_rawDescData = protoimpl.X.CompressGZIP(file_npool_inspire_mw_v1_achievement_good_coin_achievement_proto_rawDescData)
	})
	return file_npool_inspire_mw_v1_achievement_good_coin_achievement_proto_rawDescData
}

var file_npool_inspire_mw_v1_achievement_good_coin_achievement_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_npool_inspire_mw_v1_achievement_good_coin_achievement_proto_goTypes = []interface{}{
	(*AchievementReq)(nil),            // 0: inspire.middleware.achievement.good.coin.v1.AchievementReq
	(*Achievement)(nil),               // 1: inspire.middleware.achievement.good.coin.v1.Achievement
	(*Conds)(nil),                     // 2: inspire.middleware.achievement.good.coin.v1.Conds
	(*GetAchievementsRequest)(nil),    // 3: inspire.middleware.achievement.good.coin.v1.GetAchievementsRequest
	(*GetAchievementsResponse)(nil),   // 4: inspire.middleware.achievement.good.coin.v1.GetAchievementsResponse
	(*DeleteAchievementRequest)(nil),  // 5: inspire.middleware.achievement.good.coin.v1.DeleteAchievementRequest
	(*DeleteAchievementResponse)(nil), // 6: inspire.middleware.achievement.good.coin.v1.DeleteAchievementResponse
	(*v1.StringVal)(nil),              // 7: basetypes.v1.StringVal
	(*v1.StringSliceVal)(nil),         // 8: basetypes.v1.StringSliceVal
}
var file_npool_inspire_mw_v1_achievement_good_coin_achievement_proto_depIdxs = []int32{
	7,  // 0: inspire.middleware.achievement.good.coin.v1.Conds.EntID:type_name -> basetypes.v1.StringVal
	7,  // 1: inspire.middleware.achievement.good.coin.v1.Conds.AppID:type_name -> basetypes.v1.StringVal
	7,  // 2: inspire.middleware.achievement.good.coin.v1.Conds.UserID:type_name -> basetypes.v1.StringVal
	8,  // 3: inspire.middleware.achievement.good.coin.v1.Conds.UserIDs:type_name -> basetypes.v1.StringSliceVal
	7,  // 4: inspire.middleware.achievement.good.coin.v1.Conds.GoodCoinTypeID:type_name -> basetypes.v1.StringVal
	2,  // 5: inspire.middleware.achievement.good.coin.v1.GetAchievementsRequest.Conds:type_name -> inspire.middleware.achievement.good.coin.v1.Conds
	1,  // 6: inspire.middleware.achievement.good.coin.v1.GetAchievementsResponse.Infos:type_name -> inspire.middleware.achievement.good.coin.v1.Achievement
	0,  // 7: inspire.middleware.achievement.good.coin.v1.DeleteAchievementRequest.Info:type_name -> inspire.middleware.achievement.good.coin.v1.AchievementReq
	1,  // 8: inspire.middleware.achievement.good.coin.v1.DeleteAchievementResponse.Info:type_name -> inspire.middleware.achievement.good.coin.v1.Achievement
	3,  // 9: inspire.middleware.achievement.good.coin.v1.Middleware.GetAchievements:input_type -> inspire.middleware.achievement.good.coin.v1.GetAchievementsRequest
	5,  // 10: inspire.middleware.achievement.good.coin.v1.Middleware.DeleteAchievement:input_type -> inspire.middleware.achievement.good.coin.v1.DeleteAchievementRequest
	4,  // 11: inspire.middleware.achievement.good.coin.v1.Middleware.GetAchievements:output_type -> inspire.middleware.achievement.good.coin.v1.GetAchievementsResponse
	6,  // 12: inspire.middleware.achievement.good.coin.v1.Middleware.DeleteAchievement:output_type -> inspire.middleware.achievement.good.coin.v1.DeleteAchievementResponse
	11, // [11:13] is the sub-list for method output_type
	9,  // [9:11] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_npool_inspire_mw_v1_achievement_good_coin_achievement_proto_init() }
func file_npool_inspire_mw_v1_achievement_good_coin_achievement_proto_init() {
	if File_npool_inspire_mw_v1_achievement_good_coin_achievement_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_npool_inspire_mw_v1_achievement_good_coin_achievement_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AchievementReq); i {
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
		file_npool_inspire_mw_v1_achievement_good_coin_achievement_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Achievement); i {
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
		file_npool_inspire_mw_v1_achievement_good_coin_achievement_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
		file_npool_inspire_mw_v1_achievement_good_coin_achievement_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAchievementsRequest); i {
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
		file_npool_inspire_mw_v1_achievement_good_coin_achievement_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAchievementsResponse); i {
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
		file_npool_inspire_mw_v1_achievement_good_coin_achievement_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteAchievementRequest); i {
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
		file_npool_inspire_mw_v1_achievement_good_coin_achievement_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteAchievementResponse); i {
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
	file_npool_inspire_mw_v1_achievement_good_coin_achievement_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_npool_inspire_mw_v1_achievement_good_coin_achievement_proto_msgTypes[2].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_npool_inspire_mw_v1_achievement_good_coin_achievement_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_npool_inspire_mw_v1_achievement_good_coin_achievement_proto_goTypes,
		DependencyIndexes: file_npool_inspire_mw_v1_achievement_good_coin_achievement_proto_depIdxs,
		MessageInfos:      file_npool_inspire_mw_v1_achievement_good_coin_achievement_proto_msgTypes,
	}.Build()
	File_npool_inspire_mw_v1_achievement_good_coin_achievement_proto = out.File
	file_npool_inspire_mw_v1_achievement_good_coin_achievement_proto_rawDesc = nil
	file_npool_inspire_mw_v1_achievement_good_coin_achievement_proto_goTypes = nil
	file_npool_inspire_mw_v1_achievement_good_coin_achievement_proto_depIdxs = nil
}
