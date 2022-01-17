/* eslint-disable */
// @ts-nocheck
/*
* This file is a generated Typescript file for GRPC Gateway, DO NOT MODIFY
*/

export enum CoinType {
  CoinTypeUnKnow = "CoinTypeUnKnow",
  CoinTypeFIL = "CoinTypeFIL",
  CoinTypeBTC = "CoinTypeBTC",
  CoinTypeETH = "CoinTypeETH",
  CoinTypeSpaceMesh = "CoinTypeSpaceMesh",
}

export type UnsignedMessage = {
  Version?: string
  To?: string
  From?: string
  Nonce?: string
  Value?: number
  GasLimit?: string
  GasFeeCap?: string
  GasPremium?: string
  Method?: string
  Params?: Uint8Array
}

export type Signature = {
  SignType?: string
  Data?: Uint8Array
}

export type MsgTx = {
  Version?: number
  TxIn?: TxIn[]
  TxOut?: TxOut[]
  LockTime?: number
}

export type TxIn = {
  PreviousOutPoint?: OutPoint
  SignatureScript?: Uint8Array
  Witness?: Uint8Array[]
  Sequence?: number
}

export type OutPoint = {
  Hash?: Uint8Array
  Index?: number
}

export type TxOut = {
  Value?: string
  PkScript?: Uint8Array
}