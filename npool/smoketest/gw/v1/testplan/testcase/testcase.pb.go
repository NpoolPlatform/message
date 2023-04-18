// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.18.1
// source: npool/smoketest/gw/v1/testplan/testcase/testcase.proto

package testcase

import (
	testcase "github.com/NpoolPlatform/message/npool/smoketest/mgr/v1/testplan/testcase"
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

type CreatePlanTestCaseRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TestPlanID     string                  `protobuf:"bytes,10,opt,name=TestPlanID,proto3" json:"TestPlanID,omitempty"`
	TestCaseID     string                  `protobuf:"bytes,20,opt,name=TestCaseID,proto3" json:"TestCaseID,omitempty"`
	TestCaseOutput *string                 `protobuf:"bytes,30,opt,name=TestCaseOutput,proto3,oneof" json:"TestCaseOutput,omitempty"`
	Description    *string                 `protobuf:"bytes,40,opt,name=Description,proto3,oneof" json:"Description,omitempty"`
	RunDuration    *uint32                 `protobuf:"varint,50,opt,name=RunDuration,proto3,oneof" json:"RunDuration,omitempty"`
	TestUserID     *string                 `protobuf:"bytes,60,opt,name=TestUserID,proto3,oneof" json:"TestUserID,omitempty"`
	Result         testcase.TestCaseResult `protobuf:"varint,70,opt,name=Result,proto3,enum=smoketest.manager.testplan.testcase.v1.TestCaseResult" json:"Result,omitempty"`
	Index          *uint32                 `protobuf:"varint,80,opt,name=Index,proto3,oneof" json:"Index,omitempty"`
}

func (x *CreatePlanTestCaseRequest) Reset() {
	*x = CreatePlanTestCaseRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_smoketest_gw_v1_testplan_testcase_testcase_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePlanTestCaseRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePlanTestCaseRequest) ProtoMessage() {}

