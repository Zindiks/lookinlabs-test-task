package middleware

import (
    "github.com/Zindiks/lookinlabs-test-task/controller"
    "github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, userController *controller.UserController) {
    api := r.Group("/api/v1")
    setupUserRoutes(api, userController)
}