package database

import "go-sessions-authentication/model"

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
func UserByEmail(email string) (model.User, error) {
	var user model.User

	gorm := DB.First(&user, "email = ?", email)

	return user, gorm.Error
}

// Checks to see if the given email is in the user database.
// Params: email string of the user
// user *models.User to be updated with found user, if any
// Returns: true if email found, false otherwise
func IsEmailInUserDB(email string, user *model.User) bool {
	user, err := database.UserByEmail(email)

	return user.Id != 0 && err == nil
}

//***********************************UPDATE***********************************//

//***********************************DELETE***********************************//
