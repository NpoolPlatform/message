// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.18.1
// source: npool/basetypes/v1/pubsub.proto

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

type MsgID int32

const (
	MsgID_DefaultMsgID                 MsgID = 0
	MsgID_RewardEventReq               MsgID = 10
	MsgID_IncreaseUserActionCreditsReq MsgID = 20
	MsgID_CreateLoginHistoryReq        MsgID = 30
	MsgID_CreateAuthHistoryReq         MsgID = 40
	MsgID_CreateReviewReq              MsgID = 50
	MsgID_RegisterAPIsReq              MsgID = 60
	MsgID_UpdateOpLogHumanReadableReq  MsgID = 70
	MsgID_CreateNewLoginReq            MsgID = 80
	MsgID_CreateCommissionReq          MsgID = 90
	MsgID_UpdateCouponsUsedReq         MsgID = 100
	MsgID_DepositReceivedReq           MsgID = 1000
	MsgID_DepositCheckFailReq          MsgID = 1010
	MsgID_WithdrawRequestReq           MsgID = 2000
	MsgID_WithdrawSuccessReq           MsgID = 2010
	MsgID_WithdrwaFailReq              MsgID = 2020
	MsgID_CreateGoodBenefitReq         MsgID = 3000
	MsgID_OrderPaidReq                 MsgID = 4000
	MsgID_OrderChildsRenewReq          MsgID = 4010
	MsgID_WithdrawReviewNotifyReq      MsgID = 5000
	MsgID_CalculateEventRewardReq      MsgID = 6000
	MsgID_ReliableEventRewardReq       MsgID = 6010
	MsgID_UnReliableEventRewardReq     MsgID = 6020
	MsgID_CreateSampleMsgReq           MsgID = 100000000
)

// Enum value maps for MsgID.
var (
	MsgID_name = map[int32]string{
		0:         "DefaultMsgID",
		10:        "RewardEventReq",
		20:        "IncreaseUserActionCreditsReq",
		30:        "CreateLoginHistoryReq",
		40:        "CreateAuthHistoryReq",
		50:        "CreateReviewReq",
		60:        "RegisterAPIsReq",
		70:        "UpdateOpLogHumanReadableReq",
		80:        "CreateNewLoginReq",
		90:        "CreateCommissionReq",
		100:       "UpdateCouponsUsedReq",
		1000:      "DepositReceivedReq",
		1010:      "DepositCheckFailReq",
		2000:      "WithdrawRequestReq",
		2010:      "WithdrawSuccessReq",
		2020:      "WithdrwaFailReq",
		3000:      "CreateGoodBenefitReq",
		4000:      "OrderPaidReq",
		4010:      "OrderChildsRenewReq",
		5000:      "WithdrawReviewNotifyReq",
		6000:      "CalculateEventRewardReq",
		6010:      "ReliableEventRewardReq",
		6020:      "UnReliableEventRewardReq",
		100000000: "CreateSampleMsgReq",
	}
	MsgID_value = map[string]int32{
		"DefaultMsgID":                 0,
		"RewardEventReq":               10,
		"IncreaseUserActionCreditsReq": 20,
		"CreateLoginHistoryReq":        30,
		"CreateAuthHistoryReq":         40,
		"CreateReviewReq":              50,
		"RegisterAPIsReq":              60,
		"UpdateOpLogHumanReadableReq":  70,
		"CreateNewLoginReq":            80,
		"CreateCommissionReq":          90,
		"UpdateCouponsUsedReq":         100,
		"DepositReceivedReq":           1000,
		"DepositCheckFailReq":          1010,
		"WithdrawRequestReq":           2000,
		"WithdrawSuccessReq":           2010,
		"WithdrwaFailReq":              2020,
		"CreateGoodBenefitReq":         3000,
		"OrderPaidReq":                 4000,
		"OrderChildsRenewReq":          4010,
		"WithdrawReviewNotifyReq":      5000,
		"CalculateEventRewardReq":      6000,
		"ReliableEventRewardReq":       6010,
		"UnReliableEventRewardReq":     6020,
		"CreateSampleMsgReq":           100000000,
	}
)

func (x MsgID) Enum() *MsgID {
	p := new(MsgID)
	*p = x
	return p
}

func (x MsgID) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MsgID) Descriptor() protoreflect.EnumDescriptor {
	return file_npool_basetypes_v1_pubsub_proto_enumTypes[0].Descriptor()
}

func (MsgID) Type() protoreflect.EnumType {
	return &file_npool_basetypes_v1_pubsub_proto_enumTypes[0]
}

func (x MsgID) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MsgID.Descriptor instead.
func (MsgID) EnumDescriptor() ([]byte, []int) {
	return file_npool_basetypes_v1_pubsub_proto_rawDescGZIP(), []int{0}
}

type MsgState int32

