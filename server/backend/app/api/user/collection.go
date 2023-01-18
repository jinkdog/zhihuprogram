package user

import (
	"github.com/gin-gonic/gin"
	"main/app/internal/service"
	"net/http"
	"strconv"
)

type CollectApi struct{}

var insCollect = CollectApi{}

func (a *CollectApi) CreateCollect(c *gin.Context) {
	err := service.User().Collect().CreateCollection(c)
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
		"msg":  "Create Question Successfully",
		"ok":   true,
	})
}

func (a *CollectApi) DeleteCollectByID(c *gin.Context) {
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
