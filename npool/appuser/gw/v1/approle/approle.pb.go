// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.18.1
// source: npool/appuser/gw/v1/approle/approle.proto

package approle

import (
	approle "github.com/NpoolPlatform/message/npool/appuser/mgr/v2/approle"
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

type CreateRoleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID string              `protobuf:"bytes,10,opt,name=UserID,proto3" json:"UserID,omitempty"`
	Info   *approle.AppRoleReq `protobuf:"bytes,20,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (x *CreateRoleRequest) Reset() {
	*x = CreateRoleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_appuser_gw_v1_approle_approle_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateRoleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRoleRequest) ProtoMessage() {}

func (x *CreateRoleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_appuser_gw_v1_approle_approle_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRoleRequest.ProtoReflect.Descriptor instead.
func (*CreateRoleRequest) Descriptor() ([]byte, []int) {
	return file_npool_appuser_gw_v1_approle_approle_proto_rawDescGZIP(), []int{0}
}

func (x *CreateRoleRequest) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

func (x *CreateRoleRequest) GetInfo() *approle.AppRoleReq {
	if x != nil {
		return x.Info
	}
	return nil
}

type CreateRoleResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *approle.AppRole `protobuf:"bytes,10,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (x *CreateRoleResponse) Reset() {
	*x = CreateRoleResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_appuser_gw_v1_approle_approle_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateRoleResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRoleResponse) ProtoMessage() {}

func (x *CreateRoleResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_appuser_gw_v1_approle_approle_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRoleResponse.ProtoReflect.Descriptor instead.
func (*CreateRoleResponse) Descriptor() ([]byte, []int) {
	return file_npool_appuser_gw_v1_approle_approle_proto_rawDescGZIP(), []int{1}
}

func (x *CreateRoleResponse) GetInfo() *approle.AppRole {
	if x != nil {
		return x.Info
	}
	return nil
}

type CreateAppRoleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TargetAppID string              `protobuf:"bytes,10,opt,name=TargetAppID,proto3" json:"TargetAppID,omitempty"`
	UserID      string              `protobuf:"bytes,20,opt,name=UserID,proto3" json:"UserID,omitempty"`
	Info        *approle.AppRoleReq `protobuf:"bytes,30,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (x *CreateAppRoleRequest) Reset() {
	*x = CreateAppRoleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_appuser_gw_v1_approle_approle_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateAppRoleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAppRoleRequest) ProtoMessage() {}

func (x *CreateAppRoleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_appuser_gw_v1_approle_approle_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAppRoleRequest.ProtoReflect.Descriptor instead.
func (*CreateAppRoleRequest) Descriptor() ([]byte, []int) {
	return file_npool_appuser_gw_v1_approle_approle_proto_rawDescGZIP(), []int{2}
}

func (x *CreateAppRoleRequest) GetTargetAppID() string {
	if x != nil {
		return x.TargetAppID
	}
	return ""
}

func (x *CreateAppRoleRequest) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

func (x *CreateAppRoleRequest) GetInfo() *approle.AppRoleReq {
	if x != nil {
		return x.Info
	}
	return nil
}

type CreateAppRoleResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *approle.AppRole `protobuf:"bytes,10,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (x *CreateAppRoleResponse) Reset() {
	*x = CreateAppRoleResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_appuser_gw_v1_approle_approle_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateAppRoleResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAppRoleResponse) ProtoMessage() {}

func (x *CreateAppRoleResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_appuser_gw_v1_approle_approle_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAppRoleResponse.ProtoReflect.Descriptor instead.
func (*CreateAppRoleResponse) Descriptor() ([]byte, []int) {
	return file_npool_appuser_gw_v1_approle_approle_proto_rawDescGZIP(), []int{3}
}

func (x *CreateAppRoleResponse) GetInfo() *approle.AppRole {
	if x != nil {
		return x.Info
	}
	return nil
}

type GetRolesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AppID  string `protobuf:"bytes,10,opt,name=AppID,proto3" json:"AppID,omitempty"`
	Limit  int32  `protobuf:"varint,20,opt,name=Limit,proto3" json:"Limit,omitempty"`
	Offset int32  `protobuf:"varint,30,opt,name=Offset,proto3" json:"Offset,omitempty"`
}

