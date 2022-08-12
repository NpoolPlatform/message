// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.18.1
// source: npool/appuser/gw/v1/approleuser/approleuser.proto

package approleuser

import (
	approleuser "github.com/NpoolPlatform/message/npool/appuser/mgr/v2/approleuser"
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

type CreateRoleUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Permissioned by path
	TargetUserID string                      `protobuf:"bytes,10,opt,name=TargetUserID,proto3" json:"TargetUserID,omitempty"`
	Info         *approleuser.AppRoleUserReq `protobuf:"bytes,20,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (x *CreateRoleUserRequest) Reset() {
	*x = CreateRoleUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_appuser_gw_v1_approleuser_approleuser_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateRoleUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRoleUserRequest) ProtoMessage() {}

func (x *CreateRoleUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_appuser_gw_v1_approleuser_approleuser_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRoleUserRequest.ProtoReflect.Descriptor instead.
func (*CreateRoleUserRequest) Descriptor() ([]byte, []int) {
	return file_npool_appuser_gw_v1_approleuser_approleuser_proto_rawDescGZIP(), []int{0}
}

func (x *CreateRoleUserRequest) GetTargetUserID() string {
	if x != nil {
		return x.TargetUserID
	}
	return ""
}

func (x *CreateRoleUserRequest) GetInfo() *approleuser.AppRoleUserReq {
	if x != nil {
		return x.Info
	}
	return nil
}

type CreateRoleUserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *approleuser.AppRoleUser `protobuf:"bytes,10,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (x *CreateRoleUserResponse) Reset() {
	*x = CreateRoleUserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_appuser_gw_v1_approleuser_approleuser_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateRoleUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRoleUserResponse) ProtoMessage() {}

func (x *CreateRoleUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_appuser_gw_v1_approleuser_approleuser_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRoleUserResponse.ProtoReflect.Descriptor instead.
func (*CreateRoleUserResponse) Descriptor() ([]byte, []int) {
	return file_npool_appuser_gw_v1_approleuser_approleuser_proto_rawDescGZIP(), []int{1}
}

func (x *CreateRoleUserResponse) GetInfo() *approleuser.AppRoleUser {
	if x != nil {
		return x.Info
	}
	return nil
}

type GetRoleUsersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Permissioned by path
	AppID  string `protobuf:"bytes,10,opt,name=AppID,proto3" json:"AppID,omitempty"`
	RoleID string `protobuf:"bytes,20,opt,name=RoleID,proto3" json:"RoleID,omitempty"`
	Limit  int32  `protobuf:"varint,30,opt,name=Limit,proto3" json:"Limit,omitempty"`
	Offset int32  `protobuf:"varint,40,opt,name=Offset,proto3" json:"Offset,omitempty"`
}

func (x *GetRoleUsersRequest) Reset() {
	*x = GetRoleUsersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_appuser_gw_v1_approleuser_approleuser_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRoleUsersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRoleUsersRequest) ProtoMessage() {}

func (x *GetRoleUsersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_appuser_gw_v1_approleuser_approleuser_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRoleUsersRequest.ProtoReflect.Descriptor instead.
func (*GetRoleUsersRequest) Descriptor() ([]byte, []int) {
	return file_npool_appuser_gw_v1_approleuser_approleuser_proto_rawDescGZIP(), []int{2}
}

func (x *GetRoleUsersRequest) GetAppID() string {
	if x != nil {
		return x.AppID
	}
	return ""
}

func (x *GetRoleUsersRequest) GetRoleID() string {
	if x != nil {
		return x.RoleID
	}
	return ""
}

func (x *GetRoleUsersRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *GetRoleUsersRequest) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

type GetRoleUsersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Infos []*approleuser.AppRoleUser `protobuf:"bytes,10,rep,name=Infos,proto3" json:"Infos,omitempty"`
	Total uint32                     `protobuf:"varint,20,opt,name=Total,proto3" json:"Total,omitempty"`
}

