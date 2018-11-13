package main

import (
	"fmt"
	"myApp/tradingApp/App/controllers"
	"myApp/tradingApp/App/models"
	"myApp/tradingApp/config"
	"myApp/tradingApp/utils"
)

func main() {
	utils.LoggingSettings(config.Config.LogFile)
	fmt.Println(models.DBConnection)
	controllers.StreamIngestionData()
}
