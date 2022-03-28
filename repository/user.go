package repository

import (
	"errors"
	"log"

	"github.com/paudelgaurav/gin-gorm-transaction/model"
	"gorm.io/gorm"
)

type userRepository struct {
	DB *gorm.DB
}

type UserRepository interface {
	Save(model.User) (model.User, error)
	GetAllUsers() ([]model.User, error)
	GetUser(int) (model.User, error)
	WithTrx(*gorm.DB) userRepository
	IncrementMoney(uint, float64) error
	DecrementMoney(uint, float64) error
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

func (u userRepository) GetAllUsers() (users []model.User, err error) {
	log.Print("[User Repository]... Get All users")
	err = u.DB.Find(&users).Error
	return users, err
}

func (u userRepository) GetUser(userID int) (user model.User, err error) {
	log.Print("[User Repository].. Get User")
	err = u.DB.First(&user, uint(userID)).Error
	return user, err
}

func (u userRepository) WithTrx(trxHandle *gorm.DB) userRepository {
	if trxHandle == nil {
		log.Print("Transaction databse not found")
		return u
	}
	u.DB = trxHandle
	return u
}

func (u userRepository) IncrementMoney(reciever uint, amount float64) error {
	log.Print("[UserRepository]...Increment Money")
	return u.DB.Model(&model.User{}).Where("id=?", reciever).Update("wallet", gorm.Expr("wallet + ?", amount)).Error
}

func (u userRepository) DecrementMoney(giver uint, amount float64) error {
	log.Print("[UserRepository]...Decrement Money")
	// return u.DB.Model(&model.User{}).Where("id=?", giver).Update("wallet", gorm.Expr("wallet - ?", amount)).Error
	return errors.New("test Error")
}

func (u userRepository) Migrate() error {
	log.Print("[UserRepository]..migrate")
	return u.DB.AutoMigrate(&model.User{})
}
