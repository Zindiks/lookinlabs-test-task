package main

import (
	"github.com/Zindiks/lookinlabs-test-task/configs"
	"github.com/Zindiks/lookinlabs-test-task/controller"
	"github.com/Zindiks/lookinlabs-test-task/middleware"
	"github.com/Zindiks/lookinlabs-test-task/repository"
	"github.com/Zindiks/lookinlabs-test-task/service"
	"github.com/charmbracelet/log"
	"github.com/gin-gonic/gin"
)

func main() {
	configs := configs.Configs()

	// Initialize database
	db, err := repository.DB(*configs)
	if err != nil {
		log.Fatal(err)
	} else {

		log.Info("Successfully connected to the database")
	}

	userService := service.NewUserService(db)
	userController := controller.NewUserController(userService)

	gin.ForceConsoleColor()
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	middleware.SetupRoutes(r, userController)

	PORT := configs.App.Port
	if err := r.Run(":" + PORT); err != nil {
		log.Fatalf("Error starting server: %s", err)
	}
}
