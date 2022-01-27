/* eslint-disable */
// @ts-nocheck
/*
* This file is a generated Typescript file for GRPC Gateway, DO NOT MODIFY
*/

import * as fm from "../../fetch.pb"
import * as GoogleProtobufEmpty from "../../google/protobuf/empty.pb"
import * as NpoolV1Npool from "../npool.pb"
export type AuthByAppRequest = {
  AppID?: string
}

export type AuthByAppResponse = {
  Allowed?: boolean
}

export type AuthByAppRoleUserRequest = {
  AppID?: string
  UserID?: string
  Token?: string
}

export type AuthByAppRoleUserResponse = {
  Allowed?: boolean
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
}