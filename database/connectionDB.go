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
	db     *gorm.DB                         // database orm
	models = append(make([]interface{}, 0), // all entity/relation models
		model.User{},
	)
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
	protocol := config.Config("PROTOCOL")
	address := config.Config("ADDRESS")
	port := config.Config("PORT")
	dbName := config.Config("DB_NAME")
	dsn := fmt.Sprintf("%v:%v@%v(%v:%v)/%v?charset=utf8&parseTime=True&loc=Local",
		username, password, protocol, address, port, dbName)

	// Connecting to database.
	var err error
	db, err = gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,   // data source name
		DefaultStringSize:         256,   // default size for string fields
		DisableDatetimePrecision:  true,  // disable datetime precision, which not supported before MySQL 5.6
		DontSupportRenameIndex:    true,  // drop & create when rename index, rename index not supported before MySQL 5.7, MariaDB
		DontSupportRenameColumn:   true,  // `change` when rename column, rename column not supported before MySQL 8, MariaDB
		SkipInitializeWithVersion: false, // auto configure based on currently MySQL version
	}), &gorm.Config{})
	util.PanicCheck(err, "failed to connect to database")
	fmt.Printf("Connected to %v database\n", dbName)

	// Auto migrating schema to keep up to date.
	for _, model := range models {
		db.AutoMigrate(model)
	}
	fmt.Println("Automigrating database schema")
}
