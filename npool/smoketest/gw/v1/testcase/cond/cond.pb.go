// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.20.1
// source: npool/smoketest/gw/v1/testcase/cond/cond.proto

package cond

import (
	cond "github.com/NpoolPlatform/message/npool/smoketest/mgr/v1/testcase/cond"
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

type CreateCondRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TestCaseID     string        `protobuf:"bytes,10,opt,name=TestCaseID,proto3" json:"TestCaseID,omitempty"`
	CondTestCaseID string        `protobuf:"bytes,20,opt,name=CondTestCaseID,proto3" json:"CondTestCaseID,omitempty"`
	ArgumentMap    string        `protobuf:"bytes,30,opt,name=ArgumentMap,proto3" json:"ArgumentMap,omitempty"`
	Index          uint32        `protobuf:"varint,40,opt,name=Index,proto3" json:"Index,omitempty"`
	CondType       cond.CondType `protobuf:"varint,50,opt,name=CondType,proto3,enum=smoketest.manager.testcase.cond.v1.CondType" json:"CondType,omitempty"`
}

func (x *CreateCondRequest) Reset() {
	*x = CreateCondRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_smoketest_gw_v1_testcase_cond_cond_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCondRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCondRequest) ProtoMessage() {}

func (x *CreateCondRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_smoketest_gw_v1_testcase_cond_cond_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCondRequest.ProtoReflect.Descriptor instead.
func (*CreateCondRequest) Descriptor() ([]byte, []int) {
	return file_npool_smoketest_gw_v1_testcase_cond_cond_proto_rawDescGZIP(), []int{0}
}

func (x *CreateCondRequest) GetTestCaseID() string {
	if x != nil {
		return x.TestCaseID
	}
	return ""
}

func (x *CreateCondRequest) GetCondTestCaseID() string {
	if x != nil {
		return x.CondTestCaseID
	}
	return ""
}

func (x *CreateCondRequest) GetArgumentMap() string {
	if x != nil {
		return x.ArgumentMap
	}
	return ""
}

func (x *CreateCondRequest) GetIndex() uint32 {
	if x != nil {
		return x.Index
	}
	return 0
}

func (x *CreateCondRequest) GetCondType() cond.CondType {
	if x != nil {
		return x.CondType
	}
	return cond.CondType(0)
}

type CreateCondResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *cond.Cond `protobuf:"bytes,10,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (x *CreateCondResponse) Reset() {
	*x = CreateCondResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_smoketest_gw_v1_testcase_cond_cond_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCondResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCondResponse) ProtoMessage() {}

func (x *CreateCondResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_smoketest_gw_v1_testcase_cond_cond_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCondResponse.ProtoReflect.Descriptor instead.
func (*CreateCondResponse) Descriptor() ([]byte, []int) {
	return file_npool_smoketest_gw_v1_testcase_cond_cond_proto_rawDescGZIP(), []int{1}
}

func (x *CreateCondResponse) GetInfo() *cond.Cond {
	if x != nil {
		return x.Info
	}
	return nil
}

type DeleteCondRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID string `protobuf:"bytes,10,opt,name=ID,proto3" json:"ID,omitempty"`
}

func (x *DeleteCondRequest) Reset() {
	*x = DeleteCondRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_smoketest_gw_v1_testcase_cond_cond_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteCondRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteCondRequest) ProtoMessage() {}

func (x *DeleteCondRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_smoketest_gw_v1_testcase_cond_cond_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteCondRequest.ProtoReflect.Descriptor instead.
func (*DeleteCondRequest) Descriptor() ([]byte, []int) {
	return file_npool_smoketest_gw_v1_testcase_cond_cond_proto_rawDescGZIP(), []int{2}
}

func (x *DeleteCondRequest) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

type DeleteCondResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *cond.Cond `protobuf:"bytes,10,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (x *DeleteCondResponse) Reset() {
	*x = DeleteCondResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_smoketest_gw_v1_testcase_cond_cond_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteCondResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteCondResponse) ProtoMessage() {}

