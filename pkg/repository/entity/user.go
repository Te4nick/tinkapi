package entity

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name        string `gorm:"primaryKey;unique" yaml:"name"`
	FullName    string `yaml:"full_name"`
	PhoneNumber string `yaml:"phone_number"`
	Email       string `yaml:"email"`
}
