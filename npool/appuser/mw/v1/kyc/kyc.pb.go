// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.18.1
// source: npool/appuser/mw/v1/kyc/kyc.proto

package kyc

import (
	_ "github.com/NpoolPlatform/message/npool"
	kyc "github.com/NpoolPlatform/message/npool/appuser/mgr/v2/kyc"
	_ "github.com/NpoolPlatform/message/npool/review/mgr/v2"
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

type Kyc struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: sql:"id"
	ID string `protobuf:"bytes,10,opt,name=ID,proto3" json:"ID,omitempty" sql:"id"`
	// @inject_tag: sql:"app_id"
	AppID string `protobuf:"bytes,20,opt,name=AppID,proto3" json:"AppID,omitempty" sql:"app_id"`
	// @inject_tag: sql:"app_name"
	AppName string `protobuf:"bytes,30,opt,name=AppName,proto3" json:"AppName,omitempty" sql:"app_name"`
	// @inject_tag: sql:"app_logo"
	AppLogo string `protobuf:"bytes,40,opt,name=AppLogo,proto3" json:"AppLogo,omitempty" sql:"app_logo"`
	// @inject_tag: sql:"user_id"
	UserID string `protobuf:"bytes,50,opt,name=UserID,proto3" json:"UserID,omitempty" sql:"user_id"`
	// @inject_tag: sql:"email_address"
	EmailAddress string `protobuf:"bytes,60,opt,name=EmailAddress,proto3" json:"EmailAddress,omitempty" sql:"email_address"`
	// @inject_tag: sql:"phone_no"
	PhoneNO string `protobuf:"bytes,70,opt,name=PhoneNO,proto3" json:"PhoneNO,omitempty" sql:"phone_no"`
	// @inject_tag: sql:"document_type"
	DocumentTypeStr string              `protobuf:"bytes,80,opt,name=DocumentTypeStr,proto3" json:"DocumentTypeStr,omitempty" sql:"document_type"`
	DocumentType    kyc.KycDocumentType `protobuf:"varint,90,opt,name=DocumentType,proto3,enum=appuser.manager.kyc.v2.KycDocumentType" json:"DocumentType,omitempty"`
	// @inject_tag: sql:"id_number"
	IDNumber string `protobuf:"bytes,100,opt,name=IDNumber,proto3" json:"IDNumber,omitempty" sql:"id_number"`
	// @inject_tag: sql:"front_img"
	FrontImg string `protobuf:"bytes,110,opt,name=FrontImg,proto3" json:"FrontImg,omitempty" sql:"front_img"`
	// @inject_tag: sql:"back_img"
	BackImg string `protobuf:"bytes,120,opt,name=BackImg,proto3" json:"BackImg,omitempty" sql:"back_img"`
	// @inject_tag: sql:"selfie_img"
	SelfieImg string `protobuf:"bytes,130,opt,name=SelfieImg,proto3" json:"SelfieImg,omitempty" sql:"selfie_img"`
	// @inject_tag: sql:"entity_type"
	EntityTypeStr string            `protobuf:"bytes,140,opt,name=EntityTypeStr,proto3" json:"EntityTypeStr,omitempty" sql:"entity_type"`
	EntityType    kyc.KycEntityType `protobuf:"varint,150,opt,name=EntityType,proto3,enum=appuser.manager.kyc.v2.KycEntityType" json:"EntityType,omitempty"`
	// @inject_tag: sql:"review_id"
	ReviewID string `protobuf:"bytes,160,opt,name=ReviewID,proto3" json:"ReviewID,omitempty" sql:"review_id"`
	// @inject_tag: sql:"review_state"
	ReviewStateStr string       `protobuf:"bytes,170,opt,name=ReviewStateStr,proto3" json:"ReviewStateStr,omitempty" sql:"review_state"`
	State          kyc.KycState `protobuf:"varint,180,opt,name=State,proto3,enum=appuser.manager.kyc.v2.KycState" json:"State,omitempty"`
	// @inject_tag: sql:"created_at"
	CreatedAt uint32 `protobuf:"varint,190,opt,name=CreatedAt,proto3" json:"CreatedAt,omitempty" sql:"created_at"`
	// @inject_tag: sql:"updated_at"
	UpdatedAt uint32 `protobuf:"varint,200,opt,name=UpdatedAt,proto3" json:"UpdatedAt,omitempty" sql:"updated_at"`
}