func (x *DeleteCondResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_smoketest_gw_v1_testcase_cond_cond_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteCondResponse.ProtoReflect.Descriptor instead.
func (*DeleteCondResponse) Descriptor() ([]byte, []int) {
	return file_npool_smoketest_gw_v1_testcase_cond_cond_proto_rawDescGZIP(), []int{3}
}

func (x *DeleteCondResponse) GetInfo() *cond.Cond {
	if x != nil {
		return x.Info
	}
	return nil
}

type GetCondsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Offset int32 `protobuf:"varint,10,opt,name=Offset,proto3" json:"Offset,omitempty"`
	Limit  int32 `protobuf:"varint,20,opt,name=Limit,proto3" json:"Limit,omitempty"`
}

func (x *GetCondsRequest) Reset() {
	*x = GetCondsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_smoketest_gw_v1_testcase_cond_cond_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCondsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCondsRequest) ProtoMessage() {}

func (x *GetCondsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_smoketest_gw_v1_testcase_cond_cond_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCondsRequest.ProtoReflect.Descriptor instead.
func (*GetCondsRequest) Descriptor() ([]byte, []int) {
	return file_npool_smoketest_gw_v1_testcase_cond_cond_proto_rawDescGZIP(), []int{4}
}

func (x *GetCondsRequest) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *GetCondsRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type GetCondsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Infos []*cond.Cond `protobuf:"bytes,10,rep,name=Infos,proto3" json:"Infos,omitempty"`
	Total uint32       `protobuf:"varint,20,opt,name=Total,proto3" json:"Total,omitempty"`
}

func (x *GetCondsResponse) Reset() {
	*x = GetCondsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_smoketest_gw_v1_testcase_cond_cond_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCondsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCondsResponse) ProtoMessage() {}

func (x *GetCondsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_smoketest_gw_v1_testcase_cond_cond_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCondsResponse.ProtoReflect.Descriptor instead.
func (*GetCondsResponse) Descriptor() ([]byte, []int) {
	return file_npool_smoketest_gw_v1_testcase_cond_cond_proto_rawDescGZIP(), []int{5}
}

func (x *GetCondsResponse) GetInfos() []*cond.Cond {
	if x != nil {
		return x.Infos
	}
	return nil
}

func (x *GetCondsResponse) GetTotal() uint32 {
	if x != nil {
		return x.Total
	}
	return 0
}

type UpdateCondRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID          string  `protobuf:"bytes,10,opt,name=ID,proto3" json:"ID,omitempty"`
	ArgumentMap *string `protobuf:"bytes,20,opt,name=ArgumentMap,proto3,oneof" json:"ArgumentMap,omitempty"`
	Index       *uint32 `protobuf:"varint,30,opt,name=Index,proto3,oneof" json:"Index,omitempty"`
}

func (x *UpdateCondRequest) Reset() {
	*x = UpdateCondRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_smoketest_gw_v1_testcase_cond_cond_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateCondRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateCondRequest) ProtoMessage() {}

func (x *UpdateCondRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_smoketest_gw_v1_testcase_cond_cond_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateCondRequest.ProtoReflect.Descriptor instead.
func (*UpdateCondRequest) Descriptor() ([]byte, []int) {
	return file_npool_smoketest_gw_v1_testcase_cond_cond_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateCondRequest) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *UpdateCondRequest) GetArgumentMap() string {
	if x != nil && x.ArgumentMap != nil {
		return *x.ArgumentMap
	}
	return ""
}

func (x *UpdateCondRequest) GetIndex() uint32 {
	if x != nil && x.Index != nil {
		return *x.Index
	}
	return 0
}

type UpdateCondResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *cond.Cond `protobuf:"bytes,10,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (x *UpdateCondResponse) Reset() {
	*x = UpdateCondResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_smoketest_gw_v1_testcase_cond_cond_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateCondResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateCondResponse) ProtoMessage() {}

func (x *UpdateCondResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_smoketest_gw_v1_testcase_cond_cond_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateCondResponse.ProtoReflect.Descriptor instead.
func (*UpdateCondResponse) Descriptor() ([]byte, []int) {
	return file_npool_smoketest_gw_v1_testcase_cond_cond_proto_rawDescGZIP(), []int{7}
}

func (x *UpdateCondResponse) GetInfo() *cond.Cond {
	if x != nil {
		return x.Info
	}
	return nil
}

var File_npool_smoketest_gw_v1_testcase_cond_cond_proto protoreflect.FileDescriptor

var file_npool_smoketest_gw_v1_testcase_cond_cond_proto_rawDesc = []byte{
	0x0a, 0x2e, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x73, 0x6d, 0x6f, 0x6b, 0x65, 0x74, 0x65, 0x73,
	0x74, 0x2f, 0x67, 0x77, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x63, 0x61, 0x73, 0x65,
	0x2f, 0x63, 0x6f, 0x6e, 0x64, 0x2f, 0x63, 0x6f, 0x6e, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x22, 0x73, 0x6d, 0x6f, 0x6b, 0x65, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x67, 0x61, 0x74, 0x65,
	0x77, 0x61, 0x79, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x63, 0x61, 0x73, 0x65, 0x2e, 0x63, 0x6f, 0x6e,
	0x64, 0x2e, 0x76, 0x31, 0x1a, 0x2f, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x73, 0x6d, 0x6f, 0x6b,
	0x65, 0x74, 0x65, 0x73, 0x74, 0x2f, 0x6d, 0x67, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x65, 0x73,
	0x74, 0x63, 0x61, 0x73, 0x65, 0x2f, 0x63, 0x6f, 0x6e, 0x64, 0x2f, 0x63, 0x6f, 0x6e, 0x64, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xdd, 0x01, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f,
	0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x54, 0x65, 0x73,
	0x74, 0x43, 0x61, 0x73, 0x65, 0x49, 0x44, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x54,
	0x65, 0x73, 0x74, 0x43, 0x61, 0x73, 0x65, 0x49, 0x44, 0x12, 0x26, 0x0a, 0x0e, 0x43, 0x6f, 0x6e,
	0x64, 0x54, 0x65, 0x73, 0x74, 0x43, 0x61, 0x73, 0x65, 0x49, 0x44, 0x18, 0x14, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0e, 0x43, 0x6f, 0x6e, 0x64, 0x54, 0x65, 0x73, 0x74, 0x43, 0x61, 0x73, 0x65, 0x49,
	0x44, 0x12, 0x20, 0x0a, 0x0b, 0x41, 0x72, 0x67, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x4d, 0x61, 0x70,
	0x18, 0x1e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x41, 0x72, 0x67, 0x75, 0x6d, 0x65, 0x6e, 0x74,
	0x4d, 0x61, 0x70, 0x12, 0x14, 0x0a, 0x05, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x28, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x05, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x48, 0x0a, 0x08, 0x43, 0x6f, 0x6e,
	0x64, 0x54, 0x79, 0x70, 0x65, 0x18, 0x32, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2c, 0x2e, 0x73, 0x6d,
	0x6f, 0x6b, 0x65, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e,
	0x74, 0x65, 0x73, 0x74, 0x63, 0x61, 0x73, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x64, 0x2e, 0x76, 0x31,
	0x2e, 0x43, 0x6f, 0x6e, 0x64, 0x54, 0x79, 0x70, 0x65, 0x52, 0x08, 0x43, 0x6f, 0x6e, 0x64, 0x54,
	0x79, 0x70, 0x65, 0x22, 0x52, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e,
	0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3c, 0x0a, 0x04, 0x49, 0x6e, 0x66,
	0x6f, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x73, 0x6d, 0x6f, 0x6b, 0x65, 0x74,
	0x65, 0x73, 0x74, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x74, 0x65, 0x73, 0x74,
	0x63, 0x61, 0x73, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6e,
	0x64, 0x52, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x23, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x43, 0x6f, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x49, 0x44, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x22, 0x52, 0x0a, 0x12,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x3c, 0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x28, 0x2e, 0x73, 0x6d, 0x6f, 0x6b, 0x65, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x6d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x72, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x63, 0x61, 0x73, 0x65, 0x2e, 0x63, 0x6f,
	0x6e, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x64, 0x52, 0x04, 0x49, 0x6e, 0x66, 0x6f,
	0x22, 0x3f, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x06, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x4c,
	0x69, 0x6d, 0x69, 0x74, 0x18, 0x14, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x4c, 0x69, 0x6d, 0x69,
	0x74, 0x22, 0x68, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x64, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3e, 0x0a, 0x05, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x18, 0x0a,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x73, 0x6d, 0x6f, 0x6b, 0x65, 0x74, 0x65, 0x73, 0x74,
	0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x63, 0x61, 0x73,
	0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x64, 0x52, 0x05,
	0x49, 0x6e, 0x66, 0x6f, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x14,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x22, 0x7f, 0x0a, 0x11, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44,
	0x12, 0x25, 0x0a, 0x0b, 0x41, 0x72, 0x67, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x4d, 0x61, 0x70, 0x18,
	0x14, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0b, 0x41, 0x72, 0x67, 0x75, 0x6d, 0x65, 0x6e,
	0x74, 0x4d, 0x61, 0x70, 0x88, 0x01, 0x01, 0x12, 0x19, 0x0a, 0x05, 0x49, 0x6e, 0x64, 0x65, 0x78,
	0x18, 0x1e, 0x20, 0x01, 0x28, 0x0d, 0x48, 0x01, 0x52, 0x05, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x88,
	0x01, 0x01, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x41, 0x72, 0x67, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x4d,
	0x61, 0x70, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x22, 0x52, 0x0a, 0x12,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x3c, 0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x28, 0x2e, 0x73, 0x6d, 0x6f, 0x6b, 0x65, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x6d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x72, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x63, 0x61, 0x73, 0x65, 0x2e, 0x63, 0x6f,
	0x6e, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x64, 0x52, 0x04, 0x49, 0x6e, 0x66, 0x6f,
	0x32, 0x89, 0x05, 0x0a, 0x07, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x12, 0x9f, 0x01, 0x0a,
	0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x64, 0x12, 0x35, 0x2e, 0x73, 0x6d,
	0x6f, 0x6b, 0x65, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e,
	0x74, 0x65, 0x73, 0x74, 0x63, 0x61, 0x73, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x64, 0x2e, 0x76, 0x31,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x36, 0x2e, 0x73, 0x6d, 0x6f, 0x6b, 0x65, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x67,
	0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x63, 0x61, 0x73, 0x65, 0x2e,
	0x63, 0x6f, 0x6e, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f,
	0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x22, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x1c, 0x22, 0x17, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x2f, 0x74,
	0x65, 0x73, 0x74, 0x63, 0x61, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x3a, 0x01, 0x2a, 0x12, 0x9f,
	0x01, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x64, 0x12, 0x35, 0x2e,
	0x73, 0x6d, 0x6f, 0x6b, 0x65, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61,
	0x79, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x63, 0x61, 0x73, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x64, 0x2e,
	0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x64, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x36, 0x2e, 0x73, 0x6d, 0x6f, 0x6b, 0x65, 0x74, 0x65, 0x73, 0x74,
	0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x63, 0x61, 0x73,
	0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x43, 0x6f, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x22, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x1c, 0x22, 0x17, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x2f, 0x74, 0x65, 0x73, 0x74, 0x63, 0x61, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x3a, 0x01, 0x2a,
	0x12, 0x9f, 0x01, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x64, 0x12,
	0x35, 0x2e, 0x73, 0x6d, 0x6f, 0x6b, 0x65, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x67, 0x61, 0x74, 0x65,
	0x77, 0x61, 0x79, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x63, 0x61, 0x73, 0x65, 0x2e, 0x63, 0x6f, 0x6e,
	0x64, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x64, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x36, 0x2e, 0x73, 0x6d, 0x6f, 0x6b, 0x65, 0x74, 0x65,
	0x73, 0x74, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x63,
	0x61, 0x73, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x43, 0x6f, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x22,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1c, 0x22, 0x17, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x63, 0x61, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x3a,
	0x01, 0x2a, 0x12, 0x97, 0x01, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x64, 0x73, 0x12,
	0x33, 0x2e, 0x73, 0x6d, 0x6f, 0x6b, 0x65, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x67, 0x61, 0x74, 0x65,
	0x77, 0x61, 0x79, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x63, 0x61, 0x73, 0x65, 0x2e, 0x63, 0x6f, 0x6e,
	0x64, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x64, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x34, 0x2e, 0x73, 0x6d, 0x6f, 0x6b, 0x65, 0x74, 0x65, 0x73, 0x74,
	0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x63, 0x61, 0x73,
	0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e,
	0x64, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x20, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x1a, 0x22, 0x15, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x65, 0x74, 0x2f, 0x74, 0x65, 0x73, 0x74,
	0x63, 0x61, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x3a, 0x01, 0x2a, 0x42, 0x46, 0x5a, 0x44,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4e, 0x70, 0x6f, 0x6f, 0x6c,
	0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x2f, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x73, 0x6d, 0x6f, 0x6b, 0x65, 0x74, 0x65, 0x73, 0x74,
	0x2f, 0x67, 0x77, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x63, 0x61, 0x73, 0x65, 0x2f,
	0x63, 0x6f, 0x6e, 0x64, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_npool_smoketest_gw_v1_testcase_cond_cond_proto_rawDescOnce sync.Once
	file_npool_smoketest_gw_v1_testcase_cond_cond_proto_rawDescData = file_npool_smoketest_gw_v1_testcase_cond_cond_proto_rawDesc
)

func file_npool_smoketest_gw_v1_testcase_cond_cond_proto_rawDescGZIP() []byte {
	file_npool_smoketest_gw_v1_testcase_cond_cond_proto_rawDescOnce.Do(func() {
		file_npool_smoketest_gw_v1_testcase_cond_cond_proto_rawDescData = protoimpl.X.CompressGZIP(file_npool_smoketest_gw_v1_testcase_cond_cond_proto_rawDescData)
	})
	return file_npool_smoketest_gw_v1_testcase_cond_cond_proto_rawDescData
}

var file_npool_smoketest_gw_v1_testcase_cond_cond_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_npool_smoketest_gw_v1_testcase_cond_cond_proto_goTypes = []interface{}{
	(*CreateCondRequest)(nil),  // 0: smoketest.gateway.testcase.cond.v1.CreateCondRequest
	(*CreateCondResponse)(nil), // 1: smoketest.gateway.testcase.cond.v1.CreateCondResponse
	(*DeleteCondRequest)(nil),  // 2: smoketest.gateway.testcase.cond.v1.DeleteCondRequest
	(*DeleteCondResponse)(nil), // 3: smoketest.gateway.testcase.cond.v1.DeleteCondResponse
	(*GetCondsRequest)(nil),    // 4: smoketest.gateway.testcase.cond.v1.GetCondsRequest
	(*GetCondsResponse)(nil),   // 5: smoketest.gateway.testcase.cond.v1.GetCondsResponse
	(*UpdateCondRequest)(nil),  // 6: smoketest.gateway.testcase.cond.v1.UpdateCondRequest
	(*UpdateCondResponse)(nil), // 7: smoketest.gateway.testcase.cond.v1.UpdateCondResponse
	(cond.CondType)(0),         // 8: smoketest.manager.testcase.cond.v1.CondType
	(*cond.Cond)(nil),          // 9: smoketest.manager.testcase.cond.v1.Cond
}
var file_npool_smoketest_gw_v1_testcase_cond_cond_proto_depIdxs = []int32{
	8, // 0: smoketest.gateway.testcase.cond.v1.CreateCondRequest.CondType:type_name -> smoketest.manager.testcase.cond.v1.CondType
	9, // 1: smoketest.gateway.testcase.cond.v1.CreateCondResponse.Info:type_name -> smoketest.manager.testcase.cond.v1.Cond
	9, // 2: smoketest.gateway.testcase.cond.v1.DeleteCondResponse.Info:type_name -> smoketest.manager.testcase.cond.v1.Cond
	9, // 3: smoketest.gateway.testcase.cond.v1.GetCondsResponse.Infos:type_name -> smoketest.manager.testcase.cond.v1.Cond
	9, // 4: smoketest.gateway.testcase.cond.v1.UpdateCondResponse.Info:type_name -> smoketest.manager.testcase.cond.v1.Cond
	0, // 5: smoketest.gateway.testcase.cond.v1.Gateway.CreateCond:input_type -> smoketest.gateway.testcase.cond.v1.CreateCondRequest
	2, // 6: smoketest.gateway.testcase.cond.v1.Gateway.DeleteCond:input_type -> smoketest.gateway.testcase.cond.v1.DeleteCondRequest
	6, // 7: smoketest.gateway.testcase.cond.v1.Gateway.UpdateCond:input_type -> smoketest.gateway.testcase.cond.v1.UpdateCondRequest
	4, // 8: smoketest.gateway.testcase.cond.v1.Gateway.GetConds:input_type -> smoketest.gateway.testcase.cond.v1.GetCondsRequest
	1, // 9: smoketest.gateway.testcase.cond.v1.Gateway.CreateCond:output_type -> smoketest.gateway.testcase.cond.v1.CreateCondResponse
	3, // 10: smoketest.gateway.testcase.cond.v1.Gateway.DeleteCond:output_type -> smoketest.gateway.testcase.cond.v1.DeleteCondResponse
	7, // 11: smoketest.gateway.testcase.cond.v1.Gateway.UpdateCond:output_type -> smoketest.gateway.testcase.cond.v1.UpdateCondResponse
	5, // 12: smoketest.gateway.testcase.cond.v1.Gateway.GetConds:output_type -> smoketest.gateway.testcase.cond.v1.GetCondsResponse
	9, // [9:13] is the sub-list for method output_type
	5, // [5:9] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_npool_smoketest_gw_v1_testcase_cond_cond_proto_init() }
func file_npool_smoketest_gw_v1_testcase_cond_cond_proto_init() {
	if File_npool_smoketest_gw_v1_testcase_cond_cond_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_npool_smoketest_gw_v1_testcase_cond_cond_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateCondRequest); i {
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
		file_npool_smoketest_gw_v1_testcase_cond_cond_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateCondResponse); i {
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
		file_npool_smoketest_gw_v1_testcase_cond_cond_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteCondRequest); i {
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
		file_npool_smoketest_gw_v1_testcase_cond_cond_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteCondResponse); i {
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
		file_npool_smoketest_gw_v1_testcase_cond_cond_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCondsRequest); i {
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
		file_npool_smoketest_gw_v1_testcase_cond_cond_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCondsResponse); i {
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
		file_npool_smoketest_gw_v1_testcase_cond_cond_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateCondRequest); i {
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
		file_npool_smoketest_gw_v1_testcase_cond_cond_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateCondResponse); i {
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
	file_npool_smoketest_gw_v1_testcase_cond_cond_proto_msgTypes[6].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_npool_smoketest_gw_v1_testcase_cond_cond_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_npool_smoketest_gw_v1_testcase_cond_cond_proto_goTypes,
		DependencyIndexes: file_npool_smoketest_gw_v1_testcase_cond_cond_proto_depIdxs,
		MessageInfos:      file_npool_smoketest_gw_v1_testcase_cond_cond_proto_msgTypes,
	}.Build()
	File_npool_smoketest_gw_v1_testcase_cond_cond_proto = out.File
	file_npool_smoketest_gw_v1_testcase_cond_cond_proto_rawDesc = nil
	file_npool_smoketest_gw_v1_testcase_cond_cond_proto_goTypes = nil
	file_npool_smoketest_gw_v1_testcase_cond_cond_proto_depIdxs = nil
}
