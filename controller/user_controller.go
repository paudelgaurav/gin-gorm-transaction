package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/paudelgaurav/gin-gorm-transaction/model"
	"github.com/paudelgaurav/gin-gorm-transaction/service"
)

type UserController interface {
	AddUser(*gin.Context)
}

type userController struct {
	userService service.UserService
}

func NewUserController(s service.UserService) UserController {
	return userController{
		userService: s,
	}
}

func (u userController) AddUser(ctx *gin.Context) {
	var user model.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, error := u.userService.Save(user)
	if error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "User cannot be created"})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"data": user})
}
