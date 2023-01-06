// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.18.1
// source: npool/inspire/mw/v1/accounting/accounting.proto

package accounting

import (
	commission "github.com/NpoolPlatform/message/npool/inspire/mgr/v1/commission"
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

type Commission struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AppID               string  `protobuf:"bytes,10,opt,name=AppID,proto3" json:"AppID,omitempty"`
	UserID              string  `protobuf:"bytes,20,opt,name=UserID,proto3" json:"UserID,omitempty"`
	DirectContributorID *string `protobuf:"bytes,30,opt,name=DirectContributorID,proto3,oneof" json:"DirectContributorID,omitempty"`
	Amount              string  `protobuf:"bytes,40,opt,name=Amount,proto3" json:"Amount,omitempty"`
}

func (x *Commission) Reset() {
	*x = Commission{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_inspire_mw_v1_accounting_accounting_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Commission) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Commission) ProtoMessage() {}

func (x *Commission) ProtoReflect() protoreflect.Message {
	mi := &file_npool_inspire_mw_v1_accounting_accounting_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Commission.ProtoReflect.Descriptor instead.
func (*Commission) Descriptor() ([]byte, []int) {
	return file_npool_inspire_mw_v1_accounting_accounting_proto_rawDescGZIP(), []int{0}
}

func (x *Commission) GetAppID() string {
	if x != nil {
		return x.AppID
	}
	return ""
}

func (x *Commission) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

func (x *Commission) GetDirectContributorID() string {
	if x != nil && x.DirectContributorID != nil {
		return *x.DirectContributorID
	}
	return ""
}

func (x *Commission) GetAmount() string {
	if x != nil {
		return x.Amount
	}
	return ""
}

type AccountingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AppID                  string                `protobuf:"bytes,10,opt,name=AppID,proto3" json:"AppID,omitempty"`
	UserID                 string                `protobuf:"bytes,20,opt,name=UserID,proto3" json:"UserID,omitempty"`
	GoodID                 string                `protobuf:"bytes,30,opt,name=GoodID,proto3" json:"GoodID,omitempty"`
	OrderID                string                `protobuf:"bytes,40,opt,name=OrderID,proto3" json:"OrderID,omitempty"`
	PaymentID              string                `protobuf:"bytes,50,opt,name=PaymentID,proto3" json:"PaymentID,omitempty"`
	CoinTypeID             string                `protobuf:"bytes,60,opt,name=CoinTypeID,proto3" json:"CoinTypeID,omitempty"`
	PaymentCoinTypeID      string                `protobuf:"bytes,70,opt,name=PaymentCoinTypeID,proto3" json:"PaymentCoinTypeID,omitempty"`
	PaymentCoinUSDCurrency string                `protobuf:"bytes,80,opt,name=PaymentCoinUSDCurrency,proto3" json:"PaymentCoinUSDCurrency,omitempty"`
	Units                  uint32                `protobuf:"varint,90,opt,name=Units,proto3" json:"Units,omitempty"`
	PaymentAmount          string                `protobuf:"bytes,100,opt,name=PaymentAmount,proto3" json:"PaymentAmount,omitempty"`
	GoodValue              string                `protobuf:"bytes,110,opt,name=GoodValue,proto3" json:"GoodValue,omitempty"`
	SettleType             commission.SettleType `protobuf:"varint,120,opt,name=SettleType,proto3,enum=inspire.manager.commission.v1.SettleType" json:"SettleType,omitempty"`
}

func (x *AccountingRequest) Reset() {
	*x = AccountingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_inspire_mw_v1_accounting_accounting_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccountingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccountingRequest) ProtoMessage() {}

func (x *AccountingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_inspire_mw_v1_accounting_accounting_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccountingRequest.ProtoReflect.Descriptor instead.
func (*AccountingRequest) Descriptor() ([]byte, []int) {
	return file_npool_inspire_mw_v1_accounting_accounting_proto_rawDescGZIP(), []int{1}
}

func (x *AccountingRequest) GetAppID() string {
	if x != nil {
		return x.AppID
	}
	return ""
}

func (x *AccountingRequest) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

func (x *AccountingRequest) GetGoodID() string {
	if x != nil {
		return x.GoodID
	}
	return ""
}

func (x *AccountingRequest) GetOrderID() string {
	if x != nil {
		return x.OrderID
	}
	return ""
}

func (x *AccountingRequest) GetPaymentID() string {
	if x != nil {
		return x.PaymentID
	}
	return ""
}

func (x *AccountingRequest) GetCoinTypeID() string {
	if x != nil {
		return x.CoinTypeID
	}
	return ""
}

func (x *AccountingRequest) GetPaymentCoinTypeID() string {
	if x != nil {
		return x.PaymentCoinTypeID
	}
	return ""
}

func (x *AccountingRequest) GetPaymentCoinUSDCurrency() string {
	if x != nil {
		return x.PaymentCoinUSDCurrency
	}
	return ""
}

func (x *AccountingRequest) GetUnits() uint32 {
	if x != nil {
		return x.Units
	}
	return 0
}

func (x *AccountingRequest) GetPaymentAmount() string {
	if x != nil {
		return x.PaymentAmount
	}
	return ""
}

func (x *AccountingRequest) GetGoodValue() string {
	if x != nil {
		return x.GoodValue
	}
	return ""
}

func (x *AccountingRequest) GetSettleType() commission.SettleType {
	if x != nil {
		return x.SettleType
	}
	return commission.SettleType(0)
}

type AccountingResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Infos []*Commission `protobuf:"bytes,10,rep,name=Infos,proto3" json:"Infos,omitempty"`
}

