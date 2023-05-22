// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.18.1
// source: npool/basetypes/v1/txstate.proto

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

type TxState int32

const (
<<<<<<< HEAD
	TxState_DefaultTxState      TxState = 0
	TxState_TxStateCreated      TxState = 10
	TxState_TxStateWait         TxState = 20
	TxState_TxStateTransferring TxState = 30
	TxState_TxStateSuccessful   TxState = 40
	TxState_TxStateFail         TxState = 50
=======
	TxState_DefaultTxState    TxState = 0
	TxState_StateCreated      TxState = 10
	TxState_StateWait         TxState = 20
	TxState_StateTransferring TxState = 30
	TxState_StateSuccessful   TxState = 40
	TxState_StateFail         TxState = 50
>>>>>>> Add tx state to basetypes
)

// Enum value maps for TxState.
var (
	TxState_name = map[int32]string{
		0:  "DefaultTxState",
<<<<<<< HEAD
		10: "TxStateCreated",
		20: "TxStateWait",
		30: "TxStateTransferring",
		40: "TxStateSuccessful",
		50: "TxStateFail",
	}
	TxState_value = map[string]int32{
		"DefaultTxState":      0,
		"TxStateCreated":      10,
		"TxStateWait":         20,
		"TxStateTransferring": 30,
		"TxStateSuccessful":   40,
		"TxStateFail":         50,
=======
		10: "StateCreated",
		20: "StateWait",
		30: "StateTransferring",
		40: "StateSuccessful",
		50: "StateFail",
	}
	TxState_value = map[string]int32{
		"DefaultTxState":    0,
		"StateCreated":      10,
		"StateWait":         20,
		"StateTransferring": 30,
		"StateSuccessful":   40,
		"StateFail":         50,
>>>>>>> Add tx state to basetypes
	}
)

func (x TxState) Enum() *TxState {
	p := new(TxState)
	*p = x
	return p
}

func (x TxState) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TxState) Descriptor() protoreflect.EnumDescriptor {
	return file_npool_basetypes_v1_txstate_proto_enumTypes[0].Descriptor()
}

func (TxState) Type() protoreflect.EnumType {
	return &file_npool_basetypes_v1_txstate_proto_enumTypes[0]
}

func (x TxState) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TxState.Descriptor instead.
func (TxState) EnumDescriptor() ([]byte, []int) {
	return file_npool_basetypes_v1_txstate_proto_rawDescGZIP(), []int{0}
}

var File_npool_basetypes_v1_txstate_proto protoreflect.FileDescriptor

var file_npool_basetypes_v1_txstate_proto_rawDesc = []byte{
	0x0a, 0x20, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x78, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0c, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31,
<<<<<<< HEAD
	0x2a, 0x83, 0x01, 0x0a, 0x07, 0x54, 0x78, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x12, 0x0a, 0x0e,
	0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x54, 0x78, 0x53, 0x74, 0x61, 0x74, 0x65, 0x10, 0x00,
	0x12, 0x12, 0x0a, 0x0e, 0x54, 0x78, 0x53, 0x74, 0x61, 0x74, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x10, 0x0a, 0x12, 0x0f, 0x0a, 0x0b, 0x54, 0x78, 0x53, 0x74, 0x61, 0x74, 0x65, 0x57,
	0x61, 0x69, 0x74, 0x10, 0x14, 0x12, 0x17, 0x0a, 0x13, 0x54, 0x78, 0x53, 0x74, 0x61, 0x74, 0x65,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x72, 0x69, 0x6e, 0x67, 0x10, 0x1e, 0x12, 0x15,
	0x0a, 0x11, 0x54, 0x78, 0x53, 0x74, 0x61, 0x74, 0x65, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x66, 0x75, 0x6c, 0x10, 0x28, 0x12, 0x0f, 0x0a, 0x0b, 0x54, 0x78, 0x53, 0x74, 0x61, 0x74, 0x65,
	0x46, 0x61, 0x69, 0x6c, 0x10, 0x32, 0x42, 0x35, 0x5a, 0x33, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4e, 0x70, 0x6f, 0x6f, 0x6c, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f,
	0x72, 0x6d, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x6e, 0x70, 0x6f, 0x6f, 0x6c,
	0x2f, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
=======
	0x2a, 0x79, 0x0a, 0x07, 0x54, 0x78, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x12, 0x0a, 0x0e, 0x44,
	0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x54, 0x78, 0x53, 0x74, 0x61, 0x74, 0x65, 0x10, 0x00, 0x12,
	0x10, 0x0a, 0x0c, 0x53, 0x74, 0x61, 0x74, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x10,
	0x0a, 0x12, 0x0d, 0x0a, 0x09, 0x53, 0x74, 0x61, 0x74, 0x65, 0x57, 0x61, 0x69, 0x74, 0x10, 0x14,
	0x12, 0x15, 0x0a, 0x11, 0x53, 0x74, 0x61, 0x74, 0x65, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65,
	0x72, 0x72, 0x69, 0x6e, 0x67, 0x10, 0x1e, 0x12, 0x13, 0x0a, 0x0f, 0x53, 0x74, 0x61, 0x74, 0x65,
	0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x66, 0x75, 0x6c, 0x10, 0x28, 0x12, 0x0d, 0x0a, 0x09,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x46, 0x61, 0x69, 0x6c, 0x10, 0x32, 0x42, 0x35, 0x5a, 0x33, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4e, 0x70, 0x6f, 0x6f, 0x6c, 0x50,
	0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f,
	0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f,
	0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
>>>>>>> Add tx state to basetypes
}

var (
	file_npool_basetypes_v1_txstate_proto_rawDescOnce sync.Once
	file_npool_basetypes_v1_txstate_proto_rawDescData = file_npool_basetypes_v1_txstate_proto_rawDesc
)

func file_npool_basetypes_v1_txstate_proto_rawDescGZIP() []byte {
	file_npool_basetypes_v1_txstate_proto_rawDescOnce.Do(func() {
		file_npool_basetypes_v1_txstate_proto_rawDescData = protoimpl.X.CompressGZIP(file_npool_basetypes_v1_txstate_proto_rawDescData)
	})
	return file_npool_basetypes_v1_txstate_proto_rawDescData
}

var file_npool_basetypes_v1_txstate_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_npool_basetypes_v1_txstate_proto_goTypes = []interface{}{
	(TxState)(0), // 0: basetypes.v1.TxState
}
var file_npool_basetypes_v1_txstate_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_npool_basetypes_v1_txstate_proto_init() }
func file_npool_basetypes_v1_txstate_proto_init() {
	if File_npool_basetypes_v1_txstate_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_npool_basetypes_v1_txstate_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_npool_basetypes_v1_txstate_proto_goTypes,
		DependencyIndexes: file_npool_basetypes_v1_txstate_proto_depIdxs,
		EnumInfos:         file_npool_basetypes_v1_txstate_proto_enumTypes,
	}.Build()
	File_npool_basetypes_v1_txstate_proto = out.File
	file_npool_basetypes_v1_txstate_proto_rawDesc = nil
	file_npool_basetypes_v1_txstate_proto_goTypes = nil
	file_npool_basetypes_v1_txstate_proto_depIdxs = nil
}
