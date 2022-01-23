package repository

import (
	"benings/config"
	"benings/model"
	"benings/util"

	"github.com/ian-kent/go-log/log"
)

func GetListArticle(request model.ListArticle) model.ListArticleResponse {
	var response model.ListArticleResponse
	var article model.Article
	db := config.Connect(config.ReadDatabaseConfig())
	defer db.Close()

	getListArticle, err := db.Query("select * from flow.get_list_article ($1,$2,$3,$4,$5)",
		request.Token,
		request.ParentId,
		request.PageNumber,
		request.RowPerpage,
		request.Search,
	)
	if err != nil {
		response.OutError = 1
		response.OutMessage = err.Error()
		log.Error(err.Error())
	} else {
		for getListArticle.Next() {
			err = getListArticle.Scan(&article.ArticleId, &article.ArticleCode, &article.ArticleName, &article.ArticlePath)
			if err != nil {
				response.OutError = util.DefaultOutErrorFailed
				response.OutMessage = err.Error()
				log.Error(err.Error())
			}
			response.OutError = 0
			response.OutMessage = "success"
			response.OutData = append(response.OutData, article)
		}
	}
	return response
}

func DetailArticle(request model.DetailArticle) model.DetailArticleResponse {
	var response model.DetailArticleResponse
	db := config.Connect(config.ReadDatabaseConfig())
	defer db.Close()

	detailArticle, err := db.Query("select flow.get_article_description ($1)",
		request.ArticleCode,
	)
	if err != nil {
		response.OutError = 1
		response.OutMessage = err.Error()
		log.Error(err.Error())
	} else {
		for detailArticle.Next() {
			err = detailArticle.Scan(&response.ArticleDescription)
			if err != nil {
				response.OutError = util.DefaultOutErrorFailed
				response.OutMessage = err.Error()
				log.Error(err.Error())
			}
		}
	}
	return response
}
