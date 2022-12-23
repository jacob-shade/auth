// Package util offers commonly used functions to be used among other packages.
package util

import "github.com/gofiber/fiber/v2"

// ErrorCheck returns JSON message and calls fiber.StatusInternalServerError
// if an error found, or nil on no error.
func ErrorCheck(c *fiber.Ctx, err error) error {
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "something went wrong: " + err.Error(),
		})
	}
	return nil
}

// PanicCheck calls a panic with the message given if there is an error.
func PanicCheck(err error, msg string) {
	if err != nil {
		panic(msg)
	}
}

// StatusOK returns JSON of the given message and calls fiber.StatusOK.
func StatusOK(c *fiber.Ctx, message string) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": message,
	})
}

// NotAuthorized returns JSON of the message not authorized and calls
// fiber.StatusUnauthorized
func NotAuthorized(c *fiber.Ctx) error {
	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		"message": "not authorized",
	})
}
