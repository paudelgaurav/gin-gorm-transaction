package service

import (
	"github.com/paudelgaurav/gin-gorm-transaction/model"
	"github.com/paudelgaurav/gin-gorm-transaction/repository"
)

type UserService interface {
	Save(model.User) (model.User, error)
}

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(r repository.UserRepository) UserService {
	return userService{
		userRepository: r,
	}
}

func (u userService) Save(user model.User) (model.User, error) {
	return u.userRepository.Save(user)
}
