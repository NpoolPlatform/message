// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.18.1
// source: npool/signproxy/signproxy.proto

package signproxy

import (
	sphinxplugin "github.com/NpoolPlatform/message/npool/sphinxplugin"
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

type TransactionType int32

const (
	TransactionType_Invalid        TransactionType = 0
	TransactionType_WalletNew      TransactionType = 1 // proxy -> sign
	TransactionType_TransactionNew TransactionType = 2
	TransactionType_Signature      TransactionType = 3 // proxy -> sign
	TransactionType_Balance        TransactionType = 4 // proxy -> plugin
	TransactionType_PreSign        TransactionType = 5 // proxy -> pluign get nonce
	TransactionType_Broadcast      TransactionType = 6 // proxy -> plugin mpool push
	TransactionType_RegisterCoin   TransactionType = 7 // plugin -> proxy
)

// Enum value maps for TransactionType.
var (
	TransactionType_name = map[int32]string{
		0: "Invalid",
		1: "WalletNew",
		2: "TransactionNew",
		3: "Signature",
		4: "Balance",
		5: "PreSign",
		6: "Broadcast",
		7: "RegisterCoin",
	}
	TransactionType_value = map[string]int32{
		"Invalid":        0,
		"WalletNew":      1,
		"TransactionNew": 2,
		"Signature":      3,
		"Balance":        4,
		"PreSign":        5,
		"Broadcast":      6,
		"RegisterCoin":   7,
	}
)

func (x TransactionType) Enum() *TransactionType {
	p := new(TransactionType)
	*p = x
	return p
}

func (x TransactionType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TransactionType) Descriptor() protoreflect.EnumDescriptor {
	return file_npool_signproxy_signproxy_proto_enumTypes[0].Descriptor()
}

func (TransactionType) Type() protoreflect.EnumType {
	return &file_npool_signproxy_signproxy_proto_enumTypes[0]
}

func (x TransactionType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TransactionType.Descriptor instead.
func (TransactionType) EnumDescriptor() ([]byte, []int) {
	return file_npool_signproxy_signproxy_proto_rawDescGZIP(), []int{0}
}

// RegisterCoin ..
type ProxyPluginResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CoinType            sphinxplugin.CoinType         `protobuf:"varint,100,opt,name=CoinType,proto3,enum=sphinx.plugin.v1.CoinType" json:"CoinType,omitempty"`
	TransactionType     TransactionType               `protobuf:"varint,110,opt,name=TransactionType,proto3,enum=sphinx.proxy.v1.TransactionType" json:"TransactionType,omitempty"`
	TransactionIDInsite string                        `protobuf:"bytes,120,opt,name=TransactionIDInsite,proto3" json:"TransactionIDInsite,omitempty"`
	Nonce               uint64                        `protobuf:"varint,130,opt,name=Nonce,proto3" json:"Nonce,omitempty"`
	CID                 string                        `protobuf:"bytes,140,opt,name=CID,proto3" json:"CID,omitempty"`
	Balance             uint64                        `protobuf:"varint,150,opt,name=Balance,proto3" json:"Balance,omitempty"`
	Message             *sphinxplugin.UnsignedMessage `protobuf:"bytes,160,opt,name=Message,proto3" json:"Message,omitempty"`
}

func (x *ProxyPluginResponse) Reset() {
	*x = ProxyPluginResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_signproxy_signproxy_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProxyPluginResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProxyPluginResponse) ProtoMessage() {}

func (x *ProxyPluginResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_signproxy_signproxy_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProxyPluginResponse.ProtoReflect.Descriptor instead.
func (*ProxyPluginResponse) Descriptor() ([]byte, []int) {
	return file_npool_signproxy_signproxy_proto_rawDescGZIP(), []int{0}
}

func (x *ProxyPluginResponse) GetCoinType() sphinxplugin.CoinType {
	if x != nil {
		return x.CoinType
	}
	return sphinxplugin.CoinType(0)
}

func (x *ProxyPluginResponse) GetTransactionType() TransactionType {
	if x != nil {
		return x.TransactionType
	}
	return TransactionType_Invalid
}

func (x *ProxyPluginResponse) GetTransactionIDInsite() string {
	if x != nil {
		return x.TransactionIDInsite
	}
	return ""
}

func (x *ProxyPluginResponse) GetNonce() uint64 {
	if x != nil {
		return x.Nonce
	}
	return 0
}

func (x *ProxyPluginResponse) GetCID() string {
	if x != nil {
		return x.CID
	}
	return ""
}

func (x *ProxyPluginResponse) GetBalance() uint64 {
	if x != nil {
		return x.Balance
	}
	return 0
}

func (x *ProxyPluginResponse) GetMessage() *sphinxplugin.UnsignedMessage {
	if x != nil {
		return x.Message
	}
	return nil
}

// MpoolGetNonce WalletBalance MpoolPush ..
type ProxyPluginRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CoinType            sphinxplugin.CoinType         `protobuf:"varint,100,opt,name=CoinType,proto3,enum=sphinx.plugin.v1.CoinType" json:"CoinType,omitempty"`
	TransactionType     TransactionType               `protobuf:"varint,110,opt,name=TransactionType,proto3,enum=sphinx.proxy.v1.TransactionType" json:"TransactionType,omitempty"`
	TransactionIDInsite string                        `protobuf:"bytes,120,opt,name=TransactionIDInsite,proto3" json:"TransactionIDInsite,omitempty"`
	Address             string                        `protobuf:"bytes,130,opt,name=Address,proto3" json:"Address,omitempty"`
	Message             *sphinxplugin.UnsignedMessage `protobuf:"bytes,140,opt,name=Message,proto3" json:"Message,omitempty"`
	Signature           *sphinxplugin.Signature       `protobuf:"bytes,150,opt,name=Signature,proto3" json:"Signature,omitempty"`
}

func (x *ProxyPluginRequest) Reset() {
	*x = ProxyPluginRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_signproxy_signproxy_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProxyPluginRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProxyPluginRequest) ProtoMessage() {}

func (x *ProxyPluginRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_signproxy_signproxy_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProxyPluginRequest.ProtoReflect.Descriptor instead.
func (*ProxyPluginRequest) Descriptor() ([]byte, []int) {
	return file_npool_signproxy_signproxy_proto_rawDescGZIP(), []int{1}
}

func (x *ProxyPluginRequest) GetCoinType() sphinxplugin.CoinType {
	if x != nil {
		return x.CoinType
	}
	return sphinxplugin.CoinType(0)
}

func (x *ProxyPluginRequest) GetTransactionType() TransactionType {
	if x != nil {
		return x.TransactionType
	}
	return TransactionType_Invalid
}

func (x *ProxyPluginRequest) GetTransactionIDInsite() string {
	if x != nil {
		return x.TransactionIDInsite
	}
	return ""
}

func (x *ProxyPluginRequest) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *ProxyPluginRequest) GetMessage() *sphinxplugin.UnsignedMessage {
	if x != nil {
		return x.Message
	}
	return nil
}

func (x *ProxyPluginRequest) GetSignature() *sphinxplugin.Signature {
	if x != nil {
		return x.Signature
	}
	return nil
}

// Sign WalletNew ..
type ProxySignRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CoinType            sphinxplugin.CoinType         `protobuf:"varint,100,opt,name=CoinType,proto3,enum=sphinx.plugin.v1.CoinType" json:"CoinType,omitempty"`
	TransactionType     TransactionType               `protobuf:"varint,110,opt,name=TransactionType,proto3,enum=sphinx.proxy.v1.TransactionType" json:"TransactionType,omitempty"`
	TransactionIDInsite string                        `protobuf:"bytes,120,opt,name=TransactionIDInsite,proto3" json:"TransactionIDInsite,omitempty"`
	Message             *sphinxplugin.UnsignedMessage `protobuf:"bytes,130,opt,name=Message,proto3" json:"Message,omitempty"`
}

