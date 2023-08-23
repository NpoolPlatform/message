// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.18.1
// source: npool/good/gw/v1/good/required/required.proto

package required

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

type CreateRequiredRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MainGoodID     string `protobuf:"bytes,10,opt,name=MainGoodID,proto3" json:"MainGoodID,omitempty"`
	RequiredGoodID string `protobuf:"bytes,20,opt,name=RequiredGoodID,proto3" json:"RequiredGoodID,omitempty"`
	Must           bool   `protobuf:"varint,30,opt,name=Must,proto3" json:"Must,omitempty"`
}

func (x *CreateRequiredRequest) Reset() {
	*x = CreateRequiredRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_good_gw_v1_good_required_required_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateRequiredRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRequiredRequest) ProtoMessage() {}

func (x *CreateRequiredRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_good_gw_v1_good_required_required_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRequiredRequest.ProtoReflect.Descriptor instead.
func (*CreateRequiredRequest) Descriptor() ([]byte, []int) {
	return file_npool_good_gw_v1_good_required_required_proto_rawDescGZIP(), []int{0}
}

func (x *CreateRequiredRequest) GetMainGoodID() string {
	if x != nil {
		return x.MainGoodID
	}
	return ""
}

func (x *CreateRequiredRequest) GetRequiredGoodID() string {
	if x != nil {
		return x.RequiredGoodID
	}
	return ""
}

func (x *CreateRequiredRequest) GetMust() bool {
	if x != nil {
		return x.Must
	}
	return false
}

type CreateRequiredResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *Required `protobuf:"bytes,10,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (x *CreateRequiredResponse) Reset() {
	*x = CreateRequiredResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_good_gw_v1_good_required_required_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateRequiredResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRequiredResponse) ProtoMessage() {}

func (x *CreateRequiredResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_good_gw_v1_good_required_required_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRequiredResponse.ProtoReflect.Descriptor instead.
func (*CreateRequiredResponse) Descriptor() ([]byte, []int) {
	return file_npool_good_gw_v1_good_required_required_proto_rawDescGZIP(), []int{1}
}

func (x *CreateRequiredResponse) GetInfo() *Required {
	if x != nil {
		return x.Info
	}
	return nil
}

type UpdateRequiredRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID   string `protobuf:"bytes,10,opt,name=ID,proto3" json:"ID,omitempty"`
	Must string `protobuf:"bytes,20,opt,name=Must,proto3" json:"Must,omitempty"`
}

func (x *UpdateRequiredRequest) Reset() {
	*x = UpdateRequiredRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_good_gw_v1_good_required_required_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateRequiredRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateRequiredRequest) ProtoMessage() {}

func (x *UpdateRequiredRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_good_gw_v1_good_required_required_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateRequiredRequest.ProtoReflect.Descriptor instead.
func (*UpdateRequiredRequest) Descriptor() ([]byte, []int) {
	return file_npool_good_gw_v1_good_required_required_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateRequiredRequest) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *UpdateRequiredRequest) GetMust() string {
	if x != nil {
		return x.Must
	}
	return ""
}

type UpdateRequiredResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *Required `protobuf:"bytes,10,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (x *UpdateRequiredResponse) Reset() {
	*x = UpdateRequiredResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_good_gw_v1_good_required_required_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateRequiredResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateRequiredResponse) ProtoMessage() {}

func (x *UpdateRequiredResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_good_gw_v1_good_required_required_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateRequiredResponse.ProtoReflect.Descriptor instead.
func (*UpdateRequiredResponse) Descriptor() ([]byte, []int) {
	return file_npool_good_gw_v1_good_required_required_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateRequiredResponse) GetInfo() *Required {
	if x != nil {
		return x.Info
	}
	return nil
}

type GetRequiredsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GoodID *string `protobuf:"bytes,10,opt,name=GoodID,proto3,oneof" json:"GoodID,omitempty"`
	Offset int32   `protobuf:"varint,20,opt,name=Offset,proto3" json:"Offset,omitempty"`
	Limit  int32   `protobuf:"varint,30,opt,name=Limit,proto3" json:"Limit,omitempty"`
}

