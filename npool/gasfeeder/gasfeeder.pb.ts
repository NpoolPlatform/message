/* eslint-disable */
// @ts-nocheck
/*
* This file is a generated Typescript file for GRPC Gateway, DO NOT MODIFY
*/

import * as fm from "../../fetch.pb"
import * as GoogleProtobufEmpty from "../../google/protobuf/empty.pb"
import * as NpoolV1Npool from "../npool.pb"
export type CoinGas = {
  id?: string
  coinTypeID?: string
  gasCoinTypeID?: string
  depositThreshold?: string
}

export type CreateCoinGasRequest = {
  info?: CoinGas
}

export type CreateCoinGasResponse = {
  info?: CoinGas
}

export type UpdateCoinGasRequest = {
  info?: CoinGas
}

export type UpdateCoinGasResponse = {
  info?: CoinGas
}

export type GetCoinGasRequest = {
  id?: string
}

export type GetCoinGasResponse = {
  info?: CoinGas
}

export type GetCoinGasOnlyRequest = {
  conds?: {[key: string]: NpoolV1Npool.FilterCond}
}

export type GetCoinGasOnlyResponse = {
  info?: CoinGas
}

export type GetCoinGasesRequest = {
  conds?: {[key: string]: NpoolV1Npool.FilterCond}
  offset?: number
  limit?: number
}

export type GetCoinGasesResponse = {
  infos?: CoinGas[]
  total?: number
}

export type ExistCoinGasRequest = {
  id?: string
}

export type ExistCoinGasResponse = {
  result?: boolean
}

export type ExistCoinGasCondsRequest = {
  conds?: {[key: string]: NpoolV1Npool.FilterCond}
}

export type ExistCoinGasCondsResponse = {
  result?: boolean
}

export type CountCoinGasesRequest = {
  conds?: {[key: string]: NpoolV1Npool.FilterCond}
}

export type CountCoinGasesResponse = {
  result?: number
}

export type DeleteCoinGasRequest = {
  id?: string
}

export type DeleteCoinGasResponse = {
  info?: CoinGas
}

export type Deposit = {
  id?: string
  accountID?: string
  depositAmount?: number
  createdAt?: number
}

export type CreateDepositRequest = {
  info?: Deposit
}

export type CreateDepositResponse = {
  info?: Deposit
}

export type UpdateDepositRequest = {
  info?: Deposit
}

export type UpdateDepositResponse = {
  info?: Deposit
}

export type GetDepositRequest = {
  id?: string
}

export type GetDepositResponse = {
  info?: Deposit
}

export type GetDepositOnlyRequest = {
  conds?: {[key: string]: NpoolV1Npool.FilterCond}
}

export type GetDepositOnlyResponse = {
  info?: Deposit
}

export type GetDepositesRequest = {
  conds?: {[key: string]: NpoolV1Npool.FilterCond}
  offset?: number
  limit?: number
}

export type GetDepositesResponse = {
  infos?: Deposit[]
  total?: number
}

export type ExistDepositRequest = {
  id?: string
}

export type ExistDepositResponse = {
  result?: boolean
}

export type ExistDepositCondsRequest = {
  conds?: {[key: string]: NpoolV1Npool.FilterCond}
}

export type ExistDepositCondsResponse = {
  result?: boolean
}

export type CountDepositesRequest = {
  conds?: {[key: string]: NpoolV1Npool.FilterCond}
}

export type CountDepositesResponse = {
  result?: number
}

export type DeleteDepositRequest = {
  id?: string
}

export type DeleteDepositResponse = {
  info?: Deposit
}

