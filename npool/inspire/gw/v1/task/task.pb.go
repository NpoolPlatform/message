// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.18.1
// source: npool/inspire/gw/v1/task/task.proto

package task

import (
	v1 "github.com/NpoolPlatform/message/npool/basetypes/inspire/v1"
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

type UserTask struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID               uint32         `protobuf:"varint,10,opt,name=ID,proto3" json:"ID,omitempty"`
	EntID            string         `protobuf:"bytes,20,opt,name=EntID,proto3" json:"EntID,omitempty"`
	AppID            string         `protobuf:"bytes,30,opt,name=AppID,proto3" json:"AppID,omitempty"`
	EventID          string         `protobuf:"bytes,40,opt,name=EventID,proto3" json:"EventID,omitempty"`
	TaskType         v1.TaskType    `protobuf:"varint,50,opt,name=TaskType,proto3,enum=basetypes.inspire.v1.TaskType" json:"TaskType,omitempty"`
	Name             string         `protobuf:"bytes,60,opt,name=Name,proto3" json:"Name,omitempty"`
	TaskDesc         string         `protobuf:"bytes,70,opt,name=TaskDesc,proto3" json:"TaskDesc,omitempty"`
	StepGuide        string         `protobuf:"bytes,80,opt,name=StepGuide,proto3" json:"StepGuide,omitempty"`
	RecommendMessage string         `protobuf:"bytes,90,opt,name=RecommendMessage,proto3" json:"RecommendMessage,omitempty"`
	Index            uint32         `protobuf:"varint,100,opt,name=Index,proto3" json:"Index,omitempty"`
	LastTaskID       string         `protobuf:"bytes,110,opt,name=LastTaskID,proto3" json:"LastTaskID,omitempty"`
	MaxRewardCount   uint32         `protobuf:"varint,120,opt,name=MaxRewardCount,proto3" json:"MaxRewardCount,omitempty"`
	CooldownSecond   uint32         `protobuf:"varint,130,opt,name=CooldownSecond,proto3" json:"CooldownSecond,omitempty"`
	CompletionTimes  uint32         `protobuf:"varint,140,opt,name=CompletionTimes,proto3" json:"CompletionTimes,omitempty"`
	NextStartAt      uint32         `protobuf:"varint,150,opt,name=NextStartAt,proto3" json:"NextStartAt,omitempty"`
	TaskState        v1.TaskState   `protobuf:"varint,160,opt,name=TaskState,proto3,enum=basetypes.inspire.v1.TaskState" json:"TaskState,omitempty"`
	RewardState      v1.RewardState `protobuf:"varint,170,opt,name=RewardState,proto3,enum=basetypes.inspire.v1.RewardState" json:"RewardState,omitempty"`
	CreatedAt        uint32         `protobuf:"varint,1000,opt,name=CreatedAt,proto3" json:"CreatedAt,omitempty"`
	UpdatedAt        uint32         `protobuf:"varint,1010,opt,name=UpdatedAt,proto3" json:"UpdatedAt,omitempty"`
}

func (x *UserTask) Reset() {
	*x = UserTask{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_inspire_gw_v1_task_task_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserTask) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserTask) ProtoMessage() {}

func (x *UserTask) ProtoReflect() protoreflect.Message {
	mi := &file_npool_inspire_gw_v1_task_task_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserTask.ProtoReflect.Descriptor instead.
func (*UserTask) Descriptor() ([]byte, []int) {
	return file_npool_inspire_gw_v1_task_task_proto_rawDescGZIP(), []int{0}
}

func (x *UserTask) GetID() uint32 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *UserTask) GetEntID() string {
	if x != nil {
		return x.EntID
	}
	return ""
}

func (x *UserTask) GetAppID() string {
	if x != nil {
		return x.AppID
	}
	return ""
}

func (x *UserTask) GetEventID() string {
	if x != nil {
		return x.EventID
	}
	return ""
}

func (x *UserTask) GetTaskType() v1.TaskType {
	if x != nil {
		return x.TaskType
	}
	return v1.TaskType(0)
}

func (x *UserTask) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UserTask) GetTaskDesc() string {
	if x != nil {
		return x.TaskDesc
	}
	return ""
}

func (x *UserTask) GetStepGuide() string {
	if x != nil {
		return x.StepGuide
	}
	return ""
}

func (x *UserTask) GetRecommendMessage() string {
	if x != nil {
		return x.RecommendMessage
	}
	return ""
}

func (x *UserTask) GetIndex() uint32 {
	if x != nil {
		return x.Index
	}
	return 0
}

