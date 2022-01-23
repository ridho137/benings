package service

import (
	"benings/model"
	"benings/repository"
	"benings/util"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"path/filepath"
)

func Register(w http.ResponseWriter, r *http.Request) model.RegisterResponse {
	var response model.RegisterResponse
	var request model.Register
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		response.OutError = 1
		response.OutMessage = "Can't bind struct"
		return response
	} else if !util.ValidateEmail(request.Email) {
		response.OutError = 1
		response.OutMessage = "Email address is invalid"
		return response
	}
	otp := util.GenerateOtp(6)
	info := model.Info{UserName: request.UserName, Otp: otp}
	subject := "Register User"
	util.EmailSender(w, request.Email, subject, info)
	response.OutError = 0
	response.OutMessage = "send email success"
	response.Otp = otp
	return response
}

func ValidationOTP(w http.ResponseWriter, r *http.Request) model.ValidationResponse {
	var response model.ValidationResponse
	var request model.Register
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		response.OutError = 1
		response.OutMessage = "Can't bind struct"
		return response
	} else if !util.ValidateEmail(request.Email) {
		response.OutError = 1
		response.OutMessage = "Email address is invalid"
		return response
	}
	response = repository.ValidationOtp(request)
	return response
}

func Login(w http.ResponseWriter, r *http.Request) model.LoginResponse {
	var response model.LoginResponse
	var request model.Login
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		response.OutError = 1
		response.OutMessage = "Can't bind struct"
		return response
	}
	response = repository.Login(request)
	return response
}

func Logout(w http.ResponseWriter, r *http.Request) model.LogoutResponse {
	var response model.LogoutResponse
	var request model.Logout
	APIkey := r.Header.Get("API-Key")
	if APIkey == "" {
		response.OutError = 1
		response.OutMessage = "API Key can't be null"
		return response
	}
	request.Token = APIkey
	response = repository.Logout(request)
	return response
}

func ForgotPassword(w http.ResponseWriter, r *http.Request) model.ForgotPasswordResponse {
	var response model.ForgotPasswordResponse
	var request model.ForgotPassword
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		response.OutError = 1
		response.OutMessage = "Can't bind struct"
		return response
	} else if request.UserName == "" {
		response.OutError = 1
		response.OutMessage = "User Name can't be null"
		return response
	}
	response = repository.ForgotPassword(request)
	if response.OutError == 0 {
		info := model.Info{UserName: request.UserName, Otp: "OTP : " + response.OutOtp, NewPassword: "Password : " + response.OutNewRandomPassword}
		subject := "Forgot Password"
		util.EmailSender(w, request.UserName, subject, info)
	}
	return response
}

func CheckMemberPassword(w http.ResponseWriter, r *http.Request) model.CheckMemberPasswordResponse {
	var response model.CheckMemberPasswordResponse
	var request model.CheckMemberPassword
	APIkey := r.Header.Get("API-Key")
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		response.OutError = 1
		response.OutMessage = "Can't bind struct"
		return response
	} else if APIkey == "" {
		response.OutError = 1
		response.OutMessage = "API Key can't be null"
		return response
	} else if request.MemberPassword == "" {
		response.OutError = 1
		response.OutMessage = "Password can't be null"
		return response
	}
	request.Token = APIkey
	response = repository.CheckMemberPassword(request)
	if response.OutError == 0 {
		otp := util.GenerateOtp(6)
		info := model.Info{UserName: response.Email, Otp: otp}
		subject := "Confirmation Change Member Prefence"
		util.EmailSender(w, response.Email, subject, info)
		response.Otp = otp
	}
	return response
}

func UpdateMemberEmail(w http.ResponseWriter, r *http.Request) model.UpdateMemberEmailResponse {
	var response model.UpdateMemberEmailResponse
	var request model.UpdateMemberEmail
	APIkey := r.Header.Get("API-Key")
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		response.OutError = 1
		response.OutMessage = "Can't bind struct"
	} else if APIkey == "" {
		response.OutError = 1
		response.OutMessage = "API Key can't be null"
	} else if request.OldEmail == "" {
		response.OutError = 1
		response.OutMessage = "Old Email can't be null"
		return response
	} else if request.NewEmail == "" {
		response.OutError = 1
		response.OutMessage = "New Email can't be null"
		return response
	} else if request.VerivicationNewEmail == "" {
		response.OutError = 1
		response.OutMessage = "Verivication New Email can't be null"
		return response
	}
	request.Token = APIkey
	response = repository.UpdateMemberEmail(request)
	return response
}

