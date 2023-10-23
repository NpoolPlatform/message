// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.18.1
// source: npool/notif/gw/v1/notif/user/user.proto

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

type NotifUser struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID           uint32       `protobuf:"varint,9,opt,name=ID,proto3" json:"ID,omitempty"`
	EntID        string       `protobuf:"bytes,10,opt,name=EntID,proto3" json:"EntID,omitempty"`
	EventType    v1.UsedFor   `protobuf:"varint,20,opt,name=EventType,proto3,enum=basetypes.v1.UsedFor" json:"EventType,omitempty"`
	AppID        string       `protobuf:"bytes,30,opt,name=AppID,proto3" json:"AppID,omitempty"`
	UserID       string       `protobuf:"bytes,50,opt,name=UserID,proto3" json:"UserID,omitempty"`
	EmailAddress string       `protobuf:"bytes,60,opt,name=EmailAddress,proto3" json:"EmailAddress,omitempty"`
	PhoneNO      string       `protobuf:"bytes,70,opt,name=PhoneNO,proto3" json:"PhoneNO,omitempty"`
	Username     string       `protobuf:"bytes,80,opt,name=Username,proto3" json:"Username,omitempty"`
	Title        string       `protobuf:"bytes,90,opt,name=Title,proto3" json:"Title,omitempty"`
	Content      string       `protobuf:"bytes,100,opt,name=Content,proto3" json:"Content,omitempty"`
	NotifType    v1.NotifType `protobuf:"varint,110,opt,name=NotifType,proto3,enum=basetypes.v1.NotifType" json:"NotifType,omitempty"`
	CreatedAt    uint32       `protobuf:"varint,120,opt,name=CreatedAt,proto3" json:"CreatedAt,omitempty"`
	UpdatedAt    uint32       `protobuf:"varint,130,opt,name=UpdatedAt,proto3" json:"UpdatedAt,omitempty"`
}

func (x *NotifUser) Reset() {
	*x = NotifUser{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_notif_gw_v1_notif_user_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NotifUser) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NotifUser) ProtoMessage() {}

func (x *NotifUser) ProtoReflect() protoreflect.Message {
	mi := &file_npool_notif_gw_v1_notif_user_user_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NotifUser.ProtoReflect.Descriptor instead.
func (*NotifUser) Descriptor() ([]byte, []int) {
	return file_npool_notif_gw_v1_notif_user_user_proto_rawDescGZIP(), []int{0}
}

func (x *NotifUser) GetID() uint32 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *NotifUser) GetEntID() string {
	if x != nil {
		return x.EntID
	}
	return ""
}

func (x *NotifUser) GetEventType() v1.UsedFor {
	if x != nil {
		return x.EventType
	}
	return v1.UsedFor(0)
}

func (x *NotifUser) GetAppID() string {
	if x != nil {
		return x.AppID
	}
	return ""
}

func (x *NotifUser) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

func (x *NotifUser) GetEmailAddress() string {
	if x != nil {
		return x.EmailAddress
	}
	return ""
}

func (x *NotifUser) GetPhoneNO() string {
	if x != nil {
		return x.PhoneNO
	}
	return ""
}

func (x *NotifUser) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *NotifUser) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *NotifUser) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *NotifUser) GetNotifType() v1.NotifType {
	if x != nil {
		return x.NotifType
	}
	return v1.NotifType(0)
}

func (x *NotifUser) GetCreatedAt() uint32 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *NotifUser) GetUpdatedAt() uint32 {
	if x != nil {
		return x.UpdatedAt
	}
	return 0
}

type CreateNotifUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AppID        string     `protobuf:"bytes,10,opt,name=AppID,proto3" json:"AppID,omitempty"`
	TargetUserID string     `protobuf:"bytes,20,opt,name=TargetUserID,proto3" json:"TargetUserID,omitempty"`
	EventType    v1.UsedFor `protobuf:"varint,30,opt,name=EventType,proto3,enum=basetypes.v1.UsedFor" json:"EventType,omitempty"`
}

