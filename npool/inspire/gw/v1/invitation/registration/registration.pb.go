// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.18.1
// source: npool/inspire/gw/v1/invitation/registration/registration.proto

package registration

import (
	_ "github.com/NpoolPlatform/message/npool"
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

type Registration struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID                  string `protobuf:"bytes,10,opt,name=ID,proto3" json:"ID,omitempty"`
	AppID               string `protobuf:"bytes,20,opt,name=AppID,proto3" json:"AppID,omitempty"`
	InviterID           string `protobuf:"bytes,30,opt,name=InviterID,proto3" json:"InviterID,omitempty"`
	InviterEmailAddress string `protobuf:"bytes,40,opt,name=InviterEmailAddress,proto3" json:"InviterEmailAddress,omitempty"`
	InviterPhoneNO      string `protobuf:"bytes,50,opt,name=InviterPhoneNO,proto3" json:"InviterPhoneNO,omitempty"`
	InviterUsername     string `protobuf:"bytes,60,opt,name=InviterUsername,proto3" json:"InviterUsername,omitempty"`
	InviteeID           string `protobuf:"bytes,70,opt,name=InviteeID,proto3" json:"InviteeID,omitempty"`
	InviteeEmailAddress string `protobuf:"bytes,80,opt,name=InviteeEmailAddress,proto3" json:"InviteeEmailAddress,omitempty"`
	InviteePhoneNO      string `protobuf:"bytes,90,opt,name=InviteePhoneNO,proto3" json:"InviteePhoneNO,omitempty"`
	InviteeUsername     string `protobuf:"bytes,100,opt,name=InviteeUsername,proto3" json:"InviteeUsername,omitempty"`
	CreatedAt           uint32 `protobuf:"varint,110,opt,name=CreatedAt,proto3" json:"CreatedAt,omitempty"`
	UpdatedAt           uint32 `protobuf:"varint,120,opt,name=UpdatedAt,proto3" json:"UpdatedAt,omitempty"`
}

func (x *Registration) Reset() {
	*x = Registration{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_inspire_gw_v1_invitation_registration_registration_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Registration) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Registration) ProtoMessage() {}

func (x *Registration) ProtoReflect() protoreflect.Message {
	mi := &file_npool_inspire_gw_v1_invitation_registration_registration_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Registration.ProtoReflect.Descriptor instead.
func (*Registration) Descriptor() ([]byte, []int) {
	return file_npool_inspire_gw_v1_invitation_registration_registration_proto_rawDescGZIP(), []int{0}
}

func (x *Registration) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *Registration) GetAppID() string {
	if x != nil {
		return x.AppID
	}
	return ""
}

func (x *Registration) GetInviterID() string {
	if x != nil {
		return x.InviterID
	}
	return ""
}

func (x *Registration) GetInviterEmailAddress() string {
	if x != nil {
		return x.InviterEmailAddress
	}
	return ""
}

func (x *Registration) GetInviterPhoneNO() string {
	if x != nil {
		return x.InviterPhoneNO
	}
	return ""
}

func (x *Registration) GetInviterUsername() string {
	if x != nil {
		return x.InviterUsername
	}
	return ""
}

func (x *Registration) GetInviteeID() string {
	if x != nil {
		return x.InviteeID
	}
	return ""
}

func (x *Registration) GetInviteeEmailAddress() string {
	if x != nil {
		return x.InviteeEmailAddress
	}
	return ""
}

func (x *Registration) GetInviteePhoneNO() string {
	if x != nil {
		return x.InviteePhoneNO
	}
	return ""
}

func (x *Registration) GetInviteeUsername() string {
	if x != nil {
		return x.InviteeUsername
	}
	return ""
}

func (x *Registration) GetCreatedAt() uint32 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *Registration) GetUpdatedAt() uint32 {
	if x != nil {
		return x.UpdatedAt
	}
	return 0
}

type UpdateRegistrationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID        string `protobuf:"bytes,10,opt,name=ID,proto3" json:"ID,omitempty"`
	AppID     string `protobuf:"bytes,20,opt,name=AppID,proto3" json:"AppID,omitempty"`
	InviterID string `protobuf:"bytes,30,opt,name=InviterID,proto3" json:"InviterID,omitempty"`
}

func (x *UpdateRegistrationRequest) Reset() {
	*x = UpdateRegistrationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_inspire_gw_v1_invitation_registration_registration_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateRegistrationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateRegistrationRequest) ProtoMessage() {}

