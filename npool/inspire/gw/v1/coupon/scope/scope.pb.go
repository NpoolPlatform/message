// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.18.1
// source: npool/inspire/gw/v1/coupon/scope/scope.proto

package scope

import (
	v1 "github.com/NpoolPlatform/message/npool/basetypes/inspire/v1"
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

type Scope struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID                 string         `protobuf:"bytes,10,opt,name=ID,proto3" json:"ID,omitempty"`
	GoodID             string         `protobuf:"bytes,30,opt,name=GoodID,proto3" json:"GoodID,omitempty"`
	GoodTitle          string         `protobuf:"bytes,40,opt,name=GoodTitle,proto3" json:"GoodTitle,omitempty"`
	CouponID           string         `protobuf:"bytes,50,opt,name=CouponID,proto3" json:"CouponID,omitempty"`
	CouponName         string         `protobuf:"bytes,60,opt,name=CouponName,proto3" json:"CouponName,omitempty"`
	CouponType         v1.CouponType  `protobuf:"varint,70,opt,name=CouponType,proto3,enum=basetypes.inspire.v1.CouponType" json:"CouponType,omitempty"`
	CouponScope        v1.CouponScope `protobuf:"varint,80,opt,name=CouponScope,proto3,enum=basetypes.inspire.v1.CouponScope" json:"CouponScope,omitempty"`
	CouponCirculation  string         `protobuf:"bytes,90,opt,name=CouponCirculation,proto3" json:"CouponCirculation,omitempty"`
	CouponDenomination string         `protobuf:"bytes,110,opt,name=CouponDenomination,proto3" json:"CouponDenomination,omitempty"`
	CreatedAt          uint32         `protobuf:"varint,1000,opt,name=CreatedAt,proto3" json:"CreatedAt,omitempty"`
	UpdatedAt          uint32         `protobuf:"varint,1010,opt,name=UpdatedAt,proto3" json:"UpdatedAt,omitempty"`
}

func (x *Scope) Reset() {
	*x = Scope{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_inspire_gw_v1_coupon_scope_scope_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Scope) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Scope) ProtoMessage() {}

func (x *Scope) ProtoReflect() protoreflect.Message {
	mi := &file_npool_inspire_gw_v1_coupon_scope_scope_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Scope.ProtoReflect.Descriptor instead.
func (*Scope) Descriptor() ([]byte, []int) {
	return file_npool_inspire_gw_v1_coupon_scope_scope_proto_rawDescGZIP(), []int{0}
}

func (x *Scope) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *Scope) GetGoodID() string {
	if x != nil {
		return x.GoodID
	}
	return ""
}

func (x *Scope) GetGoodTitle() string {
	if x != nil {
		return x.GoodTitle
	}
	return ""
}

func (x *Scope) GetCouponID() string {
	if x != nil {
		return x.CouponID
	}
	return ""
}

func (x *Scope) GetCouponName() string {
	if x != nil {
		return x.CouponName
	}
	return ""
}

func (x *Scope) GetCouponType() v1.CouponType {
	if x != nil {
		return x.CouponType
	}
	return v1.CouponType(0)
}

func (x *Scope) GetCouponScope() v1.CouponScope {
	if x != nil {
		return x.CouponScope
	}
	return v1.CouponScope(0)
}

func (x *Scope) GetCouponCirculation() string {
	if x != nil {
		return x.CouponCirculation
	}
	return ""
}

func (x *Scope) GetCouponDenomination() string {
	if x != nil {
		return x.CouponDenomination
	}
	return ""
}

func (x *Scope) GetCreatedAt() uint32 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *Scope) GetUpdatedAt() uint32 {
	if x != nil {
		return x.UpdatedAt
	}
	return 0
}

type CreateScopeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GoodID      string         `protobuf:"bytes,10,opt,name=GoodID,proto3" json:"GoodID,omitempty"`
	CouponID    string         `protobuf:"bytes,20,opt,name=CouponID,proto3" json:"CouponID,omitempty"`
	CouponScope v1.CouponScope `protobuf:"varint,30,opt,name=CouponScope,proto3,enum=basetypes.inspire.v1.CouponScope" json:"CouponScope,omitempty"`
}

func (x *CreateScopeRequest) Reset() {
	*x = CreateScopeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_inspire_gw_v1_coupon_scope_scope_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateScopeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateScopeRequest) ProtoMessage() {}

