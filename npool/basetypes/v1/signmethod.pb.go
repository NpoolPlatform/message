// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.18.1
// source: npool/basetypes/v1/signmethod.proto

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

type SignMethod int32

const (
	SignMethod_DefaultSignMethod SignMethod = 0
	SignMethod_Mobile            SignMethod = 10 // For signup, signin, verify
	SignMethod_Email             SignMethod = 20 // For signup, signin, verify
	SignMethod_Twitter           SignMethod = 30 // For signin
	SignMethod_Github            SignMethod = 40 // For signin
	SignMethod_Facebook          SignMethod = 50 // For signin
	SignMethod_Linkedin          SignMethod = 60 // For signin
	SignMethod_Wechat            SignMethod = 70 // For signin
	SignMethod_Google            SignMethod = 80 // For signin, verify
	SignMethod_Username          SignMethod = 90 // For signup, signin
)

// Enum value maps for SignMethod.
var (
	SignMethod_name = map[int32]string{
		0:  "DefaultSignMethod",
		10: "Mobile",
		20: "Email",
		30: "Twitter",
		40: "Github",
		50: "Facebook",
		60: "Linkedin",
		70: "Wechat",
		80: "Google",
		90: "Username",
	}
	SignMethod_value = map[string]int32{
		"DefaultSignMethod": 0,
		"Mobile":            10,
		"Email":             20,
		"Twitter":           30,
		"Github":            40,
		"Facebook":          50,
		"Linkedin":          60,
		"Wechat":            70,
		"Google":            80,
		"Username":          90,
	}
)

func (x SignMethod) Enum() *SignMethod {
	p := new(SignMethod)
	*p = x
	return p
}

func (x SignMethod) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SignMethod) Descriptor() protoreflect.EnumDescriptor {
	return file_npool_basetypes_v1_signmethod_proto_enumTypes[0].Descriptor()
}

func (SignMethod) Type() protoreflect.EnumType {
	return &file_npool_basetypes_v1_signmethod_proto_enumTypes[0]
}

func (x SignMethod) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SignMethod.Descriptor instead.
func (SignMethod) EnumDescriptor() ([]byte, []int) {
	return file_npool_basetypes_v1_signmethod_proto_rawDescGZIP(), []int{0}
}

var File_npool_basetypes_v1_signmethod_proto protoreflect.FileDescriptor

var file_npool_basetypes_v1_signmethod_proto_rawDesc = []byte{
	0x0a, 0x23, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2e, 0x62, 0x61, 0x73,
	0x65, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2a, 0x95, 0x01, 0x0a, 0x0a, 0x53, 0x69,
	0x67, 0x6e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x15, 0x0a, 0x11, 0x44, 0x65, 0x66, 0x61,
	0x75, 0x6c, 0x74, 0x53, 0x69, 0x67, 0x6e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x10, 0x00, 0x12,
	0x0a, 0x0a, 0x06, 0x4d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x10, 0x0a, 0x12, 0x09, 0x0a, 0x05, 0x45,
	0x6d, 0x61, 0x69, 0x6c, 0x10, 0x14, 0x12, 0x0b, 0x0a, 0x07, 0x54, 0x77, 0x69, 0x74, 0x74, 0x65,
	0x72, 0x10, 0x1e, 0x12, 0x0a, 0x0a, 0x06, 0x47, 0x69, 0x74, 0x68, 0x75, 0x62, 0x10, 0x28, 0x12,
	0x0c, 0x0a, 0x08, 0x46, 0x61, 0x63, 0x65, 0x62, 0x6f, 0x6f, 0x6b, 0x10, 0x32, 0x12, 0x0c, 0x0a,
	0x08, 0x4c, 0x69, 0x6e, 0x6b, 0x65, 0x64, 0x69, 0x6e, 0x10, 0x3c, 0x12, 0x0a, 0x0a, 0x06, 0x57,
	0x65, 0x63, 0x68, 0x61, 0x74, 0x10, 0x46, 0x12, 0x0a, 0x0a, 0x06, 0x47, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x10, 0x50, 0x12, 0x0c, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x10,
	0x5a, 0x42, 0x35, 0x5a, 0x33, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x4e, 0x70, 0x6f, 0x6f, 0x6c, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x62, 0x61, 0x73, 0x65,
	0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_npool_basetypes_v1_signmethod_proto_rawDescOnce sync.Once
	file_npool_basetypes_v1_signmethod_proto_rawDescData = file_npool_basetypes_v1_signmethod_proto_rawDesc
)

func file_npool_basetypes_v1_signmethod_proto_rawDescGZIP() []byte {
	file_npool_basetypes_v1_signmethod_proto_rawDescOnce.Do(func() {
		file_npool_basetypes_v1_signmethod_proto_rawDescData = protoimpl.X.CompressGZIP(file_npool_basetypes_v1_signmethod_proto_rawDescData)
	})
	return file_npool_basetypes_v1_signmethod_proto_rawDescData
}

var file_npool_basetypes_v1_signmethod_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_npool_basetypes_v1_signmethod_proto_goTypes = []interface{}{
	(SignMethod)(0), // 0: npool.basetypes.v1.SignMethod
}
var file_npool_basetypes_v1_signmethod_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_npool_basetypes_v1_signmethod_proto_init() }
func file_npool_basetypes_v1_signmethod_proto_init() {
	if File_npool_basetypes_v1_signmethod_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_npool_basetypes_v1_signmethod_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_npool_basetypes_v1_signmethod_proto_goTypes,
		DependencyIndexes: file_npool_basetypes_v1_signmethod_proto_depIdxs,
		EnumInfos:         file_npool_basetypes_v1_signmethod_proto_enumTypes,
	}.Build()
	File_npool_basetypes_v1_signmethod_proto = out.File
	file_npool_basetypes_v1_signmethod_proto_rawDesc = nil
	file_npool_basetypes_v1_signmethod_proto_goTypes = nil
	file_npool_basetypes_v1_signmethod_proto_depIdxs = nil
}
