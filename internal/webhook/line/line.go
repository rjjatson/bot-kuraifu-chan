package line

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/linebot"
)

const tokenGrantEndpoint = "https://api.line.me/v2/oauth/accessToken"

// Client wraps linebot client
type Client struct {
	bot *linebot.Client
}

// New creates new linebot
func New(channelID, channelSecret string) *Client {
	accessToken, _, err := getAccessToken(channelID, channelSecret)
	if err != nil {
		log.Fatal("unable to get access token ", err.Error())
	}

	bot, err := linebot.New(channelSecret, accessToken)
	if err != nil {
		log.Fatal("unable to create line client ", err.Error())
	}

	return &Client{
		bot: bot,
	}
}

func getAccessToken(channelID, channelSecret string) (accessToken string, expirationTime time.Duration, err error) {
	httpClient := http.Client{}
	resp, err := httpClient.PostForm(tokenGrantEndpoint, url.Values{
		"grant_type":    []string{"client_credentials"},
		"client_id":     []string{channelID},
		"client_secret": []string{channelSecret},
	})
	if err != nil {
		return
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	type tokenGrantResponse struct {
		AccessToken string `json:"access_token"`
		ExpiresIn   int64  `json:"expires_in"`
	}
	var token tokenGrantResponse
	_ = json.Unmarshal(b, &token)

	return token.AccessToken, time.Duration(token.ExpiresIn) * time.Second, nil
}

// HandleWebHook handles line webhook
func (lc *Client) HandleWebHook(c *gin.Context) {
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
