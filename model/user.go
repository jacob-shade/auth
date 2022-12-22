// Package model holds all types linked to the database schema.
package model

type User struct {
	Id       uint   `json:"id" gorm:"primarykey"`
	Name     string `json:"name"`
	Email    string `json:"email" gorm:"unique"`
	Password []byte `json:"-"`
}
