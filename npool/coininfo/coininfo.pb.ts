/* eslint-disable */
// @ts-nocheck
/*
* This file is a generated Typescript file for GRPC Gateway, DO NOT MODIFY
*/

import * as fm from "../../fetch.pb"
import * as GoogleProtobufEmpty from "../../google/protobuf/empty.pb"
export type VersionResponse = {
  info?: string
}

export type CoinInfo = {
  id?: string
  preSale?: boolean
  name?: string
  unit?: string
  logo?: string
  eNV?: string
  reservedAmount?: number
  createdAt?: number
  updatedAt?: number
  forPay?: boolean
}

export type GetCoinInfoRequest = {
  id?: string
  name?: string
}

export type GetCoinInfoResponse = {
  info?: CoinInfo
}

export type CreateCoinInfoRequest = {
  preSale?: boolean
  name?: string
  unit?: string
  logo?: string
  eNV?: string
}

export type CreateCoinInfoResponse = {
  info?: CoinInfo
}

export type GetCoinInfosRequest = {
  preSale?: boolean
  name?: string
  offset?: number
  limit?: number
}

export type GetCoinInfosResponse = {
  total?: number
  infos?: CoinInfo[]
}

export type UpdateCoinInfoRequest = {
  id?: string
  preSale?: boolean
  logo?: string
  reservedAmount?: number
}

export type UpdateCoinInfoResponse = {
  info?: CoinInfo
}

export class SphinxCoinInfo {
  static Version(req: GoogleProtobufEmpty.Empty, initReq?: fm.InitReq): Promise<VersionResponse> {
    return fm.fetchReq<GoogleProtobufEmpty.Empty, VersionResponse>(`/version`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static CreateCoinInfo(req: CreateCoinInfoRequest, initReq?: fm.InitReq): Promise<CreateCoinInfoResponse> {
    return fm.fetchReq<CreateCoinInfoRequest, CreateCoinInfoResponse>(`/v1/create/coininfo`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetCoinInfo(req: GetCoinInfoRequest, initReq?: fm.InitReq): Promise<GetCoinInfoResponse> {
    return fm.fetchReq<GetCoinInfoRequest, GetCoinInfoResponse>(`/v1/get/coininfo`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetCoinInfos(req: GetCoinInfosRequest, initReq?: fm.InitReq): Promise<GetCoinInfosResponse> {
    return fm.fetchReq<GetCoinInfosRequest, GetCoinInfosResponse>(`/v1/get/coininfos`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static UpdateCoinInfo(req: UpdateCoinInfoRequest, initReq?: fm.InitReq): Promise<UpdateCoinInfoResponse> {
    return fm.fetchReq<UpdateCoinInfoRequest, UpdateCoinInfoResponse>(`/v1/update/coininfo`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
}