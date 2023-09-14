// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.18.1
// source: npool/build-chain/buildchain.proto

package buildchain

import (
	v1 "github.com/NpoolPlatform/message/npool/basetypes/v1"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type TokenInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID               string `protobuf:"bytes,10,opt,name=ID,proto3" json:"ID,omitempty"`
	Name             string `protobuf:"bytes,20,opt,name=Name,proto3" json:"Name,omitempty"`
	Unit             string `protobuf:"bytes,30,opt,name=Unit,proto3" json:"Unit,omitempty"`
	Decimal          string `protobuf:"bytes,40,opt,name=Decimal,proto3" json:"Decimal,omitempty"`
	ChainType        string `protobuf:"bytes,50,opt,name=ChainType,proto3" json:"ChainType,omitempty"`
	TokenType        string `protobuf:"bytes,60,opt,name=TokenType,proto3" json:"TokenType,omitempty"`
	OfficialContract string `protobuf:"bytes,70,opt,name=OfficialContract,proto3" json:"OfficialContract,omitempty"`
	PrivateContract  string `protobuf:"bytes,80,opt,name=PrivateContract,proto3" json:"PrivateContract,omitempty"`
	Remark           string `protobuf:"bytes,90,opt,name=Remark,proto3" json:"Remark,omitempty"`
	Data             []byte `protobuf:"bytes,100,opt,name=Data,proto3" json:"Data,omitempty"`
}

func (x *TokenInfo) Reset() {
	*x = TokenInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_build_chain_buildchain_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TokenInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TokenInfo) ProtoMessage() {}

func (x *TokenInfo) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use TokenInfo.ProtoReflect.Descriptor instead.
func (*TokenInfo) Descriptor() ([]byte, []int) {
	return file_npool_build_chain_buildchain_proto_rawDescGZIP(), []int{0}
}

func (x *TokenInfo) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *TokenInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *TokenInfo) GetUnit() string {
	if x != nil {
		return x.Unit
	}
	return ""
}

func (x *TokenInfo) GetDecimal() string {
	if x != nil {
		return x.Decimal
	}
	return ""
}

func (x *TokenInfo) GetChainType() string {
	if x != nil {
		return x.ChainType
	}
	return ""
}

func (x *TokenInfo) GetTokenType() string {
	if x != nil {
		return x.TokenType
	}
	return ""
}

func (x *TokenInfo) GetOfficialContract() string {
	if x != nil {
		return x.OfficialContract
	}
	return ""
}

func (x *TokenInfo) GetPrivateContract() string {
	if x != nil {
		return x.PrivateContract
	}
	return ""
}

func (x *TokenInfo) GetRemark() string {
	if x != nil {
		return x.Remark
	}
	return ""
}

func (x *TokenInfo) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type CreateTokenInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info  *TokenInfo `protobuf:"bytes,10,opt,name=Info,proto3" json:"Info,omitempty"`
	Force bool       `protobuf:"varint,20,opt,name=Force,proto3" json:"Force,omitempty"`
}

func (x *CreateTokenInfoRequest) Reset() {
	*x = CreateTokenInfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_build_chain_buildchain_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTokenInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTokenInfoRequest) ProtoMessage() {}

func (x *CreateTokenInfoRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use CreateTokenInfoRequest.ProtoReflect.Descriptor instead.
func (*CreateTokenInfoRequest) Descriptor() ([]byte, []int) {
	return file_npool_build_chain_buildchain_proto_rawDescGZIP(), []int{1}
}

func (x *CreateTokenInfoRequest) GetInfo() *TokenInfo {
	if x != nil {
		return x.Info
	}
	return nil
}

func (x *CreateTokenInfoRequest) GetForce() bool {
	if x != nil {
		return x.Force
	}
	return false
}

type CreateTokenInfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info    *TokenInfo `protobuf:"bytes,10,opt,name=info,proto3" json:"info,omitempty"`
	Success bool       `protobuf:"varint,20,opt,name=Success,proto3" json:"Success,omitempty"`
	Msg     string     `protobuf:"bytes,30,opt,name=Msg,proto3" json:"Msg,omitempty"`
}

func (x *CreateTokenInfoResponse) Reset() {
	*x = CreateTokenInfoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_build_chain_buildchain_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTokenInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTokenInfoResponse) ProtoMessage() {}

func (x *CreateTokenInfoResponse) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use CreateTokenInfoResponse.ProtoReflect.Descriptor instead.
func (*CreateTokenInfoResponse) Descriptor() ([]byte, []int) {
	return file_npool_build_chain_buildchain_proto_rawDescGZIP(), []int{2}
}

