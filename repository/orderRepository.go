package repository

import (
	"benings/config"
	"benings/model"
	"benings/util"

	"github.com/ian-kent/go-log/log"
)

func OrderHeader(request model.CreateOrder) model.CreateOrderResponse {
	var response model.CreateOrderResponse
	db := config.Connect(config.ReadDatabaseConfig())
	defer db.Close()
	orderHeader, err := db.Query("call trx.ins_orderheader ($1,$2,$3,$4,$5,$6)",
		request.Token,
		request.OrderHeader.MembertemplateId,
		request.OrderType,
		request.OrderHeaderId,
		request.OutError,
		request.OutMessage,
	)
	if err != nil {
		response.OutError = 1
		response.OutMessage = err.Error()
		log.Error(err.Error())
	} else {
		for orderHeader.Next() {
			err = orderHeader.Scan(&response.HeaderId, &response.OutError, &response.OutMessage)
			if err != nil {
				response.OutError = util.DefaultOutErrorFailed
				response.OutMessage = err.Error()
				log.Error(err.Error())
			}
		}
	}
	return response
}

func OrderItem(request model.CreateOrder) model.OrderResponse {
	var response model.OrderResponse
	var baseResponse model.OrderItemsResp
	for _, orderItems := range request.OrderHeader.OrderItems {
		db := config.Connect(config.ReadDatabaseConfig())
		defer db.Close()

		orderItem, err := db.Query("call trx.ins_orderitem ($1,$2,$3,$4,$5,$6)",
			request.Token,
			request.OrderHeaderId,
			orderItems.ProductCode,
			orderItems.Quantity,
			request.OutError,
			request.OutMessage,
		)
		if err != nil {
			response.OutError = 1
			response.OutMessage = err.Error()
			log.Error(err.Error())
		} else {
			for orderItem.Next() {
				err = orderItem.Scan(&baseResponse.OutError, &baseResponse.OutMessage)
				if err != nil {
					response.OutError = util.DefaultOutErrorFailed
					response.OutMessage = err.Error()
					log.Error(err.Error())
				}
			}

			response.OrderItemsResponse = append(response.OrderItemsResponse, baseResponse)
			if baseResponse.OutError == 1 {
				return response
			}
		}
	}
	return response
}

func OrderCourier(request model.CreateOrder) model.OrderResponse {
	var response model.OrderResponse
	var orderCouriers model.OrderCourierResp
	db := config.Connect(config.ReadDatabaseConfig())
	defer db.Close()

	orderCourier, err := db.Query("call trx.ins_ordercourier ($1,$2,$3,$4,$5,$6,$7)",
		request.Token,
		request.OrderHeaderId,
		request.OrderHeader.OrderCourier.CourierCode,
		request.OrderHeader.OrderCourier.DeliveryType,
		request.OrderHeader.OrderCourier.DeliveryFee,
		request.OutError,
		request.OutMessage,
	)
	if err != nil {
		orderCouriers.OutError = 1
		orderCouriers.OutMessage = err.Error()
		log.Error(err.Error())
	} else {
		for orderCourier.Next() {
			err = orderCourier.Scan(&orderCouriers.OutError, &orderCouriers.OutMessage)
			if err != nil {
				orderCouriers.OutError = util.DefaultOutErrorFailed
				orderCouriers.OutMessage = err.Error()
				log.Error(err.Error())
			}
		}
		response.OrderCourierResponse = orderCouriers
	}
	return response
}

func CreateOrderDiscountVoucher(request model.OrderVoucher) model.OrderVoucherResponse {
	var response model.OrderVoucherResponse
	db := config.Connect(config.ReadDatabaseConfig())
	defer db.Close()

	orderVoucherDiscount, err := db.Query("call trx.ins_orderdiscount ($1,$2,$3,$4,$5,$6)",
		request.Token,
		request.HeaderId,
		request.VoucherCode,
		request.VoucherAmount,
		request.OutError,
		request.OutMessage,
	)
	if err != nil {
		response.OutError = 1
		response.OutMessage = err.Error()
		log.Error(err.Error())
	} else {
		for orderVoucherDiscount.Next() {
			err = orderVoucherDiscount.Scan(&response.OutError, &response.OutMessage)
			if err != nil {
				response.OutError = util.DefaultOutErrorFailed
				response.OutMessage = err.Error()
				log.Error(err.Error())
			}
		}
	}
	return response
}

func CreateOrderPayment(request model.OrderPayment) model.OrderPaymentResponse {
	var response model.OrderPaymentResponse
	db := config.Connect(config.ReadDatabaseConfig())
	defer db.Close()
	orderPayment, err := db.Query("call trx.ins_orderpayment ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11)",
		request.Token,
		request.HeaderId,
		request.PaymentMethodId,
		request.PaymentIssuerId,
		request.TotalAmount,
		request.OutError,
		request.OutMessage,
		request.AccountNumber,
		request.ExpiredBilling,
		request.BillingAmount,
		request.OrderNumber,
	)
	if err != nil {
		response.OutError = 1
		response.OutMessage = err.Error()
		log.Error(err.Error())
	} else {
		for orderPayment.Next() {
			err = orderPayment.Scan(&response.OutError, &response.OutMessage, &response.AccountNumber, &response.ExpiredBilling, &response.BillingAmount, &response.OrderNumber)
			if err != nil {
				response.OutError = util.DefaultOutErrorFailed
				response.OutMessage = err.Error()
				log.Error(err.Error())
			}
		}
	}
	return response
}

