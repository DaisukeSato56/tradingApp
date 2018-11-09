package main

import (
	"fmt"
	"myApp/tradingApp/config"
)

func main() {
	fmt.Println(config.Config.ApiKey)
	fmt.Println(config.Config.ApiSecret)
}
