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

type CoinArchivement struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CoinTypeID      string `protobuf:"bytes,10,opt,name=CoinTypeID,proto3" json:"CoinTypeID,omitempty"`
	CoinName        string `protobuf:"bytes,20,opt,name=CoinName,proto3" json:"CoinName,omitempty"`
	CoinLogo        string `protobuf:"bytes,30,opt,name=CoinLogo,proto3" json:"CoinLogo,omitempty"`
	CoinUnit        string `protobuf:"bytes,40,opt,name=CoinUnit,proto3" json:"CoinUnit,omitempty"`
	TotalUnits      uint32 `protobuf:"varint,50,opt,name=TotalUnits,proto3" json:"TotalUnits,omitempty"`
	SelfUnits       uint32 `protobuf:"varint,60,opt,name=SelfUnits,proto3" json:"SelfUnits,omitempty"`
	TotalAmount     string `protobuf:"bytes,70,opt,name=TotalAmount,proto3" json:"TotalAmount,omitempty"`
	SelfAmount      string `protobuf:"bytes,80,opt,name=SelfAmount,proto3" json:"SelfAmount,omitempty"`
	TotalCommission string `protobuf:"bytes,90,opt,name=TotalCommission,proto3" json:"TotalCommission,omitempty"`
	SelfCommission  string `protobuf:"bytes,100,opt,name=SelfCommission,proto3" json:"SelfCommission,omitempty"`
	CurPercent      uint32 `protobuf:"varint,110,opt,name=CurPercent,proto3" json:"CurPercent,omitempty"`
	CurGoodID       string `protobuf:"bytes,120,opt,name=CurGoodID,proto3" json:"CurGoodID,omitempty"`
	CurGoodName     string `protobuf:"bytes,130,opt,name=CurGoodName,proto3" json:"CurGoodName,omitempty"`
	CurGoodUnit     string `protobuf:"bytes,140,opt,name=CurGoodUnit,proto3" json:"CurGoodUnit,omitempty"`
}

func (x *CoinArchivement) Reset() {
	*x = CoinArchivement{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_inspire_gw_v1_archivement_archivement_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CoinArchivement) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CoinArchivement) ProtoMessage() {}

func (x *CoinArchivement) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use CoinArchivement.ProtoReflect.Descriptor instead.
func (*CoinArchivement) Descriptor() ([]byte, []int) {
	return file_npool_inspire_gw_v1_archivement_archivement_proto_rawDescGZIP(), []int{0}
}

func (x *CoinArchivement) GetCoinTypeID() string {
	if x != nil {
		return x.CoinTypeID
	}
	return ""
}

func (x *CoinArchivement) GetCoinName() string {
	if x != nil {
		return x.CoinName
	}
	return ""
}

func (x *CoinArchivement) GetCoinLogo() string {
	if x != nil {
		return x.CoinLogo
	}
	return ""
}

func (x *CoinArchivement) GetCoinUnit() string {
	if x != nil {
		return x.CoinUnit
	}
	return ""
}

func (x *CoinArchivement) GetTotalUnits() uint32 {
	if x != nil {
		return x.TotalUnits
	}
	return 0
}

func (x *CoinArchivement) GetSelfUnits() uint32 {
	if x != nil {
		return x.SelfUnits
	}
	return 0
}

func (x *CoinArchivement) GetTotalAmount() string {
	if x != nil {
		return x.TotalAmount
	}
	return ""
}

func (x *CoinArchivement) GetSelfAmount() string {
	if x != nil {
		return x.SelfAmount
	}
	return ""
}

func (x *CoinArchivement) GetTotalCommission() string {
	if x != nil {
		return x.TotalCommission
	}
	return ""
}

func (x *CoinArchivement) GetSelfCommission() string {
	if x != nil {
		return x.SelfCommission
	}
	return ""
}

func (x *CoinArchivement) GetCurPercent() uint32 {
	if x != nil {
		return x.CurPercent
	}
	return 0
}

func (x *CoinArchivement) GetCurGoodID() string {
	if x != nil {
		return x.CurGoodID
	}
	return ""
}

func (x *CoinArchivement) GetCurGoodName() string {
	if x != nil {
		return x.CurGoodName
	}
	return ""
}

func (x *CoinArchivement) GetCurGoodUnit() string {
	if x != nil {
		return x.CurGoodUnit
	}
	return ""
}

type GetCoinArchivementsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AppID  string `protobuf:"bytes,10,opt,name=AppID,proto3" json:"AppID,omitempty"`
	UserID string `protobuf:"bytes,20,opt,name=UserID,proto3" json:"UserID,omitempty"`
	Offset int32  `protobuf:"varint,30,opt,name=Offset,proto3" json:"Offset,omitempty"`
	Limit  int32  `protobuf:"varint,40,opt,name=Limit,proto3" json:"Limit,omitempty"`
}

func (x *GetCoinArchivementsRequest) Reset() {
	*x = GetCoinArchivementsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_inspire_gw_v1_archivement_archivement_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCoinArchivementsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCoinArchivementsRequest) ProtoMessage() {}

func (x *GetCoinArchivementsRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use GetCoinArchivementsRequest.ProtoReflect.Descriptor instead.
func (*GetCoinArchivementsRequest) Descriptor() ([]byte, []int) {
	return file_npool_inspire_gw_v1_archivement_archivement_proto_rawDescGZIP(), []int{1}
}

func (x *GetCoinArchivementsRequest) GetAppID() string {
	if x != nil {
		return x.AppID
	}
	return ""
}

func (x *GetCoinArchivementsRequest) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

func (x *GetCoinArchivementsRequest) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *GetCoinArchivementsRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type GetCoinArchivementsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Archivements []*GetCoinArchivementsResponse_Archivement `protobuf:"bytes,10,rep,name=Archivements,proto3" json:"Archivements,omitempty"` // Coin
	Total        uint32                                     `protobuf:"varint,20,opt,name=Total,proto3" json:"Total,omitempty"`
}

func (x *GetCoinArchivementsResponse) Reset() {
	*x = GetCoinArchivementsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_inspire_gw_v1_archivement_archivement_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCoinArchivementsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCoinArchivementsResponse) ProtoMessage() {}

func (x *GetCoinArchivementsResponse) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use GetCoinArchivementsResponse.ProtoReflect.Descriptor instead.
func (*GetCoinArchivementsResponse) Descriptor() ([]byte, []int) {
	return file_npool_inspire_gw_v1_archivement_archivement_proto_rawDescGZIP(), []int{2}
}

func (x *GetCoinArchivementsResponse) GetArchivements() []*GetCoinArchivementsResponse_Archivement {
	if x != nil {
		return x.Archivements
	}
	return nil
}

func (x *GetCoinArchivementsResponse) GetTotal() uint32 {
	if x != nil {
		return x.Total
	}
	return 0
}

type GetCoinArchivementsResponse_Archivement struct {
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
	Archivements  []*CoinArchivement `protobuf:"bytes,90,rep,name=Archivements,proto3" json:"Archivements,omitempty"`
}

func (x *GetCoinArchivementsResponse_Archivement) Reset() {
	*x = GetCoinArchivementsResponse_Archivement{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_inspire_gw_v1_archivement_archivement_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCoinArchivementsResponse_Archivement) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCoinArchivementsResponse_Archivement) ProtoMessage() {}

func (x *GetCoinArchivementsResponse_Archivement) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use GetCoinArchivementsResponse_Archivement.ProtoReflect.Descriptor instead.
func (*GetCoinArchivementsResponse_Archivement) Descriptor() ([]byte, []int) {
	return file_npool_inspire_gw_v1_archivement_archivement_proto_rawDescGZIP(), []int{2, 0}
}

func (x *GetCoinArchivementsResponse_Archivement) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

func (x *GetCoinArchivementsResponse_Archivement) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *GetCoinArchivementsResponse_Archivement) GetEmailAddress() string {
	if x != nil {
		return x.EmailAddress
	}
	return ""
}

func (x *GetCoinArchivementsResponse_Archivement) GetPhoneNO() string {
	if x != nil {
		return x.PhoneNO
	}
	return ""
}

func (x *GetCoinArchivementsResponse_Archivement) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *GetCoinArchivementsResponse_Archivement) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *GetCoinArchivementsResponse_Archivement) GetKol() bool {
	if x != nil {
		return x.Kol
	}
	return false
}

func (x *GetCoinArchivementsResponse_Archivement) GetTotalInvitees() uint32 {
	if x != nil {
		return x.TotalInvitees
	}
	return 0
}

