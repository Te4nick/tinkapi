package repository

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"tinkapi/v2/pkg/repository/entity"
)

const (
	usersTable        = "user"
	teamsTable        = "team"
	dutiesTable       = "duty"
	dutyRolePrimary   = "primary"
	dutyRoleSecondary = "secondary"
)

func NewSqlite(dbPath string) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(dbPath))
	if err != nil {
		return nil, err
	}
	if err = db.AutoMigrate(&entity.User{}); err != nil {
		return nil, err
	}
	if err = db.AutoMigrate(&entity.Team{}); err != nil {
		return nil, err
	}
	if err = db.AutoMigrate(&entity.Duty{}); err != nil {
		return nil, err
	}
	return db, nil
}
