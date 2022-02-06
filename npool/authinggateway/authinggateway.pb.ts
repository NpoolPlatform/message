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

export type GetAuthHistoriesByOtherAppRequest = {
  targetAppID?: string
}

export type GetAuthHistoriesByOtherAppResponse = {
  infos?: AuthHistory[]
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
  static GetAuthHistoriesByOtherApp(req: GetAuthHistoriesByOtherAppRequest, initReq?: fm.InitReq): Promise<GetAuthHistoriesByOtherAppResponse> {
    return fm.fetchReq<GetAuthHistoriesByOtherAppRequest, GetAuthHistoriesByOtherAppResponse>(`/v1/get/auth/histories/by/other/app`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
}