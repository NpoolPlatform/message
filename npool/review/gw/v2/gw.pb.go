// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.18.1
// source: npool/review/gw/v2/gw.proto

package v2

import (
	_ "github.com/NpoolPlatform/message/npool"
	v2 "github.com/NpoolPlatform/message/npool/review/mgr/v2"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Review struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EmailAddress string               `protobuf:"bytes,10,opt,name=EmailAddress,proto3" json:"EmailAddress,omitempty"`
	PhoneNO      string               `protobuf:"bytes,20,opt,name=PhoneNO,proto3" json:"PhoneNO,omitempty"`
	Reviewer     string               `protobuf:"bytes,30,opt,name=Reviewer,proto3" json:"Reviewer,omitempty"`
	ObjectType   string               `protobuf:"bytes,40,opt,name=ObjectType,proto3" json:"ObjectType,omitempty"`
	ObjectID     string               `protobuf:"bytes,50,opt,name=ObjectID,proto3" json:"ObjectID,omitempty"`
	Domain       string               `protobuf:"bytes,60,opt,name=Domain,proto3" json:"Domain,omitempty"`
	CreatedAt    uint32               `protobuf:"varint,70,opt,name=CreatedAt,proto3" json:"CreatedAt,omitempty"`
	UpdatedAt    uint32               `protobuf:"varint,80,opt,name=UpdatedAt,proto3" json:"UpdatedAt,omitempty"`
	Message      string               `protobuf:"bytes,90,opt,name=Message,proto3" json:"Message,omitempty"`
	State        v2.ReviewState       `protobuf:"varint,100,opt,name=State,proto3,enum=review.manager.v2.ReviewState" json:"State,omitempty"`
	Trigger      v2.ReviewTriggerType `protobuf:"varint,110,opt,name=Trigger,proto3,enum=review.manager.v2.ReviewTriggerType" json:"Trigger,omitempty"`
}

func (x *Review) Reset() {
	*x = Review{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_review_gw_v2_gw_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Review) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Review) ProtoMessage() {}

func (x *Review) ProtoReflect() protoreflect.Message {
	mi := &file_npool_review_gw_v2_gw_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Review.ProtoReflect.Descriptor instead.
func (*Review) Descriptor() ([]byte, []int) {
	return file_npool_review_gw_v2_gw_proto_rawDescGZIP(), []int{0}
}

func (x *Review) GetEmailAddress() string {
	if x != nil {
		return x.EmailAddress
	}
	return ""
}

func (x *Review) GetPhoneNO() string {
	if x != nil {
		return x.PhoneNO
	}
	return ""
}

func (x *Review) GetReviewer() string {
	if x != nil {
		return x.Reviewer
	}
	return ""
}

func (x *Review) GetObjectType() string {
	if x != nil {
		return x.ObjectType
	}
	return ""
}

func (x *Review) GetObjectID() string {
	if x != nil {
		return x.ObjectID
	}
	return ""
}

func (x *Review) GetDomain() string {
	if x != nil {
		return x.Domain
	}
	return ""
}

func (x *Review) GetCreatedAt() uint32 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *Review) GetUpdatedAt() uint32 {
	if x != nil {
		return x.UpdatedAt
	}
	return 0
}

func (x *Review) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *Review) GetState() v2.ReviewState {
	if x != nil {
		return x.State
	}
	return v2.ReviewState(0)
}

func (x *Review) GetTrigger() v2.ReviewTriggerType {
	if x != nil {
		return x.Trigger
	}
	return v2.ReviewTriggerType(0)
}

type UpdateReviewRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID      string         `protobuf:"bytes,10,opt,name=ID,proto3" json:"ID,omitempty"`
	AppID   string         `protobuf:"bytes,20,opt,name=AppID,proto3" json:"AppID,omitempty"`
	UserID  string         `protobuf:"bytes,30,opt,name=UserID,proto3" json:"UserID,omitempty"` // Reviewer
	State   v2.ReviewState `protobuf:"varint,40,opt,name=State,proto3,enum=review.manager.v2.ReviewState" json:"State,omitempty"`
	Message *string        `protobuf:"bytes,50,opt,name=Message,proto3,oneof" json:"Message,omitempty"`
}

