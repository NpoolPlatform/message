// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.18.1
// source: npool/appuser/gw/v1/ga/ga.proto

package ga

import (
	_ "github.com/NpoolPlatform/message/npool/appuser/mgr/v2/banapp"
	_ "github.com/NpoolPlatform/message/npool/appuser/mgr/v2/banappuser"
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

type SetupGoogleAuthRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Permissioned by path
	AppID  string `protobuf:"bytes,10,opt,name=AppID,proto3" json:"AppID,omitempty"`
	UserID string `protobuf:"bytes,20,opt,name=UserID,proto3" json:"UserID,omitempty"`
}

func (x *SetupGoogleAuthRequest) Reset() {
	*x = SetupGoogleAuthRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_appuser_gw_v1_ga_ga_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetupGoogleAuthRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetupGoogleAuthRequest) ProtoMessage() {}

func (x *SetupGoogleAuthRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_appuser_gw_v1_ga_ga_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetupGoogleAuthRequest.ProtoReflect.Descriptor instead.
func (*SetupGoogleAuthRequest) Descriptor() ([]byte, []int) {
	return file_npool_appuser_gw_v1_ga_ga_proto_rawDescGZIP(), []int{0}
}

func (x *SetupGoogleAuthRequest) GetAppID() string {
	if x != nil {
		return x.AppID
	}
	return ""
}

func (x *SetupGoogleAuthRequest) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

type SetupGoogleAuthResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *user.User `protobuf:"bytes,10,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (x *SetupGoogleAuthResponse) Reset() {
	*x = SetupGoogleAuthResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_appuser_gw_v1_ga_ga_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetupGoogleAuthResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetupGoogleAuthResponse) ProtoMessage() {}

func (x *SetupGoogleAuthResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_appuser_gw_v1_ga_ga_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetupGoogleAuthResponse.ProtoReflect.Descriptor instead.
func (*SetupGoogleAuthResponse) Descriptor() ([]byte, []int) {
	return file_npool_appuser_gw_v1_ga_ga_proto_rawDescGZIP(), []int{1}
}

func (x *SetupGoogleAuthResponse) GetInfo() *user.User {
	if x != nil {
		return x.Info
	}
	return nil
}

type VerifyGoogleAuthRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Permissioned by path
	AppID  string `protobuf:"bytes,10,opt,name=AppID,proto3" json:"AppID,omitempty"`
	UserID string `protobuf:"bytes,20,opt,name=UserID,proto3" json:"UserID,omitempty"`
	Code   string `protobuf:"bytes,30,opt,name=Code,proto3" json:"Code,omitempty"`
}

func (x *VerifyGoogleAuthRequest) Reset() {
	*x = VerifyGoogleAuthRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_appuser_gw_v1_ga_ga_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VerifyGoogleAuthRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VerifyGoogleAuthRequest) ProtoMessage() {}

func (x *VerifyGoogleAuthRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_appuser_gw_v1_ga_ga_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VerifyGoogleAuthRequest.ProtoReflect.Descriptor instead.
func (*VerifyGoogleAuthRequest) Descriptor() ([]byte, []int) {
	return file_npool_appuser_gw_v1_ga_ga_proto_rawDescGZIP(), []int{2}
}

func (x *VerifyGoogleAuthRequest) GetAppID() string {
	if x != nil {
		return x.AppID
	}
	return ""
}

func (x *VerifyGoogleAuthRequest) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

func (x *VerifyGoogleAuthRequest) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

type VerifyGoogleAuthResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *user.User `protobuf:"bytes,10,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (x *VerifyGoogleAuthResponse) Reset() {
	*x = VerifyGoogleAuthResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_appuser_gw_v1_ga_ga_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VerifyGoogleAuthResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VerifyGoogleAuthResponse) ProtoMessage() {}

