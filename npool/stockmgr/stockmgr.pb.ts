/* eslint-disable */
// @ts-nocheck
/*
* This file is a generated Typescript file for GRPC Gateway, DO NOT MODIFY
*/

import * as fm from "../../fetch.pb"
import * as GoogleProtobufEmpty from "../../google/protobuf/empty.pb"
import * as NpoolV1Npool from "../npool.pb"
export type Stock = {
  id?: string
  goodID?: string
  total?: number
  inService?: number
  sold?: number
}

export type CreateStockRequest = {
  info?: Stock
}

export type CreateStockResponse = {
  info?: Stock
}

export class StockManager {
  static Version(req: GoogleProtobufEmpty.Empty, initReq?: fm.InitReq): Promise<NpoolV1Npool.VersionResponse> {
    return fm.fetchReq<GoogleProtobufEmpty.Empty, NpoolV1Npool.VersionResponse>(`/version`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static CreateStock(req: CreateStockRequest, initReq?: fm.InitReq): Promise<CreateStockResponse> {
    return fm.fetchReq<CreateStockRequest, CreateStockResponse>(`/v1/create/stock`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
}