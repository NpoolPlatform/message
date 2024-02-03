// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.18.1
// source: npool/scheduler/mw/v1/order/renew/renew.proto

package renew

import (
	coin "github.com/NpoolPlatform/message/npool/chain/mw/v1/app/coin"
	good "github.com/NpoolPlatform/message/npool/good/mw/v1/app/good"
	order "github.com/NpoolPlatform/message/npool/order/mw/v1/order"
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

type Deduction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// If the deduction will be used to create order, then it has app good
	AppGood     *good.Good `protobuf:"bytes,10,opt,name=AppGood,proto3,oneof" json:"AppGood,omitempty"`
	AppCoin     *coin.Coin `protobuf:"bytes,20,opt,name=AppCoin,proto3" json:"AppCoin,omitempty"`
	USDCurrency string     `protobuf:"bytes,30,opt,name=USDCurrency,proto3" json:"USDCurrency,omitempty"`
	Amount      string     `protobuf:"bytes,40,opt,name=Amount,proto3" json:"Amount,omitempty"`
}

func (x *Deduction) Reset() {
	*x = Deduction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_scheduler_mw_v1_order_renew_renew_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Deduction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Deduction) ProtoMessage() {}

func (x *Deduction) ProtoReflect() protoreflect.Message {
	mi := &file_npool_scheduler_mw_v1_order_renew_renew_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Deduction.ProtoReflect.Descriptor instead.
func (*Deduction) Descriptor() ([]byte, []int) {
	return file_npool_scheduler_mw_v1_order_renew_renew_proto_rawDescGZIP(), []int{0}
}

func (x *Deduction) GetAppGood() *good.Good {
	if x != nil {
		return x.AppGood
	}
	return nil
}

func (x *Deduction) GetAppCoin() *coin.Coin {
	if x != nil {
		return x.AppCoin
	}
	return nil
}

func (x *Deduction) GetUSDCurrency() string {
	if x != nil {
		return x.USDCurrency
	}
	return ""
}

func (x *Deduction) GetAmount() string {
	if x != nil {
		return x.Amount
	}
	return ""
}

type RenewInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AppGood *good.Good `protobuf:"bytes,10,opt,name=AppGood,proto3" json:"AppGood,omitempty"`
	EndAt   uint32     `protobuf:"varint,20,opt,name=EndAt,proto3" json:"EndAt,omitempty"`
}

func (x *RenewInfo) Reset() {
	*x = RenewInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_scheduler_mw_v1_order_renew_renew_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RenewInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RenewInfo) ProtoMessage() {}

func (x *RenewInfo) ProtoReflect() protoreflect.Message {
	mi := &file_npool_scheduler_mw_v1_order_renew_renew_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RenewInfo.ProtoReflect.Descriptor instead.
func (*RenewInfo) Descriptor() ([]byte, []int) {
	return file_npool_scheduler_mw_v1_order_renew_renew_proto_rawDescGZIP(), []int{1}
}

func (x *RenewInfo) GetAppGood() *good.Good {
	if x != nil {
		return x.AppGood
	}
	return nil
}

func (x *RenewInfo) GetEndAt() uint32 {
	if x != nil {
		return x.EndAt
	}
	return 0
}

type MsgOrderChildsRenewReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ParentOrder         *order.Order `protobuf:"bytes,10,opt,name=ParentOrder,proto3" json:"ParentOrder,omitempty"`
	Deductions          []*Deduction `protobuf:"bytes,20,rep,name=Deductions,proto3" json:"Deductions,omitempty"`
	InsufficientBalance bool         `protobuf:"varint,30,opt,name=InsufficientBalance,proto3" json:"InsufficientBalance,omitempty"`
	WillCreateOrder     bool         `protobuf:"varint,40,opt,name=WillCreateOrder,proto3" json:"WillCreateOrder,omitempty"`
	RenewInfos          []*RenewInfo `protobuf:"bytes,50,rep,name=RenewInfos,proto3" json:"RenewInfos,omitempty"`
	Error               *string      `protobuf:"bytes,60,opt,name=Error,proto3,oneof" json:"Error,omitempty"`
}

