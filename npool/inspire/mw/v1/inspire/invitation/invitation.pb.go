// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.18.1
// source: npool/inspire/mw/v1/inspire/invitation/invitation.proto

package invitation

import (
	_ "github.com/NpoolPlatform/message/npool"
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

type Invitation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: sql "inviter_id"
	InviterID string `protobuf:"bytes,10,opt,name=InviterID,proto3" json:"InviterID,omitempty"`
	// @inject_tag: sql "invitee_id"
	InviteeID string `protobuf:"bytes,20,opt,name=InviteeID,proto3" json:"InviteeID,omitempty"`
	// @inject_tag: sql "create_at"
	CreatedAt uint32 `protobuf:"varint,30,opt,name=CreatedAt,proto3" json:"CreatedAt,omitempty"`
}

func (x *Invitation) Reset() {
	*x = Invitation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_inspire_mw_v1_inspire_invitation_invitation_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Invitation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Invitation) ProtoMessage() {}

func (x *Invitation) ProtoReflect() protoreflect.Message {
	mi := &file_npool_inspire_mw_v1_inspire_invitation_invitation_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Invitation.ProtoReflect.Descriptor instead.
func (*Invitation) Descriptor() ([]byte, []int) {
	return file_npool_inspire_mw_v1_inspire_invitation_invitation_proto_rawDescGZIP(), []int{0}
}

func (x *Invitation) GetInviterID() string {
	if x != nil {
		return x.InviterID
	}
	return ""
}

func (x *Invitation) GetInviteeID() string {
	if x != nil {
		return x.InviteeID
	}
	return ""
}

func (x *Invitation) GetCreatedAt() uint32 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

type GetInviteesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AppID   string   `protobuf:"bytes,10,opt,name=AppID,proto3" json:"AppID,omitempty"`
	UserIDs []string `protobuf:"bytes,20,rep,name=UserIDs,proto3" json:"UserIDs,omitempty"`
	Offset  int32    `protobuf:"varint,30,opt,name=Offset,proto3" json:"Offset,omitempty"`
	Limit   int32    `protobuf:"varint,40,opt,name=Limit,proto3" json:"Limit,omitempty"`
}

func (x *GetInviteesRequest) Reset() {
	*x = GetInviteesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_inspire_mw_v1_inspire_invitation_invitation_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetInviteesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetInviteesRequest) ProtoMessage() {}

func (x *GetInviteesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_inspire_mw_v1_inspire_invitation_invitation_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetInviteesRequest.ProtoReflect.Descriptor instead.
func (*GetInviteesRequest) Descriptor() ([]byte, []int) {
	return file_npool_inspire_mw_v1_inspire_invitation_invitation_proto_rawDescGZIP(), []int{1}
}

func (x *GetInviteesRequest) GetAppID() string {
	if x != nil {
		return x.AppID
	}
	return ""
}

func (x *GetInviteesRequest) GetUserIDs() []string {
	if x != nil {
		return x.UserIDs
	}
	return nil
}

func (x *GetInviteesRequest) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *GetInviteesRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type GetInviteesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Infos []*Invitation `protobuf:"bytes,10,rep,name=Infos,proto3" json:"Infos,omitempty"`
	Total uint32        `protobuf:"varint,20,opt,name=Total,proto3" json:"Total,omitempty"`
}

func (x *GetInviteesResponse) Reset() {
	*x = GetInviteesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_inspire_mw_v1_inspire_invitation_invitation_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetInviteesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetInviteesResponse) ProtoMessage() {}

func (x *GetInviteesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_inspire_mw_v1_inspire_invitation_invitation_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetInviteesResponse.ProtoReflect.Descriptor instead.
func (*GetInviteesResponse) Descriptor() ([]byte, []int) {
	return file_npool_inspire_mw_v1_inspire_invitation_invitation_proto_rawDescGZIP(), []int{2}
}

func (x *GetInviteesResponse) GetInfos() []*Invitation {
	if x != nil {
		return x.Infos
	}
	return nil
}

func (x *GetInviteesResponse) GetTotal() uint32 {
	if x != nil {
		return x.Total
	}
	return 0
}

type GetInvitersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AppID  string `protobuf:"bytes,10,opt,name=AppID,proto3" json:"AppID,omitempty"`
	UserID string `protobuf:"bytes,20,opt,name=UserID,proto3" json:"UserID,omitempty"`
	Offset int32  `protobuf:"varint,30,opt,name=Offset,proto3" json:"Offset,omitempty"`
	Limit  int32  `protobuf:"varint,40,opt,name=Limit,proto3" json:"Limit,omitempty"`
}

func (x *GetInvitersRequest) Reset() {
	*x = GetInvitersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_inspire_mw_v1_inspire_invitation_invitation_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetInvitersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetInvitersRequest) ProtoMessage() {}

