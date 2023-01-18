package comment

import (
	"github.com/gin-gonic/gin"
	"main/app/internal/service"
	"net/http"
)

type ChangeApi struct{}

var insChange = ChangeApi{}

func (a *ChangeApi) CreateComment(c *gin.Context) {
	err := service.Comment().Comment().CreateCommend(c)
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
		"msg":  "Create Comment Successfully",
		"ok":   true,
	})
}
