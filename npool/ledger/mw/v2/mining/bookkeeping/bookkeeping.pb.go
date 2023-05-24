// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.18.1
// source: npool/ledger/mw/v2/mining/bookkeeping/bookkeeping.proto

package bookkeeping

import (
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

type BookKeepingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GoodID                    string `protobuf:"bytes,10,opt,name=GoodID,proto3" json:"GoodID,omitempty"`
	CoinTypeID                string `protobuf:"bytes,20,opt,name=CoinTypeID,proto3" json:"CoinTypeID,omitempty"`
	TotalAmount               string `protobuf:"bytes,30,opt,name=TotalAmount,proto3" json:"TotalAmount,omitempty"`
	UnsoldAmount              string `protobuf:"bytes,40,opt,name=UnsoldAmount,proto3" json:"UnsoldAmount,omitempty"`
	TechniqueServiceFeeAmount string `protobuf:"bytes,50,opt,name=TechniqueServiceFeeAmount,proto3" json:"TechniqueServiceFeeAmount,omitempty"`
	BenefitDate               uint32 `protobuf:"varint,60,opt,name=BenefitDate,proto3" json:"BenefitDate,omitempty"`
}

func (x *BookKeepingRequest) Reset() {
	*x = BookKeepingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_ledger_mw_v2_mining_bookkeeping_bookkeeping_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BookKeepingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BookKeepingRequest) ProtoMessage() {}

func (x *BookKeepingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_ledger_mw_v2_mining_bookkeeping_bookkeeping_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BookKeepingRequest.ProtoReflect.Descriptor instead.
func (*BookKeepingRequest) Descriptor() ([]byte, []int) {
	return file_npool_ledger_mw_v2_mining_bookkeeping_bookkeeping_proto_rawDescGZIP(), []int{0}
}

func (x *BookKeepingRequest) GetGoodID() string {
	if x != nil {
		return x.GoodID
	}
	return ""
}

func (x *BookKeepingRequest) GetCoinTypeID() string {
	if x != nil {
		return x.CoinTypeID
	}
	return ""
}

func (x *BookKeepingRequest) GetTotalAmount() string {
	if x != nil {
		return x.TotalAmount
	}
	return ""
}

func (x *BookKeepingRequest) GetUnsoldAmount() string {
	if x != nil {
		return x.UnsoldAmount
	}
	return ""
}

func (x *BookKeepingRequest) GetTechniqueServiceFeeAmount() string {
	if x != nil {
		return x.TechniqueServiceFeeAmount
	}
	return ""
}

func (x *BookKeepingRequest) GetBenefitDate() uint32 {
	if x != nil {
		return x.BenefitDate
	}
	return 0
}

type BookKeepingResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *BookKeepingResponse) Reset() {
	*x = BookKeepingResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_ledger_mw_v2_mining_bookkeeping_bookkeeping_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BookKeepingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BookKeepingResponse) ProtoMessage() {}

