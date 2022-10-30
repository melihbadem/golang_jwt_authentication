package services

import (
	"golang.org/x/crypto/bcrypt"
	"jwt-authentication/database"
	"jwt-authentication/models"
)

type UserService struct {
}

func (userService *UserService) GetUser(id interface{}) models.User {
	var user models.User
	database.DBConn.Where("id = ?", id).First(&user)
	return user
}

func (userService *UserService) CreateUser(name string, email string, password string) models.User {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 14)

	user := models.User{
		Name:     name,
		Email:    email,
		Password: hashedPassword,
	}

	database.DBConn.Create(&user)

	return user
}
