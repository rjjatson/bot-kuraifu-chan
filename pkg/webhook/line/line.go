package line

import (
	"net/http"

	"github.com/line/line-bot-sdk-go/linebot"

	"github.com/gin-gonic/gin"
)

// LineClient wraps linebot client
type LineClient struct {
	bot *linebot.Client
}

// New creates new linebot
func New(bot *linebot.Client) *LineClient {
	return &LineClient{
		bot: bot,
	}
}

// HandleWebHook handles line webhook
func (lc *LineClient) HandleWebHook(c *gin.Context) {
	c.Status(http.StatusOK)
}
