package controller

import (
	"benings/model"
	"benings/service"
	"encoding/json"
	"net/http"
)

func CreateOrder(w http.ResponseWriter, r *http.Request) {
	var response model.CreateOrderResponse
	response = service.CreateOrder(w, r)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func CreateOrderDiscountVoucher(w http.ResponseWriter, r *http.Request) {
	var response model.OrderVoucherResponse
	response = service.CreateOrderDiscountVoucher(w, r)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func CreateOrderPayment(w http.ResponseWriter, r *http.Request) {
	var response model.OrderPaymentResponse
	response = service.CreateOrderPayment(w, r)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func GetDiscountAmount(w http.ResponseWriter, r *http.Request) {
	var response model.DiscountAmountResponse
	response = service.GetDiscountAmount(w, r)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func ListCourier(w http.ResponseWriter, r *http.Request) {
	var response model.ListCourierResponse
	response = service.ListCourier(w, r)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func ListCourierService(w http.ResponseWriter, r *http.Request) {
	var response model.ListCourierServiceResponse
	response = service.ListCourierService(w, r)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func ListPaymentMethod(w http.ResponseWriter, r *http.Request) {
	var response model.ListPaymentMethodResponse
	response = service.ListPaymentMethod(w, r)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func UpdateHeaderPaid(w http.ResponseWriter, r *http.Request) {
	var response model.UpdateHeaderPaidResponse
	response = service.UpdateHeaderPaid(w, r)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func UpdateHeaderDo(w http.ResponseWriter, r *http.Request) {
	var response model.UpdateHeaderDoResponse
	response = service.UpdateHeaderDo(w, r)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
