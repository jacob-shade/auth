package router

import (
	"fmt"
	"go-sessions-authentication/config"
	"go-sessions-authentication/controller"
	"go-sessions-authentication/database"
	"go-sessions-authentication/model"
	"go-sessions-authentication/util"
	"strings"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

// AuthMiddleware checks that the user's session is logged in.
func AuthMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		sess, err := store.Get(c)

		// omits the use of the middleware for authentication routes
		if strings.Split(c.Path(), "/")[1] == "auth" {
			return c.Next()
		}

		if err != nil || sess.Get(config.Config("AUTH_KEY")) == nil {
			fmt.Println("middleware stopping exe")
			return util.NotAuthorized(c)
		}

		return c.Next()
	}
}

// Login attempts to login the user with the username and password given.
//
// Updates the session storage and client cookie.
func Login(c *fiber.Ctx) error {
	// parsing json from client
	var data map[string]string
	err := c.BodyParser(&data)
	if status := util.ErrorCheck(c, err); status != nil {
		return status
	}

	// checking if username is in db
	user, err := database.UserByUsername(data["username"])
	if user.Id == 0 {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "user not found: " + fmt.Sprint(err),
		})
	}

	// checking that password matches
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data["password"]))
	if err != nil { // not authorized
		return util.NotAuthorized(c)
	}

	// retrieves the session from the store
	sess, err := store.Get(c)
	if status := util.ErrorCheck(c, err); status != nil { // error occurred
		return status
	}
	sess.Set(config.Config("AUTH_KEY"), true)
	sess.Set(config.Config("USER_ID"), user.Id)

	// updating storage and client cookie with new session
	err = sess.Save()
	if status := util.ErrorCheck(c, err); status != nil { // error occurred
		return status
	}

	return util.StatusOK(c, "logged in")
}

// Logout logs out user, deletes session from storage, expires session cookie.
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

// checkAuthenticated verifies that the user is authorized.
func checkAuthenticated(c *fiber.Ctx) error {
	fmt.Printf("checking if authorized")
	sess, err := store.Get(c)
	if err != nil { // not authorized
		return util.NotAuthorized(c)
	}

	auth := sess.Get(config.Config("AUTH_KEY"))
	if auth != nil {
		return util.StatusOK(c, "authenticated")
	}
	return util.NotAuthorized(c)
}

// GetUser returns the JSON of the user.
func GetUser(c *fiber.Ctx) error {
	fmt.Println("in getuser")
	sess, err := store.Get(c)
	if err != nil { // not authorized
		return util.NotAuthorized(c)
	}
	fmt.Println("got session")

	auth := sess.Get(config.Config("AUTH_KEY"))
	if auth == nil {
		return util.NotAuthorized(c)
	}
	fmt.Println("got auth key")

	userId := sess.Get(config.Config("USER_ID"))
	if userId == nil { // not authorized
		return util.NotAuthorized(c)
	}
	fmt.Println("got userid")

	var user model.User
	user, err = controller.GetUser(fmt.Sprint(userId))
	if err != nil { // not authorized
		return util.NotAuthorized(c)
	}
	fmt.Println("got user")

	return c.Status(fiber.StatusOK).JSON(user)
}
