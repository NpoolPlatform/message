// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.18.1
// source: npool/inspire/mw/v1/achievement/statement/order/payment/statement.proto

package payment

import (
	v1 "github.com/NpoolPlatform/message/npool/basetypes/inspire/v1"
	v11 "github.com/NpoolPlatform/message/npool/basetypes/v1"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type StatementReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID                *uint32 `protobuf:"varint,10,opt,name=ID,proto3,oneof" json:"ID,omitempty"`
	EntID             *string `protobuf:"bytes,20,opt,name=EntID,proto3,oneof" json:"EntID,omitempty"`
	StatementID       *string `protobuf:"bytes,30,opt,name=StatementID,proto3,oneof" json:"StatementID,omitempty"`
	PaymentCoinTypeID *string `protobuf:"bytes,40,opt,name=PaymentCoinTypeID,proto3,oneof" json:"PaymentCoinTypeID,omitempty"`
	CoinUSDCurrency   *string `protobuf:"bytes,50,opt,name=CoinUSDCurrency,proto3,oneof" json:"CoinUSDCurrency,omitempty"`
	Amount            *string `protobuf:"bytes,60,opt,name=Amount,proto3,oneof" json:"Amount,omitempty"`
	CommissionAmount  *string `protobuf:"bytes,70,opt,name=CommissionAmount,proto3,oneof" json:"CommissionAmount,omitempty"`
}

func (x *StatementReq) Reset() {
	*x = StatementReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_inspire_mw_v1_achievement_statement_order_payment_statement_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatementReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatementReq) ProtoMessage() {}

