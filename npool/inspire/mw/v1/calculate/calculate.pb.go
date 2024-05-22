// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.18.1
// source: npool/inspire/mw/v1/calculate/calculate.proto

package calculate

import (
	v1 "github.com/NpoolPlatform/message/npool/basetypes/inspire/v1"
	order "github.com/NpoolPlatform/message/npool/inspire/mw/v1/achievement/statement/order"
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

type Payment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CoinTypeID      string `protobuf:"bytes,10,opt,name=CoinTypeID,proto3" json:"CoinTypeID,omitempty"`
	CoinUSDCurrency string `protobuf:"bytes,20,opt,name=CoinUSDCurrency,proto3" json:"CoinUSDCurrency,omitempty"`
	Amount          string `protobuf:"bytes,30,opt,name=Amount,proto3" json:"Amount,omitempty"`
}

func (x *Payment) Reset() {
	*x = Payment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_inspire_mw_v1_calculate_calculate_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Payment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Payment) ProtoMessage() {}

func (x *Payment) ProtoReflect() protoreflect.Message {
	mi := &file_npool_inspire_mw_v1_calculate_calculate_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Payment.ProtoReflect.Descriptor instead.
func (*Payment) Descriptor() ([]byte, []int) {
	return file_npool_inspire_mw_v1_calculate_calculate_proto_rawDescGZIP(), []int{0}
}

func (x *Payment) GetCoinTypeID() string {
	if x != nil {
		return x.CoinTypeID
	}
	return ""
}

func (x *Payment) GetCoinUSDCurrency() string {
	if x != nil {
		return x.CoinUSDCurrency
	}
	return ""
}

func (x *Payment) GetAmount() string {
	if x != nil {
		return x.Amount
	}
	return ""
}

type CalculateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AppID            string        `protobuf:"bytes,10,opt,name=AppID,proto3" json:"AppID,omitempty"`
	UserID           string        `protobuf:"bytes,20,opt,name=UserID,proto3" json:"UserID,omitempty"`
	GoodID           string        `protobuf:"bytes,30,opt,name=GoodID,proto3" json:"GoodID,omitempty"`
	AppGoodID        string        `protobuf:"bytes,40,opt,name=AppGoodID,proto3" json:"AppGoodID,omitempty"`
	OrderID          string        `protobuf:"bytes,50,opt,name=OrderID,proto3" json:"OrderID,omitempty"`
	GoodCoinTypeID   string        `protobuf:"bytes,60,opt,name=GoodCoinTypeID,proto3" json:"GoodCoinTypeID,omitempty"`
	OrderUserID      string        `protobuf:"bytes,70,opt,name=OrderUserID,proto3" json:"OrderUserID,omitempty"`
	Units            string        `protobuf:"bytes,80,opt,name=Units,proto3" json:"Units,omitempty"`
	PaymentAmountUSD string        `protobuf:"bytes,90,opt,name=PaymentAmountUSD,proto3" json:"PaymentAmountUSD,omitempty"`
	GoodValueUSD     string        `protobuf:"bytes,100,opt,name=GoodValueUSD,proto3" json:"GoodValueUSD,omitempty"`
	SettleType       v1.SettleType `protobuf:"varint,110,opt,name=SettleType,proto3,enum=basetypes.inspire.v1.SettleType" json:"SettleType,omitempty"`
	HasCommission    bool          `protobuf:"varint,120,opt,name=HasCommission,proto3" json:"HasCommission,omitempty"`
	OrderCreatedAt   uint32        `protobuf:"varint,130,opt,name=OrderCreatedAt,proto3" json:"OrderCreatedAt,omitempty"`
	Payments         []*Payment    `protobuf:"bytes,140,rep,name=Payments,proto3" json:"Payments,omitempty"`
}

func (x *CalculateRequest) Reset() {
	*x = CalculateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_inspire_mw_v1_calculate_calculate_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CalculateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CalculateRequest) ProtoMessage() {}

