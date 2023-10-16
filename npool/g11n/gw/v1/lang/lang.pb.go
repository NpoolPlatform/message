// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.18.1
// source: npool/g11n/gw/v1/lang/lang.proto

package lang

import (
	lang "github.com/NpoolPlatform/message/npool/g11n/mw/v1/lang"
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

type CreateLangRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EntID *string `protobuf:"bytes,10,opt,name=EntID,proto3,oneof" json:"EntID,omitempty"`
	Lang  string  `protobuf:"bytes,20,opt,name=Lang,proto3" json:"Lang,omitempty"`
	Logo  string  `protobuf:"bytes,30,opt,name=Logo,proto3" json:"Logo,omitempty"`
	Name  string  `protobuf:"bytes,40,opt,name=Name,proto3" json:"Name,omitempty"`
	Short string  `protobuf:"bytes,50,opt,name=Short,proto3" json:"Short,omitempty"`
}

func (x *CreateLangRequest) Reset() {
	*x = CreateLangRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_g11n_gw_v1_lang_lang_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateLangRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateLangRequest) ProtoMessage() {}

func (x *CreateLangRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_g11n_gw_v1_lang_lang_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateLangRequest.ProtoReflect.Descriptor instead.
func (*CreateLangRequest) Descriptor() ([]byte, []int) {
	return file_npool_g11n_gw_v1_lang_lang_proto_rawDescGZIP(), []int{0}
}

func (x *CreateLangRequest) GetEntID() string {
	if x != nil && x.EntID != nil {
		return *x.EntID
	}
	return 0
}

func (x *CreateLangRequest) GetLang() string {
	if x != nil {
		return x.Lang
	}
	return ""
}

func (x *CreateLangRequest) GetLogo() string {
	if x != nil {
		return x.Logo
	}
	return ""
}

func (x *CreateLangRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateLangRequest) GetShort() string {
	if x != nil {
		return x.Short
	}
	return ""
}

type CreateLangResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *lang.Lang `protobuf:"bytes,10,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (x *CreateLangResponse) Reset() {
	*x = CreateLangResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_g11n_gw_v1_lang_lang_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateLangResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateLangResponse) ProtoMessage() {}

func (x *CreateLangResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_g11n_gw_v1_lang_lang_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateLangResponse.ProtoReflect.Descriptor instead.
func (*CreateLangResponse) Descriptor() ([]byte, []int) {
	return file_npool_g11n_gw_v1_lang_lang_proto_rawDescGZIP(), []int{1}
}

func (x *CreateLangResponse) GetInfo() *lang.Lang {
	if x != nil {
		return x.Info
	}
	return nil
}

type CreateLangsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Infos []*lang.LangReq `protobuf:"bytes,10,rep,name=Infos,proto3" json:"Infos,omitempty"`
}

func (x *CreateLangsRequest) Reset() {
	*x = CreateLangsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_g11n_gw_v1_lang_lang_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateLangsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateLangsRequest) ProtoMessage() {}

func (x *CreateLangsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_g11n_gw_v1_lang_lang_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateLangsRequest.ProtoReflect.Descriptor instead.
func (*CreateLangsRequest) Descriptor() ([]byte, []int) {
	return file_npool_g11n_gw_v1_lang_lang_proto_rawDescGZIP(), []int{2}
}

func (x *CreateLangsRequest) GetInfos() []*lang.LangReq {
	if x != nil {
		return x.Infos
	}
	return nil
}

type CreateLangsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Infos []*lang.Lang `protobuf:"bytes,10,rep,name=Infos,proto3" json:"Infos,omitempty"`
}

func (x *CreateLangsResponse) Reset() {
	*x = CreateLangsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_g11n_gw_v1_lang_lang_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateLangsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateLangsResponse) ProtoMessage() {}

func (x *CreateLangsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_g11n_gw_v1_lang_lang_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateLangsResponse.ProtoReflect.Descriptor instead.
func (*CreateLangsResponse) Descriptor() ([]byte, []int) {
	return file_npool_g11n_gw_v1_lang_lang_proto_rawDescGZIP(), []int{3}
}

func (x *CreateLangsResponse) GetInfos() []*lang.Lang {
	if x != nil {
		return x.Infos
	}
	return nil
}

type UpdateLangRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID    uint32  `protobuf:"varint,10,opt,name=ID,proto3" json:"ID,omitempty"`
	Lang  *string `protobuf:"bytes,20,opt,name=Lang,proto3,oneof" json:"Lang,omitempty"`
	Logo  *string `protobuf:"bytes,30,opt,name=Logo,proto3,oneof" json:"Logo,omitempty"`
	Name  *string `protobuf:"bytes,40,opt,name=Name,proto3,oneof" json:"Name,omitempty"`
	Short *string `protobuf:"bytes,50,opt,name=Short,proto3,oneof" json:"Short,omitempty"`
}

func (x *UpdateLangRequest) Reset() {
	*x = UpdateLangRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_g11n_gw_v1_lang_lang_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateLangRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateLangRequest) ProtoMessage() {}

func (x *UpdateLangRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_g11n_gw_v1_lang_lang_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateLangRequest.ProtoReflect.Descriptor instead.
func (*UpdateLangRequest) Descriptor() ([]byte, []int) {
	return file_npool_g11n_gw_v1_lang_lang_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateLangRequest) GetID() uint32 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *UpdateLangRequest) GetLang() string {
	if x != nil && x.Lang != nil {
		return *x.Lang
	}
	return ""
}

func (x *UpdateLangRequest) GetLogo() string {
	if x != nil && x.Logo != nil {
		return *x.Logo
	}
	return ""
}

func (x *UpdateLangRequest) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *UpdateLangRequest) GetShort() string {
	if x != nil && x.Short != nil {
		return *x.Short
	}
	return ""
}

type UpdateLangResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *lang.Lang `protobuf:"bytes,10,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (x *UpdateLangResponse) Reset() {
	*x = UpdateLangResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_g11n_gw_v1_lang_lang_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateLangResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateLangResponse) ProtoMessage() {}

