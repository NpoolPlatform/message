// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.18.1
// source: npool/chain/mw/v1/fiat/fiat.proto

package fiat

import (
	v1 "github.com/NpoolPlatform/message/npool/basetypes/v1"
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

type FiatReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID   *string `protobuf:"bytes,10,opt,name=ID,proto3,oneof" json:"ID,omitempty"`
	Name *string `protobuf:"bytes,20,opt,name=Name,proto3,oneof" json:"Name,omitempty"`
	Logo *string `protobuf:"bytes,30,opt,name=Logo,proto3,oneof" json:"Logo,omitempty"`
	Unit *string `protobuf:"bytes,40,opt,name=Unit,proto3,oneof" json:"Unit,omitempty"`
}

func (x *FiatReq) Reset() {
	*x = FiatReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_chain_mw_v1_fiat_fiat_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FiatReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FiatReq) ProtoMessage() {}

func (x *FiatReq) ProtoReflect() protoreflect.Message {
	mi := &file_npool_chain_mw_v1_fiat_fiat_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FiatReq.ProtoReflect.Descriptor instead.
func (*FiatReq) Descriptor() ([]byte, []int) {
	return file_npool_chain_mw_v1_fiat_fiat_proto_rawDescGZIP(), []int{0}
}

func (x *FiatReq) GetID() string {
	if x != nil && x.ID != nil {
		return *x.ID
	}
	return ""
}

func (x *FiatReq) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *FiatReq) GetLogo() string {
	if x != nil && x.Logo != nil {
		return *x.Logo
	}
	return ""
}

func (x *FiatReq) GetUnit() string {
	if x != nil && x.Unit != nil {
		return *x.Unit
	}
	return ""
}

type Fiat struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: sql:"id"
	ID string `protobuf:"bytes,10,opt,name=ID,proto3" json:"ID,omitempty" sql:"id"`
	// @inject_tag: sql:"name"
	Name string `protobuf:"bytes,20,opt,name=Name,proto3" json:"Name,omitempty" sql:"name"`
	// @inject_tag: sql:"logo"
	Logo string `protobuf:"bytes,30,opt,name=Logo,proto3" json:"Logo,omitempty" sql:"logo"`
	// @inject_tag: sql:"unit"
	Unit string `protobuf:"bytes,40,opt,name=Unit,proto3" json:"Unit,omitempty" sql:"unit"`
	// @inject_tag: sql:"created_at"
	CreatedAt uint32 `protobuf:"varint,1000,opt,name=CreatedAt,proto3" json:"CreatedAt,omitempty" sql:"created_at"`
	// @inject_tag: sql:"updated_at"
	UpdatedAt uint32 `protobuf:"varint,1010,opt,name=UpdatedAt,proto3" json:"UpdatedAt,omitempty" sql:"updated_at"`
}

func (x *Fiat) Reset() {
	*x = Fiat{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_chain_mw_v1_fiat_fiat_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Fiat) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Fiat) ProtoMessage() {}

func (x *Fiat) ProtoReflect() protoreflect.Message {
	mi := &file_npool_chain_mw_v1_fiat_fiat_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Fiat.ProtoReflect.Descriptor instead.
func (*Fiat) Descriptor() ([]byte, []int) {
	return file_npool_chain_mw_v1_fiat_fiat_proto_rawDescGZIP(), []int{1}
}

func (x *Fiat) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *Fiat) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Fiat) GetLogo() string {
	if x != nil {
		return x.Logo
	}
	return ""
}

func (x *Fiat) GetUnit() string {
	if x != nil {
		return x.Unit
	}
	return ""
}

func (x *Fiat) GetCreatedAt() uint32 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *Fiat) GetUpdatedAt() uint32 {
	if x != nil {
		return x.UpdatedAt
	}
	return 0
}

type Conds struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID   *v1.StringVal `protobuf:"bytes,10,opt,name=ID,proto3,oneof" json:"ID,omitempty"`
	Name *v1.StringVal `protobuf:"bytes,20,opt,name=Name,proto3,oneof" json:"Name,omitempty"`
}

func (x *Conds) Reset() {
	*x = Conds{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_chain_mw_v1_fiat_fiat_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Conds) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Conds) ProtoMessage() {}

