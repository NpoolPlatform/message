// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.18.1
// source: npool/oplog/mw/v1/log/log.proto

package log

import (
	v1 "github.com/NpoolPlatform/message/npool/basetypes/v1"
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

type LogReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AppID         *string        `protobuf:"bytes,10,opt,name=AppID,proto3,oneof" json:"AppID,omitempty"`
	UserID        *string        `protobuf:"bytes,20,opt,name=UserID,proto3,oneof" json:"UserID,omitempty"`
	Path          *string        `protobuf:"bytes,30,opt,name=Path,proto3,oneof" json:"Path,omitempty"`
	Method        *v1.HTTPMethod `protobuf:"varint,40,opt,name=Method,proto3,enum=basetypes.v1.HTTPMethod,oneof" json:"Method,omitempty"`
	Arguments     *string        `protobuf:"bytes,50,opt,name=Arguments,proto3,oneof" json:"Arguments,omitempty"`
	HumanReadable *string        `protobuf:"bytes,60,opt,name=HumanReadable,proto3,oneof" json:"HumanReadable,omitempty"`
	Result        *v1.Result     `protobuf:"varint,70,opt,name=Result,proto3,enum=basetypes.v1.Result,oneof" json:"Result,omitempty"`
	FailReason    *string        `protobuf:"bytes,80,opt,name=FailReason,proto3,oneof" json:"FailReason,omitempty"`
}

func (x *LogReq) Reset() {
	*x = LogReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_oplog_mw_v1_log_log_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogReq) ProtoMessage() {}

func (x *LogReq) ProtoReflect() protoreflect.Message {
	mi := &file_npool_oplog_mw_v1_log_log_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogReq.ProtoReflect.Descriptor instead.
func (*LogReq) Descriptor() ([]byte, []int) {
	return file_npool_oplog_mw_v1_log_log_proto_rawDescGZIP(), []int{0}
}

func (x *LogReq) GetAppID() string {
	if x != nil && x.AppID != nil {
		return *x.AppID
	}
	return ""
}

func (x *LogReq) GetUserID() string {
	if x != nil && x.UserID != nil {
		return *x.UserID
	}
	return ""
}

func (x *LogReq) GetPath() string {
	if x != nil && x.Path != nil {
		return *x.Path
	}
	return ""
}

func (x *LogReq) GetMethod() v1.HTTPMethod {
	if x != nil && x.Method != nil {
		return *x.Method
	}
	return v1.HTTPMethod(0)
}

func (x *LogReq) GetArguments() string {
	if x != nil && x.Arguments != nil {
		return *x.Arguments
	}
	return ""
}

func (x *LogReq) GetHumanReadable() string {
	if x != nil && x.HumanReadable != nil {
		return *x.HumanReadable
	}
	return ""
}

func (x *LogReq) GetResult() v1.Result {
	if x != nil && x.Result != nil {
		return *x.Result
	}
	return v1.Result(0)
}

func (x *LogReq) GetFailReason() string {
	if x != nil && x.FailReason != nil {
		return *x.FailReason
	}
	return ""
}

type Log struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: sql:"id"
	ID uint32 `protobuf:"varint,10,opt,name=ID,proto3" json:"ID,omitempty" sql:"id"`
	// @inject_tag: sql:"app_id"
	AppID string `protobuf:"bytes,20,opt,name=AppID,proto3" json:"AppID,omitempty" sql:"app_id"`
	// @inject_tag: sql:"user_id"
	UserID *string `protobuf:"bytes,30,opt,name=UserID,proto3,oneof" json:"UserID,omitempty" sql:"user_id"`
	// @inject_tag: sql:"path"
	Path string `protobuf:"bytes,40,opt,name=Path,proto3" json:"Path,omitempty" sql:"path"`
	// @inject_tag: sql:"method"
	MethodStr string        `protobuf:"bytes,50,opt,name=MethodStr,proto3" json:"MethodStr,omitempty" sql:"method"`
	Method    v1.HTTPMethod `protobuf:"varint,60,opt,name=Method,proto3,enum=basetypes.v1.HTTPMethod" json:"Method,omitempty"`
	// @inject_tag: sql:"arguments"
	Arguments string `protobuf:"bytes,70,opt,name=Arguments,proto3" json:"Arguments,omitempty" sql:"arguments"`
	// @inject_tag: sql:"old_value"
	OldValue string `protobuf:"bytes,80,opt,name=OldValue,proto3" json:"OldValue,omitempty" sql:"old_value"`
	// @inject_tag: sql:"human_readable"
	HumanReadable string `protobuf:"bytes,90,opt,name=HumanReadable,proto3" json:"HumanReadable,omitempty" sql:"human_readable"`
	// @inject_tag: sql:"result"
	ResultStr string    `protobuf:"bytes,100,opt,name=ResultStr,proto3" json:"ResultStr,omitempty" sql:"result"`
	Result    v1.Result `protobuf:"varint,110,opt,name=Result,proto3,enum=basetypes.v1.Result" json:"Result,omitempty"`
	// @inject_tag: sql:"fail_reason"
	FailReason string `protobuf:"bytes,120,opt,name=FailReason,proto3" json:"FailReason,omitempty" sql:"fail_reason"`
	// @inject_tag: sql:"created_at"
	CreatedAt uint32 `protobuf:"varint,1000,opt,name=CreatedAt,proto3" json:"CreatedAt,omitempty" sql:"created_at"`
	// @inject_tag: sql:"updated_at"
	UpdatedAt uint32 `protobuf:"varint,1010,opt,name=UpdatedAt,proto3" json:"UpdatedAt,omitempty" sql:"updated_at"`
}

