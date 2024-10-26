package main

import (
	"net/http"

	"github.com/Zindiks/lookinlabs-test-task/configs"
	"github.com/Zindiks/lookinlabs-test-task/repository"
	"github.com/Zindiks/lookinlabs-test-task/routes"
	"github.com/charmbracelet/log"
	"github.com/gin-gonic/gin"
)

func main() {

	configs := configs.Configs()

	db, err := repository.DB(*configs)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Info("Successfully connected to the database")
	}

	gin.ForceConsoleColor()
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	api := r.Group("/api/v1")
	userGroup := api.Group("/users")
	routes.UserRoutes(userGroup, db)



	port := configs.App.Port

	err = r.Run(":" + port)
	if err != nil {
		log.Fatalf("Error starting server: %s", err)
	}

}
