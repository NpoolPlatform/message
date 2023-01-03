// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.18.1
// source: npool/inspire/gw/v1/archivement/archivement.proto

package archivement

import (
	_ "github.com/NpoolPlatform/message/npool"
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

type GoodArchivement struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GoodID            string `protobuf:"bytes,10,opt,name=GoodID,proto3" json:"GoodID,omitempty"`
	GoodName          string `protobuf:"bytes,20,opt,name=GoodName,proto3" json:"GoodName,omitempty"`
	GoodUnit          string `protobuf:"bytes,30,opt,name=GoodUnit,proto3" json:"GoodUnit,omitempty"`
	CommissionPercent string `protobuf:"bytes,40,opt,name=CommissionPercent,proto3" json:"CommissionPercent,omitempty"`
	CoinTypeID        string `protobuf:"bytes,50,opt,name=CoinTypeID,proto3" json:"CoinTypeID,omitempty"`
	CoinName          string `protobuf:"bytes,60,opt,name=CoinName,proto3" json:"CoinName,omitempty"`
	CoinLogo          string `protobuf:"bytes,70,opt,name=CoinLogo,proto3" json:"CoinLogo,omitempty"`
	CoinUnit          string `protobuf:"bytes,80,opt,name=CoinUnit,proto3" json:"CoinUnit,omitempty"`
	TotalUnits        uint32 `protobuf:"varint,90,opt,name=TotalUnits,proto3" json:"TotalUnits,omitempty"`
	SelfUnits         uint32 `protobuf:"varint,100,opt,name=SelfUnits,proto3" json:"SelfUnits,omitempty"`
	// In USD
	TotalAmount        string `protobuf:"bytes,110,opt,name=TotalAmount,proto3" json:"TotalAmount,omitempty"`
	SelfAmount         string `protobuf:"bytes,120,opt,name=SelfAmount,proto3" json:"SelfAmount,omitempty"`
	TotalCommission    string `protobuf:"bytes,130,opt,name=TotalCommission,proto3" json:"TotalCommission,omitempty"`
	SelfCommission     string `protobuf:"bytes,140,opt,name=SelfCommission,proto3" json:"SelfCommission,omitempty"`
	SuperiorCommission string `protobuf:"bytes,150,opt,name=SuperiorCommission,proto3" json:"SuperiorCommission,omitempty"`
}

func (x *GoodArchivement) Reset() {
	*x = GoodArchivement{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_inspire_gw_v1_archivement_archivement_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GoodArchivement) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GoodArchivement) ProtoMessage() {}

