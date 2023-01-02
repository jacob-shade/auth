package database

import (
	"go-sessions-authentication/model"
)

//***********************************CREATE***********************************//

// UserCreate creates a database entry of the given user.
func UserCreate(user *model.User) error {
	return db.Create(&user).Error
}

//***********************************QUERY************************************//

// UserById returns the user with the given id from the database.
// If no user is found, an empty user will be returned.
func UserById(id string) (user model.User, err error) {
	err = db.First(&user, "id = ?", id).Error

	return
}

// UserByEmail returns the user with the given email from the database.
// If no user is found, an empty user will be returned.
func UserByEmail(email string) (user model.User, err error) {
	err = db.First(&user, "email = ?", email).Error

	return
}

// UserByUsername returns the user with the given username from the database.
// If no user is found, an empty user will be returned.
func UserByUsername(username string) (user model.User, err error) {
	err = db.First(&user, "username = ?", username).Error

	return
}

//***********************************UPDATE***********************************//

//***********************************DELETE***********************************//
