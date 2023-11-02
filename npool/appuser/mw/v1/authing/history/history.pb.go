// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.18.1
// source: npool/appuser/mw/v1/authing/history/history.proto

package history

import (
	v1 "github.com/NpoolPlatform/message/npool/basetypes/v1"
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

type HistoryReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID       *uint32 `protobuf:"varint,9,opt,name=ID,proto3,oneof" json:"ID,omitempty"`
	EntID    *string `protobuf:"bytes,10,opt,name=EntID,proto3,oneof" json:"EntID,omitempty"`
	AppID    *string `protobuf:"bytes,20,opt,name=AppID,proto3,oneof" json:"AppID,omitempty"`
	UserID   *string `protobuf:"bytes,30,opt,name=UserID,proto3,oneof" json:"UserID,omitempty"`
	Resource *string `protobuf:"bytes,40,opt,name=Resource,proto3,oneof" json:"Resource,omitempty"`
	Method   *string `protobuf:"bytes,50,opt,name=Method,proto3,oneof" json:"Method,omitempty"`
	Allowed  *bool   `protobuf:"varint,60,opt,name=Allowed,proto3,oneof" json:"Allowed,omitempty"`
}

func (x *HistoryReq) Reset() {
	*x = HistoryReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_appuser_mw_v1_authing_history_history_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HistoryReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HistoryReq) ProtoMessage() {}

func (x *HistoryReq) ProtoReflect() protoreflect.Message {
	mi := &file_npool_appuser_mw_v1_authing_history_history_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HistoryReq.ProtoReflect.Descriptor instead.
func (*HistoryReq) Descriptor() ([]byte, []int) {
	return file_npool_appuser_mw_v1_authing_history_history_proto_rawDescGZIP(), []int{0}
}

func (x *HistoryReq) GetID() uint32 {
	if x != nil && x.ID != nil {
		return *x.ID
	}
	return 0
}

func (x *HistoryReq) GetEntID() string {
	if x != nil && x.EntID != nil {
		return *x.EntID
	}
	return ""
}

func (x *HistoryReq) GetAppID() string {
	if x != nil && x.AppID != nil {
		return *x.AppID
	}
	return ""
}

func (x *HistoryReq) GetUserID() string {
	if x != nil && x.UserID != nil {
		return *x.UserID
	}
	return ""
}

func (x *HistoryReq) GetResource() string {
	if x != nil && x.Resource != nil {
		return *x.Resource
	}
	return ""
}

func (x *HistoryReq) GetMethod() string {
	if x != nil && x.Method != nil {
		return *x.Method
	}
	return ""
}

func (x *HistoryReq) GetAllowed() bool {
	if x != nil && x.Allowed != nil {
		return *x.Allowed
	}
	return false
}

type History struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: sql:"app_id"
	AppID string `protobuf:"bytes,10,opt,name=AppID,proto3" json:"AppID,omitempty" sql:"app_id"`
	// @inject_tag: sql:"app_name"
	AppName string `protobuf:"bytes,20,opt,name=AppName,proto3" json:"AppName,omitempty" sql:"app_name"`
	// @inject_tag: sql:"app_logo"
	AppLogo string `protobuf:"bytes,30,opt,name=AppLogo,proto3" json:"AppLogo,omitempty" sql:"app_logo"`
	// @inject_tag: sql:"user_id"
	UserID string `protobuf:"bytes,40,opt,name=UserID,proto3" json:"UserID,omitempty" sql:"user_id"`
	// @inject_tag: sql:"email_address"
	EmailAddress string `protobuf:"bytes,50,opt,name=EmailAddress,proto3" json:"EmailAddress,omitempty" sql:"email_address"`
	// @inject_tag: sql:"phone_no"
	PhoneNO string `protobuf:"bytes,60,opt,name=PhoneNO,proto3" json:"PhoneNO,omitempty" sql:"phone_no"`
	// @inject_tag: sql:"resource"
	Resource string `protobuf:"bytes,70,opt,name=Resource,proto3" json:"Resource,omitempty" sql:"resource"`
	// @inject_tag: sql:"method"
	Method string `protobuf:"bytes,80,opt,name=Method,proto3" json:"Method,omitempty" sql:"method"`
	// @inject_tag: sql:"allowed"
	Allowed bool `protobuf:"varint,90,opt,name=Allowed,proto3" json:"Allowed,omitempty" sql:"allowed"`
	// @inject_tag: sql:"created_at"
	CreatedAt uint32 `protobuf:"varint,100,opt,name=CreatedAt,proto3" json:"CreatedAt,omitempty" sql:"created_at"`
}

