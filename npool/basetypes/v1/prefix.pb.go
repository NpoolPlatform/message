// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.18.1
// source: npool/basetypes/v1/prefix.proto

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

type Prefix int32

const (
	Prefix_DefaultPrefix                  Prefix = 0
	Prefix_PrefixUserCode                 Prefix = 10
	Prefix_PrefixUserAccount              Prefix = 20
	Prefix_PrefixUserLogin                Prefix = 30
	Prefix_PrefixWalletAccountLock        Prefix = 40
	Prefix_PrefixAPIRegister              Prefix = 50
	Prefix_PrefixCreateCoin               Prefix = 60
	Prefix_PrefixCreateAppCoinDescription Prefix = 70
	Prefix_PrefixCreateAppCoin            Prefix = 80
	Prefix_PrefixSetFiat                  Prefix = 90
	Prefix_PrefixCreateFiatCurrencyFeed   Prefix = 100
	Prefix_PrefixCreateFiatCurrency       Prefix = 110
	Prefix_PrefixCreateCoinCurrencyFeed   Prefix = 120
	Prefix_PrefixCreateCoinCurrency       Prefix = 130
	Prefix_PrefixCreateCoinFiat           Prefix = 140
	Prefix_PrefixCreateCoinFiatCurrency   Prefix = 150
	Prefix_PrefixCreateAppCountry         Prefix = 160
	Prefix_PrefixCreateCountry            Prefix = 170
	Prefix_PrefixCreateAppLang            Prefix = 180
	Prefix_PrefixCreateLang               Prefix = 190
	Prefix_PrefixCreateMessage            Prefix = 200
	Prefix_PrefixCreateNotifUser          Prefix = 210
	Prefix_PrefixCreateNotif              Prefix = 220
	Prefix_PrefixCreateEmailTemplate      Prefix = 230
	Prefix_PrefixCreateFrontendTemplate   Prefix = 240
	Prefix_PrefixCreateSMSTemplate        Prefix = 250
)

// Enum value maps for Prefix.
var (
	Prefix_name = map[int32]string{
		0:   "DefaultPrefix",
		10:  "PrefixUserCode",
		20:  "PrefixUserAccount",
		30:  "PrefixUserLogin",
		40:  "PrefixWalletAccountLock",
		50:  "PrefixAPIRegister",
		60:  "PrefixCreateCoin",
		70:  "PrefixCreateAppCoinDescription",
		80:  "PrefixCreateAppCoin",
		90:  "PrefixSetFiat",
		100: "PrefixCreateFiatCurrencyFeed",
		110: "PrefixCreateFiatCurrency",
		120: "PrefixCreateCoinCurrencyFeed",
		130: "PrefixCreateCoinCurrency",
		140: "PrefixCreateCoinFiat",
		150: "PrefixCreateCoinFiatCurrency",
		160: "PrefixCreateAppCountry",
		170: "PrefixCreateCountry",
		180: "PrefixCreateAppLang",
		190: "PrefixCreateLang",
		200: "PrefixCreateMessage",
		210: "PrefixCreateNotifUser",
		220: "PrefixCreateNotif",
		230: "PrefixCreateEmailTemplate",
		240: "PrefixCreateFrontendTemplate",
		250: "PrefixCreateSMSTemplate",
	}
	Prefix_value = map[string]int32{
		"DefaultPrefix":                  0,
		"PrefixUserCode":                 10,
		"PrefixUserAccount":              20,
		"PrefixUserLogin":                30,
		"PrefixWalletAccountLock":        40,
		"PrefixAPIRegister":              50,
		"PrefixCreateCoin":               60,
		"PrefixCreateAppCoinDescription": 70,
		"PrefixCreateAppCoin":            80,
		"PrefixSetFiat":                  90,
		"PrefixCreateFiatCurrencyFeed":   100,
		"PrefixCreateFiatCurrency":       110,
		"PrefixCreateCoinCurrencyFeed":   120,
		"PrefixCreateCoinCurrency":       130,
		"PrefixCreateCoinFiat":           140,
		"PrefixCreateCoinFiatCurrency":   150,
		"PrefixCreateAppCountry":         160,
		"PrefixCreateCountry":            170,
		"PrefixCreateAppLang":            180,
		"PrefixCreateLang":               190,
		"PrefixCreateMessage":            200,
		"PrefixCreateNotifUser":          210,
		"PrefixCreateNotif":              220,
		"PrefixCreateEmailTemplate":      230,
		"PrefixCreateFrontendTemplate":   240,
		"PrefixCreateSMSTemplate":        250,
	}
)

