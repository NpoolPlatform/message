// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.18.1
// source: npool/g11n/mw/v1/applang/applang.proto

package applang

import (
	_ "github.com/NpoolPlatform/message/npool"
	applang "github.com/NpoolPlatform/message/npool/g11n/mgr/v1/applang"
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

type Lang struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: sql:"id"
	ID string `protobuf:"bytes,10,opt,name=ID,proto3" json:"ID,omitempty" sql:"id"`
	// @inject_tag: sql:"app_id"
	AppID string `protobuf:"bytes,20,opt,name=AppID,proto3" json:"AppID,omitempty" sql:"app_id"`
	// @inject_tag: sql:"lang_id"
	LangID string `protobuf:"bytes,30,opt,name=LangID,proto3" json:"LangID,omitempty" sql:"lang_id"`
	// @inject_tag: sql:"lang"
	Lang string `protobuf:"bytes,40,opt,name=Lang,proto3" json:"Lang,omitempty" sql:"lang"`
	// @inject_tag: sql:"logo"
	Logo string `protobuf:"bytes,50,opt,name=Logo,proto3" json:"Logo,omitempty" sql:"logo"`
	// @inject_tag: sql:"name"
	Name string `protobuf:"bytes,60,opt,name=Name,proto3" json:"Name,omitempty" sql:"name"`
	// @inject_tag: sql:"short"
	Short string `protobuf:"bytes,70,opt,name=Short,proto3" json:"Short,omitempty" sql:"short"`
	// @inject_tag: sql:"main"
	Main bool `protobuf:"varint,80,opt,name=Main,proto3" json:"Main,omitempty" sql:"main"`
}

func (x *Lang) Reset() {
	*x = Lang{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_g11n_mw_v1_applang_applang_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Lang) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Lang) ProtoMessage() {}

func (x *Lang) ProtoReflect() protoreflect.Message {
	mi := &file_npool_g11n_mw_v1_applang_applang_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Lang.ProtoReflect.Descriptor instead.
func (*Lang) Descriptor() ([]byte, []int) {
	return file_npool_g11n_mw_v1_applang_applang_proto_rawDescGZIP(), []int{0}
}

func (x *Lang) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *Lang) GetAppID() string {
	if x != nil {
		return x.AppID
	}
	return ""
}

func (x *Lang) GetLangID() string {
	if x != nil {
		return x.LangID
	}
	return ""
}

func (x *Lang) GetLang() string {
	if x != nil {
		return x.Lang
	}
	return ""
}

func (x *Lang) GetLogo() string {
	if x != nil {
		return x.Logo
	}
	return ""
}

func (x *Lang) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Lang) GetShort() string {
	if x != nil {
		return x.Short
	}
	return ""
}

func (x *Lang) GetMain() bool {
	if x != nil {
		return x.Main
	}
	return false
}

type CreateLangRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *applang.LangReq `protobuf:"bytes,10,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (x *CreateLangRequest) Reset() {
	*x = CreateLangRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_g11n_mw_v1_applang_applang_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateLangRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateLangRequest) ProtoMessage() {}

func (x *CreateLangRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_g11n_mw_v1_applang_applang_proto_msgTypes[1]
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
	return file_npool_g11n_mw_v1_applang_applang_proto_rawDescGZIP(), []int{1}
}

func (x *CreateLangRequest) GetInfo() *applang.LangReq {
	if x != nil {
		return x.Info
	}
	return nil
}

type CreateLangResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *Lang `protobuf:"bytes,10,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (x *CreateLangResponse) Reset() {
	*x = CreateLangResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_g11n_mw_v1_applang_applang_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateLangResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateLangResponse) ProtoMessage() {}

func (x *CreateLangResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_g11n_mw_v1_applang_applang_proto_msgTypes[2]
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
	return file_npool_g11n_mw_v1_applang_applang_proto_rawDescGZIP(), []int{2}
}

func (x *CreateLangResponse) GetInfo() *Lang {
	if x != nil {
		return x.Info
	}
	return nil
}

type CreateLangsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Infos []*applang.LangReq `protobuf:"bytes,10,rep,name=Infos,proto3" json:"Infos,omitempty"`
}

