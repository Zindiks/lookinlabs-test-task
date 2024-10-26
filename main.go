package main

import (
    "github.com/Zindiks/lookinlabs-test-task/configs"
    "github.com/Zindiks/lookinlabs-test-task/controller"
    "github.com/Zindiks/lookinlabs-test-task/repository"
    "github.com/Zindiks/lookinlabs-test-task/router"
    "github.com/Zindiks/lookinlabs-test-task/service"
    "github.com/charmbracelet/log"
    "github.com/gin-gonic/gin"
)

func main() {
    // Load configs
    configs := configs.Configs()

    // Initialize database
    db, err := repository.DB(*configs)
    if err != nil {
        log.Fatal(err)
    }
    log.Info("Successfully connected to the database")

    // Initialize service and controller
    userService := service.NewUserService(db)
    userController := controller.NewUserController(userService)

    // Setup Gin
    gin.ForceConsoleColor()
    r := gin.Default()

    // Health check route
    r.GET("/ping", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "pong",
        })
    })

    // Setup routes
    router.SetupRoutes(r, userController)

	
    PORT := configs.App.Port
    if err := r.Run(":" + PORT); err != nil {
        log.Fatalf("Error starting server: %s", err)
    }
}