/* eslint-disable */
// @ts-nocheck
/*
* This file is a generated Typescript file for GRPC Gateway, DO NOT MODIFY
*/

import * as fm from "../../fetch.pb"
import * as GoogleProtobufEmpty from "../../google/protobuf/empty.pb"
import * as ApplicationManagementV1Application-management from "../application/application-management.pb"
import * as NpoolV1Npool from "../npool.pb"
export type UserBasicInfo = {
  userID?: string
  username?: string
  password?: string
  avatar?: string
  age?: number
  gender?: string
  region?: string
  birthday?: string
  country?: string
  province?: string
  city?: string
  phoneNumber?: string
  emailAddress?: string
  createAt?: number
  updateAt?: number
  signupMethod?: string
  career?: string
  displayName?: string
  firstName?: string
  lastName?: string
  streetAddress1?: string
  streetAddress2?: string
  compony?: string
  postalCode?: string
  appID?: string
}

export type SignupRequest = {
  appID?: string
  username?: string
  password?: string
  emailAddress?: string
  phoneNumber?: string
  code?: string
}

export type SignupResponse = {
  info?: UserBasicInfo
}

export type GetUserRequest = {
  appID?: string
  userID?: string
}

export type GetUserResponse = {
  info?: UserBasicInfo
}

export type GetUsersRequest = {
  appID?: string
  info?: NpoolV1Npool.PageInfo
}

export type GetUsersResponse = {
  infos?: UserBasicInfo[]
}

export type UpdateUserInfoRequest = {
  appID?: string
  userID?: string
  info?: UserBasicInfo
}

export type UpdateUserInfoResponse = {
  info?: UserBasicInfo
}

export type BindUserPhoneRequest = {
  appID?: string
  userID?: string
  phoneNumber?: string
  code?: string
}

export type BindUserPhoneResponse = {
  info?: string
}

export type BindUserEmailRequest = {
  appID?: string
  userID?: string
  emailAddress?: string
  code?: string
}

export type BindUserEmailResponse = {
  info?: string
}

export type UserProvider = {
  id?: string
  userID?: string
  providerID?: string
  providerUserID?: string
  userProviderInfo?: string
  createAt?: number
  updateAt?: number
}

export type BindThirdPartyRequest = {
  appID?: string
  userID?: string
  providerID?: string
  providerUserID?: string
  userProviderInfo?: string
}

export type BindThirdPartyResponse = {
  info?: UserProvider
}

export type UnbindThirdPartyRequest = {
  appID?: string
  userID?: string
  providerID?: string
}

export type UnbindThirdPartyResponse = {
  info?: UserProvider
}

export type ChangeUserPasswordRequest = {
  appID?: string
  userID?: string
  verifyParam?: string
  verifyType?: string
  oldPassword?: string
  password?: string
  code?: string
}

export type ChangeUserPasswordResponse = {
  info?: string
}

export type CertificateKycRequest = {
  appID?: string
  userID?: string
  firstName?: string
  lastName?: string
  frontCardImg?: string
  backCardImg?: string
  userCardImg?: string
  cardType?: string
  cardID?: string
  region?: string
}

export type CertificateKycResponse = {
  info?: string
}

export type GetGaQRCodeRequest = {
  appID?: string
  userID?: string
}

export type GetGaQRCodeResponse = {
  info?: string
}

export type ForgetPasswordRequest = {
  appID?: string
  verifyParam?: string
  verifyType?: string
  password?: string
  code?: string
}

export type ForgetPasswordResponse = {
  info?: string
}

export type DeleteUserRequest = {
  appID?: string
  deleteUserIDs?: string[]
}

export type DeleteUserResponse = {
  info?: string
}

export type AddUserRequest = {
  appID?: string
  userInfo?: UserBasicInfo
}

export type AddUserResponse = {
  info?: UserBasicInfo
}

export type FrozenUser = {
  id?: string
  userID?: string
  frozenBy?: string
  frozenCause?: string
  startAt?: number
  endAt?: number
  status?: string
  unfrozenBy?: string
}

export type FrozenUserRequest = {
  appID?: string
  userID?: string
  frozenBy?: string
  frozenCause?: string
}

