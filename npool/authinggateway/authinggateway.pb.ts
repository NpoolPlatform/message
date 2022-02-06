/* eslint-disable */
// @ts-nocheck
/*
* This file is a generated Typescript file for GRPC Gateway, DO NOT MODIFY
*/

import * as fm from "../../fetch.pb"
import * as GoogleProtobufEmpty from "../../google/protobuf/empty.pb"
import * as NpoolV1Npool from "../npool.pb"
export type AuthByAppRequest = {
  appID?: string
  resource?: string
  method?: string
}

export type AuthByAppResponse = {
  allowed?: boolean
}

export type AuthByAppRoleUserRequest = {
  appID?: string
  userID?: string
  token?: string
  resource?: string
  method?: string
}

export type AuthByAppRoleUserResponse = {
  allowed?: boolean
}

export type AuthHistory = {
  id?: string
  appID?: string
  userID?: string
  resource?: string
  method?: string
  allowed?: boolean
  createAt?: number
}

export type GetAuthHistoriesRequest = {
  appID?: string
  userID?: string
}

export type GetAuthHistoriesResponse = {
  infos?: AuthHistory[]
}

export type GetAuthHistoriesByAppRequest = {
  appID?: string
}

export type GetAuthHistoriesByAppResponse = {
  infos?: AuthHistory[]
}

export type GetAuthHistoriesByOtherAppRequest = {
  targetAppID?: string
}

export type GetAuthHistoriesByOtherAppResponse = {
  infos?: AuthHistory[]
}

export type Resource = {
  resource?: string
  method?: string
}

export type AppAuth = {
  id?: string
  appID?: string
  resources?: Resource[]
}

export type Auth = {
  id?: string
  appID?: string
  roleID?: string
  userID?: string
  resource?: string
  method?: string
}

export type CreateAppAuthForOtherAppRequest = {
  targetAppID?: string
  info?: AppAuth
}

export type CreateAppAuthForOtherAppResponse = {
  info?: AppAuth
}

export type GetAppAuthByAppResourceMethodRequest = {
  appID?: string
  resource?: string
  method?: string
}

export type GetAppAuthByAppResourceMethodResponse = {
  infos?: Auth[]
}

export type GetAppAuthByOtherAppResourceMethodRequest = {
  targetAppID?: string
  resource?: string
  method?: string
}

export type GetAppAuthByOtherAppResourceMethodResponse = {
  infos?: Auth[]
}

export type DeleteAppAuthRequest = {
  id?: string
}

export type DeleteAppAuthResponse = {
  info?: Auth
}

export type AppRoleAuth = {
  id?: string
  appID?: string
  roleID?: string
  resource?: string
  method?: string
}

export type CreateAppRoleAuthRequest = {
  info?: AppRoleAuth
}

export type CreateAppRoleAuthResponse = {
  info?: AppRoleAuth
}

export type CreateAppRoleAuthForOtherAppRequest = {
  targetAppID?: string
  info?: AppRoleAuth
}

export type CreateAppRoleAuthForOtherAppResponse = {
  info?: Auth
}

export type GetAppAuthByAppRoleResourceMethodRequest = {
  appID?: string
  roleID?: string
  resource?: string
  method?: string
}

export type GetAppAuthByAppRoleResourceMethodResponse = {
  info?: Auth
}

export type GetAppAuthByOtherAppRoleResourceMethodRequest = {
  targetAppID?: string
  roleID?: string
  resource?: string
  method?: string
}

export type GetAppAuthByOtherAppRoleResourceMethodResponse = {
  info?: Auth
}

export type DeleteAppRoleAuthRequest = {
  id?: string
}

export type DeleteAppRoleAuthResponse = {
  info?: Auth
}

export type AppUserAuth = {
  id?: string
  appID?: string
  userID?: string
  resource?: string
  method?: string
}

export type CreateAppUserAuthRequest = {
  info?: AppUserAuth
}

export type CreateAppUserAuthResponse = {
  info?: AppUserAuth
}

export type CreateAppUserAuthForOtherAppRequest = {
  targetAppID?: string
  info?: AppUserAuth
}

export type CreateAppUserAuthForOtherAppResponse = {
  info?: AppUserAuth
}

export type GetAppAuthByAppUserResourceMethodRequest = {
  appID?: string
  resource?: string
  method?: string
}

export type GetAppAuthByAppUserResourceMethodResponse = {
  info?: Auth
}

export type GetAppAuthByOtherAppUserResourceMethodRequest = {
  targetAppID?: string
  resource?: string
  method?: string
}

export type GetAppAuthByOtherAppUserResourceMethodResponse = {
  info?: Auth
}

export type DeleteAppUserAuthRequest = {
  id?: string
}

export type DeleteAppUserAuthResponse = {
  info?: Auth
}

export type GetAuthsByAppRequest = {
  appID?: string
}

export type GetAuthsByAppResponse = {
  infos?: Auth[]
}

