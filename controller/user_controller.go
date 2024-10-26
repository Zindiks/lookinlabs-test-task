package controller

import (
	"net/http"

	"github.com/Zindiks/lookinlabs-test-task/model"
	"github.com/Zindiks/lookinlabs-test-task/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserController struct {
    userOperations service.UserService
}

func NewUserController(userOperations service.UserService) *UserController {
    return &UserController{
        userOperations: userOperations,
    }
}

func (ctrl *UserController) CreateUser(ctx *gin.Context) {
    var user model.User

    if err := ctx.ShouldBindJSON(&user); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := ctrl.userOperations.CreateUser(&user); err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save user"})
        return
    }

    ctx.JSON(http.StatusOK, user)
}

func (ctrl *UserController) GetUsers(ctx *gin.Context) {
    users, err := ctrl.userOperations.GetUsers()
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch users"})
        return
    }

    ctx.JSON(http.StatusOK, users)
}

func (ctrl *UserController) GetUser(ctx *gin.Context) {
    id := ctx.Param("id")

    user, err := ctrl.userOperations.GetUserByID(id)
    if err != nil {
        if err == gorm.ErrRecordNotFound {
            ctx.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
            return
        }
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch user"})
        return
    }

    ctx.JSON(http.StatusOK, user)
}

func (ctrl *UserController) UpdateUser(ctx *gin.Context) {
    id := ctx.Param("id")

    user, err := ctrl.userOperations.GetUserByID(id)
    if err != nil {
        if err == gorm.ErrRecordNotFound {
            ctx.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
            return
        }
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch user"})
        return
    }

    if err := ctx.ShouldBindJSON(user); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := ctrl.userOperations.UpdateUser(user); err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
        return
    }

    ctx.JSON(http.StatusOK, user)
}