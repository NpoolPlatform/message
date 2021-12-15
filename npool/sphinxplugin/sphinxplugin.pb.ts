/* eslint-disable */
// @ts-nocheck
/*
* This file is a generated Typescript file for GRPC Gateway, DO NOT MODIFY
*/

export enum CoinType {
  CoinTypeUnKnow = "CoinTypeUnKnow",
  CoinTypeFIL = "CoinTypeFIL",
  CoinTypeBTC = "CoinTypeBTC",
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