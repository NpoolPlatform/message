// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v3.18.1
// source: npool/basetypes/v1/kycstate.proto

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

type KycDocumentType int32

const (
	KycDocumentType_DefaultKycDocumentType KycDocumentType = 0
	KycDocumentType_IDCard                 KycDocumentType = 10
	KycDocumentType_DriverLicense          KycDocumentType = 20
	KycDocumentType_Passport               KycDocumentType = 30
)

// Enum value maps for KycDocumentType.
var (
	KycDocumentType_name = map[int32]string{
		0:  "DefaultKycDocumentType",
		10: "IDCard",
		20: "DriverLicense",
		30: "Passport",
	}
	KycDocumentType_value = map[string]int32{
		"DefaultKycDocumentType": 0,
		"IDCard":                 10,
		"DriverLicense":          20,
		"Passport":               30,
	}
)

func (x KycDocumentType) Enum() *KycDocumentType {
	p := new(KycDocumentType)
	*p = x
	return p
}

func (x KycDocumentType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (KycDocumentType) Descriptor() protoreflect.EnumDescriptor {
	return file_npool_basetypes_v1_kycstate_proto_enumTypes[0].Descriptor()
}

func (KycDocumentType) Type() protoreflect.EnumType {
	return &file_npool_basetypes_v1_kycstate_proto_enumTypes[0]
}

func (x KycDocumentType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use KycDocumentType.Descriptor instead.
func (KycDocumentType) EnumDescriptor() ([]byte, []int) {
	return file_npool_basetypes_v1_kycstate_proto_rawDescGZIP(), []int{0}
}

type KycEntityType int32

const (
	KycEntityType_DefaultKycEntityType KycEntityType = 0
	KycEntityType_Individual           KycEntityType = 10
	KycEntityType_Organization         KycEntityType = 20
)

// Enum value maps for KycEntityType.
var (
	KycEntityType_name = map[int32]string{
		0:  "DefaultKycEntityType",
		10: "Individual",
		20: "Organization",
	}
	KycEntityType_value = map[string]int32{
		"DefaultKycEntityType": 0,
		"Individual":           10,
		"Organization":         20,
	}
)

func (x KycEntityType) Enum() *KycEntityType {
	p := new(KycEntityType)
	*p = x
	return p
}

func (x KycEntityType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (KycEntityType) Descriptor() protoreflect.EnumDescriptor {
	return file_npool_basetypes_v1_kycstate_proto_enumTypes[1].Descriptor()
}

func (KycEntityType) Type() protoreflect.EnumType {
	return &file_npool_basetypes_v1_kycstate_proto_enumTypes[1]
}

func (x KycEntityType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use KycEntityType.Descriptor instead.
func (KycEntityType) EnumDescriptor() ([]byte, []int) {
	return file_npool_basetypes_v1_kycstate_proto_rawDescGZIP(), []int{1}
}

type KycImageType int32

const (
	KycImageType_DefaultKycImageType KycImageType = 0
	KycImageType_FrontImg            KycImageType = 10
	KycImageType_BackImg             KycImageType = 20
	KycImageType_SelfieImg           KycImageType = 30
)

// Enum value maps for KycImageType.
var (
	KycImageType_name = map[int32]string{
		0:  "DefaultKycImageType",
		10: "FrontImg",
		20: "BackImg",
		30: "SelfieImg",
	}
	KycImageType_value = map[string]int32{
		"DefaultKycImageType": 0,
		"FrontImg":            10,
		"BackImg":             20,
		"SelfieImg":           30,
	}
)

func (x KycImageType) Enum() *KycImageType {
	p := new(KycImageType)
	*p = x
	return p
}

func (x KycImageType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (KycImageType) Descriptor() protoreflect.EnumDescriptor {
	return file_npool_basetypes_v1_kycstate_proto_enumTypes[2].Descriptor()
}

func (KycImageType) Type() protoreflect.EnumType {
	return &file_npool_basetypes_v1_kycstate_proto_enumTypes[2]
}

func (x KycImageType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use KycImageType.Descriptor instead.
func (KycImageType) EnumDescriptor() ([]byte, []int) {
	return file_npool_basetypes_v1_kycstate_proto_rawDescGZIP(), []int{2}
}

type KycState int32

const (
	KycState_DefaultState KycState = 0
	KycState_Approved     KycState = 10
	KycState_Reviewing    KycState = 20
	KycState_Rejected     KycState = 30
)

// Enum value maps for KycState.
var (
	KycState_name = map[int32]string{
		0:  "DefaultState",
		10: "Approved",
		20: "Reviewing",
		30: "Rejected",
	}
	KycState_value = map[string]int32{
		"DefaultState": 0,
		"Approved":     10,
		"Reviewing":    20,
		"Rejected":     30,
	}
)

func (x KycState) Enum() *KycState {
	p := new(KycState)
	*p = x
	return p
}

func (x KycState) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (KycState) Descriptor() protoreflect.EnumDescriptor {
	return file_npool_basetypes_v1_kycstate_proto_enumTypes[3].Descriptor()
}

func (KycState) Type() protoreflect.EnumType {
	return &file_npool_basetypes_v1_kycstate_proto_enumTypes[3]
}

func (x KycState) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use KycState.Descriptor instead.
func (KycState) EnumDescriptor() ([]byte, []int) {
	return file_npool_basetypes_v1_kycstate_proto_rawDescGZIP(), []int{3}
}

var File_npool_basetypes_v1_kycstate_proto protoreflect.FileDescriptor

var file_npool_basetypes_v1_kycstate_proto_rawDesc = []byte{
	0x0a, 0x21, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x2f, 0x76, 0x31, 0x2f, 0x6b, 0x79, 0x63, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76,
	0x31, 0x2a, 0x5a, 0x0a, 0x0f, 0x4b, 0x79, 0x63, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x16, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x4b,
	0x79, 0x63, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x10, 0x00,
	0x12, 0x0a, 0x0a, 0x06, 0x49, 0x44, 0x43, 0x61, 0x72, 0x64, 0x10, 0x0a, 0x12, 0x11, 0x0a, 0x0d,
	0x44, 0x72, 0x69, 0x76, 0x65, 0x72, 0x4c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x10, 0x14, 0x12,
	0x0c, 0x0a, 0x08, 0x50, 0x61, 0x73, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x10, 0x1e, 0x2a, 0x4b, 0x0a,
	0x0d, 0x4b, 0x79, 0x63, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x54, 0x79, 0x70, 0x65, 0x12, 0x18,
	0x0a, 0x14, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x4b, 0x79, 0x63, 0x45, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x54, 0x79, 0x70, 0x65, 0x10, 0x00, 0x12, 0x0e, 0x0a, 0x0a, 0x49, 0x6e, 0x64, 0x69,
	0x76, 0x69, 0x64, 0x75, 0x61, 0x6c, 0x10, 0x0a, 0x12, 0x10, 0x0a, 0x0c, 0x4f, 0x72, 0x67, 0x61,
	0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x10, 0x14, 0x2a, 0x51, 0x0a, 0x0c, 0x4b, 0x79,
	0x63, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x17, 0x0a, 0x13, 0x44, 0x65,
	0x66, 0x61, 0x75, 0x6c, 0x74, 0x4b, 0x79, 0x63, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70,
	0x65, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x46, 0x72, 0x6f, 0x6e, 0x74, 0x49, 0x6d, 0x67, 0x10,
	0x0a, 0x12, 0x0b, 0x0a, 0x07, 0x42, 0x61, 0x63, 0x6b, 0x49, 0x6d, 0x67, 0x10, 0x14, 0x12, 0x0d,
	0x0a, 0x09, 0x53, 0x65, 0x6c, 0x66, 0x69, 0x65, 0x49, 0x6d, 0x67, 0x10, 0x1e, 0x2a, 0x47, 0x0a,
	0x08, 0x4b, 0x79, 0x63, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x10, 0x0a, 0x0c, 0x44, 0x65, 0x66,
	0x61, 0x75, 0x6c, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x41,
	0x70, 0x70, 0x72, 0x6f, 0x76, 0x65, 0x64, 0x10, 0x0a, 0x12, 0x0d, 0x0a, 0x09, 0x52, 0x65, 0x76,
	0x69, 0x65, 0x77, 0x69, 0x6e, 0x67, 0x10, 0x14, 0x12, 0x0c, 0x0a, 0x08, 0x52, 0x65, 0x6a, 0x65,
	0x63, 0x74, 0x65, 0x64, 0x10, 0x1e, 0x42, 0x35, 0x5a, 0x33, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4e, 0x70, 0x6f, 0x6f, 0x6c, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f,
	0x72, 0x6d, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x6e, 0x70, 0x6f, 0x6f, 0x6c,
	0x2f, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_npool_basetypes_v1_kycstate_proto_rawDescOnce sync.Once
	file_npool_basetypes_v1_kycstate_proto_rawDescData = file_npool_basetypes_v1_kycstate_proto_rawDesc
)

func file_npool_basetypes_v1_kycstate_proto_rawDescGZIP() []byte {
	file_npool_basetypes_v1_kycstate_proto_rawDescOnce.Do(func() {
		file_npool_basetypes_v1_kycstate_proto_rawDescData = protoimpl.X.CompressGZIP(file_npool_basetypes_v1_kycstate_proto_rawDescData)
	})
	return file_npool_basetypes_v1_kycstate_proto_rawDescData
}

var file_npool_basetypes_v1_kycstate_proto_enumTypes = make([]protoimpl.EnumInfo, 4)
var file_npool_basetypes_v1_kycstate_proto_goTypes = []interface{}{
	(KycDocumentType)(0), // 0: basetypes.v1.KycDocumentType
	(KycEntityType)(0),   // 1: basetypes.v1.KycEntityType
	(KycImageType)(0),    // 2: basetypes.v1.KycImageType
	(KycState)(0),        // 3: basetypes.v1.KycState
}
var file_npool_basetypes_v1_kycstate_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_npool_basetypes_v1_kycstate_proto_init() }
func file_npool_basetypes_v1_kycstate_proto_init() {
	if File_npool_basetypes_v1_kycstate_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_npool_basetypes_v1_kycstate_proto_rawDesc,
			NumEnums:      4,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_npool_basetypes_v1_kycstate_proto_goTypes,
		DependencyIndexes: file_npool_basetypes_v1_kycstate_proto_depIdxs,
		EnumInfos:         file_npool_basetypes_v1_kycstate_proto_enumTypes,
	}.Build()
	File_npool_basetypes_v1_kycstate_proto = out.File
	file_npool_basetypes_v1_kycstate_proto_rawDesc = nil
	file_npool_basetypes_v1_kycstate_proto_goTypes = nil
	file_npool_basetypes_v1_kycstate_proto_depIdxs = nil
}
