/* eslint-disable */
// @ts-nocheck
/*
* This file is a generated Typescript file for GRPC Gateway, DO NOT MODIFY
*/

import * as fm from "../../fetch.pb"
import * as GoogleProtobufEmpty from "../../google/protobuf/empty.pb"
import * as NpoolV1Npool from "../npool.pb"
export type GoodPaying = {
  ID?: string
  OrderID?: string
  PaymentID?: string
}

export type CreateGoodPayingRequest = {
  Info?: GoodPaying
}

export type CreateGoodPayingResponse = {
  Info?: GoodPaying
}

export type GetGoodPayingByOrderRequest = {
  OrderID?: string
}

export type GetGoodPayingByOrderResponse = {
  Info?: GoodPaying
}

export type GasPaying = {
  ID?: string
  OrderID?: string
  PaymentID?: string
  DurationMinutes?: number
  FeeTypeID?: string
}

export type CreateGasPayingRequest = {
  Info?: GasPaying
}

export type CreateGasPayingResponse = {
  Info?: GasPaying
}

export type GetGasPayingsByOrderRequest = {
  OrderID?: string
}

export type GetGasPayingsByOrderResponse = {
  Infos?: GasPaying[]
}

export type Order = {
  ID?: string
  GoodID?: string
  AppID?: string
  UserID?: string
  Units?: number
  DiscountCouponID?: string
  UserSpecialReductionID?: string
  Start?: number
  End?: number
  CouponID?: string
}

export type CreateOrderRequest = {
  Info?: Order
}

export type CreateOrderResponse = {
  Info?: Order
}

export type GetOrderRequest = {
  ID?: string
}

export type GetOrderResponse = {
  Info?: Order
}

export type GetOrdersByAppUserRequest = {
  AppID?: string
  UserID?: string
}

export type GetOrdersByAppUserResponse = {
  Infos?: Order[]
}

export type GetOrdersByAppRequest = {
  AppID?: string
}

export type GetOrdersByAppResponse = {
  Infos?: Order[]
}

export type GetOrdersByGoodRequest = {
  GoodID?: string
}

export type GetOrdersByGoodResponse = {
  Infos?: Order[]
}

export type Compensate = {
  ID?: string
  OrderID?: string
  Start?: number
  End?: number
  Message?: string
}

export type CreateCompensateRequest = {
  Info?: Compensate
}

export type CreateCompensateResponse = {
  Info?: Compensate
}

export type GetCompensatesByOrderRequest = {
  OrderID?: string
}

export type GetCompensatesByOrderResponse = {
  Infos?: Compensate[]
}

export type OutOfGas = {
  ID?: string
  OrderID?: string
  Start?: number
  End?: number
}

export type CreateOutOfGasRequest = {
  Info?: OutOfGas
}

export type CreateOutOfGasResponse = {
  Info?: OutOfGas
}

export type GetOutOfGasesByOrderRequest = {
  OrderID?: string
}

export type GetOutOfGasesByOrderResponse = {
  Infos?: OutOfGas[]
}

export type Payment = {
  ID?: string
  OrderID?: string
  AccountID?: string
  StartAmount?: number
  Amount?: number
  CoinInfoID?: string
  State?: string
  ChainTransactionID?: string
  PlatformTransactionID?: string
}

export type CreatePaymentRequest = {
  Info?: Payment
}

export type CreatePaymentResponse = {
  Info?: Payment
}

export type GetPaymentRequest = {
  ID?: string
}

export type GetPaymentResponse = {
  Info?: Payment
}

export type UpdatePaymentRequest = {
  Info?: Payment
}

export type UpdatePaymentResponse = {
  Info?: Payment
}

export type GetPaymentByOrderRequest = {
  OrderID?: string
}

export type GetPaymentByOrderResponse = {
  Info?: Payment
}

