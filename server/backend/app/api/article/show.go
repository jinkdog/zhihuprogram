package article

import (
	"github.com/gin-gonic/gin"
	"main/app/internal/service"
	"net/http"
	"strconv"
)

type ShowApi struct{}

var insShow = ShowApi{}

func (a *ShowApi) GetArticleList(c *gin.Context) {
	//question := c.PostForm("question_community")
	_, err := service.Article().Article().GetArticleList(c)
	if err != nil {
		switch err.Error() {
		case "internal error":
			c.JSON(http.StatusInternalServerError, gin.H{
				"code": http.StatusInternalServerError,
				"msg":  err.Error(),
				"ok":   false,
			})
		case "no answer in the db":
			c.JSON(http.StatusBadRequest, gin.H{
				"code": http.StatusBadRequest,
				"msg":  err.Error(),
				"ok":   false,
			})

		}
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "Get Article List Successfully",
		"ok":   true,
	})
}

func (a *ShowApi) GetArticleByID(c *gin.Context) {
	articleIdStr := c.PostForm("article_id")
	articleId, err := strconv.ParseUint(articleIdStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  "Can't transfer isStr to id",
			"ok":   false,
		})
		return
	}
	err = service.Article().Article().GetArticleById(articleId, c) //顺序不要搞错了
	if err != nil {
		switch err.Error() {
		case "internal error":
			c.JSON(http.StatusInternalServerError, gin.H{
				"code": http.StatusInternalServerError,
				"msg":  err.Error(),
				"ok":   false,
			})
		case "no article_post in the db":
			c.JSON(http.StatusBadRequest, gin.H{
				"code": http.StatusBadRequest,
				"msg":  err.Error(),
				"ok":   false,
			})
		}
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "Get Article By Id Successfully",
		"ok":   false,
	})
}
