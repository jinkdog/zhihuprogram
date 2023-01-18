package router

import (
	"github.com/gin-gonic/gin"
	"main/app/api"
)

type AnswerRouter struct{}

func (r *AnswerRouter) InitAnswerShowRouter(router *gin.RouterGroup) gin.IRouter {
	answerRouter := router.Group("/answer")
	answerApi := api.Answer()
	{
		answerRouter.GET("/answer_list", answerApi.Show().GetAnswerList)
		answerRouter.POST("/answer_search", answerApi.Show().GetAnswerByID)
	}
	return answerRouter
}

func (r *AnswerRouter) InitAnswerChangeRouter(router *gin.RouterGroup) gin.IRouter {
	answerRouter := router.Group("/answer")
	//answerApi := api.Answer()
	{
		//answerRouter.GET("/answer_create", answerApi.Change().CreateAnswer)
		//answerRouter.POST("/answer_delete", answerApi.Change().DeleteAnswerByID)
		//answerRouter.POST("/answer_update", answerApi.Change().UpdateAnswerByID)
	}
	return answerRouter
}
