// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.18.1
// source: npool/appuser/gw/v1/banappuser/banappuser.proto

package banappuser

import (
	banappuser "github.com/NpoolPlatform/message/npool/appuser/mgr/v2/banappuser"
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

type CreateBanAppUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *banappuser.BanAppUserReq `protobuf:"bytes,10,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (x *CreateBanAppUserRequest) Reset() {
	*x = CreateBanAppUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_appuser_gw_v1_banappuser_banappuser_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateBanAppUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateBanAppUserRequest) ProtoMessage() {}

func (x *CreateBanAppUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_appuser_gw_v1_banappuser_banappuser_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateBanAppUserRequest.ProtoReflect.Descriptor instead.
func (*CreateBanAppUserRequest) Descriptor() ([]byte, []int) {
	return file_npool_appuser_gw_v1_banappuser_banappuser_proto_rawDescGZIP(), []int{0}
}

func (x *CreateBanAppUserRequest) GetInfo() *banappuser.BanAppUserReq {
	if x != nil {
		return x.Info
	}
	return nil
}

type CreateBanAppUserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *banappuser.BanAppUser `protobuf:"bytes,10,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (x *CreateBanAppUserResponse) Reset() {
	*x = CreateBanAppUserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_appuser_gw_v1_banappuser_banappuser_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateBanAppUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateBanAppUserResponse) ProtoMessage() {}

func (x *CreateBanAppUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_appuser_gw_v1_banappuser_banappuser_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateBanAppUserResponse.ProtoReflect.Descriptor instead.
func (*CreateBanAppUserResponse) Descriptor() ([]byte, []int) {
	return file_npool_appuser_gw_v1_banappuser_banappuser_proto_rawDescGZIP(), []int{1}
}

func (x *CreateBanAppUserResponse) GetInfo() *banappuser.BanAppUser {
	if x != nil {
		return x.Info
	}
	return nil
}

type GetBanAppUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID string `protobuf:"bytes,10,opt,name=ID,proto3" json:"ID,omitempty"`
}

func (x *GetBanAppUserRequest) Reset() {
	*x = GetBanAppUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_appuser_gw_v1_banappuser_banappuser_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBanAppUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBanAppUserRequest) ProtoMessage() {}

func (x *GetBanAppUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_appuser_gw_v1_banappuser_banappuser_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBanAppUserRequest.ProtoReflect.Descriptor instead.
func (*GetBanAppUserRequest) Descriptor() ([]byte, []int) {
	return file_npool_appuser_gw_v1_banappuser_banappuser_proto_rawDescGZIP(), []int{2}
}

func (x *GetBanAppUserRequest) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

type GetBanAppUserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *banappuser.BanAppUser `protobuf:"bytes,10,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (x *GetBanAppUserResponse) Reset() {
	*x = GetBanAppUserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_appuser_gw_v1_banappuser_banappuser_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBanAppUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBanAppUserResponse) ProtoMessage() {}

func (x *GetBanAppUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_appuser_gw_v1_banappuser_banappuser_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBanAppUserResponse.ProtoReflect.Descriptor instead.
func (*GetBanAppUserResponse) Descriptor() ([]byte, []int) {
	return file_npool_appuser_gw_v1_banappuser_banappuser_proto_rawDescGZIP(), []int{3}
}

func (x *GetBanAppUserResponse) GetInfo() *banappuser.BanAppUser {
	if x != nil {
		return x.Info
	}
	return nil
}

type GetAppUserBanAppUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TargetAppID string `protobuf:"bytes,10,opt,name=TargetAppID,proto3" json:"TargetAppID,omitempty"`
	UserID      string `protobuf:"bytes,20,opt,name=UserID,proto3" json:"UserID,omitempty"`
}

func (x *GetAppUserBanAppUserRequest) Reset() {
	*x = GetAppUserBanAppUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_appuser_gw_v1_banappuser_banappuser_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAppUserBanAppUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAppUserBanAppUserRequest) ProtoMessage() {}

func (x *GetAppUserBanAppUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_appuser_gw_v1_banappuser_banappuser_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAppUserBanAppUserRequest.ProtoReflect.Descriptor instead.
func (*GetAppUserBanAppUserRequest) Descriptor() ([]byte, []int) {
	return file_npool_appuser_gw_v1_banappuser_banappuser_proto_rawDescGZIP(), []int{4}
}

func (x *GetAppUserBanAppUserRequest) GetTargetAppID() string {
	if x != nil {
		return x.TargetAppID
	}
	return ""
}

func (x *GetAppUserBanAppUserRequest) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

type GetAppUserBanAppUserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *banappuser.BanAppUser `protobuf:"bytes,10,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (x *GetAppUserBanAppUserResponse) Reset() {
	*x = GetAppUserBanAppUserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_appuser_gw_v1_banappuser_banappuser_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAppUserBanAppUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAppUserBanAppUserResponse) ProtoMessage() {}

