package configs

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type dbConfig struct {
	host string
	port string
	user string
	pass string
	name string
}

func LoadDBConfig() *dbConfig {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Unable to load env file")
	}

	config := &dbConfig{
		host: os.Getenv("DB_HOST"),
		port: os.Getenv("DB_PORT"),
		user: os.Getenv("DB_USER"),
		pass: os.Getenv("DB_PASS"),
		name: os.Getenv("DB_NAME"),
	}

	return config
}

func DSN(c dbConfig) string {
	return fmt.Sprintf("user=%s password=%s dbname=%s port=%s sslmode=disable", c.user, c.pass, c.name, c.port)
}
