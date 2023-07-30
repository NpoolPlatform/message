// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.18.1
// source: npool/third/mw/v1/oauth/oauth.proto

package oauth

import (
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

type ThirdUserInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID        string `protobuf:"bytes,10,opt,name=ID,proto3" json:"ID,omitempty"`
	Name      string `protobuf:"bytes,20,opt,name=Name,proto3" json:"Name,omitempty"`
	Login     string `protobuf:"bytes,30,opt,name=Login,proto3" json:"Login,omitempty"`
	AvatarURL string `protobuf:"bytes,40,opt,name=AvatarURL,proto3" json:"AvatarURL,omitempty"`
	Message   string `protobuf:"bytes,50,opt,name=Message,proto3" json:"Message,omitempty"`
}

func (x *ThirdUserInfo) Reset() {
	*x = ThirdUserInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_third_mw_v1_oauth_oauth_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ThirdUserInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ThirdUserInfo) ProtoMessage() {}

func (x *ThirdUserInfo) ProtoReflect() protoreflect.Message {
	mi := &file_npool_third_mw_v1_oauth_oauth_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ThirdUserInfo.ProtoReflect.Descriptor instead.
func (*ThirdUserInfo) Descriptor() ([]byte, []int) {
	return file_npool_third_mw_v1_oauth_oauth_proto_rawDescGZIP(), []int{0}
}

func (x *ThirdUserInfo) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *ThirdUserInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ThirdUserInfo) GetLogin() string {
	if x != nil {
		return x.Login
	}
	return ""
}

func (x *ThirdUserInfo) GetAvatarURL() string {
	if x != nil {
		return x.AvatarURL
	}
	return ""
}

func (x *ThirdUserInfo) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type AccessTokenInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccessToken        string `protobuf:"bytes,10,opt,name=AccessToken,proto3" json:"AccessToken,omitempty"`
	RefreshAccessToken string `protobuf:"bytes,20,opt,name=RefreshAccessToken,proto3" json:"RefreshAccessToken,omitempty"`
	TokenType          string `protobuf:"bytes,30,opt,name=TokenType,proto3" json:"TokenType,omitempty"`
	Scope              string `protobuf:"bytes,40,opt,name=Scope,proto3" json:"Scope,omitempty"`
}

func (x *AccessTokenInfo) Reset() {
	*x = AccessTokenInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_third_mw_v1_oauth_oauth_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccessTokenInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccessTokenInfo) ProtoMessage() {}

func (x *AccessTokenInfo) ProtoReflect() protoreflect.Message {
	mi := &file_npool_third_mw_v1_oauth_oauth_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccessTokenInfo.ProtoReflect.Descriptor instead.
func (*AccessTokenInfo) Descriptor() ([]byte, []int) {
	return file_npool_third_mw_v1_oauth_oauth_proto_rawDescGZIP(), []int{1}
}

func (x *AccessTokenInfo) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

func (x *AccessTokenInfo) GetRefreshAccessToken() string {
	if x != nil {
		return x.RefreshAccessToken
	}
	return ""
}

func (x *AccessTokenInfo) GetTokenType() string {
	if x != nil {
		return x.TokenType
	}
	return ""
}

func (x *AccessTokenInfo) GetScope() string {
	if x != nil {
		return x.Scope
	}
	return ""
}

type GetOAuthAccessTokenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClientName   v1.SignMethod `protobuf:"varint,10,opt,name=ClientName,proto3,enum=basetypes.v1.SignMethod" json:"ClientName,omitempty"`
	Code         string        `protobuf:"bytes,20,opt,name=Code,proto3" json:"Code,omitempty"`
	ClientID     string        `protobuf:"bytes,30,opt,name=ClientID,proto3" json:"ClientID,omitempty"`
	ClientSecret string        `protobuf:"bytes,40,opt,name=ClientSecret,proto3" json:"ClientSecret,omitempty"`
}

func (x *GetOAuthAccessTokenRequest) Reset() {
	*x = GetOAuthAccessTokenRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_third_mw_v1_oauth_oauth_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetOAuthAccessTokenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOAuthAccessTokenRequest) ProtoMessage() {}

