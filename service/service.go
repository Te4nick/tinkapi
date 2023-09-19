package service

import "tinkapi/v2/pkg/repository"

type Authorization interface {
}

type Team interface {
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
