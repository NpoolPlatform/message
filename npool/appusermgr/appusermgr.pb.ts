/* eslint-disable */
// @ts-nocheck
/*
* This file is a generated Typescript file for GRPC Gateway, DO NOT MODIFY
*/

import * as fm from "../../fetch.pb"
import * as GoogleProtobufEmpty from "../../google/protobuf/empty.pb"
import * as NpoolV1Npool from "../npool.pb"
export type App = {
  id?: string
  createdBy?: string
  name?: string
  logo?: string
  description?: string
  createAt?: number
}

export type CreateAdminAppsRequest = {
}

export type CreateAdminAppsResponse = {
  infos?: App[]
}

export type GetAdminAppsRequest = {
}

export type GetAdminAppsResponse = {
  infos?: App[]
}

export type CreateAppRequest = {
  info?: App
}

export type CreateAppResponse = {
  info?: App
}

export type GetAppRequest = {
  id?: string
}

export type GetAppResponse = {
  info?: App
}

export type GetAppsRequest = {
  pageInfo?: NpoolV1Npool.PageInfo
}

export type GetAppsResponse = {
  infos?: App[]
  total?: number
}

export type GetAppsByCreatorRequest = {
  userID?: string
  pageInfo?: NpoolV1Npool.PageInfo
}

export type GetAppsByCreatorResponse = {
  infos?: App[]
  total?: number
}

export type UpdateAppRequest = {
  info?: App
}

export type UpdateAppResponse = {
  info?: App
}

export type AppControl = {
  id?: string
  appID?: string
  signupMethods?: string[]
  externSigninMethods?: string[]
  recaptchaMethod?: string
  kycEnable?: boolean
  signinVerifyEnable?: boolean
  invitationCodeMust?: boolean
}

export type CreateAppControlRequest = {
  info?: AppControl
}

export type CreateAppControlResponse = {
  info?: AppControl
}

export type CreateAppControlForOtherAppRequest = {
  targetAppID?: string
  info?: AppControl
}

export type CreateAppControlForOtherAppResponse = {
  info?: AppControl
}

export type GetAppControlRequest = {
  id?: string
}

export type GetAppControlResponse = {
  info?: AppControl
}

export type GetAppControlByAppRequest = {
  appID?: string
}

export type GetAppControlByAppResponse = {
  info?: AppControl
}

export type UpdateAppControlRequest = {
  info?: AppControl
}

export type UpdateAppControlResponse = {
  info?: AppControl
}

export type BanApp = {
  id?: string
  appID?: string
  message?: string
}

export type CreateBanAppRequest = {
  info?: BanApp
}

export type CreateBanAppResponse = {
  info?: BanApp
}

export type GetBanAppRequest = {
  id?: string
}

export type GetBanAppResponse = {
  info?: BanApp
}

export type GetBanAppByAppRequest = {
  appID?: string
}

export type GetBanAppByAppResponse = {
  info?: BanApp
}

export type UpdateBanAppRequest = {
  info?: BanApp
}

export type UpdateBanAppResponse = {
  info?: BanApp
}

export type DeleteBanAppRequest = {
  id?: string
}

export type DeleteBanAppResponse = {
  info?: BanApp
}

export type AppInfo = {
  app?: App
  ctrl?: AppControl
  ban?: BanApp
}

export type GetAppInfoRequest = {
  id?: string
}

export type GetAppInfoResponse = {
  info?: AppInfo
}

export type GetAppInfosRequest = {
  pageInfo?: NpoolV1Npool.PageInfo
}

export type GetAppInfosResponse = {
  infos?: AppInfo[]
  total?: number
}

export type GetAppInfosByCreatorRequest = {
  userID?: string
  pageInfo?: NpoolV1Npool.PageInfo
}

export type GetAppInfosByCreatorResponse = {
  infos?: AppInfo[]
  total?: number
}

export type AppUser = {
  id?: string
  appID?: string
  emailAddress?: string
  phoneNO?: string
  importFromApp?: string
  createAt?: number
}

export type CreateAppUserRequest = {
  info?: AppUser
}

export type CreateAppUserResponse = {
  info?: AppUser
}

