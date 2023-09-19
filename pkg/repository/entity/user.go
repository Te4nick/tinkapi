package entity

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name        string `gorm:"primaryKey;unique"`
	FullName    string
	PhoneNumber string
	Email       string
}
