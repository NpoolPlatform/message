// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.18.1
// source: npool/inspire/mw/v1/calculate/calculate.proto

package calculate

import (
	v1 "github.com/NpoolPlatform/message/npool/basetypes/inspire/v1"
<<<<<<< HEAD
<<<<<<< HEAD
	statement "github.com/NpoolPlatform/message/npool/inspire/mw/v1/achievement/statement"
=======
	statement "github.com/NpoolPlatform/message/npool/inspire/mw/v1/archivement/statement"
>>>>>>> Move accounting to calculate
=======
	statement "github.com/NpoolPlatform/message/npool/inspire/mw/v1/achivement/statement"
>>>>>>> Correct achivement
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

type CalculateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AppID                  string        `protobuf:"bytes,10,opt,name=AppID,proto3" json:"AppID,omitempty"`
	UserID                 string        `protobuf:"bytes,20,opt,name=UserID,proto3" json:"UserID,omitempty"`
	GoodID                 string        `protobuf:"bytes,30,opt,name=GoodID,proto3" json:"GoodID,omitempty"`
	OrderID                string        `protobuf:"bytes,40,opt,name=OrderID,proto3" json:"OrderID,omitempty"`
	PaymentID              string        `protobuf:"bytes,50,opt,name=PaymentID,proto3" json:"PaymentID,omitempty"`
	CoinTypeID             string        `protobuf:"bytes,60,opt,name=CoinTypeID,proto3" json:"CoinTypeID,omitempty"`
	PaymentCoinTypeID      string        `protobuf:"bytes,70,opt,name=PaymentCoinTypeID,proto3" json:"PaymentCoinTypeID,omitempty"`
	PaymentCoinUSDCurrency string        `protobuf:"bytes,80,opt,name=PaymentCoinUSDCurrency,proto3" json:"PaymentCoinUSDCurrency,omitempty"`
	Units                  string        `protobuf:"bytes,90,opt,name=Units,proto3" json:"Units,omitempty"`
	PaymentAmount          string        `protobuf:"bytes,100,opt,name=PaymentAmount,proto3" json:"PaymentAmount,omitempty"`
	GoodValue              string        `protobuf:"bytes,110,opt,name=GoodValue,proto3" json:"GoodValue,omitempty"`
	SettleType             v1.SettleType `protobuf:"varint,120,opt,name=SettleType,proto3,enum=basetypes.inspire.v1.SettleType" json:"SettleType,omitempty"`
	HasCommission          bool          `protobuf:"varint,130,opt,name=HasCommission,proto3" json:"HasCommission,omitempty"`
	OrderCreatedAt         uint32        `protobuf:"varint,140,opt,name=OrderCreatedAt,proto3" json:"OrderCreatedAt,omitempty"`
	SettleMode             v1.SettleMode `protobuf:"varint,150,opt,name=SettleMode,proto3,enum=basetypes.inspire.v1.SettleMode" json:"SettleMode,omitempty"`
}

func (x *CalculateRequest) Reset() {
	*x = CalculateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_inspire_mw_v1_calculate_calculate_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CalculateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CalculateRequest) ProtoMessage() {}

