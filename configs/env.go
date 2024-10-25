package configs

import (

	"os"

	"github.com/joho/godotenv"
	"github.com/charmbracelet/log"
)

func LoadEnv() error {
	env := os.Getenv("ENV")
	envFile := ".env"

	if env == "dev" {
		envFile = ".env.dev"
	}

	if err := godotenv.Load(envFile); err != nil {
		log.Fatalf("Failed to load %s file", envFile)
	}

	return nil
}
