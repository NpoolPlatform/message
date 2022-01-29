/* eslint-disable */
// @ts-nocheck
/*
* This file is a generated Typescript file for GRPC Gateway, DO NOT MODIFY
*/

import * as fm from "../../fetch.pb"
import * as GoogleProtobufEmpty from "../../google/protobuf/empty.pb"
import * as AppUserManagerV1Appusermgr from "../appusermgr/appusermgr.pb"
import * as NpoolV1Npool from "../npool.pb"
export type LoginRequest = {
  appID?: string
  account?: string
  passwordHash?: string
  manMachineSpec?: string
  environmentSpec?: string
  loginType?: string
  token?: string
}

export type LoginResponse = {
  info?: AppUserManagerV1Appusermgr.AppUserInfo
  token?: string
}

export type LoginedRequest = {
  appID?: string
  userID?: string
  token?: string
}

export type LoginedResponse = {
  info?: AppUserManagerV1Appusermgr.AppUserInfo
}

export type LogoutRequest = {
  appID?: string
  userID?: string
}

export type LogoutResponse = {
  info?: AppUserManagerV1Appusermgr.AppUserInfo
}

export type LoginHistory = {
  id?: string
  appID?: string
  userID?: string
  clientIP?: string
  userAgent?: string
  createAt?: number
}

export type GetLoginHistoriesRequest = {
  appID?: string
  userID?: string
  pageInfo?: NpoolV1Npool.PageInfo
}

export type GetLoginHistoriesResponse = {
  infos?: LoginHistory[]
}

export class LoginGateway {
  static Version(req: GoogleProtobufEmpty.Empty, initReq?: fm.InitReq): Promise<NpoolV1Npool.VersionResponse> {
    return fm.fetchReq<GoogleProtobufEmpty.Empty, NpoolV1Npool.VersionResponse>(`/version`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static Login(req: LoginRequest, initReq?: fm.InitReq): Promise<LoginResponse> {
    return fm.fetchReq<LoginRequest, LoginResponse>(`/v1/login`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static Logined(req: LoginedRequest, initReq?: fm.InitReq): Promise<LoginedResponse> {
    return fm.fetchReq<LoginedRequest, LoginedResponse>(`/v1/logined`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static Logout(req: LogoutRequest, initReq?: fm.InitReq): Promise<LogoutResponse> {
    return fm.fetchReq<LogoutRequest, LogoutResponse>(`/v1/logout`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetLoginHistories(req: GetLoginHistoriesRequest, initReq?: fm.InitReq): Promise<GetLoginHistoriesResponse> {
    return fm.fetchReq<GetLoginHistoriesRequest, GetLoginHistoriesResponse>(`/v1/get/login/histories`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
}