func UpdateMemberNoHp(w http.ResponseWriter, r *http.Request) model.UpdateMemberNoHpResponse {
	var response model.UpdateMemberNoHpResponse
	var request model.UpdateMemberNoHp
	APIkey := r.Header.Get("API-Key")
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		response.OutError = 1
		response.OutMessage = "Can't bind struct"
	} else if APIkey == "" {
		response.OutError = 1
		response.OutMessage = "API Key can't be null"
	} else if request.OldNoHp == "" {
		response.OutError = 1
		response.OutMessage = "Old NoHp can't be null"
		return response
	} else if request.NewNoHp == "" {
		response.OutError = 1
		response.OutMessage = "New NoHp can't be null"
		return response
	} else if request.VerivicationNewNoHp == "" {
		response.OutError = 1
		response.OutMessage = "Verivication New NoHp can't be null"
		return response
	}
	request.Token = APIkey
	response = repository.UpdateMemberNoHp(request)
	return response
}

func UpdateMemberPassword(w http.ResponseWriter, r *http.Request) model.UpdateMemberPasswordResponse {
	var response model.UpdateMemberPasswordResponse
	var request model.UpdateMemberPassword
	APIkey := r.Header.Get("API-Key")
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		response.OutError = 1
		response.OutMessage = "Can't bind struct"
	} else if APIkey == "" {
		response.OutError = 1
		response.OutMessage = "API Key can't be null"
	} else if request.OldPassword == "" {
		response.OutError = 1
		response.OutMessage = "Old Password can't be null"
		return response
	} else if request.NewPassword == "" {
		response.OutError = 1
		response.OutMessage = "New Password can't be null"
		return response
	} else if request.VerivicationNewPassword == "" {
		response.OutError = 1
		response.OutMessage = "Password New Email can't be null"
		return response
	}
	request.Token = APIkey
	response = repository.UpdateMemberPassword(request)
	return response
}

func ListMember(w http.ResponseWriter, r *http.Request) model.ListMemberResponse {
	var response model.ListMemberResponse
	var request model.ListMember
	APIkey := r.Header.Get("API-Key")
	if APIkey == "" {
		response.OutError = 1
		response.OutMessage = "API Key can't be null"
		return response
	}
	request.Token = APIkey
	response = repository.ListMember(request)
	return response
}

func InsertMemberAddress(w http.ResponseWriter, r *http.Request) model.InsertMemberAddressResponse {
	var response model.InsertMemberAddressResponse
	var request model.InsertMemberAddress
	APIkey := r.Header.Get("API-Key")
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		response.OutError = 1
		response.OutMessage = "Can't bind struct"
	} else if APIkey == "" {
		response.OutError = 1
		response.OutMessage = "API Key can't be null"
	} else if request.TemplateName == "" {
		response.OutError = 1
		response.OutMessage = "Template Name can't be null"
		return response
	} else if request.Street == "" {
		response.OutError = 1
		response.OutMessage = "Street can't be null"
		return response
	} else if request.ProvinceCode == "" {
		response.OutError = 1
		response.OutMessage = "Province Code can't be null"
		return response
	} else if request.CityCode == "" {
		response.OutError = 1
		response.OutMessage = "City Code can't be null"
		return response
	} else if request.DistrictCode == "" {
		response.OutError = 1
		response.OutMessage = "District Code can't be null"
		return response
	} else if request.VillageCode == "" {
		response.OutError = 1
		response.OutMessage = "Village Code can't be null"
		return response
	} else if request.PostalCode == "" {
		response.OutError = 1
		response.OutMessage = "Postal Code can't be null"
		return response
	}
	request.Token = APIkey
	response = repository.InsertMemberAddress(request)
	return response
}

func UpdateMemberAddress(w http.ResponseWriter, r *http.Request) model.UpdateMemberAddressResponse {
	var response model.UpdateMemberAddressResponse
	var request model.UpdateMemberAddress
	APIkey := r.Header.Get("API-Key")
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		response.OutError = 1
		response.OutMessage = "Can't bind struct"
	} else if APIkey == "" {
		response.OutError = 1
		response.OutMessage = "API Key can't be null"
	} else if request.TemplateName == "" {
		response.OutError = 1
		response.OutMessage = "Template Name can't be null"
		return response
	} else if request.Street == "" {
		response.OutError = 1
		response.OutMessage = "Street can't be null"
		return response
	} else if request.ProvinceCode == "" {
		response.OutError = 1
		response.OutMessage = "Province Code can't be null"
		return response
	} else if request.CityCode == "" {
		response.OutError = 1
		response.OutMessage = "City Code can't be null"
		return response
	} else if request.DistrictCode == "" {
		response.OutError = 1
		response.OutMessage = "District Code can't be null"
		return response
	} else if request.VillageCode == "" {
		response.OutError = 1
		response.OutMessage = "Village Code can't be null"
		return response
	} else if request.PostalCode == "" {
		response.OutError = 1
		response.OutMessage = "Postal Code can't be null"
		return response
	}
	request.Token = APIkey
	response = repository.UpdateMemberAddress(request)
	return response
}