func (x *CalculateRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use CalculateRequest.ProtoReflect.Descriptor instead.
func (*CalculateRequest) Descriptor() ([]byte, []int) {
	return file_npool_inspire_mw_v1_calculate_calculate_proto_rawDescGZIP(), []int{0}
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

func (x *CalculateRequest) GetOrderID() string {
	if x != nil {
		return x.OrderID
	}
	return ""
}

func (x *CalculateRequest) GetPaymentID() string {
	if x != nil {
		return x.PaymentID
	}
	return ""
}

func (x *CalculateRequest) GetCoinTypeID() string {
	if x != nil {
		return x.CoinTypeID
	}
	return ""
}

func (x *CalculateRequest) GetPaymentCoinTypeID() string {
	if x != nil {
		return x.PaymentCoinTypeID
	}
	return ""
}

func (x *CalculateRequest) GetPaymentCoinUSDCurrency() string {
	if x != nil {
		return x.PaymentCoinUSDCurrency
	}
	return ""
}

func (x *CalculateRequest) GetUnits() string {
	if x != nil {
		return x.Units
	}
	return ""
}

func (x *CalculateRequest) GetPaymentAmount() string {
	if x != nil {
		return x.PaymentAmount
	}
	return ""
}

func (x *CalculateRequest) GetGoodValue() string {
	if x != nil {
		return x.GoodValue
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

func (x *CalculateRequest) GetSettleMode() v1.SettleMode {
	if x != nil {
		return x.SettleMode
	}
	return v1.SettleMode(0)
}

type CalculateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Infos []*statement.Statement `protobuf:"bytes,10,rep,name=Infos,proto3" json:"Infos,omitempty"`
}

func (x *CalculateResponse) Reset() {
	*x = CalculateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_inspire_mw_v1_calculate_calculate_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CalculateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CalculateResponse) ProtoMessage() {}

func (x *CalculateResponse) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use CalculateResponse.ProtoReflect.Descriptor instead.
func (*CalculateResponse) Descriptor() ([]byte, []int) {
	return file_npool_inspire_mw_v1_calculate_calculate_proto_rawDescGZIP(), []int{1}
}

func (x *CalculateResponse) GetInfos() []*statement.Statement {
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
	0x1a, 0x26, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x2f, 0x69, 0x6e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x6e, 0x75,
<<<<<<< HEAD
	0x6d, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x39, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f,
<<<<<<< HEAD
	0x69, 0x6e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2f, 0x6d, 0x77, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x63,
	0x68, 0x69, 0x65, 0x76, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x6d,
=======
	0x69, 0x6e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2f, 0x6d, 0x77, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x72,
	0x63, 0x68, 0x69, 0x76, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x6d,
>>>>>>> Move accounting to calculate
	0x65, 0x6e, 0x74, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xc5, 0x04, 0x0a, 0x10, 0x43, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x41, 0x70, 0x70, 0x49,
	0x44, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x41, 0x70, 0x70, 0x49, 0x44, 0x12, 0x16,
	0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x47, 0x6f, 0x6f, 0x64, 0x49, 0x44,
	0x18, 0x1e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x47, 0x6f, 0x6f, 0x64, 0x49, 0x44, 0x12, 0x18,
	0x0a, 0x07, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x44, 0x18, 0x28, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x44, 0x12, 0x1c, 0x0a, 0x09, 0x50, 0x61, 0x79, 0x6d,
	0x65, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x32, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x50, 0x61, 0x79,
	0x6d, 0x65, 0x6e, 0x74, 0x49, 0x44, 0x12, 0x1e, 0x0a, 0x0a, 0x43, 0x6f, 0x69, 0x6e, 0x54, 0x79,
	0x70, 0x65, 0x49, 0x44, 0x18, 0x3c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x43, 0x6f, 0x69, 0x6e,
	0x54, 0x79, 0x70, 0x65, 0x49, 0x44, 0x12, 0x2c, 0x0a, 0x11, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e,
	0x74, 0x43, 0x6f, 0x69, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x49, 0x44, 0x18, 0x46, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x11, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x69, 0x6e, 0x54, 0x79,
	0x70, 0x65, 0x49, 0x44, 0x12, 0x36, 0x0a, 0x16, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x43,
	0x6f, 0x69, 0x6e, 0x55, 0x53, 0x44, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x50,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x16, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x69,
	0x6e, 0x55, 0x53, 0x44, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x14, 0x0a, 0x05,
	0x55, 0x6e, 0x69, 0x74, 0x73, 0x18, 0x5a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x55, 0x6e, 0x69,
	0x74, 0x73, 0x12, 0x24, 0x0a, 0x0d, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x41, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x64, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x50, 0x61, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x47, 0x6f, 0x6f, 0x64,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x6e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x47, 0x6f, 0x6f,
	0x64, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x40, 0x0a, 0x0a, 0x53, 0x65, 0x74, 0x74, 0x6c, 0x65,
	0x54, 0x79, 0x70, 0x65, 0x18, 0x78, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x20, 0x2e, 0x62, 0x61, 0x73,
	0x65, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x69, 0x6e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x53, 0x65, 0x74, 0x74, 0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0a, 0x53, 0x65,
	0x74, 0x74, 0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x25, 0x0a, 0x0d, 0x48, 0x61, 0x73, 0x43,
	0x6f, 0x6d, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x82, 0x01, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x0d, 0x48, 0x61, 0x73, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12,
	0x27, 0x0a, 0x0e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x18, 0x8c, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x41, 0x0a, 0x0a, 0x53, 0x65, 0x74, 0x74,
	0x6c, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x18, 0x96, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x20, 0x2e,
	0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x69, 0x6e, 0x73, 0x70, 0x69, 0x72,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x74, 0x74, 0x6c, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x52,
	0x0a, 0x53, 0x65, 0x74, 0x74, 0x6c, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x22, 0x61, 0x0a, 0x11, 0x43,
	0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x4c, 0x0a, 0x05, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x36, 0x2e, 0x69, 0x6e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65,
<<<<<<< HEAD
	0x77, 0x61, 0x72, 0x65, 0x2e, 0x61, 0x63, 0x68, 0x69, 0x65, 0x76, 0x65, 0x6d, 0x65, 0x6e, 0x74,
=======
	0x77, 0x61, 0x72, 0x65, 0x2e, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x6d, 0x65, 0x6e, 0x74,
>>>>>>> Move accounting to calculate
	0x2e, 0x73, 0x74, 0x61, 0x74, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74,
	0x61, 0x74, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x05, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x32, 0x82,
	0x01, 0x0a, 0x0a, 0x4d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x12, 0x74, 0x0a,
	0x09, 0x43, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x12, 0x31, 0x2e, 0x69, 0x6e, 0x73,
	0x70, 0x69, 0x72, 0x65, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e,
	0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x61, 0x6c,
	0x63, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x32, 0x2e,
	0x69, 0x6e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61,
	0x72, 0x65, 0x2e, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x43, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x42, 0x40, 0x5a, 0x3e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x4e, 0x70, 0x6f, 0x6f, 0x6c, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x69, 0x6e,
	0x73, 0x70, 0x69, 0x72, 0x65, 0x2f, 0x6d, 0x77, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x61, 0x6c, 0x63,
	0x75, 0x6c, 0x61, 0x74, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
=======
	0x6d, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x38, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f,
	0x69, 0x6e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2f, 0x6d, 0x77, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x63,
	0x68, 0x69, 0x76, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x6d, 0x65,
	0x6e, 0x74, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xc5, 0x04, 0x0a, 0x10, 0x43, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x41, 0x70, 0x70, 0x49, 0x44,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x41, 0x70, 0x70, 0x49, 0x44, 0x12, 0x16, 0x0a,
	0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x55,
	0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x47, 0x6f, 0x6f, 0x64, 0x49, 0x44, 0x18,
	0x1e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x47, 0x6f, 0x6f, 0x64, 0x49, 0x44, 0x12, 0x18, 0x0a,
	0x07, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x44, 0x18, 0x28, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x44, 0x12, 0x1c, 0x0a, 0x09, 0x50, 0x61, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x49, 0x44, 0x18, 0x32, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x50, 0x61, 0x79, 0x6d,
	0x65, 0x6e, 0x74, 0x49, 0x44, 0x12, 0x1e, 0x0a, 0x0a, 0x43, 0x6f, 0x69, 0x6e, 0x54, 0x79, 0x70,
	0x65, 0x49, 0x44, 0x18, 0x3c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x43, 0x6f, 0x69, 0x6e, 0x54,
	0x79, 0x70, 0x65, 0x49, 0x44, 0x12, 0x2c, 0x0a, 0x11, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x43, 0x6f, 0x69, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x49, 0x44, 0x18, 0x46, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x11, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x69, 0x6e, 0x54, 0x79, 0x70,
	0x65, 0x49, 0x44, 0x12, 0x36, 0x0a, 0x16, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x43, 0x6f,
	0x69, 0x6e, 0x55, 0x53, 0x44, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x50, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x16, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x69, 0x6e,
	0x55, 0x53, 0x44, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x55,
	0x6e, 0x69, 0x74, 0x73, 0x18, 0x5a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x55, 0x6e, 0x69, 0x74,
	0x73, 0x12, 0x24, 0x0a, 0x0d, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x41, 0x6d, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x64, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e,
	0x74, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x47, 0x6f, 0x6f, 0x64, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x6e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x47, 0x6f, 0x6f, 0x64,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x40, 0x0a, 0x0a, 0x53, 0x65, 0x74, 0x74, 0x6c, 0x65, 0x54,
	0x79, 0x70, 0x65, 0x18, 0x78, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x20, 0x2e, 0x62, 0x61, 0x73, 0x65,
	0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x69, 0x6e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x53, 0x65, 0x74, 0x74, 0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0a, 0x53, 0x65, 0x74,
	0x74, 0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x25, 0x0a, 0x0d, 0x48, 0x61, 0x73, 0x43, 0x6f,
	0x6d, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x82, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x0d, 0x48, 0x61, 0x73, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x27,
	0x0a, 0x0e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x18, 0x8c, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x41, 0x0a, 0x0a, 0x53, 0x65, 0x74, 0x74, 0x6c,
	0x65, 0x4d, 0x6f, 0x64, 0x65, 0x18, 0x96, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x20, 0x2e, 0x62,
	0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x69, 0x6e, 0x73, 0x70, 0x69, 0x72, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x74, 0x74, 0x6c, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x52, 0x0a,
	0x53, 0x65, 0x74, 0x74, 0x6c, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x22, 0x60, 0x0a, 0x11, 0x43, 0x61,
	0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x4b, 0x0a, 0x05, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x35,
	0x2e, 0x69, 0x6e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77,
	0x61, 0x72, 0x65, 0x2e, 0x61, 0x63, 0x68, 0x69, 0x76, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x73,
	0x74, 0x61, 0x74, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x61, 0x74,
	0x65, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x05, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x32, 0x82, 0x01, 0x0a,
	0x0a, 0x4d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x12, 0x74, 0x0a, 0x09, 0x43,
	0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x12, 0x31, 0x2e, 0x69, 0x6e, 0x73, 0x70, 0x69,
	0x72, 0x65, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x63, 0x61,
	0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x61, 0x6c, 0x63, 0x75,
	0x6c, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x32, 0x2e, 0x69, 0x6e,
	0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65,
	0x2e, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x61,
	0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x42, 0x40, 0x5a, 0x3e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x4e, 0x70, 0x6f, 0x6f, 0x6c, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x69, 0x6e, 0x73, 0x70,
	0x69, 0x72, 0x65, 0x2f, 0x6d, 0x77, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c,
	0x61, 0x74, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
>>>>>>> Correct achivement
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

var file_npool_inspire_mw_v1_calculate_calculate_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_npool_inspire_mw_v1_calculate_calculate_proto_goTypes = []interface{}{
	(*CalculateRequest)(nil),    // 0: inspire.middleware.calculate.v1.CalculateRequest
	(*CalculateResponse)(nil),   // 1: inspire.middleware.calculate.v1.CalculateResponse
	(v1.SettleType)(0),          // 2: basetypes.inspire.v1.SettleType
	(v1.SettleMode)(0),          // 3: basetypes.inspire.v1.SettleMode
<<<<<<< HEAD
<<<<<<< HEAD
	(*statement.Statement)(nil), // 4: inspire.middleware.achievement.statement.v1.Statement
=======
	(*statement.Statement)(nil), // 4: inspire.middleware.archivement.statement.v1.Statement
>>>>>>> Move accounting to calculate
=======
	(*statement.Statement)(nil), // 4: inspire.middleware.achivement.statement.v1.Statement
>>>>>>> Correct achivement
}
var file_npool_inspire_mw_v1_calculate_calculate_proto_depIdxs = []int32{
	2, // 0: inspire.middleware.calculate.v1.CalculateRequest.SettleType:type_name -> basetypes.inspire.v1.SettleType
	3, // 1: inspire.middleware.calculate.v1.CalculateRequest.SettleMode:type_name -> basetypes.inspire.v1.SettleMode
<<<<<<< HEAD
<<<<<<< HEAD
	4, // 2: inspire.middleware.calculate.v1.CalculateResponse.Infos:type_name -> inspire.middleware.achievement.statement.v1.Statement
=======
	4, // 2: inspire.middleware.calculate.v1.CalculateResponse.Infos:type_name -> inspire.middleware.archivement.statement.v1.Statement
>>>>>>> Move accounting to calculate
=======
	4, // 2: inspire.middleware.calculate.v1.CalculateResponse.Infos:type_name -> inspire.middleware.achivement.statement.v1.Statement
>>>>>>> Correct achivement
	0, // 3: inspire.middleware.calculate.v1.Middleware.Calculate:input_type -> inspire.middleware.calculate.v1.CalculateRequest
	1, // 4: inspire.middleware.calculate.v1.Middleware.Calculate:output_type -> inspire.middleware.calculate.v1.CalculateResponse
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
		file_npool_inspire_mw_v1_calculate_calculate_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
			NumMessages:   2,
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