func (x *Log) Reset() {
	*x = Log{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_oplog_mw_v1_log_log_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Log) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Log) ProtoMessage() {}

func (x *Log) ProtoReflect() protoreflect.Message {
	mi := &file_npool_oplog_mw_v1_log_log_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Log.ProtoReflect.Descriptor instead.
func (*Log) Descriptor() ([]byte, []int) {
	return file_npool_oplog_mw_v1_log_log_proto_rawDescGZIP(), []int{1}
}

func (x *Log) GetID() uint32 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *Log) GetAppID() string {
	if x != nil {
		return x.AppID
	}
	return ""
}

func (x *Log) GetUserID() string {
	if x != nil && x.UserID != nil {
		return *x.UserID
	}
	return ""
}

func (x *Log) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *Log) GetMethodStr() string {
	if x != nil {
		return x.MethodStr
	}
	return ""
}

func (x *Log) GetMethod() v1.HTTPMethod {
	if x != nil {
		return x.Method
	}
	return v1.HTTPMethod(0)
}

func (x *Log) GetArguments() string {
	if x != nil {
		return x.Arguments
	}
	return ""
}

func (x *Log) GetOldValue() string {
	if x != nil {
		return x.OldValue
	}
	return ""
}

func (x *Log) GetHumanReadable() string {
	if x != nil {
		return x.HumanReadable
	}
	return ""
}

func (x *Log) GetResultStr() string {
	if x != nil {
		return x.ResultStr
	}
	return ""
}

func (x *Log) GetResult() v1.Result {
	if x != nil {
		return x.Result
	}
	return v1.Result(0)
}

func (x *Log) GetFailReason() string {
	if x != nil {
		return x.FailReason
	}
	return ""
}

func (x *Log) GetCreatedAt() uint32 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *Log) GetUpdatedAt() uint32 {
	if x != nil {
		return x.UpdatedAt
	}
	return 0
}

type Conds struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AppID  *v1.StringVal `protobuf:"bytes,10,opt,name=AppID,proto3,oneof" json:"AppID,omitempty"`
	UserID *v1.StringVal `protobuf:"bytes,20,opt,name=UserID,proto3,oneof" json:"UserID,omitempty"`
	Result *v1.Uint32Val `protobuf:"bytes,30,opt,name=Result,proto3,oneof" json:"Result,omitempty"`
}