func (x *GetInvitersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_inspire_mw_v1_inspire_invitation_invitation_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetInvitersRequest.ProtoReflect.Descriptor instead.
func (*GetInvitersRequest) Descriptor() ([]byte, []int) {
	return file_npool_inspire_mw_v1_inspire_invitation_invitation_proto_rawDescGZIP(), []int{3}
}

func (x *GetInvitersRequest) GetAppID() string {
	if x != nil {
		return x.AppID
	}
	return ""
}

func (x *GetInvitersRequest) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

func (x *GetInvitersRequest) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *GetInvitersRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type GetInvitersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Infos []*Invitation `protobuf:"bytes,10,rep,name=Infos,proto3" json:"Infos,omitempty"`
	Total uint32        `protobuf:"varint,20,opt,name=Total,proto3" json:"Total,omitempty"`
}

func (x *GetInvitersResponse) Reset() {
	*x = GetInvitersResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_inspire_mw_v1_inspire_invitation_invitation_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetInvitersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetInvitersResponse) ProtoMessage() {}

func (x *GetInvitersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_inspire_mw_v1_inspire_invitation_invitation_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetInvitersResponse.ProtoReflect.Descriptor instead.
func (*GetInvitersResponse) Descriptor() ([]byte, []int) {
	return file_npool_inspire_mw_v1_inspire_invitation_invitation_proto_rawDescGZIP(), []int{4}
}

func (x *GetInvitersResponse) GetInfos() []*Invitation {
	if x != nil {
		return x.Infos
	}
	return nil
}

func (x *GetInvitersResponse) GetTotal() uint32 {
	if x != nil {
		return x.Total
	}
	return 0
}

var File_npool_inspire_mw_v1_inspire_invitation_invitation_proto protoreflect.FileDescriptor

var file_npool_inspire_mw_v1_inspire_invitation_invitation_proto_rawDesc = []byte{
	0x0a, 0x37, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x69, 0x6e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2f,
	0x6d, 0x77, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x6e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2f, 0x69, 0x6e,
	0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x69, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x28, 0x69, 0x6e, 0x73, 0x70, 0x69,
	0x72, 0x65, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x69, 0x6e,
	0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x69, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11,
	0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x66, 0x0a, 0x0a, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x1c, 0x0a, 0x09, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x72, 0x49, 0x44, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x72, 0x49, 0x44, 0x12, 0x1c, 0x0a,
	0x09, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x65, 0x49, 0x44, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x65, 0x49, 0x44, 0x12, 0x1c, 0x0a, 0x09, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x72, 0x0a, 0x12, 0x47, 0x65, 0x74,
	0x49, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x14, 0x0a, 0x05, 0x41, 0x70, 0x70, 0x49, 0x44, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x41, 0x70, 0x70, 0x49, 0x44, 0x12, 0x18, 0x0a, 0x07, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x73,
	0x18, 0x14, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x73, 0x12,
	0x16, 0x0a, 0x06, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x06, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x4c, 0x69, 0x6d, 0x69, 0x74,
	0x18, 0x28, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x22, 0x77, 0x0a,
	0x13, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4a, 0x0a, 0x05, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x18, 0x0a, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x34, 0x2e, 0x69, 0x6e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x6d, 0x69,
	0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x69, 0x6e, 0x73, 0x70, 0x69, 0x72, 0x65,
	0x2e, 0x69, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x49,
	0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x05, 0x49, 0x6e, 0x66, 0x6f, 0x73,
	0x12, 0x14, 0x0a, 0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x14, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x22, 0x70, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x76,
	0x69, 0x74, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05,
	0x41, 0x70, 0x70, 0x49, 0x44, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x41, 0x70, 0x70,
	0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x14, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x4f, 0x66,
	0x66, 0x73, 0x65, 0x74, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x4f, 0x66, 0x66, 0x73,
	0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x28, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x05, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x22, 0x77, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x49,
	0x6e, 0x76, 0x69, 0x74, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x4a, 0x0a, 0x05, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x34,
	0x2e, 0x69, 0x6e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77,
	0x61, 0x72, 0x65, 0x2e, 0x69, 0x6e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x69, 0x6e, 0x76, 0x69,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x05, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x54,
	0x6f, 0x74, 0x61, 0x6c, 0x18, 0x14, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x54, 0x6f, 0x74, 0x61,
	0x6c, 0x32, 0xaa, 0x02, 0x0a, 0x0a, 0x4d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65,
	0x12, 0x8c, 0x01, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x65, 0x73,
	0x12, 0x3c, 0x2e, 0x69, 0x6e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c,
	0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x69, 0x6e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x69, 0x6e,
	0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x49,
	0x6e, 0x76, 0x69, 0x74, 0x65, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3d,
	0x2e, 0x69, 0x6e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77,
	0x61, 0x72, 0x65, 0x2e, 0x69, 0x6e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x69, 0x6e, 0x76, 0x69,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x76,
	0x69, 0x74, 0x65, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x8c, 0x01, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x72, 0x73, 0x12,
	0x3c, 0x2e, 0x69, 0x6e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65,
	0x77, 0x61, 0x72, 0x65, 0x2e, 0x69, 0x6e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x69, 0x6e, 0x76,
	0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x49, 0x6e,
	0x76, 0x69, 0x74, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3d, 0x2e,
	0x69, 0x6e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61,
	0x72, 0x65, 0x2e, 0x69, 0x6e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x69, 0x6e, 0x76, 0x69, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x76, 0x69,
	0x74, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x49,
	0x5a, 0x47, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4e, 0x70, 0x6f,
	0x6f, 0x6c, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x2f, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x69, 0x6e, 0x73, 0x70, 0x69, 0x72, 0x65,
	0x2f, 0x6d, 0x77, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x6e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2f, 0x69,
	0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_npool_inspire_mw_v1_inspire_invitation_invitation_proto_rawDescOnce sync.Once
	file_npool_inspire_mw_v1_inspire_invitation_invitation_proto_rawDescData = file_npool_inspire_mw_v1_inspire_invitation_invitation_proto_rawDesc
)

func file_npool_inspire_mw_v1_inspire_invitation_invitation_proto_rawDescGZIP() []byte {
	file_npool_inspire_mw_v1_inspire_invitation_invitation_proto_rawDescOnce.Do(func() {
		file_npool_inspire_mw_v1_inspire_invitation_invitation_proto_rawDescData = protoimpl.X.CompressGZIP(file_npool_inspire_mw_v1_inspire_invitation_invitation_proto_rawDescData)
	})
	return file_npool_inspire_mw_v1_inspire_invitation_invitation_proto_rawDescData
}

var file_npool_inspire_mw_v1_inspire_invitation_invitation_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_npool_inspire_mw_v1_inspire_invitation_invitation_proto_goTypes = []interface{}{
	(*Invitation)(nil),          // 0: inspire.middleware.inspire.invitation.v1.Invitation
	(*GetInviteesRequest)(nil),  // 1: inspire.middleware.inspire.invitation.v1.GetInviteesRequest
	(*GetInviteesResponse)(nil), // 2: inspire.middleware.inspire.invitation.v1.GetInviteesResponse
	(*GetInvitersRequest)(nil),  // 3: inspire.middleware.inspire.invitation.v1.GetInvitersRequest
	(*GetInvitersResponse)(nil), // 4: inspire.middleware.inspire.invitation.v1.GetInvitersResponse
}
var file_npool_inspire_mw_v1_inspire_invitation_invitation_proto_depIdxs = []int32{
	0, // 0: inspire.middleware.inspire.invitation.v1.GetInviteesResponse.Infos:type_name -> inspire.middleware.inspire.invitation.v1.Invitation
	0, // 1: inspire.middleware.inspire.invitation.v1.GetInvitersResponse.Infos:type_name -> inspire.middleware.inspire.invitation.v1.Invitation
	1, // 2: inspire.middleware.inspire.invitation.v1.Middleware.GetInvitees:input_type -> inspire.middleware.inspire.invitation.v1.GetInviteesRequest
	3, // 3: inspire.middleware.inspire.invitation.v1.Middleware.GetInviters:input_type -> inspire.middleware.inspire.invitation.v1.GetInvitersRequest
	2, // 4: inspire.middleware.inspire.invitation.v1.Middleware.GetInvitees:output_type -> inspire.middleware.inspire.invitation.v1.GetInviteesResponse
	4, // 5: inspire.middleware.inspire.invitation.v1.Middleware.GetInviters:output_type -> inspire.middleware.inspire.invitation.v1.GetInvitersResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_npool_inspire_mw_v1_inspire_invitation_invitation_proto_init() }
func file_npool_inspire_mw_v1_inspire_invitation_invitation_proto_init() {
	if File_npool_inspire_mw_v1_inspire_invitation_invitation_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_npool_inspire_mw_v1_inspire_invitation_invitation_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Invitation); i {
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
		file_npool_inspire_mw_v1_inspire_invitation_invitation_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetInviteesRequest); i {
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
		file_npool_inspire_mw_v1_inspire_invitation_invitation_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetInviteesResponse); i {
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
		file_npool_inspire_mw_v1_inspire_invitation_invitation_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetInvitersRequest); i {
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
		file_npool_inspire_mw_v1_inspire_invitation_invitation_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetInvitersResponse); i {
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
			RawDescriptor: file_npool_inspire_mw_v1_inspire_invitation_invitation_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_npool_inspire_mw_v1_inspire_invitation_invitation_proto_goTypes,
		DependencyIndexes: file_npool_inspire_mw_v1_inspire_invitation_invitation_proto_depIdxs,
		MessageInfos:      file_npool_inspire_mw_v1_inspire_invitation_invitation_proto_msgTypes,
	}.Build()
	File_npool_inspire_mw_v1_inspire_invitation_invitation_proto = out.File
	file_npool_inspire_mw_v1_inspire_invitation_invitation_proto_rawDesc = nil
	file_npool_inspire_mw_v1_inspire_invitation_invitation_proto_goTypes = nil
	file_npool_inspire_mw_v1_inspire_invitation_invitation_proto_depIdxs = nil
}
