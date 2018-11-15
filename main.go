package main

import (
	"myApp/tradingApp/App/controllers"
	"myApp/tradingApp/config"
	"myApp/tradingApp/utils"
)

func main() {
	utils.LoggingSettings(config.Config.LogFile)
	controllers.StreamIngestionData()
	controllers.StartWebServer()
}
