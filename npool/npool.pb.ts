/* eslint-disable */
// @ts-nocheck
/*
* This file is a generated Typescript file for GRPC Gateway, DO NOT MODIFY
*/

import * as GoogleProtobufStruct from "../google/protobuf/struct.pb"
export type PageInfo = {
  offset?: number
  limit?: number
}

export type VersionResponse = {
  info?: string
}

export type FilterCond = {
  op?: string
  val?: GoogleProtobufStruct.Value
}