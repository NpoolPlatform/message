// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.18.1
// source: npool/good/gw/v1/vender/location/location.proto

package location

import (
	location "github.com/NpoolPlatform/message/npool/good/mw/v1/vender/location"
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

type CreateLocationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Country  string `protobuf:"bytes,10,opt,name=Country,proto3" json:"Country,omitempty"`
	Province string `protobuf:"bytes,20,opt,name=Province,proto3" json:"Province,omitempty"`
	City     string `protobuf:"bytes,30,opt,name=City,proto3" json:"City,omitempty"`
	Address  string `protobuf:"bytes,40,opt,name=Address,proto3" json:"Address,omitempty"`
	BrandID  string `protobuf:"bytes,50,opt,name=BrandID,proto3" json:"BrandID,omitempty"`
}

func (x *CreateLocationRequest) Reset() {
	*x = CreateLocationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_good_gw_v1_vender_location_location_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateLocationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateLocationRequest) ProtoMessage() {}

func (x *CreateLocationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_good_gw_v1_vender_location_location_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateLocationRequest.ProtoReflect.Descriptor instead.
func (*CreateLocationRequest) Descriptor() ([]byte, []int) {
	return file_npool_good_gw_v1_vender_location_location_proto_rawDescGZIP(), []int{0}
}

func (x *CreateLocationRequest) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

func (x *CreateLocationRequest) GetProvince() string {
	if x != nil {
		return x.Province
	}
	return ""
}

func (x *CreateLocationRequest) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

func (x *CreateLocationRequest) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *CreateLocationRequest) GetBrandID() string {
	if x != nil {
		return x.BrandID
	}
	return ""
}

type CreateLocationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *location.Location `protobuf:"bytes,10,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (x *CreateLocationResponse) Reset() {
	*x = CreateLocationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_good_gw_v1_vender_location_location_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateLocationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateLocationResponse) ProtoMessage() {}

func (x *CreateLocationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_good_gw_v1_vender_location_location_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateLocationResponse.ProtoReflect.Descriptor instead.
func (*CreateLocationResponse) Descriptor() ([]byte, []int) {
	return file_npool_good_gw_v1_vender_location_location_proto_rawDescGZIP(), []int{1}
}

func (x *CreateLocationResponse) GetInfo() *location.Location {
	if x != nil {
		return x.Info
	}
	return nil
}

type UpdateLocationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID       uint32  `protobuf:"varint,9,opt,name=ID,proto3" json:"ID,omitempty"`
	EntID    string  `protobuf:"bytes,10,opt,name=EntID,proto3" json:"EntID,omitempty"`
	Country  *string `protobuf:"bytes,20,opt,name=Country,proto3,oneof" json:"Country,omitempty"`
	Province *string `protobuf:"bytes,30,opt,name=Province,proto3,oneof" json:"Province,omitempty"`
	City     *string `protobuf:"bytes,40,opt,name=City,proto3,oneof" json:"City,omitempty"`
	Address  *string `protobuf:"bytes,50,opt,name=Address,proto3,oneof" json:"Address,omitempty"`
	BrandID  *string `protobuf:"bytes,60,opt,name=BrandID,proto3,oneof" json:"BrandID,omitempty"`
}

func (x *UpdateLocationRequest) Reset() {
	*x = UpdateLocationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_good_gw_v1_vender_location_location_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateLocationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateLocationRequest) ProtoMessage() {}

func (x *UpdateLocationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_good_gw_v1_vender_location_location_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateLocationRequest.ProtoReflect.Descriptor instead.
func (*UpdateLocationRequest) Descriptor() ([]byte, []int) {
	return file_npool_good_gw_v1_vender_location_location_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateLocationRequest) GetID() uint32 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *UpdateLocationRequest) GetEntID() string {
	if x != nil {
		return x.EntID
	}
	return ""
}

func (x *UpdateLocationRequest) GetCountry() string {
	if x != nil && x.Country != nil {
		return *x.Country
	}
	return ""
}

func (x *UpdateLocationRequest) GetProvince() string {
	if x != nil && x.Province != nil {
		return *x.Province
	}
	return ""
}

func (x *UpdateLocationRequest) GetCity() string {
	if x != nil && x.City != nil {
		return *x.City
	}
	return ""
}

func (x *UpdateLocationRequest) GetAddress() string {
	if x != nil && x.Address != nil {
		return *x.Address
	}
	return ""
}

func (x *UpdateLocationRequest) GetBrandID() string {
	if x != nil && x.BrandID != nil {
		return *x.BrandID
	}
	return ""
}

type UpdateLocationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *location.Location `protobuf:"bytes,10,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (x *UpdateLocationResponse) Reset() {
	*x = UpdateLocationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_good_gw_v1_vender_location_location_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateLocationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateLocationResponse) ProtoMessage() {}

