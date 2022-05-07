/* eslint-disable */
// @ts-nocheck
/*
* This file is a generated Typescript file for GRPC Gateway, DO NOT MODIFY
*/

import * as fm from "../../fetch.pb"
import * as GoogleProtobufEmpty from "../../google/protobuf/empty.pb"
import * as NpoolV1Npool from "../npool.pb"
export type CoinDescriptionInfo = {
  id?: string
  coinTypeID?: string
  title?: string
  message?: string
  usedFor?: string
  createdAt?: number
  updatedAt?: number
}

export type CreateCoinDescriptionRequest = {
  coinTypeID?: string
  title?: string
  message?: string
  usedFor?: string
}

export type CreateCoinDescriptionResponse = {
  info?: CoinDescriptionInfo
}

export type GetCoinDescriptionRequest = {
  coinTypeID?: string
  limit?: number
  offset?: number
}

export type GetCoinDescriptionResponse = {
  total?: number
  infos?: CoinDescriptionInfo[]
}

export type UpdateCoinDescriptionRequest = {
  id?: string
  coinTypeID?: string
  title?: string
  message?: string
  usedFor?: string
}

export type UpdateCoinDescriptionResponse = {
  info?: CoinDescriptionInfo
}

export class ProjectInfoManager {
  static Version(req: GoogleProtobufEmpty.Empty, initReq?: fm.InitReq): Promise<NpoolV1Npool.VersionResponse> {
    return fm.fetchReq<GoogleProtobufEmpty.Empty, NpoolV1Npool.VersionResponse>(`/version`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static CreateCoinDescription(req: CreateCoinDescriptionRequest, initReq?: fm.InitReq): Promise<CreateCoinDescriptionResponse> {
    return fm.fetchReq<CreateCoinDescriptionRequest, CreateCoinDescriptionResponse>(`/v1/create/coin/description`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetCoinDescription(req: GetCoinDescriptionRequest, initReq?: fm.InitReq): Promise<GetCoinDescriptionResponse> {
    return fm.fetchReq<GetCoinDescriptionRequest, GetCoinDescriptionResponse>(`/v1/get/coin/description`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static UpdateCoinDescription(req: UpdateCoinDescriptionRequest, initReq?: fm.InitReq): Promise<UpdateCoinDescriptionResponse> {
    return fm.fetchReq<UpdateCoinDescriptionRequest, UpdateCoinDescriptionResponse>(`/v1/update/coin/description`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
}