func (x *StatementReq) ProtoReflect() protoreflect.Message {
	mi := &file_npool_inspire_mw_v1_achievement_statement_order_payment_statement_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatementReq.ProtoReflect.Descriptor instead.
func (*StatementReq) Descriptor() ([]byte, []int) {
	return file_npool_inspire_mw_v1_achievement_statement_order_payment_statement_proto_rawDescGZIP(), []int{0}
}

func (x *StatementReq) GetID() uint32 {
	if x != nil && x.ID != nil {
		return *x.ID
	}
	return 0
}

func (x *StatementReq) GetEntID() string {
	if x != nil && x.EntID != nil {
		return *x.EntID
	}
	return ""
}

func (x *StatementReq) GetStatementID() string {
	if x != nil && x.StatementID != nil {
		return *x.StatementID
	}
	return ""
}

func (x *StatementReq) GetPaymentCoinTypeID() string {
	if x != nil && x.PaymentCoinTypeID != nil {
		return *x.PaymentCoinTypeID
	}
	return ""
}

func (x *StatementReq) GetCoinUSDCurrency() string {
	if x != nil && x.CoinUSDCurrency != nil {
		return *x.CoinUSDCurrency
	}
	return ""
}

func (x *StatementReq) GetAmount() string {
	if x != nil && x.Amount != nil {
		return *x.Amount
	}
	return ""
}

func (x *StatementReq) GetCommissionAmount() string {
	if x != nil && x.CommissionAmount != nil {
		return *x.CommissionAmount
	}
	return ""
}

type Statement struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: sql:"id"
	ID uint32 `protobuf:"varint,10,opt,name=ID,proto3" json:"ID,omitempty" sql:"id"`
	// @inject_tag: sql:"ent_id"
	EntID string `protobuf:"bytes,20,opt,name=EntID,proto3" json:"EntID,omitempty" sql:"ent_id"`
	// @inject_tag: sql:"statement_id"
	StatementID string `protobuf:"bytes,30,opt,name=StatementID,proto3" json:"StatementID,omitempty" sql:"statement_id"`
	// @inject_tag: sql:"payment_coin_type_id"
	PaymentCoinTypeID string `protobuf:"bytes,40,opt,name=PaymentCoinTypeID,proto3" json:"PaymentCoinTypeID,omitempty" sql:"payment_coin_type_id"`
	// @inject_tag: sql:"amount"
	Amount string `protobuf:"bytes,50,opt,name=Amount,proto3" json:"Amount,omitempty" sql:"amount"`
	// @inject_tag: sql:"commission_amount"
	CommissionAmount string `protobuf:"bytes,60,opt,name=CommissionAmount,proto3" json:"CommissionAmount,omitempty" sql:"commission_amount"`
	// @inject_tag: sql:"app_id"
	AppID string `protobuf:"bytes,70,opt,name=AppID,proto3" json:"AppID,omitempty" sql:"app_id"`
	// @inject_tag: sql:"user_id"
	UserID string `protobuf:"bytes,80,opt,name=UserID,proto3" json:"UserID,omitempty" sql:"user_id"`
	// @inject_tag: sql:"good_id"
	GoodID string `protobuf:"bytes,90,opt,name=GoodID,proto3" json:"GoodID,omitempty" sql:"good_id"`
	// @inject_tag: sql:"app_good_id"
	AppGoodID string `protobuf:"bytes,100,opt,name=AppGoodID,proto3" json:"AppGoodID,omitempty" sql:"app_good_id"`
	// @inject_tag: sql:"order_id"
	OrderID string `protobuf:"bytes,110,opt,name=OrderID,proto3" json:"OrderID,omitempty" sql:"order_id"`
	// @inject_tag: sql:"order_user_id"
	OrderUserID string `protobuf:"bytes,120,opt,name=OrderUserID,proto3" json:"OrderUserID,omitempty" sql:"order_user_id"`
	// @inject_tag: sql:"good_coin_type_id"
	GoodCoinTypeID string `protobuf:"bytes,130,opt,name=GoodCoinTypeID,proto3" json:"GoodCoinTypeID,omitempty" sql:"good_coin_type_id"`
	// @inject_tag: sql:"units"
	Units string `protobuf:"bytes,140,opt,name=Units,proto3" json:"Units,omitempty" sql:"units"`
	// @inject_tag: sql:"good_value_usd"
	GoodValueUSD string `protobuf:"bytes,150,opt,name=GoodValueUSD,proto3" json:"GoodValueUSD,omitempty" sql:"good_value_usd"`
	// @inject_tag: sql:"payment_amount_usd"
	PaymentAmountUSD string `protobuf:"bytes,160,opt,name=PaymentAmountUSD,proto3" json:"PaymentAmountUSD,omitempty" sql:"payment_amount_usd"`
	// @inject_tag: sql:"commission_amount_usd"
	CommissionAmountUSD string `protobuf:"bytes,170,opt,name=CommissionAmountUSD,proto3" json:"CommissionAmountUSD,omitempty" sql:"commission_amount_usd"`
	// @inject_tag: sql:"app_config_id"
	AppConfigID string `protobuf:"bytes,180,opt,name=AppConfigID,proto3" json:"AppConfigID,omitempty" sql:"app_config_id"`
	// @inject_tag: sql:"commission_config_id"
	CommissionConfigID string `protobuf:"bytes,190,opt,name=CommissionConfigID,proto3" json:"CommissionConfigID,omitempty" sql:"commission_config_id"`
	// @inject_tag: sql:"commission_config_type"
	CommissionConfigTypeStr string                  `protobuf:"bytes,200,opt,name=CommissionConfigTypeStr,proto3" json:"CommissionConfigTypeStr,omitempty" sql:"commission_config_type"`
	CommissionConfigType    v1.CommissionConfigType `protobuf:"varint,210,opt,name=CommissionConfigType,proto3,enum=basetypes.inspire.v1.CommissionConfigType" json:"CommissionConfigType,omitempty"`
	// @inject_tag: sql:"created_at"
	CreatedAt uint32 `protobuf:"varint,1000,opt,name=CreatedAt,proto3" json:"CreatedAt,omitempty" sql:"created_at"`
	// @inject_tag: sql:"updated_at"
	UpdatedAt uint32 `protobuf:"varint,1010,opt,name=UpdatedAt,proto3" json:"UpdatedAt,omitempty" sql:"updated_at"`
}

func (x *Statement) Reset() {
	*x = Statement{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_inspire_mw_v1_achievement_statement_order_payment_statement_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Statement) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Statement) ProtoMessage() {}

