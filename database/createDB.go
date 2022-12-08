package database

import "go-sessions-authentication/model"

func UserCreate(user *model.User) error {
	return DB.Create(&user).Error
}
