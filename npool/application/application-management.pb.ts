/* eslint-disable */
// @ts-nocheck
/*
* This file is a generated Typescript file for GRPC Gateway, DO NOT MODIFY
*/

import * as fm from "../../fetch.pb"
import * as GoogleProtobufEmpty from "../../google/protobuf/empty.pb"
export type VersionResponse = {
  Info?: string
}

export type PageInfo = {
  PageIndex?: number
  PageSize?: number
}

export type ApplicationInfo = {
  ID?: string
  ApplicationName?: string
  ApplicationOwner?: string
  HomepageUrl?: string
  RedirectUrl?: string
  ApplicationLogo?: string
  CreateAT?: number
  UpdateAT?: number
  ClientSecret?: string
  GoogleRecaptcha?: boolean
  SmsLogin?: boolean
  InvitationCodeMust?: boolean
}

export type CreateApplicationRequest = {
  Info?: ApplicationInfo
}

export type CreateApplicationResponse = {
  Info?: ApplicationInfo
}

export type UpdateApplicationRequest = {
  Info?: ApplicationInfo
}

export type UpdateApplicationResponse = {
  Info?: ApplicationInfo
}

export type GetApplicationRequest = {
  AppID?: string
}

export type GetApplicationResponse = {
  Info?: ApplicationInfo
}

export type GetApplicationByOwnerRequest = {
  Owner?: string
}

export type OwnerApplication = {
  Owner?: string
  Infos?: ApplicationInfo[]
}

export type GetApplicationByOwnerResponse = {
  Info?: OwnerApplication
}

export type GetApplicationsRequest = {
  Info?: PageInfo
}

export type GetApplicationsResponse = {
  Infos?: ApplicationInfo[]
}

export type DeleteApplicationRequest = {
  AppID?: string
}

export type DeleteApplicationResponse = {
  Info?: string
}

export type RoleInfo = {
  ID?: string
  AppID?: string
  RoleName?: string
  Creator?: string
  CreateAT?: number
  UpdateAT?: number
  Annotation?: string
}

export type CreateRoleRequest = {
  Info?: RoleInfo
}

export type CreateRoleResponse = {
  Info?: RoleInfo
}

export type UpdateRoleRequest = {
  Info?: RoleInfo
}

export type UpdateRoleResponse = {
  Info?: RoleInfo
}

export type GetRoleRequest = {
  AppID?: string
  RoleID?: string
}

export type GetRoleResponse = {
  Info?: RoleInfo
}

export type GetRoleByCreatorRequest = {
  AppID?: string
  Creator?: string
}

export type CreatorRole = {
  AppID?: string
  Creator?: string
  Infos?: RoleInfo[]
}

export type GetRoleByCreatorResponse = {
  Info?: CreatorRole
}

export type DeleteRoleRequest = {
  RoleID?: string
  AppID?: string
}

export type DeleteRoleResponse = {
  Info?: string
}

export type GetRolesRequest = {
  AppID?: string
}

export type GetRolesResponse = {
  Infos?: RoleInfo[]
}

export type RoleUserInfo = {
  ID?: string
  AppID?: string
  RoleID?: string
  UserID?: string
  CreateAT?: number
}

export type SetUserRoleRequest = {
  UserIDs?: string[]
  RoleID?: string
  AppID?: string
}

export type SetUserRoleResponse = {
  Infos?: RoleUserInfo[]
}

export type GetUserRoleRequest = {
  UserID?: string
  AppID?: string
}

export type UserRole = {
  Infos?: RoleInfo[]
  UserID?: string
  AppID?: string
}

export type GetUserRoleResponse = {
  Info?: UserRole
}

export type GetRoleUsersRequest = {
  AppID?: string
  RoleID?: string
}

export type GetRoleUsersResponse = {
  Infos?: RoleUserInfo[]
}

export type UnSetUserRoleRequest = {
  UserIDs?: string[]
  RoleID?: string
  AppID?: string
}

export type UnSetUserRoleResponse = {
  Info?: string
}

export type ApplicationUserInfo = {
  ID?: string
  AppID?: string
  UserID?: string
  Original?: boolean
  KycVerify?: boolean
  GAVerify?: boolean
  GALogin?: boolean
  SMSLogin?: boolean
  LoginNumber?: number
  CreateAT?: number
}