export type OrderDetail = {
  ID?: string
  GoodID?: string
  AppID?: string
  UserID?: string
  Units?: number
  DiscountCouponID?: string
  UserSpecialReductionID?: string
  GoodPaying?: GoodPaying
  GasPayings?: GasPaying[]
  Compensates?: Compensate[]
  OutOfGases?: OutOfGas[]
  Payment?: Payment
  Start?: number
  End?: number
  CouponID?: string
}

export type GetOrderDetailRequest = {
  ID?: string
}

export type GetOrderDetailResponse = {
  Detail?: OrderDetail
}

export type GetOrdersDetailByAppUserRequest = {
  AppID?: string
  UserID?: string
}

export type GetOrdersDetailByAppUserResponse = {
  Details?: OrderDetail[]
}

export type GetOrdersShortDetailByAppUserRequest = {
  AppID?: string
  UserID?: string
}

export type GetOrdersShortDetailByAppUserResponse = {
  Details?: OrderDetail[]
}

export type GetOrdersDetailByAppRequest = {
  AppID?: string
}

export type GetOrdersDetailByAppResponse = {
  Details?: OrderDetail[]
}

export type GetOrdersDetailByGoodRequest = {
  GoodID?: string
}

export type GetOrdersDetailByGoodResponse = {
  Details?: OrderDetail[]
}

