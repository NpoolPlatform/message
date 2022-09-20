// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.18.1
// source: npool/third/gw/v1/usedfor/usedfor.proto

package usedfor

import (
	usedfor "github.com/NpoolPlatform/message/npool/third/mgr/v1/usedfor"
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

type GetUsedForRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetUsedForRequest) Reset() {
	*x = GetUsedForRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_third_gw_v1_usedfor_usedfor_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUsedForRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUsedForRequest) ProtoMessage() {}

func (x *GetUsedForRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_third_gw_v1_usedfor_usedfor_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUsedForRequest.ProtoReflect.Descriptor instead.
func (*GetUsedForRequest) Descriptor() ([]byte, []int) {
	return file_npool_third_gw_v1_usedfor_usedfor_proto_rawDescGZIP(), []int{0}
}

type GetUsedForResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UsedFor []usedfor.UsedFor `protobuf:"varint,10,rep,packed,name=UsedFor,proto3,enum=third.manager.usedfor.v1.UsedFor" json:"UsedFor,omitempty"`
}

func (x *GetUsedForResponse) Reset() {
	*x = GetUsedForResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_third_gw_v1_usedfor_usedfor_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUsedForResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUsedForResponse) ProtoMessage() {}

func (x *GetUsedForResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_third_gw_v1_usedfor_usedfor_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUsedForResponse.ProtoReflect.Descriptor instead.
func (*GetUsedForResponse) Descriptor() ([]byte, []int) {
	return file_npool_third_gw_v1_usedfor_usedfor_proto_rawDescGZIP(), []int{1}
}

func (x *GetUsedForResponse) GetUsedFor() []usedfor.UsedFor {
	if x != nil {
		return x.UsedFor
	}
	return nil
}

var File_npool_third_gw_v1_usedfor_usedfor_proto protoreflect.FileDescriptor

var file_npool_third_gw_v1_usedfor_usedfor_proto_rawDesc = []byte{
	0x0a, 0x27, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x74, 0x68, 0x69, 0x72, 0x64, 0x2f, 0x67, 0x77,
	0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x64, 0x66, 0x6f, 0x72, 0x2f, 0x75, 0x73, 0x65, 0x64,
	0x66, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18, 0x74, 0x68, 0x69, 0x72, 0x64,
	0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x75, 0x73, 0x65, 0x64, 0x66, 0x6f, 0x72,
	0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x29, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x74, 0x68, 0x69, 0x72, 0x64, 0x2f, 0x6d,
	0x67, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x64, 0x66, 0x6f, 0x72, 0x2f, 0x75, 0x73,
	0x65, 0x72, 0x64, 0x66, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x13, 0x0a, 0x11,
	0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x64, 0x46, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x22, 0x51, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x64, 0x46, 0x6f, 0x72, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3b, 0x0a, 0x07, 0x55, 0x73, 0x65, 0x64, 0x46,
	0x6f, 0x72, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x21, 0x2e, 0x74, 0x68, 0x69, 0x72, 0x64,
	0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x75, 0x73, 0x65, 0x64, 0x66, 0x6f, 0x72,
	0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x64, 0x46, 0x6f, 0x72, 0x52, 0x07, 0x55, 0x73, 0x65,
	0x64, 0x46, 0x6f, 0x72, 0x32, 0x8f, 0x01, 0x0a, 0x07, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79,
	0x12, 0x83, 0x01, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x64, 0x46, 0x6f, 0x72, 0x12,
	0x2b, 0x2e, 0x74, 0x68, 0x69, 0x72, 0x64, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e,
	0x75, 0x73, 0x65, 0x64, 0x66, 0x6f, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73,
	0x65, 0x64, 0x46, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2c, 0x2e, 0x74,
	0x68, 0x69, 0x72, 0x64, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x75, 0x73, 0x65,
	0x64, 0x66, 0x6f, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x64, 0x46,
	0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1a, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x14, 0x22, 0x0f, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x65, 0x74, 0x2f, 0x75, 0x73, 0x65, 0x64,
	0x66, 0x6f, 0x72, 0x3a, 0x01, 0x2a, 0x42, 0x3c, 0x5a, 0x3a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4e, 0x70, 0x6f, 0x6f, 0x6c, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f,
	0x72, 0x6d, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x6e, 0x70, 0x6f, 0x6f, 0x6c,
	0x2f, 0x74, 0x68, 0x69, 0x72, 0x64, 0x2f, 0x67, 0x77, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65,
	0x64, 0x66, 0x6f, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_npool_third_gw_v1_usedfor_usedfor_proto_rawDescOnce sync.Once
	file_npool_third_gw_v1_usedfor_usedfor_proto_rawDescData = file_npool_third_gw_v1_usedfor_usedfor_proto_rawDesc
)

func file_npool_third_gw_v1_usedfor_usedfor_proto_rawDescGZIP() []byte {
	file_npool_third_gw_v1_usedfor_usedfor_proto_rawDescOnce.Do(func() {
		file_npool_third_gw_v1_usedfor_usedfor_proto_rawDescData = protoimpl.X.CompressGZIP(file_npool_third_gw_v1_usedfor_usedfor_proto_rawDescData)
	})
	return file_npool_third_gw_v1_usedfor_usedfor_proto_rawDescData
}

var file_npool_third_gw_v1_usedfor_usedfor_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_npool_third_gw_v1_usedfor_usedfor_proto_goTypes = []interface{}{
	(*GetUsedForRequest)(nil),  // 0: third.gateway.usedfor.v1.GetUsedForRequest
	(*GetUsedForResponse)(nil), // 1: third.gateway.usedfor.v1.GetUsedForResponse
	(usedfor.UsedFor)(0),       // 2: third.manager.usedfor.v1.UsedFor
}
var file_npool_third_gw_v1_usedfor_usedfor_proto_depIdxs = []int32{
	2, // 0: third.gateway.usedfor.v1.GetUsedForResponse.UsedFor:type_name -> third.manager.usedfor.v1.UsedFor
	0, // 1: third.gateway.usedfor.v1.Gateway.GetUsedFor:input_type -> third.gateway.usedfor.v1.GetUsedForRequest
	1, // 2: third.gateway.usedfor.v1.Gateway.GetUsedFor:output_type -> third.gateway.usedfor.v1.GetUsedForResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_npool_third_gw_v1_usedfor_usedfor_proto_init() }
func file_npool_third_gw_v1_usedfor_usedfor_proto_init() {
	if File_npool_third_gw_v1_usedfor_usedfor_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_npool_third_gw_v1_usedfor_usedfor_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUsedForRequest); i {
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
		file_npool_third_gw_v1_usedfor_usedfor_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUsedForResponse); i {
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
			RawDescriptor: file_npool_third_gw_v1_usedfor_usedfor_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_npool_third_gw_v1_usedfor_usedfor_proto_goTypes,
		DependencyIndexes: file_npool_third_gw_v1_usedfor_usedfor_proto_depIdxs,
		MessageInfos:      file_npool_third_gw_v1_usedfor_usedfor_proto_msgTypes,
	}.Build()
	File_npool_third_gw_v1_usedfor_usedfor_proto = out.File
	file_npool_third_gw_v1_usedfor_usedfor_proto_rawDesc = nil
	file_npool_third_gw_v1_usedfor_usedfor_proto_goTypes = nil
	file_npool_third_gw_v1_usedfor_usedfor_proto_depIdxs = nil
}
