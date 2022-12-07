package controllers

import (
	"go-sessions-authentication/database"
	"go-sessions-authentication/models"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *fiber.Ctx) error {
	var data map[string]string

	err := c.BodyParser(&data)
	if status := errorCheck(c, err); status != nil { // error occurred
		return status
	}

	password, err := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)
	if status := errorCheck(c, err); status != nil { // error occurred
		return status
	}

	user := models.User{
		Name:     data["name"],
		Email:    data["email"],
		Password: password,
	}

	err = database.CreateUser(&user)
	if status := errorCheck(c, err); status != nil { // error occurred
		return status
	}

	return statusOK(c, "registered")
}
