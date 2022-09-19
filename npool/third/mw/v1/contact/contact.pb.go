// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.18.1
// source: npool/third/mw/v1/contact/contact.proto

package contact

import (
	usedfor "github.com/NpoolPlatform/message/npool/third/mgr/v1/usedfor"
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

type ContactViaEmailRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AppID      string          `protobuf:"bytes,10,opt,name=AppID,proto3" json:"AppID,omitempty"`
	UsedFor    usedfor.UsedFor `protobuf:"varint,30,opt,name=UsedFor,proto3,enum=third.manager.usedfor.v1.UsedFor" json:"UsedFor,omitempty"`
	Sender     string          `protobuf:"bytes,40,opt,name=Sender,proto3" json:"Sender,omitempty"`
	Subject    string          `protobuf:"bytes,60,opt,name=Subject,proto3" json:"Subject,omitempty"`
	Body       string          `protobuf:"bytes,70,opt,name=Body,proto3" json:"Body,omitempty"`
	SenderName string          `protobuf:"bytes,80,opt,name=SenderName,proto3" json:"SenderName,omitempty"`
}

func (x *ContactViaEmailRequest) Reset() {
	*x = ContactViaEmailRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_third_mw_v1_contact_contact_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContactViaEmailRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContactViaEmailRequest) ProtoMessage() {}

func (x *ContactViaEmailRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_third_mw_v1_contact_contact_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContactViaEmailRequest.ProtoReflect.Descriptor instead.
func (*ContactViaEmailRequest) Descriptor() ([]byte, []int) {
	return file_npool_third_mw_v1_contact_contact_proto_rawDescGZIP(), []int{0}
}

func (x *ContactViaEmailRequest) GetAppID() string {
	if x != nil {
		return x.AppID
	}
	return ""
}

func (x *ContactViaEmailRequest) GetUsedFor() usedfor.UsedFor {
	if x != nil {
		return x.UsedFor
	}
	return usedfor.UsedFor(0)
}

func (x *ContactViaEmailRequest) GetSender() string {
	if x != nil {
		return x.Sender
	}
	return ""
}

func (x *ContactViaEmailRequest) GetSubject() string {
	if x != nil {
		return x.Subject
	}
	return ""
}

func (x *ContactViaEmailRequest) GetBody() string {
	if x != nil {
		return x.Body
	}
	return ""
}

func (x *ContactViaEmailRequest) GetSenderName() string {
	if x != nil {
		return x.SenderName
	}
	return ""
}

type ContactViaEmailResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ContactViaEmailResponse) Reset() {
	*x = ContactViaEmailResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_third_mw_v1_contact_contact_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContactViaEmailResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContactViaEmailResponse) ProtoMessage() {}