export type FrozenUserResponse = {
  info?: FrozenUser
}

export type QueryUserFrozenRequest = {
  appID?: string
  userID?: string
}

export type QueryUserFrozenResponse = {
  info?: FrozenUser
}

export type UnfrozenUserRequest = {
  appID?: string
  id?: string
  userID?: string
  unfrozenBy?: string
}

export type UnfrozenUserResponse = {
  info?: FrozenUser
}

export type GetFrozenUsersRequest = {
  appID?: string
  info?: NpoolV1Npool.PageInfo
}

export type GetFrozenUsersResponse = {
  infos?: FrozenUser[]
}

export type GetUserProvidersRequest = {
  appID?: string
  userID?: string
}

export type GetUserProvidersResponse = {
  infos?: UserProvider[]
}

export type QueryUserExistRequest = {
  appID?: string
  username?: string
  password?: string
}

export type QueryUserExistResponse = {
  info?: UserBasicInfo
}

export type QueryProviderUserInfo = {
  userProviderInfo?: UserProvider
  userBasicInfo?: UserBasicInfo
}

export type QueryUserByUserProviderIDRequest = {
  appID?: string
  providerID?: string
  providerUserID?: string
}

export type QueryUserByUserProviderIDResponse = {
  info?: QueryProviderUserInfo
}

export type SetPasswordRequest = {
  appID?: string
  username?: string
  password?: string
}

export type SetPasswordResponse = {
  info?: string
}

export type GetUserDetailsRequest = {
  appID?: string
  userID?: string
}

export type UserDetails = {
  userBasicInfo?: UserBasicInfo
  userAppInfo?: ApplicationManagementV1Application-management.ApplicationUserDetail
}

export type GetUserDetailsResponse = {
  info?: UserDetails
}

export type UpdateUserEmailRequest = {
  appID?: string
  userID?: string
  oldEmail?: string
  oldCode?: string
  newEmail?: string
  newCode?: string
}

export type UpdateUserEmailResponse = {
  info?: string
}

export type UpdateUserPhoneRequest = {
  appID?: string
  userID?: string
  oldPhone?: string
  oldCode?: string
  newPhone?: string
  newCode?: string
}

export type UpdateUserPhoneResponse = {
  info?: string
}

