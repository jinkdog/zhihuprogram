package question_community

import (
	"github.com/gin-gonic/gin"
	"main/app/internal/service"
	"net/http"
	"strconv"
)

type ShowApi struct{}

var insShow = ShowApi{}

func (a *ShowApi) GetQuestionList(c *gin.Context) {
	//question := c.PostForm("question_community")
	_, err := service.Question().Question().GetQuestionList(c)
	if err != nil {
		switch err.Error() {
		case "internal error":
			c.JSON(http.StatusInternalServerError, gin.H{
				"code": http.StatusInternalServerError,
				"msg":  err.Error(),
				"ok":   false,
			})
		case "no question_community in the db":
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
		"msg":  "Get Question List Successfully",
		"ok":   true,
	})
}

func (a *ShowApi) GetQuestionByID(c *gin.Context) {
	questionIdStr := c.PostForm("question_community_id")
	questionId, err := strconv.ParseUint(questionIdStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  "Can't transfer isStr to id",
			"ok":   false,
		})
		return
	}
	_, err = service.Question().Question().GetQuestionByID(questionId, c) //顺序不要搞错了
	if err != nil {
		switch err.Error() {
		case "internal error":
			c.JSON(http.StatusInternalServerError, gin.H{
				"code": http.StatusInternalServerError,
				"msg":  err.Error(),
				"ok":   false,
			})
		case "invalid question_community_id":
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
		"msg":  "Get Question By Id Successfully",
		"ok":   false,
	})
}
