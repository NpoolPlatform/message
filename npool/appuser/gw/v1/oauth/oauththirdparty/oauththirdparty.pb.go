// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v3.18.1
// source: npool/appuser/gw/v1/oauth/oauththirdparty/oauththirdparty.proto

package oauththirdparty

import (
	oauththirdparty "github.com/NpoolPlatform/message/npool/appuser/mw/v1/oauth/oauththirdparty"
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

type CreateOAuthThirdPartyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClientName     v1.SignMethod `protobuf:"varint,20,opt,name=ClientName,proto3,enum=basetypes.v1.SignMethod" json:"ClientName,omitempty"`
	ClientTag      string        `protobuf:"bytes,30,opt,name=ClientTag,proto3" json:"ClientTag,omitempty"`
	ClientLogoURL  string        `protobuf:"bytes,40,opt,name=ClientLogoURL,proto3" json:"ClientLogoURL,omitempty"`
	ClientOAuthURL string        `protobuf:"bytes,50,opt,name=ClientOAuthURL,proto3" json:"ClientOAuthURL,omitempty"`
	ResponseType   string        `protobuf:"bytes,60,opt,name=ResponseType,proto3" json:"ResponseType,omitempty"`
	Scope          string        `protobuf:"bytes,70,opt,name=Scope,proto3" json:"Scope,omitempty"`
}

func (x *CreateOAuthThirdPartyRequest) Reset() {
	*x = CreateOAuthThirdPartyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_appuser_gw_v1_oauth_oauththirdparty_oauththirdparty_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateOAuthThirdPartyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateOAuthThirdPartyRequest) ProtoMessage() {}

func (x *CreateOAuthThirdPartyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_appuser_gw_v1_oauth_oauththirdparty_oauththirdparty_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateOAuthThirdPartyRequest.ProtoReflect.Descriptor instead.
func (*CreateOAuthThirdPartyRequest) Descriptor() ([]byte, []int) {
	return file_npool_appuser_gw_v1_oauth_oauththirdparty_oauththirdparty_proto_rawDescGZIP(), []int{0}
}

func (x *CreateOAuthThirdPartyRequest) GetClientName() v1.SignMethod {
	if x != nil {
		return x.ClientName
	}
	return v1.SignMethod(0)
}

func (x *CreateOAuthThirdPartyRequest) GetClientTag() string {
	if x != nil {
		return x.ClientTag
	}
	return ""
}

func (x *CreateOAuthThirdPartyRequest) GetClientLogoURL() string {
	if x != nil {
		return x.ClientLogoURL
	}
	return ""
}

func (x *CreateOAuthThirdPartyRequest) GetClientOAuthURL() string {
	if x != nil {
		return x.ClientOAuthURL
	}
	return ""
}

func (x *CreateOAuthThirdPartyRequest) GetResponseType() string {
	if x != nil {
		return x.ResponseType
	}
	return ""
}

func (x *CreateOAuthThirdPartyRequest) GetScope() string {
	if x != nil {
		return x.Scope
	}
	return ""
}

type CreateOAuthThirdPartyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *oauththirdparty.OAuthThirdParty `protobuf:"bytes,10,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (x *CreateOAuthThirdPartyResponse) Reset() {
	*x = CreateOAuthThirdPartyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_appuser_gw_v1_oauth_oauththirdparty_oauththirdparty_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateOAuthThirdPartyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateOAuthThirdPartyResponse) ProtoMessage() {}

func (x *CreateOAuthThirdPartyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_appuser_gw_v1_oauth_oauththirdparty_oauththirdparty_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateOAuthThirdPartyResponse.ProtoReflect.Descriptor instead.
func (*CreateOAuthThirdPartyResponse) Descriptor() ([]byte, []int) {
	return file_npool_appuser_gw_v1_oauth_oauththirdparty_oauththirdparty_proto_rawDescGZIP(), []int{1}
}

func (x *CreateOAuthThirdPartyResponse) GetInfo() *oauththirdparty.OAuthThirdParty {
	if x != nil {
		return x.Info
	}
	return nil
}

type UpdateOAuthThirdPartyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID             uint32         `protobuf:"varint,10,opt,name=ID,proto3" json:"ID,omitempty"`
	EntID          string         `protobuf:"bytes,11,opt,name=EntID,proto3" json:"EntID,omitempty"`
	ClientName     *v1.SignMethod `protobuf:"varint,20,opt,name=ClientName,proto3,enum=basetypes.v1.SignMethod,oneof" json:"ClientName,omitempty"`
	ClientTag      *string        `protobuf:"bytes,30,opt,name=ClientTag,proto3,oneof" json:"ClientTag,omitempty"`
	ClientLogoURL  *string        `protobuf:"bytes,40,opt,name=ClientLogoURL,proto3,oneof" json:"ClientLogoURL,omitempty"`
	ClientOAuthURL *string        `protobuf:"bytes,50,opt,name=ClientOAuthURL,proto3,oneof" json:"ClientOAuthURL,omitempty"`
	ResponseType   *string        `protobuf:"bytes,60,opt,name=ResponseType,proto3,oneof" json:"ResponseType,omitempty"`
	Scope          *string        `protobuf:"bytes,70,opt,name=Scope,proto3,oneof" json:"Scope,omitempty"`
}

func (x *UpdateOAuthThirdPartyRequest) Reset() {
	*x = UpdateOAuthThirdPartyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_appuser_gw_v1_oauth_oauththirdparty_oauththirdparty_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateOAuthThirdPartyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateOAuthThirdPartyRequest) ProtoMessage() {}

func (x *UpdateOAuthThirdPartyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_appuser_gw_v1_oauth_oauththirdparty_oauththirdparty_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateOAuthThirdPartyRequest.ProtoReflect.Descriptor instead.
func (*UpdateOAuthThirdPartyRequest) Descriptor() ([]byte, []int) {
	return file_npool_appuser_gw_v1_oauth_oauththirdparty_oauththirdparty_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateOAuthThirdPartyRequest) GetID() uint32 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *UpdateOAuthThirdPartyRequest) GetEntID() string {
	if x != nil {
		return x.EntID
	}
	return ""
}

func (x *UpdateOAuthThirdPartyRequest) GetClientName() v1.SignMethod {
	if x != nil && x.ClientName != nil {
		return *x.ClientName
	}
	return v1.SignMethod(0)
}

func (x *UpdateOAuthThirdPartyRequest) GetClientTag() string {
	if x != nil && x.ClientTag != nil {
		return *x.ClientTag
	}
	return ""
}

func (x *UpdateOAuthThirdPartyRequest) GetClientLogoURL() string {
	if x != nil && x.ClientLogoURL != nil {
		return *x.ClientLogoURL
	}
	return ""
}

func (x *UpdateOAuthThirdPartyRequest) GetClientOAuthURL() string {
	if x != nil && x.ClientOAuthURL != nil {
		return *x.ClientOAuthURL
	}
	return ""
}

func (x *UpdateOAuthThirdPartyRequest) GetResponseType() string {
	if x != nil && x.ResponseType != nil {
		return *x.ResponseType
	}
	return ""
}

func (x *UpdateOAuthThirdPartyRequest) GetScope() string {
	if x != nil && x.Scope != nil {
		return *x.Scope
	}
	return ""
}

type UpdateOAuthThirdPartyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *oauththirdparty.OAuthThirdParty `protobuf:"bytes,10,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (x *UpdateOAuthThirdPartyResponse) Reset() {
	*x = UpdateOAuthThirdPartyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_appuser_gw_v1_oauth_oauththirdparty_oauththirdparty_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateOAuthThirdPartyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateOAuthThirdPartyResponse) ProtoMessage() {}

func (x *UpdateOAuthThirdPartyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_appuser_gw_v1_oauth_oauththirdparty_oauththirdparty_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateOAuthThirdPartyResponse.ProtoReflect.Descriptor instead.
func (*UpdateOAuthThirdPartyResponse) Descriptor() ([]byte, []int) {
	return file_npool_appuser_gw_v1_oauth_oauththirdparty_oauththirdparty_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateOAuthThirdPartyResponse) GetInfo() *oauththirdparty.OAuthThirdParty {
	if x != nil {
		return x.Info
	}
	return nil
}

type DeleteOAuthThirdPartyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID    uint32 `protobuf:"varint,10,opt,name=ID,proto3" json:"ID,omitempty"`
	EntID string `protobuf:"bytes,20,opt,name=EntID,proto3" json:"EntID,omitempty"`
}

func (x *DeleteOAuthThirdPartyRequest) Reset() {
	*x = DeleteOAuthThirdPartyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_appuser_gw_v1_oauth_oauththirdparty_oauththirdparty_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteOAuthThirdPartyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteOAuthThirdPartyRequest) ProtoMessage() {}

func (x *DeleteOAuthThirdPartyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_appuser_gw_v1_oauth_oauththirdparty_oauththirdparty_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteOAuthThirdPartyRequest.ProtoReflect.Descriptor instead.
func (*DeleteOAuthThirdPartyRequest) Descriptor() ([]byte, []int) {
	return file_npool_appuser_gw_v1_oauth_oauththirdparty_oauththirdparty_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteOAuthThirdPartyRequest) GetID() uint32 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *DeleteOAuthThirdPartyRequest) GetEntID() string {
	if x != nil {
		return x.EntID
	}
	return ""
}

type DeleteOAuthThirdPartyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *oauththirdparty.OAuthThirdParty `protobuf:"bytes,10,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (x *DeleteOAuthThirdPartyResponse) Reset() {
	*x = DeleteOAuthThirdPartyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_appuser_gw_v1_oauth_oauththirdparty_oauththirdparty_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteOAuthThirdPartyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteOAuthThirdPartyResponse) ProtoMessage() {}

func (x *DeleteOAuthThirdPartyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_appuser_gw_v1_oauth_oauththirdparty_oauththirdparty_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteOAuthThirdPartyResponse.ProtoReflect.Descriptor instead.
func (*DeleteOAuthThirdPartyResponse) Descriptor() ([]byte, []int) {
	return file_npool_appuser_gw_v1_oauth_oauththirdparty_oauththirdparty_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteOAuthThirdPartyResponse) GetInfo() *oauththirdparty.OAuthThirdParty {
	if x != nil {
		return x.Info
	}
	return nil
}

type GetOAuthThirdPartiesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Offset int32 `protobuf:"varint,10,opt,name=Offset,proto3" json:"Offset,omitempty"`
	Limit  int32 `protobuf:"varint,20,opt,name=Limit,proto3" json:"Limit,omitempty"`
}

func (x *GetOAuthThirdPartiesRequest) Reset() {
	*x = GetOAuthThirdPartiesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_appuser_gw_v1_oauth_oauththirdparty_oauththirdparty_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetOAuthThirdPartiesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOAuthThirdPartiesRequest) ProtoMessage() {}

func (x *GetOAuthThirdPartiesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_appuser_gw_v1_oauth_oauththirdparty_oauththirdparty_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOAuthThirdPartiesRequest.ProtoReflect.Descriptor instead.
func (*GetOAuthThirdPartiesRequest) Descriptor() ([]byte, []int) {
	return file_npool_appuser_gw_v1_oauth_oauththirdparty_oauththirdparty_proto_rawDescGZIP(), []int{6}
}

func (x *GetOAuthThirdPartiesRequest) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *GetOAuthThirdPartiesRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type GetOAuthThirdPartiesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Infos []*oauththirdparty.OAuthThirdParty `protobuf:"bytes,10,rep,name=Infos,proto3" json:"Infos,omitempty"`
	Total uint32                             `protobuf:"varint,20,opt,name=Total,proto3" json:"Total,omitempty"`
}

func (x *GetOAuthThirdPartiesResponse) Reset() {
	*x = GetOAuthThirdPartiesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_appuser_gw_v1_oauth_oauththirdparty_oauththirdparty_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetOAuthThirdPartiesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOAuthThirdPartiesResponse) ProtoMessage() {}

func (x *GetOAuthThirdPartiesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_appuser_gw_v1_oauth_oauththirdparty_oauththirdparty_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOAuthThirdPartiesResponse.ProtoReflect.Descriptor instead.
func (*GetOAuthThirdPartiesResponse) Descriptor() ([]byte, []int) {
	return file_npool_appuser_gw_v1_oauth_oauththirdparty_oauththirdparty_proto_rawDescGZIP(), []int{7}
}

func (x *GetOAuthThirdPartiesResponse) GetInfos() []*oauththirdparty.OAuthThirdParty {
	if x != nil {
		return x.Infos
	}
	return nil
}

func (x *GetOAuthThirdPartiesResponse) GetTotal() uint32 {
	if x != nil {
		return x.Total
	}
	return 0
}

var File_npool_appuser_gw_v1_oauth_oauththirdparty_oauththirdparty_proto protoreflect.FileDescriptor

var file_npool_appuser_gw_v1_oauth_oauththirdparty_oauththirdparty_proto_rawDesc = []byte{
	0x0a, 0x3f, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x61, 0x70, 0x70, 0x75, 0x73, 0x65, 0x72, 0x2f,
	0x67, 0x77, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x6f, 0x61, 0x75, 0x74,
	0x68, 0x74, 0x68, 0x69, 0x72, 0x64, 0x70, 0x61, 0x72, 0x74, 0x79, 0x2f, 0x6f, 0x61, 0x75, 0x74,
	0x68, 0x74, 0x68, 0x69, 0x72, 0x64, 0x70, 0x61, 0x72, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x28, 0x61, 0x70, 0x70, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77,
	0x61, 0x79, 0x2e, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x74, 0x68,
	0x69, 0x72, 0x64, 0x70, 0x61, 0x72, 0x74, 0x79, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x3f, 0x6e, 0x70, 0x6f, 0x6f, 0x6c,
	0x2f, 0x61, 0x70, 0x70, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x6d, 0x77, 0x2f, 0x76, 0x31, 0x2f, 0x6f,
	0x61, 0x75, 0x74, 0x68, 0x2f, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x74, 0x68, 0x69, 0x72, 0x64, 0x70,
	0x61, 0x72, 0x74, 0x79, 0x2f, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x74, 0x68, 0x69, 0x72, 0x64, 0x70,
	0x61, 0x72, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x23, 0x6e, 0x70, 0x6f, 0x6f,
	0x6c, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x73,
	0x69, 0x67, 0x6e, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xfe, 0x01, 0x0a, 0x1c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x41, 0x75, 0x74, 0x68, 0x54,
	0x68, 0x69, 0x72, 0x64, 0x50, 0x61, 0x72, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x38, 0x0a, 0x0a, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x14,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x18, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x52, 0x0a,
	0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x43, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x54, 0x61, 0x67, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x43,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x54, 0x61, 0x67, 0x12, 0x24, 0x0a, 0x0d, 0x43, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x4c, 0x6f, 0x67, 0x6f, 0x55, 0x52, 0x4c, 0x18, 0x28, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0d, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4c, 0x6f, 0x67, 0x6f, 0x55, 0x52, 0x4c, 0x12, 0x26,
	0x0a, 0x0e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4f, 0x41, 0x75, 0x74, 0x68, 0x55, 0x52, 0x4c,
	0x18, 0x32, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4f, 0x41,
	0x75, 0x74, 0x68, 0x55, 0x52, 0x4c, 0x12, 0x22, 0x0a, 0x0c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x54, 0x79, 0x70, 0x65, 0x18, 0x3c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x53, 0x63,
	0x6f, 0x70, 0x65, 0x18, 0x46, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x53, 0x63, 0x6f, 0x70, 0x65,
	0x22, 0x71, 0x0a, 0x1d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x41, 0x75, 0x74, 0x68, 0x54,
	0x68, 0x69, 0x72, 0x64, 0x50, 0x61, 0x72, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x50, 0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x3c, 0x2e, 0x61, 0x70, 0x70, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65,
	0x77, 0x61, 0x72, 0x65, 0x2e, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x6f, 0x61, 0x75, 0x74, 0x68,
	0x74, 0x68, 0x69, 0x72, 0x64, 0x70, 0x61, 0x72, 0x74, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x41,
	0x75, 0x74, 0x68, 0x54, 0x68, 0x69, 0x72, 0x64, 0x50, 0x61, 0x72, 0x74, 0x79, 0x52, 0x04, 0x49,
	0x6e, 0x66, 0x6f, 0x22, 0x9f, 0x03, 0x0a, 0x1c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4f, 0x41,
	0x75, 0x74, 0x68, 0x54, 0x68, 0x69, 0x72, 0x64, 0x50, 0x61, 0x72, 0x74, 0x79, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x02, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x45, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x0b, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x45, 0x6e, 0x74, 0x49, 0x44, 0x12, 0x3d, 0x0a, 0x0a, 0x43, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x14, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x18,
	0x2e, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x69,
	0x67, 0x6e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x48, 0x00, 0x52, 0x0a, 0x43, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x21, 0x0a, 0x09, 0x43, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x54, 0x61, 0x67, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x09,
	0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x54, 0x61, 0x67, 0x88, 0x01, 0x01, 0x12, 0x29, 0x0a, 0x0d,
	0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4c, 0x6f, 0x67, 0x6f, 0x55, 0x52, 0x4c, 0x18, 0x28, 0x20,
	0x01, 0x28, 0x09, 0x48, 0x02, 0x52, 0x0d, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4c, 0x6f, 0x67,
	0x6f, 0x55, 0x52, 0x4c, 0x88, 0x01, 0x01, 0x12, 0x2b, 0x0a, 0x0e, 0x43, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x4f, 0x41, 0x75, 0x74, 0x68, 0x55, 0x52, 0x4c, 0x18, 0x32, 0x20, 0x01, 0x28, 0x09, 0x48,
	0x03, 0x52, 0x0e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4f, 0x41, 0x75, 0x74, 0x68, 0x55, 0x52,
	0x4c, 0x88, 0x01, 0x01, 0x12, 0x27, 0x0a, 0x0c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x54, 0x79, 0x70, 0x65, 0x18, 0x3c, 0x20, 0x01, 0x28, 0x09, 0x48, 0x04, 0x52, 0x0c, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x54, 0x79, 0x70, 0x65, 0x88, 0x01, 0x01, 0x12, 0x19, 0x0a,
	0x05, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x18, 0x46, 0x20, 0x01, 0x28, 0x09, 0x48, 0x05, 0x52, 0x05,
	0x53, 0x63, 0x6f, 0x70, 0x65, 0x88, 0x01, 0x01, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x43, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x43, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x54, 0x61, 0x67, 0x42, 0x10, 0x0a, 0x0e, 0x5f, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x4c, 0x6f, 0x67, 0x6f, 0x55, 0x52, 0x4c, 0x42, 0x11, 0x0a, 0x0f, 0x5f, 0x43, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x4f, 0x41, 0x75, 0x74, 0x68, 0x55, 0x52, 0x4c, 0x42, 0x0f, 0x0a, 0x0d, 0x5f, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x54, 0x79, 0x70, 0x65, 0x42, 0x08, 0x0a, 0x06, 0x5f,
	0x53, 0x63, 0x6f, 0x70, 0x65, 0x22, 0x71, 0x0a, 0x1d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4f,
	0x41, 0x75, 0x74, 0x68, 0x54, 0x68, 0x69, 0x72, 0x64, 0x50, 0x61, 0x72, 0x74, 0x79, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x50, 0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x3c, 0x2e, 0x61, 0x70, 0x70, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x6d,
	0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x2e,
	0x6f, 0x61, 0x75, 0x74, 0x68, 0x74, 0x68, 0x69, 0x72, 0x64, 0x70, 0x61, 0x72, 0x74, 0x79, 0x2e,
	0x76, 0x31, 0x2e, 0x4f, 0x41, 0x75, 0x74, 0x68, 0x54, 0x68, 0x69, 0x72, 0x64, 0x50, 0x61, 0x72,
	0x74, 0x79, 0x52, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x44, 0x0a, 0x1c, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x4f, 0x41, 0x75, 0x74, 0x68, 0x54, 0x68, 0x69, 0x72, 0x64, 0x50, 0x61, 0x72, 0x74,
	0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x45, 0x6e, 0x74, 0x49,
	0x44, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x45, 0x6e, 0x74, 0x49, 0x44, 0x22, 0x71,
	0x0a, 0x1d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4f, 0x41, 0x75, 0x74, 0x68, 0x54, 0x68, 0x69,
	0x72, 0x64, 0x50, 0x61, 0x72, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x50, 0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3c, 0x2e,
	0x61, 0x70, 0x70, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61,
	0x72, 0x65, 0x2e, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x74, 0x68,
	0x69, 0x72, 0x64, 0x70, 0x61, 0x72, 0x74, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x41, 0x75, 0x74,
	0x68, 0x54, 0x68, 0x69, 0x72, 0x64, 0x50, 0x61, 0x72, 0x74, 0x79, 0x52, 0x04, 0x49, 0x6e, 0x66,
	0x6f, 0x22, 0x4b, 0x0a, 0x1b, 0x47, 0x65, 0x74, 0x4f, 0x41, 0x75, 0x74, 0x68, 0x54, 0x68, 0x69,
	0x72, 0x64, 0x50, 0x61, 0x72, 0x74, 0x69, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x16, 0x0a, 0x06, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x06, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x4c, 0x69, 0x6d, 0x69,
	0x74, 0x18, 0x14, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x22, 0x88,
	0x01, 0x0a, 0x1c, 0x47, 0x65, 0x74, 0x4f, 0x41, 0x75, 0x74, 0x68, 0x54, 0x68, 0x69, 0x72, 0x64,
	0x50, 0x61, 0x72, 0x74, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x52, 0x0a, 0x05, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3c,
	0x2e, 0x61, 0x70, 0x70, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77,
	0x61, 0x72, 0x65, 0x2e, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x74,
	0x68, 0x69, 0x72, 0x64, 0x70, 0x61, 0x72, 0x74, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x41, 0x75,
	0x74, 0x68, 0x54, 0x68, 0x69, 0x72, 0x64, 0x50, 0x61, 0x72, 0x74, 0x79, 0x52, 0x05, 0x49, 0x6e,
	0x66, 0x6f, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x14, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x32, 0xcd, 0x06, 0x0a, 0x07, 0x47, 0x61,
	0x74, 0x65, 0x77, 0x61, 0x79, 0x12, 0xcf, 0x01, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x4f, 0x41, 0x75, 0x74, 0x68, 0x54, 0x68, 0x69, 0x72, 0x64, 0x50, 0x61, 0x72, 0x74, 0x79, 0x12,
	0x46, 0x2e, 0x61, 0x70, 0x70, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61,
	0x79, 0x2e, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x74, 0x68, 0x69,
	0x72, 0x64, 0x70, 0x61, 0x72, 0x74, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x4f, 0x41, 0x75, 0x74, 0x68, 0x54, 0x68, 0x69, 0x72, 0x64, 0x50, 0x61, 0x72, 0x74, 0x79,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x47, 0x2e, 0x61, 0x70, 0x70, 0x75, 0x73, 0x65,
	0x72, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x2e,
	0x6f, 0x61, 0x75, 0x74, 0x68, 0x74, 0x68, 0x69, 0x72, 0x64, 0x70, 0x61, 0x72, 0x74, 0x79, 0x2e,
	0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x41, 0x75, 0x74, 0x68, 0x54, 0x68,
	0x69, 0x72, 0x64, 0x50, 0x61, 0x72, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x25, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1f, 0x3a, 0x01, 0x2a, 0x22, 0x1a, 0x2f, 0x76, 0x31,
	0x2f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x2f, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x74, 0x68, 0x69,
	0x72, 0x64, 0x70, 0x61, 0x72, 0x74, 0x79, 0x12, 0xcf, 0x01, 0x0a, 0x15, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x4f, 0x41, 0x75, 0x74, 0x68, 0x54, 0x68, 0x69, 0x72, 0x64, 0x50, 0x61, 0x72, 0x74,
	0x79, 0x12, 0x46, 0x2e, 0x61, 0x70, 0x70, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x67, 0x61, 0x74, 0x65,
	0x77, 0x61, 0x79, 0x2e, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x74,
	0x68, 0x69, 0x72, 0x64, 0x70, 0x61, 0x72, 0x74, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x4f, 0x41, 0x75, 0x74, 0x68, 0x54, 0x68, 0x69, 0x72, 0x64, 0x50, 0x61, 0x72,
	0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x47, 0x2e, 0x61, 0x70, 0x70, 0x75,
	0x73, 0x65, 0x72, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x6f, 0x61, 0x75, 0x74,
	0x68, 0x2e, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x74, 0x68, 0x69, 0x72, 0x64, 0x70, 0x61, 0x72, 0x74,
	0x79, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4f, 0x41, 0x75, 0x74, 0x68,
	0x54, 0x68, 0x69, 0x72, 0x64, 0x50, 0x61, 0x72, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x25, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1f, 0x3a, 0x01, 0x2a, 0x22, 0x1a, 0x2f,
	0x76, 0x31, 0x2f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x74,
	0x68, 0x69, 0x72, 0x64, 0x70, 0x61, 0x72, 0x74, 0x79, 0x12, 0xcf, 0x01, 0x0a, 0x15, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x4f, 0x41, 0x75, 0x74, 0x68, 0x54, 0x68, 0x69, 0x72, 0x64, 0x50, 0x61,
	0x72, 0x74, 0x79, 0x12, 0x46, 0x2e, 0x61, 0x70, 0x70, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x67, 0x61,
	0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x6f, 0x61, 0x75, 0x74,
	0x68, 0x74, 0x68, 0x69, 0x72, 0x64, 0x70, 0x61, 0x72, 0x74, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x4f, 0x41, 0x75, 0x74, 0x68, 0x54, 0x68, 0x69, 0x72, 0x64, 0x50,
	0x61, 0x72, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x47, 0x2e, 0x61, 0x70,
	0x70, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x6f, 0x61,
	0x75, 0x74, 0x68, 0x2e, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x74, 0x68, 0x69, 0x72, 0x64, 0x70, 0x61,
	0x72, 0x74, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4f, 0x41, 0x75,
	0x74, 0x68, 0x54, 0x68, 0x69, 0x72, 0x64, 0x50, 0x61, 0x72, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x25, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1f, 0x3a, 0x01, 0x2a, 0x22,
	0x1a, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x2f, 0x6f, 0x61, 0x75, 0x74,
	0x68, 0x74, 0x68, 0x69, 0x72, 0x64, 0x70, 0x61, 0x72, 0x74, 0x79, 0x12, 0xcb, 0x01, 0x0a, 0x14,
	0x47, 0x65, 0x74, 0x4f, 0x41, 0x75, 0x74, 0x68, 0x54, 0x68, 0x69, 0x72, 0x64, 0x50, 0x61, 0x72,
	0x74, 0x69, 0x65, 0x73, 0x12, 0x45, 0x2e, 0x61, 0x70, 0x70, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x67,
	0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x6f, 0x61, 0x75,
	0x74, 0x68, 0x74, 0x68, 0x69, 0x72, 0x64, 0x70, 0x61, 0x72, 0x74, 0x79, 0x2e, 0x76, 0x31, 0x2e,
	0x47, 0x65, 0x74, 0x4f, 0x41, 0x75, 0x74, 0x68, 0x54, 0x68, 0x69, 0x72, 0x64, 0x50, 0x61, 0x72,
	0x74, 0x69, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x46, 0x2e, 0x61, 0x70,
	0x70, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x6f, 0x61,
	0x75, 0x74, 0x68, 0x2e, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x74, 0x68, 0x69, 0x72, 0x64, 0x70, 0x61,
	0x72, 0x74, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4f, 0x41, 0x75, 0x74, 0x68, 0x54,
	0x68, 0x69, 0x72, 0x64, 0x50, 0x61, 0x72, 0x74, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x24, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1e, 0x3a, 0x01, 0x2a, 0x22, 0x19,
	0x2f, 0x76, 0x31, 0x2f, 0x67, 0x65, 0x74, 0x2f, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x74, 0x68, 0x69,
	0x72, 0x64, 0x70, 0x61, 0x72, 0x74, 0x69, 0x65, 0x73, 0x42, 0x4c, 0x5a, 0x4a, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4e, 0x70, 0x6f, 0x6f, 0x6c, 0x50, 0x6c, 0x61,
	0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x6e, 0x70,
	0x6f, 0x6f, 0x6c, 0x2f, 0x61, 0x70, 0x70, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x67, 0x77, 0x2f, 0x76,
	0x31, 0x2f, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x74, 0x68, 0x69,
	0x72, 0x64, 0x70, 0x61, 0x72, 0x74, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_npool_appuser_gw_v1_oauth_oauththirdparty_oauththirdparty_proto_rawDescOnce sync.Once
	file_npool_appuser_gw_v1_oauth_oauththirdparty_oauththirdparty_proto_rawDescData = file_npool_appuser_gw_v1_oauth_oauththirdparty_oauththirdparty_proto_rawDesc
)

func file_npool_appuser_gw_v1_oauth_oauththirdparty_oauththirdparty_proto_rawDescGZIP() []byte {
	file_npool_appuser_gw_v1_oauth_oauththirdparty_oauththirdparty_proto_rawDescOnce.Do(func() {
		file_npool_appuser_gw_v1_oauth_oauththirdparty_oauththirdparty_proto_rawDescData = protoimpl.X.CompressGZIP(file_npool_appuser_gw_v1_oauth_oauththirdparty_oauththirdparty_proto_rawDescData)
	})
	return file_npool_appuser_gw_v1_oauth_oauththirdparty_oauththirdparty_proto_rawDescData
}

var file_npool_appuser_gw_v1_oauth_oauththirdparty_oauththirdparty_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_npool_appuser_gw_v1_oauth_oauththirdparty_oauththirdparty_proto_goTypes = []interface{}{
	(*CreateOAuthThirdPartyRequest)(nil),    // 0: appuser.gateway.oauth.oauththirdparty.v1.CreateOAuthThirdPartyRequest
	(*CreateOAuthThirdPartyResponse)(nil),   // 1: appuser.gateway.oauth.oauththirdparty.v1.CreateOAuthThirdPartyResponse
	(*UpdateOAuthThirdPartyRequest)(nil),    // 2: appuser.gateway.oauth.oauththirdparty.v1.UpdateOAuthThirdPartyRequest
	(*UpdateOAuthThirdPartyResponse)(nil),   // 3: appuser.gateway.oauth.oauththirdparty.v1.UpdateOAuthThirdPartyResponse
	(*DeleteOAuthThirdPartyRequest)(nil),    // 4: appuser.gateway.oauth.oauththirdparty.v1.DeleteOAuthThirdPartyRequest
	(*DeleteOAuthThirdPartyResponse)(nil),   // 5: appuser.gateway.oauth.oauththirdparty.v1.DeleteOAuthThirdPartyResponse
	(*GetOAuthThirdPartiesRequest)(nil),     // 6: appuser.gateway.oauth.oauththirdparty.v1.GetOAuthThirdPartiesRequest
	(*GetOAuthThirdPartiesResponse)(nil),    // 7: appuser.gateway.oauth.oauththirdparty.v1.GetOAuthThirdPartiesResponse
	(v1.SignMethod)(0),                      // 8: basetypes.v1.SignMethod
	(*oauththirdparty.OAuthThirdParty)(nil), // 9: appuser.middleware.oauth.oauththirdparty.v1.OAuthThirdParty
}
var file_npool_appuser_gw_v1_oauth_oauththirdparty_oauththirdparty_proto_depIdxs = []int32{
	8,  // 0: appuser.gateway.oauth.oauththirdparty.v1.CreateOAuthThirdPartyRequest.ClientName:type_name -> basetypes.v1.SignMethod
	9,  // 1: appuser.gateway.oauth.oauththirdparty.v1.CreateOAuthThirdPartyResponse.Info:type_name -> appuser.middleware.oauth.oauththirdparty.v1.OAuthThirdParty
	8,  // 2: appuser.gateway.oauth.oauththirdparty.v1.UpdateOAuthThirdPartyRequest.ClientName:type_name -> basetypes.v1.SignMethod
	9,  // 3: appuser.gateway.oauth.oauththirdparty.v1.UpdateOAuthThirdPartyResponse.Info:type_name -> appuser.middleware.oauth.oauththirdparty.v1.OAuthThirdParty
	9,  // 4: appuser.gateway.oauth.oauththirdparty.v1.DeleteOAuthThirdPartyResponse.Info:type_name -> appuser.middleware.oauth.oauththirdparty.v1.OAuthThirdParty
	9,  // 5: appuser.gateway.oauth.oauththirdparty.v1.GetOAuthThirdPartiesResponse.Infos:type_name -> appuser.middleware.oauth.oauththirdparty.v1.OAuthThirdParty
	0,  // 6: appuser.gateway.oauth.oauththirdparty.v1.Gateway.CreateOAuthThirdParty:input_type -> appuser.gateway.oauth.oauththirdparty.v1.CreateOAuthThirdPartyRequest
	2,  // 7: appuser.gateway.oauth.oauththirdparty.v1.Gateway.UpdateOAuthThirdParty:input_type -> appuser.gateway.oauth.oauththirdparty.v1.UpdateOAuthThirdPartyRequest
	4,  // 8: appuser.gateway.oauth.oauththirdparty.v1.Gateway.DeleteOAuthThirdParty:input_type -> appuser.gateway.oauth.oauththirdparty.v1.DeleteOAuthThirdPartyRequest
	6,  // 9: appuser.gateway.oauth.oauththirdparty.v1.Gateway.GetOAuthThirdParties:input_type -> appuser.gateway.oauth.oauththirdparty.v1.GetOAuthThirdPartiesRequest
	1,  // 10: appuser.gateway.oauth.oauththirdparty.v1.Gateway.CreateOAuthThirdParty:output_type -> appuser.gateway.oauth.oauththirdparty.v1.CreateOAuthThirdPartyResponse
	3,  // 11: appuser.gateway.oauth.oauththirdparty.v1.Gateway.UpdateOAuthThirdParty:output_type -> appuser.gateway.oauth.oauththirdparty.v1.UpdateOAuthThirdPartyResponse
	5,  // 12: appuser.gateway.oauth.oauththirdparty.v1.Gateway.DeleteOAuthThirdParty:output_type -> appuser.gateway.oauth.oauththirdparty.v1.DeleteOAuthThirdPartyResponse
	7,  // 13: appuser.gateway.oauth.oauththirdparty.v1.Gateway.GetOAuthThirdParties:output_type -> appuser.gateway.oauth.oauththirdparty.v1.GetOAuthThirdPartiesResponse
	10, // [10:14] is the sub-list for method output_type
	6,  // [6:10] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_npool_appuser_gw_v1_oauth_oauththirdparty_oauththirdparty_proto_init() }
func file_npool_appuser_gw_v1_oauth_oauththirdparty_oauththirdparty_proto_init() {
	if File_npool_appuser_gw_v1_oauth_oauththirdparty_oauththirdparty_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_npool_appuser_gw_v1_oauth_oauththirdparty_oauththirdparty_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateOAuthThirdPartyRequest); i {
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
		file_npool_appuser_gw_v1_oauth_oauththirdparty_oauththirdparty_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateOAuthThirdPartyResponse); i {
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
		file_npool_appuser_gw_v1_oauth_oauththirdparty_oauththirdparty_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateOAuthThirdPartyRequest); i {
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
		file_npool_appuser_gw_v1_oauth_oauththirdparty_oauththirdparty_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateOAuthThirdPartyResponse); i {
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
		file_npool_appuser_gw_v1_oauth_oauththirdparty_oauththirdparty_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteOAuthThirdPartyRequest); i {
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
		file_npool_appuser_gw_v1_oauth_oauththirdparty_oauththirdparty_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteOAuthThirdPartyResponse); i {
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
		file_npool_appuser_gw_v1_oauth_oauththirdparty_oauththirdparty_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetOAuthThirdPartiesRequest); i {
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
		file_npool_appuser_gw_v1_oauth_oauththirdparty_oauththirdparty_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetOAuthThirdPartiesResponse); i {
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
	file_npool_appuser_gw_v1_oauth_oauththirdparty_oauththirdparty_proto_msgTypes[2].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_npool_appuser_gw_v1_oauth_oauththirdparty_oauththirdparty_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_npool_appuser_gw_v1_oauth_oauththirdparty_oauththirdparty_proto_goTypes,
		DependencyIndexes: file_npool_appuser_gw_v1_oauth_oauththirdparty_oauththirdparty_proto_depIdxs,
		MessageInfos:      file_npool_appuser_gw_v1_oauth_oauththirdparty_oauththirdparty_proto_msgTypes,
	}.Build()
	File_npool_appuser_gw_v1_oauth_oauththirdparty_oauththirdparty_proto = out.File
	file_npool_appuser_gw_v1_oauth_oauththirdparty_oauththirdparty_proto_rawDesc = nil
	file_npool_appuser_gw_v1_oauth_oauththirdparty_oauththirdparty_proto_goTypes = nil
	file_npool_appuser_gw_v1_oauth_oauththirdparty_oauththirdparty_proto_depIdxs = nil
}