export type GetAppUserRequest = {
  id?: string
}

export type GetAppUserResponse = {
  info?: AppUser
}

export type GetAppUserByAppUserRequest = {
  appID?: string
  userID?: string
}

export type GetAppUserByAppUserResponse = {
  info?: AppUser
}

export type GetAppUserByAppAccountRequest = {
  appID?: string
  account?: string
}

export type GetAppUserByAppAccountResponse = {
  info?: AppUser
}

export type GetAppUsersByAppRequest = {
  appID?: string
  pageInfo?: NpoolV1Npool.PageInfo
}

export type GetAppUsersByAppResponse = {
  infos?: AppUser[]
  total?: number
}

export type UpdateAppUserRequest = {
  info?: AppUser
}

export type UpdateAppUserResponse = {
  info?: AppUser
}

export type AppUserSecret = {
  id?: string
  appID?: string
  userID?: string
  passwordHash?: string
  salt?: string
  googleSecret?: string
}

export type AppUserSecretMap = {
  hasGoogleSecret?: boolean
}

export type CreateAppUserSecretRequest = {
  info?: AppUserSecret
}

export type CreateAppUserSecretResponse = {
  info?: AppUserSecret
}

export type GetAppUserSecretRequest = {
  id?: string
}

export type GetAppUserSecretResponse = {
  info?: AppUserSecret
}

export type GetAppUserSecretByAppUserRequest = {
  appID?: string
  userID?: string
}

export type GetAppUserSecretByAppUserResponse = {
  info?: AppUserSecret
}

export type UpdateAppUserSecretRequest = {
  info?: AppUserSecret
}

export type UpdateAppUserSecretResponse = {
  info?: AppUserSecret
}

export type AppUserExtra = {
  id?: string
  appID?: string
  userID?: string
  username?: string
  addressFields?: string[]
  gender?: string
  postalCode?: string
  age?: number
  birthday?: number
  avatar?: string
  organization?: string
  firstName?: string
  lastName?: string
  iDNumber?: string
}

export type CreateAppUserExtraRequest = {
  info?: AppUserExtra
}

export type CreateAppUserExtraResponse = {
  info?: AppUserExtra
}

export type GetAppUserExtraRequest = {
  id?: string
}

export type GetAppUserExtraResponse = {
  info?: AppUserExtra
}

export type GetAppUserExtraByAppUserRequest = {
  appID?: string
  userID?: string
}

export type GetAppUserExtraByAppUserResponse = {
  info?: AppUserExtra
}

export type UpdateAppUserExtraRequest = {
  info?: AppUserExtra
}

export type UpdateAppUserExtraResponse = {
  info?: AppUserExtra
}

export type BanAppUser = {
  id?: string
  appID?: string
  userID?: string
  message?: string
}

export type CreateBanAppUserRequest = {
  info?: BanAppUser
}

export type CreateBanAppUserResponse = {
  info?: BanAppUser
}

export type GetBanAppUserRequest = {
  id?: string
}

export type GetBanAppUserResponse = {
  info?: BanAppUser
}

export type GetBanAppUserByAppUserRequest = {
  appID?: string
  userID?: string
}

export type GetBanAppUserByAppUserResponse = {
  info?: BanAppUser
}

export type UpdateBanAppUserRequest = {
  info?: BanAppUser
}

export type UpdateBanAppUserResponse = {
  info?: BanAppUser
}

export type DeleteBanAppUserRequest = {
  id?: string
}

export type DeleteBanAppUserResponse = {
  info?: BanAppUser
}

export type AppUserControl = {
  id?: string
  appID?: string
  userID?: string
  signinVerifyByGoogleAuthentication?: boolean
  googleAuthenticationVerified?: boolean
}

export type CreateAppUserControlRequest = {
  info?: AppUserControl
}

export type CreateAppUserControlResponse = {
  info?: AppUserControl
}

export type GetAppUserControlRequest = {
  id?: string
}

export type GetAppUserControlResponse = {
  info?: AppUserControl
}

export type GetAppUserControlByAppUserRequest = {
  appID?: string
  userID?: string
}