func (x *History) Reset() {
	*x = History{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_appuser_mw_v1_authing_history_history_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *History) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*History) ProtoMessage() {}

func (x *History) ProtoReflect() protoreflect.Message {
	mi := &file_npool_appuser_mw_v1_authing_history_history_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use History.ProtoReflect.Descriptor instead.
func (*History) Descriptor() ([]byte, []int) {
	return file_npool_appuser_mw_v1_authing_history_history_proto_rawDescGZIP(), []int{1}
}

func (x *History) GetAppID() string {
	if x != nil {
		return x.AppID
	}
	return ""
}

func (x *History) GetAppName() string {
	if x != nil {
		return x.AppName
	}
	return ""
}

func (x *History) GetAppLogo() string {
	if x != nil {
		return x.AppLogo
	}
	return ""
}

func (x *History) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

func (x *History) GetEmailAddress() string {
	if x != nil {
		return x.EmailAddress
	}
	return ""
}

func (x *History) GetPhoneNO() string {
	if x != nil {
		return x.PhoneNO
	}
	return ""
}

func (x *History) GetResource() string {
	if x != nil {
		return x.Resource
	}
	return ""
}

func (x *History) GetMethod() string {
	if x != nil {
		return x.Method
	}
	return ""
}

func (x *History) GetAllowed() bool {
	if x != nil {
		return x.Allowed
	}
	return false
}

func (x *History) GetCreatedAt() uint32 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

type Conds struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EntID    *v1.StringVal `protobuf:"bytes,10,opt,name=EntID,proto3" json:"EntID,omitempty"`
	AppID    *v1.StringVal `protobuf:"bytes,20,opt,name=AppID,proto3" json:"AppID,omitempty"`
	UserID   *v1.StringVal `protobuf:"bytes,30,opt,name=UserID,proto3" json:"UserID,omitempty"`
	Resource *v1.StringVal `protobuf:"bytes,40,opt,name=Resource,proto3" json:"Resource,omitempty"`
	Method   *v1.StringVal `protobuf:"bytes,50,opt,name=Method,proto3" json:"Method,omitempty"`
	Allowed  *v1.BoolVal   `protobuf:"bytes,60,opt,name=Allowed,proto3" json:"Allowed,omitempty"`
}

func (x *Conds) Reset() {
	*x = Conds{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_appuser_mw_v1_authing_history_history_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Conds) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Conds) ProtoMessage() {}

func (x *Conds) ProtoReflect() protoreflect.Message {
	mi := &file_npool_appuser_mw_v1_authing_history_history_proto_msgTypes[2]
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
	return file_npool_appuser_mw_v1_authing_history_history_proto_rawDescGZIP(), []int{2}
}

func (x *Conds) GetEntID() *v1.StringVal {
	if x != nil {
		return x.EntID
	}
	return nil
}

func (x *Conds) GetAppID() *v1.StringVal {
	if x != nil {
		return x.AppID
	}
	return nil
}

func (x *Conds) GetUserID() *v1.StringVal {
	if x != nil {
		return x.UserID
	}
	return nil
}

func (x *Conds) GetResource() *v1.StringVal {
	if x != nil {
		return x.Resource
	}
	return nil
}

func (x *Conds) GetMethod() *v1.StringVal {
	if x != nil {
		return x.Method
	}
	return nil
}

func (x *Conds) GetAllowed() *v1.BoolVal {
	if x != nil {
		return x.Allowed
	}
	return nil
}

type CreateHistoryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *HistoryReq `protobuf:"bytes,10,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (x *CreateHistoryRequest) Reset() {
	*x = CreateHistoryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_appuser_mw_v1_authing_history_history_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateHistoryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateHistoryRequest) ProtoMessage() {}

func (x *CreateHistoryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_appuser_mw_v1_authing_history_history_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateHistoryRequest.ProtoReflect.Descriptor instead.
func (*CreateHistoryRequest) Descriptor() ([]byte, []int) {
	return file_npool_appuser_mw_v1_authing_history_history_proto_rawDescGZIP(), []int{3}
}

func (x *CreateHistoryRequest) GetInfo() *HistoryReq {
	if x != nil {
		return x.Info
	}
	return nil
}

type CreateHistoryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *History `protobuf:"bytes,10,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (x *CreateHistoryResponse) Reset() {
	*x = CreateHistoryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_appuser_mw_v1_authing_history_history_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateHistoryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateHistoryResponse) ProtoMessage() {}

func (x *CreateHistoryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_appuser_mw_v1_authing_history_history_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateHistoryResponse.ProtoReflect.Descriptor instead.
func (*CreateHistoryResponse) Descriptor() ([]byte, []int) {
	return file_npool_appuser_mw_v1_authing_history_history_proto_rawDescGZIP(), []int{4}
}

func (x *CreateHistoryResponse) GetInfo() *History {
	if x != nil {
		return x.Info
	}
	return nil
}

type GetHistoriesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Conds  *Conds `protobuf:"bytes,10,opt,name=Conds,proto3" json:"Conds,omitempty"`
	Offset int32  `protobuf:"varint,20,opt,name=Offset,proto3" json:"Offset,omitempty"`
	Limit  int32  `protobuf:"varint,30,opt,name=Limit,proto3" json:"Limit,omitempty"`
}

func (x *GetHistoriesRequest) Reset() {
	*x = GetHistoriesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_appuser_mw_v1_authing_history_history_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetHistoriesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetHistoriesRequest) ProtoMessage() {}

func (x *GetHistoriesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_appuser_mw_v1_authing_history_history_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetHistoriesRequest.ProtoReflect.Descriptor instead.
func (*GetHistoriesRequest) Descriptor() ([]byte, []int) {
	return file_npool_appuser_mw_v1_authing_history_history_proto_rawDescGZIP(), []int{5}
}

func (x *GetHistoriesRequest) GetConds() *Conds {
	if x != nil {
		return x.Conds
	}
	return nil
}

func (x *GetHistoriesRequest) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *GetHistoriesRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type GetHistoriesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Infos []*History `protobuf:"bytes,10,rep,name=Infos,proto3" json:"Infos,omitempty"`
	Total uint32     `protobuf:"varint,20,opt,name=Total,proto3" json:"Total,omitempty"`
}

func (x *GetHistoriesResponse) Reset() {
	*x = GetHistoriesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_appuser_mw_v1_authing_history_history_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetHistoriesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetHistoriesResponse) ProtoMessage() {}

