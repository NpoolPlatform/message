/* eslint-disable */
// @ts-nocheck
/*
* This file is a generated Typescript file for GRPC Gateway, DO NOT MODIFY
*/

import * as fm from "../../../fetch.pb"
import * as GoogleProtobufEmpty from "../../../google/protobuf/empty.pb"
import * as NpoolV1Npool from "../../npool.pb"
export type Stock = {
  id?: string
  goodID?: string
  inService?: number
  sold?: number
}

export type CreateRequest = {
  info?: Stock
}

export type CreateResponse = {
  info?: Stock
}

export class StockManager {
  static Version(req: GoogleProtobufEmpty.Empty, initReq?: fm.InitReq): Promise<NpoolV1Npool.VersionResponse> {
    return fm.fetchReq<GoogleProtobufEmpty.Empty, NpoolV1Npool.VersionResponse>(`/version`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static Create(req: CreateRequest, initReq?: fm.InitReq): Promise<CreateResponse> {
    return fm.fetchReq<CreateRequest, CreateResponse>(`/v1/create/stock`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
}