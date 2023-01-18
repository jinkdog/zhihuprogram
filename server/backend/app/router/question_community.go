package router

import (
	"github.com/gin-gonic/gin"
	"main/app/api"
)

type QuestionRouter struct{}

func (r *QuestionRouter) InitQuestionShowRouter(router *gin.RouterGroup) gin.IRouter {
	questionRouter := router.Group("/question")
	questionApi := api.Question()
	{
		questionRouter.GET("/question_list", questionApi.Show().GetQuestionList)
		questionRouter.POST("/question_search", questionApi.Show().GetQuestionByID)
	}
	return questionRouter
}

func (r *QuestionRouter) InitQuestionChangeRouter(router *gin.RouterGroup) gin.IRouter {
	questionRouter := router.Group("/question")
	questionApi := api.Question()
	{
		questionRouter.GET("/question_create", questionApi.Change().CreateQuestion)
		questionRouter.POST("/question_delete", questionApi.Change().DeleteQuestionByID)
		questionRouter.POST("/question_update", questionApi.Change().UpdateQuestionByID)
	}
	return questionRouter
}
