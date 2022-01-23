package controller

import (
	"benings/model"
	"benings/service"
	"encoding/json"
	"net/http"
)

func Register(w http.ResponseWriter, r *http.Request) {
	var response model.RegisterResponse
	response = service.Register(w, r)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func ValidationOTP(w http.ResponseWriter, r *http.Request) {
	var response model.ValidationResponse
	response = service.ValidationOTP(w, r)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func Login(w http.ResponseWriter, r *http.Request) {
	var response model.LoginResponse
	response = service.Login(w, r)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func Logout(w http.ResponseWriter, r *http.Request) {
	var response model.LogoutResponse
	response = service.Logout(w, r)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func ForgotPassword(w http.ResponseWriter, r *http.Request) {
	var response model.ForgotPasswordResponse
	response = service.ForgotPassword(w, r)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func CheckMemberPassword(w http.ResponseWriter, r *http.Request) {
	var response model.CheckMemberPasswordResponse
	response = service.CheckMemberPassword(w, r)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func UpdateMemberEmail(w http.ResponseWriter, r *http.Request) {
	var response model.UpdateMemberEmailResponse
	response = service.UpdateMemberEmail(w, r)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func UpdateMemberNoHp(w http.ResponseWriter, r *http.Request) {
	var response model.UpdateMemberNoHpResponse
	response = service.UpdateMemberNoHp(w, r)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func UpdateMemberPassword(w http.ResponseWriter, r *http.Request) {
	var response model.UpdateMemberPasswordResponse
	response = service.UpdateMemberPassword(w, r)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func ListMember(w http.ResponseWriter, r *http.Request) {
	var response model.ListMemberResponse
	response = service.ListMember(w, r)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func InsertMemberAddress(w http.ResponseWriter, r *http.Request) {
	var response model.InsertMemberAddressResponse
	response = service.InsertMemberAddress(w, r)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func UpdateMemberAddress(w http.ResponseWriter, r *http.Request) {
	var response model.UpdateMemberAddressResponse
	response = service.UpdateMemberAddress(w, r)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func DeleteMemberAddress(w http.ResponseWriter, r *http.Request) {
	var response model.DeleteMemberAddressResponse
	response = service.DeleteMemberAddress(w, r)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func UpdateMemberMainAddress(w http.ResponseWriter, r *http.Request) {
	var response model.UpdateMemberMainAddressResponse
	response = service.UpdateMemberMainAddress(w, r)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func UpdateMemberPhotoProfile(w http.ResponseWriter, r *http.Request) {
	var response model.UpdateMemberPhotoResponse
	response = service.UpdateMemberPhotoProfile(w, r)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func UpdateMemberPhotoBackground(w http.ResponseWriter, r *http.Request) {
	var response model.UpdateMemberPhotoResponse
	response = service.UpdateMemberPhotoBackground(w, r)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func GetMemberCart(w http.ResponseWriter, r *http.Request) {
	var response model.MemberCartResponse
	response = service.GetMemberCart(w, r)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func GetMemberWishList(w http.ResponseWriter, r *http.Request) {
	var response model.MemberWishListResponse
	response = service.GetMemberWishList(w, r)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func GetMemberAddress(w http.ResponseWriter, r *http.Request) {
	var response model.GetMemberAddressResponse
	response = service.GetMemberAddress(w, r)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func GetMemberVoucher(w http.ResponseWriter, r *http.Request) {
	var response model.GetMemberVoucherResponse
	response = service.GetMemberVoucher(w, r)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func GetMemberTransaction(w http.ResponseWriter, r *http.Request) {
	var response model.GetMemberTrxResponse
	response = service.GetMemberTransaction(w, r)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
