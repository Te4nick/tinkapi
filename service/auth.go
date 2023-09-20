package service

import (
	"tinkapi/v2/pkg/repository"
	"tinkapi/v2/pkg/repository/entity"
)

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user entity.User) (string, error) {
	return s.repo.CreateUser(user)
}
func (s *AuthService) ReadUser(name string) (entity.User, error) {
	return entity.User{}, nil
}
func (s *AuthService) UpdateUser(user entity.User) (string, error) {
	return "", nil
}
func (s *AuthService) DeleteUser(name string) (string, error) {
	return "", nil
}