func (x *GetOAuthAccessTokenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_third_mw_v1_oauth_oauth_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOAuthAccessTokenRequest.ProtoReflect.Descriptor instead.
func (*GetOAuthAccessTokenRequest) Descriptor() ([]byte, []int) {
	return file_npool_third_mw_v1_oauth_oauth_proto_rawDescGZIP(), []int{2}
}

func (x *GetOAuthAccessTokenRequest) GetClientName() v1.SignMethod {
	if x != nil {
		return x.ClientName
	}
	return v1.SignMethod(0)
}

func (x *GetOAuthAccessTokenRequest) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *GetOAuthAccessTokenRequest) GetClientID() string {
	if x != nil {
		return x.ClientID
	}
	return ""
}

func (x *GetOAuthAccessTokenRequest) GetClientSecret() string {
	if x != nil {
		return x.ClientSecret
	}
	return ""
}

type GetOAuthAccessTokenResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *AccessTokenInfo `protobuf:"bytes,10,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (x *GetOAuthAccessTokenResponse) Reset() {
	*x = GetOAuthAccessTokenResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_third_mw_v1_oauth_oauth_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetOAuthAccessTokenResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOAuthAccessTokenResponse) ProtoMessage() {}

func (x *GetOAuthAccessTokenResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_third_mw_v1_oauth_oauth_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOAuthAccessTokenResponse.ProtoReflect.Descriptor instead.
func (*GetOAuthAccessTokenResponse) Descriptor() ([]byte, []int) {
	return file_npool_third_mw_v1_oauth_oauth_proto_rawDescGZIP(), []int{3}
}

func (x *GetOAuthAccessTokenResponse) GetInfo() *AccessTokenInfo {
	if x != nil {
		return x.Info
	}
	return nil
}

type GetOAuthUserInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClientName  v1.SignMethod `protobuf:"varint,10,opt,name=ClientName,proto3,enum=basetypes.v1.SignMethod" json:"ClientName,omitempty"`
	AccessToken string        `protobuf:"bytes,20,opt,name=AccessToken,proto3" json:"AccessToken,omitempty"`
}

func (x *GetOAuthUserInfoRequest) Reset() {
	*x = GetOAuthUserInfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_third_mw_v1_oauth_oauth_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetOAuthUserInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOAuthUserInfoRequest) ProtoMessage() {}

func (x *GetOAuthUserInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_third_mw_v1_oauth_oauth_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOAuthUserInfoRequest.ProtoReflect.Descriptor instead.
func (*GetOAuthUserInfoRequest) Descriptor() ([]byte, []int) {
	return file_npool_third_mw_v1_oauth_oauth_proto_rawDescGZIP(), []int{4}
}

func (x *GetOAuthUserInfoRequest) GetClientName() v1.SignMethod {
	if x != nil {
		return x.ClientName
	}
	return v1.SignMethod(0)
}

func (x *GetOAuthUserInfoRequest) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

type GetOAuthUserInfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *ThirdUserInfo `protobuf:"bytes,10,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (x *GetOAuthUserInfoResponse) Reset() {
	*x = GetOAuthUserInfoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_third_mw_v1_oauth_oauth_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetOAuthUserInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOAuthUserInfoResponse) ProtoMessage() {}

