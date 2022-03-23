package service

import (
	"github.com/paudelgaurav/gin-gorm-transaction/model"
	"github.com/paudelgaurav/gin-gorm-transaction/repository"
)

type UserService interface {
	Save(model.User) (model.User, error)
	GetAllUsers() ([]model.User, error)
	GetUser(int) (model.User, error)
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

func (u userService) GetAllUsers() ([]model.User, error) {
	return u.userRepository.GetAllUsers()
}

func (u userService) GetUser(userID int) (model.User, error) {
	return u.userRepository.GetUser(userID)
}
