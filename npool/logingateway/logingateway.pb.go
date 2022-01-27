// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.18.1
// source: npool/logingateway/logingateway.proto

package logingateway

import (
	npool "github.com/NpoolPlatform/message/npool"
	appusermgr "github.com/NpoolPlatform/message/npool/appusermgr"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type LoginRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AppID           string `protobuf:"bytes,10,opt,name=AppID,proto3" json:"AppID,omitempty"`
	Account         string `protobuf:"bytes,20,opt,name=Account,proto3" json:"Account,omitempty"`
	PasswordHash    string `protobuf:"bytes,30,opt,name=PasswordHash,proto3" json:"PasswordHash,omitempty"`
	ManMachineSpec  string `protobuf:"bytes,40,opt,name=ManMachineSpec,proto3" json:"ManMachineSpec,omitempty"`
	EnvironmentSpec string `protobuf:"bytes,50,opt,name=EnvironmentSpec,proto3" json:"EnvironmentSpec,omitempty"`
	LoginType       string `protobuf:"bytes,60,opt,name=LoginType,proto3" json:"LoginType,omitempty"`
}

func (x *LoginRequest) Reset() {
	*x = LoginRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_logingateway_logingateway_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginRequest) ProtoMessage() {}

func (x *LoginRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_logingateway_logingateway_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginRequest.ProtoReflect.Descriptor instead.
func (*LoginRequest) Descriptor() ([]byte, []int) {
	return file_npool_logingateway_logingateway_proto_rawDescGZIP(), []int{0}
}

func (x *LoginRequest) GetAppID() string {
	if x != nil {
		return x.AppID
	}
	return ""
}

func (x *LoginRequest) GetAccount() string {
	if x != nil {
		return x.Account
	}
	return ""
}

func (x *LoginRequest) GetPasswordHash() string {
	if x != nil {
		return x.PasswordHash
	}
	return ""
}

func (x *LoginRequest) GetManMachineSpec() string {
	if x != nil {
		return x.ManMachineSpec
	}
	return ""
}

func (x *LoginRequest) GetEnvironmentSpec() string {
	if x != nil {
		return x.EnvironmentSpec
	}
	return ""
}

func (x *LoginRequest) GetLoginType() string {
	if x != nil {
		return x.LoginType
	}
	return ""
}

type LoginResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *appusermgr.AppUserInfo `protobuf:"bytes,10,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (x *LoginResponse) Reset() {
	*x = LoginResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_logingateway_logingateway_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginResponse) ProtoMessage() {}

func (x *LoginResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_logingateway_logingateway_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginResponse.ProtoReflect.Descriptor instead.
func (*LoginResponse) Descriptor() ([]byte, []int) {
	return file_npool_logingateway_logingateway_proto_rawDescGZIP(), []int{1}
}

func (x *LoginResponse) GetInfo() *appusermgr.AppUserInfo {
	if x != nil {
		return x.Info
	}
	return nil
}

type LoginedRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AppID  string `protobuf:"bytes,10,opt,name=AppID,proto3" json:"AppID,omitempty"`
	UserID string `protobuf:"bytes,20,opt,name=UserID,proto3" json:"UserID,omitempty"`
}

func (x *LoginedRequest) Reset() {
	*x = LoginedRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_logingateway_logingateway_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginedRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginedRequest) ProtoMessage() {}

func (x *LoginedRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_logingateway_logingateway_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginedRequest.ProtoReflect.Descriptor instead.
func (*LoginedRequest) Descriptor() ([]byte, []int) {
	return file_npool_logingateway_logingateway_proto_rawDescGZIP(), []int{2}
}

func (x *LoginedRequest) GetAppID() string {
	if x != nil {
		return x.AppID
	}
	return ""
}

func (x *LoginedRequest) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

type LoginedResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *appusermgr.AppUserInfo `protobuf:"bytes,10,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (x *LoginedResponse) Reset() {
	*x = LoginedResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_logingateway_logingateway_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginedResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginedResponse) ProtoMessage() {}

