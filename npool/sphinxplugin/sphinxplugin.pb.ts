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
  CoinTypetfilecoin = "CoinTypetfilecoin",
  CoinTypetbitcoin = "CoinTypetbitcoin",
  CoinTypetethereum = "CoinTypetethereum",
  CoinTypetusdterc20 = "CoinTypetusdterc20",
  CoinTypetspacemesh = "CoinTypetspacemesh",
  CoinTypetsolana = "CoinTypetsolana",
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