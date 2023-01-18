package answer

import (
	"github.com/gin-gonic/gin"
	"main/app/internal/service"
	"net/http"
	"strconv"
)

type ChangeApi struct{}

var insChange = ChangeApi{}

func (a *ChangeApi) CreateAnswer(c *gin.Context) {
	err := service.Answer().Answer().CreateAnswer(c)
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
		"msg":  "Create Answer Successfully",
		"ok":   true,
	})
}

func (a *ChangeApi) DeleteAnswerByID(c *gin.Context) {
	answerIds := c.PostForm("answer_post_id")
	answerId, err := strconv.ParseUint(answerIds, 10, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  "Can't transfer isStr to id",
			"ok":   false,
		})
		return
	}
	err = service.Answer().Answer().DeleteAnswerByID(answerId, c)
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
		"msg":  "Create Answer Successfully",
		"ok":   true,
	})
}

func (a *ChangeApi) UpdateAnswerByID(c *gin.Context) {
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
		case "invalid answer_post_id":
			c.JSON(http.StatusBadRequest, gin.H{
				"code": http.StatusBadRequest,
				"msg":  err.Error(),
				"ok":   false,
			})
		}
		return
	}
	err = service.Answer().Answer().UpdateAnswerByID(answerId, c)
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
		"msg":  "Create Answer Successfully",
		"ok":   true,
	})
}