func (x *UpdateLangResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_g11n_gw_v1_lang_lang_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateLangResponse.ProtoReflect.Descriptor instead.
func (*UpdateLangResponse) Descriptor() ([]byte, []int) {
	return file_npool_g11n_gw_v1_lang_lang_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateLangResponse) GetInfo() *lang.Lang {
	if x != nil {
		return x.Info
	}
	return nil
}

type GetLangsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Offset int32 `protobuf:"varint,10,opt,name=Offset,proto3" json:"Offset,omitempty"`
	Limit  int32 `protobuf:"varint,20,opt,name=Limit,proto3" json:"Limit,omitempty"`
}

func (x *GetLangsRequest) Reset() {
	*x = GetLangsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_g11n_gw_v1_lang_lang_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetLangsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLangsRequest) ProtoMessage() {}

func (x *GetLangsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_g11n_gw_v1_lang_lang_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLangsRequest.ProtoReflect.Descriptor instead.
func (*GetLangsRequest) Descriptor() ([]byte, []int) {
	return file_npool_g11n_gw_v1_lang_lang_proto_rawDescGZIP(), []int{6}
}

func (x *GetLangsRequest) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *GetLangsRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type GetLangsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Infos []*lang.Lang `protobuf:"bytes,10,rep,name=Infos,proto3" json:"Infos,omitempty"`
	Total uint32       `protobuf:"varint,20,opt,name=Total,proto3" json:"Total,omitempty"`
}

func (x *GetLangsResponse) Reset() {
	*x = GetLangsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_g11n_gw_v1_lang_lang_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetLangsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLangsResponse) ProtoMessage() {}

