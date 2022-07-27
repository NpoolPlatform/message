// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.18.1
// source: npool/appuser/mw/v1/app/app.proto

package app

import (
	npool "github.com/NpoolPlatform/message/npool"
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

type AppInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                  string   `protobuf:"bytes,20,opt,name=id,proto3" json:"id,omitempty"`
	CreatedBy           string   `protobuf:"bytes,30,opt,name=createdBy,proto3" json:"createdBy,omitempty"`
	Name                string   `protobuf:"bytes,40,opt,name=Name,proto3" json:"Name,omitempty"`
	Logo                string   `protobuf:"bytes,50,opt,name=Logo,proto3" json:"Logo,omitempty"`
	CreatedAt           uint32   `protobuf:"varint,60,opt,name=CreatedAt,proto3" json:"CreatedAt,omitempty"`
	Description         string   `protobuf:"bytes,10,opt,name=Description,proto3" json:"Description,omitempty"`
	IsBanApp            bool     `protobuf:"varint,80,opt,name=IsBanApp,proto3" json:"IsBanApp,omitempty"`
	BanAppMessage       string   `protobuf:"bytes,90,opt,name=BanAppMessage,proto3" json:"BanAppMessage,omitempty"`
	SignupMethods       []string `protobuf:"bytes,100,rep,name=SignupMethods,proto3" json:"SignupMethods,omitempty"`
	ExternSigninMethods []string `protobuf:"bytes,110,rep,name=ExternSigninMethods,proto3" json:"ExternSigninMethods,omitempty"`
	RecaptchaMethod     string   `protobuf:"bytes,120,opt,name=RecaptchaMethod,proto3" json:"RecaptchaMethod,omitempty"`
	KycEnable           bool     `protobuf:"varint,130,opt,name=KycEnable,proto3" json:"KycEnable,omitempty"`
	SigninVerifyEnable  bool     `protobuf:"varint,140,opt,name=SigninVerifyEnable,proto3" json:"SigninVerifyEnable,omitempty"`
	InvitationCodeMust  bool     `protobuf:"varint,150,opt,name=InvitationCodeMust,proto3" json:"InvitationCodeMust,omitempty"`
}

func (x *AppInfo) Reset() {
	*x = AppInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_appuser_mw_v1_app_app_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AppInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppInfo) ProtoMessage() {}

func (x *AppInfo) ProtoReflect() protoreflect.Message {
	mi := &file_npool_appuser_mw_v1_app_app_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AppInfo.ProtoReflect.Descriptor instead.
func (*AppInfo) Descriptor() ([]byte, []int) {
	return file_npool_appuser_mw_v1_app_app_proto_rawDescGZIP(), []int{0}
}

func (x *AppInfo) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *AppInfo) GetCreatedBy() string {
	if x != nil {
		return x.CreatedBy
	}
	return ""
}

func (x *AppInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AppInfo) GetLogo() string {
	if x != nil {
		return x.Logo
	}
	return ""
}

func (x *AppInfo) GetCreatedAt() uint32 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *AppInfo) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *AppInfo) GetIsBanApp() bool {
	if x != nil {
		return x.IsBanApp
	}
	return false
}

func (x *AppInfo) GetBanAppMessage() string {
	if x != nil {
		return x.BanAppMessage
	}
	return ""
}

func (x *AppInfo) GetSignupMethods() []string {
	if x != nil {
		return x.SignupMethods
	}
	return nil
}

func (x *AppInfo) GetExternSigninMethods() []string {
	if x != nil {
		return x.ExternSigninMethods
	}
	return nil
}

func (x *AppInfo) GetRecaptchaMethod() string {
	if x != nil {
		return x.RecaptchaMethod
	}
	return ""
}

func (x *AppInfo) GetKycEnable() bool {
	if x != nil {
		return x.KycEnable
	}
	return false
}

func (x *AppInfo) GetSigninVerifyEnable() bool {
	if x != nil {
		return x.SigninVerifyEnable
	}
	return false
}

func (x *AppInfo) GetInvitationCodeMust() bool {
	if x != nil {
		return x.InvitationCodeMust
	}
	return false
}

type GetAppInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID string `protobuf:"bytes,10,opt,name=ID,proto3" json:"ID,omitempty"`
}

func (x *GetAppInfoRequest) Reset() {
	*x = GetAppInfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_appuser_mw_v1_app_app_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAppInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAppInfoRequest) ProtoMessage() {}

