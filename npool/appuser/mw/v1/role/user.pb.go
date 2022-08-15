// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.18.1
// source: npool/appuser/mw/v1/role/user.proto

package role

import (
	approleuser "github.com/NpoolPlatform/message/npool/appuser/mgr/v2/approleuser"
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

type RoleUser struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: sql:"id"
	ID string `protobuf:"bytes,10,opt,name=ID,proto3" json:"ID,omitempty" sql:"id"`
	// @inject_tag: sql:"created_by"
	CreatedBy string `protobuf:"bytes,20,opt,name=CreatedBy,proto3" json:"CreatedBy,omitempty" sql:"created_by"`
	// @inject_tag: sql:"role"
	Role string `protobuf:"bytes,30,opt,name=Role,proto3" json:"Role,omitempty" sql:"role"`
	// @inject_tag: sql:"description"
	Description string `protobuf:"bytes,40,opt,name=Description,proto3" json:"Description,omitempty" sql:"description"`
	// @inject_tag: sql:"default"
	DefaultInt int32 `protobuf:"varint,50,opt,name=DefaultInt,proto3" json:"DefaultInt,omitempty" sql:"default"`
	Default    bool  `protobuf:"varint,60,opt,name=Default,proto3" json:"Default,omitempty"`
	// @inject_tag: sql:"app_id"
	AppID string `protobuf:"bytes,70,opt,name=AppID,proto3" json:"AppID,omitempty" sql:"app_id"`
	// @inject_tag: sql:"app_name"
	AppName string `protobuf:"bytes,80,opt,name=AppName,proto3" json:"AppName,omitempty" sql:"app_name"`
	// @inject_tag: sql:"app_logo"
	AppLogo string `protobuf:"bytes,90,opt,name=AppLogo,proto3" json:"AppLogo,omitempty" sql:"app_logo"`
	// @inject_tag: sql:"created_at"
	CreatedAt uint32 `protobuf:"varint,100,opt,name=CreatedAt,proto3" json:"CreatedAt,omitempty" sql:"created_at"`
	// @inject_tag: sql:"user_id"
	UserID string `protobuf:"bytes,110,opt,name=UserID,proto3" json:"UserID,omitempty" sql:"user_id"`
	// @inject_tag: sql:"email_address"
	EmailAddress string `protobuf:"bytes,120,opt,name=EmailAddress,proto3" json:"EmailAddress,omitempty" sql:"email_address"`
	// @inject_tag: sql:"phone_no"
	PhoneNO string `protobuf:"bytes,130,opt,name=PhoneNO,proto3" json:"PhoneNO,omitempty" sql:"phone_no"`
}

func (x *RoleUser) Reset() {
	*x = RoleUser{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_appuser_mw_v1_role_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoleUser) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoleUser) ProtoMessage() {}

func (x *RoleUser) ProtoReflect() protoreflect.Message {
	mi := &file_npool_appuser_mw_v1_role_user_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoleUser.ProtoReflect.Descriptor instead.
func (*RoleUser) Descriptor() ([]byte, []int) {
	return file_npool_appuser_mw_v1_role_user_proto_rawDescGZIP(), []int{0}
}

func (x *RoleUser) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *RoleUser) GetCreatedBy() string {
	if x != nil {
		return x.CreatedBy
	}
	return ""
}

func (x *RoleUser) GetRole() string {
	if x != nil {
		return x.Role
	}
	return ""
}

func (x *RoleUser) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *RoleUser) GetDefaultInt() int32 {
	if x != nil {
		return x.DefaultInt
	}
	return 0
}

func (x *RoleUser) GetDefault() bool {
	if x != nil {
		return x.Default
	}
	return false
}

func (x *RoleUser) GetAppID() string {
	if x != nil {
		return x.AppID
	}
	return ""
}

func (x *RoleUser) GetAppName() string {
	if x != nil {
		return x.AppName
	}
	return ""
}

func (x *RoleUser) GetAppLogo() string {
	if x != nil {
		return x.AppLogo
	}
	return ""
}

func (x *RoleUser) GetCreatedAt() uint32 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *RoleUser) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

func (x *RoleUser) GetEmailAddress() string {
	if x != nil {
		return x.EmailAddress
	}
	return ""
}

