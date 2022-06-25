/* eslint-disable */
// @ts-nocheck
/*
* This file is a generated Typescript file for GRPC Gateway, DO NOT MODIFY
*/

import * as fm from "../../fetch.pb"
import * as GoogleProtobufEmpty from "../../google/protobuf/empty.pb"
import * as NpoolV1Npool from "../npool.pb"
export type CoinDescription = {
  id?: string
  appID?: string
  coinTypeID?: string
  title?: string
  message?: string
  usedFor?: string
  createAt?: number
  updateAt?: number
}

export type CreateCoinDescriptionRequest = {
  info?: CoinDescription
}

export type CreateCoinDescriptionResponse = {
  info?: CoinDescription
}

export type CreateCoinDescriptionsRequest = {
  appID?: string
  infos?: CoinDescription[]
}

export type CreateCoinDescriptionsResponse = {
  infos?: CoinDescription[]
}

export type CreateAppCoinDescriptionRequest = {
  targetAppID?: string
  info?: CoinDescription
}

export type CreateAppCoinDescriptionResponse = {
  info?: CoinDescription
}

export type CreateAppCoinDescriptionsRequest = {
  targetAppID?: string
  infos?: CoinDescription[]
}

export type CreateAppCoinDescriptionsResponse = {
  infos?: CoinDescription[]
}

export type UpdateCoinDescriptionRequest = {
  info?: CoinDescription
}

export type UpdateCoinDescriptionResponse = {
  info?: CoinDescription
}

export type GetCoinDescriptionRequest = {
  id?: string
}

export type GetCoinDescriptionResponse = {
  info?: CoinDescription
}

export type GetCoinDescriptionsRequest = {
  appID?: string
  conds?: {[key: string]: NpoolV1Npool.FilterCond}
  offset?: number
  limit?: number
}

export type GetCoinDescriptionsResponse = {
  infos?: CoinDescription[]
  total?: number
}

export type GetCoinDescriptionOnlyRequest = {
  appID?: string
  conds?: {[key: string]: NpoolV1Npool.FilterCond}
  offset?: number
  limit?: number
}

export type GetCoinDescriptionOnlyResponse = {
  info?: CoinDescription
}

export type GetAppCoinDescriptionsRequest = {
  targetAppID?: string
  conds?: {[key: string]: NpoolV1Npool.FilterCond}
  offset?: number
  limit?: number
}

export type GetAppCoinDescriptionsResponse = {
  infos?: CoinDescription[]
  total?: number
}

export type GetAppCoinDescriptionOnlyRequest = {
  targetAppID?: string
  conds?: {[key: string]: NpoolV1Npool.FilterCond}
  offset?: number
  limit?: number
}

export type GetAppCoinDescriptionOnlyResponse = {
  info?: CoinDescription
}

export type DeleteCoinDescriptionRequest = {
  id?: string
}

export type DeleteCoinDescriptionResponse = {
  info?: CoinDescription
}

export type CoinProductInfo = {
  id?: string
  appID?: string
  coinTypeID?: string
  productPage?: string
}

export type CreateCoinProductInfoRequest = {
  info?: CoinProductInfo
}

export type CreateCoinProductInfoResponse = {
  info?: CoinProductInfo
}

export type CreateCoinProductInfosRequest = {
  appID?: string
  infos?: CoinProductInfo[]
}

export type CreateCoinProductInfosResponse = {
  infos?: CoinProductInfo[]
}

export type CreateAppCoinProductInfoRequest = {
  targetAppID?: string
  info?: CoinProductInfo
}

export type CreateAppCoinProductInfoResponse = {
  info?: CoinProductInfo
}

export type CreateAppCoinProductInfosRequest = {
  targetAppID?: string
  infos?: CoinProductInfo[]
}

export type CreateAppCoinProductInfosResponse = {
  infos?: CoinProductInfo[]
}

export type UpdateCoinProductInfoRequest = {
  info?: CoinProductInfo
}

export type UpdateCoinProductInfoResponse = {
  info?: CoinProductInfo
}

export type GetCoinProductInfoRequest = {
  id?: string
}

export type GetCoinProductInfoResponse = {
  info?: CoinProductInfo
}

export type GetCoinProductInfosRequest = {
  appID?: string
  conds?: {[key: string]: NpoolV1Npool.FilterCond}
  offset?: number
  limit?: number
}

export type GetCoinProductInfosResponse = {
  infos?: CoinProductInfo[]
  total?: number
}

export type GetCoinProductInfoOnlyRequest = {
  appID?: string
  conds?: {[key: string]: NpoolV1Npool.FilterCond}
  offset?: number
  limit?: number
}

export type GetCoinProductInfoOnlyResponse = {
  info?: CoinProductInfo
}

export type GetAppCoinProductInfosRequest = {
  targetAppID?: string
  conds?: {[key: string]: NpoolV1Npool.FilterCond}
  offset?: number
  limit?: number
}

export type GetAppCoinProductInfosResponse = {
  infos?: CoinProductInfo[]
  total?: number
}

