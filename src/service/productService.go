package service

import (
	"benings/model"
	"benings/repository"
	"encoding/json"
	"net/http"
)

func GetListBanner(w http.ResponseWriter, r *http.Request) model.ListBannerResponse {
	var response model.ListBannerResponse
	var request model.ListBanner
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
	response = repository.GetListBanner(request)
	return response
}

func GetListProduct(w http.ResponseWriter, r *http.Request) model.ListProductResponse {
	var response model.ListProductResponse
	var request model.ListProduct
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
	response = repository.GetListProduct(request)
	return response
}

func AddCart(w http.ResponseWriter, r *http.Request) model.AddCartResponse {
	var response model.AddCartResponse
	var request model.AddCart
	APIkey := r.Header.Get("API-Key")
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		response.OutError = 1
		response.OutMessage = "File can't be null"
	} else if APIkey == "" {
		response.OutError = 1
		response.OutMessage = "API Key can't be null"
	} else if request.ProductCode == "" {
		response.OutError = 1
		response.OutMessage = "Product Code can't be null"
	}
	request.Token = APIkey
	response = repository.AddCart(request)
	return response
}

func AddCartByQuantity(w http.ResponseWriter, r *http.Request) model.AddCartByQuantityResponse {
	var response model.AddCartByQuantityResponse
	var request model.AddCartByQuantity
	APIkey := r.Header.Get("API-Key")
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		response.OutError = 1
		response.OutMessage = "File can't be null"
	} else if APIkey == "" {
		response.OutError = 1
		response.OutMessage = "API Key can't be null"
	} else if request.ProductCode == "" {
		response.OutError = 1
		response.OutMessage = "Product Code can't be null"
	} else if request.Quantity == "" {
		response.OutError = 1
		response.OutMessage = "Quantity can't be null"
	}
	request.Token = APIkey
	response = repository.AddCartByQuantity(request)
	return response
}

func DeleteCart(w http.ResponseWriter, r *http.Request) model.DeleteCartResponse {
	var response model.DeleteCartResponse
	var request model.DeleteCart
	APIkey := r.Header.Get("API-Key")
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		response.OutError = 1
		response.OutMessage = "File can't be null"
	} else if APIkey == "" {
		response.OutError = 1
		response.OutMessage = "API Key can't be null"
	} else if request.ProductCode == "" {
		response.OutError = 1
		response.OutMessage = "Product Code can't be null"
	}
	request.Token = APIkey
	response = repository.DeleteCart(request)
	return response
}

func AddWishList(w http.ResponseWriter, r *http.Request) model.AddWishListResponse {
	var response model.AddWishListResponse
	var request model.AddWishList
	APIkey := r.Header.Get("API-Key")
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		response.OutError = 1
		response.OutMessage = "File can't be null"
	} else if APIkey == "" {
		response.OutError = 1
		response.OutMessage = "API Key can't be null"
	} else if request.ProductCode == "" {
		response.OutError = 1
		response.OutMessage = "Product Code can't be null"
	}
	request.Token = APIkey
	response = repository.AddWishList(request)
	return response
}

func DeleteWishList(w http.ResponseWriter, r *http.Request) model.DeleteWishListResponse {
	var response model.DeleteWishListResponse
	var request model.DeleteWishList
	APIkey := r.Header.Get("API-Key")
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		response.OutError = 1
		response.OutMessage = "File can't be null"
	} else if APIkey == "" {
		response.OutError = 1
		response.OutMessage = "API Key can't be null"
	} else if request.ProductCode == "" {
		response.OutError = 1
		response.OutMessage = "Product Code can't be null"
	}
	request.Token = APIkey
	response = repository.DeleteWishList(request)
	return response
}

func ListCategory(w http.ResponseWriter, r *http.Request) model.ListCategoryResponse {
	var response model.ListCategoryResponse
	var request model.ListCategory
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
	response = repository.ListCategory(request)
	return response
}

func GetListType(w http.ResponseWriter, r *http.Request) model.ListTypeResponse {
	var response model.ListTypeResponse
	var request model.ListType
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
	response = repository.GetListType(request)
	return response
}

// func GetCategoryProduct(w http.ResponseWriter, r *http.Request) []model.Product {
// 	var arr_response []model.Product = make([]model.Product, 0)
// 	var response model.Response
// 	var request model.ProductByCategoryRequest
// 	APIkey := r.Header.Get("API-Key")
// 	err := json.NewDecoder(r.Body).Decode(&request)
// 	if err != nil {
// 		response.OutError = 1
// 		response.OutMessage = "Can't bind struct"
// 	} else if APIkey == "" {
// 		response.OutError = 1
// 		response.OutMessage = "API Key can't be null"
// 	} else if request.Limit == "" {
// 		response.OutError = 1
// 		response.OutMessage = "Limit can't be 0"
// 	} else if request.Offset == "" {
// 		response.OutError = 1
// 		response.OutMessage = "Offset can't be null"
// 	}
// 	log.Println("user: ", request.UserId)
// 	if response.OutError == 0 {
// 		response = repository.APIkey(APIkey, request.UserId)
// 	}
// 	if response.OutError == 0 {
// 		arr_response = repository.GetProductByCategory(request)
// 	} else {
// 		var respProductPromo model.Product
// 		respProductPromo.OutError = response.OutError
// 		respProductPromo.OutMessage = response.OutMessage
// 		arr_response = append(arr_response, respProductPromo)
// 	}
// 	return arr_response
// }
