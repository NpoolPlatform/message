// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.18.1
// source: npool/notif/mw/v1/notif/tx/tx.proto

package tx

import (
	tx "github.com/NpoolPlatform/message/npool/notif/mgr/v1/notif/tx"
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

type CreateTxRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *tx.TxReq `protobuf:"bytes,10,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (x *CreateTxRequest) Reset() {
	*x = CreateTxRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_notif_mw_v1_notif_tx_tx_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTxRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTxRequest) ProtoMessage() {}

func (x *CreateTxRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_notif_mw_v1_notif_tx_tx_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTxRequest.ProtoReflect.Descriptor instead.
func (*CreateTxRequest) Descriptor() ([]byte, []int) {
	return file_npool_notif_mw_v1_notif_tx_tx_proto_rawDescGZIP(), []int{0}
}

func (x *CreateTxRequest) GetInfo() *tx.TxReq {
	if x != nil {
		return x.Info
	}
	return nil
}

type CreateTxResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *tx.Tx `protobuf:"bytes,10,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (x *CreateTxResponse) Reset() {
	*x = CreateTxResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_notif_mw_v1_notif_tx_tx_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTxResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTxResponse) ProtoMessage() {}

func (x *CreateTxResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_notif_mw_v1_notif_tx_tx_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTxResponse.ProtoReflect.Descriptor instead.
func (*CreateTxResponse) Descriptor() ([]byte, []int) {
	return file_npool_notif_mw_v1_notif_tx_tx_proto_rawDescGZIP(), []int{1}
}

func (x *CreateTxResponse) GetInfo() *tx.Tx {
	if x != nil {
		return x.Info
	}
	return nil
}

type UpdateTxRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *tx.TxReq `protobuf:"bytes,10,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (x *UpdateTxRequest) Reset() {
	*x = UpdateTxRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_notif_mw_v1_notif_tx_tx_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateTxRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateTxRequest) ProtoMessage() {}

func (x *UpdateTxRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_notif_mw_v1_notif_tx_tx_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateTxRequest.ProtoReflect.Descriptor instead.
func (*UpdateTxRequest) Descriptor() ([]byte, []int) {
	return file_npool_notif_mw_v1_notif_tx_tx_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateTxRequest) GetInfo() *tx.TxReq {
	if x != nil {
		return x.Info
	}
	return nil
}

type UpdateTxResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *tx.Tx `protobuf:"bytes,10,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (x *UpdateTxResponse) Reset() {
	*x = UpdateTxResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_notif_mw_v1_notif_tx_tx_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateTxResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateTxResponse) ProtoMessage() {}

func (x *UpdateTxResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_notif_mw_v1_notif_tx_tx_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateTxResponse.ProtoReflect.Descriptor instead.
func (*UpdateTxResponse) Descriptor() ([]byte, []int) {
	return file_npool_notif_mw_v1_notif_tx_tx_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateTxResponse) GetInfo() *tx.Tx {
	if x != nil {
		return x.Info
	}
	return nil
}

type GetTxsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Conds  *tx.Conds `protobuf:"bytes,10,opt,name=Conds,proto3" json:"Conds,omitempty"`
	Offset int32     `protobuf:"varint,20,opt,name=Offset,proto3" json:"Offset,omitempty"`
	Limit  int32     `protobuf:"varint,30,opt,name=Limit,proto3" json:"Limit,omitempty"`
}

func (x *GetTxsRequest) Reset() {
	*x = GetTxsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_notif_mw_v1_notif_tx_tx_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTxsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTxsRequest) ProtoMessage() {}

func (x *GetTxsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_notif_mw_v1_notif_tx_tx_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTxsRequest.ProtoReflect.Descriptor instead.
func (*GetTxsRequest) Descriptor() ([]byte, []int) {
	return file_npool_notif_mw_v1_notif_tx_tx_proto_rawDescGZIP(), []int{4}
}

func (x *GetTxsRequest) GetConds() *tx.Conds {
	if x != nil {
		return x.Conds
	}
	return nil
}

func (x *GetTxsRequest) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *GetTxsRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type GetTxsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Infos []*tx.Tx `protobuf:"bytes,10,rep,name=Infos,proto3" json:"Infos,omitempty"`
	Total uint32   `protobuf:"varint,20,opt,name=Total,proto3" json:"Total,omitempty"`
}

func (x *GetTxsResponse) Reset() {
	*x = GetTxsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_notif_mw_v1_notif_tx_tx_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTxsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTxsResponse) ProtoMessage() {}

