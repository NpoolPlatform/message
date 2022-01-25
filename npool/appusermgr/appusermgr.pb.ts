/* eslint-disable */
// @ts-nocheck
/*
* This file is a generated Typescript file for GRPC Gateway, DO NOT MODIFY
*/

import * as fm from "../../fetch.pb"
import * as GoogleProtobufEmpty from "../../google/protobuf/empty.pb"
import * as NpoolV1Npool from "../npool.pb"
export type App = {
  ID?: string
  CreatedBy?: string
  Name?: string
  Logo?: string
  Description?: string
  CreateAt?: number
}

export type CreateAdminAppsRequest = {
}

export type CreateAdminAppsResponse = {
  Infos?: App[]
}

export type GetAdminAppsRequest = {
}

export type GetAdminAppsResponse = {
  Infos?: App[]
}

export type CreateAppRequest = {
  Info?: App
}

export type CreateAppResponse = {
  Info?: App
}

export type GetAppRequest = {
  ID?: string
}

export type GetAppResponse = {
  Info?: App
}

export type GetAppsRequest = {
  PageInfo?: NpoolV1Npool.PageInfo
}

export type GetAppsResponse = {
  Infos?: App[]
  Total?: number
}

export type GetAppsByCreatorRequest = {
  UserID?: string
  PageInfo?: NpoolV1Npool.PageInfo
}

export type GetAppsByCreatorResponse = {
  Infos?: App[]
  Total?: number
}

export type UpdateAppRequest = {
  Info?: App
}

export type UpdateAppResponse = {
  Info?: App
}

export type AppControl = {
  ID?: string
  AppID?: string
  SignupMethods?: string[]
  ExternSigninMethods?: string[]
  RecaptchaMethod?: string
  KycEnable?: boolean
  SigninVerifyEnable?: boolean
}

export type CreateAppControlRequest = {
  Info?: AppControl
}

export type CreateAppControlResponse = {
  Info?: AppControl
}

export type GetAppControlRequest = {
  ID?: string
}

export type GetAppControlResponse = {
  Info?: AppControl
}

export type GetAppControlByAppRequest = {
  AppID?: string
}

export type GetAppControlByAppResponse = {
  Info?: AppControl
}

export type UpdateAppControlRequest = {
  Info?: AppControl
}

export type UpdateAppControlResponse = {
  Info?: AppControl
}

export type BanApp = {
  ID?: string
  AppID?: string
  Message?: string
}

export type CreateBanAppRequest = {
  Info?: BanApp
}

export type CreateBanAppResponse = {
  Info?: BanApp
}

export type GetBanAppRequest = {
  ID?: string
}

export type GetBanAppResponse = {
  Info?: BanApp
}

export type GetBanAppByAppRequest = {
  AppID?: string
}

export type GetBanAppByAppResponse = {
  Info?: BanApp
}

export type UpdateBanAppRequest = {
  Info?: BanApp
}

export type UpdateBanAppResponse = {
  Info?: BanApp
}

export type DeleteBanAppRequest = {
  ID?: string
}

export type DeleteBanAppResponse = {
  Info?: BanApp
}

export type AppInfo = {
  App?: App
  Ctrl?: AppControl
  Ban?: BanApp
}

export type GetAppInfoRequest = {
  ID?: string
}

export type GetAppInfoResponse = {
  Info?: AppInfo
}

export type GetAppInfosRequest = {
  PageInfo?: NpoolV1Npool.PageInfo
}

export type GetAppInfosResponse = {
  Infos?: AppInfo[]
  Total?: number
}

export type GetAppInfosByCreatorRequest = {
  UserID?: string
  PageInfo?: NpoolV1Npool.PageInfo
}

export type GetAppInfosByCreatorResponse = {
  Infos?: AppInfo[]
  Total?: number
}

export type AppUser = {
  ID?: string
  AppID?: string
  EmailAddress?: string
  PhoneNO?: string
  ImportFromApp?: string
  CreateAt?: number
}

