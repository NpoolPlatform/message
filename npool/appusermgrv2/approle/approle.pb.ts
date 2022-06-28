/* eslint-disable */
// @ts-nocheck
/*
* This file is a generated Typescript file for GRPC Gateway, DO NOT MODIFY
*/

import * as fm from "../../../fetch.pb"
import * as GoogleProtobufEmpty from "../../../google/protobuf/empty.pb"
import * as GoogleProtobufStruct from "../../../google/protobuf/struct.pb"
import * as NpoolV1Npool from "../../npool.pb"
export type AppRole = {
  id?: string
  appID?: string
  createdBy?: string
  role?: string
  description?: string
  default?: boolean
}

export type CreateAppRoleRequest = {
  info?: AppRole
}

export type CreateAppRoleResponse = {
  info?: AppRole
}

export type CreateAppRolesRequest = {
  infos?: AppRole[]
}

export type CreateAppRolesResponse = {
  infos?: AppRole[]
}

export type GetAppRoleRequest = {
  id?: string
}

export type GetAppRoleResponse = {
  info?: AppRole
}

export type GetAppRolesRequest = {
  conds?: {[key: string]: NpoolV1Npool.FilterCond}
  offset?: number
  limit?: number
}

export type GetAppRoleOnlyRequest = {
  conds?: {[key: string]: NpoolV1Npool.FilterCond}
}

export type GetAppRoleOnlyResponse = {
  info?: AppRole
}

export type GetAppRolesResponse = {
  infos?: AppRole[]
  total?: number
}

export type UpdateAppRoleRequest = {
  info?: AppRole
}

export type UpdateAppRoleResponse = {
  info?: AppRole
}

export type UpdateAppRoleFieldsRequest = {
  id?: string
  fields?: {[key: string]: GoogleProtobufStruct.Value}
}

export type UpdateAppRoleFieldsResponse = {
  info?: AppRole
}

export type ExistAppRoleRequest = {
  id?: string
}

export type ExistAppRoleResponse = {
  result?: boolean
}

export type ExistAppRoleCondsRequest = {
  conds?: {[key: string]: NpoolV1Npool.FilterCond}
}

export type ExistAppRoleCondsResponse = {
  result?: boolean
}

export type CountAppRolesRequest = {
  conds?: {[key: string]: NpoolV1Npool.FilterCond}
}

export type CountAppRolesResponse = {
  result?: number
}

export type DeleteAppRoleRequest = {
  id?: string
}

export type DeleteAppRoleResponse = {
  info?: AppRole
}