func (x *BookKeepingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_ledger_mw_v2_mining_bookkeeping_bookkeeping_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BookKeepingResponse.ProtoReflect.Descriptor instead.
func (*BookKeepingResponse) Descriptor() ([]byte, []int) {
	return file_npool_ledger_mw_v2_mining_bookkeeping_bookkeeping_proto_rawDescGZIP(), []int{1}
}

var File_npool_ledger_mw_v2_mining_bookkeeping_bookkeeping_proto protoreflect.FileDescriptor

var file_npool_ledger_mw_v2_mining_bookkeeping_bookkeeping_proto_rawDesc = []byte{
	0x0a, 0x37, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2f, 0x6d,
	0x77, 0x2f, 0x76, 0x32, 0x2f, 0x6d, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x2f, 0x62, 0x6f, 0x6f, 0x6b,
	0x6b, 0x65, 0x65, 0x70, 0x69, 0x6e, 0x67, 0x2f, 0x62, 0x6f, 0x6f, 0x6b, 0x6b, 0x65, 0x65, 0x70,
	0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x27, 0x6c, 0x65, 0x64, 0x67, 0x65,
	0x72, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x6d, 0x69, 0x6e,
	0x69, 0x6e, 0x67, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x6b, 0x65, 0x65, 0x70, 0x69, 0x6e, 0x67, 0x2e,
	0x76, 0x32, 0x22, 0xf2, 0x01, 0x0a, 0x12, 0x42, 0x6f, 0x6f, 0x6b, 0x4b, 0x65, 0x65, 0x70, 0x69,
	0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x47, 0x6f, 0x6f,
	0x64, 0x49, 0x44, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x47, 0x6f, 0x6f, 0x64, 0x49,
	0x44, 0x12, 0x1e, 0x0a, 0x0a, 0x43, 0x6f, 0x69, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x49, 0x44, 0x18,
	0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x43, 0x6f, 0x69, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x49,
	0x44, 0x12, 0x20, 0x0a, 0x0b, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x1e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x41, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x22, 0x0a, 0x0c, 0x55, 0x6e, 0x73, 0x6f, 0x6c, 0x64, 0x41, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x28, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x55, 0x6e, 0x73, 0x6f, 0x6c,
	0x64, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x3c, 0x0a, 0x19, 0x54, 0x65, 0x63, 0x68, 0x6e,
	0x69, 0x71, 0x75, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x46, 0x65, 0x65, 0x41, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x32, 0x20, 0x01, 0x28, 0x09, 0x52, 0x19, 0x54, 0x65, 0x63, 0x68,
	0x6e, 0x69, 0x71, 0x75, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x46, 0x65, 0x65, 0x41,
	0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x42, 0x65, 0x6e, 0x65, 0x66, 0x69, 0x74,
	0x44, 0x61, 0x74, 0x65, 0x18, 0x3c, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x42, 0x65, 0x6e, 0x65,
	0x66, 0x69, 0x74, 0x44, 0x61, 0x74, 0x65, 0x22, 0x15, 0x0a, 0x13, 0x42, 0x6f, 0x6f, 0x6b, 0x4b,
	0x65, 0x65, 0x70, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0x99,
	0x01, 0x0a, 0x0a, 0x4d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x12, 0x8a, 0x01,
	0x0a, 0x0b, 0x42, 0x6f, 0x6f, 0x6b, 0x4b, 0x65, 0x65, 0x70, 0x69, 0x6e, 0x67, 0x12, 0x3b, 0x2e,
	0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72,
	0x65, 0x2e, 0x6d, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x6b, 0x65, 0x65,
	0x70, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x32, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x4b, 0x65, 0x65, 0x70,
	0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3c, 0x2e, 0x6c, 0x65, 0x64,
	0x67, 0x65, 0x72, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x6d,
	0x69, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x6b, 0x65, 0x65, 0x70, 0x69, 0x6e,
	0x67, 0x2e, 0x76, 0x32, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x4b, 0x65, 0x65, 0x70, 0x69, 0x6e, 0x67,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x48, 0x5a, 0x46, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4e, 0x70, 0x6f, 0x6f, 0x6c, 0x50, 0x6c,
	0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x6e,
	0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2f, 0x6d, 0x77, 0x2f, 0x76,
	0x32, 0x2f, 0x6d, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x2f, 0x62, 0x6f, 0x6f, 0x6b, 0x6b, 0x65, 0x65,
	0x70, 0x69, 0x6e, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_npool_ledger_mw_v2_mining_bookkeeping_bookkeeping_proto_rawDescOnce sync.Once
	file_npool_ledger_mw_v2_mining_bookkeeping_bookkeeping_proto_rawDescData = file_npool_ledger_mw_v2_mining_bookkeeping_bookkeeping_proto_rawDesc
)

func file_npool_ledger_mw_v2_mining_bookkeeping_bookkeeping_proto_rawDescGZIP() []byte {
	file_npool_ledger_mw_v2_mining_bookkeeping_bookkeeping_proto_rawDescOnce.Do(func() {
		file_npool_ledger_mw_v2_mining_bookkeeping_bookkeeping_proto_rawDescData = protoimpl.X.CompressGZIP(file_npool_ledger_mw_v2_mining_bookkeeping_bookkeeping_proto_rawDescData)
	})
	return file_npool_ledger_mw_v2_mining_bookkeeping_bookkeeping_proto_rawDescData
}

var file_npool_ledger_mw_v2_mining_bookkeeping_bookkeeping_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_npool_ledger_mw_v2_mining_bookkeeping_bookkeeping_proto_goTypes = []interface{}{
	(*BookKeepingRequest)(nil),  // 0: ledger.middleware.mining.bookkeeping.v2.BookKeepingRequest
	(*BookKeepingResponse)(nil), // 1: ledger.middleware.mining.bookkeeping.v2.BookKeepingResponse
}
var file_npool_ledger_mw_v2_mining_bookkeeping_bookkeeping_proto_depIdxs = []int32{
	0, // 0: ledger.middleware.mining.bookkeeping.v2.Middleware.BookKeeping:input_type -> ledger.middleware.mining.bookkeeping.v2.BookKeepingRequest
	1, // 1: ledger.middleware.mining.bookkeeping.v2.Middleware.BookKeeping:output_type -> ledger.middleware.mining.bookkeeping.v2.BookKeepingResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_npool_ledger_mw_v2_mining_bookkeeping_bookkeeping_proto_init() }
func file_npool_ledger_mw_v2_mining_bookkeeping_bookkeeping_proto_init() {
	if File_npool_ledger_mw_v2_mining_bookkeeping_bookkeeping_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_npool_ledger_mw_v2_mining_bookkeeping_bookkeeping_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BookKeepingRequest); i {
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
		file_npool_ledger_mw_v2_mining_bookkeeping_bookkeeping_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BookKeepingResponse); i {
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
			RawDescriptor: file_npool_ledger_mw_v2_mining_bookkeeping_bookkeeping_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_npool_ledger_mw_v2_mining_bookkeeping_bookkeeping_proto_goTypes,
		DependencyIndexes: file_npool_ledger_mw_v2_mining_bookkeeping_bookkeeping_proto_depIdxs,
		MessageInfos:      file_npool_ledger_mw_v2_mining_bookkeeping_bookkeeping_proto_msgTypes,
	}.Build()
	File_npool_ledger_mw_v2_mining_bookkeeping_bookkeeping_proto = out.File
	file_npool_ledger_mw_v2_mining_bookkeeping_bookkeeping_proto_rawDesc = nil
	file_npool_ledger_mw_v2_mining_bookkeeping_bookkeeping_proto_goTypes = nil
	file_npool_ledger_mw_v2_mining_bookkeeping_bookkeeping_proto_depIdxs = nil
}