func (x *GetAppInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_appuser_mw_v1_app_app_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAppInfoRequest.ProtoReflect.Descriptor instead.
func (*GetAppInfoRequest) Descriptor() ([]byte, []int) {
	return file_npool_appuser_mw_v1_app_app_proto_rawDescGZIP(), []int{1}
}

func (x *GetAppInfoRequest) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

type GetAppInfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *AppInfo `protobuf:"bytes,10,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (x *GetAppInfoResponse) Reset() {
	*x = GetAppInfoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_appuser_mw_v1_app_app_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAppInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAppInfoResponse) ProtoMessage() {}

func (x *GetAppInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_appuser_mw_v1_app_app_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAppInfoResponse.ProtoReflect.Descriptor instead.
func (*GetAppInfoResponse) Descriptor() ([]byte, []int) {
	return file_npool_appuser_mw_v1_app_app_proto_rawDescGZIP(), []int{2}
}

func (x *GetAppInfoResponse) GetInfo() *AppInfo {
	if x != nil {
		return x.Info
	}
	return nil
}

type GetAppInfosRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Offset int32 `protobuf:"varint,10,opt,name=Offset,proto3" json:"Offset,omitempty"`
	Limit  int32 `protobuf:"varint,20,opt,name=Limit,proto3" json:"Limit,omitempty"`
}

func (x *GetAppInfosRequest) Reset() {
	*x = GetAppInfosRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_appuser_mw_v1_app_app_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAppInfosRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAppInfosRequest) ProtoMessage() {}

func (x *GetAppInfosRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_appuser_mw_v1_app_app_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAppInfosRequest.ProtoReflect.Descriptor instead.
func (*GetAppInfosRequest) Descriptor() ([]byte, []int) {
	return file_npool_appuser_mw_v1_app_app_proto_rawDescGZIP(), []int{3}
}

func (x *GetAppInfosRequest) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *GetAppInfosRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type GetAppInfosResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Infos []*AppInfo `protobuf:"bytes,10,rep,name=Infos,proto3" json:"Infos,omitempty"`
	Total uint32     `protobuf:"varint,20,opt,name=Total,proto3" json:"Total,omitempty"`
}

func (x *GetAppInfosResponse) Reset() {
	*x = GetAppInfosResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_appuser_mw_v1_app_app_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAppInfosResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAppInfosResponse) ProtoMessage() {}

func (x *GetAppInfosResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_appuser_mw_v1_app_app_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAppInfosResponse.ProtoReflect.Descriptor instead.
func (*GetAppInfosResponse) Descriptor() ([]byte, []int) {
	return file_npool_appuser_mw_v1_app_app_proto_rawDescGZIP(), []int{4}
}

func (x *GetAppInfosResponse) GetInfos() []*AppInfo {
	if x != nil {
		return x.Infos
	}
	return nil
}

func (x *GetAppInfosResponse) GetTotal() uint32 {
	if x != nil {
		return x.Total
	}
	return 0
}

type GetAppInfosByCreatorRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID string `protobuf:"bytes,10,opt,name=UserID,proto3" json:"UserID,omitempty"`
	Offset int32  `protobuf:"varint,20,opt,name=Offset,proto3" json:"Offset,omitempty"`
	Limit  int32  `protobuf:"varint,30,opt,name=Limit,proto3" json:"Limit,omitempty"`
}

func (x *GetAppInfosByCreatorRequest) Reset() {
	*x = GetAppInfosByCreatorRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_appuser_mw_v1_app_app_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAppInfosByCreatorRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAppInfosByCreatorRequest) ProtoMessage() {}

func (x *GetAppInfosByCreatorRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_appuser_mw_v1_app_app_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAppInfosByCreatorRequest.ProtoReflect.Descriptor instead.
func (*GetAppInfosByCreatorRequest) Descriptor() ([]byte, []int) {
	return file_npool_appuser_mw_v1_app_app_proto_rawDescGZIP(), []int{5}
}

func (x *GetAppInfosByCreatorRequest) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

func (x *GetAppInfosByCreatorRequest) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *GetAppInfosByCreatorRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type GetAppInfosByCreatorResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Infos []*AppInfo `protobuf:"bytes,10,rep,name=Infos,proto3" json:"Infos,omitempty"`
	Total uint32     `protobuf:"varint,20,opt,name=Total,proto3" json:"Total,omitempty"`
}

func (x *GetAppInfosByCreatorResponse) Reset() {
	*x = GetAppInfosByCreatorResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_appuser_mw_v1_app_app_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAppInfosByCreatorResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAppInfosByCreatorResponse) ProtoMessage() {}

func (x *GetAppInfosByCreatorResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_appuser_mw_v1_app_app_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAppInfosByCreatorResponse.ProtoReflect.Descriptor instead.
func (*GetAppInfosByCreatorResponse) Descriptor() ([]byte, []int) {
	return file_npool_appuser_mw_v1_app_app_proto_rawDescGZIP(), []int{6}
}

func (x *GetAppInfosByCreatorResponse) GetInfos() []*AppInfo {
	if x != nil {
		return x.Infos
	}
	return nil
}

func (x *GetAppInfosByCreatorResponse) GetTotal() uint32 {
	if x != nil {
		return x.Total
	}
	return 0
}

var File_npool_appuser_mw_v1_app_app_proto protoreflect.FileDescriptor

var file_npool_appuser_mw_v1_app_app_proto_rawDesc = []byte{
	0x0a, 0x21, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x61, 0x70, 0x70, 0x75, 0x73, 0x65, 0x72, 0x2f,
	0x6d, 0x77, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x70, 0x70, 0x2f, 0x61, 0x70, 0x70, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x1a, 0x61, 0x70, 0x70, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x6d, 0x69,
	0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x76, 0x32, 0x1a,
	0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x6e, 0x70,
	0x6f, 0x6f, 0x6c, 0x2f, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xe4, 0x03, 0x0a, 0x07, 0x41, 0x70, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d,
	0x65, 0x18, 0x28, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x4c, 0x6f, 0x67, 0x6f, 0x18, 0x32, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4c, 0x6f, 0x67,
	0x6f, 0x12, 0x1c, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x3c,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12,
	0x20, 0x0a, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x49, 0x73, 0x42, 0x61, 0x6e, 0x41, 0x70, 0x70, 0x18, 0x50, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x08, 0x49, 0x73, 0x42, 0x61, 0x6e, 0x41, 0x70, 0x70, 0x12, 0x24, 0x0a,
	0x0d, 0x42, 0x61, 0x6e, 0x41, 0x70, 0x70, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x5a,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x42, 0x61, 0x6e, 0x41, 0x70, 0x70, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x53, 0x69, 0x67, 0x6e, 0x75, 0x70, 0x4d, 0x65, 0x74,
	0x68, 0x6f, 0x64, 0x73, 0x18, 0x64, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0d, 0x53, 0x69, 0x67, 0x6e,
	0x75, 0x70, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x73, 0x12, 0x30, 0x0a, 0x13, 0x45, 0x78, 0x74,
	0x65, 0x72, 0x6e, 0x53, 0x69, 0x67, 0x6e, 0x69, 0x6e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x73,
	0x18, 0x6e, 0x20, 0x03, 0x28, 0x09, 0x52, 0x13, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x53, 0x69,
	0x67, 0x6e, 0x69, 0x6e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x73, 0x12, 0x28, 0x0a, 0x0f, 0x52,
	0x65, 0x63, 0x61, 0x70, 0x74, 0x63, 0x68, 0x61, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x78,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x52, 0x65, 0x63, 0x61, 0x70, 0x74, 0x63, 0x68, 0x61, 0x4d,
	0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x1d, 0x0a, 0x09, 0x4b, 0x79, 0x63, 0x45, 0x6e, 0x61, 0x62,
	0x6c, 0x65, 0x18, 0x82, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x4b, 0x79, 0x63, 0x45, 0x6e,
	0x61, 0x62, 0x6c, 0x65, 0x12, 0x2f, 0x0a, 0x12, 0x53, 0x69, 0x67, 0x6e, 0x69, 0x6e, 0x56, 0x65,
	0x72, 0x69, 0x66, 0x79, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x8c, 0x01, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x12, 0x53, 0x69, 0x67, 0x6e, 0x69, 0x6e, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x45,
	0x6e, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x2f, 0x0a, 0x12, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x64, 0x65, 0x4d, 0x75, 0x73, 0x74, 0x18, 0x96, 0x01, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x12, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f,
	0x64, 0x65, 0x4d, 0x75, 0x73, 0x74, 0x22, 0x23, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x41, 0x70, 0x70,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49,
	0x44, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x22, 0x4d, 0x0a, 0x12, 0x47,
	0x65, 0x74, 0x41, 0x70, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x37, 0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x23, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c,
	0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x76, 0x32, 0x2e, 0x41, 0x70, 0x70,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x42, 0x0a, 0x12, 0x47, 0x65,
	0x74, 0x41, 0x70, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x16, 0x0a, 0x06, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x06, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x4c, 0x69, 0x6d, 0x69,
	0x74, 0x18, 0x14, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x22, 0x66,
	0x0a, 0x13, 0x47, 0x65, 0x74, 0x41, 0x70, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x39, 0x0a, 0x05, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x18, 0x0a,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e,
	0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x76,
	0x32, 0x2e, 0x41, 0x70, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x05, 0x49, 0x6e, 0x66, 0x6f, 0x73,
	0x12, 0x14, 0x0a, 0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x14, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x22, 0x63, 0x0a, 0x1b, 0x47, 0x65, 0x74, 0x41, 0x70, 0x70,
	0x49, 0x6e, 0x66, 0x6f, 0x73, 0x42, 0x79, 0x43, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x16, 0x0a,
	0x06, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x14, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x4f,
	0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x1e,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x22, 0x6f, 0x0a, 0x1c, 0x47,
	0x65, 0x74, 0x41, 0x70, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x42, 0x79, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x39, 0x0a, 0x05, 0x49,
	0x6e, 0x66, 0x6f, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x61, 0x70, 0x70,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65,
	0x2e, 0x61, 0x70, 0x70, 0x2e, 0x76, 0x32, 0x2e, 0x41, 0x70, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x05, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x18,
	0x14, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x32, 0xc5, 0x03, 0x0a,
	0x14, 0x41, 0x70, 0x70, 0x55, 0x73, 0x65, 0x72, 0x4d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61,
	0x72, 0x65, 0x41, 0x70, 0x70, 0x12, 0x3e, 0x0a, 0x07, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x19, 0x2e, 0x6e, 0x70, 0x6f, 0x6f, 0x6c,
	0x2e, 0x76, 0x31, 0x2e, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x6d, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x41, 0x70, 0x70, 0x49,
	0x6e, 0x66, 0x6f, 0x12, 0x2d, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x6d,
	0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x76, 0x32,
	0x2e, 0x47, 0x65, 0x74, 0x41, 0x70, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x2e, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x6d, 0x69,
	0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x76, 0x32, 0x2e,
	0x47, 0x65, 0x74, 0x41, 0x70, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x70, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x41, 0x70, 0x70, 0x49, 0x6e,
	0x66, 0x6f, 0x73, 0x12, 0x2e, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x6d,
	0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x76, 0x32,
	0x2e, 0x47, 0x65, 0x74, 0x41, 0x70, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x2f, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x6d,
	0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x76, 0x32,
	0x2e, 0x47, 0x65, 0x74, 0x41, 0x70, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x8b, 0x01, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x41, 0x70,
	0x70, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x42, 0x79, 0x43, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x12,
	0x37, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c,
	0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x76, 0x32, 0x2e, 0x47, 0x65, 0x74,
	0x41, 0x70, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x42, 0x79, 0x43, 0x72, 0x65, 0x61, 0x74, 0x6f,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x38, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x75,
	0x73, 0x65, 0x72, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x61,
	0x70, 0x70, 0x2e, 0x76, 0x32, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x70, 0x70, 0x49, 0x6e, 0x66, 0x6f,
	0x73, 0x42, 0x79, 0x43, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x42, 0x36, 0x5a, 0x34, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x4e, 0x70, 0x6f, 0x6f, 0x6c, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d,
	0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x61,
	0x70, 0x70, 0x75, 0x73, 0x65, 0x72, 0x6d, 0x77, 0x2f, 0x61, 0x70, 0x70, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_npool_appuser_mw_v1_app_app_proto_rawDescOnce sync.Once
	file_npool_appuser_mw_v1_app_app_proto_rawDescData = file_npool_appuser_mw_v1_app_app_proto_rawDesc
)

func file_npool_appuser_mw_v1_app_app_proto_rawDescGZIP() []byte {
	file_npool_appuser_mw_v1_app_app_proto_rawDescOnce.Do(func() {
		file_npool_appuser_mw_v1_app_app_proto_rawDescData = protoimpl.X.CompressGZIP(file_npool_appuser_mw_v1_app_app_proto_rawDescData)
	})
	return file_npool_appuser_mw_v1_app_app_proto_rawDescData
}

var file_npool_appuser_mw_v1_app_app_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_npool_appuser_mw_v1_app_app_proto_goTypes = []interface{}{
	(*AppInfo)(nil),                      // 0: app.user.middleware.app.v2.AppInfo
	(*GetAppInfoRequest)(nil),            // 1: app.user.middleware.app.v2.GetAppInfoRequest
	(*GetAppInfoResponse)(nil),           // 2: app.user.middleware.app.v2.GetAppInfoResponse
	(*GetAppInfosRequest)(nil),           // 3: app.user.middleware.app.v2.GetAppInfosRequest
	(*GetAppInfosResponse)(nil),          // 4: app.user.middleware.app.v2.GetAppInfosResponse
	(*GetAppInfosByCreatorRequest)(nil),  // 5: app.user.middleware.app.v2.GetAppInfosByCreatorRequest
	(*GetAppInfosByCreatorResponse)(nil), // 6: app.user.middleware.app.v2.GetAppInfosByCreatorResponse
	(*emptypb.Empty)(nil),                // 7: google.protobuf.Empty
	(*npool.VersionResponse)(nil),        // 8: npool.v1.VersionResponse
}
var file_npool_appuser_mw_v1_app_app_proto_depIdxs = []int32{
	0, // 0: app.user.middleware.app.v2.GetAppInfoResponse.Info:type_name -> app.user.middleware.app.v2.AppInfo
	0, // 1: app.user.middleware.app.v2.GetAppInfosResponse.Infos:type_name -> app.user.middleware.app.v2.AppInfo
	0, // 2: app.user.middleware.app.v2.GetAppInfosByCreatorResponse.Infos:type_name -> app.user.middleware.app.v2.AppInfo
	7, // 3: app.user.middleware.app.v2.AppUserMiddlewareApp.Version:input_type -> google.protobuf.Empty
	1, // 4: app.user.middleware.app.v2.AppUserMiddlewareApp.GetAppInfo:input_type -> app.user.middleware.app.v2.GetAppInfoRequest
	3, // 5: app.user.middleware.app.v2.AppUserMiddlewareApp.GetAppInfos:input_type -> app.user.middleware.app.v2.GetAppInfosRequest
	5, // 6: app.user.middleware.app.v2.AppUserMiddlewareApp.GetAppInfosByCreator:input_type -> app.user.middleware.app.v2.GetAppInfosByCreatorRequest
	8, // 7: app.user.middleware.app.v2.AppUserMiddlewareApp.Version:output_type -> npool.v1.VersionResponse
	2, // 8: app.user.middleware.app.v2.AppUserMiddlewareApp.GetAppInfo:output_type -> app.user.middleware.app.v2.GetAppInfoResponse
	4, // 9: app.user.middleware.app.v2.AppUserMiddlewareApp.GetAppInfos:output_type -> app.user.middleware.app.v2.GetAppInfosResponse
	6, // 10: app.user.middleware.app.v2.AppUserMiddlewareApp.GetAppInfosByCreator:output_type -> app.user.middleware.app.v2.GetAppInfosByCreatorResponse
	7, // [7:11] is the sub-list for method output_type
	3, // [3:7] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_npool_appuser_mw_v1_app_app_proto_init() }
func file_npool_appuser_mw_v1_app_app_proto_init() {
	if File_npool_appuser_mw_v1_app_app_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_npool_appuser_mw_v1_app_app_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AppInfo); i {
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
		file_npool_appuser_mw_v1_app_app_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAppInfoRequest); i {
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
		file_npool_appuser_mw_v1_app_app_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAppInfoResponse); i {
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
		file_npool_appuser_mw_v1_app_app_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAppInfosRequest); i {
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
		file_npool_appuser_mw_v1_app_app_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAppInfosResponse); i {
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
		file_npool_appuser_mw_v1_app_app_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAppInfosByCreatorRequest); i {
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
		file_npool_appuser_mw_v1_app_app_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAppInfosByCreatorResponse); i {
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
			RawDescriptor: file_npool_appuser_mw_v1_app_app_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_npool_appuser_mw_v1_app_app_proto_goTypes,
		DependencyIndexes: file_npool_appuser_mw_v1_app_app_proto_depIdxs,
		MessageInfos:      file_npool_appuser_mw_v1_app_app_proto_msgTypes,
	}.Build()
	File_npool_appuser_mw_v1_app_app_proto = out.File
	file_npool_appuser_mw_v1_app_app_proto_rawDesc = nil
	file_npool_appuser_mw_v1_app_app_proto_goTypes = nil
	file_npool_appuser_mw_v1_app_app_proto_depIdxs = nil
}
