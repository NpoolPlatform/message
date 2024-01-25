// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.18.1
// source: npool/basetypes/v1/usedfor.proto

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

type UsedFor int32

const (
	UsedFor_DefaultUsedFor         UsedFor = 0
	UsedFor_Signup                 UsedFor = 10
	UsedFor_Signin                 UsedFor = 20
	UsedFor_Update                 UsedFor = 30
	UsedFor_Contact                UsedFor = 40
	UsedFor_SetWithdrawAddress     UsedFor = 50
	UsedFor_Withdraw               UsedFor = 60
	UsedFor_CreateInvitationCode   UsedFor = 70
	UsedFor_SetCommission          UsedFor = 80
	UsedFor_SetTransferTargetUser  UsedFor = 90
	UsedFor_Transfer               UsedFor = 100
	UsedFor_WithdrawalRequest      UsedFor = 110
	UsedFor_WithdrawalCompleted    UsedFor = 120
	UsedFor_DepositReceived        UsedFor = 130
	UsedFor_KYCApproved            UsedFor = 140
	UsedFor_KYCRejected            UsedFor = 150
	UsedFor_Purchase               UsedFor = 160
	UsedFor_AffiliatePurchase      UsedFor = 170
	UsedFor_AffiliateSignup        UsedFor = 180
	UsedFor_GoodBenefit1           UsedFor = 190
	UsedFor_NewLogin               UsedFor = 200
	UsedFor_UpdatePassword         UsedFor = 210
	UsedFor_UpdateEmail            UsedFor = 220
	UsedFor_UpdateMobile           UsedFor = 230
	UsedFor_UpdateGoogleAuth       UsedFor = 240
	UsedFor_OrderCompleted         UsedFor = 250
	UsedFor_OrderChildsRenewNotify UsedFor = 260
	UsedFor_OrderChildsRenew       UsedFor = 270
	UsedFor_Announcement           UsedFor = 100000
	UsedFor_ProductPage            UsedFor = 200000
)

// Enum value maps for UsedFor.
var (
	UsedFor_name = map[int32]string{
		0:      "DefaultUsedFor",
		10:     "Signup",
		20:     "Signin",
		30:     "Update",
		40:     "Contact",
		50:     "SetWithdrawAddress",
		60:     "Withdraw",
		70:     "CreateInvitationCode",
		80:     "SetCommission",
		90:     "SetTransferTargetUser",
		100:    "Transfer",
		110:    "WithdrawalRequest",
		120:    "WithdrawalCompleted",
		130:    "DepositReceived",
		140:    "KYCApproved",
		150:    "KYCRejected",
		160:    "Purchase",
		170:    "AffiliatePurchase",
		180:    "AffiliateSignup",
		190:    "GoodBenefit1",
		200:    "NewLogin",
		210:    "UpdatePassword",
		220:    "UpdateEmail",
		230:    "UpdateMobile",
		240:    "UpdateGoogleAuth",
		250:    "OrderCompleted",
		260:    "OrderChildsRenewNotify",
		270:    "OrderChildsRenew",
		100000: "Announcement",
		200000: "ProductPage",
	}
	UsedFor_value = map[string]int32{
		"DefaultUsedFor":         0,
		"Signup":                 10,
		"Signin":                 20,
		"Update":                 30,
		"Contact":                40,
		"SetWithdrawAddress":     50,
		"Withdraw":               60,
		"CreateInvitationCode":   70,
		"SetCommission":          80,
		"SetTransferTargetUser":  90,
		"Transfer":               100,
		"WithdrawalRequest":      110,
		"WithdrawalCompleted":    120,
		"DepositReceived":        130,
		"KYCApproved":            140,
		"KYCRejected":            150,
		"Purchase":               160,
		"AffiliatePurchase":      170,
		"AffiliateSignup":        180,
		"GoodBenefit1":           190,
		"NewLogin":               200,
		"UpdatePassword":         210,
		"UpdateEmail":            220,
		"UpdateMobile":           230,
		"UpdateGoogleAuth":       240,
		"OrderCompleted":         250,
		"OrderChildsRenewNotify": 260,
		"OrderChildsRenew":       270,
		"Announcement":           100000,
		"ProductPage":            200000,
	}
)