func GetDiscountAmount(request model.DiscountAmount) model.DiscountAmountResponse {
	var response model.DiscountAmountResponse
	db := config.Connect(config.ReadDatabaseConfig())
	defer db.Close()

	discountAmount, err := db.Query("select trx.get_discount_amount ($1,$2,$3)",
		request.Token,
		request.HeaderId,
		request.VoucherCode,
	)
	if err != nil {
		response.OutError = 1
		response.OutMessage = err.Error()
		log.Error(err.Error())
	} else {
		for discountAmount.Next() {
			err = discountAmount.Scan(&response.VoucherAmount)
			if err != nil {
				response.OutError = util.DefaultOutErrorFailed
				response.OutMessage = err.Error()
				log.Error(err.Error())
			}
		}
	}
	return response
}

func ListCourier(request model.ListCourier) model.ListCourierResponse {
	var response model.ListCourierResponse
	var courier model.Courier
	db := config.Connect(config.ReadDatabaseConfig())
	defer db.Close()

	listCourier, err := db.Query("select * from flow.get_list_courier ($1,$2,$3,$4)",
		request.Token,
		request.PageNumber,
		request.RowPerpage,
		request.Search,
	)
	if err != nil {
		response.OutError = 1
		response.OutMessage = err.Error()
		log.Error(err.Error())
	} else {
		for listCourier.Next() {
			err = listCourier.Scan(&courier.CourierCode, &courier.CourierName)
			if err != nil {
				response.OutError = util.DefaultOutErrorFailed
				response.OutMessage = err.Error()
				log.Error(err.Error())
			}
			response.OutError = 0
			response.OutMessage = "success"
			response.OutData = append(response.OutData, courier)
		}
	}
	return response
}

func ListCourierService(request model.ListCourierService) model.ListCourierServiceResponse {
	var response model.ListCourierServiceResponse
	var courierService model.CourierService
	db := config.Connect(config.ReadDatabaseConfig())
	defer db.Close()

	listCourierService, err := db.Query("select * from flow.get_list_courier_service ($1,$2,$3,$4,$5)",
		request.Token,
		request.CourierCode,
		request.PageNumber,
		request.RowPerpage,
		request.Search,
	)
	if err != nil {
		response.OutError = 1
		response.OutMessage = err.Error()
		log.Error(err.Error())
	} else {
		for listCourierService.Next() {
			err = listCourierService.Scan(&courierService.ServiceType, &courierService.ServiceName, &courierService.ServiceDescription)
			if err != nil {
				response.OutError = util.DefaultOutErrorFailed
				response.OutMessage = err.Error()
				log.Error(err.Error())
			}
			response.OutError = 0
			response.OutMessage = "success"
			response.OutData = append(response.OutData, courierService)
		}
	}
	return response
}

func ListPaymentMethod(request model.ListPaymentMethod) model.ListPaymentMethodResponse {
	var response model.ListPaymentMethodResponse
	var paymentMethod model.PaymentMethod
	db := config.Connect(config.ReadDatabaseConfig())
	defer db.Close()

	listPaymentMethod, err := db.Query("select * from flow.get_list_payment_method ($1,$2,$3,$4)",
		request.Token,
		request.PageNumber,
		request.RowPerpage,
		request.Search,
	)
	if err != nil {
		response.OutError = 1
		response.OutMessage = err.Error()
		log.Error(err.Error())
	} else {
		for listPaymentMethod.Next() {
			err = listPaymentMethod.Scan(&paymentMethod.PaymentHeaderCode, &paymentMethod.PaymentHeaderName, &paymentMethod.PaymentHeaderDescription, &paymentMethod.PaymentMethod)
			if err != nil {
				response.OutError = util.DefaultOutErrorFailed
				response.OutMessage = err.Error()
				log.Error(err.Error())
			}
			response.OutError = 0
			response.OutMessage = "success"
			response.OutData = append(response.OutData, paymentMethod)
		}
	}
	return response
}

func UpdateHeaderPaid(request model.UpdateHeaderPaid) model.UpdateHeaderPaidResponse {
	var response model.UpdateHeaderPaidResponse
	db := config.Connect(config.ReadDatabaseConfig())
	defer db.Close()

	updHeaderPaid, err := db.Query("call trx.upd_orderheader_paid ($1,$2,$3,$4)",
		request.Token,
		request.HeaderId,
		request.OutError,
		request.OutMessage,
	)
	if err != nil {
		response.OutError = 1
		response.OutMessage = err.Error()
		log.Error(err.Error())
	} else {
		for updHeaderPaid.Next() {
			err = updHeaderPaid.Scan(&response.OutError, &response.OutMessage)
			if err != nil {
				response.OutError = util.DefaultOutErrorFailed
				response.OutMessage = err.Error()
				log.Error(err.Error())
			}
		}
	}
	return response
}

func UpdateHeaderDo(request model.UpdateHeaderDo) model.UpdateHeaderDoResponse {
	var response model.UpdateHeaderDoResponse
	db := config.Connect(config.ReadDatabaseConfig())
	defer db.Close()

	updHeaderDo, err := db.Query("call trx.upd_orderheader_do ($1,$2,$3,$4,,$5)",
		request.Token,
		request.HeaderId,
		request.DeliveryNumber,
		request.OutError,
		request.OutMessage,
	)
	if err != nil {
		response.OutError = 1
		response.OutMessage = err.Error()
		log.Error(err.Error())
	} else {
		for updHeaderDo.Next() {
			err = updHeaderDo.Scan(&response.OutError, &response.OutMessage)
			if err != nil {
				response.OutError = util.DefaultOutErrorFailed
				response.OutMessage = err.Error()
				log.Error(err.Error())
			}
		}
	}
	return response
}
