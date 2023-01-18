package router

import (
	"github.com/gin-gonic/gin"
	"main/app/api"
)

type ArticleRouter struct{}

func (r *ArticleRouter) InitArticleShowRouter(router *gin.RouterGroup) gin.IRouter {
	articleRouter := router.Group("/article")
	articleApi := api.Article()
	{
		articleRouter.GET("/article_list", articleApi.Show().GetArticleList)
		articleRouter.POST("/article_search", articleApi.Show().GetArticleByID)
	}
	return articleRouter
}

func (r *ArticleRouter) InitArticleChangeRouter(router *gin.RouterGroup) gin.IRouter {
	articleRouter := router.Group("/article")
	articleApi := api.Article()
	{
		articleRouter.GET("/article_create", articleApi.Change().CreateArticle)
		articleRouter.POST("/article_delete", articleApi.Change().DeleteArticleByID)
		articleRouter.POST("/article_update", articleApi.Change().UpdateArticleByID)
	}
	return articleRouter
}
