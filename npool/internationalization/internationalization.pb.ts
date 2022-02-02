/* eslint-disable */
// @ts-nocheck
/*
* This file is a generated Typescript file for GRPC Gateway, DO NOT MODIFY
*/

import * as fm from "../../fetch.pb"
import * as GoogleProtobufEmpty from "../../google/protobuf/empty.pb"
import * as NpoolV1Npool from "../npool.pb"
export type Lang = {
  id?: string
  lang?: string
  logo?: string
  name?: string
}

export type AddLangRequest = {
  info?: Lang
}

export type AddLangResponse = {
  info?: Lang
}

export type GetLangRequest = {
  id?: string
}

export type GetLangResponse = {
  info?: Lang
}

export type UpdateLangRequest = {
  info?: Lang
}

export type UpdateLangResponse = {
  info?: Lang
}

export type GetLangsRequest = {
}

export type GetLangsResponse = {
  infos?: Lang[]
}

export type Message = {
  id?: string
  appID?: string
  messageID?: string
  langID?: string
  message?: string
  batchGet?: boolean
}

export type CreateMessageRequest = {
  info?: Message
}

export type CreateMessageResponse = {
  info?: Message
}

export type CreateMessagesRequest = {
  infos?: Message[]
}

export type CreateMessagesResponse = {
  infos?: Message[]
}

export type UpdateMessageRequest = {
  info?: Message
}

export type UpdateMessageResponse = {
  info?: Message
}

export type UpdateMessagesRequest = {
  infos?: Message[]
}

export type UpdateMessagesResponse = {
  infos?: Message[]
}

export type GetMessagesByLangIDRequest = {
  langID?: string
}

export type GetMessagesByLangIDResponse = {
  infos?: Message[]
}

export type GetMessageByLangIDMessageIDRequest = {
  langID?: string
  messageID?: string
}

export type GetMessageByLangIDMessageIDResponse = {
  info?: Message
}

export type AppLang = {
  id?: string
  appID?: string
  langID?: string
}

export type CreateAppLangRequest = {
  info?: AppLang
}

export type CreateAppLangResponse = {
  info?: AppLang
}

export type GetAppLangRequest = {
  id?: string
}

export type GetAppLangResponse = {
  info?: AppLang
}

export type GetAppLangsByAppRequest = {
  appID?: string
}

export type GetAppLangsByAppResponse = {
  infos?: AppLang[]
}

export type AppLangInfo = {
  appLang?: AppLang
  lang?: Lang
}

export type GetAppLangInfosByAppRequest = {
  appID?: string
  targetAppID?: string
}

export type GetAppLangInfosByAppResponse = {
  infos?: AppLangInfo[]
}

export class Internationalization {
  static Version(req: GoogleProtobufEmpty.Empty, initReq?: fm.InitReq): Promise<NpoolV1Npool.VersionResponse> {
    return fm.fetchReq<GoogleProtobufEmpty.Empty, NpoolV1Npool.VersionResponse>(`/version`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static AddLang(req: AddLangRequest, initReq?: fm.InitReq): Promise<AddLangResponse> {
    return fm.fetchReq<AddLangRequest, AddLangResponse>(`/v1/add/lang`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetLang(req: GetLangRequest, initReq?: fm.InitReq): Promise<GetLangResponse> {
    return fm.fetchReq<GetLangRequest, GetLangResponse>(`/v1/get/lang`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static UpdateLang(req: UpdateLangRequest, initReq?: fm.InitReq): Promise<UpdateLangResponse> {
    return fm.fetchReq<UpdateLangRequest, UpdateLangResponse>(`/v1/update/lang`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetLangs(req: GetLangsRequest, initReq?: fm.InitReq): Promise<GetLangsResponse> {
    return fm.fetchReq<GetLangsRequest, GetLangsResponse>(`/v1/get/langs`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static CreateAppLang(req: CreateAppLangRequest, initReq?: fm.InitReq): Promise<CreateAppLangResponse> {
    return fm.fetchReq<CreateAppLangRequest, CreateAppLangResponse>(`/v1/create/app/lang`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetAppLang(req: GetAppLangRequest, initReq?: fm.InitReq): Promise<GetAppLangResponse> {
    return fm.fetchReq<GetAppLangRequest, GetAppLangResponse>(`/v1/get/app/lang`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetAppLangsByApp(req: GetAppLangsByAppRequest, initReq?: fm.InitReq): Promise<GetAppLangsByAppResponse> {
    return fm.fetchReq<GetAppLangsByAppRequest, GetAppLangsByAppResponse>(`/v1/get/app/langs/by/app`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetAppLangInfosByApp(req: GetAppLangInfosByAppRequest, initReq?: fm.InitReq): Promise<GetAppLangInfosByAppResponse> {
    return fm.fetchReq<GetAppLangInfosByAppRequest, GetAppLangInfosByAppResponse>(`/v1/get/app/lang/infos/by/app`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static CreateMessage(req: CreateMessageRequest, initReq?: fm.InitReq): Promise<CreateMessageResponse> {
    return fm.fetchReq<CreateMessageRequest, CreateMessageResponse>(`/v1/create/message`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static CreateMessages(req: CreateMessagesRequest, initReq?: fm.InitReq): Promise<CreateMessagesResponse> {
    return fm.fetchReq<CreateMessagesRequest, CreateMessagesResponse>(`/v1/create/messages`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static UpdateMessage(req: UpdateMessageRequest, initReq?: fm.InitReq): Promise<UpdateMessageResponse> {
    return fm.fetchReq<UpdateMessageRequest, UpdateMessageResponse>(`/v1/update/message`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static UpdateMessages(req: UpdateMessagesRequest, initReq?: fm.InitReq): Promise<UpdateMessagesResponse> {
    return fm.fetchReq<UpdateMessagesRequest, UpdateMessagesResponse>(`/v1/update/messages`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetMessagesByLangID(req: GetMessagesByLangIDRequest, initReq?: fm.InitReq): Promise<GetMessagesByLangIDResponse> {
    return fm.fetchReq<GetMessagesByLangIDRequest, GetMessagesByLangIDResponse>(`/v1/get/messages/by/lang/id`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetMessageByLangIDMessageID(req: GetMessageByLangIDMessageIDRequest, initReq?: fm.InitReq): Promise<GetMessageByLangIDMessageIDResponse> {
    return fm.fetchReq<GetMessageByLangIDMessageIDRequest, GetMessageByLangIDMessageIDResponse>(`/v1/get/message/by/lang/id/message/id`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
}