func (x Prefix) Enum() *Prefix {
	p := new(Prefix)
	*p = x
	return p
}

func (x Prefix) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Prefix) Descriptor() protoreflect.EnumDescriptor {
	return file_npool_basetypes_v1_prefix_proto_enumTypes[0].Descriptor()
}

func (Prefix) Type() protoreflect.EnumType {
	return &file_npool_basetypes_v1_prefix_proto_enumTypes[0]
}

func (x Prefix) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Prefix.Descriptor instead.
func (Prefix) EnumDescriptor() ([]byte, []int) {
	return file_npool_basetypes_v1_prefix_proto_rawDescGZIP(), []int{0}
}

var File_npool_basetypes_v1_prefix_proto protoreflect.FileDescriptor

var file_npool_basetypes_v1_prefix_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0c, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2a,
	0xcb, 0x05, 0x0a, 0x06, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x12, 0x11, 0x0a, 0x0d, 0x44, 0x65,
	0x66, 0x61, 0x75, 0x6c, 0x74, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x10, 0x00, 0x12, 0x12, 0x0a,
	0x0e, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x55, 0x73, 0x65, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x10,
	0x0a, 0x12, 0x15, 0x0a, 0x11, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x55, 0x73, 0x65, 0x72, 0x41,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x10, 0x14, 0x12, 0x13, 0x0a, 0x0f, 0x50, 0x72, 0x65, 0x66,
	0x69, 0x78, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x10, 0x1e, 0x12, 0x1b, 0x0a,
	0x17, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x41, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x4c, 0x6f, 0x63, 0x6b, 0x10, 0x28, 0x12, 0x15, 0x0a, 0x11, 0x50, 0x72,
	0x65, 0x66, 0x69, 0x78, 0x41, 0x50, 0x49, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x10,
	0x32, 0x12, 0x14, 0x0a, 0x10, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x43, 0x6f, 0x69, 0x6e, 0x10, 0x3c, 0x12, 0x22, 0x0a, 0x1e, 0x50, 0x72, 0x65, 0x66, 0x69,
	0x78, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x70, 0x70, 0x43, 0x6f, 0x69, 0x6e, 0x44, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x10, 0x46, 0x12, 0x17, 0x0a, 0x13, 0x50,
	0x72, 0x65, 0x66, 0x69, 0x78, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x70, 0x70, 0x43, 0x6f,
	0x69, 0x6e, 0x10, 0x50, 0x12, 0x11, 0x0a, 0x0d, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x53, 0x65,
	0x74, 0x46, 0x69, 0x61, 0x74, 0x10, 0x5a, 0x12, 0x20, 0x0a, 0x1c, 0x50, 0x72, 0x65, 0x66, 0x69,
	0x78, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x46, 0x69, 0x61, 0x74, 0x43, 0x75, 0x72, 0x72, 0x65,
	0x6e, 0x63, 0x79, 0x46, 0x65, 0x65, 0x64, 0x10, 0x64, 0x12, 0x1c, 0x0a, 0x18, 0x50, 0x72, 0x65,
	0x66, 0x69, 0x78, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x46, 0x69, 0x61, 0x74, 0x43, 0x75, 0x72,
	0x72, 0x65, 0x6e, 0x63, 0x79, 0x10, 0x6e, 0x12, 0x20, 0x0a, 0x1c, 0x50, 0x72, 0x65, 0x66, 0x69,
	0x78, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x69, 0x6e, 0x43, 0x75, 0x72, 0x72, 0x65,
	0x6e, 0x63, 0x79, 0x46, 0x65, 0x65, 0x64, 0x10, 0x78, 0x12, 0x1d, 0x0a, 0x18, 0x50, 0x72, 0x65,
	0x66, 0x69, 0x78, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x69, 0x6e, 0x43, 0x75, 0x72,
	0x72, 0x65, 0x6e, 0x63, 0x79, 0x10, 0x82, 0x01, 0x12, 0x19, 0x0a, 0x14, 0x50, 0x72, 0x65, 0x66,
	0x69, 0x78, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x69, 0x6e, 0x46, 0x69, 0x61, 0x74,
	0x10, 0x8c, 0x01, 0x12, 0x21, 0x0a, 0x1c, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x43, 0x6f, 0x69, 0x6e, 0x46, 0x69, 0x61, 0x74, 0x43, 0x75, 0x72, 0x72, 0x65,
	0x6e, 0x63, 0x79, 0x10, 0x96, 0x01, 0x12, 0x1b, 0x0a, 0x16, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x70, 0x70, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79,
	0x10, 0xa0, 0x01, 0x12, 0x18, 0x0a, 0x13, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x10, 0xaa, 0x01, 0x12, 0x18, 0x0a,
	0x13, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x70, 0x70,
	0x4c, 0x61, 0x6e, 0x67, 0x10, 0xb4, 0x01, 0x12, 0x15, 0x0a, 0x10, 0x50, 0x72, 0x65, 0x66, 0x69,
	0x78, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x61, 0x6e, 0x67, 0x10, 0xbe, 0x01, 0x12, 0x18,
	0x0a, 0x13, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x10, 0xc8, 0x01, 0x12, 0x1a, 0x0a, 0x15, 0x50, 0x72, 0x65, 0x66,
	0x69, 0x78, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x55, 0x73, 0x65,
	0x72, 0x10, 0xd2, 0x01, 0x12, 0x16, 0x0a, 0x11, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x10, 0xdc, 0x01, 0x12, 0x1e, 0x0a, 0x19,
	0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x45, 0x6d, 0x61, 0x69,
	0x6c, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x10, 0xe6, 0x01, 0x12, 0x21, 0x0a, 0x1c,
	0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x46, 0x72, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x64, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x10, 0xf0, 0x01, 0x12,
	0x1c, 0x0a, 0x17, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53,
	0x4d, 0x53, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x10, 0xfa, 0x01, 0x42, 0x35, 0x5a,
	0x33, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4e, 0x70, 0x6f, 0x6f,
	0x6c, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x2f, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_npool_basetypes_v1_prefix_proto_rawDescOnce sync.Once
	file_npool_basetypes_v1_prefix_proto_rawDescData = file_npool_basetypes_v1_prefix_proto_rawDesc
)

