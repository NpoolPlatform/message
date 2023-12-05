// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.18.1
// source: npool/basetypes/good/v1/enums.proto

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

type BenefitType int32

const (
	BenefitType_DefaultBenefitType BenefitType = 0
	// Platform receive benefit, then distribute benefit to user
	BenefitType_BenefitTypePlatform BenefitType = 10
	// User receive benefit from pool directly
	BenefitType_BenefitTypePool BenefitType = 20
	// User do not receive any online reward
	BenefitType_BenefitTypeOffline BenefitType = 30
)

// Enum value maps for BenefitType.
var (
	BenefitType_name = map[int32]string{
		0:  "DefaultBenefitType",
		10: "BenefitTypePlatform",
		20: "BenefitTypePool",
		30: "BenefitTypeOffline",
	}
	BenefitType_value = map[string]int32{
		"DefaultBenefitType":  0,
		"BenefitTypePlatform": 10,
		"BenefitTypePool":     20,
		"BenefitTypeOffline":  30,
	}
)

func (x BenefitType) Enum() *BenefitType {
	p := new(BenefitType)
	*p = x
	return p
}

func (x BenefitType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (BenefitType) Descriptor() protoreflect.EnumDescriptor {
	return file_npool_basetypes_good_v1_enums_proto_enumTypes[0].Descriptor()
}

func (BenefitType) Type() protoreflect.EnumType {
	return &file_npool_basetypes_good_v1_enums_proto_enumTypes[0]
}

func (x BenefitType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use BenefitType.Descriptor instead.
func (BenefitType) EnumDescriptor() ([]byte, []int) {
	return file_npool_basetypes_good_v1_enums_proto_rawDescGZIP(), []int{0}
}

type GoodType int32

const (
	GoodType_DefaultGoodType     GoodType = 0
	GoodType_PowerRenting        GoodType = 10
	GoodType_MachineRenting      GoodType = 20
	GoodType_MachineHosting      GoodType = 30
	GoodType_TechniqueServiceFee GoodType = 1000
	GoodType_ElectricityFee      GoodType = 2000
)

// Enum value maps for GoodType.
var (
	GoodType_name = map[int32]string{
		0:    "DefaultGoodType",
		10:   "PowerRenting",
		20:   "MachineRenting",
		30:   "MachineHosting",
		1000: "TechniqueServiceFee",
		2000: "ElectricityFee",
	}
	GoodType_value = map[string]int32{
		"DefaultGoodType":     0,
		"PowerRenting":        10,
		"MachineRenting":      20,
		"MachineHosting":      30,
		"TechniqueServiceFee": 1000,
		"ElectricityFee":      2000,
	}
)

func (x GoodType) Enum() *GoodType {
	p := new(GoodType)
	*p = x
	return p
}

func (x GoodType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (GoodType) Descriptor() protoreflect.EnumDescriptor {
	return file_npool_basetypes_good_v1_enums_proto_enumTypes[1].Descriptor()
}

func (GoodType) Type() protoreflect.EnumType {
	return &file_npool_basetypes_good_v1_enums_proto_enumTypes[1]
}

func (x GoodType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use GoodType.Descriptor instead.
func (GoodType) EnumDescriptor() ([]byte, []int) {
	return file_npool_basetypes_good_v1_enums_proto_rawDescGZIP(), []int{1}
}

type BenefitState int32

const (
	BenefitState_DefaultBenefitState BenefitState = 0
	BenefitState_BenefitWait         BenefitState = 10
	BenefitState_BenefitTransferring BenefitState = 20
	// Good bookkeeping
	BenefitState_BenefitBookKeeping     BenefitState = 30
	BenefitState_BenefitUserBookKeeping BenefitState = 31
	BenefitState_BenefitDone            BenefitState = 40
	BenefitState_BenefitFail            BenefitState = 50
)

// Enum value maps for BenefitState.
var (
	BenefitState_name = map[int32]string{
		0:  "DefaultBenefitState",
		10: "BenefitWait",
		20: "BenefitTransferring",
		30: "BenefitBookKeeping",
		31: "BenefitUserBookKeeping",
		40: "BenefitDone",
		50: "BenefitFail",
	}
	BenefitState_value = map[string]int32{
		"DefaultBenefitState":    0,
		"BenefitWait":            10,
		"BenefitTransferring":    20,
		"BenefitBookKeeping":     30,
		"BenefitUserBookKeeping": 31,
		"BenefitDone":            40,
		"BenefitFail":            50,
	}
)

func (x BenefitState) Enum() *BenefitState {
	p := new(BenefitState)
	*p = x
	return p
}

func (x BenefitState) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (BenefitState) Descriptor() protoreflect.EnumDescriptor {
	return file_npool_basetypes_good_v1_enums_proto_enumTypes[2].Descriptor()
}

func (BenefitState) Type() protoreflect.EnumType {
	return &file_npool_basetypes_good_v1_enums_proto_enumTypes[2]
}

func (x BenefitState) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use BenefitState.Descriptor instead.
func (BenefitState) EnumDescriptor() ([]byte, []int) {
	return file_npool_basetypes_good_v1_enums_proto_rawDescGZIP(), []int{2}
}

type CancelMode int32

const (
	CancelMode_DefaultCancelMode        CancelMode = 0
	CancelMode_CancellableBeforeStart   CancelMode = 10
	CancelMode_CancellableBeforeBenefit CancelMode = 20
	CancelMode_Uncancellable            CancelMode = 30
)

// Enum value maps for CancelMode.
var (
	CancelMode_name = map[int32]string{
		0:  "DefaultCancelMode",
		10: "CancellableBeforeStart",
		20: "CancellableBeforeBenefit",
		30: "Uncancellable",
	}
	CancelMode_value = map[string]int32{
		"DefaultCancelMode":        0,
		"CancellableBeforeStart":   10,
		"CancellableBeforeBenefit": 20,
		"Uncancellable":            30,
	}
)

func (x CancelMode) Enum() *CancelMode {
	p := new(CancelMode)
	*p = x
	return p
}

func (x CancelMode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CancelMode) Descriptor() protoreflect.EnumDescriptor {
	return file_npool_basetypes_good_v1_enums_proto_enumTypes[3].Descriptor()
}

func (CancelMode) Type() protoreflect.EnumType {
	return &file_npool_basetypes_good_v1_enums_proto_enumTypes[3]
}

func (x CancelMode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CancelMode.Descriptor instead.
func (CancelMode) EnumDescriptor() ([]byte, []int) {
	return file_npool_basetypes_good_v1_enums_proto_rawDescGZIP(), []int{3}
}

type GoodLabel int32

const (
	GoodLabel_DefaultGoodLabel           GoodLabel = 0
	GoodLabel_GoodLabelPromotion         GoodLabel = 10
	GoodLabel_GoodLabelNoviceExclusive   GoodLabel = 20 // Newbie
	GoodLabel_GoodLabelInnovationStarter GoodLabel = 30 // First batch mining
	GoodLabel_GoodLabelLoyaltyExclusive  GoodLabel = 40 // User with action credis pass threshold
)

// Enum value maps for GoodLabel.
var (
	GoodLabel_name = map[int32]string{
		0:  "DefaultGoodLabel",
		10: "GoodLabelPromotion",
		20: "GoodLabelNoviceExclusive",
		30: "GoodLabelInnovationStarter",
		40: "GoodLabelLoyaltyExclusive",
	}
	GoodLabel_value = map[string]int32{
		"DefaultGoodLabel":           0,
		"GoodLabelPromotion":         10,
		"GoodLabelNoviceExclusive":   20,
		"GoodLabelInnovationStarter": 30,
		"GoodLabelLoyaltyExclusive":  40,
	}
)

func (x GoodLabel) Enum() *GoodLabel {
	p := new(GoodLabel)
	*p = x
	return p
}

func (x GoodLabel) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (GoodLabel) Descriptor() protoreflect.EnumDescriptor {
	return file_npool_basetypes_good_v1_enums_proto_enumTypes[4].Descriptor()
}

func (GoodLabel) Type() protoreflect.EnumType {
	return &file_npool_basetypes_good_v1_enums_proto_enumTypes[4]
}

func (x GoodLabel) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use GoodLabel.Descriptor instead.
func (GoodLabel) EnumDescriptor() ([]byte, []int) {
	return file_npool_basetypes_good_v1_enums_proto_rawDescGZIP(), []int{4}
}

type GoodTopMostType int32

const (
	GoodTopMostType_DefaultGoodTopMostType   GoodTopMostType = 0
	GoodTopMostType_TopMostPromotion         GoodTopMostType = 10
	GoodTopMostType_TopMostNoviceExclusive   GoodTopMostType = 20
	GoodTopMostType_TopMostBestOffer         GoodTopMostType = 30
	GoodTopMostType_TopMostInnovationStarter GoodTopMostType = 40
	GoodTopMostType_TopMostLoyaltyExclusive  GoodTopMostType = 50
)

// Enum value maps for GoodTopMostType.
var (
	GoodTopMostType_name = map[int32]string{
		0:  "DefaultGoodTopMostType",
		10: "TopMostPromotion",
		20: "TopMostNoviceExclusive",
		30: "TopMostBestOffer",
		40: "TopMostInnovationStarter",
		50: "TopMostLoyaltyExclusive",
	}
	GoodTopMostType_value = map[string]int32{
		"DefaultGoodTopMostType":   0,
		"TopMostPromotion":         10,
		"TopMostNoviceExclusive":   20,
		"TopMostBestOffer":         30,
		"TopMostInnovationStarter": 40,
		"TopMostLoyaltyExclusive":  50,
	}
)

func (x GoodTopMostType) Enum() *GoodTopMostType {
	p := new(GoodTopMostType)
	*p = x
	return p
}

func (x GoodTopMostType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (GoodTopMostType) Descriptor() protoreflect.EnumDescriptor {
	return file_npool_basetypes_good_v1_enums_proto_enumTypes[5].Descriptor()
}

func (GoodTopMostType) Type() protoreflect.EnumType {
	return &file_npool_basetypes_good_v1_enums_proto_enumTypes[5]
}

func (x GoodTopMostType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use GoodTopMostType.Descriptor instead.
func (GoodTopMostType) EnumDescriptor() ([]byte, []int) {
	return file_npool_basetypes_good_v1_enums_proto_rawDescGZIP(), []int{5}
}

type GoodStartMode int32

const (
	GoodStartMode_DefaultGoodStartMode   GoodStartMode = 0
	GoodStartMode_GoodStartModeTBD       GoodStartMode = 10
	GoodStartMode_GoodStartModeConfirmed GoodStartMode = 20
)

// Enum value maps for GoodStartMode.
var (
	GoodStartMode_name = map[int32]string{
		0:  "DefaultGoodStartMode",
		10: "GoodStartModeTBD",
		20: "GoodStartModeConfirmed",
	}
	GoodStartMode_value = map[string]int32{
		"DefaultGoodStartMode":   0,
		"GoodStartModeTBD":       10,
		"GoodStartModeConfirmed": 20,
	}
)

func (x GoodStartMode) Enum() *GoodStartMode {
	p := new(GoodStartMode)
	*p = x
	return p
}

func (x GoodStartMode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (GoodStartMode) Descriptor() protoreflect.EnumDescriptor {
	return file_npool_basetypes_good_v1_enums_proto_enumTypes[6].Descriptor()
}

func (GoodStartMode) Type() protoreflect.EnumType {
	return &file_npool_basetypes_good_v1_enums_proto_enumTypes[6]
}

func (x GoodStartMode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use GoodStartMode.Descriptor instead.
func (GoodStartMode) EnumDescriptor() ([]byte, []int) {
	return file_npool_basetypes_good_v1_enums_proto_rawDescGZIP(), []int{6}
}

type AppStockLockState int32

const (
	AppStockLockState_DefaultAppStockLockState AppStockLockState = 0
	AppStockLockState_AppStockLocked           AppStockLockState = 10
	AppStockLockState_AppStockWaitStart        AppStockLockState = 20
	AppStockLockState_AppStockInService        AppStockLockState = 30
	AppStockLockState_AppStockExpired          AppStockLockState = 40
	AppStockLockState_AppStockChargeBack       AppStockLockState = 50
	AppStockLockState_AppStockRollback         AppStockLockState = 60
	AppStockLockState_AppStockCanceled         AppStockLockState = 70
)

// Enum value maps for AppStockLockState.
var (
	AppStockLockState_name = map[int32]string{
		0:  "DefaultAppStockLockState",
		10: "AppStockLocked",
		20: "AppStockWaitStart",
		30: "AppStockInService",
		40: "AppStockExpired",
		50: "AppStockChargeBack",
		60: "AppStockRollback",
		70: "AppStockCanceled",
	}
	AppStockLockState_value = map[string]int32{
		"DefaultAppStockLockState": 0,
		"AppStockLocked":           10,
		"AppStockWaitStart":        20,
		"AppStockInService":        30,
		"AppStockExpired":          40,
		"AppStockChargeBack":       50,
		"AppStockRollback":         60,
		"AppStockCanceled":         70,
	}
)

func (x AppStockLockState) Enum() *AppStockLockState {
	p := new(AppStockLockState)
	*p = x
	return p
}

func (x AppStockLockState) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AppStockLockState) Descriptor() protoreflect.EnumDescriptor {
	return file_npool_basetypes_good_v1_enums_proto_enumTypes[7].Descriptor()
}

func (AppStockLockState) Type() protoreflect.EnumType {
	return &file_npool_basetypes_good_v1_enums_proto_enumTypes[7]
}

func (x AppStockLockState) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AppStockLockState.Descriptor instead.
func (AppStockLockState) EnumDescriptor() ([]byte, []int) {
	return file_npool_basetypes_good_v1_enums_proto_rawDescGZIP(), []int{7}
}

var File_npool_basetypes_good_v1_enums_proto protoreflect.FileDescriptor

var file_npool_basetypes_good_v1_enums_proto_rawDesc = []byte{
	0x0a, 0x23, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x64, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x2e, 0x67, 0x6f, 0x6f, 0x64, 0x2e, 0x76, 0x31, 0x2a, 0x6b, 0x0a, 0x0b, 0x42, 0x65, 0x6e, 0x65,
	0x66, 0x69, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x12, 0x44, 0x65, 0x66, 0x61, 0x75,
	0x6c, 0x74, 0x42, 0x65, 0x6e, 0x65, 0x66, 0x69, 0x74, 0x54, 0x79, 0x70, 0x65, 0x10, 0x00, 0x12,
	0x17, 0x0a, 0x13, 0x42, 0x65, 0x6e, 0x65, 0x66, 0x69, 0x74, 0x54, 0x79, 0x70, 0x65, 0x50, 0x6c,
	0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x10, 0x0a, 0x12, 0x13, 0x0a, 0x0f, 0x42, 0x65, 0x6e, 0x65,
	0x66, 0x69, 0x74, 0x54, 0x79, 0x70, 0x65, 0x50, 0x6f, 0x6f, 0x6c, 0x10, 0x14, 0x12, 0x16, 0x0a,
	0x12, 0x42, 0x65, 0x6e, 0x65, 0x66, 0x69, 0x74, 0x54, 0x79, 0x70, 0x65, 0x4f, 0x66, 0x66, 0x6c,
	0x69, 0x6e, 0x65, 0x10, 0x1e, 0x2a, 0x88, 0x01, 0x0a, 0x08, 0x47, 0x6f, 0x6f, 0x64, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x13, 0x0a, 0x0f, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x47, 0x6f, 0x6f,
	0x64, 0x54, 0x79, 0x70, 0x65, 0x10, 0x00, 0x12, 0x10, 0x0a, 0x0c, 0x50, 0x6f, 0x77, 0x65, 0x72,
	0x52, 0x65, 0x6e, 0x74, 0x69, 0x6e, 0x67, 0x10, 0x0a, 0x12, 0x12, 0x0a, 0x0e, 0x4d, 0x61, 0x63,
	0x68, 0x69, 0x6e, 0x65, 0x52, 0x65, 0x6e, 0x74, 0x69, 0x6e, 0x67, 0x10, 0x14, 0x12, 0x12, 0x0a,
	0x0e, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x48, 0x6f, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x10,
	0x1e, 0x12, 0x18, 0x0a, 0x13, 0x54, 0x65, 0x63, 0x68, 0x6e, 0x69, 0x71, 0x75, 0x65, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x46, 0x65, 0x65, 0x10, 0xe8, 0x07, 0x12, 0x13, 0x0a, 0x0e, 0x45,
	0x6c, 0x65, 0x63, 0x74, 0x72, 0x69, 0x63, 0x69, 0x74, 0x79, 0x46, 0x65, 0x65, 0x10, 0xd0, 0x0f,
	0x2a, 0xa7, 0x01, 0x0a, 0x0c, 0x42, 0x65, 0x6e, 0x65, 0x66, 0x69, 0x74, 0x53, 0x74, 0x61, 0x74,
	0x65, 0x12, 0x17, 0x0a, 0x13, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x42, 0x65, 0x6e, 0x65,
	0x66, 0x69, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x42, 0x65,
	0x6e, 0x65, 0x66, 0x69, 0x74, 0x57, 0x61, 0x69, 0x74, 0x10, 0x0a, 0x12, 0x17, 0x0a, 0x13, 0x42,
	0x65, 0x6e, 0x65, 0x66, 0x69, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x72, 0x69,
	0x6e, 0x67, 0x10, 0x14, 0x12, 0x16, 0x0a, 0x12, 0x42, 0x65, 0x6e, 0x65, 0x66, 0x69, 0x74, 0x42,
	0x6f, 0x6f, 0x6b, 0x4b, 0x65, 0x65, 0x70, 0x69, 0x6e, 0x67, 0x10, 0x1e, 0x12, 0x1a, 0x0a, 0x16,
	0x42, 0x65, 0x6e, 0x65, 0x66, 0x69, 0x74, 0x55, 0x73, 0x65, 0x72, 0x42, 0x6f, 0x6f, 0x6b, 0x4b,
	0x65, 0x65, 0x70, 0x69, 0x6e, 0x67, 0x10, 0x1f, 0x12, 0x0f, 0x0a, 0x0b, 0x42, 0x65, 0x6e, 0x65,
	0x66, 0x69, 0x74, 0x44, 0x6f, 0x6e, 0x65, 0x10, 0x28, 0x12, 0x0f, 0x0a, 0x0b, 0x42, 0x65, 0x6e,
	0x65, 0x66, 0x69, 0x74, 0x46, 0x61, 0x69, 0x6c, 0x10, 0x32, 0x2a, 0x70, 0x0a, 0x0a, 0x43, 0x61,
	0x6e, 0x63, 0x65, 0x6c, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x15, 0x0a, 0x11, 0x44, 0x65, 0x66, 0x61,
	0x75, 0x6c, 0x74, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x4d, 0x6f, 0x64, 0x65, 0x10, 0x00, 0x12,
	0x1a, 0x0a, 0x16, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x42, 0x65,
	0x66, 0x6f, 0x72, 0x65, 0x53, 0x74, 0x61, 0x72, 0x74, 0x10, 0x0a, 0x12, 0x1c, 0x0a, 0x18, 0x43,
	0x61, 0x6e, 0x63, 0x65, 0x6c, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x42, 0x65, 0x66, 0x6f, 0x72, 0x65,
	0x42, 0x65, 0x6e, 0x65, 0x66, 0x69, 0x74, 0x10, 0x14, 0x12, 0x11, 0x0a, 0x0d, 0x55, 0x6e, 0x63,
	0x61, 0x6e, 0x63, 0x65, 0x6c, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x10, 0x1e, 0x2a, 0x96, 0x01, 0x0a,
	0x09, 0x47, 0x6f, 0x6f, 0x64, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x12, 0x14, 0x0a, 0x10, 0x44, 0x65,
	0x66, 0x61, 0x75, 0x6c, 0x74, 0x47, 0x6f, 0x6f, 0x64, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x10, 0x00,
	0x12, 0x16, 0x0a, 0x12, 0x47, 0x6f, 0x6f, 0x64, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x50, 0x72, 0x6f,
	0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x10, 0x0a, 0x12, 0x1c, 0x0a, 0x18, 0x47, 0x6f, 0x6f, 0x64,
	0x4c, 0x61, 0x62, 0x65, 0x6c, 0x4e, 0x6f, 0x76, 0x69, 0x63, 0x65, 0x45, 0x78, 0x63, 0x6c, 0x75,
	0x73, 0x69, 0x76, 0x65, 0x10, 0x14, 0x12, 0x1e, 0x0a, 0x1a, 0x47, 0x6f, 0x6f, 0x64, 0x4c, 0x61,
	0x62, 0x65, 0x6c, 0x49, 0x6e, 0x6e, 0x6f, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61,
	0x72, 0x74, 0x65, 0x72, 0x10, 0x1e, 0x12, 0x1d, 0x0a, 0x19, 0x47, 0x6f, 0x6f, 0x64, 0x4c, 0x61,
	0x62, 0x65, 0x6c, 0x4c, 0x6f, 0x79, 0x61, 0x6c, 0x74, 0x79, 0x45, 0x78, 0x63, 0x6c, 0x75, 0x73,
	0x69, 0x76, 0x65, 0x10, 0x28, 0x2a, 0xb0, 0x01, 0x0a, 0x0f, 0x47, 0x6f, 0x6f, 0x64, 0x54, 0x6f,
	0x70, 0x4d, 0x6f, 0x73, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x16, 0x44, 0x65, 0x66,
	0x61, 0x75, 0x6c, 0x74, 0x47, 0x6f, 0x6f, 0x64, 0x54, 0x6f, 0x70, 0x4d, 0x6f, 0x73, 0x74, 0x54,
	0x79, 0x70, 0x65, 0x10, 0x00, 0x12, 0x14, 0x0a, 0x10, 0x54, 0x6f, 0x70, 0x4d, 0x6f, 0x73, 0x74,
	0x50, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x10, 0x0a, 0x12, 0x1a, 0x0a, 0x16, 0x54,
	0x6f, 0x70, 0x4d, 0x6f, 0x73, 0x74, 0x4e, 0x6f, 0x76, 0x69, 0x63, 0x65, 0x45, 0x78, 0x63, 0x6c,
	0x75, 0x73, 0x69, 0x76, 0x65, 0x10, 0x14, 0x12, 0x14, 0x0a, 0x10, 0x54, 0x6f, 0x70, 0x4d, 0x6f,
	0x73, 0x74, 0x42, 0x65, 0x73, 0x74, 0x4f, 0x66, 0x66, 0x65, 0x72, 0x10, 0x1e, 0x12, 0x1c, 0x0a,
	0x18, 0x54, 0x6f, 0x70, 0x4d, 0x6f, 0x73, 0x74, 0x49, 0x6e, 0x6e, 0x6f, 0x76, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x53, 0x74, 0x61, 0x72, 0x74, 0x65, 0x72, 0x10, 0x28, 0x12, 0x1b, 0x0a, 0x17, 0x54,
	0x6f, 0x70, 0x4d, 0x6f, 0x73, 0x74, 0x4c, 0x6f, 0x79, 0x61, 0x6c, 0x74, 0x79, 0x45, 0x78, 0x63,
	0x6c, 0x75, 0x73, 0x69, 0x76, 0x65, 0x10, 0x32, 0x2a, 0x5b, 0x0a, 0x0d, 0x47, 0x6f, 0x6f, 0x64,
	0x53, 0x74, 0x61, 0x72, 0x74, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x14, 0x44, 0x65, 0x66,
	0x61, 0x75, 0x6c, 0x74, 0x47, 0x6f, 0x6f, 0x64, 0x53, 0x74, 0x61, 0x72, 0x74, 0x4d, 0x6f, 0x64,
	0x65, 0x10, 0x00, 0x12, 0x14, 0x0a, 0x10, 0x47, 0x6f, 0x6f, 0x64, 0x53, 0x74, 0x61, 0x72, 0x74,
	0x4d, 0x6f, 0x64, 0x65, 0x54, 0x42, 0x44, 0x10, 0x0a, 0x12, 0x1a, 0x0a, 0x16, 0x47, 0x6f, 0x6f,
	0x64, 0x53, 0x74, 0x61, 0x72, 0x74, 0x4d, 0x6f, 0x64, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72,
	0x6d, 0x65, 0x64, 0x10, 0x14, 0x2a, 0xcc, 0x01, 0x0a, 0x11, 0x41, 0x70, 0x70, 0x53, 0x74, 0x6f,
	0x63, 0x6b, 0x4c, 0x6f, 0x63, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x1c, 0x0a, 0x18, 0x44,
	0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x41, 0x70, 0x70, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x4c, 0x6f,
	0x63, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x65, 0x10, 0x00, 0x12, 0x12, 0x0a, 0x0e, 0x41, 0x70, 0x70,
	0x53, 0x74, 0x6f, 0x63, 0x6b, 0x4c, 0x6f, 0x63, 0x6b, 0x65, 0x64, 0x10, 0x0a, 0x12, 0x15, 0x0a,
	0x11, 0x41, 0x70, 0x70, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x57, 0x61, 0x69, 0x74, 0x53, 0x74, 0x61,
	0x72, 0x74, 0x10, 0x14, 0x12, 0x15, 0x0a, 0x11, 0x41, 0x70, 0x70, 0x53, 0x74, 0x6f, 0x63, 0x6b,
	0x49, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x10, 0x1e, 0x12, 0x13, 0x0a, 0x0f, 0x41,
	0x70, 0x70, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x45, 0x78, 0x70, 0x69, 0x72, 0x65, 0x64, 0x10, 0x28,
	0x12, 0x16, 0x0a, 0x12, 0x41, 0x70, 0x70, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x43, 0x68, 0x61, 0x72,
	0x67, 0x65, 0x42, 0x61, 0x63, 0x6b, 0x10, 0x32, 0x12, 0x14, 0x0a, 0x10, 0x41, 0x70, 0x70, 0x53,
	0x74, 0x6f, 0x63, 0x6b, 0x52, 0x6f, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x10, 0x3c, 0x12, 0x14,
	0x0a, 0x10, 0x41, 0x70, 0x70, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c,
	0x65, 0x64, 0x10, 0x46, 0x42, 0x3a, 0x5a, 0x38, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x4e, 0x70, 0x6f, 0x6f, 0x6c, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d,
	0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x62,
	0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x64, 0x2f, 0x76, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_npool_basetypes_good_v1_enums_proto_rawDescOnce sync.Once
	file_npool_basetypes_good_v1_enums_proto_rawDescData = file_npool_basetypes_good_v1_enums_proto_rawDesc
)

func file_npool_basetypes_good_v1_enums_proto_rawDescGZIP() []byte {
	file_npool_basetypes_good_v1_enums_proto_rawDescOnce.Do(func() {
		file_npool_basetypes_good_v1_enums_proto_rawDescData = protoimpl.X.CompressGZIP(file_npool_basetypes_good_v1_enums_proto_rawDescData)
	})
	return file_npool_basetypes_good_v1_enums_proto_rawDescData
}

var file_npool_basetypes_good_v1_enums_proto_enumTypes = make([]protoimpl.EnumInfo, 8)
var file_npool_basetypes_good_v1_enums_proto_goTypes = []interface{}{
	(BenefitType)(0),       // 0: basetypes.good.v1.BenefitType
	(GoodType)(0),          // 1: basetypes.good.v1.GoodType
	(BenefitState)(0),      // 2: basetypes.good.v1.BenefitState
	(CancelMode)(0),        // 3: basetypes.good.v1.CancelMode
	(GoodLabel)(0),         // 4: basetypes.good.v1.GoodLabel
	(GoodTopMostType)(0),   // 5: basetypes.good.v1.GoodTopMostType
	(GoodStartMode)(0),     // 6: basetypes.good.v1.GoodStartMode
	(AppStockLockState)(0), // 7: basetypes.good.v1.AppStockLockState
}
var file_npool_basetypes_good_v1_enums_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_npool_basetypes_good_v1_enums_proto_init() }
func file_npool_basetypes_good_v1_enums_proto_init() {
	if File_npool_basetypes_good_v1_enums_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_npool_basetypes_good_v1_enums_proto_rawDesc,
			NumEnums:      8,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_npool_basetypes_good_v1_enums_proto_goTypes,
		DependencyIndexes: file_npool_basetypes_good_v1_enums_proto_depIdxs,
		EnumInfos:         file_npool_basetypes_good_v1_enums_proto_enumTypes,
	}.Build()
	File_npool_basetypes_good_v1_enums_proto = out.File
	file_npool_basetypes_good_v1_enums_proto_rawDesc = nil
	file_npool_basetypes_good_v1_enums_proto_goTypes = nil
	file_npool_basetypes_good_v1_enums_proto_depIdxs = nil
}