func (x *GetCoinArchivementsResponse_Archivement) GetArchivements() []*CoinArchivement {
	if x != nil {
		return x.Archivements
	}
	return nil
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
	0x6f, 0x22, 0xdb, 0x03, 0x0a, 0x0f, 0x43, 0x6f, 0x69, 0x6e, 0x41, 0x72, 0x63, 0x68, 0x69, 0x76,
	0x65, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x43, 0x6f, 0x69, 0x6e, 0x54, 0x79, 0x70,
	0x65, 0x49, 0x44, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x43, 0x6f, 0x69, 0x6e, 0x54,
	0x79, 0x70, 0x65, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08, 0x43, 0x6f, 0x69, 0x6e, 0x4e, 0x61, 0x6d,
	0x65, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x43, 0x6f, 0x69, 0x6e, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x43, 0x6f, 0x69, 0x6e, 0x4c, 0x6f, 0x67, 0x6f, 0x18, 0x1e, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x43, 0x6f, 0x69, 0x6e, 0x4c, 0x6f, 0x67, 0x6f, 0x12, 0x1a, 0x0a,
	0x08, 0x43, 0x6f, 0x69, 0x6e, 0x55, 0x6e, 0x69, 0x74, 0x18, 0x28, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x43, 0x6f, 0x69, 0x6e, 0x55, 0x6e, 0x69, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x54, 0x6f, 0x74,
	0x61, 0x6c, 0x55, 0x6e, 0x69, 0x74, 0x73, 0x18, 0x32, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x54,
	0x6f, 0x74, 0x61, 0x6c, 0x55, 0x6e, 0x69, 0x74, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x53, 0x65, 0x6c,
	0x66, 0x55, 0x6e, 0x69, 0x74, 0x73, 0x18, 0x3c, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x53, 0x65,
	0x6c, 0x66, 0x55, 0x6e, 0x69, 0x74, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x54, 0x6f, 0x74, 0x61, 0x6c,
	0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x46, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x54, 0x6f,
	0x74, 0x61, 0x6c, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x53, 0x65, 0x6c,
	0x66, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x50, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x53,
	0x65, 0x6c, 0x66, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x28, 0x0a, 0x0f, 0x54, 0x6f, 0x74,
	0x61, 0x6c, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x5a, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0f, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x12, 0x26, 0x0a, 0x0e, 0x53, 0x65, 0x6c, 0x66, 0x43, 0x6f, 0x6d, 0x6d, 0x69,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x64, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x53, 0x65, 0x6c,
	0x66, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1e, 0x0a, 0x0a, 0x43,
	0x75, 0x72, 0x50, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x18, 0x6e, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x0a, 0x43, 0x75, 0x72, 0x50, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x43,
	0x75, 0x72, 0x47, 0x6f, 0x6f, 0x64, 0x49, 0x44, 0x18, 0x78, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x43, 0x75, 0x72, 0x47, 0x6f, 0x6f, 0x64, 0x49, 0x44, 0x12, 0x21, 0x0a, 0x0b, 0x43, 0x75, 0x72,
	0x47, 0x6f, 0x6f, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x82, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x43, 0x75, 0x72, 0x47, 0x6f, 0x6f, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0b,
	0x43, 0x75, 0x72, 0x47, 0x6f, 0x6f, 0x64, 0x55, 0x6e, 0x69, 0x74, 0x18, 0x8c, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x43, 0x75, 0x72, 0x47, 0x6f, 0x6f, 0x64, 0x55, 0x6e, 0x69, 0x74, 0x22,
	0x78, 0x0a, 0x1a, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x69, 0x6e, 0x41, 0x72, 0x63, 0x68, 0x69, 0x76,
	0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a,
	0x05, 0x41, 0x70, 0x70, 0x49, 0x44, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x41, 0x70,
	0x70, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x14, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x4f,
	0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x4f, 0x66, 0x66,
	0x73, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x28, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x05, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x22, 0xe9, 0x03, 0x0a, 0x1b, 0x47, 0x65,
	0x74, 0x43, 0x6f, 0x69, 0x6e, 0x41, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x6d, 0x65, 0x6e, 0x74,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x6b, 0x0a, 0x0c, 0x41, 0x72, 0x63,
	0x68, 0x69, 0x76, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x47, 0x2e, 0x69, 0x6e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61,
	0x79, 0x2e, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31,
	0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x69, 0x6e, 0x41, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x6d,
	0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x41, 0x72, 0x63,
	0x68, 0x69, 0x76, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x0c, 0x41, 0x72, 0x63, 0x68, 0x69, 0x76,
	0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x18,
	0x14, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x1a, 0xc6, 0x02, 0x0a,
	0x0b, 0x41, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06,
	0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x55, 0x73,
	0x65, 0x72, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x22, 0x0a, 0x0c, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x18, 0x1e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x41, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x4f, 0x18,
	0x28, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x4f, 0x12, 0x1c,
	0x0a, 0x09, 0x46, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x32, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x46, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x4c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x3c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x4c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x4b, 0x6f, 0x6c, 0x18,
	0x46, 0x20, 0x01, 0x28, 0x08, 0x52, 0x03, 0x4b, 0x6f, 0x6c, 0x12, 0x24, 0x0a, 0x0d, 0x54, 0x6f,
	0x74, 0x61, 0x6c, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x65, 0x73, 0x18, 0x50, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x0d, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x65, 0x73,
	0x12, 0x53, 0x0a, 0x0c, 0x41, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73,
	0x18, 0x5a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x69, 0x6e, 0x73, 0x70, 0x69, 0x72, 0x65,
	0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65,
	0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x69, 0x6e, 0x41, 0x72, 0x63, 0x68,
	0x69, 0x76, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x0c, 0x41, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65,
	0x6d, 0x65, 0x6e, 0x74, 0x73, 0x32, 0xbf, 0x01, 0x0a, 0x07, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61,
	0x79, 0x12, 0xb3, 0x01, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x69, 0x6e, 0x41, 0x72, 0x63,
	0x68, 0x69, 0x76, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x3a, 0x2e, 0x69, 0x6e, 0x73, 0x70,
	0x69, 0x72, 0x65, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x61, 0x72, 0x63, 0x68,
	0x69, 0x76, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f,
	0x69, 0x6e, 0x41, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3b, 0x2e, 0x69, 0x6e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e,
	0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x6d,
	0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x69, 0x6e, 0x41, 0x72,
	0x63, 0x68, 0x69, 0x76, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x23, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1d, 0x22, 0x18, 0x2f, 0x76, 0x31, 0x2f,
	0x67, 0x65, 0x74, 0x2f, 0x63, 0x6f, 0x69, 0x6e, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x6d,
	0x65, 0x6e, 0x74, 0x73, 0x3a, 0x01, 0x2a, 0x42, 0x42, 0x5a, 0x40, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4e, 0x70, 0x6f, 0x6f, 0x6c, 0x50, 0x6c, 0x61, 0x74, 0x66,
	0x6f, 0x72, 0x6d, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x6e, 0x70, 0x6f, 0x6f,
	0x6c, 0x2f, 0x69, 0x6e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2f, 0x67, 0x77, 0x2f, 0x76, 0x31, 0x2f,
	0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
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

var file_npool_inspire_gw_v1_archivement_archivement_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_npool_inspire_gw_v1_archivement_archivement_proto_goTypes = []interface{}{
	(*CoinArchivement)(nil),                         // 0: inspire.gateway.archivement.v1.CoinArchivement
	(*GetCoinArchivementsRequest)(nil),              // 1: inspire.gateway.archivement.v1.GetCoinArchivementsRequest
	(*GetCoinArchivementsResponse)(nil),             // 2: inspire.gateway.archivement.v1.GetCoinArchivementsResponse
	(*GetCoinArchivementsResponse_Archivement)(nil), // 3: inspire.gateway.archivement.v1.GetCoinArchivementsResponse.Archivement
}
var file_npool_inspire_gw_v1_archivement_archivement_proto_depIdxs = []int32{
	3, // 0: inspire.gateway.archivement.v1.GetCoinArchivementsResponse.Archivements:type_name -> inspire.gateway.archivement.v1.GetCoinArchivementsResponse.Archivement
	0, // 1: inspire.gateway.archivement.v1.GetCoinArchivementsResponse.Archivement.Archivements:type_name -> inspire.gateway.archivement.v1.CoinArchivement
	1, // 2: inspire.gateway.archivement.v1.Gateway.GetCoinArchivements:input_type -> inspire.gateway.archivement.v1.GetCoinArchivementsRequest
	2, // 3: inspire.gateway.archivement.v1.Gateway.GetCoinArchivements:output_type -> inspire.gateway.archivement.v1.GetCoinArchivementsResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_npool_inspire_gw_v1_archivement_archivement_proto_init() }
func file_npool_inspire_gw_v1_archivement_archivement_proto_init() {
	if File_npool_inspire_gw_v1_archivement_archivement_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_npool_inspire_gw_v1_archivement_archivement_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CoinArchivement); i {
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
			switch v := v.(*GetCoinArchivementsRequest); i {
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
			switch v := v.(*GetCoinArchivementsResponse); i {
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
			switch v := v.(*GetCoinArchivementsResponse_Archivement); i {
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
			NumMessages:   4,
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
