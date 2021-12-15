/* eslint-disable */
// @ts-nocheck
/*
* This file is a generated Typescript file for GRPC Gateway, DO NOT MODIFY
*/

import * as fm from "../../fetch.pb"
export type CoinInfo = {
  ID?: string
  PreSale?: boolean
  Name?: string
  Unit?: string
  Logo?: string
  ReservedAmount?: number
  CreatedAt?: number
  UpdatedAt?: number
}

export type GetCoinInfoRequest = {
  ID?: string
  Name?: string
}

export type GetCoinInfoResponse = {
  Info?: CoinInfo
}

export type CreateCoinInfoRequest = {
  PreSale?: boolean
  Name?: string
  Unit?: string
  Logo?: string
}

export type CreateCoinInfoResponse = {
  Info?: CoinInfo
}

export type GetCoinInfosRequest = {
  PreSale?: boolean
  Name?: string
  Offset?: number
  Limit?: number
}

export type GetCoinInfosResponse = {
  Total?: number
  Infos?: CoinInfo[]
}

export type UpdateCoinInfoRequest = {
  ID?: string
  PreSale?: boolean
  Logo?: string
  ReservedAmount?: number
}

export type UpdateCoinInfoResponse = {
  Info?: CoinInfo
}

export class SphinxCoinInfo {
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