const (
	MsgState_DefaultMsgState MsgState = 0
	MsgState_StateSuccess    MsgState = 10
	MsgState_StateFail       MsgState = 20
)

// Enum value maps for MsgState.
var (
	MsgState_name = map[int32]string{
		0:  "DefaultMsgState",
		10: "StateSuccess",
		20: "StateFail",
	}
	MsgState_value = map[string]int32{
		"DefaultMsgState": 0,
		"StateSuccess":    10,
		"StateFail":       20,
	}
)

func (x MsgState) Enum() *MsgState {
	p := new(MsgState)
	*p = x
	return p
}

func (x MsgState) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MsgState) Descriptor() protoreflect.EnumDescriptor {
	return file_npool_basetypes_v1_pubsub_proto_enumTypes[1].Descriptor()
}

func (MsgState) Type() protoreflect.EnumType {
	return &file_npool_basetypes_v1_pubsub_proto_enumTypes[1]
}

func (x MsgState) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MsgState.Descriptor instead.
func (MsgState) EnumDescriptor() ([]byte, []int) {
	return file_npool_basetypes_v1_pubsub_proto_rawDescGZIP(), []int{1}
}

type MsgError struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Json value
	Value string `protobuf:"bytes,10,opt,name=Value,proto3" json:"Value,omitempty"`
	Error string `protobuf:"bytes,20,opt,name=Error,proto3" json:"Error,omitempty"`
}

func (x *MsgError) Reset() {
	*x = MsgError{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_basetypes_v1_pubsub_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgError) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgError) ProtoMessage() {}

func (x *MsgError) ProtoReflect() protoreflect.Message {
	mi := &file_npool_basetypes_v1_pubsub_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MsgError.ProtoReflect.Descriptor instead.
func (*MsgError) Descriptor() ([]byte, []int) {
	return file_npool_basetypes_v1_pubsub_proto_rawDescGZIP(), []int{0}
}

func (x *MsgError) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *MsgError) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

var File_npool_basetypes_v1_pubsub_proto protoreflect.FileDescriptor

