package database

import (
	"go-sessions-authentication/models"
)

func CreateUser(user *models.User) error {
	return DB.Create(&user).Error
}
