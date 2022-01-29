/* eslint-disable */
// @ts-nocheck
/*
* This file is a generated Typescript file for GRPC Gateway, DO NOT MODIFY
*/

import * as fm from "../../fetch.pb"
import * as GoogleProtobufEmpty from "../../google/protobuf/empty.pb"
import * as NpoolV1Npool from "../npool.pb"
export type ApplicationInfo = {
  id?: string
  applicationName?: string
  applicationOwner?: string
  homepageUrl?: string
  redirectUrl?: string
  applicationLogo?: string
  createAT?: number
  updateAT?: number
  clientSecret?: string
  googleRecaptcha?: boolean
  smsLogin?: boolean
  invitationCodeMust?: boolean
}

export type CreateApplicationRequest = {
  info?: ApplicationInfo
}

export type CreateApplicationResponse = {
  info?: ApplicationInfo
}

export type UpdateApplicationRequest = {
  info?: ApplicationInfo
}

export type UpdateApplicationResponse = {
  info?: ApplicationInfo
}

export type GetApplicationRequest = {
  appID?: string
}

export type GetApplicationResponse = {
  info?: ApplicationInfo
}

export type GetApplicationByOwnerRequest = {
  owner?: string
}

export type OwnerApplication = {
  owner?: string
  infos?: ApplicationInfo[]
}

export type GetApplicationByOwnerResponse = {
  info?: OwnerApplication
}

export type GetApplicationsRequest = {
  info?: NpoolV1Npool.PageInfo
}

export type GetApplicationsResponse = {
  infos?: ApplicationInfo[]
}

export type DeleteApplicationRequest = {
  appID?: string
}

export type DeleteApplicationResponse = {
  info?: string
}

export type RoleInfo = {
  id?: string
  appID?: string
  roleName?: string
  creator?: string
  createAT?: number
  updateAT?: number
  annotation?: string
}

export type CreateRoleRequest = {
  info?: RoleInfo
}

export type CreateRoleResponse = {
  info?: RoleInfo
}

export type UpdateRoleRequest = {
  info?: RoleInfo
}

export type UpdateRoleResponse = {
  info?: RoleInfo
}

export type GetRoleRequest = {
  appID?: string
  roleID?: string
}

export type GetRoleResponse = {
  info?: RoleInfo
}

export type GetRoleByCreatorRequest = {
  appID?: string
  creator?: string
}

export type CreatorRole = {
  appID?: string
  creator?: string
  infos?: RoleInfo[]
}

export type GetRoleByCreatorResponse = {
  info?: CreatorRole
}

export type DeleteRoleRequest = {
  roleID?: string
  appID?: string
}

export type DeleteRoleResponse = {
  info?: string
}

export type GetRolesRequest = {
  appID?: string
}

export type GetRolesResponse = {
  infos?: RoleInfo[]
}

export type RoleUserInfo = {
  id?: string
  appID?: string
  roleID?: string
  userID?: string
  createAT?: number
}

export type SetUserRoleRequest = {
  userIDs?: string[]
  roleID?: string
  appID?: string
}

export type SetUserRoleResponse = {
  infos?: RoleUserInfo[]
}

export type GetUserRoleRequest = {
  userID?: string
  appID?: string
}

export type UserRole = {
  infos?: RoleInfo[]
  userID?: string
  appID?: string
}

export type GetUserRoleResponse = {
  info?: UserRole
}

export type GetRoleUsersRequest = {
  appID?: string
  roleID?: string
}

export type GetRoleUsersResponse = {
  infos?: RoleUserInfo[]
}

export type UnSetUserRoleRequest = {
  userIDs?: string[]
  roleID?: string
  appID?: string
}

export type UnSetUserRoleResponse = {
  info?: string
}

export type ApplicationUserInfo = {
  id?: string
  appID?: string
  userID?: string
  original?: boolean
  kycVerify?: boolean
  gAVerify?: boolean
  gALogin?: boolean
  sMSLogin?: boolean
  loginNumber?: number
  createAT?: number
}

export type ApplicationUserDetail = {
  userApplicationInfo?: ApplicationUserInfo
  userGroupInfos?: GroupUserInfo[]
  userRoleInfo?: UserRole
}

export type GetApplicationUserDetailRequest = {
  appID?: string
  userID?: string
}

export type GetApplicationUserDetailResponse = {
  info?: ApplicationUserDetail
}

export type AddUsersToApplicationRequest = {
  userIDs?: string[]
  appID?: string
  original?: boolean
}

export type AddUsersToApplicationResponse = {
  infos?: ApplicationUserInfo[]
}

export type GetUserFromApplicationRequest = {
  appID?: string
  userID?: string
}

export type GetUserFromApplicationResponse = {
  info?: ApplicationUserInfo
}

export type GetUsersFromApplicationRequest = {
  appID?: string
}

export type GetUsersFromApplicationResponse = {
  infos?: ApplicationUserInfo[]
}