func (x *GetRoleUsersResponse) Reset() {
	*x = GetRoleUsersResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_appuser_gw_v1_approleuser_approleuser_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRoleUsersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRoleUsersResponse) ProtoMessage() {}

func (x *GetRoleUsersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_appuser_gw_v1_approleuser_approleuser_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRoleUsersResponse.ProtoReflect.Descriptor instead.
func (*GetRoleUsersResponse) Descriptor() ([]byte, []int) {
	return file_npool_appuser_gw_v1_approleuser_approleuser_proto_rawDescGZIP(), []int{3}
}

func (x *GetRoleUsersResponse) GetInfos() []*approleuser.AppRoleUser {
	if x != nil {
		return x.Infos
	}
	return nil
}

func (x *GetRoleUsersResponse) GetTotal() uint32 {
	if x != nil {
		return x.Total
	}
	return 0
}

type DeleteRoleUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Permissioned by path
	ID string `protobuf:"bytes,10,opt,name=ID,proto3" json:"ID,omitempty"`
}

func (x *DeleteRoleUserRequest) Reset() {
	*x = DeleteRoleUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_appuser_gw_v1_approleuser_approleuser_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteRoleUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteRoleUserRequest) ProtoMessage() {}

func (x *DeleteRoleUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_appuser_gw_v1_approleuser_approleuser_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteRoleUserRequest.ProtoReflect.Descriptor instead.
func (*DeleteRoleUserRequest) Descriptor() ([]byte, []int) {
	return file_npool_appuser_gw_v1_approleuser_approleuser_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteRoleUserRequest) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

type DeleteRoleUserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *approleuser.AppRoleUser `protobuf:"bytes,10,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (x *DeleteRoleUserResponse) Reset() {
	*x = DeleteRoleUserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_appuser_gw_v1_approleuser_approleuser_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteRoleUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteRoleUserResponse) ProtoMessage() {}

func (x *DeleteRoleUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_appuser_gw_v1_approleuser_approleuser_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteRoleUserResponse.ProtoReflect.Descriptor instead.
func (*DeleteRoleUserResponse) Descriptor() ([]byte, []int) {
	return file_npool_appuser_gw_v1_approleuser_approleuser_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteRoleUserResponse) GetInfo() *approleuser.AppRoleUser {
	if x != nil {
		return x.Info
	}
	return nil
}

type CreateAppRoleUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Permissioned by path
	TargetAppID  string                      `protobuf:"bytes,10,opt,name=TargetAppID,proto3" json:"TargetAppID,omitempty"`
	TargetUserID string                      `protobuf:"bytes,20,opt,name=TargetUserID,proto3" json:"TargetUserID,omitempty"`
	Info         *approleuser.AppRoleUserReq `protobuf:"bytes,30,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (x *CreateAppRoleUserRequest) Reset() {
	*x = CreateAppRoleUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_appuser_gw_v1_approleuser_approleuser_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateAppRoleUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAppRoleUserRequest) ProtoMessage() {}

