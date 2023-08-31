// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.18.1
// source: npool/basetypes/order/v1/enums.proto

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

type OrderType int32

const (
	OrderType_DefaultOrderType OrderType = 0
	OrderType_Normal           OrderType = 10
	OrderType_Offline          OrderType = 20
	OrderType_Airdrop          OrderType = 30
)

// Enum value maps for OrderType.
var (
	OrderType_name = map[int32]string{
		0:  "DefaultOrderType",
		10: "Normal",
		20: "Offline",
		30: "Airdrop",
	}
	OrderType_value = map[string]int32{
		"DefaultOrderType": 0,
		"Normal":           10,
		"Offline":          20,
		"Airdrop":          30,
	}
)

func (x OrderType) Enum() *OrderType {
	p := new(OrderType)
	*p = x
	return p
}

func (x OrderType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (OrderType) Descriptor() protoreflect.EnumDescriptor {
	return file_npool_basetypes_order_v1_enums_proto_enumTypes[0].Descriptor()
}

func (OrderType) Type() protoreflect.EnumType {
	return &file_npool_basetypes_order_v1_enums_proto_enumTypes[0]
}

func (x OrderType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use OrderType.Descriptor instead.
func (OrderType) EnumDescriptor() ([]byte, []int) {
	return file_npool_basetypes_order_v1_enums_proto_rawDescGZIP(), []int{0}
}

type OrderState int32

const (
	OrderState_DefaultOrderState OrderState = 0
	// Common state
	OrderState_OrderStateWaitPayment  OrderState = 10
	OrderState_OrderStateCheckPayment OrderState = 20
	OrderState_OrderStatePaid         OrderState = 30
	OrderState_OrderStateInService    OrderState = 40
	// Timeout of order
	OrderState_OrderStatePaymentTimeout OrderState = 50 // -> OrderStatePreCancel
	// End or order
<<<<<<< HEAD
	OrderState_OrderStatePreCancel       OrderState = 60 // -> OrderStatePreCancelCheck
	OrderState_OrderStatePreCancelCheck  OrderState = 70 // -> OrderStateRestoreCanceledStock
	OrderState_OrderStatePreExpired      OrderState = 80 // -> OrderStatePreExpired
	OrderState_OrderStatePreExpiredCheck OrderState = 90 // -> OrderStateRestoreExpiredStock
	// Calculation
	OrderState_OrderStateRestoreExpiredStock        OrderState = 100 // -> OrderStateRestoreExpiredStockCheck
	OrderState_OrderStateRestoreExpiredStockCheck   OrderState = 110 // -> OrderStateExpired
	OrderState_OrderStateRestoreCanceledStock       OrderState = 120 // -> OrderStateReturnCanceledBalanceCheck
	OrderState_OrderStateRestoreCanceledStockCheck  OrderState = 130 // -> OrderStateReturnCanceledBalance
	OrderState_OrderStateReturnCalceledBalance      OrderState = 140 // -> OrderStateReturnCalceledBalanceCheck
	OrderState_OrderStateReturnCalceledBalanceCheck OrderState = 150 // -> OrderStateCanceled
	// End state
	OrderState_OrderStateCanceled OrderState = 160
	OrderState_OrderStateExpired  OrderState = 170
=======
	OrderState_OrderStatePreCancel  OrderState = 60 // -> OrderStateRestoreCanceledStock
	OrderState_OrderStatePreExpired OrderState = 70 // -> OrderStateRestoreExpiredStock
	// Calculation
	OrderState_OrderStateRestoreExpiredStock   OrderState = 80  // -> OrderStateExpired
	OrderState_OrderStateRestoreCanceledStock  OrderState = 90  // -> OrderStateReturnCanceledBalance
	OrderState_OrderStateReturnCalceledBalance OrderState = 100 // -> OrderStateCanceled
	// End state
	OrderState_OrderStateCanceled OrderState = 110
	OrderState_OrderStateExpired  OrderState = 120
>>>>>>> a68c66b1588db8e3a56afcd919ff25a3008745d6
)

// Enum value maps for OrderState.
var (
	OrderState_name = map[int32]string{
		0:   "DefaultOrderState",
		10:  "OrderStateWaitPayment",
		20:  "OrderStateCheckPayment",
		30:  "OrderStatePaid",
		40:  "OrderStateInService",
		50:  "OrderStatePaymentTimeout",
		60:  "OrderStatePreCancel",
<<<<<<< HEAD
		70:  "OrderStatePreCancelCheck",
		80:  "OrderStatePreExpired",
		90:  "OrderStatePreExpiredCheck",
		100: "OrderStateRestoreExpiredStock",
		110: "OrderStateRestoreExpiredStockCheck",
		120: "OrderStateRestoreCanceledStock",
		130: "OrderStateRestoreCanceledStockCheck",
		140: "OrderStateReturnCalceledBalance",
		150: "OrderStateReturnCalceledBalanceCheck",
		160: "OrderStateCanceled",
		170: "OrderStateExpired",
	}
	OrderState_value = map[string]int32{
		"DefaultOrderState":                    0,
		"OrderStateWaitPayment":                10,
		"OrderStateCheckPayment":               20,
		"OrderStatePaid":                       30,
		"OrderStateInService":                  40,
		"OrderStatePaymentTimeout":             50,
		"OrderStatePreCancel":                  60,
		"OrderStatePreCancelCheck":             70,
		"OrderStatePreExpired":                 80,
		"OrderStatePreExpiredCheck":            90,
		"OrderStateRestoreExpiredStock":        100,
		"OrderStateRestoreExpiredStockCheck":   110,
		"OrderStateRestoreCanceledStock":       120,
		"OrderStateRestoreCanceledStockCheck":  130,
		"OrderStateReturnCalceledBalance":      140,
		"OrderStateReturnCalceledBalanceCheck": 150,
		"OrderStateCanceled":                   160,
		"OrderStateExpired":                    170,
=======
		70:  "OrderStatePreExpired",
		80:  "OrderStateRestoreExpiredStock",
		90:  "OrderStateRestoreCanceledStock",
		100: "OrderStateReturnCalceledBalance",
		110: "OrderStateCanceled",
		120: "OrderStateExpired",
	}
	OrderState_value = map[string]int32{
		"DefaultOrderState":               0,
		"OrderStateWaitPayment":           10,
		"OrderStateCheckPayment":          20,
		"OrderStatePaid":                  30,
		"OrderStateInService":             40,
		"OrderStatePaymentTimeout":        50,
		"OrderStatePreCancel":             60,
		"OrderStatePreExpired":            70,
		"OrderStateRestoreExpiredStock":   80,
		"OrderStateRestoreCanceledStock":  90,
		"OrderStateReturnCalceledBalance": 100,
		"OrderStateCanceled":              110,
		"OrderStateExpired":               120,
>>>>>>> a68c66b1588db8e3a56afcd919ff25a3008745d6
	}
)

func (x OrderState) Enum() *OrderState {
	p := new(OrderState)
	*p = x
	return p
}

func (x OrderState) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (OrderState) Descriptor() protoreflect.EnumDescriptor {
	return file_npool_basetypes_order_v1_enums_proto_enumTypes[1].Descriptor()
}

func (OrderState) Type() protoreflect.EnumType {
	return &file_npool_basetypes_order_v1_enums_proto_enumTypes[1]
}

func (x OrderState) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use OrderState.Descriptor instead.
func (OrderState) EnumDescriptor() ([]byte, []int) {
	return file_npool_basetypes_order_v1_enums_proto_rawDescGZIP(), []int{1}
}

type PaymentState int32

const (
	PaymentState_DefaultPaymentState   PaymentState = 0
	PaymentState_PaymentStateWait      PaymentState = 10
	PaymentState_PaymentStateDone      PaymentState = 20
	PaymentState_PaymentStateCanceled  PaymentState = 30
	PaymentState_PaymentStateTimeout   PaymentState = 40
	PaymentState_PaymentStateNoPayment PaymentState = 1000
)

// Enum value maps for PaymentState.
var (
	PaymentState_name = map[int32]string{
		0:    "DefaultPaymentState",
		10:   "PaymentStateWait",
		20:   "PaymentStateDone",
		30:   "PaymentStateCanceled",
		40:   "PaymentStateTimeout",
		1000: "PaymentStateNoPayment",
	}
	PaymentState_value = map[string]int32{
		"DefaultPaymentState":   0,
		"PaymentStateWait":      10,
		"PaymentStateDone":      20,
		"PaymentStateCanceled":  30,
		"PaymentStateTimeout":   40,
		"PaymentStateNoPayment": 1000,
	}
)

func (x PaymentState) Enum() *PaymentState {
	p := new(PaymentState)
	*p = x
	return p
}

func (x PaymentState) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PaymentState) Descriptor() protoreflect.EnumDescriptor {
	return file_npool_basetypes_order_v1_enums_proto_enumTypes[2].Descriptor()
}

func (PaymentState) Type() protoreflect.EnumType {
	return &file_npool_basetypes_order_v1_enums_proto_enumTypes[2]
}

func (x PaymentState) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PaymentState.Descriptor instead.
func (PaymentState) EnumDescriptor() ([]byte, []int) {
	return file_npool_basetypes_order_v1_enums_proto_rawDescGZIP(), []int{2}
}

type InvestmentType int32

const (
	InvestmentType_DefaultInvestmentType InvestmentType = 0
	InvestmentType_UnionMining           InvestmentType = 10
	InvestmentType_FullPayment           InvestmentType = 20
)

// Enum value maps for InvestmentType.
var (
	InvestmentType_name = map[int32]string{
		0:  "DefaultInvestmentType",
		10: "UnionMining",
		20: "FullPayment",
	}
	InvestmentType_value = map[string]int32{
		"DefaultInvestmentType": 0,
		"UnionMining":           10,
		"FullPayment":           20,
	}
)

func (x InvestmentType) Enum() *InvestmentType {
	p := new(InvestmentType)
	*p = x
	return p
}

func (x InvestmentType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (InvestmentType) Descriptor() protoreflect.EnumDescriptor {
	return file_npool_basetypes_order_v1_enums_proto_enumTypes[3].Descriptor()
}

func (InvestmentType) Type() protoreflect.EnumType {
	return &file_npool_basetypes_order_v1_enums_proto_enumTypes[3]
}

func (x InvestmentType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use InvestmentType.Descriptor instead.
func (InvestmentType) EnumDescriptor() ([]byte, []int) {
	return file_npool_basetypes_order_v1_enums_proto_rawDescGZIP(), []int{3}
}

type PaymentType int32

const (
	PaymentType_DefaultPaymentType        PaymentType = 0
	PaymentType_PayWithBalanceOnly        PaymentType = 10
	PaymentType_PayWithTransferOnly       PaymentType = 20
	PaymentType_PayWithTransferAndBalance PaymentType = 30
	PaymentType_PayWithParentOrder        PaymentType = 40
	PaymentType_PayWithOffline            PaymentType = 1000
	PaymentType_PayWithNoPayment          PaymentType = 1010
)

// Enum value maps for PaymentType.
var (
	PaymentType_name = map[int32]string{
		0:    "DefaultPaymentType",
		10:   "PayWithBalanceOnly",
		20:   "PayWithTransferOnly",
		30:   "PayWithTransferAndBalance",
		40:   "PayWithParentOrder",
		1000: "PayWithOffline",
		1010: "PayWithNoPayment",
	}
	PaymentType_value = map[string]int32{
		"DefaultPaymentType":        0,
		"PayWithBalanceOnly":        10,
		"PayWithTransferOnly":       20,
		"PayWithTransferAndBalance": 30,
		"PayWithParentOrder":        40,
		"PayWithOffline":            1000,
		"PayWithNoPayment":          1010,
	}
)

func (x PaymentType) Enum() *PaymentType {
	p := new(PaymentType)
	*p = x
	return p
}

func (x PaymentType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PaymentType) Descriptor() protoreflect.EnumDescriptor {
	return file_npool_basetypes_order_v1_enums_proto_enumTypes[4].Descriptor()
}

func (PaymentType) Type() protoreflect.EnumType {
	return &file_npool_basetypes_order_v1_enums_proto_enumTypes[4]
}

func (x PaymentType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PaymentType.Descriptor instead.
func (PaymentType) EnumDescriptor() ([]byte, []int) {
	return file_npool_basetypes_order_v1_enums_proto_rawDescGZIP(), []int{4}
}

type CompensateType int32

const (
	CompensateType_DefaultCompensateType  CompensateType = 0
	CompensateType_CompensateMalfunction  CompensateType = 10
	CompensateType_CompensateWalfare      CompensateType = 20
	CompensateType_CompensateStarterDelay CompensateType = 30
)

// Enum value maps for CompensateType.
var (
	CompensateType_name = map[int32]string{
		0:  "DefaultCompensateType",
		10: "CompensateMalfunction",
		20: "CompensateWalfare",
		30: "CompensateStarterDelay",
	}
	CompensateType_value = map[string]int32{
		"DefaultCompensateType":  0,
		"CompensateMalfunction":  10,
		"CompensateWalfare":      20,
		"CompensateStarterDelay": 30,
	}
)

func (x CompensateType) Enum() *CompensateType {
	p := new(CompensateType)
	*p = x
	return p
}

func (x CompensateType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CompensateType) Descriptor() protoreflect.EnumDescriptor {
	return file_npool_basetypes_order_v1_enums_proto_enumTypes[5].Descriptor()
}

func (CompensateType) Type() protoreflect.EnumType {
	return &file_npool_basetypes_order_v1_enums_proto_enumTypes[5]
}

func (x CompensateType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CompensateType.Descriptor instead.
func (CompensateType) EnumDescriptor() ([]byte, []int) {
	return file_npool_basetypes_order_v1_enums_proto_rawDescGZIP(), []int{5}
}

type BenefitState int32

const (
	BenefitState_DefaultBenefitState BenefitState = 0
	BenefitState_BenefitWait         BenefitState = 10
	BenefitState_BenefitCalculated   BenefitState = 20
	BenefitState_BenefitDone         BenefitState = 30
)

// Enum value maps for BenefitState.
var (
	BenefitState_name = map[int32]string{
		0:  "DefaultBenefitState",
		10: "BenefitWait",
		20: "BenefitCalculated",
		30: "BenefitDone",
	}
	BenefitState_value = map[string]int32{
		"DefaultBenefitState": 0,
		"BenefitWait":         10,
		"BenefitCalculated":   20,
		"BenefitDone":         30,
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
	return file_npool_basetypes_order_v1_enums_proto_enumTypes[6].Descriptor()
}

func (BenefitState) Type() protoreflect.EnumType {
	return &file_npool_basetypes_order_v1_enums_proto_enumTypes[6]
}

func (x BenefitState) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use BenefitState.Descriptor instead.
func (BenefitState) EnumDescriptor() ([]byte, []int) {
	return file_npool_basetypes_order_v1_enums_proto_rawDescGZIP(), []int{6}
}

type OrderStartMode int32

const (
	OrderStartMode_DefaultOrderStartMode OrderStartMode = 0
	OrderStartMode_OrderStartTBD         OrderStartMode = 10
	OrderStartMode_OrderStartConfirmed   OrderStartMode = 20
)

// Enum value maps for OrderStartMode.
var (
	OrderStartMode_name = map[int32]string{
		0:  "DefaultOrderStartMode",
		10: "OrderStartTBD",
		20: "OrderStartConfirmed",
	}
	OrderStartMode_value = map[string]int32{
		"DefaultOrderStartMode": 0,
		"OrderStartTBD":         10,
		"OrderStartConfirmed":   20,
	}
)

func (x OrderStartMode) Enum() *OrderStartMode {
	p := new(OrderStartMode)
	*p = x
	return p
}

func (x OrderStartMode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (OrderStartMode) Descriptor() protoreflect.EnumDescriptor {
	return file_npool_basetypes_order_v1_enums_proto_enumTypes[7].Descriptor()
}

func (OrderStartMode) Type() protoreflect.EnumType {
	return &file_npool_basetypes_order_v1_enums_proto_enumTypes[7]
}

func (x OrderStartMode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use OrderStartMode.Descriptor instead.
func (OrderStartMode) EnumDescriptor() ([]byte, []int) {
	return file_npool_basetypes_order_v1_enums_proto_rawDescGZIP(), []int{7}
}

var File_npool_basetypes_order_v1_enums_proto protoreflect.FileDescriptor

var file_npool_basetypes_order_v1_enums_proto_rawDesc = []byte{
	0x0a, 0x24, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2a, 0x47, 0x0a, 0x09, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x10, 0x44, 0x65, 0x66, 0x61, 0x75,
	0x6c, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x10, 0x00, 0x12, 0x0a, 0x0a,
	0x06, 0x4e, 0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x10, 0x0a, 0x12, 0x0b, 0x0a, 0x07, 0x4f, 0x66, 0x66,
	0x6c, 0x69, 0x6e, 0x65, 0x10, 0x14, 0x12, 0x0b, 0x0a, 0x07, 0x41, 0x69, 0x72, 0x64, 0x72, 0x6f,
<<<<<<< HEAD
	0x70, 0x10, 0x1e, 0x2a, 0xb0, 0x04, 0x0a, 0x0a, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x74, 0x61,
=======
	0x70, 0x10, 0x1e, 0x2a, 0xf3, 0x02, 0x0a, 0x0a, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x74, 0x61,
>>>>>>> a68c66b1588db8e3a56afcd919ff25a3008745d6
	0x74, 0x65, 0x12, 0x15, 0x0a, 0x11, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x65, 0x10, 0x00, 0x12, 0x19, 0x0a, 0x15, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x65, 0x57, 0x61, 0x69, 0x74, 0x50, 0x61, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x10, 0x0a, 0x12, 0x1a, 0x0a, 0x16, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x74, 0x61,
	0x74, 0x65, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x10, 0x14,
	0x12, 0x12, 0x0a, 0x0e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x65, 0x50, 0x61,
	0x69, 0x64, 0x10, 0x1e, 0x12, 0x17, 0x0a, 0x13, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x74, 0x61,
	0x74, 0x65, 0x49, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x10, 0x28, 0x12, 0x1c, 0x0a,
	0x18, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x65, 0x50, 0x61, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x10, 0x32, 0x12, 0x17, 0x0a, 0x13, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x65, 0x50, 0x72, 0x65, 0x43, 0x61, 0x6e, 0x63,
<<<<<<< HEAD
	0x65, 0x6c, 0x10, 0x3c, 0x12, 0x1c, 0x0a, 0x18, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x74, 0x61,
	0x74, 0x65, 0x50, 0x72, 0x65, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x43, 0x68, 0x65, 0x63, 0x6b,
	0x10, 0x46, 0x12, 0x18, 0x0a, 0x14, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x65,
	0x50, 0x72, 0x65, 0x45, 0x78, 0x70, 0x69, 0x72, 0x65, 0x64, 0x10, 0x50, 0x12, 0x1d, 0x0a, 0x19,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x65, 0x50, 0x72, 0x65, 0x45, 0x78, 0x70,
	0x69, 0x72, 0x65, 0x64, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x10, 0x5a, 0x12, 0x21, 0x0a, 0x1d, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x65,
	0x45, 0x78, 0x70, 0x69, 0x72, 0x65, 0x64, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x10, 0x64, 0x12, 0x26,
	0x0a, 0x22, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x74,
	0x6f, 0x72, 0x65, 0x45, 0x78, 0x70, 0x69, 0x72, 0x65, 0x64, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x43,
	0x68, 0x65, 0x63, 0x6b, 0x10, 0x6e, 0x12, 0x22, 0x0a, 0x1e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53,
	0x74, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x43, 0x61, 0x6e, 0x63, 0x65,
	0x6c, 0x65, 0x64, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x10, 0x78, 0x12, 0x28, 0x0a, 0x23, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x43,
	0x61, 0x6e, 0x63, 0x65, 0x6c, 0x65, 0x64, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x43, 0x68, 0x65, 0x63,
	0x6b, 0x10, 0x82, 0x01, 0x12, 0x24, 0x0a, 0x1f, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x74, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x43, 0x61, 0x6c, 0x63, 0x65, 0x6c, 0x65, 0x64,
	0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x10, 0x8c, 0x01, 0x12, 0x29, 0x0a, 0x24, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x43, 0x61,
	0x6c, 0x63, 0x65, 0x6c, 0x65, 0x64, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x43, 0x68, 0x65,
	0x63, 0x6b, 0x10, 0x96, 0x01, 0x12, 0x17, 0x0a, 0x12, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x74,
	0x61, 0x74, 0x65, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x65, 0x64, 0x10, 0xa0, 0x01, 0x12, 0x16,
	0x0a, 0x11, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x65, 0x45, 0x78, 0x70, 0x69,
	0x72, 0x65, 0x64, 0x10, 0xaa, 0x01, 0x2a, 0xa2, 0x01, 0x0a, 0x0c, 0x50, 0x61, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x17, 0x0a, 0x13, 0x44, 0x65, 0x66, 0x61, 0x75,
	0x6c, 0x74, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x10, 0x00,
	0x12, 0x14, 0x0a, 0x10, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65,
	0x57, 0x61, 0x69, 0x74, 0x10, 0x0a, 0x12, 0x14, 0x0a, 0x10, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e,
	0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x44, 0x6f, 0x6e, 0x65, 0x10, 0x14, 0x12, 0x18, 0x0a, 0x14,
	0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x43, 0x61, 0x6e, 0x63,
	0x65, 0x6c, 0x65, 0x64, 0x10, 0x1e, 0x12, 0x17, 0x0a, 0x13, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e,
	0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x10, 0x28, 0x12,
	0x1a, 0x0a, 0x15, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x4e,
	0x6f, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x10, 0xe8, 0x07, 0x2a, 0x4d, 0x0a, 0x0e, 0x49,
	0x6e, 0x76, 0x65, 0x73, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x19, 0x0a,
	0x15, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x49, 0x6e, 0x76, 0x65, 0x73, 0x74, 0x6d, 0x65,
	0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x6e, 0x69, 0x6f,
	0x6e, 0x4d, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x10, 0x0a, 0x12, 0x0f, 0x0a, 0x0b, 0x46, 0x75, 0x6c,
	0x6c, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x10, 0x14, 0x2a, 0xb9, 0x01, 0x0a, 0x0b, 0x50,
	0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x12, 0x44, 0x65,
	0x66, 0x61, 0x75, 0x6c, 0x74, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65,
	0x10, 0x00, 0x12, 0x16, 0x0a, 0x12, 0x50, 0x61, 0x79, 0x57, 0x69, 0x74, 0x68, 0x42, 0x61, 0x6c,
	0x61, 0x6e, 0x63, 0x65, 0x4f, 0x6e, 0x6c, 0x79, 0x10, 0x0a, 0x12, 0x17, 0x0a, 0x13, 0x50, 0x61,
	0x79, 0x57, 0x69, 0x74, 0x68, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x4f, 0x6e, 0x6c,
	0x79, 0x10, 0x14, 0x12, 0x1d, 0x0a, 0x19, 0x50, 0x61, 0x79, 0x57, 0x69, 0x74, 0x68, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x41, 0x6e, 0x64, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65,
	0x10, 0x1e, 0x12, 0x16, 0x0a, 0x12, 0x50, 0x61, 0x79, 0x57, 0x69, 0x74, 0x68, 0x50, 0x61, 0x72,
	0x65, 0x6e, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x10, 0x28, 0x12, 0x13, 0x0a, 0x0e, 0x50, 0x61,
	0x79, 0x57, 0x69, 0x74, 0x68, 0x4f, 0x66, 0x66, 0x6c, 0x69, 0x6e, 0x65, 0x10, 0xe8, 0x07, 0x12,
	0x15, 0x0a, 0x10, 0x50, 0x61, 0x79, 0x57, 0x69, 0x74, 0x68, 0x4e, 0x6f, 0x50, 0x61, 0x79, 0x6d,
	0x65, 0x6e, 0x74, 0x10, 0xf2, 0x07, 0x2a, 0x79, 0x0a, 0x0e, 0x43, 0x6f, 0x6d, 0x70, 0x65, 0x6e,
	0x73, 0x61, 0x74, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x19, 0x0a, 0x15, 0x44, 0x65, 0x66, 0x61,
	0x75, 0x6c, 0x74, 0x43, 0x6f, 0x6d, 0x70, 0x65, 0x6e, 0x73, 0x61, 0x74, 0x65, 0x54, 0x79, 0x70,
	0x65, 0x10, 0x00, 0x12, 0x19, 0x0a, 0x15, 0x43, 0x6f, 0x6d, 0x70, 0x65, 0x6e, 0x73, 0x61, 0x74,
	0x65, 0x4d, 0x61, 0x6c, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x10, 0x0a, 0x12, 0x15,
	0x0a, 0x11, 0x43, 0x6f, 0x6d, 0x70, 0x65, 0x6e, 0x73, 0x61, 0x74, 0x65, 0x57, 0x61, 0x6c, 0x66,
	0x61, 0x72, 0x65, 0x10, 0x14, 0x12, 0x1a, 0x0a, 0x16, 0x43, 0x6f, 0x6d, 0x70, 0x65, 0x6e, 0x73,
	0x61, 0x74, 0x65, 0x53, 0x74, 0x61, 0x72, 0x74, 0x65, 0x72, 0x44, 0x65, 0x6c, 0x61, 0x79, 0x10,
	0x1e, 0x2a, 0x60, 0x0a, 0x0c, 0x42, 0x65, 0x6e, 0x65, 0x66, 0x69, 0x74, 0x53, 0x74, 0x61, 0x74,
	0x65, 0x12, 0x17, 0x0a, 0x13, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x42, 0x65, 0x6e, 0x65,
	0x66, 0x69, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x42, 0x65,
	0x6e, 0x65, 0x66, 0x69, 0x74, 0x57, 0x61, 0x69, 0x74, 0x10, 0x0a, 0x12, 0x15, 0x0a, 0x11, 0x42,
	0x65, 0x6e, 0x65, 0x66, 0x69, 0x74, 0x43, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x64,
	0x10, 0x14, 0x12, 0x0f, 0x0a, 0x0b, 0x42, 0x65, 0x6e, 0x65, 0x66, 0x69, 0x74, 0x44, 0x6f, 0x6e,
	0x65, 0x10, 0x1e, 0x2a, 0x57, 0x0a, 0x0e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x74, 0x61, 0x72,
	0x74, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x19, 0x0a, 0x15, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x74, 0x61, 0x72, 0x74, 0x4d, 0x6f, 0x64, 0x65, 0x10, 0x00,
	0x12, 0x11, 0x0a, 0x0d, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x74, 0x61, 0x72, 0x74, 0x54, 0x42,
	0x44, 0x10, 0x0a, 0x12, 0x17, 0x0a, 0x13, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x74, 0x61, 0x72,
	0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x65, 0x64, 0x10, 0x14, 0x42, 0x3b, 0x5a, 0x39,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4e, 0x70, 0x6f, 0x6f, 0x6c,
	0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x2f, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
=======
	0x65, 0x6c, 0x10, 0x3c, 0x12, 0x18, 0x0a, 0x14, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x74, 0x61,
	0x74, 0x65, 0x50, 0x72, 0x65, 0x45, 0x78, 0x70, 0x69, 0x72, 0x65, 0x64, 0x10, 0x46, 0x12, 0x21,
	0x0a, 0x1d, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x74,
	0x6f, 0x72, 0x65, 0x45, 0x78, 0x70, 0x69, 0x72, 0x65, 0x64, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x10,
	0x50, 0x12, 0x22, 0x0a, 0x1e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x65, 0x64, 0x53, 0x74,
	0x6f, 0x63, 0x6b, 0x10, 0x5a, 0x12, 0x23, 0x0a, 0x1f, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x74,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x43, 0x61, 0x6c, 0x63, 0x65, 0x6c, 0x65,
	0x64, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x10, 0x64, 0x12, 0x16, 0x0a, 0x12, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x65, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x65, 0x64,
	0x10, 0x6e, 0x12, 0x15, 0x0a, 0x11, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x65,
	0x45, 0x78, 0x70, 0x69, 0x72, 0x65, 0x64, 0x10, 0x78, 0x2a, 0xa2, 0x01, 0x0a, 0x0c, 0x50, 0x61,
	0x79, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x17, 0x0a, 0x13, 0x44, 0x65,
	0x66, 0x61, 0x75, 0x6c, 0x74, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74,
	0x65, 0x10, 0x00, 0x12, 0x14, 0x0a, 0x10, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x74,
	0x61, 0x74, 0x65, 0x57, 0x61, 0x69, 0x74, 0x10, 0x0a, 0x12, 0x14, 0x0a, 0x10, 0x50, 0x61, 0x79,
	0x6d, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x44, 0x6f, 0x6e, 0x65, 0x10, 0x14, 0x12,
	0x18, 0x0a, 0x14, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x43,
	0x61, 0x6e, 0x63, 0x65, 0x6c, 0x65, 0x64, 0x10, 0x1e, 0x12, 0x17, 0x0a, 0x13, 0x50, 0x61, 0x79,
	0x6d, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74,
	0x10, 0x28, 0x12, 0x1a, 0x0a, 0x15, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61,
	0x74, 0x65, 0x4e, 0x6f, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x10, 0xe8, 0x07, 0x2a, 0x4d,
	0x0a, 0x0e, 0x49, 0x6e, 0x76, 0x65, 0x73, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x19, 0x0a, 0x15, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x49, 0x6e, 0x76, 0x65, 0x73,
	0x74, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x55,
	0x6e, 0x69, 0x6f, 0x6e, 0x4d, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x10, 0x0a, 0x12, 0x0f, 0x0a, 0x0b,
	0x46, 0x75, 0x6c, 0x6c, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x10, 0x14, 0x2a, 0xb9, 0x01,
	0x0a, 0x0b, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a,
	0x12, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x54,
	0x79, 0x70, 0x65, 0x10, 0x00, 0x12, 0x16, 0x0a, 0x12, 0x50, 0x61, 0x79, 0x57, 0x69, 0x74, 0x68,
	0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x4f, 0x6e, 0x6c, 0x79, 0x10, 0x0a, 0x12, 0x17, 0x0a,
	0x13, 0x50, 0x61, 0x79, 0x57, 0x69, 0x74, 0x68, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72,
	0x4f, 0x6e, 0x6c, 0x79, 0x10, 0x14, 0x12, 0x1d, 0x0a, 0x19, 0x50, 0x61, 0x79, 0x57, 0x69, 0x74,
	0x68, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x41, 0x6e, 0x64, 0x42, 0x61, 0x6c, 0x61,
	0x6e, 0x63, 0x65, 0x10, 0x1e, 0x12, 0x16, 0x0a, 0x12, 0x50, 0x61, 0x79, 0x57, 0x69, 0x74, 0x68,
	0x50, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x10, 0x28, 0x12, 0x13, 0x0a,
	0x0e, 0x50, 0x61, 0x79, 0x57, 0x69, 0x74, 0x68, 0x4f, 0x66, 0x66, 0x6c, 0x69, 0x6e, 0x65, 0x10,
	0xe8, 0x07, 0x12, 0x15, 0x0a, 0x10, 0x50, 0x61, 0x79, 0x57, 0x69, 0x74, 0x68, 0x4e, 0x6f, 0x50,
	0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x10, 0xf2, 0x07, 0x2a, 0x79, 0x0a, 0x0e, 0x43, 0x6f, 0x6d,
	0x70, 0x65, 0x6e, 0x73, 0x61, 0x74, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x19, 0x0a, 0x15, 0x44,
	0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x43, 0x6f, 0x6d, 0x70, 0x65, 0x6e, 0x73, 0x61, 0x74, 0x65,
	0x54, 0x79, 0x70, 0x65, 0x10, 0x00, 0x12, 0x19, 0x0a, 0x15, 0x43, 0x6f, 0x6d, 0x70, 0x65, 0x6e,
	0x73, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x6c, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x10,
	0x0a, 0x12, 0x15, 0x0a, 0x11, 0x43, 0x6f, 0x6d, 0x70, 0x65, 0x6e, 0x73, 0x61, 0x74, 0x65, 0x57,
	0x61, 0x6c, 0x66, 0x61, 0x72, 0x65, 0x10, 0x14, 0x12, 0x1a, 0x0a, 0x16, 0x43, 0x6f, 0x6d, 0x70,
	0x65, 0x6e, 0x73, 0x61, 0x74, 0x65, 0x53, 0x74, 0x61, 0x72, 0x74, 0x65, 0x72, 0x44, 0x65, 0x6c,
	0x61, 0x79, 0x10, 0x1e, 0x2a, 0x60, 0x0a, 0x0c, 0x42, 0x65, 0x6e, 0x65, 0x66, 0x69, 0x74, 0x53,
	0x74, 0x61, 0x74, 0x65, 0x12, 0x17, 0x0a, 0x13, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x42,
	0x65, 0x6e, 0x65, 0x66, 0x69, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x10, 0x00, 0x12, 0x0f, 0x0a,
	0x0b, 0x42, 0x65, 0x6e, 0x65, 0x66, 0x69, 0x74, 0x57, 0x61, 0x69, 0x74, 0x10, 0x0a, 0x12, 0x15,
	0x0a, 0x11, 0x42, 0x65, 0x6e, 0x65, 0x66, 0x69, 0x74, 0x43, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61,
	0x74, 0x65, 0x64, 0x10, 0x14, 0x12, 0x0f, 0x0a, 0x0b, 0x42, 0x65, 0x6e, 0x65, 0x66, 0x69, 0x74,
	0x44, 0x6f, 0x6e, 0x65, 0x10, 0x1e, 0x2a, 0x57, 0x0a, 0x0e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53,
	0x74, 0x61, 0x72, 0x74, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x19, 0x0a, 0x15, 0x44, 0x65, 0x66, 0x61,
	0x75, 0x6c, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x74, 0x61, 0x72, 0x74, 0x4d, 0x6f, 0x64,
	0x65, 0x10, 0x00, 0x12, 0x11, 0x0a, 0x0d, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x74, 0x61, 0x72,
	0x74, 0x54, 0x42, 0x44, 0x10, 0x0a, 0x12, 0x17, 0x0a, 0x13, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53,
	0x74, 0x61, 0x72, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x65, 0x64, 0x10, 0x14, 0x42,
	0x3b, 0x5a, 0x39, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4e, 0x70,
	0x6f, 0x6f, 0x6c, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x2f, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
>>>>>>> a68c66b1588db8e3a56afcd919ff25a3008745d6
}

var (
	file_npool_basetypes_order_v1_enums_proto_rawDescOnce sync.Once
	file_npool_basetypes_order_v1_enums_proto_rawDescData = file_npool_basetypes_order_v1_enums_proto_rawDesc
)

func file_npool_basetypes_order_v1_enums_proto_rawDescGZIP() []byte {
	file_npool_basetypes_order_v1_enums_proto_rawDescOnce.Do(func() {
		file_npool_basetypes_order_v1_enums_proto_rawDescData = protoimpl.X.CompressGZIP(file_npool_basetypes_order_v1_enums_proto_rawDescData)
	})
	return file_npool_basetypes_order_v1_enums_proto_rawDescData
}

var file_npool_basetypes_order_v1_enums_proto_enumTypes = make([]protoimpl.EnumInfo, 8)
var file_npool_basetypes_order_v1_enums_proto_goTypes = []interface{}{
	(OrderType)(0),      // 0: basetypes.order.v1.OrderType
	(OrderState)(0),     // 1: basetypes.order.v1.OrderState
	(PaymentState)(0),   // 2: basetypes.order.v1.PaymentState
	(InvestmentType)(0), // 3: basetypes.order.v1.InvestmentType
	(PaymentType)(0),    // 4: basetypes.order.v1.PaymentType
	(CompensateType)(0), // 5: basetypes.order.v1.CompensateType
	(BenefitState)(0),   // 6: basetypes.order.v1.BenefitState
	(OrderStartMode)(0), // 7: basetypes.order.v1.OrderStartMode
}
var file_npool_basetypes_order_v1_enums_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_npool_basetypes_order_v1_enums_proto_init() }
func file_npool_basetypes_order_v1_enums_proto_init() {
	if File_npool_basetypes_order_v1_enums_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_npool_basetypes_order_v1_enums_proto_rawDesc,
			NumEnums:      8,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_npool_basetypes_order_v1_enums_proto_goTypes,
		DependencyIndexes: file_npool_basetypes_order_v1_enums_proto_depIdxs,
		EnumInfos:         file_npool_basetypes_order_v1_enums_proto_enumTypes,
	}.Build()
	File_npool_basetypes_order_v1_enums_proto = out.File
	file_npool_basetypes_order_v1_enums_proto_rawDesc = nil
	file_npool_basetypes_order_v1_enums_proto_goTypes = nil
	file_npool_basetypes_order_v1_enums_proto_depIdxs = nil
}