func (x *CreateTokenInfoResponse) GetInfo() *TokenInfo {
	if x != nil {
		return x.Info
	}
	return nil
}

func (x *CreateTokenInfoResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *CreateTokenInfoResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type Conds struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChainType        *v1.StringVal `protobuf:"bytes,10,opt,name=ChainType,proto3" json:"ChainType,omitempty"`
	TokenType        *v1.StringVal `protobuf:"bytes,20,opt,name=TokenType,proto3" json:"TokenType,omitempty"`
	OfficialContract *v1.StringVal `protobuf:"bytes,30,opt,name=OfficialContract,proto3" json:"OfficialContract,omitempty"`
}

func (x *Conds) Reset() {
	*x = Conds{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_build_chain_buildchain_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Conds) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Conds) ProtoMessage() {}

func (x *Conds) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use Conds.ProtoReflect.Descriptor instead.
func (*Conds) Descriptor() ([]byte, []int) {
	return file_npool_build_chain_buildchain_proto_rawDescGZIP(), []int{3}
}

func (x *Conds) GetChainType() *v1.StringVal {
	if x != nil {
		return x.ChainType
	}
	return nil
}

func (x *Conds) GetTokenType() *v1.StringVal {
	if x != nil {
		return x.TokenType
	}
	return nil
}

func (x *Conds) GetOfficialContract() *v1.StringVal {
	if x != nil {
		return x.OfficialContract
	}
	return nil
}

type GetTokenInfosRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Conds *Conds `protobuf:"bytes,10,opt,name=Conds,proto3" json:"Conds,omitempty"`
}

func (x *GetTokenInfosRequest) Reset() {
	*x = GetTokenInfosRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_build_chain_buildchain_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTokenInfosRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTokenInfosRequest) ProtoMessage() {}

func (x *GetTokenInfosRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use GetTokenInfosRequest.ProtoReflect.Descriptor instead.
func (*GetTokenInfosRequest) Descriptor() ([]byte, []int) {
	return file_npool_build_chain_buildchain_proto_rawDescGZIP(), []int{4}
}

func (x *GetTokenInfosRequest) GetConds() *Conds {
	if x != nil {
		return x.Conds
	}
	return nil
}

type GetTokenInfosResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Infos []*TokenInfo `protobuf:"bytes,10,rep,name=Infos,proto3" json:"Infos,omitempty"`
	Total uint32       `protobuf:"varint,20,opt,name=Total,proto3" json:"Total,omitempty"`
}

func (x *GetTokenInfosResponse) Reset() {
	*x = GetTokenInfosResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_build_chain_buildchain_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTokenInfosResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTokenInfosResponse) ProtoMessage() {}

func (x *GetTokenInfosResponse) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use GetTokenInfosResponse.ProtoReflect.Descriptor instead.
func (*GetTokenInfosResponse) Descriptor() ([]byte, []int) {
	return file_npool_build_chain_buildchain_proto_rawDescGZIP(), []int{5}
}

func (x *GetTokenInfosResponse) GetInfos() []*TokenInfo {
	if x != nil {
		return x.Infos
	}
	return nil
}

func (x *GetTokenInfosResponse) GetTotal() uint32 {
	if x != nil {
		return x.Total
	}
	return 0
}

type FaucetRequst struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	To               string `protobuf:"bytes,10,opt,name=To,proto3" json:"To,omitempty"`
	Amount           string `protobuf:"bytes,20,opt,name=Amount,proto3" json:"Amount,omitempty"`
	OfficialContract string `protobuf:"bytes,30,opt,name=OfficialContract,proto3" json:"OfficialContract,omitempty"`
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

func (x *FaucetRequst) GetTo() string {
	if x != nil {
		return x.To
	}
	return ""
}

func (x *FaucetRequst) GetAmount() string {
	if x != nil {
		return x.Amount
	}
	return ""
}

