package database

import "go-sessions-authentication/models"

func GetUser(id string) models.User {
	var user models.User

	DB.First(&user, "id = ?", id) // find user with id

	return user
}
