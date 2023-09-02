// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.18.1
// source: npool/basetypes/ledger/v1/enums.proto

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

type IOType int32

const (
	IOType_DefaultType IOType = 0
	IOType_Incoming    IOType = 10
	IOType_Outcoming   IOType = 20
)

// Enum value maps for IOType.
var (
	IOType_name = map[int32]string{
		0:  "DefaultType",
		10: "Incoming",
		20: "Outcoming",
	}
	IOType_value = map[string]int32{
		"DefaultType": 0,
		"Incoming":    10,
		"Outcoming":   20,
	}
)

func (x IOType) Enum() *IOType {
	p := new(IOType)
	*p = x
	return p
}

func (x IOType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (IOType) Descriptor() protoreflect.EnumDescriptor {
	return file_npool_basetypes_ledger_v1_enums_proto_enumTypes[0].Descriptor()
}

func (IOType) Type() protoreflect.EnumType {
	return &file_npool_basetypes_ledger_v1_enums_proto_enumTypes[0]
}

func (x IOType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use IOType.Descriptor instead.
func (IOType) EnumDescriptor() ([]byte, []int) {
	return file_npool_basetypes_ledger_v1_enums_proto_rawDescGZIP(), []int{0}
}

type IOSubType int32

const (
	IOSubType_DefaultSubType IOSubType = 0
	// I/O
	// I: record tids in extra
	// O: record order id in extra
	IOSubType_Payment IOSubType = 100
	// I: record good id in extra
	IOSubType_MiningBenefit IOSubType = 110
	// I: record user id, order id in extra
	IOSubType_Commission IOSubType = 120
	// I: record good id in extra
	IOSubType_TechniqueFeeCommission IOSubType = 130
	// I: record address in extra
	IOSubType_Deposit IOSubType = 140
	// I: record from user in extra
	// O: record target user in extra
	IOSubType_Transfer IOSubType = 150
	// O
	IOSubType_Withdrawal IOSubType = 160
	// I: record user id, order id in extra
	IOSubType_OrderRevoke IOSubType = 170
	// I: record user id,order id, archivement details id in extra
	IOSubType_CommissionRevoke IOSubType = 180
)

// Enum value maps for IOSubType.
var (
	IOSubType_name = map[int32]string{
		0:   "DefaultSubType",
		100: "Payment",
		110: "MiningBenefit",
		120: "Commission",
		130: "TechniqueFeeCommission",
		140: "Deposit",
		150: "Transfer",
		160: "Withdrawal",
		170: "OrderRevoke",
		180: "CommissionRevoke",
	}
	IOSubType_value = map[string]int32{
		"DefaultSubType":         0,
		"Payment":                100,
		"MiningBenefit":          110,
		"Commission":             120,
		"TechniqueFeeCommission": 130,
		"Deposit":                140,
		"Transfer":               150,
		"Withdrawal":             160,
		"OrderRevoke":            170,
		"CommissionRevoke":       180,
	}
)

func (x IOSubType) Enum() *IOSubType {
	p := new(IOSubType)
	*p = x
	return p
}

func (x IOSubType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (IOSubType) Descriptor() protoreflect.EnumDescriptor {
	return file_npool_basetypes_ledger_v1_enums_proto_enumTypes[1].Descriptor()
}

func (IOSubType) Type() protoreflect.EnumType {
	return &file_npool_basetypes_ledger_v1_enums_proto_enumTypes[1]
}

func (x IOSubType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use IOSubType.Descriptor instead.
func (IOSubType) EnumDescriptor() ([]byte, []int) {
	return file_npool_basetypes_ledger_v1_enums_proto_rawDescGZIP(), []int{1}
}

type WithdrawState int32

const (
	WithdrawState_DefaultWithdrawState WithdrawState = 0
	WithdrawState_Created              WithdrawState = 10
	WithdrawState_Reviewing            WithdrawState = 20
	WithdrawState_Approved             WithdrawState = 30
	WithdrawState_Rejected             WithdrawState = 40
	WithdrawState_Transferring         WithdrawState = 50
	WithdrawState_TransactionFail      WithdrawState = 60
	WithdrawState_Successful           WithdrawState = 70
)

// Enum value maps for WithdrawState.
var (
	WithdrawState_name = map[int32]string{
		0:  "DefaultWithdrawState",
		10: "Created",
		20: "Reviewing",
		30: "Approved",
		40: "Rejected",
		50: "Transferring",
		60: "TransactionFail",
		70: "Successful",
	}
	WithdrawState_value = map[string]int32{
		"DefaultWithdrawState": 0,
		"Created":              10,
		"Reviewing":            20,
		"Approved":             30,
		"Rejected":             40,
		"Transferring":         50,
		"TransactionFail":      60,
		"Successful":           70,
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
	return file_npool_basetypes_ledger_v1_enums_proto_enumTypes[2].Descriptor()
}

func (WithdrawState) Type() protoreflect.EnumType {
	return &file_npool_basetypes_ledger_v1_enums_proto_enumTypes[2]
}

func (x WithdrawState) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use WithdrawState.Descriptor instead.
func (WithdrawState) EnumDescriptor() ([]byte, []int) {
	return file_npool_basetypes_ledger_v1_enums_proto_rawDescGZIP(), []int{2}
}

var File_npool_basetypes_ledger_v1_enums_proto protoreflect.FileDescriptor

var file_npool_basetypes_ledger_v1_enums_proto_rawDesc = []byte{
	0x0a, 0x25, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x2f, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x6e, 0x75, 0x6d,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70,
	0x65, 0x73, 0x2e, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2a, 0x36, 0x0a, 0x06,
	0x49, 0x4f, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0f, 0x0a, 0x0b, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c,
	0x74, 0x54, 0x79, 0x70, 0x65, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x49, 0x6e, 0x63, 0x6f, 0x6d,
	0x69, 0x6e, 0x67, 0x10, 0x0a, 0x12, 0x0d, 0x0a, 0x09, 0x4f, 0x75, 0x74, 0x63, 0x6f, 0x6d, 0x69,
	0x6e, 0x67, 0x10, 0x14, 0x2a, 0xc3, 0x01, 0x0a, 0x09, 0x49, 0x4f, 0x53, 0x75, 0x62, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x12, 0x0a, 0x0e, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x53, 0x75, 0x62,
	0x54, 0x79, 0x70, 0x65, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e,
	0x74, 0x10, 0x64, 0x12, 0x11, 0x0a, 0x0d, 0x4d, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x42, 0x65, 0x6e,
	0x65, 0x66, 0x69, 0x74, 0x10, 0x6e, 0x12, 0x0e, 0x0a, 0x0a, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x10, 0x78, 0x12, 0x1b, 0x0a, 0x16, 0x54, 0x65, 0x63, 0x68, 0x6e, 0x69,
	0x71, 0x75, 0x65, 0x46, 0x65, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x10, 0x82, 0x01, 0x12, 0x0c, 0x0a, 0x07, 0x44, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x10, 0x8c,
	0x01, 0x12, 0x0d, 0x0a, 0x08, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x10, 0x96, 0x01,
	0x12, 0x0f, 0x0a, 0x0a, 0x57, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x61, 0x6c, 0x10, 0xa0,
	0x01, 0x12, 0x10, 0x0a, 0x0b, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x76, 0x6f, 0x6b, 0x65,
	0x10, 0xaa, 0x01, 0x12, 0x15, 0x0a, 0x10, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x76, 0x6f, 0x6b, 0x65, 0x10, 0xb4, 0x01, 0x2a, 0x98, 0x01, 0x0a, 0x0d, 0x57,
	0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x18, 0x0a, 0x14,
	0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x57, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x53,
	0x74, 0x61, 0x74, 0x65, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x10, 0x0a, 0x12, 0x0d, 0x0a, 0x09, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x69, 0x6e, 0x67,
	0x10, 0x14, 0x12, 0x0c, 0x0a, 0x08, 0x41, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x65, 0x64, 0x10, 0x1e,
	0x12, 0x0c, 0x0a, 0x08, 0x52, 0x65, 0x6a, 0x65, 0x63, 0x74, 0x65, 0x64, 0x10, 0x28, 0x12, 0x10,
	0x0a, 0x0c, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x72, 0x69, 0x6e, 0x67, 0x10, 0x32,
	0x12, 0x13, 0x0a, 0x0f, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x46,
	0x61, 0x69, 0x6c, 0x10, 0x3c, 0x12, 0x0e, 0x0a, 0x0a, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x66, 0x75, 0x6c, 0x10, 0x46, 0x42, 0x3c, 0x5a, 0x3a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x4e, 0x70, 0x6f, 0x6f, 0x6c, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72,
	0x6d, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f,
	0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72,
	0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_npool_basetypes_ledger_v1_enums_proto_rawDescOnce sync.Once
	file_npool_basetypes_ledger_v1_enums_proto_rawDescData = file_npool_basetypes_ledger_v1_enums_proto_rawDesc
)

func file_npool_basetypes_ledger_v1_enums_proto_rawDescGZIP() []byte {
	file_npool_basetypes_ledger_v1_enums_proto_rawDescOnce.Do(func() {
		file_npool_basetypes_ledger_v1_enums_proto_rawDescData = protoimpl.X.CompressGZIP(file_npool_basetypes_ledger_v1_enums_proto_rawDescData)
	})
	return file_npool_basetypes_ledger_v1_enums_proto_rawDescData
}

var file_npool_basetypes_ledger_v1_enums_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_npool_basetypes_ledger_v1_enums_proto_goTypes = []interface{}{
	(IOType)(0),        // 0: basetypes.ledger.v1.IOType
	(IOSubType)(0),     // 1: basetypes.ledger.v1.IOSubType
	(WithdrawState)(0), // 2: basetypes.ledger.v1.WithdrawState
}
var file_npool_basetypes_ledger_v1_enums_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_npool_basetypes_ledger_v1_enums_proto_init() }
func file_npool_basetypes_ledger_v1_enums_proto_init() {
	if File_npool_basetypes_ledger_v1_enums_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_npool_basetypes_ledger_v1_enums_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_npool_basetypes_ledger_v1_enums_proto_goTypes,
		DependencyIndexes: file_npool_basetypes_ledger_v1_enums_proto_depIdxs,
		EnumInfos:         file_npool_basetypes_ledger_v1_enums_proto_enumTypes,
	}.Build()
	File_npool_basetypes_ledger_v1_enums_proto = out.File
	file_npool_basetypes_ledger_v1_enums_proto_rawDesc = nil
	file_npool_basetypes_ledger_v1_enums_proto_goTypes = nil
	file_npool_basetypes_ledger_v1_enums_proto_depIdxs = nil
}
