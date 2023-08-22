// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.18.1
// source: npool/good/gw/v1/deviceinfo/deviceinfo.proto

package deviceinfo

import (
	deviceinfo "github.com/NpoolPlatform/message/npool/good/mw/v1/deviceinfo"
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

type CreateDeviceInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type            string   `protobuf:"bytes,10,opt,name=Type,proto3" json:"Type,omitempty"`
	Manufacturer    string   `protobuf:"bytes,20,opt,name=Manufacturer,proto3" json:"Manufacturer,omitempty"`
	PowerComsuption uint32   `protobuf:"varint,30,opt,name=PowerComsuption,proto3" json:"PowerComsuption,omitempty"`
	ShipmentAt      uint32   `protobuf:"varint,40,opt,name=ShipmentAt,proto3" json:"ShipmentAt,omitempty"`
	Posters         []string `protobuf:"bytes,50,rep,name=Posters,proto3" json:"Posters,omitempty"`
}

func (x *CreateDeviceInfoRequest) Reset() {
	*x = CreateDeviceInfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_good_gw_v1_deviceinfo_deviceinfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateDeviceInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateDeviceInfoRequest) ProtoMessage() {}

func (x *CreateDeviceInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_good_gw_v1_deviceinfo_deviceinfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateDeviceInfoRequest.ProtoReflect.Descriptor instead.
func (*CreateDeviceInfoRequest) Descriptor() ([]byte, []int) {
	return file_npool_good_gw_v1_deviceinfo_deviceinfo_proto_rawDescGZIP(), []int{0}
}

func (x *CreateDeviceInfoRequest) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *CreateDeviceInfoRequest) GetManufacturer() string {
	if x != nil {
		return x.Manufacturer
	}
	return ""
}

func (x *CreateDeviceInfoRequest) GetPowerComsuption() uint32 {
	if x != nil {
		return x.PowerComsuption
	}
	return 0
}

func (x *CreateDeviceInfoRequest) GetShipmentAt() uint32 {
	if x != nil {
		return x.ShipmentAt
	}
	return 0
}

func (x *CreateDeviceInfoRequest) GetPosters() []string {
	if x != nil {
		return x.Posters
	}
	return nil
}

type CreateDeviceInfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *deviceinfo.DeviceInfo `protobuf:"bytes,10,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (x *CreateDeviceInfoResponse) Reset() {
	*x = CreateDeviceInfoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_good_gw_v1_deviceinfo_deviceinfo_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateDeviceInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateDeviceInfoResponse) ProtoMessage() {}

func (x *CreateDeviceInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_good_gw_v1_deviceinfo_deviceinfo_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateDeviceInfoResponse.ProtoReflect.Descriptor instead.
func (*CreateDeviceInfoResponse) Descriptor() ([]byte, []int) {
	return file_npool_good_gw_v1_deviceinfo_deviceinfo_proto_rawDescGZIP(), []int{1}
}

func (x *CreateDeviceInfoResponse) GetInfo() *deviceinfo.DeviceInfo {
	if x != nil {
		return x.Info
	}
	return nil
}

type UpdateDeviceInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID              string   `protobuf:"bytes,10,opt,name=ID,proto3" json:"ID,omitempty"`
	Type            *string  `protobuf:"bytes,20,opt,name=Type,proto3,oneof" json:"Type,omitempty"`
	Manufacturer    *string  `protobuf:"bytes,30,opt,name=Manufacturer,proto3,oneof" json:"Manufacturer,omitempty"`
	PowerComsuption *uint32  `protobuf:"varint,40,opt,name=PowerComsuption,proto3,oneof" json:"PowerComsuption,omitempty"`
	ShipmentAt      *uint32  `protobuf:"varint,50,opt,name=ShipmentAt,proto3,oneof" json:"ShipmentAt,omitempty"`
	Posters         []string `protobuf:"bytes,60,rep,name=Posters,proto3" json:"Posters,omitempty"`
}

func (x *UpdateDeviceInfoRequest) Reset() {
	*x = UpdateDeviceInfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_good_gw_v1_deviceinfo_deviceinfo_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateDeviceInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateDeviceInfoRequest) ProtoMessage() {}

func (x *UpdateDeviceInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_good_gw_v1_deviceinfo_deviceinfo_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateDeviceInfoRequest.ProtoReflect.Descriptor instead.
func (*UpdateDeviceInfoRequest) Descriptor() ([]byte, []int) {
	return file_npool_good_gw_v1_deviceinfo_deviceinfo_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateDeviceInfoRequest) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *UpdateDeviceInfoRequest) GetType() string {
	if x != nil && x.Type != nil {
		return *x.Type
	}
	return ""
}

func (x *UpdateDeviceInfoRequest) GetManufacturer() string {
	if x != nil && x.Manufacturer != nil {
		return *x.Manufacturer
	}
	return ""
}

func (x *UpdateDeviceInfoRequest) GetPowerComsuption() uint32 {
	if x != nil && x.PowerComsuption != nil {
		return *x.PowerComsuption
	}
	return 0
}

func (x *UpdateDeviceInfoRequest) GetShipmentAt() uint32 {
	if x != nil && x.ShipmentAt != nil {
		return *x.ShipmentAt
	}
	return 0
}

func (x *UpdateDeviceInfoRequest) GetPosters() []string {
	if x != nil {
		return x.Posters
	}
	return nil
}

type UpdateDeviceInfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *deviceinfo.DeviceInfo `protobuf:"bytes,10,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (x *UpdateDeviceInfoResponse) Reset() {
	*x = UpdateDeviceInfoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_good_gw_v1_deviceinfo_deviceinfo_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateDeviceInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateDeviceInfoResponse) ProtoMessage() {}

