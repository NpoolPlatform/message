// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.18.1
// source: npool/basetypes/miningpool/v1/enums.proto

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

type MiningpoolType int32

const (
	MiningpoolType_DefaultMiningpoolType MiningpoolType = 0
	MiningpoolType_F2Pool                MiningpoolType = 10
	MiningpoolType_AntPool               MiningpoolType = 20
)

// Enum value maps for MiningpoolType.
var (
	MiningpoolType_name = map[int32]string{
		0:  "DefaultMiningpoolType",
		10: "F2Pool",
		20: "AntPool",
	}
	MiningpoolType_value = map[string]int32{
		"DefaultMiningpoolType": 0,
		"F2Pool":                10,
		"AntPool":               20,
	}
)

func (x MiningpoolType) Enum() *MiningpoolType {
	p := new(MiningpoolType)
	*p = x
	return p
}

func (x MiningpoolType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MiningpoolType) Descriptor() protoreflect.EnumDescriptor {
	return file_npool_basetypes_miningpool_v1_enums_proto_enumTypes[0].Descriptor()
}

func (MiningpoolType) Type() protoreflect.EnumType {
	return &file_npool_basetypes_miningpool_v1_enums_proto_enumTypes[0]
}

func (x MiningpoolType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MiningpoolType.Descriptor instead.
func (MiningpoolType) EnumDescriptor() ([]byte, []int) {
	return file_npool_basetypes_miningpool_v1_enums_proto_rawDescGZIP(), []int{0}
}

type CoinType int32

const (
	CoinType_DefaultCoinType CoinType = 0
	CoinType_BitCoin         CoinType = 10
)

// Enum value maps for CoinType.
var (
	CoinType_name = map[int32]string{
		0:  "DefaultCoinType",
		10: "BitCoin",
	}
	CoinType_value = map[string]int32{
		"DefaultCoinType": 0,
		"BitCoin":         10,
	}
)

func (x CoinType) Enum() *CoinType {
	p := new(CoinType)
	*p = x
	return p
}

func (x CoinType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CoinType) Descriptor() protoreflect.EnumDescriptor {
	return file_npool_basetypes_miningpool_v1_enums_proto_enumTypes[1].Descriptor()
}

func (CoinType) Type() protoreflect.EnumType {
	return &file_npool_basetypes_miningpool_v1_enums_proto_enumTypes[1]
}

func (x CoinType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CoinType.Descriptor instead.
func (CoinType) EnumDescriptor() ([]byte, []int) {
	return file_npool_basetypes_miningpool_v1_enums_proto_rawDescGZIP(), []int{1}
}

type RevenueType int32

const (
	RevenueType_DefaultRevenueType RevenueType = 0
	RevenueType_FPPS               RevenueType = 10
	RevenueType_PPLNS              RevenueType = 20
)

// Enum value maps for RevenueType.
var (
	RevenueType_name = map[int32]string{
		0:  "DefaultRevenueType",
		10: "FPPS",
		20: "PPLNS",
	}
	RevenueType_value = map[string]int32{
		"DefaultRevenueType": 0,
		"FPPS":               10,
		"PPLNS":              20,
	}
)

func (x RevenueType) Enum() *RevenueType {
	p := new(RevenueType)
	*p = x
	return p
}

func (x RevenueType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RevenueType) Descriptor() protoreflect.EnumDescriptor {
	return file_npool_basetypes_miningpool_v1_enums_proto_enumTypes[2].Descriptor()
}

func (RevenueType) Type() protoreflect.EnumType {
	return &file_npool_basetypes_miningpool_v1_enums_proto_enumTypes[2]
}

func (x RevenueType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RevenueType.Descriptor instead.
func (RevenueType) EnumDescriptor() ([]byte, []int) {
	return file_npool_basetypes_miningpool_v1_enums_proto_rawDescGZIP(), []int{2}
}

type WithdrawState int32

const (
	WithdrawState_DefaultWithdrawState     WithdrawState = 0
	WithdrawState_WithdrawStateWait        WithdrawState = 10
	WithdrawState_WithdrawStateProccessing WithdrawState = 20
	WithdrawState_WithdrawStateSuccess     WithdrawState = 30
	WithdrawState_WithdrawStateFailed      WithdrawState = 40
)

// Enum value maps for WithdrawState.
var (
	WithdrawState_name = map[int32]string{
		0:  "DefaultWithdrawState",
		10: "WithdrawStateWait",
		20: "WithdrawStateProccessing",
		30: "WithdrawStateSuccess",
		40: "WithdrawStateFailed",
	}
	WithdrawState_value = map[string]int32{
		"DefaultWithdrawState":     0,
		"WithdrawStateWait":        10,
		"WithdrawStateProccessing": 20,
		"WithdrawStateSuccess":     30,
		"WithdrawStateFailed":      40,
	}
)

func (x WithdrawState) Enum() *WithdrawState {
	p := new(WithdrawState)
	*p = x
	return p
}

func (x WithdrawState) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (WithdrawState) Descriptor() protoreflect.EnumDescriptor {
	return file_npool_basetypes_miningpool_v1_enums_proto_enumTypes[3].Descriptor()
}

func (WithdrawState) Type() protoreflect.EnumType {
	return &file_npool_basetypes_miningpool_v1_enums_proto_enumTypes[3]
}

func (x WithdrawState) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use WithdrawState.Descriptor instead.
func (WithdrawState) EnumDescriptor() ([]byte, []int) {
	return file_npool_basetypes_miningpool_v1_enums_proto_rawDescGZIP(), []int{3}
}

var File_npool_basetypes_miningpool_v1_enums_proto protoreflect.FileDescriptor

var file_npool_basetypes_miningpool_v1_enums_proto_rawDesc = []byte{
	0x0a, 0x29, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x2f, 0x6d, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x76, 0x31, 0x2f,
	0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x17, 0x62, 0x61, 0x73,
	0x65, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x6d, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x70, 0x6f, 0x6f,
	0x6c, 0x2e, 0x76, 0x31, 0x2a, 0x44, 0x0a, 0x0e, 0x4d, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x70, 0x6f,
	0x6f, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x12, 0x19, 0x0a, 0x15, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c,
	0x74, 0x4d, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x70, 0x6f, 0x6f, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x10,
	0x00, 0x12, 0x0a, 0x0a, 0x06, 0x46, 0x32, 0x50, 0x6f, 0x6f, 0x6c, 0x10, 0x0a, 0x12, 0x0b, 0x0a,
	0x07, 0x41, 0x6e, 0x74, 0x50, 0x6f, 0x6f, 0x6c, 0x10, 0x14, 0x2a, 0x2c, 0x0a, 0x08, 0x43, 0x6f,
	0x69, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x13, 0x0a, 0x0f, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c,
	0x74, 0x43, 0x6f, 0x69, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x42,
	0x69, 0x74, 0x43, 0x6f, 0x69, 0x6e, 0x10, 0x0a, 0x2a, 0x3a, 0x0a, 0x0b, 0x52, 0x65, 0x76, 0x65,
	0x6e, 0x75, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x12, 0x44, 0x65, 0x66, 0x61, 0x75,
	0x6c, 0x74, 0x52, 0x65, 0x76, 0x65, 0x6e, 0x75, 0x65, 0x54, 0x79, 0x70, 0x65, 0x10, 0x00, 0x12,
	0x08, 0x0a, 0x04, 0x46, 0x50, 0x50, 0x53, 0x10, 0x0a, 0x12, 0x09, 0x0a, 0x05, 0x50, 0x50, 0x4c,
	0x4e, 0x53, 0x10, 0x14, 0x2a, 0x91, 0x01, 0x0a, 0x0d, 0x57, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61,
	0x77, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x18, 0x0a, 0x14, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c,
	0x74, 0x57, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x53, 0x74, 0x61, 0x74, 0x65, 0x10, 0x00,
	0x12, 0x15, 0x0a, 0x11, 0x57, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x53, 0x74, 0x61, 0x74,
	0x65, 0x57, 0x61, 0x69, 0x74, 0x10, 0x0a, 0x12, 0x1c, 0x0a, 0x18, 0x57, 0x69, 0x74, 0x68, 0x64,
	0x72, 0x61, 0x77, 0x53, 0x74, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x69, 0x6e, 0x67, 0x10, 0x14, 0x12, 0x18, 0x0a, 0x14, 0x57, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61,
	0x77, 0x53, 0x74, 0x61, 0x74, 0x65, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x10, 0x1e, 0x12,
	0x17, 0x0a, 0x13, 0x57, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x53, 0x74, 0x61, 0x74, 0x65,
	0x46, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x10, 0x28, 0x42, 0x40, 0x5a, 0x3e, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4e, 0x70, 0x6f, 0x6f, 0x6c, 0x50, 0x6c, 0x61, 0x74,
	0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x6e, 0x70, 0x6f,
	0x6f, 0x6c, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x6d, 0x69, 0x6e,
	0x69, 0x6e, 0x67, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_npool_basetypes_miningpool_v1_enums_proto_rawDescOnce sync.Once
	file_npool_basetypes_miningpool_v1_enums_proto_rawDescData = file_npool_basetypes_miningpool_v1_enums_proto_rawDesc
)

func file_npool_basetypes_miningpool_v1_enums_proto_rawDescGZIP() []byte {
	file_npool_basetypes_miningpool_v1_enums_proto_rawDescOnce.Do(func() {
		file_npool_basetypes_miningpool_v1_enums_proto_rawDescData = protoimpl.X.CompressGZIP(file_npool_basetypes_miningpool_v1_enums_proto_rawDescData)
	})
	return file_npool_basetypes_miningpool_v1_enums_proto_rawDescData
}

var file_npool_basetypes_miningpool_v1_enums_proto_enumTypes = make([]protoimpl.EnumInfo, 4)
var file_npool_basetypes_miningpool_v1_enums_proto_goTypes = []interface{}{
	(MiningpoolType)(0), // 0: basetypes.miningpool.v1.MiningpoolType
	(CoinType)(0),       // 1: basetypes.miningpool.v1.CoinType
	(RevenueType)(0),    // 2: basetypes.miningpool.v1.RevenueType
	(WithdrawState)(0),  // 3: basetypes.miningpool.v1.WithdrawState
}
var file_npool_basetypes_miningpool_v1_enums_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_npool_basetypes_miningpool_v1_enums_proto_init() }
func file_npool_basetypes_miningpool_v1_enums_proto_init() {
	if File_npool_basetypes_miningpool_v1_enums_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_npool_basetypes_miningpool_v1_enums_proto_rawDesc,
			NumEnums:      4,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_npool_basetypes_miningpool_v1_enums_proto_goTypes,
		DependencyIndexes: file_npool_basetypes_miningpool_v1_enums_proto_depIdxs,
		EnumInfos:         file_npool_basetypes_miningpool_v1_enums_proto_enumTypes,
	}.Build()
	File_npool_basetypes_miningpool_v1_enums_proto = out.File
	file_npool_basetypes_miningpool_v1_enums_proto_rawDesc = nil
	file_npool_basetypes_miningpool_v1_enums_proto_goTypes = nil
	file_npool_basetypes_miningpool_v1_enums_proto_depIdxs = nil
}
