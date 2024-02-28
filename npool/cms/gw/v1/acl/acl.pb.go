// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.18.1
// source: npool/cms/gw/v1/acl/acl.proto

package acl

import (
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

type ACL struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID         uint32 `protobuf:"varint,10,opt,name=ID,proto3" json:"ID,omitempty"`
	EntID      string `protobuf:"bytes,20,opt,name=EntID,proto3" json:"EntID,omitempty"`
	AppID      string `protobuf:"bytes,30,opt,name=AppID,proto3" json:"AppID,omitempty"`
	RoleID     string `protobuf:"bytes,40,opt,name=RoleID,proto3" json:"RoleID,omitempty"`
	Role       string `protobuf:"bytes,50,opt,name=Role,proto3" json:"Role,omitempty"`
	ArticleKey string `protobuf:"bytes,60,opt,name=ArticleKey,proto3" json:"ArticleKey,omitempty"`
	CreatedAt  uint32 `protobuf:"varint,1000,opt,name=CreatedAt,proto3" json:"CreatedAt,omitempty"`
	UpdatedAt  uint32 `protobuf:"varint,1010,opt,name=UpdatedAt,proto3" json:"UpdatedAt,omitempty"`
}

func (x *ACL) Reset() {
	*x = ACL{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_cms_gw_v1_acl_acl_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ACL) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ACL) ProtoMessage() {}

func (x *ACL) ProtoReflect() protoreflect.Message {
	mi := &file_npool_cms_gw_v1_acl_acl_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ACL.ProtoReflect.Descriptor instead.
func (*ACL) Descriptor() ([]byte, []int) {
	return file_npool_cms_gw_v1_acl_acl_proto_rawDescGZIP(), []int{0}
}

func (x *ACL) GetID() uint32 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *ACL) GetEntID() string {
	if x != nil {
		return x.EntID
	}
	return ""
}

func (x *ACL) GetAppID() string {
	if x != nil {
		return x.AppID
	}
	return ""
}

func (x *ACL) GetRoleID() string {
	if x != nil {
		return x.RoleID
	}
	return ""
}

func (x *ACL) GetRole() string {
	if x != nil {
		return x.Role
	}
	return ""
}

func (x *ACL) GetArticleKey() string {
	if x != nil {
		return x.ArticleKey
	}
	return ""
}

func (x *ACL) GetCreatedAt() uint32 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *ACL) GetUpdatedAt() uint32 {
	if x != nil {
		return x.UpdatedAt
	}
	return 0
}

type CreateACLRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AppID      string `protobuf:"bytes,10,opt,name=AppID,proto3" json:"AppID,omitempty"`
	RoleID     string `protobuf:"bytes,20,opt,name=RoleID,proto3" json:"RoleID,omitempty"`
	ArticleKey string `protobuf:"bytes,30,opt,name=ArticleKey,proto3" json:"ArticleKey,omitempty"`
}

func (x *CreateACLRequest) Reset() {
	*x = CreateACLRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_cms_gw_v1_acl_acl_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateACLRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateACLRequest) ProtoMessage() {}

