// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.18.1
// source: npool/good/mw/v1/good/reward/history/history.proto

package history

import (
	_ "github.com/NpoolPlatform/message/npool/basetypes/good/v1"
	v1 "github.com/NpoolPlatform/message/npool/basetypes/v1"
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

type HistoryReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID         *string `protobuf:"bytes,10,opt,name=ID,proto3,oneof" json:"ID,omitempty"`
	GoodID     *string `protobuf:"bytes,20,opt,name=GoodID,proto3,oneof" json:"GoodID,omitempty"`
	RewardDate *uint32 `protobuf:"varint,30,opt,name=RewardDate,proto3,oneof" json:"RewardDate,omitempty"`
	TID        *string `protobuf:"bytes,40,opt,name=TID,proto3,oneof" json:"TID,omitempty"`
	Amount     *string `protobuf:"bytes,50,opt,name=Amount,proto3,oneof" json:"Amount,omitempty"`
	UnitAmount *string `protobuf:"bytes,60,opt,name=UnitAmount,proto3,oneof" json:"UnitAmount,omitempty"`
}

func (x *HistoryReq) Reset() {
	*x = HistoryReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_good_mw_v1_good_reward_history_history_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HistoryReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HistoryReq) ProtoMessage() {}

func (x *HistoryReq) ProtoReflect() protoreflect.Message {
	mi := &file_npool_good_mw_v1_good_reward_history_history_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HistoryReq.ProtoReflect.Descriptor instead.
func (*HistoryReq) Descriptor() ([]byte, []int) {
	return file_npool_good_mw_v1_good_reward_history_history_proto_rawDescGZIP(), []int{0}
}

func (x *HistoryReq) GetID() string {
	if x != nil && x.ID != nil {
		return *x.ID
	}
	return ""
}

func (x *HistoryReq) GetGoodID() string {
	if x != nil && x.GoodID != nil {
		return *x.GoodID
	}
	return ""
}

func (x *HistoryReq) GetRewardDate() uint32 {
	if x != nil && x.RewardDate != nil {
		return *x.RewardDate
	}
	return 0
}

func (x *HistoryReq) GetTID() string {
	if x != nil && x.TID != nil {
		return *x.TID
	}
	return ""
}

func (x *HistoryReq) GetAmount() string {
	if x != nil && x.Amount != nil {
		return *x.Amount
	}
	return ""
}

func (x *HistoryReq) GetUnitAmount() string {
	if x != nil && x.UnitAmount != nil {
		return *x.UnitAmount
	}
	return ""
}

type History struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: sql:"id"
	ID string `protobuf:"bytes,10,opt,name=ID,proto3" json:"ID,omitempty" sql:"id"`
	// @inject_tag: sql:"good_id"
	GoodID string `protobuf:"bytes,20,opt,name=GoodID,proto3" json:"GoodID,omitempty" sql:"good_id"`
	// @inject_tag: sql:"good_name"
	GoodName string `protobuf:"bytes,30,opt,name=GoodName,proto3" json:"GoodName,omitempty" sql:"good_name"`
	// @inject_tag: sql:"reward_date"
	RewardDate uint32 `protobuf:"varint,40,opt,name=RewardDate,proto3" json:"RewardDate,omitempty" sql:"reward_date"`
	// @inject_tag: sql:"tid"
	TID string `protobuf:"bytes,50,opt,name=TID,proto3" json:"TID,omitempty" sql:"tid"`
	// @inject_tag: sql:"amount"
	Amount string `protobuf:"bytes,60,opt,name=Amount,proto3" json:"Amount,omitempty" sql:"amount"`
	// @inject_tag: sql:"unit_amount"
	UnitAmount string `protobuf:"bytes,70,opt,name=UnitAmount,proto3" json:"UnitAmount,omitempty" sql:"unit_amount"`
	// @inject_tag: sql:"created_at"
	CreatedAt uint32 `protobuf:"varint,1000,opt,name=CreatedAt,proto3" json:"CreatedAt,omitempty" sql:"created_at"`
	// @inject_tag: sql:"updated_at"
	UpdatedAt uint32 `protobuf:"varint,1010,opt,name=UpdatedAt,proto3" json:"UpdatedAt,omitempty" sql:"updated_at"`
}

func (x *History) Reset() {
	*x = History{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_good_mw_v1_good_reward_history_history_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *History) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*History) ProtoMessage() {}