export type ApplicationUserDetail = {
  UserApplicationInfo?: ApplicationUserInfo
  UserGroupInfos?: GroupUserInfo[]
  UserRoleInfo?: UserRole
}

export type GetApplicationUserDetailRequest = {
  AppID?: string
  UserID?: string
}

export type GetApplicationUserDetailResponse = {
  Info?: ApplicationUserDetail
}

export type AddUsersToApplicationRequest = {
  UserIDs?: string[]
  AppID?: string
  Original?: boolean
}

export type AddUsersToApplicationResponse = {
  Infos?: ApplicationUserInfo[]
}

export type GetUserFromApplicationRequest = {
  AppID?: string
  UserID?: string
}

export type GetUserFromApplicationResponse = {
  Info?: ApplicationUserInfo
}

export type GetUsersFromApplicationRequest = {
  AppID?: string
}

export type GetUsersFromApplicationResponse = {
  Infos?: ApplicationUserInfo[]
}

export type RemoveUsersFromApplicationRequest = {
  UserIDs?: string[]
  AppID?: string
}

export type RemoveUsersFromApplicationResponse = {
  Info?: string
}

export type GroupInfo = {
  ID?: string
  AppID?: string
  GroupName?: string
  GroupOwner?: string
  GroupLogo?: string
  Annotation?: string
  CreateAT?: number
  UpdateAT?: number
}

export type CreateGroupRequest = {
  Info?: GroupInfo
}

export type CreateGroupResponse = {
  Info?: GroupInfo
}

export type GetGroupRequest = {
  GroupID?: string
  AppID?: string
}

export type GetGroupResponse = {
  Info?: GroupInfo
}

export type GetGroupByOwnerRequest = {
  AppID?: string
  Owner?: string
}

export type OwnerGroup = {
  AppID?: string
  Owner?: string
  Infos?: GroupInfo[]
}

export type GetGroupByOwnerResponse = {
  Info?: OwnerGroup
}

export type GetAllGroupsRequest = {
  AppID?: string
}

export type GetAllGroupsResponse = {
  Infos?: GroupInfo[]
}

export type UpdateGroupRequest = {
  Info?: GroupInfo
}

export type UpdateGroupResponse = {
  Info?: GroupInfo
}

export type DeleteGroupRequest = {
  GroupID?: string
  AppID?: string
}

export type DeleteGroupResponse = {
  Info?: string
}

export type GroupUserInfo = {
  ID?: string
  GroupID?: string
  AppID?: string
  UserID?: string
  Annotation?: string
  CreateAT?: number
}

export type AddGroupUsersRequest = {
  UserIDs?: string[]
  GroupID?: string
  AppID?: string
}

export type AddGroupUsersResponse = {
  Infos?: GroupUserInfo[]
}

export type GetGroupUsersRequest = {
  GroupID?: string
  AppID?: string
}

export type GetGroupUsersResponse = {
  Infos?: GroupUserInfo[]
}

export type RemoveGroupUsersRequest = {
  UserIDs?: string[]
  GroupID?: string
  AppID?: string
}

export type RemoveGroupUsersResponse = {
  Info?: string
}

export type ResourceInfo = {
  ID?: string
  AppID?: string
  ResourceName?: string
  ResourceDescription?: string
  Type?: string
  Creator?: string
  CreateAT?: number
  UpdateAT?: number
}

export type CreateResourceRequest = {
  Info?: ResourceInfo
}

export type CreateResourceResponse = {
  Info?: ResourceInfo
}

export type UpdateResourceRequest = {
  Info?: ResourceInfo
}

export type UpdateResourceResponse = {
  Info?: ResourceInfo
}

export type GetResourceRequest = {
  ResourceID?: string
  AppID?: string
}

export type GetResourceResponse = {
  Info?: ResourceInfo
}

export type GetResourceByCreatorRequest = {
  AppID?: string
  Creator?: string
}

export type CreatorResource = {
  AppID?: string
  Creator?: string
  Infos?: ResourceInfo[]
}

export type GetResourceByCreatorResponse = {
  Info?: CreatorResource
}

export type GetResourcesRequest = {
  AppID?: string
}

export type GetResourcesResponse = {
  Infos?: ResourceInfo[]
}

export type DeleteResourceRequest = {
  ResourceID?: string
  AppID?: string
}

