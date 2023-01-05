package question_community

import (
	"github.com/gin-gonic/gin"
	"main/app/internal/service"
	"net/http"
	"strconv"
)

type ChangeApi struct{}

var insChange = ChangeApi{}

func (a *ChangeApi) CreateQuestion(c *gin.Context) {
	err := service.Question().Question().CreateQuestion(c)
	if err != nil {
		switch err.Error() {
		case "insert mysal record failed":
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
		"msg":  "Create Question Successfully",
		"ok":   true,
	})
}

func (a *ChangeApi) DeleteQuestionByID(c *gin.Context) {
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
	err = service.Question().Question().DeleteQuestionByID(questionId, c)
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
		"msg":  "Create Question Successfully",
		"ok":   true,
	})
}

func (a *ChangeApi) UpdateQuestionByID(c *gin.Context) {
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
	err = service.Question().Question().UpdateQuestionByID(questionId, c)
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
		"msg":  "Create Question Successfully",
		"ok":   true,
	})
}
