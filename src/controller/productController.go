package controller

import (
	"benings/model"
	"benings/service"
	"encoding/json"
	"net/http"
)

func ListBanner(w http.ResponseWriter, r *http.Request) {
	var response model.ListBannerResponse
	response = service.GetListBanner(w, r)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func ListProduct(w http.ResponseWriter, r *http.Request) {
	var response model.ListProductResponse
	response = service.GetListProduct(w, r)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func AddCart(w http.ResponseWriter, r *http.Request) {
	var response model.AddCartResponse
	response = service.AddCart(w, r)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func AddCartByQuantity(w http.ResponseWriter, r *http.Request) {
	var response model.AddCartByQuantityResponse
	response = service.AddCartByQuantity(w, r)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func DeleteCart(w http.ResponseWriter, r *http.Request) {
	var response model.DeleteCartResponse
	response = service.DeleteCart(w, r)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func AddWishList(w http.ResponseWriter, r *http.Request) {
	var response model.AddWishListResponse
	response = service.AddWishList(w, r)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func DeleteWishList(w http.ResponseWriter, r *http.Request) {
	var response model.DeleteWishListResponse
	response = service.DeleteWishList(w, r)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func ListCategory(w http.ResponseWriter, r *http.Request) {
	var response model.ListCategoryResponse
	response = service.ListCategory(w, r)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func ListType(w http.ResponseWriter, r *http.Request) {
	var response model.ListTypeResponse
	response = service.GetListType(w, r)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// func ListCategoryProduct(w http.ResponseWriter, r *http.Request) {
// 	var response []model.Product
// 	response = service.GetCategoryProduct(w, r)
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(response)
// }
