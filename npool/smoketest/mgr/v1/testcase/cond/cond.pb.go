// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.18.1
// source: npool/smoketest/mgr/v1/testcase/cond/cond.proto

package cond

import (
	npool "github.com/NpoolPlatform/message/npool"
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

type CondType int32

const (
	CondType_DefaultCondType CondType = 0
	CondType_PreCondition    CondType = 10
	CondType_Cleaner         CondType = 20
)

// Enum value maps for CondType.
var (
	CondType_name = map[int32]string{
		0:  "DefaultCondType",
		10: "PreCondition",
		20: "Cleaner",
	}
	CondType_value = map[string]int32{
		"DefaultCondType": 0,
		"PreCondition":    10,
		"Cleaner":         20,
	}
)

func (x CondType) Enum() *CondType {
	p := new(CondType)
	*p = x
	return p
}

func (x CondType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CondType) Descriptor() protoreflect.EnumDescriptor {
	return file_npool_smoketest_mgr_v1_testcase_cond_cond_proto_enumTypes[0].Descriptor()
}

func (CondType) Type() protoreflect.EnumType {
	return &file_npool_smoketest_mgr_v1_testcase_cond_cond_proto_enumTypes[0]
}

func (x CondType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CondType.Descriptor instead.
func (CondType) EnumDescriptor() ([]byte, []int) {
	return file_npool_smoketest_mgr_v1_testcase_cond_cond_proto_rawDescGZIP(), []int{0}
}

type CondReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID             *string   `protobuf:"bytes,10,opt,name=ID,proto3,oneof" json:"ID,omitempty"`
	CondType       *CondType `protobuf:"varint,20,opt,name=CondType,proto3,enum=smoketest.manager.testcase.cond.v1.CondType,oneof" json:"CondType,omitempty"`
	TestCaseID     *string   `protobuf:"bytes,30,opt,name=TestCaseID,proto3,oneof" json:"TestCaseID,omitempty"`
	CondTestCaseID *string   `protobuf:"bytes,40,opt,name=CondTestCaseID,proto3,oneof" json:"CondTestCaseID,omitempty"`
	Index          *uint32   `protobuf:"varint,50,opt,name=Index,proto3,oneof" json:"Index,omitempty"`
	CreatedAt      *uint32   `protobuf:"varint,60,opt,name=CreatedAt,proto3,oneof" json:"CreatedAt,omitempty"`
	ArgumentMap    *string   `protobuf:"bytes,70,opt,name=ArgumentMap,proto3,oneof" json:"ArgumentMap,omitempty"`
}

func (x *CondReq) Reset() {
	*x = CondReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_smoketest_mgr_v1_testcase_cond_cond_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CondReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CondReq) ProtoMessage() {}

func (x *CondReq) ProtoReflect() protoreflect.Message {
	mi := &file_npool_smoketest_mgr_v1_testcase_cond_cond_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CondReq.ProtoReflect.Descriptor instead.
func (*CondReq) Descriptor() ([]byte, []int) {
	return file_npool_smoketest_mgr_v1_testcase_cond_cond_proto_rawDescGZIP(), []int{0}
}

func (x *CondReq) GetID() string {
	if x != nil && x.ID != nil {
		return *x.ID
	}
	return ""
}

func (x *CondReq) GetCondType() CondType {
	if x != nil && x.CondType != nil {
		return *x.CondType
	}
	return CondType_DefaultCondType
}

func (x *CondReq) GetTestCaseID() string {
	if x != nil && x.TestCaseID != nil {
		return *x.TestCaseID
	}
	return ""
}

func (x *CondReq) GetCondTestCaseID() string {
	if x != nil && x.CondTestCaseID != nil {
		return *x.CondTestCaseID
	}
	return ""
}

func (x *CondReq) GetIndex() uint32 {
	if x != nil && x.Index != nil {
		return *x.Index
	}
	return 0
}

func (x *CondReq) GetCreatedAt() uint32 {
	if x != nil && x.CreatedAt != nil {
		return *x.CreatedAt
	}
	return 0
}

func (x *CondReq) GetArgumentMap() string {
	if x != nil && x.ArgumentMap != nil {
		return *x.ArgumentMap
	}
	return ""
}

type Cond struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: sql:"id"
	ID string `protobuf:"bytes,10,opt,name=ID,proto3" json:"ID,omitempty" sql:"id"`
	// @inject_tag: sql:"cond_type"
	CondType CondType `protobuf:"varint,20,opt,name=CondType,proto3,enum=smoketest.manager.testcase.cond.v1.CondType" json:"CondType,omitempty" sql:"cond_type"`
	// @inject_tag: sql:"test_case_id"
	TestCaseID string `protobuf:"bytes,30,opt,name=TestCaseID,proto3" json:"TestCaseID,omitempty" sql:"test_case_id"`
	// @inject_tag: sql:"cond_test_case_id"
	CondTestCaseID string `protobuf:"bytes,40,opt,name=CondTestCaseID,proto3" json:"CondTestCaseID,omitempty" sql:"cond_test_case_id"`
	// @inject_tag: sql:"index"
	Index uint32 `protobuf:"varint,50,opt,name=Index,proto3" json:"Index,omitempty" sql:"index"`
	// @inject_tag: sql:"argument_map"
	ArgumentMap string `protobuf:"bytes,60,opt,name=ArgumentMap,proto3" json:"ArgumentMap,omitempty" sql:"argument_map"`
	// @inject_tag: sql:"created_at"
	CreatedAt uint32 `protobuf:"varint,70,opt,name=CreatedAt,proto3" json:"CreatedAt,omitempty" sql:"created_at"`
	// @inject_tag: sql:"updated_at"
	UpdatedAt uint32 `protobuf:"varint,80,opt,name=UpdatedAt,proto3" json:"UpdatedAt,omitempty" sql:"updated_at"`
	// @inject_tag: sql:"deleted_at"
	DeletedAt uint32 `protobuf:"varint,90,opt,name=DeletedAt,proto3" json:"DeletedAt,omitempty" sql:"deleted_at"`
}

func (x *Cond) Reset() {
	*x = Cond{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_smoketest_mgr_v1_testcase_cond_cond_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Cond) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Cond) ProtoMessage() {}

func (x *Cond) ProtoReflect() protoreflect.Message {
	mi := &file_npool_smoketest_mgr_v1_testcase_cond_cond_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Cond.ProtoReflect.Descriptor instead.
func (*Cond) Descriptor() ([]byte, []int) {
	return file_npool_smoketest_mgr_v1_testcase_cond_cond_proto_rawDescGZIP(), []int{1}
}

func (x *Cond) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *Cond) GetCondType() CondType {
	if x != nil {
		return x.CondType
	}
	return CondType_DefaultCondType
}

func (x *Cond) GetTestCaseID() string {
	if x != nil {
		return x.TestCaseID
	}
	return ""
}

func (x *Cond) GetCondTestCaseID() string {
	if x != nil {
		return x.CondTestCaseID
	}
	return ""
}

func (x *Cond) GetIndex() uint32 {
	if x != nil {
		return x.Index
	}
	return 0
}

func (x *Cond) GetArgumentMap() string {
	if x != nil {
		return x.ArgumentMap
	}
	return ""
}

func (x *Cond) GetCreatedAt() uint32 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *Cond) GetUpdatedAt() uint32 {
	if x != nil {
		return x.UpdatedAt
	}
	return 0
}

