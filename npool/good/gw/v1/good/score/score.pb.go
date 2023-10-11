// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.18.1
// source: npool/good/gw/v1/good/score/score.proto

package score

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

type Score struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID           string  `protobuf:"bytes,10,opt,name=ID,proto3" json:"ID,omitempty"`
	AppID        string  `protobuf:"bytes,20,opt,name=AppID,proto3" json:"AppID,omitempty"`
	AppName      string  `protobuf:"bytes,30,opt,name=AppName,proto3" json:"AppName,omitempty"`
	UserID       string  `protobuf:"bytes,40,opt,name=UserID,proto3" json:"UserID,omitempty"`
	Username     *string `protobuf:"bytes,50,opt,name=Username,proto3,oneof" json:"Username,omitempty"`
	EmailAddress *string `protobuf:"bytes,60,opt,name=EmailAddress,proto3,oneof" json:"EmailAddress,omitempty"`
	PhoneNO      *string `protobuf:"bytes,70,opt,name=PhoneNO,proto3,oneof" json:"PhoneNO,omitempty"`
	AppGoodID    string  `protobuf:"bytes,80,opt,name=AppGoodID,proto3" json:"AppGoodID,omitempty"`
	GoodName     string  `protobuf:"bytes,90,opt,name=GoodName,proto3" json:"GoodName,omitempty"`
	Score        string  `protobuf:"bytes,100,opt,name=Score,proto3" json:"Score,omitempty"`
	GoodID       string  `protobuf:"bytes,110,opt,name=GoodID,proto3" json:"GoodID,omitempty"`
	CreatedAt    uint32  `protobuf:"varint,1000,opt,name=CreatedAt,proto3" json:"CreatedAt,omitempty"`
}

func (x *Score) Reset() {
	*x = Score{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_good_gw_v1_good_score_score_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Score) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Score) ProtoMessage() {}

func (x *Score) ProtoReflect() protoreflect.Message {
	mi := &file_npool_good_gw_v1_good_score_score_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Score.ProtoReflect.Descriptor instead.
func (*Score) Descriptor() ([]byte, []int) {
	return file_npool_good_gw_v1_good_score_score_proto_rawDescGZIP(), []int{0}
}

func (x *Score) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *Score) GetAppID() string {
	if x != nil {
		return x.AppID
	}
	return ""
}

func (x *Score) GetAppName() string {
	if x != nil {
		return x.AppName
	}
	return ""
}

func (x *Score) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

func (x *Score) GetUsername() string {
	if x != nil && x.Username != nil {
		return *x.Username
	}
	return ""
}

func (x *Score) GetEmailAddress() string {
	if x != nil && x.EmailAddress != nil {
		return *x.EmailAddress
	}
	return ""
}

func (x *Score) GetPhoneNO() string {
	if x != nil && x.PhoneNO != nil {
		return *x.PhoneNO
	}
	return ""
}

func (x *Score) GetAppGoodID() string {
	if x != nil {
		return x.AppGoodID
	}
	return ""
}

func (x *Score) GetGoodName() string {
	if x != nil {
		return x.GoodName
	}
	return ""
}

func (x *Score) GetScore() string {
	if x != nil {
		return x.Score
	}
	return ""
}

func (x *Score) GetGoodID() string {
	if x != nil {
		return x.GoodID
	}
	return ""
}

func (x *Score) GetCreatedAt() uint32 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

type CreateScoreRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AppID     string `protobuf:"bytes,10,opt,name=AppID,proto3" json:"AppID,omitempty"`
	UserID    string `protobuf:"bytes,20,opt,name=UserID,proto3" json:"UserID,omitempty"`
	AppGoodID string `protobuf:"bytes,30,opt,name=AppGoodID,proto3" json:"AppGoodID,omitempty"`
	Score     string `protobuf:"bytes,40,opt,name=Score,proto3" json:"Score,omitempty"`
}