func (x *LoginedResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_logingateway_logingateway_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginedResponse.ProtoReflect.Descriptor instead.
func (*LoginedResponse) Descriptor() ([]byte, []int) {
	return file_npool_logingateway_logingateway_proto_rawDescGZIP(), []int{3}
}

func (x *LoginedResponse) GetInfo() *appusermgr.AppUserInfo {
	if x != nil {
		return x.Info
	}
	return nil
}

type LogoutRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AppID  string `protobuf:"bytes,10,opt,name=AppID,proto3" json:"AppID,omitempty"`
	UserID string `protobuf:"bytes,20,opt,name=UserID,proto3" json:"UserID,omitempty"`
}

func (x *LogoutRequest) Reset() {
	*x = LogoutRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_logingateway_logingateway_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogoutRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogoutRequest) ProtoMessage() {}

func (x *LogoutRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_logingateway_logingateway_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogoutRequest.ProtoReflect.Descriptor instead.
func (*LogoutRequest) Descriptor() ([]byte, []int) {
	return file_npool_logingateway_logingateway_proto_rawDescGZIP(), []int{4}
}

func (x *LogoutRequest) GetAppID() string {
	if x != nil {
		return x.AppID
	}
	return ""
}

func (x *LogoutRequest) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

type LogoutResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *appusermgr.AppUserInfo `protobuf:"bytes,10,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (x *LogoutResponse) Reset() {
	*x = LogoutResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_logingateway_logingateway_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogoutResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogoutResponse) ProtoMessage() {}

func (x *LogoutResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_logingateway_logingateway_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogoutResponse.ProtoReflect.Descriptor instead.
func (*LogoutResponse) Descriptor() ([]byte, []int) {
	return file_npool_logingateway_logingateway_proto_rawDescGZIP(), []int{5}
}

func (x *LogoutResponse) GetInfo() *appusermgr.AppUserInfo {
	if x != nil {
		return x.Info
	}
	return nil
}

type RefreshRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AppID  string `protobuf:"bytes,10,opt,name=AppID,proto3" json:"AppID,omitempty"`
	UserID string `protobuf:"bytes,20,opt,name=UserID,proto3" json:"UserID,omitempty"`
}

func (x *RefreshRequest) Reset() {
	*x = RefreshRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_logingateway_logingateway_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RefreshRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RefreshRequest) ProtoMessage() {}

func (x *RefreshRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_logingateway_logingateway_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RefreshRequest.ProtoReflect.Descriptor instead.
func (*RefreshRequest) Descriptor() ([]byte, []int) {
	return file_npool_logingateway_logingateway_proto_rawDescGZIP(), []int{6}
}

func (x *RefreshRequest) GetAppID() string {
	if x != nil {
		return x.AppID
	}
	return ""
}

func (x *RefreshRequest) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

type RefreshResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *appusermgr.AppUserInfo `protobuf:"bytes,10,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (x *RefreshResponse) Reset() {
	*x = RefreshResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_logingateway_logingateway_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RefreshResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RefreshResponse) ProtoMessage() {}

func (x *RefreshResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_logingateway_logingateway_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RefreshResponse.ProtoReflect.Descriptor instead.
func (*RefreshResponse) Descriptor() ([]byte, []int) {
	return file_npool_logingateway_logingateway_proto_rawDescGZIP(), []int{7}
}

func (x *RefreshResponse) GetInfo() *appusermgr.AppUserInfo {
	if x != nil {
		return x.Info
	}
	return nil
}

type LoginHistory struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID        string `protobuf:"bytes,10,opt,name=ID,proto3" json:"ID,omitempty"`
	AppID     string `protobuf:"bytes,20,opt,name=AppID,proto3" json:"AppID,omitempty"`
	UserID    string `protobuf:"bytes,30,opt,name=UserID,proto3" json:"UserID,omitempty"`
	ClientIP  string `protobuf:"bytes,40,opt,name=ClientIP,proto3" json:"ClientIP,omitempty"`
	UserAgent string `protobuf:"bytes,50,opt,name=UserAgent,proto3" json:"UserAgent,omitempty"`
	CreateAt  uint32 `protobuf:"varint,60,opt,name=CreateAt,proto3" json:"CreateAt,omitempty"`
}