func (x *CreateNotifUserRequest) Reset() {
	*x = CreateNotifUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_notif_gw_v1_notif_user_user_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateNotifUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateNotifUserRequest) ProtoMessage() {}

func (x *CreateNotifUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_notif_gw_v1_notif_user_user_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateNotifUserRequest.ProtoReflect.Descriptor instead.
func (*CreateNotifUserRequest) Descriptor() ([]byte, []int) {
	return file_npool_notif_gw_v1_notif_user_user_proto_rawDescGZIP(), []int{1}
}

func (x *CreateNotifUserRequest) GetAppID() string {
	if x != nil {
		return x.AppID
	}
	return ""
}

func (x *CreateNotifUserRequest) GetTargetUserID() string {
	if x != nil {
		return x.TargetUserID
	}
	return ""
}

func (x *CreateNotifUserRequest) GetEventType() v1.UsedFor {
	if x != nil {
		return x.EventType
	}
	return v1.UsedFor(0)
}

type CreateNotifUserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *NotifUser `protobuf:"bytes,10,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (x *CreateNotifUserResponse) Reset() {
	*x = CreateNotifUserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_notif_gw_v1_notif_user_user_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateNotifUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateNotifUserResponse) ProtoMessage() {}

func (x *CreateNotifUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_notif_gw_v1_notif_user_user_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateNotifUserResponse.ProtoReflect.Descriptor instead.
func (*CreateNotifUserResponse) Descriptor() ([]byte, []int) {
	return file_npool_notif_gw_v1_notif_user_user_proto_rawDescGZIP(), []int{2}
}

func (x *CreateNotifUserResponse) GetInfo() *NotifUser {
	if x != nil {
		return x.Info
	}
	return nil
}

type DeleteNotifUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID    uint32 `protobuf:"varint,10,opt,name=ID,proto3" json:"ID,omitempty"`
	AppID string `protobuf:"bytes,20,opt,name=AppID,proto3" json:"AppID,omitempty"`
}

func (x *DeleteNotifUserRequest) Reset() {
	*x = DeleteNotifUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_notif_gw_v1_notif_user_user_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteNotifUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteNotifUserRequest) ProtoMessage() {}

func (x *DeleteNotifUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_notif_gw_v1_notif_user_user_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteNotifUserRequest.ProtoReflect.Descriptor instead.
func (*DeleteNotifUserRequest) Descriptor() ([]byte, []int) {
	return file_npool_notif_gw_v1_notif_user_user_proto_rawDescGZIP(), []int{3}
}

func (x *DeleteNotifUserRequest) GetID() uint32 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *DeleteNotifUserRequest) GetAppID() string {
	if x != nil {
		return x.AppID
	}
	return ""
}

type DeleteNotifUserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *NotifUser `protobuf:"bytes,10,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (x *DeleteNotifUserResponse) Reset() {
	*x = DeleteNotifUserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_notif_gw_v1_notif_user_user_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteNotifUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteNotifUserResponse) ProtoMessage() {}

func (x *DeleteNotifUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_notif_gw_v1_notif_user_user_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteNotifUserResponse.ProtoReflect.Descriptor instead.
func (*DeleteNotifUserResponse) Descriptor() ([]byte, []int) {
	return file_npool_notif_gw_v1_notif_user_user_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteNotifUserResponse) GetInfo() *NotifUser {
	if x != nil {
		return x.Info
	}
	return nil
}

type GetNotifUsersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AppID     string      `protobuf:"bytes,10,opt,name=AppID,proto3" json:"AppID,omitempty"`
	EventType *v1.UsedFor `protobuf:"varint,20,opt,name=EventType,proto3,enum=basetypes.v1.UsedFor,oneof" json:"EventType,omitempty"`
	Offset    int32       `protobuf:"varint,30,opt,name=Offset,proto3" json:"Offset,omitempty"`
	Limit     int32       `protobuf:"varint,40,opt,name=Limit,proto3" json:"Limit,omitempty"`
}

