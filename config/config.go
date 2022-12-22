// Package config provides a way to retrieve enviornment variables.
package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

// Config gets env value from key
func Config(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Print("Error loading .env file")
	}

	return os.Getenv(key)
}