func (x *VerifyGoogleAuthResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_appuser_gw_v1_ga_ga_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VerifyGoogleAuthResponse.ProtoReflect.Descriptor instead.
func (*VerifyGoogleAuthResponse) Descriptor() ([]byte, []int) {
	return file_npool_appuser_gw_v1_ga_ga_proto_rawDescGZIP(), []int{3}
}

func (x *VerifyGoogleAuthResponse) GetInfo() *user.User {
	if x != nil {
		return x.Info
	}
	return nil
}

var File_npool_appuser_gw_v1_ga_ga_proto protoreflect.FileDescriptor

var file_npool_appuser_gw_v1_ga_ga_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x61, 0x70, 0x70, 0x75, 0x73, 0x65, 0x72, 0x2f,
	0x67, 0x77, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x61, 0x2f, 0x67, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x15, 0x61, 0x70, 0x70, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77,
	0x61, 0x79, 0x2e, 0x67, 0x61, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x23, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x61, 0x70,
	0x70, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x6d, 0x77, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72,
	0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x28, 0x6e, 0x70, 0x6f,
	0x6f, 0x6c, 0x2f, 0x61, 0x70, 0x70, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x6d, 0x67, 0x72, 0x2f, 0x76,
	0x32, 0x2f, 0x62, 0x61, 0x6e, 0x61, 0x70, 0x70, 0x2f, 0x62, 0x61, 0x6e, 0x61, 0x70, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x30, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x61, 0x70, 0x70,
	0x75, 0x73, 0x65, 0x72, 0x2f, 0x6d, 0x67, 0x72, 0x2f, 0x76, 0x32, 0x2f, 0x62, 0x61, 0x6e, 0x61,
	0x70, 0x70, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x62, 0x61, 0x6e, 0x61, 0x70, 0x70, 0x75, 0x73, 0x65,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x46, 0x0a, 0x16, 0x53, 0x65, 0x74, 0x75, 0x70,
	0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x14, 0x0a, 0x05, 0x41, 0x70, 0x70, 0x49, 0x44, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x41, 0x70, 0x70, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49,
	0x44, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x22,
	0x4f, 0x0a, 0x17, 0x53, 0x65, 0x74, 0x75, 0x70, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x75,
	0x74, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x34, 0x0a, 0x04, 0x49, 0x6e,
	0x66, 0x6f, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x61, 0x70, 0x70, 0x75, 0x73,
	0x65, 0x72, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x49, 0x6e, 0x66, 0x6f,
	0x22, 0x5b, 0x0a, 0x17, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x41,
	0x70, 0x70, 0x49, 0x44, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x41, 0x70, 0x70, 0x49,
	0x44, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x14, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x43, 0x6f, 0x64,
	0x65, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x22, 0x50, 0x0a,
	0x18, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x75, 0x74,
	0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x34, 0x0a, 0x04, 0x49, 0x6e, 0x66,
	0x6f, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x61, 0x70, 0x70, 0x75, 0x73, 0x65,
	0x72, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x32,
	0xb5, 0x02, 0x0a, 0x07, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x12, 0x91, 0x01, 0x0a, 0x0f,
	0x53, 0x65, 0x74, 0x75, 0x70, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x75, 0x74, 0x68, 0x12,
	0x2d, 0x2e, 0x61, 0x70, 0x70, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61,
	0x79, 0x2e, 0x67, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x74, 0x75, 0x70, 0x47, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2e,
	0x2e, 0x61, 0x70, 0x70, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79,
	0x2e, 0x67, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x74, 0x75, 0x70, 0x47, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1f,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x19, 0x22, 0x14, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x74, 0x75,
	0x70, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x75, 0x74, 0x68, 0x3a, 0x01, 0x2a, 0x12,
	0x95, 0x01, 0x0a, 0x10, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x41, 0x75, 0x74, 0x68, 0x12, 0x2e, 0x2e, 0x61, 0x70, 0x70, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x67,
	0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x67, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x56, 0x65, 0x72,
	0x69, 0x66, 0x79, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x2f, 0x2e, 0x61, 0x70, 0x70, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x67,
	0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x67, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x56, 0x65, 0x72,
	0x69, 0x66, 0x79, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x20, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1a, 0x22, 0x15, 0x2f,
	0x76, 0x31, 0x2f, 0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x61, 0x75, 0x74, 0x68, 0x3a, 0x01, 0x2a, 0x42, 0x39, 0x5a, 0x37, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4e, 0x70, 0x6f, 0x6f, 0x6c, 0x50, 0x6c, 0x61, 0x74, 0x66,
	0x6f, 0x72, 0x6d, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x6e, 0x70, 0x6f, 0x6f,
	0x6c, 0x2f, 0x61, 0x70, 0x70, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x67, 0x77, 0x2f, 0x76, 0x31, 0x2f,
	0x67, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_npool_appuser_gw_v1_ga_ga_proto_rawDescOnce sync.Once
	file_npool_appuser_gw_v1_ga_ga_proto_rawDescData = file_npool_appuser_gw_v1_ga_ga_proto_rawDesc
)

