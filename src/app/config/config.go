package config

import (
	log "app/logging"
	"fmt"
	"github.com/tkanos/gonfig"
)

type DB struct {
	DbUsername string
	DbPassword string
	DbPort     string
	DbHost     string
}

func GetConfig(env string) DB {
	configuration := DB{}

	fileName := fmt.Sprintf("./config/%s_config.yml", env)
	err := gonfig.GetConf(fileName, &configuration)
	if err != nil {
		log.ErrorLogger.Println(err)
		return DB{}
	}

	return configuration
}
