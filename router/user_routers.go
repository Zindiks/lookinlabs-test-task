package router

import (
    "github.com/Zindiks/lookinlabs-test-task/controller"
    "github.com/gin-gonic/gin"
)

func setupUserRoutes(api *gin.RouterGroup, userController *controller.UserController) {
    users := api.Group("/users")
    {
        users.POST("", userController.CreateUser)
        users.GET("", userController.GetUsers)
        users.GET("/:id", userController.GetUser)
        users.PATCH("/:id", userController.UpdateUser)
    }
}