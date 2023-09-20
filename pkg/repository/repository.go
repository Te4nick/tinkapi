package repository

import (
	"gorm.io/gorm"
	"tinkapi/v2/pkg/repository/entity"
)

type Authorization interface {
	CreateUser(user entity.User) (string, error)
	ReadUser(name string) (entity.User, error)
	UpdateUser(user entity.User) (string, error)
	DeleteUser(name string) (string, error)
}

type Team interface {
}

type Duty interface {
}

type Repository struct {
	Authorization
	Team
	Duty
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		Authorization: NewAuthSqlite(db),
	}
}
