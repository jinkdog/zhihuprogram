package comment

import (
	"github.com/gin-gonic/gin"
	"main/app/internal/service"
	"net/http"
	"strconv"
)

type ShowApi struct{}

var insShow = ShowApi{}

func (a *ShowApi) GetCommentListByID(c *gin.Context) {
	commentIdStr := c.PostForm("Comment_id")
	commentId, err := strconv.ParseUint(commentIdStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  "Can't transfer isStr to id",
			"ok":   false,
		})
		return
	}
	err = service.Comment().Comment().GetCommentListById(commentId, c)
	if err != nil {
		switch err.Error() {
		case "internal error":
			c.JSON(http.StatusInternalServerError, gin.H{
				"code": http.StatusInternalServerError,
				"msg":  err.Error(),
				"ok":   false,
			})
		case "no comment in the db":
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
		"msg":  "Get Comment By Id Successfully",
		"ok":   false,
	})
}
