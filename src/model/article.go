package model

import (
	"database/sql"
	"reflect"
)

type ListArticle struct {
	Token      string `json:"token"`
	ParentId   string `json:"parentId"`
	PageNumber string `json:"pageNumber"`
	RowPerpage string `json:"rowPerpage"`
	Search     string `json:"search"`
}

type ListArticleResponse struct {
	OutData []Article `json:"outData"`
	BaseResponseModel
}

type Article struct {
	ArticleId   string `json:"articleId"`
	ArticleCode string `json:"articleCode"`
	ArticleName string `json:"articleName"`
	ArticlePath string `json:"articlePath"`
}

type DetailArticle struct {
	ArticleCode string `json:"articleCode"`
}

type DetailArticleResponse struct {
	ArticleDescription string `json:"articleDescription"`
	BaseResponseModel
}

// NullString is an alias for sql.NullString data type
type NullString sql.NullString

// Scan implements the Scanner interface for NullString
func (ns *NullString) Scan(value interface{}) error {
	var s sql.NullString
	if err := s.Scan(value); err != nil {
		return err
	}

	// if nil then make Valid false
	if reflect.TypeOf(value) == nil {
		*ns = NullString{s.String, false}
	} else {
		*ns = NullString{s.String, true}
	}

	return nil
}
