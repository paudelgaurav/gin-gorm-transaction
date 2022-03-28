package route

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/paudelgaurav/gin-gorm-transaction/controller"
	"github.com/paudelgaurav/gin-gorm-transaction/middleware"
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
	users.GET("/", userController.GetAllUsers)
	users.GET("/:id", userController.GetUser)

	httpRouter.POST("/transer-money", middleware.DBTransactionMiddleware(db), userController.TranserMoney)

	httpRouter.Run()
}
