// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: npool/smoketest/mw/v1/testplan/testplan.proto

package testplan

import (
	testplan "github.com/NpoolPlatform/message/npool/smoketest/mgr/v1/testplan"
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

type CreateTestPlanRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *testplan.TestPlanReq `protobuf:"bytes,10,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (x *CreateTestPlanRequest) Reset() {
	*x = CreateTestPlanRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_smoketest_mw_v1_testplan_testplan_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTestPlanRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTestPlanRequest) ProtoMessage() {}

func (x *CreateTestPlanRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_smoketest_mw_v1_testplan_testplan_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTestPlanRequest.ProtoReflect.Descriptor instead.
func (*CreateTestPlanRequest) Descriptor() ([]byte, []int) {
	return file_npool_smoketest_mw_v1_testplan_testplan_proto_rawDescGZIP(), []int{0}
}

func (x *CreateTestPlanRequest) GetInfo() *testplan.TestPlanReq {
	if x != nil {
		return x.Info
	}
	return nil
}

type CreateTestPlanResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *testplan.TestPlan `protobuf:"bytes,10,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (x *CreateTestPlanResponse) Reset() {
	*x = CreateTestPlanResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_smoketest_mw_v1_testplan_testplan_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTestPlanResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTestPlanResponse) ProtoMessage() {}

func (x *CreateTestPlanResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_smoketest_mw_v1_testplan_testplan_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTestPlanResponse.ProtoReflect.Descriptor instead.
func (*CreateTestPlanResponse) Descriptor() ([]byte, []int) {
	return file_npool_smoketest_mw_v1_testplan_testplan_proto_rawDescGZIP(), []int{1}
}

func (x *CreateTestPlanResponse) GetInfo() *testplan.TestPlan {
	if x != nil {
		return x.Info
	}
	return nil
}

type UpdateTestPlanRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *testplan.TestPlanReq `protobuf:"bytes,10,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (x *UpdateTestPlanRequest) Reset() {
	*x = UpdateTestPlanRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_smoketest_mw_v1_testplan_testplan_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateTestPlanRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateTestPlanRequest) ProtoMessage() {}

func (x *UpdateTestPlanRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_smoketest_mw_v1_testplan_testplan_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateTestPlanRequest.ProtoReflect.Descriptor instead.
func (*UpdateTestPlanRequest) Descriptor() ([]byte, []int) {
	return file_npool_smoketest_mw_v1_testplan_testplan_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateTestPlanRequest) GetInfo() *testplan.TestPlanReq {
	if x != nil {
		return x.Info
	}
	return nil
}

type UpdateTestPlanResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *testplan.TestPlan `protobuf:"bytes,10,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (x *UpdateTestPlanResponse) Reset() {
	*x = UpdateTestPlanResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_smoketest_mw_v1_testplan_testplan_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateTestPlanResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateTestPlanResponse) ProtoMessage() {}

func (x *UpdateTestPlanResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_smoketest_mw_v1_testplan_testplan_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateTestPlanResponse.ProtoReflect.Descriptor instead.
func (*UpdateTestPlanResponse) Descriptor() ([]byte, []int) {
	return file_npool_smoketest_mw_v1_testplan_testplan_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateTestPlanResponse) GetInfo() *testplan.TestPlan {
	if x != nil {
		return x.Info
	}
	return nil
}

type GetTestPlanRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID string `protobuf:"bytes,10,opt,name=ID,proto3" json:"ID,omitempty"`
}

func (x *GetTestPlanRequest) Reset() {
	*x = GetTestPlanRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_smoketest_mw_v1_testplan_testplan_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTestPlanRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTestPlanRequest) ProtoMessage() {}

