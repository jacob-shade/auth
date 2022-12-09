package router

import (
	"go-sessions-authentication/handler"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/session"
)

var store *session.Store

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

	router.Use(NewMiddleware(), cors.New(cors.Config{
		AllowCredentials: true,
		AllowOrigins:     "*", //so we can access from local host
		AllowHeaders:     "Access-Control-Allow-Origin, Content-Type, Origin, Accept",
	}))

	router.Post("/auth/register", handler.Register)
	router.Post("/auth/login", Login)
	router.Post("/auth/logout", Logout)
	router.Get("/auth/healthcheck", HealthCheck)

	router.Get("/user", GetUser)

	router.Listen(":5000")
}
