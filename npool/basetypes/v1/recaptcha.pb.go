// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.18.1
// source: npool/basetypes/v1/recaptcha.proto

<<<<<<< HEAD
package v1
=======
package recaptcha
>>>>>>> Move some basetype to basetypes

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

<<<<<<< HEAD
type RecaptchaMethod int32

const (
	RecaptchaMethod_DefaultRecaptchaMethod RecaptchaMethod = 0
	RecaptchaMethod_GoogleRecaptchaV3      RecaptchaMethod = 10
)

// Enum value maps for RecaptchaMethod.
var (
	RecaptchaMethod_name = map[int32]string{
		0:  "DefaultRecaptchaMethod",
		10: "GoogleRecaptchaV3",
	}
	RecaptchaMethod_value = map[string]int32{
		"DefaultRecaptchaMethod": 0,
		"GoogleRecaptchaV3":      10,
	}
)

func (x RecaptchaMethod) Enum() *RecaptchaMethod {
	p := new(RecaptchaMethod)
=======
type RecaptchaType int32

const (
	RecaptchaType_DefaultRecaptchaType RecaptchaType = 0
	RecaptchaType_GoogleRecaptchaV3    RecaptchaType = 10
)

// Enum value maps for RecaptchaType.
var (
	RecaptchaType_name = map[int32]string{
		0:  "DefaultRecaptchaType",
		10: "GoogleRecaptchaV3",
	}
	RecaptchaType_value = map[string]int32{
		"DefaultRecaptchaType": 0,
		"GoogleRecaptchaV3":    10,
	}
)

func (x RecaptchaType) Enum() *RecaptchaType {
	p := new(RecaptchaType)
>>>>>>> Move some basetype to basetypes
	*p = x
	return p
}

<<<<<<< HEAD
func (x RecaptchaMethod) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RecaptchaMethod) Descriptor() protoreflect.EnumDescriptor {
	return file_npool_basetypes_v1_recaptcha_proto_enumTypes[0].Descriptor()
}

func (RecaptchaMethod) Type() protoreflect.EnumType {
	return &file_npool_basetypes_v1_recaptcha_proto_enumTypes[0]
}

func (x RecaptchaMethod) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RecaptchaMethod.Descriptor instead.
func (RecaptchaMethod) EnumDescriptor() ([]byte, []int) {
=======
func (x RecaptchaType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RecaptchaType) Descriptor() protoreflect.EnumDescriptor {
	return file_npool_basetypes_v1_recaptcha_proto_enumTypes[0].Descriptor()
}

func (RecaptchaType) Type() protoreflect.EnumType {
	return &file_npool_basetypes_v1_recaptcha_proto_enumTypes[0]
}

func (x RecaptchaType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RecaptchaType.Descriptor instead.
func (RecaptchaType) EnumDescriptor() ([]byte, []int) {
>>>>>>> Move some basetype to basetypes
	return file_npool_basetypes_v1_recaptcha_proto_rawDescGZIP(), []int{0}
}

var File_npool_basetypes_v1_recaptcha_proto protoreflect.FileDescriptor

var file_npool_basetypes_v1_recaptcha_proto_rawDesc = []byte{
	0x0a, 0x22, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x63, 0x61, 0x70, 0x74, 0x63, 0x68, 0x61, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e,
<<<<<<< HEAD
	0x76, 0x31, 0x2a, 0x44, 0x0a, 0x0f, 0x52, 0x65, 0x63, 0x61, 0x70, 0x74, 0x63, 0x68, 0x61, 0x4d,
	0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x1a, 0x0a, 0x16, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74,
	0x52, 0x65, 0x63, 0x61, 0x70, 0x74, 0x63, 0x68, 0x61, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x10,
	0x00, 0x12, 0x15, 0x0a, 0x11, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x52, 0x65, 0x63, 0x61, 0x70,
	0x74, 0x63, 0x68, 0x61, 0x56, 0x33, 0x10, 0x0a, 0x42, 0x35, 0x5a, 0x33, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4e, 0x70, 0x6f, 0x6f, 0x6c, 0x50, 0x6c, 0x61, 0x74,
	0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x6e, 0x70, 0x6f,
	0x6f, 0x6c, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
=======
	0x76, 0x31, 0x2a, 0x40, 0x0a, 0x0d, 0x52, 0x65, 0x63, 0x61, 0x70, 0x74, 0x63, 0x68, 0x61, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x14, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x52, 0x65,
	0x63, 0x61, 0x70, 0x74, 0x63, 0x68, 0x61, 0x54, 0x79, 0x70, 0x65, 0x10, 0x00, 0x12, 0x15, 0x0a,
	0x11, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x52, 0x65, 0x63, 0x61, 0x70, 0x74, 0x63, 0x68, 0x61,
	0x56, 0x33, 0x10, 0x0a, 0x42, 0x3f, 0x5a, 0x3d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x4e, 0x70, 0x6f, 0x6f, 0x6c, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d,
	0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x62,
	0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x63, 0x61,
	0x70, 0x74, 0x63, 0x68, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
>>>>>>> Move some basetype to basetypes
}

var (
	file_npool_basetypes_v1_recaptcha_proto_rawDescOnce sync.Once
	file_npool_basetypes_v1_recaptcha_proto_rawDescData = file_npool_basetypes_v1_recaptcha_proto_rawDesc
)

func file_npool_basetypes_v1_recaptcha_proto_rawDescGZIP() []byte {
	file_npool_basetypes_v1_recaptcha_proto_rawDescOnce.Do(func() {
		file_npool_basetypes_v1_recaptcha_proto_rawDescData = protoimpl.X.CompressGZIP(file_npool_basetypes_v1_recaptcha_proto_rawDescData)
	})
	return file_npool_basetypes_v1_recaptcha_proto_rawDescData
}

var file_npool_basetypes_v1_recaptcha_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_npool_basetypes_v1_recaptcha_proto_goTypes = []interface{}{
<<<<<<< HEAD
	(RecaptchaMethod)(0), // 0: basetypes.v1.RecaptchaMethod
=======
	(RecaptchaType)(0), // 0: basetypes.v1.RecaptchaType
>>>>>>> Move some basetype to basetypes
}
var file_npool_basetypes_v1_recaptcha_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_npool_basetypes_v1_recaptcha_proto_init() }
func file_npool_basetypes_v1_recaptcha_proto_init() {
	if File_npool_basetypes_v1_recaptcha_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_npool_basetypes_v1_recaptcha_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_npool_basetypes_v1_recaptcha_proto_goTypes,
		DependencyIndexes: file_npool_basetypes_v1_recaptcha_proto_depIdxs,
		EnumInfos:         file_npool_basetypes_v1_recaptcha_proto_enumTypes,
	}.Build()
	File_npool_basetypes_v1_recaptcha_proto = out.File
	file_npool_basetypes_v1_recaptcha_proto_rawDesc = nil
	file_npool_basetypes_v1_recaptcha_proto_goTypes = nil
	file_npool_basetypes_v1_recaptcha_proto_depIdxs = nil
}
