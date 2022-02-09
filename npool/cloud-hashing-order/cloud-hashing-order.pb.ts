/* eslint-disable */
// @ts-nocheck
/*
* This file is a generated Typescript file for GRPC Gateway, DO NOT MODIFY
*/

import * as fm from "../../fetch.pb"
import * as GoogleProtobufEmpty from "../../google/protobuf/empty.pb"
import * as NpoolV1Npool from "../npool.pb"
export type GoodPaying = {
  id?: string
  orderID?: string
  paymentID?: string
}

export type CreateGoodPayingRequest = {
  info?: GoodPaying
}

export type CreateGoodPayingResponse = {
  info?: GoodPaying
}

export type GetGoodPayingByOrderRequest = {
  orderID?: string
}

export type GetGoodPayingByOrderResponse = {
  info?: GoodPaying
}

export type GasPaying = {
  id?: string
  orderID?: string
  paymentID?: string
  durationMinutes?: number
  feeTypeID?: string
}

export type CreateGasPayingRequest = {
  info?: GasPaying
}

export type CreateGasPayingResponse = {
  info?: GasPaying
}

export type GetGasPayingsByOrderRequest = {
  orderID?: string
}

export type GetGasPayingsByOrderResponse = {
  infos?: GasPaying[]
}

export type Order = {
  id?: string
  goodID?: string
  appID?: string
  userID?: string
  units?: number
  discountCouponID?: string
  userSpecialReductionID?: string
  start?: number
  end?: number
  couponID?: string
}

export type CreateOrderRequest = {
  info?: Order
}

export type CreateOrderResponse = {
  info?: Order
}

export type GetOrderRequest = {
  id?: string
}

export type GetOrderResponse = {
  info?: Order
}

export type GetOrdersByAppUserRequest = {
  appID?: string
  userID?: string
}

export type GetOrdersByAppUserResponse = {
  infos?: Order[]
}

export type GetOrdersByAppRequest = {
  appID?: string
}

export type GetOrdersByAppResponse = {
  infos?: Order[]
}

export type GetOrdersByGoodRequest = {
  goodID?: string
}

export type GetOrdersByGoodResponse = {
  infos?: Order[]
}

export type Compensate = {
  id?: string
  orderID?: string
  start?: number
  end?: number
  message?: string
}

export type CreateCompensateRequest = {
  info?: Compensate
}

export type CreateCompensateResponse = {
  info?: Compensate
}

export type GetCompensatesByOrderRequest = {
  orderID?: string
}

export type GetCompensatesByOrderResponse = {
  infos?: Compensate[]
}

export type OutOfGas = {
  id?: string
  orderID?: string
  start?: number
  end?: number
}

export type CreateOutOfGasRequest = {
  info?: OutOfGas
}

export type CreateOutOfGasResponse = {
  info?: OutOfGas
}

export type GetOutOfGasesByOrderRequest = {
  orderID?: string
}

export type GetOutOfGasesByOrderResponse = {
  infos?: OutOfGas[]
}

export type Payment = {
  id?: string
  orderID?: string
  accountID?: string
  startAmount?: number
  amount?: number
  coinInfoID?: string
  state?: string
  chainTransactionID?: string
  platformTransactionID?: string
  createAt?: number
}

export type CreatePaymentRequest = {
  info?: Payment
}

export type CreatePaymentResponse = {
  info?: Payment
}

export type GetPaymentRequest = {
  id?: string
}

export type GetPaymentResponse = {
  info?: Payment
}

export type UpdatePaymentRequest = {
  info?: Payment
}

export type UpdatePaymentResponse = {
  info?: Payment
}

export type GetPaymentByOrderRequest = {
  orderID?: string
}

export type GetPaymentByOrderResponse = {
  info?: Payment
}

export type OrderDetail = {
  id?: string
  goodID?: string
  appID?: string
  userID?: string
  units?: number
  discountCouponID?: string
  userSpecialReductionID?: string
  goodPaying?: GoodPaying
  gasPayings?: GasPaying[]
  compensates?: Compensate[]
  outOfGases?: OutOfGas[]
  payment?: Payment
  start?: number
  end?: number
  couponID?: string
}

export type GetOrderDetailRequest = {
  id?: string
}

export type GetOrderDetailResponse = {
  detail?: OrderDetail
}

export type GetOrdersDetailByAppUserRequest = {
  appID?: string
  userID?: string
}

export type GetOrdersDetailByAppUserResponse = {
  details?: OrderDetail[]
}

export type GetOrdersShortDetailByAppUserRequest = {
  appID?: string
  userID?: string
}

export type GetOrdersShortDetailByAppUserResponse = {
  details?: OrderDetail[]
}

export type GetOrdersDetailByAppRequest = {
  appID?: string
}

export type GetOrdersDetailByAppResponse = {
  details?: OrderDetail[]
}

export type GetOrdersDetailByGoodRequest = {
  goodID?: string
}

export type GetOrdersDetailByGoodResponse = {
  details?: OrderDetail[]
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