func (x *UpdateReviewRequest) Reset() {
	*x = UpdateReviewRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_review_gw_v2_gw_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateReviewRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateReviewRequest) ProtoMessage() {}

func (x *UpdateReviewRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_review_gw_v2_gw_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateReviewRequest.ProtoReflect.Descriptor instead.
func (*UpdateReviewRequest) Descriptor() ([]byte, []int) {
	return file_npool_review_gw_v2_gw_proto_rawDescGZIP(), []int{1}
}

func (x *UpdateReviewRequest) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *UpdateReviewRequest) GetAppID() string {
	if x != nil {
		return x.AppID
	}
	return ""
}

func (x *UpdateReviewRequest) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

func (x *UpdateReviewRequest) GetState() v2.ReviewState {
	if x != nil {
		return x.State
	}
	return v2.ReviewState(0)
}

func (x *UpdateReviewRequest) GetMessage() string {
	if x != nil && x.Message != nil {
		return *x.Message
	}
	return ""
}

type UpdateReviewResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *Review `protobuf:"bytes,10,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (x *UpdateReviewResponse) Reset() {
	*x = UpdateReviewResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_review_gw_v2_gw_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateReviewResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateReviewResponse) ProtoMessage() {}

func (x *UpdateReviewResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_review_gw_v2_gw_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateReviewResponse.ProtoReflect.Descriptor instead.
func (*UpdateReviewResponse) Descriptor() ([]byte, []int) {
	return file_npool_review_gw_v2_gw_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateReviewResponse) GetInfo() *Review {
	if x != nil {
		return x.Info
	}
	return nil
}

type UpdateAppReviewRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID          string         `protobuf:"bytes,10,opt,name=ID,proto3" json:"ID,omitempty"`
	TargetAppID string         `protobuf:"bytes,20,opt,name=TargetAppID,proto3" json:"TargetAppID,omitempty"`
	UserID      string         `protobuf:"bytes,30,opt,name=UserID,proto3" json:"UserID,omitempty"` // Reviewer
	State       v2.ReviewState `protobuf:"varint,40,opt,name=State,proto3,enum=review.manager.v2.ReviewState" json:"State,omitempty"`
	Message     *string        `protobuf:"bytes,50,opt,name=Message,proto3,oneof" json:"Message,omitempty"`
}

func (x *UpdateAppReviewRequest) Reset() {
	*x = UpdateAppReviewRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_review_gw_v2_gw_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateAppReviewRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateAppReviewRequest) ProtoMessage() {}

func (x *UpdateAppReviewRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_review_gw_v2_gw_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateAppReviewRequest.ProtoReflect.Descriptor instead.
func (*UpdateAppReviewRequest) Descriptor() ([]byte, []int) {
	return file_npool_review_gw_v2_gw_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateAppReviewRequest) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *UpdateAppReviewRequest) GetTargetAppID() string {
	if x != nil {
		return x.TargetAppID
	}
	return ""
}

func (x *UpdateAppReviewRequest) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

func (x *UpdateAppReviewRequest) GetState() v2.ReviewState {
	if x != nil {
		return x.State
	}
	return v2.ReviewState(0)
}

func (x *UpdateAppReviewRequest) GetMessage() string {
	if x != nil && x.Message != nil {
		return *x.Message
	}
	return ""
}

type UpdateAppReviewResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *Review `protobuf:"bytes,10,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (x *UpdateAppReviewResponse) Reset() {
	*x = UpdateAppReviewResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_review_gw_v2_gw_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateAppReviewResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateAppReviewResponse) ProtoMessage() {}

func (x *UpdateAppReviewResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_review_gw_v2_gw_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateAppReviewResponse.ProtoReflect.Descriptor instead.
func (*UpdateAppReviewResponse) Descriptor() ([]byte, []int) {
	return file_npool_review_gw_v2_gw_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateAppReviewResponse) GetInfo() *Review {
	if x != nil {
		return x.Info
	}
	return nil
}

type GetObjectTypesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetObjectTypesRequest) Reset() {
	*x = GetObjectTypesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_review_gw_v2_gw_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetObjectTypesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetObjectTypesRequest) ProtoMessage() {}

