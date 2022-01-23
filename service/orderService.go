package service

import (
	"benings/model"
	"benings/repository"
	"encoding/json"
	"log"
	"net/http"
)

func CreateOrder(w http.ResponseWriter, r *http.Request) model.CreateOrderResponse {
	var response model.CreateOrderResponse
	var responseHeader model.CreateOrderResponse
	var responseOrderItem model.OrderResponse
	var responseCourier model.OrderResponse
	var request model.CreateOrder
	var logInsOrder model.LogInsOrder
	var logInsOrderResponse model.LogInsOrderResponse
	var logUpdOrder model.LogUpdOrder
	APIkey := r.Header.Get("API-Key")
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		response.OutError = 1
		response.OutMessage = "File can't be null"
		return response
	} else if APIkey == "" {
		response.OutError = 1
		response.OutMessage = "API Key can't be null"
		return response
	} else if request.OrderHeader.MembertemplateId == "" {
		response.OutError = 1
		response.OutMessage = "Member template Id can't be null"
		return response
	} else if request.OrderHeader.OrderItems[0].ProductCode == "" {
		response.OutError = 1
		response.OutMessage = "Product code can't be null"
		return response
	} else if request.OrderHeader.OrderItems[0].Quantity == "" {
		response.OutError = 1
		response.OutMessage = "Quantity can't be null"
		return response
	} else if request.OrderHeader.OrderCourier.CourierCode == "" {
		response.OutError = 1
		response.OutMessage = "Courier code can't be null"
		return response
	} else if request.OrderHeader.OrderCourier.DeliveryType == "" {
		response.OutError = 1
		response.OutMessage = "Delivery type can't be null"
		return response
	} else if request.OrderHeader.OrderCourier.DeliveryFee == "" {
		response.OutError = 1
		response.OutMessage = "Delivery fee can't be null"
		return response
	}
	request.OrderType = 10
	request.Token = APIkey
	log.Println("orderType : ", request.OrderType)
	log.Println("MembertemplateId : ", request.OrderHeader.MembertemplateId)
	response.OrderHeaderResponse.OrderCourierResponse.OutError = 1
	response.OrderHeaderResponse.OrderCourierResponse.OutMessage = "Failed"
	if response.OutError != 1 {
		req, _ := json.MarshalIndent(&request, "", "  ")
		logInsOrder.Token = APIkey
		logInsOrder.Request = req
		logInsOrderResponse = repository.LogInsOrder(logInsOrder)
		//hit header
		responseHeader = repository.OrderHeader(request)
		response.OrderHeaderResponse.HeaderId = responseHeader.HeaderId
		response.OrderHeaderResponse.OutError = responseHeader.OutError
		response.OrderHeaderResponse.OutMessage = responseHeader.OutMessage
		response.HeaderId = responseHeader.HeaderId
		response.OutError = responseHeader.OutError
		response.OutMessage = responseHeader.OutMessage
	}
	request.OrderHeaderId = responseHeader.HeaderId
	if response.OutError != 1 {
		log.Println("getOrderItem")
		responseOrderItem = repository.OrderItem(request)
		response.OrderHeaderResponse.OrderItemsResponse = responseOrderItem.OrderItemsResponse
		for _, orderItems := range responseOrderItem.OrderItemsResponse {
			response.OutError = orderItems.OutError
			response.OutMessage = orderItems.OutMessage
		}
	}
	if response.OutError != 1 {
		log.Println("getOrderCourier")
		responseCourier = repository.OrderCourier(request)
		response.OrderHeaderResponse.OrderCourierResponse.OutError = responseCourier.OrderCourierResponse.OutError
		response.OrderHeaderResponse.OrderCourierResponse.OutMessage = responseCourier.OrderCourierResponse.OutMessage
		response.OutError = responseCourier.OrderCourierResponse.OutError
		response.OutMessage = responseCourier.OrderCourierResponse.OutMessage
	}
	resp, _ := json.MarshalIndent(&response, "", "  ")
	logUpdOrder.LogId = logInsOrderResponse.OutId
	logUpdOrder.Response = resp
	repository.LogUpdOrder(logUpdOrder)
	return response
}

func CreateOrderDiscountVoucher(w http.ResponseWriter, r *http.Request) model.OrderVoucherResponse {
	var response model.OrderVoucherResponse
	var request model.OrderVoucher
	APIkey := r.Header.Get("API-Key")
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		response.OutError = 1
		response.OutMessage = "File can't be null"
	} else if APIkey == "" {
		response.OutError = 1
		response.OutMessage = "API Key can't be null"
	} else if request.HeaderId == "" {
		response.OutError = 1
		response.OutMessage = "Header Id can't be null"
	} else if request.VoucherCode == "" {
		response.OutError = 1
		response.OutMessage = "Voucher code can't be null"
	} else if request.VoucherAmount == "" {
		response.OutError = 1
		response.OutMessage = "Voucher amount can't be null"
	}
	request.Token = APIkey
	response = repository.CreateOrderDiscountVoucher(request)
	return response
}

