// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.18.1
// source: npool/basetypes/v1/gastype.proto

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

type GasType int32

const (
	GasType_DefaultGasType GasType = 0
	GasType_FixedGas       GasType = 10
	GasType_DynamicGas     GasType = 20
	GasType_GasUnsupported GasType = 30
)

// Enum value maps for GasType.
var (
	GasType_name = map[int32]string{
		0:  "DefaultGasType",
		10: "FixedGas",
		20: "DynamicGas",
		30: "GasUnsupported",
	}
	GasType_value = map[string]int32{
		"DefaultGasType": 0,
		"FixedGas":       10,
		"DynamicGas":     20,
		"GasUnsupported": 30,
	}
)

func (x GasType) Enum() *GasType {
	p := new(GasType)
	*p = x
	return p
}

func (x GasType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (GasType) Descriptor() protoreflect.EnumDescriptor {
	return file_npool_basetypes_v1_gastype_proto_enumTypes[0].Descriptor()
}

func (GasType) Type() protoreflect.EnumType {
	return &file_npool_basetypes_v1_gastype_proto_enumTypes[0]
}

func (x GasType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use GasType.Descriptor instead.
func (GasType) EnumDescriptor() ([]byte, []int) {
	return file_npool_basetypes_v1_gastype_proto_rawDescGZIP(), []int{0}
}

var File_npool_basetypes_v1_gastype_proto protoreflect.FileDescriptor

var file_npool_basetypes_v1_gastype_proto_rawDesc = []byte{
	0x0a, 0x20, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x61, 0x73, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0c, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31,
	0x2a, 0x4f, 0x0a, 0x07, 0x47, 0x61, 0x73, 0x54, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x0e, 0x44,
	0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x47, 0x61, 0x73, 0x54, 0x79, 0x70, 0x65, 0x10, 0x00, 0x12,
	0x0c, 0x0a, 0x08, 0x46, 0x69, 0x78, 0x65, 0x64, 0x47, 0x61, 0x73, 0x10, 0x0a, 0x12, 0x0e, 0x0a,
	0x0a, 0x44, 0x79, 0x6e, 0x61, 0x6d, 0x69, 0x63, 0x47, 0x61, 0x73, 0x10, 0x14, 0x12, 0x12, 0x0a,
	0x0e, 0x47, 0x61, 0x73, 0x55, 0x6e, 0x73, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x64, 0x10,
	0x1e, 0x42, 0x35, 0x5a, 0x33, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x4e, 0x70, 0x6f, 0x6f, 0x6c, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x62, 0x61, 0x73, 0x65,
	0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_npool_basetypes_v1_gastype_proto_rawDescOnce sync.Once
	file_npool_basetypes_v1_gastype_proto_rawDescData = file_npool_basetypes_v1_gastype_proto_rawDesc
)

func file_npool_basetypes_v1_gastype_proto_rawDescGZIP() []byte {
	file_npool_basetypes_v1_gastype_proto_rawDescOnce.Do(func() {
		file_npool_basetypes_v1_gastype_proto_rawDescData = protoimpl.X.CompressGZIP(file_npool_basetypes_v1_gastype_proto_rawDescData)
	})
	return file_npool_basetypes_v1_gastype_proto_rawDescData
}

var file_npool_basetypes_v1_gastype_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_npool_basetypes_v1_gastype_proto_goTypes = []interface{}{
	(GasType)(0), // 0: basetypes.v1.GasType
}
var file_npool_basetypes_v1_gastype_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_npool_basetypes_v1_gastype_proto_init() }
func file_npool_basetypes_v1_gastype_proto_init() {
	if File_npool_basetypes_v1_gastype_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_npool_basetypes_v1_gastype_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_npool_basetypes_v1_gastype_proto_goTypes,
		DependencyIndexes: file_npool_basetypes_v1_gastype_proto_depIdxs,
		EnumInfos:         file_npool_basetypes_v1_gastype_proto_enumTypes,
	}.Build()
	File_npool_basetypes_v1_gastype_proto = out.File
	file_npool_basetypes_v1_gastype_proto_rawDesc = nil
	file_npool_basetypes_v1_gastype_proto_goTypes = nil
	file_npool_basetypes_v1_gastype_proto_depIdxs = nil
}
