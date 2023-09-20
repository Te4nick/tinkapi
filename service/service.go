package service

import (
	"tinkapi/v2/pkg/repository"
	"tinkapi/v2/pkg/repository/entity"
)

type Authorization interface {
	CreateUser(user entity.User) (string, error)
	ReadUser(name string) (entity.User, error)
	UpdateUser(user entity.User) (string, error)
	DeleteUser(name string) (string, error)
}

type Team interface {
	CreateTeam(team entity.Team) (string, error)
	CreateFromDesc(desc entity.Desc) error
}

type Duty interface {
	CreateDuty(duty entity.Duty) (string, error)
}

type Service struct {
	Authorization
	Team
	Duty
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		Team:          NewTeamService(repos.Team),
		Duty:          NewDutyService(repos.Duty),
	}
}