func (x *GetRequiredsRequest) Reset() {
	*x = GetRequiredsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_good_gw_v1_good_required_required_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRequiredsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRequiredsRequest) ProtoMessage() {}

func (x *GetRequiredsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_good_gw_v1_good_required_required_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRequiredsRequest.ProtoReflect.Descriptor instead.
func (*GetRequiredsRequest) Descriptor() ([]byte, []int) {
	return file_npool_good_gw_v1_good_required_required_proto_rawDescGZIP(), []int{4}
}

func (x *GetRequiredsRequest) GetGoodID() string {
	if x != nil && x.GoodID != nil {
		return *x.GoodID
	}
	return ""
}

func (x *GetRequiredsRequest) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *GetRequiredsRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type GetRequiredsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Infos []*Required `protobuf:"bytes,10,rep,name=Infos,proto3" json:"Infos,omitempty"`
	Total uint32      `protobuf:"varint,20,opt,name=Total,proto3" json:"Total,omitempty"`
}

func (x *GetRequiredsResponse) Reset() {
	*x = GetRequiredsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_good_gw_v1_good_required_required_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRequiredsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRequiredsResponse) ProtoMessage() {}

func (x *GetRequiredsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_good_gw_v1_good_required_required_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRequiredsResponse.ProtoReflect.Descriptor instead.
func (*GetRequiredsResponse) Descriptor() ([]byte, []int) {
	return file_npool_good_gw_v1_good_required_required_proto_rawDescGZIP(), []int{5}
}

func (x *GetRequiredsResponse) GetInfos() []*Required {
	if x != nil {
		return x.Infos
	}
	return nil
}

func (x *GetRequiredsResponse) GetTotal() uint32 {
	if x != nil {
		return x.Total
	}
	return 0
}

type DeleteRequiredRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID string `protobuf:"bytes,10,opt,name=ID,proto3" json:"ID,omitempty"`
}

func (x *DeleteRequiredRequest) Reset() {
	*x = DeleteRequiredRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_good_gw_v1_good_required_required_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteRequiredRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteRequiredRequest) ProtoMessage() {}

