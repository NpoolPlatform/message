// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.18.1
// source: npool/appuser/gw/v1/user/user.proto

package user

import (
	_ "github.com/NpoolPlatform/message/npool/appuser/mgr/v2/appuser"
	user "github.com/NpoolPlatform/message/npool/appuser/mw/v1/user"
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

type SignupRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Permissioned by path
	AppID            string `protobuf:"bytes,10,opt,name=AppID,proto3" json:"AppID,omitempty"`
	Username         string `protobuf:"bytes,20,opt,name=Username,proto3" json:"Username,omitempty"`
	PasswordHash     string `protobuf:"bytes,30,opt,name=PasswordHash,proto3" json:"PasswordHash,omitempty"`
	Account          string `protobuf:"bytes,40,opt,name=Account,proto3" json:"Account,omitempty"`
	AccountType      string `protobuf:"bytes,50,opt,name=AccountType,proto3" json:"AccountType,omitempty"`
	VerificationCode string `protobuf:"bytes,60,opt,name=VerificationCode,proto3" json:"VerificationCode,omitempty"`
	InvitationCode   string `protobuf:"bytes,70,opt,name=InvitationCode,proto3" json:"InvitationCode,omitempty"`
}

func (x *SignupRequest) Reset() {
	*x = SignupRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_appuser_gw_v1_user_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignupRequest) ProtoMessage() {}

func (x *SignupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_appuser_gw_v1_user_user_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignupRequest.ProtoReflect.Descriptor instead.
func (*SignupRequest) Descriptor() ([]byte, []int) {
	return file_npool_appuser_gw_v1_user_user_proto_rawDescGZIP(), []int{0}
}

func (x *SignupRequest) GetAppID() string {
	if x != nil {
		return x.AppID
	}
	return ""
}

func (x *SignupRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *SignupRequest) GetPasswordHash() string {
	if x != nil {
		return x.PasswordHash
	}
	return ""
}

func (x *SignupRequest) GetAccount() string {
	if x != nil {
		return x.Account
	}
	return ""
}

func (x *SignupRequest) GetAccountType() string {
	if x != nil {
		return x.AccountType
	}
	return ""
}

func (x *SignupRequest) GetVerificationCode() string {
	if x != nil {
		return x.VerificationCode
	}
	return ""
}

func (x *SignupRequest) GetInvitationCode() string {
	if x != nil {
		return x.InvitationCode
	}
	return ""
}

type SignupResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *user.User `protobuf:"bytes,10,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (x *SignupResponse) Reset() {
	*x = SignupResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_appuser_gw_v1_user_user_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignupResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignupResponse) ProtoMessage() {}

func (x *SignupResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_appuser_gw_v1_user_user_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignupResponse.ProtoReflect.Descriptor instead.
func (*SignupResponse) Descriptor() ([]byte, []int) {
	return file_npool_appuser_gw_v1_user_user_proto_rawDescGZIP(), []int{1}
}

func (x *SignupResponse) GetInfo() *user.User {
	if x != nil {
		return x.Info
	}
	return nil
}

type CreateUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *user.UserReq `protobuf:"bytes,10,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (x *CreateUserRequest) Reset() {
	*x = CreateUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_appuser_gw_v1_user_user_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUserRequest) ProtoMessage() {}

func (x *CreateUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_appuser_gw_v1_user_user_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateUserRequest.ProtoReflect.Descriptor instead.
func (*CreateUserRequest) Descriptor() ([]byte, []int) {
	return file_npool_appuser_gw_v1_user_user_proto_rawDescGZIP(), []int{2}
}

func (x *CreateUserRequest) GetInfo() *user.UserReq {
	if x != nil {
		return x.Info
	}
	return nil
}

type CreateUserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *user.User `protobuf:"bytes,10,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (x *CreateUserResponse) Reset() {
	*x = CreateUserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_appuser_gw_v1_user_user_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUserResponse) ProtoMessage() {}

func (x *CreateUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_appuser_gw_v1_user_user_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateUserResponse.ProtoReflect.Descriptor instead.
func (*CreateUserResponse) Descriptor() ([]byte, []int) {
	return file_npool_appuser_gw_v1_user_user_proto_rawDescGZIP(), []int{3}
}

func (x *CreateUserResponse) GetInfo() *user.User {
	if x != nil {
		return x.Info
	}
	return nil
}

type UpdateUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *user.UserReq `protobuf:"bytes,10,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (x *UpdateUserRequest) Reset() {
	*x = UpdateUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_appuser_gw_v1_user_user_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateUserRequest) ProtoMessage() {}