func CreateOrderPayment(w http.ResponseWriter, r *http.Request) model.OrderPaymentResponse {
	var response model.OrderPaymentResponse
	var request model.OrderPayment
	APIkey := r.Header.Get("API-Key")
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		response.OutError = 1
		response.OutMessage = "File can't be null"
	} else if APIkey == "" {
		response.OutError = 1
		response.OutMessage = "API Key can't be null"
	} else if request.HeaderId == "" {
		response.OutError = 1
		response.OutMessage = "Header Id can't be null"
	} else if request.PaymentMethodId == "" {
		response.OutError = 1
		response.OutMessage = "Payment method id can't be null"
	} else if request.PaymentIssuerId == "" {
		response.OutError = 1
		response.OutMessage = "Payment issuer id can't be null"
	} else if request.TotalAmount == "" {
		response.OutError = 1
		response.OutMessage = "Total amount can't be null"
	}
	request.Token = APIkey
	response = repository.CreateOrderPayment(request)
	return response
}

func GetDiscountAmount(w http.ResponseWriter, r *http.Request) model.DiscountAmountResponse {
	var response model.DiscountAmountResponse
	var request model.DiscountAmount
	APIkey := r.Header.Get("API-Key")
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		response.OutError = 1
		response.OutMessage = "File can't be null"
	} else if APIkey == "" {
		response.OutError = 1
		response.OutMessage = "API Key can't be null"
	} else if request.HeaderId == "" {
		response.OutError = 1
		response.OutMessage = "Order header Id can't be null"
	} else if request.VoucherCode == "" {
		response.OutError = 1
		response.OutMessage = "Voucher Code code can't be null"
	}
	request.Token = APIkey
	response = repository.GetDiscountAmount(request)
	return response
}

func ListCourier(w http.ResponseWriter, r *http.Request) model.ListCourierResponse {
	var response model.ListCourierResponse
	var request model.ListCourier
	APIkey := r.Header.Get("API-Key")
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		response.OutError = 1
		response.OutMessage = "File can't be null"
	} else if APIkey == "" {
		response.OutError = 1
		response.OutMessage = "API Key can't be null"
	}
	request.Token = APIkey
	response = repository.ListCourier(request)
	return response
}

func ListCourierService(w http.ResponseWriter, r *http.Request) model.ListCourierServiceResponse {
	var response model.ListCourierServiceResponse
	var request model.ListCourierService
	APIkey := r.Header.Get("API-Key")
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		response.OutError = 1
		response.OutMessage = "File can't be null"
	} else if APIkey == "" {
		response.OutError = 1
		response.OutMessage = "API Key can't be null"
	} else if request.CourierCode == "" {
		response.OutError = 1
		response.OutMessage = "Courier Code can't be null"
	}
	request.Token = APIkey
	response = repository.ListCourierService(request)
	return response
}

func ListPaymentMethod(w http.ResponseWriter, r *http.Request) model.ListPaymentMethodResponse {
	var response model.ListPaymentMethodResponse
	var request model.ListPaymentMethod
	APIkey := r.Header.Get("API-Key")
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		response.OutError = 1
		response.OutMessage = "File can't be null"
	} else if APIkey == "" {
		response.OutError = 1
		response.OutMessage = "API Key can't be null"
	}
	request.Token = APIkey
	response = repository.ListPaymentMethod(request)
	return response
}

func UpdateHeaderPaid(w http.ResponseWriter, r *http.Request) model.UpdateHeaderPaidResponse {
	var response model.UpdateHeaderPaidResponse
	var request model.UpdateHeaderPaid
	APIkey := r.Header.Get("API-Key")
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		response.OutError = 1
		response.OutMessage = "File can't be null"
	} else if APIkey == "" {
		response.OutError = 1
		response.OutMessage = "API Key can't be null"
	} else if request.HeaderId == "" {
		response.OutError = 1
		response.OutMessage = "Order header Id can't be null"
	}
	request.Token = APIkey
	response = repository.UpdateHeaderPaid(request)
	return response
}

func UpdateHeaderDo(w http.ResponseWriter, r *http.Request) model.UpdateHeaderDoResponse {
	var response model.UpdateHeaderDoResponse
	var request model.UpdateHeaderDo
	APIkey := r.Header.Get("API-Key")
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		response.OutError = 1
		response.OutMessage = "File can't be null"
	} else if APIkey == "" {
		response.OutError = 1
		response.OutMessage = "API Key can't be null"
	} else if request.HeaderId == "" {
		response.OutError = 1
		response.OutMessage = "Order header Id can't be null"
	} else if request.DeliveryNumber == "" {
		response.OutError = 1
		response.OutMessage = "Delivery number can't be null"
	}
	request.Token = APIkey
	response = repository.UpdateHeaderDo(request)
	return response
}
