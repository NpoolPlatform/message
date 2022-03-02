/* eslint-disable */
// @ts-nocheck
/*
* This file is a generated Typescript file for GRPC Gateway, DO NOT MODIFY
*/

import * as fm from "../../fetch.pb"
import * as GoogleProtobufEmpty from "../../google/protobuf/empty.pb"
import * as SphinxPluginV1Sphinxplugin from "../sphinxplugin/sphinxplugin.pb"

export enum TransactionType {
  Invalid = "Invalid",
  WalletNew = "WalletNew",
  TransactionNew = "TransactionNew",
  Signature = "Signature",
  Balance = "Balance",
  PreSign = "PreSign",
  Broadcast = "Broadcast",
  RegisterCoin = "RegisterCoin",
  SyncMsgState = "SyncMsgState",
}

export enum TransactionState {
  TransactionStateUnKnow = "TransactionStateUnKnow",
  TransactionStateWait = "TransactionStateWait",
  TransactionStateSign = "TransactionStateSign",
  TransactionStateSync = "TransactionStateSync",
  TransactionStateDone = "TransactionStateDone",
  TransactionStateFail = "TransactionStateFail",
}

export type VersionResponse = {
  info?: string
}

export type GetBalanceRequest = {
  name?: string
  address?: string
}

export type BalanceInfo = {
  balance?: number
  balanceStr?: string
}

export type GetBalanceResponse = {
  info?: BalanceInfo
}

export type CreateWalletRequest = {
  name?: string
}

export type WalletInfo = {
  address?: string
}

export type CreateWalletResponse = {
  info?: WalletInfo
}

export type CreateTransactionRequest = {
  name?: string
  transactionID?: string
  amount?: number
  from?: string
  to?: string
}

export type CreateTransactionResponse = {
}

export type GetTransactionRequest = {
  transactionID?: string
}

export type TransactionInfo = {
  transactionID?: string
  name?: string
  amount?: number
  from?: string
  to?: string
  transactionState?: TransactionState
  cID?: string
  exitCode?: string
  createdAt?: number
  updatedAt?: number
}

export type GetTransactionResponse = {
  info?: TransactionInfo
}

export type ProxyPluginResponse = {
  coinType?: SphinxPluginV1Sphinxplugin.CoinType
  transactionType?: TransactionType
  eNV?: string
  unit?: string
  transactionID?: string
  cID?: string
  balance?: number
  balanceStr?: string
  message?: SphinxPluginV1Sphinxplugin.UnsignedMessage
  exitCode?: string
}

export type ProxyPluginRequest = {
  coinType?: SphinxPluginV1Sphinxplugin.CoinType
  transactionType?: TransactionType
  transactionID?: string
  address?: string
  message?: SphinxPluginV1Sphinxplugin.UnsignedMessage
  signature?: SphinxPluginV1Sphinxplugin.Signature
  msgTx?: SphinxPluginV1Sphinxplugin.MsgTx
  cID?: string
}

export type ProxySignRequest = {
  netENV?: string
  coinType?: SphinxPluginV1Sphinxplugin.CoinType
  transactionType?: TransactionType
  transactionID?: string
  message?: SphinxPluginV1Sphinxplugin.UnsignedMessage
}

export type ProxySignResponse = {
  coinType?: SphinxPluginV1Sphinxplugin.CoinType
  transactionType?: TransactionType
  transactionID?: string
  info?: ProxySignResponseInfo
  msgTx?: SphinxPluginV1Sphinxplugin.MsgTx
}

export type ProxySignResponseInfo = {
  address?: string
  message?: SphinxPluginV1Sphinxplugin.UnsignedMessage
  signature?: SphinxPluginV1Sphinxplugin.Signature
}

export class SphinxProxy {
  static Version(req: GoogleProtobufEmpty.Empty, initReq?: fm.InitReq): Promise<VersionResponse> {
    return fm.fetchReq<GoogleProtobufEmpty.Empty, VersionResponse>(`/version`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetBalance(req: GetBalanceRequest, initReq?: fm.InitReq): Promise<GetBalanceResponse> {
    return fm.fetchReq<GetBalanceRequest, GetBalanceResponse>(`/v1/get/balance`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static CreateWallet(req: CreateWalletRequest, initReq?: fm.InitReq): Promise<CreateWalletResponse> {
    return fm.fetchReq<CreateWalletRequest, CreateWalletResponse>(`/v1/create/wallet`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static CreateTransaction(req: CreateTransactionRequest, initReq?: fm.InitReq): Promise<CreateTransactionResponse> {
    return fm.fetchReq<CreateTransactionRequest, CreateTransactionResponse>(`/v1/create/transaction`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetTransaction(req: GetTransactionRequest, initReq?: fm.InitReq): Promise<GetTransactionResponse> {
    return fm.fetchReq<GetTransactionRequest, GetTransactionResponse>(`/v1/get/transaction`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
}