export type RemoveUsersFromApplicationRequest = {
  userIDs?: string[]
  appID?: string
}

export type RemoveUsersFromApplicationResponse = {
  info?: string
}

export type GroupInfo = {
  id?: string
  appID?: string
  groupName?: string
  groupOwner?: string
  groupLogo?: string
  annotation?: string
  createAT?: number
  updateAT?: number
}

export type CreateGroupRequest = {
  info?: GroupInfo
}

export type CreateGroupResponse = {
  info?: GroupInfo
}

export type GetGroupRequest = {
  groupID?: string
  appID?: string
}

export type GetGroupResponse = {
  info?: GroupInfo
}

export type GetGroupByOwnerRequest = {
  appID?: string
  owner?: string
}

export type OwnerGroup = {
  appID?: string
  owner?: string
  infos?: GroupInfo[]
}

export type GetGroupByOwnerResponse = {
  info?: OwnerGroup
}

export type GetAllGroupsRequest = {
  appID?: string
}

export type GetAllGroupsResponse = {
  infos?: GroupInfo[]
}

export type UpdateGroupRequest = {
  info?: GroupInfo
}

export type UpdateGroupResponse = {
  info?: GroupInfo
}

export type DeleteGroupRequest = {
  groupID?: string
  appID?: string
}

export type DeleteGroupResponse = {
  info?: string
}

export type GroupUserInfo = {
  id?: string
  groupID?: string
  appID?: string
  userID?: string
  annotation?: string
  createAT?: number
}

export type AddGroupUsersRequest = {
  userIDs?: string[]
  groupID?: string
  appID?: string
}

export type AddGroupUsersResponse = {
  infos?: GroupUserInfo[]
}

export type GetGroupUsersRequest = {
  groupID?: string
  appID?: string
}

export type GetGroupUsersResponse = {
  infos?: GroupUserInfo[]
}

export type RemoveGroupUsersRequest = {
  userIDs?: string[]
  groupID?: string
  appID?: string
}

export type RemoveGroupUsersResponse = {
  info?: string
}

export type ResourceInfo = {
  id?: string
  appID?: string
  resourceName?: string
  resourceDescription?: string
  type?: string
  creator?: string
  createAT?: number
  updateAT?: number
}

export type CreateResourceRequest = {
  info?: ResourceInfo
}

export type CreateResourceResponse = {
  info?: ResourceInfo
}

export type UpdateResourceRequest = {
  info?: ResourceInfo
}

export type UpdateResourceResponse = {
  info?: ResourceInfo
}

export type GetResourceRequest = {
  resourceID?: string
  appID?: string
}

export type GetResourceResponse = {
  info?: ResourceInfo
}

export type GetResourceByCreatorRequest = {
  appID?: string
  creator?: string
}

export type CreatorResource = {
  appID?: string
  creator?: string
  infos?: ResourceInfo[]
}

export type GetResourceByCreatorResponse = {
  info?: CreatorResource
}

export type GetResourcesRequest = {
  appID?: string
}

export type GetResourcesResponse = {
  infos?: ResourceInfo[]
}

export type DeleteResourceRequest = {
  resourceID?: string
  appID?: string
}

export type DeleteResourceResponse = {
  info?: string
}

export type SetGALoginRequest = {
  appID?: string
  userID?: string
  set?: boolean
}

export type SetGALoginResponse = {
  info?: string
}

export type AddUserLoginTimeRequest = {
  userID?: string
  appID?: string
}

export type AddUserLoginTimeResponse = {
  info?: number
}

export type UpdateUserGAStatusRequest = {
  userID?: string
  appID?: string
  status?: boolean
}

export type UpdateUserGAStatusResponse = {
  info?: string
}

export type UpdateUserKYCStatusRequest = {
  userID?: string
  appID?: string
  status?: boolean
}

export type UpdateUserKYCStatusResponse = {
  info?: string
}

export type GetUserGroupRequest = {
  appID?: string
  userID?: string
}

export type GetUserGroupResponse = {
  infos?: GroupUserInfo[]
}

export type SetSMSLoginRequest = {
  appID?: string
  userID?: string
  set?: boolean
}

export type SetSMSLoginResponse = {
  info?: string
}

export type GetResourceByNameRequest = {
  resourceName?: string
  appID?: string
}

export type GetResourceByNameResponse = {
  info?: ResourceInfo
}

export type GetUserAppIDRequest = {
  userID?: string
}

export type GetUserAppIDResponse = {
  infos?: string[]
}

export class ApplicationManagement {
  static Version(req: GoogleProtobufEmpty.Empty, initReq?: fm.InitReq): Promise<NpoolV1Npool.VersionResponse> {
    return fm.fetchReq<GoogleProtobufEmpty.Empty, NpoolV1Npool.VersionResponse>(`/version`, {...initReq, method: "POST", body: JSON.stringify(req)})
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