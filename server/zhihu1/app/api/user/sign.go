package user

import (
	"github.com/gin-gonic/gin"
	"main/app/internal/model"
	"main/app/internal/service"
	"net/http"
)

type SignApi struct{}

var insSign = SignApi{} //自定义类型为SignApi结构体//为什么这么做？

func (a *SignApi) Register(c *gin.Context) { //注册逻辑
	username := c.PostForm("username")
	password := c.PostForm("password")

	if username == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "username cannot be null",
			"ok":   false,
		})
		return
	}
	if password == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "password cannot be null",
			"ok":   false,
		})
		return
	}

	err := service.User().User().CheckUserIsExist(c, username)
	if err != nil {
		if err.Error() == "internal err" {
			c.JSON(http.StatusInternalServerError, gin.H{
				"code": http.StatusInternalServerError,
				"msg":  err.Error(),
				"ok":   false,
			})

		} else if err.Error() == "username already exist" {
			c.JSON(http.StatusBadRequest, gin.H{
				"code": http.StatusBadRequest,
				"msg":  "username already exist",
				"ok":   false,
			})
		}
		return
	}

	userSubject := &model.UserSubject{}

	encryptedPassword := service.User().User().EncryptPassword(password)

	userSubject.Username = username
	userSubject.Password = encryptedPassword

	service.User().User().CreateUser(c, userSubject)

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "register successfully",
		"ok":   true,
	})
}

func (a *SignApi) Login(c *gin.Context) { //登录逻辑
	username := c.PostForm("username")
	password := c.PostForm("password")

	if username == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "username cannot be null",
			"ok":   false,
		})
		return
	}
	if password == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "password cannot be null",
			"ok":   false,
		})
		return
	}

	userSubject := &model.UserSubject{
		Username: username,
		Password: service.User().User().EncryptPassword(password),
	}

	err := service.User().User().CheckPassword(c, userSubject)
	if err != nil {
		switch err.Error() {
		case "internal err":
			c.JSON(http.StatusInternalServerError, gin.H{
				"code": http.StatusInternalServerError,
				"msg":  err.Error(),
				"ok":   false,
			})
		case "invalid username or password":
			c.JSON(http.StatusBadRequest, gin.H{
				"code": http.StatusBadRequest,
				"msg":  err.Error(),
				"ok":   false,
			})
		}

		return
	}

}