func (x *Conds) Reset() {
	*x = Conds{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_oplog_mw_v1_log_log_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Conds) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Conds) ProtoMessage() {}

func (x *Conds) ProtoReflect() protoreflect.Message {
	mi := &file_npool_oplog_mw_v1_log_log_proto_msgTypes[2]
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
	return file_npool_oplog_mw_v1_log_log_proto_rawDescGZIP(), []int{2}
}

func (x *Conds) GetAppID() *v1.StringVal {
	if x != nil {
		return x.AppID
	}
	return nil
}

func (x *Conds) GetUserID() *v1.StringVal {
	if x != nil {
		return x.UserID
	}
	return nil
}

func (x *Conds) GetResult() *v1.Uint32Val {
	if x != nil {
		return x.Result
	}
	return nil
}

type CreateLogRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *LogReq `protobuf:"bytes,10,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (x *CreateLogRequest) Reset() {
	*x = CreateLogRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_oplog_mw_v1_log_log_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateLogRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateLogRequest) ProtoMessage() {}

func (x *CreateLogRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_oplog_mw_v1_log_log_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateLogRequest.ProtoReflect.Descriptor instead.
func (*CreateLogRequest) Descriptor() ([]byte, []int) {
	return file_npool_oplog_mw_v1_log_log_proto_rawDescGZIP(), []int{3}
}

func (x *CreateLogRequest) GetInfo() *LogReq {
	if x != nil {
		return x.Info
	}
	return nil
}

type CreateLogResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *Log `protobuf:"bytes,10,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (x *CreateLogResponse) Reset() {
	*x = CreateLogResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_oplog_mw_v1_log_log_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateLogResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateLogResponse) ProtoMessage() {}

func (x *CreateLogResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_oplog_mw_v1_log_log_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateLogResponse.ProtoReflect.Descriptor instead.
func (*CreateLogResponse) Descriptor() ([]byte, []int) {
	return file_npool_oplog_mw_v1_log_log_proto_rawDescGZIP(), []int{4}
}

func (x *CreateLogResponse) GetInfo() *Log {
	if x != nil {
		return x.Info
	}
	return nil
}

type UpdateLogRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *LogReq `protobuf:"bytes,10,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (x *UpdateLogRequest) Reset() {
	*x = UpdateLogRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_oplog_mw_v1_log_log_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateLogRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateLogRequest) ProtoMessage() {}

func (x *UpdateLogRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_oplog_mw_v1_log_log_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateLogRequest.ProtoReflect.Descriptor instead.
func (*UpdateLogRequest) Descriptor() ([]byte, []int) {
	return file_npool_oplog_mw_v1_log_log_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateLogRequest) GetInfo() *LogReq {
	if x != nil {
		return x.Info
	}
	return nil
}

type UpdateLogResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *Log `protobuf:"bytes,10,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (x *UpdateLogResponse) Reset() {
	*x = UpdateLogResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_oplog_mw_v1_log_log_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateLogResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateLogResponse) ProtoMessage() {}

func (x *UpdateLogResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_oplog_mw_v1_log_log_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateLogResponse.ProtoReflect.Descriptor instead.
func (*UpdateLogResponse) Descriptor() ([]byte, []int) {
	return file_npool_oplog_mw_v1_log_log_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateLogResponse) GetInfo() *Log {
	if x != nil {
		return x.Info
	}
	return nil
}

type GetLogsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Conds  *Conds `protobuf:"bytes,10,opt,name=Conds,proto3" json:"Conds,omitempty"`
	Offset int32  `protobuf:"varint,20,opt,name=Offset,proto3" json:"Offset,omitempty"`
	Limit  int32  `protobuf:"varint,30,opt,name=Limit,proto3" json:"Limit,omitempty"`
}