func (x *LoginHistory) Reset() {
	*x = LoginHistory{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_logingateway_logingateway_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginHistory) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginHistory) ProtoMessage() {}

func (x *LoginHistory) ProtoReflect() protoreflect.Message {
	mi := &file_npool_logingateway_logingateway_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginHistory.ProtoReflect.Descriptor instead.
func (*LoginHistory) Descriptor() ([]byte, []int) {
	return file_npool_logingateway_logingateway_proto_rawDescGZIP(), []int{8}
}

func (x *LoginHistory) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *LoginHistory) GetAppID() string {
	if x != nil {
		return x.AppID
	}
	return ""
}

func (x *LoginHistory) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

func (x *LoginHistory) GetClientIP() string {
	if x != nil {
		return x.ClientIP
	}
	return ""
}

func (x *LoginHistory) GetUserAgent() string {
	if x != nil {
		return x.UserAgent
	}
	return ""
}

func (x *LoginHistory) GetCreateAt() uint32 {
	if x != nil {
		return x.CreateAt
	}
	return 0
}

type GetLoginHistoriesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AppID    string          `protobuf:"bytes,10,opt,name=AppID,proto3" json:"AppID,omitempty"`
	UserID   string          `protobuf:"bytes,20,opt,name=UserID,proto3" json:"UserID,omitempty"`
	PageInfo *npool.PageInfo `protobuf:"bytes,30,opt,name=PageInfo,proto3" json:"PageInfo,omitempty"`
}

func (x *GetLoginHistoriesRequest) Reset() {
	*x = GetLoginHistoriesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_logingateway_logingateway_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetLoginHistoriesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLoginHistoriesRequest) ProtoMessage() {}

func (x *GetLoginHistoriesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_logingateway_logingateway_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLoginHistoriesRequest.ProtoReflect.Descriptor instead.
func (*GetLoginHistoriesRequest) Descriptor() ([]byte, []int) {
	return file_npool_logingateway_logingateway_proto_rawDescGZIP(), []int{9}
}

func (x *GetLoginHistoriesRequest) GetAppID() string {
	if x != nil {
		return x.AppID
	}
	return ""
}

func (x *GetLoginHistoriesRequest) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

func (x *GetLoginHistoriesRequest) GetPageInfo() *npool.PageInfo {
	if x != nil {
		return x.PageInfo
	}
	return nil
}

type GetLoginHistoriesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Infos []*LoginHistory `protobuf:"bytes,10,rep,name=Infos,proto3" json:"Infos,omitempty"`
}

func (x *GetLoginHistoriesResponse) Reset() {
	*x = GetLoginHistoriesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_logingateway_logingateway_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetLoginHistoriesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLoginHistoriesResponse) ProtoMessage() {}