func file_npool_basetypes_v1_prefix_proto_rawDescGZIP() []byte {
	file_npool_basetypes_v1_prefix_proto_rawDescOnce.Do(func() {
		file_npool_basetypes_v1_prefix_proto_rawDescData = protoimpl.X.CompressGZIP(file_npool_basetypes_v1_prefix_proto_rawDescData)
	})
	return file_npool_basetypes_v1_prefix_proto_rawDescData
}

var file_npool_basetypes_v1_prefix_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_npool_basetypes_v1_prefix_proto_goTypes = []interface{}{
	(Prefix)(0), // 0: basetypes.v1.Prefix
}
var file_npool_basetypes_v1_prefix_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_npool_basetypes_v1_prefix_proto_init() }
func file_npool_basetypes_v1_prefix_proto_init() {
	if File_npool_basetypes_v1_prefix_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_npool_basetypes_v1_prefix_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_npool_basetypes_v1_prefix_proto_goTypes,
		DependencyIndexes: file_npool_basetypes_v1_prefix_proto_depIdxs,
		EnumInfos:         file_npool_basetypes_v1_prefix_proto_enumTypes,
	}.Build()
	File_npool_basetypes_v1_prefix_proto = out.File
	file_npool_basetypes_v1_prefix_proto_rawDesc = nil
	file_npool_basetypes_v1_prefix_proto_goTypes = nil
	file_npool_basetypes_v1_prefix_proto_depIdxs = nil
}