export class GasFeeder {
  static Version(req: GoogleProtobufEmpty.Empty, initReq?: fm.InitReq): Promise<NpoolV1Npool.VersionResponse> {
    return fm.fetchReq<GoogleProtobufEmpty.Empty, NpoolV1Npool.VersionResponse>(`/version`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static CreateCoinGas(req: CreateCoinGasRequest, initReq?: fm.InitReq): Promise<CreateCoinGasResponse> {
    return fm.fetchReq<CreateCoinGasRequest, CreateCoinGasResponse>(`/v1/create/coingas`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static UpdateCoinGas(req: UpdateCoinGasRequest, initReq?: fm.InitReq): Promise<UpdateCoinGasResponse> {
    return fm.fetchReq<UpdateCoinGasRequest, UpdateCoinGasResponse>(`/v1/update/coingas`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetCoinGas(req: GetCoinGasRequest, initReq?: fm.InitReq): Promise<GetCoinGasResponse> {
    return fm.fetchReq<GetCoinGasRequest, GetCoinGasResponse>(`/v1/get/coingas`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetCoinGasOnly(req: GetCoinGasOnlyRequest, initReq?: fm.InitReq): Promise<GetCoinGasOnlyResponse> {
    return fm.fetchReq<GetCoinGasOnlyRequest, GetCoinGasOnlyResponse>(`/v1/get/coingas/only`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetCoinGases(req: GetCoinGasesRequest, initReq?: fm.InitReq): Promise<GetCoinGasesResponse> {
    return fm.fetchReq<GetCoinGasesRequest, GetCoinGasesResponse>(`/v1/get/coingass`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static ExistCoinGas(req: ExistCoinGasRequest, initReq?: fm.InitReq): Promise<ExistCoinGasResponse> {
    return fm.fetchReq<ExistCoinGasRequest, ExistCoinGasResponse>(`/v1/exist/coingas`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static ExistCoinGasConds(req: ExistCoinGasCondsRequest, initReq?: fm.InitReq): Promise<ExistCoinGasCondsResponse> {
    return fm.fetchReq<ExistCoinGasCondsRequest, ExistCoinGasCondsResponse>(`/v1/exist/coingas/conds`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static DeleteCoinGas(req: DeleteCoinGasRequest, initReq?: fm.InitReq): Promise<DeleteCoinGasResponse> {
    return fm.fetchReq<DeleteCoinGasRequest, DeleteCoinGasResponse>(`/v1/delete/coingas`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static CreateDeposit(req: CreateDepositRequest, initReq?: fm.InitReq): Promise<CreateDepositResponse> {
    return fm.fetchReq<CreateDepositRequest, CreateDepositResponse>(`/v1/create/deposit`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static UpdateDeposit(req: UpdateDepositRequest, initReq?: fm.InitReq): Promise<UpdateDepositResponse> {
    return fm.fetchReq<UpdateDepositRequest, UpdateDepositResponse>(`/v1/update/deposit`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetDeposit(req: GetDepositRequest, initReq?: fm.InitReq): Promise<GetDepositResponse> {
    return fm.fetchReq<GetDepositRequest, GetDepositResponse>(`/v1/get/deposit`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetDepositOnly(req: GetDepositOnlyRequest, initReq?: fm.InitReq): Promise<GetDepositOnlyResponse> {
    return fm.fetchReq<GetDepositOnlyRequest, GetDepositOnlyResponse>(`/v1/get/deposit/only`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetDeposites(req: GetDepositesRequest, initReq?: fm.InitReq): Promise<GetDepositesResponse> {
    return fm.fetchReq<GetDepositesRequest, GetDepositesResponse>(`/v1/get/deposits`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static ExistDeposit(req: ExistDepositRequest, initReq?: fm.InitReq): Promise<ExistDepositResponse> {
    return fm.fetchReq<ExistDepositRequest, ExistDepositResponse>(`/v1/exist/deposit`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static ExistDepositConds(req: ExistDepositCondsRequest, initReq?: fm.InitReq): Promise<ExistDepositCondsResponse> {
    return fm.fetchReq<ExistDepositCondsRequest, ExistDepositCondsResponse>(`/v1/exist/deposit/conds`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static DeleteDeposit(req: DeleteDepositRequest, initReq?: fm.InitReq): Promise<DeleteDepositResponse> {
    return fm.fetchReq<DeleteDepositRequest, DeleteDepositResponse>(`/v1/delete/deposit`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
}