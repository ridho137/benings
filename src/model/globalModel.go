package model

import (
	"encoding/json"

	"github.com/golang-jwt/jwt"
)

type BaseResponseModel struct {
	OutError   int    `json:"outError"`
	OutMessage string `json:"outMessage"`
}

type Token struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

type JwtResponse struct {
	OutError   int    `json:"outError"`
	OutMessage string `json:"outMessage"`
	Token      string `json:"token"`
}

type Info struct {
	UserName    string
	Otp         string
	NewPassword string
}

type ListProvince struct {
	PageNumber string `json:"pageNumber"`
	RowPerpage string `json:"rowPerpage"`
	Search     string `json:"search"`
}

type ListProvinceResponse struct {
	OutData []Province `json:"outData"`
	BaseResponseModel
}

type Province struct {
	ProvinceCode string `json:"provinceCode"`
	ProvinceName string `json:"provinceName"`
}

type ListCity struct {
	ProvinceCode string `json:"provinceCode"`
	PageNumber   string `json:"pageNumber"`
	RowPerpage   string `json:"rowPerpage"`
	Search       string `json:"search"`
}

type ListCityResponse struct {
	OutData []City `json:"outData"`
	BaseResponseModel
}

type City struct {
	CityCode string `json:"cityCode"`
	CityName string `json:"cityName"`
}

type ListDistrict struct {
	CityCode   string `json:"cityCode"`
	PageNumber string `json:"pageNumber"`
	RowPerpage string `json:"rowPerpage"`
	Search     string `json:"search"`
}

type ListDistrictResponse struct {
	OutData []District `json:"outData"`
	BaseResponseModel
}

type District struct {
	DistrictCode string `json:"districtCode"`
	DistrictName string `json:"districtName"`
}

type ListVillage struct {
	DistrictCode string `json:"districtCode"`
	PageNumber   string `json:"pageNumber"`
	RowPerpage   string `json:"rowPerpage"`
	Search       string `json:"search"`
}

type ListVillageResponse struct {
	OutData []Village `json:"outData"`
	BaseResponseModel
}

type Village struct {
	VillageCode string `json:"villageCode"`
	VillageName string `json:"villageName"`
	PostalCode  string `json:"postalCode"`
}

type GetMemberPath struct {
	Token         string `json:"token"`
	GlobalPath    string `json:"globalPath"`
	PublicPath    string `json:"publicPath"`
	AccountId     string `json:"accountId"`
	Key           string `json:"key"`
	InitialVector string `json:"initialVector"`
	BaseResponseModel
}

type GetMemberPathResponse struct {
	GlobalPath    string `json:"globalPath"`
	PublicPath    string `json:"publicPath"`
	AccountId     string `json:"accountId"`
	Key           string `json:"key"`
	InitialVector string `json:"initialVector"`
	BaseResponseModel
}

type LogInsOrder struct {
	Token   string          `json:"token"`
	Request json.RawMessage `json:"request"`
	OutId   string          `json:"outId"`
	BaseResponseModel
}

type LogInsOrderResponse struct {
	OutId string `json:"outId"`
	BaseResponseModel
}

type LogUpdOrder struct {
	LogId    string          `json:"logId"`
	Response json.RawMessage `json:"response"`
	BaseResponseModel
}

type LogUpdOrderResponse struct {
	BaseResponseModel
}
