/* eslint-disable */
// @ts-nocheck
/*
* This file is a generated Typescript file for GRPC Gateway, DO NOT MODIFY
*/

export enum CoinType {
  CoinTypeUnKnow = "CoinTypeUnKnow",
  CoinTypefilecoin = "CoinTypefilecoin",
  CoinTypebitcoin = "CoinTypebitcoin",
  CoinTypeethereum = "CoinTypeethereum",
  CoinTypeusdterc20 = "CoinTypeusdterc20",
  CoinTypespacemesh = "CoinTypespacemesh",
  CoinTypesolana = "CoinTypesolana",
  CoinTypeusdttrc20 = "CoinTypeusdttrc20",
  CoinTypebinancecoin = "CoinTypebinancecoin",
  CoinTypetron = "CoinTypetron",
  CoinTypebinanceusd = "CoinTypebinanceusd",
  CoinTypetfilecoin = "CoinTypetfilecoin",
  CoinTypetbitcoin = "CoinTypetbitcoin",
  CoinTypetethereum = "CoinTypetethereum",
  CoinTypetusdterc20 = "CoinTypetusdterc20",
  CoinTypetspacemesh = "CoinTypetspacemesh",
  CoinTypetsolana = "CoinTypetsolana",
  CoinTypetusdttrc20 = "CoinTypetusdttrc20",
  CoinTypetbinancecoin = "CoinTypetbinancecoin",
  CoinTypettron = "CoinTypettron",
  CoinTypetbinanceusd = "CoinTypetbinanceusd",
}

export enum TransactionType {
  Invalid = "Invalid",
  WalletNew = "WalletNew",
  TransactionNew = "TransactionNew",
  Sign = "Sign",
  Balance = "Balance",
  PreSign = "PreSign",
  Broadcast = "Broadcast",
  RegisterCoin = "RegisterCoin",
  SyncMsgState = "SyncMsgState",
}

export type UnsignedMessage = {
  version?: string
  to?: string
  from?: string
  value?: number
  nonce?: string
  gasLimit?: string
  gasPrice?: string
  chainID?: string
  contractID?: string
  gasFeeCap?: string
  gasPremium?: string
  method?: string
  params?: Uint8Array
  unspent?: Unspent[]
  recentBhash?: string
  txData?: Uint8Array
}

export type Signature = {
  signType?: string
  data?: Uint8Array
}

export type Unspent = {
  txID?: string
  vout?: number
  address?: string
  account?: string
  scriptPubKey?: string
  redeemScript?: string
  amount?: number
  confirmations?: string
  spendable?: boolean
}

export type MsgTx = {
  version?: number
  txIn?: TxIn[]
  txOut?: TxOut[]
  lockTime?: number
}

export type TxIn = {
  previousOutPoint?: OutPoint
  signatureScript?: Uint8Array
  witness?: Uint8Array[]
  sequence?: number
}

export type OutPoint = {
  hash?: Uint8Array
  index?: number
}

export type TxOut = {
  value?: string
  pkScript?: Uint8Array
}