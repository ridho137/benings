package main

import (
	"benings/controller"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func RestController() {
	router := mux.NewRouter()
	headers := handlers.AllowedHeaders([]string{"X-Request-With", "Content-Type", "Authorization", "API-Key"})
	origins := handlers.AllowedOrigins([]string{"*"})
	//API MEMBER
	router.HandleFunc("/register", controller.Register)                                          //send email to validatade OTP
	router.HandleFunc("/validation-otp", controller.ValidationOTP)                               //Create to DB after OTP valid
	router.HandleFunc("/login", controller.Login)                                                //login and set token
	router.HandleFunc("/logout", controller.Logout)                                              // logout ubah setatus
	router.HandleFunc("/forgot-password", controller.ForgotPassword)                             // Lupa password dan kirim password baru
	router.HandleFunc("/check-member-password", controller.CheckMemberPassword)                  // Sebelum melakukan update data member wajib check password terlebih dahulu
	router.HandleFunc("/update-member-email", controller.UpdateMemberEmail)                      // update email
	router.HandleFunc("/update-member-nohp", controller.UpdateMemberNoHp)                        // update no hp
	router.HandleFunc("/update-member-password", controller.UpdateMemberPassword)                // update password
	router.HandleFunc("/get-list-member", controller.ListMember)                                 // get list Member
	router.HandleFunc("/get-list-province", controller.ListProvince)                             // get list province
	router.HandleFunc("/get-list-city", controller.ListCity)                                     // get list city
	router.HandleFunc("/get-list-district", controller.ListDistrict)                             // get list district
	router.HandleFunc("/get-list-village", controller.ListVillage)                               // get list village
	router.HandleFunc("/insert-member-address", controller.InsertMemberAddress)                  // InsertMemberAddress
	router.HandleFunc("/update-member-photo-profile", controller.UpdateMemberPhotoProfile)       // UpdateMemberPhotoProfile
	router.HandleFunc("/update-member-photo-background", controller.UpdateMemberPhotoBackground) // UpdateMemberPhotoBackground
	router.HandleFunc("/update-member-address", controller.UpdateMemberAddress)                  // UpdateMemberAddress
	router.HandleFunc("/update-member-main-address", controller.UpdateMemberMainAddress)         // UpdateMemberMainAddress
	router.HandleFunc("/delete-member-address", controller.DeleteMemberAddress)                  // DeleteMemberAddress
	router.HandleFunc("/get-member-address", controller.GetMemberAddress)                        // GetMemberAddress
	router.HandleFunc("/get-member-voucher", controller.GetMemberVoucher)                        // GetMemberVoucher

	//API Product
	router.HandleFunc("/get-list-banner", controller.ListBanner)            //All Benner
	router.HandleFunc("/get-list-product", controller.ListProduct)          //All product
	router.HandleFunc("/get-list-article", controller.ListArticle)          //ALL Article pertama di home page parentId di set 0 setelah di klik maka parent id balikan di request kembali
	router.HandleFunc("/get-list-category", controller.ListCategory)        //ALL Category
	router.HandleFunc("/get-detail-article", controller.DetailArticle)      //get Detail Article
	router.HandleFunc("/get-member-cart", controller.GetMemberCart)         //get Member Cart
	router.HandleFunc("/add-cart", controller.AddCart)                      //Add cart
	router.HandleFunc("/add-cart-quantity", controller.AddCartByQuantity)   //AddCartByQuantity
	router.HandleFunc("/delete-cart", controller.DeleteCart)                //DeleteCart
	router.HandleFunc("/get-member-wishlist", controller.GetMemberWishList) //get Member wishlist
	router.HandleFunc("/add-wishlist", controller.AddWishList)              //AddWishList
	router.HandleFunc("/delete-wishlist", controller.DeleteWishList)        //DeleteWishList
	router.HandleFunc("/get-list-type", controller.ListType)                //All Type

	// API checkOutProcess
	router.HandleFunc("/create-order", controller.CreateOrder)                                 //CreateOrder
	router.HandleFunc("/get-discount-amount", controller.GetDiscountAmount)                    //All GetDiscountAmount
	router.HandleFunc("/create-order-discount-voucher", controller.CreateOrderDiscountVoucher) //CreateOrderDiscountVoucher
	router.HandleFunc("/create-order-payment", controller.CreateOrderPayment)                  //CreateOrderPayment

	router.HandleFunc("/get-list-courier", controller.ListCourier)                //ListCourier
	router.HandleFunc("/get-list-courier-service", controller.ListCourierService) //ListCourierService
	router.HandleFunc("/get-list-payment-method", controller.ListPaymentMethod)   //ListPaymentMethod
	router.HandleFunc("/update-header-paid", controller.UpdateHeaderPaid)         //UpdateHeaderPaid
	//masih belum tau
	router.HandleFunc("/update-header-do", controller.UpdateHeaderDo)             //UpdateHeaderDo
	router.HandleFunc("/get-member-transaction", controller.GetMemberTransaction) //GetMemberTransaction

	// router.HandleFunc("/getDetailMember", controller.Article)                         // get Article
	// router.HandleFunc("/getRiwayatTransactions", controller.Article)                         // get Article

	http.Handle("/", router)
	fmt.Println("Successfully connected!")
	log.Fatal(http.ListenAndServe(":8822", handlers.CORS(headers, origins)(router)))
}