func (x *Statement) ProtoReflect() protoreflect.Message {
	mi := &file_npool_inspire_mw_v1_achievement_statement_order_payment_statement_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Statement.ProtoReflect.Descriptor instead.
func (*Statement) Descriptor() ([]byte, []int) {
	return file_npool_inspire_mw_v1_achievement_statement_order_payment_statement_proto_rawDescGZIP(), []int{1}
}

func (x *Statement) GetID() uint32 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *Statement) GetEntID() string {
	if x != nil {
		return x.EntID
	}
	return ""
}

func (x *Statement) GetStatementID() string {
	if x != nil {
		return x.StatementID
	}
	return ""
}

func (x *Statement) GetPaymentCoinTypeID() string {
	if x != nil {
		return x.PaymentCoinTypeID
	}
	return ""
}

func (x *Statement) GetAmount() string {
	if x != nil {
		return x.Amount
	}
	return ""
}

func (x *Statement) GetCommissionAmount() string {
	if x != nil {
		return x.CommissionAmount
	}
	return ""
}

func (x *Statement) GetAppID() string {
	if x != nil {
		return x.AppID
	}
	return ""
}

func (x *Statement) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

func (x *Statement) GetGoodID() string {
	if x != nil {
		return x.GoodID
	}
	return ""
}

func (x *Statement) GetAppGoodID() string {
	if x != nil {
		return x.AppGoodID
	}
	return ""
}

func (x *Statement) GetOrderID() string {
	if x != nil {
		return x.OrderID
	}
	return ""
}

func (x *Statement) GetOrderUserID() string {
	if x != nil {
		return x.OrderUserID
	}
	return ""
}

func (x *Statement) GetGoodCoinTypeID() string {
	if x != nil {
		return x.GoodCoinTypeID
	}
	return ""
}

func (x *Statement) GetUnits() string {
	if x != nil {
		return x.Units
	}
	return ""
}

func (x *Statement) GetGoodValueUSD() string {
	if x != nil {
		return x.GoodValueUSD
	}
	return ""
}

func (x *Statement) GetPaymentAmountUSD() string {
	if x != nil {
		return x.PaymentAmountUSD
	}
	return ""
}

func (x *Statement) GetCommissionAmountUSD() string {
	if x != nil {
		return x.CommissionAmountUSD
	}
	return ""
}

func (x *Statement) GetAppConfigID() string {
	if x != nil {
		return x.AppConfigID
	}
	return ""
}

func (x *Statement) GetCommissionConfigID() string {
	if x != nil {
		return x.CommissionConfigID
	}
	return ""
}

func (x *Statement) GetCommissionConfigTypeStr() string {
	if x != nil {
		return x.CommissionConfigTypeStr
	}
	return ""
}

func (x *Statement) GetCommissionConfigType() v1.CommissionConfigType {
	if x != nil {
		return x.CommissionConfigType
	}
	return v1.CommissionConfigType(0)
}

func (x *Statement) GetCreatedAt() uint32 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *Statement) GetUpdatedAt() uint32 {
	if x != nil {
		return x.UpdatedAt
	}
	return 0
}

type Conds struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EntID             *v11.StringVal `protobuf:"bytes,10,opt,name=EntID,proto3,oneof" json:"EntID,omitempty"`
	AppID             *v11.StringVal `protobuf:"bytes,20,opt,name=AppID,proto3,oneof" json:"AppID,omitempty"`
	UserID            *v11.StringVal `protobuf:"bytes,30,opt,name=UserID,proto3,oneof" json:"UserID,omitempty"`
	PaymentCoinTypeID *v11.StringVal `protobuf:"bytes,40,opt,name=PaymentCoinTypeID,proto3,oneof" json:"PaymentCoinTypeID,omitempty"`
}

func (x *Conds) Reset() {
	*x = Conds{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_inspire_mw_v1_achievement_statement_order_payment_statement_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Conds) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Conds) ProtoMessage() {}

