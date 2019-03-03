package webservice

import (
	"github.com/gin-gonic/gin"
)

// SetupRoute add routes to gin router
func SetupRoute(router *gin.Engine) {
	router.GET("/", func(c *gin.Context) {
		_, _ = c.Writer.Write([]byte("fuwaaaaaa"))
	})
}