func (x *GetOAuthUserInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_third_mw_v1_oauth_oauth_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOAuthUserInfoResponse.ProtoReflect.Descriptor instead.
func (*GetOAuthUserInfoResponse) Descriptor() ([]byte, []int) {
	return file_npool_third_mw_v1_oauth_oauth_proto_rawDescGZIP(), []int{5}
}

func (x *GetOAuthUserInfoResponse) GetInfo() *ThirdUserInfo {
	if x != nil {
		return x.Info
	}
	return nil
}

var File_npool_third_mw_v1_oauth_oauth_proto protoreflect.FileDescriptor

var file_npool_third_mw_v1_oauth_oauth_proto_rawDesc = []byte{
	0x0a, 0x23, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x74, 0x68, 0x69, 0x72, 0x64, 0x2f, 0x6d, 0x77,
	0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x19, 0x74, 0x68, 0x69, 0x72, 0x64, 0x2e, 0x6d, 0x69, 0x64,
	0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x76, 0x31,
	0x1a, 0x23, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x81, 0x01, 0x0a, 0x0d, 0x54, 0x68, 0x69, 0x72, 0x64, 0x55,
	0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18,
	0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x4c,
	0x6f, 0x67, 0x69, 0x6e, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x4c, 0x6f, 0x67, 0x69,
	0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x55, 0x52, 0x4c, 0x18, 0x28,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x55, 0x52, 0x4c, 0x12,
	0x18, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x32, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x97, 0x01, 0x0a, 0x0f, 0x41, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x20, 0x0a,
	0x0b, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12,
	0x2e, 0x0a, 0x12, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x52, 0x65, 0x66,
	0x72, 0x65, 0x73, 0x68, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12,
	0x1c, 0x0a, 0x09, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x18, 0x1e, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x18, 0x28, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x53, 0x63,
	0x6f, 0x70, 0x65, 0x22, 0xaa, 0x01, 0x0a, 0x1a, 0x47, 0x65, 0x74, 0x4f, 0x41, 0x75, 0x74, 0x68,
	0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x38, 0x0a, 0x0a, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x18, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70,
	0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64,
	0x52, 0x0a, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x43, 0x6f, 0x64, 0x65, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x43, 0x6f, 0x64, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x1e, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x44, 0x12, 0x22, 0x0a, 0x0c,
	0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x18, 0x28, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0c, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74,
	0x22, 0x5d, 0x0a, 0x1b, 0x47, 0x65, 0x74, 0x4f, 0x41, 0x75, 0x74, 0x68, 0x41, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x3e, 0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e,
	0x74, 0x68, 0x69, 0x72, 0x64, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65,
	0x2e, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x22,
	0x75, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x4f, 0x41, 0x75, 0x74, 0x68, 0x55, 0x73, 0x65, 0x72, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x38, 0x0a, 0x0a, 0x43, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x18,
	0x2e, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x69,
	0x67, 0x6e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x52, 0x0a, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x41, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x58, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x4f, 0x41, 0x75,
	0x74, 0x68, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x3c, 0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x28, 0x2e, 0x74, 0x68, 0x69, 0x72, 0x64, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77,
	0x61, 0x72, 0x65, 0x2e, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x68, 0x69,
	0x72, 0x64, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x49, 0x6e, 0x66, 0x6f,
	0x32, 0x94, 0x02, 0x0a, 0x0a, 0x4d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x12,
	0x86, 0x01, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x4f, 0x41, 0x75, 0x74, 0x68, 0x41, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x35, 0x2e, 0x74, 0x68, 0x69, 0x72, 0x64, 0x2e,
	0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x6f, 0x61, 0x75, 0x74, 0x68,
	0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4f, 0x41, 0x75, 0x74, 0x68, 0x41, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x36,
	0x2e, 0x74, 0x68, 0x69, 0x72, 0x64, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72,
	0x65, 0x2e, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4f, 0x41,
	0x75, 0x74, 0x68, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x7d, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x4f,
	0x41, 0x75, 0x74, 0x68, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x32, 0x2e, 0x74,
	0x68, 0x69, 0x72, 0x64, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e,
	0x6f, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4f, 0x41, 0x75, 0x74,
	0x68, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x33, 0x2e, 0x74, 0x68, 0x69, 0x72, 0x64, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77,
	0x61, 0x72, 0x65, 0x2e, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74,
	0x4f, 0x41, 0x75, 0x74, 0x68, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x3a, 0x5a, 0x38, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4e, 0x70, 0x6f, 0x6f, 0x6c, 0x50, 0x6c, 0x61, 0x74, 0x66,
	0x6f, 0x72, 0x6d, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x6e, 0x70, 0x6f, 0x6f,
	0x6c, 0x2f, 0x74, 0x68, 0x69, 0x72, 0x64, 0x2f, 0x6d, 0x77, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x61,
	0x75, 0x74, 0x68, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_npool_third_mw_v1_oauth_oauth_proto_rawDescOnce sync.Once
	file_npool_third_mw_v1_oauth_oauth_proto_rawDescData = file_npool_third_mw_v1_oauth_oauth_proto_rawDesc
)

func file_npool_third_mw_v1_oauth_oauth_proto_rawDescGZIP() []byte {
	file_npool_third_mw_v1_oauth_oauth_proto_rawDescOnce.Do(func() {
		file_npool_third_mw_v1_oauth_oauth_proto_rawDescData = protoimpl.X.CompressGZIP(file_npool_third_mw_v1_oauth_oauth_proto_rawDescData)
	})
	return file_npool_third_mw_v1_oauth_oauth_proto_rawDescData
}

var file_npool_third_mw_v1_oauth_oauth_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_npool_third_mw_v1_oauth_oauth_proto_goTypes = []interface{}{
	(*ThirdUserInfo)(nil),               // 0: third.middleware.oauth.v1.ThirdUserInfo
	(*AccessTokenInfo)(nil),             // 1: third.middleware.oauth.v1.AccessTokenInfo
	(*GetOAuthAccessTokenRequest)(nil),  // 2: third.middleware.oauth.v1.GetOAuthAccessTokenRequest
	(*GetOAuthAccessTokenResponse)(nil), // 3: third.middleware.oauth.v1.GetOAuthAccessTokenResponse
	(*GetOAuthUserInfoRequest)(nil),     // 4: third.middleware.oauth.v1.GetOAuthUserInfoRequest
	(*GetOAuthUserInfoResponse)(nil),    // 5: third.middleware.oauth.v1.GetOAuthUserInfoResponse
	(v1.SignMethod)(0),                  // 6: basetypes.v1.SignMethod
}
var file_npool_third_mw_v1_oauth_oauth_proto_depIdxs = []int32{
	6, // 0: third.middleware.oauth.v1.GetOAuthAccessTokenRequest.ClientName:type_name -> basetypes.v1.SignMethod
	1, // 1: third.middleware.oauth.v1.GetOAuthAccessTokenResponse.Info:type_name -> third.middleware.oauth.v1.AccessTokenInfo
	6, // 2: third.middleware.oauth.v1.GetOAuthUserInfoRequest.ClientName:type_name -> basetypes.v1.SignMethod
	0, // 3: third.middleware.oauth.v1.GetOAuthUserInfoResponse.Info:type_name -> third.middleware.oauth.v1.ThirdUserInfo
	2, // 4: third.middleware.oauth.v1.Middleware.GetOAuthAccessToken:input_type -> third.middleware.oauth.v1.GetOAuthAccessTokenRequest
	4, // 5: third.middleware.oauth.v1.Middleware.GetOAuthUserInfo:input_type -> third.middleware.oauth.v1.GetOAuthUserInfoRequest
	3, // 6: third.middleware.oauth.v1.Middleware.GetOAuthAccessToken:output_type -> third.middleware.oauth.v1.GetOAuthAccessTokenResponse
	5, // 7: third.middleware.oauth.v1.Middleware.GetOAuthUserInfo:output_type -> third.middleware.oauth.v1.GetOAuthUserInfoResponse
	6, // [6:8] is the sub-list for method output_type
	4, // [4:6] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_npool_third_mw_v1_oauth_oauth_proto_init() }
func file_npool_third_mw_v1_oauth_oauth_proto_init() {
	if File_npool_third_mw_v1_oauth_oauth_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_npool_third_mw_v1_oauth_oauth_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ThirdUserInfo); i {
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
		file_npool_third_mw_v1_oauth_oauth_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccessTokenInfo); i {
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
		file_npool_third_mw_v1_oauth_oauth_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetOAuthAccessTokenRequest); i {
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
		file_npool_third_mw_v1_oauth_oauth_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetOAuthAccessTokenResponse); i {
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
		file_npool_third_mw_v1_oauth_oauth_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetOAuthUserInfoRequest); i {
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
		file_npool_third_mw_v1_oauth_oauth_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetOAuthUserInfoResponse); i {
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
			RawDescriptor: file_npool_third_mw_v1_oauth_oauth_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_npool_third_mw_v1_oauth_oauth_proto_goTypes,
		DependencyIndexes: file_npool_third_mw_v1_oauth_oauth_proto_depIdxs,
		MessageInfos:      file_npool_third_mw_v1_oauth_oauth_proto_msgTypes,
	}.Build()
	File_npool_third_mw_v1_oauth_oauth_proto = out.File
	file_npool_third_mw_v1_oauth_oauth_proto_rawDesc = nil
	file_npool_third_mw_v1_oauth_oauth_proto_goTypes = nil
	file_npool_third_mw_v1_oauth_oauth_proto_depIdxs = nil
}
