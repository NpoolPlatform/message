// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: npool/basal/mw/v1/event/event.proto

package event

import (
	api "github.com/NpoolPlatform/message/npool/basal/mw/v1/api"
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

type APIRegisterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Protocol    api.Protocol `protobuf:"varint,10,opt,name=Protocol,proto3,enum=basal.middleware.api.v1.Protocol" json:"Protocol,omitempty"`
	ServiceName string       `protobuf:"bytes,20,opt,name=ServiceName,proto3" json:"ServiceName,omitempty"`
	Method      api.Method   `protobuf:"varint,30,opt,name=Method,proto3,enum=basal.middleware.api.v1.Method" json:"Method,omitempty"`
	MethodName  string       `protobuf:"bytes,40,opt,name=MethodName,proto3" json:"MethodName,omitempty"`
	Path        string       `protobuf:"bytes,50,opt,name=Path,proto3" json:"Path,omitempty"`
	PathPrefix  string       `protobuf:"bytes,60,opt,name=PathPrefix,proto3" json:"PathPrefix,omitempty"`
	Domains     []string     `protobuf:"bytes,70,rep,name=Domains,proto3" json:"Domains,omitempty"`
	Exported    *bool        `protobuf:"varint,80,opt,name=Exported,proto3,oneof" json:"Exported,omitempty"`
}

func (x *APIRegisterRequest) Reset() {
	*x = APIRegisterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_basal_mw_v1_event_event_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *APIRegisterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*APIRegisterRequest) ProtoMessage() {}

func (x *APIRegisterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_basal_mw_v1_event_event_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use APIRegisterRequest.ProtoReflect.Descriptor instead.
func (*APIRegisterRequest) Descriptor() ([]byte, []int) {
	return file_npool_basal_mw_v1_event_event_proto_rawDescGZIP(), []int{0}
}

func (x *APIRegisterRequest) GetProtocol() api.Protocol {
	if x != nil {
		return x.Protocol
	}
	return api.Protocol(0)
}

func (x *APIRegisterRequest) GetServiceName() string {
	if x != nil {
		return x.ServiceName
	}
	return ""
}

func (x *APIRegisterRequest) GetMethod() api.Method {
	if x != nil {
		return x.Method
	}
	return api.Method(0)
}

func (x *APIRegisterRequest) GetMethodName() string {
	if x != nil {
		return x.MethodName
	}
	return ""
}

func (x *APIRegisterRequest) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *APIRegisterRequest) GetPathPrefix() string {
	if x != nil {
		return x.PathPrefix
	}
	return ""
}

func (x *APIRegisterRequest) GetDomains() []string {
	if x != nil {
		return x.Domains
	}
	return nil
}

func (x *APIRegisterRequest) GetExported() bool {
	if x != nil && x.Exported != nil {
		return *x.Exported
	}
	return false
}

type APIRegisterResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Infos []*api.API `protobuf:"bytes,10,rep,name=Infos,proto3" json:"Infos,omitempty"`
}

func (x *APIRegisterResponse) Reset() {
	*x = APIRegisterResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_basal_mw_v1_event_event_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *APIRegisterResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*APIRegisterResponse) ProtoMessage() {}

