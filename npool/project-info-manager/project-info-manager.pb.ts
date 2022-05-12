/* eslint-disable */
// @ts-nocheck
/*
* This file is a generated Typescript file for GRPC Gateway, DO NOT MODIFY
*/

import * as fm from "../../fetch.pb"
import * as GoogleProtobufEmpty from "../../google/protobuf/empty.pb"
import * as NpoolV1Npool from "../npool.pb"
export type CoinDescriptionBase = {
  appID?: string
  coinTypeID?: string
  title?: string
  message?: string
  usedFor?: string
}

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
  info?: CoinDescriptionBase
}

export type CreateCoinDescriptionResponse = {
  info?: CoinDescription
}

export type CreateCoinDescriptionsRequest = {
  infos?: CoinDescriptionBase[]
}

export type CreateCoinDescriptionsResponse = {
  infos?: CoinDescription[]
}

export type UpdateCoinDescriptionRequest = {
  appID?: string
  info?: CoinDescription
}

export type UpdateCoinDescriptionResponse = {
  info?: CoinDescription
}

export type UpdateAppCoinDescriptionRequest = {
  targetAppID?: string
  info?: CoinDescription
}

export type UpdateAppCoinDescriptionResponse = {
  info?: CoinDescription
}

export type GetCoinDescriptionRequest = {
  appID?: string
  id?: string
}

export type GetCoinDescriptionResponse = {
  info?: CoinDescription
}

export type GetAppCoinDescriptionRequest = {
  targetAppID?: string
  id?: string
}

export type GetAppCoinDescriptionResponse = {
  info?: CoinDescription
}

export type GetCoinDescriptionsRequest = {
  appID?: string
  offset?: number
  limit?: number
}

export type GetCoinDescriptionsResponse = {
  infos?: CoinDescription[]
  total?: number
}

export type GetAppCoinDescriptionsRequest = {
  targetAppID?: string
  offset?: number
  limit?: number
}

export type GetAppCoinDescriptionsResponse = {
  infos?: CoinDescription[]
  total?: number
}

export type CountCoinDescriptionsRequest = {
  appID?: string
}

export type CountCoinDescriptionsResponse = {
  result?: number
}

export type CountAppCoinDescriptionsRequest = {
  targetAppID?: string
}

export type CountAppCoinDescriptionsResponse = {
  result?: number
}

export type DeleteAppCoinDescriptionRequest = {
  targetAppID?: string
  id?: string
}

export type DeleteAppCoinDescriptionResponse = {
  info?: CoinDescription
}

export class ProjectInfoManager {
  static Version(req: GoogleProtobufEmpty.Empty, initReq?: fm.InitReq): Promise<NpoolV1Npool.VersionResponse> {
    return fm.fetchReq<GoogleProtobufEmpty.Empty, NpoolV1Npool.VersionResponse>(`/version`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static CreateCoinDescription(req: CreateCoinDescriptionRequest, initReq?: fm.InitReq): Promise<CreateCoinDescriptionResponse> {
    return fm.fetchReq<CreateCoinDescriptionRequest, CreateCoinDescriptionResponse>(`/v1/create/coin/description`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static CreateCoinDescriptions(req: CreateCoinDescriptionsRequest, initReq?: fm.InitReq): Promise<CreateCoinDescriptionsResponse> {
    return fm.fetchReq<CreateCoinDescriptionsRequest, CreateCoinDescriptionsResponse>(`/v1/create/coin/descriptions`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static UpdateCoinDescription(req: UpdateCoinDescriptionRequest, initReq?: fm.InitReq): Promise<UpdateCoinDescriptionResponse> {
    return fm.fetchReq<UpdateCoinDescriptionRequest, UpdateCoinDescriptionResponse>(`/v1/update/coin/description`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static UpdateAppCoinDescription(req: UpdateAppCoinDescriptionRequest, initReq?: fm.InitReq): Promise<UpdateAppCoinDescriptionResponse> {
    return fm.fetchReq<UpdateAppCoinDescriptionRequest, UpdateAppCoinDescriptionResponse>(`/v1/update/app/coin/description`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetCoinDescription(req: GetCoinDescriptionRequest, initReq?: fm.InitReq): Promise<GetCoinDescriptionResponse> {
    return fm.fetchReq<GetCoinDescriptionRequest, GetCoinDescriptionResponse>(`/v1/get/coin/description`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetAppCoinDescription(req: GetAppCoinDescriptionRequest, initReq?: fm.InitReq): Promise<GetAppCoinDescriptionResponse> {
    return fm.fetchReq<GetAppCoinDescriptionRequest, GetAppCoinDescriptionResponse>(`/v1/get/app/coin/description`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetCoinDescriptions(req: GetCoinDescriptionsRequest, initReq?: fm.InitReq): Promise<GetCoinDescriptionsResponse> {
    return fm.fetchReq<GetCoinDescriptionsRequest, GetCoinDescriptionsResponse>(`/v1/get/coin/descriptions`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetAppCoinDescriptions(req: GetAppCoinDescriptionsRequest, initReq?: fm.InitReq): Promise<GetAppCoinDescriptionsResponse> {
    return fm.fetchReq<GetAppCoinDescriptionsRequest, GetAppCoinDescriptionsResponse>(`/v1/get/app/coin/descriptions`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static CountCoinDescriptions(req: CountCoinDescriptionsRequest, initReq?: fm.InitReq): Promise<CountCoinDescriptionsResponse> {
    return fm.fetchReq<CountCoinDescriptionsRequest, CountCoinDescriptionsResponse>(`/v1/count/coin/descriptions`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static CountAppCoinDescriptions(req: CountAppCoinDescriptionsRequest, initReq?: fm.InitReq): Promise<CountAppCoinDescriptionsResponse> {
    return fm.fetchReq<CountAppCoinDescriptionsRequest, CountAppCoinDescriptionsResponse>(`/v1/count/app/coin/descriptions`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static DeleteAppCoinDescription(req: DeleteAppCoinDescriptionRequest, initReq?: fm.InitReq): Promise<DeleteAppCoinDescriptionResponse> {
    return fm.fetchReq<DeleteAppCoinDescriptionRequest, DeleteAppCoinDescriptionResponse>(`/v1/delete/app/coin/description`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
}