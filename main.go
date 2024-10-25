package main

import (
	"net/http"

	"github.com/Zindiks/lookinlabs-test-task/configs"
	"github.com/charmbracelet/log"
	"github.com/gin-gonic/gin"
)

func main() {
	gin.ForceConsoleColor()
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	appConfig := configs.Configs()
	port := appConfig.App.Port

	err := r.Run(":" + port)
	if err != nil {
		log.Fatalf("Error starting server: %s", err)
	}

}
