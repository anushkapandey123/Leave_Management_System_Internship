package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name        string `json:"name"`
	Email       string `json:"email" gorm:"unique"`
	Password    string `json:"password"`
	PhoneNumber int    `json:"phone_number"`
}

func NewUser(name string, email string, password string, phoneNumber int) User {
	return User{
		Name:        name,
		Email:       email,
		Password:    password,
		PhoneNumber: phoneNumber,
	}
}
