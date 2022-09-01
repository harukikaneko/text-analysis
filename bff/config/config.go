package config

import (
	"fmt"
	"log"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	TextAnalysisEndpoint string `required:"true" split_words:"true"`
}

var config Config

func InitConfig() {
	if err := envconfig.Process("", &config); err != nil {
		log.Fatalf("[ERROR] Failed to process env: %s", err.Error())
	}
	fmt.Println("Config.RndManagementEndpoint:", config.TextAnalysisEndpoint)
}

func GetConfig() Config {
	return config
}