func (x *Conds) ProtoReflect() protoreflect.Message {
	mi := &file_npool_inspire_mw_v1_achievement_statement_order_payment_statement_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Conds.ProtoReflect.Descriptor instead.
func (*Conds) Descriptor() ([]byte, []int) {
	return file_npool_inspire_mw_v1_achievement_statement_order_payment_statement_proto_rawDescGZIP(), []int{2}
}

func (x *Conds) GetEntID() *v11.StringVal {
	if x != nil {
		return x.EntID
	}
	return nil
}

func (x *Conds) GetAppID() *v11.StringVal {
	if x != nil {
		return x.AppID
	}
	return nil
}

func (x *Conds) GetUserID() *v11.StringVal {
	if x != nil {
		return x.UserID
	}
	return nil
}

func (x *Conds) GetPaymentCoinTypeID() *v11.StringVal {
	if x != nil {
		return x.PaymentCoinTypeID
	}
	return nil
}

type GetStatementsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Conds  *Conds `protobuf:"bytes,10,opt,name=Conds,proto3" json:"Conds,omitempty"`
	Offset int32  `protobuf:"varint,20,opt,name=Offset,proto3" json:"Offset,omitempty"`
	Limit  int32  `protobuf:"varint,30,opt,name=Limit,proto3" json:"Limit,omitempty"`
}

func (x *GetStatementsRequest) Reset() {
	*x = GetStatementsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_inspire_mw_v1_achievement_statement_order_payment_statement_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetStatementsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetStatementsRequest) ProtoMessage() {}

func (x *GetStatementsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_inspire_mw_v1_achievement_statement_order_payment_statement_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetStatementsRequest.ProtoReflect.Descriptor instead.
func (*GetStatementsRequest) Descriptor() ([]byte, []int) {
	return file_npool_inspire_mw_v1_achievement_statement_order_payment_statement_proto_rawDescGZIP(), []int{3}
}

func (x *GetStatementsRequest) GetConds() *Conds {
	if x != nil {
		return x.Conds
	}
	return nil
}

func (x *GetStatementsRequest) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *GetStatementsRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type GetStatementsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Infos []*Statement `protobuf:"bytes,10,rep,name=Infos,proto3" json:"Infos,omitempty"`
	Total uint32       `protobuf:"varint,20,opt,name=Total,proto3" json:"Total,omitempty"`
}

func (x *GetStatementsResponse) Reset() {
	*x = GetStatementsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_inspire_mw_v1_achievement_statement_order_payment_statement_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetStatementsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetStatementsResponse) ProtoMessage() {}