func (x *RoleUser) GetPhoneNO() string {
	if x != nil {
		return x.PhoneNO
	}
	return ""
}

type CreateRoleUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Permissioned by path
	Info *approleuser.AppRoleUserReq `protobuf:"bytes,10,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (x *CreateRoleUserRequest) Reset() {
	*x = CreateRoleUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_appuser_mw_v1_role_user_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateRoleUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRoleUserRequest) ProtoMessage() {}

func (x *CreateRoleUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_appuser_mw_v1_role_user_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRoleUserRequest.ProtoReflect.Descriptor instead.
func (*CreateRoleUserRequest) Descriptor() ([]byte, []int) {
	return file_npool_appuser_mw_v1_role_user_proto_rawDescGZIP(), []int{1}
}

func (x *CreateRoleUserRequest) GetInfo() *approleuser.AppRoleUserReq {
	if x != nil {
		return x.Info
	}
	return nil
}

type CreateRoleUserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *RoleUser `protobuf:"bytes,10,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (x *CreateRoleUserResponse) Reset() {
	*x = CreateRoleUserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_appuser_mw_v1_role_user_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateRoleUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRoleUserResponse) ProtoMessage() {}

func (x *CreateRoleUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_appuser_mw_v1_role_user_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRoleUserResponse.ProtoReflect.Descriptor instead.
func (*CreateRoleUserResponse) Descriptor() ([]byte, []int) {
	return file_npool_appuser_mw_v1_role_user_proto_rawDescGZIP(), []int{2}
}

func (x *CreateRoleUserResponse) GetInfo() *RoleUser {
	if x != nil {
		return x.Info
	}
	return nil
}

type DeleteRoleUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID string `protobuf:"bytes,10,opt,name=ID,proto3" json:"ID,omitempty"`
}

func (x *DeleteRoleUserRequest) Reset() {
	*x = DeleteRoleUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_appuser_mw_v1_role_user_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteRoleUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteRoleUserRequest) ProtoMessage() {}

