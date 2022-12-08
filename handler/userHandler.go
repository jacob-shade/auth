package handler

import (
	"go-sessions-authentication/database"
	"go-sessions-authentication/model"
)

func GetUser(id string) (model.User, error) {
	return database.UserById(id)
}