func (x *UpdateUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_appuser_gw_v1_user_user_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateUserRequest.ProtoReflect.Descriptor instead.
func (*UpdateUserRequest) Descriptor() ([]byte, []int) {
	return file_npool_appuser_gw_v1_user_user_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateUserRequest) GetInfo() *user.UserReq {
	if x != nil {
		return x.Info
	}
	return nil
}

type UpdateUserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *user.User `protobuf:"bytes,10,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (x *UpdateUserResponse) Reset() {
	*x = UpdateUserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_appuser_gw_v1_user_user_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateUserResponse) ProtoMessage() {}

func (x *UpdateUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_appuser_gw_v1_user_user_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateUserResponse.ProtoReflect.Descriptor instead.
func (*UpdateUserResponse) Descriptor() ([]byte, []int) {
	return file_npool_appuser_gw_v1_user_user_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateUserResponse) GetInfo() *user.User {
	if x != nil {
		return x.Info
	}
	return nil
}

type GetUsersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AppID  string `protobuf:"bytes,10,opt,name=AppID,proto3" json:"AppID,omitempty"`
	Offset int32  `protobuf:"varint,20,opt,name=Offset,proto3" json:"Offset,omitempty"`
	Limit  int32  `protobuf:"varint,30,opt,name=Limit,proto3" json:"Limit,omitempty"`
}

func (x *GetUsersRequest) Reset() {
	*x = GetUsersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_appuser_gw_v1_user_user_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUsersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUsersRequest) ProtoMessage() {}

func (x *GetUsersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_appuser_gw_v1_user_user_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUsersRequest.ProtoReflect.Descriptor instead.
func (*GetUsersRequest) Descriptor() ([]byte, []int) {
	return file_npool_appuser_gw_v1_user_user_proto_rawDescGZIP(), []int{6}
}

func (x *GetUsersRequest) GetAppID() string {
	if x != nil {
		return x.AppID
	}
	return ""
}

func (x *GetUsersRequest) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *GetUsersRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type GetUsersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Infos []*user.User `protobuf:"bytes,10,rep,name=Infos,proto3" json:"Infos,omitempty"`
	Total uint32       `protobuf:"varint,40,opt,name=Total,proto3" json:"Total,omitempty"`
}

func (x *GetUsersResponse) Reset() {
	*x = GetUsersResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_appuser_gw_v1_user_user_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUsersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUsersResponse) ProtoMessage() {}

func (x *GetUsersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_appuser_gw_v1_user_user_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUsersResponse.ProtoReflect.Descriptor instead.
func (*GetUsersResponse) Descriptor() ([]byte, []int) {
	return file_npool_appuser_gw_v1_user_user_proto_rawDescGZIP(), []int{7}
}

func (x *GetUsersResponse) GetInfos() []*user.User {
	if x != nil {
		return x.Infos
	}
	return nil
}

func (x *GetUsersResponse) GetTotal() uint32 {
	if x != nil {
		return x.Total
	}
	return 0
}

type GetAppUsersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TargetAppID string `protobuf:"bytes,10,opt,name=TargetAppID,proto3" json:"TargetAppID,omitempty"`
	Offset      int32  `protobuf:"varint,20,opt,name=Offset,proto3" json:"Offset,omitempty"`
	Limit       int32  `protobuf:"varint,30,opt,name=Limit,proto3" json:"Limit,omitempty"`
}

func (x *GetAppUsersRequest) Reset() {
	*x = GetAppUsersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_appuser_gw_v1_user_user_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAppUsersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAppUsersRequest) ProtoMessage() {}

func (x *GetAppUsersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_appuser_gw_v1_user_user_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAppUsersRequest.ProtoReflect.Descriptor instead.
func (*GetAppUsersRequest) Descriptor() ([]byte, []int) {
	return file_npool_appuser_gw_v1_user_user_proto_rawDescGZIP(), []int{8}
}

func (x *GetAppUsersRequest) GetTargetAppID() string {
	if x != nil {
		return x.TargetAppID
	}
	return ""
}

func (x *GetAppUsersRequest) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *GetAppUsersRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type GetAppUsersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Infos []*user.User `protobuf:"bytes,10,rep,name=Infos,proto3" json:"Infos,omitempty"`
	Total uint32       `protobuf:"varint,40,opt,name=Total,proto3" json:"Total,omitempty"`
}

func (x *GetAppUsersResponse) Reset() {
	*x = GetAppUsersResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_appuser_gw_v1_user_user_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAppUsersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAppUsersResponse) ProtoMessage() {}

func (x *GetAppUsersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_appuser_gw_v1_user_user_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAppUsersResponse.ProtoReflect.Descriptor instead.
func (*GetAppUsersResponse) Descriptor() ([]byte, []int) {
	return file_npool_appuser_gw_v1_user_user_proto_rawDescGZIP(), []int{9}
}

func (x *GetAppUsersResponse) GetInfos() []*user.User {
	if x != nil {
		return x.Infos
	}
	return nil
}

func (x *GetAppUsersResponse) GetTotal() uint32 {
	if x != nil {
		return x.Total
	}
	return 0
}

var File_npool_appuser_gw_v1_user_user_proto protoreflect.FileDescriptor

var file_npool_appuser_gw_v1_user_user_proto_rawDesc = []byte{
	0x0a, 0x23, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x61, 0x70, 0x70, 0x75, 0x73, 0x65, 0x72, 0x2f,
	0x67, 0x77, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x17, 0x61, 0x70, 0x70, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x67,
	0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x1a, 0x1c,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2a, 0x6e, 0x70,
	0x6f, 0x6f, 0x6c, 0x2f, 0x61, 0x70, 0x70, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x6d, 0x67, 0x72, 0x2f,
	0x76, 0x32, 0x2f, 0x61, 0x70, 0x70, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x70, 0x75, 0x73,
	0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x23, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f,
	0x61, 0x70, 0x70, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x6d, 0x77, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73,
	0x65, 0x72, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf5, 0x01,
	0x0a, 0x0d, 0x53, 0x69, 0x67, 0x6e, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x14, 0x0a, 0x05, 0x41, 0x70, 0x70, 0x49, 0x44, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x41, 0x70, 0x70, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x22, 0x0a, 0x0c, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x48, 0x61, 0x73,
	0x68, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72,
	0x64, 0x48, 0x61, 0x73, 0x68, 0x12, 0x18, 0x0a, 0x07, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x28, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12,
	0x20, 0x0a, 0x0b, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x18, 0x32,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x2a, 0x0a, 0x10, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x3c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x56, 0x65, 0x72,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x26, 0x0a,
	0x0e, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x64, 0x65, 0x18,
	0x46, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x43, 0x6f, 0x64, 0x65, 0x22, 0x46, 0x0a, 0x0e, 0x53, 0x69, 0x67, 0x6e, 0x75, 0x70, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x34, 0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x61, 0x70, 0x70, 0x75, 0x73, 0x65, 0x72, 0x2e,
	0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e,
	0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x4c, 0x0a,
	0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x37, 0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x23, 0x2e, 0x61, 0x70, 0x70, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c,
	0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x52, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x4a, 0x0a, 0x12, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x34, 0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x20, 0x2e, 0x61, 0x70, 0x70, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65,
	0x77, 0x61, 0x72, 0x65, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65,
	0x72, 0x52, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x4c, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x37, 0x0a, 0x04,
	0x49, 0x6e, 0x66, 0x6f, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x61, 0x70, 0x70,
	0x75, 0x73, 0x65, 0x72, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x52,
	0x04, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x4a, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55,
	0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x34, 0x0a, 0x04, 0x49,
	0x6e, 0x66, 0x6f, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x61, 0x70, 0x70, 0x75,
	0x73, 0x65, 0x72, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x75,
	0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x49, 0x6e, 0x66,
	0x6f, 0x22, 0x55, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x41, 0x70, 0x70, 0x49, 0x44, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x41, 0x70, 0x70, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x4f, 0x66,
	0x66, 0x73, 0x65, 0x74, 0x18, 0x14, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x4f, 0x66, 0x66, 0x73,
	0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x1e, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x05, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x22, 0x60, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x55,
	0x73, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x36, 0x0a, 0x05,
	0x49, 0x6e, 0x66, 0x6f, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x61, 0x70,
	0x70, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x05, 0x49,
	0x6e, 0x66, 0x6f, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x28, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x22, 0x64, 0x0a, 0x12, 0x47, 0x65,
	0x74, 0x41, 0x70, 0x70, 0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x20, 0x0a, 0x0b, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x41, 0x70, 0x70, 0x49, 0x44, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x41, 0x70, 0x70,
	0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x14, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x06, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x4c, 0x69,
	0x6d, 0x69, 0x74, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x4c, 0x69, 0x6d, 0x69, 0x74,
	0x22, 0x63, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x41, 0x70, 0x70, 0x55, 0x73, 0x65, 0x72, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x36, 0x0a, 0x05, 0x49, 0x6e, 0x66, 0x6f, 0x73,
	0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x61, 0x70, 0x70, 0x75, 0x73, 0x65, 0x72,
	0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x05, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x12,
	0x14, 0x0a, 0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x28, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05,
	0x54, 0x6f, 0x74, 0x61, 0x6c, 0x32, 0x84, 0x05, 0x0a, 0x07, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61,
	0x79, 0x12, 0x6f, 0x0a, 0x06, 0x53, 0x69, 0x67, 0x6e, 0x75, 0x70, 0x12, 0x26, 0x2e, 0x61, 0x70,
	0x70, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x61, 0x70, 0x70, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x67, 0x61,
	0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x69,
	0x67, 0x6e, 0x75, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x14, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x0e, 0x22, 0x09, 0x2f, 0x31, 0x2f, 0x73, 0x69, 0x67, 0x6e, 0x75, 0x70, 0x3a,
	0x01, 0x2a, 0x12, 0x80, 0x01, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65,
	0x72, 0x12, 0x2a, 0x2e, 0x61, 0x70, 0x70, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x67, 0x61, 0x74, 0x65,
	0x77, 0x61, 0x79, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2b, 0x2e,
	0x61, 0x70, 0x70, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73,
	0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x19, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x13, 0x22, 0x0e, 0x2f, 0x31, 0x2f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x2f, 0x75, 0x73,
	0x65, 0x72, 0x3a, 0x01, 0x2a, 0x12, 0x80, 0x01, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x55, 0x73, 0x65, 0x72, 0x12, 0x2a, 0x2e, 0x61, 0x70, 0x70, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x67,
	0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x2b, 0x2e, 0x61, 0x70, 0x70, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77,
	0x61, 0x79, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x19, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x13, 0x22, 0x0e, 0x2f, 0x31, 0x2f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x2f, 0x75, 0x73, 0x65, 0x72, 0x3a, 0x01, 0x2a, 0x12, 0x79, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x55,
	0x73, 0x65, 0x72, 0x73, 0x12, 0x28, 0x2e, 0x61, 0x70, 0x70, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x67,
	0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47,
	0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29,
	0x2e, 0x61, 0x70, 0x70, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x18, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x12, 0x22, 0x0d, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x65, 0x74, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x73,
	0x3a, 0x01, 0x2a, 0x12, 0x86, 0x01, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x41, 0x70, 0x70, 0x55, 0x73,
	0x65, 0x72, 0x73, 0x12, 0x2b, 0x2e, 0x61, 0x70, 0x70, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x67, 0x61,
	0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65,
	0x74, 0x41, 0x70, 0x70, 0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x2c, 0x2e, 0x61, 0x70, 0x70, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77,
	0x61, 0x79, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x70,
	0x70, 0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1c,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x16, 0x22, 0x11, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x65, 0x74, 0x2f,
	0x61, 0x70, 0x70, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x3a, 0x01, 0x2a, 0x42, 0x3b, 0x5a, 0x39,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4e, 0x70, 0x6f, 0x6f, 0x6c,
	0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x2f, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x61, 0x70, 0x70, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x67,
	0x77, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_npool_appuser_gw_v1_user_user_proto_rawDescOnce sync.Once
	file_npool_appuser_gw_v1_user_user_proto_rawDescData = file_npool_appuser_gw_v1_user_user_proto_rawDesc
)

func file_npool_appuser_gw_v1_user_user_proto_rawDescGZIP() []byte {
	file_npool_appuser_gw_v1_user_user_proto_rawDescOnce.Do(func() {
		file_npool_appuser_gw_v1_user_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_npool_appuser_gw_v1_user_user_proto_rawDescData)
	})
	return file_npool_appuser_gw_v1_user_user_proto_rawDescData
}

var file_npool_appuser_gw_v1_user_user_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_npool_appuser_gw_v1_user_user_proto_goTypes = []interface{}{
	(*SignupRequest)(nil),       // 0: appuser.gateway.user.v1.SignupRequest
	(*SignupResponse)(nil),      // 1: appuser.gateway.user.v1.SignupResponse
	(*CreateUserRequest)(nil),   // 2: appuser.gateway.user.v1.CreateUserRequest
	(*CreateUserResponse)(nil),  // 3: appuser.gateway.user.v1.CreateUserResponse
	(*UpdateUserRequest)(nil),   // 4: appuser.gateway.user.v1.UpdateUserRequest
	(*UpdateUserResponse)(nil),  // 5: appuser.gateway.user.v1.UpdateUserResponse
	(*GetUsersRequest)(nil),     // 6: appuser.gateway.user.v1.GetUsersRequest
	(*GetUsersResponse)(nil),    // 7: appuser.gateway.user.v1.GetUsersResponse
	(*GetAppUsersRequest)(nil),  // 8: appuser.gateway.user.v1.GetAppUsersRequest
	(*GetAppUsersResponse)(nil), // 9: appuser.gateway.user.v1.GetAppUsersResponse
	(*user.User)(nil),           // 10: appuser.middleware.user.v1.User
	(*user.UserReq)(nil),        // 11: appuser.middleware.user.v1.UserReq
}
var file_npool_appuser_gw_v1_user_user_proto_depIdxs = []int32{
	10, // 0: appuser.gateway.user.v1.SignupResponse.Info:type_name -> appuser.middleware.user.v1.User
	11, // 1: appuser.gateway.user.v1.CreateUserRequest.Info:type_name -> appuser.middleware.user.v1.UserReq
	10, // 2: appuser.gateway.user.v1.CreateUserResponse.Info:type_name -> appuser.middleware.user.v1.User
	11, // 3: appuser.gateway.user.v1.UpdateUserRequest.Info:type_name -> appuser.middleware.user.v1.UserReq
	10, // 4: appuser.gateway.user.v1.UpdateUserResponse.Info:type_name -> appuser.middleware.user.v1.User
	10, // 5: appuser.gateway.user.v1.GetUsersResponse.Infos:type_name -> appuser.middleware.user.v1.User
	10, // 6: appuser.gateway.user.v1.GetAppUsersResponse.Infos:type_name -> appuser.middleware.user.v1.User
	0,  // 7: appuser.gateway.user.v1.Gateway.Signup:input_type -> appuser.gateway.user.v1.SignupRequest
	2,  // 8: appuser.gateway.user.v1.Gateway.CreateUser:input_type -> appuser.gateway.user.v1.CreateUserRequest
	4,  // 9: appuser.gateway.user.v1.Gateway.UpdateUser:input_type -> appuser.gateway.user.v1.UpdateUserRequest
	6,  // 10: appuser.gateway.user.v1.Gateway.GetUsers:input_type -> appuser.gateway.user.v1.GetUsersRequest
	8,  // 11: appuser.gateway.user.v1.Gateway.GetAppUsers:input_type -> appuser.gateway.user.v1.GetAppUsersRequest
	1,  // 12: appuser.gateway.user.v1.Gateway.Signup:output_type -> appuser.gateway.user.v1.SignupResponse
	3,  // 13: appuser.gateway.user.v1.Gateway.CreateUser:output_type -> appuser.gateway.user.v1.CreateUserResponse
	5,  // 14: appuser.gateway.user.v1.Gateway.UpdateUser:output_type -> appuser.gateway.user.v1.UpdateUserResponse
	7,  // 15: appuser.gateway.user.v1.Gateway.GetUsers:output_type -> appuser.gateway.user.v1.GetUsersResponse
	9,  // 16: appuser.gateway.user.v1.Gateway.GetAppUsers:output_type -> appuser.gateway.user.v1.GetAppUsersResponse
	12, // [12:17] is the sub-list for method output_type
	7,  // [7:12] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_npool_appuser_gw_v1_user_user_proto_init() }
func file_npool_appuser_gw_v1_user_user_proto_init() {
	if File_npool_appuser_gw_v1_user_user_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_npool_appuser_gw_v1_user_user_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignupRequest); i {
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
		file_npool_appuser_gw_v1_user_user_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignupResponse); i {
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
		file_npool_appuser_gw_v1_user_user_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateUserRequest); i {
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
		file_npool_appuser_gw_v1_user_user_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateUserResponse); i {
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
		file_npool_appuser_gw_v1_user_user_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateUserRequest); i {
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
		file_npool_appuser_gw_v1_user_user_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateUserResponse); i {
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
		file_npool_appuser_gw_v1_user_user_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUsersRequest); i {
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
		file_npool_appuser_gw_v1_user_user_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUsersResponse); i {
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
		file_npool_appuser_gw_v1_user_user_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAppUsersRequest); i {
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
		file_npool_appuser_gw_v1_user_user_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAppUsersResponse); i {
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
			RawDescriptor: file_npool_appuser_gw_v1_user_user_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_npool_appuser_gw_v1_user_user_proto_goTypes,
		DependencyIndexes: file_npool_appuser_gw_v1_user_user_proto_depIdxs,
		MessageInfos:      file_npool_appuser_gw_v1_user_user_proto_msgTypes,
	}.Build()
	File_npool_appuser_gw_v1_user_user_proto = out.File
	file_npool_appuser_gw_v1_user_user_proto_rawDesc = nil
	file_npool_appuser_gw_v1_user_user_proto_goTypes = nil
	file_npool_appuser_gw_v1_user_user_proto_depIdxs = nil
}
