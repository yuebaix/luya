package admin

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func StartAdminServer() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "farewell")
	})
	router.Run()
}
