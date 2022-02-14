/* eslint-disable */
// @ts-nocheck
/*
* This file is a generated Typescript file for GRPC Gateway, DO NOT MODIFY
*/

import * as fm from "../../fetch.pb"
import * as GoogleProtobufEmpty from "../../google/protobuf/empty.pb"
export type VersionResponse = {
  info?: string
}

export type Review = {
  id?: string
  objectType?: string
  appID?: string
  reviewerID?: string
  state?: string
  message?: string
  objectID?: string
  domain?: string
  createAt?: number
  trigger?: string
}

export type CreateReviewRequest = {
  info?: Review
}

export type CreateReviewResponse = {
  info?: Review
}

export type GetReviewRequest = {
  id?: string
}

export type GetReviewResponse = {
  info?: Review
}

export type UpdateReviewRequest = {
  info?: Review
}

export type UpdateReviewResponse = {
  info?: Review
}

export type GetReviewsByDomainRequest = {
  domain?: string
}

export type GetReviewsByDomainResponse = {
  infos?: Review[]
}

export type GetReviewsByAppDomainRequest = {
  appID?: string
  domain?: string
}

export type GetReviewsByAppDomainResponse = {
  infos?: Review[]
}

export type GetReviewsByAppDomainObjectTypeIDRequest = {
  appID?: string
  domain?: string
  objectType?: string
  objectID?: string
}

export type GetReviewsByAppDomainObjectTypeIDResponse = {
  infos?: Review[]
}

export type SubmitReviewRequest = {
  info?: Review
}

export type SubmitReviewResponse = {
  info?: Review
}

export type ReviewRule = {
  id?: string
  objectType?: string
  domain?: string
  rules?: string
}

export type CreateReviewRuleRequest = {
  info?: ReviewRule
}

export type CreateReviewRuleResponse = {
  info?: ReviewRule
}

export type UpdateReviewRuleRequest = {
  info?: ReviewRule
}

export type UpdateReviewRuleResponse = {
  info?: ReviewRule
}

export type GetReviewRuleRequest = {
  id?: string
}

export type GetReviewRuleResponse = {
  info?: ReviewRule
}

export type GetReviewRulesByDomainRequest = {
  domain?: string
}

export type GetReviewRulesByDomainResponse = {
  infos?: ReviewRule[]
}

export type GetReviewRuleByDomainObjectTypeRequest = {
  domain?: string
  objectType?: string
}

export type GetReviewRuleByDomainObjectTypeResponse = {
  info?: ReviewRule
}

export type SubmitReviewResultRequest = {
  info?: Review
}

export type SubmitReviewResultResponse = {
  info?: Review
}

export class ReviewService {
  static Version(req: GoogleProtobufEmpty.Empty, initReq?: fm.InitReq): Promise<VersionResponse> {
    return fm.fetchReq<GoogleProtobufEmpty.Empty, VersionResponse>(`/version`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static CreateReview(req: CreateReviewRequest, initReq?: fm.InitReq): Promise<CreateReviewResponse> {
    return fm.fetchReq<CreateReviewRequest, CreateReviewResponse>(`/v1/create/review`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetReview(req: GetReviewRequest, initReq?: fm.InitReq): Promise<GetReviewResponse> {
    return fm.fetchReq<GetReviewRequest, GetReviewResponse>(`/v1/get/review`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static UpdateReview(req: UpdateReviewRequest, initReq?: fm.InitReq): Promise<UpdateReviewResponse> {
    return fm.fetchReq<UpdateReviewRequest, UpdateReviewResponse>(`/v1/update/review`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetReviewsByDomain(req: GetReviewsByDomainRequest, initReq?: fm.InitReq): Promise<GetReviewsByDomainResponse> {
    return fm.fetchReq<GetReviewsByDomainRequest, GetReviewsByDomainResponse>(`/v1/get/reviews/by/domain`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetReviewsByAppDomain(req: GetReviewsByAppDomainRequest, initReq?: fm.InitReq): Promise<GetReviewsByAppDomainResponse> {
    return fm.fetchReq<GetReviewsByAppDomainRequest, GetReviewsByAppDomainResponse>(`/v1/get/reviews/by/app/domain`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetReviewsByAppDomainObjectTypeID(req: GetReviewsByAppDomainObjectTypeIDRequest, initReq?: fm.InitReq): Promise<GetReviewsByAppDomainObjectTypeIDResponse> {
    return fm.fetchReq<GetReviewsByAppDomainObjectTypeIDRequest, GetReviewsByAppDomainObjectTypeIDResponse>(`/v1/get/reviews/by/app/domain/object/type/id`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static SubmitReview(req: SubmitReviewRequest, initReq?: fm.InitReq): Promise<SubmitReviewResponse> {
    return fm.fetchReq<SubmitReviewRequest, SubmitReviewResponse>(`/v1/submit/review`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static SubmitReviewResult(req: SubmitReviewResultRequest, initReq?: fm.InitReq): Promise<SubmitReviewResultResponse> {
    return fm.fetchReq<SubmitReviewResultRequest, SubmitReviewResultResponse>(`/v1/submit/review/result`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static CreateReviewRule(req: CreateReviewRuleRequest, initReq?: fm.InitReq): Promise<CreateReviewRuleResponse> {
    return fm.fetchReq<CreateReviewRuleRequest, CreateReviewRuleResponse>(`/v1/create/review/rule`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static UpdateReviewRule(req: UpdateReviewRuleRequest, initReq?: fm.InitReq): Promise<UpdateReviewRuleResponse> {
    return fm.fetchReq<UpdateReviewRuleRequest, UpdateReviewRuleResponse>(`/v1/update/review/rule`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetReviewRule(req: GetReviewRuleRequest, initReq?: fm.InitReq): Promise<GetReviewRuleResponse> {
    return fm.fetchReq<GetReviewRuleRequest, GetReviewRuleResponse>(`/v1/get/review/rule`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetReviewRulesByDomain(req: GetReviewRulesByDomainRequest, initReq?: fm.InitReq): Promise<GetReviewRulesByDomainResponse> {
    return fm.fetchReq<GetReviewRulesByDomainRequest, GetReviewRulesByDomainResponse>(`/v1/get/review/rules/by/domain`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetReviewRuleByDomainObjectType(req: GetReviewRuleByDomainObjectTypeRequest, initReq?: fm.InitReq): Promise<GetReviewRuleByDomainObjectTypeResponse> {
    return fm.fetchReq<GetReviewRuleByDomainObjectTypeRequest, GetReviewRuleByDomainObjectTypeResponse>(`/v1/get/review/rule/by/domain/object/type`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
}