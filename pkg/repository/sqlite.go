package repository

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func NewSqlite(dbPath string) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(dbPath))
	if err != nil {
		return nil, err
	}
	return db, nil
}
