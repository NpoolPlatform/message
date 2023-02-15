// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.18.1
// source: npool/basal/mw/v1/usercode/usercode.proto

package usercode

import (
	_ "github.com/NpoolPlatform/message/npool"
	v1 "github.com/NpoolPlatform/message/npool/basetypes/v1"
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

type UserCode struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Prefix      string        `protobuf:"bytes,10,opt,name=Prefix,proto3" json:"Prefix,omitempty"`
	AppID       string        `protobuf:"bytes,20,opt,name=AppID,proto3" json:"AppID,omitempty"`
	Account     string        `protobuf:"bytes,30,opt,name=Account,proto3" json:"Account,omitempty"`
	AccountType v1.SignMethod `protobuf:"varint,40,opt,name=AccountType,proto3,enum=basetypes.v1.SignMethod" json:"AccountType,omitempty"`
	UsedFor     v1.UsedFor    `protobuf:"varint,50,opt,name=UsedFor,proto3,enum=basetypes.v1.UsedFor" json:"UsedFor,omitempty"`
	NextAt      uint32        `protobuf:"varint,60,opt,name=NextAt,proto3" json:"NextAt,omitempty"`
	ExpireAt    uint32        `protobuf:"varint,70,opt,name=ExpireAt,proto3" json:"ExpireAt,omitempty"`
	Code        string        `protobuf:"bytes,80,opt,name=Code,proto3" json:"Code,omitempty"`
}

func (x *UserCode) Reset() {
	*x = UserCode{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_basal_mw_v1_usercode_usercode_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserCode) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserCode) ProtoMessage() {}

func (x *UserCode) ProtoReflect() protoreflect.Message {
	mi := &file_npool_basal_mw_v1_usercode_usercode_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserCode.ProtoReflect.Descriptor instead.
func (*UserCode) Descriptor() ([]byte, []int) {
	return file_npool_basal_mw_v1_usercode_usercode_proto_rawDescGZIP(), []int{0}
}

func (x *UserCode) GetPrefix() string {
	if x != nil {
		return x.Prefix
	}
	return ""
}

func (x *UserCode) GetAppID() string {
	if x != nil {
		return x.AppID
	}
	return ""
}

func (x *UserCode) GetAccount() string {
	if x != nil {
		return x.Account
	}
	return ""
}

func (x *UserCode) GetAccountType() v1.SignMethod {
	if x != nil {
		return x.AccountType
	}
	return v1.SignMethod(0)
}

func (x *UserCode) GetUsedFor() v1.UsedFor {
	if x != nil {
		return x.UsedFor
	}
	return v1.UsedFor(0)
}

func (x *UserCode) GetNextAt() uint32 {
	if x != nil {
		return x.NextAt
	}
	return 0
}

func (x *UserCode) GetExpireAt() uint32 {
	if x != nil {
		return x.ExpireAt
	}
	return 0
}

func (x *UserCode) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

type CreateUserCodeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Prefix      string        `protobuf:"bytes,10,opt,name=Prefix,proto3" json:"Prefix,omitempty"`
	AppID       string        `protobuf:"bytes,20,opt,name=AppID,proto3" json:"AppID,omitempty"`
	Account     string        `protobuf:"bytes,30,opt,name=Account,proto3" json:"Account,omitempty"`
	AccountType v1.SignMethod `protobuf:"varint,40,opt,name=AccountType,proto3,enum=basetypes.v1.SignMethod" json:"AccountType,omitempty"`
	UsedFor     v1.UsedFor    `protobuf:"varint,50,opt,name=UsedFor,proto3,enum=basetypes.v1.UsedFor" json:"UsedFor,omitempty"`
}

func (x *CreateUserCodeRequest) Reset() {
	*x = CreateUserCodeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_basal_mw_v1_usercode_usercode_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateUserCodeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUserCodeRequest) ProtoMessage() {}

