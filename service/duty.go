package service

import (
	"tinkapi/v2/pkg/repository"
	"tinkapi/v2/pkg/repository/entity"
)

type DutyService struct {
	repo repository.Duty
}

func NewDutyService(repo repository.Duty) *DutyService {
	return &DutyService{repo: repo}
}

func (s *DutyService) CreateDuty(duty entity.Duty) (string, error) {
	return s.repo.CreateDuty(duty)
}
