// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.18.1
// source: npool/inspire/gw/v1/coupon/app/scope/scope.proto

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
	GoodID             string         `protobuf:"bytes,20,opt,name=GoodID,proto3" json:"GoodID,omitempty"`
	AppID              string         `protobuf:"bytes,30,opt,name=AppID,proto3" json:"AppID,omitempty"`
	AppGoodID          string         `protobuf:"bytes,40,opt,name=AppGoodID,proto3" json:"AppGoodID,omitempty"`
	GoodName           string         `protobuf:"bytes,50,opt,name=GoodName,proto3" json:"GoodName,omitempty"`
	ScopeID            string         `protobuf:"bytes,60,opt,name=ScopeID,proto3" json:"ScopeID,omitempty"`
	CouponID           string         `protobuf:"bytes,70,opt,name=CouponID,proto3" json:"CouponID,omitempty"`
	CouponName         string         `protobuf:"bytes,80,opt,name=CouponName,proto3" json:"CouponName,omitempty"`
	CouponType         v1.CouponType  `protobuf:"varint,90,opt,name=CouponType,proto3,enum=basetypes.inspire.v1.CouponType" json:"CouponType,omitempty"`
	CouponScope        v1.CouponScope `protobuf:"varint,100,opt,name=CouponScope,proto3,enum=basetypes.inspire.v1.CouponScope" json:"CouponScope,omitempty"`
	CouponDenomination string         `protobuf:"bytes,110,opt,name=CouponDenomination,proto3" json:"CouponDenomination,omitempty"`
	CreatedAt          uint32         `protobuf:"varint,1000,opt,name=CreatedAt,proto3" json:"CreatedAt,omitempty"`
	UpdatedAt          uint32         `protobuf:"varint,1010,opt,name=UpdatedAt,proto3" json:"UpdatedAt,omitempty"`
}

func (x *Scope) Reset() {
	*x = Scope{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_inspire_gw_v1_coupon_app_scope_scope_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Scope) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Scope) ProtoMessage() {}