func (x *GetLogsRequest) Reset() {
	*x = GetLogsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_oplog_mw_v1_log_log_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetLogsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLogsRequest) ProtoMessage() {}

func (x *GetLogsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_oplog_mw_v1_log_log_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLogsRequest.ProtoReflect.Descriptor instead.
func (*GetLogsRequest) Descriptor() ([]byte, []int) {
	return file_npool_oplog_mw_v1_log_log_proto_rawDescGZIP(), []int{7}
}

func (x *GetLogsRequest) GetConds() *Conds {
	if x != nil {
		return x.Conds
	}
	return nil
}

func (x *GetLogsRequest) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *GetLogsRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type GetLogsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Infos []*Log `protobuf:"bytes,10,rep,name=Infos,proto3" json:"Infos,omitempty"`
	Total uint32 `protobuf:"varint,20,opt,name=Total,proto3" json:"Total,omitempty"`
}

func (x *GetLogsResponse) Reset() {
	*x = GetLogsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_oplog_mw_v1_log_log_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetLogsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLogsResponse) ProtoMessage() {}

func (x *GetLogsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_oplog_mw_v1_log_log_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLogsResponse.ProtoReflect.Descriptor instead.
func (*GetLogsResponse) Descriptor() ([]byte, []int) {
	return file_npool_oplog_mw_v1_log_log_proto_rawDescGZIP(), []int{8}
}

func (x *GetLogsResponse) GetInfos() []*Log {
	if x != nil {
		return x.Infos
	}
	return nil
}

func (x *GetLogsResponse) GetTotal() uint32 {
	if x != nil {
		return x.Total
	}
	return 0
}

var File_npool_oplog_mw_v1_log_log_proto protoreflect.FileDescriptor

var file_npool_oplog_mw_v1_log_log_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x6f, 0x70, 0x6c, 0x6f, 0x67, 0x2f, 0x6d, 0x77,
	0x2f, 0x76, 0x31, 0x2f, 0x6c, 0x6f, 0x67, 0x2f, 0x6c, 0x6f, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x17, 0x6f, 0x70, 0x6c, 0x6f, 0x67, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77,
	0x61, 0x72, 0x65, 0x2e, 0x6c, 0x6f, 0x67, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x23, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f,
	0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x68, 0x74, 0x74,
	0x70, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x6e,
	0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x76,
	0x31, 0x2f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e,
	0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f,
	0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x99,
	0x03, 0x0a, 0x06, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x71, 0x12, 0x19, 0x0a, 0x05, 0x41, 0x70, 0x70,
	0x49, 0x44, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x05, 0x41, 0x70, 0x70, 0x49,
	0x44, 0x88, 0x01, 0x01, 0x12, 0x1b, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x14,
	0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x88, 0x01,
	0x01, 0x12, 0x17, 0x0a, 0x04, 0x50, 0x61, 0x74, 0x68, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x09, 0x48,
	0x02, 0x52, 0x04, 0x50, 0x61, 0x74, 0x68, 0x88, 0x01, 0x01, 0x12, 0x35, 0x0a, 0x06, 0x4d, 0x65,
	0x74, 0x68, 0x6f, 0x64, 0x18, 0x28, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x18, 0x2e, 0x62, 0x61, 0x73,
	0x65, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x48, 0x54, 0x54, 0x50, 0x4d, 0x65,
	0x74, 0x68, 0x6f, 0x64, 0x48, 0x03, 0x52, 0x06, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x88, 0x01,
	0x01, 0x12, 0x21, 0x0a, 0x09, 0x41, 0x72, 0x67, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x32,
	0x20, 0x01, 0x28, 0x09, 0x48, 0x04, 0x52, 0x09, 0x41, 0x72, 0x67, 0x75, 0x6d, 0x65, 0x6e, 0x74,
	0x73, 0x88, 0x01, 0x01, 0x12, 0x29, 0x0a, 0x0d, 0x48, 0x75, 0x6d, 0x61, 0x6e, 0x52, 0x65, 0x61,
	0x64, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x3c, 0x20, 0x01, 0x28, 0x09, 0x48, 0x05, 0x52, 0x0d, 0x48,
	0x75, 0x6d, 0x61, 0x6e, 0x52, 0x65, 0x61, 0x64, 0x61, 0x62, 0x6c, 0x65, 0x88, 0x01, 0x01, 0x12,
	0x31, 0x0a, 0x06, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x46, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x14, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x52,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x48, 0x06, 0x52, 0x06, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x88,
	0x01, 0x01, 0x12, 0x23, 0x0a, 0x0a, 0x46, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e,
	0x18, 0x50, 0x20, 0x01, 0x28, 0x09, 0x48, 0x07, 0x52, 0x0a, 0x46, 0x61, 0x69, 0x6c, 0x52, 0x65,
	0x61, 0x73, 0x6f, 0x6e, 0x88, 0x01, 0x01, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x41, 0x70, 0x70, 0x49,
	0x44, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x42, 0x07, 0x0a, 0x05,
	0x5f, 0x50, 0x61, 0x74, 0x68, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64,
	0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x41, 0x72, 0x67, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x42, 0x10,
	0x0a, 0x0e, 0x5f, 0x48, 0x75, 0x6d, 0x61, 0x6e, 0x52, 0x65, 0x61, 0x64, 0x61, 0x62, 0x6c, 0x65,
	0x42, 0x09, 0x0a, 0x07, 0x5f, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x42, 0x0d, 0x0a, 0x0b, 0x5f,
	0x46, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x22, 0xc1, 0x03, 0x0a, 0x03, 0x4c,
	0x6f, 0x67, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02,
	0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x41, 0x70, 0x70, 0x49, 0x44, 0x18, 0x14, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x41, 0x70, 0x70, 0x49, 0x44, 0x12, 0x1b, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72,
	0x49, 0x44, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72,
	0x49, 0x44, 0x88, 0x01, 0x01, 0x12, 0x12, 0x0a, 0x04, 0x50, 0x61, 0x74, 0x68, 0x18, 0x28, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x50, 0x61, 0x74, 0x68, 0x12, 0x1c, 0x0a, 0x09, 0x4d, 0x65, 0x74,
	0x68, 0x6f, 0x64, 0x53, 0x74, 0x72, 0x18, 0x32, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x4d, 0x65,
	0x74, 0x68, 0x6f, 0x64, 0x53, 0x74, 0x72, 0x12, 0x30, 0x0a, 0x06, 0x4d, 0x65, 0x74, 0x68, 0x6f,
	0x64, 0x18, 0x3c, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x18, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x48, 0x54, 0x54, 0x50, 0x4d, 0x65, 0x74, 0x68, 0x6f,
	0x64, 0x52, 0x06, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x41, 0x72, 0x67,
	0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x46, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x41, 0x72,
	0x67, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x4f, 0x6c, 0x64, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x50, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x4f, 0x6c, 0x64, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x48, 0x75, 0x6d, 0x61, 0x6e, 0x52, 0x65, 0x61, 0x64,
	0x61, 0x62, 0x6c, 0x65, 0x18, 0x5a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x48, 0x75, 0x6d, 0x61,
	0x6e, 0x52, 0x65, 0x61, 0x64, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x52, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x53, 0x74, 0x72, 0x18, 0x64, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x52, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x53, 0x74, 0x72, 0x12, 0x2c, 0x0a, 0x06, 0x52, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x18, 0x6e, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x14, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x06, 0x52,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x46, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x61,
	0x73, 0x6f, 0x6e, 0x18, 0x78, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x46, 0x61, 0x69, 0x6c, 0x52,
	0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x1d, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x18, 0xe8, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x18, 0xf2, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x22, 0xc7,
	0x01, 0x0a, 0x05, 0x43, 0x6f, 0x6e, 0x64, 0x73, 0x12, 0x32, 0x0a, 0x05, 0x41, 0x70, 0x70, 0x49,
	0x44, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c,
	0x48, 0x00, 0x52, 0x05, 0x41, 0x70, 0x70, 0x49, 0x44, 0x88, 0x01, 0x01, 0x12, 0x34, 0x0a, 0x06,
	0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x14, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x62,
	0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x72, 0x69,
	0x6e, 0x67, 0x56, 0x61, 0x6c, 0x48, 0x01, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x88,
	0x01, 0x01, 0x12, 0x34, 0x0a, 0x06, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x1e, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x17, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x55, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x48, 0x02, 0x52, 0x06, 0x52,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x88, 0x01, 0x01, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x41, 0x70, 0x70,
	0x49, 0x44, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x42, 0x09, 0x0a,
	0x07, 0x5f, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x47, 0x0a, 0x10, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x33, 0x0a, 0x04,
	0x49, 0x6e, 0x66, 0x6f, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x6f, 0x70, 0x6c,
	0x6f, 0x67, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x6c, 0x6f,
	0x67, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x71, 0x52, 0x04, 0x49, 0x6e, 0x66,
	0x6f, 0x22, 0x45, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x6f, 0x67, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x30, 0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x6f, 0x70, 0x6c, 0x6f, 0x67, 0x2e, 0x6d, 0x69, 0x64,
	0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x6c, 0x6f, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x4c,
	0x6f, 0x67, 0x52, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x47, 0x0a, 0x10, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x33, 0x0a, 0x04,
	0x49, 0x6e, 0x66, 0x6f, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x6f, 0x70, 0x6c,
	0x6f, 0x67, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x6c, 0x6f,
	0x67, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x71, 0x52, 0x04, 0x49, 0x6e, 0x66,
	0x6f, 0x22, 0x45, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4c, 0x6f, 0x67, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x30, 0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x6f, 0x70, 0x6c, 0x6f, 0x67, 0x2e, 0x6d, 0x69, 0x64,
	0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x6c, 0x6f, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x4c,
	0x6f, 0x67, 0x52, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x74, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x4c,
	0x6f, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x34, 0x0a, 0x05, 0x43, 0x6f,
	0x6e, 0x64, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x6f, 0x70, 0x6c, 0x6f,
	0x67, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x6c, 0x6f, 0x67,
	0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x64, 0x73, 0x52, 0x05, 0x43, 0x6f, 0x6e, 0x64, 0x73,
	0x12, 0x16, 0x0a, 0x06, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x14, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x06, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x4c, 0x69, 0x6d, 0x69,
	0x74, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x22, 0x5b,
	0x0a, 0x0f, 0x47, 0x65, 0x74, 0x4c, 0x6f, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x32, 0x0a, 0x05, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x1c, 0x2e, 0x6f, 0x70, 0x6c, 0x6f, 0x67, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77,
	0x61, 0x72, 0x65, 0x2e, 0x6c, 0x6f, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x6f, 0x67, 0x52, 0x05,
	0x49, 0x6e, 0x66, 0x6f, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x14,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x32, 0xb8, 0x02, 0x0a, 0x0a,
	0x4d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x12, 0x64, 0x0a, 0x09, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x4c, 0x6f, 0x67, 0x12, 0x29, 0x2e, 0x6f, 0x70, 0x6c, 0x6f, 0x67, 0x2e,
	0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x6c, 0x6f, 0x67, 0x2e, 0x76,
	0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x6f, 0x70, 0x6c, 0x6f, 0x67, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c,
	0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x6c, 0x6f, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x64, 0x0a, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4c, 0x6f, 0x67, 0x12, 0x29, 0x2e,
	0x6f, 0x70, 0x6c, 0x6f, 0x67, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65,
	0x2e, 0x6c, 0x6f, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4c, 0x6f,
	0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x6f, 0x70, 0x6c, 0x6f, 0x67,
	0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x6c, 0x6f, 0x67, 0x2e,
	0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x5e, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x4c, 0x6f, 0x67,
	0x73, 0x12, 0x27, 0x2e, 0x6f, 0x70, 0x6c, 0x6f, 0x67, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65,
	0x77, 0x61, 0x72, 0x65, 0x2e, 0x6c, 0x6f, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4c,
	0x6f, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x6f, 0x70, 0x6c,
	0x6f, 0x67, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x6c, 0x6f,
	0x67, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x6f, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x38, 0x5a, 0x36, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4e, 0x70, 0x6f, 0x6f, 0x6c, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f,
	0x72, 0x6d, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x6e, 0x70, 0x6f, 0x6f, 0x6c,
	0x2f, 0x6f, 0x70, 0x6c, 0x6f, 0x67, 0x2f, 0x6d, 0x77, 0x2f, 0x76, 0x31, 0x2f, 0x6c, 0x6f, 0x67,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_npool_oplog_mw_v1_log_log_proto_rawDescOnce sync.Once
	file_npool_oplog_mw_v1_log_log_proto_rawDescData = file_npool_oplog_mw_v1_log_log_proto_rawDesc
)