export type CreateAppUserRequest = {
  Info?: AppUser
}

export type CreateAppUserResponse = {
  Info?: AppUser
}

export type GetAppUserRequest = {
  ID?: string
}

export type GetAppUserResponse = {
  Info?: AppUser
}

export type GetAppUsersByAppRequest = {
  AppID?: string
  PageInfo?: NpoolV1Npool.PageInfo
}

export type GetAppUsersByAppResponse = {
  Infos?: AppUser[]
  Total?: number
}

export type UpdateAppUserRequest = {
  Info?: AppUser
}

export type UpdateAppUserResponse = {
  Info?: AppUser
}

export type AppUserSecret = {
  ID?: string
  AppID?: string
  UserID?: string
  PasswordHash?: string
  Salt?: string
  GoogleSecret?: string
}

export type CreateAppUserSecretRequest = {
  Info?: AppUserSecret
}

export type CreateAppUserSecretResponse = {
  Info?: AppUserSecret
}

export type GetAppUserSecretRequest = {
  ID?: string
}

export type GetAppUserSecretResponse = {
  Info?: AppUserSecret
}

export type GetAppUserSecretByAppUserRequest = {
  AppID?: string
  UserID?: string
}

export type GetAppUserSecretByAppUserResponse = {
  Info?: AppUserSecret
}

export type UpdateAppUserSecretRequest = {
  Info?: AppUserSecret
}

export type UpdateAppUserSecretResponse = {
  Info?: AppUserSecret
}

export type AppUserExtra = {
  ID?: string
  AppID?: string
  UserID?: string
  Username?: string
  AddressFields?: string[]
  Gender?: string
  PostalCode?: string
  Age?: number
  Birthday?: number
  Avatar?: string
  Organization?: string
}

export type CreateAppUserExtraRequest = {
  Info?: AppUserExtra
}

export type CreateAppUserExtraResponse = {
  Info?: AppUserExtra
}

export type GetAppUserExtraRequest = {
  ID?: string
}

export type GetAppUserExtraResponse = {
  Info?: AppUserExtra
}

export type GetAppUserExtraByAppUserRequest = {
  AppID?: string
  UserID?: string
}

export type GetAppUserExtraByAppUserResponse = {
  Info?: AppUserExtra
}

export type UpdateAppUserExtraRequest = {
  Info?: AppUserExtra
}

export type UpdateAppUserExtraResponse = {
  Info?: AppUserExtra
}

export type BanAppUser = {
  ID?: string
  AppID?: string
  UserID?: string
  Message?: string
}

export type CreateBanAppUserRequest = {
  Info?: BanAppUser
}

export type CreateBanAppUserResponse = {
  Info?: BanAppUser
}

export type GetBanAppUserRequest = {
  ID?: string
}

export type GetBanAppUserResponse = {
  Info?: BanAppUser
}

export type GetBanAppUserByAppUserRequest = {
  AppID?: string
  UserID?: string
}

export type GetBanAppUserByAppUserResponse = {
  Info?: BanAppUser
}

export type UpdateBanAppUserRequest = {
  Info?: BanAppUser
}

export type UpdateBanAppUserResponse = {
  Info?: BanAppUser
}

export type DeleteBanAppUserRequest = {
  ID?: string
}

export type DeleteBanAppUserResponse = {
  Info?: BanAppUser
}

export type AppUserControl = {
  ID?: string
  AppID?: string
  UserID?: string
}

export type CreateAppUserControlRequest = {
  Info?: AppUserControl
}

export type CreateAppUserControlResponse = {
  Info?: AppUserControl
}

export type GetAppUserControlRequest = {
  ID?: string
}

export type GetAppUserControlResponse = {
  Info?: AppUserControl
}

export type GetAppUserControlByAppUserRequest = {
  AppID?: string
  UserID?: string
}