func (x *DeleteRoleUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_appuser_mw_v1_role_user_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteRoleUserRequest.ProtoReflect.Descriptor instead.
func (*DeleteRoleUserRequest) Descriptor() ([]byte, []int) {
	return file_npool_appuser_mw_v1_role_user_proto_rawDescGZIP(), []int{3}
}

func (x *DeleteRoleUserRequest) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

type DeleteRoleUserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *RoleUser `protobuf:"bytes,10,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (x *DeleteRoleUserResponse) Reset() {
	*x = DeleteRoleUserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_appuser_mw_v1_role_user_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteRoleUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteRoleUserResponse) ProtoMessage() {}

func (x *DeleteRoleUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_appuser_mw_v1_role_user_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteRoleUserResponse.ProtoReflect.Descriptor instead.
func (*DeleteRoleUserResponse) Descriptor() ([]byte, []int) {
	return file_npool_appuser_mw_v1_role_user_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteRoleUserResponse) GetInfo() *RoleUser {
	if x != nil {
		return x.Info
	}
	return nil
}

type GetRoleUsersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Permissioned by path
	AppID  string `protobuf:"bytes,10,opt,name=AppID,proto3" json:"AppID,omitempty"`
	Offset int32  `protobuf:"varint,20,opt,name=Offset,proto3" json:"Offset,omitempty"`
	Limit  int32  `protobuf:"varint,30,opt,name=Limit,proto3" json:"Limit,omitempty"`
}

func (x *GetRoleUsersRequest) Reset() {
	*x = GetRoleUsersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_appuser_mw_v1_role_user_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRoleUsersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRoleUsersRequest) ProtoMessage() {}

func (x *GetRoleUsersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_appuser_mw_v1_role_user_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRoleUsersRequest.ProtoReflect.Descriptor instead.
func (*GetRoleUsersRequest) Descriptor() ([]byte, []int) {
	return file_npool_appuser_mw_v1_role_user_proto_rawDescGZIP(), []int{5}
}

func (x *GetRoleUsersRequest) GetAppID() string {
	if x != nil {
		return x.AppID
	}
	return ""
}

func (x *GetRoleUsersRequest) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *GetRoleUsersRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type GetRoleUsersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Infos []*RoleUser `protobuf:"bytes,10,rep,name=Infos,proto3" json:"Infos,omitempty"`
	Total uint32      `protobuf:"varint,20,opt,name=Total,proto3" json:"Total,omitempty"`
}

func (x *GetRoleUsersResponse) Reset() {
	*x = GetRoleUsersResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_appuser_mw_v1_role_user_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRoleUsersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRoleUsersResponse) ProtoMessage() {}

func (x *GetRoleUsersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_appuser_mw_v1_role_user_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRoleUsersResponse.ProtoReflect.Descriptor instead.
func (*GetRoleUsersResponse) Descriptor() ([]byte, []int) {
	return file_npool_appuser_mw_v1_role_user_proto_rawDescGZIP(), []int{6}
}

func (x *GetRoleUsersResponse) GetInfos() []*RoleUser {
	if x != nil {
		return x.Infos
	}
	return nil
}

func (x *GetRoleUsersResponse) GetTotal() uint32 {
	if x != nil {
		return x.Total
	}
	return 0
}

type GetAppRoleUsersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Permissioned by path
	TargetAppID string `protobuf:"bytes,10,opt,name=TargetAppID,proto3" json:"TargetAppID,omitempty"`
	Offset      int32  `protobuf:"varint,20,opt,name=Offset,proto3" json:"Offset,omitempty"`
	Limit       int32  `protobuf:"varint,30,opt,name=Limit,proto3" json:"Limit,omitempty"`
}

func (x *GetAppRoleUsersRequest) Reset() {
	*x = GetAppRoleUsersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_appuser_mw_v1_role_user_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAppRoleUsersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAppRoleUsersRequest) ProtoMessage() {}

func (x *GetAppRoleUsersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_npool_appuser_mw_v1_role_user_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAppRoleUsersRequest.ProtoReflect.Descriptor instead.
func (*GetAppRoleUsersRequest) Descriptor() ([]byte, []int) {
	return file_npool_appuser_mw_v1_role_user_proto_rawDescGZIP(), []int{7}
}

func (x *GetAppRoleUsersRequest) GetTargetAppID() string {
	if x != nil {
		return x.TargetAppID
	}
	return ""
}

func (x *GetAppRoleUsersRequest) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *GetAppRoleUsersRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type GetAppRoleUsersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Infos []*RoleUser `protobuf:"bytes,10,rep,name=Infos,proto3" json:"Infos,omitempty"`
	Total uint32      `protobuf:"varint,20,opt,name=Total,proto3" json:"Total,omitempty"`
}

func (x *GetAppRoleUsersResponse) Reset() {
	*x = GetAppRoleUsersResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_npool_appuser_mw_v1_role_user_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAppRoleUsersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAppRoleUsersResponse) ProtoMessage() {}

func (x *GetAppRoleUsersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_npool_appuser_mw_v1_role_user_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAppRoleUsersResponse.ProtoReflect.Descriptor instead.
func (*GetAppRoleUsersResponse) Descriptor() ([]byte, []int) {
	return file_npool_appuser_mw_v1_role_user_proto_rawDescGZIP(), []int{8}
}

func (x *GetAppRoleUsersResponse) GetInfos() []*RoleUser {
	if x != nil {
		return x.Infos
	}
	return nil
}

func (x *GetAppRoleUsersResponse) GetTotal() uint32 {
	if x != nil {
		return x.Total
	}
	return 0
}

var File_npool_appuser_mw_v1_role_user_proto protoreflect.FileDescriptor

var file_npool_appuser_mw_v1_role_user_proto_rawDesc = []byte{
	0x0a, 0x23, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x61, 0x70, 0x70, 0x75, 0x73, 0x65, 0x72, 0x2f,
	0x6d, 0x77, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x6f, 0x6c, 0x65, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1a, 0x61, 0x70, 0x70, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x6d,
	0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x76,
	0x31, 0x1a, 0x32, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x61, 0x70, 0x70, 0x75, 0x73, 0x65, 0x72,
	0x2f, 0x6d, 0x67, 0x72, 0x2f, 0x76, 0x32, 0x2f, 0x61, 0x70, 0x70, 0x72, 0x6f, 0x6c, 0x65, 0x75,
	0x73, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x70, 0x72, 0x6f, 0x6c, 0x65, 0x75, 0x73, 0x65, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe7, 0x02, 0x0a, 0x08, 0x52, 0x6f, 0x6c, 0x65, 0x55, 0x73,
	0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x49, 0x44, 0x12, 0x1c, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x18,
	0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79,
	0x12, 0x12, 0x0a, 0x04, 0x52, 0x6f, 0x6c, 0x65, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x52, 0x6f, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x28, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1e, 0x0a, 0x0a, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c,
	0x74, 0x49, 0x6e, 0x74, 0x18, 0x32, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x44, 0x65, 0x66, 0x61,
	0x75, 0x6c, 0x74, 0x49, 0x6e, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c,
	0x74, 0x18, 0x3c, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74,
	0x12, 0x14, 0x0a, 0x05, 0x41, 0x70, 0x70, 0x49, 0x44, 0x18, 0x46, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x41, 0x70, 0x70, 0x49, 0x44, 0x12, 0x18, 0x0a, 0x07, 0x41, 0x70, 0x70, 0x4e, 0x61, 0x6d,
	0x65, 0x18, 0x50, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x41, 0x70, 0x70, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x41, 0x70, 0x70, 0x4c, 0x6f, 0x67, 0x6f, 0x18, 0x5a, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x41, 0x70, 0x70, 0x4c, 0x6f, 0x67, 0x6f, 0x12, 0x1c, 0x0a, 0x09, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x64, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72,
	0x49, 0x44, 0x18, 0x6e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44,
	0x12, 0x22, 0x0a, 0x0c, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x18, 0x78, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x41, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x12, 0x19, 0x0a, 0x07, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x4f, 0x18,
	0x82, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x4f, 0x22,
	0x5b, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x55, 0x73, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x42, 0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x61, 0x70, 0x70, 0x75, 0x73, 0x65, 0x72,
	0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x70, 0x72, 0x6f, 0x6c, 0x65,
	0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x32, 0x2e, 0x41, 0x70, 0x70, 0x52, 0x6f, 0x6c, 0x65, 0x55,
	0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x52, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x52, 0x0a, 0x16,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x38, 0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x61, 0x70, 0x70, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x6d,
	0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x49, 0x6e, 0x66, 0x6f,
	0x22, 0x27, 0x0a, 0x15, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x55, 0x73,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x22, 0x52, 0x0a, 0x16, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x38, 0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x24, 0x2e, 0x61, 0x70, 0x70, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x6d, 0x69, 0x64, 0x64,
	0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x52,
	0x6f, 0x6c, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x59, 0x0a,
	0x13, 0x47, 0x65, 0x74, 0x52, 0x6f, 0x6c, 0x65, 0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x41, 0x70, 0x70, 0x49, 0x44, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x41, 0x70, 0x70, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x4f, 0x66,
	0x66, 0x73, 0x65, 0x74, 0x18, 0x14, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x4f, 0x66, 0x66, 0x73,
	0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x1e, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x05, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x22, 0x68, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x52,
	0x6f, 0x6c, 0x65, 0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x3a, 0x0a, 0x05, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x24, 0x2e, 0x61, 0x70, 0x70, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65,
	0x77, 0x61, 0x72, 0x65, 0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x6f, 0x6c,
	0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x05, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x12, 0x14, 0x0a, 0x05,
	0x54, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x14, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x54, 0x6f, 0x74,
	0x61, 0x6c, 0x22, 0x68, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x41, 0x70, 0x70, 0x52, 0x6f, 0x6c, 0x65,
	0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x0b,
	0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x41, 0x70, 0x70, 0x49, 0x44, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x41, 0x70, 0x70, 0x49, 0x44, 0x12, 0x16,
	0x0a, 0x06, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x14, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06,
	0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x18,
	0x1e, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x22, 0x6b, 0x0a, 0x17,
	0x47, 0x65, 0x74, 0x41, 0x70, 0x70, 0x52, 0x6f, 0x6c, 0x65, 0x55, 0x73, 0x65, 0x72, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3a, 0x0a, 0x05, 0x49, 0x6e, 0x66, 0x6f, 0x73,
	0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x61, 0x70, 0x70, 0x75, 0x73, 0x65, 0x72,
	0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x72, 0x6f, 0x6c, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x05, 0x49, 0x6e,
	0x66, 0x6f, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x14, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x32, 0xf5, 0x03, 0x0a, 0x0a, 0x4d, 0x69,
	0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x12, 0x79, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x31, 0x2e, 0x61, 0x70, 0x70,
	0x75, 0x73, 0x65, 0x72, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e,
	0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x6f,
	0x6c, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x32, 0x2e,
	0x61, 0x70, 0x70, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61,
	0x72, 0x65, 0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x52, 0x6f, 0x6c, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x79, 0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x6f, 0x6c,
	0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x31, 0x2e, 0x61, 0x70, 0x70, 0x75, 0x73, 0x65, 0x72, 0x2e,
	0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x55, 0x73, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x32, 0x2e, 0x61, 0x70, 0x70, 0x75, 0x73,
	0x65, 0x72, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x72, 0x6f,
	0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x6f, 0x6c, 0x65,
	0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x73,
	0x0a, 0x0c, 0x47, 0x65, 0x74, 0x52, 0x6f, 0x6c, 0x65, 0x55, 0x73, 0x65, 0x72, 0x73, 0x12, 0x2f,
	0x2e, 0x61, 0x70, 0x70, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77,
	0x61, 0x72, 0x65, 0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x52,
	0x6f, 0x6c, 0x65, 0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x30, 0x2e, 0x61, 0x70, 0x70, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65,
	0x77, 0x61, 0x72, 0x65, 0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74,
	0x52, 0x6f, 0x6c, 0x65, 0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x7c, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x41, 0x70, 0x70, 0x52, 0x6f, 0x6c,
	0x65, 0x55, 0x73, 0x65, 0x72, 0x73, 0x12, 0x32, 0x2e, 0x61, 0x70, 0x70, 0x75, 0x73, 0x65, 0x72,
	0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x72, 0x6f, 0x6c, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x70, 0x70, 0x52, 0x6f, 0x6c, 0x65, 0x55, 0x73,
	0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x33, 0x2e, 0x61, 0x70, 0x70,
	0x75, 0x73, 0x65, 0x72, 0x2e, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e,
	0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x70, 0x70, 0x52, 0x6f,
	0x6c, 0x65, 0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x42, 0x3b, 0x5a, 0x39, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x4e, 0x70, 0x6f, 0x6f, 0x6c, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x6e, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x61, 0x70, 0x70, 0x75,
	0x73, 0x65, 0x72, 0x2f, 0x6d, 0x77, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x6f, 0x6c, 0x65, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_npool_appuser_mw_v1_role_user_proto_rawDescOnce sync.Once
	file_npool_appuser_mw_v1_role_user_proto_rawDescData = file_npool_appuser_mw_v1_role_user_proto_rawDesc
)

func file_npool_appuser_mw_v1_role_user_proto_rawDescGZIP() []byte {
	file_npool_appuser_mw_v1_role_user_proto_rawDescOnce.Do(func() {
		file_npool_appuser_mw_v1_role_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_npool_appuser_mw_v1_role_user_proto_rawDescData)
	})
	return file_npool_appuser_mw_v1_role_user_proto_rawDescData
}

var file_npool_appuser_mw_v1_role_user_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_npool_appuser_mw_v1_role_user_proto_goTypes = []interface{}{
	(*RoleUser)(nil),                   // 0: appuser.middleware.role.v1.RoleUser
	(*CreateRoleUserRequest)(nil),      // 1: appuser.middleware.role.v1.CreateRoleUserRequest
	(*CreateRoleUserResponse)(nil),     // 2: appuser.middleware.role.v1.CreateRoleUserResponse
	(*DeleteRoleUserRequest)(nil),      // 3: appuser.middleware.role.v1.DeleteRoleUserRequest
	(*DeleteRoleUserResponse)(nil),     // 4: appuser.middleware.role.v1.DeleteRoleUserResponse
	(*GetRoleUsersRequest)(nil),        // 5: appuser.middleware.role.v1.GetRoleUsersRequest
	(*GetRoleUsersResponse)(nil),       // 6: appuser.middleware.role.v1.GetRoleUsersResponse
	(*GetAppRoleUsersRequest)(nil),     // 7: appuser.middleware.role.v1.GetAppRoleUsersRequest
	(*GetAppRoleUsersResponse)(nil),    // 8: appuser.middleware.role.v1.GetAppRoleUsersResponse
	(*approleuser.AppRoleUserReq)(nil), // 9: appuser.manager.approleuser.v2.AppRoleUserReq
}
var file_npool_appuser_mw_v1_role_user_proto_depIdxs = []int32{
	9, // 0: appuser.middleware.role.v1.CreateRoleUserRequest.Info:type_name -> appuser.manager.approleuser.v2.AppRoleUserReq
	0, // 1: appuser.middleware.role.v1.CreateRoleUserResponse.Info:type_name -> appuser.middleware.role.v1.RoleUser
	0, // 2: appuser.middleware.role.v1.DeleteRoleUserResponse.Info:type_name -> appuser.middleware.role.v1.RoleUser
	0, // 3: appuser.middleware.role.v1.GetRoleUsersResponse.Infos:type_name -> appuser.middleware.role.v1.RoleUser
	0, // 4: appuser.middleware.role.v1.GetAppRoleUsersResponse.Infos:type_name -> appuser.middleware.role.v1.RoleUser
	1, // 5: appuser.middleware.role.v1.Middleware.CreateRoleUser:input_type -> appuser.middleware.role.v1.CreateRoleUserRequest
	3, // 6: appuser.middleware.role.v1.Middleware.DeleteRoleUser:input_type -> appuser.middleware.role.v1.DeleteRoleUserRequest
	5, // 7: appuser.middleware.role.v1.Middleware.GetRoleUsers:input_type -> appuser.middleware.role.v1.GetRoleUsersRequest
	7, // 8: appuser.middleware.role.v1.Middleware.GetAppRoleUsers:input_type -> appuser.middleware.role.v1.GetAppRoleUsersRequest
	2, // 9: appuser.middleware.role.v1.Middleware.CreateRoleUser:output_type -> appuser.middleware.role.v1.CreateRoleUserResponse
	4, // 10: appuser.middleware.role.v1.Middleware.DeleteRoleUser:output_type -> appuser.middleware.role.v1.DeleteRoleUserResponse
	6, // 11: appuser.middleware.role.v1.Middleware.GetRoleUsers:output_type -> appuser.middleware.role.v1.GetRoleUsersResponse
	8, // 12: appuser.middleware.role.v1.Middleware.GetAppRoleUsers:output_type -> appuser.middleware.role.v1.GetAppRoleUsersResponse
	9, // [9:13] is the sub-list for method output_type
	5, // [5:9] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_npool_appuser_mw_v1_role_user_proto_init() }
func file_npool_appuser_mw_v1_role_user_proto_init() {
	if File_npool_appuser_mw_v1_role_user_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_npool_appuser_mw_v1_role_user_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoleUser); i {
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
		file_npool_appuser_mw_v1_role_user_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateRoleUserRequest); i {
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
		file_npool_appuser_mw_v1_role_user_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateRoleUserResponse); i {
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
		file_npool_appuser_mw_v1_role_user_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteRoleUserRequest); i {
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
		file_npool_appuser_mw_v1_role_user_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteRoleUserResponse); i {
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
		file_npool_appuser_mw_v1_role_user_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRoleUsersRequest); i {
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
		file_npool_appuser_mw_v1_role_user_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRoleUsersResponse); i {
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
		file_npool_appuser_mw_v1_role_user_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAppRoleUsersRequest); i {
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
		file_npool_appuser_mw_v1_role_user_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAppRoleUsersResponse); i {
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
			RawDescriptor: file_npool_appuser_mw_v1_role_user_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_npool_appuser_mw_v1_role_user_proto_goTypes,
		DependencyIndexes: file_npool_appuser_mw_v1_role_user_proto_depIdxs,
		MessageInfos:      file_npool_appuser_mw_v1_role_user_proto_msgTypes,
	}.Build()
	File_npool_appuser_mw_v1_role_user_proto = out.File
	file_npool_appuser_mw_v1_role_user_proto_rawDesc = nil
	file_npool_appuser_mw_v1_role_user_proto_goTypes = nil
	file_npool_appuser_mw_v1_role_user_proto_depIdxs = nil
}