export type GetAppUserControlByAppUserResponse = {
  info?: AppUserControl
}

export type UpdateAppUserControlRequest = {
  info?: AppUserControl
}

export type UpdateAppUserControlResponse = {
  info?: AppUserControl
}

export type AppRole = {
  id?: string
  appID?: string
  createdBy?: string
  role?: string
  description?: string
  default?: boolean
}

export type CreateGenesisRoleRequest = {
}

export type CreateGenesisRoleResponse = {
  info?: AppRole
}

export type GetGenesisRoleRequest = {
}

export type GetGenesisRoleResponse = {
  info?: AppRole
}

export type GetGenesisAppRoleUsersByOtherAppRequest = {
  appID?: string
  targetAppID?: string
}

export type GetGenesisAppRoleUsersByOtherAppResponse = {
  infos?: AppRoleUser[]
  total?: number
}

export type CreateAppRoleRequest = {
  userID?: string
  info?: AppRole
}

export type CreateAppRoleResponse = {
  info?: AppRole
}

export type CreateAppRoleForOtherAppRequest = {
  targetAppID?: string
  userID?: string
  info?: AppRole
}

export type CreateAppRoleForOtherAppResponse = {
  info?: AppRole
}

export type GetAppRoleRequest = {
  id?: string
}

export type GetAppRoleResponse = {
  info?: AppRole
}

export type GetAppRoleByAppRoleRequest = {
  appID?: string
  role?: string
}

export type GetAppRoleByAppRoleResponse = {
  info?: AppRole
}

export type GetAppRolesByAppRequest = {
  appID?: string
  pageInfo?: NpoolV1Npool.PageInfo
}

export type GetAppRolesByAppResponse = {
  infos?: AppRole[]
  total?: number
}

export type GetAppRolesByOtherAppRequest = {
  targetAppID?: string
  pageInfo?: NpoolV1Npool.PageInfo
}

export type GetAppRolesByOtherAppResponse = {
  infos?: AppRole[]
  total?: number
}

export type UpdateAppRoleRequest = {
  info?: AppRole
}

export type UpdateAppRoleResponse = {
  info?: AppRole
}

export type AppRoleUser = {
  id?: string
  appID?: string
  roleID?: string
  userID?: string
}

export type CreateGenesisRoleUserRequest = {
  user?: AppUser
  secret?: AppUserSecret
}

export type CreateGenesisRoleUserResponse = {
  user?: AppUser
  roleUser?: AppRoleUser
}

export type CreateAppRoleUserRequest = {
  info?: AppRoleUser
}

export type CreateAppRoleUserResponse = {
  info?: AppRoleUser
}

export type CreateAppRoleUserForOtherAppUserRequest = {
  targetAppID?: string
  targetUserID?: string
  info?: AppRoleUser
}

export type CreateAppRoleUserForOtherAppUserResponse = {
  info?: AppRoleUser
}

export type CreateAppRoleUserForAppOtherUserRequest = {
  targetUserID?: string
  info?: AppRoleUser
}

export type CreateAppRoleUserForAppOtherUserResponse = {
  info?: AppRoleUser
}

export type GetAppRoleUserRequest = {
  id?: string
}

export type GetAppRoleUserResponse = {
  info?: AppRoleUser
}

export type GetAppRoleUserByAppUserRequest = {
  appID?: string
  userID?: string
}

export type GetAppRoleUserByAppUserResponse = {
  infos?: AppRoleUser[]
}

export type GetAppRoleUsersByAppRoleRequest = {
  appID?: string
  roleID?: string
  pageInfo?: NpoolV1Npool.PageInfo
}

export type GetAppRoleUsersByAppRoleResponse = {
  infos?: AppRoleUser[]
  total?: number
}

export type GetAppRoleUsersByOtherAppRoleRequest = {
  appID?: string
  userID?: string
  roleID?: string
  targetAppID?: string
  pageInfo?: NpoolV1Npool.PageInfo
}

export type GetAppRoleUsersByOtherAppRoleResponse = {
  infos?: AppRoleUser[]
  total?: number
}

