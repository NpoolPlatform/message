// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.18.1
// source: npool/inspire/mw/v1/inspire/coupon/coupon.proto

package coupon

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

type CouponType int32

const (
	CouponType_DefaultCouponType  CouponType = 0
	CouponType_FixAmount          CouponType = 10
	CouponType_Discount           CouponType = 20
	CouponType_SpecialOffer       CouponType = 30
	CouponType_ThresholdReduction CouponType = 40
)

// Enum value maps for CouponType.
var (
	CouponType_name = map[int32]string{
		0:  "DefaultCouponType",
		10: "FixAmount",
		20: "Discount",
		30: "SpecialOffer",
		40: "ThresholdReduction",
	}
	CouponType_value = map[string]int32{
		"DefaultCouponType":  0,
		"FixAmount":          10,
		"Discount":           20,
		"SpecialOffer":       30,
		"ThresholdReduction": 40,
	}
)

func (x CouponType) Enum() *CouponType {
	p := new(CouponType)
	*p = x
	return p
}

func (x CouponType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CouponType) Descriptor() protoreflect.EnumDescriptor {
	return file_npool_inspire_mw_v1_inspire_coupon_coupon_proto_enumTypes[0].Descriptor()
}

func (CouponType) Type() protoreflect.EnumType {
	return &file_npool_inspire_mw_v1_inspire_coupon_coupon_proto_enumTypes[0]
}

func (x CouponType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CouponType.Descriptor instead.
func (CouponType) EnumDescriptor() ([]byte, []int) {
	return file_npool_inspire_mw_v1_inspire_coupon_coupon_proto_rawDescGZIP(), []int{0}
}

type Coupon struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID              string     `protobuf:"bytes,10,opt,name=ID,proto3" json:"ID,omitempty"`
	Type            CouponType `protobuf:"varint,20,opt,name=Type,proto3,enum=inspire.middleware.inspire.coupon.v1.CouponType" json:"Type,omitempty"`
	Value           string     `protobuf:"bytes,30,opt,name=Value,proto3" json:"Value,omitempty"`
	ThresholdAmount string     `protobuf:"bytes,40,opt,name=ThresholdAmount,proto3" json:"ThresholdAmount,omitempty"`
	Start           uint32     `protobuf:"varint,50,opt,name=Start,proto3" json:"Start,omitempty"`
	End             uint32     `protobuf:"varint,60,opt,name=End,proto3" json:"End,omitempty"`
	Name            string     `protobuf:"bytes,70,opt,name=Name,proto3" json:"Name,omitempty"`
	Message         string     `protobuf:"bytes,80,opt,name=Message,proto3" json:"Message,omitempty"`
}

func (x *Coupon) Reset() {
	*x = Coupon{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_inspire_mw_v1_inspire_coupon_coupon_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Coupon) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Coupon) ProtoMessage() {}

func (x *Coupon) ProtoReflect() protoreflect.Message {
	mi := &file_npool_inspire_mw_v1_inspire_coupon_coupon_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Coupon.ProtoReflect.Descriptor instead.
func (*Coupon) Descriptor() ([]byte, []int) {
	return file_npool_inspire_mw_v1_inspire_coupon_coupon_proto_rawDescGZIP(), []int{0}
}

func (x *Coupon) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *Coupon) GetType() CouponType {
	if x != nil {
		return x.Type
	}
	return CouponType_DefaultCouponType
}

func (x *Coupon) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *Coupon) GetThresholdAmount() string {
	if x != nil {
		return x.ThresholdAmount
	}
	return ""
}

func (x *Coupon) GetStart() uint32 {
	if x != nil {
		return x.Start
	}
	return 0
}

func (x *Coupon) GetEnd() uint32 {
	if x != nil {
		return x.End
	}
	return 0
}

func (x *Coupon) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Coupon) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type GetCouponRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID   string     `protobuf:"bytes,10,opt,name=ID,proto3" json:"ID,omitempty"`
	Type CouponType `protobuf:"varint,20,opt,name=Type,proto3,enum=inspire.middleware.inspire.coupon.v1.CouponType" json:"Type,omitempty"`
}

func (x *GetCouponRequest) Reset() {
	*x = GetCouponRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_inspire_mw_v1_inspire_coupon_coupon_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCouponRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCouponRequest) ProtoMessage() {}