func UpdateMemberMainAddress(w http.ResponseWriter, r *http.Request) model.UpdateMemberMainAddressResponse {
	var response model.UpdateMemberMainAddressResponse
	var request model.UpdateMemberMainAddress
	APIkey := r.Header.Get("API-Key")
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		response.OutError = 1
		response.OutMessage = "Can't bind struct"
	} else if APIkey == "" {
		response.OutError = 1
		response.OutMessage = "API Key can't be null"
	} else if request.TemplateName == "" {
		response.OutError = 1
		response.OutMessage = "Template Name can't be null"
		return response
	}
	request.Token = APIkey
	response = repository.UpdateMemberMainAddress(request)
	return response
}

func DeleteMemberAddress(w http.ResponseWriter, r *http.Request) model.DeleteMemberAddressResponse {
	var response model.DeleteMemberAddressResponse
	var request model.DeleteMemberAddress
	APIkey := r.Header.Get("API-Key")
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		response.OutError = 1
		response.OutMessage = "Can't bind struct"
	} else if APIkey == "" {
		response.OutError = 1
		response.OutMessage = "API Key can't be null"
	} else if request.TemplateName == "" {
		response.OutError = 1
		response.OutMessage = "Template Name can't be null"
		return response
	}
	request.Token = APIkey
	response = repository.DeleteMemberAddress(request)
	return response
}

func UpdateMemberPhotoProfile(w http.ResponseWriter, r *http.Request) model.UpdateMemberPhotoResponse {
	var response model.UpdateMemberPhotoResponse
	var request model.UpdateMemberPhoto
	var getMemberPath model.GetMemberPathResponse
	APIkey := r.Header.Get("API-Key")
	//get a ref to the parsed multipart form
	//get the *fileheaders
	_, h, err := r.FormFile("photoFile")
	if err != nil {
		response.OutError = 1
		response.OutMessage = "File can't be null"
	} else if APIkey == "" {
		response.OutError = 1
		response.OutMessage = "API Key can't be null"
	}
	extention := filepath.Ext(h.Filename)
	tittleDirectory := "/profile/"
	//getMember Path Ke Procedure
	getMemberPath = GetMemberPath(APIkey)
	//decrypted global path
	decryptedGlobal, err := hex.DecodeString(getMemberPath.GlobalPath)
	decryptedPathGlobal := util.Decrypt(decryptedGlobal, []byte(getMemberPath.Key), getMemberPath.InitialVector)
	fmt.Println(string(decryptedPathGlobal))
	//decrypted public path
	decryptedPublic, err := hex.DecodeString(getMemberPath.PublicPath)
	decryptedPathPublic := util.Decrypt(decryptedPublic, []byte(getMemberPath.Key), getMemberPath.InitialVector)
	fmt.Println(string(decryptedPathPublic))
	pathGlobal := string(decryptedPathGlobal) + tittleDirectory + getMemberPath.AccountId + extention
	pathPublic := string(decryptedPathPublic) + tittleDirectory + getMemberPath.AccountId + extention
	//check file Exist
	status, err := util.Exists(pathGlobal)
	if err != nil {
		response.OutError = 1
		response.OutMessage = "error directory"
	} else if status == false {
		log.Println("kosong")
		log.Println(pathGlobal)
		response = util.UploadHandler(w, r, pathGlobal)
	} else {
		log.Println("udah ada")
		util.DeleteHandler(pathGlobal)
		response = util.UploadHandler(w, r, pathGlobal)
	}

	if response.OutError != 1 {
		request.Token = APIkey
		request.ProfilePath = tittleDirectory + getMemberPath.AccountId + extention
		response = repository.UpdateMemberPhoto(request)
		if response.OutError != 1 {
			response.ProfilePath = pathPublic
		}
	}
	return response
}

