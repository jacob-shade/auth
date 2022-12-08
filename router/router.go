package router

import (
	"go-sessions-authentication/handler"
	"go-sessions-authentication/model"
	"go-sessions-authentication/services"
	"go-sessions-authentication/util"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/session"
	"golang.org/x/crypto/bcrypt"
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

	router.Post("/auth/register", handler.Register)
	router.Post("/auth/login", Login)
	//router.Post("/auth/logout", handler.Logout)
	//router.Get("/auth/healthcheck", handler.HealthCheck)

	//router.Get("/user", handler.GetUser)
}

func Login(c *fiber.Ctx) error {
	var data map[string]string

	err := c.BodyParser(&data)
	if status := util.ErrorCheck(c, err); status != nil { // error occurred
		return status
	}

	var user model.User
	if !services.IsEmailInUserDB(data["email"], &user) {
		return util.NotAuthorized(c)
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data.Password))
	if err != nil {
		return util.NotAuthorized(c)
	}

	sess, err := store.Get(c)
	if status := util.ErrorCheck(c, err); status != nil { // error occurred
		return status
	}

	sess.Get(AUTH_KEY)
	sess.Set(USER_ID, user.Id)

	err = sess.Save()
	if status := util.ErrorCheck(c, err); status != nil { // error occurred
		return status
	}

	return util.StatusOK(c, "logged in")
}

func Logout(c *fiber.Ctx) error {
	sess, err := store.Get(c)
	if err != nil { // error occurred
		return util.StatusOK(c, "logged out (no session)")
	}

	err = sess.Destroy()
	if status := util.ErrorCheck(c, err); status != nil { // error occurred
		return status
	}

	return util.StatusOK(c, "logged out")
}
