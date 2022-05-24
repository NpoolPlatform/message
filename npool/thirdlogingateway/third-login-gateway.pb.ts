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
  createdAt?: string
  updateAt?: string
  logoUrl?: string
  redirectUrl?: string
}

export type GetAuthsRequest = {
  appID?: string
}

export type GetAuthsResponse = {
  infos?: Auth[]
}

export type GetAuthsByAppRequest = {
  targetAppID?: string
}

export type GetAuthsByAppResponse = {
  infos?: Auth[]
}

export type CreateAuthRequest = {
  appID?: string
  info?: ThirdAuth
}

export type CreateAuthResponse = {
  info?: Auth
}

export type CreateAuthsRequest = {
  appID?: string
  infos?: ThirdAuth[]
}

export type CreateAuthsResponse = {
  infos?: Auth[]
}

export type CreateAppAuthRequest = {
  targetAppID?: string
  info?: ThirdAuth
}

export type CreateAppAuthResponse = {
  info?: Auth
}

export type CreateAppAuthsRequest = {
  targetAppID?: string
  infos?: ThirdAuth[]
}

export type CreateAppAuthsResponse = {
  infos?: Auth[]
}

export type Auth = {
  authUrl?: string
  logoUrl?: string
  third?: string
}

export type LoginRequest = {
  code?: string
  appID?: string
  third?: string
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
  static GetAuthsByApp(req: GetAuthsByAppRequest, initReq?: fm.InitReq): Promise<GetAuthsByAppResponse> {
    return fm.fetchReq<GetAuthsByAppRequest, GetAuthsByAppResponse>(`/v1/get/auths/by/app`, {...initReq, method: "POST", body: JSON.stringify(req)})
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
  static Login(req: LoginRequest, initReq?: fm.InitReq): Promise<LoginResponse> {
    return fm.fetchReq<LoginRequest, LoginResponse>(`/v1/login`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
}