func (x UsedFor) Enum() *UsedFor {
	p := new(UsedFor)
	*p = x
	return p
}

func (x UsedFor) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (UsedFor) Descriptor() protoreflect.EnumDescriptor {
	return file_npool_basetypes_v1_usedfor_proto_enumTypes[0].Descriptor()
}

func (UsedFor) Type() protoreflect.EnumType {
	return &file_npool_basetypes_v1_usedfor_proto_enumTypes[0]
}

func (x UsedFor) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use UsedFor.Descriptor instead.
func (UsedFor) EnumDescriptor() ([]byte, []int) {
	return file_npool_basetypes_v1_usedfor_proto_rawDescGZIP(), []int{0}
}

var File_npool_basetypes_v1_usedfor_proto protoreflect.FileDescriptor

var file_npool_basetypes_v1_usedfor_proto_rawDesc = []byte{
	0x0a, 0x20, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x64, 0x66, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0c, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31,
	0x2a, 0xd4, 0x04, 0x0a, 0x07, 0x55, 0x73, 0x65, 0x64, 0x46, 0x6f, 0x72, 0x12, 0x12, 0x0a, 0x0e,
	0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x55, 0x73, 0x65, 0x64, 0x46, 0x6f, 0x72, 0x10, 0x00,
	0x12, 0x0a, 0x0a, 0x06, 0x53, 0x69, 0x67, 0x6e, 0x75, 0x70, 0x10, 0x0a, 0x12, 0x0a, 0x0a, 0x06,
	0x53, 0x69, 0x67, 0x6e, 0x69, 0x6e, 0x10, 0x14, 0x12, 0x0a, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x10, 0x1e, 0x12, 0x0b, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x10,
	0x28, 0x12, 0x16, 0x0a, 0x12, 0x53, 0x65, 0x74, 0x57, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77,
	0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x10, 0x32, 0x12, 0x0c, 0x0a, 0x08, 0x57, 0x69, 0x74,
	0x68, 0x64, 0x72, 0x61, 0x77, 0x10, 0x3c, 0x12, 0x18, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x64, 0x65, 0x10,
	0x46, 0x12, 0x11, 0x0a, 0x0d, 0x53, 0x65, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x10, 0x50, 0x12, 0x19, 0x0a, 0x15, 0x53, 0x65, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x66, 0x65, 0x72, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x10, 0x5a, 0x12,
	0x0c, 0x0a, 0x08, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x10, 0x64, 0x12, 0x15, 0x0a,
	0x11, 0x57, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x61, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x10, 0x6e, 0x12, 0x17, 0x0a, 0x13, 0x57, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77,
	0x61, 0x6c, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x10, 0x78, 0x12, 0x14, 0x0a,
	0x0f, 0x44, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x64,
	0x10, 0x82, 0x01, 0x12, 0x10, 0x0a, 0x0b, 0x4b, 0x59, 0x43, 0x41, 0x70, 0x70, 0x72, 0x6f, 0x76,
	0x65, 0x64, 0x10, 0x8c, 0x01, 0x12, 0x10, 0x0a, 0x0b, 0x4b, 0x59, 0x43, 0x52, 0x65, 0x6a, 0x65,
	0x63, 0x74, 0x65, 0x64, 0x10, 0x96, 0x01, 0x12, 0x0d, 0x0a, 0x08, 0x50, 0x75, 0x72, 0x63, 0x68,
	0x61, 0x73, 0x65, 0x10, 0xa0, 0x01, 0x12, 0x16, 0x0a, 0x11, 0x41, 0x66, 0x66, 0x69, 0x6c, 0x69,
	0x61, 0x74, 0x65, 0x50, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x10, 0xaa, 0x01, 0x12, 0x14,
	0x0a, 0x0f, 0x41, 0x66, 0x66, 0x69, 0x6c, 0x69, 0x61, 0x74, 0x65, 0x53, 0x69, 0x67, 0x6e, 0x75,
	0x70, 0x10, 0xb4, 0x01, 0x12, 0x11, 0x0a, 0x0c, 0x47, 0x6f, 0x6f, 0x64, 0x42, 0x65, 0x6e, 0x65,
	0x66, 0x69, 0x74, 0x31, 0x10, 0xbe, 0x01, 0x12, 0x0d, 0x0a, 0x08, 0x4e, 0x65, 0x77, 0x4c, 0x6f,
	0x67, 0x69, 0x6e, 0x10, 0xc8, 0x01, 0x12, 0x13, 0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x10, 0xd2, 0x01, 0x12, 0x10, 0x0a, 0x0b, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x10, 0xdc, 0x01, 0x12, 0x11, 0x0a,
	0x0c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x10, 0xe6, 0x01,
	0x12, 0x15, 0x0a, 0x10, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x41, 0x75, 0x74, 0x68, 0x10, 0xf0, 0x01, 0x12, 0x13, 0x0a, 0x0e, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x10, 0xfa, 0x01, 0x12, 0x1b, 0x0a, 0x16,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x43, 0x68, 0x69, 0x6c, 0x64, 0x73, 0x52, 0x65, 0x6e, 0x65, 0x77,
	0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x10, 0x84, 0x02, 0x12, 0x15, 0x0a, 0x10, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x43, 0x68, 0x69, 0x6c, 0x64, 0x73, 0x52, 0x65, 0x6e, 0x65, 0x77, 0x10, 0x8e, 0x02,
	0x12, 0x12, 0x0a, 0x0c, 0x41, 0x6e, 0x6e, 0x6f, 0x75, 0x6e, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74,
	0x10, 0xa0, 0x8d, 0x06, 0x12, 0x11, 0x0a, 0x0b, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x50,
	0x61, 0x67, 0x65, 0x10, 0xc0, 0x9a, 0x0c, 0x42, 0x35, 0x5a, 0x33, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4e, 0x70, 0x6f, 0x6f, 0x6c, 0x50, 0x6c, 0x61, 0x74, 0x66,
	0x6f, 0x72, 0x6d, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x6e, 0x70, 0x6f, 0x6f,
	0x6c, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_npool_basetypes_v1_usedfor_proto_rawDescOnce sync.Once
	file_npool_basetypes_v1_usedfor_proto_rawDescData = file_npool_basetypes_v1_usedfor_proto_rawDesc
)

