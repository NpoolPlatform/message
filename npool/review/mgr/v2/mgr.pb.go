// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.18.1
// source: npool/review/mgr/v2/mgr.proto

package v1

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

type ReviewObjectType int32

const (
	ReviewObjectType_DefaultObjectType ReviewObjectType = 0
	ReviewObjectType_ObjectKyc         ReviewObjectType = 10
	ReviewObjectType_ObjectWithdrawal  ReviewObjectType = 20
)

// Enum value maps for ReviewObjectType.
var (
	ReviewObjectType_name = map[int32]string{
		0:  "DefaultObjectType",
		10: "ObjectKyc",
		20: "ObjectWithdrawal",
	}
	ReviewObjectType_value = map[string]int32{
		"DefaultObjectType": 0,
		"ObjectKyc":         10,
		"ObjectWithdrawal":  20,
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
	return file_npool_review_mgr_v2_mgr_proto_enumTypes[0].Descriptor()
}

func (ReviewObjectType) Type() protoreflect.EnumType {
	return &file_npool_review_mgr_v2_mgr_proto_enumTypes[0]
}

func (x ReviewObjectType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ReviewObjectType.Descriptor instead.
func (ReviewObjectType) EnumDescriptor() ([]byte, []int) {
	return file_npool_review_mgr_v2_mgr_proto_rawDescGZIP(), []int{0}
}

var File_npool_review_mgr_v2_mgr_proto protoreflect.FileDescriptor

var file_npool_review_mgr_v2_mgr_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x2f, 0x6d,
	0x67, 0x72, 0x2f, 0x76, 0x32, 0x2f, 0x6d, 0x67, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x11, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e,
	0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x6e,
	0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2a, 0x4e, 0x0a, 0x10, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x15, 0x0a, 0x11, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x4f,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x54, 0x79, 0x70, 0x65, 0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09, 0x4f,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x4b, 0x79, 0x63, 0x10, 0x0a, 0x12, 0x14, 0x0a, 0x10, 0x4f, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x57, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x61, 0x6c, 0x10, 0x14,
	0x32, 0x09, 0x0a, 0x07, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x42, 0x36, 0x5a, 0x34, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4e, 0x70, 0x6f, 0x6f, 0x6c, 0x50,
	0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f,
	0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x2f, 0x6d, 0x67, 0x72,
	0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_npool_review_mgr_v2_mgr_proto_rawDescOnce sync.Once
	file_npool_review_mgr_v2_mgr_proto_rawDescData = file_npool_review_mgr_v2_mgr_proto_rawDesc
)

func file_npool_review_mgr_v2_mgr_proto_rawDescGZIP() []byte {
	file_npool_review_mgr_v2_mgr_proto_rawDescOnce.Do(func() {
		file_npool_review_mgr_v2_mgr_proto_rawDescData = protoimpl.X.CompressGZIP(file_npool_review_mgr_v2_mgr_proto_rawDescData)
	})
	return file_npool_review_mgr_v2_mgr_proto_rawDescData
}

var file_npool_review_mgr_v2_mgr_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_npool_review_mgr_v2_mgr_proto_goTypes = []interface{}{
	(ReviewObjectType)(0), // 0: review.manager.v1.ReviewObjectType
}
var file_npool_review_mgr_v2_mgr_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_npool_review_mgr_v2_mgr_proto_init() }
func file_npool_review_mgr_v2_mgr_proto_init() {
	if File_npool_review_mgr_v2_mgr_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_npool_review_mgr_v2_mgr_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_npool_review_mgr_v2_mgr_proto_goTypes,
		DependencyIndexes: file_npool_review_mgr_v2_mgr_proto_depIdxs,
		EnumInfos:         file_npool_review_mgr_v2_mgr_proto_enumTypes,
	}.Build()
	File_npool_review_mgr_v2_mgr_proto = out.File
	file_npool_review_mgr_v2_mgr_proto_rawDesc = nil
	file_npool_review_mgr_v2_mgr_proto_goTypes = nil
	file_npool_review_mgr_v2_mgr_proto_depIdxs = nil
}