func (x *CreatePlanTestCaseRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_smoketest_gw_v1_testplan_testcase_testcase_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePlanTestCaseRequest.ProtoReflect.Descriptor instead.
func (*CreatePlanTestCaseRequest) Descriptor() ([]byte, []int) {
	return file_npool_smoketest_gw_v1_testplan_testcase_testcase_proto_rawDescGZIP(), []int{0}
}

func (x *CreatePlanTestCaseRequest) GetTestPlanID() string {
	if x != nil {
		return x.TestPlanID
	}
	return ""
}

func (x *CreatePlanTestCaseRequest) GetTestCaseID() string {
	if x != nil {
		return x.TestCaseID
	}
	return ""
}

func (x *CreatePlanTestCaseRequest) GetTestCaseOutput() string {
	if x != nil && x.TestCaseOutput != nil {
		return *x.TestCaseOutput
	}
	return ""
}

func (x *CreatePlanTestCaseRequest) GetDescription() string {
	if x != nil && x.Description != nil {
		return *x.Description
	}
	return ""
}

func (x *CreatePlanTestCaseRequest) GetRunDuration() uint32 {
	if x != nil && x.RunDuration != nil {
		return *x.RunDuration
	}
	return 0
}

func (x *CreatePlanTestCaseRequest) GetTestUserID() string {
	if x != nil && x.TestUserID != nil {
		return *x.TestUserID
	}
	return ""
}

func (x *CreatePlanTestCaseRequest) GetResult() testcase.TestCaseResult {
	if x != nil {
		return x.Result
	}
	return testcase.TestCaseResult(0)
}

func (x *CreatePlanTestCaseRequest) GetIndex() uint32 {
	if x != nil && x.Index != nil {
		return *x.Index
	}
	return 0
}

type CreatePlanTestCaseResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *testcase.PlanTestCase `protobuf:"bytes,10,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (x *CreatePlanTestCaseResponse) Reset() {
	*x = CreatePlanTestCaseResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_smoketest_gw_v1_testplan_testcase_testcase_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePlanTestCaseResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePlanTestCaseResponse) ProtoMessage() {}

func (x *CreatePlanTestCaseResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_smoketest_gw_v1_testplan_testcase_testcase_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePlanTestCaseResponse.ProtoReflect.Descriptor instead.
func (*CreatePlanTestCaseResponse) Descriptor() ([]byte, []int) {
	return file_npool_smoketest_gw_v1_testplan_testcase_testcase_proto_rawDescGZIP(), []int{1}
}

func (x *CreatePlanTestCaseResponse) GetInfo() *testcase.PlanTestCase {
	if x != nil {
		return x.Info
	}
	return nil
}

type GetPlanTestCasesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TestPlanID string `protobuf:"bytes,10,opt,name=TestPlanID,proto3" json:"TestPlanID,omitempty"`
	Offset     *int32 `protobuf:"varint,20,opt,name=Offset,proto3,oneof" json:"Offset,omitempty"`
	Limit      *int32 `protobuf:"varint,30,opt,name=Limit,proto3,oneof" json:"Limit,omitempty"`
}

func (x *GetPlanTestCasesRequest) Reset() {
	*x = GetPlanTestCasesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_smoketest_gw_v1_testplan_testcase_testcase_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPlanTestCasesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPlanTestCasesRequest) ProtoMessage() {}

func (x *GetPlanTestCasesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_smoketest_gw_v1_testplan_testcase_testcase_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPlanTestCasesRequest.ProtoReflect.Descriptor instead.
func (*GetPlanTestCasesRequest) Descriptor() ([]byte, []int) {
	return file_npool_smoketest_gw_v1_testplan_testcase_testcase_proto_rawDescGZIP(), []int{2}
}

func (x *GetPlanTestCasesRequest) GetTestPlanID() string {
	if x != nil {
		return x.TestPlanID
	}
	return ""
}

func (x *GetPlanTestCasesRequest) GetOffset() int32 {
	if x != nil && x.Offset != nil {
		return *x.Offset
	}
	return 0
}

func (x *GetPlanTestCasesRequest) GetLimit() int32 {
	if x != nil && x.Limit != nil {
		return *x.Limit
	}
	return 0
}

type GetPlanTestCasesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Infos []*testcase.PlanTestCase `protobuf:"bytes,10,rep,name=Infos,proto3" json:"Infos,omitempty"`
	Total uint32                   `protobuf:"varint,20,opt,name=Total,proto3" json:"Total,omitempty"`
}

func (x *GetPlanTestCasesResponse) Reset() {
	*x = GetPlanTestCasesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_smoketest_gw_v1_testplan_testcase_testcase_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPlanTestCasesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPlanTestCasesResponse) ProtoMessage() {}

func (x *GetPlanTestCasesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_smoketest_gw_v1_testplan_testcase_testcase_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPlanTestCasesResponse.ProtoReflect.Descriptor instead.
func (*GetPlanTestCasesResponse) Descriptor() ([]byte, []int) {
	return file_npool_smoketest_gw_v1_testplan_testcase_testcase_proto_rawDescGZIP(), []int{3}
}

func (x *GetPlanTestCasesResponse) GetInfos() []*testcase.PlanTestCase {
	if x != nil {
		return x.Infos
	}
	return nil
}

func (x *GetPlanTestCasesResponse) GetTotal() uint32 {
	if x != nil {
		return x.Total
	}
	return 0
}

type GetPlanTestCaseRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID string `protobuf:"bytes,10,opt,name=ID,proto3" json:"ID,omitempty"`
}

func (x *GetPlanTestCaseRequest) Reset() {
	*x = GetPlanTestCaseRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_smoketest_gw_v1_testplan_testcase_testcase_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPlanTestCaseRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPlanTestCaseRequest) ProtoMessage() {}

func (x *GetPlanTestCaseRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_smoketest_gw_v1_testplan_testcase_testcase_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPlanTestCaseRequest.ProtoReflect.Descriptor instead.
func (*GetPlanTestCaseRequest) Descriptor() ([]byte, []int) {
	return file_npool_smoketest_gw_v1_testplan_testcase_testcase_proto_rawDescGZIP(), []int{4}
}

func (x *GetPlanTestCaseRequest) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

type GetPlanTestCaseResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *testcase.PlanTestCase `protobuf:"bytes,10,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (x *GetPlanTestCaseResponse) Reset() {
	*x = GetPlanTestCaseResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_smoketest_gw_v1_testplan_testcase_testcase_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPlanTestCaseResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPlanTestCaseResponse) ProtoMessage() {}