export class User {
  static Version(req: GoogleProtobufEmpty.Empty, initReq?: fm.InitReq): Promise<NpoolV1Npool.VersionResponse> {
    return fm.fetchReq<GoogleProtobufEmpty.Empty, NpoolV1Npool.VersionResponse>(`/version`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static SignUp(req: SignupRequest, initReq?: fm.InitReq): Promise<SignupResponse> {
    return fm.fetchReq<SignupRequest, SignupResponse>(`/v1/signup`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetUser(req: GetUserRequest, initReq?: fm.InitReq): Promise<GetUserResponse> {
    return fm.fetchReq<GetUserRequest, GetUserResponse>(`/v1/get/user`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetUserDetails(req: GetUserDetailsRequest, initReq?: fm.InitReq): Promise<GetUserDetailsResponse> {
    return fm.fetchReq<GetUserDetailsRequest, GetUserDetailsResponse>(`/v1/get/user/details`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetUsers(req: GetUsersRequest, initReq?: fm.InitReq): Promise<GetUsersResponse> {
    return fm.fetchReq<GetUsersRequest, GetUsersResponse>(`/v1/get/users`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static UpdateUserInfo(req: UpdateUserInfoRequest, initReq?: fm.InitReq): Promise<UpdateUserInfoResponse> {
    return fm.fetchReq<UpdateUserInfoRequest, UpdateUserInfoResponse>(`/v1/update/user`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static BindUserPhone(req: BindUserPhoneRequest, initReq?: fm.InitReq): Promise<BindUserPhoneResponse> {
    return fm.fetchReq<BindUserPhoneRequest, BindUserPhoneResponse>(`/v1/bind/phone`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static BindUserEmail(req: BindUserEmailRequest, initReq?: fm.InitReq): Promise<BindUserEmailResponse> {
    return fm.fetchReq<BindUserEmailRequest, BindUserEmailResponse>(`/v1/bind/email`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static UpdateUserEmail(req: UpdateUserEmailRequest, initReq?: fm.InitReq): Promise<UpdateUserEmailResponse> {
    return fm.fetchReq<UpdateUserEmailRequest, UpdateUserEmailResponse>(`/v1/update/user/email`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static UpdateUserPhone(req: UpdateUserPhoneRequest, initReq?: fm.InitReq): Promise<UpdateUserPhoneResponse> {
    return fm.fetchReq<UpdateUserPhoneRequest, UpdateUserPhoneResponse>(`/v1/update/user/phone`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static BindThirdParty(req: BindThirdPartyRequest, initReq?: fm.InitReq): Promise<BindThirdPartyResponse> {
    return fm.fetchReq<BindThirdPartyRequest, BindThirdPartyResponse>(`/v1/bind/thirdparty`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static UnbindThirdParty(req: UnbindThirdPartyRequest, initReq?: fm.InitReq): Promise<UnbindThirdPartyResponse> {
    return fm.fetchReq<UnbindThirdPartyRequest, UnbindThirdPartyResponse>(`/v1/unbind/thirdparty`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static ChangeUserPassword(req: ChangeUserPasswordRequest, initReq?: fm.InitReq): Promise<ChangeUserPasswordResponse> {
    return fm.fetchReq<ChangeUserPasswordRequest, ChangeUserPasswordResponse>(`/v1/change/password`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static ForgetPassword(req: ForgetPasswordRequest, initReq?: fm.InitReq): Promise<ForgetPasswordResponse> {
    return fm.fetchReq<ForgetPasswordRequest, ForgetPasswordResponse>(`/v1/forget/password`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static SetPassword(req: SetPasswordRequest, initReq?: fm.InitReq): Promise<SetPasswordResponse> {
    return fm.fetchReq<SetPasswordRequest, SetPasswordResponse>(`/v1/set/password`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static AddUser(req: AddUserRequest, initReq?: fm.InitReq): Promise<AddUserResponse> {
    return fm.fetchReq<AddUserRequest, AddUserResponse>(`/v1/add/user`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static DeleteUser(req: DeleteUserRequest, initReq?: fm.InitReq): Promise<DeleteUserResponse> {
    return fm.fetchReq<DeleteUserRequest, DeleteUserResponse>(`/v1/delete/users`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static FrozenUser(req: FrozenUserRequest, initReq?: fm.InitReq): Promise<FrozenUserResponse> {
    return fm.fetchReq<FrozenUserRequest, FrozenUserResponse>(`/v1/frozen/user`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static UnfrozenUser(req: UnfrozenUserRequest, initReq?: fm.InitReq): Promise<UnfrozenUserResponse> {
    return fm.fetchReq<UnfrozenUserRequest, UnfrozenUserResponse>(`/v1/unfrozen/user`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static QueryUserFrozen(req: QueryUserFrozenRequest, initReq?: fm.InitReq): Promise<QueryUserFrozenResponse> {
    return fm.fetchReq<QueryUserFrozenRequest, QueryUserFrozenResponse>(`/v1/query/user/frozen`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetFrozenUsers(req: GetFrozenUsersRequest, initReq?: fm.InitReq): Promise<GetFrozenUsersResponse> {
    return fm.fetchReq<GetFrozenUsersRequest, GetFrozenUsersResponse>(`/v1/get/frozen/user`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetUserProviders(req: GetUserProvidersRequest, initReq?: fm.InitReq): Promise<GetUserProvidersResponse> {
    return fm.fetchReq<GetUserProvidersRequest, GetUserProvidersResponse>(`/v1/get/user/providers`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static QueryUserExist(req: QueryUserExistRequest, initReq?: fm.InitReq): Promise<QueryUserExistResponse> {
    return fm.fetchReq<QueryUserExistRequest, QueryUserExistResponse>(`/v1/query/user/exist`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static QueryUserByUserProviderID(req: QueryUserByUserProviderIDRequest, initReq?: fm.InitReq): Promise<QueryUserByUserProviderIDResponse> {
    return fm.fetchReq<QueryUserByUserProviderIDRequest, QueryUserByUserProviderIDResponse>(`/v1/query/user/by/userproviderid`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
}