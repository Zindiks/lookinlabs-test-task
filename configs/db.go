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

	config := &DBConfig{
		Host: os.Getenv("DB_HOST"),
		Port: os.Getenv("DB_PORT"),
		User: os.Getenv("DB_USER"),
		Pass: os.Getenv("DB_PASS"),
		Name: os.Getenv("DB_NAME"),
	}

	return config
}

// func DSN(c DBConfig) string {
// 	return fmt.Sprintf("user=%s password=%s dbname=%s port=%s sslmode=disable", c.User, c.Pass, c.Name, c.Port)
// }
