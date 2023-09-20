package service

import (
	"tinkapi/v2/pkg/repository"
	"tinkapi/v2/pkg/repository/entity"
)

type TeamService struct {
	repo repository.Team
}

func NewTeamService(repo repository.Team) *TeamService {
	return &TeamService{repo: repo}
}

func (s *TeamService) CreateFromDesc(desc entity.Desc) error {
	return s.repo.CreateFromDesc(desc)
}

func (s *TeamService) CreateTeam(team entity.Team) (string, error) {
	return s.repo.CreateTeam(team)
}
