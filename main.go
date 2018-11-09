package main

import (
	"log"
	"myApp/tradingApp/config"
	"myApp/tradingApp/utils"
)

func main() {
	utils.LoggingSettings(config.Config.LogFile)
	log.Println("test")
}