export type DeleteResourceResponse = {
  Info?: string
}

export type SetGALoginRequest = {
  AppID?: string
  UserID?: string
  Set?: boolean
}

export type SetGALoginResponse = {
  Info?: string
}

export type AddUserLoginTimeRequest = {
  UserID?: string
  AppID?: string
}

export type AddUserLoginTimeResponse = {
  Info?: number
}

export type UpdateUserGAStatusRequest = {
  UserID?: string
  AppID?: string
  Status?: boolean
}

export type UpdateUserGAStatusResponse = {
  Info?: string
}

export type UpdateUserKYCStatusRequest = {
  UserID?: string
  AppID?: string
  Status?: boolean
}

export type UpdateUserKYCStatusResponse = {
  Info?: string
}

export type GetUserGroupRequest = {
  AppID?: string
  UserID?: string
}

export type GetUserGroupResponse = {
  Infos?: GroupUserInfo[]
}

export type SetSMSLoginRequest = {
  AppID?: string
  UserID?: string
  Set?: boolean
}

export type SetSMSLoginResponse = {
  Info?: string
}

export type GetResourceByNameRequest = {
  ResourceName?: string
  AppID?: string
}

export type GetResourceByNameResponse = {
  Info?: ResourceInfo
}

export type GetUserAppIDRequest = {
  UserID?: string
}

export type GetUserAppIDResponse = {
  Infos?: string[]
}

