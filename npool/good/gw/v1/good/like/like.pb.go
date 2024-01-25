// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.18.1
// source: npool/good/gw/v1/good/like/like.proto

package like

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

type Like struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID           uint32  `protobuf:"varint,9,opt,name=ID,proto3" json:"ID,omitempty"`
	EntID        string  `protobuf:"bytes,10,opt,name=EntID,proto3" json:"EntID,omitempty"`
	AppID        string  `protobuf:"bytes,20,opt,name=AppID,proto3" json:"AppID,omitempty"`
	AppName      string  `protobuf:"bytes,30,opt,name=AppName,proto3" json:"AppName,omitempty"`
	UserID       string  `protobuf:"bytes,40,opt,name=UserID,proto3" json:"UserID,omitempty"`
	Username     *string `protobuf:"bytes,50,opt,name=Username,proto3,oneof" json:"Username,omitempty"`
	EmailAddress *string `protobuf:"bytes,60,opt,name=EmailAddress,proto3,oneof" json:"EmailAddress,omitempty"`
	PhoneNO      *string `protobuf:"bytes,70,opt,name=PhoneNO,proto3,oneof" json:"PhoneNO,omitempty"`
	AppGoodID    string  `protobuf:"bytes,80,opt,name=AppGoodID,proto3" json:"AppGoodID,omitempty"`
	GoodName     string  `protobuf:"bytes,90,opt,name=GoodName,proto3" json:"GoodName,omitempty"`
	Like         bool    `protobuf:"varint,100,opt,name=Like,proto3" json:"Like,omitempty"`
	GoodID       string  `protobuf:"bytes,110,opt,name=GoodID,proto3" json:"GoodID,omitempty"`
	CreatedAt    uint32  `protobuf:"varint,1000,opt,name=CreatedAt,proto3" json:"CreatedAt,omitempty"`
}

func (x *Like) Reset() {
	*x = Like{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_good_gw_v1_good_like_like_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Like) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Like) ProtoMessage() {}

func (x *Like) ProtoReflect() protoreflect.Message {
	mi := &file_npool_good_gw_v1_good_like_like_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Like.ProtoReflect.Descriptor instead.
func (*Like) Descriptor() ([]byte, []int) {
	return file_npool_good_gw_v1_good_like_like_proto_rawDescGZIP(), []int{0}
}

func (x *Like) GetID() uint32 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *Like) GetEntID() string {
	if x != nil {
		return x.EntID
	}
	return ""
}

func (x *Like) GetAppID() string {
	if x != nil {
		return x.AppID
	}
	return ""
}

func (x *Like) GetAppName() string {
	if x != nil {
		return x.AppName
	}
	return ""
}

func (x *Like) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

func (x *Like) GetUsername() string {
	if x != nil && x.Username != nil {
		return *x.Username
	}
	return ""
}

func (x *Like) GetEmailAddress() string {
	if x != nil && x.EmailAddress != nil {
		return *x.EmailAddress
	}
	return ""
}

func (x *Like) GetPhoneNO() string {
	if x != nil && x.PhoneNO != nil {
		return *x.PhoneNO
	}
	return ""
}

func (x *Like) GetAppGoodID() string {
	if x != nil {
		return x.AppGoodID
	}
	return ""
}

func (x *Like) GetGoodName() string {
	if x != nil {
		return x.GoodName
	}
	return ""
}

func (x *Like) GetLike() bool {
	if x != nil {
		return x.Like
	}
	return false
}

func (x *Like) GetGoodID() string {
	if x != nil {
		return x.GoodID
	}
	return ""
}

func (x *Like) GetCreatedAt() uint32 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

type CreateLikeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AppID     string `protobuf:"bytes,10,opt,name=AppID,proto3" json:"AppID,omitempty"`
	UserID    string `protobuf:"bytes,20,opt,name=UserID,proto3" json:"UserID,omitempty"`
	AppGoodID string `protobuf:"bytes,30,opt,name=AppGoodID,proto3" json:"AppGoodID,omitempty"`
	Like      bool   `protobuf:"varint,40,opt,name=Like,proto3" json:"Like,omitempty"`
}

func (x *CreateLikeRequest) Reset() {
	*x = CreateLikeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_good_gw_v1_good_like_like_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateLikeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateLikeRequest) ProtoMessage() {}

