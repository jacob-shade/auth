package router

import (
	"fmt"
	"go-sessions-authentication/database"
	"go-sessions-authentication/handler"
	"go-sessions-authentication/model"
	"go-sessions-authentication/util"
	"strings"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func NewMiddleware() fiber.Handler {
	return AuthMiddleware
}

func AuthMiddleware(c *fiber.Ctx) error {
	sess, err := store.Get(c)

	// can modify later to only check for authorization
	// for pages necessary to be signed in
	if strings.Split(c.Path(), "/")[1] == "auth" {
		return c.Next()
	}

	if err != nil || sess.Get(AUTH_KEY) == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "not authorized",
		})
	}

	return c.Next()
}

func Login(c *fiber.Ctx) error {
	var data map[string]string

	err := c.BodyParser(&data)
	if status := util.ErrorCheck(c, err); status != nil { // error occurred
		return status
	}

	var user model.User
	err = database.UserByEmail(data["email"], &user)
	if err != nil { // not authorized
		return util.NotAuthorized(c)
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data["password"]))
	if err != nil { // not authorized
		return util.NotAuthorized(c)
	}

	sess, err := store.Get(c)
	if status := util.ErrorCheck(c, err); status != nil { // error occurred
		return status
	}

	sess.Get(AUTH_KEY)
	sess.Set(USER_ID, user.Id)

	err = sess.Save()
	if status := util.ErrorCheck(c, err); status != nil { // error occurred
		return status
	}

	return util.StatusOK(c, "logged in")
}

func Logout(c *fiber.Ctx) error {
	sess, err := store.Get(c)
	if err != nil { // error occurred
		return util.StatusOK(c, "logged out (no session)")
	}

	err = sess.Destroy()
	if status := util.ErrorCheck(c, err); status != nil { // error occurred
		return status
	}

	return util.StatusOK(c, "logged out")
}

func HealthCheck(c *fiber.Ctx) error {
	sess, err := store.Get(c)
	if err != nil { // not authorized
		return util.NotAuthorized(c)
	}

	auth := sess.Get(AUTH_KEY)
	if auth != nil {
		return util.StatusOK(c, "authenticated")
	}
	return util.NotAuthorized(c)
}

func GetUser(c *fiber.Ctx) error {
	sess, err := store.Get(c)
	if err != nil { // not authorized
		return util.NotAuthorized(c)
	}

	auth := sess.Get(AUTH_KEY)
	if auth != nil {
		return util.NotAuthorized(c)
	}

	userId := sess.Get(USER_ID)
	if userId != nil { // not authorized
		return util.NotAuthorized(c)
	}

	var user model.User
	user, err = handler.GetUser(fmt.Sprint(userId))
	if err != nil { // not authorized
		return util.NotAuthorized(c)
	}

	return c.Status(fiber.StatusOK).JSON(user)
}
