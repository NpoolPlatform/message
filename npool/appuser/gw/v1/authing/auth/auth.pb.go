// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.18.1
// source: npool/appuser/gw/v1/authing/auth/auth.proto

package auth

import (
	auth "github.com/NpoolPlatform/message/npool/appuser/mw/v1/authing/auth"
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

type AuthenticateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AppID    string  `protobuf:"bytes,10,opt,name=AppID,proto3" json:"AppID,omitempty"`
	UserID   *string `protobuf:"bytes,20,opt,name=UserID,proto3,oneof" json:"UserID,omitempty"`
	Token    *string `protobuf:"bytes,30,opt,name=Token,proto3,oneof" json:"Token,omitempty"`
	Resource string  `protobuf:"bytes,40,opt,name=Resource,proto3" json:"Resource,omitempty"`
	Method   string  `protobuf:"bytes,50,opt,name=Method,proto3" json:"Method,omitempty"`
}

func (x *AuthenticateRequest) Reset() {
	*x = AuthenticateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_appuser_gw_v1_authing_auth_auth_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthenticateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthenticateRequest) ProtoMessage() {}

func (x *AuthenticateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_appuser_gw_v1_authing_auth_auth_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthenticateRequest.ProtoReflect.Descriptor instead.
func (*AuthenticateRequest) Descriptor() ([]byte, []int) {
	return file_npool_appuser_gw_v1_authing_auth_auth_proto_rawDescGZIP(), []int{0}
}

func (x *AuthenticateRequest) GetAppID() string {
	if x != nil {
		return x.AppID
	}
	return ""
}

func (x *AuthenticateRequest) GetUserID() string {
	if x != nil && x.UserID != nil {
		return *x.UserID
	}
	return ""
}

func (x *AuthenticateRequest) GetToken() string {
	if x != nil && x.Token != nil {
		return *x.Token
	}
	return ""
}

func (x *AuthenticateRequest) GetResource() string {
	if x != nil {
		return x.Resource
	}
	return ""
}

func (x *AuthenticateRequest) GetMethod() string {
	if x != nil {
		return x.Method
	}
	return ""
}

type AuthenticateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info bool `protobuf:"varint,10,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (x *AuthenticateResponse) Reset() {
	*x = AuthenticateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_appuser_gw_v1_authing_auth_auth_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthenticateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthenticateResponse) ProtoMessage() {}

func (x *AuthenticateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_appuser_gw_v1_authing_auth_auth_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthenticateResponse.ProtoReflect.Descriptor instead.
func (*AuthenticateResponse) Descriptor() ([]byte, []int) {
	return file_npool_appuser_gw_v1_authing_auth_auth_proto_rawDescGZIP(), []int{1}
}

func (x *AuthenticateResponse) GetInfo() bool {
	if x != nil {
		return x.Info
	}
	return false
}

type CreateAppAuthRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TargetAppID  string  `protobuf:"bytes,10,opt,name=TargetAppID,proto3" json:"TargetAppID,omitempty"`
	TargetUserID *string `protobuf:"bytes,20,opt,name=TargetUserID,proto3,oneof" json:"TargetUserID,omitempty"`
	RoleID       *string `protobuf:"bytes,30,opt,name=RoleID,proto3,oneof" json:"RoleID,omitempty"`
	Resource     string  `protobuf:"bytes,40,opt,name=Resource,proto3" json:"Resource,omitempty"`
	Method       string  `protobuf:"bytes,50,opt,name=Method,proto3" json:"Method,omitempty"`
}

func (x *CreateAppAuthRequest) Reset() {
	*x = CreateAppAuthRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_appuser_gw_v1_authing_auth_auth_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateAppAuthRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAppAuthRequest) ProtoMessage() {}

func (x *CreateAppAuthRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_appuser_gw_v1_authing_auth_auth_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAppAuthRequest.ProtoReflect.Descriptor instead.
func (*CreateAppAuthRequest) Descriptor() ([]byte, []int) {
	return file_npool_appuser_gw_v1_authing_auth_auth_proto_rawDescGZIP(), []int{2}
}

func (x *CreateAppAuthRequest) GetTargetAppID() string {
	if x != nil {
		return x.TargetAppID
	}
	return ""
}

func (x *CreateAppAuthRequest) GetTargetUserID() string {
	if x != nil && x.TargetUserID != nil {
		return *x.TargetUserID
	}
	return ""
}

func (x *CreateAppAuthRequest) GetRoleID() string {
	if x != nil && x.RoleID != nil {
		return *x.RoleID
	}
	return ""
}

func (x *CreateAppAuthRequest) GetResource() string {
	if x != nil {
		return x.Resource
	}
	return ""
}

func (x *CreateAppAuthRequest) GetMethod() string {
	if x != nil {
		return x.Method
	}
	return ""
}

type CreateAppAuthResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *auth.Auth `protobuf:"bytes,10,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (x *CreateAppAuthResponse) Reset() {
	*x = CreateAppAuthResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_appuser_gw_v1_authing_auth_auth_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateAppAuthResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAppAuthResponse) ProtoMessage() {}