func (x *GetAppUserBanAppUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_appuser_gw_v1_banappuser_banappuser_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAppUserBanAppUserResponse.ProtoReflect.Descriptor instead.
func (*GetAppUserBanAppUserResponse) Descriptor() ([]byte, []int) {
	return file_npool_appuser_gw_v1_banappuser_banappuser_proto_rawDescGZIP(), []int{5}
}

func (x *GetAppUserBanAppUserResponse) GetInfo() *banappuser.BanAppUser {
	if x != nil {
		return x.Info
	}
	return nil
}

type UpdateBanAppUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *banappuser.BanAppUserReq `protobuf:"bytes,10,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (x *UpdateBanAppUserRequest) Reset() {
	*x = UpdateBanAppUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_appuser_gw_v1_banappuser_banappuser_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateBanAppUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateBanAppUserRequest) ProtoMessage() {}

func (x *UpdateBanAppUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_appuser_gw_v1_banappuser_banappuser_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateBanAppUserRequest.ProtoReflect.Descriptor instead.
func (*UpdateBanAppUserRequest) Descriptor() ([]byte, []int) {
	return file_npool_appuser_gw_v1_banappuser_banappuser_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateBanAppUserRequest) GetInfo() *banappuser.BanAppUserReq {
	if x != nil {
		return x.Info
	}
	return nil
}

type UpdateBanAppUserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *banappuser.BanAppUser `protobuf:"bytes,10,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (x *UpdateBanAppUserResponse) Reset() {
	*x = UpdateBanAppUserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_appuser_gw_v1_banappuser_banappuser_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateBanAppUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateBanAppUserResponse) ProtoMessage() {}

func (x *UpdateBanAppUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_appuser_gw_v1_banappuser_banappuser_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateBanAppUserResponse.ProtoReflect.Descriptor instead.
func (*UpdateBanAppUserResponse) Descriptor() ([]byte, []int) {
	return file_npool_appuser_gw_v1_banappuser_banappuser_proto_rawDescGZIP(), []int{7}
}

func (x *UpdateBanAppUserResponse) GetInfo() *banappuser.BanAppUser {
	if x != nil {
		return x.Info
	}
	return nil
}

type DeleteBanAppUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID string `protobuf:"bytes,10,opt,name=ID,proto3" json:"ID,omitempty"`
}

func (x *DeleteBanAppUserRequest) Reset() {
	*x = DeleteBanAppUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_appuser_gw_v1_banappuser_banappuser_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteBanAppUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteBanAppUserRequest) ProtoMessage() {}