func (x *History) ProtoReflect() protoreflect.Message {
	mi := &file_npool_good_mw_v1_good_reward_history_history_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use History.ProtoReflect.Descriptor instead.
func (*History) Descriptor() ([]byte, []int) {
	return file_npool_good_mw_v1_good_reward_history_history_proto_rawDescGZIP(), []int{1}
}

func (x *History) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *History) GetGoodID() string {
	if x != nil {
		return x.GoodID
	}
	return ""
}

func (x *History) GetGoodName() string {
	if x != nil {
		return x.GoodName
	}
	return ""
}

func (x *History) GetRewardDate() uint32 {
	if x != nil {
		return x.RewardDate
	}
	return 0
}

func (x *History) GetTID() string {
	if x != nil {
		return x.TID
	}
	return ""
}

func (x *History) GetAmount() string {
	if x != nil {
		return x.Amount
	}
	return ""
}

func (x *History) GetUnitAmount() string {
	if x != nil {
		return x.UnitAmount
	}
	return ""
}

func (x *History) GetCreatedAt() uint32 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *History) GetUpdatedAt() uint32 {
	if x != nil {
		return x.UpdatedAt
	}
	return 0
}

type Conds struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID         *v1.StringVal      `protobuf:"bytes,10,opt,name=ID,proto3,oneof" json:"ID,omitempty"`
	GoodID     *v1.StringVal      `protobuf:"bytes,20,opt,name=GoodID,proto3,oneof" json:"GoodID,omitempty"`
	GoodIDs    *v1.StringSliceVal `protobuf:"bytes,30,opt,name=GoodIDs,proto3,oneof" json:"GoodIDs,omitempty"`
	RewardDate *v1.Uint32Val      `protobuf:"bytes,40,opt,name=RewardDate,proto3,oneof" json:"RewardDate,omitempty"`
}

func (x *Conds) Reset() {
	*x = Conds{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_good_mw_v1_good_reward_history_history_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Conds) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Conds) ProtoMessage() {}

func (x *Conds) ProtoReflect() protoreflect.Message {
	mi := &file_npool_good_mw_v1_good_reward_history_history_proto_msgTypes[2]
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
	return file_npool_good_mw_v1_good_reward_history_history_proto_rawDescGZIP(), []int{2}
}

func (x *Conds) GetID() *v1.StringVal {
	if x != nil {
		return x.ID
	}
	return nil
}

func (x *Conds) GetGoodID() *v1.StringVal {
	if x != nil {
		return x.GoodID
	}
	return nil
}

func (x *Conds) GetGoodIDs() *v1.StringSliceVal {
	if x != nil {
		return x.GoodIDs
	}
	return nil
}

func (x *Conds) GetRewardDate() *v1.Uint32Val {
	if x != nil {
		return x.RewardDate
	}
	return nil
}

type GetHistoriesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Conds  *Conds `protobuf:"bytes,10,opt,name=Conds,proto3" json:"Conds,omitempty"`
	Offset int32  `protobuf:"varint,20,opt,name=Offset,proto3" json:"Offset,omitempty"`
	Limit  int32  `protobuf:"varint,30,opt,name=Limit,proto3" json:"Limit,omitempty"`
}

func (x *GetHistoriesRequest) Reset() {
	*x = GetHistoriesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_good_mw_v1_good_reward_history_history_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetHistoriesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetHistoriesRequest) ProtoMessage() {}

func (x *GetHistoriesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_good_mw_v1_good_reward_history_history_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetHistoriesRequest.ProtoReflect.Descriptor instead.
func (*GetHistoriesRequest) Descriptor() ([]byte, []int) {
	return file_npool_good_mw_v1_good_reward_history_history_proto_rawDescGZIP(), []int{3}
}

func (x *GetHistoriesRequest) GetConds() *Conds {
	if x != nil {
		return x.Conds
	}
	return nil
}

func (x *GetHistoriesRequest) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *GetHistoriesRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type GetHistoriesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Infos []*History `protobuf:"bytes,10,rep,name=Infos,proto3" json:"Infos,omitempty"`
	Total uint32     `protobuf:"varint,20,opt,name=Total,proto3" json:"Total,omitempty"`
}

func (x *GetHistoriesResponse) Reset() {
	*x = GetHistoriesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_good_mw_v1_good_reward_history_history_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetHistoriesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetHistoriesResponse) ProtoMessage() {}

