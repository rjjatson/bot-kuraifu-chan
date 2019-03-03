package webservice

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rjjatson/bot-kuraifu-chan/pkg/webhook/line"
)

// SetupRoute add routes to gin router
func SetupRoute(basePath string,
	lineClient *line.LineClient,
	router *gin.Engine) {

	router.GET(basePath+"/ping", func(c *gin.Context) {
		c.Writer.WriteHeader(http.StatusOK)
	})

	router.POST(basePath+"/webhook/line", lineClient.HandleWebHook)
}
