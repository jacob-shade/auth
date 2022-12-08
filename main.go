package main

import (
	"go-sessions-authentication/database"
	"go-sessions-authentication/router"
)

func main() {
	database.Connect()

	//app.Use(middleware.Logger())

	router.Setup()

	//listen
}