func (x *GetLoginHistoriesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_logingateway_logingateway_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLoginHistoriesResponse.ProtoReflect.Descriptor instead.
func (*GetLoginHistoriesResponse) Descriptor() ([]byte, []int) {
	return file_npool_logingateway_logingateway_proto_rawDescGZIP(), []int{10}
}

func (x *GetLoginHistoriesResponse) GetInfos() []*LoginHistory {
	if x != nil {
		return x.Infos
	}
	return nil
}

var File_npool_logingateway_logingateway_proto protoreflect.FileDescriptor

var file_npool_logingateway_logingateway_proto_rawDesc = []byte{
	0x0a, 0x25, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x67, 0x61, 0x74,
	0x65, 0x77, 0x61, 0x79, 0x2f, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x2e, 0x67,
	0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x6e, 0x70, 0x6f, 0x6f,
	0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x61,
	0x70, 0x70, 0x75, 0x73, 0x65, 0x72, 0x6d, 0x67, 0x72, 0x2f, 0x61, 0x70, 0x70, 0x75, 0x73, 0x65,
	0x72, 0x6d, 0x67, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd2, 0x01, 0x0a, 0x0c, 0x4c,
	0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x41,
	0x70, 0x70, 0x49, 0x44, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x41, 0x70, 0x70, 0x49,
	0x44, 0x12, 0x18, 0x0a, 0x07, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x14, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x22, 0x0a, 0x0c, 0x50,
	0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x48, 0x61, 0x73, 0x68, 0x18, 0x1e, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0c, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x48, 0x61, 0x73, 0x68, 0x12,
	0x26, 0x0a, 0x0e, 0x4d, 0x61, 0x6e, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x53, 0x70, 0x65,
	0x63, 0x18, 0x28, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x4d, 0x61, 0x6e, 0x4d, 0x61, 0x63, 0x68,
	0x69, 0x6e, 0x65, 0x53, 0x70, 0x65, 0x63, 0x12, 0x28, 0x0a, 0x0f, 0x45, 0x6e, 0x76, 0x69, 0x72,
	0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x70, 0x65, 0x63, 0x18, 0x32, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0f, 0x45, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x70, 0x65,
	0x63, 0x12, 0x1c, 0x0a, 0x09, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x18, 0x3c,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x22,
	0x45, 0x0a, 0x0d, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x34, 0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20,
	0x2e, 0x61, 0x70, 0x70, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65,
	0x72, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x70, 0x70, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x3e, 0x0a, 0x0e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x65,
	0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x41, 0x70, 0x70, 0x49,
	0x44, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x41, 0x70, 0x70, 0x49, 0x44, 0x12, 0x16,
	0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x22, 0x47, 0x0a, 0x0f, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x65,
	0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x34, 0x0a, 0x04, 0x49, 0x6e, 0x66,
	0x6f, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x70,
	0x70, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x22,
	0x3d, 0x0a, 0x0d, 0x4c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x14, 0x0a, 0x05, 0x41, 0x70, 0x70, 0x49, 0x44, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x41, 0x70, 0x70, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44,
	0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x22, 0x46,
	0x0a, 0x0e, 0x4c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x34, 0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20,
	0x2e, 0x61, 0x70, 0x70, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65,
	0x72, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x70, 0x70, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x3e, 0x0a, 0x0e, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73,
	0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x41, 0x70, 0x70, 0x49,
	0x44, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x41, 0x70, 0x70, 0x49, 0x44, 0x12, 0x16,
	0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x22, 0x47, 0x0a, 0x0f, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73,
	0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x34, 0x0a, 0x04, 0x49, 0x6e, 0x66,
	0x6f, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x70,
	0x70, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x22,
	0xa2, 0x01, 0x0a, 0x0c, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79,
	0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44,
	0x12, 0x14, 0x0a, 0x05, 0x41, 0x70, 0x70, 0x49, 0x44, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x41, 0x70, 0x70, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44,
	0x18, 0x1e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x1a,
	0x0a, 0x08, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x50, 0x18, 0x28, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x50, 0x12, 0x1c, 0x0a, 0x09, 0x55, 0x73,
	0x65, 0x72, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x18, 0x32, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x55,
	0x73, 0x65, 0x72, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x41, 0x74, 0x18, 0x3c, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x41, 0x74, 0x22, 0x78, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x4c, 0x6f, 0x67, 0x69, 0x6e,
	0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x14, 0x0a, 0x05, 0x41, 0x70, 0x70, 0x49, 0x44, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x41, 0x70, 0x70, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44,
	0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x2e,
	0x0a, 0x08, 0x50, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x12, 0x2e, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x67, 0x65,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x08, 0x50, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x51,
	0x0a, 0x19, 0x47, 0x65, 0x74, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72,
	0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x34, 0x0a, 0x05, 0x49,
	0x6e, 0x66, 0x6f, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x6c, 0x6f, 0x67,
	0x69, 0x6e, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x6f,
	0x67, 0x69, 0x6e, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x05, 0x49, 0x6e, 0x66, 0x6f,
	0x73, 0x32, 0x88, 0x05, 0x0a, 0x0c, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x47, 0x61, 0x74, 0x65, 0x77,
	0x61, 0x79, 0x12, 0x51, 0x0a, 0x07, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x19, 0x2e, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2e, 0x76, 0x31,
	0x2e, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x13, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0d, 0x22, 0x08, 0x2f, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x3a, 0x01, 0x2a, 0x12, 0x5e, 0x0a, 0x05, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x1e,
	0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x76,
	0x31, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f,
	0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x76,
	0x31, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x14, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0e, 0x22, 0x09, 0x2f, 0x76, 0x31, 0x2f, 0x6c, 0x6f, 0x67,
	0x69, 0x6e, 0x3a, 0x01, 0x2a, 0x12, 0x66, 0x0a, 0x07, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x65, 0x64,
	0x12, 0x20, 0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79,
	0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x21, 0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77,
	0x61, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x65, 0x64, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x16, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x10, 0x22, 0x0b, 0x2f,
	0x76, 0x31, 0x2f, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x65, 0x64, 0x3a, 0x01, 0x2a, 0x12, 0x62, 0x0a,
	0x06, 0x4c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x12, 0x1f, 0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x2e,
	0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x6f, 0x67, 0x6f, 0x75,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x6e,
	0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x6f, 0x67, 0x6f,
	0x75, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x15, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x0f, 0x22, 0x0a, 0x2f, 0x76, 0x31, 0x2f, 0x6c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x3a, 0x01,
	0x2a, 0x12, 0x66, 0x0a, 0x07, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x12, 0x20, 0x2e, 0x6c,
	0x6f, 0x67, 0x69, 0x6e, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x76, 0x31, 0x2e,
	0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21,
	0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x76,
	0x31, 0x2e, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x16, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x10, 0x22, 0x0b, 0x2f, 0x76, 0x31, 0x2f, 0x72,
	0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x3a, 0x01, 0x2a, 0x12, 0x90, 0x01, 0x0a, 0x11, 0x47, 0x65,
	0x74, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x12,
	0x2a, 0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e,
	0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x48, 0x69, 0x73, 0x74, 0x6f,
	0x72, 0x69, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2b, 0x2e, 0x6c, 0x6f,
	0x67, 0x69, 0x6e, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x47,
	0x65, 0x74, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x69, 0x65, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x22, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1c,
	0x22, 0x17, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x65, 0x74, 0x2f, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x2f,
	0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x3a, 0x01, 0x2a, 0x42, 0x35, 0x5a, 0x33,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4e, 0x70, 0x6f, 0x6f, 0x6c,
	0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x2f, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x67, 0x61, 0x74, 0x65,
	0x77, 0x61, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_npool_logingateway_logingateway_proto_rawDescOnce sync.Once
	file_npool_logingateway_logingateway_proto_rawDescData = file_npool_logingateway_logingateway_proto_rawDesc
)

func file_npool_logingateway_logingateway_proto_rawDescGZIP() []byte {
	file_npool_logingateway_logingateway_proto_rawDescOnce.Do(func() {
		file_npool_logingateway_logingateway_proto_rawDescData = protoimpl.X.CompressGZIP(file_npool_logingateway_logingateway_proto_rawDescData)
	})
	return file_npool_logingateway_logingateway_proto_rawDescData
}

var file_npool_logingateway_logingateway_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_npool_logingateway_logingateway_proto_goTypes = []interface{}{
	(*LoginRequest)(nil),              // 0: login.gateway.v1.LoginRequest
	(*LoginResponse)(nil),             // 1: login.gateway.v1.LoginResponse
	(*LoginedRequest)(nil),            // 2: login.gateway.v1.LoginedRequest
	(*LoginedResponse)(nil),           // 3: login.gateway.v1.LoginedResponse
	(*LogoutRequest)(nil),             // 4: login.gateway.v1.LogoutRequest
	(*LogoutResponse)(nil),            // 5: login.gateway.v1.LogoutResponse
	(*RefreshRequest)(nil),            // 6: login.gateway.v1.RefreshRequest
	(*RefreshResponse)(nil),           // 7: login.gateway.v1.RefreshResponse
	(*LoginHistory)(nil),              // 8: login.gateway.v1.LoginHistory
	(*GetLoginHistoriesRequest)(nil),  // 9: login.gateway.v1.GetLoginHistoriesRequest
	(*GetLoginHistoriesResponse)(nil), // 10: login.gateway.v1.GetLoginHistoriesResponse
	(*appusermgr.AppUserInfo)(nil),    // 11: app.user.manager.v1.AppUserInfo
	(*npool.PageInfo)(nil),            // 12: npool.v1.PageInfo
	(*emptypb.Empty)(nil),             // 13: google.protobuf.Empty
	(*npool.VersionResponse)(nil),     // 14: npool.v1.VersionResponse
}
var file_npool_logingateway_logingateway_proto_depIdxs = []int32{
	11, // 0: login.gateway.v1.LoginResponse.Info:type_name -> app.user.manager.v1.AppUserInfo
	11, // 1: login.gateway.v1.LoginedResponse.Info:type_name -> app.user.manager.v1.AppUserInfo
	11, // 2: login.gateway.v1.LogoutResponse.Info:type_name -> app.user.manager.v1.AppUserInfo
	11, // 3: login.gateway.v1.RefreshResponse.Info:type_name -> app.user.manager.v1.AppUserInfo
	12, // 4: login.gateway.v1.GetLoginHistoriesRequest.PageInfo:type_name -> npool.v1.PageInfo
	8,  // 5: login.gateway.v1.GetLoginHistoriesResponse.Infos:type_name -> login.gateway.v1.LoginHistory
	13, // 6: login.gateway.v1.LoginGateway.Version:input_type -> google.protobuf.Empty
	0,  // 7: login.gateway.v1.LoginGateway.Login:input_type -> login.gateway.v1.LoginRequest
	2,  // 8: login.gateway.v1.LoginGateway.Logined:input_type -> login.gateway.v1.LoginedRequest
	4,  // 9: login.gateway.v1.LoginGateway.Logout:input_type -> login.gateway.v1.LogoutRequest
	6,  // 10: login.gateway.v1.LoginGateway.Refresh:input_type -> login.gateway.v1.RefreshRequest
	9,  // 11: login.gateway.v1.LoginGateway.GetLoginHistories:input_type -> login.gateway.v1.GetLoginHistoriesRequest
	14, // 12: login.gateway.v1.LoginGateway.Version:output_type -> npool.v1.VersionResponse
	1,  // 13: login.gateway.v1.LoginGateway.Login:output_type -> login.gateway.v1.LoginResponse
	3,  // 14: login.gateway.v1.LoginGateway.Logined:output_type -> login.gateway.v1.LoginedResponse
	5,  // 15: login.gateway.v1.LoginGateway.Logout:output_type -> login.gateway.v1.LogoutResponse
	7,  // 16: login.gateway.v1.LoginGateway.Refresh:output_type -> login.gateway.v1.RefreshResponse
	10, // 17: login.gateway.v1.LoginGateway.GetLoginHistories:output_type -> login.gateway.v1.GetLoginHistoriesResponse
	12, // [12:18] is the sub-list for method output_type
	6,  // [6:12] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_npool_logingateway_logingateway_proto_init() }
func file_npool_logingateway_logingateway_proto_init() {
	if File_npool_logingateway_logingateway_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_npool_logingateway_logingateway_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginRequest); i {
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
		file_npool_logingateway_logingateway_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginResponse); i {
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
		file_npool_logingateway_logingateway_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginedRequest); i {
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
		file_npool_logingateway_logingateway_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginedResponse); i {
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
		file_npool_logingateway_logingateway_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogoutRequest); i {
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
		file_npool_logingateway_logingateway_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogoutResponse); i {
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
		file_npool_logingateway_logingateway_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RefreshRequest); i {
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
		file_npool_logingateway_logingateway_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RefreshResponse); i {
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
		file_npool_logingateway_logingateway_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginHistory); i {
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
		file_npool_logingateway_logingateway_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetLoginHistoriesRequest); i {
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
		file_npool_logingateway_logingateway_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetLoginHistoriesResponse); i {
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
			RawDescriptor: file_npool_logingateway_logingateway_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_npool_logingateway_logingateway_proto_goTypes,
		DependencyIndexes: file_npool_logingateway_logingateway_proto_depIdxs,
		MessageInfos:      file_npool_logingateway_logingateway_proto_msgTypes,
	}.Build()
	File_npool_logingateway_logingateway_proto = out.File
	file_npool_logingateway_logingateway_proto_rawDesc = nil
	file_npool_logingateway_logingateway_proto_goTypes = nil
	file_npool_logingateway_logingateway_proto_depIdxs = nil
}
