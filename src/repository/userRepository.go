package repository

import (
	"benings/config"
	"benings/model"
	"benings/util"
	"fmt"
	"time"

	"github.com/ian-kent/go-log/log"
)

func ValidationOtp(request model.Register) model.ValidationResponse {
	var response model.ValidationResponse
	db := config.Connect(config.ReadDatabaseConfig())
	defer db.Close()

	insmember, err := db.Query("call flow.ins_member ($1,$2,$3,$4,$5,$6,$7,$8,$9)",
		request.UserName,
		request.Email,
		request.PhoneNumber,
		request.Gender,
		request.Bod,
		request.Password,
		request.OutId,
		request.OutError,
		request.OutMessage,
	)
	if err != nil {
		response.OutError = 1
		response.OutMessage = err.Error()
		log.Error(err.Error())
	} else {
		for insmember.Next() {
			err = insmember.Scan(&response.OutId, &response.OutError, &response.OutMessage)
			if err != nil {
				response.OutError = 1
				response.OutMessage = err.Error()
				log.Error(err.Error())
			}
		}
	}
	return response
}

func Login(request model.Login) model.LoginResponse {
	var response model.LoginResponse
	var jwtResponse model.JwtResponse
	db := config.Connect(config.ReadDatabaseConfig())
	defer db.Close()
	dt := time.Now()
	jwtResponse = util.TokenUsingJWT(request.UserName + dt.String())
	insmember, err := db.Query("call flow.login_member ($1,$2,$3,$4,$5)",
		request.UserName,
		request.Password,
		jwtResponse.Token,
		request.OutError,
		request.OutMessage,
	)
	if err != nil {
		response.OutError = 1
		response.OutMessage = err.Error()
		log.Error(err.Error())
	} else {
		for insmember.Next() {
			err = insmember.Scan(&response.Token, &response.OutError, &response.OutMessage)
			if err != nil {
				response.OutError = 1
				response.OutMessage = err.Error()
				log.Error(err.Error())
			}
		}
	}
	return response
}

func Logout(request model.Logout) model.LogoutResponse {
	var response model.LogoutResponse
	db := config.Connect(config.ReadDatabaseConfig())
	defer db.Close()

	logOutMember, err := db.Query("call flow.logout_member ($1,$2,$3)",
		request.Token,
		request.OutError,
		request.OutMessage,
	)
	if err != nil {
		response.OutError = 1
		response.OutMessage = err.Error()
		log.Error(err.Error())
	} else {
		for logOutMember.Next() {
			err = logOutMember.Scan(&response.OutError, &response.OutMessage)
			if err != nil {
				response.OutError = 1
				response.OutMessage = err.Error()
				log.Error(err.Error())
			}
		}
	}

	return response
}

func ForgotPassword(request model.ForgotPassword) model.ForgotPasswordResponse {
	var response model.ForgotPasswordResponse
	db := config.Connect(config.ReadDatabaseConfig())
	defer db.Close()

	forgotMemberPassword, err := db.Query("call flow.forgot_member_password ($1,$2,$3,$4,$5,$6,$7)",
		request.UserName,
		request.Otp,
		request.OutNewRandomPassword,
		request.OutToken,
		request.OutMemberId,
		request.OutError,
		request.OutMessage,
	)
	if err != nil {
		response.OutError = 1
		response.OutMessage = err.Error()
		log.Error(err.Error())
	} else {
		for forgotMemberPassword.Next() {
			err = forgotMemberPassword.Scan(&response.OutOtp, &response.OutNewRandomPassword, &response.OutToken, &response.OutMemberId, &response.OutError, &response.OutMessage)
			if err != nil {
				response.OutError = 1
				response.OutMessage = err.Error()
				log.Error(err.Error())
			}
		}
	}
	return response
}

