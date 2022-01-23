package repository

import (
	"benings/config"
	"benings/model"
	"benings/util"

	"github.com/ian-kent/go-log/log"
)

func ListProvince(request model.ListProvince) model.ListProvinceResponse {
	var response model.ListProvinceResponse
	var province model.Province
	db := config.Connect(config.ReadDatabaseConfig())
	defer db.Close()
	listProvince, err := db.Query("select * from flow.get_lov_province ($1,$2,$3)",
		request.PageNumber,
		request.RowPerpage,
		request.Search,
	)
	if err != nil {
		response.OutError = 1
		response.OutMessage = err.Error()
		log.Error(err.Error())
	} else {
		for listProvince.Next() {
			err = listProvince.Scan(&province.ProvinceCode, &province.ProvinceName)
			if err != nil {
				response.OutError = util.DefaultOutErrorFailed
				response.OutMessage = err.Error()
				log.Error(err.Error())
			}
			response.OutError = 0
			response.OutMessage = "success"
			response.OutData = append(response.OutData, province)
		}
	}
	return response
}

func ListCity(request model.ListCity) model.ListCityResponse {
	var response model.ListCityResponse
	var city model.City
	db := config.Connect(config.ReadDatabaseConfig())
	defer db.Close()
	listCity, err := db.Query("select * from flow.get_lov_city ($1,$2,$3,$4)",
		request.ProvinceCode,
		request.PageNumber,
		request.RowPerpage,
		request.Search,
	)
	if err != nil {
		response.OutError = 1
		response.OutMessage = err.Error()
		log.Error(err.Error())
	} else {
		for listCity.Next() {
			err = listCity.Scan(&city.CityCode, &city.CityName)
			if err != nil {
				response.OutError = util.DefaultOutErrorFailed
				response.OutMessage = err.Error()
				log.Error(err.Error())
			}
			response.OutError = 0
			response.OutMessage = "success"
			response.OutData = append(response.OutData, city)
		}
	}
	return response
}

func ListDistrict(request model.ListDistrict) model.ListDistrictResponse {
	var response model.ListDistrictResponse
	var district model.District
	db := config.Connect(config.ReadDatabaseConfig())
	defer db.Close()
	listDistrict, err := db.Query("select * from flow.get_lov_district ($1,$2,$3,$4)",
		request.CityCode,
		request.PageNumber,
		request.RowPerpage,
		request.Search,
	)
	if err != nil {
		response.OutError = 1
		response.OutMessage = err.Error()
		log.Error(err.Error())
	} else {
		for listDistrict.Next() {
			err = listDistrict.Scan(&district.DistrictCode, &district.DistrictName)
			if err != nil {
				response.OutError = util.DefaultOutErrorFailed
				response.OutMessage = err.Error()
				log.Error(err.Error())
			}
			response.OutError = 0
			response.OutMessage = "success"
			response.OutData = append(response.OutData, district)
		}
	}
	return response
}

func ListVillage(request model.ListVillage) model.ListVillageResponse {
	var response model.ListVillageResponse
	var village model.Village
	db := config.Connect(config.ReadDatabaseConfig())
	defer db.Close()
	listVillage, err := db.Query("select * from flow.get_lov_village ($1,$2,$3,$4)",
		request.DistrictCode,
		request.PageNumber,
		request.RowPerpage,
		request.Search,
	)
	if err != nil {
		response.OutError = 1
		response.OutMessage = err.Error()
		log.Error(err.Error())
	} else {
		for listVillage.Next() {
			err = listVillage.Scan(&village.VillageCode, &village.VillageName, &village.PostalCode)
			if err != nil {
				response.OutError = util.DefaultOutErrorFailed
				response.OutMessage = err.Error()
				log.Error(err.Error())
			}
			response.OutError = 0
			response.OutMessage = "success"
			response.OutData = append(response.OutData, village)
		}
	}
	return response
}

func GetMemberPath(token string) model.GetMemberPathResponse {
	var response model.GetMemberPathResponse
	var request model.GetMemberPath
	db := config.Connect(config.ReadDatabaseConfig())
	defer db.Close()
	getMember, err := db.Query("call flow.get_member_path ($1,$2,$3,$4,$5,$6,$7,$8)",
		token,
		request.Key,
		request.InitialVector,
		request.GlobalPath,
		request.PublicPath,
		request.AccountId,
		request.OutError,
		request.OutMessage,
	)
	if err != nil {
		response.OutError = 1
		response.OutMessage = err.Error()
		log.Error(err.Error())
	} else {
		for getMember.Next() {
			err = getMember.Scan(&response.Key, &response.InitialVector, &response.GlobalPath, &response.PublicPath, &response.AccountId, &response.OutError, &response.OutMessage)
			if err != nil {
				response.OutError = util.DefaultOutErrorFailed
				response.OutMessage = err.Error()
				log.Error(err.Error())
			} else {
				response.OutError = 0
				response.OutMessage = "success"
			}
		}
	}
	return response
}

func LogInsOrder(request model.LogInsOrder) model.LogInsOrderResponse {
	var response model.LogInsOrderResponse
	db := config.Connect(config.ReadDatabaseConfig())
	defer db.Close()
	logInsOrder, err := db.Query("call trx.ins_order_log ($1,$2,$3,$4,$5)",
		request.Token,
		request.Request,
		request.OutId,
		request.OutError,
		request.OutMessage,
	)
	if err != nil {
		response.OutError = 1
		response.OutMessage = err.Error()
		log.Error(err.Error())
	} else {
		for logInsOrder.Next() {
			err = logInsOrder.Scan(&response.OutId, &response.OutError, &response.OutMessage)
			if err != nil {
				response.OutError = util.DefaultOutErrorFailed
				response.OutMessage = err.Error()
				log.Error(err.Error())
			}
		}
	}
	return response
}

func LogUpdOrder(request model.LogUpdOrder) model.LogUpdOrderResponse {
	var response model.LogUpdOrderResponse
	db := config.Connect(config.ReadDatabaseConfig())
	defer db.Close()
	logUpdOrder, err := db.Query("call trx.upd_order_log ($1,$2,$3,$4)",
		request.LogId,
		request.Response,
		request.OutError,
		request.OutMessage,
	)
	if err != nil {
		response.OutError = 1
		response.OutMessage = err.Error()
		log.Error(err.Error())
	} else {
		for logUpdOrder.Next() {
			err = logUpdOrder.Scan(&response.OutError, &response.OutMessage)
			if err != nil {
				response.OutError = util.DefaultOutErrorFailed
				response.OutMessage = err.Error()
				log.Error(err.Error())
			}
		}
	}
	return response
}