func (x *GetTestPlanRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_smoketest_mw_v1_testplan_testplan_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTestPlanRequest.ProtoReflect.Descriptor instead.
func (*GetTestPlanRequest) Descriptor() ([]byte, []int) {
	return file_npool_smoketest_mw_v1_testplan_testplan_proto_rawDescGZIP(), []int{4}
}

func (x *GetTestPlanRequest) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

type GetTestPlanResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *testplan.TestPlan `protobuf:"bytes,10,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (x *GetTestPlanResponse) Reset() {
	*x = GetTestPlanResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_smoketest_mw_v1_testplan_testplan_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTestPlanResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTestPlanResponse) ProtoMessage() {}

func (x *GetTestPlanResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_smoketest_mw_v1_testplan_testplan_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTestPlanResponse.ProtoReflect.Descriptor instead.
func (*GetTestPlanResponse) Descriptor() ([]byte, []int) {
	return file_npool_smoketest_mw_v1_testplan_testplan_proto_rawDescGZIP(), []int{5}
}

func (x *GetTestPlanResponse) GetInfo() *testplan.TestPlan {
	if x != nil {
		return x.Info
	}
	return nil
}

type GetTestPlansRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Conds  *testplan.Conds `protobuf:"bytes,10,opt,name=Conds,proto3" json:"Conds,omitempty"`
	Offset int32           `protobuf:"varint,20,opt,name=Offset,proto3" json:"Offset,omitempty"`
	Limit  int32           `protobuf:"varint,30,opt,name=Limit,proto3" json:"Limit,omitempty"`
}

func (x *GetTestPlansRequest) Reset() {
	*x = GetTestPlansRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_smoketest_mw_v1_testplan_testplan_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTestPlansRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTestPlansRequest) ProtoMessage() {}

func (x *GetTestPlansRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_smoketest_mw_v1_testplan_testplan_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTestPlansRequest.ProtoReflect.Descriptor instead.
func (*GetTestPlansRequest) Descriptor() ([]byte, []int) {
	return file_npool_smoketest_mw_v1_testplan_testplan_proto_rawDescGZIP(), []int{6}
}

func (x *GetTestPlansRequest) GetConds() *testplan.Conds {
	if x != nil {
		return x.Conds
	}
	return nil
}

func (x *GetTestPlansRequest) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *GetTestPlansRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type GetTestPlansResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Infos []*testplan.TestPlan `protobuf:"bytes,10,rep,name=Infos,proto3" json:"Infos,omitempty"`
	Total uint32               `protobuf:"varint,20,opt,name=Total,proto3" json:"Total,omitempty"`
}

func (x *GetTestPlansResponse) Reset() {
	*x = GetTestPlansResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_smoketest_mw_v1_testplan_testplan_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTestPlansResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTestPlansResponse) ProtoMessage() {}

func (x *GetTestPlansResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_smoketest_mw_v1_testplan_testplan_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTestPlansResponse.ProtoReflect.Descriptor instead.
func (*GetTestPlansResponse) Descriptor() ([]byte, []int) {
	return file_npool_smoketest_mw_v1_testplan_testplan_proto_rawDescGZIP(), []int{7}
}

func (x *GetTestPlansResponse) GetInfos() []*testplan.TestPlan {
	if x != nil {
		return x.Infos
	}
	return nil
}

func (x *GetTestPlansResponse) GetTotal() uint32 {
	if x != nil {
		return x.Total
	}
	return 0
}

type ExistTestPlanRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID string `protobuf:"bytes,10,opt,name=ID,proto3" json:"ID,omitempty"`
}

func (x *ExistTestPlanRequest) Reset() {
	*x = ExistTestPlanRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_smoketest_mw_v1_testplan_testplan_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExistTestPlanRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExistTestPlanRequest) ProtoMessage() {}

func (x *ExistTestPlanRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_smoketest_mw_v1_testplan_testplan_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExistTestPlanRequest.ProtoReflect.Descriptor instead.
func (*ExistTestPlanRequest) Descriptor() ([]byte, []int) {
	return file_npool_smoketest_mw_v1_testplan_testplan_proto_rawDescGZIP(), []int{8}
}

func (x *ExistTestPlanRequest) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

type ExistTestPlanResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info bool `protobuf:"varint,10,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (x *ExistTestPlanResponse) Reset() {
	*x = ExistTestPlanResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_smoketest_mw_v1_testplan_testplan_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExistTestPlanResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExistTestPlanResponse) ProtoMessage() {}

