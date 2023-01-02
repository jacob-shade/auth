// Package model holds all types linked to the database schema.
package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Id       uint   `json:"id" gorm:"primarykey"`
	Username string `json:"name"`
	Email    string `json:"email" gorm:"unique"`
	Password []byte `json:"-"`
}