func (x *CreateUserCodeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_basal_mw_v1_usercode_usercode_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateUserCodeRequest.ProtoReflect.Descriptor instead.
func (*CreateUserCodeRequest) Descriptor() ([]byte, []int) {
	return file_npool_basal_mw_v1_usercode_usercode_proto_rawDescGZIP(), []int{1}
}

func (x *CreateUserCodeRequest) GetPrefix() string {
	if x != nil {
		return x.Prefix
	}
	return ""
}

func (x *CreateUserCodeRequest) GetAppID() string {
	if x != nil {
		return x.AppID
	}
	return ""
}

func (x *CreateUserCodeRequest) GetAccount() string {
	if x != nil {
		return x.Account
	}
	return ""
}

func (x *CreateUserCodeRequest) GetAccountType() v1.SignMethod {
	if x != nil {
		return x.AccountType
	}
	return v1.SignMethod(0)
}

func (x *CreateUserCodeRequest) GetUsedFor() v1.UsedFor {
	if x != nil {
		return x.UsedFor
	}
	return v1.UsedFor(0)
}

type CreateUserCodeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *UserCode `protobuf:"bytes,10,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (x *CreateUserCodeResponse) Reset() {
	*x = CreateUserCodeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_basal_mw_v1_usercode_usercode_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateUserCodeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUserCodeResponse) ProtoMessage() {}

func (x *CreateUserCodeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_basal_mw_v1_usercode_usercode_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateUserCodeResponse.ProtoReflect.Descriptor instead.
func (*CreateUserCodeResponse) Descriptor() ([]byte, []int) {
	return file_npool_basal_mw_v1_usercode_usercode_proto_rawDescGZIP(), []int{2}
}

func (x *CreateUserCodeResponse) GetInfo() *UserCode {
	if x != nil {
		return x.Info
	}
	return nil
}

type VerifyUserCodeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Prefix      string        `protobuf:"bytes,10,opt,name=Prefix,proto3" json:"Prefix,omitempty"`
	AppID       string        `protobuf:"bytes,20,opt,name=AppID,proto3" json:"AppID,omitempty"`
	Account     string        `protobuf:"bytes,30,opt,name=Account,proto3" json:"Account,omitempty"`
	AccountType v1.SignMethod `protobuf:"varint,40,opt,name=AccountType,proto3,enum=basetypes.v1.SignMethod" json:"AccountType,omitempty"`
	UsedFor     v1.UsedFor    `protobuf:"varint,50,opt,name=UsedFor,proto3,enum=basetypes.v1.UsedFor" json:"UsedFor,omitempty"`
	Code        string        `protobuf:"bytes,60,opt,name=Code,proto3" json:"Code,omitempty"`
}

func (x *VerifyUserCodeRequest) Reset() {
	*x = VerifyUserCodeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_basal_mw_v1_usercode_usercode_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VerifyUserCodeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VerifyUserCodeRequest) ProtoMessage() {}

func (x *VerifyUserCodeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_basal_mw_v1_usercode_usercode_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VerifyUserCodeRequest.ProtoReflect.Descriptor instead.
func (*VerifyUserCodeRequest) Descriptor() ([]byte, []int) {
	return file_npool_basal_mw_v1_usercode_usercode_proto_rawDescGZIP(), []int{3}
}

func (x *VerifyUserCodeRequest) GetPrefix() string {
	if x != nil {
		return x.Prefix
	}
	return ""
}

func (x *VerifyUserCodeRequest) GetAppID() string {
	if x != nil {
		return x.AppID
	}
	return ""
}

func (x *VerifyUserCodeRequest) GetAccount() string {
	if x != nil {
		return x.Account
	}
	return ""
}

func (x *VerifyUserCodeRequest) GetAccountType() v1.SignMethod {
	if x != nil {
		return x.AccountType
	}
	return v1.SignMethod(0)
}

func (x *VerifyUserCodeRequest) GetUsedFor() v1.UsedFor {
	if x != nil {
		return x.UsedFor
	}
	return v1.UsedFor(0)
}

func (x *VerifyUserCodeRequest) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

type VerifyUserCodeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *UserCode `protobuf:"bytes,10,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (x *VerifyUserCodeResponse) Reset() {
	*x = VerifyUserCodeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_basal_mw_v1_usercode_usercode_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VerifyUserCodeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VerifyUserCodeResponse) ProtoMessage() {}

