// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.18.1
// source: npool/basetypes/review/v1/enums.proto

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

type ReviewObjectType int32

const (
	ReviewObjectType_DefaultObjectType      ReviewObjectType = 0
	ReviewObjectType_ObjectKyc              ReviewObjectType = 10
	ReviewObjectType_ObjectWithdrawal       ReviewObjectType = 20
	ReviewObjectType_ObjectRandomCouponCash ReviewObjectType = 30
)

// Enum value maps for ReviewObjectType.
var (
	ReviewObjectType_name = map[int32]string{
		0:  "DefaultObjectType",
		10: "ObjectKyc",
		20: "ObjectWithdrawal",
		30: "ObjectRandomCouponCash",
	}
	ReviewObjectType_value = map[string]int32{
		"DefaultObjectType":      0,
		"ObjectKyc":              10,
		"ObjectWithdrawal":       20,
		"ObjectRandomCouponCash": 30,
	}
)

func (x ReviewObjectType) Enum() *ReviewObjectType {
	p := new(ReviewObjectType)
	*p = x
	return p
}

func (x ReviewObjectType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ReviewObjectType) Descriptor() protoreflect.EnumDescriptor {
	return file_npool_basetypes_review_v1_enums_proto_enumTypes[0].Descriptor()
}

func (ReviewObjectType) Type() protoreflect.EnumType {
	return &file_npool_basetypes_review_v1_enums_proto_enumTypes[0]
}

func (x ReviewObjectType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ReviewObjectType.Descriptor instead.
func (ReviewObjectType) EnumDescriptor() ([]byte, []int) {
	return file_npool_basetypes_review_v1_enums_proto_rawDescGZIP(), []int{0}
}

type ReviewState int32

const (
	ReviewState_DefaultReviewState ReviewState = 0
	ReviewState_Approved           ReviewState = 10
	ReviewState_Wait               ReviewState = 20
	ReviewState_Rejected           ReviewState = 30
)

// Enum value maps for ReviewState.
var (
	ReviewState_name = map[int32]string{
		0:  "DefaultReviewState",
		10: "Approved",
		20: "Wait",
		30: "Rejected",
	}
	ReviewState_value = map[string]int32{
		"DefaultReviewState": 0,
		"Approved":           10,
		"Wait":               20,
		"Rejected":           30,
	}
)

func (x ReviewState) Enum() *ReviewState {
	p := new(ReviewState)
	*p = x
	return p
}

func (x ReviewState) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ReviewState) Descriptor() protoreflect.EnumDescriptor {
	return file_npool_basetypes_review_v1_enums_proto_enumTypes[1].Descriptor()
}

func (ReviewState) Type() protoreflect.EnumType {
	return &file_npool_basetypes_review_v1_enums_proto_enumTypes[1]
}

func (x ReviewState) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ReviewState.Descriptor instead.
func (ReviewState) EnumDescriptor() ([]byte, []int) {
	return file_npool_basetypes_review_v1_enums_proto_rawDescGZIP(), []int{1}
}

type ReviewTriggerType int32

const (
	ReviewTriggerType_DefaultTriggerType ReviewTriggerType = 0
	// Withdrawal triggers
	ReviewTriggerType_AutoReviewed         ReviewTriggerType = 10
	ReviewTriggerType_LargeAmount          ReviewTriggerType = 20
	ReviewTriggerType_InsufficientFunds    ReviewTriggerType = 30
	ReviewTriggerType_InsufficientGas      ReviewTriggerType = 40
	ReviewTriggerType_InsufficientFundsGas ReviewTriggerType = 50
)

// Enum value maps for ReviewTriggerType.
var (
	ReviewTriggerType_name = map[int32]string{
		0:  "DefaultTriggerType",
		10: "AutoReviewed",
		20: "LargeAmount",
		30: "InsufficientFunds",
		40: "InsufficientGas",
		50: "InsufficientFundsGas",
	}
	ReviewTriggerType_value = map[string]int32{
		"DefaultTriggerType":   0,
		"AutoReviewed":         10,
		"LargeAmount":          20,
		"InsufficientFunds":    30,
		"InsufficientGas":      40,
		"InsufficientFundsGas": 50,
	}
)

func (x ReviewTriggerType) Enum() *ReviewTriggerType {
	p := new(ReviewTriggerType)
	*p = x
	return p
}

func (x ReviewTriggerType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ReviewTriggerType) Descriptor() protoreflect.EnumDescriptor {
	return file_npool_basetypes_review_v1_enums_proto_enumTypes[2].Descriptor()
}

func (ReviewTriggerType) Type() protoreflect.EnumType {
	return &file_npool_basetypes_review_v1_enums_proto_enumTypes[2]
}