func (x *CreateACLRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_cms_gw_v1_acl_acl_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateACLRequest.ProtoReflect.Descriptor instead.
func (*CreateACLRequest) Descriptor() ([]byte, []int) {
	return file_npool_cms_gw_v1_acl_acl_proto_rawDescGZIP(), []int{1}
}

func (x *CreateACLRequest) GetAppID() string {
	if x != nil {
		return x.AppID
	}
	return ""
}

func (x *CreateACLRequest) GetRoleID() string {
	if x != nil {
		return x.RoleID
	}
	return ""
}

func (x *CreateACLRequest) GetArticleKey() string {
	if x != nil {
		return x.ArticleKey
	}
	return ""
}

type CreateACLResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *ACL `protobuf:"bytes,10,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (x *CreateACLResponse) Reset() {
	*x = CreateACLResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_cms_gw_v1_acl_acl_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateACLResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateACLResponse) ProtoMessage() {}

func (x *CreateACLResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_cms_gw_v1_acl_acl_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateACLResponse.ProtoReflect.Descriptor instead.
func (*CreateACLResponse) Descriptor() ([]byte, []int) {
	return file_npool_cms_gw_v1_acl_acl_proto_rawDescGZIP(), []int{2}
}

func (x *CreateACLResponse) GetInfo() *ACL {
	if x != nil {
		return x.Info
	}
	return nil
}

type GetACLsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AppID      string `protobuf:"bytes,10,opt,name=AppID,proto3" json:"AppID,omitempty"`
	ArticleKey string `protobuf:"bytes,20,opt,name=ArticleKey,proto3" json:"ArticleKey,omitempty"`
	Offset     int32  `protobuf:"varint,30,opt,name=Offset,proto3" json:"Offset,omitempty"`
	Limit      int32  `protobuf:"varint,40,opt,name=Limit,proto3" json:"Limit,omitempty"`
}

func (x *GetACLsRequest) Reset() {
	*x = GetACLsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_cms_gw_v1_acl_acl_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetACLsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetACLsRequest) ProtoMessage() {}

func (x *GetACLsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_cms_gw_v1_acl_acl_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetACLsRequest.ProtoReflect.Descriptor instead.
func (*GetACLsRequest) Descriptor() ([]byte, []int) {
	return file_npool_cms_gw_v1_acl_acl_proto_rawDescGZIP(), []int{3}
}

func (x *GetACLsRequest) GetAppID() string {
	if x != nil {
		return x.AppID
	}
	return ""
}

func (x *GetACLsRequest) GetArticleKey() string {
	if x != nil {
		return x.ArticleKey
	}
	return ""
}

func (x *GetACLsRequest) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *GetACLsRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type GetACLsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Infos []*ACL `protobuf:"bytes,10,rep,name=Infos,proto3" json:"Infos,omitempty"`
	Total uint32 `protobuf:"varint,20,opt,name=Total,proto3" json:"Total,omitempty"`
}

func (x *GetACLsResponse) Reset() {
	*x = GetACLsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_cms_gw_v1_acl_acl_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetACLsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetACLsResponse) ProtoMessage() {}

func (x *GetACLsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_cms_gw_v1_acl_acl_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetACLsResponse.ProtoReflect.Descriptor instead.
func (*GetACLsResponse) Descriptor() ([]byte, []int) {
	return file_npool_cms_gw_v1_acl_acl_proto_rawDescGZIP(), []int{4}
}

func (x *GetACLsResponse) GetInfos() []*ACL {
	if x != nil {
		return x.Infos
	}
	return nil
}

func (x *GetACLsResponse) GetTotal() uint32 {
	if x != nil {
		return x.Total
	}
	return 0
}

type DeleteACLRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID    uint32 `protobuf:"varint,10,opt,name=ID,proto3" json:"ID,omitempty"`
	EntID string `protobuf:"bytes,20,opt,name=EntID,proto3" json:"EntID,omitempty"`
	AppID string `protobuf:"bytes,30,opt,name=AppID,proto3" json:"AppID,omitempty"`
}

func (x *DeleteACLRequest) Reset() {
	*x = DeleteACLRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_cms_gw_v1_acl_acl_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteACLRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteACLRequest) ProtoMessage() {}

func (x *DeleteACLRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_cms_gw_v1_acl_acl_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteACLRequest.ProtoReflect.Descriptor instead.
func (*DeleteACLRequest) Descriptor() ([]byte, []int) {
	return file_npool_cms_gw_v1_acl_acl_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteACLRequest) GetID() uint32 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *DeleteACLRequest) GetEntID() string {
	if x != nil {
		return x.EntID
	}
	return ""
}

func (x *DeleteACLRequest) GetAppID() string {
	if x != nil {
		return x.AppID
	}
	return ""
}

type DeleteACLResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *ACL `protobuf:"bytes,10,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (x *DeleteACLResponse) Reset() {
	*x = DeleteACLResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_cms_gw_v1_acl_acl_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteACLResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteACLResponse) ProtoMessage() {}