export type GetAppCoinProductInfoOnlyRequest = {
  targetAppID?: string
  conds?: {[key: string]: NpoolV1Npool.FilterCond}
  offset?: number
  limit?: number
}

export type GetAppCoinProductInfoOnlyResponse = {
  info?: CoinProductInfo
}

export type DeleteCoinProductInfoRequest = {
  id?: string
}

export type DeleteCoinProductInfoResponse = {
  info?: CoinProductInfo
}

export class ProjectInfoManager {
  static Version(req: GoogleProtobufEmpty.Empty, initReq?: fm.InitReq): Promise<NpoolV1Npool.VersionResponse> {
    return fm.fetchReq<GoogleProtobufEmpty.Empty, NpoolV1Npool.VersionResponse>(`/version`, {...initReq, method: "POST", body: JSON.stringify(req, fm.replacer)})
  }
  static CreateCoinDescription(req: CreateCoinDescriptionRequest, initReq?: fm.InitReq): Promise<CreateCoinDescriptionResponse> {
    return fm.fetchReq<CreateCoinDescriptionRequest, CreateCoinDescriptionResponse>(`/v1/create/coin/description`, {...initReq, method: "POST", body: JSON.stringify(req, fm.replacer)})
  }
  static CreateCoinDescriptions(req: CreateCoinDescriptionsRequest, initReq?: fm.InitReq): Promise<CreateCoinDescriptionsResponse> {
    return fm.fetchReq<CreateCoinDescriptionsRequest, CreateCoinDescriptionsResponse>(`/v1/create/coin/descriptions`, {...initReq, method: "POST", body: JSON.stringify(req, fm.replacer)})
  }
  static CreateAppCoinDescription(req: CreateAppCoinDescriptionRequest, initReq?: fm.InitReq): Promise<CreateAppCoinDescriptionResponse> {
    return fm.fetchReq<CreateAppCoinDescriptionRequest, CreateAppCoinDescriptionResponse>(`/v1/create/app/coin/description`, {...initReq, method: "POST", body: JSON.stringify(req, fm.replacer)})
  }
  static CreateAppCoinDescriptions(req: CreateAppCoinDescriptionsRequest, initReq?: fm.InitReq): Promise<CreateAppCoinDescriptionsResponse> {
    return fm.fetchReq<CreateAppCoinDescriptionsRequest, CreateAppCoinDescriptionsResponse>(`/v1/create/app/coin/descriptions`, {...initReq, method: "POST", body: JSON.stringify(req, fm.replacer)})
  }
  static UpdateCoinDescription(req: UpdateCoinDescriptionRequest, initReq?: fm.InitReq): Promise<UpdateCoinDescriptionResponse> {
    return fm.fetchReq<UpdateCoinDescriptionRequest, UpdateCoinDescriptionResponse>(`/v1/update/coin/description`, {...initReq, method: "POST", body: JSON.stringify(req, fm.replacer)})
  }
  static GetCoinDescription(req: GetCoinDescriptionRequest, initReq?: fm.InitReq): Promise<GetCoinDescriptionResponse> {
    return fm.fetchReq<GetCoinDescriptionRequest, GetCoinDescriptionResponse>(`/v1/get/coin/description`, {...initReq, method: "POST", body: JSON.stringify(req, fm.replacer)})
  }
  static GetCoinDescriptions(req: GetCoinDescriptionsRequest, initReq?: fm.InitReq): Promise<GetCoinDescriptionsResponse> {
    return fm.fetchReq<GetCoinDescriptionsRequest, GetCoinDescriptionsResponse>(`/v1/get/coin/descriptions`, {...initReq, method: "POST", body: JSON.stringify(req, fm.replacer)})
  }
  static GetCoinDescriptionOnly(req: GetCoinDescriptionOnlyRequest, initReq?: fm.InitReq): Promise<GetCoinDescriptionOnlyResponse> {
    return fm.fetchReq<GetCoinDescriptionOnlyRequest, GetCoinDescriptionOnlyResponse>(`/v1/get/coin/description/only`, {...initReq, method: "POST", body: JSON.stringify(req, fm.replacer)})
  }
  static GetAppCoinDescriptions(req: GetAppCoinDescriptionsRequest, initReq?: fm.InitReq): Promise<GetAppCoinDescriptionsResponse> {
    return fm.fetchReq<GetAppCoinDescriptionsRequest, GetAppCoinDescriptionsResponse>(`/v1/get/app/coin/descriptions`, {...initReq, method: "POST", body: JSON.stringify(req, fm.replacer)})
  }
  static GetAppCoinDescriptionOnly(req: GetAppCoinDescriptionOnlyRequest, initReq?: fm.InitReq): Promise<GetAppCoinDescriptionOnlyResponse> {
    return fm.fetchReq<GetAppCoinDescriptionOnlyRequest, GetAppCoinDescriptionOnlyResponse>(`/v1/get/app/coin/description/only`, {...initReq, method: "POST", body: JSON.stringify(req, fm.replacer)})
  }
  static DeleteCoinDescription(req: DeleteCoinDescriptionRequest, initReq?: fm.InitReq): Promise<DeleteCoinDescriptionResponse> {
    return fm.fetchReq<DeleteCoinDescriptionRequest, DeleteCoinDescriptionResponse>(`/v1/delete/coin/description`, {...initReq, method: "POST", body: JSON.stringify(req, fm.replacer)})
  }
  static CreateCoinProductInfo(req: CreateCoinProductInfoRequest, initReq?: fm.InitReq): Promise<CreateCoinProductInfoResponse> {
    return fm.fetchReq<CreateCoinProductInfoRequest, CreateCoinProductInfoResponse>(`/v1/create/coin/productinfo`, {...initReq, method: "POST", body: JSON.stringify(req, fm.replacer)})
  }
  static CreateCoinProductInfos(req: CreateCoinProductInfosRequest, initReq?: fm.InitReq): Promise<CreateCoinProductInfosResponse> {
    return fm.fetchReq<CreateCoinProductInfosRequest, CreateCoinProductInfosResponse>(`/v1/create/coin/productinfos`, {...initReq, method: "POST", body: JSON.stringify(req, fm.replacer)})
  }
  static CreateAppCoinProductInfo(req: CreateAppCoinProductInfoRequest, initReq?: fm.InitReq): Promise<CreateAppCoinProductInfoResponse> {
    return fm.fetchReq<CreateAppCoinProductInfoRequest, CreateAppCoinProductInfoResponse>(`/v1/create/app/coin/productinfo`, {...initReq, method: "POST", body: JSON.stringify(req, fm.replacer)})
  }
  static CreateAppCoinProductInfos(req: CreateAppCoinProductInfosRequest, initReq?: fm.InitReq): Promise<CreateAppCoinProductInfosResponse> {
    return fm.fetchReq<CreateAppCoinProductInfosRequest, CreateAppCoinProductInfosResponse>(`/v1/create/app/coin/productinfos`, {...initReq, method: "POST", body: JSON.stringify(req, fm.replacer)})
  }
  static UpdateCoinProductInfo(req: UpdateCoinProductInfoRequest, initReq?: fm.InitReq): Promise<UpdateCoinProductInfoResponse> {
    return fm.fetchReq<UpdateCoinProductInfoRequest, UpdateCoinProductInfoResponse>(`/v1/update/coin/productinfo`, {...initReq, method: "POST", body: JSON.stringify(req, fm.replacer)})
  }
  static GetCoinProductInfo(req: GetCoinProductInfoRequest, initReq?: fm.InitReq): Promise<GetCoinProductInfoResponse> {
    return fm.fetchReq<GetCoinProductInfoRequest, GetCoinProductInfoResponse>(`/v1/get/coin/productinfo`, {...initReq, method: "POST", body: JSON.stringify(req, fm.replacer)})
  }
  static GetCoinProductInfos(req: GetCoinProductInfosRequest, initReq?: fm.InitReq): Promise<GetCoinProductInfosResponse> {
    return fm.fetchReq<GetCoinProductInfosRequest, GetCoinProductInfosResponse>(`/v1/get/coin/productinfos`, {...initReq, method: "POST", body: JSON.stringify(req, fm.replacer)})
  }
  static GetCoinProductInfoOnly(req: GetCoinProductInfoOnlyRequest, initReq?: fm.InitReq): Promise<GetCoinProductInfoOnlyResponse> {
    return fm.fetchReq<GetCoinProductInfoOnlyRequest, GetCoinProductInfoOnlyResponse>(`/v1/get/coin/productinfo/only`, {...initReq, method: "POST", body: JSON.stringify(req, fm.replacer)})
  }
  static GetAppCoinProductInfos(req: GetAppCoinProductInfosRequest, initReq?: fm.InitReq): Promise<GetAppCoinProductInfosResponse> {
    return fm.fetchReq<GetAppCoinProductInfosRequest, GetAppCoinProductInfosResponse>(`/v1/get/app/coin/productinfos`, {...initReq, method: "POST", body: JSON.stringify(req, fm.replacer)})
  }
  static GetAppCoinProductInfoOnly(req: GetAppCoinProductInfoOnlyRequest, initReq?: fm.InitReq): Promise<GetAppCoinProductInfoOnlyResponse> {
    return fm.fetchReq<GetAppCoinProductInfoOnlyRequest, GetAppCoinProductInfoOnlyResponse>(`/v1/get/app/coin/productinfo/only`, {...initReq, method: "POST", body: JSON.stringify(req, fm.replacer)})
  }
  static DeleteCoinProductInfo(req: DeleteCoinProductInfoRequest, initReq?: fm.InitReq): Promise<DeleteCoinProductInfoResponse> {
    return fm.fetchReq<DeleteCoinProductInfoRequest, DeleteCoinProductInfoResponse>(`/v1/delete/coin/productinfo`, {...initReq, method: "POST", body: JSON.stringify(req, fm.replacer)})
  }
}