func (x *GetNotifUsersRequest) Reset() {
	*x = GetNotifUsersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_notif_gw_v1_notif_user_user_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetNotifUsersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetNotifUsersRequest) ProtoMessage() {}

func (x *GetNotifUsersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_notif_gw_v1_notif_user_user_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetNotifUsersRequest.ProtoReflect.Descriptor instead.
func (*GetNotifUsersRequest) Descriptor() ([]byte, []int) {
	return file_npool_notif_gw_v1_notif_user_user_proto_rawDescGZIP(), []int{5}
}

func (x *GetNotifUsersRequest) GetAppID() string {
	if x != nil {
		return x.AppID
	}
	return ""
}

func (x *GetNotifUsersRequest) GetEventType() v1.UsedFor {
	if x != nil && x.EventType != nil {
		return *x.EventType
	}
	return v1.UsedFor(0)
}

func (x *GetNotifUsersRequest) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *GetNotifUsersRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type GetNotifUsersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Infos []*NotifUser `protobuf:"bytes,10,rep,name=Infos,proto3" json:"Infos,omitempty"`
	Total uint32       `protobuf:"varint,20,opt,name=Total,proto3" json:"Total,omitempty"`
}

func (x *GetNotifUsersResponse) Reset() {
	*x = GetNotifUsersResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_notif_gw_v1_notif_user_user_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetNotifUsersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetNotifUsersResponse) ProtoMessage() {}

func (x *GetNotifUsersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_notif_gw_v1_notif_user_user_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetNotifUsersResponse.ProtoReflect.Descriptor instead.
func (*GetNotifUsersResponse) Descriptor() ([]byte, []int) {
	return file_npool_notif_gw_v1_notif_user_user_proto_rawDescGZIP(), []int{6}
}

func (x *GetNotifUsersResponse) GetInfos() []*NotifUser {
	if x != nil {
		return x.Infos
	}
	return nil
}

func (x *GetNotifUsersResponse) GetTotal() uint32 {
	if x != nil {
		return x.Total
	}
	return 0
}

type GetAppNotifUsersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TargetAppID string `protobuf:"bytes,10,opt,name=TargetAppID,proto3" json:"TargetAppID,omitempty"`
	Offset      int32  `protobuf:"varint,20,opt,name=Offset,proto3" json:"Offset,omitempty"`
	Limit       int32  `protobuf:"varint,30,opt,name=Limit,proto3" json:"Limit,omitempty"`
}

func (x *GetAppNotifUsersRequest) Reset() {
	*x = GetAppNotifUsersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_notif_gw_v1_notif_user_user_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAppNotifUsersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAppNotifUsersRequest) ProtoMessage() {}

func (x *GetAppNotifUsersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_notif_gw_v1_notif_user_user_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAppNotifUsersRequest.ProtoReflect.Descriptor instead.
func (*GetAppNotifUsersRequest) Descriptor() ([]byte, []int) {
	return file_npool_notif_gw_v1_notif_user_user_proto_rawDescGZIP(), []int{7}
}

func (x *GetAppNotifUsersRequest) GetTargetAppID() string {
	if x != nil {
		return x.TargetAppID
	}
	return ""
}

func (x *GetAppNotifUsersRequest) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *GetAppNotifUsersRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type GetAppNotifUsersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Infos []*NotifUser `protobuf:"bytes,10,rep,name=Infos,proto3" json:"Infos,omitempty"`
	Total uint32       `protobuf:"varint,20,opt,name=Total,proto3" json:"Total,omitempty"`
}

func (x *GetAppNotifUsersResponse) Reset() {
	*x = GetAppNotifUsersResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_notif_gw_v1_notif_user_user_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAppNotifUsersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAppNotifUsersResponse) ProtoMessage() {}

