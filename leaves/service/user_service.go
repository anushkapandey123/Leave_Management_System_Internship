package service

import (
	"golang.org/x/crypto/bcrypt"
	"main.go/leaves/model"
	"main.go/leaves/repository"
)

type UserService struct {
	UserRepo *repository.UserRepository
}

func NewUserService(userRepo *repository.UserRepository) *UserService {
	return &UserService{UserRepo: userRepo}
}

func (service *UserService) RegisterUser(user model.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
	return service.UserRepo.CreateUser(user)
}

func (service *UserService) AuthenticateUser(email, password string) (model.User, error) {
	user, err := service.UserRepo.GetUserByEmail(email)
	if err != nil {
		return user, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return user, err
	}
	return user, nil
}
