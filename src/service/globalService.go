package service

import (
	"benings/model"
	"benings/repository"
	"encoding/json"
	"net/http"
)

func ListProvince(w http.ResponseWriter, r *http.Request) model.ListProvinceResponse {
	var response model.ListProvinceResponse
	var request model.ListProvince
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		response.OutError = 1
		response.OutMessage = "Can't bind struct"
	} else if request.PageNumber == "" {
		response.OutError = 1
		response.OutMessage = "Page Number can't be null"
	} else if request.RowPerpage == "" {
		response.OutError = 1
		response.OutMessage = "Row Perpage can't be null"
	}
	response = repository.ListProvince(request)
	return response
}

func ListCity(w http.ResponseWriter, r *http.Request) model.ListCityResponse {
	var response model.ListCityResponse
	var request model.ListCity
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		response.OutError = 1
		response.OutMessage = "Can't bind struct"
	} else if request.ProvinceCode == "" {
		response.OutError = 1
		response.OutMessage = "Province Code can't be null"
	} else if request.PageNumber == "" {
		response.OutError = 1
		response.OutMessage = "Page Number can't be null"
	} else if request.RowPerpage == "" {
		response.OutError = 1
		response.OutMessage = "Row Perpage can't be null"
	}
	response = repository.ListCity(request)
	return response
}

func ListDistrict(w http.ResponseWriter, r *http.Request) model.ListDistrictResponse {
	var response model.ListDistrictResponse
	var request model.ListDistrict
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		response.OutError = 1
		response.OutMessage = "Can't bind struct"
	} else if request.CityCode == "" {
		response.OutError = 1
		response.OutMessage = "City Code can't be null"
	} else if request.PageNumber == "" {
		response.OutError = 1
		response.OutMessage = "Page Number can't be null"
	} else if request.RowPerpage == "" {
		response.OutError = 1
		response.OutMessage = "Row Perpage can't be null"
	}
	response = repository.ListDistrict(request)
	return response
}

func ListVillage(w http.ResponseWriter, r *http.Request) model.ListVillageResponse {
	var response model.ListVillageResponse
	var request model.ListVillage
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		response.OutError = 1
		response.OutMessage = "Can't bind struct"
	} else if request.DistrictCode == "" {
		response.OutError = 1
		response.OutMessage = "DistrictCode Code can't be null"
	} else if request.PageNumber == "" {
		response.OutError = 1
		response.OutMessage = "Page Number can't be null"
	} else if request.RowPerpage == "" {
		response.OutError = 1
		response.OutMessage = "Row Perpage can't be null"
	}
	response = repository.ListVillage(request)
	return response
}

func GetMemberPath(token string) model.GetMemberPathResponse {
	var response model.GetMemberPathResponse
	response = repository.GetMemberPath(token)
	return response
}