func (x *GoodArchivement) ProtoReflect() protoreflect.Message {
	mi := &file_npool_inspire_gw_v1_archivement_archivement_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GoodArchivement.ProtoReflect.Descriptor instead.
func (*GoodArchivement) Descriptor() ([]byte, []int) {
	return file_npool_inspire_gw_v1_archivement_archivement_proto_rawDescGZIP(), []int{0}
}

func (x *GoodArchivement) GetGoodID() string {
	if x != nil {
		return x.GoodID
	}
	return ""
}

func (x *GoodArchivement) GetGoodName() string {
	if x != nil {
		return x.GoodName
	}
	return ""
}

func (x *GoodArchivement) GetGoodUnit() string {
	if x != nil {
		return x.GoodUnit
	}
	return ""
}

func (x *GoodArchivement) GetCommissionPercent() string {
	if x != nil {
		return x.CommissionPercent
	}
	return ""
}

func (x *GoodArchivement) GetCoinTypeID() string {
	if x != nil {
		return x.CoinTypeID
	}
	return ""
}

func (x *GoodArchivement) GetCoinName() string {
	if x != nil {
		return x.CoinName
	}
	return ""
}

func (x *GoodArchivement) GetCoinLogo() string {
	if x != nil {
		return x.CoinLogo
	}
	return ""
}

func (x *GoodArchivement) GetCoinUnit() string {
	if x != nil {
		return x.CoinUnit
	}
	return ""
}

func (x *GoodArchivement) GetTotalUnits() uint32 {
	if x != nil {
		return x.TotalUnits
	}
	return 0
}

func (x *GoodArchivement) GetSelfUnits() uint32 {
	if x != nil {
		return x.SelfUnits
	}
	return 0
}

func (x *GoodArchivement) GetTotalAmount() string {
	if x != nil {
		return x.TotalAmount
	}
	return ""
}

func (x *GoodArchivement) GetSelfAmount() string {
	if x != nil {
		return x.SelfAmount
	}
	return ""
}

func (x *GoodArchivement) GetTotalCommission() string {
	if x != nil {
		return x.TotalCommission
	}
	return ""
}

func (x *GoodArchivement) GetSelfCommission() string {
	if x != nil {
		return x.SelfCommission
	}
	return ""
}

func (x *GoodArchivement) GetSuperiorCommission() string {
	if x != nil {
		return x.SuperiorCommission
	}
	return ""
}

type UserArchivement struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID        string             `protobuf:"bytes,10,opt,name=UserID,proto3" json:"UserID,omitempty"`
	Username      string             `protobuf:"bytes,20,opt,name=Username,proto3" json:"Username,omitempty"`
	EmailAddress  string             `protobuf:"bytes,30,opt,name=EmailAddress,proto3" json:"EmailAddress,omitempty"`
	PhoneNO       string             `protobuf:"bytes,40,opt,name=PhoneNO,proto3" json:"PhoneNO,omitempty"`
	FirstName     string             `protobuf:"bytes,50,opt,name=FirstName,proto3" json:"FirstName,omitempty"`
	LastName      string             `protobuf:"bytes,60,opt,name=LastName,proto3" json:"LastName,omitempty"`
	Kol           bool               `protobuf:"varint,70,opt,name=Kol,proto3" json:"Kol,omitempty"`
	TotalInvitees uint32             `protobuf:"varint,80,opt,name=TotalInvitees,proto3" json:"TotalInvitees,omitempty"`
	Archivements  []*GoodArchivement `protobuf:"bytes,90,rep,name=Archivements,proto3" json:"Archivements,omitempty"`
	CreatedAt     uint32             `protobuf:"varint,100,opt,name=CreatedAt,proto3" json:"CreatedAt,omitempty"`
	InvitedAt     uint32             `protobuf:"varint,110,opt,name=InvitedAt,proto3" json:"InvitedAt,omitempty"`
}

func (x *UserArchivement) Reset() {
	*x = UserArchivement{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_inspire_gw_v1_archivement_archivement_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserArchivement) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserArchivement) ProtoMessage() {}

func (x *UserArchivement) ProtoReflect() protoreflect.Message {
	mi := &file_npool_inspire_gw_v1_archivement_archivement_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserArchivement.ProtoReflect.Descriptor instead.
func (*UserArchivement) Descriptor() ([]byte, []int) {
	return file_npool_inspire_gw_v1_archivement_archivement_proto_rawDescGZIP(), []int{1}
}

func (x *UserArchivement) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

func (x *UserArchivement) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *UserArchivement) GetEmailAddress() string {
	if x != nil {
		return x.EmailAddress
	}
	return ""
}

func (x *UserArchivement) GetPhoneNO() string {
	if x != nil {
		return x.PhoneNO
	}
	return ""
}

func (x *UserArchivement) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *UserArchivement) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *UserArchivement) GetKol() bool {
	if x != nil {
		return x.Kol
	}
	return false
}

func (x *UserArchivement) GetTotalInvitees() uint32 {
	if x != nil {
		return x.TotalInvitees
	}
	return 0
}

func (x *UserArchivement) GetArchivements() []*GoodArchivement {
	if x != nil {
		return x.Archivements
	}
	return nil
}

