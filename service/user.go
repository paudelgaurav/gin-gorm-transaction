package service

import (
	"github.com/paudelgaurav/gin-gorm-transaction/model"
	"github.com/paudelgaurav/gin-gorm-transaction/repository"
	"gorm.io/gorm"
)

type UserService interface {
	Save(model.User) (model.User, error)
	GetAllUsers() ([]model.User, error)
	GetUser(int) (model.User, error)
	WithTrx(*gorm.DB) userService
	IncrementMoney(uint, float64) error
	DecrementMoney(uint, float64) error
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

func (u userService) WithTrx(trxHandle *gorm.DB) userService {
	u.userRepository = u.userRepository.WithTrx(trxHandle)
	return u
}

func (u userService) IncrementMoney(reciever uint, amount float64) error {
	return u.userRepository.IncrementMoney(reciever, amount)
}

func (u userService) DecrementMoney(giver uint, amount float64) error {
	return u.userRepository.DecrementMoney(giver, amount)
}