func CheckMemberPassword(request model.CheckMemberPassword) model.CheckMemberPasswordResponse {
	var response model.CheckMemberPasswordResponse
	db := config.Connect(config.ReadDatabaseConfig())
	defer db.Close()

	forgotMemberPassword, err := db.Query("call flow.check_member_password ($1,$2,$3,$4,$5)",
		request.Token,
		request.MemberPassword,
		request.Email,
		request.OutError,
		request.OutMessage,
	)
	if err != nil {
		response.OutError = 1
		response.OutMessage = err.Error()
		log.Error(err.Error())
	} else {
		for forgotMemberPassword.Next() {
			err = forgotMemberPassword.Scan(&response.Email, &response.OutError, &response.OutMessage)
			if err != nil {
				response.OutError = 1
				response.OutMessage = err.Error()
				log.Error(err.Error())
			}
		}
	}
	return response
}

func UpdateMemberEmail(request model.UpdateMemberEmail) model.UpdateMemberEmailResponse {
	var response model.UpdateMemberEmailResponse
	db := config.Connect(config.ReadDatabaseConfig())
	defer db.Close()

	updMemberEmail, err := db.Query("call flow.upd_member_email ($1,$2,$3,$4,$5,$6)",
		request.Token,
		request.OldEmail,
		request.NewEmail,
		request.VerivicationNewEmail,
		request.OutError,
		request.OutMessage,
	)
	if err != nil {
		response.OutError = 1
		response.OutMessage = err.Error()
		log.Error(err.Error())
	} else {
		for updMemberEmail.Next() {
			err = updMemberEmail.Scan(&response.OutError, &response.OutMessage)
			if err != nil {
				response.OutError = 1
				response.OutMessage = err.Error()
				log.Error(err.Error())
			}
		}
	}
	return response
}

func UpdateMemberNoHp(request model.UpdateMemberNoHp) model.UpdateMemberNoHpResponse {
	var response model.UpdateMemberNoHpResponse
	db := config.Connect(config.ReadDatabaseConfig())
	defer db.Close()

	updMemberNoHp, err := db.Query("call flow.upd_member_nohp ($1,$2,$3,$4,$5,$6)",
		request.Token,
		request.OldNoHp,
		request.NewNoHp,
		request.VerivicationNewNoHp,
		request.OutError,
		request.OutMessage,
	)
	if err != nil {
		response.OutError = 1
		response.OutMessage = err.Error()
		log.Error(err.Error())
	} else {
		for updMemberNoHp.Next() {
			err = updMemberNoHp.Scan(&response.OutError, &response.OutMessage)
			if err != nil {
				response.OutError = 1
				response.OutMessage = err.Error()
				log.Error(err.Error())
			}
		}
	}
	return response
}

func UpdateMemberPassword(request model.UpdateMemberPassword) model.UpdateMemberPasswordResponse {
	var response model.UpdateMemberPasswordResponse
	db := config.Connect(config.ReadDatabaseConfig())
	defer db.Close()
	updMemberPassword, err := db.Query("call flow.upd_member_password ($1,$2,$3,$4,$5,$6)",
		request.Token,
		request.OldPassword,
		request.NewPassword,
		request.VerivicationNewPassword,
		request.OutError,
		request.OutMessage,
	)
	if err != nil {
		response.OutError = 1
		response.OutMessage = err.Error()
		log.Error(err.Error())
	} else {
		for updMemberPassword.Next() {
			err = updMemberPassword.Scan(&response.OutError, &response.OutMessage)
			if err != nil {
				response.OutError = 1
				response.OutMessage = err.Error()
				log.Error(err.Error())
			}
		}
	}
	return response
}

func ListMember(request model.ListMember) model.ListMemberResponse {
	var response model.ListMemberResponse
	var member model.Member
	db := config.Connect(config.ReadDatabaseConfig())
	defer db.Close()
	listMember, err := db.Query("select * from flow.get_list_member ($1)",
		request.Token,
	)
	if err != nil {
		response.OutError = 1
		response.OutMessage = err.Error()
		log.Error(err.Error())
	} else {
		for listMember.Next() {
			err = listMember.Scan(&member.AccountId, &member.MemberName, &member.MemberEmail, &member.MemberPhone, &member.MemberGender, &member.MemberBod, &member.PhotoPath, &member.BackgroundPhotoPath, &member.MemberAddress)
			if err != nil {
				response.OutError = 1
				response.OutMessage = err.Error()
				log.Error(err.Error())
			}
			response.OutError = 0
			response.OutMessage = "success"
			response.OutData = append(response.OutData, member)
		}
	}
	return response
}