func (x *GetPlanTestCaseResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_smoketest_gw_v1_testplan_testcase_testcase_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPlanTestCaseResponse.ProtoReflect.Descriptor instead.
func (*GetPlanTestCaseResponse) Descriptor() ([]byte, []int) {
	return file_npool_smoketest_gw_v1_testplan_testcase_testcase_proto_rawDescGZIP(), []int{5}
}

func (x *GetPlanTestCaseResponse) GetInfo() *testcase.PlanTestCase {
	if x != nil {
		return x.Info
	}
	return nil
}

type DeletePlanTestCaseRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID string `protobuf:"bytes,10,opt,name=ID,proto3" json:"ID,omitempty"`
}

func (x *DeletePlanTestCaseRequest) Reset() {
	*x = DeletePlanTestCaseRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_smoketest_gw_v1_testplan_testcase_testcase_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeletePlanTestCaseRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeletePlanTestCaseRequest) ProtoMessage() {}

func (x *DeletePlanTestCaseRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_smoketest_gw_v1_testplan_testcase_testcase_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeletePlanTestCaseRequest.ProtoReflect.Descriptor instead.
func (*DeletePlanTestCaseRequest) Descriptor() ([]byte, []int) {
	return file_npool_smoketest_gw_v1_testplan_testcase_testcase_proto_rawDescGZIP(), []int{6}
}

func (x *DeletePlanTestCaseRequest) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

type DeletePlanTestCaseResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *testcase.PlanTestCase `protobuf:"bytes,10,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (x *DeletePlanTestCaseResponse) Reset() {
	*x = DeletePlanTestCaseResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_smoketest_gw_v1_testplan_testcase_testcase_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeletePlanTestCaseResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeletePlanTestCaseResponse) ProtoMessage() {}

func (x *DeletePlanTestCaseResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_smoketest_gw_v1_testplan_testcase_testcase_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeletePlanTestCaseResponse.ProtoReflect.Descriptor instead.
func (*DeletePlanTestCaseResponse) Descriptor() ([]byte, []int) {
	return file_npool_smoketest_gw_v1_testplan_testcase_testcase_proto_rawDescGZIP(), []int{7}
}

func (x *DeletePlanTestCaseResponse) GetInfo() *testcase.PlanTestCase {
	if x != nil {
		return x.Info
	}
	return nil
}

type DeletePlanTestCasesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TestPlanID string `protobuf:"bytes,20,opt,name=TestPlanID,proto3" json:"TestPlanID,omitempty"`
}

func (x *DeletePlanTestCasesRequest) Reset() {
	*x = DeletePlanTestCasesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_smoketest_gw_v1_testplan_testcase_testcase_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeletePlanTestCasesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeletePlanTestCasesRequest) ProtoMessage() {}

func (x *DeletePlanTestCasesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_smoketest_gw_v1_testplan_testcase_testcase_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeletePlanTestCasesRequest.ProtoReflect.Descriptor instead.
func (*DeletePlanTestCasesRequest) Descriptor() ([]byte, []int) {
	return file_npool_smoketest_gw_v1_testplan_testcase_testcase_proto_rawDescGZIP(), []int{8}
}

func (x *DeletePlanTestCasesRequest) GetTestPlanID() string {
	if x != nil {
		return x.TestPlanID
	}
	return ""
}

type DeletePlanTestCasesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info []*testcase.PlanTestCase `protobuf:"bytes,10,rep,name=Info,proto3" json:"Info,omitempty"`
}

func (x *DeletePlanTestCasesResponse) Reset() {
	*x = DeletePlanTestCasesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_smoketest_gw_v1_testplan_testcase_testcase_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeletePlanTestCasesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeletePlanTestCasesResponse) ProtoMessage() {}

func (x *DeletePlanTestCasesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_smoketest_gw_v1_testplan_testcase_testcase_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeletePlanTestCasesResponse.ProtoReflect.Descriptor instead.
func (*DeletePlanTestCasesResponse) Descriptor() ([]byte, []int) {
	return file_npool_smoketest_gw_v1_testplan_testcase_testcase_proto_rawDescGZIP(), []int{9}
}

func (x *DeletePlanTestCasesResponse) GetInfo() []*testcase.PlanTestCase {
	if x != nil {
		return x.Info
	}
	return nil
}

var File_npool_smoketest_gw_v1_testplan_testcase_testcase_proto protoreflect.FileDescriptor

var file_npool_smoketest_gw_v1_testplan_testcase_testcase_proto_rawDesc = []byte{
	0x0a, 0x36, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x73, 0x6d, 0x6f, 0x6b, 0x65, 0x74, 0x65, 0x73,
	0x74, 0x2f, 0x67, 0x77, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x70, 0x6c, 0x61, 0x6e,
	0x2f, 0x74, 0x65, 0x73, 0x74, 0x63, 0x61, 0x73, 0x65, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x63, 0x61,
	0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x26, 0x73, 0x6d, 0x6f, 0x6b, 0x65, 0x74,
	0x65, 0x73, 0x74, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x74, 0x65, 0x73, 0x74,
	0x70, 0x6c, 0x61, 0x6e, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x63, 0x61, 0x73, 0x65, 0x2e, 0x76, 0x31,
	0x1a, 0x37, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x73, 0x6d, 0x6f, 0x6b, 0x65, 0x74, 0x65, 0x73,
	0x74, 0x2f, 0x6d, 0x67, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x70, 0x6c, 0x61,
	0x6e, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x63, 0x61, 0x73, 0x65, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x63,
	0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb2, 0x03, 0x0a, 0x19, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x50, 0x6c, 0x61, 0x6e, 0x54, 0x65, 0x73, 0x74, 0x43, 0x61, 0x73, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x54, 0x65, 0x73, 0x74, 0x50, 0x6c, 0x61,
	0x6e, 0x49, 0x44, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x54, 0x65, 0x73, 0x74, 0x50,
	0x6c, 0x61, 0x6e, 0x49, 0x44, 0x12, 0x1e, 0x0a, 0x0a, 0x54, 0x65, 0x73, 0x74, 0x43, 0x61, 0x73,
	0x65, 0x49, 0x44, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x54, 0x65, 0x73, 0x74, 0x43,
	0x61, 0x73, 0x65, 0x49, 0x44, 0x12, 0x2b, 0x0a, 0x0e, 0x54, 0x65, 0x73, 0x74, 0x43, 0x61, 0x73,
	0x65, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52,
	0x0e, 0x54, 0x65, 0x73, 0x74, 0x43, 0x61, 0x73, 0x65, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x88,
	0x01, 0x01, 0x12, 0x25, 0x0a, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x28, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x88, 0x01, 0x01, 0x12, 0x25, 0x0a, 0x0b, 0x52, 0x75, 0x6e,
	0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x32, 0x20, 0x01, 0x28, 0x0d, 0x48, 0x02,
	0x52, 0x0b, 0x52, 0x75, 0x6e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x88, 0x01, 0x01,
	0x12, 0x23, 0x0a, 0x0a, 0x54, 0x65, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x3c,
	0x20, 0x01, 0x28, 0x09, 0x48, 0x03, 0x52, 0x0a, 0x54, 0x65, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72,
	0x49, 0x44, 0x88, 0x01, 0x01, 0x12, 0x4e, 0x0a, 0x06, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18,
	0x46, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x36, 0x2e, 0x73, 0x6d, 0x6f, 0x6b, 0x65, 0x74, 0x65, 0x73,
	0x74, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x70, 0x6c,
	0x61, 0x6e, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x63, 0x61, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x54,
	0x65, 0x73, 0x74, 0x43, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x06, 0x52,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x19, 0x0a, 0x05, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x50,
	0x20, 0x01, 0x28, 0x0d, 0x48, 0x04, 0x52, 0x05, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x88, 0x01, 0x01,
	0x42, 0x11, 0x0a, 0x0f, 0x5f, 0x54, 0x65, 0x73, 0x74, 0x43, 0x61, 0x73, 0x65, 0x4f, 0x75, 0x74,
	0x70, 0x75, 0x74, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x52, 0x75, 0x6e, 0x44, 0x75, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x54, 0x65, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72,
	0x49, 0x44, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x22, 0x66, 0x0a, 0x1a,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x6c, 0x61, 0x6e, 0x54, 0x65, 0x73, 0x74, 0x43, 0x61,
	0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x48, 0x0a, 0x04, 0x49, 0x6e,
	0x66, 0x6f, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x34, 0x2e, 0x73, 0x6d, 0x6f, 0x6b, 0x65,
	0x74, 0x65, 0x73, 0x74, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x74, 0x65, 0x73,
	0x74, 0x70, 0x6c, 0x61, 0x6e, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x63, 0x61, 0x73, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x50, 0x6c, 0x61, 0x6e, 0x54, 0x65, 0x73, 0x74, 0x43, 0x61, 0x73, 0x65, 0x52, 0x04,
	0x49, 0x6e, 0x66, 0x6f, 0x22, 0x86, 0x01, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x50, 0x6c, 0x61, 0x6e,
	0x54, 0x65, 0x73, 0x74, 0x43, 0x61, 0x73, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1e, 0x0a, 0x0a, 0x54, 0x65, 0x73, 0x74, 0x50, 0x6c, 0x61, 0x6e, 0x49, 0x44, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x54, 0x65, 0x73, 0x74, 0x50, 0x6c, 0x61, 0x6e, 0x49, 0x44,
	0x12, 0x1b, 0x0a, 0x06, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x14, 0x20, 0x01, 0x28, 0x05,
	0x48, 0x00, 0x52, 0x06, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x88, 0x01, 0x01, 0x12, 0x19, 0x0a,
	0x05, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x05, 0x48, 0x01, 0x52, 0x05,
	0x4c, 0x69, 0x6d, 0x69, 0x74, 0x88, 0x01, 0x01, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x4f, 0x66, 0x66,
	0x73, 0x65, 0x74, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x22, 0x7c, 0x0a,
	0x18, 0x47, 0x65, 0x74, 0x50, 0x6c, 0x61, 0x6e, 0x54, 0x65, 0x73, 0x74, 0x43, 0x61, 0x73, 0x65,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4a, 0x0a, 0x05, 0x49, 0x6e, 0x66,
	0x6f, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x34, 0x2e, 0x73, 0x6d, 0x6f, 0x6b, 0x65,
	0x74, 0x65, 0x73, 0x74, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x74, 0x65, 0x73,
	0x74, 0x70, 0x6c, 0x61, 0x6e, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x63, 0x61, 0x73, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x50, 0x6c, 0x61, 0x6e, 0x54, 0x65, 0x73, 0x74, 0x43, 0x61, 0x73, 0x65, 0x52, 0x05,
	0x49, 0x6e, 0x66, 0x6f, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x14,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x22, 0x28, 0x0a, 0x16, 0x47,
	0x65, 0x74, 0x50, 0x6c, 0x61, 0x6e, 0x54, 0x65, 0x73, 0x74, 0x43, 0x61, 0x73, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x49, 0x44, 0x22, 0x63, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x50, 0x6c, 0x61, 0x6e,
	0x54, 0x65, 0x73, 0x74, 0x43, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x48, 0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x34,
	0x2e, 0x73, 0x6d, 0x6f, 0x6b, 0x65, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67,
	0x65, 0x72, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x70, 0x6c, 0x61, 0x6e, 0x2e, 0x74, 0x65, 0x73, 0x74,
	0x63, 0x61, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6c, 0x61, 0x6e, 0x54, 0x65, 0x73, 0x74,
	0x43, 0x61, 0x73, 0x65, 0x52, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x2b, 0x0a, 0x19, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x50, 0x6c, 0x61, 0x6e, 0x54, 0x65, 0x73, 0x74, 0x43, 0x61, 0x73, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x22, 0x66, 0x0a, 0x1a, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x50, 0x6c, 0x61, 0x6e, 0x54, 0x65, 0x73, 0x74, 0x43, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x48, 0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x34, 0x2e, 0x73, 0x6d, 0x6f, 0x6b, 0x65, 0x74, 0x65, 0x73, 0x74, 0x2e,
	0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x70, 0x6c, 0x61, 0x6e,
	0x2e, 0x74, 0x65, 0x73, 0x74, 0x63, 0x61, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6c, 0x61,
	0x6e, 0x54, 0x65, 0x73, 0x74, 0x43, 0x61, 0x73, 0x65, 0x52, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x22,
	0x3c, 0x0a, 0x1a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x6c, 0x61, 0x6e, 0x54, 0x65, 0x73,
	0x74, 0x43, 0x61, 0x73, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a,
	0x0a, 0x54, 0x65, 0x73, 0x74, 0x50, 0x6c, 0x61, 0x6e, 0x49, 0x44, 0x18, 0x14, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x54, 0x65, 0x73, 0x74, 0x50, 0x6c, 0x61, 0x6e, 0x49, 0x44, 0x22, 0x67, 0x0a,
	0x1b, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x6c, 0x61, 0x6e, 0x54, 0x65, 0x73, 0x74, 0x43,
	0x61, 0x73, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x48, 0x0a, 0x04,
	0x49, 0x6e, 0x66, 0x6f, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x34, 0x2e, 0x73, 0x6d, 0x6f,
	0x6b, 0x65, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x74,
	0x65, 0x73, 0x74, 0x70, 0x6c, 0x61, 0x6e, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x63, 0x61, 0x73, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6c, 0x61, 0x6e, 0x54, 0x65, 0x73, 0x74, 0x43, 0x61, 0x73, 0x65,
	0x52, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x32, 0xd3, 0x04, 0x0a, 0x07, 0x47, 0x61, 0x74, 0x65, 0x77,
	0x61, 0x79, 0x12, 0xbf, 0x01, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x6c, 0x61,
	0x6e, 0x54, 0x65, 0x73, 0x74, 0x43, 0x61, 0x73, 0x65, 0x12, 0x41, 0x2e, 0x73, 0x6d, 0x6f, 0x6b,
	0x65, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x74, 0x65,
	0x73, 0x74, 0x70, 0x6c, 0x61, 0x6e, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x63, 0x61, 0x73, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x6c, 0x61, 0x6e, 0x54, 0x65, 0x73,
	0x74, 0x43, 0x61, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x42, 0x2e, 0x73,
	0x6d, 0x6f, 0x6b, 0x65, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79,
	0x2e, 0x74, 0x65, 0x73, 0x74, 0x70, 0x6c, 0x61, 0x6e, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x63, 0x61,
	0x73, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x6c, 0x61, 0x6e,
	0x54, 0x65, 0x73, 0x74, 0x43, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x22, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1c, 0x22, 0x17, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x65, 0x73, 0x74, 0x63, 0x61, 0x73,
	0x65, 0x3a, 0x01, 0x2a, 0x12, 0xbf, 0x01, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50,
	0x6c, 0x61, 0x6e, 0x54, 0x65, 0x73, 0x74, 0x43, 0x61, 0x73, 0x65, 0x12, 0x41, 0x2e, 0x73, 0x6d,
	0x6f, 0x6b, 0x65, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e,
	0x74, 0x65, 0x73, 0x74, 0x70, 0x6c, 0x61, 0x6e, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x63, 0x61, 0x73,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x6c, 0x61, 0x6e, 0x54,
	0x65, 0x73, 0x74, 0x43, 0x61, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x42,
	0x2e, 0x73, 0x6d, 0x6f, 0x6b, 0x65, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77,
	0x61, 0x79, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x70, 0x6c, 0x61, 0x6e, 0x2e, 0x74, 0x65, 0x73, 0x74,
	0x63, 0x61, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x6c,
	0x61, 0x6e, 0x54, 0x65, 0x73, 0x74, 0x43, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x22, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1c, 0x22, 0x17, 0x2f, 0x76, 0x31, 0x2f,
	0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x65, 0x73, 0x74, 0x63,
	0x61, 0x73, 0x65, 0x3a, 0x01, 0x2a, 0x12, 0xc3, 0x01, 0x0a, 0x13, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x50, 0x6c, 0x61, 0x6e, 0x54, 0x65, 0x73, 0x74, 0x43, 0x61, 0x73, 0x65, 0x73, 0x12, 0x42,
	0x2e, 0x73, 0x6d, 0x6f, 0x6b, 0x65, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77,
	0x61, 0x79, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x70, 0x6c, 0x61, 0x6e, 0x2e, 0x74, 0x65, 0x73, 0x74,
	0x63, 0x61, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x6c,
	0x61, 0x6e, 0x54, 0x65, 0x73, 0x74, 0x43, 0x61, 0x73, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x43, 0x2e, 0x73, 0x6d, 0x6f, 0x6b, 0x65, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x67,
	0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x70, 0x6c, 0x61, 0x6e, 0x2e,
	0x74, 0x65, 0x73, 0x74, 0x63, 0x61, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x50, 0x6c, 0x61, 0x6e, 0x54, 0x65, 0x73, 0x74, 0x43, 0x61, 0x73, 0x65, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x23, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1d, 0x22,
	0x18, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x2f, 0x70, 0x6c, 0x61, 0x6e,
	0x74, 0x65, 0x73, 0x74, 0x63, 0x61, 0x73, 0x65, 0x73, 0x3a, 0x01, 0x2a, 0x42, 0x4a, 0x5a, 0x48,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4e, 0x70, 0x6f, 0x6f, 0x6c,
	0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x2f, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x73, 0x6d, 0x6f, 0x6b, 0x65, 0x74, 0x65, 0x73, 0x74,
	0x2f, 0x67, 0x77, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x70, 0x6c, 0x61, 0x6e, 0x2f,
	0x74, 0x65, 0x73, 0x74, 0x63, 0x61, 0x73, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_npool_smoketest_gw_v1_testplan_testcase_testcase_proto_rawDescOnce sync.Once
	file_npool_smoketest_gw_v1_testplan_testcase_testcase_proto_rawDescData = file_npool_smoketest_gw_v1_testplan_testcase_testcase_proto_rawDesc
)

func file_npool_smoketest_gw_v1_testplan_testcase_testcase_proto_rawDescGZIP() []byte {
	file_npool_smoketest_gw_v1_testplan_testcase_testcase_proto_rawDescOnce.Do(func() {
		file_npool_smoketest_gw_v1_testplan_testcase_testcase_proto_rawDescData = protoimpl.X.CompressGZIP(file_npool_smoketest_gw_v1_testplan_testcase_testcase_proto_rawDescData)
	})
	return file_npool_smoketest_gw_v1_testplan_testcase_testcase_proto_rawDescData
}

var file_npool_smoketest_gw_v1_testplan_testcase_testcase_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_npool_smoketest_gw_v1_testplan_testcase_testcase_proto_goTypes = []interface{}{
	(*CreatePlanTestCaseRequest)(nil),   // 0: smoketest.gateway.testplan.testcase.v1.CreatePlanTestCaseRequest
	(*CreatePlanTestCaseResponse)(nil),  // 1: smoketest.gateway.testplan.testcase.v1.CreatePlanTestCaseResponse
	(*GetPlanTestCasesRequest)(nil),     // 2: smoketest.gateway.testplan.testcase.v1.GetPlanTestCasesRequest
	(*GetPlanTestCasesResponse)(nil),    // 3: smoketest.gateway.testplan.testcase.v1.GetPlanTestCasesResponse
	(*GetPlanTestCaseRequest)(nil),      // 4: smoketest.gateway.testplan.testcase.v1.GetPlanTestCaseRequest
	(*GetPlanTestCaseResponse)(nil),     // 5: smoketest.gateway.testplan.testcase.v1.GetPlanTestCaseResponse
	(*DeletePlanTestCaseRequest)(nil),   // 6: smoketest.gateway.testplan.testcase.v1.DeletePlanTestCaseRequest
	(*DeletePlanTestCaseResponse)(nil),  // 7: smoketest.gateway.testplan.testcase.v1.DeletePlanTestCaseResponse
	(*DeletePlanTestCasesRequest)(nil),  // 8: smoketest.gateway.testplan.testcase.v1.DeletePlanTestCasesRequest
	(*DeletePlanTestCasesResponse)(nil), // 9: smoketest.gateway.testplan.testcase.v1.DeletePlanTestCasesResponse
	(testcase.TestCaseResult)(0),        // 10: smoketest.manager.testplan.testcase.v1.TestCaseResult
	(*testcase.PlanTestCase)(nil),       // 11: smoketest.manager.testplan.testcase.v1.PlanTestCase
}
var file_npool_smoketest_gw_v1_testplan_testcase_testcase_proto_depIdxs = []int32{
	10, // 0: smoketest.gateway.testplan.testcase.v1.CreatePlanTestCaseRequest.Result:type_name -> smoketest.manager.testplan.testcase.v1.TestCaseResult
	11, // 1: smoketest.gateway.testplan.testcase.v1.CreatePlanTestCaseResponse.Info:type_name -> smoketest.manager.testplan.testcase.v1.PlanTestCase
	11, // 2: smoketest.gateway.testplan.testcase.v1.GetPlanTestCasesResponse.Infos:type_name -> smoketest.manager.testplan.testcase.v1.PlanTestCase
	11, // 3: smoketest.gateway.testplan.testcase.v1.GetPlanTestCaseResponse.Info:type_name -> smoketest.manager.testplan.testcase.v1.PlanTestCase
	11, // 4: smoketest.gateway.testplan.testcase.v1.DeletePlanTestCaseResponse.Info:type_name -> smoketest.manager.testplan.testcase.v1.PlanTestCase
	11, // 5: smoketest.gateway.testplan.testcase.v1.DeletePlanTestCasesResponse.Info:type_name -> smoketest.manager.testplan.testcase.v1.PlanTestCase
	0,  // 6: smoketest.gateway.testplan.testcase.v1.Gateway.CreatePlanTestCase:input_type -> smoketest.gateway.testplan.testcase.v1.CreatePlanTestCaseRequest
	6,  // 7: smoketest.gateway.testplan.testcase.v1.Gateway.DeletePlanTestCase:input_type -> smoketest.gateway.testplan.testcase.v1.DeletePlanTestCaseRequest
	8,  // 8: smoketest.gateway.testplan.testcase.v1.Gateway.DeletePlanTestCases:input_type -> smoketest.gateway.testplan.testcase.v1.DeletePlanTestCasesRequest
	1,  // 9: smoketest.gateway.testplan.testcase.v1.Gateway.CreatePlanTestCase:output_type -> smoketest.gateway.testplan.testcase.v1.CreatePlanTestCaseResponse
	7,  // 10: smoketest.gateway.testplan.testcase.v1.Gateway.DeletePlanTestCase:output_type -> smoketest.gateway.testplan.testcase.v1.DeletePlanTestCaseResponse
	9,  // 11: smoketest.gateway.testplan.testcase.v1.Gateway.DeletePlanTestCases:output_type -> smoketest.gateway.testplan.testcase.v1.DeletePlanTestCasesResponse
	9,  // [9:12] is the sub-list for method output_type
	6,  // [6:9] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_npool_smoketest_gw_v1_testplan_testcase_testcase_proto_init() }
func file_npool_smoketest_gw_v1_testplan_testcase_testcase_proto_init() {
	if File_npool_smoketest_gw_v1_testplan_testcase_testcase_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_npool_smoketest_gw_v1_testplan_testcase_testcase_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePlanTestCaseRequest); i {
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
		file_npool_smoketest_gw_v1_testplan_testcase_testcase_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePlanTestCaseResponse); i {
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
		file_npool_smoketest_gw_v1_testplan_testcase_testcase_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPlanTestCasesRequest); i {
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
		file_npool_smoketest_gw_v1_testplan_testcase_testcase_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPlanTestCasesResponse); i {
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
		file_npool_smoketest_gw_v1_testplan_testcase_testcase_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPlanTestCaseRequest); i {
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
		file_npool_smoketest_gw_v1_testplan_testcase_testcase_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPlanTestCaseResponse); i {
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
		file_npool_smoketest_gw_v1_testplan_testcase_testcase_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeletePlanTestCaseRequest); i {
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
		file_npool_smoketest_gw_v1_testplan_testcase_testcase_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeletePlanTestCaseResponse); i {
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
		file_npool_smoketest_gw_v1_testplan_testcase_testcase_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeletePlanTestCasesRequest); i {
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
		file_npool_smoketest_gw_v1_testplan_testcase_testcase_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeletePlanTestCasesResponse); i {
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
	file_npool_smoketest_gw_v1_testplan_testcase_testcase_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_npool_smoketest_gw_v1_testplan_testcase_testcase_proto_msgTypes[2].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_npool_smoketest_gw_v1_testplan_testcase_testcase_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_npool_smoketest_gw_v1_testplan_testcase_testcase_proto_goTypes,
		DependencyIndexes: file_npool_smoketest_gw_v1_testplan_testcase_testcase_proto_depIdxs,
		MessageInfos:      file_npool_smoketest_gw_v1_testplan_testcase_testcase_proto_msgTypes,
	}.Build()
	File_npool_smoketest_gw_v1_testplan_testcase_testcase_proto = out.File
	file_npool_smoketest_gw_v1_testplan_testcase_testcase_proto_rawDesc = nil
	file_npool_smoketest_gw_v1_testplan_testcase_testcase_proto_goTypes = nil
	file_npool_smoketest_gw_v1_testplan_testcase_testcase_proto_depIdxs = nil
}
