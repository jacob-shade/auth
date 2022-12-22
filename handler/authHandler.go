// Package handler provides functionality for routes.
//
// This package is implemented by manipulating the recieved "[Fiber]" Contexts.
//
// [Fiber]: https://pkg.go.dev/github.com/gofiber/fiber/v2@v2.40.1
package handler

import (
	"go-sessions-authentication/database"
	"go-sessions-authentication/model"
	"go-sessions-authentication/util"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

// Register allows a user to register with the service.
//
// Uses "[bcrypt]" to hash the password before storing in the database.
//
// [bcrypt]: https://pkg.go.dev/golang.org/x/crypto@v0.4.0/bcrypt
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

	err = database.UserCreate(&user)
	if status := util.ErrorCheck(c, err); status != nil { // error occurred
		return status
	}

	return util.StatusOK(c, "registered")
}