func (x *DeleteRequiredRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_good_gw_v1_good_required_required_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteRequiredRequest.ProtoReflect.Descriptor instead.
func (*DeleteRequiredRequest) Descriptor() ([]byte, []int) {
	return file_npool_good_gw_v1_good_required_required_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteRequiredRequest) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

type DeleteRequiredResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *Required `protobuf:"bytes,10,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (x *DeleteRequiredResponse) Reset() {
	*x = DeleteRequiredResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_good_gw_v1_good_required_required_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteRequiredResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteRequiredResponse) ProtoMessage() {}

func (x *DeleteRequiredResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_good_gw_v1_good_required_required_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteRequiredResponse.ProtoReflect.Descriptor instead.
func (*DeleteRequiredResponse) Descriptor() ([]byte, []int) {
	return file_npool_good_gw_v1_good_required_required_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteRequiredResponse) GetInfo() *Required {
	if x != nil {
		return x.Info
	}
	return nil
}

var File_npool_good_gw_v1_good_required_required_proto protoreflect.FileDescriptor

var file_npool_good_gw_v1_good_required_required_proto_rawDesc = []byte{
	0x0a, 0x2d, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x67, 0x6f, 0x6f, 0x64, 0x2f, 0x67, 0x77, 0x2f,
	0x76, 0x31, 0x2f, 0x67, 0x6f, 0x6f, 0x64, 0x2f, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64,
	0x2f, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x1f, 0x67, 0x6f, 0x6f, 0x64, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x67, 0x6f,
	0x6f, 0x64, 0x31, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x31, 0x2e, 0x76, 0x31,
	0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2d,
	0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x67, 0x6f, 0x6f, 0x64, 0x2f, 0x6d, 0x77, 0x2f, 0x76, 0x31,
	0x2f, 0x67, 0x6f, 0x6f, 0x64, 0x2f, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x2f, 0x72,
	0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x73, 0x0a,
	0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x4d, 0x61, 0x69, 0x6e, 0x47, 0x6f,
	0x6f, 0x64, 0x49, 0x44, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x4d, 0x61, 0x69, 0x6e,
	0x47, 0x6f, 0x6f, 0x64, 0x49, 0x44, 0x12, 0x26, 0x0a, 0x0e, 0x52, 0x65, 0x71, 0x75, 0x69, 0x72,
	0x65, 0x64, 0x47, 0x6f, 0x6f, 0x64, 0x49, 0x44, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e,
	0x52, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x47, 0x6f, 0x6f, 0x64, 0x49, 0x44, 0x12, 0x12,
	0x0a, 0x04, 0x4d, 0x75, 0x73, 0x74, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x4d, 0x75,
	0x73, 0x74, 0x22, 0x5a, 0x0a, 0x16, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x69, 0x72, 0x65, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x40, 0x0a, 0x04,
	0x49, 0x6e, 0x66, 0x6f, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x67, 0x6f, 0x6f,
	0x64, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x67, 0x6f, 0x6f,
	0x64, 0x31, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x31, 0x2e, 0x76, 0x31, 0x2e,
	0x52, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x52, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x3b,
	0x0a, 0x15, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x4d, 0x75, 0x73, 0x74, 0x18,
	0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4d, 0x75, 0x73, 0x74, 0x22, 0x5a, 0x0a, 0x16, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x40, 0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x67, 0x6f, 0x6f, 0x64, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c,
	0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x67, 0x6f, 0x6f, 0x64, 0x31, 0x2e, 0x72, 0x65, 0x71, 0x75,
	0x69, 0x72, 0x65, 0x64, 0x31, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65,
	0x64, 0x52, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x6b, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b,
	0x0a, 0x06, 0x47, 0x6f, 0x6f, 0x64, 0x49, 0x44, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00,
	0x52, 0x06, 0x47, 0x6f, 0x6f, 0x64, 0x49, 0x44, 0x88, 0x01, 0x01, 0x12, 0x16, 0x0a, 0x06, 0x4f,
	0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x14, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x4f, 0x66, 0x66,
	0x73, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x1e, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x05, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x47, 0x6f,
	0x6f, 0x64, 0x49, 0x44, 0x22, 0x70, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x69,
	0x72, 0x65, 0x64, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x42, 0x0a, 0x05,
	0x49, 0x6e, 0x66, 0x6f, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x67, 0x6f,
	0x6f, 0x64, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x67, 0x6f,
	0x6f, 0x64, 0x31, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x31, 0x2e, 0x76, 0x31,
	0x2e, 0x52, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x52, 0x05, 0x49, 0x6e, 0x66, 0x6f, 0x73,
	0x12, 0x14, 0x0a, 0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x14, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x22, 0x27, 0x0a, 0x15, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x22,
	0x5a, 0x0a, 0x16, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65,
	0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x40, 0x0a, 0x04, 0x49, 0x6e, 0x66,
	0x6f, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x67, 0x6f, 0x6f, 0x64, 0x2e, 0x6d,
	0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x67, 0x6f, 0x6f, 0x64, 0x31, 0x2e,
	0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x31, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x71,
	0x75, 0x69, 0x72, 0x65, 0x64, 0x52, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x32, 0xa1, 0x05, 0x0a, 0x07,
	0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x12, 0xa5, 0x01, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x12, 0x36, 0x2e, 0x67, 0x6f, 0x6f,
	0x64, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x67, 0x6f, 0x6f, 0x64, 0x31, 0x2e,
	0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x31, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x37, 0x2e, 0x67, 0x6f, 0x6f, 0x64, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61,
	0x79, 0x2e, 0x67, 0x6f, 0x6f, 0x64, 0x31, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64,
	0x31, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x69,
	0x72, 0x65, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x22, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x1c, 0x3a, 0x01, 0x2a, 0x22, 0x17, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x2f, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x67, 0x6f, 0x6f, 0x64, 0x12,
	0xa5, 0x01, 0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x69, 0x72,
	0x65, 0x64, 0x12, 0x36, 0x2e, 0x67, 0x6f, 0x6f, 0x64, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61,
	0x79, 0x2e, 0x67, 0x6f, 0x6f, 0x64, 0x31, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64,
	0x31, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x69,
	0x72, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x37, 0x2e, 0x67, 0x6f, 0x6f,
	0x64, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x67, 0x6f, 0x6f, 0x64, 0x31, 0x2e,
	0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x31, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x22, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1c, 0x3a, 0x01, 0x2a, 0x22, 0x17,
	0x2f, 0x76, 0x31, 0x2f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x72, 0x65, 0x71, 0x75, 0x69,
	0x72, 0x65, 0x64, 0x67, 0x6f, 0x6f, 0x64, 0x12, 0x9d, 0x01, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x73, 0x12, 0x34, 0x2e, 0x67, 0x6f, 0x6f, 0x64, 0x2e,
	0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x67, 0x6f, 0x6f, 0x64, 0x31, 0x2e, 0x72, 0x65,
	0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x31, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x35,
	0x2e, 0x67, 0x6f, 0x6f, 0x64, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x67, 0x6f,
	0x6f, 0x64, 0x31, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x31, 0x2e, 0x76, 0x31,
	0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x20, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1a, 0x3a, 0x01, 0x2a,
	0x22, 0x15, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x65, 0x74, 0x2f, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72,
	0x65, 0x64, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x12, 0xa5, 0x01, 0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x12, 0x36, 0x2e, 0x67, 0x6f, 0x6f,
	0x64, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x67, 0x6f, 0x6f, 0x64, 0x31, 0x2e,
	0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x31, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x37, 0x2e, 0x67, 0x6f, 0x6f, 0x64, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61,
	0x79, 0x2e, 0x67, 0x6f, 0x6f, 0x64, 0x31, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64,
	0x31, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x69,
	0x72, 0x65, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x22, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x1c, 0x3a, 0x01, 0x2a, 0x22, 0x17, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x2f, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x67, 0x6f, 0x6f, 0x64, 0x42,
	0x41, 0x5a, 0x3f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4e, 0x70,
	0x6f, 0x6f, 0x6c, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x2f, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x67, 0x6f, 0x6f, 0x64, 0x2f, 0x6d,
	0x77, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x6f, 0x6f, 0x64, 0x2f, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72,
	0x65, 0x64, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_npool_good_gw_v1_good_required_required_proto_rawDescOnce sync.Once
	file_npool_good_gw_v1_good_required_required_proto_rawDescData = file_npool_good_gw_v1_good_required_required_proto_rawDesc
)

func file_npool_good_gw_v1_good_required_required_proto_rawDescGZIP() []byte {
	file_npool_good_gw_v1_good_required_required_proto_rawDescOnce.Do(func() {
		file_npool_good_gw_v1_good_required_required_proto_rawDescData = protoimpl.X.CompressGZIP(file_npool_good_gw_v1_good_required_required_proto_rawDescData)
	})
	return file_npool_good_gw_v1_good_required_required_proto_rawDescData
}

var file_npool_good_gw_v1_good_required_required_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_npool_good_gw_v1_good_required_required_proto_goTypes = []interface{}{
	(*CreateRequiredRequest)(nil),  // 0: good.gateway.good1.required1.v1.CreateRequiredRequest
	(*CreateRequiredResponse)(nil), // 1: good.gateway.good1.required1.v1.CreateRequiredResponse
	(*UpdateRequiredRequest)(nil),  // 2: good.gateway.good1.required1.v1.UpdateRequiredRequest
	(*UpdateRequiredResponse)(nil), // 3: good.gateway.good1.required1.v1.UpdateRequiredResponse
	(*GetRequiredsRequest)(nil),    // 4: good.gateway.good1.required1.v1.GetRequiredsRequest
	(*GetRequiredsResponse)(nil),   // 5: good.gateway.good1.required1.v1.GetRequiredsResponse
	(*DeleteRequiredRequest)(nil),  // 6: good.gateway.good1.required1.v1.DeleteRequiredRequest
	(*DeleteRequiredResponse)(nil), // 7: good.gateway.good1.required1.v1.DeleteRequiredResponse
	(*Required)(nil),               // 8: good.middleware.good1.required1.v1.Required
}
var file_npool_good_gw_v1_good_required_required_proto_depIdxs = []int32{
	8, // 0: good.gateway.good1.required1.v1.CreateRequiredResponse.Info:type_name -> good.middleware.good1.required1.v1.Required
	8, // 1: good.gateway.good1.required1.v1.UpdateRequiredResponse.Info:type_name -> good.middleware.good1.required1.v1.Required
	8, // 2: good.gateway.good1.required1.v1.GetRequiredsResponse.Infos:type_name -> good.middleware.good1.required1.v1.Required
	8, // 3: good.gateway.good1.required1.v1.DeleteRequiredResponse.Info:type_name -> good.middleware.good1.required1.v1.Required
	0, // 4: good.gateway.good1.required1.v1.Gateway.CreateRequired:input_type -> good.gateway.good1.required1.v1.CreateRequiredRequest
	2, // 5: good.gateway.good1.required1.v1.Gateway.UpdateRequired:input_type -> good.gateway.good1.required1.v1.UpdateRequiredRequest
	4, // 6: good.gateway.good1.required1.v1.Gateway.GetRequireds:input_type -> good.gateway.good1.required1.v1.GetRequiredsRequest
	6, // 7: good.gateway.good1.required1.v1.Gateway.DeleteRequired:input_type -> good.gateway.good1.required1.v1.DeleteRequiredRequest
	1, // 8: good.gateway.good1.required1.v1.Gateway.CreateRequired:output_type -> good.gateway.good1.required1.v1.CreateRequiredResponse
	3, // 9: good.gateway.good1.required1.v1.Gateway.UpdateRequired:output_type -> good.gateway.good1.required1.v1.UpdateRequiredResponse
	5, // 10: good.gateway.good1.required1.v1.Gateway.GetRequireds:output_type -> good.gateway.good1.required1.v1.GetRequiredsResponse
	7, // 11: good.gateway.good1.required1.v1.Gateway.DeleteRequired:output_type -> good.gateway.good1.required1.v1.DeleteRequiredResponse
	8, // [8:12] is the sub-list for method output_type
	4, // [4:8] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_npool_good_gw_v1_good_required_required_proto_init() }
func file_npool_good_gw_v1_good_required_required_proto_init() {
	if File_npool_good_gw_v1_good_required_required_proto != nil {
		return
	}
	file_npool_good_mw_v1_good_required_required_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_npool_good_gw_v1_good_required_required_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateRequiredRequest); i {
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
		file_npool_good_gw_v1_good_required_required_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateRequiredResponse); i {
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
		file_npool_good_gw_v1_good_required_required_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateRequiredRequest); i {
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
		file_npool_good_gw_v1_good_required_required_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateRequiredResponse); i {
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
		file_npool_good_gw_v1_good_required_required_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRequiredsRequest); i {
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
		file_npool_good_gw_v1_good_required_required_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRequiredsResponse); i {
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
		file_npool_good_gw_v1_good_required_required_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteRequiredRequest); i {
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
		file_npool_good_gw_v1_good_required_required_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteRequiredResponse); i {
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
	file_npool_good_gw_v1_good_required_required_proto_msgTypes[4].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_npool_good_gw_v1_good_required_required_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_npool_good_gw_v1_good_required_required_proto_goTypes,
		DependencyIndexes: file_npool_good_gw_v1_good_required_required_proto_depIdxs,
		MessageInfos:      file_npool_good_gw_v1_good_required_required_proto_msgTypes,
	}.Build()
	File_npool_good_gw_v1_good_required_required_proto = out.File
	file_npool_good_gw_v1_good_required_required_proto_rawDesc = nil
	file_npool_good_gw_v1_good_required_required_proto_goTypes = nil
	file_npool_good_gw_v1_good_required_required_proto_depIdxs = nil
}
