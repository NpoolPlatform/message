// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.18.1
// source: npool/chain/gw/v1/coin/currency/history/history.proto

package history

import (
	currency "github.com/NpoolPlatform/message/npool/chain/mw/v1/coin/currency"
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

type GetCurrenciesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CoinTypeIDs []string `protobuf:"bytes,10,rep,name=CoinTypeIDs,proto3" json:"CoinTypeIDs,omitempty"`
	Offset      int32    `protobuf:"varint,20,opt,name=Offset,proto3" json:"Offset,omitempty"`
	Limit       int32    `protobuf:"varint,30,opt,name=Limit,proto3" json:"Limit,omitempty"`
}

func (x *GetCurrenciesRequest) Reset() {
	*x = GetCurrenciesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_chain_gw_v1_coin_currency_history_history_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCurrenciesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCurrenciesRequest) ProtoMessage() {}

func (x *GetCurrenciesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_chain_gw_v1_coin_currency_history_history_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCurrenciesRequest.ProtoReflect.Descriptor instead.
func (*GetCurrenciesRequest) Descriptor() ([]byte, []int) {
	return file_npool_chain_gw_v1_coin_currency_history_history_proto_rawDescGZIP(), []int{0}
}

func (x *GetCurrenciesRequest) GetCoinTypeIDs() []string {
	if x != nil {
		return x.CoinTypeIDs
	}
	return nil
}

func (x *GetCurrenciesRequest) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *GetCurrenciesRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type GetCurrenciesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Infos []*currency.Currency `protobuf:"bytes,10,rep,name=Infos,proto3" json:"Infos,omitempty"`
	Total uint32               `protobuf:"varint,20,opt,name=Total,proto3" json:"Total,omitempty"`
}

func (x *GetCurrenciesResponse) Reset() {
	*x = GetCurrenciesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_chain_gw_v1_coin_currency_history_history_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCurrenciesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCurrenciesResponse) ProtoMessage() {}

