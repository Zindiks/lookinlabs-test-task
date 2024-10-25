package main

import (
	"net/http"

	"github.com/Zindiks/lookinlabs-test-task/configs"
	"github.com/Zindiks/lookinlabs-test-task/repository"
	"github.com/charmbracelet/log"
	"github.com/gin-gonic/gin"
)

func main() {

	configs := configs.Configs()


	db, err := repository.DB(*configs)
	if err != nil {
		log.Fatal(err)
	}

	log.Info(db)

	gin.ForceConsoleColor()
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	port := configs.App.Port

	err2 := r.Run(":" + port)
	if err2 != nil {
		log.Fatalf("Error starting server: %s", err)
	}

}
