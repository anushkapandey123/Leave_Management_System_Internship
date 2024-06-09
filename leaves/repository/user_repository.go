package repository

import (
	"errors"

	"main.go/leaves/model"

	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (repo *UserRepository) CreateUser(user model.User) error {
	return repo.DB.Create(&user).Error
}

func (repo *UserRepository) GetUserByEmail(email string) (model.User, error) {
	var user model.User
	result := repo.DB.Where("email = ?", email).First(&user)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return user, errors.New("user not found")
	}
	return user, result.Error
}
