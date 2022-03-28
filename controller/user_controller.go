package controller

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/paudelgaurav/gin-gorm-transaction/model"
	"github.com/paudelgaurav/gin-gorm-transaction/service"
	"gorm.io/gorm"
)

type UserController interface {
	AddUser(*gin.Context)
	GetAllUsers(*gin.Context)
	GetUser(*gin.Context)
	TranserMoney(*gin.Context)
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

func (u userController) GetAllUsers(ctx *gin.Context) {
	users, error := u.userService.GetAllUsers()
	if error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": error})
	}
	ctx.JSON(http.StatusOK, gin.H{"data": users})
}

func (u userController) GetUser(ctx *gin.Context) {
	userID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, err := u.userService.GetUser(userID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": user})
}

func (u userController) TranserMoney(ctx *gin.Context) {
	log.Print("[UserController]...get all Users")

	txHandle := ctx.MustGet("db_trx").(*gorm.DB)
	var moneytransfer model.MoneyTransfer

	if err := ctx.ShouldBindJSON(&moneytransfer); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"erorr": err.Error()})
		return
	}

	if err := u.userService.WithTrx(txHandle).IncrementMoney(moneytransfer.Receiver, moneytransfer.Amount); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Error while incrementing money"})
		txHandle.Rollback()
		return
	}

	if err := u.userService.WithTrx(txHandle).DecrementMoney(moneytransfer.Giver, moneytransfer.Amount); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Error while decrementing money"})
		txHandle.Rollback()
		return
	}

	if err := txHandle.Commit().Error; err != nil {
		log.Print("Trx commit error")
	}

	ctx.JSON(http.StatusOK, gin.H{"detail": "Money transfer successful"})

}
