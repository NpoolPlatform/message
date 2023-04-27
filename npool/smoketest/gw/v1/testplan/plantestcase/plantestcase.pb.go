// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: npool/smoketest/gw/v1/testplan/plantestcase/plantestcase.proto

package plantestcase

import (
	plantestcase "github.com/NpoolPlatform/message/npool/smoketest/mw/v1/testplan/plantestcase"
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

	TestPlanID string  `protobuf:"bytes,10,opt,name=TestPlanID,proto3" json:"TestPlanID,omitempty"`
	TestCaseID string  `protobuf:"bytes,20,opt,name=TestCaseID,proto3" json:"TestCaseID,omitempty"`
	Index      *uint32 `protobuf:"varint,30,opt,name=Index,proto3,oneof" json:"Index,omitempty"`
}

func (x *CreatePlanTestCaseRequest) Reset() {
	*x = CreatePlanTestCaseRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_smoketest_gw_v1_testplan_plantestcase_plantestcase_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePlanTestCaseRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePlanTestCaseRequest) ProtoMessage() {}

func (x *CreatePlanTestCaseRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_smoketest_gw_v1_testplan_plantestcase_plantestcase_proto_msgTypes[0]
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
	return file_npool_smoketest_gw_v1_testplan_plantestcase_plantestcase_proto_rawDescGZIP(), []int{0}
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

	Info *plantestcase.PlanTestCase `protobuf:"bytes,10,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (x *CreatePlanTestCaseResponse) Reset() {
	*x = CreatePlanTestCaseResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_smoketest_gw_v1_testplan_plantestcase_plantestcase_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePlanTestCaseResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePlanTestCaseResponse) ProtoMessage() {}

func (x *CreatePlanTestCaseResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_smoketest_gw_v1_testplan_plantestcase_plantestcase_proto_msgTypes[1]
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
	return file_npool_smoketest_gw_v1_testplan_plantestcase_plantestcase_proto_rawDescGZIP(), []int{1}
}

func (x *CreatePlanTestCaseResponse) GetInfo() *plantestcase.PlanTestCase {
	if x != nil {
		return x.Info
	}
	return nil
}

type UpdatePlanTestCaseRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID             string                       `protobuf:"bytes,10,opt,name=ID,proto3" json:"ID,omitempty"`
	TestCaseOutput *string                      `protobuf:"bytes,20,opt,name=TestCaseOutput,proto3,oneof" json:"TestCaseOutput,omitempty"`
	Description    *string                      `protobuf:"bytes,30,opt,name=Description,proto3,oneof" json:"Description,omitempty"`
	RunDuration    *uint32                      `protobuf:"varint,40,opt,name=RunDuration,proto3,oneof" json:"RunDuration,omitempty"`
	TestUserID     *string                      `protobuf:"bytes,50,opt,name=TestUserID,proto3,oneof" json:"TestUserID,omitempty"`
	AppID          *string                      `protobuf:"bytes,60,opt,name=AppID,proto3,oneof" json:"AppID,omitempty"`
	Result         *plantestcase.TestCaseResult `protobuf:"varint,70,opt,name=Result,proto3,enum=smoketest.middleware.testplan.plantestcase.v1.TestCaseResult,oneof" json:"Result,omitempty"`
	Index          *uint32                      `protobuf:"varint,80,opt,name=Index,proto3,oneof" json:"Index,omitempty"`
}

func (x *UpdatePlanTestCaseRequest) Reset() {
	*x = UpdatePlanTestCaseRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_smoketest_gw_v1_testplan_plantestcase_plantestcase_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdatePlanTestCaseRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePlanTestCaseRequest) ProtoMessage() {}

func (x *UpdatePlanTestCaseRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_smoketest_gw_v1_testplan_plantestcase_plantestcase_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatePlanTestCaseRequest.ProtoReflect.Descriptor instead.
func (*UpdatePlanTestCaseRequest) Descriptor() ([]byte, []int) {
	return file_npool_smoketest_gw_v1_testplan_plantestcase_plantestcase_proto_rawDescGZIP(), []int{2}
}

func (x *UpdatePlanTestCaseRequest) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *UpdatePlanTestCaseRequest) GetTestCaseOutput() string {
	if x != nil && x.TestCaseOutput != nil {
		return *x.TestCaseOutput
	}
	return ""
}

func (x *UpdatePlanTestCaseRequest) GetDescription() string {
	if x != nil && x.Description != nil {
		return *x.Description
	}
	return ""
}

func (x *UpdatePlanTestCaseRequest) GetRunDuration() uint32 {
	if x != nil && x.RunDuration != nil {
		return *x.RunDuration
	}
	return 0
}

func (x *UpdatePlanTestCaseRequest) GetTestUserID() string {
	if x != nil && x.TestUserID != nil {
		return *x.TestUserID
	}
	return ""
}

func (x *UpdatePlanTestCaseRequest) GetAppID() string {
	if x != nil && x.AppID != nil {
		return *x.AppID
	}
	return ""
}

func (x *UpdatePlanTestCaseRequest) GetResult() plantestcase.TestCaseResult {
	if x != nil && x.Result != nil {
		return *x.Result
	}
	return plantestcase.TestCaseResult(0)
}

func (x *UpdatePlanTestCaseRequest) GetIndex() uint32 {
	if x != nil && x.Index != nil {
		return *x.Index
	}
	return 0
}

type UpdatePlanTestCaseResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *plantestcase.PlanTestCase `protobuf:"bytes,10,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (x *UpdatePlanTestCaseResponse) Reset() {
	*x = UpdatePlanTestCaseResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_smoketest_gw_v1_testplan_plantestcase_plantestcase_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdatePlanTestCaseResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePlanTestCaseResponse) ProtoMessage() {}

func (x *UpdatePlanTestCaseResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_smoketest_gw_v1_testplan_plantestcase_plantestcase_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatePlanTestCaseResponse.ProtoReflect.Descriptor instead.
func (*UpdatePlanTestCaseResponse) Descriptor() ([]byte, []int) {
	return file_npool_smoketest_gw_v1_testplan_plantestcase_plantestcase_proto_rawDescGZIP(), []int{3}
}

func (x *UpdatePlanTestCaseResponse) GetInfo() *plantestcase.PlanTestCase {
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
	Offset     int32  `protobuf:"varint,20,opt,name=Offset,proto3" json:"Offset,omitempty"`
	Limit      int32  `protobuf:"varint,30,opt,name=Limit,proto3" json:"Limit,omitempty"`
}

func (x *GetPlanTestCasesRequest) Reset() {
	*x = GetPlanTestCasesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_smoketest_gw_v1_testplan_plantestcase_plantestcase_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPlanTestCasesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPlanTestCasesRequest) ProtoMessage() {}

func (x *GetPlanTestCasesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_smoketest_gw_v1_testplan_plantestcase_plantestcase_proto_msgTypes[4]
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
	return file_npool_smoketest_gw_v1_testplan_plantestcase_plantestcase_proto_rawDescGZIP(), []int{4}
}

func (x *GetPlanTestCasesRequest) GetTestPlanID() string {
	if x != nil {
		return x.TestPlanID
	}
	return ""
}

func (x *GetPlanTestCasesRequest) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *GetPlanTestCasesRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type GetPlanTestCasesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Infos []*plantestcase.PlanTestCase `protobuf:"bytes,10,rep,name=Infos,proto3" json:"Infos,omitempty"`
	Total uint32                       `protobuf:"varint,20,opt,name=Total,proto3" json:"Total,omitempty"`
}

func (x *GetPlanTestCasesResponse) Reset() {
	*x = GetPlanTestCasesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_smoketest_gw_v1_testplan_plantestcase_plantestcase_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPlanTestCasesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPlanTestCasesResponse) ProtoMessage() {}

func (x *GetPlanTestCasesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_smoketest_gw_v1_testplan_plantestcase_plantestcase_proto_msgTypes[5]
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
	return file_npool_smoketest_gw_v1_testplan_plantestcase_plantestcase_proto_rawDescGZIP(), []int{5}
}