func (x *Kyc) Reset() {
	*x = Kyc{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_appuser_mw_v1_kyc_kyc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Kyc) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Kyc) ProtoMessage() {}

func (x *Kyc) ProtoReflect() protoreflect.Message {
	mi := &file_npool_appuser_mw_v1_kyc_kyc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Kyc.ProtoReflect.Descriptor instead.
func (*Kyc) Descriptor() ([]byte, []int) {
	return file_npool_appuser_mw_v1_kyc_kyc_proto_rawDescGZIP(), []int{0}
}

func (x *Kyc) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *Kyc) GetAppID() string {
	if x != nil {
		return x.AppID
	}
	return ""
}

func (x *Kyc) GetAppName() string {
	if x != nil {
		return x.AppName
	}
	return ""
}

func (x *Kyc) GetAppLogo() string {
	if x != nil {
		return x.AppLogo
	}
	return ""
}

func (x *Kyc) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

func (x *Kyc) GetEmailAddress() string {
	if x != nil {
		return x.EmailAddress
	}
	return ""
}

func (x *Kyc) GetPhoneNO() string {
	if x != nil {
		return x.PhoneNO
	}
	return ""
}

func (x *Kyc) GetDocumentTypeStr() string {
	if x != nil {
		return x.DocumentTypeStr
	}
	return ""
}

func (x *Kyc) GetDocumentType() kyc.KycDocumentType {
	if x != nil {
		return x.DocumentType
	}
	return kyc.KycDocumentType(0)
}

func (x *Kyc) GetIDNumber() string {
	if x != nil {
		return x.IDNumber
	}
	return ""
}

func (x *Kyc) GetFrontImg() string {
	if x != nil {
		return x.FrontImg
	}
	return ""
}

func (x *Kyc) GetBackImg() string {
	if x != nil {
		return x.BackImg
	}
	return ""
}

func (x *Kyc) GetSelfieImg() string {
	if x != nil {
		return x.SelfieImg
	}
	return ""
}

func (x *Kyc) GetEntityTypeStr() string {
	if x != nil {
		return x.EntityTypeStr
	}
	return ""
}

func (x *Kyc) GetEntityType() kyc.KycEntityType {
	if x != nil {
		return x.EntityType
	}
	return kyc.KycEntityType(0)
}

func (x *Kyc) GetReviewID() string {
	if x != nil {
		return x.ReviewID
	}
	return ""
}

func (x *Kyc) GetReviewStateStr() string {
	if x != nil {
		return x.ReviewStateStr
	}
	return ""
}

func (x *Kyc) GetState() kyc.KycState {
	if x != nil {
		return x.State
	}
	return kyc.KycState(0)
}

func (x *Kyc) GetCreatedAt() uint32 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *Kyc) GetUpdatedAt() uint32 {
	if x != nil {
		return x.UpdatedAt
	}
	return 0
}

type Conds struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Conds *kyc.Conds `protobuf:"bytes,10,opt,name=Conds,proto3" json:"Conds,omitempty"`
}

func (x *Conds) Reset() {
	*x = Conds{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_appuser_mw_v1_kyc_kyc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Conds) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Conds) ProtoMessage() {}

func (x *Conds) ProtoReflect() protoreflect.Message {
	mi := &file_npool_appuser_mw_v1_kyc_kyc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Conds.ProtoReflect.Descriptor instead.
func (*Conds) Descriptor() ([]byte, []int) {
	return file_npool_appuser_mw_v1_kyc_kyc_proto_rawDescGZIP(), []int{1}
}

func (x *Conds) GetConds() *kyc.Conds {
	if x != nil {
		return x.Conds
	}
	return nil
}

type GetKycRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID string `protobuf:"bytes,10,opt,name=ID,proto3" json:"ID,omitempty"`
}

func (x *GetKycRequest) Reset() {
	*x = GetKycRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_appuser_mw_v1_kyc_kyc_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetKycRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetKycRequest) ProtoMessage() {}

