// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.18.1
// source: npool/basetypes/v1/createinvitationcodewhen.proto

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

type CreateInvitationCodeWhen int32

const (
	CreateInvitationCodeWhen_DefaultWhen  CreateInvitationCodeWhen = 0
	CreateInvitationCodeWhen_Registration CreateInvitationCodeWhen = 10
	CreateInvitationCodeWhen_SetToKol     CreateInvitationCodeWhen = 20
	CreateInvitationCodeWhen_HasPaidOrder CreateInvitationCodeWhen = 30
)

// Enum value maps for CreateInvitationCodeWhen.
var (
	CreateInvitationCodeWhen_name = map[int32]string{
		0:  "DefaultWhen",
		10: "Registration",
		20: "SetToKol",
		30: "HasPaidOrder",
	}
	CreateInvitationCodeWhen_value = map[string]int32{
		"DefaultWhen":  0,
		"Registration": 10,
		"SetToKol":     20,
		"HasPaidOrder": 30,
	}
)

func (x CreateInvitationCodeWhen) Enum() *CreateInvitationCodeWhen {
	p := new(CreateInvitationCodeWhen)
	*p = x
	return p
}

func (x CreateInvitationCodeWhen) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CreateInvitationCodeWhen) Descriptor() protoreflect.EnumDescriptor {
	return file_npool_basetypes_v1_createinvitationcodewhen_proto_enumTypes[0].Descriptor()
}

func (CreateInvitationCodeWhen) Type() protoreflect.EnumType {
	return &file_npool_basetypes_v1_createinvitationcodewhen_proto_enumTypes[0]
}

func (x CreateInvitationCodeWhen) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CreateInvitationCodeWhen.Descriptor instead.
func (CreateInvitationCodeWhen) EnumDescriptor() ([]byte, []int) {
	return file_npool_basetypes_v1_createinvitationcodewhen_proto_rawDescGZIP(), []int{0}
}

var File_npool_basetypes_v1_createinvitationcodewhen_proto protoreflect.FileDescriptor

var file_npool_basetypes_v1_createinvitationcodewhen_proto_rawDesc = []byte{
	0x0a, 0x31, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x69, 0x6e, 0x76, 0x69, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x63, 0x6f, 0x64, 0x65, 0x77, 0x68, 0x65, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76,
	0x31, 0x2a, 0x5d, 0x0a, 0x18, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x76, 0x69, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x64, 0x65, 0x57, 0x68, 0x65, 0x6e, 0x12, 0x0f, 0x0a,
	0x0b, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x57, 0x68, 0x65, 0x6e, 0x10, 0x00, 0x12, 0x10,
	0x0a, 0x0c, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x10, 0x0a,
	0x12, 0x0c, 0x0a, 0x08, 0x53, 0x65, 0x74, 0x54, 0x6f, 0x4b, 0x6f, 0x6c, 0x10, 0x14, 0x12, 0x10,
	0x0a, 0x0c, 0x48, 0x61, 0x73, 0x50, 0x61, 0x69, 0x64, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x10, 0x1e,
	0x42, 0x35, 0x5a, 0x33, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4e,
	0x70, 0x6f, 0x6f, 0x6c, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x2f, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_npool_basetypes_v1_createinvitationcodewhen_proto_rawDescOnce sync.Once
	file_npool_basetypes_v1_createinvitationcodewhen_proto_rawDescData = file_npool_basetypes_v1_createinvitationcodewhen_proto_rawDesc
)

func file_npool_basetypes_v1_createinvitationcodewhen_proto_rawDescGZIP() []byte {
	file_npool_basetypes_v1_createinvitationcodewhen_proto_rawDescOnce.Do(func() {
		file_npool_basetypes_v1_createinvitationcodewhen_proto_rawDescData = protoimpl.X.CompressGZIP(file_npool_basetypes_v1_createinvitationcodewhen_proto_rawDescData)
	})
	return file_npool_basetypes_v1_createinvitationcodewhen_proto_rawDescData
}

var file_npool_basetypes_v1_createinvitationcodewhen_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_npool_basetypes_v1_createinvitationcodewhen_proto_goTypes = []interface{}{
	(CreateInvitationCodeWhen)(0), // 0: basetypes.v1.CreateInvitationCodeWhen
}
var file_npool_basetypes_v1_createinvitationcodewhen_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_npool_basetypes_v1_createinvitationcodewhen_proto_init() }
func file_npool_basetypes_v1_createinvitationcodewhen_proto_init() {
	if File_npool_basetypes_v1_createinvitationcodewhen_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_npool_basetypes_v1_createinvitationcodewhen_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_npool_basetypes_v1_createinvitationcodewhen_proto_goTypes,
		DependencyIndexes: file_npool_basetypes_v1_createinvitationcodewhen_proto_depIdxs,
		EnumInfos:         file_npool_basetypes_v1_createinvitationcodewhen_proto_enumTypes,
	}.Build()
	File_npool_basetypes_v1_createinvitationcodewhen_proto = out.File
	file_npool_basetypes_v1_createinvitationcodewhen_proto_rawDesc = nil
	file_npool_basetypes_v1_createinvitationcodewhen_proto_goTypes = nil
	file_npool_basetypes_v1_createinvitationcodewhen_proto_depIdxs = nil
}