func (x *CreateScoreRequest) Reset() {
	*x = CreateScoreRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_good_gw_v1_good_score_score_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateScoreRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateScoreRequest) ProtoMessage() {}

func (x *CreateScoreRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_good_gw_v1_good_score_score_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateScoreRequest.ProtoReflect.Descriptor instead.
func (*CreateScoreRequest) Descriptor() ([]byte, []int) {
	return file_npool_good_gw_v1_good_score_score_proto_rawDescGZIP(), []int{1}
}

func (x *CreateScoreRequest) GetAppID() string {
	if x != nil {
		return x.AppID
	}
	return ""
}

func (x *CreateScoreRequest) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

func (x *CreateScoreRequest) GetAppGoodID() string {
	if x != nil {
		return x.AppGoodID
	}
	return ""
}

func (x *CreateScoreRequest) GetScore() string {
	if x != nil {
		return x.Score
	}
	return ""
}

type CreateScoreResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *Score `protobuf:"bytes,10,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (x *CreateScoreResponse) Reset() {
	*x = CreateScoreResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_good_gw_v1_good_score_score_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateScoreResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateScoreResponse) ProtoMessage() {}

func (x *CreateScoreResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_good_gw_v1_good_score_score_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateScoreResponse.ProtoReflect.Descriptor instead.
func (*CreateScoreResponse) Descriptor() ([]byte, []int) {
	return file_npool_good_gw_v1_good_score_score_proto_rawDescGZIP(), []int{2}
}

func (x *CreateScoreResponse) GetInfo() *Score {
	if x != nil {
		return x.Info
	}
	return nil
}

type GetScoresRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AppID     string  `protobuf:"bytes,10,opt,name=AppID,proto3" json:"AppID,omitempty"`
	AppGoodID *string `protobuf:"bytes,20,opt,name=AppGoodID,proto3,oneof" json:"AppGoodID,omitempty"`
	Offset    int32   `protobuf:"varint,30,opt,name=Offset,proto3" json:"Offset,omitempty"`
	Limit     int32   `protobuf:"varint,40,opt,name=Limit,proto3" json:"Limit,omitempty"`
}

func (x *GetScoresRequest) Reset() {
	*x = GetScoresRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_good_gw_v1_good_score_score_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetScoresRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetScoresRequest) ProtoMessage() {}

func (x *GetScoresRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_good_gw_v1_good_score_score_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetScoresRequest.ProtoReflect.Descriptor instead.
func (*GetScoresRequest) Descriptor() ([]byte, []int) {
	return file_npool_good_gw_v1_good_score_score_proto_rawDescGZIP(), []int{3}
}

func (x *GetScoresRequest) GetAppID() string {
	if x != nil {
		return x.AppID
	}
	return ""
}

func (x *GetScoresRequest) GetAppGoodID() string {
	if x != nil && x.AppGoodID != nil {
		return *x.AppGoodID
	}
	return ""
}

func (x *GetScoresRequest) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *GetScoresRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type GetScoresResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Infos []*Score `protobuf:"bytes,10,rep,name=Infos,proto3" json:"Infos,omitempty"`
	Total uint32   `protobuf:"varint,20,opt,name=Total,proto3" json:"Total,omitempty"`
}

func (x *GetScoresResponse) Reset() {
	*x = GetScoresResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_good_gw_v1_good_score_score_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetScoresResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetScoresResponse) ProtoMessage() {}

func (x *GetScoresResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_good_gw_v1_good_score_score_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetScoresResponse.ProtoReflect.Descriptor instead.
func (*GetScoresResponse) Descriptor() ([]byte, []int) {
	return file_npool_good_gw_v1_good_score_score_proto_rawDescGZIP(), []int{4}
}

func (x *GetScoresResponse) GetInfos() []*Score {
	if x != nil {
		return x.Infos
	}
	return nil
}

func (x *GetScoresResponse) GetTotal() uint32 {
	if x != nil {
		return x.Total
	}
	return 0
}

type GetMyScoresRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AppID  string `protobuf:"bytes,10,opt,name=AppID,proto3" json:"AppID,omitempty"`
	UserID string `protobuf:"bytes,20,opt,name=UserID,proto3" json:"UserID,omitempty"`
	Offset int32  `protobuf:"varint,30,opt,name=Offset,proto3" json:"Offset,omitempty"`
	Limit  int32  `protobuf:"varint,40,opt,name=Limit,proto3" json:"Limit,omitempty"`
}

func (x *GetMyScoresRequest) Reset() {
	*x = GetMyScoresRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_good_gw_v1_good_score_score_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMyScoresRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMyScoresRequest) ProtoMessage() {}

func (x *GetMyScoresRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_good_gw_v1_good_score_score_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMyScoresRequest.ProtoReflect.Descriptor instead.
func (*GetMyScoresRequest) Descriptor() ([]byte, []int) {
	return file_npool_good_gw_v1_good_score_score_proto_rawDescGZIP(), []int{5}
}

func (x *GetMyScoresRequest) GetAppID() string {
	if x != nil {
		return x.AppID
	}
	return ""
}

func (x *GetMyScoresRequest) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

func (x *GetMyScoresRequest) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *GetMyScoresRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type GetMyScoresResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Infos []*Score `protobuf:"bytes,10,rep,name=Infos,proto3" json:"Infos,omitempty"`
	Total uint32   `protobuf:"varint,20,opt,name=Total,proto3" json:"Total,omitempty"`
}

func (x *GetMyScoresResponse) Reset() {
	*x = GetMyScoresResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_good_gw_v1_good_score_score_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMyScoresResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMyScoresResponse) ProtoMessage() {}

func (x *GetMyScoresResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_good_gw_v1_good_score_score_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMyScoresResponse.ProtoReflect.Descriptor instead.
func (*GetMyScoresResponse) Descriptor() ([]byte, []int) {
	return file_npool_good_gw_v1_good_score_score_proto_rawDescGZIP(), []int{6}
}

func (x *GetMyScoresResponse) GetInfos() []*Score {
	if x != nil {
		return x.Infos
	}
	return nil
}

func (x *GetMyScoresResponse) GetTotal() uint32 {
	if x != nil {
		return x.Total
	}
	return 0
}

type DeleteScoreRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID     string `protobuf:"bytes,10,opt,name=ID,proto3" json:"ID,omitempty"`
	AppID  string `protobuf:"bytes,20,opt,name=AppID,proto3" json:"AppID,omitempty"`
	UserID string `protobuf:"bytes,30,opt,name=UserID,proto3" json:"UserID,omitempty"`
}

func (x *DeleteScoreRequest) Reset() {
	*x = DeleteScoreRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_good_gw_v1_good_score_score_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteScoreRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteScoreRequest) ProtoMessage() {}

func (x *DeleteScoreRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_good_gw_v1_good_score_score_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteScoreRequest.ProtoReflect.Descriptor instead.
func (*DeleteScoreRequest) Descriptor() ([]byte, []int) {
	return file_npool_good_gw_v1_good_score_score_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteScoreRequest) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *DeleteScoreRequest) GetAppID() string {
	if x != nil {
		return x.AppID
	}
	return ""
}

func (x *DeleteScoreRequest) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

type DeleteScoreResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *Score `protobuf:"bytes,10,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (x *DeleteScoreResponse) Reset() {
	*x = DeleteScoreResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_good_gw_v1_good_score_score_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteScoreResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteScoreResponse) ProtoMessage() {}