func (x *FaucetRequst) GetOfficialContract() string {
	if x != nil {
		return x.OfficialContract
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
	0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x20, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x2f, 0x76, 0x31, 0x2f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1e, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70,
	0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x9b, 0x02, 0x0a, 0x09, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x12,
	0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x12,
	0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x55, 0x6e, 0x69, 0x74, 0x18, 0x1e, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x55, 0x6e, 0x69, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x44, 0x65, 0x63, 0x69, 0x6d,
	0x61, 0x6c, 0x18, 0x28, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x44, 0x65, 0x63, 0x69, 0x6d, 0x61,
	0x6c, 0x12, 0x1c, 0x0a, 0x09, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x18, 0x32,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x1c, 0x0a, 0x09, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x18, 0x3c, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x2a, 0x0a,
	0x10, 0x4f, 0x66, 0x66, 0x69, 0x63, 0x69, 0x61, 0x6c, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63,
	0x74, 0x18, 0x46, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x4f, 0x66, 0x66, 0x69, 0x63, 0x69, 0x61,
	0x6c, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x12, 0x28, 0x0a, 0x0f, 0x50, 0x72, 0x69,
	0x76, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x18, 0x50, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0f, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x72,
	0x61, 0x63, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x52, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x18, 0x5a, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x52, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x12, 0x12, 0x0a, 0x04, 0x44,
	0x61, 0x74, 0x61, 0x18, 0x64, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x44, 0x61, 0x74, 0x61, 0x22,
	0x5d, 0x0a, 0x16, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2d, 0x0a, 0x04, 0x49, 0x6e, 0x66,
	0x6f, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2e,
	0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x14, 0x0a, 0x05, 0x46, 0x6f, 0x72, 0x63,
	0x65, 0x18, 0x14, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x46, 0x6f, 0x72, 0x63, 0x65, 0x22, 0x74,
	0x0a, 0x17, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2d, 0x0a, 0x04, 0x69, 0x6e, 0x66,
	0x6f, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2e,
	0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x12, 0x18, 0x0a, 0x07, 0x53, 0x75, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x18, 0x14, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x53, 0x75, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x4d, 0x73, 0x67, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x4d, 0x73, 0x67, 0x22, 0xba, 0x01, 0x0a, 0x05, 0x43, 0x6f, 0x6e, 0x64, 0x73, 0x12, 0x35,
	0x0a, 0x09, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x17, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x52, 0x09, 0x43, 0x68, 0x61, 0x69,
	0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x35, 0x0a, 0x09, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x54, 0x79,
	0x70, 0x65, 0x18, 0x14, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61,
	0x6c, 0x52, 0x09, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x43, 0x0a, 0x10,
	0x4f, 0x66, 0x66, 0x69, 0x63, 0x69, 0x61, 0x6c, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74,
	0x18, 0x1e, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70,
	0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x52,
	0x10, 0x4f, 0x66, 0x66, 0x69, 0x63, 0x69, 0x61, 0x6c, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63,
	0x74, 0x22, 0x43, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x49, 0x6e, 0x66,
	0x6f, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2b, 0x0a, 0x05, 0x43, 0x6f, 0x6e,
	0x64, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x62, 0x75, 0x69, 0x6c, 0x64,
	0x2e, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x64, 0x73, 0x52,
	0x05, 0x43, 0x6f, 0x6e, 0x64, 0x73, 0x22, 0x5e, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x2f, 0x0a, 0x05, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19,
	0x2e, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2e, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x05, 0x49, 0x6e, 0x66, 0x6f, 0x73,
	0x12, 0x14, 0x0a, 0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x14, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x22, 0x62, 0x0a, 0x0c, 0x46, 0x61, 0x75, 0x63, 0x65, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x54, 0x6f, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x54, 0x6f, 0x12, 0x16, 0x0a, 0x06, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x2a,
	0x0a, 0x10, 0x4f, 0x66, 0x66, 0x69, 0x63, 0x69, 0x61, 0x6c, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61,
	0x63, 0x74, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x4f, 0x66, 0x66, 0x69, 0x63, 0x69,
	0x61, 0x6c, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x22, 0x3c, 0x0a, 0x0e, 0x46, 0x61,
	0x75, 0x63, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x53,
	0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x4d, 0x73, 0x67, 0x18, 0x14, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x4d, 0x73, 0x67, 0x32, 0xc8, 0x03, 0x0a, 0x0a, 0x42, 0x75, 0x69,
	0x6c, 0x64, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x12, 0x58, 0x0a, 0x07, 0x56, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1d, 0x2e, 0x62, 0x61, 0x73,
	0x65, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x16, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x10, 0x3a, 0x01, 0x2a, 0x22, 0x0b, 0x2f, 0x76, 0x31, 0x2f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x12, 0x7b, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x49, 0x6e, 0x66,
	0x6f, 0x73, 0x12, 0x24, 0x2e, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2e, 0x63, 0x68, 0x61, 0x69, 0x6e,
	0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x49, 0x6e, 0x66, 0x6f,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x62, 0x75, 0x69, 0x6c, 0x64,
	0x2e, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x1d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x17, 0x3a, 0x01, 0x2a, 0x22, 0x12, 0x2f, 0x76, 0x31, 0x2f,
	0x67, 0x65, 0x74, 0x2f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x69, 0x6e, 0x66, 0x6f, 0x73, 0x12, 0x83,
	0x01, 0x0a, 0x0f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x49, 0x6e,
	0x66, 0x6f, 0x12, 0x26, 0x2e, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2e, 0x63, 0x68, 0x61, 0x69, 0x6e,
	0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x62, 0x75, 0x69,
	0x6c, 0x64, 0x2e, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x1f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x19, 0x3a, 0x01, 0x2a, 0x22, 0x14,
	0x2f, 0x76, 0x31, 0x2f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x2f, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x69, 0x6e, 0x66, 0x6f, 0x12, 0x5d, 0x0a, 0x06, 0x46, 0x61, 0x75, 0x63, 0x65, 0x74, 0x12, 0x1c,
	0x2e, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2e, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e,
	0x46, 0x61, 0x75, 0x63, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x62,
	0x75, 0x69, 0x6c, 0x64, 0x2e, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x61,
	0x75, 0x63, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x15, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x0f, 0x3a, 0x01, 0x2a, 0x22, 0x0a, 0x2f, 0x76, 0x31, 0x2f, 0x66, 0x61, 0x75,
	0x63, 0x65, 0x74, 0x42, 0x29, 0x5a, 0x27, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x4e, 0x70, 0x6f, 0x6f, 0x6c, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f,
	0x70, 0x6b, 0x67, 0x2f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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
	(*TokenInfo)(nil),               // 0: build.chain.v1.TokenInfo
	(*CreateTokenInfoRequest)(nil),  // 1: build.chain.v1.CreateTokenInfoRequest
	(*CreateTokenInfoResponse)(nil), // 2: build.chain.v1.CreateTokenInfoResponse
	(*Conds)(nil),                   // 3: build.chain.v1.Conds
	(*GetTokenInfosRequest)(nil),    // 4: build.chain.v1.GetTokenInfosRequest
	(*GetTokenInfosResponse)(nil),   // 5: build.chain.v1.GetTokenInfosResponse
	(*FaucetRequst)(nil),            // 6: build.chain.v1.FaucetRequst
	(*FaucetResponse)(nil),          // 7: build.chain.v1.FaucetResponse
	(*v1.StringVal)(nil),            // 8: basetypes.v1.StringVal
	(*emptypb.Empty)(nil),           // 9: google.protobuf.Empty
	(*v1.VersionResponse)(nil),      // 10: basetypes.v1.VersionResponse
}
var file_npool_build_chain_buildchain_proto_depIdxs = []int32{
	0,  // 0: build.chain.v1.CreateTokenInfoRequest.Info:type_name -> build.chain.v1.TokenInfo
	0,  // 1: build.chain.v1.CreateTokenInfoResponse.info:type_name -> build.chain.v1.TokenInfo
	8,  // 2: build.chain.v1.Conds.ChainType:type_name -> basetypes.v1.StringVal
	8,  // 3: build.chain.v1.Conds.TokenType:type_name -> basetypes.v1.StringVal
	8,  // 4: build.chain.v1.Conds.OfficialContract:type_name -> basetypes.v1.StringVal
	3,  // 5: build.chain.v1.GetTokenInfosRequest.Conds:type_name -> build.chain.v1.Conds
	0,  // 6: build.chain.v1.GetTokenInfosResponse.Infos:type_name -> build.chain.v1.TokenInfo
	9,  // 7: build.chain.v1.BuildChain.Version:input_type -> google.protobuf.Empty
	4,  // 8: build.chain.v1.BuildChain.GetTokenInfos:input_type -> build.chain.v1.GetTokenInfosRequest
	1,  // 9: build.chain.v1.BuildChain.CreateTokenInfo:input_type -> build.chain.v1.CreateTokenInfoRequest
	6,  // 10: build.chain.v1.BuildChain.Faucet:input_type -> build.chain.v1.FaucetRequst
	10, // 11: build.chain.v1.BuildChain.Version:output_type -> basetypes.v1.VersionResponse
	5,  // 12: build.chain.v1.BuildChain.GetTokenInfos:output_type -> build.chain.v1.GetTokenInfosResponse
	2,  // 13: build.chain.v1.BuildChain.CreateTokenInfo:output_type -> build.chain.v1.CreateTokenInfoResponse
	7,  // 14: build.chain.v1.BuildChain.Faucet:output_type -> build.chain.v1.FaucetResponse
	11, // [11:15] is the sub-list for method output_type
	7,  // [7:11] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_npool_build_chain_buildchain_proto_init() }
func file_npool_build_chain_buildchain_proto_init() {
	if File_npool_build_chain_buildchain_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_npool_build_chain_buildchain_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TokenInfo); i {
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
			switch v := v.(*CreateTokenInfoRequest); i {
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
			switch v := v.(*CreateTokenInfoResponse); i {
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
		file_npool_build_chain_buildchain_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTokenInfosRequest); i {
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
			switch v := v.(*GetTokenInfosResponse); i {
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
