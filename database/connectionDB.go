package database

import (
	"fmt"
	"go-sessions-authentication/config"
	"go-sessions-authentication/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	// dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"

	username := config.Config("DB_USERNAME")
	password := config.Config("DB_PASSWORD")
	serverIP := config.Config("SERVER_IP")
	dbName := config.Config("DB_NAME")
	dsn := fmt.Sprintf("%v:%v@%v/%v", username, password, serverIP, dbName)

	DB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}

	DB.AutoMigrate(&model.User{}) //migrates schema to server
	fmt.Printf("Connected to %v database\n", dbName)
}
