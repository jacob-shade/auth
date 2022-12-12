package util

import "github.com/gofiber/fiber/v2"

func ErrorCheck(c *fiber.Ctx, err error) error {
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "something went wrong: " + err.Error(),
		})
	}
	return nil
}

func PanicCheck(err error, msg string) {
	if err != nil {
		panic(msg)
	}
}

func StatusOK(c *fiber.Ctx, message string) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": message,
	})
}

func NotAuthorized(c *fiber.Ctx) error {
	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		"message": "not authorized",
	})
}