func (x *CreateScopeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_inspire_gw_v1_coupon_scope_scope_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateScopeRequest.ProtoReflect.Descriptor instead.
func (*CreateScopeRequest) Descriptor() ([]byte, []int) {
	return file_npool_inspire_gw_v1_coupon_scope_scope_proto_rawDescGZIP(), []int{1}
}

func (x *CreateScopeRequest) GetGoodID() string {
	if x != nil {
		return x.GoodID
	}
	return ""
}

func (x *CreateScopeRequest) GetCouponID() string {
	if x != nil {
		return x.CouponID
	}
	return ""
}

func (x *CreateScopeRequest) GetCouponScope() v1.CouponScope {
	if x != nil {
		return x.CouponScope
	}
	return v1.CouponScope(0)
}

type CreateScopeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *Scope `protobuf:"bytes,10,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (x *CreateScopeResponse) Reset() {
	*x = CreateScopeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_inspire_gw_v1_coupon_scope_scope_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateScopeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateScopeResponse) ProtoMessage() {}

func (x *CreateScopeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_inspire_gw_v1_coupon_scope_scope_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateScopeResponse.ProtoReflect.Descriptor instead.
func (*CreateScopeResponse) Descriptor() ([]byte, []int) {
	return file_npool_inspire_gw_v1_coupon_scope_scope_proto_rawDescGZIP(), []int{2}
}

func (x *CreateScopeResponse) GetInfo() *Scope {
	if x != nil {
		return x.Info
	}
	return nil
}

type DeleteScopeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID string `protobuf:"bytes,10,opt,name=ID,proto3" json:"ID,omitempty"`
}

func (x *DeleteScopeRequest) Reset() {
	*x = DeleteScopeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_inspire_gw_v1_coupon_scope_scope_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteScopeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteScopeRequest) ProtoMessage() {}

func (x *DeleteScopeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_inspire_gw_v1_coupon_scope_scope_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteScopeRequest.ProtoReflect.Descriptor instead.
func (*DeleteScopeRequest) Descriptor() ([]byte, []int) {
	return file_npool_inspire_gw_v1_coupon_scope_scope_proto_rawDescGZIP(), []int{3}
}

func (x *DeleteScopeRequest) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

type DeleteScopeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *Scope `protobuf:"bytes,10,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (x *DeleteScopeResponse) Reset() {
	*x = DeleteScopeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_inspire_gw_v1_coupon_scope_scope_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteScopeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteScopeResponse) ProtoMessage() {}

func (x *DeleteScopeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_inspire_gw_v1_coupon_scope_scope_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteScopeResponse.ProtoReflect.Descriptor instead.
func (*DeleteScopeResponse) Descriptor() ([]byte, []int) {
	return file_npool_inspire_gw_v1_coupon_scope_scope_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteScopeResponse) GetInfo() *Scope {
	if x != nil {
		return x.Info
	}
	return nil
}

type GetScopesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Offset int32 `protobuf:"varint,30,opt,name=Offset,proto3" json:"Offset,omitempty"`
	Limit  int32 `protobuf:"varint,40,opt,name=Limit,proto3" json:"Limit,omitempty"`
}

func (x *GetScopesRequest) Reset() {
	*x = GetScopesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_inspire_gw_v1_coupon_scope_scope_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetScopesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetScopesRequest) ProtoMessage() {}

func (x *GetScopesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_inspire_gw_v1_coupon_scope_scope_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetScopesRequest.ProtoReflect.Descriptor instead.
func (*GetScopesRequest) Descriptor() ([]byte, []int) {
	return file_npool_inspire_gw_v1_coupon_scope_scope_proto_rawDescGZIP(), []int{5}
}

func (x *GetScopesRequest) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *GetScopesRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type GetScopesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Infos []*Scope `protobuf:"bytes,10,rep,name=Infos,proto3" json:"Infos,omitempty"`
	Total uint32   `protobuf:"varint,20,opt,name=Total,proto3" json:"Total,omitempty"`
}

func (x *GetScopesResponse) Reset() {
	*x = GetScopesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_inspire_gw_v1_coupon_scope_scope_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetScopesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetScopesResponse) ProtoMessage() {}