func (x *Scope) ProtoReflect() protoreflect.Message {
	mi := &file_npool_inspire_gw_v1_coupon_app_scope_scope_proto_msgTypes[0]
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
	return file_npool_inspire_gw_v1_coupon_app_scope_scope_proto_rawDescGZIP(), []int{0}
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

func (x *Scope) GetAppID() string {
	if x != nil {
		return x.AppID
	}
	return ""
}

func (x *Scope) GetAppGoodID() string {
	if x != nil {
		return x.AppGoodID
	}
	return ""
}

func (x *Scope) GetGoodName() string {
	if x != nil {
		return x.GoodName
	}
	return ""
}

func (x *Scope) GetScopeID() string {
	if x != nil {
		return x.ScopeID
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

type GetAppGoodScopesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AppID  string `protobuf:"bytes,10,opt,name=AppID,proto3" json:"AppID,omitempty"`
	Offset int32  `protobuf:"varint,20,opt,name=Offset,proto3" json:"Offset,omitempty"`
	Limit  int32  `protobuf:"varint,30,opt,name=Limit,proto3" json:"Limit,omitempty"`
}

func (x *GetAppGoodScopesRequest) Reset() {
	*x = GetAppGoodScopesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_inspire_gw_v1_coupon_app_scope_scope_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAppGoodScopesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAppGoodScopesRequest) ProtoMessage() {}

func (x *GetAppGoodScopesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_inspire_gw_v1_coupon_app_scope_scope_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAppGoodScopesRequest.ProtoReflect.Descriptor instead.
func (*GetAppGoodScopesRequest) Descriptor() ([]byte, []int) {
	return file_npool_inspire_gw_v1_coupon_app_scope_scope_proto_rawDescGZIP(), []int{1}
}

func (x *GetAppGoodScopesRequest) GetAppID() string {
	if x != nil {
		return x.AppID
	}
	return ""
}

func (x *GetAppGoodScopesRequest) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *GetAppGoodScopesRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type GetAppGoodScopesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Infos []*Scope `protobuf:"bytes,10,rep,name=Infos,proto3" json:"Infos,omitempty"`
	Total uint32   `protobuf:"varint,20,opt,name=Total,proto3" json:"Total,omitempty"`
}

func (x *GetAppGoodScopesResponse) Reset() {
	*x = GetAppGoodScopesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_inspire_gw_v1_coupon_app_scope_scope_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAppGoodScopesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAppGoodScopesResponse) ProtoMessage() {}

func (x *GetAppGoodScopesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_inspire_gw_v1_coupon_app_scope_scope_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAppGoodScopesResponse.ProtoReflect.Descriptor instead.
func (*GetAppGoodScopesResponse) Descriptor() ([]byte, []int) {
	return file_npool_inspire_gw_v1_coupon_app_scope_scope_proto_rawDescGZIP(), []int{2}
}

func (x *GetAppGoodScopesResponse) GetInfos() []*Scope {
	if x != nil {
		return x.Infos
	}
	return nil
}

func (x *GetAppGoodScopesResponse) GetTotal() uint32 {
	if x != nil {
		return x.Total
	}
	return 0
}

type CreateAppGoodScopeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AppID       string          `protobuf:"bytes,10,opt,name=AppID,proto3" json:"AppID,omitempty"`
	ScopeID     string          `protobuf:"bytes,20,opt,name=ScopeID,proto3" json:"ScopeID,omitempty"`
	AppGoodID   string          `protobuf:"bytes,30,opt,name=AppGoodID,proto3" json:"AppGoodID,omitempty"`
	CouponScope *v1.CouponScope `protobuf:"varint,40,opt,name=CouponScope,proto3,enum=basetypes.inspire.v1.CouponScope,oneof" json:"CouponScope,omitempty"`
}

func (x *CreateAppGoodScopeRequest) Reset() {
	*x = CreateAppGoodScopeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_inspire_gw_v1_coupon_app_scope_scope_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateAppGoodScopeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAppGoodScopeRequest) ProtoMessage() {}

func (x *CreateAppGoodScopeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_inspire_gw_v1_coupon_app_scope_scope_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAppGoodScopeRequest.ProtoReflect.Descriptor instead.
func (*CreateAppGoodScopeRequest) Descriptor() ([]byte, []int) {
	return file_npool_inspire_gw_v1_coupon_app_scope_scope_proto_rawDescGZIP(), []int{3}
}

func (x *CreateAppGoodScopeRequest) GetAppID() string {
	if x != nil {
		return x.AppID
	}
	return ""
}

func (x *CreateAppGoodScopeRequest) GetScopeID() string {
	if x != nil {
		return x.ScopeID
	}
	return ""
}

func (x *CreateAppGoodScopeRequest) GetAppGoodID() string {
	if x != nil {
		return x.AppGoodID
	}
	return ""
}

func (x *CreateAppGoodScopeRequest) GetCouponScope() v1.CouponScope {
	if x != nil && x.CouponScope != nil {
		return *x.CouponScope
	}
	return v1.CouponScope(0)
}

type CreateAppGoodScopeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *Scope `protobuf:"bytes,10,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (x *CreateAppGoodScopeResponse) Reset() {
	*x = CreateAppGoodScopeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_inspire_gw_v1_coupon_app_scope_scope_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateAppGoodScopeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAppGoodScopeResponse) ProtoMessage() {}

func (x *CreateAppGoodScopeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_inspire_gw_v1_coupon_app_scope_scope_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAppGoodScopeResponse.ProtoReflect.Descriptor instead.
func (*CreateAppGoodScopeResponse) Descriptor() ([]byte, []int) {
	return file_npool_inspire_gw_v1_coupon_app_scope_scope_proto_rawDescGZIP(), []int{4}
}

func (x *CreateAppGoodScopeResponse) GetInfo() *Scope {
	if x != nil {
		return x.Info
	}
	return nil
}

type DeleteAppGoodScopeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID    string `protobuf:"bytes,10,opt,name=ID,proto3" json:"ID,omitempty"`
	AppID string `protobuf:"bytes,20,opt,name=AppID,proto3" json:"AppID,omitempty"`
}

func (x *DeleteAppGoodScopeRequest) Reset() {
	*x = DeleteAppGoodScopeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_inspire_gw_v1_coupon_app_scope_scope_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteAppGoodScopeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteAppGoodScopeRequest) ProtoMessage() {}

func (x *DeleteAppGoodScopeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_inspire_gw_v1_coupon_app_scope_scope_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteAppGoodScopeRequest.ProtoReflect.Descriptor instead.
func (*DeleteAppGoodScopeRequest) Descriptor() ([]byte, []int) {
	return file_npool_inspire_gw_v1_coupon_app_scope_scope_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteAppGoodScopeRequest) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *DeleteAppGoodScopeRequest) GetAppID() string {
	if x != nil {
		return x.AppID
	}
	return ""
}

type DeleteAppGoodScopeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *Scope `protobuf:"bytes,10,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (x *DeleteAppGoodScopeResponse) Reset() {
	*x = DeleteAppGoodScopeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_inspire_gw_v1_coupon_app_scope_scope_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteAppGoodScopeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteAppGoodScopeResponse) ProtoMessage() {}

func (x *DeleteAppGoodScopeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_inspire_gw_v1_coupon_app_scope_scope_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteAppGoodScopeResponse.ProtoReflect.Descriptor instead.
func (*DeleteAppGoodScopeResponse) Descriptor() ([]byte, []int) {
	return file_npool_inspire_gw_v1_coupon_app_scope_scope_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteAppGoodScopeResponse) GetInfo() *Scope {
	if x != nil {
		return x.Info
	}
	return nil
}

var File_npool_inspire_gw_v1_coupon_app_scope_scope_proto protoreflect.FileDescriptor

var file_npool_inspire_gw_v1_coupon_app_scope_scope_proto_rawDesc = []byte{
	0x0a, 0x30, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x69, 0x6e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2f,
	0x67, 0x77, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x70,
	0x2f, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x2f, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x23, 0x69, 0x6e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x67, 0x61, 0x74, 0x65,
	0x77, 0x61, 0x79, 0x2e, 0x63, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x73,
	0x63, 0x6f, 0x70, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x26, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x62, 0x61, 0x73,
	0x65, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x69, 0x6e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2f, 0x76,
	0x31, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xca, 0x03,
	0x0a, 0x05, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x47, 0x6f, 0x6f, 0x64, 0x49,
	0x44, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x47, 0x6f, 0x6f, 0x64, 0x49, 0x44, 0x12,
	0x14, 0x0a, 0x05, 0x41, 0x70, 0x70, 0x49, 0x44, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x41, 0x70, 0x70, 0x49, 0x44, 0x12, 0x1c, 0x0a, 0x09, 0x41, 0x70, 0x70, 0x47, 0x6f, 0x6f, 0x64,
	0x49, 0x44, 0x18, 0x28, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x41, 0x70, 0x70, 0x47, 0x6f, 0x6f,
	0x64, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08, 0x47, 0x6f, 0x6f, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x18,
	0x32, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x47, 0x6f, 0x6f, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x49, 0x44, 0x18, 0x3c, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08, 0x43, 0x6f, 0x75,
	0x70, 0x6f, 0x6e, 0x49, 0x44, 0x18, 0x46, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x43, 0x6f, 0x75,
	0x70, 0x6f, 0x6e, 0x49, 0x44, 0x12, 0x1e, 0x0a, 0x0a, 0x43, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x4e,
	0x61, 0x6d, 0x65, 0x18, 0x50, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x43, 0x6f, 0x75, 0x70, 0x6f,
	0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x40, 0x0a, 0x0a, 0x43, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x54,
	0x79, 0x70, 0x65, 0x18, 0x5a, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x20, 0x2e, 0x62, 0x61, 0x73, 0x65,
	0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x69, 0x6e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x43, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0a, 0x43, 0x6f, 0x75,
	0x70, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x43, 0x0a, 0x0b, 0x43, 0x6f, 0x75, 0x70, 0x6f,
	0x6e, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x18, 0x64, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x21, 0x2e, 0x62,
	0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x69, 0x6e, 0x73, 0x70, 0x69, 0x72, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x52,
	0x0b, 0x43, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x12, 0x2e, 0x0a, 0x12,
	0x43, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x44, 0x65, 0x6e, 0x6f, 0x6d, 0x69, 0x6e, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x6e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x43, 0x6f, 0x75, 0x70, 0x6f, 0x6e,
	0x44, 0x65, 0x6e, 0x6f, 0x6d, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1d, 0x0a, 0x09,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0xe8, 0x07, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x09, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0xf2, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x5d, 0x0a, 0x17, 0x47, 0x65,
	0x74, 0x41, 0x70, 0x70, 0x47, 0x6f, 0x6f, 0x64, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x41, 0x70, 0x70, 0x49, 0x44, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x41, 0x70, 0x70, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x4f,
	0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x14, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x4f, 0x66, 0x66,
	0x73, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x1e, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x05, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x22, 0x72, 0x0a, 0x18, 0x47, 0x65, 0x74,
	0x41, 0x70, 0x70, 0x47, 0x6f, 0x6f, 0x64, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x40, 0x0a, 0x05, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x18, 0x0a,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x69, 0x6e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x67,
	0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x63, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x2e, 0x61, 0x70,
	0x70, 0x2e, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x63, 0x6f, 0x70, 0x65,
	0x52, 0x05, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c,
	0x18, 0x14, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x22, 0xc3, 0x01,
	0x0a, 0x19, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x70, 0x70, 0x47, 0x6f, 0x6f, 0x64, 0x53,
	0x63, 0x6f, 0x70, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x41,
	0x70, 0x70, 0x49, 0x44, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x41, 0x70, 0x70, 0x49,
	0x44, 0x12, 0x18, 0x0a, 0x07, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x49, 0x44, 0x18, 0x14, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x49, 0x44, 0x12, 0x1c, 0x0a, 0x09, 0x41,
	0x70, 0x70, 0x47, 0x6f, 0x6f, 0x64, 0x49, 0x44, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x41, 0x70, 0x70, 0x47, 0x6f, 0x6f, 0x64, 0x49, 0x44, 0x12, 0x48, 0x0a, 0x0b, 0x43, 0x6f, 0x75,
	0x70, 0x6f, 0x6e, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x18, 0x28, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x21,
	0x2e, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x69, 0x6e, 0x73, 0x70, 0x69,
	0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x53, 0x63, 0x6f, 0x70,
	0x65, 0x48, 0x00, 0x52, 0x0b, 0x43, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x53, 0x63, 0x6f, 0x70, 0x65,
	0x88, 0x01, 0x01, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x43, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x53, 0x63,
	0x6f, 0x70, 0x65, 0x22, 0x5c, 0x0a, 0x1a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x70, 0x70,
	0x47, 0x6f, 0x6f, 0x64, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x3e, 0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x2a, 0x2e, 0x69, 0x6e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61,
	0x79, 0x2e, 0x63, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x73, 0x63, 0x6f,
	0x70, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x52, 0x04, 0x49, 0x6e, 0x66,
	0x6f, 0x22, 0x41, 0x0a, 0x19, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x70, 0x70, 0x47, 0x6f,
	0x6f, 0x64, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x49, 0x44, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x12, 0x14,
	0x0a, 0x05, 0x41, 0x70, 0x70, 0x49, 0x44, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x41,
	0x70, 0x70, 0x49, 0x44, 0x22, 0x5c, 0x0a, 0x1a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x70,
	0x70, 0x47, 0x6f, 0x6f, 0x64, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x3e, 0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x2a, 0x2e, 0x69, 0x6e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77,
	0x61, 0x79, 0x2e, 0x63, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x73, 0x63,
	0x6f, 0x70, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x52, 0x04, 0x49, 0x6e,
	0x66, 0x6f, 0x32, 0xb5, 0x04, 0x0a, 0x07, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x12, 0xb1,
	0x01, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x41, 0x70, 0x70, 0x47, 0x6f, 0x6f, 0x64, 0x53, 0x63, 0x6f,
	0x70, 0x65, 0x73, 0x12, 0x3c, 0x2e, 0x69, 0x6e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x67, 0x61,
	0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x63, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x70,
	0x2e, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x70, 0x70,
	0x47, 0x6f, 0x6f, 0x64, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x3d, 0x2e, 0x69, 0x6e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x67, 0x61, 0x74, 0x65,
	0x77, 0x61, 0x79, 0x2e, 0x63, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x73,
	0x63, 0x6f, 0x70, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x70, 0x70, 0x47, 0x6f,
	0x6f, 0x64, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x20, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1a, 0x3a, 0x01, 0x2a, 0x22, 0x15, 0x2f, 0x76, 0x31,
	0x2f, 0x67, 0x65, 0x74, 0x2f, 0x61, 0x70, 0x70, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x63, 0x6f, 0x70,
	0x65, 0x73, 0x12, 0xb9, 0x01, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x70, 0x70,
	0x47, 0x6f, 0x6f, 0x64, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x12, 0x3e, 0x2e, 0x69, 0x6e, 0x73, 0x70,
	0x69, 0x72, 0x65, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x63, 0x6f, 0x75, 0x70,
	0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x70, 0x70, 0x47, 0x6f, 0x6f, 0x64, 0x53, 0x63, 0x6f,
	0x70, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3f, 0x2e, 0x69, 0x6e, 0x73, 0x70,
	0x69, 0x72, 0x65, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x63, 0x6f, 0x75, 0x70,
	0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x70, 0x70, 0x47, 0x6f, 0x6f, 0x64, 0x53, 0x63, 0x6f,
	0x70, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x22, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x1c, 0x3a, 0x01, 0x2a, 0x22, 0x17, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x2f, 0x61, 0x70, 0x70, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x12, 0xb9,
	0x01, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x70, 0x70, 0x47, 0x6f, 0x6f, 0x64,
	0x53, 0x63, 0x6f, 0x70, 0x65, 0x12, 0x3e, 0x2e, 0x69, 0x6e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e,
	0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x63, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x2e, 0x61,
	0x70, 0x70, 0x2e, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x41, 0x70, 0x70, 0x47, 0x6f, 0x6f, 0x64, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3f, 0x2e, 0x69, 0x6e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e,
	0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x63, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x2e, 0x61,
	0x70, 0x70, 0x2e, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x41, 0x70, 0x70, 0x47, 0x6f, 0x6f, 0x64, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x22, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1c, 0x3a, 0x01,
	0x2a, 0x22, 0x17, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x2f, 0x61, 0x70,
	0x70, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x42, 0x47, 0x5a, 0x45, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4e, 0x70, 0x6f, 0x6f, 0x6c, 0x50, 0x6c,
	0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x6e,
	0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x69, 0x6e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2f, 0x67, 0x77, 0x2f,
	0x76, 0x31, 0x2f, 0x63, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x70, 0x2f, 0x73, 0x63,
	0x6f, 0x70, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_npool_inspire_gw_v1_coupon_app_scope_scope_proto_rawDescOnce sync.Once
	file_npool_inspire_gw_v1_coupon_app_scope_scope_proto_rawDescData = file_npool_inspire_gw_v1_coupon_app_scope_scope_proto_rawDesc
)

func file_npool_inspire_gw_v1_coupon_app_scope_scope_proto_rawDescGZIP() []byte {
	file_npool_inspire_gw_v1_coupon_app_scope_scope_proto_rawDescOnce.Do(func() {
		file_npool_inspire_gw_v1_coupon_app_scope_scope_proto_rawDescData = protoimpl.X.CompressGZIP(file_npool_inspire_gw_v1_coupon_app_scope_scope_proto_rawDescData)
	})
	return file_npool_inspire_gw_v1_coupon_app_scope_scope_proto_rawDescData
}

var file_npool_inspire_gw_v1_coupon_app_scope_scope_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_npool_inspire_gw_v1_coupon_app_scope_scope_proto_goTypes = []interface{}{
	(*Scope)(nil),                      // 0: inspire.gateway.coupon.app.scope.v1.Scope
	(*GetAppGoodScopesRequest)(nil),    // 1: inspire.gateway.coupon.app.scope.v1.GetAppGoodScopesRequest
	(*GetAppGoodScopesResponse)(nil),   // 2: inspire.gateway.coupon.app.scope.v1.GetAppGoodScopesResponse
	(*CreateAppGoodScopeRequest)(nil),  // 3: inspire.gateway.coupon.app.scope.v1.CreateAppGoodScopeRequest
	(*CreateAppGoodScopeResponse)(nil), // 4: inspire.gateway.coupon.app.scope.v1.CreateAppGoodScopeResponse
	(*DeleteAppGoodScopeRequest)(nil),  // 5: inspire.gateway.coupon.app.scope.v1.DeleteAppGoodScopeRequest
	(*DeleteAppGoodScopeResponse)(nil), // 6: inspire.gateway.coupon.app.scope.v1.DeleteAppGoodScopeResponse
	(v1.CouponType)(0),                 // 7: basetypes.inspire.v1.CouponType
	(v1.CouponScope)(0),                // 8: basetypes.inspire.v1.CouponScope
}
var file_npool_inspire_gw_v1_coupon_app_scope_scope_proto_depIdxs = []int32{
	7, // 0: inspire.gateway.coupon.app.scope.v1.Scope.CouponType:type_name -> basetypes.inspire.v1.CouponType
	8, // 1: inspire.gateway.coupon.app.scope.v1.Scope.CouponScope:type_name -> basetypes.inspire.v1.CouponScope
	0, // 2: inspire.gateway.coupon.app.scope.v1.GetAppGoodScopesResponse.Infos:type_name -> inspire.gateway.coupon.app.scope.v1.Scope
	8, // 3: inspire.gateway.coupon.app.scope.v1.CreateAppGoodScopeRequest.CouponScope:type_name -> basetypes.inspire.v1.CouponScope
	0, // 4: inspire.gateway.coupon.app.scope.v1.CreateAppGoodScopeResponse.Info:type_name -> inspire.gateway.coupon.app.scope.v1.Scope
	0, // 5: inspire.gateway.coupon.app.scope.v1.DeleteAppGoodScopeResponse.Info:type_name -> inspire.gateway.coupon.app.scope.v1.Scope
	1, // 6: inspire.gateway.coupon.app.scope.v1.Gateway.GetAppGoodScopes:input_type -> inspire.gateway.coupon.app.scope.v1.GetAppGoodScopesRequest
	3, // 7: inspire.gateway.coupon.app.scope.v1.Gateway.CreateAppGoodScope:input_type -> inspire.gateway.coupon.app.scope.v1.CreateAppGoodScopeRequest
	5, // 8: inspire.gateway.coupon.app.scope.v1.Gateway.DeleteAppGoodScope:input_type -> inspire.gateway.coupon.app.scope.v1.DeleteAppGoodScopeRequest
	2, // 9: inspire.gateway.coupon.app.scope.v1.Gateway.GetAppGoodScopes:output_type -> inspire.gateway.coupon.app.scope.v1.GetAppGoodScopesResponse
	4, // 10: inspire.gateway.coupon.app.scope.v1.Gateway.CreateAppGoodScope:output_type -> inspire.gateway.coupon.app.scope.v1.CreateAppGoodScopeResponse
	6, // 11: inspire.gateway.coupon.app.scope.v1.Gateway.DeleteAppGoodScope:output_type -> inspire.gateway.coupon.app.scope.v1.DeleteAppGoodScopeResponse
	9, // [9:12] is the sub-list for method output_type
	6, // [6:9] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_npool_inspire_gw_v1_coupon_app_scope_scope_proto_init() }
func file_npool_inspire_gw_v1_coupon_app_scope_scope_proto_init() {
	if File_npool_inspire_gw_v1_coupon_app_scope_scope_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_npool_inspire_gw_v1_coupon_app_scope_scope_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_npool_inspire_gw_v1_coupon_app_scope_scope_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAppGoodScopesRequest); i {
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
		file_npool_inspire_gw_v1_coupon_app_scope_scope_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAppGoodScopesResponse); i {
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
		file_npool_inspire_gw_v1_coupon_app_scope_scope_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateAppGoodScopeRequest); i {
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
		file_npool_inspire_gw_v1_coupon_app_scope_scope_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateAppGoodScopeResponse); i {
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
		file_npool_inspire_gw_v1_coupon_app_scope_scope_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteAppGoodScopeRequest); i {
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
		file_npool_inspire_gw_v1_coupon_app_scope_scope_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteAppGoodScopeResponse); i {
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
	file_npool_inspire_gw_v1_coupon_app_scope_scope_proto_msgTypes[3].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_npool_inspire_gw_v1_coupon_app_scope_scope_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_npool_inspire_gw_v1_coupon_app_scope_scope_proto_goTypes,
		DependencyIndexes: file_npool_inspire_gw_v1_coupon_app_scope_scope_proto_depIdxs,
		MessageInfos:      file_npool_inspire_gw_v1_coupon_app_scope_scope_proto_msgTypes,
	}.Build()
	File_npool_inspire_gw_v1_coupon_app_scope_scope_proto = out.File
	file_npool_inspire_gw_v1_coupon_app_scope_scope_proto_rawDesc = nil
	file_npool_inspire_gw_v1_coupon_app_scope_scope_proto_goTypes = nil
	file_npool_inspire_gw_v1_coupon_app_scope_scope_proto_depIdxs = nil
}
