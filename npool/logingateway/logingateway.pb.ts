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
  AppID?: string
  Account?: string
  PasswordHash?: string
  ManMachineSpec?: string
  EnvironmentSpec?: string
  LoginType?: string
}

export type LoginResponse = {
  Info?: AppUserManagerV1Appusermgr.AppUserInfo
}

export type LoginedRequest = {
  AppID?: string
  UserID?: string
}

export type LoginedResponse = {
  Info?: AppUserManagerV1Appusermgr.AppUserInfo
}

export type LogoutRequest = {
  AppID?: string
  UserID?: string
}

export type LogoutResponse = {
  Info?: AppUserManagerV1Appusermgr.AppUserInfo
}

export type RefreshRequest = {
  AppID?: string
  UserID?: string
}

export type RefreshResponse = {
  Info?: AppUserManagerV1Appusermgr.AppUserInfo
}

export type LoginHistory = {
  ID?: string
  AppID?: string
  UserID?: string
  ClientIP?: string
  UserAgent?: string
  CreateAt?: number
}

export type GetLoginHistoriesRequest = {
  AppID?: string
  UserID?: string
}

export type GetLoginHistoriesResponse = {
  Infos?: LoginHistory[]
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
  static Refresh(req: RefreshRequest, initReq?: fm.InitReq): Promise<RefreshResponse> {
    return fm.fetchReq<RefreshRequest, RefreshResponse>(`/v1/refresh`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetLoginHistories(req: GetLoginHistoriesRequest, initReq?: fm.InitReq): Promise<GetLoginHistoriesResponse> {
    return fm.fetchReq<GetLoginHistoriesRequest, GetLoginHistoriesResponse>(`/v1/get/login/histories`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
}