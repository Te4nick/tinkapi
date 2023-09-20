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
	CreateFromDesc(desc *entity.Desc) error
}

type Duty interface {
}

type Service struct {
	Authorization
	Team
	Duty
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}