func (x *GetAppNotifUsersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_notif_gw_v1_notif_user_user_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAppNotifUsersResponse.ProtoReflect.Descriptor instead.
func (*GetAppNotifUsersResponse) Descriptor() ([]byte, []int) {
	return file_npool_notif_gw_v1_notif_user_user_proto_rawDescGZIP(), []int{8}
}

func (x *GetAppNotifUsersResponse) GetInfos() []*NotifUser {
	if x != nil {
		return x.Infos
	}
	return nil
}

func (x *GetAppNotifUsersResponse) GetTotal() uint32 {
	if x != nil {
		return x.Total
	}
	return 0
}

var File_npool_notif_gw_v1_notif_user_user_proto protoreflect.FileDescriptor

var file_npool_notif_gw_v1_notif_user_user_proto_rawDesc = []byte{
	0x0a, 0x27, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x2f, 0x67, 0x77,
	0x2f, 0x76, 0x31, 0x2f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x75,
	0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1b, 0x6e, 0x6f, 0x74, 0x69, 0x66,
	0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x2e, 0x75,
	0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x62, 0x61, 0x73, 0x65,
	0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x64, 0x66, 0x6f, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x22, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x62, 0x61,
	0x73, 0x65, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x6e, 0x6f, 0x74, 0x69, 0x66,
	0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x92, 0x03, 0x0a, 0x09, 0x4e,
	0x6f, 0x74, 0x69, 0x66, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x45, 0x6e, 0x74, 0x49,
	0x44, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x45, 0x6e, 0x74, 0x49, 0x44, 0x12, 0x33,
	0x0a, 0x09, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x18, 0x14, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x15, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x55, 0x73, 0x65, 0x64, 0x46, 0x6f, 0x72, 0x52, 0x09, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x41, 0x70, 0x70, 0x49, 0x44, 0x18, 0x1e, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x41, 0x70, 0x70, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65,
	0x72, 0x49, 0x44, 0x18, 0x32, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49,
	0x44, 0x12, 0x22, 0x0a, 0x0c, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x18, 0x3c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x41, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x4f,
	0x18, 0x46, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x4f, 0x12,
	0x1a, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x50, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x54,
	0x69, 0x74, 0x6c, 0x65, 0x18, 0x5a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x54, 0x69, 0x74, 0x6c,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x64, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x35, 0x0a, 0x09, 0x4e,
	0x6f, 0x74, 0x69, 0x66, 0x54, 0x79, 0x70, 0x65, 0x18, 0x6e, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x17,
	0x2e, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x6f,
	0x74, 0x69, 0x66, 0x54, 0x79, 0x70, 0x65, 0x52, 0x09, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18,
	0x78, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x12, 0x1d, 0x0a, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x82, 0x01,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22,
	0x87, 0x01, 0x0a, 0x16, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x55,
	0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x41, 0x70,
	0x70, 0x49, 0x44, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x41, 0x70, 0x70, 0x49, 0x44,
	0x12, 0x22, 0x0a, 0x0c, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44,
	0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x55, 0x73,
	0x65, 0x72, 0x49, 0x44, 0x12, 0x33, 0x0a, 0x09, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70,
	0x65, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x15, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x64, 0x46, 0x6f, 0x72, 0x52, 0x09,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x22, 0x55, 0x0a, 0x17, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3a, 0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x26, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77,
	0x61, 0x79, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31,
	0x2e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x49, 0x6e, 0x66, 0x6f,
	0x22, 0x3e, 0x0a, 0x16, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x55,
	0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x41, 0x70,
	0x70, 0x49, 0x44, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x41, 0x70, 0x70, 0x49, 0x44,
	0x22, 0x55, 0x0a, 0x17, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x55,
	0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3a, 0x0a, 0x04, 0x49,
	0x6e, 0x66, 0x6f, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x6e, 0x6f, 0x74, 0x69,
	0x66, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x55, 0x73, 0x65,
	0x72, 0x52, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0xa2, 0x01, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x4e,
	0x6f, 0x74, 0x69, 0x66, 0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x14, 0x0a, 0x05, 0x41, 0x70, 0x70, 0x49, 0x44, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x41, 0x70, 0x70, 0x49, 0x44, 0x12, 0x38, 0x0a, 0x09, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54,
	0x79, 0x70, 0x65, 0x18, 0x14, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x15, 0x2e, 0x62, 0x61, 0x73, 0x65,
	0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x64, 0x46, 0x6f, 0x72,
	0x48, 0x00, 0x52, 0x09, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x88, 0x01, 0x01,
	0x12, 0x16, 0x0a, 0x06, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x06, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x4c, 0x69, 0x6d, 0x69,
	0x74, 0x18, 0x28, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x42, 0x0c,
	0x0a, 0x0a, 0x5f, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x22, 0x6b, 0x0a, 0x15,
	0x47, 0x65, 0x74, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3c, 0x0a, 0x05, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x18, 0x0a,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x2e, 0x67, 0x61, 0x74,
	0x65, 0x77, 0x61, 0x79, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e,
	0x76, 0x31, 0x2e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x55, 0x73, 0x65, 0x72, 0x52, 0x05, 0x49, 0x6e,
	0x66, 0x6f, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x14, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x22, 0x69, 0x0a, 0x17, 0x47, 0x65, 0x74,
	0x41, 0x70, 0x70, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x41, 0x70,
	0x70, 0x49, 0x44, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x54, 0x61, 0x72, 0x67, 0x65,
	0x74, 0x41, 0x70, 0x70, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74,
	0x18, 0x14, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x14,
	0x0a, 0x05, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x4c,
	0x69, 0x6d, 0x69, 0x74, 0x22, 0x6e, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x41, 0x70, 0x70, 0x4e, 0x6f,
	0x74, 0x69, 0x66, 0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x3c, 0x0a, 0x05, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x26, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e,
	0x6e, 0x6f, 0x74, 0x69, 0x66, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x6f,
	0x74, 0x69, 0x66, 0x55, 0x73, 0x65, 0x72, 0x52, 0x05, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x12, 0x14,
	0x0a, 0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x14, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x54,
	0x6f, 0x74, 0x61, 0x6c, 0x32, 0x86, 0x05, 0x0a, 0x07, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79,
	0x12, 0x9d, 0x01, 0x0a, 0x0f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4e, 0x6f, 0x74, 0x69, 0x66,
	0x55, 0x73, 0x65, 0x72, 0x12, 0x33, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x2e, 0x67, 0x61, 0x74,
	0x65, 0x77, 0x61, 0x79, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e,
	0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x55, 0x73,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x34, 0x2e, 0x6e, 0x6f, 0x74, 0x69,
	0x66, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4e, 0x6f,
	0x74, 0x69, 0x66, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x1f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x19, 0x22, 0x14, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x2f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x75, 0x73, 0x65, 0x72, 0x3a, 0x01, 0x2a,
	0x12, 0x9d, 0x01, 0x0a, 0x0f, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4e, 0x6f, 0x74, 0x69, 0x66,
	0x55, 0x73, 0x65, 0x72, 0x12, 0x33, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x2e, 0x67, 0x61, 0x74,
	0x65, 0x77, 0x61, 0x79, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e,
	0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x55, 0x73,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x34, 0x2e, 0x6e, 0x6f, 0x74, 0x69,
	0x66, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4e, 0x6f,
	0x74, 0x69, 0x66, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x1f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x19, 0x22, 0x14, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x2f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x75, 0x73, 0x65, 0x72, 0x3a, 0x01, 0x2a,
	0x12, 0x95, 0x01, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x55, 0x73, 0x65,
	0x72, 0x73, 0x12, 0x31, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77,
	0x61, 0x79, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31,
	0x2e, 0x47, 0x65, 0x74, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x32, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x2e, 0x67, 0x61,
	0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x55, 0x73, 0x65, 0x72,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1d, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x17, 0x22, 0x12, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x65, 0x74, 0x2f, 0x6e, 0x6f, 0x74, 0x69, 0x66,
	0x75, 0x73, 0x65, 0x72, 0x73, 0x3a, 0x01, 0x2a, 0x12, 0xa2, 0x01, 0x0a, 0x10, 0x47, 0x65, 0x74,
	0x41, 0x70, 0x70, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x55, 0x73, 0x65, 0x72, 0x73, 0x12, 0x34, 0x2e,
	0x6e, 0x6f, 0x74, 0x69, 0x66, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x6e, 0x6f,
	0x74, 0x69, 0x66, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41,
	0x70, 0x70, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x35, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x2e, 0x67, 0x61, 0x74, 0x65,
	0x77, 0x61, 0x79, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76,
	0x31, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x70, 0x70, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x55, 0x73, 0x65,
	0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x21, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x1b, 0x22, 0x16, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x65, 0x74, 0x2f, 0x61, 0x70, 0x70, 0x2f,
	0x6e, 0x6f, 0x74, 0x69, 0x66, 0x75, 0x73, 0x65, 0x72, 0x73, 0x3a, 0x01, 0x2a, 0x42, 0x3f, 0x5a,
	0x3d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4e, 0x70, 0x6f, 0x6f,
	0x6c, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x2f, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x2f, 0x67, 0x77,
	0x2f, 0x76, 0x31, 0x2f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_npool_notif_gw_v1_notif_user_user_proto_rawDescOnce sync.Once
	file_npool_notif_gw_v1_notif_user_user_proto_rawDescData = file_npool_notif_gw_v1_notif_user_user_proto_rawDesc
)

func file_npool_notif_gw_v1_notif_user_user_proto_rawDescGZIP() []byte {
	file_npool_notif_gw_v1_notif_user_user_proto_rawDescOnce.Do(func() {
		file_npool_notif_gw_v1_notif_user_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_npool_notif_gw_v1_notif_user_user_proto_rawDescData)
	})
	return file_npool_notif_gw_v1_notif_user_user_proto_rawDescData
}