export class CloudHashingOrder {
  static Version(req: GoogleProtobufEmpty.Empty, initReq?: fm.InitReq): Promise<NpoolV1Npool.VersionResponse> {
    return fm.fetchReq<GoogleProtobufEmpty.Empty, NpoolV1Npool.VersionResponse>(`/version`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static CreateGoodPaying(req: CreateGoodPayingRequest, initReq?: fm.InitReq): Promise<CreateGoodPayingResponse> {
    return fm.fetchReq<CreateGoodPayingRequest, CreateGoodPayingResponse>(`/v1/create/good/paying`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetGoodPayingByOrder(req: GetGoodPayingByOrderRequest, initReq?: fm.InitReq): Promise<GetGoodPayingByOrderResponse> {
    return fm.fetchReq<GetGoodPayingByOrderRequest, GetGoodPayingByOrderResponse>(`/v1/get/good/paying/by/order`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static CreateGasPaying(req: CreateGasPayingRequest, initReq?: fm.InitReq): Promise<CreateGasPayingResponse> {
    return fm.fetchReq<CreateGasPayingRequest, CreateGasPayingResponse>(`/v1/create/gas/paying`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetGasPayingsByOrder(req: GetGasPayingsByOrderRequest, initReq?: fm.InitReq): Promise<GetGasPayingsByOrderResponse> {
    return fm.fetchReq<GetGasPayingsByOrderRequest, GetGasPayingsByOrderResponse>(`/v1/get/gas/payings/by/order`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static CreateOrder(req: CreateOrderRequest, initReq?: fm.InitReq): Promise<CreateOrderResponse> {
    return fm.fetchReq<CreateOrderRequest, CreateOrderResponse>(`/v1/create/order`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetOrder(req: GetOrderRequest, initReq?: fm.InitReq): Promise<GetOrderResponse> {
    return fm.fetchReq<GetOrderRequest, GetOrderResponse>(`/v1/get/order`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetOrderDetail(req: GetOrderDetailRequest, initReq?: fm.InitReq): Promise<GetOrderDetailResponse> {
    return fm.fetchReq<GetOrderDetailRequest, GetOrderDetailResponse>(`/v1/get/order/detail`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetOrdersByAppUser(req: GetOrdersByAppUserRequest, initReq?: fm.InitReq): Promise<GetOrdersByAppUserResponse> {
    return fm.fetchReq<GetOrdersByAppUserRequest, GetOrdersByAppUserResponse>(`/v1/get/orders/by/app/user`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetOrdersByApp(req: GetOrdersByAppRequest, initReq?: fm.InitReq): Promise<GetOrdersByAppResponse> {
    return fm.fetchReq<GetOrdersByAppRequest, GetOrdersByAppResponse>(`/v1/get/orders/by/app`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetOrdersByGood(req: GetOrdersByGoodRequest, initReq?: fm.InitReq): Promise<GetOrdersByGoodResponse> {
    return fm.fetchReq<GetOrdersByGoodRequest, GetOrdersByGoodResponse>(`/v1/get/orders/by/good`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static CreateCompensate(req: CreateCompensateRequest, initReq?: fm.InitReq): Promise<CreateCompensateResponse> {
    return fm.fetchReq<CreateCompensateRequest, CreateCompensateResponse>(`/v1/create/compensate`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetCompensatesByOrder(req: GetCompensatesByOrderRequest, initReq?: fm.InitReq): Promise<GetCompensatesByOrderResponse> {
    return fm.fetchReq<GetCompensatesByOrderRequest, GetCompensatesByOrderResponse>(`/v1/get/compensates/by/order`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static CreateOutOfGas(req: CreateOutOfGasRequest, initReq?: fm.InitReq): Promise<CreateOutOfGasResponse> {
    return fm.fetchReq<CreateOutOfGasRequest, CreateOutOfGasResponse>(`/v1/create/out-of-gas`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetOutOfGasesByOrder(req: GetOutOfGasesByOrderRequest, initReq?: fm.InitReq): Promise<GetOutOfGasesByOrderResponse> {
    return fm.fetchReq<GetOutOfGasesByOrderRequest, GetOutOfGasesByOrderResponse>(`/v1/get/out-of-gases/by/order`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static CreatePayment(req: CreatePaymentRequest, initReq?: fm.InitReq): Promise<CreatePaymentResponse> {
    return fm.fetchReq<CreatePaymentRequest, CreatePaymentResponse>(`/v1/create/payment`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetPayment(req: GetPaymentRequest, initReq?: fm.InitReq): Promise<GetPaymentResponse> {
    return fm.fetchReq<GetPaymentRequest, GetPaymentResponse>(`/v1/get/payment`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static UpdatePayment(req: UpdatePaymentRequest, initReq?: fm.InitReq): Promise<UpdatePaymentResponse> {
    return fm.fetchReq<UpdatePaymentRequest, UpdatePaymentResponse>(`/v1/update/payment`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetPaymentByOrder(req: GetPaymentByOrderRequest, initReq?: fm.InitReq): Promise<GetPaymentByOrderResponse> {
    return fm.fetchReq<GetPaymentByOrderRequest, GetPaymentByOrderResponse>(`/v1/get/payment/by/order`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetOrdersDetailByAppUser(req: GetOrdersDetailByAppUserRequest, initReq?: fm.InitReq): Promise<GetOrdersDetailByAppUserResponse> {
    return fm.fetchReq<GetOrdersDetailByAppUserRequest, GetOrdersDetailByAppUserResponse>(`/v1/get/orders/detail/by/app/user`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetOrdersShortDetailByAppUser(req: GetOrdersShortDetailByAppUserRequest, initReq?: fm.InitReq): Promise<GetOrdersShortDetailByAppUserResponse> {
    return fm.fetchReq<GetOrdersShortDetailByAppUserRequest, GetOrdersShortDetailByAppUserResponse>(`/v1/get/orders/short/detail/by/app/user`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetOrdersDetailByApp(req: GetOrdersDetailByAppRequest, initReq?: fm.InitReq): Promise<GetOrdersDetailByAppResponse> {
    return fm.fetchReq<GetOrdersDetailByAppRequest, GetOrdersDetailByAppResponse>(`/v1/get/orders/detail/by/app`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
  static GetOrdersDetailByGood(req: GetOrdersDetailByGoodRequest, initReq?: fm.InitReq): Promise<GetOrdersDetailByGoodResponse> {
    return fm.fetchReq<GetOrdersDetailByGoodRequest, GetOrdersDetailByGoodResponse>(`/v1/get/orders/detail/by/good`, {...initReq, method: "POST", body: JSON.stringify(req)})
  }
}