var file_npool_basetypes_v1_pubsub_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x75, 0x62, 0x73, 0x75, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0c, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x22,
	0x36, 0x0a, 0x08, 0x4d, 0x73, 0x67, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x2a, 0xef, 0x04, 0x0a, 0x05, 0x4d, 0x73, 0x67, 0x49,
	0x44, 0x12, 0x10, 0x0a, 0x0c, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x4d, 0x73, 0x67, 0x49,
	0x44, 0x10, 0x00, 0x12, 0x12, 0x0a, 0x0e, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x52, 0x65, 0x71, 0x10, 0x0a, 0x12, 0x20, 0x0a, 0x1c, 0x49, 0x6e, 0x63, 0x72, 0x65,
	0x61, 0x73, 0x65, 0x55, 0x73, 0x65, 0x72, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x72, 0x65,
	0x64, 0x69, 0x74, 0x73, 0x52, 0x65, 0x71, 0x10, 0x14, 0x12, 0x19, 0x0a, 0x15, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x52,
	0x65, 0x71, 0x10, 0x1e, 0x12, 0x18, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x75,
	0x74, 0x68, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x10, 0x28, 0x12, 0x13,
	0x0a, 0x0f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x52, 0x65,
	0x71, 0x10, 0x32, 0x12, 0x13, 0x0a, 0x0f, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x41,
	0x50, 0x49, 0x73, 0x52, 0x65, 0x71, 0x10, 0x3c, 0x12, 0x1f, 0x0a, 0x1b, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x4f, 0x70, 0x4c, 0x6f, 0x67, 0x48, 0x75, 0x6d, 0x61, 0x6e, 0x52, 0x65, 0x61, 0x64,
	0x61, 0x62, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x10, 0x46, 0x12, 0x15, 0x0a, 0x11, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x4e, 0x65, 0x77, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x10, 0x50,
	0x12, 0x17, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x10, 0x5a, 0x12, 0x18, 0x0a, 0x14, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x43, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x73, 0x55, 0x73, 0x65, 0x64, 0x52, 0x65,
	0x71, 0x10, 0x64, 0x12, 0x17, 0x0a, 0x12, 0x44, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x52, 0x65,
	0x63, 0x65, 0x69, 0x76, 0x65, 0x64, 0x52, 0x65, 0x71, 0x10, 0xe8, 0x07, 0x12, 0x18, 0x0a, 0x13,
	0x44, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x46, 0x61, 0x69, 0x6c,
	0x52, 0x65, 0x71, 0x10, 0xf2, 0x07, 0x12, 0x17, 0x0a, 0x12, 0x57, 0x69, 0x74, 0x68, 0x64, 0x72,
	0x61, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x65, 0x71, 0x10, 0xd0, 0x0f, 0x12,
	0x17, 0x0a, 0x12, 0x57, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x53, 0x75, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x52, 0x65, 0x71, 0x10, 0xda, 0x0f, 0x12, 0x14, 0x0a, 0x0f, 0x57, 0x69, 0x74, 0x68,
	0x64, 0x72, 0x77, 0x61, 0x46, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x10, 0xe4, 0x0f, 0x12, 0x19,
	0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x47, 0x6f, 0x6f, 0x64, 0x42, 0x65, 0x6e, 0x65,
	0x66, 0x69, 0x74, 0x52, 0x65, 0x71, 0x10, 0xb8, 0x17, 0x12, 0x11, 0x0a, 0x0c, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x50, 0x61, 0x69, 0x64, 0x52, 0x65, 0x71, 0x10, 0xa0, 0x1f, 0x12, 0x18, 0x0a, 0x13,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x43, 0x68, 0x69, 0x6c, 0x64, 0x73, 0x52, 0x65, 0x6e, 0x65, 0x77,
	0x52, 0x65, 0x71, 0x10, 0xaa, 0x1f, 0x12, 0x1c, 0x0a, 0x17, 0x57, 0x69, 0x74, 0x68, 0x64, 0x72,
	0x61, 0x77, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x52, 0x65,
	0x71, 0x10, 0x88, 0x27, 0x12, 0x1c, 0x0a, 0x17, 0x43, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74,
	0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x52, 0x65, 0x71, 0x10,
	0xf0, 0x2e, 0x12, 0x1b, 0x0a, 0x16, 0x52, 0x65, 0x6c, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x52, 0x65, 0x71, 0x10, 0xfa, 0x2e, 0x12,
	0x1d, 0x0a, 0x18, 0x55, 0x6e, 0x52, 0x65, 0x6c, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x52, 0x65, 0x71, 0x10, 0x84, 0x2f, 0x12, 0x19,
	0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x4d, 0x73,
	0x67, 0x52, 0x65, 0x71, 0x10, 0x80, 0xc2, 0xd7, 0x2f, 0x2a, 0x40, 0x0a, 0x08, 0x4d, 0x73, 0x67,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x13, 0x0a, 0x0f, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74,
	0x4d, 0x73, 0x67, 0x53, 0x74, 0x61, 0x74, 0x65, 0x10, 0x00, 0x12, 0x10, 0x0a, 0x0c, 0x53, 0x74,
	0x61, 0x74, 0x65, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x10, 0x0a, 0x12, 0x0d, 0x0a, 0x09,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x46, 0x61, 0x69, 0x6c, 0x10, 0x14, 0x42, 0x35, 0x5a, 0x33, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4e, 0x70, 0x6f, 0x6f, 0x6c, 0x50,
	0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f,
	0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f,
	0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_npool_basetypes_v1_pubsub_proto_rawDescOnce sync.Once
	file_npool_basetypes_v1_pubsub_proto_rawDescData = file_npool_basetypes_v1_pubsub_proto_rawDesc
)

func file_npool_basetypes_v1_pubsub_proto_rawDescGZIP() []byte {
	file_npool_basetypes_v1_pubsub_proto_rawDescOnce.Do(func() {
		file_npool_basetypes_v1_pubsub_proto_rawDescData = protoimpl.X.CompressGZIP(file_npool_basetypes_v1_pubsub_proto_rawDescData)
	})
	return file_npool_basetypes_v1_pubsub_proto_rawDescData
}

var file_npool_basetypes_v1_pubsub_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_npool_basetypes_v1_pubsub_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_npool_basetypes_v1_pubsub_proto_goTypes = []interface{}{
	(MsgID)(0),       // 0: basetypes.v1.MsgID
	(MsgState)(0),    // 1: basetypes.v1.MsgState
	(*MsgError)(nil), // 2: basetypes.v1.MsgError
}
var file_npool_basetypes_v1_pubsub_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_npool_basetypes_v1_pubsub_proto_init() }
func file_npool_basetypes_v1_pubsub_proto_init() {
	if File_npool_basetypes_v1_pubsub_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_npool_basetypes_v1_pubsub_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgError); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_npool_basetypes_v1_pubsub_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_npool_basetypes_v1_pubsub_proto_goTypes,
		DependencyIndexes: file_npool_basetypes_v1_pubsub_proto_depIdxs,
		EnumInfos:         file_npool_basetypes_v1_pubsub_proto_enumTypes,
		MessageInfos:      file_npool_basetypes_v1_pubsub_proto_msgTypes,
	}.Build()
	File_npool_basetypes_v1_pubsub_proto = out.File
	file_npool_basetypes_v1_pubsub_proto_rawDesc = nil
	file_npool_basetypes_v1_pubsub_proto_goTypes = nil
	file_npool_basetypes_v1_pubsub_proto_depIdxs = nil
}