func (x *GetCurrenciesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_chain_gw_v1_coin_currency_history_history_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCurrenciesResponse.ProtoReflect.Descriptor instead.
func (*GetCurrenciesResponse) Descriptor() ([]byte, []int) {
	return file_npool_chain_gw_v1_coin_currency_history_history_proto_rawDescGZIP(), []int{1}
}

func (x *GetCurrenciesResponse) GetInfos() []*currency.Currency {
	if x != nil {
		return x.Infos
	}
	return nil
}

func (x *GetCurrenciesResponse) GetTotal() uint32 {
	if x != nil {
		return x.Total
	}
	return 0
}

var File_npool_chain_gw_v1_coin_currency_history_history_proto protoreflect.FileDescriptor

var file_npool_chain_gw_v1_coin_currency_history_history_proto_rawDesc = []byte{
	0x0a, 0x35, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2f, 0x67, 0x77,
	0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x69, 0x6e, 0x2f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63,
	0x79, 0x2f, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x2f, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x26, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x67,
	0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x63, 0x6f, 0x69, 0x6e, 0x2e, 0x63, 0x75, 0x72, 0x72,
	0x65, 0x6e, 0x63, 0x79, 0x2e, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x1a,
	0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x6e,
	0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2f, 0x6d, 0x77, 0x2f, 0x76, 0x31,
	0x2f, 0x63, 0x6f, 0x69, 0x6e, 0x2f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x2f, 0x63,
	0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x66, 0x0a,
	0x14, 0x47, 0x65, 0x74, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x69, 0x65, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x43, 0x6f, 0x69, 0x6e, 0x54, 0x79, 0x70,
	0x65, 0x49, 0x44, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x43, 0x6f, 0x69, 0x6e,
	0x54, 0x79, 0x70, 0x65, 0x49, 0x44, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x4f, 0x66, 0x66, 0x73, 0x65,
	0x74, 0x18, 0x14, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12,
	0x14, 0x0a, 0x05, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05,
	0x4c, 0x69, 0x6d, 0x69, 0x74, 0x22, 0x70, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x43, 0x75, 0x72, 0x72,
	0x65, 0x6e, 0x63, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x41,
	0x0a, 0x05, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2b, 0x2e,
	0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65,
	0x2e, 0x63, 0x6f, 0x69, 0x6e, 0x2e, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x2e, 0x76,
	0x31, 0x2e, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x52, 0x05, 0x49, 0x6e, 0x66, 0x6f,
	0x73, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x14, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x32, 0xc2, 0x01, 0x0a, 0x07, 0x47, 0x61, 0x74, 0x65,
	0x77, 0x61, 0x79, 0x12, 0xb6, 0x01, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x43, 0x75, 0x72, 0x72, 0x65,
	0x6e, 0x63, 0x69, 0x65, 0x73, 0x12, 0x3c, 0x2e, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x67, 0x61,
	0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x63, 0x6f, 0x69, 0x6e, 0x2e, 0x63, 0x75, 0x72, 0x72, 0x65,
	0x6e, 0x63, 0x79, 0x2e, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x47,
	0x65, 0x74, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x69, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x3d, 0x2e, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x67, 0x61, 0x74, 0x65,
	0x77, 0x61, 0x79, 0x2e, 0x63, 0x6f, 0x69, 0x6e, 0x2e, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63,
	0x79, 0x2e, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74,
	0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x28, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x22, 0x3a, 0x01, 0x2a, 0x22, 0x1d, 0x2f,
	0x76, 0x31, 0x2f, 0x67, 0x65, 0x74, 0x2f, 0x63, 0x6f, 0x69, 0x6e, 0x63, 0x75, 0x72, 0x72, 0x65,
	0x6e, 0x63, 0x79, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x42, 0x4a, 0x5a, 0x48,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4e, 0x70, 0x6f, 0x6f, 0x6c,
	0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x2f, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2f, 0x67, 0x77, 0x2f,
	0x76, 0x31, 0x2f, 0x63, 0x6f, 0x69, 0x6e, 0x2f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79,
	0x2f, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_npool_chain_gw_v1_coin_currency_history_history_proto_rawDescOnce sync.Once
	file_npool_chain_gw_v1_coin_currency_history_history_proto_rawDescData = file_npool_chain_gw_v1_coin_currency_history_history_proto_rawDesc
)

func file_npool_chain_gw_v1_coin_currency_history_history_proto_rawDescGZIP() []byte {
	file_npool_chain_gw_v1_coin_currency_history_history_proto_rawDescOnce.Do(func() {
		file_npool_chain_gw_v1_coin_currency_history_history_proto_rawDescData = protoimpl.X.CompressGZIP(file_npool_chain_gw_v1_coin_currency_history_history_proto_rawDescData)
	})
	return file_npool_chain_gw_v1_coin_currency_history_history_proto_rawDescData
}

var file_npool_chain_gw_v1_coin_currency_history_history_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_npool_chain_gw_v1_coin_currency_history_history_proto_goTypes = []interface{}{
	(*GetCurrenciesRequest)(nil),  // 0: chain.gateway.coin.currency.history.v1.GetCurrenciesRequest
	(*GetCurrenciesResponse)(nil), // 1: chain.gateway.coin.currency.history.v1.GetCurrenciesResponse
	(*currency.Currency)(nil),     // 2: chain.middleware.coin.currency.v1.Currency
}
var file_npool_chain_gw_v1_coin_currency_history_history_proto_depIdxs = []int32{
	2, // 0: chain.gateway.coin.currency.history.v1.GetCurrenciesResponse.Infos:type_name -> chain.middleware.coin.currency.v1.Currency
	0, // 1: chain.gateway.coin.currency.history.v1.Gateway.GetCurrencies:input_type -> chain.gateway.coin.currency.history.v1.GetCurrenciesRequest
	1, // 2: chain.gateway.coin.currency.history.v1.Gateway.GetCurrencies:output_type -> chain.gateway.coin.currency.history.v1.GetCurrenciesResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_npool_chain_gw_v1_coin_currency_history_history_proto_init() }
func file_npool_chain_gw_v1_coin_currency_history_history_proto_init() {
	if File_npool_chain_gw_v1_coin_currency_history_history_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_npool_chain_gw_v1_coin_currency_history_history_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCurrenciesRequest); i {
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
		file_npool_chain_gw_v1_coin_currency_history_history_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCurrenciesResponse); i {
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
			RawDescriptor: file_npool_chain_gw_v1_coin_currency_history_history_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_npool_chain_gw_v1_coin_currency_history_history_proto_goTypes,
		DependencyIndexes: file_npool_chain_gw_v1_coin_currency_history_history_proto_depIdxs,
		MessageInfos:      file_npool_chain_gw_v1_coin_currency_history_history_proto_msgTypes,
	}.Build()
	File_npool_chain_gw_v1_coin_currency_history_history_proto = out.File
	file_npool_chain_gw_v1_coin_currency_history_history_proto_rawDesc = nil
	file_npool_chain_gw_v1_coin_currency_history_history_proto_goTypes = nil
	file_npool_chain_gw_v1_coin_currency_history_history_proto_depIdxs = nil
}
