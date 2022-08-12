// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.18.1
// source: npool/build-chain/buildchain.proto

package proto

import (
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

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_build_chain_buildchain_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_npool_build_chain_buildchain_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_npool_build_chain_buildchain_proto_rawDescGZIP(), []int{0}
}

type VersionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info string `protobuf:"bytes,100,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (x *VersionResponse) Reset() {
	*x = VersionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_build_chain_buildchain_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VersionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VersionResponse) ProtoMessage() {}

func (x *VersionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_build_chain_buildchain_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VersionResponse.ProtoReflect.Descriptor instead.
func (*VersionResponse) Descriptor() ([]byte, []int) {
	return file_npool_build_chain_buildchain_proto_rawDescGZIP(), []int{1}
}

func (x *VersionResponse) GetInfo() string {
	if x != nil {
		return x.Info
	}
	return ""
}

type DeployedCoin struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID       string `protobuf:"bytes,10,opt,name=ID,proto3" json:"ID,omitempty"`
	CoinID   string `protobuf:"bytes,20,opt,name=CoinID,proto3" json:"CoinID,omitempty"`
	Contract string `protobuf:"bytes,30,opt,name=Contract,proto3" json:"Contract,omitempty"`
}

func (x *DeployedCoin) Reset() {
	*x = DeployedCoin{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_build_chain_buildchain_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeployedCoin) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeployedCoin) ProtoMessage() {}

func (x *DeployedCoin) ProtoReflect() protoreflect.Message {
	mi := &file_npool_build_chain_buildchain_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeployedCoin.ProtoReflect.Descriptor instead.
func (*DeployedCoin) Descriptor() ([]byte, []int) {
	return file_npool_build_chain_buildchain_proto_rawDescGZIP(), []int{2}
}

func (x *DeployedCoin) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *DeployedCoin) GetCoinID() string {
	if x != nil {
		return x.CoinID
	}
	return ""
}

func (x *DeployedCoin) GetContract() string {
	if x != nil {
		return x.Contract
	}
	return ""
}

type CoinsInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID         string `protobuf:"bytes,10,opt,name=ID,proto3" json:"ID,omitempty"`
	Name       string `protobuf:"bytes,20,opt,name=Name,proto3" json:"Name,omitempty"`
	ChainType  string `protobuf:"bytes,30,opt,name=ChainType,proto3" json:"ChainType,omitempty"`
	TokenType  string `protobuf:"bytes,40,opt,name=TokenType,proto3" json:"TokenType,omitempty"`
	Contract   string `protobuf:"bytes,50,opt,name=Contract,proto3" json:"Contract,omitempty"`
	Similarity int32  `protobuf:"varint,60,opt,name=Similarity,proto3" json:"Similarity,omitempty"`
	Remark     string `protobuf:"bytes,70,opt,name=Remark,proto3" json:"Remark,omitempty"`
	Data       []byte `protobuf:"bytes,80,opt,name=Data,proto3" json:"Data,omitempty"`
}

func (x *CoinsInfo) Reset() {
	*x = CoinsInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_build_chain_buildchain_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CoinsInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CoinsInfo) ProtoMessage() {}

func (x *CoinsInfo) ProtoReflect() protoreflect.Message {
	mi := &file_npool_build_chain_buildchain_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CoinsInfo.ProtoReflect.Descriptor instead.
func (*CoinsInfo) Descriptor() ([]byte, []int) {
	return file_npool_build_chain_buildchain_proto_rawDescGZIP(), []int{3}
}

func (x *CoinsInfo) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *CoinsInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CoinsInfo) GetChainType() string {
	if x != nil {
		return x.ChainType
	}
	return ""
}

func (x *CoinsInfo) GetTokenType() string {
	if x != nil {
		return x.TokenType
	}
	return ""
}

func (x *CoinsInfo) GetContract() string {
	if x != nil {
		return x.Contract
	}
	return ""
}

func (x *CoinsInfo) GetSimilarity() int32 {
	if x != nil {
		return x.Similarity
	}
	return 0
}

func (x *CoinsInfo) GetRemark() string {
	if x != nil {
		return x.Remark
	}
	return ""
}

func (x *CoinsInfo) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type DeployedCoinInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID        string `protobuf:"bytes,10,opt,name=ID,proto3" json:"ID,omitempty"`
	CoinID    string `protobuf:"bytes,20,opt,name=CoinID,proto3" json:"CoinID,omitempty"`
	Contract  string `protobuf:"bytes,30,opt,name=Contract,proto3" json:"Contract,omitempty"`
	Name      string `protobuf:"bytes,40,opt,name=Name,proto3" json:"Name,omitempty"`
	ChainType string `protobuf:"bytes,50,opt,name=ChainType,proto3" json:"ChainType,omitempty"`
	TokenType string `protobuf:"bytes,60,opt,name=TokenType,proto3" json:"TokenType,omitempty"`
}

func (x *DeployedCoinInfo) Reset() {
	*x = DeployedCoinInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_build_chain_buildchain_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeployedCoinInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeployedCoinInfo) ProtoMessage() {}

func (x *DeployedCoinInfo) ProtoReflect() protoreflect.Message {
	mi := &file_npool_build_chain_buildchain_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeployedCoinInfo.ProtoReflect.Descriptor instead.
func (*DeployedCoinInfo) Descriptor() ([]byte, []int) {
	return file_npool_build_chain_buildchain_proto_rawDescGZIP(), []int{4}
}

func (x *DeployedCoinInfo) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *DeployedCoinInfo) GetCoinID() string {
	if x != nil {
		return x.CoinID
	}
	return ""
}

func (x *DeployedCoinInfo) GetContract() string {
	if x != nil {
		return x.Contract
	}
	return ""
}

func (x *DeployedCoinInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *DeployedCoinInfo) GetChainType() string {
	if x != nil {
		return x.ChainType
	}
	return ""
}

func (x *DeployedCoinInfo) GetTokenType() string {
	if x != nil {
		return x.TokenType
	}
	return ""
}

type GetDeployedCoinsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Infos []*DeployedCoinInfo `protobuf:"bytes,10,rep,name=Infos,proto3" json:"Infos,omitempty"`
	Total uint32              `protobuf:"varint,20,opt,name=Total,proto3" json:"Total,omitempty"`
}

func (x *GetDeployedCoinsResponse) Reset() {
	*x = GetDeployedCoinsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_build_chain_buildchain_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDeployedCoinsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDeployedCoinsResponse) ProtoMessage() {}

func (x *GetDeployedCoinsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_build_chain_buildchain_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDeployedCoinsResponse.ProtoReflect.Descriptor instead.
func (*GetDeployedCoinsResponse) Descriptor() ([]byte, []int) {
	return file_npool_build_chain_buildchain_proto_rawDescGZIP(), []int{5}
}

func (x *GetDeployedCoinsResponse) GetInfos() []*DeployedCoinInfo {
	if x != nil {
		return x.Infos
	}
	return nil
}

func (x *GetDeployedCoinsResponse) GetTotal() uint32 {
	if x != nil {
		return x.Total
	}
	return 0
}

type FaucetRequst struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID       string  `protobuf:"bytes,10,opt,name=ID,proto3" json:"ID,omitempty"`
	Amount   float64 `protobuf:"fixed64,20,opt,name=Amount,proto3" json:"Amount,omitempty"`
	Contract string  `protobuf:"bytes,30,opt,name=Contract,proto3" json:"Contract,omitempty"`
}

func (x *FaucetRequst) Reset() {
	*x = FaucetRequst{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_build_chain_buildchain_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FaucetRequst) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FaucetRequst) ProtoMessage() {}

func (x *FaucetRequst) ProtoReflect() protoreflect.Message {
	mi := &file_npool_build_chain_buildchain_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FaucetRequst.ProtoReflect.Descriptor instead.
func (*FaucetRequst) Descriptor() ([]byte, []int) {
	return file_npool_build_chain_buildchain_proto_rawDescGZIP(), []int{6}
}

func (x *FaucetRequst) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *FaucetRequst) GetAmount() float64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *FaucetRequst) GetContract() string {
	if x != nil {
		return x.Contract
	}
	return ""
}

type FaucetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool   `protobuf:"varint,10,opt,name=Success,proto3" json:"Success,omitempty"`
	Msg     string `protobuf:"bytes,20,opt,name=Msg,proto3" json:"Msg,omitempty"`
}

func (x *FaucetResponse) Reset() {
	*x = FaucetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_build_chain_buildchain_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FaucetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FaucetResponse) ProtoMessage() {}

func (x *FaucetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_build_chain_buildchain_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FaucetResponse.ProtoReflect.Descriptor instead.
func (*FaucetResponse) Descriptor() ([]byte, []int) {
	return file_npool_build_chain_buildchain_proto_rawDescGZIP(), []int{7}
}

func (x *FaucetResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *FaucetResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

var File_npool_build_chain_buildchain_proto protoreflect.FileDescriptor

var file_npool_build_chain_buildchain_proto_rawDesc = []byte{
	0x0a, 0x22, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2d, 0x63, 0x68,
	0x61, 0x69, 0x6e, 0x2f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2e, 0x63, 0x68, 0x61, 0x69,
	0x6e, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x25, 0x0a, 0x0f, 0x56,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x64, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x49, 0x6e,
	0x66, 0x6f, 0x22, 0x52, 0x0a, 0x0c, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x64, 0x43, 0x6f,
	0x69, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x43, 0x6f, 0x69, 0x6e, 0x49, 0x44, 0x18, 0x14, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x43, 0x6f, 0x69, 0x6e, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08, 0x43, 0x6f,
	0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x43, 0x6f,
	0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x22, 0xd3, 0x01, 0x0a, 0x09, 0x43, 0x6f, 0x69, 0x6e, 0x73,
	0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x14, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x43, 0x68, 0x61, 0x69,
	0x6e, 0x54, 0x79, 0x70, 0x65, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x43, 0x68, 0x61,
	0x69, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x54,
	0x79, 0x70, 0x65, 0x18, 0x28, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74,
	0x18, 0x32, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74,
	0x12, 0x1e, 0x0a, 0x0a, 0x53, 0x69, 0x6d, 0x69, 0x6c, 0x61, 0x72, 0x69, 0x74, 0x79, 0x18, 0x3c,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x53, 0x69, 0x6d, 0x69, 0x6c, 0x61, 0x72, 0x69, 0x74, 0x79,
	0x12, 0x16, 0x0a, 0x06, 0x52, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x18, 0x46, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x52, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x12, 0x12, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x61,
	0x18, 0x50, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x44, 0x61, 0x74, 0x61, 0x22, 0xa6, 0x01, 0x0a,
	0x10, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x64, 0x43, 0x6f, 0x69, 0x6e, 0x49, 0x6e, 0x66,
	0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49,
	0x44, 0x12, 0x16, 0x0a, 0x06, 0x43, 0x6f, 0x69, 0x6e, 0x49, 0x44, 0x18, 0x14, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x43, 0x6f, 0x69, 0x6e, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08, 0x43, 0x6f, 0x6e,
	0x74, 0x72, 0x61, 0x63, 0x74, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x43, 0x6f, 0x6e,
	0x74, 0x72, 0x61, 0x63, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x28, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x43, 0x68, 0x61,
	0x69, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x18, 0x32, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x43, 0x68,
	0x61, 0x69, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x54, 0x79, 0x70, 0x65, 0x18, 0x3c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x54, 0x79, 0x70, 0x65, 0x22, 0x68, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x44, 0x65, 0x70, 0x6c,
	0x6f, 0x79, 0x65, 0x64, 0x43, 0x6f, 0x69, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x36, 0x0a, 0x05, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x20, 0x2e, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2e, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x76,
	0x31, 0x2e, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x64, 0x43, 0x6f, 0x69, 0x6e, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x05, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x6f, 0x74,
	0x61, 0x6c, 0x18, 0x14, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x22,
	0x52, 0x0a, 0x0c, 0x46, 0x61, 0x75, 0x63, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x73, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x12,
	0x16, 0x0a, 0x06, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x14, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x06, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x43, 0x6f, 0x6e, 0x74, 0x72,
	0x61, 0x63, 0x74, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x43, 0x6f, 0x6e, 0x74, 0x72,
	0x61, 0x63, 0x74, 0x22, 0x3c, 0x0a, 0x0e, 0x46, 0x61, 0x75, 0x63, 0x65, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12,
	0x10, 0x0a, 0x03, 0x4d, 0x73, 0x67, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x4d, 0x73,
	0x67, 0x32, 0xbb, 0x02, 0x0a, 0x0a, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x43, 0x68, 0x61, 0x69, 0x6e,
	0x12, 0x56, 0x0a, 0x07, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x15, 0x2e, 0x62, 0x75,
	0x69, 0x6c, 0x64, 0x2e, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x1a, 0x1f, 0x2e, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2e, 0x63, 0x68, 0x61, 0x69, 0x6e,
	0x2e, 0x76, 0x31, 0x2e, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x13, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0d, 0x22, 0x08, 0x2f, 0x76, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x3a, 0x01, 0x2a, 0x12, 0x76, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x44,
	0x65, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x64, 0x43, 0x6f, 0x69, 0x6e, 0x73, 0x12, 0x15, 0x2e, 0x62,
	0x75, 0x69, 0x6c, 0x64, 0x2e, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x1a, 0x28, 0x2e, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2e, 0x63, 0x68, 0x61, 0x69,
	0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x64,
	0x43, 0x6f, 0x69, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x21, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x1b, 0x22, 0x16, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x65, 0x74, 0x2f, 0x64,
	0x65, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x64, 0x2f, 0x63, 0x6f, 0x69, 0x6e, 0x73, 0x3a, 0x01, 0x2a,
	0x12, 0x5d, 0x0a, 0x06, 0x46, 0x61, 0x75, 0x63, 0x65, 0x74, 0x12, 0x1c, 0x2e, 0x62, 0x75, 0x69,
	0x6c, 0x64, 0x2e, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x61, 0x75, 0x63,
	0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x62, 0x75, 0x69, 0x6c, 0x64,
	0x2e, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x61, 0x75, 0x63, 0x65, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x15, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0f,
	0x22, 0x0a, 0x2f, 0x76, 0x31, 0x2f, 0x66, 0x61, 0x75, 0x63, 0x65, 0x74, 0x3a, 0x01, 0x2a, 0x42,
	0x24, 0x5a, 0x22, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4e, 0x70,
	0x6f, 0x6f, 0x6c, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x70, 0x6b, 0x67, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_npool_build_chain_buildchain_proto_rawDescOnce sync.Once
	file_npool_build_chain_buildchain_proto_rawDescData = file_npool_build_chain_buildchain_proto_rawDesc
)

func file_npool_build_chain_buildchain_proto_rawDescGZIP() []byte {
	file_npool_build_chain_buildchain_proto_rawDescOnce.Do(func() {
		file_npool_build_chain_buildchain_proto_rawDescData = protoimpl.X.CompressGZIP(file_npool_build_chain_buildchain_proto_rawDescData)
	})
	return file_npool_build_chain_buildchain_proto_rawDescData
}

var file_npool_build_chain_buildchain_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_npool_build_chain_buildchain_proto_goTypes = []interface{}{
	(*Empty)(nil),                    // 0: build.chain.v1.Empty
	(*VersionResponse)(nil),          // 1: build.chain.v1.VersionResponse
	(*DeployedCoin)(nil),             // 2: build.chain.v1.DeployedCoin
	(*CoinsInfo)(nil),                // 3: build.chain.v1.CoinsInfo
	(*DeployedCoinInfo)(nil),         // 4: build.chain.v1.DeployedCoinInfo
	(*GetDeployedCoinsResponse)(nil), // 5: build.chain.v1.GetDeployedCoinsResponse
	(*FaucetRequst)(nil),             // 6: build.chain.v1.FaucetRequst
	(*FaucetResponse)(nil),           // 7: build.chain.v1.FaucetResponse
}
var file_npool_build_chain_buildchain_proto_depIdxs = []int32{
	4, // 0: build.chain.v1.GetDeployedCoinsResponse.Infos:type_name -> build.chain.v1.DeployedCoinInfo
	0, // 1: build.chain.v1.BuildChain.Version:input_type -> build.chain.v1.Empty
	0, // 2: build.chain.v1.BuildChain.GetDeployedCoins:input_type -> build.chain.v1.Empty
	6, // 3: build.chain.v1.BuildChain.Faucet:input_type -> build.chain.v1.FaucetRequst
	1, // 4: build.chain.v1.BuildChain.Version:output_type -> build.chain.v1.VersionResponse
	5, // 5: build.chain.v1.BuildChain.GetDeployedCoins:output_type -> build.chain.v1.GetDeployedCoinsResponse
	7, // 6: build.chain.v1.BuildChain.Faucet:output_type -> build.chain.v1.FaucetResponse
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_npool_build_chain_buildchain_proto_init() }
func file_npool_build_chain_buildchain_proto_init() {
	if File_npool_build_chain_buildchain_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_npool_build_chain_buildchain_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
		file_npool_build_chain_buildchain_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VersionResponse); i {
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
		file_npool_build_chain_buildchain_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeployedCoin); i {
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
		file_npool_build_chain_buildchain_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CoinsInfo); i {
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
		file_npool_build_chain_buildchain_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeployedCoinInfo); i {
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
		file_npool_build_chain_buildchain_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDeployedCoinsResponse); i {
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
		file_npool_build_chain_buildchain_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FaucetRequst); i {
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
		file_npool_build_chain_buildchain_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FaucetResponse); i {
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
			RawDescriptor: file_npool_build_chain_buildchain_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_npool_build_chain_buildchain_proto_goTypes,
		DependencyIndexes: file_npool_build_chain_buildchain_proto_depIdxs,
		MessageInfos:      file_npool_build_chain_buildchain_proto_msgTypes,
	}.Build()
	File_npool_build_chain_buildchain_proto = out.File
	file_npool_build_chain_buildchain_proto_rawDesc = nil
	file_npool_build_chain_buildchain_proto_goTypes = nil
	file_npool_build_chain_buildchain_proto_depIdxs = nil
}