func (x *CreateLangsRequest) Reset() {
	*x = CreateLangsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_g11n_mw_v1_applang_applang_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateLangsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateLangsRequest) ProtoMessage() {}

func (x *CreateLangsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_g11n_mw_v1_applang_applang_proto_msgTypes[3]
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
	return file_npool_g11n_mw_v1_applang_applang_proto_rawDescGZIP(), []int{3}
}

func (x *CreateLangsRequest) GetInfos() []*applang.LangReq {
	if x != nil {
		return x.Infos
	}
	return nil
}

type CreateLangsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Infos []*Lang `protobuf:"bytes,10,rep,name=Infos,proto3" json:"Infos,omitempty"`
}

func (x *CreateLangsResponse) Reset() {
	*x = CreateLangsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_g11n_mw_v1_applang_applang_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateLangsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateLangsResponse) ProtoMessage() {}

func (x *CreateLangsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_g11n_mw_v1_applang_applang_proto_msgTypes[4]
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
	return file_npool_g11n_mw_v1_applang_applang_proto_rawDescGZIP(), []int{4}
}

func (x *CreateLangsResponse) GetInfos() []*Lang {
	if x != nil {
		return x.Infos
	}
	return nil
}

type UpdateLangRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *applang.LangReq `protobuf:"bytes,10,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (x *UpdateLangRequest) Reset() {
	*x = UpdateLangRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_g11n_mw_v1_applang_applang_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateLangRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateLangRequest) ProtoMessage() {}

func (x *UpdateLangRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_g11n_mw_v1_applang_applang_proto_msgTypes[5]
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
	return file_npool_g11n_mw_v1_applang_applang_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateLangRequest) GetInfo() *applang.LangReq {
	if x != nil {
		return x.Info
	}
	return nil
}

type UpdateLangResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *Lang `protobuf:"bytes,10,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (x *UpdateLangResponse) Reset() {
	*x = UpdateLangResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_g11n_mw_v1_applang_applang_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateLangResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateLangResponse) ProtoMessage() {}

func (x *UpdateLangResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_g11n_mw_v1_applang_applang_proto_msgTypes[6]
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
	return file_npool_g11n_mw_v1_applang_applang_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateLangResponse) GetInfo() *Lang {
	if x != nil {
		return x.Info
	}
	return nil
}

type GetLangsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Conds  *applang.Conds `protobuf:"bytes,10,opt,name=Conds,proto3" json:"Conds,omitempty"`
	Offset int32          `protobuf:"varint,20,opt,name=Offset,proto3" json:"Offset,omitempty"`
	Limit  int32          `protobuf:"varint,30,opt,name=Limit,proto3" json:"Limit,omitempty"`
}

func (x *GetLangsRequest) Reset() {
	*x = GetLangsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_g11n_mw_v1_applang_applang_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetLangsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLangsRequest) ProtoMessage() {}

func (x *GetLangsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_g11n_mw_v1_applang_applang_proto_msgTypes[7]
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
	return file_npool_g11n_mw_v1_applang_applang_proto_rawDescGZIP(), []int{7}
}

func (x *GetLangsRequest) GetConds() *applang.Conds {
	if x != nil {
		return x.Conds
	}
	return nil
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

	Infos []*Lang `protobuf:"bytes,10,rep,name=Infos,proto3" json:"Infos,omitempty"`
	Total uint32  `protobuf:"varint,20,opt,name=Total,proto3" json:"Total,omitempty"`
}

func (x *GetLangsResponse) Reset() {
	*x = GetLangsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_g11n_mw_v1_applang_applang_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetLangsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLangsResponse) ProtoMessage() {}

func (x *GetLangsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_g11n_mw_v1_applang_applang_proto_msgTypes[8]
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
	return file_npool_g11n_mw_v1_applang_applang_proto_rawDescGZIP(), []int{8}
}

func (x *GetLangsResponse) GetInfos() []*Lang {
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

type DeleteLangRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID string `protobuf:"bytes,10,opt,name=ID,proto3" json:"ID,omitempty"`
}

