package repository

import (
	"log"

	"github.com/paudelgaurav/gin-gorm-transaction/model"
	"gorm.io/gorm"
)

type userRepository struct {
	DB *gorm.DB
}

type UserRepository interface {
	Save(model.User) (model.User, error)
	Migrate() error
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return userRepository{
		DB: db,
	}
}

func (u userRepository) Save(user model.User) (model.User, error) {
	log.Print("[UserRepository]...Save")
	err := u.DB.Create(&user).Error
	return user, err
}

func (u userRepository) Migrate() error {
	log.Print("[UserRepository]..migrate")
	return u.DB.AutoMigrate(&model.User{})
}
