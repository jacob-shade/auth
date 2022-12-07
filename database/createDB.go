package database

import "go-sessions-authentication/models"

func CreateUser(user *models.User) {
	DB.Create(&user)
}
