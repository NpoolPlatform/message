/* eslint-disable */
// @ts-nocheck
/*
* This file is a generated Typescript file for GRPC Gateway, DO NOT MODIFY
*/

import * as fm from "../../fetch.pb"
import * as GoogleProtobufEmpty from "../../google/protobuf/empty.pb"
import * as NpoolV1Npool from "../npool.pb"
export type EmailSubscriber = {
  id?: string
  appID?: string
  emailAddress?: string
}

export type CreateEmailSubscriberRequest = {
  info?: EmailSubscriber
}

export type CreateEmailSubscriberResponse = {
  info?: EmailSubscriber
}

export type CreateEmailSubscribersRequest = {
  infos?: EmailSubscriber[]
}

export type CreateEmailSubscribersResponse = {
  infos?: EmailSubscriber[]
}

export type GetEmailSubscribersRequest = {
  appID?: string
  conds?: {[key: string]: NpoolV1Npool.FilterCond}
  offset?: number
  limit?: number
}

export type GetEmailSubscribersResponse = {
  infos?: EmailSubscriber[]
}

export type GetAppEmailSubscribersRequest = {
  targetAppID?: string
  conds?: {[key: string]: NpoolV1Npool.FilterCond}
  offset?: number
  limit?: number
}

export type GetAppEmailSubscribersResponse = {
  infos?: EmailSubscriber[]
}

export class SubscribeManager {
  static Version(req: GoogleProtobufEmpty.Empty, initReq?: fm.InitReq): Promise<NpoolV1Npool.VersionResponse> {
    return fm.fetchReq<GoogleProtobufEmpty.Empty, NpoolV1Npool.VersionResponse>(`/version`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static CreateEmailSubscriber(req: CreateEmailSubscriberRequest, initReq?: fm.InitReq): Promise<CreateEmailSubscriberResponse> {
    return fm.fetchReq<CreateEmailSubscriberRequest, CreateEmailSubscriberResponse>(`/v1/create/email/subscriber`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static CreateEmailSubscribers(req: CreateEmailSubscribersRequest, initReq?: fm.InitReq): Promise<CreateEmailSubscribersResponse> {
    return fm.fetchReq<CreateEmailSubscribersRequest, CreateEmailSubscribersResponse>(`/v1/create/email/subscribers`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetEmailSubscribers(req: GetEmailSubscribersRequest, initReq?: fm.InitReq): Promise<GetEmailSubscribersResponse> {
    return fm.fetchReq<GetEmailSubscribersRequest, GetEmailSubscribersResponse>(`/v1/get/email/subscribers`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetAppEmailSubscribers(req: GetAppEmailSubscribersRequest, initReq?: fm.InitReq): Promise<GetAppEmailSubscribersResponse> {
    return fm.fetchReq<GetAppEmailSubscribersRequest, GetAppEmailSubscribersResponse>(`/v1/get/app/email/subscribers`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
}