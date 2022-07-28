// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.18.1
// source: npool/appuser/gw/v1/app/app.proto

package app

import (
	_ "github.com/NpoolPlatform/message/npool"
	app "github.com/NpoolPlatform/message/npool/appuser/mgr/v2/app"
	app1 "github.com/NpoolPlatform/message/npool/appuser/mw/v1/app"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CreateAppRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Permissioned by path
	Info *app.AppReq `protobuf:"bytes,10,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (x *CreateAppRequest) Reset() {
	*x = CreateAppRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_appuser_gw_v1_app_app_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateAppRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAppRequest) ProtoMessage() {}

func (x *CreateAppRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_appuser_gw_v1_app_app_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAppRequest.ProtoReflect.Descriptor instead.
func (*CreateAppRequest) Descriptor() ([]byte, []int) {
	return file_npool_appuser_gw_v1_app_app_proto_rawDescGZIP(), []int{0}
}

func (x *CreateAppRequest) GetInfo() *app.AppReq {
	if x != nil {
		return x.Info
	}
	return nil
}

type CreateAppResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *app.App `protobuf:"bytes,10,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (x *CreateAppResponse) Reset() {
	*x = CreateAppResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_appuser_gw_v1_app_app_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateAppResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAppResponse) ProtoMessage() {}

func (x *CreateAppResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_appuser_gw_v1_app_app_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAppResponse.ProtoReflect.Descriptor instead.
func (*CreateAppResponse) Descriptor() ([]byte, []int) {
	return file_npool_appuser_gw_v1_app_app_proto_rawDescGZIP(), []int{1}
}

func (x *CreateAppResponse) GetInfo() *app.App {
	if x != nil {
		return x.Info
	}
	return nil
}

type UpdateAppRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Permissioned by path (may by user in future if we support user create app)
	Info *app.AppReq `protobuf:"bytes,10,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (x *UpdateAppRequest) Reset() {
	*x = UpdateAppRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_appuser_gw_v1_app_app_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateAppRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateAppRequest) ProtoMessage() {}

func (x *UpdateAppRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_appuser_gw_v1_app_app_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateAppRequest.ProtoReflect.Descriptor instead.
func (*UpdateAppRequest) Descriptor() ([]byte, []int) {
	return file_npool_appuser_gw_v1_app_app_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateAppRequest) GetInfo() *app.AppReq {
	if x != nil {
		return x.Info
	}
	return nil
}

type UpdateAppResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *app.App `protobuf:"bytes,10,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (x *UpdateAppResponse) Reset() {
	*x = UpdateAppResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_appuser_gw_v1_app_app_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateAppResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateAppResponse) ProtoMessage() {}

func (x *UpdateAppResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_appuser_gw_v1_app_app_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateAppResponse.ProtoReflect.Descriptor instead.
func (*UpdateAppResponse) Descriptor() ([]byte, []int) {
	return file_npool_appuser_gw_v1_app_app_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateAppResponse) GetInfo() *app.App {
	if x != nil {
		return x.Info
	}
	return nil
}

type GetAppRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Permissioned by app
	AppID string `protobuf:"bytes,10,opt,name=AppID,proto3" json:"AppID,omitempty"`
}

func (x *GetAppRequest) Reset() {
	*x = GetAppRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_appuser_gw_v1_app_app_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAppRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAppRequest) ProtoMessage() {}

