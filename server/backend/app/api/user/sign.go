package user

import (
	"github.com/gin-gonic/gin"
	g "main/app/global"
	"main/app/internal/model"
	"main/app/internal/service"
	"main/utils/cookie"
	"net/http"
	"strconv"
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
	tokenString, err := service.User().User().GenerateToken(c, userSubject)
	if err != nil {
		switch err.Error() {
		case "generate token failed internal err":
			c.JSON(http.StatusInternalServerError, gin.H{
				"code": http.StatusInternalServerError,
				"msg":  err.Error(),
				"ok":   false,
			})
		case "set redis cache failed internal err":
			c.JSON(http.StatusInternalServerError, gin.H{
				"code": http.StatusInternalServerError,
				"msg":  err.Error(),
				"ok":   false,
			})
		}

	}

	cookieConfig := g.Config.App.Cookie
	cookieWriter := cookie.NewCookieWriter(cookieConfig.Secret, cookie.Option{
		Config: http.Cookie{
			Path:     "/",
			Domain:   cookieConfig.Domain,
			MaxAge:   cookieConfig.MaxAge,
			Secure:   cookieConfig.Secure,
			HttpOnly: cookieConfig.HttpOnly,
			SameSite: cookieConfig.SameSite,
		},
		Ctx: c,
	})

	cookieWriter.Set("x-token", tokenString)

	c.JSON(http.StatusOK, gin.H{
		"code":  http.StatusOK,
		"msg":   "login successfully",
		"token": tokenString,
		"ok":    true,
	})

}
func (a *SignApi) UpdateUser(c *gin.Context) {
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
	NewUserName := c.PostForm("username")
	NewPassWord := c.PostForm("password")
	UserIdStr := c.PostForm("user_id")

	if NewUserName == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "username cannot be null",
			"ok":   false,
		})
		return
	}
	if NewPassWord == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "password cannot be null",
			"ok":   false,
		})
		return
	}

	UserId, err1 := strconv.ParseUint(UserIdStr, 10, 64)
	if err1 != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  "Can't transfer UserIdStr to UserId",
			"ok":   false,
		})
		return
	}

	err = service.User().User().UpdateUserByID(UserId, c)
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
