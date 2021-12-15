/* eslint-disable */
// @ts-nocheck
/*
* This file is a generated Typescript file for GRPC Gateway, DO NOT MODIFY
*/

import * as fm from "../../fetch.pb"
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
  TransactionStatePendingReview = "TransactionStatePendingReview",
  TransactionStateRejected = "TransactionStateRejected",
  TransactionStateWait = "TransactionStateWait",
  TransactionStateSign = "TransactionStateSign",
  TransactionStateSync = "TransactionStateSync",
  TransactionStateDone = "TransactionStateDone",
  TransactionStateFail = "TransactionStateFail",
}

export type GetBalanceRequest = {
  Name?: string
  Address?: string
}

export type BalanceInfo = {
  Balance?: number
  BalanceStr?: string
}

export type GetBalanceResponse = {
  Info?: BalanceInfo
}

export type CreateWalletRequest = {
  Name?: string
}

export type WalletInfo = {
  Address?: string
}

export type CreateWalletResponse = {
  Info?: WalletInfo
}

export type CreateTransactionRequest = {
  Name?: string
  TransactionID?: string
  Amount?: number
  From?: string
  To?: string
}

export type CreateTransactionResponse = {
}

export type GetTransactionRequest = {
  TransactionID?: string
}

export type TransactionInfo = {
  TransactionID?: string
  Name?: string
  Amount?: number
  From?: string
  To?: string
  TransactionState?: TransactionState
  CID?: string
  ExitCode?: string
  CreatedAt?: number
  UpdatedAt?: number
}

export type GetTransactionResponse = {
  Info?: TransactionInfo
}

export type ReviewTransactionRequest = {
  TransactionID?: string
  TransactionState?: TransactionState
}

export type ReviewTransactionResponse = {
}

export type ProxyPluginResponse = {
  CoinType?: SphinxPluginV1Sphinxplugin.CoinType
  TransactionType?: TransactionType
  TransactionID?: string
  Nonce?: string
  CID?: string
  Balance?: number
  BalanceStr?: string
  Message?: SphinxPluginV1Sphinxplugin.UnsignedMessage
  ExitCode?: string
}

export type ProxyPluginRequest = {
  CoinType?: SphinxPluginV1Sphinxplugin.CoinType
  TransactionType?: TransactionType
  TransactionID?: string
  Address?: string
  Message?: SphinxPluginV1Sphinxplugin.UnsignedMessage
  Signature?: SphinxPluginV1Sphinxplugin.Signature
  CID?: string
}

export type ProxySignRequest = {
  CoinType?: SphinxPluginV1Sphinxplugin.CoinType
  TransactionType?: TransactionType
  TransactionID?: string
  Message?: SphinxPluginV1Sphinxplugin.UnsignedMessage
}

export type ProxySignResponse = {
  CoinType?: SphinxPluginV1Sphinxplugin.CoinType
  TransactionType?: TransactionType
  TransactionID?: string
  Info?: ProxySignResponseInfo
}

export type ProxySignResponseInfo = {
  Address?: string
  Message?: SphinxPluginV1Sphinxplugin.UnsignedMessage
  Signature?: SphinxPluginV1Sphinxplugin.Signature
}

export class SphinxProxy {
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
  static ReviewTransaction(req: ReviewTransactionRequest, initReq?: fm.InitReq): Promise<ReviewTransactionResponse> {
    return fm.fetchReq<ReviewTransactionRequest, ReviewTransactionResponse>(`/sphinx.proxy.v1.SphinxProxy/ReviewTransaction`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
}