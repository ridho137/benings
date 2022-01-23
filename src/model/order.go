package model

import "encoding/json"

type CreateOrder struct {
	Token         string `json:"token"`
	OrderType     int    `json:"orderType"`
	OrderHeaderId int    `json:"orderHeaderId"`
	OrderHeader   Order  `json:"orderHeader"`
	BaseResponseModel
}

type Order struct {
	MembertemplateId string       `json:"membertemplateId"`
	OrderItems       []OrderItems `json:"orderItems"`
	OrderCourier     OrderCourier `json:"orderCourier"`
}

type OrderItems struct {
	ProductCode string `json:"productCode"`
	Quantity    string `json:"quantity"`
}

type OrderCourier struct {
	CourierCode  string `json:"courierCode"`
	DeliveryType string `json:"deliveryType"`
	DeliveryFee  string `json:"deliveryFee"`
}

type CreateOrderResponse struct {
	HeaderId int `json:"headerId"`
	BaseResponseModel
	OrderHeaderResponse OrderResponse `json:"orderHeader"`
}

type OrderResponse struct {
	HeaderId int `json:"headerId"`
	BaseResponseModel
	OrderItemsResponse   []OrderItemsResp `json:"orderItems"`
	OrderCourierResponse OrderCourierResp `json:"orderCourier"`
}

type OrderItemsResp struct {
	BaseResponseModel
}

type OrderCourierResp struct {
	BaseResponseModel
}

type OrderVoucher struct {
	Token         string `json:"token"`
	HeaderId      string `json:"headerId"`
	VoucherCode   string `json:"voucherCode"`
	VoucherAmount string `json:"voucherAmount"`
	BaseResponseModel
}

type OrderVoucherResponse struct {
	BaseResponseModel
}

type OrderPayment struct {
	Token           string `json:"token"`
	HeaderId        string `json:"headerId"`
	PaymentMethodId string `json:"paymentMethodId"`
	PaymentIssuerId string `json:"paymentIssuerId"`
	TotalAmount     string `json:"totalAmount"`
	AccountNumber   string `json:"accountNumber"`
	ExpiredBilling  string `json:"expiredBilling"`
	BillingAmount   int    `json:"billingAmount"`
	OrderNumber     string `json:"orderNumber"`
	BaseResponseModel
}

type OrderPaymentResponse struct {
	AccountNumber  string `json:"accountNumber"`
	ExpiredBilling string `json:"expiredBilling"`
	BillingAmount  int    `json:"billingAmount"`
	OrderNumber    string `json:"orderNumber"`
	BaseResponseModel
}

type DiscountAmount struct {
	Token       string `json:"token"`
	HeaderId    string `json:"HeaderId"`
	VoucherCode string `json:"voucherCode"`
}

type DiscountAmountResponse struct {
	VoucherAmount string `json:"voucherAmount"`
	BaseResponseModel
}

type ListCourier struct {
	Token      string `json:"token"`
	PageNumber string `json:"pageNumber"`
	RowPerpage string `json:"rowPerpage"`
	Search     string `json:"search"`
}

type ListCourierResponse struct {
	OutData []Courier `json:"outData"`
	BaseResponseModel
}

type Courier struct {
	CourierCode string `json:"courierCode"`
	CourierName string `json:"courierName"`
}

type ListCourierService struct {
	Token       string `json:"token"`
	CourierCode string `json:"courierCode"`
	PageNumber  string `json:"pageNumber"`
	RowPerpage  string `json:"rowPerpage"`
	Search      string `json:"search"`
}

type ListCourierServiceResponse struct {
	OutData []CourierService `json:"outData"`
	BaseResponseModel
}

type CourierService struct {
	ServiceType        string `json:"serviceType"`
	ServiceName        string `json:"serviceName"`
	ServiceDescription string `json:"serviceDescription"`
}

type ListPaymentMethod struct {
	Token      string `json:"token"`
	PageNumber string `json:"pageNumber"`
	RowPerpage string `json:"rowPerpage"`
	Search     string `json:"search"`
}

type ListPaymentMethodResponse struct {
	OutData []PaymentMethod `json:"outData"`
	BaseResponseModel
}

type PaymentMethod struct {
	PaymentHeaderCode        string          `json:"paymentHeaderCode"`
	PaymentHeaderName        string          `json:"paymentHeaderName"`
	PaymentHeaderDescription string          `json:"paymentHeaderDescription"`
	PaymentMethod            json.RawMessage `json:"paymentMethod"`
}

type UpdateHeaderPaid struct {
	Token    string `json:"token"`
	HeaderId string `json:"headerId"`
	BaseResponseModel
}

type UpdateHeaderPaidResponse struct {
	BaseResponseModel
}

type UpdateHeaderDo struct {
	Token          string `json:"token"`
	HeaderId       string `json:"headerId"`
	DeliveryNumber string `json:"deliveryNumber"`
	BaseResponseModel
}

type UpdateHeaderDoResponse struct {
	BaseResponseModel
}