func (x *GetHistoriesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_appuser_mw_v1_authing_history_history_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetHistoriesResponse.ProtoReflect.Descriptor instead.
func (*GetHistoriesResponse) Descriptor() ([]byte, []int) {
	return file_npool_appuser_mw_v1_authing_history_history_proto_rawDescGZIP(), []int{6}
}

func (x *GetHistoriesResponse) GetInfos() []*History {
	if x != nil {
		return x.Infos
	}
	return nil
}

func (x *GetHistoriesResponse) GetTotal() uint32 {
	if x != nil {
		return x.Total
	}
	return 0
}

var File_npool_appuser_mw_v1_authing_history_history_proto protoreflect.FileDescriptor

var file_npool_appuser_mw_v1_authing_history_history_proto_rawDesc = []byte{
	0x0a, 0x31, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x61, 0x70, 0x70, 0x75, 0x73, 0x65, 0x72, 0x2f,
	0x6d, 0x77, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x2f, 0x68, 0x69,
	0x73, 0x74, 0x6f, 0x72, 0x79, 0x2f, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x25, 0x61, 0x70, 0x70, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x6d, 0x69, 0x64,
	0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x2e,
	0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x1a, 0x1e, 0x6e, 0x70, 0x6f, 0x6f,
	0x6c, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63,
	0x6f, 0x6e, 0x64, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9b, 0x02, 0x0a, 0x0a, 0x48,
	0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x12, 0x13, 0x0a, 0x02, 0x49, 0x44, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x0d, 0x48, 0x00, 0x52, 0x02, 0x49, 0x44, 0x88, 0x01, 0x01, 0x12, 0x19,
	0x0a, 0x05, 0x45, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52,
	0x05, 0x45, 0x6e, 0x74, 0x49, 0x44, 0x88, 0x01, 0x01, 0x12, 0x19, 0x0a, 0x05, 0x41, 0x70, 0x70,
	0x49, 0x44, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x48, 0x02, 0x52, 0x05, 0x41, 0x70, 0x70, 0x49,
	0x44, 0x88, 0x01, 0x01, 0x12, 0x1b, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x1e,
	0x20, 0x01, 0x28, 0x09, 0x48, 0x03, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x88, 0x01,
	0x01, 0x12, 0x1f, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x28, 0x20,
	0x01, 0x28, 0x09, 0x48, 0x04, 0x52, 0x08, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x88,
	0x01, 0x01, 0x12, 0x1b, 0x0a, 0x06, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x32, 0x20, 0x01,
	0x28, 0x09, 0x48, 0x05, 0x52, 0x06, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x88, 0x01, 0x01, 0x12,
	0x1d, 0x0a, 0x07, 0x41, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x18, 0x3c, 0x20, 0x01, 0x28, 0x08,
	0x48, 0x06, 0x52, 0x07, 0x41, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x88, 0x01, 0x01, 0x42, 0x05,
	0x0a, 0x03, 0x5f, 0x49, 0x44, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x45, 0x6e, 0x74, 0x49, 0x44, 0x42,
	0x08, 0x0a, 0x06, 0x5f, 0x41, 0x70, 0x70, 0x49, 0x44, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x55, 0x73,
	0x65, 0x72, 0x49, 0x44, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x42, 0x0a, 0x0a, 0x08,
	0x5f, 0x41, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x22, 0x95, 0x02, 0x0a, 0x07, 0x48, 0x69, 0x73,
	0x74, 0x6f, 0x72, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x41, 0x70, 0x70, 0x49, 0x44, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x41, 0x70, 0x70, 0x49, 0x44, 0x12, 0x18, 0x0a, 0x07, 0x41, 0x70,
	0x70, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x41, 0x70, 0x70,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x41, 0x70, 0x70, 0x4c, 0x6f, 0x67, 0x6f, 0x18,
	0x1e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x41, 0x70, 0x70, 0x4c, 0x6f, 0x67, 0x6f, 0x12, 0x16,
	0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x28, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x22, 0x0a, 0x0c, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x41,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x32, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x45, 0x6d,
	0x61, 0x69, 0x6c, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x50, 0x68,
	0x6f, 0x6e, 0x65, 0x4e, 0x4f, 0x18, 0x3c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x50, 0x68, 0x6f,
	0x6e, 0x65, 0x4e, 0x4f, 0x12, 0x1a, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x18, 0x46, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x50, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x41, 0x6c, 0x6c, 0x6f,
	0x77, 0x65, 0x64, 0x18, 0x5a, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x41, 0x6c, 0x6c, 0x6f, 0x77,
	0x65, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18,
	0x64, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x22, 0xad, 0x02, 0x0a, 0x05, 0x43, 0x6f, 0x6e, 0x64, 0x73, 0x12, 0x2d, 0x0a, 0x05, 0x45, 0x6e,
	0x74, 0x49, 0x44, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x62, 0x61, 0x73, 0x65,
	0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56,
	0x61, 0x6c, 0x52, 0x05, 0x45, 0x6e, 0x74, 0x49, 0x44, 0x12, 0x2d, 0x0a, 0x05, 0x41, 0x70, 0x70,
	0x49, 0x44, 0x18, 0x14, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61,
	0x6c, 0x52, 0x05, 0x41, 0x70, 0x70, 0x49, 0x44, 0x12, 0x2f, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72,
	0x49, 0x44, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61,
	0x6c, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x33, 0x0a, 0x08, 0x52, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x28, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x62, 0x61,
	0x73, 0x65, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e,
	0x67, 0x56, 0x61, 0x6c, 0x52, 0x08, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x2f,
	0x0a, 0x06, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x32, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17,
	0x2e, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74,
	0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x52, 0x06, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12,
	0x2f, 0x0a, 0x07, 0x41, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x18, 0x3c, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x15, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x42, 0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x52, 0x07, 0x41, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64,
	0x22, 0x5d, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72,
	0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x45, 0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x31, 0x2e, 0x61, 0x70, 0x70, 0x75, 0x73, 0x65, 0x72,
	0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x61, 0x75, 0x74, 0x68,
	0x69, 0x6e, 0x67, 0x2e, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x48,
	0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x52, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x22,
	0x5b, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x42, 0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x61, 0x70, 0x70, 0x75, 0x73, 0x65, 0x72,
	0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x61, 0x75, 0x74, 0x68,
	0x69, 0x6e, 0x67, 0x2e, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x48,
	0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x87, 0x01, 0x0a,
	0x13, 0x47, 0x65, 0x74, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x42, 0x0a, 0x05, 0x43, 0x6f, 0x6e, 0x64, 0x73, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x61, 0x70, 0x70, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x6d, 0x69,
	0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x69, 0x6e, 0x67,
	0x2e, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x64,
	0x73, 0x52, 0x05, 0x43, 0x6f, 0x6e, 0x64, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x4f, 0x66, 0x66, 0x73,
	0x65, 0x74, 0x18, 0x14, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74,
	0x12, 0x14, 0x0a, 0x05, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x05, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x22, 0x72, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x48, 0x69, 0x73,
	0x74, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x44,
	0x0a, 0x05, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2e, 0x2e,
	0x61, 0x70, 0x70, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61,
	0x72, 0x65, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x2e, 0x68, 0x69, 0x73, 0x74, 0x6f,
	0x72, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x05, 0x49,
	0x6e, 0x66, 0x6f, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x14, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x32, 0xa7, 0x02, 0x0a, 0x0a, 0x4d,
	0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x12, 0x8c, 0x01, 0x0a, 0x0d, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x3b, 0x2e, 0x61, 0x70,
	0x70, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65,
	0x2e, 0x61, 0x75, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x2e, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79,
	0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72,
	0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3c, 0x2e, 0x61, 0x70, 0x70, 0x75, 0x73,
	0x65, 0x72, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x61, 0x75,
	0x74, 0x68, 0x69, 0x6e, 0x67, 0x2e, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x76, 0x31,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x89, 0x01, 0x0a, 0x0c, 0x47, 0x65, 0x74,
	0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x12, 0x3a, 0x2e, 0x61, 0x70, 0x70, 0x75,
	0x73, 0x65, 0x72, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x61,
	0x75, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x2e, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x76,
	0x31, 0x2e, 0x47, 0x65, 0x74, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3b, 0x2e, 0x61, 0x70, 0x70, 0x75, 0x73, 0x65, 0x72, 0x2e,
	0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x69,
	0x6e, 0x67, 0x2e, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65,
	0x74, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x42, 0x46, 0x5a, 0x44, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x4e, 0x70, 0x6f, 0x6f, 0x6c, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d,
	0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x61,
	0x70, 0x70, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x6d, 0x77, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x75, 0x74,
	0x68, 0x69, 0x6e, 0x67, 0x2f, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_npool_appuser_mw_v1_authing_history_history_proto_rawDescOnce sync.Once
	file_npool_appuser_mw_v1_authing_history_history_proto_rawDescData = file_npool_appuser_mw_v1_authing_history_history_proto_rawDesc
)

func file_npool_appuser_mw_v1_authing_history_history_proto_rawDescGZIP() []byte {
	file_npool_appuser_mw_v1_authing_history_history_proto_rawDescOnce.Do(func() {
		file_npool_appuser_mw_v1_authing_history_history_proto_rawDescData = protoimpl.X.CompressGZIP(file_npool_appuser_mw_v1_authing_history_history_proto_rawDescData)
	})
	return file_npool_appuser_mw_v1_authing_history_history_proto_rawDescData
}

var file_npool_appuser_mw_v1_authing_history_history_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_npool_appuser_mw_v1_authing_history_history_proto_goTypes = []interface{}{
	(*HistoryReq)(nil),            // 0: appuser.middleware.authing.history.v1.HistoryReq
	(*History)(nil),               // 1: appuser.middleware.authing.history.v1.History
	(*Conds)(nil),                 // 2: appuser.middleware.authing.history.v1.Conds
	(*CreateHistoryRequest)(nil),  // 3: appuser.middleware.authing.history.v1.CreateHistoryRequest
	(*CreateHistoryResponse)(nil), // 4: appuser.middleware.authing.history.v1.CreateHistoryResponse
	(*GetHistoriesRequest)(nil),   // 5: appuser.middleware.authing.history.v1.GetHistoriesRequest
	(*GetHistoriesResponse)(nil),  // 6: appuser.middleware.authing.history.v1.GetHistoriesResponse
	(*v1.StringVal)(nil),          // 7: basetypes.v1.StringVal
	(*v1.BoolVal)(nil),            // 8: basetypes.v1.BoolVal
}
var file_npool_appuser_mw_v1_authing_history_history_proto_depIdxs = []int32{
	7,  // 0: appuser.middleware.authing.history.v1.Conds.EntID:type_name -> basetypes.v1.StringVal
	7,  // 1: appuser.middleware.authing.history.v1.Conds.AppID:type_name -> basetypes.v1.StringVal
	7,  // 2: appuser.middleware.authing.history.v1.Conds.UserID:type_name -> basetypes.v1.StringVal
	7,  // 3: appuser.middleware.authing.history.v1.Conds.Resource:type_name -> basetypes.v1.StringVal
	7,  // 4: appuser.middleware.authing.history.v1.Conds.Method:type_name -> basetypes.v1.StringVal
	8,  // 5: appuser.middleware.authing.history.v1.Conds.Allowed:type_name -> basetypes.v1.BoolVal
	0,  // 6: appuser.middleware.authing.history.v1.CreateHistoryRequest.Info:type_name -> appuser.middleware.authing.history.v1.HistoryReq
	1,  // 7: appuser.middleware.authing.history.v1.CreateHistoryResponse.Info:type_name -> appuser.middleware.authing.history.v1.History
	2,  // 8: appuser.middleware.authing.history.v1.GetHistoriesRequest.Conds:type_name -> appuser.middleware.authing.history.v1.Conds
	1,  // 9: appuser.middleware.authing.history.v1.GetHistoriesResponse.Infos:type_name -> appuser.middleware.authing.history.v1.History
	3,  // 10: appuser.middleware.authing.history.v1.Middleware.CreateHistory:input_type -> appuser.middleware.authing.history.v1.CreateHistoryRequest
	5,  // 11: appuser.middleware.authing.history.v1.Middleware.GetHistories:input_type -> appuser.middleware.authing.history.v1.GetHistoriesRequest
	4,  // 12: appuser.middleware.authing.history.v1.Middleware.CreateHistory:output_type -> appuser.middleware.authing.history.v1.CreateHistoryResponse
	6,  // 13: appuser.middleware.authing.history.v1.Middleware.GetHistories:output_type -> appuser.middleware.authing.history.v1.GetHistoriesResponse
	12, // [12:14] is the sub-list for method output_type
	10, // [10:12] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_npool_appuser_mw_v1_authing_history_history_proto_init() }
func file_npool_appuser_mw_v1_authing_history_history_proto_init() {
	if File_npool_appuser_mw_v1_authing_history_history_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_npool_appuser_mw_v1_authing_history_history_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HistoryReq); i {
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
		file_npool_appuser_mw_v1_authing_history_history_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*History); i {
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
		file_npool_appuser_mw_v1_authing_history_history_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
		file_npool_appuser_mw_v1_authing_history_history_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateHistoryRequest); i {
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
		file_npool_appuser_mw_v1_authing_history_history_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateHistoryResponse); i {
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
		file_npool_appuser_mw_v1_authing_history_history_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetHistoriesRequest); i {
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
		file_npool_appuser_mw_v1_authing_history_history_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetHistoriesResponse); i {
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
	file_npool_appuser_mw_v1_authing_history_history_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_npool_appuser_mw_v1_authing_history_history_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_npool_appuser_mw_v1_authing_history_history_proto_goTypes,
		DependencyIndexes: file_npool_appuser_mw_v1_authing_history_history_proto_depIdxs,
		MessageInfos:      file_npool_appuser_mw_v1_authing_history_history_proto_msgTypes,
	}.Build()
	File_npool_appuser_mw_v1_authing_history_history_proto = out.File
	file_npool_appuser_mw_v1_authing_history_history_proto_rawDesc = nil
	file_npool_appuser_mw_v1_authing_history_history_proto_goTypes = nil
	file_npool_appuser_mw_v1_authing_history_history_proto_depIdxs = nil
}
