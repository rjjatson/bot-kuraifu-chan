package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/kelseyhightower/envconfig"
	"github.com/rjjatson/bot-kuraifu-chan/api/webservice"
	"github.com/rjjatson/bot-kuraifu-chan/config"
	"github.com/rjjatson/bot-kuraifu-chan/internal/webhook/line"
)

func main() {
	fmt.Printf("bot-kuraifu-chan is running...")

	var cfg config.Configuration
	err := envconfig.Process("kuraifu", &cfg)
	if err != nil {
		log.Fatal("unable to process env variable ", err.Error())
	}

	cf, _ := json.MarshalIndent(cfg, "", "    ")
	fmt.Printf("config : \n%v\n", string(cf))

	linClient := line.New(cfg.LineChannelID, cfg.LineChannelSecret)

	r := gin.Default()
	webservice.SetupRoute(cfg.BasePath,
		linClient,
		r)
	err = r.Run(":" + cfg.BasePort)
	if err != nil {
		fmt.Printf("bot-kuraifu-chan was stopped...")
		log.Fatal(err.Error())
	}
}