func (x *ProxySignRequest) Reset() {
	*x = ProxySignRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_signproxy_signproxy_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProxySignRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProxySignRequest) ProtoMessage() {}

func (x *ProxySignRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_signproxy_signproxy_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProxySignRequest.ProtoReflect.Descriptor instead.
func (*ProxySignRequest) Descriptor() ([]byte, []int) {
	return file_npool_signproxy_signproxy_proto_rawDescGZIP(), []int{2}
}

func (x *ProxySignRequest) GetCoinType() sphinxplugin.CoinType {
	if x != nil {
		return x.CoinType
	}
	return sphinxplugin.CoinType(0)
}

func (x *ProxySignRequest) GetTransactionType() TransactionType {
	if x != nil {
		return x.TransactionType
	}
	return TransactionType_Invalid
}

func (x *ProxySignRequest) GetTransactionIDInsite() string {
	if x != nil {
		return x.TransactionIDInsite
	}
	return ""
}

func (x *ProxySignRequest) GetMessage() *sphinxplugin.UnsignedMessage {
	if x != nil {
		return x.Message
	}
	return nil
}

type ProxySignResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CoinType            sphinxplugin.CoinType  `protobuf:"varint,100,opt,name=CoinType,proto3,enum=sphinx.plugin.v1.CoinType" json:"CoinType,omitempty"`
	TransactionType     TransactionType        `protobuf:"varint,110,opt,name=TransactionType,proto3,enum=sphinx.proxy.v1.TransactionType" json:"TransactionType,omitempty"`
	TransactionIDInsite string                 `protobuf:"bytes,120,opt,name=TransactionIDInsite,proto3" json:"TransactionIDInsite,omitempty"`
	Info                *ProxySignResponseInfo `protobuf:"bytes,130,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (x *ProxySignResponse) Reset() {
	*x = ProxySignResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_signproxy_signproxy_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProxySignResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProxySignResponse) ProtoMessage() {}

func (x *ProxySignResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_signproxy_signproxy_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProxySignResponse.ProtoReflect.Descriptor instead.
func (*ProxySignResponse) Descriptor() ([]byte, []int) {
	return file_npool_signproxy_signproxy_proto_rawDescGZIP(), []int{3}
}

func (x *ProxySignResponse) GetCoinType() sphinxplugin.CoinType {
	if x != nil {
		return x.CoinType
	}
	return sphinxplugin.CoinType(0)
}

func (x *ProxySignResponse) GetTransactionType() TransactionType {
	if x != nil {
		return x.TransactionType
	}
	return TransactionType_Invalid
}

func (x *ProxySignResponse) GetTransactionIDInsite() string {
	if x != nil {
		return x.TransactionIDInsite
	}
	return ""
}

func (x *ProxySignResponse) GetInfo() *ProxySignResponseInfo {
	if x != nil {
		return x.Info
	}
	return nil
}

type ProxySignResponseInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address   string                        `protobuf:"bytes,100,opt,name=Address,proto3" json:"Address,omitempty"` // create new account address
	Message   *sphinxplugin.UnsignedMessage `protobuf:"bytes,110,opt,name=Message,proto3" json:"Message,omitempty"`
	Signature *sphinxplugin.Signature       `protobuf:"bytes,120,opt,name=Signature,proto3" json:"Signature,omitempty"`
}

func (x *ProxySignResponseInfo) Reset() {
	*x = ProxySignResponseInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_signproxy_signproxy_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProxySignResponseInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProxySignResponseInfo) ProtoMessage() {}

func (x *ProxySignResponseInfo) ProtoReflect() protoreflect.Message {
	mi := &file_npool_signproxy_signproxy_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProxySignResponseInfo.ProtoReflect.Descriptor instead.
func (*ProxySignResponseInfo) Descriptor() ([]byte, []int) {
	return file_npool_signproxy_signproxy_proto_rawDescGZIP(), []int{4}
}

func (x *ProxySignResponseInfo) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *ProxySignResponseInfo) GetMessage() *sphinxplugin.UnsignedMessage {
	if x != nil {
		return x.Message
	}
	return nil
}

func (x *ProxySignResponseInfo) GetSignature() *sphinxplugin.Signature {
	if x != nil {
		return x.Signature
	}
	return nil
}

var File_npool_signproxy_signproxy_proto protoreflect.FileDescriptor

var file_npool_signproxy_signproxy_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x73, 0x69, 0x67, 0x6e, 0x70, 0x72, 0x6f, 0x78,
	0x79, 0x2f, 0x73, 0x69, 0x67, 0x6e, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0f, 0x73, 0x70, 0x68, 0x69, 0x6e, 0x78, 0x2e, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e,
	0x76, 0x31, 0x1a, 0x25, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x73, 0x70, 0x68, 0x69, 0x6e, 0x78,
	0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2f, 0x73, 0x70, 0x68, 0x69, 0x6e, 0x78, 0x70, 0x6c, 0x75,
	0x67, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xce, 0x02, 0x0a, 0x13, 0x50, 0x72,
	0x6f, 0x78, 0x79, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x36, 0x0a, 0x08, 0x43, 0x6f, 0x69, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x18, 0x64, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x73, 0x70, 0x68, 0x69, 0x6e, 0x78, 0x2e, 0x70, 0x6c, 0x75,
	0x67, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x69, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52,
	0x08, 0x43, 0x6f, 0x69, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x4a, 0x0a, 0x0f, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x18, 0x6e, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x20, 0x2e, 0x73, 0x70, 0x68, 0x69, 0x6e, 0x78, 0x2e, 0x70, 0x72, 0x6f, 0x78,
	0x79, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x54, 0x79, 0x70, 0x65, 0x52, 0x0f, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x30, 0x0a, 0x13, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x49, 0x6e, 0x73, 0x69, 0x74, 0x65, 0x18, 0x78, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x13, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49,
	0x44, 0x49, 0x6e, 0x73, 0x69, 0x74, 0x65, 0x12, 0x15, 0x0a, 0x05, 0x4e, 0x6f, 0x6e, 0x63, 0x65,
	0x18, 0x82, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x4e, 0x6f, 0x6e, 0x63, 0x65, 0x12, 0x11,
	0x0a, 0x03, 0x43, 0x49, 0x44, 0x18, 0x8c, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x43, 0x49,
	0x44, 0x12, 0x19, 0x0a, 0x07, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x96, 0x01, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x07, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x3c, 0x0a, 0x07,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0xa0, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21,
	0x2e, 0x73, 0x70, 0x68, 0x69, 0x6e, 0x78, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x76,
	0x31, 0x2e, 0x55, 0x6e, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x52, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0xdf, 0x02, 0x0a, 0x12, 0x50,
	0x72, 0x6f, 0x78, 0x79, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x36, 0x0a, 0x08, 0x43, 0x6f, 0x69, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x18, 0x64, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x73, 0x70, 0x68, 0x69, 0x6e, 0x78, 0x2e, 0x70, 0x6c, 0x75,
	0x67, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x69, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52,
	0x08, 0x43, 0x6f, 0x69, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x4a, 0x0a, 0x0f, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x18, 0x6e, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x20, 0x2e, 0x73, 0x70, 0x68, 0x69, 0x6e, 0x78, 0x2e, 0x70, 0x72, 0x6f, 0x78,
	0x79, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x54, 0x79, 0x70, 0x65, 0x52, 0x0f, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x30, 0x0a, 0x13, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x49, 0x6e, 0x73, 0x69, 0x74, 0x65, 0x18, 0x78, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x13, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49,
	0x44, 0x49, 0x6e, 0x73, 0x69, 0x74, 0x65, 0x12, 0x19, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x18, 0x82, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x41, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x12, 0x3c, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x8c, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x73, 0x70, 0x68, 0x69, 0x6e, 0x78, 0x2e, 0x70, 0x6c,
	0x75, 0x67, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x6e, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x64,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x12, 0x3a, 0x0a, 0x09, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x96, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x73, 0x70, 0x68, 0x69, 0x6e, 0x78, 0x2e, 0x70, 0x6c,
	0x75, 0x67, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72,
	0x65, 0x52, 0x09, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x22, 0x86, 0x02, 0x0a,
	0x10, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x53, 0x69, 0x67, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x36, 0x0a, 0x08, 0x43, 0x6f, 0x69, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x18, 0x64, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x73, 0x70, 0x68, 0x69, 0x6e, 0x78, 0x2e, 0x70, 0x6c, 0x75,
	0x67, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x69, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52,
	0x08, 0x43, 0x6f, 0x69, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x4a, 0x0a, 0x0f, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x18, 0x6e, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x20, 0x2e, 0x73, 0x70, 0x68, 0x69, 0x6e, 0x78, 0x2e, 0x70, 0x72, 0x6f, 0x78,
	0x79, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x54, 0x79, 0x70, 0x65, 0x52, 0x0f, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x30, 0x0a, 0x13, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x49, 0x6e, 0x73, 0x69, 0x74, 0x65, 0x18, 0x78, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x13, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49,
	0x44, 0x49, 0x6e, 0x73, 0x69, 0x74, 0x65, 0x12, 0x3c, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x18, 0x82, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x73, 0x70, 0x68, 0x69,
	0x6e, 0x78, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x6e, 0x73,
	0x69, 0x67, 0x6e, 0x65, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x07, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x86, 0x02, 0x0a, 0x11, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x53,
	0x69, 0x67, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x36, 0x0a, 0x08, 0x43,
	0x6f, 0x69, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x18, 0x64, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1a, 0x2e,
	0x73, 0x70, 0x68, 0x69, 0x6e, 0x78, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x76, 0x31,
	0x2e, 0x43, 0x6f, 0x69, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x08, 0x43, 0x6f, 0x69, 0x6e, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x4a, 0x0a, 0x0f, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x18, 0x6e, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x20, 0x2e, 0x73,
	0x70, 0x68, 0x69, 0x6e, 0x78, 0x2e, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0f,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x30, 0x0a, 0x13, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44,
	0x49, 0x6e, 0x73, 0x69, 0x74, 0x65, 0x18, 0x78, 0x20, 0x01, 0x28, 0x09, 0x52, 0x13, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x49, 0x6e, 0x73, 0x69, 0x74,
	0x65, 0x12, 0x3b, 0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x82, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x26, 0x2e, 0x73, 0x70, 0x68, 0x69, 0x6e, 0x78, 0x2e, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e,
	0x76, 0x31, 0x2e, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x53, 0x69, 0x67, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0xa9,
	0x01, 0x0a, 0x15, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x53, 0x69, 0x67, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x18, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x18, 0x64, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x41, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x12, 0x3b, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x6e, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x73, 0x70, 0x68, 0x69, 0x6e, 0x78, 0x2e, 0x70, 0x6c, 0x75,
	0x67, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x6e, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12,
	0x39, 0x0a, 0x09, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x78, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x73, 0x70, 0x68, 0x69, 0x6e, 0x78, 0x2e, 0x70, 0x6c, 0x75, 0x67,
	0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x52,
	0x09, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x2a, 0x8b, 0x01, 0x0a, 0x0f, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0b,
	0x0a, 0x07, 0x49, 0x6e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09, 0x57,
	0x61, 0x6c, 0x6c, 0x65, 0x74, 0x4e, 0x65, 0x77, 0x10, 0x01, 0x12, 0x12, 0x0a, 0x0e, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x65, 0x77, 0x10, 0x02, 0x12, 0x0d,
	0x0a, 0x09, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x10, 0x03, 0x12, 0x0b, 0x0a,
	0x07, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x10, 0x04, 0x12, 0x0b, 0x0a, 0x07, 0x50, 0x72,
	0x65, 0x53, 0x69, 0x67, 0x6e, 0x10, 0x05, 0x12, 0x0d, 0x0a, 0x09, 0x42, 0x72, 0x6f, 0x61, 0x64,
	0x63, 0x61, 0x73, 0x74, 0x10, 0x06, 0x12, 0x10, 0x0a, 0x0c, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x65, 0x72, 0x43, 0x6f, 0x69, 0x6e, 0x10, 0x07, 0x32, 0xc5, 0x01, 0x0a, 0x09, 0x53, 0x69, 0x67,
	0x6e, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x12, 0x5e, 0x0a, 0x0b, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x50,
	0x6c, 0x75, 0x67, 0x69, 0x6e, 0x12, 0x24, 0x2e, 0x73, 0x70, 0x68, 0x69, 0x6e, 0x78, 0x2e, 0x70,
	0x72, 0x6f, 0x78, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x50, 0x6c, 0x75,
	0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x1a, 0x23, 0x2e, 0x73, 0x70,
	0x68, 0x69, 0x6e, 0x78, 0x2e, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72,
	0x6f, 0x78, 0x79, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x12, 0x58, 0x0a, 0x09, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x53,
	0x69, 0x67, 0x6e, 0x12, 0x22, 0x2e, 0x73, 0x70, 0x68, 0x69, 0x6e, 0x78, 0x2e, 0x70, 0x72, 0x6f,
	0x78, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x53, 0x69, 0x67, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x1a, 0x21, 0x2e, 0x73, 0x70, 0x68, 0x69, 0x6e, 0x78,
	0x2e, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x53,
	0x69, 0x67, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01,
	0x42, 0x32, 0x5a, 0x30, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4e,
	0x70, 0x6f, 0x6f, 0x6c, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x2f, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x73, 0x69, 0x67, 0x6e, 0x70,
	0x72, 0x6f, 0x78, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_npool_signproxy_signproxy_proto_rawDescOnce sync.Once
	file_npool_signproxy_signproxy_proto_rawDescData = file_npool_signproxy_signproxy_proto_rawDesc
)

func file_npool_signproxy_signproxy_proto_rawDescGZIP() []byte {
	file_npool_signproxy_signproxy_proto_rawDescOnce.Do(func() {
		file_npool_signproxy_signproxy_proto_rawDescData = protoimpl.X.CompressGZIP(file_npool_signproxy_signproxy_proto_rawDescData)
	})
	return file_npool_signproxy_signproxy_proto_rawDescData
}

var file_npool_signproxy_signproxy_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_npool_signproxy_signproxy_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_npool_signproxy_signproxy_proto_goTypes = []interface{}{
	(TransactionType)(0),                 // 0: sphinx.proxy.v1.TransactionType
	(*ProxyPluginResponse)(nil),          // 1: sphinx.proxy.v1.ProxyPluginResponse
	(*ProxyPluginRequest)(nil),           // 2: sphinx.proxy.v1.ProxyPluginRequest
	(*ProxySignRequest)(nil),             // 3: sphinx.proxy.v1.ProxySignRequest
	(*ProxySignResponse)(nil),            // 4: sphinx.proxy.v1.ProxySignResponse
	(*ProxySignResponseInfo)(nil),        // 5: sphinx.proxy.v1.ProxySignResponseInfo
	(sphinxplugin.CoinType)(0),           // 6: sphinx.plugin.v1.CoinType
	(*sphinxplugin.UnsignedMessage)(nil), // 7: sphinx.plugin.v1.UnsignedMessage
	(*sphinxplugin.Signature)(nil),       // 8: sphinx.plugin.v1.Signature
}
var file_npool_signproxy_signproxy_proto_depIdxs = []int32{
	6,  // 0: sphinx.proxy.v1.ProxyPluginResponse.CoinType:type_name -> sphinx.plugin.v1.CoinType
	0,  // 1: sphinx.proxy.v1.ProxyPluginResponse.TransactionType:type_name -> sphinx.proxy.v1.TransactionType
	7,  // 2: sphinx.proxy.v1.ProxyPluginResponse.Message:type_name -> sphinx.plugin.v1.UnsignedMessage
	6,  // 3: sphinx.proxy.v1.ProxyPluginRequest.CoinType:type_name -> sphinx.plugin.v1.CoinType
	0,  // 4: sphinx.proxy.v1.ProxyPluginRequest.TransactionType:type_name -> sphinx.proxy.v1.TransactionType
	7,  // 5: sphinx.proxy.v1.ProxyPluginRequest.Message:type_name -> sphinx.plugin.v1.UnsignedMessage
	8,  // 6: sphinx.proxy.v1.ProxyPluginRequest.Signature:type_name -> sphinx.plugin.v1.Signature
	6,  // 7: sphinx.proxy.v1.ProxySignRequest.CoinType:type_name -> sphinx.plugin.v1.CoinType
	0,  // 8: sphinx.proxy.v1.ProxySignRequest.TransactionType:type_name -> sphinx.proxy.v1.TransactionType
	7,  // 9: sphinx.proxy.v1.ProxySignRequest.Message:type_name -> sphinx.plugin.v1.UnsignedMessage
	6,  // 10: sphinx.proxy.v1.ProxySignResponse.CoinType:type_name -> sphinx.plugin.v1.CoinType
	0,  // 11: sphinx.proxy.v1.ProxySignResponse.TransactionType:type_name -> sphinx.proxy.v1.TransactionType
	5,  // 12: sphinx.proxy.v1.ProxySignResponse.Info:type_name -> sphinx.proxy.v1.ProxySignResponseInfo
	7,  // 13: sphinx.proxy.v1.ProxySignResponseInfo.Message:type_name -> sphinx.plugin.v1.UnsignedMessage
	8,  // 14: sphinx.proxy.v1.ProxySignResponseInfo.Signature:type_name -> sphinx.plugin.v1.Signature
	1,  // 15: sphinx.proxy.v1.SignProxy.ProxyPlugin:input_type -> sphinx.proxy.v1.ProxyPluginResponse
	4,  // 16: sphinx.proxy.v1.SignProxy.ProxySign:input_type -> sphinx.proxy.v1.ProxySignResponse
	2,  // 17: sphinx.proxy.v1.SignProxy.ProxyPlugin:output_type -> sphinx.proxy.v1.ProxyPluginRequest
	3,  // 18: sphinx.proxy.v1.SignProxy.ProxySign:output_type -> sphinx.proxy.v1.ProxySignRequest
	17, // [17:19] is the sub-list for method output_type
	15, // [15:17] is the sub-list for method input_type
	15, // [15:15] is the sub-list for extension type_name
	15, // [15:15] is the sub-list for extension extendee
	0,  // [0:15] is the sub-list for field type_name
}

func init() { file_npool_signproxy_signproxy_proto_init() }
func file_npool_signproxy_signproxy_proto_init() {
	if File_npool_signproxy_signproxy_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_npool_signproxy_signproxy_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProxyPluginResponse); i {
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
		file_npool_signproxy_signproxy_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProxyPluginRequest); i {
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
		file_npool_signproxy_signproxy_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProxySignRequest); i {
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
		file_npool_signproxy_signproxy_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProxySignResponse); i {
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
		file_npool_signproxy_signproxy_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProxySignResponseInfo); i {
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
			RawDescriptor: file_npool_signproxy_signproxy_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_npool_signproxy_signproxy_proto_goTypes,
		DependencyIndexes: file_npool_signproxy_signproxy_proto_depIdxs,
		EnumInfos:         file_npool_signproxy_signproxy_proto_enumTypes,
		MessageInfos:      file_npool_signproxy_signproxy_proto_msgTypes,
	}.Build()
	File_npool_signproxy_signproxy_proto = out.File
	file_npool_signproxy_signproxy_proto_rawDesc = nil
	file_npool_signproxy_signproxy_proto_goTypes = nil
	file_npool_signproxy_signproxy_proto_depIdxs = nil
}
