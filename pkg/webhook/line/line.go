package line

import (
	"github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/linebot"
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
	events, err := lc.bot.ParseRequest(c.Request)
	if err != nil {
		logrus.Error("unable to parse request: ", err)
	}

	for _, evt := range events {
		switch evt.Type {
		case linebot.EventTypeMessage:
			msg := evt.Message.(*linebot.TextMessage)
			logrus.Info("text message received: ", msg.Text)
			_, err = lc.bot.ReplyMessage(evt.ReplyToken, linebot.NewTextMessage(msg.Text+" too")).Do()
			if err != nil {
				logrus.Error("unable to send reply: ", err)
			}
		}
	}
}