func (x *GetKycRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_appuser_mw_v1_kyc_kyc_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetKycRequest.ProtoReflect.Descriptor instead.
func (*GetKycRequest) Descriptor() ([]byte, []int) {
	return file_npool_appuser_mw_v1_kyc_kyc_proto_rawDescGZIP(), []int{2}
}

func (x *GetKycRequest) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

type GetKycResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *Kyc `protobuf:"bytes,10,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (x *GetKycResponse) Reset() {
	*x = GetKycResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_appuser_mw_v1_kyc_kyc_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetKycResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetKycResponse) ProtoMessage() {}

func (x *GetKycResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_appuser_mw_v1_kyc_kyc_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetKycResponse.ProtoReflect.Descriptor instead.
func (*GetKycResponse) Descriptor() ([]byte, []int) {
	return file_npool_appuser_mw_v1_kyc_kyc_proto_rawDescGZIP(), []int{3}
}

func (x *GetKycResponse) GetInfo() *Kyc {
	if x != nil {
		return x.Info
	}
	return nil
}

type GetKycsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Conds  *Conds `protobuf:"bytes,10,opt,name=Conds,proto3" json:"Conds,omitempty"`
	Offset int32  `protobuf:"varint,20,opt,name=Offset,proto3" json:"Offset,omitempty"`
	Limit  int32  `protobuf:"varint,30,opt,name=Limit,proto3" json:"Limit,omitempty"`
}

func (x *GetKycsRequest) Reset() {
	*x = GetKycsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_appuser_mw_v1_kyc_kyc_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetKycsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetKycsRequest) ProtoMessage() {}

func (x *GetKycsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_appuser_mw_v1_kyc_kyc_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetKycsRequest.ProtoReflect.Descriptor instead.
func (*GetKycsRequest) Descriptor() ([]byte, []int) {
	return file_npool_appuser_mw_v1_kyc_kyc_proto_rawDescGZIP(), []int{4}
}

func (x *GetKycsRequest) GetConds() *Conds {
	if x != nil {
		return x.Conds
	}
	return nil
}

func (x *GetKycsRequest) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *GetKycsRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type GetKycsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Infos []*Kyc `protobuf:"bytes,10,rep,name=Infos,proto3" json:"Infos,omitempty"`
	Total uint32 `protobuf:"varint,20,opt,name=Total,proto3" json:"Total,omitempty"`
}

func (x *GetKycsResponse) Reset() {
	*x = GetKycsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_appuser_mw_v1_kyc_kyc_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetKycsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetKycsResponse) ProtoMessage() {}

func (x *GetKycsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_appuser_mw_v1_kyc_kyc_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetKycsResponse.ProtoReflect.Descriptor instead.
func (*GetKycsResponse) Descriptor() ([]byte, []int) {
	return file_npool_appuser_mw_v1_kyc_kyc_proto_rawDescGZIP(), []int{5}
}

func (x *GetKycsResponse) GetInfos() []*Kyc {
	if x != nil {
		return x.Infos
	}
	return nil
}

func (x *GetKycsResponse) GetTotal() uint32 {
	if x != nil {
		return x.Total
	}
	return 0
}

var File_npool_appuser_mw_v1_kyc_kyc_proto protoreflect.FileDescriptor

var file_npool_appuser_mw_v1_kyc_kyc_proto_rawDesc = []byte{
	0x0a, 0x21, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x61, 0x70, 0x70, 0x75, 0x73, 0x65, 0x72, 0x2f,
	0x6d, 0x77, 0x2f, 0x76, 0x31, 0x2f, 0x6b, 0x79, 0x63, 0x2f, 0x6b, 0x79, 0x63, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x19, 0x61, 0x70, 0x70, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x6d, 0x69, 0x64,
	0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x6b, 0x79, 0x63, 0x2e, 0x76, 0x31, 0x1a, 0x1c,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d,
	0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x6e, 0x70, 0x6f, 0x6f, 0x6c,
	0x2f, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x22, 0x6e, 0x70,
	0x6f, 0x6f, 0x6c, 0x2f, 0x61, 0x70, 0x70, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x6d, 0x67, 0x72, 0x2f,
	0x76, 0x32, 0x2f, 0x6b, 0x79, 0x63, 0x2f, 0x6b, 0x79, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1d, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x2f, 0x6d,
	0x67, 0x72, 0x2f, 0x76, 0x32, 0x2f, 0x6d, 0x67, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xc9, 0x05, 0x0a, 0x03, 0x4b, 0x79, 0x63, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x41, 0x70, 0x70, 0x49, 0x44,
	0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x41, 0x70, 0x70, 0x49, 0x44, 0x12, 0x18, 0x0a,
	0x07, 0x41, 0x70, 0x70, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x41, 0x70, 0x70, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x41, 0x70, 0x70, 0x4c, 0x6f,
	0x67, 0x6f, 0x18, 0x28, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x41, 0x70, 0x70, 0x4c, 0x6f, 0x67,
	0x6f, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x32, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x22, 0x0a, 0x0c, 0x45, 0x6d, 0x61,
	0x69, 0x6c, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x3c, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x18, 0x0a,
	0x07, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x4f, 0x18, 0x46, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x50, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x4f, 0x12, 0x28, 0x0a, 0x0f, 0x44, 0x6f, 0x63, 0x75, 0x6d,
	0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x53, 0x74, 0x72, 0x18, 0x50, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0f, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x53, 0x74,
	0x72, 0x12, 0x4b, 0x0a, 0x0c, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70,
	0x65, 0x18, 0x5a, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x27, 0x2e, 0x61, 0x70, 0x70, 0x75, 0x73, 0x65,
	0x72, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x6b, 0x79, 0x63, 0x2e, 0x76, 0x32,
	0x2e, 0x4b, 0x79, 0x63, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65,
	0x52, 0x0c, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x49, 0x44, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x64, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x49, 0x44, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x46, 0x72,
	0x6f, 0x6e, 0x74, 0x49, 0x6d, 0x67, 0x18, 0x6e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x46, 0x72,
	0x6f, 0x6e, 0x74, 0x49, 0x6d, 0x67, 0x12, 0x18, 0x0a, 0x07, 0x42, 0x61, 0x63, 0x6b, 0x49, 0x6d,
	0x67, 0x18, 0x78, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x42, 0x61, 0x63, 0x6b, 0x49, 0x6d, 0x67,
	0x12, 0x1d, 0x0a, 0x09, 0x53, 0x65, 0x6c, 0x66, 0x69, 0x65, 0x49, 0x6d, 0x67, 0x18, 0x82, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x53, 0x65, 0x6c, 0x66, 0x69, 0x65, 0x49, 0x6d, 0x67, 0x12,
	0x25, 0x0a, 0x0d, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x54, 0x79, 0x70, 0x65, 0x53, 0x74, 0x72,
	0x18, 0x8c, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x54,
	0x79, 0x70, 0x65, 0x53, 0x74, 0x72, 0x12, 0x46, 0x0a, 0x0a, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x54, 0x79, 0x70, 0x65, 0x18, 0x96, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x25, 0x2e, 0x61, 0x70,
	0x70, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x6b, 0x79,
	0x63, 0x2e, 0x76, 0x32, 0x2e, 0x4b, 0x79, 0x63, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x54, 0x79,
	0x70, 0x65, 0x52, 0x0a, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1b,
	0x0a, 0x08, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x49, 0x44, 0x18, 0xa0, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x49, 0x44, 0x12, 0x27, 0x0a, 0x0e, 0x52,
	0x65, 0x76, 0x69, 0x65, 0x77, 0x53, 0x74, 0x61, 0x74, 0x65, 0x53, 0x74, 0x72, 0x18, 0xaa, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x53, 0x74, 0x61, 0x74,
	0x65, 0x53, 0x74, 0x72, 0x12, 0x37, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x74, 0x65, 0x18, 0xb4, 0x01,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x20, 0x2e, 0x61, 0x70, 0x70, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x6d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x6b, 0x79, 0x63, 0x2e, 0x76, 0x32, 0x2e, 0x4b, 0x79,
	0x63, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x1d, 0x0a,
	0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0xbe, 0x01, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x09,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0xc8, 0x01, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x3c, 0x0a, 0x05, 0x43,
	0x6f, 0x6e, 0x64, 0x73, 0x12, 0x33, 0x0a, 0x05, 0x43, 0x6f, 0x6e, 0x64, 0x73, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x61, 0x70, 0x70, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x6d, 0x61,
	0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x6b, 0x79, 0x63, 0x2e, 0x76, 0x32, 0x2e, 0x43, 0x6f, 0x6e,
	0x64, 0x73, 0x52, 0x05, 0x43, 0x6f, 0x6e, 0x64, 0x73, 0x22, 0x1f, 0x0a, 0x0d, 0x47, 0x65, 0x74,
	0x4b, 0x79, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x22, 0x44, 0x0a, 0x0e, 0x47, 0x65,
	0x74, 0x4b, 0x79, 0x63, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x32, 0x0a, 0x04,
	0x49, 0x6e, 0x66, 0x6f, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x61, 0x70, 0x70,
	0x75, 0x73, 0x65, 0x72, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e,
	0x6b, 0x79, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x4b, 0x79, 0x63, 0x52, 0x04, 0x49, 0x6e, 0x66, 0x6f,
	0x22, 0x76, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x4b, 0x79, 0x63, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x36, 0x0a, 0x05, 0x43, 0x6f, 0x6e, 0x64, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x20, 0x2e, 0x61, 0x70, 0x70, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x6d, 0x69, 0x64, 0x64,
	0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x6b, 0x79, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f,
	0x6e, 0x64, 0x73, 0x52, 0x05, 0x43, 0x6f, 0x6e, 0x64, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x4f, 0x66,
	0x66, 0x73, 0x65, 0x74, 0x18, 0x14, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x4f, 0x66, 0x66, 0x73,
	0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x1e, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x05, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x22, 0x5d, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x4b,
	0x79, 0x63, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x34, 0x0a, 0x05, 0x49,
	0x6e, 0x66, 0x6f, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x61, 0x70, 0x70,
	0x75, 0x73, 0x65, 0x72, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e,
	0x6b, 0x79, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x4b, 0x79, 0x63, 0x52, 0x05, 0x49, 0x6e, 0x66, 0x6f,
	0x73, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x14, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x32, 0xd1, 0x01, 0x0a, 0x0a, 0x4d, 0x69, 0x64, 0x64,
	0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x12, 0x5f, 0x0a, 0x06, 0x47, 0x65, 0x74, 0x4b, 0x79, 0x63,
	0x12, 0x28, 0x2e, 0x61, 0x70, 0x70, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c,
	0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x6b, 0x79, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74,
	0x4b, 0x79, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x61, 0x70, 0x70,
	0x75, 0x73, 0x65, 0x72, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e,
	0x6b, 0x79, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4b, 0x79, 0x63, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x62, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x4b, 0x79,
	0x63, 0x73, 0x12, 0x29, 0x2e, 0x61, 0x70, 0x70, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x6d, 0x69, 0x64,
	0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x6b, 0x79, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x47,
	0x65, 0x74, 0x4b, 0x79, 0x63, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2a, 0x2e,
	0x61, 0x70, 0x70, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61,
	0x72, 0x65, 0x2e, 0x6b, 0x79, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4b, 0x79, 0x63,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x3a, 0x5a, 0x38, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4e, 0x70, 0x6f, 0x6f, 0x6c, 0x50,
	0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f,
	0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x61, 0x70, 0x70, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x6d, 0x77,
	0x2f, 0x76, 0x31, 0x2f, 0x6b, 0x79, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_npool_appuser_mw_v1_kyc_kyc_proto_rawDescOnce sync.Once
	file_npool_appuser_mw_v1_kyc_kyc_proto_rawDescData = file_npool_appuser_mw_v1_kyc_kyc_proto_rawDesc
)

func file_npool_appuser_mw_v1_kyc_kyc_proto_rawDescGZIP() []byte {
	file_npool_appuser_mw_v1_kyc_kyc_proto_rawDescOnce.Do(func() {
		file_npool_appuser_mw_v1_kyc_kyc_proto_rawDescData = protoimpl.X.CompressGZIP(file_npool_appuser_mw_v1_kyc_kyc_proto_rawDescData)
	})
	return file_npool_appuser_mw_v1_kyc_kyc_proto_rawDescData
}

var file_npool_appuser_mw_v1_kyc_kyc_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_npool_appuser_mw_v1_kyc_kyc_proto_goTypes = []interface{}{
	(*Kyc)(nil),              // 0: appuser.middleware.kyc.v1.Kyc
	(*Conds)(nil),            // 1: appuser.middleware.kyc.v1.Conds
	(*GetKycRequest)(nil),    // 2: appuser.middleware.kyc.v1.GetKycRequest
	(*GetKycResponse)(nil),   // 3: appuser.middleware.kyc.v1.GetKycResponse
	(*GetKycsRequest)(nil),   // 4: appuser.middleware.kyc.v1.GetKycsRequest
	(*GetKycsResponse)(nil),  // 5: appuser.middleware.kyc.v1.GetKycsResponse
	(kyc.KycDocumentType)(0), // 6: appuser.manager.kyc.v2.KycDocumentType
	(kyc.KycEntityType)(0),   // 7: appuser.manager.kyc.v2.KycEntityType
	(kyc.KycState)(0),        // 8: appuser.manager.kyc.v2.KycState
	(*kyc.Conds)(nil),        // 9: appuser.manager.kyc.v2.Conds
}
var file_npool_appuser_mw_v1_kyc_kyc_proto_depIdxs = []int32{
	6, // 0: appuser.middleware.kyc.v1.Kyc.DocumentType:type_name -> appuser.manager.kyc.v2.KycDocumentType
	7, // 1: appuser.middleware.kyc.v1.Kyc.EntityType:type_name -> appuser.manager.kyc.v2.KycEntityType
	8, // 2: appuser.middleware.kyc.v1.Kyc.State:type_name -> appuser.manager.kyc.v2.KycState
	9, // 3: appuser.middleware.kyc.v1.Conds.Conds:type_name -> appuser.manager.kyc.v2.Conds
	0, // 4: appuser.middleware.kyc.v1.GetKycResponse.Info:type_name -> appuser.middleware.kyc.v1.Kyc
	1, // 5: appuser.middleware.kyc.v1.GetKycsRequest.Conds:type_name -> appuser.middleware.kyc.v1.Conds
	0, // 6: appuser.middleware.kyc.v1.GetKycsResponse.Infos:type_name -> appuser.middleware.kyc.v1.Kyc
	2, // 7: appuser.middleware.kyc.v1.Middleware.GetKyc:input_type -> appuser.middleware.kyc.v1.GetKycRequest
	4, // 8: appuser.middleware.kyc.v1.Middleware.GetKycs:input_type -> appuser.middleware.kyc.v1.GetKycsRequest
	3, // 9: appuser.middleware.kyc.v1.Middleware.GetKyc:output_type -> appuser.middleware.kyc.v1.GetKycResponse
	5, // 10: appuser.middleware.kyc.v1.Middleware.GetKycs:output_type -> appuser.middleware.kyc.v1.GetKycsResponse
	9, // [9:11] is the sub-list for method output_type
	7, // [7:9] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_npool_appuser_mw_v1_kyc_kyc_proto_init() }
func file_npool_appuser_mw_v1_kyc_kyc_proto_init() {
	if File_npool_appuser_mw_v1_kyc_kyc_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_npool_appuser_mw_v1_kyc_kyc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Kyc); i {
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
		file_npool_appuser_mw_v1_kyc_kyc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Conds); i {
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
		file_npool_appuser_mw_v1_kyc_kyc_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetKycRequest); i {
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
		file_npool_appuser_mw_v1_kyc_kyc_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetKycResponse); i {
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
		file_npool_appuser_mw_v1_kyc_kyc_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetKycsRequest); i {
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
		file_npool_appuser_mw_v1_kyc_kyc_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetKycsResponse); i {
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
			RawDescriptor: file_npool_appuser_mw_v1_kyc_kyc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_npool_appuser_mw_v1_kyc_kyc_proto_goTypes,
		DependencyIndexes: file_npool_appuser_mw_v1_kyc_kyc_proto_depIdxs,
		MessageInfos:      file_npool_appuser_mw_v1_kyc_kyc_proto_msgTypes,
	}.Build()
	File_npool_appuser_mw_v1_kyc_kyc_proto = out.File
	file_npool_appuser_mw_v1_kyc_kyc_proto_rawDesc = nil
	file_npool_appuser_mw_v1_kyc_kyc_proto_goTypes = nil
	file_npool_appuser_mw_v1_kyc_kyc_proto_depIdxs = nil
}