func (x *UpdateLocationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_good_gw_v1_vender_location_location_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateLocationResponse.ProtoReflect.Descriptor instead.
func (*UpdateLocationResponse) Descriptor() ([]byte, []int) {
	return file_npool_good_gw_v1_vender_location_location_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateLocationResponse) GetInfo() *location.Location {
	if x != nil {
		return x.Info
	}
	return nil
}

type GetLocationsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BrandID *string `protobuf:"bytes,10,opt,name=BrandID,proto3,oneof" json:"BrandID,omitempty"`
	Offset  int32   `protobuf:"varint,20,opt,name=Offset,proto3" json:"Offset,omitempty"`
	Limit   int32   `protobuf:"varint,30,opt,name=Limit,proto3" json:"Limit,omitempty"`
}

func (x *GetLocationsRequest) Reset() {
	*x = GetLocationsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_good_gw_v1_vender_location_location_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetLocationsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLocationsRequest) ProtoMessage() {}

func (x *GetLocationsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_good_gw_v1_vender_location_location_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLocationsRequest.ProtoReflect.Descriptor instead.
func (*GetLocationsRequest) Descriptor() ([]byte, []int) {
	return file_npool_good_gw_v1_vender_location_location_proto_rawDescGZIP(), []int{4}
}

func (x *GetLocationsRequest) GetBrandID() string {
	if x != nil && x.BrandID != nil {
		return *x.BrandID
	}
	return ""
}

func (x *GetLocationsRequest) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *GetLocationsRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type GetLocationsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Infos []*location.Location `protobuf:"bytes,10,rep,name=Infos,proto3" json:"Infos,omitempty"`
	Total uint32               `protobuf:"varint,20,opt,name=Total,proto3" json:"Total,omitempty"`
}

func (x *GetLocationsResponse) Reset() {
	*x = GetLocationsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_good_gw_v1_vender_location_location_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetLocationsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLocationsResponse) ProtoMessage() {}

func (x *GetLocationsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_good_gw_v1_vender_location_location_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLocationsResponse.ProtoReflect.Descriptor instead.
func (*GetLocationsResponse) Descriptor() ([]byte, []int) {
	return file_npool_good_gw_v1_vender_location_location_proto_rawDescGZIP(), []int{5}
}

func (x *GetLocationsResponse) GetInfos() []*location.Location {
	if x != nil {
		return x.Infos
	}
	return nil
}

func (x *GetLocationsResponse) GetTotal() uint32 {
	if x != nil {
		return x.Total
	}
	return 0
}

type DeleteLocationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID    uint32 `protobuf:"varint,9,opt,name=ID,proto3" json:"ID,omitempty"`
	EntID string `protobuf:"bytes,10,opt,name=EntID,proto3" json:"EntID,omitempty"`
}

func (x *DeleteLocationRequest) Reset() {
	*x = DeleteLocationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_good_gw_v1_vender_location_location_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteLocationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteLocationRequest) ProtoMessage() {}

func (x *DeleteLocationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_good_gw_v1_vender_location_location_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteLocationRequest.ProtoReflect.Descriptor instead.
func (*DeleteLocationRequest) Descriptor() ([]byte, []int) {
	return file_npool_good_gw_v1_vender_location_location_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteLocationRequest) GetID() uint32 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *DeleteLocationRequest) GetEntID() string {
	if x != nil {
		return x.EntID
	}
	return ""
}

type DeleteLocationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *location.Location `protobuf:"bytes,10,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (x *DeleteLocationResponse) Reset() {
	*x = DeleteLocationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_good_gw_v1_vender_location_location_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteLocationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteLocationResponse) ProtoMessage() {}

func (x *DeleteLocationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_good_gw_v1_vender_location_location_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteLocationResponse.ProtoReflect.Descriptor instead.
func (*DeleteLocationResponse) Descriptor() ([]byte, []int) {
	return file_npool_good_gw_v1_vender_location_location_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteLocationResponse) GetInfo() *location.Location {
	if x != nil {
		return x.Info
	}
	return nil
}

var File_npool_good_gw_v1_vender_location_location_proto protoreflect.FileDescriptor

var file_npool_good_gw_v1_vender_location_location_proto_rawDesc = []byte{
	0x0a, 0x2f, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x67, 0x6f, 0x6f, 0x64, 0x2f, 0x67, 0x77, 0x2f,
	0x76, 0x31, 0x2f, 0x76, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x1f, 0x67, 0x6f, 0x6f, 0x64, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e,
	0x76, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x2e, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x76, 0x31, 0x1a, 0x2f, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x67, 0x6f, 0x6f, 0x64, 0x2f, 0x6d,
	0x77, 0x2f, 0x76, 0x31, 0x2f, 0x76, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x2f, 0x6c, 0x6f, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x95, 0x01, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x6f, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63,
	0x65, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x43, 0x69, 0x74, 0x79, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x43, 0x69, 0x74, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x18, 0x28, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12,
	0x18, 0x0a, 0x07, 0x42, 0x72, 0x61, 0x6e, 0x64, 0x49, 0x44, 0x18, 0x32, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x42, 0x72, 0x61, 0x6e, 0x64, 0x49, 0x44, 0x22, 0x5a, 0x0a, 0x16, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x40, 0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x2c, 0x2e, 0x67, 0x6f, 0x6f, 0x64, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77,
	0x61, 0x72, 0x65, 0x2e, 0x76, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x2e, 0x6c, 0x6f, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x04, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x8e, 0x02, 0x0a, 0x15, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x49, 0x44, 0x12,
	0x14, 0x0a, 0x05, 0x45, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x45, 0x6e, 0x74, 0x49, 0x44, 0x12, 0x1d, 0x0a, 0x07, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79,
	0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x07, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72,
	0x79, 0x88, 0x01, 0x01, 0x12, 0x1f, 0x0a, 0x08, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65,
	0x18, 0x1e, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x08, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x6e,
	0x63, 0x65, 0x88, 0x01, 0x01, 0x12, 0x17, 0x0a, 0x04, 0x43, 0x69, 0x74, 0x79, 0x18, 0x28, 0x20,
	0x01, 0x28, 0x09, 0x48, 0x02, 0x52, 0x04, 0x43, 0x69, 0x74, 0x79, 0x88, 0x01, 0x01, 0x12, 0x1d,
	0x0a, 0x07, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x32, 0x20, 0x01, 0x28, 0x09, 0x48,
	0x03, 0x52, 0x07, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x88, 0x01, 0x01, 0x12, 0x1d, 0x0a,
	0x07, 0x42, 0x72, 0x61, 0x6e, 0x64, 0x49, 0x44, 0x18, 0x3c, 0x20, 0x01, 0x28, 0x09, 0x48, 0x04,
	0x52, 0x07, 0x42, 0x72, 0x61, 0x6e, 0x64, 0x49, 0x44, 0x88, 0x01, 0x01, 0x42, 0x0a, 0x0a, 0x08,
	0x5f, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x50, 0x72, 0x6f,
	0x76, 0x69, 0x6e, 0x63, 0x65, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x43, 0x69, 0x74, 0x79, 0x42, 0x0a,
	0x0a, 0x08, 0x5f, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x42,
	0x72, 0x61, 0x6e, 0x64, 0x49, 0x44, 0x22, 0x5a, 0x0a, 0x16, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x40, 0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2c,
	0x2e, 0x67, 0x6f, 0x6f, 0x64, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65,
	0x2e, 0x76, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x2e, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x04, 0x49, 0x6e,
	0x66, 0x6f, 0x22, 0x6e, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x07, 0x42, 0x72, 0x61,
	0x6e, 0x64, 0x49, 0x44, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x07, 0x42, 0x72,
	0x61, 0x6e, 0x64, 0x49, 0x44, 0x88, 0x01, 0x01, 0x12, 0x16, 0x0a, 0x06, 0x4f, 0x66, 0x66, 0x73,
	0x65, 0x74, 0x18, 0x14, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74,
	0x12, 0x14, 0x0a, 0x05, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x05, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x42, 0x72, 0x61, 0x6e, 0x64,
	0x49, 0x44, 0x22, 0x70, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x42, 0x0a, 0x05, 0x49, 0x6e,
	0x66, 0x6f, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x67, 0x6f, 0x6f, 0x64,
	0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x76, 0x65, 0x6e, 0x64,
	0x6f, 0x72, 0x2e, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x4c,
	0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x05, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x12, 0x14,
	0x0a, 0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x14, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x54,
	0x6f, 0x74, 0x61, 0x6c, 0x22, 0x3d, 0x0a, 0x15, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4c, 0x6f,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x49, 0x44, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x49, 0x44, 0x12, 0x14, 0x0a,
	0x05, 0x45, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x45, 0x6e,
	0x74, 0x49, 0x44, 0x22, 0x5a, 0x0a, 0x16, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4c, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x40, 0x0a,
	0x04, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x67, 0x6f,
	0x6f, 0x64, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x76, 0x65,
	0x6e, 0x64, 0x6f, 0x72, 0x2e, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31,
	0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x32,
	0xa9, 0x05, 0x0a, 0x07, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x12, 0xa7, 0x01, 0x0a, 0x0e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x36,
	0x2e, 0x67, 0x6f, 0x6f, 0x64, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x76, 0x65,
	0x6e, 0x64, 0x6f, 0x72, 0x2e, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x37, 0x2e, 0x67, 0x6f, 0x6f, 0x64, 0x2e, 0x67, 0x61,
	0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x76, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x2e, 0x6c, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c,
	0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x24, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1e, 0x3a, 0x01, 0x2a, 0x22, 0x19, 0x2f, 0x76, 0x31, 0x2f,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x6c, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0xa7, 0x01, 0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x36, 0x2e, 0x67, 0x6f, 0x6f, 0x64, 0x2e,
	0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x76, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x2e, 0x6c,
	0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x37, 0x2e, 0x67, 0x6f, 0x6f, 0x64, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e,
	0x76, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x2e, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x24, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x1e, 0x3a, 0x01, 0x2a, 0x22, 0x19, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x2f, 0x76, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x9f, 0x01, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x12, 0x34, 0x2e, 0x67, 0x6f, 0x6f, 0x64, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e,
	0x76, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x2e, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x35, 0x2e, 0x67, 0x6f, 0x6f, 0x64, 0x2e, 0x67, 0x61,
	0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x76, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x2e, 0x6c, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x6f, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x22, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x1c, 0x3a, 0x01, 0x2a, 0x22, 0x17, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x65,
	0x74, 0x2f, 0x76, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x12, 0xa7, 0x01, 0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4c, 0x6f, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x36, 0x2e, 0x67, 0x6f, 0x6f, 0x64, 0x2e, 0x67, 0x61, 0x74, 0x65,
	0x77, 0x61, 0x79, 0x2e, 0x76, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x2e, 0x6c, 0x6f, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4c, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x37, 0x2e, 0x67,
	0x6f, 0x6f, 0x64, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x76, 0x65, 0x6e, 0x64,
	0x6f, 0x72, 0x2e, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x24, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1e, 0x3a, 0x01, 0x2a,
	0x22, 0x19, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x2f, 0x76, 0x65, 0x6e,
	0x64, 0x6f, 0x72, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x43, 0x5a, 0x41, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4e, 0x70, 0x6f, 0x6f, 0x6c, 0x50,
	0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f,
	0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x67, 0x6f, 0x6f, 0x64, 0x2f, 0x67, 0x77, 0x2f, 0x76, 0x31,
	0x2f, 0x76, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_npool_good_gw_v1_vender_location_location_proto_rawDescOnce sync.Once
	file_npool_good_gw_v1_vender_location_location_proto_rawDescData = file_npool_good_gw_v1_vender_location_location_proto_rawDesc
)

func file_npool_good_gw_v1_vender_location_location_proto_rawDescGZIP() []byte {
	file_npool_good_gw_v1_vender_location_location_proto_rawDescOnce.Do(func() {
		file_npool_good_gw_v1_vender_location_location_proto_rawDescData = protoimpl.X.CompressGZIP(file_npool_good_gw_v1_vender_location_location_proto_rawDescData)
	})
	return file_npool_good_gw_v1_vender_location_location_proto_rawDescData
}

var file_npool_good_gw_v1_vender_location_location_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_npool_good_gw_v1_vender_location_location_proto_goTypes = []interface{}{
	(*CreateLocationRequest)(nil),  // 0: good.gateway.vendor.location.v1.CreateLocationRequest
	(*CreateLocationResponse)(nil), // 1: good.gateway.vendor.location.v1.CreateLocationResponse
	(*UpdateLocationRequest)(nil),  // 2: good.gateway.vendor.location.v1.UpdateLocationRequest
	(*UpdateLocationResponse)(nil), // 3: good.gateway.vendor.location.v1.UpdateLocationResponse
	(*GetLocationsRequest)(nil),    // 4: good.gateway.vendor.location.v1.GetLocationsRequest
	(*GetLocationsResponse)(nil),   // 5: good.gateway.vendor.location.v1.GetLocationsResponse
	(*DeleteLocationRequest)(nil),  // 6: good.gateway.vendor.location.v1.DeleteLocationRequest
	(*DeleteLocationResponse)(nil), // 7: good.gateway.vendor.location.v1.DeleteLocationResponse
	(*location.Location)(nil),      // 8: good.middleware.vendor.location.v1.Location
}
var file_npool_good_gw_v1_vender_location_location_proto_depIdxs = []int32{
	8, // 0: good.gateway.vendor.location.v1.CreateLocationResponse.Info:type_name -> good.middleware.vendor.location.v1.Location
	8, // 1: good.gateway.vendor.location.v1.UpdateLocationResponse.Info:type_name -> good.middleware.vendor.location.v1.Location
	8, // 2: good.gateway.vendor.location.v1.GetLocationsResponse.Infos:type_name -> good.middleware.vendor.location.v1.Location
	8, // 3: good.gateway.vendor.location.v1.DeleteLocationResponse.Info:type_name -> good.middleware.vendor.location.v1.Location
	0, // 4: good.gateway.vendor.location.v1.Gateway.CreateLocation:input_type -> good.gateway.vendor.location.v1.CreateLocationRequest
	2, // 5: good.gateway.vendor.location.v1.Gateway.UpdateLocation:input_type -> good.gateway.vendor.location.v1.UpdateLocationRequest
	4, // 6: good.gateway.vendor.location.v1.Gateway.GetLocations:input_type -> good.gateway.vendor.location.v1.GetLocationsRequest
	6, // 7: good.gateway.vendor.location.v1.Gateway.DeleteLocation:input_type -> good.gateway.vendor.location.v1.DeleteLocationRequest
	1, // 8: good.gateway.vendor.location.v1.Gateway.CreateLocation:output_type -> good.gateway.vendor.location.v1.CreateLocationResponse
	3, // 9: good.gateway.vendor.location.v1.Gateway.UpdateLocation:output_type -> good.gateway.vendor.location.v1.UpdateLocationResponse
	5, // 10: good.gateway.vendor.location.v1.Gateway.GetLocations:output_type -> good.gateway.vendor.location.v1.GetLocationsResponse
	7, // 11: good.gateway.vendor.location.v1.Gateway.DeleteLocation:output_type -> good.gateway.vendor.location.v1.DeleteLocationResponse
	8, // [8:12] is the sub-list for method output_type
	4, // [4:8] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_npool_good_gw_v1_vender_location_location_proto_init() }
func file_npool_good_gw_v1_vender_location_location_proto_init() {
	if File_npool_good_gw_v1_vender_location_location_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_npool_good_gw_v1_vender_location_location_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateLocationRequest); i {
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
		file_npool_good_gw_v1_vender_location_location_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateLocationResponse); i {
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
		file_npool_good_gw_v1_vender_location_location_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateLocationRequest); i {
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
		file_npool_good_gw_v1_vender_location_location_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateLocationResponse); i {
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
		file_npool_good_gw_v1_vender_location_location_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetLocationsRequest); i {
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
		file_npool_good_gw_v1_vender_location_location_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetLocationsResponse); i {
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
		file_npool_good_gw_v1_vender_location_location_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteLocationRequest); i {
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
		file_npool_good_gw_v1_vender_location_location_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteLocationResponse); i {
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
	file_npool_good_gw_v1_vender_location_location_proto_msgTypes[2].OneofWrappers = []interface{}{}
	file_npool_good_gw_v1_vender_location_location_proto_msgTypes[4].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_npool_good_gw_v1_vender_location_location_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_npool_good_gw_v1_vender_location_location_proto_goTypes,
		DependencyIndexes: file_npool_good_gw_v1_vender_location_location_proto_depIdxs,
		MessageInfos:      file_npool_good_gw_v1_vender_location_location_proto_msgTypes,
	}.Build()
	File_npool_good_gw_v1_vender_location_location_proto = out.File
	file_npool_good_gw_v1_vender_location_location_proto_rawDesc = nil
	file_npool_good_gw_v1_vender_location_location_proto_goTypes = nil
	file_npool_good_gw_v1_vender_location_location_proto_depIdxs = nil
}
