// Package router implements a request router using the [Fiber] app.
//
// [Fiber]: https://pkg.go.dev/github.com/gofiber/fiber/v2@v2.40.1
package router

import (
	"go-sessions-authentication/controller"

	"github.com/gofiber/fiber/v2"
)

// Setup sets up routes to use HTTP methods and calls the appropriate function
// to handle each route.
func Setup(app *fiber.App) {
	api := app.Group("/api")
	api.Get("/user", controller.GetUser)

	auth := api.Group("/auth")
	auth.Post("/register", controller.Register)
	auth.Post("/login", controller.Login)
	auth.Post("/logout", controller.Logout)
	auth.Get("/loginStatus", controller.LoginStatus)
}
