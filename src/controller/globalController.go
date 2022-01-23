package controller

import (
	"benings/model"
	"benings/service"
	"encoding/json"
	"net/http"
)

func ListProvince(w http.ResponseWriter, r *http.Request) {
	var response model.ListProvinceResponse
	response = service.ListProvince(w, r)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func ListCity(w http.ResponseWriter, r *http.Request) {
	var response model.ListCityResponse
	response = service.ListCity(w, r)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func ListDistrict(w http.ResponseWriter, r *http.Request) {
	var response model.ListDistrictResponse
	response = service.ListDistrict(w, r)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func ListVillage(w http.ResponseWriter, r *http.Request) {
	var response model.ListVillageResponse
	response = service.ListVillage(w, r)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