func (x *UserTask) GetLastTaskID() string {
	if x != nil {
		return x.LastTaskID
	}
	return ""
}

func (x *UserTask) GetMaxRewardCount() uint32 {
	if x != nil {
		return x.MaxRewardCount
	}
	return 0
}

func (x *UserTask) GetCooldownSecond() uint32 {
	if x != nil {
		return x.CooldownSecond
	}
	return 0
}

func (x *UserTask) GetCompletionTimes() uint32 {
	if x != nil {
		return x.CompletionTimes
	}
	return 0
}

func (x *UserTask) GetNextStartAt() uint32 {
	if x != nil {
		return x.NextStartAt
	}
	return 0
}

func (x *UserTask) GetTaskState() v1.TaskState {
	if x != nil {
		return x.TaskState
	}
	return v1.TaskState(0)
}

func (x *UserTask) GetRewardState() v1.RewardState {
	if x != nil {
		return x.RewardState
	}
	return v1.RewardState(0)
}

func (x *UserTask) GetCreatedAt() uint32 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *UserTask) GetUpdatedAt() uint32 {
	if x != nil {
		return x.UpdatedAt
	}
	return 0
}

type GetMyTasksRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AppID  string `protobuf:"bytes,10,opt,name=AppID,proto3" json:"AppID,omitempty"`
	UserID string `protobuf:"bytes,20,opt,name=UserID,proto3" json:"UserID,omitempty"`
	Offset int32  `protobuf:"varint,30,opt,name=Offset,proto3" json:"Offset,omitempty"`
	Limit  int32  `protobuf:"varint,40,opt,name=Limit,proto3" json:"Limit,omitempty"`
}

func (x *GetMyTasksRequest) Reset() {
	*x = GetMyTasksRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_inspire_gw_v1_task_task_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMyTasksRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMyTasksRequest) ProtoMessage() {}

func (x *GetMyTasksRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_inspire_gw_v1_task_task_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMyTasksRequest.ProtoReflect.Descriptor instead.
func (*GetMyTasksRequest) Descriptor() ([]byte, []int) {
	return file_npool_inspire_gw_v1_task_task_proto_rawDescGZIP(), []int{1}
}

func (x *GetMyTasksRequest) GetAppID() string {
	if x != nil {
		return x.AppID
	}
	return ""
}

func (x *GetMyTasksRequest) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

func (x *GetMyTasksRequest) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *GetMyTasksRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type GetMyTasksResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Infos []*UserTask `protobuf:"bytes,10,rep,name=Infos,proto3" json:"Infos,omitempty"`
	Total uint32      `protobuf:"varint,20,opt,name=Total,proto3" json:"Total,omitempty"`
}

func (x *GetMyTasksResponse) Reset() {
	*x = GetMyTasksResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_inspire_gw_v1_task_task_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMyTasksResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMyTasksResponse) ProtoMessage() {}

func (x *GetMyTasksResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_inspire_gw_v1_task_task_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMyTasksResponse.ProtoReflect.Descriptor instead.
func (*GetMyTasksResponse) Descriptor() ([]byte, []int) {
	return file_npool_inspire_gw_v1_task_task_proto_rawDescGZIP(), []int{2}
}

func (x *GetMyTasksResponse) GetInfos() []*UserTask {
	if x != nil {
		return x.Infos
	}
	return nil
}

func (x *GetMyTasksResponse) GetTotal() uint32 {
	if x != nil {
		return x.Total
	}
	return 0
}

type AdminGetTasksRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TargetAppID  string `protobuf:"bytes,10,opt,name=TargetAppID,proto3" json:"TargetAppID,omitempty"`
	TargetUserID string `protobuf:"bytes,20,opt,name=TargetUserID,proto3" json:"TargetUserID,omitempty"`
	Offset       int32  `protobuf:"varint,30,opt,name=Offset,proto3" json:"Offset,omitempty"`
	Limit        int32  `protobuf:"varint,40,opt,name=Limit,proto3" json:"Limit,omitempty"`
}

func (x *AdminGetTasksRequest) Reset() {
	*x = AdminGetTasksRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_inspire_gw_v1_task_task_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdminGetTasksRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdminGetTasksRequest) ProtoMessage() {}

