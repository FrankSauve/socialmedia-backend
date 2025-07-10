package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Routes(route *gin.Engine) {
	user := route.Group("/api/user")
	user.GET("/", getUser)
}

func getUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"hello": "world",
	})
}
