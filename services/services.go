package services

import (
	"go-sessions-authentication/database"
	"go-sessions-authentication/model"
)

// Checks to see if the given email is in the user database.
// Params: email string of the user
// user *models.User to be updated with found user, if any
// Returns: true if email found, false otherwise
func IsEmailInUserDB(email string, user *model.User) bool {
	*user = database.User(email)

	return user.Id != 0
}

func RegisterUser(user *model.User) error {
	return database.UserCreate(user)
}