func (x *GetAppRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_appuser_gw_v1_app_app_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAppRequest.ProtoReflect.Descriptor instead.
func (*GetAppRequest) Descriptor() ([]byte, []int) {
	return file_npool_appuser_gw_v1_app_app_proto_rawDescGZIP(), []int{4}
}

func (x *GetAppRequest) GetAppID() string {
	if x != nil {
		return x.AppID
	}
	return ""
}

type GetAppResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *app1.App `protobuf:"bytes,10,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (x *GetAppResponse) Reset() {
	*x = GetAppResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_appuser_gw_v1_app_app_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAppResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAppResponse) ProtoMessage() {}

func (x *GetAppResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_appuser_gw_v1_app_app_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAppResponse.ProtoReflect.Descriptor instead.
func (*GetAppResponse) Descriptor() ([]byte, []int) {
	return file_npool_appuser_gw_v1_app_app_proto_rawDescGZIP(), []int{5}
}

func (x *GetAppResponse) GetInfo() *app1.App {
	if x != nil {
		return x.Info
	}
	return nil
}

type GetAppsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Permissioned by path
	Offset int32 `protobuf:"varint,30,opt,name=Offset,proto3" json:"Offset,omitempty"`
	Limit  int32 `protobuf:"varint,40,opt,name=Limit,proto3" json:"Limit,omitempty"`
}

func (x *GetAppsRequest) Reset() {
	*x = GetAppsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_appuser_gw_v1_app_app_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAppsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAppsRequest) ProtoMessage() {}

func (x *GetAppsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_appuser_gw_v1_app_app_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAppsRequest.ProtoReflect.Descriptor instead.
func (*GetAppsRequest) Descriptor() ([]byte, []int) {
	return file_npool_appuser_gw_v1_app_app_proto_rawDescGZIP(), []int{6}
}

func (x *GetAppsRequest) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *GetAppsRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type GetAppsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Infos []*app1.App `protobuf:"bytes,10,rep,name=Infos,proto3" json:"Infos,omitempty"`
	Total uint32      `protobuf:"varint,20,opt,name=Total,proto3" json:"Total,omitempty"`
}

func (x *GetAppsResponse) Reset() {
	*x = GetAppsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_appuser_gw_v1_app_app_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAppsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAppsResponse) ProtoMessage() {}

func (x *GetAppsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_appuser_gw_v1_app_app_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAppsResponse.ProtoReflect.Descriptor instead.
func (*GetAppsResponse) Descriptor() ([]byte, []int) {
	return file_npool_appuser_gw_v1_app_app_proto_rawDescGZIP(), []int{7}
}

func (x *GetAppsResponse) GetInfos() []*app1.App {
	if x != nil {
		return x.Infos
	}
	return nil
}

func (x *GetAppsResponse) GetTotal() uint32 {
	if x != nil {
		return x.Total
	}
	return 0
}

type GetUserAppsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Permissioned by path
	TargetUserID string `protobuf:"bytes,10,opt,name=TargetUserID,proto3" json:"TargetUserID,omitempty"`
	Offset       int32  `protobuf:"varint,20,opt,name=Offset,proto3" json:"Offset,omitempty"`
	Limit        int32  `protobuf:"varint,30,opt,name=Limit,proto3" json:"Limit,omitempty"`
}

func (x *GetUserAppsRequest) Reset() {
	*x = GetUserAppsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_appuser_gw_v1_app_app_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserAppsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserAppsRequest) ProtoMessage() {}

func (x *GetUserAppsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_appuser_gw_v1_app_app_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserAppsRequest.ProtoReflect.Descriptor instead.
func (*GetUserAppsRequest) Descriptor() ([]byte, []int) {
	return file_npool_appuser_gw_v1_app_app_proto_rawDescGZIP(), []int{8}
}

func (x *GetUserAppsRequest) GetTargetUserID() string {
	if x != nil {
		return x.TargetUserID
	}
	return ""
}

func (x *GetUserAppsRequest) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *GetUserAppsRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type GetUserAppsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Infos []*app1.App `protobuf:"bytes,10,rep,name=Infos,proto3" json:"Infos,omitempty"`
	Total uint32      `protobuf:"varint,20,opt,name=Total,proto3" json:"Total,omitempty"`
}

func (x *GetUserAppsResponse) Reset() {
	*x = GetUserAppsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_appuser_gw_v1_app_app_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserAppsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserAppsResponse) ProtoMessage() {}

func (x *GetUserAppsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_appuser_gw_v1_app_app_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserAppsResponse.ProtoReflect.Descriptor instead.
func (*GetUserAppsResponse) Descriptor() ([]byte, []int) {
	return file_npool_appuser_gw_v1_app_app_proto_rawDescGZIP(), []int{9}
}

func (x *GetUserAppsResponse) GetInfos() []*app1.App {
	if x != nil {
		return x.Infos
	}
	return nil
}

func (x *GetUserAppsResponse) GetTotal() uint32 {
	if x != nil {
		return x.Total
	}
	return 0
}

var File_npool_appuser_gw_v1_app_app_proto protoreflect.FileDescriptor

var file_npool_appuser_gw_v1_app_app_proto_rawDesc = []byte{
	0x0a, 0x21, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x61, 0x70, 0x70, 0x75, 0x73, 0x65, 0x72, 0x2f,
	0x67, 0x77, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x70, 0x70, 0x2f, 0x61, 0x70, 0x70, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x16, 0x61, 0x70, 0x70, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x67, 0x61, 0x74,
	0x65, 0x77, 0x61, 0x79, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x6e, 0x70,
	0x6f, 0x6f, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x22, 0x6e, 0x70, 0x6f, 0x6f, 0x6c,
	0x2f, 0x61, 0x70, 0x70, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x6d, 0x67, 0x72, 0x2f, 0x76, 0x32, 0x2f,
	0x61, 0x70, 0x70, 0x2f, 0x61, 0x70, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x6e,
	0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x61, 0x70, 0x70, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x6d, 0x77, 0x2f,
	0x76, 0x31, 0x2f, 0x61, 0x70, 0x70, 0x2f, 0x61, 0x70, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x46, 0x0a, 0x10, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x70, 0x70, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x32, 0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x61, 0x70, 0x70, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x6d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x76, 0x32, 0x2e, 0x41, 0x70, 0x70, 0x52,
	0x65, 0x71, 0x52, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x44, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x41, 0x70, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2f, 0x0a,
	0x04, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x61, 0x70,
	0x70, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x61, 0x70,
	0x70, 0x2e, 0x76, 0x32, 0x2e, 0x41, 0x70, 0x70, 0x52, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x46,
	0x0a, 0x10, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x70, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x32, 0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1e, 0x2e, 0x61, 0x70, 0x70, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67,
	0x65, 0x72, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x76, 0x32, 0x2e, 0x41, 0x70, 0x70, 0x52, 0x65, 0x71,
	0x52, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x44, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x41, 0x70, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2f, 0x0a, 0x04, 0x49,
	0x6e, 0x66, 0x6f, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x61, 0x70, 0x70, 0x75,
	0x73, 0x65, 0x72, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x70, 0x2e,
	0x76, 0x32, 0x2e, 0x41, 0x70, 0x70, 0x52, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x25, 0x0a, 0x0d,
	0x47, 0x65, 0x74, 0x41, 0x70, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a,
	0x05, 0x41, 0x70, 0x70, 0x49, 0x44, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x41, 0x70,
	0x70, 0x49, 0x44, 0x22, 0x44, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x41, 0x70, 0x70, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x32, 0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x61, 0x70, 0x70, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x6d, 0x69,
	0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x76, 0x31, 0x2e,
	0x41, 0x70, 0x70, 0x52, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x3e, 0x0a, 0x0e, 0x47, 0x65, 0x74,
	0x41, 0x70, 0x70, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x4f,
	0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x4f, 0x66, 0x66,
	0x73, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x28, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x05, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x22, 0x5d, 0x0a, 0x0f, 0x47, 0x65, 0x74,
	0x41, 0x70, 0x70, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x34, 0x0a, 0x05,
	0x49, 0x6e, 0x66, 0x6f, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x61, 0x70,
	0x70, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65,
	0x2e, 0x61, 0x70, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x70, 0x70, 0x52, 0x05, 0x49, 0x6e, 0x66,
	0x6f, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x14, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x22, 0x66, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x55,
	0x73, 0x65, 0x72, 0x41, 0x70, 0x70, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x22,
	0x0a, 0x0c, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72,
	0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x14, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x06, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x4c, 0x69,
	0x6d, 0x69, 0x74, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x4c, 0x69, 0x6d, 0x69, 0x74,
	0x22, 0x61, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x41, 0x70, 0x70, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x34, 0x0a, 0x05, 0x49, 0x6e, 0x66, 0x6f, 0x73,
	0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x61, 0x70, 0x70, 0x75, 0x73, 0x65, 0x72,
	0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x61, 0x70, 0x70, 0x2e,
	0x76, 0x31, 0x2e, 0x41, 0x70, 0x70, 0x52, 0x05, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x12, 0x14, 0x0a,
	0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x14, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x54, 0x6f,
	0x74, 0x61, 0x6c, 0x32, 0xee, 0x04, 0x0a, 0x05, 0x41, 0x70, 0x70, 0x47, 0x77, 0x12, 0x7b, 0x0a,
	0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x70, 0x70, 0x12, 0x28, 0x2e, 0x61, 0x70, 0x70,
	0x75, 0x73, 0x65, 0x72, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x61, 0x70, 0x70,
	0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x70, 0x70, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x61, 0x70, 0x70, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x67,
	0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x41, 0x70, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x19, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x13, 0x22, 0x0e, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x2f, 0x61, 0x70, 0x70, 0x3a, 0x01, 0x2a, 0x12, 0x7b, 0x0a, 0x09, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x41, 0x70, 0x70, 0x12, 0x28, 0x2e, 0x61, 0x70, 0x70, 0x75, 0x73, 0x65,
	0x72, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x76, 0x31,
	0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x70, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x29, 0x2e, 0x61, 0x70, 0x70, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x67, 0x61, 0x74, 0x65,
	0x77, 0x61, 0x79, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x41, 0x70, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x19, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x13, 0x22, 0x0e, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x2f, 0x61, 0x70, 0x70, 0x3a, 0x01, 0x2a, 0x12, 0x6f, 0x0a, 0x06, 0x47, 0x65, 0x74, 0x41, 0x70,
	0x70, 0x12, 0x25, 0x2e, 0x61, 0x70, 0x70, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x67, 0x61, 0x74, 0x65,
	0x77, 0x61, 0x79, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x70,
	0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x61, 0x70, 0x70, 0x75, 0x73,
	0x65, 0x72, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x76,
	0x31, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x70, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x16, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x10, 0x22, 0x0b, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x65,
	0x74, 0x2f, 0x61, 0x70, 0x70, 0x3a, 0x01, 0x2a, 0x12, 0x73, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x41,
	0x70, 0x70, 0x73, 0x12, 0x26, 0x2e, 0x61, 0x70, 0x70, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x67, 0x61,
	0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74,
	0x41, 0x70, 0x70, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x61, 0x70,
	0x70, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x61, 0x70,
	0x70, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x70, 0x70, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x17, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x11, 0x22, 0x0c, 0x2f, 0x76,
	0x31, 0x2f, 0x67, 0x65, 0x74, 0x2f, 0x61, 0x70, 0x70, 0x73, 0x3a, 0x01, 0x2a, 0x12, 0x84, 0x01,
	0x0a, 0x0b, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x41, 0x70, 0x70, 0x73, 0x12, 0x2a, 0x2e,
	0x61, 0x70, 0x70, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e,
	0x61, 0x70, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x41, 0x70,
	0x70, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2b, 0x2e, 0x61, 0x70, 0x70, 0x75,
	0x73, 0x65, 0x72, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x61, 0x70, 0x70, 0x2e,
	0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x41, 0x70, 0x70, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x16, 0x22, 0x11,
	0x2f, 0x76, 0x31, 0x2f, 0x67, 0x65, 0x74, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x70,
	0x73, 0x3a, 0x01, 0x2a, 0x42, 0x3a, 0x5a, 0x38, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x4e, 0x70, 0x6f, 0x6f, 0x6c, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d,
	0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x61,
	0x70, 0x70, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x67, 0x77, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x70, 0x70,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_npool_appuser_gw_v1_app_app_proto_rawDescOnce sync.Once
	file_npool_appuser_gw_v1_app_app_proto_rawDescData = file_npool_appuser_gw_v1_app_app_proto_rawDesc
)

func file_npool_appuser_gw_v1_app_app_proto_rawDescGZIP() []byte {
	file_npool_appuser_gw_v1_app_app_proto_rawDescOnce.Do(func() {
		file_npool_appuser_gw_v1_app_app_proto_rawDescData = protoimpl.X.CompressGZIP(file_npool_appuser_gw_v1_app_app_proto_rawDescData)
	})
	return file_npool_appuser_gw_v1_app_app_proto_rawDescData
}

var file_npool_appuser_gw_v1_app_app_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_npool_appuser_gw_v1_app_app_proto_goTypes = []interface{}{
	(*CreateAppRequest)(nil),    // 0: appuser.gateway.app.v1.CreateAppRequest
	(*CreateAppResponse)(nil),   // 1: appuser.gateway.app.v1.CreateAppResponse
	(*UpdateAppRequest)(nil),    // 2: appuser.gateway.app.v1.UpdateAppRequest
	(*UpdateAppResponse)(nil),   // 3: appuser.gateway.app.v1.UpdateAppResponse
	(*GetAppRequest)(nil),       // 4: appuser.gateway.app.v1.GetAppRequest
	(*GetAppResponse)(nil),      // 5: appuser.gateway.app.v1.GetAppResponse
	(*GetAppsRequest)(nil),      // 6: appuser.gateway.app.v1.GetAppsRequest
	(*GetAppsResponse)(nil),     // 7: appuser.gateway.app.v1.GetAppsResponse
	(*GetUserAppsRequest)(nil),  // 8: appuser.gateway.app.v1.GetUserAppsRequest
	(*GetUserAppsResponse)(nil), // 9: appuser.gateway.app.v1.GetUserAppsResponse
	(*app.AppReq)(nil),          // 10: appuser.manager.app.v2.AppReq
	(*app.App)(nil),             // 11: appuser.manager.app.v2.App
	(*app1.App)(nil),            // 12: appuser.middleware.app.v1.App
}
var file_npool_appuser_gw_v1_app_app_proto_depIdxs = []int32{
	10, // 0: appuser.gateway.app.v1.CreateAppRequest.Info:type_name -> appuser.manager.app.v2.AppReq
	11, // 1: appuser.gateway.app.v1.CreateAppResponse.Info:type_name -> appuser.manager.app.v2.App
	10, // 2: appuser.gateway.app.v1.UpdateAppRequest.Info:type_name -> appuser.manager.app.v2.AppReq
	11, // 3: appuser.gateway.app.v1.UpdateAppResponse.Info:type_name -> appuser.manager.app.v2.App
	12, // 4: appuser.gateway.app.v1.GetAppResponse.Info:type_name -> appuser.middleware.app.v1.App
	12, // 5: appuser.gateway.app.v1.GetAppsResponse.Infos:type_name -> appuser.middleware.app.v1.App
	12, // 6: appuser.gateway.app.v1.GetUserAppsResponse.Infos:type_name -> appuser.middleware.app.v1.App
	0,  // 7: appuser.gateway.app.v1.AppGw.CreateApp:input_type -> appuser.gateway.app.v1.CreateAppRequest
	2,  // 8: appuser.gateway.app.v1.AppGw.UpdateApp:input_type -> appuser.gateway.app.v1.UpdateAppRequest
	4,  // 9: appuser.gateway.app.v1.AppGw.GetApp:input_type -> appuser.gateway.app.v1.GetAppRequest
	6,  // 10: appuser.gateway.app.v1.AppGw.GetApps:input_type -> appuser.gateway.app.v1.GetAppsRequest
	8,  // 11: appuser.gateway.app.v1.AppGw.GetUserApps:input_type -> appuser.gateway.app.v1.GetUserAppsRequest
	1,  // 12: appuser.gateway.app.v1.AppGw.CreateApp:output_type -> appuser.gateway.app.v1.CreateAppResponse
	3,  // 13: appuser.gateway.app.v1.AppGw.UpdateApp:output_type -> appuser.gateway.app.v1.UpdateAppResponse
	5,  // 14: appuser.gateway.app.v1.AppGw.GetApp:output_type -> appuser.gateway.app.v1.GetAppResponse
	7,  // 15: appuser.gateway.app.v1.AppGw.GetApps:output_type -> appuser.gateway.app.v1.GetAppsResponse
	9,  // 16: appuser.gateway.app.v1.AppGw.GetUserApps:output_type -> appuser.gateway.app.v1.GetUserAppsResponse
	12, // [12:17] is the sub-list for method output_type
	7,  // [7:12] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_npool_appuser_gw_v1_app_app_proto_init() }
func file_npool_appuser_gw_v1_app_app_proto_init() {
	if File_npool_appuser_gw_v1_app_app_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_npool_appuser_gw_v1_app_app_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateAppRequest); i {
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
		file_npool_appuser_gw_v1_app_app_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateAppResponse); i {
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
		file_npool_appuser_gw_v1_app_app_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateAppRequest); i {
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
		file_npool_appuser_gw_v1_app_app_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateAppResponse); i {
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
		file_npool_appuser_gw_v1_app_app_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAppRequest); i {
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
		file_npool_appuser_gw_v1_app_app_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAppResponse); i {
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
		file_npool_appuser_gw_v1_app_app_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAppsRequest); i {
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
		file_npool_appuser_gw_v1_app_app_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAppsResponse); i {
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
		file_npool_appuser_gw_v1_app_app_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserAppsRequest); i {
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
		file_npool_appuser_gw_v1_app_app_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserAppsResponse); i {
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
			RawDescriptor: file_npool_appuser_gw_v1_app_app_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_npool_appuser_gw_v1_app_app_proto_goTypes,
		DependencyIndexes: file_npool_appuser_gw_v1_app_app_proto_depIdxs,
		MessageInfos:      file_npool_appuser_gw_v1_app_app_proto_msgTypes,
	}.Build()
	File_npool_appuser_gw_v1_app_app_proto = out.File
	file_npool_appuser_gw_v1_app_app_proto_rawDesc = nil
	file_npool_appuser_gw_v1_app_app_proto_goTypes = nil
	file_npool_appuser_gw_v1_app_app_proto_depIdxs = nil
}
