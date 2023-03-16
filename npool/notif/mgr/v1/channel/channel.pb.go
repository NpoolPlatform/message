// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.18.1
// source: npool/notif/mgr/v1/channel/channel.proto

package channel

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

type NotifChannel int32

const (
	NotifChannel_DefaultChannel  NotifChannel = 0
	NotifChannel_ChannelEmail    NotifChannel = 10
	NotifChannel_ChannelSMS      NotifChannel = 20
	NotifChannel_ChannelFrontend NotifChannel = 30
)

// Enum value maps for NotifChannel.
var (
	NotifChannel_name = map[int32]string{
		0:  "DefaultChannel",
		10: "ChannelEmail",
		20: "ChannelSMS",
		30: "ChannelFrontend",
	}
	NotifChannel_value = map[string]int32{
		"DefaultChannel":  0,
		"ChannelEmail":    10,
		"ChannelSMS":      20,
		"ChannelFrontend": 30,
	}
)

func (x NotifChannel) Enum() *NotifChannel {
	p := new(NotifChannel)
	*p = x
	return p
}

func (x NotifChannel) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (NotifChannel) Descriptor() protoreflect.EnumDescriptor {
	return file_npool_notif_mgr_v1_channel_channel_proto_enumTypes[0].Descriptor()
}

func (NotifChannel) Type() protoreflect.EnumType {
	return &file_npool_notif_mgr_v1_channel_channel_proto_enumTypes[0]
}

func (x NotifChannel) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use NotifChannel.Descriptor instead.
func (NotifChannel) EnumDescriptor() ([]byte, []int) {
	return file_npool_notif_mgr_v1_channel_channel_proto_rawDescGZIP(), []int{0}
}

var File_npool_notif_mgr_v1_channel_channel_proto protoreflect.FileDescriptor

var file_npool_notif_mgr_v1_channel_channel_proto_rawDesc = []byte{
	0x0a, 0x28, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x2f, 0x6d, 0x67,
	0x72, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x2f, 0x63, 0x68, 0x61,
	0x6e, 0x6e, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18, 0x6e, 0x6f, 0x74, 0x69,
	0x66, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65,
	0x6c, 0x2e, 0x76, 0x31, 0x2a, 0x59, 0x0a, 0x0c, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x43, 0x68, 0x61,
	0x6e, 0x6e, 0x65, 0x6c, 0x12, 0x12, 0x0a, 0x0e, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x43,
	0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x10, 0x00, 0x12, 0x10, 0x0a, 0x0c, 0x43, 0x68, 0x61, 0x6e,
	0x6e, 0x65, 0x6c, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x10, 0x0a, 0x12, 0x0e, 0x0a, 0x0a, 0x43, 0x68,
	0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x53, 0x4d, 0x53, 0x10, 0x14, 0x12, 0x13, 0x0a, 0x0f, 0x43, 0x68,
	0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x46, 0x72, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x64, 0x10, 0x1e, 0x42,
	0x3d, 0x5a, 0x3b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4e, 0x70,
	0x6f, 0x6f, 0x6c, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x2f, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x2f,
	0x6d, 0x67, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_npool_notif_mgr_v1_channel_channel_proto_rawDescOnce sync.Once
	file_npool_notif_mgr_v1_channel_channel_proto_rawDescData = file_npool_notif_mgr_v1_channel_channel_proto_rawDesc
)

func file_npool_notif_mgr_v1_channel_channel_proto_rawDescGZIP() []byte {
	file_npool_notif_mgr_v1_channel_channel_proto_rawDescOnce.Do(func() {
		file_npool_notif_mgr_v1_channel_channel_proto_rawDescData = protoimpl.X.CompressGZIP(file_npool_notif_mgr_v1_channel_channel_proto_rawDescData)
	})
	return file_npool_notif_mgr_v1_channel_channel_proto_rawDescData
}

var file_npool_notif_mgr_v1_channel_channel_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_npool_notif_mgr_v1_channel_channel_proto_goTypes = []interface{}{
	(NotifChannel)(0), // 0: notif.manager.channel.v1.NotifChannel
}
var file_npool_notif_mgr_v1_channel_channel_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_npool_notif_mgr_v1_channel_channel_proto_init() }
func file_npool_notif_mgr_v1_channel_channel_proto_init() {
	if File_npool_notif_mgr_v1_channel_channel_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_npool_notif_mgr_v1_channel_channel_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_npool_notif_mgr_v1_channel_channel_proto_goTypes,
		DependencyIndexes: file_npool_notif_mgr_v1_channel_channel_proto_depIdxs,
		EnumInfos:         file_npool_notif_mgr_v1_channel_channel_proto_enumTypes,
	}.Build()
	File_npool_notif_mgr_v1_channel_channel_proto = out.File
	file_npool_notif_mgr_v1_channel_channel_proto_rawDesc = nil
	file_npool_notif_mgr_v1_channel_channel_proto_goTypes = nil
	file_npool_notif_mgr_v1_channel_channel_proto_depIdxs = nil
}
