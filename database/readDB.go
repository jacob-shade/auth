package database

import "go-sessions-authentication/model"

// Gets the user with the given email.
// Params: email string of the user
// Returns: user with the given email, if any, otherwise empty user
func User(email string) model.User {
	var user model.User

	DB.First(&user, "email = ?", email)

	return user
}