func (x *UpdateDeviceInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_good_gw_v1_deviceinfo_deviceinfo_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateDeviceInfoResponse.ProtoReflect.Descriptor instead.
func (*UpdateDeviceInfoResponse) Descriptor() ([]byte, []int) {
	return file_npool_good_gw_v1_deviceinfo_deviceinfo_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateDeviceInfoResponse) GetInfo() *deviceinfo.DeviceInfo {
	if x != nil {
		return x.Info
	}
	return nil
}

type GetDeviceInfosRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Offset int32 `protobuf:"varint,10,opt,name=Offset,proto3" json:"Offset,omitempty"`
	Limit  int32 `protobuf:"varint,20,opt,name=Limit,proto3" json:"Limit,omitempty"`
}

func (x *GetDeviceInfosRequest) Reset() {
	*x = GetDeviceInfosRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_good_gw_v1_deviceinfo_deviceinfo_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDeviceInfosRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDeviceInfosRequest) ProtoMessage() {}

func (x *GetDeviceInfosRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_good_gw_v1_deviceinfo_deviceinfo_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDeviceInfosRequest.ProtoReflect.Descriptor instead.
func (*GetDeviceInfosRequest) Descriptor() ([]byte, []int) {
	return file_npool_good_gw_v1_deviceinfo_deviceinfo_proto_rawDescGZIP(), []int{4}
}

func (x *GetDeviceInfosRequest) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *GetDeviceInfosRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type GetDeviceInfosResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Infos []*deviceinfo.DeviceInfo `protobuf:"bytes,10,rep,name=Infos,proto3" json:"Infos,omitempty"`
	Total uint32                   `protobuf:"varint,20,opt,name=Total,proto3" json:"Total,omitempty"`
}

func (x *GetDeviceInfosResponse) Reset() {
	*x = GetDeviceInfosResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_good_gw_v1_deviceinfo_deviceinfo_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDeviceInfosResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDeviceInfosResponse) ProtoMessage() {}

func (x *GetDeviceInfosResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_good_gw_v1_deviceinfo_deviceinfo_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDeviceInfosResponse.ProtoReflect.Descriptor instead.
func (*GetDeviceInfosResponse) Descriptor() ([]byte, []int) {
	return file_npool_good_gw_v1_deviceinfo_deviceinfo_proto_rawDescGZIP(), []int{5}
}

func (x *GetDeviceInfosResponse) GetInfos() []*deviceinfo.DeviceInfo {
	if x != nil {
		return x.Infos
	}
	return nil
}

func (x *GetDeviceInfosResponse) GetTotal() uint32 {
	if x != nil {
		return x.Total
	}
	return 0
}

type DeleteDeviceInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID string `protobuf:"bytes,10,opt,name=ID,proto3" json:"ID,omitempty"`
}

