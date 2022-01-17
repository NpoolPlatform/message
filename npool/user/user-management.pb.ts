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
  UserID?: string
  Username?: string
  Password?: string
  Avatar?: string
  Age?: number
  Gender?: string
  Region?: string
  Birthday?: string
  Country?: string
  Province?: string
  City?: string
  PhoneNumber?: string
  EmailAddress?: string
  CreateAt?: number
  UpdateAt?: number
  SignupMethod?: string
  Career?: string
  DisplayName?: string
  FirstName?: string
  LastName?: string
  StreetAddress1?: string
  StreetAddress2?: string
  Compony?: string
  PostalCode?: string
}

export type SignupRequest = {
  Username?: string
  Password?: string
  EmailAddress?: string
  PhoneNumber?: string
  Code?: string
  AppID?: string
}

export type SignupResponse = {
  Info?: UserBasicInfo
}

export type GetUserRequest = {
  UserID?: string
  AppID?: string
}

export type GetUserResponse = {
  Info?: UserBasicInfo
}

export type GetUsersRequest = {
  Info?: NpoolV1Npool.PageInfo
  AppID?: string
}

export type GetUsersResponse = {
  Infos?: UserBasicInfo[]
}

export type UpdateUserInfoRequest = {
  Info?: UserBasicInfo
  AppID?: string
  UserID?: string
}

export type UpdateUserInfoResponse = {
  Info?: UserBasicInfo
}

export type BindUserPhoneRequest = {
  UserID?: string
  PhoneNumber?: string
  AppID?: string
  Code?: string
}

export type BindUserPhoneResponse = {
  Info?: string
}

export type BindUserEmailRequest = {
  UserID?: string
  EmailAddress?: string
  Code?: string
  AppID?: string
}

export type BindUserEmailResponse = {
  Info?: string
}

export type UserProvider = {
  ID?: string
  UserID?: string
  ProviderID?: string
  ProviderUserID?: string
  UserProviderInfo?: string
  CreateAt?: number
  UpdateAt?: number
}

export type BindThirdPartyRequest = {
  UserID?: string
  AppID?: string
  ProviderID?: string
  ProviderUserID?: string
  UserProviderInfo?: string
}

export type BindThirdPartyResponse = {
  Info?: UserProvider
}

export type UnbindThirdPartyRequest = {
  UserID?: string
  AppID?: string
  ProviderID?: string
}

export type UnbindThirdPartyResponse = {
  Info?: UserProvider
}

export type ChangeUserPasswordRequest = {
  UserID?: string
  AppID?: string
  VerifyParam?: string
  VerifyType?: string
  OldPassword?: string
  Password?: string
  Code?: string
}

export type ChangeUserPasswordResponse = {
  Info?: string
}

export type CertificateKycRequest = {
  UserID?: string
  AppID?: string
  FirstName?: string
  LastName?: string
  FrontCardImg?: string
  BackCardImg?: string
  UserCardImg?: string
  CardType?: string
  CardID?: string
  Region?: string
}

export type CertificateKycResponse = {
  Info?: string
}

export type GetGaQRCodeRequest = {
  UserID?: string
  AppID?: string
}

export type GetGaQRCodeResponse = {
  Info?: string
}

export type ForgetPasswordRequest = {
  VerifyParam?: string
  VerifyType?: string
  Password?: string
  Code?: string
  AppID?: string
}

export type ForgetPasswordResponse = {
  Info?: string
}

export type DeleteUserRequest = {
  DeleteUserIDs?: string[]
  AppID?: string
}

export type DeleteUserResponse = {
  Info?: string
}

export type AddUserRequest = {
  AppID?: string
  UserInfo?: UserBasicInfo
}

export type AddUserResponse = {
  Info?: UserBasicInfo
}

export type FrozenUser = {
  ID?: string
  UserID?: string
  FrozenBy?: string
  FrozenCause?: string
  StartAt?: number
  EndAt?: number
  Status?: string
  UnfrozenBy?: string
}

export type FrozenUserRequest = {
  UserID?: string
  FrozenBy?: string
  FrozenCause?: string
  AppID?: string
}

export type FrozenUserResponse = {
  Info?: FrozenUser
}

export type QueryUserFrozenRequest = {
  UserID?: string
  AppID?: string
}

export type QueryUserFrozenResponse = {
  Info?: FrozenUser
}

export type UnfrozenUserRequest = {
  ID?: string
  UserID?: string
  UnfrozenBy?: string
  AppID?: string
}

export type UnfrozenUserResponse = {
  Info?: FrozenUser
}

export type GetFrozenUsersRequest = {
  Info?: NpoolV1Npool.PageInfo
  AppID?: string
}

export type GetFrozenUsersResponse = {
  Infos?: FrozenUser[]
}

export type GetUserProvidersRequest = {
  UserID?: string
  AppID?: string
}

export type GetUserProvidersResponse = {
  Infos?: UserProvider[]
}

export type QueryUserExistRequest = {
  Username?: string
  Password?: string
  AppID?: string
}

export type QueryUserExistResponse = {
  Info?: UserBasicInfo
}

export type QueryProviderUserInfo = {
  UserProviderInfo?: UserProvider
  UserBasicInfo?: UserBasicInfo
}

export type QueryUserByUserProviderIDRequest = {
  ProviderID?: string
  ProviderUserID?: string
  AppID?: string
}

export type QueryUserByUserProviderIDResponse = {
  Info?: QueryProviderUserInfo
}

export type SetPasswordRequest = {
  Username?: string
  Password?: string
  AppID?: string
}

export type SetPasswordResponse = {
  Info?: string
}

export type GetUserDetailsRequest = {
  AppID?: string
  UserID?: string
}

export type UserDetails = {
  UserBasicInfo?: UserBasicInfo
  UserAppInfo?: ApplicationManagementV1Application-management.ApplicationUserDetail
}

export type GetUserDetailsResponse = {
  Info?: UserDetails
}

export type UpdateUserEmailRequest = {
  UserID?: string
  AppID?: string
  OldEmail?: string
  OldCode?: string
  NewEmail?: string
  NewCode?: string
}

export type UpdateUserEmailResponse = {
  Info?: string
}

export type UpdateUserPhoneRequest = {
  UserID?: string
  AppID?: string
  OldPhone?: string
  OldCode?: string
  NewPhone?: string
  NewCode?: string
}

export type UpdateUserPhoneResponse = {
  Info?: string
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