func (x *ExistTestPlanResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_smoketest_mw_v1_testplan_testplan_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExistTestPlanResponse.ProtoReflect.Descriptor instead.
func (*ExistTestPlanResponse) Descriptor() ([]byte, []int) {
	return file_npool_smoketest_mw_v1_testplan_testplan_proto_rawDescGZIP(), []int{9}
}

func (x *ExistTestPlanResponse) GetInfo() bool {
	if x != nil {
		return x.Info
	}
	return false
}

type DeleteTestPlanRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *testplan.TestPlanReq `protobuf:"bytes,10,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (x *DeleteTestPlanRequest) Reset() {
	*x = DeleteTestPlanRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_smoketest_mw_v1_testplan_testplan_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteTestPlanRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteTestPlanRequest) ProtoMessage() {}

func (x *DeleteTestPlanRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_smoketest_mw_v1_testplan_testplan_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteTestPlanRequest.ProtoReflect.Descriptor instead.
func (*DeleteTestPlanRequest) Descriptor() ([]byte, []int) {
	return file_npool_smoketest_mw_v1_testplan_testplan_proto_rawDescGZIP(), []int{10}
}

func (x *DeleteTestPlanRequest) GetInfo() *testplan.TestPlanReq {
	if x != nil {
		return x.Info
	}
	return nil
}

type DeleteTestPlanResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *testplan.TestPlan `protobuf:"bytes,10,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (x *DeleteTestPlanResponse) Reset() {
	*x = DeleteTestPlanResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_smoketest_mw_v1_testplan_testplan_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteTestPlanResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteTestPlanResponse) ProtoMessage() {}

