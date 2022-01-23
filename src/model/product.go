package model

import "encoding/json"

type ListBanner struct {
	Token      string `json:"token"`
	PageNumber string `json:"pageNumber"`
	RowPerpage string `json:"rowPerpage"`
}

type ListBannerResponse struct {
	OutData []Banner `json:"outData"`
	BaseResponseModel
}

type Banner struct {
	BannerCode    string `json:"bannerCode"`
	BannerName    string `json:"bannerName"`
	BannerPath    string `json:"bannerPath"`
	EffectiveDate string `json:"effectiveDate"`
	ExpiredDate   string `json:"expiredDate"`
}

type ListProduct struct {
	Token        string `json:"token"`
	BannerCode   string `json:"bannerCode"`
	CategoryCode string `json:"categoryCode"`
	TypeId       string `json:"typeId"`
	PageNumber   string `json:"pageNumber"`
	RowPerpage   string `json:"rowPerpage"`
	Search       string `json:"search"`
}

type ListProductResponse struct {
	OutData []Product `json:"outData"`
	BaseResponseModel
}

type Product struct {
	ProductCode      string          `json:"productCode"`
	ProductName      string          `json:"productName"`
	ProductPath      string          `json:"productPath"`
	ProductPrice     string          `json:"productPrice"`
	PromoPrice       string          `json:"promoPrice"`
	IsPromo          string          `json:"isPromo"`
	Netto            string          `json:"netto"`
	Ratting          string          `json:"ratting"`
	SoldByMember     string          `json:"soldByMember"`
	TypeDescription  string          `json:"typeDescription"`
	Categories       json.RawMessage `json:"categories"`
	ProductAttribute json.RawMessage `json:"productAttribute"`
	IsWishlist       string          `json:"isWishlist"`
	ProductWeight    string          `json:"productWeight"`
}

type ListCategory struct {
	Token      string `json:"token"`
	PageNumber string `json:"pageNumber"`
	RowPerpage string `json:"rowPerpage"`
	Search     string `json:"search"`
}

type ListCategoryResponse struct {
	OutData []Category `json:"outData"`
	BaseResponseModel
}

type Category struct { 
	CategoryCode        string `json:"categoryCode"`
	CategoryName        string `json:"categoryName"`
	CategoryPath        string `json:"categoryPath"`
	CategoryDescription string `json:"categoryDescription"`
}

type AddCart struct {
	Token       string `json:"token"`
	ProductCode string `json:"productCode"`
	BaseResponseModel
}

type AddCartResponse struct {
	BaseResponseModel
}

type AddCartByQuantity struct {
	Token       string `json:"token"`
	ProductCode string `json:"productCode"`
	Quantity    string `json:"quantity"`
	BaseResponseModel
}

type AddCartByQuantityResponse struct {
	BaseResponseModel
}

type DeleteCart struct {
	Token       string `json:"token"`
	ProductCode string `json:"productCode"`
	BaseResponseModel
}

type DeleteCartResponse struct {
	BaseResponseModel
}

type AddWishList struct {
	Token       string `json:"token"`
	ProductCode string `json:"productCode"`
	BaseResponseModel
}

type AddWishListResponse struct {
	BaseResponseModel
}

type DeleteWishList struct {
	Token       string `json:"token"`
	ProductCode string `json:"productCode"`
	BaseResponseModel
}

type DeleteWishListResponse struct {
	BaseResponseModel
}

type ListType struct {
	Token      string `json:"token"`
	PageNumber string `json:"pageNumber"`
	RowPerpage string `json:"rowPerpage"`
	Search     string `json:"search"`
}

type ListTypeResponse struct {
	OutData []Type `json:"outData"`
	BaseResponseModel
}

type Type struct {
	TypeId string `json:"tpyeId"`
	Type   string `json:"type"`
}