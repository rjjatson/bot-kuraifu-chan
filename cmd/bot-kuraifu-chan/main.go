package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/kelseyhightower/envconfig"
	"github.com/line/line-bot-sdk-go/linebot"
	"github.com/rjjatson/bot-kuraifu-chan/api/webservice"
	"github.com/rjjatson/bot-kuraifu-chan/config"
	"github.com/rjjatson/bot-kuraifu-chan/pkg/webhook/line"
)

func main() {
	fmt.Printf("bot-kuraifu-chan is running...")

	var cfg config.Configuration
	err := envconfig.Process("kuraifu", &cfg)
	if err != nil {
		log.Fatal("unable to process env variable ", err.Error())
	}

	bot, err := linebot.New(cfg.LineChannelSecret, cfg.LineChannelSecret)
	if err != nil {
		//log.Fatal("unable to create line client ", err.Error())
	}

	router := gin.Default()
	webservice.SetupRoute(cfg.BasePath,
		line.New(bot),
		router)

	err = router.Run(":" + cfg.BasePort)
	if err != nil {
		fmt.Printf("bot-kuraifu-chan was stopped...")
		log.Fatal(err.Error())
	}
}
