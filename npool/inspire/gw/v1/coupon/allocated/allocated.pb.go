// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.18.1
// source: npool/inspire/gw/v1/coupon/allocated/allocated.proto

package allocated

import (
	allocated "github.com/NpoolPlatform/message/npool/inspire/mw/v1/coupon/allocated"
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

type CreateCouponRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AppID        string `protobuf:"bytes,10,opt,name=AppID,proto3" json:"AppID,omitempty"`
	TargetUserID string `protobuf:"bytes,20,opt,name=TargetUserID,proto3" json:"TargetUserID,omitempty"`
	CouponID     string `protobuf:"bytes,30,opt,name=CouponID,proto3" json:"CouponID,omitempty"`
}

func (x *CreateCouponRequest) Reset() {
	*x = CreateCouponRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_inspire_gw_v1_coupon_allocated_allocated_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCouponRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCouponRequest) ProtoMessage() {}

func (x *CreateCouponRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_inspire_gw_v1_coupon_allocated_allocated_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCouponRequest.ProtoReflect.Descriptor instead.
func (*CreateCouponRequest) Descriptor() ([]byte, []int) {
	return file_npool_inspire_gw_v1_coupon_allocated_allocated_proto_rawDescGZIP(), []int{0}
}

func (x *CreateCouponRequest) GetAppID() string {
	if x != nil {
		return x.AppID
	}
	return ""
}

func (x *CreateCouponRequest) GetTargetUserID() string {
	if x != nil {
		return x.TargetUserID
	}
	return ""
}

func (x *CreateCouponRequest) GetCouponID() string {
	if x != nil {
		return x.CouponID
	}
	return ""
}

type CreateCouponResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *allocated.Coupon `protobuf:"bytes,10,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (x *CreateCouponResponse) Reset() {
	*x = CreateCouponResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_inspire_gw_v1_coupon_allocated_allocated_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCouponResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCouponResponse) ProtoMessage() {}

func (x *CreateCouponResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_inspire_gw_v1_coupon_allocated_allocated_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCouponResponse.ProtoReflect.Descriptor instead.
func (*CreateCouponResponse) Descriptor() ([]byte, []int) {
	return file_npool_inspire_gw_v1_coupon_allocated_allocated_proto_rawDescGZIP(), []int{1}
}

func (x *CreateCouponResponse) GetInfo() *allocated.Coupon {
	if x != nil {
		return x.Info
	}
	return nil
}

type GetCouponsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AppID  string `protobuf:"bytes,10,opt,name=AppID,proto3" json:"AppID,omitempty"`
	UserID string `protobuf:"bytes,20,opt,name=UserID,proto3" json:"UserID,omitempty"`
	Offset int32  `protobuf:"varint,30,opt,name=Offset,proto3" json:"Offset,omitempty"`
	Limit  int32  `protobuf:"varint,40,opt,name=Limit,proto3" json:"Limit,omitempty"`
}

func (x *GetCouponsRequest) Reset() {
	*x = GetCouponsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_inspire_gw_v1_coupon_allocated_allocated_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCouponsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCouponsRequest) ProtoMessage() {}

func (x *GetCouponsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_inspire_gw_v1_coupon_allocated_allocated_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCouponsRequest.ProtoReflect.Descriptor instead.
func (*GetCouponsRequest) Descriptor() ([]byte, []int) {
	return file_npool_inspire_gw_v1_coupon_allocated_allocated_proto_rawDescGZIP(), []int{2}
}

func (x *GetCouponsRequest) GetAppID() string {
	if x != nil {
		return x.AppID
	}
	return ""
}

func (x *GetCouponsRequest) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

func (x *GetCouponsRequest) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *GetCouponsRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type GetCouponsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Infos []*allocated.Coupon `protobuf:"bytes,10,rep,name=Infos,proto3" json:"Infos,omitempty"`
	Total uint32              `protobuf:"varint,20,opt,name=Total,proto3" json:"Total,omitempty"`
}

func (x *GetCouponsResponse) Reset() {
	*x = GetCouponsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_inspire_gw_v1_coupon_allocated_allocated_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCouponsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCouponsResponse) ProtoMessage() {}

func (x *GetCouponsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_inspire_gw_v1_coupon_allocated_allocated_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCouponsResponse.ProtoReflect.Descriptor instead.
func (*GetCouponsResponse) Descriptor() ([]byte, []int) {
	return file_npool_inspire_gw_v1_coupon_allocated_allocated_proto_rawDescGZIP(), []int{3}
}

func (x *GetCouponsResponse) GetInfos() []*allocated.Coupon {
	if x != nil {
		return x.Infos
	}
	return nil
}

func (x *GetCouponsResponse) GetTotal() uint32 {
	if x != nil {
		return x.Total
	}
	return 0
}

type GetAppCouponsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AppID  string `protobuf:"bytes,10,opt,name=AppID,proto3" json:"AppID,omitempty"`
	Offset int32  `protobuf:"varint,20,opt,name=Offset,proto3" json:"Offset,omitempty"`
	Limit  int32  `protobuf:"varint,30,opt,name=Limit,proto3" json:"Limit,omitempty"`
}

func (x *GetAppCouponsRequest) Reset() {
	*x = GetAppCouponsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_inspire_gw_v1_coupon_allocated_allocated_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAppCouponsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAppCouponsRequest) ProtoMessage() {}

func (x *GetAppCouponsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_inspire_gw_v1_coupon_allocated_allocated_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAppCouponsRequest.ProtoReflect.Descriptor instead.
func (*GetAppCouponsRequest) Descriptor() ([]byte, []int) {
	return file_npool_inspire_gw_v1_coupon_allocated_allocated_proto_rawDescGZIP(), []int{4}
}

func (x *GetAppCouponsRequest) GetAppID() string {
	if x != nil {
		return x.AppID
	}
	return ""
}

func (x *GetAppCouponsRequest) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *GetAppCouponsRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type GetAppCouponsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Infos []*allocated.Coupon `protobuf:"bytes,10,rep,name=Infos,proto3" json:"Infos,omitempty"`
	Total uint32              `protobuf:"varint,20,opt,name=Total,proto3" json:"Total,omitempty"`
}

func (x *GetAppCouponsResponse) Reset() {
	*x = GetAppCouponsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_inspire_gw_v1_coupon_allocated_allocated_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAppCouponsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAppCouponsResponse) ProtoMessage() {}

func (x *GetAppCouponsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_inspire_gw_v1_coupon_allocated_allocated_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAppCouponsResponse.ProtoReflect.Descriptor instead.
func (*GetAppCouponsResponse) Descriptor() ([]byte, []int) {
	return file_npool_inspire_gw_v1_coupon_allocated_allocated_proto_rawDescGZIP(), []int{5}
}

func (x *GetAppCouponsResponse) GetInfos() []*allocated.Coupon {
	if x != nil {
		return x.Infos
	}
	return nil
}

func (x *GetAppCouponsResponse) GetTotal() uint32 {
	if x != nil {
		return x.Total
	}
	return 0
}

var File_npool_inspire_gw_v1_coupon_allocated_allocated_proto protoreflect.FileDescriptor

var file_npool_inspire_gw_v1_coupon_allocated_allocated_proto_rawDesc = []byte{
	0x0a, 0x34, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x69, 0x6e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2f,
	0x67, 0x77, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x2f, 0x61, 0x6c, 0x6c,
	0x6f, 0x63, 0x61, 0x74, 0x65, 0x64, 0x2f, 0x61, 0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x65, 0x64,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x23, 0x69, 0x6e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e,
	0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x63, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x2e, 0x61,
	0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x65, 0x64, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x34, 0x6e, 0x70, 0x6f, 0x6f, 0x6c,
	0x2f, 0x69, 0x6e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2f, 0x6d, 0x77, 0x2f, 0x76, 0x31, 0x2f, 0x63,
	0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x2f, 0x61, 0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x65, 0x64, 0x2f,
	0x61, 0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x6b, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x41, 0x70, 0x70, 0x49, 0x44, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x41, 0x70, 0x70, 0x49, 0x44, 0x12, 0x22, 0x0a, 0x0c,
	0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x14, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0c, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44,
	0x12, 0x1a, 0x0a, 0x08, 0x43, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x49, 0x44, 0x18, 0x1e, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x43, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x49, 0x44, 0x22, 0x5a, 0x0a, 0x14,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x42, 0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x69, 0x6e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x6d, 0x69, 0x64,
	0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x63, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x2e, 0x61,
	0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x65, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x75, 0x70,
	0x6f, 0x6e, 0x52, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x6f, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x43,
	0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a,
	0x05, 0x41, 0x70, 0x70, 0x49, 0x44, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x41, 0x70,
	0x70, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x14, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x4f,
	0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x4f, 0x66, 0x66,
	0x73, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x28, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x05, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x22, 0x70, 0x0a, 0x12, 0x47, 0x65, 0x74,
	0x43, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x44, 0x0a, 0x05, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2e,
	0x2e, 0x69, 0x6e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77,
	0x61, 0x72, 0x65, 0x2e, 0x63, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x2e, 0x61, 0x6c, 0x6c, 0x6f, 0x63,
	0x61, 0x74, 0x65, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x52, 0x05,
	0x49, 0x6e, 0x66, 0x6f, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x14,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x22, 0x5a, 0x0a, 0x14, 0x47,
	0x65, 0x74, 0x41, 0x70, 0x70, 0x43, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x41, 0x70, 0x70, 0x49, 0x44, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x41, 0x70, 0x70, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x4f, 0x66, 0x66,
	0x73, 0x65, 0x74, 0x18, 0x14, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x4f, 0x66, 0x66, 0x73, 0x65,
	0x74, 0x12, 0x14, 0x0a, 0x05, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x05, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x22, 0x73, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x41, 0x70,
	0x70, 0x43, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x44, 0x0a, 0x05, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x2e, 0x2e, 0x69, 0x6e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65,
	0x77, 0x61, 0x72, 0x65, 0x2e, 0x63, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x2e, 0x61, 0x6c, 0x6c, 0x6f,
	0x63, 0x61, 0x74, 0x65, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x52,
	0x05, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x18,
	0x14, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x32, 0x8d, 0x04, 0x0a,
	0x07, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x12, 0xaa, 0x01, 0x0a, 0x0c, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x43, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x12, 0x38, 0x2e, 0x69, 0x6e, 0x73, 0x70,
	0x69, 0x72, 0x65, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x63, 0x6f, 0x75, 0x70,
	0x6f, 0x6e, 0x2e, 0x61, 0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x65, 0x64, 0x2e, 0x76, 0x31, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x39, 0x2e, 0x69, 0x6e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x67, 0x61,
	0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x63, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x2e, 0x61, 0x6c, 0x6c,
	0x6f, 0x63, 0x61, 0x74, 0x65, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x43, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x25,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1f, 0x3a, 0x01, 0x2a, 0x22, 0x1a, 0x2f, 0x76, 0x31, 0x2f, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x2f, 0x61, 0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x65, 0x64, 0x63,
	0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x12, 0xa2, 0x01, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x75,
	0x70, 0x6f, 0x6e, 0x73, 0x12, 0x36, 0x2e, 0x69, 0x6e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x67,
	0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x63, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x2e, 0x61, 0x6c,
	0x6c, 0x6f, 0x63, 0x61, 0x74, 0x65, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f,
	0x75, 0x70, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x37, 0x2e, 0x69,
	0x6e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x63,
	0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x2e, 0x61, 0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x65, 0x64, 0x2e,
	0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x23, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1d, 0x3a, 0x01, 0x2a,
	0x22, 0x18, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x65, 0x74, 0x2f, 0x61, 0x6c, 0x6c, 0x6f, 0x63, 0x61,
	0x74, 0x65, 0x64, 0x63, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x73, 0x12, 0xaf, 0x01, 0x0a, 0x0d, 0x47,
	0x65, 0x74, 0x41, 0x70, 0x70, 0x43, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x73, 0x12, 0x39, 0x2e, 0x69,
	0x6e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x63,
	0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x2e, 0x61, 0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x65, 0x64, 0x2e,
	0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x70, 0x70, 0x43, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3a, 0x2e, 0x69, 0x6e, 0x73, 0x70, 0x69, 0x72,
	0x65, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x63, 0x6f, 0x75, 0x70, 0x6f, 0x6e,
	0x2e, 0x61, 0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x65, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65,
	0x74, 0x41, 0x70, 0x70, 0x43, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x27, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x21, 0x3a, 0x01, 0x2a, 0x22, 0x1c,
	0x2f, 0x76, 0x31, 0x2f, 0x67, 0x65, 0x74, 0x2f, 0x61, 0x70, 0x70, 0x2f, 0x61, 0x6c, 0x6c, 0x6f,
	0x63, 0x61, 0x74, 0x65, 0x64, 0x63, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x73, 0x42, 0x47, 0x5a, 0x45,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4e, 0x70, 0x6f, 0x6f, 0x6c,
	0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x2f, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x69, 0x6e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2f, 0x67,
	0x77, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x2f, 0x61, 0x6c, 0x6c, 0x6f,
	0x63, 0x61, 0x74, 0x65, 0x64, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_npool_inspire_gw_v1_coupon_allocated_allocated_proto_rawDescOnce sync.Once
	file_npool_inspire_gw_v1_coupon_allocated_allocated_proto_rawDescData = file_npool_inspire_gw_v1_coupon_allocated_allocated_proto_rawDesc
)

func file_npool_inspire_gw_v1_coupon_allocated_allocated_proto_rawDescGZIP() []byte {
	file_npool_inspire_gw_v1_coupon_allocated_allocated_proto_rawDescOnce.Do(func() {
		file_npool_inspire_gw_v1_coupon_allocated_allocated_proto_rawDescData = protoimpl.X.CompressGZIP(file_npool_inspire_gw_v1_coupon_allocated_allocated_proto_rawDescData)
	})
	return file_npool_inspire_gw_v1_coupon_allocated_allocated_proto_rawDescData
}

var file_npool_inspire_gw_v1_coupon_allocated_allocated_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_npool_inspire_gw_v1_coupon_allocated_allocated_proto_goTypes = []interface{}{
	(*CreateCouponRequest)(nil),   // 0: inspire.gateway.coupon.allocated.v1.CreateCouponRequest
	(*CreateCouponResponse)(nil),  // 1: inspire.gateway.coupon.allocated.v1.CreateCouponResponse
	(*GetCouponsRequest)(nil),     // 2: inspire.gateway.coupon.allocated.v1.GetCouponsRequest
	(*GetCouponsResponse)(nil),    // 3: inspire.gateway.coupon.allocated.v1.GetCouponsResponse
	(*GetAppCouponsRequest)(nil),  // 4: inspire.gateway.coupon.allocated.v1.GetAppCouponsRequest
	(*GetAppCouponsResponse)(nil), // 5: inspire.gateway.coupon.allocated.v1.GetAppCouponsResponse
	(*allocated.Coupon)(nil),      // 6: inspire.middleware.coupon.allocated.v1.Coupon
}
var file_npool_inspire_gw_v1_coupon_allocated_allocated_proto_depIdxs = []int32{
	6, // 0: inspire.gateway.coupon.allocated.v1.CreateCouponResponse.Info:type_name -> inspire.middleware.coupon.allocated.v1.Coupon
	6, // 1: inspire.gateway.coupon.allocated.v1.GetCouponsResponse.Infos:type_name -> inspire.middleware.coupon.allocated.v1.Coupon
	6, // 2: inspire.gateway.coupon.allocated.v1.GetAppCouponsResponse.Infos:type_name -> inspire.middleware.coupon.allocated.v1.Coupon
	0, // 3: inspire.gateway.coupon.allocated.v1.Gateway.CreateCoupon:input_type -> inspire.gateway.coupon.allocated.v1.CreateCouponRequest
	2, // 4: inspire.gateway.coupon.allocated.v1.Gateway.GetCoupons:input_type -> inspire.gateway.coupon.allocated.v1.GetCouponsRequest
	4, // 5: inspire.gateway.coupon.allocated.v1.Gateway.GetAppCoupons:input_type -> inspire.gateway.coupon.allocated.v1.GetAppCouponsRequest
	1, // 6: inspire.gateway.coupon.allocated.v1.Gateway.CreateCoupon:output_type -> inspire.gateway.coupon.allocated.v1.CreateCouponResponse
	3, // 7: inspire.gateway.coupon.allocated.v1.Gateway.GetCoupons:output_type -> inspire.gateway.coupon.allocated.v1.GetCouponsResponse
	5, // 8: inspire.gateway.coupon.allocated.v1.Gateway.GetAppCoupons:output_type -> inspire.gateway.coupon.allocated.v1.GetAppCouponsResponse
	6, // [6:9] is the sub-list for method output_type
	3, // [3:6] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_npool_inspire_gw_v1_coupon_allocated_allocated_proto_init() }
func file_npool_inspire_gw_v1_coupon_allocated_allocated_proto_init() {
	if File_npool_inspire_gw_v1_coupon_allocated_allocated_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_npool_inspire_gw_v1_coupon_allocated_allocated_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateCouponRequest); i {
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
		file_npool_inspire_gw_v1_coupon_allocated_allocated_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateCouponResponse); i {
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
		file_npool_inspire_gw_v1_coupon_allocated_allocated_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCouponsRequest); i {
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
		file_npool_inspire_gw_v1_coupon_allocated_allocated_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCouponsResponse); i {
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
		file_npool_inspire_gw_v1_coupon_allocated_allocated_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAppCouponsRequest); i {
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
		file_npool_inspire_gw_v1_coupon_allocated_allocated_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAppCouponsResponse); i {
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
			RawDescriptor: file_npool_inspire_gw_v1_coupon_allocated_allocated_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_npool_inspire_gw_v1_coupon_allocated_allocated_proto_goTypes,
		DependencyIndexes: file_npool_inspire_gw_v1_coupon_allocated_allocated_proto_depIdxs,
		MessageInfos:      file_npool_inspire_gw_v1_coupon_allocated_allocated_proto_msgTypes,
	}.Build()
	File_npool_inspire_gw_v1_coupon_allocated_allocated_proto = out.File
	file_npool_inspire_gw_v1_coupon_allocated_allocated_proto_rawDesc = nil
	file_npool_inspire_gw_v1_coupon_allocated_allocated_proto_goTypes = nil
	file_npool_inspire_gw_v1_coupon_allocated_allocated_proto_depIdxs = nil
}