func (x *UserArchivement) GetCreatedAt() uint32 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *UserArchivement) GetInvitedAt() uint32 {
	if x != nil {
		return x.InvitedAt
	}
	return 0
}

type GetGoodArchivementsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AppID  string `protobuf:"bytes,10,opt,name=AppID,proto3" json:"AppID,omitempty"`
	UserID string `protobuf:"bytes,20,opt,name=UserID,proto3" json:"UserID,omitempty"`
	Offset int32  `protobuf:"varint,30,opt,name=Offset,proto3" json:"Offset,omitempty"`
	Limit  int32  `protobuf:"varint,40,opt,name=Limit,proto3" json:"Limit,omitempty"`
}

func (x *GetGoodArchivementsRequest) Reset() {
	*x = GetGoodArchivementsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_inspire_gw_v1_archivement_archivement_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetGoodArchivementsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetGoodArchivementsRequest) ProtoMessage() {}

func (x *GetGoodArchivementsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_inspire_gw_v1_archivement_archivement_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetGoodArchivementsRequest.ProtoReflect.Descriptor instead.
func (*GetGoodArchivementsRequest) Descriptor() ([]byte, []int) {
	return file_npool_inspire_gw_v1_archivement_archivement_proto_rawDescGZIP(), []int{2}
}

func (x *GetGoodArchivementsRequest) GetAppID() string {
	if x != nil {
		return x.AppID
	}
	return ""
}

func (x *GetGoodArchivementsRequest) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

func (x *GetGoodArchivementsRequest) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *GetGoodArchivementsRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type GetGoodArchivementsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Archivements []*UserArchivement `protobuf:"bytes,10,rep,name=Archivements,proto3" json:"Archivements,omitempty"` // Each user
	Total        uint32             `protobuf:"varint,20,opt,name=Total,proto3" json:"Total,omitempty"`
}

func (x *GetGoodArchivementsResponse) Reset() {
	*x = GetGoodArchivementsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_inspire_gw_v1_archivement_archivement_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetGoodArchivementsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetGoodArchivementsResponse) ProtoMessage() {}

func (x *GetGoodArchivementsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_inspire_gw_v1_archivement_archivement_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetGoodArchivementsResponse.ProtoReflect.Descriptor instead.
func (*GetGoodArchivementsResponse) Descriptor() ([]byte, []int) {
	return file_npool_inspire_gw_v1_archivement_archivement_proto_rawDescGZIP(), []int{3}
}

func (x *GetGoodArchivementsResponse) GetArchivements() []*UserArchivement {
	if x != nil {
		return x.Archivements
	}
	return nil
}

func (x *GetGoodArchivementsResponse) GetTotal() uint32 {
	if x != nil {
		return x.Total
	}
	return 0
}

type GetUserGoodArchivementsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AppID   string   `protobuf:"bytes,10,opt,name=AppID,proto3" json:"AppID,omitempty"`
	UserIDs []string `protobuf:"bytes,20,rep,name=UserIDs,proto3" json:"UserIDs,omitempty"`
	Offset  int32    `protobuf:"varint,30,opt,name=Offset,proto3" json:"Offset,omitempty"`
	Limit   int32    `protobuf:"varint,40,opt,name=Limit,proto3" json:"Limit,omitempty"`
}

func (x *GetUserGoodArchivementsRequest) Reset() {
	*x = GetUserGoodArchivementsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_inspire_gw_v1_archivement_archivement_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserGoodArchivementsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserGoodArchivementsRequest) ProtoMessage() {}