func (x *MsgOrderChildsRenewReq) Reset() {
	*x = MsgOrderChildsRenewReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_scheduler_mw_v1_order_renew_renew_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgOrderChildsRenewReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgOrderChildsRenewReq) ProtoMessage() {}

func (x *MsgOrderChildsRenewReq) ProtoReflect() protoreflect.Message {
	mi := &file_npool_scheduler_mw_v1_order_renew_renew_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MsgOrderChildsRenewReq.ProtoReflect.Descriptor instead.
func (*MsgOrderChildsRenewReq) Descriptor() ([]byte, []int) {
	return file_npool_scheduler_mw_v1_order_renew_renew_proto_rawDescGZIP(), []int{2}
}

func (x *MsgOrderChildsRenewReq) GetParentOrder() *order.Order {
	if x != nil {
		return x.ParentOrder
	}
	return nil
}

func (x *MsgOrderChildsRenewReq) GetDeductions() []*Deduction {
	if x != nil {
		return x.Deductions
	}
	return nil
}

func (x *MsgOrderChildsRenewReq) GetInsufficientBalance() bool {
	if x != nil {
		return x.InsufficientBalance
	}
	return false
}

func (x *MsgOrderChildsRenewReq) GetWillCreateOrder() bool {
	if x != nil {
		return x.WillCreateOrder
	}
	return false
}

func (x *MsgOrderChildsRenewReq) GetRenewInfos() []*RenewInfo {
	if x != nil {
		return x.RenewInfos
	}
	return nil
}

func (x *MsgOrderChildsRenewReq) GetError() string {
	if x != nil && x.Error != nil {
		return *x.Error
	}
	return ""
}

var File_npool_scheduler_mw_v1_order_renew_renew_proto protoreflect.FileDescriptor

var file_npool_scheduler_mw_v1_order_renew_renew_proto_rawDesc = []byte{
	0x0a, 0x2d, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65,
	0x72, 0x2f, 0x6d, 0x77, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2f, 0x72, 0x65,
	0x6e, 0x65, 0x77, 0x2f, 0x72, 0x65, 0x6e, 0x65, 0x77, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x24, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c,
	0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x32, 0x2e, 0x72, 0x65, 0x6e,
	0x65, 0x77, 0x2e, 0x76, 0x31, 0x1a, 0x23, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x2f, 0x6d, 0x77, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2f, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x24, 0x6e, 0x70, 0x6f, 0x6f,
	0x6c, 0x2f, 0x67, 0x6f, 0x6f, 0x64, 0x2f, 0x6d, 0x77, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x70, 0x70,
	0x2f, 0x67, 0x6f, 0x6f, 0x64, 0x2f, 0x67, 0x6f, 0x6f, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x25, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2f, 0x6d, 0x77,
	0x2f, 0x76, 0x31, 0x2f, 0x61, 0x70, 0x70, 0x2f, 0x63, 0x6f, 0x69, 0x6e, 0x2f, 0x63, 0x6f, 0x69,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd2, 0x01, 0x0a, 0x09, 0x44, 0x65, 0x64, 0x75,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x41, 0x0a, 0x07, 0x41, 0x70, 0x70, 0x47, 0x6f, 0x6f, 0x64,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x67, 0x6f, 0x6f, 0x64, 0x2e, 0x6d, 0x69,
	0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x67, 0x6f, 0x6f,
	0x64, 0x31, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x6f, 0x6f, 0x64, 0x48, 0x00, 0x52, 0x07, 0x41, 0x70,
	0x70, 0x47, 0x6f, 0x6f, 0x64, 0x88, 0x01, 0x01, 0x12, 0x3c, 0x0a, 0x07, 0x41, 0x70, 0x70, 0x43,
	0x6f, 0x69, 0x6e, 0x18, 0x14, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x63, 0x68, 0x61, 0x69,
	0x6e, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x61, 0x70, 0x70,
	0x2e, 0x63, 0x6f, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x69, 0x6e, 0x52, 0x07, 0x41,
	0x70, 0x70, 0x43, 0x6f, 0x69, 0x6e, 0x12, 0x20, 0x0a, 0x0b, 0x55, 0x53, 0x44, 0x43, 0x75, 0x72,
	0x72, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x55, 0x53, 0x44,
	0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x41, 0x6d, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x28, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x41, 0x70, 0x70, 0x47, 0x6f, 0x6f, 0x64, 0x22, 0x5f, 0x0a, 0x09,
	0x52, 0x65, 0x6e, 0x65, 0x77, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x3c, 0x0a, 0x07, 0x41, 0x70, 0x70,
	0x47, 0x6f, 0x6f, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x67, 0x6f, 0x6f,
	0x64, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x61, 0x70, 0x70,
	0x2e, 0x67, 0x6f, 0x6f, 0x64, 0x31, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x6f, 0x6f, 0x64, 0x52, 0x07,
	0x41, 0x70, 0x70, 0x47, 0x6f, 0x6f, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x45, 0x6e, 0x64, 0x41, 0x74,
	0x18, 0x14, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x45, 0x6e, 0x64, 0x41, 0x74, 0x22, 0x80, 0x03,
	0x0a, 0x16, 0x4d, 0x73, 0x67, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x43, 0x68, 0x69, 0x6c, 0x64, 0x73,
	0x52, 0x65, 0x6e, 0x65, 0x77, 0x52, 0x65, 0x71, 0x12, 0x43, 0x0a, 0x0b, 0x50, 0x61, 0x72, 0x65,
	0x6e, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65,
	0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x31, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x52, 0x0b, 0x50, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x4f, 0x0a,
	0x0a, 0x44, 0x65, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x14, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x2f, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x2e, 0x6d, 0x69,
	0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x32, 0x2e,
	0x72, 0x65, 0x6e, 0x65, 0x77, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x64, 0x75, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x0a, 0x44, 0x65, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x30,
	0x0a, 0x13, 0x49, 0x6e, 0x73, 0x75, 0x66, 0x66, 0x69, 0x63, 0x69, 0x65, 0x6e, 0x74, 0x42, 0x61,
	0x6c, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x08, 0x52, 0x13, 0x49, 0x6e, 0x73,
	0x75, 0x66, 0x66, 0x69, 0x63, 0x69, 0x65, 0x6e, 0x74, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65,
	0x12, 0x28, 0x0a, 0x0f, 0x57, 0x69, 0x6c, 0x6c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x18, 0x28, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0f, 0x57, 0x69, 0x6c, 0x6c, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x4f, 0x0a, 0x0a, 0x52, 0x65,
	0x6e, 0x65, 0x77, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x18, 0x32, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2f,
	0x2e, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c,
	0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x32, 0x2e, 0x72, 0x65, 0x6e,
	0x65, 0x77, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x6e, 0x65, 0x77, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x0a, 0x52, 0x65, 0x6e, 0x65, 0x77, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x12, 0x19, 0x0a, 0x05, 0x45,
	0x72, 0x72, 0x6f, 0x72, 0x18, 0x3c, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x05, 0x45, 0x72,
	0x72, 0x6f, 0x72, 0x88, 0x01, 0x01, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x45, 0x72, 0x72, 0x6f, 0x72,
	0x42, 0x44, 0x5a, 0x42, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4e,
	0x70, 0x6f, 0x6f, 0x6c, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x2f, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x64,
	0x75, 0x6c, 0x65, 0x72, 0x2f, 0x6d, 0x77, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x2f, 0x72, 0x65, 0x6e, 0x65, 0x77, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_npool_scheduler_mw_v1_order_renew_renew_proto_rawDescOnce sync.Once
	file_npool_scheduler_mw_v1_order_renew_renew_proto_rawDescData = file_npool_scheduler_mw_v1_order_renew_renew_proto_rawDesc
)

func file_npool_scheduler_mw_v1_order_renew_renew_proto_rawDescGZIP() []byte {
	file_npool_scheduler_mw_v1_order_renew_renew_proto_rawDescOnce.Do(func() {
		file_npool_scheduler_mw_v1_order_renew_renew_proto_rawDescData = protoimpl.X.CompressGZIP(file_npool_scheduler_mw_v1_order_renew_renew_proto_rawDescData)
	})
	return file_npool_scheduler_mw_v1_order_renew_renew_proto_rawDescData
}

var file_npool_scheduler_mw_v1_order_renew_renew_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_npool_scheduler_mw_v1_order_renew_renew_proto_goTypes = []interface{}{
	(*Deduction)(nil),              // 0: scheduler.middleware.order2.renew.v1.Deduction
	(*RenewInfo)(nil),              // 1: scheduler.middleware.order2.renew.v1.RenewInfo
	(*MsgOrderChildsRenewReq)(nil), // 2: scheduler.middleware.order2.renew.v1.MsgOrderChildsRenewReq
	(*good.Good)(nil),              // 3: good.middleware.app.good1.v1.Good
	(*coin.Coin)(nil),              // 4: chain.middleware.app.coin.v1.Coin
	(*order.Order)(nil),            // 5: order.middleware.order1.v1.Order
}
var file_npool_scheduler_mw_v1_order_renew_renew_proto_depIdxs = []int32{
	3, // 0: scheduler.middleware.order2.renew.v1.Deduction.AppGood:type_name -> good.middleware.app.good1.v1.Good
	4, // 1: scheduler.middleware.order2.renew.v1.Deduction.AppCoin:type_name -> chain.middleware.app.coin.v1.Coin
	3, // 2: scheduler.middleware.order2.renew.v1.RenewInfo.AppGood:type_name -> good.middleware.app.good1.v1.Good
	5, // 3: scheduler.middleware.order2.renew.v1.MsgOrderChildsRenewReq.ParentOrder:type_name -> order.middleware.order1.v1.Order
	0, // 4: scheduler.middleware.order2.renew.v1.MsgOrderChildsRenewReq.Deductions:type_name -> scheduler.middleware.order2.renew.v1.Deduction
	1, // 5: scheduler.middleware.order2.renew.v1.MsgOrderChildsRenewReq.RenewInfos:type_name -> scheduler.middleware.order2.renew.v1.RenewInfo
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_npool_scheduler_mw_v1_order_renew_renew_proto_init() }
func file_npool_scheduler_mw_v1_order_renew_renew_proto_init() {
	if File_npool_scheduler_mw_v1_order_renew_renew_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_npool_scheduler_mw_v1_order_renew_renew_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Deduction); i {
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
		file_npool_scheduler_mw_v1_order_renew_renew_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RenewInfo); i {
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
		file_npool_scheduler_mw_v1_order_renew_renew_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgOrderChildsRenewReq); i {
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
	file_npool_scheduler_mw_v1_order_renew_renew_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_npool_scheduler_mw_v1_order_renew_renew_proto_msgTypes[2].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_npool_scheduler_mw_v1_order_renew_renew_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_npool_scheduler_mw_v1_order_renew_renew_proto_goTypes,
		DependencyIndexes: file_npool_scheduler_mw_v1_order_renew_renew_proto_depIdxs,
		MessageInfos:      file_npool_scheduler_mw_v1_order_renew_renew_proto_msgTypes,
	}.Build()
	File_npool_scheduler_mw_v1_order_renew_renew_proto = out.File
	file_npool_scheduler_mw_v1_order_renew_renew_proto_rawDesc = nil
	file_npool_scheduler_mw_v1_order_renew_renew_proto_goTypes = nil
	file_npool_scheduler_mw_v1_order_renew_renew_proto_depIdxs = nil
}