func (x *Conds) ProtoReflect() protoreflect.Message {
	mi := &file_npool_chain_mw_v1_fiat_fiat_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Conds.ProtoReflect.Descriptor instead.
func (*Conds) Descriptor() ([]byte, []int) {
	return file_npool_chain_mw_v1_fiat_fiat_proto_rawDescGZIP(), []int{2}
}

func (x *Conds) GetID() *v1.StringVal {
	if x != nil {
		return x.ID
	}
	return nil
}

func (x *Conds) GetName() *v1.StringVal {
	if x != nil {
		return x.Name
	}
	return nil
}

type CreateFiatRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *FiatReq `protobuf:"bytes,10,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (x *CreateFiatRequest) Reset() {
	*x = CreateFiatRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_chain_mw_v1_fiat_fiat_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateFiatRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateFiatRequest) ProtoMessage() {}

func (x *CreateFiatRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_chain_mw_v1_fiat_fiat_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateFiatRequest.ProtoReflect.Descriptor instead.
func (*CreateFiatRequest) Descriptor() ([]byte, []int) {
	return file_npool_chain_mw_v1_fiat_fiat_proto_rawDescGZIP(), []int{3}
}

func (x *CreateFiatRequest) GetInfo() *FiatReq {
	if x != nil {
		return x.Info
	}
	return nil
}

type CreateFiatResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *Fiat `protobuf:"bytes,10,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (x *CreateFiatResponse) Reset() {
	*x = CreateFiatResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_chain_mw_v1_fiat_fiat_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateFiatResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateFiatResponse) ProtoMessage() {}

func (x *CreateFiatResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_chain_mw_v1_fiat_fiat_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateFiatResponse.ProtoReflect.Descriptor instead.
func (*CreateFiatResponse) Descriptor() ([]byte, []int) {
	return file_npool_chain_mw_v1_fiat_fiat_proto_rawDescGZIP(), []int{4}
}

func (x *CreateFiatResponse) GetInfo() *Fiat {
	if x != nil {
		return x.Info
	}
	return nil
}

type GetFiatRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID string `protobuf:"bytes,10,opt,name=ID,proto3" json:"ID,omitempty"`
}

func (x *GetFiatRequest) Reset() {
	*x = GetFiatRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_chain_mw_v1_fiat_fiat_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFiatRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFiatRequest) ProtoMessage() {}

func (x *GetFiatRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_chain_mw_v1_fiat_fiat_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFiatRequest.ProtoReflect.Descriptor instead.
func (*GetFiatRequest) Descriptor() ([]byte, []int) {
	return file_npool_chain_mw_v1_fiat_fiat_proto_rawDescGZIP(), []int{5}
}

func (x *GetFiatRequest) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

type GetFiatResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *Fiat `protobuf:"bytes,10,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (x *GetFiatResponse) Reset() {
	*x = GetFiatResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_chain_mw_v1_fiat_fiat_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFiatResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFiatResponse) ProtoMessage() {}

func (x *GetFiatResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_chain_mw_v1_fiat_fiat_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFiatResponse.ProtoReflect.Descriptor instead.
func (*GetFiatResponse) Descriptor() ([]byte, []int) {
	return file_npool_chain_mw_v1_fiat_fiat_proto_rawDescGZIP(), []int{6}
}

func (x *GetFiatResponse) GetInfo() *Fiat {
	if x != nil {
		return x.Info
	}
	return nil
}

type GetFiatsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Conds  *Conds `protobuf:"bytes,10,opt,name=Conds,proto3" json:"Conds,omitempty"`
	Offset int32  `protobuf:"varint,20,opt,name=Offset,proto3" json:"Offset,omitempty"`
	Limit  int32  `protobuf:"varint,30,opt,name=Limit,proto3" json:"Limit,omitempty"`
}

func (x *GetFiatsRequest) Reset() {
	*x = GetFiatsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_chain_mw_v1_fiat_fiat_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFiatsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFiatsRequest) ProtoMessage() {}

func (x *GetFiatsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_chain_mw_v1_fiat_fiat_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFiatsRequest.ProtoReflect.Descriptor instead.
func (*GetFiatsRequest) Descriptor() ([]byte, []int) {
	return file_npool_chain_mw_v1_fiat_fiat_proto_rawDescGZIP(), []int{7}
}

func (x *GetFiatsRequest) GetConds() *Conds {
	if x != nil {
		return x.Conds
	}
	return nil
}

func (x *GetFiatsRequest) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *GetFiatsRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type GetFiatsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Infos []*Fiat `protobuf:"bytes,10,rep,name=Infos,proto3" json:"Infos,omitempty"`
	Total uint32  `protobuf:"varint,20,opt,name=Total,proto3" json:"Total,omitempty"`
}

func (x *GetFiatsResponse) Reset() {
	*x = GetFiatsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_chain_mw_v1_fiat_fiat_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFiatsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFiatsResponse) ProtoMessage() {}

func (x *GetFiatsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_chain_mw_v1_fiat_fiat_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFiatsResponse.ProtoReflect.Descriptor instead.
func (*GetFiatsResponse) Descriptor() ([]byte, []int) {
	return file_npool_chain_mw_v1_fiat_fiat_proto_rawDescGZIP(), []int{8}
}

func (x *GetFiatsResponse) GetInfos() []*Fiat {
	if x != nil {
		return x.Infos
	}
	return nil
}

func (x *GetFiatsResponse) GetTotal() uint32 {
	if x != nil {
		return x.Total
	}
	return 0
}

type UpdateFiatRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *FiatReq `protobuf:"bytes,10,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (x *UpdateFiatRequest) Reset() {
	*x = UpdateFiatRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_chain_mw_v1_fiat_fiat_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateFiatRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateFiatRequest) ProtoMessage() {}