var file_npool_notif_gw_v1_notif_user_user_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_npool_notif_gw_v1_notif_user_user_proto_goTypes = []interface{}{
	(*NotifUser)(nil),                // 0: notif.gateway.notif.user.v1.NotifUser
	(*CreateNotifUserRequest)(nil),   // 1: notif.gateway.notif.user.v1.CreateNotifUserRequest
	(*CreateNotifUserResponse)(nil),  // 2: notif.gateway.notif.user.v1.CreateNotifUserResponse
	(*DeleteNotifUserRequest)(nil),   // 3: notif.gateway.notif.user.v1.DeleteNotifUserRequest
	(*DeleteNotifUserResponse)(nil),  // 4: notif.gateway.notif.user.v1.DeleteNotifUserResponse
	(*GetNotifUsersRequest)(nil),     // 5: notif.gateway.notif.user.v1.GetNotifUsersRequest
	(*GetNotifUsersResponse)(nil),    // 6: notif.gateway.notif.user.v1.GetNotifUsersResponse
	(*GetAppNotifUsersRequest)(nil),  // 7: notif.gateway.notif.user.v1.GetAppNotifUsersRequest
	(*GetAppNotifUsersResponse)(nil), // 8: notif.gateway.notif.user.v1.GetAppNotifUsersResponse
	(v1.UsedFor)(0),                  // 9: basetypes.v1.UsedFor
	(v1.NotifType)(0),                // 10: basetypes.v1.NotifType
}
var file_npool_notif_gw_v1_notif_user_user_proto_depIdxs = []int32{
	9,  // 0: notif.gateway.notif.user.v1.NotifUser.EventType:type_name -> basetypes.v1.UsedFor
	10, // 1: notif.gateway.notif.user.v1.NotifUser.NotifType:type_name -> basetypes.v1.NotifType
	9,  // 2: notif.gateway.notif.user.v1.CreateNotifUserRequest.EventType:type_name -> basetypes.v1.UsedFor
	0,  // 3: notif.gateway.notif.user.v1.CreateNotifUserResponse.Info:type_name -> notif.gateway.notif.user.v1.NotifUser
	0,  // 4: notif.gateway.notif.user.v1.DeleteNotifUserResponse.Info:type_name -> notif.gateway.notif.user.v1.NotifUser
	9,  // 5: notif.gateway.notif.user.v1.GetNotifUsersRequest.EventType:type_name -> basetypes.v1.UsedFor
	0,  // 6: notif.gateway.notif.user.v1.GetNotifUsersResponse.Infos:type_name -> notif.gateway.notif.user.v1.NotifUser
	0,  // 7: notif.gateway.notif.user.v1.GetAppNotifUsersResponse.Infos:type_name -> notif.gateway.notif.user.v1.NotifUser
	1,  // 8: notif.gateway.notif.user.v1.Gateway.CreateNotifUser:input_type -> notif.gateway.notif.user.v1.CreateNotifUserRequest
	3,  // 9: notif.gateway.notif.user.v1.Gateway.DeleteNotifUser:input_type -> notif.gateway.notif.user.v1.DeleteNotifUserRequest
	5,  // 10: notif.gateway.notif.user.v1.Gateway.GetNotifUsers:input_type -> notif.gateway.notif.user.v1.GetNotifUsersRequest
	7,  // 11: notif.gateway.notif.user.v1.Gateway.GetAppNotifUsers:input_type -> notif.gateway.notif.user.v1.GetAppNotifUsersRequest
	2,  // 12: notif.gateway.notif.user.v1.Gateway.CreateNotifUser:output_type -> notif.gateway.notif.user.v1.CreateNotifUserResponse
	4,  // 13: notif.gateway.notif.user.v1.Gateway.DeleteNotifUser:output_type -> notif.gateway.notif.user.v1.DeleteNotifUserResponse
	6,  // 14: notif.gateway.notif.user.v1.Gateway.GetNotifUsers:output_type -> notif.gateway.notif.user.v1.GetNotifUsersResponse
	8,  // 15: notif.gateway.notif.user.v1.Gateway.GetAppNotifUsers:output_type -> notif.gateway.notif.user.v1.GetAppNotifUsersResponse
	12, // [12:16] is the sub-list for method output_type
	8,  // [8:12] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_npool_notif_gw_v1_notif_user_user_proto_init() }
func file_npool_notif_gw_v1_notif_user_user_proto_init() {
	if File_npool_notif_gw_v1_notif_user_user_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_npool_notif_gw_v1_notif_user_user_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NotifUser); i {
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
		file_npool_notif_gw_v1_notif_user_user_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateNotifUserRequest); i {
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
		file_npool_notif_gw_v1_notif_user_user_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateNotifUserResponse); i {
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
		file_npool_notif_gw_v1_notif_user_user_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteNotifUserRequest); i {
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
		file_npool_notif_gw_v1_notif_user_user_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteNotifUserResponse); i {
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
		file_npool_notif_gw_v1_notif_user_user_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetNotifUsersRequest); i {
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
		file_npool_notif_gw_v1_notif_user_user_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetNotifUsersResponse); i {
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
		file_npool_notif_gw_v1_notif_user_user_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAppNotifUsersRequest); i {
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
		file_npool_notif_gw_v1_notif_user_user_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAppNotifUsersResponse); i {
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
	file_npool_notif_gw_v1_notif_user_user_proto_msgTypes[5].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_npool_notif_gw_v1_notif_user_user_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_npool_notif_gw_v1_notif_user_user_proto_goTypes,
		DependencyIndexes: file_npool_notif_gw_v1_notif_user_user_proto_depIdxs,
		MessageInfos:      file_npool_notif_gw_v1_notif_user_user_proto_msgTypes,
	}.Build()
	File_npool_notif_gw_v1_notif_user_user_proto = out.File
	file_npool_notif_gw_v1_notif_user_user_proto_rawDesc = nil
	file_npool_notif_gw_v1_notif_user_user_proto_goTypes = nil
	file_npool_notif_gw_v1_notif_user_user_proto_depIdxs = nil
}