func (x *GetObjectTypesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_review_gw_v2_gw_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetObjectTypesRequest.ProtoReflect.Descriptor instead.
func (*GetObjectTypesRequest) Descriptor() ([]byte, []int) {
	return file_npool_review_gw_v2_gw_proto_rawDescGZIP(), []int{5}
}

type GetObjectTypesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Infos []string `protobuf:"bytes,10,rep,name=Infos,proto3" json:"Infos,omitempty"`
}

func (x *GetObjectTypesResponse) Reset() {
	*x = GetObjectTypesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_review_gw_v2_gw_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetObjectTypesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetObjectTypesResponse) ProtoMessage() {}

func (x *GetObjectTypesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_review_gw_v2_gw_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetObjectTypesResponse.ProtoReflect.Descriptor instead.
func (*GetObjectTypesResponse) Descriptor() ([]byte, []int) {
	return file_npool_review_gw_v2_gw_proto_rawDescGZIP(), []int{6}
}

func (x *GetObjectTypesResponse) GetInfos() []string {
	if x != nil {
		return x.Infos
	}
	return nil
}

var File_npool_review_gw_v2_gw_proto protoreflect.FileDescriptor

var file_npool_review_gw_v2_gw_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x2f, 0x67,
	0x77, 0x2f, 0x76, 0x32, 0x2f, 0x67, 0x77, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x72,
	0x65, 0x76, 0x69, 0x65, 0x77, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x76, 0x32,
	0x1a, 0x1d, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x2f, 0x6d,
	0x67, 0x72, 0x2f, 0x76, 0x32, 0x2f, 0x6d, 0x67, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65,
	0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x6e, 0x70, 0x6f, 0x6f,
	0x6c, 0x2f, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x82, 0x03,
	0x0a, 0x06, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x12, 0x22, 0x0a, 0x0c, 0x45, 0x6d, 0x61, 0x69,
	0x6c, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c,
	0x45, 0x6d, 0x61, 0x69, 0x6c, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x18, 0x0a, 0x07,
	0x50, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x4f, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x50,
	0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x4f, 0x12, 0x1a, 0x0a, 0x08, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77,
	0x65, 0x72, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77,
	0x65, 0x72, 0x12, 0x1e, 0x0a, 0x0a, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x54, 0x79, 0x70, 0x65,
	0x18, 0x28, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x44, 0x18, 0x32,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x44, 0x12, 0x16,
	0x0a, 0x06, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x18, 0x3c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x18, 0x46, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x18, 0x50, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x5a, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x34, 0x0a, 0x05,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x18, 0x64, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1e, 0x2e, 0x72, 0x65,
	0x76, 0x69, 0x65, 0x77, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x32, 0x2e,
	0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x53, 0x74, 0x61,
	0x74, 0x65, 0x12, 0x3e, 0x0a, 0x07, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x18, 0x6e, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x24, 0x2e, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x2e, 0x6d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x32, 0x2e, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x54, 0x72,
	0x69, 0x67, 0x67, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x52, 0x07, 0x54, 0x72, 0x69, 0x67, 0x67,
	0x65, 0x72, 0x22, 0xb4, 0x01, 0x0a, 0x13, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x76,
	0x69, 0x65, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x41, 0x70,
	0x70, 0x49, 0x44, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x41, 0x70, 0x70, 0x49, 0x44,
	0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x34, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x74,
	0x65, 0x18, 0x28, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1e, 0x2e, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77,
	0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x32, 0x2e, 0x52, 0x65, 0x76, 0x69,
	0x65, 0x77, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x1d,
	0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x32, 0x20, 0x01, 0x28, 0x09, 0x48,
	0x00, 0x52, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x88, 0x01, 0x01, 0x42, 0x0a, 0x0a,
	0x08, 0x5f, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x45, 0x0a, 0x14, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x2d, 0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x19, 0x2e, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79,
	0x2e, 0x76, 0x32, 0x2e, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x52, 0x04, 0x49, 0x6e, 0x66, 0x6f,
	0x22, 0xc3, 0x01, 0x0a, 0x16, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x70, 0x70, 0x52, 0x65,
	0x76, 0x69, 0x65, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49,
	0x44, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x12, 0x20, 0x0a, 0x0b, 0x54,
	0x61, 0x72, 0x67, 0x65, 0x74, 0x41, 0x70, 0x70, 0x49, 0x44, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x41, 0x70, 0x70, 0x49, 0x44, 0x12, 0x16, 0x0a,
	0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x55,
	0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x34, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x74, 0x65, 0x18, 0x28,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x1e, 0x2e, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x2e, 0x6d, 0x61,
	0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x32, 0x2e, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x53,
	0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x1d, 0x0a, 0x07, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x32, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x07,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x88, 0x01, 0x01, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x48, 0x0a, 0x17, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x41, 0x70, 0x70, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x2d, 0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x19, 0x2e, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79,
	0x2e, 0x76, 0x32, 0x2e, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x52, 0x04, 0x49, 0x6e, 0x66, 0x6f,
	0x22, 0x17, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x54, 0x79, 0x70,
	0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x2e, 0x0a, 0x16, 0x47, 0x65, 0x74,
	0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x54, 0x79, 0x70, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x18, 0x0a, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x05, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x32, 0x99, 0x03, 0x0a, 0x07, 0x47, 0x61,
	0x74, 0x65, 0x77, 0x61, 0x79, 0x12, 0x7d, 0x0a, 0x0c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x76, 0x69, 0x65, 0x77, 0x12, 0x26, 0x2e, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x2e, 0x67,
	0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x76, 0x32, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e,
	0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x76,
	0x32, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x16, 0x22, 0x11,
	0x2f, 0x76, 0x32, 0x2f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x72, 0x65, 0x76, 0x69, 0x65,
	0x77, 0x3a, 0x01, 0x2a, 0x12, 0x89, 0x01, 0x0a, 0x0f, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41,
	0x70, 0x70, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x12, 0x29, 0x2e, 0x72, 0x65, 0x76, 0x69, 0x65,
	0x77, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x76, 0x32, 0x2e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x41, 0x70, 0x70, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x2e, 0x67, 0x61, 0x74,
	0x65, 0x77, 0x61, 0x79, 0x2e, 0x76, 0x32, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x70,
	0x70, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x1f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x19, 0x22, 0x14, 0x2f, 0x76, 0x32, 0x2f, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x2f, 0x61, 0x70, 0x70, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x3a, 0x01, 0x2a,
	0x12, 0x82, 0x01, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x54, 0x79,
	0x70, 0x65, 0x73, 0x12, 0x28, 0x2e, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x2e, 0x67, 0x61, 0x74,
	0x65, 0x77, 0x61, 0x79, 0x2e, 0x76, 0x32, 0x2e, 0x47, 0x65, 0x74, 0x4f, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x54, 0x79, 0x70, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e,
	0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x76,
	0x32, 0x2e, 0x47, 0x65, 0x74, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x54, 0x79, 0x70, 0x65, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15,
	0x22, 0x13, 0x2f, 0x76, 0x32, 0x2f, 0x67, 0x65, 0x74, 0x2f, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x74, 0x79, 0x70, 0x65, 0x73, 0x42, 0x35, 0x5a, 0x33, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x4e, 0x70, 0x6f, 0x6f, 0x6c, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72,
	0x6d, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f,
	0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x2f, 0x67, 0x77, 0x2f, 0x76, 0x32, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_npool_review_gw_v2_gw_proto_rawDescOnce sync.Once
	file_npool_review_gw_v2_gw_proto_rawDescData = file_npool_review_gw_v2_gw_proto_rawDesc
)