export type GetAppRoleUsersByAppRequest = {
  appID?: string
  pageInfo?: NpoolV1Npool.PageInfo
}

export type GetAppRoleUsersByAppResponse = {
  infos?: AppRoleUser[]
  total?: number
}

export type GetAppRoleUsersByOtherAppRequest = {
  targetAppID?: string
  pageInfo?: NpoolV1Npool.PageInfo
}

export type GetAppRoleUsersByOtherAppResponse = {
  infos?: AppRoleUser[]
  total?: number
}

export type GetUserRolesByAppUserRequest = {
  appID?: string
  userID?: string
  pageInfo?: NpoolV1Npool.PageInfo
}

export type GetUserRolesByAppUserResponse = {
  infos?: AppRole[]
  total?: number
}

export type DeleteAppRoleUserRequest = {
  id?: string
}

export type DeleteAppRoleUserResponse = {
  info?: AppRoleUser
}

export type AppUserInfo = {
  user?: AppUser
  extra?: AppUserExtra
  ctrl?: AppUserControl
  ban?: BanAppUser
  roles?: AppRole[]
  secretMap?: AppUserSecretMap
}

export type GetAppUserInfoRequest = {
  id?: string
}

export type GetAppUserInfoResponse = {
  info?: AppUserInfo
}

export type GetAppUserInfoByAppUserRequest = {
  appID?: string
  userID?: string
}

export type GetAppUserInfoByAppUserResponse = {
  info?: AppUserInfo
}

export type GetAppUserInfosByAppRequest = {
  appID?: string
  pageInfo?: NpoolV1Npool.PageInfo
}

export type GetAppUserInfosByAppResponse = {
  infos?: AppUserInfo[]
  total?: number
}

export type GetAppUserInfosByOtherAppRequest = {
  targetAppID?: string
  pageInfo?: NpoolV1Npool.PageInfo
}

export type GetAppUserInfosByOtherAppResponse = {
  infos?: AppUserInfo[]
  total?: number
}

export type CreateAppUserWithSecretRequest = {
  user?: AppUser
  secret?: AppUserSecret
}

export type CreateAppUserWithSecretResponse = {
  info?: AppUser
}

export type VerifyAppUserByAppAccountPasswordRequest = {
  appID?: string
  account?: string
  passwordHash?: string
}

export type VerifyAppUserByAppAccountPasswordResponse = {
  info?: AppUserInfo
}

export type AppUserThird = {
  id?: string
  appID?: string
  userID?: string
  thirdUserId?: string
  third?: string
  thirdUserName?: string
  thirdUserPicture?: string
  thirdExtra?: string
  thirdId?: string
}

export type CreateAppUserThirdRequest = {
  info?: AppUserThird
}

export type CreateAppUserThirdResponse = {
  info?: AppUserThird
}

export type GetAppUserThirdByAppThirdRequest = {
  appID?: string
  thirdId?: string
  thirdUserId?: string
}

export type GetAppUserThirdByAppThirdResponse = {
  info?: AppUserThird
}

export type CreateAppUserWithThirdRequest = {
  user?: AppUser
  third?: AppUserThird
}