func (x *DeleteDeviceInfoRequest) Reset() {
	*x = DeleteDeviceInfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_good_gw_v1_deviceinfo_deviceinfo_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteDeviceInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteDeviceInfoRequest) ProtoMessage() {}

func (x *DeleteDeviceInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_good_gw_v1_deviceinfo_deviceinfo_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteDeviceInfoRequest.ProtoReflect.Descriptor instead.
func (*DeleteDeviceInfoRequest) Descriptor() ([]byte, []int) {
	return file_npool_good_gw_v1_deviceinfo_deviceinfo_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteDeviceInfoRequest) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

type DeleteDeviceInfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *deviceinfo.DeviceInfo `protobuf:"bytes,10,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (x *DeleteDeviceInfoResponse) Reset() {
	*x = DeleteDeviceInfoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_good_gw_v1_deviceinfo_deviceinfo_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteDeviceInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteDeviceInfoResponse) ProtoMessage() {}

func (x *DeleteDeviceInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_good_gw_v1_deviceinfo_deviceinfo_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteDeviceInfoResponse.ProtoReflect.Descriptor instead.
func (*DeleteDeviceInfoResponse) Descriptor() ([]byte, []int) {
	return file_npool_good_gw_v1_deviceinfo_deviceinfo_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteDeviceInfoResponse) GetInfo() *deviceinfo.DeviceInfo {
	if x != nil {
		return x.Info
	}
	return nil
}

var File_npool_good_gw_v1_deviceinfo_deviceinfo_proto protoreflect.FileDescriptor

var file_npool_good_gw_v1_deviceinfo_deviceinfo_proto_rawDesc = []byte{
	0x0a, 0x2c, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x67, 0x6f, 0x6f, 0x64, 0x2f, 0x67, 0x77, 0x2f,
	0x76, 0x31, 0x2f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x69, 0x6e, 0x66, 0x6f, 0x2f, 0x64, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1a,
	0x67, 0x6f, 0x6f, 0x64, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x64, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x76, 0x31, 0x1a, 0x2c, 0x6e, 0x70, 0x6f, 0x6f,
	0x6c, 0x2f, 0x67, 0x6f, 0x6f, 0x64, 0x2f, 0x6d, 0x77, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x69, 0x6e, 0x66, 0x6f, 0x2f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x69, 0x6e,
	0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb5, 0x01, 0x0a, 0x17, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x4d, 0x61, 0x6e, 0x75, 0x66, 0x61,
	0x63, 0x74, 0x75, 0x72, 0x65, 0x72, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x4d, 0x61,
	0x6e, 0x75, 0x66, 0x61, 0x63, 0x74, 0x75, 0x72, 0x65, 0x72, 0x12, 0x28, 0x0a, 0x0f, 0x50, 0x6f,
	0x77, 0x65, 0x72, 0x43, 0x6f, 0x6d, 0x73, 0x75, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x1e, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x0f, 0x50, 0x6f, 0x77, 0x65, 0x72, 0x43, 0x6f, 0x6d, 0x73, 0x75, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1e, 0x0a, 0x0a, 0x53, 0x68, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74,
	0x41, 0x74, 0x18, 0x28, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x53, 0x68, 0x69, 0x70, 0x6d, 0x65,
	0x6e, 0x74, 0x41, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x50, 0x6f, 0x73, 0x74, 0x65, 0x72, 0x73, 0x18,
	0x32, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x50, 0x6f, 0x73, 0x74, 0x65, 0x72, 0x73, 0x22, 0x59,
	0x0a, 0x18, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3d, 0x0a, 0x04, 0x49, 0x6e,
	0x66, 0x6f, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x67, 0x6f, 0x6f, 0x64, 0x2e,
	0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x64, 0x65, 0x76, 0x69, 0x63,
	0x65, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x96, 0x02, 0x0a, 0x17, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x49, 0x44, 0x12, 0x17, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x18, 0x14, 0x20,
	0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x04, 0x54, 0x79, 0x70, 0x65, 0x88, 0x01, 0x01, 0x12, 0x27,
	0x0a, 0x0c, 0x4d, 0x61, 0x6e, 0x75, 0x66, 0x61, 0x63, 0x74, 0x75, 0x72, 0x65, 0x72, 0x18, 0x1e,
	0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x0c, 0x4d, 0x61, 0x6e, 0x75, 0x66, 0x61, 0x63, 0x74,
	0x75, 0x72, 0x65, 0x72, 0x88, 0x01, 0x01, 0x12, 0x2d, 0x0a, 0x0f, 0x50, 0x6f, 0x77, 0x65, 0x72,
	0x43, 0x6f, 0x6d, 0x73, 0x75, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x28, 0x20, 0x01, 0x28, 0x0d,
	0x48, 0x02, 0x52, 0x0f, 0x50, 0x6f, 0x77, 0x65, 0x72, 0x43, 0x6f, 0x6d, 0x73, 0x75, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x88, 0x01, 0x01, 0x12, 0x23, 0x0a, 0x0a, 0x53, 0x68, 0x69, 0x70, 0x6d, 0x65,
	0x6e, 0x74, 0x41, 0x74, 0x18, 0x32, 0x20, 0x01, 0x28, 0x0d, 0x48, 0x03, 0x52, 0x0a, 0x53, 0x68,
	0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x41, 0x74, 0x88, 0x01, 0x01, 0x12, 0x18, 0x0a, 0x07, 0x50,
	0x6f, 0x73, 0x74, 0x65, 0x72, 0x73, 0x18, 0x3c, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x50, 0x6f,
	0x73, 0x74, 0x65, 0x72, 0x73, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x54, 0x79, 0x70, 0x65, 0x42, 0x0f,
	0x0a, 0x0d, 0x5f, 0x4d, 0x61, 0x6e, 0x75, 0x66, 0x61, 0x63, 0x74, 0x75, 0x72, 0x65, 0x72, 0x42,
	0x12, 0x0a, 0x10, 0x5f, 0x50, 0x6f, 0x77, 0x65, 0x72, 0x43, 0x6f, 0x6d, 0x73, 0x75, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x53, 0x68, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74,
	0x41, 0x74, 0x22, 0x59, 0x0a, 0x18, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x44, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3d,
	0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x67,
	0x6f, 0x6f, 0x64, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x64,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x45, 0x0a,
	0x15, 0x47, 0x65, 0x74, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x14,
	0x0a, 0x05, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x14, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x4c,
	0x69, 0x6d, 0x69, 0x74, 0x22, 0x6f, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x44, 0x65, 0x76, 0x69, 0x63,
	0x65, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3f,
	0x0a, 0x05, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x29, 0x2e,
	0x67, 0x6f, 0x6f, 0x64, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e,
	0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x05, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x12,
	0x14, 0x0a, 0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x14, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05,
	0x54, 0x6f, 0x74, 0x61, 0x6c, 0x22, 0x29, 0x0a, 0x17, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x44,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44,
	0x22, 0x59, 0x0a, 0x18, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3d, 0x0a, 0x04,
	0x49, 0x6e, 0x66, 0x6f, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x67, 0x6f, 0x6f,
	0x64, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x64, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x63,
	0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x32, 0x89, 0x05, 0x0a, 0x07,
	0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x12, 0x9f, 0x01, 0x0a, 0x10, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x33, 0x2e, 0x67,
	0x6f, 0x6f, 0x64, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x64, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x34, 0x2e, 0x67, 0x6f, 0x6f, 0x64, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79,
	0x2e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x20, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1a, 0x3a,
	0x01, 0x2a, 0x22, 0x15, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x2f, 0x64,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x69, 0x6e, 0x66, 0x6f, 0x12, 0x9f, 0x01, 0x0a, 0x10, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x33,
	0x2e, 0x67, 0x6f, 0x6f, 0x64, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x64, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x34, 0x2e, 0x67, 0x6f, 0x6f, 0x64, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77,
	0x61, 0x79, 0x2e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x76, 0x31,
	0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x20, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x1a, 0x3a, 0x01, 0x2a, 0x22, 0x15, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x2f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x69, 0x6e, 0x66, 0x6f, 0x12, 0x97, 0x01, 0x0a, 0x0e,
	0x47, 0x65, 0x74, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x12, 0x31,
	0x2e, 0x67, 0x6f, 0x6f, 0x64, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x64, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x44,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x32, 0x2e, 0x67, 0x6f, 0x6f, 0x64, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79,
	0x2e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x47,
	0x65, 0x74, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x18, 0x3a, 0x01, 0x2a,
	0x22, 0x13, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x65, 0x74, 0x2f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x69, 0x6e, 0x66, 0x6f, 0x73, 0x12, 0x9f, 0x01, 0x0a, 0x10, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x33, 0x2e, 0x67, 0x6f, 0x6f,
	0x64, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x44, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x34, 0x2e, 0x67, 0x6f, 0x6f, 0x64, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x64,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x20, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1a, 0x3a, 0x01, 0x2a,
	0x22, 0x15, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x2f, 0x64, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x69, 0x6e, 0x66, 0x6f, 0x42, 0x3e, 0x5a, 0x3c, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4e, 0x70, 0x6f, 0x6f, 0x6c, 0x50, 0x6c, 0x61, 0x74, 0x66,
	0x6f, 0x72, 0x6d, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x6e, 0x70, 0x6f, 0x6f,
	0x6c, 0x2f, 0x67, 0x6f, 0x6f, 0x64, 0x2f, 0x67, 0x77, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x69, 0x6e, 0x66, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_npool_good_gw_v1_deviceinfo_deviceinfo_proto_rawDescOnce sync.Once
	file_npool_good_gw_v1_deviceinfo_deviceinfo_proto_rawDescData = file_npool_good_gw_v1_deviceinfo_deviceinfo_proto_rawDesc
)

func file_npool_good_gw_v1_deviceinfo_deviceinfo_proto_rawDescGZIP() []byte {
	file_npool_good_gw_v1_deviceinfo_deviceinfo_proto_rawDescOnce.Do(func() {
		file_npool_good_gw_v1_deviceinfo_deviceinfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_npool_good_gw_v1_deviceinfo_deviceinfo_proto_rawDescData)
	})
	return file_npool_good_gw_v1_deviceinfo_deviceinfo_proto_rawDescData
}

var file_npool_good_gw_v1_deviceinfo_deviceinfo_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_npool_good_gw_v1_deviceinfo_deviceinfo_proto_goTypes = []interface{}{
	(*CreateDeviceInfoRequest)(nil),  // 0: good.gateway.deviceinfo.v1.CreateDeviceInfoRequest
	(*CreateDeviceInfoResponse)(nil), // 1: good.gateway.deviceinfo.v1.CreateDeviceInfoResponse
	(*UpdateDeviceInfoRequest)(nil),  // 2: good.gateway.deviceinfo.v1.UpdateDeviceInfoRequest
	(*UpdateDeviceInfoResponse)(nil), // 3: good.gateway.deviceinfo.v1.UpdateDeviceInfoResponse
	(*GetDeviceInfosRequest)(nil),    // 4: good.gateway.deviceinfo.v1.GetDeviceInfosRequest
	(*GetDeviceInfosResponse)(nil),   // 5: good.gateway.deviceinfo.v1.GetDeviceInfosResponse
	(*DeleteDeviceInfoRequest)(nil),  // 6: good.gateway.deviceinfo.v1.DeleteDeviceInfoRequest
	(*DeleteDeviceInfoResponse)(nil), // 7: good.gateway.deviceinfo.v1.DeleteDeviceInfoResponse
	(*deviceinfo.DeviceInfo)(nil),    // 8: good.middleware.deviceinfo.v1.DeviceInfo
}
var file_npool_good_gw_v1_deviceinfo_deviceinfo_proto_depIdxs = []int32{
	8, // 0: good.gateway.deviceinfo.v1.CreateDeviceInfoResponse.Info:type_name -> good.middleware.deviceinfo.v1.DeviceInfo
	8, // 1: good.gateway.deviceinfo.v1.UpdateDeviceInfoResponse.Info:type_name -> good.middleware.deviceinfo.v1.DeviceInfo
	8, // 2: good.gateway.deviceinfo.v1.GetDeviceInfosResponse.Infos:type_name -> good.middleware.deviceinfo.v1.DeviceInfo
	8, // 3: good.gateway.deviceinfo.v1.DeleteDeviceInfoResponse.Info:type_name -> good.middleware.deviceinfo.v1.DeviceInfo
	0, // 4: good.gateway.deviceinfo.v1.Gateway.CreateDeviceInfo:input_type -> good.gateway.deviceinfo.v1.CreateDeviceInfoRequest
	2, // 5: good.gateway.deviceinfo.v1.Gateway.UpdateDeviceInfo:input_type -> good.gateway.deviceinfo.v1.UpdateDeviceInfoRequest
	4, // 6: good.gateway.deviceinfo.v1.Gateway.GetDeviceInfos:input_type -> good.gateway.deviceinfo.v1.GetDeviceInfosRequest
	6, // 7: good.gateway.deviceinfo.v1.Gateway.DeleteDeviceInfo:input_type -> good.gateway.deviceinfo.v1.DeleteDeviceInfoRequest
	1, // 8: good.gateway.deviceinfo.v1.Gateway.CreateDeviceInfo:output_type -> good.gateway.deviceinfo.v1.CreateDeviceInfoResponse
	3, // 9: good.gateway.deviceinfo.v1.Gateway.UpdateDeviceInfo:output_type -> good.gateway.deviceinfo.v1.UpdateDeviceInfoResponse
	5, // 10: good.gateway.deviceinfo.v1.Gateway.GetDeviceInfos:output_type -> good.gateway.deviceinfo.v1.GetDeviceInfosResponse
	7, // 11: good.gateway.deviceinfo.v1.Gateway.DeleteDeviceInfo:output_type -> good.gateway.deviceinfo.v1.DeleteDeviceInfoResponse
	8, // [8:12] is the sub-list for method output_type
	4, // [4:8] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_npool_good_gw_v1_deviceinfo_deviceinfo_proto_init() }
func file_npool_good_gw_v1_deviceinfo_deviceinfo_proto_init() {
	if File_npool_good_gw_v1_deviceinfo_deviceinfo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_npool_good_gw_v1_deviceinfo_deviceinfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateDeviceInfoRequest); i {
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
		file_npool_good_gw_v1_deviceinfo_deviceinfo_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateDeviceInfoResponse); i {
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
		file_npool_good_gw_v1_deviceinfo_deviceinfo_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateDeviceInfoRequest); i {
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
		file_npool_good_gw_v1_deviceinfo_deviceinfo_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateDeviceInfoResponse); i {
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
		file_npool_good_gw_v1_deviceinfo_deviceinfo_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDeviceInfosRequest); i {
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
		file_npool_good_gw_v1_deviceinfo_deviceinfo_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDeviceInfosResponse); i {
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
		file_npool_good_gw_v1_deviceinfo_deviceinfo_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteDeviceInfoRequest); i {
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
		file_npool_good_gw_v1_deviceinfo_deviceinfo_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteDeviceInfoResponse); i {
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
	file_npool_good_gw_v1_deviceinfo_deviceinfo_proto_msgTypes[2].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_npool_good_gw_v1_deviceinfo_deviceinfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_npool_good_gw_v1_deviceinfo_deviceinfo_proto_goTypes,
		DependencyIndexes: file_npool_good_gw_v1_deviceinfo_deviceinfo_proto_depIdxs,
		MessageInfos:      file_npool_good_gw_v1_deviceinfo_deviceinfo_proto_msgTypes,
	}.Build()
	File_npool_good_gw_v1_deviceinfo_deviceinfo_proto = out.File
	file_npool_good_gw_v1_deviceinfo_deviceinfo_proto_rawDesc = nil
	file_npool_good_gw_v1_deviceinfo_deviceinfo_proto_goTypes = nil
	file_npool_good_gw_v1_deviceinfo_deviceinfo_proto_depIdxs = nil
}