func (x *CreateLikeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_good_gw_v1_good_like_like_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateLikeRequest.ProtoReflect.Descriptor instead.
func (*CreateLikeRequest) Descriptor() ([]byte, []int) {
	return file_npool_good_gw_v1_good_like_like_proto_rawDescGZIP(), []int{1}
}

func (x *CreateLikeRequest) GetAppID() string {
	if x != nil {
		return x.AppID
	}
	return ""
}

func (x *CreateLikeRequest) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

func (x *CreateLikeRequest) GetAppGoodID() string {
	if x != nil {
		return x.AppGoodID
	}
	return ""
}

func (x *CreateLikeRequest) GetLike() bool {
	if x != nil {
		return x.Like
	}
	return false
}

type CreateLikeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *Like `protobuf:"bytes,10,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (x *CreateLikeResponse) Reset() {
	*x = CreateLikeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_good_gw_v1_good_like_like_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateLikeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateLikeResponse) ProtoMessage() {}

func (x *CreateLikeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_good_gw_v1_good_like_like_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateLikeResponse.ProtoReflect.Descriptor instead.
func (*CreateLikeResponse) Descriptor() ([]byte, []int) {
	return file_npool_good_gw_v1_good_like_like_proto_rawDescGZIP(), []int{2}
}

func (x *CreateLikeResponse) GetInfo() *Like {
	if x != nil {
		return x.Info
	}
	return nil
}

type GetLikesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AppID     string  `protobuf:"bytes,10,opt,name=AppID,proto3" json:"AppID,omitempty"`
	AppGoodID *string `protobuf:"bytes,20,opt,name=AppGoodID,proto3,oneof" json:"AppGoodID,omitempty"`
	GoodID    *string `protobuf:"bytes,30,opt,name=GoodID,proto3,oneof" json:"GoodID,omitempty"`
	Offset    int32   `protobuf:"varint,40,opt,name=Offset,proto3" json:"Offset,omitempty"`
	Limit     int32   `protobuf:"varint,50,opt,name=Limit,proto3" json:"Limit,omitempty"`
}

func (x *GetLikesRequest) Reset() {
	*x = GetLikesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_good_gw_v1_good_like_like_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetLikesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLikesRequest) ProtoMessage() {}

func (x *GetLikesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_good_gw_v1_good_like_like_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLikesRequest.ProtoReflect.Descriptor instead.
func (*GetLikesRequest) Descriptor() ([]byte, []int) {
	return file_npool_good_gw_v1_good_like_like_proto_rawDescGZIP(), []int{3}
}

func (x *GetLikesRequest) GetAppID() string {
	if x != nil {
		return x.AppID
	}
	return ""
}

func (x *GetLikesRequest) GetAppGoodID() string {
	if x != nil && x.AppGoodID != nil {
		return *x.AppGoodID
	}
	return ""
}

func (x *GetLikesRequest) GetGoodID() string {
	if x != nil && x.GoodID != nil {
		return *x.GoodID
	}
	return ""
}

func (x *GetLikesRequest) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *GetLikesRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type GetLikesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Infos []*Like `protobuf:"bytes,10,rep,name=Infos,proto3" json:"Infos,omitempty"`
	Total uint32  `protobuf:"varint,20,opt,name=Total,proto3" json:"Total,omitempty"`
}

func (x *GetLikesResponse) Reset() {
	*x = GetLikesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_good_gw_v1_good_like_like_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetLikesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLikesResponse) ProtoMessage() {}

func (x *GetLikesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_good_gw_v1_good_like_like_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLikesResponse.ProtoReflect.Descriptor instead.
func (*GetLikesResponse) Descriptor() ([]byte, []int) {
	return file_npool_good_gw_v1_good_like_like_proto_rawDescGZIP(), []int{4}
}

func (x *GetLikesResponse) GetInfos() []*Like {
	if x != nil {
		return x.Infos
	}
	return nil
}

func (x *GetLikesResponse) GetTotal() uint32 {
	if x != nil {
		return x.Total
	}
	return 0
}

type GetMyLikesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AppID  string `protobuf:"bytes,10,opt,name=AppID,proto3" json:"AppID,omitempty"`
	UserID string `protobuf:"bytes,20,opt,name=UserID,proto3" json:"UserID,omitempty"`
	Offset int32  `protobuf:"varint,30,opt,name=Offset,proto3" json:"Offset,omitempty"`
	Limit  int32  `protobuf:"varint,40,opt,name=Limit,proto3" json:"Limit,omitempty"`
}

func (x *GetMyLikesRequest) Reset() {
	*x = GetMyLikesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_good_gw_v1_good_like_like_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMyLikesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMyLikesRequest) ProtoMessage() {}

func (x *GetMyLikesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_good_gw_v1_good_like_like_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMyLikesRequest.ProtoReflect.Descriptor instead.
func (*GetMyLikesRequest) Descriptor() ([]byte, []int) {
	return file_npool_good_gw_v1_good_like_like_proto_rawDescGZIP(), []int{5}
}

func (x *GetMyLikesRequest) GetAppID() string {
	if x != nil {
		return x.AppID
	}
	return ""
}

func (x *GetMyLikesRequest) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

func (x *GetMyLikesRequest) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *GetMyLikesRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type GetMyLikesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Infos []*Like `protobuf:"bytes,10,rep,name=Infos,proto3" json:"Infos,omitempty"`
	Total uint32  `protobuf:"varint,20,opt,name=Total,proto3" json:"Total,omitempty"`
}

func (x *GetMyLikesResponse) Reset() {
	*x = GetMyLikesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_good_gw_v1_good_like_like_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMyLikesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMyLikesResponse) ProtoMessage() {}

func (x *GetMyLikesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_good_gw_v1_good_like_like_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMyLikesResponse.ProtoReflect.Descriptor instead.
func (*GetMyLikesResponse) Descriptor() ([]byte, []int) {
	return file_npool_good_gw_v1_good_like_like_proto_rawDescGZIP(), []int{6}
}

func (x *GetMyLikesResponse) GetInfos() []*Like {
	if x != nil {
		return x.Infos
	}
	return nil
}

func (x *GetMyLikesResponse) GetTotal() uint32 {
	if x != nil {
		return x.Total
	}
	return 0
}

type DeleteLikeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID     uint32 `protobuf:"varint,9,opt,name=ID,proto3" json:"ID,omitempty"`
	EntID  string `protobuf:"bytes,10,opt,name=EntID,proto3" json:"EntID,omitempty"`
	AppID  string `protobuf:"bytes,20,opt,name=AppID,proto3" json:"AppID,omitempty"`
	UserID string `protobuf:"bytes,30,opt,name=UserID,proto3" json:"UserID,omitempty"`
}

func (x *DeleteLikeRequest) Reset() {
	*x = DeleteLikeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_good_gw_v1_good_like_like_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteLikeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteLikeRequest) ProtoMessage() {}

func (x *DeleteLikeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_good_gw_v1_good_like_like_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteLikeRequest.ProtoReflect.Descriptor instead.
func (*DeleteLikeRequest) Descriptor() ([]byte, []int) {
	return file_npool_good_gw_v1_good_like_like_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteLikeRequest) GetID() uint32 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *DeleteLikeRequest) GetEntID() string {
	if x != nil {
		return x.EntID
	}
	return ""
}

func (x *DeleteLikeRequest) GetAppID() string {
	if x != nil {
		return x.AppID
	}
	return ""
}

func (x *DeleteLikeRequest) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

type DeleteLikeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *Like `protobuf:"bytes,10,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (x *DeleteLikeResponse) Reset() {
	*x = DeleteLikeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_good_gw_v1_good_like_like_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteLikeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteLikeResponse) ProtoMessage() {}

