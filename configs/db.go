package configs

import (
	"os"

	"github.com/charmbracelet/log"
)

type DBConfig struct {
	Host string
	Port string
	User string
	Pass string
	Name string
}

func LoadDBConfig() (*DBConfig) {
	err := LoadEnv()
	if err != nil {
		log.Fatal("Failed to load env")
		
	}

	return &DBConfig{
		Host: os.Getenv("DB_HOST"),
		Port: os.Getenv("DB_PORT"),
		User: os.Getenv("DB_USER"),
		Pass: os.Getenv("DB_PASS"),
		Name: os.Getenv("DB_NAME"),
	}

	
}
