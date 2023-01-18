package answer

import (
	"github.com/gin-gonic/gin"
	"main/app/internal/service"
	"net/http"
	"strconv"
)

type ShowApi struct{}

var insShow = ShowApi{}

func (a *ShowApi) GetAnswerList(c *gin.Context) {
	//question := c.PostForm("question_community")
	_, err := service.Answer().Answer().GetAnswerList(c)
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
		"msg":  "Get Answer List Successfully",
		"ok":   true,
	})
}

func (a *ShowApi) GetAnswerByID(c *gin.Context) {
	answerIdStr := c.PostForm("answer_post_id")
	answerId, err := strconv.ParseUint(answerIdStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  "Can't transfer isStr to id",
			"ok":   false,
		})
		return
	}
	err = service.Answer().Answer().GetAnswerById(answerId, c) //顺序不要搞错了
	if err != nil {
		switch err.Error() {
		case "internal error":
			c.JSON(http.StatusInternalServerError, gin.H{
				"code": http.StatusInternalServerError,
				"msg":  err.Error(),
				"ok":   false,
			})
		case "no answer_post in the db":
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
		"msg":  "Get Answer By Id Successfully",
		"ok":   false,
	})
}
