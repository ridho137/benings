package model

import (
	"encoding/json"
)

type Register struct {
	UserName    string `json:"userName"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	PhoneNumber string `json:"phoneNumber"`
	Gender      string `json:"gender"`
	Bod         string `json:"bod"`
	OutId       int    `json:"outId"`
	BaseResponseModel
}

type RegisterResponse struct {
	BaseResponseModel
	Otp string `json:"otp"`
}

type ValidationResponse struct {
	OutId string `json:"outId"`
	BaseResponseModel
}

type Login struct {
	UserName string `json:"userName"`
	Password string `json:"password"`
	BaseResponseModel
}

type LoginResponse struct {
	Token string `json:"token"`
	BaseResponseModel
}

type ForgotPassword struct {
	UserName             string `json:"userName"`
	Otp                  string `json:"otp"`
	OutNewRandomPassword string `json:"outNewRandomPassword"`
	OutToken             string `json:"outToken"`
	OutMemberId          int    `json:"outMemberId"`
	BaseResponseModel
}

type ForgotPasswordResponse struct {
	OutOtp               string `json:"otp"`
	OutNewRandomPassword string `json:"outNewRandomPassword"`
	OutToken             string `json:"outToken"`
	OutMemberId          int    `json:"outMemberId"`
	BaseResponseModel
}

type Logout struct {
	Token string `json:"token"`
	BaseResponseModel
}

type LogoutResponse struct {
	BaseResponseModel
}

type CheckMemberPassword struct {
	Token          string `json:"token"`
	MemberPassword string `json:"MemberPassword"`
	Email          string `json:"email"`
	BaseResponseModel
}

type CheckMemberPasswordResponse struct {
	Otp   string `json:"otp"`
	Email string `json:"email"`
	BaseResponseModel
}

type UpdateMemberEmail struct {
	Token                string `json:"token"`
	OldEmail             string `json:"oldEmail"`
	NewEmail             string `json:"newEmail"`
	VerivicationNewEmail string `json:"verivicationNewEmail"`
	BaseResponseModel
}

type UpdateMemberEmailResponse struct {
	BaseResponseModel
}

type UpdateMemberNoHp struct {
	Token               string `json:"token"`
	OldNoHp             string `json:"oldNoHp"`
	NewNoHp             string `json:"newNoHp"`
	VerivicationNewNoHp string `json:"verivicationNewNoHp"`
	BaseResponseModel
}

type UpdateMemberNoHpResponse struct {
	BaseResponseModel
}

type UpdateMemberPassword struct {
	Token                   string `json:"token"`
	OldPassword             string `json:"oldPassword"`
	NewPassword             string `json:"newPassword"`
	VerivicationNewPassword string `json:"verivicationNewPassword"`
	BaseResponseModel
}

type UpdateMemberPasswordResponse struct {
	BaseResponseModel
}

type InsertMemberAddress struct {
	Token               string `json:"token"`
	TemplateName        string `json:"templateName"`
	Street              string `json:"street"`
	ProvinceCode        string `json:"provinceCode"`
	CityCode            string `json:"cityCode"`
	DistrictCode        string `json:"districtCode"`
	VillageCode         string `json:"villageCode"`
	PostalCode          string `json:"postalCode"`
	Note                string `json:"note"`
	Latitude            string `json:"latitude"`
	ReceiverName        string `json:"receiverName"`
	ReceiverPhoneNumber string `json:"receiverPhoneNumber"`
	Longitude           string `json:"longitude"`
	OutId               int    `json:"outId"`
	BaseResponseModel
}

type InsertMemberAddressResponse struct {
	OutId string `json:"outId"`
	BaseResponseModel
}

type UpdateMemberAddress struct {
	Token               string `json:"token"`
	TemplateName        string `json:"templateName"`
	ReceiverName        string `json:"receiverName"`
	ReceiverPhoneNumber string `json:"receiverPhoneNumber"`
	Street              string `json:"street"`
	ProvinceCode        string `json:"provinceCode"`
	CityCode            string `json:"cityCode"`
	DistrictCode        string `json:"districtCode"`
	VillageCode         string `json:"villageCode"`
	PostalCode          string `json:"postalCode"`
	Note                string `json:"note"`
	Latitude            string `json:"latitude"`
	Longitude           string `json:"longitude"`
	BaseResponseModel
}

type UpdateMemberAddressResponse struct {
	BaseResponseModel
}

type UpdateMemberMainAddress struct {
	Token        string `json:"token"`
	TemplateName string `json:"templateName"`
	BaseResponseModel
}

type UpdateMemberMainAddressResponse struct {
	BaseResponseModel
}

type DeleteMemberAddress struct {
	Token        string `json:"token"`
	TemplateName string `json:"templateName"`
	BaseResponseModel
}

type DeleteMemberAddressResponse struct {
	BaseResponseModel
}

type ListMember struct {
	Token string `json:"token"`
}

type ListMemberResponse struct {
	OutData []Member `json:"outData"`
	BaseResponseModel
}

type Member struct {
	AccountId           string          `json:"accountId"`
	MemberName          string          `json:"memberName"`
	MemberEmail         string          `json:"memberEmail"`
	MemberPhone         string          `json:"memberPhone"`
	MemberGender        string          `json:"memberGender"`
	MemberBod           string          `json:"memberBod"`
	PhotoPath           string          `json:"photoPath"`
	BackgroundPhotoPath string          `json:"backgroundPhotoPath"`
	MemberAddress       json.RawMessage `json:"memberAddress"`
}

type ResponseLogin struct {
	Token       string `json:"token"`
	OutMemberId int    `json:"outMemberId"`
	BaseResponseModel
}

type UpdateMemberPhoto struct {
	Token          string `json:"token"`
	ProfilePath    string `json:"profilePath"`
	BackgroundPath string `json:"backgroundPath"`
	BaseResponseModel
}

type UpdateMemberPhotoResponse struct {
	ProfilePath    string `json:"profilePath"`
	BackgroundPath string `json:"backgroundPath"`
	BaseResponseModel
}

type MemberCart struct {
	Token string `json:"token"`
}

type MemberCartResponse struct {
	ProductId   string `json:"productId"`
	ProductCode string `json:"productCode"`
	ProductName string `json:"productName"`
	Quantity    string `json:"quantity"`
	TotalPrice  string `json:"totalPrice"`
	BaseResponseModel
}

type MemberWishList struct {
	Token string `json:"token"`
}

type MemberWishListResponse struct {
	OutData []WishList `json:"outData"`
	BaseResponseModel
}

type WishList struct {
	ProductId   string `json:"productId"`
	ProductCode string `json:"productCode"`
	ProductName string `json:"productName"`
	Price       string `json:"price"`
}

type User struct {
	UserId     string `json:"userId"`
	UserName   string `json:"userName"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	Address    string `json:"address"`
	OutError   int    `json:"outError"`
	OutMessage string `json:"outMessage"`
}