export type GetAppUserControlByAppUserResponse = {
  Info?: AppUserControl
}

export type UpdateAppUserControlRequest = {
  Info?: AppUserControl
}

export type UpdateAppUserControlResponse = {
  Info?: AppUserControl
}

export type AppRole = {
  ID?: string
  AppID?: string
  CreatedBy?: string
  Role?: string
  Description?: string
  Default?: boolean
}

export type CreateGenesisRoleRequest = {
}

export type CreateGenesisRoleResponse = {
  Info?: AppRole
}

export type GetGenesisRoleRequest = {
}

export type GetGenesisRoleResponse = {
  Info?: AppRole
}

export type CreateAppRoleRequest = {
  Info?: AppRole
}

export type CreateAppRoleResponse = {
  Info?: AppRole
}

export type GetAppRoleRequest = {
  ID?: string
}

export type GetAppRoleResponse = {
  Info?: AppRole
}

export type GetAppRoleByAppRoleRequest = {
  AppID?: string
  Role?: string
}

export type GetAppRoleByAppRoleResponse = {
  Info?: AppRole
}

export type GetAppRolesByAppRequest = {
  AppID?: string
  PageInfo?: NpoolV1Npool.PageInfo
}

export type GetAppRolesByAppResponse = {
  Infos?: AppRole[]
  Total?: number
}

export type UpdateAppRoleRequest = {
  Info?: AppRole
}

export type UpdateAppRoleResponse = {
  Info?: AppRole
}

export type AppRoleUser = {
  ID?: string
  AppID?: string
  RoleID?: string
  UserID?: string
}

export type CreateGenesisRoleUserRequest = {
  AppID?: string
  UserID?: string
}

export type CreateGenesisRoleUserResponse = {
  Info?: AppRoleUser
}

export type CreateAppRoleUserRequest = {
  Info?: AppRoleUser
}

export type CreateAppRoleUserResponse = {
  Info?: AppRoleUser
}

export type GetAppRoleUserRequest = {
  ID?: string
}

export type GetAppRoleUserResponse = {
  Info?: AppRoleUser
}

export type GetAppRoleUsersByAppRoleRequest = {
  AppID?: string
  RoleID?: string
  PageInfo?: NpoolV1Npool.PageInfo
}

export type GetAppRoleUsersByAppRoleResponse = {
  Infos?: AppRoleUser[]
  Total?: number
}

export type GetUserRolesByAppUserRequest = {
  AppID?: string
  UserID?: string
  PageInfo?: NpoolV1Npool.PageInfo
}

export type GetUserRolesByAppUserResponse = {
  Infos?: AppRole[]
  Total?: number
}

export type DeleteAppRoleUserRequest = {
  ID?: string
}

export type DeleteAppRoleUserResponse = {
  Info?: AppRoleUser
}

export type AppUserInfo = {
  User?: AppUser
  Extra?: AppUserExtra
  Ctrl?: AppUserControl
  Ban?: BanAppUser
  Roles?: AppRole[]
}

export type GetAppUserInfoRequest = {
  ID?: string
}

export type GetAppUserInfoResponse = {
  Info?: AppUserInfo
}

export type GetAppUserInfosRequest = {
  PageInfo?: NpoolV1Npool.PageInfo
}

export type GetAppUserInfosResponse = {
  Infos?: AppUserInfo[]
  Total?: number
}

export type GetAppUserInfosByAppRequest = {
  AppID?: string
  PageInfo?: NpoolV1Npool.PageInfo
}

export type GetAppUserInfosByAppResponse = {
  Infos?: AppUserInfo[]
  Total?: number
}

