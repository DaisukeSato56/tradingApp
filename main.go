package main

import (
	"fmt"
	"gotrading_application/110_sma/bitflyer"
	"myApp/tradingApp/config"
	"myApp/tradingApp/utils"
)

func main() {
	utils.LoggingSettings(config.Config.LogFile)
	apiClient := bitflyer.New(config.Config.ApiKey, config.Config.ApiSecret)
	fmt.Println(apiClient.GetBalance())
}
