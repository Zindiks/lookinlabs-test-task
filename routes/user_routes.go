package routes

import (
	"github.com/Zindiks/lookinlabs-test-task/controller"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UserRoutes(router *gin.RouterGroup, db *gorm.DB) {

	router.POST("/", func(c *gin.Context) {
		controller.CreateUser(c, db)
	})

	router.GET("/", func(c *gin.Context) {
		controller.GetUsers(c, db)
	})

	router.GET("/:id", func(c *gin.Context) {
		controller.GetUser(c, db)
	})

	router.PATCH("/:id", func(c *gin.Context) {
		controller.UpdateUser(c, db)
	})

}