func (x *ContactViaEmailResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_third_mw_v1_contact_contact_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContactViaEmailResponse.ProtoReflect.Descriptor instead.
func (*ContactViaEmailResponse) Descriptor() ([]byte, []int) {
	return file_npool_third_mw_v1_contact_contact_proto_rawDescGZIP(), []int{1}
}

var File_npool_third_mw_v1_contact_contact_proto protoreflect.FileDescriptor

var file_npool_third_mw_v1_contact_contact_proto_rawDesc = []byte{
	0x0a, 0x27, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x74, 0x68, 0x69, 0x72, 0x64, 0x2f, 0x6d, 0x77,
	0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x2f, 0x63, 0x6f, 0x6e, 0x74,
	0x61, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1b, 0x74, 0x68, 0x69, 0x72, 0x64,
	0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x74,
	0x61, 0x63, 0x74, 0x2e, 0x76, 0x31, 0x1a, 0x29, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x74, 0x68,
	0x69, 0x72, 0x64, 0x2f, 0x6d, 0x67, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x64, 0x66,
	0x6f, 0x72, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x64, 0x66, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xd1, 0x01, 0x0a, 0x16, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x56, 0x69, 0x61,
	0x45, 0x6d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05,
	0x41, 0x70, 0x70, 0x49, 0x44, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x41, 0x70, 0x70,
	0x49, 0x44, 0x12, 0x3b, 0x0a, 0x07, 0x55, 0x73, 0x65, 0x64, 0x46, 0x6f, 0x72, 0x18, 0x1e, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x21, 0x2e, 0x74, 0x68, 0x69, 0x72, 0x64, 0x2e, 0x6d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x72, 0x2e, 0x75, 0x73, 0x65, 0x64, 0x66, 0x6f, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x55,
	0x73, 0x65, 0x64, 0x46, 0x6f, 0x72, 0x52, 0x07, 0x55, 0x73, 0x65, 0x64, 0x46, 0x6f, 0x72, 0x12,
	0x16, 0x0a, 0x06, 0x53, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x28, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x53, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x53, 0x75, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x18, 0x3c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x53, 0x75, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x42, 0x6f, 0x64, 0x79, 0x18, 0x46, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x42, 0x6f, 0x64, 0x79, 0x12, 0x1e, 0x0a, 0x0a, 0x53, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x4e,
	0x61, 0x6d, 0x65, 0x18, 0x50, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x53, 0x65, 0x6e, 0x64, 0x65,
	0x72, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x19, 0x0a, 0x17, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74,
	0x56, 0x69, 0x61, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x32, 0x8c, 0x01, 0x0a, 0x0a, 0x4d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x12,
	0x7e, 0x0a, 0x0f, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x56, 0x69, 0x61, 0x45, 0x6d, 0x61,
	0x69, 0x6c, 0x12, 0x33, 0x2e, 0x74, 0x68, 0x69, 0x72, 0x64, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c,
	0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x2e, 0x76, 0x31,
	0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x56, 0x69, 0x61, 0x45, 0x6d, 0x61, 0x69, 0x6c,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x34, 0x2e, 0x74, 0x68, 0x69, 0x72, 0x64, 0x2e,
	0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x61,
	0x63, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x56, 0x69, 0x61,
	0x45, 0x6d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42,
	0x3c, 0x5a, 0x3a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4e, 0x70,
	0x6f, 0x6f, 0x6c, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x2f, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x74, 0x68, 0x69, 0x72, 0x64, 0x2f,
	0x6d, 0x77, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_npool_third_mw_v1_contact_contact_proto_rawDescOnce sync.Once
	file_npool_third_mw_v1_contact_contact_proto_rawDescData = file_npool_third_mw_v1_contact_contact_proto_rawDesc
)

func file_npool_third_mw_v1_contact_contact_proto_rawDescGZIP() []byte {
	file_npool_third_mw_v1_contact_contact_proto_rawDescOnce.Do(func() {
		file_npool_third_mw_v1_contact_contact_proto_rawDescData = protoimpl.X.CompressGZIP(file_npool_third_mw_v1_contact_contact_proto_rawDescData)
	})
	return file_npool_third_mw_v1_contact_contact_proto_rawDescData
}

var file_npool_third_mw_v1_contact_contact_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_npool_third_mw_v1_contact_contact_proto_goTypes = []interface{}{
	(*ContactViaEmailRequest)(nil),  // 0: third.middleware.contact.v1.ContactViaEmailRequest
	(*ContactViaEmailResponse)(nil), // 1: third.middleware.contact.v1.ContactViaEmailResponse
	(usedfor.UsedFor)(0),            // 2: third.manager.usedfor.v1.UsedFor
}
var file_npool_third_mw_v1_contact_contact_proto_depIdxs = []int32{
	2, // 0: third.middleware.contact.v1.ContactViaEmailRequest.UsedFor:type_name -> third.manager.usedfor.v1.UsedFor
	0, // 1: third.middleware.contact.v1.Middleware.ContactViaEmail:input_type -> third.middleware.contact.v1.ContactViaEmailRequest
	1, // 2: third.middleware.contact.v1.Middleware.ContactViaEmail:output_type -> third.middleware.contact.v1.ContactViaEmailResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_npool_third_mw_v1_contact_contact_proto_init() }
func file_npool_third_mw_v1_contact_contact_proto_init() {
	if File_npool_third_mw_v1_contact_contact_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_npool_third_mw_v1_contact_contact_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ContactViaEmailRequest); i {
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
		file_npool_third_mw_v1_contact_contact_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ContactViaEmailResponse); i {
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
			RawDescriptor: file_npool_third_mw_v1_contact_contact_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_npool_third_mw_v1_contact_contact_proto_goTypes,
		DependencyIndexes: file_npool_third_mw_v1_contact_contact_proto_depIdxs,
		MessageInfos:      file_npool_third_mw_v1_contact_contact_proto_msgTypes,
	}.Build()
	File_npool_third_mw_v1_contact_contact_proto = out.File
	file_npool_third_mw_v1_contact_contact_proto_rawDesc = nil
	file_npool_third_mw_v1_contact_contact_proto_goTypes = nil
	file_npool_third_mw_v1_contact_contact_proto_depIdxs = nil
}
