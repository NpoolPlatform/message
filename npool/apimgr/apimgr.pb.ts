/* eslint-disable */
// @ts-nocheck
/*
* This file is a generated Typescript file for GRPC Gateway, DO NOT MODIFY
*/

import * as fm from "../../fetch.pb"
import * as GoogleProtobufEmpty from "../../google/protobuf/empty.pb"
import * as NpoolV1Npool from "../npool.pb"
export type Path = {
  Method?: string
  Path?: string
  Exported?: boolean
}

export type ServiceApis = {
  ServiceName?: string
  Paths?: Path[]
}

export type RegisterRequest = {
  Info?: ServiceApis
}

export type RegisterResponse = {
  Info?: ServiceApis
}

export type GetApisRequest = {
}

export type GetApisResponse = {
  Infos?: ServiceApis[]
}

export class ApiManager {
  static Version(req: GoogleProtobufEmpty.Empty, initReq?: fm.InitReq): Promise<NpoolV1Npool.VersionResponse> {
    return fm.fetchReq<GoogleProtobufEmpty.Empty, NpoolV1Npool.VersionResponse>(`/version`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static Register(req: RegisterRequest, initReq?: fm.InitReq): Promise<RegisterResponse> {
    return fm.fetchReq<RegisterRequest, RegisterResponse>(`/v1/register`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetApis(req: GetApisRequest, initReq?: fm.InitReq): Promise<GetApisResponse> {
    return fm.fetchReq<GetApisRequest, GetApisResponse>(`/v1/get/apis`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
}