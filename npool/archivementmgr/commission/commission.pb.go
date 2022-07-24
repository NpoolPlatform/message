// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.18.1
// source: npool/archivementmgr/commission/commission.proto

package commission

import (
	npool "github.com/NpoolPlatform/message/npool"
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

type CalculateOrderCommissionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderID string `protobuf:"bytes,10,opt,name=OrderID,proto3" json:"OrderID,omitempty"`
}

func (x *CalculateOrderCommissionRequest) Reset() {
	*x = CalculateOrderCommissionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_archivementmgr_commission_commission_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CalculateOrderCommissionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CalculateOrderCommissionRequest) ProtoMessage() {}

func (x *CalculateOrderCommissionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_archivementmgr_commission_commission_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CalculateOrderCommissionRequest.ProtoReflect.Descriptor instead.
func (*CalculateOrderCommissionRequest) Descriptor() ([]byte, []int) {
	return file_npool_archivementmgr_commission_commission_proto_rawDescGZIP(), []int{0}
}

func (x *CalculateOrderCommissionRequest) GetOrderID() string {
	if x != nil {
		return x.OrderID
	}
	return ""
}

type CalculateOrderCommissionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CalculateOrderCommissionResponse) Reset() {
	*x = CalculateOrderCommissionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_archivementmgr_commission_commission_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CalculateOrderCommissionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CalculateOrderCommissionResponse) ProtoMessage() {}

func (x *CalculateOrderCommissionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_archivementmgr_commission_commission_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CalculateOrderCommissionResponse.ProtoReflect.Descriptor instead.
func (*CalculateOrderCommissionResponse) Descriptor() ([]byte, []int) {
	return file_npool_archivementmgr_commission_commission_proto_rawDescGZIP(), []int{1}
}

type CreateAppUserGoodCommissionsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AppID  string `protobuf:"bytes,10,opt,name=AppID,proto3" json:"AppID,omitempty"`
	UserID string `protobuf:"bytes,20,opt,name=UserID,proto3" json:"UserID,omitempty"`
	GoodID string `protobuf:"bytes,30,opt,name=GoodID,proto3" json:"GoodID,omitempty"`
}

func (x *CreateAppUserGoodCommissionsRequest) Reset() {
	*x = CreateAppUserGoodCommissionsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_archivementmgr_commission_commission_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateAppUserGoodCommissionsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAppUserGoodCommissionsRequest) ProtoMessage() {}

func (x *CreateAppUserGoodCommissionsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_archivementmgr_commission_commission_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAppUserGoodCommissionsRequest.ProtoReflect.Descriptor instead.
func (*CreateAppUserGoodCommissionsRequest) Descriptor() ([]byte, []int) {
	return file_npool_archivementmgr_commission_commission_proto_rawDescGZIP(), []int{2}
}

func (x *CreateAppUserGoodCommissionsRequest) GetAppID() string {
	if x != nil {
		return x.AppID
	}
	return ""
}

func (x *CreateAppUserGoodCommissionsRequest) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

func (x *CreateAppUserGoodCommissionsRequest) GetGoodID() string {
	if x != nil {
		return x.GoodID
	}
	return ""
}

type CreateAppUserGoodCommissionsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateAppUserGoodCommissionsResponse) Reset() {
	*x = CreateAppUserGoodCommissionsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_archivementmgr_commission_commission_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateAppUserGoodCommissionsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAppUserGoodCommissionsResponse) ProtoMessage() {}

func (x *CreateAppUserGoodCommissionsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_archivementmgr_commission_commission_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAppUserGoodCommissionsResponse.ProtoReflect.Descriptor instead.
func (*CreateAppUserGoodCommissionsResponse) Descriptor() ([]byte, []int) {
	return file_npool_archivementmgr_commission_commission_proto_rawDescGZIP(), []int{3}
}

var File_npool_archivementmgr_commission_commission_proto protoreflect.FileDescriptor

var file_npool_archivementmgr_commission_commission_proto_rawDesc = []byte{
	0x0a, 0x30, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x6d,
	0x65, 0x6e, 0x74, 0x6d, 0x67, 0x72, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x21, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e,
	0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x11, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x3b, 0x0a, 0x1f, 0x43, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x65,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49,
	0x44, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x44,
	0x22, 0x22, 0x0a, 0x20, 0x43, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x6b, 0x0a, 0x23, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x70,
	0x70, 0x55, 0x73, 0x65, 0x72, 0x47, 0x6f, 0x6f, 0x64, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x41,
	0x70, 0x70, 0x49, 0x44, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x41, 0x70, 0x70, 0x49,
	0x44, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x14, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x47, 0x6f, 0x6f,
	0x64, 0x49, 0x44, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x47, 0x6f, 0x6f, 0x64, 0x49,
	0x44, 0x22, 0x26, 0x0a, 0x24, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x70, 0x70, 0x55, 0x73,
	0x65, 0x72, 0x47, 0x6f, 0x6f, 0x64, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0xd7, 0x03, 0x0a, 0x0a, 0x43, 0x6f,
	0x6d, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x3e, 0x0a, 0x07, 0x56, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x19, 0x2e, 0x6e, 0x70,
	0x6f, 0x6f, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0xa5, 0x01, 0x0a, 0x18, 0x43, 0x61, 0x6c,
	0x63, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x43, 0x6f, 0x6d, 0x6d, 0x69,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x42, 0x2e, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x6d,
	0x65, 0x6e, 0x74, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x61, 0x6c, 0x63, 0x75, 0x6c,
	0x61, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x43, 0x2e, 0x61, 0x72, 0x63, 0x68,
	0x69, 0x76, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x61,
	0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x43, 0x6f, 0x6d, 0x6d,
	0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0xe0, 0x01, 0x0a, 0x1c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x70, 0x70, 0x55, 0x73,
	0x65, 0x72, 0x47, 0x6f, 0x6f, 0x64, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x73, 0x12, 0x46, 0x2e, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e,
	0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x70, 0x70, 0x55,
	0x73, 0x65, 0x72, 0x47, 0x6f, 0x6f, 0x64, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x47, 0x2e, 0x61, 0x72, 0x63, 0x68,
	0x69, 0x76, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x41, 0x70, 0x70, 0x55, 0x73, 0x65, 0x72, 0x47, 0x6f, 0x6f, 0x64, 0x43,
	0x6f, 0x6d, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x2f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x29, 0x22, 0x24, 0x2f, 0x76, 0x31, 0x2f,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x2f, 0x61, 0x70, 0x70, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f,
	0x67, 0x6f, 0x6f, 0x64, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73,
	0x3a, 0x01, 0x2a, 0x42, 0x42, 0x5a, 0x40, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x4e, 0x70, 0x6f, 0x6f, 0x6c, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x61, 0x72,
	0x63, 0x68, 0x69, 0x76, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x6d, 0x67, 0x72, 0x2f, 0x63, 0x6f, 0x6d,
	0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_npool_archivementmgr_commission_commission_proto_rawDescOnce sync.Once
	file_npool_archivementmgr_commission_commission_proto_rawDescData = file_npool_archivementmgr_commission_commission_proto_rawDesc
)

func file_npool_archivementmgr_commission_commission_proto_rawDescGZIP() []byte {
	file_npool_archivementmgr_commission_commission_proto_rawDescOnce.Do(func() {
		file_npool_archivementmgr_commission_commission_proto_rawDescData = protoimpl.X.CompressGZIP(file_npool_archivementmgr_commission_commission_proto_rawDescData)
	})
	return file_npool_archivementmgr_commission_commission_proto_rawDescData
}

var file_npool_archivementmgr_commission_commission_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_npool_archivementmgr_commission_commission_proto_goTypes = []interface{}{
	(*CalculateOrderCommissionRequest)(nil),      // 0: archivement.manager.commission.v1.CalculateOrderCommissionRequest
	(*CalculateOrderCommissionResponse)(nil),     // 1: archivement.manager.commission.v1.CalculateOrderCommissionResponse
	(*CreateAppUserGoodCommissionsRequest)(nil),  // 2: archivement.manager.commission.v1.CreateAppUserGoodCommissionsRequest
	(*CreateAppUserGoodCommissionsResponse)(nil), // 3: archivement.manager.commission.v1.CreateAppUserGoodCommissionsResponse
	(*emptypb.Empty)(nil),                        // 4: google.protobuf.Empty
	(*npool.VersionResponse)(nil),                // 5: npool.v1.VersionResponse
}
var file_npool_archivementmgr_commission_commission_proto_depIdxs = []int32{
	4, // 0: archivement.manager.commission.v1.Commission.Version:input_type -> google.protobuf.Empty
	0, // 1: archivement.manager.commission.v1.Commission.CalculateOrderCommission:input_type -> archivement.manager.commission.v1.CalculateOrderCommissionRequest
	2, // 2: archivement.manager.commission.v1.Commission.CreateAppUserGoodCommissions:input_type -> archivement.manager.commission.v1.CreateAppUserGoodCommissionsRequest
	5, // 3: archivement.manager.commission.v1.Commission.Version:output_type -> npool.v1.VersionResponse
	1, // 4: archivement.manager.commission.v1.Commission.CalculateOrderCommission:output_type -> archivement.manager.commission.v1.CalculateOrderCommissionResponse
	3, // 5: archivement.manager.commission.v1.Commission.CreateAppUserGoodCommissions:output_type -> archivement.manager.commission.v1.CreateAppUserGoodCommissionsResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_npool_archivementmgr_commission_commission_proto_init() }
func file_npool_archivementmgr_commission_commission_proto_init() {
	if File_npool_archivementmgr_commission_commission_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_npool_archivementmgr_commission_commission_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CalculateOrderCommissionRequest); i {
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
		file_npool_archivementmgr_commission_commission_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CalculateOrderCommissionResponse); i {
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
		file_npool_archivementmgr_commission_commission_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateAppUserGoodCommissionsRequest); i {
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
		file_npool_archivementmgr_commission_commission_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateAppUserGoodCommissionsResponse); i {
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
			RawDescriptor: file_npool_archivementmgr_commission_commission_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_npool_archivementmgr_commission_commission_proto_goTypes,
		DependencyIndexes: file_npool_archivementmgr_commission_commission_proto_depIdxs,
		MessageInfos:      file_npool_archivementmgr_commission_commission_proto_msgTypes,
	}.Build()
	File_npool_archivementmgr_commission_commission_proto = out.File
	file_npool_archivementmgr_commission_commission_proto_rawDesc = nil
	file_npool_archivementmgr_commission_commission_proto_goTypes = nil
	file_npool_archivementmgr_commission_commission_proto_depIdxs = nil
}
