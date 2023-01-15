package middleware

import (
	"fmt"
	"go-sessions-authentication/config"
	"go-sessions-authentication/util"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/session"
)

var store *session.Store

func Setup(app *fiber.App) {
	store = session.New(
		session.Config{
			CookieHTTPOnly: true,
			// CookieSecure: true, //for https
			Expiration: time.Hour * 24 * 30, // 1 Week
			// Storage interface to store the session data
			// Optional. Default value memory.New()
			// Storage fiber.Storage
			// storage := sqlite3.New() // From github.com/gofiber/storage/sqlite3
			// store := session.New(session.Config{
			//   Storage: storage,
			// })
		},
	)

	app.Use(authMiddleware(),
		cors.New(cors.Config{
			AllowOrigins:     config.Config("DOMAIN"),
			AllowHeaders:     "Access-Control-Allow-Origin, Content-Type, Origin, Accept",
			AllowCredentials: true,
		}),
	)

	// app.Use(AuthMiddleware(),
	// 	csrf.New(csrf.Config{
	// 		CookieDomain:   config.Config("DOMAIN"),
	// 		CookieHTTPOnly: true,
	// 		Expiration:     time.Hour * 24 * 30,
	// 		//Storage
	// 	}),
	// )
}

func GetSession(c *fiber.Ctx) (*session.Session, error) {
	sess, err := store.Get(c)
	return sess, err
}

// AuthMiddleware checks that the user's session is logged in.
func authMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		sess, err := store.Get(c)

		// omits the use of the middleware for authentication routes
		if strings.Split(c.Path(), "/")[2] == "auth" {
			return c.Next()
		}

		if err != nil || sess.Get(config.Config("AUTH_KEY")) == nil {
			fmt.Println("middleware stopping exe")
			return util.NotAuthorized(c)
		}

		return c.Next()
	}
}