func (x *DeleteScoreResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_good_gw_v1_good_score_score_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteScoreResponse.ProtoReflect.Descriptor instead.
func (*DeleteScoreResponse) Descriptor() ([]byte, []int) {
	return file_npool_good_gw_v1_good_score_score_proto_rawDescGZIP(), []int{8}
}

func (x *DeleteScoreResponse) GetInfo() *Score {
	if x != nil {
		return x.Info
	}
	return nil
}

var File_npool_good_gw_v1_good_score_score_proto protoreflect.FileDescriptor

var file_npool_good_gw_v1_good_score_score_proto_rawDesc = []byte{
	0x0a, 0x27, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x67, 0x6f, 0x6f, 0x64, 0x2f, 0x67, 0x77, 0x2f,
	0x76, 0x31, 0x2f, 0x67, 0x6f, 0x6f, 0x64, 0x2f, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x73, 0x63,
	0x6f, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1b, 0x67, 0x6f, 0x6f, 0x64, 0x2e,
	0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x67, 0x6f, 0x6f, 0x64, 0x31, 0x2e, 0x73, 0x63,
	0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe1, 0x02, 0x0a, 0x05, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x0e,
	0x0a, 0x02, 0x49, 0x44, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x12, 0x14,
	0x0a, 0x05, 0x41, 0x70, 0x70, 0x49, 0x44, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x41,
	0x70, 0x70, 0x49, 0x44, 0x12, 0x18, 0x0a, 0x07, 0x41, 0x70, 0x70, 0x4e, 0x61, 0x6d, 0x65, 0x18,
	0x1e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x41, 0x70, 0x70, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x28, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x1f, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x32, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x08, 0x55, 0x73, 0x65, 0x72,
	0x6e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x27, 0x0a, 0x0c, 0x45, 0x6d, 0x61, 0x69, 0x6c,
	0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x3c, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52,
	0x0c, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x88, 0x01, 0x01,
	0x12, 0x1d, 0x0a, 0x07, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x4f, 0x18, 0x46, 0x20, 0x01, 0x28,
	0x09, 0x48, 0x02, 0x52, 0x07, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x4f, 0x88, 0x01, 0x01, 0x12,
	0x1c, 0x0a, 0x09, 0x41, 0x70, 0x70, 0x47, 0x6f, 0x6f, 0x64, 0x49, 0x44, 0x18, 0x50, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x41, 0x70, 0x70, 0x47, 0x6f, 0x6f, 0x64, 0x49, 0x44, 0x12, 0x1a, 0x0a,
	0x08, 0x47, 0x6f, 0x6f, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x5a, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x47, 0x6f, 0x6f, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x53, 0x63, 0x6f,
	0x72, 0x65, 0x18, 0x64, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x12,
	0x1d, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0xe8, 0x07, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x42, 0x0b,
	0x0a, 0x09, 0x5f, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x0f, 0x0a, 0x0d, 0x5f,
	0x45, 0x6d, 0x61, 0x69, 0x6c, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x42, 0x0a, 0x0a, 0x08,
	0x5f, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x4f, 0x22, 0x76, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14,
	0x0a, 0x05, 0x41, 0x70, 0x70, 0x49, 0x44, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x41,
	0x70, 0x70, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x14,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x1c, 0x0a, 0x09,
	0x41, 0x70, 0x70, 0x47, 0x6f, 0x6f, 0x64, 0x49, 0x44, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x41, 0x70, 0x70, 0x47, 0x6f, 0x6f, 0x64, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x53, 0x63,
	0x6f, 0x72, 0x65, 0x18, 0x28, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x53, 0x63, 0x6f, 0x72, 0x65,
	0x22, 0x4d, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x36, 0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x67, 0x6f, 0x6f, 0x64, 0x2e, 0x67, 0x61, 0x74,
	0x65, 0x77, 0x61, 0x79, 0x2e, 0x67, 0x6f, 0x6f, 0x64, 0x31, 0x2e, 0x73, 0x63, 0x6f, 0x72, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x52, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x22,
	0x87, 0x01, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x41, 0x70, 0x70, 0x49, 0x44, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x41, 0x70, 0x70, 0x49, 0x44, 0x12, 0x21, 0x0a, 0x09, 0x41, 0x70,
	0x70, 0x47, 0x6f, 0x6f, 0x64, 0x49, 0x44, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52,
	0x09, 0x41, 0x70, 0x70, 0x47, 0x6f, 0x6f, 0x64, 0x49, 0x44, 0x88, 0x01, 0x01, 0x12, 0x16, 0x0a,
	0x06, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x4f,
	0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x28,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x42, 0x0c, 0x0a, 0x0a, 0x5f,
	0x41, 0x70, 0x70, 0x47, 0x6f, 0x6f, 0x64, 0x49, 0x44, 0x22, 0x63, 0x0a, 0x11, 0x47, 0x65, 0x74,
	0x53, 0x63, 0x6f, 0x72, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x38,
	0x0a, 0x05, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e,
	0x67, 0x6f, 0x6f, 0x64, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x67, 0x6f, 0x6f,
	0x64, 0x31, 0x2e, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x63, 0x6f, 0x72,
	0x65, 0x52, 0x05, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x6f, 0x74, 0x61,
	0x6c, 0x18, 0x14, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x22, 0x70,
	0x0a, 0x12, 0x47, 0x65, 0x74, 0x4d, 0x79, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x41, 0x70, 0x70, 0x49, 0x44, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x41, 0x70, 0x70, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73,
	0x65, 0x72, 0x49, 0x44, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72,
	0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x1e, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x06, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x4c, 0x69,
	0x6d, 0x69, 0x74, 0x18, 0x28, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x4c, 0x69, 0x6d, 0x69, 0x74,
	0x22, 0x65, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x4d, 0x79, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x38, 0x0a, 0x05, 0x49, 0x6e, 0x66, 0x6f, 0x73,
	0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x67, 0x6f, 0x6f, 0x64, 0x2e, 0x67, 0x61,
	0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x67, 0x6f, 0x6f, 0x64, 0x31, 0x2e, 0x73, 0x63, 0x6f, 0x72,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x52, 0x05, 0x49, 0x6e, 0x66, 0x6f,
	0x73, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x14, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x22, 0x52, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x49, 0x44, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x12, 0x14, 0x0a,
	0x05, 0x41, 0x70, 0x70, 0x49, 0x44, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x41, 0x70,
	0x70, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x1e, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x22, 0x4d, 0x0a, 0x13, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x36, 0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x22, 0x2e, 0x67, 0x6f, 0x6f, 0x64, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e,
	0x67, 0x6f, 0x6f, 0x64, 0x31, 0x2e, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53,
	0x63, 0x6f, 0x72, 0x65, 0x52, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x32, 0xd1, 0x04, 0x0a, 0x07, 0x47,
	0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x12, 0x91, 0x01, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x2f, 0x2e, 0x67, 0x6f, 0x6f, 0x64, 0x2e, 0x67, 0x61,
	0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x67, 0x6f, 0x6f, 0x64, 0x31, 0x2e, 0x73, 0x63, 0x6f, 0x72,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x63, 0x6f, 0x72, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x30, 0x2e, 0x67, 0x6f, 0x6f, 0x64, 0x2e, 0x67,
	0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x67, 0x6f, 0x6f, 0x64, 0x31, 0x2e, 0x73, 0x63, 0x6f,
	0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x63, 0x6f, 0x72,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1f, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x19, 0x3a, 0x01, 0x2a, 0x22, 0x14, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x2f, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x89, 0x01, 0x0a, 0x09, 0x47,
	0x65, 0x74, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x73, 0x12, 0x2d, 0x2e, 0x67, 0x6f, 0x6f, 0x64, 0x2e,
	0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x67, 0x6f, 0x6f, 0x64, 0x31, 0x2e, 0x73, 0x63,
	0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2e, 0x2e, 0x67, 0x6f, 0x6f, 0x64, 0x2e, 0x67,
	0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x67, 0x6f, 0x6f, 0x64, 0x31, 0x2e, 0x73, 0x63, 0x6f,
	0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x17, 0x3a,
	0x01, 0x2a, 0x22, 0x12, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x65, 0x74, 0x2f, 0x67, 0x6f, 0x6f, 0x64,
	0x73, 0x63, 0x6f, 0x72, 0x65, 0x73, 0x12, 0x91, 0x01, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x4d, 0x79,
	0x53, 0x63, 0x6f, 0x72, 0x65, 0x73, 0x12, 0x2f, 0x2e, 0x67, 0x6f, 0x6f, 0x64, 0x2e, 0x67, 0x61,
	0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x67, 0x6f, 0x6f, 0x64, 0x31, 0x2e, 0x73, 0x63, 0x6f, 0x72,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x79, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x30, 0x2e, 0x67, 0x6f, 0x6f, 0x64, 0x2e, 0x67,
	0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x67, 0x6f, 0x6f, 0x64, 0x31, 0x2e, 0x73, 0x63, 0x6f,
	0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x79, 0x53, 0x63, 0x6f, 0x72, 0x65,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1f, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x19, 0x3a, 0x01, 0x2a, 0x22, 0x14, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x65, 0x74, 0x2f, 0x6d, 0x79,
	0x67, 0x6f, 0x6f, 0x64, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x73, 0x12, 0x91, 0x01, 0x0a, 0x0b, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x2f, 0x2e, 0x67, 0x6f, 0x6f,
	0x64, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x67, 0x6f, 0x6f, 0x64, 0x31, 0x2e,
	0x73, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53,
	0x63, 0x6f, 0x72, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x30, 0x2e, 0x67, 0x6f,
	0x6f, 0x64, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x67, 0x6f, 0x6f, 0x64, 0x31,
	0x2e, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x53, 0x63, 0x6f, 0x72, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1f, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x19, 0x3a, 0x01, 0x2a, 0x22, 0x14, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x2f, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x42, 0x3e,
	0x5a, 0x3c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4e, 0x70, 0x6f,
	0x6f, 0x6c, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x2f, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x67, 0x6f, 0x6f, 0x64, 0x2f, 0x67, 0x77,
	0x2f, 0x76, 0x31, 0x2f, 0x67, 0x6f, 0x6f, 0x64, 0x2f, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_npool_good_gw_v1_good_score_score_proto_rawDescOnce sync.Once
	file_npool_good_gw_v1_good_score_score_proto_rawDescData = file_npool_good_gw_v1_good_score_score_proto_rawDesc
)

func file_npool_good_gw_v1_good_score_score_proto_rawDescGZIP() []byte {
	file_npool_good_gw_v1_good_score_score_proto_rawDescOnce.Do(func() {
		file_npool_good_gw_v1_good_score_score_proto_rawDescData = protoimpl.X.CompressGZIP(file_npool_good_gw_v1_good_score_score_proto_rawDescData)
	})
	return file_npool_good_gw_v1_good_score_score_proto_rawDescData
}

var file_npool_good_gw_v1_good_score_score_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_npool_good_gw_v1_good_score_score_proto_goTypes = []interface{}{
	(*Score)(nil),               // 0: good.gateway.good1.score.v1.Score
	(*CreateScoreRequest)(nil),  // 1: good.gateway.good1.score.v1.CreateScoreRequest
	(*CreateScoreResponse)(nil), // 2: good.gateway.good1.score.v1.CreateScoreResponse
	(*GetScoresRequest)(nil),    // 3: good.gateway.good1.score.v1.GetScoresRequest
	(*GetScoresResponse)(nil),   // 4: good.gateway.good1.score.v1.GetScoresResponse
	(*GetMyScoresRequest)(nil),  // 5: good.gateway.good1.score.v1.GetMyScoresRequest
	(*GetMyScoresResponse)(nil), // 6: good.gateway.good1.score.v1.GetMyScoresResponse
	(*DeleteScoreRequest)(nil),  // 7: good.gateway.good1.score.v1.DeleteScoreRequest
	(*DeleteScoreResponse)(nil), // 8: good.gateway.good1.score.v1.DeleteScoreResponse
}
var file_npool_good_gw_v1_good_score_score_proto_depIdxs = []int32{
	0, // 0: good.gateway.good1.score.v1.CreateScoreResponse.Info:type_name -> good.gateway.good1.score.v1.Score
	0, // 1: good.gateway.good1.score.v1.GetScoresResponse.Infos:type_name -> good.gateway.good1.score.v1.Score
	0, // 2: good.gateway.good1.score.v1.GetMyScoresResponse.Infos:type_name -> good.gateway.good1.score.v1.Score
	0, // 3: good.gateway.good1.score.v1.DeleteScoreResponse.Info:type_name -> good.gateway.good1.score.v1.Score
	1, // 4: good.gateway.good1.score.v1.Gateway.CreateScore:input_type -> good.gateway.good1.score.v1.CreateScoreRequest
	3, // 5: good.gateway.good1.score.v1.Gateway.GetScores:input_type -> good.gateway.good1.score.v1.GetScoresRequest
	5, // 6: good.gateway.good1.score.v1.Gateway.GetMyScores:input_type -> good.gateway.good1.score.v1.GetMyScoresRequest
	7, // 7: good.gateway.good1.score.v1.Gateway.DeleteScore:input_type -> good.gateway.good1.score.v1.DeleteScoreRequest
	2, // 8: good.gateway.good1.score.v1.Gateway.CreateScore:output_type -> good.gateway.good1.score.v1.CreateScoreResponse
	4, // 9: good.gateway.good1.score.v1.Gateway.GetScores:output_type -> good.gateway.good1.score.v1.GetScoresResponse
	6, // 10: good.gateway.good1.score.v1.Gateway.GetMyScores:output_type -> good.gateway.good1.score.v1.GetMyScoresResponse
	8, // 11: good.gateway.good1.score.v1.Gateway.DeleteScore:output_type -> good.gateway.good1.score.v1.DeleteScoreResponse
	8, // [8:12] is the sub-list for method output_type
	4, // [4:8] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_npool_good_gw_v1_good_score_score_proto_init() }
func file_npool_good_gw_v1_good_score_score_proto_init() {
	if File_npool_good_gw_v1_good_score_score_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_npool_good_gw_v1_good_score_score_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Score); i {
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
		file_npool_good_gw_v1_good_score_score_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateScoreRequest); i {
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
		file_npool_good_gw_v1_good_score_score_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateScoreResponse); i {
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
		file_npool_good_gw_v1_good_score_score_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetScoresRequest); i {
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
		file_npool_good_gw_v1_good_score_score_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetScoresResponse); i {
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
		file_npool_good_gw_v1_good_score_score_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMyScoresRequest); i {
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
		file_npool_good_gw_v1_good_score_score_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMyScoresResponse); i {
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
		file_npool_good_gw_v1_good_score_score_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteScoreRequest); i {
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
		file_npool_good_gw_v1_good_score_score_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteScoreResponse); i {
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
	file_npool_good_gw_v1_good_score_score_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_npool_good_gw_v1_good_score_score_proto_msgTypes[3].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_npool_good_gw_v1_good_score_score_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_npool_good_gw_v1_good_score_score_proto_goTypes,
		DependencyIndexes: file_npool_good_gw_v1_good_score_score_proto_depIdxs,
		MessageInfos:      file_npool_good_gw_v1_good_score_score_proto_msgTypes,
	}.Build()
	File_npool_good_gw_v1_good_score_score_proto = out.File
	file_npool_good_gw_v1_good_score_score_proto_rawDesc = nil
	file_npool_good_gw_v1_good_score_score_proto_goTypes = nil
	file_npool_good_gw_v1_good_score_score_proto_depIdxs = nil
}