func (x *APIRegisterResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_basal_mw_v1_event_event_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use APIRegisterResponse.ProtoReflect.Descriptor instead.
func (*APIRegisterResponse) Descriptor() ([]byte, []int) {
	return file_npool_basal_mw_v1_event_event_proto_rawDescGZIP(), []int{1}
}

func (x *APIRegisterResponse) GetInfos() []*api.API {
	if x != nil {
		return x.Infos
	}
	return nil
}

var File_npool_basal_mw_v1_event_event_proto protoreflect.FileDescriptor

var file_npool_basal_mw_v1_event_event_proto_rawDesc = []byte{
	0x0a, 0x23, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x62, 0x61, 0x73, 0x61, 0x6c, 0x2f, 0x6d, 0x77,
	0x2f, 0x76, 0x31, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x19, 0x62, 0x61, 0x73, 0x61, 0x6c, 0x2e, 0x6d, 0x69, 0x64,
	0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31,
	0x1a, 0x1f, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x62, 0x61, 0x73, 0x61, 0x6c, 0x2f, 0x6d, 0x77,
	0x2f, 0x76, 0x31, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xca, 0x02, 0x0a, 0x12, 0x41, 0x50, 0x49, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3d, 0x0a, 0x08, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x6f, 0x6c, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x21, 0x2e, 0x62, 0x61, 0x73,
	0x61, 0x6c, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x52, 0x08, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x12, 0x20, 0x0a, 0x0b, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x37, 0x0a, 0x06, 0x4d, 0x65, 0x74,
	0x68, 0x6f, 0x64, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1f, 0x2e, 0x62, 0x61, 0x73, 0x61,
	0x6c, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x52, 0x06, 0x4d, 0x65, 0x74, 0x68,
	0x6f, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x4e, 0x61, 0x6d, 0x65,
	0x18, 0x28, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x50, 0x61, 0x74, 0x68, 0x18, 0x32, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x50, 0x61, 0x74, 0x68, 0x12, 0x1e, 0x0a, 0x0a, 0x50, 0x61, 0x74, 0x68, 0x50, 0x72,
	0x65, 0x66, 0x69, 0x78, 0x18, 0x3c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x50, 0x61, 0x74, 0x68,
	0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x12, 0x18, 0x0a, 0x07, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e,
	0x73, 0x18, 0x46, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x73,
	0x12, 0x1f, 0x0a, 0x08, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x64, 0x18, 0x50, 0x20, 0x01,
	0x28, 0x08, 0x48, 0x00, 0x52, 0x08, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x64, 0x88, 0x01,
	0x01, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x64, 0x22, 0x49,
	0x0a, 0x13, 0x41, 0x50, 0x49, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x32, 0x0a, 0x05, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x18, 0x0a,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x62, 0x61, 0x73, 0x61, 0x6c, 0x2e, 0x6d, 0x69, 0x64,
	0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x41,
	0x50, 0x49, 0x52, 0x05, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x42, 0x3a, 0x5a, 0x38, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4e, 0x70, 0x6f, 0x6f, 0x6c, 0x50, 0x6c, 0x61,
	0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x6e, 0x70,
	0x6f, 0x6f, 0x6c, 0x2f, 0x62, 0x61, 0x73, 0x61, 0x6c, 0x2f, 0x6d, 0x77, 0x2f, 0x76, 0x31, 0x2f,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_npool_basal_mw_v1_event_event_proto_rawDescOnce sync.Once
	file_npool_basal_mw_v1_event_event_proto_rawDescData = file_npool_basal_mw_v1_event_event_proto_rawDesc
)

func file_npool_basal_mw_v1_event_event_proto_rawDescGZIP() []byte {
	file_npool_basal_mw_v1_event_event_proto_rawDescOnce.Do(func() {
		file_npool_basal_mw_v1_event_event_proto_rawDescData = protoimpl.X.CompressGZIP(file_npool_basal_mw_v1_event_event_proto_rawDescData)
	})
	return file_npool_basal_mw_v1_event_event_proto_rawDescData
}

var file_npool_basal_mw_v1_event_event_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_npool_basal_mw_v1_event_event_proto_goTypes = []interface{}{
	(*APIRegisterRequest)(nil),  // 0: basal.middleware.event.v1.APIRegisterRequest
	(*APIRegisterResponse)(nil), // 1: basal.middleware.event.v1.APIRegisterResponse
	(api.Protocol)(0),           // 2: basal.middleware.api.v1.Protocol
	(api.Method)(0),             // 3: basal.middleware.api.v1.Method
	(*api.API)(nil),             // 4: basal.middleware.api.v1.API
}
var file_npool_basal_mw_v1_event_event_proto_depIdxs = []int32{
	2, // 0: basal.middleware.event.v1.APIRegisterRequest.Protocol:type_name -> basal.middleware.api.v1.Protocol
	3, // 1: basal.middleware.event.v1.APIRegisterRequest.Method:type_name -> basal.middleware.api.v1.Method
	4, // 2: basal.middleware.event.v1.APIRegisterResponse.Infos:type_name -> basal.middleware.api.v1.API
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_npool_basal_mw_v1_event_event_proto_init() }
func file_npool_basal_mw_v1_event_event_proto_init() {
	if File_npool_basal_mw_v1_event_event_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_npool_basal_mw_v1_event_event_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*APIRegisterRequest); i {
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
		file_npool_basal_mw_v1_event_event_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*APIRegisterResponse); i {
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
	file_npool_basal_mw_v1_event_event_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_npool_basal_mw_v1_event_event_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_npool_basal_mw_v1_event_event_proto_goTypes,
		DependencyIndexes: file_npool_basal_mw_v1_event_event_proto_depIdxs,
		MessageInfos:      file_npool_basal_mw_v1_event_event_proto_msgTypes,
	}.Build()
	File_npool_basal_mw_v1_event_event_proto = out.File
	file_npool_basal_mw_v1_event_event_proto_rawDesc = nil
	file_npool_basal_mw_v1_event_event_proto_goTypes = nil
	file_npool_basal_mw_v1_event_event_proto_depIdxs = nil
}