func (x *CreateAppAuthResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_appuser_gw_v1_authing_auth_auth_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAppAuthResponse.ProtoReflect.Descriptor instead.
func (*CreateAppAuthResponse) Descriptor() ([]byte, []int) {
	return file_npool_appuser_gw_v1_authing_auth_auth_proto_rawDescGZIP(), []int{3}
}

func (x *CreateAppAuthResponse) GetInfo() *auth.Auth {
	if x != nil {
		return x.Info
	}
	return nil
}

type DeleteAppAuthRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TargetAppID string `protobuf:"bytes,10,opt,name=TargetAppID,proto3" json:"TargetAppID,omitempty"`
	ID          uint32 `protobuf:"varint,20,opt,name=ID,proto3" json:"ID,omitempty"`
	EntID       string `protobuf:"bytes,30,opt,name=EntID,proto3" json:"EntID,omitempty"`
}

func (x *DeleteAppAuthRequest) Reset() {
	*x = DeleteAppAuthRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_appuser_gw_v1_authing_auth_auth_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteAppAuthRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteAppAuthRequest) ProtoMessage() {}

func (x *DeleteAppAuthRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_appuser_gw_v1_authing_auth_auth_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteAppAuthRequest.ProtoReflect.Descriptor instead.
func (*DeleteAppAuthRequest) Descriptor() ([]byte, []int) {
	return file_npool_appuser_gw_v1_authing_auth_auth_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteAppAuthRequest) GetTargetAppID() string {
	if x != nil {
		return x.TargetAppID
	}
	return ""
}

func (x *DeleteAppAuthRequest) GetID() uint32 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *DeleteAppAuthRequest) GetEntID() string {
	if x != nil {
		return x.EntID
	}
	return ""
}

type DeleteAppAuthResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *auth.Auth `protobuf:"bytes,10,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (x *DeleteAppAuthResponse) Reset() {
	*x = DeleteAppAuthResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_appuser_gw_v1_authing_auth_auth_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteAppAuthResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteAppAuthResponse) ProtoMessage() {}

func (x *DeleteAppAuthResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_appuser_gw_v1_authing_auth_auth_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteAppAuthResponse.ProtoReflect.Descriptor instead.
func (*DeleteAppAuthResponse) Descriptor() ([]byte, []int) {
	return file_npool_appuser_gw_v1_authing_auth_auth_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteAppAuthResponse) GetInfo() *auth.Auth {
	if x != nil {
		return x.Info
	}
	return nil
}

type GetAppAuthsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TargetAppID string `protobuf:"bytes,10,opt,name=TargetAppID,proto3" json:"TargetAppID,omitempty"`
	Offset      int32  `protobuf:"varint,20,opt,name=Offset,proto3" json:"Offset,omitempty"`
	Limit       int32  `protobuf:"varint,30,opt,name=Limit,proto3" json:"Limit,omitempty"`
}

func (x *GetAppAuthsRequest) Reset() {
	*x = GetAppAuthsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_appuser_gw_v1_authing_auth_auth_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAppAuthsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAppAuthsRequest) ProtoMessage() {}

func (x *GetAppAuthsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_appuser_gw_v1_authing_auth_auth_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAppAuthsRequest.ProtoReflect.Descriptor instead.
func (*GetAppAuthsRequest) Descriptor() ([]byte, []int) {
	return file_npool_appuser_gw_v1_authing_auth_auth_proto_rawDescGZIP(), []int{6}
}

func (x *GetAppAuthsRequest) GetTargetAppID() string {
	if x != nil {
		return x.TargetAppID
	}
	return ""
}

func (x *GetAppAuthsRequest) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *GetAppAuthsRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type GetAppAuthsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Infos []*auth.Auth `protobuf:"bytes,10,rep,name=Infos,proto3" json:"Infos,omitempty"`
	Total uint32       `protobuf:"varint,20,opt,name=Total,proto3" json:"Total,omitempty"`
}

func (x *GetAppAuthsResponse) Reset() {
	*x = GetAppAuthsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_appuser_gw_v1_authing_auth_auth_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAppAuthsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAppAuthsResponse) ProtoMessage() {}

func (x *GetAppAuthsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_appuser_gw_v1_authing_auth_auth_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAppAuthsResponse.ProtoReflect.Descriptor instead.
func (*GetAppAuthsResponse) Descriptor() ([]byte, []int) {
	return file_npool_appuser_gw_v1_authing_auth_auth_proto_rawDescGZIP(), []int{7}
}

func (x *GetAppAuthsResponse) GetInfos() []*auth.Auth {
	if x != nil {
		return x.Infos
	}
	return nil
}

func (x *GetAppAuthsResponse) GetTotal() uint32 {
	if x != nil {
		return x.Total
	}
	return 0
}

var File_npool_appuser_gw_v1_authing_auth_auth_proto protoreflect.FileDescriptor

var file_npool_appuser_gw_v1_authing_auth_auth_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x61, 0x70, 0x70, 0x75, 0x73, 0x65, 0x72, 0x2f,
	0x67, 0x77, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x2f, 0x61, 0x75,
	0x74, 0x68, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1f, 0x61,
	0x70, 0x70, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x61,
	0x75, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x1a, 0x1c,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2b, 0x6e, 0x70,
	0x6f, 0x6f, 0x6c, 0x2f, 0x61, 0x70, 0x70, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x6d, 0x77, 0x2f, 0x76,
	0x31, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x61,
	0x75, 0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xac, 0x01, 0x0a, 0x13, 0x41, 0x75,
	0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x14, 0x0a, 0x05, 0x41, 0x70, 0x70, 0x49, 0x44, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x41, 0x70, 0x70, 0x49, 0x44, 0x12, 0x1b, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49,
	0x44, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49,
	0x44, 0x88, 0x01, 0x01, 0x12, 0x19, 0x0a, 0x05, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x1e, 0x20,
	0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x05, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x88, 0x01, 0x01, 0x12,
	0x1a, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x28, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x4d,
	0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x32, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x4d, 0x65, 0x74,
	0x68, 0x6f, 0x64, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x42, 0x08,
	0x0a, 0x06, 0x5f, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x2a, 0x0a, 0x14, 0x41, 0x75, 0x74, 0x68,
	0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04,
	0x49, 0x6e, 0x66, 0x6f, 0x22, 0xce, 0x01, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41,
	0x70, 0x70, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x20, 0x0a,
	0x0b, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x41, 0x70, 0x70, 0x49, 0x44, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x41, 0x70, 0x70, 0x49, 0x44, 0x12,
	0x27, 0x0a, 0x0c, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18,
	0x14, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0c, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x55,
	0x73, 0x65, 0x72, 0x49, 0x44, 0x88, 0x01, 0x01, 0x12, 0x1b, 0x0a, 0x06, 0x52, 0x6f, 0x6c, 0x65,
	0x49, 0x44, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x06, 0x52, 0x6f, 0x6c, 0x65,
	0x49, 0x44, 0x88, 0x01, 0x01, 0x12, 0x1a, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x18, 0x28, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x32, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x42, 0x0f, 0x0a, 0x0d, 0x5f, 0x54, 0x61,
	0x72, 0x67, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x52,
	0x6f, 0x6c, 0x65, 0x49, 0x44, 0x22, 0x55, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41,
	0x70, 0x70, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3c,
	0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x61,
	0x70, 0x70, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72,
	0x65, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x76,
	0x31, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x52, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x5e, 0x0a, 0x14,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x70, 0x70, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x41, 0x70,
	0x70, 0x49, 0x44, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x54, 0x61, 0x72, 0x67, 0x65,
	0x74, 0x41, 0x70, 0x70, 0x49, 0x44, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x14, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x02, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x45, 0x6e, 0x74, 0x49, 0x44, 0x18,
	0x1e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x45, 0x6e, 0x74, 0x49, 0x44, 0x22, 0x55, 0x0a, 0x15,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x70, 0x70, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3c, 0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x61, 0x70, 0x70, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x6d, 0x69,
	0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x69, 0x6e, 0x67,
	0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x52, 0x04, 0x49,
	0x6e, 0x66, 0x6f, 0x22, 0x64, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x41, 0x70, 0x70, 0x41, 0x75, 0x74,
	0x68, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x54, 0x61, 0x72,
	0x67, 0x65, 0x74, 0x41, 0x70, 0x70, 0x49, 0x44, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x41, 0x70, 0x70, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x4f,
	0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x14, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x4f, 0x66, 0x66,
	0x73, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x1e, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x05, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x22, 0x6b, 0x0a, 0x13, 0x47, 0x65, 0x74,
	0x41, 0x70, 0x70, 0x41, 0x75, 0x74, 0x68, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x3e, 0x0a, 0x05, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x28, 0x2e, 0x61, 0x70, 0x70, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65,
	0x77, 0x61, 0x72, 0x65, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x2e, 0x61, 0x75, 0x74,
	0x68, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x52, 0x05, 0x49, 0x6e, 0x66, 0x6f, 0x73,
	0x12, 0x14, 0x0a, 0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x14, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x32, 0xff, 0x04, 0x0a, 0x07, 0x47, 0x61, 0x74, 0x65, 0x77,
	0x61, 0x79, 0x12, 0x98, 0x01, 0x0a, 0x0c, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63,
	0x61, 0x74, 0x65, 0x12, 0x34, 0x2e, 0x61, 0x70, 0x70, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x67, 0x61,
	0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x2e, 0x61, 0x75,
	0x74, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x35, 0x2e, 0x61, 0x70, 0x70, 0x75,
	0x73, 0x65, 0x72, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x61, 0x75, 0x74, 0x68,
	0x69, 0x6e, 0x67, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x75, 0x74, 0x68,
	0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x1b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15, 0x3a, 0x01, 0x2a, 0x22, 0x10, 0x2f, 0x76, 0x31,
	0x2f, 0x61, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x65, 0x12, 0x9e, 0x01,
	0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x70, 0x70, 0x41, 0x75, 0x74, 0x68, 0x12,
	0x35, 0x2e, 0x61, 0x70, 0x70, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61,
	0x79, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x76,
	0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x70, 0x70, 0x41, 0x75, 0x74, 0x68, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x36, 0x2e, 0x61, 0x70, 0x70, 0x75, 0x73, 0x65, 0x72,
	0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x69, 0x6e, 0x67,
	0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41,
	0x70, 0x70, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1e,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x18, 0x3a, 0x01, 0x2a, 0x22, 0x13, 0x2f, 0x76, 0x31, 0x2f, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x2f, 0x61, 0x70, 0x70, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x12, 0x9e,
	0x01, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x70, 0x70, 0x41, 0x75, 0x74, 0x68,
	0x12, 0x35, 0x2e, 0x61, 0x70, 0x70, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77,
	0x61, 0x79, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e,
	0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x70, 0x70, 0x41, 0x75, 0x74, 0x68,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x36, 0x2e, 0x61, 0x70, 0x70, 0x75, 0x73, 0x65,
	0x72, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x69, 0x6e,
	0x67, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x41, 0x70, 0x70, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x1e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x18, 0x3a, 0x01, 0x2a, 0x22, 0x13, 0x2f, 0x76, 0x31, 0x2f,
	0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x2f, 0x61, 0x70, 0x70, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x12,
	0x96, 0x01, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x41, 0x70, 0x70, 0x41, 0x75, 0x74, 0x68, 0x73, 0x12,
	0x33, 0x2e, 0x61, 0x70, 0x70, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61,
	0x79, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x76,
	0x31, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x70, 0x70, 0x41, 0x75, 0x74, 0x68, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x34, 0x2e, 0x61, 0x70, 0x70, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x67,
	0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x2e, 0x61,
	0x75, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x70, 0x70, 0x41, 0x75, 0x74,
	0x68, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1c, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x16, 0x3a, 0x01, 0x2a, 0x22, 0x11, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x65, 0x74, 0x2f, 0x61,
	0x70, 0x70, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x73, 0x42, 0x43, 0x5a, 0x41, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4e, 0x70, 0x6f, 0x6f, 0x6c, 0x50, 0x6c, 0x61, 0x74,
	0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x6e, 0x70, 0x6f,
	0x6f, 0x6c, 0x2f, 0x61, 0x70, 0x70, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x67, 0x77, 0x2f, 0x76, 0x31,
	0x2f, 0x61, 0x75, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_npool_appuser_gw_v1_authing_auth_auth_proto_rawDescOnce sync.Once
	file_npool_appuser_gw_v1_authing_auth_auth_proto_rawDescData = file_npool_appuser_gw_v1_authing_auth_auth_proto_rawDesc
)

func file_npool_appuser_gw_v1_authing_auth_auth_proto_rawDescGZIP() []byte {
	file_npool_appuser_gw_v1_authing_auth_auth_proto_rawDescOnce.Do(func() {
		file_npool_appuser_gw_v1_authing_auth_auth_proto_rawDescData = protoimpl.X.CompressGZIP(file_npool_appuser_gw_v1_authing_auth_auth_proto_rawDescData)
	})
	return file_npool_appuser_gw_v1_authing_auth_auth_proto_rawDescData
}

var file_npool_appuser_gw_v1_authing_auth_auth_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_npool_appuser_gw_v1_authing_auth_auth_proto_goTypes = []interface{}{
	(*AuthenticateRequest)(nil),   // 0: appuser.gateway.authing.auth.v1.AuthenticateRequest
	(*AuthenticateResponse)(nil),  // 1: appuser.gateway.authing.auth.v1.AuthenticateResponse
	(*CreateAppAuthRequest)(nil),  // 2: appuser.gateway.authing.auth.v1.CreateAppAuthRequest
	(*CreateAppAuthResponse)(nil), // 3: appuser.gateway.authing.auth.v1.CreateAppAuthResponse
	(*DeleteAppAuthRequest)(nil),  // 4: appuser.gateway.authing.auth.v1.DeleteAppAuthRequest
	(*DeleteAppAuthResponse)(nil), // 5: appuser.gateway.authing.auth.v1.DeleteAppAuthResponse
	(*GetAppAuthsRequest)(nil),    // 6: appuser.gateway.authing.auth.v1.GetAppAuthsRequest
	(*GetAppAuthsResponse)(nil),   // 7: appuser.gateway.authing.auth.v1.GetAppAuthsResponse
	(*auth.Auth)(nil),             // 8: appuser.middleware.authing.auth.v1.Auth
}
var file_npool_appuser_gw_v1_authing_auth_auth_proto_depIdxs = []int32{
	8, // 0: appuser.gateway.authing.auth.v1.CreateAppAuthResponse.Info:type_name -> appuser.middleware.authing.auth.v1.Auth
	8, // 1: appuser.gateway.authing.auth.v1.DeleteAppAuthResponse.Info:type_name -> appuser.middleware.authing.auth.v1.Auth
	8, // 2: appuser.gateway.authing.auth.v1.GetAppAuthsResponse.Infos:type_name -> appuser.middleware.authing.auth.v1.Auth
	0, // 3: appuser.gateway.authing.auth.v1.Gateway.Authenticate:input_type -> appuser.gateway.authing.auth.v1.AuthenticateRequest
	2, // 4: appuser.gateway.authing.auth.v1.Gateway.CreateAppAuth:input_type -> appuser.gateway.authing.auth.v1.CreateAppAuthRequest
	4, // 5: appuser.gateway.authing.auth.v1.Gateway.DeleteAppAuth:input_type -> appuser.gateway.authing.auth.v1.DeleteAppAuthRequest
	6, // 6: appuser.gateway.authing.auth.v1.Gateway.GetAppAuths:input_type -> appuser.gateway.authing.auth.v1.GetAppAuthsRequest
	1, // 7: appuser.gateway.authing.auth.v1.Gateway.Authenticate:output_type -> appuser.gateway.authing.auth.v1.AuthenticateResponse
	3, // 8: appuser.gateway.authing.auth.v1.Gateway.CreateAppAuth:output_type -> appuser.gateway.authing.auth.v1.CreateAppAuthResponse
	5, // 9: appuser.gateway.authing.auth.v1.Gateway.DeleteAppAuth:output_type -> appuser.gateway.authing.auth.v1.DeleteAppAuthResponse
	7, // 10: appuser.gateway.authing.auth.v1.Gateway.GetAppAuths:output_type -> appuser.gateway.authing.auth.v1.GetAppAuthsResponse
	7, // [7:11] is the sub-list for method output_type
	3, // [3:7] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_npool_appuser_gw_v1_authing_auth_auth_proto_init() }
func file_npool_appuser_gw_v1_authing_auth_auth_proto_init() {
	if File_npool_appuser_gw_v1_authing_auth_auth_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_npool_appuser_gw_v1_authing_auth_auth_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthenticateRequest); i {
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
		file_npool_appuser_gw_v1_authing_auth_auth_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthenticateResponse); i {
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
		file_npool_appuser_gw_v1_authing_auth_auth_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateAppAuthRequest); i {
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
		file_npool_appuser_gw_v1_authing_auth_auth_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateAppAuthResponse); i {
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
		file_npool_appuser_gw_v1_authing_auth_auth_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteAppAuthRequest); i {
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
		file_npool_appuser_gw_v1_authing_auth_auth_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteAppAuthResponse); i {
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
		file_npool_appuser_gw_v1_authing_auth_auth_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAppAuthsRequest); i {
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
		file_npool_appuser_gw_v1_authing_auth_auth_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAppAuthsResponse); i {
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
	file_npool_appuser_gw_v1_authing_auth_auth_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_npool_appuser_gw_v1_authing_auth_auth_proto_msgTypes[2].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_npool_appuser_gw_v1_authing_auth_auth_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_npool_appuser_gw_v1_authing_auth_auth_proto_goTypes,
		DependencyIndexes: file_npool_appuser_gw_v1_authing_auth_auth_proto_depIdxs,
		MessageInfos:      file_npool_appuser_gw_v1_authing_auth_auth_proto_msgTypes,
	}.Build()
	File_npool_appuser_gw_v1_authing_auth_auth_proto = out.File
	file_npool_appuser_gw_v1_authing_auth_auth_proto_rawDesc = nil
	file_npool_appuser_gw_v1_authing_auth_auth_proto_goTypes = nil
	file_npool_appuser_gw_v1_authing_auth_auth_proto_depIdxs = nil
}
