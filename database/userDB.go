package database

import (
	"go-sessions-authentication/model"
)

//***********************************CREATE***********************************//

func UserCreate(user *model.User) error {
	return DB.Create(&user).Error
}

//************************************READ************************************//

// Gets the user with the given id.
// Params: email string of the user
// Returns: user with the given email, if any, otherwise empty user
func UserById(id string) (model.User, error) {
	var user model.User

	gorm := DB.First(&user, "id = ?", id)

	return user, gorm.Error
}

// Gets the user with the given email.
// Params: email string of the user
// Returns: user with the given email, if any, otherwise empty user
func UserByEmail(email string, user *model.User) error {
	return DB.First(&user, "email = ?", email).Error
}

//***********************************UPDATE***********************************//

//***********************************DELETE***********************************//
