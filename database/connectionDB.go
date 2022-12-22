// Package database allows services or handlers to connect and manipulate
// mySQL databases.
//
// Allows for:
//  1. Database connection
//  2. User CRUD
//
// This package is implemented using the "[GORM]" ORM.
//
// [GORM]: https://gorm.io/docs/
package database

import (
	"fmt"
	"go-sessions-authentication/config"
	"go-sessions-authentication/model"
	"go-sessions-authentication/util"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db     *gorm.DB         // database orm
	models = []interface{}{ // slice of all entity/relation models
		model.User{},
	}
)

// Connect connects to the database with the config given in the .env file and
// the models given in the database package.
//
// For details on dsn(data source name) formating, refer to [go-sql-driver]
// docs.
//
// [go-sql-driver]: https://github.com/go-sql-driver/mysql#dsn-data-source-name
func Connect() {
	// Loading in database enviornment variables.
	username := config.Config("DB_USERNAME")
	password := config.Config("DB_PASSWORD")
	serverIP := config.Config("SERVER_IP")
	dbName := config.Config("DB_NAME")
	dsn := fmt.Sprintf("%v:%v@%v/%v", username, password, serverIP, dbName)

	// Connecting to database.
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	util.PanicCheck(err, "failed to connect to database")
	fmt.Printf("Connected to %v database\n", dbName)

	// Auto migrating schema to keep up to date.
	db.AutoMigrate(models)
}