func (x *GetScopesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_inspire_gw_v1_coupon_scope_scope_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetScopesResponse.ProtoReflect.Descriptor instead.
func (*GetScopesResponse) Descriptor() ([]byte, []int) {
	return file_npool_inspire_gw_v1_coupon_scope_scope_proto_rawDescGZIP(), []int{6}
}

func (x *GetScopesResponse) GetInfos() []*Scope {
	if x != nil {
		return x.Infos
	}
	return nil
}

func (x *GetScopesResponse) GetTotal() uint32 {
	if x != nil {
		return x.Total
	}
	return 0
}

var File_npool_inspire_gw_v1_coupon_scope_scope_proto protoreflect.FileDescriptor

var file_npool_inspire_gw_v1_coupon_scope_scope_proto_rawDesc = []byte{
	0x0a, 0x2c, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x69, 0x6e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2f,
	0x67, 0x77, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x2f, 0x73, 0x63, 0x6f,
	0x70, 0x65, 0x2f, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1f,
	0x69, 0x6e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e,
	0x63, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x2e, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x2e, 0x76, 0x31, 0x1a,
	0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x26, 0x6e,
	0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x69,
	0x6e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xac, 0x03, 0x0a, 0x05, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x12,
	0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x12,
	0x16, 0x0a, 0x06, 0x47, 0x6f, 0x6f, 0x64, 0x49, 0x44, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x47, 0x6f, 0x6f, 0x64, 0x49, 0x44, 0x12, 0x1c, 0x0a, 0x09, 0x47, 0x6f, 0x6f, 0x64, 0x54,
	0x69, 0x74, 0x6c, 0x65, 0x18, 0x28, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x47, 0x6f, 0x6f, 0x64,
	0x54, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x43, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x49,
	0x44, 0x18, 0x32, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x43, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x49,
	0x44, 0x12, 0x1e, 0x0a, 0x0a, 0x43, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x18,
	0x3c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x43, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x40, 0x0a, 0x0a, 0x43, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x18,
	0x46, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x20, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x2e, 0x69, 0x6e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x75,
	0x70, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0a, 0x43, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x43, 0x0a, 0x0b, 0x43, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x53, 0x63, 0x6f,
	0x70, 0x65, 0x18, 0x50, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x21, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x2e, 0x69, 0x6e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x43, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x52, 0x0b, 0x43, 0x6f, 0x75,
	0x70, 0x6f, 0x6e, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x12, 0x2c, 0x0a, 0x11, 0x43, 0x6f, 0x75, 0x70,
	0x6f, 0x6e, 0x43, 0x69, 0x72, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x5a, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x11, 0x43, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x43, 0x69, 0x72, 0x63, 0x75,
	0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2e, 0x0a, 0x12, 0x43, 0x6f, 0x75, 0x70, 0x6f, 0x6e,
	0x44, 0x65, 0x6e, 0x6f, 0x6d, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x6e, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x12, 0x43, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x44, 0x65, 0x6e, 0x6f, 0x6d, 0x69,
	0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1d, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x18, 0xe8, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x18, 0xf2, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x22, 0x8d, 0x01, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53,
	0x63, 0x6f, 0x70, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x47,
	0x6f, 0x6f, 0x64, 0x49, 0x44, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x47, 0x6f, 0x6f,
	0x64, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08, 0x43, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x49, 0x44, 0x18,
	0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x43, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x49, 0x44, 0x12,
	0x43, 0x0a, 0x0b, 0x43, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x18, 0x1e,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x21, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x2e, 0x69, 0x6e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x75, 0x70,
	0x6f, 0x6e, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x52, 0x0b, 0x43, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x53,
	0x63, 0x6f, 0x70, 0x65, 0x22, 0x51, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x63,
	0x6f, 0x70, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3a, 0x0a, 0x04, 0x49,
	0x6e, 0x66, 0x6f, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x69, 0x6e, 0x73, 0x70,
	0x69, 0x72, 0x65, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x63, 0x6f, 0x75, 0x70,
	0x6f, 0x6e, 0x2e, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x63, 0x6f, 0x70,
	0x65, 0x52, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x24, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x49, 0x44, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x22, 0x51, 0x0a,
	0x13, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3a, 0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x26, 0x2e, 0x69, 0x6e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x67, 0x61, 0x74,
	0x65, 0x77, 0x61, 0x79, 0x2e, 0x63, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x2e, 0x73, 0x63, 0x6f, 0x70,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x52, 0x04, 0x49, 0x6e, 0x66, 0x6f,
	0x22, 0x40, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x1e,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05,
	0x4c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x28, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x4c, 0x69, 0x6d,
	0x69, 0x74, 0x22, 0x67, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3c, 0x0a, 0x05, 0x49, 0x6e, 0x66, 0x6f, 0x73,
	0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x69, 0x6e, 0x73, 0x70, 0x69, 0x72, 0x65,
	0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x63, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x2e,
	0x73, 0x63, 0x6f, 0x70, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x52, 0x05,
	0x49, 0x6e, 0x66, 0x6f, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x14,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x32, 0xc9, 0x03, 0x0a, 0x07,
	0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x12, 0x95, 0x01, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x12, 0x33, 0x2e, 0x69, 0x6e, 0x73, 0x70, 0x69, 0x72,
	0x65, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x63, 0x6f, 0x75, 0x70, 0x6f, 0x6e,
	0x2e, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x53, 0x63, 0x6f, 0x70, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x34, 0x2e, 0x69,
	0x6e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x63,
	0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x2e, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x1b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15, 0x22, 0x10, 0x2f, 0x76, 0x31, 0x2f,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x2f, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x3a, 0x01, 0x2a, 0x12,
	0x95, 0x01, 0x0a, 0x0b, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x12,
	0x33, 0x2e, 0x69, 0x6e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61,
	0x79, 0x2e, 0x63, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x2e, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x34, 0x2e, 0x69, 0x6e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x67,
	0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x63, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x2e, 0x73, 0x63,
	0x6f, 0x70, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x63, 0x6f,
	0x70, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1b, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x15, 0x22, 0x10, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x2f, 0x73,
	0x63, 0x6f, 0x70, 0x65, 0x3a, 0x01, 0x2a, 0x12, 0x8d, 0x01, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x53,
	0x63, 0x6f, 0x70, 0x65, 0x73, 0x12, 0x31, 0x2e, 0x69, 0x6e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e,
	0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x63, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x2e, 0x73,
	0x63, 0x6f, 0x70, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x63, 0x6f, 0x70, 0x65,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x32, 0x2e, 0x69, 0x6e, 0x73, 0x70, 0x69,
	0x72, 0x65, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x63, 0x6f, 0x75, 0x70, 0x6f,
	0x6e, 0x2e, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x63,
	0x6f, 0x70, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x19, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x13, 0x22, 0x0e, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x65, 0x74, 0x2f, 0x73, 0x63,
	0x6f, 0x70, 0x65, 0x73, 0x3a, 0x01, 0x2a, 0x42, 0x43, 0x5a, 0x41, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4e, 0x70, 0x6f, 0x6f, 0x6c, 0x50, 0x6c, 0x61, 0x74, 0x66,
	0x6f, 0x72, 0x6d, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x6e, 0x70, 0x6f, 0x6f,
	0x6c, 0x2f, 0x69, 0x6e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2f, 0x67, 0x77, 0x2f, 0x76, 0x31, 0x2f,
	0x63, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x2f, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_npool_inspire_gw_v1_coupon_scope_scope_proto_rawDescOnce sync.Once
	file_npool_inspire_gw_v1_coupon_scope_scope_proto_rawDescData = file_npool_inspire_gw_v1_coupon_scope_scope_proto_rawDesc
)

func file_npool_inspire_gw_v1_coupon_scope_scope_proto_rawDescGZIP() []byte {
	file_npool_inspire_gw_v1_coupon_scope_scope_proto_rawDescOnce.Do(func() {
		file_npool_inspire_gw_v1_coupon_scope_scope_proto_rawDescData = protoimpl.X.CompressGZIP(file_npool_inspire_gw_v1_coupon_scope_scope_proto_rawDescData)
	})
	return file_npool_inspire_gw_v1_coupon_scope_scope_proto_rawDescData
}

var file_npool_inspire_gw_v1_coupon_scope_scope_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_npool_inspire_gw_v1_coupon_scope_scope_proto_goTypes = []interface{}{
	(*Scope)(nil),               // 0: inspire.gateway.coupon.scope.v1.Scope
	(*CreateScopeRequest)(nil),  // 1: inspire.gateway.coupon.scope.v1.CreateScopeRequest
	(*CreateScopeResponse)(nil), // 2: inspire.gateway.coupon.scope.v1.CreateScopeResponse
	(*DeleteScopeRequest)(nil),  // 3: inspire.gateway.coupon.scope.v1.DeleteScopeRequest
	(*DeleteScopeResponse)(nil), // 4: inspire.gateway.coupon.scope.v1.DeleteScopeResponse
	(*GetScopesRequest)(nil),    // 5: inspire.gateway.coupon.scope.v1.GetScopesRequest
	(*GetScopesResponse)(nil),   // 6: inspire.gateway.coupon.scope.v1.GetScopesResponse
	(v1.CouponType)(0),          // 7: basetypes.inspire.v1.CouponType
	(v1.CouponScope)(0),         // 8: basetypes.inspire.v1.CouponScope
}
var file_npool_inspire_gw_v1_coupon_scope_scope_proto_depIdxs = []int32{
	7, // 0: inspire.gateway.coupon.scope.v1.Scope.CouponType:type_name -> basetypes.inspire.v1.CouponType
	8, // 1: inspire.gateway.coupon.scope.v1.Scope.CouponScope:type_name -> basetypes.inspire.v1.CouponScope
	8, // 2: inspire.gateway.coupon.scope.v1.CreateScopeRequest.CouponScope:type_name -> basetypes.inspire.v1.CouponScope
	0, // 3: inspire.gateway.coupon.scope.v1.CreateScopeResponse.Info:type_name -> inspire.gateway.coupon.scope.v1.Scope
	0, // 4: inspire.gateway.coupon.scope.v1.DeleteScopeResponse.Info:type_name -> inspire.gateway.coupon.scope.v1.Scope
	0, // 5: inspire.gateway.coupon.scope.v1.GetScopesResponse.Infos:type_name -> inspire.gateway.coupon.scope.v1.Scope
	1, // 6: inspire.gateway.coupon.scope.v1.Gateway.CreateScope:input_type -> inspire.gateway.coupon.scope.v1.CreateScopeRequest
	3, // 7: inspire.gateway.coupon.scope.v1.Gateway.DeleteScope:input_type -> inspire.gateway.coupon.scope.v1.DeleteScopeRequest
	5, // 8: inspire.gateway.coupon.scope.v1.Gateway.GetScopes:input_type -> inspire.gateway.coupon.scope.v1.GetScopesRequest
	2, // 9: inspire.gateway.coupon.scope.v1.Gateway.CreateScope:output_type -> inspire.gateway.coupon.scope.v1.CreateScopeResponse
	4, // 10: inspire.gateway.coupon.scope.v1.Gateway.DeleteScope:output_type -> inspire.gateway.coupon.scope.v1.DeleteScopeResponse
	6, // 11: inspire.gateway.coupon.scope.v1.Gateway.GetScopes:output_type -> inspire.gateway.coupon.scope.v1.GetScopesResponse
	9, // [9:12] is the sub-list for method output_type
	6, // [6:9] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_npool_inspire_gw_v1_coupon_scope_scope_proto_init() }
func file_npool_inspire_gw_v1_coupon_scope_scope_proto_init() {
	if File_npool_inspire_gw_v1_coupon_scope_scope_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_npool_inspire_gw_v1_coupon_scope_scope_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Scope); i {
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
		file_npool_inspire_gw_v1_coupon_scope_scope_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateScopeRequest); i {
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
		file_npool_inspire_gw_v1_coupon_scope_scope_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateScopeResponse); i {
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
		file_npool_inspire_gw_v1_coupon_scope_scope_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteScopeRequest); i {
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
		file_npool_inspire_gw_v1_coupon_scope_scope_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteScopeResponse); i {
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
		file_npool_inspire_gw_v1_coupon_scope_scope_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetScopesRequest); i {
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
		file_npool_inspire_gw_v1_coupon_scope_scope_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetScopesResponse); i {
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
			RawDescriptor: file_npool_inspire_gw_v1_coupon_scope_scope_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_npool_inspire_gw_v1_coupon_scope_scope_proto_goTypes,
		DependencyIndexes: file_npool_inspire_gw_v1_coupon_scope_scope_proto_depIdxs,
		MessageInfos:      file_npool_inspire_gw_v1_coupon_scope_scope_proto_msgTypes,
	}.Build()
	File_npool_inspire_gw_v1_coupon_scope_scope_proto = out.File
	file_npool_inspire_gw_v1_coupon_scope_scope_proto_rawDesc = nil
	file_npool_inspire_gw_v1_coupon_scope_scope_proto_goTypes = nil
	file_npool_inspire_gw_v1_coupon_scope_scope_proto_depIdxs = nil
}
