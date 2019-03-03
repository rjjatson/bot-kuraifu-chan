package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/rjjatson/bot-kuraifu-chan/api/webservice"
)

func main() {
	fmt.Printf("bot-kuraifu-chan is running...")
	router := gin.Default()

	webservice.SetupRoute(router)

	err := router.Run(":8080")
	if err != nil {
		fmt.Printf("bot-kuraifu-chan was stopped...")
		panic(err)
	}
}
