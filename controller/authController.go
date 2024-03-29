// Package handler provides functionality for routes.
//
// This package is implemented by manipulating the recieved [Fiber] Contexts.
//
// [Fiber]: https://pkg.go.dev/github.com/gofiber/fiber/v2@v2.40.1
package controller

import (
	"fmt"
	"go-sessions-authentication/config"
	"go-sessions-authentication/database"
	"go-sessions-authentication/middleware"
	"go-sessions-authentication/model"
	"go-sessions-authentication/util"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

// Login attempts to login the user with the username and password given.
//
// Updates the session storage and client cookie.
func Login(c *fiber.Ctx) error {
	// parsing json from client
	var data map[string]string
	err := c.BodyParser(&data)
	if status := util.ErrorCheck(c, err); status != nil {
		return status
	}

	// checking if username is in db
	user, err := database.UserByUsername(data["username"])
	if user.Id == 0 {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "user not found: " + fmt.Sprint(err),
		})
	}

	// checking that password matches
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data["password"]))
	if err != nil { // not authorized
		return util.NotAuthorized(c)
	}

	// retrieves the session from the store
	sess, err := middleware.GetSession(c)
	if status := util.ErrorCheck(c, err); status != nil { // error occurred
		return status
	}
	sess.Set(config.Config("AUTH_KEY"), true)
	sess.Set(config.Config("USER_ID"), user.Id)

	// updating storage and client cookie with new session
	err = sess.Save()
	if status := util.ErrorCheck(c, err); status != nil { // error occurred
		return status
	}

	return util.StatusOK(c, "logged in")
}

// Logout logs out user, deletes session from storage, expires session cookie.
func Logout(c *fiber.Ctx) error {
	sess, err := middleware.GetSession(c)
	if err != nil {
		return util.StatusOK(c, "logged out (no session)")
	}

	err = sess.Destroy()
	if status := util.ErrorCheck(c, err); status != nil {
		return status
	}

	return util.StatusOK(c, "logged out")
}

// LoginStatus verifies that the user is authorized.
func LoginStatus(c *fiber.Ctx) error {
	sess, err := middleware.GetSession(c)
	if err != nil { // not authorized
		return util.NotAuthorized(c)
	}

	auth := sess.Get(config.Config("AUTH_KEY"))
	if auth != nil {
		return util.StatusOK(c, "authenticated")
	}
	return util.NotAuthorized(c)
}

// Register allows a user to register with the service.
//
// Uses [bcrypt] to hash the password before storing in the database.
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
		Username: data["username"],
		Email:    data["email"],
		Password: password,
	}

	err = database.UserCreate(&user)
	if status := util.ErrorCheck(c, err); status != nil { // error occurred
		return status
	}

	return util.StatusOK(c, "registered")
}
