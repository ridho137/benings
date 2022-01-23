package repository

import (
	"benings/config"
	"benings/model"
	"benings/util"

	"github.com/ian-kent/go-log/log"
)

func GetListBanner(request model.ListBanner) model.ListBannerResponse {
	var response model.ListBannerResponse
	var banner model.Banner
	db := config.Connect(config.ReadDatabaseConfig())
	defer db.Close()
	getListBanner, err := db.Query("select * from flow.get_list_banner ($1,$2,$3)",
		request.Token,
		request.PageNumber,
		request.RowPerpage,
	)
	if err != nil {
		response.OutError = 1
		response.OutMessage = err.Error()
		log.Error(err.Error())
	} else {
		for getListBanner.Next() {
			err = getListBanner.Scan(&banner.BannerCode, &banner.BannerName, &banner.BannerPath, &banner.EffectiveDate, &banner.ExpiredDate)
			if err != nil {
				response.OutError = util.DefaultOutErrorFailed
				response.OutMessage = err.Error()
				log.Error(err.Error())
			}
			response.OutError = 0
			response.OutMessage = "success"
			response.OutData = append(response.OutData, banner)
		}
	}
	return response
}

func GetListProduct(request model.ListProduct) model.ListProductResponse {
	var response model.ListProductResponse
	var product model.Product
	db := config.Connect(config.ReadDatabaseConfig())
	defer db.Close()

	getListProduct, err := db.Query("select * from flow.get_list_product ($1,$2,$3,$4,$5,$6,$7)",
		request.Token,
		request.BannerCode,
		request.CategoryCode,
		request.TypeId,
		request.PageNumber,
		request.RowPerpage,
		request.Search,
	)
	if err != nil {
		response.OutError = 1
		response.OutMessage = err.Error()
		log.Error(err.Error())
	} else {
		for getListProduct.Next() {
			err = getListProduct.Scan(&product.ProductCode, &product.ProductName, &product.ProductPath, &product.ProductPrice, &product.PromoPrice, &product.IsPromo, &product.Netto, &product.Ratting, &product.SoldByMember, &product.TypeDescription, &product.Categories, &product.ProductAttribute, &product.IsWishlist, &product.ProductWeight)
			if err != nil {
				response.OutError = util.DefaultOutErrorFailed
				response.OutMessage = err.Error()
				log.Error(err.Error())
			}
			response.OutError = 0
			response.OutMessage = "success"
			response.OutData = append(response.OutData, product)
		}
	}
	return response
}

func AddCart(request model.AddCart) model.AddCartResponse {
	var response model.AddCartResponse
	db := config.Connect(config.ReadDatabaseConfig())
	defer db.Close()

	addCart, err := db.Query("call flow.add_cart ($1,$2,$3,$4)",
		request.Token,
		request.ProductCode,
		request.OutError,
		request.OutMessage,
	)
	if err != nil {
		response.OutError = 1
		response.OutMessage = err.Error()
		log.Error(err.Error())
	} else {
		for addCart.Next() {
			err = addCart.Scan(&response.OutError, &response.OutMessage)
			if err != nil {
				response.OutError = util.DefaultOutErrorFailed
				response.OutMessage = err.Error()
				log.Error(err.Error())
			}
		}
	}
	return response
}

func AddCartByQuantity(request model.AddCartByQuantity) model.AddCartByQuantityResponse {
	var response model.AddCartByQuantityResponse
	db := config.Connect(config.ReadDatabaseConfig())
	defer db.Close()

	addCartByQty, err := db.Query("call flow.add_cart_by_qty ($1,$2,$3,$4,$5)",
		request.Token,
		request.ProductCode,
		request.Quantity,
		request.OutError,
		request.OutMessage,
	)
	if err != nil {
		response.OutError = 1
		response.OutMessage = err.Error()
		log.Error(err.Error())
	} else {
		for addCartByQty.Next() {
			err = addCartByQty.Scan(&response.OutError, &response.OutMessage)
			if err != nil {
				response.OutError = util.DefaultOutErrorFailed
				response.OutMessage = err.Error()
				log.Error(err.Error())
			}
		}
	}
	return response
}

func DeleteCart(request model.DeleteCart) model.DeleteCartResponse {
	var response model.DeleteCartResponse
	db := config.Connect(config.ReadDatabaseConfig())
	defer db.Close()

	delCart, err := db.Query("call flow.del_cart ($1,$2,$3,$4)",
		request.Token,
		request.ProductCode,
		request.OutError,
		request.OutMessage,
	)
	if err != nil {
		response.OutError = 1
		response.OutMessage = err.Error()
		log.Error(err.Error())
	} else {
		for delCart.Next() {
			err = delCart.Scan(&response.OutError, &response.OutMessage)
			if err != nil {
				response.OutError = util.DefaultOutErrorFailed
				response.OutMessage = err.Error()
				log.Error(err.Error())
			}
		}
	}
	return response
}