func (x *DeleteTestPlanResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_smoketest_mw_v1_testplan_testplan_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteTestPlanResponse.ProtoReflect.Descriptor instead.
func (*DeleteTestPlanResponse) Descriptor() ([]byte, []int) {
	return file_npool_smoketest_mw_v1_testplan_testplan_proto_rawDescGZIP(), []int{11}
}

func (x *DeleteTestPlanResponse) GetInfo() *testplan.TestPlan {
	if x != nil {
		return x.Info
	}
	return nil
}

var File_npool_smoketest_mw_v1_testplan_testplan_proto protoreflect.FileDescriptor

var file_npool_smoketest_mw_v1_testplan_testplan_proto_rawDesc = []byte{
	0x0a, 0x2d, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x73, 0x6d, 0x6f, 0x6b, 0x65, 0x74, 0x65, 0x73,
	0x74, 0x2f, 0x6d, 0x77, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x70, 0x6c, 0x61, 0x6e,
	0x2f, 0x74, 0x65, 0x73, 0x74, 0x70, 0x6c, 0x61, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x20, 0x73, 0x6d, 0x6f, 0x6b, 0x65, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c,
	0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x70, 0x6c, 0x61, 0x6e, 0x2e, 0x76,
	0x31, 0x1a, 0x2e, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x73, 0x6d, 0x6f, 0x6b, 0x65, 0x74, 0x65,
	0x73, 0x74, 0x2f, 0x6d, 0x67, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x70, 0x6c,
	0x61, 0x6e, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x70, 0x6c, 0x61, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x57, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x65, 0x73, 0x74, 0x50,
	0x6c, 0x61, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3e, 0x0a, 0x04, 0x49, 0x6e,
	0x66, 0x6f, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x73, 0x6d, 0x6f, 0x6b, 0x65,
	0x74, 0x65, 0x73, 0x74, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x74, 0x65, 0x73,
	0x74, 0x70, 0x6c, 0x61, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x50, 0x6c, 0x61,
	0x6e, 0x52, 0x65, 0x71, 0x52, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x55, 0x0a, 0x16, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x54, 0x65, 0x73, 0x74, 0x50, 0x6c, 0x61, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3b, 0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x27, 0x2e, 0x73, 0x6d, 0x6f, 0x6b, 0x65, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x6d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x70, 0x6c, 0x61, 0x6e, 0x2e,
	0x76, 0x31, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x50, 0x6c, 0x61, 0x6e, 0x52, 0x04, 0x49, 0x6e, 0x66,
	0x6f, 0x22, 0x57, 0x0a, 0x15, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x65, 0x73, 0x74, 0x50,
	0x6c, 0x61, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3e, 0x0a, 0x04, 0x49, 0x6e,
	0x66, 0x6f, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x73, 0x6d, 0x6f, 0x6b, 0x65,
	0x74, 0x65, 0x73, 0x74, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x74, 0x65, 0x73,
	0x74, 0x70, 0x6c, 0x61, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x50, 0x6c, 0x61,
	0x6e, 0x52, 0x65, 0x71, 0x52, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x55, 0x0a, 0x16, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x54, 0x65, 0x73, 0x74, 0x50, 0x6c, 0x61, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3b, 0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x27, 0x2e, 0x73, 0x6d, 0x6f, 0x6b, 0x65, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x6d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x70, 0x6c, 0x61, 0x6e, 0x2e,
	0x76, 0x31, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x50, 0x6c, 0x61, 0x6e, 0x52, 0x04, 0x49, 0x6e, 0x66,
	0x6f, 0x22, 0x24, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x54, 0x65, 0x73, 0x74, 0x50, 0x6c, 0x61, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x22, 0x52, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x54, 0x65,
	0x73, 0x74, 0x50, 0x6c, 0x61, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3b,
	0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x73,
	0x6d, 0x6f, 0x6b, 0x65, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72,
	0x2e, 0x74, 0x65, 0x73, 0x74, 0x70, 0x6c, 0x61, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x65, 0x73,
	0x74, 0x50, 0x6c, 0x61, 0x6e, 0x52, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x7f, 0x0a, 0x13, 0x47,
	0x65, 0x74, 0x54, 0x65, 0x73, 0x74, 0x50, 0x6c, 0x61, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x3a, 0x0a, 0x05, 0x43, 0x6f, 0x6e, 0x64, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x24, 0x2e, 0x73, 0x6d, 0x6f, 0x6b, 0x65, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x6d, 0x61,
	0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x70, 0x6c, 0x61, 0x6e, 0x2e, 0x76,
	0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x64, 0x73, 0x52, 0x05, 0x43, 0x6f, 0x6e, 0x64, 0x73, 0x12, 0x16,
	0x0a, 0x06, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x14, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06,
	0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x18,
	0x1e, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x22, 0x6b, 0x0a, 0x14,
	0x47, 0x65, 0x74, 0x54, 0x65, 0x73, 0x74, 0x50, 0x6c, 0x61, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3d, 0x0a, 0x05, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x18, 0x0a, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x73, 0x6d, 0x6f, 0x6b, 0x65, 0x74, 0x65, 0x73, 0x74, 0x2e,
	0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x70, 0x6c, 0x61, 0x6e,
	0x2e, 0x76, 0x31, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x50, 0x6c, 0x61, 0x6e, 0x52, 0x05, 0x49, 0x6e,
	0x66, 0x6f, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x14, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x22, 0x26, 0x0a, 0x14, 0x45, 0x78, 0x69,
	0x73, 0x74, 0x54, 0x65, 0x73, 0x74, 0x50, 0x6c, 0x61, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49,
	0x44, 0x22, 0x2b, 0x0a, 0x15, 0x45, 0x78, 0x69, 0x73, 0x74, 0x54, 0x65, 0x73, 0x74, 0x50, 0x6c,
	0x61, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x49, 0x6e,
	0x66, 0x6f, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x57,
	0x0a, 0x15, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x65, 0x73, 0x74, 0x50, 0x6c, 0x61, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3e, 0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x73, 0x6d, 0x6f, 0x6b, 0x65, 0x74, 0x65, 0x73,
	0x74, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x70, 0x6c,
	0x61, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x50, 0x6c, 0x61, 0x6e, 0x52, 0x65,
	0x71, 0x52, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x55, 0x0a, 0x16, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x54, 0x65, 0x73, 0x74, 0x50, 0x6c, 0x61, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x3b, 0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x27, 0x2e, 0x73, 0x6d, 0x6f, 0x6b, 0x65, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x6d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x72, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x70, 0x6c, 0x61, 0x6e, 0x2e, 0x76, 0x31, 0x2e,
	0x54, 0x65, 0x73, 0x74, 0x50, 0x6c, 0x61, 0x6e, 0x52, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x32, 0xa8,
	0x06, 0x0a, 0x0a, 0x4d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x12, 0x85, 0x01,
	0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x65, 0x73, 0x74, 0x50, 0x6c, 0x61, 0x6e,
	0x12, 0x37, 0x2e, 0x73, 0x6d, 0x6f, 0x6b, 0x65, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x6d, 0x69, 0x64,
	0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x70, 0x6c, 0x61, 0x6e,
	0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x65, 0x73, 0x74, 0x50, 0x6c,
	0x61, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x38, 0x2e, 0x73, 0x6d, 0x6f, 0x6b,
	0x65, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65,
	0x2e, 0x74, 0x65, 0x73, 0x74, 0x70, 0x6c, 0x61, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x54, 0x65, 0x73, 0x74, 0x50, 0x6c, 0x61, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x85, 0x01, 0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x54, 0x65, 0x73, 0x74, 0x50, 0x6c, 0x61, 0x6e, 0x12, 0x37, 0x2e, 0x73, 0x6d, 0x6f, 0x6b, 0x65,
	0x74, 0x65, 0x73, 0x74, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e,
	0x74, 0x65, 0x73, 0x74, 0x70, 0x6c, 0x61, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x54, 0x65, 0x73, 0x74, 0x50, 0x6c, 0x61, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x38, 0x2e, 0x73, 0x6d, 0x6f, 0x6b, 0x65, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x6d, 0x69,
	0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x70, 0x6c, 0x61,
	0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x65, 0x73, 0x74, 0x50,
	0x6c, 0x61, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x7c, 0x0a,
	0x0b, 0x47, 0x65, 0x74, 0x54, 0x65, 0x73, 0x74, 0x50, 0x6c, 0x61, 0x6e, 0x12, 0x34, 0x2e, 0x73,
	0x6d, 0x6f, 0x6b, 0x65, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77,
	0x61, 0x72, 0x65, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x70, 0x6c, 0x61, 0x6e, 0x2e, 0x76, 0x31, 0x2e,
	0x47, 0x65, 0x74, 0x54, 0x65, 0x73, 0x74, 0x50, 0x6c, 0x61, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x35, 0x2e, 0x73, 0x6d, 0x6f, 0x6b, 0x65, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x6d,
	0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x70, 0x6c,
	0x61, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x65, 0x73, 0x74, 0x50, 0x6c, 0x61,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x7f, 0x0a, 0x0c, 0x47,
	0x65, 0x74, 0x54, 0x65, 0x73, 0x74, 0x50, 0x6c, 0x61, 0x6e, 0x73, 0x12, 0x35, 0x2e, 0x73, 0x6d,
	0x6f, 0x6b, 0x65, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61,
	0x72, 0x65, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x70, 0x6c, 0x61, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x47,
	0x65, 0x74, 0x54, 0x65, 0x73, 0x74, 0x50, 0x6c, 0x61, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x36, 0x2e, 0x73, 0x6d, 0x6f, 0x6b, 0x65, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x6d,
	0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x70, 0x6c,
	0x61, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x65, 0x73, 0x74, 0x50, 0x6c, 0x61,
	0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x82, 0x01, 0x0a,
	0x0d, 0x45, 0x78, 0x69, 0x73, 0x74, 0x54, 0x65, 0x73, 0x74, 0x50, 0x6c, 0x61, 0x6e, 0x12, 0x36,
	0x2e, 0x73, 0x6d, 0x6f, 0x6b, 0x65, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c,
	0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x70, 0x6c, 0x61, 0x6e, 0x2e, 0x76,
	0x31, 0x2e, 0x45, 0x78, 0x69, 0x73, 0x74, 0x54, 0x65, 0x73, 0x74, 0x50, 0x6c, 0x61, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x37, 0x2e, 0x73, 0x6d, 0x6f, 0x6b, 0x65, 0x74, 0x65,
	0x73, 0x74, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x74, 0x65,
	0x73, 0x74, 0x70, 0x6c, 0x61, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x78, 0x69, 0x73, 0x74, 0x54,
	0x65, 0x73, 0x74, 0x50, 0x6c, 0x61, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x85, 0x01, 0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x65, 0x73, 0x74,
	0x50, 0x6c, 0x61, 0x6e, 0x12, 0x37, 0x2e, 0x73, 0x6d, 0x6f, 0x6b, 0x65, 0x74, 0x65, 0x73, 0x74,
	0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x74, 0x65, 0x73, 0x74,
	0x70, 0x6c, 0x61, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x65,
	0x73, 0x74, 0x50, 0x6c, 0x61, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x38, 0x2e,
	0x73, 0x6d, 0x6f, 0x6b, 0x65, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65,
	0x77, 0x61, 0x72, 0x65, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x70, 0x6c, 0x61, 0x6e, 0x2e, 0x76, 0x31,
	0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x65, 0x73, 0x74, 0x50, 0x6c, 0x61, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x41, 0x5a, 0x3f, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4e, 0x70, 0x6f, 0x6f, 0x6c, 0x50, 0x6c, 0x61,
	0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x6e, 0x70,
	0x6f, 0x6f, 0x6c, 0x2f, 0x73, 0x6d, 0x6f, 0x6b, 0x65, 0x74, 0x65, 0x73, 0x74, 0x2f, 0x6d, 0x77,
	0x2f, 0x76, 0x31, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x70, 0x6c, 0x61, 0x6e, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_npool_smoketest_mw_v1_testplan_testplan_proto_rawDescOnce sync.Once
	file_npool_smoketest_mw_v1_testplan_testplan_proto_rawDescData = file_npool_smoketest_mw_v1_testplan_testplan_proto_rawDesc
)

func file_npool_smoketest_mw_v1_testplan_testplan_proto_rawDescGZIP() []byte {
	file_npool_smoketest_mw_v1_testplan_testplan_proto_rawDescOnce.Do(func() {
		file_npool_smoketest_mw_v1_testplan_testplan_proto_rawDescData = protoimpl.X.CompressGZIP(file_npool_smoketest_mw_v1_testplan_testplan_proto_rawDescData)
	})
	return file_npool_smoketest_mw_v1_testplan_testplan_proto_rawDescData
}

var file_npool_smoketest_mw_v1_testplan_testplan_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_npool_smoketest_mw_v1_testplan_testplan_proto_goTypes = []interface{}{
	(*CreateTestPlanRequest)(nil),  // 0: smoketest.middleware.testplan.v1.CreateTestPlanRequest
	(*CreateTestPlanResponse)(nil), // 1: smoketest.middleware.testplan.v1.CreateTestPlanResponse
	(*UpdateTestPlanRequest)(nil),  // 2: smoketest.middleware.testplan.v1.UpdateTestPlanRequest
	(*UpdateTestPlanResponse)(nil), // 3: smoketest.middleware.testplan.v1.UpdateTestPlanResponse
	(*GetTestPlanRequest)(nil),     // 4: smoketest.middleware.testplan.v1.GetTestPlanRequest
	(*GetTestPlanResponse)(nil),    // 5: smoketest.middleware.testplan.v1.GetTestPlanResponse
	(*GetTestPlansRequest)(nil),    // 6: smoketest.middleware.testplan.v1.GetTestPlansRequest
	(*GetTestPlansResponse)(nil),   // 7: smoketest.middleware.testplan.v1.GetTestPlansResponse
	(*ExistTestPlanRequest)(nil),   // 8: smoketest.middleware.testplan.v1.ExistTestPlanRequest
	(*ExistTestPlanResponse)(nil),  // 9: smoketest.middleware.testplan.v1.ExistTestPlanResponse
	(*DeleteTestPlanRequest)(nil),  // 10: smoketest.middleware.testplan.v1.DeleteTestPlanRequest
	(*DeleteTestPlanResponse)(nil), // 11: smoketest.middleware.testplan.v1.DeleteTestPlanResponse
	(*testplan.TestPlanReq)(nil),   // 12: smoketest.manager.testplan.v1.TestPlanReq
	(*testplan.TestPlan)(nil),      // 13: smoketest.manager.testplan.v1.TestPlan
	(*testplan.Conds)(nil),         // 14: smoketest.manager.testplan.v1.Conds
}
var file_npool_smoketest_mw_v1_testplan_testplan_proto_depIdxs = []int32{
	12, // 0: smoketest.middleware.testplan.v1.CreateTestPlanRequest.Info:type_name -> smoketest.manager.testplan.v1.TestPlanReq
	13, // 1: smoketest.middleware.testplan.v1.CreateTestPlanResponse.Info:type_name -> smoketest.manager.testplan.v1.TestPlan
	12, // 2: smoketest.middleware.testplan.v1.UpdateTestPlanRequest.Info:type_name -> smoketest.manager.testplan.v1.TestPlanReq
	13, // 3: smoketest.middleware.testplan.v1.UpdateTestPlanResponse.Info:type_name -> smoketest.manager.testplan.v1.TestPlan
	13, // 4: smoketest.middleware.testplan.v1.GetTestPlanResponse.Info:type_name -> smoketest.manager.testplan.v1.TestPlan
	14, // 5: smoketest.middleware.testplan.v1.GetTestPlansRequest.Conds:type_name -> smoketest.manager.testplan.v1.Conds
	13, // 6: smoketest.middleware.testplan.v1.GetTestPlansResponse.Infos:type_name -> smoketest.manager.testplan.v1.TestPlan
	12, // 7: smoketest.middleware.testplan.v1.DeleteTestPlanRequest.Info:type_name -> smoketest.manager.testplan.v1.TestPlanReq
	13, // 8: smoketest.middleware.testplan.v1.DeleteTestPlanResponse.Info:type_name -> smoketest.manager.testplan.v1.TestPlan
	0,  // 9: smoketest.middleware.testplan.v1.Middleware.CreateTestPlan:input_type -> smoketest.middleware.testplan.v1.CreateTestPlanRequest
	2,  // 10: smoketest.middleware.testplan.v1.Middleware.UpdateTestPlan:input_type -> smoketest.middleware.testplan.v1.UpdateTestPlanRequest
	4,  // 11: smoketest.middleware.testplan.v1.Middleware.GetTestPlan:input_type -> smoketest.middleware.testplan.v1.GetTestPlanRequest
	6,  // 12: smoketest.middleware.testplan.v1.Middleware.GetTestPlans:input_type -> smoketest.middleware.testplan.v1.GetTestPlansRequest
	8,  // 13: smoketest.middleware.testplan.v1.Middleware.ExistTestPlan:input_type -> smoketest.middleware.testplan.v1.ExistTestPlanRequest
	10, // 14: smoketest.middleware.testplan.v1.Middleware.DeleteTestPlan:input_type -> smoketest.middleware.testplan.v1.DeleteTestPlanRequest
	1,  // 15: smoketest.middleware.testplan.v1.Middleware.CreateTestPlan:output_type -> smoketest.middleware.testplan.v1.CreateTestPlanResponse
	3,  // 16: smoketest.middleware.testplan.v1.Middleware.UpdateTestPlan:output_type -> smoketest.middleware.testplan.v1.UpdateTestPlanResponse
	5,  // 17: smoketest.middleware.testplan.v1.Middleware.GetTestPlan:output_type -> smoketest.middleware.testplan.v1.GetTestPlanResponse
	7,  // 18: smoketest.middleware.testplan.v1.Middleware.GetTestPlans:output_type -> smoketest.middleware.testplan.v1.GetTestPlansResponse
	9,  // 19: smoketest.middleware.testplan.v1.Middleware.ExistTestPlan:output_type -> smoketest.middleware.testplan.v1.ExistTestPlanResponse
	11, // 20: smoketest.middleware.testplan.v1.Middleware.DeleteTestPlan:output_type -> smoketest.middleware.testplan.v1.DeleteTestPlanResponse
	15, // [15:21] is the sub-list for method output_type
	9,  // [9:15] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_npool_smoketest_mw_v1_testplan_testplan_proto_init() }
func file_npool_smoketest_mw_v1_testplan_testplan_proto_init() {
	if File_npool_smoketest_mw_v1_testplan_testplan_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_npool_smoketest_mw_v1_testplan_testplan_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateTestPlanRequest); i {
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
		file_npool_smoketest_mw_v1_testplan_testplan_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateTestPlanResponse); i {
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
		file_npool_smoketest_mw_v1_testplan_testplan_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateTestPlanRequest); i {
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
		file_npool_smoketest_mw_v1_testplan_testplan_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateTestPlanResponse); i {
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
		file_npool_smoketest_mw_v1_testplan_testplan_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTestPlanRequest); i {
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
		file_npool_smoketest_mw_v1_testplan_testplan_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTestPlanResponse); i {
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
		file_npool_smoketest_mw_v1_testplan_testplan_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTestPlansRequest); i {
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
		file_npool_smoketest_mw_v1_testplan_testplan_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTestPlansResponse); i {
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
		file_npool_smoketest_mw_v1_testplan_testplan_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExistTestPlanRequest); i {
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
		file_npool_smoketest_mw_v1_testplan_testplan_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExistTestPlanResponse); i {
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
		file_npool_smoketest_mw_v1_testplan_testplan_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteTestPlanRequest); i {
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
		file_npool_smoketest_mw_v1_testplan_testplan_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteTestPlanResponse); i {
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
			RawDescriptor: file_npool_smoketest_mw_v1_testplan_testplan_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_npool_smoketest_mw_v1_testplan_testplan_proto_goTypes,
		DependencyIndexes: file_npool_smoketest_mw_v1_testplan_testplan_proto_depIdxs,
		MessageInfos:      file_npool_smoketest_mw_v1_testplan_testplan_proto_msgTypes,
	}.Build()
	File_npool_smoketest_mw_v1_testplan_testplan_proto = out.File
	file_npool_smoketest_mw_v1_testplan_testplan_proto_rawDesc = nil
	file_npool_smoketest_mw_v1_testplan_testplan_proto_goTypes = nil
	file_npool_smoketest_mw_v1_testplan_testplan_proto_depIdxs = nil
}