func (x *CreateAppRoleUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_appuser_gw_v1_approleuser_approleuser_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAppRoleUserRequest.ProtoReflect.Descriptor instead.
func (*CreateAppRoleUserRequest) Descriptor() ([]byte, []int) {
	return file_npool_appuser_gw_v1_approleuser_approleuser_proto_rawDescGZIP(), []int{6}
}

func (x *CreateAppRoleUserRequest) GetTargetAppID() string {
	if x != nil {
		return x.TargetAppID
	}
	return ""
}

func (x *CreateAppRoleUserRequest) GetTargetUserID() string {
	if x != nil {
		return x.TargetUserID
	}
	return ""
}

func (x *CreateAppRoleUserRequest) GetInfo() *approleuser.AppRoleUserReq {
	if x != nil {
		return x.Info
	}
	return nil
}

type CreateAppRoleUserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *approleuser.AppRoleUser `protobuf:"bytes,10,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (x *CreateAppRoleUserResponse) Reset() {
	*x = CreateAppRoleUserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_appuser_gw_v1_approleuser_approleuser_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateAppRoleUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAppRoleUserResponse) ProtoMessage() {}

func (x *CreateAppRoleUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_appuser_gw_v1_approleuser_approleuser_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAppRoleUserResponse.ProtoReflect.Descriptor instead.
func (*CreateAppRoleUserResponse) Descriptor() ([]byte, []int) {
	return file_npool_appuser_gw_v1_approleuser_approleuser_proto_rawDescGZIP(), []int{7}
}

func (x *CreateAppRoleUserResponse) GetInfo() *approleuser.AppRoleUser {
	if x != nil {
		return x.Info
	}
	return nil
}

type GetAppRoleUsersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Permissioned by path
	TargetAppID string `protobuf:"bytes,10,opt,name=TargetAppID,proto3" json:"TargetAppID,omitempty"`
	RoleID      string `protobuf:"bytes,20,opt,name=RoleID,proto3" json:"RoleID,omitempty"`
	Limit       int32  `protobuf:"varint,30,opt,name=Limit,proto3" json:"Limit,omitempty"`
	Offset      int32  `protobuf:"varint,40,opt,name=Offset,proto3" json:"Offset,omitempty"`
}

func (x *GetAppRoleUsersRequest) Reset() {
	*x = GetAppRoleUsersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_appuser_gw_v1_approleuser_approleuser_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAppRoleUsersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAppRoleUsersRequest) ProtoMessage() {}

func (x *GetAppRoleUsersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_appuser_gw_v1_approleuser_approleuser_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAppRoleUsersRequest.ProtoReflect.Descriptor instead.
func (*GetAppRoleUsersRequest) Descriptor() ([]byte, []int) {
	return file_npool_appuser_gw_v1_approleuser_approleuser_proto_rawDescGZIP(), []int{8}
}

func (x *GetAppRoleUsersRequest) GetTargetAppID() string {
	if x != nil {
		return x.TargetAppID
	}
	return ""
}

func (x *GetAppRoleUsersRequest) GetRoleID() string {
	if x != nil {
		return x.RoleID
	}
	return ""
}

func (x *GetAppRoleUsersRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *GetAppRoleUsersRequest) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

type GetAppRoleUsersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Infos []*approleuser.AppRoleUser `protobuf:"bytes,10,rep,name=Infos,proto3" json:"Infos,omitempty"`
	Total uint32                     `protobuf:"varint,20,opt,name=Total,proto3" json:"Total,omitempty"`
}

func (x *GetAppRoleUsersResponse) Reset() {
	*x = GetAppRoleUsersResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_appuser_gw_v1_approleuser_approleuser_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAppRoleUsersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAppRoleUsersResponse) ProtoMessage() {}

func (x *GetAppRoleUsersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_appuser_gw_v1_approleuser_approleuser_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAppRoleUsersResponse.ProtoReflect.Descriptor instead.
func (*GetAppRoleUsersResponse) Descriptor() ([]byte, []int) {
	return file_npool_appuser_gw_v1_approleuser_approleuser_proto_rawDescGZIP(), []int{9}
}

func (x *GetAppRoleUsersResponse) GetInfos() []*approleuser.AppRoleUser {
	if x != nil {
		return x.Infos
	}
	return nil
}

func (x *GetAppRoleUsersResponse) GetTotal() uint32 {
	if x != nil {
		return x.Total
	}
	return 0
}

var File_npool_appuser_gw_v1_approleuser_approleuser_proto protoreflect.FileDescriptor

var file_npool_appuser_gw_v1_approleuser_approleuser_proto_rawDesc = []byte{
	0x0a, 0x31, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x61, 0x70, 0x70, 0x75, 0x73, 0x65, 0x72, 0x2f,
	0x67, 0x77, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x70, 0x70, 0x72, 0x6f, 0x6c, 0x65, 0x75, 0x73, 0x65,
	0x72, 0x2f, 0x61, 0x70, 0x70, 0x72, 0x6f, 0x6c, 0x65, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x1e, 0x61, 0x70, 0x70, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x67, 0x61, 0x74,
	0x65, 0x77, 0x61, 0x79, 0x2e, 0x61, 0x70, 0x70, 0x72, 0x6f, 0x6c, 0x65, 0x75, 0x73, 0x65, 0x72,
	0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x32, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x61, 0x70, 0x70, 0x75, 0x73, 0x65, 0x72,
	0x2f, 0x6d, 0x67, 0x72, 0x2f, 0x76, 0x32, 0x2f, 0x61, 0x70, 0x70, 0x72, 0x6f, 0x6c, 0x65, 0x75,
	0x73, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x70, 0x72, 0x6f, 0x6c, 0x65, 0x75, 0x73, 0x65, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x7f, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52,
	0x6f, 0x6c, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x22,
	0x0a, 0x0c, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72,
	0x49, 0x44, 0x12, 0x42, 0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x14, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x2e, 0x2e, 0x61, 0x70, 0x70, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67,
	0x65, 0x72, 0x2e, 0x61, 0x70, 0x70, 0x72, 0x6f, 0x6c, 0x65, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76,
	0x32, 0x2e, 0x41, 0x70, 0x70, 0x52, 0x6f, 0x6c, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x52, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x59, 0x0a, 0x16, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x52, 0x6f, 0x6c, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x3f, 0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b,
	0x2e, 0x61, 0x70, 0x70, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72,
	0x2e, 0x61, 0x70, 0x70, 0x72, 0x6f, 0x6c, 0x65, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x32, 0x2e,
	0x41, 0x70, 0x70, 0x52, 0x6f, 0x6c, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x49, 0x6e, 0x66,
	0x6f, 0x22, 0x71, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x52, 0x6f, 0x6c, 0x65, 0x55, 0x73, 0x65, 0x72,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x41, 0x70, 0x70, 0x49,
	0x44, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x41, 0x70, 0x70, 0x49, 0x44, 0x12, 0x16,
	0x0a, 0x06, 0x52, 0x6f, 0x6c, 0x65, 0x49, 0x44, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x52, 0x6f, 0x6c, 0x65, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x18,
	0x1e, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x16, 0x0a, 0x06,
	0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x28, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x4f, 0x66,
	0x66, 0x73, 0x65, 0x74, 0x22, 0x6f, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x52, 0x6f, 0x6c, 0x65, 0x55,
	0x73, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x41, 0x0a, 0x05,
	0x49, 0x6e, 0x66, 0x6f, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x61, 0x70,
	0x70, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x61, 0x70,
	0x70, 0x72, 0x6f, 0x6c, 0x65, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x32, 0x2e, 0x41, 0x70, 0x70,
	0x52, 0x6f, 0x6c, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x05, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x12,
	0x14, 0x0a, 0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x14, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05,
	0x54, 0x6f, 0x74, 0x61, 0x6c, 0x22, 0x27, 0x0a, 0x15, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52,
	0x6f, 0x6c, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x49, 0x44, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x22, 0x59,
	0x0a, 0x16, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x55, 0x73, 0x65, 0x72,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3f, 0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x61, 0x70, 0x70, 0x75, 0x73, 0x65, 0x72,
	0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x70, 0x72, 0x6f, 0x6c, 0x65,
	0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x32, 0x2e, 0x41, 0x70, 0x70, 0x52, 0x6f, 0x6c, 0x65, 0x55,
	0x73, 0x65, 0x72, 0x52, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0xa4, 0x01, 0x0a, 0x18, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x41, 0x70, 0x70, 0x52, 0x6f, 0x6c, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74,
	0x41, 0x70, 0x70, 0x49, 0x44, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x54, 0x61, 0x72,
	0x67, 0x65, 0x74, 0x41, 0x70, 0x70, 0x49, 0x44, 0x12, 0x22, 0x0a, 0x0c, 0x54, 0x61, 0x72, 0x67,
	0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c,
	0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x42, 0x0a, 0x04,
	0x49, 0x6e, 0x66, 0x6f, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x61, 0x70, 0x70,
	0x75, 0x73, 0x65, 0x72, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x70,
	0x72, 0x6f, 0x6c, 0x65, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x32, 0x2e, 0x41, 0x70, 0x70, 0x52,
	0x6f, 0x6c, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x52, 0x04, 0x49, 0x6e, 0x66, 0x6f,
	0x22, 0x5c, 0x0a, 0x19, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x70, 0x70, 0x52, 0x6f, 0x6c,
	0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3f, 0x0a,
	0x04, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x61, 0x70,
	0x70, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x61, 0x70,
	0x70, 0x72, 0x6f, 0x6c, 0x65, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x32, 0x2e, 0x41, 0x70, 0x70,
	0x52, 0x6f, 0x6c, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x80,
	0x01, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x41, 0x70, 0x70, 0x52, 0x6f, 0x6c, 0x65, 0x55, 0x73, 0x65,
	0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x54, 0x61, 0x72,
	0x67, 0x65, 0x74, 0x41, 0x70, 0x70, 0x49, 0x44, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x41, 0x70, 0x70, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x52,
	0x6f, 0x6c, 0x65, 0x49, 0x44, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x52, 0x6f, 0x6c,
	0x65, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x1e, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x05, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x4f, 0x66, 0x66,
	0x73, 0x65, 0x74, 0x18, 0x28, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x4f, 0x66, 0x66, 0x73, 0x65,
	0x74, 0x22, 0x72, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x41, 0x70, 0x70, 0x52, 0x6f, 0x6c, 0x65, 0x55,
	0x73, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x41, 0x0a, 0x05,
	0x49, 0x6e, 0x66, 0x6f, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x61, 0x70,
	0x70, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x61, 0x70,
	0x70, 0x72, 0x6f, 0x6c, 0x65, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x32, 0x2e, 0x41, 0x70, 0x70,
	0x52, 0x6f, 0x6c, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x05, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x12,
	0x14, 0x0a, 0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x14, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05,
	0x54, 0x6f, 0x74, 0x61, 0x6c, 0x32, 0xc8, 0x06, 0x0a, 0x0d, 0x41, 0x70, 0x70, 0x52, 0x6f, 0x6c,
	0x65, 0x55, 0x73, 0x65, 0x72, 0x47, 0x77, 0x12, 0xa0, 0x01, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x35, 0x2e, 0x61, 0x70, 0x70,
	0x75, 0x73, 0x65, 0x72, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x61, 0x70, 0x70,
	0x72, 0x6f, 0x6c, 0x65, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x36, 0x2e, 0x61, 0x70, 0x70, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x67, 0x61, 0x74, 0x65,
	0x77, 0x61, 0x79, 0x2e, 0x61, 0x70, 0x70, 0x72, 0x6f, 0x6c, 0x65, 0x75, 0x73, 0x65, 0x72, 0x2e,
	0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x55, 0x73, 0x65,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1f, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x19, 0x22, 0x14, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x2f, 0x72, 0x6f,
	0x6c, 0x65, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x3a, 0x01, 0x2a, 0x12, 0x98, 0x01, 0x0a, 0x0c, 0x47,
	0x65, 0x74, 0x52, 0x6f, 0x6c, 0x65, 0x55, 0x73, 0x65, 0x72, 0x73, 0x12, 0x33, 0x2e, 0x61, 0x70,
	0x70, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x61, 0x70,
	0x70, 0x72, 0x6f, 0x6c, 0x65, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74,
	0x52, 0x6f, 0x6c, 0x65, 0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x34, 0x2e, 0x61, 0x70, 0x70, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77,
	0x61, 0x79, 0x2e, 0x61, 0x70, 0x70, 0x72, 0x6f, 0x6c, 0x65, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76,
	0x31, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x6f, 0x6c, 0x65, 0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x17, 0x22, 0x12,
	0x2f, 0x76, 0x31, 0x2f, 0x67, 0x65, 0x74, 0x2f, 0x72, 0x6f, 0x6c, 0x65, 0x2f, 0x75, 0x73, 0x65,
	0x72, 0x73, 0x3a, 0x01, 0x2a, 0x12, 0xa0, 0x01, 0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x52, 0x6f, 0x6c, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x35, 0x2e, 0x61, 0x70, 0x70, 0x75, 0x73,
	0x65, 0x72, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x61, 0x70, 0x70, 0x72, 0x6f,
	0x6c, 0x65, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x52, 0x6f, 0x6c, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x36, 0x2e, 0x61, 0x70, 0x70, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61,
	0x79, 0x2e, 0x61, 0x70, 0x70, 0x72, 0x6f, 0x6c, 0x65, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31,
	0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x19, 0x22,
	0x14, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x2f, 0x72, 0x6f, 0x6c, 0x65,
	0x2f, 0x75, 0x73, 0x65, 0x72, 0x3a, 0x01, 0x2a, 0x12, 0xad, 0x01, 0x0a, 0x11, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x41, 0x70, 0x70, 0x52, 0x6f, 0x6c, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x38,
	0x2e, 0x61, 0x70, 0x70, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79,
	0x2e, 0x61, 0x70, 0x70, 0x72, 0x6f, 0x6c, 0x65, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x70, 0x70, 0x52, 0x6f, 0x6c, 0x65, 0x55, 0x73, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x39, 0x2e, 0x61, 0x70, 0x70, 0x75, 0x73,
	0x65, 0x72, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x61, 0x70, 0x70, 0x72, 0x6f,
	0x6c, 0x65, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x41, 0x70, 0x70, 0x52, 0x6f, 0x6c, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x23, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1d, 0x22, 0x18, 0x2f, 0x76, 0x31,
	0x2f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x2f, 0x61, 0x70, 0x70, 0x2f, 0x72, 0x6f, 0x6c, 0x65,
	0x2f, 0x75, 0x73, 0x65, 0x72, 0x3a, 0x01, 0x2a, 0x12, 0xa5, 0x01, 0x0a, 0x0f, 0x47, 0x65, 0x74,
	0x41, 0x70, 0x70, 0x52, 0x6f, 0x6c, 0x65, 0x55, 0x73, 0x65, 0x72, 0x73, 0x12, 0x36, 0x2e, 0x61,
	0x70, 0x70, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x61,
	0x70, 0x70, 0x72, 0x6f, 0x6c, 0x65, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65,
	0x74, 0x41, 0x70, 0x70, 0x52, 0x6f, 0x6c, 0x65, 0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x37, 0x2e, 0x61, 0x70, 0x70, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x67,
	0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x61, 0x70, 0x70, 0x72, 0x6f, 0x6c, 0x65, 0x75, 0x73,
	0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x70, 0x70, 0x52, 0x6f, 0x6c, 0x65,
	0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x21, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x1b, 0x22, 0x16, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x65, 0x74, 0x2f, 0x61,
	0x70, 0x70, 0x2f, 0x72, 0x6f, 0x6c, 0x65, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x3a, 0x01, 0x2a,
	0x42, 0x42, 0x5a, 0x40, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4e,
	0x70, 0x6f, 0x6f, 0x6c, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x2f, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x61, 0x70, 0x70, 0x75, 0x73,
	0x65, 0x72, 0x2f, 0x67, 0x77, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x70, 0x70, 0x72, 0x6f, 0x6c, 0x65,
	0x75, 0x73, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_npool_appuser_gw_v1_approleuser_approleuser_proto_rawDescOnce sync.Once
	file_npool_appuser_gw_v1_approleuser_approleuser_proto_rawDescData = file_npool_appuser_gw_v1_approleuser_approleuser_proto_rawDesc
)

func file_npool_appuser_gw_v1_approleuser_approleuser_proto_rawDescGZIP() []byte {
	file_npool_appuser_gw_v1_approleuser_approleuser_proto_rawDescOnce.Do(func() {
		file_npool_appuser_gw_v1_approleuser_approleuser_proto_rawDescData = protoimpl.X.CompressGZIP(file_npool_appuser_gw_v1_approleuser_approleuser_proto_rawDescData)
	})
	return file_npool_appuser_gw_v1_approleuser_approleuser_proto_rawDescData
}

var file_npool_appuser_gw_v1_approleuser_approleuser_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_npool_appuser_gw_v1_approleuser_approleuser_proto_goTypes = []interface{}{
	(*CreateRoleUserRequest)(nil),      // 0: appuser.gateway.approleuser.v1.CreateRoleUserRequest
	(*CreateRoleUserResponse)(nil),     // 1: appuser.gateway.approleuser.v1.CreateRoleUserResponse
	(*GetRoleUsersRequest)(nil),        // 2: appuser.gateway.approleuser.v1.GetRoleUsersRequest
	(*GetRoleUsersResponse)(nil),       // 3: appuser.gateway.approleuser.v1.GetRoleUsersResponse
	(*DeleteRoleUserRequest)(nil),      // 4: appuser.gateway.approleuser.v1.DeleteRoleUserRequest
	(*DeleteRoleUserResponse)(nil),     // 5: appuser.gateway.approleuser.v1.DeleteRoleUserResponse
	(*CreateAppRoleUserRequest)(nil),   // 6: appuser.gateway.approleuser.v1.CreateAppRoleUserRequest
	(*CreateAppRoleUserResponse)(nil),  // 7: appuser.gateway.approleuser.v1.CreateAppRoleUserResponse
	(*GetAppRoleUsersRequest)(nil),     // 8: appuser.gateway.approleuser.v1.GetAppRoleUsersRequest
	(*GetAppRoleUsersResponse)(nil),    // 9: appuser.gateway.approleuser.v1.GetAppRoleUsersResponse
	(*approleuser.AppRoleUserReq)(nil), // 10: appuser.manager.approleuser.v2.AppRoleUserReq
	(*approleuser.AppRoleUser)(nil),    // 11: appuser.manager.approleuser.v2.AppRoleUser
}
var file_npool_appuser_gw_v1_approleuser_approleuser_proto_depIdxs = []int32{
	10, // 0: appuser.gateway.approleuser.v1.CreateRoleUserRequest.Info:type_name -> appuser.manager.approleuser.v2.AppRoleUserReq
	11, // 1: appuser.gateway.approleuser.v1.CreateRoleUserResponse.Info:type_name -> appuser.manager.approleuser.v2.AppRoleUser
	11, // 2: appuser.gateway.approleuser.v1.GetRoleUsersResponse.Infos:type_name -> appuser.manager.approleuser.v2.AppRoleUser
	11, // 3: appuser.gateway.approleuser.v1.DeleteRoleUserResponse.Info:type_name -> appuser.manager.approleuser.v2.AppRoleUser
	10, // 4: appuser.gateway.approleuser.v1.CreateAppRoleUserRequest.Info:type_name -> appuser.manager.approleuser.v2.AppRoleUserReq
	11, // 5: appuser.gateway.approleuser.v1.CreateAppRoleUserResponse.Info:type_name -> appuser.manager.approleuser.v2.AppRoleUser
	11, // 6: appuser.gateway.approleuser.v1.GetAppRoleUsersResponse.Infos:type_name -> appuser.manager.approleuser.v2.AppRoleUser
	0,  // 7: appuser.gateway.approleuser.v1.AppRoleUserGw.CreateRoleUser:input_type -> appuser.gateway.approleuser.v1.CreateRoleUserRequest
	2,  // 8: appuser.gateway.approleuser.v1.AppRoleUserGw.GetRoleUsers:input_type -> appuser.gateway.approleuser.v1.GetRoleUsersRequest
	4,  // 9: appuser.gateway.approleuser.v1.AppRoleUserGw.DeleteRoleUser:input_type -> appuser.gateway.approleuser.v1.DeleteRoleUserRequest
	6,  // 10: appuser.gateway.approleuser.v1.AppRoleUserGw.CreateAppRoleUser:input_type -> appuser.gateway.approleuser.v1.CreateAppRoleUserRequest
	8,  // 11: appuser.gateway.approleuser.v1.AppRoleUserGw.GetAppRoleUsers:input_type -> appuser.gateway.approleuser.v1.GetAppRoleUsersRequest
	1,  // 12: appuser.gateway.approleuser.v1.AppRoleUserGw.CreateRoleUser:output_type -> appuser.gateway.approleuser.v1.CreateRoleUserResponse
	3,  // 13: appuser.gateway.approleuser.v1.AppRoleUserGw.GetRoleUsers:output_type -> appuser.gateway.approleuser.v1.GetRoleUsersResponse
	5,  // 14: appuser.gateway.approleuser.v1.AppRoleUserGw.DeleteRoleUser:output_type -> appuser.gateway.approleuser.v1.DeleteRoleUserResponse
	7,  // 15: appuser.gateway.approleuser.v1.AppRoleUserGw.CreateAppRoleUser:output_type -> appuser.gateway.approleuser.v1.CreateAppRoleUserResponse
	9,  // 16: appuser.gateway.approleuser.v1.AppRoleUserGw.GetAppRoleUsers:output_type -> appuser.gateway.approleuser.v1.GetAppRoleUsersResponse
	12, // [12:17] is the sub-list for method output_type
	7,  // [7:12] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_npool_appuser_gw_v1_approleuser_approleuser_proto_init() }
func file_npool_appuser_gw_v1_approleuser_approleuser_proto_init() {
	if File_npool_appuser_gw_v1_approleuser_approleuser_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_npool_appuser_gw_v1_approleuser_approleuser_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateRoleUserRequest); i {
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
		file_npool_appuser_gw_v1_approleuser_approleuser_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateRoleUserResponse); i {
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
		file_npool_appuser_gw_v1_approleuser_approleuser_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRoleUsersRequest); i {
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
		file_npool_appuser_gw_v1_approleuser_approleuser_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRoleUsersResponse); i {
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
		file_npool_appuser_gw_v1_approleuser_approleuser_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteRoleUserRequest); i {
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
		file_npool_appuser_gw_v1_approleuser_approleuser_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteRoleUserResponse); i {
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
		file_npool_appuser_gw_v1_approleuser_approleuser_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateAppRoleUserRequest); i {
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
		file_npool_appuser_gw_v1_approleuser_approleuser_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateAppRoleUserResponse); i {
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
		file_npool_appuser_gw_v1_approleuser_approleuser_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAppRoleUsersRequest); i {
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
		file_npool_appuser_gw_v1_approleuser_approleuser_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAppRoleUsersResponse); i {
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
			RawDescriptor: file_npool_appuser_gw_v1_approleuser_approleuser_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_npool_appuser_gw_v1_approleuser_approleuser_proto_goTypes,
		DependencyIndexes: file_npool_appuser_gw_v1_approleuser_approleuser_proto_depIdxs,
		MessageInfos:      file_npool_appuser_gw_v1_approleuser_approleuser_proto_msgTypes,
	}.Build()
	File_npool_appuser_gw_v1_approleuser_approleuser_proto = out.File
	file_npool_appuser_gw_v1_approleuser_approleuser_proto_rawDesc = nil
	file_npool_appuser_gw_v1_approleuser_approleuser_proto_goTypes = nil
	file_npool_appuser_gw_v1_approleuser_approleuser_proto_depIdxs = nil
}
