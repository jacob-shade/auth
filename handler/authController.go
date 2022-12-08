package handler

import (
	"go-sessions-authentication/model"
	"go-sessions-authentication/services"
	"go-sessions-authentication/util"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

// Allows a user to
func Register(c *fiber.Ctx) error {
	var data map[string]string

	err := c.BodyParser(&data)
	status := util.ErrorCheck(c, err)
	if status != nil { // error occurred
		return status
	}

	password, err := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)
	if status := util.ErrorCheck(c, err); status != nil { // error occurred
		return status
	}

	user := model.User{
		Name:     data["name"],
		Email:    data["email"],
		Password: password,
	}

	err = services.RegisterUser(&user)
	if status := util.ErrorCheck(c, err); status != nil { // error occurred
		return status
	}

	return util.StatusOK(c, "registered")
}
