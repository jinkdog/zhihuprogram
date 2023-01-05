package router

import (
	"github.com/gin-gonic/gin"
	g "main/app/global"
	"main/app/internal/middleware"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	r.Use(middleware.ZapLogger(g.Logger), middleware.ZapRecovery(g.Logger, true))
	// 使用其他的中间件(跨域,限流...)
	r.Use(middleware.CorsByRules())

	routerGroup := new(Group)

	publicGroup := r.Group("/api")
	{
		routerGroup.InitUserSignRouter(publicGroup)
		routerGroup.InitQuestionShowRouter(publicGroup)
	}

	privateGroup := r.Group("/api")
	privateGroup.Use(middleware.JWTAuthMiddleware())
	{

		routerGroup.InitQuestionChangeRouter(privateGroup)
	}

	g.Logger.Info("initialize routers successfully!")

	return r
}

//publicgroup和privategroup分别是什么意思？