func (x *CalculateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_inspire_mw_v1_calculate_calculate_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CalculateRequest.ProtoReflect.Descriptor instead.
func (*CalculateRequest) Descriptor() ([]byte, []int) {
	return file_npool_inspire_mw_v1_calculate_calculate_proto_rawDescGZIP(), []int{1}
}

func (x *CalculateRequest) GetAppID() string {
	if x != nil {
		return x.AppID
	}
	return ""
}

func (x *CalculateRequest) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

func (x *CalculateRequest) GetGoodID() string {
	if x != nil {
		return x.GoodID
	}
	return ""
}

func (x *CalculateRequest) GetAppGoodID() string {
	if x != nil {
		return x.AppGoodID
	}
	return ""
}

func (x *CalculateRequest) GetOrderID() string {
	if x != nil {
		return x.OrderID
	}
	return ""
}

func (x *CalculateRequest) GetGoodCoinTypeID() string {
	if x != nil {
		return x.GoodCoinTypeID
	}
	return ""
}

func (x *CalculateRequest) GetOrderUserID() string {
	if x != nil {
		return x.OrderUserID
	}
	return ""
}

func (x *CalculateRequest) GetUnits() string {
	if x != nil {
		return x.Units
	}
	return ""
}

func (x *CalculateRequest) GetPaymentAmountUSD() string {
	if x != nil {
		return x.PaymentAmountUSD
	}
	return ""
}

func (x *CalculateRequest) GetGoodValueUSD() string {
	if x != nil {
		return x.GoodValueUSD
	}
	return ""
}

func (x *CalculateRequest) GetSettleType() v1.SettleType {
	if x != nil {
		return x.SettleType
	}
	return v1.SettleType(0)
}

func (x *CalculateRequest) GetHasCommission() bool {
	if x != nil {
		return x.HasCommission
	}
	return false
}

func (x *CalculateRequest) GetOrderCreatedAt() uint32 {
	if x != nil {
		return x.OrderCreatedAt
	}
	return 0
}

func (x *CalculateRequest) GetPayments() []*Payment {
	if x != nil {
		return x.Payments
	}
	return nil
}

type CalculateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Infos []*order.StatementReq `protobuf:"bytes,10,rep,name=Infos,proto3" json:"Infos,omitempty"`
}

func (x *CalculateResponse) Reset() {
	*x = CalculateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_inspire_mw_v1_calculate_calculate_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CalculateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CalculateResponse) ProtoMessage() {}

