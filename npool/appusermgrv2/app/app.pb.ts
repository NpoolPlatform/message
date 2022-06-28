/* eslint-disable */
// @ts-nocheck
/*
* This file is a generated Typescript file for GRPC Gateway, DO NOT MODIFY
*/

import * as fm from "../../../fetch.pb"
import * as GoogleProtobufEmpty from "../../../google/protobuf/empty.pb"
import * as GoogleProtobufStruct from "../../../google/protobuf/struct.pb"
import * as NpoolV1Npool from "../../npool.pb"
export type App = {
  id?: string
  createdBy?: string
  name?: string
  logo?: string
  description?: string
  createAt?: number
}

export type CreateAppRequest = {
  info?: App
}

export type CreateAppResponse = {
  info?: App
}

export type CreateAppsRequest = {
  infos?: App[]
}

export type CreateAppsResponse = {
  infos?: App[]
}

export type GetAppRequest = {
  id?: string
}

export type GetAppResponse = {
  info?: App
}

export type GetAppsRequest = {
  conds?: {[key: string]: NpoolV1Npool.FilterCond}
  offset?: number
  limit?: number
}

export type GetAppOnlyRequest = {
  conds?: {[key: string]: NpoolV1Npool.FilterCond}
}

export type GetAppOnlyResponse = {
  info?: App
}

export type GetAppsResponse = {
  infos?: App[]
  total?: number
}

export type UpdateAppRequest = {
  info?: App
}

export type UpdateAppResponse = {
  info?: App
}

export type UpdateAppFieldsRequest = {
  id?: string
  fields?: {[key: string]: GoogleProtobufStruct.Value}
}

export type UpdateAppFieldsResponse = {
  info?: App
}

export type ExistAppRequest = {
  id?: string
}

export type ExistAppResponse = {
  result?: boolean
}

export type ExistAppCondsRequest = {
  conds?: {[key: string]: NpoolV1Npool.FilterCond}
}

export type ExistAppCondsResponse = {
  result?: boolean
}

export type CountAppsRequest = {
  conds?: {[key: string]: NpoolV1Npool.FilterCond}
}

export type CountAppsResponse = {
  result?: number
}

export type DeleteAppRequest = {
  id?: string
}

export type DeleteAppResponse = {
  info?: App
}

export class AppUserManagerApp {
  static Version(req: GoogleProtobufEmpty.Empty, initReq?: fm.InitReq): Promise<NpoolV1Npool.VersionResponse> {
    return fm.fetchReq<GoogleProtobufEmpty.Empty, NpoolV1Npool.VersionResponse>(`/app.user.manager.v2.AppUserManagerApp/Version`, {...initReq, method: "POST", body: JSON.stringify(req, fm.replacer)})
  }
  static CreateAppV2(req: CreateAppRequest, initReq?: fm.InitReq): Promise<CreateAppResponse> {
    return fm.fetchReq<CreateAppRequest, CreateAppResponse>(`/app.user.manager.v2.AppUserManagerApp/CreateAppV2`, {...initReq, method: "POST", body: JSON.stringify(req, fm.replacer)})
  }
  static CreateAppsV2(req: CreateAppsRequest, initReq?: fm.InitReq): Promise<CreateAppsResponse> {
    return fm.fetchReq<CreateAppsRequest, CreateAppsResponse>(`/app.user.manager.v2.AppUserManagerApp/CreateAppsV2`, {...initReq, method: "POST", body: JSON.stringify(req, fm.replacer)})
  }
  static UpdateAppV2(req: UpdateAppRequest, initReq?: fm.InitReq): Promise<UpdateAppResponse> {
    return fm.fetchReq<UpdateAppRequest, UpdateAppResponse>(`/app.user.manager.v2.AppUserManagerApp/UpdateAppV2`, {...initReq, method: "POST", body: JSON.stringify(req, fm.replacer)})
  }
  static UpdateAppFieldsV2(req: UpdateAppFieldsRequest, initReq?: fm.InitReq): Promise<UpdateAppFieldsResponse> {
    return fm.fetchReq<UpdateAppFieldsRequest, UpdateAppFieldsResponse>(`/app.user.manager.v2.AppUserManagerApp/UpdateAppFieldsV2`, {...initReq, method: "POST", body: JSON.stringify(req, fm.replacer)})
  }
  static GetAppV2(req: GetAppRequest, initReq?: fm.InitReq): Promise<GetAppResponse> {
    return fm.fetchReq<GetAppRequest, GetAppResponse>(`/app.user.manager.v2.AppUserManagerApp/GetAppV2`, {...initReq, method: "POST", body: JSON.stringify(req, fm.replacer)})
  }
  static GetAppOnlyV2(req: GetAppOnlyRequest, initReq?: fm.InitReq): Promise<GetAppOnlyResponse> {
    return fm.fetchReq<GetAppOnlyRequest, GetAppOnlyResponse>(`/app.user.manager.v2.AppUserManagerApp/GetAppOnlyV2`, {...initReq, method: "POST", body: JSON.stringify(req, fm.replacer)})
  }
  static GetAppsV2(req: GetAppsRequest, initReq?: fm.InitReq): Promise<GetAppsResponse> {
    return fm.fetchReq<GetAppsRequest, GetAppsResponse>(`/app.user.manager.v2.AppUserManagerApp/GetAppsV2`, {...initReq, method: "POST", body: JSON.stringify(req, fm.replacer)})
  }
  static ExistAppV2(req: ExistAppRequest, initReq?: fm.InitReq): Promise<ExistAppResponse> {
    return fm.fetchReq<ExistAppRequest, ExistAppResponse>(`/app.user.manager.v2.AppUserManagerApp/ExistAppV2`, {...initReq, method: "POST", body: JSON.stringify(req, fm.replacer)})
  }
  static ExistAppCondsV2(req: ExistAppCondsRequest, initReq?: fm.InitReq): Promise<ExistAppCondsResponse> {
    return fm.fetchReq<ExistAppCondsRequest, ExistAppCondsResponse>(`/app.user.manager.v2.AppUserManagerApp/ExistAppCondsV2`, {...initReq, method: "POST", body: JSON.stringify(req, fm.replacer)})
  }
  static CountAppsV2(req: CountAppsRequest, initReq?: fm.InitReq): Promise<CountAppsResponse> {
    return fm.fetchReq<CountAppsRequest, CountAppsResponse>(`/app.user.manager.v2.AppUserManagerApp/CountAppsV2`, {...initReq, method: "POST", body: JSON.stringify(req, fm.replacer)})
  }
  static DeleteAppV2(req: DeleteAppRequest, initReq?: fm.InitReq): Promise<DeleteAppResponse> {
    return fm.fetchReq<DeleteAppRequest, DeleteAppResponse>(`/app.user.manager.v2.AppUserManagerApp/DeleteAppV2`, {...initReq, method: "POST", body: JSON.stringify(req, fm.replacer)})
  }
}