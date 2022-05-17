/* eslint-disable */
// @ts-nocheck
/*
* This file is a generated Typescript file for GRPC Gateway, DO NOT MODIFY
*/

import * as fm from "../../fetch.pb"
import * as GoogleProtobufEmpty from "../../google/protobuf/empty.pb"
import * as AppUserManagerV1Appusermgr from "../appusermgr/appusermgr.pb"
import * as NpoolV1Npool from "../npool.pb"
export type ThirdAuth = {
  id?: string
  appID?: string
  third?: string
  thirdAppKey?: string
  thirdAppSecret?: string
  createAt?: string
  updateAt?: string
  logoUrl?: string
  redirectUrl?: string
}

export type GetThirdAuthByAppRequest = {
  appID?: string
}

export type GetThirdAuthByAppResponse = {
  infos?: Auth[]
}

export type Auth = {
  authUrl?: string
  logoUrl?: string
  third?: string
}

export type AuthLoginRequest = {
  code?: string
  appID?: string
  third?: string
}

export type AuthLoginResponse = {
  info?: AppUserManagerV1Appusermgr.AppUserInfo
}

export class ThirdLoginGateway {
  static Version(req: GoogleProtobufEmpty.Empty, initReq?: fm.InitReq): Promise<NpoolV1Npool.VersionResponse> {
    return fm.fetchReq<GoogleProtobufEmpty.Empty, NpoolV1Npool.VersionResponse>(`/version`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetThirdAuthByApp(req: GetThirdAuthByAppRequest, initReq?: fm.InitReq): Promise<GetThirdAuthByAppResponse> {
    return fm.fetchReq<GetThirdAuthByAppRequest, GetThirdAuthByAppResponse>(`/v1/get/third/auth/by/app`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static AuthLogin(req: AuthLoginRequest, initReq?: fm.InitReq): Promise<AuthLoginResponse> {
    return fm.fetchReq<AuthLoginRequest, AuthLoginResponse>(`/v1/auth/login`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
}