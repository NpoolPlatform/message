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

type MsgID int32

const (
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
=======
>>>>>>> Add create user and create invitation code tcc message
	MsgID_DefaultMsgID                            MsgID = 0
	MsgID_CreateRegistrationInvitationTry         MsgID = 1000010
	MsgID_CreateRegistrationInvitationTryResp     MsgID = 1000020
	MsgID_CreateRegistrationInvitationConfirm     MsgID = 1000030
	MsgID_CreateRegistrationInvitationConfirmResp MsgID = 1000040
	MsgID_CreateRegistrationInvitationCancel      MsgID = 1000050
	MsgID_CreateRegistrationInvitationCancelResp  MsgID = 1000060
	MsgID_CreateInvitationCodeTry                 MsgID = 1000070
	MsgID_CreateInvitationCodeTryResp             MsgID = 1000080
	MsgID_CreateInvitationCodeConfirm             MsgID = 1000090
	MsgID_CreateInvitationCodeConfirmResp         MsgID = 1000100
	MsgID_CreateInvitationCodeCancel              MsgID = 1000110
	MsgID_CreateInvitationCodeCancelResp          MsgID = 1000120
	MsgID_CreateUserTry                           MsgID = 1000130
	MsgID_CreateUserTryResp                       MsgID = 1000140
	MsgID_CreateUserConfirm                       MsgID = 1000150
	MsgID_CreateUserConfirmResp                   MsgID = 1000160
	MsgID_CreateUserCancel                        MsgID = 1000170
	MsgID_CreateUserCancelResp                    MsgID = 1000180
<<<<<<< HEAD
=======
	MsgID_DefaultMsgID MsgID = 0
	// For binding invitations during registration
<<<<<<< HEAD
	MsgID_CreateRegistrationInvitationTry     MsgID = 10010
	MsgID_CreateRegistrationInvitationConfirm MsgID = 10020
	MsgID_CreateRegistrationInvitationCancel  MsgID = 10030
>>>>>>> Remove processing state
=======
	MsgID_CreateRegistrationInvitationTry         MsgID = 10010
	MsgID_CreateRegistrationInvitationTryResp     MsgID = 10020
	MsgID_CreateRegistrationInvitationConfirm     MsgID = 10030
	MsgID_CreateRegistrationInvitationConfirmResp MsgID = 10040
	MsgID_CreateRegistrationInvitationCancel      MsgID = 10050
	MsgID_CreateRegistrationInvitationCancelResp  MsgID = 10060
>>>>>>> Add response
=======
>>>>>>> Add create user and create invitation code tcc message
=======
	MsgID_DefaultMsgID    MsgID = 0
	MsgID_CreateCouponReq MsgID = 10
>>>>>>> Support create coupon
)