func (x *VerifyUserCodeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_basal_mw_v1_usercode_usercode_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VerifyUserCodeResponse.ProtoReflect.Descriptor instead.
func (*VerifyUserCodeResponse) Descriptor() ([]byte, []int) {
	return file_npool_basal_mw_v1_usercode_usercode_proto_rawDescGZIP(), []int{4}
}

func (x *VerifyUserCodeResponse) GetInfo() *UserCode {
	if x != nil {
		return x.Info
	}
	return nil
}

var File_npool_basal_mw_v1_usercode_usercode_proto protoreflect.FileDescriptor

var file_npool_basal_mw_v1_usercode_usercode_proto_rawDesc = []byte{
	0x0a, 0x29, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x62, 0x61, 0x73, 0x61, 0x6c, 0x2f, 0x6d, 0x77,
	0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x63, 0x6f, 0x64, 0x65, 0x2f, 0x75, 0x73, 0x65,
	0x72, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1c, 0x62, 0x61, 0x73,
	0x61, 0x6c, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x23, 0x6e, 0x70, 0x6f, 0x6f, 0x6c,
	0x2f, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x69,
	0x67, 0x6e, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20,
	0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f,
	0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x64, 0x66, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x11, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x87, 0x02, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x43, 0x6f, 0x64, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x12, 0x14, 0x0a, 0x05, 0x41, 0x70, 0x70, 0x49,
	0x44, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x41, 0x70, 0x70, 0x49, 0x44, 0x12, 0x18,
	0x0a, 0x07, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x3a, 0x0a, 0x0b, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x18, 0x28, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x18, 0x2e,
	0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x69, 0x67,
	0x6e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x52, 0x0b, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x2f, 0x0a, 0x07, 0x55, 0x73, 0x65, 0x64, 0x46, 0x6f, 0x72, 0x18,
	0x32, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x15, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x64, 0x46, 0x6f, 0x72, 0x52, 0x07, 0x55, 0x73,
	0x65, 0x64, 0x46, 0x6f, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x4e, 0x65, 0x78, 0x74, 0x41, 0x74, 0x18,
	0x3c, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x4e, 0x65, 0x78, 0x74, 0x41, 0x74, 0x12, 0x1a, 0x0a,
	0x08, 0x45, 0x78, 0x70, 0x69, 0x72, 0x65, 0x41, 0x74, 0x18, 0x46, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x08, 0x45, 0x78, 0x70, 0x69, 0x72, 0x65, 0x41, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x43, 0x6f, 0x64,
	0x65, 0x18, 0x50, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x22, 0xcc, 0x01,
	0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x43, 0x6f, 0x64, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x50, 0x72, 0x65, 0x66, 0x69,
	0x78, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x12,
	0x14, 0x0a, 0x05, 0x41, 0x70, 0x70, 0x49, 0x44, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x41, 0x70, 0x70, 0x49, 0x44, 0x12, 0x18, 0x0a, 0x07, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x1e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12,
	0x3a, 0x0a, 0x0b, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x18, 0x28,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x18, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x52, 0x0b,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x2f, 0x0a, 0x07, 0x55,
	0x73, 0x65, 0x64, 0x46, 0x6f, 0x72, 0x18, 0x32, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x15, 0x2e, 0x62,
	0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x64,
	0x46, 0x6f, 0x72, 0x52, 0x07, 0x55, 0x73, 0x65, 0x64, 0x46, 0x6f, 0x72, 0x22, 0x54, 0x0a, 0x16,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3a, 0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x62, 0x61, 0x73, 0x61, 0x6c, 0x2e, 0x6d, 0x69, 0x64,
	0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x63, 0x6f, 0x64, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x49, 0x6e,
	0x66, 0x6f, 0x22, 0xe0, 0x01, 0x0a, 0x15, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x55, 0x73, 0x65,
	0x72, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06,
	0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x50, 0x72,
	0x65, 0x66, 0x69, 0x78, 0x12, 0x14, 0x0a, 0x05, 0x41, 0x70, 0x70, 0x49, 0x44, 0x18, 0x14, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x41, 0x70, 0x70, 0x49, 0x44, 0x12, 0x18, 0x0a, 0x07, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x41, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x12, 0x3a, 0x0a, 0x0b, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x54,
	0x79, 0x70, 0x65, 0x18, 0x28, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x18, 0x2e, 0x62, 0x61, 0x73, 0x65,
	0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x4d, 0x65, 0x74,
	0x68, 0x6f, 0x64, 0x52, 0x0b, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x2f, 0x0a, 0x07, 0x55, 0x73, 0x65, 0x64, 0x46, 0x6f, 0x72, 0x18, 0x32, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x15, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x55, 0x73, 0x65, 0x64, 0x46, 0x6f, 0x72, 0x52, 0x07, 0x55, 0x73, 0x65, 0x64, 0x46, 0x6f,
	0x72, 0x12, 0x12, 0x0a, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x3c, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x43, 0x6f, 0x64, 0x65, 0x22, 0x54, 0x0a, 0x16, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x55,
	0x73, 0x65, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x3a, 0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e,
	0x62, 0x61, 0x73, 0x61, 0x6c, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65,
	0x72, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x32, 0x8a, 0x02, 0x0a, 0x0a,
	0x4d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x12, 0x7d, 0x0a, 0x0e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x33, 0x2e, 0x62,
	0x61, 0x73, 0x61, 0x6c, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x34, 0x2e, 0x62, 0x61, 0x73, 0x61, 0x6c, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65,
	0x77, 0x61, 0x72, 0x65, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x7d, 0x0a, 0x0e, 0x56, 0x65, 0x72,
	0x69, 0x66, 0x79, 0x55, 0x73, 0x65, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x33, 0x2e, 0x62, 0x61,
	0x73, 0x61, 0x6c, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x75,
	0x73, 0x65, 0x72, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x56, 0x65, 0x72, 0x69, 0x66,
	0x79, 0x55, 0x73, 0x65, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x34, 0x2e, 0x62, 0x61, 0x73, 0x61, 0x6c, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77,
	0x61, 0x72, 0x65, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x55, 0x73, 0x65, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x3d, 0x5a, 0x3b, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4e, 0x70, 0x6f, 0x6f, 0x6c, 0x50, 0x6c, 0x61, 0x74,
	0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x6e, 0x70, 0x6f,
	0x6f, 0x6c, 0x2f, 0x62, 0x61, 0x73, 0x61, 0x6c, 0x2f, 0x6d, 0x77, 0x2f, 0x76, 0x31, 0x2f, 0x75,
	0x73, 0x65, 0x72, 0x63, 0x6f, 0x64, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_npool_basal_mw_v1_usercode_usercode_proto_rawDescOnce sync.Once
	file_npool_basal_mw_v1_usercode_usercode_proto_rawDescData = file_npool_basal_mw_v1_usercode_usercode_proto_rawDesc
)

func file_npool_basal_mw_v1_usercode_usercode_proto_rawDescGZIP() []byte {
	file_npool_basal_mw_v1_usercode_usercode_proto_rawDescOnce.Do(func() {
		file_npool_basal_mw_v1_usercode_usercode_proto_rawDescData = protoimpl.X.CompressGZIP(file_npool_basal_mw_v1_usercode_usercode_proto_rawDescData)
	})
	return file_npool_basal_mw_v1_usercode_usercode_proto_rawDescData
}

var file_npool_basal_mw_v1_usercode_usercode_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_npool_basal_mw_v1_usercode_usercode_proto_goTypes = []interface{}{
	(*UserCode)(nil),               // 0: basal.middleware.usercode.v1.UserCode
	(*CreateUserCodeRequest)(nil),  // 1: basal.middleware.usercode.v1.CreateUserCodeRequest
	(*CreateUserCodeResponse)(nil), // 2: basal.middleware.usercode.v1.CreateUserCodeResponse
	(*VerifyUserCodeRequest)(nil),  // 3: basal.middleware.usercode.v1.VerifyUserCodeRequest
	(*VerifyUserCodeResponse)(nil), // 4: basal.middleware.usercode.v1.VerifyUserCodeResponse
	(v1.SignMethod)(0),             // 5: basetypes.v1.SignMethod
	(v1.UsedFor)(0),                // 6: basetypes.v1.UsedFor
}
var file_npool_basal_mw_v1_usercode_usercode_proto_depIdxs = []int32{
	5,  // 0: basal.middleware.usercode.v1.UserCode.AccountType:type_name -> basetypes.v1.SignMethod
	6,  // 1: basal.middleware.usercode.v1.UserCode.UsedFor:type_name -> basetypes.v1.UsedFor
	5,  // 2: basal.middleware.usercode.v1.CreateUserCodeRequest.AccountType:type_name -> basetypes.v1.SignMethod
	6,  // 3: basal.middleware.usercode.v1.CreateUserCodeRequest.UsedFor:type_name -> basetypes.v1.UsedFor
	0,  // 4: basal.middleware.usercode.v1.CreateUserCodeResponse.Info:type_name -> basal.middleware.usercode.v1.UserCode
	5,  // 5: basal.middleware.usercode.v1.VerifyUserCodeRequest.AccountType:type_name -> basetypes.v1.SignMethod
	6,  // 6: basal.middleware.usercode.v1.VerifyUserCodeRequest.UsedFor:type_name -> basetypes.v1.UsedFor
	0,  // 7: basal.middleware.usercode.v1.VerifyUserCodeResponse.Info:type_name -> basal.middleware.usercode.v1.UserCode
	1,  // 8: basal.middleware.usercode.v1.Middleware.CreateUserCode:input_type -> basal.middleware.usercode.v1.CreateUserCodeRequest
	3,  // 9: basal.middleware.usercode.v1.Middleware.VerifyUserCode:input_type -> basal.middleware.usercode.v1.VerifyUserCodeRequest
	2,  // 10: basal.middleware.usercode.v1.Middleware.CreateUserCode:output_type -> basal.middleware.usercode.v1.CreateUserCodeResponse
	4,  // 11: basal.middleware.usercode.v1.Middleware.VerifyUserCode:output_type -> basal.middleware.usercode.v1.VerifyUserCodeResponse
	10, // [10:12] is the sub-list for method output_type
	8,  // [8:10] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_npool_basal_mw_v1_usercode_usercode_proto_init() }
func file_npool_basal_mw_v1_usercode_usercode_proto_init() {
	if File_npool_basal_mw_v1_usercode_usercode_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_npool_basal_mw_v1_usercode_usercode_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserCode); i {
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
		file_npool_basal_mw_v1_usercode_usercode_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateUserCodeRequest); i {
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
		file_npool_basal_mw_v1_usercode_usercode_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateUserCodeResponse); i {
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
		file_npool_basal_mw_v1_usercode_usercode_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VerifyUserCodeRequest); i {
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
		file_npool_basal_mw_v1_usercode_usercode_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VerifyUserCodeResponse); i {
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
			RawDescriptor: file_npool_basal_mw_v1_usercode_usercode_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_npool_basal_mw_v1_usercode_usercode_proto_goTypes,
		DependencyIndexes: file_npool_basal_mw_v1_usercode_usercode_proto_depIdxs,
		MessageInfos:      file_npool_basal_mw_v1_usercode_usercode_proto_msgTypes,
	}.Build()
	File_npool_basal_mw_v1_usercode_usercode_proto = out.File
	file_npool_basal_mw_v1_usercode_usercode_proto_rawDesc = nil
	file_npool_basal_mw_v1_usercode_usercode_proto_goTypes = nil
	file_npool_basal_mw_v1_usercode_usercode_proto_depIdxs = nil
}