func (x *GetTxsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_notif_mw_v1_notif_tx_tx_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTxsResponse.ProtoReflect.Descriptor instead.
func (*GetTxsResponse) Descriptor() ([]byte, []int) {
	return file_npool_notif_mw_v1_notif_tx_tx_proto_rawDescGZIP(), []int{5}
}

func (x *GetTxsResponse) GetInfos() []*tx.Tx {
	if x != nil {
		return x.Infos
	}
	return nil
}

func (x *GetTxsResponse) GetTotal() uint32 {
	if x != nil {
		return x.Total
	}
	return 0
}

type GetTxOnlyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Conds *tx.Conds `protobuf:"bytes,10,opt,name=Conds,proto3" json:"Conds,omitempty"`
}

func (x *GetTxOnlyRequest) Reset() {
	*x = GetTxOnlyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_notif_mw_v1_notif_tx_tx_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTxOnlyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTxOnlyRequest) ProtoMessage() {}

func (x *GetTxOnlyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_notif_mw_v1_notif_tx_tx_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTxOnlyRequest.ProtoReflect.Descriptor instead.
func (*GetTxOnlyRequest) Descriptor() ([]byte, []int) {
	return file_npool_notif_mw_v1_notif_tx_tx_proto_rawDescGZIP(), []int{6}
}

func (x *GetTxOnlyRequest) GetConds() *tx.Conds {
	if x != nil {
		return x.Conds
	}
	return nil
}

type GetTxOnlyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *tx.Tx `protobuf:"bytes,10,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (x *GetTxOnlyResponse) Reset() {
	*x = GetTxOnlyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_notif_mw_v1_notif_tx_tx_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTxOnlyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTxOnlyResponse) ProtoMessage() {}

func (x *GetTxOnlyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_notif_mw_v1_notif_tx_tx_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTxOnlyResponse.ProtoReflect.Descriptor instead.
func (*GetTxOnlyResponse) Descriptor() ([]byte, []int) {
	return file_npool_notif_mw_v1_notif_tx_tx_proto_rawDescGZIP(), []int{7}
}

func (x *GetTxOnlyResponse) GetInfo() *tx.Tx {
	if x != nil {
		return x.Info
	}
	return nil
}

var File_npool_notif_mw_v1_notif_tx_tx_proto protoreflect.FileDescriptor

var file_npool_notif_mw_v1_notif_tx_tx_proto_rawDesc = []byte{
	0x0a, 0x23, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x2f, 0x6d, 0x77,
	0x2f, 0x76, 0x31, 0x2f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x2f, 0x74, 0x78, 0x2f, 0x74, 0x78, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x2e, 0x6d, 0x69, 0x64,
	0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x74, 0x78, 0x2e, 0x76, 0x31, 0x1a, 0x24, 0x6e,
	0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x2f, 0x6d, 0x67, 0x72, 0x2f, 0x76,
	0x31, 0x2f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x2f, 0x74, 0x78, 0x2f, 0x74, 0x78, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x47, 0x0a, 0x0f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x78, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x34, 0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x2e, 0x6d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x72, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x2e, 0x74, 0x78, 0x2e, 0x76, 0x31,
	0x2e, 0x54, 0x78, 0x52, 0x65, 0x71, 0x52, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x45, 0x0a, 0x10,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x78, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x31, 0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d,
	0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x6e,
	0x6f, 0x74, 0x69, 0x66, 0x2e, 0x74, 0x78, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x78, 0x52, 0x04, 0x49,
	0x6e, 0x66, 0x6f, 0x22, 0x47, 0x0a, 0x0f, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x78, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x34, 0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x2e, 0x6d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x72, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x2e, 0x74, 0x78, 0x2e, 0x76, 0x31,
	0x2e, 0x54, 0x78, 0x52, 0x65, 0x71, 0x52, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x45, 0x0a, 0x10,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x78, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x31, 0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d,
	0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x6e,
	0x6f, 0x74, 0x69, 0x66, 0x2e, 0x74, 0x78, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x78, 0x52, 0x04, 0x49,
	0x6e, 0x66, 0x6f, 0x22, 0x75, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x54, 0x78, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x36, 0x0a, 0x05, 0x43, 0x6f, 0x6e, 0x64, 0x73, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x2e, 0x6d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x72, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x2e, 0x74, 0x78, 0x2e, 0x76, 0x31, 0x2e,
	0x43, 0x6f, 0x6e, 0x64, 0x73, 0x52, 0x05, 0x43, 0x6f, 0x6e, 0x64, 0x73, 0x12, 0x16, 0x0a, 0x06,
	0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x14, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x4f, 0x66,
	0x66, 0x73, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x1e, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x05, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x22, 0x5b, 0x0a, 0x0e, 0x47, 0x65,
	0x74, 0x54, 0x78, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x33, 0x0a, 0x05,
	0x49, 0x6e, 0x66, 0x6f, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x6e, 0x6f,
	0x74, 0x69, 0x66, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x6e, 0x6f, 0x74, 0x69,
	0x66, 0x2e, 0x74, 0x78, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x78, 0x52, 0x05, 0x49, 0x6e, 0x66, 0x6f,
	0x73, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x14, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x22, 0x4a, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x54, 0x78,
	0x4f, 0x6e, 0x6c, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x36, 0x0a, 0x05, 0x43,
	0x6f, 0x6e, 0x64, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x6e, 0x6f, 0x74,
	0x69, 0x66, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66,
	0x2e, 0x74, 0x78, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x64, 0x73, 0x52, 0x05, 0x43, 0x6f,
	0x6e, 0x64, 0x73, 0x22, 0x46, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x54, 0x78, 0x4f, 0x6e, 0x6c, 0x79,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x31, 0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x2e, 0x6d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x2e, 0x74, 0x78, 0x2e,
	0x76, 0x31, 0x2e, 0x54, 0x78, 0x52, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x32, 0x8d, 0x03, 0x0a, 0x0a,
	0x4d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x12, 0x5f, 0x0a, 0x08, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x54, 0x78, 0x12, 0x27, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x2e, 0x6d,
	0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x74, 0x78, 0x2e, 0x76, 0x31, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x78, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x28, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61,
	0x72, 0x65, 0x2e, 0x74, 0x78, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54,
	0x78, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x5f, 0x0a, 0x08, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x78, 0x12, 0x27, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x2e,
	0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x74, 0x78, 0x2e, 0x76, 0x31,
	0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x78, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x28, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77,
	0x61, 0x72, 0x65, 0x2e, 0x74, 0x78, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x54, 0x78, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x59, 0x0a, 0x06,
	0x47, 0x65, 0x74, 0x54, 0x78, 0x73, 0x12, 0x25, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x2e, 0x6d,
	0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x74, 0x78, 0x2e, 0x76, 0x31, 0x2e,
	0x47, 0x65, 0x74, 0x54, 0x78, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e,
	0x6e, 0x6f, 0x74, 0x69, 0x66, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65,
	0x2e, 0x74, 0x78, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x78, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x62, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x54, 0x78,
	0x4f, 0x6e, 0x6c, 0x79, 0x12, 0x28, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x2e, 0x6d, 0x69, 0x64,
	0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x74, 0x78, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65,
	0x74, 0x54, 0x78, 0x4f, 0x6e, 0x6c, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29,
	0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72,
	0x65, 0x2e, 0x74, 0x78, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x78, 0x4f, 0x6e, 0x6c,
	0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x3d, 0x5a, 0x3b, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4e, 0x70, 0x6f, 0x6f, 0x6c, 0x50,
	0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f,
	0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x2f, 0x6d, 0x77, 0x2f, 0x76,
	0x31, 0x2f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x2f, 0x74, 0x78, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_npool_notif_mw_v1_notif_tx_tx_proto_rawDescOnce sync.Once
	file_npool_notif_mw_v1_notif_tx_tx_proto_rawDescData = file_npool_notif_mw_v1_notif_tx_tx_proto_rawDesc
)

func file_npool_notif_mw_v1_notif_tx_tx_proto_rawDescGZIP() []byte {
	file_npool_notif_mw_v1_notif_tx_tx_proto_rawDescOnce.Do(func() {
		file_npool_notif_mw_v1_notif_tx_tx_proto_rawDescData = protoimpl.X.CompressGZIP(file_npool_notif_mw_v1_notif_tx_tx_proto_rawDescData)
	})
	return file_npool_notif_mw_v1_notif_tx_tx_proto_rawDescData
}

var file_npool_notif_mw_v1_notif_tx_tx_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_npool_notif_mw_v1_notif_tx_tx_proto_goTypes = []interface{}{
	(*CreateTxRequest)(nil),   // 0: notif.middleware.tx.v1.CreateTxRequest
	(*CreateTxResponse)(nil),  // 1: notif.middleware.tx.v1.CreateTxResponse
	(*UpdateTxRequest)(nil),   // 2: notif.middleware.tx.v1.UpdateTxRequest
	(*UpdateTxResponse)(nil),  // 3: notif.middleware.tx.v1.UpdateTxResponse
	(*GetTxsRequest)(nil),     // 4: notif.middleware.tx.v1.GetTxsRequest
	(*GetTxsResponse)(nil),    // 5: notif.middleware.tx.v1.GetTxsResponse
	(*GetTxOnlyRequest)(nil),  // 6: notif.middleware.tx.v1.GetTxOnlyRequest
	(*GetTxOnlyResponse)(nil), // 7: notif.middleware.tx.v1.GetTxOnlyResponse
	(*tx.TxReq)(nil),          // 8: notif.manager.notif.tx.v1.TxReq
	(*tx.Tx)(nil),             // 9: notif.manager.notif.tx.v1.Tx
	(*tx.Conds)(nil),          // 10: notif.manager.notif.tx.v1.Conds
}
var file_npool_notif_mw_v1_notif_tx_tx_proto_depIdxs = []int32{
	8,  // 0: notif.middleware.tx.v1.CreateTxRequest.Info:type_name -> notif.manager.notif.tx.v1.TxReq
	9,  // 1: notif.middleware.tx.v1.CreateTxResponse.Info:type_name -> notif.manager.notif.tx.v1.Tx
	8,  // 2: notif.middleware.tx.v1.UpdateTxRequest.Info:type_name -> notif.manager.notif.tx.v1.TxReq
	9,  // 3: notif.middleware.tx.v1.UpdateTxResponse.Info:type_name -> notif.manager.notif.tx.v1.Tx
	10, // 4: notif.middleware.tx.v1.GetTxsRequest.Conds:type_name -> notif.manager.notif.tx.v1.Conds
	9,  // 5: notif.middleware.tx.v1.GetTxsResponse.Infos:type_name -> notif.manager.notif.tx.v1.Tx
	10, // 6: notif.middleware.tx.v1.GetTxOnlyRequest.Conds:type_name -> notif.manager.notif.tx.v1.Conds
	9,  // 7: notif.middleware.tx.v1.GetTxOnlyResponse.Info:type_name -> notif.manager.notif.tx.v1.Tx
	0,  // 8: notif.middleware.tx.v1.Middleware.CreateTx:input_type -> notif.middleware.tx.v1.CreateTxRequest
	2,  // 9: notif.middleware.tx.v1.Middleware.UpdateTx:input_type -> notif.middleware.tx.v1.UpdateTxRequest
	4,  // 10: notif.middleware.tx.v1.Middleware.GetTxs:input_type -> notif.middleware.tx.v1.GetTxsRequest
	6,  // 11: notif.middleware.tx.v1.Middleware.GetTxOnly:input_type -> notif.middleware.tx.v1.GetTxOnlyRequest
	1,  // 12: notif.middleware.tx.v1.Middleware.CreateTx:output_type -> notif.middleware.tx.v1.CreateTxResponse
	3,  // 13: notif.middleware.tx.v1.Middleware.UpdateTx:output_type -> notif.middleware.tx.v1.UpdateTxResponse
	5,  // 14: notif.middleware.tx.v1.Middleware.GetTxs:output_type -> notif.middleware.tx.v1.GetTxsResponse
	7,  // 15: notif.middleware.tx.v1.Middleware.GetTxOnly:output_type -> notif.middleware.tx.v1.GetTxOnlyResponse
	12, // [12:16] is the sub-list for method output_type
	8,  // [8:12] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_npool_notif_mw_v1_notif_tx_tx_proto_init() }
func file_npool_notif_mw_v1_notif_tx_tx_proto_init() {
	if File_npool_notif_mw_v1_notif_tx_tx_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_npool_notif_mw_v1_notif_tx_tx_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateTxRequest); i {
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
		file_npool_notif_mw_v1_notif_tx_tx_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateTxResponse); i {
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
		file_npool_notif_mw_v1_notif_tx_tx_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateTxRequest); i {
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
		file_npool_notif_mw_v1_notif_tx_tx_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateTxResponse); i {
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
		file_npool_notif_mw_v1_notif_tx_tx_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTxsRequest); i {
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
		file_npool_notif_mw_v1_notif_tx_tx_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTxsResponse); i {
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
		file_npool_notif_mw_v1_notif_tx_tx_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTxOnlyRequest); i {
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
		file_npool_notif_mw_v1_notif_tx_tx_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTxOnlyResponse); i {
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
			RawDescriptor: file_npool_notif_mw_v1_notif_tx_tx_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_npool_notif_mw_v1_notif_tx_tx_proto_goTypes,
		DependencyIndexes: file_npool_notif_mw_v1_notif_tx_tx_proto_depIdxs,
		MessageInfos:      file_npool_notif_mw_v1_notif_tx_tx_proto_msgTypes,
	}.Build()
	File_npool_notif_mw_v1_notif_tx_tx_proto = out.File
	file_npool_notif_mw_v1_notif_tx_tx_proto_rawDesc = nil
	file_npool_notif_mw_v1_notif_tx_tx_proto_goTypes = nil
	file_npool_notif_mw_v1_notif_tx_tx_proto_depIdxs = nil
}