func (x *DeleteACLResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_cms_gw_v1_acl_acl_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteACLResponse.ProtoReflect.Descriptor instead.
func (*DeleteACLResponse) Descriptor() ([]byte, []int) {
	return file_npool_cms_gw_v1_acl_acl_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteACLResponse) GetInfo() *ACL {
	if x != nil {
		return x.Info
	}
	return nil
}

var File_npool_cms_gw_v1_acl_acl_proto protoreflect.FileDescriptor

var file_npool_cms_gw_v1_acl_acl_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x63, 0x6d, 0x73, 0x2f, 0x67, 0x77, 0x2f, 0x76,
	0x31, 0x2f, 0x61, 0x63, 0x6c, 0x2f, 0x61, 0x63, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x12, 0x63, 0x6d, 0x73, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x61, 0x63, 0x6c,
	0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xcb, 0x01, 0x0a, 0x03, 0x41, 0x43, 0x4c, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x45, 0x6e, 0x74,
	0x49, 0x44, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x45, 0x6e, 0x74, 0x49, 0x44, 0x12,
	0x14, 0x0a, 0x05, 0x41, 0x70, 0x70, 0x49, 0x44, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x41, 0x70, 0x70, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x52, 0x6f, 0x6c, 0x65, 0x49, 0x44, 0x18,
	0x28, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x52, 0x6f, 0x6c, 0x65, 0x49, 0x44, 0x12, 0x12, 0x0a,
	0x04, 0x52, 0x6f, 0x6c, 0x65, 0x18, 0x32, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x52, 0x6f, 0x6c,
	0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x4b, 0x65, 0x79, 0x18,
	0x3c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x4b, 0x65,
	0x79, 0x12, 0x1d, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0xe8,
	0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x12, 0x1d, 0x0a, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0xf2, 0x07,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22,
	0x60, 0x0a, 0x10, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x43, 0x4c, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x41, 0x70, 0x70, 0x49, 0x44, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x41, 0x70, 0x70, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x52, 0x6f, 0x6c,
	0x65, 0x49, 0x44, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x52, 0x6f, 0x6c, 0x65, 0x49,
	0x44, 0x12, 0x1e, 0x0a, 0x0a, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x4b, 0x65, 0x79, 0x18,
	0x1e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x4b, 0x65,
	0x79, 0x22, 0x40, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x43, 0x4c, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2b, 0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x63, 0x6d, 0x73, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77,
	0x61, 0x79, 0x2e, 0x61, 0x63, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x43, 0x4c, 0x52, 0x04, 0x49,
	0x6e, 0x66, 0x6f, 0x22, 0x74, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x41, 0x43, 0x4c, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x41, 0x70, 0x70, 0x49, 0x44, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x41, 0x70, 0x70, 0x49, 0x44, 0x12, 0x1e, 0x0a, 0x0a, 0x41,
	0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x4b, 0x65, 0x79, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x4b, 0x65, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x4f,
	0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x4f, 0x66, 0x66,
	0x73, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x28, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x05, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x22, 0x56, 0x0a, 0x0f, 0x47, 0x65, 0x74,
	0x41, 0x43, 0x4c, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2d, 0x0a, 0x05,
	0x49, 0x6e, 0x66, 0x6f, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x63, 0x6d,
	0x73, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x61, 0x63, 0x6c, 0x2e, 0x76, 0x31,
	0x2e, 0x41, 0x43, 0x4c, 0x52, 0x05, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x54,
	0x6f, 0x74, 0x61, 0x6c, 0x18, 0x14, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x54, 0x6f, 0x74, 0x61,
	0x6c, 0x22, 0x4e, 0x0a, 0x10, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x43, 0x4c, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x02, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x45, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x14,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x45, 0x6e, 0x74, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x41,
	0x70, 0x70, 0x49, 0x44, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x41, 0x70, 0x70, 0x49,
	0x44, 0x22, 0x40, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x43, 0x4c, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2b, 0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x63, 0x6d, 0x73, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77,
	0x61, 0x79, 0x2e, 0x61, 0x63, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x43, 0x4c, 0x52, 0x04, 0x49,
	0x6e, 0x66, 0x6f, 0x32, 0xe0, 0x02, 0x0a, 0x07, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x12,
	0x73, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x43, 0x4c, 0x12, 0x24, 0x2e, 0x63,
	0x6d, 0x73, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x61, 0x63, 0x6c, 0x2e, 0x76,
	0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x43, 0x4c, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x25, 0x2e, 0x63, 0x6d, 0x73, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79,
	0x2e, 0x61, 0x63, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x43,
	0x4c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x19, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x13, 0x3a, 0x01, 0x2a, 0x22, 0x0e, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x2f, 0x61, 0x63, 0x6c, 0x12, 0x6b, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x41, 0x43, 0x4c, 0x73, 0x12,
	0x22, 0x2e, 0x63, 0x6d, 0x73, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x61, 0x63,
	0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x43, 0x4c, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x63, 0x6d, 0x73, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61,
	0x79, 0x2e, 0x61, 0x63, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x43, 0x4c, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x17, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x11,
	0x3a, 0x01, 0x2a, 0x22, 0x0c, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x65, 0x74, 0x2f, 0x61, 0x63, 0x6c,
	0x73, 0x12, 0x73, 0x0a, 0x09, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x43, 0x4c, 0x12, 0x24,
	0x2e, 0x63, 0x6d, 0x73, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x61, 0x63, 0x6c,
	0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x43, 0x4c, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x63, 0x6d, 0x73, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77,
	0x61, 0x79, 0x2e, 0x61, 0x63, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x41, 0x43, 0x4c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x19, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x13, 0x3a, 0x01, 0x2a, 0x22, 0x0e, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x2f, 0x61, 0x63, 0x6c, 0x42, 0x36, 0x5a, 0x34, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4e, 0x70, 0x6f, 0x6f, 0x6c, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f,
	0x72, 0x6d, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x6e, 0x70, 0x6f, 0x6f, 0x6c,
	0x2f, 0x63, 0x6d, 0x73, 0x2f, 0x67, 0x77, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x63, 0x6c, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_npool_cms_gw_v1_acl_acl_proto_rawDescOnce sync.Once
	file_npool_cms_gw_v1_acl_acl_proto_rawDescData = file_npool_cms_gw_v1_acl_acl_proto_rawDesc
)