// Enum value maps for MsgID.
var (
	MsgID_name = map[int32]string{
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
=======
>>>>>>> Add create user and create invitation code tcc message
		0:       "DefaultMsgID",
		1000010: "CreateRegistrationInvitationTry",
		1000020: "CreateRegistrationInvitationTryResp",
		1000030: "CreateRegistrationInvitationConfirm",
		1000040: "CreateRegistrationInvitationConfirmResp",
		1000050: "CreateRegistrationInvitationCancel",
		1000060: "CreateRegistrationInvitationCancelResp",
		1000070: "CreateInvitationCodeTry",
		1000080: "CreateInvitationCodeTryResp",
		1000090: "CreateInvitationCodeConfirm",
		1000100: "CreateInvitationCodeConfirmResp",
		1000110: "CreateInvitationCodeCancel",
		1000120: "CreateInvitationCodeCancelResp",
		1000130: "CreateUserTry",
		1000140: "CreateUserTryResp",
		1000150: "CreateUserConfirm",
		1000160: "CreateUserConfirmResp",
		1000170: "CreateUserCancel",
		1000180: "CreateUserCancelResp",
<<<<<<< HEAD
	}
	MsgID_value = map[string]int32{
		"DefaultMsgID":                            0,
		"CreateRegistrationInvitationTry":         1000010,
		"CreateRegistrationInvitationTryResp":     1000020,
		"CreateRegistrationInvitationConfirm":     1000030,
		"CreateRegistrationInvitationConfirmResp": 1000040,
		"CreateRegistrationInvitationCancel":      1000050,
		"CreateRegistrationInvitationCancelResp":  1000060,
		"CreateInvitationCodeTry":                 1000070,
		"CreateInvitationCodeTryResp":             1000080,
		"CreateInvitationCodeConfirm":             1000090,
		"CreateInvitationCodeConfirmResp":         1000100,
		"CreateInvitationCodeCancel":              1000110,
		"CreateInvitationCodeCancelResp":          1000120,
		"CreateUserTry":                           1000130,
		"CreateUserTryResp":                       1000140,
		"CreateUserConfirm":                       1000150,
		"CreateUserConfirmResp":                   1000160,
		"CreateUserCancel":                        1000170,
		"CreateUserCancelResp":                    1000180,
=======
		0:     "DefaultMsgID",
		10010: "CreateRegistrationInvitationTry",
		10020: "CreateRegistrationInvitationTryResp",
		10030: "CreateRegistrationInvitationConfirm",
		10040: "CreateRegistrationInvitationConfirmResp",
		10050: "CreateRegistrationInvitationCancel",
		10060: "CreateRegistrationInvitationCancelResp",
=======
>>>>>>> Add create user and create invitation code tcc message
	}
	MsgID_value = map[string]int32{
<<<<<<< HEAD
		"DefaultMsgID":                        0,
		"CreateRegistrationInvitationTry":     10010,
		"CreateRegistrationInvitationConfirm": 10020,
		"CreateRegistrationInvitationCancel":  10030,
>>>>>>> Remove processing state
=======
		"DefaultMsgID":                            0,
<<<<<<< HEAD
		"CreateRegistrationInvitationTry":         10010,
		"CreateRegistrationInvitationTryResp":     10020,
		"CreateRegistrationInvitationConfirm":     10030,
		"CreateRegistrationInvitationConfirmResp": 10040,
		"CreateRegistrationInvitationCancel":      10050,
		"CreateRegistrationInvitationCancelResp":  10060,
>>>>>>> Add response
=======
		"CreateRegistrationInvitationTry":         1000010,
		"CreateRegistrationInvitationTryResp":     1000020,
		"CreateRegistrationInvitationConfirm":     1000030,
		"CreateRegistrationInvitationConfirmResp": 1000040,
		"CreateRegistrationInvitationCancel":      1000050,
		"CreateRegistrationInvitationCancelResp":  1000060,
		"CreateInvitationCodeTry":                 1000070,
		"CreateInvitationCodeTryResp":             1000080,
		"CreateInvitationCodeConfirm":             1000090,
		"CreateInvitationCodeConfirmResp":         1000100,
		"CreateInvitationCodeCancel":              1000110,
		"CreateInvitationCodeCancelResp":          1000120,
		"CreateUserTry":                           1000130,
		"CreateUserTryResp":                       1000140,
		"CreateUserConfirm":                       1000150,
		"CreateUserConfirmResp":                   1000160,
		"CreateUserCancel":                        1000170,
		"CreateUserCancelResp":                    1000180,
>>>>>>> Add create user and create invitation code tcc message
=======
		0:  "DefaultMsgID",
		10: "CreateCouponReq",
	}
	MsgID_value = map[string]int32{
		"DefaultMsgID":    0,
		"CreateCouponReq": 10,
>>>>>>> Support create coupon
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
<<<<<<< HEAD
<<<<<<< HEAD
	MsgState_DefaultMsgState MsgState = 0
	MsgState_StateSuccess    MsgState = 10
	MsgState_StateFail       MsgState = 20
=======
	MsgState_DefaultMessageState MsgState = 0
	MsgState_StateSuccess        MsgState = 10
	MsgState_StateFail           MsgState = 20
>>>>>>> Remove processing state
=======
	MsgState_DefaultMsgState MsgState = 0
	MsgState_StateSuccess    MsgState = 10
	MsgState_StateFail       MsgState = 20
>>>>>>> Refactor name
)

// Enum value maps for MsgState.
var (
	MsgState_name = map[int32]string{
<<<<<<< HEAD
<<<<<<< HEAD
		0:  "DefaultMsgState",
=======
		0:  "DefaultMessageState",
>>>>>>> Remove processing state
=======
		0:  "DefaultMsgState",
>>>>>>> Refactor name
		10: "StateSuccess",
		20: "StateFail",
	}
	MsgState_value = map[string]int32{
<<<<<<< HEAD
<<<<<<< HEAD
		"DefaultMsgState": 0,
		"StateSuccess":    10,
		"StateFail":       20,
=======
		"DefaultMessageState": 0,
		"StateSuccess":        10,
		"StateFail":           20,
>>>>>>> Remove processing state
=======
		"DefaultMsgState": 0,
		"StateSuccess":    10,
		"StateFail":       20,
>>>>>>> Refactor name
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

var File_npool_basetypes_v1_pubsub_proto protoreflect.FileDescriptor

var file_npool_basetypes_v1_pubsub_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x75, 0x62, 0x73, 0x75, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0c, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2a,
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
	0x89, 0x05, 0x0a, 0x05, 0x4d, 0x73, 0x67, 0x49, 0x44, 0x12, 0x10, 0x0a, 0x0c, 0x44, 0x65, 0x66,
	0x61, 0x75, 0x6c, 0x74, 0x4d, 0x73, 0x67, 0x49, 0x44, 0x10, 0x00, 0x12, 0x25, 0x0a, 0x1f, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x72, 0x79, 0x10, 0xca,
	0x84, 0x3d, 0x12, 0x29, 0x0a, 0x23, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x54, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x10, 0xd4, 0x84, 0x3d, 0x12, 0x29, 0x0a,
	0x23, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x72, 0x6d, 0x10, 0xde, 0x84, 0x3d, 0x12, 0x2d, 0x0a, 0x27, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e,
	0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x52,
	0x65, 0x73, 0x70, 0x10, 0xe8, 0x84, 0x3d, 0x12, 0x28, 0x0a, 0x22, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x76,
	0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x10, 0xf2, 0x84,
	0x3d, 0x12, 0x2c, 0x0a, 0x26, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x67, 0x69, 0x73,
	0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x10, 0xfc, 0x84, 0x3d, 0x12,
	0x1d, 0x0a, 0x17, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x64, 0x65, 0x54, 0x72, 0x79, 0x10, 0x86, 0x85, 0x3d, 0x12, 0x21,
	0x0a, 0x1b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x43, 0x6f, 0x64, 0x65, 0x54, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x10, 0x90, 0x85,
	0x3d, 0x12, 0x21, 0x0a, 0x1b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x76, 0x69, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x64, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d,
	0x10, 0x9a, 0x85, 0x3d, 0x12, 0x25, 0x0a, 0x1f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x49, 0x6e,
	0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x64, 0x65, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x72, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x10, 0xa4, 0x85, 0x3d, 0x12, 0x20, 0x0a, 0x1a, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43,
	0x6f, 0x64, 0x65, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x10, 0xae, 0x85, 0x3d, 0x12, 0x24, 0x0a,
	0x1e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x43, 0x6f, 0x64, 0x65, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x10,
	0xb8, 0x85, 0x3d, 0x12, 0x13, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65,
	0x72, 0x54, 0x72, 0x79, 0x10, 0xc2, 0x85, 0x3d, 0x12, 0x17, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x54, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x10, 0xcc, 0x85,
	0x3d, 0x12, 0x17, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x10, 0xd6, 0x85, 0x3d, 0x12, 0x1b, 0x0a, 0x15, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x52,
	0x65, 0x73, 0x70, 0x10, 0xe0, 0x85, 0x3d, 0x12, 0x16, 0x0a, 0x10, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x55, 0x73, 0x65, 0x72, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x10, 0xea, 0x85, 0x3d, 0x12,
	0x1a, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x43, 0x61, 0x6e,
	0x63, 0x65, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x10, 0xf4, 0x85, 0x3d, 0x2a, 0x40, 0x0a, 0x08, 0x4d,
	0x73, 0x67, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x13, 0x0a, 0x0f, 0x44, 0x65, 0x66, 0x61, 0x75,
	0x6c, 0x74, 0x4d, 0x73, 0x67, 0x53, 0x74, 0x61, 0x74, 0x65, 0x10, 0x00, 0x12, 0x10, 0x0a, 0x0c,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x10, 0x0a, 0x12, 0x0d,
	0x0a, 0x09, 0x53, 0x74, 0x61, 0x74, 0x65, 0x46, 0x61, 0x69, 0x6c, 0x10, 0x14, 0x42, 0x35, 0x5a,
	0x33, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4e, 0x70, 0x6f, 0x6f,
	0x6c, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x2f, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
=======
	0x92, 0x01, 0x0a, 0x05, 0x4d, 0x73, 0x67, 0x49, 0x44, 0x12, 0x10, 0x0a, 0x0c, 0x44, 0x65, 0x66,
=======
	0x97, 0x02, 0x0a, 0x05, 0x4d, 0x73, 0x67, 0x49, 0x44, 0x12, 0x10, 0x0a, 0x0c, 0x44, 0x65, 0x66,
>>>>>>> Add response
	0x61, 0x75, 0x6c, 0x74, 0x4d, 0x73, 0x67, 0x49, 0x44, 0x10, 0x00, 0x12, 0x24, 0x0a, 0x1f, 0x43,
=======
	0x89, 0x05, 0x0a, 0x05, 0x4d, 0x73, 0x67, 0x49, 0x44, 0x12, 0x10, 0x0a, 0x0c, 0x44, 0x65, 0x66,
	0x61, 0x75, 0x6c, 0x74, 0x4d, 0x73, 0x67, 0x49, 0x44, 0x10, 0x00, 0x12, 0x25, 0x0a, 0x1f, 0x43,
>>>>>>> Add create user and create invitation code tcc message
	0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x72, 0x79, 0x10, 0xca,
	0x84, 0x3d, 0x12, 0x29, 0x0a, 0x23, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x54, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x10, 0xd4, 0x84, 0x3d, 0x12, 0x29, 0x0a,
	0x23, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x72, 0x6d, 0x10, 0xde, 0x84, 0x3d, 0x12, 0x2d, 0x0a, 0x27, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e,
	0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x52,
	0x65, 0x73, 0x70, 0x10, 0xe8, 0x84, 0x3d, 0x12, 0x28, 0x0a, 0x22, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x76,
	0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x10, 0xf2, 0x84,
	0x3d, 0x12, 0x2c, 0x0a, 0x26, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x67, 0x69, 0x73,
	0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f,
<<<<<<< HEAD
	0x6e, 0x54, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x10, 0xa4, 0x4e, 0x12, 0x28, 0x0a, 0x23, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f,
<<<<<<< HEAD
	0x6e, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x61, 0x6e, 0x63, 0x65,
<<<<<<< HEAD
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
=======
	0x6c, 0x10, 0xae, 0x4e, 0x2a, 0x40, 0x0a, 0x08, 0x4d, 0x73, 0x67, 0x53, 0x74, 0x61, 0x74, 0x65,
	0x12, 0x13, 0x0a, 0x0f, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x4d, 0x73, 0x67, 0x53, 0x74,
	0x61, 0x74, 0x65, 0x10, 0x00, 0x12, 0x10, 0x0a, 0x0c, 0x53, 0x74, 0x61, 0x74, 0x65, 0x53, 0x75,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x10, 0x0a, 0x12, 0x0d, 0x0a, 0x09, 0x53, 0x74, 0x61, 0x74, 0x65,
	0x46, 0x61, 0x69, 0x6c, 0x10, 0x14, 0x42, 0x35, 0x5a, 0x33, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4e, 0x70, 0x6f, 0x6f, 0x6c, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f,
	0x72, 0x6d, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x6e, 0x70, 0x6f, 0x6f, 0x6c,
	0x2f, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
>>>>>>> Refactor name
=======
	0x6e, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x72, 0x6d, 0x10, 0xae, 0x4e, 0x12, 0x2c, 0x0a, 0x27, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x76, 0x69, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x52, 0x65, 0x73, 0x70,
	0x10, 0xb8, 0x4e, 0x12, 0x27, 0x0a, 0x22, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x10, 0xc2, 0x4e, 0x12, 0x2b, 0x0a, 0x26,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x61, 0x6e, 0x63,
	0x65, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x10, 0xcc, 0x4e, 0x2a, 0x40, 0x0a, 0x08, 0x4d, 0x73, 0x67,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x13, 0x0a, 0x0f, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74,
	0x4d, 0x73, 0x67, 0x53, 0x74, 0x61, 0x74, 0x65, 0x10, 0x00, 0x12, 0x10, 0x0a, 0x0c, 0x53, 0x74,
	0x61, 0x74, 0x65, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x10, 0x0a, 0x12, 0x0d, 0x0a, 0x09,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x46, 0x61, 0x69, 0x6c, 0x10, 0x14, 0x42, 0x35, 0x5a, 0x33, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4e, 0x70, 0x6f, 0x6f, 0x6c, 0x50,
	0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f,
	0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f,
	0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
>>>>>>> Add response
=======
	0x6e, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x10, 0xfc, 0x84, 0x3d, 0x12,
	0x1d, 0x0a, 0x17, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x64, 0x65, 0x54, 0x72, 0x79, 0x10, 0x86, 0x85, 0x3d, 0x12, 0x21,
	0x0a, 0x1b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x43, 0x6f, 0x64, 0x65, 0x54, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x10, 0x90, 0x85,
	0x3d, 0x12, 0x21, 0x0a, 0x1b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x76, 0x69, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x64, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d,
	0x10, 0x9a, 0x85, 0x3d, 0x12, 0x25, 0x0a, 0x1f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x49, 0x6e,
	0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x64, 0x65, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x72, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x10, 0xa4, 0x85, 0x3d, 0x12, 0x20, 0x0a, 0x1a, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43,
	0x6f, 0x64, 0x65, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x10, 0xae, 0x85, 0x3d, 0x12, 0x24, 0x0a,
	0x1e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x43, 0x6f, 0x64, 0x65, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x10,
	0xb8, 0x85, 0x3d, 0x12, 0x13, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65,
	0x72, 0x54, 0x72, 0x79, 0x10, 0xc2, 0x85, 0x3d, 0x12, 0x17, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x54, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x10, 0xcc, 0x85,
	0x3d, 0x12, 0x17, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x10, 0xd6, 0x85, 0x3d, 0x12, 0x1b, 0x0a, 0x15, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x52,
	0x65, 0x73, 0x70, 0x10, 0xe0, 0x85, 0x3d, 0x12, 0x16, 0x0a, 0x10, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x55, 0x73, 0x65, 0x72, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x10, 0xea, 0x85, 0x3d, 0x12,
	0x1a, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x43, 0x61, 0x6e,
	0x63, 0x65, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x10, 0xf4, 0x85, 0x3d, 0x2a, 0x40, 0x0a, 0x08, 0x4d,
	0x73, 0x67, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x13, 0x0a, 0x0f, 0x44, 0x65, 0x66, 0x61, 0x75,
	0x6c, 0x74, 0x4d, 0x73, 0x67, 0x53, 0x74, 0x61, 0x74, 0x65, 0x10, 0x00, 0x12, 0x10, 0x0a, 0x0c,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x10, 0x0a, 0x12, 0x0d,
	0x0a, 0x09, 0x53, 0x74, 0x61, 0x74, 0x65, 0x46, 0x61, 0x69, 0x6c, 0x10, 0x14, 0x42, 0x35, 0x5a,
	0x33, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4e, 0x70, 0x6f, 0x6f,
	0x6c, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x2f, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
>>>>>>> Add create user and create invitation code tcc message
=======
	0x2e, 0x0a, 0x05, 0x4d, 0x73, 0x67, 0x49, 0x44, 0x12, 0x10, 0x0a, 0x0c, 0x44, 0x65, 0x66, 0x61,
	0x75, 0x6c, 0x74, 0x4d, 0x73, 0x67, 0x49, 0x44, 0x10, 0x00, 0x12, 0x13, 0x0a, 0x0f, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x10, 0x0a, 0x2a,
	0x40, 0x0a, 0x08, 0x4d, 0x73, 0x67, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x13, 0x0a, 0x0f, 0x44,
	0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x4d, 0x73, 0x67, 0x53, 0x74, 0x61, 0x74, 0x65, 0x10, 0x00,
	0x12, 0x10, 0x0a, 0x0c, 0x53, 0x74, 0x61, 0x74, 0x65, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x10, 0x0a, 0x12, 0x0d, 0x0a, 0x09, 0x53, 0x74, 0x61, 0x74, 0x65, 0x46, 0x61, 0x69, 0x6c, 0x10,
	0x14, 0x42, 0x35, 0x5a, 0x33, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x4e, 0x70, 0x6f, 0x6f, 0x6c, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x62, 0x61, 0x73, 0x65,
	0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
>>>>>>> Support create coupon
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
	(MsgID)(0),    // 0: basetypes.v1.MsgID
	(MsgState)(0), // 1: basetypes.v1.MsgState
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
