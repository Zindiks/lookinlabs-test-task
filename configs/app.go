package configs

import (

	"os"

	"github.com/charmbracelet/log"
)

type AppConfig struct {
	Port string
}

func LoadAppConfig() *AppConfig {
	err := LoadEnv()
	if err != nil {
		log.Fatal("Failed to load env")
	}

	return &AppConfig{
		Port: os.Getenv("API_PORT"),
	}
}