export class AppUserManagerAppRole {
  static Version(req: GoogleProtobufEmpty.Empty, initReq?: fm.InitReq): Promise<NpoolV1Npool.VersionResponse> {
    return fm.fetchReq<GoogleProtobufEmpty.Empty, NpoolV1Npool.VersionResponse>(`/app.user.manager.v2.AppUserManagerAppRole/Version`, {...initReq, method: "POST", body: JSON.stringify(req, fm.replacer)})
  }
  static CreateAppRoleV2(req: CreateAppRoleRequest, initReq?: fm.InitReq): Promise<CreateAppRoleResponse> {
    return fm.fetchReq<CreateAppRoleRequest, CreateAppRoleResponse>(`/app.user.manager.v2.AppUserManagerAppRole/CreateAppRoleV2`, {...initReq, method: "POST", body: JSON.stringify(req, fm.replacer)})
  }
  static CreateAppRolesV2(req: CreateAppRolesRequest, initReq?: fm.InitReq): Promise<CreateAppRolesResponse> {
    return fm.fetchReq<CreateAppRolesRequest, CreateAppRolesResponse>(`/app.user.manager.v2.AppUserManagerAppRole/CreateAppRolesV2`, {...initReq, method: "POST", body: JSON.stringify(req, fm.replacer)})
  }
  static UpdateAppRoleV2(req: UpdateAppRoleRequest, initReq?: fm.InitReq): Promise<UpdateAppRoleResponse> {
    return fm.fetchReq<UpdateAppRoleRequest, UpdateAppRoleResponse>(`/app.user.manager.v2.AppUserManagerAppRole/UpdateAppRoleV2`, {...initReq, method: "POST", body: JSON.stringify(req, fm.replacer)})
  }
  static UpdateAppRoleFieldsV2(req: UpdateAppRoleFieldsRequest, initReq?: fm.InitReq): Promise<UpdateAppRoleFieldsResponse> {
    return fm.fetchReq<UpdateAppRoleFieldsRequest, UpdateAppRoleFieldsResponse>(`/app.user.manager.v2.AppUserManagerAppRole/UpdateAppRoleFieldsV2`, {...initReq, method: "POST", body: JSON.stringify(req, fm.replacer)})
  }
  static GetAppRoleV2(req: GetAppRoleRequest, initReq?: fm.InitReq): Promise<GetAppRoleResponse> {
    return fm.fetchReq<GetAppRoleRequest, GetAppRoleResponse>(`/app.user.manager.v2.AppUserManagerAppRole/GetAppRoleV2`, {...initReq, method: "POST", body: JSON.stringify(req, fm.replacer)})
  }
  static GetAppRoleOnlyV2(req: GetAppRoleOnlyRequest, initReq?: fm.InitReq): Promise<GetAppRoleOnlyResponse> {
    return fm.fetchReq<GetAppRoleOnlyRequest, GetAppRoleOnlyResponse>(`/app.user.manager.v2.AppUserManagerAppRole/GetAppRoleOnlyV2`, {...initReq, method: "POST", body: JSON.stringify(req, fm.replacer)})
  }
  static GetAppRolesV2(req: GetAppRolesRequest, initReq?: fm.InitReq): Promise<GetAppRolesResponse> {
    return fm.fetchReq<GetAppRolesRequest, GetAppRolesResponse>(`/app.user.manager.v2.AppUserManagerAppRole/GetAppRolesV2`, {...initReq, method: "POST", body: JSON.stringify(req, fm.replacer)})
  }
  static ExistAppRoleV2(req: ExistAppRoleRequest, initReq?: fm.InitReq): Promise<ExistAppRoleResponse> {
    return fm.fetchReq<ExistAppRoleRequest, ExistAppRoleResponse>(`/app.user.manager.v2.AppUserManagerAppRole/ExistAppRoleV2`, {...initReq, method: "POST", body: JSON.stringify(req, fm.replacer)})
  }
  static ExistAppRoleCondsV2(req: ExistAppRoleCondsRequest, initReq?: fm.InitReq): Promise<ExistAppRoleCondsResponse> {
    return fm.fetchReq<ExistAppRoleCondsRequest, ExistAppRoleCondsResponse>(`/app.user.manager.v2.AppUserManagerAppRole/ExistAppRoleCondsV2`, {...initReq, method: "POST", body: JSON.stringify(req, fm.replacer)})
  }
  static CountAppRolesV2(req: CountAppRolesRequest, initReq?: fm.InitReq): Promise<CountAppRolesResponse> {
    return fm.fetchReq<CountAppRolesRequest, CountAppRolesResponse>(`/app.user.manager.v2.AppUserManagerAppRole/CountAppRolesV2`, {...initReq, method: "POST", body: JSON.stringify(req, fm.replacer)})
  }
  static DeleteAppRoleV2(req: DeleteAppRoleRequest, initReq?: fm.InitReq): Promise<DeleteAppRoleResponse> {
    return fm.fetchReq<DeleteAppRoleRequest, DeleteAppRoleResponse>(`/app.user.manager.v2.AppUserManagerAppRole/DeleteAppRoleV2`, {...initReq, method: "POST", body: JSON.stringify(req, fm.replacer)})
  }
}