func (x *Cond) GetDeletedAt() uint32 {
	if x != nil {
		return x.DeletedAt
	}
	return 0
}

type Conds struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID             *npool.StringVal `protobuf:"bytes,10,opt,name=ID,proto3,oneof" json:"ID,omitempty"`
	CondType       *npool.StringVal `protobuf:"bytes,20,opt,name=CondType,proto3,oneof" json:"CondType,omitempty"`
	TestCaseID     *npool.StringVal `protobuf:"bytes,30,opt,name=TestCaseID,proto3,oneof" json:"TestCaseID,omitempty"`
	CondTestCaseID *npool.StringVal `protobuf:"bytes,40,opt,name=CondTestCaseID,proto3,oneof" json:"CondTestCaseID,omitempty"`
}

func (x *Conds) Reset() {
	*x = Conds{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_smoketest_mgr_v1_testcase_cond_cond_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Conds) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Conds) ProtoMessage() {}

func (x *Conds) ProtoReflect() protoreflect.Message {
	mi := &file_npool_smoketest_mgr_v1_testcase_cond_cond_proto_msgTypes[2]
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
	return file_npool_smoketest_mgr_v1_testcase_cond_cond_proto_rawDescGZIP(), []int{2}
}

func (x *Conds) GetID() *npool.StringVal {
	if x != nil {
		return x.ID
	}
	return nil
}

