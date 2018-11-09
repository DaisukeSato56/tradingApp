package config

import (
	"log"
	"os"

	ini "gopkg.in/ini.v1"
)

type ConfigList struct {
	ApiKey    string
	ApiSecret string
	LogFile   string
}

var Config ConfigList

func init() {
	cfg, err := ini.Load("config.ini")
	if err != nil {
		log.Printf("faild to load file %v", err)
		os.Exit(1)
	}

	Config = ConfigList{
		ApiKey:    cfg.Section("bitflyer").Key("api_key").String(),
		ApiSecret: cfg.Section("bitflyer").Key("api_secret").String(),
		LogFile:   cfg.Section("tradingApp").Key("log_file").String(),
	}
}
