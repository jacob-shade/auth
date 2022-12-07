package routes

import (
	"go-sessions-authentication/controllers"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/session"
)

var (
	store    *session.Store
	AUTH_KEY string = "authenticated"
	USER_ID  string = "user_id"
)

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

	router.Post("/auth/register", controllers.Register)
	router.Post("/auth/login", controllers.Login)
	router.Post("/auth/logout", controllers.Logout)
	router.Get("/auth/healthcheck", controllers.HealthCheck)

	router.Get("/user", controllers.GetUser)
}