func file_npool_appuser_gw_v1_ga_ga_proto_rawDescGZIP() []byte {
	file_npool_appuser_gw_v1_ga_ga_proto_rawDescOnce.Do(func() {
		file_npool_appuser_gw_v1_ga_ga_proto_rawDescData = protoimpl.X.CompressGZIP(file_npool_appuser_gw_v1_ga_ga_proto_rawDescData)
	})
	return file_npool_appuser_gw_v1_ga_ga_proto_rawDescData
}

var file_npool_appuser_gw_v1_ga_ga_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_npool_appuser_gw_v1_ga_ga_proto_goTypes = []interface{}{
	(*SetupGoogleAuthRequest)(nil),   // 0: appuser.gateway.ga.v1.SetupGoogleAuthRequest
	(*SetupGoogleAuthResponse)(nil),  // 1: appuser.gateway.ga.v1.SetupGoogleAuthResponse
	(*VerifyGoogleAuthRequest)(nil),  // 2: appuser.gateway.ga.v1.VerifyGoogleAuthRequest
	(*VerifyGoogleAuthResponse)(nil), // 3: appuser.gateway.ga.v1.VerifyGoogleAuthResponse
	(*user.User)(nil),                // 4: appuser.middleware.user.v1.User
}
var file_npool_appuser_gw_v1_ga_ga_proto_depIdxs = []int32{
	4, // 0: appuser.gateway.ga.v1.SetupGoogleAuthResponse.Info:type_name -> appuser.middleware.user.v1.User
	4, // 1: appuser.gateway.ga.v1.VerifyGoogleAuthResponse.Info:type_name -> appuser.middleware.user.v1.User
	0, // 2: appuser.gateway.ga.v1.Gateway.SetupGoogleAuth:input_type -> appuser.gateway.ga.v1.SetupGoogleAuthRequest
	2, // 3: appuser.gateway.ga.v1.Gateway.VerifyGoogleAuth:input_type -> appuser.gateway.ga.v1.VerifyGoogleAuthRequest
	1, // 4: appuser.gateway.ga.v1.Gateway.SetupGoogleAuth:output_type -> appuser.gateway.ga.v1.SetupGoogleAuthResponse
	3, // 5: appuser.gateway.ga.v1.Gateway.VerifyGoogleAuth:output_type -> appuser.gateway.ga.v1.VerifyGoogleAuthResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_npool_appuser_gw_v1_ga_ga_proto_init() }
func file_npool_appuser_gw_v1_ga_ga_proto_init() {
	if File_npool_appuser_gw_v1_ga_ga_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_npool_appuser_gw_v1_ga_ga_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetupGoogleAuthRequest); i {
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
		file_npool_appuser_gw_v1_ga_ga_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetupGoogleAuthResponse); i {
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
		file_npool_appuser_gw_v1_ga_ga_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VerifyGoogleAuthRequest); i {
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
		file_npool_appuser_gw_v1_ga_ga_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VerifyGoogleAuthResponse); i {
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
			RawDescriptor: file_npool_appuser_gw_v1_ga_ga_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_npool_appuser_gw_v1_ga_ga_proto_goTypes,
		DependencyIndexes: file_npool_appuser_gw_v1_ga_ga_proto_depIdxs,
		MessageInfos:      file_npool_appuser_gw_v1_ga_ga_proto_msgTypes,
	}.Build()
	File_npool_appuser_gw_v1_ga_ga_proto = out.File
	file_npool_appuser_gw_v1_ga_ga_proto_rawDesc = nil
	file_npool_appuser_gw_v1_ga_ga_proto_goTypes = nil
	file_npool_appuser_gw_v1_ga_ga_proto_depIdxs = nil
}
