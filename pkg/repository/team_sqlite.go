package repository

import (
	"gorm.io/gorm"
	"time"
	"tinkapi/v2/pkg/repository/entity"
)

type TeamSqlite struct {
	db *gorm.DB
}

func NewTeamSqlite(db *gorm.DB) *TeamSqlite {
	return &TeamSqlite{db: db}
}

func (r *TeamSqlite) CreateFromDesc(desc entity.Desc) error {
	var teams []entity.Team
	var users []entity.User
	var duties []entity.Duty
	for _, team := range desc.Teams {
		dbTeam := entity.Team{
			Name:               team.Name,
			SchedulingTimezone: team.SchedulingTimezone,
			Email:              team.Email,
			SlackChannel:       team.SlackChannel,
		}
		teams = append(teams, dbTeam)
		for _, user := range team.Users {
			dbUser := entity.User{
				Name:        user.Name,
				FullName:    user.FullName,
				PhoneNumber: user.PhoneNumber,
				Email:       user.Email,
			}
			users = append(users, dbUser)
			for _, duty := range user.Duty {
				dbDate, _ := time.Parse("02/01/2006", duty.Date)
				dbDuty := entity.Duty{
					User: dbUser,
					Team: dbTeam,
					Date: dbDate,
					Role: duty.Role,
				}
				duties = append(duties, dbDuty)
			}
		}
	}
	return nil
}

func (r *TeamSqlite) CreateTeam(team entity.Team) (string, error) {
	result := r.db.Create(&team)
	if result.Error != nil {
		return "", result.Error
	}
	return team.Name, result.Error
}
