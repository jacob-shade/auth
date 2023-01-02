package controller

import (
	"go-sessions-authentication/database"
	"go-sessions-authentication/model"
)

// GetUser returns the user with the given id.
func GetUser(id string) (model.User, error) {
	return database.UserById(id)
}
