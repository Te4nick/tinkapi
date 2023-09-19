package repository

import (
	_ "gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type AuthSqlite struct {
	db *gorm.DB
}

func NewAuthSqlite(db *gorm.DB) AuthSqlite {
	return AuthSqlite{db: db}
}

func (r *AuthSqlite) CreateUser() {

}