func file_npool_basetypes_v1_usedfor_proto_rawDescGZIP() []byte {
	file_npool_basetypes_v1_usedfor_proto_rawDescOnce.Do(func() {
		file_npool_basetypes_v1_usedfor_proto_rawDescData = protoimpl.X.CompressGZIP(file_npool_basetypes_v1_usedfor_proto_rawDescData)
	})
	return file_npool_basetypes_v1_usedfor_proto_rawDescData
}

var file_npool_basetypes_v1_usedfor_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_npool_basetypes_v1_usedfor_proto_goTypes = []interface{}{
	(UsedFor)(0), // 0: basetypes.v1.UsedFor
}
var file_npool_basetypes_v1_usedfor_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_npool_basetypes_v1_usedfor_proto_init() }
func file_npool_basetypes_v1_usedfor_proto_init() {
	if File_npool_basetypes_v1_usedfor_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_npool_basetypes_v1_usedfor_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_npool_basetypes_v1_usedfor_proto_goTypes,
		DependencyIndexes: file_npool_basetypes_v1_usedfor_proto_depIdxs,
		EnumInfos:         file_npool_basetypes_v1_usedfor_proto_enumTypes,
	}.Build()
	File_npool_basetypes_v1_usedfor_proto = out.File
	file_npool_basetypes_v1_usedfor_proto_rawDesc = nil
	file_npool_basetypes_v1_usedfor_proto_goTypes = nil
	file_npool_basetypes_v1_usedfor_proto_depIdxs = nil
}