func (x *DeleteLangRequest) Reset() {
	*x = DeleteLangRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_g11n_mw_v1_applang_applang_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteLangRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteLangRequest) ProtoMessage() {}

func (x *DeleteLangRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_g11n_mw_v1_applang_applang_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteLangRequest.ProtoReflect.Descriptor instead.
func (*DeleteLangRequest) Descriptor() ([]byte, []int) {
	return file_npool_g11n_mw_v1_applang_applang_proto_rawDescGZIP(), []int{9}
}

func (x *DeleteLangRequest) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

type DeleteLangResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *Lang `protobuf:"bytes,10,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (x *DeleteLangResponse) Reset() {
	*x = DeleteLangResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_g11n_mw_v1_applang_applang_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteLangResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteLangResponse) ProtoMessage() {}

func (x *DeleteLangResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_g11n_mw_v1_applang_applang_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteLangResponse.ProtoReflect.Descriptor instead.
func (*DeleteLangResponse) Descriptor() ([]byte, []int) {
	return file_npool_g11n_mw_v1_applang_applang_proto_rawDescGZIP(), []int{10}
}

func (x *DeleteLangResponse) GetInfo() *Lang {
	if x != nil {
		return x.Info
	}
	return nil
}

var File_npool_g11n_mw_v1_applang_applang_proto protoreflect.FileDescriptor

var file_npool_g11n_mw_v1_applang_applang_proto_rawDesc = []byte{
	0x0a, 0x26, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x67, 0x31, 0x31, 0x6e, 0x2f, 0x6d, 0x77, 0x2f,
	0x76, 0x31, 0x2f, 0x61, 0x70, 0x70, 0x6c, 0x61, 0x6e, 0x67, 0x2f, 0x61, 0x70, 0x70, 0x6c, 0x61,
	0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1a, 0x67, 0x31, 0x31, 0x6e, 0x2e, 0x6d,
	0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x61, 0x70, 0x70, 0x6c, 0x61, 0x6e,
	0x67, 0x2e, 0x76, 0x31, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x11, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x27, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x67, 0x31, 0x31, 0x6e,
	0x2f, 0x6d, 0x67, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x70, 0x70, 0x6c, 0x61, 0x6e, 0x67, 0x2f,
	0x61, 0x70, 0x70, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xaa, 0x01,
	0x0a, 0x04, 0x4c, 0x61, 0x6e, 0x67, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x41, 0x70, 0x70, 0x49, 0x44, 0x18,
	0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x41, 0x70, 0x70, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06,
	0x4c, 0x61, 0x6e, 0x67, 0x49, 0x44, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x4c, 0x61,
	0x6e, 0x67, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x4c, 0x61, 0x6e, 0x67, 0x18, 0x28, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x4c, 0x61, 0x6e, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x4c, 0x6f, 0x67, 0x6f,
	0x18, 0x32, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4c, 0x6f, 0x67, 0x6f, 0x12, 0x12, 0x0a, 0x04,
	0x4e, 0x61, 0x6d, 0x65, 0x18, 0x3c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x18, 0x46, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x4d, 0x61, 0x69, 0x6e, 0x18, 0x50,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x4d, 0x61, 0x69, 0x6e, 0x22, 0x49, 0x0a, 0x11, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x4c, 0x61, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x34, 0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e,
	0x67, 0x31, 0x31, 0x6e, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x70,
	0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x61, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x52,
	0x04, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x4a, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c,
	0x61, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x34, 0x0a, 0x04, 0x49,
	0x6e, 0x66, 0x6f, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x67, 0x31, 0x31, 0x6e,
	0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x61, 0x70, 0x70, 0x6c,
	0x61, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x61, 0x6e, 0x67, 0x52, 0x04, 0x49, 0x6e, 0x66,
	0x6f, 0x22, 0x4c, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x61, 0x6e, 0x67, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x36, 0x0a, 0x05, 0x49, 0x6e, 0x66, 0x6f, 0x73,
	0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x67, 0x31, 0x31, 0x6e, 0x2e, 0x6d, 0x61,
	0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x70, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x76, 0x31,
	0x2e, 0x4c, 0x61, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x52, 0x05, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x22,
	0x4d, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x61, 0x6e, 0x67, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x36, 0x0a, 0x05, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x18,
	0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x67, 0x31, 0x31, 0x6e, 0x2e, 0x6d, 0x69, 0x64,
	0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x61, 0x70, 0x70, 0x6c, 0x61, 0x6e, 0x67, 0x2e,
	0x76, 0x31, 0x2e, 0x4c, 0x61, 0x6e, 0x67, 0x52, 0x05, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x22, 0x49,
	0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4c, 0x61, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x34, 0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x20, 0x2e, 0x67, 0x31, 0x31, 0x6e, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72,
	0x2e, 0x61, 0x70, 0x70, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x61, 0x6e, 0x67,
	0x52, 0x65, 0x71, 0x52, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x4a, 0x0a, 0x12, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x4c, 0x61, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x34, 0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e,
	0x67, 0x31, 0x31, 0x6e, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e,
	0x61, 0x70, 0x70, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x61, 0x6e, 0x67, 0x52,
	0x04, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x75, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x4c, 0x61, 0x6e, 0x67,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x34, 0x0a, 0x05, 0x43, 0x6f, 0x6e, 0x64,
	0x73, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x67, 0x31, 0x31, 0x6e, 0x2e, 0x6d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x70, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x76,
	0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x64, 0x73, 0x52, 0x05, 0x43, 0x6f, 0x6e, 0x64, 0x73, 0x12, 0x16,
	0x0a, 0x06, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x14, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06,
	0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x18,
	0x1e, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x22, 0x60, 0x0a, 0x10,
	0x47, 0x65, 0x74, 0x4c, 0x61, 0x6e, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x36, 0x0a, 0x05, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x20, 0x2e, 0x67, 0x31, 0x31, 0x6e, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72,
	0x65, 0x2e, 0x61, 0x70, 0x70, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x61, 0x6e,
	0x67, 0x52, 0x05, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x6f, 0x74, 0x61,
	0x6c, 0x18, 0x14, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x22, 0x23,
	0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4c, 0x61, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x49, 0x44, 0x22, 0x4a, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4c, 0x61, 0x6e,
	0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x34, 0x0a, 0x04, 0x49, 0x6e, 0x66,
	0x6f, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x67, 0x31, 0x31, 0x6e, 0x2e, 0x6d,
	0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x61, 0x70, 0x70, 0x6c, 0x61, 0x6e,
	0x67, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x61, 0x6e, 0x67, 0x52, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x32,
	0xb4, 0x04, 0x0a, 0x0a, 0x4d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x12, 0x6d,
	0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x61, 0x6e, 0x67, 0x12, 0x2d, 0x2e, 0x67,
	0x31, 0x31, 0x6e, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x61,
	0x70, 0x70, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x4c, 0x61, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2e, 0x2e, 0x67, 0x31,
	0x31, 0x6e, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x61, 0x70,
	0x70, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c,
	0x61, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x70, 0x0a,
	0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x61, 0x6e, 0x67, 0x73, 0x12, 0x2e, 0x2e, 0x67,
	0x31, 0x31, 0x6e, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x61,
	0x70, 0x70, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x4c, 0x61, 0x6e, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2f, 0x2e, 0x67,
	0x31, 0x31, 0x6e, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x61,
	0x70, 0x70, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x4c, 0x61, 0x6e, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x6d, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4c, 0x61, 0x6e, 0x67, 0x12, 0x2d, 0x2e,
	0x67, 0x31, 0x31, 0x6e, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e,
	0x61, 0x70, 0x70, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x4c, 0x61, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2e, 0x2e, 0x67,
	0x31, 0x31, 0x6e, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x61,
	0x70, 0x70, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x4c, 0x61, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x67,
	0x0a, 0x08, 0x47, 0x65, 0x74, 0x4c, 0x61, 0x6e, 0x67, 0x73, 0x12, 0x2b, 0x2e, 0x67, 0x31, 0x31,
	0x6e, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x61, 0x70, 0x70,
	0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x61, 0x6e, 0x67, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2c, 0x2e, 0x67, 0x31, 0x31, 0x6e, 0x2e, 0x6d,
	0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x61, 0x70, 0x70, 0x6c, 0x61, 0x6e,
	0x67, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x61, 0x6e, 0x67, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x6d, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x4c, 0x61, 0x6e, 0x67, 0x12, 0x2d, 0x2e, 0x67, 0x31, 0x31, 0x6e, 0x2e, 0x6d, 0x69, 0x64,
	0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x61, 0x70, 0x70, 0x6c, 0x61, 0x6e, 0x67, 0x2e,
	0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4c, 0x61, 0x6e, 0x67, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x2e, 0x2e, 0x67, 0x31, 0x31, 0x6e, 0x2e, 0x6d, 0x69, 0x64, 0x64,
	0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x61, 0x70, 0x70, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x76,
	0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4c, 0x61, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x3b, 0x5a, 0x39, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4e, 0x70, 0x6f, 0x6f, 0x6c, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f,
	0x72, 0x6d, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x6e, 0x70, 0x6f, 0x6f, 0x6c,
	0x2f, 0x67, 0x31, 0x31, 0x6e, 0x2f, 0x6d, 0x77, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x70, 0x70, 0x6c,
	0x61, 0x6e, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_npool_g11n_mw_v1_applang_applang_proto_rawDescOnce sync.Once
	file_npool_g11n_mw_v1_applang_applang_proto_rawDescData = file_npool_g11n_mw_v1_applang_applang_proto_rawDesc
)

func file_npool_g11n_mw_v1_applang_applang_proto_rawDescGZIP() []byte {
	file_npool_g11n_mw_v1_applang_applang_proto_rawDescOnce.Do(func() {
		file_npool_g11n_mw_v1_applang_applang_proto_rawDescData = protoimpl.X.CompressGZIP(file_npool_g11n_mw_v1_applang_applang_proto_rawDescData)
	})
	return file_npool_g11n_mw_v1_applang_applang_proto_rawDescData
}

var file_npool_g11n_mw_v1_applang_applang_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_npool_g11n_mw_v1_applang_applang_proto_goTypes = []interface{}{
	(*Lang)(nil),                // 0: g11n.middleware.applang.v1.Lang
	(*CreateLangRequest)(nil),   // 1: g11n.middleware.applang.v1.CreateLangRequest
	(*CreateLangResponse)(nil),  // 2: g11n.middleware.applang.v1.CreateLangResponse
	(*CreateLangsRequest)(nil),  // 3: g11n.middleware.applang.v1.CreateLangsRequest
	(*CreateLangsResponse)(nil), // 4: g11n.middleware.applang.v1.CreateLangsResponse
	(*UpdateLangRequest)(nil),   // 5: g11n.middleware.applang.v1.UpdateLangRequest
	(*UpdateLangResponse)(nil),  // 6: g11n.middleware.applang.v1.UpdateLangResponse
	(*GetLangsRequest)(nil),     // 7: g11n.middleware.applang.v1.GetLangsRequest
	(*GetLangsResponse)(nil),    // 8: g11n.middleware.applang.v1.GetLangsResponse
	(*DeleteLangRequest)(nil),   // 9: g11n.middleware.applang.v1.DeleteLangRequest
	(*DeleteLangResponse)(nil),  // 10: g11n.middleware.applang.v1.DeleteLangResponse
	(*applang.LangReq)(nil),     // 11: g11n.manager.applang.v1.LangReq
	(*applang.Conds)(nil),       // 12: g11n.manager.applang.v1.Conds
}
var file_npool_g11n_mw_v1_applang_applang_proto_depIdxs = []int32{
	11, // 0: g11n.middleware.applang.v1.CreateLangRequest.Info:type_name -> g11n.manager.applang.v1.LangReq
	0,  // 1: g11n.middleware.applang.v1.CreateLangResponse.Info:type_name -> g11n.middleware.applang.v1.Lang
	11, // 2: g11n.middleware.applang.v1.CreateLangsRequest.Infos:type_name -> g11n.manager.applang.v1.LangReq
	0,  // 3: g11n.middleware.applang.v1.CreateLangsResponse.Infos:type_name -> g11n.middleware.applang.v1.Lang
	11, // 4: g11n.middleware.applang.v1.UpdateLangRequest.Info:type_name -> g11n.manager.applang.v1.LangReq
	0,  // 5: g11n.middleware.applang.v1.UpdateLangResponse.Info:type_name -> g11n.middleware.applang.v1.Lang
	12, // 6: g11n.middleware.applang.v1.GetLangsRequest.Conds:type_name -> g11n.manager.applang.v1.Conds
	0,  // 7: g11n.middleware.applang.v1.GetLangsResponse.Infos:type_name -> g11n.middleware.applang.v1.Lang
	0,  // 8: g11n.middleware.applang.v1.DeleteLangResponse.Info:type_name -> g11n.middleware.applang.v1.Lang
	1,  // 9: g11n.middleware.applang.v1.Middleware.CreateLang:input_type -> g11n.middleware.applang.v1.CreateLangRequest
	3,  // 10: g11n.middleware.applang.v1.Middleware.CreateLangs:input_type -> g11n.middleware.applang.v1.CreateLangsRequest
	5,  // 11: g11n.middleware.applang.v1.Middleware.UpdateLang:input_type -> g11n.middleware.applang.v1.UpdateLangRequest
	7,  // 12: g11n.middleware.applang.v1.Middleware.GetLangs:input_type -> g11n.middleware.applang.v1.GetLangsRequest
	9,  // 13: g11n.middleware.applang.v1.Middleware.DeleteLang:input_type -> g11n.middleware.applang.v1.DeleteLangRequest
	2,  // 14: g11n.middleware.applang.v1.Middleware.CreateLang:output_type -> g11n.middleware.applang.v1.CreateLangResponse
	4,  // 15: g11n.middleware.applang.v1.Middleware.CreateLangs:output_type -> g11n.middleware.applang.v1.CreateLangsResponse
	6,  // 16: g11n.middleware.applang.v1.Middleware.UpdateLang:output_type -> g11n.middleware.applang.v1.UpdateLangResponse
	8,  // 17: g11n.middleware.applang.v1.Middleware.GetLangs:output_type -> g11n.middleware.applang.v1.GetLangsResponse
	10, // 18: g11n.middleware.applang.v1.Middleware.DeleteLang:output_type -> g11n.middleware.applang.v1.DeleteLangResponse
	14, // [14:19] is the sub-list for method output_type
	9,  // [9:14] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_npool_g11n_mw_v1_applang_applang_proto_init() }
func file_npool_g11n_mw_v1_applang_applang_proto_init() {
	if File_npool_g11n_mw_v1_applang_applang_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_npool_g11n_mw_v1_applang_applang_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Lang); i {
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
		file_npool_g11n_mw_v1_applang_applang_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_npool_g11n_mw_v1_applang_applang_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
		file_npool_g11n_mw_v1_applang_applang_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
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
		file_npool_g11n_mw_v1_applang_applang_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
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
		file_npool_g11n_mw_v1_applang_applang_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
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
		file_npool_g11n_mw_v1_applang_applang_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
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
		file_npool_g11n_mw_v1_applang_applang_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
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
		file_npool_g11n_mw_v1_applang_applang_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
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
		file_npool_g11n_mw_v1_applang_applang_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteLangRequest); i {
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
		file_npool_g11n_mw_v1_applang_applang_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteLangResponse); i {
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
			RawDescriptor: file_npool_g11n_mw_v1_applang_applang_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_npool_g11n_mw_v1_applang_applang_proto_goTypes,
		DependencyIndexes: file_npool_g11n_mw_v1_applang_applang_proto_depIdxs,
		MessageInfos:      file_npool_g11n_mw_v1_applang_applang_proto_msgTypes,
	}.Build()
	File_npool_g11n_mw_v1_applang_applang_proto = out.File
	file_npool_g11n_mw_v1_applang_applang_proto_rawDesc = nil
	file_npool_g11n_mw_v1_applang_applang_proto_goTypes = nil
	file_npool_g11n_mw_v1_applang_applang_proto_depIdxs = nil
}
