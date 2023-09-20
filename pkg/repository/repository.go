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
	CreateFromDesc(desc entity.Desc) error
	CreateTeam(team entity.Team) (string, error)
}

type Duty interface {
	CreateDuty(duty entity.Duty) (string, error)
}

type Repository struct {
	Authorization
	Team
	Duty
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		Authorization: NewAuthSqlite(db),
		Team:          NewTeamSqlite(db),
		Duty:          NewDutySqlite(db),
	}
}