func (x ReviewTriggerType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ReviewTriggerType.Descriptor instead.
func (ReviewTriggerType) EnumDescriptor() ([]byte, []int) {
	return file_npool_basetypes_review_v1_enums_proto_rawDescGZIP(), []int{2}
}

var File_npool_basetypes_review_v1_enums_proto protoreflect.FileDescriptor

var file_npool_basetypes_review_v1_enums_proto_rawDesc = []byte{
	0x0a, 0x25, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x2f, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x6e, 0x75, 0x6d,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70,
	0x65, 0x73, 0x2e, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x2e, 0x76, 0x31, 0x2a, 0x6a, 0x0a, 0x10,
	0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x15, 0x0a, 0x11, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x4f, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x54, 0x79, 0x70, 0x65, 0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09, 0x4f, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x4b, 0x79, 0x63, 0x10, 0x0a, 0x12, 0x14, 0x0a, 0x10, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x57, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x61, 0x6c, 0x10, 0x14, 0x12, 0x1a, 0x0a, 0x16,
	0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x61, 0x6e, 0x64, 0x6f, 0x6d, 0x43, 0x6f, 0x75, 0x70,
	0x6f, 0x6e, 0x43, 0x61, 0x73, 0x68, 0x10, 0x1e, 0x2a, 0x4b, 0x0a, 0x0b, 0x52, 0x65, 0x76, 0x69,
	0x65, 0x77, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x16, 0x0a, 0x12, 0x44, 0x65, 0x66, 0x61, 0x75,
	0x6c, 0x74, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x53, 0x74, 0x61, 0x74, 0x65, 0x10, 0x00, 0x12,
	0x0c, 0x0a, 0x08, 0x41, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x65, 0x64, 0x10, 0x0a, 0x12, 0x08, 0x0a,
	0x04, 0x57, 0x61, 0x69, 0x74, 0x10, 0x14, 0x12, 0x0c, 0x0a, 0x08, 0x52, 0x65, 0x6a, 0x65, 0x63,
	0x74, 0x65, 0x64, 0x10, 0x1e, 0x2a, 0x94, 0x01, 0x0a, 0x11, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77,
	0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x12, 0x44,
	0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x54, 0x79, 0x70,
	0x65, 0x10, 0x00, 0x12, 0x10, 0x0a, 0x0c, 0x41, 0x75, 0x74, 0x6f, 0x52, 0x65, 0x76, 0x69, 0x65,
	0x77, 0x65, 0x64, 0x10, 0x0a, 0x12, 0x0f, 0x0a, 0x0b, 0x4c, 0x61, 0x72, 0x67, 0x65, 0x41, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x10, 0x14, 0x12, 0x15, 0x0a, 0x11, 0x49, 0x6e, 0x73, 0x75, 0x66, 0x66,
	0x69, 0x63, 0x69, 0x65, 0x6e, 0x74, 0x46, 0x75, 0x6e, 0x64, 0x73, 0x10, 0x1e, 0x12, 0x13, 0x0a,
	0x0f, 0x49, 0x6e, 0x73, 0x75, 0x66, 0x66, 0x69, 0x63, 0x69, 0x65, 0x6e, 0x74, 0x47, 0x61, 0x73,
	0x10, 0x28, 0x12, 0x18, 0x0a, 0x14, 0x49, 0x6e, 0x73, 0x75, 0x66, 0x66, 0x69, 0x63, 0x69, 0x65,
	0x6e, 0x74, 0x46, 0x75, 0x6e, 0x64, 0x73, 0x47, 0x61, 0x73, 0x10, 0x32, 0x42, 0x3c, 0x5a, 0x3a,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4e, 0x70, 0x6f, 0x6f, 0x6c,
	0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x2f, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x2f, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_npool_basetypes_review_v1_enums_proto_rawDescOnce sync.Once
	file_npool_basetypes_review_v1_enums_proto_rawDescData = file_npool_basetypes_review_v1_enums_proto_rawDesc
)

func file_npool_basetypes_review_v1_enums_proto_rawDescGZIP() []byte {
	file_npool_basetypes_review_v1_enums_proto_rawDescOnce.Do(func() {
		file_npool_basetypes_review_v1_enums_proto_rawDescData = protoimpl.X.CompressGZIP(file_npool_basetypes_review_v1_enums_proto_rawDescData)
	})
	return file_npool_basetypes_review_v1_enums_proto_rawDescData
}

var file_npool_basetypes_review_v1_enums_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_npool_basetypes_review_v1_enums_proto_goTypes = []interface{}{
	(ReviewObjectType)(0),  // 0: basetypes.review.v1.ReviewObjectType
	(ReviewState)(0),       // 1: basetypes.review.v1.ReviewState
	(ReviewTriggerType)(0), // 2: basetypes.review.v1.ReviewTriggerType
}
var file_npool_basetypes_review_v1_enums_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_npool_basetypes_review_v1_enums_proto_init() }
func file_npool_basetypes_review_v1_enums_proto_init() {
	if File_npool_basetypes_review_v1_enums_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_npool_basetypes_review_v1_enums_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_npool_basetypes_review_v1_enums_proto_goTypes,
		DependencyIndexes: file_npool_basetypes_review_v1_enums_proto_depIdxs,
		EnumInfos:         file_npool_basetypes_review_v1_enums_proto_enumTypes,
	}.Build()
	File_npool_basetypes_review_v1_enums_proto = out.File
	file_npool_basetypes_review_v1_enums_proto_rawDesc = nil
	file_npool_basetypes_review_v1_enums_proto_goTypes = nil
	file_npool_basetypes_review_v1_enums_proto_depIdxs = nil
}