func (x *Conds) GetCondType() *npool.StringVal {
	if x != nil {
		return x.CondType
	}
	return nil
}

func (x *Conds) GetTestCaseID() *npool.StringVal {
	if x != nil {
		return x.TestCaseID
	}
	return nil
}

func (x *Conds) GetCondTestCaseID() *npool.StringVal {
	if x != nil {
		return x.CondTestCaseID
	}
	return nil
}

var File_npool_smoketest_mgr_v1_testcase_cond_cond_proto protoreflect.FileDescriptor

var file_npool_smoketest_mgr_v1_testcase_cond_cond_proto_rawDesc = []byte{
	0x0a, 0x2f, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x73, 0x6d, 0x6f, 0x6b, 0x65, 0x74, 0x65, 0x73,
	0x74, 0x2f, 0x6d, 0x67, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x63, 0x61, 0x73,
	0x65, 0x2f, 0x63, 0x6f, 0x6e, 0x64, 0x2f, 0x63, 0x6f, 0x6e, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x22, 0x73, 0x6d, 0x6f, 0x6b, 0x65, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x6d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x72, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x63, 0x61, 0x73, 0x65, 0x2e, 0x63, 0x6f,
	0x6e, 0x64, 0x2e, 0x76, 0x31, 0x1a, 0x11, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x6e, 0x70, 0x6f,
	0x6f, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x82, 0x03, 0x0a, 0x07, 0x43, 0x6f, 0x6e,
	0x64, 0x52, 0x65, 0x71, 0x12, 0x13, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09,
	0x48, 0x00, 0x52, 0x02, 0x49, 0x44, 0x88, 0x01, 0x01, 0x12, 0x4d, 0x0a, 0x08, 0x43, 0x6f, 0x6e,
	0x64, 0x54, 0x79, 0x70, 0x65, 0x18, 0x14, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2c, 0x2e, 0x73, 0x6d,
	0x6f, 0x6b, 0x65, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e,
	0x74, 0x65, 0x73, 0x74, 0x63, 0x61, 0x73, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x64, 0x2e, 0x76, 0x31,
	0x2e, 0x43, 0x6f, 0x6e, 0x64, 0x54, 0x79, 0x70, 0x65, 0x48, 0x01, 0x52, 0x08, 0x43, 0x6f, 0x6e,
	0x64, 0x54, 0x79, 0x70, 0x65, 0x88, 0x01, 0x01, 0x12, 0x23, 0x0a, 0x0a, 0x54, 0x65, 0x73, 0x74,
	0x43, 0x61, 0x73, 0x65, 0x49, 0x44, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x09, 0x48, 0x02, 0x52, 0x0a,
	0x54, 0x65, 0x73, 0x74, 0x43, 0x61, 0x73, 0x65, 0x49, 0x44, 0x88, 0x01, 0x01, 0x12, 0x2b, 0x0a,
	0x0e, 0x43, 0x6f, 0x6e, 0x64, 0x54, 0x65, 0x73, 0x74, 0x43, 0x61, 0x73, 0x65, 0x49, 0x44, 0x18,
	0x28, 0x20, 0x01, 0x28, 0x09, 0x48, 0x03, 0x52, 0x0e, 0x43, 0x6f, 0x6e, 0x64, 0x54, 0x65, 0x73,
	0x74, 0x43, 0x61, 0x73, 0x65, 0x49, 0x44, 0x88, 0x01, 0x01, 0x12, 0x19, 0x0a, 0x05, 0x49, 0x6e,
	0x64, 0x65, 0x78, 0x18, 0x32, 0x20, 0x01, 0x28, 0x0d, 0x48, 0x04, 0x52, 0x05, 0x49, 0x6e, 0x64,
	0x65, 0x78, 0x88, 0x01, 0x01, 0x12, 0x21, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x18, 0x3c, 0x20, 0x01, 0x28, 0x0d, 0x48, 0x05, 0x52, 0x09, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x88, 0x01, 0x01, 0x12, 0x25, 0x0a, 0x0b, 0x41, 0x72, 0x67, 0x75,
	0x6d, 0x65, 0x6e, 0x74, 0x4d, 0x61, 0x70, 0x18, 0x46, 0x20, 0x01, 0x28, 0x09, 0x48, 0x06, 0x52,
	0x0b, 0x41, 0x72, 0x67, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x4d, 0x61, 0x70, 0x88, 0x01, 0x01, 0x42,
	0x05, 0x0a, 0x03, 0x5f, 0x49, 0x44, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x43, 0x6f, 0x6e, 0x64, 0x54,
	0x79, 0x70, 0x65, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x54, 0x65, 0x73, 0x74, 0x43, 0x61, 0x73, 0x65,
	0x49, 0x44, 0x42, 0x11, 0x0a, 0x0f, 0x5f, 0x43, 0x6f, 0x6e, 0x64, 0x54, 0x65, 0x73, 0x74, 0x43,
	0x61, 0x73, 0x65, 0x49, 0x44, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x42,
	0x0c, 0x0a, 0x0a, 0x5f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x42, 0x0e, 0x0a,
	0x0c, 0x5f, 0x41, 0x72, 0x67, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x4d, 0x61, 0x70, 0x22, 0xba, 0x02,
	0x0a, 0x04, 0x43, 0x6f, 0x6e, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x12, 0x48, 0x0a, 0x08, 0x43, 0x6f, 0x6e, 0x64, 0x54, 0x79,
	0x70, 0x65, 0x18, 0x14, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2c, 0x2e, 0x73, 0x6d, 0x6f, 0x6b, 0x65,
	0x74, 0x65, 0x73, 0x74, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x74, 0x65, 0x73,
	0x74, 0x63, 0x61, 0x73, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f,
	0x6e, 0x64, 0x54, 0x79, 0x70, 0x65, 0x52, 0x08, 0x43, 0x6f, 0x6e, 0x64, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x1e, 0x0a, 0x0a, 0x54, 0x65, 0x73, 0x74, 0x43, 0x61, 0x73, 0x65, 0x49, 0x44, 0x18, 0x1e,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x54, 0x65, 0x73, 0x74, 0x43, 0x61, 0x73, 0x65, 0x49, 0x44,
	0x12, 0x26, 0x0a, 0x0e, 0x43, 0x6f, 0x6e, 0x64, 0x54, 0x65, 0x73, 0x74, 0x43, 0x61, 0x73, 0x65,
	0x49, 0x44, 0x18, 0x28, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x43, 0x6f, 0x6e, 0x64, 0x54, 0x65,
	0x73, 0x74, 0x43, 0x61, 0x73, 0x65, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x49, 0x6e, 0x64, 0x65,
	0x78, 0x18, 0x32, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x20,
	0x0a, 0x0b, 0x41, 0x72, 0x67, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x4d, 0x61, 0x70, 0x18, 0x3c, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x41, 0x72, 0x67, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x4d, 0x61, 0x70,
	0x12, 0x1c, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x46, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1c,
	0x0a, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x50, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1c, 0x0a, 0x09,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x5a, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x09, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x99, 0x02, 0x0a, 0x05, 0x43,
	0x6f, 0x6e, 0x64, 0x73, 0x12, 0x28, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x13, 0x2e, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x72, 0x69,
	0x6e, 0x67, 0x56, 0x61, 0x6c, 0x48, 0x00, 0x52, 0x02, 0x49, 0x44, 0x88, 0x01, 0x01, 0x12, 0x34,
	0x0a, 0x08, 0x43, 0x6f, 0x6e, 0x64, 0x54, 0x79, 0x70, 0x65, 0x18, 0x14, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x13, 0x2e, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x72, 0x69,
	0x6e, 0x67, 0x56, 0x61, 0x6c, 0x48, 0x01, 0x52, 0x08, 0x43, 0x6f, 0x6e, 0x64, 0x54, 0x79, 0x70,
	0x65, 0x88, 0x01, 0x01, 0x12, 0x38, 0x0a, 0x0a, 0x54, 0x65, 0x73, 0x74, 0x43, 0x61, 0x73, 0x65,
	0x49, 0x44, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x6e, 0x70, 0x6f, 0x6f, 0x6c,
	0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x48, 0x02, 0x52,
	0x0a, 0x54, 0x65, 0x73, 0x74, 0x43, 0x61, 0x73, 0x65, 0x49, 0x44, 0x88, 0x01, 0x01, 0x12, 0x40,
	0x0a, 0x0e, 0x43, 0x6f, 0x6e, 0x64, 0x54, 0x65, 0x73, 0x74, 0x43, 0x61, 0x73, 0x65, 0x49, 0x44,
	0x18, 0x28, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2e, 0x76,
	0x31, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x48, 0x03, 0x52, 0x0e, 0x43,
	0x6f, 0x6e, 0x64, 0x54, 0x65, 0x73, 0x74, 0x43, 0x61, 0x73, 0x65, 0x49, 0x44, 0x88, 0x01, 0x01,
	0x42, 0x05, 0x0a, 0x03, 0x5f, 0x49, 0x44, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x43, 0x6f, 0x6e, 0x64,
	0x54, 0x79, 0x70, 0x65, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x54, 0x65, 0x73, 0x74, 0x43, 0x61, 0x73,
	0x65, 0x49, 0x44, 0x42, 0x11, 0x0a, 0x0f, 0x5f, 0x43, 0x6f, 0x6e, 0x64, 0x54, 0x65, 0x73, 0x74,
	0x43, 0x61, 0x73, 0x65, 0x49, 0x44, 0x2a, 0x3e, 0x0a, 0x08, 0x43, 0x6f, 0x6e, 0x64, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x13, 0x0a, 0x0f, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x43, 0x6f, 0x6e,
	0x64, 0x54, 0x79, 0x70, 0x65, 0x10, 0x00, 0x12, 0x10, 0x0a, 0x0c, 0x50, 0x72, 0x65, 0x43, 0x6f,
	0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x10, 0x0a, 0x12, 0x0b, 0x0a, 0x07, 0x43, 0x6c, 0x65,
	0x61, 0x6e, 0x65, 0x72, 0x10, 0x14, 0x42, 0x47, 0x5a, 0x45, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4e, 0x70, 0x6f, 0x6f, 0x6c, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f,
	0x72, 0x6d, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x6e, 0x70, 0x6f, 0x6f, 0x6c,
	0x2f, 0x73, 0x6d, 0x6f, 0x6b, 0x65, 0x74, 0x65, 0x73, 0x74, 0x2f, 0x6d, 0x67, 0x72, 0x2f, 0x76,
	0x31, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x63, 0x61, 0x73, 0x65, 0x2f, 0x63, 0x6f, 0x6e, 0x64, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_npool_smoketest_mgr_v1_testcase_cond_cond_proto_rawDescOnce sync.Once
	file_npool_smoketest_mgr_v1_testcase_cond_cond_proto_rawDescData = file_npool_smoketest_mgr_v1_testcase_cond_cond_proto_rawDesc
)

