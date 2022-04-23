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
  short?: string
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
  targetLangID?: string
  info?: Message
}

export type CreateMessageResponse = {
  info?: Message
}

export type CreateMessageForOtherAppRequest = {
  targetAppID?: string
  targetLangID?: string
  info?: Message
}

export type CreateMessageForOtherAppResponse = {
  info?: Message
}

export type CreateMessagesRequest = {
  targetLangID?: string
  infos?: Message[]
}

export type CreateMessagesResponse = {
  infos?: Message[]
}

export type CreateMessagesForOtherAppRequest = {
  targetAppID?: string
  targetLangID?: string
  infos?: Message[]
}

export type CreateMessagesForOtherAppResponse = {
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

export type GetMessagesByAppLangRequest = {
  appID?: string
  langID?: string
  targetLangID?: string
}

export type GetMessagesByAppLangResponse = {
  infos?: Message[]
}

export type GetMessagesByOtherAppLangRequest = {
  targetAppID?: string
  langID?: string
  targetLangID?: string
}

export type GetMessagesByOtherAppLangResponse = {
  infos?: Message[]
}

export type GetMessageByAppLangMessageRequest = {
  appID?: string
  langID?: string
  messageID?: string
}

export type GetMessageByAppLangMessageResponse = {
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

export type CreateAppLangForOtherAppRequest = {
  targetAppID?: string
  info?: AppLang
}

export type CreateAppLangForOtherAppResponse = {
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
}

export type GetAppLangInfosByAppResponse = {
  infos?: AppLangInfo[]
}

export type GetAppLangInfosByOtherAppRequest = {
  targetAppID?: string
}

export type GetAppLangInfosByOtherAppResponse = {
  infos?: AppLangInfo[]
}

export type Country = {
  id?: string
  country?: string
  flag?: string
  code?: string
  short?: string
}

export type CreateCountryRequest = {
  info?: Country
}

export type CreateCountryResponse = {
  info?: Country
}

export type CreateCountriesRequest = {
  infos?: Country[]
}

export type CreateCountriesResponse = {
  infos?: Country[]
}

export type GetCountryRequest = {
  id?: string
}

export type GetCountryResponse = {
  info?: Country
}

export type UpdateCountryRequest = {
  info?: Country
}

export type UpdateCountryResponse = {
  info?: Country
}

export type GetCountriesRequest = {
}

export type GetCountriesResponse = {
  infos?: Country[]
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
  static CreateAppLangForOtherApp(req: CreateAppLangForOtherAppRequest, initReq?: fm.InitReq): Promise<CreateAppLangForOtherAppResponse> {
    return fm.fetchReq<CreateAppLangForOtherAppRequest, CreateAppLangForOtherAppResponse>(`/v1/create/app/lang/for/other/app`, {...initReq, method: "POST", body: JSON.stringify(req)})
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
  static GetAppLangInfosByOtherApp(req: GetAppLangInfosByOtherAppRequest, initReq?: fm.InitReq): Promise<GetAppLangInfosByOtherAppResponse> {
    return fm.fetchReq<GetAppLangInfosByOtherAppRequest, GetAppLangInfosByOtherAppResponse>(`/v1/get/app/lang/infos/by/other/app`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static CreateMessage(req: CreateMessageRequest, initReq?: fm.InitReq): Promise<CreateMessageResponse> {
    return fm.fetchReq<CreateMessageRequest, CreateMessageResponse>(`/v1/create/message`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static CreateMessageForOtherApp(req: CreateMessageForOtherAppRequest, initReq?: fm.InitReq): Promise<CreateMessageForOtherAppResponse> {
    return fm.fetchReq<CreateMessageForOtherAppRequest, CreateMessageForOtherAppResponse>(`/v1/create/message/for/other/app`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static CreateMessages(req: CreateMessagesRequest, initReq?: fm.InitReq): Promise<CreateMessagesResponse> {
    return fm.fetchReq<CreateMessagesRequest, CreateMessagesResponse>(`/v1/create/messages`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static CreateMessagesForOtherApp(req: CreateMessagesForOtherAppRequest, initReq?: fm.InitReq): Promise<CreateMessagesForOtherAppResponse> {
    return fm.fetchReq<CreateMessagesForOtherAppRequest, CreateMessagesForOtherAppResponse>(`/v1/create/messages/for/other/app`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static UpdateMessage(req: UpdateMessageRequest, initReq?: fm.InitReq): Promise<UpdateMessageResponse> {
    return fm.fetchReq<UpdateMessageRequest, UpdateMessageResponse>(`/v1/update/message`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static UpdateMessages(req: UpdateMessagesRequest, initReq?: fm.InitReq): Promise<UpdateMessagesResponse> {
    return fm.fetchReq<UpdateMessagesRequest, UpdateMessagesResponse>(`/v1/update/messages`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetMessagesByAppLang(req: GetMessagesByAppLangRequest, initReq?: fm.InitReq): Promise<GetMessagesByAppLangResponse> {
    return fm.fetchReq<GetMessagesByAppLangRequest, GetMessagesByAppLangResponse>(`/v1/get/messages/by/app/lang`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetMessagesByOtherAppLang(req: GetMessagesByOtherAppLangRequest, initReq?: fm.InitReq): Promise<GetMessagesByOtherAppLangResponse> {
    return fm.fetchReq<GetMessagesByOtherAppLangRequest, GetMessagesByOtherAppLangResponse>(`/v1/get/messages/by/other/app/lang`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetMessageByAppLangMessage(req: GetMessageByAppLangMessageRequest, initReq?: fm.InitReq): Promise<GetMessageByAppLangMessageResponse> {
    return fm.fetchReq<GetMessageByAppLangMessageRequest, GetMessageByAppLangMessageResponse>(`/v1/get/message/by/app/lang/message`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static CreateCountry(req: CreateCountryRequest, initReq?: fm.InitReq): Promise<CreateCountryResponse> {
    return fm.fetchReq<CreateCountryRequest, CreateCountryResponse>(`/v1/create/country`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static CreateCountries(req: CreateCountriesRequest, initReq?: fm.InitReq): Promise<CreateCountriesResponse> {
    return fm.fetchReq<CreateCountriesRequest, CreateCountriesResponse>(`/v1/create/countries`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetCountry(req: GetCountryRequest, initReq?: fm.InitReq): Promise<GetCountryResponse> {
    return fm.fetchReq<GetCountryRequest, GetCountryResponse>(`/v1/get/country`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static UpdateCountry(req: UpdateCountryRequest, initReq?: fm.InitReq): Promise<UpdateCountryResponse> {
    return fm.fetchReq<UpdateCountryRequest, UpdateCountryResponse>(`/v1/update/country`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetCountries(req: GetCountriesRequest, initReq?: fm.InitReq): Promise<GetCountriesResponse> {
    return fm.fetchReq<GetCountriesRequest, GetCountriesResponse>(`/v1/get/countries`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
}