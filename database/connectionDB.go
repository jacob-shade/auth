package database

import (
	"fmt"
	"go-sessions-authentication/config"
	"go-sessions-authentication/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func Connect() {
	username := config.Config("DB_USERNAME")
	password := config.Config("DB_PASSWORD")
	serverIP := config.Config("SERVER_IP")
	dbName := config.Config("DB_NAME")
	dsn := fmt.Sprintf("%v:%v@%v/%v", username, password, serverIP, dbName)

	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}

	db.AutoMigrate(&model.User{}) //migrates schema to server
	fmt.Printf("Connected to %v database\n", dbName)
}
