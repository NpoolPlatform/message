// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.18.1
// source: npool/basetypes/v1/result.proto

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

type Result int32

const (
	Result_DefaultResult Result = 0
	Result_Pass          Result = 10
	Result_Fail          Result = 20
)

// Enum value maps for Result.
var (
	Result_name = map[int32]string{
		0:  "DefaultResult",
		10: "Pass",
		20: "Fail",
	}
	Result_value = map[string]int32{
		"DefaultResult": 0,
		"Pass":          10,
		"Fail":          20,
	}
)

func (x Result) Enum() *Result {
	p := new(Result)
	*p = x
	return p
}

func (x Result) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Result) Descriptor() protoreflect.EnumDescriptor {
	return file_npool_basetypes_v1_result_proto_enumTypes[0].Descriptor()
}

func (Result) Type() protoreflect.EnumType {
	return &file_npool_basetypes_v1_result_proto_enumTypes[0]
}

func (x Result) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Result.Descriptor instead.
func (Result) EnumDescriptor() ([]byte, []int) {
	return file_npool_basetypes_v1_result_proto_rawDescGZIP(), []int{0}
}

var File_npool_basetypes_v1_result_proto protoreflect.FileDescriptor

var file_npool_basetypes_v1_result_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0c, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2a,
	0x2f, 0x0a, 0x06, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x11, 0x0a, 0x0d, 0x44, 0x65, 0x66,
	0x61, 0x75, 0x6c, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04,
	0x50, 0x61, 0x73, 0x73, 0x10, 0x0a, 0x12, 0x08, 0x0a, 0x04, 0x46, 0x61, 0x69, 0x6c, 0x10, 0x14,
	0x42, 0x35, 0x5a, 0x33, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4e,
	0x70, 0x6f, 0x6f, 0x6c, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x2f, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_npool_basetypes_v1_result_proto_rawDescOnce sync.Once
	file_npool_basetypes_v1_result_proto_rawDescData = file_npool_basetypes_v1_result_proto_rawDesc
)

func file_npool_basetypes_v1_result_proto_rawDescGZIP() []byte {
	file_npool_basetypes_v1_result_proto_rawDescOnce.Do(func() {
		file_npool_basetypes_v1_result_proto_rawDescData = protoimpl.X.CompressGZIP(file_npool_basetypes_v1_result_proto_rawDescData)
	})
	return file_npool_basetypes_v1_result_proto_rawDescData
}

var file_npool_basetypes_v1_result_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_npool_basetypes_v1_result_proto_goTypes = []interface{}{
	(Result)(0), // 0: basetypes.v1.Result
}
var file_npool_basetypes_v1_result_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_npool_basetypes_v1_result_proto_init() }
func file_npool_basetypes_v1_result_proto_init() {
	if File_npool_basetypes_v1_result_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_npool_basetypes_v1_result_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_npool_basetypes_v1_result_proto_goTypes,
		DependencyIndexes: file_npool_basetypes_v1_result_proto_depIdxs,
		EnumInfos:         file_npool_basetypes_v1_result_proto_enumTypes,
	}.Build()
	File_npool_basetypes_v1_result_proto = out.File
	file_npool_basetypes_v1_result_proto_rawDesc = nil
	file_npool_basetypes_v1_result_proto_goTypes = nil
	file_npool_basetypes_v1_result_proto_depIdxs = nil
}