type GetMemberAddress struct {
	Token      string `json:"token"`
	PageNumber string `json:"pageNumber"`
	RowPerpage string `json:"rowPerpage"`
}

type GetMemberAddressResponse struct {
	OutData []MemberAddress `json:"outData"`
	BaseResponseModel
}

type MemberAddress struct {
	TemplateId          string `json:"templateId"`
	TemplateName        string `json:"templateName"`
	Street              string `json:"street"`
	Province            string `json:"province"`
	City                string `json:"city"`
	District            string `json:"district"`
	Village             string `json:"village"`
	PostalCode          string `json:"postalCode"`
	Note                string `json:"note"`
	Latitude            string `json:"latitude"`
	ReceiverName        string `json:"receiverName"`
	ReceiverPhoneNumber string `json:"receiverPhoneNumber"`
	Longitude           string `json:"longitude"`
	MainAddress         string `json:"mainAddress"`
}

type GetMemberVoucher struct {
	Token      string `json:"token"`
	PageNumber string `json:"pageNumber"`
	RowPerpage string `json:"rowPerpage"`
}

type GetMemberVoucherResponse struct {
	OutData []MemberVoucher `json:"outData"`
	BaseResponseModel
}

type MemberVoucher struct {
	VoucherCode        string `json:"voucherCode"`
	VoucherDescription string `json:"voucherDescription"`
	DiscountAmount     string `json:"discountAmount"`
	DiscountPercent    string `json:"discountPercent"`
	MaximumDiscount    string `json:"maximumDiscount"`
	ExpiredDate        string `json:"expiredDate"`
}

type Response struct {
	BaseResponseModel
}

type GetMemberTrx struct {
	Token       string `json:"token"`
	PageNumber  string `json:"pageNumber"`
	RowPerpage  string `json:"rowPerpage"`
	OrderStatus string `json:"orderStatus"`
}

type GetMemberTrxResponse struct {
	OutData []MemberTrx `json:"outData"`
	BaseResponseModel
}

type MemberTrx struct {
	HeaderId      string          `json:"headerId"`
	OrderNumber   string          `json:"orderNumber"`
	OrderStatus   string          `json:"orderStatus"`
	OrderDate     string          `json:"orderDate"`
	PaidDate      string          `json:"paidDate"`
	RejectDate    string          `json:"rejectDate"`
	ValidDate     string          `json:"validDate"`
	OrderItems    json.RawMessage `json:"orderItems"`
	OrderCourier  json.RawMessage `json:"orderCourier"`
	OrderDiscount json.RawMessage `json:"orderDiscount"`
	OrderPayment  json.RawMessage `json:"orderPayment"`
}
