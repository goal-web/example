package controllers

import (
	"github.com/goal-web/contracts"
)

func CreateArticle(request contracts.HttpRequest) interface{} {
	return "发布文章" + request.GetString("content")
}