func InsertMemberAddress(request model.InsertMemberAddress) model.InsertMemberAddressResponse {
	var response model.InsertMemberAddressResponse
	db := config.Connect(config.ReadDatabaseConfig())
	defer db.Close()
	insMemberAddress, err := db.Query("call flow.ins_member_address ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15,$16)",
		request.Token,
		request.TemplateName,
		request.ReceiverName,
		request.ReceiverPhoneNumber,
		request.Street,
		request.ProvinceCode,
		request.CityCode,
		request.DistrictCode,
		request.VillageCode,
		request.PostalCode,
		request.Note,
		request.Longitude,
		request.Latitude,
		request.OutId,
		request.OutError,
		request.OutMessage,
	)
	if err != nil {
		response.OutError = 1
		response.OutMessage = err.Error()
		log.Error(err.Error())
	} else {
		for insMemberAddress.Next() {
			err = insMemberAddress.Scan(&response.OutId, &response.OutError, &response.OutMessage)
			if err != nil {
				response.OutError = 1
				response.OutMessage = err.Error()
				log.Error(err.Error())
			}
		}
	}
	return response
}

func UpdateMemberAddress(request model.UpdateMemberAddress) model.UpdateMemberAddressResponse {
	var response model.UpdateMemberAddressResponse
	db := config.Connect(config.ReadDatabaseConfig())
	defer db.Close()
	updMemberAddress, err := db.Query("call flow.upd_member_address ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15)",
		request.Token,
		request.TemplateName,
		request.ReceiverName,
		request.ReceiverPhoneNumber,
		request.Street,
		request.ProvinceCode,
		request.CityCode,
		request.DistrictCode,
		request.VillageCode,
		request.PostalCode,
		request.Note,
		request.Longitude,
		request.Latitude,
		request.OutError,
		request.OutMessage,
	)
	if err != nil {
		response.OutError = 1
		response.OutMessage = err.Error()
		log.Error(err.Error())
	} else {
		for updMemberAddress.Next() {
			err = updMemberAddress.Scan(&response.OutError, &response.OutMessage)
			if err != nil {
				response.OutError = 1
				response.OutMessage = err.Error()
				log.Error(err.Error())
			}
		}
	}
	return response
}

func UpdateMemberMainAddress(request model.UpdateMemberMainAddress) model.UpdateMemberMainAddressResponse {
	var response model.UpdateMemberMainAddressResponse
	db := config.Connect(config.ReadDatabaseConfig())
	defer db.Close()
	updMemberMainAddress, err := db.Query("call flow.upd_member_main_address ($1,$2,$3,$4)",
		request.Token,
		request.TemplateName,
		request.OutError,
		request.OutMessage,
	)
	if err != nil {
		response.OutError = 1
		response.OutMessage = err.Error()
		log.Error(err.Error())
	} else {
		for updMemberMainAddress.Next() {
			err = updMemberMainAddress.Scan(&response.OutError, &response.OutMessage)
			if err != nil {
				response.OutError = 1
				response.OutMessage = err.Error()
				log.Error(err.Error())
			}
		}
	}
	return response
}

func DeleteMemberAddress(request model.DeleteMemberAddress) model.DeleteMemberAddressResponse {
	var response model.DeleteMemberAddressResponse
	db := config.Connect(config.ReadDatabaseConfig())
	defer db.Close()
	delMemberAddress, err := db.Query("call flow.del_member_address ($1,$2,$3,$4)",
		request.Token,
		request.TemplateName,
		request.OutError,
		request.OutMessage,
	)
	if err != nil {
		response.OutError = 1
		response.OutMessage = err.Error()
		log.Error(err.Error())
	} else {
		for delMemberAddress.Next() {
			err = delMemberAddress.Scan(&response.OutError, &response.OutMessage)
			if err != nil {
				response.OutError = 1
				response.OutMessage = err.Error()
				log.Error(err.Error())
			}
		}
	}
	return response
}

