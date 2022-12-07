package routes

import (
	"strings"

	"github.com/gofiber/fiber/v2"
)

func NewMiddleware() fiber.Handler {
	return AuthMiddleware
}

func AuthMiddleware(c *fiber.Ctx) error {
	sess, err := store.Get(c)

	// can modify later to only check for authorization
	// for pages necessary to be signed in
	if strings.Split(c.Path(), "/")[0] == "auth" {
		return c.Next()
	}

	if err != nil || sess.Get(AUTH_KEY) == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "not authorized",
		})
	}

	return c.Next()
}