export type CreateAppUserWithThirdResponse = {
  info?: AppUser
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
  static CreateAppControlForOtherApp(req: CreateAppControlForOtherAppRequest, initReq?: fm.InitReq): Promise<CreateAppControlForOtherAppResponse> {
    return fm.fetchReq<CreateAppControlForOtherAppRequest, CreateAppControlForOtherAppResponse>(`/v1/create/app/control/for/other/app`, {...initReq, method: "POST", body: JSON.stringify(req)})
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
  static GetAppUserByAppUser(req: GetAppUserByAppUserRequest, initReq?: fm.InitReq): Promise<GetAppUserByAppUserResponse> {
    return fm.fetchReq<GetAppUserByAppUserRequest, GetAppUserByAppUserResponse>(`/v1/get/app/user/by/app/user`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetAppUserByAppAccount(req: GetAppUserByAppAccountRequest, initReq?: fm.InitReq): Promise<GetAppUserByAppAccountResponse> {
    return fm.fetchReq<GetAppUserByAppAccountRequest, GetAppUserByAppAccountResponse>(`/v1/get/app/user/by/account`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static VerifyAppUserByAppAccountPassword(req: VerifyAppUserByAppAccountPasswordRequest, initReq?: fm.InitReq): Promise<VerifyAppUserByAppAccountPasswordResponse> {
    return fm.fetchReq<VerifyAppUserByAppAccountPasswordRequest, VerifyAppUserByAppAccountPasswordResponse>(`/v1/verify/app/user/by/account/password`, {...initReq, method: "POST", body: JSON.stringify(req)})
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
  static GetGenesisAppRoleUsersByOtherApp(req: GetGenesisAppRoleUsersByOtherAppRequest, initReq?: fm.InitReq): Promise<GetGenesisAppRoleUsersByOtherAppResponse> {
    return fm.fetchReq<GetGenesisAppRoleUsersByOtherAppRequest, GetGenesisAppRoleUsersByOtherAppResponse>(`/v1/get/genesis/app/role/users/by/other/app`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static CreateGenesisRoleUser(req: CreateGenesisRoleUserRequest, initReq?: fm.InitReq): Promise<CreateGenesisRoleUserResponse> {
    return fm.fetchReq<CreateGenesisRoleUserRequest, CreateGenesisRoleUserResponse>(`/v1/create/genesis/role/user`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static CreateAppRole(req: CreateAppRoleRequest, initReq?: fm.InitReq): Promise<CreateAppRoleResponse> {
    return fm.fetchReq<CreateAppRoleRequest, CreateAppRoleResponse>(`/v1/create/app/role`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static CreateAppRoleForOtherApp(req: CreateAppRoleForOtherAppRequest, initReq?: fm.InitReq): Promise<CreateAppRoleForOtherAppResponse> {
    return fm.fetchReq<CreateAppRoleForOtherAppRequest, CreateAppRoleForOtherAppResponse>(`/v1/create/app/role/for/other/app`, {...initReq, method: "POST", body: JSON.stringify(req)})
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
  static GetAppRolesByOtherApp(req: GetAppRolesByOtherAppRequest, initReq?: fm.InitReq): Promise<GetAppRolesByOtherAppResponse> {
    return fm.fetchReq<GetAppRolesByOtherAppRequest, GetAppRolesByOtherAppResponse>(`/v1/get/app/roles/by/other/app`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static UpdateAppRole(req: UpdateAppRoleRequest, initReq?: fm.InitReq): Promise<UpdateAppRoleResponse> {
    return fm.fetchReq<UpdateAppRoleRequest, UpdateAppRoleResponse>(`/v1/update/app/role`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static CreateAppRoleUser(req: CreateAppRoleUserRequest, initReq?: fm.InitReq): Promise<CreateAppRoleUserResponse> {
    return fm.fetchReq<CreateAppRoleUserRequest, CreateAppRoleUserResponse>(`/v1/create/app/role/user`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static CreateAppRoleUserForOtherAppUser(req: CreateAppRoleUserForOtherAppUserRequest, initReq?: fm.InitReq): Promise<CreateAppRoleUserForOtherAppUserResponse> {
    return fm.fetchReq<CreateAppRoleUserForOtherAppUserRequest, CreateAppRoleUserForOtherAppUserResponse>(`/v1/create/app/role/user/for/other/app/user`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static CreateAppRoleUserForAppOtherUser(req: CreateAppRoleUserForAppOtherUserRequest, initReq?: fm.InitReq): Promise<CreateAppRoleUserForAppOtherUserResponse> {
    return fm.fetchReq<CreateAppRoleUserForAppOtherUserRequest, CreateAppRoleUserForAppOtherUserResponse>(`/v1/create/app/role/user/for/app/other/user`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetAppRoleUser(req: GetAppRoleUserRequest, initReq?: fm.InitReq): Promise<GetAppRoleUserResponse> {
    return fm.fetchReq<GetAppRoleUserRequest, GetAppRoleUserResponse>(`/v1/get/app/role/user`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetAppRoleUserByAppUser(req: GetAppRoleUserByAppUserRequest, initReq?: fm.InitReq): Promise<GetAppRoleUserByAppUserResponse> {
    return fm.fetchReq<GetAppRoleUserByAppUserRequest, GetAppRoleUserByAppUserResponse>(`/v1/get/app/role/user/by/app/user`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetAppRoleUsersByAppRole(req: GetAppRoleUsersByAppRoleRequest, initReq?: fm.InitReq): Promise<GetAppRoleUsersByAppRoleResponse> {
    return fm.fetchReq<GetAppRoleUsersByAppRoleRequest, GetAppRoleUsersByAppRoleResponse>(`/v1/get/app/role/users/by/app/role`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetAppRoleUsersByOtherAppRole(req: GetAppRoleUsersByOtherAppRoleRequest, initReq?: fm.InitReq): Promise<GetAppRoleUsersByOtherAppRoleResponse> {
    return fm.fetchReq<GetAppRoleUsersByOtherAppRoleRequest, GetAppRoleUsersByOtherAppRoleResponse>(`/v1/get/app/role/users/by/other/app/role`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetAppRoleUsersByApp(req: GetAppRoleUsersByAppRequest, initReq?: fm.InitReq): Promise<GetAppRoleUsersByAppResponse> {
    return fm.fetchReq<GetAppRoleUsersByAppRequest, GetAppRoleUsersByAppResponse>(`/v1/get/app/role/users/by/app`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetAppRoleUsersByOtherApp(req: GetAppRoleUsersByOtherAppRequest, initReq?: fm.InitReq): Promise<GetAppRoleUsersByOtherAppResponse> {
    return fm.fetchReq<GetAppRoleUsersByOtherAppRequest, GetAppRoleUsersByOtherAppResponse>(`/v1/get/app/role/users/by/other/app`, {...initReq, method: "POST", body: JSON.stringify(req)})
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
  static GetAppUserInfoByAppUser(req: GetAppUserInfoByAppUserRequest, initReq?: fm.InitReq): Promise<GetAppUserInfoByAppUserResponse> {
    return fm.fetchReq<GetAppUserInfoByAppUserRequest, GetAppUserInfoByAppUserResponse>(`/v1/get/app/userinfo/by/app/user`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetAppUserInfosByApp(req: GetAppUserInfosByAppRequest, initReq?: fm.InitReq): Promise<GetAppUserInfosByAppResponse> {
    return fm.fetchReq<GetAppUserInfosByAppRequest, GetAppUserInfosByAppResponse>(`/v1/get/app/userinfos/by/app`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetAppUserInfosByOtherApp(req: GetAppUserInfosByOtherAppRequest, initReq?: fm.InitReq): Promise<GetAppUserInfosByOtherAppResponse> {
    return fm.fetchReq<GetAppUserInfosByOtherAppRequest, GetAppUserInfosByOtherAppResponse>(`/v1/get/app/userinfos/by/other/app`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static CreateAppUserWithSecret(req: CreateAppUserWithSecretRequest, initReq?: fm.InitReq): Promise<CreateAppUserWithSecretResponse> {
    return fm.fetchReq<CreateAppUserWithSecretRequest, CreateAppUserWithSecretResponse>(`/v1/create/app/user/with/secret`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static CreateAppUserWithThird(req: CreateAppUserWithThirdRequest, initReq?: fm.InitReq): Promise<CreateAppUserWithThirdResponse> {
    return fm.fetchReq<CreateAppUserWithThirdRequest, CreateAppUserWithThirdResponse>(`/v1/create/app/user/with/third`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static CreateAppUserThird(req: CreateAppUserThirdRequest, initReq?: fm.InitReq): Promise<CreateAppUserThirdResponse> {
    return fm.fetchReq<CreateAppUserThirdRequest, CreateAppUserThirdResponse>(`/v1/create/app/user/third`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetAppUserThirdByAppThird(req: GetAppUserThirdByAppThirdRequest, initReq?: fm.InitReq): Promise<GetAppUserThirdByAppThirdResponse> {
    return fm.fetchReq<GetAppUserThirdByAppThirdRequest, GetAppUserThirdByAppThirdResponse>(`/v1/get/app/user/third/by/app/user`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
}