func (x *GetRolesRequest) Reset() {
	*x = GetRolesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_appuser_gw_v1_approle_approle_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRolesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRolesRequest) ProtoMessage() {}

func (x *GetRolesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_appuser_gw_v1_approle_approle_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRolesRequest.ProtoReflect.Descriptor instead.
func (*GetRolesRequest) Descriptor() ([]byte, []int) {
	return file_npool_appuser_gw_v1_approle_approle_proto_rawDescGZIP(), []int{4}
}

func (x *GetRolesRequest) GetAppID() string {
	if x != nil {
		return x.AppID
	}
	return ""
}

func (x *GetRolesRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *GetRolesRequest) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

type GetRolesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Infos []*approle.AppRole `protobuf:"bytes,10,rep,name=Infos,proto3" json:"Infos,omitempty"`
	Total uint32             `protobuf:"varint,20,opt,name=Total,proto3" json:"Total,omitempty"`
}

func (x *GetRolesResponse) Reset() {
	*x = GetRolesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_appuser_gw_v1_approle_approle_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRolesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRolesResponse) ProtoMessage() {}

func (x *GetRolesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_appuser_gw_v1_approle_approle_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRolesResponse.ProtoReflect.Descriptor instead.
func (*GetRolesResponse) Descriptor() ([]byte, []int) {
	return file_npool_appuser_gw_v1_approle_approle_proto_rawDescGZIP(), []int{5}
}

func (x *GetRolesResponse) GetInfos() []*approle.AppRole {
	if x != nil {
		return x.Infos
	}
	return nil
}

func (x *GetRolesResponse) GetTotal() uint32 {
	if x != nil {
		return x.Total
	}
	return 0
}

type GetAppRolesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TargetAppID string `protobuf:"bytes,10,opt,name=TargetAppID,proto3" json:"TargetAppID,omitempty"`
	Limit       int32  `protobuf:"varint,20,opt,name=Limit,proto3" json:"Limit,omitempty"`
	Offset      int32  `protobuf:"varint,30,opt,name=Offset,proto3" json:"Offset,omitempty"`
}

func (x *GetAppRolesRequest) Reset() {
	*x = GetAppRolesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_appuser_gw_v1_approle_approle_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAppRolesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAppRolesRequest) ProtoMessage() {}

func (x *GetAppRolesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_appuser_gw_v1_approle_approle_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAppRolesRequest.ProtoReflect.Descriptor instead.
func (*GetAppRolesRequest) Descriptor() ([]byte, []int) {
	return file_npool_appuser_gw_v1_approle_approle_proto_rawDescGZIP(), []int{6}
}

func (x *GetAppRolesRequest) GetTargetAppID() string {
	if x != nil {
		return x.TargetAppID
	}
	return ""
}

func (x *GetAppRolesRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *GetAppRolesRequest) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

type GetAppRolesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Infos []*approle.AppRole `protobuf:"bytes,10,rep,name=Infos,proto3" json:"Infos,omitempty"`
	Total uint32             `protobuf:"varint,20,opt,name=Total,proto3" json:"Total,omitempty"`
}

func (x *GetAppRolesResponse) Reset() {
	*x = GetAppRolesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_appuser_gw_v1_approle_approle_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAppRolesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAppRolesResponse) ProtoMessage() {}

func (x *GetAppRolesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_appuser_gw_v1_approle_approle_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAppRolesResponse.ProtoReflect.Descriptor instead.
func (*GetAppRolesResponse) Descriptor() ([]byte, []int) {
	return file_npool_appuser_gw_v1_approle_approle_proto_rawDescGZIP(), []int{7}
}

func (x *GetAppRolesResponse) GetInfos() []*approle.AppRole {
	if x != nil {
		return x.Infos
	}
	return nil
}

func (x *GetAppRolesResponse) GetTotal() uint32 {
	if x != nil {
		return x.Total
	}
	return 0
}

type UpdateRoleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *approle.AppRoleReq `protobuf:"bytes,10,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (x *UpdateRoleRequest) Reset() {
	*x = UpdateRoleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_appuser_gw_v1_approle_approle_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateRoleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateRoleRequest) ProtoMessage() {}

func (x *UpdateRoleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_appuser_gw_v1_approle_approle_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateRoleRequest.ProtoReflect.Descriptor instead.
func (*UpdateRoleRequest) Descriptor() ([]byte, []int) {
	return file_npool_appuser_gw_v1_approle_approle_proto_rawDescGZIP(), []int{8}
}

func (x *UpdateRoleRequest) GetInfo() *approle.AppRoleReq {
	if x != nil {
		return x.Info
	}
	return nil
}

type UpdateRoleResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *approle.AppRole `protobuf:"bytes,10,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (x *UpdateRoleResponse) Reset() {
	*x = UpdateRoleResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_appuser_gw_v1_approle_approle_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateRoleResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateRoleResponse) ProtoMessage() {}