func file_npool_smoketest_mgr_v1_testcase_cond_cond_proto_rawDescGZIP() []byte {
	file_npool_smoketest_mgr_v1_testcase_cond_cond_proto_rawDescOnce.Do(func() {
		file_npool_smoketest_mgr_v1_testcase_cond_cond_proto_rawDescData = protoimpl.X.CompressGZIP(file_npool_smoketest_mgr_v1_testcase_cond_cond_proto_rawDescData)
	})
	return file_npool_smoketest_mgr_v1_testcase_cond_cond_proto_rawDescData
}

var file_npool_smoketest_mgr_v1_testcase_cond_cond_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_npool_smoketest_mgr_v1_testcase_cond_cond_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_npool_smoketest_mgr_v1_testcase_cond_cond_proto_goTypes = []interface{}{
	(CondType)(0),           // 0: smoketest.manager.testcase.cond.v1.CondType
	(*CondReq)(nil),         // 1: smoketest.manager.testcase.cond.v1.CondReq
	(*Cond)(nil),            // 2: smoketest.manager.testcase.cond.v1.Cond
	(*Conds)(nil),           // 3: smoketest.manager.testcase.cond.v1.Conds
	(*npool.StringVal)(nil), // 4: npool.v1.StringVal
}
var file_npool_smoketest_mgr_v1_testcase_cond_cond_proto_depIdxs = []int32{
	0, // 0: smoketest.manager.testcase.cond.v1.CondReq.CondType:type_name -> smoketest.manager.testcase.cond.v1.CondType
	0, // 1: smoketest.manager.testcase.cond.v1.Cond.CondType:type_name -> smoketest.manager.testcase.cond.v1.CondType
	4, // 2: smoketest.manager.testcase.cond.v1.Conds.ID:type_name -> npool.v1.StringVal
	4, // 3: smoketest.manager.testcase.cond.v1.Conds.CondType:type_name -> npool.v1.StringVal
	4, // 4: smoketest.manager.testcase.cond.v1.Conds.TestCaseID:type_name -> npool.v1.StringVal
	4, // 5: smoketest.manager.testcase.cond.v1.Conds.CondTestCaseID:type_name -> npool.v1.StringVal
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_npool_smoketest_mgr_v1_testcase_cond_cond_proto_init() }
func file_npool_smoketest_mgr_v1_testcase_cond_cond_proto_init() {
	if File_npool_smoketest_mgr_v1_testcase_cond_cond_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_npool_smoketest_mgr_v1_testcase_cond_cond_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CondReq); i {
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
		file_npool_smoketest_mgr_v1_testcase_cond_cond_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Cond); i {
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
		file_npool_smoketest_mgr_v1_testcase_cond_cond_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
	}
	file_npool_smoketest_mgr_v1_testcase_cond_cond_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_npool_smoketest_mgr_v1_testcase_cond_cond_proto_msgTypes[2].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_npool_smoketest_mgr_v1_testcase_cond_cond_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_npool_smoketest_mgr_v1_testcase_cond_cond_proto_goTypes,
		DependencyIndexes: file_npool_smoketest_mgr_v1_testcase_cond_cond_proto_depIdxs,
		EnumInfos:         file_npool_smoketest_mgr_v1_testcase_cond_cond_proto_enumTypes,
		MessageInfos:      file_npool_smoketest_mgr_v1_testcase_cond_cond_proto_msgTypes,
	}.Build()
	File_npool_smoketest_mgr_v1_testcase_cond_cond_proto = out.File
	file_npool_smoketest_mgr_v1_testcase_cond_cond_proto_rawDesc = nil
	file_npool_smoketest_mgr_v1_testcase_cond_cond_proto_goTypes = nil
	file_npool_smoketest_mgr_v1_testcase_cond_cond_proto_depIdxs = nil
}