func (x *CalculateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_inspire_mw_v1_calculate_calculate_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CalculateResponse.ProtoReflect.Descriptor instead.
func (*CalculateResponse) Descriptor() ([]byte, []int) {
	return file_npool_inspire_mw_v1_calculate_calculate_proto_rawDescGZIP(), []int{2}
}

func (x *CalculateResponse) GetInfos() []*order.StatementReq {
	if x != nil {
		return x.Infos
	}
	return nil
}

var File_npool_inspire_mw_v1_calculate_calculate_proto protoreflect.FileDescriptor

var file_npool_inspire_mw_v1_calculate_calculate_proto_rawDesc = []byte{
	0x0a, 0x2d, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x69, 0x6e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2f,
	0x6d, 0x77, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x2f,
	0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x1f, 0x69, 0x6e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77,
	0x61, 0x72, 0x65, 0x2e, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x2e, 0x76, 0x31,
	0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x26,
	0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f,
	0x69, 0x6e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x3f, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x69, 0x6e,
	0x73, 0x70, 0x69, 0x72, 0x65, 0x2f, 0x6d, 0x77, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x63, 0x68, 0x69,
	0x65, 0x76, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6b, 0x0a, 0x07, 0x50, 0x61, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x43, 0x6f, 0x69, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x49, 0x44,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x43, 0x6f, 0x69, 0x6e, 0x54, 0x79, 0x70, 0x65,
	0x49, 0x44, 0x12, 0x28, 0x0a, 0x0f, 0x43, 0x6f, 0x69, 0x6e, 0x55, 0x53, 0x44, 0x43, 0x75, 0x72,
	0x72, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x43, 0x6f, 0x69,
	0x6e, 0x55, 0x53, 0x44, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x16, 0x0a, 0x06,
	0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x41, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x22, 0x98, 0x04, 0x0a, 0x10, 0x43, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x41, 0x70, 0x70,
	0x49, 0x44, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x41, 0x70, 0x70, 0x49, 0x44, 0x12,
	0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x47, 0x6f, 0x6f, 0x64, 0x49,
	0x44, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x47, 0x6f, 0x6f, 0x64, 0x49, 0x44, 0x12,
	0x1c, 0x0a, 0x09, 0x41, 0x70, 0x70, 0x47, 0x6f, 0x6f, 0x64, 0x49, 0x44, 0x18, 0x28, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x41, 0x70, 0x70, 0x47, 0x6f, 0x6f, 0x64, 0x49, 0x44, 0x12, 0x18, 0x0a,
	0x07, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x44, 0x18, 0x32, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x44, 0x12, 0x26, 0x0a, 0x0e, 0x47, 0x6f, 0x6f, 0x64, 0x43,
	0x6f, 0x69, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x49, 0x44, 0x18, 0x3c, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0e, 0x47, 0x6f, 0x6f, 0x64, 0x43, 0x6f, 0x69, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x49, 0x44, 0x12,
	0x20, 0x0a, 0x0b, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x46,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x49,
	0x44, 0x12, 0x14, 0x0a, 0x05, 0x55, 0x6e, 0x69, 0x74, 0x73, 0x18, 0x50, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x55, 0x6e, 0x69, 0x74, 0x73, 0x12, 0x2a, 0x0a, 0x10, 0x50, 0x61, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x55, 0x53, 0x44, 0x18, 0x5a, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x10, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x55, 0x53, 0x44, 0x12, 0x22, 0x0a, 0x0c, 0x47, 0x6f, 0x6f, 0x64, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x55, 0x53, 0x44, 0x18, 0x64, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x47, 0x6f, 0x6f, 0x64, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x55, 0x53, 0x44, 0x12, 0x40, 0x0a, 0x0a, 0x53, 0x65, 0x74, 0x74, 0x6c,
	0x65, 0x54, 0x79, 0x70, 0x65, 0x18, 0x6e, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x20, 0x2e, 0x62, 0x61,
	0x73, 0x65, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x69, 0x6e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x53, 0x65, 0x74, 0x74, 0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0a, 0x53,
	0x65, 0x74, 0x74, 0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x48, 0x61, 0x73,
	0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x78, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x0d, 0x48, 0x61, 0x73, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12,
	0x27, 0x0a, 0x0e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x18, 0x82, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x45, 0x0a, 0x08, 0x50, 0x61, 0x79, 0x6d,
	0x65, 0x6e, 0x74, 0x73, 0x18, 0x8c, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x69, 0x6e,
	0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65,
	0x2e, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61,
	0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x08, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x22,
	0x6a, 0x0a, 0x11, 0x43, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x55, 0x0a, 0x05, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x18, 0x0a, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x3f, 0x2e, 0x69, 0x6e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x6d, 0x69,
	0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x61, 0x63, 0x68, 0x69, 0x65, 0x76, 0x65,
	0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x52, 0x65, 0x71, 0x52, 0x05, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x32, 0xa5, 0x01, 0x0a, 0x0a,
	0x4d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x12, 0x96, 0x01, 0x0a, 0x09, 0x43,
	0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x12, 0x31, 0x2e, 0x69, 0x6e, 0x73, 0x70, 0x69,
	0x72, 0x65, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x63, 0x61,
	0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x61, 0x6c, 0x63, 0x75,
	0x6c, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x32, 0x2e, 0x69, 0x6e,
	0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65,
	0x2e, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x61,
	0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x22, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1c, 0x3a, 0x01, 0x2a, 0x22, 0x17, 0x2f, 0x76, 0x31, 0x2f,
	0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x5f, 0x43, 0x61, 0x6c, 0x63, 0x75, 0x6c,
	0x61, 0x74, 0x65, 0x42, 0x40, 0x5a, 0x3e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x4e, 0x70, 0x6f, 0x6f, 0x6c, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x69, 0x6e,
	0x73, 0x70, 0x69, 0x72, 0x65, 0x2f, 0x6d, 0x77, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x61, 0x6c, 0x63,
	0x75, 0x6c, 0x61, 0x74, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_npool_inspire_mw_v1_calculate_calculate_proto_rawDescOnce sync.Once
	file_npool_inspire_mw_v1_calculate_calculate_proto_rawDescData = file_npool_inspire_mw_v1_calculate_calculate_proto_rawDesc
)

func file_npool_inspire_mw_v1_calculate_calculate_proto_rawDescGZIP() []byte {
	file_npool_inspire_mw_v1_calculate_calculate_proto_rawDescOnce.Do(func() {
		file_npool_inspire_mw_v1_calculate_calculate_proto_rawDescData = protoimpl.X.CompressGZIP(file_npool_inspire_mw_v1_calculate_calculate_proto_rawDescData)
	})
	return file_npool_inspire_mw_v1_calculate_calculate_proto_rawDescData
}

var file_npool_inspire_mw_v1_calculate_calculate_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_npool_inspire_mw_v1_calculate_calculate_proto_goTypes = []interface{}{
	(*Payment)(nil),            // 0: inspire.middleware.calculate.v1.Payment
	(*CalculateRequest)(nil),   // 1: inspire.middleware.calculate.v1.CalculateRequest
	(*CalculateResponse)(nil),  // 2: inspire.middleware.calculate.v1.CalculateResponse
	(v1.SettleType)(0),         // 3: basetypes.inspire.v1.SettleType
	(*order.StatementReq)(nil), // 4: inspire.middleware.achievement.statement.order.v1.StatementReq
}
var file_npool_inspire_mw_v1_calculate_calculate_proto_depIdxs = []int32{
	3, // 0: inspire.middleware.calculate.v1.CalculateRequest.SettleType:type_name -> basetypes.inspire.v1.SettleType
	0, // 1: inspire.middleware.calculate.v1.CalculateRequest.Payments:type_name -> inspire.middleware.calculate.v1.Payment
	4, // 2: inspire.middleware.calculate.v1.CalculateResponse.Infos:type_name -> inspire.middleware.achievement.statement.order.v1.StatementReq
	1, // 3: inspire.middleware.calculate.v1.Middleware.Calculate:input_type -> inspire.middleware.calculate.v1.CalculateRequest
	2, // 4: inspire.middleware.calculate.v1.Middleware.Calculate:output_type -> inspire.middleware.calculate.v1.CalculateResponse
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_npool_inspire_mw_v1_calculate_calculate_proto_init() }
func file_npool_inspire_mw_v1_calculate_calculate_proto_init() {
	if File_npool_inspire_mw_v1_calculate_calculate_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_npool_inspire_mw_v1_calculate_calculate_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Payment); i {
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
		file_npool_inspire_mw_v1_calculate_calculate_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CalculateRequest); i {
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
		file_npool_inspire_mw_v1_calculate_calculate_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CalculateResponse); i {
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
			RawDescriptor: file_npool_inspire_mw_v1_calculate_calculate_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_npool_inspire_mw_v1_calculate_calculate_proto_goTypes,
		DependencyIndexes: file_npool_inspire_mw_v1_calculate_calculate_proto_depIdxs,
		MessageInfos:      file_npool_inspire_mw_v1_calculate_calculate_proto_msgTypes,
	}.Build()
	File_npool_inspire_mw_v1_calculate_calculate_proto = out.File
	file_npool_inspire_mw_v1_calculate_calculate_proto_rawDesc = nil
	file_npool_inspire_mw_v1_calculate_calculate_proto_goTypes = nil
	file_npool_inspire_mw_v1_calculate_calculate_proto_depIdxs = nil
}
