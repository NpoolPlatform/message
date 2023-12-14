// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.18.1
// source: npool/basetypes/inspire/v1/enums.proto

package v1

import (
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

type SettleType int32

const (
	SettleType_DefaultSettleType SettleType = 0
	// Divide commission according to order payment amount or good value
	SettleType_GoodOrderPayment SettleType = 10
	// Divice commission according to order technique fee
	SettleType_TechniqueServiceFee SettleType = 20
)

// Enum value maps for SettleType.
var (
	SettleType_name = map[int32]string{
		0:  "DefaultSettleType",
		10: "GoodOrderPayment",
		20: "TechniqueServiceFee",
	}
	SettleType_value = map[string]int32{
		"DefaultSettleType":   0,
		"GoodOrderPayment":    10,
		"TechniqueServiceFee": 20,
	}
)

func (x SettleType) Enum() *SettleType {
	p := new(SettleType)
	*p = x
	return p
}

func (x SettleType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SettleType) Descriptor() protoreflect.EnumDescriptor {
	return file_npool_basetypes_inspire_v1_enums_proto_enumTypes[0].Descriptor()
}

func (SettleType) Type() protoreflect.EnumType {
	return &file_npool_basetypes_inspire_v1_enums_proto_enumTypes[0]
}

func (x SettleType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SettleType.Descriptor instead.
func (SettleType) EnumDescriptor() ([]byte, []int) {
	return file_npool_basetypes_inspire_v1_enums_proto_rawDescGZIP(), []int{0}
}

type SettleMode int32

const (
	SettleMode_DefaultSettleMode       SettleMode = 0
	SettleMode_SettleWithGoodValue     SettleMode = 10
	SettleMode_SettleWithPaymentAmount SettleMode = 20
)

// Enum value maps for SettleMode.
var (
	SettleMode_name = map[int32]string{
		0:  "DefaultSettleMode",
		10: "SettleWithGoodValue",
		20: "SettleWithPaymentAmount",
	}
	SettleMode_value = map[string]int32{
		"DefaultSettleMode":       0,
		"SettleWithGoodValue":     10,
		"SettleWithPaymentAmount": 20,
	}
)

func (x SettleMode) Enum() *SettleMode {
	p := new(SettleMode)
	*p = x
	return p
}

func (x SettleMode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SettleMode) Descriptor() protoreflect.EnumDescriptor {
	return file_npool_basetypes_inspire_v1_enums_proto_enumTypes[1].Descriptor()
}

func (SettleMode) Type() protoreflect.EnumType {
	return &file_npool_basetypes_inspire_v1_enums_proto_enumTypes[1]
}

func (x SettleMode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SettleMode.Descriptor instead.
func (SettleMode) EnumDescriptor() ([]byte, []int) {
	return file_npool_basetypes_inspire_v1_enums_proto_rawDescGZIP(), []int{1}
}

type SettleAmountType int32

const (
	SettleAmountType_DefaultSettleAmountType SettleAmountType = 0
	SettleAmountType_SettleByPercent         SettleAmountType = 10
	SettleAmountType_SettleByAmount          SettleAmountType = 20
)

// Enum value maps for SettleAmountType.
var (
	SettleAmountType_name = map[int32]string{
		0:  "DefaultSettleAmountType",
		10: "SettleByPercent",
		20: "SettleByAmount",
	}
	SettleAmountType_value = map[string]int32{
		"DefaultSettleAmountType": 0,
		"SettleByPercent":         10,
		"SettleByAmount":          20,
	}
)

func (x SettleAmountType) Enum() *SettleAmountType {
	p := new(SettleAmountType)
	*p = x
	return p
}

func (x SettleAmountType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SettleAmountType) Descriptor() protoreflect.EnumDescriptor {
	return file_npool_basetypes_inspire_v1_enums_proto_enumTypes[2].Descriptor()
}

func (SettleAmountType) Type() protoreflect.EnumType {
	return &file_npool_basetypes_inspire_v1_enums_proto_enumTypes[2]
}

func (x SettleAmountType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SettleAmountType.Descriptor instead.
func (SettleAmountType) EnumDescriptor() ([]byte, []int) {
	return file_npool_basetypes_inspire_v1_enums_proto_rawDescGZIP(), []int{2}
}

type SettleInterval int32

const (
	SettleInterval_DefaultSettleInterval SettleInterval = 0
	SettleInterval_SettleAggregate       SettleInterval = 10
	SettleInterval_SettleMonthly         SettleInterval = 20
	SettleInterval_SettleYearly          SettleInterval = 30
	SettleInterval_SettleEveryOrder      SettleInterval = 40
)

// Enum value maps for SettleInterval.
var (
	SettleInterval_name = map[int32]string{
		0:  "DefaultSettleInterval",
		10: "SettleAggregate",
		20: "SettleMonthly",
		30: "SettleYearly",
		40: "SettleEveryOrder",
	}
	SettleInterval_value = map[string]int32{
		"DefaultSettleInterval": 0,
		"SettleAggregate":       10,
		"SettleMonthly":         20,
		"SettleYearly":          30,
		"SettleEveryOrder":      40,
	}
)

func (x SettleInterval) Enum() *SettleInterval {
	p := new(SettleInterval)
	*p = x
	return p
}

func (x SettleInterval) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SettleInterval) Descriptor() protoreflect.EnumDescriptor {
	return file_npool_basetypes_inspire_v1_enums_proto_enumTypes[3].Descriptor()
}

func (SettleInterval) Type() protoreflect.EnumType {
	return &file_npool_basetypes_inspire_v1_enums_proto_enumTypes[3]
}

func (x SettleInterval) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SettleInterval.Descriptor instead.
func (SettleInterval) EnumDescriptor() ([]byte, []int) {
	return file_npool_basetypes_inspire_v1_enums_proto_rawDescGZIP(), []int{3}
}

type CouponType int32

const (
	CouponType_DefaultCouponType CouponType = 0
	CouponType_FixAmount         CouponType = 10
	CouponType_Discount          CouponType = 20
)

// Enum value maps for CouponType.
var (
	CouponType_name = map[int32]string{
		0:  "DefaultCouponType",
		10: "FixAmount",
		20: "Discount",
	}
	CouponType_value = map[string]int32{
		"DefaultCouponType": 0,
		"FixAmount":         10,
		"Discount":          20,
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
	return file_npool_basetypes_inspire_v1_enums_proto_enumTypes[4].Descriptor()
}

func (CouponType) Type() protoreflect.EnumType {
	return &file_npool_basetypes_inspire_v1_enums_proto_enumTypes[4]
}

func (x CouponType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CouponType.Descriptor instead.
func (CouponType) EnumDescriptor() ([]byte, []int) {
	return file_npool_basetypes_inspire_v1_enums_proto_rawDescGZIP(), []int{4}
}

type CouponConstraint int32

const (
	CouponConstraint_DefaultCouponConstraint CouponConstraint = 0
	CouponConstraint_Normal                  CouponConstraint = 10
	CouponConstraint_PaymentThreshold        CouponConstraint = 20
)

// Enum value maps for CouponConstraint.
var (
	CouponConstraint_name = map[int32]string{
		0:  "DefaultCouponConstraint",
		10: "Normal",
		20: "PaymentThreshold",
	}
	CouponConstraint_value = map[string]int32{
		"DefaultCouponConstraint": 0,
		"Normal":                  10,
		"PaymentThreshold":        20,
	}
)

func (x CouponConstraint) Enum() *CouponConstraint {
	p := new(CouponConstraint)
	*p = x
	return p
}

func (x CouponConstraint) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CouponConstraint) Descriptor() protoreflect.EnumDescriptor {
	return file_npool_basetypes_inspire_v1_enums_proto_enumTypes[5].Descriptor()
}

func (CouponConstraint) Type() protoreflect.EnumType {
	return &file_npool_basetypes_inspire_v1_enums_proto_enumTypes[5]
}

func (x CouponConstraint) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CouponConstraint.Descriptor instead.
func (CouponConstraint) EnumDescriptor() ([]byte, []int) {
	return file_npool_basetypes_inspire_v1_enums_proto_rawDescGZIP(), []int{5}
}

type CouponScope int32

const (
	CouponScope_DefaultCouponScope CouponScope = 0
	CouponScope_AllGood            CouponScope = 10
	CouponScope_Whitelist          CouponScope = 20
	CouponScope_Blacklist          CouponScope = 30
)

// Enum value maps for CouponScope.
var (
	CouponScope_name = map[int32]string{
		0:  "DefaultCouponScope",
		10: "AllGood",
		20: "Whitelist",
		30: "Blacklist",
	}
	CouponScope_value = map[string]int32{
		"DefaultCouponScope": 0,
		"AllGood":            10,
		"Whitelist":          20,
		"Blacklist":          30,
	}
)

func (x CouponScope) Enum() *CouponScope {
	p := new(CouponScope)
	*p = x
	return p
}

func (x CouponScope) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CouponScope) Descriptor() protoreflect.EnumDescriptor {
	return file_npool_basetypes_inspire_v1_enums_proto_enumTypes[6].Descriptor()
}

func (CouponScope) Type() protoreflect.EnumType {
	return &file_npool_basetypes_inspire_v1_enums_proto_enumTypes[6]
}

func (x CouponScope) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CouponScope.Descriptor instead.
func (CouponScope) EnumDescriptor() ([]byte, []int) {
	return file_npool_basetypes_inspire_v1_enums_proto_rawDescGZIP(), []int{6}
}

var File_npool_basetypes_inspire_v1_enums_proto protoreflect.FileDescriptor

var file_npool_basetypes_inspire_v1_enums_proto_rawDesc = []byte{
	0x0a, 0x26, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x2f, 0x69, 0x6e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x6e, 0x75,
	0x6d, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x14, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x2e, 0x69, 0x6e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2a, 0x52,
	0x0a, 0x0a, 0x53, 0x65, 0x74, 0x74, 0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x15, 0x0a, 0x11,
	0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x53, 0x65, 0x74, 0x74, 0x6c, 0x65, 0x54, 0x79, 0x70,
	0x65, 0x10, 0x00, 0x12, 0x14, 0x0a, 0x10, 0x47, 0x6f, 0x6f, 0x64, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x10, 0x0a, 0x12, 0x17, 0x0a, 0x13, 0x54, 0x65, 0x63,
	0x68, 0x6e, 0x69, 0x71, 0x75, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x46, 0x65, 0x65,
	0x10, 0x14, 0x2a, 0x59, 0x0a, 0x0a, 0x53, 0x65, 0x74, 0x74, 0x6c, 0x65, 0x4d, 0x6f, 0x64, 0x65,
	0x12, 0x15, 0x0a, 0x11, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x53, 0x65, 0x74, 0x74, 0x6c,
	0x65, 0x4d, 0x6f, 0x64, 0x65, 0x10, 0x00, 0x12, 0x17, 0x0a, 0x13, 0x53, 0x65, 0x74, 0x74, 0x6c,
	0x65, 0x57, 0x69, 0x74, 0x68, 0x47, 0x6f, 0x6f, 0x64, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x10, 0x0a,
	0x12, 0x1b, 0x0a, 0x17, 0x53, 0x65, 0x74, 0x74, 0x6c, 0x65, 0x57, 0x69, 0x74, 0x68, 0x50, 0x61,
	0x79, 0x6d, 0x65, 0x6e, 0x74, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x10, 0x14, 0x2a, 0x58, 0x0a,
	0x10, 0x53, 0x65, 0x74, 0x74, 0x6c, 0x65, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x1b, 0x0a, 0x17, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x53, 0x65, 0x74, 0x74,
	0x6c, 0x65, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x10, 0x00, 0x12, 0x13,
	0x0a, 0x0f, 0x53, 0x65, 0x74, 0x74, 0x6c, 0x65, 0x42, 0x79, 0x50, 0x65, 0x72, 0x63, 0x65, 0x6e,
	0x74, 0x10, 0x0a, 0x12, 0x12, 0x0a, 0x0e, 0x53, 0x65, 0x74, 0x74, 0x6c, 0x65, 0x42, 0x79, 0x41,
	0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x10, 0x14, 0x2a, 0x7b, 0x0a, 0x0e, 0x53, 0x65, 0x74, 0x74, 0x6c,
	0x65, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x12, 0x19, 0x0a, 0x15, 0x44, 0x65, 0x66,
	0x61, 0x75, 0x6c, 0x74, 0x53, 0x65, 0x74, 0x74, 0x6c, 0x65, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76,
	0x61, 0x6c, 0x10, 0x00, 0x12, 0x13, 0x0a, 0x0f, 0x53, 0x65, 0x74, 0x74, 0x6c, 0x65, 0x41, 0x67,
	0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x10, 0x0a, 0x12, 0x11, 0x0a, 0x0d, 0x53, 0x65, 0x74,
	0x74, 0x6c, 0x65, 0x4d, 0x6f, 0x6e, 0x74, 0x68, 0x6c, 0x79, 0x10, 0x14, 0x12, 0x10, 0x0a, 0x0c,
	0x53, 0x65, 0x74, 0x74, 0x6c, 0x65, 0x59, 0x65, 0x61, 0x72, 0x6c, 0x79, 0x10, 0x1e, 0x12, 0x14,
	0x0a, 0x10, 0x53, 0x65, 0x74, 0x74, 0x6c, 0x65, 0x45, 0x76, 0x65, 0x72, 0x79, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x10, 0x28, 0x2a, 0x40, 0x0a, 0x0a, 0x43, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x15, 0x0a, 0x11, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x43, 0x6f, 0x75,
	0x70, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09, 0x46, 0x69, 0x78,
	0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x10, 0x0a, 0x12, 0x0c, 0x0a, 0x08, 0x44, 0x69, 0x73, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x10, 0x14, 0x2a, 0x51, 0x0a, 0x10, 0x43, 0x6f, 0x75, 0x70, 0x6f, 0x6e,
	0x43, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x74, 0x12, 0x1b, 0x0a, 0x17, 0x44, 0x65,
	0x66, 0x61, 0x75, 0x6c, 0x74, 0x43, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x73, 0x74,
	0x72, 0x61, 0x69, 0x6e, 0x74, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x4e, 0x6f, 0x72, 0x6d, 0x61,
	0x6c, 0x10, 0x0a, 0x12, 0x14, 0x0a, 0x10, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x68,
	0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x10, 0x14, 0x2a, 0x50, 0x0a, 0x0b, 0x43, 0x6f, 0x75,
	0x70, 0x6f, 0x6e, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x12, 0x44, 0x65, 0x66, 0x61,
	0x75, 0x6c, 0x74, 0x43, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x10, 0x00,
	0x12, 0x0b, 0x0a, 0x07, 0x41, 0x6c, 0x6c, 0x47, 0x6f, 0x6f, 0x64, 0x10, 0x0a, 0x12, 0x0d, 0x0a,
	0x09, 0x57, 0x68, 0x69, 0x74, 0x65, 0x6c, 0x69, 0x73, 0x74, 0x10, 0x14, 0x12, 0x0d, 0x0a, 0x09,
	0x42, 0x6c, 0x61, 0x63, 0x6b, 0x6c, 0x69, 0x73, 0x74, 0x10, 0x1e, 0x42, 0x3d, 0x5a, 0x3b, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4e, 0x70, 0x6f, 0x6f, 0x6c, 0x50,
	0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f,
	0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f,
	0x69, 0x6e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_npool_basetypes_inspire_v1_enums_proto_rawDescOnce sync.Once
	file_npool_basetypes_inspire_v1_enums_proto_rawDescData = file_npool_basetypes_inspire_v1_enums_proto_rawDesc
)

func file_npool_basetypes_inspire_v1_enums_proto_rawDescGZIP() []byte {
	file_npool_basetypes_inspire_v1_enums_proto_rawDescOnce.Do(func() {
		file_npool_basetypes_inspire_v1_enums_proto_rawDescData = protoimpl.X.CompressGZIP(file_npool_basetypes_inspire_v1_enums_proto_rawDescData)
	})
	return file_npool_basetypes_inspire_v1_enums_proto_rawDescData
}

var file_npool_basetypes_inspire_v1_enums_proto_enumTypes = make([]protoimpl.EnumInfo, 7)
var file_npool_basetypes_inspire_v1_enums_proto_goTypes = []interface{}{
	(SettleType)(0),       // 0: basetypes.inspire.v1.SettleType
	(SettleMode)(0),       // 1: basetypes.inspire.v1.SettleMode
	(SettleAmountType)(0), // 2: basetypes.inspire.v1.SettleAmountType
	(SettleInterval)(0),   // 3: basetypes.inspire.v1.SettleInterval
	(CouponType)(0),       // 4: basetypes.inspire.v1.CouponType
	(CouponConstraint)(0), // 5: basetypes.inspire.v1.CouponConstraint
	(CouponScope)(0),      // 6: basetypes.inspire.v1.CouponScope
}
var file_npool_basetypes_inspire_v1_enums_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_npool_basetypes_inspire_v1_enums_proto_init() }
func file_npool_basetypes_inspire_v1_enums_proto_init() {
	if File_npool_basetypes_inspire_v1_enums_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_npool_basetypes_inspire_v1_enums_proto_rawDesc,
			NumEnums:      7,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_npool_basetypes_inspire_v1_enums_proto_goTypes,
		DependencyIndexes: file_npool_basetypes_inspire_v1_enums_proto_depIdxs,
		EnumInfos:         file_npool_basetypes_inspire_v1_enums_proto_enumTypes,
	}.Build()
	File_npool_basetypes_inspire_v1_enums_proto = out.File
	file_npool_basetypes_inspire_v1_enums_proto_rawDesc = nil
	file_npool_basetypes_inspire_v1_enums_proto_goTypes = nil
	file_npool_basetypes_inspire_v1_enums_proto_depIdxs = nil
}
