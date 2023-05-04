package storage

import (
	"github.com/thiccpan/go-online-lawyer/entities"
	"gorm.io/gorm"
)

type UserStorer interface {
	Add(data entities.User) (entities.User, error)
	GetByEmail(email string) (entities.User, error)
}

type User struct {
	DB *gorm.DB
}

func NewUserStorer(DB *gorm.DB) *User {
	return &User{
		DB: DB,
	}
}

func (u *User) Add(data entities.User) (entities.User, error) {
	err := u.DB.Create(&data).Error
	return data, err
}

func (u *User) GetByEmail(email string) (entities.User, error) {
	user := entities.User{}
	err := u.DB.Where("email = ?", email).First(&user).Error
	return user, err
}

func (u *User) GetAll() ([]entities.User, error) {
	var users []entities.User
	res := u.DB.Find(&users)
	if res.Error != nil {
		return nil, res.Error
	}
	return users, nil
}