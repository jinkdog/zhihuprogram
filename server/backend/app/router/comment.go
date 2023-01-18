package router

import (
	"github.com/gin-gonic/gin"
	"main/app/api"
)

type CommentRouter struct{}

func (r *CommentRouter) InitCommentShowRouter(router *gin.RouterGroup) gin.IRouter {
	commentRouter := router.Group("/comment")
	commentApi := api.Comment()
	{
		commentRouter.POST("/comment_list", commentApi.Show().GetCommentListByID)
	}
	return commentRouter
}

func (r *CommentRouter) InitCommentChangeRouter(router *gin.RouterGroup) gin.IRouter {
	commentRouter := router.Group("/comment")
	commentApi := api.Comment()
	{
		commentRouter.POST("/comment_create", commentApi.Change().CreateComment)
	}
	return commentRouter
}