func (x *UpdateRegistrationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_inspire_gw_v1_invitation_registration_registration_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateRegistrationRequest.ProtoReflect.Descriptor instead.
func (*UpdateRegistrationRequest) Descriptor() ([]byte, []int) {
	return file_npool_inspire_gw_v1_invitation_registration_registration_proto_rawDescGZIP(), []int{1}
}

func (x *UpdateRegistrationRequest) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *UpdateRegistrationRequest) GetAppID() string {
	if x != nil {
		return x.AppID
	}
	return ""
}

func (x *UpdateRegistrationRequest) GetInviterID() string {
	if x != nil {
		return x.InviterID
	}
	return ""
}

type UpdateRegistrationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *Registration `protobuf:"bytes,10,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (x *UpdateRegistrationResponse) Reset() {
	*x = UpdateRegistrationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_inspire_gw_v1_invitation_registration_registration_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateRegistrationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateRegistrationResponse) ProtoMessage() {}

func (x *UpdateRegistrationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_inspire_gw_v1_invitation_registration_registration_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateRegistrationResponse.ProtoReflect.Descriptor instead.
func (*UpdateRegistrationResponse) Descriptor() ([]byte, []int) {
	return file_npool_inspire_gw_v1_invitation_registration_registration_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateRegistrationResponse) GetInfo() *Registration {
	if x != nil {
		return x.Info
	}
	return nil
}

type GetRegistrationsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AppID  string `protobuf:"bytes,10,opt,name=AppID,proto3" json:"AppID,omitempty"`
	Offset int32  `protobuf:"varint,20,opt,name=Offset,proto3" json:"Offset,omitempty"`
	Limit  int32  `protobuf:"varint,30,opt,name=Limit,proto3" json:"Limit,omitempty"`
}

func (x *GetRegistrationsRequest) Reset() {
	*x = GetRegistrationsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_inspire_gw_v1_invitation_registration_registration_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRegistrationsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRegistrationsRequest) ProtoMessage() {}

func (x *GetRegistrationsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_inspire_gw_v1_invitation_registration_registration_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRegistrationsRequest.ProtoReflect.Descriptor instead.
func (*GetRegistrationsRequest) Descriptor() ([]byte, []int) {
	return file_npool_inspire_gw_v1_invitation_registration_registration_proto_rawDescGZIP(), []int{3}
}

func (x *GetRegistrationsRequest) GetAppID() string {
	if x != nil {
		return x.AppID
	}
	return ""
}

func (x *GetRegistrationsRequest) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *GetRegistrationsRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type GetRegistrationsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Infos []*Registration `protobuf:"bytes,10,rep,name=Infos,proto3" json:"Infos,omitempty"`
	Total uint32          `protobuf:"varint,20,opt,name=Total,proto3" json:"Total,omitempty"`
}

func (x *GetRegistrationsResponse) Reset() {
	*x = GetRegistrationsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_inspire_gw_v1_invitation_registration_registration_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRegistrationsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRegistrationsResponse) ProtoMessage() {}

func (x *GetRegistrationsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_inspire_gw_v1_invitation_registration_registration_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRegistrationsResponse.ProtoReflect.Descriptor instead.
func (*GetRegistrationsResponse) Descriptor() ([]byte, []int) {
	return file_npool_inspire_gw_v1_invitation_registration_registration_proto_rawDescGZIP(), []int{4}
}

func (x *GetRegistrationsResponse) GetInfos() []*Registration {
	if x != nil {
		return x.Infos
	}
	return nil
}

func (x *GetRegistrationsResponse) GetTotal() uint32 {
	if x != nil {
		return x.Total
	}
	return 0
}

type GetAppRegistrationsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TargetAppID string `protobuf:"bytes,10,opt,name=TargetAppID,proto3" json:"TargetAppID,omitempty"`
	Offset      int32  `protobuf:"varint,20,opt,name=Offset,proto3" json:"Offset,omitempty"`
	Limit       int32  `protobuf:"varint,30,opt,name=Limit,proto3" json:"Limit,omitempty"`
}

func (x *GetAppRegistrationsRequest) Reset() {
	*x = GetAppRegistrationsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_inspire_gw_v1_invitation_registration_registration_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAppRegistrationsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAppRegistrationsRequest) ProtoMessage() {}

func (x *GetAppRegistrationsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_inspire_gw_v1_invitation_registration_registration_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAppRegistrationsRequest.ProtoReflect.Descriptor instead.
func (*GetAppRegistrationsRequest) Descriptor() ([]byte, []int) {
	return file_npool_inspire_gw_v1_invitation_registration_registration_proto_rawDescGZIP(), []int{5}
}

func (x *GetAppRegistrationsRequest) GetTargetAppID() string {
	if x != nil {
		return x.TargetAppID
	}
	return ""
}

func (x *GetAppRegistrationsRequest) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *GetAppRegistrationsRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type GetAppRegistrationsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Infos []*Registration `protobuf:"bytes,10,rep,name=Infos,proto3" json:"Infos,omitempty"`
	Total uint32          `protobuf:"varint,20,opt,name=Total,proto3" json:"Total,omitempty"`
}

func (x *GetAppRegistrationsResponse) Reset() {
	*x = GetAppRegistrationsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_inspire_gw_v1_invitation_registration_registration_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAppRegistrationsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAppRegistrationsResponse) ProtoMessage() {}

func (x *GetAppRegistrationsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_inspire_gw_v1_invitation_registration_registration_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAppRegistrationsResponse.ProtoReflect.Descriptor instead.
func (*GetAppRegistrationsResponse) Descriptor() ([]byte, []int) {
	return file_npool_inspire_gw_v1_invitation_registration_registration_proto_rawDescGZIP(), []int{6}
}

func (x *GetAppRegistrationsResponse) GetInfos() []*Registration {
	if x != nil {
		return x.Infos
	}
	return nil
}

func (x *GetAppRegistrationsResponse) GetTotal() uint32 {
	if x != nil {
		return x.Total
	}
	return 0
}

var File_npool_inspire_gw_v1_invitation_registration_registration_proto protoreflect.FileDescriptor

var file_npool_inspire_gw_v1_invitation_registration_registration_proto_rawDesc = []byte{
	0x0a, 0x3e, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x69, 0x6e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2f,
	0x67, 0x77, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2f, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x72, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x2a, 0x69, 0x6e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61,
	0x79, 0x2e, 0x69, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x72, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x6e, 0x70, 0x6f, 0x6f,
	0x6c, 0x2f, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb4, 0x03,
	0x0a, 0x0c, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e,
	0x0a, 0x02, 0x49, 0x44, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x12, 0x14,
	0x0a, 0x05, 0x41, 0x70, 0x70, 0x49, 0x44, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x41,
	0x70, 0x70, 0x49, 0x44, 0x12, 0x1c, 0x0a, 0x09, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x72, 0x49,
	0x44, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x72,
	0x49, 0x44, 0x12, 0x30, 0x0a, 0x13, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x72, 0x45, 0x6d, 0x61,
	0x69, 0x6c, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x28, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x13, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x72, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x41, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x12, 0x26, 0x0a, 0x0e, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x72, 0x50,
	0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x4f, 0x18, 0x32, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x49, 0x6e,
	0x76, 0x69, 0x74, 0x65, 0x72, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x4f, 0x12, 0x28, 0x0a, 0x0f,
	0x49, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x3c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x72, 0x55, 0x73,
	0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x65,
	0x65, 0x49, 0x44, 0x18, 0x46, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x49, 0x6e, 0x76, 0x69, 0x74,
	0x65, 0x65, 0x49, 0x44, 0x12, 0x30, 0x0a, 0x13, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x65, 0x45,
	0x6d, 0x61, 0x69, 0x6c, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x50, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x13, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x65, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x41,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x26, 0x0a, 0x0e, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x65,
	0x65, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x4f, 0x18, 0x5a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e,
	0x49, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x65, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x4f, 0x12, 0x28,
	0x0a, 0x0f, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x65, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x64, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x65,
	0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x6e, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x18, 0x78, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x22, 0x5f, 0x0a, 0x19, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49,
	0x44, 0x12, 0x14, 0x0a, 0x05, 0x41, 0x70, 0x70, 0x49, 0x44, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x41, 0x70, 0x70, 0x49, 0x44, 0x12, 0x1c, 0x0a, 0x09, 0x49, 0x6e, 0x76, 0x69, 0x74,
	0x65, 0x72, 0x49, 0x44, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x49, 0x6e, 0x76, 0x69,
	0x74, 0x65, 0x72, 0x49, 0x44, 0x22, 0x6a, 0x0a, 0x1a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x4c, 0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x38, 0x2e, 0x69, 0x6e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x67, 0x61, 0x74, 0x65,
	0x77, 0x61, 0x79, 0x2e, 0x69, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x72,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x52,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x04, 0x49, 0x6e, 0x66,
	0x6f, 0x22, 0x5d, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05,
	0x41, 0x70, 0x70, 0x49, 0x44, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x41, 0x70, 0x70,
	0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x14, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x06, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x4c, 0x69,
	0x6d, 0x69, 0x74, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x4c, 0x69, 0x6d, 0x69, 0x74,
	0x22, 0x80, 0x01, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4e, 0x0a,
	0x05, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x38, 0x2e, 0x69,
	0x6e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x69,
	0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x05, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x12, 0x14, 0x0a,
	0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x14, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x54, 0x6f,
	0x74, 0x61, 0x6c, 0x22, 0x6c, 0x0a, 0x1a, 0x47, 0x65, 0x74, 0x41, 0x70, 0x70, 0x52, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x20, 0x0a, 0x0b, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x41, 0x70, 0x70, 0x49, 0x44,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x41, 0x70,
	0x70, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x14, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x06, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x4c,
	0x69, 0x6d, 0x69, 0x74, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x4c, 0x69, 0x6d, 0x69,
	0x74, 0x22, 0x83, 0x01, 0x0a, 0x1b, 0x47, 0x65, 0x74, 0x41, 0x70, 0x70, 0x52, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x4e, 0x0a, 0x05, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x38, 0x2e, 0x69, 0x6e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77,
	0x61, 0x79, 0x2e, 0x69, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x72, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x05, 0x49, 0x6e, 0x66, 0x6f,
	0x73, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x14, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x32, 0x82, 0x05, 0x0a, 0x07, 0x47, 0x61, 0x74, 0x65,
	0x77, 0x61, 0x79, 0x12, 0xd1, 0x01, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x45, 0x2e, 0x69, 0x6e, 0x73,
	0x70, 0x69, 0x72, 0x65, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x69, 0x6e, 0x76,
	0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x46, 0x2e, 0x69, 0x6e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x67, 0x61, 0x74, 0x65,
	0x77, 0x61, 0x79, 0x2e, 0x69, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x72,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2c, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x26, 0x3a, 0x01, 0x2a, 0x22, 0x21, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x2f, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x69, 0x6e, 0x76,
	0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0xc9, 0x01, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x52,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x43, 0x2e, 0x69,
	0x6e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x69,
	0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x44, 0x2e, 0x69, 0x6e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x67, 0x61, 0x74, 0x65,
	0x77, 0x61, 0x79, 0x2e, 0x69, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x72,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x47,
	0x65, 0x74, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x24, 0x3a,
	0x01, 0x2a, 0x22, 0x1f, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x65, 0x74, 0x2f, 0x72, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x69, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x12, 0xd6, 0x01, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x41, 0x70, 0x70, 0x52, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x46, 0x2e, 0x69, 0x6e,
	0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x69, 0x6e,
	0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x70, 0x70, 0x52,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x47, 0x2e, 0x69, 0x6e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x67, 0x61,
	0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x69, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31,
	0x2e, 0x47, 0x65, 0x74, 0x41, 0x70, 0x70, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2e, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x28, 0x3a, 0x01, 0x2a, 0x22, 0x23, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x65, 0x74,
	0x2f, 0x61, 0x70, 0x70, 0x2f, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x69, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x42, 0x4e, 0x5a, 0x4c,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4e, 0x70, 0x6f, 0x6f, 0x6c,
	0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x2f, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x69, 0x6e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2f, 0x67,
	0x77, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f,
	0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_npool_inspire_gw_v1_invitation_registration_registration_proto_rawDescOnce sync.Once
	file_npool_inspire_gw_v1_invitation_registration_registration_proto_rawDescData = file_npool_inspire_gw_v1_invitation_registration_registration_proto_rawDesc
)

func file_npool_inspire_gw_v1_invitation_registration_registration_proto_rawDescGZIP() []byte {
	file_npool_inspire_gw_v1_invitation_registration_registration_proto_rawDescOnce.Do(func() {
		file_npool_inspire_gw_v1_invitation_registration_registration_proto_rawDescData = protoimpl.X.CompressGZIP(file_npool_inspire_gw_v1_invitation_registration_registration_proto_rawDescData)
	})
	return file_npool_inspire_gw_v1_invitation_registration_registration_proto_rawDescData
}

var file_npool_inspire_gw_v1_invitation_registration_registration_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_npool_inspire_gw_v1_invitation_registration_registration_proto_goTypes = []interface{}{
	(*Registration)(nil),                // 0: inspire.gateway.invitation.registration.v1.Registration
	(*UpdateRegistrationRequest)(nil),   // 1: inspire.gateway.invitation.registration.v1.UpdateRegistrationRequest
	(*UpdateRegistrationResponse)(nil),  // 2: inspire.gateway.invitation.registration.v1.UpdateRegistrationResponse
	(*GetRegistrationsRequest)(nil),     // 3: inspire.gateway.invitation.registration.v1.GetRegistrationsRequest
	(*GetRegistrationsResponse)(nil),    // 4: inspire.gateway.invitation.registration.v1.GetRegistrationsResponse
	(*GetAppRegistrationsRequest)(nil),  // 5: inspire.gateway.invitation.registration.v1.GetAppRegistrationsRequest
	(*GetAppRegistrationsResponse)(nil), // 6: inspire.gateway.invitation.registration.v1.GetAppRegistrationsResponse
}
var file_npool_inspire_gw_v1_invitation_registration_registration_proto_depIdxs = []int32{
	0, // 0: inspire.gateway.invitation.registration.v1.UpdateRegistrationResponse.Info:type_name -> inspire.gateway.invitation.registration.v1.Registration
	0, // 1: inspire.gateway.invitation.registration.v1.GetRegistrationsResponse.Infos:type_name -> inspire.gateway.invitation.registration.v1.Registration
	0, // 2: inspire.gateway.invitation.registration.v1.GetAppRegistrationsResponse.Infos:type_name -> inspire.gateway.invitation.registration.v1.Registration
	1, // 3: inspire.gateway.invitation.registration.v1.Gateway.UpdateRegistration:input_type -> inspire.gateway.invitation.registration.v1.UpdateRegistrationRequest
	3, // 4: inspire.gateway.invitation.registration.v1.Gateway.GetRegistrations:input_type -> inspire.gateway.invitation.registration.v1.GetRegistrationsRequest
	5, // 5: inspire.gateway.invitation.registration.v1.Gateway.GetAppRegistrations:input_type -> inspire.gateway.invitation.registration.v1.GetAppRegistrationsRequest
	2, // 6: inspire.gateway.invitation.registration.v1.Gateway.UpdateRegistration:output_type -> inspire.gateway.invitation.registration.v1.UpdateRegistrationResponse
	4, // 7: inspire.gateway.invitation.registration.v1.Gateway.GetRegistrations:output_type -> inspire.gateway.invitation.registration.v1.GetRegistrationsResponse
	6, // 8: inspire.gateway.invitation.registration.v1.Gateway.GetAppRegistrations:output_type -> inspire.gateway.invitation.registration.v1.GetAppRegistrationsResponse
	6, // [6:9] is the sub-list for method output_type
	3, // [3:6] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_npool_inspire_gw_v1_invitation_registration_registration_proto_init() }
func file_npool_inspire_gw_v1_invitation_registration_registration_proto_init() {
	if File_npool_inspire_gw_v1_invitation_registration_registration_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_npool_inspire_gw_v1_invitation_registration_registration_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Registration); i {
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
		file_npool_inspire_gw_v1_invitation_registration_registration_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateRegistrationRequest); i {
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
		file_npool_inspire_gw_v1_invitation_registration_registration_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateRegistrationResponse); i {
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
		file_npool_inspire_gw_v1_invitation_registration_registration_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRegistrationsRequest); i {
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
		file_npool_inspire_gw_v1_invitation_registration_registration_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRegistrationsResponse); i {
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
		file_npool_inspire_gw_v1_invitation_registration_registration_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAppRegistrationsRequest); i {
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
		file_npool_inspire_gw_v1_invitation_registration_registration_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAppRegistrationsResponse); i {
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
			RawDescriptor: file_npool_inspire_gw_v1_invitation_registration_registration_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_npool_inspire_gw_v1_invitation_registration_registration_proto_goTypes,
		DependencyIndexes: file_npool_inspire_gw_v1_invitation_registration_registration_proto_depIdxs,
		MessageInfos:      file_npool_inspire_gw_v1_invitation_registration_registration_proto_msgTypes,
	}.Build()
	File_npool_inspire_gw_v1_invitation_registration_registration_proto = out.File
	file_npool_inspire_gw_v1_invitation_registration_registration_proto_rawDesc = nil
	file_npool_inspire_gw_v1_invitation_registration_registration_proto_goTypes = nil
	file_npool_inspire_gw_v1_invitation_registration_registration_proto_depIdxs = nil
}