func (x *AdminGetTasksRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_inspire_gw_v1_task_task_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdminGetTasksRequest.ProtoReflect.Descriptor instead.
func (*AdminGetTasksRequest) Descriptor() ([]byte, []int) {
	return file_npool_inspire_gw_v1_task_task_proto_rawDescGZIP(), []int{3}
}

func (x *AdminGetTasksRequest) GetTargetAppID() string {
	if x != nil {
		return x.TargetAppID
	}
	return ""
}

func (x *AdminGetTasksRequest) GetTargetUserID() string {
	if x != nil {
		return x.TargetUserID
	}
	return ""
}

func (x *AdminGetTasksRequest) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *AdminGetTasksRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type AdminGetTasksResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Infos []*UserTask `protobuf:"bytes,10,rep,name=Infos,proto3" json:"Infos,omitempty"`
	Total uint32      `protobuf:"varint,20,opt,name=Total,proto3" json:"Total,omitempty"`
}

func (x *AdminGetTasksResponse) Reset() {
	*x = AdminGetTasksResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_inspire_gw_v1_task_task_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdminGetTasksResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdminGetTasksResponse) ProtoMessage() {}

func (x *AdminGetTasksResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_inspire_gw_v1_task_task_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdminGetTasksResponse.ProtoReflect.Descriptor instead.
func (*AdminGetTasksResponse) Descriptor() ([]byte, []int) {
	return file_npool_inspire_gw_v1_task_task_proto_rawDescGZIP(), []int{4}
}

func (x *AdminGetTasksResponse) GetInfos() []*UserTask {
	if x != nil {
		return x.Infos
	}
	return nil
}

func (x *AdminGetTasksResponse) GetTotal() uint32 {
	if x != nil {
		return x.Total
	}
	return 0
}

var File_npool_inspire_gw_v1_task_task_proto protoreflect.FileDescriptor

var file_npool_inspire_gw_v1_task_task_proto_rawDesc = []byte{
	0x0a, 0x23, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x69, 0x6e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2f,
	0x67, 0x77, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x17, 0x69, 0x6e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x67,
	0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x76, 0x31, 0x1a, 0x1c,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x26, 0x6e, 0x70,
	0x6f, 0x6f, 0x6c, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x69, 0x6e,
	0x73, 0x70, 0x69, 0x72, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xaf, 0x05, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x54, 0x61, 0x73,
	0x6b, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x49,
	0x44, 0x12, 0x14, 0x0a, 0x05, 0x45, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x45, 0x6e, 0x74, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x41, 0x70, 0x70, 0x49, 0x44,
	0x18, 0x1e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x41, 0x70, 0x70, 0x49, 0x44, 0x12, 0x18, 0x0a,
	0x07, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x28, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x49, 0x44, 0x12, 0x3a, 0x0a, 0x08, 0x54, 0x61, 0x73, 0x6b, 0x54,
	0x79, 0x70, 0x65, 0x18, 0x32, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1e, 0x2e, 0x62, 0x61, 0x73, 0x65,
	0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x69, 0x6e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x54, 0x61, 0x73, 0x6b, 0x54, 0x79, 0x70, 0x65, 0x52, 0x08, 0x54, 0x61, 0x73, 0x6b, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x3c, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x54, 0x61, 0x73, 0x6b, 0x44,
	0x65, 0x73, 0x63, 0x18, 0x46, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x54, 0x61, 0x73, 0x6b, 0x44,
	0x65, 0x73, 0x63, 0x12, 0x1c, 0x0a, 0x09, 0x53, 0x74, 0x65, 0x70, 0x47, 0x75, 0x69, 0x64, 0x65,
	0x18, 0x50, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x53, 0x74, 0x65, 0x70, 0x47, 0x75, 0x69, 0x64,
	0x65, 0x12, 0x2a, 0x0a, 0x10, 0x52, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x5a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x52, 0x65, 0x63,
	0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x64, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x49, 0x6e,
	0x64, 0x65, 0x78, 0x12, 0x1e, 0x0a, 0x0a, 0x4c, 0x61, 0x73, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x49,
	0x44, 0x18, 0x6e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x4c, 0x61, 0x73, 0x74, 0x54, 0x61, 0x73,
	0x6b, 0x49, 0x44, 0x12, 0x26, 0x0a, 0x0e, 0x4d, 0x61, 0x78, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x78, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0e, 0x4d, 0x61, 0x78,
	0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x27, 0x0a, 0x0e, 0x43,
	0x6f, 0x6f, 0x6c, 0x64, 0x6f, 0x77, 0x6e, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x18, 0x82, 0x01,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x0e, 0x43, 0x6f, 0x6f, 0x6c, 0x64, 0x6f, 0x77, 0x6e, 0x53, 0x65,
	0x63, 0x6f, 0x6e, 0x64, 0x12, 0x29, 0x0a, 0x0f, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x69,
	0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x18, 0x8c, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0f,
	0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x12,
	0x21, 0x0a, 0x0b, 0x4e, 0x65, 0x78, 0x74, 0x53, 0x74, 0x61, 0x72, 0x74, 0x41, 0x74, 0x18, 0x96,
	0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x4e, 0x65, 0x78, 0x74, 0x53, 0x74, 0x61, 0x72, 0x74,
	0x41, 0x74, 0x12, 0x3e, 0x0a, 0x09, 0x54, 0x61, 0x73, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x65, 0x18,
	0xa0, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1f, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70,
	0x65, 0x73, 0x2e, 0x69, 0x6e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x61,
	0x73, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x09, 0x54, 0x61, 0x73, 0x6b, 0x53, 0x74, 0x61,
	0x74, 0x65, 0x12, 0x44, 0x0a, 0x0b, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x53, 0x74, 0x61, 0x74,
	0x65, 0x18, 0xaa, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x21, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x2e, 0x69, 0x6e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x0b, 0x52, 0x65, 0x77,
	0x61, 0x72, 0x64, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x1d, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0xe8, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x18, 0xf2, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x6f, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x4d, 0x79, 0x54,
	0x61, 0x73, 0x6b, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x41,
	0x70, 0x70, 0x49, 0x44, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x41, 0x70, 0x70, 0x49,
	0x44, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x14, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x4f, 0x66, 0x66,
	0x73, 0x65, 0x74, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x4f, 0x66, 0x66, 0x73, 0x65,
	0x74, 0x12, 0x14, 0x0a, 0x05, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x28, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x05, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x22, 0x63, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x4d, 0x79,
	0x54, 0x61, 0x73, 0x6b, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x37, 0x0a,
	0x05, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x69,
	0x6e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x74,
	0x61, 0x73, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x54, 0x61, 0x73, 0x6b, 0x52,
	0x05, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x18,
	0x14, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x22, 0x8a, 0x01, 0x0a,
	0x14, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x47, 0x65, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x41,
	0x70, 0x70, 0x49, 0x44, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x54, 0x61, 0x72, 0x67,
	0x65, 0x74, 0x41, 0x70, 0x70, 0x49, 0x44, 0x12, 0x22, 0x0a, 0x0c, 0x54, 0x61, 0x72, 0x67, 0x65,
	0x74, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x54,
	0x61, 0x72, 0x67, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x4f,
	0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x4f, 0x66, 0x66,
	0x73, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x28, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x05, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x22, 0x66, 0x0a, 0x15, 0x41, 0x64, 0x6d,
	0x69, 0x6e, 0x47, 0x65, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x37, 0x0a, 0x05, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x21, 0x2e, 0x69, 0x6e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x67, 0x61, 0x74, 0x65,
	0x77, 0x61, 0x79, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72,
	0x54, 0x61, 0x73, 0x6b, 0x52, 0x05, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x54,
	0x6f, 0x74, 0x61, 0x6c, 0x18, 0x14, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x54, 0x6f, 0x74, 0x61,
	0x6c, 0x32, 0x9f, 0x02, 0x0a, 0x07, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x12, 0x8e, 0x01,
	0x0a, 0x0d, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x47, 0x65, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x12,
	0x2d, 0x2e, 0x69, 0x6e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61,
	0x79, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x47,
	0x65, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2e,
	0x2e, 0x69, 0x6e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79,
	0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x47, 0x65,
	0x74, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1e,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x18, 0x3a, 0x01, 0x2a, 0x22, 0x13, 0x2f, 0x76, 0x31, 0x2f, 0x61,
	0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x67, 0x65, 0x74, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x12, 0x82,
	0x01, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x4d, 0x79, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x12, 0x2a, 0x2e,
	0x69, 0x6e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e,
	0x74, 0x61, 0x73, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x79, 0x54, 0x61, 0x73,
	0x6b, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2b, 0x2e, 0x69, 0x6e, 0x73, 0x70,
	0x69, 0x72, 0x65, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x74, 0x61, 0x73, 0x6b,
	0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x79, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15, 0x3a, 0x01,
	0x2a, 0x22, 0x10, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x65, 0x74, 0x2f, 0x6d, 0x79, 0x2f, 0x74, 0x61,
	0x73, 0x6b, 0x73, 0x42, 0x3b, 0x5a, 0x39, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x4e, 0x70, 0x6f, 0x6f, 0x6c, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x69, 0x6e,
	0x73, 0x70, 0x69, 0x72, 0x65, 0x2f, 0x67, 0x77, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x61, 0x73, 0x6b,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_npool_inspire_gw_v1_task_task_proto_rawDescOnce sync.Once
	file_npool_inspire_gw_v1_task_task_proto_rawDescData = file_npool_inspire_gw_v1_task_task_proto_rawDesc
)

func file_npool_inspire_gw_v1_task_task_proto_rawDescGZIP() []byte {
	file_npool_inspire_gw_v1_task_task_proto_rawDescOnce.Do(func() {
		file_npool_inspire_gw_v1_task_task_proto_rawDescData = protoimpl.X.CompressGZIP(file_npool_inspire_gw_v1_task_task_proto_rawDescData)
	})
	return file_npool_inspire_gw_v1_task_task_proto_rawDescData
}

var file_npool_inspire_gw_v1_task_task_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_npool_inspire_gw_v1_task_task_proto_goTypes = []interface{}{
	(*UserTask)(nil),              // 0: inspire.gateway.task.v1.UserTask
	(*GetMyTasksRequest)(nil),     // 1: inspire.gateway.task.v1.GetMyTasksRequest
	(*GetMyTasksResponse)(nil),    // 2: inspire.gateway.task.v1.GetMyTasksResponse
	(*AdminGetTasksRequest)(nil),  // 3: inspire.gateway.task.v1.AdminGetTasksRequest
	(*AdminGetTasksResponse)(nil), // 4: inspire.gateway.task.v1.AdminGetTasksResponse
	(v1.TaskType)(0),              // 5: basetypes.inspire.v1.TaskType
	(v1.TaskState)(0),             // 6: basetypes.inspire.v1.TaskState
	(v1.RewardState)(0),           // 7: basetypes.inspire.v1.RewardState
}
var file_npool_inspire_gw_v1_task_task_proto_depIdxs = []int32{
	5, // 0: inspire.gateway.task.v1.UserTask.TaskType:type_name -> basetypes.inspire.v1.TaskType
	6, // 1: inspire.gateway.task.v1.UserTask.TaskState:type_name -> basetypes.inspire.v1.TaskState
	7, // 2: inspire.gateway.task.v1.UserTask.RewardState:type_name -> basetypes.inspire.v1.RewardState
	0, // 3: inspire.gateway.task.v1.GetMyTasksResponse.Infos:type_name -> inspire.gateway.task.v1.UserTask
	0, // 4: inspire.gateway.task.v1.AdminGetTasksResponse.Infos:type_name -> inspire.gateway.task.v1.UserTask
	3, // 5: inspire.gateway.task.v1.Gateway.AdminGetTasks:input_type -> inspire.gateway.task.v1.AdminGetTasksRequest
	1, // 6: inspire.gateway.task.v1.Gateway.GetMyTasks:input_type -> inspire.gateway.task.v1.GetMyTasksRequest
	4, // 7: inspire.gateway.task.v1.Gateway.AdminGetTasks:output_type -> inspire.gateway.task.v1.AdminGetTasksResponse
	2, // 8: inspire.gateway.task.v1.Gateway.GetMyTasks:output_type -> inspire.gateway.task.v1.GetMyTasksResponse
	7, // [7:9] is the sub-list for method output_type
	5, // [5:7] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_npool_inspire_gw_v1_task_task_proto_init() }
func file_npool_inspire_gw_v1_task_task_proto_init() {
	if File_npool_inspire_gw_v1_task_task_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_npool_inspire_gw_v1_task_task_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserTask); i {
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
		file_npool_inspire_gw_v1_task_task_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMyTasksRequest); i {
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
		file_npool_inspire_gw_v1_task_task_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMyTasksResponse); i {
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
		file_npool_inspire_gw_v1_task_task_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdminGetTasksRequest); i {
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
		file_npool_inspire_gw_v1_task_task_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdminGetTasksResponse); i {
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
			RawDescriptor: file_npool_inspire_gw_v1_task_task_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_npool_inspire_gw_v1_task_task_proto_goTypes,
		DependencyIndexes: file_npool_inspire_gw_v1_task_task_proto_depIdxs,
		MessageInfos:      file_npool_inspire_gw_v1_task_task_proto_msgTypes,
	}.Build()
	File_npool_inspire_gw_v1_task_task_proto = out.File
	file_npool_inspire_gw_v1_task_task_proto_rawDesc = nil
	file_npool_inspire_gw_v1_task_task_proto_goTypes = nil
	file_npool_inspire_gw_v1_task_task_proto_depIdxs = nil
}