func UpdateMemberPhoto(request model.UpdateMemberPhoto) model.UpdateMemberPhotoResponse {
	var response model.UpdateMemberPhotoResponse
	db := config.Connect(config.ReadDatabaseConfig())
	defer db.Close()
	updMemberPhoto, err := db.Query("call flow.upd_member_photo ($1,$2,$3,$4,$5)",
		request.Token,
		request.ProfilePath,
		request.BackgroundPath,
		request.OutError,
		request.OutMessage,
	)
	if err != nil {
		response.OutError = 1
		response.OutMessage = err.Error()
		log.Error(err.Error())
	} else {
		for updMemberPhoto.Next() {
			err = updMemberPhoto.Scan(&response.OutError, &response.OutMessage)
			if err != nil {
				response.OutError = 1
				response.OutMessage = err.Error()
				log.Error(err.Error())
			}
		}
	}
	return response
}

func APIkey(APIkey string, UserId string) model.Response {
	var response model.Response
	var validasiToken model.ResponseLogin
	db := config.Connect(config.ReadDatabaseConfig())
	defer db.Close()

	value, err := db.Query(`Select token from  public.benings_user_login Where user_id = $1 and token = $2`, UserId, APIkey)
	if err != nil {
		fmt.Println("Error select by token:", err)
		response.OutError = 1
		response.OutMessage = "Failed Token"
	} else {
		fmt.Println(value)
		for value.Next() {
			err = value.Scan(&validasiToken.Token)
			if err != nil {
				panic(err)
			}
		}
		if validasiToken.Token != "" {
			response.OutError = 0
			response.OutMessage = "success"
		} else {
			response.OutError = 1
			response.OutMessage = "Wrong API-Key or userId"
		}

	}
	return response
}

func GetUserByUserName(userName string) model.User {
	var response model.User
	db := config.Connect(config.ReadDatabaseConfig())
	defer db.Close()
	value, err := db.Query(`Select user_id, user_name, email, password, address from  public.benings_user Where user_name = $1`, userName)
	if err != nil {
		fmt.Println("Error select:", err)
		response.OutError = 1
		response.OutMessage = "Failed"
	} else {
		fmt.Println(value)
		for value.Next() {
			err = value.Scan(&response.UserId, &response.UserName, &response.Email, &response.Password, &response.Address)
			if err != nil {
				panic(err)
			}
		}
		if response.UserName != "" {
			response.OutError = 1
			response.OutMessage = "Name Is already Exist"
		} else {
			response.OutError = 0
			response.OutMessage = "success"
		}
	}
	return response
}

func GetMemberCart(request model.MemberCart) model.MemberCartResponse {
	var response model.MemberCartResponse
	db := config.Connect(config.ReadDatabaseConfig())
	defer db.Close()
	memberCart, err := db.Query("select * from flow.get_member_cart ($1)",
		request.Token,
	)
	if err != nil {
		response.OutError = 1
		response.OutMessage = err.Error()
		log.Error(err.Error())
	} else {
		for memberCart.Next() {
			err = memberCart.Scan(&response.ProductId, &response.ProductCode, &response.ProductName, &response.Quantity, &response.TotalPrice)
			if err != nil {
				response.OutError = 1
				response.OutMessage = err.Error()
				log.Error(err.Error())
			}
			response.OutError = 0
			response.OutMessage = "success"
		}
	}
	return response
}

func GetMemberWishList(request model.MemberWishList) model.MemberWishListResponse {
	var response model.MemberWishListResponse
	var wishList model.WishList
	db := config.Connect(config.ReadDatabaseConfig())
	defer db.Close()
	memberWishList, err := db.Query("select * from flow.get_member_wishlist ($1)",
		request.Token,
	)
	if err != nil {
		response.OutError = 1
		response.OutMessage = err.Error()
		log.Error(err.Error())
	} else {
		for memberWishList.Next() {
			err = memberWishList.Scan(&wishList.ProductId, &wishList.ProductCode, &wishList.ProductName, &wishList.Price)
			if err != nil {
				response.OutError = 1
				response.OutMessage = err.Error()
				log.Error(err.Error())
			}
			response.OutError = 0
			response.OutMessage = "success"
			response.OutData = append(response.OutData, wishList)
		}
	}
	return response
}