func (x *GetCouponRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_inspire_mw_v1_inspire_coupon_coupon_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCouponRequest.ProtoReflect.Descriptor instead.
func (*GetCouponRequest) Descriptor() ([]byte, []int) {
	return file_npool_inspire_mw_v1_inspire_coupon_coupon_proto_rawDescGZIP(), []int{1}
}

func (x *GetCouponRequest) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *GetCouponRequest) GetType() CouponType {
	if x != nil {
		return x.Type
	}
	return CouponType_DefaultCouponType
}

type GetCouponResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *Coupon `protobuf:"bytes,10,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (x *GetCouponResponse) Reset() {
	*x = GetCouponResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_inspire_mw_v1_inspire_coupon_coupon_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCouponResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCouponResponse) ProtoMessage() {}

func (x *GetCouponResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_inspire_mw_v1_inspire_coupon_coupon_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCouponResponse.ProtoReflect.Descriptor instead.
func (*GetCouponResponse) Descriptor() ([]byte, []int) {
	return file_npool_inspire_mw_v1_inspire_coupon_coupon_proto_rawDescGZIP(), []int{2}
}

func (x *GetCouponResponse) GetInfo() *Coupon {
	if x != nil {
		return x.Info
	}
	return nil
}

var File_npool_inspire_mw_v1_inspire_coupon_coupon_proto protoreflect.FileDescriptor

var file_npool_inspire_mw_v1_inspire_coupon_coupon_proto_rawDesc = []byte{
	0x0a, 0x2f, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x69, 0x6e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2f,
	0x6d, 0x77, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x6e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2f, 0x63, 0x6f,
	0x75, 0x70, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x24, 0x69, 0x6e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c,
	0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x69, 0x6e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x63, 0x6f,
	0x75, 0x70, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x11, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf4, 0x01, 0x0a, 0x06, 0x43, 0x6f, 0x75, 0x70, 0x6f, 0x6e,
	0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44,
	0x12, 0x44, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x18, 0x14, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x30,
	0x2e, 0x69, 0x6e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77,
	0x61, 0x72, 0x65, 0x2e, 0x69, 0x6e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x63, 0x6f, 0x75, 0x70,
	0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65,
	0x52, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x1e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x28, 0x0a, 0x0f,
	0x54, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x28, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x54, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64,
	0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x72, 0x74, 0x18,
	0x32, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x53, 0x74, 0x61, 0x72, 0x74, 0x12, 0x10, 0x0a, 0x03,
	0x45, 0x6e, 0x64, 0x18, 0x3c, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x45, 0x6e, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x46, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x50, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x68, 0x0a, 0x10,
	0x47, 0x65, 0x74, 0x43, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44,
	0x12, 0x44, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x18, 0x14, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x30,
	0x2e, 0x69, 0x6e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77,
	0x61, 0x72, 0x65, 0x2e, 0x69, 0x6e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x63, 0x6f, 0x75, 0x70,
	0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65,
	0x52, 0x04, 0x54, 0x79, 0x70, 0x65, 0x22, 0x55, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x75,
	0x70, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x40, 0x0a, 0x04, 0x49,
	0x6e, 0x66, 0x6f, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x69, 0x6e, 0x73, 0x70,
	0x69, 0x72, 0x65, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x69,
	0x6e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x63, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x2e, 0x76, 0x31,
	0x2e, 0x43, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x52, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x2a, 0x6a, 0x0a,
	0x0a, 0x43, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x15, 0x0a, 0x11, 0x44,
	0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x43, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65,
	0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09, 0x46, 0x69, 0x78, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x10,
	0x0a, 0x12, 0x0c, 0x0a, 0x08, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x10, 0x14, 0x12,
	0x10, 0x0a, 0x0c, 0x53, 0x70, 0x65, 0x63, 0x69, 0x61, 0x6c, 0x4f, 0x66, 0x66, 0x65, 0x72, 0x10,
	0x1e, 0x12, 0x16, 0x0a, 0x12, 0x54, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x52, 0x65,
	0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x10, 0x28, 0x32, 0x8c, 0x01, 0x0a, 0x0a, 0x4d, 0x69,
	0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x12, 0x7e, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x43,
	0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x12, 0x36, 0x2e, 0x69, 0x6e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e,
	0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x69, 0x6e, 0x73, 0x70, 0x69,
	0x72, 0x65, 0x2e, 0x63, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74,
	0x43, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x37, 0x2e,
	0x69, 0x6e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61,
	0x72, 0x65, 0x2e, 0x69, 0x6e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x63, 0x6f, 0x75, 0x70, 0x6f,
	0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x45, 0x5a, 0x43, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4e, 0x70, 0x6f, 0x6f, 0x6c, 0x50, 0x6c, 0x61, 0x74,
	0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x6e, 0x70, 0x6f,
	0x6f, 0x6c, 0x2f, 0x69, 0x6e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2f, 0x6d, 0x77, 0x2f, 0x76, 0x31,
	0x2f, 0x69, 0x6e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2f, 0x63, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_npool_inspire_mw_v1_inspire_coupon_coupon_proto_rawDescOnce sync.Once
	file_npool_inspire_mw_v1_inspire_coupon_coupon_proto_rawDescData = file_npool_inspire_mw_v1_inspire_coupon_coupon_proto_rawDesc
)

func file_npool_inspire_mw_v1_inspire_coupon_coupon_proto_rawDescGZIP() []byte {
	file_npool_inspire_mw_v1_inspire_coupon_coupon_proto_rawDescOnce.Do(func() {
		file_npool_inspire_mw_v1_inspire_coupon_coupon_proto_rawDescData = protoimpl.X.CompressGZIP(file_npool_inspire_mw_v1_inspire_coupon_coupon_proto_rawDescData)
	})
	return file_npool_inspire_mw_v1_inspire_coupon_coupon_proto_rawDescData
}

var file_npool_inspire_mw_v1_inspire_coupon_coupon_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_npool_inspire_mw_v1_inspire_coupon_coupon_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_npool_inspire_mw_v1_inspire_coupon_coupon_proto_goTypes = []interface{}{
	(CouponType)(0),           // 0: inspire.middleware.inspire.coupon.v1.CouponType
	(*Coupon)(nil),            // 1: inspire.middleware.inspire.coupon.v1.Coupon
	(*GetCouponRequest)(nil),  // 2: inspire.middleware.inspire.coupon.v1.GetCouponRequest
	(*GetCouponResponse)(nil), // 3: inspire.middleware.inspire.coupon.v1.GetCouponResponse
}
var file_npool_inspire_mw_v1_inspire_coupon_coupon_proto_depIdxs = []int32{
	0, // 0: inspire.middleware.inspire.coupon.v1.Coupon.Type:type_name -> inspire.middleware.inspire.coupon.v1.CouponType
	0, // 1: inspire.middleware.inspire.coupon.v1.GetCouponRequest.Type:type_name -> inspire.middleware.inspire.coupon.v1.CouponType
	1, // 2: inspire.middleware.inspire.coupon.v1.GetCouponResponse.Info:type_name -> inspire.middleware.inspire.coupon.v1.Coupon
	2, // 3: inspire.middleware.inspire.coupon.v1.Middleware.GetCoupon:input_type -> inspire.middleware.inspire.coupon.v1.GetCouponRequest
	3, // 4: inspire.middleware.inspire.coupon.v1.Middleware.GetCoupon:output_type -> inspire.middleware.inspire.coupon.v1.GetCouponResponse
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_npool_inspire_mw_v1_inspire_coupon_coupon_proto_init() }
func file_npool_inspire_mw_v1_inspire_coupon_coupon_proto_init() {
	if File_npool_inspire_mw_v1_inspire_coupon_coupon_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_npool_inspire_mw_v1_inspire_coupon_coupon_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Coupon); i {
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
		file_npool_inspire_mw_v1_inspire_coupon_coupon_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCouponRequest); i {
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
		file_npool_inspire_mw_v1_inspire_coupon_coupon_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCouponResponse); i {
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
			RawDescriptor: file_npool_inspire_mw_v1_inspire_coupon_coupon_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_npool_inspire_mw_v1_inspire_coupon_coupon_proto_goTypes,
		DependencyIndexes: file_npool_inspire_mw_v1_inspire_coupon_coupon_proto_depIdxs,
		EnumInfos:         file_npool_inspire_mw_v1_inspire_coupon_coupon_proto_enumTypes,
		MessageInfos:      file_npool_inspire_mw_v1_inspire_coupon_coupon_proto_msgTypes,
	}.Build()
	File_npool_inspire_mw_v1_inspire_coupon_coupon_proto = out.File
	file_npool_inspire_mw_v1_inspire_coupon_coupon_proto_rawDesc = nil
	file_npool_inspire_mw_v1_inspire_coupon_coupon_proto_goTypes = nil
	file_npool_inspire_mw_v1_inspire_coupon_coupon_proto_depIdxs = nil
}