func file_npool_cms_gw_v1_acl_acl_proto_rawDescGZIP() []byte {
	file_npool_cms_gw_v1_acl_acl_proto_rawDescOnce.Do(func() {
		file_npool_cms_gw_v1_acl_acl_proto_rawDescData = protoimpl.X.CompressGZIP(file_npool_cms_gw_v1_acl_acl_proto_rawDescData)
	})
	return file_npool_cms_gw_v1_acl_acl_proto_rawDescData
}

var file_npool_cms_gw_v1_acl_acl_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_npool_cms_gw_v1_acl_acl_proto_goTypes = []interface{}{
	(*ACL)(nil),               // 0: cms.gateway.acl.v1.ACL
	(*CreateACLRequest)(nil),  // 1: cms.gateway.acl.v1.CreateACLRequest
	(*CreateACLResponse)(nil), // 2: cms.gateway.acl.v1.CreateACLResponse
	(*GetACLsRequest)(nil),    // 3: cms.gateway.acl.v1.GetACLsRequest
	(*GetACLsResponse)(nil),   // 4: cms.gateway.acl.v1.GetACLsResponse
	(*DeleteACLRequest)(nil),  // 5: cms.gateway.acl.v1.DeleteACLRequest
	(*DeleteACLResponse)(nil), // 6: cms.gateway.acl.v1.DeleteACLResponse
}
var file_npool_cms_gw_v1_acl_acl_proto_depIdxs = []int32{
	0, // 0: cms.gateway.acl.v1.CreateACLResponse.Info:type_name -> cms.gateway.acl.v1.ACL
	0, // 1: cms.gateway.acl.v1.GetACLsResponse.Infos:type_name -> cms.gateway.acl.v1.ACL
	0, // 2: cms.gateway.acl.v1.DeleteACLResponse.Info:type_name -> cms.gateway.acl.v1.ACL
	1, // 3: cms.gateway.acl.v1.Gateway.CreateACL:input_type -> cms.gateway.acl.v1.CreateACLRequest
	3, // 4: cms.gateway.acl.v1.Gateway.GetACLs:input_type -> cms.gateway.acl.v1.GetACLsRequest
	5, // 5: cms.gateway.acl.v1.Gateway.DeleteACL:input_type -> cms.gateway.acl.v1.DeleteACLRequest
	2, // 6: cms.gateway.acl.v1.Gateway.CreateACL:output_type -> cms.gateway.acl.v1.CreateACLResponse
	4, // 7: cms.gateway.acl.v1.Gateway.GetACLs:output_type -> cms.gateway.acl.v1.GetACLsResponse
	6, // 8: cms.gateway.acl.v1.Gateway.DeleteACL:output_type -> cms.gateway.acl.v1.DeleteACLResponse
	6, // [6:9] is the sub-list for method output_type
	3, // [3:6] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_npool_cms_gw_v1_acl_acl_proto_init() }
func file_npool_cms_gw_v1_acl_acl_proto_init() {
	if File_npool_cms_gw_v1_acl_acl_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_npool_cms_gw_v1_acl_acl_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ACL); i {
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
		file_npool_cms_gw_v1_acl_acl_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateACLRequest); i {
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
		file_npool_cms_gw_v1_acl_acl_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateACLResponse); i {
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
		file_npool_cms_gw_v1_acl_acl_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetACLsRequest); i {
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
		file_npool_cms_gw_v1_acl_acl_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetACLsResponse); i {
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
		file_npool_cms_gw_v1_acl_acl_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteACLRequest); i {
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
		file_npool_cms_gw_v1_acl_acl_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteACLResponse); i {
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
			RawDescriptor: file_npool_cms_gw_v1_acl_acl_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_npool_cms_gw_v1_acl_acl_proto_goTypes,
		DependencyIndexes: file_npool_cms_gw_v1_acl_acl_proto_depIdxs,
		MessageInfos:      file_npool_cms_gw_v1_acl_acl_proto_msgTypes,
	}.Build()
	File_npool_cms_gw_v1_acl_acl_proto = out.File
	file_npool_cms_gw_v1_acl_acl_proto_rawDesc = nil
	file_npool_cms_gw_v1_acl_acl_proto_goTypes = nil
	file_npool_cms_gw_v1_acl_acl_proto_depIdxs = nil
}
