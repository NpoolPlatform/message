// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.18.1
// source: npool/inspire/gw/v1/reconcile/reconcile.proto

package reconciliation

import (
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

type ReconcileRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AppID     string `protobuf:"bytes,10,opt,name=AppID,proto3" json:"AppID,omitempty"`
	AppGoodID string `protobuf:"bytes,20,opt,name=AppGoodID,proto3" json:"AppGoodID,omitempty"`
}

func (x *ReconcileRequest) Reset() {
	*x = ReconcileRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_inspire_gw_v1_reconcile_reconcile_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReconcileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReconcileRequest) ProtoMessage() {}

func (x *ReconcileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_inspire_gw_v1_reconcile_reconcile_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReconcileRequest.ProtoReflect.Descriptor instead.
func (*ReconcileRequest) Descriptor() ([]byte, []int) {
	return file_npool_inspire_gw_v1_reconcile_reconcile_proto_rawDescGZIP(), []int{0}
}

func (x *ReconcileRequest) GetAppID() string {
	if x != nil {
		return x.AppID
	}
	return ""
}

func (x *ReconcileRequest) GetAppGoodID() string {
	if x != nil {
		return x.AppGoodID
	}
	return ""
}

type ReconcileResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ReconcileResponse) Reset() {
	*x = ReconcileResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_inspire_gw_v1_reconcile_reconcile_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReconcileResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReconcileResponse) ProtoMessage() {}

func (x *ReconcileResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_inspire_gw_v1_reconcile_reconcile_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReconcileResponse.ProtoReflect.Descriptor instead.
func (*ReconcileResponse) Descriptor() ([]byte, []int) {
	return file_npool_inspire_gw_v1_reconcile_reconcile_proto_rawDescGZIP(), []int{1}
}

var File_npool_inspire_gw_v1_reconcile_reconcile_proto protoreflect.FileDescriptor

var file_npool_inspire_gw_v1_reconcile_reconcile_proto_rawDesc = []byte{
	0x0a, 0x2d, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x69, 0x6e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2f,
	0x67, 0x77, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x63, 0x6f, 0x6e, 0x63, 0x69, 0x6c, 0x65, 0x2f,
	0x72, 0x65, 0x63, 0x6f, 0x6e, 0x63, 0x69, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x21, 0x69, 0x6e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79,
	0x2e, 0x72, 0x65, 0x63, 0x6f, 0x6e, 0x63, 0x69, 0x6c, 0x69, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x46, 0x0a, 0x10, 0x52, 0x65, 0x63, 0x6f, 0x6e, 0x63, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x41, 0x70, 0x70, 0x49, 0x44, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x41, 0x70, 0x70, 0x49, 0x44, 0x12, 0x1c, 0x0a, 0x09, 0x41, 0x70,
	0x70, 0x47, 0x6f, 0x6f, 0x64, 0x49, 0x44, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x41,
	0x70, 0x70, 0x47, 0x6f, 0x6f, 0x64, 0x49, 0x44, 0x22, 0x13, 0x0a, 0x11, 0x52, 0x65, 0x63, 0x6f,
	0x6e, 0x63, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0x9c, 0x01,
	0x0a, 0x07, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x12, 0x90, 0x01, 0x0a, 0x09, 0x52, 0x65,
	0x63, 0x6f, 0x6e, 0x63, 0x69, 0x6c, 0x65, 0x12, 0x33, 0x2e, 0x69, 0x6e, 0x73, 0x70, 0x69, 0x72,
	0x65, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x72, 0x65, 0x63, 0x6f, 0x6e, 0x63,
	0x69, 0x6c, 0x69, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x63, 0x6f,
	0x6e, 0x63, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x34, 0x2e, 0x69,
	0x6e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x72,
	0x65, 0x63, 0x6f, 0x6e, 0x63, 0x69, 0x6c, 0x69, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31,
	0x2e, 0x52, 0x65, 0x63, 0x6f, 0x6e, 0x63, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x18, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x12, 0x3a, 0x01, 0x2a, 0x22, 0x0d, 0x2f,
	0x76, 0x31, 0x2f, 0x72, 0x65, 0x63, 0x6f, 0x6e, 0x63, 0x69, 0x6c, 0x65, 0x42, 0x45, 0x5a, 0x43,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4e, 0x70, 0x6f, 0x6f, 0x6c,
	0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x2f, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x69, 0x6e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2f, 0x67,
	0x77, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x63, 0x6f, 0x6e, 0x63, 0x69, 0x6c, 0x69, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_npool_inspire_gw_v1_reconcile_reconcile_proto_rawDescOnce sync.Once
	file_npool_inspire_gw_v1_reconcile_reconcile_proto_rawDescData = file_npool_inspire_gw_v1_reconcile_reconcile_proto_rawDesc
)

func file_npool_inspire_gw_v1_reconcile_reconcile_proto_rawDescGZIP() []byte {
	file_npool_inspire_gw_v1_reconcile_reconcile_proto_rawDescOnce.Do(func() {
		file_npool_inspire_gw_v1_reconcile_reconcile_proto_rawDescData = protoimpl.X.CompressGZIP(file_npool_inspire_gw_v1_reconcile_reconcile_proto_rawDescData)
	})
	return file_npool_inspire_gw_v1_reconcile_reconcile_proto_rawDescData
}

var file_npool_inspire_gw_v1_reconcile_reconcile_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_npool_inspire_gw_v1_reconcile_reconcile_proto_goTypes = []interface{}{
	(*ReconcileRequest)(nil),  // 0: inspire.gateway.reconciliation.v1.ReconcileRequest
	(*ReconcileResponse)(nil), // 1: inspire.gateway.reconciliation.v1.ReconcileResponse
}
var file_npool_inspire_gw_v1_reconcile_reconcile_proto_depIdxs = []int32{
	0, // 0: inspire.gateway.reconciliation.v1.Gateway.Reconcile:input_type -> inspire.gateway.reconciliation.v1.ReconcileRequest
	1, // 1: inspire.gateway.reconciliation.v1.Gateway.Reconcile:output_type -> inspire.gateway.reconciliation.v1.ReconcileResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_npool_inspire_gw_v1_reconcile_reconcile_proto_init() }
func file_npool_inspire_gw_v1_reconcile_reconcile_proto_init() {
	if File_npool_inspire_gw_v1_reconcile_reconcile_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_npool_inspire_gw_v1_reconcile_reconcile_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReconcileRequest); i {
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
		file_npool_inspire_gw_v1_reconcile_reconcile_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReconcileResponse); i {
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
			RawDescriptor: file_npool_inspire_gw_v1_reconcile_reconcile_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_npool_inspire_gw_v1_reconcile_reconcile_proto_goTypes,
		DependencyIndexes: file_npool_inspire_gw_v1_reconcile_reconcile_proto_depIdxs,
		MessageInfos:      file_npool_inspire_gw_v1_reconcile_reconcile_proto_msgTypes,
	}.Build()
	File_npool_inspire_gw_v1_reconcile_reconcile_proto = out.File
	file_npool_inspire_gw_v1_reconcile_reconcile_proto_rawDesc = nil
	file_npool_inspire_gw_v1_reconcile_reconcile_proto_goTypes = nil
	file_npool_inspire_gw_v1_reconcile_reconcile_proto_depIdxs = nil
}
