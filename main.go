package main

import (
	"go-sessions-authentication/database"
	"go-sessions-authentication/middleware"
	"go-sessions-authentication/router"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	database.Connect()

	middleware.Setup(app)

	router.Setup(app)

	log.Fatal(app.Listen(":5000"))
}