func (x *AccountingResponse) Reset() {
	*x = AccountingResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_inspire_mw_v1_accounting_accounting_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccountingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccountingResponse) ProtoMessage() {}

func (x *AccountingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_inspire_mw_v1_accounting_accounting_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccountingResponse.ProtoReflect.Descriptor instead.
func (*AccountingResponse) Descriptor() ([]byte, []int) {
	return file_npool_inspire_mw_v1_accounting_accounting_proto_rawDescGZIP(), []int{2}
}

func (x *AccountingResponse) GetInfos() []*Commission {
	if x != nil {
		return x.Infos
	}
	return nil
}

var File_npool_inspire_mw_v1_accounting_accounting_proto protoreflect.FileDescriptor

var file_npool_inspire_mw_v1_accounting_accounting_proto_rawDesc = []byte{
	0x0a, 0x2f, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x69, 0x6e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2f,
	0x6d, 0x77, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x69, 0x6e, 0x67,
	0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x20, 0x69, 0x6e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c,
	0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x69, 0x6e, 0x67,
	0x2e, 0x76, 0x31, 0x1a, 0x30, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x69, 0x6e, 0x73, 0x70, 0x69,
	0x72, 0x65, 0x2f, 0x6d, 0x67, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa1, 0x01, 0x0a, 0x0a, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x41, 0x70, 0x70, 0x49, 0x44, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x41, 0x70, 0x70, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73,
	0x65, 0x72, 0x49, 0x44, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72,
	0x49, 0x44, 0x12, 0x35, 0x0a, 0x13, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x43, 0x6f, 0x6e, 0x74,
	0x72, 0x69, 0x62, 0x75, 0x74, 0x6f, 0x72, 0x49, 0x44, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x09, 0x48,
	0x00, 0x52, 0x13, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x69, 0x62,
	0x75, 0x74, 0x6f, 0x72, 0x49, 0x44, 0x88, 0x01, 0x01, 0x12, 0x16, 0x0a, 0x06, 0x41, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x28, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x41, 0x6d, 0x6f, 0x75, 0x6e,
	0x74, 0x42, 0x16, 0x0a, 0x14, 0x5f, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x43, 0x6f, 0x6e, 0x74,
	0x72, 0x69, 0x62, 0x75, 0x74, 0x6f, 0x72, 0x49, 0x44, 0x22, 0xbc, 0x03, 0x0a, 0x11, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x14, 0x0a, 0x05, 0x41, 0x70, 0x70, 0x49, 0x44, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x41, 0x70, 0x70, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18,
	0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x16, 0x0a,
	0x06, 0x47, 0x6f, 0x6f, 0x64, 0x49, 0x44, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x47,
	0x6f, 0x6f, 0x64, 0x49, 0x44, 0x12, 0x18, 0x0a, 0x07, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x44,
	0x18, 0x28, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x44, 0x12,
	0x1c, 0x0a, 0x09, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x32, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x44, 0x12, 0x1e, 0x0a,
	0x0a, 0x43, 0x6f, 0x69, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x49, 0x44, 0x18, 0x3c, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x43, 0x6f, 0x69, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x49, 0x44, 0x12, 0x2c, 0x0a,
	0x11, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x69, 0x6e, 0x54, 0x79, 0x70, 0x65,
	0x49, 0x44, 0x18, 0x46, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e,
	0x74, 0x43, 0x6f, 0x69, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x49, 0x44, 0x12, 0x36, 0x0a, 0x16, 0x50,
	0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x69, 0x6e, 0x55, 0x53, 0x44, 0x43, 0x75, 0x72,
	0x72, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x50, 0x20, 0x01, 0x28, 0x09, 0x52, 0x16, 0x50, 0x61, 0x79,
	0x6d, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x69, 0x6e, 0x55, 0x53, 0x44, 0x43, 0x75, 0x72, 0x72, 0x65,
	0x6e, 0x63, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x55, 0x6e, 0x69, 0x74, 0x73, 0x18, 0x5a, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x05, 0x55, 0x6e, 0x69, 0x74, 0x73, 0x12, 0x24, 0x0a, 0x0d, 0x50, 0x61, 0x79,
	0x6d, 0x65, 0x6e, 0x74, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x64, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0d, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12,
	0x1c, 0x0a, 0x09, 0x47, 0x6f, 0x6f, 0x64, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x6e, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x47, 0x6f, 0x6f, 0x64, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x49, 0x0a,
	0x0a, 0x53, 0x65, 0x74, 0x74, 0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x18, 0x78, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x29, 0x2e, 0x69, 0x6e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x6d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x72, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x76,
	0x31, 0x2e, 0x53, 0x65, 0x74, 0x74, 0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0a, 0x53, 0x65,
	0x74, 0x74, 0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x22, 0x58, 0x0a, 0x12, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x42,
	0x0a, 0x05, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2c, 0x2e,
	0x69, 0x6e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61,
	0x72, 0x65, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31,
	0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x05, 0x49, 0x6e, 0x66,
	0x6f, 0x73, 0x32, 0x87, 0x01, 0x0a, 0x0a, 0x4d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72,
	0x65, 0x12, 0x79, 0x0a, 0x0a, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x69, 0x6e, 0x67, 0x12,
	0x33, 0x2e, 0x69, 0x6e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65,
	0x77, 0x61, 0x72, 0x65, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x69, 0x6e, 0x67, 0x2e,
	0x76, 0x31, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x34, 0x2e, 0x69, 0x6e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x6d,
	0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x69,
	0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x41, 0x5a, 0x3f,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4e, 0x70, 0x6f, 0x6f, 0x6c,
	0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x2f, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x69, 0x6e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2f, 0x6d,
	0x77, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x69, 0x6e, 0x67, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_npool_inspire_mw_v1_accounting_accounting_proto_rawDescOnce sync.Once
	file_npool_inspire_mw_v1_accounting_accounting_proto_rawDescData = file_npool_inspire_mw_v1_accounting_accounting_proto_rawDesc
)

func file_npool_inspire_mw_v1_accounting_accounting_proto_rawDescGZIP() []byte {
	file_npool_inspire_mw_v1_accounting_accounting_proto_rawDescOnce.Do(func() {
		file_npool_inspire_mw_v1_accounting_accounting_proto_rawDescData = protoimpl.X.CompressGZIP(file_npool_inspire_mw_v1_accounting_accounting_proto_rawDescData)
	})
	return file_npool_inspire_mw_v1_accounting_accounting_proto_rawDescData
}

var file_npool_inspire_mw_v1_accounting_accounting_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_npool_inspire_mw_v1_accounting_accounting_proto_goTypes = []interface{}{
	(*Commission)(nil),         // 0: inspire.middleware.accounting.v1.Commission
	(*AccountingRequest)(nil),  // 1: inspire.middleware.accounting.v1.AccountingRequest
	(*AccountingResponse)(nil), // 2: inspire.middleware.accounting.v1.AccountingResponse
	(commission.SettleType)(0), // 3: inspire.manager.commission.v1.SettleType
}
var file_npool_inspire_mw_v1_accounting_accounting_proto_depIdxs = []int32{
	3, // 0: inspire.middleware.accounting.v1.AccountingRequest.SettleType:type_name -> inspire.manager.commission.v1.SettleType
	0, // 1: inspire.middleware.accounting.v1.AccountingResponse.Infos:type_name -> inspire.middleware.accounting.v1.Commission
	1, // 2: inspire.middleware.accounting.v1.Middleware.Accounting:input_type -> inspire.middleware.accounting.v1.AccountingRequest
	2, // 3: inspire.middleware.accounting.v1.Middleware.Accounting:output_type -> inspire.middleware.accounting.v1.AccountingResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_npool_inspire_mw_v1_accounting_accounting_proto_init() }
func file_npool_inspire_mw_v1_accounting_accounting_proto_init() {
	if File_npool_inspire_mw_v1_accounting_accounting_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_npool_inspire_mw_v1_accounting_accounting_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Commission); i {
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
		file_npool_inspire_mw_v1_accounting_accounting_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccountingRequest); i {
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
		file_npool_inspire_mw_v1_accounting_accounting_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccountingResponse); i {
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
	file_npool_inspire_mw_v1_accounting_accounting_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_npool_inspire_mw_v1_accounting_accounting_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_npool_inspire_mw_v1_accounting_accounting_proto_goTypes,
		DependencyIndexes: file_npool_inspire_mw_v1_accounting_accounting_proto_depIdxs,
		MessageInfos:      file_npool_inspire_mw_v1_accounting_accounting_proto_msgTypes,
	}.Build()
	File_npool_inspire_mw_v1_accounting_accounting_proto = out.File
	file_npool_inspire_mw_v1_accounting_accounting_proto_rawDesc = nil
	file_npool_inspire_mw_v1_accounting_accounting_proto_goTypes = nil
	file_npool_inspire_mw_v1_accounting_accounting_proto_depIdxs = nil
}