func (x *GetUserGoodArchivementsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_inspire_gw_v1_archivement_archivement_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserGoodArchivementsRequest.ProtoReflect.Descriptor instead.
func (*GetUserGoodArchivementsRequest) Descriptor() ([]byte, []int) {
	return file_npool_inspire_gw_v1_archivement_archivement_proto_rawDescGZIP(), []int{4}
}

func (x *GetUserGoodArchivementsRequest) GetAppID() string {
	if x != nil {
		return x.AppID
	}
	return ""
}

func (x *GetUserGoodArchivementsRequest) GetUserIDs() []string {
	if x != nil {
		return x.UserIDs
	}
	return nil
}

func (x *GetUserGoodArchivementsRequest) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *GetUserGoodArchivementsRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type GetUserGoodArchivementsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Archivements []*UserArchivement `protobuf:"bytes,10,rep,name=Archivements,proto3" json:"Archivements,omitempty"`
	Total        uint32             `protobuf:"varint,20,opt,name=Total,proto3" json:"Total,omitempty"`
}

func (x *GetUserGoodArchivementsResponse) Reset() {
	*x = GetUserGoodArchivementsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_inspire_gw_v1_archivement_archivement_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserGoodArchivementsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserGoodArchivementsResponse) ProtoMessage() {}

func (x *GetUserGoodArchivementsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_inspire_gw_v1_archivement_archivement_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserGoodArchivementsResponse.ProtoReflect.Descriptor instead.
func (*GetUserGoodArchivementsResponse) Descriptor() ([]byte, []int) {
	return file_npool_inspire_gw_v1_archivement_archivement_proto_rawDescGZIP(), []int{5}
}

func (x *GetUserGoodArchivementsResponse) GetArchivements() []*UserArchivement {
	if x != nil {
		return x.Archivements
	}
	return nil
}

func (x *GetUserGoodArchivementsResponse) GetTotal() uint32 {
	if x != nil {
		return x.Total
	}
	return 0
}

var File_npool_inspire_gw_v1_archivement_archivement_proto protoreflect.FileDescriptor

var file_npool_inspire_gw_v1_archivement_archivement_proto_rawDesc = []byte{
	0x0a, 0x31, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x69, 0x6e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2f,
	0x67, 0x77, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x2f, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x1e, 0x69, 0x6e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x67, 0x61, 0x74,
	0x65, 0x77, 0x61, 0x79, 0x2e, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x6d, 0x65, 0x6e, 0x74,
	0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11,
	0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x88, 0x04, 0x0a, 0x0f, 0x47, 0x6f, 0x6f, 0x64, 0x41, 0x72, 0x63, 0x68, 0x69, 0x76,
	0x65, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x47, 0x6f, 0x6f, 0x64, 0x49, 0x44, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x47, 0x6f, 0x6f, 0x64, 0x49, 0x44, 0x12, 0x1a, 0x0a,
	0x08, 0x47, 0x6f, 0x6f, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x47, 0x6f, 0x6f, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x47, 0x6f, 0x6f,
	0x64, 0x55, 0x6e, 0x69, 0x74, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x47, 0x6f, 0x6f,
	0x64, 0x55, 0x6e, 0x69, 0x74, 0x12, 0x2c, 0x0a, 0x11, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x50, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x18, 0x28, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x11, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x50, 0x65, 0x72, 0x63,
	0x65, 0x6e, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x43, 0x6f, 0x69, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x49,
	0x44, 0x18, 0x32, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x43, 0x6f, 0x69, 0x6e, 0x54, 0x79, 0x70,
	0x65, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08, 0x43, 0x6f, 0x69, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x18,
	0x3c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x43, 0x6f, 0x69, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x43, 0x6f, 0x69, 0x6e, 0x4c, 0x6f, 0x67, 0x6f, 0x18, 0x46, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x43, 0x6f, 0x69, 0x6e, 0x4c, 0x6f, 0x67, 0x6f, 0x12, 0x1a, 0x0a, 0x08, 0x43,
	0x6f, 0x69, 0x6e, 0x55, 0x6e, 0x69, 0x74, 0x18, 0x50, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x43,
	0x6f, 0x69, 0x6e, 0x55, 0x6e, 0x69, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x54, 0x6f, 0x74, 0x61, 0x6c,
	0x55, 0x6e, 0x69, 0x74, 0x73, 0x18, 0x5a, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x54, 0x6f, 0x74,
	0x61, 0x6c, 0x55, 0x6e, 0x69, 0x74, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x53, 0x65, 0x6c, 0x66, 0x55,
	0x6e, 0x69, 0x74, 0x73, 0x18, 0x64, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x53, 0x65, 0x6c, 0x66,
	0x55, 0x6e, 0x69, 0x74, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x41, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x6e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x54, 0x6f, 0x74, 0x61,
	0x6c, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x53, 0x65, 0x6c, 0x66, 0x41,
	0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x78, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x53, 0x65, 0x6c,
	0x66, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x29, 0x0a, 0x0f, 0x54, 0x6f, 0x74, 0x61, 0x6c,
	0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x82, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0f, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x12, 0x27, 0x0a, 0x0e, 0x53, 0x65, 0x6c, 0x66, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x18, 0x8c, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x53, 0x65, 0x6c,
	0x66, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x2f, 0x0a, 0x12, 0x53,
	0x75, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x72, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x18, 0x96, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x53, 0x75, 0x70, 0x65, 0x72, 0x69,
	0x6f, 0x72, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x86, 0x03, 0x0a,
	0x0f, 0x55, 0x73, 0x65, 0x72, 0x41, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x6d, 0x65, 0x6e, 0x74,
	0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x55, 0x73, 0x65, 0x72,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x41, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x45, 0x6d, 0x61, 0x69,
	0x6c, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x50, 0x68, 0x6f, 0x6e,
	0x65, 0x4e, 0x4f, 0x18, 0x28, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x50, 0x68, 0x6f, 0x6e, 0x65,
	0x4e, 0x4f, 0x12, 0x1c, 0x0a, 0x09, 0x46, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18,
	0x32, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x46, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x4c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x3c, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x4c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03,
	0x4b, 0x6f, 0x6c, 0x18, 0x46, 0x20, 0x01, 0x28, 0x08, 0x52, 0x03, 0x4b, 0x6f, 0x6c, 0x12, 0x24,
	0x0a, 0x0d, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x65, 0x73, 0x18,
	0x50, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0d, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x49, 0x6e, 0x76, 0x69,
	0x74, 0x65, 0x65, 0x73, 0x12, 0x53, 0x0a, 0x0c, 0x41, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x6d,
	0x65, 0x6e, 0x74, 0x73, 0x18, 0x5a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x69, 0x6e, 0x73,
	0x70, 0x69, 0x72, 0x65, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x61, 0x72, 0x63,
	0x68, 0x69, 0x76, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x6f, 0x6f, 0x64,
	0x41, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x0c, 0x41, 0x72, 0x63,
	0x68, 0x69, 0x76, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x64, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x49, 0x6e, 0x76, 0x69, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x18, 0x6e, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x49, 0x6e, 0x76, 0x69,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x78, 0x0a, 0x1a, 0x47, 0x65, 0x74, 0x47, 0x6f, 0x6f, 0x64,
	0x41, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x41, 0x70, 0x70, 0x49, 0x44, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x41, 0x70, 0x70, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65,
	0x72, 0x49, 0x44, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49,
	0x44, 0x12, 0x16, 0x0a, 0x06, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x1e, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x06, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x4c, 0x69, 0x6d,
	0x69, 0x74, 0x18, 0x28, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x22,
	0x88, 0x01, 0x0a, 0x1b, 0x47, 0x65, 0x74, 0x47, 0x6f, 0x6f, 0x64, 0x41, 0x72, 0x63, 0x68, 0x69,
	0x76, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x53, 0x0a, 0x0c, 0x41, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18,
	0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x69, 0x6e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e,
	0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x6d,
	0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x41, 0x72, 0x63, 0x68, 0x69,
	0x76, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x0c, 0x41, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x6d,
	0x65, 0x6e, 0x74, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x14, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x22, 0x7e, 0x0a, 0x1e, 0x47, 0x65,
	0x74, 0x55, 0x73, 0x65, 0x72, 0x47, 0x6f, 0x6f, 0x64, 0x41, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65,
	0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05,
	0x41, 0x70, 0x70, 0x49, 0x44, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x41, 0x70, 0x70,
	0x49, 0x44, 0x12, 0x18, 0x0a, 0x07, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x73, 0x18, 0x14, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x07, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x73, 0x12, 0x16, 0x0a, 0x06,
	0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x4f, 0x66,
	0x66, 0x73, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x28, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x05, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x22, 0x8c, 0x01, 0x0a, 0x1f, 0x47,
	0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x47, 0x6f, 0x6f, 0x64, 0x41, 0x72, 0x63, 0x68, 0x69, 0x76,
	0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x53,
	0x0a, 0x0c, 0x41, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x0a,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x69, 0x6e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x67,
	0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x6d, 0x65,
	0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x41, 0x72, 0x63, 0x68, 0x69, 0x76,
	0x65, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x0c, 0x41, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x6d, 0x65,
	0x6e, 0x74, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x14, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x32, 0x86, 0x03, 0x0a, 0x07, 0x47, 0x61,
	0x74, 0x65, 0x77, 0x61, 0x79, 0x12, 0xb3, 0x01, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x47, 0x6f, 0x6f,
	0x64, 0x41, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x3a, 0x2e,
	0x69, 0x6e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e,
	0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x47,
	0x65, 0x74, 0x47, 0x6f, 0x6f, 0x64, 0x41, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3b, 0x2e, 0x69, 0x6e, 0x73, 0x70,
	0x69, 0x72, 0x65, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x61, 0x72, 0x63, 0x68,
	0x69, 0x76, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x47, 0x6f,
	0x6f, 0x64, 0x41, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x23, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1d, 0x22, 0x18,
	0x2f, 0x76, 0x31, 0x2f, 0x67, 0x65, 0x74, 0x2f, 0x67, 0x6f, 0x6f, 0x64, 0x61, 0x72, 0x63, 0x68,
	0x69, 0x76, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x3a, 0x01, 0x2a, 0x12, 0xc4, 0x01, 0x0a, 0x17,
	0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x47, 0x6f, 0x6f, 0x64, 0x41, 0x72, 0x63, 0x68, 0x69,
	0x76, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x3e, 0x2e, 0x69, 0x6e, 0x73, 0x70, 0x69, 0x72,
	0x65, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76,
	0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72,
	0x47, 0x6f, 0x6f, 0x64, 0x41, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3f, 0x2e, 0x69, 0x6e, 0x73, 0x70, 0x69, 0x72,
	0x65, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76,
	0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72,
	0x47, 0x6f, 0x6f, 0x64, 0x41, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x28, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x22,
	0x22, 0x1d, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x65, 0x74, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x67,
	0x6f, 0x6f, 0x64, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x3a,
	0x01, 0x2a, 0x42, 0x42, 0x5a, 0x40, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x4e, 0x70, 0x6f, 0x6f, 0x6c, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x69, 0x6e, 0x73,
	0x70, 0x69, 0x72, 0x65, 0x2f, 0x67, 0x77, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x72, 0x63, 0x68, 0x69,
	0x76, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_npool_inspire_gw_v1_archivement_archivement_proto_rawDescOnce sync.Once
	file_npool_inspire_gw_v1_archivement_archivement_proto_rawDescData = file_npool_inspire_gw_v1_archivement_archivement_proto_rawDesc
)

func file_npool_inspire_gw_v1_archivement_archivement_proto_rawDescGZIP() []byte {
	file_npool_inspire_gw_v1_archivement_archivement_proto_rawDescOnce.Do(func() {
		file_npool_inspire_gw_v1_archivement_archivement_proto_rawDescData = protoimpl.X.CompressGZIP(file_npool_inspire_gw_v1_archivement_archivement_proto_rawDescData)
	})
	return file_npool_inspire_gw_v1_archivement_archivement_proto_rawDescData
}

var file_npool_inspire_gw_v1_archivement_archivement_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_npool_inspire_gw_v1_archivement_archivement_proto_goTypes = []interface{}{
	(*GoodArchivement)(nil),                 // 0: inspire.gateway.archivement.v1.GoodArchivement
	(*UserArchivement)(nil),                 // 1: inspire.gateway.archivement.v1.UserArchivement
	(*GetGoodArchivementsRequest)(nil),      // 2: inspire.gateway.archivement.v1.GetGoodArchivementsRequest
	(*GetGoodArchivementsResponse)(nil),     // 3: inspire.gateway.archivement.v1.GetGoodArchivementsResponse
	(*GetUserGoodArchivementsRequest)(nil),  // 4: inspire.gateway.archivement.v1.GetUserGoodArchivementsRequest
	(*GetUserGoodArchivementsResponse)(nil), // 5: inspire.gateway.archivement.v1.GetUserGoodArchivementsResponse
}
var file_npool_inspire_gw_v1_archivement_archivement_proto_depIdxs = []int32{
	0, // 0: inspire.gateway.archivement.v1.UserArchivement.Archivements:type_name -> inspire.gateway.archivement.v1.GoodArchivement
	1, // 1: inspire.gateway.archivement.v1.GetGoodArchivementsResponse.Archivements:type_name -> inspire.gateway.archivement.v1.UserArchivement
	1, // 2: inspire.gateway.archivement.v1.GetUserGoodArchivementsResponse.Archivements:type_name -> inspire.gateway.archivement.v1.UserArchivement
	2, // 3: inspire.gateway.archivement.v1.Gateway.GetGoodArchivements:input_type -> inspire.gateway.archivement.v1.GetGoodArchivementsRequest
	4, // 4: inspire.gateway.archivement.v1.Gateway.GetUserGoodArchivements:input_type -> inspire.gateway.archivement.v1.GetUserGoodArchivementsRequest
	3, // 5: inspire.gateway.archivement.v1.Gateway.GetGoodArchivements:output_type -> inspire.gateway.archivement.v1.GetGoodArchivementsResponse
	5, // 6: inspire.gateway.archivement.v1.Gateway.GetUserGoodArchivements:output_type -> inspire.gateway.archivement.v1.GetUserGoodArchivementsResponse
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_npool_inspire_gw_v1_archivement_archivement_proto_init() }
func file_npool_inspire_gw_v1_archivement_archivement_proto_init() {
	if File_npool_inspire_gw_v1_archivement_archivement_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_npool_inspire_gw_v1_archivement_archivement_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GoodArchivement); i {
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
		file_npool_inspire_gw_v1_archivement_archivement_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserArchivement); i {
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
		file_npool_inspire_gw_v1_archivement_archivement_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetGoodArchivementsRequest); i {
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
		file_npool_inspire_gw_v1_archivement_archivement_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetGoodArchivementsResponse); i {
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
		file_npool_inspire_gw_v1_archivement_archivement_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserGoodArchivementsRequest); i {
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
		file_npool_inspire_gw_v1_archivement_archivement_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserGoodArchivementsResponse); i {
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
			RawDescriptor: file_npool_inspire_gw_v1_archivement_archivement_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_npool_inspire_gw_v1_archivement_archivement_proto_goTypes,
		DependencyIndexes: file_npool_inspire_gw_v1_archivement_archivement_proto_depIdxs,
		MessageInfos:      file_npool_inspire_gw_v1_archivement_archivement_proto_msgTypes,
	}.Build()
	File_npool_inspire_gw_v1_archivement_archivement_proto = out.File
	file_npool_inspire_gw_v1_archivement_archivement_proto_rawDesc = nil
	file_npool_inspire_gw_v1_archivement_archivement_proto_goTypes = nil
	file_npool_inspire_gw_v1_archivement_archivement_proto_depIdxs = nil
}