func (x *GetLangsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_g11n_gw_v1_lang_lang_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLangsResponse.ProtoReflect.Descriptor instead.
func (*GetLangsResponse) Descriptor() ([]byte, []int) {
	return file_npool_g11n_gw_v1_lang_lang_proto_rawDescGZIP(), []int{7}
}

func (x *GetLangsResponse) GetInfos() []*lang.Lang {
	if x != nil {
		return x.Infos
	}
	return nil
}

func (x *GetLangsResponse) GetTotal() uint32 {
	if x != nil {
		return x.Total
	}
	return 0
}

var File_npool_g11n_gw_v1_lang_lang_proto protoreflect.FileDescriptor

var file_npool_g11n_gw_v1_lang_lang_proto_rawDesc = []byte{
	0x0a, 0x20, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x67, 0x31, 0x31, 0x6e, 0x2f, 0x67, 0x77, 0x2f,
	0x76, 0x31, 0x2f, 0x6c, 0x61, 0x6e, 0x67, 0x2f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x14, 0x67, 0x31, 0x31, 0x6e, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79,
	0x2e, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x67, 0x31,
	0x31, 0x6e, 0x2f, 0x6d, 0x77, 0x2f, 0x76, 0x31, 0x2f, 0x6c, 0x61, 0x6e, 0x67, 0x2f, 0x6c, 0x61,
	0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8a, 0x01, 0x0a, 0x11, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x4c, 0x61, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19,
	0x0a, 0x05, 0x45, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52,
	0x05, 0x45, 0x6e, 0x74, 0x49, 0x44, 0x88, 0x01, 0x01, 0x12, 0x12, 0x0a, 0x04, 0x4c, 0x61, 0x6e,
	0x67, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4c, 0x61, 0x6e, 0x67, 0x12, 0x12, 0x0a,
	0x04, 0x4c, 0x6f, 0x67, 0x6f, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4c, 0x6f, 0x67,
	0x6f, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x28, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x18, 0x32,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x42, 0x08, 0x0a, 0x06, 0x5f,
	0x45, 0x6e, 0x74, 0x49, 0x44, 0x22, 0x47, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c,
	0x61, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x31, 0x0a, 0x04, 0x49,
	0x6e, 0x66, 0x6f, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x67, 0x31, 0x31, 0x6e,
	0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x6c, 0x61, 0x6e, 0x67,
	0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x61, 0x6e, 0x67, 0x52, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x4c,
	0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x61, 0x6e, 0x67, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x36, 0x0a, 0x05, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x18, 0x0a, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x67, 0x31, 0x31, 0x6e, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c,
	0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x61,
	0x6e, 0x67, 0x52, 0x65, 0x71, 0x52, 0x05, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x22, 0x4a, 0x0a, 0x13,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x61, 0x6e, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x33, 0x0a, 0x05, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x18, 0x0a, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x67, 0x31, 0x31, 0x6e, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65,
	0x77, 0x61, 0x72, 0x65, 0x2e, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x61, 0x6e,
	0x67, 0x52, 0x05, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x22, 0xae, 0x01, 0x0a, 0x11, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x4c, 0x61, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x49, 0x44, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x49, 0x44, 0x12, 0x17,
	0x0a, 0x04, 0x4c, 0x61, 0x6e, 0x67, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x04,
	0x4c, 0x61, 0x6e, 0x67, 0x88, 0x01, 0x01, 0x12, 0x17, 0x0a, 0x04, 0x4c, 0x6f, 0x67, 0x6f, 0x18,
	0x1e, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x04, 0x4c, 0x6f, 0x67, 0x6f, 0x88, 0x01, 0x01,
	0x12, 0x17, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x28, 0x20, 0x01, 0x28, 0x09, 0x48, 0x02,
	0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x19, 0x0a, 0x05, 0x53, 0x68, 0x6f,
	0x72, 0x74, 0x18, 0x32, 0x20, 0x01, 0x28, 0x09, 0x48, 0x03, 0x52, 0x05, 0x53, 0x68, 0x6f, 0x72,
	0x74, 0x88, 0x01, 0x01, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x4c, 0x61, 0x6e, 0x67, 0x42, 0x07, 0x0a,
	0x05, 0x5f, 0x4c, 0x6f, 0x67, 0x6f, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x4e, 0x61, 0x6d, 0x65, 0x42,
	0x08, 0x0a, 0x06, 0x5f, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x22, 0x47, 0x0a, 0x12, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x4c, 0x61, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x31, 0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e,
	0x67, 0x31, 0x31, 0x6e, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e,
	0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x61, 0x6e, 0x67, 0x52, 0x04, 0x49, 0x6e,
	0x66, 0x6f, 0x22, 0x3f, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x4c, 0x61, 0x6e, 0x67, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x14, 0x0a,
	0x05, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x14, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x4c, 0x69,
	0x6d, 0x69, 0x74, 0x22, 0x5d, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x4c, 0x61, 0x6e, 0x67, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x33, 0x0a, 0x05, 0x49, 0x6e, 0x66, 0x6f, 0x73,
	0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x67, 0x31, 0x31, 0x6e, 0x2e, 0x6d, 0x69,
	0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x76, 0x31,
	0x2e, 0x4c, 0x61, 0x6e, 0x67, 0x52, 0x05, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x12, 0x14, 0x0a, 0x05,
	0x54, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x14, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x54, 0x6f, 0x74,
	0x61, 0x6c, 0x32, 0xf9, 0x03, 0x0a, 0x07, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x12, 0x7b,
	0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x61, 0x6e, 0x67, 0x12, 0x27, 0x2e, 0x67,
	0x31, 0x31, 0x6e, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x6c, 0x61, 0x6e, 0x67,
	0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x61, 0x6e, 0x67, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x67, 0x31, 0x31, 0x6e, 0x2e, 0x67, 0x61, 0x74,
	0x65, 0x77, 0x61, 0x79, 0x2e, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x4c, 0x61, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x1a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14, 0x3a, 0x01, 0x2a, 0x22, 0x0f, 0x2f, 0x76, 0x31, 0x2f,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x2f, 0x6c, 0x61, 0x6e, 0x67, 0x12, 0x7f, 0x0a, 0x0b, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x61, 0x6e, 0x67, 0x73, 0x12, 0x28, 0x2e, 0x67, 0x31, 0x31,
	0x6e, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x76,
	0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x61, 0x6e, 0x67, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x67, 0x31, 0x31, 0x6e, 0x2e, 0x67, 0x61, 0x74, 0x65,
	0x77, 0x61, 0x79, 0x2e, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x4c, 0x61, 0x6e, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x1b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15, 0x3a, 0x01, 0x2a, 0x22, 0x10, 0x2f, 0x76, 0x31, 0x2f,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x2f, 0x6c, 0x61, 0x6e, 0x67, 0x73, 0x12, 0x7b, 0x0a, 0x0a,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4c, 0x61, 0x6e, 0x67, 0x12, 0x27, 0x2e, 0x67, 0x31, 0x31,
	0x6e, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x76,
	0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4c, 0x61, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x67, 0x31, 0x31, 0x6e, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77,
	0x61, 0x79, 0x2e, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x4c, 0x61, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1a, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x14, 0x3a, 0x01, 0x2a, 0x22, 0x0f, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x2f, 0x6c, 0x61, 0x6e, 0x67, 0x12, 0x73, 0x0a, 0x08, 0x47, 0x65, 0x74,
	0x4c, 0x61, 0x6e, 0x67, 0x73, 0x12, 0x25, 0x2e, 0x67, 0x31, 0x31, 0x6e, 0x2e, 0x67, 0x61, 0x74,
	0x65, 0x77, 0x61, 0x79, 0x2e, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74,
	0x4c, 0x61, 0x6e, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x67,
	0x31, 0x31, 0x6e, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x6c, 0x61, 0x6e, 0x67,
	0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x61, 0x6e, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x18, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x12, 0x3a, 0x01, 0x2a, 0x22,
	0x0d, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x65, 0x74, 0x2f, 0x6c, 0x61, 0x6e, 0x67, 0x73, 0x42, 0x38,
	0x5a, 0x36, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4e, 0x70, 0x6f,
	0x6f, 0x6c, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x2f, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x67, 0x31, 0x31, 0x6e, 0x2f, 0x67, 0x77,
	0x2f, 0x76, 0x31, 0x2f, 0x6c, 0x61, 0x6e, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_npool_g11n_gw_v1_lang_lang_proto_rawDescOnce sync.Once
	file_npool_g11n_gw_v1_lang_lang_proto_rawDescData = file_npool_g11n_gw_v1_lang_lang_proto_rawDesc
)

func file_npool_g11n_gw_v1_lang_lang_proto_rawDescGZIP() []byte {
	file_npool_g11n_gw_v1_lang_lang_proto_rawDescOnce.Do(func() {
		file_npool_g11n_gw_v1_lang_lang_proto_rawDescData = protoimpl.X.CompressGZIP(file_npool_g11n_gw_v1_lang_lang_proto_rawDescData)
	})
	return file_npool_g11n_gw_v1_lang_lang_proto_rawDescData
}

var file_npool_g11n_gw_v1_lang_lang_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_npool_g11n_gw_v1_lang_lang_proto_goTypes = []interface{}{
	(*CreateLangRequest)(nil),   // 0: g11n.gateway.lang.v1.CreateLangRequest
	(*CreateLangResponse)(nil),  // 1: g11n.gateway.lang.v1.CreateLangResponse
	(*CreateLangsRequest)(nil),  // 2: g11n.gateway.lang.v1.CreateLangsRequest
	(*CreateLangsResponse)(nil), // 3: g11n.gateway.lang.v1.CreateLangsResponse
	(*UpdateLangRequest)(nil),   // 4: g11n.gateway.lang.v1.UpdateLangRequest
	(*UpdateLangResponse)(nil),  // 5: g11n.gateway.lang.v1.UpdateLangResponse
	(*GetLangsRequest)(nil),     // 6: g11n.gateway.lang.v1.GetLangsRequest
	(*GetLangsResponse)(nil),    // 7: g11n.gateway.lang.v1.GetLangsResponse
	(*lang.Lang)(nil),           // 8: g11n.middleware.lang.v1.Lang
	(*lang.LangReq)(nil),        // 9: g11n.middleware.lang.v1.LangReq
}
var file_npool_g11n_gw_v1_lang_lang_proto_depIdxs = []int32{
	8, // 0: g11n.gateway.lang.v1.CreateLangResponse.Info:type_name -> g11n.middleware.lang.v1.Lang
	9, // 1: g11n.gateway.lang.v1.CreateLangsRequest.Infos:type_name -> g11n.middleware.lang.v1.LangReq
	8, // 2: g11n.gateway.lang.v1.CreateLangsResponse.Infos:type_name -> g11n.middleware.lang.v1.Lang
	8, // 3: g11n.gateway.lang.v1.UpdateLangResponse.Info:type_name -> g11n.middleware.lang.v1.Lang
	8, // 4: g11n.gateway.lang.v1.GetLangsResponse.Infos:type_name -> g11n.middleware.lang.v1.Lang
	0, // 5: g11n.gateway.lang.v1.Gateway.CreateLang:input_type -> g11n.gateway.lang.v1.CreateLangRequest
	2, // 6: g11n.gateway.lang.v1.Gateway.CreateLangs:input_type -> g11n.gateway.lang.v1.CreateLangsRequest
	4, // 7: g11n.gateway.lang.v1.Gateway.UpdateLang:input_type -> g11n.gateway.lang.v1.UpdateLangRequest
	6, // 8: g11n.gateway.lang.v1.Gateway.GetLangs:input_type -> g11n.gateway.lang.v1.GetLangsRequest
	1, // 9: g11n.gateway.lang.v1.Gateway.CreateLang:output_type -> g11n.gateway.lang.v1.CreateLangResponse
	3, // 10: g11n.gateway.lang.v1.Gateway.CreateLangs:output_type -> g11n.gateway.lang.v1.CreateLangsResponse
	5, // 11: g11n.gateway.lang.v1.Gateway.UpdateLang:output_type -> g11n.gateway.lang.v1.UpdateLangResponse
	7, // 12: g11n.gateway.lang.v1.Gateway.GetLangs:output_type -> g11n.gateway.lang.v1.GetLangsResponse
	9, // [9:13] is the sub-list for method output_type
	5, // [5:9] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_npool_g11n_gw_v1_lang_lang_proto_init() }
func file_npool_g11n_gw_v1_lang_lang_proto_init() {
	if File_npool_g11n_gw_v1_lang_lang_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_npool_g11n_gw_v1_lang_lang_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateLangRequest); i {
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
		file_npool_g11n_gw_v1_lang_lang_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateLangResponse); i {
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
		file_npool_g11n_gw_v1_lang_lang_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateLangsRequest); i {
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
		file_npool_g11n_gw_v1_lang_lang_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateLangsResponse); i {
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
		file_npool_g11n_gw_v1_lang_lang_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateLangRequest); i {
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
		file_npool_g11n_gw_v1_lang_lang_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateLangResponse); i {
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
		file_npool_g11n_gw_v1_lang_lang_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetLangsRequest); i {
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
		file_npool_g11n_gw_v1_lang_lang_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetLangsResponse); i {
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
	file_npool_g11n_gw_v1_lang_lang_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_npool_g11n_gw_v1_lang_lang_proto_msgTypes[4].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_npool_g11n_gw_v1_lang_lang_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_npool_g11n_gw_v1_lang_lang_proto_goTypes,
		DependencyIndexes: file_npool_g11n_gw_v1_lang_lang_proto_depIdxs,
		MessageInfos:      file_npool_g11n_gw_v1_lang_lang_proto_msgTypes,
	}.Build()
	File_npool_g11n_gw_v1_lang_lang_proto = out.File
	file_npool_g11n_gw_v1_lang_lang_proto_rawDesc = nil
	file_npool_g11n_gw_v1_lang_lang_proto_goTypes = nil
	file_npool_g11n_gw_v1_lang_lang_proto_depIdxs = nil
}
