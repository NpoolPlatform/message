// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.18.1
// source: npool/good/mw/v1/good/stock/stock.proto

package stock

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

type MiningGoodStockReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MiningPoolID   *string `protobuf:"bytes,10,opt,name=MiningPoolID,proto3,oneof" json:"MiningPoolID,omitempty"`
	PoolGoodUserID *string `protobuf:"bytes,20,opt,name=PoolGoodUserID,proto3,oneof" json:"PoolGoodUserID,omitempty"`
	Total          *string `protobuf:"bytes,30,opt,name=Total,proto3,oneof" json:"Total,omitempty"`
}

func (x *MiningGoodStockReq) Reset() {
	*x = MiningGoodStockReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_good_mw_v1_good_stock_stock_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MiningGoodStockReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MiningGoodStockReq) ProtoMessage() {}

func (x *MiningGoodStockReq) ProtoReflect() protoreflect.Message {
	mi := &file_npool_good_mw_v1_good_stock_stock_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MiningGoodStockReq.ProtoReflect.Descriptor instead.
func (*MiningGoodStockReq) Descriptor() ([]byte, []int) {
	return file_npool_good_mw_v1_good_stock_stock_proto_rawDescGZIP(), []int{0}
}

func (x *MiningGoodStockReq) GetMiningPoolID() string {
	if x != nil && x.MiningPoolID != nil {
		return *x.MiningPoolID
	}
	return ""
}

func (x *MiningGoodStockReq) GetPoolGoodUserID() string {
	if x != nil && x.PoolGoodUserID != nil {
		return *x.PoolGoodUserID
	}
	return ""
}

func (x *MiningGoodStockReq) GetTotal() string {
	if x != nil && x.Total != nil {
		return *x.Total
	}
	return ""
}

type MiningGoodStock struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: sql:"good_stock_id"
	GoodStockID string `protobuf:"bytes,20,opt,name=GoodStockID,proto3" json:"GoodStockID,omitempty" sql:"good_stock_id"`
}

func (x *MiningGoodStock) Reset() {
	*x = MiningGoodStock{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_good_mw_v1_good_stock_stock_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MiningGoodStock) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MiningGoodStock) ProtoMessage() {}