func (x *UpdateRoleResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_appuser_gw_v1_approle_approle_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateRoleResponse.ProtoReflect.Descriptor instead.
func (*UpdateRoleResponse) Descriptor() ([]byte, []int) {
	return file_npool_appuser_gw_v1_approle_approle_proto_rawDescGZIP(), []int{9}
}

func (x *UpdateRoleResponse) GetInfo() *approle.AppRole {
	if x != nil {
		return x.Info
	}
	return nil
}

var File_npool_appuser_gw_v1_approle_approle_proto protoreflect.FileDescriptor

var file_npool_appuser_gw_v1_approle_approle_proto_rawDesc = []byte{
	0x0a, 0x29, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x61, 0x70, 0x70, 0x75, 0x73, 0x65, 0x72, 0x2f,
	0x67, 0x77, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x70, 0x70, 0x72, 0x6f, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x70, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1a, 0x61, 0x70, 0x70,
	0x75, 0x73, 0x65, 0x72, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x61, 0x70, 0x70,
	0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2a, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x61, 0x70, 0x70,
	0x75, 0x73, 0x65, 0x72, 0x2f, 0x6d, 0x67, 0x72, 0x2f, 0x76, 0x32, 0x2f, 0x61, 0x70, 0x70, 0x72,
	0x6f, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x70, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x67, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x3a,
	0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x14, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x61,
	0x70, 0x70, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x61,
	0x70, 0x70, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x76, 0x32, 0x2e, 0x41, 0x70, 0x70, 0x52, 0x6f, 0x6c,
	0x65, 0x52, 0x65, 0x71, 0x52, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x4d, 0x0a, 0x12, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x37, 0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23,
	0x2e, 0x61, 0x70, 0x70, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72,
	0x2e, 0x61, 0x70, 0x70, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x76, 0x32, 0x2e, 0x41, 0x70, 0x70, 0x52,
	0x6f, 0x6c, 0x65, 0x52, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x8c, 0x01, 0x0a, 0x14, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x41, 0x70, 0x70, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x41, 0x70, 0x70, 0x49,
	0x44, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x41,
	0x70, 0x70, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x14,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x3a, 0x0a, 0x04,
	0x49, 0x6e, 0x66, 0x6f, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x61, 0x70, 0x70,
	0x75, 0x73, 0x65, 0x72, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x70,
	0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x76, 0x32, 0x2e, 0x41, 0x70, 0x70, 0x52, 0x6f, 0x6c, 0x65, 0x52,
	0x65, 0x71, 0x52, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x50, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x41, 0x70, 0x70, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x37, 0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x23, 0x2e, 0x61, 0x70, 0x70, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65,
	0x72, 0x2e, 0x61, 0x70, 0x70, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x76, 0x32, 0x2e, 0x41, 0x70, 0x70,
	0x52, 0x6f, 0x6c, 0x65, 0x52, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x55, 0x0a, 0x0f, 0x47, 0x65,
	0x74, 0x52, 0x6f, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a,
	0x05, 0x41, 0x70, 0x70, 0x49, 0x44, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x41, 0x70,
	0x70, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x14, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x05, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x4f, 0x66, 0x66,
	0x73, 0x65, 0x74, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x4f, 0x66, 0x66, 0x73, 0x65,
	0x74, 0x22, 0x63, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x52, 0x6f, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x39, 0x0a, 0x05, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x18, 0x0a,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x61, 0x70, 0x70, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x6d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x70, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x76,
	0x32, 0x2e, 0x41, 0x70, 0x70, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x05, 0x49, 0x6e, 0x66, 0x6f, 0x73,
	0x12, 0x14, 0x0a, 0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x14, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x22, 0x64, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x41, 0x70, 0x70,
	0x52, 0x6f, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x0b,
	0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x41, 0x70, 0x70, 0x49, 0x44, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x41, 0x70, 0x70, 0x49, 0x44, 0x12, 0x14,
	0x0a, 0x05, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x14, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x4c,
	0x69, 0x6d, 0x69, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x1e,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x22, 0x66, 0x0a, 0x13,
	0x47, 0x65, 0x74, 0x41, 0x70, 0x70, 0x52, 0x6f, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x39, 0x0a, 0x05, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x18, 0x0a, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x23, 0x2e, 0x61, 0x70, 0x70, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x6d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x70, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x76, 0x32, 0x2e,
	0x41, 0x70, 0x70, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x05, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x12, 0x14,
	0x0a, 0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x14, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x54,
	0x6f, 0x74, 0x61, 0x6c, 0x22, 0x4f, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x6f,
	0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3a, 0x0a, 0x04, 0x49, 0x6e, 0x66,
	0x6f, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x61, 0x70, 0x70, 0x75, 0x73, 0x65,
	0x72, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x70, 0x72, 0x6f, 0x6c,
	0x65, 0x2e, 0x76, 0x32, 0x2e, 0x41, 0x70, 0x70, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x52,
	0x04, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x4d, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52,
	0x6f, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x37, 0x0a, 0x04, 0x49,
	0x6e, 0x66, 0x6f, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x61, 0x70, 0x70, 0x75,
	0x73, 0x65, 0x72, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x70, 0x72,
	0x6f, 0x6c, 0x65, 0x2e, 0x76, 0x32, 0x2e, 0x41, 0x70, 0x70, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x04,
	0x49, 0x6e, 0x66, 0x6f, 0x32, 0xca, 0x05, 0x0a, 0x09, 0x41, 0x70, 0x70, 0x52, 0x6f, 0x6c, 0x65,
	0x47, 0x77, 0x12, 0x87, 0x01, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x6c,
	0x65, 0x12, 0x2d, 0x2e, 0x61, 0x70, 0x70, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x67, 0x61, 0x74, 0x65,
	0x77, 0x61, 0x79, 0x2e, 0x61, 0x70, 0x70, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x2e, 0x2e, 0x61, 0x70, 0x70, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77,
	0x61, 0x79, 0x2e, 0x61, 0x70, 0x70, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x1a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14, 0x22, 0x0f, 0x2f, 0x76, 0x32, 0x2f, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x2f, 0x72, 0x6f, 0x6c, 0x65, 0x3a, 0x01, 0x2a, 0x12, 0x94, 0x01, 0x0a,
	0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x70, 0x70, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x30,
	0x2e, 0x61, 0x70, 0x70, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79,
	0x2e, 0x61, 0x70, 0x70, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x41, 0x70, 0x70, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x31, 0x2e, 0x61, 0x70, 0x70, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77,
	0x61, 0x79, 0x2e, 0x61, 0x70, 0x70, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x41, 0x70, 0x70, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x1e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x18, 0x22, 0x13, 0x2f, 0x76, 0x32,
	0x2f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x2f, 0x61, 0x70, 0x70, 0x2f, 0x72, 0x6f, 0x6c, 0x65,
	0x3a, 0x01, 0x2a, 0x12, 0x7f, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x52, 0x6f, 0x6c, 0x65, 0x73, 0x12,
	0x2b, 0x2e, 0x61, 0x70, 0x70, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61,
	0x79, 0x2e, 0x61, 0x70, 0x70, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74,
	0x52, 0x6f, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2c, 0x2e, 0x61,
	0x70, 0x70, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x61,
	0x70, 0x70, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x6f, 0x6c,
	0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x18, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x12, 0x22, 0x0d, 0x2f, 0x76, 0x32, 0x2f, 0x67, 0x65, 0x74, 0x2f, 0x72, 0x6f, 0x6c, 0x65,
	0x73, 0x3a, 0x01, 0x2a, 0x12, 0x8c, 0x01, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x41, 0x70, 0x70, 0x52,
	0x6f, 0x6c, 0x65, 0x73, 0x12, 0x2e, 0x2e, 0x61, 0x70, 0x70, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x67,
	0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x61, 0x70, 0x70, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x70, 0x70, 0x52, 0x6f, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x2f, 0x2e, 0x61, 0x70, 0x70, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x67,
	0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x61, 0x70, 0x70, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x70, 0x70, 0x52, 0x6f, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x16, 0x22, 0x11, 0x2f,
	0x76, 0x32, 0x2f, 0x67, 0x65, 0x74, 0x2f, 0x61, 0x70, 0x70, 0x2f, 0x72, 0x6f, 0x6c, 0x65, 0x73,
	0x3a, 0x01, 0x2a, 0x12, 0x8b, 0x01, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x6f,
	0x6c, 0x65, 0x12, 0x2d, 0x2e, 0x61, 0x70, 0x70, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x67, 0x61, 0x74,
	0x65, 0x77, 0x61, 0x79, 0x2e, 0x61, 0x70, 0x70, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x2e, 0x2e, 0x61, 0x70, 0x70, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x67, 0x61, 0x74, 0x65,
	0x77, 0x61, 0x79, 0x2e, 0x61, 0x70, 0x70, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x1e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x18, 0x22, 0x13, 0x2f, 0x76, 0x32, 0x2f, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x61, 0x70, 0x70, 0x2f, 0x72, 0x6f, 0x6c, 0x65, 0x3a, 0x01,
	0x2a, 0x42, 0x3e, 0x5a, 0x3c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x4e, 0x70, 0x6f, 0x6f, 0x6c, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x61, 0x70, 0x70, 0x75,
	0x73, 0x65, 0x72, 0x2f, 0x67, 0x77, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x70, 0x70, 0x72, 0x6f, 0x6c,
	0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_npool_appuser_gw_v1_approle_approle_proto_rawDescOnce sync.Once
	file_npool_appuser_gw_v1_approle_approle_proto_rawDescData = file_npool_appuser_gw_v1_approle_approle_proto_rawDesc
)

func file_npool_appuser_gw_v1_approle_approle_proto_rawDescGZIP() []byte {
	file_npool_appuser_gw_v1_approle_approle_proto_rawDescOnce.Do(func() {
		file_npool_appuser_gw_v1_approle_approle_proto_rawDescData = protoimpl.X.CompressGZIP(file_npool_appuser_gw_v1_approle_approle_proto_rawDescData)
	})
	return file_npool_appuser_gw_v1_approle_approle_proto_rawDescData
}

var file_npool_appuser_gw_v1_approle_approle_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_npool_appuser_gw_v1_approle_approle_proto_goTypes = []interface{}{
	(*CreateRoleRequest)(nil),     // 0: appuser.gateway.approle.v1.CreateRoleRequest
	(*CreateRoleResponse)(nil),    // 1: appuser.gateway.approle.v1.CreateRoleResponse
	(*CreateAppRoleRequest)(nil),  // 2: appuser.gateway.approle.v1.CreateAppRoleRequest
	(*CreateAppRoleResponse)(nil), // 3: appuser.gateway.approle.v1.CreateAppRoleResponse
	(*GetRolesRequest)(nil),       // 4: appuser.gateway.approle.v1.GetRolesRequest
	(*GetRolesResponse)(nil),      // 5: appuser.gateway.approle.v1.GetRolesResponse
	(*GetAppRolesRequest)(nil),    // 6: appuser.gateway.approle.v1.GetAppRolesRequest
	(*GetAppRolesResponse)(nil),   // 7: appuser.gateway.approle.v1.GetAppRolesResponse
	(*UpdateRoleRequest)(nil),     // 8: appuser.gateway.approle.v1.UpdateRoleRequest
	(*UpdateRoleResponse)(nil),    // 9: appuser.gateway.approle.v1.UpdateRoleResponse
	(*approle.AppRoleReq)(nil),    // 10: appuser.manager.approle.v2.AppRoleReq
	(*approle.AppRole)(nil),       // 11: appuser.manager.approle.v2.AppRole
}
var file_npool_appuser_gw_v1_approle_approle_proto_depIdxs = []int32{
	10, // 0: appuser.gateway.approle.v1.CreateRoleRequest.Info:type_name -> appuser.manager.approle.v2.AppRoleReq
	11, // 1: appuser.gateway.approle.v1.CreateRoleResponse.Info:type_name -> appuser.manager.approle.v2.AppRole
	10, // 2: appuser.gateway.approle.v1.CreateAppRoleRequest.Info:type_name -> appuser.manager.approle.v2.AppRoleReq
	11, // 3: appuser.gateway.approle.v1.CreateAppRoleResponse.Info:type_name -> appuser.manager.approle.v2.AppRole
	11, // 4: appuser.gateway.approle.v1.GetRolesResponse.Infos:type_name -> appuser.manager.approle.v2.AppRole
	11, // 5: appuser.gateway.approle.v1.GetAppRolesResponse.Infos:type_name -> appuser.manager.approle.v2.AppRole
	10, // 6: appuser.gateway.approle.v1.UpdateRoleRequest.Info:type_name -> appuser.manager.approle.v2.AppRoleReq
	11, // 7: appuser.gateway.approle.v1.UpdateRoleResponse.Info:type_name -> appuser.manager.approle.v2.AppRole
	0,  // 8: appuser.gateway.approle.v1.AppRoleGw.CreateRole:input_type -> appuser.gateway.approle.v1.CreateRoleRequest
	2,  // 9: appuser.gateway.approle.v1.AppRoleGw.CreateAppRole:input_type -> appuser.gateway.approle.v1.CreateAppRoleRequest
	4,  // 10: appuser.gateway.approle.v1.AppRoleGw.GetRoles:input_type -> appuser.gateway.approle.v1.GetRolesRequest
	6,  // 11: appuser.gateway.approle.v1.AppRoleGw.GetAppRoles:input_type -> appuser.gateway.approle.v1.GetAppRolesRequest
	8,  // 12: appuser.gateway.approle.v1.AppRoleGw.UpdateRole:input_type -> appuser.gateway.approle.v1.UpdateRoleRequest
	1,  // 13: appuser.gateway.approle.v1.AppRoleGw.CreateRole:output_type -> appuser.gateway.approle.v1.CreateRoleResponse
	3,  // 14: appuser.gateway.approle.v1.AppRoleGw.CreateAppRole:output_type -> appuser.gateway.approle.v1.CreateAppRoleResponse
	5,  // 15: appuser.gateway.approle.v1.AppRoleGw.GetRoles:output_type -> appuser.gateway.approle.v1.GetRolesResponse
	7,  // 16: appuser.gateway.approle.v1.AppRoleGw.GetAppRoles:output_type -> appuser.gateway.approle.v1.GetAppRolesResponse
	9,  // 17: appuser.gateway.approle.v1.AppRoleGw.UpdateRole:output_type -> appuser.gateway.approle.v1.UpdateRoleResponse
	13, // [13:18] is the sub-list for method output_type
	8,  // [8:13] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_npool_appuser_gw_v1_approle_approle_proto_init() }
func file_npool_appuser_gw_v1_approle_approle_proto_init() {
	if File_npool_appuser_gw_v1_approle_approle_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_npool_appuser_gw_v1_approle_approle_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateRoleRequest); i {
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
		file_npool_appuser_gw_v1_approle_approle_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateRoleResponse); i {
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
		file_npool_appuser_gw_v1_approle_approle_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateAppRoleRequest); i {
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
		file_npool_appuser_gw_v1_approle_approle_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateAppRoleResponse); i {
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
		file_npool_appuser_gw_v1_approle_approle_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRolesRequest); i {
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
		file_npool_appuser_gw_v1_approle_approle_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRolesResponse); i {
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
		file_npool_appuser_gw_v1_approle_approle_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAppRolesRequest); i {
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
		file_npool_appuser_gw_v1_approle_approle_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAppRolesResponse); i {
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
		file_npool_appuser_gw_v1_approle_approle_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateRoleRequest); i {
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
		file_npool_appuser_gw_v1_approle_approle_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateRoleResponse); i {
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
			RawDescriptor: file_npool_appuser_gw_v1_approle_approle_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_npool_appuser_gw_v1_approle_approle_proto_goTypes,
		DependencyIndexes: file_npool_appuser_gw_v1_approle_approle_proto_depIdxs,
		MessageInfos:      file_npool_appuser_gw_v1_approle_approle_proto_msgTypes,
	}.Build()
	File_npool_appuser_gw_v1_approle_approle_proto = out.File
	file_npool_appuser_gw_v1_approle_approle_proto_rawDesc = nil
	file_npool_appuser_gw_v1_approle_approle_proto_goTypes = nil
	file_npool_appuser_gw_v1_approle_approle_proto_depIdxs = nil
}
