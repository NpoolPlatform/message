/* eslint-disable */
// @ts-nocheck
/*
* This file is a generated Typescript file for GRPC Gateway, DO NOT MODIFY
*/

import * as fm from "../../fetch.pb"
import * as GoogleProtobufEmpty from "../../google/protobuf/empty.pb"
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

export type ThirdLoginRequest = {
  code?: string
  state?: string
}

export type ThirdLoginResponse = {
}

export class ThirdLoginGateway {
  static Version(req: GoogleProtobufEmpty.Empty, initReq?: fm.InitReq): Promise<NpoolV1Npool.VersionResponse> {
    return fm.fetchReq<GoogleProtobufEmpty.Empty, NpoolV1Npool.VersionResponse>(`/version`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetPlatformsByApp(req: GetPlatformsByAppRequest, initReq?: fm.InitReq): Promise<GetPlatformsByAppResponse> {
    return fm.fetchReq<GetPlatformsByAppRequest, GetPlatformsByAppResponse>(`/v1/get/platforms/by/app`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static ThirdRedirect(req: ThirdLoginRequest, initReq?: fm.InitReq): Promise<ThirdLoginResponse> {
    return fm.fetchReq<ThirdLoginRequest, ThirdLoginResponse>(`/v1/third/redirect?${fm.renderURLSearchParams(req, [])}`, {...initReq, method: "GET"})
  }
}