func (x *MiningGoodStock) ProtoReflect() protoreflect.Message {
	mi := &file_npool_good_mw_v1_good_stock_stock_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MiningGoodStock.ProtoReflect.Descriptor instead.
func (*MiningGoodStock) Descriptor() ([]byte, []int) {
	return file_npool_good_mw_v1_good_stock_stock_proto_rawDescGZIP(), []int{1}
}

func (x *MiningGoodStock) GetGoodStockID() string {
	if x != nil {
		return x.GoodStockID
	}
	return ""
}

var File_npool_good_mw_v1_good_stock_stock_proto protoreflect.FileDescriptor

var file_npool_good_mw_v1_good_stock_stock_proto_rawDesc = []byte{
	0x0a, 0x27, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x67, 0x6f, 0x6f, 0x64, 0x2f, 0x6d, 0x77, 0x2f,
	0x76, 0x31, 0x2f, 0x67, 0x6f, 0x6f, 0x64, 0x2f, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x2f, 0x73, 0x74,
	0x6f, 0x63, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1e, 0x67, 0x6f, 0x6f, 0x64, 0x2e,
	0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x67, 0x6f, 0x6f, 0x64, 0x31,
	0x2e, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x2e, 0x76, 0x31, 0x22, 0xb3, 0x01, 0x0a, 0x12, 0x4d, 0x69,
	0x6e, 0x69, 0x6e, 0x67, 0x47, 0x6f, 0x6f, 0x64, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x71,
	0x12, 0x27, 0x0a, 0x0c, 0x4d, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x50, 0x6f, 0x6f, 0x6c, 0x49, 0x44,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0c, 0x4d, 0x69, 0x6e, 0x69, 0x6e, 0x67,
	0x50, 0x6f, 0x6f, 0x6c, 0x49, 0x44, 0x88, 0x01, 0x01, 0x12, 0x2b, 0x0a, 0x0e, 0x50, 0x6f, 0x6f,
	0x6c, 0x47, 0x6f, 0x6f, 0x64, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x14, 0x20, 0x01, 0x28,
	0x09, 0x48, 0x01, 0x52, 0x0e, 0x50, 0x6f, 0x6f, 0x6c, 0x47, 0x6f, 0x6f, 0x64, 0x55, 0x73, 0x65,
	0x72, 0x49, 0x44, 0x88, 0x01, 0x01, 0x12, 0x19, 0x0a, 0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x18,
	0x1e, 0x20, 0x01, 0x28, 0x09, 0x48, 0x02, 0x52, 0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x88, 0x01,
	0x01, 0x42, 0x0f, 0x0a, 0x0d, 0x5f, 0x4d, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x50, 0x6f, 0x6f, 0x6c,
	0x49, 0x44, 0x42, 0x11, 0x0a, 0x0f, 0x5f, 0x50, 0x6f, 0x6f, 0x6c, 0x47, 0x6f, 0x6f, 0x64, 0x55,
	0x73, 0x65, 0x72, 0x49, 0x44, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x22,
	0x33, 0x0a, 0x0f, 0x4d, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x47, 0x6f, 0x6f, 0x64, 0x53, 0x74, 0x6f,
	0x63, 0x6b, 0x12, 0x20, 0x0a, 0x0b, 0x47, 0x6f, 0x6f, 0x64, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x49,
	0x44, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x47, 0x6f, 0x6f, 0x64, 0x53, 0x74, 0x6f,
	0x63, 0x6b, 0x49, 0x44, 0x42, 0x3e, 0x5a, 0x3c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x4e, 0x70, 0x6f, 0x6f, 0x6c, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d,
	0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x67,
	0x6f, 0x6f, 0x64, 0x2f, 0x6d, 0x77, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x6f, 0x6f, 0x64, 0x2f, 0x73,
	0x74, 0x6f, 0x63, 0x6b, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_npool_good_mw_v1_good_stock_stock_proto_rawDescOnce sync.Once
	file_npool_good_mw_v1_good_stock_stock_proto_rawDescData = file_npool_good_mw_v1_good_stock_stock_proto_rawDesc
)

func file_npool_good_mw_v1_good_stock_stock_proto_rawDescGZIP() []byte {
	file_npool_good_mw_v1_good_stock_stock_proto_rawDescOnce.Do(func() {
		file_npool_good_mw_v1_good_stock_stock_proto_rawDescData = protoimpl.X.CompressGZIP(file_npool_good_mw_v1_good_stock_stock_proto_rawDescData)
	})
	return file_npool_good_mw_v1_good_stock_stock_proto_rawDescData
}

var file_npool_good_mw_v1_good_stock_stock_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_npool_good_mw_v1_good_stock_stock_proto_goTypes = []interface{}{
	(*MiningGoodStockReq)(nil), // 0: good.middleware.good1.stock.v1.MiningGoodStockReq
	(*MiningGoodStock)(nil),    // 1: good.middleware.good1.stock.v1.MiningGoodStock
}
var file_npool_good_mw_v1_good_stock_stock_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_npool_good_mw_v1_good_stock_stock_proto_init() }
func file_npool_good_mw_v1_good_stock_stock_proto_init() {
	if File_npool_good_mw_v1_good_stock_stock_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_npool_good_mw_v1_good_stock_stock_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MiningGoodStockReq); i {
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
		file_npool_good_mw_v1_good_stock_stock_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MiningGoodStock); i {
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
	file_npool_good_mw_v1_good_stock_stock_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_npool_good_mw_v1_good_stock_stock_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_npool_good_mw_v1_good_stock_stock_proto_goTypes,
		DependencyIndexes: file_npool_good_mw_v1_good_stock_stock_proto_depIdxs,
		MessageInfos:      file_npool_good_mw_v1_good_stock_stock_proto_msgTypes,
	}.Build()
	File_npool_good_mw_v1_good_stock_stock_proto = out.File
	file_npool_good_mw_v1_good_stock_stock_proto_rawDesc = nil
	file_npool_good_mw_v1_good_stock_stock_proto_goTypes = nil
	file_npool_good_mw_v1_good_stock_stock_proto_depIdxs = nil
}