func file_npool_oplog_mw_v1_log_log_proto_rawDescGZIP() []byte {
	file_npool_oplog_mw_v1_log_log_proto_rawDescOnce.Do(func() {
		file_npool_oplog_mw_v1_log_log_proto_rawDescData = protoimpl.X.CompressGZIP(file_npool_oplog_mw_v1_log_log_proto_rawDescData)
	})
	return file_npool_oplog_mw_v1_log_log_proto_rawDescData
}

var file_npool_oplog_mw_v1_log_log_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_npool_oplog_mw_v1_log_log_proto_goTypes = []interface{}{
	(*LogReq)(nil),            // 0: oplog.middleware.log.v1.LogReq
	(*Log)(nil),               // 1: oplog.middleware.log.v1.Log
	(*Conds)(nil),             // 2: oplog.middleware.log.v1.Conds
	(*CreateLogRequest)(nil),  // 3: oplog.middleware.log.v1.CreateLogRequest
	(*CreateLogResponse)(nil), // 4: oplog.middleware.log.v1.CreateLogResponse
	(*UpdateLogRequest)(nil),  // 5: oplog.middleware.log.v1.UpdateLogRequest
	(*UpdateLogResponse)(nil), // 6: oplog.middleware.log.v1.UpdateLogResponse
	(*GetLogsRequest)(nil),    // 7: oplog.middleware.log.v1.GetLogsRequest
	(*GetLogsResponse)(nil),   // 8: oplog.middleware.log.v1.GetLogsResponse
	(v1.HTTPMethod)(0),        // 9: basetypes.v1.HTTPMethod
	(v1.Result)(0),            // 10: basetypes.v1.Result
	(*v1.StringVal)(nil),      // 11: basetypes.v1.StringVal
	(*v1.Uint32Val)(nil),      // 12: basetypes.v1.Uint32Val
}
var file_npool_oplog_mw_v1_log_log_proto_depIdxs = []int32{
	9,  // 0: oplog.middleware.log.v1.LogReq.Method:type_name -> basetypes.v1.HTTPMethod
	10, // 1: oplog.middleware.log.v1.LogReq.Result:type_name -> basetypes.v1.Result
	9,  // 2: oplog.middleware.log.v1.Log.Method:type_name -> basetypes.v1.HTTPMethod
	10, // 3: oplog.middleware.log.v1.Log.Result:type_name -> basetypes.v1.Result
	11, // 4: oplog.middleware.log.v1.Conds.AppID:type_name -> basetypes.v1.StringVal
	11, // 5: oplog.middleware.log.v1.Conds.UserID:type_name -> basetypes.v1.StringVal
	12, // 6: oplog.middleware.log.v1.Conds.Result:type_name -> basetypes.v1.Uint32Val
	0,  // 7: oplog.middleware.log.v1.CreateLogRequest.Info:type_name -> oplog.middleware.log.v1.LogReq
	1,  // 8: oplog.middleware.log.v1.CreateLogResponse.Info:type_name -> oplog.middleware.log.v1.Log
	0,  // 9: oplog.middleware.log.v1.UpdateLogRequest.Info:type_name -> oplog.middleware.log.v1.LogReq
	1,  // 10: oplog.middleware.log.v1.UpdateLogResponse.Info:type_name -> oplog.middleware.log.v1.Log
	2,  // 11: oplog.middleware.log.v1.GetLogsRequest.Conds:type_name -> oplog.middleware.log.v1.Conds
	1,  // 12: oplog.middleware.log.v1.GetLogsResponse.Infos:type_name -> oplog.middleware.log.v1.Log
	3,  // 13: oplog.middleware.log.v1.Middleware.CreateLog:input_type -> oplog.middleware.log.v1.CreateLogRequest
	5,  // 14: oplog.middleware.log.v1.Middleware.UpdateLog:input_type -> oplog.middleware.log.v1.UpdateLogRequest
	7,  // 15: oplog.middleware.log.v1.Middleware.GetLogs:input_type -> oplog.middleware.log.v1.GetLogsRequest
	4,  // 16: oplog.middleware.log.v1.Middleware.CreateLog:output_type -> oplog.middleware.log.v1.CreateLogResponse
	6,  // 17: oplog.middleware.log.v1.Middleware.UpdateLog:output_type -> oplog.middleware.log.v1.UpdateLogResponse
	8,  // 18: oplog.middleware.log.v1.Middleware.GetLogs:output_type -> oplog.middleware.log.v1.GetLogsResponse
	16, // [16:19] is the sub-list for method output_type
	13, // [13:16] is the sub-list for method input_type
	13, // [13:13] is the sub-list for extension type_name
	13, // [13:13] is the sub-list for extension extendee
	0,  // [0:13] is the sub-list for field type_name
}

