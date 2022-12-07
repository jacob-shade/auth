package database

import (
	"fmt"
	"go-sessions-authentication/models"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	//dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	envMap, envErr := godotenv.Read(".env")
	if envErr != nil {
		panic("error loading .env into map[string]string")
	}

	username := envMap["DB_USERNAME"]
	password := envMap["DB_PASSWORD"]
	serverIP := envMap["SERVER_IP"]
	dbName := envMap["DB_NAME"]
	dsn := username + ":" + password + "@" + serverIP + "/" + dbName
	DB, conErr := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if conErr != nil {
		panic("failed to connect to database")
	}

	DB.AutoMigrate(&models.User{}) //migrates schema to server
	fmt.Printf("Connected to %v database\n", dbName)
}
