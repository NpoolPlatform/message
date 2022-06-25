/* eslint-disable */
// @ts-nocheck
/*
* This file is a generated Typescript file for GRPC Gateway, DO NOT MODIFY
*/

import * as fm from "../../fetch.pb"
import * as GoogleProtobufEmpty from "../../google/protobuf/empty.pb"
import * as NpoolV1Npool from "../npool.pb"
export type Path = {
  method?: string
  path?: string
  exported?: boolean
  methodName?: string
}

export type ServiceApis = {
  protocol?: string
  serviceName?: string
  pathPrefix?: string
  paths?: Path[]
}

export type RegisterRequest = {
  info?: ServiceApis
}

export type RegisterResponse = {
  info?: ServiceApis
}

export type ServicePath = {
  id?: string
  protocol?: string
  serviceName?: string
  domains?: string[]
  pathPrefix?: string
  method?: string
  path?: string
  exported?: boolean
  createAt?: number
  updateAt?: number
  methodName?: string
}

export type GetApisRequest = {
}

export type GetApisResponse = {
  infos?: ServicePath[]
}

export type GetServiceMethodAPIRequest = {
  serviceName?: string
  methodName?: string
}

export type GetServiceMethodAPIResponse = {
  info?: ServicePath
}

export class ApiManager {
  static Version(req: GoogleProtobufEmpty.Empty, initReq?: fm.InitReq): Promise<NpoolV1Npool.VersionResponse> {
    return fm.fetchReq<GoogleProtobufEmpty.Empty, NpoolV1Npool.VersionResponse>(`/version`, {...initReq, method: "POST", body: JSON.stringify(req, fm.replacer)})
  }
  static Register(req: RegisterRequest, initReq?: fm.InitReq): Promise<RegisterResponse> {
    return fm.fetchReq<RegisterRequest, RegisterResponse>(`/v1/register`, {...initReq, method: "POST", body: JSON.stringify(req, fm.replacer)})
  }
  static GetApis(req: GetApisRequest, initReq?: fm.InitReq): Promise<GetApisResponse> {
    return fm.fetchReq<GetApisRequest, GetApisResponse>(`/v1/get/apis`, {...initReq, method: "POST", body: JSON.stringify(req, fm.replacer)})
  }
  static GetServiceMethodAPI(req: GetServiceMethodAPIRequest, initReq?: fm.InitReq): Promise<GetServiceMethodAPIResponse> {
    return fm.fetchReq<GetServiceMethodAPIRequest, GetServiceMethodAPIResponse>(`/v1/get/service/method/api`, {...initReq, method: "POST", body: JSON.stringify(req, fm.replacer)})
  }
}