func (x *GetHistoriesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_good_mw_v1_good_reward_history_history_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetHistoriesResponse.ProtoReflect.Descriptor instead.
func (*GetHistoriesResponse) Descriptor() ([]byte, []int) {
	return file_npool_good_mw_v1_good_reward_history_history_proto_rawDescGZIP(), []int{4}
}

func (x *GetHistoriesResponse) GetInfos() []*History {
	if x != nil {
		return x.Infos
	}
	return nil
}

func (x *GetHistoriesResponse) GetTotal() uint32 {
	if x != nil {
		return x.Total
	}
	return 0
}

var File_npool_good_mw_v1_good_reward_history_history_proto protoreflect.FileDescriptor

var file_npool_good_mw_v1_good_reward_history_history_proto_rawDesc = []byte{
	0x0a, 0x32, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x67, 0x6f, 0x6f, 0x64, 0x2f, 0x6d, 0x77, 0x2f,
	0x76, 0x31, 0x2f, 0x67, 0x6f, 0x6f, 0x64, 0x2f, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x2f, 0x68,
	0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x2f, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x27, 0x67, 0x6f, 0x6f, 0x64, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c,
	0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x67, 0x6f, 0x6f, 0x64, 0x31, 0x2e, 0x72, 0x65, 0x77, 0x61,
	0x72, 0x64, 0x2e, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x1a, 0x1e, 0x6e,
	0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x76,
	0x31, 0x2f, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x23, 0x6e,
	0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x67,
	0x6f, 0x6f, 0x64, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xff, 0x01, 0x0a, 0x0a, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65,
	0x71, 0x12, 0x13, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52,
	0x02, 0x49, 0x44, 0x88, 0x01, 0x01, 0x12, 0x1b, 0x0a, 0x06, 0x47, 0x6f, 0x6f, 0x64, 0x49, 0x44,
	0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x06, 0x47, 0x6f, 0x6f, 0x64, 0x49, 0x44,
	0x88, 0x01, 0x01, 0x12, 0x23, 0x0a, 0x0a, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x44, 0x61, 0x74,
	0x65, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x0d, 0x48, 0x02, 0x52, 0x0a, 0x52, 0x65, 0x77, 0x61, 0x72,
	0x64, 0x44, 0x61, 0x74, 0x65, 0x88, 0x01, 0x01, 0x12, 0x15, 0x0a, 0x03, 0x54, 0x49, 0x44, 0x18,
	0x28, 0x20, 0x01, 0x28, 0x09, 0x48, 0x03, 0x52, 0x03, 0x54, 0x49, 0x44, 0x88, 0x01, 0x01, 0x12,
	0x1b, 0x0a, 0x06, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x32, 0x20, 0x01, 0x28, 0x09, 0x48,
	0x04, 0x52, 0x06, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x88, 0x01, 0x01, 0x12, 0x23, 0x0a, 0x0a,
	0x55, 0x6e, 0x69, 0x74, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x3c, 0x20, 0x01, 0x28, 0x09,
	0x48, 0x05, 0x52, 0x0a, 0x55, 0x6e, 0x69, 0x74, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x88, 0x01,
	0x01, 0x42, 0x05, 0x0a, 0x03, 0x5f, 0x49, 0x44, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x47, 0x6f, 0x6f,
	0x64, 0x49, 0x44, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x44, 0x61,
	0x74, 0x65, 0x42, 0x06, 0x0a, 0x04, 0x5f, 0x54, 0x49, 0x44, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x41,
	0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x55, 0x6e, 0x69, 0x74, 0x41, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x22, 0xf5, 0x01, 0x0a, 0x07, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79,
	0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44,
	0x12, 0x16, 0x0a, 0x06, 0x47, 0x6f, 0x6f, 0x64, 0x49, 0x44, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x47, 0x6f, 0x6f, 0x64, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08, 0x47, 0x6f, 0x6f, 0x64,
	0x4e, 0x61, 0x6d, 0x65, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x47, 0x6f, 0x6f, 0x64,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x44, 0x61,
	0x74, 0x65, 0x18, 0x28, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64,
	0x44, 0x61, 0x74, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x54, 0x49, 0x44, 0x18, 0x32, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x54, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x3c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1e,
	0x0a, 0x0a, 0x55, 0x6e, 0x69, 0x74, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x46, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x55, 0x6e, 0x69, 0x74, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1d,
	0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0xe8, 0x07, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a,
	0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0xf2, 0x07, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x93, 0x02, 0x0a,
	0x05, 0x43, 0x6f, 0x6e, 0x64, 0x73, 0x12, 0x2c, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x17, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x48, 0x00, 0x52, 0x02, 0x49,
	0x44, 0x88, 0x01, 0x01, 0x12, 0x34, 0x0a, 0x06, 0x47, 0x6f, 0x6f, 0x64, 0x49, 0x44, 0x18, 0x14,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x48, 0x01, 0x52,
	0x06, 0x47, 0x6f, 0x6f, 0x64, 0x49, 0x44, 0x88, 0x01, 0x01, 0x12, 0x3b, 0x0a, 0x07, 0x47, 0x6f,
	0x6f, 0x64, 0x49, 0x44, 0x73, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x62, 0x61,
	0x73, 0x65, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e,
	0x67, 0x53, 0x6c, 0x69, 0x63, 0x65, 0x56, 0x61, 0x6c, 0x48, 0x02, 0x52, 0x07, 0x47, 0x6f, 0x6f,
	0x64, 0x49, 0x44, 0x73, 0x88, 0x01, 0x01, 0x12, 0x3c, 0x0a, 0x0a, 0x52, 0x65, 0x77, 0x61, 0x72,
	0x64, 0x44, 0x61, 0x74, 0x65, 0x18, 0x28, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x62, 0x61,
	0x73, 0x65, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x69, 0x6e, 0x74, 0x33,
	0x32, 0x56, 0x61, 0x6c, 0x48, 0x03, 0x52, 0x0a, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x44, 0x61,
	0x74, 0x65, 0x88, 0x01, 0x01, 0x42, 0x05, 0x0a, 0x03, 0x5f, 0x49, 0x44, 0x42, 0x09, 0x0a, 0x07,
	0x5f, 0x47, 0x6f, 0x6f, 0x64, 0x49, 0x44, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x47, 0x6f, 0x6f, 0x64,
	0x49, 0x44, 0x73, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x44, 0x61,
	0x74, 0x65, 0x22, 0x89, 0x01, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72,
	0x69, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x44, 0x0a, 0x05, 0x43, 0x6f,
	0x6e, 0x64, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x67, 0x6f, 0x6f, 0x64,
	0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x67, 0x6f, 0x6f, 0x64,
	0x31, 0x2e, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x2e, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79,
	0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x64, 0x73, 0x52, 0x05, 0x43, 0x6f, 0x6e, 0x64, 0x73,
	0x12, 0x16, 0x0a, 0x06, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x14, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x06, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x4c, 0x69, 0x6d, 0x69,
	0x74, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x22, 0x74,
	0x0a, 0x14, 0x47, 0x65, 0x74, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x46, 0x0a, 0x05, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x18,
	0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x67, 0x6f, 0x6f, 0x64, 0x2e, 0x6d, 0x69, 0x64,
	0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x67, 0x6f, 0x6f, 0x64, 0x31, 0x2e, 0x72, 0x65,
	0x77, 0x61, 0x72, 0x64, 0x2e, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x2e,
	0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x05, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x12, 0x14,
	0x0a, 0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x14, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x54,
	0x6f, 0x74, 0x61, 0x6c, 0x32, 0x9c, 0x01, 0x0a, 0x0a, 0x4d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77,
	0x61, 0x72, 0x65, 0x12, 0x8d, 0x01, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x48, 0x69, 0x73, 0x74, 0x6f,
	0x72, 0x69, 0x65, 0x73, 0x12, 0x3c, 0x2e, 0x67, 0x6f, 0x6f, 0x64, 0x2e, 0x6d, 0x69, 0x64, 0x64,
	0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x67, 0x6f, 0x6f, 0x64, 0x31, 0x2e, 0x72, 0x65, 0x77,
	0x61, 0x72, 0x64, 0x2e, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x47,
	0x65, 0x74, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x3d, 0x2e, 0x67, 0x6f, 0x6f, 0x64, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65,
	0x77, 0x61, 0x72, 0x65, 0x2e, 0x67, 0x6f, 0x6f, 0x64, 0x31, 0x2e, 0x72, 0x65, 0x77, 0x61, 0x72,
	0x64, 0x2e, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74,
	0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x42, 0x47, 0x5a, 0x45, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x4e, 0x70, 0x6f, 0x6f, 0x6c, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x67, 0x6f,
	0x6f, 0x64, 0x2f, 0x6d, 0x77, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x6f, 0x6f, 0x64, 0x2f, 0x72, 0x65,
	0x77, 0x61, 0x72, 0x64, 0x2f, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_npool_good_mw_v1_good_reward_history_history_proto_rawDescOnce sync.Once
	file_npool_good_mw_v1_good_reward_history_history_proto_rawDescData = file_npool_good_mw_v1_good_reward_history_history_proto_rawDesc
)

func file_npool_good_mw_v1_good_reward_history_history_proto_rawDescGZIP() []byte {
	file_npool_good_mw_v1_good_reward_history_history_proto_rawDescOnce.Do(func() {
		file_npool_good_mw_v1_good_reward_history_history_proto_rawDescData = protoimpl.X.CompressGZIP(file_npool_good_mw_v1_good_reward_history_history_proto_rawDescData)
	})
	return file_npool_good_mw_v1_good_reward_history_history_proto_rawDescData
}

var file_npool_good_mw_v1_good_reward_history_history_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_npool_good_mw_v1_good_reward_history_history_proto_goTypes = []interface{}{
	(*HistoryReq)(nil),           // 0: good.middleware.good1.reward.history.v1.HistoryReq
	(*History)(nil),              // 1: good.middleware.good1.reward.history.v1.History
	(*Conds)(nil),                // 2: good.middleware.good1.reward.history.v1.Conds
	(*GetHistoriesRequest)(nil),  // 3: good.middleware.good1.reward.history.v1.GetHistoriesRequest
	(*GetHistoriesResponse)(nil), // 4: good.middleware.good1.reward.history.v1.GetHistoriesResponse
	(*v1.StringVal)(nil),         // 5: basetypes.v1.StringVal
	(*v1.StringSliceVal)(nil),    // 6: basetypes.v1.StringSliceVal
	(*v1.Uint32Val)(nil),         // 7: basetypes.v1.Uint32Val
}
var file_npool_good_mw_v1_good_reward_history_history_proto_depIdxs = []int32{
	5, // 0: good.middleware.good1.reward.history.v1.Conds.ID:type_name -> basetypes.v1.StringVal
	5, // 1: good.middleware.good1.reward.history.v1.Conds.GoodID:type_name -> basetypes.v1.StringVal
	6, // 2: good.middleware.good1.reward.history.v1.Conds.GoodIDs:type_name -> basetypes.v1.StringSliceVal
	7, // 3: good.middleware.good1.reward.history.v1.Conds.RewardDate:type_name -> basetypes.v1.Uint32Val
	2, // 4: good.middleware.good1.reward.history.v1.GetHistoriesRequest.Conds:type_name -> good.middleware.good1.reward.history.v1.Conds
	1, // 5: good.middleware.good1.reward.history.v1.GetHistoriesResponse.Infos:type_name -> good.middleware.good1.reward.history.v1.History
	3, // 6: good.middleware.good1.reward.history.v1.Middleware.GetHistories:input_type -> good.middleware.good1.reward.history.v1.GetHistoriesRequest
	4, // 7: good.middleware.good1.reward.history.v1.Middleware.GetHistories:output_type -> good.middleware.good1.reward.history.v1.GetHistoriesResponse
	7, // [7:8] is the sub-list for method output_type
	6, // [6:7] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_npool_good_mw_v1_good_reward_history_history_proto_init() }
func file_npool_good_mw_v1_good_reward_history_history_proto_init() {
	if File_npool_good_mw_v1_good_reward_history_history_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_npool_good_mw_v1_good_reward_history_history_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HistoryReq); i {
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
		file_npool_good_mw_v1_good_reward_history_history_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*History); i {
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
		file_npool_good_mw_v1_good_reward_history_history_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
		file_npool_good_mw_v1_good_reward_history_history_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetHistoriesRequest); i {
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
		file_npool_good_mw_v1_good_reward_history_history_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetHistoriesResponse); i {
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
	file_npool_good_mw_v1_good_reward_history_history_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_npool_good_mw_v1_good_reward_history_history_proto_msgTypes[2].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_npool_good_mw_v1_good_reward_history_history_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_npool_good_mw_v1_good_reward_history_history_proto_goTypes,
		DependencyIndexes: file_npool_good_mw_v1_good_reward_history_history_proto_depIdxs,
		MessageInfos:      file_npool_good_mw_v1_good_reward_history_history_proto_msgTypes,
	}.Build()
	File_npool_good_mw_v1_good_reward_history_history_proto = out.File
	file_npool_good_mw_v1_good_reward_history_history_proto_rawDesc = nil
	file_npool_good_mw_v1_good_reward_history_history_proto_goTypes = nil
	file_npool_good_mw_v1_good_reward_history_history_proto_depIdxs = nil
}