func (x *GetStatementsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_inspire_mw_v1_achievement_statement_order_payment_statement_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetStatementsResponse.ProtoReflect.Descriptor instead.
func (*GetStatementsResponse) Descriptor() ([]byte, []int) {
	return file_npool_inspire_mw_v1_achievement_statement_order_payment_statement_proto_rawDescGZIP(), []int{4}
}

func (x *GetStatementsResponse) GetInfos() []*Statement {
	if x != nil {
		return x.Infos
	}
	return nil
}

func (x *GetStatementsResponse) GetTotal() uint32 {
	if x != nil {
		return x.Total
	}
	return 0
}

var File_npool_inspire_mw_v1_achievement_statement_order_payment_statement_proto protoreflect.FileDescriptor

var file_npool_inspire_mw_v1_achievement_statement_order_payment_statement_proto_rawDesc = []byte{
	0x0a, 0x47, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x69, 0x6e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2f,
	0x6d, 0x77, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x63, 0x68, 0x69, 0x65, 0x76, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x2f, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x6d,
	0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x39, 0x69, 0x6e, 0x73, 0x70, 0x69,
	0x72, 0x65, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x61, 0x63,
	0x68, 0x69, 0x65, 0x76, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x65, 0x6d,
	0x65, 0x6e, 0x74, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e,
	0x74, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1e, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x26, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x2f, 0x69, 0x6e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x65,
	0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x80, 0x03, 0x0a, 0x0c, 0x53,
	0x74, 0x61, 0x74, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x12, 0x13, 0x0a, 0x02, 0x49,
	0x44, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0d, 0x48, 0x00, 0x52, 0x02, 0x49, 0x44, 0x88, 0x01, 0x01,
	0x12, 0x19, 0x0a, 0x05, 0x45, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x48,
	0x01, 0x52, 0x05, 0x45, 0x6e, 0x74, 0x49, 0x44, 0x88, 0x01, 0x01, 0x12, 0x25, 0x0a, 0x0b, 0x53,
	0x74, 0x61, 0x74, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x09,
	0x48, 0x02, 0x52, 0x0b, 0x53, 0x74, 0x61, 0x74, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x44, 0x88,
	0x01, 0x01, 0x12, 0x31, 0x0a, 0x11, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x69,
	0x6e, 0x54, 0x79, 0x70, 0x65, 0x49, 0x44, 0x18, 0x28, 0x20, 0x01, 0x28, 0x09, 0x48, 0x03, 0x52,
	0x11, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x69, 0x6e, 0x54, 0x79, 0x70, 0x65,
	0x49, 0x44, 0x88, 0x01, 0x01, 0x12, 0x2d, 0x0a, 0x0f, 0x43, 0x6f, 0x69, 0x6e, 0x55, 0x53, 0x44,
	0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x32, 0x20, 0x01, 0x28, 0x09, 0x48, 0x04,
	0x52, 0x0f, 0x43, 0x6f, 0x69, 0x6e, 0x55, 0x53, 0x44, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63,
	0x79, 0x88, 0x01, 0x01, 0x12, 0x1b, 0x0a, 0x06, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x3c,
	0x20, 0x01, 0x28, 0x09, 0x48, 0x05, 0x52, 0x06, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x88, 0x01,
	0x01, 0x12, 0x2f, 0x0a, 0x10, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x41,
	0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x46, 0x20, 0x01, 0x28, 0x09, 0x48, 0x06, 0x52, 0x10, 0x43,
	0x6f, 0x6d, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x88,
	0x01, 0x01, 0x42, 0x05, 0x0a, 0x03, 0x5f, 0x49, 0x44, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x45, 0x6e,
	0x74, 0x49, 0x44, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x53, 0x74, 0x61, 0x74, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x49, 0x44, 0x42, 0x14, 0x0a, 0x12, 0x5f, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x43,
	0x6f, 0x69, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x49, 0x44, 0x42, 0x12, 0x0a, 0x10, 0x5f, 0x43, 0x6f,
	0x69, 0x6e, 0x55, 0x53, 0x44, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x42, 0x09, 0x0a,
	0x07, 0x5f, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x13, 0x0a, 0x11, 0x5f, 0x43, 0x6f, 0x6d,
	0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0xd8, 0x06,
	0x0a, 0x09, 0x53, 0x74, 0x61, 0x74, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49,
	0x44, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x45,
	0x6e, 0x74, 0x49, 0x44, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x45, 0x6e, 0x74, 0x49,
	0x44, 0x12, 0x20, 0x0a, 0x0b, 0x53, 0x74, 0x61, 0x74, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x44,
	0x18, 0x1e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x53, 0x74, 0x61, 0x74, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x49, 0x44, 0x12, 0x2c, 0x0a, 0x11, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x43, 0x6f,
	0x69, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x49, 0x44, 0x18, 0x28, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11,
	0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x69, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x49,
	0x44, 0x12, 0x16, 0x0a, 0x06, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x32, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x2a, 0x0a, 0x10, 0x43, 0x6f, 0x6d,
	0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x3c, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x10, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x41,
	0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x41, 0x70, 0x70, 0x49, 0x44, 0x18, 0x46,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x41, 0x70, 0x70, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x55,
	0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x50, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x55, 0x73, 0x65,
	0x72, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x47, 0x6f, 0x6f, 0x64, 0x49, 0x44, 0x18, 0x5a, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x47, 0x6f, 0x6f, 0x64, 0x49, 0x44, 0x12, 0x1c, 0x0a, 0x09, 0x41,
	0x70, 0x70, 0x47, 0x6f, 0x6f, 0x64, 0x49, 0x44, 0x18, 0x64, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x41, 0x70, 0x70, 0x47, 0x6f, 0x6f, 0x64, 0x49, 0x44, 0x12, 0x18, 0x0a, 0x07, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x49, 0x44, 0x18, 0x6e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x49, 0x44, 0x12, 0x20, 0x0a, 0x0b, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72,
	0x49, 0x44, 0x18, 0x78, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x55,
	0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x27, 0x0a, 0x0e, 0x47, 0x6f, 0x6f, 0x64, 0x43, 0x6f, 0x69,
	0x6e, 0x54, 0x79, 0x70, 0x65, 0x49, 0x44, 0x18, 0x82, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e,
	0x47, 0x6f, 0x6f, 0x64, 0x43, 0x6f, 0x69, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x49, 0x44, 0x12, 0x15,
	0x0a, 0x05, 0x55, 0x6e, 0x69, 0x74, 0x73, 0x18, 0x8c, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x55, 0x6e, 0x69, 0x74, 0x73, 0x12, 0x23, 0x0a, 0x0c, 0x47, 0x6f, 0x6f, 0x64, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x55, 0x53, 0x44, 0x18, 0x96, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x47, 0x6f,
	0x6f, 0x64, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x55, 0x53, 0x44, 0x12, 0x2b, 0x0a, 0x10, 0x50, 0x61,
	0x79, 0x6d, 0x65, 0x6e, 0x74, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x55, 0x53, 0x44, 0x18, 0xa0,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x41, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x55, 0x53, 0x44, 0x12, 0x31, 0x0a, 0x13, 0x43, 0x6f, 0x6d, 0x6d, 0x69,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x55, 0x53, 0x44, 0x18, 0xaa,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x13, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x55, 0x53, 0x44, 0x12, 0x21, 0x0a, 0x0b, 0x41, 0x70,
	0x70, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x49, 0x44, 0x18, 0xb4, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x41, 0x70, 0x70, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x49, 0x44, 0x12, 0x2f, 0x0a,
	0x12, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x49, 0x44, 0x18, 0xbe, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x43, 0x6f, 0x6d, 0x6d,
	0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x49, 0x44, 0x12, 0x39,
	0x0a, 0x17, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x54, 0x79, 0x70, 0x65, 0x53, 0x74, 0x72, 0x18, 0xc8, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x17, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x54, 0x79, 0x70, 0x65, 0x53, 0x74, 0x72, 0x12, 0x5f, 0x0a, 0x14, 0x43, 0x6f, 0x6d,
	0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x54, 0x79, 0x70,
	0x65, 0x18, 0xd2, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2a, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x2e, 0x69, 0x6e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x54, 0x79, 0x70, 0x65, 0x52, 0x14, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1d, 0x0a, 0x09, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0xe8, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x09, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0xf2, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0xa6, 0x02, 0x0a, 0x05, 0x43, 0x6f, 0x6e,
	0x64, 0x73, 0x12, 0x32, 0x0a, 0x05, 0x45, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x17, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x48, 0x00, 0x52, 0x05, 0x45, 0x6e,
	0x74, 0x49, 0x44, 0x88, 0x01, 0x01, 0x12, 0x32, 0x0a, 0x05, 0x41, 0x70, 0x70, 0x49, 0x44, 0x18,
	0x14, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x48, 0x01,
	0x52, 0x05, 0x41, 0x70, 0x70, 0x49, 0x44, 0x88, 0x01, 0x01, 0x12, 0x34, 0x0a, 0x06, 0x55, 0x73,
	0x65, 0x72, 0x49, 0x44, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x62, 0x61, 0x73,
	0x65, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67,
	0x56, 0x61, 0x6c, 0x48, 0x02, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x88, 0x01, 0x01,
	0x12, 0x4a, 0x0a, 0x11, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x69, 0x6e, 0x54,
	0x79, 0x70, 0x65, 0x49, 0x44, 0x18, 0x28, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x62, 0x61,
	0x73, 0x65, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e,
	0x67, 0x56, 0x61, 0x6c, 0x48, 0x03, 0x52, 0x11, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x43,
	0x6f, 0x69, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x49, 0x44, 0x88, 0x01, 0x01, 0x42, 0x08, 0x0a, 0x06,
	0x5f, 0x45, 0x6e, 0x74, 0x49, 0x44, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x41, 0x70, 0x70, 0x49, 0x44,
	0x42, 0x09, 0x0a, 0x07, 0x5f, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x42, 0x14, 0x0a, 0x12, 0x5f,
	0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x69, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x49,
	0x44, 0x22, 0x9c, 0x01, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x6d, 0x65,
	0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x56, 0x0a, 0x05, 0x43, 0x6f,
	0x6e, 0x64, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x40, 0x2e, 0x69, 0x6e, 0x73, 0x70,
	0x69, 0x72, 0x65, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x61,
	0x63, 0x68, 0x69, 0x65, 0x76, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x65,
	0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x64, 0x73, 0x52, 0x05, 0x43, 0x6f, 0x6e,
	0x64, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x14, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x06, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x4c, 0x69,
	0x6d, 0x69, 0x74, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x4c, 0x69, 0x6d, 0x69, 0x74,
	0x22, 0x89, 0x01, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5a, 0x0a, 0x05, 0x49, 0x6e,
	0x66, 0x6f, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x44, 0x2e, 0x69, 0x6e, 0x73, 0x70,
	0x69, 0x72, 0x65, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x61,
	0x63, 0x68, 0x69, 0x65, 0x76, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x65,
	0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x52,
	0x05, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x18,
	0x14, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x32, 0xf7, 0x01, 0x0a,
	0x0a, 0x4d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x12, 0xe8, 0x01, 0x0a, 0x0d,
	0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x4f, 0x2e,
	0x69, 0x6e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61,
	0x72, 0x65, 0x2e, 0x61, 0x63, 0x68, 0x69, 0x65, 0x76, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x73,
	0x74, 0x61, 0x74, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x70,
	0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x74, 0x61,
	0x74, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x50,
	0x2e, 0x69, 0x6e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77,
	0x61, 0x72, 0x65, 0x2e, 0x61, 0x63, 0x68, 0x69, 0x65, 0x76, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e,
	0x73, 0x74, 0x61, 0x74, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e,
	0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x74,
	0x61, 0x74, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x34, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2e, 0x3a, 0x01, 0x2a, 0x22, 0x29, 0x2f, 0x76, 0x31,
	0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x73,
	0x74, 0x61, 0x74, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74,
	0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x42, 0x5a, 0x5a, 0x58, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4e, 0x70, 0x6f, 0x6f, 0x6c, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f,
	0x72, 0x6d, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x6e, 0x70, 0x6f, 0x6f, 0x6c,
	0x2f, 0x69, 0x6e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2f, 0x6d, 0x77, 0x2f, 0x76, 0x31, 0x2f, 0x61,
	0x63, 0x68, 0x69, 0x65, 0x76, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x65,
	0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2f, 0x70, 0x61, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_npool_inspire_mw_v1_achievement_statement_order_payment_statement_proto_rawDescOnce sync.Once
	file_npool_inspire_mw_v1_achievement_statement_order_payment_statement_proto_rawDescData = file_npool_inspire_mw_v1_achievement_statement_order_payment_statement_proto_rawDesc
)

func file_npool_inspire_mw_v1_achievement_statement_order_payment_statement_proto_rawDescGZIP() []byte {
	file_npool_inspire_mw_v1_achievement_statement_order_payment_statement_proto_rawDescOnce.Do(func() {
		file_npool_inspire_mw_v1_achievement_statement_order_payment_statement_proto_rawDescData = protoimpl.X.CompressGZIP(file_npool_inspire_mw_v1_achievement_statement_order_payment_statement_proto_rawDescData)
	})
	return file_npool_inspire_mw_v1_achievement_statement_order_payment_statement_proto_rawDescData
}

var file_npool_inspire_mw_v1_achievement_statement_order_payment_statement_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_npool_inspire_mw_v1_achievement_statement_order_payment_statement_proto_goTypes = []interface{}{
	(*StatementReq)(nil),          // 0: inspire.middleware.achievement.statement.order.payment.v1.StatementReq
	(*Statement)(nil),             // 1: inspire.middleware.achievement.statement.order.payment.v1.Statement
	(*Conds)(nil),                 // 2: inspire.middleware.achievement.statement.order.payment.v1.Conds
	(*GetStatementsRequest)(nil),  // 3: inspire.middleware.achievement.statement.order.payment.v1.GetStatementsRequest
	(*GetStatementsResponse)(nil), // 4: inspire.middleware.achievement.statement.order.payment.v1.GetStatementsResponse
	(v1.CommissionConfigType)(0),  // 5: basetypes.inspire.v1.CommissionConfigType
	(*v11.StringVal)(nil),         // 6: basetypes.v1.StringVal
}
var file_npool_inspire_mw_v1_achievement_statement_order_payment_statement_proto_depIdxs = []int32{
	5, // 0: inspire.middleware.achievement.statement.order.payment.v1.Statement.CommissionConfigType:type_name -> basetypes.inspire.v1.CommissionConfigType
	6, // 1: inspire.middleware.achievement.statement.order.payment.v1.Conds.EntID:type_name -> basetypes.v1.StringVal
	6, // 2: inspire.middleware.achievement.statement.order.payment.v1.Conds.AppID:type_name -> basetypes.v1.StringVal
	6, // 3: inspire.middleware.achievement.statement.order.payment.v1.Conds.UserID:type_name -> basetypes.v1.StringVal
	6, // 4: inspire.middleware.achievement.statement.order.payment.v1.Conds.PaymentCoinTypeID:type_name -> basetypes.v1.StringVal
	2, // 5: inspire.middleware.achievement.statement.order.payment.v1.GetStatementsRequest.Conds:type_name -> inspire.middleware.achievement.statement.order.payment.v1.Conds
	1, // 6: inspire.middleware.achievement.statement.order.payment.v1.GetStatementsResponse.Infos:type_name -> inspire.middleware.achievement.statement.order.payment.v1.Statement
	3, // 7: inspire.middleware.achievement.statement.order.payment.v1.Middleware.GetStatements:input_type -> inspire.middleware.achievement.statement.order.payment.v1.GetStatementsRequest
	4, // 8: inspire.middleware.achievement.statement.order.payment.v1.Middleware.GetStatements:output_type -> inspire.middleware.achievement.statement.order.payment.v1.GetStatementsResponse
	8, // [8:9] is the sub-list for method output_type
	7, // [7:8] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_npool_inspire_mw_v1_achievement_statement_order_payment_statement_proto_init() }
func file_npool_inspire_mw_v1_achievement_statement_order_payment_statement_proto_init() {
	if File_npool_inspire_mw_v1_achievement_statement_order_payment_statement_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_npool_inspire_mw_v1_achievement_statement_order_payment_statement_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StatementReq); i {
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
		file_npool_inspire_mw_v1_achievement_statement_order_payment_statement_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Statement); i {
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
		file_npool_inspire_mw_v1_achievement_statement_order_payment_statement_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Conds); i {
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
		file_npool_inspire_mw_v1_achievement_statement_order_payment_statement_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetStatementsRequest); i {
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
		file_npool_inspire_mw_v1_achievement_statement_order_payment_statement_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetStatementsResponse); i {
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
	file_npool_inspire_mw_v1_achievement_statement_order_payment_statement_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_npool_inspire_mw_v1_achievement_statement_order_payment_statement_proto_msgTypes[2].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_npool_inspire_mw_v1_achievement_statement_order_payment_statement_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_npool_inspire_mw_v1_achievement_statement_order_payment_statement_proto_goTypes,
		DependencyIndexes: file_npool_inspire_mw_v1_achievement_statement_order_payment_statement_proto_depIdxs,
		MessageInfos:      file_npool_inspire_mw_v1_achievement_statement_order_payment_statement_proto_msgTypes,
	}.Build()
	File_npool_inspire_mw_v1_achievement_statement_order_payment_statement_proto = out.File
	file_npool_inspire_mw_v1_achievement_statement_order_payment_statement_proto_rawDesc = nil
	file_npool_inspire_mw_v1_achievement_statement_order_payment_statement_proto_goTypes = nil
	file_npool_inspire_mw_v1_achievement_statement_order_payment_statement_proto_depIdxs = nil
}