func (x *GetPlanTestCasesResponse) GetInfos() []*plantestcase.PlanTestCase {
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

type DeletePlanTestCaseRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID string `protobuf:"bytes,10,opt,name=ID,proto3" json:"ID,omitempty"`
}

func (x *DeletePlanTestCaseRequest) Reset() {
	*x = DeletePlanTestCaseRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_smoketest_gw_v1_testplan_plantestcase_plantestcase_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeletePlanTestCaseRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeletePlanTestCaseRequest) ProtoMessage() {}

func (x *DeletePlanTestCaseRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_smoketest_gw_v1_testplan_plantestcase_plantestcase_proto_msgTypes[6]
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
	return file_npool_smoketest_gw_v1_testplan_plantestcase_plantestcase_proto_rawDescGZIP(), []int{6}
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

	Info *plantestcase.PlanTestCase `protobuf:"bytes,10,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (x *DeletePlanTestCaseResponse) Reset() {
	*x = DeletePlanTestCaseResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_smoketest_gw_v1_testplan_plantestcase_plantestcase_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeletePlanTestCaseResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeletePlanTestCaseResponse) ProtoMessage() {}

func (x *DeletePlanTestCaseResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_smoketest_gw_v1_testplan_plantestcase_plantestcase_proto_msgTypes[7]
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
	return file_npool_smoketest_gw_v1_testplan_plantestcase_plantestcase_proto_rawDescGZIP(), []int{7}
}

func (x *DeletePlanTestCaseResponse) GetInfo() *plantestcase.PlanTestCase {
	if x != nil {
		return x.Info
	}
	return nil
}

var File_npool_smoketest_gw_v1_testplan_plantestcase_plantestcase_proto protoreflect.FileDescriptor

var file_npool_smoketest_gw_v1_testplan_plantestcase_plantestcase_proto_rawDesc = []byte{
	0x0a, 0x3e, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x73, 0x6d, 0x6f, 0x6b, 0x65, 0x74, 0x65, 0x73,
	0x74, 0x2f, 0x67, 0x77, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x70, 0x6c, 0x61, 0x6e,
	0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x65, 0x73, 0x74, 0x63, 0x61, 0x73, 0x65, 0x2f, 0x70, 0x6c,
	0x61, 0x6e, 0x74, 0x65, 0x73, 0x74, 0x63, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x2a, 0x73, 0x6d, 0x6f, 0x6b, 0x65, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x67, 0x61, 0x74, 0x65,
	0x77, 0x61, 0x79, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x70, 0x6c, 0x61, 0x6e, 0x2e, 0x70, 0x6c, 0x61,
	0x6e, 0x74, 0x65, 0x73, 0x74, 0x63, 0x61, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x3e, 0x6e, 0x70,
	0x6f, 0x6f, 0x6c, 0x2f, 0x73, 0x6d, 0x6f, 0x6b, 0x65, 0x74, 0x65, 0x73, 0x74, 0x2f, 0x6d, 0x77,
	0x2f, 0x76, 0x31, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x70, 0x6c, 0x61, 0x6e, 0x2f, 0x70, 0x6c, 0x61,
	0x6e, 0x74, 0x65, 0x73, 0x74, 0x63, 0x61, 0x73, 0x65, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x65,
	0x73, 0x74, 0x63, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x80, 0x01, 0x0a, 0x19, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x6c, 0x61, 0x6e, 0x54, 0x65, 0x73, 0x74, 0x43, 0x61, 0x73,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x54, 0x65, 0x73, 0x74,
	0x50, 0x6c, 0x61, 0x6e, 0x49, 0x44, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x54, 0x65,
	0x73, 0x74, 0x50, 0x6c, 0x61, 0x6e, 0x49, 0x44, 0x12, 0x1e, 0x0a, 0x0a, 0x54, 0x65, 0x73, 0x74,
	0x43, 0x61, 0x73, 0x65, 0x49, 0x44, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x54, 0x65,
	0x73, 0x74, 0x43, 0x61, 0x73, 0x65, 0x49, 0x44, 0x12, 0x19, 0x0a, 0x05, 0x49, 0x6e, 0x64, 0x65,
	0x78, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x0d, 0x48, 0x00, 0x52, 0x05, 0x49, 0x6e, 0x64, 0x65, 0x78,
	0x88, 0x01, 0x01, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x22, 0x6d, 0x0a,
	0x1a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x6c, 0x61, 0x6e, 0x54, 0x65, 0x73, 0x74, 0x43,
	0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4f, 0x0a, 0x04, 0x49,
	0x6e, 0x66, 0x6f, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3b, 0x2e, 0x73, 0x6d, 0x6f, 0x6b,
	0x65, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65,
	0x2e, 0x74, 0x65, 0x73, 0x74, 0x70, 0x6c, 0x61, 0x6e, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x65,
	0x73, 0x74, 0x63, 0x61, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6c, 0x61, 0x6e, 0x54, 0x65,
	0x73, 0x74, 0x43, 0x61, 0x73, 0x65, 0x52, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0xbe, 0x03, 0x0a,
	0x19, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x6c, 0x61, 0x6e, 0x54, 0x65, 0x73, 0x74, 0x43,
	0x61, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x12, 0x2b, 0x0a, 0x0e, 0x54, 0x65,
	0x73, 0x74, 0x43, 0x61, 0x73, 0x65, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x18, 0x14, 0x20, 0x01,
	0x28, 0x09, 0x48, 0x00, 0x52, 0x0e, 0x54, 0x65, 0x73, 0x74, 0x43, 0x61, 0x73, 0x65, 0x4f, 0x75,
	0x74, 0x70, 0x75, 0x74, 0x88, 0x01, 0x01, 0x12, 0x25, 0x0a, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x0b,
	0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x88, 0x01, 0x01, 0x12, 0x25,
	0x0a, 0x0b, 0x52, 0x75, 0x6e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x28, 0x20,
	0x01, 0x28, 0x0d, 0x48, 0x02, 0x52, 0x0b, 0x52, 0x75, 0x6e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x88, 0x01, 0x01, 0x12, 0x23, 0x0a, 0x0a, 0x54, 0x65, 0x73, 0x74, 0x55, 0x73, 0x65,
	0x72, 0x49, 0x44, 0x18, 0x32, 0x20, 0x01, 0x28, 0x09, 0x48, 0x03, 0x52, 0x0a, 0x54, 0x65, 0x73,
	0x74, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x88, 0x01, 0x01, 0x12, 0x19, 0x0a, 0x05, 0x41, 0x70,
	0x70, 0x49, 0x44, 0x18, 0x3c, 0x20, 0x01, 0x28, 0x09, 0x48, 0x04, 0x52, 0x05, 0x41, 0x70, 0x70,
	0x49, 0x44, 0x88, 0x01, 0x01, 0x12, 0x5a, 0x0a, 0x06, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18,
	0x46, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x3d, 0x2e, 0x73, 0x6d, 0x6f, 0x6b, 0x65, 0x74, 0x65, 0x73,
	0x74, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x74, 0x65, 0x73,
	0x74, 0x70, 0x6c, 0x61, 0x6e, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x65, 0x73, 0x74, 0x63, 0x61,
	0x73, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x43, 0x61, 0x73, 0x65, 0x52, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x48, 0x05, 0x52, 0x06, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x88, 0x01,
	0x01, 0x12, 0x19, 0x0a, 0x05, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x50, 0x20, 0x01, 0x28, 0x0d,
	0x48, 0x06, 0x52, 0x05, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x88, 0x01, 0x01, 0x42, 0x11, 0x0a, 0x0f,
	0x5f, 0x54, 0x65, 0x73, 0x74, 0x43, 0x61, 0x73, 0x65, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x42,
	0x0e, 0x0a, 0x0c, 0x5f, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x42,
	0x0e, 0x0a, 0x0c, 0x5f, 0x52, 0x75, 0x6e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42,
	0x0d, 0x0a, 0x0b, 0x5f, 0x54, 0x65, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x42, 0x08,
	0x0a, 0x06, 0x5f, 0x41, 0x70, 0x70, 0x49, 0x44, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x52, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x22, 0x6d, 0x0a,
	0x1a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x6c, 0x61, 0x6e, 0x54, 0x65, 0x73, 0x74, 0x43,
	0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4f, 0x0a, 0x04, 0x49,
	0x6e, 0x66, 0x6f, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3b, 0x2e, 0x73, 0x6d, 0x6f, 0x6b,
	0x65, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65,
	0x2e, 0x74, 0x65, 0x73, 0x74, 0x70, 0x6c, 0x61, 0x6e, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x65,
	0x73, 0x74, 0x63, 0x61, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6c, 0x61, 0x6e, 0x54, 0x65,
	0x73, 0x74, 0x43, 0x61, 0x73, 0x65, 0x52, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x67, 0x0a, 0x17,
	0x47, 0x65, 0x74, 0x50, 0x6c, 0x61, 0x6e, 0x54, 0x65, 0x73, 0x74, 0x43, 0x61, 0x73, 0x65, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x54, 0x65, 0x73, 0x74, 0x50,
	0x6c, 0x61, 0x6e, 0x49, 0x44, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x54, 0x65, 0x73,
	0x74, 0x50, 0x6c, 0x61, 0x6e, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x4f, 0x66, 0x66, 0x73, 0x65,
	0x74, 0x18, 0x14, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12,
	0x14, 0x0a, 0x05, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05,
	0x4c, 0x69, 0x6d, 0x69, 0x74, 0x22, 0x83, 0x01, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x50, 0x6c, 0x61,
	0x6e, 0x54, 0x65, 0x73, 0x74, 0x43, 0x61, 0x73, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x51, 0x0a, 0x05, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x3b, 0x2e, 0x73, 0x6d, 0x6f, 0x6b, 0x65, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x6d, 0x69,
	0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x70, 0x6c, 0x61,
	0x6e, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x65, 0x73, 0x74, 0x63, 0x61, 0x73, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x50, 0x6c, 0x61, 0x6e, 0x54, 0x65, 0x73, 0x74, 0x43, 0x61, 0x73, 0x65, 0x52, 0x05,
	0x49, 0x6e, 0x66, 0x6f, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x14,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x22, 0x2b, 0x0a, 0x19, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x6c, 0x61, 0x6e, 0x54, 0x65, 0x73, 0x74, 0x43, 0x61, 0x73,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x22, 0x6d, 0x0a, 0x1a, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x50, 0x6c, 0x61, 0x6e, 0x54, 0x65, 0x73, 0x74, 0x43, 0x61, 0x73, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4f, 0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x3b, 0x2e, 0x73, 0x6d, 0x6f, 0x6b, 0x65, 0x74, 0x65, 0x73, 0x74,
	0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x74, 0x65, 0x73, 0x74,
	0x70, 0x6c, 0x61, 0x6e, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x65, 0x73, 0x74, 0x63, 0x61, 0x73,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6c, 0x61, 0x6e, 0x54, 0x65, 0x73, 0x74, 0x43, 0x61, 0x73,
	0x65, 0x52, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x32, 0xa9, 0x06, 0x0a, 0x07, 0x47, 0x61, 0x74, 0x65,
	0x77, 0x61, 0x79, 0x12, 0xc7, 0x01, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x6c,
	0x61, 0x6e, 0x54, 0x65, 0x73, 0x74, 0x43, 0x61, 0x73, 0x65, 0x12, 0x45, 0x2e, 0x73, 0x6d, 0x6f,
	0x6b, 0x65, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x74,
	0x65, 0x73, 0x74, 0x70, 0x6c, 0x61, 0x6e, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x65, 0x73, 0x74,
	0x63, 0x61, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x6c,
	0x61, 0x6e, 0x54, 0x65, 0x73, 0x74, 0x43, 0x61, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x46, 0x2e, 0x73, 0x6d, 0x6f, 0x6b, 0x65, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x67, 0x61,
	0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x70, 0x6c, 0x61, 0x6e, 0x2e, 0x70,
	0x6c, 0x61, 0x6e, 0x74, 0x65, 0x73, 0x74, 0x63, 0x61, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x6c, 0x61, 0x6e, 0x54, 0x65, 0x73, 0x74, 0x43, 0x61, 0x73,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x22, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x1c, 0x22, 0x17, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x2f, 0x70, 0x6c,
	0x61, 0x6e, 0x74, 0x65, 0x73, 0x74, 0x63, 0x61, 0x73, 0x65, 0x3a, 0x01, 0x2a, 0x12, 0xc7, 0x01,
	0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x6c, 0x61, 0x6e, 0x54, 0x65, 0x73, 0x74,
	0x43, 0x61, 0x73, 0x65, 0x12, 0x45, 0x2e, 0x73, 0x6d, 0x6f, 0x6b, 0x65, 0x74, 0x65, 0x73, 0x74,
	0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x70, 0x6c, 0x61,
	0x6e, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x65, 0x73, 0x74, 0x63, 0x61, 0x73, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x6c, 0x61, 0x6e, 0x54, 0x65, 0x73, 0x74,
	0x43, 0x61, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x46, 0x2e, 0x73, 0x6d,
	0x6f, 0x6b, 0x65, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e,
	0x74, 0x65, 0x73, 0x74, 0x70, 0x6c, 0x61, 0x6e, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x65, 0x73,
	0x74, 0x63, 0x61, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50,
	0x6c, 0x61, 0x6e, 0x54, 0x65, 0x73, 0x74, 0x43, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x22, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1c, 0x22, 0x17, 0x2f, 0x76, 0x31,
	0x2f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x65, 0x73, 0x74,
	0x63, 0x61, 0x73, 0x65, 0x3a, 0x01, 0x2a, 0x12, 0xc7, 0x01, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x50, 0x6c, 0x61, 0x6e, 0x54, 0x65, 0x73, 0x74, 0x43, 0x61, 0x73, 0x65, 0x12, 0x45,
	0x2e, 0x73, 0x6d, 0x6f, 0x6b, 0x65, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77,
	0x61, 0x79, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x70, 0x6c, 0x61, 0x6e, 0x2e, 0x70, 0x6c, 0x61, 0x6e,
	0x74, 0x65, 0x73, 0x74, 0x63, 0x61, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x50, 0x6c, 0x61, 0x6e, 0x54, 0x65, 0x73, 0x74, 0x43, 0x61, 0x73, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x46, 0x2e, 0x73, 0x6d, 0x6f, 0x6b, 0x65, 0x74, 0x65, 0x73,
	0x74, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x70, 0x6c,
	0x61, 0x6e, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x65, 0x73, 0x74, 0x63, 0x61, 0x73, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x6c, 0x61, 0x6e, 0x54, 0x65, 0x73,
	0x74, 0x43, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x22, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x1c, 0x22, 0x17, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x65, 0x73, 0x74, 0x63, 0x61, 0x73, 0x65, 0x3a, 0x01,
	0x2a, 0x12, 0xbf, 0x01, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x50, 0x6c, 0x61, 0x6e, 0x54, 0x65, 0x73,
	0x74, 0x43, 0x61, 0x73, 0x65, 0x73, 0x12, 0x43, 0x2e, 0x73, 0x6d, 0x6f, 0x6b, 0x65, 0x74, 0x65,
	0x73, 0x74, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x70,
	0x6c, 0x61, 0x6e, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x65, 0x73, 0x74, 0x63, 0x61, 0x73, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x6c, 0x61, 0x6e, 0x54, 0x65, 0x73, 0x74, 0x43,
	0x61, 0x73, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x44, 0x2e, 0x73, 0x6d,
	0x6f, 0x6b, 0x65, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e,
	0x74, 0x65, 0x73, 0x74, 0x70, 0x6c, 0x61, 0x6e, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x65, 0x73,
	0x74, 0x63, 0x61, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x6c, 0x61, 0x6e,
	0x54, 0x65, 0x73, 0x74, 0x43, 0x61, 0x73, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x20, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1a, 0x22, 0x15, 0x2f, 0x76, 0x31, 0x2f, 0x67,
	0x65, 0x74, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x65, 0x73, 0x74, 0x63, 0x61, 0x73, 0x65, 0x73,
	0x3a, 0x01, 0x2a, 0x42, 0x4e, 0x5a, 0x4c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x4e, 0x70, 0x6f, 0x6f, 0x6c, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x73, 0x6d,
	0x6f, 0x6b, 0x65, 0x74, 0x65, 0x73, 0x74, 0x2f, 0x67, 0x77, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x65,
	0x73, 0x74, 0x70, 0x6c, 0x61, 0x6e, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x65, 0x73, 0x74, 0x63,
	0x61, 0x73, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_npool_smoketest_gw_v1_testplan_plantestcase_plantestcase_proto_rawDescOnce sync.Once
	file_npool_smoketest_gw_v1_testplan_plantestcase_plantestcase_proto_rawDescData = file_npool_smoketest_gw_v1_testplan_plantestcase_plantestcase_proto_rawDesc
)

func file_npool_smoketest_gw_v1_testplan_plantestcase_plantestcase_proto_rawDescGZIP() []byte {
	file_npool_smoketest_gw_v1_testplan_plantestcase_plantestcase_proto_rawDescOnce.Do(func() {
		file_npool_smoketest_gw_v1_testplan_plantestcase_plantestcase_proto_rawDescData = protoimpl.X.CompressGZIP(file_npool_smoketest_gw_v1_testplan_plantestcase_plantestcase_proto_rawDescData)
	})
	return file_npool_smoketest_gw_v1_testplan_plantestcase_plantestcase_proto_rawDescData
}

var file_npool_smoketest_gw_v1_testplan_plantestcase_plantestcase_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_npool_smoketest_gw_v1_testplan_plantestcase_plantestcase_proto_goTypes = []interface{}{
	(*CreatePlanTestCaseRequest)(nil),  // 0: smoketest.gateway.testplan.plantestcase.v1.CreatePlanTestCaseRequest
	(*CreatePlanTestCaseResponse)(nil), // 1: smoketest.gateway.testplan.plantestcase.v1.CreatePlanTestCaseResponse
	(*UpdatePlanTestCaseRequest)(nil),  // 2: smoketest.gateway.testplan.plantestcase.v1.UpdatePlanTestCaseRequest
	(*UpdatePlanTestCaseResponse)(nil), // 3: smoketest.gateway.testplan.plantestcase.v1.UpdatePlanTestCaseResponse
	(*GetPlanTestCasesRequest)(nil),    // 4: smoketest.gateway.testplan.plantestcase.v1.GetPlanTestCasesRequest
	(*GetPlanTestCasesResponse)(nil),   // 5: smoketest.gateway.testplan.plantestcase.v1.GetPlanTestCasesResponse
	(*DeletePlanTestCaseRequest)(nil),  // 6: smoketest.gateway.testplan.plantestcase.v1.DeletePlanTestCaseRequest
	(*DeletePlanTestCaseResponse)(nil), // 7: smoketest.gateway.testplan.plantestcase.v1.DeletePlanTestCaseResponse
	(*plantestcase.PlanTestCase)(nil),  // 8: smoketest.middleware.testplan.plantestcase.v1.PlanTestCase
	(plantestcase.TestCaseResult)(0),   // 9: smoketest.middleware.testplan.plantestcase.v1.TestCaseResult
}
var file_npool_smoketest_gw_v1_testplan_plantestcase_plantestcase_proto_depIdxs = []int32{
	8, // 0: smoketest.gateway.testplan.plantestcase.v1.CreatePlanTestCaseResponse.Info:type_name -> smoketest.middleware.testplan.plantestcase.v1.PlanTestCase
	9, // 1: smoketest.gateway.testplan.plantestcase.v1.UpdatePlanTestCaseRequest.Result:type_name -> smoketest.middleware.testplan.plantestcase.v1.TestCaseResult
	8, // 2: smoketest.gateway.testplan.plantestcase.v1.UpdatePlanTestCaseResponse.Info:type_name -> smoketest.middleware.testplan.plantestcase.v1.PlanTestCase
	8, // 3: smoketest.gateway.testplan.plantestcase.v1.GetPlanTestCasesResponse.Infos:type_name -> smoketest.middleware.testplan.plantestcase.v1.PlanTestCase
	8, // 4: smoketest.gateway.testplan.plantestcase.v1.DeletePlanTestCaseResponse.Info:type_name -> smoketest.middleware.testplan.plantestcase.v1.PlanTestCase
	0, // 5: smoketest.gateway.testplan.plantestcase.v1.Gateway.CreatePlanTestCase:input_type -> smoketest.gateway.testplan.plantestcase.v1.CreatePlanTestCaseRequest
	6, // 6: smoketest.gateway.testplan.plantestcase.v1.Gateway.DeletePlanTestCase:input_type -> smoketest.gateway.testplan.plantestcase.v1.DeletePlanTestCaseRequest
	2, // 7: smoketest.gateway.testplan.plantestcase.v1.Gateway.UpdatePlanTestCase:input_type -> smoketest.gateway.testplan.plantestcase.v1.UpdatePlanTestCaseRequest
	4, // 8: smoketest.gateway.testplan.plantestcase.v1.Gateway.GetPlanTestCases:input_type -> smoketest.gateway.testplan.plantestcase.v1.GetPlanTestCasesRequest
	1, // 9: smoketest.gateway.testplan.plantestcase.v1.Gateway.CreatePlanTestCase:output_type -> smoketest.gateway.testplan.plantestcase.v1.CreatePlanTestCaseResponse
	7, // 10: smoketest.gateway.testplan.plantestcase.v1.Gateway.DeletePlanTestCase:output_type -> smoketest.gateway.testplan.plantestcase.v1.DeletePlanTestCaseResponse
	3, // 11: smoketest.gateway.testplan.plantestcase.v1.Gateway.UpdatePlanTestCase:output_type -> smoketest.gateway.testplan.plantestcase.v1.UpdatePlanTestCaseResponse
	5, // 12: smoketest.gateway.testplan.plantestcase.v1.Gateway.GetPlanTestCases:output_type -> smoketest.gateway.testplan.plantestcase.v1.GetPlanTestCasesResponse
	9, // [9:13] is the sub-list for method output_type
	5, // [5:9] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_npool_smoketest_gw_v1_testplan_plantestcase_plantestcase_proto_init() }
func file_npool_smoketest_gw_v1_testplan_plantestcase_plantestcase_proto_init() {
	if File_npool_smoketest_gw_v1_testplan_plantestcase_plantestcase_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_npool_smoketest_gw_v1_testplan_plantestcase_plantestcase_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_npool_smoketest_gw_v1_testplan_plantestcase_plantestcase_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_npool_smoketest_gw_v1_testplan_plantestcase_plantestcase_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdatePlanTestCaseRequest); i {
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
		file_npool_smoketest_gw_v1_testplan_plantestcase_plantestcase_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdatePlanTestCaseResponse); i {
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
		file_npool_smoketest_gw_v1_testplan_plantestcase_plantestcase_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
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
		file_npool_smoketest_gw_v1_testplan_plantestcase_plantestcase_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
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
		file_npool_smoketest_gw_v1_testplan_plantestcase_plantestcase_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
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
		file_npool_smoketest_gw_v1_testplan_plantestcase_plantestcase_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
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
	}
	file_npool_smoketest_gw_v1_testplan_plantestcase_plantestcase_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_npool_smoketest_gw_v1_testplan_plantestcase_plantestcase_proto_msgTypes[2].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_npool_smoketest_gw_v1_testplan_plantestcase_plantestcase_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_npool_smoketest_gw_v1_testplan_plantestcase_plantestcase_proto_goTypes,
		DependencyIndexes: file_npool_smoketest_gw_v1_testplan_plantestcase_plantestcase_proto_depIdxs,
		MessageInfos:      file_npool_smoketest_gw_v1_testplan_plantestcase_plantestcase_proto_msgTypes,
	}.Build()
	File_npool_smoketest_gw_v1_testplan_plantestcase_plantestcase_proto = out.File
	file_npool_smoketest_gw_v1_testplan_plantestcase_plantestcase_proto_rawDesc = nil
	file_npool_smoketest_gw_v1_testplan_plantestcase_plantestcase_proto_goTypes = nil
	file_npool_smoketest_gw_v1_testplan_plantestcase_plantestcase_proto_depIdxs = nil
}
