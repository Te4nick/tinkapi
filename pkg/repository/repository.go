package repository

import "gorm.io/gorm"

type Authorization interface {
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