export type GetAuthsByOtherAppRequest = {
  targetAppID?: string
}

export type GetAuthsByOtherAppResponse = {
  infos?: Auth[]
}

export type GetAuthsByAppRoleRequest = {
  appID?: string
  roleID?: string
}

export type GetAuthsByAppRoleResponse = {
  infos?: Auth[]
}

export type GetAuthsByOtherAppRoleRequest = {
  targetAppID?: string
  roleID?: string
}

export type GetAuthsByOtherAppRoleResponse = {
  infos?: Auth[]
}

export type GetAuthsByAppUserRequest = {
  appID?: string
  userID?: string
}

export type GetAuthsByAppUserResponse = {
  infos?: Auth[]
}

export type GetAuthsByOtherAppUserRequest = {
  targetAppID?: string
  targetUserID?: string
}

export type GetAuthsByOtherAppUserResponse = {
  infos?: Auth[]
}

export class AuthingGateway {
  static Version(req: GoogleProtobufEmpty.Empty, initReq?: fm.InitReq): Promise<NpoolV1Npool.VersionResponse> {
    return fm.fetchReq<GoogleProtobufEmpty.Empty, NpoolV1Npool.VersionResponse>(`/version`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static AuthByApp(req: AuthByAppRequest, initReq?: fm.InitReq): Promise<AuthByAppResponse> {
    return fm.fetchReq<AuthByAppRequest, AuthByAppResponse>(`/v1/auth/by/app`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static AuthByAppRoleUser(req: AuthByAppRoleUserRequest, initReq?: fm.InitReq): Promise<AuthByAppRoleUserResponse> {
    return fm.fetchReq<AuthByAppRoleUserRequest, AuthByAppRoleUserResponse>(`/v1/auth/by/app/role/user`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetAuthHistories(req: GetAuthHistoriesRequest, initReq?: fm.InitReq): Promise<GetAuthHistoriesResponse> {
    return fm.fetchReq<GetAuthHistoriesRequest, GetAuthHistoriesResponse>(`/v1/get/auth/histories`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetAuthHistoriesByApp(req: GetAuthHistoriesByAppRequest, initReq?: fm.InitReq): Promise<GetAuthHistoriesByAppResponse> {
    return fm.fetchReq<GetAuthHistoriesByAppRequest, GetAuthHistoriesByAppResponse>(`/v1/get/auth/histories/by/app`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetAuthHistoriesByOtherApp(req: GetAuthHistoriesByOtherAppRequest, initReq?: fm.InitReq): Promise<GetAuthHistoriesByOtherAppResponse> {
    return fm.fetchReq<GetAuthHistoriesByOtherAppRequest, GetAuthHistoriesByOtherAppResponse>(`/v1/get/auth/histories/by/other/app`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static CreateAppAuthForOtherApp(req: CreateAppAuthForOtherAppRequest, initReq?: fm.InitReq): Promise<CreateAppAuthForOtherAppResponse> {
    return fm.fetchReq<CreateAppAuthForOtherAppRequest, CreateAppAuthForOtherAppResponse>(`/v1/create/app/auth/for/other/app`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetAppAuthByAppResourceMethod(req: GetAppAuthByAppResourceMethodRequest, initReq?: fm.InitReq): Promise<GetAppAuthByAppResourceMethodResponse> {
    return fm.fetchReq<GetAppAuthByAppResourceMethodRequest, GetAppAuthByAppResourceMethodResponse>(`/v1/get/app/auth/by/app/resource/method`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetAppAuthByOtherAppResourceMethod(req: GetAppAuthByOtherAppResourceMethodRequest, initReq?: fm.InitReq): Promise<GetAppAuthByOtherAppResourceMethodResponse> {
    return fm.fetchReq<GetAppAuthByOtherAppResourceMethodRequest, GetAppAuthByOtherAppResourceMethodResponse>(`/v1/get/app/auth/by/other/app/resource/method`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static DeleteAppAuth(req: DeleteAppAuthRequest, initReq?: fm.InitReq): Promise<DeleteAppAuthResponse> {
    return fm.fetchReq<DeleteAppAuthRequest, DeleteAppAuthResponse>(`/v1/delete/app/auth`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static CreateAppRoleAuth(req: CreateAppRoleAuthRequest, initReq?: fm.InitReq): Promise<CreateAppRoleAuthResponse> {
    return fm.fetchReq<CreateAppRoleAuthRequest, CreateAppRoleAuthResponse>(`/v1/create/app/role/auth`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static CreateAppRoleAuthForOtherApp(req: CreateAppRoleAuthForOtherAppRequest, initReq?: fm.InitReq): Promise<CreateAppRoleAuthForOtherAppResponse> {
    return fm.fetchReq<CreateAppRoleAuthForOtherAppRequest, CreateAppRoleAuthForOtherAppResponse>(`/v1/create/app/role/auth/for/other/app`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetAppAuthByAppRoleResourceMethod(req: GetAppAuthByAppRoleResourceMethodRequest, initReq?: fm.InitReq): Promise<GetAppAuthByAppRoleResourceMethodResponse> {
    return fm.fetchReq<GetAppAuthByAppRoleResourceMethodRequest, GetAppAuthByAppRoleResourceMethodResponse>(`/v1/get/app/auth/by/app/role/resource/method`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetAppAuthByOtherAppRoleResourceMethod(req: GetAppAuthByOtherAppRoleResourceMethodRequest, initReq?: fm.InitReq): Promise<GetAppAuthByOtherAppRoleResourceMethodResponse> {
    return fm.fetchReq<GetAppAuthByOtherAppRoleResourceMethodRequest, GetAppAuthByOtherAppRoleResourceMethodResponse>(`/v1/get/app/auth/by/other/app/role/resource/method`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static DeleteAppRoleAuth(req: DeleteAppRoleAuthRequest, initReq?: fm.InitReq): Promise<DeleteAppRoleAuthResponse> {
    return fm.fetchReq<DeleteAppRoleAuthRequest, DeleteAppRoleAuthResponse>(`/v1/delete/app/role/auth`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static CreateAppUserAuth(req: CreateAppUserAuthRequest, initReq?: fm.InitReq): Promise<CreateAppUserAuthResponse> {
    return fm.fetchReq<CreateAppUserAuthRequest, CreateAppUserAuthResponse>(`/v1/create/app/user/auth`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static CreateAppUserAuthForOtherApp(req: CreateAppUserAuthForOtherAppRequest, initReq?: fm.InitReq): Promise<CreateAppUserAuthForOtherAppResponse> {
    return fm.fetchReq<CreateAppUserAuthForOtherAppRequest, CreateAppUserAuthForOtherAppResponse>(`/v1/create/app/user/auth/for/other/app`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetAppAuthByAppUserResourceMethod(req: GetAppAuthByAppUserResourceMethodRequest, initReq?: fm.InitReq): Promise<GetAppAuthByAppUserResourceMethodResponse> {
    return fm.fetchReq<GetAppAuthByAppUserResourceMethodRequest, GetAppAuthByAppUserResourceMethodResponse>(`/v1/get/app/auth/by/app/user/resource/method`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetAppAuthByOtherAppUserResourceMethod(req: GetAppAuthByOtherAppUserResourceMethodRequest, initReq?: fm.InitReq): Promise<GetAppAuthByOtherAppUserResourceMethodResponse> {
    return fm.fetchReq<GetAppAuthByOtherAppUserResourceMethodRequest, GetAppAuthByOtherAppUserResourceMethodResponse>(`/v1/get/app/auth/by/other/app/user/resource/method`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static DeleteAppUserAuth(req: DeleteAppUserAuthRequest, initReq?: fm.InitReq): Promise<DeleteAppUserAuthResponse> {
    return fm.fetchReq<DeleteAppUserAuthRequest, DeleteAppUserAuthResponse>(`/v1/delete/app/user/auth`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetAuthsByApp(req: GetAuthsByAppRequest, initReq?: fm.InitReq): Promise<GetAuthsByAppResponse> {
    return fm.fetchReq<GetAuthsByAppRequest, GetAuthsByAppResponse>(`/v1/get/auths/by/app`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetAuthsByOtherApp(req: GetAuthsByOtherAppRequest, initReq?: fm.InitReq): Promise<GetAuthsByOtherAppResponse> {
    return fm.fetchReq<GetAuthsByOtherAppRequest, GetAuthsByOtherAppResponse>(`/v1/get/auths/by/other/app`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetAuthsByAppRole(req: GetAuthsByAppRoleRequest, initReq?: fm.InitReq): Promise<GetAuthsByAppRoleResponse> {
    return fm.fetchReq<GetAuthsByAppRoleRequest, GetAuthsByAppRoleResponse>(`/v1/get/auths/by/app/role`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetAuthsByOtherAppRole(req: GetAuthsByOtherAppRoleRequest, initReq?: fm.InitReq): Promise<GetAuthsByOtherAppRoleResponse> {
    return fm.fetchReq<GetAuthsByOtherAppRoleRequest, GetAuthsByOtherAppRoleResponse>(`/v1/get/auths/by/other/app/role`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetAuthsByAppUser(req: GetAuthsByAppUserRequest, initReq?: fm.InitReq): Promise<GetAuthsByAppUserResponse> {
    return fm.fetchReq<GetAuthsByAppUserRequest, GetAuthsByAppUserResponse>(`/v1/get/auths/by/app/user`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetAuthsByOtherAppUser(req: GetAuthsByOtherAppUserRequest, initReq?: fm.InitReq): Promise<GetAuthsByOtherAppUserResponse> {
    return fm.fetchReq<GetAuthsByOtherAppUserRequest, GetAuthsByOtherAppUserResponse>(`/v1/get/auths/by/other/app/user`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
}