func init() { file_npool_oplog_mw_v1_log_log_proto_init() }
func file_npool_oplog_mw_v1_log_log_proto_init() {
	if File_npool_oplog_mw_v1_log_log_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_npool_oplog_mw_v1_log_log_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogReq); i {
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
		file_npool_oplog_mw_v1_log_log_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Log); i {
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
		file_npool_oplog_mw_v1_log_log_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
		file_npool_oplog_mw_v1_log_log_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateLogRequest); i {
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
		file_npool_oplog_mw_v1_log_log_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateLogResponse); i {
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
		file_npool_oplog_mw_v1_log_log_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateLogRequest); i {
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
		file_npool_oplog_mw_v1_log_log_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateLogResponse); i {
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
		file_npool_oplog_mw_v1_log_log_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetLogsRequest); i {
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
		file_npool_oplog_mw_v1_log_log_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetLogsResponse); i {
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
	file_npool_oplog_mw_v1_log_log_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_npool_oplog_mw_v1_log_log_proto_msgTypes[1].OneofWrappers = []interface{}{}
	file_npool_oplog_mw_v1_log_log_proto_msgTypes[2].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_npool_oplog_mw_v1_log_log_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_npool_oplog_mw_v1_log_log_proto_goTypes,
		DependencyIndexes: file_npool_oplog_mw_v1_log_log_proto_depIdxs,
		MessageInfos:      file_npool_oplog_mw_v1_log_log_proto_msgTypes,
	}.Build()
	File_npool_oplog_mw_v1_log_log_proto = out.File
	file_npool_oplog_mw_v1_log_log_proto_rawDesc = nil
	file_npool_oplog_mw_v1_log_log_proto_goTypes = nil
	file_npool_oplog_mw_v1_log_log_proto_depIdxs = nil
}
