package service

import (
	"benings/model"
	"benings/repository"
	"encoding/json"
	"net/http"
)

func GetListArticle(w http.ResponseWriter, r *http.Request) model.ListArticleResponse {
	var response model.ListArticleResponse
	var request model.ListArticle
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
	response = repository.GetListArticle(request)
	return response
}

func DetailArticle(w http.ResponseWriter, r *http.Request) model.DetailArticleResponse {
	var response model.DetailArticleResponse
	var request model.DetailArticle
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		response.OutError = 1
		response.OutMessage = "File can't be null"
	} else if request.ArticleCode == "" {
		response.OutError = 1
		response.OutMessage = "Article Code can't be null"
	}
	response = repository.DetailArticle(request)
	return response
}
