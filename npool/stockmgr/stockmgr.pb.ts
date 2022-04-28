/* eslint-disable */
// @ts-nocheck
/*
* This file is a generated Typescript file for GRPC Gateway, DO NOT MODIFY
*/

import * as fm from "../../fetch.pb"
import * as GoogleProtobufAny from "../../google/protobuf/any.pb"
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

export type CreateStocksRequest = {
  infos?: Stock[]
}

export type CreateStocksResponse = {
  infos?: Stock[]
}

export type UpdateStockRequest = {
  info?: Stock
}

export type UpdateStockResponse = {
  info?: Stock
}

export type UpdateStockFieldsRequest = {
  id?: string
  fields?: {[key: string]: GoogleProtobufAny.Any}
}

export type UpdateStockFieldsResponse = {
  info?: Stock
}

export type AtomicIncStockRequest = {
  id?: string
  fields?: string[]
}

export type AtomicIncStockResponse = {
  info?: Stock
}

export type AtomicSubStockRequest = {
  id?: string
  fields?: string[]
}

export type AtomicSubStockResponse = {
  info?: Stock
}

export type AtomicSetStockRequest = {
  id?: string
  fields?: {[key: string]: GoogleProtobufAny.Any}
}

export type AtomicSetStockResponse = {
  info?: Stock
}

export type ExistStockRequest = {
  fields?: {[key: string]: GoogleProtobufAny.Any}
}

export type ExistStockResponse = {
  result?: boolean
}

export type CountStocksRequest = {
  fields?: {[key: string]: GoogleProtobufAny.Any}
}

export type CountStocksResponse = {
  result?: number
}

export type DeleteStockRequest = {
  id?: string
}

export type DeleteStockResponse = {
  info?: Stock
}

export class StockManager {
  static Version(req: GoogleProtobufEmpty.Empty, initReq?: fm.InitReq): Promise<NpoolV1Npool.VersionResponse> {
    return fm.fetchReq<GoogleProtobufEmpty.Empty, NpoolV1Npool.VersionResponse>(`/version`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static CreateStock(req: CreateStockRequest, initReq?: fm.InitReq): Promise<CreateStockResponse> {
    return fm.fetchReq<CreateStockRequest, CreateStockResponse>(`/v1/create/stock`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static CreateStocks(req: CreateStocksRequest, initReq?: fm.InitReq): Promise<CreateStocksResponse> {
    return fm.fetchReq<CreateStocksRequest, CreateStocksResponse>(`/v1/create/stocks`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static UpdateStock(req: UpdateStockRequest, initReq?: fm.InitReq): Promise<UpdateStockResponse> {
    return fm.fetchReq<UpdateStockRequest, UpdateStockResponse>(`/v1/update/stock`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static UpdateStockFields(req: UpdateStockFieldsRequest, initReq?: fm.InitReq): Promise<UpdateStockFieldsResponse> {
    return fm.fetchReq<UpdateStockFieldsRequest, UpdateStockFieldsResponse>(`/v1/update/stock/fields`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static AtomicIncStock(req: AtomicIncStockRequest, initReq?: fm.InitReq): Promise<AtomicIncStockResponse> {
    return fm.fetchReq<AtomicIncStockRequest, AtomicIncStockResponse>(`/v1/atomic/inc/stock/fields`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static AtomicSubStock(req: AtomicSubStockRequest, initReq?: fm.InitReq): Promise<AtomicSubStockResponse> {
    return fm.fetchReq<AtomicSubStockRequest, AtomicSubStockResponse>(`/v1/atomic/sub/stock/fields`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static AtomicSetStock(req: AtomicSetStockRequest, initReq?: fm.InitReq): Promise<AtomicSetStockResponse> {
    return fm.fetchReq<AtomicSetStockRequest, AtomicSetStockResponse>(`/v1/atomic/set/stock/fields`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static ExistStock(req: ExistStockRequest, initReq?: fm.InitReq): Promise<ExistStockResponse> {
    return fm.fetchReq<ExistStockRequest, ExistStockResponse>(`/v1/exist/stock`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static CountStocks(req: CountStocksRequest, initReq?: fm.InitReq): Promise<CountStocksResponse> {
    return fm.fetchReq<CountStocksRequest, CountStocksResponse>(`/v1/count/stocks`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static DeleteStock(req: DeleteStockRequest, initReq?: fm.InitReq): Promise<DeleteStockResponse> {
    return fm.fetchReq<DeleteStockRequest, DeleteStockResponse>(`/v1/delete/stock`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
}