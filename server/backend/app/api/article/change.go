package article

import (
	"github.com/gin-gonic/gin"
	"main/app/internal/service"
	"net/http"
	"strconv"
)

type ChangeApi struct{}

var insChange = ChangeApi{}

func (a *ChangeApi) CreateArticle(c *gin.Context) {
	err := service.Article().Article().CreateArticle(c)
	if err != nil {
		switch err.Error() {
		case "insert mysql record failed":
			c.JSON(http.StatusInternalServerError, gin.H{
				"code": http.StatusInternalServerError,
				"msg":  err.Error(),
				"ok":   false,
			})

		}
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "Create Article Successfully",
		"ok":   true,
	})
}

func (a *ChangeApi) DeleteArticleByID(c *gin.Context) {
	articleIds := c.PostForm("article_id")
	articleId, err := strconv.ParseUint(articleIds, 10, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  "Can't transfer isStr to id",
			"ok":   false,
		})
		return
	}
	err = service.Article().Article().DeleteArticleByID(articleId, c)
	if err != nil {
		switch err.Error() {
		case "delete mysql record failed":
			c.JSON(http.StatusInternalServerError, gin.H{
				"code": http.StatusInternalServerError,
				"msg":  err.Error(),
				"ok":   false,
			})
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "Create Article Successfully",
		"ok":   true,
	})
}

func (a *ChangeApi) UpdateArticleByID(c *gin.Context) {
	articleIdStr := c.PostForm("answer_post_id")
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
		case "invalid answer_post_id":
			c.JSON(http.StatusBadRequest, gin.H{
				"code": http.StatusBadRequest,
				"msg":  err.Error(),
				"ok":   false,
			})
		}
		return
	}
	err = service.Article().Article().UpdateArticleByID(articleId, c)
	if err != nil {
		switch err.Error() {
		case "update mysql record failed":
			c.JSON(http.StatusInternalServerError, gin.H{
				"code": http.StatusInternalServerError,
				"msg":  err.Error(),
				"ok":   false,
			})
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "Create Article Successfully",
		"ok":   true,
	})
}
