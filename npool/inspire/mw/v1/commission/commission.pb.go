// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.18.1
// source: npool/inspire/mw/v1/commission/commission.proto

package commission

import (
	_ "github.com/NpoolPlatform/message/npool"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/emptypb"
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
		mi := &file_npool_inspire_mw_v1_commission_commission_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CalculateOrderCommissionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CalculateOrderCommissionRequest) ProtoMessage() {}

func (x *CalculateOrderCommissionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_inspire_mw_v1_commission_commission_proto_msgTypes[0]
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
	return file_npool_inspire_mw_v1_commission_commission_proto_rawDescGZIP(), []int{0}
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
		mi := &file_npool_inspire_mw_v1_commission_commission_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CalculateOrderCommissionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CalculateOrderCommissionResponse) ProtoMessage() {}

func (x *CalculateOrderCommissionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_inspire_mw_v1_commission_commission_proto_msgTypes[1]
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
	return file_npool_inspire_mw_v1_commission_commission_proto_rawDescGZIP(), []int{1}
}

type CreateUserGoodCommissionsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AppID  string `protobuf:"bytes,10,opt,name=AppID,proto3" json:"AppID,omitempty"`
	UserID string `protobuf:"bytes,20,opt,name=UserID,proto3" json:"UserID,omitempty"`
	GoodID string `protobuf:"bytes,30,opt,name=GoodID,proto3" json:"GoodID,omitempty"`
}

func (x *CreateUserGoodCommissionsRequest) Reset() {
	*x = CreateUserGoodCommissionsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_inspire_mw_v1_commission_commission_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateUserGoodCommissionsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUserGoodCommissionsRequest) ProtoMessage() {}

func (x *CreateUserGoodCommissionsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_inspire_mw_v1_commission_commission_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateUserGoodCommissionsRequest.ProtoReflect.Descriptor instead.
func (*CreateUserGoodCommissionsRequest) Descriptor() ([]byte, []int) {
	return file_npool_inspire_mw_v1_commission_commission_proto_rawDescGZIP(), []int{2}
}

func (x *CreateUserGoodCommissionsRequest) GetAppID() string {
	if x != nil {
		return x.AppID
	}
	return ""
}

func (x *CreateUserGoodCommissionsRequest) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

func (x *CreateUserGoodCommissionsRequest) GetGoodID() string {
	if x != nil {
		return x.GoodID
	}
	return ""
}

type CreateUserGoodCommissionsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateUserGoodCommissionsResponse) Reset() {
	*x = CreateUserGoodCommissionsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_inspire_mw_v1_commission_commission_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateUserGoodCommissionsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUserGoodCommissionsResponse) ProtoMessage() {}

func (x *CreateUserGoodCommissionsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_inspire_mw_v1_commission_commission_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateUserGoodCommissionsResponse.ProtoReflect.Descriptor instead.
func (*CreateUserGoodCommissionsResponse) Descriptor() ([]byte, []int) {
	return file_npool_inspire_mw_v1_commission_commission_proto_rawDescGZIP(), []int{3}
}

var File_npool_inspire_mw_v1_commission_commission_proto protoreflect.FileDescriptor

var file_npool_inspire_mw_v1_commission_commission_proto_rawDesc = []byte{
	0x0a, 0x2f, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x69, 0x6e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2f,
	0x6d, 0x77, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x20, 0x69, 0x6e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c,
	0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11,
	0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x3b, 0x0a, 0x1f, 0x43, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x44, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x44, 0x22, 0x22,
	0x0a, 0x20, 0x43, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x68, 0x0a, 0x20, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72,
	0x47, 0x6f, 0x6f, 0x64, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x41, 0x70, 0x70, 0x49, 0x44, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x41, 0x70, 0x70, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06,
	0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x55, 0x73,
	0x65, 0x72, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x47, 0x6f, 0x6f, 0x64, 0x49, 0x44, 0x18, 0x1e,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x47, 0x6f, 0x6f, 0x64, 0x49, 0x44, 0x22, 0x23, 0x0a, 0x21,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x47, 0x6f, 0x6f, 0x64, 0x43, 0x6f,
	0x6d, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x32, 0xdd, 0x02, 0x0a, 0x0c, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x4d, 0x77, 0x12, 0xa3, 0x01, 0x0a, 0x18, 0x43, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x65,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12,
	0x41, 0x2e, 0x69, 0x6e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65,
	0x77, 0x61, 0x72, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e,
	0x76, 0x31, 0x2e, 0x43, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x42, 0x2e, 0x69, 0x6e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x6d, 0x69, 0x64,
	0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0xa6, 0x01, 0x0a, 0x19, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x47, 0x6f, 0x6f, 0x64, 0x43, 0x6f, 0x6d, 0x6d, 0x69,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x42, 0x2e, 0x69, 0x6e, 0x73, 0x70, 0x69, 0x72, 0x65,
	0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x55, 0x73, 0x65, 0x72, 0x47, 0x6f, 0x6f, 0x64, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x43, 0x2e, 0x69, 0x6e, 0x73,
	0x70, 0x69, 0x72, 0x65, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x47, 0x6f, 0x6f, 0x64, 0x43, 0x6f, 0x6d, 0x6d,
	0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x42, 0x41, 0x5a, 0x3f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x4e, 0x70, 0x6f, 0x6f, 0x6c, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x69, 0x6e, 0x73, 0x70,
	0x69, 0x72, 0x65, 0x2f, 0x6d, 0x77, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_npool_inspire_mw_v1_commission_commission_proto_rawDescOnce sync.Once
	file_npool_inspire_mw_v1_commission_commission_proto_rawDescData = file_npool_inspire_mw_v1_commission_commission_proto_rawDesc
)

func file_npool_inspire_mw_v1_commission_commission_proto_rawDescGZIP() []byte {
	file_npool_inspire_mw_v1_commission_commission_proto_rawDescOnce.Do(func() {
		file_npool_inspire_mw_v1_commission_commission_proto_rawDescData = protoimpl.X.CompressGZIP(file_npool_inspire_mw_v1_commission_commission_proto_rawDescData)
	})
	return file_npool_inspire_mw_v1_commission_commission_proto_rawDescData
}

var file_npool_inspire_mw_v1_commission_commission_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_npool_inspire_mw_v1_commission_commission_proto_goTypes = []interface{}{
	(*CalculateOrderCommissionRequest)(nil),   // 0: inspire.middleware.commission.v1.CalculateOrderCommissionRequest
	(*CalculateOrderCommissionResponse)(nil),  // 1: inspire.middleware.commission.v1.CalculateOrderCommissionResponse
	(*CreateUserGoodCommissionsRequest)(nil),  // 2: inspire.middleware.commission.v1.CreateUserGoodCommissionsRequest
	(*CreateUserGoodCommissionsResponse)(nil), // 3: inspire.middleware.commission.v1.CreateUserGoodCommissionsResponse
}
var file_npool_inspire_mw_v1_commission_commission_proto_depIdxs = []int32{
	0, // 0: inspire.middleware.commission.v1.CommissionMw.CalculateOrderCommission:input_type -> inspire.middleware.commission.v1.CalculateOrderCommissionRequest
	2, // 1: inspire.middleware.commission.v1.CommissionMw.CreateUserGoodCommissions:input_type -> inspire.middleware.commission.v1.CreateUserGoodCommissionsRequest
	1, // 2: inspire.middleware.commission.v1.CommissionMw.CalculateOrderCommission:output_type -> inspire.middleware.commission.v1.CalculateOrderCommissionResponse
	3, // 3: inspire.middleware.commission.v1.CommissionMw.CreateUserGoodCommissions:output_type -> inspire.middleware.commission.v1.CreateUserGoodCommissionsResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_npool_inspire_mw_v1_commission_commission_proto_init() }
func file_npool_inspire_mw_v1_commission_commission_proto_init() {
	if File_npool_inspire_mw_v1_commission_commission_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_npool_inspire_mw_v1_commission_commission_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_npool_inspire_mw_v1_commission_commission_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_npool_inspire_mw_v1_commission_commission_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateUserGoodCommissionsRequest); i {
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
		file_npool_inspire_mw_v1_commission_commission_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateUserGoodCommissionsResponse); i {
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
			RawDescriptor: file_npool_inspire_mw_v1_commission_commission_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_npool_inspire_mw_v1_commission_commission_proto_goTypes,
		DependencyIndexes: file_npool_inspire_mw_v1_commission_commission_proto_depIdxs,
		MessageInfos:      file_npool_inspire_mw_v1_commission_commission_proto_msgTypes,
	}.Build()
	File_npool_inspire_mw_v1_commission_commission_proto = out.File
	file_npool_inspire_mw_v1_commission_commission_proto_rawDesc = nil
	file_npool_inspire_mw_v1_commission_commission_proto_goTypes = nil
	file_npool_inspire_mw_v1_commission_commission_proto_depIdxs = nil
}