func (x *UpdateFiatRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_chain_mw_v1_fiat_fiat_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateFiatRequest.ProtoReflect.Descriptor instead.
func (*UpdateFiatRequest) Descriptor() ([]byte, []int) {
	return file_npool_chain_mw_v1_fiat_fiat_proto_rawDescGZIP(), []int{9}
}

func (x *UpdateFiatRequest) GetInfo() *FiatReq {
	if x != nil {
		return x.Info
	}
	return nil
}

type UpdateFiatResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *Fiat `protobuf:"bytes,10,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (x *UpdateFiatResponse) Reset() {
	*x = UpdateFiatResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_chain_mw_v1_fiat_fiat_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateFiatResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateFiatResponse) ProtoMessage() {}

func (x *UpdateFiatResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_chain_mw_v1_fiat_fiat_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateFiatResponse.ProtoReflect.Descriptor instead.
func (*UpdateFiatResponse) Descriptor() ([]byte, []int) {
	return file_npool_chain_mw_v1_fiat_fiat_proto_rawDescGZIP(), []int{10}
}

func (x *UpdateFiatResponse) GetInfo() *Fiat {
	if x != nil {
		return x.Info
	}
	return nil
}

var File_npool_chain_mw_v1_fiat_fiat_proto protoreflect.FileDescriptor

var file_npool_chain_mw_v1_fiat_fiat_proto_rawDesc = []byte{
	0x0a, 0x21, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2f, 0x6d, 0x77,
	0x2f, 0x76, 0x31, 0x2f, 0x66, 0x69, 0x61, 0x74, 0x2f, 0x66, 0x69, 0x61, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x18, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c,
	0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x66, 0x69, 0x61, 0x74, 0x2e, 0x76, 0x31, 0x1a, 0x1e, 0x6e,
	0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x76,
	0x31, 0x2f, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8b, 0x01,
	0x0a, 0x07, 0x46, 0x69, 0x61, 0x74, 0x52, 0x65, 0x71, 0x12, 0x13, 0x0a, 0x02, 0x49, 0x44, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x02, 0x49, 0x44, 0x88, 0x01, 0x01, 0x12, 0x17,
	0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x04,
	0x4e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x17, 0x0a, 0x04, 0x4c, 0x6f, 0x67, 0x6f, 0x18,
	0x1e, 0x20, 0x01, 0x28, 0x09, 0x48, 0x02, 0x52, 0x04, 0x4c, 0x6f, 0x67, 0x6f, 0x88, 0x01, 0x01,
	0x12, 0x17, 0x0a, 0x04, 0x55, 0x6e, 0x69, 0x74, 0x18, 0x28, 0x20, 0x01, 0x28, 0x09, 0x48, 0x03,
	0x52, 0x04, 0x55, 0x6e, 0x69, 0x74, 0x88, 0x01, 0x01, 0x42, 0x05, 0x0a, 0x03, 0x5f, 0x49, 0x44,
	0x42, 0x07, 0x0a, 0x05, 0x5f, 0x4e, 0x61, 0x6d, 0x65, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x4c, 0x6f,
	0x67, 0x6f, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x55, 0x6e, 0x69, 0x74, 0x22, 0x90, 0x01, 0x0a, 0x04,
	0x46, 0x69, 0x61, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x14, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x4c, 0x6f, 0x67, 0x6f,
	0x18, 0x1e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4c, 0x6f, 0x67, 0x6f, 0x12, 0x12, 0x0a, 0x04,
	0x55, 0x6e, 0x69, 0x74, 0x18, 0x28, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x55, 0x6e, 0x69, 0x74,
	0x12, 0x1d, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0xe8, 0x07,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12,
	0x1d, 0x0a, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0xf2, 0x07, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x77,
	0x0a, 0x05, 0x43, 0x6f, 0x6e, 0x64, 0x73, 0x12, 0x2c, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x48, 0x00, 0x52, 0x02,
	0x49, 0x44, 0x88, 0x01, 0x01, 0x12, 0x30, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x14, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x48, 0x01, 0x52, 0x04,
	0x4e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x42, 0x05, 0x0a, 0x03, 0x5f, 0x49, 0x44, 0x42, 0x07,
	0x0a, 0x05, 0x5f, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x4a, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x46, 0x69, 0x61, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x35, 0x0a, 0x04,
	0x49, 0x6e, 0x66, 0x6f, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x63, 0x68, 0x61,
	0x69, 0x6e, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x66, 0x69,
	0x61, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x69, 0x61, 0x74, 0x52, 0x65, 0x71, 0x52, 0x04, 0x49,
	0x6e, 0x66, 0x6f, 0x22, 0x48, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x46, 0x69, 0x61,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x32, 0x0a, 0x04, 0x49, 0x6e, 0x66,
	0x6f, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e,
	0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x66, 0x69, 0x61, 0x74, 0x2e,
	0x76, 0x31, 0x2e, 0x46, 0x69, 0x61, 0x74, 0x52, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x20, 0x0a,
	0x0e, 0x47, 0x65, 0x74, 0x46, 0x69, 0x61, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x22,
	0x45, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x46, 0x69, 0x61, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x32, 0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1e, 0x2e, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77,
	0x61, 0x72, 0x65, 0x2e, 0x66, 0x69, 0x61, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x69, 0x61, 0x74,
	0x52, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x76, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x46, 0x69, 0x61,
	0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x35, 0x0a, 0x05, 0x43, 0x6f, 0x6e,
	0x64, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x63, 0x68, 0x61, 0x69, 0x6e,
	0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x66, 0x69, 0x61, 0x74,
	0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x64, 0x73, 0x52, 0x05, 0x43, 0x6f, 0x6e, 0x64, 0x73,
	0x12, 0x16, 0x0a, 0x06, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x14, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x06, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x4c, 0x69, 0x6d, 0x69,
	0x74, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x22, 0x5e,
	0x0a, 0x10, 0x47, 0x65, 0x74, 0x46, 0x69, 0x61, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x34, 0x0a, 0x05, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x1e, 0x2e, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65,
	0x77, 0x61, 0x72, 0x65, 0x2e, 0x66, 0x69, 0x61, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x69, 0x61,
	0x74, 0x52, 0x05, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x6f, 0x74, 0x61,
	0x6c, 0x18, 0x14, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x22, 0x4a,
	0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x46, 0x69, 0x61, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x35, 0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x21, 0x2e, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65,
	0x77, 0x61, 0x72, 0x65, 0x2e, 0x66, 0x69, 0x61, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x69, 0x61,
	0x74, 0x52, 0x65, 0x71, 0x52, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x48, 0x0a, 0x12, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x46, 0x69, 0x61, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x32, 0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e,
	0x2e, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72,
	0x65, 0x2e, 0x66, 0x69, 0x61, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x69, 0x61, 0x74, 0x52, 0x04,
	0x49, 0x6e, 0x66, 0x6f, 0x32, 0xa9, 0x03, 0x0a, 0x0a, 0x4d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77,
	0x61, 0x72, 0x65, 0x12, 0x69, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x46, 0x69, 0x61,
	0x74, 0x12, 0x2b, 0x2e, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65,
	0x77, 0x61, 0x72, 0x65, 0x2e, 0x66, 0x69, 0x61, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x46, 0x69, 0x61, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2c,
	0x2e, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72,
	0x65, 0x2e, 0x66, 0x69, 0x61, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x46, 0x69, 0x61, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x60,
	0x0a, 0x07, 0x47, 0x65, 0x74, 0x46, 0x69, 0x61, 0x74, 0x12, 0x28, 0x2e, 0x63, 0x68, 0x61, 0x69,
	0x6e, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x66, 0x69, 0x61,
	0x74, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x69, 0x61, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x6d, 0x69, 0x64, 0x64,
	0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x66, 0x69, 0x61, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x47,
	0x65, 0x74, 0x46, 0x69, 0x61, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x63, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x46, 0x69, 0x61, 0x74, 0x73, 0x12, 0x29, 0x2e, 0x63,
	0x68, 0x61, 0x69, 0x6e, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e,
	0x66, 0x69, 0x61, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x69, 0x61, 0x74, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e,
	0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x66, 0x69, 0x61, 0x74, 0x2e,
	0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x69, 0x61, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x69, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x46,
	0x69, 0x61, 0x74, 0x12, 0x2b, 0x2e, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x6d, 0x69, 0x64, 0x64,
	0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x66, 0x69, 0x61, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x46, 0x69, 0x61, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x2c, 0x2e, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77,
	0x61, 0x72, 0x65, 0x2e, 0x66, 0x69, 0x61, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x46, 0x69, 0x61, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x42, 0x39, 0x5a, 0x37, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4e,
	0x70, 0x6f, 0x6f, 0x6c, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x2f, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x63, 0x68, 0x61, 0x69, 0x6e,
	0x2f, 0x6d, 0x77, 0x2f, 0x76, 0x31, 0x2f, 0x66, 0x69, 0x61, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_npool_chain_mw_v1_fiat_fiat_proto_rawDescOnce sync.Once
	file_npool_chain_mw_v1_fiat_fiat_proto_rawDescData = file_npool_chain_mw_v1_fiat_fiat_proto_rawDesc
)

func file_npool_chain_mw_v1_fiat_fiat_proto_rawDescGZIP() []byte {
	file_npool_chain_mw_v1_fiat_fiat_proto_rawDescOnce.Do(func() {
		file_npool_chain_mw_v1_fiat_fiat_proto_rawDescData = protoimpl.X.CompressGZIP(file_npool_chain_mw_v1_fiat_fiat_proto_rawDescData)
	})
	return file_npool_chain_mw_v1_fiat_fiat_proto_rawDescData
}

var file_npool_chain_mw_v1_fiat_fiat_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_npool_chain_mw_v1_fiat_fiat_proto_goTypes = []interface{}{
	(*FiatReq)(nil),            // 0: chain.middleware.fiat.v1.FiatReq
	(*Fiat)(nil),               // 1: chain.middleware.fiat.v1.Fiat
	(*Conds)(nil),              // 2: chain.middleware.fiat.v1.Conds
	(*CreateFiatRequest)(nil),  // 3: chain.middleware.fiat.v1.CreateFiatRequest
	(*CreateFiatResponse)(nil), // 4: chain.middleware.fiat.v1.CreateFiatResponse
	(*GetFiatRequest)(nil),     // 5: chain.middleware.fiat.v1.GetFiatRequest
	(*GetFiatResponse)(nil),    // 6: chain.middleware.fiat.v1.GetFiatResponse
	(*GetFiatsRequest)(nil),    // 7: chain.middleware.fiat.v1.GetFiatsRequest
	(*GetFiatsResponse)(nil),   // 8: chain.middleware.fiat.v1.GetFiatsResponse
	(*UpdateFiatRequest)(nil),  // 9: chain.middleware.fiat.v1.UpdateFiatRequest
	(*UpdateFiatResponse)(nil), // 10: chain.middleware.fiat.v1.UpdateFiatResponse
	(*v1.StringVal)(nil),       // 11: basetypes.v1.StringVal
}
var file_npool_chain_mw_v1_fiat_fiat_proto_depIdxs = []int32{
	11, // 0: chain.middleware.fiat.v1.Conds.ID:type_name -> basetypes.v1.StringVal
	11, // 1: chain.middleware.fiat.v1.Conds.Name:type_name -> basetypes.v1.StringVal
	0,  // 2: chain.middleware.fiat.v1.CreateFiatRequest.Info:type_name -> chain.middleware.fiat.v1.FiatReq
	1,  // 3: chain.middleware.fiat.v1.CreateFiatResponse.Info:type_name -> chain.middleware.fiat.v1.Fiat
	1,  // 4: chain.middleware.fiat.v1.GetFiatResponse.Info:type_name -> chain.middleware.fiat.v1.Fiat
	2,  // 5: chain.middleware.fiat.v1.GetFiatsRequest.Conds:type_name -> chain.middleware.fiat.v1.Conds
	1,  // 6: chain.middleware.fiat.v1.GetFiatsResponse.Infos:type_name -> chain.middleware.fiat.v1.Fiat
	0,  // 7: chain.middleware.fiat.v1.UpdateFiatRequest.Info:type_name -> chain.middleware.fiat.v1.FiatReq
	1,  // 8: chain.middleware.fiat.v1.UpdateFiatResponse.Info:type_name -> chain.middleware.fiat.v1.Fiat
	3,  // 9: chain.middleware.fiat.v1.Middleware.CreateFiat:input_type -> chain.middleware.fiat.v1.CreateFiatRequest
	5,  // 10: chain.middleware.fiat.v1.Middleware.GetFiat:input_type -> chain.middleware.fiat.v1.GetFiatRequest
	7,  // 11: chain.middleware.fiat.v1.Middleware.GetFiats:input_type -> chain.middleware.fiat.v1.GetFiatsRequest
	9,  // 12: chain.middleware.fiat.v1.Middleware.UpdateFiat:input_type -> chain.middleware.fiat.v1.UpdateFiatRequest
	4,  // 13: chain.middleware.fiat.v1.Middleware.CreateFiat:output_type -> chain.middleware.fiat.v1.CreateFiatResponse
	6,  // 14: chain.middleware.fiat.v1.Middleware.GetFiat:output_type -> chain.middleware.fiat.v1.GetFiatResponse
	8,  // 15: chain.middleware.fiat.v1.Middleware.GetFiats:output_type -> chain.middleware.fiat.v1.GetFiatsResponse
	10, // 16: chain.middleware.fiat.v1.Middleware.UpdateFiat:output_type -> chain.middleware.fiat.v1.UpdateFiatResponse
	13, // [13:17] is the sub-list for method output_type
	9,  // [9:13] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_npool_chain_mw_v1_fiat_fiat_proto_init() }
func file_npool_chain_mw_v1_fiat_fiat_proto_init() {
	if File_npool_chain_mw_v1_fiat_fiat_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_npool_chain_mw_v1_fiat_fiat_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FiatReq); i {
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
		file_npool_chain_mw_v1_fiat_fiat_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Fiat); i {
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
		file_npool_chain_mw_v1_fiat_fiat_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Conds); i {
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
		file_npool_chain_mw_v1_fiat_fiat_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateFiatRequest); i {
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
		file_npool_chain_mw_v1_fiat_fiat_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateFiatResponse); i {
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
		file_npool_chain_mw_v1_fiat_fiat_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFiatRequest); i {
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
		file_npool_chain_mw_v1_fiat_fiat_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFiatResponse); i {
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
		file_npool_chain_mw_v1_fiat_fiat_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFiatsRequest); i {
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
		file_npool_chain_mw_v1_fiat_fiat_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFiatsResponse); i {
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
		file_npool_chain_mw_v1_fiat_fiat_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateFiatRequest); i {
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
		file_npool_chain_mw_v1_fiat_fiat_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateFiatResponse); i {
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
	file_npool_chain_mw_v1_fiat_fiat_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_npool_chain_mw_v1_fiat_fiat_proto_msgTypes[2].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_npool_chain_mw_v1_fiat_fiat_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_npool_chain_mw_v1_fiat_fiat_proto_goTypes,
		DependencyIndexes: file_npool_chain_mw_v1_fiat_fiat_proto_depIdxs,
		MessageInfos:      file_npool_chain_mw_v1_fiat_fiat_proto_msgTypes,
	}.Build()
	File_npool_chain_mw_v1_fiat_fiat_proto = out.File
	file_npool_chain_mw_v1_fiat_fiat_proto_rawDesc = nil
	file_npool_chain_mw_v1_fiat_fiat_proto_goTypes = nil
	file_npool_chain_mw_v1_fiat_fiat_proto_depIdxs = nil
}