func (x *DeleteLikeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_good_gw_v1_good_like_like_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteLikeResponse.ProtoReflect.Descriptor instead.
func (*DeleteLikeResponse) Descriptor() ([]byte, []int) {
	return file_npool_good_gw_v1_good_like_like_proto_rawDescGZIP(), []int{8}
}

func (x *DeleteLikeResponse) GetInfo() *Like {
	if x != nil {
		return x.Info
	}
	return nil
}

var File_npool_good_gw_v1_good_like_like_proto protoreflect.FileDescriptor

var file_npool_good_gw_v1_good_like_like_proto_rawDesc = []byte{
	0x0a, 0x25, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x67, 0x6f, 0x6f, 0x64, 0x2f, 0x67, 0x77, 0x2f,
	0x76, 0x31, 0x2f, 0x67, 0x6f, 0x6f, 0x64, 0x2f, 0x6c, 0x69, 0x6b, 0x65, 0x2f, 0x6c, 0x69, 0x6b,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1a, 0x67, 0x6f, 0x6f, 0x64, 0x2e, 0x67, 0x61,
	0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x67, 0x6f, 0x6f, 0x64, 0x31, 0x2e, 0x6c, 0x69, 0x6b, 0x65,
	0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x8c, 0x03, 0x0a, 0x04, 0x4c, 0x69, 0x6b, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x45, 0x6e,
	0x74, 0x49, 0x44, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x45, 0x6e, 0x74, 0x49, 0x44,
	0x12, 0x14, 0x0a, 0x05, 0x41, 0x70, 0x70, 0x49, 0x44, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x41, 0x70, 0x70, 0x49, 0x44, 0x12, 0x18, 0x0a, 0x07, 0x41, 0x70, 0x70, 0x4e, 0x61, 0x6d,
	0x65, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x41, 0x70, 0x70, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x28, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x1f, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x32, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x08, 0x55, 0x73,
	0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x27, 0x0a, 0x0c, 0x45, 0x6d, 0x61,
	0x69, 0x6c, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x3c, 0x20, 0x01, 0x28, 0x09, 0x48,
	0x01, 0x52, 0x0c, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x88,
	0x01, 0x01, 0x12, 0x1d, 0x0a, 0x07, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x4f, 0x18, 0x46, 0x20,
	0x01, 0x28, 0x09, 0x48, 0x02, 0x52, 0x07, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x4f, 0x88, 0x01,
	0x01, 0x12, 0x1c, 0x0a, 0x09, 0x41, 0x70, 0x70, 0x47, 0x6f, 0x6f, 0x64, 0x49, 0x44, 0x18, 0x50,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x41, 0x70, 0x70, 0x47, 0x6f, 0x6f, 0x64, 0x49, 0x44, 0x12,
	0x1a, 0x0a, 0x08, 0x47, 0x6f, 0x6f, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x5a, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x47, 0x6f, 0x6f, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x4c,
	0x69, 0x6b, 0x65, 0x18, 0x64, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x4c, 0x69, 0x6b, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x47, 0x6f, 0x6f, 0x64, 0x49, 0x44, 0x18, 0x6e, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x47, 0x6f, 0x6f, 0x64, 0x49, 0x44, 0x12, 0x1d, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x18, 0xe8, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x55, 0x73, 0x65, 0x72, 0x6e,
	0x61, 0x6d, 0x65, 0x42, 0x0f, 0x0a, 0x0d, 0x5f, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x41, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x4f,
	0x22, 0x73, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6b, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x41, 0x70, 0x70, 0x49, 0x44, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x41, 0x70, 0x70, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x55,
	0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x55, 0x73, 0x65,
	0x72, 0x49, 0x44, 0x12, 0x1c, 0x0a, 0x09, 0x41, 0x70, 0x70, 0x47, 0x6f, 0x6f, 0x64, 0x49, 0x44,
	0x18, 0x1e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x41, 0x70, 0x70, 0x47, 0x6f, 0x6f, 0x64, 0x49,
	0x44, 0x12, 0x12, 0x0a, 0x04, 0x4c, 0x69, 0x6b, 0x65, 0x18, 0x28, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x04, 0x4c, 0x69, 0x6b, 0x65, 0x22, 0x4a, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c,
	0x69, 0x6b, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x34, 0x0a, 0x04, 0x49,
	0x6e, 0x66, 0x6f, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x67, 0x6f, 0x6f, 0x64,
	0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x67, 0x6f, 0x6f, 0x64, 0x31, 0x2e, 0x6c,
	0x69, 0x6b, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x6b, 0x65, 0x52, 0x04, 0x49, 0x6e, 0x66,
	0x6f, 0x22, 0xae, 0x01, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x6b, 0x65, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x41, 0x70, 0x70, 0x49, 0x44, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x41, 0x70, 0x70, 0x49, 0x44, 0x12, 0x21, 0x0a, 0x09, 0x41,
	0x70, 0x70, 0x47, 0x6f, 0x6f, 0x64, 0x49, 0x44, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00,
	0x52, 0x09, 0x41, 0x70, 0x70, 0x47, 0x6f, 0x6f, 0x64, 0x49, 0x44, 0x88, 0x01, 0x01, 0x12, 0x1b,
	0x0a, 0x06, 0x47, 0x6f, 0x6f, 0x64, 0x49, 0x44, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01,
	0x52, 0x06, 0x47, 0x6f, 0x6f, 0x64, 0x49, 0x44, 0x88, 0x01, 0x01, 0x12, 0x16, 0x0a, 0x06, 0x4f,
	0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x28, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x4f, 0x66, 0x66,
	0x73, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x32, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x05, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x41, 0x70,
	0x70, 0x47, 0x6f, 0x6f, 0x64, 0x49, 0x44, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x47, 0x6f, 0x6f, 0x64,
	0x49, 0x44, 0x22, 0x60, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x6b, 0x65, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x36, 0x0a, 0x05, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x18,
	0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x67, 0x6f, 0x6f, 0x64, 0x2e, 0x67, 0x61, 0x74,
	0x65, 0x77, 0x61, 0x79, 0x2e, 0x67, 0x6f, 0x6f, 0x64, 0x31, 0x2e, 0x6c, 0x69, 0x6b, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x4c, 0x69, 0x6b, 0x65, 0x52, 0x05, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x12, 0x14,
	0x0a, 0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x14, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x54,
	0x6f, 0x74, 0x61, 0x6c, 0x22, 0x6f, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x4d, 0x79, 0x4c, 0x69, 0x6b,
	0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x41, 0x70, 0x70,
	0x49, 0x44, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x41, 0x70, 0x70, 0x49, 0x44, 0x12,
	0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x4f, 0x66, 0x66, 0x73, 0x65,
	0x74, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12,
	0x14, 0x0a, 0x05, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x28, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05,
	0x4c, 0x69, 0x6d, 0x69, 0x74, 0x22, 0x62, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x4d, 0x79, 0x4c, 0x69,
	0x6b, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x36, 0x0a, 0x05, 0x49,
	0x6e, 0x66, 0x6f, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x67, 0x6f, 0x6f,
	0x64, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x67, 0x6f, 0x6f, 0x64, 0x31, 0x2e,
	0x6c, 0x69, 0x6b, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x6b, 0x65, 0x52, 0x05, 0x49, 0x6e,
	0x66, 0x6f, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x14, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x22, 0x67, 0x0a, 0x11, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x4c, 0x69, 0x6b, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x49, 0x44, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x49, 0x44, 0x12, 0x14,
	0x0a, 0x05, 0x45, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x45,
	0x6e, 0x74, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x41, 0x70, 0x70, 0x49, 0x44, 0x18, 0x14, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x41, 0x70, 0x70, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73,
	0x65, 0x72, 0x49, 0x44, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72,
	0x49, 0x44, 0x22, 0x4a, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4c, 0x69, 0x6b, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x34, 0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x67, 0x6f, 0x6f, 0x64, 0x2e, 0x67, 0x61,
	0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x67, 0x6f, 0x6f, 0x64, 0x31, 0x2e, 0x6c, 0x69, 0x6b, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x6b, 0x65, 0x52, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x32, 0xb9,
	0x04, 0x0a, 0x07, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x12, 0x8b, 0x01, 0x0a, 0x0a, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6b, 0x65, 0x12, 0x2d, 0x2e, 0x67, 0x6f, 0x6f, 0x64,
	0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x67, 0x6f, 0x6f, 0x64, 0x31, 0x2e, 0x6c,
	0x69, 0x6b, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6b,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2e, 0x2e, 0x67, 0x6f, 0x6f, 0x64, 0x2e,
	0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x67, 0x6f, 0x6f, 0x64, 0x31, 0x2e, 0x6c, 0x69,
	0x6b, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6b, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x18,
	0x3a, 0x01, 0x2a, 0x22, 0x13, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x2f,
	0x67, 0x6f, 0x6f, 0x64, 0x6c, 0x69, 0x6b, 0x65, 0x12, 0x83, 0x01, 0x0a, 0x08, 0x47, 0x65, 0x74,
	0x4c, 0x69, 0x6b, 0x65, 0x73, 0x12, 0x2b, 0x2e, 0x67, 0x6f, 0x6f, 0x64, 0x2e, 0x67, 0x61, 0x74,
	0x65, 0x77, 0x61, 0x79, 0x2e, 0x67, 0x6f, 0x6f, 0x64, 0x31, 0x2e, 0x6c, 0x69, 0x6b, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x6b, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x2c, 0x2e, 0x67, 0x6f, 0x6f, 0x64, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61,
	0x79, 0x2e, 0x67, 0x6f, 0x6f, 0x64, 0x31, 0x2e, 0x6c, 0x69, 0x6b, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x47, 0x65, 0x74, 0x4c, 0x69, 0x6b, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x1c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x16, 0x3a, 0x01, 0x2a, 0x22, 0x11, 0x2f, 0x76, 0x31,
	0x2f, 0x67, 0x65, 0x74, 0x2f, 0x67, 0x6f, 0x6f, 0x64, 0x6c, 0x69, 0x6b, 0x65, 0x73, 0x12, 0x8b,
	0x01, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x4d, 0x79, 0x4c, 0x69, 0x6b, 0x65, 0x73, 0x12, 0x2d, 0x2e,
	0x67, 0x6f, 0x6f, 0x64, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x67, 0x6f, 0x6f,
	0x64, 0x31, 0x2e, 0x6c, 0x69, 0x6b, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x79,
	0x4c, 0x69, 0x6b, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2e, 0x2e, 0x67,
	0x6f, 0x6f, 0x64, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x67, 0x6f, 0x6f, 0x64,
	0x31, 0x2e, 0x6c, 0x69, 0x6b, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x79, 0x4c,
	0x69, 0x6b, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1e, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x18, 0x3a, 0x01, 0x2a, 0x22, 0x13, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x65, 0x74,
	0x2f, 0x6d, 0x79, 0x67, 0x6f, 0x6f, 0x64, 0x6c, 0x69, 0x6b, 0x65, 0x73, 0x12, 0x8b, 0x01, 0x0a,
	0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4c, 0x69, 0x6b, 0x65, 0x12, 0x2d, 0x2e, 0x67, 0x6f,
	0x6f, 0x64, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x67, 0x6f, 0x6f, 0x64, 0x31,
	0x2e, 0x6c, 0x69, 0x6b, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4c,
	0x69, 0x6b, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2e, 0x2e, 0x67, 0x6f, 0x6f,
	0x64, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x67, 0x6f, 0x6f, 0x64, 0x31, 0x2e,
	0x6c, 0x69, 0x6b, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4c, 0x69,
	0x6b, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1e, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x18, 0x3a, 0x01, 0x2a, 0x22, 0x13, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x2f, 0x67, 0x6f, 0x6f, 0x64, 0x6c, 0x69, 0x6b, 0x65, 0x42, 0x3d, 0x5a, 0x3b, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4e, 0x70, 0x6f, 0x6f, 0x6c, 0x50, 0x6c,
	0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x6e,
	0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x67, 0x6f, 0x6f, 0x64, 0x2f, 0x67, 0x77, 0x2f, 0x76, 0x31, 0x2f,
	0x67, 0x6f, 0x6f, 0x64, 0x2f, 0x6c, 0x69, 0x6b, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_npool_good_gw_v1_good_like_like_proto_rawDescOnce sync.Once
	file_npool_good_gw_v1_good_like_like_proto_rawDescData = file_npool_good_gw_v1_good_like_like_proto_rawDesc
)

func file_npool_good_gw_v1_good_like_like_proto_rawDescGZIP() []byte {
	file_npool_good_gw_v1_good_like_like_proto_rawDescOnce.Do(func() {
		file_npool_good_gw_v1_good_like_like_proto_rawDescData = protoimpl.X.CompressGZIP(file_npool_good_gw_v1_good_like_like_proto_rawDescData)
	})
	return file_npool_good_gw_v1_good_like_like_proto_rawDescData
}

var file_npool_good_gw_v1_good_like_like_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_npool_good_gw_v1_good_like_like_proto_goTypes = []interface{}{
	(*Like)(nil),               // 0: good.gateway.good1.like.v1.Like
	(*CreateLikeRequest)(nil),  // 1: good.gateway.good1.like.v1.CreateLikeRequest
	(*CreateLikeResponse)(nil), // 2: good.gateway.good1.like.v1.CreateLikeResponse
	(*GetLikesRequest)(nil),    // 3: good.gateway.good1.like.v1.GetLikesRequest
	(*GetLikesResponse)(nil),   // 4: good.gateway.good1.like.v1.GetLikesResponse
	(*GetMyLikesRequest)(nil),  // 5: good.gateway.good1.like.v1.GetMyLikesRequest
	(*GetMyLikesResponse)(nil), // 6: good.gateway.good1.like.v1.GetMyLikesResponse
	(*DeleteLikeRequest)(nil),  // 7: good.gateway.good1.like.v1.DeleteLikeRequest
	(*DeleteLikeResponse)(nil), // 8: good.gateway.good1.like.v1.DeleteLikeResponse
}
var file_npool_good_gw_v1_good_like_like_proto_depIdxs = []int32{
	0, // 0: good.gateway.good1.like.v1.CreateLikeResponse.Info:type_name -> good.gateway.good1.like.v1.Like
	0, // 1: good.gateway.good1.like.v1.GetLikesResponse.Infos:type_name -> good.gateway.good1.like.v1.Like
	0, // 2: good.gateway.good1.like.v1.GetMyLikesResponse.Infos:type_name -> good.gateway.good1.like.v1.Like
	0, // 3: good.gateway.good1.like.v1.DeleteLikeResponse.Info:type_name -> good.gateway.good1.like.v1.Like
	1, // 4: good.gateway.good1.like.v1.Gateway.CreateLike:input_type -> good.gateway.good1.like.v1.CreateLikeRequest
	3, // 5: good.gateway.good1.like.v1.Gateway.GetLikes:input_type -> good.gateway.good1.like.v1.GetLikesRequest
	5, // 6: good.gateway.good1.like.v1.Gateway.GetMyLikes:input_type -> good.gateway.good1.like.v1.GetMyLikesRequest
	7, // 7: good.gateway.good1.like.v1.Gateway.DeleteLike:input_type -> good.gateway.good1.like.v1.DeleteLikeRequest
	2, // 8: good.gateway.good1.like.v1.Gateway.CreateLike:output_type -> good.gateway.good1.like.v1.CreateLikeResponse
	4, // 9: good.gateway.good1.like.v1.Gateway.GetLikes:output_type -> good.gateway.good1.like.v1.GetLikesResponse
	6, // 10: good.gateway.good1.like.v1.Gateway.GetMyLikes:output_type -> good.gateway.good1.like.v1.GetMyLikesResponse
	8, // 11: good.gateway.good1.like.v1.Gateway.DeleteLike:output_type -> good.gateway.good1.like.v1.DeleteLikeResponse
	8, // [8:12] is the sub-list for method output_type
	4, // [4:8] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_npool_good_gw_v1_good_like_like_proto_init() }
func file_npool_good_gw_v1_good_like_like_proto_init() {
	if File_npool_good_gw_v1_good_like_like_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_npool_good_gw_v1_good_like_like_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Like); i {
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
		file_npool_good_gw_v1_good_like_like_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateLikeRequest); i {
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
		file_npool_good_gw_v1_good_like_like_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateLikeResponse); i {
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
		file_npool_good_gw_v1_good_like_like_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetLikesRequest); i {
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
		file_npool_good_gw_v1_good_like_like_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetLikesResponse); i {
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
		file_npool_good_gw_v1_good_like_like_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMyLikesRequest); i {
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
		file_npool_good_gw_v1_good_like_like_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMyLikesResponse); i {
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
		file_npool_good_gw_v1_good_like_like_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteLikeRequest); i {
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
		file_npool_good_gw_v1_good_like_like_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteLikeResponse); i {
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
	file_npool_good_gw_v1_good_like_like_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_npool_good_gw_v1_good_like_like_proto_msgTypes[3].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_npool_good_gw_v1_good_like_like_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_npool_good_gw_v1_good_like_like_proto_goTypes,
		DependencyIndexes: file_npool_good_gw_v1_good_like_like_proto_depIdxs,
		MessageInfos:      file_npool_good_gw_v1_good_like_like_proto_msgTypes,
	}.Build()
	File_npool_good_gw_v1_good_like_like_proto = out.File
	file_npool_good_gw_v1_good_like_like_proto_rawDesc = nil
	file_npool_good_gw_v1_good_like_like_proto_goTypes = nil
	file_npool_good_gw_v1_good_like_like_proto_depIdxs = nil
}
