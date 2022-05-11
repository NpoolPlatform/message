/* eslint-disable */
// @ts-nocheck
/*
* This file is a generated Typescript file for GRPC Gateway, DO NOT MODIFY
*/

import * as fm from "../../fetch.pb"
import * as GoogleProtobufEmpty from "../../google/protobuf/empty.pb"
import * as AppUserManagerV1Appusermgr from "../appusermgr/appusermgr.pb"
import * as NpoolV1Npool from "../npool.pb"
export type Platform = {
  id?: string
  appID?: string
  platform?: string
  platformAppKey?: string
  platformAppSecret?: string
  createAt?: string
  updateAt?: string
  logoUrl?: string
  redirectUrl?: string
}

export type GetPlatformsByAppRequest = {
  appID?: string
}

export type GetPlatformsByAppResponse = {
  infos?: Auth[]
}

export type Auth = {
  authUrl?: string
  logoUrl?: string
  platform?: string
}

export type AuthLoginRequest = {
  code?: string
  appID?: string
  platform?: string
}

export type AuthLoginResponse = {
  info?: AppUserManagerV1Appusermgr.AppUserInfo
  token?: string
}

export class ThirdLoginGateway {
  static Version(req: GoogleProtobufEmpty.Empty, initReq?: fm.InitReq): Promise<NpoolV1Npool.VersionResponse> {
    return fm.fetchReq<GoogleProtobufEmpty.Empty, NpoolV1Npool.VersionResponse>(`/version`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetPlatformsByApp(req: GetPlatformsByAppRequest, initReq?: fm.InitReq): Promise<GetPlatformsByAppResponse> {
    return fm.fetchReq<GetPlatformsByAppRequest, GetPlatformsByAppResponse>(`/v1/get/platforms/by/app`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static AuthLogin(req: AuthLoginRequest, initReq?: fm.InitReq): Promise<AuthLoginResponse> {
    return fm.fetchReq<AuthLoginRequest, AuthLoginResponse>(`/v1/auth/login`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
}