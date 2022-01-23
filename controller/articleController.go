package controller

import (
	"benings/model"
	"benings/service"
	"encoding/json"
	"net/http"
)

func ListArticle(w http.ResponseWriter, r *http.Request) {
	var response model.ListArticleResponse
	response = service.GetListArticle(w, r)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func DetailArticle(w http.ResponseWriter, r *http.Request) {
	var response model.DetailArticleResponse
	response = service.DetailArticle(w, r)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
