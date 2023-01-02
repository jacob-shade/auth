// Package router implements a request router using the "[Fiber]" app.
//
// [Fiber]: https://pkg.go.dev/github.com/gofiber/fiber/v2@v2.40.1
package router

import (
	controller "go-sessions-authentication/controller"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/session"
)

var store *session.Store

// Setup sets up routes to use HTTP methods and calls the appropriate function
// to handle each route.
func Setup() {
	router := fiber.New()

	store = session.New(session.Config{
		CookieHTTPOnly: true,
		// CookieSecure: true, //for https
		// CookieDomain: ,
		// CookiePath: ,
		// CookieSameSite: ,
		Expiration: time.Hour * 5,
	})

	router.Use(AuthMiddleware(), cors.New(cors.Config{
		AllowCredentials: true,
		AllowOrigins:     "*", //so we can access from local host
		AllowHeaders:     "Access-Control-Allow-Origin, Content-Type, Origin, Accept",
	}))

	router.Post("/auth/register", controller.Register)
	router.Post("/auth/login", Login)
	router.Post("/auth/logout", Logout)
	router.Get("/auth/checkAuthenticated", checkAuthenticated)

	router.Get("/user", GetUser)

	router.Listen(":5000")
}
