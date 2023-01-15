package controller

import (
	"fmt"
	"go-sessions-authentication/config"
	"go-sessions-authentication/database"
	"go-sessions-authentication/middleware"
	"go-sessions-authentication/model"
	"go-sessions-authentication/util"

	"github.com/gofiber/fiber/v2"
)

// GetUser returns the JSON of the user.
func GetUser(c *fiber.Ctx) error {
	sess, err := middleware.GetSession(c)
	if err != nil { // not authorized
		return util.NotAuthorized(c)
	}

	auth := sess.Get(config.Config("AUTH_KEY"))
	if auth == nil {
		return util.NotAuthorized(c)
	}

	userId := sess.Get(config.Config("USER_ID"))
	if userId == nil { // not authorized
		return util.NotAuthorized(c)
	}

	var user model.User
	user, err = database.UserById(fmt.Sprint(userId))
	if err != nil { // not authorized
		return util.NotAuthorized(c)
	}

	return c.Status(fiber.StatusOK).JSON(user)
}
