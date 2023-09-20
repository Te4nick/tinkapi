package repository

import (
	"gorm.io/gorm"
	"tinkapi/v2/pkg/repository/entity"
)

type DutySqlite struct {
	db *gorm.DB
}

func NewDutySqlite(db *gorm.DB) *DutySqlite {
	return &DutySqlite{db: db}
}

func (r *DutySqlite) CreateDuty(duty entity.Duty) (string, error) {
	result := r.db.Create(&duty)
	if result.Error != nil {
		return "", result.Error
	}
	return duty.User.Name, result.Error
}
