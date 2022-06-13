/* eslint-disable */
// @ts-nocheck
/*
* This file is a generated Typescript file for GRPC Gateway, DO NOT MODIFY
*/

import * as fm from "../../fetch.pb"
import * as GoogleProtobufEmpty from "../../google/protobuf/empty.pb"
import * as GoogleProtobufStruct from "../../google/protobuf/struct.pb"
import * as NpoolV1Npool from "../npool.pb"
export type Stock = {
  id?: string
  goodID?: string
  total?: number
  inService?: number
  sold?: number
  locked?: number
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
  fields?: {[key: string]: GoogleProtobufStruct.Value}
}

export type UpdateStockFieldsResponse = {
  info?: Stock
}

export type AddStockFieldsRequest = {
  id?: string
  fields?: {[key: string]: GoogleProtobufStruct.Value}
}

export type AddStockFieldsResponse = {
  info?: Stock
}

export type GetStockRequest = {
  id?: string
}

export type GetStockResponse = {
  info?: Stock
}

export type GetStockOnlyRequest = {
  conds?: {[key: string]: NpoolV1Npool.FilterCond}
}

export type GetStockOnlyResponse = {
  info?: Stock
}

export type GetStocksRequest = {
  conds?: {[key: string]: NpoolV1Npool.FilterCond}
  offset?: number
  limit?: number
}

export type GetStocksResponse = {
  infos?: Stock[]
  total?: number
}

export type ExistStockRequest = {
  id?: string
}

export type ExistStockResponse = {
  result?: boolean
}

export type ExistStockCondsRequest = {
  conds?: {[key: string]: NpoolV1Npool.FilterCond}
}

export type ExistStockCondsResponse = {
  result?: boolean
}

export type CountStocksRequest = {
  conds?: {[key: string]: NpoolV1Npool.FilterCond}
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
  static AddStockFields(req: AddStockFieldsRequest, initReq?: fm.InitReq): Promise<AddStockFieldsResponse> {
    return fm.fetchReq<AddStockFieldsRequest, AddStockFieldsResponse>(`/v1/add/stock/fields`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetStock(req: GetStockRequest, initReq?: fm.InitReq): Promise<GetStockResponse> {
    return fm.fetchReq<GetStockRequest, GetStockResponse>(`/v1/get/stock`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetStockOnly(req: GetStockOnlyRequest, initReq?: fm.InitReq): Promise<GetStockOnlyResponse> {
    return fm.fetchReq<GetStockOnlyRequest, GetStockOnlyResponse>(`/v1/get/stock/only`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetStocks(req: GetStocksRequest, initReq?: fm.InitReq): Promise<GetStocksResponse> {
    return fm.fetchReq<GetStocksRequest, GetStocksResponse>(`/v1/get/stocks`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static ExistStock(req: ExistStockRequest, initReq?: fm.InitReq): Promise<ExistStockResponse> {
    return fm.fetchReq<ExistStockRequest, ExistStockResponse>(`/v1/exist/stock`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static ExistStockConds(req: ExistStockCondsRequest, initReq?: fm.InitReq): Promise<ExistStockCondsResponse> {
    return fm.fetchReq<ExistStockCondsRequest, ExistStockCondsResponse>(`/v1/exist/stock/conds`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static CountStocks(req: CountStocksRequest, initReq?: fm.InitReq): Promise<CountStocksResponse> {
    return fm.fetchReq<CountStocksRequest, CountStocksResponse>(`/v1/count/stocks`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static DeleteStock(req: DeleteStockRequest, initReq?: fm.InitReq): Promise<DeleteStockResponse> {
    return fm.fetchReq<DeleteStockRequest, DeleteStockResponse>(`/v1/delete/stock`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
}