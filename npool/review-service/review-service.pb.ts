/* eslint-disable */
// @ts-nocheck
/*
* This file is a generated Typescript file for GRPC Gateway, DO NOT MODIFY
*/

import * as fm from "../../fetch.pb"
import * as GoogleProtobufEmpty from "../../google/protobuf/empty.pb"
export type VersionResponse = {
  Info?: string
}

export type Review = {
  ID?: string
  ObjectType?: string
  AppID?: string
  ReviewerID?: string
  State?: string
  Message?: string
  ObjectID?: string
  Domain?: string
  CreateAt?: number
}

export type CreateReviewRequest = {
  Info?: Review
}

export type CreateReviewResponse = {
  Info?: Review
}

export type UpdateReviewRequest = {
  Info?: Review
}

export type UpdateReviewResponse = {
  Info?: Review
}

export type GetReviewsByDomainRequest = {
  Domain?: string
}

export type GetReviewsByDomainResponse = {
  Infos?: Review[]
}

export type GetReviewsByAppDomainRequest = {
  AppID?: string
  Domain?: string
}

export type GetReviewsByAppDomainResponse = {
  Infos?: Review[]
}

export type GetReviewsByAppDomainObjectTypeIDRequest = {
  AppID?: string
  Domain?: string
  ObjectType?: string
  ObjectID?: string
}

export type GetReviewsByAppDomainObjectTypeIDResponse = {
  Infos?: Review[]
}

export type SubmitReviewRequest = {
  Info?: Review
}

export type SubmitReviewResponse = {
  Info?: Review
}

export type ReviewRule = {
  ID?: string
  ObjectType?: string
  Domain?: string
  Rules?: string
}

export type CreateReviewRuleRequest = {
  Info?: ReviewRule
}

export type CreateReviewRuleResponse = {
  Info?: ReviewRule
}

export type UpdateReviewRuleRequest = {
  Info?: ReviewRule
}

export type UpdateReviewRuleResponse = {
  Info?: ReviewRule
}

export type GetReviewRuleRequest = {
  ID?: string
}

export type GetReviewRuleResponse = {
  Info?: ReviewRule
}

export type GetReviewRulesByDomainRequest = {
  Domain?: string
}

export type GetReviewRulesByDomainResponse = {
  Infos?: ReviewRule[]
}

export type GetReviewRuleByDomainObjectTypeRequest = {
  Domain?: string
  ObjectType?: string
}

export type GetReviewRuleByDomainObjectTypeResponse = {
  Info?: ReviewRule
}

export type SubmitReviewResultRequest = {
  Info?: Review
}

export type SubmitReviewResultResponse = {
  Info?: Review
}

export class ReviewService {
  static Version(req: GoogleProtobufEmpty.Empty, initReq?: fm.InitReq): Promise<VersionResponse> {
    return fm.fetchReq<GoogleProtobufEmpty.Empty, VersionResponse>(`/version`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static CreateReview(req: CreateReviewRequest, initReq?: fm.InitReq): Promise<CreateReviewResponse> {
    return fm.fetchReq<CreateReviewRequest, CreateReviewResponse>(`/v1/create/review`, {...initReq, method: "POST", body: JSON.stringify(req)})
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