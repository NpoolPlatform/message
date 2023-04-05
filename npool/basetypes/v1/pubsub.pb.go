// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
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

<<<<<<< HEAD
<<<<<<< HEAD
type MsgID int32

const (
	MsgID_DefaultMsgID    MsgID = 0
	MsgID_RewardEventReq  MsgID = 10
	MsgID_RewardEventResp MsgID = 20
)

// Enum value maps for MsgID.
var (
	MsgID_name = map[int32]string{
		0:  "DefaultMsgID",
		10: "RewardEventReq",
		20: "RewardEventResp",
	}
	MsgID_value = map[string]int32{
		"DefaultMsgID":    0,
		"RewardEventReq":  10,
		"RewardEventResp": 20,
	}
)

func (x MsgID) Enum() *MsgID {
	p := new(MsgID)
=======
type MessageID int32
=======
type MsgID int32
>>>>>>> Remove processing state

const (
	MsgID_DefaultMsgID MsgID = 0
	// For binding invitations during registration
	MsgID_CreateRegistrationInvitationTry     MsgID = 10010
	MsgID_CreateRegistrationInvitationConfirm MsgID = 10020
	MsgID_CreateRegistrationInvitationCancel  MsgID = 10030
)

// Enum value maps for MsgID.
var (
	MsgID_name = map[int32]string{
		0:     "DefaultMsgID",
		10010: "CreateRegistrationInvitationTry",
		10020: "CreateRegistrationInvitationConfirm",
		10030: "CreateRegistrationInvitationCancel",
	}
	MsgID_value = map[string]int32{
		"DefaultMsgID":                        0,
		"CreateRegistrationInvitationTry":     10010,
		"CreateRegistrationInvitationConfirm": 10020,
		"CreateRegistrationInvitationCancel":  10030,
	}
)

<<<<<<< HEAD
func (x MessageID) Enum() *MessageID {
	p := new(MessageID)
>>>>>>>  update pubsub
=======
func (x MsgID) Enum() *MsgID {
	p := new(MsgID)
>>>>>>> Remove processing state
	*p = x
	return p
}

<<<<<<< HEAD
<<<<<<< HEAD
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
=======
func (x MessageID) String() string {
=======
func (x MsgID) String() string {
>>>>>>> Remove processing state
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
	MsgState_DefaultMessageState MsgState = 0
	MsgState_StateSuccess        MsgState = 10
	MsgState_StateFail           MsgState = 20
)

// Enum value maps for MsgState.
var (
	MsgState_name = map[int32]string{
		0:  "DefaultMessageState",
		10: "StateSuccess",
		20: "StateFail",
	}
	MsgState_value = map[string]int32{
		"DefaultMessageState": 0,
		"StateSuccess":        10,
		"StateFail":           20,
	}
)

<<<<<<< HEAD
func (x MessageState) Enum() *MessageState {
	p := new(MessageState)
>>>>>>>  update pubsub
=======
func (x MsgState) Enum() *MsgState {
	p := new(MsgState)
>>>>>>> Remove processing state
	*p = x
	return p
}

<<<<<<< HEAD
<<<<<<< HEAD
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
=======
func (x MessageState) String() string {
=======
func (x MsgState) String() string {
>>>>>>> Remove processing state
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

<<<<<<< HEAD
// Deprecated: Use MessageState.Descriptor instead.
func (MessageState) EnumDescriptor() ([]byte, []int) {
>>>>>>>  update pubsub
=======
// Deprecated: Use MsgState.Descriptor instead.
func (MsgState) EnumDescriptor() ([]byte, []int) {
>>>>>>> Remove processing state
	return file_npool_basetypes_v1_pubsub_proto_rawDescGZIP(), []int{1}
}

var File_npool_basetypes_v1_pubsub_proto protoreflect.FileDescriptor

var file_npool_basetypes_v1_pubsub_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x75, 0x62, 0x73, 0x75, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0c, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2a,
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
	0x42, 0x0a, 0x05, 0x4d, 0x73, 0x67, 0x49, 0x44, 0x12, 0x10, 0x0a, 0x0c, 0x44, 0x65, 0x66, 0x61,
	0x75, 0x6c, 0x74, 0x4d, 0x73, 0x67, 0x49, 0x44, 0x10, 0x00, 0x12, 0x12, 0x0a, 0x0e, 0x52, 0x65,
	0x77, 0x61, 0x72, 0x64, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x10, 0x0a, 0x12, 0x13,
	0x0a, 0x0f, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x10, 0x14, 0x2a, 0x40, 0x0a, 0x08, 0x4d, 0x73, 0x67, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12,
	0x13, 0x0a, 0x0f, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x4d, 0x73, 0x67, 0x53, 0x74, 0x61,
	0x74, 0x65, 0x10, 0x00, 0x12, 0x10, 0x0a, 0x0c, 0x53, 0x74, 0x61, 0x74, 0x65, 0x53, 0x75, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x10, 0x0a, 0x12, 0x0d, 0x0a, 0x09, 0x53, 0x74, 0x61, 0x74, 0x65, 0x46,
	0x61, 0x69, 0x6c, 0x10, 0x14, 0x42, 0x35, 0x5a, 0x33, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x4e, 0x70, 0x6f, 0x6f, 0x6c, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72,
	0x6d, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f,
	0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
=======
	0x5c, 0x0a, 0x09, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x10,
	0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x49, 0x44,
	0x10, 0x00, 0x12, 0x1b, 0x0a, 0x17, 0x53, 0x69, 0x67, 0x6e, 0x75, 0x70, 0x49, 0x6e, 0x76, 0x69,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x10, 0x0a, 0x12,
	0x1c, 0x0a, 0x18, 0x53, 0x69, 0x67, 0x6e, 0x75, 0x70, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x10, 0x14, 0x2a, 0x4d, 0x0a,
	0x0c, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x17, 0x0a,
	0x13, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x53,
	0x74, 0x61, 0x74, 0x65, 0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09, 0x49, 0x6e, 0x50, 0x72, 0x6f, 0x63,
	0x65, 0x73, 0x73, 0x10, 0x0a, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x10, 0x14, 0x12, 0x08, 0x0a, 0x04, 0x46, 0x61, 0x69, 0x6c, 0x10, 0x1e, 0x42, 0x35, 0x5a, 0x33,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4e, 0x70, 0x6f, 0x6f, 0x6c,
	0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x2f, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
>>>>>>>  update pubsub
=======
	0x6e, 0x0a, 0x09, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x10,
	0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x49, 0x44,
	0x10, 0x00, 0x12, 0x24, 0x0a, 0x1f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x71, 0x10, 0x9a, 0x4e, 0x12, 0x25, 0x0a, 0x20, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e,
	0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x10, 0xa4, 0x4e, 0x2a,
	0x4e, 0x0a, 0x0c, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12,
	0x17, 0x0a, 0x13, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x10, 0x00, 0x12, 0x0e, 0x0a, 0x0a, 0x50, 0x72, 0x6f, 0x63,
	0x65, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x10, 0x0a, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x75, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x10, 0x14, 0x12, 0x08, 0x0a, 0x04, 0x46, 0x61, 0x69, 0x6c, 0x10, 0x1e, 0x42,
	0x35, 0x5a, 0x33, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4e, 0x70,
	0x6f, 0x6f, 0x6c, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x2f, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
>>>>>>> Refactor message definitions
=======
	0x92, 0x01, 0x0a, 0x05, 0x4d, 0x73, 0x67, 0x49, 0x44, 0x12, 0x10, 0x0a, 0x0c, 0x44, 0x65, 0x66,
	0x61, 0x75, 0x6c, 0x74, 0x4d, 0x73, 0x67, 0x49, 0x44, 0x10, 0x00, 0x12, 0x24, 0x0a, 0x1f, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x72, 0x79, 0x10, 0x9a,
	0x4e, 0x12, 0x28, 0x0a, 0x23, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x67, 0x69, 0x73,
	0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x10, 0xa4, 0x4e, 0x12, 0x27, 0x0a, 0x22, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x61, 0x6e, 0x63, 0x65,
	0x6c, 0x10, 0xae, 0x4e, 0x2a, 0x44, 0x0a, 0x08, 0x4d, 0x73, 0x67, 0x53, 0x74, 0x61, 0x74, 0x65,
	0x12, 0x17, 0x0a, 0x13, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x10, 0x00, 0x12, 0x10, 0x0a, 0x0c, 0x53, 0x74, 0x61,
	0x74, 0x65, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x10, 0x0a, 0x12, 0x0d, 0x0a, 0x09, 0x53,
	0x74, 0x61, 0x74, 0x65, 0x46, 0x61, 0x69, 0x6c, 0x10, 0x14, 0x42, 0x35, 0x5a, 0x33, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4e, 0x70, 0x6f, 0x6f, 0x6c, 0x50, 0x6c,
	0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x6e,
	0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x76,
	0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
>>>>>>> Remove processing state
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
var file_npool_basetypes_v1_pubsub_proto_goTypes = []interface{}{
<<<<<<< HEAD
<<<<<<< HEAD
	(MsgID)(0),    // 0: basetypes.v1.MsgID
	(MsgState)(0), // 1: basetypes.v1.MsgState
=======
	(MessageID)(0),    // 0: basetypes.v1.MessageID
	(MessageState)(0), // 1: basetypes.v1.MessageState
>>>>>>>  update pubsub
=======
	(MsgID)(0),    // 0: basetypes.v1.MsgID
	(MsgState)(0), // 1: basetypes.v1.MsgState
>>>>>>> Remove processing state
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_npool_basetypes_v1_pubsub_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_npool_basetypes_v1_pubsub_proto_goTypes,
		DependencyIndexes: file_npool_basetypes_v1_pubsub_proto_depIdxs,
		EnumInfos:         file_npool_basetypes_v1_pubsub_proto_enumTypes,
	}.Build()
	File_npool_basetypes_v1_pubsub_proto = out.File
	file_npool_basetypes_v1_pubsub_proto_rawDesc = nil
	file_npool_basetypes_v1_pubsub_proto_goTypes = nil
	file_npool_basetypes_v1_pubsub_proto_depIdxs = nil
}
