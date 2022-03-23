package route

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/paudelgaurav/gin-gorm-transaction/controller"
	"github.com/paudelgaurav/gin-gorm-transaction/repository"
	"github.com/paudelgaurav/gin-gorm-transaction/service"
	"gorm.io/gorm"
)

func SetUpRoute(db *gorm.DB) {
	httpRouter := gin.Default()

	userRepository := repository.NewUserRepository(db)

	if err := userRepository.Migrate(); err != nil {
		log.Fatal("User migrate error", err)
	}

	userService := service.NewUserService(userRepository)

	userController := controller.NewUserController(userService)

	users := httpRouter.Group("users")
	users.POST("/", userController.AddUser)

	httpRouter.Run()
}