func UpdateMemberPhotoBackground(w http.ResponseWriter, r *http.Request) model.UpdateMemberPhotoResponse {
	var response model.UpdateMemberPhotoResponse
	var request model.UpdateMemberPhoto
	var getMemberPath model.GetMemberPathResponse
	APIkey := r.Header.Get("API-Key")
	//get a ref to the parsed multipart form
	//get the *fileheaders
	_, h, err := r.FormFile("photoFile")
	if err != nil {
		response.OutError = 1
		response.OutMessage = "File can't be null"
	} else if APIkey == "" {
		response.OutError = 1
		response.OutMessage = "API Key can't be null"
	}
	extention := filepath.Ext(h.Filename)
	tittleDirectory := "/background/"
	//getMember Path Ke Procedure
	getMemberPath = GetMemberPath(APIkey)
	//decrypted global path
	decryptedGlobal, err := hex.DecodeString(getMemberPath.GlobalPath)
	decryptedPathGlobal := util.Decrypt(decryptedGlobal, []byte(getMemberPath.Key), getMemberPath.InitialVector)
	fmt.Println(string(decryptedPathGlobal))
	//decrypted public path
	decryptedPublic, err := hex.DecodeString(getMemberPath.PublicPath)
	decryptedPathPublic := util.Decrypt(decryptedPublic, []byte(getMemberPath.Key), getMemberPath.InitialVector)
	fmt.Println(string(decryptedPathPublic))
	pathGlobal := string(decryptedPathGlobal) + tittleDirectory + getMemberPath.AccountId + extention
	pathPublic := string(decryptedPathPublic) + tittleDirectory + getMemberPath.AccountId + extention
	//check file Exist
	status, err := util.Exists(pathGlobal)
	if err != nil {
		response.OutError = 1
		response.OutMessage = "error directory"
	} else if status == false {
		log.Println("kosong")
		log.Println(pathGlobal)
		response = util.UploadHandler(w, r, pathGlobal)
	} else {
		log.Println("udah ada")
		util.DeleteHandler(pathGlobal)
		response = util.UploadHandler(w, r, pathGlobal)
	}

	if response.OutError != 1 {
		request.Token = APIkey
		request.BackgroundPath = tittleDirectory + getMemberPath.AccountId + extention
		response = repository.UpdateMemberPhoto(request)
		if response.OutError != 1 {
			response.BackgroundPath = pathPublic
		}
	}
	return response
}

func GetMemberCart(w http.ResponseWriter, r *http.Request) model.MemberCartResponse {
	var response model.MemberCartResponse
	var request model.MemberCart
	APIkey := r.Header.Get("API-Key")
	if APIkey == "" {
		response.OutError = 1
		response.OutMessage = "API Key can't be null"
		return response
	}
	request.Token = APIkey
	response = repository.GetMemberCart(request)
	return response
}

func GetMemberWishList(w http.ResponseWriter, r *http.Request) model.MemberWishListResponse {
	var response model.MemberWishListResponse
	var request model.MemberWishList
	APIkey := r.Header.Get("API-Key")
	if APIkey == "" {
		response.OutError = 1
		response.OutMessage = "API Key can't be null"
		return response
	}
	request.Token = APIkey
	response = repository.GetMemberWishList(request)
	return response
}

func GetMemberAddress(w http.ResponseWriter, r *http.Request) model.GetMemberAddressResponse {
	var response model.GetMemberAddressResponse
	var request model.GetMemberAddress
	APIkey := r.Header.Get("API-Key")
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		response.OutError = 1
		response.OutMessage = "File can't be null"
	} else if APIkey == "" {
		response.OutError = 1
		response.OutMessage = "API Key can't be null"
		return response
	}
	request.Token = APIkey
	response = repository.GetMemberAddress(request)
	return response
}

func GetMemberVoucher(w http.ResponseWriter, r *http.Request) model.GetMemberVoucherResponse {
	var response model.GetMemberVoucherResponse
	var request model.GetMemberVoucher
	APIkey := r.Header.Get("API-Key")
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		response.OutError = 1
		response.OutMessage = "File can't be null"
	} else if APIkey == "" {
		response.OutError = 1
		response.OutMessage = "API Key can't be null"
		return response
	}
	request.Token = APIkey
	response = repository.GetMemberVoucher(request)
	return response
}

func GetMemberTransaction(w http.ResponseWriter, r *http.Request) model.GetMemberTrxResponse {
	var response model.GetMemberTrxResponse
	var request model.GetMemberTrx
	APIkey := r.Header.Get("API-Key")
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		response.OutError = 1
		response.OutMessage = "Can't bind struct"
	} else if APIkey == "" {
		response.OutError = 1
		response.OutMessage = "API Key can't be null"
		return response
	} else if request.OrderStatus == "" {
		response.OutError = 1
		response.OutMessage = "Order status can't be null"
		return response
	}
	request.Token = APIkey
	response = repository.GetMemberTransaction(request)
	return response
}

// noted for encrypt
// var plainText = "hello world"
// encryptedData := util.Encrypt(plainText, []byte(passphrase))
// encryptedString := base64.StdEncoding.EncodeToString(encryptedData)
// fmt.Println(encryptedString)