export class ApplicationManagement {
  static Version(req: GoogleProtobufEmpty.Empty, initReq?: fm.InitReq): Promise<VersionResponse> {
    return fm.fetchReq<GoogleProtobufEmpty.Empty, VersionResponse>(`/version`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static CreateApplication(req: CreateApplicationRequest, initReq?: fm.InitReq): Promise<CreateApplicationResponse> {
    return fm.fetchReq<CreateApplicationRequest, CreateApplicationResponse>(`/v1/create/app`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static UpdateApplication(req: UpdateApplicationRequest, initReq?: fm.InitReq): Promise<UpdateApplicationResponse> {
    return fm.fetchReq<UpdateApplicationRequest, UpdateApplicationResponse>(`/v1/update/app`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetApplication(req: GetApplicationRequest, initReq?: fm.InitReq): Promise<GetApplicationResponse> {
    return fm.fetchReq<GetApplicationRequest, GetApplicationResponse>(`/v1/get/app`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetApplicationByOwner(req: GetApplicationByOwnerRequest, initReq?: fm.InitReq): Promise<GetApplicationByOwnerResponse> {
    return fm.fetchReq<GetApplicationByOwnerRequest, GetApplicationByOwnerResponse>(`/v1/get/app/by/owner`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetApplications(req: GetApplicationsRequest, initReq?: fm.InitReq): Promise<GetApplicationsResponse> {
    return fm.fetchReq<GetApplicationsRequest, GetApplicationsResponse>(`/v1/get/apps`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static DeleteApplication(req: DeleteApplicationRequest, initReq?: fm.InitReq): Promise<DeleteApplicationResponse> {
    return fm.fetchReq<DeleteApplicationRequest, DeleteApplicationResponse>(`/v1/delete/app`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static CreateRole(req: CreateRoleRequest, initReq?: fm.InitReq): Promise<CreateRoleResponse> {
    return fm.fetchReq<CreateRoleRequest, CreateRoleResponse>(`/v1/create/role`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static UpdateRole(req: UpdateRoleRequest, initReq?: fm.InitReq): Promise<UpdateRoleResponse> {
    return fm.fetchReq<UpdateRoleRequest, UpdateRoleResponse>(`/v1/update/role`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetRole(req: GetRoleRequest, initReq?: fm.InitReq): Promise<GetRoleResponse> {
    return fm.fetchReq<GetRoleRequest, GetRoleResponse>(`/v1/get/role`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetRoleByCreator(req: GetRoleByCreatorRequest, initReq?: fm.InitReq): Promise<GetRoleByCreatorResponse> {
    return fm.fetchReq<GetRoleByCreatorRequest, GetRoleByCreatorResponse>(`/v1/get/role/by/creator`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetRoles(req: GetRolesRequest, initReq?: fm.InitReq): Promise<GetRolesResponse> {
    return fm.fetchReq<GetRolesRequest, GetRolesResponse>(`/v1/get/roles`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static DeleteRole(req: DeleteRoleRequest, initReq?: fm.InitReq): Promise<DeleteRoleResponse> {
    return fm.fetchReq<DeleteRoleRequest, DeleteRoleResponse>(`/v1/delete/role`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static SetUserRole(req: SetUserRoleRequest, initReq?: fm.InitReq): Promise<SetUserRoleResponse> {
    return fm.fetchReq<SetUserRoleRequest, SetUserRoleResponse>(`/v1/set/user/role`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetUserRole(req: GetUserRoleRequest, initReq?: fm.InitReq): Promise<GetUserRoleResponse> {
    return fm.fetchReq<GetUserRoleRequest, GetUserRoleResponse>(`/v1/get/user/role`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetRoleUsers(req: GetRoleUsersRequest, initReq?: fm.InitReq): Promise<GetRoleUsersResponse> {
    return fm.fetchReq<GetRoleUsersRequest, GetRoleUsersResponse>(`/v1/get/role/users`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static UnSetUserRole(req: UnSetUserRoleRequest, initReq?: fm.InitReq): Promise<UnSetUserRoleResponse> {
    return fm.fetchReq<UnSetUserRoleRequest, UnSetUserRoleResponse>(`/v1/unset/user/role`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static AddUsersToApplication(req: AddUsersToApplicationRequest, initReq?: fm.InitReq): Promise<AddUsersToApplicationResponse> {
    return fm.fetchReq<AddUsersToApplicationRequest, AddUsersToApplicationResponse>(`/v1/add/users/to/app`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetUserFromApplication(req: GetUserFromApplicationRequest, initReq?: fm.InitReq): Promise<GetUserFromApplicationResponse> {
    return fm.fetchReq<GetUserFromApplicationRequest, GetUserFromApplicationResponse>(`/v1/get/user/from/app`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetUsersFromApplication(req: GetUsersFromApplicationRequest, initReq?: fm.InitReq): Promise<GetUsersFromApplicationResponse> {
    return fm.fetchReq<GetUsersFromApplicationRequest, GetUsersFromApplicationResponse>(`/v1/get/users/from/app`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetUserAppID(req: GetUserAppIDRequest, initReq?: fm.InitReq): Promise<GetUserAppIDResponse> {
    return fm.fetchReq<GetUserAppIDRequest, GetUserAppIDResponse>(`/v1/get/user/appid`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static RemoveUsersFromApplication(req: RemoveUsersFromApplicationRequest, initReq?: fm.InitReq): Promise<RemoveUsersFromApplicationResponse> {
    return fm.fetchReq<RemoveUsersFromApplicationRequest, RemoveUsersFromApplicationResponse>(`/v1/remove/users/from/app`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static CreateGroup(req: CreateGroupRequest, initReq?: fm.InitReq): Promise<CreateGroupResponse> {
    return fm.fetchReq<CreateGroupRequest, CreateGroupResponse>(`/v1/create/group`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetGroup(req: GetGroupRequest, initReq?: fm.InitReq): Promise<GetGroupResponse> {
    return fm.fetchReq<GetGroupRequest, GetGroupResponse>(`/v1/get/group`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetGroupByOwner(req: GetGroupByOwnerRequest, initReq?: fm.InitReq): Promise<GetGroupByOwnerResponse> {
    return fm.fetchReq<GetGroupByOwnerRequest, GetGroupByOwnerResponse>(`/v1/get/group/by/owner`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetAllGroups(req: GetAllGroupsRequest, initReq?: fm.InitReq): Promise<GetAllGroupsResponse> {
    return fm.fetchReq<GetAllGroupsRequest, GetAllGroupsResponse>(`/v1/get/all/groups`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static UpdateGroup(req: UpdateGroupRequest, initReq?: fm.InitReq): Promise<UpdateGroupResponse> {
    return fm.fetchReq<UpdateGroupRequest, UpdateGroupResponse>(`/v1/update/group`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static DeleteGroup(req: DeleteGroupRequest, initReq?: fm.InitReq): Promise<DeleteGroupResponse> {
    return fm.fetchReq<DeleteGroupRequest, DeleteGroupResponse>(`/v1/delete/group`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static AddGroupUsers(req: AddGroupUsersRequest, initReq?: fm.InitReq): Promise<AddGroupUsersResponse> {
    return fm.fetchReq<AddGroupUsersRequest, AddGroupUsersResponse>(`/v1/add/group/users`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetGroupUsers(req: GetGroupUsersRequest, initReq?: fm.InitReq): Promise<GetGroupUsersResponse> {
    return fm.fetchReq<GetGroupUsersRequest, GetGroupUsersResponse>(`/v1/get/group/users`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetUserGroup(req: GetUserGroupRequest, initReq?: fm.InitReq): Promise<GetUserGroupResponse> {
    return fm.fetchReq<GetUserGroupRequest, GetUserGroupResponse>(`/v1/get/user/group`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static RemoveGroupUsers(req: RemoveGroupUsersRequest, initReq?: fm.InitReq): Promise<RemoveGroupUsersResponse> {
    return fm.fetchReq<RemoveGroupUsersRequest, RemoveGroupUsersResponse>(`/v1/remove/group/users`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static CreateResource(req: CreateResourceRequest, initReq?: fm.InitReq): Promise<CreateResourceResponse> {
    return fm.fetchReq<CreateResourceRequest, CreateResourceResponse>(`/v1/create/resource`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static UpdateResource(req: UpdateResourceRequest, initReq?: fm.InitReq): Promise<UpdateResourceResponse> {
    return fm.fetchReq<UpdateResourceRequest, UpdateResourceResponse>(`/v1/update/resource`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetResource(req: GetResourceRequest, initReq?: fm.InitReq): Promise<GetResourceResponse> {
    return fm.fetchReq<GetResourceRequest, GetResourceResponse>(`/v1/get/resource`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetResourceByName(req: GetResourceByNameRequest, initReq?: fm.InitReq): Promise<GetResourceByNameResponse> {
    return fm.fetchReq<GetResourceByNameRequest, GetResourceByNameResponse>(`/v1/get/resource/by/name`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetResourceByCreator(req: GetResourceByCreatorRequest, initReq?: fm.InitReq): Promise<GetResourceByCreatorResponse> {
    return fm.fetchReq<GetResourceByCreatorRequest, GetResourceByCreatorResponse>(`/v1/get/resource/by/creator`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetResources(req: GetResourcesRequest, initReq?: fm.InitReq): Promise<GetResourcesResponse> {
    return fm.fetchReq<GetResourcesRequest, GetResourcesResponse>(`/v1/get/resources`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static DeleteResource(req: DeleteResourceRequest, initReq?: fm.InitReq): Promise<DeleteResourceResponse> {
    return fm.fetchReq<DeleteResourceRequest, DeleteResourceResponse>(`/v1/delete/resource`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static SetGALogin(req: SetGALoginRequest, initReq?: fm.InitReq): Promise<SetGALoginResponse> {
    return fm.fetchReq<SetGALoginRequest, SetGALoginResponse>(`/v1/set/ga/login`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static SetSMSLogin(req: SetSMSLoginRequest, initReq?: fm.InitReq): Promise<SetSMSLoginResponse> {
    return fm.fetchReq<SetSMSLoginRequest, SetSMSLoginResponse>(`/v1/set/sms/login`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static AddUserLoginTime(req: AddUserLoginTimeRequest, initReq?: fm.InitReq): Promise<AddUserLoginTimeResponse> {
    return fm.fetchReq<AddUserLoginTimeRequest, AddUserLoginTimeResponse>(`/v1/add/user/login/time`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static UpdateUserGAStatus(req: UpdateUserGAStatusRequest, initReq?: fm.InitReq): Promise<UpdateUserGAStatusResponse> {
    return fm.fetchReq<UpdateUserGAStatusRequest, UpdateUserGAStatusResponse>(`/v1/update/user/ga/status`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static UpdateUserKYCStatus(req: UpdateUserKYCStatusRequest, initReq?: fm.InitReq): Promise<UpdateUserKYCStatusResponse> {
    return fm.fetchReq<UpdateUserKYCStatusRequest, UpdateUserKYCStatusResponse>(`/v1/update/user/kyc/status`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetApplicationUserDetail(req: GetApplicationUserDetailRequest, initReq?: fm.InitReq): Promise<GetApplicationUserDetailResponse> {
    return fm.fetchReq<GetApplicationUserDetailRequest, GetApplicationUserDetailResponse>(`/v1/get/app/user/detail`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
}