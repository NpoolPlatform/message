/* eslint-disable */
// @ts-nocheck
/*
* This file is a generated Typescript file for GRPC Gateway, DO NOT MODIFY
*/

import * as fm from "../../fetch.pb"
import * as GoogleProtobufEmpty from "../../google/protobuf/empty.pb"
import * as AppUserManagerV1Appusermgr from "../appusermgr/appusermgr.pb"
import * as NpoolV1Npool from "../npool.pb"
export type ThirdParty = {
  id?: string
  brandName?: string
  logo?: string
  domain?: string
}

export type CreateThirdPartyRequest = {
  info?: ThirdParty
}

export type CreateThirdPartyResponse = {
  info?: ThirdParty
}

export type UpdateThirdPartyRequest = {
  info?: ThirdParty
}

export type UpdateThirdPartyResponse = {
  info?: ThirdParty
}

export type GetThirdPartiesRequest = {
}

export type GetThirdPartiesResponse = {
  infos?: ThirdParty[]
}

export type Auth = {
  id?: string
  appID?: string
  thirdPartyID?: string
  appKey?: string
  appSecret?: string
  redirectURL?: string
  authURL?: string
}

export type GetAuthsRequest = {
  appID?: string
}

export type GetAuthsResponse = {
  infos?: Auth[]
}

export type GetAppAuthsRequest = {
  targetAppID?: string
}

export type GetAppAuthsResponse = {
  infos?: Auth[]
}

export type CreateAuthRequest = {
  info?: Auth
}

export type CreateAuthResponse = {
  info?: Auth
}

export type CreateAuthsRequest = {
  appID?: string
  infos?: Auth[]
}

export type CreateAuthsResponse = {
  infos?: Auth[]
}

export type CreateAppAuthRequest = {
  targetAppID?: string
  info?: Auth
}

export type CreateAppAuthResponse = {
  info?: Auth
}

export type CreateAppAuthsRequest = {
  targetAppID?: string
  infos?: Auth[]
}

export type CreateAppAuthsResponse = {
  infos?: Auth[]
}

export type LoginRequest = {
  code?: string
  appID?: string
  thirdPartyID?: string
}

export type LoginResponse = {
  info?: AppUserManagerV1Appusermgr.AppUserInfo
}

export class ThirdLoginGateway {
  static Version(req: GoogleProtobufEmpty.Empty, initReq?: fm.InitReq): Promise<NpoolV1Npool.VersionResponse> {
    return fm.fetchReq<GoogleProtobufEmpty.Empty, NpoolV1Npool.VersionResponse>(`/version`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetAuths(req: GetAuthsRequest, initReq?: fm.InitReq): Promise<GetAuthsResponse> {
    return fm.fetchReq<GetAuthsRequest, GetAuthsResponse>(`/v1/get/auths`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetAppAuths(req: GetAppAuthsRequest, initReq?: fm.InitReq): Promise<GetAppAuthsResponse> {
    return fm.fetchReq<GetAppAuthsRequest, GetAppAuthsResponse>(`/v1/get/app/auths`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static CreateAuth(req: CreateAuthRequest, initReq?: fm.InitReq): Promise<CreateAuthResponse> {
    return fm.fetchReq<CreateAuthRequest, CreateAuthResponse>(`/v1/create/auth`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static CreateAuths(req: CreateAuthsRequest, initReq?: fm.InitReq): Promise<CreateAuthsResponse> {
    return fm.fetchReq<CreateAuthsRequest, CreateAuthsResponse>(`/v1/create/auths`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static CreateAppAuth(req: CreateAppAuthRequest, initReq?: fm.InitReq): Promise<CreateAppAuthResponse> {
    return fm.fetchReq<CreateAppAuthRequest, CreateAppAuthResponse>(`/v1/create/app/auth`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static CreateAppAuths(req: CreateAppAuthsRequest, initReq?: fm.InitReq): Promise<CreateAppAuthsResponse> {
    return fm.fetchReq<CreateAppAuthsRequest, CreateAppAuthsResponse>(`/v1/create/app/auths`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static CreateThirdParty(req: CreateThirdPartyRequest, initReq?: fm.InitReq): Promise<CreateThirdPartyResponse> {
    return fm.fetchReq<CreateThirdPartyRequest, CreateThirdPartyResponse>(`/v1/create/third/party`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static UpdateThirdParty(req: UpdateThirdPartyRequest, initReq?: fm.InitReq): Promise<UpdateThirdPartyResponse> {
    return fm.fetchReq<UpdateThirdPartyRequest, UpdateThirdPartyResponse>(`/v1/update/third/party`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetThirdParties(req: GetThirdPartiesRequest, initReq?: fm.InitReq): Promise<GetThirdPartiesResponse> {
    return fm.fetchReq<GetThirdPartiesRequest, GetThirdPartiesResponse>(`/v1/get/third/parties`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static Login(req: LoginRequest, initReq?: fm.InitReq): Promise<LoginResponse> {
    return fm.fetchReq<LoginRequest, LoginResponse>(`/v1/login`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
}