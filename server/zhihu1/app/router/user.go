package router

import (
	"github.com/gin-gonic/gin"
	"main/app/api"
)

type UserRouter struct{} //将以下两个函数绑定为UserRouter的两个方法

func (r *UserRouter) InitUserSignRouter(router *gin.RouterGroup) gin.IRouter { //创建登录注册的路由
	userRouter := router.Group("/user")
	userApi := api.User()
	{
		userRouter.POST("/register", userApi.Sign().Register) //调用Register中间件
		userRouter.POST("/login", userApi.Sign().Login)       //调用登陆中间件
	}

	return userRouter
}

func (r *UserRouter) InitUserInfoRouter(router *gin.RouterGroup) gin.IRoutes { //创建实现用户其他功能的路由
	userRouter := router.Group("/user")

	return userRouter
}