func file_npool_review_gw_v2_gw_proto_rawDescGZIP() []byte {
	file_npool_review_gw_v2_gw_proto_rawDescOnce.Do(func() {
		file_npool_review_gw_v2_gw_proto_rawDescData = protoimpl.X.CompressGZIP(file_npool_review_gw_v2_gw_proto_rawDescData)
	})
	return file_npool_review_gw_v2_gw_proto_rawDescData
}

var file_npool_review_gw_v2_gw_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_npool_review_gw_v2_gw_proto_goTypes = []interface{}{
	(*Review)(nil),                  // 0: review.gateway.v2.Review
	(*UpdateReviewRequest)(nil),     // 1: review.gateway.v2.UpdateReviewRequest
	(*UpdateReviewResponse)(nil),    // 2: review.gateway.v2.UpdateReviewResponse
	(*UpdateAppReviewRequest)(nil),  // 3: review.gateway.v2.UpdateAppReviewRequest
	(*UpdateAppReviewResponse)(nil), // 4: review.gateway.v2.UpdateAppReviewResponse
	(*GetObjectTypesRequest)(nil),   // 5: review.gateway.v2.GetObjectTypesRequest
	(*GetObjectTypesResponse)(nil),  // 6: review.gateway.v2.GetObjectTypesResponse
	(v2.ReviewState)(0),             // 7: review.manager.v2.ReviewState
	(v2.ReviewTriggerType)(0),       // 8: review.manager.v2.ReviewTriggerType
}
var file_npool_review_gw_v2_gw_proto_depIdxs = []int32{
	7, // 0: review.gateway.v2.Review.State:type_name -> review.manager.v2.ReviewState
	8, // 1: review.gateway.v2.Review.Trigger:type_name -> review.manager.v2.ReviewTriggerType
	7, // 2: review.gateway.v2.UpdateReviewRequest.State:type_name -> review.manager.v2.ReviewState
	0, // 3: review.gateway.v2.UpdateReviewResponse.Info:type_name -> review.gateway.v2.Review
	7, // 4: review.gateway.v2.UpdateAppReviewRequest.State:type_name -> review.manager.v2.ReviewState
	0, // 5: review.gateway.v2.UpdateAppReviewResponse.Info:type_name -> review.gateway.v2.Review
	1, // 6: review.gateway.v2.Gateway.UpdateReview:input_type -> review.gateway.v2.UpdateReviewRequest
	3, // 7: review.gateway.v2.Gateway.UpdateAppReview:input_type -> review.gateway.v2.UpdateAppReviewRequest
	5, // 8: review.gateway.v2.Gateway.GetObjectTypes:input_type -> review.gateway.v2.GetObjectTypesRequest
	2, // 9: review.gateway.v2.Gateway.UpdateReview:output_type -> review.gateway.v2.UpdateReviewResponse
	4, // 10: review.gateway.v2.Gateway.UpdateAppReview:output_type -> review.gateway.v2.UpdateAppReviewResponse
	6, // 11: review.gateway.v2.Gateway.GetObjectTypes:output_type -> review.gateway.v2.GetObjectTypesResponse
	9, // [9:12] is the sub-list for method output_type
	6, // [6:9] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_npool_review_gw_v2_gw_proto_init() }
func file_npool_review_gw_v2_gw_proto_init() {
	if File_npool_review_gw_v2_gw_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_npool_review_gw_v2_gw_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Review); i {
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
		file_npool_review_gw_v2_gw_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateReviewRequest); i {
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
		file_npool_review_gw_v2_gw_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateReviewResponse); i {
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
		file_npool_review_gw_v2_gw_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateAppReviewRequest); i {
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
		file_npool_review_gw_v2_gw_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateAppReviewResponse); i {
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
		file_npool_review_gw_v2_gw_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetObjectTypesRequest); i {
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
		file_npool_review_gw_v2_gw_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetObjectTypesResponse); i {
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
	file_npool_review_gw_v2_gw_proto_msgTypes[1].OneofWrappers = []interface{}{}
	file_npool_review_gw_v2_gw_proto_msgTypes[3].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_npool_review_gw_v2_gw_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_npool_review_gw_v2_gw_proto_goTypes,
		DependencyIndexes: file_npool_review_gw_v2_gw_proto_depIdxs,
		MessageInfos:      file_npool_review_gw_v2_gw_proto_msgTypes,
	}.Build()
	File_npool_review_gw_v2_gw_proto = out.File
	file_npool_review_gw_v2_gw_proto_rawDesc = nil
	file_npool_review_gw_v2_gw_proto_goTypes = nil
	file_npool_review_gw_v2_gw_proto_depIdxs = nil
}