func AddWishList(request model.AddWishList) model.AddWishListResponse {
	var response model.AddWishListResponse
	db := config.Connect(config.ReadDatabaseConfig())
	defer db.Close()

	addWishList, err := db.Query("call flow.add_wishlist ($1,$2,$3,$4)",
		request.Token,
		request.ProductCode,
		request.OutError,
		request.OutMessage,
	)
	if err != nil {
		response.OutError = 1
		response.OutMessage = err.Error()
		log.Error(err.Error())
	} else {
		for addWishList.Next() {
			err = addWishList.Scan(&response.OutError, &response.OutMessage)
			if err != nil {
				response.OutError = util.DefaultOutErrorFailed
				response.OutMessage = err.Error()
				log.Error(err.Error())
			}
		}
	}
	return response
}

func DeleteWishList(request model.DeleteWishList) model.DeleteWishListResponse {
	var response model.DeleteWishListResponse
	db := config.Connect(config.ReadDatabaseConfig())
	defer db.Close()

	delWishList, err := db.Query("call flow.del_wishlist ($1,$2,$3,$4)",
		request.Token,
		request.ProductCode,
		request.OutError,
		request.OutMessage,
	)
	if err != nil {
		response.OutError = 1
		response.OutMessage = err.Error()
		log.Error(err.Error())
	} else {
		for delWishList.Next() {
			err = delWishList.Scan(&response.OutError, &response.OutMessage)
			if err != nil {
				response.OutError = util.DefaultOutErrorFailed
				response.OutMessage = err.Error()
				log.Error(err.Error())
			}
		}
	}
	return response
}

func ListCategory(request model.ListCategory) model.ListCategoryResponse {
	var response model.ListCategoryResponse
	var category model.Category
	db := config.Connect(config.ReadDatabaseConfig())
	defer db.Close()

	listCategory, err := db.Query("select * from flow.get_list_category ($1,$2,$3,$4)",
		request.Token,
		request.PageNumber,
		request.RowPerpage,
		request.Search,
	)
	if err != nil {
		response.OutError = 1
		response.OutMessage = err.Error()
		log.Error(err.Error())
	} else {
		for listCategory.Next() {
			err = listCategory.Scan(&category.CategoryCode, &category.CategoryName, &category.CategoryPath, &category.CategoryDescription)
			if err != nil {
				response.OutError = util.DefaultOutErrorFailed
				response.OutMessage = err.Error()
				log.Error(err.Error())
			}
			response.OutError = 0
			response.OutMessage = "success"
			response.OutData = append(response.OutData, category)
		}
	}
	return response
}

func GetListType(request model.ListType) model.ListTypeResponse {
	var response model.ListTypeResponse
	var resType model.Type
	db := config.Connect(config.ReadDatabaseConfig())
	defer db.Close()

	getListType, err := db.Query("select * from flow.get_list_type ($1,$2,$3,$4)",
		request.Token,
		request.PageNumber,
		request.RowPerpage,
		request.Search,
	)
	if err != nil {
		response.OutError = 1
		response.OutMessage = err.Error()
		log.Error(err.Error())
	} else {
		for getListType.Next() {
			err = getListType.Scan(&resType.TypeId, &resType.Type)
			if err != nil {
				response.OutError = util.DefaultOutErrorFailed
				response.OutMessage = err.Error()
				log.Error(err.Error())
			}
			response.OutError = 0
			response.OutMessage = "success"
			response.OutData = append(response.OutData, resType)
		}
	}
	return response
}

// func GetProductByCategory(request model.ProductByCategoryRequest) []model.Product {
// 	var response model.Product
// 	var arr_response []model.Product = make([]model.Product, 0)
// 	db := config.Connect(config.ReadDatabaseConfig())
// 	defer db.Close()

// 	value, err := db.Query(`select product_id,product_category_id,product_name, product_code, description, product_path, product_price, product_status_id, netto, expired from benings_product WHERE product_category_id = $1 ORDER BY product_id LIMIT $2 OFFSET  $3`, request.CategoryProduct, request.Limit, request.Offset)
// 	if err != nil {
// 		fmt.Println("Error select benings_product: ", err)
// 		response.OutError = 1
// 		response.OutMessage = "Null Data"
// 		arr_response = append(arr_response, response)
// 	} else {
// 		fmt.Println(value)
// 		for value.Next() {
// 			err = value.Scan(&response.ProductId, &response.ProductCategoryId, &response.ProductName, &response.ProductCode, &response.Description, &response.ProductPath, &response.ProductPrice, &response.ProductStatus, &response.Netto, &response.Expired)
// 			if err != nil {
// 				response.OutError = 1
// 				response.OutMessage = "Null data"
// 				panic(err)
// 			} else {
// 				response.OutError = 0
// 				response.OutMessage = "success"
// 			}
// 			arr_response = append(arr_response, response)
// 		}
// 	}
// 	return arr_response
// }