export class AppUserManager {
  static Version(req: GoogleProtobufEmpty.Empty, initReq?: fm.InitReq): Promise<NpoolV1Npool.VersionResponse> {
    return fm.fetchReq<GoogleProtobufEmpty.Empty, NpoolV1Npool.VersionResponse>(`/version`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static CreateAdminApps(req: CreateAdminAppsRequest, initReq?: fm.InitReq): Promise<CreateAdminAppsResponse> {
    return fm.fetchReq<CreateAdminAppsRequest, CreateAdminAppsResponse>(`/v1/create/admin/apps`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetAdminApps(req: GetAdminAppsRequest, initReq?: fm.InitReq): Promise<GetAdminAppsResponse> {
    return fm.fetchReq<GetAdminAppsRequest, GetAdminAppsResponse>(`/v1/get/admin/apps`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static CreateApp(req: CreateAppRequest, initReq?: fm.InitReq): Promise<CreateAppResponse> {
    return fm.fetchReq<CreateAppRequest, CreateAppResponse>(`/v1/create/app`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetApp(req: GetAppRequest, initReq?: fm.InitReq): Promise<GetAppResponse> {
    return fm.fetchReq<GetAppRequest, GetAppResponse>(`/v1/get/app`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetApps(req: GetAppsRequest, initReq?: fm.InitReq): Promise<GetAppsResponse> {
    return fm.fetchReq<GetAppsRequest, GetAppsResponse>(`/v1/get/apps`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetAppsByCreator(req: GetAppsByCreatorRequest, initReq?: fm.InitReq): Promise<GetAppsByCreatorResponse> {
    return fm.fetchReq<GetAppsByCreatorRequest, GetAppsByCreatorResponse>(`/v1/get/apps/by/creator`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static UpdateApp(req: UpdateAppRequest, initReq?: fm.InitReq): Promise<UpdateAppResponse> {
    return fm.fetchReq<UpdateAppRequest, UpdateAppResponse>(`/v1/update/app`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static CreateAppControl(req: CreateAppControlRequest, initReq?: fm.InitReq): Promise<CreateAppControlResponse> {
    return fm.fetchReq<CreateAppControlRequest, CreateAppControlResponse>(`/v1/create/app/control`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetAppControl(req: GetAppControlRequest, initReq?: fm.InitReq): Promise<GetAppControlResponse> {
    return fm.fetchReq<GetAppControlRequest, GetAppControlResponse>(`/v1/get/app/control`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetAppControlByApp(req: GetAppControlByAppRequest, initReq?: fm.InitReq): Promise<GetAppControlByAppResponse> {
    return fm.fetchReq<GetAppControlByAppRequest, GetAppControlByAppResponse>(`/v1/get/app/control/by/app`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static UpdateAppControl(req: UpdateAppControlRequest, initReq?: fm.InitReq): Promise<UpdateAppControlResponse> {
    return fm.fetchReq<UpdateAppControlRequest, UpdateAppControlResponse>(`/v1/update/app/control`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static CreateBanApp(req: CreateBanAppRequest, initReq?: fm.InitReq): Promise<CreateBanAppResponse> {
    return fm.fetchReq<CreateBanAppRequest, CreateBanAppResponse>(`/v1/create/ban/app`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetBanApp(req: GetBanAppRequest, initReq?: fm.InitReq): Promise<GetBanAppResponse> {
    return fm.fetchReq<GetBanAppRequest, GetBanAppResponse>(`/v1/get/ban/app`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetBanAppByApp(req: GetBanAppByAppRequest, initReq?: fm.InitReq): Promise<GetBanAppByAppResponse> {
    return fm.fetchReq<GetBanAppByAppRequest, GetBanAppByAppResponse>(`/v1/get/ban/app/by/app`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static UpdateBanApp(req: UpdateBanAppRequest, initReq?: fm.InitReq): Promise<UpdateBanAppResponse> {
    return fm.fetchReq<UpdateBanAppRequest, UpdateBanAppResponse>(`/v1/update/ban/app`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static DeleteBanApp(req: DeleteBanAppRequest, initReq?: fm.InitReq): Promise<DeleteBanAppResponse> {
    return fm.fetchReq<DeleteBanAppRequest, DeleteBanAppResponse>(`/v1/delete/ban/app`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetAppInfo(req: GetAppInfoRequest, initReq?: fm.InitReq): Promise<GetAppInfoResponse> {
    return fm.fetchReq<GetAppInfoRequest, GetAppInfoResponse>(`/v1/get/appinfo`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetAppInfos(req: GetAppInfosRequest, initReq?: fm.InitReq): Promise<GetAppInfosResponse> {
    return fm.fetchReq<GetAppInfosRequest, GetAppInfosResponse>(`/v1/get/appinfos`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetAppInfosByCreator(req: GetAppInfosByCreatorRequest, initReq?: fm.InitReq): Promise<GetAppInfosByCreatorResponse> {
    return fm.fetchReq<GetAppInfosByCreatorRequest, GetAppInfosByCreatorResponse>(`/v1/get/appinfos/by/creator`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static CreateAppUser(req: CreateAppUserRequest, initReq?: fm.InitReq): Promise<CreateAppUserResponse> {
    return fm.fetchReq<CreateAppUserRequest, CreateAppUserResponse>(`/v1/create/app/user`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetAppUser(req: GetAppUserRequest, initReq?: fm.InitReq): Promise<GetAppUserResponse> {
    return fm.fetchReq<GetAppUserRequest, GetAppUserResponse>(`/v1/get/app/user`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetAppUsersByApp(req: GetAppUsersByAppRequest, initReq?: fm.InitReq): Promise<GetAppUsersByAppResponse> {
    return fm.fetchReq<GetAppUsersByAppRequest, GetAppUsersByAppResponse>(`/v1/get/app/users/by/app`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static UpdateAppUser(req: UpdateAppUserRequest, initReq?: fm.InitReq): Promise<UpdateAppUserResponse> {
    return fm.fetchReq<UpdateAppUserRequest, UpdateAppUserResponse>(`/v1/update/app/user`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static CreateAppUserSecret(req: CreateAppUserSecretRequest, initReq?: fm.InitReq): Promise<CreateAppUserSecretResponse> {
    return fm.fetchReq<CreateAppUserSecretRequest, CreateAppUserSecretResponse>(`/v1/create/app/user/secret`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetAppUserSecret(req: GetAppUserSecretRequest, initReq?: fm.InitReq): Promise<GetAppUserSecretResponse> {
    return fm.fetchReq<GetAppUserSecretRequest, GetAppUserSecretResponse>(`/v1/get/app/user/secret`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetAppUserSecretByAppUser(req: GetAppUserSecretByAppUserRequest, initReq?: fm.InitReq): Promise<GetAppUserSecretByAppUserResponse> {
    return fm.fetchReq<GetAppUserSecretByAppUserRequest, GetAppUserSecretByAppUserResponse>(`/v1/get/app/user/secret/by/app/user`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static UpdateAppUserSecret(req: UpdateAppUserSecretRequest, initReq?: fm.InitReq): Promise<UpdateAppUserSecretResponse> {
    return fm.fetchReq<UpdateAppUserSecretRequest, UpdateAppUserSecretResponse>(`/v1/update/app/user/secret`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static CreateAppUserExtra(req: CreateAppUserExtraRequest, initReq?: fm.InitReq): Promise<CreateAppUserExtraResponse> {
    return fm.fetchReq<CreateAppUserExtraRequest, CreateAppUserExtraResponse>(`/v1/create/app/user/extra`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetAppUserExtra(req: GetAppUserExtraRequest, initReq?: fm.InitReq): Promise<GetAppUserExtraResponse> {
    return fm.fetchReq<GetAppUserExtraRequest, GetAppUserExtraResponse>(`/v1/get/app/user/extra`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetAppUserExtraByAppUser(req: GetAppUserExtraByAppUserRequest, initReq?: fm.InitReq): Promise<GetAppUserExtraByAppUserResponse> {
    return fm.fetchReq<GetAppUserExtraByAppUserRequest, GetAppUserExtraByAppUserResponse>(`/v1/get/app/user/extra/by/app/user`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static UpdateAppUserExtra(req: UpdateAppUserExtraRequest, initReq?: fm.InitReq): Promise<UpdateAppUserExtraResponse> {
    return fm.fetchReq<UpdateAppUserExtraRequest, UpdateAppUserExtraResponse>(`/v1/update/app/user/extra`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static CreateBanAppUser(req: CreateBanAppUserRequest, initReq?: fm.InitReq): Promise<CreateBanAppUserResponse> {
    return fm.fetchReq<CreateBanAppUserRequest, CreateBanAppUserResponse>(`/v1/create/ban/app/user`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetBanAppUser(req: GetBanAppUserRequest, initReq?: fm.InitReq): Promise<GetBanAppUserResponse> {
    return fm.fetchReq<GetBanAppUserRequest, GetBanAppUserResponse>(`/v1/get/ban/app/user`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetBanAppUserByAppUser(req: GetBanAppUserByAppUserRequest, initReq?: fm.InitReq): Promise<GetBanAppUserByAppUserResponse> {
    return fm.fetchReq<GetBanAppUserByAppUserRequest, GetBanAppUserByAppUserResponse>(`/v1/get/ban/app/user/by/app/user`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static UpdateBanAppUser(req: UpdateBanAppUserRequest, initReq?: fm.InitReq): Promise<UpdateBanAppUserResponse> {
    return fm.fetchReq<UpdateBanAppUserRequest, UpdateBanAppUserResponse>(`/v1/update/ban/app/user`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static DeleteBanAppUser(req: DeleteBanAppUserRequest, initReq?: fm.InitReq): Promise<DeleteBanAppUserResponse> {
    return fm.fetchReq<DeleteBanAppUserRequest, DeleteBanAppUserResponse>(`/v1/delete/ban/app/user`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static CreateAppUserControl(req: CreateAppUserControlRequest, initReq?: fm.InitReq): Promise<CreateAppUserControlResponse> {
    return fm.fetchReq<CreateAppUserControlRequest, CreateAppUserControlResponse>(`/v1/create/app/user/control`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetAppUserControl(req: GetAppUserControlRequest, initReq?: fm.InitReq): Promise<GetAppUserControlResponse> {
    return fm.fetchReq<GetAppUserControlRequest, GetAppUserControlResponse>(`/v1/get/app/user/control`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetAppUserControlByAppUser(req: GetAppUserControlByAppUserRequest, initReq?: fm.InitReq): Promise<GetAppUserControlByAppUserResponse> {
    return fm.fetchReq<GetAppUserControlByAppUserRequest, GetAppUserControlByAppUserResponse>(`/v1/get/app/user/control/by/app/user`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static UpdateAppUserControl(req: UpdateAppUserControlRequest, initReq?: fm.InitReq): Promise<UpdateAppUserControlResponse> {
    return fm.fetchReq<UpdateAppUserControlRequest, UpdateAppUserControlResponse>(`/v1/update/app/user/control`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static CreateGenesisRole(req: CreateGenesisRoleRequest, initReq?: fm.InitReq): Promise<CreateGenesisRoleResponse> {
    return fm.fetchReq<CreateGenesisRoleRequest, CreateGenesisRoleResponse>(`/v1/create/genesis/role`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetGenesisRole(req: GetGenesisRoleRequest, initReq?: fm.InitReq): Promise<GetGenesisRoleResponse> {
    return fm.fetchReq<GetGenesisRoleRequest, GetGenesisRoleResponse>(`/v1/get/genesis/role`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static CreateGenesisRoleUser(req: CreateGenesisRoleUserRequest, initReq?: fm.InitReq): Promise<CreateGenesisRoleUserResponse> {
    return fm.fetchReq<CreateGenesisRoleUserRequest, CreateGenesisRoleUserResponse>(`/v1/create/genesis/role/user`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static CreateAppRole(req: CreateAppRoleRequest, initReq?: fm.InitReq): Promise<CreateAppRoleResponse> {
    return fm.fetchReq<CreateAppRoleRequest, CreateAppRoleResponse>(`/v1/create/app/role`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetAppRole(req: GetAppRoleRequest, initReq?: fm.InitReq): Promise<GetAppRoleResponse> {
    return fm.fetchReq<GetAppRoleRequest, GetAppRoleResponse>(`/v1/get/app/role`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetAppRoleByAppRole(req: GetAppRoleByAppRoleRequest, initReq?: fm.InitReq): Promise<GetAppRoleByAppRoleResponse> {
    return fm.fetchReq<GetAppRoleByAppRoleRequest, GetAppRoleByAppRoleResponse>(`/v1/get/app/role/by/app/role`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetAppRolesByApp(req: GetAppRolesByAppRequest, initReq?: fm.InitReq): Promise<GetAppRolesByAppResponse> {
    return fm.fetchReq<GetAppRolesByAppRequest, GetAppRolesByAppResponse>(`/v1/get/app/roles/by/app`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static UpdateAppRole(req: UpdateAppRoleRequest, initReq?: fm.InitReq): Promise<UpdateAppRoleResponse> {
    return fm.fetchReq<UpdateAppRoleRequest, UpdateAppRoleResponse>(`/v1/update/app/role`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static CreateAppRoleUser(req: CreateAppRoleUserRequest, initReq?: fm.InitReq): Promise<CreateAppRoleUserResponse> {
    return fm.fetchReq<CreateAppRoleUserRequest, CreateAppRoleUserResponse>(`/v1/create/app/role/user`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetAppRoleUser(req: GetAppRoleUserRequest, initReq?: fm.InitReq): Promise<GetAppRoleUserResponse> {
    return fm.fetchReq<GetAppRoleUserRequest, GetAppRoleUserResponse>(`/v1/get/app/role/user`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetAppRoleUsersByAppRole(req: GetAppRoleUsersByAppRoleRequest, initReq?: fm.InitReq): Promise<GetAppRoleUsersByAppRoleResponse> {
    return fm.fetchReq<GetAppRoleUsersByAppRoleRequest, GetAppRoleUsersByAppRoleResponse>(`/v1/get/app/role/users/by/app/role`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetUserRolesByAppUser(req: GetUserRolesByAppUserRequest, initReq?: fm.InitReq): Promise<GetUserRolesByAppUserResponse> {
    return fm.fetchReq<GetUserRolesByAppUserRequest, GetUserRolesByAppUserResponse>(`/v1/get/user/roles/by/app/user`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static DeleteAppRoleUser(req: DeleteAppRoleUserRequest, initReq?: fm.InitReq): Promise<DeleteAppRoleUserResponse> {
    return fm.fetchReq<DeleteAppRoleUserRequest, DeleteAppRoleUserResponse>(`/v1/delete/app/role/user`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetAppUserInfo(req: GetAppUserInfoRequest, initReq?: fm.InitReq): Promise<GetAppUserInfoResponse> {
    return fm.fetchReq<GetAppUserInfoRequest, GetAppUserInfoResponse>(`/v1/get/app/userinfo`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetAppUserInfos(req: GetAppUserInfosRequest, initReq?: fm.InitReq): Promise<GetAppUserInfosResponse> {
    return fm.fetchReq<GetAppUserInfosRequest, GetAppUserInfosResponse>(`/v1/get/app/userinfos`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetAppUserInfosByApp(req: GetAppUserInfosByAppRequest, initReq?: fm.InitReq): Promise<GetAppUserInfosByAppResponse> {
    return fm.fetchReq<GetAppUserInfosByAppRequest, GetAppUserInfosByAppResponse>(`/v1/get/app/userinfos/by/app`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
}