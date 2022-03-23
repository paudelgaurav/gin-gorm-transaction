package route

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/paudelgaurav/gin-gorm-transaction/service"
	"github.com/paudelgaurav/go-repo/repository"
	"gorm.io/gorm"
)

func SetUpRoute(db *gorm.DB) {
	httpRouter := gin.Default()

	userRepository := repository.NewUserRepository(db)

	if err := userRepository.Migrate(); err != nil {
		log.Fatal("User migrate error", err)
	}

	userService := service.NewUserService(userRepository)

}