func GetMemberAddress(request model.GetMemberAddress) model.GetMemberAddressResponse {
	var response model.GetMemberAddressResponse
	var memberAddress model.MemberAddress
	db := config.Connect(config.ReadDatabaseConfig())
	defer db.Close()
	log.Println("token : ", request.Token)
	log.Println("PageNumber : ", request.PageNumber)
	log.Println("RowPerpage : ", request.RowPerpage)
	getMemberAddress, err := db.Query("select * from flow.get_member_address ($1,$2,$3)",
		request.Token,
		request.PageNumber,
		request.RowPerpage,
	)
	if err != nil {
		response.OutError = 1
		response.OutMessage = err.Error()
		log.Error(err.Error())
	} else {
		for getMemberAddress.Next() {
			err = getMemberAddress.Scan(&memberAddress.TemplateId, &memberAddress.TemplateName, &memberAddress.ReceiverName, &memberAddress.ReceiverPhoneNumber, &memberAddress.Street, &memberAddress.Province, &memberAddress.City, &memberAddress.District, &memberAddress.Village, &memberAddress.PostalCode, &memberAddress.Note, &memberAddress.Longitude, &memberAddress.Latitude, &memberAddress.MainAddress)
			if err != nil {
				response.OutError = 1
				response.OutMessage = err.Error()
				log.Error(err.Error())
			}
			response.OutError = 0
			response.OutMessage = "success"
			response.OutData = append(response.OutData, memberAddress)
		}
	}
	return response
}

func GetMemberVoucher(request model.GetMemberVoucher) model.GetMemberVoucherResponse {
	var response model.GetMemberVoucherResponse
	var memberVoucher model.MemberVoucher
	db := config.Connect(config.ReadDatabaseConfig())
	defer db.Close()
	getMemberVoucher, err := db.Query("select * from flow.get_member_voucher ($1,$2,$3)",
		request.Token,
		request.PageNumber,
		request.RowPerpage,
	)
	if err != nil {
		response.OutError = 1
		response.OutMessage = err.Error()
		log.Error(err.Error())
	} else {
		for getMemberVoucher.Next() {
			err = getMemberVoucher.Scan(&memberVoucher.VoucherCode, &memberVoucher.VoucherDescription, &memberVoucher.DiscountAmount, &memberVoucher.DiscountPercent, &memberVoucher.MaximumDiscount, &memberVoucher.ExpiredDate)
			if err != nil {
				response.OutError = 1
				response.OutMessage = err.Error()
				log.Error(err.Error())
			}
			response.OutError = 0
			response.OutMessage = "success"
			response.OutData = append(response.OutData, memberVoucher)
		}
	}
	return response
}

func GetMemberTransaction(request model.GetMemberTrx) model.GetMemberTrxResponse {
	var response model.GetMemberTrxResponse
	var memberTrx model.MemberTrx
	db := config.Connect(config.ReadDatabaseConfig())
	defer db.Close()
	getMemberTrx, err := db.Query("select * from trx.get_member_trx ($1,$2,$3,$4)",
		request.Token,
		request.PageNumber,
		request.RowPerpage,
		request.OrderStatus,
	)
	if err != nil {
		response.OutError = 1
		response.OutMessage = err.Error()
		log.Error(err.Error())
	} else {
		for getMemberTrx.Next() {
			err = getMemberTrx.Scan(&memberTrx.HeaderId, &memberTrx.OrderNumber, &memberTrx.OrderStatus, &memberTrx.OrderDate, &memberTrx.PaidDate, &memberTrx.RejectDate, &memberTrx.ValidDate, &memberTrx.OrderItems, &memberTrx.OrderCourier, &memberTrx.OrderDiscount, &memberTrx.OrderPayment)
			if err != nil {
				response.OutError = 1
				response.OutMessage = err.Error()
				log.Error(err.Error())
			}
			response.OutError = 0
			response.OutMessage = "success"
			response.OutData = append(response.OutData, memberTrx)
		}
	}
	return response
}