func (x *DeleteBanAppUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_appuser_gw_v1_banappuser_banappuser_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteBanAppUserRequest.ProtoReflect.Descriptor instead.
func (*DeleteBanAppUserRequest) Descriptor() ([]byte, []int) {
	return file_npool_appuser_gw_v1_banappuser_banappuser_proto_rawDescGZIP(), []int{8}
}

func (x *DeleteBanAppUserRequest) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

type DeleteBanAppUserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *banappuser.BanAppUser `protobuf:"bytes,10,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (x *DeleteBanAppUserResponse) Reset() {
	*x = DeleteBanAppUserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_appuser_gw_v1_banappuser_banappuser_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteBanAppUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteBanAppUserResponse) ProtoMessage() {}

func (x *DeleteBanAppUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_appuser_gw_v1_banappuser_banappuser_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteBanAppUserResponse.ProtoReflect.Descriptor instead.
func (*DeleteBanAppUserResponse) Descriptor() ([]byte, []int) {
	return file_npool_appuser_gw_v1_banappuser_banappuser_proto_rawDescGZIP(), []int{9}
}

func (x *DeleteBanAppUserResponse) GetInfo() *banappuser.BanAppUser {
	if x != nil {
		return x.Info
	}
	return nil
}

var File_npool_appuser_gw_v1_banappuser_banappuser_proto protoreflect.FileDescriptor

var file_npool_appuser_gw_v1_banappuser_banappuser_proto_rawDesc = []byte{
	0x0a, 0x2f, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x61, 0x70, 0x70, 0x75, 0x73, 0x65, 0x72, 0x2f,
	0x67, 0x77, 0x2f, 0x76, 0x31, 0x2f, 0x62, 0x61, 0x6e, 0x61, 0x70, 0x70, 0x75, 0x73, 0x65, 0x72,
	0x2f, 0x62, 0x61, 0x6e, 0x61, 0x70, 0x70, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x1d, 0x61, 0x70, 0x70, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77,
	0x61, 0x79, 0x2e, 0x62, 0x61, 0x6e, 0x61, 0x70, 0x70, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31,
	0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x30,
	0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x61, 0x70, 0x70, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x6d, 0x67,
	0x72, 0x2f, 0x76, 0x32, 0x2f, 0x62, 0x61, 0x6e, 0x61, 0x70, 0x70, 0x75, 0x73, 0x65, 0x72, 0x2f,
	0x62, 0x61, 0x6e, 0x61, 0x70, 0x70, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x5b, 0x0a, 0x17, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x61, 0x6e, 0x41, 0x70, 0x70,
	0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x40, 0x0a, 0x04, 0x49,
	0x6e, 0x66, 0x6f, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x61, 0x70, 0x70, 0x75,
	0x73, 0x65, 0x72, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x62, 0x61, 0x6e, 0x61,
	0x70, 0x70, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x32, 0x2e, 0x42, 0x61, 0x6e, 0x41, 0x70, 0x70,
	0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x52, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x59, 0x0a,
	0x18, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x61, 0x6e, 0x41, 0x70, 0x70, 0x55, 0x73, 0x65,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3d, 0x0a, 0x04, 0x49, 0x6e, 0x66,
	0x6f, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x61, 0x70, 0x70, 0x75, 0x73, 0x65,
	0x72, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x62, 0x61, 0x6e, 0x61, 0x70, 0x70,
	0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x32, 0x2e, 0x42, 0x61, 0x6e, 0x41, 0x70, 0x70, 0x55, 0x73,
	0x65, 0x72, 0x52, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x26, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x42,
	0x61, 0x6e, 0x41, 0x70, 0x70, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44,
	0x22, 0x56, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x42, 0x61, 0x6e, 0x41, 0x70, 0x70, 0x55, 0x73, 0x65,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3d, 0x0a, 0x04, 0x49, 0x6e, 0x66,
	0x6f, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x61, 0x70, 0x70, 0x75, 0x73, 0x65,
	0x72, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x62, 0x61, 0x6e, 0x61, 0x70, 0x70,
	0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x32, 0x2e, 0x42, 0x61, 0x6e, 0x41, 0x70, 0x70, 0x55, 0x73,
	0x65, 0x72, 0x52, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x57, 0x0a, 0x1b, 0x47, 0x65, 0x74, 0x41,
	0x70, 0x70, 0x55, 0x73, 0x65, 0x72, 0x42, 0x61, 0x6e, 0x41, 0x70, 0x70, 0x55, 0x73, 0x65, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x54, 0x61, 0x72, 0x67, 0x65,
	0x74, 0x41, 0x70, 0x70, 0x49, 0x44, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x54, 0x61,
	0x72, 0x67, 0x65, 0x74, 0x41, 0x70, 0x70, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65,
	0x72, 0x49, 0x44, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49,
	0x44, 0x22, 0x5d, 0x0a, 0x1c, 0x47, 0x65, 0x74, 0x41, 0x70, 0x70, 0x55, 0x73, 0x65, 0x72, 0x42,
	0x61, 0x6e, 0x41, 0x70, 0x70, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x3d, 0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x29, 0x2e, 0x61, 0x70, 0x70, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65,
	0x72, 0x2e, 0x62, 0x61, 0x6e, 0x61, 0x70, 0x70, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x32, 0x2e,
	0x42, 0x61, 0x6e, 0x41, 0x70, 0x70, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x49, 0x6e, 0x66, 0x6f,
	0x22, 0x5b, 0x0a, 0x17, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x61, 0x6e, 0x41, 0x70, 0x70,
	0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x40, 0x0a, 0x04, 0x49,
	0x6e, 0x66, 0x6f, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x61, 0x70, 0x70, 0x75,
	0x73, 0x65, 0x72, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x62, 0x61, 0x6e, 0x61,
	0x70, 0x70, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x32, 0x2e, 0x42, 0x61, 0x6e, 0x41, 0x70, 0x70,
	0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x52, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x59, 0x0a,
	0x18, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x61, 0x6e, 0x41, 0x70, 0x70, 0x55, 0x73, 0x65,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3d, 0x0a, 0x04, 0x49, 0x6e, 0x66,
	0x6f, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x61, 0x70, 0x70, 0x75, 0x73, 0x65,
	0x72, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x62, 0x61, 0x6e, 0x61, 0x70, 0x70,
	0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x32, 0x2e, 0x42, 0x61, 0x6e, 0x41, 0x70, 0x70, 0x55, 0x73,
	0x65, 0x72, 0x52, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x29, 0x0a, 0x17, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x42, 0x61, 0x6e, 0x41, 0x70, 0x70, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x49, 0x44, 0x22, 0x59, 0x0a, 0x18, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x61, 0x6e,
	0x41, 0x70, 0x70, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x3d, 0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e,
	0x61, 0x70, 0x70, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e,
	0x62, 0x61, 0x6e, 0x61, 0x70, 0x70, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x32, 0x2e, 0x42, 0x61,
	0x6e, 0x41, 0x70, 0x70, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x32, 0xe6,
	0x06, 0x0a, 0x0c, 0x42, 0x61, 0x6e, 0x41, 0x70, 0x70, 0x55, 0x73, 0x65, 0x72, 0x47, 0x77, 0x12,
	0xa7, 0x01, 0x0a, 0x10, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x61, 0x6e, 0x41, 0x70, 0x70,
	0x55, 0x73, 0x65, 0x72, 0x12, 0x36, 0x2e, 0x61, 0x70, 0x70, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x67,
	0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x62, 0x61, 0x6e, 0x61, 0x70, 0x70, 0x75, 0x73, 0x65,
	0x72, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x61, 0x6e, 0x41, 0x70,
	0x70, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x37, 0x2e, 0x61,
	0x70, 0x70, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x62,
	0x61, 0x6e, 0x61, 0x70, 0x70, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x42, 0x61, 0x6e, 0x41, 0x70, 0x70, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x22, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1c, 0x22, 0x17, 0x2f,
	0x76, 0x32, 0x2f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x2f, 0x62, 0x61, 0x6e, 0x2f, 0x61, 0x70,
	0x70, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x3a, 0x01, 0x2a, 0x12, 0x9b, 0x01, 0x0a, 0x0d, 0x47, 0x65,
	0x74, 0x42, 0x61, 0x6e, 0x41, 0x70, 0x70, 0x55, 0x73, 0x65, 0x72, 0x12, 0x33, 0x2e, 0x61, 0x70,
	0x70, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x62, 0x61,
	0x6e, 0x61, 0x70, 0x70, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x42,
	0x61, 0x6e, 0x41, 0x70, 0x70, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x34, 0x2e, 0x61, 0x70, 0x70, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77,
	0x61, 0x79, 0x2e, 0x62, 0x61, 0x6e, 0x61, 0x70, 0x70, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31,
	0x2e, 0x47, 0x65, 0x74, 0x42, 0x61, 0x6e, 0x41, 0x70, 0x70, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x19, 0x22, 0x14,
	0x2f, 0x76, 0x32, 0x2f, 0x67, 0x65, 0x74, 0x2f, 0x62, 0x61, 0x6e, 0x2f, 0x61, 0x70, 0x70, 0x2f,
	0x75, 0x73, 0x65, 0x72, 0x3a, 0x01, 0x2a, 0x12, 0xb9, 0x01, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x41,
	0x70, 0x70, 0x55, 0x73, 0x65, 0x72, 0x42, 0x61, 0x6e, 0x41, 0x70, 0x70, 0x55, 0x73, 0x65, 0x72,
	0x12, 0x3a, 0x2e, 0x61, 0x70, 0x70, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77,
	0x61, 0x79, 0x2e, 0x62, 0x61, 0x6e, 0x61, 0x70, 0x70, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31,
	0x2e, 0x47, 0x65, 0x74, 0x41, 0x70, 0x70, 0x55, 0x73, 0x65, 0x72, 0x42, 0x61, 0x6e, 0x41, 0x70,
	0x70, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3b, 0x2e, 0x61,
	0x70, 0x70, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x62,
	0x61, 0x6e, 0x61, 0x70, 0x70, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74,
	0x41, 0x70, 0x70, 0x55, 0x73, 0x65, 0x72, 0x42, 0x61, 0x6e, 0x41, 0x70, 0x70, 0x55, 0x73, 0x65,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x28, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x22, 0x22, 0x1d, 0x2f, 0x76, 0x32, 0x2f, 0x67, 0x65, 0x74, 0x2f, 0x61, 0x70, 0x70, 0x2f, 0x75,
	0x73, 0x65, 0x72, 0x2f, 0x62, 0x61, 0x6e, 0x2f, 0x61, 0x70, 0x70, 0x2f, 0x75, 0x73, 0x65, 0x72,
	0x3a, 0x01, 0x2a, 0x12, 0xa7, 0x01, 0x0a, 0x10, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x61,
	0x6e, 0x41, 0x70, 0x70, 0x55, 0x73, 0x65, 0x72, 0x12, 0x36, 0x2e, 0x61, 0x70, 0x70, 0x75, 0x73,
	0x65, 0x72, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x62, 0x61, 0x6e, 0x61, 0x70,
	0x70, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42,
	0x61, 0x6e, 0x41, 0x70, 0x70, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x37, 0x2e, 0x61, 0x70, 0x70, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77,
	0x61, 0x79, 0x2e, 0x62, 0x61, 0x6e, 0x61, 0x70, 0x70, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31,
	0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x61, 0x6e, 0x41, 0x70, 0x70, 0x55, 0x73, 0x65,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x22, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x1c, 0x22, 0x17, 0x2f, 0x76, 0x32, 0x2f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x62, 0x61,
	0x6e, 0x2f, 0x61, 0x70, 0x70, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x3a, 0x01, 0x2a, 0x12, 0xa7, 0x01,
	0x0a, 0x10, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x61, 0x6e, 0x41, 0x70, 0x70, 0x55, 0x73,
	0x65, 0x72, 0x12, 0x36, 0x2e, 0x61, 0x70, 0x70, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x67, 0x61, 0x74,
	0x65, 0x77, 0x61, 0x79, 0x2e, 0x62, 0x61, 0x6e, 0x61, 0x70, 0x70, 0x75, 0x73, 0x65, 0x72, 0x2e,
	0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x61, 0x6e, 0x41, 0x70, 0x70, 0x55,
	0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x37, 0x2e, 0x61, 0x70, 0x70,
	0x75, 0x73, 0x65, 0x72, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x62, 0x61, 0x6e,
	0x61, 0x70, 0x70, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x42, 0x61, 0x6e, 0x41, 0x70, 0x70, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x22, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1c, 0x22, 0x17, 0x2f, 0x76, 0x32,
	0x2f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x2f, 0x62, 0x61, 0x6e, 0x2f, 0x61, 0x70, 0x70, 0x2f,
	0x75, 0x73, 0x65, 0x72, 0x3a, 0x01, 0x2a, 0x42, 0x41, 0x5a, 0x3f, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4e, 0x70, 0x6f, 0x6f, 0x6c, 0x50, 0x6c, 0x61, 0x74, 0x66,
	0x6f, 0x72, 0x6d, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x6e, 0x70, 0x6f, 0x6f,
	0x6c, 0x2f, 0x61, 0x70, 0x70, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x67, 0x77, 0x2f, 0x76, 0x31, 0x2f,
	0x62, 0x61, 0x6e, 0x61, 0x70, 0x70, 0x75, 0x73, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_npool_appuser_gw_v1_banappuser_banappuser_proto_rawDescOnce sync.Once
	file_npool_appuser_gw_v1_banappuser_banappuser_proto_rawDescData = file_npool_appuser_gw_v1_banappuser_banappuser_proto_rawDesc
)

func file_npool_appuser_gw_v1_banappuser_banappuser_proto_rawDescGZIP() []byte {
	file_npool_appuser_gw_v1_banappuser_banappuser_proto_rawDescOnce.Do(func() {
		file_npool_appuser_gw_v1_banappuser_banappuser_proto_rawDescData = protoimpl.X.CompressGZIP(file_npool_appuser_gw_v1_banappuser_banappuser_proto_rawDescData)
	})
	return file_npool_appuser_gw_v1_banappuser_banappuser_proto_rawDescData
}

var file_npool_appuser_gw_v1_banappuser_banappuser_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_npool_appuser_gw_v1_banappuser_banappuser_proto_goTypes = []interface{}{
	(*CreateBanAppUserRequest)(nil),      // 0: appuser.gateway.banappuser.v1.CreateBanAppUserRequest
	(*CreateBanAppUserResponse)(nil),     // 1: appuser.gateway.banappuser.v1.CreateBanAppUserResponse
	(*GetBanAppUserRequest)(nil),         // 2: appuser.gateway.banappuser.v1.GetBanAppUserRequest
	(*GetBanAppUserResponse)(nil),        // 3: appuser.gateway.banappuser.v1.GetBanAppUserResponse
	(*GetAppUserBanAppUserRequest)(nil),  // 4: appuser.gateway.banappuser.v1.GetAppUserBanAppUserRequest
	(*GetAppUserBanAppUserResponse)(nil), // 5: appuser.gateway.banappuser.v1.GetAppUserBanAppUserResponse
	(*UpdateBanAppUserRequest)(nil),      // 6: appuser.gateway.banappuser.v1.UpdateBanAppUserRequest
	(*UpdateBanAppUserResponse)(nil),     // 7: appuser.gateway.banappuser.v1.UpdateBanAppUserResponse
	(*DeleteBanAppUserRequest)(nil),      // 8: appuser.gateway.banappuser.v1.DeleteBanAppUserRequest
	(*DeleteBanAppUserResponse)(nil),     // 9: appuser.gateway.banappuser.v1.DeleteBanAppUserResponse
	(*banappuser.BanAppUserReq)(nil),     // 10: appuser.manager.banappuser.v2.BanAppUserReq
	(*banappuser.BanAppUser)(nil),        // 11: appuser.manager.banappuser.v2.BanAppUser
}
var file_npool_appuser_gw_v1_banappuser_banappuser_proto_depIdxs = []int32{
	10, // 0: appuser.gateway.banappuser.v1.CreateBanAppUserRequest.Info:type_name -> appuser.manager.banappuser.v2.BanAppUserReq
	11, // 1: appuser.gateway.banappuser.v1.CreateBanAppUserResponse.Info:type_name -> appuser.manager.banappuser.v2.BanAppUser
	11, // 2: appuser.gateway.banappuser.v1.GetBanAppUserResponse.Info:type_name -> appuser.manager.banappuser.v2.BanAppUser
	11, // 3: appuser.gateway.banappuser.v1.GetAppUserBanAppUserResponse.Info:type_name -> appuser.manager.banappuser.v2.BanAppUser
	10, // 4: appuser.gateway.banappuser.v1.UpdateBanAppUserRequest.Info:type_name -> appuser.manager.banappuser.v2.BanAppUserReq
	11, // 5: appuser.gateway.banappuser.v1.UpdateBanAppUserResponse.Info:type_name -> appuser.manager.banappuser.v2.BanAppUser
	11, // 6: appuser.gateway.banappuser.v1.DeleteBanAppUserResponse.Info:type_name -> appuser.manager.banappuser.v2.BanAppUser
	0,  // 7: appuser.gateway.banappuser.v1.BanAppUserGw.CreateBanAppUser:input_type -> appuser.gateway.banappuser.v1.CreateBanAppUserRequest
	2,  // 8: appuser.gateway.banappuser.v1.BanAppUserGw.GetBanAppUser:input_type -> appuser.gateway.banappuser.v1.GetBanAppUserRequest
	4,  // 9: appuser.gateway.banappuser.v1.BanAppUserGw.GetAppUserBanAppUser:input_type -> appuser.gateway.banappuser.v1.GetAppUserBanAppUserRequest
	6,  // 10: appuser.gateway.banappuser.v1.BanAppUserGw.UpdateBanAppUser:input_type -> appuser.gateway.banappuser.v1.UpdateBanAppUserRequest
	8,  // 11: appuser.gateway.banappuser.v1.BanAppUserGw.DeleteBanAppUser:input_type -> appuser.gateway.banappuser.v1.DeleteBanAppUserRequest
	1,  // 12: appuser.gateway.banappuser.v1.BanAppUserGw.CreateBanAppUser:output_type -> appuser.gateway.banappuser.v1.CreateBanAppUserResponse
	3,  // 13: appuser.gateway.banappuser.v1.BanAppUserGw.GetBanAppUser:output_type -> appuser.gateway.banappuser.v1.GetBanAppUserResponse
	5,  // 14: appuser.gateway.banappuser.v1.BanAppUserGw.GetAppUserBanAppUser:output_type -> appuser.gateway.banappuser.v1.GetAppUserBanAppUserResponse
	7,  // 15: appuser.gateway.banappuser.v1.BanAppUserGw.UpdateBanAppUser:output_type -> appuser.gateway.banappuser.v1.UpdateBanAppUserResponse
	9,  // 16: appuser.gateway.banappuser.v1.BanAppUserGw.DeleteBanAppUser:output_type -> appuser.gateway.banappuser.v1.DeleteBanAppUserResponse
	12, // [12:17] is the sub-list for method output_type
	7,  // [7:12] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_npool_appuser_gw_v1_banappuser_banappuser_proto_init() }
func file_npool_appuser_gw_v1_banappuser_banappuser_proto_init() {
	if File_npool_appuser_gw_v1_banappuser_banappuser_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_npool_appuser_gw_v1_banappuser_banappuser_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateBanAppUserRequest); i {
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
		file_npool_appuser_gw_v1_banappuser_banappuser_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateBanAppUserResponse); i {
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
		file_npool_appuser_gw_v1_banappuser_banappuser_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBanAppUserRequest); i {
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
		file_npool_appuser_gw_v1_banappuser_banappuser_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBanAppUserResponse); i {
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
		file_npool_appuser_gw_v1_banappuser_banappuser_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAppUserBanAppUserRequest); i {
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
		file_npool_appuser_gw_v1_banappuser_banappuser_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAppUserBanAppUserResponse); i {
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
		file_npool_appuser_gw_v1_banappuser_banappuser_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateBanAppUserRequest); i {
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
		file_npool_appuser_gw_v1_banappuser_banappuser_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateBanAppUserResponse); i {
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
		file_npool_appuser_gw_v1_banappuser_banappuser_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteBanAppUserRequest); i {
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
		file_npool_appuser_gw_v1_banappuser_banappuser_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteBanAppUserResponse); i {
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
			RawDescriptor: file_npool_appuser_gw_v1_banappuser_banappuser_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_npool_appuser_gw_v1_banappuser_banappuser_proto_goTypes,
		DependencyIndexes: file_npool_appuser_gw_v1_banappuser_banappuser_proto_depIdxs,
		MessageInfos:      file_npool_appuser_gw_v1_banappuser_banappuser_proto_msgTypes,
	}.Build()
	File_npool_appuser_gw_v1_banappuser_banappuser_proto = out.File
	file_npool_appuser_gw_v1_banappuser_banappuser_proto_rawDesc = nil
	file_npool_appuser_gw_v1_banappuser_banappuser_proto_goTypes = nil
	file_npool_appuser_gw_v1_banappuser_banappuser_proto_depIdxs = nil
}
